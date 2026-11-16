package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int }
type GenS2 struct { F0 uint; F1 uint }

func main() {
    v1 := []int{127, -2147483648}
    var v2 uint = 42
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]string)
    var v5 uint8 = 1
    if true {
        var v6 float32 = float32(-2.25)
        _ = v6
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi7 := 0; fi7 < 9; fi7++ {
            delete(v3, 477)
            _ = fi7
        }
        v5 -= (^(v5))
    }
    for fi8 := 0; fi8 < 17; fi8++ {
        fmt.Println(v5, v5, fi8)
        var v9 string = ("\n\t" + "\n\t")
        switch fi8 {
        case -2147483648:
            var v10 float64 = float64(3.14159)
            _ = v10
        case 0:
            var v11 float64 = -2.25
            _ = v11
        case 533:
            v12 := GenS1{F0: 32767, F1: 65535, F2: (^(fi8))}
            _ = v12
        default:
            if false {
                v5 += (v5 % (((v5 & (v5 ^ (v5 - v5)))) | 1))
                fmt.Println(v5, v9, v9)
            }
        }
        _ = fi8
        _ = v9
    }
    delete(v4, 4294967295)
    fmt.Println(v2)
    var v13 int64 = -986
    if (v13 >= v13) {
        v14 := GenS2{F0: (^(v2)), F1: uint(v5)}
        _ = v14
    }
    v5 = v5
    switch ("\n\t" + "") {
    case "foo bar":
        var v15 int16 = 300
        _ = v15
    case "\n\t":
        switch -691 {
        case 1:
            fmt.Println(v13, v2)
        default:
            v13 = v13
        }
    default:
        v16 := GenS1{F0: -686, F1: 0, F2: 9223372036854775807}
        _ = v16
    }
    if (("" + "\n\t") < "hello") {
        fmt.Println(v2, v2)
    }
    for fi17 := 0; fi17 < 6; fi17++ {
        v18 := GenS2{F0: v2, F1: uint(v5)}
        _ = fi17
        _ = v18
    }
    switch ("foo bar" + "hello") {
    case "foo bar":
        v19 := GenS1{F0: -1, F1: 65535, F2: 42}
        _ = v19
    case "":
        v2 = (^(v2))
    case "\n\t":
        v5 = (v5 * (v5 * (v5 << (((v2 / (((v2 ^ uint(v5))) | 1))) & 63))))
    default:
        v3[2] = 65535
    }
    var v20 uint64 = 146
    if (818 > (v13 / (((-(v13))) | 1))) {
        switch -160 {
        case -2147483648:
            if ((1e10 - 1e10) == 3.14159) {
                v21 := v3[0]
                v13 -= 381
                v22 := GenS1{F0: 1000, F1: 32767, F2: 32768}
                _ = v21
                _ = v22
            } else {
                v23 := GenS1{F0: -741, F1: 255, F2: 375}
                _ = v23
            }
        }
    } else {
        v2 -= v2
        v13 = v13
        fmt.Println(v2)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
