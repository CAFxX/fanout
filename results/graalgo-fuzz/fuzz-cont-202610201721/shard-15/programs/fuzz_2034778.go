package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint, p1 int16, p2 string) int {
    p1 = (p1 >> ((255) & 63))
    v5 := GenS1{F0: 1}
    var v6 float64 = (1e10 * (3.14159 - (3.14159 / 1.0)))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 132
}

func main() {
    v1 := []uint32{2, 311, 486, 107}
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]string)
    v4 := []int{2147483647}
    var v7 string = "foo bar"
    fmt.Println(v7)
    var v8 float64 = 3.14159
    fmt.Println(v7, v7)
    var v9 uint = 18446744073709551615
    var v10 float32 = 0.1
    fmt.Println(v10, v8)
    if (197 >= 2) {
        var v11 int32 = 2147483647
        v8 = v8
        fmt.Println(v7, v8, v7)
        _ = v11
    }
    v12 := GenS2{F0: 9223372036854775807}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
