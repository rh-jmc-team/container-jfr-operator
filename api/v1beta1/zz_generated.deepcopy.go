// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateSecret) DeepCopyInto(out *CertificateSecret) {
	*out = *in
	if in.CertificateKey != nil {
		in, out := &in.CertificateKey, &out.CertificateKey
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateSecret.
func (in *CertificateSecret) DeepCopy() *CertificateSecret {
	if in == nil {
		return nil
	}
	out := new(CertificateSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerJFR) DeepCopyInto(out *ContainerJFR) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerJFR.
func (in *ContainerJFR) DeepCopy() *ContainerJFR {
	if in == nil {
		return nil
	}
	out := new(ContainerJFR)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ContainerJFR) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerJFRList) DeepCopyInto(out *ContainerJFRList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ContainerJFR, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerJFRList.
func (in *ContainerJFRList) DeepCopy() *ContainerJFRList {
	if in == nil {
		return nil
	}
	out := new(ContainerJFRList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ContainerJFRList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerJFRSpec) DeepCopyInto(out *ContainerJFRSpec) {
	*out = *in
	if in.TrustedCertSecrets != nil {
		in, out := &in.TrustedCertSecrets, &out.TrustedCertSecrets
		*out = make([]CertificateSecret, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StorageOptions != nil {
		in, out := &in.StorageOptions, &out.StorageOptions
		*out = new(StorageConfiguration)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerJFRSpec.
func (in *ContainerJFRSpec) DeepCopy() *ContainerJFRSpec {
	if in == nil {
		return nil
	}
	out := new(ContainerJFRSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerJFRStatus) DeepCopyInto(out *ContainerJFRStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerJFRStatus.
func (in *ContainerJFRStatus) DeepCopy() *ContainerJFRStatus {
	if in == nil {
		return nil
	}
	out := new(ContainerJFRStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventInfo) DeepCopyInto(out *EventInfo) {
	*out = *in
	if in.Category != nil {
		in, out := &in.Category, &out.Category
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make(map[string]OptionDescriptor, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventInfo.
func (in *EventInfo) DeepCopy() *EventInfo {
	if in == nil {
		return nil
	}
	out := new(EventInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlightRecorder) DeepCopyInto(out *FlightRecorder) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlightRecorder.
func (in *FlightRecorder) DeepCopy() *FlightRecorder {
	if in == nil {
		return nil
	}
	out := new(FlightRecorder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlightRecorder) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlightRecorderList) DeepCopyInto(out *FlightRecorderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FlightRecorder, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlightRecorderList.
func (in *FlightRecorderList) DeepCopy() *FlightRecorderList {
	if in == nil {
		return nil
	}
	out := new(FlightRecorderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlightRecorderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlightRecorderSpec) DeepCopyInto(out *FlightRecorderSpec) {
	*out = *in
	if in.RecordingSelector != nil {
		in, out := &in.RecordingSelector, &out.RecordingSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.JMXCredentials != nil {
		in, out := &in.JMXCredentials, &out.JMXCredentials
		*out = new(JMXAuthSecret)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlightRecorderSpec.
func (in *FlightRecorderSpec) DeepCopy() *FlightRecorderSpec {
	if in == nil {
		return nil
	}
	out := new(FlightRecorderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlightRecorderStatus) DeepCopyInto(out *FlightRecorderStatus) {
	*out = *in
	if in.Events != nil {
		in, out := &in.Events, &out.Events
		*out = make([]EventInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Templates != nil {
		in, out := &in.Templates, &out.Templates
		*out = make([]TemplateInfo, len(*in))
		copy(*out, *in)
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(v1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlightRecorderStatus.
func (in *FlightRecorderStatus) DeepCopy() *FlightRecorderStatus {
	if in == nil {
		return nil
	}
	out := new(FlightRecorderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JMXAuthSecret) DeepCopyInto(out *JMXAuthSecret) {
	*out = *in
	if in.UsernameKey != nil {
		in, out := &in.UsernameKey, &out.UsernameKey
		*out = new(string)
		**out = **in
	}
	if in.PasswordKey != nil {
		in, out := &in.PasswordKey, &out.PasswordKey
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JMXAuthSecret.
func (in *JMXAuthSecret) DeepCopy() *JMXAuthSecret {
	if in == nil {
		return nil
	}
	out := new(JMXAuthSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OptionDescriptor) DeepCopyInto(out *OptionDescriptor) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OptionDescriptor.
func (in *OptionDescriptor) DeepCopy() *OptionDescriptor {
	if in == nil {
		return nil
	}
	out := new(OptionDescriptor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistentVolumeClaimConfig) DeepCopyInto(out *PersistentVolumeClaimConfig) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(v1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistentVolumeClaimConfig.
func (in *PersistentVolumeClaimConfig) DeepCopy() *PersistentVolumeClaimConfig {
	if in == nil {
		return nil
	}
	out := new(PersistentVolumeClaimConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Recording) DeepCopyInto(out *Recording) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Recording.
func (in *Recording) DeepCopy() *Recording {
	if in == nil {
		return nil
	}
	out := new(Recording)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Recording) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecordingList) DeepCopyInto(out *RecordingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Recording, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecordingList.
func (in *RecordingList) DeepCopy() *RecordingList {
	if in == nil {
		return nil
	}
	out := new(RecordingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RecordingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecordingSpec) DeepCopyInto(out *RecordingSpec) {
	*out = *in
	if in.EventOptions != nil {
		in, out := &in.EventOptions, &out.EventOptions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Duration = in.Duration
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(RecordingState)
		**out = **in
	}
	if in.FlightRecorder != nil {
		in, out := &in.FlightRecorder, &out.FlightRecorder
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecordingSpec.
func (in *RecordingSpec) DeepCopy() *RecordingSpec {
	if in == nil {
		return nil
	}
	out := new(RecordingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecordingStatus) DeepCopyInto(out *RecordingStatus) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(RecordingState)
		**out = **in
	}
	in.StartTime.DeepCopyInto(&out.StartTime)
	out.Duration = in.Duration
	if in.DownloadURL != nil {
		in, out := &in.DownloadURL, &out.DownloadURL
		*out = new(string)
		**out = **in
	}
	if in.ReportURL != nil {
		in, out := &in.ReportURL, &out.ReportURL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecordingStatus.
func (in *RecordingStatus) DeepCopy() *RecordingStatus {
	if in == nil {
		return nil
	}
	out := new(RecordingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageConfiguration) DeepCopyInto(out *StorageConfiguration) {
	*out = *in
	if in.PVC != nil {
		in, out := &in.PVC, &out.PVC
		*out = new(PersistentVolumeClaimConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageConfiguration.
func (in *StorageConfiguration) DeepCopy() *StorageConfiguration {
	if in == nil {
		return nil
	}
	out := new(StorageConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateInfo) DeepCopyInto(out *TemplateInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateInfo.
func (in *TemplateInfo) DeepCopy() *TemplateInfo {
	if in == nil {
		return nil
	}
	out := new(TemplateInfo)
	in.DeepCopyInto(out)
	return out
}
