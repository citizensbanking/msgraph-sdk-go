package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i078dc43cfacb492d7c14c400159d76effb3ddc82fd0975278340abf380ca4ac0 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/createforward"
    i1e2e330cf63776672ee46cb29c0931a99d88b111c325d0eddcba469f590284ef "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/value"
    i283292ee0714f6d62978117a50ae89004640243f72c98fb77b3d3ad5d895724b "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/multivalueextendedproperties"
    i46afd8dc1f68253c5ec58fffe851f31d4606c9896eeca162270731a770163159 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/move"
    i4d24d04db91a658e0a0039257207563b182ae368b5564bc2651941a177fcec1f "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/send"
    i5a8c865ef21d4e26d59de283ac0bf75d8fa86b7ef5610dbf01669ac1a6a6426c "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/singlevalueextendedproperties"
    i601006ca5c8d4a2bd32a308480c958aa565af8b7eef89ca92de125fddcd8df37 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/createreply"
    i8bf596c0443d82bd098088cf313820c77518a4903e0996592bf7aa9bb0fc6204 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/forward"
    ia2b7d74de2d2ae40f811c9f5ca231ae8bcd6383bd423cbdadabf72fbdbfb6e62 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/replyall"
    ia3fdc47a4b68aebb3754bf14fd0396a90c0336bfcb6d210afc93388c25c3f2df "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/calendarsharingmessage"
    ia7dce7bbbad8834f61a47bf8d2d11a325ca6cb60f92e7e712275f7e8c25fc693 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/reply"
    ib2ca11a51f3340d7f385f534ea5827066380a80aa53f7fb73586d468aad2cf61 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/createreplyall"
    id3a837a6f0f6a0a991f73cd274bed4374b20b7ab621b89f0c8b54728f1438971 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/copy"
    id871d310ec2178d7aee89ca6f65864dd837e1389228be08e4040bb3486b367d4 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/extensions"
    if72e21ea0481f36a049b51f38179e467041e8723702278110ed61ad1e65dde50 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/attachments"
    i5b5457ef2553e64079400434de96e0e9776d9caa96dc2683cdee41584b29dcf5 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/multivalueextendedproperties/item"
    ic067f029b1404c524dee30760b61043b1957ec2a9b5b5f2858424190465ad31a "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/singlevalueextendedproperties/item"
    ie43dfbb21dca01399bcbaa4461736852466b5a2f62ecd895499c91d040c14a01 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/attachments/item"
    ief0529f669908aba2ad97bcc88a9d64281988927f10a7814f2b993f88439be1b "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/extensions/item"
)

// MessageItemRequestBuilder provides operations to manage the messages property of the microsoft.graph.mailFolder entity.
type MessageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// MessageItemRequestBuilderDeleteOptions options for Delete
type MessageItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// MessageItemRequestBuilderGetOptions options for Get
type MessageItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *MessageItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// MessageItemRequestBuilderGetQueryParameters the collection of messages in the mailFolder.
type MessageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// MessageItemRequestBuilderPatchOptions options for Patch
type MessageItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// Attachments the attachments property
func (m *MessageItemRequestBuilder) Attachments()(*if72e21ea0481f36a049b51f38179e467041e8723702278110ed61ad1e65dde50.AttachmentsRequestBuilder) {
    return if72e21ea0481f36a049b51f38179e467041e8723702278110ed61ad1e65dde50.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.mailFolders.item.childFolders.item.messages.item.attachments.item collection
func (m *MessageItemRequestBuilder) AttachmentsById(id string)(*ie43dfbb21dca01399bcbaa4461736852466b5a2f62ecd895499c91d040c14a01.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return ie43dfbb21dca01399bcbaa4461736852466b5a2f62ecd895499c91d040c14a01.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CalendarSharingMessage the calendarSharingMessage property
func (m *MessageItemRequestBuilder) CalendarSharingMessage()(*ia3fdc47a4b68aebb3754bf14fd0396a90c0336bfcb6d210afc93388c25c3f2df.CalendarSharingMessageRequestBuilder) {
    return ia3fdc47a4b68aebb3754bf14fd0396a90c0336bfcb6d210afc93388c25c3f2df.NewCalendarSharingMessageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMessageItemRequestBuilderInternal instantiates a new MessageItemRequestBuilder and sets the default values.
func NewMessageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MessageItemRequestBuilder) {
    m := &MessageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/mailFolders/{mailFolder_id}/childFolders/{mailFolder_id1}/messages/{message_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMessageItemRequestBuilder instantiates a new MessageItemRequestBuilder and sets the default values.
func NewMessageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MessageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMessageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content the Content property
func (m *MessageItemRequestBuilder) Content()(*i1e2e330cf63776672ee46cb29c0931a99d88b111c325d0eddcba469f590284ef.ContentRequestBuilder) {
    return i1e2e330cf63776672ee46cb29c0931a99d88b111c325d0eddcba469f590284ef.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Copy the copy property
func (m *MessageItemRequestBuilder) Copy()(*id3a837a6f0f6a0a991f73cd274bed4374b20b7ab621b89f0c8b54728f1438971.CopyRequestBuilder) {
    return id3a837a6f0f6a0a991f73cd274bed4374b20b7ab621b89f0c8b54728f1438971.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property messages for users
func (m *MessageItemRequestBuilder) CreateDeleteRequestInformation(options *MessageItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateForward the createForward property
func (m *MessageItemRequestBuilder) CreateForward()(*i078dc43cfacb492d7c14c400159d76effb3ddc82fd0975278340abf380ca4ac0.CreateForwardRequestBuilder) {
    return i078dc43cfacb492d7c14c400159d76effb3ddc82fd0975278340abf380ca4ac0.NewCreateForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the collection of messages in the mailFolder.
func (m *MessageItemRequestBuilder) CreateGetRequestInformation(options *MessageItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property messages in users
func (m *MessageItemRequestBuilder) CreatePatchRequestInformation(options *MessageItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateReply the createReply property
func (m *MessageItemRequestBuilder) CreateReply()(*i601006ca5c8d4a2bd32a308480c958aa565af8b7eef89ca92de125fddcd8df37.CreateReplyRequestBuilder) {
    return i601006ca5c8d4a2bd32a308480c958aa565af8b7eef89ca92de125fddcd8df37.NewCreateReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateReplyAll the createReplyAll property
func (m *MessageItemRequestBuilder) CreateReplyAll()(*ib2ca11a51f3340d7f385f534ea5827066380a80aa53f7fb73586d468aad2cf61.CreateReplyAllRequestBuilder) {
    return ib2ca11a51f3340d7f385f534ea5827066380a80aa53f7fb73586d468aad2cf61.NewCreateReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property messages for users
func (m *MessageItemRequestBuilder) Delete(options *MessageItemRequestBuilderDeleteOptions)(error) {
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
func (m *MessageItemRequestBuilder) Extensions()(*id871d310ec2178d7aee89ca6f65864dd837e1389228be08e4040bb3486b367d4.ExtensionsRequestBuilder) {
    return id871d310ec2178d7aee89ca6f65864dd837e1389228be08e4040bb3486b367d4.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.mailFolders.item.childFolders.item.messages.item.extensions.item collection
func (m *MessageItemRequestBuilder) ExtensionsById(id string)(*ief0529f669908aba2ad97bcc88a9d64281988927f10a7814f2b993f88439be1b.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return ief0529f669908aba2ad97bcc88a9d64281988927f10a7814f2b993f88439be1b.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *MessageItemRequestBuilder) Forward()(*i8bf596c0443d82bd098088cf313820c77518a4903e0996592bf7aa9bb0fc6204.ForwardRequestBuilder) {
    return i8bf596c0443d82bd098088cf313820c77518a4903e0996592bf7aa9bb0fc6204.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the collection of messages in the mailFolder.
func (m *MessageItemRequestBuilder) Get(options *MessageItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMessageFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable), nil
}
// Move the move property
func (m *MessageItemRequestBuilder) Move()(*i46afd8dc1f68253c5ec58fffe851f31d4606c9896eeca162270731a770163159.MoveRequestBuilder) {
    return i46afd8dc1f68253c5ec58fffe851f31d4606c9896eeca162270731a770163159.NewMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *MessageItemRequestBuilder) MultiValueExtendedProperties()(*i283292ee0714f6d62978117a50ae89004640243f72c98fb77b3d3ad5d895724b.MultiValueExtendedPropertiesRequestBuilder) {
    return i283292ee0714f6d62978117a50ae89004640243f72c98fb77b3d3ad5d895724b.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.mailFolders.item.childFolders.item.messages.item.multiValueExtendedProperties.item collection
func (m *MessageItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i5b5457ef2553e64079400434de96e0e9776d9caa96dc2683cdee41584b29dcf5.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i5b5457ef2553e64079400434de96e0e9776d9caa96dc2683cdee41584b29dcf5.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property messages in users
func (m *MessageItemRequestBuilder) Patch(options *MessageItemRequestBuilderPatchOptions)(error) {
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
func (m *MessageItemRequestBuilder) Reply()(*ia7dce7bbbad8834f61a47bf8d2d11a325ca6cb60f92e7e712275f7e8c25fc693.ReplyRequestBuilder) {
    return ia7dce7bbbad8834f61a47bf8d2d11a325ca6cb60f92e7e712275f7e8c25fc693.NewReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReplyAll the replyAll property
func (m *MessageItemRequestBuilder) ReplyAll()(*ia2b7d74de2d2ae40f811c9f5ca231ae8bcd6383bd423cbdadabf72fbdbfb6e62.ReplyAllRequestBuilder) {
    return ia2b7d74de2d2ae40f811c9f5ca231ae8bcd6383bd423cbdadabf72fbdbfb6e62.NewReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Send the send property
func (m *MessageItemRequestBuilder) Send()(*i4d24d04db91a658e0a0039257207563b182ae368b5564bc2651941a177fcec1f.SendRequestBuilder) {
    return i4d24d04db91a658e0a0039257207563b182ae368b5564bc2651941a177fcec1f.NewSendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *MessageItemRequestBuilder) SingleValueExtendedProperties()(*i5a8c865ef21d4e26d59de283ac0bf75d8fa86b7ef5610dbf01669ac1a6a6426c.SingleValueExtendedPropertiesRequestBuilder) {
    return i5a8c865ef21d4e26d59de283ac0bf75d8fa86b7ef5610dbf01669ac1a6a6426c.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.mailFolders.item.childFolders.item.messages.item.singleValueExtendedProperties.item collection
func (m *MessageItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ic067f029b1404c524dee30760b61043b1957ec2a9b5b5f2858424190465ad31a.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ic067f029b1404c524dee30760b61043b1957ec2a9b5b5f2858424190465ad31a.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
