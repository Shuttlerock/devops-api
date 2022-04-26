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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ReleaseSpec defines the desired state of Release
type ReleaseSpec struct {
	ReleaseNotesURL string         `json:"releaseNotesURL,omitempty" protobuf:"bytes,8,opt,name=releaseNotesURL"`
	Issues          []IssueSummary `json:"issues,omitempty" protobuf:"bytes,6,opt,name=issues"`
}

// IssueSummary is the summary of an issue
type IssueSummary struct {
	ID                string        `json:"id,omitempty"  protobuf:"bytes,1,opt,name=id"`
	URL               string        `json:"url,omitempty"  protobuf:"bytes,2,opt,name=url"`
	Title             string        `json:"title,omitempty"  protobuf:"bytes,3,opt,name=title"`
	Body              string        `json:"body,omitempty"  protobuf:"bytes,4,opt,name=body"`
	State             string        `json:"state,omitempty"  protobuf:"bytes,5,opt,name=state"`
	Message           string        `json:"message,omitempty"  protobuf:"bytes,6,opt,name=message"`
	User              *UserDetails  `json:"user,omitempty"  protobuf:"bytes,7,opt,name=user"`
	Assignees         []UserDetails `json:"assignees,omitempty"  protobuf:"bytes,8,opt,name=assignees"`
	ClosedBy          *UserDetails  `json:"closedBy,omitempty"  protobuf:"bytes,9,opt,name=closedBy"`
	CreationTimestamp *metav1.Time  `json:"creationTimestamp,omitempty" protobuf:"bytes,10,opt,name=creationTimestamp"`
	Labels            []IssueLabel  `json:"labels,omitempty" protobuf:"bytes,11,opt,name=labels"`
}

type IssueLabel struct {
	//URL   string `json:"name,omitempty"  protobuf:"bytes,1,opt,name=name"`
	//Name  string `json:"url,omitempty"  protobuf:"bytes,2,opt,name=url"`
	//Color string `json:"color,omitempty"  protobuf:"bytes,3,opt,name=color"`
}

// UserDetails contains details of a user
type UserDetails struct {
	//Login             string             `json:"login,omitempty"  protobuf:"bytes,1,opt,name=login"`
	//Name              string             `json:"name,omitempty"  protobuf:"bytes,2,opt,name=name"`
	//Email             string             `json:"email,omitempty"  protobuf:"bytes,3,opt,name=email"`
	//CreationTimestamp *metav1.Time       `json:"creationTimestamp,omitempty" protobuf:"bytes,4,opt,name=creationTimestamp"`
	//URL               string             `json:"url,omitempty"  protobuf:"bytes,5,opt,name=url"`
	//AvatarURL         string             `json:"avatarUrl,omitempty"  protobuf:"bytes,6,opt,name=avatarUrl"`
	//ServiceAccount    string             `json:"serviceAccount,omitempty"  protobuf:"bytes,7,opt,name=serviceAccount"`
	//Accounts          []AccountReference `json:"accountReference,omitempty"  protobuf:"bytes,8,opt,name=accountReference"`
	//ExternalUser      bool               `json:"externalUser,omitempty"  protobuf:"bytes,9,opt,name=externalUser"`
}

// AccountReference is a reference to a user account in another system that is attached to this user
type AccountReference struct {
	//Provider string `json:"provider,omitempty"  protobuf:"bytes,1,opt,name=provider"`
	//ID       string `json:"id,omitempty"  protobuf:"bytes,2,opt,name=id"`
}

// ReleaseStatus defines the observed state of Release
type ReleaseStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Release is the Schema for the releases API
type Release struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ReleaseSpec   `json:"spec,omitempty"`
	Status ReleaseStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ReleaseList contains a list of Release
type ReleaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Release `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Release{}, &ReleaseList{})
}
