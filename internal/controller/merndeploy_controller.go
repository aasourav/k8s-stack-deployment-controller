/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	controllerapi "sandtech.io/stackdeployment/api/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	stackdeploymentsv1 "sandtech.io/stackdeployment/api/v1"
)

// MernDeployReconciler reconciles a MernDeploy object
type MernDeployReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Log logr.Logger
	KubeClients
}

// +kubebuilder:rbac:groups=stackdeployments.sandtech.io,resources=merndeploys,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=stackdeployments.sandtech.io,resources=merndeploys/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=stackdeployments.sandtech.io,resources=merndeploys/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the MernDeploy object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.18.4/pkg/reconcile
func (r *MernDeployReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := r.Log.WithValues("reconciling: ",req.NamespacedName)

	mernDeploy := &controllerapi.MernDeploy{}
	err := r.Get(ctx, types.NamespacedName{Namespace: req.Namespace, Name: req.Name}, mernDeploy)

	if err !=nil{
		if errors.IsNotFound(err){
			l.Info(fmt.Sprintf("could not find: %s/%s",req.Name, req.Namespace))
			return ctrl.Result{},nil
		}
		l.Error(err, fmt.Sprintf("failed to get: %s/%s",req.Name, req.Namespace))
		return ctrl.Result{}, err
	}

	


	// TODO(user): your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *MernDeployReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&stackdeploymentsv1.MernDeploy{}).WithOptions(controller.Options{MaxConcurrentReconciles: 2}).
		Complete(r)
}
