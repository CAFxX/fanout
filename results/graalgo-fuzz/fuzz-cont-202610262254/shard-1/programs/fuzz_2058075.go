package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint32; F1 uint64 }
type GenS6 struct { F0 uint16; F1 int64; F2 uint }
type GenS7 struct { F0 uint64; F1 uint }
type GenS8 struct { F0 int8; F1 uint16; F2 uint }
type GenS9 struct { F0 string; F1 uint16 }
type GenS10 struct { F0 uint32; F1 uint8; F2 uint64 }
type GenS11 struct { F0 uint8; F1 uint; F2 uint16 }
type GenS12 struct { F0 int8; F1 int; F2 int16 }
type GenS13 struct { F0 bool; F1 int }
type GenS14 struct { F0 int8; F1 uint8; F2 string }
type GenS15 struct { F0 int64; F1 uint32 }
type GenS16 struct { F0 int16; F1 bool }
func genFn0(p0 int64) uint {
    switch 394 {
    case 0:
        p0 = 0
    case 2147483647:
        switch 32767 {
        case -2147483648:
            v7 := GenS1{F0: 147}
            _ = v7
        case 128:
            fmt.Println(p0)
        default:
            switch "" {
            case "foo bar":
                var v8 int = 2147483647
                _ = v8
            case "":
                switch -6 {
                case -2147483648:
                    v9 := GenS2{F0: 160, F1: "a"}
                    _ = v9
                case 263:
                    for fi10 := 0; fi10 < 19; fi10++ {
                        if false {
                            var v11 uint8 = 128
                            v11 += 42
                            var v12 float32 = 0.1
                            _ = v11
                            _ = v12
                        }
                        var v13 int64 = 42
                        _ = fi10
                        _ = v13
                    }
                default:
                    fmt.Println(p0)
                }
            }
        }
    default:
        for fi14 := 0; fi14 < 20; fi14++ {
            v15 := GenS1{F0: 0}
            switch "" {
            case "a":
                fmt.Println(fi14, p0)
            }
            v16 := []int{(^(fi14)), (^(fi14)), fi14, fi14}
            _ = fi14
            _ = v15
            _ = v16
        }
    }
    var v17 uint16 = 1
    var v18 int8 = 114
    v19 := []int{}
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    _ = p0
    _ = v17
    _ = v18
    _ = v19
    return 1
}
func genFn1() int {
    var v20 uint = 18446744073709551615
    var v21 uint = genFn0(-1)
    switch "foo bar" {
    case "foo bar":
        for fi22 := 0; fi22 < 18; fi22++ {
            switch fi22 {
            case 32768:
                v20 += genFn0(581)
            case 4294967295:
                if true {
                    var v23 int64 = -9223372036854775808
                    if false {
                        v24 := []int{32767, -988, (fi22 - (fi22 * fi22))}
                        var v25 int64 = v23
                        for fi26 := 0; fi26 < 13; fi26++ {
                            v25 = v25
                            _ = fi26
                        }
                        _ = v24
                        _ = v25
                    }
                    v23 = 42
                    _ = v23
                } else {
                    v27 := []int{fi22, fi22, (-(fi22))}
                    v28 := GenS3{F0: 2147483647}
                    v21 = 255
                    _ = v27
                    _ = v28
                }
            case 96:
                v20 -= genFn0(147)
            default:
                v29 := GenS2{F0: 0, F1: "a"}
                _ = v29
            }
            _ = fi22
        }
    }
    _ = v20
    _ = v21
    return -2147483648
}

func main() {
    var v1 uint32 = 4294967295
    var v2 float32 = 0.0
    var v3 string = "a"
    var v4 string = v3
    v5 := []uint32{545, (v1 ^ v1), (v1 / (((v1 + (v1 * (v1 & (v1 | (v1 / ((2147483647) | 1))))))) | 1)), v1}
    v6 := make(map[string]uint32)
    for fi30 := 0; fi30 < 20; fi30++ {
        v31 := GenS4{F0: 18446744073709551615}
        fmt.Println(v1, v2, v4)
        v32 := GenS5{F0: (v1 >> ((65535) & 63)), F1: 18446744073709551615}
        _ = fi30
        _ = v31
        _ = v32
    }
    if false {
        v4 = v4
        v1 = v1
        var v33 uint = 18446744073709551615
        _ = v33
    } else {
        if true {
            for fi34 := 0; fi34 < 1; fi34++ {
                v35 := GenS6{F0: uint16(v1), F1: -11, F2: 18446744073709551615}
                switch "\n\t" {
                case "\n\t":
                    for fi36 := 0; fi36 < 10; fi36++ {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        v37 := GenS6{F0: 255, F1: 986, F2: 42}
                        _ = fi36
                        _ = v37
                    }
                case "foo bar":
                    for fi38 := 0; fi38 < 19; fi38++ {
                        if (!(false || (401 <= 0))) {
                            var v39 int64 = 42
                            _ = v39
                        } else {
                            delete(v6, "a")
                            v3 = ("foo bar" + "\n\t")
                            if (18446744073709551615 == 2) {
                                switch v4 {
                                case "\n\t":
                                    v3 = "\n\t"
                                default:
                                    var v40 uint64 = 808
                                    _ = v40
                                }
                                v41 := GenS6{F0: 0, F1: 9223372036854775807, F2: 65535}
                                _ = v41
                            }
                        }
                        switch -278 {
                        case -9223372036854775808:
                            switch (v4 + (v3 + v4)) {
                            case "foo bar":
                                var v42 bool = false
                                _ = v42
                            case "hello":
                                v3 = ((v4 + "\n\t") + "a")
                            default:
                                v43 := GenS2{F0: 1, F1: "a"}
                                _ = v43
                            }
                        default:
                            var v44 int8 = 70
                            _ = v44
                        }
                        v2 = (-2.25 - (v2 + 0.1))
                        _ = fi38
                    }
                default:
                    v45 := GenS5{F0: 2147483647, F1: 4294967295}
                    _ = v45
                }
                for fi46 := 0; fi46 < 3; fi46++ {
                    v1 = v1
                    _ = fi46
                }
                _ = fi34
                _ = v35
            }
            if (false || false) {
                v3 = v3
                _ = len(v5)
            } else {
                if false {
                    v1 = v1
                    var v47 uint = 65535
                    if false {
                        _ = len(v6)
                        fmt.Println(v4)
                    } else {
                        var v48 int = 42
                        switch "a" {
                        case "":
                            fmt.Println(v1, v3, v1)
                        case "a":
                            var v49 string = ""
                            _ = v49
                        }
                        _ = v48
                    }
                    _ = v47
                } else {
                    var v50 int = 280
                    _ = len(v6)
                    fmt.Println(v50)
                    _ = v50
                }
                for fi51 := 0; fi51 < 1; fi51++ {
                    if (!true) {
                        var v52 int64 = -669
                        v52 -= (v52 ^ 842)
                        _ = v52
                    }
                    if true {
                        if true {
                            for fi53 := 0; fi53 < 18; fi53++ {
                                fmt.Println(v2, v3, v2)
                                if ((true && false) && (!(true || (float32(0.0) != v2)))) {
                                    v5 = append(v5, 65535)
                                    var v54 bool = true
                                    _ = v54
                                }
                                _ = fi53
                            }
                        }
                    }
                    _ = fi51
                }
                var v55 bool = (914 < 160)
                _ = v55
            }
        }
        v56 := GenS7{F0: 18446744073709551615, F1: 2}
        _ = len(v6)
        _ = v56
    }
    v57 := GenS8{F0: 119, F1: 379, F2: 255}
    var v58 string = v3
    var v59 bool = (-128 <= 1)
    fmt.Println(v3)
    if (-28 != -2147483648) {
        var v60 int16 = 1
        var v61 int32 = 42
        _ = v60
        _ = v61
    }
    v62 := GenS9{F0: v58, F1: 2}
    var v63 uint8 = 1
    if (!(32 <= -128)) {
        fmt.Println(v58)
        v3 = ("foo bar" + "\n\t")
    } else {
        var v64 int8 = 96
        for fi65 := 0; fi65 < 19; fi65++ {
            if (!(!v59)) {
                for fi66 := 0; fi66 < 18; fi66++ {
                    switch v4 {
                    case "foo bar":
                        var v67 float32 = 1.5
                        _ = v67
                    default:
                        var v68 float64 = ((float64(v2) - 0.0) + (((float64(((v2 - v2) + v2)) / (0.1 / ((0.0 * 0.0) + 1e10))) * float64(v2)) - (0.1 * 1.5)))
                        _ = v68
                    }
                    if ((0.1 + (((3.14159 - 0.0) / (1.5 / ((-2.25 + ((0.0 - 1e10) * float64(v2))) / -2.25))) - (-2.25 / 1.0))) > 1.5) {
                        var v69 float32 = v2
                        if (true || (!((^(v63)) >= 127))) {
                            v70 := GenS10{F0: (v1 - (v1 - v1)), F1: v63, F2: 4294967295}
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            if v59 {
                                var v71 uint32 = v1
                                if v59 {
                                    v72 := GenS1{F0: (v63 << ((255) & 63))}
                                    _ = v72
                                }
                                _ = v71
                            } else {
                                v73 := GenS11{F0: v63, F1: 2, F2: 65535}
                                v74 := GenS12{F0: v64, F1: fi66, F2: 0}
                                _ = len(v6)
                                _ = v73
                                _ = v74
                            }
                            _ = v70
                        }
                        _ = v69
                    } else {
                        if (v59 || v59) {
                            var v75 uint16 = 42
                            var v76 int32 = 2147483647
                            _ = v75
                            _ = v76
                        } else {
                            v5 = append(v5, (v1 & v1))
                            v77 := GenS13{F0: v59, F1: (fi66 << ((1) & 63))}
                            _ = v77
                        }
                        v59 = v59
                        fmt.Println(v3)
                        var v78 int = 32767
                        v3 = (v3 + v4)
                        _ = v78
                    }
                    var v79 uint8 = 128
                    _ = fi66
                    _ = v79
                }
                for fi80 := 0; fi80 < 20; fi80++ {
                    var v81 uint16 = 144
                    fmt.Println(v58)
                    _ = fi80
                    _ = v81
                }
                for fi82 := 0; fi82 < 11; fi82++ {
                    v83 := GenS2{F0: (v1 * (v1 >> ((255) & 63))), F1: v58}
                    v3 = v3
                    v4 = (v4 + v4)
                    v2 -= (float32(1e10) / 3.14159)
                    if v59 {
                        v63 = (v63 >> ((42) & 63))
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                    v84 := GenS1{F0: 255}
                    _ = fi82
                    _ = v83
                    _ = v84
                }
            } else {
                switch "foo bar" {
                case "hello":
                    v4 = ("" + "hello")
                case "":
                    delete(v6, "a")
                case "foo bar":
                    fmt.Println(v59, v63, v63)
                default:
                    switch "foo bar" {
                    case "a":
                        v85 := GenS14{F0: 103, F1: (v63 & (v63 - uint8(v1))), F2: "hello"}
                        _ = v85
                    case "":
                        fmt.Println(v58, v59)
                    case "\n\t":
                        if v59 {
                            v59 = v59
                        } else {
                            var v86 int16 = -923
                            _ = v86
                        }
                    }
                }
                var v87 int = (fi65 << ((0) & 63))
                _ = v87
            }
            _ = fi65
        }
        _ = v64
    }
    _ = len(v5)
    var v88 uint32 = v1
    var v89 int8 = -1
    _ = len(v6)
    v90 := GenS15{F0: -9223372036854775808, F1: 42}
    for fi91 := 0; fi91 < 11; fi91++ {
        if (!v59) {
            v92 := GenS10{F0: 65535, F1: (v63 | 0), F2: 0}
            for fi93 := 0; fi93 < 16; fi93++ {
                var v94 int16 = -687
                _ = fi93
                _ = v94
            }
            _ = v92
        } else {
            switch "hello" {
            case "":
                for fi95 := 0; fi95 < 16; fi95++ {
                    fmt.Println(v63)
                    _ = fi95
                }
            case "hello":
                switch v58 {
                case "a":
                    if ((v1 ^ 4294967295) <= (v1 << ((338) & 63))) {
                        v59 = v59
                    } else {
                        if false {
                            if false {
                                var v96 bool = (!(false && (v88 > v1)))
                                for fi97 := 0; fi97 < 19; fi97++ {
                                    v89 += v89
                                    var v98 int8 = v89
                                    if (int(v63) > (fi97 / (((fi97 / ((fi97) | 1))) | 1))) {
                                        fmt.Println(v4)
                                    }
                                    _ = fi97
                                    _ = v98
                                }
                                _ = v96
                            } else {
                                var v99 float64 = 0.1
                                fmt.Println(v59)
                                v100 := GenS6{F0: 0, F1: 9223372036854775807, F2: 18446744073709551615}
                                _ = v100
                                _ = v99
                            }
                            v101 := GenS1{F0: v63}
                            _ = len(v5)
                            _ = v101
                        }
                        switch 2147483647 {
                        case 9223372036854775807:
                            var v102 int64 = -170
                            _ = v102
                        case 1:
                            var v103 int32 = -2147483648
                            _ = v103
                        default:
                            for fi104 := 0; fi104 < 8; fi104++ {
                                var v105 int64 = -9223372036854775808
                                _ = fi104
                                _ = v105
                            }
                        }
                        v59 = (v63 != v63)
                    }
                case "foo bar":
                    switch v3 {
                    case "":
                        fmt.Println(v1)
                    }
                case "hello":
                    if (true && v59) {
                        if v59 {
                            for fi106 := 0; fi106 < 18; fi106++ {
                                if false {
                                    var v107 uint8 = v63
                                    _ = v107
                                }
                                _ = fi106
                            }
                            var v108 int = genFn1()
                            var v109 uint = 4294967295
                            _ = v108
                            _ = v109
                        }
                        delete(v6, v58)
                    }
                default:
                    if (false || ((v63 % (((v63 / ((91) | 1))) | 1)) <= (v63 ^ 128))) {
                        fmt.Println(v58, v4, v89)
                        v4 = v4
                    }
                }
            case "foo bar":
                if v59 {
                    var v110 int = 816
                    var v111 uint32 = v88
                    _ = v110
                    _ = v111
                } else {
                    v59 = (false && ((18 >= 520) && (5 < 156)))
                    var v112 uint = 4294967295
                    if v59 {
                        v3 = v3
                        v6[v58] = (v1 >> (((v112 + 1)) & 63))
                    } else {
                        if (!v59) {
                            v4 = "foo bar"
                            var v113 int64 = -629
                            v4 = v3
                            _ = v113
                        } else {
                            switch "a" {
                            case "a":
                                for fi114 := 0; fi114 < 7; fi114++ {
                                    var v115 uint32 = (v88 ^ v1)
                                    _ = fi114
                                    _ = v115
                                }
                            }
                            if ((v1 / (((v88 / (((v1 * 304)) | 1))) | 1)) == v1) {
                                for fi116 := 0; fi116 < 18; fi116++ {
                                    v89 += (v89 / (((v89 / ((-1) | 1))) | 1))
                                    if len(v5) > 0 {
                                        v5[0] = v88
                                    }
                                    _ = fi116
                                }
                                _ = len(v5)
                                for fi117 := 0; fi117 < 1; fi117++ {
                                    v118 := GenS16{F0: 0, F1: (!true)}
                                    var v119 uint32 = (v1 % (((v88 - v1)) | 1))
                                    v2 += v2
                                    _ = fi117
                                    _ = v118
                                    _ = v119
                                }
                            }
                        }
                    }
                    _ = v112
                }
            }
        }
        v120 := GenS2{F0: uint32(v63), F1: ((("\n\t" + v4) + v3) + v4)}
        if len(v5) > 0 {
            v5[0] = v88
        }
        _ = fi91
        _ = v120
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v63:", v63)
    fmt.Println("v88:", v88)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v58
    _ = v59
    _ = v6
    _ = v62
    _ = v63
    _ = v88
    _ = v89
    _ = v90
}
