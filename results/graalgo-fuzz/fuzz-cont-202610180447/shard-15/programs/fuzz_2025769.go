package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
func genFn0(p0 uint8, p1 bool) uint16 {
    p0 = (^(p0))
    v4 := GenS1{F0: 0, F1: -1}
    switch 547 {
    case 55:
        fmt.Println(p1, p1)
    }
    _ = p0
    _ = p1
    _ = v4
    return 42
}

func main() {
    var v1 float64 = 1.5
    v2 := []string{("hello" + "\n\t"), "hello", "\n\t"}
    var v3 uint8 = 0
    var v5 int16 = -221
    v5 -= (v5 >> ((2) & 63))
    var v6 string = "foo bar"
    var v7 uint64 = 0
    var v8 float64 = v1
    fmt.Println(v3, v1)
    fmt.Println(v5, v6)
    var v9 bool = (!(((!true) || (true && false)) || false))
    fmt.Println(v7, v9)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
