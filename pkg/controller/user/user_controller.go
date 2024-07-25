package user

import (
	"context"
	"fmt"
	"reflect"

	ctlrbacv1 "github.com/rancher/wrangler/v3/pkg/generated/controllers/rbac/v1"
	"github.com/sirupsen/logrus"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	mgmtv1 "github.com/llmos-ai/llmos-operator/pkg/apis/management.llmos.ai/v1"
	"github.com/llmos-ai/llmos-operator/pkg/constant"
	ctlmgmtv1 "github.com/llmos-ai/llmos-operator/pkg/generated/controllers/management.llmos.ai/v1"
	"github.com/llmos-ai/llmos-operator/pkg/indexeres"

	"github.com/llmos-ai/llmos-operator/pkg/server/config"
)

const (
	usernameLabelKey     = "management.llmos.ai/username"
	publicInfoViewerRole = "system:public-info-viewer"
	userControllerName   = "user.onChange"
)

// userHandler reconcile the user's clusterRole and clusterRoleBinding
type userHandler struct {
	users                   ctlmgmtv1.UserClient
	clusterRoleBindings     ctlrbacv1.ClusterRoleBindingClient
	clusterRoleBindingCache ctlrbacv1.ClusterRoleBindingCache
}

func Register(ctx context.Context, management *config.Management) error {
	users := management.MgmtFactory.Management().V1().User()

	userRBACController := &userHandler{
		users:                   users,
		clusterRoleBindings:     management.RbacFactory.Rbac().V1().ClusterRoleBinding(),
		clusterRoleBindingCache: management.RbacFactory.Rbac().V1().ClusterRoleBinding().Cache(),
	}

	users.OnChange(ctx, userControllerName, userRBACController.OnChanged)
	return nil
}

func (h *userHandler) OnChanged(_ string, user *mgmtv1.User) (*mgmtv1.User, error) {
	if user == nil || user.DeletionTimestamp != nil {
		return user, nil
	}

	roleName := publicInfoViewerRole
	toUpdate := user.DeepCopy()
	if toUpdate.Spec.IsAdmin {
		roleName = constant.AdminRole
	}

	if err := h.ensureClusterBinding(roleName, toUpdate); err != nil {
		return user, err
	}

	return h.updateStatus(user, toUpdate)
}

func (h *userHandler) updateStatus(user *mgmtv1.User, toUpdate *mgmtv1.User) (*mgmtv1.User, error) {
	toUpdate.Status.IsAdmin = toUpdate.Spec.IsAdmin
	toUpdate.Status.IsActive = toUpdate.Spec.IsActive
	if !reflect.DeepEqual(user.Status, toUpdate.Status) {
		toUpdate.Status.LastUpdateTime = metav1.Now().Format(constant.TimeLayout)
		return h.users.UpdateStatus(toUpdate)
	}
	return nil, nil
}

func (h *userHandler) ensureClusterBinding(roleName string, user *mgmtv1.User) error {
	subject := rbacv1.Subject{
		Kind: "User",
		Name: user.Name,
	}

	// find if there is a clusterRoleBinding with the same role and subject
	key := indexeres.GetCrbKey(roleName, subject)
	crbs, err := h.clusterRoleBindingCache.GetByIndex(indexeres.ClusterRoleBindingNameIndex, key)
	if err != nil {
		return err
	}
	if len(crbs) > 0 {
		logrus.Infof("ClusterRoleBinding with role %v for subject %v already exists, skip creating.", roleName, subject.Name)
		return nil
	}

	logrus.Infof("Creating clusterRoleBinding with role %v for subject %v", roleName, subject.Name)
	_, err = h.clusterRoleBindings.Create(&rbacv1.ClusterRoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: fmt.Sprintf("%s-", user.Name),
			Labels: map[string]string{
				usernameLabelKey: user.Name,
			},
			OwnerReferences: []metav1.OwnerReference{
				{
					APIVersion: mgmtv1.SchemeGroupVersion.String(),
					Kind:       "User",
					Name:       user.Name,
					UID:        user.UID,
				},
			},
		},
		Subjects: []rbacv1.Subject{subject},
		RoleRef: rbacv1.RoleRef{
			Kind: "ClusterRole",
			Name: roleName,
		},
	})

	return err
}
