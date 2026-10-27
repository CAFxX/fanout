package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 int16 }

func main() {
    var v1 int32 = 1
    var v2 int64 = 1
    var v3 float64 = 1e10
    var v4 string = "a"
    var v5 float64 = (1e10 - float64((1.5 * ((float32(v3) + ((0.0 - ((1.5 / -2.25) * 3.14159)) / 1.0)) + 1e10))))
    if false {
        if false {
            var v6 int32 = v1
            _ = v6
        }
        for fi7 := 0; fi7 < 1; fi7++ {
            if (!false) {
                fmt.Println(v4, v1)
                v3 = v3
            } else {
                var v8 int16 = 1
                _ = v8
            }
            v4 = v4
            _ = fi7
        }
        v3 = 0.0
    } else {
        var v9 int32 = v1
        v10 := []int{128, 127, -248, 32767}
        v1 += v9
        _ = v10
        _ = v9
    }
    v11 := []int{-538, -185, 127, 128}
    if len(v11) > 0 {
        v11[0] = 32767
    }
    v12 := GenS1{F0: 0, F1: 1, F2: 0}
    v11 = append(v11, 128)
    if len(v11) > 0 {
        v11[0] = -210
    }
    _ = len(v11)
    v5 -= (0.1 * (v5 - 0.0))
    if len(v11) > 0 {
        v11[0] = -385
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
