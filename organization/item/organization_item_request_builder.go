package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i2c3e7e24e2730178382d4adade068738ce6025ae386d5724c24a3a268a81525c "github.com/microsoftgraph/msgraph-sdk-go/organization/item/getmembergroups"
    i3d6521f9c6af801263d8177b58bd3f14c44b32d31606bcefa9d9284325e9ae05 "github.com/microsoftgraph/msgraph-sdk-go/organization/item/checkmemberobjects"
    i62fcec4c98d8f7e9caa7bd4b260ea9d20f54a28de6683ff87d4101316c41f38a "github.com/microsoftgraph/msgraph-sdk-go/organization/item/extensions"
    i6be0b4e2753e0c9bfc4e1158ef7d0165dbba9087c2342ddf87c85eb2609049e8 "github.com/microsoftgraph/msgraph-sdk-go/organization/item/branding"
    i6d71976a15ac6f607366bcc5307ea5cc985e145be0c3586755194917f7fb4169 "github.com/microsoftgraph/msgraph-sdk-go/organization/item/checkmembergroups"
    i8885d5cee4c4d28a727da7b8bbf4e0a830b801820a872d899634b98dea40da8c "github.com/microsoftgraph/msgraph-sdk-go/organization/item/getmemberobjects"
    id705df728183d1caa2cd0a8e71b54c1f25c4bf34951aa4d58d74debd1d27c3d8 "github.com/microsoftgraph/msgraph-sdk-go/organization/item/restore"
    ieb9078c75a28c74bd2df523a3384718424b13d8e1f2c2fb11f57ab22d0557625 "github.com/microsoftgraph/msgraph-sdk-go/organization/item/setmobiledevicemanagementauthority"
    iee043428b7b7f62f2e5471e1c8a277c3f71551ba1ea8cb7e6ba5dbb3e2b0607d "github.com/microsoftgraph/msgraph-sdk-go/organization/item/certificatebasedauthconfiguration"
    i326a7c120a79ddb25aaa2231980538e1771104ba214e60fe5dac09d717662be7 "github.com/microsoftgraph/msgraph-sdk-go/organization/item/certificatebasedauthconfiguration/item"
    iead5f0a6000982961cff6f0e297f9427177afc718378c21027082d54c9ed6105 "github.com/microsoftgraph/msgraph-sdk-go/organization/item/extensions/item"
)

// OrganizationItemRequestBuilder provides operations to manage the collection of organization entities.
type OrganizationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OrganizationItemRequestBuilderDeleteOptions options for Delete
type OrganizationItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// OrganizationItemRequestBuilderGetOptions options for Get
type OrganizationItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *OrganizationItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// OrganizationItemRequestBuilderGetQueryParameters get entity from organization by key
type OrganizationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// OrganizationItemRequestBuilderPatchOptions options for Patch
type OrganizationItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Organizationable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// Branding the branding property
func (m *OrganizationItemRequestBuilder) Branding()(*i6be0b4e2753e0c9bfc4e1158ef7d0165dbba9087c2342ddf87c85eb2609049e8.BrandingRequestBuilder) {
    return i6be0b4e2753e0c9bfc4e1158ef7d0165dbba9087c2342ddf87c85eb2609049e8.NewBrandingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CertificateBasedAuthConfiguration the certificateBasedAuthConfiguration property
func (m *OrganizationItemRequestBuilder) CertificateBasedAuthConfiguration()(*iee043428b7b7f62f2e5471e1c8a277c3f71551ba1ea8cb7e6ba5dbb3e2b0607d.CertificateBasedAuthConfigurationRequestBuilder) {
    return iee043428b7b7f62f2e5471e1c8a277c3f71551ba1ea8cb7e6ba5dbb3e2b0607d.NewCertificateBasedAuthConfigurationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CertificateBasedAuthConfigurationById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.organization.item.certificateBasedAuthConfiguration.item collection
func (m *OrganizationItemRequestBuilder) CertificateBasedAuthConfigurationById(id string)(*i326a7c120a79ddb25aaa2231980538e1771104ba214e60fe5dac09d717662be7.CertificateBasedAuthConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["certificateBasedAuthConfiguration_id"] = id
    }
    return i326a7c120a79ddb25aaa2231980538e1771104ba214e60fe5dac09d717662be7.NewCertificateBasedAuthConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *OrganizationItemRequestBuilder) CheckMemberGroups()(*i6d71976a15ac6f607366bcc5307ea5cc985e145be0c3586755194917f7fb4169.CheckMemberGroupsRequestBuilder) {
    return i6d71976a15ac6f607366bcc5307ea5cc985e145be0c3586755194917f7fb4169.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *OrganizationItemRequestBuilder) CheckMemberObjects()(*i3d6521f9c6af801263d8177b58bd3f14c44b32d31606bcefa9d9284325e9ae05.CheckMemberObjectsRequestBuilder) {
    return i3d6521f9c6af801263d8177b58bd3f14c44b32d31606bcefa9d9284325e9ae05.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewOrganizationItemRequestBuilderInternal instantiates a new OrganizationItemRequestBuilder and sets the default values.
func NewOrganizationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrganizationItemRequestBuilder) {
    m := &OrganizationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/organization/{organization_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOrganizationItemRequestBuilder instantiates a new OrganizationItemRequestBuilder and sets the default values.
func NewOrganizationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrganizationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOrganizationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from organization
func (m *OrganizationItemRequestBuilder) CreateDeleteRequestInformation(options *OrganizationItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get entity from organization by key
func (m *OrganizationItemRequestBuilder) CreateGetRequestInformation(options *OrganizationItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update entity in organization
func (m *OrganizationItemRequestBuilder) CreatePatchRequestInformation(options *OrganizationItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete entity from organization
func (m *OrganizationItemRequestBuilder) Delete(options *OrganizationItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Extensions the extensions property
func (m *OrganizationItemRequestBuilder) Extensions()(*i62fcec4c98d8f7e9caa7bd4b260ea9d20f54a28de6683ff87d4101316c41f38a.ExtensionsRequestBuilder) {
    return i62fcec4c98d8f7e9caa7bd4b260ea9d20f54a28de6683ff87d4101316c41f38a.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.organization.item.extensions.item collection
func (m *OrganizationItemRequestBuilder) ExtensionsById(id string)(*iead5f0a6000982961cff6f0e297f9427177afc718378c21027082d54c9ed6105.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return iead5f0a6000982961cff6f0e297f9427177afc718378c21027082d54c9ed6105.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get entity from organization by key
func (m *OrganizationItemRequestBuilder) Get(options *OrganizationItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Organizationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOrganizationFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Organizationable), nil
}
// GetMemberGroups the getMemberGroups property
func (m *OrganizationItemRequestBuilder) GetMemberGroups()(*i2c3e7e24e2730178382d4adade068738ce6025ae386d5724c24a3a268a81525c.GetMemberGroupsRequestBuilder) {
    return i2c3e7e24e2730178382d4adade068738ce6025ae386d5724c24a3a268a81525c.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *OrganizationItemRequestBuilder) GetMemberObjects()(*i8885d5cee4c4d28a727da7b8bbf4e0a830b801820a872d899634b98dea40da8c.GetMemberObjectsRequestBuilder) {
    return i8885d5cee4c4d28a727da7b8bbf4e0a830b801820a872d899634b98dea40da8c.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update entity in organization
func (m *OrganizationItemRequestBuilder) Patch(options *OrganizationItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Restore the restore property
func (m *OrganizationItemRequestBuilder) Restore()(*id705df728183d1caa2cd0a8e71b54c1f25c4bf34951aa4d58d74debd1d27c3d8.RestoreRequestBuilder) {
    return id705df728183d1caa2cd0a8e71b54c1f25c4bf34951aa4d58d74debd1d27c3d8.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetMobileDeviceManagementAuthority the setMobileDeviceManagementAuthority property
func (m *OrganizationItemRequestBuilder) SetMobileDeviceManagementAuthority()(*ieb9078c75a28c74bd2df523a3384718424b13d8e1f2c2fb11f57ab22d0557625.SetMobileDeviceManagementAuthorityRequestBuilder) {
    return ieb9078c75a28c74bd2df523a3384718424b13d8e1f2c2fb11f57ab22d0557625.NewSetMobileDeviceManagementAuthorityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
