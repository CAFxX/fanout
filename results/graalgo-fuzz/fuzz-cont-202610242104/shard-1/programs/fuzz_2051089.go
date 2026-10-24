package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int32 }
type GenS2 struct { F0 int32; F1 int8 }
func genFn0(p0 int32, p1 uint64, p2 uint8) int64 {
    switch "foo bar" {
    case "hello":
        p0 -= (p0 & (^(p0)))
    case "foo bar":
        for fi4 := 0; fi4 < 14; fi4++ {
            var v5 uint32 = 989
            fmt.Println(v5, p0, fi4)
            _ = fi4
            _ = v5
        }
    default:
        p2 += p2
    }
    var v6 uint64 = (p1 * (p1 / (((p1 ^ p1)) | 1)))
    switch 788 {
    case 747:
        var v7 int = 429
        _ = v7
    case 618:
        v8 := GenS1{F0: -887, F1: 127, F2: -2147483648}
        _ = v8
    case 127:
        switch (("hello" + "hello") + "foo bar") {
        case "foo bar":
            var v9 float64 = -2.25
            _ = v9
        default:
            var v10 uint = 42
            _ = v10
        }
    default:
        var v11 int = 1
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -463
}

func main() {
    var v1 int64 = 477
    v2 := []uint32{2, 2147483648}
    var v3 float64 = (((0.0 + ((0.0 + 0.0) - 1e10)) * ((float64((3.14159 / 3.14159)) - (1e10 - float64(3.14159))) - (-2.25 / 0.1))) / (0.1 / 1.0))
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    v2 = append(v2, 65535)
    v2 = append(v2, 190)
    v2 = append(v2, 491)
    var v12 bool = (621 > 65535)
    var v13 uint = 4294967295
    v14 := GenS2{F0: -615, F1: -128}
    fmt.Println(v13, v1)
    v15 := GenS1{F0: -29, F1: -1, F2: -1}
    var v16 int64 = (-(v1))
    var v17 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
