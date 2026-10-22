package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8; F1 uint16; F2 string }
func genFn0(p0 uint, p1 uint64) uint32 {
    fmt.Println(p0)
    switch "\n\t" {
    case "foo bar":
        if (265 == 2147483647) {
            p0 = (p0 * (p0 - p0))
        } else {
            p1 += p1
            p1 = (^(p1))
        }
    case "a":
        var v4 int8 = 42
        _ = v4
    }
    if false {
        var v5 float64 = float64(0.1)
        if true {
            var v6 uint64 = 1
            if false {
                if true {
                    var v7 uint16 = 42
                    switch "\n\t" {
                    case "":
                        fmt.Println(v6, v7, v5)
                    case "a":
                        var v8 float64 = v5
                        _ = v8
                    }
                    v9 := GenS1{F0: "hello"}
                    _ = v7
                    _ = v9
                }
            } else {
                for fi10 := 0; fi10 < 7; fi10++ {
                    p1 = p1
                    v11 := GenS2{F0: v6, F1: 255}
                    v6 = p1
                    _ = fi10
                    _ = v11
                }
                v12 := []int{845, 128, -9223372036854775808, -1}
                var v13 int64 = 1
                _ = v12
                _ = v13
            }
            _ = v6
        } else {
            var v14 uint = p0
            fmt.Println(v5, p1, p0)
            p1 -= p1
            _ = v14
        }
        var v15 bool = true
        _ = v15
        _ = v5
    } else {
        v16 := []int{-1, 2147483647}
        var v17 float32 = (1.5 * (1e10 + 0.0))
        _ = v16
        _ = v17
    }
    p1 = p1
    p1 -= 368
    _ = p0
    _ = p1
    return 2147483647
}
func genFn1() bool {
    v18 := GenS3{F0: 128}
    var v19 uint32 = 2147483647
    switch 128 {
    case 9223372036854775807:
        v19 = (v19 / ((v19) | 1))
    case 42:
        v19 -= 756
    case 127:
        v19 = v19
    }
    _ = v18
    _ = v19
    return false
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]string)
    var v3 uint8 = 42
    fmt.Println(v3)
    for fi20 := 0; fi20 < 19; fi20++ {
        v21 := GenS1{F0: ""}
        _ = fi20
        _ = v21
    }
    v3 -= v3
    var v22 int64 = -1
    var v23 uint32 = 206
    v24 := []int{}
    for fi25 := 0; fi25 < 6; fi25++ {
        for fi26 := 0; fi26 < 5; fi26++ {
            var v27 int = fi25
            _ = fi26
            _ = v27
        }
        var v28 uint32 = 42
        _ = fi25
        _ = v28
    }
    if true {
        if false {
            switch "a" {
            case "\n\t":
                v24 = append(v24, -132)
            default:
                v29 := GenS4{F0: (v3 * (^(v3))), F1: 42, F2: "\n\t"}
                _ = v29
            }
            var v30 bool = false
            _ = v30
        }
        _ = len(v2)
    } else {
        v2[("hello" + ("" + "\n\t"))] = "\n\t"
        v23 += (v23 % ((genFn0(872, 81)) | 1))
    }
    var v31 uint8 = v3
    v23 = 42
    if len(v24) > 0 {
        v24[0] = 9223372036854775807
    }
    for fi32 := 0; fi32 < 16; fi32++ {
        fmt.Println(v3, v23)
        _ = fi32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v31
}
