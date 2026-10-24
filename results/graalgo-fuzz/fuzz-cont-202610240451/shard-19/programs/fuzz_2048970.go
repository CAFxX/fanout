package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int8 }

func main() {
    var v1 uint64 = 162
    v2 := []string{"foo bar"}
    var v3 int = 4294967295
    v4 := make(map[string]bool)
    var v5 int8 = 0
    var v6 uint8 = 101
    var v7 int64 = 42
    switch (v3 * (-(v3))) {
    case 605:
        v6 -= 195
    case -1:
        _ = len(v4)
    default:
        if (!(159 == v1)) {
            v6 += v6
            var v8 int16 = 977
            _ = v8
        }
    }
    if (2 >= 0) {
        v7 = v7
        v5 -= (v5 >> ((2) & 63))
    } else {
        for fi9 := 0; fi9 < 4; fi9++ {
            v10 := GenS1{F0: 2147483647, F1: -419, F2: -128}
            v11 := GenS1{F0: 8, F1: -1, F2: v5}
            if (!((v5 / ((v5) | 1)) == (v5 - v5))) {
                fmt.Println(v7)
                var v12 float32 = (float32(-2.25) - 1.5)
                var v13 int16 = int16(v1)
                _ = v12
                _ = v13
            }
            _ = fi9
            _ = v10
            _ = v11
        }
        v3 -= (v3 >> ((940) & 63))
        switch (v3 ^ (v3 + (-(v3)))) {
        case 42:
            v14 := GenS1{F0: 42, F1: 2147483647, F2: v5}
            _ = v14
        }
    }
    v3 += (-(v3))
    v15 := GenS1{F0: 42, F1: -1, F2: v5}
    fmt.Println(v6, v3, v7)
    var v16 string = ("foo bar" + "a")
    var v17 float64 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
