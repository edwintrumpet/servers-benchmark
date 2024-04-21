package api_test

import (
	"net/http"
	"testing"
	"time"

	echoserver "github.com/edwintrumpet/servers-benchmark/api/echo"
	ginserver "github.com/edwintrumpet/servers-benchmark/api/gin"
	muxserver "github.com/edwintrumpet/servers-benchmark/api/mux"
	standardserver "github.com/edwintrumpet/servers-benchmark/api/standard"
)

func init() {
	go standardserver.Start(":3000")
	go muxserver.Start(":3001")
	go ginserver.Start(":3002")
	go echoserver.Start(":3003")

	time.Sleep(1 * time.Second)
}

func BenchmarkMuxParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			http.Get("http://localhost:3001/ping")
		}
	})
}

func BenchmarkStandarParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			http.Get("http://localhost:3000/ping")
		}
	})
}

func BenchmarkGinParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			http.Get("http://localhost:3002/ping")
		}
	})
}

func BenchmarkEchoParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			http.Get("http://localhost:3003/ping")
		}
	})
}
