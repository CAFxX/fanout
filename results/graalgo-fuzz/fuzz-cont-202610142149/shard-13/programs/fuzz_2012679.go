package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 uint32; F1 bool; F2 uint }
type GenS4 struct { F0 int64; F1 uint32 }
type GenS5 struct { F0 uint; F1 uint64; F2 bool }
type GenS6 struct { F0 uint32; F1 bool }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int32; F1 int16; F2 bool }
type GenS9 struct { F0 bool; F1 int32 }
type GenS10 struct { F0 int; F1 uint16; F2 uint16 }
type GenS11 struct { F0 uint64; F1 uint8 }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 uint; F1 uint32; F2 int32 }
type GenS14 struct { F0 uint64 }
func genFn0() int {
    for fi6 := 0; fi6 < 6; fi6++ {
        v7 := GenS1{F0: 950, F1: 181}
        _ = fi6
        _ = v7
    }
    switch 0 {
    case 9223372036854775807:
        var v8 float32 = 0.0
        _ = v8
    case -710:
        if false {
            v9 := GenS1{F0: 150, F1: 255}
            var v10 uint8 = 236
            var v11 uint8 = (v10 * (v10 / ((255) | 1)))
            _ = v10
            _ = v11
            _ = v9
        } else {
            v12 := GenS1{F0: 4294967295, F1: 65535}
            v13 := []int{2147483647, 0, 402, 128}
            _ = v12
            _ = v13
        }
    }
    var v14 int8 = 0
    _ = v14
    return -828
}
func genFn1(p0 string) uint {
    v15 := []int{-2147483648}
    var v16 uint64 = 0
    v16 = 4294967295
    for fi17 := 0; fi17 < 19; fi17++ {
        var v18 float32 = 1.5
        for fi19 := 0; fi19 < 1; fi19++ {
            fmt.Println(p0, v18)
            if true {
                v18 = float32(-2.25)
                if len(v15) > 0 {
                    v15[0] = (fi19 ^ 136)
                }
            }
            _ = fi19
        }
        var v20 float64 = 1.5
        _ = fi17
        _ = v18
        _ = v20
    }
    var v21 bool = false
    _ = p0
    _ = v15
    _ = v16
    _ = v21
    return 42
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint = 65535
    v3 := make(map[uint32]uint32)
    v4 := []uint32{4294967295, 42, 0}
    v5 := make(map[int]string)
    for fi22 := 0; fi22 < 15; fi22++ {
        _ = len(v3)
        v4 = append(v4, 65535)
        _ = len(v1)
        _ = fi22
    }
    switch ("hello" + "a") {
    case "a":
        v23 := GenS1{F0: 4294967295, F1: v2}
        _ = v23
    case "foo bar":
        var v24 int8 = 0
        _ = v24
    case "":
        var v25 float32 = 0.0
        _ = v25
    }
    v26 := GenS1{F0: 39, F1: (^(v2))}
    if (42 < 868) {
        v2 = (v2 ^ (v2 | (v2 % ((genFn1(("foo bar" + ""))) | 1))))
        var v27 uint = v2
        if ((v27 - v27) >= (v2 ^ 832)) {
            v28 := GenS2{F0: "", F1: 127}
            switch "foo bar" {
            case "\n\t":
                if true {
                    v27 = genFn1("hello")
                    var v29 int = -9223372036854775808
                    _ = v29
                }
            case "a":
                var v30 uint32 = 65535
                _ = v30
            case "hello":
                v27 += v27
            }
            for fi31 := 0; fi31 < 13; fi31++ {
                if (!(!false)) {
                    v32 := GenS1{F0: 65535, F1: (v2 ^ (v2 - v27))}
                    v2 = genFn1("a")
                    for fi33 := 0; fi33 < 19; fi33++ {
                        var v34 uint16 = 194
                        switch "\n\t" {
                        case "":
                            fmt.Println(fi33, v2)
                        case "foo bar":
                            if true {
                                fmt.Println(fi31, v2, v2)
                                _ = len(v3)
                            } else {
                                var v35 float64 = -2.25
                                v36 := v1["foo bar"]
                                fmt.Println(fi31)
                                _ = v35
                                _ = v36
                            }
                        }
                        _ = fi33
                        _ = v34
                    }
                    _ = v32
                }
                _ = fi31
            }
            _ = v28
        } else {
            for fi37 := 0; fi37 < 2; fi37++ {
                for fi38 := 0; fi38 < 5; fi38++ {
                    _ = len(v1)
                    switch "\n\t" {
                    case "":
                        v39 := GenS1{F0: 280, F1: v27}
                        _ = v39
                    case "hello":
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    default:
                        switch "foo bar" {
                        case "\n\t":
                            fmt.Println(v27)
                        case "a":
                            if true {
                                v40 := GenS2{F0: "a", F1: 0}
                                v27 = v2
                                _ = v40
                            }
                        case "foo bar":
                            var v41 uint64 = 2
                            var v42 int = fi37
                            v43 := GenS1{F0: 2, F1: (v27 | (v27 - v2))}
                            _ = v41
                            _ = v42
                            _ = v43
                        default:
                            switch fi38 {
                            case 501:
                                v44 := GenS3{F0: 2147483648, F1: true, F2: (v2 + (v2 & v2))}
                                _ = v44
                            case -2147483648:
                                var v45 int64 = int64(fi38)
                                _ = v45
                            default:
                                fmt.Println(v2, fi37, fi38)
                            }
                        }
                    }
                    v27 = genFn1(("\n\t" + "\n\t"))
                    _ = len(v3)
                    v2 -= genFn1("a")
                    if false {
                        if len(v4) > 0 {
                            v4[0] = 2147483647
                        }
                        v27 = v27
                    }
                    _ = fi38
                }
                var v46 bool = (!(944 < 42))
                var v47 int64 = -34
                _ = fi37
                _ = v46
                _ = v47
            }
            switch 1 {
            case 42:
                var v48 int16 = 99
                _ = v48
            case 32767:
                if false {
                    if false {
                        v2 = 255
                    } else {
                        _ = len(v4)
                        var v49 uint32 = 252
                        var v50 string = ("\n\t" + "")
                        _ = v49
                        _ = v50
                    }
                } else {
                    v27 = (v27 ^ 18446744073709551615)
                    if false {
                        v27 += 65535
                    } else {
                        fmt.Println(v2, v2)
                        var v51 int8 = 1
                        if ((-128 >= v51) || (0 < 197)) {
                            var v52 uint64 = 42
                            v53 := GenS4{F0: -603, F1: 415}
                            _ = v52
                            _ = v53
                        }
                        _ = v51
                    }
                    v2 += (v2 - (v27 << ((18446744073709551615) & 63)))
                }
            }
            switch 211 {
            case 986:
                v54 := GenS5{F0: v27, F1: 0, F2: true}
                _ = v54
            case -9223372036854775808:
                if false {
                    if false {
                        v55 := GenS3{F0: 4294967295, F1: ((0.1 != float64(1e10)) || (2 <= 802)), F2: v27}
                        if true {
                            switch "hello" {
                            case "hello":
                                if (((0.1 * 3.14159) / 1.5) <= -2.25) {
                                    _ = len(v5)
                                    v27 = 1
                                    var v56 int = -2147483648
                                    _ = v56
                                } else {
                                    delete(v5, 4294967295)
                                }
                            case "":
                                fmt.Println(v2)
                            case "foo bar":
                                fmt.Println(v27, v27)
                            default:
                                if true {
                                    var v57 int16 = 99
                                    fmt.Println(v57, v27, v2)
                                    _ = v57
                                } else {
                                    var v58 uint8 = 203
                                    v58 += 0
                                    _ = v58
                                }
                            }
                            for fi59 := 0; fi59 < 9; fi59++ {
                                var v60 uint = v2
                                v2 += (^(v2))
                                switch ("foo bar" + ("" + ("" + "foo bar"))) {
                                case "a":
                                    for fi61 := 0; fi61 < 2; fi61++ {
                                        v62 := GenS6{F0: 2147483647, F1: (71 < 0)}
                                        _ = fi61
                                        _ = v62
                                    }
                                case "foo bar":
                                    if (371 != -1) {
                                        v3[0] = 42
                                    } else {
                                        if false {
                                            if false {
                                                delete(v3, 1)
                                                v63 := v5[(^(fi59))]
                                                _ = v63
                                            }
                                            if (!(176 >= 42)) {
                                                v64 := GenS7{F0: uint16(v60)}
                                                v65 := GenS8{F0: 0, F1: -1, F2: (!false)}
                                                var v66 string = "foo bar"
                                                _ = v64
                                                _ = v65
                                                _ = v66
                                            }
                                        } else {
                                            if false {
                                                v60 = (^(v60))
                                            }
                                        }
                                        for fi67 := 0; fi67 < 20; fi67++ {
                                            _ = len(v4)
                                            var v68 int8 = 100
                                            v27 -= (^(v60))
                                            _ = fi67
                                            _ = v68
                                        }
                                        if false {
                                            fmt.Println(v27, v27, v27)
                                            var v69 int16 = 488
                                            _ = v69
                                        } else {
                                            v70 := GenS9{F0: ((true && (fi59 > (fi59 / ((genFn0()) | 1)))) || (!false)), F1: 819}
                                            switch fi59 {
                                            case -2147483648:
                                                v27 += v27
                                            case -306:
                                                v71 := GenS2{F0: "\n\t", F1: (-(fi59))}
                                                _ = v71
                                            case -9223372036854775808:
                                                var v72 int16 = 1000
                                                _ = v72
                                            default:
                                                v27 -= v2
                                            }
                                            _ = v70
                                        }
                                    }
                                case "":
                                    if true {
                                        v73 := GenS10{F0: (fi59 | (-(fi59))), F1: 32767, F2: 344}
                                        for fi74 := 0; fi74 < 16; fi74++ {
                                            fmt.Println(v60, v60, v60)
                                            _ = fi74
                                        }
                                        if (42 != 569) {
                                            var v75 float32 = -2.25
                                            _ = v75
                                        }
                                        _ = v73
                                    } else {
                                        v60 = (v2 + (v60 | genFn1((("hello" + ("" + "")) + ("foo bar" + "")))))
                                        for fi76 := 0; fi76 < 6; fi76++ {
                                            var v77 uint8 = 127
                                            var v78 uint32 = 547
                                            _ = fi76
                                            _ = v77
                                            _ = v78
                                        }
                                        if false {
                                            fmt.Println(v27, v2, fi59)
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            var v79 float32 = -2.25
                                            _ = v79
                                        } else {
                                            switch "foo bar" {
                                            case "a":
                                                var v80 bool = false
                                                _ = v80
                                            default:
                                                if false {
                                                    if len(v4) > 0 {
                                                        v4[0] = 383
                                                    }
                                                    v81 := GenS5{F0: (v27 + 42), F1: 18446744073709551615, F2: false}
                                                    _ = v81
                                                }
                                            }
                                            var v82 int8 = 0
                                            _ = v82
                                        }
                                    }
                                default:
                                    for fi83 := 0; fi83 < 4; fi83++ {
                                        v4 = append(v4, 42)
                                        _ = fi83
                                    }
                                }
                                _ = fi59
                                _ = v60
                            }
                        } else {
                            v2 = (v27 % ((v27) | 1))
                            var v84 int64 = 1
                            v84 = (v84 << ((genFn1("hello")) & 63))
                            _ = v84
                        }
                        var v85 float32 = (((3.14159 - 0.1) - 0.0) * float32(3.14159))
                        _ = v55
                        _ = v85
                    } else {
                        v27 += (v2 / ((4294967295) | 1))
                        v86 := GenS11{F0: 2, F1: 255}
                        _ = v86
                    }
                    fmt.Println(v2, v2)
                } else {
                    v2 += (^(v27))
                    if true {
                        if len(v4) > 0 {
                            v4[0] = 4294967295
                        }
                    } else {
                        v27 -= (v2 & v27)
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v2 += v27
                    }
                    var v87 int32 = -1
                    _ = v87
                }
            }
        }
        _ = v27
    } else {
        var v88 uint32 = 2147483648
        v2 += (v2 * (v2 * 1))
        _ = v88
    }
    if false {
        var v89 int16 = 1000
        if false {
            if true {
                if true {
                    v90 := GenS6{F0: 2147483648, F1: false}
                    v91 := GenS12{F0: 2}
                    for fi92 := 0; fi92 < 18; fi92++ {
                        v2 = v2
                        _ = fi92
                    }
                    _ = v90
                    _ = v91
                }
                v2 = (v2 >> ((v2) & 63))
                var v93 int32 = 49
                _ = v93
            }
            fmt.Println(v2, v89)
            var v94 uint = (v2 + 637)
            _ = v94
        }
        if (666 > 42) {
            if false {
                var v95 string = "foo bar"
                _ = v95
            }
            fmt.Println(v89)
            switch -9223372036854775808 {
            case 2147483647:
                var v96 int16 = v89
                _ = v96
            case 0:
                if false {
                    var v97 uint16 = 42
                    if true {
                        v97 -= (v97 / ((65535) | 1))
                        v98 := GenS3{F0: 4294967295, F1: false, F2: v2}
                        _ = v98
                    } else {
                        v99 := GenS4{F0: 1, F1: 801}
                        var v100 float32 = 3.14159
                        v89 = 0
                        _ = v100
                        _ = v99
                    }
                    _ = v97
                } else {
                    var v101 int16 = v89
                    v102 := GenS13{F0: 255, F1: 0, F2: 1}
                    v89 = (v101 >> ((1) & 63))
                    _ = v101
                    _ = v102
                }
            case 9223372036854775807:
                v89 -= (-(v89))
            }
        }
        _ = v89
    } else {
        v103 := GenS14{F0: 42}
        _ = v103
    }
    v2 += v2
    v4 = append(v4, 0)
    _ = len(v3)
    if (!((3.14159 * 3.14159) <= 1e10)) {
        for fi104 := 0; fi104 < 16; fi104++ {
            if false {
                var v105 bool = ("hello" <= "foo bar")
                v106 := GenS14{F0: 1}
                v105 = false
                _ = v105
                _ = v106
            } else {
                if false {
                    switch ("" + "a") {
                    case "hello":
                        v107 := GenS6{F0: 0, F1: ((-(fi104)) >= (fi104 & 1))}
                        _ = v107
                    }
                    var v108 bool = (1.5 == (float32(0.0) + (((float32(0.1) / (((0.0 - 0.1) + (0.0 * -2.25)) + (1e10 - -2.25))) / 1.5) / 0.1)))
                    v109 := GenS2{F0: ("\n\t" + ("foo bar" + "a")), F1: genFn0()}
                    _ = v108
                    _ = v109
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            _ = fi104
        }
    } else {
        v110 := v3[2147483647]
        if (1 <= 958) {
            _ = len(v5)
            v111 := GenS10{F0: 32768, F1: uint16(v2), F2: 42}
            _ = v111
        } else {
            v110 -= (v110 & v110)
        }
        v112 := GenS13{F0: (^(v2)), F1: 4294967295, F2: -1}
        _ = v110
        _ = v112
    }
    var v113 int8 = 100
    fmt.Println(v2, v113)
    v2 = (^(v2))
    v114 := GenS13{F0: v2, F1: 0, F2: 42}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if false {
        fmt.Println(v113)
        v113 -= -77
    }
    v2 += v2
    v115 := GenS6{F0: 2147483647, F1: (4294967295 < 18446744073709551615)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v113:", v113)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v113
    _ = v114
    _ = v115
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
