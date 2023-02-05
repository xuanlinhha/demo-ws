package service

import (
	"context"
	"os"

	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	apiextv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes"
	scheme "k8s.io/client-go/kubernetes/scheme"
)

type ResourceService interface {
	CreatePodWithConfigmapFromFiles(podFile, configMapFile string) error
}

type resourceService struct {
	clientset kubernetes.Interface
}

func NewResourceService(clientset kubernetes.Interface) ResourceService {
	return &resourceService{clientset: clientset}
}

func (r *resourceService) CreatePodWithConfigmapFromFiles(configMapFile, podFile string) error {
	sch := runtime.NewScheme()
	_ = scheme.AddToScheme(sch)
	_ = apiextv1beta1.AddToScheme(sch)
	decode := serializer.NewCodecFactory(sch).UniversalDeserializer().Decode
	// add config map
	stream, _ := os.ReadFile(configMapFile)
	obj, gKV, _ := decode(stream, nil, nil)
	if gKV.Kind == "ConfigMap" {
		cm := obj.(*corev1.ConfigMap)
		if _, err := r.clientset.CoreV1().ConfigMaps("default").Create(context.TODO(), cm, metav1.CreateOptions{}); err != nil {
			return errors.Wrap(err, "Create")
		}
	}

	// add pod
	stream, _ = os.ReadFile(podFile)
	obj, gKV, _ = decode(stream, nil, nil)
	if gKV.Kind == "Pod" {
		pod := obj.(*corev1.Pod)
		if _, err := r.clientset.CoreV1().Pods("default").Create(context.TODO(), pod, metav1.CreateOptions{}); err != nil {
			return errors.Wrap(err, "Create")
		}
	}
	return nil
}
