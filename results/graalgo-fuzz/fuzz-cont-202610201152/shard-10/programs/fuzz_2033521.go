package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }

func main() {
    var v1 uint8 = 42
    v2 := make(map[int]uint32)
    var v3 int64 = -1
    v4 := []int{2147483647, -2147483648, int(v1)}
    v5 := make(map[string]bool)
    v1 -= (v1 % ((102) | 1))
    v1 = 2
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v1 += 128
    v3 = -580
    for fi6 := 0; fi6 < 14; fi6++ {
        if true {
            v7 := GenS1{F0: ("hello" + "\n\t"), F1: -60}
            delete(v2, 32767)
            _ = v7
        }
        _ = fi6
    }
    var v8 float32 = ((-2.25 + -2.25) + 0.1)
    switch 42 {
    case -635:
        for fi9 := 0; fi9 < 15; fi9++ {
            v1 = v1
            v1 += 255
            _ = fi9
        }
    default:
        fmt.Println(v3)
    }
    var v10 uint32 = 1
    var v11 int16 = 42
    v12 := GenS1{F0: "\n\t", F1: -128}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
