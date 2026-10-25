package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint32 }
func genFn0(p0 uint64, p1 int8, p2 uint) uint64 {
    fmt.Println(p1)
    v6 := GenS1{F0: false, F1: -16, F2: 42}
    var v7 float32 = (-2.25 / 1.0)
    v7 += (v7 + 0.0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return p0
}
func genFn1(p0 string) bool {
    var v8 int32 = -1
    v8 = (^(v8))
    var v9 int8 = 112
    var v10 int16 = -32768
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return (!true)
}

func main() {
    var v1 uint64 = 531
    var v2 int64 = -211
    var v3 int = 128
    var v4 float64 = float64(3.14159)
    v5 := make(map[uint32]string)
    v11 := GenS1{F0: true, F1: 32767, F2: 260}
    v4 = v4
    var v12 bool = true
    v13 := []int{(v3 % ((v3) | 1)), -400}
    v3 += v3
    if v12 {
        for fi14 := 0; fi14 < 19; fi14++ {
            v12 = (4294967295 == 4294967295)
            for fi15 := 0; fi15 < 5; fi15++ {
                var v16 int16 = 42
                _ = fi15
                _ = v16
            }
            _ = fi14
        }
        v17 := GenS1{F0: (1 > (v2 / (((^(v2))) | 1))), F1: 656, F2: 42}
        _ = v17
    }
    delete(v5, 0)
    var v18 uint32 = 4294967295
    v19 := GenS1{F0: ((v2 / ((v2) | 1)) > 9223372036854775807), F1: 1, F2: (^(v18))}
    for fi20 := 0; fi20 < 16; fi20++ {
        if len(v13) > 0 {
            v13[0] = v3
        }
        _ = fi20
    }
    var v21 uint32 = (v18 % ((2147483648) | 1))
    delete(v5, 42)
    v22 := GenS1{F0: false, F1: 595, F2: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
