package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 string }

func main() {
    var v1 uint16 = 353
    var v2 int8 = 113
    v3 := make(map[string]string)
    var v4 int32 = -488
    v5 := make(map[string]int)
    v6 := GenS1{F0: 1, F1: 0}
    v7 := []int{-32, 2147483647, 127}
    fmt.Println(v2, v1, v1)
    if (uint(v1) != 255) {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        if (42 == 42) {
            switch 4294967295 {
            case -2147483648:
                var v8 uint32 = 4294967295
                _ = v8
            case 127:
                v2 -= 100
            case 128:
                fmt.Println(v4, v4, v1)
            default:
                v2 -= int8(v4)
            }
            v2 += v2
        }
    } else {
        fmt.Println(v2, v4, v2)
    }
    var v9 uint64 = 2
    var v10 uint16 = (v1 << ((1) & 63))
    if (!(!false)) {
        v11 := GenS1{F0: (v10 * v10), F1: v4}
        _ = v11
    }
    v4 = v4
    v10 = (^(v1))
    v12 := GenS2{F0: "foo bar"}
    for fi13 := 0; fi13 < 15; fi13++ {
        fmt.Println(v1)
        v1 = v10
        fmt.Println(v9)
        _ = fi13
    }
    fmt.Println(v10)
    for fi14 := 0; fi14 < 13; fi14++ {
        v9 -= (v9 + v9)
        var v15 int16 = 55
        switch (("foo bar" + "hello") + "a") {
        case "a":
            v9 = v9
        case "hello":
            v4 = (v4 + (v4 | v4))
        }
        _ = fi14
        _ = v15
    }
    v4 -= v4
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
