package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int8 }
type GenS2 struct { F0 bool; F1 uint8 }

func main() {
    var v1 int16 = 220
    var v2 uint16 = 2
    var v3 int8 = 1
    var v4 string = "a"
    var v5 string = (v4 + v4)
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := []int{-9223372036854775808}
        _ = fi6
        _ = v7
    }
    v8 := []int{-9223372036854775808}
    v4 = v4
    for fi9 := 0; fi9 < 19; fi9++ {
        v10 := GenS1{F0: 246, F1: (v2 << ((1) & 63)), F2: 100}
        var v11 int = (fi9 % ((fi9) | 1))
        _ = fi9
        _ = v10
        _ = v11
    }
    switch -9223372036854775808 {
    case -2147483648:
        v2 += 255
    }
    var v12 float32 = float32((1e10 + ((0.0 - float64(float32(3.14159))) * float64(1.5))))
    switch ("\n\t" + v4) {
    case "":
        for fi13 := 0; fi13 < 19; fi13++ {
            v14 := GenS1{F0: 743, F1: 1, F2: (v3 * (v3 | v3))}
            for fi15 := 0; fi15 < 1; fi15++ {
                if false {
                    fmt.Println(fi13)
                }
                _ = fi15
            }
            var v16 uint8 = 1
            _ = fi13
            _ = v14
            _ = v16
        }
    case "hello":
        v17 := GenS1{F0: 18446744073709551615, F1: 429, F2: 1}
        _ = v17
    case "foo bar":
        var v18 float64 = (1.5 + (0.0 - (1.5 * float64(-2.25))))
        _ = v18
    default:
        fmt.Println(v1, v5, v2)
    }
    switch -9223372036854775808 {
    case 32768:
        _ = len(v8)
    default:
        v5 = v5
    }
    if len(v8) > 0 {
        v8[0] = 128
    }
    if (v3 == 0) {
        var v19 int = 42
        v20 := GenS1{F0: 18446744073709551615, F1: 609, F2: (^(v3))}
        v21 := GenS2{F0: true, F1: 114}
        _ = v19
        _ = v20
        _ = v21
    } else {
        var v22 int = 730
        v1 = -424
        _ = v22
    }
    fmt.Println(v3)
    if true {
        if true {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
        }
        v5 = v5
    }
    var v23 int = -355
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
