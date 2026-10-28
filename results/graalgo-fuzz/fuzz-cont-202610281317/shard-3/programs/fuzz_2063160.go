package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 bool; F2 uint }
func genFn0(p0 uint64) uint32 {
    fmt.Println(p0)
    var v5 uint16 = uint16(p0)
    var v6 bool = true
    if (!(687 <= 0)) {
        v7 := GenS1{F0: "\n\t"}
        v8 := GenS2{F0: 470, F1: v6, F2: 255}
        _ = v7
        _ = v8
    } else {
        v9 := []int{4294967295, 862, 4294967295}
        _ = v9
    }
    v5 = 42
    _ = p0
    _ = v5
    _ = v6
    return 99
}

func main() {
    var v1 float32 = float32(1e10)
    var v2 uint64 = 2
    v3 := make(map[uint32]int)
    var v4 float32 = -2.25
    var v10 uint8 = 14
    v11 := []int{2147483647, 2147483647, 727, 32767}
    v3[42] = 9223372036854775807
    v10 = (^(v10))
    delete(v3, 2147483648)
    for fi12 := 0; fi12 < 5; fi12++ {
        v2 -= (v2 << ((797) & 63))
        for fi13 := 0; fi13 < 7; fi13++ {
            fmt.Println(v2, fi12, v2)
            v2 = uint64(v10)
            v1 += v1
            _ = fi13
        }
        _ = fi12
    }
    v14 := v3[2147483647]
    v10 = v10
    if false {
        var v15 uint16 = 0
        _ = v15
    }
    v4 -= (v1 - ((v4 - (float32(0.0) * v4)) / 1e10))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
