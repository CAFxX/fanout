package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 uint; F2 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint; F1 int32; F2 uint64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int32; F1 int; F2 uint64 }
type GenS7 struct { F0 uint; F1 int16 }
type GenS8 struct { F0 string; F1 uint64; F2 int16 }
type GenS9 struct { F0 uint8; F1 int32 }
type GenS10 struct { F0 int8; F1 int32 }
func genFn0() uint8 {
    var v5 uint16 = 42
    fmt.Println(v5)
    _ = v5
    return 2
}
func genFn1(p0 uint8) string {
    fmt.Println(p0)
    v6 := GenS1{F0: 160}
    p0 = p0
    _ = p0
    _ = v6
    return "foo bar"
}

func main() {
    var v1 float32 = 1e10
    var v2 int = 9
    var v3 int32 = -2147483648
    v4 := []string{}
    v3 = (v3 << ((0) & 63))
    var v7 float32 = ((v1 / 3.14159) + 0.0)
    switch 128 {
    case -9223372036854775808:
        fmt.Println(v7, v1)
    }
    if len(v4) > 0 {
        v4[0] = ("hello" + ("\n\t" + "a"))
    }
    switch (v2 % ((128) | 1)) {
    case 32768:
        v7 = 0.1
    case -423:
        for fi8 := 0; fi8 < 2; fi8++ {
            v9 := GenS1{F0: (fi8 % ((fi8) | 1))}
            _ = len(v4)
            v10 := GenS1{F0: (v2 | fi8)}
            _ = fi8
            _ = v10
            _ = v9
        }
    case 32767:
        switch v2 {
        case -1:
            v11 := GenS1{F0: v2}
            _ = v11
        case -9223372036854775808:
            switch v2 {
            case -851:
                if false {
                    v3 += (v3 << ((255) & 63))
                    v1 += float32((((1.5 / (1e10 * (1.5 + -2.25))) - 0.1) - (0.0 * 3.14159)))
                } else {
                    v4 = append(v4, "foo bar")
                    v3 = (v3 | v3)
                }
            case 32768:
                for fi12 := 0; fi12 < 13; fi12++ {
                    for fi13 := 0; fi13 < 1; fi13++ {
                        var v14 float32 = v7
                        _ = fi13
                        _ = v14
                    }
                    _ = fi12
                }
            default:
                switch (((("a" + "hello") + ("\n\t" + "")) + "foo bar") + "\n\t") {
                case "hello":
                    fmt.Println(v3, v3)
                default:
                    v1 -= v7
                }
            }
        case 32767:
            if (1 == 374) {
                v7 = v1
                switch v2 {
                case 672:
                    if ((("foo bar" + (("a" + ("" + "\n\t")) + ("a" + ("foo bar" + "a")))) + "") < "") {
                        v7 += (v1 / v1)
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v15 float64 = 0.1
                        _ = v15
                    }
                default:
                    for fi16 := 0; fi16 < 7; fi16++ {
                        if true {
                            v17 := GenS2{F0: -128, F1: 1, F2: 2}
                            v18 := GenS2{F0: 42, F1: 4294967295, F2: 42}
                            var v19 bool = false
                            _ = v17
                            _ = v18
                            _ = v19
                        } else {
                            if (!(255 < 65535)) {
                                fmt.Println(v3, fi16)
                                fmt.Println(v2, v2, v2)
                                var v20 int8 = -128
                                _ = v20
                            }
                            v21 := GenS1{F0: (^(fi16))}
                            if false {
                                fmt.Println(v3, fi16, v7)
                                switch "" {
                                case "":
                                    v7 = (3.14159 * 1e10)
                                default:
                                    for fi22 := 0; fi22 < 5; fi22++ {
                                        v2 = (fi22 << ((657) & 63))
                                        v2 = fi22
                                        v4 = append(v4, "\n\t")
                                        _ = fi22
                                    }
                                }
                            }
                            _ = v21
                        }
                        if true {
                            var v23 string = "a"
                            switch (fi16 | v2) {
                            case 9223372036854775807:
                                for fi24 := 0; fi24 < 4; fi24++ {
                                    switch v23 {
                                    case "hello":
                                        if (v1 >= v7) {
                                            v7 = (v1 - ((v7 - v1) * (3.14159 - v7)))
                                            if false {
                                                if (255 == 18446744073709551615) {
                                                    v2 = (fi16 % ((4294967295) | 1))
                                                    v3 -= v3
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                } else {
                                                    v23 = ""
                                                    if (false && (!(!true))) {
                                                        switch v23 {
                                                        case "\n\t":
                                                            v1 += float32(3.14159)
                                                        case "a":
                                                            if true {
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                                switch fi16 {
                                                                case 256:
                                                                    if ((v7 * v7) != (((0.1 + 1.5) / v7) / 1.0)) {
                                                                        v3 += v3
                                                                        switch v23 {
                                                                        case "\n\t":
                                                                            var v25 string = v23
                                                                            _ = v25
                                                                        case "hello":
                                                                            var v26 bool = false
                                                                            _ = v26
                                                                        case "a":
                                                                            if true {
                                                                                v4 = append(v4, (("foo bar" + "") + v23))
                                                                                var v27 uint32 = 533
                                                                                _ = v27
                                                                            } else {
                                                                                v4 = append(v4, v23)
                                                                            }
                                                                        default:
                                                                            if true {
                                                                                if ((32767 >= 1) || (true && (553 == 0))) {
                                                                                    var v28 float32 = v1
                                                                                    if (1 >= 638) {
                                                                                        v29 := GenS2{F0: -1, F1: 4294967295, F2: 42}
                                                                                        v4 = append(v4, v23)
                                                                                        _ = v29
                                                                                    }
                                                                                    _ = v28
                                                                                } else {
                                                                                    v30 := GenS3{F0: 69}
                                                                                    _ = v30
                                                                                }
                                                                                switch v23 {
                                                                                case "\n\t":
                                                                                    fmt.Println(v23, v23, v7)
                                                                                case "":
                                                                                    if len(v4) > 0 {
                                                                                        v4[0] = (v23 + "a")
                                                                                    }
                                                                                case "foo bar":
                                                                                    var v31 int64 = -624
                                                                                    _ = v31
                                                                                default:
                                                                                    fmt.Println(fi24, fi24)
                                                                                }
                                                                                v4 = append(v4, v23)
                                                                            }
                                                                        }
                                                                    }
                                                                case 9223372036854775807:
                                                                    v23 = "foo bar"
                                                                case 32767:
                                                                    v1 = (v1 + v1)
                                                                }
                                                                _ = len(v4)
                                                            } else {
                                                                fmt.Println(v23, v3)
                                                                if (566 < 4294967295) {
                                                                    v2 += (v2 & v2)
                                                                    var v32 float32 = 0.0
                                                                    _ = v32
                                                                } else {
                                                                    var v33 int = (fi24 >> ((4294967295) & 63))
                                                                    switch (v33 << ((806) & 63)) {
                                                                    case 2147483647:
                                                                        var v34 int8 = 59
                                                                        _ = v34
                                                                    case 871:
                                                                        switch (v23 + "") {
                                                                        case "hello":
                                                                            v33 += fi24
                                                                        case "":
                                                                            switch -9223372036854775808 {
                                                                            case -516:
                                                                                v3 = (v3 * v3)
                                                                            case 943:
                                                                                if len(v4) >= 2 {
                                                                                    v4 = v4[1:len(v4)]
                                                                                }
                                                                            case -830:
                                                                                v33 -= 993
                                                                            }
                                                                        case "\n\t":
                                                                            v3 = v3
                                                                        }
                                                                    default:
                                                                        _ = len(v4)
                                                                    }
                                                                    if (-1 <= 1000) {
                                                                        v35 := GenS2{F0: 127, F1: 994, F2: 832}
                                                                        v33 = v33
                                                                        _ = v35
                                                                    } else {
                                                                        v36 := GenS4{F0: 673, F1: -1, F2: 1}
                                                                        switch 32768 {
                                                                        case 32768:
                                                                            var v37 int16 = int16(v3)
                                                                            _ = v37
                                                                        }
                                                                        _ = v36
                                                                    }
                                                                    var v38 bool = ((v3 << ((18446744073709551615) & 63)) != v3)
                                                                    v3 += 42
                                                                    _ = v33
                                                                    _ = v38
                                                                }
                                                            }
                                                        }
                                                        switch v23 {
                                                        case "foo bar":
                                                            if len(v4) > 0 {
                                                                v4[0] = v23
                                                            }
                                                        case "":
                                                            v1 = ((((v1 / (1.5 * ((float32(3.14159) / float32(1e10)) + (-2.25 * v1)))) / ((v7 * ((v7 * v7) + v7)) + ((float32(3.14159) + float32(-2.25)) / (((v7 / 1.5) / float32(((float64(1.5) - float64(((v7 - 0.1) * 1.5))) - 0.0))) / v1)))) + (float32((1e10 - (3.14159 * 1e10))) + (v1 * v7))) + v1)
                                                        default:
                                                            v39 := GenS2{F0: -128, F1: 603, F2: 0}
                                                            _ = v39
                                                        }
                                                    }
                                                    v40 := GenS5{F0: (v3 ^ v3)}
                                                    _ = v40
                                                }
                                                v41 := GenS6{F0: 42, F1: fi24, F2: 42}
                                                _ = v41
                                            } else {
                                                switch v2 {
                                                case 42:
                                                    var v42 uint32 = 2147483648
                                                    _ = v42
                                                case 351:
                                                    v7 += v7
                                                case -1:
                                                    v4 = append(v4, "hello")
                                                    var v43 int16 = 1000
                                                    var v44 uint = 1
                                                    _ = v43
                                                    _ = v44
                                                default:
                                                    var v45 float64 = (float64(v1) * -2.25)
                                                    _ = v45
                                                }
                                            }
                                        }
                                    case "\n\t":
                                        if (v7 == float32(0.1)) {
                                            if len(v4) > 0 {
                                                v4[0] = ((v23 + ("\n\t" + (v23 + ("hello" + "")))) + (v23 + (v23 + "hello")))
                                            }
                                            _ = len(v4)
                                        }
                                    case "a":
                                        if true {
                                            if true {
                                                fmt.Println(fi24, fi24, v23)
                                                v4 = append(v4, "hello")
                                            } else {
                                                switch (fi24 | (fi24 | (fi16 % (((fi24 % (((-(fi16))) | 1))) | 1)))) {
                                                case -639:
                                                    if false {
                                                        var v46 string = "a"
                                                        fmt.Println(v46)
                                                        _ = v46
                                                    }
                                                case 128:
                                                    switch (((v23 + v23) + "hello") + "hello") {
                                                    case "\n\t":
                                                        v47 := GenS7{F0: 65535, F1: -538}
                                                        _ = v47
                                                    case "hello":
                                                        var v48 uint8 = 1
                                                        _ = v48
                                                    }
                                                default:
                                                    if true {
                                                        v49 := GenS6{F0: 42, F1: 2147483647, F2: 4294967295}
                                                        _ = v49
                                                    } else {
                                                        fmt.Println(v3, fi16)
                                                    }
                                                }
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                v50 := GenS8{F0: v23, F1: 698, F2: 32767}
                                                _ = v50
                                            }
                                            fmt.Println(v1)
                                        } else {
                                            var v51 uint8 = 1
                                            v52 := GenS6{F0: v3, F1: v2, F2: 1}
                                            v1 = (0.0 - (v1 / 0.1))
                                            _ = v51
                                            _ = v52
                                        }
                                    default:
                                        switch v2 {
                                        case 9223372036854775807:
                                            var v53 uint16 = 32767
                                            v54 := GenS5{F0: int32(v53)}
                                            _ = v53
                                            _ = v54
                                        case 32767:
                                            if true {
                                                switch fi16 {
                                                case 1:
                                                    fmt.Println(fi24, v7)
                                                case 0:
                                                    if false {
                                                        var v55 int16 = 860
                                                        if true {
                                                            if len(v4) > 0 {
                                                                v4[0] = v23
                                                            }
                                                            fmt.Println(v23)
                                                            if (false && (1 > 4294967295)) {
                                                                if ((418 < 4294967295) && true) {
                                                                    var v56 uint16 = 1
                                                                    switch 128 {
                                                                    case -2147483648:
                                                                        fmt.Println(fi24, v2)
                                                                    case 32768:
                                                                        if len(v4) > 0 {
                                                                            v4[0] = v23
                                                                        }
                                                                        v4 = append(v4, "\n\t")
                                                                    case 1:
                                                                        var v57 uint32 = 0
                                                                        _ = v57
                                                                    default:
                                                                        _ = len(v4)
                                                                    }
                                                                    _ = v56
                                                                }
                                                                var v58 uint32 = 2147483648
                                                                var v59 uint64 = 0
                                                                _ = v58
                                                                _ = v59
                                                            } else {
                                                                v55 -= -292
                                                                var v60 uint64 = 264
                                                                v61 := GenS4{F0: 18446744073709551615, F1: 42, F2: 18446744073709551615}
                                                                _ = v60
                                                                _ = v61
                                                            }
                                                        } else {
                                                            var v62 int8 = -1
                                                            switch "foo bar" {
                                                            case "hello":
                                                                if false {
                                                                    v3 = v3
                                                                } else {
                                                                    if (42 == 255) {
                                                                        var v63 string = ""
                                                                        _ = v63
                                                                    }
                                                                }
                                                            }
                                                            _ = v62
                                                        }
                                                        switch ("" + v23) {
                                                        case "hello":
                                                            v7 += v7
                                                        case "foo bar":
                                                            fmt.Println(v55, fi16)
                                                        }
                                                        _ = v55
                                                    }
                                                case 32767:
                                                    switch v2 {
                                                    case -281:
                                                        v7 = (0.1 + (((3.14159 * v7) / (((1.5 / (float32((0.1 * (0.0 / 1.5))) + (1.5 - v1))) - (float32(((1e10 - 3.14159) + (1.5 - (1.5 * 1e10)))) * v7)) * v1)) * (v1 - v7)))
                                                    default:
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                    }
                                                }
                                                fmt.Println(fi16, fi16, v3)
                                                var v64 int = v2
                                                _ = v64
                                            } else {
                                                var v65 float64 = -2.25
                                                var v66 int = (fi16 + fi24)
                                                _ = v65
                                                _ = v66
                                            }
                                        case 238:
                                            switch (v2 % ((-2147483648) | 1)) {
                                            case 127:
                                                v23 = "foo bar"
                                                var v67 int64 = 1
                                                v1 += v1
                                                _ = v67
                                            }
                                            var v68 uint = 18446744073709551615
                                            v69 := GenS7{F0: 255, F1: 0}
                                            v2 -= (fi24 >> ((v68) & 63))
                                            switch (v23 + ((v23 + "\n\t") + v23)) {
                                            case "hello":
                                                var v70 int16 = 606
                                                _ = v70
                                            case "a":
                                                var v71 uint64 = 340
                                                _ = v71
                                            case "":
                                                v72 := GenS5{F0: v3}
                                                _ = v72
                                            default:
                                                if ((275 > v68) && true) {
                                                    fmt.Println(v2)
                                                    switch v2 {
                                                    case 2147483647:
                                                        if (v3 <= (v3 / ((42) | 1))) {
                                                            v4 = append(v4, v23)
                                                        }
                                                    case -607:
                                                        if (!(((1 == 2147483647) || ("a" == "\n\t")) && (100 < 42))) {
                                                            if (float64(3.14159) < 3.14159) {
                                                                fmt.Println(v68, fi24)
                                                            } else {
                                                                v23 = "\n\t"
                                                                if ((false || (fi16 == fi24)) && (false && (true || (!true)))) {
                                                                    v4 = append(v4, ("\n\t" + v23))
                                                                    fmt.Println(v1, fi16)
                                                                    var v73 float32 = v1
                                                                    _ = v73
                                                                }
                                                            }
                                                            v1 -= (v7 / v7)
                                                            if false {
                                                                v1 = v1
                                                                v74 := GenS4{F0: v68, F1: v3, F2: 645}
                                                                _ = v74
                                                            } else {
                                                                v1 = (-2.25 + 0.0)
                                                                if ((-1 != 1) && (!false)) {
                                                                    fmt.Println(v23, v7, v2)
                                                                } else {
                                                                    if (v23 < (v23 + v23)) {
                                                                        var v75 bool = (-496 == -9223372036854775808)
                                                                        v76 := GenS3{F0: 0}
                                                                        v7 -= 0.1
                                                                        _ = v75
                                                                        _ = v76
                                                                    }
                                                                    v77 := GenS9{F0: genFn0(), F1: v3}
                                                                    _ = v77
                                                                }
                                                                var v78 bool = true
                                                                _ = v78
                                                            }
                                                        } else {
                                                            var v79 int8 = 100
                                                            fmt.Println(v7, v23)
                                                            v80 := GenS2{F0: (-(v79)), F1: 602, F2: 0}
                                                            _ = v79
                                                            _ = v80
                                                        }
                                                    case -150:
                                                        var v81 uint8 = 128
                                                        _ = v81
                                                    default:
                                                        if true {
                                                            fmt.Println(v7, v3, fi16)
                                                        } else {
                                                            v82 := GenS3{F0: -132}
                                                            _ = v82
                                                        }
                                                        v3 -= 2147483647
                                                        var v83 uint = 4294967295
                                                        var v84 int32 = v3
                                                        v3 = v84
                                                        _ = v83
                                                        _ = v84
                                                    }
                                                }
                                            }
                                            var v85 float32 = (float32(float64(float32(((-2.25 * (1.5 / 0.1)) + 1e10)))) / v1)
                                            _ = v68
                                            _ = v69
                                            _ = v85
                                        default:
                                            _ = len(v4)
                                        }
                                    }
                                    v7 = 1e10
                                    _ = fi24
                                }
                            case 2147483647:
                                fmt.Println(v3, v1)
                            default:
                                for fi86 := 0; fi86 < 11; fi86++ {
                                    v7 = (v7 + v1)
                                    v87 := GenS5{F0: (v3 * 608)}
                                    _ = fi86
                                    _ = v87
                                }
                            }
                            fmt.Println(v2)
                            _ = v23
                        }
                        var v88 int16 = 760
                        _ = fi16
                        _ = v88
                    }
                }
            } else {
                if len(v4) > 0 {
                    v4[0] = "foo bar"
                }
                var v89 uint64 = 1
                _ = v89
            }
        default:
            if false {
                fmt.Println(v2)
                var v90 uint32 = 2147483648
                v91 := GenS6{F0: v3, F1: v2, F2: 648}
                _ = v90
                _ = v91
            }
        }
    }
    v92 := GenS10{F0: 0, F1: v3}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v93 float64 = 1.5
    if (66 == 1) {
        v3 = (v3 / (((v3 >> ((814) & 63))) | 1))
    }
    var v94 uint = 0
    var v95 int16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v93:", v93)
    fmt.Println("v94:", v94)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v92
    _ = v93
    _ = v94
    _ = v95
}
