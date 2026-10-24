package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint64 }
type GenS2 struct { F0 int32; F1 bool }
func genFn0(p0 uint16, p1 uint8, p2 int64) bool {
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 int = 128
        _ = fi6
        _ = v7
    }
    var v8 float64 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return false
}

func main() {
    var v1 string = "a"
    var v2 int64 = 1
    var v3 uint64 = 4294967295
    var v4 int8 = 127
    var v5 int64 = v2
    var v9 float32 = 0.1
    var v10 uint8 = 188
    v11 := []int{}
    if (v9 == (((float32(1.5) / v9) + float32((1e10 + (0.0 + (3.14159 - (0.1 + 1e10)))))) * v9)) {
        fmt.Println(v9, v4, v9)
        for fi12 := 0; fi12 < 17; fi12++ {
            if (583 < 2147483648) {
                if true {
                    v13 := GenS1{F0: fi12, F1: (v1 + "hello"), F2: 4294967295}
                    var v14 uint8 = (v10 + v10)
                    var v15 string = v1
                    _ = v13
                    _ = v14
                    _ = v15
                } else {
                    v4 -= (v4 / ((100) | 1))
                    v4 = (v4 + -1)
                }
                if len(v11) > 0 {
                    v11[0] = (fi12 & (fi12 & fi12))
                }
                v4 += (^(v4))
            } else {
                fmt.Println(v3, v2, v10)
                if false {
                    var v16 uint = 913
                    var v17 float32 = v9
                    _ = v16
                    _ = v17
                } else {
                    if len(v11) > 0 {
                        v11[0] = (fi12 & (fi12 - (fi12 | -2147483648)))
                    }
                }
            }
            _ = fi12
        }
        if len(v11) > 0 {
            v11[0] = -94
        }
    }
    v1 = v1
    _ = len(v11)
    switch 32768 {
    case 32768:
        var v18 string = (v1 + v1)
        _ = v18
    case 127:
        v3 -= v3
    }
    var v19 int32 = -2147483648
    fmt.Println(v19)
    v20 := GenS1{F0: 32768, F1: (v1 + ((v1 + v1) + (v1 + "\n\t"))), F2: v3}
    v9 = (((v9 - float32((1e10 + (0.1 + 0.0)))) + v9) + 1.5)
    fmt.Println(v10)
    v21 := GenS2{F0: (v19 / ((-447) | 1)), F1: true}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
