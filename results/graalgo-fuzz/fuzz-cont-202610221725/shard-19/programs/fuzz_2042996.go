package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int; F1 uint; F2 int32 }
type GenS6 struct { F0 bool; F1 int8; F2 uint16 }
type GenS7 struct { F0 int8; F1 int }
type GenS8 struct { F0 int8; F1 int32 }
type GenS9 struct { F0 uint64; F1 uint16 }
type GenS10 struct { F0 int8; F1 int16; F2 int64 }
type GenS11 struct { F0 uint; F1 uint64; F2 int32 }
type GenS12 struct { F0 int }
type GenS13 struct { F0 uint32; F1 uint8; F2 int64 }
type GenS14 struct { F0 int; F1 uint; F2 int32 }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 int }
func genFn0() uint16 {
    if true {
        if false {
            if (!true) {
                var v6 uint32 = 1
                _ = v6
            }
        }
    } else {
        var v7 uint8 = 184
        _ = v7
    }
    var v8 bool = (false || (512 != 0))
    v9 := []int{128}
    fmt.Println(v8)
    _ = v8
    _ = v9
    return 639
}

func main() {
    var v1 uint16 = 2
    var v2 uint16 = v1
    v3 := make(map[uint32]string)
    v4 := []string{"a", "hello"}
    var v5 bool = true
    delete(v3, 65535)
    for fi10 := 0; fi10 < 10; fi10++ {
        v1 -= v2
        if (v5 || v5) {
            for fi11 := 0; fi11 < 15; fi11++ {
                var v12 string = "foo bar"
                fmt.Println(v5)
                delete(v3, 2147483647)
                var v13 uint64 = 376
                _ = fi11
                _ = v12
                _ = v13
            }
            v3[2147483648] = "a"
        } else {
            if v5 {
                if true {
                    v4 = append(v4, "a")
                    switch "" {
                    case "hello":
                        delete(v3, 65535)
                    case "\n\t":
                        var v14 int16 = 1000
                        _ = v14
                    }
                    for fi15 := 0; fi15 < 19; fi15++ {
                        _ = len(v3)
                        _ = fi15
                    }
                } else {
                    v2 -= v1
                    v1 = genFn0()
                }
            }
            v16 := GenS1{F0: 1}
            if v5 {
                v17 := GenS2{F0: 0, F1: 42}
                _ = v17
            }
            _ = v16
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi10
    }
    var v18 bool = v5
    v4 = append(v4, "hello")
    if v18 {
        _ = len(v3)
    } else {
        v2 += 0
    }
    v19 := v3[726]
    v2 = (v2 / (((v1 / (((v1 % ((v1) | 1))) | 1))) | 1))
    if ((618 <= 347) && (v19 == v19)) {
        fmt.Println(v5, v18)
        var v20 float64 = (0.0 - 0.0)
        _ = v20
    } else {
        for fi21 := 0; fi21 < 19; fi21++ {
            for fi22 := 0; fi22 < 11; fi22++ {
                if (((-9223372036854775808 <= 365) || (1000 < -1)) || (!(389 == 536))) {
                    v23 := GenS3{F0: -128}
                    _ = v23
                } else {
                    v5 = v18
                    if false {
                        if (!((100 > 1) && true)) {
                            switch fi22 {
                            case -585:
                                var v24 uint8 = 128
                                _ = v24
                            case 498:
                                var v25 int16 = 1
                                _ = v25
                            default:
                                switch "hello" {
                                case "hello":
                                    fmt.Println(fi22)
                                case "a":
                                    switch v19 {
                                    case "foo bar":
                                        fmt.Println(fi22, fi21, v5)
                                    case "\n\t":
                                        var v26 int = int(v2)
                                        _ = v26
                                    default:
                                        if v18 {
                                            if v18 {
                                                v1 = 32767
                                                v27 := v3[2147483648]
                                                if (fi22 <= (fi22 | (fi21 >> ((0) & 63)))) {
                                                    v2 -= (v1 >> ((741) & 63))
                                                    var v28 int = fi21
                                                    _ = v28
                                                } else {
                                                    v18 = (((v5 || (255 <= 255)) && false) || v18)
                                                }
                                                switch ("foo bar" + ((v27 + "a") + "a")) {
                                                case "":
                                                    v29 := v3[2]
                                                    _ = v29
                                                }
                                                _ = v27
                                            } else {
                                                switch v19 {
                                                case "a":
                                                    v19 = "\n\t"
                                                default:
                                                    fmt.Println(v19, v5, v1)
                                                }
                                                v19 = v19
                                            }
                                            v1 = 32767
                                            v30 := GenS4{F0: 42}
                                            _ = v30
                                        } else {
                                            v19 = v19
                                            v31 := GenS2{F0: 986, F1: 18446744073709551615}
                                            _ = v31
                                        }
                                        v18 = v18
                                        var v32 uint = 391
                                        _ = v32
                                    }
                                }
                            }
                            v33 := v3[4294967295]
                            switch fi22 {
                            case 310:
                                if true {
                                    v34 := GenS5{F0: (fi21 | -441), F1: 255, F2: 285}
                                    if (42 < -301) {
                                        v35 := GenS6{F0: v5, F1: -5, F2: (v1 & (^(v1)))}
                                        if len(v4) > 0 {
                                            v4[0] = v19
                                        }
                                        _ = v35
                                    } else {
                                        if v18 {
                                            var v36 uint = 2
                                            v33 = "foo bar"
                                            var v37 int = fi21
                                            _ = v36
                                            _ = v37
                                        } else {
                                            if true {
                                                v38 := GenS4{F0: 425}
                                                var v39 bool = true
                                                _ = v38
                                                _ = v39
                                            } else {
                                                if true {
                                                    var v40 uint32 = 42
                                                    fmt.Println(v2, v40, v2)
                                                    _ = v40
                                                } else {
                                                    var v41 uint = 895
                                                    if false {
                                                        fmt.Println(v1)
                                                        if (4294967295 != 42) {
                                                            v41 += (v41 ^ v41)
                                                            v42 := GenS5{F0: fi21, F1: (v41 / (((v41 * (v41 / (((v41 % ((v41) | 1))) | 1)))) | 1)), F2: 2147483647}
                                                            _ = v42
                                                        } else {
                                                            if v18 {
                                                                delete(v3, 2)
                                                                if false {
                                                                    if len(v4) > 0 {
                                                                        v4[0] = v33
                                                                    }
                                                                    if true {
                                                                        if (127 < -1) {
                                                                            v43 := GenS7{F0: 0, F1: int(v1)}
                                                                            if v5 {
                                                                                v4 = append(v4, (v33 + (v19 + v19)))
                                                                                v44 := GenS2{F0: (v41 * 4294967295), F1: 698}
                                                                                v1 -= (^(v1))
                                                                                var v45 uint16 = (v1 >> ((v41) & 63))
                                                                                _ = v44
                                                                                _ = v45
                                                                            }
                                                                            _ = v43
                                                                        } else {
                                                                            v1 += (v2 & v2)
                                                                            var v46 uint16 = (v2 & 42)
                                                                            _ = v46
                                                                        }
                                                                        switch fi21 {
                                                                        case 32767:
                                                                            _ = len(v3)
                                                                        case 182:
                                                                            v47 := GenS8{F0: 0, F1: 1}
                                                                            switch (fi21 << (((v41 / ((18446744073709551615) | 1))) & 63)) {
                                                                            case 4294967295:
                                                                                fmt.Println(v18)
                                                                            case -9223372036854775808:
                                                                                fmt.Println(v5)
                                                                                v48 := GenS4{F0: -32768}
                                                                                v49 := GenS7{F0: 127, F1: (fi22 + fi21)}
                                                                                _ = v48
                                                                                _ = v49
                                                                            case 596:
                                                                                if false {
                                                                                    var v50 uint16 = (v1 / ((v1) | 1))
                                                                                    var v51 uint32 = 2
                                                                                    v33 = (v19 + v33)
                                                                                    v52 := GenS9{F0: 4294967295, F1: (v50 + 255)}
                                                                                    v1 = v2
                                                                                    _ = v50
                                                                                    _ = v51
                                                                                    _ = v52
                                                                                } else {
                                                                                    v33 = v33
                                                                                    var v53 uint8 = 186
                                                                                    var v54 float64 = 1.5
                                                                                    _ = v53
                                                                                    _ = v54
                                                                                }
                                                                            default:
                                                                                var v55 int = 0
                                                                                _ = v55
                                                                            }
                                                                            var v56 uint8 = 0
                                                                            _ = v47
                                                                            _ = v56
                                                                        case 9223372036854775807:
                                                                            switch (v19 + v33) {
                                                                            case "a":
                                                                                _ = len(v4)
                                                                            case "\n\t":
                                                                                switch v19 {
                                                                                case "":
                                                                                    v57 := GenS10{F0: 100, F1: -93, F2: -243}
                                                                                    _ = v57
                                                                                default:
                                                                                    v41 += 846
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                    var v58 float32 = ((0.1 / 1.0) / 1e10)
                                                                    switch fi22 {
                                                                    case 127:
                                                                        var v59 int = 2147483647
                                                                        _ = v59
                                                                    default:
                                                                        var v60 float32 = 1.5
                                                                        _ = v60
                                                                    }
                                                                    v61 := v3[2147483647]
                                                                    _ = v58
                                                                    _ = v61
                                                                } else {
                                                                    delete(v3, 42)
                                                                    v62 := GenS11{F0: 113, F1: 42, F2: 1}
                                                                    v1 -= (v1 >> ((uint(v1)) & 63))
                                                                    _ = v62
                                                                }
                                                            } else {
                                                                v63 := GenS5{F0: fi21, F1: v41, F2: -211}
                                                                v64 := GenS3{F0: 83}
                                                                var v65 uint64 = 890
                                                                _ = v63
                                                                _ = v64
                                                                _ = v65
                                                            }
                                                            var v66 float32 = ((-2.25 / float32((1e10 / 1e10))) * -2.25)
                                                            _ = v66
                                                        }
                                                        v18 = (true && (2 > 85))
                                                        fmt.Println(v2, fi21)
                                                    } else {
                                                        var v67 uint8 = 128
                                                        _ = v67
                                                    }
                                                    _ = v41
                                                }
                                                var v68 string = "a"
                                                _ = v68
                                            }
                                            v19 = v19
                                        }
                                    }
                                    v69 := v3[867]
                                    _ = v34
                                    _ = v69
                                }
                            case -392:
                                v5 = (128 > 128)
                                if true {
                                    var v70 bool = v5
                                    _ = v70
                                }
                                if (!(v5 || ((!(!false)) && v5))) {
                                    var v71 uint = 343
                                    var v72 float32 = ((1e10 + 0.1) / 0.1)
                                    if v5 {
                                        v73 := GenS12{F0: fi21}
                                        _ = v73
                                    } else {
                                        var v74 float64 = 0.0
                                        if (2 == 42) {
                                            v75 := GenS13{F0: 4294967295, F1: 42, F2: 1}
                                            _ = v75
                                        } else {
                                            v76 := GenS8{F0: 0, F1: 185}
                                            _ = v76
                                        }
                                        if (32767 < -1) {
                                            if (1.5 >= (v72 / 3.14159)) {
                                                var v77 int8 = 127
                                                v4 = append(v4, v19)
                                                v78 := GenS14{F0: fi21, F1: v71, F2: 42}
                                                _ = v77
                                                _ = v78
                                            } else {
                                                var v79 uint64 = 0
                                                if v5 {
                                                    switch v19 {
                                                    case "a":
                                                        v74 += 0.1
                                                        var v80 int16 = int16(v2)
                                                        _ = v80
                                                    default:
                                                        if false {
                                                            v18 = ((0 != 0) || (v5 && ((true || (false || false)) && false)))
                                                        }
                                                    }
                                                    if v5 {
                                                        v19 = ((v19 + v19) + v33)
                                                        var v81 int = (fi22 / ((0) | 1))
                                                        _ = v81
                                                    }
                                                    v82 := GenS15{F0: 0}
                                                    _ = v82
                                                }
                                                _ = v79
                                            }
                                            v74 += v74
                                            v74 = (float64((v72 - 0.0)) + (v74 - (v74 * ((v74 * float64(1e10)) * v74))))
                                        }
                                        if (("foo bar" + v19) <= v33) {
                                            fmt.Println(v71, v19, v19)
                                            var v83 int32 = -320
                                            _ = v83
                                        }
                                        if len(v4) > 0 {
                                            v4[0] = v19
                                        }
                                        _ = v74
                                    }
                                    _ = v71
                                    _ = v72
                                }
                            }
                            if (18446744073709551615 > 1) {
                                v18 = (true || true)
                            }
                            _ = v33
                        }
                        v18 = false
                        v84 := GenS16{F0: (^(fi22))}
                        fmt.Println(v19)
                        _ = v84
                    } else {
                        var v85 float32 = 3.14159
                        v5 = (("" + v19) != v19)
                        var v86 uint8 = 128
                        _ = v85
                        _ = v86
                    }
                }
                v87 := v3[2]
                if (fi22 >= (fi21 >> ((0) & 63))) {
                    var v88 uint = 255
                    _ = v88
                } else {
                    var v89 int16 = 1
                    v87 = "hello"
                    v5 = v5
                    v90 := v3[2]
                    _ = v89
                    _ = v90
                }
                _ = fi22
                _ = v87
            }
            v5 = (!(652 >= 906))
            _ = fi21
        }
    }
    v1 -= genFn0()
    var v91 uint8 = 2
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v91
}
