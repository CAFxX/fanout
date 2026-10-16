package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint32, p1 string) int8 {
    var v6 uint = 4294967295
    var v7 int8 = 1
    v8 := []int{4294967295, 127}
    if (p0 < (^(p0))) {
        if ((^(p0)) > (p0 >> ((v6) & 63))) {
            for fi9 := 0; fi9 < 9; fi9++ {
                v10 := GenS1{F0: 42, F1: 11}
                _ = fi9
                _ = v10
            }
        } else {
            var v11 uint32 = (p0 / ((4294967295) | 1))
            _ = v11
        }
    }
    v7 += v7
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return v7
}

func main() {
    v1 := []string{}
    v2 := make(map[string]string)
    var v3 float32 = ((0.1 + float32(-2.25)) / ((((3.14159 / 1e10) * float32(1e10)) / 3.14159) * (0.0 - 3.14159)))
    var v4 float64 = 1.5
    var v5 int16 = 0
    v12 := GenS2{F0: 127}
    for fi13 := 0; fi13 < 1; fi13++ {
        for fi14 := 0; fi14 < 6; fi14++ {
            delete(v2, "a")
            _ = fi14
        }
        _ = len(v1)
        _ = fi13
    }
    var v15 bool = false
    if v15 {
        v4 = float64(0.1)
    }
    for fi16 := 0; fi16 < 13; fi16++ {
        var v17 uint64 = 243
        fmt.Println(v17, v3)
        var v18 int64 = 0
        _ = fi16
        _ = v17
        _ = v18
    }
    v15 = v15
    var v19 int = 128
    for fi20 := 0; fi20 < 10; fi20++ {
        v19 += fi20
        v3 = v3
        switch "foo bar" {
        case "foo bar":
            var v21 float64 = float64((((v3 * (((3.14159 + v3) * v3) / 1.0)) / (v3 + v3)) * 1.5))
            _ = v21
        default:
            if (!v15) {
                v19 += (^(v19))
            }
        }
        _ = fi20
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
