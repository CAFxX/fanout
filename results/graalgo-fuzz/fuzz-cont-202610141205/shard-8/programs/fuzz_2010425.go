package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 int64 }

func main() {
    var v1 uint64 = 611
    v2 := make(map[string]string)
    v3 := make(map[string]int)
    v1 = (v1 & v1)
    switch 0 {
    case 127:
        var v4 string = ("foo bar" + "")
        _ = v4
    }
    var v5 float64 = -2.25
    fmt.Println(v1, v5)
    v6 := GenS1{F0: 1, F1: 255, F2: 42}
    v7 := v3[""]
    var v8 bool = true
    var v9 int16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
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
