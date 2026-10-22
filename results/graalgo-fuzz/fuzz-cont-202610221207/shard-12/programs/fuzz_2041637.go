package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 uint; F1 uint8; F2 int }
type GenS4 struct { F0 string; F1 uint64 }
type GenS5 struct { F0 int8; F1 int; F2 uint8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 string; F1 uint64; F2 uint8 }
type GenS8 struct { F0 int16; F1 uint64 }
type GenS9 struct { F0 uint16; F1 uint16 }
type GenS10 struct { F0 int8; F1 int }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 int16; F1 string }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 int; F1 string; F2 bool }
type GenS15 struct { F0 int16; F1 uint; F2 int32 }
type GenS16 struct { F0 uint16; F1 int64; F2 int16 }
type GenS17 struct { F0 bool }
type GenS18 struct { F0 uint32 }
type GenS19 struct { F0 uint32; F1 uint32; F2 int16 }
type GenS20 struct { F0 int8; F1 uint; F2 int8 }
type GenS21 struct { F0 uint8; F1 int32; F2 uint16 }
func genFn0(p0 int, p1 int32) int64 {
    for fi4 := 0; fi4 < 15; fi4++ {
        for fi5 := 0; fi5 < 5; fi5++ {
            if (false && (2 != 18446744073709551615)) {
                switch "foo bar" {
                case "hello":
                    if (false || true) {
                        p0 -= -331
                        if true {
                            var v6 float32 = float32((0.1 + (1.5 + 0.1)))
                            var v7 int64 = 1
                            _ = v6
                            _ = v7
                        }
                        p1 = p1
                    } else {
                        switch "hello" {
                        case "foo bar":
                            var v8 string = "a"
                            _ = v8
                        }
                        var v9 int = (fi5 >> ((4294967295) & 63))
                        _ = v9
                    }
                }
                var v10 bool = true
                var v11 uint16 = 255
                _ = v10
                _ = v11
            } else {
                v12 := []int{(fi5 % (((fi5 % (((^(fi4))) | 1))) | 1)), (p0 - fi5), fi5, (fi5 & fi5)}
                var v13 int8 = 100
                if len(v12) > 0 {
                    v12[0] = (fi4 | fi5)
                }
                _ = v12
                _ = v13
            }
            if false {
                v14 := GenS1{F0: ("foo bar" + "foo bar"), F1: 42}
                _ = v14
            }
            _ = fi5
        }
        v15 := []int{(p0 << ((661) & 63)), 42}
        _ = fi4
        _ = v15
    }
    if (42 == 2) {
        var v16 string = "hello"
        for fi17 := 0; fi17 < 7; fi17++ {
            var v18 string = (v16 + ((v16 + v16) + v16))
            _ = fi17
            _ = v18
        }
        _ = v16
    }
    p1 = (-(p1))
    for fi19 := 0; fi19 < 6; fi19++ {
        v20 := GenS1{F0: "hello", F1: 42}
        _ = fi19
        _ = v20
    }
    if (65535 != 65535) {
        var v21 string = "\n\t"
        fmt.Println(v21)
        _ = v21
    } else {
        if true {
            for fi22 := 0; fi22 < 1; fi22++ {
                p1 += 0
                p0 += p0
                _ = fi22
            }
            for fi23 := 0; fi23 < 15; fi23++ {
                p1 = (p1 | (p1 % ((p1) | 1)))
                v24 := GenS2{F0: true, F1: 472}
                _ = fi23
                _ = v24
            }
            p1 += (p1 / ((p1) | 1))
        }
    }
    _ = p0
    _ = p1
    return 49
}
func genFn1() int64 {
    var v25 int32 = -2147483648
    if (0.1 == 1.5) {
        var v26 uint32 = 6
        _ = v26
    } else {
        v27 := GenS1{F0: "hello", F1: -9223372036854775808}
        _ = v27
    }
    if (3.14159 != float32(float64((1.5 * float32(1.5))))) {
        v25 = v25
    } else {
        v28 := GenS1{F0: "", F1: 42}
        v29 := GenS1{F0: "a", F1: 0}
        _ = v28
        _ = v29
    }
    switch 32768 {
    case -879:
        switch "\n\t" {
        case "\n\t":
            v30 := GenS2{F0: (v25 == 2147483647), F1: 65535}
            _ = v30
        default:
            fmt.Println(v25)
        }
    case 910:
        v31 := GenS3{F0: 62, F1: 127, F2: -52}
        _ = v31
    }
    v32 := []int{}
    _ = v25
    _ = v32
    return -9223372036854775808
}

func main() {
    var v1 uint64 = 594
    var v2 float32 = 0.0
    v3 := make(map[uint32]int)
    var v33 int8 = 42
    var v34 int64 = 1
    v35 := GenS1{F0: "hello", F1: (v34 * v34)}
    fmt.Println(v33)
    v36 := GenS1{F0: ("foo bar" + "a"), F1: v34}
    var v37 string = ("\n\t" + "a")
    for fi38 := 0; fi38 < 14; fi38++ {
        v39 := GenS4{F0: v37, F1: v1}
        _ = fi38
        _ = v39
    }
    var v40 int8 = (v33 - (v33 + v33))
    var v41 uint32 = 4294967295
    v42 := []int{4294967295, 42}
    var v43 bool = true
    if (!v43) {
        switch "hello" {
        case "a":
            fmt.Println(v41)
        default:
            if v43 {
                v44 := GenS5{F0: (v40 * (v33 - v40)), F1: -2147483648, F2: 128}
                v42 = append(v42, 9223372036854775807)
                _ = v44
            } else {
                fmt.Println(v2)
                v2 -= float32(3.14159)
            }
        }
    } else {
        var v45 int8 = (v33 % ((v33) | 1))
        v40 += v40
        _ = v45
    }
    if (71 < 135) {
        v41 += 614
    } else {
        v34 += (v34 / ((v34) | 1))
    }
    if len(v42) >= 2 {
        v42 = v42[1:len(v42)]
    }
    if false {
        v46 := v3[922]
        _ = v46
    } else {
        if v43 {
            var v47 int = -1
            _ = v47
        }
        _ = len(v3)
        var v48 float64 = ((1e10 + 0.1) * (1.5 + (1e10 * (1e10 * 0.1))))
        _ = v48
    }
    for fi49 := 0; fi49 < 14; fi49++ {
        v43 = true
        v50 := GenS6{F0: 4294967295}
        _ = fi49
        _ = v50
    }
    for fi51 := 0; fi51 < 11; fi51++ {
        _ = len(v42)
        v52 := v3[(^(v41))]
        for fi53 := 0; fi53 < 11; fi53++ {
            var v54 bool = v43
            var v55 uint16 = 255
            if v54 {
                v56 := GenS2{F0: true, F1: 111}
                switch -886 {
                case 32768:
                    if v43 {
                        switch (-(fi51)) {
                        case 32767:
                            var v57 int16 = 1000
                            _ = v57
                        case 343:
                            var v58 float64 = (-2.25 * -2.25)
                            _ = v58
                        case 32768:
                            if ((v43 || v43) && false) {
                                var v59 uint = 460
                                v59 += 0
                                _ = v59
                            }
                        }
                        var v60 uint = 709
                        if len(v42) >= 2 {
                            v42 = v42[1:len(v42)]
                        }
                        if ((v55 > v55) && ((uint16(v41) == (v55 + 65535)) || (((!(!(!(31 == v55)))) || v54) && true))) {
                            v43 = v43
                            if true {
                                v33 -= (v33 % (((-(v40))) | 1))
                                fmt.Println(v1)
                                if (((1.5 + 3.14159) + (3.14159 + (-2.25 + 0.1))) < ((1e10 / 1.5) - 1.5)) {
                                    if true {
                                        if v54 {
                                            v41 = v41
                                        }
                                    } else {
                                        v41 -= 635
                                    }
                                }
                            } else {
                                if (v37 < v37) {
                                    var v61 uint16 = 32767
                                    _ = v61
                                } else {
                                    fmt.Println(v52, v34, v33)
                                    v62 := GenS5{F0: (v40 - 100), F1: fi53, F2: 255}
                                    v63 := GenS2{F0: false, F1: (v60 * 2)}
                                    switch "" {
                                    case "hello":
                                        v34 = (v34 | (v34 ^ (v34 & (-(v34)))))
                                    case "":
                                        v52 -= int(v41)
                                    case "foo bar":
                                        var v64 float64 = 1.5
                                        _ = v64
                                    }
                                    _ = v62
                                    _ = v63
                                }
                            }
                            if v43 {
                                if v54 {
                                    if ((v60 * (v60 ^ v60)) == v60) {
                                        var v65 uint16 = (v55 & 2)
                                        _ = v65
                                    }
                                    v66 := GenS7{F0: (v37 + (((v37 + "\n\t") + (v37 + v37)) + "hello")), F1: v1, F2: 35}
                                    _ = v66
                                } else {
                                    if (((((v37 + "a") + "") + v37) + "\n\t") <= ("foo bar" + ("foo bar" + v37))) {
                                        if false {
                                            if (!(!true)) {
                                                if len(v42) > 0 {
                                                    v42[0] = -9223372036854775808
                                                }
                                                v67 := GenS2{F0: v43, F1: (v60 - (v60 + uint(v41)))}
                                                fmt.Println(v33, v41)
                                                if v54 {
                                                    v68 := GenS8{F0: 1, F1: v1}
                                                    _ = v68
                                                }
                                                _ = v67
                                            }
                                        }
                                    }
                                    if (v54 || (!v54)) {
                                        if v54 {
                                            if true {
                                                v55 = (^(v55))
                                            }
                                            var v69 int64 = v34
                                            if (false || v54) {
                                                v70 := GenS3{F0: (v60 - (v60 << (((v60 % (((v60 % ((v60) | 1))) | 1))) & 63))), F1: 127, F2: 1}
                                                if len(v42) >= 2 {
                                                    v42 = v42[1:len(v42)]
                                                }
                                                v71 := GenS9{F0: v55, F1: (v55 ^ 0)}
                                                _ = v70
                                                _ = v71
                                            }
                                            _ = v69
                                        } else {
                                            var v72 int = int(v60)
                                            fmt.Println(v72)
                                            _ = v72
                                        }
                                        v60 -= v60
                                        var v73 int8 = -102
                                        _ = v73
                                    } else {
                                        v74 := GenS10{F0: (v40 % ((v40) | 1)), F1: v52}
                                        v42 = append(v42, (^(fi51)))
                                        switch fi53 {
                                        case 356:
                                            var v75 uint16 = (^(v55))
                                            _ = v75
                                        case 2147483647:
                                            v1 = uint64(v60)
                                        default:
                                            v40 += (v33 | int8(v1))
                                        }
                                        _ = v74
                                    }
                                }
                                switch 127 {
                                case 2147483647:
                                    v40 -= 0
                                }
                            } else {
                                v76 := GenS10{F0: (-(v33)), F1: (fi51 * (fi51 >> ((42) & 63)))}
                                if (v43 || (v34 <= (v34 & v34))) {
                                    if v43 {
                                        v77 := GenS4{F0: ((v37 + v37) + v37), F1: (v1 * 0)}
                                        switch (^(fi53)) {
                                        case -64:
                                            v34 = 1
                                        }
                                        _ = v77
                                    }
                                    switch "\n\t" {
                                    case "hello":
                                        _ = len(v42)
                                    case "":
                                        switch fi51 {
                                        case 4294967295:
                                            fmt.Println(v54, fi53)
                                        case 2147483647:
                                            if v54 {
                                                v60 = v60
                                                v78 := GenS5{F0: v33, F1: (fi51 | fi51), F2: 19}
                                                _ = v78
                                            }
                                            var v79 uint16 = v55
                                            _ = v79
                                        case -1:
                                            switch "" {
                                            case "\n\t":
                                                v34 += genFn1()
                                            case "":
                                                if len(v42) >= 2 {
                                                    v42 = v42[1:len(v42)]
                                                }
                                            default:
                                                var v80 int32 = 0
                                                _ = v80
                                            }
                                        }
                                    default:
                                        switch v52 {
                                        case 1:
                                            if v43 {
                                                v41 += 2147483648
                                            } else {
                                                var v81 int32 = 2147483647
                                                if v43 {
                                                    fmt.Println(v34, v54)
                                                    var v82 float32 = v2
                                                    _ = v82
                                                }
                                                v83 := GenS10{F0: (v40 / ((v33) | 1)), F1: fi53}
                                                v84 := GenS5{F0: v40, F1: 461, F2: 0}
                                                v55 += v55
                                                _ = v81
                                                _ = v83
                                                _ = v84
                                            }
                                        case 4294967295:
                                            if ((v60 & (v60 << (((v60 % (((v60 / ((v60) | 1))) | 1))) & 63))) == (v60 << (((v60 << ((698) & 63))) & 63))) {
                                                v34 = v34
                                                switch (fi51 << ((699) & 63)) {
                                                case 2147483647:
                                                    var v85 uint8 = 90
                                                    _ = v85
                                                case 127:
                                                    var v86 float64 = 0.1
                                                    v3[(v41 + v41)] = (v52 | 512)
                                                    _ = v86
                                                case -432:
                                                    v40 += (v33 & (v40 >> ((v60) & 63)))
                                                default:
                                                    var v87 bool = (!v43)
                                                    _ = v87
                                                }
                                                var v88 uint16 = 0
                                                _ = v88
                                            } else {
                                                switch "\n\t" {
                                                case "hello":
                                                    fmt.Println(v55, v55, v52)
                                                case "foo bar":
                                                    switch (((v37 + v37) + ("a" + v37)) + ("hello" + "\n\t")) {
                                                    case "a":
                                                        if v43 {
                                                            if v43 {
                                                                fmt.Println(v55)
                                                                fmt.Println(v2, v2)
                                                            } else {
                                                                var v89 uint = (v60 & (v60 - 430))
                                                                var v90 uint64 = v1
                                                                _ = len(v3)
                                                                if v43 {
                                                                    v91 := v3[165]
                                                                    var v92 int64 = v34
                                                                    if true {
                                                                        if v54 {
                                                                            if v43 {
                                                                                switch v52 {
                                                                                case 0:
                                                                                    if (v41 < (v41 >> ((336) & 63))) {
                                                                                        v91 -= (^(fi51))
                                                                                        var v93 float64 = float64((v2 / v2))
                                                                                        var v94 uint32 = v41
                                                                                        _ = v93
                                                                                        _ = v94
                                                                                    }
                                                                                case 4294967295:
                                                                                    switch v37 {
                                                                                    case "\n\t":
                                                                                        v3[(v41 ^ (v41 * (v41 - (v41 / (((v41 >> (((v89 & (v60 | 679))) & 63))) | 1)))))] = (-(v52))
                                                                                        if ((^(fi51)) != (fi53 + v52)) {
                                                                                            fmt.Println(v33, v60, v1)
                                                                                            switch 349 {
                                                                                            case -344:
                                                                                                v95 := GenS6{F0: (v41 ^ v41)}
                                                                                                _ = v95
                                                                                            }
                                                                                            fmt.Println(v2)
                                                                                            switch fi53 {
                                                                                            case 764:
                                                                                                var v96 string = ("foo bar" + v37)
                                                                                                _ = v96
                                                                                            case 32767:
                                                                                                v41 -= 0
                                                                                            case 1:
                                                                                                if v43 {
                                                                                                    if false {
                                                                                                        if v43 {
                                                                                                            fmt.Println(v89, v52)
                                                                                                        }
                                                                                                        fmt.Println(v55, v43)
                                                                                                        var v97 uint32 = v41
                                                                                                        _ = v97
                                                                                                    }
                                                                                                    v98 := v3[(v41 >> (((v89 % (((v89 + (v89 & v60))) | 1))) & 63))]
                                                                                                    _ = v98
                                                                                                } else {
                                                                                                    v91 += 0
                                                                                                    if len(v42) >= 2 {
                                                                                                        v42 = v42[1:len(v42)]
                                                                                                    }
                                                                                                }
                                                                                            default:
                                                                                                fmt.Println(v43, v60)
                                                                                            }
                                                                                        }
                                                                                        _ = len(v3)
                                                                                        _ = len(v3)
                                                                                        v99 := GenS6{F0: uint32(v1)}
                                                                                        _ = v99
                                                                                    }
                                                                                case -9223372036854775808:
                                                                                    if ((32767 == -1) || v43) {
                                                                                        delete(v3, (v41 - 4294967295))
                                                                                    } else {
                                                                                        var v100 uint64 = (^(v1))
                                                                                        _ = v100
                                                                                    }
                                                                                default:
                                                                                    if false {
                                                                                        var v101 bool = (-128 <= -1)
                                                                                        _ = v101
                                                                                    }
                                                                                }
                                                                                v102 := GenS11{F0: (v40 * v33)}
                                                                                fmt.Println(v89, v1, v43)
                                                                                var v103 uint16 = uint16(v89)
                                                                                _ = v102
                                                                                _ = v103
                                                                            } else {
                                                                                switch ((v37 + ("hello" + v37)) + "a") {
                                                                                case "\n\t":
                                                                                    var v104 uint8 = uint8(v55)
                                                                                    _ = v104
                                                                                case "a":
                                                                                    var v105 int = -742
                                                                                    _ = v105
                                                                                default:
                                                                                    fmt.Println(v37, v52)
                                                                                }
                                                                                fmt.Println(fi51)
                                                                            }
                                                                            if (127 != 0) {
                                                                                v106 := GenS8{F0: 32767, F1: (^(v1))}
                                                                                _ = v106
                                                                            }
                                                                            v107 := GenS5{F0: (v33 - 86), F1: (v91 + v52), F2: 0}
                                                                            _ = v107
                                                                        } else {
                                                                            v60 -= v89
                                                                        }
                                                                        v108 := GenS12{F0: -1, F1: ""}
                                                                        _ = v108
                                                                    } else {
                                                                        fmt.Println(v92, v1)
                                                                        if (true || false) {
                                                                            switch "" {
                                                                            case "a":
                                                                                v109 := GenS13{F0: genFn1()}
                                                                                switch "" {
                                                                                case "":
                                                                                    if v54 {
                                                                                        v89 = 42
                                                                                        if v43 {
                                                                                            var v110 uint = (v89 & 112)
                                                                                            _ = v110
                                                                                        } else {
                                                                                            v55 -= uint16(v89)
                                                                                            v111 := v3[(v41 ^ 845)]
                                                                                            _ = v111
                                                                                        }
                                                                                    } else {
                                                                                        v112 := GenS14{F0: fi51, F1: "a", F2: false}
                                                                                        var v113 int16 = int16(v60)
                                                                                        var v114 uint8 = 127
                                                                                        _ = v112
                                                                                        _ = v113
                                                                                        _ = v114
                                                                                    }
                                                                                case "hello":
                                                                                    v115 := GenS5{F0: 100, F1: (fi51 - v52), F2: 127}
                                                                                    _ = v115
                                                                                }
                                                                                var v116 uint64 = v90
                                                                                v89 = (v89 >> (((^(v89))) & 63))
                                                                                var v117 float64 = 0.1
                                                                                fmt.Println(v1)
                                                                                _ = v109
                                                                                _ = v116
                                                                                _ = v117
                                                                            case "hello":
                                                                                v118 := GenS6{F0: v41}
                                                                                _ = v118
                                                                            default:
                                                                                if len(v42) > 0 {
                                                                                    v42[0] = (^(v52))
                                                                                }
                                                                            }
                                                                            v42 = append(v42, v91)
                                                                            if (v54 && true) {
                                                                                v1 = 510
                                                                                v89 = (v60 ^ v89)
                                                                                v42 = append(v42, (v52 % ((-9223372036854775808) | 1)))
                                                                            }
                                                                            if (-2147483648 != 2147483647) {
                                                                                v89 = (v89 ^ 1)
                                                                            } else {
                                                                                if ((v41 | 1) != (v41 % ((2147483648) | 1))) {
                                                                                    var v119 bool = (v43 || ((v40 & (v40 & 1)) >= v33))
                                                                                    var v120 int32 = -1
                                                                                    fmt.Println(v52)
                                                                                    _ = v119
                                                                                    _ = v120
                                                                                } else {
                                                                                    fmt.Println(v34)
                                                                                    var v121 bool = ((v1 - 18446744073709551615) <= 42)
                                                                                    _ = v121
                                                                                }
                                                                                if v54 {
                                                                                    delete(v3, v41)
                                                                                    fmt.Println(v33)
                                                                                } else {
                                                                                    fmt.Println(v37)
                                                                                    if len(v42) >= 2 {
                                                                                        v42 = v42[1:len(v42)]
                                                                                    }
                                                                                    v42 = append(v42, v52)
                                                                                }
                                                                            }
                                                                        }
                                                                        v122 := GenS15{F0: -283, F1: 42, F2: 255}
                                                                        _ = v122
                                                                    }
                                                                    v123 := GenS16{F0: (v55 | v55), F1: genFn1(), F2: 32767}
                                                                    var v124 uint64 = (v1 | (^(v90)))
                                                                    fmt.Println(v37, v91, v43)
                                                                    var v125 uint = (v60 ^ (v89 * v89))
                                                                    _ = v123
                                                                    _ = v124
                                                                    _ = v125
                                                                    _ = v91
                                                                    _ = v92
                                                                } else {
                                                                    v126 := GenS12{F0: 371, F1: v37}
                                                                    var v127 bool = (42 <= -1)
                                                                    v128 := GenS17{F0: (v34 > int64(v90))}
                                                                    _ = v126
                                                                    _ = v127
                                                                    _ = v128
                                                                }
                                                                v129 := GenS5{F0: -36, F1: (fi51 * 298), F2: 1}
                                                                v42 = append(v42, (v52 + 4294967295))
                                                                if ((v33 / ((v40) | 1)) > v33) {
                                                                    fmt.Println(v33, v2)
                                                                    if true {
                                                                        var v130 uint64 = v90
                                                                        switch int(v60) {
                                                                        case -2147483648:
                                                                            switch (fi53 << ((18446744073709551615) & 63)) {
                                                                            case -1:
                                                                                if (0.0 == -2.25) {
                                                                                    v130 -= v1
                                                                                }
                                                                            case -2147483648:
                                                                                fmt.Println(v43)
                                                                            case 4294967295:
                                                                                switch "" {
                                                                                case "foo bar":
                                                                                    v60 -= uint(v130)
                                                                                    if v43 {
                                                                                        if ((v41 / ((1) | 1)) > 2147483647) {
                                                                                            v131 := GenS18{F0: (v41 - (v41 >> ((v60) & 63)))}
                                                                                            v55 -= v55
                                                                                            fmt.Println(fi53, v55)
                                                                                            _ = v131
                                                                                        }
                                                                                        v54 = (v60 == v60)
                                                                                        v41 -= v41
                                                                                        var v132 int32 = -750
                                                                                        _ = v132
                                                                                    } else {
                                                                                        fmt.Println(fi53, v37, fi53)
                                                                                        var v133 int32 = -1
                                                                                        v41 -= v41
                                                                                        _ = v133
                                                                                    }
                                                                                    var v134 bool = v54
                                                                                    _ = v134
                                                                                default:
                                                                                    var v135 uint16 = (v55 & v55)
                                                                                    _ = v135
                                                                                }
                                                                            }
                                                                        }
                                                                        fmt.Println(v89, v34)
                                                                        _ = v130
                                                                    }
                                                                } else {
                                                                    if v54 {
                                                                        var v136 uint16 = v55
                                                                        switch (^(fi51)) {
                                                                        case -2147483648:
                                                                            switch -886 {
                                                                            case 4294967295:
                                                                                v1 = v90
                                                                            case -529:
                                                                                if v43 {
                                                                                    if true {
                                                                                        v33 += int8(v90)
                                                                                    } else {
                                                                                        _ = len(v3)
                                                                                    }
                                                                                    switch ("\n\t" + ("a" + "\n\t")) {
                                                                                    case "\n\t":
                                                                                        v3[uint32(v60)] = (fi51 + int(v33))
                                                                                    default:
                                                                                        if (v54 && v54) {
                                                                                            if (true || v54) {
                                                                                                if (!(!(5 != (v34 << ((v60) & 63))))) {
                                                                                                    v42 = append(v42, fi51)
                                                                                                    v136 = (v136 % ((255) | 1))
                                                                                                } else {
                                                                                                    if len(v42) >= 2 {
                                                                                                        v42 = v42[1:len(v42)]
                                                                                                    }
                                                                                                }
                                                                                            } else {
                                                                                                var v137 int = fi51
                                                                                                _ = v137
                                                                                            }
                                                                                            if ((fi53 - -506) < (fi51 + fi51)) {
                                                                                                v138 := GenS19{F0: v41, F1: v41, F2: int16(v40)}
                                                                                                var v139 uint32 = (^(v41))
                                                                                                _ = v138
                                                                                                _ = v139
                                                                                            }
                                                                                            v140 := GenS20{F0: (v33 & v40), F1: 596, F2: -128}
                                                                                            _ = v140
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    v2 = float32(1.5)
                                                                                    v141 := v3[(v41 & (v41 & v41))]
                                                                                    _ = v141
                                                                                }
                                                                            case -9223372036854775808:
                                                                                var v142 int8 = (v33 % (((v40 ^ v40)) | 1))
                                                                                _ = v142
                                                                            }
                                                                        }
                                                                        _ = v136
                                                                    } else {
                                                                        fmt.Println(v43, v33, v33)
                                                                        var v143 string = v37
                                                                        v3[v41] = fi51
                                                                        _ = v143
                                                                    }
                                                                }
                                                                v3[v41] = (v52 / (((v52 * v52)) | 1))
                                                                _ = v129
                                                                _ = v89
                                                                _ = v90
                                                            }
                                                            if len(v42) > 0 {
                                                                v42[0] = (v52 | fi53)
                                                            }
                                                        } else {
                                                            if v54 {
                                                                v144 := GenS21{F0: 0, F1: 145, F2: 65535}
                                                                fmt.Println(v60, fi53, v41)
                                                                var v145 int = 42
                                                                _ = v144
                                                                _ = v145
                                                            } else {
                                                                var v146 uint64 = uint64(v41)
                                                                _ = v146
                                                            }
                                                            if len(v42) > 0 {
                                                                v42[0] = (-(fi51))
                                                            }
                                                        }
                                                    }
                                                default:
                                                    fmt.Println(v1)
                                                }
                                            }
                                        case -439:
                                            var v147 int32 = -415
                                            _ = v147
                                        }
                                    }
                                }
                                if v54 {
                                    var v148 uint64 = 149
                                    _ = v148
                                }
                                v149 := v3[v41]
                                _ = v149
                                _ = v76
                            }
                        } else {
                            if len(v42) > 0 {
                                v42[0] = (-(fi53))
                            }
                            var v150 uint64 = (^(v1))
                            _ = v150
                        }
                        switch v37 {
                        case "a":
                            fmt.Println(v34, v41)
                        }
                        var v151 int16 = -32768
                        _ = v151
                        _ = v60
                    } else {
                        _ = len(v42)
                        v152 := GenS21{F0: 255, F1: -2147483648, F2: 602}
                        _ = v152
                    }
                case 2147483647:
                    v41 += (v41 << ((255) & 63))
                }
                _ = v56
            }
            _ = fi53
            _ = v54
            _ = v55
        }
        _ = fi51
        _ = v52
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v40
    _ = v41
    _ = v42
    _ = v43
}
