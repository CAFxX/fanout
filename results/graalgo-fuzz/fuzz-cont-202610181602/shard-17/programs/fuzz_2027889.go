package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 uint32 }

func main() {
    var v1 uint8 = 255
    v2 := []int{1, 281, -2147483648, 42}
    v3 := []string{"hello", "a", "foo bar"}
    var v4 int64 = -16
    var v5 float32 = ((3.14159 / 1.0) + 1.5)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v5 = (-2.25 / 1e10)
    fmt.Println(v4)
    var v6 float64 = (-2.25 + 0.0)
    var v7 uint32 = 42
    if false {
        fmt.Println(v6, v5)
        v5 = (((3.14159 - v5) - (v5 / v5)) / float32((((v6 * float64((1e10 + v5))) * v6) + 1.5)))
    }
    switch "\n\t" {
    case "\n\t":
        var v8 uint8 = 128
        _ = v8
    default:
        v9 := GenS1{F0: (v4 << ((961) & 63)), F1: 4294967295, F2: 0}
        _ = v9
    }
    var v10 uint16 = 376
    var v11 int16 = 1000
    if false {
        fmt.Println(v4, v7, v11)
        var v12 string = ("a" + "")
        var v13 uint64 = 18446744073709551615
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: (v4 % ((-1) | 1)), F1: 42, F2: 1}
    v1 = (v1 / ((v1) | 1))
    var v15 bool = (!(!false))
    v5 += v5
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
