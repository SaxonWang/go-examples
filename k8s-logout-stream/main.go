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

	ctx := context.TODO()
	byteReader, err := clientset.CoreV1().Pods("hidevopsio-alpha").
		GetLogs("hidevopsio-log-6b94b49dbc-xr27t", &v1.PodLogOptions{Follow: true}).Context(ctx).Stream()
	if err != nil {
		return
	}

	reader := bufio.NewReader(byteReader)
	err = nil
	for err == nil {
		str, err := reader.ReadString('\n')
		fmt.Println(str)
		if err != nil {
			break
		}
		if err != nil {
			break
		}
	}
	if err == io.EOF {
		return
	}
}