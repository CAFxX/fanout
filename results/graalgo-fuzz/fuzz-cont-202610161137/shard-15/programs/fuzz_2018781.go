package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 int64; F1 bool; F2 uint8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 bool; F1 int; F2 int8 }
type GenS7 struct { F0 uint8; F1 uint8 }
type GenS8 struct { F0 uint16; F1 int8; F2 int16 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 bool; F1 uint }
type GenS11 struct { F0 int64; F1 int8; F2 bool }
type GenS12 struct { F0 uint64; F1 uint32; F2 uint16 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 string; F1 uint32; F2 uint16 }
type GenS15 struct { F0 string; F1 int16; F2 uint32 }
type GenS16 struct { F0 string; F1 uint16; F2 uint64 }
type GenS17 struct { F0 int32; F1 uint32; F2 bool }
type GenS18 struct { F0 bool; F1 string }
type GenS19 struct { F0 uint32; F1 int16; F2 int8 }
func genFn0(p0 int32) uint8 {
    v6 := GenS1{F0: (p0 ^ p0)}
    var v7 float64 = 1.5
    v7 += ((((v7 * ((-2.25 - (0.1 - (1.5 + v7))) * ((v7 + float64(-2.25)) + v7))) * float64(float32(0.0))) + 1e10) + (float64(float32((v7 + v7))) + 1e10))
    switch -400 {
    case 2147483647:
        v8 := []int{794, 2147483647, 0}
        _ = v8
    case 738:
        var v9 int32 = (p0 << ((42) & 63))
        _ = v9
    }
    _ = p0
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 uint8, p1 uint64) uint64 {
    v10 := GenS2{F0: uint32(p0), F1: -9223372036854775808}
    v11 := GenS1{F0: -501}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return (p1 % (((p1 & (p1 << ((0) & 63)))) | 1))
}

func main() {
    var v1 uint8 = 255
    var v2 uint64 = 921
    var v3 int16 = 1
    var v4 uint = 728
    v5 := make(map[string]string)
    v12 := []int{int(v3), 2147483647, -9223372036854775808}
    var v13 int = 0
    var v14 string = ("hello" + "a")
    for fi15 := 0; fi15 < 4; fi15++ {
        if true {
            if (1 >= 255) {
                var v16 uint16 = 42
                v17 := GenS3{F0: 0, F1: (!(!false)), F2: v1}
                _ = v16
                _ = v17
            }
            if (0 < -579) {
                v18 := GenS4{F0: (v2 % (((v2 - 822)) | 1))}
                var v19 int32 = -334
                for fi20 := 0; fi20 < 18; fi20++ {
                    if false {
                        var v21 float64 = -2.25
                        var v22 uint32 = 23
                        var v23 float64 = 1e10
                        _ = v21
                        _ = v22
                        _ = v23
                    } else {
                        var v24 int64 = -1
                        if (!(2 < 4294967295)) {
                            var v25 bool = true
                            _ = v25
                        } else {
                            if false {
                                v26 := GenS5{F0: 305}
                                v19 = (v19 + 42)
                                v2 += v2
                                _ = v26
                            }
                            if ((-2.25 - 1e10) <= 0.0) {
                                v12 = append(v12, fi20)
                                if (v3 <= (^(v3))) {
                                    if false {
                                        v12 = append(v12, (v13 << ((361) & 63)))
                                        switch fi15 {
                                        case 578:
                                            if len(v12) >= 2 {
                                                v12 = v12[1:len(v12)]
                                            }
                                        }
                                        var v27 string = v14
                                        _ = v27
                                    }
                                    fmt.Println(v13)
                                } else {
                                    if false {
                                        var v28 uint64 = 1
                                        _ = v28
                                    } else {
                                        v1 -= 0
                                        var v29 float64 = 3.14159
                                        switch (fi20 & -786) {
                                        case 127:
                                            if ((v24 >> (((v4 << ((v4) & 63))) & 63)) >= (v24 + 42)) {
                                                v29 -= (((v29 * 1e10) / (v29 / v29)) * float64((1e10 / float32(0.1))))
                                                v30 := v5[(v14 + "\n\t")]
                                                v31 := GenS6{F0: ((v19 << ((42) & 63)) > 42), F1: v13, F2: -1}
                                                if (true && true) {
                                                    delete(v5, v30)
                                                    v13 -= int(v1)
                                                    delete(v5, (("" + (((v30 + (((v14 + ("\n\t" + "")) + v30) + v30)) + (v14 + (("" + ((v30 + "\n\t") + v30)) + v30))) + v30)) + v14))
                                                } else {
                                                    v14 = ("\n\t" + (("hello" + v14) + v14))
                                                    v32 := GenS7{F0: v1, F1: 255}
                                                    _ = v32
                                                }
                                                _ = v30
                                                _ = v31
                                            }
                                        case -688:
                                            switch v14 {
                                            case "":
                                                var v33 uint32 = 503
                                                _ = v33
                                            case "a":
                                                fmt.Println(v14, v24, v1)
                                            }
                                        default:
                                            if len(v12) > 0 {
                                                v12[0] = 128
                                            }
                                        }
                                        _ = v29
                                    }
                                }
                                var v34 int8 = 79
                                _ = v34
                            }
                            var v35 float32 = -2.25
                            _ = v35
                        }
                        _ = v24
                    }
                    _ = fi20
                }
                _ = v18
                _ = v19
            } else {
                delete(v5, v14)
                if false {
                    switch ("foo bar" + "foo bar") {
                    case "foo bar":
                        for fi36 := 0; fi36 < 13; fi36++ {
                            switch (fi36 >> ((v4) & 63)) {
                            case 128:
                                if true {
                                    if true {
                                        if false {
                                            switch fi36 {
                                            case -115:
                                                if (v3 > v3) {
                                                    v37 := GenS8{F0: 2, F1: 100, F2: int16(v4)}
                                                    v3 -= v3
                                                    v5[v14] = "hello"
                                                    _ = v37
                                                } else {
                                                    if len(v12) >= 2 {
                                                        v12 = v12[1:len(v12)]
                                                    }
                                                }
                                            case 32767:
                                                if (((true && false) && ((!((v3 << ((v4) & 63)) == (^(v3)))) || false)) || true) {
                                                    var v38 float32 = 1.5
                                                    if true {
                                                        switch 510 {
                                                        case 9223372036854775807:
                                                            var v39 string = "foo bar"
                                                            v5[("foo bar" + ("foo bar" + v39))] = v39
                                                            _ = v39
                                                        }
                                                        if false {
                                                            v38 -= (v38 / 3.14159)
                                                            v13 = 4294967295
                                                        } else {
                                                            _ = len(v12)
                                                            v40 := GenS8{F0: uint16(v4), F1: -78, F2: (v3 % (((v3 << (((^(v4))) & 63))) | 1))}
                                                            v5[(("a" + v14) + "a")] = v14
                                                            _ = v40
                                                        }
                                                        fmt.Println(fi15)
                                                    }
                                                    _ = v38
                                                }
                                            }
                                        }
                                        if true {
                                            var v41 float64 = float64(float32(-2.25))
                                            _ = v41
                                        } else {
                                            _ = len(v5)
                                            if false {
                                                var v42 uint8 = v1
                                                _ = v42
                                            } else {
                                                switch fi15 {
                                                case -9223372036854775808:
                                                    var v43 uint32 = 2147483647
                                                    _ = v43
                                                case 1:
                                                    if false {
                                                        switch -2147483648 {
                                                        case 4294967295:
                                                            var v44 float64 = 0.1
                                                            _ = v44
                                                        case 42:
                                                            switch v14 {
                                                            case "foo bar":
                                                                _ = len(v12)
                                                            case "\n\t":
                                                                var v45 int64 = 886
                                                                _ = v45
                                                            }
                                                        case 2147483647:
                                                            v2 = v2
                                                        }
                                                        v3 = v3
                                                        v13 = (fi36 | fi36)
                                                    }
                                                case 0:
                                                    var v46 int32 = 1
                                                    _ = v46
                                                default:
                                                    var v47 uint8 = (^(v1))
                                                    _ = v47
                                                }
                                            }
                                            var v48 float64 = 0.0
                                            fmt.Println(v4, v1, v4)
                                            _ = v48
                                        }
                                    } else {
                                        if (57 >= 321) {
                                            fmt.Println(v4, v4)
                                            var v49 uint32 = 2
                                            v4 += (v4 + 357)
                                            _ = v49
                                        } else {
                                            if (!false) {
                                                if false {
                                                    v1 = genFn0(-1)
                                                }
                                            }
                                            if ((!false) && (127 < 32767)) {
                                                if true {
                                                    if (!(!((65535 >= uint16(v4)) || false))) {
                                                        v12 = append(v12, fi36)
                                                        if true {
                                                            v50 := GenS9{F0: v4}
                                                            delete(v5, v14)
                                                            v2 = v2
                                                            _ = v50
                                                        } else {
                                                            var v51 uint = (v4 - (v4 / ((42) | 1)))
                                                            _ = len(v5)
                                                            _ = v51
                                                        }
                                                        if (!true) {
                                                            if true {
                                                                if true {
                                                                    _ = len(v12)
                                                                    var v52 int16 = 32767
                                                                    var v53 uint = (v4 & (v4 % (((v4 - v4)) | 1)))
                                                                    _ = v52
                                                                    _ = v53
                                                                }
                                                            }
                                                            v12 = append(v12, -9223372036854775808)
                                                            var v54 uint = v4
                                                            _ = v54
                                                        } else {
                                                            fmt.Println(v2)
                                                            var v55 int8 = -128
                                                            _ = v55
                                                        }
                                                        switch v14 {
                                                        case "":
                                                            if (v2 <= v2) {
                                                                v56 := v5[""]
                                                                if false {
                                                                    var v57 uint8 = (v1 % (((v1 / ((genFn0(-968)) | 1))) | 1))
                                                                    var v58 uint = (v4 % ((v4) | 1))
                                                                    v13 += fi36
                                                                    _ = v57
                                                                    _ = v58
                                                                }
                                                                if ((^(v2)) > v2) {
                                                                    var v59 int32 = 0
                                                                    var v60 int16 = v3
                                                                    if false {
                                                                        v2 -= 1
                                                                        var v61 bool = (222 >= 2)
                                                                        _ = v61
                                                                    }
                                                                    _ = v59
                                                                    _ = v60
                                                                }
                                                                _ = v56
                                                            } else {
                                                                fmt.Println(v4, v14)
                                                                var v62 int32 = -2147483648
                                                                if len(v12) >= 2 {
                                                                    v12 = v12[1:len(v12)]
                                                                }
                                                                v13 = -9223372036854775808
                                                                _ = v62
                                                            }
                                                        case "\n\t":
                                                            var v63 int16 = -860
                                                            _ = v63
                                                        }
                                                        fmt.Println(v14)
                                                        v64 := GenS5{F0: 42}
                                                        _ = v64
                                                    } else {
                                                        var v65 int16 = 42
                                                        _ = v65
                                                    }
                                                }
                                                v66 := v5[v14]
                                                var v67 float32 = 3.14159
                                                if true {
                                                    if len(v12) >= 2 {
                                                        v12 = v12[1:len(v12)]
                                                    }
                                                    if len(v12) > 0 {
                                                        v12[0] = fi36
                                                    }
                                                    delete(v5, ("foo bar" + v14))
                                                    fmt.Println(v1, v1, v14)
                                                }
                                                v5[v14] = "hello"
                                                _ = v66
                                                _ = v67
                                            } else {
                                                v68 := GenS10{F0: false, F1: (^(v4))}
                                                fmt.Println(v13, v13)
                                                _ = v68
                                            }
                                            v69 := GenS11{F0: 42, F1: 0, F2: (true && ((uint64(v1) == uint64(v4)) || true))}
                                            fmt.Println(v1)
                                            _ = v69
                                        }
                                        fmt.Println(v4, v13)
                                    }
                                    var v70 float32 = 0.0
                                    if (255 >= 1) {
                                        var v71 int32 = 421
                                        v3 += 1000
                                        _ = v71
                                    } else {
                                        if false {
                                            switch v14 {
                                            case "a":
                                                v72 := GenS4{F0: genFn1(v1, (v2 ^ (v2 + 4294967295)))}
                                                _ = v72
                                            case "":
                                                _ = len(v5)
                                            default:
                                                v73 := GenS1{F0: 130}
                                                _ = v73
                                            }
                                            v3 += int16(v4)
                                            var v74 uint = (v4 << (((^(v4))) & 63))
                                            v75 := GenS12{F0: (v2 | genFn1((v1 + v1), (v2 << ((4294967295) & 63)))), F1: 0, F2: 32767}
                                            v1 -= 127
                                            _ = v74
                                            _ = v75
                                        } else {
                                            switch 9223372036854775807 {
                                            case 32767:
                                                v4 -= (^(v4))
                                            case -389:
                                                v1 = v1
                                            case 277:
                                                switch (fi36 / ((fi36) | 1)) {
                                                case 204:
                                                    var v76 bool = (0 >= -30)
                                                    _ = v76
                                                case 128:
                                                    fmt.Println(v70)
                                                case 669:
                                                    if len(v12) >= 2 {
                                                        v12 = v12[1:len(v12)]
                                                    }
                                                default:
                                                    v14 = "hello"
                                                }
                                            }
                                            v77 := GenS13{F0: (v2 >> ((2) & 63))}
                                            if (false && ((fi15 | fi36) <= fi15)) {
                                                var v78 int8 = 39
                                                _ = v78
                                            }
                                            _ = v77
                                        }
                                        v1 += 255
                                        var v79 int16 = (v3 * (v3 | (v3 ^ (v3 - v3))))
                                        var v80 uint32 = 500
                                        _ = v79
                                        _ = v80
                                    }
                                    v3 += -32768
                                    _ = v70
                                } else {
                                    fmt.Println(v4)
                                }
                            default:
                                if false {
                                    v13 = (v13 & fi15)
                                    switch (fi36 + v13) {
                                    case -75:
                                        v14 = v14
                                    case 502:
                                        fmt.Println(fi15, fi36, fi36)
                                    }
                                }
                            }
                            v81 := GenS14{F0: v14, F1: 42, F2: 715}
                            _ = fi36
                            _ = v81
                        }
                    default:
                        var v82 int16 = (-(v3))
                        _ = v82
                    }
                    for fi83 := 0; fi83 < 9; fi83++ {
                        fmt.Println(v14)
                        _ = len(v5)
                        var v84 float32 = (float32(0.1) * 1.5)
                        _ = fi83
                        _ = v84
                    }
                }
            }
            for fi85 := 0; fi85 < 5; fi85++ {
                var v86 int8 = 0
                var v87 int8 = 127
                v88 := GenS5{F0: 357}
                _ = fi85
                _ = v86
                _ = v87
                _ = v88
            }
        }
        switch v13 {
        case 127:
            var v89 uint64 = (v2 * v2)
            _ = v89
        default:
            v12 = append(v12, -2147483648)
        }
        _ = fi15
    }
    v90 := GenS15{F0: (((v14 + ("hello" + v14)) + v14) + v14), F1: (v3 | 32767), F2: 298}
    v12 = append(v12, v13)
    var v91 int16 = v3
    if (!true) {
        var v92 int32 = 42
        v93 := GenS16{F0: v14, F1: 791, F2: v2}
        v94 := GenS17{F0: (v92 * 42), F1: 65535, F2: false}
        _ = v92
        _ = v93
        _ = v94
    } else {
        v95 := GenS18{F0: (uint64(v4) != 2), F1: v14}
        _ = v95
    }
    if ((42 == (v4 ^ (v4 ^ 65535))) && true) {
        var v96 int8 = 0
        _ = v96
    } else {
        fmt.Println(v2, v3)
        for fi97 := 0; fi97 < 1; fi97++ {
            v98 := GenS19{F0: 1, F1: v3, F2: 127}
            v5["a"] = v14
            _ = fi97
            _ = v98
        }
        fmt.Println(v13, v13, v2)
    }
    if false {
        for fi99 := 0; fi99 < 7; fi99++ {
            v5[((v14 + "") + (v14 + "\n\t"))] = v14
            _ = fi99
        }
        v100 := GenS9{F0: (v4 & (v4 >> ((uint(v1)) & 63)))}
        _ = v100
    } else {
        if true {
            v101 := GenS18{F0: true, F1: v14}
            var v102 uint32 = 1
            _ = v101
            _ = v102
        } else {
            v103 := v5[v14]
            for fi104 := 0; fi104 < 18; fi104++ {
                fmt.Println(v4)
                _ = fi104
            }
            v105 := GenS1{F0: 903}
            _ = v103
            _ = v105
        }
    }
    v106 := GenS12{F0: (^(v2)), F1: 1, F2: uint16(v4)}
    v13 = v13
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v91:", v91)
    _ = v1
    _ = v106
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v90
    _ = v91
}
