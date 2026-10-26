package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }

func main() {
    var v1 int64 = -630
    var v2 int32 = 1
    var v3 uint = 0
    var v4 int16 = 0
    v5 := make(map[uint32]int)
    v6 := v5[42]
    v5[2] = 2147483647
    delete(v5, 1)
    var v7 string = (("hello" + "foo bar") + ("foo bar" + "hello"))
    var v8 int8 = 100
    var v9 float32 = -2.25
    v8 = 127
    v10 := GenS1{F0: 0, F1: v2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
