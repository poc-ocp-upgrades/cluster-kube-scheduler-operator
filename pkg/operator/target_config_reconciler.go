package operator

import (
	"fmt"
	"reflect"
	"time"
	"k8s.io/klog"
	operatorv1 "github.com/openshift/api/operator/v1"
	configinformers "github.com/openshift/client-go/config/informers/externalversions"
	configlistersv1 "github.com/openshift/client-go/config/listers/config/v1"
	operatorv1client "github.com/openshift/client-go/operator/clientset/versioned/typed/operator/v1"
	operatorv1informers "github.com/openshift/client-go/operator/informers/externalversions/operator/v1"
	"github.com/openshift/cluster-kube-scheduler-operator/pkg/operator/operatorclient"
	"github.com/openshift/library-go/pkg/operator/events"
	"github.com/openshift/library-go/pkg/operator/v1helpers"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	corev1listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

type TargetConfigReconciler struct {
	targetImagePullSpec	string
	operatorConfigClient	operatorv1client.KubeSchedulersGetter
	kubeClient		kubernetes.Interface
	eventRecorder		events.Recorder
	configMapLister		corev1listers.ConfigMapLister
	SchedulerLister		configlistersv1.SchedulerLister
	SchedulingCacheSync	cache.InformerSynced
	featureGateLister	configlistersv1.FeatureGateLister
	featureGateCacheSync	cache.InformerSynced
	queue			workqueue.RateLimitingInterface
}

func NewTargetConfigReconciler(targetImagePullSpec string, operatorConfigInformer operatorv1informers.KubeSchedulerInformer, namespacedKubeInformers informers.SharedInformerFactory, kubeInformersForNamespaces v1helpers.KubeInformersForNamespaces, configInformer configinformers.SharedInformerFactory, operatorConfigClient operatorv1client.KubeSchedulersGetter, kubeClient kubernetes.Interface, eventRecorder events.Recorder) *TargetConfigReconciler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	c := &TargetConfigReconciler{targetImagePullSpec: targetImagePullSpec, operatorConfigClient: operatorConfigClient, kubeClient: kubeClient, configMapLister: kubeInformersForNamespaces.ConfigMapLister(), eventRecorder: eventRecorder, SchedulerLister: configInformer.Config().V1().Schedulers().Lister(), SchedulingCacheSync: configInformer.Config().V1().Schedulers().Informer().HasSynced, featureGateLister: configInformer.Config().V1().FeatureGates().Lister(), featureGateCacheSync: configInformer.Config().V1().FeatureGates().Informer().HasSynced, queue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "TargetConfigReconciler")}
	configInformer.Config().V1().Schedulers().Informer().AddEventHandler(c.eventHandler())
	operatorConfigInformer.Informer().AddEventHandler(c.eventHandler())
	namespacedKubeInformers.Rbac().V1().Roles().Informer().AddEventHandler(c.eventHandler())
	namespacedKubeInformers.Rbac().V1().RoleBindings().Informer().AddEventHandler(c.eventHandler())
	namespacedKubeInformers.Core().V1().ConfigMaps().Informer().AddEventHandler(c.eventHandler())
	namespacedKubeInformers.Core().V1().Secrets().Informer().AddEventHandler(c.eventHandler())
	namespacedKubeInformers.Core().V1().ServiceAccounts().Informer().AddEventHandler(c.eventHandler())
	namespacedKubeInformers.Core().V1().Services().Informer().AddEventHandler(c.eventHandler())
	configInformer.Config().V1().Schedulers().Informer().AddEventHandler(c.eventHandler())
	configInformer.Config().V1().FeatureGates().Informer().AddEventHandler(c.eventHandler())
	kubeInformersForNamespaces.InformersFor(operatorclient.GlobalUserSpecifiedConfigNamespace).Core().V1().ConfigMaps().Informer().AddEventHandler(c.eventHandler())
	kubeInformersForNamespaces.InformersFor(operatorclient.TargetNamespace).Core().V1().ConfigMaps().Informer().AddEventHandler(c.eventHandler())
	namespacedKubeInformers.Core().V1().Namespaces().Informer().AddEventHandler(c.namespaceEventHandler())
	return c
}
func (c TargetConfigReconciler) sync() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	operatorConfig, err := c.operatorConfigClient.KubeSchedulers().Get("cluster", metav1.GetOptions{})
	if err != nil {
		return err
	}
	operatorConfigOriginal := operatorConfig.DeepCopy()
	switch operatorConfig.Spec.ManagementState {
	case operatorv1.Unmanaged:
		return nil
	case operatorv1.Removed:
		return nil
	}
	requeue, err := createTargetConfigReconciler_v311_00_to_latest(c, c.eventRecorder, operatorConfig)
	if requeue && err == nil {
		return fmt.Errorf("synthetic requeue request")
	}
	if err != nil {
		if !reflect.DeepEqual(operatorConfigOriginal, operatorConfig) {
			v1helpers.SetOperatorCondition(&operatorConfig.Status.Conditions, operatorv1.OperatorCondition{Type: operatorv1.OperatorStatusTypeDegraded, Status: operatorv1.ConditionTrue, Reason: "StatusUpdateError", Message: err.Error()})
			if _, updateError := c.operatorConfigClient.KubeSchedulers().UpdateStatus(operatorConfig); updateError != nil {
				klog.Error(updateError)
			}
		}
		return err
	}
	return nil
}
func (c *TargetConfigReconciler) Run(workers int, stopCh <-chan struct{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer utilruntime.HandleCrash()
	defer c.queue.ShutDown()
	klog.Infof("Starting TargetConfigReconciler")
	defer klog.Infof("Shutting down TargetConfigReconciler")
	if !cache.WaitForCacheSync(stopCh, c.SchedulingCacheSync) {
		utilruntime.HandleError(fmt.Errorf("scheduler caches did not sync"))
		return
	}
	go wait.Until(c.runWorker, time.Second, stopCh)
	<-stopCh
}
func (c *TargetConfigReconciler) runWorker() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for c.processNextWorkItem() {
	}
}
func (c *TargetConfigReconciler) processNextWorkItem() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	dsKey, quit := c.queue.Get()
	if quit {
		return false
	}
	defer c.queue.Done(dsKey)
	err := c.sync()
	if err == nil {
		c.queue.Forget(dsKey)
		return true
	}
	utilruntime.HandleError(fmt.Errorf("%v failed with : %v", dsKey, err))
	c.queue.AddRateLimited(dsKey)
	return true
}
func (c *TargetConfigReconciler) eventHandler() cache.ResourceEventHandler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return cache.ResourceEventHandlerFuncs{AddFunc: func(obj interface{}) {
		c.queue.Add(workQueueKey)
	}, UpdateFunc: func(old, new interface{}) {
		c.queue.Add(workQueueKey)
	}, DeleteFunc: func(obj interface{}) {
		c.queue.Add(workQueueKey)
	}}
}

var interestingNamespaces = sets.NewString(operatorclient.TargetNamespace)

func (c *TargetConfigReconciler) namespaceEventHandler() cache.ResourceEventHandler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return cache.ResourceEventHandlerFuncs{AddFunc: func(obj interface{}) {
		ns, ok := obj.(*corev1.Namespace)
		if !ok {
			c.queue.Add(workQueueKey)
		}
		if ns.Name == operatorclient.TargetNamespace {
			c.queue.Add(workQueueKey)
		}
	}, UpdateFunc: func(old, new interface{}) {
		ns, ok := old.(*corev1.Namespace)
		if !ok {
			c.queue.Add(workQueueKey)
		}
		if ns.Name == operatorclient.TargetNamespace {
			c.queue.Add(workQueueKey)
		}
	}, DeleteFunc: func(obj interface{}) {
		ns, ok := obj.(*corev1.Namespace)
		if !ok {
			tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
			if !ok {
				utilruntime.HandleError(fmt.Errorf("couldn't get object from tombstone %#v", obj))
				return
			}
			ns, ok = tombstone.Obj.(*corev1.Namespace)
			if !ok {
				utilruntime.HandleError(fmt.Errorf("tombstone contained object that is not a Namespace %#v", obj))
				return
			}
		}
		if ns.Name == operatorclient.TargetNamespace {
			c.queue.Add(workQueueKey)
		}
	}}
}
