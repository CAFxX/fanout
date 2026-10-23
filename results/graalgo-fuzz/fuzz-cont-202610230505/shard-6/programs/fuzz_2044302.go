package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 uint32; F1 uint8; F2 string }

func main() {
    v1 := []uint32{123, 0}
    v2 := make(map[int]int)
    v3 := make(map[string]int)
    var v4 int16 = -935
    _ = len(v1)
    v5 := GenS1{F0: "foo bar", F1: -960}
    v4 = (v4 | v4)
    v4 = (v4 ^ v4)
    for fi6 := 0; fi6 < 17; fi6++ {
        var v7 float64 = 0.1
        var v8 string = "\n\t"
        var v9 uint = 652
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 uint8 = 78
        _ = fi10
        _ = v11
    }
    var v12 float32 = 1.5
    if (0 >= 18446744073709551615) {
        v12 = v12
        v4 = (v4 * (v4 + v4))
        v12 -= (float32(1e10) - v12)
    } else {
        _ = len(v1)
        fmt.Println(v12)
    }
    if (128 >= 1) {
        v4 = -208
        fmt.Println(v4, v12)
    } else {
        fmt.Println(v4)
        v13 := GenS2{F0: 42, F1: 2, F2: "\n\t"}
        _ = v13
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        v15 := GenS2{F0: 0, F1: 42, F2: "a"}
        var v16 bool = true
        _ = fi14
        _ = v15
        _ = v16
    }
    fmt.Println(v12)
    if (37 <= 1) {
        v12 += (v12 / 1.0)
        if (0 > 2147483647) {
            for fi17 := 0; fi17 < 4; fi17++ {
                v18 := GenS2{F0: 2, F1: 127, F2: ""}
                v12 -= v12
                fmt.Println(v4)
                _ = fi17
                _ = v18
            }
        }
    }
    if (60 == 2) {
        switch ("" + "\n\t") {
        case "\n\t":
            v19 := GenS1{F0: "a", F1: 9223372036854775807}
            _ = v19
        }
    } else {
        var v20 int32 = -2147483648
        switch "hello" {
        case "a":
            for fi21 := 0; fi21 < 9; fi21++ {
                v22 := GenS1{F0: "a", F1: 42}
                fmt.Println(v4)
                _ = fi21
                _ = v22
            }
        }
        var v23 uint = 1
        _ = v20
        _ = v23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
