// Copyright 2019 Intel Corporation and Smart-Edge.com, Inc. All rights reserved
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import "encoding/json"

// NotificationDescriptor describes a type used in EAA API
type NotificationDescriptor struct {
	// Name of notification
	Name string `json:"name,omitempty"`
	// Version of notification
	Version string `json:"version,omitempty"`
	// Human readable description of notification
	Description string `json:"description,omitempty"`
}

// NotificationToConsumer describes a type used in EAA API
type NotificationToConsumer struct {
	// Name of notification
	Name string `json:"name,omitempty"`
	// Version of notification
	Version string `json:"version,omitempty"`
	// The payload can be any JSON object with a name
	// and version-specific schema.
	Payload json.RawMessage `json:"payload,omitempty"`
	// URN of the producer
	URN URN `json:"producer,omitempty"`
}

// ServiceList JSON struct
type ServiceList struct {
	Services []Service `json:"services,omitempty"`
}

// Service JSON struct
type Service struct {
	URN           *URN                     `json:"urn,omitempty"`
	Description   string                   `json:"description,omitempty"`
	EndpointURI   string                   `json:"endpoint_uri,omitempty"`
	Status        string                   `json:"status,omitempty"`
	Notifications []NotificationDescriptor `json:"notifications,omitempty"`
}

// SubscriptionList JSON struct
type SubscriptionList struct {
	Subscriptions []Subscription `json:"subscriptions,omitempty"`
}

// Subscription describes a type used in EAA API
type Subscription struct {

	// The name of the producer app. The unique ID is optional for
	// subscribing and if not given will subscribe to any producer in the
	// namespace.
	URN *URN `json:"urn,omitempty"`

	// The list of all notification types registered by all producers in
	// this namespace.
	Notifications []NotificationDescriptor `json:"notifications,omitempty"`
}

// URN describes a type used in EAA API
type URN struct {

	// The per-namespace unique portion of the URN that when appended to
	// the namespace with a separator forms the complete URN.
	ID string `json:"id,omitempty"`

	// The non-unique portion of the URN that identifies the class excluding
	// a trailing separator.
	Namespace string `json:"namespace,omitempty"`
}

// AuthCredentials defines a response for a request to obtain authentication
// credentials. These credentials may be used to further communicate with
// endpoint(s) that are protected by a form of authentication.
//
// Any strings that are annotated as "PEM-encoded" implies that encoding format
// is used, with any newlines indicated with `\n` characters. Most languages
// provide encoders that correctly marshal this out. For more information,
// see the RFC here: https://tools.ietf.org/html/rfc7468
type AuthCredentials struct {
	ID          string   `json:"id,omitempty"`
	Certificate string   `json:"certificate,omitempty"`
	CaChain     []string `json:"ca_chain,omitempty"`
	CaPool      []string `json:"ca_pool,omitempty"`
}

// AuthIdentity defines a request to obtain authentication credentials. These
// credentials would be used to further communicate with endpoint(s) that are
// protected by a form of authentication.
//
// Any strings that are annotated as "PEM-encoded" implies that encoding format
// is used, with any newlines indicated with `\n` characters. Most languages
// provide encoders that correctly marshal this out. For more information,
// see the RFC here: https://tools.ietf.org/html/rfc7468
type AuthIdentity struct {
	Csr string `json:"csr,omitempty"`
}
