package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 int32; F2 uint8 }
type GenS4 struct { F0 uint8; F1 uint16 }
type GenS5 struct { F0 int32; F1 int8 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int16; F1 int64; F2 uint16 }
type GenS8 struct { F0 int64; F1 uint16; F2 uint8 }
type GenS9 struct { F0 int; F1 int64 }
func genFn0() int16 {
    var v4 bool = true
    var v5 uint16 = 65535
    if v4 {
        fmt.Println(v5)
        if (-128 > -128) {
            var v6 int64 = -9223372036854775808
            for fi7 := 0; fi7 < 10; fi7++ {
                var v8 float64 = -2.25
                _ = fi7
                _ = v8
            }
            _ = v6
        } else {
            if ("" > (("" + "a") + "hello")) {
                var v9 float64 = ((float64(1e10) / 0.1) + 1.5)
                v4 = v4
                _ = v9
            }
        }
        v10 := GenS1{F0: 1, F1: 218, F2: 1}
        _ = v10
    } else {
        for fi11 := 0; fi11 < 10; fi11++ {
            v4 = (!(2147483648 <= 42))
            if false {
                v5 = 32767
                if v4 {
                    for fi12 := 0; fi12 < 19; fi12++ {
                        v13 := GenS1{F0: 18446744073709551615, F1: 255, F2: -339}
                        if (!(!(128 > 218))) {
                            fmt.Println(fi12)
                            if (fi11 <= 622) {
                                v4 = true
                            }
                        } else {
                            v14 := GenS2{F0: 128}
                            _ = v14
                        }
                        if (47 >= 1) {
                            var v15 float32 = 0.0
                            v4 = ((1 != 65535) && (99 > -1))
                            _ = v15
                        }
                        _ = fi12
                        _ = v13
                    }
                } else {
                    v5 -= (v5 * v5)
                    var v16 uint8 = 208
                    _ = v16
                }
            } else {
                var v17 int8 = int8(fi11)
                if ((!v4) && (!false)) {
                    v17 = 1
                    v18 := []int{fi11, fi11, int(v17)}
                    if len(v18) >= 2 {
                        v18 = v18[1:len(v18)]
                    }
                    _ = v18
                } else {
                    if true {
                        v19 := GenS1{F0: 2, F1: 2, F2: 902}
                        _ = v19
                    } else {
                        v4 = (!(v17 > v17))
                    }
                    for fi20 := 0; fi20 < 9; fi20++ {
                        var v21 int8 = (v17 % (((v17 << ((613) & 63))) | 1))
                        _ = fi20
                        _ = v21
                    }
                }
                for fi22 := 0; fi22 < 3; fi22++ {
                    var v23 uint8 = 127
                    fmt.Println(v17, v23)
                    _ = fi22
                    _ = v23
                }
                _ = v17
            }
            var v24 int32 = 471
            _ = fi11
            _ = v24
        }
        for fi25 := 0; fi25 < 9; fi25++ {
            v26 := GenS1{F0: 2, F1: 1, F2: 1}
            v5 = v5
            v27 := GenS3{F0: 0, F1: 2147483647, F2: 0}
            _ = fi25
            _ = v26
            _ = v27
        }
        for fi28 := 0; fi28 < 3; fi28++ {
            switch fi28 {
            case 32768:
                var v29 string = ("hello" + (("foo bar" + "hello") + ""))
                _ = v29
            case -2147483648:
                var v30 float32 = -2.25
                _ = v30
            case 42:
                var v31 uint8 = 81
                _ = v31
            default:
                var v32 float32 = 0.0
                _ = v32
            }
            var v33 float64 = 0.1
            switch fi28 {
            case 32768:
                if (203 <= 0) {
                    for fi34 := 0; fi34 < 4; fi34++ {
                        v35 := GenS4{F0: 1, F1: 493}
                        _ = fi34
                        _ = v35
                    }
                    v4 = v4
                }
            case -1:
                v36 := []int{-169, 9223372036854775807, 140, 0}
                _ = v36
            }
            _ = fi28
            _ = v33
        }
    }
    _ = v4
    _ = v5
    return -32768
}
func genFn1(p0 int64, p1 int32) int16 {
    p0 = (p0 & p0)
    p1 = p1
    p0 = (-(p0))
    if (true && true) {
        v37 := []int{781}
        for fi38 := 0; fi38 < 10; fi38++ {
            v39 := GenS5{F0: (p1 << ((255) & 63)), F1: 100}
            if false {
                var v40 float64 = 1e10
                p0 = -1
                _ = v40
            } else {
                p0 -= p0
            }
            var v41 int32 = p1
            _ = fi38
            _ = v39
            _ = v41
        }
        for fi42 := 0; fi42 < 16; fi42++ {
            var v43 uint32 = 42
            _ = fi42
            _ = v43
        }
        _ = v37
    }
    _ = p0
    _ = p1
    return 32767
}

func main() {
    var v1 uint16 = 327
    var v2 uint64 = 710
    v3 := make(map[string]string)
    fmt.Println(v1, v2)
    var v44 uint8 = 42
    switch ("a" + "\n\t") {
    case "\n\t":
        for fi45 := 0; fi45 < 3; fi45++ {
            fmt.Println(v2, v2)
            _ = fi45
        }
    case "a":
        v44 = (v44 >> ((2) & 63))
    case "foo bar":
        for fi46 := 0; fi46 < 17; fi46++ {
            v47 := v3["foo bar"]
            _ = fi46
            _ = v47
        }
    default:
        switch (((("a" + "") + (("foo bar" + "a") + (("" + "hello") + "a"))) + "") + ("foo bar" + "hello")) {
        case "foo bar":
            v48 := GenS5{F0: 568, F1: -1}
            _ = v48
        default:
            fmt.Println(v1)
        }
    }
    var v49 int16 = -878
    v1 += v1
    v44 -= (v44 * v44)
    v3[""] = "hello"
    if true {
        if (((true || false) && true) && true) {
            var v50 uint8 = v44
            var v51 float64 = ((float64(0.0) / 1e10) + 1.5)
            for fi52 := 0; fi52 < 4; fi52++ {
                v2 += 4294967295
                v50 = (v50 ^ (^(v50)))
                v51 += v51
                _ = fi52
            }
            _ = v50
            _ = v51
        }
        v53 := GenS6{F0: 42}
        _ = v53
    } else {
        var v54 string = ""
        var v55 uint = 255
        for fi56 := 0; fi56 < 13; fi56++ {
            if true {
                var v57 uint = v55
                fmt.Println(v1)
                switch fi56 {
                case 231:
                    for fi58 := 0; fi58 < 18; fi58++ {
                        if true {
                            v59 := []int{(fi58 - (fi58 + 1))}
                            v60 := GenS2{F0: v44}
                            fmt.Println(fi56, fi58, v1)
                            v3[v54] = (v54 + v54)
                            _ = v59
                            _ = v60
                        } else {
                            v44 += 128
                        }
                        _ = fi58
                    }
                case 127:
                    for fi61 := 0; fi61 < 2; fi61++ {
                        var v62 uint16 = 65535
                        _ = fi61
                        _ = v62
                    }
                case -1:
                    switch (fi56 - fi56) {
                    case 1:
                        for fi63 := 0; fi63 < 19; fi63++ {
                            v57 += uint(v1)
                            v64 := GenS7{F0: -554, F1: 1, F2: v1}
                            var v65 uint64 = (v2 ^ v2)
                            _ = fi63
                            _ = v64
                            _ = v65
                        }
                    case 32768:
                        switch ("a" + v54) {
                        case "foo bar":
                            fmt.Println(v57)
                        }
                    default:
                        v66 := GenS5{F0: int32(v55), F1: -128}
                        _ = v66
                    }
                }
                _ = v57
            }
            delete(v3, v54)
            _ = fi56
        }
        _ = v54
        _ = v55
    }
    fmt.Println(v1, v49)
    switch -250 {
    case 42:
        var v67 uint8 = 1
        _ = v67
    case 4294967295:
        for fi68 := 0; fi68 < 10; fi68++ {
            if true {
                var v69 uint8 = v44
                if (v1 == v1) {
                    fmt.Println(v44)
                    if true {
                        switch (fi68 + fi68) {
                        case -937:
                            var v70 uint = 687
                            _ = v70
                        default:
                            if false {
                                fmt.Println(v69, v49, v1)
                                v1 += v1
                                for fi71 := 0; fi71 < 12; fi71++ {
                                    if ("a" < "hello") {
                                        var v72 uint16 = (v1 * v1)
                                        v44 = v44
                                        v44 = v44
                                        _ = v72
                                    }
                                    _ = fi71
                                }
                            } else {
                                fmt.Println(fi68, v69)
                            }
                        }
                    } else {
                        if true {
                            var v73 int16 = (v49 & (v49 >> ((514) & 63)))
                            for fi74 := 0; fi74 < 7; fi74++ {
                                if (1 != -9223372036854775808) {
                                    var v75 float64 = 0.1
                                    _ = v75
                                }
                                _ = fi74
                            }
                            _ = v73
                        }
                        var v76 uint = 4294967295
                        v77 := GenS8{F0: 42, F1: v1, F2: 230}
                        _ = v76
                        _ = v77
                    }
                    var v78 float32 = 1.5
                    _ = v78
                } else {
                    if (42 != -75) {
                        v49 = genFn0()
                        v79 := GenS2{F0: 2}
                        _ = v79
                    } else {
                        for fi80 := 0; fi80 < 8; fi80++ {
                            fmt.Println(v44, fi68, v1)
                            fmt.Println(v49, fi80)
                            if false {
                                var v81 int8 = -1
                                v82 := v3[((("\n\t" + (("\n\t" + "\n\t") + ("a" + ("a" + ("" + ("foo bar" + (("a" + "hello") + ""))))))) + "") + ("\n\t" + "foo bar"))]
                                _ = v81
                                _ = v82
                            } else {
                                fmt.Println(v49)
                                if ((!true) || false) {
                                    var v83 float32 = 3.14159
                                    var v84 float32 = (3.14159 + 1.5)
                                    v1 = (v1 | (v1 * v1))
                                    var v85 uint = 1
                                    _ = v83
                                    _ = v84
                                    _ = v85
                                } else {
                                    v86 := GenS5{F0: -1, F1: 0}
                                    v87 := []int{(fi68 << ((0) & 63)), -568, (fi68 | (fi68 - (fi68 % (((fi80 & fi68)) | 1))))}
                                    fmt.Println(fi68, v1, v1)
                                    _ = v86
                                    _ = v87
                                }
                            }
                            v88 := GenS9{F0: (-(fi80)), F1: 821}
                            var v89 string = "foo bar"
                            v90 := []int{fi68, (-(fi80)), -9223372036854775808, (fi80 << ((266) & 63))}
                            if false {
                                _ = len(v3)
                            } else {
                                v49 -= int16(fi68)
                                var v91 float32 = (3.14159 / 1.0)
                                switch -25 {
                                case 504:
                                    if (!true) {
                                        var v92 uint8 = v44
                                        _ = len(v3)
                                        _ = v92
                                    } else {
                                        _ = len(v3)
                                        var v93 int64 = 42
                                        if true {
                                            v69 = 23
                                        }
                                        _ = v93
                                    }
                                    v94 := GenS9{F0: (fi68 | fi68), F1: 999}
                                    switch v89 {
                                    case "\n\t":
                                        if (v1 != 255) {
                                            fmt.Println(v44)
                                        }
                                        delete(v3, (((v89 + "a") + (v89 + "\n\t")) + ((("\n\t" + "foo bar") + v89) + v89)))
                                        var v95 uint16 = v1
                                        _ = v95
                                    case "a":
                                        v2 = 969
                                    case "":
                                        delete(v3, v89)
                                    }
                                    _ = v94
                                case 991:
                                    delete(v3, "foo bar")
                                default:
                                    if (255 <= (v69 | 239)) {
                                        switch -9223372036854775808 {
                                        case 42:
                                            v96 := GenS4{F0: (^(v44)), F1: v1}
                                            _ = v96
                                        }
                                    } else {
                                        v97 := GenS5{F0: -1, F1: -46}
                                        _ = v97
                                    }
                                }
                                _ = v91
                            }
                            var v98 int16 = -341
                            fmt.Println(fi80, v2)
                            _ = fi80
                            _ = v88
                            _ = v89
                            _ = v90
                            _ = v98
                        }
                        if (v1 >= (v1 >> ((uint(v69)) & 63))) {
                            var v99 bool = (3.14159 >= float64(0.0))
                            _ = v99
                        } else {
                            v49 += -1
                            v69 += v69
                        }
                    }
                }
                fmt.Println(fi68, fi68, v69)
                _ = v69
            } else {
                if true {
                    if (1e10 != 1e10) {
                        if false {
                            for fi100 := 0; fi100 < 6; fi100++ {
                                var v101 uint8 = v44
                                _ = fi100
                                _ = v101
                            }
                        }
                        v2 = (v2 ^ (^(v2)))
                    } else {
                        switch ("" + "") {
                        case "\n\t":
                            var v102 bool = true
                            _ = v102
                        case "hello":
                            var v103 string = ("\n\t" + "foo bar")
                            _ = v103
                        case "":
                            fmt.Println(fi68, v49)
                        }
                        var v104 int64 = -1
                        _ = v104
                    }
                }
            }
            var v105 int16 = v49
            _ = fi68
            _ = v105
        }
    }
    for fi106 := 0; fi106 < 8; fi106++ {
        if (!true) {
            var v107 float64 = -2.25
            _ = v107
        } else {
            v108 := GenS9{F0: (fi106 & fi106), F1: 9223372036854775807}
            _ = v108
        }
        v3["a"] = "a"
        _ = fi106
    }
    fmt.Println(v1)
    var v109 float32 = 0.1
    fmt.Println(v49)
    if true {
        fmt.Println(v44, v1, v2)
        v2 -= v2
        if true {
            var v110 float32 = 3.14159
            _ = v110
        } else {
            for fi111 := 0; fi111 < 18; fi111++ {
                var v112 uint8 = (^(v44))
                v49 += v49
                _ = fi111
                _ = v112
            }
            var v113 int16 = (v49 & (^(v49)))
            _ = v113
        }
    } else {
        v44 = v44
        switch 1 {
        case 951:
            if false {
                v114 := []int{2147483647}
                var v115 float32 = (v109 - ((1.5 * float32((0.0 + float64(((float32(float64(float32(1e10))) + v109) + v109))))) * (3.14159 / (float32(0.0) - (v109 + 0.1)))))
                _ = v114
                _ = v115
            } else {
                var v116 uint64 = (v2 ^ (v2 | 56))
                v117 := []int{}
                _ = v116
                _ = v117
            }
        }
    }
    fmt.Println(v44)
    fmt.Println("v1:", v1)
    fmt.Println("v109:", v109)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v44:", v44)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v109
    _ = v2
    _ = v3
    _ = v44
    _ = v49
}
