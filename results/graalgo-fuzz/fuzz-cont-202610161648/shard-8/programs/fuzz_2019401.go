package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 int }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 uint64; F1 uint8; F2 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int16; F1 int64 }
func genFn0() int8 {
    for fi4 := 0; fi4 < 20; fi4++ {
        v5 := GenS1{F0: 197, F1: 127, F2: -983}
        if false {
            var v6 int = 2147483647
            v7 := GenS2{F0: 0, F1: 18446744073709551615}
            v6 = (fi4 / (((fi4 ^ fi4)) | 1))
            _ = v6
            _ = v7
        }
        _ = fi4
        _ = v5
    }
    switch 127 {
    case 2147483647:
        var v8 uint8 = 127
        _ = v8
    case 42:
        v9 := GenS1{F0: -468, F1: 2, F2: -871}
        _ = v9
    case 1:
        var v10 float32 = 0.1
        _ = v10
    default:
        var v11 int16 = 0
        _ = v11
    }
    v12 := GenS2{F0: 0, F1: 2}
    var v13 int32 = 348
    _ = v12
    _ = v13
    return 42
}
func genFn1(p0 int32, p1 int64, p2 bool) uint16 {
    p1 += (p1 + -1)
    v14 := GenS3{F0: 325, F1: 2, F2: 76}
    p1 = p1
    switch (("hello" + "hello") + ("\n\t" + "a")) {
    case "a":
        v15 := []int{4294967295, -299, 128, -439}
        _ = v15
    case "hello":
        v16 := GenS4{F0: ((988 == 32767) && (2 == 2))}
        _ = v16
    case "\n\t":
        fmt.Println(p0)
    default:
        switch "\n\t" {
        case "a":
            v17 := GenS2{F0: 42, F1: 18446744073709551615}
            _ = v17
        case "\n\t":
            p0 = (p0 % ((p0) | 1))
        case "hello":
            fmt.Println(p0, p0, p2)
        default:
            v18 := GenS5{F0: -646, F1: p1}
            _ = v18
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 32767
}

func main() {
    var v1 float32 = 0.1
    var v2 uint32 = 1
    var v3 uint32 = (v2 / (((v2 << ((65535) & 63))) | 1))
    for fi19 := 0; fi19 < 5; fi19++ {
        v20 := []int{fi19}
        v2 = 42
        _ = fi19
        _ = v20
    }
    var v21 uint16 = genFn1(0, -1, false)
    v22 := GenS3{F0: 1, F1: 1, F2: 1}
    var v23 int64 = -9223372036854775808
    var v24 int = 1
    v25 := []int{}
    if len(v25) >= 2 {
        v25 = v25[1:len(v25)]
    }
    _ = len(v25)
    v26 := GenS3{F0: 1, F1: 128, F2: 0}
    if false {
        v25 = append(v25, -9223372036854775808)
        v27 := GenS5{F0: 1000, F1: 42}
        _ = v27
    } else {
        v3 = v2
        if true {
            var v28 int8 = -1
            var v29 int32 = -2147483648
            _ = v28
            _ = v29
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
