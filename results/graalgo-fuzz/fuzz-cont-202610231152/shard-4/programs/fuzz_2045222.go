package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 bool; F1 uint; F2 bool }
func genFn0() int16 {
    var v7 uint = 672
    v7 = (v7 / (((^(v7))) | 1))
    var v8 uint8 = 2
    _ = v7
    _ = v8
    return -1
}
func genFn1(p0 uint64) int {
    v9 := []int{-1, 4294967295, -9223372036854775808, 32767}
    p0 -= p0
    if len(v9) > 0 {
        v9[0] = 2147483647
    }
    p0 -= 790
    _ = p0
    _ = v9
    return -453
}

func main() {
    var v1 uint16 = 65535
    var v2 uint64 = 860
    var v3 float64 = 0.0
    var v4 int = 737
    v5 := make(map[string]bool)
    v6 := []int{}
    switch -9223372036854775808 {
    case 2147483647:
        for fi10 := 0; fi10 < 8; fi10++ {
            fmt.Println(v4, v3)
            _ = fi10
        }
    }
    if true {
        if true {
            v11 := GenS1{F0: false, F1: (3.14159 <= 0.0)}
            _ = v11
        } else {
            fmt.Println(v3, v1, v2)
            v4 = (v4 & (v4 ^ (v4 << ((616) & 63))))
        }
    }
    v12 := GenS1{F0: (false && ((v4 / (((v4 >> ((897) & 63))) | 1)) == (v4 ^ (v4 ^ (v4 - -1))))), F1: (v3 == (v3 + float64((-2.25 * 3.14159))))}
    var v13 int16 = 0
    for fi14 := 0; fi14 < 3; fi14++ {
        switch genFn1(42) {
        case 377:
            fmt.Println(v1)
        case 127:
            var v15 float32 = float32(1e10)
            _ = v15
        }
        _ = fi14
    }
    v16 := GenS2{F0: (799 >= -9223372036854775808), F1: 42, F2: (v2 <= v2)}
    v5[("a" + "hello")] = (((v4 << ((0) & 63)) >= v4) && ((((v3 / v3) <= v3) || false) || ((v13 % (((v13 % (((^(v13))) | 1))) | 1)) >= v13)))
    var v17 uint32 = 2
    v13 = (v13 + (v13 * genFn0()))
    for fi18 := 0; fi18 < 18; fi18++ {
        fmt.Println(v1, v1)
        fmt.Println(v4, v3)
        _ = fi18
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        var v20 bool = (26 < 100)
        switch 9223372036854775807 {
        case 1:
            v3 += 0.1
        }
        var v21 int16 = (v13 - 1)
        _ = fi19
        _ = v20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
