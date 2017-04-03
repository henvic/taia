package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

// NetworkSize is the amount of nodes used
const NetworkSize = 200

// Edge linking two nodes
type Edge struct {
	Node1 int
	Node2 int
}

// Node of the network
type Node struct {
	ID    int
	Level int
	Edges []Edge
}

// FreeScale network return an array of edges
func FreeScale() []Edge {
	var nodes = []Node{}
	var edges = []Edge{}

	for i := 0; i < NetworkSize; i++ {
		nodes = append(nodes, Node{
			ID:    i,
			Level: 1,
		})

		var x = rand.Intn(i + 1)

		if float32(nodes[x].Level)*rand.Float32() > 0.5 {
			edges = append(edges, Edge{
				Node1: i,
				Node2: x,
			})
		}

		var y = rand.Intn(i + 1)

		if float32(nodes[y].Level)*rand.Float32() > 0.5 {
			edges = append(edges, Edge{
				Node1: i,
				Node2: y,
			})
		}
	}

	return edges
}

// RandomNetwork returns an array of edges
func RandomNetwork() []Edge {
	var edges = []Edge{}

	for i := 0; i < NetworkSize; i++ {
		edges = append(edges, Edge{
			Node1: rand.Intn(NetworkSize),
			Node2: rand.Intn(NetworkSize),
		})
	}

	return edges
}

// Print network
func Print(w io.Writer, edges []Edge) {
	fmt.Fprintf(w, "Source;Target;Type\n")

	for i := 0; i < len(edges); i++ {
		fmt.Fprintf(w, "%v;%v;undirected\n", edges[i].Node1, edges[i].Node2)
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	switch args := strings.Join(os.Args, " "); {
	case strings.Contains(args, "freescale"):
		Print(os.Stdout, FreeScale())
	case strings.Contains(args, "random"):
		Print(os.Stdout, RandomNetwork())
	default:
		println(`Use "freescale" or "random" as arguments`)
		os.Exit(1)
	}
}
