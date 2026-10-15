package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint16 }
func genFn0() uint64 {
    var v5 int = 32768
    fmt.Println(v5)
    fmt.Println(v5)
    v6 := GenS1{F0: v5, F1: 1}
    var v7 uint8 = 191
    _ = v5
    _ = v6
    _ = v7
    return 695
}
func genFn1(p0 int16, p1 uint8) string {
    var v8 int16 = (p0 - 32767)
    v9 := GenS2{F0: 32767}
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 uint32 = 0
        for fi12 := 0; fi12 < 9; fi12++ {
            v8 = 0
            _ = fi12
        }
        p0 = (v8 ^ p0)
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return "a"
}

func main() {
    var v1 int64 = 0
    v2 := make(map[string]uint32)
    var v3 uint = 4294967295
    v4 := make(map[int]bool)
    var v13 uint8 = 73
    var v14 uint64 = 4294967295
    fmt.Println(v13)
    var v15 int = 1
    for fi16 := 0; fi16 < 9; fi16++ {
        v17 := GenS1{F0: -7, F1: 1}
        for fi18 := 0; fi18 < 11; fi18++ {
            v19 := GenS1{F0: fi16, F1: 42}
            v20 := []int{(v15 / (((fi18 ^ 0)) | 1)), -2147483648}
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = fi16
        _ = v17
    }
    v21 := []int{v15, -721}
    _ = len(v2)
    if ((true && (true || true)) || (42 < int64(v14))) {
        var v22 bool = false
        v1 -= v1
        _ = len(v4)
        _ = v22
    }
    fmt.Println(v1)
    v15 += v15
    v13 = 128
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
