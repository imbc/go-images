package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"os"
	"sort"
	"time"
)

func main() {
	width, height := 640, 640
	canvas := NewCanvas(image.Rect(0, 0, width, height))
	canvas.DrawRect(color.RGBA{0, 0, 0, 255}, Vector{0, 0}, Vector{float64(width), float64(height)})
	rand.Seed(time.Now().UTC().UnixNano())

	// Create and populate the slice of Nodes
	n := 250
	peers := 6
	nodes := make([]*Node, n)
	for i := 0; i < n; i++ {
		nodes[i] = NewNode(peers, canvas)
	}

	// Randomly point Nodes at each other
	// for _, node := range nodes {
	//   for _, j := range rand.Perm(n)[:peers] {
	//     node.Peers = append(node.Peers, nodes[j])
	//   }
	// }

	// Calculate nearest peers for each node
	// This is pretty ineffecient for large n
	nodesCopy := make([]*Node, n)
	copy(nodesCopy, nodes)
	log.Print("Sorting nodes...")
	for _, node := range nodes {
		// Sort the nodes by distance
		sorter := NodeSorter{nodesCopy, node}
		sort.Sort(sorter)
		node.Peers = append(node.Peers, nodesCopy[1:peers+1]...)
	}
	log.Print("Nodes sorted")

	// Draw on circles representing nodes
	for _, node := range nodes {
		canvas.DrawCircle(color.RGBA{22, 131, 201, 255}, node.Position, 5)
	}
	canvas.Blur(3, new(WeightFunctionDist))

	// Draw connections between nodes
	for _, node := range nodes {
		for _, peer := range node.Peers[:3] {
			canvas.DrawLine(color.RGBA{0, 0, 0, 10}, node.Position, peer.Position)
		}
	}

	// Start sending messages between nodes
	for i := 0; i < 1; i++ {
		nodes[i].Power = 255
		go nodes[i].Send()
	}
	time.Sleep(time.Second)

	// Write out image
	dirName := "img"
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		os.Mkdir(dirName, 0755)
		return
	}
	outFilename := dirName + "/nodes.png"
	outFile, err := os.Create(outFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	log.Print("Saving image to: ", outFilename)
	png.Encode(outFile, canvas)
}
