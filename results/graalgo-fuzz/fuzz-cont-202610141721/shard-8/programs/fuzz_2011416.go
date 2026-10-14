package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }

func main() {
    v1 := []uint32{}
    var v2 int32 = -1
    var v3 int16 = -970
    var v4 int16 = (v3 / ((42) | 1))
    var v5 int16 = (-(v4))
    v1 = append(v1, 65535)
    switch "a" {
    case "a":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v2 -= v2
    switch 9223372036854775807 {
    case -896:
        var v6 int64 = 881
        _ = v6
    case 9223372036854775807:
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 bool = (2 < 132)
            _ = fi7
            _ = v8
        }
    case 32768:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        var v9 float32 = 1.5
        _ = v9
    }
    var v10 uint = 527
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        var v12 int32 = v2
        fmt.Println(v5, fi11, v5)
        _ = fi11
        _ = v12
    }
    var v13 bool = false
    v3 = (-(v4))
    switch "foo bar" {
    case "":
        v14 := GenS1{F0: 0, F1: false}
        _ = v14
    case "a":
        v10 -= v10
    case "\n\t":
        v3 -= (v5 + (v5 << (((v10 % ((v10) | 1))) & 63)))
    default:
        v13 = v13
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        for fi16 := 0; fi16 < 14; fi16++ {
            fmt.Println(v4, v10)
            _ = fi16
        }
        if true {
            var v17 uint8 = 44
            _ = v17
        }
        fmt.Println(v10, v4, v13)
        _ = fi15
    }
    var v18 uint = (^(v10))
    var v19 string = "foo bar"
    v5 = (v5 * (v5 + (v5 + v4)))
    v18 += 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
