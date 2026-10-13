package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint16 }
type GenS2 struct { F0 uint }
func genFn0(p0 bool, p1 int32, p2 uint16) int16 {
    p1 = (p1 << ((211) & 63))
    var v5 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -44
}
func genFn1(p0 uint16, p1 uint32) uint64 {
    switch -2147483648 {
    case -1:
        if true {
            for fi6 := 0; fi6 < 20; fi6++ {
                var v7 float32 = (-2.25 - 0.0)
                _ = fi6
                _ = v7
            }
            fmt.Println(p0, p0)
            v8 := GenS1{F0: uint64(p1), F1: 2147483647, F2: p0}
            _ = v8
        }
    case -189:
        var v9 int = 1
        _ = v9
    }
    fmt.Println(p0, p0)
    v10 := GenS1{F0: 4294967295, F1: 0, F2: p0}
    _ = p0
    _ = p1
    _ = v10
    return 1
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint32 = 563
    v3 := make(map[int]uint32)
    var v4 int = 32767
    for fi11 := 0; fi11 < 1; fi11++ {
        v2 -= v2
        if (false && (-2147483648 >= 1)) {
            fmt.Println(fi11, v2, v4)
            var v12 uint64 = 18446744073709551615
            _ = v12
        }
        v4 += fi11
        _ = fi11
    }
    v13 := GenS1{F0: 1, F1: 388, F2: 170}
    v2 -= (v2 / (((v2 + v2)) | 1))
    v14 := GenS1{F0: 94, F1: 392, F2: 32767}
    v15 := []int{v4, -265, v4, (v4 + 2147483647)}
    switch v4 {
    case 270:
        fmt.Println(v2, v4)
    default:
        v16 := GenS1{F0: 483, F1: 691, F2: 996}
        _ = v16
    }
    var v17 int8 = -1
    for fi18 := 0; fi18 < 4; fi18++ {
        v17 = v17
        for fi19 := 0; fi19 < 13; fi19++ {
            fmt.Println(fi19, v2, v17)
            v20 := GenS2{F0: 2}
            _ = fi19
            _ = v20
        }
        var v21 uint16 = 2
        _ = fi18
        _ = v21
    }
    v17 -= v17
    v22 := GenS1{F0: 703, F1: 761, F2: 0}
    fmt.Println(v2, v4, v4)
    fmt.Println(v17, v17)
    var v23 uint = 2
    for fi24 := 0; fi24 < 2; fi24++ {
        _ = len(v3)
        v3[(fi24 - (v4 | (v4 % (((v4 - 1)) | 1))))] = 0
        _ = fi24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
