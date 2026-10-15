package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 int16 }
func genFn0(p0 uint16) int64 {
    p0 -= p0
    p0 = 42
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    return -9223372036854775808
}
func genFn1(p0 uint16, p1 int32) uint16 {
    for fi7 := 0; fi7 < 17; fi7++ {
        p1 += 42
        _ = fi7
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        v9 := []int{(fi8 | fi8), (fi8 & (fi8 & fi8)), fi8, fi8}
        _ = fi8
        _ = v9
    }
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    return (p0 / ((2) | 1))
}

func main() {
    v1 := make(map[string]string)
    var v2 string = ("hello" + "")
    var v3 float32 = 1.5
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]string)
    v6 := make(map[int]string)
    v3 += (v3 * v3)
    v3 = 3.14159
    v2 = v2
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 uint32 = 2147483648
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 8; fi12++ {
        v2 = (v2 + v2)
        _ = fi12
    }
    fmt.Println(v2)
    for fi13 := 0; fi13 < 13; fi13++ {
        var v14 uint16 = 121
        var v15 int64 = -1
        _ = fi13
        _ = v14
        _ = v15
    }
    v16 := GenS1{F0: 0, F1: 65535, F2: -36}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
