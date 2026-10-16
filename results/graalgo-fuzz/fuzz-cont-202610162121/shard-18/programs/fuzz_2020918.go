package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 int64 }
type GenS2 struct { F0 int; F1 int64; F2 uint64 }

func main() {
    var v1 int64 = 0
    var v2 float64 = 0.0
    var v3 int32 = -57
    var v4 string = "foo bar"
    fmt.Println(v1)
    var v5 int = 9223372036854775807
    var v6 string = ("foo bar" + ("hello" + ("foo bar" + "")))
    if (-9223372036854775808 >= 42) {
        fmt.Println(v3, v4)
    }
    var v7 int64 = (v1 >> ((1) & 63))
    fmt.Println(v7)
    v8 := []int{(v5 >> ((2) & 63)), (v5 ^ v5), int(v7), -139}
    if len(v8) > 0 {
        v8[0] = -2147483648
    }
    v9 := GenS1{F0: v7, F1: "hello", F2: (v7 & (^(v7)))}
    var v10 float32 = 0.1
    v11 := GenS2{F0: v5, F1: -1, F2: 248}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
