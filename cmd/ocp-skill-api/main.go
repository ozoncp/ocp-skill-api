package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"context"
	"google.golang.org/grpc"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	desc "github.com/ozoncp/ocp-skill-api/pkg/ocp-skill-api"
	api "github.com/ozoncp/ocp-skill-api/internal/api"
)

const (
	grpcPort = ":82"
	grpcServerEndpoint = "localhost:82"
)

func run() error {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	desc.RegisterOcpSkillApiServer(s, api.NewSkillAPI())

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func runJSON() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := desc.RegisterOcpSkillApiHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		panic(err)
	}

	err = http.ListenAndServe(":8081", mux)
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("This is Skill API")
	go runJSON()
	if err := run(); err != nil {
		log.Fatal(err)
	}
	output, _ := ReadConfig([]string{"conf.json"})
	fmt.Println(output)
}

func ReadConfig(paths []string) ([]string, error) {
	if err := run(); err != nil {
		log.Fatal(err)
	}
	if len(paths) < 1 {
		return nil, errors.New("empty config list")
	}

	readFile := func(path string) (string, error) {
		file, error := os.Open(path)

		if error != nil {
			return "", error
		}
		defer func() {
			if error := file.Close(); error != nil {
				log.Println("can't close file")
			}
		}()

		data := new(bytes.Buffer)

		if _, error = data.ReadFrom(file); error != nil {
			return "", error
		}

		return data.String(), nil
	}

	output := make([]string, 0)

	for _, path := range paths {
		data, error := readFile(path)

		if error != nil {
			log.Println(fmt.Sprintf("problems with %v, skip", path))
		} else {
			output = append(output, data)
		}
	}

	return output, nil
}