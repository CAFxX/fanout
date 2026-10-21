package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 uint64) int16 {
    v5 := GenS1{F0: 0}
    p0 += 4294967295
    _ = p0
    _ = v5
    return 1
}

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = 42
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]int)
    v2 = v2
    for fi6 := 0; fi6 < 4; fi6++ {
        v7 := []int{fi6, fi6}
        var v8 int8 = -52
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 16; fi9++ {
        for fi10 := 0; fi10 < 3; fi10++ {
            var v11 bool = (!false)
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    fmt.Println(v2)
    if (true || (true && ((!false) || false))) {
        v12 := GenS1{F0: 1}
        v2 = v2
        delete(v3, 1)
        _ = v12
    }
    _ = len(v3)
    v2 += v2
    var v13 uint16 = 430
    if (4294967295 >= 42) {
        for fi14 := 0; fi14 < 2; fi14++ {
            for fi15 := 0; fi15 < 8; fi15++ {
                v16 := []int{}
                _ = fi15
                _ = v16
            }
            v17 := []int{(-(fi14)), 25, (fi14 - (fi14 | (fi14 * (fi14 / (((fi14 >> ((42) & 63))) | 1))))), fi14}
            var v18 int = (fi14 + (-(fi14)))
            _ = fi14
            _ = v17
            _ = v18
        }
    }
    v19 := GenS1{F0: -1}
    v20 := []int{}
    for fi21 := 0; fi21 < 8; fi21++ {
        v13 += 0
        _ = fi21
    }
    v13 -= (^(v13))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
