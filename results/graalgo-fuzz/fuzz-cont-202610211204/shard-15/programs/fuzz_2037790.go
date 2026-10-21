package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 int8 }
func genFn0() int32 {
    var v6 uint64 = 0
    v6 = v6
    _ = v6
    return -2147483648
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = 78
    var v3 int64 = -9223372036854775808
    var v4 string = ("foo bar" + "hello")
    v5 := make(map[int]uint32)
    for fi7 := 0; fi7 < 17; fi7++ {
        v4 = v4
        fmt.Println(v2, v4)
        _ = fi7
    }
    if true {
        var v8 int = 2147483647
        if true {
            v9 := []int{}
            _ = v9
        }
        _ = v8
    }
    v10 := GenS1{F0: 697, F1: v4}
    fmt.Println(v4, v2)
    var v11 int32 = 2147483647
    if false {
        var v12 int8 = v2
        fmt.Println(v12, v11, v12)
        fmt.Println(v12, v3, v4)
        _ = v12
    } else {
        for fi13 := 0; fi13 < 11; fi13++ {
            var v14 string = (v4 + "a")
            _ = fi13
            _ = v14
        }
        for fi15 := 0; fi15 < 8; fi15++ {
            if false {
                var v16 int8 = v2
                if ((!false) || true) {
                    if false {
                        var v17 uint16 = 2
                        for fi18 := 0; fi18 < 17; fi18++ {
                            var v19 float32 = 0.0
                            _ = fi18
                            _ = v19
                        }
                        _ = v17
                    } else {
                        for fi20 := 0; fi20 < 2; fi20++ {
                            v21 := GenS2{F0: 18446744073709551615}
                            var v22 uint32 = 2147483647
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                        v4 = v4
                        fmt.Println(v2)
                    }
                    for fi23 := 0; fi23 < 20; fi23++ {
                        var v24 int8 = -128
                        _ = len(v1)
                        _ = fi23
                        _ = v24
                    }
                } else {
                    var v25 uint = 2
                    for fi26 := 0; fi26 < 9; fi26++ {
                        v3 = (v3 | (v3 & -89))
                        v11 = (v11 * 42)
                        fmt.Println(v3)
                        v25 -= v25
                        v3 -= v3
                        _ = fi26
                    }
                    for fi27 := 0; fi27 < 8; fi27++ {
                        var v28 uint = (^(v25))
                        _ = fi27
                        _ = v28
                    }
                    _ = v25
                }
                _ = v16
            } else {
                var v29 int8 = int8(fi15)
                _ = v29
            }
            _ = fi15
        }
        v30 := GenS1{F0: -567, F1: "foo bar"}
        _ = v30
    }
    var v31 uint32 = 109
    if false {
        v32 := GenS3{F0: 1, F1: v2}
        if ((!false) || true) {
            var v33 int16 = 545
            _ = v33
        }
        _ = v32
    } else {
        var v34 uint32 = 0
        delete(v5, -839)
        var v35 int8 = v2
        _ = v34
        _ = v35
    }
    switch 689 {
    case -2147483648:
        var v36 uint16 = 2
        _ = v36
    case -9223372036854775808:
        v37 := GenS3{F0: 65535, F1: (v2 / (((^(v2))) | 1))}
        _ = v37
    default:
        var v38 float32 = 1e10
        _ = v38
    }
    v39 := []int{9223372036854775807, -1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v31
    _ = v39
    _ = v4
    _ = v5
}
