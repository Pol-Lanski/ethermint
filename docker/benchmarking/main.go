package main

import (
	"log"
	"os"

	"github.com/ChainSafe/ethermint/docker/benchmarking/benchmark"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Version = "1.0.0"
	app.Name = "benchmark"
	app.Usage = "Benchmarking suite for ethermint"
	app.Commands = []cli.Command{
		benchmark.SendTx,
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}