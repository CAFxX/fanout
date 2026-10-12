package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 int32 }

func main() {
    var v1 int16 = 32767
    var v2 string = ("\n\t" + "hello")
    v3 := make(map[int]string)
    v4 := make(map[string]string)
    var v5 uint64 = 191
    v2 = v2
    for fi6 := 0; fi6 < 5; fi6++ {
        v1 += v1
        v5 -= (v5 + (v5 << ((403) & 63)))
        _ = fi6
    }
    var v7 string = v2
    v8 := []int{-9223372036854775808, 2147483647, 9223372036854775807, -268}
    var v9 float64 = 0.1
    v9 = (v9 - (((v9 - 0.0) + (v9 / 1e10)) - ((float64(((0.0 * (((1.5 / -2.25) - (-2.25 / 1.5)) * (float32(v9) / float32(1e10)))) + float32(v9))) - v9) * -2.25)))
    if (2147483647 != 0) {
        fmt.Println(v1, v7, v1)
    }
    var v10 uint = 69
    v9 = (1.5 / (0.0 / float64(float32(v9))))
    if false {
        var v11 uint16 = 344
        for fi12 := 0; fi12 < 16; fi12++ {
            fmt.Println(v9, v1)
            _ = fi12
        }
        if true {
            v8 = append(v8, 0)
            if len(v8) > 0 {
                v8[0] = 489
            }
            v2 = v7
        }
        _ = v11
    } else {
        v7 = (("a" + v2) + (v2 + "hello"))
        for fi13 := 0; fi13 < 18; fi13++ {
            v14 := GenS1{F0: v5, F1: true, F2: 42}
            if ((1 != v1) && false) {
                var v15 float64 = 1.5
                _ = v15
            }
            var v16 string = v2
            _ = fi13
            _ = v14
            _ = v16
        }
    }
    v5 -= (v5 << ((255) & 63))
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    var v17 int32 = -79
    v17 = (^(v17))
    var v18 int = 8
    var v19 string = (("hello" + ("" + (v2 + v7))) + v7)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
