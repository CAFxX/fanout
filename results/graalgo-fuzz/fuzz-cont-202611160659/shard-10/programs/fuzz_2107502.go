package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int; F1 int64; F2 int32 }
type GenS5 struct { F0 uint64; F1 bool }
type GenS6 struct { F0 int8; F1 int8 }
type GenS7 struct { F0 uint64; F1 uint8 }
type GenS8 struct { F0 uint64; F1 uint32 }
func genFn0(p0 uint) int {
    v7 := []int{0, 890, 185}
    for fi8 := 0; fi8 < 12; fi8++ {
        _ = len(v7)
        _ = fi8
    }
    v9 := GenS1{F0: -654}
    var v10 float32 = 0.1
    fmt.Println(p0, p0)
    _ = p0
    _ = v10
    _ = v7
    _ = v9
    return -419
}
func genFn1(p0 int, p1 uint8) uint16 {
    var v11 string = ""
    if true {
        var v12 int32 = -983
        var v13 bool = ((0 <= -122) || false)
        switch "hello" {
        case "\n\t":
            var v14 uint16 = 569
            _ = v14
        default:
            v15 := GenS2{F0: 42}
            _ = v15
        }
        _ = v12
        _ = v13
    } else {
        var v16 float64 = (3.14159 * (0.1 - 1e10))
        _ = v16
    }
    v17 := GenS2{F0: -128}
    _ = p0
    _ = p1
    _ = v11
    _ = v17
    return 255
}

func main() {
    var v1 uint8 = 1
    v2 := make(map[uint32]int)
    var v3 int64 = 1
    var v4 float32 = ((-2.25 - 0.1) + -2.25)
    v5 := []uint32{2, 923}
    v6 := make(map[int]int)
    v18 := GenS3{F0: v3}
    v19 := GenS4{F0: -1, F1: 42, F2: 2147483647}
    v1 += 255
    _ = len(v6)
    for fi20 := 0; fi20 < 2; fi20++ {
        for fi21 := 0; fi21 < 16; fi21++ {
            switch genFn0(31) {
            case -26:
                v1 -= (v1 ^ (v1 % (((v1 >> ((864) & 63))) | 1)))
                var v22 uint16 = 42
                var v23 int8 = -128
                _ = v22
                _ = v23
            case 128:
                v24 := GenS5{F0: 850, F1: (!(127 == -280))}
                _ = v24
            }
            v5 = append(v5, 806)
            fmt.Println(v3, v4)
            _ = fi21
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v4 -= -2.25
        _ = fi20
    }
    switch "\n\t" {
    case "\n\t":
        delete(v2, 2)
    case "a":
        v25 := GenS6{F0: 111, F1: -1}
        _ = v25
    default:
        for fi26 := 0; fi26 < 4; fi26++ {
            v27 := GenS5{F0: 874, F1: false}
            if false {
                v28 := GenS1{F0: (fi26 & (fi26 + (fi26 >> ((511) & 63))))}
                for fi29 := 0; fi29 < 6; fi29++ {
                    v30 := GenS5{F0: 1, F1: (!true)}
                    v6[genFn0(65535)] = fi29
                    var v31 int16 = -32768
                    v32 := GenS4{F0: (^(fi29)), F1: v3, F2: 42}
                    _ = fi29
                    _ = v30
                    _ = v31
                    _ = v32
                }
                _ = v28
            }
            _ = fi26
            _ = v27
        }
    }
    if (!true) {
        var v33 float32 = v4
        _ = v33
    }
    v4 -= (v4 + (((0.0 * (3.14159 - (v4 * ((0.0 / 3.14159) - (3.14159 - float32((-2.25 - 1e10))))))) - v4) / v4))
    _ = len(v2)
    v34 := GenS7{F0: 982, F1: 255}
    switch (("a" + "a") + ("a" + "\n\t")) {
    case "\n\t":
        for fi35 := 0; fi35 < 1; fi35++ {
            v36 := GenS8{F0: uint64(v1), F1: 42}
            fmt.Println(v3, v3)
            _ = fi35
            _ = v36
        }
    case "":
        if true {
            _ = len(v5)
        } else {
            if (-107 != 316) {
                var v37 uint16 = 414
                var v38 bool = (1.5 < v4)
                _ = v37
                _ = v38
            } else {
                _ = len(v5)
                if ((42 != 2147483648) && (1 == -32768)) {
                    for fi39 := 0; fi39 < 15; fi39++ {
                        switch ("foo bar" + "") {
                        case "hello":
                            v1 = (v1 >> ((4294967295) & 63))
                        }
                        for fi40 := 0; fi40 < 5; fi40++ {
                            var v41 uint8 = v1
                            if len(v5) > 0 {
                                v5[0] = 1
                            }
                            v4 += 1e10
                            if true {
                                var v42 int8 = -1
                                v42 += (v42 & 42)
                                _ = v42
                            }
                            _ = fi40
                            _ = v41
                        }
                        if (-800 != -327) {
                            v3 += (v3 ^ -1)
                            var v43 int16 = -67
                            _ = v43
                        } else {
                            v1 = (v1 >> ((1) & 63))
                        }
                        _ = fi39
                    }
                }
                v4 = (v4 + v4)
            }
        }
    case "a":
        for fi44 := 0; fi44 < 4; fi44++ {
            v3 -= v3
            _ = fi44
        }
    default:
        v3 = 0
    }
    v4 = (v4 + ((v4 - v4) * 1e10))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
