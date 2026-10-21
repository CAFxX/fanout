package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 string, p1 uint8, p2 uint64) uint {
    var v6 uint16 = 42
    v7 := GenS1{F0: 2147483647}
    v8 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 714
}

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{0, 309}
    var v3 uint32 = 1
    var v4 uint16 = 564
    v5 := make(map[uint32]uint32)
    if false {
        v3 = (^(v3))
        v9 := GenS1{F0: 1}
        v3 = (v3 << ((2) & 63))
        _ = v9
    } else {
        v4 -= v4
    }
    v4 += v4
    v4 += 391
    v4 += (^(v4))
    v10 := v5[v3]
    _ = len(v2)
    v11 := v1[-9223372036854775808]
    if len(v2) > 0 {
        v2[0] = v3
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        fmt.Println(v3, v11, v3)
        var v13 int16 = 0
        _ = fi12
        _ = v13
    }
    v5[v3] = (v10 & v3)
    fmt.Println(v3, v4)
    var v14 uint32 = (v3 / ((2) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
