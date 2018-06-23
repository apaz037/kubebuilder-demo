package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!
// Created by "kubebuilder create resource" for you to implement the Sloth resource schema definition
// as a go struct.
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SlothSpec defines the desired state of Sloth
type SlothSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "kubebuilder generate" to regenerate code after modifying this file
	Color string `json:"color,omitempty"`

	//  Define relevant sloth config bits here.
}

// SlothStatus defines the observed state of Sloth
type SlothStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "kubebuilder generate" to regenerate code after modifying this file
	Message string `json:"message,omitempty"` 
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Sloth
// +k8s:openapi-gen=true
// +kubebuilder:resource:path=sloths
type Sloth struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SlothSpec   `json:"spec,omitempty"`
	Status SlothStatus `json:"status,omitempty"`
}
