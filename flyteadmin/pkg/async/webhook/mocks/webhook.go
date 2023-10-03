// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	admin "github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/admin"

	interfaces "github.com/flyteorg/flyteadmin/pkg/runtime/interfaces"

	mock "github.com/stretchr/testify/mock"
)

// Webhook is an autogenerated mock type for the Webhook type
type Webhook struct {
	mock.Mock
}

type Webhook_GetConfig struct {
	*mock.Call
}

func (_m Webhook_GetConfig) Return(_a0 interfaces.WebHookConfig) *Webhook_GetConfig {
	return &Webhook_GetConfig{Call: _m.Call.Return(_a0)}
}

func (_m *Webhook) OnGetConfig() *Webhook_GetConfig {
	c_call := _m.On("GetConfig")
	return &Webhook_GetConfig{Call: c_call}
}

func (_m *Webhook) OnGetConfigMatch(matchers ...interface{}) *Webhook_GetConfig {
	c_call := _m.On("GetConfig", matchers...)
	return &Webhook_GetConfig{Call: c_call}
}

// GetConfig provides a mock function with given fields:
func (_m *Webhook) GetConfig() interfaces.WebHookConfig {
	ret := _m.Called()

	var r0 interfaces.WebHookConfig
	if rf, ok := ret.Get(0).(func() interfaces.WebHookConfig); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(interfaces.WebHookConfig)
	}

	return r0
}

type Webhook_Post struct {
	*mock.Call
}

func (_m Webhook_Post) Return(_a0 error) *Webhook_Post {
	return &Webhook_Post{Call: _m.Call.Return(_a0)}
}

func (_m *Webhook) OnPost(ctx context.Context, payload admin.WebhookPayload) *Webhook_Post {
	c_call := _m.On("Post", ctx, payload)
	return &Webhook_Post{Call: c_call}
}

func (_m *Webhook) OnPostMatch(matchers ...interface{}) *Webhook_Post {
	c_call := _m.On("Post", matchers...)
	return &Webhook_Post{Call: c_call}
}

// Post provides a mock function with given fields: ctx, payload
func (_m *Webhook) Post(ctx context.Context, payload admin.WebhookPayload) error {
	ret := _m.Called(ctx, payload)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, admin.WebhookPayload) error); ok {
		r0 = rf(ctx, payload)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}