package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 int }

func main() {
    var v1 uint32 = 983
    var v2 float64 = 3.14159
    var v3 int16 = -255
    v4 := make(map[uint32]string)
    var v5 string = ("foo bar" + "")
    v6 := make(map[string]bool)
    fmt.Println(v3, v3)
    v3 -= v3
    var v7 uint16 = 1
    var v8 uint64 = 42
    for fi9 := 0; fi9 < 3; fi9++ {
        if false {
            var v10 int16 = int16(v8)
            _ = v10
        } else {
            var v11 int = fi9
            _ = v11
        }
        v5 = "foo bar"
        _ = fi9
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        fmt.Println(v2, fi12)
        _ = fi12
    }
    v8 += (v8 % ((1) | 1))
    v13 := GenS1{F0: v5, F1: v5, F2: -5}
    v5 = "a"
    v2 += -2.25
    switch 2147483647 {
    case 127:
        if (2 == 127) {
            var v14 int8 = 100
            _ = v14
        } else {
            _ = len(v4)
            var v15 uint64 = v8
            _ = len(v6)
            _ = v15
        }
    default:
        var v16 float32 = 0.0
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
