package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]uint32)
    var v3 int32 = 1
    var v4 float64 = float64(((1e10 + 1.5) + 1e10))
    v5 := make(map[string]string)
    v6 := make(map[string]uint32)
    if true {
        fmt.Println(v3, v3)
        var v7 float64 = v4
        _ = v7
    } else {
        v5[((("hello" + "hello") + ("a" + (("foo bar" + "") + ("\n\t" + "hello")))) + "hello")] = "foo bar"
    }
    switch ((("hello" + "hello") + ("" + ("hello" + ""))) + "a") {
    case "\n\t":
        fmt.Println(v3)
    }
    switch -502 {
    case 128:
        for fi8 := 0; fi8 < 10; fi8++ {
            v4 += v4
            _ = fi8
        }
    case 9223372036854775807:
        for fi9 := 0; fi9 < 10; fi9++ {
            switch ("\n\t" + (("hello" + "") + "")) {
            case "foo bar":
                v10 := GenS1{F0: 1, F1: 209, F2: (fi9 - fi9)}
                _ = v10
            default:
                fmt.Println(v3, v3)
            }
            if false {
                var v11 int64 = -652
                var v12 uint16 = 0
                _ = v11
                _ = v12
            } else {
                _ = len(v2)
                fmt.Println(fi9, v3, v3)
            }
            v13 := []int{}
            _ = fi9
            _ = v13
        }
    case 41:
        var v14 int = 9223372036854775807
        _ = v14
    default:
        v15 := GenS1{F0: 128, F1: 1000, F2: -810}
        _ = v15
    }
    var v16 int32 = v3
    switch -342 {
    case 128:
        var v17 int8 = 24
        _ = v17
    default:
        var v18 uint = 1
        _ = v18
    }
    var v19 float64 = v4
    var v20 uint16 = 904
    var v21 int64 = -440
    for fi22 := 0; fi22 < 14; fi22++ {
        fmt.Println(fi22)
        for fi23 := 0; fi23 < 19; fi23++ {
            v24 := []int{-105, fi23, -2147483648, (fi22 << ((2) & 63))}
            _ = fi23
            _ = v24
        }
        fmt.Println(fi22, v16, v20)
        _ = fi22
    }
    v25 := GenS1{F0: 5, F1: 0, F2: 1}
    fmt.Println(v21, v19)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
