package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int16 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64; F1 bool }
func genFn0() bool {
    switch -202 {
    case 435:
        v4 := GenS1{F0: -897, F1: -128, F2: 1}
        _ = v4
    }
    for fi5 := 0; fi5 < 12; fi5++ {
        var v6 float32 = (1.5 * 1e10)
        var v7 uint8 = 2
        v8 := GenS1{F0: 42, F1: 9, F2: -1}
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    if true {
        for fi9 := 0; fi9 < 8; fi9++ {
            v10 := GenS1{F0: 9223372036854775807, F1: -27, F2: 1000}
            var v11 int8 = -1
            _ = fi9
            _ = v10
            _ = v11
        }
        v12 := []int{-1, 2147483647}
        _ = v12
    } else {
        v13 := GenS2{F0: ("\n\t" + "hello"), F1: 42}
        _ = v13
    }
    v14 := []int{127, -2147483648}
    switch 359 {
    case 128:
        if false {
            var v15 uint8 = 66
            v15 += v15
            _ = v15
        }
    }
    _ = v14
    return true
}

func main() {
    v1 := []int{-1, 4294967295, 971}
    var v2 uint16 = 42
    v3 := make(map[uint32]string)
    var v16 bool = ((true && (1 == 0)) || (("a" != "") && false))
    switch -210 {
    case 9223372036854775807:
        fmt.Println(v2, v2)
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        v18 := GenS1{F0: 9223372036854775807, F1: 100, F2: -1}
        for fi19 := 0; fi19 < 18; fi19++ {
            if false {
                delete(v3, 2)
                var v20 uint16 = (v2 ^ (v2 / ((v2) | 1)))
                _ = v20
            }
            if len(v1) > 0 {
                v1[0] = (^(fi19))
            }
            _ = fi19
        }
        _ = fi17
        _ = v18
    }
    for fi21 := 0; fi21 < 19; fi21++ {
        v22 := GenS3{F0: 746}
        delete(v3, 1)
        var v23 float32 = ((float32(float64(((1.5 * 0.1) - (((1e10 + (-2.25 / (0.1 + (3.14159 / (((float32(1.5) + -2.25) + 1.5) + (1.5 + float32((1.5 * 1e10)))))))) / 1.5) * 0.1)))) / -2.25) + -2.25)
        _ = fi21
        _ = v22
        _ = v23
    }
    v24 := GenS4{F0: 42, F1: v16}
    if true {
        if (!v16) {
            fmt.Println(v2)
            var v25 int64 = 280
            if true {
                v16 = ("hello" == "foo bar")
                fmt.Println(v16, v25, v25)
                fmt.Println(v2, v25, v2)
            }
            _ = v25
        } else {
            v2 = v2
        }
        for fi26 := 0; fi26 < 20; fi26++ {
            if v16 {
                v16 = true
            } else {
                v16 = (fi26 != fi26)
                var v27 string = ("foo bar" + "foo bar")
                v28 := GenS4{F0: 869, F1: (!true)}
                _ = v27
                _ = v28
            }
            _ = fi26
        }
    }
    var v29 int = 2147483647
    for fi30 := 0; fi30 < 7; fi30++ {
        var v31 uint8 = 0
        _ = fi30
        _ = v31
    }
    var v32 int64 = 1
    var v33 int = v29
    fmt.Println(v32, v32, v29)
    v16 = (!false)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v16
    _ = v2
    _ = v24
    _ = v29
    _ = v3
    _ = v32
    _ = v33
}
