// -------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// --------------------------------------------------------------------------------------------

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzureBackendPool is the resource AGIC is watching on for any backend IPs change
type AzureBackendPool struct {
	metav1.TypeMeta `json:",inline"`

	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
    // +optional
	Spec AzureBackendPoolSpec `json:"spec"`
}

// AzureBackendPoolSpec defines the info object of backend pool
type AzureBackendPoolSpec struct {
	metav1.TypeMeta `json:",inline"`

	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// backend pool info including id and ip addresses
	backendPoolInfo []BackendPool `json:"backendPoolInfo"`
}

type BackendPool struct {
	backendPoolID string   `json:"backendPoolID"`
	ipAddresses   []string `json:"ipAddresses,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzureBackendPoolList is the list of backend pool
type AzureBackendPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []AzureBackendPool `json:"items"`
}
