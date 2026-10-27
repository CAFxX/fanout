package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 int64; F1 uint32; F2 int32 }
func genFn0(p0 int64, p1 int32, p2 uint8) bool {
    p2 -= p2
    v7 := []int{478, 128, -1, 32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return true
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int16 = 32767
    var v3 int32 = 165
    var v4 int = 2147483647
    var v5 int = v4
    v6 := make(map[string]uint32)
    if (1 >= 1) {
        v8 := v6[""]
        v8 = (v8 << ((1) & 63))
        _ = v8
    } else {
        v3 -= (v3 | v3)
        var v9 bool = (!((0.0 + ((((1e10 - -2.25) - ((0.1 - 3.14159) + 0.0)) - 1e10) + 0.1)) >= -2.25))
        _ = v9
    }
    v10 := GenS1{F0: true, F1: 32767}
    var v11 string = "\n\t"
    var v12 int8 = 0
    fmt.Println(v2, v11, v12)
    var v13 int64 = 485
    var v14 float64 = (0.1 + 1e10)
    var v15 uint = 42
    fmt.Println(v15, v12, v5)
    for fi16 := 0; fi16 < 3; fi16++ {
        v17 := []int{}
        delete(v6, v11)
        _ = fi16
        _ = v17
    }
    fmt.Println(v2, v5)
    switch (v5 + 32767) {
    case 128:
        for fi18 := 0; fi18 < 18; fi18++ {
            v19 := GenS2{F0: (v13 | v13), F1: 42, F2: (v3 / ((v3) | 1))}
            v20 := GenS1{F0: ((v3 <= 783) && false), F1: 42}
            _ = fi18
            _ = v19
            _ = v20
        }
    case -2147483648:
        var v21 uint64 = 18446744073709551615
        _ = v21
    case -216:
        if (v13 > (v13 * (-(v13)))) {
            for fi22 := 0; fi22 < 15; fi22++ {
                fmt.Println(v15, v2, v3)
                v23 := GenS2{F0: -1, F1: 65535, F2: (v3 | 0)}
                _ = fi22
                _ = v23
            }
            _ = len(v6)
        }
    default:
        fmt.Println(v4)
    }
    v3 = 618
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
