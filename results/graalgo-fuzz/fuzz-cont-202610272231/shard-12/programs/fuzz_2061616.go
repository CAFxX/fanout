package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int32 }
func genFn0() int8 {
    var v7 int = 127
    if true {
        switch v7 {
        case 42:
            v8 := []int{32768, -2147483648}
            _ = v8
        }
        v7 = v7
        fmt.Println(v7)
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 float64 = 0.1
        v11 := GenS1{F0: 0, F1: 0, F2: 0}
        var v12 int16 = 1000
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    var v13 bool = false
    _ = v13
    _ = v7
    return -128
}

func main() {
    var v1 uint8 = 255
    var v2 float64 = -2.25
    var v3 float64 = (1e10 * 0.0)
    var v4 int64 = -869
    v5 := make(map[int]bool)
    v6 := make(map[string]uint32)
    var v14 uint = 4294967295
    var v15 int64 = int64(v14)
    var v16 bool = (!(true && false))
    if true {
        v17 := GenS1{F0: v14, F1: 20, F2: -787}
        _ = v17
    }
    var v18 float32 = (-2.25 + (3.14159 * -2.25))
    v19 := GenS1{F0: (v14 ^ 551), F1: 127, F2: 0}
    fmt.Println(v2, v2, v14)
    v4 -= v15
    fmt.Println(v2, v1, v2)
    v20 := []int{int(v15)}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
