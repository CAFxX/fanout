package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 int32; F1 uint16; F2 string }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32; F1 int; F2 int }
type GenS6 struct { F0 int64; F1 uint }
type GenS7 struct { F0 uint; F1 uint16 }
type GenS8 struct { F0 string; F1 int16; F2 uint32 }
type GenS9 struct { F0 uint64; F1 uint8; F2 int64 }
type GenS10 struct { F0 int; F1 uint16 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 string; F1 int32 }
type GenS14 struct { F0 int16; F1 uint8 }
type GenS15 struct { F0 int32; F1 uint; F2 int }
type GenS16 struct { F0 bool; F1 int64; F2 uint32 }
func genFn0(p0 bool, p1 uint8, p2 uint16) int8 {
    var v7 int = 2147483647
    if ((3.14159 / 1.5) < ((float64(-2.25) - 1.5) / ((((((0.1 * 0.1) * 1.5) + float64(1.5)) + (1e10 - (1e10 * float64(float32(1.5))))) + (-2.25 - float64(-2.25))) * -2.25))) {
        var v8 int64 = 9223372036854775807
        _ = v8
    } else {
        v9 := GenS1{F0: (^(p2)), F1: 42}
        _ = v9
    }
    p1 += 255
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 62
}
func genFn1(p0 uint32, p1 uint64) int64 {
    var v10 string = "a"
    for fi11 := 0; fi11 < 3; fi11++ {
        v12 := GenS2{F0: v10, F1: 2}
        _ = fi11
        _ = v12
    }
    v13 := []int{}
    for fi14 := 0; fi14 < 13; fi14++ {
        switch v10 {
        case "":
            fmt.Println(v10, v10)
        default:
            if true {
                v15 := GenS1{F0: 32767, F1: 0}
                _ = v15
            } else {
                for fi16 := 0; fi16 < 8; fi16++ {
                    switch ("foo bar" + ("\n\t" + "hello")) {
                    case "foo bar":
                        v17 := GenS3{F0: 42, F1: 42, F2: v10}
                        _ = v17
                    case "a":
                        switch 718 {
                        case 42:
                            v18 := GenS3{F0: 1, F1: 0, F2: v10}
                            var v19 float32 = (-2.25 / ((1.5 - float32(3.14159)) - 0.0))
                            var v20 int = (fi16 >> ((2) & 63))
                            _ = v18
                            _ = v19
                            _ = v20
                        case 32767:
                            var v21 uint8 = 20
                            _ = v21
                        case -9223372036854775808:
                            var v22 float32 = 3.14159
                            _ = v22
                        }
                    case "hello":
                        var v23 float32 = 0.0
                        _ = v23
                    default:
                        v24 := GenS3{F0: -466, F1: 516, F2: v10}
                        _ = v24
                    }
                    fmt.Println(fi16)
                    if true {
                        _ = len(v13)
                        v25 := GenS2{F0: "\n\t", F1: 2}
                        if ((true && true) && true) {
                            switch fi14 {
                            case 128:
                                v26 := GenS4{F0: (3.14159 > 1e10)}
                                _ = v26
                            default:
                                if len(v13) >= 2 {
                                    v13 = v13[1:len(v13)]
                                }
                                v13 = append(v13, (fi14 ^ (-(fi16))))
                                v27 := GenS1{F0: 65535, F1: 1000}
                                _ = v27
                            }
                        }
                        p0 = p0
                        _ = v25
                    }
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                    var v28 uint = 18446744073709551615
                    _ = fi16
                    _ = v28
                }
            }
        }
        for fi29 := 0; fi29 < 17; fi29++ {
            fmt.Println(p1, v10, fi14)
            p1 = 528
            _ = fi29
        }
        _ = fi14
    }
    for fi30 := 0; fi30 < 10; fi30++ {
        var v31 uint32 = (p0 - p0)
        fmt.Println(v10, fi30)
        _ = fi30
        _ = v31
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v13
    return -1
}

func main() {
    var v1 int64 = 42
    var v2 int8 = -1
    v3 := make(map[uint32]int)
    var v4 uint8 = 42
    var v5 int16 = -32768
    var v6 uint64 = 541
    var v32 int64 = (v1 ^ (v1 % ((v1) | 1)))
    if false {
        for fi33 := 0; fi33 < 15; fi33++ {
            fmt.Println(v4, fi33)
            v4 -= 255
            _ = fi33
        }
        var v34 bool = ((42 > 1) && (-78 != genFn0((!(!false)), (v4 * (v4 + 2)), 0)))
        for fi35 := 0; fi35 < 13; fi35++ {
            if (v5 != v5) {
                fmt.Println(v34, v4)
                v36 := []int{}
                _ = v36
            } else {
                v37 := GenS1{F0: 65535, F1: (v5 >> ((65535) & 63))}
                _ = v37
            }
            _ = fi35
        }
        _ = v34
    }
    if (32767 <= -1) {
        v6 += (^(v6))
        for fi38 := 0; fi38 < 18; fi38++ {
            fmt.Println(v5)
            v5 = v5
            v6 -= v6
            _ = fi38
        }
        switch "\n\t" {
        case "a":
            var v39 bool = false
            _ = v39
        case "":
            v5 += v5
        default:
            if (1e10 > 1e10) {
                switch -2147483648 {
                case -538:
                    v40 := []int{0}
                    _ = v40
                case 127:
                    var v41 bool = true
                    _ = v41
                }
                for fi42 := 0; fi42 < 17; fi42++ {
                    var v43 int8 = (v2 - v2)
                    for fi44 := 0; fi44 < 14; fi44++ {
                        var v45 int16 = 451
                        v46 := GenS5{F0: 376, F1: (fi44 | fi42), F2: (fi42 | -9223372036854775808)}
                        var v47 int32 = 1
                        _ = fi44
                        _ = v45
                        _ = v46
                        _ = v47
                    }
                    var v48 string = "a"
                    _ = fi42
                    _ = v43
                    _ = v48
                }
                v6 += 18446744073709551615
            } else {
                if false {
                    v32 += v1
                } else {
                    v49 := GenS6{F0: v32, F1: 1}
                    var v50 int32 = 1
                    switch 483 {
                    case 0:
                        if (0 < 0) {
                            v32 += -140
                            switch ("" + (("hello" + "a") + "foo bar")) {
                            case "":
                                v50 -= v50
                            default:
                                switch "\n\t" {
                                case "a":
                                    var v51 int64 = v1
                                    _ = v51
                                case "\n\t":
                                    v52 := GenS1{F0: 2, F1: v5}
                                    _ = v52
                                case "hello":
                                    for fi53 := 0; fi53 < 19; fi53++ {
                                        if (v2 == (v2 + -128)) {
                                            fmt.Println(v2, v50, v50)
                                            switch "\n\t" {
                                            case "\n\t":
                                                v54 := GenS7{F0: 4294967295, F1: 32767}
                                                _ = v54
                                            case "foo bar":
                                                var v55 int16 = (v5 ^ v5)
                                                _ = v55
                                            default:
                                                switch (fi53 ^ fi53) {
                                                case 863:
                                                    var v56 uint = 425
                                                    _ = v56
                                                case -2147483648:
                                                    v32 = v1
                                                default:
                                                    v57 := []int{127, fi53, (fi53 / ((fi53) | 1)), fi53}
                                                    _ = v57
                                                }
                                            }
                                        } else {
                                            v58 := GenS6{F0: 0, F1: 1}
                                            _ = v58
                                        }
                                        if false {
                                            var v59 bool = (((0.1 / (3.14159 - 1e10)) * 1.5) > (1.5 * 0.1))
                                            for fi60 := 0; fi60 < 14; fi60++ {
                                                var v61 float64 = 1.5
                                                if ((v61 / (-2.25 * (v61 - 0.0))) >= (3.14159 / 0.1)) {
                                                    v62 := GenS8{F0: "a", F1: v5, F2: 2}
                                                    v3[930] = 32767
                                                    _ = v62
                                                } else {
                                                    switch 0 {
                                                    case -1:
                                                        var v63 float64 = 0.1
                                                        _ = v63
                                                    case 771:
                                                        fmt.Println(v4, v1, fi60)
                                                    default:
                                                        v64 := []int{(fi53 | (fi53 | -451)), 9223372036854775807, fi53}
                                                        _ = v64
                                                    }
                                                    v65 := []int{fi53, -927, fi60}
                                                    if len(v65) > 0 {
                                                        v65[0] = (fi53 << ((294) & 63))
                                                    }
                                                    _ = v65
                                                }
                                                _ = fi60
                                                _ = v61
                                            }
                                            _ = v59
                                        }
                                        v1 = 42
                                        _ = fi53
                                    }
                                }
                            }
                        }
                    case 128:
                        var v66 uint8 = 1
                        _ = v66
                    case -900:
                        switch -9223372036854775808 {
                        case -103:
                            fmt.Println(v4)
                        }
                    }
                    _ = v49
                    _ = v50
                }
                v67 := GenS8{F0: "hello", F1: v5, F2: 4294967295}
                _ = v67
            }
        }
    } else {
        if ((!false) || (!false)) {
            v68 := GenS9{F0: (v6 >> ((4294967295) & 63)), F1: 42, F2: genFn1(0, 42)}
            v69 := GenS10{F0: 32768, F1: 65535}
            _ = v68
            _ = v69
        }
        v32 = v1
        var v70 uint16 = 2
        _ = v70
    }
    switch "foo bar" {
    case "hello":
        var v71 uint32 = 1
        _ = v71
    }
    for fi72 := 0; fi72 < 15; fi72++ {
        v73 := GenS6{F0: (v1 << ((255) & 63)), F1: 1}
        switch (fi72 / (((fi72 | -9223372036854775808)) | 1)) {
        case 2147483647:
            switch "hello" {
            case "foo bar":
                v74 := []int{42, fi72, -2147483648, (fi72 >> ((734) & 63))}
                _ = v74
            case "":
                var v75 uint16 = 161
                _ = v75
            default:
                v1 = 0
            }
        default:
            v76 := GenS11{F0: 2}
            _ = v76
        }
        _ = fi72
        _ = v73
    }
    var v77 int8 = (v2 << ((0) & 63))
    var v78 float32 = (((1.5 * (1e10 * (float32(3.14159) - 1.5))) / 0.1) + 3.14159)
    var v79 int8 = (v77 % ((-128) | 1))
    v79 = genFn0((-2147483648 <= 2147483647), v4, 255)
    fmt.Println(v77, v5, v2)
    v80 := GenS11{F0: 65535}
    if (!(uint32(v6) >= 2)) {
        if ((v6 << ((70) & 63)) != v6) {
            v81 := GenS12{F0: "foo bar"}
            _ = v81
        } else {
            for fi82 := 0; fi82 < 19; fi82++ {
                for fi83 := 0; fi83 < 10; fi83++ {
                    switch (fi82 >> ((65535) & 63)) {
                    case 9223372036854775807:
                        delete(v3, 42)
                    default:
                        v84 := []int{fi83}
                        v85 := GenS9{F0: 1, F1: (v4 >> ((2) & 63)), F2: (v32 + v1)}
                        _ = v84
                        _ = v85
                    }
                    _ = fi83
                }
                var v86 uint64 = (^(v6))
                _ = fi82
                _ = v86
            }
            v87 := GenS1{F0: 507, F1: 667}
            v88 := v3[2147483648]
            _ = v87
            _ = v88
        }
    }
    fmt.Println(v78, v79, v32)
    if (!(((v78 * -2.25) + v78) >= v78)) {
        for fi89 := 0; fi89 < 18; fi89++ {
            fmt.Println(v5)
            _ = fi89
        }
    }
    switch 916 {
    case 283:
        var v90 int = -374
        _ = v90
    case 128:
        v91 := GenS13{F0: "", F1: -285}
        _ = v91
    }
    if true {
        if (3.14159 != float64(v78)) {
            v32 = (v1 + 1)
            var v92 uint = uint(v6)
            if ((0.1 + 0.1) >= float64((float32((1e10 * -2.25)) + 1.5))) {
                if (!false) {
                    v79 -= 42
                    if false {
                        switch 32768 {
                        case -649:
                            v32 -= (v32 >> ((0) & 63))
                        case 127:
                            var v93 float64 = -2.25
                            _ = v93
                        }
                    }
                }
                v94 := GenS3{F0: 1, F1: 19, F2: "foo bar"}
                var v95 uint = (^(v92))
                _ = v94
                _ = v95
            } else {
                v96 := v3[2147483648]
                _ = v96
            }
            _ = v92
        } else {
            if (0 != (v4 & v4)) {
                for fi97 := 0; fi97 < 12; fi97++ {
                    v2 -= v77
                    if (false && true) {
                        for fi98 := 0; fi98 < 7; fi98++ {
                            switch fi98 {
                            case -2147483648:
                                var v99 uint64 = (v6 % ((v6) | 1))
                                _ = v99
                            case 0:
                                var v100 uint64 = (v6 | (v6 / (((v6 % ((42) | 1))) | 1)))
                                _ = v100
                            case 4294967295:
                                if (false && ((v6 ^ v6) == (v6 % ((v6) | 1)))) {
                                    v5 -= (v5 << ((107) & 63))
                                } else {
                                    v101 := GenS14{F0: (v5 - (v5 * (v5 | v5))), F1: v4}
                                    _ = v101
                                }
                            default:
                                v102 := GenS14{F0: v5, F1: (v4 + v4)}
                                _ = v102
                            }
                            switch -876 {
                            case 9223372036854775807:
                                var v103 int64 = v32
                                _ = v103
                            case 4294967295:
                                v5 -= v5
                            default:
                                switch "" {
                                case "a":
                                    var v104 uint32 = 2147483648
                                    _ = v104
                                case "hello":
                                    v32 += v1
                                }
                            }
                            _ = fi98
                        }
                        v105 := GenS15{F0: -1, F1: 2, F2: fi97}
                        v106 := GenS16{F0: true, F1: (v32 << ((42) & 63)), F2: 2147483647}
                        _ = v105
                        _ = v106
                    } else {
                        v32 = (v1 - v1)
                        v107 := GenS5{F0: 1, F1: 842, F2: 0}
                        _ = v107
                    }
                    _ = fi97
                }
                for fi108 := 0; fi108 < 14; fi108++ {
                    v78 -= (v78 / (v78 + (v78 * (0.0 - v78))))
                    var v109 float64 = -2.25
                    _ = fi108
                    _ = v109
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v77
    _ = v78
    _ = v79
    _ = v80
}
