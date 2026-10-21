package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint32 }
type GenS2 struct { F0 uint64; F1 uint16; F2 int16 }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 string; F1 uint16 }
type GenS5 struct { F0 int64; F1 uint }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int8; F1 uint64 }
type GenS8 struct { F0 int64; F1 uint }
type GenS9 struct { F0 string; F1 int8; F2 uint32 }
type GenS10 struct { F0 uint16; F1 int32 }
type GenS11 struct { F0 int8; F1 string; F2 int32 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 bool; F1 int64; F2 int }
type GenS14 struct { F0 int32; F1 uint32 }
type GenS15 struct { F0 int64; F1 bool; F2 uint32 }
type GenS16 struct { F0 int32 }
type GenS17 struct { F0 int16 }
type GenS18 struct { F0 int; F1 int16; F2 uint32 }
type GenS19 struct { F0 int64 }
type GenS20 struct { F0 int8; F1 int8; F2 uint32 }
type GenS21 struct { F0 int64; F1 uint; F2 uint }
type GenS22 struct { F0 string; F1 bool; F2 int }
type GenS23 struct { F0 uint8; F1 int8 }
type GenS24 struct { F0 uint }

func main() {
    v1 := []string{"", (("\n\t" + ("\n\t" + ("foo bar" + "hello"))) + "hello")}
    var v2 int8 = -91
    v3 := make(map[int]string)
    switch "foo bar" {
    case "a":
        switch 1 {
        case 4294967295:
            for fi4 := 0; fi4 < 2; fi4++ {
                for fi5 := 0; fi5 < 6; fi5++ {
                    fmt.Println(v2)
                    v6 := GenS1{F0: v2, F1: (v2 & (v2 % ((127) | 1))), F2: 2147483648}
                    _ = fi5
                    _ = v6
                }
                for fi7 := 0; fi7 < 17; fi7++ {
                    if (-483 != -1) {
                        v8 := GenS2{F0: 4294967295, F1: 411, F2: 1000}
                        v9 := GenS3{F0: 42, F1: 367}
                        v10 := GenS2{F0: 42, F1: 42, F2: 0}
                        _ = v10
                        _ = v8
                        _ = v9
                    } else {
                        var v11 uint16 = 32767
                        _ = v11
                    }
                    var v12 int16 = 0
                    var v13 float32 = (3.14159 / (1e10 * (1.5 + (((0.1 / (0.1 + float32(3.14159))) * 3.14159) - 0.0))))
                    var v14 int8 = 44
                    v15 := GenS3{F0: 847, F1: 2}
                    _ = fi7
                    _ = v12
                    _ = v13
                    _ = v14
                    _ = v15
                }
                _ = fi4
            }
        default:
            if true {
                if true {
                    v2 = v2
                    for fi16 := 0; fi16 < 18; fi16++ {
                        _ = len(v1)
                        _ = fi16
                    }
                } else {
                    var v17 uint = 65535
                    v3[923] = "hello"
                    _ = v17
                }
            } else {
                if false {
                    _ = len(v1)
                    if false {
                        v18 := GenS4{F0: "", F1: 1}
                        _ = v18
                    } else {
                        var v19 int = 925
                        _ = v19
                    }
                    if true {
                        var v20 uint = 2
                        v21 := GenS2{F0: 2, F1: 65535, F2: 1}
                        _ = v20
                        _ = v21
                    } else {
                        switch -774 {
                        case 868:
                            for fi22 := 0; fi22 < 7; fi22++ {
                                v23 := GenS3{F0: 42, F1: 42}
                                v24 := GenS4{F0: ("hello" + ""), F1: 106}
                                if (2147483647 <= -2147483648) {
                                    if (934 <= 333) {
                                        for fi25 := 0; fi25 < 5; fi25++ {
                                            _ = len(v3)
                                            v26 := v3[127]
                                            if false {
                                                if true {
                                                    fmt.Println(v26, v26, fi22)
                                                }
                                            }
                                            _ = fi25
                                            _ = v26
                                        }
                                    }
                                    for fi27 := 0; fi27 < 11; fi27++ {
                                        var v28 int32 = 1
                                        v28 -= int32(v2)
                                        _ = fi27
                                        _ = v28
                                    }
                                    if true {
                                        v29 := GenS2{F0: 1, F1: 2, F2: 1000}
                                        if ((42 > 2147483647) && (-2.25 > 1e10)) {
                                            if (-2.25 != ((((0.1 + (0.0 / 3.14159)) * 1.5) / 0.1) / float32((0.1 * ((1.5 * 0.0) + (-2.25 * 3.14159)))))) {
                                                v30 := GenS4{F0: ("" + ("\n\t" + ("" + "a"))), F1: 32767}
                                                var v31 int32 = -2147483648
                                                v32 := GenS1{F0: int8(fi22), F1: 100, F2: 559}
                                                _ = v30
                                                _ = v31
                                                _ = v32
                                            } else {
                                                for fi33 := 0; fi33 < 4; fi33++ {
                                                    if false {
                                                        switch (^(fi22)) {
                                                        case 798:
                                                            var v34 uint16 = 526
                                                            _ = v34
                                                        case -1:
                                                            if true {
                                                                v35 := GenS1{F0: (v2 - v2), F1: v2, F2: 42}
                                                                if false {
                                                                    var v36 uint8 = 4
                                                                    switch (fi33 ^ -2147483648) {
                                                                    case -9223372036854775808:
                                                                        switch "foo bar" {
                                                                        case "hello":
                                                                            v36 = v36
                                                                        }
                                                                    case 344:
                                                                        var v37 uint = 4294967295
                                                                        _ = v37
                                                                    case 32768:
                                                                        var v38 int32 = -715
                                                                        _ = v38
                                                                    default:
                                                                        switch fi33 {
                                                                        case 42:
                                                                            var v39 uint = 18446744073709551615
                                                                            _ = v39
                                                                        case 4294967295:
                                                                            switch (fi33 ^ (fi33 << ((340) & 63))) {
                                                                            case 401:
                                                                                v40 := GenS2{F0: 492, F1: 823, F2: -32768}
                                                                                _ = v40
                                                                            default:
                                                                                switch "hello" {
                                                                                case "\n\t":
                                                                                    v41 := GenS5{F0: 1, F1: 0}
                                                                                    _ = v41
                                                                                default:
                                                                                    fmt.Println(v2, fi33, v36)
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                    var v42 int16 = -446
                                                                    _ = v36
                                                                    _ = v42
                                                                } else {
                                                                    if false {
                                                                        if false {
                                                                            var v43 uint16 = 0
                                                                            _ = v43
                                                                        } else {
                                                                            v44 := GenS1{F0: (v2 / (((v2 / (((v2 & (v2 * v2))) | 1))) | 1)), F1: -29, F2: 310}
                                                                            var v45 uint32 = 258
                                                                            _ = v44
                                                                            _ = v45
                                                                        }
                                                                        var v46 int16 = 32767
                                                                        _ = v46
                                                                    } else {
                                                                        if true {
                                                                            var v47 int16 = 0
                                                                            _ = v47
                                                                        } else {
                                                                            v1 = append(v1, "a")
                                                                            var v48 float64 = 1.5
                                                                            v2 = v2
                                                                            _ = v48
                                                                        }
                                                                        v1 = append(v1, "hello")
                                                                    }
                                                                    if (!(false && true)) {
                                                                        v2 -= (v2 / ((v2) | 1))
                                                                    }
                                                                    delete(v3, fi22)
                                                                }
                                                                var v49 float32 = (1e10 + 1e10)
                                                                var v50 bool = true
                                                                v49 += (0.0 / 3.14159)
                                                                _ = v35
                                                                _ = v49
                                                                _ = v50
                                                            } else {
                                                                v2 = v2
                                                                v2 = v2
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            }
                                                        case -9223372036854775808:
                                                            v51 := GenS6{F0: 32767}
                                                            _ = v51
                                                        default:
                                                            switch (fi33 << ((255) & 63)) {
                                                            case 1:
                                                                v52 := GenS7{F0: v2, F1: 2}
                                                                _ = v52
                                                            }
                                                        }
                                                    } else {
                                                        switch (fi33 ^ fi33) {
                                                        case 127:
                                                            var v53 int64 = -670
                                                            switch fi22 {
                                                            case 127:
                                                                var v54 uint16 = 1
                                                                _ = v54
                                                            case 2147483647:
                                                                delete(v3, (fi22 << ((65535) & 63)))
                                                            default:
                                                                switch -632 {
                                                                case 42:
                                                                    v2 -= (v2 / (((v2 - -1)) | 1))
                                                                case 730:
                                                                    fmt.Println(v2, fi22)
                                                                case 32768:
                                                                    if (fi22 == (fi22 * fi22)) {
                                                                        if (true || true) {
                                                                            fmt.Println(fi33)
                                                                        } else {
                                                                            v55 := GenS8{F0: v53, F1: 0}
                                                                            switch fi22 {
                                                                            case 0:
                                                                                v2 = v2
                                                                            case 127:
                                                                                switch ("" + "hello") {
                                                                                case "":
                                                                                    var v56 uint8 = 2
                                                                                    _ = v56
                                                                                default:
                                                                                    v3[fi33] = ("" + "hello")
                                                                                }
                                                                            case -9223372036854775808:
                                                                                if ((v53 >= v53) || (4294967295 < 59)) {
                                                                                    v57 := GenS8{F0: int64(fi33), F1: 533}
                                                                                    if (!false) {
                                                                                        v53 -= (v53 - v53)
                                                                                        if ((v2 * (v2 * (v2 ^ (^(v2))))) != v2) {
                                                                                            v2 = -1
                                                                                            fmt.Println(fi22)
                                                                                            if ((!(3.14159 <= (1e10 - (1.5 * 1.5)))) && (!true)) {
                                                                                                switch "" {
                                                                                                case "\n\t":
                                                                                                    switch (fi22 + fi33) {
                                                                                                    case 782:
                                                                                                        if false {
                                                                                                            var v58 bool = true
                                                                                                            fmt.Println(v58)
                                                                                                            v58 = v58
                                                                                                            _ = v58
                                                                                                        } else {
                                                                                                            v53 += (v53 * 734)
                                                                                                            var v59 uint64 = 42
                                                                                                            _ = v59
                                                                                                        }
                                                                                                    default:
                                                                                                        var v60 uint8 = 255
                                                                                                        v53 += v53
                                                                                                        _ = v60
                                                                                                    }
                                                                                                case "hello":
                                                                                                    fmt.Println(fi33, v53)
                                                                                                }
                                                                                                if (((0.1 + 0.1) + (((1e10 * 3.14159) * (1e10 + 3.14159)) / 3.14159)) < -2.25) {
                                                                                                    var v61 uint = 70
                                                                                                    var v62 int32 = 1
                                                                                                    _ = v61
                                                                                                    _ = v62
                                                                                                }
                                                                                            } else {
                                                                                                fmt.Println(v2)
                                                                                                v63 := v3[-2147483648]
                                                                                                _ = v63
                                                                                            }
                                                                                        } else {
                                                                                            v2 = v2
                                                                                            if len(v1) > 0 {
                                                                                                v1[0] = ("foo bar" + "a")
                                                                                            }
                                                                                            var v64 int = (fi33 << ((42) & 63))
                                                                                            _ = v64
                                                                                        }
                                                                                        v65 := GenS4{F0: "", F1: 811}
                                                                                        _ = v65
                                                                                    }
                                                                                    var v66 uint64 = 42
                                                                                    if true {
                                                                                        var v67 int = 4294967295
                                                                                        v2 -= v2
                                                                                        _ = v67
                                                                                    } else {
                                                                                        fmt.Println(v53, v66, fi22)
                                                                                    }
                                                                                    v68 := GenS7{F0: (-(v2)), F1: 1}
                                                                                    v3[(fi22 << ((708) & 63))] = "\n\t"
                                                                                    switch "" {
                                                                                    case "":
                                                                                        v53 += v53
                                                                                    case "\n\t":
                                                                                        v66 += v66
                                                                                    default:
                                                                                        if (v53 >= (v53 >> ((375) & 63))) {
                                                                                            var v69 float64 = 3.14159
                                                                                            fmt.Println(v53, fi22)
                                                                                            v69 -= 0.1
                                                                                            v70 := GenS7{F0: (v2 + 100), F1: v66}
                                                                                            _ = v69
                                                                                            _ = v70
                                                                                        } else {
                                                                                            var v71 uint = 1
                                                                                            if ((v2 - v2) > v2) {
                                                                                                v72 := GenS8{F0: 1, F1: 0}
                                                                                                var v73 string = "\n\t"
                                                                                                _ = v72
                                                                                                _ = v73
                                                                                            }
                                                                                            v66 -= v66
                                                                                            v74 := GenS9{F0: "a", F1: (v2 / (((v2 * (v2 | (-(v2))))) | 1)), F2: 2}
                                                                                            v75 := GenS10{F0: 1, F1: 2147483647}
                                                                                            _ = v71
                                                                                            _ = v74
                                                                                            _ = v75
                                                                                        }
                                                                                        switch "" {
                                                                                        case "hello":
                                                                                            if true {
                                                                                                var v76 float32 = 1e10
                                                                                                _ = v76
                                                                                            } else {
                                                                                                v3[fi22] = "hello"
                                                                                                var v77 bool = (!true)
                                                                                                var v78 int16 = int16(v53)
                                                                                                switch "a" {
                                                                                                case "\n\t":
                                                                                                    v79 := GenS11{F0: int8(v78), F1: "foo bar", F2: 752}
                                                                                                    _ = v79
                                                                                                }
                                                                                                _ = v77
                                                                                                _ = v78
                                                                                            }
                                                                                        case "a":
                                                                                            var v80 int32 = -326
                                                                                            _ = v80
                                                                                        default:
                                                                                            fmt.Println(v2, v66, v66)
                                                                                        }
                                                                                    }
                                                                                    _ = v57
                                                                                    _ = v66
                                                                                    _ = v68
                                                                                } else {
                                                                                    v2 = 42
                                                                                }
                                                                                var v81 uint16 = 42
                                                                                var v82 float64 = float64(float32(((3.14159 / float64(((float32(1.5) + float32(float64(1.5))) * 1e10))) - (0.0 - -2.25))))
                                                                                if true {
                                                                                    var v83 int16 = 42
                                                                                    if false {
                                                                                        var v84 int32 = 649
                                                                                        if (v82 != (v82 + v82)) {
                                                                                            if true {
                                                                                                v85 := GenS12{F0: 255}
                                                                                                _ = v85
                                                                                            } else {
                                                                                                var v86 uint16 = 608
                                                                                                v87 := GenS1{F0: 127, F1: (-(v2)), F2: 176}
                                                                                                _ = v86
                                                                                                _ = v87
                                                                                            }
                                                                                            if false {
                                                                                                fmt.Println(v53)
                                                                                                if len(v1) > 0 {
                                                                                                    v1[0] = "a"
                                                                                                }
                                                                                            } else {
                                                                                                var v88 uint64 = 221
                                                                                                fmt.Println(v84, fi22, fi33)
                                                                                                v84 = -597
                                                                                                _ = v88
                                                                                            }
                                                                                        }
                                                                                        _ = len(v1)
                                                                                        _ = v84
                                                                                    } else {
                                                                                        v89 := GenS3{F0: 2, F1: (^(v81))}
                                                                                        v81 += v81
                                                                                        v90 := GenS13{F0: false, F1: (v53 % ((600) | 1)), F2: (fi33 * -2147483648)}
                                                                                        _ = v89
                                                                                        _ = v90
                                                                                    }
                                                                                    v81 += v81
                                                                                    _ = v83
                                                                                }
                                                                                _ = v81
                                                                                _ = v82
                                                                            default:
                                                                                v91 := GenS6{F0: 1}
                                                                                _ = v91
                                                                            }
                                                                            _ = v55
                                                                        }
                                                                        fmt.Println(fi33, v2)
                                                                    } else {
                                                                        v2 = (v2 & v2)
                                                                        v92 := GenS12{F0: 0}
                                                                        var v93 int32 = 1
                                                                        _ = v92
                                                                        _ = v93
                                                                    }
                                                                }
                                                                fmt.Println(fi33)
                                                            }
                                                            var v94 float64 = 0.1
                                                            _ = v53
                                                            _ = v94
                                                        case 2147483647:
                                                            v2 += 0
                                                        case 9223372036854775807:
                                                            switch "\n\t" {
                                                            case "a":
                                                                switch (fi33 | fi22) {
                                                                case 1:
                                                                    switch fi22 {
                                                                    case -2147483648:
                                                                        if false {
                                                                            fmt.Println(fi22, fi33, fi22)
                                                                            switch fi33 {
                                                                            case -19:
                                                                                v2 -= (v2 * v2)
                                                                            }
                                                                        }
                                                                    default:
                                                                        var v95 uint8 = 33
                                                                        v95 = v95
                                                                        _ = v95
                                                                    }
                                                                case 0:
                                                                    var v96 uint16 = 82
                                                                    _ = v96
                                                                case -2147483648:
                                                                    delete(v3, (fi33 % ((32767) | 1)))
                                                                default:
                                                                    switch fi33 {
                                                                    case 127:
                                                                        var v97 uint64 = 1
                                                                        _ = v97
                                                                    case 32768:
                                                                        v98 := GenS1{F0: (v2 / ((-7) | 1)), F1: (v2 % ((v2) | 1)), F2: 2147483647}
                                                                        _ = v98
                                                                    default:
                                                                        if true {
                                                                            var v99 int64 = -9223372036854775808
                                                                            if true {
                                                                                v100 := GenS11{F0: (v2 * 127), F1: "hello", F2: 2147483647}
                                                                                var v101 uint64 = 42
                                                                                var v102 int32 = 42
                                                                                _ = v100
                                                                                _ = v101
                                                                                _ = v102
                                                                            } else {
                                                                                v103 := GenS7{F0: (v2 - (v2 * v2)), F1: 42}
                                                                                var v104 string = ""
                                                                                v105 := GenS9{F0: (v104 + v104), F1: v2, F2: 2147483647}
                                                                                _ = v103
                                                                                _ = v104
                                                                                _ = v105
                                                                            }
                                                                            _ = v99
                                                                        } else {
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                            case "":
                                                                delete(v3, (fi22 + (fi22 + (fi22 / (((fi33 / ((fi22) | 1))) | 1)))))
                                                            default:
                                                                var v106 float32 = (float32(-2.25) - float32((3.14159 - ((-2.25 - 0.1) / (1e10 / 0.1)))))
                                                                _ = v106
                                                            }
                                                        default:
                                                            if len(v1) > 0 {
                                                                v1[0] = "\n\t"
                                                            }
                                                        }
                                                        if false {
                                                            var v107 int16 = -256
                                                            var v108 int = (fi33 << ((0) & 63))
                                                            switch (^(fi22)) {
                                                            case 128:
                                                                v109 := GenS5{F0: 42, F1: 939}
                                                                _ = v109
                                                            case -303:
                                                                var v110 int64 = 1
                                                                _ = v110
                                                            }
                                                            _ = v107
                                                            _ = v108
                                                        } else {
                                                            var v111 float32 = 0.0
                                                            _ = v111
                                                        }
                                                    }
                                                    _ = fi33
                                                }
                                                v112 := GenS14{F0: 1, F1: 42}
                                                v113 := GenS15{F0: 0, F1: ((float32(0.1) + 0.0) <= float32(float64(float32(-2.25)))), F2: 280}
                                                _ = v112
                                                _ = v113
                                            }
                                            fmt.Println(fi22)
                                        }
                                        _ = v29
                                    } else {
                                        v2 = v2
                                    }
                                } else {
                                    v2 += (v2 * v2)
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    var v114 int64 = -751
                                    _ = v114
                                }
                                _ = fi22
                                _ = v23
                                _ = v24
                            }
                        case -24:
                            for fi115 := 0; fi115 < 2; fi115++ {
                                for fi116 := 0; fi116 < 12; fi116++ {
                                    if (538 < 1) {
                                        v117 := GenS16{F0: 850}
                                        _ = v117
                                    } else {
                                        v2 = 0
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        if false {
                                            fmt.Println(v2, fi115, fi116)
                                            v118 := v3[fi116]
                                            _ = v118
                                        } else {
                                            if (!true) {
                                                fmt.Println(fi116, v2)
                                            } else {
                                                v119 := GenS17{F0: 896}
                                                v2 = (v2 ^ (v2 ^ (-(v2))))
                                                _ = v119
                                            }
                                        }
                                        if (true && (!(!(2147483647 == 1)))) {
                                            var v120 uint32 = 42
                                            var v121 int = 9223372036854775807
                                            _ = v120
                                            _ = v121
                                        } else {
                                            v122 := GenS16{F0: 0}
                                            switch "" {
                                            case "\n\t":
                                                if (42 >= 103) {
                                                    var v123 int64 = -9223372036854775808
                                                    _ = len(v3)
                                                    v124 := GenS18{F0: fi115, F1: -32768, F2: 2147483647}
                                                    _ = v123
                                                    _ = v124
                                                } else {
                                                    fmt.Println(v2, fi116, v2)
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    var v125 int8 = (v2 ^ 1)
                                                    _ = v125
                                                }
                                                _ = len(v3)
                                            case "":
                                                var v126 int32 = 518
                                                _ = v126
                                            }
                                            _ = v122
                                        }
                                    }
                                    v127 := v3[fi115]
                                    _ = fi116
                                    _ = v127
                                }
                                _ = fi115
                            }
                        case 128:
                            delete(v3, 4294967295)
                        }
                        if false {
                            for fi128 := 0; fi128 < 7; fi128++ {
                                var v129 float32 = 0.0
                                for fi130 := 0; fi130 < 11; fi130++ {
                                    var v131 int = fi130
                                    _ = fi130
                                    _ = v131
                                }
                                fmt.Println(v2, v2, v2)
                                _ = fi128
                                _ = v129
                            }
                            for fi132 := 0; fi132 < 13; fi132++ {
                                if false {
                                    v2 += v2
                                } else {
                                    var v133 float32 = 0.0
                                    var v134 string = "\n\t"
                                    _ = v133
                                    _ = v134
                                }
                                if false {
                                    fmt.Println(fi132, fi132)
                                    var v135 int32 = 0
                                    _ = v135
                                }
                                var v136 int16 = 42
                                _ = fi132
                                _ = v136
                            }
                        }
                    }
                } else {
                    if (-9223372036854775808 >= 0) {
                        if (0.1 <= 0.1) {
                            v137 := GenS16{F0: -1}
                            if (3.14159 != float32(0.1)) {
                                var v138 uint64 = 2
                                v138 -= (v138 % ((v138) | 1))
                                _ = v138
                            }
                            _ = v137
                        } else {
                            for fi139 := 0; fi139 < 7; fi139++ {
                                v140 := GenS12{F0: 306}
                                _ = fi139
                                _ = v140
                            }
                            fmt.Println(v2)
                            for fi141 := 0; fi141 < 12; fi141++ {
                                v2 = (-(v2))
                                var v142 bool = true
                                _ = fi141
                                _ = v142
                            }
                        }
                        v2 = 46
                        var v143 int16 = 42
                        _ = v143
                    } else {
                        fmt.Println(v2)
                    }
                }
                v144 := GenS19{F0: 691}
                if true {
                    v145 := GenS20{F0: (v2 % ((v2) | 1)), F1: (-(v2)), F2: 42}
                    _ = v145
                } else {
                    var v146 float64 = 1e10
                    v146 = 1e10
                    _ = v146
                }
                _ = v144
            }
        }
    case "\n\t":
        switch "foo bar" {
        case "":
            switch 1 {
            case 533:
                for fi147 := 0; fi147 < 16; fi147++ {
                    for fi148 := 0; fi148 < 19; fi148++ {
                        fmt.Println(v2)
                        if true {
                            if (!(-128 > v2)) {
                                v149 := GenS6{F0: 584}
                                var v150 uint32 = 2147483647
                                _ = len(v1)
                                _ = v149
                                _ = v150
                            } else {
                                if true {
                                    switch (fi147 >> ((18446744073709551615) & 63)) {
                                    case 127:
                                        v151 := v3[fi147]
                                        _ = v151
                                    default:
                                        v2 = v2
                                    }
                                    v152 := GenS1{F0: v2, F1: 42, F2: 770}
                                    v2 -= v2
                                    _ = v152
                                } else {
                                    var v153 int = (fi148 & 127)
                                    _ = v153
                                }
                                v154 := v3[(fi147 % (((fi147 / ((fi148) | 1))) | 1))]
                                if (!false) {
                                    fmt.Println(v154, fi148)
                                    v1 = append(v1, "\n\t")
                                }
                                _ = v154
                            }
                            var v155 uint16 = 255
                            _ = len(v3)
                            _ = v155
                        } else {
                            fmt.Println(fi147)
                        }
                        _ = fi148
                    }
                    var v156 int8 = v2
                    v2 = (v156 / ((v156) | 1))
                    _ = fi147
                    _ = v156
                }
            }
        case "foo bar":
            v2 = -75
        }
    }
    v157 := GenS13{F0: (false || true), F1: 0, F2: -105}
    v1 = append(v1, "a")
    v2 -= (^(v2))
    v2 += 127
    switch 0 {
    case 127:
        for fi158 := 0; fi158 < 20; fi158++ {
            if true {
                if false {
                    fmt.Println(v2, fi158)
                    _ = len(v1)
                    var v159 uint16 = 255
                    _ = v159
                } else {
                    var v160 int8 = v2
                    if (true || false) {
                        v3[(fi158 % (((fi158 >> ((702) & 63))) | 1))] = ""
                        var v161 string = ("hello" + "foo bar")
                        _ = v161
                    } else {
                        if len(v1) > 0 {
                            v1[0] = "hello"
                        }
                    }
                    _ = v160
                }
            } else {
                v162 := GenS14{F0: -1, F1: 2147483648}
                var v163 uint16 = 255
                var v164 uint64 = 42
                _ = v162
                _ = v163
                _ = v164
            }
            switch "\n\t" {
            case "":
                var v165 int64 = -1
                _ = v165
            default:
                for fi166 := 0; fi166 < 3; fi166++ {
                    v2 -= 127
                    var v167 uint16 = 0
                    _ = fi166
                    _ = v167
                }
            }
            _ = fi158
        }
    case 128:
        v168 := GenS21{F0: 42, F1: 0, F2: 586}
        _ = v168
    case -9223372036854775808:
        for fi169 := 0; fi169 < 18; fi169++ {
            v170 := GenS5{F0: -9223372036854775808, F1: 0}
            v3[fi169] = ("foo bar" + "a")
            _ = fi169
            _ = v170
        }
    default:
        v2 -= v2
    }
    v2 += -128
    delete(v3, -9223372036854775808)
    _ = len(v3)
    var v171 int16 = -1
    var v172 bool = false
    switch (("\n\t" + "foo bar") + "\n\t") {
    case "a":
        switch "hello" {
        case "\n\t":
            switch "" {
            case "a":
                switch "foo bar" {
                case "foo bar":
                    for fi173 := 0; fi173 < 20; fi173++ {
                        for fi174 := 0; fi174 < 12; fi174++ {
                            switch fi174 {
                            case 32768:
                                v175 := GenS22{F0: "\n\t", F1: ((v2 << ((255) & 63)) == v2), F2: fi174}
                                _ = v175
                            case 53:
                                fmt.Println(v172)
                            }
                            switch fi174 {
                            case -594:
                                fmt.Println(fi173, fi173, v172)
                            }
                            _ = fi174
                        }
                        if (!true) {
                            v176 := GenS8{F0: -9223372036854775808, F1: 65535}
                            _ = v176
                        } else {
                            fmt.Println(v172, fi173)
                            v177 := GenS10{F0: 0, F1: 0}
                            if v172 {
                                for fi178 := 0; fi178 < 8; fi178++ {
                                    v2 += int8(fi178)
                                    v179 := GenS23{F0: 41, F1: v2}
                                    _ = fi178
                                    _ = v179
                                }
                                if (158 == 1) {
                                    var v180 bool = v172
                                    for fi181 := 0; fi181 < 17; fi181++ {
                                        v171 = 1
                                        _ = fi181
                                    }
                                    _ = v180
                                }
                                for fi182 := 0; fi182 < 3; fi182++ {
                                    if false {
                                        v171 = (^(v171))
                                        v171 += v171
                                    }
                                    var v183 float32 = 1e10
                                    switch (fi182 | (^(fi182))) {
                                    case 32767:
                                        v184 := GenS19{F0: -304}
                                        _ = v184
                                    case -2147483648:
                                        var v185 uint = 1
                                        _ = v185
                                    case 2147483647:
                                        v186 := v3[(fi173 * fi173)]
                                        _ = v186
                                    }
                                    _ = fi182
                                    _ = v183
                                }
                            } else {
                                switch "\n\t" {
                                case "a":
                                    delete(v3, (fi173 >> ((844) & 63)))
                                case "foo bar":
                                    var v187 string = ""
                                    _ = v187
                                default:
                                    var v188 int8 = v2
                                    _ = v188
                                }
                                if v172 {
                                    fmt.Println(v2, v171, v172)
                                    v172 = (v172 || true)
                                    if (v172 || false) {
                                        v189 := GenS24{F0: 706}
                                        _ = v189
                                    }
                                }
                            }
                            _ = v177
                        }
                        fmt.Println(v171)
                        _ = fi173
                    }
                case "hello":
                    for fi190 := 0; fi190 < 14; fi190++ {
                        var v191 int16 = 0
                        _ = len(v3)
                        for fi192 := 0; fi192 < 12; fi192++ {
                            _ = len(v3)
                            v172 = true
                            _ = fi192
                        }
                        _ = fi190
                        _ = v191
                    }
                case "":
                    switch ("" + "foo bar") {
                    case "a":
                        var v193 string = ("" + "")
                        _ = v193
                    case "hello":
                        delete(v3, 42)
                    }
                }
            case "":
                v2 -= (-(v2))
            }
        case "":
            v194 := GenS17{F0: v171}
            _ = v194
        }
    default:
        v1 = append(v1, "foo bar")
    }
    fmt.Println(v2)
    v195 := GenS9{F0: "a", F1: (v2 ^ 1), F2: 2147483648}
    v171 = v171
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v171:", v171)
    fmt.Println("v172:", v172)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v157
    _ = v171
    _ = v172
    _ = v195
    _ = v2
    _ = v3
}
