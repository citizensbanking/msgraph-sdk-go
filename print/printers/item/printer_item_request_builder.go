package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i05f5154d5dc8e1096ef3ba18fb5bdb890712250391686b0ffd44f043e05db1fb "github.com/microsoftgraph/msgraph-sdk-go/print/printers/item/tasktriggers"
    i6ef3a73ca24cdfe4f01546d744e3ad5d5f17082ec1283caf5b2176d392cb9cd9 "github.com/microsoftgraph/msgraph-sdk-go/print/printers/item/connectors"
    i74b470ddc01590a04dfc5ba245b7ccb094a2f7b819d0a4ea94773869da1de8bc "github.com/microsoftgraph/msgraph-sdk-go/print/printers/item/shares"
    i95e425711e2e59b85807bfbd36bd18a7893a62957fccbb0dcc17647c9c3f5e90 "github.com/microsoftgraph/msgraph-sdk-go/print/printers/item/restorefactorydefaults"
    i02acec2b11a647c1fc0601a6c9f6283bf36064facb17fdc38860e7fc02ca0db0 "github.com/microsoftgraph/msgraph-sdk-go/print/printers/item/connectors/item"
    i652aa29a7f8828f592dc5bf3e9ada513fbcda2a67c9d060f001cc99e28d86159 "github.com/microsoftgraph/msgraph-sdk-go/print/printers/item/shares/item"
    i8a35ad54635f07016096adf7157efc92ede0057cd4a790d06a32332ce94f22cc "github.com/microsoftgraph/msgraph-sdk-go/print/printers/item/tasktriggers/item"
)

// PrinterItemRequestBuilder provides operations to manage the printers property of the microsoft.graph.print entity.
type PrinterItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// PrinterItemRequestBuilderDeleteOptions options for Delete
type PrinterItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// PrinterItemRequestBuilderGetOptions options for Get
type PrinterItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *PrinterItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// PrinterItemRequestBuilderGetQueryParameters the list of printers registered in the tenant.
type PrinterItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// PrinterItemRequestBuilderPatchOptions options for Patch
type PrinterItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Printerable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// Connectors the connectors property
func (m *PrinterItemRequestBuilder) Connectors()(*i6ef3a73ca24cdfe4f01546d744e3ad5d5f17082ec1283caf5b2176d392cb9cd9.ConnectorsRequestBuilder) {
    return i6ef3a73ca24cdfe4f01546d744e3ad5d5f17082ec1283caf5b2176d392cb9cd9.NewConnectorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConnectorsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.print.printers.item.connectors.item collection
func (m *PrinterItemRequestBuilder) ConnectorsById(id string)(*i02acec2b11a647c1fc0601a6c9f6283bf36064facb17fdc38860e7fc02ca0db0.PrintConnectorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printConnector_id"] = id
    }
    return i02acec2b11a647c1fc0601a6c9f6283bf36064facb17fdc38860e7fc02ca0db0.NewPrintConnectorItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewPrinterItemRequestBuilderInternal instantiates a new PrinterItemRequestBuilder and sets the default values.
func NewPrinterItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrinterItemRequestBuilder) {
    m := &PrinterItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/print/printers/{printer_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPrinterItemRequestBuilder instantiates a new PrinterItemRequestBuilder and sets the default values.
func NewPrinterItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrinterItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrinterItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property printers for print
func (m *PrinterItemRequestBuilder) CreateDeleteRequestInformation(options *PrinterItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the list of printers registered in the tenant.
func (m *PrinterItemRequestBuilder) CreateGetRequestInformation(options *PrinterItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property printers in print
func (m *PrinterItemRequestBuilder) CreatePatchRequestInformation(options *PrinterItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property printers for print
func (m *PrinterItemRequestBuilder) Delete(options *PrinterItemRequestBuilderDeleteOptions)(error) {
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
// Get the list of printers registered in the tenant.
func (m *PrinterItemRequestBuilder) Get(options *PrinterItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Printerable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePrinterFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Printerable), nil
}
// Patch update the navigation property printers in print
func (m *PrinterItemRequestBuilder) Patch(options *PrinterItemRequestBuilderPatchOptions)(error) {
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
// RestoreFactoryDefaults the restoreFactoryDefaults property
func (m *PrinterItemRequestBuilder) RestoreFactoryDefaults()(*i95e425711e2e59b85807bfbd36bd18a7893a62957fccbb0dcc17647c9c3f5e90.RestoreFactoryDefaultsRequestBuilder) {
    return i95e425711e2e59b85807bfbd36bd18a7893a62957fccbb0dcc17647c9c3f5e90.NewRestoreFactoryDefaultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Shares the shares property
func (m *PrinterItemRequestBuilder) Shares()(*i74b470ddc01590a04dfc5ba245b7ccb094a2f7b819d0a4ea94773869da1de8bc.SharesRequestBuilder) {
    return i74b470ddc01590a04dfc5ba245b7ccb094a2f7b819d0a4ea94773869da1de8bc.NewSharesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.print.printers.item.shares.item collection
func (m *PrinterItemRequestBuilder) SharesById(id string)(*i652aa29a7f8828f592dc5bf3e9ada513fbcda2a67c9d060f001cc99e28d86159.PrinterShareItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printerShare_id"] = id
    }
    return i652aa29a7f8828f592dc5bf3e9ada513fbcda2a67c9d060f001cc99e28d86159.NewPrinterShareItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TaskTriggers the taskTriggers property
func (m *PrinterItemRequestBuilder) TaskTriggers()(*i05f5154d5dc8e1096ef3ba18fb5bdb890712250391686b0ffd44f043e05db1fb.TaskTriggersRequestBuilder) {
    return i05f5154d5dc8e1096ef3ba18fb5bdb890712250391686b0ffd44f043e05db1fb.NewTaskTriggersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaskTriggersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.print.printers.item.taskTriggers.item collection
func (m *PrinterItemRequestBuilder) TaskTriggersById(id string)(*i8a35ad54635f07016096adf7157efc92ede0057cd4a790d06a32332ce94f22cc.PrintTaskTriggerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printTaskTrigger_id"] = id
    }
    return i8a35ad54635f07016096adf7157efc92ede0057cd4a790d06a32332ce94f22cc.NewPrintTaskTriggerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
