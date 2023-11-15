//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetadataOverride) DeepCopyInto(out *MetadataOverride) {
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetadataOverride.
func (in *MetadataOverride) DeepCopy() *MetadataOverride {
	if in == nil {
		return nil
	}
	out := new(MetadataOverride)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnErrorRemediationStrategy) DeepCopyInto(out *OnErrorRemediationStrategy) {
	*out = *in
	if in.MaxRetries != nil {
		in, out := &in.MaxRetries, &out.MaxRetries
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnErrorRemediationStrategy.
func (in *OnErrorRemediationStrategy) DeepCopy() *OnErrorRemediationStrategy {
	if in == nil {
		return nil
	}
	out := new(OnErrorRemediationStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OverrideRunnerSpec) DeepCopyInto(out *OverrideRunnerSpec) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EnvFrom != nil {
		in, out := &in.EnvFrom, &out.EnvFrom
		*out = make([]v1.EnvFromSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]v1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Metadata.DeepCopyInto(&out.Metadata)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OverrideRunnerSpec.
func (in *OverrideRunnerSpec) DeepCopy() *OverrideRunnerSpec {
	if in == nil {
		return nil
	}
	out := new(OverrideRunnerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemediationStrategy) DeepCopyInto(out *RemediationStrategy) {
	*out = *in
	in.OnError.DeepCopyInto(&out.OnError)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemediationStrategy.
func (in *RemediationStrategy) DeepCopy() *RemediationStrategy {
	if in == nil {
		return nil
	}
	out := new(RemediationStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunHistoryPolicy) DeepCopyInto(out *RunHistoryPolicy) {
	*out = *in
	if in.KeepLastPlanRuns != nil {
		in, out := &in.KeepLastPlanRuns, &out.KeepLastPlanRuns
		*out = new(int)
		**out = **in
	}
	if in.KeepLastApplyRuns != nil {
		in, out := &in.KeepLastApplyRuns, &out.KeepLastApplyRuns
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunHistoryPolicy.
func (in *RunHistoryPolicy) DeepCopy() *RunHistoryPolicy {
	if in == nil {
		return nil
	}
	out := new(RunHistoryPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TerraformConfig) DeepCopyInto(out *TerraformConfig) {
	*out = *in
	in.TerragruntConfig.DeepCopyInto(&out.TerragruntConfig)
	if in.NoLock != nil {
		in, out := &in.NoLock, &out.NoLock
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TerraformConfig.
func (in *TerraformConfig) DeepCopy() *TerraformConfig {
	if in == nil {
		return nil
	}
	out := new(TerraformConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TerraformLayer) DeepCopyInto(out *TerraformLayer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TerraformLayer.
func (in *TerraformLayer) DeepCopy() *TerraformLayer {
	if in == nil {
		return nil
	}
	out := new(TerraformLayer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TerraformLayer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TerraformLayerList) DeepCopyInto(out *TerraformLayerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TerraformLayer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TerraformLayerList.
func (in *TerraformLayerList) DeepCopy() *TerraformLayerList {
	if in == nil {
		return nil
	}
	out := new(TerraformLayerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TerraformLayerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TerraformLayerRepository) DeepCopyInto(out *TerraformLayerRepository) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TerraformLayerRepository.
func (in *TerraformLayerRepository) DeepCopy() *TerraformLayerRepository {
	if in == nil {
		return nil
	}
	out := new(TerraformLayerRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TerraformLayerSpec) DeepCopyInto(out *TerraformLayerSpec) {
	*out = *in
	in.TerraformConfig.DeepCopyInto(&out.TerraformConfig)
	out.Repository = in.Repository
	in.RemediationStrategy.DeepCopyInto(&out.RemediationStrategy)
	in.OverrideRunnerSpec.DeepCopyInto(&out.OverrideRunnerSpec)
	in.RunHistoryPolicy.DeepCopyInto(&out.RunHistoryPolicy)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TerraformLayerSpec.
func (in *TerraformLayerSpec) DeepCopy() *TerraformLayerSpec {
	if in == nil {
		return nil
	}
	out := new(TerraformLayerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TerraformLayerStatus) DeepCopyInto(out *TerraformLayerStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TerraformLayerStatus.
func (in *TerraformLayerStatus) DeepCopy() *TerraformLayerStatus {
	if in == nil {
		return nil
	}
	out := new(TerraformLayerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TerraformPullRequest) DeepCopyInto(out *TerraformPullRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TerraformPullRequest.
func (in *TerraformPullRequest) DeepCopy() *TerraformPullRequest {
	if in == nil {
		return nil
	}
	out := new(TerraformPullRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TerraformPullRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TerraformPullRequestList) DeepCopyInto(out *TerraformPullRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TerraformPullRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TerraformPullRequestList.
func (in *TerraformPullRequestList) DeepCopy() *TerraformPullRequestList {
	if in == nil {
		return nil
	}
	out := new(TerraformPullRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TerraformPullRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TerraformPullRequestSpec) DeepCopyInto(out *TerraformPullRequestSpec) {
	*out = *in
	out.Repository = in.Repository
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TerraformPullRequestSpec.
func (in *TerraformPullRequestSpec) DeepCopy() *TerraformPullRequestSpec {
	if in == nil {
		return nil
	}
	out := new(TerraformPullRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TerraformPullRequestStatus) DeepCopyInto(out *TerraformPullRequestStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TerraformPullRequestStatus.
func (in *TerraformPullRequestStatus) DeepCopy() *TerraformPullRequestStatus {
	if in == nil {
		return nil
	}
	out := new(TerraformPullRequestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TerraformRepository) DeepCopyInto(out *TerraformRepository) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TerraformRepository.
func (in *TerraformRepository) DeepCopy() *TerraformRepository {
	if in == nil {
		return nil
	}
	out := new(TerraformRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TerraformRepository) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TerraformRepositoryList) DeepCopyInto(out *TerraformRepositoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TerraformRepository, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TerraformRepositoryList.
func (in *TerraformRepositoryList) DeepCopy() *TerraformRepositoryList {
	if in == nil {
		return nil
	}
	out := new(TerraformRepositoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TerraformRepositoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TerraformRepositoryRepository) DeepCopyInto(out *TerraformRepositoryRepository) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TerraformRepositoryRepository.
func (in *TerraformRepositoryRepository) DeepCopy() *TerraformRepositoryRepository {
	if in == nil {
		return nil
	}
	out := new(TerraformRepositoryRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TerraformRepositorySpec) DeepCopyInto(out *TerraformRepositorySpec) {
	*out = *in
	out.Repository = in.Repository
	in.TerraformConfig.DeepCopyInto(&out.TerraformConfig)
	in.RemediationStrategy.DeepCopyInto(&out.RemediationStrategy)
	in.OverrideRunnerSpec.DeepCopyInto(&out.OverrideRunnerSpec)
	in.RunHistoryPolicy.DeepCopyInto(&out.RunHistoryPolicy)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TerraformRepositorySpec.
func (in *TerraformRepositorySpec) DeepCopy() *TerraformRepositorySpec {
	if in == nil {
		return nil
	}
	out := new(TerraformRepositorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TerraformRepositoryStatus) DeepCopyInto(out *TerraformRepositoryStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TerraformRepositoryStatus.
func (in *TerraformRepositoryStatus) DeepCopy() *TerraformRepositoryStatus {
	if in == nil {
		return nil
	}
	out := new(TerraformRepositoryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TerraformRun) DeepCopyInto(out *TerraformRun) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TerraformRun.
func (in *TerraformRun) DeepCopy() *TerraformRun {
	if in == nil {
		return nil
	}
	out := new(TerraformRun)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TerraformRun) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TerraformRunLayer) DeepCopyInto(out *TerraformRunLayer) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TerraformRunLayer.
func (in *TerraformRunLayer) DeepCopy() *TerraformRunLayer {
	if in == nil {
		return nil
	}
	out := new(TerraformRunLayer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TerraformRunList) DeepCopyInto(out *TerraformRunList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TerraformRun, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TerraformRunList.
func (in *TerraformRunList) DeepCopy() *TerraformRunList {
	if in == nil {
		return nil
	}
	out := new(TerraformRunList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TerraformRunList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TerraformRunSpec) DeepCopyInto(out *TerraformRunSpec) {
	*out = *in
	out.Layer = in.Layer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TerraformRunSpec.
func (in *TerraformRunSpec) DeepCopy() *TerraformRunSpec {
	if in == nil {
		return nil
	}
	out := new(TerraformRunSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TerraformRunStatus) DeepCopyInto(out *TerraformRunStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TerraformRunStatus.
func (in *TerraformRunStatus) DeepCopy() *TerraformRunStatus {
	if in == nil {
		return nil
	}
	out := new(TerraformRunStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TerragruntConfig) DeepCopyInto(out *TerragruntConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TerragruntConfig.
func (in *TerragruntConfig) DeepCopy() *TerragruntConfig {
	if in == nil {
		return nil
	}
	out := new(TerragruntConfig)
	in.DeepCopyInto(out)
	return out
}
