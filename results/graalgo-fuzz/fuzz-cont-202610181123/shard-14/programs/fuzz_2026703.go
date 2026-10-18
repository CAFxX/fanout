package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 int16; F1 uint16 }
func genFn0(p0 int, p1 int16) uint {
    fmt.Println(p1, p0)
    p0 += p0
    p1 = p1
    var v4 uint = 1
    v5 := []int{p0, (^(p0)), (p0 + p0)}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return (v4 * v4)
}

func main() {
    var v1 int8 = 1
    var v2 int8 = (^(v1))
    var v3 uint32 = 977
    v6 := GenS1{F0: 728, F1: 0}
    v2 = (v1 | v2)
    var v7 int64 = 1
    v3 = v3
    if true {
        fmt.Println(v1, v2)
    } else {
        for fi8 := 0; fi8 < 6; fi8++ {
            var v9 uint = genFn0(fi8, 0)
            v3 -= v3
            v10 := GenS1{F0: (fi8 | fi8), F1: 4294967295}
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    fmt.Println(v2, v1)
    var v11 uint16 = 1
    v12 := GenS2{F0: 727, F1: (v11 & 65535)}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
