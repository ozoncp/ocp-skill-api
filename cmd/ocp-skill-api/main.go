package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"context"
	"google.golang.org/grpc"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	desc "github.com/ozoncp/ocp-skill-api/pkg/ocp-skill-api"
	"github.com/opentracing/opentracing-go"
	"github.com/ozoncp/ocp-skill-api/internal/producer"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/ozoncp/ocp-skill-api/internal/repo"
	"github.com/uber/jaeger-client-go"
	api "github.com/ozoncp/ocp-skill-api/internal/api"
	"github.com/jmoiron/sqlx"
	_ "github.com/jackc/pgx/v4/stdlib"
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

	connectString := "host=db port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"

	db, err := sqlx.Connect("pgx", connectString)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	kafka, err := producer.NewProducer([]string{"kafka:9092"}, "skills")
	if err != nil {
		log.Fatalf("failed to connect kafka: %v", err)
	}

	s := grpc.NewServer()
	desc.RegisterOcpSkillApiServer(s, api.NewSkillAPI(repo.NewRepo(db), kafka))

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

func initJaeger(address string) (io.Closer, error) {
	cfgMetrics := &jaegercfg.Configuration{
		ServiceName: "ocp-skill-api",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}
	tracer, closer, err := cfgMetrics.NewTracer(
		jaegercfg.Logger(jaeger.StdLogger),
	)
	if err != nil {
		return nil, err
	}
	opentracing.SetGlobalTracer(tracer)

	return closer, nil
}

func createMetricsServer() *http.Server {

	mux := http.DefaultServeMux
	mux.Handle("/metrics", promhttp.Handler())

	metricsServer := &http.Server{
		Addr:    ":9100",
		Handler: mux,
	}

	return metricsServer
}



func main() {
	fmt.Println("This is Skill API")
	metricsServer := createMetricsServer()

	go func() {
		if err := metricsServer.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	jaegerRunner, err := initJaeger("0.0.0.0:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer jaegerRunner.Close()
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