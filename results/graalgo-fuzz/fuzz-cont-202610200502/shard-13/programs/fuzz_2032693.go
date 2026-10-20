package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 string; F1 string }
type GenS6 struct { F0 uint16; F1 uint64 }
type GenS7 struct { F0 string; F1 int; F2 uint16 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 bool; F1 uint32 }
type GenS10 struct { F0 int64; F1 int64; F2 string }
type GenS11 struct { F0 uint64; F1 int16 }
func genFn0(p0 int8, p1 uint8) uint8 {
    v6 := []int{}
    fmt.Println(p0, p1)
    p0 = 0
    _ = p0
    _ = p1
    _ = v6
    return p1
}

func main() {
    var v1 uint8 = 127
    var v2 string = (("" + "hello") + (("" + (("a" + "hello") + (("" + (("a" + "foo bar") + "")) + ""))) + ("\n\t" + ("a" + ("a" + "foo bar")))))
    var v3 bool = (691 < 221)
    var v4 int64 = -9223372036854775808
    v5 := make(map[uint32]string)
    fmt.Println(v1, v3, v3)
    for fi7 := 0; fi7 < 11; fi7++ {
        for fi8 := 0; fi8 < 5; fi8++ {
            var v9 int64 = (v4 / ((v4) | 1))
            _ = fi8
            _ = v9
        }
        if v3 {
            if (v3 && ((((fi7 & fi7) == fi7) && true) || false)) {
                switch v2 {
                case "hello":
                    v10 := []int{fi7, (-(fi7)), fi7, fi7}
                    _ = v10
                case "a":
                    v3 = (65535 != 4294967295)
                default:
                    switch "foo bar" {
                    case "hello":
                        if true {
                            v11 := []int{fi7}
                            switch fi7 {
                            case -9223372036854775808:
                                if v3 {
                                    v12 := GenS1{F0: -1}
                                    _ = v12
                                } else {
                                    var v13 uint8 = 2
                                    switch v2 {
                                    case "hello":
                                        var v14 bool = v3
                                        _ = v14
                                    case "":
                                        v15 := GenS1{F0: -332}
                                        _ = v15
                                    }
                                    _ = v13
                                }
                            case 127:
                                var v16 float32 = 0.1
                                _ = v16
                            }
                            _ = v11
                        } else {
                            fmt.Println(v3, v3, v3)
                            switch v2 {
                            case "a":
                                switch fi7 {
                                case -2147483648:
                                    for fi17 := 0; fi17 < 11; fi17++ {
                                        v2 = v2
                                        var v18 int32 = 2147483647
                                        _ = fi17
                                        _ = v18
                                    }
                                case 32768:
                                    v19 := []int{fi7, (fi7 ^ (fi7 + (fi7 - fi7)))}
                                    _ = v19
                                }
                            case "\n\t":
                                switch 32768 {
                                case 0:
                                    switch v2 {
                                    case "hello":
                                        for fi20 := 0; fi20 < 12; fi20++ {
                                            switch int(v1) {
                                            case 127:
                                                var v21 uint32 = 8
                                                switch v2 {
                                                case "\n\t":
                                                    v22 := []int{(fi20 % (((fi20 % (((fi7 - fi7)) | 1))) | 1)), fi7, (fi7 << ((42) & 63))}
                                                    _ = v22
                                                }
                                                if (-81 < -116) {
                                                    v23 := []int{fi7, (fi7 ^ (-(fi7)))}
                                                    var v24 uint8 = genFn0(0, v1)
                                                    _ = v23
                                                    _ = v24
                                                }
                                                _ = v21
                                            case 32768:
                                                var v25 uint64 = 4294967295
                                                _ = v25
                                            }
                                            var v26 uint64 = 328
                                            _ = fi20
                                            _ = v26
                                        }
                                    case "foo bar":
                                        switch fi7 {
                                        case 1:
                                            var v27 uint8 = genFn0(100, 0)
                                            _ = v27
                                        case 2147483647:
                                            fmt.Println(v4, v1)
                                        case 42:
                                            fmt.Println(v4, v1)
                                        default:
                                            var v28 int16 = -1
                                            _ = v28
                                        }
                                    default:
                                        if (32767 > fi7) {
                                            if v3 {
                                                for fi29 := 0; fi29 < 18; fi29++ {
                                                    v4 = (v4 << ((1) & 63))
                                                    var v30 string = "a"
                                                    v1 = (v1 - (v1 + genFn0(-101, (^(v1)))))
                                                    _ = fi29
                                                    _ = v30
                                                }
                                                var v31 float64 = ((((3.14159 / 1.0) + 0.0) - 0.1) / -2.25)
                                                _ = v31
                                            } else {
                                                v1 += (^(v1))
                                                var v32 int64 = 42
                                                v2 = ((((v2 + (v2 + "foo bar")) + v2) + v2) + (v2 + "\n\t"))
                                                _ = v32
                                            }
                                            if v3 {
                                                var v33 uint32 = 42
                                                if false {
                                                    var v34 int32 = 1
                                                    for fi35 := 0; fi35 < 1; fi35++ {
                                                        var v36 bool = v3
                                                        switch fi35 {
                                                        case 32767:
                                                            v2 = (v2 + v2)
                                                        case -9223372036854775808:
                                                            fmt.Println(v1, fi7)
                                                        case -2147483648:
                                                            var v37 float64 = -2.25
                                                            if false {
                                                                var v38 uint32 = (v33 % ((v33) | 1))
                                                                _ = v38
                                                            }
                                                            switch (("a" + "hello") + v2) {
                                                            case "":
                                                                fmt.Println(v3, v33, v36)
                                                                switch v2 {
                                                                case "":
                                                                    v3 = (v3 && ((42 == v33) || ((0.0 * 1e10) < 0.1)))
                                                                case "\n\t":
                                                                    switch 4294967295 {
                                                                    case -1:
                                                                        switch v2 {
                                                                        case "a":
                                                                            _ = len(v5)
                                                                        case "hello":
                                                                            var v39 uint16 = 1
                                                                            _ = v39
                                                                        case "foo bar":
                                                                            fmt.Println(v1)
                                                                        default:
                                                                            v33 = 2147483647
                                                                        }
                                                                    case -9223372036854775808:
                                                                        v5[(v33 >> ((42) & 63))] = "\n\t"
                                                                    case 4294967295:
                                                                        if (v33 <= 189) {
                                                                            switch v2 {
                                                                            case "foo bar":
                                                                                var v40 string = v2
                                                                                _ = v40
                                                                            case "":
                                                                                v41 := GenS2{F0: 2}
                                                                                _ = v41
                                                                            case "\n\t":
                                                                                switch "a" {
                                                                                case "hello":
                                                                                    if ((!(((v2 + ((v2 + v2) + "foo bar")) + v2) >= v2)) || (false && v3)) {
                                                                                        v42 := GenS2{F0: 4294967295}
                                                                                        v43 := []int{-451, fi35}
                                                                                        v33 = v33
                                                                                        _ = v42
                                                                                        _ = v43
                                                                                    } else {
                                                                                        var v44 float32 = 0.0
                                                                                        _ = v44
                                                                                    }
                                                                                case "\n\t":
                                                                                    var v45 float32 = 1.5
                                                                                    _ = v45
                                                                                case "foo bar":
                                                                                    var v46 string = v2
                                                                                    _ = v46
                                                                                }
                                                                            }
                                                                        } else {
                                                                            switch 829 {
                                                                            case 2147483647:
                                                                                var v47 bool = ((((v37 / (((v37 + 1e10) + v37) / 0.1)) / 0.1) != v37) && ((v34 * (v34 >> ((18446744073709551615) & 63))) > v34))
                                                                                _ = v47
                                                                            case 127:
                                                                                fmt.Println(fi35)
                                                                            case 0:
                                                                                var v48 int = fi7
                                                                                _ = v48
                                                                            default:
                                                                                if (42 == (v4 & (v4 << ((271) & 63)))) {
                                                                                    var v49 uint64 = uint64(v33)
                                                                                    _ = v49
                                                                                } else {
                                                                                    if v36 {
                                                                                        var v50 float32 = float32(float64(((3.14159 + (-2.25 * (0.0 / -2.25))) / 0.1)))
                                                                                        v4 -= v4
                                                                                        _ = v50
                                                                                    }
                                                                                    if false {
                                                                                        v51 := GenS3{F0: (v4 * 9223372036854775807), F1: (v1 / ((v1) | 1))}
                                                                                        _ = v51
                                                                                    } else {
                                                                                        var v52 int64 = (v4 ^ (v4 | v4))
                                                                                        v4 = (v4 << ((4294967295) & 63))
                                                                                        var v53 uint64 = 2
                                                                                        _ = v52
                                                                                        _ = v53
                                                                                    }
                                                                                    fmt.Println(v2)
                                                                                }
                                                                            }
                                                                            v36 = (0 <= 42)
                                                                        }
                                                                        v54 := GenS1{F0: v34}
                                                                        var v55 bool = v36
                                                                        _ = v54
                                                                        _ = v55
                                                                    default:
                                                                        if (v3 || v36) {
                                                                            if (!v36) {
                                                                                v4 -= v4
                                                                                v34 -= (v34 / ((-939) | 1))
                                                                                v56 := GenS3{F0: (v4 << ((uint(v1)) & 63)), F1: (^(v1))}
                                                                                _ = v56
                                                                            } else {
                                                                                if ((3.14159 / (v37 + 0.0)) < 1.5) {
                                                                                    if v36 {
                                                                                        switch v2 {
                                                                                        case "foo bar":
                                                                                            v1 = v1
                                                                                            var v57 int = fi35
                                                                                            _ = v57
                                                                                        }
                                                                                        v3 = true
                                                                                    } else {
                                                                                        v4 += -190
                                                                                        var v58 int32 = 2147483647
                                                                                        _ = v58
                                                                                    }
                                                                                    fmt.Println(v36, v1, fi7)
                                                                                    v4 += 1
                                                                                } else {
                                                                                    if false {
                                                                                        v59 := []int{}
                                                                                        if (-2147483648 != v34) {
                                                                                            if (false && v3) {
                                                                                                var v60 float64 = 3.14159
                                                                                                v61 := GenS3{F0: (v4 | v4), F1: v1}
                                                                                                _ = v60
                                                                                                _ = v61
                                                                                            } else {
                                                                                                if ((0 == 42) && v36) {
                                                                                                    switch (fi35 * 1) {
                                                                                                    case -995:
                                                                                                        if false {
                                                                                                            var v62 int16 = -1
                                                                                                            v59 = append(v59, 4294967295)
                                                                                                            if len(v59) > 0 {
                                                                                                                v59[0] = fi7
                                                                                                            }
                                                                                                            if v36 {
                                                                                                                v3 = ((fi7 & 42) > (fi35 + fi35))
                                                                                                                var v63 int16 = 75
                                                                                                                _ = v63
                                                                                                            }
                                                                                                            _ = v62
                                                                                                        } else {
                                                                                                            if len(v59) >= 2 {
                                                                                                                v59 = v59[1:len(v59)]
                                                                                                            }
                                                                                                            v33 = v33
                                                                                                        }
                                                                                                    case 32768:
                                                                                                        var v64 uint64 = 918
                                                                                                        _ = v64
                                                                                                    case 485:
                                                                                                        if ((1.5 + ((3.14159 / (3.14159 / (v37 * 0.1))) + (v37 + v37))) > v37) {
                                                                                                            v65 := v5[240]
                                                                                                            switch v65 {
                                                                                                            case "\n\t":
                                                                                                                fmt.Println(v33, v33, v36)
                                                                                                            }
                                                                                                            v2 = (v2 + v2)
                                                                                                            _ = v65
                                                                                                        }
                                                                                                    }
                                                                                                } else {
                                                                                                    var v66 int64 = v4
                                                                                                    if ((v34 <= (v34 / ((-2147483648) | 1))) && false) {
                                                                                                        if (!(8 != -28)) {
                                                                                                            v66 -= (v66 - (v4 - v66))
                                                                                                        }
                                                                                                        var v67 int64 = 0
                                                                                                        _ = v67
                                                                                                    } else {
                                                                                                        if v3 {
                                                                                                            var v68 int8 = 85
                                                                                                            v2 = ("" + v2)
                                                                                                            _ = v68
                                                                                                        } else {
                                                                                                            switch "a" {
                                                                                                            case "a":
                                                                                                                if false {
                                                                                                                    v59 = append(v59, fi35)
                                                                                                                    v69 := GenS1{F0: 289}
                                                                                                                    if (0 != (v4 >> ((2) & 63))) {
                                                                                                                        if (v4 > 1) {
                                                                                                                            v70 := GenS2{F0: 1}
                                                                                                                            var v71 int16 = 32767
                                                                                                                            _ = v70
                                                                                                                            _ = v71
                                                                                                                        } else {
                                                                                                                            var v72 int = fi35
                                                                                                                            if v3 {
                                                                                                                                v73 := GenS2{F0: 1}
                                                                                                                                _ = len(v59)
                                                                                                                                _ = v73
                                                                                                                            } else {
                                                                                                                                v34 = 1
                                                                                                                                var v74 float32 = float32(v37)
                                                                                                                                fmt.Println(v4, v4)
                                                                                                                                v66 += (v4 >> ((42) & 63))
                                                                                                                                _ = v74
                                                                                                                            }
                                                                                                                            _ = v72
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        var v75 int64 = int64(v34)
                                                                                                                        _ = v75
                                                                                                                    }
                                                                                                                    _ = v69
                                                                                                                } else {
                                                                                                                    if (9223372036854775807 == fi35) {
                                                                                                                        v66 -= v66
                                                                                                                    } else {
                                                                                                                        v4 = (v66 * v66)
                                                                                                                        if (!v36) {
                                                                                                                            var v76 uint64 = 0
                                                                                                                            _ = v76
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                            case "hello":
                                                                                                                var v77 int32 = (v34 | 1)
                                                                                                                _ = v77
                                                                                                            default:
                                                                                                                switch "\n\t" {
                                                                                                                case "\n\t":
                                                                                                                    if ((!(!(2 != 42))) && v3) {
                                                                                                                        if (0 >= 551) {
                                                                                                                            if (false && (!false)) {
                                                                                                                                v78 := GenS3{F0: (v66 & (v66 + (v4 * (v66 | 9223372036854775807)))), F1: v1}
                                                                                                                                _ = len(v59)
                                                                                                                                if (false || v3) {
                                                                                                                                    v2 = ((((v2 + ((("foo bar" + v2) + v2) + v2)) + "") + "a") + v2)
                                                                                                                                    v66 -= v66
                                                                                                                                    if false {
                                                                                                                                        if ((fi7 >> ((4294967295) & 63)) < (fi35 & (fi7 / ((fi7) | 1)))) {
                                                                                                                                            if false {
                                                                                                                                                if ((!v3) || ((0.1 != 1.5) && ((v4 / ((v66) | 1)) > v66))) {
                                                                                                                                                    v79 := GenS1{F0: -487}
                                                                                                                                                    var v80 int8 = 79
                                                                                                                                                    var v81 float32 = -2.25
                                                                                                                                                    _ = v79
                                                                                                                                                    _ = v80
                                                                                                                                                    _ = v81
                                                                                                                                                }
                                                                                                                                                var v82 int32 = (-(v34))
                                                                                                                                                _ = v82
                                                                                                                                            } else {
                                                                                                                                                _ = len(v59)
                                                                                                                                                if len(v59) >= 2 {
                                                                                                                                                    v59 = v59[1:len(v59)]
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                            fmt.Println(v66)
                                                                                                                                            delete(v5, v33)
                                                                                                                                        }
                                                                                                                                        var v83 int = fi7
                                                                                                                                        _ = v83
                                                                                                                                    }
                                                                                                                                } else {
                                                                                                                                    v36 = ("" < v2)
                                                                                                                                    v34 = -1
                                                                                                                                    var v84 uint32 = 0
                                                                                                                                    var v85 uint8 = genFn0(1, (v1 - 106))
                                                                                                                                    _ = v84
                                                                                                                                    _ = v85
                                                                                                                                }
                                                                                                                                _ = v78
                                                                                                                            }
                                                                                                                        }
                                                                                                                        v1 -= (^(v1))
                                                                                                                        v1 += v1
                                                                                                                    }
                                                                                                                case "foo bar":
                                                                                                                    _ = len(v5)
                                                                                                                case "a":
                                                                                                                    fmt.Println(v37, v2)
                                                                                                                }
                                                                                                            }
                                                                                                            if (42 < 509) {
                                                                                                                v3 = (0.0 > ((v37 / 1.5) + 1e10))
                                                                                                                if len(v59) > 0 {
                                                                                                                    v59[0] = fi7
                                                                                                                }
                                                                                                                var v86 uint16 = 1
                                                                                                                _ = v86
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    switch ("hello" + ("hello" + "\n\t")) {
                                                                                                    case "":
                                                                                                        v87 := GenS4{F0: 2147483648, F1: 940}
                                                                                                        _ = v87
                                                                                                    case "hello":
                                                                                                        switch v2 {
                                                                                                        case "":
                                                                                                            if (!v3) {
                                                                                                                var v88 int64 = (v66 & v4)
                                                                                                                _ = v88
                                                                                                            } else {
                                                                                                                var v89 int8 = 100
                                                                                                                v4 += v66
                                                                                                                fmt.Println(v89)
                                                                                                                v90 := GenS2{F0: 765}
                                                                                                                _ = v89
                                                                                                                _ = v90
                                                                                                            }
                                                                                                        case "foo bar":
                                                                                                            var v91 int64 = v4
                                                                                                            _ = v91
                                                                                                        default:
                                                                                                            var v92 float64 = 1.5
                                                                                                            _ = v92
                                                                                                        }
                                                                                                    default:
                                                                                                        var v93 int16 = -32768
                                                                                                        _ = v93
                                                                                                    }
                                                                                                    _ = v66
                                                                                                }
                                                                                            }
                                                                                            var v94 int8 = -1
                                                                                            fmt.Println(v36, v94)
                                                                                            _ = v94
                                                                                        } else {
                                                                                            _ = len(v59)
                                                                                            v5[v33] = "a"
                                                                                        }
                                                                                        if v36 {
                                                                                            switch fi35 {
                                                                                            case 4294967295:
                                                                                                v95 := GenS4{F0: v33, F1: 753}
                                                                                                _ = v95
                                                                                            case 42:
                                                                                                v5[v33] = "foo bar"
                                                                                            }
                                                                                            if v3 {
                                                                                                switch 9223372036854775807 {
                                                                                                case -1:
                                                                                                    var v96 int64 = 0
                                                                                                    _ = v96
                                                                                                default:
                                                                                                    if (1 >= 992) {
                                                                                                        v34 -= int32(v4)
                                                                                                        var v97 uint = 205
                                                                                                        v98 := GenS3{F0: -1, F1: (v1 * (v1 | v1))}
                                                                                                        _ = v97
                                                                                                        _ = v98
                                                                                                    } else {
                                                                                                        v99 := GenS5{F0: "foo bar", F1: v2}
                                                                                                        var v100 int16 = 32767
                                                                                                        if len(v59) >= 2 {
                                                                                                            v59 = v59[1:len(v59)]
                                                                                                        }
                                                                                                        _ = v100
                                                                                                        _ = v99
                                                                                                    }
                                                                                                }
                                                                                            } else {
                                                                                                switch fi7 {
                                                                                                case 2147483647:
                                                                                                    v37 = 1e10
                                                                                                case 42:
                                                                                                    switch 32768 {
                                                                                                    case 127:
                                                                                                        fmt.Println(v36)
                                                                                                    case -250:
                                                                                                        v101 := GenS1{F0: v34}
                                                                                                        _ = v101
                                                                                                    default:
                                                                                                        v102 := GenS4{F0: v33, F1: 0}
                                                                                                        _ = v102
                                                                                                    }
                                                                                                default:
                                                                                                    if v36 {
                                                                                                        v36 = (0 == 18446744073709551615)
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            var v103 float32 = 0.1
                                                                                            v34 += 2147483647
                                                                                            v4 += -336
                                                                                            _ = v103
                                                                                        }
                                                                                        _ = v59
                                                                                    } else {
                                                                                        delete(v5, (v33 >> ((0) & 63)))
                                                                                        delete(v5, (v33 - (v33 | 0)))
                                                                                        v3 = v36
                                                                                    }
                                                                                }
                                                                                v33 -= v33
                                                                                v104 := GenS6{F0: 0, F1: uint64(v1)}
                                                                                _ = v104
                                                                            }
                                                                            v36 = v36
                                                                            v33 = (v33 >> ((437) & 63))
                                                                        } else {
                                                                            if false {
                                                                                switch ((v2 + "hello") + "foo bar") {
                                                                                case "":
                                                                                    v105 := GenS7{F0: "\n\t", F1: (^(fi35)), F2: 886}
                                                                                    _ = v105
                                                                                case "foo bar":
                                                                                    if (!false) {
                                                                                        if (!false) {
                                                                                            v5[(v33 + v33)] = ((v2 + v2) + "")
                                                                                            var v106 uint = 255
                                                                                            _ = v106
                                                                                        }
                                                                                    }
                                                                                    var v107 bool = v3
                                                                                    _ = v107
                                                                                case "\n\t":
                                                                                    v108 := GenS2{F0: 4294967295}
                                                                                    var v109 uint = 457
                                                                                    v110 := GenS8{F0: (v109 >> (((v109 - v109)) & 63))}
                                                                                    v3 = (((!v3) || (uint(v33) <= (v109 / (((^(v109))) | 1)))) || v3)
                                                                                    var v111 int32 = v34
                                                                                    _ = v108
                                                                                    _ = v109
                                                                                    _ = v110
                                                                                    _ = v111
                                                                                }
                                                                                var v112 float32 = (((1.5 / float32(((((v37 * v37) + v37) / -2.25) - v37))) - 0.1) - -2.25)
                                                                                v1 = v1
                                                                                _ = v112
                                                                            }
                                                                            var v113 string = (v2 + v2)
                                                                            var v114 int64 = (-(v4))
                                                                            _ = v113
                                                                            _ = v114
                                                                        }
                                                                    }
                                                                case "hello":
                                                                    v3 = true
                                                                default:
                                                                    v33 = (v33 / (((^(v33))) | 1))
                                                                }
                                                            case "\n\t":
                                                                v115 := GenS9{F0: v3, F1: 65535}
                                                                _ = v115
                                                            }
                                                            _ = v37
                                                        default:
                                                            switch int(v1) {
                                                            case 4294967295:
                                                                if (!((fi35 == (fi35 + fi35)) && v3)) {
                                                                    v3 = v36
                                                                } else {
                                                                    v116 := GenS10{F0: v4, F1: (v4 >> ((660) & 63)), F2: "foo bar"}
                                                                    _ = v116
                                                                }
                                                            case 32767:
                                                                var v117 float64 = 1e10
                                                                _ = v117
                                                            default:
                                                                var v118 float64 = (float64(((-2.25 / (3.14159 - (((((-2.25 - 1e10) - 3.14159) - 3.14159) / (float32(0.0) + -2.25)) - ((3.14159 / 1.0) / 1.0)))) + (1e10 + 0.0))) / -2.25)
                                                                _ = v118
                                                            }
                                                        }
                                                        var v119 bool = (0 == 100)
                                                        _ = fi35
                                                        _ = v119
                                                        _ = v36
                                                    }
                                                    _ = v34
                                                }
                                                _ = v33
                                            } else {
                                                var v120 uint8 = v1
                                                if false {
                                                    var v121 uint32 = 4294967295
                                                    _ = v121
                                                } else {
                                                    var v122 uint = 876
                                                    _ = v122
                                                }
                                                v120 -= v1
                                                _ = v120
                                            }
                                            delete(v5, 2147483648)
                                        } else {
                                            v2 = v2
                                            for fi123 := 0; fi123 < 14; fi123++ {
                                                var v124 uint16 = 64
                                                v3 = v3
                                                _ = fi123
                                                _ = v124
                                            }
                                        }
                                    }
                                default:
                                    v2 = ""
                                }
                            default:
                                for fi125 := 0; fi125 < 3; fi125++ {
                                    v3 = v3
                                    _ = fi125
                                }
                            }
                            v3 = (((float32(0.0) / (0.1 * 1.5)) * float32(((0.0 / 0.1) + 0.1))) == (1.5 - 3.14159))
                        }
                    case "":
                        if ((v4 > int64(v1)) || false) {
                            v126 := GenS10{F0: 0, F1: v4, F2: "\n\t"}
                            fmt.Println(v4)
                            _ = v126
                        }
                    case "\n\t":
                        if v3 {
                            switch v2 {
                            case "hello":
                                var v127 float64 = ((1e10 - float64((0.1 / (1.5 * 1e10)))) - float64(1.5))
                                _ = v127
                            case "a":
                                v4 = v4
                            default:
                                switch v2 {
                                case "hello":
                                    fmt.Println(v2)
                                case "a":
                                    switch "" {
                                    case "":
                                        if v3 {
                                            v1 = (v1 + 2)
                                            if ((true && false) || ((1.5 == (float32(1.5) / 1e10)) && true)) {
                                                for fi128 := 0; fi128 < 1; fi128++ {
                                                    var v129 int8 = -128
                                                    fmt.Println(v4, fi7)
                                                    v4 -= (v4 / (((^(v4))) | 1))
                                                    _ = fi128
                                                    _ = v129
                                                }
                                                if true {
                                                    var v130 float64 = ((3.14159 / (3.14159 / 0.1)) + (0.0 - (1e10 * 3.14159)))
                                                    _ = v130
                                                }
                                                var v131 uint32 = 100
                                                _ = v131
                                            } else {
                                                var v132 float64 = 1e10
                                                v1 += v1
                                                v5[65535] = "a"
                                                _ = v132
                                            }
                                            for fi133 := 0; fi133 < 18; fi133++ {
                                                v134 := GenS3{F0: (-(v4)), F1: genFn0(-128, v1)}
                                                if v3 {
                                                    fmt.Println(fi7, fi7)
                                                } else {
                                                    v135 := GenS10{F0: (v4 + (v4 ^ (-(v4)))), F1: v4, F2: (v2 + "hello")}
                                                    switch 986 {
                                                    case -1:
                                                        v136 := GenS8{F0: 4294967295}
                                                        _ = v136
                                                    case 2147483647:
                                                        if v3 {
                                                            v137 := []int{int(v1), (fi133 % (((fi7 ^ (fi7 * -2147483648))) | 1))}
                                                            switch ("a" + "") {
                                                            case "foo bar":
                                                                fmt.Println(v2, fi133)
                                                            case "hello":
                                                                if v3 {
                                                                    var v138 uint8 = v1
                                                                    v2 = v2
                                                                    v139 := GenS10{F0: v4, F1: (v4 * v4), F2: ("\n\t" + ("\n\t" + "\n\t"))}
                                                                    _ = v138
                                                                    _ = v139
                                                                } else {
                                                                    switch ("\n\t" + (v2 + "\n\t")) {
                                                                    case "a":
                                                                        switch v2 {
                                                                        case "foo bar":
                                                                            switch v2 {
                                                                            case "foo bar":
                                                                                if v3 {
                                                                                    fmt.Println(fi7, v2)
                                                                                } else {
                                                                                    if true {
                                                                                        fmt.Println(v1)
                                                                                    } else {
                                                                                        if (!v3) {
                                                                                            fmt.Println(v3, v4)
                                                                                        }
                                                                                        delete(v5, 306)
                                                                                    }
                                                                                    switch (fi133 | -1) {
                                                                                    case -585:
                                                                                        fmt.Println(v4, v1)
                                                                                    default:
                                                                                        v140 := GenS9{F0: ("hello" == ""), F1: 906}
                                                                                        _ = v140
                                                                                    }
                                                                                }
                                                                            }
                                                                        default:
                                                                            var v141 int32 = 1
                                                                            _ = v141
                                                                        }
                                                                    case "hello":
                                                                        v142 := v5[65535]
                                                                        _ = v142
                                                                    case "":
                                                                        v2 = v2
                                                                    }
                                                                }
                                                            case "a":
                                                                switch fi133 {
                                                                case 2147483647:
                                                                    var v143 int = fi133
                                                                    _ = v143
                                                                }
                                                            default:
                                                                var v144 int32 = 1
                                                                _ = v144
                                                            }
                                                            v3 = v3
                                                            v145 := v5[968]
                                                            var v146 int16 = -948
                                                            _ = v137
                                                            _ = v145
                                                            _ = v146
                                                        } else {
                                                            v2 = "\n\t"
                                                        }
                                                    }
                                                    v2 = "\n\t"
                                                    _ = v135
                                                }
                                                fmt.Println(v2, v3, v2)
                                                _ = fi133
                                                _ = v134
                                            }
                                        }
                                    case "\n\t":
                                        for fi147 := 0; fi147 < 14; fi147++ {
                                            v4 = (v4 | v4)
                                            v148 := GenS10{F0: (v4 ^ (v4 - (v4 - -965))), F1: (^(v4)), F2: "foo bar"}
                                            _ = fi147
                                            _ = v148
                                        }
                                    case "a":
                                        v1 -= v1
                                    default:
                                        v149 := GenS10{F0: (v4 / (((v4 ^ 0)) | 1)), F1: v4, F2: v2}
                                        _ = v149
                                    }
                                case "\n\t":
                                    var v150 uint32 = 2147483647
                                    _ = v150
                                }
                            }
                            for fi151 := 0; fi151 < 3; fi151++ {
                                v1 = (v1 >> ((42) & 63))
                                _ = fi151
                            }
                        } else {
                            v152 := v5[65535]
                            _ = v152
                        }
                    }
                }
            }
        }
        _ = fi7
    }
    for fi153 := 0; fi153 < 4; fi153++ {
        v154 := GenS3{F0: (v4 - (v4 >> ((518) & 63))), F1: v1}
        v2 = (v2 + "a")
        _ = fi153
        _ = v154
    }
    v155 := GenS11{F0: 514, F1: 483}
    v1 += 70
    for fi156 := 0; fi156 < 14; fi156++ {
        for fi157 := 0; fi157 < 10; fi157++ {
            v158 := []int{702, (fi156 >> ((255) & 63)), fi157}
            delete(v5, 0)
            _ = fi157
            _ = v158
        }
        v5[201] = v2
        var v159 uint32 = 2
        _ = fi156
        _ = v159
    }
    fmt.Println(v4)
    fmt.Println(v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v155
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
