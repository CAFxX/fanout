package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 uint; F1 int32; F2 int32 }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 bool; F1 uint8 }
type GenS6 struct { F0 uint; F1 int8; F2 int16 }
type GenS7 struct { F0 int; F1 int64 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint; F1 int64; F2 bool }
type GenS10 struct { F0 int16; F1 uint32 }
func genFn0() int8 {
    for fi7 := 0; fi7 < 15; fi7++ {
        if (!true) {
            var v8 uint16 = 42
            _ = v8
        }
        var v9 int64 = -9223372036854775808
        var v10 int32 = -2147483648
        _ = fi7
        _ = v10
        _ = v9
    }
    var v11 uint32 = 532
    fmt.Println(v11)
    _ = v11
    return -4
}

func main() {
    var v1 uint16 = 2
    var v2 uint64 = 2
    var v3 string = ("foo bar" + "")
    var v4 int16 = -1
    var v5 float32 = (-2.25 / 3.14159)
    var v6 int8 = -25
    switch ("" + v3) {
    case "a":
        for fi12 := 0; fi12 < 8; fi12++ {
            v13 := GenS1{F0: 9223372036854775807, F1: 42}
            _ = fi12
            _ = v13
        }
    case "\n\t":
        switch int(v6) {
        case 42:
            v6 += v6
        default:
            switch v3 {
            case "\n\t":
                if false {
                    switch 918 {
                    case 32767:
                        if true {
                            v2 += v2
                            var v14 bool = false
                            _ = v14
                        }
                    default:
                        v4 += (-(v4))
                    }
                    var v15 int64 = 42
                    v16 := []int{326, 32767, 4294967295, 127}
                    _ = v15
                    _ = v16
                } else {
                    var v17 float64 = -2.25
                    for fi18 := 0; fi18 < 19; fi18++ {
                        v19 := GenS2{F0: 2, F1: -2147483648, F2: 2147483647}
                        v20 := GenS1{F0: 1, F1: 316}
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                    _ = v17
                }
            default:
                if true {
                    if true {
                        v3 = v3
                        v21 := []int{-167, 128}
                        _ = v21
                    }
                    var v22 int = 32767
                    _ = v22
                } else {
                    switch v3 {
                    case "\n\t":
                        var v23 int64 = -9223372036854775808
                        _ = v23
                    }
                }
            }
        }
    default:
        switch ("a" + v3) {
        case "a":
            for fi24 := 0; fi24 < 11; fi24++ {
                v25 := []int{(fi24 ^ fi24), fi24, (fi24 << ((18446744073709551615) & 63)), (fi24 >> ((42) & 63))}
                v1 = 42
                v3 = ""
                _ = fi24
                _ = v25
            }
        case "\n\t":
            if false {
                for fi26 := 0; fi26 < 11; fi26++ {
                    for fi27 := 0; fi27 < 4; fi27++ {
                        switch fi27 {
                        case -346:
                            var v28 uint64 = v2
                            _ = v28
                        case 32768:
                            fmt.Println(fi26, fi27, v2)
                        case -9223372036854775808:
                            v29 := []int{fi27, (fi27 % ((fi27) | 1)), (fi26 & (fi26 * (fi27 << ((605) & 63))))}
                            _ = v29
                        }
                        fmt.Println(v1)
                        _ = fi27
                    }
                    for fi30 := 0; fi30 < 7; fi30++ {
                        var v31 int16 = (^(v4))
                        var v32 uint32 = 574
                        v33 := GenS1{F0: 45, F1: 1}
                        switch v3 {
                        case "":
                            var v34 float32 = ((float32((((1.5 * 1.5) - (0.0 + (1e10 * 1e10))) / 1e10)) - 3.14159) * v5)
                            _ = v34
                        case "\n\t":
                            if true {
                                v32 -= (v32 + (v32 + uint32(v1)))
                            } else {
                                var v35 uint = 2
                                _ = v35
                            }
                        case "hello":
                            if ((v32 >= v32) && (!true)) {
                                var v36 int64 = 820
                                var v37 uint = 2
                                v38 := GenS3{F0: 522, F1: (v32 * v32)}
                                if ((v3 + "hello") <= (v3 + v3)) {
                                    v39 := GenS1{F0: v36, F1: (v37 - (^(v37)))}
                                    if (("\n\t" + (v3 + (v3 + v3))) == v3) {
                                        v6 += v6
                                    } else {
                                        v40 := []int{}
                                        v41 := GenS4{F0: 529}
                                        _ = v40
                                        _ = v41
                                    }
                                    _ = v39
                                }
                                _ = v36
                                _ = v37
                                _ = v38
                            }
                        default:
                            if ((127 == (v6 ^ v6)) || (false && (int(v32) > fi26))) {
                                v42 := GenS3{F0: v1, F1: (v32 >> ((42) & 63))}
                                _ = v42
                            }
                        }
                        _ = fi30
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                    _ = fi26
                }
                var v43 float64 = 1e10
                _ = v43
            } else {
                var v44 uint64 = 0
                _ = v44
            }
        case "foo bar":
            v5 = ((v5 * -2.25) * v5)
        }
    }
    v6 = genFn0()
    switch -9223372036854775808 {
    case -619:
        v45 := GenS5{F0: false, F1: 182}
        _ = v45
    case 128:
        fmt.Println(v5, v2, v3)
    default:
        v1 = (v1 << ((0) & 63))
    }
    if (!(true || (!(((127 == 2) && false) && (v6 == (v6 / ((0) | 1))))))) {
        v6 -= -56
        var v46 uint32 = uint32(v1)
        fmt.Println(v6, v46, v6)
        _ = v46
    } else {
        v47 := []int{1, 32768, -2147483648}
        _ = v47
    }
    var v48 int8 = genFn0()
    if true {
        fmt.Println(v4, v4)
        v49 := []int{-1, 127}
        _ = v49
    } else {
        v50 := []int{}
        _ = v50
    }
    var v51 uint = 65535
    v48 = v48
    switch "foo bar" {
    case "hello":
        if true {
            for fi52 := 0; fi52 < 6; fi52++ {
                v51 -= (v51 >> ((65535) & 63))
                for fi53 := 0; fi53 < 4; fi53++ {
                    v54 := []int{}
                    if true {
                        switch v3 {
                        case "a":
                            switch (fi53 << ((v51) & 63)) {
                            case -9223372036854775808:
                                var v55 uint8 = 51
                                _ = v55
                            case 9223372036854775807:
                                var v56 int32 = 0
                                if len(v54) >= 2 {
                                    v54 = v54[1:len(v54)]
                                }
                                _ = v56
                            case 32767:
                                _ = len(v54)
                            }
                        case "foo bar":
                            if true {
                                fmt.Println(v3)
                                var v57 bool = false
                                switch 128 {
                                case 9223372036854775807:
                                    switch v3 {
                                    case "a":
                                        v58 := GenS5{F0: v57, F1: uint8(v1)}
                                        _ = v58
                                    case "hello":
                                        v54 = append(v54, (fi53 % ((61) | 1)))
                                    }
                                case 32768:
                                    if (0 < int64(v1)) {
                                        v54 = append(v54, (fi52 >> ((v51) & 63)))
                                        switch fi53 {
                                        case 1:
                                            var v59 uint16 = v1
                                            _ = v59
                                        case 0:
                                            if v57 {
                                                switch 632 {
                                                case -556:
                                                    switch v3 {
                                                    case "hello":
                                                        v54 = append(v54, fi52)
                                                    case "foo bar":
                                                        if true {
                                                            var v60 int32 = 242
                                                            var v61 uint16 = (v1 - v1)
                                                            v54 = append(v54, (fi52 >> (((^(v51))) & 63)))
                                                            _ = v60
                                                            _ = v61
                                                        } else {
                                                            if v57 {
                                                                fmt.Println(v1)
                                                                var v62 uint = v51
                                                                v51 = (v51 | v62)
                                                                _ = v62
                                                            }
                                                            var v63 int8 = (-(v48))
                                                            var v64 bool = v57
                                                            _ = v63
                                                            _ = v64
                                                        }
                                                        v65 := GenS2{F0: 4294967295, F1: -837, F2: -411}
                                                        _ = v65
                                                    }
                                                case 9223372036854775807:
                                                    if true {
                                                        switch (-(fi52)) {
                                                        case -645:
                                                            v57 = v57
                                                        default:
                                                            v54 = append(v54, (^(fi53)))
                                                            v66 := GenS6{F0: (v51 << ((1) & 63)), F1: genFn0(), F2: v4}
                                                            var v67 int = fi52
                                                            _ = v66
                                                            _ = v67
                                                        }
                                                        v68 := GenS7{F0: int(v1), F1: -9223372036854775808}
                                                        v57 = v57
                                                        _ = v68
                                                    }
                                                case 4294967295:
                                                    var v69 float64 = 3.14159
                                                    _ = v69
                                                default:
                                                    v6 += (^(v48))
                                                }
                                                v70 := GenS6{F0: v51, F1: v48, F2: (-(v4))}
                                                v6 -= (-(v48))
                                                _ = v70
                                            } else {
                                                if true {
                                                    var v71 uint16 = uint16(v2)
                                                    v54 = append(v54, (^(fi53)))
                                                    if (((int8(fi53) == (v6 << (((^(v51))) & 63))) || (0 == 0)) || (v2 != (^(v2)))) {
                                                        fmt.Println(fi53)
                                                        v57 = (0 == 2147483648)
                                                        var v72 uint8 = 42
                                                        _ = v72
                                                    } else {
                                                        if (-2.25 <= 1.5) {
                                                            fmt.Println(v57)
                                                            v4 = (v4 % (((v4 + -171)) | 1))
                                                            fmt.Println(v1, v6)
                                                        } else {
                                                            v54 = append(v54, fi52)
                                                            v51 = (^(v51))
                                                        }
                                                        if false {
                                                            v51 = (v51 & v51)
                                                            switch 534 {
                                                            case 42:
                                                                if len(v54) >= 2 {
                                                                    v54 = v54[1:len(v54)]
                                                                }
                                                            case 541:
                                                                fmt.Println(v48)
                                                            case -2147483648:
                                                                v73 := GenS1{F0: 9223372036854775807, F1: 42}
                                                                _ = v73
                                                            }
                                                        }
                                                        var v74 int16 = (v4 >> (((v51 & uint(v2))) & 63))
                                                        _ = v74
                                                    }
                                                    _ = v71
                                                }
                                                v75 := GenS6{F0: v51, F1: v48, F2: v4}
                                                _ = v75
                                            }
                                        case 128:
                                            v76 := GenS2{F0: (v51 / (((v51 / (((v51 >> ((v51) & 63))) | 1))) | 1)), F1: 349, F2: 1}
                                            _ = v76
                                        }
                                        if false {
                                            if len(v54) >= 2 {
                                                v54 = v54[1:len(v54)]
                                            }
                                            if v57 {
                                                if v57 {
                                                    v77 := GenS3{F0: v1, F1: 2147483647}
                                                    if v57 {
                                                        if len(v54) > 0 {
                                                            v54[0] = 128
                                                        }
                                                        var v78 uint = 302
                                                        v51 = (v51 << ((v51) & 63))
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            var v79 int32 = -1
                                                            _ = v79
                                                        case "hello":
                                                            v51 = v78
                                                        case "a":
                                                            v78 = v78
                                                        }
                                                        _ = v78
                                                    } else {
                                                        v2 += v2
                                                        v80 := GenS8{F0: v4}
                                                        if (v57 && ((42 == 2) && (0 > 689))) {
                                                            v57 = (255 <= 86)
                                                            var v81 int64 = int64(v1)
                                                            v82 := GenS3{F0: (^(v1)), F1: 4294967295}
                                                            _ = v81
                                                            _ = v82
                                                        }
                                                        var v83 int32 = int32(v51)
                                                        _ = v80
                                                        _ = v83
                                                    }
                                                    fmt.Println(v57, v48)
                                                    _ = v77
                                                }
                                            }
                                        }
                                    } else {
                                        v6 -= (v6 & v6)
                                        var v84 int = (fi52 >> ((v51) & 63))
                                        var v85 uint16 = 32767
                                        _ = v84
                                        _ = v85
                                    }
                                case -2147483648:
                                    var v86 uint16 = (^(v1))
                                    _ = v86
                                }
                                _ = v57
                            }
                        case "\n\t":
                            v87 := GenS5{F0: (!((-(fi53)) >= -9223372036854775808)), F1: 82}
                            _ = v87
                        default:
                            v1 += 255
                        }
                        if false {
                            var v88 string = "\n\t"
                            _ = v88
                        }
                        fmt.Println(v4, v48)
                    }
                    if false {
                        v51 = v51
                        v2 -= (v2 << ((v51) & 63))
                    } else {
                        v3 = v3
                        v89 := GenS8{F0: (^(v4))}
                        _ = v89
                    }
                    _ = fi53
                    _ = v54
                }
                v90 := GenS9{F0: (v51 >> ((v51) & 63)), F1: -1, F2: true}
                _ = fi52
                _ = v90
            }
            var v91 uint8 = uint8(v2)
            var v92 int = -610
            _ = v91
            _ = v92
        }
    default:
        v93 := GenS10{F0: -1, F1: 2}
        _ = v93
    }
    if ((v48 << ((v51) & 63)) >= 100) {
        for fi94 := 0; fi94 < 13; fi94++ {
            var v95 uint32 = 0
            _ = fi94
            _ = v95
        }
    } else {
        if false {
            v96 := GenS6{F0: (v51 << ((v51) & 63)), F1: genFn0(), F2: v4}
            v6 += (v6 - -128)
            var v97 uint32 = 0
            _ = v96
            _ = v97
        } else {
            var v98 uint8 = 1
            _ = v98
        }
        fmt.Println(v3)
        v51 -= 2
    }
    for fi99 := 0; fi99 < 10; fi99++ {
        for fi100 := 0; fi100 < 6; fi100++ {
            v101 := GenS9{F0: (v51 * v51), F1: -35, F2: ((v1 >= v1) && true)}
            var v102 uint8 = 218
            if true {
                fmt.Println(v2, v102, v1)
                if (3.14159 >= ((((((1e10 - 1.5) - 0.1) + 0.0) - 3.14159) - ((3.14159 * float64(v5)) / ((1e10 + (float64(v5) + (1.5 / 1.5))) + (float64((v5 * v5)) + float64(float32(((float64(((v5 + float32((1e10 / 1.0))) * v5)) * -2.25) - 0.1))))))) * -2.25)) {
                    v3 = v3
                    v6 = (v48 >> ((255) & 63))
                    fmt.Println(v5, v5, fi100)
                    v6 = (v48 & (^(v6)))
                    var v103 string = "\n\t"
                    _ = v103
                } else {
                    v104 := []int{(fi99 - int(v6)), fi99, 4294967295}
                    fmt.Println(fi99)
                    if (!(((((v5 / 1e10) + (((((v5 / (v5 + (float32(((0.0 + -2.25) * float64(-2.25))) / -2.25))) - v5) + v5) / v5) * v5)) / v5) * 0.1) != v5)) {
                        var v105 uint32 = 896
                        var v106 int = fi100
                        _ = v105
                        _ = v106
                    }
                    _ = v104
                }
            }
            _ = fi100
            _ = v101
            _ = v102
        }
        if (false || (555 == uint32(v1))) {
            v107 := GenS8{F0: (v4 >> ((784) & 63))}
            var v108 uint32 = 316
            fmt.Println(v3)
            _ = v107
            _ = v108
        }
        _ = fi99
    }
    var v109 uint64 = 217
    fmt.Println("v1:", v1)
    fmt.Println("v109:", v109)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v109
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v51
    _ = v6
}
