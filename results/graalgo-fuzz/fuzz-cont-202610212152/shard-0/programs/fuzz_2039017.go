package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int16; F1 int16; F2 uint16 }
type GenS5 struct { F0 bool; F1 bool; F2 string }
type GenS6 struct { F0 uint64; F1 int }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint; F1 uint64; F2 int16 }
type GenS9 struct { F0 int16; F1 int8 }
type GenS10 struct { F0 uint32; F1 int8; F2 bool }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 uint32; F1 uint64 }
func genFn0(p0 int32) uint16 {
    for fi6 := 0; fi6 < 8; fi6++ {
        p0 += p0
        fmt.Println(p0)
        _ = fi6
    }
    for fi7 := 0; fi7 < 12; fi7++ {
        p0 -= (p0 | p0)
        p0 -= (p0 & p0)
        _ = fi7
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        if (1 <= 4) {
            switch fi8 {
            case -499:
                var v9 int = fi8
                _ = v9
            case -451:
                switch (fi8 * fi8) {
                case 654:
                    switch "foo bar" {
                    case "":
                        if true {
                            var v10 uint64 = 0
                            if ((!false) && false) {
                                for fi11 := 0; fi11 < 17; fi11++ {
                                    var v12 int8 = -1
                                    _ = fi11
                                    _ = v12
                                }
                                if ((false && false) || (!(0.0 >= 1.5))) {
                                    var v13 uint8 = 128
                                    if (!(("\n\t" + "a") <= "hello")) {
                                        if true {
                                            var v14 int16 = -1
                                            _ = v14
                                        }
                                    } else {
                                        var v15 float32 = (0.1 / (0.1 + 1e10))
                                        for fi16 := 0; fi16 < 1; fi16++ {
                                            if ((v15 - 0.1) <= 3.14159) {
                                                var v17 int32 = p0
                                                v18 := []int{-104, (-(fi8))}
                                                _ = v17
                                                _ = v18
                                            }
                                            v19 := []int{0, fi16, (fi8 ^ fi16), fi8}
                                            if len(v19) > 0 {
                                                v19[0] = fi16
                                            }
                                            _ = fi16
                                            _ = v19
                                        }
                                        v15 = v15
                                        _ = v15
                                    }
                                    var v20 float64 = (3.14159 * -2.25)
                                    _ = v13
                                    _ = v20
                                }
                                for fi21 := 0; fi21 < 7; fi21++ {
                                    v22 := []int{fi8, fi8}
                                    var v23 uint32 = 707
                                    _ = fi21
                                    _ = v22
                                    _ = v23
                                }
                            } else {
                                var v24 float32 = (0.0 - 1.5)
                                var v25 float32 = v24
                                _ = v24
                                _ = v25
                            }
                            _ = v10
                        } else {
                            v26 := GenS1{F0: 0, F1: 222, F2: 2}
                            fmt.Println(fi8, fi8)
                            fmt.Println(p0, p0)
                            _ = v26
                        }
                    default:
                        if true {
                            var v27 uint8 = 128
                            var v28 int64 = -493
                            _ = v27
                            _ = v28
                        }
                    }
                }
            }
            if (false || true) {
                if false {
                    for fi29 := 0; fi29 < 13; fi29++ {
                        if (0 < 65535) {
                            fmt.Println(p0)
                        } else {
                            p0 = (p0 & 27)
                        }
                        _ = fi29
                    }
                }
            }
        } else {
            v30 := GenS2{F0: 0}
            if false {
                fmt.Println(p0)
                var v31 uint32 = 0
                _ = v31
            } else {
                p0 = 1
            }
            v32 := []int{fi8}
            _ = v30
            _ = v32
        }
        if true {
            if false {
                p0 += (p0 + p0)
                switch 0 {
                case 742:
                    var v33 string = ("" + ("" + ("foo bar" + "a")))
                    _ = v33
                case 635:
                    switch fi8 {
                    case 32767:
                        v34 := GenS2{F0: 18446744073709551615}
                        _ = v34
                    case 4294967295:
                        v35 := []int{}
                        _ = v35
                    }
                }
            }
        } else {
            var v36 uint64 = 1
            _ = v36
        }
        if false {
            var v37 int64 = 9223372036854775807
            var v38 float64 = 1e10
            p0 = 42
            _ = v37
            _ = v38
        } else {
            v39 := GenS1{F0: -333, F1: 0, F2: 2}
            _ = v39
        }
        _ = fi8
    }
    switch ("\n\t" + "\n\t") {
    case "hello":
        var v40 bool = (!("" == "foo bar"))
        _ = v40
    case "":
        p0 -= p0
    case "\n\t":
        v41 := GenS3{F0: p0}
        _ = v41
    default:
        v42 := GenS3{F0: p0}
        _ = v42
    }
    for fi43 := 0; fi43 < 5; fi43++ {
        if (721 < p0) {
            var v44 float64 = float64(((-2.25 - -2.25) / -2.25))
            var v45 int64 = 0
            var v46 int32 = 0
            _ = v44
            _ = v45
            _ = v46
        }
        p0 = p0
        v47 := GenS3{F0: p0}
        _ = fi43
        _ = v47
    }
    _ = p0
    return 32767
}

func main() {
    var v1 uint32 = 0
    v2 := make(map[uint32]uint32)
    var v3 bool = true
    var v4 uint16 = 255
    var v5 int64 = 9223372036854775807
    if v3 {
        v48 := GenS4{F0: 1, F1: 994, F2: (v4 + (v4 >> ((18446744073709551615) & 63)))}
        _ = v48
    }
    v49 := []int{4294967295}
    if v3 {
        if len(v49) > 0 {
            v49[0] = 32767
        }
        switch -303 {
        case -9223372036854775808:
            fmt.Println(v1, v3)
        case -1:
            if true {
                v5 = (^(v5))
                if false {
                    v3 = v3
                }
                for fi50 := 0; fi50 < 18; fi50++ {
                    if v3 {
                        for fi51 := 0; fi51 < 12; fi51++ {
                            v3 = (v3 && (false || v3))
                            if v3 {
                                if (!(0.1 <= -2.25)) {
                                    var v52 uint8 = 42
                                    _ = v52
                                }
                                var v53 uint16 = 803
                                _ = v53
                            }
                            _ = fi51
                        }
                        _ = len(v49)
                    }
                    delete(v2, (v1 & 4294967295))
                    v3 = (v3 || v3)
                    _ = fi50
                }
            } else {
                v54 := GenS5{F0: true, F1: (0.0 != (float64((1.5 * 3.14159)) + (float64((0.0 + (-2.25 - 1.5))) + (float64(0.1) * (0.0 * 0.0))))), F2: (("\n\t" + ("foo bar" + (("\n\t" + ("foo bar" + "\n\t")) + "hello"))) + (("" + "\n\t") + "foo bar"))}
                v5 -= v5
                _ = v54
            }
        }
        _ = len(v2)
    } else {
        v2[v1] = 65535
    }
    if false {
        v5 = v5
        fmt.Println(v4, v3, v5)
    } else {
        fmt.Println(v1)
    }
    if (v4 > 0) {
        fmt.Println(v5)
        delete(v2, 523)
    }
    v2[v1] = (v1 & (v1 ^ (v1 ^ (v1 % ((466) | 1)))))
    if v3 {
        if v3 {
            fmt.Println(v4, v1, v1)
        } else {
            v55 := GenS6{F0: 2, F1: -9223372036854775808}
            for fi56 := 0; fi56 < 18; fi56++ {
                if (32767 != (v4 + (v4 % ((v4) | 1)))) {
                    v57 := GenS7{F0: -1}
                    for fi58 := 0; fi58 < 10; fi58++ {
                        v5 -= v5
                        if (!true) {
                            var v59 int8 = -1
                            _ = v59
                        }
                        v3 = v3
                        _ = fi58
                    }
                    _ = v57
                }
                if (622 > v1) {
                    var v60 int32 = -357
                    _ = v60
                } else {
                    v61 := v2[2147483647]
                    if (0 == 0) {
                        if (-736 == 32767) {
                            switch int(v4) {
                            case 32767:
                                switch "" {
                                case "foo bar":
                                    var v62 uint8 = 234
                                    _ = v62
                                case "a":
                                    v1 = (v61 % (((v1 / ((v1) | 1))) | 1))
                                }
                            }
                            v61 += v1
                        } else {
                            for fi63 := 0; fi63 < 19; fi63++ {
                                v1 -= v1
                                _ = fi63
                            }
                            switch "foo bar" {
                            case "":
                                v2[808] = 4294967295
                            default:
                                for fi64 := 0; fi64 < 1; fi64++ {
                                    var v65 string = ("a" + "a")
                                    _ = fi64
                                    _ = v65
                                }
                            }
                            switch "hello" {
                            case "hello":
                                if v3 {
                                    for fi66 := 0; fi66 < 13; fi66++ {
                                        v67 := GenS8{F0: 4294967295, F1: 42, F2: 0}
                                        _ = fi66
                                        _ = v67
                                    }
                                    switch 1 {
                                    case 672:
                                        var v68 float32 = (0.0 + -2.25)
                                        _ = v68
                                    case -120:
                                        v69 := GenS6{F0: 2, F1: (^(fi56))}
                                        _ = v69
                                    default:
                                        var v70 bool = true
                                        _ = v70
                                    }
                                }
                            case "a":
                                _ = len(v2)
                            case "foo bar":
                                var v71 uint16 = 0
                                _ = v71
                            }
                        }
                        for fi72 := 0; fi72 < 8; fi72++ {
                            var v73 int32 = 0
                            var v74 uint64 = uint64(v61)
                            fmt.Println(v5, v1, fi56)
                            switch "hello" {
                            case "foo bar":
                                if v3 {
                                    switch "\n\t" {
                                    case "\n\t":
                                        var v75 string = "foo bar"
                                        if (true && (127 >= 128)) {
                                            v49 = append(v49, (fi72 ^ int(v74)))
                                            if ((v73 == (v73 * v73)) || ((v5 ^ (v5 ^ v5)) <= 42)) {
                                                switch 892 {
                                                case 21:
                                                    if true {
                                                        v1 = (v1 * v1)
                                                    }
                                                case -1:
                                                    v61 = (v61 / (((^(v61))) | 1))
                                                case 42:
                                                    var v76 uint32 = v1
                                                    _ = v76
                                                default:
                                                    var v77 float64 = ((3.14159 * 1e10) * float64(0.0))
                                                    _ = v77
                                                }
                                                var v78 int64 = 9223372036854775807
                                                _ = v78
                                            } else {
                                                var v79 uint = 255
                                                if v3 {
                                                    v1 = (v61 | uint32(v4))
                                                } else {
                                                    if v3 {
                                                        _ = len(v2)
                                                    } else {
                                                        var v80 int16 = 1000
                                                        if v3 {
                                                            v49 = append(v49, fi56)
                                                            fmt.Println(v73, v75)
                                                            if v3 {
                                                                v75 = v75
                                                            }
                                                        } else {
                                                            v81 := GenS1{F0: 42, F1: v61, F2: (v4 - 65535)}
                                                            v80 = -356
                                                            if (v75 == (((v75 + v75) + v75) + (v75 + v75))) {
                                                                if ((^(v61)) > (v1 * v1)) {
                                                                    v3 = v3
                                                                    var v82 uint = (v79 & 65535)
                                                                    _ = v82
                                                                }
                                                                v74 = (v74 & (v74 + v74))
                                                                v83 := v2[uint32(v4)]
                                                                _ = v83
                                                            } else {
                                                                v3 = v3
                                                                v84 := GenS2{F0: (v79 | (^(v79)))}
                                                                if v3 {
                                                                    if (170 < 128) {
                                                                        if v3 {
                                                                            if (v4 != v4) {
                                                                                var v85 uint8 = 255
                                                                                _ = v85
                                                                            } else {
                                                                                switch fi56 {
                                                                                case 32767:
                                                                                    var v86 uint16 = 0
                                                                                    _ = v86
                                                                                case 175:
                                                                                    switch "" {
                                                                                    case "":
                                                                                        v87 := GenS2{F0: (v79 / (((v79 | v79)) | 1))}
                                                                                        _ = v87
                                                                                    default:
                                                                                        v88 := GenS3{F0: 0}
                                                                                        _ = v88
                                                                                    }
                                                                                case 680:
                                                                                    v89 := GenS9{F0: v80, F1: int8(fi72)}
                                                                                    _ = v89
                                                                                }
                                                                                var v90 string = "foo bar"
                                                                                if v3 {
                                                                                    v49 = append(v49, (fi72 - (fi72 | fi72)))
                                                                                    v2[(v61 ^ (v61 ^ (v1 ^ uint32(v74))))] = 65535
                                                                                    var v91 uint = (v79 ^ v79)
                                                                                    if len(v49) > 0 {
                                                                                        v49[0] = int(v4)
                                                                                    }
                                                                                    _ = v91
                                                                                }
                                                                                _ = v90
                                                                            }
                                                                            if false {
                                                                                v92 := GenS7{F0: v73}
                                                                                v4 -= uint16(v79)
                                                                                var v93 int64 = (^(v5))
                                                                                _ = v92
                                                                                _ = v93
                                                                            }
                                                                            v74 = 0
                                                                        } else {
                                                                            var v94 float32 = (float32(1e10) * 1e10)
                                                                            var v95 uint32 = 639
                                                                            _ = v94
                                                                            _ = v95
                                                                        }
                                                                        var v96 int32 = int32(v4)
                                                                        _ = v96
                                                                    } else {
                                                                        if v3 {
                                                                            if ((v61 - (^(v1))) >= v1) {
                                                                                var v97 uint16 = genFn0((v73 % (((v73 ^ 2147483647)) | 1)))
                                                                                _ = v97
                                                                            }
                                                                        }
                                                                        if v3 {
                                                                            switch fi72 {
                                                                            case 594:
                                                                                if len(v49) > 0 {
                                                                                    v49[0] = (fi56 ^ fi56)
                                                                                }
                                                                            case -549:
                                                                                fmt.Println(v79, v61, fi72)
                                                                            default:
                                                                                fmt.Println(fi72)
                                                                            }
                                                                            v79 = v79
                                                                        }
                                                                        if true {
                                                                            v98 := GenS1{F0: v5, F1: (v61 % ((v1) | 1)), F2: (^(v4))}
                                                                            _ = v98
                                                                        } else {
                                                                            _ = len(v49)
                                                                        }
                                                                    }
                                                                    fmt.Println(v74)
                                                                }
                                                                var v99 int64 = (v5 * (v5 >> ((v79) & 63)))
                                                                _ = v84
                                                                _ = v99
                                                            }
                                                            _ = v81
                                                        }
                                                        _ = v80
                                                    }
                                                    v100 := GenS9{F0: 32767, F1: 1}
                                                    var v101 float32 = -2.25
                                                    if len(v49) >= 2 {
                                                        v49 = v49[1:len(v49)]
                                                    }
                                                    _ = len(v2)
                                                    _ = v100
                                                    _ = v101
                                                }
                                                v74 = 479
                                                _ = v79
                                            }
                                            _ = len(v49)
                                        } else {
                                            v102 := GenS6{F0: 2, F1: fi72}
                                            v73 -= (v73 % ((v73) | 1))
                                            _ = v102
                                        }
                                        _ = v75
                                    case "":
                                        v74 -= 0
                                    default:
                                        var v103 uint16 = (v4 >> ((858) & 63))
                                        _ = v103
                                    }
                                }
                            case "":
                                v104 := GenS5{F0: (v73 < -701), F1: v3, F2: "\n\t"}
                                _ = v104
                            default:
                                switch -9223372036854775808 {
                                case 128:
                                    v105 := GenS10{F0: uint32(v4), F1: int8(fi56), F2: v3}
                                    _ = v105
                                case -125:
                                    var v106 string = "foo bar"
                                    _ = v106
                                default:
                                    v49 = append(v49, fi72)
                                }
                            }
                            v107 := GenS11{F0: -681}
                            var v108 float32 = 1e10
                            _ = fi72
                            _ = v107
                            _ = v108
                            _ = v73
                            _ = v74
                        }
                    }
                    var v109 bool = v3
                    _ = v109
                    _ = v61
                }
                var v110 int16 = -952
                _ = fi56
                _ = v110
            }
            switch "\n\t" {
            case "\n\t":
                v1 += 0
            }
            _ = v55
        }
    }
    if true {
        for fi111 := 0; fi111 < 10; fi111++ {
            if (true && (-14 >= 0)) {
                v112 := GenS4{F0: 414, F1: -677, F2: (v4 % ((v4) | 1))}
                var v113 bool = v3
                _ = v112
                _ = v113
            } else {
                v4 -= (v4 - 255)
            }
            for fi114 := 0; fi114 < 1; fi114++ {
                delete(v2, (v1 / ((42) | 1)))
                v115 := GenS7{F0: -2147483648}
                var v116 int = fi111
                var v117 int = (fi114 & (v116 * v116))
                var v118 int64 = (v5 - v5)
                _ = fi114
                _ = v115
                _ = v116
                _ = v117
                _ = v118
            }
            v119 := GenS12{F0: (v1 ^ (v1 + (v1 % (((v1 << ((979) & 63))) | 1)))), F1: 4294967295}
            _ = fi111
            _ = v119
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v49:", len(v49), v49)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
}
