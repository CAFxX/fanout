package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int16; F1 int32 }

func main() {
    v1 := []int{4294967295}
    var v2 bool = false
    var v3 float32 = (0.1 - 0.1)
    var v4 float32 = v3
    v5 := []string{}
    v5 = append(v5, "a")
    var v6 float32 = (v4 - v3)
    var v7 int32 = 933
    if v2 {
        v6 = float32(3.14159)
    }
    v3 -= v4
    v4 = v3
    if v2 {
        fmt.Println(v7)
    }
    if (42 != 1000) {
        fmt.Println(v7)
        v8 := GenS1{F0: 0, F1: 35}
        v9 := GenS2{F0: 898, F1: v7}
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        v11 := GenS1{F0: 2, F1: 42}
        _ = fi10
        _ = v11
    }
    switch "" {
    case "\n\t":
        for fi12 := 0; fi12 < 14; fi12++ {
            v7 += v7
            v13 := GenS2{F0: 1, F1: (^(v7))}
            _ = fi12
            _ = v13
        }
    case "a":
        _ = len(v1)
    case "foo bar":
        if len(v1) > 0 {
            v1[0] = -838
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
