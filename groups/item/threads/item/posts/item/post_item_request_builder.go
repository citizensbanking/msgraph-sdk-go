package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i1d5e0e0253764c43a7170c4dae6d6dada63a8577871ccfb6dbfa89ffcbb240cd "github.com/microsoftgraph/msgraph-sdk-go/groups/item/threads/item/posts/item/multivalueextendedproperties"
    i5134eaa0bf73502180126762f8dba15d49dd36677713ba778154e0a60592a5bd "github.com/microsoftgraph/msgraph-sdk-go/groups/item/threads/item/posts/item/forward"
    i68a4055bc841e0456a2e6639a547657de33eb9b48f925db253682599604a51e3 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/threads/item/posts/item/singlevalueextendedproperties"
    i75174cbcb1b898c11858f5f96b2d8847c330a38149bcd9be83e30ad77ed43c22 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/threads/item/posts/item/attachments"
    i87cc194ce5c8b6a9311a66fb5a22e3ba337812de9de5c80802d09bdb94b863aa "github.com/microsoftgraph/msgraph-sdk-go/groups/item/threads/item/posts/item/reply"
    ic5c310ab811a9efd80ec215346a10a39c86ebea5d6fadba8de5ac53f120fbdc1 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/threads/item/posts/item/extensions"
    ie143dd9268acc57eb633fcf1a20193d5695c0268106b5f598eb2c2f1e7a8a611 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/threads/item/posts/item/inreplyto"
    i27611e4379a9836b89c70ec5ef6aa8f3e16e6cc0deccd831272023f099a54b9a "github.com/microsoftgraph/msgraph-sdk-go/groups/item/threads/item/posts/item/multivalueextendedproperties/item"
    i8ec098680e04f7df4b9ae597c1037f94b0a4442db8e1280e22d5710ee0602d81 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/threads/item/posts/item/attachments/item"
    ie25007fdca9d41586081045acb368264715f3e0d2b48bebac108a93b0d11781a "github.com/microsoftgraph/msgraph-sdk-go/groups/item/threads/item/posts/item/extensions/item"
    if0f66d4e50c563fdb7e3c99373c63bc365617d543a82f7c19e2cfc84e3dcb3a8 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/threads/item/posts/item/singlevalueextendedproperties/item"
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
func (m *PostItemRequestBuilder) Attachments()(*i75174cbcb1b898c11858f5f96b2d8847c330a38149bcd9be83e30ad77ed43c22.AttachmentsRequestBuilder) {
    return i75174cbcb1b898c11858f5f96b2d8847c330a38149bcd9be83e30ad77ed43c22.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.threads.item.posts.item.attachments.item collection
func (m *PostItemRequestBuilder) AttachmentsById(id string)(*i8ec098680e04f7df4b9ae597c1037f94b0a4442db8e1280e22d5710ee0602d81.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i8ec098680e04f7df4b9ae597c1037f94b0a4442db8e1280e22d5710ee0602d81.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewPostItemRequestBuilderInternal instantiates a new PostItemRequestBuilder and sets the default values.
func NewPostItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PostItemRequestBuilder) {
    m := &PostItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/threads/{conversationThread_id}/posts/{post_id}{?select,expand}";
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
func (m *PostItemRequestBuilder) Extensions()(*ic5c310ab811a9efd80ec215346a10a39c86ebea5d6fadba8de5ac53f120fbdc1.ExtensionsRequestBuilder) {
    return ic5c310ab811a9efd80ec215346a10a39c86ebea5d6fadba8de5ac53f120fbdc1.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.threads.item.posts.item.extensions.item collection
func (m *PostItemRequestBuilder) ExtensionsById(id string)(*ie25007fdca9d41586081045acb368264715f3e0d2b48bebac108a93b0d11781a.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return ie25007fdca9d41586081045acb368264715f3e0d2b48bebac108a93b0d11781a.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *PostItemRequestBuilder) Forward()(*i5134eaa0bf73502180126762f8dba15d49dd36677713ba778154e0a60592a5bd.ForwardRequestBuilder) {
    return i5134eaa0bf73502180126762f8dba15d49dd36677713ba778154e0a60592a5bd.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *PostItemRequestBuilder) InReplyTo()(*ie143dd9268acc57eb633fcf1a20193d5695c0268106b5f598eb2c2f1e7a8a611.InReplyToRequestBuilder) {
    return ie143dd9268acc57eb633fcf1a20193d5695c0268106b5f598eb2c2f1e7a8a611.NewInReplyToRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *PostItemRequestBuilder) MultiValueExtendedProperties()(*i1d5e0e0253764c43a7170c4dae6d6dada63a8577871ccfb6dbfa89ffcbb240cd.MultiValueExtendedPropertiesRequestBuilder) {
    return i1d5e0e0253764c43a7170c4dae6d6dada63a8577871ccfb6dbfa89ffcbb240cd.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.threads.item.posts.item.multiValueExtendedProperties.item collection
func (m *PostItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i27611e4379a9836b89c70ec5ef6aa8f3e16e6cc0deccd831272023f099a54b9a.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i27611e4379a9836b89c70ec5ef6aa8f3e16e6cc0deccd831272023f099a54b9a.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *PostItemRequestBuilder) Reply()(*i87cc194ce5c8b6a9311a66fb5a22e3ba337812de9de5c80802d09bdb94b863aa.ReplyRequestBuilder) {
    return i87cc194ce5c8b6a9311a66fb5a22e3ba337812de9de5c80802d09bdb94b863aa.NewReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *PostItemRequestBuilder) SingleValueExtendedProperties()(*i68a4055bc841e0456a2e6639a547657de33eb9b48f925db253682599604a51e3.SingleValueExtendedPropertiesRequestBuilder) {
    return i68a4055bc841e0456a2e6639a547657de33eb9b48f925db253682599604a51e3.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.threads.item.posts.item.singleValueExtendedProperties.item collection
func (m *PostItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*if0f66d4e50c563fdb7e3c99373c63bc365617d543a82f7c19e2cfc84e3dcb3a8.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return if0f66d4e50c563fdb7e3c99373c63bc365617d543a82f7c19e2cfc84e3dcb3a8.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
