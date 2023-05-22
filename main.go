package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/user"
	"time"
)

func main() {
	fmt.Println("Acceptance tests repo for goreleaser")

	// for gitlab
	// export GITLAB_HOME=$HOME/gitlab
	val, present := os.LookupEnv("GITLAB_HOME")
	if present {
		log.Printf("GITLAB_HOME was already set to: %s", val)
	}

	if !present {
		user, err := user.Current()
		if err != nil {
			log.Fatalf("failed to get current user")
		}
		homeDir := user.HomeDir
		gitlabHome := fmt.Sprintf("%s/gitlab", homeDir)
		os.Setenv("GITLAB_HOME", gitlabHome)
		log.Printf("GITLAB_HOME set to: %s", gitlabHome)

		defer os.Unsetenv("GITLAB_HOME")
	}

	// docker compose -f docker-compose-gitlab.yml up -d

	// wait until localhost return 200 (~10-13 min)
	expectedStatusCode := 200
	retrievedStatusCode := -1

	for expectedStatusCode != retrievedStatusCode {
		resp, err := http.Get("http://localhost/users/sign_in")
		if err != nil {
			log.Fatalf("failed to fetch '/users/sign_in' path: %s", err)
		}

		retrievedStatusCode = resp.StatusCode
		log.Printf("for '/users/sign_in' got statusCode: %v", retrievedStatusCode)
		time.Sleep(10 * time.Second)
	}

	log.Printf("for '/users/sign_in' got expected statusCode: %v", expectedStatusCode)
}
