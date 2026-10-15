package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 int8; F1 int8; F2 int16 }
func genFn0(p0 int32) int32 {
    var v7 int32 = 1
    var v8 string = "\n\t"
    var v9 int16 = 42
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return (-(v7))
}

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = false
    var v3 int8 = 109
    var v4 uint8 = 0
    var v5 int32 = -724
    var v6 uint32 = 65535
    v3 -= (v3 + (v3 % ((v3) | 1)))
    v10 := GenS1{F0: 1}
    var v11 uint8 = 42
    for fi12 := 0; fi12 < 12; fi12++ {
        v13 := []int{}
        _ = fi12
        _ = v13
    }
    v11 = v4
    v2 = v2
    var v14 float64 = (((-2.25 + (1e10 / 1.5)) * 3.14159) - ((float64(-2.25) + (1.5 + 0.0)) - 3.14159))
    var v15 bool = ((^(v3)) < (v3 >> ((1) & 63)))
    v16 := []int{-9223372036854775808, 1, -2147483648, -1}
    v17 := GenS1{F0: 65535}
    for fi18 := 0; fi18 < 7; fi18++ {
        v19 := GenS2{F0: 100, F1: 42}
        _ = fi18
        _ = v19
    }
    if v2 {
        var v20 uint64 = 175
        fmt.Println(v5, v15, v2)
        for fi21 := 0; fi21 < 10; fi21++ {
            v1[fi21] = (v6 + v6)
            var v22 bool = (42 >= -128)
            _ = fi21
            _ = v22
        }
        _ = v20
    } else {
        for fi23 := 0; fi23 < 1; fi23++ {
            v24 := GenS1{F0: 0}
            if v2 {
                var v25 int8 = v3
                _ = v25
            } else {
                v15 = true
                fmt.Println(v5)
                v6 = v6
            }
            v26 := GenS3{F0: v3, F1: (v3 << ((4294967295) & 63)), F2: 399}
            _ = fi23
            _ = v24
            _ = v26
        }
        fmt.Println(v4)
        var v27 int = -2147483648
        _ = v27
    }
    var v28 int8 = v3
    for fi29 := 0; fi29 < 10; fi29++ {
        v2 = (!true)
        _ = fi29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
