package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementManagementResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilder provides operations to manage the environment property of the microsoft.graph.accessPackageResource entity.
type EntitlementManagementResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementManagementResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilderGetQueryParameters contains the environment information for the resource. This can be set using either the @odata.bind annotation or the environment's originId.Supports $expand.
type EntitlementManagementResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementManagementResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilderGetQueryParameters
}
// NewEntitlementManagementResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilderInternal instantiates a new EnvironmentRequestBuilder and sets the default values.
func NewEntitlementManagementResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilder) {
    m := &EntitlementManagementResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/resources/{accessPackageResource%2Did}/roles/{accessPackageResourceRole%2Did}/resource/scopes/{accessPackageResourceScope%2Did}/resource/environment{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewEntitlementManagementResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilder instantiates a new EnvironmentRequestBuilder and sets the default values.
func NewEntitlementManagementResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilderInternal(urlParams, requestAdapter)
}
// Get contains the environment information for the resource. This can be set using either the @odata.bind annotation or the environment's originId.Supports $expand.
func (m *EntitlementManagementResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceEnvironmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageResourceEnvironmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceEnvironmentable), nil
}
// ToGetRequestInformation contains the environment information for the resource. This can be set using either the @odata.bind annotation or the environment's originId.Supports $expand.
func (m *EntitlementManagementResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
