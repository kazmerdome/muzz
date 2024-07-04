// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	decision "github.com/kazmerdome/muzz/internal/module/decision"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// DecisionRepository is an autogenerated mock type for the DecisionRepository type
type DecisionRepository struct {
	mock.Mock
}

type DecisionRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *DecisionRepository) EXPECT() *DecisionRepository_Expecter {
	return &DecisionRepository_Expecter{mock: &_m.Mock}
}

// Count provides a mock function with given fields: ctx, where
func (_m *DecisionRepository) Count(ctx context.Context, where *decision.WhereDto) (int64, error) {
	ret := _m.Called(ctx, where)

	if len(ret) == 0 {
		panic("no return value specified for Count")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *decision.WhereDto) (int64, error)); ok {
		return rf(ctx, where)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *decision.WhereDto) int64); ok {
		r0 = rf(ctx, where)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *decision.WhereDto) error); ok {
		r1 = rf(ctx, where)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DecisionRepository_Count_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Count'
type DecisionRepository_Count_Call struct {
	*mock.Call
}

// Count is a helper method to define mock.On call
//   - ctx context.Context
//   - where *decision.WhereDto
func (_e *DecisionRepository_Expecter) Count(ctx interface{}, where interface{}) *DecisionRepository_Count_Call {
	return &DecisionRepository_Count_Call{Call: _e.mock.On("Count", ctx, where)}
}

func (_c *DecisionRepository_Count_Call) Run(run func(ctx context.Context, where *decision.WhereDto)) *DecisionRepository_Count_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*decision.WhereDto))
	})
	return _c
}

func (_c *DecisionRepository_Count_Call) Return(_a0 int64, _a1 error) *DecisionRepository_Count_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DecisionRepository_Count_Call) RunAndReturn(run func(context.Context, *decision.WhereDto) (int64, error)) *DecisionRepository_Count_Call {
	_c.Call.Return(run)
	return _c
}

// GetOneByActorUserId provides a mock function with given fields: ctx, actorUserID
func (_m *DecisionRepository) GetOneByActorUserId(ctx context.Context, actorUserID uuid.UUID) (*decision.Decision, error) {
	ret := _m.Called(ctx, actorUserID)

	if len(ret) == 0 {
		panic("no return value specified for GetOneByActorUserId")
	}

	var r0 *decision.Decision
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*decision.Decision, error)); ok {
		return rf(ctx, actorUserID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *decision.Decision); ok {
		r0 = rf(ctx, actorUserID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*decision.Decision)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, actorUserID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DecisionRepository_GetOneByActorUserId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOneByActorUserId'
type DecisionRepository_GetOneByActorUserId_Call struct {
	*mock.Call
}

// GetOneByActorUserId is a helper method to define mock.On call
//   - ctx context.Context
//   - actorUserID uuid.UUID
func (_e *DecisionRepository_Expecter) GetOneByActorUserId(ctx interface{}, actorUserID interface{}) *DecisionRepository_GetOneByActorUserId_Call {
	return &DecisionRepository_GetOneByActorUserId_Call{Call: _e.mock.On("GetOneByActorUserId", ctx, actorUserID)}
}

func (_c *DecisionRepository_GetOneByActorUserId_Call) Run(run func(ctx context.Context, actorUserID uuid.UUID)) *DecisionRepository_GetOneByActorUserId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *DecisionRepository_GetOneByActorUserId_Call) Return(_a0 *decision.Decision, _a1 error) *DecisionRepository_GetOneByActorUserId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DecisionRepository_GetOneByActorUserId_Call) RunAndReturn(run func(context.Context, uuid.UUID) (*decision.Decision, error)) *DecisionRepository_GetOneByActorUserId_Call {
	_c.Call.Return(run)
	return _c
}

// GetOneByRecipientUserID provides a mock function with given fields: ctx, recipientUserID
func (_m *DecisionRepository) GetOneByRecipientUserID(ctx context.Context, recipientUserID uuid.UUID) (*decision.Decision, error) {
	ret := _m.Called(ctx, recipientUserID)

	if len(ret) == 0 {
		panic("no return value specified for GetOneByRecipientUserID")
	}

	var r0 *decision.Decision
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*decision.Decision, error)); ok {
		return rf(ctx, recipientUserID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *decision.Decision); ok {
		r0 = rf(ctx, recipientUserID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*decision.Decision)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, recipientUserID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DecisionRepository_GetOneByRecipientUserID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOneByRecipientUserID'
type DecisionRepository_GetOneByRecipientUserID_Call struct {
	*mock.Call
}

// GetOneByRecipientUserID is a helper method to define mock.On call
//   - ctx context.Context
//   - recipientUserID uuid.UUID
func (_e *DecisionRepository_Expecter) GetOneByRecipientUserID(ctx interface{}, recipientUserID interface{}) *DecisionRepository_GetOneByRecipientUserID_Call {
	return &DecisionRepository_GetOneByRecipientUserID_Call{Call: _e.mock.On("GetOneByRecipientUserID", ctx, recipientUserID)}
}

func (_c *DecisionRepository_GetOneByRecipientUserID_Call) Run(run func(ctx context.Context, recipientUserID uuid.UUID)) *DecisionRepository_GetOneByRecipientUserID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *DecisionRepository_GetOneByRecipientUserID_Call) Return(_a0 *decision.Decision, _a1 error) *DecisionRepository_GetOneByRecipientUserID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DecisionRepository_GetOneByRecipientUserID_Call) RunAndReturn(run func(context.Context, uuid.UUID) (*decision.Decision, error)) *DecisionRepository_GetOneByRecipientUserID_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, where
func (_m *DecisionRepository) List(ctx context.Context, where *decision.WhereDto) ([]decision.Decision, error) {
	ret := _m.Called(ctx, where)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []decision.Decision
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *decision.WhereDto) ([]decision.Decision, error)); ok {
		return rf(ctx, where)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *decision.WhereDto) []decision.Decision); ok {
		r0 = rf(ctx, where)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]decision.Decision)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *decision.WhereDto) error); ok {
		r1 = rf(ctx, where)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DecisionRepository_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type DecisionRepository_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - where *decision.WhereDto
func (_e *DecisionRepository_Expecter) List(ctx interface{}, where interface{}) *DecisionRepository_List_Call {
	return &DecisionRepository_List_Call{Call: _e.mock.On("List", ctx, where)}
}

func (_c *DecisionRepository_List_Call) Run(run func(ctx context.Context, where *decision.WhereDto)) *DecisionRepository_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*decision.WhereDto))
	})
	return _c
}

func (_c *DecisionRepository_List_Call) Return(_a0 []decision.Decision, _a1 error) *DecisionRepository_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DecisionRepository_List_Call) RunAndReturn(run func(context.Context, *decision.WhereDto) ([]decision.Decision, error)) *DecisionRepository_List_Call {
	_c.Call.Return(run)
	return _c
}

// UpsertOne provides a mock function with given fields: ctx, dto
func (_m *DecisionRepository) UpsertOne(ctx context.Context, dto decision.UpsertOneDto) (*decision.Decision, error) {
	ret := _m.Called(ctx, dto)

	if len(ret) == 0 {
		panic("no return value specified for UpsertOne")
	}

	var r0 *decision.Decision
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, decision.UpsertOneDto) (*decision.Decision, error)); ok {
		return rf(ctx, dto)
	}
	if rf, ok := ret.Get(0).(func(context.Context, decision.UpsertOneDto) *decision.Decision); ok {
		r0 = rf(ctx, dto)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*decision.Decision)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, decision.UpsertOneDto) error); ok {
		r1 = rf(ctx, dto)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DecisionRepository_UpsertOne_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpsertOne'
type DecisionRepository_UpsertOne_Call struct {
	*mock.Call
}

// UpsertOne is a helper method to define mock.On call
//   - ctx context.Context
//   - dto decision.UpsertOneDto
func (_e *DecisionRepository_Expecter) UpsertOne(ctx interface{}, dto interface{}) *DecisionRepository_UpsertOne_Call {
	return &DecisionRepository_UpsertOne_Call{Call: _e.mock.On("UpsertOne", ctx, dto)}
}

func (_c *DecisionRepository_UpsertOne_Call) Run(run func(ctx context.Context, dto decision.UpsertOneDto)) *DecisionRepository_UpsertOne_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(decision.UpsertOneDto))
	})
	return _c
}

func (_c *DecisionRepository_UpsertOne_Call) Return(_a0 *decision.Decision, _a1 error) *DecisionRepository_UpsertOne_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DecisionRepository_UpsertOne_Call) RunAndReturn(run func(context.Context, decision.UpsertOneDto) (*decision.Decision, error)) *DecisionRepository_UpsertOne_Call {
	_c.Call.Return(run)
	return _c
}

// NewDecisionRepository creates a new instance of DecisionRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDecisionRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *DecisionRepository {
	mock := &DecisionRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
