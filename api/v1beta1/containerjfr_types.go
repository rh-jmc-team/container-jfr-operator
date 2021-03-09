// Copyright (c) 2021 Red Hat, Inc.
//
// The Universal Permissive License (UPL), Version 1.0
//
// Subject to the condition set forth below, permission is hereby granted to any
// person obtaining a copy of this software, associated documentation and/or data
// (collectively the "Software"), free of charge and under any and all copyright
// rights in the Software, and any and all patent rights owned or freely
// licensable by each licensor hereunder covering either (i) the unmodified
// Software as contributed to or provided by such licensor, or (ii) the Larger
// Works (as defined below), to deal in both
//
// (a) the Software, and
// (b) any piece of software and/or hardware listed in the lrgrwrks.txt file if
// one is included with the Software (each a "Larger Work" to which the Software
// is contributed by such licensors),
//
// without restriction, including without limitation the rights to copy, create
// derivative works of, display, perform, and distribute the Software and make,
// use, sell, offer for sale, import, export, have made, and have sold the
// Software and the Larger Work(s), and to sublicense the foregoing rights on
// either these or other terms.
//
// This license is subject to the following condition:
// The above copyright notice and either this complete permission notice or at
// a minimum a reference to the UPL must be included in all copies or
// substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ContainerJFRSpec defines the desired state of ContainerJFR
type ContainerJFRSpec struct {
	Minimal bool `json:"minimal"`
	// List of certificates to enable tls when connecting to targets
	TrustedCertSecrets []CertificateSecret `json:"trustedCertSecrets"`
}

// ContainerJFRStatus defines the observed state of ContainerJFR
type ContainerJFRStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:path=containerjfrs,scope=Namespaced

// ContainerJFR is the Schema for the containerjfrs API
type ContainerJFR struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ContainerJFRSpec   `json:"spec,omitempty"`
	Status ContainerJFRStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ContainerJFRList contains a list of ContainerJFR
type ContainerJFRList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContainerJFR `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ContainerJFR{}, &ContainerJFRList{})
}

// DefaultCertificateKey will be used when looking up the certificate within a secret,
// if a key is not manually specified
const DefaultCertificateKey = corev1.TLSCertKey

type CertificateSecret struct {
	// Name of secret in the local namespace
	SecretName string `json:"secretName"`
	// Key within secret containing the certificate
	// +optional
	CertificateKey *string `json:"certificateKey,omitempty"`
}