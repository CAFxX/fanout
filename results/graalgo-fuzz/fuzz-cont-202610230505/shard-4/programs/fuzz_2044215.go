package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 int; F2 uint32 }
func genFn0(p0 uint, p1 int) bool {
    v7 := []int{}
    for fi8 := 0; fi8 < 15; fi8++ {
        for fi9 := 0; fi9 < 20; fi9++ {
            switch p1 {
            case 127:
                v10 := GenS1{F0: false, F1: (fi8 << (((p0 % ((29) | 1))) & 63))}
                _ = v10
            default:
                var v11 string = "hello"
                p1 -= p1
                _ = v11
            }
            _ = fi9
        }
        switch (("foo bar" + "\n\t") + ("" + "hello")) {
        case "foo bar":
            if false {
                for fi12 := 0; fi12 < 1; fi12++ {
                    switch fi8 {
                    case 32767:
                        if true {
                            if true {
                                v13 := GenS1{F0: ((p0 ^ p0) >= 55), F1: 32768}
                                v7 = append(v7, fi12)
                                _ = v13
                            }
                            if len(v7) > 0 {
                                v7[0] = (fi8 % (((fi12 % ((32768) | 1))) | 1))
                            }
                            var v14 uint64 = 18446744073709551615
                            _ = v14
                        } else {
                            _ = len(v7)
                            var v15 string = ("" + "foo bar")
                            _ = v15
                        }
                    }
                    var v16 int = fi8
                    p0 = p0
                    _ = fi12
                    _ = v16
                }
                if len(v7) > 0 {
                    v7[0] = (^(fi8))
                }
                switch ("foo bar" + "\n\t") {
                case "a":
                    var v17 uint = p0
                    _ = v17
                default:
                    var v18 uint = (p0 & p0)
                    _ = v18
                }
            } else {
                for fi19 := 0; fi19 < 9; fi19++ {
                    fmt.Println(fi8, fi19)
                    v20 := GenS2{F0: -750}
                    _ = fi19
                    _ = v20
                }
            }
        case "\n\t":
            fmt.Println(fi8, p0)
        case "a":
            for fi21 := 0; fi21 < 14; fi21++ {
                p1 -= (^(fi8))
                _ = fi21
            }
        default:
            var v22 int32 = 1
            _ = v22
        }
        _ = fi8
    }
    if ((!(-9223372036854775808 <= 742)) || (177 > 5)) {
        var v23 int = (p1 % ((p1) | 1))
        _ = v23
    }
    _ = p0
    _ = p1
    _ = v7
    return (0 <= 72)
}
func genFn1(p0 int, p1 uint8, p2 uint32) int8 {
    if (true || false) {
        v24 := []int{(p0 - p0)}
        p1 += 87
        _ = v24
    }
    v25 := GenS1{F0: (false || false), F1: (p0 << ((209) & 63))}
    p1 = p1
    var v26 int = (-(p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    _ = v26
    return 100
}

func main() {
    var v1 int8 = 127
    var v2 int32 = 2147483647
    var v3 uint16 = 65535
    v4 := make(map[uint32]string)
    var v5 int = -1
    v6 := make(map[string]uint32)
    var v27 int8 = genFn1(int(v3), 0, 2)
    v2 = (v2 & v2)
    v27 += (-(v1))
    v28 := GenS3{F0: 9223372036854775807, F1: -159, F2: 557}
    var v29 uint = 961
    var v30 uint32 = 1
    v2 = (v2 & (v2 * v2))
    v31 := []int{-415, v5}
    v2 -= -615
    for fi32 := 0; fi32 < 16; fi32++ {
        var v33 float64 = (((float64((float32(3.14159) + float32((0.1 * -2.25)))) + (3.14159 + 0.0)) * 1e10) / 3.14159)
        _ = fi32
        _ = v33
    }
    switch "hello" {
    case "foo bar":
        v3 = 472
    default:
        v30 = 338
    }
    v34 := GenS3{F0: 1, F1: v5, F2: v30}
    switch "\n\t" {
    case "\n\t":
        for fi35 := 0; fi35 < 7; fi35++ {
            fmt.Println(v29, v5)
            _ = fi35
        }
    case "foo bar":
        var v36 bool = true
        _ = v36
    case "a":
        v37 := GenS3{F0: -746, F1: (v5 ^ (v5 / ((v5) | 1))), F2: (^(v30))}
        _ = v37
    }
    var v38 uint8 = 127
    v5 = v5
    if (false || ((false && false) && (18446744073709551615 < v29))) {
        if (!(uint(v3) >= v29)) {
            v29 -= v29
            v29 += v29
            fmt.Println(v1)
        }
        v31 = append(v31, v5)
    } else {
        var v39 uint16 = 593
        delete(v4, v30)
        switch (("\n\t" + "") + "") {
        case "a":
            v40 := GenS2{F0: 32767}
            _ = v40
        default:
            var v41 int8 = v1
            _ = v41
        }
        _ = v39
    }
    fmt.Println(v5, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
