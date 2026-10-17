package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16; F1 bool }
type GenS5 struct { F0 bool; F1 int16 }
type GenS6 struct { F0 uint16; F1 int8 }
type GenS7 struct { F0 uint8; F1 uint; F2 uint64 }
type GenS8 struct { F0 int8; F1 string }
type GenS9 struct { F0 int8; F1 uint; F2 int64 }
func genFn0(p0 uint) int {
    v4 := []int{2147483647}
    p0 = (p0 << ((65535) & 63))
    var v5 int = -9223372036854775808
    if len(v4) > 0 {
        v4[0] = (v5 / (((v5 - (v5 & v5))) | 1))
    }
    _ = p0
    _ = v4
    _ = v5
    return 4294967295
}
func genFn1(p0 uint64, p1 uint16, p2 int32) int8 {
    v6 := GenS1{F0: true, F1: true}
    if (true || (p2 < (p2 % (((p2 * p2)) | 1)))) {
        p0 += p0
        switch genFn0(509) {
        case 32768:
            var v7 int16 = -269
            _ = v7
        case 32767:
            v8 := GenS1{F0: true, F1: false}
            _ = v8
        case 283:
            var v9 uint32 = 2
            _ = v9
        default:
            var v10 uint = 982
            _ = v10
        }
        var v11 float64 = 0.0
        _ = v11
    } else {
        var v12 int32 = p2
        fmt.Println(p0)
        _ = v12
    }
    v13 := GenS1{F0: false, F1: (p1 <= 446)}
    v14 := []int{114, 596, 127, -2147483648}
    v15 := GenS1{F0: true, F1: (!true)}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v15
    _ = v6
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{}
    var v3 int64 = -23
    if true {
        for fi16 := 0; fi16 < 8; fi16++ {
            var v17 uint32 = 935
            if ((!(!(!false))) || (false || false)) {
                if (1 == 0) {
                    v18 := GenS1{F0: true, F1: true}
                    switch -894 {
                    case 652:
                        v3 += v3
                    case 9223372036854775807:
                        if true {
                            for fi19 := 0; fi19 < 18; fi19++ {
                                if true {
                                    if ((float64(-2.25) * 1.5) != float64(-2.25)) {
                                        v17 += (v17 >> ((60) & 63))
                                        fmt.Println(v17)
                                    }
                                    v17 = v17
                                } else {
                                    v17 = v17
                                }
                                v20 := GenS1{F0: true, F1: false}
                                _ = fi19
                                _ = v20
                            }
                            var v21 int32 = -2147483648
                            _ = v21
                        }
                    default:
                        switch genFn0(811) {
                        case 32767:
                            v22 := GenS2{F0: 136, F1: ("a" + "\n\t")}
                            _ = v22
                        case 127:
                            for fi23 := 0; fi23 < 12; fi23++ {
                                v3 -= (-(v3))
                                var v24 bool = (-13 > -128)
                                _ = fi23
                                _ = v24
                            }
                        case 128:
                            var v25 int64 = 9223372036854775807
                            _ = v25
                        }
                    }
                    _ = v18
                }
                if false {
                    var v26 string = ""
                    _ = len(v2)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v26
                } else {
                    var v27 string = ("foo bar" + "hello")
                    delete(v1, ((((v27 + "a") + "a") + v27) + (v27 + "foo bar")))
                    _ = v27
                }
                fmt.Println(fi16, v17, fi16)
            }
            var v28 uint32 = v17
            _ = fi16
            _ = v17
            _ = v28
        }
    } else {
        v3 += v3
        for fi29 := 0; fi29 < 13; fi29++ {
            for fi30 := 0; fi30 < 1; fi30++ {
                var v31 uint = 328
                fmt.Println(fi29, fi29, fi29)
                v2 = append(v2, fi29)
                _ = fi30
                _ = v31
            }
            _ = fi29
        }
    }
    v32 := GenS1{F0: true, F1: ((false || false) || (false || ((-128 == 0) && true)))}
    if true {
        var v33 string = ("hello" + ("\n\t" + "foo bar"))
        fmt.Println(v3, v3)
        if (!(true || false)) {
            fmt.Println(v33, v33)
            var v34 int64 = v3
            var v35 string = ("" + ("" + "hello"))
            _ = v34
            _ = v35
        }
        _ = v33
    }
    if (42 <= 1) {
        var v36 uint8 = 1
        var v37 uint32 = 206
        _ = v36
        _ = v37
    }
    switch -1 {
    case -685:
        if true {
            fmt.Println(v3)
            v3 = 774
        } else {
            if (((899 == 4294967295) && (!true)) && (979 != 1)) {
                v3 = v3
                var v38 float64 = 1e10
                _ = v38
            }
            v39 := GenS2{F0: 1, F1: "a"}
            switch 0 {
            case 32768:
                for fi40 := 0; fi40 < 13; fi40++ {
                    v3 = -9223372036854775808
                    _ = fi40
                }
            case 13:
                v3 = (v3 - v3)
            case 42:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            _ = v39
        }
    case 2147483647:
        v3 = (v3 / (((v3 + (v3 << ((42) & 63)))) | 1))
    default:
        v41 := GenS1{F0: false, F1: (true || true)}
        _ = v41
    }
    v42 := GenS3{F0: (true || true)}
    var v43 bool = false
    var v44 uint16 = 42
    v45 := GenS4{F0: 470, F1: (1.5 >= float32(3.14159))}
    if (((^(v44)) != (v44 * v44)) || true) {
        fmt.Println(v44, v44, v44)
        v3 += (^(v3))
    } else {
        var v46 uint = 223
        switch 127 {
        case 42:
            fmt.Println(v44)
        case -1:
            fmt.Println(v44)
        case 32767:
            if v43 {
                if false {
                    for fi47 := 0; fi47 < 10; fi47++ {
                        v43 = (!(false && ((v44 * 2) == 121)))
                        var v48 bool = v43
                        _ = fi47
                        _ = v48
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                } else {
                    if v43 {
                        v3 += (v3 << ((v46) & 63))
                        for fi49 := 0; fi49 < 5; fi49++ {
                            v46 = 255
                            _ = fi49
                        }
                        var v50 int = -9223372036854775808
                        _ = v50
                    } else {
                        if (32767 >= (v44 | (v44 ^ (v44 | 478)))) {
                            fmt.Println(v3, v43)
                            v46 = v46
                            for fi51 := 0; fi51 < 9; fi51++ {
                                var v52 uint8 = 42
                                v53 := GenS3{F0: ((("a" + "hello") + ("foo bar" + "foo bar")) > "")}
                                _ = fi51
                                _ = v52
                                _ = v53
                            }
                        }
                    }
                }
                v54 := GenS5{F0: true, F1: -1}
                _ = v54
            }
        default:
            var v55 uint = 354
            _ = v55
        }
        for fi56 := 0; fi56 < 7; fi56++ {
            for fi57 := 0; fi57 < 1; fi57++ {
                v58 := GenS3{F0: v43}
                fmt.Println(fi57)
                _ = fi57
                _ = v58
            }
            _ = fi56
        }
        _ = v46
    }
    v59 := GenS6{F0: 2, F1: 62}
    fmt.Println(v43, v44, v44)
    if v43 {
        v44 += (v44 / ((v44) | 1))
    }
    if v43 {
        for fi60 := 0; fi60 < 10; fi60++ {
            switch (fi60 * (fi60 << ((268) & 63))) {
            case -2147483648:
                if (!false) {
                    v61 := GenS1{F0: ((v44 / ((v44) | 1)) < v44), F1: ((v3 * -9223372036854775808) > (-(v3)))}
                    _ = v61
                }
            case 127:
                if (1e10 >= (-2.25 / 1.5)) {
                    v62 := GenS7{F0: 42, F1: 1, F2: 1}
                    v63 := GenS6{F0: (v44 << ((2) & 63)), F1: -24}
                    _ = v62
                    _ = v63
                }
            case 317:
                if ((fi60 % ((fi60) | 1)) < (fi60 + fi60)) {
                    var v64 int64 = v3
                    _ = v64
                }
            default:
                v44 = (v44 + (^(v44)))
            }
            _ = fi60
        }
        fmt.Println(v3, v43)
        switch 4294967295 {
        case 4294967295:
            if true {
                var v65 int16 = 1
                v43 = v43
                for fi66 := 0; fi66 < 20; fi66++ {
                    v67 := GenS6{F0: (^(v44)), F1: -5}
                    if v43 {
                        if v43 {
                            v43 = (65535 != 202)
                        }
                        _ = len(v1)
                        if ((v3 | 42) == (v3 >> ((138) & 63))) {
                            v65 = v65
                            v2 = append(v2, (fi66 ^ 1))
                            for fi68 := 0; fi68 < 2; fi68++ {
                                var v69 int = (fi68 & fi68)
                                v44 = (^(v44))
                                fmt.Println(v65, v44, v65)
                                _ = fi68
                                _ = v69
                            }
                        }
                    }
                    v70 := GenS1{F0: true, F1: v43}
                    _ = fi66
                    _ = v67
                    _ = v70
                }
                _ = v65
            } else {
                if len(v2) > 0 {
                    v2[0] = 4294967295
                }
            }
        }
    } else {
        if v43 {
            var v71 uint = 475
            v72 := GenS2{F0: 325, F1: "a"}
            for fi73 := 0; fi73 < 13; fi73++ {
                var v74 uint16 = (^(v44))
                v44 -= v44
                switch genFn0(213) {
                case -960:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case 2147483647:
                    var v75 int32 = 597
                    _ = v75
                default:
                    v3 = (v3 - 73)
                }
                _ = fi73
                _ = v74
            }
            _ = v71
            _ = v72
        } else {
            v44 = 255
            switch ("\n\t" + ("\n\t" + (("\n\t" + "hello") + "foo bar"))) {
            case "foo bar":
                var v76 uint16 = v44
                _ = v76
            case "hello":
                fmt.Println(v43, v3)
            default:
                var v77 uint16 = 290
                _ = v77
            }
        }
        for fi78 := 0; fi78 < 14; fi78++ {
            fmt.Println(v3)
            v3 = (-(v3))
            fmt.Println(v44, fi78)
            _ = fi78
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    if v43 {
        var v79 float64 = 0.1
        switch "\n\t" {
        case "hello":
            if (genFn0(838) >= -652) {
                var v80 int = genFn0(534)
                if (128 != 42) {
                    for fi81 := 0; fi81 < 12; fi81++ {
                        for fi82 := 0; fi82 < 6; fi82++ {
                            switch "foo bar" {
                            case "":
                                if (v43 || v43) {
                                    var v83 int64 = 9223372036854775807
                                    if v43 {
                                        var v84 uint16 = 42
                                        var v85 uint32 = 2147483647
                                        _ = v84
                                        _ = v85
                                    }
                                    _ = v83
                                } else {
                                    v2 = append(v2, 127)
                                    v80 = (fi81 - (^(v80)))
                                }
                                if (4294967295 > 236) {
                                    v86 := v1[(("hello" + "foo bar") + "\n\t")]
                                    _ = v86
                                }
                                fmt.Println(fi81, v79)
                            case "hello":
                                var v87 uint = 255
                                _ = v87
                            case "\n\t":
                                v80 -= (fi81 >> ((42) & 63))
                                v80 += fi81
                                switch fi81 {
                                case 42:
                                    v88 := GenS7{F0: 127, F1: 65535, F2: 4294967295}
                                    _ = v88
                                case 127:
                                    if v43 {
                                        if len(v2) > 0 {
                                            v2[0] = 32767
                                        }
                                        if (!(v44 == (v44 * (^(v44))))) {
                                            var v89 uint16 = 65535
                                            _ = len(v2)
                                            _ = v89
                                        } else {
                                            v90 := GenS3{F0: v43}
                                            var v91 uint32 = 654
                                            if (v3 <= 9223372036854775807) {
                                                v43 = false
                                            }
                                            delete(v1, "hello")
                                            _ = v90
                                            _ = v91
                                        }
                                    }
                                case -2147483648:
                                    v44 = 203
                                }
                            default:
                                delete(v1, ("hello" + ("foo bar" + "hello")))
                                var v92 float32 = (0.0 * 0.0)
                                var v93 int16 = -1
                                if len(v2) > 0 {
                                    v2[0] = 127
                                }
                                var v94 uint = 65535
                                _ = v92
                                _ = v93
                                _ = v94
                            }
                            _ = fi82
                        }
                        var v95 uint64 = 601
                        var v96 uint64 = 964
                        _ = fi81
                        _ = v95
                        _ = v96
                    }
                }
                _ = v80
            } else {
                var v97 float64 = v79
                _ = v97
            }
        case "a":
            if false {
                switch "\n\t" {
                case "\n\t":
                    if len(v2) > 0 {
                        v2[0] = 42
                    }
                case "hello":
                    for fi98 := 0; fi98 < 4; fi98++ {
                        for fi99 := 0; fi99 < 15; fi99++ {
                            v100 := GenS3{F0: (0 > 1)}
                            var v101 int64 = 830
                            _ = fi99
                            _ = v100
                            _ = v101
                        }
                        v102 := GenS1{F0: true, F1: (295 > 1)}
                        if false {
                            if (fi98 >= (fi98 | 2147483647)) {
                                var v103 bool = true
                                for fi104 := 0; fi104 < 14; fi104++ {
                                    v2 = append(v2, (fi98 << ((947) & 63)))
                                    _ = len(v2)
                                    var v105 int32 = 2147483647
                                    _ = fi104
                                    _ = v105
                                }
                                var v106 int8 = -128
                                _ = v103
                                _ = v106
                            }
                            v107 := GenS8{F0: -23, F1: "foo bar"}
                            switch "a" {
                            case "":
                                var v108 uint16 = 32767
                                _ = v108
                            case "\n\t":
                                for fi109 := 0; fi109 < 4; fi109++ {
                                    if (v43 && v43) {
                                        switch (-(fi98)) {
                                        case 32767:
                                            v110 := GenS5{F0: (42 < -2147483648), F1: -1}
                                            _ = v110
                                        case -308:
                                            var v111 float32 = -2.25
                                            _ = v111
                                        case -2147483648:
                                            var v112 float32 = ((1.5 + -2.25) + -2.25)
                                            _ = v112
                                        default:
                                            fmt.Println(v43, v3, fi109)
                                        }
                                        var v113 string = ("foo bar" + (("hello" + ("foo bar" + "hello")) + "foo bar"))
                                        _ = v113
                                    } else {
                                        var v114 uint16 = v44
                                        var v115 int64 = (v3 | (v3 & -37))
                                        _ = v114
                                        _ = v115
                                    }
                                    _ = fi109
                                }
                            case "a":
                                v116 := GenS9{F0: 0, F1: 42, F2: (v3 + v3)}
                                _ = v116
                            }
                            _ = v107
                        } else {
                            var v117 uint8 = 2
                            _ = v117
                        }
                        _ = fi98
                        _ = v102
                    }
                case "foo bar":
                    if len(v2) > 0 {
                        v2[0] = genFn0(4294967295)
                    }
                default:
                    var v118 uint16 = (v44 + 42)
                    _ = v118
                }
                v119 := v1["a"]
                _ = v119
            }
        }
        _ = v79
    } else {
        var v120 float32 = 1.5
        for fi121 := 0; fi121 < 2; fi121++ {
            v122 := GenS7{F0: 240, F1: 255, F2: 48}
            v44 -= 32767
            v3 = v3
            _ = fi121
            _ = v122
        }
        _ = v120
    }
    switch -1 {
    case 725:
        var v123 int64 = 42
        _ = v123
    case 9223372036854775807:
        v3 += v3
    case 994:
        fmt.Println(v44)
    default:
        var v124 uint = 65535
        _ = v124
    }
    v1[""] = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v59
}
