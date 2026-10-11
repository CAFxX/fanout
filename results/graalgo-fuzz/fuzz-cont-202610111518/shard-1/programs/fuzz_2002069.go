package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 bool }
type GenS2 struct { F0 int64; F1 int64 }
func genFn0(p0 int32, p1 string, p2 uint) int16 {
    for fi5 := 0; fi5 < 5; fi5++ {
        v6 := GenS1{F0: 2, F1: 255, F2: false}
        var v7 int = (fi5 << (((p2 - 1)) & 63))
        var v8 uint8 = 128
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        var v10 int16 = -370
        _ = fi9
        _ = v10
    }
    var v11 float64 = -2.25
    p2 += p2
    var v12 uint32 = 2147483648
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 0
}
func genFn1() uint16 {
    if (32767 >= 0) {
        var v13 uint = 433
        v14 := []int{4294967295, 42, 138}
        var v15 int16 = 1
        _ = v13
        _ = v14
        _ = v15
    } else {
        if (("foo bar" <= "") && true) {
            v16 := []int{1, 9223372036854775807, 32767}
            _ = v16
        }
    }
    var v17 int = -2147483648
    _ = v17
    return 32767
}

func main() {
    var v1 uint8 = 128
    var v2 uint = 0
    var v3 float32 = (1.5 / 0.1)
    var v4 int = 127
    var v18 float64 = -2.25
    v19 := GenS1{F0: 255, F1: 469, F2: true}
    fmt.Println(v2, v1)
    v20 := []int{42, (v4 >> (((v2 & (v2 >> ((4294967295) & 63)))) & 63)), v4}
    v21 := GenS2{F0: 42, F1: 1}
    for fi22 := 0; fi22 < 7; fi22++ {
        var v23 int8 = -128
        fmt.Println(v23)
        _ = fi22
        _ = v23
    }
    switch (^(v4)) {
    case -343:
        v1 = (v1 & (^(v1)))
    case 9223372036854775807:
        var v24 int16 = 0
        _ = v24
    }
    v20 = append(v20, v4)
    var v25 uint64 = 210
    switch "hello" {
    case "":
        _ = len(v20)
    case "hello":
        fmt.Println(v25, v2)
    }
    var v26 int64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
