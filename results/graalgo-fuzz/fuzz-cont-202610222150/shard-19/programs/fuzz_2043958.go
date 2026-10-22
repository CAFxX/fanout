package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 uint32; F2 uint16 }
type GenS4 struct { F0 bool; F1 uint32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int32; F1 uint64; F2 uint16 }
type GenS7 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS8 struct { F0 int64; F1 uint; F2 int8 }
func genFn0(p0 int8, p1 int8, p2 int64) uint {
    var v6 uint = 228
    if true {
        for fi7 := 0; fi7 < 3; fi7++ {
            var v8 uint = (v6 % (((v6 >> ((v6) & 63))) | 1))
            if ("foo bar" <= "\n\t") {
                var v9 uint64 = 2
                _ = v9
            }
            v6 += v8
            _ = fi7
            _ = v8
        }
    } else {
        var v10 uint = v6
        p2 = (p2 ^ 9223372036854775807)
        if true {
            if (240 <= 42) {
                var v11 int32 = 0
                var v12 uint64 = 0
                var v13 int16 = -32768
                _ = v11
                _ = v12
                _ = v13
            }
        }
        _ = v10
    }
    switch "\n\t" {
    case "":
        v14 := []int{127}
        _ = v14
    }
    var v15 uint16 = 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v6
    return (v6 & v6)
}

func main() {
    var v1 int8 = -116
    var v2 int64 = 654
    var v3 string = "foo bar"
    v4 := []int{127, 32767, -1}
    var v5 uint8 = 0
    var v16 uint = 810
    for fi17 := 0; fi17 < 7; fi17++ {
        for fi18 := 0; fi18 < 16; fi18++ {
            v19 := GenS1{F0: ((fi18 << ((v16) & 63)) < 626)}
            fmt.Println(fi17, v3)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi18
            _ = v19
        }
        v3 = ((v3 + v3) + v3)
        var v20 float64 = float64(((-2.25 - (0.0 * float32((float64(float32(3.14159)) - 0.1)))) / float32(1.5)))
        _ = fi17
        _ = v20
    }
    for fi21 := 0; fi21 < 6; fi21++ {
        if len(v4) > 0 {
            v4[0] = int(v5)
        }
        if (v3 > (v3 + v3)) {
            if len(v4) > 0 {
                v4[0] = -992
            }
        } else {
            var v22 int = fi21
            var v23 uint64 = 42
            fmt.Println(v2, v23)
            _ = v22
            _ = v23
        }
        if false {
            v3 = "\n\t"
            if (false || ((v2 >> (((v16 / ((0) | 1))) & 63)) != int64(v16))) {
                var v24 int32 = 0
                v1 = v1
                if (!(-2147483648 <= (v24 | v24))) {
                    v1 = (v1 % (((v1 % (((-(v1))) | 1))) | 1))
                    fmt.Println(v1, v3, v3)
                } else {
                    var v25 bool = (!(true && true))
                    _ = v25
                }
                _ = v24
            } else {
                var v26 bool = true
                switch (fi21 >> ((genFn0(v1, (v1 >> ((genFn0((v1 - (v1 & (v1 * v1))), int8(fi21), 42)) & 63)), v2)) & 63)) {
                case 0:
                    if true {
                        fmt.Println(fi21, v16)
                        v27 := GenS2{F0: (fi21 + (fi21 | (fi21 << (((v16 ^ genFn0(v1, (v1 % (((v1 % ((v1) | 1))) | 1)), -1))) & 63))))}
                        _ = v27
                    }
                default:
                    fmt.Println(v16)
                }
                _ = v26
            }
        } else {
            if ((v3 + "foo bar") != "") {
                for fi28 := 0; fi28 < 9; fi28++ {
                    switch (^(fi21)) {
                    case 32767:
                        v2 += v2
                    default:
                        if (((false || false) || (!true)) && true) {
                            v2 = 42
                        }
                    }
                    switch "\n\t" {
                    case "":
                        var v29 uint64 = 18446744073709551615
                        _ = v29
                    default:
                        if (!false) {
                            if (!((^(v16)) <= v16)) {
                                _ = len(v4)
                                var v30 uint8 = 0
                                if true {
                                    var v31 uint64 = uint64(v16)
                                    switch "\n\t" {
                                    case "\n\t":
                                        var v32 uint64 = v31
                                        _ = v32
                                    case "foo bar":
                                        if len(v4) > 0 {
                                            v4[0] = -86
                                        }
                                    }
                                    var v33 uint64 = (v31 | 4294967295)
                                    var v34 string = "a"
                                    if true {
                                        v3 = ""
                                        var v35 int16 = 32767
                                        fmt.Println(v1, v2, v30)
                                        if (true || (1e10 > (3.14159 + (1e10 - ((0.1 / float64((3.14159 / 1.5))) / (-2.25 - 0.1)))))) {
                                            v36 := GenS2{F0: fi28}
                                            _ = v36
                                        }
                                        v2 = int64(v5)
                                        var v37 uint8 = (v5 ^ uint8(v31))
                                        _ = v35
                                        _ = v37
                                    }
                                    _ = v31
                                    _ = v33
                                    _ = v34
                                }
                                _ = v30
                            } else {
                                var v38 uint8 = 65
                                var v39 float32 = 1e10
                                _ = v38
                                _ = v39
                            }
                            var v40 uint = v16
                            _ = v40
                        } else {
                            var v41 uint16 = 770
                            fmt.Println(v41, v16)
                            if (!true) {
                                if true {
                                    v1 = (v1 * (-(v1)))
                                    var v42 string = "\n\t"
                                    if len(v4) > 0 {
                                        v4[0] = fi21
                                    }
                                    if (false && (true && false)) {
                                        fmt.Println(fi28, v41, fi28)
                                        if (-1 > -891) {
                                            if (2147483647 == 1) {
                                                v42 = v42
                                            } else {
                                                v42 = v3
                                                v3 = (("hello" + v3) + "a")
                                                v1 = v1
                                                v4 = append(v4, (fi21 | fi28))
                                                _ = len(v4)
                                            }
                                        }
                                        v1 += (-(v1))
                                    }
                                    var v43 int64 = v2
                                    _ = v42
                                    _ = v43
                                } else {
                                    var v44 string = ("\n\t" + "a")
                                    v16 = uint(v5)
                                    if true {
                                        if false {
                                            if ((fi28 ^ fi21) < fi28) {
                                                var v45 float32 = 3.14159
                                                v41 -= (^(v41))
                                                v46 := GenS2{F0: fi21}
                                                v47 := GenS3{F0: (fi21 >> ((0) & 63)), F1: uint32(v41), F2: v41}
                                                if true {
                                                    fmt.Println(v2)
                                                    switch fi21 {
                                                    case -9223372036854775808:
                                                        v48 := GenS4{F0: true, F1: 4294967295}
                                                        _ = v48
                                                    case -134:
                                                        _ = len(v4)
                                                    case -138:
                                                        fmt.Println(v41)
                                                    }
                                                    v2 += 502
                                                }
                                                _ = v45
                                                _ = v46
                                                _ = v47
                                            } else {
                                                _ = len(v4)
                                                var v49 bool = (!false)
                                                v2 += (^(v2))
                                                _ = v49
                                            }
                                        } else {
                                            fmt.Println(fi21, v41)
                                            var v50 bool = (v16 == (^(v16)))
                                            switch (v3 + (((v3 + ("" + "hello")) + ("" + v44)) + "")) {
                                            case "\n\t":
                                                fmt.Println(fi21, v2)
                                            default:
                                                fmt.Println(v50, fi21)
                                            }
                                            _ = v50
                                        }
                                        if (fi28 > fi28) {
                                            var v51 uint32 = 1
                                            var v52 uint8 = 42
                                            switch "hello" {
                                            case "foo bar":
                                                v4 = append(v4, (fi28 << ((v16) & 63)))
                                            case "hello":
                                                if len(v4) > 0 {
                                                    v4[0] = -688
                                                }
                                            }
                                            _ = v51
                                            _ = v52
                                        }
                                    } else {
                                        var v53 uint32 = 637
                                        var v54 string = (v3 + ((((v44 + "") + v44) + "foo bar") + v44))
                                        if len(v4) > 0 {
                                            v4[0] = (fi21 & fi21)
                                        }
                                        switch (fi28 >> ((176) & 63)) {
                                        case 42:
                                            var v55 int64 = (^(v2))
                                            var v56 float32 = 0.0
                                            _ = v55
                                            _ = v56
                                        case 2147483647:
                                            v1 += -52
                                        case -564:
                                            if len(v4) > 0 {
                                                v4[0] = 128
                                            }
                                        default:
                                            fmt.Println(v54, v16)
                                        }
                                        _ = v53
                                        _ = v54
                                    }
                                    _ = v44
                                }
                            } else {
                                v2 += -9223372036854775808
                            }
                            _ = v41
                        }
                    }
                    var v57 uint16 = 527
                    var v58 uint = (^(v16))
                    v59 := GenS3{F0: -46, F1: 2147483647, F2: v57}
                    _ = fi28
                    _ = v57
                    _ = v58
                    _ = v59
                }
                v3 = v3
            }
        }
        _ = fi21
    }
    fmt.Println(v16, v2)
    fmt.Println(v3, v2)
    var v60 int32 = 42
    v5 = v5
    v61 := GenS5{F0: v2}
    v1 -= (^(v1))
    v5 -= v5
    var v62 int16 = 1000
    v62 = v62
    var v63 uint16 = 967
    if (!(220 != 18446744073709551615)) {
        if (!true) {
            if false {
                var v64 uint = genFn0(v1, (v1 << (((^(v16))) & 63)), (^(v2)))
                fmt.Println(v2, v2)
                v64 += v64
                _ = v64
            } else {
                for fi65 := 0; fi65 < 11; fi65++ {
                    for fi66 := 0; fi66 < 4; fi66++ {
                        var v67 uint16 = 42
                        switch ((v3 + v3) + (v3 + "foo bar")) {
                        case "a":
                            v60 = -1
                        case "hello":
                            var v68 uint = (v16 & v16)
                            _ = v68
                        case "foo bar":
                            var v69 int8 = v1
                            _ = v69
                        default:
                            if false {
                                var v70 float64 = 0.0
                                _ = v70
                            } else {
                                if false {
                                    v63 = 42
                                    var v71 bool = (v3 < "a")
                                    var v72 float32 = 0.1
                                    _ = v71
                                    _ = v72
                                } else {
                                    if (!false) {
                                        var v73 string = v3
                                        v74 := GenS4{F0: ((!(true || ((true && (v16 <= v16)) && false))) || false), F1: 0}
                                        fmt.Println(v1, v5)
                                        _ = v73
                                        _ = v74
                                    }
                                    if true {
                                        if (2147483647 <= 4294967295) {
                                            v5 = v5
                                            if ((!((v2 << ((uint(v5)) & 63)) > (v2 - -9223372036854775808))) || (-2.25 != 0.1)) {
                                                v75 := GenS1{F0: true}
                                                _ = v75
                                            }
                                            var v76 float32 = -2.25
                                            _ = v76
                                        }
                                        if (((18446744073709551615 == 411) && (false || (v3 <= "\n\t"))) || ((143 >= 160) && false)) {
                                            if (false || (0 > -32768)) {
                                                switch fi66 {
                                                case 809:
                                                    fmt.Println(v67, v60, v16)
                                                default:
                                                    v2 += (v2 & int64(v60))
                                                }
                                            } else {
                                                if false {
                                                    if (v63 == v67) {
                                                        var v77 float32 = 3.14159
                                                        _ = v77
                                                    }
                                                    v78 := GenS1{F0: true}
                                                    _ = v78
                                                } else {
                                                    v79 := GenS6{F0: (v60 ^ -1), F1: 757, F2: 345}
                                                    var v80 int64 = (v2 >> ((18446744073709551615) & 63))
                                                    switch (fi65 * fi65) {
                                                    case 32768:
                                                        v80 -= (v2 % ((-1) | 1))
                                                    case -1:
                                                        v81 := GenS7{F0: 4294967295, F1: v60, F2: v63}
                                                        _ = v81
                                                    }
                                                    _ = v79
                                                    _ = v80
                                                }
                                            }
                                        }
                                        v1 = v1
                                    }
                                    var v82 float32 = ((0.1 / float32(3.14159)) * 3.14159)
                                    _ = v82
                                }
                                v4 = append(v4, (-(fi66)))
                                if (4294967295 > 42) {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    v83 := GenS1{F0: true}
                                    v84 := GenS4{F0: true, F1: 0}
                                    _ = v83
                                    _ = v84
                                } else {
                                    var v85 int8 = v1
                                    var v86 uint = 0
                                    fmt.Println(v1)
                                    _ = v85
                                    _ = v86
                                }
                            }
                            _ = len(v4)
                        }
                        _ = fi66
                        _ = v67
                    }
                    for fi87 := 0; fi87 < 6; fi87++ {
                        if len(v4) > 0 {
                            v4[0] = (fi87 & fi65)
                        }
                        var v88 int = -9223372036854775808
                        _ = fi87
                        _ = v88
                    }
                    v89 := GenS7{F0: 65535, F1: v60, F2: v63}
                    _ = fi65
                    _ = v89
                }
                v3 = v3
            }
            v63 += 0
        } else {
            _ = len(v4)
            v60 = 0
            var v90 float32 = 0.1
            _ = v90
        }
        var v91 uint = (v16 * v16)
        _ = v91
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if (!(42 >= 2147483648)) {
            v92 := GenS8{F0: (-(v2)), F1: (v16 >> ((genFn0(v1, v1, (v2 | v2))) & 63)), F2: v1}
            if (3.14159 == ((((0.1 / 0.1) / (0.0 - 0.1)) - (((1e10 / float32(1e10)) * 0.0) + (0.0 / 1.5))) / 1.5)) {
                v93 := GenS1{F0: ((v2 << ((v16) & 63)) <= (v2 >> ((1) & 63)))}
                if false {
                    for fi94 := 0; fi94 < 3; fi94++ {
                        if len(v4) > 0 {
                            v4[0] = int(v63)
                        }
                        v62 = v62
                        var v95 uint16 = (v63 | (v63 >> ((2) & 63)))
                        _ = fi94
                        _ = v95
                    }
                }
                v96 := GenS6{F0: 0, F1: 18446744073709551615, F2: v63}
                _ = v93
                _ = v96
            }
            for fi97 := 0; fi97 < 16; fi97++ {
                var v98 bool = ((!(v62 > v62)) || (!(v3 >= (v3 + (v3 + v3)))))
                v4 = append(v4, (fi97 * fi97))
                _ = fi97
                _ = v98
            }
            _ = v92
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v60:", v60)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v60
    _ = v61
    _ = v62
    _ = v63
}
