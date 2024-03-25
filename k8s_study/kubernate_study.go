package k8s_study

import (
	"context"
	"fmt"
	"io/ioutil"
	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	BearerToken = "eyJhbGciOiJSUzI1NiIsImtpZCI6IjZBNEt4ZUJaWG9OU2FCa0RUU2s0NkxhdXJPTTI1SlB2X2dyWXo1SndVVFkifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJkZWZhdWx0Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6ImRldi10b2tlbi00MjhzciIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50Lm5hbWUiOiJkZXYiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC51aWQiOiIxNmFmZjc4ZC01ODhmLTQ5ZDUtYTNjMy1mODA2OGIxMDUxMzgiLCJzdWIiOiJzeXN0ZW06c2VydmljZWFjY291bnQ6ZGVmYXVsdDpkZXYifQ.jKZ6c6vUJZEGF7fZrBRY3hcLlBZBNqMRa3qNNIe3F6SyPLyUPMFOFxoga5kYvw6angbjGsVE-7ZG9BVm9cboQgDghsKr5diVvCgAbcsiRdSem4KzaiGE-mVrZO1LcWC0df4_n56kp5E-csGuM9tp-pViMZ3Nz3lRcMsnmds0DYHDMCqup164uEP9AEPlKN42qvgBBjJPCENGvaGmtiBXIBCSBn1xJw3WQXvErFp42-eIEG0HuR8zl9cRPDm8dl5vRDzpQZpS2eHvhhhGHXgTbZqXqb0jncJxmNotpgXqFlsGFAixkZo33QguEXd8C62g1n3uN08WocKL7ol79YCzhA"

	CAPEMString = `-----BEGIN CERTIFICATE-----
MIIC5zCCAc+gAwIBAgIBADANBgkqhkiG9w0BAQsFADAVMRMwEQYDVQQDEwprdWJl
cm5ldGVzMB4XDTIyMDYwNjAxNTkyMFoXDTMyMDYwMzAxNTkyMFowFTETMBEGA1UE
AxMKa3ViZXJuZXRlczCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBALEN
VyFapvQKMOfvo0sPlRpcNcgYZCv/0pvYPohtltIYbUNDXDIjRJ7Ccay6R5P220Ws
E9rIoiqWMOHpnd8huMhBXXrgCQXKXBGp8rfVqf94oo/1dyiFGnGilFf4hVPBwd9B
TdwhOKQEv4lm/NsT2VABrz/Xh92adNjtpH0MEYXkUb00ocE8CNitWyu0hejgpBpp
awMyCDL5K8SJ9XQm2Ehi53YePoxykvQw1U+6/QJGQ0F8JuHMa1x6rxKBNHqKSqYH
1XuQK6wjjmzslubxtG4kI19YE+Kqm6mFQahkXtbSISPB1Rwt7kqSs7xxKDwItOCu
FjoWGFLcmigq6EpMWc0CAwEAAaNCMEAwDgYDVR0PAQH/BAQDAgKkMA8GA1UdEwEB
/wQFMAMBAf8wHQYDVR0OBBYEFMnpS4x3ADeUaX54djL63jIlJ/6gMA0GCSqGSIb3
DQEBCwUAA4IBAQBMj04FAx5D8RhqQ0P0jti1sqUrQ0r6uc73CY9tXXG07pYEuU9G
WNsRH5Ouy3d78qlTeyjV2ionZGKoM0MySzBIEjzElGsU8rM3vBa/pjAzSlViovSl
EAtRYxOMjp0rZjnp3KYsNfDac9HLo/ICZ14NOonHyTQunQClCWQGtIBMX/D30CEP
tUTOI0+DXETtJdefeJ/YbjIHSC80Pjhl9NY3/A+O3txz/UpzK48QgiCXgytpt3MB
HRrjhJigjJ/sBcPRMT3O9NJQIc+jcA3aNSCM/GzwdHjQSJI6L2AvR/JZXirZJXIy
SvZRT5oQPEKqf80z/tSKDELVhJaklZ65Bp5R
-----END CERTIFICATE-----`
)

func GetNodeResource() {
	log.SetFlags(log.Llongfile | log.Ltime)

	tlsClientConfig := rest.TLSClientConfig{
		CAData: []byte(CAPEMString),
	}
	config := rest.Config{
		Host:            "lb.kubesphere.local:6443",
		BearerToken:     BearerToken,
		TLSClientConfig: tlsClientConfig,
		Timeout:         20 * time.Second,
	}
	clientSet, err := kubernetes.NewForConfig(&config)
	if err != nil {
		log.Println(err)
		return
	}
	nodes, err := clientSet.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})

	if err != nil {
		fmt.Println("11111", err)
		return
	}
	for _, nds := range nodes.Items {
		//fmt.Println("-------Labels", nds.Labels)
		fmt.Println("-------Allocatable", nds.Status.Allocatable.Memory().String())
		fmt.Println("-------Allocatable", nds.Status.Allocatable.Storage().String())
		fmt.Println("-------------------Allocatable", fmt.Sprintf("%+v", nds.Status.Allocatable))
		fmt.Println("-------------------Capacity", fmt.Sprintf("%+v", nds.Status.Capacity))
	}
}

func StudyK8s() {
	log.SetFlags(log.Llongfile | log.Ltime)

	wg := sync.WaitGroup{}

	tlsClientConfig := rest.TLSClientConfig{
		CAData: []byte(CAPEMString),
	}
	config := rest.Config{
		Host:            "lb.kubesphere.local:6443",
		BearerToken:     BearerToken,
		TLSClientConfig: tlsClientConfig,
		Timeout:         20 * time.Second,
	}
	clientSet, err := kubernetes.NewForConfig(&config)
	if err != nil {
		log.Println(err)
		return
	}

	//开启一个集群监听
	go func() {
		wg.Add(1)
		defer wg.Done()
		watcher, err := clientSet.CoreV1().Events("default").Watch(context.Background(), metav1.ListOptions{})
		if err != nil {
			fmt.Println("-----", err)
			return
		}
		eventChan := watcher.ResultChan()
		var i = 0
		for {
			event := <-eventChan
			fmt.Println(event.Type, i)
			i++
			time.Sleep(3 * time.Second)
		}
	}()

	//一个简单的例子：创建一个 deployment，启动3个 busybox 副本
	container := corev1.Container{
		Name:    "busybox",
		Image:   "busybox:latest",
		Command: []string{"top"},
	}
	podSpec := corev1.PodSpec{
		Containers: []corev1.Container{container},
	}
	podTemplateSpec := corev1.PodTemplateSpec{
		ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{"dev": "busybox"}},
		Spec:       podSpec,
	}
	spec := appv1.DeploymentSpec{
		Selector: &metav1.LabelSelector{
			MatchLabels: map[string]string{
				"dev": "busybox",
			},
		},
		Replicas: Int32Value(1),
		Template: podTemplateSpec,
	}
	deployment := appv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "busybox-test",
		},
		Spec: spec,
	}
	resultDeployment, err := clientSet.AppsV1().Deployments("default").Get(context.Background(), "busybox-test", metav1.GetOptions{})
	fmt.Println("+++++", err)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			resultDeployment, err = clientSet.AppsV1().Deployments("default").Create(context.Background(), &deployment, metav1.CreateOptions{})
			if err != nil {
				log.Println(err)
				return
			}
		}
		return
	} else {
		resultDeployment, err = clientSet.AppsV1().Deployments("default").Update(context.Background(), &deployment, metav1.UpdateOptions{})
		if err != nil {
			log.Println(err)
			return
		}
	}

	log.Println(resultDeployment.Name)
	wg.Wait()
	time.Sleep(5 * time.Minute)
}

func Int32Value(a int32) *int32 {
	return &a
}

func GetK8sLog() {
	tlsClientConfig := rest.TLSClientConfig{
		CAData: []byte(CAPEMString),
	}
	config := rest.Config{
		Host:            "lb.kubesphere.local:6443",
		BearerToken:     BearerToken,
		TLSClientConfig: tlsClientConfig,
		Timeout:         20 * time.Second,
	}
	clientSet, err := kubernetes.NewForConfig(&config)
	if err != nil {
		log.Println(err)
		return
	}

	var tailLines int64
	// 生成获取POD日志请求
	req := clientSet.CoreV1().Pods("default").GetLogs("busybox-test-784f8874f7-k9kf6", &corev1.PodLogOptions{Container: "busybox", TailLines: &tailLines})

	// req.Stream()也可以实现Do的效果

	// 发送请求
	res := req.Do(context.Background())
	if res.Error() != nil {
		fmt.Println("wetefsadfsdgsdgsdgsdgsdf")
		err = res.Error()
		fmt.Println("--------1", err)
	}

	// 获取结果
	logs, err := res.Raw()
	if err != nil {
		fmt.Println("---------2", err)
	}

	fmt.Println("ssssssssssslllllllllllll", string(logs))
}

func GetPodInfo() {
	tlsClientConfig := rest.TLSClientConfig{
		CAData: []byte(CAPEMString),
	}
	config := rest.Config{
		Host:            "lb.kubesphere.local:6443",
		BearerToken:     BearerToken,
		TLSClientConfig: tlsClientConfig,
		Timeout:         20 * time.Second,
	}
	clientSet, err := kubernetes.NewForConfig(&config)
	if err != nil {
		log.Println(err)
		return
	}

	clientSet.CoreV1().Pods("default").Watch(context.Background(), metav1.ListOptions{})

	podInfo, err := clientSet.CoreV1().Pods("default").Get(context.TODO(), "busybox-test-784f8874f7-k9kf6", metav1.GetOptions{})
	if err != nil {
		fmt.Println("111111", err)
	}

	for _, status := range podInfo.Status.Conditions {
		fmt.Println("----", fmt.Sprintf("%+v", status))
	}
}

func GetPodLog() {
	tlsClientConfig := rest.TLSClientConfig{
		CAData: []byte(CAPEMString),
	}
	config := rest.Config{
		Host:            "lb.kubesphere.local:6443",
		BearerToken:     BearerToken,
		TLSClientConfig: tlsClientConfig,
		Timeout:         20 * time.Second,
	}
	clientSet, err := kubernetes.NewForConfig(&config)
	if err != nil {
		log.Println(err)
		return
	}
	var tailLine int64
	logOpt := &corev1.PodLogOptions{
		Container: "busybox",
		Follow:    true,
		TailLines: &tailLine,
		Previous:  false,
	}

	req := clientSet.CoreV1().Pods("default").GetLogs("busybox-test-784f8874f7-k9kf6", logOpt)
	stream, err := req.Stream(context.TODO())

	if err != nil {
		fmt.Println("222222", err)
	}

	all, err := ioutil.ReadAll(stream)

	if err != nil {
		fmt.Println("3333333", err)
	}

	fmt.Println("uuuuuuuuu", string(all))
}

func GetK8sResource() {
	log.SetFlags(log.Llongfile | log.Ltime)

	tlsClientConfig := rest.TLSClientConfig{
		CAData: []byte(CAPEMString),
	}
	config := rest.Config{
		Host:            "lb.kubesphere.local:6443",
		BearerToken:     BearerToken,
		TLSClientConfig: tlsClientConfig,
		Timeout:         20 * time.Second,
	}
	clientset, err := kubernetes.NewForConfig(&config)

	if err != nil {
		fmt.Println("5555", err)
		return
	}

	restClient := clientset.RESTClient()

	discoveryClient := discovery.NewDiscoveryClient(restClient)

	resources, err := discoveryClient.ServerPreferredResources()

	if err != nil {
		fmt.Println("6666", err)
	}

	fmt.Println("---------------", fmt.Sprintf("%+v", resources))
}

func NewK8sResource() {
	// 创建一个Kubernetes客户端
	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Printf("Failed to create Kubernetes client config: %v\n", err)
		os.Exit(1)
	}

	// 2. 创建Kubernetes客户端
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	// 3. 通过nodeName获取机器的可使用资源
	nodeName := "devnode37"
	node, err := clientset.CoreV1().Nodes().Get(ctx, nodeName, metav1.GetOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for {
		fmt.Printf("Node name: %s\n", node.Name)
		fmt.Printf("Node CPU: %v\n", node.Status.Allocatable.Cpu().MilliValue())
		fmt.Printf("Node Memory: %v\n", node.Status.Allocatable.Memory().Value())

		// 获取节点的使用情况
		nodeUsage, err := clientset.CoreV1().Nodes().Get(ctx, nodeName, metav1.GetOptions{})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Node name: %s\n", node.Name)
		fmt.Printf("Node CPU usage: %v\n", nodeUsage.Status.Allocatable.Cpu().MilliValue()-node.Status.Allocatable.Cpu().MilliValue())
		fmt.Printf("Node Memory usage: %v\n", nodeUsage.Status.Allocatable.Memory().Value()-node.Status.Allocatable.Memory().Value())
		fmt.Println("-----------------------------------------------------------------")
		time.Sleep(1 * time.Minute)
	}
}
