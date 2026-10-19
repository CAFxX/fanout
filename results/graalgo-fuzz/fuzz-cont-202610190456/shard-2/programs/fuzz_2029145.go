package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 string }
type GenS5 struct { F0 int; F1 int16; F2 int }
type GenS6 struct { F0 uint16; F1 string; F2 uint16 }
type GenS7 struct { F0 uint32; F1 int }
type GenS8 struct { F0 uint32; F1 int16 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int64; F1 int }
type GenS11 struct { F0 int16; F1 uint8 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 uint; F1 int32 }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 int64; F1 int8; F2 uint32 }
type GenS16 struct { F0 int64; F1 int; F2 int64 }
type GenS17 struct { F0 int16; F1 int64; F2 int16 }
type GenS18 struct { F0 string }
func genFn0() uint64 {
    if (-2147483648 >= 433) {
        v5 := GenS1{F0: -1, F1: "hello"}
        var v6 int16 = 42
        var v7 uint64 = 18446744073709551615
        _ = v5
        _ = v6
        _ = v7
    }
    if (1 == 1) {
        if false {
            if (-2147483648 <= -2147483648) {
                var v8 int64 = 42
                v9 := []int{127, 1, 446}
                _ = v8
                _ = v9
            }
            switch "a" {
            case "hello":
                v10 := GenS1{F0: 32767, F1: "foo bar"}
                _ = v10
            case "a":
                v11 := []int{128, -2147483648, 4294967295, 32768}
                _ = v11
            default:
                if (!(!true)) {
                    var v12 int32 = 1
                    if true {
                        var v13 int32 = (v12 - v12)
                        _ = v13
                    } else {
                        v14 := GenS2{F0: 0}
                        fmt.Println(v12)
                        var v15 uint16 = 1
                        _ = v14
                        _ = v15
                    }
                    _ = v12
                } else {
                    var v16 float32 = (1.5 * (1.5 + -2.25))
                    _ = v16
                }
            }
        } else {
            var v17 uint8 = 128
            v17 = v17
            var v18 int16 = -1
            _ = v17
            _ = v18
        }
        v19 := []int{-498}
        if (163 > 179) {
            var v20 float64 = 0.1
            v19 = append(v19, -9223372036854775808)
            _ = v20
        }
        _ = v19
    }
    var v21 uint32 = 2147483647
    v22 := []int{4294967295, 128}
    _ = v21
    _ = v22
    return 0
}
func genFn1() uint16 {
    for fi23 := 0; fi23 < 9; fi23++ {
        for fi24 := 0; fi24 < 3; fi24++ {
            var v25 float64 = -2.25
            _ = fi24
            _ = v25
        }
        _ = fi23
    }
    var v26 bool = (!(0.1 != ((((1e10 - 0.0) * 0.1) + 1e10) * -2.25)))
    v27 := GenS1{F0: -1, F1: "\n\t"}
    v26 = v26
    _ = v26
    _ = v27
    return 42
}

func main() {
    v1 := []uint32{}
    var v2 uint8 = 163
    var v3 uint64 = 18446744073709551615
    v4 := make(map[uint32]bool)
    var v28 uint8 = (^(v2))
    v1 = append(v1, 2147483647)
    v3 += v3
    for fi29 := 0; fi29 < 20; fi29++ {
        v30 := GenS3{F0: false}
        _ = fi29
        _ = v30
    }
    if len(v1) > 0 {
        v1[0] = uint32(v2)
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if true {
        for fi31 := 0; fi31 < 13; fi31++ {
            v32 := GenS1{F0: -1, F1: "foo bar"}
            if true {
                v2 = 128
                v3 = v3
            }
            _ = fi31
            _ = v32
        }
        var v33 int8 = 100
        v3 = v3
        _ = v33
    }
    var v34 int8 = 0
    if true {
        for fi35 := 0; fi35 < 5; fi35++ {
            var v36 string = "hello"
            for fi37 := 0; fi37 < 12; fi37++ {
                var v38 uint = 255
                _ = fi37
                _ = v38
            }
            if (((!(255 < 255)) && true) || ((-(fi35)) == 1)) {
                v2 -= 252
                for fi39 := 0; fi39 < 13; fi39++ {
                    if true {
                        v40 := GenS4{F0: v3, F1: (v36 + ((("a" + "") + "hello") + v36))}
                        _ = v40
                    } else {
                        if ((1.5 * float32(0.1)) < 1.5) {
                            delete(v4, 1)
                            if true {
                                var v41 int = 32768
                                _ = v41
                            } else {
                                var v42 int8 = 127
                                var v43 float32 = (1e10 * ((((3.14159 + float32(float64(0.1))) * -2.25) - float32(3.14159)) + (float32(1.5) / 3.14159)))
                                var v44 uint16 = 78
                                _ = v42
                                _ = v43
                                _ = v44
                            }
                            var v45 int64 = 9223372036854775807
                            v28 -= (v2 / ((128) | 1))
                            _ = v45
                        }
                    }
                    fmt.Println(v28)
                    v36 = "a"
                    v46 := GenS1{F0: int16(v28), F1: (v36 + "a")}
                    var v47 int8 = v34
                    _ = fi39
                    _ = v46
                    _ = v47
                }
                v2 -= (v28 & 2)
            }
            _ = fi35
            _ = v36
        }
        if false {
            for fi48 := 0; fi48 < 1; fi48++ {
                for fi49 := 0; fi49 < 15; fi49++ {
                    if ((0.0 - -2.25) >= 3.14159) {
                        if false {
                            var v50 float64 = 1e10
                            if true {
                                v3 -= (^(v3))
                                if (((v34 & 1) != (-(v34))) && (true || true)) {
                                    v28 += (v2 / ((uint8(v3)) | 1))
                                    v51 := GenS5{F0: (-(fi48)), F1: int16(v2), F2: (fi49 % ((fi49) | 1))}
                                    v2 += (v2 / ((v28) | 1))
                                    _ = v51
                                } else {
                                    switch fi49 {
                                    case 42:
                                        v50 = v50
                                    }
                                    v2 = v28
                                    v28 = 254
                                }
                            } else {
                                v52 := GenS1{F0: 1000, F1: ""}
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = v52
                            }
                            _ = v50
                        } else {
                            v53 := v4[457]
                            v4[4294967295] = v53
                            if false {
                                var v54 float64 = (0.1 + (float64(3.14159) + (1e10 + float64(3.14159))))
                                _ = v54
                            } else {
                                v53 = false
                                if v53 {
                                    if false {
                                        var v55 int = (-(fi48))
                                        switch "" {
                                        case "a":
                                            fmt.Println(v2, v55)
                                        default:
                                            v34 = v34
                                            var v56 int = fi48
                                            var v57 float64 = (1e10 / (-2.25 - 3.14159))
                                            _ = v56
                                            _ = v57
                                        }
                                        if v53 {
                                            if v53 {
                                                v2 = (v2 | (v2 | (^(v2))))
                                                var v58 uint16 = 49
                                                if v53 {
                                                    v55 += 356
                                                    var v59 int = (v55 | 42)
                                                    var v60 int8 = v34
                                                    fmt.Println(v53, v2, v28)
                                                    _ = v59
                                                    _ = v60
                                                } else {
                                                    v2 = (v2 >> ((65535) & 63))
                                                    fmt.Println(v2)
                                                    fmt.Println(v34, v28)
                                                }
                                                _ = v58
                                            }
                                        } else {
                                            v61 := GenS4{F0: (v3 * 42), F1: ("foo bar" + (("\n\t" + "foo bar") + "\n\t"))}
                                            fmt.Println(fi48, fi49, v2)
                                            _ = v61
                                        }
                                        v2 += v28
                                        v28 -= 2
                                        _ = v55
                                    } else {
                                        fmt.Println(v53)
                                        v3 -= (v3 << ((2) & 63))
                                    }
                                    v3 = (v3 % (((v3 + genFn0())) | 1))
                                    var v62 bool = v53
                                    _ = v62
                                }
                                if ("" == "a") {
                                    var v63 string = "foo bar"
                                    v28 = (^(v28))
                                    if true {
                                        if (-9223372036854775808 <= 42) {
                                            v2 = 0
                                            _ = len(v1)
                                            v53 = (42 < 0)
                                            v3 += genFn0()
                                            fmt.Println(v34, v63)
                                            fmt.Println(fi48, v3, fi48)
                                        }
                                        fmt.Println(fi49, v28)
                                    } else {
                                        v3 += v3
                                        v28 = 255
                                        v1 = append(v1, 65535)
                                    }
                                    _ = v63
                                } else {
                                    v64 := GenS6{F0: 65535, F1: "\n\t", F2: 65535}
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = v64
                                }
                            }
                            _ = v53
                        }
                    }
                    switch -593 {
                    case 544:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case 32768:
                        fmt.Println(fi48, v3)
                    case 2147483647:
                        v65 := GenS7{F0: 0, F1: 32767}
                        _ = v65
                    default:
                        if true {
                            if (-332 < 386) {
                                if false {
                                    v66 := v4[970]
                                    if len(v1) > 0 {
                                        v1[0] = 618
                                    }
                                    _ = v66
                                }
                                var v67 int32 = -1
                                if true {
                                    if (false || false) {
                                        fmt.Println(v2, fi48, v28)
                                        if true {
                                            switch "foo bar" {
                                            case "\n\t":
                                                fmt.Println(fi49)
                                            case "foo bar":
                                                v1 = append(v1, 2147483647)
                                                var v68 float32 = 1.5
                                                _ = v68
                                            case "hello":
                                                v4[0] = false
                                            default:
                                                if false {
                                                    if false {
                                                        fmt.Println(v28, v28, v67)
                                                    } else {
                                                        var v69 int8 = -49
                                                        fmt.Println(v3, v34, v2)
                                                        v70 := v4[1]
                                                        _ = v69
                                                        _ = v70
                                                    }
                                                } else {
                                                    fmt.Println(v67, fi49)
                                                    if (-2.25 != 0.0) {
                                                        switch ("foo bar" + "foo bar") {
                                                        case "foo bar":
                                                            fmt.Println(fi48, v67, fi48)
                                                        default:
                                                            v71 := GenS4{F0: v3, F1: ((("hello" + ("" + ("a" + ""))) + "foo bar") + "\n\t")}
                                                            _ = v71
                                                        }
                                                        var v72 uint16 = 2
                                                        _ = v72
                                                    }
                                                    fmt.Println(v67, v2)
                                                    v3 = (^(v3))
                                                }
                                            }
                                        }
                                        var v73 int32 = v67
                                        _ = v73
                                    } else {
                                        fmt.Println(fi48, fi49)
                                    }
                                    v34 = (^(v34))
                                    v74 := GenS8{F0: 2, F1: 601}
                                    _ = v74
                                } else {
                                    v2 += (v28 ^ v2)
                                    if true {
                                        if ((v34 | v34) < (v34 & 1)) {
                                            var v75 int16 = 42
                                            var v76 uint8 = 0
                                            v77 := GenS3{F0: ((v67 <= (v67 + int32(v34))) || false)}
                                            _ = v75
                                            _ = v76
                                            _ = v77
                                        } else {
                                            var v78 int = fi48
                                            _ = v78
                                        }
                                        v1 = append(v1, 65535)
                                        v2 = 0
                                        if ((fi49 <= int(v34)) && false) {
                                            if (2 != 2) {
                                                v79 := GenS5{F0: fi48, F1: 649, F2: (fi48 % ((32768) | 1))}
                                                if true {
                                                    v80 := GenS4{F0: v3, F1: "hello"}
                                                    var v81 uint32 = 2147483648
                                                    if true {
                                                        v4[(v81 - 0)] = ((3.14159 != (1.5 / 3.14159)) || (v67 <= v67))
                                                        if false {
                                                            switch (fi48 - (-(fi48))) {
                                                            case -9223372036854775808:
                                                                var v82 int = fi48
                                                                _ = v82
                                                            case -2147483648:
                                                                v83 := GenS4{F0: v3, F1: "foo bar"}
                                                                _ = v83
                                                            default:
                                                                switch (fi49 | (fi48 | (fi49 * 32767))) {
                                                                case 42:
                                                                    v28 = 127
                                                                }
                                                            }
                                                            _ = len(v4)
                                                            if (("\n\t" + ((("\n\t" + "a") + "\n\t") + "hello")) != (("a" + ("a" + "")) + "\n\t")) {
                                                                var v84 bool = true
                                                                _ = v84
                                                            } else {
                                                                switch (fi48 - 9223372036854775807) {
                                                                case 1:
                                                                    var v85 uint64 = v3
                                                                    _ = v85
                                                                case -9223372036854775808:
                                                                    v86 := GenS3{F0: (false || false)}
                                                                    _ = v86
                                                                case -168:
                                                                    v4[v81] = (v67 > (v67 ^ (v67 >> ((781) & 63))))
                                                                }
                                                                var v87 uint64 = v3
                                                                v28 = v2
                                                                _ = v87
                                                            }
                                                        }
                                                        var v88 uint32 = v81
                                                        if (32767 > 32767) {
                                                            var v89 float32 = 0.1
                                                            v3 = 2
                                                            _ = v89
                                                        } else {
                                                            fmt.Println(v67, v88, v28)
                                                        }
                                                        _ = v88
                                                    }
                                                    _ = v80
                                                    _ = v81
                                                }
                                                var v90 bool = true
                                                _ = v79
                                                _ = v90
                                            }
                                            v34 += v34
                                            var v91 uint16 = genFn1()
                                            _ = v91
                                        } else {
                                            v92 := v4[73]
                                            var v93 int8 = (v34 << ((42) & 63))
                                            _ = v92
                                            _ = v93
                                        }
                                        if false {
                                            v67 = 448
                                        } else {
                                            v94 := GenS9{F0: true}
                                            fmt.Println(v3, v28)
                                            delete(v4, 626)
                                            _ = v94
                                        }
                                    } else {
                                        if (!false) {
                                            _ = len(v1)
                                            v3 -= 2
                                            if true {
                                                fmt.Println(v34, v28)
                                                if (128 <= 243) {
                                                    v95 := GenS10{F0: 1, F1: -2147483648}
                                                    switch 32767 {
                                                    case 9223372036854775807:
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            v67 -= v67
                                                        }
                                                    case 186:
                                                        var v96 int64 = -726
                                                        switch 127 {
                                                        case 9223372036854775807:
                                                            if true {
                                                                v2 = (v28 + (v28 << ((255) & 63)))
                                                            } else {
                                                                v2 += 128
                                                                v34 = (v34 % ((v34) | 1))
                                                                v28 += 0
                                                            }
                                                        case -9223372036854775808:
                                                            var v97 uint16 = 740
                                                            _ = v97
                                                        case -782:
                                                            var v98 uint64 = (v3 / ((v3) | 1))
                                                            _ = v98
                                                        }
                                                        _ = v96
                                                    case -523:
                                                        v2 += 128
                                                    default:
                                                        if true {
                                                            var v99 string = ("a" + ("\n\t" + (("\n\t" + (("hello" + "a") + "\n\t")) + "hello")))
                                                            _ = len(v4)
                                                            v100 := GenS11{F0: 1000, F1: v2}
                                                            var v101 uint64 = v3
                                                            fmt.Println(v99, fi49, v67)
                                                            _ = v100
                                                            _ = v101
                                                            _ = v99
                                                        }
                                                    }
                                                    v67 = (v67 >> ((uint(v3)) & 63))
                                                    var v102 uint64 = v3
                                                    v3 = (^(v3))
                                                    delete(v4, 2)
                                                    _ = v102
                                                    _ = v95
                                                }
                                                v4[979] = true
                                            }
                                        }
                                        switch ("" + (("\n\t" + "") + (("\n\t" + "") + "\n\t"))) {
                                        case "\n\t":
                                            var v103 uint32 = 860
                                            _ = v103
                                        case "foo bar":
                                            if ((^(v34)) > 104) {
                                                var v104 int = (fi49 >> ((65535) & 63))
                                                if true {
                                                    v105 := GenS12{F0: (v3 / ((genFn0()) | 1))}
                                                    v106 := GenS7{F0: 362, F1: (fi48 & (v104 / (((v104 % (((-(fi48))) | 1))) | 1)))}
                                                    var v107 uint = 255
                                                    _ = v105
                                                    _ = v106
                                                    _ = v107
                                                } else {
                                                    var v108 uint64 = v3
                                                    var v109 uint16 = 0
                                                    _ = v108
                                                    _ = v109
                                                }
                                                switch 42 {
                                                case -153:
                                                    if false {
                                                        var v110 int16 = 1000
                                                        _ = len(v1)
                                                        if ("hello" >= "foo bar") {
                                                            if false {
                                                                var v111 int8 = 127
                                                                var v112 bool = ((fi49 < v104) || true)
                                                                fmt.Println(v28, v110, v104)
                                                                fmt.Println(v3, v3)
                                                                _ = v111
                                                                _ = v112
                                                            } else {
                                                                if (42 == 0) {
                                                                    v110 -= -164
                                                                    v113 := GenS9{F0: ((v67 / ((42) | 1)) > (v67 << ((527) & 63)))}
                                                                    if false {
                                                                        var v114 uint8 = 255
                                                                        switch v104 {
                                                                        case 776:
                                                                            v115 := GenS13{F0: 65535, F1: (v67 / (((v67 - (-(v67)))) | 1))}
                                                                            _ = v115
                                                                        case 0:
                                                                            var v116 int8 = v34
                                                                            _ = v116
                                                                        default:
                                                                            var v117 uint = 0
                                                                            switch "\n\t" {
                                                                            case "hello":
                                                                                if ((v3 << ((107) & 63)) != v3) {
                                                                                    var v118 float32 = -2.25
                                                                                    _ = len(v1)
                                                                                    _ = v118
                                                                                }
                                                                            default:
                                                                                var v119 uint8 = 2
                                                                                _ = v119
                                                                            }
                                                                            var v120 bool = false
                                                                            _ = v117
                                                                            _ = v120
                                                                        }
                                                                        _ = v114
                                                                    } else {
                                                                        var v121 uint16 = 32767
                                                                        v122 := GenS5{F0: fi48, F1: v110, F2: (-(v104))}
                                                                        var v123 int = fi48
                                                                        _ = v121
                                                                        _ = v122
                                                                        _ = v123
                                                                    }
                                                                    _ = v113
                                                                }
                                                                if false {
                                                                    var v124 bool = (!true)
                                                                    v104 = v104
                                                                    if (!v124) {
                                                                        if len(v1) > 0 {
                                                                            v1[0] = 650
                                                                        }
                                                                        if v124 {
                                                                            var v125 int = (v104 + 127)
                                                                            _ = v125
                                                                        } else {
                                                                            if (32767 != (v104 % (((v104 - 42)) | 1))) {
                                                                                var v126 uint16 = genFn1()
                                                                                _ = v126
                                                                            } else {
                                                                                var v127 float64 = 1e10
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                                _ = v127
                                                                            }
                                                                            v2 = (^(v2))
                                                                            v2 -= v2
                                                                        }
                                                                        fmt.Println(v104, v2, v28)
                                                                        var v128 uint8 = 1
                                                                        _ = v128
                                                                    }
                                                                    _ = v124
                                                                } else {
                                                                    switch "\n\t" {
                                                                    case "":
                                                                        if true {
                                                                            switch "foo bar" {
                                                                            case "a":
                                                                                v3 = 770
                                                                            case "\n\t":
                                                                                var v129 bool = false
                                                                                _ = v129
                                                                            default:
                                                                                v130 := GenS14{F0: 128}
                                                                                fmt.Println(v28, v3, fi49)
                                                                                _ = v130
                                                                            }
                                                                        } else {
                                                                            if false {
                                                                                if (414 < 2) {
                                                                                    v131 := GenS5{F0: 770, F1: v110, F2: 178}
                                                                                    var v132 uint8 = 49
                                                                                    if true {
                                                                                        var v133 uint64 = (v3 * (v3 << ((1) & 63)))
                                                                                        _ = v133
                                                                                    } else {
                                                                                        if false {
                                                                                            v34 += (v34 >> ((18446744073709551615) & 63))
                                                                                            var v134 uint16 = 567
                                                                                            _ = v134
                                                                                        }
                                                                                    }
                                                                                    _ = v131
                                                                                    _ = v132
                                                                                }
                                                                            }
                                                                            v104 -= int(v34)
                                                                            v67 -= v67
                                                                        }
                                                                        fmt.Println(fi49, v2, v28)
                                                                        v135 := GenS9{F0: false}
                                                                        _ = v135
                                                                    case "\n\t":
                                                                        if true {
                                                                            v136 := GenS6{F0: 255, F1: "a", F2: genFn1()}
                                                                            fmt.Println(fi48, v104, v67)
                                                                            if false {
                                                                                if true {
                                                                                    fmt.Println(fi49, v67, v3)
                                                                                    v137 := GenS15{F0: -751, F1: -1, F2: 518}
                                                                                    _ = v137
                                                                                }
                                                                                v138 := GenS7{F0: 56, F1: fi49}
                                                                                _ = v138
                                                                            }
                                                                            _ = v136
                                                                        }
                                                                    }
                                                                    var v139 uint = uint(v3)
                                                                    _ = v139
                                                                }
                                                                v140 := GenS16{F0: 136, F1: v104, F2: -9223372036854775808}
                                                                _ = v140
                                                            }
                                                            var v141 int16 = -1
                                                            v1 = append(v1, 0)
                                                            _ = v141
                                                        }
                                                        _ = v110
                                                    } else {
                                                        var v142 uint32 = 689
                                                        switch -555 {
                                                        case -1:
                                                            v1 = append(v1, (v142 + 201))
                                                        case -2147483648:
                                                            var v143 int32 = (v67 ^ 984)
                                                            _ = v143
                                                        case 0:
                                                            v144 := GenS17{F0: -1, F1: -9223372036854775808, F2: 1}
                                                            _ = v144
                                                        }
                                                        switch v104 {
                                                        case -882:
                                                            v28 -= (v28 + v2)
                                                        case 176:
                                                            v3 = (v3 >> ((4294967295) & 63))
                                                        default:
                                                            switch 4294967295 {
                                                            case 0:
                                                                var v145 int64 = -997
                                                                _ = v145
                                                            case 1:
                                                                if len(v1) > 0 {
                                                                    v1[0] = (^(v142))
                                                                }
                                                            }
                                                        }
                                                        _ = v142
                                                    }
                                                case 292:
                                                    switch -902 {
                                                    case 127:
                                                        v3 += (v3 / ((42) | 1))
                                                    case 31:
                                                        v2 = v2
                                                    case 34:
                                                        var v146 string = ((("" + (("hello" + ("" + ((("a" + "") + "foo bar") + ((("\n\t" + "a") + "a") + "foo bar")))) + "hello")) + "hello") + "\n\t")
                                                        _ = v146
                                                    default:
                                                        var v147 int = fi49
                                                        _ = v147
                                                    }
                                                case 4294967295:
                                                    v148 := v4[4294967295]
                                                    _ = v148
                                                }
                                                _ = v104
                                            } else {
                                                v149 := GenS1{F0: -1, F1: "\n\t"}
                                                v34 = v34
                                                var v150 uint8 = 2
                                                v151 := GenS8{F0: 2147483647, F1: 1000}
                                                var v152 int8 = v34
                                                var v153 int16 = 32767
                                                var v154 int = (fi49 | fi48)
                                                v155 := GenS18{F0: "a"}
                                                v154 += (fi48 * (fi48 & fi49))
                                                _ = v149
                                                _ = v150
                                                _ = v151
                                                _ = v152
                                                _ = v153
                                                _ = v154
                                                _ = v155
                                            }
                                        case "a":
                                            v1 = append(v1, 2)
                                        }
                                        v2 = 120
                                    }
                                }
                                _ = v67
                            }
                        } else {
                            var v156 uint64 = genFn0()
                            _ = v156
                        }
                    }
                    _ = fi49
                }
                v157 := GenS16{F0: 1, F1: -95, F2: 9223372036854775807}
                _ = fi48
                _ = v157
            }
            switch "" {
            case "hello":
                v34 += v34
            case "\n\t":
                var v158 int = -2147483648
                _ = v158
            }
        }
    }
    v34 += 28
    var v159 int64 = -516
    v2 = v28
    v159 += 9223372036854775807
    v160 := v4[0]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v159:", v159)
    fmt.Println("v160:", v160)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v159
    _ = v160
    _ = v2
    _ = v28
    _ = v3
    _ = v34
    _ = v4
}
