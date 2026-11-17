package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint }

func main() {
    var v1 float32 = (float32(0.1) * -2.25)
    var v2 uint64 = 2
    v3 := []int{-20, 1}
    var v4 int16 = 1
    var v5 uint8 = 1
    _ = len(v3)
    v3 = append(v3, 1)
    fmt.Println(v2)
    var v6 bool = (-1 < -783)
    var v7 uint8 = uint8(v2)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if v6 {
        for fi8 := 0; fi8 < 8; fi8++ {
            if len(v3) > 0 {
                v3[0] = -9223372036854775808
            }
            switch ("a" + ("" + "")) {
            case "":
                fmt.Println(v7, v2)
            }
            _ = fi8
        }
        var v9 bool = (v6 || v6)
        _ = v9
    }
    v1 -= float32(-2.25)
    v10 := GenS1{F0: 42, F1: v2, F2: 595}
    v11 := GenS1{F0: -81, F1: v2, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
