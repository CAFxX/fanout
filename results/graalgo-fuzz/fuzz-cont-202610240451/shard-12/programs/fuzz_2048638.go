package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int64 }
type GenS2 struct { F0 int32; F1 int }
func genFn0() uint {
    v4 := []int{42, 9223372036854775807}
    for fi5 := 0; fi5 < 8; fi5++ {
        fmt.Println(fi5)
        _ = fi5
    }
    _ = v4
    return 255
}
func genFn1(p0 int32) int64 {
    fmt.Println(p0)
    if (true && (-1 == 9223372036854775807)) {
        v6 := []int{-907, 2147483647}
        _ = v6
    }
    p0 += p0
    _ = p0
    return -129
}

func main() {
    var v1 int64 = 42
    var v2 uint8 = 255
    var v3 uint64 = 1
    for fi7 := 0; fi7 < 2; fi7++ {
        v1 = (v1 >> ((genFn0()) & 63))
        v8 := GenS1{F0: v1, F1: 100, F2: v1}
        v1 -= 42
        _ = fi7
        _ = v8
    }
    var v9 int32 = -2147483648
    v1 -= genFn1((v9 % ((v9) | 1)))
    if true {
        v2 += v2
    } else {
        var v10 int64 = (^(v1))
        v9 += 786
        var v11 bool = (!false)
        _ = v10
        _ = v11
    }
    v12 := []int{-296, 127, 4294967295}
    var v13 uint32 = 2147483647
    if ((((v3 % (((^(v3))) | 1)) != v3) || (-1 != -128)) || (!false)) {
        var v14 string = "a"
        v12 = append(v12, -472)
        _ = v14
    }
    v1 = (-(v1))
    switch "\n\t" {
    case "foo bar":
        v15 := GenS2{F0: (v9 & (v9 % (((v9 ^ (v9 % (((v9 % ((v9) | 1))) | 1)))) | 1))), F1: 128}
        _ = v15
    case "\n\t":
        v16 := GenS2{F0: v9, F1: -2147483648}
        _ = v16
    case "a":
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
    }
    var v17 uint64 = (v3 | 18446744073709551615)
    v1 = 1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v9
}
