package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 int; F1 uint16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint64; F1 bool }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int64; F1 int16 }
type GenS9 struct { F0 uint16; F1 uint64; F2 uint8 }
func genFn0(p0 string) uint32 {
    p0 = ((("a" + ((p0 + "foo bar") + p0)) + ((p0 + "") + p0)) + "foo bar")
    var v6 uint16 = 1
    if (-168 == 1) {
        p0 = (p0 + p0)
    }
    _ = p0
    _ = v6
    return 1
}
func genFn1(p0 uint, p1 int8) uint32 {
    var v7 int16 = 0
    var v8 int32 = 2147483647
    v9 := GenS1{F0: ""}
    v10 := GenS1{F0: ""}
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := []string{"hello", "\n\t", "foo bar", "foo bar"}
    var v2 uint = 2
    var v3 int16 = 1000
    var v4 uint = v2
    v5 := []int{127, 340, -2147483648}
    fmt.Println(v2, v3)
    fmt.Println(v2, v2, v2)
    var v11 int64 = 0
    for fi12 := 0; fi12 < 1; fi12++ {
        switch fi12 {
        case 9223372036854775807:
            v13 := GenS2{F0: false, F1: false}
            _ = v13
        case 2147483647:
            if ((true && (!(1.5 <= 0.0))) && false) {
                if true {
                    for fi14 := 0; fi14 < 7; fi14++ {
                        switch "" {
                        case "foo bar":
                            var v15 int8 = 127
                            _ = v15
                        default:
                            if ((true && false) && (-936 >= -1)) {
                                v4 = (^(v2))
                            } else {
                                v16 := GenS3{F0: 81, F1: 635}
                                v3 = 1000
                                v3 = 907
                                _ = v16
                            }
                        }
                        v11 = (v11 / ((9223372036854775807) | 1))
                        v1 = append(v1, "a")
                        _ = fi14
                    }
                }
            }
        default:
            fmt.Println(v4)
        }
        _ = len(v5)
        _ = fi12
    }
    v11 -= -9223372036854775808
    switch -572 {
    case -871:
        for fi17 := 0; fi17 < 7; fi17++ {
            v18 := GenS3{F0: 178, F1: 1}
            _ = fi17
            _ = v18
        }
    case 32768:
        for fi19 := 0; fi19 < 8; fi19++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi19
        }
    case -451:
        for fi20 := 0; fi20 < 10; fi20++ {
            if (false || false) {
                v4 = 42
            } else {
                if false {
                    v4 -= (v2 << ((v2) & 63))
                    v2 = (v4 & v2)
                }
                v21 := GenS2{F0: ((1e10 - ((0.0 * 1e10) + -2.25)) >= 3.14159), F1: false}
                for fi22 := 0; fi22 < 7; fi22++ {
                    if false {
                        v23 := GenS4{F0: (fi22 % (((fi22 + (fi22 | fi20))) | 1)), F1: 124}
                        _ = v23
                    }
                    var v24 int8 = -1
                    var v25 int = fi22
                    _ = fi22
                    _ = v24
                    _ = v25
                }
                _ = v21
            }
            for fi26 := 0; fi26 < 3; fi26++ {
                switch (fi20 ^ (fi26 % ((fi20) | 1))) {
                case -1:
                    v11 -= v11
                    v4 = v4
                    v4 += v4
                case 0:
                    var v27 int8 = 42
                    _ = v27
                }
                var v28 int64 = v11
                var v29 int32 = -1
                v2 = 255
                switch -1 {
                case 829:
                    v28 = v11
                case -162:
                    var v30 int = fi20
                    _ = v30
                default:
                    if false {
                        var v31 bool = (396 == (v4 >> ((v4) & 63)))
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = v31
                    } else {
                        fmt.Println(v29, v4)
                        if true {
                            var v32 float32 = 3.14159
                            _ = v32
                        } else {
                            if len(v5) > 0 {
                                v5[0] = fi26
                            }
                            fmt.Println(v28)
                            switch "hello" {
                            case "hello":
                                v33 := GenS4{F0: fi20, F1: 928}
                                _ = v33
                            case "\n\t":
                                if (((("foo bar" + ("a" + (((("\n\t" + ("foo bar" + "")) + "a") + "hello") + "\n\t"))) + ("foo bar" + ("\n\t" + ("hello" + ("a" + ("" + "hello")))))) + "") == ("hello" + ((("\n\t" + "hello") + ("\n\t" + ("" + "hello"))) + (("hello" + "foo bar") + "\n\t")))) {
                                    var v34 string = "hello"
                                    var v35 uint32 = uint32(v4)
                                    if (v2 > (v4 + v4)) {
                                        if len(v1) > 0 {
                                            v1[0] = (v34 + v34)
                                        }
                                    } else {
                                        v36 := GenS5{F0: uint64(v35)}
                                        v2 += v4
                                        _ = v36
                                    }
                                    _ = v34
                                    _ = v35
                                } else {
                                    switch "foo bar" {
                                    case "\n\t":
                                        v4 -= (v4 + (v4 % (((v4 ^ v4)) | 1)))
                                    case "foo bar":
                                        if false {
                                            if (-1 == v3) {
                                                fmt.Println(fi20, v3)
                                                v28 -= v28
                                                v5 = append(v5, fi26)
                                            }
                                            v5 = append(v5, fi26)
                                            v11 -= int64(v3)
                                            v37 := GenS6{F0: 765, F1: false}
                                            var v38 uint64 = 4294967295
                                            _ = v37
                                            _ = v38
                                        }
                                    }
                                    v39 := GenS4{F0: 2147483647, F1: 65535}
                                    fmt.Println(v11, v4)
                                    if (!(164 <= 65535)) {
                                        v28 = v11
                                        v40 := GenS2{F0: true, F1: (v29 >= (v29 & (-(v29))))}
                                        if true {
                                            fmt.Println(v29, fi20, fi26)
                                            if true {
                                                v41 := GenS2{F0: false, F1: false}
                                                var v42 uint = 394
                                                _ = v41
                                                _ = v42
                                            }
                                            v3 += (v3 >> ((18446744073709551615) & 63))
                                            var v43 int8 = 0
                                            _ = v43
                                        }
                                        var v44 int8 = 0
                                        fmt.Println(v4, fi20, fi26)
                                        _ = v40
                                        _ = v44
                                    }
                                    if (1 == 2147483647) {
                                        v11 = v28
                                    } else {
                                        var v45 float64 = (float64(1e10) - 1.5)
                                        _ = v45
                                    }
                                    fmt.Println(v11, v29)
                                    _ = v39
                                }
                            case "":
                                v46 := GenS3{F0: 141, F1: 2147483648}
                                _ = v46
                            }
                        }
                    }
                }
                _ = fi26
                _ = v28
                _ = v29
            }
            _ = fi20
        }
    }
    _ = len(v5)
    v47 := GenS7{F0: true}
    v2 = 974
    if (float64(-2.25) > -2.25) {
        var v48 float32 = 0.0
        _ = v48
    } else {
        if (!(((255 <= 32767) && (false || true)) && false)) {
            var v49 int16 = -32768
            var v50 uint8 = 127
            _ = v49
            _ = v50
        }
    }
    v51 := GenS2{F0: (false && ((1.5 <= 1.5) || (int32(v3) == 2147483647))), F1: (1e10 <= -2.25)}
    var v52 float64 = 0.1
    var v53 uint16 = 32767
    if true {
        var v54 float32 = 0.1
        var v55 uint32 = 156
        _ = v54
        _ = v55
    } else {
        _ = len(v5)
        v2 = uint(v53)
    }
    for fi56 := 0; fi56 < 19; fi56++ {
        v53 -= v53
        if false {
            v57 := GenS4{F0: fi56, F1: (v53 ^ (^(v53)))}
            v58 := GenS2{F0: ("\n\t" > "\n\t"), F1: true}
            v52 = (((v52 * 0.0) + (v52 - -2.25)) - float64(((((3.14159 * (0.0 + 3.14159)) + 1.5) * 3.14159) + -2.25)))
            _ = v57
            _ = v58
        } else {
            for fi59 := 0; fi59 < 9; fi59++ {
                var v60 int32 = -2147483648
                v61 := GenS5{F0: 4294967295}
                _ = fi59
                _ = v60
                _ = v61
            }
            v62 := GenS8{F0: 9223372036854775807, F1: (v3 - v3)}
            if ("a" <= "\n\t") {
                _ = len(v5)
                v2 -= v2
            } else {
                for fi63 := 0; fi63 < 18; fi63++ {
                    var v64 uint64 = 0
                    v65 := GenS4{F0: fi56, F1: (v53 / ((v53) | 1))}
                    if false {
                        v2 = v2
                        switch "a" {
                        case "":
                            switch "a" {
                            case "\n\t":
                                if (false && (-634 < v3)) {
                                    var v66 float64 = (v52 / (0.0 * v52))
                                    var v67 uint64 = uint64(v4)
                                    _ = v66
                                    _ = v67
                                } else {
                                    fmt.Println(v64, v2, v64)
                                    if (!false) {
                                        var v68 bool = (false || true)
                                        var v69 uint = (v4 | (v2 % (((v4 | (v4 << ((v2) & 63)))) | 1)))
                                        v68 = v68
                                        _ = v68
                                        _ = v69
                                    }
                                    var v70 int = fi63
                                    _ = v70
                                }
                            case "hello":
                                v53 += (v53 >> ((v4) & 63))
                            case "":
                                var v71 int64 = v11
                                v53 -= (v53 ^ v53)
                                if false {
                                    var v72 int8 = -128
                                    _ = v72
                                } else {
                                    var v73 float64 = ((float64(3.14159) - v52) + float64((((0.1 - (0.1 / float32(3.14159))) - float32((float64(1.5) / 1.0))) * 3.14159)))
                                    var v74 int16 = v3
                                    v75 := GenS9{F0: (v53 << ((v2) & 63)), F1: v64, F2: 0}
                                    _ = v73
                                    _ = v74
                                    _ = v75
                                }
                                _ = v71
                            default:
                                v53 = v53
                            }
                        case "a":
                            var v76 int32 = -819
                            _ = v76
                        }
                    } else {
                        var v77 uint = 192
                        _ = v77
                    }
                    _ = fi63
                    _ = v64
                    _ = v65
                }
                for fi78 := 0; fi78 < 12; fi78++ {
                    v79 := GenS6{F0: 78, F1: (2147483647 == -530)}
                    fmt.Println(fi78, v52, v3)
                    v2 += (v2 | (v2 << ((815) & 63)))
                    _ = fi78
                    _ = v79
                }
                if ((fi56 << ((516) & 63)) < (fi56 * (fi56 ^ fi56))) {
                    v52 = (v52 - 3.14159)
                }
            }
            _ = v62
        }
        v80 := GenS5{F0: 1}
        _ = fi56
        _ = v80
    }
    var v81 int8 = 127
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v51
    _ = v52
    _ = v53
    _ = v81
}
