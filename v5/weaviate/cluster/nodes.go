package cluster

import (
	"context"
	"net/http"

	"github.com/semi-technologies/weaviate-go-client/v5/weaviate/connection"
	"github.com/semi-technologies/weaviate-go-client/v5/weaviate/except"
	"github.com/semi-technologies/weaviate/entities/models"
)

type NodesStatusGetter struct {
	connection *connection.Connection
}

// Do get the nodes endpoint
func (nsg *NodesStatusGetter) Do(ctx context.Context) (*models.NodesStatusResponse, error) {
	responseData, responseErr := nsg.connection.RunREST(ctx, "/nodes", http.MethodGet, nil)
	err := except.CheckResponseDataErrorAndStatusCode(responseData, responseErr, 200)
	if err != nil {
		return nil, err
	}
	var nodesStatus models.NodesStatusResponse
	parseErr := responseData.DecodeBodyIntoTarget(&nodesStatus)
	return &nodesStatus, parseErr
}
