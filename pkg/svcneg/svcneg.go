/*
Copyright 2020 The Kubernetes Authors.

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
package svcneg

import (
	apisneg "k8s.io/ingress-gce/pkg/apis/svcneg"
	negv1beta1 "k8s.io/ingress-gce/pkg/apis/svcneg/v1beta1"
	"k8s.io/ingress-gce/pkg/crd"
)

func CRDMeta() *crd.CRDMeta {
	meta := crd.NewCRDMeta(
		apisneg.GroupName,
		"ServiceNetworkEndpointGroup",
		"ServiceNetworkEndpointGroupList",
		"servicenetworkendpointgroup",
		"servicenetworkendpointgroups",
		[]*crd.Version{
			crd.NewVersion("v1beta1", "k8s.io/ingress-gce/pkg/apis/svcneg/v1beta1.ServiceNetworkEndpointGroup", negv1beta1.GetOpenAPIDefinitions, false),
		},
		"svcneg",
	)
	return meta
}
