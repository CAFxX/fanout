package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 uint64; F1 uint32; F2 int16 }
type GenS4 struct { F0 uint }

func main() {
    v1 := []uint32{}
    var v2 int16 = -32768
    var v3 uint32 = 2147483647
    v4 := make(map[string]bool)
    v5 := make(map[string]bool)
    switch "\n\t" {
    case "":
        if ("hello" != "foo bar") {
            fmt.Println(v2)
            v6 := GenS1{F0: v2}
            for fi7 := 0; fi7 < 19; fi7++ {
                v8 := GenS1{F0: v2}
                var v9 int = fi7
                _ = fi7
                _ = v8
                _ = v9
            }
            _ = v6
        } else {
            v10 := GenS1{F0: (v2 >> ((0) & 63))}
            v11 := GenS2{F0: v2, F1: -297}
            _ = v10
            _ = v11
        }
    default:
        var v12 uint64 = 18446744073709551615
        _ = v12
    }
    v2 = v2
    if (704 != -634) {
        var v13 uint = 1
        _ = v13
    }
    var v14 string = "foo bar"
    switch v14 {
    case "\n\t":
        fmt.Println(v3, v14, v2)
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        if false {
            for fi16 := 0; fi16 < 9; fi16++ {
                var v17 uint8 = 255
                if len(v1) > 0 {
                    v1[0] = v3
                }
                var v18 float32 = ((0.1 - float32(1e10)) - (0.0 + 0.1))
                _ = fi16
                _ = v17
                _ = v18
            }
            v19 := GenS3{F0: 351, F1: (v3 * (v3 | (v3 * v3))), F2: 32767}
            _ = v19
        } else {
            for fi20 := 0; fi20 < 5; fi20++ {
                v2 = (v2 ^ v2)
                _ = fi20
            }
            var v21 int64 = 846
            v22 := v5[v14]
            _ = v21
            _ = v22
        }
        var v23 int = fi15
        _ = fi15
        _ = v23
    }
    v24 := GenS4{F0: 255}
    var v25 uint16 = 557
    v14 = "\n\t"
    var v26 int16 = v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
