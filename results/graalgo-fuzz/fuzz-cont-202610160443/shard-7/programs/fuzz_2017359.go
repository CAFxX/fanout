package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = (0.1 + 1e10)
    var v3 float64 = (-2.25 + float64(1e10))
    v4 := make(map[uint32]bool)
    v5 := make(map[int]string)
    var v6 uint = 2
    v2 = (float64(float32(-2.25)) - -2.25)
    v7 := GenS1{F0: 1}
    var v8 uint = (v6 | v6)
    fmt.Println(v8, v6)
    v3 = v2
    for fi9 := 0; fi9 < 10; fi9++ {
        for fi10 := 0; fi10 < 4; fi10++ {
            v11 := v1[2147483648]
            v6 += (^(v6))
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    v4[555] = (v6 != v8)
    v12 := GenS1{F0: 65535}
    fmt.Println(v2, v2)
    _ = len(v1)
    for fi13 := 0; fi13 < 14; fi13++ {
        if true {
            var v14 uint = (v8 * (v6 / ((v8) | 1)))
            v8 += 784
            var v15 uint32 = 1
            _ = v14
            _ = v15
        }
        if false {
            var v16 float64 = 1e10
            v17 := v5[127]
            _ = v16
            _ = v17
        }
        _ = fi13
    }
    v2 = v2
    if true {
        v18 := []int{-2147483648}
        fmt.Println(v2, v3)
        _ = v18
    }
    var v19 int32 = 42
    fmt.Println(v6, v19, v6)
    var v20 uint = v8
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
