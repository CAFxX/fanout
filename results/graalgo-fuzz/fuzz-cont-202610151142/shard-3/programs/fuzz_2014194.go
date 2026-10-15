package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }

func main() {
    var v1 int8 = 94
    v2 := make(map[int]string)
    v3 := []int{}
    var v4 int64 = -794
    v5 := make(map[string]uint32)
    var v6 float64 = 0.1
    var v7 float32 = (3.14159 * ((0.1 - 3.14159) * (((float32(-2.25) * -2.25) - 1.5) - (0.1 * 1.5))))
    v1 -= (-(v1))
    v7 -= float32(v6)
    switch ("\n\t" + "foo bar") {
    case "a":
        v1 = v1
    default:
        fmt.Println(v6)
    }
    v4 -= -9223372036854775808
    var v8 int8 = -128
    var v9 bool = false
    fmt.Println(v6, v9)
    v8 = 127
    switch "a" {
    case "\n\t":
        _ = len(v2)
    }
    fmt.Println(v8)
    v4 += 0
    if len(v3) > 0 {
        v3[0] = 128
    }
    switch ("hello" + "hello") {
    case "a":
        if v9 {
            _ = len(v5)
            for fi10 := 0; fi10 < 20; fi10++ {
                v11 := GenS1{F0: 42, F1: v4}
                _ = fi10
                _ = v11
            }
        } else {
            switch "" {
            case "":
                if len(v3) > 0 {
                    v3[0] = -683
                }
            case "a":
                v9 = v9
            case "foo bar":
                fmt.Println(v8, v8)
            }
        }
    case "\n\t":
        v1 += v8
    case "hello":
        var v12 int32 = 940
        _ = v12
    default:
        var v13 uint8 = 2
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
