package random

import "time"

// Implementation of a basic randomized gossip protocol.
// At each round, each node in the network will randomly a neighbour
// and send information to it.

type MessageId struct {
	NodeId         uint
	SequenceNumber uint
}

type Message struct {
	Id    MessageId
	MyId  uint
	Round uint
	Data  interface{}
}

type GossipNode struct {
	NodeId uint

	GossipNodeState

	tickInterval time.Duration

	doneC chan struct{}
	stopC chan struct{}
}

type GossipNodeState struct {
	RemotePeers         []RemotePeer
	PropagetionProgress map[uint]map[uint]bool
	Messages            map[uint]Message
}

type RemotePeer struct {
	host string
	port uint16
}

type GossipOptions struct {
	TickInterval time.Duration
}

func NewGossipNode(host string, port uint16, peers []RemotePeer, opts GossipOptions) *GossipNode {
	node := &GossipNode{

		GossipNodeState: GossipNodeState{
			RemotePeers: peers,
		},

		tickInterval: opts.TickInterval,

		doneC: make(chan struct{}),
		stopC: make(chan struct{}),
	}

	return node
}

func (n *GossipNode) ProcessMessage(msg interface{}) {

}

func (n *GossipNode) Start() {

}

func (n *GossipNode) GracefulStop() {

}

func (n *GossipNode) Crash() {

}

func (n *GossipNode) run() {

}

func (n *GossipNode) Join() {

}

func (n *GossipNode) Leave() {

}

func (n *GossipNode) round() {

}

func (n *GossipNode) Propose(msg interface{}) {

}
