// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/BoggerByte/Sentinel-backend.git/pkg/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	db "github.com/BoggerByte/Sentinel-backend.git/pkg/db/sqlc"
	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateOrUpdateGuild mocks base method.
func (m *MockStore) CreateOrUpdateGuild(arg0 context.Context, arg1 db.CreateOrUpdateGuildParams) (db.Guild, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateGuild", arg0, arg1)
	ret0, _ := ret[0].(db.Guild)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdateGuild indicates an expected call of CreateOrUpdateGuild.
func (mr *MockStoreMockRecorder) CreateOrUpdateGuild(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateGuild", reflect.TypeOf((*MockStore)(nil).CreateOrUpdateGuild), arg0, arg1)
}

// CreateOrUpdateGuildConfig mocks base method.
func (m *MockStore) CreateOrUpdateGuildConfig(arg0 context.Context, arg1 db.CreateOrUpdateGuildConfigParams) (db.GuildConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateGuildConfig", arg0, arg1)
	ret0, _ := ret[0].(db.GuildConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdateGuildConfig indicates an expected call of CreateOrUpdateGuildConfig.
func (mr *MockStoreMockRecorder) CreateOrUpdateGuildConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateGuildConfig", reflect.TypeOf((*MockStore)(nil).CreateOrUpdateGuildConfig), arg0, arg1)
}

// CreateOrUpdateUser mocks base method.
func (m *MockStore) CreateOrUpdateUser(arg0 context.Context, arg1 db.CreateOrUpdateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdateUser indicates an expected call of CreateOrUpdateUser.
func (mr *MockStoreMockRecorder) CreateOrUpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateUser", reflect.TypeOf((*MockStore)(nil).CreateOrUpdateUser), arg0, arg1)
}

// CreateOrUpdateUserGuildRel mocks base method.
func (m *MockStore) CreateOrUpdateUserGuildRel(arg0 context.Context, arg1 db.CreateOrUpdateUserGuildRelParams) (db.UserGuild, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateUserGuildRel", arg0, arg1)
	ret0, _ := ret[0].(db.UserGuild)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdateUserGuildRel indicates an expected call of CreateOrUpdateUserGuildRel.
func (mr *MockStoreMockRecorder) CreateOrUpdateUserGuildRel(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateUserGuildRel", reflect.TypeOf((*MockStore)(nil).CreateOrUpdateUserGuildRel), arg0, arg1)
}

// CreateUserGuildRel mocks base method.
func (m *MockStore) CreateUserGuildRel(arg0 context.Context, arg1 db.CreateUserGuildRelParams) (db.UserGuild, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserGuildRel", arg0, arg1)
	ret0, _ := ret[0].(db.UserGuild)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUserGuildRel indicates an expected call of CreateUserGuildRel.
func (mr *MockStoreMockRecorder) CreateUserGuildRel(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserGuildRel", reflect.TypeOf((*MockStore)(nil).CreateUserGuildRel), arg0, arg1)
}

// ExecTx mocks base method.
func (m *MockStore) ExecTx(arg0 context.Context, arg1 func(*db.Queries) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecTx", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExecTx indicates an expected call of ExecTx.
func (mr *MockStoreMockRecorder) ExecTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecTx", reflect.TypeOf((*MockStore)(nil).ExecTx), arg0, arg1)
}

// GetGuild mocks base method.
func (m *MockStore) GetGuild(arg0 context.Context, arg1 string) (db.GetGuildRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGuild", arg0, arg1)
	ret0, _ := ret[0].(db.GetGuildRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGuild indicates an expected call of GetGuild.
func (mr *MockStoreMockRecorder) GetGuild(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGuild", reflect.TypeOf((*MockStore)(nil).GetGuild), arg0, arg1)
}

// GetGuildConfig mocks base method.
func (m *MockStore) GetGuildConfig(arg0 context.Context, arg1 string) (db.GuildConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGuildConfig", arg0, arg1)
	ret0, _ := ret[0].(db.GuildConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGuildConfig indicates an expected call of GetGuildConfig.
func (mr *MockStoreMockRecorder) GetGuildConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGuildConfig", reflect.TypeOf((*MockStore)(nil).GetGuildConfig), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGuild", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGuild", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// GetUserGuild mocks base method.
func (m *MockStore) GetUserGuild(arg0 context.Context, arg1 db.GetUserGuildParams) (db.GetUserGuildRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserGuild", arg0, arg1)
	ret0, _ := ret[0].(db.GetUserGuildRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserGuild indicates an expected call of GetUserGuild.
func (mr *MockStoreMockRecorder) GetUserGuild(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserGuild", reflect.TypeOf((*MockStore)(nil).GetUserGuild), arg0, arg1)
}

// GetUserGuildRel mocks base method.
func (m *MockStore) GetUserGuildRel(arg0 context.Context, arg1 db.GetUserGuildRelParams) (db.UserGuild, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserGuildRel", arg0, arg1)
	ret0, _ := ret[0].(db.UserGuild)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserGuildRel indicates an expected call of GetUserGuildRel.
func (mr *MockStoreMockRecorder) GetUserGuildRel(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserGuildRel", reflect.TypeOf((*MockStore)(nil).GetUserGuildRel), arg0, arg1)
}

// GetUserGuilds mocks base method.
func (m *MockStore) GetUserGuilds(arg0 context.Context, arg1 string) ([]db.GetUserGuildsRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserGuilds", arg0, arg1)
	ret0, _ := ret[0].([]db.GetUserGuildsRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserGuilds indicates an expected call of GetUserGuilds.
func (mr *MockStoreMockRecorder) GetUserGuilds(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserGuilds", reflect.TypeOf((*MockStore)(nil).GetUserGuilds), arg0, arg1)
}

// TryCreateGuildConfig mocks base method.
func (m *MockStore) TryCreateGuildConfig(arg0 context.Context, arg1 db.TryCreateGuildConfigParams) (db.GuildConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TryCreateGuildConfig", arg0, arg1)
	ret0, _ := ret[0].(db.GuildConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TryCreateGuildConfig indicates an expected call of TryCreateGuildConfig.
func (mr *MockStoreMockRecorder) TryCreateGuildConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TryCreateGuildConfig", reflect.TypeOf((*MockStore)(nil).TryCreateGuildConfig), arg0, arg1)
}

// UpdateGuildConfig mocks base method.
func (m *MockStore) UpdateGuildConfig(arg0 context.Context, arg1 db.UpdateGuildConfigParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGuildConfig", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateGuildConfig indicates an expected call of UpdateGuildConfig.
func (mr *MockStoreMockRecorder) UpdateGuildConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGuildConfig", reflect.TypeOf((*MockStore)(nil).UpdateGuildConfig), arg0, arg1)
}
