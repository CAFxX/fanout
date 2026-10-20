package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 int32; F2 uint }

func main() {
    var v1 uint64 = 4294967295
    var v2 uint8 = 0
    v3 := make(map[uint32]string)
    v4 := make(map[string]uint32)
    var v5 int32 = -575
    v6 := []int{}
    v2 += (v2 * v2)
    v7 := GenS1{F0: 255}
    switch 128 {
    case 42:
        for fi8 := 0; fi8 < 13; fi8++ {
            var v9 int64 = 941
            var v10 uint8 = v2
            v6 = append(v6, (-(fi8)))
            _ = fi8
            _ = v10
            _ = v9
        }
    case 0:
        var v11 uint32 = 38
        _ = v11
    case -1:
        v12 := GenS1{F0: 255}
        _ = v12
    default:
        v1 = 2
    }
    var v13 string = "foo bar"
    for fi14 := 0; fi14 < 14; fi14++ {
        fmt.Println(v5)
        v2 -= v2
        for fi15 := 0; fi15 < 2; fi15++ {
            v1 = v1
            var v16 uint = 2
            var v17 uint8 = (v2 << ((730) & 63))
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = fi14
    }
    var v18 int16 = 42
    for fi19 := 0; fi19 < 12; fi19++ {
        for fi20 := 0; fi20 < 13; fi20++ {
            fmt.Println(v5, v18, v2)
            _ = fi20
        }
        _ = fi19
    }
    for fi21 := 0; fi21 < 17; fi21++ {
        for fi22 := 0; fi22 < 17; fi22++ {
            var v23 uint8 = v2
            _ = fi22
            _ = v23
        }
        var v24 int64 = 0
        _ = fi21
        _ = v24
    }
    v25 := GenS1{F0: 0}
    delete(v4, v13)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v18 = 32767
    v26 := GenS2{F0: 255, F1: v5, F2: 0}
    if (!(!false)) {
        var v27 uint32 = 512
        var v28 uint64 = (v1 + (^(v1)))
        _ = v27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
