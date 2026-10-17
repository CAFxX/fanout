package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 uint; F1 int32; F2 int }
func genFn0(p0 bool) uint16 {
    var v5 float32 = 0.1
    var v6 uint32 = 2147483647
    var v7 int32 = -946
    if (42 <= -9223372036854775808) {
        fmt.Println(v6)
        switch "a" {
        case "\n\t":
            for fi8 := 0; fi8 < 5; fi8++ {
                p0 = (v6 != v6)
                v9 := GenS1{F0: (v7 * v7), F1: ("" + "hello")}
                for fi10 := 0; fi10 < 15; fi10++ {
                    p0 = true
                    v11 := []int{32768, (fi8 * fi10), 9223372036854775807, 0}
                    if p0 {
                        v12 := GenS2{F0: 2147483647, F1: v7}
                        switch fi10 {
                        case -9223372036854775808:
                            var v13 float64 = -2.25
                            _ = v13
                        case 9223372036854775807:
                            v6 = (v6 + v6)
                        }
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                        _ = v12
                    }
                    _ = fi10
                    _ = v11
                }
                _ = fi8
                _ = v9
            }
        case "foo bar":
            v14 := []int{4294967295, -9, 4294967295, 9223372036854775807}
            _ = v14
        case "hello":
            v15 := []int{}
            _ = v15
        default:
            v16 := []int{4294967295, 4294967295, 32768}
            _ = v16
        }
    } else {
        if p0 {
            v17 := []int{-774, 128, 1, -9223372036854775808}
            _ = v17
        } else {
            if ((v7 - (-(v7))) == (-(v7))) {
                v18 := GenS2{F0: (v6 ^ v6), F1: v7}
                for fi19 := 0; fi19 < 9; fi19++ {
                    switch "foo bar" {
                    case "\n\t":
                        v5 = -2.25
                    case "":
                        fmt.Println(v5, v5, v6)
                    }
                    var v20 uint64 = 1
                    _ = fi19
                    _ = v20
                }
                v7 -= (v7 * v7)
                _ = v18
            }
            var v21 int = 127
            var v22 int = (^(v21))
            _ = v21
            _ = v22
        }
        fmt.Println(v5, p0, v5)
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 uint64) uint32 {
    var v23 int = -2147483648
    v24 := GenS3{F0: 0, F1: 1, F2: 4294967295}
    for fi25 := 0; fi25 < 11; fi25++ {
        var v26 int64 = -9223372036854775808
        _ = fi25
        _ = v26
    }
    _ = p0
    _ = v23
    _ = v24
    return 2
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[int]string)
    var v3 uint8 = 255
    var v4 uint8 = (v3 | v3)
    v27 := []int{-1}
    fmt.Println(v4, v1, v4)
    var v28 bool = false
    switch "" {
    case "foo bar":
        fmt.Println(v4, v1, v4)
    default:
        switch "a" {
        case "a":
            v29 := GenS2{F0: 2147483648, F1: 42}
            _ = v29
        case "foo bar":
            for fi30 := 0; fi30 < 6; fi30++ {
                v4 = v4
                if ((fi30 & (fi30 + (fi30 ^ fi30))) >= -1) {
                    fmt.Println(v4, v28)
                } else {
                    var v31 uint8 = (v3 / ((v3) | 1))
                    _ = v31
                }
                for fi32 := 0; fi32 < 17; fi32++ {
                    v3 += v3
                    v33 := GenS3{F0: 536, F1: -1, F2: int(v1)}
                    _ = fi32
                    _ = v33
                }
                _ = fi30
            }
        }
    }
    for fi34 := 0; fi34 < 16; fi34++ {
        delete(v2, (-(fi34)))
        var v35 uint16 = 614
        for fi36 := 0; fi36 < 3; fi36++ {
            var v37 uint32 = 65535
            var v38 float32 = 0.0
            v2[fi36] = "foo bar"
            _ = fi36
            _ = v37
            _ = v38
        }
        _ = fi34
        _ = v35
    }
    _ = len(v27)
    fmt.Println(v4, v4)
    if v28 {
        for fi39 := 0; fi39 < 7; fi39++ {
            var v40 uint16 = 1
            var v41 float32 = -2.25
            if (!true) {
                if true {
                    _ = len(v2)
                    switch "hello" {
                    case "a":
                        v1 -= v1
                    case "hello":
                        for fi42 := 0; fi42 < 4; fi42++ {
                            var v43 bool = ((-187 <= 1) || ((fi42 / (((fi39 | (fi39 & 660))) | 1)) < (^(fi39))))
                            _ = fi42
                            _ = v43
                        }
                    }
                    v3 = v4
                }
                if len(v27) >= 2 {
                    v27 = v27[1:len(v27)]
                }
                v41 -= (v41 * v41)
            }
            _ = fi39
            _ = v40
            _ = v41
        }
    }
    v4 += 21
    var v44 float64 = 3.14159
    v45 := v2[498]
    fmt.Println(v28, v4, v45)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v44
    _ = v45
}
