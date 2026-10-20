package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 string }

func main() {
    v1 := make(map[string]bool)
    v2 := []int{}
    var v3 string = "foo bar"
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v4 float64 = 0.1
    if false {
        fmt.Println(v4, v4)
    } else {
        switch 32767 {
        case 128:
            v3 = "a"
        default:
            var v5 bool = false
            _ = v5
        }
        if (!true) {
            v6 := GenS1{F0: 42, F1: 42, F2: "hello"}
            fmt.Println(v3, v3)
            _ = v6
        }
    }
    v4 = (v4 + float64(3.14159))
    var v7 uint16 = 255
    v3 = v3
    var v8 int8 = -128
    var v9 uint8 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
