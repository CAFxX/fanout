package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 uint16; F1 uint8; F2 bool }

func main() {
    v1 := []uint32{2, 4294967295, 1, 2147483648}
    v2 := []uint32{0, 2147483647}
    v3 := make(map[string]uint32)
    var v4 uint16 = 32767
    v3["foo bar"] = 509
    var v5 int = -1
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 float32 = (((1.5 - 3.14159) / (-2.25 + 3.14159)) / (((1e10 / ((1e10 + (0.0 * 0.1)) / 1.0)) / (-2.25 / ((0.1 * 0.0) - (3.14159 + 0.0)))) * float32(1e10)))
        v8 := GenS1{F0: 127, F1: -959}
        _ = fi6
        _ = v7
        _ = v8
    }
    v2 = append(v2, 0)
    switch (-(v5)) {
    case 127:
        v5 -= v5
    case -1:
        v5 = (v5 + (v5 + v5))
    }
    var v9 bool = ("" <= "")
    v9 = ((v5 ^ (v5 * 4294967295)) < (v5 & (v5 + (v5 + v5))))
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 string = ""
        switch v11 {
        case "foo bar":
            var v12 float64 = 1e10
            _ = v12
        case "\n\t":
            var v13 int32 = 42
            _ = v13
        case "a":
            var v14 float64 = 0.0
            _ = v14
        default:
            v4 += 545
        }
        _ = fi10
        _ = v11
    }
    switch 4294967295 {
    case 127:
        var v15 float64 = 1e10
        _ = v15
    case 32768:
        v4 += v4
    case -2147483648:
        for fi16 := 0; fi16 < 14; fi16++ {
            v9 = (!false)
            for fi17 := 0; fi17 < 19; fi17++ {
                v5 -= (v5 ^ (v5 % ((-2147483648) | 1)))
                v18 := GenS1{F0: 127, F1: -202}
                if (1 != 65535) {
                    v19 := GenS1{F0: 128, F1: -311}
                    var v20 int16 = 32767
                    _ = v19
                    _ = v20
                }
                _ = fi17
                _ = v18
            }
            var v21 bool = v9
            _ = fi16
            _ = v21
        }
    }
    if v9 {
        var v22 bool = ((1.5 == (0.1 - 1e10)) || false)
        _ = v22
    } else {
        fmt.Println(v4)
        var v23 uint32 = 65535
        _ = v23
    }
    if v9 {
        _ = len(v3)
    } else {
        switch "hello" {
        case "hello":
            if ((v4 - (v4 >> ((42) & 63))) != v4) {
                fmt.Println(v5)
            } else {
                fmt.Println(v4, v5, v5)
                for fi24 := 0; fi24 < 20; fi24++ {
                    v25 := GenS1{F0: 42, F1: 0}
                    _ = fi24
                    _ = v25
                }
            }
        case "\n\t":
            v26 := GenS2{F0: v4, F1: 0, F2: (float64(1.5) <= ((1e10 - -2.25) - -2.25))}
            _ = v26
        case "foo bar":
            if len(v2) > 0 {
                v2[0] = 65535
            }
        }
        var v27 int = (v5 << ((311) & 63))
        v28 := GenS2{F0: v4, F1: 127, F2: v9}
        _ = v27
        _ = v28
    }
    fmt.Println(v9)
    var v29 float64 = 1e10
    var v30 int = (v5 / ((0) | 1))
    if len(v2) > 0 {
        v2[0] = 839
    }
    for fi31 := 0; fi31 < 15; fi31++ {
        switch "hello" {
        case "hello":
            var v32 bool = v9
            _ = v32
        default:
            fmt.Println(v29, v29)
        }
        _ = fi31
    }
    v5 -= 127
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v9
}
