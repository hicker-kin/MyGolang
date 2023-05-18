package deployment

import (
	appsV1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	appV1 "app.deploy/api/v1"
	"app.deploy/constant"
	"app.deploy/resource"
)

func New(app *appV1.App) *appsV1.Deployment {
	labels := map[string]string{resource.GetAppLabelKey(): app.Name}
	selector := &metaV1.LabelSelector{MatchLabels: labels}
	return &appsV1.Deployment{
		TypeMeta: metaV1.TypeMeta{
			APIVersion: constant.DeploymentVersion,
			Kind:       constant.DeploymentKind,
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name:      app.Name,
			Namespace: app.Namespace,
			OwnerReferences: []metaV1.OwnerReference{
				*metaV1.NewControllerRef(app, schema.GroupVersionKind{
					Group:   appV1.GroupVersion.Group,
					Version: appV1.GroupVersion.Version,
					Kind:    constant.APPName,
				}),
			},
		},
		Spec: appsV1.DeploymentSpec{
			Replicas: &app.Spec.Replicas,
			Selector: selector,
			Template: coreV1.PodTemplateSpec{
				ObjectMeta: metaV1.ObjectMeta{
					Labels: labels,
				},
				Spec: coreV1.PodSpec{
					Containers: newContainers(app),
				},
			},
		},
	}
}

func newContainers(app *appV1.App) []coreV1.Container {
	var containerPorts []coreV1.ContainerPort
	for _, servicePort := range app.Spec.Ports {
		var port coreV1.ContainerPort
		port.ContainerPort = servicePort.TargetPort.IntVal
		containerPorts = append(containerPorts, port)
	}
	return []coreV1.Container{
		{
			Name:            app.Name,
			Image:           app.Spec.Image,
			Ports:           containerPorts,
			Env:             app.Spec.Envs,
			Resources:       app.Spec.Resources,
			ImagePullPolicy: coreV1.PullIfNotPresent,
		},
	}
}
