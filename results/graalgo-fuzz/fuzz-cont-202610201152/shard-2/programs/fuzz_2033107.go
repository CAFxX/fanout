package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 int8; F2 uint16 }
type GenS4 struct { F0 uint8; F1 uint64; F2 uint16 }
type GenS5 struct { F0 string; F1 uint8 }
type GenS6 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint16; F1 uint32; F2 int16 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int32 }
func genFn0(p0 uint) string {
    fmt.Println(p0)
    for fi4 := 0; fi4 < 2; fi4++ {
        p0 -= p0
        if (p0 >= p0) {
            v5 := []int{fi4, (fi4 ^ (-(fi4))), fi4, (fi4 << (((p0 + p0)) & 63))}
            _ = v5
        } else {
            p0 = 1
            var v6 uint8 = 49
            _ = v6
        }
        _ = fi4
    }
    for fi7 := 0; fi7 < 10; fi7++ {
        switch "a" {
        case "":
            for fi8 := 0; fi8 < 4; fi8++ {
                p0 = p0
                _ = fi8
            }
        case "\n\t":
            fmt.Println(p0)
        default:
            if true {
                p0 += (p0 & 2)
                p0 = (^(p0))
            } else {
                var v9 uint = (p0 + (p0 << ((321) & 63)))
                v9 = v9
                _ = v9
            }
        }
        var v10 float64 = (3.14159 - float64((float32((1e10 / 3.14159)) + (1.5 * 1e10))))
        if false {
            var v11 uint = (^(p0))
            for fi12 := 0; fi12 < 12; fi12++ {
                fmt.Println(fi7)
                _ = fi12
            }
            _ = v11
        } else {
            if ((127 > 100) && false) {
                if false {
                    var v13 uint8 = 207
                    switch -479 {
                    case 32767:
                        v14 := []int{(fi7 % (((fi7 * fi7)) | 1)), fi7, (fi7 >> ((p0) & 63))}
                        _ = v14
                    case 128:
                        if ((892 == (^(p0))) && (42 <= 18446744073709551615)) {
                            for fi15 := 0; fi15 < 15; fi15++ {
                                var v16 int8 = -1
                                var v17 uint16 = 42
                                var v18 int = fi7
                                _ = fi15
                                _ = v16
                                _ = v17
                                _ = v18
                            }
                            fmt.Println(p0, v10)
                            fmt.Println(fi7, fi7)
                        } else {
                            for fi19 := 0; fi19 < 13; fi19++ {
                                if (!false) {
                                    p0 = p0
                                } else {
                                    v13 = v13
                                    p0 = (p0 / (((^(p0))) | 1))
                                    if true {
                                        fmt.Println(v13)
                                        var v20 uint64 = 1
                                        _ = v20
                                    }
                                    v21 := GenS1{F0: 648, F1: 39, F2: -1}
                                    var v22 float32 = 1e10
                                    _ = v21
                                    _ = v22
                                }
                                if false {
                                    fmt.Println(p0, fi7, p0)
                                }
                                v23 := []int{fi19, fi7, (fi7 & (fi7 + (fi7 - 32767))), fi19}
                                fmt.Println(v13, v13)
                                _ = fi19
                                _ = v23
                            }
                            v24 := []int{(fi7 << (((p0 << (((p0 + (p0 >> (((p0 % ((p0) | 1))) & 63)))) & 63))) & 63))}
                            var v25 int = fi7
                            _ = v24
                            _ = v25
                        }
                    }
                    for fi26 := 0; fi26 < 16; fi26++ {
                        var v27 uint16 = 1
                        switch fi26 {
                        case 0:
                            v13 = v13
                            v28 := GenS2{F0: 32767}
                            var v29 int16 = -32768
                            _ = v28
                            _ = v29
                        default:
                            v30 := []int{fi7, int(v27)}
                            _ = v30
                        }
                        switch fi7 {
                        case -1:
                            if true {
                                if true {
                                    if (p0 < (p0 ^ (p0 << ((862) & 63)))) {
                                        if (!((0 >= 42) || false)) {
                                            v13 += (^(v13))
                                        }
                                        if true {
                                            v31 := GenS3{F0: ("foo bar" + "\n\t"), F1: 127, F2: v27}
                                            _ = v31
                                        }
                                        var v32 string = "hello"
                                        _ = v32
                                    } else {
                                        var v33 int8 = 111
                                        var v34 uint64 = 334
                                        var v35 uint32 = 594
                                        _ = v33
                                        _ = v34
                                        _ = v35
                                    }
                                    var v36 uint8 = (v13 / (((v13 << ((p0) & 63))) | 1))
                                    v27 -= v27
                                    v37 := []int{fi26, fi7, -2147483648}
                                    _ = v36
                                    _ = v37
                                }
                                if (582 < 281) {
                                    var v38 float64 = 1.5
                                    _ = v38
                                }
                                v27 += 255
                            }
                        case -936:
                            if ((42 < -940) || (true || ((2147483647 > 306) || (!true)))) {
                                p0 = (p0 % (((p0 + 2)) | 1))
                                v39 := GenS4{F0: (v13 / ((42) | 1)), F1: 286, F2: (v27 | 3)}
                                if false {
                                    var v40 uint16 = (^(v27))
                                    v41 := GenS4{F0: (v13 << ((255) & 63)), F1: 42, F2: (v40 >> (((p0 / ((p0) | 1))) & 63))}
                                    _ = v40
                                    _ = v41
                                }
                                _ = v39
                            } else {
                                v42 := GenS5{F0: "a", F1: 2}
                                v43 := []int{fi26, (fi7 % ((fi7) | 1)), (fi7 ^ (-(fi7))), (fi7 | (fi26 % (((fi7 ^ (fi7 >> ((p0) & 63)))) | 1)))}
                                _ = v42
                                _ = v43
                            }
                        case 128:
                            p0 -= p0
                        }
                        _ = fi26
                        _ = v27
                    }
                    _ = v13
                }
                var v44 uint16 = 32767
                if false {
                    v44 += 42
                    for fi45 := 0; fi45 < 7; fi45++ {
                        var v46 uint32 = 1
                        var v47 uint64 = 4294967295
                        _ = fi45
                        _ = v46
                        _ = v47
                    }
                    for fi48 := 0; fi48 < 6; fi48++ {
                        var v49 int8 = 0
                        _ = fi48
                        _ = v49
                    }
                }
                _ = v44
            }
            for fi50 := 0; fi50 < 14; fi50++ {
                v51 := []int{(fi7 + (fi50 / (((fi50 + fi50)) | 1))), (fi50 >> (((p0 & p0)) & 63)), (fi7 / ((32767) | 1)), fi7}
                _ = fi50
                _ = v51
            }
            var v52 float32 = float32(1.5)
            _ = v52
        }
        _ = fi7
        _ = v10
    }
    var v53 int64 = -89
    _ = p0
    _ = v53
    return ("\n\t" + ("hello" + ""))
}

func main() {
    var v1 int = 128
    v2 := []string{}
    var v3 string = ""
    var v54 uint16 = 2
    var v55 uint64 = 1
    switch v1 {
    case -9223372036854775808:
        if false {
            if true {
                if (0.1 != -2.25) {
                    fmt.Println(v55, v54)
                    v56 := GenS6{F0: 407, F1: 2, F2: v55}
                    fmt.Println(v3, v55, v54)
                    _ = v56
                } else {
                    for fi57 := 0; fi57 < 10; fi57++ {
                        var v58 int = (v1 * fi57)
                        if false {
                            _ = len(v2)
                        }
                        for fi59 := 0; fi59 < 10; fi59++ {
                            if (!(fi59 < fi57)) {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                fmt.Println(v58, fi57, v55)
                                if (!(1e10 > (3.14159 + 1e10))) {
                                    v55 -= 202
                                    switch "foo bar" {
                                    case "hello":
                                        v55 += (v55 ^ (v55 | 4294967295))
                                    case "a":
                                        v60 := GenS7{F0: 103}
                                        _ = v60
                                    default:
                                        v3 = (v3 + (((v3 + v3) + v3) + v3))
                                    }
                                }
                                var v61 int16 = -1
                                _ = v61
                            }
                            if true {
                                fmt.Println(v3)
                                v2 = append(v2, v3)
                                v55 = 42
                            } else {
                                var v62 int = (v58 | int(v55))
                                v63 := GenS7{F0: 633}
                                v2 = append(v2, "")
                                _ = v62
                                _ = v63
                            }
                            v3 = v3
                            _ = fi59
                        }
                        _ = fi57
                        _ = v58
                    }
                    for fi64 := 0; fi64 < 12; fi64++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v65 uint8 = 127
                        v55 = (^(v55))
                        _ = fi64
                        _ = v65
                    }
                }
            } else {
                v66 := GenS8{F0: (v54 | 639), F1: uint32(v54), F2: int16(v54)}
                _ = v66
            }
        }
    case 9223372036854775807:
        var v67 uint16 = 32767
        _ = v67
    case -556:
        switch v1 {
        case 539:
            if true {
                switch v3 {
                case "a":
                    switch -2147483648 {
                    case -2147483648:
                        v68 := GenS9{F0: 79}
                        _ = v68
                    case 32768:
                        fmt.Println(v3)
                    }
                default:
                    if false {
                        fmt.Println(v55, v54)
                        v2 = append(v2, "a")
                        if true {
                            v69 := GenS8{F0: (v54 - v54), F1: 42, F2: -1}
                            _ = v69
                        }
                    } else {
                        var v70 uint16 = (v54 | v54)
                        if (false && (!false)) {
                            if (967 >= 1) {
                                if (-2.25 < 0.0) {
                                    var v71 int64 = -9223372036854775808
                                    if len(v2) > 0 {
                                        v2[0] = v3
                                    }
                                    _ = v71
                                }
                                var v72 int = (v1 | -816)
                                _ = v72
                            }
                            v3 = ((v3 + v3) + "foo bar")
                        }
                        _ = v70
                    }
                }
            }
        case 2147483647:
            fmt.Println(v55, v3, v1)
        default:
            v54 = (v54 << ((1) & 63))
        }
    default:
        v3 = (("hello" + v3) + v3)
    }
    var v73 int8 = 100
    for fi74 := 0; fi74 < 4; fi74++ {
        v2 = append(v2, (v3 + v3))
        _ = fi74
    }
    if true {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v73 += v73
        switch 127 {
        case 9223372036854775807:
            v75 := GenS10{F0: 244}
            _ = v75
        case 309:
            _ = len(v2)
        case 665:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    } else {
        switch "hello" {
        case "\n\t":
            switch v3 {
            case "\n\t":
                var v76 uint16 = 255
                _ = v76
            case "a":
                var v77 int32 = -285
                _ = v77
            }
        case "":
            var v78 uint16 = (v54 | v54)
            _ = v78
        case "foo bar":
            v73 = (v73 + 127)
        default:
            v73 = (^(v73))
        }
        v3 = ("" + (((v3 + (v3 + "\n\t")) + (("a" + "foo bar") + ("a" + v3))) + v3))
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    if (0 > 1) {
        v79 := GenS7{F0: 412}
        v80 := GenS8{F0: 65535, F1: 42, F2: 1}
        v2 = append(v2, ((((("hello" + ("hello" + "\n\t")) + v3) + "\n\t") + (("foo bar" + ("hello" + v3)) + v3)) + v3))
        _ = v79
        _ = v80
    } else {
        switch 0 {
        case 32767:
            v3 = v3
        case 184:
            v81 := GenS6{F0: 40, F1: 2, F2: 8}
            _ = v81
        default:
            var v82 float64 = (0.1 / -2.25)
            _ = v82
        }
        switch v3 {
        case "":
            v73 -= 42
        case "a":
            var v83 int16 = 32767
            _ = v83
        }
        var v84 float32 = 0.1
        _ = v84
    }
    v55 += 42
    for fi85 := 0; fi85 < 4; fi85++ {
        if (((0.0 / 1.0) <= -2.25) || true) {
            for fi86 := 0; fi86 < 14; fi86++ {
                if ((v73 & 127) != 127) {
                    if ("foo bar" != v3) {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = len(v2)
                        v3 = ""
                    }
                    fmt.Println(v1, v3, v1)
                } else {
                    var v87 uint32 = 0
                    var v88 string = v3
                    _ = v87
                    _ = v88
                }
                v89 := GenS6{F0: 0, F1: 0, F2: v55}
                v73 -= (v73 - v73)
                _ = fi86
                _ = v89
            }
            v2 = append(v2, (v3 + v3))
        } else {
            v2 = append(v2, v3)
        }
        _ = fi85
    }
    v3 = v3
    if false {
        if true {
            var v90 bool = false
            _ = v90
        } else {
            v55 = 855
        }
    } else {
        v73 = 39
        fmt.Println(v54)
    }
    if (!false) {
        fmt.Println(v3, v3, v54)
        if true {
            v3 = (v3 + v3)
        }
        for fi91 := 0; fi91 < 6; fi91++ {
            v1 = (v1 << ((451) & 63))
            v92 := GenS2{F0: -1}
            _ = fi91
            _ = v92
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v54
    _ = v55
    _ = v73
}
