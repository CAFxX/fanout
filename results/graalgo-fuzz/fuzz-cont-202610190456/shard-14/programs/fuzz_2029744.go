package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 int64; F2 int16 }

func main() {
    v1 := []uint32{0, 4294967295, 2147483647}
    var v2 int8 = -65
    v3 := make(map[string]uint32)
    var v4 uint64 = 2
    var v5 uint8 = 199
    switch 4294967295 {
    case 239:
        switch "a" {
        case "hello":
            var v6 int16 = -32768
            _ = v6
        case "foo bar":
            switch ("hello" + "foo bar") {
            case "\n\t":
                switch -9223372036854775808 {
                case -959:
                    var v7 float64 = float64(3.14159)
                    _ = v7
                case 32767:
                    if ("" > "\n\t") {
                        var v8 uint32 = 0
                        var v9 int = -9223372036854775808
                        _ = v8
                        _ = v9
                    } else {
                        var v10 int = -2147483648
                        v11 := GenS1{F0: v2, F1: v2, F2: 480}
                        var v12 int8 = v2
                        _ = v10
                        _ = v11
                        _ = v12
                    }
                case 1:
                    for fi13 := 0; fi13 < 15; fi13++ {
                        fmt.Println(fi13)
                        _ = fi13
                    }
                }
            case "hello":
                for fi14 := 0; fi14 < 1; fi14++ {
                    for fi15 := 0; fi15 < 9; fi15++ {
                        var v16 uint8 = v5
                        v2 = 1
                        v17 := GenS2{F0: -85}
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    _ = fi14
                }
            case "":
                v18 := GenS2{F0: -941}
                _ = v18
            default:
                v5 = v5
            }
        case "":
            fmt.Println(v5, v2, v2)
        }
    }
    for fi19 := 0; fi19 < 16; fi19++ {
        var v20 int = fi19
        v5 += v5
        _ = fi19
        _ = v20
    }
    v5 = v5
    var v21 int = 32767
    v2 += (v2 | 1)
    var v22 float64 = (-2.25 - (float64(-2.25) - 3.14159))
    var v23 bool = false
    v24 := GenS2{F0: 2147483647}
    switch (^(v21)) {
    case -1:
        var v25 uint32 = 2147483648
        _ = v25
    case 0:
        v26 := GenS1{F0: (v2 / (((v2 / ((v2) | 1))) | 1)), F1: (v2 << ((42) & 63)), F2: 9223372036854775807}
        _ = v26
    case -428:
        v27 := GenS3{F0: -116, F1: 42, F2: 1}
        _ = v27
    default:
        var v28 float64 = ((v22 * -2.25) / v22)
        _ = v28
    }
    v4 += (v4 % ((v4) | 1))
    var v29 int64 = -1
    fmt.Println(v4)
    v30 := GenS1{F0: (-(v2)), F1: v2, F2: -9223372036854775808}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
