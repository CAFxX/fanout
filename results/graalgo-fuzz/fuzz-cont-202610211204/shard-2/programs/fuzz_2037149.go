package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 int64; F1 uint64 }

func main() {
    var v1 int = 128
    v2 := make(map[string]bool)
    var v3 int16 = 221
    v4 := []uint32{2147483647, 2147483648, 2147483647}
    var v5 int64 = 42
    var v6 uint8 = 253
    var v7 int8 = 0
    fmt.Println(v3, v1, v3)
    var v8 int8 = v7
    v2["foo bar"] = false
    var v9 float32 = 1.5
    v10 := GenS1{F0: 2147483647, F1: (2 >= (v6 % ((v6) | 1)))}
    switch "foo bar" {
    case "hello":
        _ = len(v2)
    case "foo bar":
        switch "foo bar" {
        case "":
            var v11 string = "a"
            _ = v11
        case "foo bar":
            v6 = 255
        case "a":
            if (false && false) {
                v12 := v2[((("" + ("" + ("\n\t" + ("hello" + "\n\t")))) + "hello") + ("" + ("" + "")))]
                if v12 {
                    for fi13 := 0; fi13 < 14; fi13++ {
                        v8 = -31
                        _ = fi13
                    }
                } else {
                    fmt.Println(v5)
                }
                for fi14 := 0; fi14 < 5; fi14++ {
                    v15 := GenS2{F0: (v6 / (((v6 ^ v6)) | 1)), F1: 1}
                    _ = fi14
                    _ = v15
                }
                _ = v12
            }
        }
    case "":
        fmt.Println(v6, v5, v1)
    default:
        v16 := GenS3{F0: (v5 + (v5 | v5)), F1: 2}
        _ = v16
    }
    if (v3 <= (-(v3))) {
        v17 := GenS4{F0: v5, F1: 904}
        _ = v17
    } else {
        v9 -= ((v9 * 3.14159) - (v9 + v9))
        v5 += 0
        if (((1 > 4294967295) || (255 <= 888)) || true) {
            v6 -= (v6 - v6)
            var v18 int8 = v7
            _ = v18
        }
    }
    v5 = v5
    v4 = append(v4, 32)
    v6 -= 219
    if ((v6 + v6) >= 128) {
        var v19 float32 = 1.5
        _ = v19
    }
    v20 := GenS1{F0: 2147483647, F1: (243 == v5)}
    for fi21 := 0; fi21 < 12; fi21++ {
        var v22 uint8 = v6
        _ = fi21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
