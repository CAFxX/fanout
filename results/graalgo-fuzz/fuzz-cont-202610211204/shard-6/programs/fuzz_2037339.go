package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }

func main() {
    var v1 bool = false
    var v2 int = 602
    var v3 float32 = (3.14159 / (0.1 / float32((1.5 - ((0.0 - float64(((0.0 * 0.0) + ((float32(0.0) - 1.5) + (1.5 / 1e10))))) - ((3.14159 - 0.0) * 1.5))))))
    var v4 uint = 42
    v5 := []int{(v2 | (v2 / ((v2) | 1))), (-(v2))}
    var v6 bool = false
    for fi7 := 0; fi7 < 9; fi7++ {
        _ = len(v5)
        var v8 int32 = -201
        _ = fi7
        _ = v8
    }
    if v1 {
        var v9 bool = ("" >= ("" + ""))
        var v10 int32 = 1
        v10 = (v10 >> ((v4) & 63))
        _ = v10
        _ = v9
    } else {
        var v11 int32 = -1
        for fi12 := 0; fi12 < 20; fi12++ {
            if true {
                fmt.Println(fi12)
                var v13 uint64 = 463
                _ = v13
            } else {
                var v14 float64 = (-2.25 + 1.5)
                v2 = (fi12 >> ((v4) & 63))
                if false {
                    if len(v5) > 0 {
                        v5[0] = -1
                    }
                    v3 += 3.14159
                }
                _ = v14
            }
            var v15 float32 = v3
            _ = len(v5)
            _ = fi12
            _ = v15
        }
        _ = v11
    }
    if (!v1) {
        var v16 int32 = 0
        v1 = v1
        _ = v16
    }
    fmt.Println(v3)
    var v17 uint8 = 255
    v18 := GenS1{F0: 42, F1: v17}
    var v19 int32 = 0
    v17 -= (v17 * (v17 | v17))
    var v20 uint16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
