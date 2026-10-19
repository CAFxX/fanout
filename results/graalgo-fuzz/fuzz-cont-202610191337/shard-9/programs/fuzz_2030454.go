package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 uint32 }
func genFn0(p0 int32, p1 int64, p2 int64) uint32 {
    v4 := []int{-642}
    fmt.Println(p1, p1, p2)
    for fi5 := 0; fi5 < 10; fi5++ {
        if (true || ((p0 == p0) && false)) {
            _ = len(v4)
        }
        _ = fi5
    }
    fmt.Println(p0)
    for fi6 := 0; fi6 < 9; fi6++ {
        for fi7 := 0; fi7 < 15; fi7++ {
            p0 -= p0
            p1 = (p1 >> ((1) & 63))
            var v8 uint8 = 164
            _ = fi7
            _ = v8
        }
        var v9 float64 = 0.0
        p0 -= (p0 + p0)
        _ = fi6
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 1
}

func main() {
    var v1 int8 = -58
    var v2 uint8 = 127
    var v3 uint64 = 2
    var v10 uint32 = 2147483647
    fmt.Println(v3, v3, v2)
    v1 = (v1 ^ (v1 + (v1 & (-(v1)))))
    v11 := []int{-811, -274, -168}
    v12 := GenS1{F0: 0, F1: 32768, F2: genFn0(0, 163, 0)}
    v10 = 415
    v1 = v1
    fmt.Println(v3, v3, v1)
    v1 -= v1
    v10 = (v10 << ((1) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
}
