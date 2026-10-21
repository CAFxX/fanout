package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint64 }

func main() {
    v1 := make(map[string]string)
    var v2 float64 = 3.14159
    v3 := make(map[int]bool)
    var v4 uint32 = 893
    if (4294967295 > 1) {
        if ("hello" >= "\n\t") {
            fmt.Println(v4, v4)
            v5 := []int{-2147483648, 857}
            _ = v5
        } else {
            v4 = (v4 & (v4 + v4))
        }
    }
    v4 += v4
    v6 := GenS1{F0: 42, F1: 9223372036854775807, F2: 886}
    var v7 string = "a"
    var v8 int64 = 42
    if true {
        v4 += v4
        fmt.Println(v7)
        delete(v3, 9223372036854775807)
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 uint8 = 190
        _ = fi9
        _ = v10
    }
    var v11 int = -634
    v12 := GenS1{F0: 1, F1: (v8 | (-(v8))), F2: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
