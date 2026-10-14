package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
func genFn0(p0 int32, p1 int16, p2 int64) int16 {
    fmt.Println(p1)
    switch "\n\t" {
    case "a":
        p0 = -419
    case "hello":
        v7 := GenS1{F0: (p0 + (p0 - (p0 ^ -694))), F1: 1}
        _ = v7
    }
    switch 0 {
    case 35:
        fmt.Println(p1)
    default:
        for fi8 := 0; fi8 < 18; fi8++ {
            fmt.Println(p1, fi8, p1)
            _ = fi8
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return p1
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]string)
    var v3 int32 = -951
    var v4 int8 = 47
    var v5 uint32 = 894
    var v6 uint64 = 52
    delete(v1, 861)
    v5 += (v5 | (v5 / (((v5 << ((0) & 63))) | 1)))
    v4 += v4
    v9 := v2[-786]
    var v10 int = -1
    v11 := GenS1{F0: (v3 >> ((0) & 63)), F1: (v4 ^ (v4 << ((2) & 63)))}
    var v12 uint32 = (v5 ^ 2147483648)
    fmt.Println(v5, v5, v6)
    v10 = (v10 << ((42) & 63))
    var v13 bool = true
    var v14 float64 = 0.0
    v15 := []int{595, (v10 % ((128) | 1))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
