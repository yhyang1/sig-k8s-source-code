package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "custom-contexts/kubernetes/kubernetes.yml"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	config.APIPath = "api"
	config.GroupVersion = &corev1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs

	DemoRestClient(config)
	DemoClientSet(config)
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func DemoRestClient(config *rest.Config) {
	// create the clientset
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err.Error())
	}

	result := &corev1.PodList{}
	err = restClient.Get().Namespace("kube-system").Resource("pods").
		VersionedParams(&metav1.ListOptions{Limit: 500}, scheme.ParameterCodec).
		Do(context.TODO()).Into(result)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Demo RESTClient:")
	for _, d := range result.Items {
		fmt.Println(d.Namespace, d.Name, d.Status.Phase)
	}
}

func DemoClientSet(config *rest.Config) {
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	podClient := clientset.CoreV1().Pods(corev1.NamespaceAll) // k8s.io/core/v1.NamespaceAll
	list, err := podClient.List(context.TODO(), metav1.ListOptions{Limit: 500})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Demo ClientSet:")
	for _, d := range list.Items {
		fmt.Println(d.Namespace, d.Name, d.Status.Phase)
	}
}

func DynamicClient(config *rest.Config) {
	dynamicClient, err := dynamec.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

}

func DiscoveryClient(config *rest.Config) {

}
