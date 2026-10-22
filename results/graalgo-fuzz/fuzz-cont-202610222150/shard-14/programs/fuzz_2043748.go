package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 int8) int64 {
    p0 = p0
    if ((!((("a" + "hello") < "") && (255 < 42))) && true) {
        var v7 float32 = 0.0
        _ = v7
    }
    p0 -= 42
    v8 := []int{32767, 32768, -9223372036854775808, -181}
    fmt.Println(p0)
    _ = p0
    _ = v8
    return 121
}

func main() {
    var v1 uint = 2
    v2 := []int{-2147483648, 32767, -309}
    v3 := make(map[int]int)
    var v4 int64 = -9223372036854775808
    v5 := make(map[uint32]uint32)
    var v6 uint16 = 32767
    fmt.Println(v6, v1)
    v1 += v1
    if (-135 >= -32768) {
        var v9 float32 = (3.14159 - (1.5 - -2.25))
        _ = v9
    } else {
        v1 = (v1 | uint(v6))
        if true {
            var v10 uint16 = v6
            v10 = (v10 | (^(v6)))
            _ = v10
        } else {
            var v11 int = -861
            var v12 string = "a"
            _ = v11
            _ = v12
        }
    }
    var v13 uint64 = 2
    var v14 int = -1
    if (499 >= (v4 / ((v4) | 1))) {
        v14 += 818
    } else {
        v15 := GenS1{F0: -2147483648, F1: 100}
        v16 := GenS2{F0: 0}
        _ = v15
        _ = v16
    }
    if ((v13 ^ v13) >= 1) {
        fmt.Println(v1, v6, v1)
        v17 := GenS3{F0: (^(v13))}
        _ = v17
    }
    v13 = uint64(v6)
    v18 := GenS1{F0: 279, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
