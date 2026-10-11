package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 string; F1 int16; F2 string }

func main() {
    var v1 uint8 = 255
    var v2 string = "a"
    var v3 bool = (!(false && (-1 <= -1)))
    var v4 string = "foo bar"
    v3 = v3
    v5 := GenS1{F0: 643, F1: v1}
    v4 = v4
    fmt.Println(v3)
    v6 := []int{388, 9223372036854775807}
    v1 = 42
    v6 = append(v6, -1)
    v7 := GenS2{F0: ((((v2 + ("" + v2)) + "foo bar") + "hello") + v4), F1: 1000, F2: v4}
    if len(v6) > 0 {
        v6[0] = 278
    }
    v8 := GenS2{F0: v2, F1: 42, F2: ""}
    fmt.Println(v3, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
