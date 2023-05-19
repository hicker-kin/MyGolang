package service

import (
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	appV1 "app.deploy/api/v1"
	"app.deploy/constant"
	"app.deploy/resource"
)

func New(app *appV1.App) *coreV1.Service {
	return &coreV1.Service{
		TypeMeta: metaV1.TypeMeta{
			Kind:       constant.ServiceKind,
			APIVersion: constant.ServiceVersion,
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name:      app.Name,
			Namespace: app.Namespace,
			OwnerReferences: []metaV1.OwnerReference{
				*metaV1.NewControllerRef(app, schema.GroupVersionKind{
					Group:   appV1.GroupVersion.Group,
					Version: appV1.GroupVersion.Version,
					Kind:    constant.ControllerKind,
				}),
			},
		},
		Spec: coreV1.ServiceSpec{
			Ports: app.Spec.Ports,
			Selector: map[string]string{
				resource.GetAppLabelKey(): app.Name,
			},
		},
	}
}
