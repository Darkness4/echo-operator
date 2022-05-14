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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MessageSpec defines the desired state of Message
type MessageSpec struct {
	// Value is an example field of Message.
	Value string `json:"value,omitempty"`
}

// MessageStatus defines the observed state of Message
type MessageStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Message is the Schema for the messages API
type Message struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MessageSpec   `json:"spec,omitempty"`
	Status MessageStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MessageList contains a list of Message
type MessageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Message `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Message{}, &MessageList{})
}
