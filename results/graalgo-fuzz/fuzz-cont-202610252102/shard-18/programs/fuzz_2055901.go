package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }

func main() {
    v1 := []string{}
    var v2 float32 = 0.0
    v3 := []int{2147483647, -593}
    _ = len(v3)
    v2 -= float32(((0.1 * float64(1e10)) - 1.5))
    var v4 int = 793
    for fi5 := 0; fi5 < 10; fi5++ {
        for fi6 := 0; fi6 < 10; fi6++ {
            v7 := GenS1{F0: 9, F1: 1}
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    v8 := GenS1{F0: 255, F1: -1}
    v4 += v4
    for fi9 := 0; fi9 < 1; fi9++ {
        v10 := GenS1{F0: 0, F1: -580}
        v1 = append(v1, "foo bar")
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: 32767, F1: 0}
    v12 := GenS1{F0: 255, F1: 219}
    switch (v4 >> ((0) & 63)) {
    case 0:
        fmt.Println(v4, v4)
    case 1:
        var v13 float32 = float32((((0.1 - 3.14159) * 1.5) * (1e10 / -2.25)))
        _ = v13
    }
    v4 = (v4 * 127)
    v14 := GenS1{F0: 0, F1: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
