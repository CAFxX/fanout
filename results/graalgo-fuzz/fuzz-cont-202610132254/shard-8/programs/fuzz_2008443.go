package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 uint; F1 int16 }
type GenS6 struct { F0 int64; F1 int }
type GenS7 struct { F0 int; F1 string; F2 int32 }
type GenS8 struct { F0 string; F1 bool }
type GenS9 struct { F0 bool; F1 int64 }
type GenS10 struct { F0 int; F1 uint32 }
type GenS11 struct { F0 uint16; F1 int; F2 string }
func genFn0(p0 int8) string {
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 float64 = 3.14159
        v8 := GenS1{F0: false}
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 bool = true
    p0 += p0
    _ = p0
    _ = v9
    return "hello"
}
func genFn1() int8 {
    if false {
        var v10 uint8 = 0
        fmt.Println(v10)
        _ = v10
    } else {
        var v11 int64 = -842
        v11 = -474
        _ = v11
    }
    var v12 float32 = 0.0
    var v13 int16 = 1000
    for fi14 := 0; fi14 < 16; fi14++ {
        for fi15 := 0; fi15 < 6; fi15++ {
            if false {
                switch -343 {
                case 2147483647:
                    fmt.Println(v13, fi14)
                case 195:
                    switch ("" + "hello") {
                    case "foo bar":
                        switch "a" {
                        case "hello":
                            switch fi15 {
                            case 42:
                                fmt.Println(fi14, fi15)
                            case 4294967295:
                                switch -1 {
                                case 32767:
                                    var v16 int64 = -1
                                    _ = v16
                                case 42:
                                    var v17 uint8 = 128
                                    _ = v17
                                }
                            case 32767:
                                switch fi14 {
                                case 32768:
                                    v18 := GenS2{F0: 42}
                                    _ = v18
                                case 42:
                                    v13 = (v13 & v13)
                                }
                            default:
                                var v19 uint32 = 679
                                _ = v19
                            }
                        case "foo bar":
                            v20 := []int{(fi14 - fi15), (-(fi14)), (fi14 / ((fi15) | 1))}
                            _ = v20
                        default:
                            v12 = 0.0
                        }
                    case "\n\t":
                        v13 += 42
                    default:
                        v21 := []int{fi15}
                        _ = v21
                    }
                case 127:
                    if (74 < 18446744073709551615) {
                        fmt.Println(v13, fi15, fi14)
                    }
                }
                var v22 int64 = 1
                v23 := GenS3{F0: 1, F1: "a"}
                _ = v22
                _ = v23
            } else {
                v12 -= 0.0
                v24 := GenS4{F0: "a", F1: 0}
                _ = v24
            }
            fmt.Println(fi15, v12)
            _ = fi15
        }
        _ = fi14
    }
    _ = v12
    _ = v13
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int = -2147483648
    var v3 uint16 = 1
    var v4 int16 = 32767
    var v5 uint = 12
    if false {
        v25 := GenS5{F0: (v5 << ((v5) & 63)), F1: (v4 + v4)}
        _ = v25
    } else {
        v26 := GenS2{F0: 0}
        _ = v26
    }
    v5 -= v5
    for fi27 := 0; fi27 < 14; fi27++ {
        var v28 int64 = 505
        _ = fi27
        _ = v28
    }
    v29 := GenS6{F0: -9223372036854775808, F1: (v2 & -846)}
    var v30 uint8 = 0
    var v31 int16 = (v4 / ((v4) | 1))
    if false {
        for fi32 := 0; fi32 < 7; fi32++ {
            v33 := []int{(fi32 + 128), (fi32 << ((18446744073709551615) & 63))}
            var v34 int64 = 394
            _ = fi32
            _ = v33
            _ = v34
        }
    } else {
        fmt.Println(v31, v3, v30)
        var v35 int8 = genFn1()
        _ = v35
    }
    v31 -= (v31 - (v4 | v4))
    v36 := GenS7{F0: v2, F1: ((((((((("" + "") + "foo bar") + "\n\t") + ("foo bar" + (((("\n\t" + "hello") + ("a" + ("foo bar" + ""))) + "") + "\n\t"))) + "") + ("" + (("a" + "foo bar") + ""))) + "\n\t") + "a") + "\n\t"), F2: -2147483648}
    var v37 bool = false
    var v38 uint64 = 690
    v39 := GenS4{F0: ("foo bar" + "hello"), F1: v30}
    v37 = v37
    switch v2 {
    case 32767:
        fmt.Println(v2, v5)
    case 1:
        if false {
            switch "" {
            case "":
                if false {
                    if true {
                        v40 := []int{4294967295, (v2 - v2), v2, (v2 | v2)}
                        fmt.Println(v3, v5)
                        fmt.Println(v5, v38, v30)
                        _ = v40
                    }
                } else {
                    switch v2 {
                    case 802:
                        v41 := GenS1{F0: (v38 > (^(v38)))}
                        _ = v41
                    case 32767:
                        fmt.Println(v2)
                    case 127:
                        v4 = v4
                    }
                    for fi42 := 0; fi42 < 9; fi42++ {
                        if (v5 < 983) {
                            v43 := GenS7{F0: (v2 & (v2 / ((fi42) | 1))), F1: ("" + "\n\t"), F2: 42}
                            for fi44 := 0; fi44 < 9; fi44++ {
                                var v45 int = fi42
                                _ = fi44
                                _ = v45
                            }
                            v46 := GenS6{F0: 653, F1: (-(fi42))}
                            _ = v43
                            _ = v46
                        }
                        if v37 {
                            switch "a" {
                            case "\n\t":
                                if (v37 && v37) {
                                    v47 := GenS7{F0: (fi42 - fi42), F1: "hello", F2: -2147483648}
                                    _ = v47
                                }
                            case "foo bar":
                                v5 = 42
                            default:
                                switch v2 {
                                case 1:
                                    if (!v37) {
                                        fmt.Println(v5, v3)
                                        if v37 {
                                            v48 := []int{int(v31), v2, fi42, (fi42 ^ -110)}
                                            fmt.Println(fi42, v30, fi42)
                                            _ = v48
                                        } else {
                                            v4 -= 0
                                        }
                                        if v37 {
                                            fmt.Println(v30, v31)
                                            v49 := []int{}
                                            _ = v49
                                        } else {
                                            fmt.Println(v38)
                                            v2 = (fi42 + (v2 & (fi42 + (v2 % ((fi42) | 1)))))
                                        }
                                    } else {
                                        var v50 uint32 = 65535
                                        switch ("\n\t" + "foo bar") {
                                        case "hello":
                                            v31 = -633
                                        case "":
                                            fmt.Println(v5, v30)
                                        }
                                        for fi51 := 0; fi51 < 17; fi51++ {
                                            switch fi51 {
                                            case 929:
                                                v1[(v50 * v50)] = true
                                            case -448:
                                                fmt.Println(v30)
                                            }
                                            _ = fi51
                                        }
                                        _ = v50
                                    }
                                case 9223372036854775807:
                                    var v52 int64 = 529
                                    _ = v52
                                case -9223372036854775808:
                                    fmt.Println(v2, fi42, v30)
                                default:
                                    var v53 string = "a"
                                    _ = v53
                                }
                            }
                        } else {
                            v54 := GenS8{F0: "a", F1: v37}
                            _ = v54
                        }
                        _ = fi42
                    }
                }
            }
            for fi55 := 0; fi55 < 6; fi55++ {
                v31 -= (v31 | (-(v31)))
                var v56 float64 = (float64((3.14159 / 1.0)) + 1e10)
                if v37 {
                    var v57 int8 = int8(v4)
                    _ = v57
                } else {
                    var v58 int32 = -1
                    _ = v58
                }
                _ = fi55
                _ = v56
            }
            fmt.Println(v3, v5, v38)
        } else {
            var v59 uint8 = v30
            v60 := []int{v2, (v2 & v2), (v2 / ((v2) | 1)), (v2 << ((v5) & 63))}
            v61 := GenS9{F0: (v59 > uint8(v5)), F1: 9223372036854775807}
            _ = v59
            _ = v60
            _ = v61
        }
    case 9223372036854775807:
        v31 += (^(v31))
    default:
        v3 = v3
    }
    v1[729] = (!v37)
    for fi62 := 0; fi62 < 16; fi62++ {
        if v37 {
            for fi63 := 0; fi63 < 8; fi63++ {
                var v64 int = fi63
                switch 42 {
                case 127:
                    switch 32767 {
                    case 2147483647:
                        v3 = (^(v3))
                    case -534:
                        v31 += (v31 - 1)
                        var v65 string = "a"
                        _ = v65
                    case 0:
                        v66 := GenS8{F0: "foo bar", F1: false}
                        _ = v66
                    }
                    var v67 float64 = ((0.1 * 0.1) + 1.5)
                    v31 = (v4 ^ v31)
                    _ = v67
                case -9223372036854775808:
                    v68 := []int{-422, (fi63 % ((fi63) | 1)), v64, -2147483648}
                    _ = v68
                case 32768:
                    if (((!v37) || false) || true) {
                        if false {
                            if v37 {
                                v69 := GenS9{F0: ((v38 >> ((v5) & 63)) <= 18446744073709551615), F1: -1}
                                if (!(-2.25 > (((0.0 / 1e10) - -2.25) * (1e10 / 0.1)))) {
                                    var v70 int64 = -9223372036854775808
                                    _ = v70
                                }
                                v71 := v1[0]
                                v64 += v2
                                fmt.Println(v4, fi62)
                                _ = v69
                                _ = v71
                            } else {
                                v38 += (v38 * (v38 + 1))
                                var v72 string = ("" + "\n\t")
                                v73 := GenS5{F0: (v5 * (v5 >> (((v5 << (((v5 % ((v5) | 1))) & 63))) & 63))), F1: -214}
                                var v74 int8 = -1
                                _ = v72
                                _ = v73
                                _ = v74
                            }
                        }
                        var v75 int = (fi63 & -67)
                        var v76 string = "foo bar"
                        var v77 uint32 = 4294967295
                        _ = v75
                        _ = v76
                        _ = v77
                    } else {
                        switch 163 {
                        case 32768:
                            var v78 bool = ((!v37) && ((v37 && v37) && (0.1 >= 1.5)))
                            _ = v78
                        }
                        switch fi63 {
                        case 1:
                            v64 = v64
                        case -240:
                            var v79 int32 = 714
                            _ = v79
                        case 0:
                            if v37 {
                                if v37 {
                                    v80 := GenS7{F0: -9223372036854775808, F1: "", F2: -1}
                                    _ = v80
                                } else {
                                    var v81 string = (("foo bar" + ((("a" + "") + ((("a" + "foo bar") + "") + (("\n\t" + "foo bar") + ""))) + "")) + "a")
                                    if v37 {
                                        fmt.Println(v37, v37)
                                    } else {
                                        v82 := []int{int(v31)}
                                        _ = v82
                                    }
                                    if v37 {
                                        switch ("hello" + v81) {
                                        case "a":
                                            var v83 bool = (!(float32(1e10) < (1e10 + (0.0 / (float32((float64(float32(-2.25)) / ((0.1 * 3.14159) * 3.14159))) * float32((((0.1 + -2.25) + (1.5 * -2.25)) / 1.5)))))))
                                            _ = v83
                                        case "hello":
                                            fmt.Println(v5)
                                        }
                                    }
                                    var v84 float32 = (0.1 / 1e10)
                                    _ = v81
                                    _ = v84
                                }
                                fmt.Println(v3)
                                var v85 float64 = -2.25
                                _ = v85
                            } else {
                                _ = len(v1)
                                v86 := []int{127, (fi62 | fi63), 9223372036854775807}
                                v86 = append(v86, (v64 % ((9223372036854775807) | 1)))
                                if ((!(!(127 != 45))) || (false || (v5 <= v5))) {
                                    var v87 float32 = 1e10
                                    v4 = -71
                                    _ = v87
                                } else {
                                    var v88 string = (("\n\t" + "\n\t") + "")
                                    if len(v86) >= 2 {
                                        v86 = v86[1:len(v86)]
                                    }
                                    _ = v88
                                }
                                if (((3.14159 <= (0.1 - (1.5 + 1.5))) && (!true)) && (1 == 4294967295)) {
                                    switch 0 {
                                    case 32767:
                                        switch fi62 {
                                        case 861:
                                            v86 = append(v86, v64)
                                        case -2147483648:
                                            v38 = v38
                                        case 42:
                                            switch int(v31) {
                                            case 2147483647:
                                                if v37 {
                                                    if true {
                                                        var v89 uint32 = 65535
                                                        if (623 != -1) {
                                                            var v90 int8 = int8(v38)
                                                            if (978 != (^(v4))) {
                                                                if (v3 != uint16(v5)) {
                                                                    switch (^(fi62)) {
                                                                    case 42:
                                                                        var v91 uint = 18446744073709551615
                                                                        _ = v91
                                                                    }
                                                                    v37 = ((v30 << (((v5 % ((v5) | 1))) & 63)) > v30)
                                                                    v31 -= 1
                                                                } else {
                                                                    fmt.Println(v4)
                                                                    v92 := GenS6{F0: 1, F1: fi62}
                                                                    if v37 {
                                                                        v64 -= (fi63 << ((v5) & 63))
                                                                    } else {
                                                                        v86 = append(v86, (v64 % ((125) | 1)))
                                                                        switch v2 {
                                                                        case -1:
                                                                            var v93 int8 = v90
                                                                            _ = v93
                                                                        }
                                                                        var v94 int = int(v30)
                                                                        _ = v94
                                                                    }
                                                                    _ = v92
                                                                }
                                                                v90 = (v90 << ((65535) & 63))
                                                                v31 += -32768
                                                                if len(v86) > 0 {
                                                                    v86[0] = (v2 | int(v89))
                                                                }
                                                                v89 -= (v89 >> (((^(v5))) & 63))
                                                                v64 -= 128
                                                            } else {
                                                                fmt.Println(v89)
                                                                if (!(!v37)) {
                                                                    fmt.Println(v89)
                                                                    fmt.Println(v38)
                                                                    var v95 int16 = (v31 << ((v5) & 63))
                                                                    _ = v95
                                                                }
                                                            }
                                                            _ = v90
                                                        } else {
                                                            v5 = uint(v38)
                                                            v96 := GenS1{F0: (0.0 <= 0.1)}
                                                            _ = v96
                                                        }
                                                        if ((fi62 << ((v5) & 63)) < (v2 % ((fi63) | 1))) {
                                                            switch 32768 {
                                                            case 128:
                                                                if (v5 != (v5 + 18446744073709551615)) {
                                                                    if (!v37) {
                                                                        if ((!(100 == 100)) && v37) {
                                                                            v86 = append(v86, (-(fi62)))
                                                                            v64 = v64
                                                                        } else {
                                                                            v86 = append(v86, (fi63 & int(v38)))
                                                                        }
                                                                    } else {
                                                                        fmt.Println(v31)
                                                                        switch (-(v64)) {
                                                                        case 28:
                                                                            v31 = (v4 - (v31 << ((v5) & 63)))
                                                                        case 9223372036854775807:
                                                                            var v97 uint8 = v30
                                                                            _ = v97
                                                                        }
                                                                        v31 = (v31 << (((v5 / (((^(v5))) | 1))) & 63))
                                                                    }
                                                                }
                                                            default:
                                                                if v37 {
                                                                    v37 = v37
                                                                    fmt.Println(fi63, v37)
                                                                }
                                                            }
                                                            var v98 int64 = 9223372036854775807
                                                            if (3.14159 > 3.14159) {
                                                                fmt.Println(v30, v31, fi62)
                                                                var v99 bool = true
                                                                fmt.Println(v89)
                                                                if v37 {
                                                                    switch 0 {
                                                                    case -52:
                                                                        if (1 < (-(v31))) {
                                                                            var v100 float32 = 0.1
                                                                            switch -2147483648 {
                                                                            case -27:
                                                                                if ((^(v5)) != uint(v3)) {
                                                                                    switch -422 {
                                                                                    case 127:
                                                                                        v37 = (!true)
                                                                                    case -9223372036854775808:
                                                                                        v99 = (!(v38 >= 42))
                                                                                    }
                                                                                    if false {
                                                                                        var v101 float32 = v100
                                                                                        var v102 int32 = 42
                                                                                        _ = v101
                                                                                        _ = v102
                                                                                    }
                                                                                }
                                                                            case 32768:
                                                                                v103 := v1[(v89 + 4294967295)]
                                                                                _ = v103
                                                                            }
                                                                            _ = v100
                                                                        }
                                                                    }
                                                                    v5 += (v5 * 915)
                                                                    if len(v86) > 0 {
                                                                        v86[0] = (^(fi62))
                                                                    }
                                                                    var v104 int64 = (-(v98))
                                                                    _ = v104
                                                                } else {
                                                                    var v105 bool = v99
                                                                    _ = len(v86)
                                                                    v106 := GenS8{F0: "a", F1: v99}
                                                                    _ = v105
                                                                    _ = v106
                                                                }
                                                                _ = v99
                                                            }
                                                            _ = v98
                                                        } else {
                                                            v86 = append(v86, (v64 >> (((v5 | (v5 & v5))) & 63)))
                                                            switch (-(v2)) {
                                                            case -899:
                                                                v86 = append(v86, (v2 >> ((uint(v89)) & 63)))
                                                            case 243:
                                                                var v107 uint16 = v3
                                                                _ = v107
                                                            case 9223372036854775807:
                                                                var v108 int8 = int8(v89)
                                                                _ = v108
                                                            default:
                                                                if len(v86) > 0 {
                                                                    v86[0] = v2
                                                                }
                                                            }
                                                            fmt.Println(fi63, v89, v37)
                                                        }
                                                        _ = v89
                                                    } else {
                                                        if v37 {
                                                            if (v3 > 188) {
                                                                v109 := GenS10{F0: -846, F1: 65535}
                                                                v4 = 569
                                                                _ = v109
                                                            } else {
                                                                fmt.Println(v30, v31)
                                                            }
                                                            if (v37 || v37) {
                                                                v110 := GenS9{F0: (!v37), F1: 9223372036854775807}
                                                                var v111 uint32 = 2147483648
                                                                _ = v110
                                                                _ = v111
                                                            }
                                                        }
                                                    }
                                                }
                                            }
                                        }
                                    default:
                                        var v112 uint32 = 2147483648
                                        _ = v112
                                    }
                                }
                                v86 = append(v86, v2)
                                fmt.Println(v37, v4)
                                _ = v86
                            }
                        default:
                            if ((127 == (-(fi63))) && v37) {
                                var v113 uint = 18446744073709551615
                                var v114 uint = (v5 / ((v113) | 1))
                                _ = v113
                                _ = v114
                            } else {
                                v4 = v4
                            }
                        }
                        var v115 int8 = genFn1()
                        _ = v115
                    }
                default:
                    v3 = (v3 ^ 65535)
                }
                v116 := GenS5{F0: (v5 & 869), F1: (v4 / (((v31 | (v31 * (-(v4))))) | 1))}
                var v117 uint = v5
                _ = fi63
                _ = v116
                _ = v117
                _ = v64
            }
            for fi118 := 0; fi118 < 14; fi118++ {
                _ = len(v1)
                _ = fi118
            }
            switch (("hello" + (("a" + "a") + "foo bar")) + ("\n\t" + "foo bar")) {
            case "":
                for fi119 := 0; fi119 < 9; fi119++ {
                    var v120 string = "foo bar"
                    v3 = v3
                    if (!true) {
                        v120 = ("\n\t" + ((v120 + v120) + v120))
                    } else {
                        var v121 uint = v5
                        var v122 int32 = 0
                        var v123 int64 = -442
                        _ = v121
                        _ = v122
                        _ = v123
                    }
                    _ = fi119
                    _ = v120
                }
            default:
                fmt.Println(v3, v31, v38)
            }
        }
        switch "foo bar" {
        case "foo bar":
            if (!(v37 || false)) {
                var v124 bool = v37
                v125 := GenS6{F0: 463, F1: -657}
                _ = v124
                _ = v125
            } else {
                fmt.Println(fi62)
            }
        case "hello":
            for fi126 := 0; fi126 < 14; fi126++ {
                v127 := GenS11{F0: 65535, F1: fi62, F2: ("foo bar" + "foo bar")}
                v128 := GenS9{F0: (-1 < (v2 & 9223372036854775807)), F1: -1}
                v5 += v5
                _ = fi126
                _ = v127
                _ = v128
            }
        case "\n\t":
            v30 += (v30 >> (((^(v5))) & 63))
        }
        for fi129 := 0; fi129 < 16; fi129++ {
            v130 := v1[65535]
            v131 := GenS10{F0: 2147483647, F1: 0}
            v132 := []int{(v2 | fi62), 42}
            v30 -= (^(v30))
            var v133 int16 = (v31 * v31)
            _ = fi129
            _ = v130
            _ = v131
            _ = v132
            _ = v133
        }
        _ = fi62
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
