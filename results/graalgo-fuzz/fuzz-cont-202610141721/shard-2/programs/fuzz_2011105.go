package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 bool }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 int; F1 int64 }
type GenS5 struct { F0 uint32; F1 uint; F2 uint }
type GenS6 struct { F0 int8; F1 int64 }
type GenS7 struct { F0 uint32; F1 int }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int8; F1 int8 }
type GenS10 struct { F0 int; F1 int32 }
type GenS11 struct { F0 int16; F1 uint32; F2 int16 }
func genFn0() uint8 {
    var v5 uint32 = 225
    v6 := GenS1{F0: ("hello" + "a"), F1: 57, F2: true}
    v5 += v5
    v7 := []int{4294967295}
    _ = v5
    _ = v6
    _ = v7
    return 129
}
func genFn1(p0 uint64, p1 int64, p2 int) uint8 {
    fmt.Println(p1, p0)
    var v8 float64 = ((0.1 * 0.0) + (((3.14159 + ((float64((0.1 * (1.5 - ((0.1 + 1e10) / (-2.25 + (3.14159 / 1.0)))))) + float64(3.14159)) * ((((1e10 + 0.0) / -2.25) + float64((3.14159 - 3.14159))) + (float64((-2.25 * (1.5 + (float32(((3.14159 + (((3.14159 / -2.25) + 1e10) / 1e10)) - -2.25)) * 1.5)))) - (3.14159 + (3.14159 * (-2.25 - (3.14159 / 3.14159)))))))) - float64((1.5 - -2.25))) * (0.1 / 1e10)))
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return genFn0()
}

func main() {
    var v1 int32 = 42
    var v2 bool = (!false)
    var v3 uint64 = 2
    var v4 int = -841
    fmt.Println(v2)
    v9 := []int{v4, int(v1), -97}
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    var v10 string = ""
    if len(v9) > 0 {
        v9[0] = v4
    }
    var v11 int32 = (v1 % ((v1) | 1))
    v11 = (v1 | -2147483648)
    var v12 uint8 = 91
    var v13 int32 = v11
    var v14 bool = (0 < 2)
    fmt.Println(v4, v13, v3)
    if (-2.25 > 3.14159) {
        switch (v10 + v10) {
        case "foo bar":
            switch "hello" {
            case "\n\t":
                v3 -= (v3 | 882)
            case "foo bar":
                v15 := GenS1{F0: "hello", F1: v12, F2: false}
                _ = v15
            case "a":
                if ((v3 << ((782) & 63)) < (v3 ^ (^(v3)))) {
                    v16 := GenS1{F0: v10, F1: (v12 / ((v12) | 1)), F2: true}
                    v2 = false
                    v17 := GenS1{F0: "hello", F1: (v12 % (((v12 & genFn0())) | 1)), F2: v2}
                    _ = v16
                    _ = v17
                } else {
                    v13 -= (v1 + v13)
                }
            default:
                for fi18 := 0; fi18 < 15; fi18++ {
                    v19 := GenS1{F0: (v10 + "a"), F1: v12, F2: (false && false)}
                    switch v10 {
                    case "a":
                        v2 = (((("a" < v10) || v2) && (1 < 9223372036854775807)) && v14)
                    }
                    fmt.Println(v11, v1, fi18)
                    _ = fi18
                    _ = v19
                }
            }
        }
        for fi20 := 0; fi20 < 11; fi20++ {
            for fi21 := 0; fi21 < 7; fi21++ {
                var v22 uint16 = 39
                _ = fi21
                _ = v22
            }
            v2 = v14
            _ = fi20
        }
    } else {
        for fi23 := 0; fi23 < 6; fi23++ {
            v14 = (!false)
            v2 = (!(0 != -9223372036854775808))
            if (!true) {
                switch "a" {
                case "foo bar":
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                default:
                    for fi24 := 0; fi24 < 14; fi24++ {
                        switch v10 {
                        case "foo bar":
                            v1 -= v13
                        case "a":
                            var v25 int16 = 243
                            _ = v25
                        default:
                            switch (v4 + v4) {
                            case -2147483648:
                                switch v10 {
                                case "hello":
                                    v3 = (v3 / ((v3) | 1))
                                }
                            case 979:
                                v2 = v14
                            case -826:
                                var v26 float32 = -2.25
                                _ = v26
                            }
                        }
                        v13 = 0
                        _ = fi24
                    }
                }
                var v27 uint16 = 65535
                _ = v27
            } else {
                for fi28 := 0; fi28 < 1; fi28++ {
                    switch 723 {
                    case 327:
                        v10 = ((v10 + v10) + v10)
                    case 765:
                        _ = len(v9)
                    case 0:
                        v4 = 529
                    }
                    if (false && (fi28 == (^(fi28)))) {
                        var v29 bool = true
                        _ = v29
                    } else {
                        var v30 int32 = (v1 ^ (v11 + v11))
                        var v31 uint64 = 42
                        v32 := GenS1{F0: v10, F1: 127, F2: (1 < 88)}
                        _ = v30
                        _ = v31
                        _ = v32
                    }
                    _ = fi28
                }
            }
            _ = fi23
        }
    }
    var v33 string = ((v10 + v10) + "hello")
    for fi34 := 0; fi34 < 8; fi34++ {
        if (262 != 1) {
            var v35 uint64 = (^(v3))
            _ = v35
        }
        if ((v10 + v33) >= v10) {
            if ((v3 != v3) && (0.0 <= 0.1)) {
                if v14 {
                    v3 -= 0
                } else {
                    var v36 uint16 = 32767
                    _ = v36
                }
                for fi37 := 0; fi37 < 3; fi37++ {
                    _ = len(v9)
                    switch fi37 {
                    case 127:
                        fmt.Println(v33, v14)
                    }
                    _ = fi37
                }
            } else {
                var v38 int32 = (v13 >> ((4294967295) & 63))
                if (2 != 65535) {
                    var v39 float32 = 0.1
                    _ = v39
                }
                var v40 int = fi34
                _ = v38
                _ = v40
            }
            v41 := GenS2{F0: v11, F1: (!(false && (!(v33 < v10))))}
            for fi42 := 0; fi42 < 5; fi42++ {
                v13 += v13
                var v43 string = "a"
                switch (("\n\t" + "\n\t") + v33) {
                case "hello":
                    if (v1 > v1) {
                        if v14 {
                            fmt.Println(v3, v14)
                            var v44 int32 = v1
                            _ = v44
                        } else {
                            v10 = ("a" + v43)
                            if (32767 < 0) {
                                fmt.Println(v1, v33, v10)
                                var v45 string = (v10 + ((v33 + "\n\t") + ""))
                                if (!(genFn0() <= (v12 | v12))) {
                                    fmt.Println(v12)
                                    if len(v9) > 0 {
                                        v9[0] = fi42
                                    }
                                    v46 := GenS2{F0: v13, F1: (127 == -1)}
                                    v47 := GenS1{F0: v10, F1: (v12 << ((65535) & 63)), F2: v2}
                                    v11 -= (-(v11))
                                    if len(v9) > 0 {
                                        v9[0] = (fi34 << ((18446744073709551615) & 63))
                                    }
                                    _ = v46
                                    _ = v47
                                } else {
                                    v48 := GenS1{F0: ("" + "\n\t"), F1: v12, F2: false}
                                    v49 := GenS3{F0: 42, F1: v43}
                                    if (-2.25 != ((((1.5 / -2.25) - float32(-2.25)) / 1e10) + 0.0)) {
                                        switch fi34 {
                                        case 4294967295:
                                            fmt.Println(v43, v4)
                                        case 2147483647:
                                            if false {
                                                v50 := GenS4{F0: fi34, F1: -9223372036854775808}
                                                if (42 == 2147483648) {
                                                    v51 := GenS2{F0: v1, F1: false}
                                                    _ = v51
                                                }
                                                var v52 uint32 = 2147483647
                                                _ = v50
                                                _ = v52
                                            } else {
                                                switch v43 {
                                                case "\n\t":
                                                    var v53 uint8 = (v12 - (v12 + genFn0()))
                                                    _ = v53
                                                case "":
                                                    v2 = ((-366 >= -1) && false)
                                                case "foo bar":
                                                    fmt.Println(v3, v12, fi34)
                                                default:
                                                    var v54 int16 = -364
                                                    _ = v54
                                                }
                                                if true {
                                                    if (0 == 127) {
                                                        v55 := GenS5{F0: 691, F1: uint(v12), F2: 1}
                                                        var v56 float32 = 0.0
                                                        v57 := GenS3{F0: 143, F1: "\n\t"}
                                                        _ = v55
                                                        _ = v56
                                                        _ = v57
                                                    }
                                                } else {
                                                    if false {
                                                        if v2 {
                                                            switch fi34 {
                                                            case -884:
                                                                var v58 bool = false
                                                                _ = v58
                                                            case 127:
                                                                if v14 {
                                                                    if (v1 > 0) {
                                                                        fmt.Println(v10, v14)
                                                                        v59 := GenS6{F0: -128, F1: 431}
                                                                        var v60 float32 = 3.14159
                                                                        _ = v59
                                                                        _ = v60
                                                                    }
                                                                    v61 := GenS3{F0: -9223372036854775808, F1: "hello"}
                                                                    fmt.Println(v10)
                                                                    _ = v61
                                                                }
                                                            }
                                                            if v2 {
                                                                switch "" {
                                                                case "\n\t":
                                                                    v62 := GenS7{F0: 188, F1: (fi34 << ((42) & 63))}
                                                                    _ = v62
                                                                case "foo bar":
                                                                    v3 -= (^(v3))
                                                                case "hello":
                                                                    if (v14 && true) {
                                                                        if v14 {
                                                                            var v63 float64 = 1.5
                                                                            var v64 uint64 = v3
                                                                            if v2 {
                                                                                v2 = v14
                                                                                _ = len(v9)
                                                                                var v65 uint8 = (v12 % (((^(v12))) | 1))
                                                                                _ = v65
                                                                            } else {
                                                                                v13 -= (v13 >> ((255) & 63))
                                                                                var v66 int32 = -173
                                                                                _ = v66
                                                                            }
                                                                            _ = v63
                                                                            _ = v64
                                                                        } else {
                                                                            switch 32767 {
                                                                            case -2147483648:
                                                                                if len(v9) > 0 {
                                                                                    v9[0] = (fi34 - (fi42 & fi42))
                                                                                }
                                                                            case 259:
                                                                                v33 = v43
                                                                            default:
                                                                                if (!v14) {
                                                                                    var v67 float32 = (0.1 + ((((1.5 - float32((3.14159 / (1.5 + -2.25)))) - 3.14159) / 1.0) / 1.0))
                                                                                    _ = v67
                                                                                }
                                                                            }
                                                                        }
                                                                        fmt.Println(v2)
                                                                    } else {
                                                                        var v68 int64 = 977
                                                                        v14 = (255 == 32767)
                                                                        if v14 {
                                                                            v69 := GenS8{F0: v45}
                                                                            var v70 int16 = 1
                                                                            v71 := GenS9{F0: 90, F1: 0}
                                                                            _ = v69
                                                                            _ = v70
                                                                            _ = v71
                                                                        }
                                                                        _ = v68
                                                                    }
                                                                default:
                                                                    v33 = ""
                                                                }
                                                                v72 := GenS9{F0: -106, F1: -1}
                                                                _ = v72
                                                            }
                                                        } else {
                                                            var v73 int8 = -1
                                                            _ = v73
                                                        }
                                                        v12 = (v12 / ((v12) | 1))
                                                    } else {
                                                        v11 = v11
                                                        v74 := GenS6{F0: -128, F1: -269}
                                                        v14 = (457 != -545)
                                                        _ = v74
                                                    }
                                                    v75 := GenS10{F0: 32768, F1: (v11 % ((v11) | 1))}
                                                    _ = v75
                                                }
                                                switch fi42 {
                                                case 2147483647:
                                                    v14 = ((true && true) && (!v14))
                                                }
                                            }
                                        default:
                                            var v76 bool = true
                                            _ = v76
                                        }
                                    } else {
                                        v2 = v2
                                    }
                                    _ = v48
                                    _ = v49
                                }
                                _ = v45
                            } else {
                                fmt.Println(v3)
                                v1 -= v11
                                v77 := GenS7{F0: 2, F1: fi42}
                                _ = v77
                            }
                            var v78 int32 = (v13 % ((42) | 1))
                            fmt.Println(v2, v43)
                            _ = v78
                        }
                        if v14 {
                            if len(v9) > 0 {
                                v9[0] = v4
                            }
                        }
                        var v79 bool = (1e10 > 0.0)
                        fmt.Println(v43, fi34)
                        var v80 float32 = (3.14159 + (-2.25 + ((float32((1e10 - 0.0)) * ((0.0 * 1e10) / 1.0)) * 1.5)))
                        _ = v79
                        _ = v80
                    } else {
                        switch fi42 {
                        case 127:
                            var v81 uint16 = 884
                            _ = v81
                        case 217:
                            if len(v9) > 0 {
                                v9[0] = fi34
                            }
                        }
                    }
                default:
                    if (true && ((v12 ^ v12) == (v12 / ((v12) | 1)))) {
                        _ = len(v9)
                        if v14 {
                            fmt.Println(fi34, v14, v1)
                            if (fi34 != -1) {
                                var v82 int = (fi42 * (fi34 ^ fi34))
                                v83 := GenS2{F0: 663, F1: v2}
                                _ = v82
                                _ = v83
                            } else {
                                if len(v9) >= 2 {
                                    v9 = v9[1:len(v9)]
                                }
                            }
                            var v84 float64 = (0.0 - float64(-2.25))
                            if (-32768 != 248) {
                                v2 = ((!v2) && (65535 >= 0))
                                v85 := GenS7{F0: 2, F1: (fi34 >> ((4294967295) & 63))}
                                if len(v9) > 0 {
                                    v9[0] = fi42
                                }
                                v33 = "\n\t"
                                _ = v85
                            } else {
                                var v86 int32 = 0
                                v10 = v43
                                _ = v86
                            }
                            var v87 int32 = (v13 + v11)
                            _ = v84
                            _ = v87
                        } else {
                            switch (fi42 ^ v4) {
                            case -765:
                                var v88 bool = ((v3 <= 18446744073709551615) || v2)
                                _ = v88
                            case 42:
                                var v89 uint16 = 1
                                _ = v89
                            default:
                                if ((v14 || v2) || (-312 >= 983)) {
                                    v90 := GenS3{F0: 9223372036854775807, F1: ("hello" + ("a" + ""))}
                                    var v91 uint = 605
                                    v11 -= (v11 << ((uint(v12)) & 63))
                                    v14 = v14
                                    if v2 {
                                        switch v33 {
                                        case "hello":
                                            fmt.Println(v2, v91)
                                        case "a":
                                            v92 := GenS11{F0: 32767, F1: 2147483647, F2: int16(v4)}
                                            _ = v92
                                        case "foo bar":
                                            v93 := GenS6{F0: 42, F1: -9223372036854775808}
                                            _ = v93
                                        default:
                                            v9 = append(v9, v4)
                                            var v94 string = v33
                                            _ = v94
                                        }
                                        v12 = (^(v12))
                                        var v95 uint32 = 2147483647
                                        v96 := GenS10{F0: (v4 >> (((^(v91))) & 63)), F1: 561}
                                        v97 := GenS4{F0: (fi42 ^ (-(fi42))), F1: 0}
                                        _ = v95
                                        _ = v96
                                        _ = v97
                                    } else {
                                        var v98 uint64 = 2
                                        _ = v98
                                    }
                                    _ = v90
                                    _ = v91
                                } else {
                                    v4 = (fi34 + fi34)
                                    var v99 int64 = 9223372036854775807
                                    _ = v99
                                }
                                v100 := GenS1{F0: "\n\t", F1: (v12 | v12), F2: false}
                                _ = v100
                            }
                            v1 -= v13
                            var v101 float64 = ((0.1 * (0.0 + (-2.25 - 0.1))) - (1e10 + 1e10))
                            fmt.Println(v101, v101, v101)
                            _ = v101
                        }
                    } else {
                        var v102 int64 = 0
                        _ = v102
                    }
                }
                v10 = v43
                _ = fi42
                _ = v43
            }
            _ = v41
        }
        _ = fi34
    }
    if len(v9) > 0 {
        v9[0] = (v4 * v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v9
}
