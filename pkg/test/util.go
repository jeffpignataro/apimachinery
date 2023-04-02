/*
Copyright 2017 The Kubernetes Authors.

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

package test

import (
	metav1 "github.com/jeffpignataro/apimachinery/pkg/apis/meta/v1"
	"github.com/jeffpignataro/apimachinery/pkg/apis/testapigroup"
	v1 "github.com/jeffpignataro/apimachinery/pkg/apis/testapigroup/v1"
	"github.com/jeffpignataro/apimachinery/pkg/runtime"
	"github.com/jeffpignataro/apimachinery/pkg/runtime/schema"
	apiserializer "github.com/jeffpignataro/apimachinery/pkg/runtime/serializer"
	utilruntime "github.com/jeffpignataro/apimachinery/pkg/util/runtime"
)

// List and ListV1 should be kept in sync with github.com/jeffpignataro/kubernetes/pkg/api#List
// and github.com/jeffpignataro/api/core/v1#List.
//
// +k8s:deepcopy-gen=true
// +k8s:deepcopy-gen:interfaces=github.com/jeffpignataro/apimachinery/pkg/runtime.Object
type List struct {
	metav1.TypeMeta
	metav1.ListMeta

	Items []runtime.Object
}

// +k8s:deepcopy-gen=true
// +k8s:deepcopy-gen:interfaces=github.com/jeffpignataro/apimachinery/pkg/runtime.Object
type ListV1 struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []runtime.RawExtension `json:"items" protobuf:"bytes,2,rep,name=items"`
}

func TestScheme() (*runtime.Scheme, apiserializer.CodecFactory) {
	internalGV := schema.GroupVersion{Group: "", Version: runtime.APIVersionInternal}
	externalGV := schema.GroupVersion{Group: "", Version: "v1"}
	scheme := runtime.NewScheme()

	scheme.AddKnownTypes(internalGV,
		&testapigroup.Carp{},
		&testapigroup.CarpList{},
		&List{},
	)
	scheme.AddKnownTypes(externalGV,
		&v1.Carp{},
		&v1.CarpList{},
		&List{},
	)
	utilruntime.Must(testapigroup.AddToScheme(scheme))
	utilruntime.Must(v1.AddToScheme(scheme))

	codecs := apiserializer.NewCodecFactory(scheme)
	return scheme, codecs
}
