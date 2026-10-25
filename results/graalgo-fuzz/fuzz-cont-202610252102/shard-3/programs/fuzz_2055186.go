package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 uint64 = 219
    v2 := []string{"foo bar", "foo bar"}
    v3 := make(map[uint32]bool)
    var v4 int64 = -9223372036854775808
    if true {
        fmt.Println(v1)
    }
    v1 = (v1 * v1)
    v1 += (v1 + (v1 - v1))
    v5 := v3[2147483647]
    v6 := GenS1{F0: -507}
    for fi7 := 0; fi7 < 13; fi7++ {
        _ = len(v2)
        _ = fi7
    }
    fmt.Println(v1, v5, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
