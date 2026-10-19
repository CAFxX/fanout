package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16; F1 uint8; F2 uint32 }
type GenS5 struct { F0 int8; F1 bool; F2 uint16 }
type GenS6 struct { F0 int32; F1 uint }
func genFn0(p0 uint8, p1 string, p2 int64) int8 {
    v6 := GenS1{F0: 383}
    if (false && false) {
        p2 = -1
        for fi7 := 0; fi7 < 6; fi7++ {
            fmt.Println(p2)
            var v8 int = (fi7 | fi7)
            var v9 string = ""
            _ = fi7
            _ = v8
            _ = v9
        }
        var v10 float32 = 1e10
        _ = v10
    } else {
        p2 -= (p2 & (p2 - p2))
        v11 := []int{}
        _ = v11
    }
    if false {
        var v12 int64 = 952
        for fi13 := 0; fi13 < 1; fi13++ {
            v12 += v12
            p2 = v12
            v12 = p2
            _ = fi13
        }
        _ = v12
    } else {
        if (false && (1 == 143)) {
            p2 -= 722
            for fi14 := 0; fi14 < 3; fi14++ {
                var v15 int8 = -1
                if (1 <= 878) {
                    v16 := []int{776, 32767, 828}
                    var v17 int32 = 0
                    fmt.Println(p2, p0, fi14)
                    _ = v16
                    _ = v17
                }
                _ = fi14
                _ = v15
            }
        } else {
            v18 := GenS1{F0: -1}
            v19 := GenS2{F0: p0}
            _ = v18
            _ = v19
        }
        switch p1 {
        case "foo bar":
            for fi20 := 0; fi20 < 5; fi20++ {
                fmt.Println(fi20, p2, fi20)
                _ = fi20
            }
        }
        switch (("\n\t" + "") + (p1 + p1)) {
        case "":
            v21 := GenS1{F0: int16(p0)}
            _ = v21
        default:
            switch -1 {
            case -9223372036854775808:
                if ((!true) && true) {
                    p1 = (p1 + "foo bar")
                    for fi22 := 0; fi22 < 18; fi22++ {
                        v23 := []int{(fi22 + (fi22 << ((1) & 63))), fi22, (-(fi22)), 127}
                        _ = fi22
                        _ = v23
                    }
                } else {
                    var v24 float32 = (0.0 + 3.14159)
                    var v25 int64 = (p2 % ((p2) | 1))
                    p1 = (p1 + (p1 + p1))
                    _ = v24
                    _ = v25
                }
            case -991:
                var v26 string = p1
                _ = v26
            case 32768:
                p1 = p1
            default:
                p0 += p0
            }
        }
    }
    if true {
        v27 := GenS2{F0: p0}
        for fi28 := 0; fi28 < 5; fi28++ {
            v29 := GenS1{F0: -587}
            for fi30 := 0; fi30 < 17; fi30++ {
                var v31 uint = 255
                p2 += -1
                p1 = p1
                _ = fi30
                _ = v31
            }
            p0 += p0
            _ = fi28
            _ = v29
        }
        if (2 < (p0 * (p0 >> ((403) & 63)))) {
            var v32 int8 = -128
            _ = v32
        } else {
            p2 -= p2
            var v33 uint = 4294967295
            _ = v33
        }
        _ = v27
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 42
}
func genFn1(p0 uint16) bool {
    switch 42 {
    case 32768:
        for fi34 := 0; fi34 < 15; fi34++ {
            var v35 uint = 42
            switch "" {
            case "foo bar":
                fmt.Println(v35, fi34, p0)
            case "\n\t":
                for fi36 := 0; fi36 < 15; fi36++ {
                    if (!true) {
                        fmt.Println(p0, v35)
                    }
                    v37 := []int{(fi34 & fi36), 643, (fi36 ^ fi36)}
                    var v38 bool = true
                    _ = fi36
                    _ = v37
                    _ = v38
                }
            default:
                v35 += (v35 - 4294967295)
            }
            _ = fi34
            _ = v35
        }
    default:
        p0 += (p0 + p0)
    }
    var v39 uint64 = 479
    switch 189 {
    case -152:
        switch "hello" {
        case "":
            fmt.Println(p0, v39)
        default:
            for fi40 := 0; fi40 < 18; fi40++ {
                var v41 int32 = -1
                v41 -= v41
                _ = fi40
                _ = v41
            }
        }
    case -9223372036854775808:
        var v42 uint8 = 110
        _ = v42
    case -858:
        v43 := []int{-792, 2147483647}
        _ = v43
    default:
        switch -2147483648 {
        case 966:
            if (p0 < (p0 >> ((uint(p0)) & 63))) {
                v44 := GenS1{F0: -1}
                switch 4294967295 {
                case -2147483648:
                    p0 = 176
                default:
                    v45 := []int{}
                    _ = v45
                }
                _ = v44
            }
        case 32768:
            var v46 uint32 = 2147483647
            _ = v46
        case 2147483647:
            v47 := []int{127, 211, -272}
            _ = v47
        default:
            if false {
                v39 -= v39
                if (-128 == genFn0(49, "\n\t", -9223372036854775808)) {
                    v48 := []int{}
                    p0 -= p0
                    _ = len(v48)
                    _ = v48
                }
            }
        }
    }
    v49 := []int{}
    p0 += p0
    _ = p0
    _ = v39
    _ = v49
    return true
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int64 = -1
    var v3 float32 = -2.25
    var v4 int16 = 533
    v5 := []uint32{2, 4294967295, 671, 747}
    var v50 bool = false
    switch -1 {
    case 1:
        v50 = v50
    }
    v51 := GenS2{F0: 1}
    v4 -= 1000
    var v52 uint64 = 2
    v53 := v1[42]
    v54 := GenS2{F0: 42}
    v2 = v2
    v55 := GenS2{F0: 0}
    switch "" {
    case "":
        v53 = (!(v2 >= (-(v2))))
    case "a":
        if (-829 != int(v52)) {
            var v56 bool = ((v2 ^ -9223372036854775808) != v2)
            v57 := GenS3{F0: 997}
            for fi58 := 0; fi58 < 15; fi58++ {
                var v59 string = ("foo bar" + "")
                for fi60 := 0; fi60 < 11; fi60++ {
                    var v61 uint32 = 473
                    switch "foo bar" {
                    case "foo bar":
                        switch (fi58 & (^(fi60))) {
                        case 128:
                            switch v59 {
                            case "":
                                if (-810 >= (v2 << ((65535) & 63))) {
                                    v62 := GenS4{F0: 472, F1: 127, F2: v61}
                                    _ = v62
                                }
                            case "hello":
                                if v50 {
                                    var v63 uint = 42
                                    var v64 bool = false
                                    v65 := GenS5{F0: 100, F1: v56, F2: 0}
                                    _ = v63
                                    _ = v64
                                    _ = v65
                                } else {
                                    if ((true && (!v50)) && (!false)) {
                                        if (32 >= -1) {
                                            switch -2147483648 {
                                            case 32768:
                                                switch v59 {
                                                case "foo bar":
                                                    var v66 int32 = int32(fi58)
                                                    _ = v66
                                                case "a":
                                                    var v67 int32 = 0
                                                    v3 -= v3
                                                    delete(v1, (v61 * v61))
                                                    _ = v67
                                                default:
                                                    v68 := v1[2147483648]
                                                    _ = v68
                                                }
                                            case 1:
                                                delete(v1, v61)
                                            case 9223372036854775807:
                                                if len(v5) > 0 {
                                                    v5[0] = (^(v61))
                                                }
                                            }
                                            v2 = v2
                                        }
                                        fmt.Println(fi58, v52)
                                        v69 := GenS6{F0: -1, F1: 18446744073709551615}
                                        var v70 string = ((("\n\t" + "hello") + ("\n\t" + v59)) + "foo bar")
                                        _ = v69
                                        _ = v70
                                    }
                                    v5 = append(v5, (^(v61)))
                                }
                            case "a":
                                v71 := GenS5{F0: 42, F1: (653 < 2), F2: 341}
                                _ = v71
                            }
                        case 32768:
                            v56 = (!(!true))
                        case 32767:
                            var v72 float32 = (0.1 - v3)
                            _ = v72
                        default:
                            v52 = v52
                        }
                    default:
                        fmt.Println(v56, v3, v59)
                    }
                    switch fi60 {
                    case 127:
                        var v73 uint8 = 127
                        _ = v73
                    case -1:
                        var v74 float64 = 1e10
                        _ = v74
                    case 4294967295:
                        switch (fi58 | fi58) {
                        case -448:
                            v50 = (0 != -2147483648)
                        case 1:
                            v3 = -2.25
                        case 0:
                            _ = len(v1)
                        }
                    default:
                        if true {
                            v52 = (v52 << ((0) & 63))
                            v75 := GenS1{F0: v4}
                            _ = v75
                        }
                    }
                    if true {
                        v59 = (v59 + "hello")
                        var v76 uint = 826
                        v52 += v52
                        _ = v76
                    } else {
                        v77 := v1[(v61 + 2147483648)]
                        _ = v77
                    }
                    v4 += (v4 & (v4 >> ((39) & 63)))
                    _ = fi60
                    _ = v61
                }
                _ = fi58
                _ = v59
            }
            _ = v56
            _ = v57
        } else {
            if v53 {
                v50 = (601 == 0)
                var v78 uint64 = v52
                _ = v78
            }
        }
    case "hello":
        var v79 uint = 2
        _ = v79
    default:
        var v80 int = -835
        _ = v80
    }
    v4 += -32768
    v2 = (v2 + -1)
    switch 0 {
    case 127:
        switch (("foo bar" + "") + ("\n\t" + "a")) {
        case "a":
            switch 475 {
            case 4294967295:
                var v81 float32 = (v3 - float32(3.14159))
                _ = v81
            case 222:
                if (0 <= 482) {
                    if (v50 || false) {
                        v53 = true
                        for fi82 := 0; fi82 < 18; fi82++ {
                            v83 := GenS3{F0: 65535}
                            var v84 float32 = v3
                            for fi85 := 0; fi85 < 13; fi85++ {
                                var v86 string = ("foo bar" + "\n\t")
                                v87 := GenS5{F0: -1, F1: v53, F2: 65535}
                                var v88 int8 = 127
                                _ = fi85
                                _ = v86
                                _ = v87
                                _ = v88
                            }
                            _ = fi82
                            _ = v83
                            _ = v84
                        }
                    }
                    for fi89 := 0; fi89 < 13; fi89++ {
                        v2 -= v2
                        _ = fi89
                    }
                }
            case 32767:
                var v90 float64 = 1.5
                _ = v90
            default:
                v91 := GenS3{F0: 0}
                _ = v91
            }
        case "":
            var v92 string = ("hello" + "foo bar")
            _ = v92
        }
    case -2147483648:
        if ((true && (((-95 <= -128) || (13 < 65535)) && v50)) && (true || ((-1 >= (v2 ^ v2)) && (true || (true || ((3.14159 - 1e10) > (0.1 - (1e10 / 1.0)))))))) {
            v53 = (v2 > (v2 >> ((768) & 63)))
            v5 = append(v5, 42)
        }
    case 690:
        _ = len(v5)
    default:
        if true {
            var v93 string = "\n\t"
            var v94 int8 = -1
            _ = v93
            _ = v94
        }
    }
    _ = len(v5)
    fmt.Println(v53, v3, v3)
    if true {
        v1[985] = v53
    }
    fmt.Println(v2, v53)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v55
}
