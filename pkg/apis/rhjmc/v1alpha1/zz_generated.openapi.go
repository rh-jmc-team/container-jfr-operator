// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/rhjmc/v1alpha1.ContainerJFR":         schema_pkg_apis_rhjmc_v1alpha1_ContainerJFR(ref),
		"./pkg/apis/rhjmc/v1alpha1.ContainerJFRSpec":     schema_pkg_apis_rhjmc_v1alpha1_ContainerJFRSpec(ref),
		"./pkg/apis/rhjmc/v1alpha1.ContainerJFRStatus":   schema_pkg_apis_rhjmc_v1alpha1_ContainerJFRStatus(ref),
		"./pkg/apis/rhjmc/v1alpha1.FlightRecorder":       schema_pkg_apis_rhjmc_v1alpha1_FlightRecorder(ref),
		"./pkg/apis/rhjmc/v1alpha1.FlightRecorderSpec":   schema_pkg_apis_rhjmc_v1alpha1_FlightRecorderSpec(ref),
		"./pkg/apis/rhjmc/v1alpha1.FlightRecorderStatus": schema_pkg_apis_rhjmc_v1alpha1_FlightRecorderStatus(ref),
	}
}

func schema_pkg_apis_rhjmc_v1alpha1_ContainerJFR(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ContainerJFR is the Schema for the containerjfrs API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/rhjmc/v1alpha1.ContainerJFRSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/rhjmc/v1alpha1.ContainerJFRStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/rhjmc/v1alpha1.ContainerJFRSpec", "./pkg/apis/rhjmc/v1alpha1.ContainerJFRStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_rhjmc_v1alpha1_ContainerJFRSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ContainerJFRSpec defines the desired state of ContainerJFR",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"minimal": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
				},
				Required: []string{"minimal"},
			},
		},
	}
}

func schema_pkg_apis_rhjmc_v1alpha1_ContainerJFRStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ContainerJFRStatus defines the observed state of ContainerJFR",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_rhjmc_v1alpha1_FlightRecorder(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "FlightRecorder is the Schema for the flightrecorders API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/rhjmc/v1alpha1.FlightRecorderSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/rhjmc/v1alpha1.FlightRecorderStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/rhjmc/v1alpha1.FlightRecorderSpec", "./pkg/apis/rhjmc/v1alpha1.FlightRecorderStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_rhjmc_v1alpha1_FlightRecorderSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "FlightRecorderSpec defines the desired state of FlightRecorder",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"recordingActive": {
						SchemaProps: spec.SchemaProps{
							Description: "Specifies whether a JFR recording should be started or stopped",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
				},
				Required: []string{"recordingActive"},
			},
		},
	}
}

func schema_pkg_apis_rhjmc_v1alpha1_FlightRecorderStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "FlightRecorderStatus defines the observed state of FlightRecorder",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"target": {
						SchemaProps: spec.SchemaProps{
							Description: "Reference to the pod/service that this object controls JFR for",
							Ref:         ref("k8s.io/api/core/v1.ObjectReference"),
						},
					},
					"recordingActive": {
						SchemaProps: spec.SchemaProps{
							Description: "Whether the pod/service is currently recording",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"recordings": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Lists all recordings for the pod/service that may be downloaded",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
				Required: []string{"target", "recordingActive", "recordings"},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.ObjectReference"},
	}
}
