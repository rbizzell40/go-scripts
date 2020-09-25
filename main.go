package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
     // Checking Cluster Status
	out, err := exec.Command("kubectl", "cluster-info", "--context", "kind-kind").Output()
	if err != nil {
		log.Fatal(err)
		fmt.Println("kind cluster does not exist, will deploy cluster")
		output := string(out[:])
		fmt.Println(output)

		createCluster := exec.Command("kind", "create", "cluster")

		createCluster.Stdout = os.Stdout
		createCluster.Stderr = os.Stdout

		if err := createCluster.Run(); err != nil {
			log.Fatal(err)
			fmt.Println("Error:", err)
		}
	}else {
			fmt.Println("kind cluster exists, will delete cluster")
			out, err := exec.Command("kind", "delete", "cluster").Output()
			if err != nil {
				log.Fatal(err)
				fmt.Printf("%s", err)
			}
			fmt.Println("Deploying new cluster")
			output := string(out[:])
			fmt.Println(output)

			createCluster := exec.Command("kind", "create", "cluster")

			createCluster.Stdout = os.Stdout
			createCluster.Stderr = os.Stdout

			if err := createCluster.Run(); err != nil {
				log.Fatal(err)
				fmt.Println("Error", err)
			}
		}

	}


