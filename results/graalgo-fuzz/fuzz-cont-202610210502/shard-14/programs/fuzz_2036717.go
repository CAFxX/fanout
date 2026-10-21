package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 string; F1 uint64; F2 uint8 }
func genFn0(p0 uint16, p1 uint8) bool {
    v6 := GenS1{F0: 42, F1: "a"}
    v7 := []int{-250}
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return true
}

func main() {
    var v1 uint32 = 42
    v2 := make(map[string]int)
    v3 := make(map[uint32]int)
    v4 := []uint32{(v1 & (v1 << ((4294967295) & 63))), (^(v1))}
    v5 := make(map[int]uint32)
    fmt.Println(v1)
    v1 = (v1 % ((2) | 1))
    var v8 int32 = 0
    v9 := GenS2{F0: "hello", F1: 51, F2: 100}
    v8 = v8
    v8 = (v8 / (((v8 / ((1) | 1))) | 1))
    for fi10 := 0; fi10 < 16; fi10++ {
        fmt.Println(fi10, fi10, v8)
        fmt.Println(v1)
        _ = fi10
    }
    var v11 int16 = 541
    v3[2147483648] = 0
    var v12 uint16 = 32767
    for fi13 := 0; fi13 < 11; fi13++ {
        delete(v2, "foo bar")
        _ = fi13
    }
    v14 := GenS1{F0: 233, F1: "hello"}
    for fi15 := 0; fi15 < 7; fi15++ {
        if (0 < 4294967295) {
            var v16 int32 = (v8 ^ v8)
            _ = v16
        }
        fmt.Println(v12, v8, v8)
        _ = fi15
    }
    fmt.Println(v11, v11, v12)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
