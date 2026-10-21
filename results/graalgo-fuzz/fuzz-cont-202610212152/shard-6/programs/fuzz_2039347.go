package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }

func main() {
    var v1 float32 = 1e10
    var v2 bool = ((0 >= 0) && (18446744073709551615 < 255))
    var v3 int = 274
    v4 := make(map[string]string)
    if (!(("\n\t" > "hello") && (1 != 1))) {
        fmt.Println(v2, v1)
    } else {
        if (255 >= 878) {
            v5 := GenS1{F0: -64, F1: (v3 + v3)}
            v3 += (v3 / (((v3 / ((-9223372036854775808) | 1))) | 1))
            fmt.Println(v3, v2)
            _ = v5
        }
        var v6 bool = ((-2.25 + -2.25) > 0.1)
        _ = v6
    }
    v2 = (("hello" == "foo bar") || false)
    var v7 int64 = 426
    fmt.Println(v2, v1, v2)
    v1 = (v1 * 1.5)
    var v8 float64 = -2.25
    v1 = 1.5
    for fi9 := 0; fi9 < 11; fi9++ {
        v10 := []int{(fi9 + fi9)}
        _ = fi9
        _ = v10
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
