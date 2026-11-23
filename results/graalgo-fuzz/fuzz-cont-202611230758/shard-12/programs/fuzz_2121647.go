package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int32 }
type GenS2 struct { F0 string; F1 int32; F2 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int8; F1 bool; F2 string }
type GenS5 struct { F0 int64; F1 int64; F2 uint16 }
type GenS6 struct { F0 int16; F1 bool }
func genFn0(p0 int32) int8 {
    fmt.Println(p0)
    v6 := GenS1{F0: 32768, F1: p0, F2: (p0 + 2147483647)}
    _ = p0
    _ = v6
    return 127
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 float32 = -2.25
    var v3 uint64 = 42
    var v4 int64 = (^(v1))
    v5 := make(map[uint32]int)
    v7 := GenS2{F0: "\n\t", F1: int32(v3), F2: 255}
    switch ("" + "foo bar") {
    case "hello":
        if ((!true) && true) {
            fmt.Println(v3)
        } else {
            var v8 uint32 = 2147483648
            v3 = v3
            for fi9 := 0; fi9 < 12; fi9++ {
                v1 -= (v4 & -1)
                v10 := GenS2{F0: "\n\t", F1: -1, F2: 2}
                _ = len(v5)
                _ = fi9
                _ = v10
            }
            _ = v8
        }
    case "":
        v4 = v4
    default:
        for fi11 := 0; fi11 < 15; fi11++ {
            var v12 uint16 = 2
            _ = fi11
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        if ((true || false) && true) {
            var v14 float64 = 3.14159
            v15 := GenS3{F0: "\n\t"}
            var v16 int = fi13
            _ = v14
            _ = v15
            _ = v16
        } else {
            switch fi13 {
            case -900:
                var v17 string = "\n\t"
                _ = v17
            case 128:
                var v18 uint32 = 2147483647
                _ = v18
            }
            var v19 int32 = 86
            _ = v19
        }
        _ = fi13
    }
    var v20 uint32 = 65535
    if (v2 > 0.0) {
        var v21 int8 = 100
        _ = v21
    }
    v4 = (v1 % ((42) | 1))
    var v22 uint16 = 391
    switch "" {
    case "foo bar":
        v4 = (-(v4))
    }
    if false {
        v23 := GenS4{F0: 1, F1: false, F2: (("a" + "a") + ("a" + ""))}
        _ = v23
    }
    for fi24 := 0; fi24 < 1; fi24++ {
        if (int16(v3) >= 1000) {
            var v25 uint8 = 42
            if (v25 != v25) {
                if ((!true) && false) {
                    switch "foo bar" {
                    case "foo bar":
                        fmt.Println(v4)
                    case "a":
                        fmt.Println(v22, v1, v20)
                    }
                    var v26 float32 = 0.0
                    _ = v26
                }
            }
            _ = v25
        }
        v22 -= 255
        v27 := GenS5{F0: v4, F1: (v1 % ((v4) | 1)), F2: v22}
        _ = fi24
        _ = v27
    }
    v28 := v5[v20]
    v29 := GenS6{F0: 696, F1: true}
    v30 := GenS2{F0: "hello", F1: -2147483648, F2: 987}
    var v31 string = "hello"
    v4 -= (v1 - (v1 >> ((255) & 63)))
    v32 := GenS5{F0: 42, F1: v1, F2: v22}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v7
}
