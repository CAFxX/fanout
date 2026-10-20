package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
func genFn0() string {
    v5 := []int{-512, -9223372036854775808}
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := GenS1{F0: 1, F1: 255}
        fmt.Println(fi6)
        _ = fi6
        _ = v7
    }
    _ = v5
    return ("a" + "hello")
}

func main() {
    var v1 float32 = ((3.14159 + 1.5) / 1.0)
    var v2 uint8 = 127
    var v3 float64 = 0.0
    var v4 bool = true
    v8 := []int{128, 2147483647, -650}
    v1 = v1
    v4 = (((("hello" + "foo bar") + "foo bar") + "\n\t") < "a")
    fmt.Println(v4)
    v1 += v1
    var v9 string = "hello"
    fmt.Println(v9, v1)
    v9 = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
