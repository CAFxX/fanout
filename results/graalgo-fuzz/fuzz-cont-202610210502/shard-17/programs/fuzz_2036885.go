package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 bool; F1 int; F2 int32 }
func genFn0(p0 uint32) int32 {
    v6 := []int{727, 369, -868}
    if len(v6) > 0 {
        v6[0] = -2147483648
    }
    _ = p0
    _ = v6
    return 389
}
func genFn1() int16 {
    v7 := []int{42, -9223372036854775808, -266, 0}
    if false {
        v7 = append(v7, 127)
    }
    var v8 int = 2147483647
    _ = v7
    _ = v8
    return 32767
}

func main() {
    var v1 uint64 = 0
    var v2 float32 = 0.0
    v3 := make(map[uint32]uint32)
    var v4 string = ("hello" + "a")
    v5 := make(map[string]int)
    var v9 int64 = 315
    v10 := GenS1{F0: 127, F1: 2}
    var v11 int64 = v9
    v1 -= (^(v1))
    if true {
        var v12 int8 = 42
        delete(v5, "\n\t")
        v13 := GenS1{F0: v12, F1: 255}
        _ = v12
        _ = v13
    } else {
        for fi14 := 0; fi14 < 17; fi14++ {
            for fi15 := 0; fi15 < 9; fi15++ {
                if true {
                    var v16 bool = true
                    _ = len(v5)
                    var v17 float64 = 0.0
                    v2 -= float32((((1.5 / float64(v2)) + v17) + v17))
                    var v18 uint32 = 611
                    _ = v16
                    _ = v17
                    _ = v18
                }
                var v19 uint8 = 127
                v20 := v3[2147483647]
                _ = fi15
                _ = v19
                _ = v20
            }
            var v21 bool = false
            _ = fi14
            _ = v21
        }
        if false {
            if ((0 < genFn0(0)) && false) {
                var v22 int64 = v11
                v23 := GenS1{F0: 127, F1: 76}
                var v24 uint32 = 597
                _ = v22
                _ = v23
                _ = v24
            } else {
                fmt.Println(v9)
                _ = len(v5)
            }
            v25 := GenS2{F0: (!true), F1: 127, F2: -888}
            _ = v25
        } else {
            v1 -= v1
        }
    }
    v11 += (v9 % ((v11) | 1))
    fmt.Println(v1, v11, v9)
    fmt.Println(v4)
    v4 = v4
    var v26 int = 42
    v27 := GenS2{F0: true, F1: (v26 | (v26 & (^(v26)))), F2: -1}
    var v28 int = -221
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
