package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 int; F2 int16 }
type GenS4 struct { F0 int16; F1 uint8 }
func genFn0(p0 int16, p1 uint) int32 {
    if (!true) {
        p0 = p0
        if (2 > 35) {
            if false {
                var v7 int32 = -2147483648
                _ = v7
            }
            for fi8 := 0; fi8 < 4; fi8++ {
                if ((!false) && false) {
                    var v9 float64 = 0.0
                    fmt.Println(p1)
                    _ = v9
                } else {
                    var v10 int8 = -1
                    _ = v10
                }
                if (int64(fi8) > 9223372036854775807) {
                    if (1.5 >= ((0.0 - (0.1 * 1e10)) + -2.25)) {
                        if (851 <= 42) {
                            switch "\n\t" {
                            case "":
                                p0 = p0
                            case "foo bar":
                                v11 := GenS1{F0: 65535}
                                _ = v11
                            case "a":
                                p0 -= 32767
                            }
                            switch "foo bar" {
                            case "a":
                                fmt.Println(p1, p0, fi8)
                            case "":
                                p1 -= p1
                            case "\n\t":
                                p1 += 1
                            default:
                                var v12 int16 = 861
                                _ = v12
                            }
                            switch fi8 {
                            case 128:
                                v13 := []int{int(p1)}
                                _ = v13
                            case 9223372036854775807:
                                p1 += (p1 << ((p1) & 63))
                            default:
                                v14 := GenS2{F0: 173}
                                _ = v14
                            }
                        }
                        var v15 int16 = (p0 & (p0 & (p0 - (-(p0)))))
                        p1 = 969
                        _ = v15
                    }
                    var v16 uint8 = 100
                    _ = v16
                }
                var v17 uint = p1
                _ = fi8
                _ = v17
            }
        }
    }
    p1 = p1
    var v18 int64 = 0
    var v19 int = 1
    switch 818 {
    case 32768:
        switch "hello" {
        case "foo bar":
            fmt.Println(p1, p0)
        default:
            for fi20 := 0; fi20 < 8; fi20++ {
                var v21 float64 = 1e10
                _ = fi20
                _ = v21
            }
        }
    }
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    return -2147483648
}
func genFn1(p0 uint32) int {
    if (18446744073709551615 == 4294967295) {
        p0 -= p0
        v22 := GenS3{F0: (p0 - 2147483648), F1: 640, F2: 42}
        _ = v22
    } else {
        var v23 uint32 = 879
        var v24 int8 = -128
        fmt.Println(p0, p0, v23)
        _ = v23
        _ = v24
    }
    fmt.Println(p0)
    var v25 uint64 = 0
    _ = p0
    _ = v25
    return -735
}

func main() {
    var v1 int64 = -1
    v2 := make(map[int]int)
    var v3 uint = 990
    var v4 uint8 = 255
    var v5 bool = true
    v6 := []int{2147483647, 128, 168, -9223372036854775808}
    v4 = v4
    var v26 uint16 = 2
    var v27 bool = false
    for fi28 := 0; fi28 < 19; fi28++ {
        var v29 int64 = v1
        _ = fi28
        _ = v29
    }
    _ = len(v6)
    var v30 uint64 = 18446744073709551615
    for fi31 := 0; fi31 < 3; fi31++ {
        fmt.Println(v30, fi31, fi31)
        if v27 {
            for fi32 := 0; fi32 < 19; fi32++ {
                v33 := GenS4{F0: -513, F1: v4}
                _ = fi32
                _ = v33
            }
        } else {
            var v34 uint8 = (v4 * (v4 + (v4 % (((v4 * uint8(v26))) | 1))))
            _ = v34
        }
        _ = fi31
    }
    if len(v6) > 0 {
        v6[0] = genFn1(2)
    }
    v2[32767] = 42
    for fi35 := 0; fi35 < 12; fi35++ {
        if (!(!v27)) {
            var v36 bool = (true || (957 > (v1 >> ((1) & 63))))
            _ = v36
        } else {
            for fi37 := 0; fi37 < 8; fi37++ {
                var v38 float32 = 1e10
                v3 += (v3 - 2)
                _ = fi37
                _ = v38
            }
            var v39 int64 = (v1 | (v1 | (v1 ^ (v1 << ((0) & 63)))))
            v4 = 127
            _ = v39
        }
        v1 -= v1
        _ = fi35
    }
    var v40 int = 356
    if v27 {
        if len(v6) > 0 {
            v6[0] = (v40 | 127)
        }
        for fi41 := 0; fi41 < 20; fi41++ {
            v40 += v40
            var v42 int16 = -53
            _ = fi41
            _ = v42
        }
        v3 -= v3
    } else {
        var v43 uint64 = 275
        _ = len(v6)
        _ = v43
    }
    v5 = false
    var v44 uint64 = uint64(v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v40
    _ = v44
    _ = v5
    _ = v6
}
