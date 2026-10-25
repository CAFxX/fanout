package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 bool }

func main() {
    var v1 int16 = 1000
    v2 := make(map[uint32]bool)
    v3 := []uint32{2}
    var v4 int = -2147483648
    v1 += 42
    v4 = (v4 % (((v4 & 440)) | 1))
    if (738 == 2) {
        switch "\n\t" {
        case "\n\t":
            v5 := GenS1{F0: -1, F1: 2147483648, F2: (false || false)}
            _ = v5
        default:
            fmt.Println(v4, v4)
        }
    } else {
        v6 := v2[139]
        var v7 uint8 = 127
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        v9 := GenS1{F0: -1, F1: 157, F2: (4294967295 != 4294967295)}
        v10 := v2[280]
        _ = fi8
        _ = v10
        _ = v9
    }
    if true {
        v1 -= v1
        v4 -= 4294967295
    }
    v3 = append(v3, 230)
    var v11 uint8 = 128
    if (-1 == v4) {
        if false {
            fmt.Println(v4, v11, v1)
        } else {
            switch "" {
            case "foo bar":
                v3 = append(v3, 990)
            case "\n\t":
                for fi12 := 0; fi12 < 4; fi12++ {
                    v4 = (-(fi12))
                    var v13 float64 = 0.1
                    _ = fi12
                    _ = v13
                }
            default:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        }
    }
    v1 -= v1
    fmt.Println(v11, v1, v11)
    fmt.Println(v1, v1)
    if (1 < 4294967295) {
        v1 = v1
        var v14 float64 = -2.25
        _ = v14
    } else {
        var v15 int = 32767
        _ = v15
    }
    for fi16 := 0; fi16 < 13; fi16++ {
        v4 = (-(fi16))
        _ = fi16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
