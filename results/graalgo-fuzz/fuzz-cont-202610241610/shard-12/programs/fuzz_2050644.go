package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8; F1 uint }
type GenS5 struct { F0 string }
type GenS6 struct { F0 bool; F1 uint64; F2 uint64 }
func genFn0(p0 uint, p1 int32) bool {
    v6 := []int{926}
    var v7 float64 = (float64(3.14159) / 1e10)
    if (-1 != -888) {
        v8 := GenS1{F0: 42}
        for fi9 := 0; fi9 < 17; fi9++ {
            switch "\n\t" {
            case "foo bar":
                for fi10 := 0; fi10 < 15; fi10++ {
                    p1 += p1
                    _ = fi10
                }
            default:
                p1 = 522
            }
            if (!("hello" <= ("foo bar" + ("" + "foo bar")))) {
                v11 := GenS2{F0: (1 <= 0)}
                _ = v11
            } else {
                if false {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    switch (fi9 >> (((p0 << ((p0) & 63))) & 63)) {
                    case 128:
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    case 0:
                        v7 = (v7 * v7)
                    default:
                        if false {
                            for fi12 := 0; fi12 < 20; fi12++ {
                                switch "" {
                                case "hello":
                                    p0 = 0
                                default:
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    var v13 uint16 = 530
                                    _ = v13
                                }
                                var v14 uint16 = 42
                                var v15 uint64 = 18446744073709551615
                                _ = fi12
                                _ = v14
                                _ = v15
                            }
                            var v16 float32 = 1e10
                            v17 := GenS3{F0: p0}
                            _ = v16
                            _ = v17
                        } else {
                            var v18 int8 = -128
                            var v19 uint64 = 0
                            for fi20 := 0; fi20 < 8; fi20++ {
                                var v21 uint16 = 251
                                fmt.Println(v21)
                                var v22 float64 = (v7 - float64(-2.25))
                                switch (fi20 & 128) {
                                case 0:
                                    fmt.Println(v18, fi9, v18)
                                default:
                                    switch "\n\t" {
                                    case "\n\t":
                                        v21 = 2
                                    case "a":
                                        p1 += p1
                                    }
                                }
                                p1 = p1
                                v23 := GenS3{F0: (p0 / (((p0 << ((p0) & 63))) | 1))}
                                var v24 int = fi9
                                _ = fi20
                                _ = v21
                                _ = v22
                                _ = v23
                                _ = v24
                            }
                            _ = v18
                            _ = v19
                        }
                    }
                } else {
                    for fi25 := 0; fi25 < 17; fi25++ {
                        v7 = (0.0 * ((v7 / (1e10 * v7)) * v7))
                        _ = fi25
                    }
                }
            }
            _ = fi9
        }
        var v26 int64 = 48
        _ = v26
        _ = v8
    }
    switch "hello" {
    case "a":
        var v27 int64 = 42
        _ = v27
    default:
        var v28 bool = ("" != ("a" + ("\n\t" + "\n\t")))
        _ = v28
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return ((^(p0)) < 355)
}
func genFn1() int32 {
    for fi29 := 0; fi29 < 16; fi29++ {
        for fi30 := 0; fi30 < 17; fi30++ {
            var v31 int8 = 42
            v31 -= int8(fi30)
            _ = fi30
            _ = v31
        }
        if true {
            for fi32 := 0; fi32 < 9; fi32++ {
                var v33 uint32 = 42
                _ = fi32
                _ = v33
            }
        }
        _ = fi29
    }
    v34 := []int{32768, -2147483648, 32767, 9223372036854775807}
    for fi35 := 0; fi35 < 15; fi35++ {
        fmt.Println(fi35)
        var v36 bool = true
        switch (fi35 >> ((1) & 63)) {
        case 1:
            var v37 uint = 42
            _ = v37
        case -657:
            var v38 int = fi35
            _ = v38
        }
        _ = fi35
        _ = v36
    }
    _ = v34
    return -1
}

func main() {
    var v1 int32 = 354
    v2 := make(map[uint32]string)
    v3 := make(map[string]uint32)
    v4 := []string{}
    var v5 int = 2147483647
    if true {
        v39 := GenS2{F0: (1000 > 1000)}
        if (1 == 2) {
            switch 0 {
            case 32768:
                _ = len(v4)
            }
        }
        _ = v39
    } else {
        fmt.Println(v5, v5)
    }
    for fi40 := 0; fi40 < 3; fi40++ {
        var v41 int64 = 0
        var v42 int8 = 127
        v41 = (v41 + int64(v1))
        _ = fi40
        _ = v41
        _ = v42
    }
    v43 := GenS4{F0: 1, F1: 2}
    for fi44 := 0; fi44 < 14; fi44++ {
        var v45 int32 = v1
        v1 = (-(v45))
        v46 := GenS5{F0: "hello"}
        _ = fi44
        _ = v45
        _ = v46
    }
    var v47 uint64 = 0
    for fi48 := 0; fi48 < 15; fi48++ {
        if (!(951 < 65535)) {
            if true {
                fmt.Println(v1)
            }
            var v49 float32 = 0.1
            var v50 uint64 = v47
            _ = v49
            _ = v50
        } else {
            v1 = genFn1()
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        _ = fi48
    }
    if true {
        v51 := GenS6{F0: false, F1: v47, F2: v47}
        if len(v4) > 0 {
            v4[0] = "\n\t"
        }
        _ = v51
    }
    v5 += v5
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v47
    _ = v5
}
