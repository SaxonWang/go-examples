package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/Users/wang/.kube/config", "(optional) absolute path to the kubeconfig file")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	pod,err := clientset.CoreV1().Pods("demo").Get("hiadmin-26-cqmjd",meta_v1.GetOptions{})
	if err != nil {
		fmt.Println("Error ",err)
		return
	}

	fmt.Println("pod:",pod)

	ctx := context.TODO()
	byteReader, err := clientset.CoreV1().Pods("demo").
		GetLogs("hiadmin-26-cqmjd", &v1.PodLogOptions{Follow: true}).Context(ctx).Stream()
	if err != nil {
		fmt.Println("Error ",err)
		return
	}

	reader := bufio.NewReader(byteReader)
	err = nil
	for err == nil {
		str, err := reader.ReadString('\n')
		fmt.Println(str)
		if err != nil {
			fmt.Println("Error ",err)
			break
		}
		if err != nil {
			fmt.Println("Error ",err)
			break
		}
	}
	if err == io.EOF {
		fmt.Println("Error ",err)
		return
	}
}