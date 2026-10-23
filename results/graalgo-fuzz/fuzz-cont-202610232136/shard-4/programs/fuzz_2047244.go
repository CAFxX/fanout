package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 uint8 }

func main() {
    var v1 int16 = 0
    var v2 bool = true
    var v3 uint8 = 255
    v4 := []int{2147483647, 127, 421, -9223372036854775808}
    var v5 uint = 1
    v6 := []string{"", "hello"}
    v7 := GenS1{F0: v1}
    var v8 int = -2147483648
    var v9 int64 = 9223372036854775807
    v10 := GenS2{F0: 226, F1: v3}
    var v11 uint64 = 2
    fmt.Println(v5, v11)
    var v12 int = (^(v8))
    if len(v6) > 0 {
        v6[0] = "a"
    }
    v8 = (v12 / ((v8) | 1))
    v11 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
