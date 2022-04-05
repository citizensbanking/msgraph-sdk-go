package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i524bf6311b24b6f141d885728d606289b8ee24cbd67c3599ccb2ec6084cb17bd "github.com/microsoftgraph/msgraph-sdk-go/groups/item/conversations/item/threads/item/posts/item/reply"
    i78e7a8cdddabb1f91bbd12d096fb8e9c132c2ec07d32fd988ce7552cfd89352b "github.com/microsoftgraph/msgraph-sdk-go/groups/item/conversations/item/threads/item/posts/item/forward"
    i9884a55ef28596c5da82f51922aa27afafadf6bc4977af30511de7bda6744049 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/conversations/item/threads/item/posts/item/singlevalueextendedproperties"
    ia669f29093557a86be3b54973c4aba4e78b09891dc40e9d967d02b3938cbb08d "github.com/microsoftgraph/msgraph-sdk-go/groups/item/conversations/item/threads/item/posts/item/inreplyto"
    ie5a656c04186903bcf4da8c7768bc7acdefb45b9b92af375c844ab6dde924c3b "github.com/microsoftgraph/msgraph-sdk-go/groups/item/conversations/item/threads/item/posts/item/multivalueextendedproperties"
    if00c8974308ec134f8b54a4b85c9b080451e578b97ae45f3a96f2983e98b235b "github.com/microsoftgraph/msgraph-sdk-go/groups/item/conversations/item/threads/item/posts/item/attachments"
    if33d4aff747bbefe22c9d4c0581a32241f79f8c00b65004202a691770673b2d8 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/conversations/item/threads/item/posts/item/extensions"
    i81cf5d0e294523ab35bef33098c9a475ca60e40c2f945ab2e768942447194086 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/conversations/item/threads/item/posts/item/multivalueextendedproperties/item"
    i85387a3ef1c4b59c2db207f407864b3a4fca529975d344ab83cd21fd615de2e6 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/conversations/item/threads/item/posts/item/singlevalueextendedproperties/item"
    i98e0b63002f7d7bb02acdac8815a30c2058502950f307d6652a0b2d766b7f02b "github.com/microsoftgraph/msgraph-sdk-go/groups/item/conversations/item/threads/item/posts/item/attachments/item"
    i9deed3605dbdfd0f6b5413628ad1be773f48a2ee7f6266b85266a4248af81477 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/conversations/item/threads/item/posts/item/extensions/item"
)

// PostItemRequestBuilder provides operations to manage the posts property of the microsoft.graph.conversationThread entity.
type PostItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// PostItemRequestBuilderDeleteOptions options for Delete
type PostItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// PostItemRequestBuilderGetOptions options for Get
type PostItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *PostItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// PostItemRequestBuilderGetQueryParameters read-only. Nullable.
type PostItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// PostItemRequestBuilderPatchOptions options for Patch
type PostItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Postable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// Attachments the attachments property
func (m *PostItemRequestBuilder) Attachments()(*if00c8974308ec134f8b54a4b85c9b080451e578b97ae45f3a96f2983e98b235b.AttachmentsRequestBuilder) {
    return if00c8974308ec134f8b54a4b85c9b080451e578b97ae45f3a96f2983e98b235b.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.conversations.item.threads.item.posts.item.attachments.item collection
func (m *PostItemRequestBuilder) AttachmentsById(id string)(*i98e0b63002f7d7bb02acdac8815a30c2058502950f307d6652a0b2d766b7f02b.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i98e0b63002f7d7bb02acdac8815a30c2058502950f307d6652a0b2d766b7f02b.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewPostItemRequestBuilderInternal instantiates a new PostItemRequestBuilder and sets the default values.
func NewPostItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PostItemRequestBuilder) {
    m := &PostItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/conversations/{conversation_id}/threads/{conversationThread_id}/posts/{post_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPostItemRequestBuilder instantiates a new PostItemRequestBuilder and sets the default values.
func NewPostItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PostItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPostItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property posts for groups
func (m *PostItemRequestBuilder) CreateDeleteRequestInformation(options *PostItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation read-only. Nullable.
func (m *PostItemRequestBuilder) CreateGetRequestInformation(options *PostItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property posts in groups
func (m *PostItemRequestBuilder) CreatePatchRequestInformation(options *PostItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property posts for groups
func (m *PostItemRequestBuilder) Delete(options *PostItemRequestBuilderDeleteOptions)(error) {
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
func (m *PostItemRequestBuilder) Extensions()(*if33d4aff747bbefe22c9d4c0581a32241f79f8c00b65004202a691770673b2d8.ExtensionsRequestBuilder) {
    return if33d4aff747bbefe22c9d4c0581a32241f79f8c00b65004202a691770673b2d8.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.conversations.item.threads.item.posts.item.extensions.item collection
func (m *PostItemRequestBuilder) ExtensionsById(id string)(*i9deed3605dbdfd0f6b5413628ad1be773f48a2ee7f6266b85266a4248af81477.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i9deed3605dbdfd0f6b5413628ad1be773f48a2ee7f6266b85266a4248af81477.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *PostItemRequestBuilder) Forward()(*i78e7a8cdddabb1f91bbd12d096fb8e9c132c2ec07d32fd988ce7552cfd89352b.ForwardRequestBuilder) {
    return i78e7a8cdddabb1f91bbd12d096fb8e9c132c2ec07d32fd988ce7552cfd89352b.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get read-only. Nullable.
func (m *PostItemRequestBuilder) Get(options *PostItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Postable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePostFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Postable), nil
}
// InReplyTo the inReplyTo property
func (m *PostItemRequestBuilder) InReplyTo()(*ia669f29093557a86be3b54973c4aba4e78b09891dc40e9d967d02b3938cbb08d.InReplyToRequestBuilder) {
    return ia669f29093557a86be3b54973c4aba4e78b09891dc40e9d967d02b3938cbb08d.NewInReplyToRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *PostItemRequestBuilder) MultiValueExtendedProperties()(*ie5a656c04186903bcf4da8c7768bc7acdefb45b9b92af375c844ab6dde924c3b.MultiValueExtendedPropertiesRequestBuilder) {
    return ie5a656c04186903bcf4da8c7768bc7acdefb45b9b92af375c844ab6dde924c3b.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.conversations.item.threads.item.posts.item.multiValueExtendedProperties.item collection
func (m *PostItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i81cf5d0e294523ab35bef33098c9a475ca60e40c2f945ab2e768942447194086.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i81cf5d0e294523ab35bef33098c9a475ca60e40c2f945ab2e768942447194086.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property posts in groups
func (m *PostItemRequestBuilder) Patch(options *PostItemRequestBuilderPatchOptions)(error) {
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
// Reply the reply property
func (m *PostItemRequestBuilder) Reply()(*i524bf6311b24b6f141d885728d606289b8ee24cbd67c3599ccb2ec6084cb17bd.ReplyRequestBuilder) {
    return i524bf6311b24b6f141d885728d606289b8ee24cbd67c3599ccb2ec6084cb17bd.NewReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *PostItemRequestBuilder) SingleValueExtendedProperties()(*i9884a55ef28596c5da82f51922aa27afafadf6bc4977af30511de7bda6744049.SingleValueExtendedPropertiesRequestBuilder) {
    return i9884a55ef28596c5da82f51922aa27afafadf6bc4977af30511de7bda6744049.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.conversations.item.threads.item.posts.item.singleValueExtendedProperties.item collection
func (m *PostItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i85387a3ef1c4b59c2db207f407864b3a4fca529975d344ab83cd21fd615de2e6.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i85387a3ef1c4b59c2db207f407864b3a4fca529975d344ab83cd21fd615de2e6.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
