package mock_repository

import (
	"context"
	"reflect"

	"github.com/golang/mock/gomock"
	"github.com/mrsubudei/task_for_golang_dev/users-service/internal/entity"
)

type MockUsersRepo struct {
	ctrl     *gomock.Controller
	recorder *MockUsersRepoRecorder
}

type MockUsersRepoRecorder struct {
	mock *MockUsersRepo
}

func NewMockUsersRepo(ctrl *gomock.Controller) *MockUsersRepo {
	mock := &MockUsersRepo{ctrl: ctrl}
	mock.recorder = &MockUsersRepoRecorder{mock}
	return mock
}

func (m *MockUsersRepo) EXPECT() *MockUsersRepoRecorder {
	return m.recorder
}

func (m *MockUsersRepo) Create(ctx context.Context, user entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockUsersRepoRecorder) Create(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create",
		reflect.TypeOf((*MockUsersRepo)(nil).Create), ctx, user)
}

func (m *MockUsersRepo) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByEmail", ctx, email)
	ret0, _ := ret[0].(entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockUsersRepoRecorder) GetByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmail",
		reflect.TypeOf((*MockUsersRepo)(nil).GetByEmail), ctx, email)
}
