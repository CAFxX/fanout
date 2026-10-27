package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }

func main() {
    var v1 int32 = 42
    v2 := make(map[string]int)
    v3 := make(map[uint32]int)
    v4 := []int{32767}
    var v5 int16 = -32768
    v1 = v1
    switch -2147483648 {
    case 0:
        for fi6 := 0; fi6 < 18; fi6++ {
            fmt.Println(v1, v5)
            _ = fi6
        }
    case 32768:
        fmt.Println(v1, v1)
    }
    if (true && false) {
        v7 := GenS1{F0: v5, F1: false}
        var v8 int = 9223372036854775807
        fmt.Println(v1, v8, v8)
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: (-(v5)), F1: true}
    v5 = (v5 >> ((794) & 63))
    if (v1 <= v1) {
        switch "" {
        case "":
            v10 := GenS1{F0: (v5 | v5), F1: true}
            _ = v10
        case "hello":
            v1 -= 639
        case "foo bar":
            v11 := GenS1{F0: (v5 ^ v5), F1: (65535 > 1)}
            _ = v11
        default:
            v5 += (v5 % ((v5) | 1))
        }
    }
    for fi12 := 0; fi12 < 8; fi12++ {
        var v13 int8 = 0
        var v14 int32 = (^(v1))
        for fi15 := 0; fi15 < 20; fi15++ {
            v14 += (v14 % (((^(v1))) | 1))
            v16 := GenS1{F0: (v5 & v5), F1: true}
            fmt.Println(v5)
            _ = fi15
            _ = v16
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
