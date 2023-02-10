package main

import "fmt"

type peer struct {
	id          uint8
	primary     string
	information string
	location    string
}

type load interface {
	reload(peer)
}

func (receiver *peer) reload(peer peer) {
	receiver.id = peer.id
	receiver.primary = peer.primary
	receiver.information = peer.information
	receiver.location = peer.location
}

func main() {
	m := peer{
		id:          1,
		primary:     "abc",
		information: "he",
		location:    "x",
	}
	fmt.Println(m)
}
