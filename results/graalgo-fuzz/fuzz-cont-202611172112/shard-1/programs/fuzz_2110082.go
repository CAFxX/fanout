package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 int64 }
type GenS2 struct { F0 string; F1 string; F2 bool }
type GenS3 struct { F0 uint; F1 int32; F2 int16 }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 uint8; F1 int16; F2 int8 }
type GenS6 struct { F0 int16; F1 bool }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint8 }
func genFn0(p0 int32) bool {
    v7 := []int{0}
    fmt.Println(p0)
    switch ("a" + "hello") {
    case "a":
        var v8 string = "foo bar"
        _ = v8
    case "hello":
        var v9 uint8 = 148
        _ = v9
    case "foo bar":
        var v10 float32 = ((1e10 * (float32(((0.1 / (1.5 + -2.25)) - (3.14159 / 3.14159))) + 0.1)) / -2.25)
        _ = v10
    default:
        _ = len(v7)
    }
    var v11 int8 = 42
    p0 = p0
    _ = p0
    _ = v11
    _ = v7
    return (false && (42 >= -364))
}
func genFn1(p0 int64) int32 {
    if false {
        for fi12 := 0; fi12 < 18; fi12++ {
            if false {
                switch "" {
                case "a":
                    for fi13 := 0; fi13 < 2; fi13++ {
                        p0 += (p0 * 42)
                        _ = fi13
                    }
                case "hello":
                    switch "a" {
                    case "foo bar":
                        for fi14 := 0; fi14 < 19; fi14++ {
                            v15 := GenS1{F0: (false || true), F1: -32768, F2: (p0 * -1)}
                            _ = fi14
                            _ = v15
                        }
                    default:
                        p0 = 703
                    }
                default:
                    if (true && (-2.25 >= 0.1)) {
                        if true {
                            v16 := GenS1{F0: ((true && (!(((fi12 >> ((901) & 63)) < 0) && true))) && (2 >= 65535)), F1: 42, F2: (p0 >> ((18446744073709551615) & 63))}
                            _ = v16
                        }
                        if (-602 <= -271) {
                            var v17 int = fi12
                            for fi18 := 0; fi18 < 15; fi18++ {
                                v19 := []int{32768}
                                var v20 string = ("foo bar" + "hello")
                                v21 := GenS1{F0: (127 >= 0), F1: -1, F2: -1}
                                _ = fi18
                                _ = v19
                                _ = v20
                                _ = v21
                            }
                            _ = v17
                        }
                        var v22 int8 = 51
                        _ = v22
                    }
                }
                p0 = (p0 & (p0 / (((p0 << ((255) & 63))) | 1)))
                p0 -= p0
            }
            _ = fi12
        }
        if (!((1 <= 2147483647) && false)) {
            switch -1 {
            case 0:
                p0 = p0
            case 2147483647:
                fmt.Println(p0)
            }
            var v23 uint32 = 4294967295
            v24 := GenS1{F0: (v23 == 2147483647), F1: -1, F2: (p0 | p0)}
            _ = v23
            _ = v24
        }
        v25 := GenS1{F0: true, F1: 0, F2: (p0 - p0)}
        _ = v25
    }
    p0 = 9223372036854775807
    fmt.Println(p0)
    p0 -= (p0 | p0)
    _ = p0
    return 1
}

func main() {
    var v1 uint = 42
    var v2 int16 = -1
    var v3 uint16 = 728
    var v4 uint64 = 322
    var v5 uint8 = 127
    v6 := []uint32{0, 2147483648, 4294967295}
    if len(v6) > 0 {
        v6[0] = 65535
    }
    switch ((("foo bar" + "\n\t") + (((((("hello" + ("foo bar" + "foo bar")) + "foo bar") + "\n\t") + ("a" + "")) + "a") + "foo bar")) + "") {
    case "":
        var v26 uint32 = 4294967295
        _ = v26
    case "foo bar":
        var v27 string = ("foo bar" + "hello")
        _ = v27
    case "hello":
        v28 := GenS1{F0: (!false), F1: (v2 % ((-1) | 1)), F2: 1}
        _ = v28
    }
    if (!false) {
        v2 += (v2 | 32767)
    } else {
        v29 := GenS1{F0: true, F1: v2, F2: -9223372036854775808}
        if len(v6) > 0 {
            v6[0] = 559
        }
        _ = v29
    }
    switch "a" {
    case "hello":
        if false {
            v1 += (v1 / (((v1 >> (((^(v1))) & 63))) | 1))
            if len(v6) > 0 {
                v6[0] = 65535
            }
            v3 += uint16(v4)
        } else {
            v2 = (v2 & (v2 * v2))
        }
    case "\n\t":
        fmt.Println(v5)
    }
    if true {
        switch 32768 {
        case 755:
            for fi30 := 0; fi30 < 12; fi30++ {
                fmt.Println(fi30, v5, v4)
                fmt.Println(fi30)
                v1 += 18446744073709551615
                _ = fi30
            }
        case -9223372036854775808:
            _ = len(v6)
        case 9223372036854775807:
            v4 = uint64(v1)
        }
    } else {
        if false {
            v31 := GenS1{F0: (((v2 ^ (-(v2))) < (^(v2))) || true), F1: v2, F2: 917}
            var v32 int64 = 0
            fmt.Println(v32, v3, v4)
            _ = v31
            _ = v32
        }
        if true {
            v33 := GenS2{F0: "foo bar", F1: ("foo bar" + "foo bar"), F2: (false && ((1 == 4294967295) && ((!false) || true)))}
            v34 := GenS3{F0: 0, F1: 0, F2: (v2 + (v2 % ((42) | 1)))}
            switch "" {
            case "":
                var v35 int64 = 1
                _ = v35
            case "a":
                for fi36 := 0; fi36 < 17; fi36++ {
                    for fi37 := 0; fi37 < 3; fi37++ {
                        v38 := GenS4{F0: 42, F1: v4}
                        v1 = (v1 | (v1 % (((v1 * (v1 ^ 500))) | 1)))
                        _ = len(v6)
                        _ = fi37
                        _ = v38
                    }
                    _ = fi36
                }
            default:
                if (1.5 >= 3.14159) {
                    _ = len(v6)
                    if len(v6) > 0 {
                        v6[0] = 280
                    }
                    var v39 int = -1
                    _ = v39
                }
            }
            _ = v33
            _ = v34
        }
        var v40 int8 = 127
        _ = v40
    }
    if false {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        v41 := GenS5{F0: 2, F1: v2, F2: 0}
        var v42 int32 = 42
        _ = v41
        _ = v42
    } else {
        var v43 int = 1
        v44 := GenS6{F0: (v2 ^ v2), F1: (!true)}
        _ = v43
        _ = v44
    }
    fmt.Println(v3, v1)
    if ((v4 ^ v4) < v4) {
        v2 += (v2 | 662)
        if (0 != int(v3)) {
            for fi45 := 0; fi45 < 7; fi45++ {
                var v46 string = "\n\t"
                if (100 > -1) {
                    v47 := GenS7{F0: genFn1(-155)}
                    var v48 float64 = 0.0
                    _ = v47
                    _ = v48
                }
                _ = fi45
                _ = v46
            }
        }
    }
    if false {
        if false {
            v49 := GenS5{F0: uint8(v1), F1: (v2 << ((v1) & 63)), F2: 37}
            _ = v49
        }
        for fi50 := 0; fi50 < 13; fi50++ {
            v2 = v2
            for fi51 := 0; fi51 < 6; fi51++ {
                if (!(0 < 127)) {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    v52 := GenS8{F0: (v2 | v2)}
                    _ = v52
                } else {
                    switch "foo bar" {
                    case "hello":
                        v4 = (v4 + v4)
                    case "":
                        if ((!true) || true) {
                            fmt.Println(v2, v2, v2)
                            if ((!(true || false)) || ((1 <= (v2 | int16(v4))) || (86 < 57))) {
                                fmt.Println(v3, v4, v3)
                            }
                        }
                    }
                }
                var v53 float64 = 0.1
                v1 -= 760
                _ = fi51
                _ = v53
            }
            if false {
                var v54 int64 = -972
                v6 = append(v6, 958)
                _ = v54
            } else {
                v55 := GenS8{F0: (v2 << ((v1) & 63))}
                switch (fi50 + fi50) {
                case -542:
                    v1 = (^(v1))
                case -9223372036854775808:
                    var v56 bool = (1e10 >= 1e10)
                    _ = v56
                case 1:
                    fmt.Println(fi50, v2)
                }
                _ = v55
            }
            _ = fi50
        }
        switch "" {
        case "\n\t":
            var v57 uint16 = (v3 + (v3 & v3))
            _ = v57
        case "hello":
            var v58 int = 9223372036854775807
            _ = v58
        default:
            if len(v6) > 0 {
                v6[0] = 1
            }
        }
    } else {
        _ = len(v6)
        v3 = (v3 ^ 1)
    }
    v59 := GenS5{F0: v5, F1: int16(v4), F2: 127}
    if (854 < 9223372036854775807) {
        var v60 bool = (-1 > (v2 % (((v2 + v2)) | 1)))
        var v61 int32 = 2147483647
        var v62 string = ""
        _ = v60
        _ = v61
        _ = v62
    }
    for fi63 := 0; fi63 < 9; fi63++ {
        v4 += v4
        v64 := GenS9{F0: 255}
        _ = fi63
        _ = v64
    }
    v65 := GenS3{F0: 65535, F1: -776, F2: v2}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if (12 > 0) {
        v66 := GenS2{F0: "a", F1: "\n\t", F2: false}
        for fi67 := 0; fi67 < 7; fi67++ {
            var v68 int32 = genFn1(42)
            var v69 int32 = (^(v68))
            var v70 float64 = (1e10 * 1e10)
            _ = fi67
            _ = v68
            _ = v69
            _ = v70
        }
        _ = v66
    }
    v5 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v6
    _ = v65
}
