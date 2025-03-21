//go:build !ignore_autogenerated

/*
Copyright 2025.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Apperance) DeepCopyInto(out *Apperance) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Apperance.
func (in *Apperance) DeepCopy() *Apperance {
	if in == nil {
		return nil
	}
	out := new(Apperance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualPet) DeepCopyInto(out *VirtualPet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualPet.
func (in *VirtualPet) DeepCopy() *VirtualPet {
	if in == nil {
		return nil
	}
	out := new(VirtualPet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualPet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualPetList) DeepCopyInto(out *VirtualPetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualPet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualPetList.
func (in *VirtualPetList) DeepCopy() *VirtualPetList {
	if in == nil {
		return nil
	}
	out := new(VirtualPetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualPetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualPetSpec) DeepCopyInto(out *VirtualPetSpec) {
	*out = *in
	if in.Appearance != nil {
		in, out := &in.Appearance, &out.Appearance
		*out = new(Apperance)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualPetSpec.
func (in *VirtualPetSpec) DeepCopy() *VirtualPetSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualPetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualPetStatus) DeepCopyInto(out *VirtualPetStatus) {
	*out = *in
	in.LastInteraction.DeepCopyInto(&out.LastInteraction)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualPetStatus.
func (in *VirtualPetStatus) DeepCopy() *VirtualPetStatus {
	if in == nil {
		return nil
	}
	out := new(VirtualPetStatus)
	in.DeepCopyInto(out)
	return out
}
