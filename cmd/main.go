package main

import (
	standardserver "github.com/edwintrumpet/servers-benchmark/api/standard"
)

func main() {
	standardserver.Start(":3000")
}
