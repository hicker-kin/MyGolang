package main

import (
	"flag"
	"time"

	"github.com/golang/glog"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	// Uncomment the following line to load the gcp plugin (only required to authenticate against GKE clusters).
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	clientset "k8s_customize_controller/pkg/client/clientset/versioned"
	informers "k8s_customize_controller/pkg/client/informers/externalversions"
	contrller "k8s_customize_controller/pkg/controller"
	"k8s_customize_controller/pkg/signals"
)

var (
	masterURL  string
	kubeConfig string
)

// go build
// ./k8s_customize_controller --kubeConfig=$HOME/.kube/config -alsologtostderr=true
func main() {
	flag.Parse()

	glog.Infof("param %s, %s", masterURL, kubeConfig)

	// 处理信号量
	stopCh := signals.SetupSignalHandler()

	// 处理入参
	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeConfig)
	// local debug
	// cfg, err := clientcmd.BuildConfigFromFlags(masterURL, "/Users/my/.kube/config")
	if err != nil {
		glog.Fatalf("Error building kubeConfig: %s", err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	studentClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building example clientset: %s", err.Error())
	}

	studentInformerFactory := informers.NewSharedInformerFactory(studentClient, time.Second*30)

	// 得到controller
	controller := contrller.NewController(kubeClient, studentClient,
		studentInformerFactory.Stable().V1().Students())

	// 启动informer
	go studentInformerFactory.Start(stopCh)

	// controller开始处理消息
	if err = controller.Run(2, stopCh); err != nil {
		glog.Fatalf("Error running controller: %s", err.Error())
	}
}

func init() {
	flag.StringVar(&kubeConfig, "kubeConfig", "", "Path to a kubeConfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeConfig. Only required if out-of-cluster.")
}
