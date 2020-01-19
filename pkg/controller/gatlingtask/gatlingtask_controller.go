package gatlingtask

import (
	"context"
	"reflect"

	logr "github.com/go-logr/logr"
	tpokkiv1alpha1 "github.com/tpokki/gatling-operator/pkg/apis/tpokki/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_gatlingtask")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new GatlingTask Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileGatlingTask{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("gatlingtask-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource GatlingTask
	err = c.Watch(&source.Kind{Type: &tpokkiv1alpha1.GatlingTask{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// Watch for changes to secondary resource Pods and requeue the owner GatlingTask
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &tpokkiv1alpha1.GatlingTask{},
	})
	if err != nil {
		return err
	}

	// Watch for changes to secondary resource ConfigMap and requeue the owner GatlingTask
	err = c.Watch(&source.Kind{Type: &corev1.ConfigMap{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &tpokkiv1alpha1.GatlingTask{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileGatlingTask implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileGatlingTask{}

// ReconcileGatlingTask reconciles a GatlingTask object
type ReconcileGatlingTask struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a GatlingTask object and makes changes based on the state read
// and what is in the GatlingTask.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileGatlingTask) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling GatlingTask")

	// Fetch the GatlingTask instance
	instance := &tpokkiv1alpha1.GatlingTask{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// Define new Configmap object
	configMap := newConfigMapForCR(instance)

	//configMap, err := r.updateConfigMap(reqLogger, instance)
	err = r.updateObject(reqLogger, instance, &TaskObject{configMap})
	if err != nil {
		return reconcile.Result{}, err
	}

	// Define a new Pod object
	pod := newPodForCR(instance, configMap)

	// Set GatlingTask instance as the owner and controller
	if err := controllerutil.SetControllerReference(instance, pod, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	// Check if this Pod already exists
	found := &corev1.Pod{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: pod.Name, Namespace: pod.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new Pod", "Pod.Namespace", pod.Namespace, "Pod.Name", pod.Name)
		err = r.client.Create(context.TODO(), configMap)
		if err != nil {
			return reconcile.Result{}, err
		}

		err = r.client.Create(context.TODO(), pod)
		if err != nil {
			return reconcile.Result{}, err
		}

		// Pod created successfully - don't requeue
		return reconcile.Result{}, nil
	} else if err != nil {
		return reconcile.Result{}, err
	}

	// Pod already exists - don't requeue
	reqLogger.Info("Skip reconcile: Pod already exists", "Pod.Namespace", found.Namespace, "Pod.Name", found.Name)
	return reconcile.Result{}, nil
}

type TaskObject struct {
	object interface{}
}

func (t *TaskObject) runtimeObject() runtime.Object {
	return t.object.(runtime.Object)
}

func (t *TaskObject) metav1Object() metav1.Object {
	return t.object.(metav1.Object)
}

func (r *ReconcileGatlingTask) updateObject(logr logr.Logger, cr *tpokkiv1alpha1.GatlingTask, object *TaskObject) error {

	// Define new Configmap object
	//	configMap := newConfigMapForCR(cr)

	// Set GatlingTask instance as the owner and controller
	if err := controllerutil.SetControllerReference(cr, object.metav1Object(), r.scheme); err != nil {
		return err
	}

	// Check if this object already exists
	found := object.runtimeObject()
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: object.metav1Object().GetName(), Namespace: object.metav1Object().GetNamespace()}, found)
	if err != nil && errors.IsNotFound(err) {
		logr.Info("Creating new Object", "Type", reflect.TypeOf(object.object), "Namespace", object.metav1Object().GetNamespace(), "Name", object.metav1Object().GetName())
		err = r.client.Create(context.TODO(), object.runtimeObject())
		if err != nil {
			return err
		}

		// Pod created successfully - don't requeue
		return nil
	} else if err != nil {
		return err
	}

	// ConfigMap already exists - don't requeue
	logr.Info("Skip reconcile: Object already exists", "Type", reflect.TypeOf(object.object), "Namespace", object.metav1Object().GetNamespace(), "Name", object.metav1Object().GetName())
	return nil
}

func (r *ReconcileGatlingTask) updateConfigMap(logr logr.Logger, cr *tpokkiv1alpha1.GatlingTask) (*corev1.ConfigMap, error) {

	// Define new Configmap object
	configMap := newConfigMapForCR(cr)

	// Set GatlingTask instance as the owner and controller
	if err := controllerutil.SetControllerReference(cr, configMap, r.scheme); err != nil {
		return nil, err
	}

	// Check if this ConfigMap already exists
	found := &corev1.ConfigMap{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: configMap.Name, Namespace: configMap.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		logr.Info("Creating new ConfigMap", "ConfigMap.Namespace", configMap.Namespace, "ConfigMap.Name", configMap.Name)
		err = r.client.Create(context.TODO(), configMap)
		if err != nil {
			return nil, err
		}

		// Pod created successfully - don't requeue
		return configMap, nil
	} else if err != nil {
		return nil, err
	}

	// ConfigMap already exists - don't requeue
	logr.Info("Skip reconcile: Pod already exists", "Pod.Namespace", found.Namespace, "Pod.Name", found.Name)
	return configMap, nil
}

func newConfigMapForCR(cr *tpokkiv1alpha1.GatlingTask) *corev1.ConfigMap {
	labels := map[string]string{
		"app": cr.Name,
	}
	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-configmap",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Data: map[string]string{
			cr.Spec.ScenarioSpec.Name: cr.Spec.ScenarioSpec.Definition,
		},
	}
}

// newPodForCR returns a busybox pod with the same name/namespace as the cr
func newPodForCR(cr *tpokkiv1alpha1.GatlingTask, cm *corev1.ConfigMap) *corev1.Pod {
	labels := map[string]string{
		"app": cr.Name,
	}
	volumeName := "configmap-scenario"
	volumePath := "/scenario/input"

	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-pod",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: corev1.PodSpec{
			Volumes: []corev1.Volume{
				{
					Name: volumeName,
					VolumeSource: corev1.VolumeSource{
						ConfigMap: &corev1.ConfigMapVolumeSource{
							LocalObjectReference: corev1.LocalObjectReference{
								Name: cm.ObjectMeta.Name,
							},
						},
					},
				},
			},
			RestartPolicy: "Never",
			Containers: []corev1.Container{
				{
					Name:      "gatling",
					Image:     "busybox",
					Command:   []string{"sleep", "3600"},
					Resources: cr.Spec.ResourceRequirements,
					VolumeMounts: []corev1.VolumeMount{
						{
							Name:      volumeName,
							MountPath: volumePath,
						},
					},
				},
			},
		},
	}
}
