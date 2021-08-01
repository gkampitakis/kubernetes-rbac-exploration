package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	log.Println("Welcome to K8s Runner")

	namespace := "default"
	if len(os.Args) >= 3 {
		namespace = os.Args[2]
	}

	log.Println(fmt.Sprintf("Selected namespace %s", namespace))

	client, err := getK8sClient()
	if err != nil {
		panic(err.Error())
	}

	switch os.Args[1] {
	case "create":
		cmd := exec.Command("kubectl", "run", "busybox", "--image=busybox", "--restart=Never", "-n", namespace)
		cmd.Stderr = os.Stdout
		err := cmd.Run()
		if err != nil {
			panic(err.Error())
		}
	case "delete":
		cmd := exec.Command("kubectl", "delete", "--all", "pods", "-n", namespace)
		cmd.Stderr = os.Stdout
		err := cmd.Run()
		if err != nil {
			panic(err.Error())
		}
	case "list":

		pods, err := client.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}

		log.Println(fmt.Sprintf("There are %d pods in the cluster\n", len(pods.Items)))
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
