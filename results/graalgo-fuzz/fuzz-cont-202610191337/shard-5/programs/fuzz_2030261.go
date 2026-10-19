package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint64 }
func genFn0() int16 {
    v7 := []int{163, -2147483648}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    _ = v7
    return 205
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]int)
    v3 := []int{32768, 9223372036854775807, 4294967295}
    v4 := make(map[string]int)
    var v5 uint = 42
    var v6 uint64 = 894
    var v8 int8 = 89
    fmt.Println(v5, v6)
    fmt.Println(v6, v8)
    var v9 bool = ((true || (v8 == v8)) || false)
    v6 = (v6 + 1)
    var v10 int64 = 26
    fmt.Println(v5, v6, v6)
    v11 := GenS1{F0: "", F1: 1, F2: (v6 % ((297) | 1))}
    _ = len(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
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
    _ = v8
    _ = v9
}
