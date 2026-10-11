package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint64 }
type GenS2 struct { F0 int8; F1 bool; F2 bool }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int }
func genFn0(p0 bool, p1 uint8) uint64 {
    p1 -= (p1 & (p1 << ((18446744073709551615) & 63)))
    if p0 {
        switch 32768 {
        case 0:
            p0 = false
        case -2147483648:
            v4 := []int{}
            _ = v4
        }
        p1 = p1
    } else {
        p0 = (32767 != 255)
    }
    p1 = p1
    if p0 {
        p0 = ((false && (29 >= 13)) && (-145 >= 0))
        var v5 bool = p0
        _ = v5
    } else {
        switch 4294967295 {
        case 4294967295:
            p0 = (((!p0) || (181 >= 32767)) || true)
        case -960:
            if ("foo bar" == "hello") {
                v6 := GenS1{F0: 2, F1: ((!(32768 == 9223372036854775807)) && p0), F2: 121}
                fmt.Println(p0)
                _ = v6
            } else {
                p0 = (313 == -1)
                if p0 {
                    v7 := GenS2{F0: 29, F1: p0, F2: p0}
                    switch -1 {
                    case 9223372036854775807:
                        v8 := GenS1{F0: 552, F1: p0, F2: 0}
                        _ = v8
                    case 330:
                        var v9 uint32 = 2
                        _ = v9
                    case -9223372036854775808:
                        var v10 int = -9223372036854775808
                        _ = v10
                    }
                    if (128 > -2147483648) {
                        var v11 int8 = -128
                        _ = v11
                    } else {
                        p1 += 102
                        p0 = p0
                    }
                    _ = v7
                } else {
                    fmt.Println(p1, p1)
                    var v12 int8 = -49
                    _ = v12
                }
            }
        default:
            if (!false) {
                var v13 int32 = int32(p1)
                v14 := GenS3{F0: 1}
                _ = v13
                _ = v14
            }
        }
    }
    v15 := GenS4{F0: 100}
    _ = p0
    _ = p1
    _ = v15
    return 2
}

func main() {
    var v1 uint8 = 230
    var v2 float64 = 0.1
    v3 := make(map[string]string)
    v2 += ((v2 - 1e10) / (v2 - v2))
    for fi16 := 0; fi16 < 3; fi16++ {
        if true {
            if (!true) {
                v17 := []int{32767, -2147483648, (fi16 & fi16)}
                fmt.Println(v1, v1)
                var v18 uint = 1
                _ = v17
                _ = v18
            }
            fmt.Println(v2, v2)
            var v19 uint32 = 2147483648
            _ = v19
        }
        v2 = 1.5
        _ = fi16
    }
    v1 = (v1 * 255)
    v2 += (((v2 + ((v2 - v2) + (3.14159 - v2))) - (0.1 - float64(1.5))) * v2)
    for fi20 := 0; fi20 < 6; fi20++ {
        _ = len(v3)
        _ = fi20
    }
    for fi21 := 0; fi21 < 4; fi21++ {
        for fi22 := 0; fi22 < 2; fi22++ {
            if false {
                var v23 uint16 = 42
                _ = v23
            }
            v2 = (float64(1.5) / 0.1)
            if true {
                var v24 uint32 = 356
                _ = v24
            }
            var v25 uint64 = 1
            fmt.Println(v2, v1, fi21)
            _ = fi22
            _ = v25
        }
        v26 := []int{}
        _ = fi21
        _ = v26
    }
    if (313 <= 4294967295) {
        v1 = 128
        if (1 > 521) {
            v1 -= (v1 & (v1 >> ((18446744073709551615) & 63)))
            var v27 uint = 2
            _ = v27
        }
    }
    v28 := []int{42, 32767, 4294967295, 32767}
    v29 := GenS5{F0: 32768}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
}
