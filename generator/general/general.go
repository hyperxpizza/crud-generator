package general

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"log"

	"github.com/hyperxpizza/crud-generator/templates"
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
		if dir == "./generated" {
			err := os.Mkdir("generated", 0755)
			if err != nil {
				log.Fatalf("creating generated directory failed: %v\n", err)
				return err
			}
		} else {
			log.Fatal("provided output directory does not exists")
			return err
		}
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

func setupDockerFiles(dir string) error {
	log.Println("Setting up Dockerfile...")
	dockerfile, err := os.Create(fmt.Sprintf("%s/Dockerfile", dir))
	if err != nil {
		log.Printf("os.Create Dockerfile failed: %v\n", err)
		return err
	}

	err = templates.DockerfileTemplate.Execute(dockerfile, struct {
		Timestamp time.Time
	}{
		Timestamp: time.Now(),
	})
	if err != nil {
		log.Fatalf("dockerfileTemplate.Execute failed: %v\n", err)
		return err
	}
	defer dockerfile.Close()
	log.Println("Created Dockerfile")

	log.Println("Setting up docker-compose...")
	dockerCompose, err := os.Create(fmt.Sprintf("%s/docker-compose.yml", dir))
	if err != nil {
		log.Printf("os.Create docker-compose failed: %v\n", err)
		return err
	}

	err = templates.DockerComposeTemplate.Execute(dockerCompose, struct {
		Timestamp time.Time
	}{
		Timestamp: time.Now(),
	})
	if err != nil {
		log.Fatalf("dockerComposeTemplate.Execute failed: %v\n", err)
		return err
	}

	defer dockerCompose.Close()
	log.Println("Created docker-compose.yml")

	return nil
}

func SetUpMainTemplate() {

}
