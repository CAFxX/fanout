package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int64 }
type GenS2 struct { F0 uint32; F1 uint; F2 int64 }
func genFn0(p0 uint64, p1 bool, p2 uint8) int8 {
    v5 := GenS1{F0: 9223372036854775807, F1: -986, F2: 1}
    p2 += p2
    fmt.Println(p0, p2)
    var v6 uint32 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 1
}
func genFn1() uint64 {
    var v7 uint16 = 373
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 int32 = 429
        _ = fi8
        _ = v9
    }
    fmt.Println(v7)
    _ = v7
    return 4294967295
}

func main() {
    var v1 uint32 = 4294967295
    var v2 uint16 = 255
    v3 := []int{127, -363}
    var v4 int32 = -190
    var v10 string = "\n\t"
    fmt.Println(v2, v2, v1)
    switch -1 {
    case -1:
        fmt.Println(v1, v10)
    default:
        var v11 uint32 = (v1 * v1)
        _ = v11
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        v3 = append(v3, (fi12 | 32768))
        for fi13 := 0; fi13 < 7; fi13++ {
            fmt.Println(v10)
            if false {
                v14 := GenS2{F0: v1, F1: 1, F2: -264}
                _ = v14
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            fmt.Println(v4, v1, v2)
            _ = fi13
        }
        _ = fi12
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        for fi16 := 0; fi16 < 14; fi16++ {
            fmt.Println(v2)
            _ = fi16
        }
        v4 -= v4
        _ = fi15
    }
    switch 2147483647 {
    case 260:
        fmt.Println(v2, v4)
    }
    v3 = append(v3, 128)
    if len(v3) > 0 {
        v3[0] = 4294967295
    }
    v17 := GenS2{F0: v1, F1: 42, F2: -1}
    _ = len(v3)
    v2 -= v2
    for fi18 := 0; fi18 < 9; fi18++ {
        for fi19 := 0; fi19 < 4; fi19++ {
            var v20 uint = 0
            _ = fi19
            _ = v20
        }
        v3 = append(v3, (fi18 ^ 1))
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
