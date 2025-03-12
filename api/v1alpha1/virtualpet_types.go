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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// Apperance defines the visual characteristics of the pet.
type Apperance struct {
	// Color of the pet, e.g. black, white, etc.
	Color string `json:"color,omitempty"`
	// Variant of the pet, e.g. siamese, persian, etc.
	Variant string `json:"variant,omitempty"`
}

// VirtualPetSpec defines the desired state of VirtualPet.
type VirtualPetSpec struct {
	// Type of pet, e.g. cat, dog, etc.
	// +kubebuilder:validation:Enum=cat;dog;bird;dragon
	PetType string `json:"petType"`
	// Name of the pet
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=20
	PetName string `json:"petName"`
	// Appearance details of the pet
	Appearance *Apperance `json:"appearance,omitempty"`
}

// VirtualPetStatus defines the observed state of VirtualPet.
type VirtualPetStatus struct {
	// Hunger level of the pet (0-100)
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	Hunger int32 `json:"hunger,omitempty"`
	// Mood level of the pet (0-100)
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	Mood int32 `json:"mood,omitempty"`
	// Current life stage of the pet, e.g. baby, adult, etc.
	// +kubebuilder:validation:Enum=baby;young;adult;senior
	Stage string `json:"stage,omitempty"`
	// Timestamp of the last interaction with the pet
	LastInteraction metav1.Time `json:"lastInteraction,omitempty"`
	// Status message of the pet
	StatusMessage string `json:"statusMessage,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="PetType",type="string",JSONPath=".spec.petType"
// +kubebuilder:printcolumn:name="PetName",type="string",JSONPath=".spec.petName"
// +kubebuilder:printcolumn:name="Hunger",type="integer",JSONPath=".status.hunger"
// +kubebuilder:printcolumn:name="Mood",type="integer",JSONPath=".status.mood"
// +kubebuilder:printcolumn:name="Stage",type="string",JSONPath=".status.stage"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// VirtualPet is the Schema for the virtualpets API.
type VirtualPet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualPetSpec   `json:"spec,omitempty"`
	Status VirtualPetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualPetList contains a list of VirtualPet.
type VirtualPetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualPet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VirtualPet{}, &VirtualPetList{})
}
