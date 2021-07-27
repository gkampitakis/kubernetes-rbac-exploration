package main

import (
	"context"
	"fmt"
	"log"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	log.Println("Welcome to K8s Runner")

	client, err := getK8sClient()
	if err != nil {
		panic(err.Error())
	}

	pods, err := client.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	log.Println(fmt.Sprintf("There are %d pods in the cluster\n", len(pods.Items)))

	switch os.Args[1] {
	case "create":

	case "delete":

	case "list":

	default:
		log.Fatal("Unknown operation")
	}
}

func getK8sClient() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		return nil, err
	}

	return clientset, nil
}
