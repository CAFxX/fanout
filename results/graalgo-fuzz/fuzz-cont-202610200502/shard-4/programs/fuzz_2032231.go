package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 int32; F1 string; F2 uint16 }

func main() {
    var v1 int64 = -1
    v2 := make(map[string]int)
    v3 := make(map[string]bool)
    if (42 != -490) {
        fmt.Println(v1)
        v4 := GenS1{F0: 1, F1: 2147483648}
        _ = v4
    }
    if true {
        var v5 uint16 = 1
        var v6 string = "foo bar"
        _ = v5
        _ = v6
    } else {
        fmt.Println(v1)
        for fi7 := 0; fi7 < 16; fi7++ {
            v8 := v3[("\n\t" + "\n\t")]
            _ = fi7
            _ = v8
        }
        v9 := GenS1{F0: -40, F1: 902}
        _ = v9
    }
    var v10 uint32 = 383
    fmt.Println(v10, v1)
    switch "foo bar" {
    case "hello":
        v1 = (v1 >> ((0) & 63))
    case "\n\t":
        fmt.Println(v10, v10)
    default:
        var v11 float64 = float64(-2.25)
        _ = v11
    }
    fmt.Println(v10)
    v10 = (v10 << ((uint(v10)) & 63))
    v10 -= (^(v10))
    v12 := GenS2{F0: 0, F1: (("a" + "\n\t") + "\n\t"), F2: 52}
    if true {
        for fi13 := 0; fi13 < 6; fi13++ {
            v14 := GenS1{F0: -1, F1: v10}
            _ = fi13
            _ = v14
        }
    } else {
        for fi15 := 0; fi15 < 20; fi15++ {
            v16 := GenS1{F0: 0, F1: (v10 >> ((77) & 63))}
            v1 += v1
            _ = fi15
            _ = v16
        }
    }
    v10 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
}
