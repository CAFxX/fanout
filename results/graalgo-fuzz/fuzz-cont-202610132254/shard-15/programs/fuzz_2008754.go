package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 bool }

func main() {
    var v1 string = "a"
    var v2 int64 = 0
    var v3 float64 = (float64(float32(1e10)) * float64((1.5 * (0.1 + (((0.1 / 1.5) + (-2.25 / 1.0)) * 1.5)))))
    var v4 bool = false
    v1 = (v1 + v1)
    v5 := []int{-223, 1}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if v4 {
        var v6 float32 = 1e10
        var v7 uint32 = 2
        fmt.Println(v2, v6, v3)
        _ = v6
        _ = v7
    }
    v2 -= (-(v2))
    var v8 int32 = 2147483647
    v8 = v8
    fmt.Println(v8)
    var v9 int16 = 241
    var v10 int64 = -617
    v5 = append(v5, -9223372036854775808)
    v11 := GenS1{F0: 0, F1: -13, F2: v4}
    fmt.Println(v3, v8)
    switch (("foo bar" + "\n\t") + "foo bar") {
    case "":
        v10 += 0
    default:
        if v4 {
            var v12 uint32 = 0
            var v13 float64 = 0.1
            var v14 string = "\n\t"
            _ = v12
            _ = v13
            _ = v14
        } else {
            var v15 uint = 2
            var v16 int32 = 2147483647
            v1 = "\n\t"
            _ = v15
            _ = v16
        }
    }
    switch (v1 + v1) {
    case "hello":
        if len(v5) > 0 {
            v5[0] = 32768
        }
    }
    v9 -= 1000
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
