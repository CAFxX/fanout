package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int32 }

func main() {
    v1 := []int{-1, -9223372036854775808, 9223372036854775807, 32767}
    v2 := make(map[string]int)
    v3 := []int{127}
    v4 := make(map[uint32]bool)
    var v5 uint16 = 842
    var v6 string = ""
    if false {
        v7 := GenS1{F0: 973, F1: 4294967295, F2: 2147483647}
        var v8 int = 881
        delete(v2, v6)
        _ = v7
        _ = v8
    }
    v5 -= v5
    for fi9 := 0; fi9 < 3; fi9++ {
        var v10 int16 = 26
        v10 = (v10 + (v10 + v10))
        _ = fi9
        _ = v10
    }
    var v11 int8 = 0
    if true {
        v6 = v6
        v11 = (v11 % ((int8(v5)) | 1))
        v5 += v5
    } else {
        var v12 int = -2147483648
        _ = v12
    }
    v11 = v11
    v3 = append(v3, 42)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
