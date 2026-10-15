package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 int16 }
func genFn0() int {
    switch 4294967295 {
    case 0:
        var v6 int8 = 127
        _ = v6
    }
    var v7 int32 = 0
    v7 = v7
    var v8 int8 = 127
    _ = v7
    _ = v8
    return -9223372036854775808
}

func main() {
    var v1 int = -930
    v2 := make(map[int]uint32)
    var v3 float32 = (1.5 * 3.14159)
    var v4 int64 = 0
    var v5 float64 = float64(-2.25)
    v9 := v2[v1]
    v10 := []int{-2147483648, (v1 / ((v1) | 1)), (v1 ^ 309)}
    delete(v2, v1)
    v3 = 1e10
    var v11 string = "foo bar"
    var v12 uint64 = 4294967295
    switch (v1 >> ((0) & 63)) {
    case -550:
        v5 = (((1.5 * v5) - float64(v3)) + v5)
    }
    _ = len(v10)
    var v13 int16 = 188
    var v14 float32 = -2.25
    var v15 int8 = 100
    switch (v11 + v11) {
    case "hello":
        var v16 uint64 = v12
        _ = v16
    case "\n\t":
        fmt.Println(v3)
    default:
        for fi17 := 0; fi17 < 14; fi17++ {
            var v18 float64 = (0.1 / ((-2.25 * v5) + v5))
            v19 := GenS1{F0: v4}
            _ = fi17
            _ = v18
            _ = v19
        }
    }
    var v20 int64 = int64(v1)
    v21 := GenS1{F0: 9223372036854775807}
    var v22 uint32 = v9
    v23 := GenS2{F0: uint32(v12), F1: (v13 * 0)}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
