package main

import (
	"image/color"
	"math/rand"
)

type Node struct {
	Position Vector
	Ch       chan *Node
	Peers    []*Node
	Canvas   *Canvas
	Power    uint8
	Group    int
}

func NewNode(peers int, canvas *Canvas) *Node {
	node := new(Node)
	node.Peers = make([]*Node, 0, peers)
	size := canvas.Bounds().Size()
	x := float64(size.X) * rand.Float64()
	y := float64(size.Y) * rand.Float64()
	node.Position = Vector{x, y}
	node.Canvas = canvas
	node.Ch = make(chan *Node)
	node.Power = 0
	go node.Listen()
	return node
}

func (n *Node) Listen() {
	// Listen for incoming connection on node's channel
	for {
		peer := <-n.Ch
		peer.Power -= 5
		n.Power = peer.Power
		n.Canvas.DrawLine(color.RGBA{255, n.Power, 0, 255}, n.Position, peer.Position)
		// Retransmit to random node
		if n.Power > 0 {
			go n.Send()
		}
	}
}

func (n *Node) Send() {
	for _, target := range n.Peers {
		if target.Power == 0 {
			target.Ch <- n
			break
		}
	}
}

type NodeSorter struct {
	data   []*Node
	target *Node
}

func (sorter NodeSorter) Len() int {
	return len(sorter.data)
}

func (sorter NodeSorter) Less(i, j int) bool {
	iDelta := sorter.data[i].Position.Sub(sorter.target.Position)
	jDelta := sorter.data[j].Position.Sub(sorter.target.Position)
	return iDelta.Length() < jDelta.Length()
}

func (sorter NodeSorter) Swap(i, j int) {
	sorter.data[i], sorter.data[j] = sorter.data[j], sorter.data[i]
}
