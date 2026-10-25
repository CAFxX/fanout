package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }

func main() {
    var v1 float32 = (1e10 / 1.0)
    v2 := []int{42, 0, -1, -737}
    var v3 int64 = 9223372036854775807
    v4 := make(map[string]int)
    v5 := []int{-376, -1}
    fmt.Println(v1, v3)
    var v6 uint16 = 42
    v7 := GenS1{F0: 32767, F1: 986}
    v4["foo bar"] = 2147483647
    v1 += v1
    if true {
        switch "a" {
        case "\n\t":
            v1 = 1.5
        }
    }
    var v8 bool = (2 >= 2)
    v3 += v3
    v1 -= v1
    v1 = v1
    v2 = append(v2, -2147483648)
    v8 = (v8 || true)
    var v9 int64 = v3
    v10 := GenS1{F0: (^(v6)), F1: 353}
    v1 -= 1.5
    var v11 uint16 = v6
    for fi12 := 0; fi12 < 8; fi12++ {
        fmt.Println(v11, v6, fi12)
        _ = fi12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
