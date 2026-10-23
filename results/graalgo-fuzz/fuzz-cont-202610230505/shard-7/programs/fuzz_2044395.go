package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int16; F1 int64; F2 uint64 }
type GenS3 struct { F0 uint }
func genFn0(p0 uint16, p1 int8, p2 int) uint {
    if (true || true) {
        for fi7 := 0; fi7 < 2; fi7++ {
            var v8 uint = 1
            var v9 int16 = -1
            _ = fi7
            _ = v8
            _ = v9
        }
        p2 = (p2 - p2)
    }
    p1 = (p1 / ((p1) | 1))
    v10 := []int{(p2 & p2), p2}
    var v11 int32 = -550
    for fi12 := 0; fi12 < 2; fi12++ {
        _ = len(v10)
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 480
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{127}
    var v3 int16 = 42
    v4 := make(map[int]string)
    var v5 bool = ((-2.25 + (1.5 / 1.0)) == 1.5)
    v6 := make(map[uint32]string)
    if len(v2) > 0 {
        v2[0] = -832
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        switch fi13 {
        case 4294967295:
            var v14 int16 = v3
            _ = v14
        case 32768:
            switch fi13 {
            case 127:
                fmt.Println(v3, fi13, v3)
            case -270:
                v15 := GenS1{F0: fi13, F1: -21}
                _ = v15
            case 32768:
                var v16 uint64 = 42
                _ = v16
            }
        default:
            v5 = false
        }
        v17 := GenS1{F0: 127, F1: -1}
        fmt.Println(fi13, v3)
        _ = fi13
        _ = v17
    }
    v5 = v5
    v18 := GenS2{F0: (^(v3)), F1: 1, F2: 4294967295}
    if (790 != 1) {
        v3 -= v3
        v5 = (42 >= 36)
    }
    switch 42 {
    case -1:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        v3 -= (v3 << ((255) & 63))
    }
    _ = len(v2)
    v3 += 365
    v19 := GenS3{F0: 994}
    if false {
        for fi20 := 0; fi20 < 5; fi20++ {
            var v21 uint64 = 4294967295
            v21 = v21
            _ = fi20
            _ = v21
        }
        for fi22 := 0; fi22 < 8; fi22++ {
            if v5 {
                v3 = (^(v3))
            }
            v3 = (v3 ^ v3)
            _ = fi22
        }
    }
    for fi23 := 0; fi23 < 18; fi23++ {
        v5 = (!v5)
        v3 = v3
        v24 := GenS1{F0: 32768, F1: -1}
        _ = fi23
        _ = v24
    }
    v3 += (v3 | (^(v3)))
    v5 = false
    v25 := GenS3{F0: 255}
    v6[99] = ""
    var v26 int16 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
