package main

import (
	"fmt"
	"math"
)

type Vec2 struct {
	X, Y float32
}

// Add 加
func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{
		v.X + other.X,
		v.Y + other.Y,
	}
}

// Sub 减
func (v Vec2) Sub(other Vec2) Vec2 {
	return Vec2{
		v.X - other.X,
		v.Y - other.Y,
	}
}

// Scale 乘 缩放或者叫矢量乘法，是对矢量的每个分量乘上缩放比，Scale() 方法传入一个参数同时乘两个分量，表示这个缩放是一个等比缩放
func (v Vec2) Scale(s float32) Vec2 {
	return Vec2{v.X * s, v.Y * s}
}

// DistanceTo 距离 计算两个矢量的距离，math.Sqrt() 是开方函数，参数是 float64，在使用时需要转换，返回值也是 float64，需要转换回 float32
func (v Vec2) DistanceTo(other Vec2) float32 {
	dx := v.X - other.X
	dy := v.Y - other.Y
	return float32(math.Sqrt(float64(dx*dx + dy*dy)))
}

// Normalize 矢量单位化
func (v Vec2) Normalize() Vec2 {
	mag := v.X*v.X + v.Y*v.Y
	if mag > 0 {
		oneOverMag := 1 / float32(math.Sqrt(float64(mag)))
		return Vec2{v.X * oneOverMag, v.Y * oneOverMag}
	}
	return Vec2{0, 0}
}
func main() {
	v1 := Vec2{1, 2}
	v2 := Vec2{3, 4}
	v3 := v1.Add(v2)
	v4 := v1.Sub(v2)
	v5 := v1.Scale(2)
	v6 := v1.DistanceTo(v2)
	v7 := v1.Normalize()
	fmt.Println(v3)
	fmt.Println(v4)
	fmt.Println(v5)
	fmt.Println(v6)
	fmt.Println("Cos X = ", v7.X, "Sin X = ", v7.Y)
}
