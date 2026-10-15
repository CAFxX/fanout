package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint64 }
func genFn0(p0 int) int16 {
    fmt.Println(p0)
    v4 := []int{-2147483648}
    switch "" {
    case "\n\t":
        p0 -= p0
    }
    _ = p0
    _ = v4
    return -426
}
func genFn1() uint8 {
    var v5 int32 = -1
    v6 := GenS1{F0: 1, F1: -27, F2: 18446744073709551615}
    fmt.Println(v5)
    _ = v5
    _ = v6
    return 128
}

func main() {
    var v1 int32 = 0
    v2 := []int{568, 0}
    var v3 uint64 = 1
    fmt.Println(v3)
    for fi7 := 0; fi7 < 10; fi7++ {
        fmt.Println(fi7, v3, fi7)
        _ = fi7
    }
    v8 := GenS1{F0: 4294967295, F1: 0, F2: (v3 + v3)}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi9 := 0; fi9 < 16; fi9++ {
        v10 := GenS1{F0: 2, F1: 1, F2: v3}
        for fi11 := 0; fi11 < 9; fi11++ {
            var v12 float32 = 3.14159
            _ = fi11
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    v3 = (^(v3))
    if len(v2) > 0 {
        v2[0] = 973
    }
    fmt.Println(v1, v3)
    v3 = (v3 >> ((255) & 63))
    for fi13 := 0; fi13 < 1; fi13++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v14 string = "\n\t"
        var v15 uint64 = 4294967295
        _ = fi13
        _ = v14
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v8
}
