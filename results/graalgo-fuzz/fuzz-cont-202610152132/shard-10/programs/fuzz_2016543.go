package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }

func main() {
    var v1 uint8 = 127
    var v2 string = "\n\t"
    v3 := make(map[string]int)
    v4 := make(map[string]bool)
    v5 := make(map[uint32]string)
    v6 := make(map[uint32]int)
    v2 = "a"
    v7 := GenS1{F0: true, F1: true}
    var v8 float32 = 3.14159
    v9 := []int{42, 42, -313}
    if (-9223372036854775808 <= 0) {
        var v10 float64 = (3.14159 + 0.1)
        if false {
            v8 -= -2.25
            v11 := GenS1{F0: false, F1: true}
            v12 := GenS1{F0: (-1 > -1), F1: false}
            _ = v11
            _ = v12
        }
        v6[42] = 4294967295
        _ = v10
    } else {
        v8 -= 0.1
        v8 = float32(float64(v8))
        v1 -= v1
    }
    var v13 uint = 255
    if len(v9) > 0 {
        v9[0] = 127
    }
    var v14 uint64 = 794
    for fi15 := 0; fi15 < 17; fi15++ {
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        _ = fi15
    }
    fmt.Println(v8)
    v3[(v2 + "a")] = 42
    var v16 bool = true
    v14 += v14
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
