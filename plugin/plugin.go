package plugin

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	framework "k8s.io/kubernetes/pkg/scheduler/framework/v1alpha1"
)

const (
	Name             = "Rabbit"
)

type RabbitPlugin struct {
	handle    framework.FrameworkHandle
	clientset kubernetes.Interface
}

var _ framework.PreFilterPlugin = &RabbitPlugin{}
var _ framework.FilterPlugin = &RabbitPlugin{}
func (p *RabbitPlugin) Name() string {
	return Name
}

func (p *RabbitPlugin) PreFilter(pc *framework.PluginContext, pod *corev1.Pod) *framework.Status {
	var ownerReference *metav1.OwnerReference

	ownerReference, _ = GetOwnerReference("ReplicaSet",pod.GetOwnerReferences())

	if ownerReference == nil {
		return framework.NewStatus(framework.Error, "could not find owning deployment")
	}

	rs, err := p.clientset.AppsV1().ReplicaSets(pod.Namespace).Get(ownerReference.Name, metav1.GetOptions{})
	if err != nil {
		return framework.NewStatus(framework.Error, fmt.Errorf("could not find owning deployment: %w", err).Error())
	}

	ownerReference, _ = GetOwnerReference("Deployment",rs.GetOwnerReferences())

	deployment, err := p.clientset.AppsV1().Deployments(pod.Namespace).Get(ownerReference.Name, metav1.GetOptions{})
	if err != nil {
		return framework.NewStatus(framework.Error, fmt.Errorf("could not find owning deployment: %w", err).Error())
	}

	if _, ok := deployment.Labels["rabbitCompatible"];  !ok {
		return framework.NewStatus(framework.Unschedulable, "")
	}
	// Do other compute and store in pc.Write...

	return framework.NewStatus(framework.Success, "Successfully scheduled")
}

func (p *RabbitPlugin) Filter(pc *framework.PluginContext, pod *corev1.Pod, nodeName string) *framework.Status {


	return framework.NewStatus(framework.Success, "")
}

func New(configuration *runtime.Unknown, f framework.FrameworkHandle) (framework.Plugin, error) {

	cfg, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("could not build in cluster rest config: %w", err)
	}

	clientSet, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return nil, fmt.Errorf("could not build clientset with in cluster rest config: %w", err)
	}

	return &RabbitPlugin{clientset: clientSet, handle:f }, nil
}
