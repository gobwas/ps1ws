// START OMIT
func (ch *Channel) Send(pkt Packet) {
	ch.queued++
	ch.spawnWriterMaybe()
	ch.queue <- pkt
}

func (ch *Channel) spawnWriterMaybe() {
	ch.mu.Lock()
	defer ch.mu.Unlock()

	if !ch.writing {
		ch.writing = true
		go ch.writer()
	}
}

// END OMIT

func (ch *Channel) writer() {
	var n int
	var stop bool
	for !stop {
		for i := 0; i < n; i++ {
			writePacket(<-ch.queue)
		}

		ch.mu.Lock()
		n, ch.queued = ch.queued, 0
		if n == 0 {
			ch.writing = false
			stop = true
		}
		ch.mu.Unlock()
	}
}
