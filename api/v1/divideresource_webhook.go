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

package v1

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// log is for logging in this package.
var divideresourcelog = logf.Log.WithName("divideresource-resource")

// SetupWebhookWithManager will setup the manager to manage the webhooks
func (r *DivideResource) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// NOTE: The 'path' attribute must follow a specific pattern and should not be modified directly here.
// Modifying the path for an invalid path can cause API server errors; failing to locate the webhook.
//+kubebuilder:webhook:path=/validate-division-division-io-v1-divideresource,mutating=false,failurePolicy=fail,sideEffects=None,groups=division.division.io,resources=divideresources,verbs=create;update,versions=v1,name=vdivideresource.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &DivideResource{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *DivideResource) ValidateCreate() (admission.Warnings, error) {
	divideresourcelog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	klog.Infof("In validate create of %s", r.Name)
	if r.Spec.NumberOne < 0 || r.Spec.NumberTwo < 0 {
		return nil, fmt.Errorf("input values Number One: %d Number Two: %d cannot be negative ", r.Spec.NumberOne, r.Spec.NumberTwo)
	}
	if r.Spec.NumberTwo == 0 {
		return nil, fmt.Errorf("input value Number Two: %d cannot be zero", r.Spec.NumberTwo)
	}

	return nil, nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *DivideResource) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	divideresourcelog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil, nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *DivideResource) ValidateDelete() (admission.Warnings, error) {
	divideresourcelog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil, nil
}
