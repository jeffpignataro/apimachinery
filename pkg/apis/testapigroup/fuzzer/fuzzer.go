/*
Copyright 2015 The Kubernetes Authors.

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

package fuzzer

import (
	"fmt"

	fuzz "github.com/google/gofuzz"

	apitesting "github.com/jeffpignataro/apimachinery/pkg/api/apitesting"
	"github.com/jeffpignataro/apimachinery/pkg/api/apitesting/fuzzer"
	"github.com/jeffpignataro/apimachinery/pkg/apis/testapigroup"
	v1 "github.com/jeffpignataro/apimachinery/pkg/apis/testapigroup/v1"
	"github.com/jeffpignataro/apimachinery/pkg/runtime"
	runtimeserializer "github.com/jeffpignataro/apimachinery/pkg/runtime/serializer"
)

// overrideMetaFuncs override some generic fuzzer funcs from github.com/jeffpignataro/apimachinery in order to have more realistic
// values in a Kubernetes context.
func overrideMetaFuncs(codecs runtimeserializer.CodecFactory) []interface{} {
	return []interface{}{
		func(j *runtime.Object, c fuzz.Continue) {
			// TODO: uncomment when round trip starts from a versioned object
			if true { //c.RandBool() {
				*j = &runtime.Unknown{
					// We do not set TypeMeta here because it is not carried through a round trip
					Raw:         []byte(`{"apiVersion":"unknown.group/unknown","kind":"Something","someKey":"someValue"}`),
					ContentType: runtime.ContentTypeJSON,
				}
			} else {
				types := []runtime.Object{&testapigroup.Carp{}}
				t := types[c.Rand.Intn(len(types))]
				c.Fuzz(t)
				*j = t
			}
		},
		func(r *runtime.RawExtension, c fuzz.Continue) {
			// Pick an arbitrary type and fuzz it
			types := []runtime.Object{&testapigroup.Carp{}}
			obj := types[c.Rand.Intn(len(types))]
			c.Fuzz(obj)

			// Convert the object to raw bytes
			bytes, err := runtime.Encode(apitesting.TestCodec(codecs, v1.SchemeGroupVersion), obj)
			if err != nil {
				panic(fmt.Sprintf("Failed to encode object: %v", err))
			}

			// Set the bytes field on the RawExtension
			r.Raw = bytes
		},
	}
}

func testapigroupFuncs(codecs runtimeserializer.CodecFactory) []interface{} {
	return []interface{}{
		func(s *testapigroup.CarpSpec, c fuzz.Continue) {
			c.FuzzNoCustom(s)
			// has a default value
			ttl := int64(30)
			if c.RandBool() {
				ttl = int64(c.Uint32())
			}
			s.TerminationGracePeriodSeconds = &ttl

			if s.SchedulerName == "" {
				s.SchedulerName = "default-scheduler"
			}
		},
		func(j *testapigroup.CarpPhase, c fuzz.Continue) {
			statuses := []testapigroup.CarpPhase{"Pending", "Running", "Succeeded", "Failed", "Unknown"}
			*j = statuses[c.Rand.Intn(len(statuses))]
		},
		func(rp *testapigroup.RestartPolicy, c fuzz.Continue) {
			policies := []testapigroup.RestartPolicy{"Always", "Never", "OnFailure"}
			*rp = policies[c.Rand.Intn(len(policies))]
		},
	}
}

// Funcs returns the fuzzer functions for the testapigroup.
var Funcs = fuzzer.MergeFuzzerFuncs(
	overrideMetaFuncs,
	testapigroupFuncs,
)
