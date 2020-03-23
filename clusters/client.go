// Package clusters implements the Azure ARM Clusters service API version 0.1.0.
//
// Databricks REST API
package clusters

// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
    "context"
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/azure"
    "github.com/Azure/go-autorest/autorest/validation"
    "github.com/Azure/go-autorest/tracing"
    "net/http"
)

const (
// DefaultBaseURI is the default URI used for the service Clusters
DefaultBaseURI = "/api/2.0")

// BaseClient is the base client for Clusters.
type BaseClient struct {
    autorest.Client
    BaseURI string
}

// New creates an instance of the BaseClient client.
func New()BaseClient {
    return NewWithBaseURI(DefaultBaseURI, )
}

// NewWithBaseURI creates an instance of the BaseClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWithBaseURI(baseURI string, ) BaseClient {
    return BaseClient{
        Client: autorest.NewClientWithUserAgent(UserAgent()),
        BaseURI: baseURI,
    }
}

    // Create sends the create request.
    func (client BaseClient) Create(ctx context.Context, body Attributes) (result CreateResult, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/BaseClient.Create")
            defer func() {
                sc := -1
                if result.Response.Response != nil {
                    sc = result.Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
        }
                if err := validation.Validate([]validation.Validation{
                { TargetValue: body,
                 Constraints: []validation.Constraint{	{Target: "body.Autoscale", Name: validation.Null, Rule: false ,
                Chain: []validation.Constraint{	{Target: "body.Autoscale.MinWorkers", Name: validation.Null, Rule: true, Chain: nil },
                	{Target: "body.Autoscale.MaxWorkers", Name: validation.Null, Rule: true, Chain: nil },
                }},
                	{Target: "body.SparkVersion", Name: validation.Null, Rule: true, Chain: nil },
                	{Target: "body.NodeTypeID", Name: validation.Null, Rule: true, Chain: nil }}}}); err != nil {
                return result, validation.NewError("clusters.BaseClient", "Create", err.Error())
                }

                    req, err := client.CreatePreparer(ctx, body)
        if err != nil {
        err = autorest.NewErrorWithError(err, "clusters.BaseClient", "Create", nil , "Failure preparing request")
        return
        }

                resp, err := client.CreateSender(req)
                if err != nil {
                result.Response = autorest.Response{Response: resp}
                err = autorest.NewErrorWithError(err, "clusters.BaseClient", "Create", resp, "Failure sending request")
                return
                }

                result, err = client.CreateResponder(resp)
                if err != nil {
                err = autorest.NewErrorWithError(err, "clusters.BaseClient", "Create", resp, "Failure responding to request")
                }

        return
        }

        // CreatePreparer prepares the Create request.
        func (client BaseClient) CreatePreparer(ctx context.Context, body Attributes) (*http.Request, error) {
            preparer := autorest.CreatePreparer(
        autorest.AsContentType("application/json; charset=utf-8"),
        autorest.AsPost(),
        autorest.WithBaseURL(client.BaseURI),
        autorest.WithPath("/clusters/create"),
        autorest.WithJSON(body))
        return preparer.Prepare((&http.Request{}).WithContext(ctx))
        }

        // CreateSender sends the Create request. The method will close the
        // http.Response Body if it receives an error.
        func (client BaseClient) CreateSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
                }

    // CreateResponder handles the response to the Create request. The method always
    // closes the http.Response Body.
    func (client BaseClient) CreateResponder(resp *http.Response) (result CreateResult, err error) {
        err = autorest.Respond(
        resp,
        client.ByInspecting(),
        azure.WithErrorUnlessStatusCode(http.StatusOK),
        autorest.ByUnmarshallingJSON(&result),
        autorest.ByClosing())
        result.Response = autorest.Response{Response: resp}
            return
        }

    // Delete sends the delete request.
    func (client BaseClient) Delete(ctx context.Context, body DeleteAttributes) (result autorest.Response, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/BaseClient.Delete")
            defer func() {
                sc := -1
                if result.Response != nil {
                    sc = result.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
        }
                if err := validation.Validate([]validation.Validation{
                { TargetValue: body,
                 Constraints: []validation.Constraint{	{Target: "body.ClusterID", Name: validation.Null, Rule: true, Chain: nil }}}}); err != nil {
                return result, validation.NewError("clusters.BaseClient", "Delete", err.Error())
                }

                    req, err := client.DeletePreparer(ctx, body)
        if err != nil {
        err = autorest.NewErrorWithError(err, "clusters.BaseClient", "Delete", nil , "Failure preparing request")
        return
        }

                resp, err := client.DeleteSender(req)
                if err != nil {
                result.Response = resp
                err = autorest.NewErrorWithError(err, "clusters.BaseClient", "Delete", resp, "Failure sending request")
                return
                }

                result, err = client.DeleteResponder(resp)
                if err != nil {
                err = autorest.NewErrorWithError(err, "clusters.BaseClient", "Delete", resp, "Failure responding to request")
                }

        return
        }

        // DeletePreparer prepares the Delete request.
        func (client BaseClient) DeletePreparer(ctx context.Context, body DeleteAttributes) (*http.Request, error) {
            preparer := autorest.CreatePreparer(
        autorest.AsContentType("application/json; charset=utf-8"),
        autorest.AsPost(),
        autorest.WithBaseURL(client.BaseURI),
        autorest.WithPath("/clusters/delete"),
        autorest.WithJSON(body))
        return preparer.Prepare((&http.Request{}).WithContext(ctx))
        }

        // DeleteSender sends the Delete request. The method will close the
        // http.Response Body if it receives an error.
        func (client BaseClient) DeleteSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
                }

    // DeleteResponder handles the response to the Delete request. The method always
    // closes the http.Response Body.
    func (client BaseClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
        err = autorest.Respond(
        resp,
        client.ByInspecting(),
        azure.WithErrorUnlessStatusCode(http.StatusOK),
        autorest.ByClosing())
        result.Response = resp
            return
        }

    // Edit sends the edit request.
    func (client BaseClient) Edit(ctx context.Context, body EditAttributes) (result autorest.Response, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/BaseClient.Edit")
            defer func() {
                sc := -1
                if result.Response != nil {
                    sc = result.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
        }
                if err := validation.Validate([]validation.Validation{
                { TargetValue: body,
                 Constraints: []validation.Constraint{	{Target: "body.Autoscale", Name: validation.Null, Rule: false ,
                Chain: []validation.Constraint{	{Target: "body.Autoscale.MinWorkers", Name: validation.Null, Rule: true, Chain: nil },
                	{Target: "body.Autoscale.MaxWorkers", Name: validation.Null, Rule: true, Chain: nil },
                }},
                	{Target: "body.SparkVersion", Name: validation.Null, Rule: true, Chain: nil },
                	{Target: "body.NodeTypeID", Name: validation.Null, Rule: true, Chain: nil }}}}); err != nil {
                return result, validation.NewError("clusters.BaseClient", "Edit", err.Error())
                }

                    req, err := client.EditPreparer(ctx, body)
        if err != nil {
        err = autorest.NewErrorWithError(err, "clusters.BaseClient", "Edit", nil , "Failure preparing request")
        return
        }

                resp, err := client.EditSender(req)
                if err != nil {
                result.Response = resp
                err = autorest.NewErrorWithError(err, "clusters.BaseClient", "Edit", resp, "Failure sending request")
                return
                }

                result, err = client.EditResponder(resp)
                if err != nil {
                err = autorest.NewErrorWithError(err, "clusters.BaseClient", "Edit", resp, "Failure responding to request")
                }

        return
        }

        // EditPreparer prepares the Edit request.
        func (client BaseClient) EditPreparer(ctx context.Context, body EditAttributes) (*http.Request, error) {
            preparer := autorest.CreatePreparer(
        autorest.AsContentType("application/json; charset=utf-8"),
        autorest.AsPost(),
        autorest.WithBaseURL(client.BaseURI),
        autorest.WithPath("/clusters/edit"),
        autorest.WithJSON(body))
        return preparer.Prepare((&http.Request{}).WithContext(ctx))
        }

        // EditSender sends the Edit request. The method will close the
        // http.Response Body if it receives an error.
        func (client BaseClient) EditSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
                }

    // EditResponder handles the response to the Edit request. The method always
    // closes the http.Response Body.
    func (client BaseClient) EditResponder(resp *http.Response) (result autorest.Response, err error) {
        err = autorest.Respond(
        resp,
        client.ByInspecting(),
        azure.WithErrorUnlessStatusCode(http.StatusOK),
        autorest.ByClosing())
        result.Response = resp
            return
        }

    // Get sends the get request.
    func (client BaseClient) Get(ctx context.Context, clusterID string) (result Info, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/BaseClient.Get")
            defer func() {
                sc := -1
                if result.Response.Response != nil {
                    sc = result.Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
        }
            req, err := client.GetPreparer(ctx, clusterID)
        if err != nil {
        err = autorest.NewErrorWithError(err, "clusters.BaseClient", "Get", nil , "Failure preparing request")
        return
        }

                resp, err := client.GetSender(req)
                if err != nil {
                result.Response = autorest.Response{Response: resp}
                err = autorest.NewErrorWithError(err, "clusters.BaseClient", "Get", resp, "Failure sending request")
                return
                }

                result, err = client.GetResponder(resp)
                if err != nil {
                err = autorest.NewErrorWithError(err, "clusters.BaseClient", "Get", resp, "Failure responding to request")
                }

        return
        }

        // GetPreparer prepares the Get request.
        func (client BaseClient) GetPreparer(ctx context.Context, clusterID string) (*http.Request, error) {
                    queryParameters := map[string]interface{} {
            "cluster_id": clusterID,
            }

            preparer := autorest.CreatePreparer(
        autorest.AsGet(),
        autorest.WithBaseURL(client.BaseURI),
        autorest.WithPath("/clusters/get"),
        autorest.WithQueryParameters(queryParameters))
        return preparer.Prepare((&http.Request{}).WithContext(ctx))
        }

        // GetSender sends the Get request. The method will close the
        // http.Response Body if it receives an error.
        func (client BaseClient) GetSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
                }

    // GetResponder handles the response to the Get request. The method always
    // closes the http.Response Body.
    func (client BaseClient) GetResponder(resp *http.Response) (result Info, err error) {
        err = autorest.Respond(
        resp,
        client.ByInspecting(),
        azure.WithErrorUnlessStatusCode(http.StatusOK),
        autorest.ByUnmarshallingJSON(&result),
        autorest.ByClosing())
        result.Response = autorest.Response{Response: resp}
            return
        }

    // List sends the list request.
    func (client BaseClient) List(ctx context.Context) (result ListInfo, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/BaseClient.List")
            defer func() {
                sc := -1
                if result.Response.Response != nil {
                    sc = result.Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
        }
            req, err := client.ListPreparer(ctx)
        if err != nil {
        err = autorest.NewErrorWithError(err, "clusters.BaseClient", "List", nil , "Failure preparing request")
        return
        }

                resp, err := client.ListSender(req)
                if err != nil {
                result.Response = autorest.Response{Response: resp}
                err = autorest.NewErrorWithError(err, "clusters.BaseClient", "List", resp, "Failure sending request")
                return
                }

                result, err = client.ListResponder(resp)
                if err != nil {
                err = autorest.NewErrorWithError(err, "clusters.BaseClient", "List", resp, "Failure responding to request")
                }

        return
        }

        // ListPreparer prepares the List request.
        func (client BaseClient) ListPreparer(ctx context.Context) (*http.Request, error) {
            preparer := autorest.CreatePreparer(
        autorest.AsGet(),
        autorest.WithBaseURL(client.BaseURI),
        autorest.WithPath("/clusters/list"))
        return preparer.Prepare((&http.Request{}).WithContext(ctx))
        }

        // ListSender sends the List request. The method will close the
        // http.Response Body if it receives an error.
        func (client BaseClient) ListSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
                }

    // ListResponder handles the response to the List request. The method always
    // closes the http.Response Body.
    func (client BaseClient) ListResponder(resp *http.Response) (result ListInfo, err error) {
        err = autorest.Respond(
        resp,
        client.ByInspecting(),
        azure.WithErrorUnlessStatusCode(http.StatusOK),
        autorest.ByUnmarshallingJSON(&result.Value),
        autorest.ByClosing())
        result.Response = autorest.Response{Response: resp}
            return
        }

    // ListNodeTypes sends the list node types request.
    func (client BaseClient) ListNodeTypes(ctx context.Context) (result ListNodeType, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/BaseClient.ListNodeTypes")
            defer func() {
                sc := -1
                if result.Response.Response != nil {
                    sc = result.Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
        }
            req, err := client.ListNodeTypesPreparer(ctx)
        if err != nil {
        err = autorest.NewErrorWithError(err, "clusters.BaseClient", "ListNodeTypes", nil , "Failure preparing request")
        return
        }

                resp, err := client.ListNodeTypesSender(req)
                if err != nil {
                result.Response = autorest.Response{Response: resp}
                err = autorest.NewErrorWithError(err, "clusters.BaseClient", "ListNodeTypes", resp, "Failure sending request")
                return
                }

                result, err = client.ListNodeTypesResponder(resp)
                if err != nil {
                err = autorest.NewErrorWithError(err, "clusters.BaseClient", "ListNodeTypes", resp, "Failure responding to request")
                }

        return
        }

        // ListNodeTypesPreparer prepares the ListNodeTypes request.
        func (client BaseClient) ListNodeTypesPreparer(ctx context.Context) (*http.Request, error) {
            preparer := autorest.CreatePreparer(
        autorest.AsGet(),
        autorest.WithBaseURL(client.BaseURI),
        autorest.WithPath("/clusters/list-node-types"))
        return preparer.Prepare((&http.Request{}).WithContext(ctx))
        }

        // ListNodeTypesSender sends the ListNodeTypes request. The method will close the
        // http.Response Body if it receives an error.
        func (client BaseClient) ListNodeTypesSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
                }

    // ListNodeTypesResponder handles the response to the ListNodeTypes request. The method always
    // closes the http.Response Body.
    func (client BaseClient) ListNodeTypesResponder(resp *http.Response) (result ListNodeType, err error) {
        err = autorest.Respond(
        resp,
        client.ByInspecting(),
        azure.WithErrorUnlessStatusCode(http.StatusOK),
        autorest.ByUnmarshallingJSON(&result.Value),
        autorest.ByClosing())
        result.Response = autorest.Response{Response: resp}
            return
        }

    // PermanentDelete sends the permanent delete request.
    func (client BaseClient) PermanentDelete(ctx context.Context, body PermanentDeleteAttributes) (result autorest.Response, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/BaseClient.PermanentDelete")
            defer func() {
                sc := -1
                if result.Response != nil {
                    sc = result.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
        }
                if err := validation.Validate([]validation.Validation{
                { TargetValue: body,
                 Constraints: []validation.Constraint{	{Target: "body.ClusterID", Name: validation.Null, Rule: true, Chain: nil }}}}); err != nil {
                return result, validation.NewError("clusters.BaseClient", "PermanentDelete", err.Error())
                }

                    req, err := client.PermanentDeletePreparer(ctx, body)
        if err != nil {
        err = autorest.NewErrorWithError(err, "clusters.BaseClient", "PermanentDelete", nil , "Failure preparing request")
        return
        }

                resp, err := client.PermanentDeleteSender(req)
                if err != nil {
                result.Response = resp
                err = autorest.NewErrorWithError(err, "clusters.BaseClient", "PermanentDelete", resp, "Failure sending request")
                return
                }

                result, err = client.PermanentDeleteResponder(resp)
                if err != nil {
                err = autorest.NewErrorWithError(err, "clusters.BaseClient", "PermanentDelete", resp, "Failure responding to request")
                }

        return
        }

        // PermanentDeletePreparer prepares the PermanentDelete request.
        func (client BaseClient) PermanentDeletePreparer(ctx context.Context, body PermanentDeleteAttributes) (*http.Request, error) {
            preparer := autorest.CreatePreparer(
        autorest.AsContentType("application/json; charset=utf-8"),
        autorest.AsPost(),
        autorest.WithBaseURL(client.BaseURI),
        autorest.WithPath("/clusters/permanent-delete"),
        autorest.WithJSON(body))
        return preparer.Prepare((&http.Request{}).WithContext(ctx))
        }

        // PermanentDeleteSender sends the PermanentDelete request. The method will close the
        // http.Response Body if it receives an error.
        func (client BaseClient) PermanentDeleteSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
                }

    // PermanentDeleteResponder handles the response to the PermanentDelete request. The method always
    // closes the http.Response Body.
    func (client BaseClient) PermanentDeleteResponder(resp *http.Response) (result autorest.Response, err error) {
        err = autorest.Respond(
        resp,
        client.ByInspecting(),
        azure.WithErrorUnlessStatusCode(http.StatusOK),
        autorest.ByClosing())
        result.Response = resp
            return
        }

    // Resize sends the resize request.
    func (client BaseClient) Resize(ctx context.Context, body ResizeAttributes) (result autorest.Response, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/BaseClient.Resize")
            defer func() {
                sc := -1
                if result.Response != nil {
                    sc = result.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
        }
                if err := validation.Validate([]validation.Validation{
                { TargetValue: body,
                 Constraints: []validation.Constraint{	{Target: "body.Autoscale", Name: validation.Null, Rule: false ,
                Chain: []validation.Constraint{	{Target: "body.Autoscale.MinWorkers", Name: validation.Null, Rule: true, Chain: nil },
                	{Target: "body.Autoscale.MaxWorkers", Name: validation.Null, Rule: true, Chain: nil },
                }},
                	{Target: "body.ClusterID", Name: validation.Null, Rule: true, Chain: nil }}}}); err != nil {
                return result, validation.NewError("clusters.BaseClient", "Resize", err.Error())
                }

                    req, err := client.ResizePreparer(ctx, body)
        if err != nil {
        err = autorest.NewErrorWithError(err, "clusters.BaseClient", "Resize", nil , "Failure preparing request")
        return
        }

                resp, err := client.ResizeSender(req)
                if err != nil {
                result.Response = resp
                err = autorest.NewErrorWithError(err, "clusters.BaseClient", "Resize", resp, "Failure sending request")
                return
                }

                result, err = client.ResizeResponder(resp)
                if err != nil {
                err = autorest.NewErrorWithError(err, "clusters.BaseClient", "Resize", resp, "Failure responding to request")
                }

        return
        }

        // ResizePreparer prepares the Resize request.
        func (client BaseClient) ResizePreparer(ctx context.Context, body ResizeAttributes) (*http.Request, error) {
            preparer := autorest.CreatePreparer(
        autorest.AsContentType("application/json; charset=utf-8"),
        autorest.AsPost(),
        autorest.WithBaseURL(client.BaseURI),
        autorest.WithPath("/clusters/resize"),
        autorest.WithJSON(body))
        return preparer.Prepare((&http.Request{}).WithContext(ctx))
        }

        // ResizeSender sends the Resize request. The method will close the
        // http.Response Body if it receives an error.
        func (client BaseClient) ResizeSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
                }

    // ResizeResponder handles the response to the Resize request. The method always
    // closes the http.Response Body.
    func (client BaseClient) ResizeResponder(resp *http.Response) (result autorest.Response, err error) {
        err = autorest.Respond(
        resp,
        client.ByInspecting(),
        azure.WithErrorUnlessStatusCode(http.StatusOK),
        autorest.ByClosing())
        result.Response = resp
            return
        }

    // Restart sends the restart request.
    func (client BaseClient) Restart(ctx context.Context, body RestartAttributes) (result autorest.Response, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/BaseClient.Restart")
            defer func() {
                sc := -1
                if result.Response != nil {
                    sc = result.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
        }
                if err := validation.Validate([]validation.Validation{
                { TargetValue: body,
                 Constraints: []validation.Constraint{	{Target: "body.ClusterID", Name: validation.Null, Rule: true, Chain: nil }}}}); err != nil {
                return result, validation.NewError("clusters.BaseClient", "Restart", err.Error())
                }

                    req, err := client.RestartPreparer(ctx, body)
        if err != nil {
        err = autorest.NewErrorWithError(err, "clusters.BaseClient", "Restart", nil , "Failure preparing request")
        return
        }

                resp, err := client.RestartSender(req)
                if err != nil {
                result.Response = resp
                err = autorest.NewErrorWithError(err, "clusters.BaseClient", "Restart", resp, "Failure sending request")
                return
                }

                result, err = client.RestartResponder(resp)
                if err != nil {
                err = autorest.NewErrorWithError(err, "clusters.BaseClient", "Restart", resp, "Failure responding to request")
                }

        return
        }

        // RestartPreparer prepares the Restart request.
        func (client BaseClient) RestartPreparer(ctx context.Context, body RestartAttributes) (*http.Request, error) {
            preparer := autorest.CreatePreparer(
        autorest.AsContentType("application/json; charset=utf-8"),
        autorest.AsPost(),
        autorest.WithBaseURL(client.BaseURI),
        autorest.WithPath("/clusters/restart"),
        autorest.WithJSON(body))
        return preparer.Prepare((&http.Request{}).WithContext(ctx))
        }

        // RestartSender sends the Restart request. The method will close the
        // http.Response Body if it receives an error.
        func (client BaseClient) RestartSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
                }

    // RestartResponder handles the response to the Restart request. The method always
    // closes the http.Response Body.
    func (client BaseClient) RestartResponder(resp *http.Response) (result autorest.Response, err error) {
        err = autorest.Respond(
        resp,
        client.ByInspecting(),
        azure.WithErrorUnlessStatusCode(http.StatusOK),
        autorest.ByClosing())
        result.Response = resp
            return
        }
