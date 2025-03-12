/*
Copyright 2025.

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
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	kubepetv1alpha1 "github.com/GrapefruitCat030/KubePet/api/v1alpha1"
)

// VirtualPetReconciler reconciles a VirtualPet object
type VirtualPetReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=kubepet.grapefruitcat030.github.io,resources=virtualpets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=kubepet.grapefruitcat030.github.io,resources=virtualpets/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=kubepet.grapefruitcat030.github.io,resources=virtualpets/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the VirtualPet object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.20.2/pkg/reconcile
func (r *VirtualPetReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here
	virtualPet := &kubepetv1alpha1.VirtualPet{}
	if err := r.Get(ctx, req.NamespacedName, virtualPet); err != nil {
		log.Log.Error(err, "Failed to get VirtualPet")
		return ctrl.Result{}, err
	}
	log.Log.Info("Reconciling VirtualPet", "Name", virtualPet.Name)

	// Update the status of the VirtualPet
	virtualPet.Status.Hunger = 50
	virtualPet.Status.Mood = 50
	virtualPet.Status.Stage = "baby"
	virtualPet.Status.LastInteraction = metav1.Now()
	virtualPet.Status.StatusMessage = "VirtualPet is healthy"

	if err := r.Status().Update(ctx, virtualPet); err != nil {
		log.Log.Error(err, "unable to update VirtualPet status")
		return ctrl.Result{}, err
	}

	// Requeue the request after a certain duration to simulate the VirtualPet's status change
	return ctrl.Result{RequeueAfter: time.Minute}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *VirtualPetReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&kubepetv1alpha1.VirtualPet{}).
		Named("virtualpet").
		Complete(r)
}
