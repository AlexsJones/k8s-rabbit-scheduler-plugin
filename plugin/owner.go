package plugin

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"

)

func GetOwnerReference(name string,ref []metav1.OwnerReference) (*metav1.OwnerReference, error){
	var ownerReference *metav1.OwnerReference
	for _, or := range ref {
		klog.Infof("Kind: %s", or.Kind)
		if or.Kind == name {
			ownerReference = &or
			break
		}
	}

	return ownerReference,nil
}


