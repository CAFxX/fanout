package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 uint16; F1 uint; F2 int64 }
func genFn0(p0 int32, p1 uint64) uint16 {
    for fi4 := 0; fi4 < 11; fi4++ {
        var v5 float64 = (-2.25 / 1e10)
        if (("hello" + "\n\t") <= "foo bar") {
            switch (fi4 + 2147483647) {
            case -399:
                var v6 float64 = v5
                _ = v6
            case -9223372036854775808:
                v7 := []int{(fi4 | -1), (fi4 - fi4), fi4}
                _ = v7
            }
        }
        _ = fi4
        _ = v5
    }
    v8 := GenS1{F0: p0, F1: 127, F2: 1}
    v9 := GenS2{F0: -2147483648, F1: 65535}
    if false {
        switch 0 {
        case 9223372036854775807:
            v10 := GenS3{F0: 48, F1: 2, F2: -9223372036854775808}
            _ = v10
        }
        for fi11 := 0; fi11 < 13; fi11++ {
            var v12 uint8 = 1
            v13 := []int{fi11}
            v14 := GenS3{F0: uint16(v12), F1: 65535, F2: 42}
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        p0 += (p0 - -1)
    }
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 2
}
func genFn1(p0 uint) int8 {
    v15 := []int{4294967295}
    v15 = append(v15, -349)
    _ = p0
    _ = v15
    return 63
}

func main() {
    var v1 int64 = -271
    v2 := []uint32{2147483647, 41}
    v3 := make(map[int]bool)
    var v16 bool = (false && true)
    for fi17 := 0; fi17 < 3; fi17++ {
        fmt.Println(fi17, v16, v1)
        var v18 int64 = v1
        _ = fi17
        _ = v18
    }
    switch "\n\t" {
    case "\n\t":
        _ = len(v2)
    case "hello":
        v16 = true
    case "foo bar":
        v19 := GenS1{F0: -1, F1: 0, F2: 668}
        _ = v19
    default:
        var v20 uint = 1
        _ = v20
    }
    if (65535 <= 0) {
        _ = len(v2)
        fmt.Println(v1, v16)
    } else {
        v16 = v16
    }
    fmt.Println(v16, v16)
    var v21 int8 = 0
    if v16 {
        v21 = -128
    }
    var v22 string = "hello"
    var v23 uint64 = 1
    var v24 float32 = (float32(float64((-2.25 * (1.5 + (float32(-2.25) + 1e10))))) - -2.25)
    v25 := GenS2{F0: 1, F1: 2}
    v26 := GenS2{F0: -1, F1: 607}
    v1 -= v1
    v24 -= (((v24 + v24) / v24) / v24)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
