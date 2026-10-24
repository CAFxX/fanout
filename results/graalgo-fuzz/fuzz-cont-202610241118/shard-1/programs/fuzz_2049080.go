package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 uint32 }

func main() {
    v1 := []int{}
    var v2 uint = 255
    var v3 int8 = int8(v2)
    if false {
        if ((true || (v3 > (v3 / ((-1) | 1)))) || (828 < 1)) {
            fmt.Println(v2, v2)
        } else {
            v2 += v2
        }
        fmt.Println(v2)
    }
    for fi4 := 0; fi4 < 6; fi4++ {
        var v5 float64 = -2.25
        for fi6 := 0; fi6 < 19; fi6++ {
            v7 := GenS1{F0: 0}
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v5 += (v5 - v5)
            var v8 uint16 = 496
            _ = fi6
            _ = v7
            _ = v8
        }
        fmt.Println(v3, v5, v5)
        _ = fi4
        _ = v5
    }
    fmt.Println(v3)
    v9 := GenS2{F0: uint32(v2), F1: 91}
    switch (((("" + "hello") + "\n\t") + ("foo bar" + ("" + ("" + "foo bar")))) + "foo bar") {
    case "\n\t":
        var v10 uint16 = 0
        _ = v10
    }
    v3 = -66
    for fi11 := 0; fi11 < 14; fi11++ {
        if true {
            switch fi11 {
            case -698:
                fmt.Println(v3, v2)
            case -128:
                v3 = v3
            default:
                var v12 bool = (843 > 65535)
                _ = v12
            }
            v3 += (v3 >> (((v2 * v2)) & 63))
        }
        fmt.Println(v3, v2, v2)
        _ = fi11
    }
    if len(v1) > 0 {
        v1[0] = 553
    }
    for fi13 := 0; fi13 < 17; fi13++ {
        var v14 int8 = -128
        for fi15 := 0; fi15 < 11; fi15++ {
            switch "" {
            case "a":
                fmt.Println(fi13, fi13, fi13)
            }
            _ = fi15
        }
        v2 -= (v2 >> ((v2) & 63))
        _ = fi13
        _ = v14
    }
    v3 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v9
}
