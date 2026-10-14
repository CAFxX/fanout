package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = 0
    var v3 int64 = 9223372036854775807
    v4 := make(map[int]int)
    var v5 int = -2147483648
    v6 := []string{}
    var v7 int64 = (-(v2))
    _ = len(v1)
    var v8 float32 = 3.14159
    v9 := GenS1{F0: true, F1: v5}
    _ = len(v4)
    var v10 uint32 = 2147483648
    fmt.Println(v8, v10, v10)
    for fi11 := 0; fi11 < 5; fi11++ {
        v12 := GenS1{F0: (-2.25 < 1.5), F1: v5}
        v3 += v7
        if false {
            var v13 uint = 1
            _ = v13
        } else {
            v5 = v5
        }
        _ = fi11
        _ = v12
    }
    fmt.Println(v8)
    if len(v6) > 0 {
        v6[0] = "a"
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
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
