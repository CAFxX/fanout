package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 int64; F2 uint16 }
type GenS3 struct { F0 bool }

func main() {
    var v1 uint64 = 1
    v2 := make(map[int]int)
    v3 := make(map[int]string)
    var v4 uint8 = 255
    var v5 uint8 = v4
    if true {
        v4 -= (v4 & v5)
        if false {
            var v6 float64 = (1.5 - 1.5)
            var v7 float32 = 1e10
            v6 -= -2.25
            _ = v6
            _ = v7
        } else {
            fmt.Println(v1)
        }
    }
    v8 := v2[0]
    v9 := GenS1{F0: 0}
    v10 := v2[(v8 - (^(v8)))]
    var v11 float32 = 0.1
    var v12 int32 = -1
    v4 = v4
    if (255 > 1) {
        v12 = v12
        fmt.Println(v4, v1, v1)
        _ = len(v2)
    }
    if (836 >= 255) {
        var v13 uint8 = v4
        if false {
            for fi14 := 0; fi14 < 7; fi14++ {
                var v15 float64 = 0.1
                _ = fi14
                _ = v15
            }
        } else {
            fmt.Println(v10, v11)
            v16 := []int{int(v1), (v10 - (v10 & v10)), 128, 9223372036854775807}
            _ = v16
        }
        v17 := GenS2{F0: 314, F1: 9223372036854775807, F2: 55}
        _ = v13
        _ = v17
    } else {
        v18 := GenS3{F0: false}
        var v19 int = v8
        v12 = v12
        _ = v18
        _ = v19
    }
    var v20 int64 = -867
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
