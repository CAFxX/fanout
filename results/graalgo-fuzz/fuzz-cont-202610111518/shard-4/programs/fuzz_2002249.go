package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 string }
type GenS2 struct { F0 int; F1 string; F2 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint }

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]string)
    var v4 uint64 = 0
    var v5 int16 = -1
    fmt.Println(v5)
    v5 += -1
    v6 := []int{-2147483648, -1, 127, -2147483648}
    switch "" {
    case "a":
        if (-1 != 54) {
            v4 += 0
            v7 := GenS1{F0: 2, F1: 2147483648, F2: "\n\t"}
            v8 := GenS2{F0: 1000, F1: "hello", F2: -827}
            _ = v7
            _ = v8
        }
    }
    var v9 int = 1
    v10 := GenS1{F0: 42, F1: 2147483647, F2: "\n\t"}
    v11 := GenS2{F0: v9, F1: (("\n\t" + "foo bar") + "\n\t"), F2: -9223372036854775808}
    _ = len(v3)
    fmt.Println(v4, v4, v5)
    switch "hello" {
    case "hello":
        if len(v6) > 0 {
            v6[0] = 4294967295
        }
    case "a":
        if (true && true) {
            fmt.Println(v1, v5)
            if true {
                var v12 float64 = 1e10
                var v13 uint64 = (^(v4))
                v12 += 0.1
                _ = v12
                _ = v13
            }
        } else {
            fmt.Println(v5, v1, v4)
        }
    }
    var v14 float32 = (1.5 * float32(3.14159))
    var v15 int32 = 42
    v16 := GenS3{F0: "hello"}
    var v17 float32 = ((0.0 - 0.1) + 3.14159)
    if (9223372036854775807 != (v1 % ((870) | 1))) {
        fmt.Println(v1)
    } else {
        v18 := GenS4{F0: 748}
        v17 = 0.1
        fmt.Println(v1, v15, v4)
        _ = v18
    }
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
