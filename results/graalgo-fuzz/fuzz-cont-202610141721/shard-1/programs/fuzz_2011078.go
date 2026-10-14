package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 string; F2 uint64 }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 int32 }
func genFn0(p0 int64) int {
    var v4 bool = true
    p0 += p0
    switch "\n\t" {
    case "hello":
        var v5 float64 = -2.25
        _ = v5
    }
    fmt.Println(p0, v4)
    v6 := []int{}
    _ = p0
    _ = v4
    _ = v6
    return 32768
}

func main() {
    v1 := make(map[int]string)
    var v2 int8 = 1
    v3 := []string{"foo bar", "foo bar", "hello"}
    v2 += v2
    var v7 uint32 = 0
    v7 -= v7
    if true {
        for fi8 := 0; fi8 < 17; fi8++ {
            v7 -= (v7 % (((v7 + (v7 & 506))) | 1))
            _ = fi8
        }
        v7 = 4294967295
    }
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    fmt.Println(v7, v7)
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 float64 = 0.1
        for fi11 := 0; fi11 < 18; fi11++ {
            fmt.Println(v7, v2, v7)
            _ = fi11
        }
        _ = fi9
        _ = v10
    }
    if true {
        v7 = v7
        var v12 int64 = 0
        _ = v12
    }
    _ = len(v3)
    for fi13 := 0; fi13 < 10; fi13++ {
        v1[genFn0(971)] = "\n\t"
        _ = fi13
    }
    if (true || false) {
        switch "hello" {
        case "\n\t":
            v14 := GenS1{F0: false}
            _ = v14
        case "hello":
            var v15 uint = 754
            _ = v15
        case "":
            fmt.Println(v7)
        default:
            var v16 float64 = ((0.0 + 0.0) * 3.14159)
            _ = v16
        }
        for fi17 := 0; fi17 < 12; fi17++ {
            v18 := GenS2{F0: 9223372036854775807, F1: "hello", F2: 1}
            if (!(false && false)) {
                switch "foo bar" {
                case "a":
                    v2 = v2
                default:
                    var v19 float32 = 3.14159
                    _ = v19
                }
                var v20 string = ("" + "")
                v21 := GenS3{F0: v7, F1: "a"}
                _ = v20
                _ = v21
            } else {
                v7 = v7
            }
            _ = fi17
            _ = v18
        }
    }
    fmt.Println(v7)
    _ = len(v3)
    for fi22 := 0; fi22 < 8; fi22++ {
        for fi23 := 0; fi23 < 9; fi23++ {
            if false {
                var v24 float64 = 1e10
                if (!true) {
                    v2 += -128
                    v2 += v2
                } else {
                    if (v2 > v2) {
                        var v25 uint64 = 2
                        v3 = append(v3, "")
                        var v26 int = fi23
                        _ = v25
                        _ = v26
                    } else {
                        var v27 uint8 = 40
                        _ = len(v1)
                        _ = v27
                    }
                    var v28 int16 = 0
                    _ = v28
                }
                v24 = v24
                _ = v24
            } else {
                if (false || false) {
                    v2 = (^(v2))
                    var v29 float64 = 0.1
                    _ = v29
                }
                var v30 int = (fi22 % (((fi23 & genFn0(-1))) | 1))
                _ = v30
            }
            v31 := GenS4{F0: -2147483648}
            _ = fi23
            _ = v31
        }
        v7 -= (v7 / ((v7) | 1))
        _ = fi22
    }
    var v32 uint8 = 255
    var v33 int = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v7
}
