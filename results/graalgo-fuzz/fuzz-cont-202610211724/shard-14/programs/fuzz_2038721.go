package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 uint32; F1 bool; F2 uint }
func genFn0(p0 int32, p1 string) int16 {
    p1 = ((p1 + ((("a" + p1) + "foo bar") + p1)) + p1)
    if (28 <= 55) {
        fmt.Println(p0, p1)
    } else {
        p1 = ("hello" + "\n\t")
        var v4 int32 = p0
        fmt.Println(p1)
        _ = v4
    }
    v5 := GenS1{F0: 10, F1: 711}
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := []int{fi6}
        var v8 uint64 = 4294967295
        _ = len(v7)
        _ = fi6
        _ = v7
        _ = v8
    }
    if ((127 > 238) && true) {
        v9 := []int{-248, 32767, 42, 32767}
        var v10 string = (p1 + (p1 + p1))
        _ = v10
        _ = v9
    } else {
        switch 128 {
        case 2147483647:
            v11 := GenS2{F0: -958, F1: (p0 * p0)}
            _ = v11
        case 42:
            if (-2.25 == 0.1) {
                if (269 < -1) {
                    if (1 > 255) {
                        var v12 int32 = p0
                        _ = v12
                    }
                    p1 = p1
                } else {
                    for fi13 := 0; fi13 < 7; fi13++ {
                        switch -1 {
                        case 4294967295:
                            if (false || false) {
                                switch p1 {
                                case "a":
                                    var v14 uint64 = 4294967295
                                    _ = v14
                                default:
                                    var v15 int64 = -1
                                    _ = v15
                                }
                            } else {
                                var v16 int8 = 127
                                if (!(42 > 42)) {
                                    fmt.Println(v16, p0, v16)
                                    p1 = p1
                                }
                                _ = v16
                            }
                        }
                        _ = fi13
                    }
                }
                for fi17 := 0; fi17 < 17; fi17++ {
                    v18 := []int{-2147483648}
                    _ = fi17
                    _ = v18
                }
                fmt.Println(p0, p1)
            }
        case -9223372036854775808:
            p0 -= (p0 * p0)
        default:
            var v19 bool = (!(-406 < 9223372036854775807))
            _ = v19
        }
        if true {
            v20 := []int{-9223372036854775808, 9223372036854775807, -945}
            _ = v20
        } else {
            p0 += p0
        }
        if (!(42 != 2)) {
            var v21 float32 = 1.5
            _ = v21
        }
    }
    _ = p0
    _ = p1
    _ = v5
    return -186
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int8 = 23
    v3 := []int{9223372036854775807, 0, 194, 32768}
    if len(v3) > 0 {
        v3[0] = 163
    }
    var v22 uint8 = 127
    if (false && false) {
        var v23 int = 400
        if false {
            var v24 float32 = 0.1
            _ = len(v3)
            _ = v24
        }
        _ = v23
    }
    if false {
        for fi25 := 0; fi25 < 2; fi25++ {
            fmt.Println(v2)
            _ = fi25
        }
    }
    v2 -= 0
    var v26 float64 = 0.1
    var v27 int = 517
    if len(v3) > 0 {
        v3[0] = (v27 >> ((659) & 63))
    }
    v2 = v2
    var v28 uint = 65535
    for fi29 := 0; fi29 < 12; fi29++ {
        var v30 uint = 18446744073709551615
        _ = fi29
        _ = v30
    }
    if ((^(v22)) == (v22 % ((127) | 1))) {
        v31 := GenS2{F0: v1, F1: -2147483648}
        _ = v31
    }
    _ = len(v3)
    v27 -= v27
    for fi32 := 0; fi32 < 4; fi32++ {
        v33 := GenS3{F0: 0, F1: false, F2: v28}
        _ = fi32
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
