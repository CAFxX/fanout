package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 bool; F1 int }
func genFn0(p0 uint32) int8 {
    v4 := []int{42, 1, 1}
    v5 := GenS1{F0: 128, F1: 221}
    _ = p0
    _ = v4
    _ = v5
    return -11
}
func genFn1(p0 int, p1 int32) bool {
    v6 := GenS1{F0: 0, F1: 42}
    var v7 uint16 = 1
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (68 > 1)
}

func main() {
    var v1 float64 = 0.0
    v2 := []int{32768}
    var v3 float32 = (0.0 / (float32(1e10) / 0.1))
    v8 := GenS1{F0: 255, F1: 42}
    if (!true) {
        v9 := GenS1{F0: 42, F1: 85}
        v1 = v1
        _ = v9
    }
    v1 = (1e10 * v1)
    switch "a" {
    case "\n\t":
        v3 = float32(float64(v3))
    case "hello":
        fmt.Println(v1, v1)
    }
    v10 := GenS2{F0: true, F1: -215}
    v1 += v1
    for fi11 := 0; fi11 < 14; fi11++ {
        if (0 == 65535) {
            v1 -= v1
            v3 = v3
        } else {
            v1 = v1
            fmt.Println(v1)
            var v12 int8 = genFn0(65535)
            _ = v12
        }
        _ = len(v2)
        _ = fi11
    }
    var v13 string = ("foo bar" + "\n\t")
    if len(v2) > 0 {
        v2[0] = -467
    }
    v1 = (v1 - ((((3.14159 / (1.5 * (1.5 - (v1 + -2.25)))) - v1) - v1) + (v1 + (v1 * (v1 - v1)))))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v8
}
