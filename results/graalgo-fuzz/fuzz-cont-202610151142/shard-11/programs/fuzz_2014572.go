package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint16 }
func genFn0(p0 bool, p1 int64) int8 {
    v7 := GenS1{F0: p1, F1: 1, F2: 32767}
    if p0 {
        var v8 int32 = 1
        _ = v8
    }
    if (8 != 221) {
        for fi9 := 0; fi9 < 15; fi9++ {
            p0 = p0
            if false {
                p1 -= p1
            }
            _ = fi9
        }
    } else {
        for fi10 := 0; fi10 < 19; fi10++ {
            for fi11 := 0; fi11 < 5; fi11++ {
                fmt.Println(p1, fi10, fi11)
                p1 += p1
                v12 := []int{(-(fi10))}
                _ = fi11
                _ = v12
            }
            v13 := GenS1{F0: (^(p1)), F1: 866, F2: 92}
            _ = fi10
            _ = v13
        }
        v14 := GenS1{F0: p1, F1: 828, F2: 255}
        _ = v14
    }
    p1 += 0
    _ = p0
    _ = p1
    _ = v7
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{}
    var v3 uint8 = 230
    v4 := make(map[int]uint32)
    var v5 uint = 0
    var v6 int32 = 850
    v15 := GenS1{F0: 1, F1: -833, F2: 390}
    var v16 int16 = 1000
    v16 -= (v16 - v16)
    var v17 int16 = (v16 + (v16 / ((-500) | 1)))
    switch "foo bar" {
    case "\n\t":
        var v18 int16 = 363
        _ = v18
    case "a":
        v19 := v1[4294967295]
        _ = v19
    case "foo bar":
        var v20 uint8 = (v3 & v3)
        _ = v20
    default:
        v5 += (^(v5))
    }
    fmt.Println(v16, v5, v6)
    if (-756 != -472) {
        v6 = v6
        var v21 int8 = 42
        _ = v21
    }
    v22 := GenS1{F0: 1, F1: -1, F2: 830}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
