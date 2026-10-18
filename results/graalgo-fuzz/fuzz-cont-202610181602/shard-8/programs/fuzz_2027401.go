package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string; F1 int8; F2 int8 }
type GenS4 struct { F0 string; F1 string; F2 int64 }
type GenS5 struct { F0 uint; F1 int16 }
type GenS6 struct { F0 uint8; F1 string; F2 int32 }
type GenS7 struct { F0 int64; F1 uint8 }
type GenS8 struct { F0 int64; F1 uint64; F2 int8 }
type GenS9 struct { F0 uint32; F1 int32; F2 int64 }
type GenS10 struct { F0 uint8; F1 string; F2 uint32 }
type GenS11 struct { F0 int16; F1 uint64; F2 uint8 }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 uint64; F1 uint16 }
type GenS14 struct { F0 int16; F1 uint8 }
type GenS15 struct { F0 int16; F1 int8 }
func genFn0(p0 int, p1 int, p2 uint) int64 {
    fmt.Println(p0, p1)
    if true {
        p0 += p0
        for fi4 := 0; fi4 < 3; fi4++ {
            v5 := []int{}
            for fi6 := 0; fi6 < 11; fi6++ {
                p1 -= (fi4 + p1)
                var v7 uint = 18446744073709551615
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi6
                _ = v7
            }
            var v8 uint = 18446744073709551615
            _ = fi4
            _ = v5
            _ = v8
        }
        v9 := GenS1{F0: 9223372036854775807, F1: p1}
        _ = v9
    } else {
        v10 := GenS2{F0: 42}
        p2 += p2
        var v11 uint16 = 675
        _ = v10
        _ = v11
    }
    var v12 uint16 = 32767
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 int64 = 9223372036854775807
        if (!(!(42 <= 0))) {
            var v15 int = fi13
            fmt.Println(p1, v12, fi13)
            v16 := []int{(-(fi13)), 128, 32768}
            _ = v15
            _ = v16
        }
        v17 := []int{(p1 | -307), p0, fi13}
        _ = fi13
        _ = v14
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 1
}
func genFn1(p0 uint32, p1 uint, p2 string) int8 {
    var v18 uint = (^(p1))
    if ((p0 & p0) == (p0 >> (((^(v18))) & 63))) {
        v19 := GenS3{F0: ("a" + p2), F1: 0, F2: 127}
        fmt.Println(p2)
        _ = v19
    } else {
        v18 = (^(v18))
    }
    fmt.Println(p2, p2)
    p0 = p0
    if (!false) {
        fmt.Println(p0)
        switch "a" {
        case "hello":
            for fi20 := 0; fi20 < 10; fi20++ {
                if true {
                    var v21 uint = v18
                    p0 = p0
                    _ = v21
                } else {
                    var v22 int32 = 0
                    _ = v22
                }
                switch (^(fi20)) {
                case 900:
                    var v23 int8 = 0
                    _ = v23
                case -1:
                    fmt.Println(p1, p0)
                case 458:
                    v24 := GenS1{F0: -1, F1: (fi20 & fi20)}
                    _ = v24
                }
                _ = fi20
            }
        case "foo bar":
            for fi25 := 0; fi25 < 11; fi25++ {
                for fi26 := 0; fi26 < 11; fi26++ {
                    v27 := GenS4{F0: p2, F1: p2, F2: genFn0((fi25 | fi26), (fi25 & (fi26 ^ (fi26 | (fi25 * (fi26 << (((v18 >> ((p1) & 63))) & 63)))))), (v18 | (v18 + p1)))}
                    var v28 float32 = 1e10
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                var v29 uint = v18
                _ = fi25
                _ = v29
            }
        case "a":
            v18 -= p1
        default:
            var v30 bool = false
            _ = v30
        }
    } else {
        for fi31 := 0; fi31 < 17; fi31++ {
            var v32 bool = (true || (-257 >= 42))
            var v33 int = fi31
            _ = fi31
            _ = v32
            _ = v33
        }
        for fi34 := 0; fi34 < 7; fi34++ {
            if (46 >= 127) {
                for fi35 := 0; fi35 < 17; fi35++ {
                    if true {
                        if true {
                            if (507 >= 42) {
                                var v36 float32 = 1e10
                                _ = v36
                            } else {
                                v37 := []int{32768, (fi35 - (fi34 - (-(fi34)))), (fi35 | (fi35 | -9223372036854775808)), fi35}
                                var v38 int = (fi35 & (fi35 + (fi35 ^ 32767)))
                                _ = len(v37)
                                _ = v37
                                _ = v38
                            }
                        }
                        p2 = p2
                        p2 = "foo bar"
                    }
                    v39 := GenS5{F0: (p1 >> ((v18) & 63)), F1: 32767}
                    if true {
                        fmt.Println(fi34)
                    }
                    switch 9223372036854775807 {
                    case 32768:
                        fmt.Println(fi35)
                        p1 += v18
                        if false {
                            v40 := []int{(fi35 * fi34), fi34}
                            v41 := GenS5{F0: 18446744073709551615, F1: -1}
                            _ = v40
                            _ = v41
                        } else {
                            p2 = "hello"
                            v42 := GenS1{F0: 42, F1: 749}
                            v43 := GenS6{F0: 148, F1: p2, F2: -953}
                            _ = v42
                            _ = v43
                        }
                    case 2147483647:
                        p1 -= v18
                    case 42:
                        switch (fi34 >> ((65535) & 63)) {
                        case -1:
                            v44 := []int{(fi35 / ((fi35) | 1)), (fi34 * (fi35 / ((fi34) | 1))), (fi34 & fi34), (fi35 / ((fi35) | 1))}
                            _ = v44
                        case 128:
                            if (!false) {
                                if ((p0 / ((p0) | 1)) <= 0) {
                                    var v45 bool = (1e10 > (0.0 + float64(1.5)))
                                    _ = v45
                                }
                                if true {
                                    v46 := []int{(fi34 / (((fi34 - fi34)) | 1)), (fi34 % ((759) | 1))}
                                    switch "\n\t" {
                                    case "\n\t":
                                        v47 := GenS7{F0: 9223372036854775807, F1: 42}
                                        _ = v47
                                    }
                                    _ = v46
                                } else {
                                    var v48 int = 702
                                    v49 := GenS3{F0: "foo bar", F1: -128, F2: 100}
                                    _ = v48
                                    _ = v49
                                }
                            } else {
                                v50 := []int{fi34, fi35, 32768}
                                v51 := GenS1{F0: 9223372036854775807, F1: fi35}
                                var v52 bool = false
                                _ = v50
                                _ = v51
                                _ = v52
                            }
                        default:
                            v53 := GenS5{F0: (^(v18)), F1: 1}
                            _ = v53
                        }
                    }
                    _ = fi35
                    _ = v39
                }
                switch -9223372036854775808 {
                case -1:
                    var v54 int32 = -1
                    _ = v54
                case 9223372036854775807:
                    v55 := GenS1{F0: 42, F1: fi34}
                    _ = v55
                default:
                    v56 := GenS8{F0: 580, F1: 501, F2: -128}
                    _ = v56
                }
                for fi57 := 0; fi57 < 12; fi57++ {
                    var v58 string = ((p2 + "foo bar") + p2)
                    if (!false) {
                        if true {
                            v59 := []int{}
                            if len(v59) >= 2 {
                                v59 = v59[1:len(v59)]
                            }
                            _ = v59
                        }
                        v60 := GenS6{F0: 0, F1: v58, F2: 183}
                        _ = v60
                    } else {
                        var v61 int32 = -1
                        p0 = (^(p0))
                        switch p2 {
                        case "":
                            v62 := []int{-9223372036854775808}
                            _ = v62
                        case "\n\t":
                            switch ((((p2 + p2) + ((("hello" + (((("a" + v58) + "a") + v58) + "\n\t")) + p2) + (v58 + "\n\t"))) + "foo bar") + ("\n\t" + p2)) {
                            case "":
                                v63 := GenS8{F0: -1, F1: 0, F2: 1}
                                _ = v63
                            case "hello":
                                var v64 int8 = 0
                                _ = v64
                            case "a":
                                if (!false) {
                                    var v65 string = p2
                                    var v66 int8 = 0
                                    _ = v65
                                    _ = v66
                                } else {
                                    var v67 uint64 = 2
                                    _ = v67
                                }
                            default:
                                if (false || (0.1 == float32(1.5))) {
                                    v68 := []int{(fi34 + (^(fi34))), fi34}
                                    switch "foo bar" {
                                    case "\n\t":
                                        switch 687 {
                                        case 436:
                                            if ((!false) && (int8(p1) > 127)) {
                                                var v69 string = v58
                                                _ = v69
                                            } else {
                                                if false {
                                                    var v70 int8 = 0
                                                    if true {
                                                        if true {
                                                            var v71 int16 = 42
                                                            if ((true || (p2 != ((v58 + v58) + "foo bar"))) && true) {
                                                                fmt.Println(p2)
                                                            }
                                                            var v72 float64 = 1e10
                                                            var v73 int = 2147483647
                                                            switch v73 {
                                                            case 128:
                                                                var v74 int32 = (-(v61))
                                                                _ = v74
                                                            }
                                                            v72 += ((v72 + (v72 / 1e10)) / 0.1)
                                                            _ = v71
                                                            _ = v72
                                                            _ = v73
                                                        } else {
                                                            p2 = "foo bar"
                                                            var v75 uint16 = 255
                                                            _ = v75
                                                        }
                                                        v18 += v18
                                                        v68 = append(v68, int(p1))
                                                        v76 := GenS4{F0: ("a" + v58), F1: ("hello" + p2), F2: int64(v61)}
                                                        fmt.Println(p0)
                                                        _ = len(v68)
                                                        var v77 uint8 = 87
                                                        _ = v76
                                                        _ = v77
                                                    } else {
                                                        v78 := GenS3{F0: (v58 + "foo bar"), F1: -128, F2: v70}
                                                        if len(v68) > 0 {
                                                            v68[0] = (fi57 - (fi57 | (fi34 % (((fi57 | (fi57 % (((fi34 * fi57)) | 1)))) | 1))))
                                                        }
                                                        _ = v78
                                                    }
                                                    var v79 int8 = (v70 << ((1) & 63))
                                                    var v80 int32 = -1
                                                    v18 = v18
                                                    _ = v70
                                                    _ = v79
                                                    _ = v80
                                                }
                                                if len(v68) > 0 {
                                                    v68[0] = (fi57 / ((fi57) | 1))
                                                }
                                            }
                                        case 444:
                                            var v81 bool = (42 != 6)
                                            _ = v81
                                        case 128:
                                            v68 = append(v68, fi57)
                                        default:
                                            var v82 int16 = 32767
                                            fmt.Println(fi57, fi57, p1)
                                            _ = v82
                                        }
                                    case "foo bar":
                                        var v83 uint = p1
                                        _ = v83
                                    case "a":
                                        switch ("hello" + (v58 + "\n\t")) {
                                        case "foo bar":
                                            var v84 uint = 242
                                            _ = v84
                                        case "":
                                            v68 = append(v68, (fi57 | int(v61)))
                                        case "a":
                                            v18 += (^(p1))
                                        }
                                    default:
                                        if true {
                                            _ = len(v68)
                                            var v85 int = fi57
                                            var v86 uint64 = 2
                                            _ = v85
                                            _ = v86
                                        }
                                    }
                                    _ = v68
                                } else {
                                    var v87 uint = p1
                                    v88 := []int{127, fi57, (fi34 * (fi34 >> (((v18 + (v18 & 65535))) & 63)))}
                                    _ = v87
                                    _ = v88
                                }
                            }
                        case "foo bar":
                            var v89 int64 = 0
                            _ = v89
                        }
                        v90 := []int{}
                        v61 += 2147483647
                        if false {
                            p0 += (p0 << (((^(p1))) & 63))
                            fmt.Println(fi57, v18, p2)
                            if len(v90) > 0 {
                                v90[0] = (fi57 / (((fi34 + 32768)) | 1))
                            }
                        }
                        switch (fi34 % ((fi57) | 1)) {
                        case -471:
                            switch (v58 + "foo bar") {
                            case "a":
                                p2 = p2
                            case "hello":
                                var v91 float64 = (3.14159 + (-2.25 * 0.0))
                                _ = v91
                            case "foo bar":
                                fmt.Println(p1)
                            }
                        case -1:
                            if ((!((uint16(v18) > 162) || true)) || false) {
                                switch v58 {
                                case "foo bar":
                                    switch (p2 + p2) {
                                    case "\n\t":
                                        var v92 uint32 = (p0 + p0)
                                        _ = v92
                                    case "":
                                        switch "foo bar" {
                                        case "foo bar":
                                            v93 := GenS6{F0: 127, F1: (p2 + "a"), F2: v61}
                                            _ = v93
                                        case "a":
                                            var v94 uint16 = 65535
                                            v95 := GenS3{F0: "foo bar", F1: 77, F2: 50}
                                            _ = v94
                                            _ = v95
                                        }
                                    default:
                                        var v96 uint16 = 333
                                        v97 := GenS9{F0: 919, F1: -287, F2: -9223372036854775808}
                                        v90 = append(v90, 32767)
                                        _ = v96
                                        _ = v97
                                    }
                                case "\n\t":
                                    if true {
                                        if len(v90) >= 2 {
                                            v90 = v90[1:len(v90)]
                                        }
                                        fmt.Println(p0, v18, p0)
                                    }
                                case "a":
                                    if true {
                                        p2 = v58
                                        var v98 int8 = -1
                                        _ = v98
                                    } else {
                                        _ = len(v90)
                                        switch v58 {
                                        case "\n\t":
                                            var v99 string = (v58 + "foo bar")
                                            _ = v99
                                        }
                                    }
                                default:
                                    if len(v90) > 0 {
                                        v90[0] = (fi34 << ((1) & 63))
                                    }
                                }
                                p0 += p0
                            } else {
                                if len(v90) >= 2 {
                                    v90 = v90[1:len(v90)]
                                }
                            }
                        default:
                            fmt.Println(v61, fi57)
                        }
                        _ = v61
                        _ = v90
                    }
                    v100 := GenS10{F0: 255, F1: v58, F2: 0}
                    _ = fi57
                    _ = v100
                    _ = v58
                }
            }
            for fi101 := 0; fi101 < 5; fi101++ {
                v102 := GenS11{F0: -32768, F1: 18446744073709551615, F2: 255}
                v103 := GenS3{F0: "foo bar", F1: 1, F2: 1}
                if (true && (p1 < p1)) {
                    switch 308 {
                    case -1:
                        fmt.Println(fi34)
                    case 32767:
                        var v104 uint32 = p0
                        _ = v104
                    case -9223372036854775808:
                        switch (p2 + (p2 + p2)) {
                        case "\n\t":
                            p2 = p2
                        case "a":
                            fmt.Println(fi34)
                        }
                    default:
                        v18 = (p1 % ((1) | 1))
                    }
                    var v105 bool = (!(0 == 42))
                    _ = v105
                }
                _ = fi101
                _ = v102
                _ = v103
            }
            switch p2 {
            case "foo bar":
                v106 := GenS12{F0: 0}
                _ = v106
            case "hello":
                if true {
                    var v107 uint32 = (p0 << (((v18 * v18)) & 63))
                    v108 := GenS13{F0: 0, F1: 42}
                    v109 := GenS1{F0: 0, F1: -864}
                    _ = v107
                    _ = v108
                    _ = v109
                } else {
                    for fi110 := 0; fi110 < 6; fi110++ {
                        v111 := []int{394, fi34, fi110}
                        switch fi34 {
                        case 0:
                            v18 += v18
                        case -9223372036854775808:
                            switch "hello" {
                            case "":
                                v112 := GenS14{F0: 1000, F1: 93}
                                _ = v112
                            case "foo bar":
                                v113 := GenS9{F0: p0, F1: 2147483647, F2: 42}
                                _ = v113
                            }
                        case -2147483648:
                            var v114 int = (fi110 % ((fi34) | 1))
                            _ = v114
                        default:
                            if ((fi34 & fi110) != 42) {
                                if len(v111) > 0 {
                                    v111[0] = fi110
                                }
                                _ = len(v111)
                            }
                        }
                        _ = fi110
                        _ = v111
                    }
                }
            }
            _ = fi34
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return -114
}

func main() {
    var v1 uint = 255
    var v2 int = 0
    v3 := []uint32{2147483647, uint32(v1)}
    fmt.Println(v2, v2)
    var v115 uint16 = 2
    fmt.Println(v115, v2)
    for fi116 := 0; fi116 < 10; fi116++ {
        v2 = (v2 - v2)
        v117 := GenS15{F0: 42, F1: -37}
        v118 := GenS4{F0: ("\n\t" + "hello"), F1: "\n\t", F2: 1}
        _ = fi116
        _ = v117
        _ = v118
    }
    if (true && (1.5 != 3.14159)) {
        var v119 uint32 = 911
        _ = v119
    } else {
        var v120 uint32 = uint32(v115)
        _ = len(v3)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v120
    }
    if false {
        v121 := GenS15{F0: 0, F1: -1}
        v122 := GenS4{F0: "a", F1: ((("hello" + ("\n\t" + "\n\t")) + "foo bar") + "foo bar"), F2: -1}
        v2 += (v2 >> (((v1 * (v1 << ((v1) & 63)))) & 63))
        _ = v121
        _ = v122
    }
    v123 := GenS6{F0: 255, F1: "", F2: -538}
    var v124 int16 = 0
    fmt.Println(v2, v115, v115)
    for fi125 := 0; fi125 < 18; fi125++ {
        var v126 string = "foo bar"
        _ = fi125
        _ = v126
    }
    var v127 float32 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v115:", v115)
    fmt.Println("v124:", v124)
    fmt.Println("v127:", v127)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v115
    _ = v123
    _ = v124
    _ = v127
    _ = v2
    _ = v3
}
