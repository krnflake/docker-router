package main

import (
	"github.com/fsouza/go-dockerclient"
	"log"
	"time"
)

var dockerCache = NewLRUCache(1000, 5*time.Minute)

func inspectCachedContainer(id string) (*docker.Container, error) {
	client, err := docker.NewClient(*dockerAddr)
	if err != nil {
		log.Fatal(err)
	}

	cContainer, ok := dockerCache.Get(id)
	if ok {
		return cContainer.(*docker.Container), nil
	}

	container, err := client.InspectContainer(id)
	dockerCache.Add(id, container)
	if err != nil {
		return nil, err
	}

	return container, nil
}
