package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 uint16 = 32767
    var v2 float64 = (((1.5 + (((3.14159 * (-2.25 * (3.14159 / 1e10))) / 1.5) - 0.0)) + 3.14159) * (3.14159 - (((1e10 - (1.5 * float64((float32(3.14159) * ((1.5 / float32((1e10 / float64(3.14159)))) / (3.14159 + 0.0)))))) / -2.25) / (0.1 / 0.1))))
    var v3 uint16 = (^(v1))
    v4 := make(map[uint32]bool)
    var v5 int32 = -2147483648
    var v6 bool = (!false)
    var v7 uint8 = 0
    var v8 uint = 65535
    var v9 int = 127
    var v10 float64 = (float64(float32(v2)) - (1e10 / v2))
    v11 := []int{0, (-(v9)), v9}
    v8 -= 2
    v12 := GenS1{F0: 9223372036854775807}
    fmt.Println(v7)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
