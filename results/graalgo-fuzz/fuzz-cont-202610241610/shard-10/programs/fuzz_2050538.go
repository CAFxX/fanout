package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int }
func genFn0(p0 string, p1 int8, p2 int16) int32 {
    fmt.Println(p2, p0)
    v6 := GenS1{F0: -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -2147483648
}

func main() {
    var v1 int = 1
    var v2 uint32 = 126
    var v3 uint32 = (^(v2))
    v4 := make(map[string]int)
    v5 := make(map[int]uint32)
    var v7 int = (v1 << ((471) & 63))
    v7 += (v7 - (v1 & v1))
    v8 := []int{}
    v3 = (v2 ^ v2)
    fmt.Println(v2, v2, v3)
    var v9 int32 = -67
    if false {
        var v10 string = "\n\t"
        var v11 bool = true
        if v11 {
            _ = len(v8)
        }
        _ = v10
        _ = v11
    }
    var v12 uint8 = 128
    v3 -= 4294967295
    var v13 float64 = ((1e10 - 1.5) + (float64(float32(1.5)) - 0.0))
    v14 := GenS2{F0: (-(v1))}
    v13 = v13
    for fi15 := 0; fi15 < 9; fi15++ {
        v2 -= v3
        delete(v5, v1)
        _ = fi15
    }
    v1 = v7
    v13 -= (v13 + v13)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
