package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 int64; F1 string }

func main() {
    var v1 uint16 = 2
    v2 := make(map[int]uint32)
    v3 := make(map[string]int)
    var v4 string = "foo bar"
    v5 := []string{v4, v4, v4, v4}
    var v6 uint16 = (v1 % ((255) | 1))
    var v7 int32 = -293
    if len(v5) > 0 {
        v5[0] = (("a" + (v4 + (v4 + "foo bar"))) + v4)
    }
    switch 720 {
    case 0:
        for fi8 := 0; fi8 < 10; fi8++ {
            v7 = (v7 % ((2147483647) | 1))
            v7 += (v7 % (((-(v7))) | 1))
            _ = fi8
        }
    case 9223372036854775807:
        fmt.Println(v6)
    case 42:
        var v9 bool = (!false)
        _ = v9
    }
    fmt.Println(v1)
    fmt.Println(v1)
    v5 = append(v5, (v4 + v4))
    var v10 float32 = -2.25
    v6 -= 0
    var v11 bool = (true && (false && true))
    v1 = (v1 - v1)
    v12 := GenS1{F0: 495, F1: v6}
    var v13 int64 = 1
    v14 := GenS2{F0: 252, F1: v4}
    v15 := GenS1{F0: 1, F1: (^(v1))}
    for fi16 := 0; fi16 < 3; fi16++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v17 int8 = -128
        _ = len(v2)
        _ = fi16
        _ = v17
    }
    switch v4 {
    case "a":
        var v18 int32 = v7
        _ = v18
    case "\n\t":
        var v19 bool = (4294967295 <= 32768)
        _ = v19
    default:
        var v20 uint32 = 444
        _ = v20
    }
    v11 = v11
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
