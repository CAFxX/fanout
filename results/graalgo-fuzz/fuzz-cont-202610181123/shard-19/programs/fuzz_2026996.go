package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 bool }
type GenS2 struct { F0 string; F1 int8; F2 string }
type GenS3 struct { F0 uint16; F1 string; F2 uint16 }
type GenS4 struct { F0 uint64; F1 uint32; F2 uint16 }
type GenS5 struct { F0 int16; F1 uint }
func genFn0(p0 int16) uint {
    v4 := GenS1{F0: 337, F1: 0, F2: true}
    v5 := []int{7, 912, 127}
    _ = p0
    _ = v4
    _ = v5
    return 253
}
func genFn1() uint64 {
    if false {
        var v6 float64 = ((float64((((-2.25 + ((1e10 * (1e10 * 3.14159)) - float32((0.0 - 0.0)))) - 1e10) * 1e10)) / (((1e10 * ((3.14159 / 1.0) / (1e10 * 3.14159))) + ((1e10 * 0.1) / (0.0 - float64(((1e10 + ((0.0 * 3.14159) - 1e10)) - -2.25))))) - 3.14159)) * 1e10)
        var v7 uint16 = 1
        _ = v6
        _ = v7
    } else {
        switch "hello" {
        case "foo bar":
            switch "hello" {
            case "hello":
                var v8 string = ""
                _ = v8
            default:
                var v9 string = ("" + ("" + "a"))
                _ = v9
            }
        default:
            var v10 int8 = 0
            _ = v10
        }
        var v11 int64 = 765
        fmt.Println(v11)
        _ = v11
    }
    var v12 uint32 = 2147483648
    for fi13 := 0; fi13 < 13; fi13++ {
        var v14 int = (-(fi13))
        v12 -= v12
        for fi15 := 0; fi15 < 1; fi15++ {
            switch v14 {
            case 18:
                v16 := GenS2{F0: "a", F1: 100, F2: ""}
                _ = v16
            case 1:
                var v17 uint16 = 320
                _ = v17
            default:
                fmt.Println(v12, v12, fi13)
            }
            _ = fi15
        }
        _ = fi13
        _ = v14
    }
    if false {
        v12 -= v12
        v12 -= v12
    }
    _ = v12
    return 0
}

func main() {
    v1 := []uint32{905}
    var v2 uint16 = 0
    v3 := make(map[int]string)
    v18 := GenS3{F0: v2, F1: ("" + ""), F2: (v2 ^ 42)}
    v2 -= v2
    switch 9223372036854775807 {
    case 631:
        fmt.Println(v2)
    case 32767:
        v2 -= (v2 >> ((genFn0(465)) & 63))
    case -1:
        fmt.Println(v2)
    }
    var v19 uint32 = 2
    var v20 int32 = 1
    if false {
        var v21 uint64 = uint64(v19)
        _ = v21
    }
    if (float32(((3.14159 * 3.14159) - 3.14159)) != 3.14159) {
        v22 := GenS4{F0: uint64(v2), F1: (v19 * 2), F2: v2}
        v23 := GenS5{F0: 514, F1: 2}
        v24 := GenS4{F0: 159, F1: 640, F2: (v2 >> ((857) & 63))}
        _ = v22
        _ = v23
        _ = v24
    }
    v2 = 42
    var v25 uint64 = 435
    for fi26 := 0; fi26 < 7; fi26++ {
        var v27 int64 = 1
        fmt.Println(v27, v2)
        _ = fi26
        _ = v27
    }
    var v28 float32 = 0.0
    v28 = (v28 + 0.1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v28
    _ = v3
}
