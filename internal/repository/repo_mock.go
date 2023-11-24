// Code generated by MockGen. DO NOT EDIT.
// Source: repo.go
//
// Generated by this command:
//
//	mockgen -source=repo.go -destination=repo_mock.go -package=repository
//
// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	models "job-portal-api/internal/models"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockUserRepo is a mock of UserRepo interface.
type MockUserRepo struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepoMockRecorder
}

// MockUserRepoMockRecorder is the mock recorder for MockUserRepo.
type MockUserRepoMockRecorder struct {
	mock *MockUserRepo
}

// NewMockUserRepo creates a new mock instance.
func NewMockUserRepo(ctrl *gomock.Controller) *MockUserRepo {
	mock := &MockUserRepo{ctrl: ctrl}
	mock.recorder = &MockUserRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepo) EXPECT() *MockUserRepoMockRecorder {
	return m.recorder
}

// CheckEmail mocks base method.
func (m *MockUserRepo) CheckEmail(ctx context.Context, email string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckEmail", ctx, email)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckEmail indicates an expected call of CheckEmail.
func (mr *MockUserRepoMockRecorder) CheckEmail(ctx, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckEmail", reflect.TypeOf((*MockUserRepo)(nil).CheckEmail), ctx, email)
}

// CreateCom mocks base method.
func (m *MockUserRepo) CreateCom(nc models.Company) (models.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCom", nc)
	ret0, _ := ret[0].(models.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCom indicates an expected call of CreateCom.
func (mr *MockUserRepoMockRecorder) CreateCom(nc any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCom", reflect.TypeOf((*MockUserRepo)(nil).CreateCom), nc)
}

// CreateUser mocks base method.
func (m *MockUserRepo) CreateUser(ctx context.Context, nu models.User) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, nu)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepoMockRecorder) CreateUser(ctx, nu any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepo)(nil).CreateUser), ctx, nu)
}

// FetchJobData mocks base method.
func (m *MockUserRepo) FetchJobData(jid uint64) (models.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchJobData", jid)
	ret0, _ := ret[0].(models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchJobData indicates an expected call of FetchJobData.
func (mr *MockUserRepoMockRecorder) FetchJobData(jid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchJobData", reflect.TypeOf((*MockUserRepo)(nil).FetchJobData), jid)
}

// GetAllJobs mocks base method.
func (m *MockUserRepo) GetAllJobs() ([]models.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllJobs")
	ret0, _ := ret[0].([]models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllJobs indicates an expected call of GetAllJobs.
func (mr *MockUserRepoMockRecorder) GetAllJobs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllJobs", reflect.TypeOf((*MockUserRepo)(nil).GetAllJobs))
}

// GetAllTheCompanies mocks base method.
func (m *MockUserRepo) GetAllTheCompanies() ([]models.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTheCompanies")
	ret0, _ := ret[0].([]models.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTheCompanies indicates an expected call of GetAllTheCompanies.
func (mr *MockUserRepoMockRecorder) GetAllTheCompanies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTheCompanies", reflect.TypeOf((*MockUserRepo)(nil).GetAllTheCompanies))
}

// GetCompany mocks base method.
func (m *MockUserRepo) GetCompany(id uint64) (models.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCompany", id)
	ret0, _ := ret[0].(models.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompany indicates an expected call of GetCompany.
func (mr *MockUserRepoMockRecorder) GetCompany(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompany", reflect.TypeOf((*MockUserRepo)(nil).GetCompany), id)
}

// GetJobsFromCompany mocks base method.
func (m *MockUserRepo) GetJobsFromCompany(comapny_id uint64) ([]models.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJobsFromCompany", comapny_id)
	ret0, _ := ret[0].([]models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobsFromCompany indicates an expected call of GetJobsFromCompany.
func (mr *MockUserRepoMockRecorder) GetJobsFromCompany(comapny_id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobsFromCompany", reflect.TypeOf((*MockUserRepo)(nil).GetJobsFromCompany), comapny_id)
}

// GetOneJob mocks base method.
func (m *MockUserRepo) GetOneJob(id uint64) ([]models.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOneJob", id)
	ret0, _ := ret[0].([]models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOneJob indicates an expected call of GetOneJob.
func (mr *MockUserRepoMockRecorder) GetOneJob(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOneJob", reflect.TypeOf((*MockUserRepo)(nil).GetOneJob), id)
}

// PostJob mocks base method.
func (m *MockUserRepo) PostJob(nj models.Job) (models.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostJob", nj)
	ret0, _ := ret[0].(models.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostJob indicates an expected call of PostJob.
func (mr *MockUserRepoMockRecorder) PostJob(nj any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostJob", reflect.TypeOf((*MockUserRepo)(nil).PostJob), nj)
}