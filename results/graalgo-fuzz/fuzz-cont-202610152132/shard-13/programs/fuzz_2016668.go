package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 uint8; F2 uint16 }
func genFn0(p0 int16) uint16 {
    p0 = (p0 | p0)
    v4 := []int{0, 584}
    for fi5 := 0; fi5 < 8; fi5++ {
        p0 = p0
        _ = fi5
    }
    v6 := GenS1{F0: 2147483648}
    v4 = append(v4, 88)
    _ = p0
    _ = v4
    _ = v6
    return 1
}
func genFn1(p0 string) uint64 {
    v7 := []int{4294967295, -124, 42, 38}
    var v8 uint8 = 2
    switch p0 {
    case "\n\t":
        v9 := GenS2{F0: 2, F1: (518 >= 1)}
        _ = v9
    default:
        v8 -= v8
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = v7
    _ = v8
    return 2
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{2147483647}
    v3 := []string{}
    var v10 float32 = 0.1
    v10 -= v10
    _ = len(v3)
    fmt.Println(v10)
    if ((356 == 42) && false) {
        if len(v2) > 0 {
            v2[0] = -592
        }
    }
    var v11 int8 = 1
    fmt.Println(v11)
    v10 += (1.5 / 1e10)
    v10 += v10
    fmt.Println(v11)
    v12 := GenS3{F0: 1}
    v13 := GenS1{F0: 1}
    v11 = -45
    v1[""] = 171
    v14 := GenS4{F0: 65535, F1: 255, F2: genFn0(0)}
    var v15 bool = false
    var v16 uint8 = 127
    switch "hello" {
    case "\n\t":
        v17 := GenS3{F0: 111}
        _ = v17
    default:
        v18 := GenS3{F0: (v16 << ((42) & 63))}
        _ = v18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
