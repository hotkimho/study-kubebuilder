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

	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	batchv1 "tutorial.kubebuilder.io/project/api/v1"
)

// TestPodReconciler reconciles a TestPod object
type TestPodReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=batch.tutorial.kubebuilder.io,resources=testpods,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=batch.tutorial.kubebuilder.io,resources=testpods/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=batch.tutorial.kubebuilder.io,resources=testpods/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the TestPod object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.18.4/pkg/reconcile
func (r *TestPodReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	//var testPod batchv1.TestPod
	//if err := r.Get(ctx, req.NamespacedName, &testPod); err != nil {
	//	log.Log.Error(err, "unable to fetch TestPod")
	//	return ctrl.Result{}, client.IgnoreNotFound(err)
	//}
	//
	//testPod.Status.Phase = "Pending"
	//if err := r.Status().Update(ctx, &testPod); err != nil {
	//	log.Log.Error(err, "unable to update TestPod status")
	//	return ctrl.Result{}, err
	//}

	var testPod batchv1.TestPod
	if err := r.Get(ctx, req.NamespacedName, &testPod); err != nil {
		log.Log.Error(err, "unable to fetch TestPod")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	switch testPod.Status.Phase {
	case "":
		testPod.Status.Phase = "Pending"
	case "Pending":
		testPod.Status.Phase = "Running"
	case "Running":
		testPod.Status.Phase = "Succeeded"
	default:
		// do error handling
	}
	if err := r.Status().Update(ctx, &testPod); err != nil {
		log.Log.Error(err, "unable to update TestPod status")
		return ctrl.Result{}, err
	}
	//
	//newTestPod := batchv1.TestPod{
	//	ObjectMeta: metav1.ObjectMeta{
	//		Namespace: "default",
	//		Name:      "testpod-example-1",
	//	},
	//	TypeMeta: metav1.TypeMeta{
	//		Kind:       "TestPod",
	//		APIVersion: "v1",
	//	},
	//	Spec: batchv1.TestPodSpec{
	//		Replicas: 3,
	//		Creator:  "kimho",
	//	},
	//	Status: batchv1.TestPodStatus{},
	//}
	//newTestPod.Status.Phase = "Running"
	//if err := r.Create(ctx, &newTestPod); err != nil {
	//	log.Log.Error(err, "unable to create TestPod")
	//	return ctrl.Result{}, err
	//}
	//
	//newTestPod2 := batchv1.TestPod{
	//	ObjectMeta: metav1.ObjectMeta{
	//		Namespace: "default",
	//		Name:      "testpod-example-1",
	//	},
	//	TypeMeta: metav1.TypeMeta{
	//		Kind:       "TestPod",
	//		APIVersion: "v1",
	//	},
	//	Spec: batchv1.TestPodSpec{
	//		Replicas: 3,
	//		Creator:  "kimho",
	//	},
	//	Status: batchv1.TestPodStatus{},
	//}
	//
	//newTestPod2.Status.Phase = ""
	//if err := r.Create(ctx, &newTestPod2); err != nil {
	//	log.Log.Error(err, "unable to create TestPod")
	//	return ctrl.Result{}, err
	//}
	//
	//time.Sleep(10 * time.Second)
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *TestPodReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&batchv1.TestPod{}).
		WithEventFilter(predicate.Funcs{
			CreateFunc: createPredicate,
			UpdateFunc: updatePredicate,
			DeleteFunc: deletePredicate,
		}).
		Complete(r)
}

func createPredicate(e event.CreateEvent) bool {
	createdTestPod := e.Object.(*batchv1.TestPod)

	// 이름을 검사하여 특정 이름의 리소스만 처리
	if createdTestPod.Name != "testpod-example-1" {
		return false
	}

	// 리소스 이벤트 핸들러에 전달
	return true
}
func updatePredicate(e event.UpdateEvent) bool {

	oldTestPod := e.ObjectNew.(*batchv1.TestPod)
	newTestPod := e.ObjectNew.(*batchv1.TestPod)

	fmt.Println("old predicate ", oldTestPod.Spec.Creator)
	fmt.Println("new predicate ", newTestPod.Spec.Creator)
	if oldTestPod.Spec.Creator == "hoho" {
		fmt.Println("wow it's hoho")
		return false
	}

	return true
}
func deletePredicate(e event.DeleteEvent) bool {
	fmt.Println("delete predicate ")
	return true
}
