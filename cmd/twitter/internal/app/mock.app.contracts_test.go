// Code generated by MockGen. DO NOT EDIT.
// Source: contracts.go
//
// Generated by this command:
//
//	mockgen -source=contracts.go -destination mock.app.contracts_test.go -package app_test
//

// Package app_test is a generated GoMock package.
package app_test

import (
	context "context"
	reflect "reflect"

	app "github.com/ZergsLaw/back-template/cmd/twitter/internal/app"
	uuid "github.com/gofrs/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance.
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// ByAuthor mocks base method.
func (m *MockRepo) ByAuthor(arg0 context.Context, arg1 uuid.UUID, arg2, arg3 int) ([]app.Twit, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ByAuthor", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]app.Twit)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ByAuthor indicates an expected call of ByAuthor.
func (mr *MockRepoMockRecorder) ByAuthor(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ByAuthor", reflect.TypeOf((*MockRepo)(nil).ByAuthor), arg0, arg1, arg2, arg3)
}

// Create mocks base method.
func (m *MockRepo) Create(arg0 context.Context, arg1 app.Twit) (*app.Twit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*app.Twit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRepoMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepo)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockRepo) Delete(ctx context.Context, ID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRepoMockRecorder) Delete(ctx, ID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepo)(nil).Delete), ctx, ID)
}

// Update mocks base method.
func (m *MockRepo) Update(ctx context.Context, twit app.Twit) (*app.Twit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, twit)
	ret0, _ := ret[0].(*app.Twit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockRepoMockRecorder) Update(ctx, twit any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepo)(nil).Update), ctx, twit)
}