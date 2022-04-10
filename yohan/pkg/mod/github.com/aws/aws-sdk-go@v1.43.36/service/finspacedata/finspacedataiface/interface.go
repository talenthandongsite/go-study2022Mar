// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package finspacedataiface provides an interface to enable mocking the FinSpace Public API service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package finspacedataiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/finspacedata"
)

// FinSpaceDataAPI provides an interface to enable mocking the
// finspacedata.FinSpaceData service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // FinSpace Public API.
//    func myFunc(svc finspacedataiface.FinSpaceDataAPI) bool {
//        // Make svc.CreateChangeset request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := finspacedata.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockFinSpaceDataClient struct {
//        finspacedataiface.FinSpaceDataAPI
//    }
//    func (m *mockFinSpaceDataClient) CreateChangeset(input *finspacedata.CreateChangesetInput) (*finspacedata.CreateChangesetOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockFinSpaceDataClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type FinSpaceDataAPI interface {
	CreateChangeset(*finspacedata.CreateChangesetInput) (*finspacedata.CreateChangesetOutput, error)
	CreateChangesetWithContext(aws.Context, *finspacedata.CreateChangesetInput, ...request.Option) (*finspacedata.CreateChangesetOutput, error)
	CreateChangesetRequest(*finspacedata.CreateChangesetInput) (*request.Request, *finspacedata.CreateChangesetOutput)

	CreateDataView(*finspacedata.CreateDataViewInput) (*finspacedata.CreateDataViewOutput, error)
	CreateDataViewWithContext(aws.Context, *finspacedata.CreateDataViewInput, ...request.Option) (*finspacedata.CreateDataViewOutput, error)
	CreateDataViewRequest(*finspacedata.CreateDataViewInput) (*request.Request, *finspacedata.CreateDataViewOutput)

	CreateDataset(*finspacedata.CreateDatasetInput) (*finspacedata.CreateDatasetOutput, error)
	CreateDatasetWithContext(aws.Context, *finspacedata.CreateDatasetInput, ...request.Option) (*finspacedata.CreateDatasetOutput, error)
	CreateDatasetRequest(*finspacedata.CreateDatasetInput) (*request.Request, *finspacedata.CreateDatasetOutput)

	CreatePermissionGroup(*finspacedata.CreatePermissionGroupInput) (*finspacedata.CreatePermissionGroupOutput, error)
	CreatePermissionGroupWithContext(aws.Context, *finspacedata.CreatePermissionGroupInput, ...request.Option) (*finspacedata.CreatePermissionGroupOutput, error)
	CreatePermissionGroupRequest(*finspacedata.CreatePermissionGroupInput) (*request.Request, *finspacedata.CreatePermissionGroupOutput)

	CreateUser(*finspacedata.CreateUserInput) (*finspacedata.CreateUserOutput, error)
	CreateUserWithContext(aws.Context, *finspacedata.CreateUserInput, ...request.Option) (*finspacedata.CreateUserOutput, error)
	CreateUserRequest(*finspacedata.CreateUserInput) (*request.Request, *finspacedata.CreateUserOutput)

	DeleteDataset(*finspacedata.DeleteDatasetInput) (*finspacedata.DeleteDatasetOutput, error)
	DeleteDatasetWithContext(aws.Context, *finspacedata.DeleteDatasetInput, ...request.Option) (*finspacedata.DeleteDatasetOutput, error)
	DeleteDatasetRequest(*finspacedata.DeleteDatasetInput) (*request.Request, *finspacedata.DeleteDatasetOutput)

	DeletePermissionGroup(*finspacedata.DeletePermissionGroupInput) (*finspacedata.DeletePermissionGroupOutput, error)
	DeletePermissionGroupWithContext(aws.Context, *finspacedata.DeletePermissionGroupInput, ...request.Option) (*finspacedata.DeletePermissionGroupOutput, error)
	DeletePermissionGroupRequest(*finspacedata.DeletePermissionGroupInput) (*request.Request, *finspacedata.DeletePermissionGroupOutput)

	DisableUser(*finspacedata.DisableUserInput) (*finspacedata.DisableUserOutput, error)
	DisableUserWithContext(aws.Context, *finspacedata.DisableUserInput, ...request.Option) (*finspacedata.DisableUserOutput, error)
	DisableUserRequest(*finspacedata.DisableUserInput) (*request.Request, *finspacedata.DisableUserOutput)

	EnableUser(*finspacedata.EnableUserInput) (*finspacedata.EnableUserOutput, error)
	EnableUserWithContext(aws.Context, *finspacedata.EnableUserInput, ...request.Option) (*finspacedata.EnableUserOutput, error)
	EnableUserRequest(*finspacedata.EnableUserInput) (*request.Request, *finspacedata.EnableUserOutput)

	GetChangeset(*finspacedata.GetChangesetInput) (*finspacedata.GetChangesetOutput, error)
	GetChangesetWithContext(aws.Context, *finspacedata.GetChangesetInput, ...request.Option) (*finspacedata.GetChangesetOutput, error)
	GetChangesetRequest(*finspacedata.GetChangesetInput) (*request.Request, *finspacedata.GetChangesetOutput)

	GetDataView(*finspacedata.GetDataViewInput) (*finspacedata.GetDataViewOutput, error)
	GetDataViewWithContext(aws.Context, *finspacedata.GetDataViewInput, ...request.Option) (*finspacedata.GetDataViewOutput, error)
	GetDataViewRequest(*finspacedata.GetDataViewInput) (*request.Request, *finspacedata.GetDataViewOutput)

	GetDataset(*finspacedata.GetDatasetInput) (*finspacedata.GetDatasetOutput, error)
	GetDatasetWithContext(aws.Context, *finspacedata.GetDatasetInput, ...request.Option) (*finspacedata.GetDatasetOutput, error)
	GetDatasetRequest(*finspacedata.GetDatasetInput) (*request.Request, *finspacedata.GetDatasetOutput)

	GetProgrammaticAccessCredentials(*finspacedata.GetProgrammaticAccessCredentialsInput) (*finspacedata.GetProgrammaticAccessCredentialsOutput, error)
	GetProgrammaticAccessCredentialsWithContext(aws.Context, *finspacedata.GetProgrammaticAccessCredentialsInput, ...request.Option) (*finspacedata.GetProgrammaticAccessCredentialsOutput, error)
	GetProgrammaticAccessCredentialsRequest(*finspacedata.GetProgrammaticAccessCredentialsInput) (*request.Request, *finspacedata.GetProgrammaticAccessCredentialsOutput)

	GetUser(*finspacedata.GetUserInput) (*finspacedata.GetUserOutput, error)
	GetUserWithContext(aws.Context, *finspacedata.GetUserInput, ...request.Option) (*finspacedata.GetUserOutput, error)
	GetUserRequest(*finspacedata.GetUserInput) (*request.Request, *finspacedata.GetUserOutput)

	GetWorkingLocation(*finspacedata.GetWorkingLocationInput) (*finspacedata.GetWorkingLocationOutput, error)
	GetWorkingLocationWithContext(aws.Context, *finspacedata.GetWorkingLocationInput, ...request.Option) (*finspacedata.GetWorkingLocationOutput, error)
	GetWorkingLocationRequest(*finspacedata.GetWorkingLocationInput) (*request.Request, *finspacedata.GetWorkingLocationOutput)

	ListChangesets(*finspacedata.ListChangesetsInput) (*finspacedata.ListChangesetsOutput, error)
	ListChangesetsWithContext(aws.Context, *finspacedata.ListChangesetsInput, ...request.Option) (*finspacedata.ListChangesetsOutput, error)
	ListChangesetsRequest(*finspacedata.ListChangesetsInput) (*request.Request, *finspacedata.ListChangesetsOutput)

	ListChangesetsPages(*finspacedata.ListChangesetsInput, func(*finspacedata.ListChangesetsOutput, bool) bool) error
	ListChangesetsPagesWithContext(aws.Context, *finspacedata.ListChangesetsInput, func(*finspacedata.ListChangesetsOutput, bool) bool, ...request.Option) error

	ListDataViews(*finspacedata.ListDataViewsInput) (*finspacedata.ListDataViewsOutput, error)
	ListDataViewsWithContext(aws.Context, *finspacedata.ListDataViewsInput, ...request.Option) (*finspacedata.ListDataViewsOutput, error)
	ListDataViewsRequest(*finspacedata.ListDataViewsInput) (*request.Request, *finspacedata.ListDataViewsOutput)

	ListDataViewsPages(*finspacedata.ListDataViewsInput, func(*finspacedata.ListDataViewsOutput, bool) bool) error
	ListDataViewsPagesWithContext(aws.Context, *finspacedata.ListDataViewsInput, func(*finspacedata.ListDataViewsOutput, bool) bool, ...request.Option) error

	ListDatasets(*finspacedata.ListDatasetsInput) (*finspacedata.ListDatasetsOutput, error)
	ListDatasetsWithContext(aws.Context, *finspacedata.ListDatasetsInput, ...request.Option) (*finspacedata.ListDatasetsOutput, error)
	ListDatasetsRequest(*finspacedata.ListDatasetsInput) (*request.Request, *finspacedata.ListDatasetsOutput)

	ListDatasetsPages(*finspacedata.ListDatasetsInput, func(*finspacedata.ListDatasetsOutput, bool) bool) error
	ListDatasetsPagesWithContext(aws.Context, *finspacedata.ListDatasetsInput, func(*finspacedata.ListDatasetsOutput, bool) bool, ...request.Option) error

	ListPermissionGroups(*finspacedata.ListPermissionGroupsInput) (*finspacedata.ListPermissionGroupsOutput, error)
	ListPermissionGroupsWithContext(aws.Context, *finspacedata.ListPermissionGroupsInput, ...request.Option) (*finspacedata.ListPermissionGroupsOutput, error)
	ListPermissionGroupsRequest(*finspacedata.ListPermissionGroupsInput) (*request.Request, *finspacedata.ListPermissionGroupsOutput)

	ListPermissionGroupsPages(*finspacedata.ListPermissionGroupsInput, func(*finspacedata.ListPermissionGroupsOutput, bool) bool) error
	ListPermissionGroupsPagesWithContext(aws.Context, *finspacedata.ListPermissionGroupsInput, func(*finspacedata.ListPermissionGroupsOutput, bool) bool, ...request.Option) error

	ListUsers(*finspacedata.ListUsersInput) (*finspacedata.ListUsersOutput, error)
	ListUsersWithContext(aws.Context, *finspacedata.ListUsersInput, ...request.Option) (*finspacedata.ListUsersOutput, error)
	ListUsersRequest(*finspacedata.ListUsersInput) (*request.Request, *finspacedata.ListUsersOutput)

	ListUsersPages(*finspacedata.ListUsersInput, func(*finspacedata.ListUsersOutput, bool) bool) error
	ListUsersPagesWithContext(aws.Context, *finspacedata.ListUsersInput, func(*finspacedata.ListUsersOutput, bool) bool, ...request.Option) error

	ResetUserPassword(*finspacedata.ResetUserPasswordInput) (*finspacedata.ResetUserPasswordOutput, error)
	ResetUserPasswordWithContext(aws.Context, *finspacedata.ResetUserPasswordInput, ...request.Option) (*finspacedata.ResetUserPasswordOutput, error)
	ResetUserPasswordRequest(*finspacedata.ResetUserPasswordInput) (*request.Request, *finspacedata.ResetUserPasswordOutput)

	UpdateChangeset(*finspacedata.UpdateChangesetInput) (*finspacedata.UpdateChangesetOutput, error)
	UpdateChangesetWithContext(aws.Context, *finspacedata.UpdateChangesetInput, ...request.Option) (*finspacedata.UpdateChangesetOutput, error)
	UpdateChangesetRequest(*finspacedata.UpdateChangesetInput) (*request.Request, *finspacedata.UpdateChangesetOutput)

	UpdateDataset(*finspacedata.UpdateDatasetInput) (*finspacedata.UpdateDatasetOutput, error)
	UpdateDatasetWithContext(aws.Context, *finspacedata.UpdateDatasetInput, ...request.Option) (*finspacedata.UpdateDatasetOutput, error)
	UpdateDatasetRequest(*finspacedata.UpdateDatasetInput) (*request.Request, *finspacedata.UpdateDatasetOutput)

	UpdatePermissionGroup(*finspacedata.UpdatePermissionGroupInput) (*finspacedata.UpdatePermissionGroupOutput, error)
	UpdatePermissionGroupWithContext(aws.Context, *finspacedata.UpdatePermissionGroupInput, ...request.Option) (*finspacedata.UpdatePermissionGroupOutput, error)
	UpdatePermissionGroupRequest(*finspacedata.UpdatePermissionGroupInput) (*request.Request, *finspacedata.UpdatePermissionGroupOutput)

	UpdateUser(*finspacedata.UpdateUserInput) (*finspacedata.UpdateUserOutput, error)
	UpdateUserWithContext(aws.Context, *finspacedata.UpdateUserInput, ...request.Option) (*finspacedata.UpdateUserOutput, error)
	UpdateUserRequest(*finspacedata.UpdateUserInput) (*request.Request, *finspacedata.UpdateUserOutput)
}

var _ FinSpaceDataAPI = (*finspacedata.FinSpaceData)(nil)
