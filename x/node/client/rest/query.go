package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
)

func queryNodeHandlerFunc(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func queryNodesHandlerFunc(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
