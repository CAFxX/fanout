package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int8 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 string; F1 int; F2 uint8 }
type GenS7 struct { F0 int64; F1 uint16; F2 int }
type GenS8 struct { F0 uint64; F1 int }
type GenS9 struct { F0 uint16; F1 int }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 int64; F1 uint32 }
type GenS12 struct { F0 string; F1 string }
type GenS13 struct { F0 bool; F1 uint32 }
func genFn0(p0 int32, p1 uint32, p2 int64) uint {
    var v7 uint16 = 255
    for fi8 := 0; fi8 < 12; fi8++ {
        v9 := GenS1{F0: 1, F1: 2, F2: 42}
        _ = fi8
        _ = v9
    }
    v10 := []int{-412, -693, 42, -2147483648}
    for fi11 := 0; fi11 < 16; fi11++ {
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        for fi12 := 0; fi12 < 10; fi12++ {
            var v13 int = fi11
            fmt.Println(p0)
            p2 = (^(p2))
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    v7 += v7
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    return 595
}
func genFn1() uint16 {
    if (0.0 <= (((((((0.0 - (0.1 * 0.0)) + (3.14159 * (0.0 * 3.14159))) * 1.5) / 3.14159) * (3.14159 - (1e10 + (0.0 + (3.14159 * 0.1))))) / (1e10 + 3.14159)) * (1.5 * 0.0))) {
        for fi14 := 0; fi14 < 14; fi14++ {
            var v15 int16 = -433
            v16 := GenS2{F0: (v15 >> ((0) & 63)), F1: 380}
            _ = fi14
            _ = v15
            _ = v16
        }
        v17 := GenS3{F0: 666}
        switch ("a" + "\n\t") {
        case "\n\t":
            if true {
                var v18 int64 = -1
                v18 -= v18
                switch 32768 {
                case 32768:
                    fmt.Println(v18)
                case -1:
                    fmt.Println(v18)
                case 42:
                    fmt.Println(v18)
                }
                _ = v18
            }
        case "a":
            v19 := []int{}
            _ = v19
        case "hello":
            if (!false) {
                var v20 uint32 = 2147483647
                _ = v20
            } else {
                if (!(!false)) {
                    var v21 uint8 = 42
                    _ = v21
                } else {
                    if (-672 < -1) {
                        var v22 uint32 = 2147483647
                        _ = v22
                    } else {
                        var v23 int8 = -1
                        for fi24 := 0; fi24 < 10; fi24++ {
                            switch fi24 {
                            case -9223372036854775808:
                                switch (fi24 & fi24) {
                                case -2147483648:
                                    v25 := GenS4{F0: 32767}
                                    _ = v25
                                case 2147483647:
                                    v23 += v23
                                }
                            }
                            var v26 int64 = -1
                            v23 = 0
                            _ = fi24
                            _ = v26
                        }
                        fmt.Println(v23)
                        _ = v23
                    }
                    if (false && false) {
                        if ((404 == 367) && false) {
                            if (18446744073709551615 != 18446744073709551615) {
                                v27 := GenS2{F0: 1, F1: 4294967295}
                                _ = v27
                            } else {
                                v28 := GenS5{F0: -539}
                                var v29 int64 = 0
                                _ = v28
                                _ = v29
                            }
                            if false {
                                var v30 bool = false
                                for fi31 := 0; fi31 < 19; fi31++ {
                                    fmt.Println(fi31)
                                    v30 = ("\n\t" >= ("hello" + ((("a" + "a") + "hello") + "foo bar")))
                                    _ = fi31
                                }
                                _ = v30
                            }
                            v32 := GenS6{F0: "", F1: 42, F2: 89}
                            _ = v32
                        }
                        var v33 float32 = 3.14159
                        _ = v33
                    }
                }
                v34 := GenS7{F0: 0, F1: 32767, F2: 0}
                for fi35 := 0; fi35 < 19; fi35++ {
                    v36 := []int{}
                    var v37 uint8 = 42
                    v37 = (v37 ^ v37)
                    _ = fi35
                    _ = v36
                    _ = v37
                }
                _ = v34
            }
        }
        _ = v17
    }
    if false {
        switch 110 {
        case 338:
            for fi38 := 0; fi38 < 7; fi38++ {
                switch (fi38 | (-(fi38))) {
                case -906:
                    var v39 int = (fi38 - -297)
                    _ = v39
                case 0:
                    if (-9223372036854775808 == -1) {
                        if true {
                            for fi40 := 0; fi40 < 15; fi40++ {
                                var v41 int16 = -32768
                                v41 = v41
                                if true {
                                    fmt.Println(v41, fi38)
                                    var v42 uint64 = 42
                                    v41 = v41
                                    _ = v42
                                }
                                _ = fi40
                                _ = v41
                            }
                            var v43 uint32 = 764
                            _ = v43
                        }
                    } else {
                        if true {
                            fmt.Println(fi38)
                            fmt.Println(fi38)
                        }
                        var v44 float32 = 3.14159
                        fmt.Println(fi38)
                        _ = v44
                    }
                case -422:
                    fmt.Println(fi38)
                }
                var v45 bool = false
                v46 := GenS5{F0: -1}
                _ = fi38
                _ = v45
                _ = v46
            }
        }
        v47 := GenS7{F0: 1, F1: 42, F2: 42}
        _ = v47
    }
    var v48 string = ""
    v48 = (v48 + (v48 + ((((v48 + "foo bar") + v48) + v48) + (v48 + v48))))
    _ = v48
    return 76
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int8 = -128
    v3 := []string{("" + ""), "\n\t", "foo bar"}
    var v4 bool = (!true)
    v5 := []uint32{984, 1}
    var v6 uint16 = 583
    var v49 uint64 = 541
    v50 := GenS3{F0: 1}
    v1["foo bar"] = 4294967295
    v51 := GenS4{F0: 1000}
    v52 := GenS6{F0: ("foo bar" + "\n\t"), F1: -717, F2: 1}
    v4 = v4
    v6 = v6
    v4 = ((255 >= 1) || v4)
    var v53 uint32 = 2
    v54 := GenS8{F0: 1, F1: 32767}
    v55 := GenS9{F0: 2, F1: 637}
    v56 := GenS10{F0: (v49 * (v49 + (v49 - (v49 & 42))))}
    switch -677 {
    case -1:
        var v57 uint16 = (v6 - (v6 << ((1) & 63)))
        _ = v57
    case -2147483648:
        if (-9223372036854775808 <= 42) {
            if (v4 && v4) {
                var v58 int8 = (-(v2))
                _ = v58
            }
        } else {
            switch "a" {
            case "foo bar":
                var v59 uint16 = genFn1()
                _ = v59
            case "a":
                v53 += v53
            case "hello":
                v4 = (224 <= uint8(v53))
            default:
                var v60 uint8 = 248
                _ = v60
            }
            if v4 {
                switch "a" {
                case "foo bar":
                    if v4 {
                        var v61 int16 = 1
                        if len(v5) > 0 {
                            v5[0] = v53
                        }
                        _ = v61
                    } else {
                        var v62 uint64 = (v49 + (v49 % ((376) | 1)))
                        fmt.Println(v49, v2)
                        var v63 int32 = -1
                        _ = v62
                        _ = v63
                    }
                case "":
                    switch "hello" {
                    case "hello":
                        switch "" {
                        case "a":
                            v64 := v1["foo bar"]
                            _ = v64
                        case "\n\t":
                            switch 128 {
                            case 0:
                                var v65 int64 = 42
                                _ = v65
                            case 396:
                                v53 += (v53 / ((v53) | 1))
                            case 32767:
                                fmt.Println(v49, v2, v49)
                            default:
                                v66 := GenS11{F0: -9223372036854775808, F1: v53}
                                _ = v66
                            }
                        case "":
                            _ = len(v3)
                        default:
                            if (true && ((v49 | (v49 | v49)) >= (v49 | (^(v49))))) {
                                switch "\n\t" {
                                case "":
                                    v67 := GenS6{F0: "\n\t", F1: -255, F2: 2}
                                    _ = v67
                                default:
                                    v53 = (v53 + v53)
                                }
                            } else {
                                var v68 uint = 4294967295
                                if v4 {
                                    if len(v5) > 0 {
                                        v5[0] = 2
                                    }
                                } else {
                                    v69 := GenS8{F0: v49, F1: 32768}
                                    fmt.Println(v4, v68)
                                    if true {
                                        if (1000 < 32767) {
                                            v70 := GenS11{F0: 42, F1: (v53 * (v53 >> (((v68 | 2)) & 63)))}
                                            _ = v70
                                        } else {
                                            v5 = append(v5, 1)
                                            if ("hello" <= "\n\t") {
                                                v71 := GenS10{F0: 13}
                                                _ = v71
                                            }
                                        }
                                        if ((v6 << (((v68 | v68)) & 63)) <= (v6 & (^(v6)))) {
                                            v6 += (v6 - v6)
                                            fmt.Println(v49, v4)
                                            var v72 float64 = ((((3.14159 - float64((-2.25 * 0.1))) / float64((1e10 + 0.0))) - -2.25) + 0.0)
                                            _ = v72
                                        }
                                    }
                                    _ = v69
                                }
                                for fi73 := 0; fi73 < 3; fi73++ {
                                    for fi74 := 0; fi74 < 11; fi74++ {
                                        v68 += 751
                                        var v75 int = (^(fi73))
                                        var v76 int64 = 1
                                        _ = fi74
                                        _ = v75
                                        _ = v76
                                    }
                                    _ = fi73
                                }
                                _ = v68
                            }
                        }
                    case "foo bar":
                        v77 := GenS1{F0: 760, F1: 65535, F2: v2}
                        _ = v77
                    case "a":
                        if v4 {
                            v78 := GenS12{F0: ("" + "hello"), F1: "hello"}
                            for fi79 := 0; fi79 < 6; fi79++ {
                                var v80 bool = (!false)
                                v6 += v6
                                _ = fi79
                                _ = v80
                            }
                            _ = v78
                        }
                    }
                case "a":
                    if (18446744073709551615 != 2) {
                        if len(v5) > 0 {
                            v5[0] = (v53 >> ((uint(v53)) & 63))
                        }
                        v53 += v53
                    } else {
                        v6 -= (v6 + (v6 >> ((4294967295) & 63)))
                        v1["hello"] = (v53 % ((v53) | 1))
                    }
                default:
                    for fi81 := 0; fi81 < 4; fi81++ {
                        var v82 int64 = -204
                        for fi83 := 0; fi83 < 6; fi83++ {
                            v84 := GenS6{F0: "foo bar", F1: fi81, F2: 0}
                            if true {
                                v85 := GenS13{F0: v4, F1: uint32(v49)}
                                if (false || v4) {
                                    v53 -= (v53 / (((^(v53))) | 1))
                                    if v4 {
                                        switch fi81 {
                                        case -2147483648:
                                            var v86 uint = 4294967295
                                            _ = v86
                                        }
                                        var v87 uint32 = v53
                                        _ = v87
                                    }
                                    var v88 uint16 = (v6 | v6)
                                    _ = v88
                                }
                                _ = v85
                            }
                            var v89 uint16 = genFn1()
                            switch "a" {
                            case "hello":
                                v5 = append(v5, 2147483647)
                            case "a":
                                v90 := GenS9{F0: (v89 % ((v89) | 1)), F1: (fi81 + int(v2))}
                                _ = v90
                            }
                            _ = fi83
                            _ = v84
                            _ = v89
                        }
                        _ = fi81
                        _ = v82
                    }
                }
                var v91 string = "foo bar"
                fmt.Println(v53)
                _ = v91
            }
            var v92 int = 32767
            _ = v92
        }
    default:
        var v93 int64 = 392
        _ = v93
    }
    v53 += v53
    var v94 float64 = 1.5
    if false {
        _ = len(v3)
    } else {
        var v95 int8 = v2
        v53 = (v53 ^ v53)
        for fi96 := 0; fi96 < 9; fi96++ {
            v6 = (^(v6))
            v97 := GenS13{F0: ((v2 * v2) < 64), F1: 42}
            _ = fi96
            _ = v97
        }
        _ = v95
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v55
    _ = v56
    _ = v6
    _ = v94
}
