package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint8 }
type GenS2 struct { F0 int }
func genFn0(p0 uint32) uint8 {
    p0 -= p0
    var v6 float64 = (((((-2.25 * 3.14159) * -2.25) * 1.5) * 1.5) + (-2.25 + 1.5))
    _ = p0
    _ = v6
    return 158
}
func genFn1() uint32 {
    v7 := []int{32768, 127, -268, 421}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    _ = v7
    return 30
}

func main() {
    var v1 int16 = -32768
    v2 := []int{}
    var v3 float64 = (3.14159 * 3.14159)
    v4 := make(map[uint32]int)
    var v5 int16 = v1
    for fi8 := 0; fi8 < 5; fi8++ {
        switch "a" {
        case "foo bar":
            v3 -= -2.25
        case "":
            v9 := GenS1{F0: 994, F1: -2147483648, F2: 19}
            _ = v9
        case "\n\t":
            v1 -= v5
        }
        switch fi8 {
        case 459:
            v1 += v5
        case 440:
            var v10 int8 = -128
            _ = v10
        }
        _ = fi8
    }
    var v11 int32 = 42
    if (!((v11 % ((v11) | 1)) <= 42)) {
        var v12 uint8 = 74
        _ = v12
    } else {
        var v13 uint = 709
        _ = v13
    }
    var v14 int64 = 663
    var v15 float64 = 0.0
    if false {
        if ("hello" > "") {
            v16 := GenS1{F0: 0, F1: (v11 ^ v11), F2: genFn0(1)}
            _ = v16
        }
        var v17 int8 = 127
        _ = v17
    } else {
        var v18 uint = 255
        switch 42 {
        case -9223372036854775808:
            var v19 int32 = v11
            _ = v19
        case 9223372036854775807:
            var v20 int = int(v14)
            _ = v20
        case -513:
            v21 := GenS1{F0: v18, F1: 2147483647, F2: 255}
            _ = v21
        }
        _ = v18
    }
    v5 -= v5
    fmt.Println(v11, v1, v11)
    var v22 float64 = -2.25
    v5 = v1
    _ = len(v2)
    for fi23 := 0; fi23 < 12; fi23++ {
        v14 += (v14 - (v14 * (v14 * (^(v14)))))
        v5 -= -32768
        switch (fi23 & 143) {
        case -943:
            var v24 int = fi23
            _ = v24
        case -2147483648:
            var v25 uint64 = 18446744073709551615
            _ = v25
        }
        _ = fi23
    }
    v14 = v14
    v15 -= (v15 + float64(1e10))
    switch "\n\t" {
    case "hello":
        switch "\n\t" {
        case "foo bar":
            for fi26 := 0; fi26 < 9; fi26++ {
                v11 += (v11 << ((368) & 63))
                _ = fi26
            }
        case "a":
            for fi27 := 0; fi27 < 9; fi27++ {
                v15 += 3.14159
                fmt.Println(v1, v22)
                _ = fi27
            }
        }
    default:
        var v28 bool = true
        _ = v28
    }
    v29 := GenS2{F0: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
