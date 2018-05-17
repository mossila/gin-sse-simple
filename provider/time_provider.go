package provider

import (
	"time"

	"github.com/grafov/bcast"
)

func main() {
	group := bcast.NewGroup()
	go group.Broadcast(0)
	go TimeProvider(group, 100*time.Millisecond)
	go newReciever(group)
	newReciever(group)
}
func newReciever(group *bcast.Group) {
	member := group.Join()
	for {
		println("message", member.Recv().(string))
	}
}

//TimeProvider provide time stamp every configured time
func TimeProvider(group *bcast.Group, d time.Duration) {
	defer group.Close()
	for x := range time.Tick(d) {
		group.Send(x.String())
	}
}
