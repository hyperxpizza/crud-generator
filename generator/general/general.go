package general

import (
	"fmt"
	"os"
	"os/exec"

	"log"
)

func SetUpBoilerplate(dir, module string) error {
	log.Println("Setting up boilerplate...")
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Fatal("provided output directory does not exists")
		return err
	}

	log.Println("Generating go.mod...")

	// create go mod init
	cmd := exec.Command("go", "mod", "init", module)
	cmd.Dir = dir
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run failed: %v\n", err)
		return err
	}

	log.Printf("Generated module %s\n", module)

	err = setupDirs(dir)
	if err != nil {
		log.Fatalf("setupDirs failed: %v\n", err)
		return err
	}

	return nil
}

func setupDirs(path string) error {
	log.Println("Setting up directories...")
	err := os.Mkdir(fmt.Sprintf("%s/database", path), 1)
	if err != nil {
		log.Fatalf("os.Mkdir database failed: %v\n", err)
		return err
	}

	log.Println("Created /database directory")

	return nil
}
