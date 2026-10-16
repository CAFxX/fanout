package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 int32 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]string)
    var v3 int8 = -1
    var v4 uint = 2
    v5 := []int{-9223372036854775808, -2147483648, 4294967295}
    var v6 int = 9223372036854775807
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 uint64 = 117
        if true {
            v4 += v4
        }
        v4 = (v4 % ((uint(v8)) | 1))
        _ = fi7
        _ = v8
    }
    v3 = (v3 * -34)
    v6 += 32768
    switch v6 {
    case 0:
        var v9 float64 = (1.5 * 1e10)
        _ = v9
    }
    var v10 int16 = 1
    v1[(v6 >> ((v4) & 63))] = 2147483647
    v10 -= v10
    fmt.Println(v6, v4, v10)
    var v11 float32 = 1.5
    v3 += v3
    v12 := GenS1{F0: 42, F1: 1, F2: 978}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
