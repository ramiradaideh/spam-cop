// package main

// import (
// 	"context"
// 	"fmt"
// 	"os"
// 	"path/filepath"

// 	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
// 	"k8s.io/client-go/kubernetes"
// 	"k8s.io/client-go/tools/clientcmd"
// 	"sigs.k8s.io/yaml"
// )

// func getClientset() *kubernetes.Clientset {
// 	kubeconfig := os.Getenv("KUBECONFIG")
// 	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	clientset, err := kubernetes.NewForConfig(config)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return clientset
// }

// func runSnapshot() {
// 	clientset := getClientset()
// 	namespace := "default"  ///Dynamically get all the namespaces 

// 	deployments, err := clientset.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	for _, d := range deployments.Items {
// 		saveResourceToFile("snapshots", fmt.Sprintf("%s-deployment.yaml", d.Name), d)
// 	}
// 	fmt.Println("Snapshots saved successfully.")
// }

// func saveResourceToFile(dir, filename string, resource interface{}) {
// 	data, err := yaml.Marshal(resource)
// 	if err != nil {
// 		panic(err)
// 	}

// 	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
// 		panic(err)
// 	}

// 	if err := os.WriteFile(filepath.Join(dir, filename), data, 0644); err != nil {
// 		panic(err)
// 	}
// }
