package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 int64; F2 int64 }
type GenS3 struct { F0 int8; F1 uint32; F2 bool }
type GenS4 struct { F0 uint8 }
func genFn0(p0 int32, p1 string, p2 bool) int64 {
    var v5 uint32 = 108
    v6 := []int{1, 9223372036854775807, 2147483647, 579}
    var v7 int8 = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 262
}

func main() {
    var v1 int16 = 1
    var v2 uint8 = 1
    var v3 int8 = -128
    var v4 uint16 = 2
    var v8 float64 = 1.5
    for fi9 := 0; fi9 < 15; fi9++ {
        v10 := GenS1{F0: 65535}
        var v11 int32 = 0
        v12 := GenS2{F0: 42, F1: genFn0(v11, "a", true), F2: 9223372036854775807}
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    v13 := GenS1{F0: 42}
    if (genFn0(652, ((("a" + "") + "foo bar") + ""), (2147483647 <= 4294967295)) <= 0) {
        fmt.Println(v2)
        v14 := []int{123, -387}
        for fi15 := 0; fi15 < 12; fi15++ {
            v14 = append(v14, fi15)
            var v16 string = ("" + ((("\n\t" + (((("a" + "foo bar") + ("hello" + "")) + "a") + "a")) + "a") + "foo bar"))
            if (!(2 == 65535)) {
                v16 = ("a" + ("a" + v16))
            }
            _ = fi15
            _ = v16
        }
        _ = v14
    } else {
        if (((0.1 * 3.14159) * ((1e10 + 1.5) - (1.5 * (0.1 - float32((v8 * ((1e10 - v8) / 3.14159))))))) <= (float32(-2.25) - 0.0)) {
            v17 := []int{-2147483648, -564, -1}
            var v18 float64 = float64((-2.25 * -2.25))
            if len(v17) > 0 {
                v17[0] = 4294967295
            }
            _ = v17
            _ = v18
        }
    }
    fmt.Println(v3, v3)
    var v19 uint32 = 52
    var v20 int16 = (v1 - -1)
    v21 := []int{817}
    var v22 uint = 2
    v21 = append(v21, 160)
    fmt.Println(v3, v4, v8)
    v23 := GenS3{F0: 100, F1: uint32(v4), F2: (true || (!(false || false)))}
    var v24 string = "hello"
    if (false || false) {
        if true {
            var v25 uint = (v22 + 65535)
            _ = v25
        }
        v4 = (v4 | (^(v4)))
    }
    var v26 uint16 = v4
    var v27 float32 = (1e10 * 0.0)
    if false {
        for fi28 := 0; fi28 < 14; fi28++ {
            v29 := GenS4{F0: (v2 ^ 215)}
            _ = fi28
            _ = v29
        }
    } else {
        var v30 int = 64
        _ = v30
    }
    fmt.Println(v27)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v8
}
