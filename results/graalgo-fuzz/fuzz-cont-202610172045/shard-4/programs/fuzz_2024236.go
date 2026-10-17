package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int8 }
type GenS2 struct { F0 int; F1 uint64; F2 int8 }
type GenS3 struct { F0 bool; F1 string; F2 string }
type GenS4 struct { F0 int16; F1 int8 }
func genFn0(p0 int64) int64 {
    switch ("\n\t" + "a") {
    case "foo bar":
        v7 := []int{-2147483648, 0, 127, 42}
        _ = v7
    default:
        var v8 int32 = 2147483647
        _ = v8
    }
    var v9 float64 = float64(0.0)
    v10 := []int{4294967295, 706, 2147483647, 2147483647}
    var v11 bool = false
    _ = p0
    _ = v10
    _ = v11
    _ = v9
    return -376
}

func main() {
    v1 := []string{"", "hello"}
    var v2 int16 = 32767
    v3 := []string{"foo bar"}
    var v4 int8 = 98
    var v5 bool = (358 <= 2147483647)
    var v6 int = 32768
    v2 = (v2 % ((v2) | 1))
    fmt.Println(v5, v4)
    v4 -= (v4 / (((-(v4))) | 1))
    var v12 uint = 559
    var v13 int = v6
    v14 := GenS1{F0: v12, F1: uint8(v12), F2: (^(v4))}
    switch "foo bar" {
    case "foo bar":
        v4 = 100
    case "\n\t":
        for fi15 := 0; fi15 < 10; fi15++ {
            v12 = v12
            v16 := GenS1{F0: v12, F1: 149, F2: (v4 >> (((v12 % (((v12 - (v12 << (((v12 ^ v12)) & 63)))) | 1))) & 63))}
            _ = fi15
            _ = v16
        }
    case "hello":
        for fi17 := 0; fi17 < 15; fi17++ {
            v18 := GenS1{F0: v12, F1: 220, F2: v4}
            v13 = v13
            _ = fi17
            _ = v18
        }
    default:
        switch (-(v13)) {
        case -426:
            v19 := GenS2{F0: (v6 << ((255) & 63)), F1: 18446744073709551615, F2: v4}
            _ = v19
        default:
            v13 -= 483
        }
    }
    v20 := GenS3{F0: (65535 > 246), F1: ("\n\t" + "a"), F2: ""}
    var v21 uint16 = 1
    if v5 {
        var v22 uint32 = 1
        if len(v3) > 0 {
            v3[0] = ""
        }
        fmt.Println(v4, v21)
        _ = v22
    } else {
        var v23 bool = false
        switch (v13 + v6) {
        case -1:
            v6 = (v13 * v13)
        }
        _ = v23
    }
    for fi24 := 0; fi24 < 11; fi24++ {
        var v25 uint8 = 0
        _ = fi24
        _ = v25
    }
    var v26 int16 = v2
    v27 := GenS4{F0: (^(v26)), F1: (v4 / (((v4 + (-(v4)))) | 1))}
    v13 -= v13
    v5 = (-2147483648 == -720)
    var v28 uint32 = 2
    v4 -= int8(v12)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
