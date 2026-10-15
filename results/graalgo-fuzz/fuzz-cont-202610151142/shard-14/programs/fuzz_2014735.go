package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }

func main() {
    var v1 uint8 = 94
    var v2 int16 = 0
    var v3 float32 = (1.5 * float32(0.0))
    v4 := []int{128}
    v3 = 0.0
    if (float64(0.0) < (1e10 / ((1e10 + 1e10) * 3.14159))) {
        v1 = (v1 % (((v1 >> ((255) & 63))) | 1))
        v3 = ((0.1 / v3) + 3.14159)
    }
    fmt.Println(v1, v3, v1)
    var v5 string = ("a" + "foo bar")
    for fi6 := 0; fi6 < 10; fi6++ {
        switch (fi6 + 884) {
        case -2147483648:
            v7 := GenS1{F0: v5, F1: 0}
            _ = v7
        }
        _ = fi6
    }
    fmt.Println(v1, v3)
    fmt.Println(v3)
    switch "" {
    case "hello":
        for fi8 := 0; fi8 < 5; fi8++ {
            v2 -= v2
            _ = fi8
        }
    case "":
        v5 = v5
    default:
        var v9 float64 = ((0.1 - 3.14159) / (0.1 - 0.0))
        _ = v9
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        v11 := GenS1{F0: v5, F1: -1}
        _ = fi10
        _ = v11
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
