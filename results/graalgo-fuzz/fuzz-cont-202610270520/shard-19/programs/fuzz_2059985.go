package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint32 }
type GenS2 struct { F0 int16 }
func genFn0() string {
    for fi7 := 0; fi7 < 4; fi7++ {
        fmt.Println(fi7)
        _ = fi7
    }
    var v8 string = "a"
    for fi9 := 0; fi9 < 16; fi9++ {
        var v10 int = 42
        var v11 float32 = ((1.5 / 0.1) / 0.1)
        v11 = ((v11 * v11) - (0.1 / v11))
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = v8
    return v8
}

func main() {
    var v1 int8 = -128
    var v2 uint = 276
    var v3 uint64 = 2
    var v4 string = "a"
    var v5 string = v4
    var v6 uint8 = 212
    v12 := GenS1{F0: 1, F1: (v2 % ((v2) | 1)), F2: 0}
    v1 = (v1 >> ((1) & 63))
    fmt.Println(v3, v6, v1)
    var v13 uint8 = (v6 + v6)
    v3 += (^(v3))
    v1 = (v1 / ((0) | 1))
    var v14 uint32 = 2
    var v15 int32 = 42
    v16 := []int{-545, -1, 128}
    if len(v16) > 0 {
        v16[0] = -2147483648
    }
    if ((true && (-32768 != 529)) && ((-128 <= v1) && (!((((1e10 / -2.25) - 0.1) - 3.14159) != 0.1)))) {
        var v17 int = 32768
        _ = v17
    }
    switch "a" {
    case "hello":
        if (((0 > 1) || false) || (!true)) {
            v1 = (v1 % (((^(v1))) | 1))
            v1 = int8(v15)
        }
    case "a":
        var v18 uint32 = v14
        _ = v18
    case "\n\t":
        if len(v16) > 0 {
            v16[0] = 202
        }
    default:
        switch "" {
        case "a":
            v5 = (("hello" + v5) + v5)
        case "foo bar":
            fmt.Println(v2, v3, v1)
        default:
            for fi19 := 0; fi19 < 13; fi19++ {
                var v20 uint8 = (v6 << ((v2) & 63))
                var v21 int8 = (v1 & -128)
                _ = fi19
                _ = v20
                _ = v21
            }
        }
    }
    var v22 uint8 = (^(v6))
    v23 := GenS1{F0: (v15 ^ v15), F1: v2, F2: (v14 - v14)}
    for fi24 := 0; fi24 < 17; fi24++ {
        var v25 string = (("" + ((v4 + v5) + ("" + "a"))) + ((v5 + v4) + (("" + (v5 + v5)) + v5)))
        v3 += (v3 * v3)
        v26 := GenS2{F0: -32768}
        _ = fi24
        _ = v25
        _ = v26
    }
    v16 = append(v16, -2147483648)
    v6 -= v13
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
