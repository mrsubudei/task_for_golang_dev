// package mock_repository mocks users reository
package mock_repository

import (
	"context"
	"reflect"

	"github.com/golang/mock/gomock"
	"github.com/mrsubudei/task_for_golang_dev/users-service/internal/entity"
)

// MockUsersRepo is a mock of UsersRepo Interface
type MockUsersRepo struct {
	ctrl     *gomock.Controller
	recorder *MockUsersRepoRecorder
}

// MockUsersRepoRecorder is a mock recorder for MockUsersRepo
type MockUsersRepoRecorder struct {
	mock *MockUsersRepo
}

// NewMockUsersRepo creates new MockUsersRepo
func NewMockUsersRepo(ctrl *gomock.Controller) *MockUsersRepo {
	mock := &MockUsersRepo{ctrl: ctrl}
	mock.recorder = &MockUsersRepoRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsersRepo) EXPECT() *MockUsersRepoRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockUsersRepo) Create(ctx context.Context, user entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUsersRepoRecorder) Create(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create",
		reflect.TypeOf((*MockUsersRepo)(nil).Create), ctx, user)
}

// GetByEmail mocks base method
func (m *MockUsersRepo) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByEmail", ctx, email)
	ret0, _ := ret[0].(entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByEmail indicates an expected call of GetByEmail.
func (mr *MockUsersRepoRecorder) GetByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmail",
		reflect.TypeOf((*MockUsersRepo)(nil).GetByEmail), ctx, email)
}
