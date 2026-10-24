package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64; F1 string; F2 string }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 uint16; F1 string; F2 int64 }
type GenS6 struct { F0 uint32; F1 uint }
type GenS7 struct { F0 uint; F1 string; F2 uint16 }
type GenS8 struct { F0 uint32; F1 uint8; F2 uint }
type GenS9 struct { F0 uint64; F1 uint16; F2 uint }
type GenS10 struct { F0 uint16; F1 uint; F2 int8 }
type GenS11 struct { F0 int8; F1 uint32; F2 uint }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 uint32; F1 uint32; F2 uint8 }
type GenS14 struct { F0 string; F1 int; F2 int16 }
func genFn0(p0 uint64, p1 int8) int32 {
    if (208 >= 255) {
        var v7 int8 = p1
        _ = v7
    } else {
        var v8 float64 = float64(1e10)
        fmt.Println(v8)
        if (false && (42 < 42)) {
            p0 = 2
        } else {
            v9 := GenS1{F0: -456, F1: 42, F2: (-2.25 < 0.1)}
            _ = v9
        }
        _ = v8
    }
    var v10 int32 = 603
    _ = p0
    _ = p1
    _ = v10
    return v10
}
func genFn1(p0 uint, p1 bool, p2 uint16) bool {
    var v11 int8 = 127
    switch 24 {
    case 32768:
        for fi12 := 0; fi12 < 1; fi12++ {
            if p1 {
                p2 -= 516
                v13 := GenS1{F0: -717, F1: -32768, F2: p1}
                _ = v13
            }
            v11 = (-(v11))
            var v14 uint = p0
            _ = fi12
            _ = v14
        }
    case 42:
        var v15 int16 = 1000
        _ = v15
    case 2147483647:
        var v16 int8 = 1
        _ = v16
    }
    var v17 string = "\n\t"
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v17
    return p1
}

func main() {
    var v1 int = 2147483647
    var v2 uint32 = 2147483647
    var v3 string = "a"
    var v4 uint64 = 1
    v5 := make(map[string]uint32)
    var v6 uint8 = 127
    v18 := GenS1{F0: 32767, F1: -32768, F2: ((346 < -508) || (v2 != v2))}
    for fi19 := 0; fi19 < 12; fi19++ {
        fmt.Println(v2)
        v3 = "a"
        for fi20 := 0; fi20 < 16; fi20++ {
            var v21 uint8 = uint8(v2)
            var v22 bool = false
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = fi19
    }
    fmt.Println(v2, v4, v4)
    v23 := GenS1{F0: -679, F1: 412, F2: false}
    fmt.Println(v3)
    for fi24 := 0; fi24 < 8; fi24++ {
        switch ((v3 + v3) + "hello") {
        case "foo bar":
            if false {
                v1 += fi24
                for fi25 := 0; fi25 < 20; fi25++ {
                    v26 := GenS2{F0: (^(fi25))}
                    delete(v5, v3)
                    _ = fi25
                    _ = v26
                }
            } else {
                for fi27 := 0; fi27 < 8; fi27++ {
                    var v28 string = v3
                    v29 := []int{}
                    if true {
                        if (-1 != -32768) {
                            var v30 uint64 = (v4 * v4)
                            if (2 > 32767) {
                                fmt.Println(fi24, v6)
                            } else {
                                if (414 < (^(v2))) {
                                    fmt.Println(v30, v3, v2)
                                }
                                v31 := v5[v28]
                                v32 := GenS2{F0: (fi24 | fi27)}
                                fmt.Println(fi27, v31)
                                v31 += (v31 << ((391) & 63))
                                _ = v31
                                _ = v32
                            }
                            if true {
                                _ = len(v5)
                                v33 := GenS3{F0: -9223372036854775808, F1: v3, F2: v3}
                                v30 = (v30 >> ((65535) & 63))
                                _ = v33
                            }
                            _ = v30
                        } else {
                            if true {
                                var v34 int32 = -788
                                v35 := GenS1{F0: -32768, F1: 42, F2: false}
                                switch fi24 {
                                case -2147483648:
                                    v29 = append(v29, 32767)
                                default:
                                    fmt.Println(v28, v3, v34)
                                }
                                _ = v34
                                _ = v35
                            }
                            switch 4294967295 {
                            case 2147483647:
                                v36 := GenS1{F0: -32768, F1: -714, F2: (1 <= 0)}
                                v4 -= (v4 ^ (v4 % (((v4 + 31)) | 1)))
                                _ = v36
                            default:
                                var v37 uint8 = v6
                                switch v3 {
                                case "":
                                    var v38 int64 = 0
                                    _ = v38
                                case "\n\t":
                                    var v39 int = 42
                                    _ = v39
                                case "foo bar":
                                    v40 := GenS3{F0: 42, F1: v28, F2: (v3 + v3)}
                                    _ = v40
                                default:
                                    var v41 int64 = 0
                                    _ = v41
                                }
                                if (445 < 439) {
                                    v42 := GenS4{F0: (v4 != 1), F1: -2147483648}
                                    v37 += (v37 << ((uint(v37)) & 63))
                                    v4 = (v4 % ((v4) | 1))
                                    fmt.Println(v37, v37, v6)
                                    _ = v42
                                }
                                if false {
                                    fmt.Println(v1)
                                } else {
                                    if (!false) {
                                        var v43 int = (fi24 / ((v1) | 1))
                                        _ = v43
                                    }
                                }
                                _ = v37
                            }
                        }
                        v44 := GenS4{F0: ((v2 + v2) >= v2), F1: -330}
                        if true {
                            fmt.Println(fi27, v2, v3)
                            var v45 int64 = -1
                            if ((v45 & -792) <= 42) {
                                v45 = v45
                            }
                            _ = v45
                        } else {
                            if (!(-1 != 9223372036854775807)) {
                                v2 += 2147483648
                            }
                            v5[v3] = (v2 << ((721) & 63))
                            var v46 int = fi24
                            _ = v46
                        }
                        _ = v44
                    } else {
                        var v47 int64 = 42
                        _ = len(v29)
                        _ = v47
                    }
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                for fi48 := 0; fi48 < 12; fi48++ {
                    v3 = v3
                    v2 += v2
                    _ = fi48
                }
            }
        case "a":
            v49 := GenS3{F0: -879, F1: (v3 + ("foo bar" + v3)), F2: "foo bar"}
            _ = v49
        default:
            var v50 bool = true
            _ = v50
        }
        if (!((2147483647 <= int32(v4)) || false)) {
            v3 = v3
        }
        v6 += 127
        _ = fi24
    }
    for fi51 := 0; fi51 < 8; fi51++ {
        v2 = 42
        _ = fi51
    }
    var v52 int64 = 997
    if true {
        if true {
            switch -1 {
            case 2147483647:
                if true {
                    v2 = (^(v2))
                    for fi53 := 0; fi53 < 14; fi53++ {
                        if (true || true) {
                            var v54 uint16 = 611
                            _ = v54
                        } else {
                            v6 += (v6 - v6)
                            if true {
                                if (v4 >= v4) {
                                    for fi55 := 0; fi55 < 17; fi55++ {
                                        v1 += 4294967295
                                        _ = fi55
                                    }
                                    if ((v6 % ((v6) | 1)) <= 1) {
                                        v1 = 9223372036854775807
                                    } else {
                                        v56 := GenS5{F0: 42, F1: v3, F2: v52}
                                        var v57 int16 = -1
                                        _ = v56
                                        _ = v57
                                    }
                                }
                                if true {
                                    for fi58 := 0; fi58 < 11; fi58++ {
                                        delete(v5, v3)
                                        delete(v5, "hello")
                                        _ = fi58
                                    }
                                    v4 = 18446744073709551615
                                } else {
                                    fmt.Println(v1)
                                    v59 := []int{(fi53 ^ -671), (fi53 | v1)}
                                    if false {
                                        v3 = v3
                                        v4 = (v4 ^ v4)
                                    } else {
                                        v60 := GenS6{F0: v2, F1: 453}
                                        switch (v3 + ((("a" + ("foo bar" + ("\n\t" + v3))) + v3) + (v3 + v3))) {
                                        case "a":
                                            var v61 float64 = 1.5
                                            _ = v61
                                        case "foo bar":
                                            for fi62 := 0; fi62 < 8; fi62++ {
                                                if false {
                                                    switch (fi62 + 0) {
                                                    case -622:
                                                        var v63 uint16 = 1
                                                        _ = v63
                                                    case -9223372036854775808:
                                                        _ = len(v5)
                                                    default:
                                                        if false {
                                                            switch "foo bar" {
                                                            case "a":
                                                                var v64 uint32 = v2
                                                                _ = v64
                                                            case "\n\t":
                                                                if false {
                                                                    switch "\n\t" {
                                                                    case "a":
                                                                        if ((0.0 - 3.14159) <= 1.5) {
                                                                            var v65 bool = true
                                                                            switch v3 {
                                                                            case "\n\t":
                                                                                v66 := GenS2{F0: 263}
                                                                                _ = v66
                                                                            case "":
                                                                                v3 = ""
                                                                            default:
                                                                                switch v3 {
                                                                                case "foo bar":
                                                                                    if v65 {
                                                                                        var v67 int64 = (v52 | v52)
                                                                                        var v68 float32 = 3.14159
                                                                                        var v69 uint16 = 1
                                                                                        _ = v67
                                                                                        _ = v68
                                                                                        _ = v69
                                                                                    } else {
                                                                                        if true {
                                                                                            var v70 bool = true
                                                                                            if (!v65) {
                                                                                                v71 := GenS7{F0: 2, F1: v3, F2: 2}
                                                                                                v4 -= (^(v4))
                                                                                                _ = v71
                                                                                            }
                                                                                            _ = v70
                                                                                        }
                                                                                        v4 = v4
                                                                                    }
                                                                                case "a":
                                                                                    var v72 uint64 = (v4 | v4)
                                                                                    _ = v72
                                                                                }
                                                                            }
                                                                            var v73 int64 = v52
                                                                            v73 -= (v52 - v73)
                                                                            _ = v65
                                                                            _ = v73
                                                                        }
                                                                    case "":
                                                                        var v74 string = v3
                                                                        var v75 int8 = 0
                                                                        _ = v74
                                                                        _ = v75
                                                                    case "hello":
                                                                        v4 = v4
                                                                    default:
                                                                        var v76 string = "hello"
                                                                        _ = v76
                                                                    }
                                                                    v77 := GenS2{F0: fi62}
                                                                    _ = v77
                                                                } else {
                                                                    switch v3 {
                                                                    case "a":
                                                                        if (0 < -45) {
                                                                            delete(v5, ("hello" + ("\n\t" + ((((v3 + v3) + v3) + v3) + "hello"))))
                                                                        } else {
                                                                            v5[""] = 87
                                                                            var v78 float64 = 0.0
                                                                            switch v3 {
                                                                            case "":
                                                                                fmt.Println(v3, v52)
                                                                                switch fi62 {
                                                                                case 42:
                                                                                    v2 -= v2
                                                                                }
                                                                            }
                                                                            _ = v78
                                                                        }
                                                                    case "hello":
                                                                        if (1 >= v52) {
                                                                            if (false || (127 == 127)) {
                                                                                _ = len(v5)
                                                                                if true {
                                                                                    v79 := GenS4{F0: true, F1: 160}
                                                                                    if ((0.1 + float64(0.0)) <= (float64(float32((-2.25 / 1.0))) * 0.0)) {
                                                                                        v2 = (v2 % ((v2) | 1))
                                                                                        v80 := GenS8{F0: 42, F1: v6, F2: 18446744073709551615}
                                                                                        _ = v80
                                                                                    } else {
                                                                                        if (false || (v6 != (v6 << ((638) & 63)))) {
                                                                                            var v81 float32 = 0.1
                                                                                            v6 -= (v6 / ((v6) | 1))
                                                                                            _ = v81
                                                                                        }
                                                                                        v82 := GenS8{F0: (v2 % ((2147483647) | 1)), F1: v6, F2: 262}
                                                                                        _ = v82
                                                                                    }
                                                                                    _ = v79
                                                                                } else {
                                                                                    var v83 uint32 = 0
                                                                                    _ = v83
                                                                                }
                                                                                v84 := GenS3{F0: (v52 << ((157) & 63)), F1: v3, F2: v3}
                                                                                _ = v84
                                                                            } else {
                                                                                if (!true) {
                                                                                    v85 := GenS5{F0: 2, F1: ((v3 + (v3 + (v3 + ("" + "")))) + v3), F2: 42}
                                                                                    switch "" {
                                                                                    case "foo bar":
                                                                                        var v86 uint16 = 255
                                                                                        _ = v86
                                                                                    default:
                                                                                        v87 := GenS9{F0: uint64(v2), F1: 32767, F2: 42}
                                                                                        _ = v87
                                                                                    }
                                                                                    if false {
                                                                                        v1 -= 932
                                                                                    }
                                                                                    _ = v85
                                                                                } else {
                                                                                    if true {
                                                                                        v59 = append(v59, v1)
                                                                                        if false {
                                                                                            var v88 int8 = 127
                                                                                            if (v88 != v88) {
                                                                                                var v89 int = int(v6)
                                                                                                if true {
                                                                                                    if false {
                                                                                                        var v90 uint32 = v2
                                                                                                        var v91 int = (-(v1))
                                                                                                        v2 = uint32(v6)
                                                                                                        _ = v90
                                                                                                        _ = v91
                                                                                                    }
                                                                                                    var v92 uint = 770
                                                                                                    switch (fi53 & (v1 / ((v89) | 1))) {
                                                                                                    case 42:
                                                                                                        var v93 int = (fi53 << ((v92) & 63))
                                                                                                        _ = v93
                                                                                                    case 128:
                                                                                                        v52 += (v52 ^ (v52 << (((^(v92))) & 63)))
                                                                                                    }
                                                                                                    _ = v92
                                                                                                } else {
                                                                                                    var v94 uint64 = (v4 & (v4 - (v4 - v4)))
                                                                                                    var v95 int = (fi53 + 128)
                                                                                                    v94 += (v4 * v4)
                                                                                                    _ = v94
                                                                                                    _ = v95
                                                                                                }
                                                                                                _ = v89
                                                                                            }
                                                                                            v5["\n\t"] = v2
                                                                                            v2 -= 65535
                                                                                            _ = v88
                                                                                        } else {
                                                                                            v6 = (^(v6))
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                            v96 := GenS10{F0: 65535, F1: 255, F2: 127}
                                                                            v1 = fi62
                                                                            _ = v96
                                                                        }
                                                                    case "foo bar":
                                                                        var v97 uint64 = uint64(v2)
                                                                        _ = v97
                                                                    }
                                                                    v2 += 1
                                                                    v1 += fi53
                                                                    if (!(v3 < v3)) {
                                                                        var v98 int16 = 1000
                                                                        v99 := GenS8{F0: (v2 | (v2 >> ((42) & 63))), F1: (v6 + (v6 * (^(v6)))), F2: 2}
                                                                        v98 -= (v98 ^ 1)
                                                                        _ = v98
                                                                        _ = v99
                                                                    }
                                                                    v100 := GenS7{F0: 0, F1: ((("\n\t" + "foo bar") + "a") + (("foo bar" + v3) + ("" + v3))), F2: 2}
                                                                    _ = v100
                                                                }
                                                            default:
                                                                var v101 int32 = -2147483648
                                                                _ = v101
                                                            }
                                                            var v102 int16 = 1
                                                            v103 := GenS11{F0: 42, F1: (v2 / ((v2) | 1)), F2: 42}
                                                            var v104 uint32 = v2
                                                            v102 = 32767
                                                            _ = v102
                                                            _ = v103
                                                            _ = v104
                                                        }
                                                    }
                                                    var v105 float64 = float64(1.5)
                                                    if false {
                                                        fmt.Println(v52, v52)
                                                        var v106 int64 = (v52 + -641)
                                                        _ = v106
                                                    } else {
                                                        v6 = v6
                                                        v107 := GenS9{F0: (^(v4)), F1: 423, F2: 2}
                                                        _ = v107
                                                    }
                                                    _ = v105
                                                } else {
                                                    switch 9223372036854775807 {
                                                    case 32768:
                                                        v3 = v3
                                                    case 0:
                                                        v108 := GenS12{F0: -1}
                                                        _ = v108
                                                    }
                                                }
                                                _ = fi62
                                            }
                                        }
                                        switch (v3 + v3) {
                                        case "a":
                                            var v109 int16 = -869
                                            _ = v109
                                        case "foo bar":
                                            v52 += 9223372036854775807
                                        }
                                        _ = v60
                                    }
                                    _ = v59
                                }
                                var v110 uint = 1
                                _ = v110
                            } else {
                                fmt.Println(v6)
                                if false {
                                    if true {
                                        var v111 int16 = 767
                                        if false {
                                            v112 := GenS13{F0: (v2 + 985), F1: (v2 ^ uint32(v4)), F2: v6}
                                            v6 = 42
                                            _ = v112
                                        }
                                        _ = v111
                                    } else {
                                        var v113 uint32 = (v2 * 348)
                                        fmt.Println(v3, v4)
                                        v6 += v6
                                        _ = v113
                                    }
                                } else {
                                    v114 := v5[(v3 + "hello")]
                                    _ = v114
                                }
                            }
                            v115 := GenS14{F0: v3, F1: int(v52), F2: 724}
                            _ = v115
                        }
                        fmt.Println(fi53)
                        var v116 int64 = -283
                        _ = fi53
                        _ = v116
                    }
                    if (v6 == v6) {
                        v3 = (v3 + ("a" + "foo bar"))
                        switch 128 {
                        case 4294967295:
                            v2 -= 4294967295
                        }
                        v2 = 553
                    }
                }
            case -1:
                v5[(v3 + "a")] = 657
            }
        } else {
            fmt.Println(v3, v3)
            _ = len(v5)
        }
    } else {
        var v117 uint16 = 2
        var v118 bool = (!false)
        v117 -= 42
        _ = v117
        _ = v118
    }
    v119 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v119:", len(v119), v119)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v119
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
}
