package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 int8; F1 int8; F2 uint }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint; F1 int }
type GenS8 struct { F0 uint64; F1 uint32 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int32; F1 uint16 }
type GenS12 struct { F0 bool; F1 uint16 }
type GenS13 struct { F0 string; F1 int8; F2 uint16 }
type GenS14 struct { F0 int }
func genFn0() uint {
    var v7 int32 = -56
    switch "a" {
    case "":
        v7 -= (v7 * (^(v7)))
    case "\n\t":
        v7 -= 0
    case "a":
        var v8 int = 32768
        _ = v8
    }
    var v9 int32 = (v7 / ((v7) | 1))
    _ = v7
    _ = v9
    return 1
}
func genFn1(p0 uint64, p1 uint, p2 int16) uint8 {
    p0 += (^(p0))
    switch "" {
    case "\n\t":
        p1 += (p1 / ((255) | 1))
    case "a":
        p1 = (p1 << ((989) & 63))
    case "foo bar":
        if (false && true) {
            for fi10 := 0; fi10 < 3; fi10++ {
                p1 = 762
                var v11 uint8 = 205
                _ = fi10
                _ = v11
            }
        }
    }
    v12 := GenS1{F0: 255}
    var v13 uint32 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int16 = 890
    var v3 float64 = 0.1
    var v4 int32 = 2147483647
    v5 := make(map[uint32]uint32)
    v6 := []int{127, 4294967295, 0}
    _ = len(v6)
    if len(v6) > 0 {
        v6[0] = 128
    }
    v3 = (v3 + v3)
    switch "\n\t" {
    case "\n\t":
        for fi14 := 0; fi14 < 1; fi14++ {
            v3 -= ((1.5 / v3) + (v3 + (((float64(float32(v3)) + v3) + 3.14159) - float64(((1e10 - 0.1) * float32((float64(0.0) + -2.25)))))))
            var v15 int64 = -9223372036854775808
            if false {
                v15 += (^(v15))
                var v16 int64 = (^(v15))
                _ = v16
            } else {
                v3 = ((v3 + (float64(float32(-2.25)) - v3)) / 1e10)
                _ = len(v1)
                fmt.Println(v4)
            }
            _ = fi14
            _ = v15
        }
    default:
        for fi17 := 0; fi17 < 9; fi17++ {
            if false {
                v18 := GenS1{F0: 25}
                for fi19 := 0; fi19 < 11; fi19++ {
                    var v20 uint32 = 42
                    var v21 int8 = 42
                    switch -552 {
                    case 1:
                        if ((v4 / (((v4 | v4)) | 1)) != v4) {
                            var v22 float32 = 1e10
                            v23 := GenS2{F0: "a", F1: 1}
                            var v24 int32 = v4
                            _ = v22
                            _ = v23
                            _ = v24
                        } else {
                            _ = len(v5)
                            if false {
                                delete(v1, (v20 / (((v20 << ((805) & 63))) | 1)))
                            }
                            v20 = (v20 * (v20 >> ((1) & 63)))
                            v25 := GenS1{F0: 0}
                            _ = v25
                        }
                    }
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                v26 := GenS1{F0: genFn1(4294967295, 255, (v2 + 0))}
                _ = v18
                _ = v26
            }
            var v27 uint64 = 42
            var v28 int16 = 42
            _ = fi17
            _ = v27
            _ = v28
        }
    }
    if len(v6) > 0 {
        v6[0] = 127
    }
    var v29 float32 = 0.1
    v30 := GenS3{F0: ((!false) && (127 >= 99)), F1: ("" + ("" + "\n\t"))}
    v31 := v1[0]
    _ = len(v6)
    var v32 uint8 = 1
    for fi33 := 0; fi33 < 13; fi33++ {
        for fi34 := 0; fi34 < 2; fi34++ {
            var v35 int16 = v2
            if ((((v32 > 54) && true) || false) && (!false)) {
                var v36 bool = ((0.1 / (float64(v29) - v3)) > 0.1)
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = v36
            } else {
                var v37 uint16 = 980
                v38 := GenS4{F0: 32, F1: 42, F2: 544}
                switch "a" {
                case "hello":
                    fmt.Println(v35, v35, v35)
                    var v39 bool = true
                    if (!(116 != int64(v32))) {
                        v40 := GenS1{F0: v32}
                        var v41 int16 = 794
                        _ = v40
                        _ = v41
                    } else {
                        var v42 int32 = v4
                        v2 = int16(v4)
                        _ = v42
                    }
                    _ = v39
                case "a":
                    if true {
                        switch ("" + ((("a" + "hello") + "\n\t") + "")) {
                        case "foo bar":
                            if ("foo bar" < "") {
                                if true {
                                    v43 := GenS5{F0: (-(v4))}
                                    v4 += v4
                                    if true {
                                        if true {
                                            if (false || false) {
                                                var v44 uint16 = uint16(v32)
                                                _ = v44
                                            }
                                            v45 := GenS3{F0: (v32 != v32), F1: "hello"}
                                            var v46 uint32 = 725
                                            _ = v45
                                            _ = v46
                                        }
                                        v1[42] = v31
                                        var v47 uint64 = 42
                                        _ = v47
                                    } else {
                                        var v48 int32 = (-(v4))
                                        var v49 uint = 803
                                        v3 += v3
                                        _ = v48
                                        _ = v49
                                    }
                                    var v50 uint64 = 0
                                    switch fi33 {
                                    case 119:
                                        var v51 uint = 219
                                        _ = v51
                                    case 739:
                                        v35 -= v35
                                    case -9223372036854775808:
                                        if (((3.14159 - v3) * (v3 / v3)) == 1e10) {
                                            var v52 int64 = 358
                                            _ = v52
                                        } else {
                                            var v53 uint64 = (v50 + v50)
                                            var v54 int64 = -1
                                            if len(v6) > 0 {
                                                v6[0] = fi34
                                            }
                                            _ = v53
                                            _ = v54
                                        }
                                    default:
                                        var v55 bool = (int(v4) >= 1)
                                        v29 += 0.1
                                        _ = v55
                                    }
                                    if false {
                                        v56 := v1[2]
                                        v4 -= (v4 / ((v4) | 1))
                                        var v57 uint8 = 2
                                        if true {
                                            if true {
                                                if len(v6) > 0 {
                                                    v6[0] = 158
                                                }
                                                if len(v6) > 0 {
                                                    v6[0] = -2147483648
                                                }
                                            }
                                            v29 = 3.14159
                                        } else {
                                            var v58 int8 = 42
                                            if (0 == (v58 * v58)) {
                                                v58 += 102
                                            }
                                            var v59 uint8 = (v32 << ((0) & 63))
                                            _ = v58
                                            _ = v59
                                        }
                                        _ = v56
                                        _ = v57
                                    } else {
                                        v60 := GenS1{F0: (v32 * v32)}
                                        fmt.Println(v31)
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        _ = v60
                                    }
                                    v61 := GenS2{F0: "\n\t", F1: 255}
                                    _ = v43
                                    _ = v50
                                    _ = v61
                                }
                                v62 := GenS2{F0: "a", F1: uint16(v32)}
                                _ = v62
                            }
                            v32 = (v32 % (((^(v32))) | 1))
                        case "\n\t":
                            if (18446744073709551615 < 48) {
                                switch -268 {
                                case 1:
                                    v63 := GenS6{F0: "\n\t"}
                                    _ = v63
                                case 32767:
                                    v3 = v3
                                case -121:
                                    _ = len(v6)
                                }
                            } else {
                                switch (v31 & (v31 & (fi33 * v31))) {
                                case 799:
                                    v31 -= (fi34 | fi33)
                                case -105:
                                    var v64 int64 = 353
                                    _ = v64
                                case 32767:
                                    switch "\n\t" {
                                    case "a":
                                        v29 -= (v29 * v29)
                                    case "\n\t":
                                        v31 += -186
                                        var v65 uint32 = 4294967295
                                        switch v31 {
                                        case 2147483647:
                                            v65 -= v65
                                        }
                                        _ = v65
                                    case "hello":
                                        fmt.Println(v2, v2)
                                    }
                                default:
                                    switch (v31 / ((fi33) | 1)) {
                                    case -9223372036854775808:
                                        if (!(genFn0() <= 1)) {
                                            v66 := GenS6{F0: "a"}
                                            v67 := GenS7{F0: 0, F1: (-(fi33))}
                                            _ = v66
                                            _ = v67
                                        } else {
                                            fmt.Println(v29, fi33)
                                            v37 = (v37 % (((v37 & (v37 >> ((974) & 63)))) | 1))
                                            var v68 int16 = v35
                                            _ = v68
                                        }
                                    case 2147483647:
                                        var v69 string = ""
                                        _ = v69
                                    default:
                                        v4 = v4
                                        if true {
                                            var v70 uint16 = v37
                                            var v71 uint64 = 18446744073709551615
                                            v1[718] = v31
                                            _ = v70
                                            _ = v71
                                        }
                                    }
                                    v72 := v5[226]
                                    v5[0] = 2147483648
                                    _ = v72
                                }
                                v73 := GenS7{F0: 516, F1: v31}
                                switch (fi33 & (fi34 + fi34)) {
                                case -2147483648:
                                    fmt.Println(v4, v35)
                                case 128:
                                    if false {
                                        var v74 uint32 = 42
                                        var v75 int32 = v4
                                        if true {
                                            v76 := GenS2{F0: "\n\t", F1: uint16(v74)}
                                            if (v37 < 32767) {
                                                var v77 float32 = float32(float64(float32(v3)))
                                                if (0.1 > (float64(v77) * (float64(0.0) + v3))) {
                                                    var v78 uint16 = (v37 >> ((genFn0()) & 63))
                                                    var v79 uint32 = (^(v74))
                                                    var v80 uint32 = v74
                                                    _ = v78
                                                    _ = v79
                                                    _ = v80
                                                }
                                                switch 128 {
                                                case 4294967295:
                                                    fmt.Println(v35)
                                                case -875:
                                                    var v81 float32 = (v29 / (((v29 - v29) * (v29 - 1.5)) / 1.0))
                                                    _ = v81
                                                case 128:
                                                    v35 = 870
                                                }
                                                _ = v77
                                            } else {
                                                if (1 <= v4) {
                                                    fmt.Println(fi33)
                                                    var v82 float64 = float64((1e10 + v29))
                                                    var v83 uint8 = v32
                                                    if false {
                                                        delete(v5, (v74 - (v74 / (((v74 - (v74 >> ((889) & 63)))) | 1))))
                                                        v84 := GenS6{F0: ("hello" + ("" + ("a" + "a")))}
                                                        var v85 int = fi34
                                                        _ = v84
                                                        _ = v85
                                                    }
                                                    _ = v82
                                                    _ = v83
                                                } else {
                                                    var v86 int32 = v4
                                                    var v87 uint64 = 1
                                                    _ = v86
                                                    _ = v87
                                                }
                                            }
                                            _ = v76
                                        }
                                        _ = v74
                                        _ = v75
                                    }
                                default:
                                    var v88 uint = 738
                                    _ = v88
                                }
                                var v89 uint8 = (v32 - 255)
                                _ = v73
                                _ = v89
                            }
                        }
                    } else {
                        if ("\n\t" < "\n\t") {
                            v90 := GenS8{F0: 42, F1: 65535}
                            switch ("foo bar" + ("foo bar" + ("\n\t" + "a"))) {
                            case "hello":
                                switch (("\n\t" + "hello") + "\n\t") {
                                case "foo bar":
                                    v91 := GenS9{F0: (v37 % ((v37) | 1))}
                                    if ((!(("hello" + ("\n\t" + "hello")) > "hello")) && (v29 != v29)) {
                                        v92 := GenS10{F0: v32}
                                        _ = v92
                                    } else {
                                        switch "\n\t" {
                                        case "foo bar":
                                            var v93 int = -9223372036854775808
                                            _ = v93
                                        default:
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                        }
                                        v6 = append(v6, fi33)
                                        if true {
                                            switch -752 {
                                            case 128:
                                                if true {
                                                    fmt.Println(v4, v29)
                                                    if false {
                                                        if false {
                                                            v3 = (v3 + v3)
                                                            switch (fi33 - v31) {
                                                            case -638:
                                                                v94 := GenS11{F0: (v4 << ((4294967295) & 63)), F1: (^(v37))}
                                                                _ = v94
                                                            }
                                                            v95 := GenS12{F0: false, F1: 32767}
                                                            _ = v95
                                                        }
                                                        var v96 bool = true
                                                        _ = v96
                                                    }
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                } else {
                                                    var v97 float32 = v29
                                                    fmt.Println(v35)
                                                    var v98 float32 = ((1.5 - v97) - ((1.5 - (float32(((1e10 - 1e10) + (1e10 + ((v3 + v3) + (v3 / v3))))) + v97)) + v97))
                                                    var v99 uint64 = 18446744073709551615
                                                    v35 += (v35 << ((4294967295) & 63))
                                                    _ = v97
                                                    _ = v98
                                                    _ = v99
                                                }
                                            }
                                        } else {
                                            switch ("" + "foo bar") {
                                            case "hello":
                                                if len(v6) > 0 {
                                                    v6[0] = fi33
                                                }
                                            case "foo bar":
                                                v100 := GenS13{F0: (("foo bar" + "a") + ""), F1: -128, F2: (v37 + (v37 & v37))}
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                                v3 = -2.25
                                                _ = v100
                                            case "":
                                                switch fi34 {
                                                case -1:
                                                    var v101 float32 = v29
                                                    _ = v101
                                                case 128:
                                                    v4 = v4
                                                    v3 = 3.14159
                                                case -250:
                                                    var v102 int = (-(fi34))
                                                    _ = v102
                                                default:
                                                    fmt.Println(v4, v35)
                                                }
                                            }
                                        }
                                    }
                                    _ = v91
                                case "\n\t":
                                    var v103 int8 = 0
                                    _ = v103
                                default:
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                }
                            default:
                                var v104 int16 = 1000
                                _ = v104
                            }
                            if false {
                                var v105 float64 = v3
                                _ = v105
                            } else {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                            }
                            _ = v90
                        } else {
                            var v106 bool = ((v32 >> ((genFn0()) & 63)) == (v32 % ((v32) | 1)))
                            v107 := GenS13{F0: "foo bar", F1: 86, F2: v37}
                            v106 = v106
                            _ = v106
                            _ = v107
                        }
                        var v108 uint = 4294967295
                        _ = v108
                    }
                case "":
                    delete(v5, 2147483648)
                default:
                    var v109 float64 = 0.0
                    switch "a" {
                    case "hello":
                        var v110 int64 = -1
                        _ = v110
                    default:
                        var v111 int64 = 544
                        _ = v111
                    }
                    _ = v109
                }
                _ = v37
                _ = v38
            }
            _ = fi34
            _ = v35
        }
        for fi112 := 0; fi112 < 18; fi112++ {
            _ = len(v5)
            if (-128 <= -93) {
                if ((!false) && false) {
                    _ = len(v6)
                } else {
                    fmt.Println(v3, v3)
                }
                if false {
                    v4 -= v4
                    var v113 uint = genFn0()
                    _ = v113
                }
                v29 += 1e10
            } else {
                v2 = (v2 & v2)
            }
            var v114 float64 = (v3 + (((v3 - 1.5) - v3) - (0.1 * (v3 - (((v3 + v3) / -2.25) * 1e10)))))
            _ = len(v1)
            fmt.Println(v31)
            _ = fi112
            _ = v114
        }
        switch "a" {
        case "hello":
            fmt.Println(v29)
        default:
            v4 += 933
        }
        _ = fi33
    }
    var v115 int64 = 1
    if (!true) {
        _ = len(v6)
    } else {
        v116 := GenS14{F0: v31}
        var v117 uint = genFn0()
        v3 = v3
        _ = v116
        _ = v117
    }
    if true {
        v2 = (-(v2))
    } else {
        v6 = append(v6, (v31 << ((42) & 63)))
        var v118 uint8 = (v32 / ((v32) | 1))
        _ = v118
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v115:", v115)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v115
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
