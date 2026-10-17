package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 int16; F1 int16; F2 uint32 }
func genFn0() string {
    var v6 int32 = 0
    switch "foo bar" {
    case "foo bar":
        v7 := []int{-368, int(v6), 32768}
        _ = v7
    case "":
        switch "a" {
        case "a":
            v6 += (v6 << ((138) & 63))
        case "foo bar":
            var v8 uint64 = 803
            _ = v8
        }
    default:
        var v9 int = 82
        _ = v9
    }
    v10 := []int{-2147483648, -787, 917, 128}
    _ = v10
    _ = v6
    return "a"
}

func main() {
    var v1 int16 = 0
    v2 := []string{"a", "hello", "hello", "foo bar"}
    var v3 int = 32768
    v4 := make(map[string]string)
    var v5 uint16 = 32767
    var v11 bool = (false || (!false))
    var v12 string = (((("a" + "foo bar") + "a") + "foo bar") + ("\n\t" + "\n\t"))
    var v13 int16 = (v1 / (((v1 ^ (-(v1)))) | 1))
    for fi14 := 0; fi14 < 15; fi14++ {
        for fi15 := 0; fi15 < 17; fi15++ {
            v3 = (v3 | 32768)
            v16 := GenS1{F0: 1, F1: uint8(v5)}
            v1 -= -32768
            _ = fi15
            _ = v16
        }
        v2 = append(v2, v12)
        var v17 int8 = 127
        _ = fi14
        _ = v17
    }
    switch v3 {
    case 42:
        v18 := GenS2{F0: (v3 - v3), F1: ""}
        _ = v18
    case 801:
        fmt.Println(v12, v3, v5)
    case 924:
        v5 += (v5 % ((v5) | 1))
    default:
        v19 := GenS3{F0: (v1 / (((v1 / (((-(v13))) | 1))) | 1)), F1: v13, F2: 632}
        _ = v19
    }
    v1 = -185
    fmt.Println(v5)
    for fi20 := 0; fi20 < 15; fi20++ {
        if (0.0 == -2.25) {
            v13 = 42
            var v21 uint16 = (v5 + v5)
            v2 = append(v2, v12)
            _ = v21
        } else {
            var v22 string = "foo bar"
            delete(v4, (v22 + v12))
            _ = v22
        }
        var v23 uint64 = 2
        v12 = (((v12 + v12) + (v12 + (v12 + "foo bar"))) + ((v12 + v12) + v12))
        _ = fi20
        _ = v23
    }
    _ = len(v4)
    v2 = append(v2, (v12 + "hello"))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
