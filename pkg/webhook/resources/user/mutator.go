package user

import (
	"github.com/oneblock-ai/webhook/pkg/server/admission"
	"github.com/sirupsen/logrus"
	admissionregv1 "k8s.io/api/admissionregistration/v1"
	"k8s.io/apimachinery/pkg/runtime"

	mgmtv1 "github.com/llmos-ai/llmos-operator/pkg/apis/management.llmos.ai/v1"
	"github.com/llmos-ai/llmos-operator/pkg/auth"
	"github.com/llmos-ai/llmos-operator/pkg/constant"
)

type mutator struct {
	admission.DefaultMutator
}

var _ admission.Mutator = &mutator{}

func NewMutator() admission.Mutator {
	return &mutator{}
}

func (m *mutator) Create(_ *admission.Request, newObj runtime.Object) (admission.Patch, error) {
	user := newObj.(*mgmtv1.User)
	logrus.Infof("[webhook mutating]user %s is created", user.Name)

	patchOps := make([]admission.PatchOp, 0)

	patchOps = append(patchOps, patchLabels(user.Labels))

	// skip default admin password hash
	if user.Labels != nil && user.Labels[constant.DefaultAdminLabelKey] != "true" {
		// hash password
		if passPatch, err := patchPassword(user.Spec.Password); err != nil {
			return nil, err
		} else {
			patchOps = append(patchOps, passPatch)
		}
	}

	return patchOps, nil
}

func patchLabels(labels map[string]string) admission.PatchOp {
	if labels == nil {
		labels = map[string]string{}
	}
	labels["llmos.ai/creator"] = "llmos-operator"
	return admission.PatchOp{
		Op:    admission.PatchOpReplace,
		Path:  "/metadata/labels",
		Value: labels,
	}
}

func patchPassword(password string) (admission.PatchOp, error) {
	hash, err := auth.HashPassword(password)
	if err != nil {
		return admission.PatchOp{}, err
	}
	return admission.PatchOp{
		Op:    admission.PatchOpReplace,
		Path:  "/spec/password",
		Value: hash,
	}, nil
}

func (m *mutator) Update(_ *admission.Request, oldObj, newObj runtime.Object) (admission.Patch, error) {
	oldUSer := oldObj.(*mgmtv1.User)
	user := newObj.(*mgmtv1.User)
	logrus.Debugf("user %s is updated", user.Name)

	patchOps := make([]admission.PatchOp, 0)

	if (oldUSer.Spec.Password != user.Spec.Password) && user.Spec.Password != "" {
		logrus.Debugf("updating password to: %s from %s", user.Spec.Password, oldUSer.Spec.Password)
		if passPatch, err := patchPassword(user.Spec.Password); err != nil {
			return nil, err
		} else {
			patchOps = append(patchOps, passPatch)
		}
	}

	return patchOps, nil
}

func (m *mutator) Resource() admission.Resource {
	return admission.Resource{
		Names:      []string{"users"},
		Scope:      admissionregv1.ClusterScope,
		APIGroup:   mgmtv1.SchemeGroupVersion.Group,
		APIVersion: mgmtv1.SchemeGroupVersion.Version,
		ObjectType: &mgmtv1.User{},
		OperationTypes: []admissionregv1.OperationType{
			admissionregv1.Create,
			admissionregv1.Update,
		},
	}
}
