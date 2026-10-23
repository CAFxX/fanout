package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 bool; F2 uint32 }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 string; F1 int64; F2 int16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 bool; F1 string }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int32; F1 int32; F2 bool }
type GenS9 struct { F0 int8; F1 uint32 }
type GenS10 struct { F0 int8; F1 int8 }
type GenS11 struct { F0 uint8; F1 string; F2 uint }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 int64; F1 string; F2 int8 }
type GenS14 struct { F0 int64; F1 uint16 }

func main() {
    var v1 int16 = -920
    var v2 uint16 = 32767
    var v3 float32 = ((3.14159 + 1.5) + 1e10)
    var v4 bool = (0.0 != 0.0)
    var v5 int64 = 385
    var v6 int16 = (v1 << ((409) & 63))
    v5 = -9223372036854775808
    var v7 int64 = v5
    for fi8 := 0; fi8 < 19; fi8++ {
        fmt.Println(v5)
        _ = fi8
    }
    v9 := []int{128}
    switch 32768 {
    case 32768:
        v7 += v7
    }
    v9 = append(v9, -1)
    var v10 uint16 = 778
    fmt.Println(v10, v7, v6)
    _ = len(v9)
    var v11 int = 32768
    if (v1 <= v1) {
        if true {
            fmt.Println(v11, v7)
            switch (("hello" + ("\n\t" + "\n\t")) + "a") {
            case "hello":
                for fi12 := 0; fi12 < 11; fi12++ {
                    var v13 float64 = 3.14159
                    var v14 int8 = -87
                    switch 115 {
                    case -9223372036854775808:
                        if ((v5 << ((701) & 63)) == v5) {
                            var v15 uint32 = 65535
                            v11 = fi12
                            v3 -= 0.1
                            _ = v15
                        }
                    case -311:
                        if v4 {
                            if true {
                                if v4 {
                                    if false {
                                        var v16 uint32 = 65535
                                        var v17 int8 = (v14 << ((18446744073709551615) & 63))
                                        _ = v16
                                        _ = v17
                                    }
                                }
                            } else {
                                var v18 int32 = 2147483647
                                _ = v18
                            }
                            if (2 <= 2) {
                                if len(v9) >= 2 {
                                    v9 = v9[1:len(v9)]
                                }
                            } else {
                                if v4 {
                                    var v19 string = "a"
                                    _ = v19
                                } else {
                                    if (42 > 42) {
                                        v20 := GenS1{F0: 2147483648}
                                        var v21 int8 = (v14 % (((v14 >> ((18446744073709551615) & 63))) | 1))
                                        _ = v20
                                        _ = v21
                                    }
                                    fmt.Println(v13, v7)
                                }
                                for fi22 := 0; fi22 < 2; fi22++ {
                                    v23 := GenS2{F0: ("\n\t" + (("hello" + "hello") + "hello")), F1: (0 <= 4294967295), F2: 2147483647}
                                    fmt.Println(v13)
                                    _ = fi22
                                    _ = v23
                                }
                                v11 = (fi12 << ((4294967295) & 63))
                            }
                        } else {
                            for fi24 := 0; fi24 < 5; fi24++ {
                                v2 += (v2 >> ((42) & 63))
                                _ = fi24
                            }
                            if len(v9) > 0 {
                                v9[0] = 1
                            }
                        }
                    default:
                        var v25 string = "\n\t"
                        _ = v25
                    }
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            default:
                var v26 uint32 = 4294967295
                _ = v26
            }
        }
        switch 9223372036854775807 {
        case 1:
            if (v3 == v3) {
                switch 9223372036854775807 {
                case 9223372036854775807:
                    v27 := GenS2{F0: "a", F1: v4, F2: 2}
                    _ = v27
                case -2147483648:
                    var v28 uint64 = 460
                    _ = v28
                case 1:
                    for fi29 := 0; fi29 < 13; fi29++ {
                        v5 = v5
                        v1 += 1
                        _ = fi29
                    }
                default:
                    fmt.Println(v1)
                }
                _ = len(v9)
            } else {
                var v30 string = "\n\t"
                _ = v30
            }
        default:
            if (true && ((3.14159 + (3.14159 + 3.14159)) > 1.5)) {
                var v31 int64 = 0
                v6 = v1
                var v32 uint32 = 0
                _ = v31
                _ = v32
            } else {
                switch "" {
                case "foo bar":
                    fmt.Println(v2)
                default:
                    v33 := GenS3{F0: -2147483648, F1: ((0.0 / -2.25) <= v3)}
                    _ = v33
                }
            }
        }
    }
    var v34 string = "a"
    v1 += 32767
    for fi35 := 0; fi35 < 8; fi35++ {
        v36 := GenS2{F0: v34, F1: false, F2: 2147483647}
        for fi37 := 0; fi37 < 18; fi37++ {
            v6 = 1
            switch fi37 {
            case 32767:
                if v4 {
                    if v4 {
                        var v38 int32 = 0
                        if len(v9) > 0 {
                            v9[0] = (-(v11))
                        }
                        _ = v38
                    } else {
                        if v4 {
                            fmt.Println(v5)
                            if len(v9) > 0 {
                                v9[0] = 160
                            }
                            switch "" {
                            case "foo bar":
                                if true {
                                    var v39 float32 = v3
                                    var v40 int8 = 43
                                    v41 := GenS1{F0: 0}
                                    _ = v39
                                    _ = v40
                                    _ = v41
                                } else {
                                    v42 := GenS3{F0: 1, F1: v4}
                                    var v43 float32 = v3
                                    _ = v42
                                    _ = v43
                                }
                            case "a":
                                if (v4 || v4) {
                                    if ((!v4) || v4) {
                                        var v44 int32 = 42
                                        _ = v44
                                    } else {
                                        if v4 {
                                            var v45 float32 = v3
                                            var v46 uint16 = (^(v10))
                                            switch (v34 + ("foo bar" + v34)) {
                                            case "a":
                                                if (0 >= 1) {
                                                    var v47 bool = true
                                                    switch (((v34 + v34) + ((v34 + "\n\t") + v34)) + v34) {
                                                    case "\n\t":
                                                        v48 := GenS1{F0: 921}
                                                        _ = v48
                                                    }
                                                    var v49 int64 = -893
                                                    _ = v47
                                                    _ = v49
                                                } else {
                                                    var v50 float64 = 1.5
                                                    _ = v50
                                                }
                                            case "foo bar":
                                                var v51 int64 = v5
                                                _ = v51
                                            case "hello":
                                                fmt.Println(fi35, v4, v34)
                                            }
                                            _ = v45
                                            _ = v46
                                        }
                                    }
                                    fmt.Println(v3)
                                    v52 := GenS2{F0: "hello", F1: true, F2: 65535}
                                    _ = v52
                                } else {
                                    var v53 float32 = ((float32(3.14159) * ((((v3 * 1.5) + ((v3 * 3.14159) / (float32(-2.25) / -2.25))) - v3) - 3.14159)) + 0.0)
                                    _ = v53
                                }
                                if len(v9) > 0 {
                                    v9[0] = 9223372036854775807
                                }
                                var v54 bool = true
                                _ = v54
                            default:
                                var v55 string = v34
                                _ = v55
                            }
                        }
                        var v56 uint16 = 0
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                        _ = v56
                    }
                    if v4 {
                        var v57 int32 = 42
                        _ = v57
                    }
                    if (173 > 2) {
                        if (-897 < -648) {
                            var v58 uint8 = 128
                            _ = v58
                        } else {
                            var v59 uint32 = 2147483647
                            _ = v59
                        }
                    }
                }
            case 1:
                v6 = v6
            default:
                switch (v34 + v34) {
                case "\n\t":
                    var v60 uint = 65535
                    _ = v60
                case "hello":
                    if (v4 && false) {
                        if v4 {
                            var v61 string = ((v34 + "foo bar") + v34)
                            v11 = fi37
                            var v62 int16 = (v1 + v1)
                            switch (v61 + "a") {
                            case "hello":
                                var v63 int64 = v7
                                _ = v63
                            default:
                                switch "foo bar" {
                                case "":
                                    if v4 {
                                        fmt.Println(v1, v11)
                                        var v64 uint32 = 640
                                        v61 = "\n\t"
                                        _ = v64
                                    }
                                case "foo bar":
                                    v65 := GenS4{F0: v61, F1: v5, F2: v6}
                                    _ = v65
                                default:
                                    switch -464 {
                                    case 4294967295:
                                        v9 = append(v9, int(v6))
                                    }
                                }
                            }
                            if (-2147483648 < 1) {
                                if len(v9) >= 2 {
                                    v9 = v9[1:len(v9)]
                                }
                                switch v61 {
                                case "":
                                    _ = len(v9)
                                default:
                                    if len(v9) >= 2 {
                                        v9 = v9[1:len(v9)]
                                    }
                                }
                                _ = len(v9)
                            }
                            v66 := GenS4{F0: ((v61 + v34) + v34), F1: (v5 % ((-1) | 1)), F2: (-(v6))}
                            fmt.Println(v5)
                            if len(v9) > 0 {
                                v9[0] = (v11 % ((fi37) | 1))
                            }
                            var v67 int64 = (v7 << ((129) & 63))
                            _ = v61
                            _ = v62
                            _ = v66
                            _ = v67
                        } else {
                            var v68 string = v34
                            fmt.Println(v34, v68)
                            v5 -= v7
                            _ = v68
                        }
                        v6 += -32768
                        if v4 {
                            v4 = true
                            if false {
                                fmt.Println(v34, fi35, v5)
                                switch (fi37 - fi37) {
                                case 32767:
                                    fmt.Println(v5, fi35, v1)
                                default:
                                    var v69 int64 = 0
                                    _ = v69
                                }
                            } else {
                                var v70 uint16 = (v10 ^ (v2 & v10))
                                var v71 int16 = v6
                                _ = v70
                                _ = v71
                            }
                            v9 = append(v9, (fi35 - (fi35 ^ 32767)))
                            if v4 {
                                if (0 > 65535) {
                                    if (!(!(2147483647 <= 922))) {
                                        v72 := GenS1{F0: 2147483647}
                                        v73 := GenS3{F0: 1, F1: true}
                                        if (3.14159 < 1.5) {
                                            v1 += v1
                                            if true {
                                                var v74 uint16 = v2
                                                var v75 uint = 0
                                                _ = v74
                                                _ = v75
                                            } else {
                                                var v76 uint32 = 42
                                                if len(v9) >= 2 {
                                                    v9 = v9[1:len(v9)]
                                                }
                                                v77 := GenS3{F0: -282, F1: (!v4)}
                                                switch fi35 {
                                                case 32767:
                                                    if ((v5 % ((int64(v76)) | 1)) > v5) {
                                                        if (!true) {
                                                            v7 -= (v5 - -210)
                                                        }
                                                        v3 = v3
                                                        switch "\n\t" {
                                                        case "hello":
                                                            switch ("\n\t" + "hello") {
                                                            case "\n\t":
                                                                fmt.Println(v34, v11, v1)
                                                                switch fi37 {
                                                                case -202:
                                                                    v78 := GenS2{F0: v34, F1: (!(0 < 1)), F2: (v76 * v76)}
                                                                    _ = v78
                                                                }
                                                            case "":
                                                                if v4 {
                                                                    v11 = fi37
                                                                    fmt.Println(v5)
                                                                    v3 = float32(-2.25)
                                                                } else {
                                                                    v34 = "a"
                                                                    v79 := GenS5{F0: 126}
                                                                    v4 = (v4 || (!true))
                                                                    _ = v79
                                                                }
                                                            }
                                                        case "":
                                                            if v4 {
                                                                if true {
                                                                    v80 := GenS4{F0: (((((("" + v34) + v34) + v34) + (("\n\t" + ("a" + "a")) + "")) + v34) + v34), F1: (^(v5)), F2: v1}
                                                                    _ = v80
                                                                }
                                                            }
                                                        default:
                                                            _ = len(v9)
                                                        }
                                                    } else {
                                                        var v81 uint8 = 255
                                                        _ = v81
                                                    }
                                                case 113:
                                                    var v82 float32 = v3
                                                    _ = v82
                                                }
                                                _ = v76
                                                _ = v77
                                            }
                                        } else {
                                            var v83 int16 = v1
                                            v3 += (float32(0.1) * v3)
                                            switch v34 {
                                            case "a":
                                                v84 := GenS6{F0: (v34 == "\n\t"), F1: "hello"}
                                                _ = v84
                                            case "":
                                                v85 := GenS6{F0: v4, F1: ((v34 + "hello") + "")}
                                                _ = v85
                                            case "foo bar":
                                                v86 := GenS5{F0: uint8(v2)}
                                                _ = v86
                                            }
                                            _ = v83
                                        }
                                        _ = v72
                                        _ = v73
                                    } else {
                                        if true {
                                            v9 = append(v9, (fi37 * (v11 >> ((255) & 63))))
                                            if v4 {
                                                if len(v9) >= 2 {
                                                    v9 = v9[1:len(v9)]
                                                }
                                                var v87 uint32 = 629
                                                var v88 int = fi35
                                                _ = v87
                                                _ = v88
                                            } else {
                                                v89 := GenS7{F0: 0}
                                                v1 = v1
                                                v9 = append(v9, 32768)
                                                _ = v89
                                            }
                                            var v90 int = v11
                                            _ = v90
                                        } else {
                                            v5 = (^(v5))
                                        }
                                        v91 := GenS8{F0: 98, F1: 2147483647, F2: (100 != -1)}
                                        var v92 int64 = 9223372036854775807
                                        _ = v91
                                        _ = v92
                                    }
                                } else {
                                    v34 = (("a" + v34) + "foo bar")
                                    var v93 uint16 = (^(v10))
                                    _ = v93
                                }
                            } else {
                                v9 = append(v9, fi35)
                                if len(v9) > 0 {
                                    v9[0] = fi35
                                }
                                switch ("hello" + ("" + "\n\t")) {
                                case "hello":
                                    switch fi37 {
                                    case 32768:
                                        if len(v9) >= 2 {
                                            v9 = v9[1:len(v9)]
                                        }
                                    case 1:
                                        v6 = v6
                                    }
                                    v94 := GenS2{F0: ((v34 + ("hello" + v34)) + v34), F1: ((fi35 * (fi35 % ((v11) | 1))) <= v11), F2: 2147483647}
                                    switch (fi35 | v11) {
                                    case -419:
                                        var v95 float32 = v3
                                        _ = v95
                                    default:
                                        if true {
                                            switch "foo bar" {
                                            case "\n\t":
                                                var v96 float64 = -2.25
                                                _ = v96
                                            default:
                                                var v97 int16 = 32767
                                                _ = v97
                                            }
                                            if v4 {
                                                var v98 float64 = 3.14159
                                                v7 -= (v7 * v5)
                                                var v99 uint16 = (v2 - (v10 + 42))
                                                _ = v98
                                                _ = v99
                                            } else {
                                                v11 -= (fi37 >> ((0) & 63))
                                                v100 := GenS9{F0: 42, F1: 112}
                                                v9 = append(v9, 2147483647)
                                                _ = v100
                                            }
                                        } else {
                                            var v101 int16 = (v6 ^ 1000)
                                            v7 -= (v7 - (v7 >> ((42) & 63)))
                                            var v102 float32 = (1e10 * 1.5)
                                            _ = v101
                                            _ = v102
                                        }
                                    }
                                    _ = v94
                                case "":
                                    if (fi37 > (fi35 + v11)) {
                                        switch v34 {
                                        case "\n\t":
                                            var v103 int32 = -310
                                            _ = v103
                                        case "a":
                                            var v104 uint32 = 0
                                            _ = v104
                                        case "foo bar":
                                            switch fi37 {
                                            case -9223372036854775808:
                                                v105 := GenS4{F0: "foo bar", F1: (v7 - v5), F2: v1}
                                                if true {
                                                    fmt.Println(v10)
                                                    if len(v9) > 0 {
                                                        v9[0] = (fi37 / ((-2147483648) | 1))
                                                    }
                                                }
                                                v9 = append(v9, (fi37 / ((fi35) | 1)))
                                                var v106 float32 = v3
                                                var v107 uint32 = 930
                                                _ = v105
                                                _ = v106
                                                _ = v107
                                            case -536:
                                                _ = len(v9)
                                            case 9223372036854775807:
                                                if true {
                                                    var v108 uint = 4294967295
                                                    var v109 int16 = int16(v108)
                                                    v4 = v4
                                                    var v110 uint8 = 0
                                                    v9 = append(v9, fi37)
                                                    _ = v108
                                                    _ = v109
                                                    _ = v110
                                                }
                                            }
                                        }
                                        if (v4 || v4) {
                                            v111 := GenS10{F0: 61, F1: 1}
                                            if v4 {
                                                v112 := GenS11{F0: 127, F1: (((("a" + v34) + v34) + v34) + v34), F2: 42}
                                                _ = v112
                                            } else {
                                                v113 := GenS12{F0: ((v11 + -1) >= (fi35 - (fi37 & v11)))}
                                                var v114 uint = 217
                                                switch 128 {
                                                case -2147483648:
                                                    var v115 uint64 = 0
                                                    _ = v115
                                                case 2147483647:
                                                    v34 = v34
                                                default:
                                                    fmt.Println(v1)
                                                    var v116 uint64 = 773
                                                    switch 32768 {
                                                    case 1:
                                                        if len(v9) > 0 {
                                                            v9[0] = fi35
                                                        }
                                                    default:
                                                        if v4 {
                                                            if len(v9) > 0 {
                                                                v9[0] = fi37
                                                            }
                                                            if v4 {
                                                                v117 := GenS13{F0: int64(v2), F1: ("" + ((v34 + (((v34 + ((v34 + (v34 + v34)) + (v34 + ""))) + "\n\t") + v34)) + v34)), F2: -25}
                                                                _ = len(v9)
                                                                v118 := GenS2{F0: (v34 + (((v34 + v34) + "\n\t") + v34)), F1: true, F2: 1}
                                                                _ = v117
                                                                _ = v118
                                                            }
                                                            if ((-(fi37)) >= fi35) {
                                                                v11 += (fi37 << ((v114) & 63))
                                                            } else {
                                                                if (v1 < (v1 - v1)) {
                                                                    _ = len(v9)
                                                                    fmt.Println(v11, v5, v1)
                                                                } else {
                                                                    var v119 uint16 = (v10 * 189)
                                                                    v3 = v3
                                                                    fmt.Println(fi37)
                                                                    _ = v119
                                                                }
                                                            }
                                                            if len(v9) > 0 {
                                                                v9[0] = (fi35 * 128)
                                                            }
                                                            var v120 uint64 = v116
                                                            _ = v120
                                                        } else {
                                                            v121 := GenS5{F0: 127}
                                                            _ = v121
                                                        }
                                                    }
                                                    _ = v116
                                                }
                                                if ((^(v10)) > (^(v10))) {
                                                    _ = len(v9)
                                                }
                                                v1 = (v6 << ((v114) & 63))
                                                v9 = append(v9, v11)
                                                switch 32768 {
                                                case 566:
                                                    if (!v4) {
                                                        v3 += (v3 + float32(-2.25))
                                                        v6 += v6
                                                    }
                                                case 32768:
                                                    v122 := GenS14{F0: (-(v7)), F1: v2}
                                                    _ = v122
                                                default:
                                                    var v123 int16 = v1
                                                    _ = v123
                                                }
                                                var v124 uint64 = 2
                                                _ = v113
                                                _ = v114
                                                _ = v124
                                            }
                                            v34 = (v34 + v34)
                                            if (((1e10 + 0.0) > -2.25) || (!v4)) {
                                                v125 := GenS2{F0: (v34 + v34), F1: true, F2: 2}
                                                v9 = append(v9, (fi35 % (((fi35 >> ((18446744073709551615) & 63))) | 1)))
                                                _ = v125
                                            } else {
                                                fmt.Println(v3, v6, v5)
                                                var v126 string = "hello"
                                                fmt.Println(v2, v6)
                                                _ = v126
                                            }
                                            _ = v111
                                        } else {
                                            switch (fi37 + fi37) {
                                            case 32767:
                                                v10 += (v2 * (v2 * 665))
                                            case 1:
                                                v3 += float32(-2.25)
                                            }
                                            v7 = -505
                                            v7 -= v7
                                        }
                                    }
                                case "a":
                                    var v127 uint = 134
                                    _ = v127
                                default:
                                    var v128 float32 = (-2.25 * v3)
                                    var v129 uint64 = 2
                                    _ = v128
                                    _ = v129
                                }
                                var v130 string = v34
                                v6 += 140
                                if (!v4) {
                                    v131 := GenS6{F0: (((v11 | -9223372036854775808) == v11) && ((0.0 * float64(v3)) == ((-2.25 * 1e10) / 1.5))), F1: v34}
                                    _ = v131
                                } else {
                                    var v132 uint8 = 1
                                    v2 = 32767
                                    if len(v9) >= 2 {
                                        v9 = v9[1:len(v9)]
                                    }
                                    if ((v6 * 42) != int16(v10)) {
                                        var v133 bool = true
                                        v6 = v1
                                        _ = v133
                                    }
                                    _ = v132
                                }
                                v130 = v130
                                v6 = 42
                                _ = v130
                            }
                        } else {
                            var v134 string = ("hello" + v34)
                            switch "\n\t" {
                            case "foo bar":
                                if (v4 || false) {
                                    var v135 int16 = v6
                                    v136 := GenS2{F0: (v34 + "a"), F1: (((("\n\t" + "\n\t") + ("a" + v34)) + ("foo bar" + "foo bar")) < v134), F2: 65535}
                                    _ = len(v9)
                                    _ = v135
                                    _ = v136
                                }
                            default:
                                var v137 float32 = float32(3.14159)
                                _ = v137
                            }
                            _ = v134
                        }
                        var v138 int64 = (v7 % ((int64(v2)) | 1))
                        _ = v138
                    }
                default:
                    if false {
                        var v139 uint64 = 42
                        _ = v139
                    }
                }
            }
            _ = fi37
        }
        _ = fi35
        _ = v36
    }
    v140 := GenS1{F0: 137}
    switch (^(v11)) {
    case 607:
        v11 = (v11 | (v11 % (((v11 >> ((2) & 63))) | 1)))
    case 4294967295:
        _ = len(v9)
    case 1:
        var v141 int64 = -1
        _ = v141
    default:
        for fi142 := 0; fi142 < 3; fi142++ {
            fmt.Println(v5, v6, v7)
            v143 := GenS11{F0: 128, F1: (v34 + ((v34 + (v34 + (("a" + (((v34 + (v34 + v34)) + v34) + ((v34 + ((v34 + (v34 + v34)) + v34)) + v34))) + v34))) + "foo bar")), F2: 668}
            _ = fi142
            _ = v143
        }
    }
    var v144 uint64 = 136
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v144:", v144)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v140
    _ = v144
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
