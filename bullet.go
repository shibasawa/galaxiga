package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const BULLET_SPEED = 0.15

func CreateBullet(renderer *sdl.Renderer) *Element {
	bullet := &Element{}

	sr := CreateSpriteRenderer(
		bullet,
		renderer,
		"sprites/laser-bolts.png",
		Rect{p: Vector{x: 0.0, y: 16.0}, s: Size{w: 16.0, h: 16.0}})
	bullet.addComponent(sr)
	bullet.addComponent(CreateBulletMover(bullet))

	bullet.tag = "bullet"

	return bullet
}
