package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }

func main() {
    var v1 uint16 = 509
    v2 := make(map[int]uint32)
    v3 := make(map[string]uint32)
    if false {
        for fi4 := 0; fi4 < 7; fi4++ {
            v5 := []int{-94, fi4}
            v6 := GenS1{F0: 2147483648, F1: (fi4 & (fi4 % ((-2147483648) | 1)))}
            delete(v3, "a")
            _ = fi4
            _ = v5
            _ = v6
        }
        v1 += (v1 | v1)
    }
    if (100 != -1) {
        fmt.Println(v1)
    }
    var v7 uint32 = 65535
    v7 += v7
    v7 += v7
    v8 := GenS1{F0: 910, F1: 4294967295}
    v9 := []int{-1, 42, 32768}
    v10 := v2[-33]
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
