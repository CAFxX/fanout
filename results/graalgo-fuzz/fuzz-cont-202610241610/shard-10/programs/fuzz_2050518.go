package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int64 }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint32 }
func genFn0(p0 string, p1 int64) string {
    p1 += -2
    for fi7 := 0; fi7 < 16; fi7++ {
        for fi8 := 0; fi8 < 13; fi8++ {
            p0 = p0
            var v9 uint16 = 42
            if (true && (127 <= 100)) {
                fmt.Println(fi7, p0, fi8)
            } else {
                p0 = "a"
            }
            p1 = p1
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    p0 = p0
    _ = p0
    _ = p1
    return ((("" + "") + p0) + p0)
}

func main() {
    var v1 int64 = 0
    var v2 bool = ("hello" < "hello")
    var v3 float32 = float32(0.0)
    var v4 int64 = -683
    v5 := make(map[int]uint32)
    var v6 int16 = 32767
    for fi10 := 0; fi10 < 9; fi10++ {
        v2 = v2
        _ = fi10
    }
    for fi11 := 0; fi11 < 18; fi11++ {
        v12 := []int{}
        _ = fi11
        _ = v12
    }
    v13 := v5[-1]
    fmt.Println(v6, v4, v6)
    v14 := v5[-837]
    v15 := GenS1{F0: 1, F1: 1, F2: (v4 + v4)}
    fmt.Println(v13)
    if v2 {
        var v16 uint = 856
        _ = v16
    } else {
        if false {
            delete(v5, -747)
            v17 := GenS2{F0: 42, F1: 1, F2: 862}
            _ = v17
        }
        v18 := GenS2{F0: 255, F1: 255, F2: (v13 | v13)}
        v19 := GenS1{F0: v14, F1: 42, F2: (-(v1))}
        _ = v18
        _ = v19
    }
    var v20 int16 = (v6 / ((v6) | 1))
    if v2 {
        v13 = 305
    }
    switch 128 {
    case 4294967295:
        v21 := GenS2{F0: 1, F1: 622, F2: (v14 * v14)}
        _ = v21
    case 787:
        fmt.Println(v2)
    default:
        for fi22 := 0; fi22 < 14; fi22++ {
            var v23 bool = v2
            var v24 uint = 18446744073709551615
            _ = fi22
            _ = v23
            _ = v24
        }
    }
    if v2 {
        v25 := GenS1{F0: v14, F1: 904, F2: (v1 + (v1 << ((507) & 63)))}
        fmt.Println(v6)
        _ = v25
    }
    v5[3] = v14
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
