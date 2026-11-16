package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 int64; F1 uint32; F2 uint64 }
func genFn0(p0 int32) int64 {
    p0 = p0
    if (false || true) {
        fmt.Println(p0)
        fmt.Println(p0)
    } else {
        fmt.Println(p0)
    }
    var v4 uint = 0
    switch int(v4) {
    case 106:
        v4 += v4
    }
    v5 := GenS1{F0: -128, F1: (p0 % (((p0 % ((p0) | 1))) | 1))}
    _ = p0
    _ = v4
    _ = v5
    return 42
}
func genFn1(p0 int64, p1 bool) int8 {
    var v6 int16 = 207
    v7 := []int{32767, 826, 32767, 2147483647}
    switch "\n\t" {
    case "hello":
        var v8 uint8 = 128
        _ = v8
    case "":
        var v9 int64 = p0
        _ = v9
    case "\n\t":
        var v10 uint16 = 0
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -128
}

func main() {
    v1 := []uint32{42, 4294967295}
    var v2 int64 = -1
    v3 := make(map[int]uint32)
    switch ("hello" + "a") {
    case "a":
        var v11 uint8 = 250
        _ = v11
    case "\n\t":
        for fi12 := 0; fi12 < 18; fi12++ {
            var v13 uint32 = 2147483648
            v14 := GenS1{F0: -128, F1: -2147483648}
            var v15 uint64 = 18446744073709551615
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
    }
    fmt.Println(v2)
    var v16 int8 = 100
    _ = len(v1)
    fmt.Println(v16)
    fmt.Println(v2)
    var v17 uint8 = 2
    if (true || (v2 < v2)) {
        if (("\n\t" == "foo bar") && false) {
            var v18 int32 = 1
            _ = v18
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v19 int64 = 9223372036854775807
            switch "a" {
            case "a":
                v17 = 2
            case "hello":
                v17 = (v17 / (((^(v17))) | 1))
            case "\n\t":
                for fi20 := 0; fi20 < 19; fi20++ {
                    var v21 float64 = 1.5
                    _ = fi20
                    _ = v21
                }
            }
            _ = v19
        }
    } else {
        for fi22 := 0; fi22 < 20; fi22++ {
            var v23 int32 = 42
            var v24 int32 = (v23 % (((v23 + 577)) | 1))
            _ = fi22
            _ = v23
            _ = v24
        }
        var v25 uint64 = 2
        _ = v25
    }
    for fi26 := 0; fi26 < 16; fi26++ {
        for fi27 := 0; fi27 < 14; fi27++ {
            v28 := GenS1{F0: (v16 << ((4294967295) & 63)), F1: -479}
            v16 += (v16 ^ (-(v16)))
            _ = fi27
            _ = v28
        }
        v29 := GenS2{F0: (v2 << ((1) & 63)), F1: 36, F2: 18446744073709551615}
        _ = fi26
        _ = v29
    }
    v30 := GenS2{F0: genFn0(2147483647), F1: 0, F2: 4294967295}
    var v31 int16 = 1000
    v31 += 0
    var v32 string = "foo bar"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
}
