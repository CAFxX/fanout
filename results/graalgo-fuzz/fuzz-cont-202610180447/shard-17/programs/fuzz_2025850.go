package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 int }

func main() {
    v1 := []string{"a", "", "hello"}
    var v2 int32 = 2147483647
    v3 := []string{"a"}
    fmt.Println(v2)
    fmt.Println(v2)
    v2 -= (v2 + -337)
    v2 = 2147483647
    v4 := GenS1{F0: 42, F1: "a", F2: 4294967295}
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    v2 -= 42
    var v5 int32 = (v2 | (v2 % ((v2) | 1)))
    for fi6 := 0; fi6 < 4; fi6++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v7 int = (fi6 / ((fi6) | 1))
        _ = fi6
        _ = v7
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
