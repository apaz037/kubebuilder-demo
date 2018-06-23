// Api versions allow the api contract for a resource to be changed while keeping
// backward compatibility by support multiple concurrent versions
// of the same resource

// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=package,register
// +k8s:conversion-gen=kubebuilder-demo/pkg/apis/mammals
// +k8s:defaulter-gen=TypeMeta
// +groupName=mammals.chill.af
package v1beta1 // import "kubebuilder-demo/pkg/apis/mammals/v1beta1"
