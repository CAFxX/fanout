package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 uint }

func main() {
    v1 := []int{470}
    v2 := make(map[string]int)
    v3 := make(map[string]int)
    v4 := make(map[int]string)
    var v5 uint64 = 852
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 bool = false
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        v5 -= 18446744073709551615
        _ = fi8
    }
    v5 += (v5 | v5)
    fmt.Println(v5)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v9 := GenS1{F0: -2147483648, F1: -1, F2: 18446744073709551615}
    v1 = append(v1, -1)
    v10 := GenS1{F0: -2147483648, F1: 100, F2: 65535}
    var v11 bool = (!(!false))
    var v12 bool = v11
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
