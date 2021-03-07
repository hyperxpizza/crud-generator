package general

import (
	"fmt"
	"os"
	"os/exec"

	"log"
)

var packages = []string{
	"github.com/lib/pq",
	"github.com/gin-gonic/gin",
	"github.com/joho/godotenv",
}

//SetUpBoilerplate generates boilerplate for the whole project
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
	log.Printf("Installing packages...")
	//install all packages
	for _, v := range packages {
		log.Printf("Getting %s\n", v)
		cmd := exec.Command("go", "get", v)
		cmd.Dir = dir
		err := cmd.Run()
		if err != nil {
			log.Fatalf("installation of package: %s failed: %v\n", v, err)
			return err
		}
	}

	err = setupDirs(dir)
	if err != nil {
		log.Fatalf("setupDirs failed: %v\n", err)
		return err
	}

	return nil
}

func setupDirs(path string) error {
	log.Println("Setting up directories...")
	err := os.Mkdir(fmt.Sprintf("%s/database", path), 0755)
	if err != nil {
		log.Fatalf("os.Mkdir database failed: %v\n", err)
		return err
	}

	log.Println("Created /database directory")

	err = os.Mkdir(fmt.Sprintf("%s/handlers", path), 0755)
	if err != nil {
		log.Fatalf("os.Mkdir handlers failed: %v\n", err)
		return err
	}

	log.Println("Created /handlers directory")

	return nil
}

func SetUpMainTemplate() {

}
