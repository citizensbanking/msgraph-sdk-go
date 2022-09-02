package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i0f7a3c302b1a158434cb422453f66a05f39699c51e438b2a1d0eabcfa4bba8a8 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/serviceprincipal"
    i165529ec4316235bfb940a989eb188c85c6648c79edcf72dd2df66a0bb7a1575 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user"
    i6d9c969a7247ebc608622ea741daf6521f2190d70c12ef5b7e3649820fe8dabd "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/application"
    i6ead2f90c7af014f21ce5508f15b42df1693b449024f90c152b9bcd2c5750218 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/device"
    ia3f7e359fe1fc2870203c8395458f6cccc3032f97372914eb0d74926b610b98a "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/group"
    ifb0097e00ea120142aad1c9afc070fc7793e093ab7607e87726f6d165bbc72f0 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/orgcontact"
)

// DirectoryObjectItemRequestBuilder provides operations to manage the transitiveMemberOf property of the microsoft.graph.device entity.
type DirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DirectoryObjectItemRequestBuilderGetQueryParameters groups and administrative units that the device is a member of. This operation is transitive. Supports $expand.
type DirectoryObjectItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DirectoryObjectItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryObjectItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryObjectItemRequestBuilderGetQueryParameters
}
// Application the application property
func (m *DirectoryObjectItemRequestBuilder) Application()(*i6d9c969a7247ebc608622ea741daf6521f2190d70c12ef5b7e3649820fe8dabd.ApplicationRequestBuilder) {
    return i6d9c969a7247ebc608622ea741daf6521f2190d70c12ef5b7e3649820fe8dabd.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryObjectItemRequestBuilder) {
    m := &DirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod%2Did}/device/transitiveMemberOf/{directoryObject%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation groups and administrative units that the device is a member of. This operation is transitive. Supports $expand.
func (m *DirectoryObjectItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration groups and administrative units that the device is a member of. This operation is transitive. Supports $expand.
func (m *DirectoryObjectItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *DirectoryObjectItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Device the device property
func (m *DirectoryObjectItemRequestBuilder) Device()(*i6ead2f90c7af014f21ce5508f15b42df1693b449024f90c152b9bcd2c5750218.DeviceRequestBuilder) {
    return i6ead2f90c7af014f21ce5508f15b42df1693b449024f90c152b9bcd2c5750218.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get groups and administrative units that the device is a member of. This operation is transitive. Supports $expand.
func (m *DirectoryObjectItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectoryObjectItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDirectoryObjectFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable), nil
}
// Group the group property
func (m *DirectoryObjectItemRequestBuilder) Group()(*ia3f7e359fe1fc2870203c8395458f6cccc3032f97372914eb0d74926b610b98a.GroupRequestBuilder) {
    return ia3f7e359fe1fc2870203c8395458f6cccc3032f97372914eb0d74926b610b98a.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact the orgContact property
func (m *DirectoryObjectItemRequestBuilder) OrgContact()(*ifb0097e00ea120142aad1c9afc070fc7793e093ab7607e87726f6d165bbc72f0.OrgContactRequestBuilder) {
    return ifb0097e00ea120142aad1c9afc070fc7793e093ab7607e87726f6d165bbc72f0.NewOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *DirectoryObjectItemRequestBuilder) ServicePrincipal()(*i0f7a3c302b1a158434cb422453f66a05f39699c51e438b2a1d0eabcfa4bba8a8.ServicePrincipalRequestBuilder) {
    return i0f7a3c302b1a158434cb422453f66a05f39699c51e438b2a1d0eabcfa4bba8a8.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *DirectoryObjectItemRequestBuilder) User()(*i165529ec4316235bfb940a989eb188c85c6648c79edcf72dd2df66a0bb7a1575.UserRequestBuilder) {
    return i165529ec4316235bfb940a989eb188c85c6648c79edcf72dd2df66a0bb7a1575.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
