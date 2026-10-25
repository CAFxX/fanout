package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int; F1 uint }

func main() {
    var v1 uint16 = 252
    var v2 uint = 18446744073709551615
    var v3 bool = false
    v4 := []string{}
    var v5 uint = (v2 * (v2 + v2))
    if true {
        v1 = 65535
        fmt.Println(v5, v5)
    } else {
        v5 += 146
        v6 := GenS1{F0: 2147483648}
        for fi7 := 0; fi7 < 9; fi7++ {
            if false {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi7
        }
        _ = v6
    }
    _ = len(v4)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v2 += (v2 / ((758) | 1))
    v1 -= (v1 + v1)
    v2 += 174
    for fi8 := 0; fi8 < 2; fi8++ {
        v9 := GenS2{F0: (0 > 2147483647), F1: ("foo bar" + "foo bar")}
        var v10 int = (fi8 >> ((v2) & 63))
        var v11 string = "a"
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    var v12 int32 = -736
    v13 := GenS3{F0: -1}
    for fi14 := 0; fi14 < 15; fi14++ {
        v12 -= v12
        _ = fi14
    }
    var v15 float32 = ((1e10 * 0.1) + -2.25)
    var v16 uint = (^(v5))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v17 uint = 2
    v18 := GenS4{F0: 2147483647, F1: uint(v1)}
    switch 1 {
    case -456:
        var v19 string = "\n\t"
        _ = v19
    case -2147483648:
        fmt.Println(v5)
    default:
        v2 = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
