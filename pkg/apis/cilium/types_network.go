// Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cilium

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IdentityAllocationMode selects how identities are shared between cilium
// nodes by setting how they are stored. The options are "crd" or "kvstore".
type IdentityAllocationMode string

const (
	// CRD defines the crd IdentityAllocationMode type.
	CRD IdentityAllocationMode = "crd"
	// KVStore defines the kvstore IdentityAllocationMode type.
	KVStore IdentityAllocationMode = "kvstore"
)

// TunnelMode defines what tunnel mode to use for Cilium.
type TunnelMode string

const (
	// VXLan defines the vxlan tunnel mode
	VXLan TunnelMode = "vxlan"
	// Geneve defines the geneve tunnel mode.
	Geneve TunnelMode = "geneve"
	// Disabled defines the disabled tunnel mode.
	Disabled TunnelMode = "disabled"
)

// KubeProxyReplacementMode defines which mode should kube-proxy run in.
// More infromation here: https://docs.cilium.io/en/v1.7/gettingstarted/kubeproxy-free/
type KubeProxyReplacementMode string

const (
	// Strict defines the strict kube-proxy replacement mode
	Strict KubeProxyReplacementMode = "strict"
	// Probe defines the probe kube-proxy replacement mode
	Probe KubeProxyReplacementMode = "probe"
	// Partial defines the partial kube-proxy replacement mode
	Partial KubeProxyReplacementMode = "partial"
)

// NodePortMode defines how NodePort services are enabled.
type NodePortMode string

const (
	Hybird NodePortMode = "hybrid"
)

// Store defines the kubernetes storage backend
type Store string

const (
	// Kubernetes defines the kubernetes CRD store type
	Kubernetes Store = "kubernetes"
	// ETCD defines the ETCD store type
	ETCD Store = "etcd"
)

// InstallIPTableRules configuration for cilium
type InstallIPTableRules struct {
	Enabled bool
}

// ExternalIPs configuration for cilium
type ExternalIP struct {
	// ExternalIPenabled is used to define whether ExternalIP address is required or not.
	Enabled bool
}

// Hubble enablement for cilium
type Hubble struct {
	// Enabled defines whether hubble will be enabled for the cluster.
	Enabled bool
	// UI defines whether Hubble UI is enabled or not.
	UI bool
	// Metrics defined what metrics will be reported by hubble
	Metrics []string
}

// Nodeport enablement for cilium
type Nodeport struct {
	// Enabled is used to define whether Nodeport is required or not.
	Enabled bool
	// Mode is the mode of NodePort feature
	Mode NodePortMode
}

// KubeProxy configuration for cilium
type KubeProxy struct {
	// Enabled specifies whether kubeproxy is disabled.
	Enabled *bool
	// ServiceHost specify the controlplane node IP Address.
	ServiceHost *string
	// ServicePort specify the kube-apiserver port number.
	ServicePort *int32
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetworkConfig is a struct representing the configmap for the cilium
// networking plugin
type NetworkConfig struct {
	metav1.TypeMeta
	// Debug configuration to be enabled or not
	Debug *bool
	// PSPEnabled configuration
	PSPEnabled *bool
	// KubeProxy configuration to be enabled or not
	KubeProxy *KubeProxy
	// Hubble configuration to be enabled or not
	Hubble *Hubble
	// TunnelMode configuration, it should be 'vxlan', 'geneve' or 'disabled'
	TunnelMode *TunnelMode
	// Store can be either Kubernetes or etcd.
	Store *Store
}
