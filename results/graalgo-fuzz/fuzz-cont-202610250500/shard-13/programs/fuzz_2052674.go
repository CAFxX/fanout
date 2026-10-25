package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 int32 = -852
    var v2 int8 = 100
    v3 := []string{"hello", "foo bar"}
    var v4 int = -1
    var v5 uint16 = 212
    var v6 uint64 = 1
    v7 := GenS1{F0: v4}
    switch ("a" + "a") {
    case "a":
        var v8 uint32 = 2
        _ = v8
    case "":
        fmt.Println(v2, v6)
    default:
        v5 -= (v5 + v5)
    }
    fmt.Println(v5, v4, v6)
    for fi9 := 0; fi9 < 15; fi9++ {
        if ((1 == 2) && ((0.1 / (1e10 - -2.25)) < (float32((1e10 * 3.14159)) / 1e10))) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v10 int32 = v1
            fmt.Println(v2, v10, v1)
            _ = v10
        }
        _ = fi9
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        v6 -= (v6 + 676)
        _ = len(v3)
        switch ("foo bar" + "foo bar") {
        case "hello":
            var v12 float64 = (((1.5 / 1.5) - -2.25) / 1e10)
            _ = v12
        case "":
            v4 += (v4 | (v4 / (((fi11 >> ((186) & 63))) | 1)))
        case "a":
            var v13 int32 = (v1 - 148)
            _ = v13
        default:
            v2 += (v2 / (((v2 >> ((4294967295) & 63))) | 1))
        }
        _ = fi11
    }
    if (v5 >= uint16(v6)) {
        v4 = (v4 ^ 127)
        v14 := GenS1{F0: v4}
        var v15 uint8 = 14
        _ = v14
        _ = v15
    } else {
        v16 := GenS1{F0: (v4 ^ v4)}
        v17 := GenS2{F0: 0}
        fmt.Println(v4, v5)
        _ = v16
        _ = v17
    }
    v1 = 0
    var v18 uint16 = uint16(v6)
    v4 += (v4 - (v4 % (((v4 + (v4 >> ((0) & 63)))) | 1)))
    v4 = v4
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
