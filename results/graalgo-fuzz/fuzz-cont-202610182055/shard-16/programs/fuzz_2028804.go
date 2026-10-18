package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{4294967295, 348, 4294967295, 32768}
    var v3 float64 = 1e10
    var v4 float32 = 3.14159
    fmt.Println(v4)
    v3 = v3
    v4 += v4
    var v5 float32 = float32((v3 / v3))
    v5 -= 1e10
    fmt.Println(v5, v5, v5)
    switch 127 {
    case 1:
        switch 128 {
        case 42:
            v3 += (v3 + v3)
        default:
            var v6 uint16 = 0
            _ = v6
        }
    case 2147483647:
        v4 = v5
    }
    v7 := GenS1{F0: 42, F1: -32768}
    if (903 >= 65535) {
        v8 := GenS1{F0: -2147483648, F1: -662}
        _ = v8
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        for fi10 := 0; fi10 < 18; fi10++ {
            var v11 int16 = 32767
            _ = fi10
            _ = v11
        }
        switch (fi9 & (fi9 * fi9)) {
        case -1:
            var v12 uint = 2
            _ = v12
        case 32767:
            v3 = (-2.25 * v3)
        case 42:
            v4 = (v5 * (v4 * 0.1))
        }
        _ = fi9
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
