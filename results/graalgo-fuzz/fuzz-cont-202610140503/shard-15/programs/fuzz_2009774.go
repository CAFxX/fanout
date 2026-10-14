package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint64 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 uint16 = 42
    v2 := []string{("hello" + "a"), ("hello" + "a"), (("\n\t" + "hello") + "a")}
    var v3 int64 = 42
    var v4 float32 = (0.0 + -2.25)
    v5 := make(map[string]int)
    v6 := GenS1{F0: v3, F1: 677, F2: 94}
    v4 = v4
    v7 := GenS2{F0: 255}
    v5[("hello" + "a")] = 901
    v3 = v3
    if true {
        var v8 float64 = (-2.25 - 0.0)
        v4 = float32((v8 / float64(0.1)))
        if false {
            var v9 int64 = (v3 * 995)
            var v10 int = 2147483647
            _ = v10
            _ = v9
        }
        _ = v8
    }
    var v11 uint32 = 65535
    switch "" {
    case "hello":
        v12 := v5[("\n\t" + "foo bar")]
        _ = v12
    case "a":
        var v13 int = 1
        _ = v13
    case "foo bar":
        fmt.Println(v11, v4, v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
