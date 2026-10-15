package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 uint16; F1 int8 }

func main() {
    var v1 uint64 = 173
    var v2 bool = ((853 >= 937) && false)
    var v3 uint8 = 77
    v2 = (("a" == "foo bar") || (-684 <= 689))
    v4 := GenS1{F0: (((("hello" + "a") + ("a" + ("" + "hello"))) + ("a" + "a")) + ("foo bar" + "\n\t")), F1: int16(v1)}
    fmt.Println(v3)
    var v5 float32 = 1e10
    var v6 uint = 255
    fmt.Println(v6)
    v7 := GenS1{F0: "\n\t", F1: 0}
    v8 := []int{127, -905, -2147483648, -2147483648}
    v9 := GenS2{F0: 1, F1: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
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
