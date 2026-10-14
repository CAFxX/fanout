package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 int64 }
type GenS2 struct { F0 uint }
func genFn0(p0 uint64) int64 {
    if (0 < 0) {
        v6 := []int{-9223372036854775808, -10}
        var v7 uint64 = p0
        _ = v6
        _ = v7
    }
    if false {
        v8 := GenS1{F0: 128, F1: -938, F2: -1}
        _ = v8
    }
    p0 = 4294967295
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 int8 = -128
        _ = fi9
        _ = v10
    }
    if (32767 <= 0) {
        p0 = (^(p0))
        p0 = 2
        var v11 uint16 = 2
        _ = v11
    } else {
        p0 += (p0 | 0)
    }
    _ = p0
    return 228
}
func genFn1(p0 uint64) int {
    var v12 uint32 = 2
    fmt.Println(p0)
    if (!(false && (!((0 <= 127) && false)))) {
        var v13 string = "a"
        _ = v13
    }
    var v14 uint = 891
    var v15 uint32 = (v12 << ((v14) & 63))
    _ = p0
    _ = v12
    _ = v14
    _ = v15
    return -9223372036854775808
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]int)
    var v3 float64 = -2.25
    var v4 uint64 = 154
    var v5 uint16 = 723
    v3 = v3
    switch "foo bar" {
    case "":
        var v16 int64 = -1
        _ = v16
    case "foo bar":
        var v17 int32 = 0
        _ = v17
    }
    var v18 uint = 635
    if false {
        var v19 uint32 = 406
        v20 := []int{}
        _ = v19
        _ = v20
    } else {
        delete(v1, "foo bar")
        v4 = v4
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        var v22 int16 = 0
        v23 := GenS2{F0: 42}
        _ = fi21
        _ = v22
        _ = v23
    }
    v24 := []int{471, genFn1((v4 >> ((v18) & 63)))}
    if (-2147483648 != -1) {
        var v25 int = 127
        v25 = v25
        v4 += v4
        _ = v25
    }
    for fi26 := 0; fi26 < 18; fi26++ {
        for fi27 := 0; fi27 < 19; fi27++ {
            fmt.Println(fi27, v18, fi26)
            var v28 float32 = 1e10
            _ = fi27
            _ = v28
        }
        fmt.Println(v18, v3, v18)
        v18 = v18
        _ = fi26
    }
    if (false && true) {
        var v29 int32 = 0
        if len(v24) >= 2 {
            v24 = v24[1:len(v24)]
        }
        v3 -= (v3 + v3)
        _ = v29
    } else {
        fmt.Println(v18)
        var v30 int32 = -760
        _ = v30
    }
    fmt.Println(v4, v18, v3)
    var v31 string = ("" + "")
    for fi32 := 0; fi32 < 12; fi32++ {
        var v33 int16 = -32768
        v33 = (v33 % (((v33 % (((v33 * v33)) | 1))) | 1))
        _ = fi32
        _ = v33
    }
    fmt.Println(v4)
    switch 1 {
    case -1:
        v34 := GenS1{F0: 133, F1: 2147483647, F2: -678}
        _ = v34
    default:
        var v35 float64 = -2.25
        _ = v35
    }
    for fi36 := 0; fi36 < 10; fi36++ {
        var v37 uint = v18
        _ = fi36
        _ = v37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
