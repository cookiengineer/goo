package main

import "fmt"
import "goo"
import "os"
import "os/exec"
import "path/filepath"

func main() {

	cwd, err1 := os.Getwd()

	if err1 == nil {

		goo_folder, err2 := filepath.Abs(cwd + "/../../")

		if err2 == nil {
			os.Setenv("GOO_FOLDER", goo_folder)
		}

		result := goo.Build(goo.AppLayout, goo.DefaultTheme, cwd + "/public/goo")

		if result == true {

			cmd := exec.Command("go", "build", "-o", cwd + "/basic.exe", cwd + "/main.go")

			err3 := cmd.Run()

			if err3 == nil {
				fmt.Println("Build successful!")
			}

		} else {
			fmt.Println("Build failed! :(")
		}

	}

}
