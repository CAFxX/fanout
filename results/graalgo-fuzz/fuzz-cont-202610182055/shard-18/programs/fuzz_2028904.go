package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 int32 }

func main() {
    var v1 float32 = 1.5
    v2 := []uint32{554}
    v3 := make(map[int]string)
    v4 := []string{}
    var v5 int8 = 100
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 float64 = 1.5
        for fi8 := 0; fi8 < 4; fi8++ {
            switch "" {
            case "\n\t":
                fmt.Println(fi8, v5)
            case "foo bar":
                v9 := v3[fi6]
                _ = v9
            case "hello":
                v1 = v1
            }
            fmt.Println(fi6)
            var v10 int16 = -32768
            _ = fi8
            _ = v10
        }
        v7 += float64((v1 * v1))
        _ = fi6
        _ = v7
    }
    v3[375] = ("foo bar" + "a")
    v5 -= v5
    for fi11 := 0; fi11 < 11; fi11++ {
        v1 -= (1.5 - ((v1 / v1) + float32(((-2.25 / 1.0) * 3.14159))))
        var v12 float64 = ((float64((float32((1.5 / 1e10)) + v1)) * 3.14159) + -2.25)
        _ = fi11
        _ = v12
    }
    _ = len(v3)
    v13 := v3[570]
    v14 := GenS1{F0: 1000, F1: -942, F2: 47}
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
