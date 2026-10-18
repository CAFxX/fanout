package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int64 }
type GenS2 struct { F0 uint8; F1 int; F2 string }

func main() {
    v1 := make(map[int]bool)
    var v2 float64 = ((0.0 - 0.1) * 0.1)
    var v3 uint = 825
    v4 := v1[-9223372036854775808]
    var v5 uint = 255
    var v6 uint16 = 32767
    v3 = v5
    var v7 int32 = 0
    v1[4294967295] = (!v4)
    var v8 float64 = v2
    v9 := GenS1{F0: (v7 + 2147483647), F1: "foo bar", F2: -9223372036854775808}
    v2 -= ((float64(3.14159) / (((((-2.25 / float64(1.5)) * v2) - v8) * (float64(3.14159) / v2)) * (float64(1e10) / v8))) / 3.14159)
    v5 -= v5
    if (!v4) {
        var v10 string = "hello"
        _ = v10
    } else {
        v11 := GenS2{F0: 127, F1: -913, F2: ("a" + "a")}
        _ = v11
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
