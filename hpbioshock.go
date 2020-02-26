package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func errcheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	file, err := os.Open("brutalist.txt")
	errcheck(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pass := scanner.Text()
		log.Println("Testing " + pass + "...")
		out, err := exec.Command("echo", pass).CombinedOutput()
		errcheck(err)
		fmt.Println(string(out))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
