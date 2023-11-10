package db

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/jackc/pgx/v5/pgtype"
	"gorm.io/gorm"
)

var Migrations = []*gormigrate.Migration{
	{
		ID: "2023-10-12-inits",
		Migrate: func(tx *gorm.DB) error {
			type ArtifactKey struct {
				gorm.Model
				Project string `gorm:"uniqueIndex:idx_pdn;index:idx_proj;type:varchar(64)"`
				Domain  string `gorm:"uniqueIndex:idx_pdn;index:idx_dom;type:varchar(64)"`
				Name    string `gorm:"uniqueIndex:idx_pdn;index:idx_name;type:varchar(255)"`
			}
			type Artifact struct {
				gorm.Model
				ArtifactKeyID uint           `gorm:"uniqueIndex:idx_pdnv"`
				ArtifactKey   ArtifactKey    `gorm:"foreignKey:ArtifactKeyID;references:ID"`
				Version       string         `gorm:"type:varchar(255);index:idx_artifact_version;uniqueIndex:idx_pdnv"`
				Partitions    *pgtype.Hstore `gorm:"type:hstore;index:idx_artifact_partitions"`

				LiteralType  []byte `gorm:"not null"`
				LiteralValue []byte `gorm:"not null"`

				Description           string `gorm:"type:varchar(255)"`
				MetadataType          string `gorm:"type:varchar(64)"`
				OffloadedUserMetadata string `gorm:"type:varchar(255)"`

				// Project/Domain assumed to always be the same as the Artifact
				ExecutionName   string `gorm:"type:varchar(255)"`
				WorkflowProject string `gorm:"type:varchar(64)"`
				WorkflowDomain  string `gorm:"type:varchar(64)"`
				WorkflowName    string `gorm:"type:varchar(255)"`
				WorkflowVersion string `gorm:"type:varchar(255)"`
				TaskProject     string `gorm:"type:varchar(64)"`
				TaskDomain      string `gorm:"type:varchar(64)"`
				TaskName        string `gorm:"type:varchar(255)"`
				TaskVersion     string `gorm:"type:varchar(255)"`
				NodeID          string `gorm:"type:varchar(64)"`
				// See Admin migration for note.
				// Here nullable in the case of workflow output.
				RetryAttempt *uint32
			}
			return tx.AutoMigrate(
				&ArtifactKey{}, &Artifact{},
			)
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(
				"artifacts", "artifact_keys",
			)
		},
	},
	{
		ID: "2023-10-22-trigger",
		Migrate: func(tx *gorm.DB) error {
			type TriggerKey struct {
				gorm.Model
				Project string        `gorm:"index:idx_t_pdn;index:idx_t_proj;type:varchar(64)"`
				Domain  string        `gorm:"index:idx_t_pdn;index:idx_t_dom;type:varchar(64)"`
				Name    string        `gorm:"index:idx_t_pdn;index:idx_t_name;type:varchar(255)"`
				RunsOn  []ArtifactKey `gorm:"many2many:active_trigger_artifact_keys;"`
			}

			type LaunchPlanID struct {
				Name    string `gorm:"not null;index:idx_lp_id;type:varchar(255)"`
				Version string `gorm:"not null;type:varchar(255);index:idx_launch_plan_version"`
			}

			type Trigger struct {
				gorm.Model
				TriggerKeyID uint       `gorm:"uniqueIndex:idx_trigger_pdnv"`
				TriggerKey   TriggerKey `gorm:"foreignKey:TriggerKeyID;references:ID"`
				Version      string     `gorm:"not null;type:varchar(255);index:idx_trigger_version;uniqueIndex:idx_trigger_pdnv"`

				// Unlike the one in the TriggerKey table, these are the list of artifact keys as specified by the user
				// for this specific version. Currently just the key but can add additional fields in the future.
				RunsOn []ArtifactKey `gorm:"many2many:trigger_ids_artifact_keys;"`

				Active            bool         `gorm:"index:idx_t_active"`
				LaunchPlanID      LaunchPlanID `gorm:"embedded"`
				LaunchPlanSpec    []byte       `gorm:"not null"`
				LaunchPlanClosure []byte       `gorm:"not null"`
			}
			return tx.AutoMigrate(
				&TriggerKey{}, &Trigger{},
			)
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(
				"triggers", "trigger_keys",
			)
		},
	},
}