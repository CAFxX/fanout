package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
func genFn0() uint8 {
    var v7 string = "foo bar"
    v7 = v7
    var v8 float64 = 1e10
    var v9 uint = 4294967295
    v8 = 1e10
    _ = v7
    _ = v8
    _ = v9
    return uint8(v9)
}

func main() {
    var v1 string = "hello"
    v2 := make(map[uint32]string)
    v3 := make(map[int]int)
    var v4 uint32 = 4294967295
    var v5 float64 = float64((float32(1e10) / 1.5))
    var v6 float32 = (((1.5 + (0.0 - 0.1)) + -2.25) * 3.14159)
    v10 := v2[v4]
    if (-961 > 42) {
        if (v6 <= 1.5) {
            v11 := GenS1{F0: 1, F1: 2}
            if false {
                v6 -= (v6 - ((-2.25 * v6) + v6))
                var v12 uint64 = 18446744073709551615
                if true {
                    v13 := GenS1{F0: v12, F1: 128}
                    v14 := GenS1{F0: (v12 / ((uint64(v4)) | 1)), F1: 235}
                    _ = v13
                    _ = v14
                }
                _ = v12
            } else {
                v15 := []int{128, 4294967295, 42, 93}
                _ = v15
            }
            v16 := GenS1{F0: 4294967295, F1: 42}
            _ = v11
            _ = v16
        } else {
            v17 := []int{32767, 4294967295, 107, 1}
            var v18 int16 = -539
            _ = v17
            _ = v18
        }
        v19 := GenS1{F0: 1, F1: 0}
        v6 = v6
        _ = v19
    }
    fmt.Println(v5, v10, v1)
    var v20 int8 = 100
    var v21 int16 = 1
    fmt.Println(v10, v1)
    _ = len(v3)
    _ = len(v2)
    v22 := []int{}
    v20 += v20
    var v23 int32 = -1
    switch 1 {
    case 717:
        var v24 bool = (951 == v4)
        _ = v24
    }
    switch v1 {
    case "a":
        v20 -= int8(v4)
    }
    var v25 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
