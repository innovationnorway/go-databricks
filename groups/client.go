// Package groups implements the Azure ARM Groups service API version 0.1.0.
//
// Databricks REST API
package groups

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
// DefaultBaseURI is the default URI used for the service Groups
DefaultBaseURI = "/api/2.0")

// BaseClient is the base client for Groups.
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

    // AddMember sends the add member request.
    func (client BaseClient) AddMember(ctx context.Context, body MemberAttributes) (result autorest.Response, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/BaseClient.AddMember")
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
                 Constraints: []validation.Constraint{	{Target: "body.ParentName", Name: validation.Null, Rule: true, Chain: nil }}}}); err != nil {
                return result, validation.NewError("groups.BaseClient", "AddMember", err.Error())
                }

                    req, err := client.AddMemberPreparer(ctx, body)
        if err != nil {
        err = autorest.NewErrorWithError(err, "groups.BaseClient", "AddMember", nil , "Failure preparing request")
        return
        }

                resp, err := client.AddMemberSender(req)
                if err != nil {
                result.Response = resp
                err = autorest.NewErrorWithError(err, "groups.BaseClient", "AddMember", resp, "Failure sending request")
                return
                }

                result, err = client.AddMemberResponder(resp)
                if err != nil {
                err = autorest.NewErrorWithError(err, "groups.BaseClient", "AddMember", resp, "Failure responding to request")
                }

        return
        }

        // AddMemberPreparer prepares the AddMember request.
        func (client BaseClient) AddMemberPreparer(ctx context.Context, body MemberAttributes) (*http.Request, error) {
            preparer := autorest.CreatePreparer(
        autorest.AsContentType("application/json; charset=utf-8"),
        autorest.AsPost(),
        autorest.WithBaseURL(client.BaseURI),
        autorest.WithPath("/groups/add-member"),
        autorest.WithJSON(body))
        return preparer.Prepare((&http.Request{}).WithContext(ctx))
        }

        // AddMemberSender sends the AddMember request. The method will close the
        // http.Response Body if it receives an error.
        func (client BaseClient) AddMemberSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
                }

    // AddMemberResponder handles the response to the AddMember request. The method always
    // closes the http.Response Body.
    func (client BaseClient) AddMemberResponder(resp *http.Response) (result autorest.Response, err error) {
        err = autorest.Respond(
        resp,
        client.ByInspecting(),
        azure.WithErrorUnlessStatusCode(http.StatusOK),
        autorest.ByClosing())
        result.Response = resp
            return
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
                 Constraints: []validation.Constraint{	{Target: "body.GroupName", Name: validation.Null, Rule: true, Chain: nil }}}}); err != nil {
                return result, validation.NewError("groups.BaseClient", "Create", err.Error())
                }

                    req, err := client.CreatePreparer(ctx, body)
        if err != nil {
        err = autorest.NewErrorWithError(err, "groups.BaseClient", "Create", nil , "Failure preparing request")
        return
        }

                resp, err := client.CreateSender(req)
                if err != nil {
                result.Response = autorest.Response{Response: resp}
                err = autorest.NewErrorWithError(err, "groups.BaseClient", "Create", resp, "Failure sending request")
                return
                }

                result, err = client.CreateResponder(resp)
                if err != nil {
                err = autorest.NewErrorWithError(err, "groups.BaseClient", "Create", resp, "Failure responding to request")
                }

        return
        }

        // CreatePreparer prepares the Create request.
        func (client BaseClient) CreatePreparer(ctx context.Context, body Attributes) (*http.Request, error) {
            preparer := autorest.CreatePreparer(
        autorest.AsContentType("application/json; charset=utf-8"),
        autorest.AsPost(),
        autorest.WithBaseURL(client.BaseURI),
        autorest.WithPath("/groups/create"),
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
            req, err := client.DeletePreparer(ctx, body)
        if err != nil {
        err = autorest.NewErrorWithError(err, "groups.BaseClient", "Delete", nil , "Failure preparing request")
        return
        }

                resp, err := client.DeleteSender(req)
                if err != nil {
                result.Response = resp
                err = autorest.NewErrorWithError(err, "groups.BaseClient", "Delete", resp, "Failure sending request")
                return
                }

                result, err = client.DeleteResponder(resp)
                if err != nil {
                err = autorest.NewErrorWithError(err, "groups.BaseClient", "Delete", resp, "Failure responding to request")
                }

        return
        }

        // DeletePreparer prepares the Delete request.
        func (client BaseClient) DeletePreparer(ctx context.Context, body DeleteAttributes) (*http.Request, error) {
            preparer := autorest.CreatePreparer(
        autorest.AsContentType("application/json; charset=utf-8"),
        autorest.AsPost(),
        autorest.WithBaseURL(client.BaseURI),
        autorest.WithPath("/groups/delete"),
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

    // List sends the list request.
    func (client BaseClient) List(ctx context.Context) (result ListListResult, err error) {
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
        err = autorest.NewErrorWithError(err, "groups.BaseClient", "List", nil , "Failure preparing request")
        return
        }

                resp, err := client.ListSender(req)
                if err != nil {
                result.Response = autorest.Response{Response: resp}
                err = autorest.NewErrorWithError(err, "groups.BaseClient", "List", resp, "Failure sending request")
                return
                }

                result, err = client.ListResponder(resp)
                if err != nil {
                err = autorest.NewErrorWithError(err, "groups.BaseClient", "List", resp, "Failure responding to request")
                }

        return
        }

        // ListPreparer prepares the List request.
        func (client BaseClient) ListPreparer(ctx context.Context) (*http.Request, error) {
            preparer := autorest.CreatePreparer(
        autorest.AsGet(),
        autorest.WithBaseURL(client.BaseURI),
        autorest.WithPath("/groups/list"))
        return preparer.Prepare((&http.Request{}).WithContext(ctx))
        }

        // ListSender sends the List request. The method will close the
        // http.Response Body if it receives an error.
        func (client BaseClient) ListSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
                }

    // ListResponder handles the response to the List request. The method always
    // closes the http.Response Body.
    func (client BaseClient) ListResponder(resp *http.Response) (result ListListResult, err error) {
        err = autorest.Respond(
        resp,
        client.ByInspecting(),
        azure.WithErrorUnlessStatusCode(http.StatusOK),
        autorest.ByUnmarshallingJSON(&result.Value),
        autorest.ByClosing())
        result.Response = autorest.Response{Response: resp}
            return
        }

    // ListMembers sends the list members request.
    func (client BaseClient) ListMembers(ctx context.Context, groupName string) (result ListMembersResult, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/BaseClient.ListMembers")
            defer func() {
                sc := -1
                if result.Response.Response != nil {
                    sc = result.Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
        }
            req, err := client.ListMembersPreparer(ctx, groupName)
        if err != nil {
        err = autorest.NewErrorWithError(err, "groups.BaseClient", "ListMembers", nil , "Failure preparing request")
        return
        }

                resp, err := client.ListMembersSender(req)
                if err != nil {
                result.Response = autorest.Response{Response: resp}
                err = autorest.NewErrorWithError(err, "groups.BaseClient", "ListMembers", resp, "Failure sending request")
                return
                }

                result, err = client.ListMembersResponder(resp)
                if err != nil {
                err = autorest.NewErrorWithError(err, "groups.BaseClient", "ListMembers", resp, "Failure responding to request")
                }

        return
        }

        // ListMembersPreparer prepares the ListMembers request.
        func (client BaseClient) ListMembersPreparer(ctx context.Context, groupName string) (*http.Request, error) {
                    queryParameters := map[string]interface{} {
            "group_name": autorest.Encode("query",groupName),
            }

            preparer := autorest.CreatePreparer(
        autorest.AsGet(),
        autorest.WithBaseURL(client.BaseURI),
        autorest.WithPath("/groups/list-members"),
        autorest.WithQueryParameters(queryParameters))
        return preparer.Prepare((&http.Request{}).WithContext(ctx))
        }

        // ListMembersSender sends the ListMembers request. The method will close the
        // http.Response Body if it receives an error.
        func (client BaseClient) ListMembersSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
                }

    // ListMembersResponder handles the response to the ListMembers request. The method always
    // closes the http.Response Body.
    func (client BaseClient) ListMembersResponder(resp *http.Response) (result ListMembersResult, err error) {
        err = autorest.Respond(
        resp,
        client.ByInspecting(),
        azure.WithErrorUnlessStatusCode(http.StatusOK),
        autorest.ByUnmarshallingJSON(&result),
        autorest.ByClosing())
        result.Response = autorest.Response{Response: resp}
            return
        }

    // ListParents sends the list parents request.
    func (client BaseClient) ListParents(ctx context.Context) (result ListListResult, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/BaseClient.ListParents")
            defer func() {
                sc := -1
                if result.Response.Response != nil {
                    sc = result.Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
        }
            req, err := client.ListParentsPreparer(ctx)
        if err != nil {
        err = autorest.NewErrorWithError(err, "groups.BaseClient", "ListParents", nil , "Failure preparing request")
        return
        }

                resp, err := client.ListParentsSender(req)
                if err != nil {
                result.Response = autorest.Response{Response: resp}
                err = autorest.NewErrorWithError(err, "groups.BaseClient", "ListParents", resp, "Failure sending request")
                return
                }

                result, err = client.ListParentsResponder(resp)
                if err != nil {
                err = autorest.NewErrorWithError(err, "groups.BaseClient", "ListParents", resp, "Failure responding to request")
                }

        return
        }

        // ListParentsPreparer prepares the ListParents request.
        func (client BaseClient) ListParentsPreparer(ctx context.Context) (*http.Request, error) {
            preparer := autorest.CreatePreparer(
        autorest.AsGet(),
        autorest.WithBaseURL(client.BaseURI),
        autorest.WithPath("/groups/list-parents"))
        return preparer.Prepare((&http.Request{}).WithContext(ctx))
        }

        // ListParentsSender sends the ListParents request. The method will close the
        // http.Response Body if it receives an error.
        func (client BaseClient) ListParentsSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
                }

    // ListParentsResponder handles the response to the ListParents request. The method always
    // closes the http.Response Body.
    func (client BaseClient) ListParentsResponder(resp *http.Response) (result ListListResult, err error) {
        err = autorest.Respond(
        resp,
        client.ByInspecting(),
        azure.WithErrorUnlessStatusCode(http.StatusOK),
        autorest.ByUnmarshallingJSON(&result.Value),
        autorest.ByClosing())
        result.Response = autorest.Response{Response: resp}
            return
        }

    // RemoveMember sends the remove member request.
    func (client BaseClient) RemoveMember(ctx context.Context, body MemberAttributes) (result autorest.Response, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/BaseClient.RemoveMember")
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
                 Constraints: []validation.Constraint{	{Target: "body.ParentName", Name: validation.Null, Rule: true, Chain: nil }}}}); err != nil {
                return result, validation.NewError("groups.BaseClient", "RemoveMember", err.Error())
                }

                    req, err := client.RemoveMemberPreparer(ctx, body)
        if err != nil {
        err = autorest.NewErrorWithError(err, "groups.BaseClient", "RemoveMember", nil , "Failure preparing request")
        return
        }

                resp, err := client.RemoveMemberSender(req)
                if err != nil {
                result.Response = resp
                err = autorest.NewErrorWithError(err, "groups.BaseClient", "RemoveMember", resp, "Failure sending request")
                return
                }

                result, err = client.RemoveMemberResponder(resp)
                if err != nil {
                err = autorest.NewErrorWithError(err, "groups.BaseClient", "RemoveMember", resp, "Failure responding to request")
                }

        return
        }

        // RemoveMemberPreparer prepares the RemoveMember request.
        func (client BaseClient) RemoveMemberPreparer(ctx context.Context, body MemberAttributes) (*http.Request, error) {
            preparer := autorest.CreatePreparer(
        autorest.AsContentType("application/json; charset=utf-8"),
        autorest.AsPost(),
        autorest.WithBaseURL(client.BaseURI),
        autorest.WithPath("/groups/remove-member"),
        autorest.WithJSON(body))
        return preparer.Prepare((&http.Request{}).WithContext(ctx))
        }

        // RemoveMemberSender sends the RemoveMember request. The method will close the
        // http.Response Body if it receives an error.
        func (client BaseClient) RemoveMemberSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
                }

    // RemoveMemberResponder handles the response to the RemoveMember request. The method always
    // closes the http.Response Body.
    func (client BaseClient) RemoveMemberResponder(resp *http.Response) (result autorest.Response, err error) {
        err = autorest.Respond(
        resp,
        client.ByInspecting(),
        azure.WithErrorUnlessStatusCode(http.StatusOK),
        autorest.ByClosing())
        result.Response = resp
            return
        }

