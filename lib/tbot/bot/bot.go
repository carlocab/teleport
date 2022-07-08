/*
Copyright 2022 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package bot

import (
	"context"

	"github.com/gravitational/teleport/api/client/proto"
	"github.com/gravitational/teleport/api/client/webclient"
	"github.com/gravitational/teleport/lib/auth"
)

// B is an interface covering various public B methods to circumvent
// import cycle issues.
type B interface {
	// AuthPing pings the auth server and returns the (possibly cached) response.
	AuthPing(ctx context.Context) (*proto.PingResponse, error)

	// ProxyPing returns a (possibly cached) ping response from the Teleport proxy.
	// Note that it relies on the auth server being configured with a sane proxy
	// public address.
	ProxyPing(ctx context.Context) (*webclient.PingResponse, error)

	// Client retrieves the current auth client.
	Client() auth.ClientI
}
