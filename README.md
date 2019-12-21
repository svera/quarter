# Quarter

A minimalistic arcade-oriented toolkit on top of [Pixel](https://github.com/faiface/pixel).

## Why another game framework / engine?

Because I love Go, I love Pixel and I love classic arcade games and I wanted to combine all these 3 passions in a simple package 
that allowed me and others to have fun developing arcade games.

## Objectives

* Simple, so it is easy to learn and use, providing a coherent API.
* Focused, as it is oriented to build 2D arcade games only.
* Modular, just include those packages you're interested in.
* Well tested, providing a rock-solid foundation for your games.
* Fast, performance should be the best it can, provided that this goal does not collide with the previous ones.

## Features

* Animated sprites with `animation` subpackage.
* Collision detection and handling using AABB in `collision` subpackage.
* Level loading with layering, grid and custom properties support in `level` subpackage.
* Simple game status management with `scene` subpackage.
* Text effects with `textfx` subpackage.
