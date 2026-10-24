package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int16 }
func genFn0(p0 bool) uint16 {
    var v7 int64 = 639
    var v8 float64 = float64((((1.5 * 1e10) / (3.14159 * 3.14159)) - (3.14159 * -2.25)))
    _ = p0
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 float32 = (3.14159 + 0.1)
    var v2 uint64 = 292
    var v3 int16 = 32767
    v4 := []string{}
    var v5 int8 = 15
    v6 := make(map[uint32]bool)
    var v9 float32 = float32(((((1.5 - 1.5) - ((((0.0 - (1.5 * 0.0)) * 3.14159) + (1e10 + 1e10)) - 0.0)) / (0.1 - 0.0)) / 1.0))
    fmt.Println(v9, v3, v9)
    var v10 float32 = v1
    fmt.Println(v5)
    v11 := GenS1{F0: 127, F1: 662, F2: (^(v3))}
    switch "" {
    case "foo bar":
        v1 += (v10 * ((v10 * v9) + float32((1.5 - 3.14159))))
    case "":
        fmt.Println(v1, v1)
    default:
        v1 -= float32((0.0 - (0.0 - ((3.14159 - ((3.14159 / 1e10) + (-2.25 * 3.14159))) - 1e10))))
    }
    switch "\n\t" {
    case "a":
        var v12 uint64 = v2
        _ = v12
    }
    v1 += (((-2.25 / ((1.5 / v9) - v10)) / v9) + float32(-2.25))
    v9 = float32(0.0)
    for fi13 := 0; fi13 < 11; fi13++ {
        for fi14 := 0; fi14 < 11; fi14++ {
            v1 = v10
            _ = fi14
        }
        v6[65535] = false
        _ = fi13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
