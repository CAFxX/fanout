package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint }
func genFn0(p0 int32) uint64 {
    if true {
        v6 := []int{2147483647, 506, 128, 745}
        var v7 int16 = 441
        _ = v6
        _ = v7
    }
    v8 := []int{1, -280}
    var v9 uint8 = 1
    p0 = 1
    p0 += (p0 & (p0 + p0))
    _ = p0
    _ = v8
    _ = v9
    return 4294967295
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]uint32)
    var v5 int = 2147483647
    v5 += (v5 >> ((545) & 63))
    var v10 int32 = 2147483647
    var v11 int = (^(v5))
    fmt.Println(v5, v11, v5)
    for fi12 := 0; fi12 < 5; fi12++ {
        fmt.Println(v11, v5)
        var v13 uint16 = 32767
        var v14 int8 = 127
        _ = fi12
        _ = v13
        _ = v14
    }
    fmt.Println(v5, v10)
    v15 := GenS1{F0: 1000}
    if (-2.25 > 1e10) {
        var v16 int = v11
        fmt.Println(v16, v11, v5)
        v11 = (v16 + 4294967295)
        _ = v16
    } else {
        for fi17 := 0; fi17 < 18; fi17++ {
            v11 += 355
            if true {
                v1[(v11 & -1)] = true
            } else {
                v5 = fi17
                v11 -= 297
            }
            _ = fi17
        }
        v11 -= v5
        var v18 int64 = 4
        _ = v18
    }
    for fi19 := 0; fi19 < 15; fi19++ {
        if (true && (false && ((fi19 <= v11) || (4294967295 < 334)))) {
            fmt.Println(v10)
            v20 := GenS2{F0: 4294967295}
            _ = v20
        }
        v21 := []int{}
        _ = fi19
        _ = v21
    }
    v22 := GenS3{F0: 495}
    fmt.Println(v11, v11)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
