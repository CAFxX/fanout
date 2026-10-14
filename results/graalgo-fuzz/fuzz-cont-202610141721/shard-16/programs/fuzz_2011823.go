package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 uint16 = 109
    var v2 string = "foo bar"
    v3 := make(map[int]int)
    var v4 float64 = 0.1
    v5 := make(map[string]string)
    v6 := []string{}
    var v7 uint = 357
    fmt.Println(v7)
    if len(v6) > 0 {
        v6[0] = (v2 + (v2 + "foo bar"))
    }
    v8 := GenS1{F0: -1}
    switch "a" {
    case "":
        fmt.Println(v4)
    case "hello":
        var v9 uint = (v7 - uint(v1))
        _ = v9
    case "a":
        v1 += v1
    }
    v10 := GenS1{F0: 128}
    fmt.Println(v4, v2)
    v1 += v1
    v11 := GenS1{F0: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
