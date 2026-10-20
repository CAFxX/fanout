package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 bool }

func main() {
    var v1 uint = 2
    v2 := make(map[uint32]int)
    var v3 int8 = -80
    v4 := make(map[int]bool)
    var v5 float64 = -2.25
    var v6 bool = (0 != 42)
    for fi7 := 0; fi7 < 6; fi7++ {
        v8 := GenS1{F0: -1, F1: v1, F2: (((32767 < 0) && v6) && (42 == 255))}
        var v9 string = "\n\t"
        for fi10 := 0; fi10 < 8; fi10++ {
            if (false || true) {
                if true {
                    var v11 int8 = v3
                    var v12 float32 = ((1e10 - -2.25) * 0.1)
                    v13 := GenS1{F0: 1, F1: v1, F2: false}
                    _ = v11
                    _ = v12
                    _ = v13
                }
            }
            v14 := GenS1{F0: -89, F1: v1, F2: v6}
            v3 += (v3 ^ (^(v3)))
            _ = fi10
            _ = v14
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    v15 := []int{42}
    if len(v15) > 0 {
        v15[0] = 0
    }
    v5 += ((v5 * v5) + (float64((((0.1 + (0.0 * 0.1)) / 1e10) / ((0.0 * 1.5) * (0.1 + (0.0 + float32((v5 + v5))))))) / (float64((((-2.25 - float32(v5)) / float32(((v5 / ((1e10 * -2.25) + v5)) + v5))) + (3.14159 / 0.1))) + (0.1 - 3.14159))))
    v16 := v4[127]
    v16 = v16
    var v17 uint64 = 189
    var v18 float32 = ((1e10 * 1e10) - float32(float64(((0.1 - 0.1) - 1e10))))
    v1 += v1
    if (42 >= -1) {
        var v19 string = "\n\t"
        if v6 {
            v5 = v5
        } else {
            v5 = v5
        }
        v18 += v18
        _ = v19
    } else {
        _ = len(v15)
        v20 := v4[2147483647]
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
