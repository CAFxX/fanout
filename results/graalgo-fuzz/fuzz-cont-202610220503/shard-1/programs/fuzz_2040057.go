package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
func genFn0(p0 int) uint32 {
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 uint16 = 32767
        _ = fi6
        _ = v7
    }
    var v8 uint64 = 247
    _ = p0
    _ = v8
    return 206
}

func main() {
    var v1 uint8 = 2
    v2 := []int{}
    var v3 uint8 = v1
    var v4 uint32 = 2147483648
    var v5 int64 = 1
    if (!false) {
        v9 := GenS1{F0: -39, F1: (v4 << ((255) & 63))}
        _ = v9
    }
    v3 -= v3
    if ((1 != 18446744073709551615) && false) {
        v1 += 1
        var v10 float64 = 3.14159
        _ = v10
    }
    if false {
        fmt.Println(v3)
        var v11 uint32 = (^(v4))
        for fi12 := 0; fi12 < 10; fi12++ {
            v13 := GenS1{F0: 2147483647, F1: uint32(v1)}
            _ = fi12
            _ = v13
        }
        _ = v11
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v14 uint = 319
    if (false && (true && true)) {
        fmt.Println(v3, v4)
        fmt.Println(v5, v5, v3)
    } else {
        var v15 uint8 = 128
        if false {
            var v16 uint64 = 796
            v16 = (v16 & v16)
            v4 -= (v4 / (((v4 & 42)) | 1))
            _ = v16
        }
        _ = v15
    }
    if (-196 < 1) {
        switch "hello" {
        case "":
            for fi17 := 0; fi17 < 6; fi17++ {
                var v18 int32 = -2147483648
                if len(v2) > 0 {
                    v2[0] = (fi17 & 224)
                }
                _ = fi17
                _ = v18
            }
        default:
            if true {
                switch "a" {
                case "foo bar":
                    var v19 int32 = 2147483647
                    _ = v19
                }
            }
        }
        fmt.Println(v3, v4)
    } else {
        if (false && ((genFn0(127) == (v4 % ((476) | 1))) && false)) {
            for fi20 := 0; fi20 < 2; fi20++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi20
            }
            if len(v2) > 0 {
                v2[0] = -608
            }
        }
        var v21 int64 = 1
        _ = v21
    }
    var v22 uint64 = 741
    switch "a" {
    case "\n\t":
        v14 = (v14 ^ (v14 & v14))
    case "foo bar":
        fmt.Println(v3)
    case "a":
        v2 = append(v2, -2147483648)
    default:
        v14 = 250
    }
    v1 = 144
    v23 := GenS1{F0: -2147483648, F1: uint32(v22)}
    v4 = (^(v4))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
