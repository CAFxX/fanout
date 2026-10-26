package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int8 }
func genFn0(p0 string, p1 string, p2 string) uint32 {
    var v4 uint16 = 2
    fmt.Println(p1, v4, v4)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 4294967295
}

func main() {
    var v1 float32 = ((float32(float64(1.5)) - (1e10 / 3.14159)) + (0.0 * 3.14159))
    v2 := make(map[int]string)
    v3 := []uint32{}
    for fi5 := 0; fi5 < 18; fi5++ {
        var v6 uint16 = 360
        _ = fi5
        _ = v6
    }
    v1 = 3.14159
    v7 := GenS1{F0: 478, F1: 1, F2: -128}
    for fi8 := 0; fi8 < 17; fi8++ {
        fmt.Println(fi8, v1)
        var v9 uint64 = 42
        _ = fi8
        _ = v9
    }
    var v10 uint64 = 0
    var v11 float64 = float64((float32(0.0) / ((0.0 + (v1 - 0.0)) / v1)))
    v11 = v11
    if false {
        v11 -= v11
        v1 = (v1 - (v1 + 0.1))
        v11 = (((v11 - 0.1) - 0.1) - ((v11 / -2.25) / (float64((v1 * v1)) - v11)))
    }
    switch "hello" {
    case "hello":
        v2[42] = (("hello" + ("hello" + "")) + "a")
    default:
        v11 += v11
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v7
}
