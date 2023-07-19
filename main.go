package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
)

func tryPullFromGit(path string, wg *sync.WaitGroup) {

	wg.Add(1)
	defer wg.Done()

	cwd , _ := os.Getwd()
	err := os.Chdir( cwd + path)
	if err != nil {
		fmt.Printf("Error with chdir:%v \n" , err)
	}

    cwd2 , _ := os.Getwd()
	fmt.Printf("123 , %v \n", cwd2)

	cmd := exec.Command("git", "pull", "origin", "master")
	output , err := cmd.Output();

	if  err != nil {
		log.Printf("Error in file : %v , %v", path, err)
	}
	log.Println(string(output))
}

func main() {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatalf("%v", err)
	}

	var jobWG sync.WaitGroup

	for _, file := range files {
		fmt.Println(file.Name())
		if file.IsDir() {
			fmt.Println("yes")
			go func() {
				folderPath := "./" + file.Name()
				tryPullFromGit(folderPath, &jobWG)
			}()
		}
	}
	jobWG.Wait()

}
