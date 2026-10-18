package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32 }
func genFn0(p0 uint16, p1 int16) int32 {
    for fi5 := 0; fi5 < 11; fi5++ {
        if false {
            for fi6 := 0; fi6 < 13; fi6++ {
                if false {
                    fmt.Println(p1, fi6, fi6)
                } else {
                    if true {
                        v7 := GenS1{F0: fi6}
                        _ = v7
                    } else {
                        v8 := []int{}
                        _ = v8
                    }
                    v9 := []int{}
                    if (true || false) {
                        p0 -= 1
                        p1 = p1
                        if len(v9) > 0 {
                            v9[0] = -2147483648
                        }
                    }
                    _ = v9
                }
                fmt.Println(fi5, fi6, p1)
                _ = fi6
            }
        }
        p0 = p0
        var v10 bool = (821 < 2147483647)
        _ = fi5
        _ = v10
    }
    v11 := GenS2{F0: (!(!false)), F1: 221}
    for fi12 := 0; fi12 < 12; fi12++ {
        var v13 bool = (42 <= 51)
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v11
    return 2147483647
}
func genFn1(p0 int8, p1 uint8) bool {
    if (true || true) {
        v14 := []int{128, 128, -1}
        var v15 int8 = 0
        _ = v14
        _ = v15
    }
    p1 += p1
    _ = p0
    _ = p1
    return false
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = -595
    v3 := make(map[string]bool)
    v4 := []uint32{2}
    v2 = v2
    v2 -= 405
    if false {
        fmt.Println(v2)
    }
    var v16 uint64 = 0
    v17 := GenS3{F0: 860}
    v2 = (v2 - int64(v16))
    v2 = v2
    v2 += 0
    switch -9223372036854775808 {
    case 42:
        v2 += (^(v2))
    case 9223372036854775807:
        switch "hello" {
        case "foo bar":
            fmt.Println(v2)
        case "":
            var v18 int32 = genFn0(1, 32767)
            _ = v18
        default:
            if true {
                var v19 int32 = -1
                for fi20 := 0; fi20 < 13; fi20++ {
                    var v21 float64 = (-2.25 / ((0.1 + (3.14159 / (0.1 - 1.5))) - float64(1.5)))
                    fmt.Println(v2, v2, v2)
                    _ = fi20
                    _ = v21
                }
                _ = v19
            }
        }
    }
    v22 := GenS2{F0: (!true), F1: 128}
    for fi23 := 0; fi23 < 13; fi23++ {
        v16 = v16
        v24 := GenS2{F0: false, F1: 2}
        fmt.Println(fi23)
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 7; fi25++ {
        v2 += v2
        if true {
            var v26 int8 = 1
            _ = v26
        } else {
            for fi27 := 0; fi27 < 3; fi27++ {
                if (true || false) {
                    fmt.Println(fi25, fi27, fi25)
                    v28 := GenS4{F0: 996}
                    if (1 < -128) {
                        fmt.Println(fi25, fi25)
                        var v29 uint32 = 61
                        _ = v29
                    } else {
                        v2 -= 1
                        var v30 int = (fi27 * (fi27 ^ (fi27 >> ((798) & 63))))
                        _ = v30
                    }
                    _ = v28
                } else {
                    v2 = (v2 / (((-(v2))) | 1))
                    v2 -= (-(v2))
                }
                fmt.Println(v16, fi27)
                var v31 int32 = 2147483647
                _ = fi27
                _ = v31
            }
            fmt.Println(fi25, fi25, v16)
        }
        v32 := v3["foo bar"]
        _ = fi25
        _ = v32
    }
    v2 = (v2 ^ (v2 | (^(v2))))
    var v33 bool = (65 > 803)
    var v34 uint8 = 67
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v4 = append(v4, 4294967295)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v33
    _ = v34
    _ = v4
}
