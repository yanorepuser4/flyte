package entrypoints

import (
	"context"
	"fmt"
	"runtime/debug"

	"github.com/flyteorg/flyteadmin/pkg/repositories"
	"github.com/flyteorg/flyteadmin/pkg/repositories/errors"

	"github.com/flyteorg/flyteadmin/pkg/common"
	"github.com/flyteorg/flyteadmin/pkg/runtime"
	"github.com/flyteorg/flyteadmin/scheduler"
	"github.com/flyteorg/flyteidl/clients/go/admin"
	"github.com/flyteorg/flytestdlib/contextutils"
	"github.com/flyteorg/flytestdlib/logger"
	"github.com/flyteorg/flytestdlib/profutils"
	"github.com/flyteorg/flytestdlib/promutils"
	"github.com/flyteorg/flytestdlib/promutils/labeled"

	"github.com/spf13/cobra"
	_ "gorm.io/driver/postgres" // Required to import database driver.
)

var schedulerRunCmd = &cobra.Command{
	Use:   "run",
	Short: "This command will start the flyte native scheduler and periodically get new schedules from the db for scheduling",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		configuration := runtime.NewConfigurationProvider()
		applicationConfiguration := configuration.ApplicationConfiguration().GetTopLevelConfig()
		schedulerConfiguration := configuration.ApplicationConfiguration().GetSchedulerConfig()

		// Define the schedulerScope for prometheus metrics
		schedulerScope := promutils.NewScope(applicationConfiguration.MetricsScope).NewSubScope("flytescheduler")
		schedulerPanics := schedulerScope.MustNewCounter("initialization_panic",
			"panics encountered initializing the flyte native scheduler")

		defer func() {
			if err := recover(); err != nil {
				schedulerPanics.Inc()
				logger.Fatalf(ctx, fmt.Sprintf("caught panic: %v [%+v]", err, string(debug.Stack())))
			}
		}()

		databaseConfig := configuration.ApplicationConfiguration().GetDbConfig()
		logConfig := logger.GetConfig()

		db, err := repositories.GetDB(ctx, databaseConfig, logConfig)
		if err != nil {
			logger.Fatal(ctx, err)
		}
		dbScope := schedulerScope.NewSubScope("database")
		repo := repositories.NewGormRepo(
			db, errors.NewPostgresErrorTransformer(schedulerScope.NewSubScope("errors")), dbScope)

		clientSet, err := admin.ClientSetBuilder().WithConfig(admin.GetConfig(ctx)).Build(ctx)
		if err != nil {
			logger.Fatalf(ctx, "Flyte native scheduler failed to start due to %v", err)
			return err
		}
		adminServiceClient := clientSet.AdminClient()

		scheduleExecutor := scheduler.NewScheduledExecutor(repo,
			configuration.ApplicationConfiguration().GetSchedulerConfig().GetWorkflowExecutorConfig(), schedulerScope, adminServiceClient)

		logger.Info(ctx, "Successfully initialized a native flyte scheduler")

		// Serve profiling endpoints.
		go func() {
			err := profutils.StartProfilingServerWithDefaultHandlers(
				ctx, schedulerConfiguration.ProfilerPort.Port, nil)
			if err != nil {
				logger.Panicf(ctx, "Failed to Start profiling and Metrics server. Error, %v", err)
			}
		}()

		err = scheduleExecutor.Run(ctx)
		if err != nil {
			logger.Fatalf(ctx, "Flyte native scheduler failed to start due to %v", err)
			return err
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(schedulerRunCmd)

	// Set Keys
	labeled.SetMetricKeys(contextutils.AppNameKey, contextutils.ProjectKey, contextutils.DomainKey,
		contextutils.ExecIDKey, contextutils.WorkflowIDKey, contextutils.NodeIDKey, contextutils.TaskIDKey,
		contextutils.TaskTypeKey, common.RuntimeTypeKey, common.RuntimeVersionKey)
}
