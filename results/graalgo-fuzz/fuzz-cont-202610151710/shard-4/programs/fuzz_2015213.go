package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 bool; F1 uint64; F2 int16 }
type GenS3 struct { F0 string; F1 int8; F2 bool }
type GenS4 struct { F0 int8; F1 int8; F2 int32 }
type GenS5 struct { F0 bool; F1 int8; F2 uint32 }
type GenS6 struct { F0 int8; F1 int; F2 uint32 }
type GenS7 struct { F0 uint8; F1 int16; F2 uint64 }
func genFn0(p0 int16, p1 int64, p2 int32) int {
    p0 = p0
    if true {
        p1 = (-(p1))
    } else {
        v6 := []int{4294967295, 9223372036854775807, 725, 751}
        switch -461 {
        case 42:
            v7 := GenS1{F0: p2, F1: -1}
            _ = v7
        case -9223372036854775808:
            var v8 uint64 = 0
            _ = v8
        case 47:
            fmt.Println(p1)
        }
        _ = v6
    }
    switch int(p1) {
    case -1:
        p2 = p2
    default:
        var v9 int16 = p0
        _ = v9
    }
    p0 = (p0 * (-(p0)))
    _ = p0
    _ = p1
    _ = p2
    return 259
}
func genFn1(p0 uint16, p1 int16) uint64 {
    var v10 uint16 = 0
    switch (("foo bar" + "\n\t") + "a") {
    case "foo bar":
        v10 -= v10
    default:
        v11 := GenS2{F0: (!false), F1: 42, F2: (p1 | -1)}
        _ = v11
    }
    var v12 float64 = 3.14159
    switch genFn0((p1 / ((p1) | 1)), 0, 2147483647) {
    case -9223372036854775808:
        v13 := GenS3{F0: ((((("\n\t" + "\n\t") + "hello") + "\n\t") + (("" + "a") + "hello")) + ""), F1: 0, F2: (!true)}
        _ = v13
    case 375:
        p0 += (p0 / ((v10) | 1))
    default:
        v14 := GenS2{F0: (false || (!true)), F1: 15, F2: p1}
        _ = v14
    }
    fmt.Println(v12, p0)
    _ = p0
    _ = p1
    _ = v10
    _ = v12
    return 711
}

func main() {
    var v1 uint = 876
    var v2 int8 = 100
    var v3 uint16 = 957
    v4 := []string{"a"}
    v5 := make(map[uint32]int)
    fmt.Println(v2)
    v3 = 65535
    for fi15 := 0; fi15 < 13; fi15++ {
        v2 += (v2 * (v2 ^ v2))
        v3 -= v3
        _ = len(v4)
        _ = fi15
    }
    if (255 < 170) {
        v2 -= v2
        v16 := GenS2{F0: false, F1: 42, F2: 1000}
        _ = v16
    } else {
        var v17 uint16 = v3
        fmt.Println(v1, v1, v2)
        for fi18 := 0; fi18 < 18; fi18++ {
            var v19 uint16 = 319
            _ = fi18
            _ = v19
        }
        _ = v17
    }
    var v20 string = "hello"
    for fi21 := 0; fi21 < 12; fi21++ {
        var v22 int8 = -1
        var v23 string = (v20 + v20)
        _ = fi21
        _ = v22
        _ = v23
    }
    switch 721 {
    case 127:
        v24 := GenS2{F0: ((v3 / (((^(v3))) | 1)) == (v3 - uint16(v1))), F1: 764, F2: 1000}
        _ = v24
    case 32767:
        v25 := GenS2{F0: (-32768 <= 653), F1: 18446744073709551615, F2: 239}
        _ = v25
    case 42:
        if (false && ((198 <= 1) || true)) {
            for fi26 := 0; fi26 < 6; fi26++ {
                v5[0] = (fi26 | -2147483648)
                var v27 int64 = 0
                fmt.Println(v3, v1, v1)
                _ = fi26
                _ = v27
            }
            var v28 int = int(v3)
            _ = v28
        } else {
            var v29 uint16 = (v3 / ((v3) | 1))
            _ = v29
        }
    default:
        switch v20 {
        case "":
            switch ("\n\t" + "foo bar") {
            case "foo bar":
                switch -272 {
                case 2147483647:
                    v2 += 100
                case -195:
                    if (!(!true)) {
                        v30 := GenS1{F0: -2147483648, F1: 761}
                        _ = v30
                    }
                case -589:
                    if (false && false) {
                        v31 := GenS2{F0: false, F1: 2, F2: 1000}
                        _ = v31
                    } else {
                        if false {
                            var v32 int32 = -2147483648
                            _ = v32
                        } else {
                            var v33 uint = v1
                            fmt.Println(v20, v33)
                            _ = v33
                        }
                        for fi34 := 0; fi34 < 19; fi34++ {
                            for fi35 := 0; fi35 < 1; fi35++ {
                                var v36 int32 = int32(v2)
                                v37 := GenS3{F0: v20, F1: v2, F2: (0 != -32768)}
                                if ((!false) && (!(!false))) {
                                    v1 -= v1
                                    switch v20 {
                                    case "hello":
                                        if true {
                                            var v38 int8 = (v2 << (((v1 - 255)) & 63))
                                            v39 := GenS4{F0: -128, F1: int8(v36), F2: -953}
                                            var v40 bool = ((^(fi35)) == (fi35 + (fi35 + (-(fi34)))))
                                            _ = v38
                                            _ = v39
                                            _ = v40
                                        }
                                    case "\n\t":
                                        var v41 uint = v1
                                        _ = v41
                                    default:
                                        v42 := GenS5{F0: (1 >= 127), F1: v2, F2: 1}
                                        _ = v42
                                    }
                                    if len(v4) > 0 {
                                        v4[0] = "a"
                                    }
                                }
                                _ = fi35
                                _ = v36
                                _ = v37
                            }
                            _ = fi34
                        }
                    }
                }
            case "":
                for fi43 := 0; fi43 < 12; fi43++ {
                    v3 += v3
                    _ = fi43
                }
            default:
                v3 += (^(v3))
            }
        default:
            fmt.Println(v2, v2)
        }
    }
    fmt.Println(v2, v20, v20)
    v2 = (-(v2))
    v20 = v20
    for fi44 := 0; fi44 < 3; fi44++ {
        if ((!(0.1 == 3.14159)) || true) {
            for fi45 := 0; fi45 < 2; fi45++ {
                _ = len(v4)
                _ = fi45
            }
        }
        var v46 int32 = 2147483647
        var v47 float64 = 3.14159
        _ = fi44
        _ = v46
        _ = v47
    }
    v2 += v2
    var v48 uint = 18446744073709551615
    fmt.Println(v48, v20)
    var v49 int64 = 9223372036854775807
    switch "a" {
    case "":
        fmt.Println(v1)
    }
    switch 91 {
    case 128:
        v50 := GenS6{F0: -128, F1: 9223372036854775807, F2: 42}
        _ = v50
    }
    v51 := GenS7{F0: uint8(v1), F1: -463, F2: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v51
}
