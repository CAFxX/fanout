package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 uint; F2 uint16 }
type GenS4 struct { F0 uint; F1 uint32; F2 uint64 }
type GenS5 struct { F0 int32; F1 uint16 }
type GenS6 struct { F0 int32; F1 uint16 }
type GenS7 struct { F0 uint32; F1 int16 }
type GenS8 struct { F0 string; F1 bool; F2 bool }
type GenS9 struct { F0 int32; F1 uint64 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int32; F1 bool }
type GenS13 struct { F0 uint64; F1 int }
type GenS14 struct { F0 int64; F1 int64 }
type GenS15 struct { F0 bool; F1 int64; F2 int64 }
type GenS16 struct { F0 int16; F1 uint64 }
type GenS17 struct { F0 int8 }
type GenS18 struct { F0 int32; F1 uint16; F2 int16 }
type GenS19 struct { F0 uint8; F1 uint16; F2 uint32 }
func genFn0(p0 int, p1 string) int8 {
    v6 := GenS1{F0: 100, F1: 255}
    if (163 <= 78) {
        var v7 uint8 = 2
        switch (p0 << ((uint(v7)) & 63)) {
        case 1:
            if ((-9223372036854775808 == 640) && false) {
                for fi8 := 0; fi8 < 3; fi8++ {
                    v9 := []int{128, p0}
                    if len(v9) > 0 {
                        v9[0] = p0
                    }
                    fmt.Println(p0)
                    _ = fi8
                    _ = v9
                }
            } else {
                v7 -= 0
                var v10 int32 = -963
                p0 = (^(p0))
                _ = v10
            }
        }
        v11 := GenS2{F0: true}
        _ = v11
        _ = v7
    }
    if (!true) {
        v12 := GenS1{F0: 100, F1: 1}
        _ = v12
    }
    if (1 >= p0) {
        v13 := []int{4294967295, (p0 >> ((18446744073709551615) & 63)), (p0 >> ((65535) & 63))}
        switch p1 {
        case "":
            p0 = -2147483648
        }
        p1 = p1
        _ = v13
    } else {
        v14 := []int{p0, (p0 ^ (-(p0))), (p0 >> ((787) & 63)), 219}
        _ = v14
    }
    p0 -= (p0 / (((p0 % ((442) | 1))) | 1))
    _ = p0
    _ = p1
    _ = v6
    return 42
}
func genFn1(p0 int, p1 string, p2 uint32) uint32 {
    var v15 int8 = -128
    p2 -= 917
    var v16 int64 = 0
    fmt.Println(v15, v15, p1)
    v17 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    _ = v17
    return (p2 >> ((2) & 63))
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[string]string)
    var v3 bool = true
    var v4 float32 = v1
    var v5 uint8 = 0
    v3 = (true || ((-1 < 127) || (!(9223372036854775807 > -1))))
    switch "hello" {
    case "a":
        if v3 {
            if (true && v3) {
                if (!(-1 <= 0)) {
                    v5 += (^(v5))
                    v2["foo bar"] = "foo bar"
                    if (42 > 0) {
                        var v18 float64 = 3.14159
                        _ = len(v2)
                        delete(v2, "foo bar")
                        _ = v18
                    }
                }
            } else {
                switch -1 {
                case -804:
                    var v19 uint8 = v5
                    _ = v19
                }
                for fi20 := 0; fi20 < 1; fi20++ {
                    v21 := GenS2{F0: (980 >= fi20)}
                    if (20 >= 2) {
                        for fi22 := 0; fi22 < 3; fi22++ {
                            v1 -= 1e10
                            v4 = float32(1e10)
                            v23 := GenS1{F0: 42, F1: uint(v5)}
                            _ = fi22
                            _ = v23
                        }
                        for fi24 := 0; fi24 < 19; fi24++ {
                            if (517 <= 961) {
                                fmt.Println(v3)
                                var v25 int8 = 42
                                v3 = ((1 == 18446744073709551615) && (v3 && true))
                                _ = v25
                            } else {
                                switch (fi24 ^ -754) {
                                case 32768:
                                    v26 := GenS1{F0: 78, F1: 65535}
                                    _ = v26
                                case 127:
                                    v5 = 243
                                case 128:
                                    var v27 float32 = (v1 - 1e10)
                                    _ = v27
                                default:
                                    if (!true) {
                                        var v28 int64 = 613
                                        fmt.Println(v28, fi24, fi24)
                                        _ = v28
                                    } else {
                                        var v29 int = fi24
                                        var v30 float32 = (v4 * v1)
                                        _ = v29
                                        _ = v30
                                    }
                                }
                            }
                            var v31 uint8 = v5
                            var v32 int = (fi20 % (((-(fi20))) | 1))
                            v33 := GenS1{F0: -1, F1: 928}
                            _ = fi24
                            _ = v31
                            _ = v32
                            _ = v33
                        }
                    }
                    _ = fi20
                    _ = v21
                }
            }
            v34 := GenS1{F0: -57, F1: 42}
            _ = v34
        }
    }
    v3 = true
    if (2 > 0) {
        v35 := GenS2{F0: true}
        var v36 uint8 = v5
        _ = v35
        _ = v36
    } else {
        if (1 > 1) {
            v4 += (v4 / (v1 * 0.1))
            if false {
                if false {
                    var v37 uint8 = v5
                    _ = v37
                }
            }
            v38 := GenS1{F0: -1, F1: 18446744073709551615}
            _ = v38
        }
        if v3 {
            v39 := []int{2147483647, int(v5), -9223372036854775808}
            if (!false) {
                if true {
                    v3 = v3
                    v39 = append(v39, 315)
                    var v40 int = -800
                    _ = v40
                }
                v5 = v5
                fmt.Println(v1, v4, v5)
            }
            _ = v39
        } else {
            delete(v2, "\n\t")
            if v3 {
                if true {
                    v41 := GenS2{F0: v3}
                    var v42 uint64 = 18446744073709551615
                    for fi43 := 0; fi43 < 14; fi43++ {
                        v44 := []int{0, 32767, -2147483648, (fi43 - fi43)}
                        _ = fi43
                        _ = v44
                    }
                    _ = v41
                    _ = v42
                } else {
                    switch 32768 {
                    case 127:
                        if v3 {
                            if (!v3) {
                                var v45 int = 9223372036854775807
                                _ = v45
                            }
                            var v46 float64 = 0.1
                            _ = v46
                        } else {
                            if v3 {
                                v47 := GenS2{F0: v3}
                                for fi48 := 0; fi48 < 12; fi48++ {
                                    var v49 uint8 = v5
                                    _ = fi48
                                    _ = v49
                                }
                                _ = v47
                            }
                            v4 += 0.1
                            fmt.Println(v1, v1, v5)
                        }
                    case 32767:
                        if (!((((float64(float32(((3.14159 + (0.0 / (3.14159 + 0.0))) - 3.14159))) - ((-2.25 * -2.25) * 0.0)) - 0.0) + -2.25) != -2.25)) {
                            if (-66 != -1) {
                                if (1 > 92) {
                                    var v50 int = 128
                                    _ = v50
                                } else {
                                    if (!(!v3)) {
                                        fmt.Println(v1, v4, v1)
                                    } else {
                                        if (!(false || v3)) {
                                            for fi51 := 0; fi51 < 11; fi51++ {
                                                for fi52 := 0; fi52 < 12; fi52++ {
                                                    v1 = ((v1 * 0.0) - (v4 + 0.1))
                                                    v53 := GenS3{F0: 4294967295, F1: 65535, F2: 1}
                                                    if (!(v5 == (v5 << ((255) & 63)))) {
                                                        v54 := []int{(fi51 * (fi52 / (((fi51 | (fi52 >> ((42) & 63)))) | 1))), (fi51 * (fi51 % ((fi51) | 1))), fi52}
                                                        var v55 bool = (4294967295 == 2)
                                                        var v56 uint = 255
                                                        switch "" {
                                                        case "a":
                                                            v4 = 0.0
                                                        case "foo bar":
                                                            var v57 uint32 = 1
                                                            _ = v57
                                                        case "hello":
                                                            if (-2.25 > (float64(v4) * -2.25)) {
                                                                v58 := GenS4{F0: (^(v56)), F1: 479, F2: 858}
                                                                v54 = append(v54, -2147483648)
                                                                if v55 {
                                                                    v59 := GenS1{F0: -1, F1: v56}
                                                                    if v3 {
                                                                        var v60 int16 = -684
                                                                        var v61 bool = v3
                                                                        if false {
                                                                            switch "hello" {
                                                                            case "hello":
                                                                                switch fi52 {
                                                                                case -354:
                                                                                    switch "hello" {
                                                                                    case "foo bar":
                                                                                        if (264 != 65535) {
                                                                                            v61 = true
                                                                                            if (!false) {
                                                                                                v2["a"] = ""
                                                                                                v54 = append(v54, int(v5))
                                                                                            } else {
                                                                                                fmt.Println(fi52, v60)
                                                                                                switch "hello" {
                                                                                                case "hello":
                                                                                                    if v3 {
                                                                                                        if (1.5 > v4) {
                                                                                                            _ = len(v54)
                                                                                                            var v62 bool = (!(!false))
                                                                                                            _ = v62
                                                                                                        }
                                                                                                        v54 = append(v54, (fi52 + fi51))
                                                                                                    }
                                                                                                default:
                                                                                                    v63 := GenS5{F0: -950, F1: 32767}
                                                                                                    _ = v63
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    case "\n\t":
                                                                                        v64 := GenS6{F0: 1, F1: 464}
                                                                                        _ = v64
                                                                                    case "":
                                                                                        v65 := GenS3{F0: fi51, F1: v56, F2: 32767}
                                                                                        switch fi52 {
                                                                                        case 2147483647:
                                                                                            var v66 uint64 = 682
                                                                                            _ = v66
                                                                                        case 128:
                                                                                            fmt.Println(v3)
                                                                                            var v67 uint16 = 1
                                                                                            _ = v67
                                                                                        default:
                                                                                            var v68 uint8 = 0
                                                                                            _ = v68
                                                                                        }
                                                                                        if false {
                                                                                            if len(v54) >= 2 {
                                                                                                v54 = v54[1:len(v54)]
                                                                                            }
                                                                                            var v69 uint16 = 280
                                                                                            _ = v69
                                                                                        }
                                                                                        _ = v65
                                                                                    default:
                                                                                        fmt.Println(v61, fi52, v56)
                                                                                    }
                                                                                    if (42 != -2147483648) {
                                                                                        v60 = 0
                                                                                    }
                                                                                }
                                                                            default:
                                                                                var v70 uint = (^(v56))
                                                                                _ = v70
                                                                            }
                                                                            switch "" {
                                                                            case "hello":
                                                                                v1 = (v1 / (v1 + v1))
                                                                            case "foo bar":
                                                                                var v71 uint = (^(v56))
                                                                                var v72 float64 = 1e10
                                                                                fmt.Println(v60, fi52)
                                                                                v55 = ("foo bar" < ("hello" + "a"))
                                                                                v5 += 1
                                                                                _ = v71
                                                                                _ = v72
                                                                            case "a":
                                                                                fmt.Println(v60, v55, v60)
                                                                            }
                                                                            if (2147483647 > -1) {
                                                                                if (fi51 >= fi51) {
                                                                                    v60 = -1
                                                                                    if v55 {
                                                                                        switch fi51 {
                                                                                        case 2147483647:
                                                                                            switch "\n\t" {
                                                                                            case "hello":
                                                                                                v73 := GenS4{F0: (v56 % (((^(v56))) | 1)), F1: uint32(v56), F2: 981}
                                                                                                _ = v73
                                                                                            case "foo bar":
                                                                                                if v55 {
                                                                                                    v74 := v2[("hello" + "")]
                                                                                                    _ = v74
                                                                                                }
                                                                                            case "\n\t":
                                                                                                switch (fi51 - fi51) {
                                                                                                case 4294967295:
                                                                                                    v75 := GenS6{F0: 293, F1: 0}
                                                                                                    var v76 int8 = 42
                                                                                                    var v77 float64 = ((1.5 * 0.1) - (1.5 - 1.5))
                                                                                                    _ = v75
                                                                                                    _ = v76
                                                                                                    _ = v77
                                                                                                case 2147483647:
                                                                                                    if ("foo bar" <= "hello") {
                                                                                                        if (v60 > (v60 >> ((v56) & 63))) {
                                                                                                            v1 = v4
                                                                                                            v78 := GenS7{F0: 0, F1: 42}
                                                                                                            _ = v78
                                                                                                        } else {
                                                                                                            fmt.Println(v4)
                                                                                                            var v79 uint8 = v5
                                                                                                            fmt.Println(v4, v61)
                                                                                                            if v61 {
                                                                                                                var v80 uint = (v56 | v56)
                                                                                                                var v81 uint32 = 2147483648
                                                                                                                v82 := GenS8{F0: "", F1: ((!("hello" != ((("hello" + "") + "") + ""))) && ((v80 | 1) > (v56 - uint(v81)))), F2: v3}
                                                                                                                v83 := GenS5{F0: -1, F1: 0}
                                                                                                                _ = v80
                                                                                                                _ = v81
                                                                                                                _ = v82
                                                                                                                _ = v83
                                                                                                            }
                                                                                                            switch -9223372036854775808 {
                                                                                                            case 2147483647:
                                                                                                                var v84 int32 = 2147483647
                                                                                                                _ = v84
                                                                                                            default:
                                                                                                                v5 -= (v79 ^ (v79 - v5))
                                                                                                                v79 = (^(v79))
                                                                                                            }
                                                                                                            _ = v79
                                                                                                        }
                                                                                                        if len(v54) > 0 {
                                                                                                            v54[0] = 0
                                                                                                        }
                                                                                                        var v85 float32 = v1
                                                                                                        v2["foo bar"] = ("" + "\n\t")
                                                                                                        _ = v85
                                                                                                    } else {
                                                                                                        _ = len(v2)
                                                                                                        if (3.14159 != (float64((float32((1.5 / 1.0)) / 1.0)) - 0.1)) {
                                                                                                            var v86 string = ""
                                                                                                            v61 = v55
                                                                                                            v55 = (1 <= (^(fi52)))
                                                                                                            _ = v86
                                                                                                        }
                                                                                                        v55 = (42 <= -353)
                                                                                                        v87 := GenS9{F0: -532, F1: 4294967295}
                                                                                                        _ = v87
                                                                                                    }
                                                                                                    v88 := GenS1{F0: genFn0((fi51 - fi51), (("hello" + "") + "\n\t")), F1: 65535}
                                                                                                    var v89 int16 = v60
                                                                                                    _ = v88
                                                                                                    _ = v89
                                                                                                default:
                                                                                                    switch ("a" + ("a" + ("a" + "a"))) {
                                                                                                    case "\n\t":
                                                                                                        v1 -= float32(0.1)
                                                                                                    case "foo bar":
                                                                                                        fmt.Println(fi51, fi52, fi51)
                                                                                                    case "hello":
                                                                                                        if v61 {
                                                                                                            if ((fi51 - fi52) != fi51) {
                                                                                                                v56 = (^(v56))
                                                                                                                if ((!(0.1 < 1.5)) && (v1 < v4)) {
                                                                                                                    v90 := GenS4{F0: 990, F1: 0, F2: 42}
                                                                                                                    _ = v90
                                                                                                                }
                                                                                                            } else {
                                                                                                                var v91 uint8 = uint8(v56)
                                                                                                                fmt.Println(fi52, v61, v3)
                                                                                                                switch "a" {
                                                                                                                case "":
                                                                                                                    v5 -= v91
                                                                                                                case "a":
                                                                                                                    v61 = false
                                                                                                                case "foo bar":
                                                                                                                    var v92 int = fi52
                                                                                                                    _ = v92
                                                                                                                }
                                                                                                                v93 := GenS8{F0: "hello", F1: false, F2: v3}
                                                                                                                _ = v91
                                                                                                                _ = v93
                                                                                                            }
                                                                                                            v3 = true
                                                                                                        } else {
                                                                                                            var v94 float64 = 0.1
                                                                                                            _ = v94
                                                                                                        }
                                                                                                    default:
                                                                                                        v2[("hello" + ("" + "hello"))] = "a"
                                                                                                    }
                                                                                                }
                                                                                            default:
                                                                                                fmt.Println(v60, v4)
                                                                                            }
                                                                                        case 0:
                                                                                            v1 -= float32(3.14159)
                                                                                        default:
                                                                                            v95 := GenS10{F0: 2147483647}
                                                                                            _ = v95
                                                                                        }
                                                                                        v4 -= ((-2.25 + 3.14159) + float32((1.5 - 1e10)))
                                                                                        v96 := v2[("\n\t" + "hello")]
                                                                                        if ("" >= (v96 + ("a" + (v96 + v96)))) {
                                                                                            v97 := GenS11{F0: fi51}
                                                                                            v98 := GenS11{F0: 90}
                                                                                            var v99 string = (v96 + "foo bar")
                                                                                            _ = v97
                                                                                            _ = v98
                                                                                            _ = v99
                                                                                        } else {
                                                                                            v100 := GenS8{F0: "\n\t", F1: v3, F2: ("a" == v96)}
                                                                                            v5 = (v5 % ((v5) | 1))
                                                                                            _ = v100
                                                                                        }
                                                                                        var v101 int32 = -1
                                                                                        _ = v101
                                                                                        _ = v96
                                                                                    }
                                                                                    v60 -= v60
                                                                                } else {
                                                                                    if len(v54) > 0 {
                                                                                        v54[0] = fi52
                                                                                    }
                                                                                    if true {
                                                                                        var v102 int8 = 0
                                                                                        _ = v102
                                                                                    }
                                                                                }
                                                                            }
                                                                            if v55 {
                                                                                v4 = ((v1 * (float32(float64(v1)) - (v4 / (0.1 * v4)))) - float32(float64(v4)))
                                                                                fmt.Println(v3, v1)
                                                                                v103 := GenS8{F0: ("" + ("\n\t" + "")), F1: (175 > 42), F2: (v4 == 3.14159)}
                                                                                _ = v103
                                                                            }
                                                                            var v104 uint64 = 200
                                                                            _ = v104
                                                                        } else {
                                                                            if len(v54) > 0 {
                                                                                v54[0] = (fi52 % ((fi52) | 1))
                                                                            }
                                                                            v61 = ((^(v5)) <= 0)
                                                                            if v3 {
                                                                                var v105 float64 = 0.1
                                                                                switch ("hello" + "foo bar") {
                                                                                case "\n\t":
                                                                                    var v106 uint16 = 0
                                                                                    _ = v106
                                                                                case "foo bar":
                                                                                    delete(v2, "\n\t")
                                                                                case "":
                                                                                    var v107 uint = v56
                                                                                    _ = v107
                                                                                }
                                                                                switch -683 {
                                                                                case 190:
                                                                                    fmt.Println(v55, v56, v4)
                                                                                    delete(v2, "foo bar")
                                                                                    var v108 uint8 = 205
                                                                                    _ = v108
                                                                                }
                                                                                v109 := GenS8{F0: "foo bar", F1: (0 == 600), F2: v61}
                                                                                _ = v105
                                                                                _ = v109
                                                                            }
                                                                        }
                                                                        v2["foo bar"] = "a"
                                                                        v4 -= v4
                                                                        if v55 {
                                                                            if (((!((0.0 >= 3.14159) && (!v61))) && (42 < 0)) || true) {
                                                                                v110 := GenS8{F0: (("a" + ("" + ("a" + "a"))) + "\n\t"), F1: v3, F2: (4294967295 > 2147483647)}
                                                                                _ = v110
                                                                            } else {
                                                                                v55 = v3
                                                                            }
                                                                            if v55 {
                                                                                fmt.Println(v5, v1)
                                                                                var v111 int32 = -2147483648
                                                                                v60 += -32768
                                                                                _ = v111
                                                                            }
                                                                        }
                                                                        _ = v60
                                                                        _ = v61
                                                                    } else {
                                                                        var v112 uint16 = 268
                                                                        v113 := GenS12{F0: 1, F1: v55}
                                                                        v112 = v112
                                                                        _ = v112
                                                                        _ = v113
                                                                    }
                                                                    var v114 int16 = 42
                                                                    _ = v114
                                                                    _ = v59
                                                                } else {
                                                                    if (0 <= 0) {
                                                                        if len(v54) >= 2 {
                                                                            v54 = v54[1:len(v54)]
                                                                        }
                                                                        var v115 float64 = 1e10
                                                                        var v116 int32 = -2147483648
                                                                        fmt.Println(v3, v116, v55)
                                                                        _ = len(v54)
                                                                        _ = v115
                                                                        _ = v116
                                                                    }
                                                                    v117 := GenS7{F0: 0, F1: 1}
                                                                    _ = v117
                                                                }
                                                                switch ("\n\t" + ("hello" + "hello")) {
                                                                case "hello":
                                                                    _ = len(v54)
                                                                }
                                                                _ = v58
                                                            }
                                                        }
                                                        _ = v54
                                                        _ = v55
                                                        _ = v56
                                                    }
                                                    _ = fi52
                                                    _ = v53
                                                }
                                                v3 = false
                                                v118 := GenS1{F0: genFn0(fi51, ("\n\t" + "")), F1: 18446744073709551615}
                                                _ = fi51
                                                _ = v118
                                            }
                                            v119 := []int{0}
                                            _ = v119
                                        } else {
                                            var v120 uint = 4294967295
                                            v5 = (v5 >> ((1) & 63))
                                            _ = v120
                                        }
                                        switch -9223372036854775808 {
                                        case -343:
                                            for fi121 := 0; fi121 < 13; fi121++ {
                                                fmt.Println(v5, v1)
                                                for fi122 := 0; fi122 < 7; fi122++ {
                                                    v1 += v1
                                                    if (v3 && false) {
                                                        var v123 int64 = 177
                                                        if (674 != 2147483647) {
                                                            if true {
                                                                if (!((fi122 / (((fi122 | fi122)) | 1)) != fi122)) {
                                                                    if (1 <= 255) {
                                                                        if (false && v3) {
                                                                            var v124 uint8 = 127
                                                                            if v3 {
                                                                                v4 -= -2.25
                                                                                v124 = v124
                                                                                fmt.Println(fi121, fi122, v3)
                                                                            }
                                                                            var v125 int64 = v123
                                                                            _ = v124
                                                                            _ = v125
                                                                        } else {
                                                                            delete(v2, "\n\t")
                                                                            v4 -= float32(0.0)
                                                                            v5 -= 44
                                                                        }
                                                                        fmt.Println(fi122)
                                                                    }
                                                                }
                                                                _ = len(v2)
                                                            } else {
                                                                var v126 uint8 = (^(v5))
                                                                if true {
                                                                    v5 = (v5 / (((v5 % ((v126) | 1))) | 1))
                                                                    fmt.Println(fi122, v5, v3)
                                                                    v5 -= (^(v126))
                                                                } else {
                                                                    v4 = 0.0
                                                                }
                                                                v127 := []int{fi122}
                                                                v128 := GenS11{F0: fi121}
                                                                if false {
                                                                    if v3 {
                                                                        var v129 uint64 = 1
                                                                        var v130 int32 = 2147483647
                                                                        _ = v129
                                                                        _ = v130
                                                                    }
                                                                    v1 = float32(1.5)
                                                                    v1 -= v4
                                                                }
                                                                _ = v126
                                                                _ = v127
                                                                _ = v128
                                                            }
                                                        } else {
                                                            var v131 int = fi121
                                                            if true {
                                                                v1 = ((0.1 + (3.14159 + (1e10 * v1))) + (v1 / (float32(1e10) + 0.0)))
                                                                var v132 uint = 1
                                                                v123 = v123
                                                                var v133 int8 = 42
                                                                _ = v132
                                                                _ = v133
                                                            } else {
                                                                if (int(v123) <= v131) {
                                                                    v134 := []int{fi122}
                                                                    v2["a"] = "hello"
                                                                    var v135 int64 = 1
                                                                    _ = v134
                                                                    _ = v135
                                                                }
                                                                switch fi121 {
                                                                case 32767:
                                                                    v123 = v123
                                                                case 1:
                                                                    v4 = 1.5
                                                                }
                                                            }
                                                            v136 := GenS12{F0: -2147483648, F1: (v3 && (-1 > 878))}
                                                            v137 := []int{}
                                                            v131 = v131
                                                            _ = v131
                                                            _ = v136
                                                            _ = v137
                                                        }
                                                        _ = v123
                                                    } else {
                                                        v138 := GenS6{F0: -118, F1: 42}
                                                        delete(v2, "foo bar")
                                                        switch "a" {
                                                        case "a":
                                                            var v139 float32 = (v1 * (v1 / ((float32(((3.14159 / float64((v1 - v1))) * 1.5)) - v4) - float32((((-2.25 - 3.14159) / 1.5) - 0.1)))))
                                                            _ = v139
                                                        case "":
                                                            fmt.Println(fi122)
                                                            switch ("" + "\n\t") {
                                                            case "foo bar":
                                                                fmt.Println(v1, fi122)
                                                            default:
                                                                v140 := []int{(fi122 << ((658) & 63)), 42, (-(fi122))}
                                                                _ = v140
                                                            }
                                                        }
                                                        if (v3 && (((!v3) && v3) && v3)) {
                                                            v141 := []int{fi122, (fi121 << ((255) & 63)), fi122, (fi121 >> ((1) & 63))}
                                                            switch "\n\t" {
                                                            case "hello":
                                                                switch "a" {
                                                                case "":
                                                                    if (fi122 > (^(fi122))) {
                                                                        v3 = v3
                                                                        if (("\n\t" + "") > "foo bar") {
                                                                            var v142 uint = 1
                                                                            _ = v142
                                                                        } else {
                                                                            if (v3 || (-1 < 0)) {
                                                                                v4 = v1
                                                                            }
                                                                            v3 = v3
                                                                        }
                                                                        if true {
                                                                            _ = len(v141)
                                                                            var v143 float64 = float64(-2.25)
                                                                            switch ("a" + (("\n\t" + "foo bar") + "hello")) {
                                                                            case "a":
                                                                                var v144 int8 = -18
                                                                                _ = v144
                                                                            case "":
                                                                                fmt.Println(fi121, fi122)
                                                                            }
                                                                            v145 := GenS3{F0: fi121, F1: 2, F2: 1}
                                                                            fmt.Println(v1)
                                                                            var v146 int64 = 9223372036854775807
                                                                            _ = v143
                                                                            _ = v145
                                                                            _ = v146
                                                                        } else {
                                                                            var v147 int32 = 1
                                                                            v148 := GenS5{F0: (v147 & v147), F1: 255}
                                                                            v5 = v5
                                                                            v149 := GenS5{F0: v147, F1: 1}
                                                                            _ = v147
                                                                            _ = v148
                                                                            _ = v149
                                                                        }
                                                                    } else {
                                                                        fmt.Println(fi121, v3, v5)
                                                                    }
                                                                case "hello":
                                                                    switch ("hello" + "") {
                                                                    case "a":
                                                                        v3 = (v3 || (127 <= genFn0(fi121, "\n\t")))
                                                                    }
                                                                default:
                                                                    var v150 int8 = 127
                                                                    _ = v150
                                                                }
                                                            case "\n\t":
                                                                if v3 {
                                                                    v151 := GenS13{F0: 0, F1: fi122}
                                                                    if (-1 >= -2147483648) {
                                                                        if (v3 && ((504 == 423) && (4294967295 > 2))) {
                                                                            v3 = (v3 && (v5 <= (v5 - (v5 - 255))))
                                                                            var v152 uint32 = 0
                                                                            _ = v152
                                                                        }
                                                                        v153 := GenS2{F0: true}
                                                                        var v154 int = 209
                                                                        if (1 > 665) {
                                                                            var v155 bool = v3
                                                                            _ = v155
                                                                        }
                                                                        _ = v153
                                                                        _ = v154
                                                                    } else {
                                                                        if (2147483647 < 130) {
                                                                            if (703 <= 0) {
                                                                                v156 := GenS9{F0: -126, F1: 1}
                                                                                v3 = (false || (v3 && (!true)))
                                                                                if (!(-2147483648 > 0)) {
                                                                                    v4 = v1
                                                                                    var v157 string = ""
                                                                                    _ = v157
                                                                                } else {
                                                                                    if len(v141) >= 2 {
                                                                                        v141 = v141[1:len(v141)]
                                                                                    }
                                                                                    v2[("hello" + "a")] = ""
                                                                                }
                                                                                _ = v156
                                                                            }
                                                                            v3 = v3
                                                                        }
                                                                        var v158 int8 = -1
                                                                        switch fi122 {
                                                                        case 0:
                                                                            var v159 float64 = (-2.25 * 0.0)
                                                                            _ = v159
                                                                        default:
                                                                            var v160 bool = (v3 && true)
                                                                            _ = v160
                                                                        }
                                                                        _ = v158
                                                                    }
                                                                    v161 := GenS11{F0: int(v5)}
                                                                    _ = v151
                                                                    _ = v161
                                                                } else {
                                                                    v5 -= v5
                                                                    if v3 {
                                                                        v1 = 0.1
                                                                        v4 -= v1
                                                                        fmt.Println(fi121, v4)
                                                                    } else {
                                                                        if (v5 < v5) {
                                                                            v162 := GenS14{F0: -9223372036854775808, F1: -614}
                                                                            var v163 uint64 = 310
                                                                            _ = v162
                                                                            _ = v163
                                                                        } else {
                                                                            var v164 uint32 = 2147483648
                                                                            fmt.Println(fi121, v164, v164)
                                                                            var v165 float64 = -2.25
                                                                            _ = v164
                                                                            _ = v165
                                                                        }
                                                                    }
                                                                }
                                                                if (v5 < (v5 - v5)) {
                                                                    v166 := GenS2{F0: true}
                                                                    _ = v166
                                                                } else {
                                                                    v5 = v5
                                                                    v2["\n\t"] = "hello"
                                                                }
                                                                if (0.1 >= 1.5) {
                                                                    v1 = 3.14159
                                                                }
                                                            }
                                                            var v167 int32 = 201
                                                            v3 = v3
                                                            if (v167 < (v167 % ((int32(v5)) | 1))) {
                                                                if true {
                                                                    if (!(!false)) {
                                                                        v5 = 2
                                                                    } else {
                                                                        v168 := GenS9{F0: v167, F1: 42}
                                                                        _ = v168
                                                                    }
                                                                    v5 = v5
                                                                    v169 := v2["hello"]
                                                                    _ = v169
                                                                }
                                                                _ = len(v2)
                                                                v170 := GenS15{F0: v3, F1: 9223372036854775807, F2: -280}
                                                                _ = v170
                                                            } else {
                                                                switch "hello" {
                                                                case "hello":
                                                                    fmt.Println(v1, v1, v4)
                                                                }
                                                                if len(v141) >= 2 {
                                                                    v141 = v141[1:len(v141)]
                                                                }
                                                            }
                                                            var v171 float64 = 0.0
                                                            var v172 int16 = -1
                                                            _ = v141
                                                            _ = v167
                                                            _ = v171
                                                            _ = v172
                                                        }
                                                        _ = v138
                                                    }
                                                    fmt.Println(fi121, v1, v5)
                                                    var v173 float32 = (((v1 / v4) * 0.1) + -2.25)
                                                    _ = fi122
                                                    _ = v173
                                                }
                                                for fi174 := 0; fi174 < 17; fi174++ {
                                                    v4 += (((v4 / 0.1) / ((v1 / (v1 - ((1.5 + (-2.25 * v4)) - 0.0))) + float32((3.14159 / -2.25)))) * v4)
                                                    v4 = float32(1.5)
                                                    fmt.Println(v1, fi121)
                                                    v175 := []int{128}
                                                    v176 := GenS16{F0: 32767, F1: 2}
                                                    _ = fi174
                                                    _ = v175
                                                    _ = v176
                                                }
                                                _ = fi121
                                            }
                                        case 32768:
                                            if (127 <= 116) {
                                                var v177 string = ("" + "a")
                                                var v178 uint16 = 307
                                                _ = v177
                                                _ = v178
                                            }
                                        case -76:
                                            fmt.Println(v3, v1, v3)
                                        }
                                        var v179 string = ("hello" + "foo bar")
                                        _ = v179
                                    }
                                    v4 = 1.5
                                }
                            }
                            if ("\n\t" > ("hello" + ("a" + ("foo bar" + "")))) {
                                var v180 int8 = 126
                                v180 = genFn0(int(v5), "foo bar")
                                _ = v180
                            }
                        } else {
                            v181 := GenS6{F0: 2147483647, F1: 1}
                            for fi182 := 0; fi182 < 1; fi182++ {
                                for fi183 := 0; fi183 < 8; fi183++ {
                                    var v184 bool = (fi182 >= (-(fi183)))
                                    v185 := v2["foo bar"]
                                    _ = fi183
                                    _ = v184
                                    _ = v185
                                }
                                var v186 float32 = v1
                                v187 := []int{(fi182 << ((136) & 63)), fi182}
                                _ = fi182
                                _ = v186
                                _ = v187
                            }
                            var v188 float64 = 0.1
                            _ = v181
                            _ = v188
                        }
                    case 32768:
                        v1 -= v1
                    }
                    var v189 string = ("\n\t" + "hello")
                    if (!((!(96 > 100)) || v3)) {
                        if (((v3 || true) || v3) && (float64(3.14159) < (0.1 * -2.25))) {
                            switch -1 {
                            case 127:
                                for fi190 := 0; fi190 < 11; fi190++ {
                                    var v191 bool = v3
                                    _ = fi190
                                    _ = v191
                                }
                            case 0:
                                var v192 uint32 = 65535
                                _ = v192
                            case 880:
                                v189 = v189
                            }
                            var v193 int64 = -9223372036854775808
                            _ = v193
                        }
                    }
                    _ = v189
                }
            }
            v194 := v2["a"]
            _ = v194
        }
    }
    v195 := GenS7{F0: 65535, F1: -707}
    v196 := []int{-1, -936, -655}
    if (-1 > -484) {
        var v197 int = 32767
        v4 = v4
        _ = v197
    }
    var v198 uint16 = 255
    v2[(("hello" + "foo bar") + "\n\t")] = (("a" + "foo bar") + "hello")
    v198 -= v198
    if true {
        v5 += v5
        if (-2.25 > 3.14159) {
            v199 := GenS17{F0: 100}
            _ = v199
        }
    }
    fmt.Println(v3, v198, v1)
    if (!v3) {
        var v200 int32 = 42
        switch 32768 {
        case -2147483648:
            var v201 bool = (!v3)
            _ = v201
        case 42:
            if (1.5 >= 3.14159) {
                var v202 int = 32768
                var v203 float64 = -2.25
                switch (v202 >> ((65535) & 63)) {
                case -222:
                    fmt.Println(v198, v200, v1)
                case 274:
                    var v204 int = v202
                    _ = v204
                case -500:
                    switch ("foo bar" + "\n\t") {
                    case "":
                        for fi205 := 0; fi205 < 5; fi205++ {
                            fmt.Println(v198)
                            _ = fi205
                        }
                    }
                default:
                    var v206 int16 = -32768
                    _ = v206
                }
                _ = v202
                _ = v203
            }
        }
        v198 += (v198 / ((v198) | 1))
        _ = v200
    } else {
        for fi207 := 0; fi207 < 6; fi207++ {
            var v208 bool = true
            v209 := GenS12{F0: -2147483648, F1: (!(-349 >= 0))}
            v208 = (!v208)
            _ = fi207
            _ = v208
            _ = v209
        }
        switch ("\n\t" + "hello") {
        case "hello":
            switch 703 {
            case 4294967295:
                if (v3 || false) {
                    fmt.Println(v1, v1)
                }
            case -9223372036854775808:
                v210 := GenS18{F0: 873, F1: v198, F2: 1}
                _ = v210
            default:
                for fi211 := 0; fi211 < 10; fi211++ {
                    v4 -= ((3.14159 - 0.0) / v4)
                    delete(v2, "")
                    for fi212 := 0; fi212 < 18; fi212++ {
                        v213 := GenS11{F0: -2147483648}
                        var v214 uint32 = 2147483647
                        if (((!(!v3)) || true) || v3) {
                            v214 += genFn1(fi212, "a", genFn1(748, "hello", (v214 % ((v214) | 1))))
                        } else {
                            v214 += genFn1((fi211 & (fi211 << ((18446744073709551615) & 63))), ("foo bar" + "foo bar"), v214)
                            switch fi211 {
                            case 4294967295:
                                switch "a" {
                                case "a":
                                    var v215 int64 = int64(v214)
                                    _ = v215
                                case "":
                                    if (!(330 != (^(v214)))) {
                                        if v3 {
                                            var v216 uint32 = (v214 / ((1) | 1))
                                            v217 := GenS12{F0: 2147483647, F1: ((v4 == float32((1e10 / float64(3.14159)))) || v3)}
                                            _ = v216
                                            _ = v217
                                        }
                                    }
                                    v218 := GenS14{F0: -956, F1: 1}
                                    v198 = (v198 - 42)
                                    _ = v218
                                case "\n\t":
                                    fmt.Println(v214, v5, fi211)
                                default:
                                    switch "\n\t" {
                                    case "a":
                                        v219 := GenS16{F0: 1000, F1: uint64(v5)}
                                        _ = v219
                                    case "":
                                        if len(v196) >= 2 {
                                            v196 = v196[1:len(v196)]
                                        }
                                    default:
                                        var v220 bool = v3
                                        _ = v220
                                    }
                                }
                            case -995:
                                v5 = 128
                            }
                        }
                        _ = fi212
                        _ = v213
                        _ = v214
                    }
                    _ = fi211
                }
            }
        default:
            if ((float32((1.5 / 3.14159)) / 0.1) == (0.0 - v4)) {
                if v3 {
                    var v221 float32 = float32((float64(v1) + 0.0))
                    v221 -= 1e10
                    var v222 uint16 = (v198 >> ((1) & 63))
                    _ = v221
                    _ = v222
                } else {
                    switch "\n\t" {
                    case "foo bar":
                        var v223 uint64 = 0
                        _ = v223
                    case "a":
                        v198 = v198
                    }
                    fmt.Println(v5, v5, v1)
                    v1 = 0.1
                }
            }
        }
        v224 := GenS17{F0: 94}
        _ = v224
    }
    for fi225 := 0; fi225 < 12; fi225++ {
        switch (fi225 | -33) {
        case -467:
            switch (fi225 & (fi225 - (fi225 & (-(fi225))))) {
            case 1:
                var v226 int16 = -1
                _ = v226
            case 127:
                if (!true) {
                    var v227 float32 = v4
                    v5 = (v5 << ((60) & 63))
                    var v228 int16 = 817
                    _ = v227
                    _ = v228
                }
            default:
                v3 = true
            }
        default:
            if v3 {
                v4 -= (-2.25 / (v4 / float32(3.14159)))
                var v229 int16 = 32767
                var v230 int8 = genFn0((fi225 ^ fi225), ("a" + "hello"))
                _ = v229
                _ = v230
            }
        }
        v5 += v5
        _ = fi225
    }
    switch "a" {
    case "\n\t":
        if (-1 == 1) {
            var v231 uint8 = (v5 ^ uint8(v198))
            v198 -= 65535
            var v232 bool = (!v3)
            _ = v231
            _ = v232
        } else {
            fmt.Println(v4, v5, v198)
            v5 = (v5 | (^(v5)))
        }
    case "":
        _ = len(v196)
    default:
        v4 += v1
    }
    var v233 string = "a"
    for fi234 := 0; fi234 < 20; fi234++ {
        for fi235 := 0; fi235 < 15; fi235++ {
            v236 := GenS14{F0: 9223372036854775807, F1: 612}
            _ = fi235
            _ = v236
        }
        _ = fi234
    }
    if (!v3) {
        v233 = ""
        var v237 int16 = -1
        for fi238 := 0; fi238 < 3; fi238++ {
            _ = len(v196)
            v239 := GenS19{F0: 42, F1: (v198 | (v198 % (((v198 | v198)) | 1))), F2: 2}
            _ = fi238
            _ = v239
        }
        _ = v237
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v196:", len(v196), v196)
    fmt.Println("v198:", v198)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v233:", v233)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v195
    _ = v196
    _ = v198
    _ = v2
    _ = v233
    _ = v3
    _ = v4
    _ = v5
}
