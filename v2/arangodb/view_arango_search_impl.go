//
// DISCLAIMER
//
// Copyright 2023 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany

package arangodb

import (
	"context"
	"net/http"

	"github.com/pkg/errors"

	"github.com/arangodb/go-driver/v2/arangodb/shared"
	"github.com/arangodb/go-driver/v2/connection"
)

// viewArangoSearch implements ArangoSearchView
type viewArangoSearch struct {
	*view
}

// Properties fetches extended information about the view.
func (v *viewArangoSearch) Properties(ctx context.Context) (ArangoSearchViewProperties, error) {
	url := v.db.url("_api", "view", v.name, "properties")

	var response struct {
		shared.ResponseStruct `json:",inline"`
		ArangoSearchViewProperties
	}

	resp, err := connection.CallGet(ctx, v.db.connection(), url, &response)
	if err != nil {
		return ArangoSearchViewProperties{}, errors.WithStack(err)
	}

	switch code := resp.Code(); code {
	case http.StatusOK:
		return response.ArangoSearchViewProperties, nil
	default:
		return ArangoSearchViewProperties{}, response.AsArangoErrorWithCode(code)
	}
}

// SetProperties changes properties of the view.
func (v *viewArangoSearch) SetProperties(ctx context.Context, options ArangoSearchViewProperties) error {
	url := v.db.url("_api", "view", v.name, "properties")
	var response struct {
		shared.ResponseStruct `json:",inline"`
	}

	resp, err := connection.CallPut(ctx, v.db.connection(), url, &response, options)
	if err != nil {
		return errors.WithStack(err)
	}

	switch code := resp.Code(); code {
	case http.StatusOK:
		return nil
	default:
		return response.AsArangoErrorWithCode(code)
	}
}
