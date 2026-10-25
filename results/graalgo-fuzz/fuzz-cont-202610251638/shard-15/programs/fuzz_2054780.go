package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 string; F1 uint; F2 int32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint32; F1 uint32 }
type GenS8 struct { F0 uint8; F1 int8; F2 uint16 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 bool; F1 uint8 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 uint8; F1 uint32; F2 int16 }
type GenS13 struct { F0 uint8; F1 uint16 }
type GenS14 struct { F0 bool; F1 bool }
func genFn0(p0 int) uint8 {
    v5 := GenS1{F0: 44}
    for fi6 := 0; fi6 < 6; fi6++ {
        v7 := []int{}
        v8 := GenS1{F0: 39}
        p0 -= fi6
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := []int{(p0 | (^(p0)))}
    switch "foo bar" {
    case "foo bar":
        if true {
            var v10 float64 = 1.5
            _ = v10
        }
    case "":
        var v11 uint32 = 2147483648
        _ = v11
    default:
        var v12 uint = 42
        _ = v12
    }
    _ = p0
    _ = v5
    _ = v9
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int = 291
    var v3 float32 = (3.14159 / 1.0)
    var v4 uint64 = 18446744073709551615
    for fi13 := 0; fi13 < 10; fi13++ {
        v14 := []int{(v2 % (((-(v2))) | 1)), v2}
        v3 = v3
        for fi15 := 0; fi15 < 6; fi15++ {
            if false {
                if (!(-482 >= 0)) {
                    if false {
                        switch (fi15 | fi15) {
                        case -846:
                            v3 = 1e10
                        case -2147483648:
                            if false {
                                v4 += 4294967295
                            } else {
                                var v16 uint = 604
                                if len(v14) >= 2 {
                                    v14 = v14[1:len(v14)]
                                }
                                var v17 int32 = 2147483647
                                switch (("\n\t" + "foo bar") + "\n\t") {
                                case "":
                                    if true {
                                        var v18 uint64 = v4
                                        var v19 float32 = v3
                                        _ = v18
                                        _ = v19
                                    } else {
                                        fmt.Println(v16, v3, v3)
                                    }
                                case "a":
                                    v3 += 0.1
                                case "\n\t":
                                    v20 := GenS1{F0: 128}
                                    if false {
                                        var v21 uint16 = 32767
                                        fmt.Println(v17, fi15)
                                        v21 += v21
                                        _ = v21
                                    } else {
                                        var v22 int32 = 42
                                        var v23 int32 = (^(v17))
                                        v4 -= (v4 / ((v4) | 1))
                                        _ = v22
                                        _ = v23
                                    }
                                    var v24 bool = true
                                    _ = v20
                                    _ = v24
                                default:
                                    v17 -= (v17 ^ (v17 ^ v17))
                                }
                                _ = v16
                                _ = v17
                            }
                        default:
                            if false {
                                if len(v14) >= 2 {
                                    v14 = v14[1:len(v14)]
                                }
                                fmt.Println(fi13, v3)
                                var v25 float64 = (((1.5 - (float64(0.0) - (0.0 + (0.0 / (3.14159 + ((((0.0 - 1e10) + 3.14159) / 1.0) * 1e10)))))) - (1.5 * (float64(float32(1.5)) * 1.5))) * -2.25)
                                v26 := GenS1{F0: 0}
                                v1[58] = false
                                _ = v25
                                _ = v26
                            } else {
                                switch (v2 << ((65535) & 63)) {
                                case -9223372036854775808:
                                    var v27 string = "hello"
                                    _ = v27
                                case 42:
                                    var v28 int8 = -17
                                    _ = v28
                                case 9223372036854775807:
                                    var v29 uint = 255
                                    _ = v29
                                default:
                                    var v30 uint8 = 226
                                    _ = v30
                                }
                            }
                        }
                        fmt.Println(v2)
                        v4 += (v4 * (v4 - v4))
                    }
                    switch ("hello" + "") {
                    case "a":
                        switch 163 {
                        case -1:
                            if false {
                                v3 += (0.1 / float32(0.1))
                                if false {
                                    switch 4294967295 {
                                    case 395:
                                        if false {
                                            fmt.Println(fi13, fi13)
                                            var v31 uint32 = 2147483648
                                            v4 = v4
                                            _ = v31
                                        } else {
                                            if ((100 > 0) && ((v3 + 1.5) < v3)) {
                                                v4 -= (v4 ^ v4)
                                                if (false || true) {
                                                    v2 = fi13
                                                    if false {
                                                        v3 += 0.1
                                                    }
                                                }
                                                fmt.Println(v2)
                                            } else {
                                                v32 := GenS2{F0: 127, F1: (v2 + 32768)}
                                                v33 := GenS3{F0: 255, F1: 2}
                                                v4 -= 4294967295
                                                _ = v32
                                                _ = v33
                                            }
                                            var v34 int16 = -696
                                            _ = v34
                                        }
                                    case 4294967295:
                                        delete(v1, 0)
                                        var v35 float32 = v3
                                        _ = v35
                                    default:
                                        v36 := GenS4{F0: "foo bar", F1: 606, F2: -481}
                                        _ = v36
                                    }
                                    v2 = fi15
                                }
                                v4 = v4
                            }
                        case 0:
                            if true {
                                if (!((2147483648 > 157) || (float32((1.5 * (0.0 - 3.14159))) >= 0.1))) {
                                    switch (fi15 << ((0) & 63)) {
                                    case 254:
                                        v37 := GenS5{F0: "foo bar"}
                                        _ = v37
                                    case 32768:
                                        var v38 float64 = 1.5
                                        _ = v38
                                    case 1:
                                        var v39 uint64 = v4
                                        _ = v39
                                    }
                                } else {
                                    if false {
                                        if len(v14) > 0 {
                                            v14[0] = (fi15 * fi15)
                                        }
                                        v40 := GenS4{F0: "\n\t", F1: 37, F2: 42}
                                        v4 = v4
                                        _ = v40
                                    } else {
                                        var v41 float64 = (-2.25 - 3.14159)
                                        var v42 bool = (v3 <= 0.1)
                                        delete(v1, 2147483647)
                                        _ = v41
                                        _ = v42
                                    }
                                    var v43 float32 = ((v3 / 1e10) * float32(-2.25))
                                    var v44 uint8 = 1
                                    _ = v43
                                    _ = v44
                                }
                                fmt.Println(fi13)
                                if true {
                                    v2 -= -578
                                } else {
                                    fmt.Println(fi13)
                                    v3 = v3
                                }
                                var v45 int8 = 127
                                switch ("\n\t" + ("a" + ("hello" + ("a" + "hello")))) {
                                case "hello":
                                    _ = len(v1)
                                default:
                                    fmt.Println(fi13, fi13, v2)
                                }
                                _ = v45
                            }
                        case 32767:
                            if true {
                                fmt.Println(fi15)
                            } else {
                                v2 -= 0
                                v2 += (v2 * -99)
                            }
                        default:
                            if (false && (!true)) {
                                v2 -= v2
                                if true {
                                    v46 := GenS4{F0: "", F1: 4294967295, F2: 42}
                                    if ((-206 > -9223372036854775808) || (!true)) {
                                        var v47 int8 = 113
                                        _ = v47
                                    }
                                    if (!true) {
                                        var v48 int8 = -113
                                        v49 := GenS4{F0: (("a" + "") + ("foo bar" + "a")), F1: 0, F2: -1}
                                        switch ("hello" + "\n\t") {
                                        case "\n\t":
                                            v3 = v3
                                        default:
                                            switch fi15 {
                                            case -9223372036854775808:
                                                _ = len(v14)
                                            default:
                                                v50 := GenS6{F0: (^(v4))}
                                                _ = v50
                                            }
                                        }
                                        _ = v48
                                        _ = v49
                                    } else {
                                        v2 = v2
                                        if (1e10 < (1e10 - (0.1 / (-2.25 * (3.14159 + -2.25))))) {
                                            v2 += v2
                                        } else {
                                            switch ("foo bar" + ((((("hello" + "hello") + (("hello" + "foo bar") + ("" + ""))) + "") + "\n\t") + "foo bar")) {
                                            case "":
                                                v14 = append(v14, fi13)
                                            case "a":
                                                v4 = (^(v4))
                                            }
                                        }
                                        v51 := GenS7{F0: 477, F1: 137}
                                        _ = v51
                                    }
                                    _ = v46
                                } else {
                                    fmt.Println(fi15)
                                }
                            }
                        }
                    default:
                        v52 := GenS7{F0: 42, F1: 2}
                        _ = v52
                    }
                    if true {
                        if true {
                            v53 := GenS7{F0: 2, F1: 65535}
                            v3 = float32(0.0)
                            _ = v53
                        } else {
                            v3 = 0.1
                        }
                    } else {
                        delete(v1, 2)
                        switch ("foo bar" + "a") {
                        case "a":
                            v14 = append(v14, -1)
                        case "\n\t":
                            switch (fi15 << ((4294967295) & 63)) {
                            case 32768:
                                fmt.Println(fi15)
                            }
                        default:
                            var v54 uint64 = 826
                            v55 := GenS1{F0: 128}
                            _ = v54
                            _ = v55
                        }
                        if (314 > 42) {
                            v56 := GenS8{F0: uint8(v4), F1: 100, F2: 2}
                            _ = v56
                        } else {
                            if ((v4 / ((v4) | 1)) == v4) {
                                v57 := GenS4{F0: "", F1: 578, F2: -449}
                                var v58 uint = 18446744073709551615
                                fmt.Println(fi15, v2)
                                if false {
                                    if (v3 >= 1e10) {
                                        var v59 string = "\n\t"
                                        _ = v59
                                    } else {
                                        if false {
                                            v60 := GenS7{F0: 204, F1: 2}
                                            _ = len(v14)
                                            _ = v60
                                        }
                                        v14 = append(v14, 1)
                                    }
                                    v4 = v4
                                    v61 := GenS9{F0: -1}
                                    _ = v61
                                } else {
                                    var v62 uint64 = (v4 >> ((v58) & 63))
                                    _ = v62
                                }
                                if true {
                                    v63 := GenS10{F0: true, F1: 42}
                                    _ = v63
                                } else {
                                    if (v3 < float32((0.1 - 1e10))) {
                                        var v64 uint32 = 361
                                        var v65 uint8 = 127
                                        v66 := GenS4{F0: "a", F1: 748, F2: 2147483647}
                                        v2 += (fi15 | (fi13 - fi13))
                                        _ = v64
                                        _ = v65
                                        _ = v66
                                    }
                                    var v67 int16 = 32767
                                    v2 = (-(v2))
                                    if false {
                                        switch (fi13 | (v2 - (-(fi15)))) {
                                        case -910:
                                            var v68 string = ("a" + ("a" + "a"))
                                            _ = v68
                                        case -183:
                                            if len(v14) > 0 {
                                                v14[0] = (fi13 % (((v2 % ((456) | 1))) | 1))
                                            }
                                        }
                                        v69 := GenS11{F0: true}
                                        _ = v69
                                    } else {
                                        if len(v14) >= 2 {
                                            v14 = v14[1:len(v14)]
                                        }
                                        fmt.Println(v58)
                                        fmt.Println(v67, v58, fi13)
                                    }
                                    var v70 float32 = 1.5
                                    var v71 uint64 = (v4 | v4)
                                    _ = v67
                                    _ = v70
                                    _ = v71
                                }
                                var v72 uint32 = 979
                                _ = v57
                                _ = v58
                                _ = v72
                            }
                        }
                    }
                    v73 := GenS12{F0: 255, F1: 65535, F2: -32768}
                    _ = v73
                } else {
                    if (v3 <= float32(-2.25)) {
                        fmt.Println(v4, v2, fi13)
                    } else {
                        v4 += (^(v4))
                        if false {
                            if (true && (!true)) {
                                var v74 bool = false
                                _ = v74
                            } else {
                                if true {
                                    var v75 uint16 = 255
                                    v75 += 2
                                    _ = v75
                                }
                                var v76 int32 = 0
                                _ = v76
                            }
                        } else {
                            fmt.Println(fi13)
                        }
                    }
                    if (v4 > v4) {
                        v2 -= 4294967295
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                        switch 32767 {
                        case 1:
                            fmt.Println(v4)
                        default:
                            if (true || true) {
                                var v77 int32 = 0
                                switch (-(v2)) {
                                case -290:
                                    fmt.Println(fi15)
                                case -1:
                                    var v78 int64 = -108
                                    _ = v78
                                case 128:
                                    var v79 float64 = -2.25
                                    _ = v79
                                default:
                                    var v80 uint64 = v4
                                    _ = v80
                                }
                                _ = v77
                            }
                        }
                    }
                    var v81 bool = false
                    fmt.Println(fi15, v3)
                    v3 = v3
                    _ = v81
                }
                if false {
                    v4 -= (v4 ^ (v4 / (((v4 + v4)) | 1)))
                    if false {
                        fmt.Println(v4)
                        var v82 uint8 = 45
                        _ = v82
                    } else {
                        v83 := GenS10{F0: false, F1: genFn0(fi13)}
                        v84 := GenS13{F0: genFn0(v2), F1: 42}
                        var v85 int64 = 285
                        v4 = (v4 - 42)
                        fmt.Println(v85)
                        _ = v83
                        _ = v84
                        _ = v85
                    }
                    var v86 int16 = 42
                    _ = v86
                } else {
                    var v87 int16 = 1
                    var v88 int64 = -9223372036854775808
                    _ = v87
                    _ = v88
                }
                v89 := GenS14{F0: (4294967295 == 59), F1: false}
                _ = v89
            }
            var v90 float64 = 3.14159
            _ = fi15
            _ = v90
        }
        _ = fi13
        _ = v14
    }
    for fi91 := 0; fi91 < 20; fi91++ {
        if ("" != (("hello" + "\n\t") + "")) {
            v92 := []int{(fi91 << ((18446744073709551615) & 63)), (v2 | (v2 >> ((42) & 63))), v2}
            _ = v92
        }
        _ = fi91
    }
    fmt.Println(v2, v3, v2)
    v4 -= (^(v4))
    switch v2 {
    case 42:
        if true {
            v2 += v2
        }
    case 991:
        fmt.Println(v3, v2, v3)
    }
    var v93 int32 = 322
    for fi94 := 0; fi94 < 4; fi94++ {
        if true {
            v95 := GenS9{F0: 930}
            var v96 int32 = -2147483648
            if false {
                v97 := []int{fi94, (^(v2))}
                var v98 int32 = (v96 & -2147483648)
                if ((v2 >> ((0) & 63)) >= 128) {
                    fmt.Println(v3, v4, v4)
                    if false {
                        v96 = v98
                        v99 := GenS13{F0: 20, F1: 611}
                        var v100 float64 = ((-2.25 * 3.14159) * float64(v3))
                        _ = v100
                        _ = v99
                    }
                    var v101 uint = 2
                    _ = v101
                }
                _ = v97
                _ = v98
            }
            _ = v95
            _ = v96
        } else {
            v102 := []int{fi94}
            var v103 int64 = -9223372036854775808
            _ = v102
            _ = v103
        }
        if (128 > 1) {
            for fi104 := 0; fi104 < 18; fi104++ {
                v93 = (v93 | v93)
                _ = fi104
            }
            var v105 uint = 4294967295
            var v106 uint32 = 933
            _ = v105
            _ = v106
        } else {
            v3 = v3
        }
        switch 4294967295 {
        case 127:
            if true {
                v3 -= (v3 + (v3 / 1.0))
            }
        default:
            switch "\n\t" {
            case "hello":
                v1[769] = true
            case "foo bar":
                v107 := []int{}
                _ = v107
            default:
                v108 := []int{128}
                _ = v108
            }
        }
        _ = fi94
    }
    v3 -= -2.25
    for fi109 := 0; fi109 < 14; fi109++ {
        var v110 uint64 = v4
        v111 := v1[2147483648]
        switch "" {
        case "a":
            v110 += v4
        case "foo bar":
            v110 = v110
        }
        _ = fi109
        _ = v110
        _ = v111
    }
    v93 += (-(v93))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v93
}
