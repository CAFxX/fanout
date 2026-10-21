package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint64 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 uint32; F1 int16 }
func genFn0(p0 int16) uint {
    v6 := GenS1{F0: 67, F1: 123, F2: 0}
    p0 = p0
    var v7 int64 = -9223372036854775808
    v8 := GenS1{F0: 790, F1: -1, F2: 42}
    for fi9 := 0; fi9 < 8; fi9++ {
        if (p0 <= 42) {
            v10 := []int{fi9, (-(fi9))}
            _ = v10
        }
        fmt.Println(fi9)
        var v11 int64 = v7
        _ = fi9
        _ = v11
    }
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 int64 = 395
    v2 := make(map[uint32]int)
    var v3 int8 = -1
    v4 := []string{((("\n\t" + "") + "hello") + "\n\t"), ""}
    var v5 int64 = 668
    v12 := GenS1{F0: 4294967295, F1: v3, F2: 728}
    v13 := GenS2{F0: (!false), F1: 920}
    v3 = v3
    v5 -= v1
    for fi14 := 0; fi14 < 17; fi14++ {
        if (false || true) {
            var v15 int8 = (v3 | -128)
            _ = v15
        }
        _ = fi14
    }
    v16 := GenS1{F0: 128, F1: (v3 % (((-(v3))) | 1)), F2: 0}
    if ("foo bar" > "foo bar") {
        v5 = 361
    } else {
        var v17 float32 = 0.0
        _ = v17
    }
    _ = len(v2)
    var v18 int64 = (^(v5))
    v5 -= 9223372036854775807
    v19 := GenS3{F0: 14, F1: 32767}
    switch ("foo bar" + ("a" + "")) {
    case "foo bar":
        var v20 uint32 = 65535
        _ = v20
    case "\n\t":
        v4 = append(v4, "hello")
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (0 == 188) {
        var v21 float64 = 1.5
        _ = v21
    } else {
        fmt.Println(v5, v1, v18)
        var v22 uint8 = 249
        _ = v22
    }
    v23 := GenS3{F0: 2, F1: -32768}
    for fi24 := 0; fi24 < 15; fi24++ {
        v25 := GenS2{F0: (65535 >= 1), F1: 255}
        _ = fi24
        _ = v25
    }
    var v26 int16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
