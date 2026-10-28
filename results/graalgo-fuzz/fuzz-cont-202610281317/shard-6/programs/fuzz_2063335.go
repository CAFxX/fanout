package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]string)
    var v3 int64 = 9223372036854775807
    var v4 int16 = 993
    v5 := []string{}
    v4 = (^(v4))
    v4 = v4
    v4 += v4
    v6 := GenS1{F0: (v3 - v3)}
    var v7 uint16 = 787
    v8 := GenS1{F0: -412}
    _ = len(v2)
    fmt.Println(v7)
    var v9 uint32 = 2147483648
    v3 -= v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
