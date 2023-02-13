package mock_service

import (
	"context"
	"reflect"

	"github.com/golang/mock/gomock"
	"github.com/mrsubudei/task_for_golang_dev/users-service/internal/entity"
)

type MockUsers struct {
	ctrl     *gomock.Controller
	recorder *MockUsersRecorder
}

type MockUsersRecorder struct {
	mock *MockUsers
}

func NewMockUsers(ctrl *gomock.Controller) *MockUsers {
	mock := &MockUsers{ctrl: ctrl}
	mock.recorder = &MockUsersRecorder{mock}
	return mock
}

func (m *MockUsers) EXPECT() *MockUsersRecorder {
	return m.recorder
}

func (m *MockUsers) CreateUser(ctx context.Context, user entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockUsersRecorder) CreateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser",
		reflect.TypeOf((*MockUsers)(nil).CreateUser), ctx, user)
}

func (m *MockUsers) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByEmail", ctx, email)
	ret0, _ := ret[0].(entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockUsersRecorder) GetByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmail",
		reflect.TypeOf((*MockUsers)(nil).GetByEmail), ctx, email)
}
