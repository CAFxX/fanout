package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }

func main() {
    v1 := []string{""}
    var v2 int8 = -106
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]bool)
    for fi5 := 0; fi5 < 5; fi5++ {
        v6 := GenS1{F0: 497, F1: 18446744073709551615}
        for fi7 := 0; fi7 < 20; fi7++ {
            var v8 float64 = ((1.5 / 1.0) * ((1.5 + 1e10) + (0.0 / -2.25)))
            _ = fi7
            _ = v8
        }
        _ = fi5
        _ = v6
    }
    v2 = v2
    v2 = v2
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 uint8 = 98
        var v11 int8 = (v2 - (-(v2)))
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 8; fi12++ {
        delete(v3, 32768)
        for fi13 := 0; fi13 < 17; fi13++ {
            v4[79] = true
            var v14 uint16 = 0
            var v15 int8 = (v2 / ((v2) | 1))
            _ = fi13
            _ = v14
            _ = v15
        }
        fmt.Println(fi12)
        _ = fi12
    }
    fmt.Println(v2)
    if false {
        if len(v1) > 0 {
            v1[0] = (("a" + "") + "")
        }
    }
    if true {
        if false {
            var v16 uint16 = 65535
            _ = v16
        }
    }
    var v17 int = 646
    var v18 uint64 = 4294967295
    v17 += 291
    var v19 uint = 2
    var v20 bool = true
    v2 -= -25
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
