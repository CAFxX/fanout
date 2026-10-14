package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 uint = 42
    var v2 uint64 = 42
    v3 := make(map[string]string)
    v4 := make(map[int]bool)
    var v5 int8 = 127
    var v6 int64 = 0
    fmt.Println(v5, v6)
    v5 = (v5 / ((v5) | 1))
    v2 += v2
    switch "" {
    case "hello":
        var v7 int8 = 50
        _ = v7
    case "":
        if (2 <= v2) {
            v1 = uint(v2)
            v1 += (^(v1))
        }
    case "foo bar":
        v8 := []int{0, 128, 127, 57}
        _ = v8
    }
    v5 = 100
    if ((v5 + -23) < v5) {
        _ = len(v3)
        v9 := []int{42}
        v1 = v1
        _ = v9
    } else {
        for fi10 := 0; fi10 < 7; fi10++ {
            fmt.Println(v6, v5)
            v11 := []int{fi10, 0}
            _ = fi10
            _ = v11
        }
        fmt.Println(v2)
        v12 := GenS1{F0: (v1 >> (((v1 / (((v1 * (v1 | (v1 % ((397) | 1))))) | 1))) & 63))}
        _ = v12
    }
    var v13 uint64 = 463
    v2 = (v13 - 42)
    v5 += v5
    fmt.Println(v13, v6, v6)
    var v14 float64 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
