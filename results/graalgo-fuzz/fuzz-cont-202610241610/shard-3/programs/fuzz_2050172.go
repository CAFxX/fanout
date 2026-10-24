package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 int64; F1 uint; F2 uint64 }
type GenS4 struct { F0 string; F1 string; F2 int64 }
type GenS5 struct { F0 int8 }
func genFn0(p0 uint32, p1 int16) uint16 {
    for fi4 := 0; fi4 < 18; fi4++ {
        switch fi4 {
        case 128:
            if (false && false) {
                for fi5 := 0; fi5 < 18; fi5++ {
                    v6 := GenS1{F0: (((("" + "") + ("foo bar" + "hello")) + "hello") + ("hello" + "foo bar"))}
                    _ = fi5
                    _ = v6
                }
                var v7 uint32 = (p0 ^ p0)
                _ = v7
            }
        case 127:
            if (true && (true && ((p0 >> ((584) & 63)) <= 2147483648))) {
                var v8 int16 = p1
                var v9 uint8 = 175
                _ = v8
                _ = v9
            } else {
                switch "foo bar" {
                case "a":
                    v10 := GenS1{F0: "foo bar"}
                    _ = v10
                case "\n\t":
                    v11 := []int{-2147483648}
                    _ = v11
                default:
                    p0 += p0
                }
            }
        case 32768:
            if (-558 > 1) {
                p0 = p0
                p0 = p0
            }
        default:
            var v12 int64 = 1
            _ = v12
        }
        v13 := GenS2{F0: (false || true), F1: (^(p0))}
        _ = fi4
        _ = v13
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    return 2
}
func genFn1(p0 string, p1 int32) uint {
    fmt.Println(p0)
    var v14 float32 = 0.1
    _ = p0
    _ = p1
    _ = v14
    return 0
}

func main() {
    var v1 int8 = -128
    var v2 bool = true
    v3 := make(map[uint32]int)
    v3[2147483647] = 42
    for fi15 := 0; fi15 < 5; fi15++ {
        switch ("hello" + "\n\t") {
        case "hello":
            var v16 int16 = -421
            _ = v16
        case "":
            if (10 <= (-(v1))) {
                fmt.Println(v2, v2, v1)
                v17 := GenS1{F0: "hello"}
                var v18 float64 = 0.1
                _ = v17
                _ = v18
            }
        case "foo bar":
            var v19 float64 = 0.0
            _ = v19
        default:
            for fi20 := 0; fi20 < 11; fi20++ {
                if (531 > 18446744073709551615) {
                    v21 := GenS2{F0: false, F1: 4294967295}
                    _ = v21
                }
                if (128 <= 184) {
                    v1 = int8(fi20)
                    var v22 int64 = -822
                    _ = v22
                } else {
                    fmt.Println(v2)
                    v23 := GenS3{F0: -9223372036854775808, F1: 255, F2: 4294967295}
                    v24 := GenS2{F0: (true || ((255 <= 65535) || (!(true && v2)))), F1: 42}
                    var v25 uint32 = 2
                    _ = v23
                    _ = v24
                    _ = v25
                }
                _ = fi20
            }
        }
        for fi26 := 0; fi26 < 14; fi26++ {
            if v2 {
                var v27 int8 = 100
                if false {
                    var v28 uint16 = 456
                    if v2 {
                        switch 320 {
                        case 868:
                            if false {
                                v1 += v1
                                v29 := v3[876]
                                v30 := GenS3{F0: 1, F1: 42, F2: 1}
                                _ = v29
                                _ = v30
                            }
                        case 213:
                            fmt.Println(fi15)
                        }
                    } else {
                        v2 = ((fi15 * fi26) <= fi15)
                        v2 = false
                    }
                    v28 -= v28
                    _ = v28
                }
                _ = v27
            } else {
                v2 = (127 <= fi15)
            }
            _ = fi26
        }
        if v2 {
            var v31 float64 = float64(0.1)
            var v32 uint8 = 0
            _ = v31
            _ = v32
        } else {
            for fi33 := 0; fi33 < 13; fi33++ {
                v34 := GenS4{F0: "", F1: "foo bar", F2: 552}
                switch "" {
                case "foo bar":
                    v1 += v1
                case "":
                    if v2 {
                        if v2 {
                            v1 = 2
                            fmt.Println(v1)
                            v35 := GenS5{F0: 100}
                            _ = v35
                        }
                    }
                case "\n\t":
                    v36 := GenS5{F0: (-(v1))}
                    _ = v36
                default:
                    var v37 string = "a"
                    _ = v37
                }
                var v38 uint8 = 149
                _ = fi33
                _ = v34
                _ = v38
            }
        }
        _ = fi15
    }
    v1 += (v1 % ((v1) | 1))
    fmt.Println(v1, v2)
    v1 = v1
    v39 := []int{183, -1}
    var v40 uint8 = 128
    v40 = v40
    switch 1 {
    case 0:
        var v41 uint32 = 829
        _ = v41
    case -1:
        v1 = (-(v1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v40
}
