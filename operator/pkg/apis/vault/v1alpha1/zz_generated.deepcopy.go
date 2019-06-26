// +build !ignore_autogenerated

// Copyright © 2019 Banzai Cloud
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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSUnsealConfig) DeepCopyInto(out *AWSUnsealConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSUnsealConfig.
func (in *AWSUnsealConfig) DeepCopy() *AWSUnsealConfig {
	if in == nil {
		return nil
	}
	out := new(AWSUnsealConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlibabaUnsealConfig) DeepCopyInto(out *AlibabaUnsealConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlibabaUnsealConfig.
func (in *AlibabaUnsealConfig) DeepCopy() *AlibabaUnsealConfig {
	if in == nil {
		return nil
	}
	out := new(AlibabaUnsealConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureUnsealConfig) DeepCopyInto(out *AzureUnsealConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureUnsealConfig.
func (in *AzureUnsealConfig) DeepCopy() *AzureUnsealConfig {
	if in == nil {
		return nil
	}
	out := new(AzureUnsealConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialsConfig) DeepCopyInto(out *CredentialsConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialsConfig.
func (in *CredentialsConfig) DeepCopy() *CredentialsConfig {
	if in == nil {
		return nil
	}
	out := new(CredentialsConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GoogleUnsealConfig) DeepCopyInto(out *GoogleUnsealConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GoogleUnsealConfig.
func (in *GoogleUnsealConfig) DeepCopy() *GoogleUnsealConfig {
	if in == nil {
		return nil
	}
	out := new(GoogleUnsealConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ingress) DeepCopyInto(out *Ingress) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ingress.
func (in *Ingress) DeepCopy() *Ingress {
	if in == nil {
		return nil
	}
	out := new(Ingress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesUnsealConfig) DeepCopyInto(out *KubernetesUnsealConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesUnsealConfig.
func (in *KubernetesUnsealConfig) DeepCopy() *KubernetesUnsealConfig {
	if in == nil {
		return nil
	}
	out := new(KubernetesUnsealConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resources) DeepCopyInto(out *Resources) {
	*out = *in
	if in.Vault != nil {
		in, out := &in.Vault, &out.Vault
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.BankVaults != nil {
		in, out := &in.BankVaults, &out.BankVaults
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Etcd != nil {
		in, out := &in.Etcd, &out.Etcd
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.PrometheusExporter != nil {
		in, out := &in.PrometheusExporter, &out.PrometheusExporter
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resources.
func (in *Resources) DeepCopy() *Resources {
	if in == nil {
		return nil
	}
	out := new(Resources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnsealConfig) DeepCopyInto(out *UnsealConfig) {
	*out = *in
	out.Options = in.Options
	if in.Kubernetes != nil {
		in, out := &in.Kubernetes, &out.Kubernetes
		*out = new(KubernetesUnsealConfig)
		**out = **in
	}
	if in.Google != nil {
		in, out := &in.Google, &out.Google
		*out = new(GoogleUnsealConfig)
		**out = **in
	}
	if in.Alibaba != nil {
		in, out := &in.Alibaba, &out.Alibaba
		*out = new(AlibabaUnsealConfig)
		**out = **in
	}
	if in.Azure != nil {
		in, out := &in.Azure, &out.Azure
		*out = new(AzureUnsealConfig)
		**out = **in
	}
	if in.AWS != nil {
		in, out := &in.AWS, &out.AWS
		*out = new(AWSUnsealConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnsealConfig.
func (in *UnsealConfig) DeepCopy() *UnsealConfig {
	if in == nil {
		return nil
	}
	out := new(UnsealConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnsealOptions) DeepCopyInto(out *UnsealOptions) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnsealOptions.
func (in *UnsealOptions) DeepCopy() *UnsealOptions {
	if in == nil {
		return nil
	}
	out := new(UnsealOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Vault) DeepCopyInto(out *Vault) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Vault.
func (in *Vault) DeepCopy() *Vault {
	if in == nil {
		return nil
	}
	out := new(Vault)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Vault) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in VaultConfig) DeepCopyInto(out *VaultConfig) {
	{
		in := &in
		*out = in.DeepCopy()
		return
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in VaultExternalConfig) DeepCopyInto(out *VaultExternalConfig) {
	{
		in := &in
		*out = in.DeepCopy()
		return
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultList) DeepCopyInto(out *VaultList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Vault, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultList.
func (in *VaultList) DeepCopy() *VaultList {
	if in == nil {
		return nil
	}
	out := new(VaultList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VaultList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultSpec) DeepCopyInto(out *VaultSpec) {
	*out = *in
	if in.WatchedSecretsLabels != nil {
		in, out := &in.WatchedSecretsLabels, &out.WatchedSecretsLabels
		*out = make([]map[string]string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make(map[string]string, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.VaultAnnotations != nil {
		in, out := &in.VaultAnnotations, &out.VaultAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.VaultLabels != nil {
		in, out := &in.VaultLabels, &out.VaultLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.VaultPodSpec.DeepCopyInto(&out.VaultPodSpec)
	if in.VaultConfigurerAnnotations != nil {
		in, out := &in.VaultConfigurerAnnotations, &out.VaultConfigurerAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.VaultConfigurerLabels != nil {
		in, out := &in.VaultConfigurerLabels, &out.VaultConfigurerLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.VaultConfigurerPodSpec.DeepCopyInto(&out.VaultConfigurerPodSpec)
	out.Config = in.Config.DeepCopy()
	out.ExternalConfig = in.ExternalConfig.DeepCopy()
	in.UnsealConfig.DeepCopyInto(&out.UnsealConfig)
	out.CredentialsConfig = in.CredentialsConfig
	if in.EnvsConfig != nil {
		in, out := &in.EnvsConfig, &out.EnvsConfig
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.SecurityContext.DeepCopyInto(&out.SecurityContext)
	if in.EtcdAnnotations != nil {
		in, out := &in.EtcdAnnotations, &out.EtcdAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EtcdPodAnnotations != nil {
		in, out := &in.EtcdPodAnnotations, &out.EtcdPodAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EtcdPVCSpec != nil {
		in, out := &in.EtcdPVCSpec, &out.EtcdPVCSpec
		*out = new(v1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ServicePorts != nil {
		in, out := &in.ServicePorts, &out.ServicePorts
		*out = make(map[string]int32, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.NodeAffinity.DeepCopyInto(&out.NodeAffinity)
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
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
	if in.VaultEnvsConfig != nil {
		in, out := &in.VaultEnvsConfig, &out.VaultEnvsConfig
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(Resources)
		(*in).DeepCopyInto(*out)
	}
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = new(Ingress)
		(*in).DeepCopyInto(*out)
	}
	if in.TLSExpiryThreshold != nil {
		in, out := &in.TLSExpiryThreshold, &out.TLSExpiryThreshold
		*out = new(time.Duration)
		**out = **in
	}
	if in.CANamespaces != nil {
		in, out := &in.CANamespaces, &out.CANamespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultSpec.
func (in *VaultSpec) DeepCopy() *VaultSpec {
	if in == nil {
		return nil
	}
	out := new(VaultSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultStatus) DeepCopyInto(out *VaultStatus) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultStatus.
func (in *VaultStatus) DeepCopy() *VaultStatus {
	if in == nil {
		return nil
	}
	out := new(VaultStatus)
	in.DeepCopyInto(out)
	return out
}
