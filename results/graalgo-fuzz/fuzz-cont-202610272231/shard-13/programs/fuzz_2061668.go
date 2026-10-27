package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 int8 }
type GenS2 struct { F0 uint; F1 uint8; F2 uint }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{}
    var v3 int16 = 1000
    v4 := make(map[int]bool)
    v5 := make(map[int]uint32)
    if true {
        v3 += (^(v3))
        for fi6 := 0; fi6 < 3; fi6++ {
            var v7 int8 = 0
            v7 += (v7 / ((-128) | 1))
            _ = fi6
            _ = v7
        }
        v3 -= (v3 & v3)
    }
    v8 := GenS1{F0: 42, F1: 127, F2: -128}
    var v9 uint = 255
    v10 := v5[0]
    for fi11 := 0; fi11 < 19; fi11++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v12 uint64 = 266
        v10 += (^(v10))
        _ = fi11
        _ = v12
    }
    fmt.Println(v3, v3)
    v13 := GenS2{F0: (v9 | (v9 % ((v9) | 1))), F1: 0, F2: 984}
    v9 += 859
    v14 := GenS1{F0: 42, F1: 73, F2: 0}
    v10 = (v10 * (v10 + (v10 | (^(v10)))))
    fmt.Println(v10, v3)
    if false {
        v15 := GenS2{F0: (v9 - v9), F1: uint8(v10), F2: 552}
        switch -1 {
        case 127:
            var v16 uint16 = 164
            _ = v16
        case 164:
            v9 = (v9 >> ((v9) & 63))
        case -909:
            var v17 int16 = (v3 * 32767)
            _ = v17
        }
        _ = v15
    } else {
        switch 2147483647 {
        case 936:
            var v18 float64 = (-2.25 * 0.0)
            _ = v18
        default:
            switch "a" {
            case "":
                v9 = (^(v9))
            case "a":
                var v19 int64 = 9223372036854775807
                _ = v19
            }
        }
        if (42 == -2147483648) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        var v20 int32 = 840
        _ = v20
    }
    for fi21 := 0; fi21 < 15; fi21++ {
        _ = len(v2)
        _ = fi21
    }
    fmt.Println(v9, v3)
    v10 += (v10 & (v10 & 270))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
