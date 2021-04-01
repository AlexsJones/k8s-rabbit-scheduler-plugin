package plugin

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	framework "k8s.io/kubernetes/pkg/scheduler/framework/v1alpha1"
)

const (
	Name             = "RabbitPlugin"
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


	return framework.NewStatus(framework.Success, "")
}

func (p *RabbitPlugin) Filter(pc *framework.PluginContext, pod *corev1.Pod, nodeName string) *framework.Status {

	// Ignore pods that are not labelled as rabbitCompatible
	if _, ok := pod.Labels["rabbitCompatible"];  !ok {
		return framework.NewStatus(framework.Unschedulable, "")
	}


	return framework.NewStatus(framework.Success, "")
}

func New(configuration *runtime.Unknown, f framework.FrameworkHandle) (framework.Plugin, error) {
	// This is effectively a hack. Newer versions (1.17+) of the scheduler runtime provide access
	// to watchers via the FrameworkHandle. For now, inject our own k8s clientset set to gain access
	// to volume information
	// https://github.com/cockroachlabs/crl-scheduler/blob/master/plugin/plugin.go
	cfg, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("could not build in cluster rest config: %w", err)
	}

	clientset, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return nil, fmt.Errorf("could not build clientset with in cluster rest config: %w", err)
	}

	return &RabbitPlugin{handle: f, clientset: clientset}, nil
}
