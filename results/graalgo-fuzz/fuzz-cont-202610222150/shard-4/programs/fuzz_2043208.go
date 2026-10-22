package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 int8; F1 int; F2 uint32 }
func genFn0(p0 uint) int {
    switch "foo bar" {
    case "hello":
        v7 := GenS1{F0: 65535, F1: 156}
        _ = v7
    case "\n\t":
        p0 += (^(p0))
    case "":
        var v8 int16 = 1000
        _ = v8
    }
    p0 += (p0 & p0)
    fmt.Println(p0)
    _ = p0
    return 0
}

func main() {
    var v1 uint = 1
    var v2 uint64 = 42
    var v3 uint8 = 2
    v4 := make(map[string]uint32)
    v5 := make(map[int]bool)
    v6 := []int{-1, 32767, 4294967295}
    v9 := GenS2{F0: 288, F1: (-1 >= 0)}
    for fi10 := 0; fi10 < 16; fi10++ {
        var v11 int32 = -1
        delete(v4, "")
        _ = fi10
        _ = v11
    }
    v12 := GenS3{F0: 0, F1: genFn0(uint(v2)), F2: 0}
    for fi13 := 0; fi13 < 17; fi13++ {
        var v14 bool = true
        _ = fi13
        _ = v14
    }
    v1 -= (^(v1))
    v3 = (v3 % ((v3) | 1))
    for fi15 := 0; fi15 < 17; fi15++ {
        var v16 uint64 = (v2 + 0)
        _ = fi15
        _ = v16
    }
    v1 = 65535
    v17 := GenS3{F0: 0, F1: 1, F2: 4294967295}
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
