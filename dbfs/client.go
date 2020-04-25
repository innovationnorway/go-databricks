// Package dbfs implements the Azure ARM Dbfs service API version 0.1.0.
//
// Databricks REST API
package dbfs

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
// DefaultBaseURI is the default URI used for the service Dbfs
DefaultBaseURI = "/api/2.0")

// BaseClient is the base client for Dbfs.
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
                 Constraints: []validation.Constraint{	{Target: "body.Path", Name: validation.Null, Rule: true, Chain: nil }}}}); err != nil {
                return result, validation.NewError("dbfs.BaseClient", "Delete", err.Error())
                }

                    req, err := client.DeletePreparer(ctx, body)
        if err != nil {
        err = autorest.NewErrorWithError(err, "dbfs.BaseClient", "Delete", nil , "Failure preparing request")
        return
        }

                resp, err := client.DeleteSender(req)
                if err != nil {
                result.Response = resp
                err = autorest.NewErrorWithError(err, "dbfs.BaseClient", "Delete", resp, "Failure sending request")
                return
                }

                result, err = client.DeleteResponder(resp)
                if err != nil {
                err = autorest.NewErrorWithError(err, "dbfs.BaseClient", "Delete", resp, "Failure responding to request")
                }

        return
        }

        // DeletePreparer prepares the Delete request.
        func (client BaseClient) DeletePreparer(ctx context.Context, body DeleteAttributes) (*http.Request, error) {
            preparer := autorest.CreatePreparer(
        autorest.AsContentType("application/json; charset=utf-8"),
        autorest.AsPost(),
        autorest.WithBaseURL(client.BaseURI),
        autorest.WithPath("/dbfs/delete"),
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

    // GetStatus sends the get status request.
    func (client BaseClient) GetStatus(ctx context.Context, pathParameter string) (result GetStatusResult, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/BaseClient.GetStatus")
            defer func() {
                sc := -1
                if result.Response.Response != nil {
                    sc = result.Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
        }
            req, err := client.GetStatusPreparer(ctx, pathParameter)
        if err != nil {
        err = autorest.NewErrorWithError(err, "dbfs.BaseClient", "GetStatus", nil , "Failure preparing request")
        return
        }

                resp, err := client.GetStatusSender(req)
                if err != nil {
                result.Response = autorest.Response{Response: resp}
                err = autorest.NewErrorWithError(err, "dbfs.BaseClient", "GetStatus", resp, "Failure sending request")
                return
                }

                result, err = client.GetStatusResponder(resp)
                if err != nil {
                err = autorest.NewErrorWithError(err, "dbfs.BaseClient", "GetStatus", resp, "Failure responding to request")
                }

        return
        }

        // GetStatusPreparer prepares the GetStatus request.
        func (client BaseClient) GetStatusPreparer(ctx context.Context, pathParameter string) (*http.Request, error) {
                    queryParameters := map[string]interface{} {
            "path": autorest.Encode("query",pathParameter),
            }

            preparer := autorest.CreatePreparer(
        autorest.AsGet(),
        autorest.WithBaseURL(client.BaseURI),
        autorest.WithPath("/dbfs/get-status"),
        autorest.WithQueryParameters(queryParameters))
        return preparer.Prepare((&http.Request{}).WithContext(ctx))
        }

        // GetStatusSender sends the GetStatus request. The method will close the
        // http.Response Body if it receives an error.
        func (client BaseClient) GetStatusSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
                }

    // GetStatusResponder handles the response to the GetStatus request. The method always
    // closes the http.Response Body.
    func (client BaseClient) GetStatusResponder(resp *http.Response) (result GetStatusResult, err error) {
        err = autorest.Respond(
        resp,
        client.ByInspecting(),
        azure.WithErrorUnlessStatusCode(http.StatusOK),
        autorest.ByUnmarshallingJSON(&result),
        autorest.ByClosing())
        result.Response = autorest.Response{Response: resp}
            return
        }

    // Mkdirs sends the mkdirs request.
    func (client BaseClient) Mkdirs(ctx context.Context, body MkdirsAttributes) (result autorest.Response, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/BaseClient.Mkdirs")
            defer func() {
                sc := -1
                if result.Response != nil {
                    sc = result.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
        }
            req, err := client.MkdirsPreparer(ctx, body)
        if err != nil {
        err = autorest.NewErrorWithError(err, "dbfs.BaseClient", "Mkdirs", nil , "Failure preparing request")
        return
        }

                resp, err := client.MkdirsSender(req)
                if err != nil {
                result.Response = resp
                err = autorest.NewErrorWithError(err, "dbfs.BaseClient", "Mkdirs", resp, "Failure sending request")
                return
                }

                result, err = client.MkdirsResponder(resp)
                if err != nil {
                err = autorest.NewErrorWithError(err, "dbfs.BaseClient", "Mkdirs", resp, "Failure responding to request")
                }

        return
        }

        // MkdirsPreparer prepares the Mkdirs request.
        func (client BaseClient) MkdirsPreparer(ctx context.Context, body MkdirsAttributes) (*http.Request, error) {
            preparer := autorest.CreatePreparer(
        autorest.AsContentType("application/json; charset=utf-8"),
        autorest.AsPost(),
        autorest.WithBaseURL(client.BaseURI),
        autorest.WithPath("/dbfs/mkdirs"),
        autorest.WithJSON(body))
        return preparer.Prepare((&http.Request{}).WithContext(ctx))
        }

        // MkdirsSender sends the Mkdirs request. The method will close the
        // http.Response Body if it receives an error.
        func (client BaseClient) MkdirsSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
                }

    // MkdirsResponder handles the response to the Mkdirs request. The method always
    // closes the http.Response Body.
    func (client BaseClient) MkdirsResponder(resp *http.Response) (result autorest.Response, err error) {
        err = autorest.Respond(
        resp,
        client.ByInspecting(),
        azure.WithErrorUnlessStatusCode(http.StatusOK),
        autorest.ByClosing())
        result.Response = resp
            return
        }

    // Put sends the put request.
    func (client BaseClient) Put(ctx context.Context, body PutAttributes) (result autorest.Response, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/BaseClient.Put")
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
                 Constraints: []validation.Constraint{	{Target: "body.Path", Name: validation.Null, Rule: true, Chain: nil }}}}); err != nil {
                return result, validation.NewError("dbfs.BaseClient", "Put", err.Error())
                }

                    req, err := client.PutPreparer(ctx, body)
        if err != nil {
        err = autorest.NewErrorWithError(err, "dbfs.BaseClient", "Put", nil , "Failure preparing request")
        return
        }

                resp, err := client.PutSender(req)
                if err != nil {
                result.Response = resp
                err = autorest.NewErrorWithError(err, "dbfs.BaseClient", "Put", resp, "Failure sending request")
                return
                }

                result, err = client.PutResponder(resp)
                if err != nil {
                err = autorest.NewErrorWithError(err, "dbfs.BaseClient", "Put", resp, "Failure responding to request")
                }

        return
        }

        // PutPreparer prepares the Put request.
        func (client BaseClient) PutPreparer(ctx context.Context, body PutAttributes) (*http.Request, error) {
            preparer := autorest.CreatePreparer(
        autorest.AsContentType("application/json; charset=utf-8"),
        autorest.AsPost(),
        autorest.WithBaseURL(client.BaseURI),
        autorest.WithPath("/dbfs/put"),
        autorest.WithJSON(body))
        return preparer.Prepare((&http.Request{}).WithContext(ctx))
        }

        // PutSender sends the Put request. The method will close the
        // http.Response Body if it receives an error.
        func (client BaseClient) PutSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
                }

    // PutResponder handles the response to the Put request. The method always
    // closes the http.Response Body.
    func (client BaseClient) PutResponder(resp *http.Response) (result autorest.Response, err error) {
        err = autorest.Respond(
        resp,
        client.ByInspecting(),
        azure.WithErrorUnlessStatusCode(http.StatusOK),
        autorest.ByClosing())
        result.Response = resp
            return
        }

