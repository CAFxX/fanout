package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    v1 := []string{}
    v2 := make(map[string]bool)
    var v3 int64 = 1
    v4 := GenS1{F0: "foo bar"}
    fmt.Println(v3)
    v3 -= -172
    _ = len(v1)
    for fi5 := 0; fi5 < 15; fi5++ {
        _ = len(v2)
        var v6 float64 = 0.1
        _ = fi5
        _ = v6
    }
    v3 -= (v3 ^ (v3 / ((-1) | 1)))
    v3 += -1
    var v7 int16 = 344
    var v8 uint8 = 42
    if false {
        fmt.Println(v3, v7)
        var v9 int32 = 1
        _ = v9
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
