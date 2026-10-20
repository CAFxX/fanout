package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int }
type GenS2 struct { F0 string; F1 int64; F2 uint64 }

func main() {
    v1 := []string{}
    var v2 int32 = -2147483648
    v3 := []uint32{}
    fmt.Println(v2)
    fmt.Println(v2)
    v4 := GenS1{F0: 42, F1: 127, F2: -2147483648}
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v2)
    default:
        v1 = append(v1, ("" + "hello"))
    }
    v2 -= (v2 & (v2 >> ((255) & 63)))
    if (216 == 2) {
        var v5 float64 = float64(0.1)
        if len(v3) > 0 {
            v3[0] = 0
        }
        _ = v5
    } else {
        v6 := GenS1{F0: 878, F1: 32768, F2: 418}
        _ = len(v1)
        _ = v6
    }
    var v7 uint8 = 160
    v8 := GenS2{F0: ("\n\t" + ""), F1: -9223372036854775808, F2: 18446744073709551615}
    fmt.Println(v7)
    var v9 uint16 = 65535
    for fi10 := 0; fi10 < 9; fi10++ {
        var v11 uint = 121
        for fi12 := 0; fi12 < 6; fi12++ {
            var v13 int = 0
            _ = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 2; fi14++ {
            v11 = 235
            if len(v1) > 0 {
                v1[0] = ("hello" + ("\n\t" + ""))
            }
            _ = fi14
        }
        _ = fi10
        _ = v11
    }
    v7 += v7
    for fi15 := 0; fi15 < 15; fi15++ {
        var v16 float64 = 0.0
        _ = fi15
        _ = v16
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
