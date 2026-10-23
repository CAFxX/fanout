package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint64 }
type GenS2 struct { F0 uint16; F1 int8 }
func genFn0(p0 uint, p1 int8) string {
    var v4 int8 = 112
    v5 := GenS1{F0: 1000, F1: (!false), F2: 1}
    var v6 float32 = 1.5
    for fi7 := 0; fi7 < 7; fi7++ {
        fmt.Println(p0)
        p1 = v4
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return "hello"
}

func main() {
    v1 := []int{-916, 128}
    var v2 uint64 = 1
    var v3 uint16 = 1
    var v8 uint16 = v3
    fmt.Println(v2)
    var v9 string = ("hello" + ("foo bar" + ""))
    var v10 uint64 = 0
    v1 = append(v1, 800)
    v11 := GenS2{F0: (^(v8)), F1: 0}
    fmt.Println(v3, v10, v3)
    var v12 bool = true
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
