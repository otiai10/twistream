package twistream

import "regexp"
import "encoding/json"

type parser struct {
	streamProxy chan Status
	pool        []byte
	triggered   bool
	trigger     *regexp.Regexp
	ready       bool
}

// Fulfil io.Writer implementaion.
func (p *parser) Write(message []byte) (n int, err error) {
	if !p.isTriggered(message) {
		return
	}
	p.pool = append(p.pool, message...)
	if status, ok := p.buildStatus(); ok {
		p.streamProxy <- status
		return p.reset(n, err)
	}
	return
}

// Accept only the next message matches to the trigger.
// If it's already triggered, turn off the flag.
func (p *parser) isTriggered(message []byte) bool {
	if p.triggered {
		p.triggered = false
		return true
	}
	if p.trigger.Match(message) {
		p.triggered = true
	}
	return false
}

// Try to build status from pooled bytes array.
// If it fails, this is uncompleted chunked status.
func (p *parser) buildStatus() (status Status, ok bool) {
	if err := json.Unmarshal(p.pool, &status); err != nil {
		// It failed, because it's still chunked.
		return status, false
	}
	// Ignore the first JSON entry.
	if !p.ready {
		p.ready = true
		p.pool = []byte{}
		return status, false
	}
	return status, true
}

// Reset pooled bytes array, and of course turn off the flag.
func (p *parser) reset(n int, err error) (int, error) {
	p.pool = []byte{}
	p.triggered = false
	return n, err
}
