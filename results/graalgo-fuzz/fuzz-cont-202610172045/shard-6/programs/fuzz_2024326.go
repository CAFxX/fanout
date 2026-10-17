package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 int64 }

func main() {
    v1 := []uint32{405}
    var v2 uint8 = 1
    var v3 int = -929
    v4 := make(map[string]uint32)
    var v5 uint32 = 2147483647
    switch v3 {
    case 127:
        v3 = -486
    default:
        var v6 int = v3
        _ = v6
    }
    v7 := GenS1{F0: 42}
    for fi8 := 0; fi8 < 15; fi8++ {
        for fi9 := 0; fi9 < 18; fi9++ {
            v3 = (^(v3))
            if len(v1) > 0 {
                v1[0] = v5
            }
            var v10 float32 = (-2.25 - -2.25)
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    for fi11 := 0; fi11 < 6; fi11++ {
        switch ("a" + "\n\t") {
        case "\n\t":
            var v12 int32 = 1
            _ = v12
        case "hello":
            fmt.Println(v2, v5)
        }
        _ = fi11
    }
    v3 = v3
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v13 := v4["foo bar"]
    for fi14 := 0; fi14 < 19; fi14++ {
        switch 32768 {
        case 645:
            v13 += 1
        default:
            if true {
                _ = len(v1)
                for fi15 := 0; fi15 < 20; fi15++ {
                    v16 := GenS2{F0: false, F1: -1}
                    _ = fi15
                    _ = v16
                }
            }
        }
        v2 -= uint8(v5)
        if len(v1) > 0 {
            v1[0] = (v13 >> ((852) & 63))
        }
        _ = fi14
    }
    var v17 uint32 = (^(v5))
    if (-32768 >= 32767) {
        v5 += v17
        v18 := v4["foo bar"]
        _ = v18
    } else {
        v17 -= v13
        v2 = (v2 | (v2 * (^(v2))))
        delete(v4, "\n\t")
    }
    if true {
        v5 += v17
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
