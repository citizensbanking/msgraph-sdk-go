package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i0f4b35070ae3fb3343fe356519400fd185fdc99c05b968ef8e5ff1f5951748ae "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/list/contenttypes/item/columnpositions"
    i32d56bc0af1bf5c57d90c3ff019f678d038230a9bfe4aae915812c67795e46ec "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/list/contenttypes/item/basetypes"
    i6e95862d0c985a15ed1b633c992759780a2e299b4630622464198db77beca3f5 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/list/contenttypes/item/ispublished"
    i83b288356df169bec7c416a5445858c6b0c17e5aec455ccded399f5883855181 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/list/contenttypes/item/base"
    i95e2f2ece45392396d7ed96f3c7c97e84076b4dfe1334b84148a29d5f47a116b "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/list/contenttypes/item/columns"
    i9830bec44bb107f7d7b788fa185d37f3515f38f70df85c780f3baa07ce0333cd "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/list/contenttypes/item/publish"
    i9a4eeebb28c8455c06b95e00db7ffbb4cba1af454d9b5fae7735786753aa0a10 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/list/contenttypes/item/copytodefaultcontentlocation"
    ib012a85189ac10702a8a311452bfa7b69e8ed59ec11032d5089b2a9cacd1e309 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/list/contenttypes/item/columnlinks"
    ie20358a85b39ee8ba95363a576c6e46d8800c6b70c5da48febb10bffc24bd2c5 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/list/contenttypes/item/associatewithhubsites"
    if2999294f099d5c8f5ba396893319eb983c4f664d57ca41b1f865fe0776e16b5 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/list/contenttypes/item/unpublish"
    i4c516dcbb87e6f393fb823d5471fd76632fb899dab27b71dc31731ea51d24234 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/list/contenttypes/item/columnpositions/item"
    ia4250265f061fb1dcec42f660bcab9c2e9fa481c6195e5768a60613bab0ca83f "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/list/contenttypes/item/columnlinks/item"
    icc9a32d3c9a2d87f0bc748ba2498d23f75930fab9f3749d72ef6c4c2a1768058 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/list/contenttypes/item/basetypes/item"
    id5f8ab988842a4e86e9ebbc90fc22f5ea8ba07bf75f2d44d596174e9e1cb94ef "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/list/contenttypes/item/columns/item"
)

// ContentTypeItemRequestBuilder provides operations to manage the contentTypes property of the microsoft.graph.list entity.
type ContentTypeItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ContentTypeItemRequestBuilderDeleteOptions options for Delete
type ContentTypeItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// ContentTypeItemRequestBuilderGetOptions options for Get
type ContentTypeItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *ContentTypeItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// ContentTypeItemRequestBuilderGetQueryParameters the collection of content types present in this list.
type ContentTypeItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ContentTypeItemRequestBuilderPatchOptions options for Patch
type ContentTypeItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// AssociateWithHubSites the associateWithHubSites property
func (m *ContentTypeItemRequestBuilder) AssociateWithHubSites()(*ie20358a85b39ee8ba95363a576c6e46d8800c6b70c5da48febb10bffc24bd2c5.AssociateWithHubSitesRequestBuilder) {
    return ie20358a85b39ee8ba95363a576c6e46d8800c6b70c5da48febb10bffc24bd2c5.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Base the base property
func (m *ContentTypeItemRequestBuilder) Base()(*i83b288356df169bec7c416a5445858c6b0c17e5aec455ccded399f5883855181.BaseRequestBuilder) {
    return i83b288356df169bec7c416a5445858c6b0c17e5aec455ccded399f5883855181.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypes the baseTypes property
func (m *ContentTypeItemRequestBuilder) BaseTypes()(*i32d56bc0af1bf5c57d90c3ff019f678d038230a9bfe4aae915812c67795e46ec.BaseTypesRequestBuilder) {
    return i32d56bc0af1bf5c57d90c3ff019f678d038230a9bfe4aae915812c67795e46ec.NewBaseTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.drives.item.list.contentTypes.item.baseTypes.item collection
func (m *ContentTypeItemRequestBuilder) BaseTypesById(id string)(*icc9a32d3c9a2d87f0bc748ba2498d23f75930fab9f3749d72ef6c4c2a1768058.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType_id1"] = id
    }
    return icc9a32d3c9a2d87f0bc748ba2498d23f75930fab9f3749d72ef6c4c2a1768058.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ColumnLinks the columnLinks property
func (m *ContentTypeItemRequestBuilder) ColumnLinks()(*ib012a85189ac10702a8a311452bfa7b69e8ed59ec11032d5089b2a9cacd1e309.ColumnLinksRequestBuilder) {
    return ib012a85189ac10702a8a311452bfa7b69e8ed59ec11032d5089b2a9cacd1e309.NewColumnLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnLinksById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.drives.item.list.contentTypes.item.columnLinks.item collection
func (m *ContentTypeItemRequestBuilder) ColumnLinksById(id string)(*ia4250265f061fb1dcec42f660bcab9c2e9fa481c6195e5768a60613bab0ca83f.ColumnLinkItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnLink_id"] = id
    }
    return ia4250265f061fb1dcec42f660bcab9c2e9fa481c6195e5768a60613bab0ca83f.NewColumnLinkItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ColumnPositions the columnPositions property
func (m *ContentTypeItemRequestBuilder) ColumnPositions()(*i0f4b35070ae3fb3343fe356519400fd185fdc99c05b968ef8e5ff1f5951748ae.ColumnPositionsRequestBuilder) {
    return i0f4b35070ae3fb3343fe356519400fd185fdc99c05b968ef8e5ff1f5951748ae.NewColumnPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnPositionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.drives.item.list.contentTypes.item.columnPositions.item collection
func (m *ContentTypeItemRequestBuilder) ColumnPositionsById(id string)(*i4c516dcbb87e6f393fb823d5471fd76632fb899dab27b71dc31731ea51d24234.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i4c516dcbb87e6f393fb823d5471fd76632fb899dab27b71dc31731ea51d24234.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Columns the columns property
func (m *ContentTypeItemRequestBuilder) Columns()(*i95e2f2ece45392396d7ed96f3c7c97e84076b4dfe1334b84148a29d5f47a116b.ColumnsRequestBuilder) {
    return i95e2f2ece45392396d7ed96f3c7c97e84076b4dfe1334b84148a29d5f47a116b.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.drives.item.list.contentTypes.item.columns.item collection
func (m *ContentTypeItemRequestBuilder) ColumnsById(id string)(*id5f8ab988842a4e86e9ebbc90fc22f5ea8ba07bf75f2d44d596174e9e1cb94ef.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return id5f8ab988842a4e86e9ebbc90fc22f5ea8ba07bf75f2d44d596174e9e1cb94ef.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewContentTypeItemRequestBuilderInternal instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewContentTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContentTypeItemRequestBuilder) {
    m := &ContentTypeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/drives/{drive_id}/list/contentTypes/{contentType_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewContentTypeItemRequestBuilder instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewContentTypeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContentTypeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContentTypeItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CopyToDefaultContentLocation the copyToDefaultContentLocation property
func (m *ContentTypeItemRequestBuilder) CopyToDefaultContentLocation()(*i9a4eeebb28c8455c06b95e00db7ffbb4cba1af454d9b5fae7735786753aa0a10.CopyToDefaultContentLocationRequestBuilder) {
    return i9a4eeebb28c8455c06b95e00db7ffbb4cba1af454d9b5fae7735786753aa0a10.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property contentTypes for groups
func (m *ContentTypeItemRequestBuilder) CreateDeleteRequestInformation(options *ContentTypeItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the collection of content types present in this list.
func (m *ContentTypeItemRequestBuilder) CreateGetRequestInformation(options *ContentTypeItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property contentTypes in groups
func (m *ContentTypeItemRequestBuilder) CreatePatchRequestInformation(options *ContentTypeItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property contentTypes for groups
func (m *ContentTypeItemRequestBuilder) Delete(options *ContentTypeItemRequestBuilderDeleteOptions)(error) {
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
// Get the collection of content types present in this list.
func (m *ContentTypeItemRequestBuilder) Get(options *ContentTypeItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateContentTypeFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable), nil
}
// IsPublished provides operations to call the isPublished method.
func (m *ContentTypeItemRequestBuilder) IsPublished()(*i6e95862d0c985a15ed1b633c992759780a2e299b4630622464198db77beca3f5.IsPublishedRequestBuilder) {
    return i6e95862d0c985a15ed1b633c992759780a2e299b4630622464198db77beca3f5.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property contentTypes in groups
func (m *ContentTypeItemRequestBuilder) Patch(options *ContentTypeItemRequestBuilderPatchOptions)(error) {
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
// Publish the publish property
func (m *ContentTypeItemRequestBuilder) Publish()(*i9830bec44bb107f7d7b788fa185d37f3515f38f70df85c780f3baa07ce0333cd.PublishRequestBuilder) {
    return i9830bec44bb107f7d7b788fa185d37f3515f38f70df85c780f3baa07ce0333cd.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Unpublish the unpublish property
func (m *ContentTypeItemRequestBuilder) Unpublish()(*if2999294f099d5c8f5ba396893319eb983c4f664d57ca41b1f865fe0776e16b5.UnpublishRequestBuilder) {
    return if2999294f099d5c8f5ba396893319eb983c4f664d57ca41b1f865fe0776e16b5.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
