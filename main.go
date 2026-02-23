package main

import (
	"context"
	"fmt"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	ctx := context.Background()
	fmt.Println("=== Kubernetes Security Audit ===\n")

	// Check for resource limits
	pods, _ := clientset.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	noLimits := 0
	for _, p := range pods.Items {
		for _, c := range p.Spec.Containers {
			if c.Resources.Limits == nil || len(c.Resources.Limits) == 0 {
				noLimits++
			}
		}
	}
	fmt.Printf("Pods without resource limits: %d\n", noLimits)

	// Check namespaces
	ns, _ := clientset.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	fmt.Printf("Namespaces: %d\n", len(ns.Items))

	// Check network policies
	np, _ := clientset.NetworkingV1().NetworkPolicies("").List(ctx, metav1.ListOptions{})
	fmt.Printf("Network policies: %d\n", len(np.Items))

	fmt.Println("\nAudit complete.")
}
