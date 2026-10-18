package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
func genFn0(p0 uint, p1 bool) uint8 {
    p0 = 0
    fmt.Println(p0, p1)
    fmt.Println(p1)
    _ = p0
    _ = p1
    return 0
}

func main() {
    v1 := []uint32{2147483647, 721}
    var v2 int32 = 248
    v3 := make(map[int]bool)
    if true {
        fmt.Println(v2)
        v2 -= v2
        var v4 uint32 = 65535
        _ = v4
    } else {
        v2 += v2
        for fi5 := 0; fi5 < 7; fi5++ {
            var v6 float64 = 0.1
            _ = fi5
            _ = v6
        }
        var v7 int64 = 1
        _ = v7
    }
    _ = len(v1)
    var v8 int32 = 212
    var v9 uint = 255
    var v10 uint32 = 2147483648
    v11 := GenS1{F0: 0, F1: 4294967295}
    _ = len(v1)
    v12 := GenS1{F0: 100, F1: 2147483648}
    var v13 int = 2147483647
    var v14 uint64 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
