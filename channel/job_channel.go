package channel

import "sync"

// Packet is a single packet, a wrapper around a Job for transport
type Packet struct {
	Data []byte
}

// PacketChan provides an async method for passing multiple ordered packets between goroutines
type PacketChan struct {
	mu   sync.Mutex
	c    chan *Packet
	C    chan<- *Packet
	err  error
	done chan struct{}
}

// Receive provides the channel from which to read packets.  It always
// returns the same channel.
func (p *PacketChan) Receive() <-chan *Packet { return p.c }

// Send sends a single packet on the channel to the receiver.
func (p *PacketChan) Send(pkt *Packet) { p.c <- pkt }

// Close closes the sending channel and sets the PacketChan's error based
// in its input.
func (p *PacketChan) Close(err error) {
	p.mu.Lock()
	p.err = err
	p.mu.Unlock()
	close(p.c)
	close(p.done)
}

// Done returns a channel that is closed when this packet channel is complete.
func (p *PacketChan) Done() <-chan struct{} {
	return p.done
}

// NewPacketChan returns a new PacketChan channel for passing packets around.
func NewPacketChan(buffer int) *PacketChan {
	pc := &PacketChan{
		c:    make(chan *Packet, buffer),
		done: make(chan struct{}),
	}
	pc.C = pc.c
	return pc
}

// Discard discards all remaining packets on the receiving end.  If you stop
// using the channel before reading all packets, you must call this function.
// It's a good idea to defer this regardless.
func (p *PacketChan) Discard() {
	go func() {
		discarded := 0
		for _ = range p.c {
			discarded++
		}
	}()
}

// Err gets the current error for the channel, if any exists.  This may be
// called during Next(), but if an error occurs it may only be set after Next()
// returns false the first time.
func (p *PacketChan) Err() error {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.err
}
