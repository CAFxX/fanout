package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint8; F1 int }
func genFn0(p0 uint16, p1 uint32, p2 int16) int16 {
    var v4 float32 = 0.1
    fmt.Println(v4, v4)
    var v5 uint8 = 42
    v6 := []int{-2147483648, -1, 1, 812}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return p2
}

func main() {
    var v1 int8 = -95
    var v2 int32 = 430
    v3 := make(map[string]bool)
    v2 -= -1
    v7 := v3["foo bar"]
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 uint = 257
        v10 := GenS1{F0: 1, F1: 0}
        if false {
            switch 127 {
            case 32768:
                var v11 int16 = 1
                _ = v11
            }
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    v12 := GenS2{F0: 127, F1: 2147483647}
    v7 = (v7 || v7)
    fmt.Println(v2, v2)
    fmt.Println(v7, v2)
    var v13 float32 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v7
}
