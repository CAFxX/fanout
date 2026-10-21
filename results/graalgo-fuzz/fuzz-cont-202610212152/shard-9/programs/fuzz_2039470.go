package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int }

func main() {
    v1 := []int{-1, 452, 2147483647, -2147483648}
    var v2 int8 = 25
    v3 := make(map[int]string)
    var v4 int32 = 0
    var v5 uint32 = 0
    var v6 int16 = 1000
    var v7 int32 = v4
    fmt.Println(v6)
    fmt.Println(v2)
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 int64 = 42
        for fi10 := 0; fi10 < 5; fi10++ {
            v2 -= v2
            var v11 uint16 = 255
            _ = fi10
            _ = v11
        }
        for fi12 := 0; fi12 < 13; fi12++ {
            v13 := GenS1{F0: (v6 * (v6 % ((v6) | 1))), F1: (v6 + (v6 >> ((65535) & 63))), F2: (fi12 + fi12)}
            var v14 float32 = 0.0
            v15 := GenS1{F0: int16(v2), F1: (v6 * (v6 & (v6 & v6))), F2: fi8}
            v5 = 0
            v16 := v3[fi8]
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
            _ = v16
        }
        _ = fi8
        _ = v9
    }
    var v17 float64 = float64(-2.25)
    var v18 float32 = float32((-2.25 + v17))
    var v19 int16 = 698
    v20 := GenS1{F0: v19, F1: (v6 / ((-791) | 1)), F2: 531}
    var v21 bool = (true || ((true && true) && false))
    var v22 int64 = 9223372036854775807
    v22 -= v22
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
