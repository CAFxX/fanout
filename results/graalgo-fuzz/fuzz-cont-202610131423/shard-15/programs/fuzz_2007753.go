package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
func genFn0(p0 string) int64 {
    fmt.Println(p0)
    var v5 uint8 = 255
    var v6 float64 = -2.25
    _ = p0
    _ = v5
    _ = v6
    return -356
}
func genFn1() uint16 {
    v7 := []int{-96}
    var v8 float64 = float64(3.14159)
    var v9 int = -230
    for fi10 := 0; fi10 < 12; fi10++ {
        fmt.Println(v8, fi10)
        _ = fi10
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        _ = len(v7)
        _ = fi11
    }
    _ = v7
    _ = v8
    _ = v9
    return 558
}

func main() {
    var v1 float64 = 1e10
    var v2 uint32 = 42
    v3 := make(map[int]string)
    var v4 uint16 = 410
    var v12 float32 = ((0.1 + -2.25) * (float32(float64(0.1)) * 0.0))
    var v13 uint8 = 42
    for fi14 := 0; fi14 < 17; fi14++ {
        v13 = (^(v13))
        _ = fi14
    }
    _ = len(v3)
    fmt.Println(v1, v1)
    for fi15 := 0; fi15 < 2; fi15++ {
        v16 := GenS1{F0: 0, F1: 1}
        _ = fi15
        _ = v16
    }
    v1 = 1.5
    if true {
        v17 := []int{-916, 0, -2147483648}
        _ = v17
    }
    var v18 int64 = -703
    v13 = v13
    var v19 int = -130
    if false {
        var v20 float32 = 0.0
        v4 = 2
        _ = v20
    }
    fmt.Println(v2, v1, v18)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
