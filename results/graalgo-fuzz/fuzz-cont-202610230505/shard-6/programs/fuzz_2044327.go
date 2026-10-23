package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 uint32, p1 int32) uint8 {
    var v7 uint32 = (p0 & p0)
    if (!false) {
        var v8 bool = false
        v9 := []int{}
        _ = v8
        _ = v9
    }
    if true {
        v7 += v7
        v7 -= p0
    } else {
        var v10 float32 = 0.0
        _ = v10
    }
    v11 := GenS1{F0: 622}
    var v12 uint32 = p0
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v7
    return 77
}

func main() {
    var v1 int16 = -1
    var v2 uint = 767
    v3 := make(map[string]int)
    v4 := make(map[int]bool)
    var v5 int64 = 700
    v6 := make(map[uint32]uint32)
    v5 = (v5 * -1)
    for fi13 := 0; fi13 < 17; fi13++ {
        var v14 uint16 = uint16(v2)
        for fi15 := 0; fi15 < 6; fi15++ {
            fmt.Println(fi15, v2)
            _ = fi15
        }
        v2 += 779
        _ = fi13
        _ = v14
    }
    var v16 int8 = -128
    v17 := []int{42, 4294967295}
    if (56 > 152) {
        for fi18 := 0; fi18 < 2; fi18++ {
            for fi19 := 0; fi19 < 9; fi19++ {
                v20 := GenS1{F0: v2}
                v2 = v2
                _ = fi19
                _ = v20
            }
            _ = fi18
        }
        fmt.Println(v2, v16)
    }
    _ = len(v17)
    v16 -= (v16 * 37)
    fmt.Println(v2, v16, v2)
    v5 += v5
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
