package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 int8; F1 int32 }

func main() {
    var v1 int16 = 0
    v2 := make(map[int]bool)
    var v3 int = -523
    v4 := make(map[uint32]int)
    for fi5 := 0; fi5 < 14; fi5++ {
        v6 := v4[4294967295]
        var v7 int64 = 871
        _ = fi5
        _ = v6
        _ = v7
    }
    v3 = v3
    v8 := []int{1, (v3 ^ v3), 4294967295, v3}
    v9 := GenS1{F0: true, F1: 0}
    v1 = (-(v1))
    switch (v3 / ((v3) | 1)) {
    case 1:
        if true {
            var v10 uint32 = 2
            for fi11 := 0; fi11 < 6; fi11++ {
                v3 += (v3 * v3)
                v12 := GenS2{F0: 42, F1: -682}
                v2[fi11] = (!false)
                _ = fi11
                _ = v12
            }
            _ = v10
        }
    default:
        delete(v4, 238)
    }
    switch "a" {
    case "\n\t":
        v1 += v1
    }
    v3 = v3
    if false {
        v1 = (v1 << ((255) & 63))
        switch ("a" + "\n\t") {
        case "a":
            for fi13 := 0; fi13 < 16; fi13++ {
                v14 := v4[4294967295]
                var v15 int8 = 127
                _ = fi13
                _ = v14
                _ = v15
            }
        case "\n\t":
            if len(v8) > 0 {
                v8[0] = (^(v3))
            }
        case "":
            v1 -= -32768
        }
    } else {
        v3 = (v3 | 42)
        v1 -= -1
        var v16 uint8 = 255
        _ = v16
    }
    var v17 int8 = -1
    var v18 float64 = 0.0
    var v19 int = -988
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
