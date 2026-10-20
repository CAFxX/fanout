package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64; F1 bool; F2 int16 }
type GenS5 struct { F0 string; F1 uint64; F2 uint8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint64; F1 uint64; F2 uint }
func genFn0() bool {
    var v4 uint8 = 164
    v4 = 1
    v5 := GenS1{F0: (!(-1 < 100)), F1: 9223372036854775807}
    _ = v4
    _ = v5
    return (!((1e10 * (1.5 + -2.25)) <= ((0.0 / (1e10 * 1e10)) / -2.25)))
}
func genFn1() uint {
    if false {
        var v6 int64 = -9223372036854775808
        _ = v6
    }
    var v7 int16 = -180
    if ((false || ((!(-79 < -26)) || (42 != -43))) || false) {
        var v8 float64 = 1.5
        _ = v8
    }
    switch "hello" {
    case "a":
        if (4294967295 > 4294967295) {
            if (!("a" > "foo bar")) {
                v9 := []int{9223372036854775807, 880, -2147483648}
                fmt.Println(v7)
                var v10 bool = (-1 != -624)
                _ = v10
                _ = v9
            } else {
                fmt.Println(v7)
                v7 += 1000
            }
            v11 := GenS1{F0: false, F1: -1}
            _ = v11
        } else {
            var v12 int32 = 42
            var v13 int16 = v7
            switch -9223372036854775808 {
            case -2147483648:
                v13 += v13
            }
            _ = v12
            _ = v13
        }
    case "":
        switch "\n\t" {
        case "hello":
            var v14 int32 = -117
            _ = v14
        default:
            if true {
                var v15 uint8 = 0
                if (!(((v15 - (v15 * (v15 * (v15 >> ((0) & 63))))) < (v15 & 128)) || true)) {
                    switch "\n\t" {
                    case "foo bar":
                        for fi16 := 0; fi16 < 13; fi16++ {
                            switch (^(fi16)) {
                            case -9223372036854775808:
                                switch fi16 {
                                case -1:
                                    var v17 int16 = (v7 - (v7 | 1000))
                                    _ = v17
                                case 485:
                                    for fi18 := 0; fi18 < 16; fi18++ {
                                        v19 := GenS2{F0: (fi16 * -465)}
                                        _ = fi18
                                        _ = v19
                                    }
                                }
                            case 4294967295:
                                for fi20 := 0; fi20 < 10; fi20++ {
                                    var v21 string = "hello"
                                    _ = fi20
                                    _ = v21
                                }
                            case 221:
                                v22 := GenS3{F0: v7}
                                _ = v22
                            }
                            _ = fi16
                        }
                    }
                    if (!true) {
                        for fi23 := 0; fi23 < 14; fi23++ {
                            v24 := []int{fi23, 32767, (fi23 / ((fi23) | 1))}
                            if ((false && false) || (v15 == (v15 + (v15 << ((0) & 63))))) {
                                v25 := GenS2{F0: fi23}
                                v26 := GenS4{F0: 2, F1: true, F2: v7}
                                if ((0.1 + 1.5) < 0.0) {
                                    var v27 uint8 = v15
                                    if false {
                                        v28 := GenS5{F0: ("foo bar" + ("a" + "hello")), F1: 0, F2: (v27 - v15)}
                                        _ = v28
                                    }
                                    _ = v27
                                } else {
                                    if false {
                                        for fi29 := 0; fi29 < 1; fi29++ {
                                            v30 := GenS2{F0: (^(fi29))}
                                            _ = fi29
                                            _ = v30
                                        }
                                        switch ("foo bar" + "") {
                                        case "hello":
                                            var v31 string = ("a" + "foo bar")
                                            _ = v31
                                        case "a":
                                            v24 = append(v24, fi23)
                                        }
                                    }
                                    for fi32 := 0; fi32 < 10; fi32++ {
                                        v33 := GenS5{F0: "foo bar", F1: 0, F2: v15}
                                        _ = fi32
                                        _ = v33
                                    }
                                    if ((!false) && false) {
                                        for fi34 := 0; fi34 < 4; fi34++ {
                                            var v35 bool = false
                                            var v36 int32 = 42
                                            v37 := GenS6{F0: 0}
                                            _ = fi34
                                            _ = v35
                                            _ = v36
                                            _ = v37
                                        }
                                        fmt.Println(fi23, v7, v7)
                                    } else {
                                        var v38 int16 = (v7 % (((^(v7))) | 1))
                                        _ = v38
                                    }
                                }
                                _ = v25
                                _ = v26
                            } else {
                                if (!true) {
                                    v15 = v15
                                    _ = len(v24)
                                } else {
                                    var v39 uint64 = 122
                                    _ = v39
                                }
                            }
                            var v40 int = fi23
                            _ = fi23
                            _ = v24
                            _ = v40
                        }
                        v7 = -837
                    }
                }
                _ = v15
            }
        }
    case "foo bar":
        fmt.Println(v7)
    default:
        fmt.Println(v7)
    }
    for fi41 := 0; fi41 < 15; fi41++ {
        v42 := GenS7{F0: 65535}
        if false {
            v7 = (-(v7))
        }
        var v43 int = (^(fi41))
        _ = fi41
        _ = v42
        _ = v43
    }
    _ = v7
    return 4294967295
}

func main() {
    v1 := make(map[int]int)
    var v2 string = (("hello" + ("\n\t" + "\n\t")) + "foo bar")
    var v3 int32 = -1
    v44 := []int{21, 308, -1, 32767}
    var v45 uint64 = 42
    var v46 int = 32767
    v47 := GenS7{F0: 914}
    for fi48 := 0; fi48 < 17; fi48++ {
        v45 = v45
        v49 := v1[(fi48 * int(v3))]
        var v50 uint = 80
        _ = fi48
        _ = v49
        _ = v50
    }
    if false {
        v46 -= v46
        switch (v2 + (("foo bar" + "foo bar") + (((((("a" + v2) + "") + ((v2 + (v2 + "\n\t")) + "hello")) + v2) + "hello") + v2))) {
        case "foo bar":
            v51 := v1[v46]
            _ = v51
        case "\n\t":
            for fi52 := 0; fi52 < 1; fi52++ {
                var v53 int16 = 545
                _ = fi52
                _ = v53
            }
        }
    } else {
        v46 += v46
        if (1.5 >= 0.1) {
            switch -100 {
            case 9223372036854775807:
                if true {
                    v54 := GenS1{F0: ("" != "hello"), F1: -233}
                    if len(v44) > 0 {
                        v44[0] = 128
                    }
                    _ = v54
                } else {
                    for fi55 := 0; fi55 < 7; fi55++ {
                        switch v2 {
                        case "\n\t":
                            switch "\n\t" {
                            case "a":
                                v46 = -130
                            default:
                                switch ("" + "\n\t") {
                                case "foo bar":
                                    if (false && false) {
                                        _ = len(v1)
                                    } else {
                                        v3 = (v3 * (v3 >> ((0) & 63)))
                                    }
                                case "\n\t":
                                    for fi56 := 0; fi56 < 19; fi56++ {
                                        if ((float64(((3.14159 + 0.1) / (-2.25 / 1.0))) / 1.0) >= 0.1) {
                                            delete(v1, fi56)
                                        }
                                        if len(v44) > 0 {
                                            v44[0] = (fi56 / (((fi56 / (((-(v46))) | 1))) | 1))
                                        }
                                        _ = fi56
                                    }
                                default:
                                    fmt.Println(v2, v45, v46)
                                }
                            }
                        case "a":
                            v45 -= 1
                        case "hello":
                            switch v2 {
                            case "":
                                for fi57 := 0; fi57 < 10; fi57++ {
                                    switch v2 {
                                    case "foo bar":
                                        v3 += (v3 | (v3 ^ (^(v3))))
                                    default:
                                        var v58 uint8 = 128
                                        _ = v58
                                    }
                                    _ = fi57
                                }
                            default:
                                v59 := v1[(^(fi55))]
                                _ = v59
                            }
                        default:
                            var v60 int64 = -1
                            _ = v60
                        }
                        _ = fi55
                    }
                    v61 := GenS8{F0: (v45 ^ (v45 - 503)), F1: v45, F2: 18446744073709551615}
                    fmt.Println(v2)
                    _ = v61
                }
            case 127:
                if false {
                    delete(v1, v46)
                    if true {
                        for fi62 := 0; fi62 < 20; fi62++ {
                            for fi63 := 0; fi63 < 17; fi63++ {
                                var v64 string = v2
                                v65 := GenS3{F0: 0}
                                _ = len(v44)
                                _ = fi63
                                _ = v64
                                _ = v65
                            }
                            v66 := GenS8{F0: 4294967295, F1: (^(v45)), F2: 18446744073709551615}
                            var v67 string = "foo bar"
                            _ = fi62
                            _ = v66
                            _ = v67
                        }
                        v46 -= v46
                        var v68 uint8 = 127
                        _ = v68
                    }
                }
            case -910:
                fmt.Println(v45)
            default:
                for fi69 := 0; fi69 < 9; fi69++ {
                    var v70 float32 = (((3.14159 * (1e10 - float32(0.1))) / 1.0) * (float32(-2.25) * 1.5))
                    _ = fi69
                    _ = v70
                }
            }
        }
        switch v2 {
        case "hello":
            var v71 bool = false
            _ = v71
        case "a":
            v72 := GenS6{F0: 2}
            _ = v72
        case "foo bar":
            if (v45 <= (v45 / ((18446744073709551615) | 1))) {
                switch 42 {
                case -2147483648:
                    var v73 int32 = int32(v45)
                    _ = v73
                case 1:
                    fmt.Println(v46)
                }
                fmt.Println(v3, v46)
                v2 = "foo bar"
            } else {
                var v74 float32 = 1.5
                _ = v74
            }
        }
    }
    var v75 int8 = 1
    var v76 uint8 = 186
    if false {
        for fi77 := 0; fi77 < 11; fi77++ {
            _ = len(v44)
            _ = fi77
        }
        fmt.Println(v76, v45, v3)
        if len(v44) > 0 {
            v44[0] = 128
        }
    } else {
        var v78 uint16 = 742
        delete(v1, 127)
        _ = v78
    }
    var v79 int16 = -32768
    v80 := v1[(v46 & (^(v46)))]
    for fi81 := 0; fi81 < 8; fi81++ {
        for fi82 := 0; fi82 < 13; fi82++ {
            _ = len(v1)
            _ = fi82
        }
        fmt.Println(v80)
        v45 = 42
        _ = fi81
    }
    v75 = v75
    v83 := GenS3{F0: v79}
    var v84 string = v2
    for fi85 := 0; fi85 < 11; fi85++ {
        switch v84 {
        case "\n\t":
            var v86 uint8 = (v76 % ((v76) | 1))
            _ = v86
        }
        if true {
            v45 -= v45
            var v87 uint = 255
            var v88 string = ""
            _ = v87
            _ = v88
        } else {
            switch ("foo bar" + v84) {
            case "":
                if (!false) {
                    fmt.Println(v79)
                    var v89 int64 = -9223372036854775808
                    _ = v89
                } else {
                    if true {
                        var v90 uint8 = (v76 % ((v76) | 1))
                        v1[-827] = -453
                        _ = v90
                    } else {
                        for fi91 := 0; fi91 < 6; fi91++ {
                            v92 := GenS8{F0: v45, F1: 2, F2: 42}
                            delete(v1, (^(fi91)))
                            fmt.Println(v79)
                            v1[(fi91 * fi85)] = fi85
                            _ = len(v44)
                            _ = fi91
                            _ = v92
                        }
                    }
                    var v93 uint8 = v76
                    _ = v93
                }
            case "hello":
                var v94 uint16 = 42
                _ = v94
            default:
                v95 := GenS8{F0: 352, F1: (^(v45)), F2: genFn1()}
                _ = v95
            }
            v80 = (fi85 | (fi85 << ((769) & 63)))
        }
        v2 = "\n\t"
        _ = fi85
    }
    for fi96 := 0; fi96 < 10; fi96++ {
        for fi97 := 0; fi97 < 17; fi97++ {
            var v98 int = v80
            _ = fi97
            _ = v98
        }
        if true {
            var v99 bool = (true || false)
            v100 := GenS5{F0: v2, F1: v45, F2: uint8(v45)}
            _ = v100
            _ = v99
        } else {
            var v101 int32 = (v3 | (v3 * -2147483648))
            if false {
                for fi102 := 0; fi102 < 20; fi102++ {
                    if (v75 < v75) {
                        v2 = (v2 + v2)
                        if ((!false) && (-128 != v75)) {
                            var v103 int8 = (v75 ^ v75)
                            if (float32(1e10) >= (0.1 * 1.5)) {
                                var v104 uint = 2
                                v75 -= (v75 + v75)
                                fmt.Println(v80, v45, v75)
                                var v105 bool = false
                                var v106 int64 = -9223372036854775808
                                v107 := GenS3{F0: (v79 % ((42) | 1))}
                                v108 := v1[32767]
                                _ = v104
                                _ = v105
                                _ = v106
                                _ = v107
                                _ = v108
                            }
                            if ((^(v3)) > 42) {
                                v2 = "\n\t"
                                delete(v1, (fi102 + (^(fi96))))
                            }
                            _ = v103
                        }
                        var v109 float32 = ((0.0 + (((-2.25 + (float32(((1.5 / 1e10) * 0.1)) / ((float32((1e10 * -2.25)) + 1e10) / 1.5))) + -2.25) - (1e10 / -2.25))) + 1e10)
                        v110 := v1[int(v45)]
                        _ = v109
                        _ = v110
                    }
                    v75 = int8(v76)
                    _ = fi102
                }
            }
            if len(v44) > 0 {
                v44[0] = (fi96 + v80)
            }
            _ = v101
        }
        delete(v1, int(v79))
        _ = fi96
    }
    var v111 string = ((v84 + "a") + (v2 + "hello"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v111:", v111)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    fmt.Println("v79:", v79)
    fmt.Println("v80:", v80)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v111
    _ = v2
    _ = v3
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v75
    _ = v76
    _ = v79
    _ = v80
    _ = v83
    _ = v84
}
