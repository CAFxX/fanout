package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 int; F2 bool }

func main() {
    v1 := []int{127, 2147483647, -104}
    var v2 float64 = (((-2.25 * (0.1 / 1.5)) - 1.5) + -2.25)
    v3 := make(map[int]int)
    var v4 int32 = 0
    var v5 float64 = (v2 + v2)
    var v6 int16 = -32768
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v5, v6, v5)
    delete(v3, 1)
    if true {
        v7 := GenS1{F0: (v6 / (((v6 >> ((854) & 63))) | 1))}
        _ = v7
    } else {
        fmt.Println(v4, v5, v6)
        if false {
            var v8 uint16 = 32767
            _ = v8
        }
        v9 := GenS1{F0: (^(v6))}
        _ = v9
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        v5 = (float64(((0.0 + 1e10) / 1.0)) * 0.1)
        switch -9223372036854775808 {
        case -9223372036854775808:
            _ = len(v3)
        case 9223372036854775807:
            var v11 int8 = -128
            _ = v11
        default:
            if ("\n\t" > "\n\t") {
                switch -73 {
                case 846:
                    var v12 float32 = ((float32(-2.25) / 3.14159) / 0.1)
                    _ = v12
                case -9223372036854775808:
                    v1 = append(v1, (fi10 << ((325) & 63)))
                case 940:
                    var v13 int8 = -18
                    _ = v13
                }
                v6 -= (v6 & (v6 << ((386) & 63)))
            }
        }
        v2 = ((1.5 + 1.5) - ((v5 + 0.0) - 0.1))
        _ = fi10
    }
    if (true || (4294967295 >= 0)) {
        switch "foo bar" {
        case "\n\t":
            var v14 bool = false
            _ = v14
        case "a":
            var v15 uint64 = 494
            _ = v15
        default:
            v16 := GenS2{F0: 255, F1: -1, F2: true}
            _ = v16
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v5 = v2
    }
    v17 := GenS2{F0: 797, F1: 893, F2: (42 > (^(v4)))}
    v18 := GenS1{F0: (v6 & (v6 >> ((255) & 63)))}
    v1 = append(v1, -1)
    v5 += -2.25
    var v19 uint8 = 8
    v4 = (v4 >> ((4294967295) & 63))
    if ((((1e10 * v5) + (0.1 / -2.25)) - 0.1) <= v2) {
        var v20 uint16 = 308
        _ = v20
    } else {
        var v21 uint32 = 3
        var v22 float64 = (v5 + (0.1 / v5))
        _ = v21
        _ = v22
    }
    delete(v3, 804)
    for fi23 := 0; fi23 < 18; fi23++ {
        fmt.Println(v6)
        _ = fi23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
