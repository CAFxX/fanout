package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 int64; F1 int }
type GenS5 struct { F0 int32; F1 uint; F2 int64 }
type GenS6 struct { F0 int64; F1 bool; F2 uint32 }
type GenS7 struct { F0 string; F1 uint32 }
type GenS8 struct { F0 int64; F1 string }
type GenS9 struct { F0 uint32; F1 string; F2 uint }
func genFn0(p0 int64, p1 uint8, p2 int32) uint8 {
    switch "foo bar" {
    case "a":
        var v4 int32 = p2
        _ = v4
    case "\n\t":
        for fi5 := 0; fi5 < 7; fi5++ {
            v6 := []int{(fi5 & (-(fi5))), fi5, fi5, (fi5 % ((fi5) | 1))}
            fmt.Println(fi5, p0, p1)
            v7 := GenS1{F0: "\n\t"}
            _ = fi5
            _ = v6
            _ = v7
        }
    case "":
        if true {
            var v8 uint16 = uint16(p1)
            if false {
                if (127 <= p1) {
                    if false {
                        if false {
                            v9 := GenS1{F0: (((("\n\t" + "a") + "foo bar") + "\n\t") + "\n\t")}
                            _ = v9
                        } else {
                            v10 := []int{2147483647, 127}
                            v10 = append(v10, 128)
                            _ = v10
                        }
                    }
                    switch "foo bar" {
                    case "":
                        v11 := GenS1{F0: "hello"}
                        _ = v11
                    default:
                        v8 += v8
                    }
                }
            }
            v12 := GenS1{F0: ""}
            _ = v12
            _ = v8
        } else {
            var v13 int = 9223372036854775807
            var v14 uint64 = 4294967295
            _ = v13
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 float64 = 1.5
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 18; fi17++ {
        var v18 bool = (947 < 1)
        _ = fi17
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    return p1
}

func main() {
    var v1 uint8 = 42
    var v2 int64 = -1
    var v3 uint16 = 42
    v19 := GenS2{F0: -53}
    fmt.Println(v3, v1, v3)
    if (!((0.0 < 0.1) || false)) {
        var v20 int8 = 61
        var v21 float64 = -2.25
        fmt.Println(v21, v1)
        _ = v20
        _ = v21
    }
    v2 -= (-(v2))
    v22 := GenS3{F0: (v3 & v3), F1: genFn0(v2, 1, int32(v2))}
    if false {
        var v23 uint8 = 1
        for fi24 := 0; fi24 < 14; fi24++ {
            if false {
                v25 := []int{int(v1), fi24, fi24}
                v23 -= genFn0(134, genFn0(v2, (v23 << ((1) & 63)), 42), -31)
                var v26 uint32 = 4294967295
                _ = v25
                _ = v26
            }
            var v27 int16 = 815
            if false {
                if (915 > v3) {
                    switch "hello" {
                    case "a":
                        if (2147483647 <= 1) {
                            v28 := GenS4{F0: v2, F1: -9223372036854775808}
                            if false {
                                var v29 float32 = (1e10 / 1.0)
                                _ = v29
                            } else {
                                fmt.Println(v27)
                                if (18446744073709551615 >= 537) {
                                    v23 -= v1
                                    fmt.Println(v1)
                                } else {
                                    switch "hello" {
                                    case "hello":
                                        v23 += 128
                                    }
                                    v27 -= -32768
                                    if (255 < 771) {
                                        v3 += v3
                                        var v30 uint16 = (v3 | 1)
                                        _ = v30
                                    } else {
                                        v27 = v27
                                        if false {
                                            switch (fi24 / (((^(fi24))) | 1)) {
                                            case 32768:
                                                switch "\n\t" {
                                                case "hello":
                                                    v31 := GenS1{F0: "a"}
                                                    _ = v31
                                                default:
                                                    v1 += uint8(v3)
                                                }
                                            }
                                            var v32 float32 = 0.0
                                            for fi33 := 0; fi33 < 7; fi33++ {
                                                v34 := GenS3{F0: (v3 >> ((2) & 63)), F1: 1}
                                                v35 := GenS5{F0: -2147483648, F1: 736, F2: 42}
                                                _ = fi33
                                                _ = v34
                                                _ = v35
                                            }
                                            _ = v32
                                        } else {
                                            v36 := GenS1{F0: (("\n\t" + "a") + (("a" + ("foo bar" + ("a" + "foo bar"))) + "\n\t"))}
                                            var v37 string = ("a" + ("foo bar" + "foo bar"))
                                            if true {
                                                v23 -= 0
                                                v38 := GenS6{F0: 42, F1: true, F2: 1}
                                                v39 := GenS1{F0: v37}
                                                _ = v38
                                                _ = v39
                                            }
                                            _ = v36
                                            _ = v37
                                        }
                                        v2 += -9223372036854775808
                                    }
                                }
                            }
                            _ = v28
                        } else {
                            switch "\n\t" {
                            case "a":
                                v1 = v1
                            default:
                                v40 := GenS6{F0: (v2 + (v2 - 42)), F1: true, F2: 57}
                                _ = v40
                            }
                            if ((fi24 ^ fi24) <= 0) {
                                fmt.Println(v1)
                                for fi41 := 0; fi41 < 20; fi41++ {
                                    fmt.Println(v2, v23, fi24)
                                    if true {
                                        v3 = v3
                                    } else {
                                        v42 := []int{}
                                        v27 += v27
                                        var v43 int64 = 42
                                        fmt.Println(fi41, fi24)
                                        _ = v42
                                        _ = v43
                                    }
                                    v44 := []int{}
                                    _ = fi41
                                    _ = v44
                                }
                                v45 := []int{}
                                _ = v45
                            } else {
                                v46 := []int{fi24, 32768}
                                if len(v46) > 0 {
                                    v46[0] = fi24
                                }
                                _ = v46
                            }
                        }
                    case "":
                        var v47 float32 = 1e10
                        _ = v47
                    case "foo bar":
                        v2 = (v2 << ((uint(v23)) & 63))
                    }
                    v48 := GenS4{F0: 0, F1: -9223372036854775808}
                    _ = v48
                } else {
                    v49 := []int{fi24, (fi24 | fi24)}
                    _ = v49
                }
                if true {
                    v27 = (v27 / ((-32768) | 1))
                } else {
                    for fi50 := 0; fi50 < 5; fi50++ {
                        v51 := GenS7{F0: "", F1: 1}
                        var v52 uint8 = (v23 * (v1 * v23))
                        _ = fi50
                        _ = v51
                        _ = v52
                    }
                    v53 := GenS7{F0: "\n\t", F1: 42}
                    v3 = 0
                    _ = v53
                }
                v2 += -193
            } else {
                v54 := []int{fi24, (fi24 - int(v23)), (fi24 >> ((255) & 63)), (fi24 >> ((809) & 63))}
                if (-1 == v27) {
                    if true {
                        v23 = v23
                    } else {
                        v23 = 1
                        v23 = uint8(v3)
                        var v55 int = fi24
                        _ = v55
                    }
                } else {
                    var v56 uint = 337
                    _ = v56
                }
                if false {
                    var v57 uint16 = (^(v3))
                    _ = v57
                }
                _ = v54
            }
            _ = fi24
            _ = v27
        }
        _ = v23
    } else {
        switch -741 {
        case 32768:
            var v58 int = 9223372036854775807
            _ = v58
        }
        v3 += 675
        v1 = (v1 / (((v1 / ((v1) | 1))) | 1))
    }
    v59 := GenS3{F0: 255, F1: (v1 | 2)}
    for fi60 := 0; fi60 < 5; fi60++ {
        if (v1 <= 1) {
            switch (fi60 + (fi60 / (((fi60 * fi60)) | 1))) {
            case 127:
                for fi61 := 0; fi61 < 16; fi61++ {
                    var v62 int16 = 701
                    var v63 uint = 255
                    var v64 uint16 = v3
                    _ = fi61
                    _ = v62
                    _ = v63
                    _ = v64
                }
            case -777:
                switch fi60 {
                case 32768:
                    var v65 bool = true
                    _ = v65
                }
            case 32767:
                v2 = (-(v2))
            default:
                v2 = (v2 | v2)
            }
            if false {
                v66 := []int{(fi60 & (fi60 / ((fi60) | 1))), (^(fi60)), (fi60 + (-(fi60)))}
                if false {
                    for fi67 := 0; fi67 < 4; fi67++ {
                        var v68 float32 = 1e10
                        var v69 uint64 = 717
                        switch (fi67 << ((18446744073709551615) & 63)) {
                        case 128:
                            var v70 uint = 0
                            _ = v70
                        case -145:
                            fmt.Println(fi60, v3)
                        default:
                            v69 -= (v69 ^ v69)
                        }
                        _ = fi67
                        _ = v68
                        _ = v69
                    }
                    if len(v66) > 0 {
                        v66[0] = fi60
                    }
                }
                fmt.Println(v2, fi60)
                _ = v66
            }
        }
        _ = fi60
    }
    fmt.Println(v1, v1)
    if true {
        if (true || true) {
            v3 = v3
        }
    } else {
        for fi71 := 0; fi71 < 5; fi71++ {
            var v72 uint = 589
            _ = fi71
            _ = v72
        }
    }
    if (((-2.25 - (-2.25 + float64(1.5))) + 1e10) == (3.14159 * float64(1.5))) {
        for fi73 := 0; fi73 < 13; fi73++ {
            var v74 int = fi73
            _ = fi73
            _ = v74
        }
        fmt.Println(v3)
    }
    for fi75 := 0; fi75 < 12; fi75++ {
        fmt.Println(fi75, v3, v2)
        for fi76 := 0; fi76 < 11; fi76++ {
            v77 := []int{0}
            fmt.Println(fi76, fi75, v1)
            _ = fi76
            _ = v77
        }
        if (-32768 >= 953) {
            if (false && false) {
                var v78 string = "a"
                _ = v78
            } else {
                for fi79 := 0; fi79 < 18; fi79++ {
                    fmt.Println(fi75)
                    _ = fi79
                }
                fmt.Println(v1, v2)
                var v80 int8 = 100
                _ = v80
            }
            v81 := GenS2{F0: 100}
            if (-813 >= v2) {
                for fi82 := 0; fi82 < 2; fi82++ {
                    switch (fi82 * 32767) {
                    case -1:
                        v83 := []int{(fi82 - fi75), 127, fi75}
                        _ = v83
                    case 618:
                        v84 := GenS4{F0: 9223372036854775807, F1: int(v3)}
                        _ = v84
                    case -610:
                        fmt.Println(v3, v1, v2)
                    default:
                        var v85 string = ""
                        _ = v85
                    }
                    var v86 bool = (!true)
                    var v87 uint64 = 382
                    var v88 int8 = 42
                    v89 := GenS5{F0: 42, F1: 730, F2: (v2 >> ((1) & 63))}
                    _ = fi82
                    _ = v86
                    _ = v87
                    _ = v88
                    _ = v89
                }
                if (((false && true) || false) || true) {
                    fmt.Println(v2)
                    var v90 float32 = 0.1
                    v1 -= (v1 % ((v1) | 1))
                    _ = v90
                } else {
                    if true {
                        switch ("a" + "") {
                        case "":
                            switch fi75 {
                            case 9223372036854775807:
                                for fi91 := 0; fi91 < 7; fi91++ {
                                    var v92 bool = (!(0.0 != ((1.5 + float32(1e10)) * float32(3.14159))))
                                    _ = fi91
                                    _ = v92
                                }
                            }
                        case "a":
                            v3 = (v3 / (((v3 % ((v3) | 1))) | 1))
                        case "hello":
                            var v93 int64 = v2
                            _ = v93
                        }
                    } else {
                        for fi94 := 0; fi94 < 14; fi94++ {
                            v95 := []int{(fi94 * fi94), 127, fi75, (fi75 / ((fi94) | 1))}
                            if true {
                                switch "a" {
                                case "a":
                                    v96 := GenS1{F0: "\n\t"}
                                    _ = v96
                                default:
                                    var v97 uint32 = 65535
                                    _ = v97
                                }
                                v95 = append(v95, (fi94 & fi75))
                                var v98 float64 = 1e10
                                _ = v98
                            }
                            var v99 uint32 = 2
                            switch "\n\t" {
                            case "a":
                                v100 := GenS4{F0: int64(fi75), F1: (-(fi94))}
                                _ = v100
                            }
                            var v101 uint8 = genFn0(42, (v1 << ((683) & 63)), 0)
                            fmt.Println(v1)
                            _ = fi94
                            _ = v101
                            _ = v95
                            _ = v99
                        }
                        switch -972 {
                        case -834:
                            if (642 < (^(v3))) {
                                fmt.Println(v3)
                            }
                        case -74:
                            switch "" {
                            case "\n\t":
                                v2 = (v2 + 0)
                            case "":
                                var v102 float64 = -2.25
                                _ = v102
                            default:
                                fmt.Println(v1, v3, fi75)
                            }
                        }
                    }
                    v103 := GenS5{F0: 42, F1: 4294967295, F2: (^(v2))}
                    _ = v103
                }
            } else {
                fmt.Println(v1, v2, fi75)
            }
            _ = v81
        } else {
            if false {
                var v104 int64 = int64(v3)
                if true {
                    fmt.Println(v104)
                } else {
                    fmt.Println(fi75, v104)
                    v105 := []int{fi75, (fi75 | fi75), 128}
                    v104 += (v2 - -649)
                    _ = v105
                }
                var v106 uint64 = 0
                _ = v104
                _ = v106
            }
            if true {
                var v107 int16 = 42
                _ = v107
            } else {
                fmt.Println(fi75, v2, fi75)
            }
            for fi108 := 0; fi108 < 18; fi108++ {
                if false {
                    var v109 float64 = -2.25
                    v3 += v3
                    _ = v109
                } else {
                    v110 := GenS8{F0: v2, F1: ""}
                    var v111 float32 = ((3.14159 * 1.5) + 1.5)
                    v112 := []int{(fi108 % ((fi75) | 1)), (fi75 >> ((255) & 63))}
                    if (false && true) {
                        if (false || true) {
                            switch "\n\t" {
                            case "foo bar":
                                var v113 string = (("foo bar" + "foo bar") + ("hello" + "a"))
                                _ = v113
                            default:
                                if false {
                                    v114 := GenS3{F0: (v3 & v3), F1: genFn0((^(v2)), 128, 42)}
                                    _ = v114
                                }
                            }
                        } else {
                            var v115 float32 = (0.1 * v111)
                            _ = v115
                        }
                        v2 -= v2
                        if len(v112) > 0 {
                            v112[0] = (fi108 / ((fi75) | 1))
                        }
                    } else {
                        var v116 int8 = -128
                        v1 = v1
                        _ = v116
                    }
                    if false {
                        switch "hello" {
                        case "foo bar":
                            v111 = (v111 * (float32(0.0) / v111))
                        case "a":
                            if (-381 < 42) {
                                v117 := GenS8{F0: v2, F1: ("\n\t" + "hello")}
                                fmt.Println(v111)
                                _ = v117
                            }
                        case "hello":
                            fmt.Println(v111, v3, fi108)
                        }
                        if len(v112) >= 2 {
                            v112 = v112[1:len(v112)]
                        }
                    }
                    var v118 int8 = 18
                    switch -2147483648 {
                    case 603:
                        fmt.Println(v2)
                        var v119 int16 = int16(v2)
                        _ = v119
                    case 9223372036854775807:
                        v3 += v3
                    }
                    _ = v110
                    _ = v111
                    _ = v112
                    _ = v118
                }
                var v120 uint32 = 2147483647
                var v121 uint16 = (v3 + (v3 - (^(v3))))
                _ = fi108
                _ = v120
                _ = v121
            }
        }
        _ = fi75
    }
    v1 = (^(v1))
    v122 := GenS9{F0: 2, F1: "foo bar", F2: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v122
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v59
}
