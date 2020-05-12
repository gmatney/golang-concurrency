package main

import (
	"sync/atomic"
)

// BulbBrightness - Shows how much life is left
type BulbBrightness int32

//LightSocket - Where the lightbulb goes
type LightSocket *int32

// LockedRoom - A safe place for all your concurrent light bulb changing needs.
type LockedRoom struct{}

// ReplaceLightBulb - Safely replace a light with a fresh 100 life bulb.
func (lr LockedRoom) ReplaceLightBulb(socket LightSocket) {
	atomic.StoreInt32(socket, 100)
}

// GetLightBrightness - Test how bright a light is
func (lr LockedRoom) GetLightBrightness(socket *LightSocket) BulbBrightness {
	return BulbBrightness(atomic.LoadInt32(*socket))
}
