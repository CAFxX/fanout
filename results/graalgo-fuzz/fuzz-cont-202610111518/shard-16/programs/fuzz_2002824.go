package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint8 }

func main() {
    var v1 int16 = 0
    var v2 float64 = (3.14159 / -2.25)
    var v3 float32 = float32((((v2 / ((v2 * 1e10) * (v2 + (0.0 / (float64(float32((3.14159 / v2))) - v2))))) / ((float64(1.5) - (v2 + (v2 + v2))) * 1.5)) * v2))
    v4 := []int{0}
    var v5 float32 = v3
    v2 += v2
    var v6 uint32 = 2147483647
    var v7 int8 = 19
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if false {
        v3 += v5
        v3 = v3
        var v8 int32 = 1
        _ = v8
    }
    v4 = append(v4, 4294967295)
    var v9 float64 = 0.1
    v9 += ((v2 + (((v9 * v2) - v9) * 3.14159)) - 1.5)
    v10 := GenS1{F0: 2147483647, F1: -1, F2: 255}
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 int = fi11
        v9 = v9
        _ = fi11
        _ = v12
    }
    v13 := GenS1{F0: 1, F1: (v1 + (v1 / ((v1) | 1))), F2: 47}
    for fi14 := 0; fi14 < 15; fi14++ {
        fmt.Println(v6, v9, v7)
        _ = fi14
    }
    var v15 uint32 = v6
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
