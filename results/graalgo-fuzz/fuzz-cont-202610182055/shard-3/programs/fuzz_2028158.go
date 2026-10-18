package main

import (
    "fmt"
)


func main() {
    var v1 uint8 = 127
    var v2 uint16 = 246
    var v3 int32 = 450
    var v4 int = -2147483648
    var v5 float64 = (0.1 / 1.0)
    for fi6 := 0; fi6 < 15; fi6++ {
        v7 := []int{}
        fmt.Println(v3)
        if len(v7) > 0 {
            v7[0] = 127
        }
        _ = fi6
        _ = v7
    }
    v3 = (v3 + 689)
    for fi8 := 0; fi8 < 15; fi8++ {
        v3 -= -747
        _ = fi8
    }
    v9 := []int{v4, (^(v4)), (v4 & (v4 | 0)), (-(v4))}
    var v10 string = "a"
    _ = len(v9)
    fmt.Println(v4, v5)
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    var v11 int = 2147483647
    fmt.Println(v3, v10)
    var v12 float32 = (float32((v5 - v5)) - 1.5)
    _ = len(v9)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
