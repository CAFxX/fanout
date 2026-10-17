package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[string]string)
    v3 := make(map[uint32]int)
    v4 := []int{42, 4294967295, -930}
    switch "\n\t" {
    case "":
        v2["foo bar"] = "hello"
    case "a":
        var v5 uint32 = 154
        _ = v5
    }
    v1 = (((v1 / 1.0) / v1) - v1)
    var v6 uint32 = 109
    var v7 int = 356
    v1 -= v1
    v7 = v7
    var v8 uint = 42
    for fi9 := 0; fi9 < 11; fi9++ {
        v10 := GenS1{F0: 2}
        var v11 string = "foo bar"
        fmt.Println(v11)
        _ = fi9
        _ = v10
        _ = v11
    }
    v12 := GenS1{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
