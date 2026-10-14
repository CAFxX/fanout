package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint8; F1 uint64; F2 int8 }
type GenS3 struct { F0 uint64; F1 int32 }
func genFn0(p0 int16, p1 int16) int {
    p0 = (p1 + 1000)
    var v4 int16 = (p0 / (((p0 + 1)) | 1))
    fmt.Println(v4)
    v4 += (p0 * (p1 + p1))
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v4
    return 1
}
func genFn1(p0 uint64) int16 {
    p0 = (^(p0))
    var v5 bool = (2147483647 > 128)
    switch ("" + "foo bar") {
    case "a":
        fmt.Println(v5)
    }
    v6 := GenS1{F0: 42, F1: -1, F2: 2}
    var v7 float32 = 0.1
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 int32 = 2147483647
    var v2 int16 = 42
    var v3 int = -119
    for fi8 := 0; fi8 < 13; fi8++ {
        for fi9 := 0; fi9 < 1; fi9++ {
            v10 := []int{genFn0((^(v2)), (v2 % ((v2) | 1))), fi8, (fi8 - v3)}
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    if false {
        for fi11 := 0; fi11 < 18; fi11++ {
            v3 -= (-(fi11))
            v12 := []int{}
            var v13 uint64 = 0
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    var v14 uint32 = 2147483648
    for fi15 := 0; fi15 < 14; fi15++ {
        var v16 uint = 4294967295
        v17 := GenS2{F0: 84, F1: 42, F2: -128}
        _ = fi15
        _ = v16
        _ = v17
    }
    switch ("foo bar" + "") {
    case "\n\t":
        if false {
            fmt.Println(v14, v3, v14)
            var v18 float64 = (float64(-2.25) - 0.1)
            v18 -= 3.14159
            _ = v18
        }
    }
    v14 = 240
    v1 = (v1 ^ v1)
    v19 := GenS3{F0: 4294967295, F1: (v1 ^ (v1 >> ((0) & 63)))}
    if ((v1 * (v1 * (v1 << ((65535) & 63)))) == (v1 & 0)) {
        v20 := []int{(v3 % ((v3) | 1)), v3}
        _ = v20
    } else {
        v21 := GenS1{F0: 0, F1: int32(v14), F2: (v14 + v14)}
        v1 = v1
        _ = v21
    }
    var v22 uint = 820
    v23 := []int{570, genFn0(1, v2), (v3 * genFn0((v2 << ((255) & 63)), v2)), v3}
    var v24 uint32 = v14
    if len(v23) > 0 {
        v23[0] = (v3 - -1)
    }
    v14 = v24
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
