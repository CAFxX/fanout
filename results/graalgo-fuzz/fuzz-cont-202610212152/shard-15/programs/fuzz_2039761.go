package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 int32; F1 uint32; F2 int32 }
type GenS5 struct { F0 int32; F1 int32 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int16; F1 int8 }
type GenS8 struct { F0 int16; F1 bool }
type GenS9 struct { F0 int32; F1 int16; F2 int32 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 uint8; F1 uint32 }
type GenS14 struct { F0 uint; F1 int16; F2 int8 }
type GenS15 struct { F0 int16; F1 int16; F2 int32 }
type GenS16 struct { F0 uint8; F1 uint8 }
type GenS17 struct { F0 int8; F1 bool }
type GenS18 struct { F0 uint8; F1 bool; F2 bool }
type GenS19 struct { F0 int8; F1 int }
func genFn0() uint64 {
    var v6 string = "foo bar"
    v6 = (v6 + (v6 + v6))
    v7 := GenS1{F0: 2, F1: 1000}
    v6 = (v6 + ((("\n\t" + (("" + (((v6 + ("" + v6)) + v6) + v6)) + (v6 + "foo bar"))) + v6) + v6))
    if true {
        var v8 int32 = 801
        v9 := GenS1{F0: 2147483648, F1: 660}
        _ = v8
        _ = v9
    }
    _ = v6
    _ = v7
    return 291
}
func genFn1(p0 uint, p1 int64, p2 int16) uint {
    v10 := GenS2{F0: 2}
    for fi11 := 0; fi11 < 3; fi11++ {
        v12 := GenS1{F0: 42, F1: p2}
        v13 := []int{(fi11 * (fi11 / (((fi11 / ((int(p2)) | 1))) | 1))), (fi11 / (((fi11 % ((int(p2)) | 1))) | 1)), 32767}
        if (-1 < fi11) {
            _ = len(v13)
            for fi14 := 0; fi14 < 12; fi14++ {
                switch (fi11 + int(p1)) {
                case 42:
                    p2 = (p2 - 32767)
                case 4294967295:
                    if true {
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                        if true {
                            var v15 uint8 = 127
                            if false {
                                v13 = append(v13, fi11)
                                if true {
                                    switch (fi14 >> ((p0) & 63)) {
                                    case -9223372036854775808:
                                        if len(v13) >= 2 {
                                            v13 = v13[1:len(v13)]
                                        }
                                    case 1:
                                        var v16 int8 = 127
                                        _ = v16
                                    }
                                    fmt.Println(fi14, p1, p0)
                                    v13 = append(v13, -9223372036854775808)
                                    switch 379 {
                                    case 1:
                                        p2 = 32767
                                    case 4294967295:
                                        var v17 int32 = 42
                                        _ = v17
                                    }
                                } else {
                                    var v18 int64 = p1
                                    _ = v18
                                }
                                var v19 int32 = -319
                                _ = v19
                            }
                            v20 := GenS1{F0: uint32(p0), F1: p2}
                            _ = v15
                            _ = v20
                        } else {
                            var v21 uint8 = 73
                            var v22 uint16 = 65535
                            _ = v21
                            _ = v22
                        }
                        var v23 string = "a"
                        _ = v23
                    } else {
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                        fmt.Println(fi11, p1)
                        if (false || (-1 >= 42)) {
                            var v24 float64 = ((1e10 - 0.1) - 0.1)
                            var v25 float32 = (0.1 / float32((1e10 * ((v24 * v24) / -2.25))))
                            v26 := GenS3{F0: -1, F1: true}
                            _ = v24
                            _ = v25
                            _ = v26
                        } else {
                            if len(v13) > 0 {
                                v13[0] = (fi14 / ((fi14) | 1))
                            }
                        }
                        _ = len(v13)
                    }
                default:
                    var v27 uint64 = 873
                    _ = v27
                }
                var v28 int = fi14
                p1 = (p1 - p1)
                _ = fi14
                _ = v28
            }
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    if (("hello" + ("hello" + (("" + "\n\t") + ""))) <= ("hello" + "")) {
        p2 -= (p2 - -1)
        v29 := GenS3{F0: 0, F1: false}
        _ = v29
    } else {
        v30 := GenS2{F0: 241}
        for fi31 := 0; fi31 < 17; fi31++ {
            fmt.Println(p0)
            p0 = (p0 % ((2) | 1))
            _ = fi31
        }
        p0 += p0
        _ = v30
    }
    p2 = p2
    var v32 string = "foo bar"
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v32
    return (p0 / ((p0) | 1))
}

func main() {
    var v1 uint16 = 403
    var v2 int64 = 9223372036854775807
    v3 := make(map[string]int)
    var v4 int8 = 42
    var v5 uint8 = 42
    v1 = v1
    v2 += -229
    var v33 bool = true
    v34 := []int{}
    var v35 float32 = 1.5
    v35 = ((3.14159 + float32(0.1)) - (v35 / v35))
    if (!true) {
        var v36 uint16 = 2
        if v33 {
            if true {
                var v37 uint64 = 368
                v38 := GenS4{F0: -1, F1: uint32(v5), F2: 1}
                if len(v34) > 0 {
                    v34[0] = -9223372036854775808
                }
                _ = v37
                _ = v38
            }
            if ((0.1 <= (0.1 / ((3.14159 * (1e10 - float32((0.1 + 0.0)))) - v35))) || true) {
                for fi39 := 0; fi39 < 11; fi39++ {
                    switch "" {
                    case "hello":
                        if (!v33) {
                            v40 := GenS1{F0: 2147483648, F1: 32767}
                            if ((!true) && (false || v33)) {
                                v3[("\n\t" + "a")] = fi39
                                v41 := GenS5{F0: -409, F1: -216}
                                if v33 {
                                    fmt.Println(v33, v35)
                                    if ((v4 - v4) > (^(v4))) {
                                        var v42 uint32 = 827
                                        var v43 uint32 = (v42 + v42)
                                        var v44 int64 = v2
                                        _ = v42
                                        _ = v43
                                        _ = v44
                                    } else {
                                        _ = len(v34)
                                        v4 -= (v4 >> ((2) & 63))
                                    }
                                } else {
                                    if true {
                                        v45 := GenS6{F0: 4294967295}
                                        _ = v45
                                    }
                                    switch ("\n\t" + ("" + (("hello" + "foo bar") + "hello"))) {
                                    case "":
                                        var v46 float64 = -2.25
                                        _ = v46
                                    case "foo bar":
                                        if v33 {
                                            var v47 uint = 255
                                            var v48 float64 = (1.5 + 1.5)
                                            _ = v47
                                            _ = v48
                                        } else {
                                            for fi49 := 0; fi49 < 19; fi49++ {
                                                if len(v34) > 0 {
                                                    v34[0] = fi49
                                                }
                                                var v50 string = ("\n\t" + "foo bar")
                                                _ = fi49
                                                _ = v50
                                            }
                                            fmt.Println(fi39, v4)
                                        }
                                    default:
                                        if ((v35 - v35) > v35) {
                                            if ((^(v2)) <= 42) {
                                                switch int(v4) {
                                                case -806:
                                                    for fi51 := 0; fi51 < 12; fi51++ {
                                                        if len(v34) > 0 {
                                                            v34[0] = fi51
                                                        }
                                                        if true {
                                                            if v33 {
                                                                v36 = 42
                                                                if (true && ((839 == 2) && (0 <= 688))) {
                                                                    if (int(v1) != (fi51 << ((255) & 63))) {
                                                                        if true {
                                                                            fmt.Println(v36, v35)
                                                                        } else {
                                                                            var v52 uint16 = v1
                                                                            var v53 uint64 = 2
                                                                            _ = v52
                                                                            _ = v53
                                                                        }
                                                                        var v54 uint8 = 213
                                                                        _ = v54
                                                                    } else {
                                                                        if v33 {
                                                                            var v55 bool = (!true)
                                                                            _ = v55
                                                                        }
                                                                    }
                                                                    switch 128 {
                                                                    case -506:
                                                                        if v33 {
                                                                            v2 = v2
                                                                        } else {
                                                                            if len(v34) >= 2 {
                                                                                v34 = v34[1:len(v34)]
                                                                            }
                                                                            fmt.Println(v33, v2, v35)
                                                                            var v56 int64 = 9223372036854775807
                                                                            v35 = v35
                                                                            if v33 {
                                                                                var v57 int32 = 42
                                                                                if v33 {
                                                                                    var v58 uint64 = 0
                                                                                    v59 := GenS7{F0: 591, F1: (v4 | int8(v1))}
                                                                                    _ = v58
                                                                                    _ = v59
                                                                                }
                                                                                v36 -= (v1 << ((4294967295) & 63))
                                                                                _ = v57
                                                                            } else {
                                                                                fmt.Println(fi51, v4)
                                                                                fmt.Println(v2)
                                                                            }
                                                                            _ = v56
                                                                        }
                                                                    case -1:
                                                                        if v33 {
                                                                            v35 += v35
                                                                            var v60 string = (("" + ("foo bar" + "hello")) + "a")
                                                                            var v61 float32 = float32(0.0)
                                                                            v62 := GenS5{F0: -2147483648, F1: -2147483648}
                                                                            switch v60 {
                                                                            case "a":
                                                                                v63 := v3[v60]
                                                                                _ = v63
                                                                            }
                                                                            _ = v60
                                                                            _ = v61
                                                                            _ = v62
                                                                        } else {
                                                                            v33 = true
                                                                            v1 = (v1 ^ v1)
                                                                            var v64 string = ""
                                                                            v65 := GenS8{F0: -32768, F1: (v5 < (v5 + (v5 | 255)))}
                                                                            _ = v64
                                                                            _ = v65
                                                                        }
                                                                    }
                                                                }
                                                                _ = len(v3)
                                                            }
                                                            if (!(float64((float32(float64((v35 / v35))) / (1.5 + 3.14159))) != 1e10)) {
                                                                v66 := GenS9{F0: -1, F1: 32767, F2: -1}
                                                                _ = v66
                                                            } else {
                                                                if v33 {
                                                                    if (!(-45 == v4)) {
                                                                        if false {
                                                                            v67 := GenS8{F0: 1, F1: (!(-849 != 622))}
                                                                            _ = len(v34)
                                                                            v34 = append(v34, fi51)
                                                                            _ = v67
                                                                        } else {
                                                                            v2 = -9223372036854775808
                                                                        }
                                                                        var v68 int16 = 0
                                                                        var v69 uint = 255
                                                                        _ = v68
                                                                        _ = v69
                                                                    } else {
                                                                        _ = len(v3)
                                                                        var v70 uint32 = uint32(v1)
                                                                        _ = v70
                                                                    }
                                                                } else {
                                                                    var v71 uint32 = 1
                                                                    var v72 int16 = -520
                                                                    v73 := GenS10{F0: v72}
                                                                    _ = v71
                                                                    _ = v72
                                                                    _ = v73
                                                                }
                                                            }
                                                            var v74 uint64 = 2
                                                            _ = v74
                                                        }
                                                        if v33 {
                                                            if (v33 || true) {
                                                                v1 -= (v36 - (v1 - v36))
                                                                v3[""] = (fi39 * (^(fi51)))
                                                                switch 9223372036854775807 {
                                                                case 1:
                                                                    var v75 int32 = 1
                                                                    _ = v75
                                                                default:
                                                                    v76 := GenS11{F0: 65535}
                                                                    _ = v76
                                                                }
                                                            }
                                                            if v33 {
                                                                v77 := v3["\n\t"]
                                                                delete(v3, "hello")
                                                                v35 = v35
                                                                var v78 uint8 = v5
                                                                _ = len(v3)
                                                                _ = v77
                                                                _ = v78
                                                            } else {
                                                                switch (("hello" + "") + "hello") {
                                                                case "a":
                                                                    var v79 int16 = 42
                                                                    _ = v79
                                                                case "foo bar":
                                                                    v80 := GenS11{F0: uint(v36)}
                                                                    _ = v80
                                                                }
                                                                v81 := GenS4{F0: 2147483647, F1: 0, F2: -2147483648}
                                                                _ = v81
                                                            }
                                                        }
                                                        _ = fi51
                                                    }
                                                case 127:
                                                    v35 = v35
                                                default:
                                                    var v82 uint16 = (v36 - v1)
                                                    _ = v82
                                                }
                                                if (v33 || ((^(v2)) == v2)) {
                                                    v1 = 0
                                                }
                                            }
                                            for fi83 := 0; fi83 < 7; fi83++ {
                                                switch fi83 {
                                                case 707:
                                                    v84 := GenS12{F0: 2147483648}
                                                    _ = v84
                                                case -848:
                                                    v36 -= (^(v1))
                                                }
                                                if (!(v33 || v33)) {
                                                    if (!false) {
                                                        v36 = v36
                                                        fmt.Println(fi39, v33, v4)
                                                    }
                                                    if len(v34) > 0 {
                                                        v34[0] = fi39
                                                    }
                                                    if ("hello" < "hello") {
                                                        fmt.Println(v1, fi39, v1)
                                                        if v33 {
                                                            v33 = true
                                                        }
                                                        v85 := GenS6{F0: 0}
                                                        v35 -= 1.5
                                                        _ = v85
                                                    }
                                                    fmt.Println(v2)
                                                    switch (("\n\t" + (("a" + ("\n\t" + ("a" + "\n\t"))) + "\n\t")) + ("a" + "a")) {
                                                    case "\n\t":
                                                        fmt.Println(v2, v1)
                                                        if v33 {
                                                            v36 = v1
                                                            if len(v34) > 0 {
                                                                v34[0] = 42
                                                            }
                                                        } else {
                                                            if (833 != 1) {
                                                                var v86 int16 = 155
                                                                _ = v86
                                                            } else {
                                                                v33 = v33
                                                                var v87 int64 = (-(v2))
                                                                if v33 {
                                                                    var v88 uint64 = 574
                                                                    var v89 int64 = 379
                                                                    if len(v34) > 0 {
                                                                        v34[0] = fi83
                                                                    }
                                                                    _ = v88
                                                                    _ = v89
                                                                }
                                                                _ = v87
                                                            }
                                                            v90 := GenS13{F0: (v5 << ((705) & 63)), F1: 42}
                                                            if false {
                                                                if len(v34) >= 2 {
                                                                    v34 = v34[1:len(v34)]
                                                                }
                                                            } else {
                                                                if v33 {
                                                                    v33 = (v2 >= -1)
                                                                    fmt.Println(v33)
                                                                }
                                                                if (v2 != (v2 & (v2 | v2))) {
                                                                    var v91 uint = 42
                                                                    v92 := GenS14{F0: 0, F1: -1, F2: v4}
                                                                    _ = v91
                                                                    _ = v92
                                                                } else {
                                                                    v34 = append(v34, fi39)
                                                                    v5 = (^(v5))
                                                                }
                                                            }
                                                            _ = v90
                                                        }
                                                        delete(v3, "foo bar")
                                                        var v93 float64 = 1.5
                                                        switch ("\n\t" + "\n\t") {
                                                        case "hello":
                                                            switch "\n\t" {
                                                            case "a":
                                                                if len(v34) >= 2 {
                                                                    v34 = v34[1:len(v34)]
                                                                }
                                                            default:
                                                                v5 -= 255
                                                            }
                                                        case "foo bar":
                                                            v94 := GenS15{F0: -32768, F1: 32767, F2: 42}
                                                            _ = v94
                                                        case "\n\t":
                                                            if true {
                                                                v36 = v1
                                                                v33 = true
                                                                var v95 uint8 = 1
                                                                _ = v95
                                                            }
                                                        }
                                                        var v96 int64 = (v2 & (v2 & (v2 / ((v2) | 1))))
                                                        var v97 int16 = 415
                                                        var v98 uint = 4294967295
                                                        var v99 int8 = v4
                                                        var v100 bool = (!v33)
                                                        v101 := GenS16{F0: (v5 + 42), F1: 1}
                                                        _ = v100
                                                        _ = v101
                                                        _ = v93
                                                        _ = v96
                                                        _ = v97
                                                        _ = v98
                                                        _ = v99
                                                    case "":
                                                        v102 := GenS17{F0: v4, F1: ((-2.25 - float32(-2.25)) != (((v35 + 3.14159) / 1.0) + (0.1 + v35)))}
                                                        _ = v102
                                                    }
                                                }
                                                _ = fi83
                                            }
                                        } else {
                                            fmt.Println(v4, v33)
                                        }
                                    }
                                }
                                _ = v41
                            }
                            _ = v40
                        } else {
                            switch "foo bar" {
                            case "hello":
                                v103 := GenS6{F0: 627}
                                _ = v103
                            default:
                                var v104 uint64 = 922
                                _ = v104
                            }
                        }
                    case "\n\t":
                        v105 := GenS18{F0: (^(v5)), F1: v33, F2: (-128 == v4)}
                        _ = v105
                    default:
                        v4 += 104
                    }
                    v106 := v3[""]
                    _ = fi39
                    _ = v106
                }
                v107 := GenS9{F0: -1, F1: 32767, F2: -1}
                _ = v107
            } else {
                var v108 int16 = 32767
                v109 := GenS19{F0: v4, F1: -856}
                _ = v108
                _ = v109
            }
            v110 := GenS6{F0: 1}
            _ = v110
        }
        _ = v36
    } else {
        v3["a"] = 32768
        var v111 int8 = v4
        fmt.Println(v1, v33)
        _ = v111
    }
    var v112 int16 = 1000
    v113 := GenS18{F0: 128, F1: (2 > 66), F2: (("hello" + "\n\t") < (((("foo bar" + ("hello" + (((("\n\t" + "hello") + "\n\t") + "") + "\n\t"))) + "a") + "") + "foo bar"))}
    var v114 float64 = 0.0
    delete(v3, ("hello" + "\n\t"))
    fmt.Println("v1:", v1)
    fmt.Println("v112:", v112)
    fmt.Println("v114:", v114)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v112
    _ = v113
    _ = v114
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
