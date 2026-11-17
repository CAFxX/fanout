package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]int)
    var v3 float64 = 0.1
    v3 = 0.1
    if false {
        var v4 int16 = -261
        _ = v4
    } else {
        v5 := GenS1{F0: 646, F1: 68, F2: 2}
        v3 = 0.1
        _ = v5
    }
    fmt.Println(v3)
    v6 := v1[(("hello" + "") + ("foo bar" + "\n\t"))]
    v7 := GenS2{F0: 697}
    v3 = v3
    if (!(!v6)) {
        v8 := GenS1{F0: 2, F1: 1, F2: 662}
        _ = v8
    }
    v9 := GenS3{F0: 127}
    for fi10 := 0; fi10 < 6; fi10++ {
        fmt.Println(v6, v6, fi10)
        v6 = (57 == 0)
        _ = fi10
    }
    v3 -= v3
    if (!false) {
        var v11 uint16 = 751
        v12 := GenS3{F0: 42}
        _ = v11
        _ = v12
    } else {
        var v13 float32 = ((1e10 * 0.1) + (0.1 / 3.14159))
        for fi14 := 0; fi14 < 13; fi14++ {
            fmt.Println(v6, fi14)
            _ = fi14
        }
        var v15 float64 = (0.0 - v3)
        _ = v13
        _ = v15
    }
    v3 = (float64(1e10) + (0.1 / (v3 * v3)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v9
}
