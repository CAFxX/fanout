package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }

func main() {
    var v1 bool = (!true)
    var v2 string = "hello"
    v3 := make(map[uint32]int)
    var v4 int16 = 32767
    var v5 float64 = 0.1
    v4 = -79
    var v6 uint16 = 835
    fmt.Println(v4)
    if (127 != 127) {
        v3[2] = -2147483648
        var v7 int8 = 9
        fmt.Println(v1, v5, v1)
        _ = v7
    }
    v8 := []int{}
    var v9 int8 = 127
    v5 = ((0.1 * (v5 + 0.1)) / v5)
    for fi10 := 0; fi10 < 9; fi10++ {
        if v1 {
            var v11 string = "\n\t"
            _ = v11
        }
        var v12 int16 = v4
        v6 += v6
        _ = fi10
        _ = v12
    }
    if (-128 != (v9 - v9)) {
        var v13 uint = 876
        if false {
            var v14 float32 = float32(v5)
            _ = v14
        }
        _ = v13
    } else {
        v2 = v2
        v15 := GenS1{F0: v1, F1: 661}
        _ = v15
    }
    var v16 int8 = -65
    fmt.Println(v1, v16, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
