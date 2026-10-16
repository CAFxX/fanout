package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 int16 }
type GenS5 struct { F0 int; F1 uint32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint32; F1 bool }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint64; F1 uint }
type GenS11 struct { F0 int64; F1 int16 }
type GenS12 struct { F0 uint32; F1 int8; F2 bool }
type GenS13 struct { F0 uint8; F1 int64; F2 bool }
type GenS14 struct { F0 int32 }
func genFn0(p0 bool, p1 uint64, p2 uint) uint16 {
    p0 = (float32(0.0) == 3.14159)
    var v6 uint32 = 1
    for fi7 := 0; fi7 < 11; fi7++ {
        if p0 {
            fmt.Println(p2, v6)
            for fi8 := 0; fi8 < 15; fi8++ {
                v6 = (v6 >> ((p2) & 63))
                _ = fi8
            }
        }
        var v9 uint = (p2 ^ 42)
        _ = fi7
        _ = v9
    }
    var v10 int32 = 42
    v11 := GenS1{F0: 127}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v6
    return 155
}
func genFn1(p0 uint, p1 uint32) int {
    p1 = (p1 & p1)
    if false {
        p1 = 976
        if true {
            for fi12 := 0; fi12 < 10; fi12++ {
                for fi13 := 0; fi13 < 11; fi13++ {
                    switch "" {
                    case "hello":
                        if false {
                            p1 = (p1 / ((p1) | 1))
                        }
                    default:
                        var v14 int16 = 32767
                        _ = v14
                    }
                    _ = fi13
                }
                switch fi12 {
                case 4294967295:
                    p1 += uint32(p0)
                case 127:
                    var v15 string = "\n\t"
                    _ = v15
                default:
                    if true {
                        fmt.Println(fi12, p1)
                        v16 := []int{(fi12 / (((^(fi12))) | 1))}
                        _ = v16
                    } else {
                        if (p0 != 42) {
                            switch 127 {
                            case -752:
                                p0 -= 2
                            case 127:
                                var v17 uint8 = 42
                                _ = v17
                            case 32768:
                                v18 := []int{fi12, (fi12 % ((fi12) | 1))}
                                _ = v18
                            }
                            p1 -= p1
                        }
                        for fi19 := 0; fi19 < 8; fi19++ {
                            v20 := GenS1{F0: 255}
                            _ = fi19
                            _ = v20
                        }
                    }
                }
                _ = fi12
            }
            var v21 uint8 = 0
            _ = v21
        } else {
            for fi22 := 0; fi22 < 20; fi22++ {
                v23 := GenS2{F0: "foo bar", F1: fi22}
                v24 := GenS2{F0: ("\n\t" + "a"), F1: fi22}
                _ = fi22
                _ = v23
                _ = v24
            }
        }
        if (!(0.0 == 0.1)) {
            var v25 bool = true
            if true {
                var v26 uint64 = 18446744073709551615
                v27 := []int{32768}
                _ = v26
                _ = v27
            } else {
                var v28 string = "\n\t"
                _ = v28
            }
            _ = v25
        }
    }
    switch -1 {
    case 127:
        p1 += p1
    case 171:
        p0 = p0
    case 42:
        p1 += p1
    default:
        v29 := []int{42, 128, 0, -1}
        _ = v29
    }
    var v30 int16 = 1000
    for fi31 := 0; fi31 < 20; fi31++ {
        v32 := GenS3{F0: "\n\t"}
        v33 := GenS4{F0: (65535 != 440), F1: 32767}
        fmt.Println(fi31)
        _ = fi31
        _ = v32
        _ = v33
    }
    _ = p0
    _ = p1
    _ = v30
    return 4294967295
}

func main() {
    var v1 string = ("a" + "a")
    v2 := make(map[int]string)
    v3 := []int{32767}
    var v4 uint16 = 1
    v5 := make(map[int]int)
    for fi34 := 0; fi34 < 8; fi34++ {
        v35 := GenS5{F0: fi34, F1: 1}
        v1 = "a"
        var v36 string = ""
        _ = fi34
        _ = v35
        _ = v36
    }
    var v37 uint32 = 0
    for fi38 := 0; fi38 < 9; fi38++ {
        if (!false) {
            switch genFn1(255, 4294967295) {
            case -509:
                for fi39 := 0; fi39 < 14; fi39++ {
                    var v40 uint = 547
                    v41 := GenS2{F0: ((v1 + v1) + (("hello" + v1) + ("foo bar" + ""))), F1: (fi39 >> ((935) & 63))}
                    v40 = (v40 - v40)
                    _ = fi39
                    _ = v40
                    _ = v41
                }
            case -864:
                if false {
                    for fi42 := 0; fi42 < 17; fi42++ {
                        v2[fi42] = v1
                        if (1 < 97) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            if true {
                                var v43 int32 = 2147483647
                                switch int(v4) {
                                case 943:
                                    switch "" {
                                    case "":
                                        switch v1 {
                                        case "foo bar":
                                            v44 := GenS3{F0: "hello"}
                                            _ = v44
                                        }
                                    default:
                                        var v45 uint = 1
                                        _ = v45
                                    }
                                case 444:
                                    if false {
                                        var v46 int = genFn1(492, (^(v37)))
                                        v1 = ("\n\t" + v1)
                                        var v47 uint32 = 188
                                        var v48 int32 = v43
                                        _ = v46
                                        _ = v47
                                        _ = v48
                                    }
                                }
                                var v49 int8 = 42
                                _ = v43
                                _ = v49
                            } else {
                                if true {
                                    if len(v3) > 0 {
                                        v3[0] = (fi42 - (fi38 | fi38))
                                    }
                                    switch "foo bar" {
                                    case "hello":
                                        switch v1 {
                                        case "hello":
                                            v4 += (v4 - (v4 << ((2) & 63)))
                                        case "a":
                                            var v50 string = v1
                                            _ = v50
                                        case "\n\t":
                                            v1 = ("foo bar" + v1)
                                        default:
                                            if ((v37 & v37) != v37) {
                                                var v51 int8 = -25
                                                _ = len(v3)
                                                v51 = 0
                                                _ = v51
                                            }
                                        }
                                    case "foo bar":
                                        switch v1 {
                                        case "foo bar":
                                            v4 -= v4
                                        case "":
                                            v52 := GenS6{F0: (32767 >= 0)}
                                            _ = v52
                                        case "hello":
                                            v2[fi42] = (v1 + v1)
                                        default:
                                            v4 -= v4
                                        }
                                    case "a":
                                        if (!(((^(v4)) != (v4 >> ((0) & 63))) || true)) {
                                            fmt.Println(v4, v1, fi38)
                                        } else {
                                            if (v1 == (((((v1 + v1) + v1) + v1) + v1) + v1)) {
                                                v4 = 0
                                                v1 = (v1 + v1)
                                            } else {
                                                if (false || false) {
                                                    if (42 != 2) {
                                                        v4 = 356
                                                    } else {
                                                        v53 := GenS7{F0: 0, F1: (!true)}
                                                        switch (^(fi38)) {
                                                        case -809:
                                                            var v54 float32 = (float32(0.1) - 1e10)
                                                            _ = v54
                                                        case 32767:
                                                            fmt.Println(fi38, v1)
                                                        }
                                                        _ = v53
                                                    }
                                                    if true {
                                                        v5[(fi42 % (((-(fi38))) | 1))] = fi38
                                                        v5[(fi38 + fi42)] = -971
                                                        if (!(!((1 < -168) && (v4 <= (v4 - (v4 << ((0) & 63))))))) {
                                                            v55 := GenS8{F0: "a"}
                                                            v1 = (v1 + ("" + (v1 + (v1 + ("foo bar" + (v1 + v1))))))
                                                            var v56 uint16 = (v4 + genFn0((3.14159 <= (0.1 / (-2.25 * float64((0.1 - float32((1e10 - 1.5))))))), 823, 18446744073709551615))
                                                            v57 := GenS8{F0: (("a" + v1) + (v1 + v1))}
                                                            v58 := v2[(fi38 >> ((65535) & 63))]
                                                            switch v1 {
                                                            case "hello":
                                                                v4 = (v56 | v4)
                                                            case "\n\t":
                                                                switch fi38 {
                                                                case 32768:
                                                                    fmt.Println(v37)
                                                                case 0:
                                                                    v56 += genFn0((-13 == -1), 42, 65535)
                                                                }
                                                            case "foo bar":
                                                                var v59 uint8 = 255
                                                                _ = v59
                                                            }
                                                            var v60 uint32 = (v37 + v37)
                                                            _ = v55
                                                            _ = v56
                                                            _ = v57
                                                            _ = v58
                                                            _ = v60
                                                        } else {
                                                            v1 = ("foo bar" + "hello")
                                                        }
                                                    }
                                                    if false {
                                                        fmt.Println(v1)
                                                        switch v1 {
                                                        case "\n\t":
                                                            var v61 uint16 = (v4 | (v4 << ((837) & 63)))
                                                            _ = v61
                                                        default:
                                                            var v62 string = v1
                                                            _ = v62
                                                        }
                                                        switch v1 {
                                                        case "":
                                                            v63 := GenS9{F0: 1}
                                                            _ = v63
                                                        case "a":
                                                            var v64 int32 = -530
                                                            _ = v64
                                                        case "foo bar":
                                                            switch (fi38 << ((42) & 63)) {
                                                            case -822:
                                                                var v65 uint64 = 1
                                                                _ = v65
                                                            case 4294967295:
                                                                v66 := GenS9{F0: genFn1(126, (v37 - (v37 * v37)))}
                                                                _ = v66
                                                            case 1:
                                                                if true {
                                                                    v1 = (("" + (((((("hello" + (v1 + (((v1 + v1) + "\n\t") + ("foo bar" + v1)))) + "hello") + "\n\t") + (v1 + (("" + ("\n\t" + ((v1 + v1) + "a"))) + ("foo bar" + v1)))) + v1) + ("\n\t" + v1))) + "hello")
                                                                    switch v1 {
                                                                    case "foo bar":
                                                                        v37 = 65535
                                                                    case "":
                                                                        switch fi38 {
                                                                        case -312:
                                                                            v67 := GenS10{F0: 987, F1: uint(v4)}
                                                                            _ = v67
                                                                        case -509:
                                                                            if true {
                                                                                fmt.Println(v4)
                                                                                fmt.Println(fi38, fi42)
                                                                            } else {
                                                                                v3 = append(v3, fi38)
                                                                                v37 += (^(v37))
                                                                            }
                                                                        case 42:
                                                                            var v68 uint = uint(v37)
                                                                            _ = v68
                                                                        }
                                                                    case "\n\t":
                                                                        if true {
                                                                            var v69 uint32 = v37
                                                                            v4 = (v4 & (^(v4)))
                                                                            _ = v69
                                                                        } else {
                                                                            v1 = (v1 + v1)
                                                                            if true {
                                                                                _ = len(v5)
                                                                                switch 127 {
                                                                                case 0:
                                                                                    v70 := GenS4{F0: true, F1: -1}
                                                                                    _ = v70
                                                                                case 146:
                                                                                    v1 = v1
                                                                                case 2147483647:
                                                                                    var v71 string = (((v1 + v1) + "") + ("\n\t" + v1))
                                                                                    _ = v71
                                                                                default:
                                                                                    var v72 int8 = 0
                                                                                    _ = v72
                                                                                }
                                                                                if ((1 == 113) && true) {
                                                                                    fmt.Println(v1)
                                                                                    var v73 uint = 937
                                                                                    _ = v73
                                                                                }
                                                                                v74 := GenS4{F0: false, F1: 183}
                                                                                v75 := GenS5{F0: fi38, F1: (v37 ^ (v37 + v37))}
                                                                                _ = v74
                                                                                _ = v75
                                                                            }
                                                                            if false {
                                                                                if (0.0 <= 3.14159) {
                                                                                    v76 := GenS4{F0: false, F1: 32767}
                                                                                    fmt.Println(v1)
                                                                                    switch (v1 + (v1 + "a")) {
                                                                                    case "\n\t":
                                                                                        v1 = v1
                                                                                        fmt.Println(v37)
                                                                                        if (!true) {
                                                                                            if true {
                                                                                                var v77 uint16 = (v4 - (^(v4)))
                                                                                                _ = v77
                                                                                            } else {
                                                                                                switch ("foo bar" + (v1 + (v1 + (v1 + v1)))) {
                                                                                                case "\n\t":
                                                                                                    v4 += genFn0((1 == 42), 751, 0)
                                                                                                default:
                                                                                                    v78 := v5[genFn1(604, (v37 - (^(v37))))]
                                                                                                    _ = v78
                                                                                                }
                                                                                            }
                                                                                            if (4294967295 == 907) {
                                                                                                if len(v3) >= 2 {
                                                                                                    v3 = v3[1:len(v3)]
                                                                                                }
                                                                                                v37 = v37
                                                                                                if len(v3) > 0 {
                                                                                                    v3[0] = (fi42 * -1)
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            v79 := GenS7{F0: 571, F1: (42 <= 1000)}
                                                                                            _ = len(v2)
                                                                                            switch fi38 {
                                                                                            case 128:
                                                                                                switch v1 {
                                                                                                case "a":
                                                                                                    v80 := GenS1{F0: 127}
                                                                                                    _ = v80
                                                                                                case "":
                                                                                                    if (!true) {
                                                                                                        if len(v3) > 0 {
                                                                                                            v3[0] = fi42
                                                                                                        }
                                                                                                        fmt.Println(fi38)
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            fmt.Println(fi38)
                                                                                            _ = v79
                                                                                        }
                                                                                    case "":
                                                                                        if len(v3) >= 2 {
                                                                                            v3 = v3[1:len(v3)]
                                                                                        }
                                                                                    case "a":
                                                                                        if len(v3) >= 2 {
                                                                                            v3 = v3[1:len(v3)]
                                                                                        }
                                                                                    }
                                                                                    _ = v76
                                                                                } else {
                                                                                    v1 = "a"
                                                                                    v1 = "\n\t"
                                                                                }
                                                                                if (false && (-9223372036854775808 >= -9223372036854775808)) {
                                                                                    delete(v5, fi42)
                                                                                    var v81 float64 = 1e10
                                                                                    _ = v81
                                                                                } else {
                                                                                    fmt.Println(v37)
                                                                                    var v82 float32 = ((3.14159 / 0.1) - 0.1)
                                                                                    v4 += 393
                                                                                    _ = v82
                                                                                }
                                                                                switch (fi38 / (((fi38 / (((fi42 >> ((1) & 63))) | 1))) | 1)) {
                                                                                case 32767:
                                                                                    fmt.Println(fi42, fi38, fi42)
                                                                                }
                                                                                v1 = v1
                                                                                if true {
                                                                                    var v83 uint = 255
                                                                                    var v84 uint16 = (v4 % ((v4) | 1))
                                                                                    var v85 int32 = -193
                                                                                    _ = v83
                                                                                    _ = v84
                                                                                    _ = v85
                                                                                } else {
                                                                                    var v86 float32 = (1.5 - 3.14159)
                                                                                    _ = v86
                                                                                }
                                                                            } else {
                                                                                v4 += genFn0(false, 4294967295, 2)
                                                                                v37 = (v37 + 905)
                                                                                v87 := GenS2{F0: v1, F1: (fi42 * (fi38 >> ((18446744073709551615) & 63)))}
                                                                                _ = v87
                                                                            }
                                                                        }
                                                                        if (100 > -1) {
                                                                            if (-2.25 != 0.0) {
                                                                                fmt.Println(v1, fi42)
                                                                                v37 += v37
                                                                            } else {
                                                                                switch (fi42 & fi42) {
                                                                                case 32767:
                                                                                    v88 := v5[fi38]
                                                                                    _ = v88
                                                                                case 128:
                                                                                    if true {
                                                                                        v89 := GenS11{F0: 728, F1: -32768}
                                                                                        _ = v89
                                                                                    }
                                                                                case -2147483648:
                                                                                    fmt.Println(v1, fi42)
                                                                                default:
                                                                                    var v90 uint64 = 4294967295
                                                                                    _ = v90
                                                                                }
                                                                            }
                                                                            v3 = append(v3, (fi38 / ((fi42) | 1)))
                                                                            if true {
                                                                                v91 := v2[(^(fi38))]
                                                                                v91 = v1
                                                                                var v92 uint8 = 2
                                                                                _ = v91
                                                                                _ = v92
                                                                            }
                                                                        } else {
                                                                            if true {
                                                                                v93 := GenS2{F0: v1, F1: (fi38 & fi38)}
                                                                                if true {
                                                                                    var v94 bool = (2147483647 >= 1)
                                                                                    _ = v94
                                                                                } else {
                                                                                    var v95 int32 = 42
                                                                                    _ = v95
                                                                                }
                                                                                _ = v93
                                                                            }
                                                                        }
                                                                    }
                                                                    v4 = (v4 ^ v4)
                                                                } else {
                                                                    var v96 uint = 951
                                                                    v1 = v1
                                                                    var v97 int = (-(fi38))
                                                                    _ = v96
                                                                    _ = v97
                                                                }
                                                            default:
                                                                var v98 bool = true
                                                                _ = v98
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    v99 := GenS5{F0: int(v4), F1: v37}
                                                    v37 = (v37 / ((2147483648) | 1))
                                                    _ = v99
                                                }
                                                v4 += (v4 | (v4 ^ v4))
                                            }
                                            var v100 int = fi42
                                            var v101 float32 = float32(float64(1e10))
                                            _ = v100
                                            _ = v101
                                        }
                                    default:
                                        v4 = 196
                                    }
                                } else {
                                    v102 := GenS12{F0: 0, F1: 100, F2: (-2.25 > (((((0.1 / 1.0) * 1.5) + 3.14159) / 0.1) - (-2.25 - (1.5 * -2.25))))}
                                    _ = v102
                                }
                                if false {
                                    if (true || ((v37 & v37) < v37)) {
                                        v103 := v5[fi38]
                                        v37 += 555
                                        _ = v103
                                    }
                                    v104 := GenS13{F0: 255, F1: 1, F2: false}
                                    _ = v104
                                } else {
                                    var v105 uint = uint(v4)
                                    var v106 uint = uint(v4)
                                    v106 = (v105 ^ 42)
                                    switch genFn1(308, 2147483648) {
                                    case 4294967295:
                                        v107 := GenS9{F0: (fi42 / ((fi38) | 1))}
                                        _ = v107
                                    }
                                    v105 += (v106 - (^(v105)))
                                    var v108 int8 = 100
                                    _ = v105
                                    _ = v106
                                    _ = v108
                                }
                            }
                            v109 := GenS7{F0: (v37 * (v37 & (v37 | 1))), F1: true}
                            _ = v109
                        } else {
                            _ = len(v2)
                            v37 += v37
                            var v110 uint64 = 454
                            _ = v110
                        }
                        _ = fi42
                    }
                    fmt.Println(fi38, fi38)
                    v1 = (v1 + "")
                } else {
                    v37 -= v37
                    fmt.Println(v37, v37)
                }
            case -2147483648:
                var v111 uint = 255
                _ = v111
            default:
                var v112 uint8 = 44
                _ = v112
            }
            v113 := GenS5{F0: fi38, F1: (v37 / ((v37) | 1))}
            _ = v113
        } else {
            v3 = append(v3, (fi38 + (-(fi38))))
            switch "a" {
            case "a":
                v114 := GenS14{F0: -1}
                _ = v114
            case "":
                var v115 float64 = 1e10
                _ = v115
            }
        }
        if len(v3) > 0 {
            v3[0] = genFn1(2, (v37 % ((65535) | 1)))
        }
        _ = fi38
    }
    fmt.Println(v1, v4)
    switch ("\n\t" + (v1 + "a")) {
    case "\n\t":
        var v116 int8 = 127
        _ = v116
    }
    v117 := v5[-9223372036854775808]
    v1 = v1
    v117 = (-(v117))
    var v118 float64 = (float64(1.5) - 3.14159)
    v2[v117] = v1
    v119 := GenS8{F0: v1}
    var v120 uint64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v117:", v117)
    fmt.Println("v118:", v118)
    fmt.Println("v120:", v120)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v117
    _ = v118
    _ = v119
    _ = v120
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
