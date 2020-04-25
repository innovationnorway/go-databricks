package dbfsapi

// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
    "context"
    "github.com/innovationnorway/go-databricks/dbfs"
    "github.com/Azure/go-autorest/autorest"
)

        // BaseClientAPI contains the set of methods on the BaseClient type.
        type BaseClientAPI interface {
            Delete(ctx context.Context, body dbfs.DeleteAttributes) (result autorest.Response, err error)
            GetStatus(ctx context.Context, pathParameter string) (result dbfs.GetStatusResult, err error)
            Mkdirs(ctx context.Context, body dbfs.MkdirsAttributes) (result autorest.Response, err error)
            Put(ctx context.Context, body dbfs.PutAttributes) (result autorest.Response, err error)
        }

        var _ BaseClientAPI = (*dbfs.BaseClient)(nil)
