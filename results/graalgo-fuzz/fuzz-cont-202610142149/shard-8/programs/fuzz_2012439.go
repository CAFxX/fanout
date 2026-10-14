package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
func genFn0(p0 uint32, p1 int16, p2 uint32) uint32 {
    p0 += p0
    p1 += p1
    v5 := []int{2147483647, -18, 797}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (p2 >> ((717) & 63))
}
func genFn1(p0 int32, p1 int32, p2 bool) int8 {
    v6 := GenS1{F0: 20, F1: 42}
    v7 := []int{-315}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 uint32 = 562
    v2 := make(map[int]string)
    var v3 int32 = 0
    v4 := make(map[string]bool)
    if (0 >= 260) {
        var v8 int = 966
        _ = v8
    } else {
        v1 -= v1
        var v9 uint8 = 1
        for fi10 := 0; fi10 < 1; fi10++ {
            var v11 uint64 = 42
            _ = fi10
            _ = v11
        }
        _ = v9
    }
    _ = len(v2)
    v12 := GenS1{F0: 127, F1: genFn1((v3 + v3), -1, false)}
    fmt.Println(v1, v3)
    v13 := GenS1{F0: 2, F1: genFn1((-(v3)), 2147483647, true)}
    for fi14 := 0; fi14 < 8; fi14++ {
        var v15 string = "hello"
        var v16 bool = (genFn1(v3, (v3 << ((42) & 63)), true) > 82)
        v3 = (-(v3))
        _ = fi14
        _ = v15
        _ = v16
    }
    v3 = v3
    if true {
        fmt.Println(v1)
        v1 = 2
    } else {
        v17 := []int{-9223372036854775808}
        v3 = -2147483648
        fmt.Println(v3)
        _ = v17
    }
    v3 += v3
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
