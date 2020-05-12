package main

// BulbBrightness - Shows how much life is left
type BulbBrightness int32

// oneUseLamp - A Lamp where you can't touch the bulb.
type oneUseLamp struct {
	safeBulb BulbBrightness
}

// Lamp - what else could you want from it?
type Lamp interface {
	GetOriginalBrightness() BulbBrightness
}

// GetOriginalBrightness - The bulb multiple threads can't break
func (lamp oneUseLamp) GetOriginalBrightness() BulbBrightness {
	return lamp.safeBulb
}

// NewOneUseLamp - the only way to get oneUseLamp
func NewOneUseLamp(brightness BulbBrightness) Lamp {
	return oneUseLamp{brightness}
}

//  What type of object does this refer to ?
