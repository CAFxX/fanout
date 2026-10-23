package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 int; F2 int32 }
type GenS4 struct { F0 int; F1 int32; F2 uint64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int32; F1 int32; F2 int32 }
type GenS10 struct { F0 bool; F1 uint8; F2 uint }
type GenS11 struct { F0 int32; F1 uint32 }
type GenS12 struct { F0 int8; F1 uint8 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 int64; F1 int16 }
func genFn0() uint32 {
    for fi4 := 0; fi4 < 12; fi4++ {
        var v5 int32 = 1
        v6 := []int{1, (fi4 - fi4)}
        _ = fi4
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: 0, F1: false}
    var v8 uint32 = 1
    if true {
        var v9 uint16 = 755
        var v10 uint = 4294967295
        _ = v10
        _ = v9
    }
    _ = v7
    _ = v8
    return (^(v8))
}
func genFn1() string {
    for fi11 := 0; fi11 < 17; fi11++ {
        if true {
            switch 32767 {
            case 32767:
                var v12 uint8 = 128
                _ = v12
            case 9223372036854775807:
                fmt.Println(fi11)
            case -643:
                for fi13 := 0; fi13 < 19; fi13++ {
                    var v14 uint64 = 18446744073709551615
                    if false {
                        var v15 float32 = 0.1
                        _ = v15
                    } else {
                        var v16 int8 = 100
                        var v17 float32 = (1e10 / float32(0.1))
                        var v18 uint64 = v14
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                    v14 += (v14 ^ (v14 + v14))
                    _ = fi13
                    _ = v14
                }
            }
            fmt.Println(fi11)
            v19 := []int{2147483647, -232, (fi11 & fi11)}
            _ = v19
        } else {
            v20 := GenS1{F0: -2147483648, F1: (!(9223372036854775807 != 1))}
            var v21 int8 = 127
            switch "foo bar" {
            case "foo bar":
                fmt.Println(v21)
            case "":
                if ((202 >= 0) && false) {
                    var v22 uint16 = 1
                    v23 := GenS1{F0: (fi11 ^ (fi11 * (fi11 << ((65535) & 63)))), F1: true}
                    _ = v22
                    _ = v23
                }
            case "a":
                var v24 uint64 = 34
                _ = v24
            default:
                if ((-88 <= 0) && (true || true)) {
                    v25 := []int{-1, 32767, fi11}
                    for fi26 := 0; fi26 < 13; fi26++ {
                        v21 = v21
                        var v27 int = fi26
                        v28 := GenS1{F0: (-(v27)), F1: (935 < -924)}
                        var v29 int32 = -933
                        if (!false) {
                            fmt.Println(v29)
                            switch ("" + (("" + ((("" + "\n\t") + "a") + ("hello" + ("hello" + "hello")))) + ("a" + "a"))) {
                            case "a":
                                v30 := GenS1{F0: 9223372036854775807, F1: true}
                                _ = v30
                            default:
                                if true {
                                    switch "" {
                                    case "hello":
                                        if true {
                                            if (v27 != (^(fi26))) {
                                                var v31 uint16 = 2
                                                var v32 uint16 = v31
                                                if len(v25) > 0 {
                                                    v25[0] = int(v31)
                                                }
                                                _ = v31
                                                _ = v32
                                            }
                                            v27 = (v27 | v27)
                                            v33 := GenS1{F0: 9223372036854775807, F1: true}
                                            _ = v33
                                        } else {
                                            if len(v25) >= 2 {
                                                v25 = v25[1:len(v25)]
                                            }
                                            v29 -= -2147483648
                                            v34 := GenS2{F0: 42}
                                            _ = v34
                                        }
                                    case "":
                                        if ((1 != 732) || false) {
                                            if len(v25) > 0 {
                                                v25[0] = (-(fi26))
                                            }
                                            switch (v27 * 42) {
                                            case 1:
                                                if false {
                                                    v21 += v21
                                                    if len(v25) > 0 {
                                                        v25[0] = (^(fi11))
                                                    }
                                                    v29 += -1
                                                } else {
                                                    var v35 uint8 = 2
                                                    _ = v35
                                                }
                                                fmt.Println(fi26, fi11)
                                            case 602:
                                                if ((!((1.5 / 1e10) != ((((-2.25 + (float64(0.1) * (1.5 - ((-2.25 + 3.14159) / ((1.5 - 1e10) / ((0.1 / float64((3.14159 / 0.1))) * -2.25)))))) - ((-2.25 - 1.5) * float64((1e10 * 3.14159)))) + (1e10 / 1.0)) - float64((((1.5 * 1.5) + 3.14159) * (float32(0.1) * 3.14159)))))) || false) {
                                                    if false {
                                                        var v36 int16 = int16(fi11)
                                                        if false {
                                                            var v37 int64 = 9223372036854775807
                                                            if len(v25) >= 2 {
                                                                v25 = v25[1:len(v25)]
                                                            }
                                                            if ((^(v36)) <= (v36 % ((-1) | 1))) {
                                                                if (v37 > (v37 << ((0) & 63))) {
                                                                    switch 127 {
                                                                    case 2147483647:
                                                                        var v38 uint16 = 1
                                                                        _ = v38
                                                                    case -691:
                                                                        var v39 uint32 = 2
                                                                        _ = v39
                                                                    default:
                                                                        v40 := GenS2{F0: 2}
                                                                        _ = v40
                                                                    }
                                                                } else {
                                                                    v41 := GenS3{F0: 42, F1: fi26, F2: (v29 / ((v29) | 1))}
                                                                    v21 = (v21 / (((v21 - int8(fi26))) | 1))
                                                                    if (1 > 255) {
                                                                        fmt.Println(v21, v29, v27)
                                                                        if len(v25) > 0 {
                                                                            v25[0] = v27
                                                                        }
                                                                    }
                                                                    _ = v41
                                                                }
                                                                if true {
                                                                    v36 += (v36 + (v36 | 1000))
                                                                }
                                                            }
                                                            _ = v37
                                                        }
                                                        _ = v36
                                                    } else {
                                                        switch "" {
                                                        case "a":
                                                            var v42 string = "a"
                                                            switch -60 {
                                                            case -231:
                                                                if ((true && (95 >= 128)) || true) {
                                                                    var v43 uint64 = 1
                                                                    _ = v43
                                                                }
                                                            case 127:
                                                                switch "foo bar" {
                                                                case "a":
                                                                    if true {
                                                                        _ = len(v25)
                                                                    } else {
                                                                        var v44 float32 = (0.0 / 3.14159)
                                                                        fmt.Println(fi11, fi26)
                                                                        v45 := GenS4{F0: v27, F1: -627, F2: 0}
                                                                        _ = v44
                                                                        _ = v45
                                                                    }
                                                                default:
                                                                    v27 = (v27 % (((fi11 % (((-(fi11))) | 1))) | 1))
                                                                }
                                                                fmt.Println(fi11, fi26)
                                                                fmt.Println(fi11, v42, v27)
                                                                fmt.Println(v29, fi11)
                                                                fmt.Println(v42, fi26, fi11)
                                                            default:
                                                                var v46 string = ((("hello" + "hello") + "foo bar") + ("\n\t" + (("a" + v42) + v42)))
                                                                _ = v46
                                                            }
                                                            _ = v42
                                                        case "\n\t":
                                                            if false {
                                                                v47 := GenS5{F0: 4294967295}
                                                                _ = v47
                                                            } else {
                                                                v27 = -480
                                                            }
                                                        case "hello":
                                                            v21 += v21
                                                        default:
                                                            var v48 int8 = v21
                                                            _ = v48
                                                        }
                                                        v49 := GenS5{F0: 18446744073709551615}
                                                        _ = v49
                                                    }
                                                    v29 += v29
                                                }
                                            default:
                                                v50 := GenS1{F0: 32768, F1: false}
                                                _ = v50
                                            }
                                            v25 = append(v25, -34)
                                        } else {
                                            _ = len(v25)
                                            v51 := GenS6{F0: "hello"}
                                            _ = len(v25)
                                            _ = v51
                                        }
                                    case "\n\t":
                                        var v52 float32 = 1e10
                                        _ = v52
                                    default:
                                        var v53 int16 = 1
                                        var v54 int64 = -9223372036854775808
                                        v21 = v21
                                        _ = v53
                                        _ = v54
                                    }
                                    if false {
                                        v21 -= -1
                                        if true {
                                            v55 := GenS7{F0: 2}
                                            switch -560 {
                                            case 1:
                                                _ = len(v25)
                                                var v56 int32 = v29
                                                if ((1.5 - -2.25) <= float32(-2.25)) {
                                                    if false {
                                                        v57 := GenS8{F0: -2147483648}
                                                        _ = v57
                                                    }
                                                }
                                                _ = v56
                                            case 4294967295:
                                                if true {
                                                    v29 = (v29 & -2147483648)
                                                } else {
                                                    var v58 string = "hello"
                                                    _ = v58
                                                }
                                            case -642:
                                                v29 -= -1
                                            }
                                            _ = v55
                                        } else {
                                            v21 += v21
                                        }
                                        fmt.Println(v21, v21)
                                    }
                                    v27 = fi11
                                } else {
                                    var v59 uint32 = 0
                                    switch 0 {
                                    case 32767:
                                        switch "a" {
                                        case "foo bar":
                                            switch "foo bar" {
                                            case "a":
                                                _ = len(v25)
                                                if len(v25) >= 2 {
                                                    v25 = v25[1:len(v25)]
                                                }
                                                v21 = (^(v21))
                                            case "\n\t":
                                                fmt.Println(fi11, fi26, v59)
                                            case "hello":
                                                var v60 int64 = 170
                                                _ = v60
                                            }
                                        }
                                        v29 += 1
                                        if true {
                                            v21 -= v21
                                            var v61 uint32 = v59
                                            fmt.Println(v61, v59)
                                            var v62 int8 = 127
                                            var v63 bool = (!(((v29 << ((304) & 63)) == (v29 + -135)) || false))
                                            _ = v61
                                            _ = v62
                                            _ = v63
                                        } else {
                                            var v64 int32 = v29
                                            if (-2.25 <= ((1e10 * -2.25) * (-2.25 + (((-2.25 / float64(-2.25)) / -2.25) - -2.25)))) {
                                                var v65 int32 = 42
                                                if (((((2 != 222) && ((v29 == (v64 >> ((255) & 63))) || (42 <= -556))) && true) && true) || false) {
                                                    v66 := GenS7{F0: 42}
                                                    _ = v66
                                                } else {
                                                    v64 = (v29 % (((^(v65))) | 1))
                                                }
                                                v67 := GenS2{F0: 255}
                                                v59 = v59
                                                var v68 uint32 = (v59 + (v59 & v59))
                                                v65 += (v64 | -201)
                                                var v69 int = v27
                                                _ = v65
                                                _ = v67
                                                _ = v68
                                                _ = v69
                                            }
                                            if len(v25) > 0 {
                                                v25[0] = 820
                                            }
                                            _ = v64
                                        }
                                    case -2147483648:
                                        var v70 int32 = (v29 ^ v29)
                                        _ = v70
                                    }
                                    _ = v59
                                }
                            }
                        }
                        _ = fi26
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                    var v71 string = ""
                    _ = v25
                    _ = v71
                } else {
                    v72 := GenS7{F0: 0}
                    fmt.Println(fi11, fi11)
                    _ = v72
                }
            }
            _ = v20
            _ = v21
        }
        _ = fi11
    }
    if true {
        v73 := []int{4294967295, 2147483647, 42}
        _ = len(v73)
        _ = v73
    }
    if true {
        if (-128 >= -9) {
            v74 := GenS1{F0: -9223372036854775808, F1: true}
            _ = v74
        } else {
            v75 := []int{2147483647, -67, 42, -1}
            _ = v75
        }
        var v76 uint8 = 255
        _ = v76
    }
    return "\n\t"
}

func main() {
    var v1 string = (("" + "\n\t") + "\n\t")
    v2 := make(map[string]bool)
    v3 := make(map[uint32]string)
    var v77 string = v1
    for fi78 := 0; fi78 < 6; fi78++ {
        v77 = v77
        v1 = v1
        v3[65535] = v1
        _ = fi78
    }
    v79 := GenS9{F0: 1, F1: 2147483647, F2: -2147483648}
    for fi80 := 0; fi80 < 16; fi80++ {
        for fi81 := 0; fi81 < 15; fi81++ {
            switch (^(fi81)) {
            case -254:
                fmt.Println(v77, fi81)
            }
            var v82 uint = 2
            fmt.Println(fi80)
            _ = fi81
            _ = v82
        }
        _ = fi80
    }
    if false {
        v1 = v1
        var v83 int8 = 100
        _ = v83
    }
    var v84 int = 32767
    v84 += v84
    var v85 int64 = -9223372036854775808
    if (0 > -151) {
        var v86 uint64 = 2
        for fi87 := 0; fi87 < 6; fi87++ {
            v86 -= 2
            for fi88 := 0; fi88 < 6; fi88++ {
                v89 := GenS10{F0: false, F1: 235, F2: 351}
                if true {
                    fmt.Println(v77, v77)
                    var v90 int64 = v85
                    var v91 int16 = 1000
                    v86 += (v86 << ((2) & 63))
                    _ = v90
                    _ = v91
                }
                fmt.Println(v86)
                _ = fi88
                _ = v89
            }
            if (42 <= 42) {
                v85 += v85
            } else {
                v77 = "\n\t"
                var v92 uint = 42
                _ = v92
            }
            _ = fi87
        }
        v93 := GenS5{F0: (v86 - (v86 ^ (v86 | v86)))}
        _ = v86
        _ = v93
    } else {
        v94 := GenS2{F0: 25}
        v84 -= (v84 - (-(v84)))
        var v95 int8 = 30
        _ = v94
        _ = v95
    }
    v96 := []int{}
    for fi97 := 0; fi97 < 19; fi97++ {
        var v98 int8 = 1
        _ = fi97
        _ = v98
    }
    fmt.Println(v85, v1, v84)
    fmt.Println(v1)
    var v99 uint32 = 679
    if len(v96) >= 2 {
        v96 = v96[1:len(v96)]
    }
    if (18446744073709551615 != 18446744073709551615) {
        v100 := GenS11{F0: -2147483648, F1: 1}
        var v101 float64 = 0.0
        _ = v100
        _ = v101
    } else {
        delete(v2, (v77 + v1))
        var v102 int16 = -1
        _ = v102
    }
    for fi103 := 0; fi103 < 16; fi103++ {
        var v104 uint16 = 0
        switch -699 {
        case -319:
            switch v77 {
            case "":
                v84 = 0
            case "foo bar":
                v105 := GenS12{F0: 1, F1: 103}
                _ = v105
            case "a":
                switch fi103 {
                case -850:
                    v2[("\n\t" + "hello")] = false
                case 4294967295:
                    v99 -= 304
                default:
                    if ((v84 << ((65535) & 63)) < 511) {
                        v2[v1] = true
                        v84 = -625
                    } else {
                        if false {
                            switch v1 {
                            case "hello":
                                var v106 int8 = 113
                                _ = v106
                            }
                            switch (("\n\t" + v1) + "") {
                            case "\n\t":
                                delete(v3, v99)
                            case "":
                                var v107 float32 = (0.0 * 1.5)
                                _ = v107
                            case "hello":
                                v85 -= (v85 & (v85 / ((v85) | 1)))
                            default:
                                fmt.Println(v99, fi103, fi103)
                            }
                        } else {
                            v108 := GenS13{F0: (v104 ^ v104)}
                            if true {
                                v104 = v104
                                v84 = 128
                                v84 -= fi103
                            } else {
                                var v109 string = v77
                                var v110 string = (v109 + "")
                                _ = v109
                                _ = v110
                            }
                            _ = v108
                        }
                        v99 = 42
                        v84 = (fi103 >> ((255) & 63))
                    }
                }
            }
        case -9223372036854775808:
            for fi111 := 0; fi111 < 11; fi111++ {
                if ((v104 >> ((1) & 63)) >= (v104 - v104)) {
                    if ((2 >= 42) || ((589 == 4294967295) && true)) {
                        v96 = append(v96, 4294967295)
                        if false {
                            if false {
                                v99 += 2
                            } else {
                                v104 -= v104
                            }
                            v112 := GenS14{F0: -514}
                            delete(v2, v1)
                            fmt.Println(v85)
                            var v113 uint32 = (v99 & 231)
                            _ = v112
                            _ = v113
                        } else {
                            v84 = 2147483647
                            fmt.Println(v85)
                            var v114 string = (((v1 + v77) + ("a" + v77)) + (v1 + "foo bar"))
                            v115 := GenS15{F0: (-(v85)), F1: -511}
                            v1 = ("foo bar" + (v77 + "foo bar"))
                            _ = v114
                            _ = v115
                        }
                        fmt.Println(v77, fi103)
                        v99 -= (v99 % ((genFn0()) | 1))
                        switch (fi111 >> ((65535) & 63)) {
                        case 32768:
                            v116 := GenS2{F0: 255}
                            _ = v116
                        case -1:
                            var v117 int64 = (v85 >> ((255) & 63))
                            _ = v117
                        }
                    }
                    fmt.Println(fi111, v104)
                }
                if (-32768 <= 1000) {
                    fmt.Println(v85)
                    var v118 int8 = 42
                    v119 := v3[genFn0()]
                    _ = v118
                    _ = v119
                }
                _ = fi111
            }
        default:
            v1 = (v1 + v77)
        }
        var v120 uint64 = 18446744073709551615
        _ = fi103
        _ = v104
        _ = v120
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v77:", v77)
    fmt.Println("v84:", v84)
    fmt.Println("v85:", v85)
    fmt.Println("slicelen v96:", len(v96), v96)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v2
    _ = v3
    _ = v77
    _ = v79
    _ = v84
    _ = v85
    _ = v96
    _ = v99
}
