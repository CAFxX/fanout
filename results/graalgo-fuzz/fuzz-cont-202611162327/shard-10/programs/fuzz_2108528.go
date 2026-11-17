package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 int32; F1 bool }

func main() {
    var v1 string = "foo bar"
    var v2 int64 = -803
    var v3 uint16 = 10
    v1 = "foo bar"
    for fi4 := 0; fi4 < 15; fi4++ {
        v5 := []int{-9223372036854775808, 789, 1}
        for fi6 := 0; fi6 < 20; fi6++ {
            fmt.Println(fi6)
            _ = fi6
        }
        _ = fi4
        _ = v5
    }
    v7 := GenS1{F0: 296, F1: 0}
    fmt.Println(v3, v3)
    fmt.Println(v3, v1, v2)
    if false {
        var v8 float32 = (float32(0.1) - 1.5)
        switch v1 {
        case "":
            var v9 float32 = 0.0
            _ = v9
        }
        var v10 int16 = 32767
        _ = v10
        _ = v8
    } else {
        v11 := GenS1{F0: 0, F1: 42}
        _ = v11
    }
    if false {
        v2 = v2
    } else {
        var v12 uint64 = 0
        switch 42 {
        case 1:
            var v13 float32 = (3.14159 + 3.14159)
            _ = v13
        case 9223372036854775807:
            var v14 int8 = 42
            _ = v14
        case 32767:
            fmt.Println(v3, v1)
        default:
            var v15 float64 = (-2.25 - (0.1 + ((1e10 / float64(-2.25)) - 0.0)))
            _ = v15
        }
        _ = v12
    }
    v1 = "hello"
    for fi16 := 0; fi16 < 9; fi16++ {
        var v17 int8 = -1
        v2 += v2
        v2 -= v2
        _ = fi16
        _ = v17
    }
    if false {
        v18 := []int{42, -985, 478, -1}
        _ = v18
    } else {
        var v19 string = (v1 + ((v1 + v1) + "hello"))
        var v20 uint64 = 4294967295
        _ = v19
        _ = v20
    }
    v3 += (v3 + 2)
    fmt.Println(v3)
    v21 := []int{320, -1, 4294967295}
    v2 = v2
    v22 := GenS2{F0: -580, F1: false}
    for fi23 := 0; fi23 < 6; fi23++ {
        v24 := GenS1{F0: 2147483648, F1: -2147483648}
        var v25 string = v1
        v3 = (v3 - 2)
        _ = fi23
        _ = v24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v7
}
