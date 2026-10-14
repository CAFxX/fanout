package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint32 }
type GenS2 struct { F0 uint; F1 int32; F2 int }
type GenS3 struct { F0 string; F1 int32 }
type GenS4 struct { F0 string }

func main() {
    var v1 int16 = -848
    var v2 uint32 = 408
    v3 := []uint32{}
    v4 := make(map[uint32]int)
    var v5 int = 154
    var v6 float32 = (0.0 * 3.14159)
    switch ("foo bar" + ((((("\n\t" + "hello") + "a") + ("hello" + "")) + ("hello" + "\n\t")) + (("" + "") + ("a" + ("" + ""))))) {
    case "":
        v6 -= 0.0
    default:
        switch "foo bar" {
        case "":
            var v7 int = v5
            _ = v7
        case "a":
            delete(v4, v2)
        case "hello":
            var v8 float32 = float32((-2.25 + 1.5))
            _ = v8
        }
    }
    if true {
        switch "\n\t" {
        case "hello":
            var v9 uint16 = 42
            _ = v9
        case "foo bar":
            var v10 int8 = 1
            _ = v10
        case "\n\t":
            _ = len(v3)
        }
    }
    v4[(v2 + (v2 & v2))] = (v5 / (((v5 / ((-1) | 1))) | 1))
    v5 += -9223372036854775808
    var v11 uint64 = uint64(v2)
    var v12 uint16 = 322
    var v13 float64 = (3.14159 - ((0.1 + -2.25) * (((0.0 - (((3.14159 * ((((((1e10 + 3.14159) - 3.14159) - ((1.5 + ((1.5 + 0.1) * 1.5)) * 0.0)) - 1.5) + (1e10 + float64((1e10 - v6)))) - 3.14159)) / ((((-2.25 / 1.0) / 1.0) * 0.1) / ((float64(v6) * (1e10 * -2.25)) - ((1e10 / (0.1 + (0.1 / (3.14159 + 0.0)))) - -2.25)))) + (1e10 + 3.14159))) - 0.1) * 1.5)))
    v11 = v11
    fmt.Println(v1, v6, v12)
    var v14 uint16 = (^(v12))
    v1 -= v1
    for fi15 := 0; fi15 < 15; fi15++ {
        v16 := GenS1{F0: -1, F1: (v1 % ((-32768) | 1)), F2: v2}
        if true {
            var v17 int16 = v1
            switch (^(v5)) {
            case 9223372036854775807:
                v18 := GenS2{F0: 65535, F1: 0, F2: 2147483647}
                _ = v18
            case -2147483648:
                var v19 uint32 = 4294967295
                _ = v19
            case 127:
                v20 := GenS2{F0: 2, F1: 2147483647, F2: (fi15 & fi15)}
                _ = v20
            default:
                var v21 int8 = -128
                _ = v21
            }
            _ = v17
        }
        var v22 float64 = 1.5
        _ = fi15
        _ = v16
        _ = v22
    }
    v23 := GenS3{F0: ("foo bar" + "a"), F1: 0}
    var v24 uint = 2
    v25 := GenS4{F0: ("hello" + ((("" + "\n\t") + "\n\t") + ("\n\t" + "hello")))}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
