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

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	divisionv1 "k8-controller/api/v1"
)

// DivideResourceReconciler reconciles a DivideResource object
type DivideResourceReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=division.division.io,resources=divideresources,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=division.division.io,resources=divideresources/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=division.division.io,resources=divideresources/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the DivideResource object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.17.3/pkg/reconcile
func (r *DivideResourceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here
	var divide divisionv1.DivideResource
	err := r.Get(ctx, req.NamespacedName, &divide)
	if err != nil {
		klog.Error(err, "unable to fetch sum")
		return ctrl.Result{}, client.IgnoreNotFound(err)

	}
	result := divide.Spec.NumberOne / divide.Spec.NumberTwo
	divide.Status.Result = result
	klog.Info("Updating the result ")
	if err := r.Status().Update(ctx, &divide); err != nil {
		klog.Error(err, "Unable to update status")
		return ctrl.Result{}, err
	}

	klog.Infof("Successfully updated the status with result %d", result)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *DivideResourceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&divisionv1.DivideResource{}).
		Complete(r)
}
