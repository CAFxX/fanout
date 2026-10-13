package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 uint16; F2 int8 }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 int64; F1 uint16 }
type GenS5 struct { F0 int64; F1 uint16 }
type GenS6 struct { F0 uint8; F1 bool }
type GenS7 struct { F0 uint }
func genFn0(p0 int16, p1 int64, p2 uint16) uint32 {
    v5 := GenS1{F0: p1}
    var v6 uint = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 846
}
func genFn1(p0 int) string {
    p0 += (p0 / ((p0) | 1))
    p0 += (-(p0))
    var v7 string = ("hello" + "hello")
    var v8 int8 = -70
    v8 += 100
    _ = p0
    _ = v7
    _ = v8
    return v7
}

func main() {
    var v1 float32 = 0.0
    v2 := []uint32{2147483648, 31, 1, 880}
    v3 := make(map[int]int)
    v4 := []uint32{65535, 0}
    switch "hello" {
    case "hello":
        for fi9 := 0; fi9 < 4; fi9++ {
            v1 = v1
            v1 = 0.0
            _ = fi9
        }
    }
    v1 -= v1
    v10 := GenS2{F0: 4294967295, F1: 42, F2: 1}
    if true {
        if (!true) {
            v1 = v1
            if len(v4) > 0 {
                v4[0] = 4294967295
            }
            v1 = v1
        } else {
            for fi11 := 0; fi11 < 3; fi11++ {
                v1 = 1.5
                v12 := GenS2{F0: (-(fi11)), F1: 42, F2: 102}
                v1 += 1.5
                _ = fi11
                _ = v12
            }
            v13 := GenS3{F0: 1, F1: 1}
            _ = v13
        }
        switch "foo bar" {
        case "":
            if (4294967295 < 237) {
                v14 := GenS4{F0: 366, F1: 65535}
                fmt.Println(v1)
                _ = v14
            }
        case "\n\t":
            fmt.Println(v1)
        case "hello":
            for fi15 := 0; fi15 < 4; fi15++ {
                v1 = (v1 / 1.5)
                fmt.Println(fi15, v1)
                _ = fi15
            }
        }
        var v16 int32 = 506
        _ = v16
    }
    if (772 < -157) {
        fmt.Println(v1)
    }
    switch -9223372036854775808 {
    case -1:
        fmt.Println(v1)
    case -9223372036854775808:
        for fi17 := 0; fi17 < 13; fi17++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi17
        }
    }
    var v18 int16 = 1000
    switch 32768 {
    case -885:
        for fi19 := 0; fi19 < 1; fi19++ {
            var v20 int64 = 9223372036854775807
            _ = fi19
            _ = v20
        }
    default:
        v1 = v1
    }
    v18 -= v18
    var v21 bool = (18446744073709551615 != 42)
    var v22 uint64 = 18446744073709551615
    if (1 == 42) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if ((^(v22)) > v22) {
            v23 := GenS4{F0: 1, F1: 2}
            fmt.Println(v22, v22)
            _ = v23
        } else {
            v24 := GenS4{F0: -36, F1: 42}
            for fi25 := 0; fi25 < 7; fi25++ {
                _ = len(v2)
                for fi26 := 0; fi26 < 2; fi26++ {
                    v27 := GenS4{F0: 1, F1: 2}
                    fmt.Println(fi26)
                    v18 -= (^(v18))
                    _ = fi26
                    _ = v27
                }
                _ = fi25
            }
            _ = v24
        }
        v18 += -32768
    }
    var v28 uint16 = 255
    var v29 uint32 = 42
    switch -628 {
    case 601:
        var v30 int8 = 77
        _ = v30
    case 1:
        for fi31 := 0; fi31 < 7; fi31++ {
            for fi32 := 0; fi32 < 9; fi32++ {
                var v33 int16 = 915
                _ = fi32
                _ = v33
            }
            v34 := GenS5{F0: 42, F1: (v28 / ((138) | 1))}
            _ = fi31
            _ = v34
        }
    case 707:
        switch -432 {
        case -9223372036854775808:
            v35 := GenS6{F0: 42, F1: v21}
            _ = v35
        default:
            if v21 {
                v36 := GenS6{F0: 255, F1: false}
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v36
            } else {
                var v37 float32 = v1
                if true {
                    var v38 uint32 = genFn0(32767, int64(v22), 42)
                    _ = len(v3)
                    v37 = (v37 * v37)
                    _ = v38
                }
                _ = v37
            }
        }
    default:
        var v39 int8 = -118
        _ = v39
    }
    if v21 {
        v1 += (v1 / (float32((float64(v1) + -2.25)) / v1))
    } else {
        if ((!(((v18 + (^(v18))) < (v18 | 42)) && v21)) && (-2147483648 != 1)) {
            _ = len(v3)
            v40 := v3[32767]
            _ = v40
        } else {
            if (int(v28) < -1) {
                if (true || v21) {
                    var v41 int32 = 42
                    _ = v41
                } else {
                    if (!v21) {
                        var v42 uint64 = (v22 * 451)
                        var v43 int = 2147483647
                        _ = v42
                        _ = v43
                    }
                }
            }
        }
        if true {
            for fi44 := 0; fi44 < 8; fi44++ {
                if v21 {
                    v22 = (^(v22))
                    var v45 uint8 = 122
                    _ = v45
                }
                _ = fi44
            }
        } else {
            for fi46 := 0; fi46 < 3; fi46++ {
                for fi47 := 0; fi47 < 1; fi47++ {
                    _ = len(v2)
                    _ = fi47
                }
                v48 := GenS5{F0: 1, F1: (^(v28))}
                _ = fi46
                _ = v48
            }
            _ = len(v3)
            if (0 > 2) {
                fmt.Println(v18, v21)
            }
        }
        v49 := GenS7{F0: 354}
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
