package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool }
func genFn0(p0 string, p1 uint16, p2 int16) uint8 {
    var v6 float64 = 3.14159
    if false {
        fmt.Println(p2, p0)
    } else {
        for fi7 := 0; fi7 < 17; fi7++ {
            for fi8 := 0; fi8 < 14; fi8++ {
                p1 -= 717
                v9 := GenS1{F0: 0}
                var v10 int16 = -557
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = fi7
        }
    }
    p2 += p2
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 127
}

func main() {
    var v1 uint16 = 886
    v2 := []uint32{775}
    v3 := make(map[string]int)
    v4 := make(map[uint32]int)
    v5 := make(map[int]int)
    v5[128] = -819
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch -1 {
    case -899:
        var v11 int = 0
        _ = v11
    case -367:
        for fi12 := 0; fi12 < 17; fi12++ {
            fmt.Println(v1)
            if false {
                v2 = append(v2, 799)
                for fi13 := 0; fi13 < 16; fi13++ {
                    var v14 int32 = -1
                    var v15 uint8 = uint8(v1)
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                v16 := GenS1{F0: 42}
                _ = v16
            }
            var v17 int32 = -1
            _ = fi12
            _ = v17
        }
    default:
        var v18 int16 = 1
        _ = v18
    }
    fmt.Println(v1)
    v1 = v1
    for fi19 := 0; fi19 < 8; fi19++ {
        var v20 int16 = -487
        _ = fi19
        _ = v20
    }
    switch -469 {
    case 11:
        v1 -= (v1 & 492)
    case 32767:
        v21 := v5[32767]
        _ = v21
    default:
        v22 := GenS2{F0: true}
        _ = v22
    }
    _ = len(v5)
    var v23 int = 0
    for fi24 := 0; fi24 < 6; fi24++ {
        v23 = v23
        if false {
            var v25 uint32 = 0
            var v26 int16 = 693
            _ = v25
            _ = v26
        }
        for fi27 := 0; fi27 < 15; fi27++ {
            var v28 uint8 = 42
            if true {
                _ = len(v4)
            } else {
                v23 -= v23
            }
            _ = fi27
            _ = v28
        }
        _ = fi24
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 += v1
    var v29 float64 = (1e10 / -2.25)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
