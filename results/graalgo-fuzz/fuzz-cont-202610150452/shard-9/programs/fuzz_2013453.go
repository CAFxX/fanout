package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint16 }

func main() {
    v1 := []uint32{729, 65535, 334}
    v2 := make(map[int]string)
    v3 := make(map[int]string)
    var v4 string = ""
    v5 := make(map[uint32]bool)
    v4 = "\n\t"
    var v6 float32 = (1e10 + -2.25)
    for fi7 := 0; fi7 < 14; fi7++ {
        v6 = ((v6 / v6) - 3.14159)
        _ = fi7
    }
    _ = len(v1)
    if len(v1) > 0 {
        v1[0] = 197
    }
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 int32 = 0
        switch v4 {
        case "a":
            fmt.Println(v9, v9, v6)
        default:
            v10 := GenS1{F0: -599, F1: v4, F2: 65535}
            _ = v10
        }
        _ = fi8
        _ = v9
    }
    if len(v1) > 0 {
        v1[0] = 65535
    }
    var v11 uint = 42
    var v12 bool = true
    v6 -= ((0.1 * 1e10) + float32(0.1))
    fmt.Println(v6, v11, v6)
    if v12 {
        for fi13 := 0; fi13 < 5; fi13++ {
            var v14 uint16 = 0
            for fi15 := 0; fi15 < 13; fi15++ {
                if v12 {
                    v16 := GenS1{F0: fi15, F1: v4, F2: v14}
                    _ = v16
                } else {
                    switch ((v4 + v4) + v4) {
                    case "hello":
                        v3[fi15] = "foo bar"
                    case "\n\t":
                        v11 = 481
                    default:
                        v4 = v4
                    }
                    var v17 float64 = (1.5 * (0.0 * 1e10))
                    _ = v17
                }
                _ = fi15
            }
            _ = fi13
            _ = v14
        }
    } else {
        if false {
            delete(v5, 331)
        }
        var v18 float64 = (0.0 + 1e10)
        _ = v18
    }
    v19 := GenS1{F0: 9223372036854775807, F1: ((((("hello" + "a") + v4) + v4) + "") + "foo bar"), F2: 612}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
