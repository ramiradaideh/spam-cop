package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/yaml"
)

func main() {
	// Ensure correct number of arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [snapshot|deploy|rollback]")
		os.Exit(1)
	}

	command := os.Args[1]
	// Handle different commands
	switch command {
	case "snapshot":
		runSnapshot()
	// case "deploy":
	// 	runDeploy()
	// case "rollback":
	// 	runRollback()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}


func getClientset() *kubernetes.Clientset {
	kubeconfig := os.Getenv("KUBECONFIG")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientset
}

func runSnapshot() {
	clientset := getClientset()
	namespace := "default"

	deployments, err := clientset.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for _, d := range deployments.Items {
		saveResourceToFile("snapshots", fmt.Sprintf("%s-deployment.yaml", d.Name), d)
	}
	fmt.Println("Snapshots saved successfully.")
}

func saveResourceToFile(dir, filename string, resource interface{}) {
	data, err := yaml.Marshal(resource)
	if err != nil {
		panic(err)
	}

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		panic(err)
	}

	if err := os.WriteFile(filepath.Join(dir, filename), data, 0644); err != nil {
		panic(err)
	}
}
