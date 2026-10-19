package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    v1 := []string{"a", "hello", "hello", "foo bar"}
    v2 := make(map[string]bool)
    v3 := make(map[int]bool)
    var v4 float32 = 1e10
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v4 -= (v4 / v4)
    for fi5 := 0; fi5 < 9; fi5++ {
        v6 := GenS1{F0: 250}
        _ = fi5
        _ = v6
    }
    var v7 uint16 = 65535
    var v8 float64 = 3.14159
    v7 += 32767
    var v9 float32 = 1e10
    v8 = v8
    v7 = 255
    fmt.Println(v8)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
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
