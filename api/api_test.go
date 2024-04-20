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

func BenchmarkStandarParallel(b *testing.B) {
	go standardserver.Start(":3000")

	time.Sleep(1 * time.Second)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			http.Get("http://localhost:3000/ping")
		}
	})
}

func BenchmarkMuxParallel(b *testing.B) {
	go muxserver.Start(":3001")

	time.Sleep(1 * time.Second)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			http.Get("http://localhost:3001/ping")
		}
	})
}

func BenchmarkGinParallel(b *testing.B) {
	go ginserver.Start(":3002")

	time.Sleep(1 * time.Second)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			http.Get("http://localhost:3002/ping")
		}
	})
}

func BenchmarkEchoParallel(b *testing.B) {
	go echoserver.Start(":3003")

	time.Sleep(1 * time.Second)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			http.Get("http://localhost:3003/ping")
		}
	})
}
