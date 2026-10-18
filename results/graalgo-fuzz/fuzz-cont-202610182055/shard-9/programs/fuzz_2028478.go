package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 int64 }
func genFn0(p0 int16, p1 uint8, p2 bool) string {
    var v6 float32 = -2.25
    p2 = false
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return "foo bar"
}

func main() {
    var v1 int = -1
    var v2 int64 = -1
    var v3 int = (-(v1))
    v4 := make(map[int]uint32)
    var v5 uint64 = 4294967295
    v7 := v4[(v3 - 635)]
    var v8 int16 = -32768
    if true {
        v1 = v3
        v9 := GenS1{F0: 1, F1: 954, F2: v2}
        _ = v9
    }
    fmt.Println(v1, v8, v7)
    var v10 uint = 2
    v11 := []int{(v3 % ((v1) | 1)), (v3 & v3), 32768, v3}
    if len(v11) > 0 {
        v11[0] = (v1 * (v3 << ((v10) & 63)))
    }
    var v12 float32 = (float32(-2.25) * float32((float64(((3.14159 + (float32((1.5 - 1e10)) + ((0.1 / float32(float64(float32((3.14159 / ((3.14159 - (1e10 / (1.5 + -2.25))) * 1.5)))))) - 1.5))) * 3.14159)) / 1.0)))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
