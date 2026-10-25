package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 int8; F1 string; F2 int64 }
type GenS4 struct { F0 uint32; F1 int32 }
type GenS5 struct { F0 uint; F1 int8 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint64; F1 int64; F2 uint64 }
type GenS8 struct { F0 uint32; F1 uint }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint8 }

func main() {
    var v1 int = -2147483648
    var v2 uint = 255
    var v3 int32 = -285
    v4 := make(map[string]int)
    var v5 int32 = (v3 << ((1) & 63))
    var v6 uint8 = 113
    for fi7 := 0; fi7 < 5; fi7++ {
        switch "\n\t" {
        case "foo bar":
            var v8 uint16 = 255
            _ = v8
        case "hello":
            var v9 uint = v2
            _ = v9
        case "a":
            for fi10 := 0; fi10 < 1; fi10++ {
                v5 += -1
                _ = fi10
            }
        default:
            v11 := GenS1{F0: 32767}
            _ = v11
        }
        _ = fi7
    }
    switch v1 {
    case 127:
        var v12 uint = 862
        _ = v12
    }
    var v13 float64 = ((0.0 / 0.1) * 0.0)
    switch ("hello" + ("\n\t" + "foo bar")) {
    case "":
        for fi14 := 0; fi14 < 8; fi14++ {
            for fi15 := 0; fi15 < 10; fi15++ {
                var v16 int32 = (-(v3))
                if true {
                    var v17 bool = (true && false)
                    _ = v17
                } else {
                    var v18 float32 = ((-2.25 / 1e10) - ((1.5 - 3.14159) * 0.1))
                    v2 = (v2 | 4294967295)
                    _ = v18
                }
                if false {
                    fmt.Println(fi15)
                }
                _ = fi15
                _ = v16
            }
            _ = fi14
        }
    case "hello":
        v19 := GenS1{F0: 1}
        _ = v19
    }
    v3 = -873
    v20 := GenS2{F0: 999, F1: 610}
    var v21 uint32 = 4294967295
    if ((!false) && (false && false)) {
        if true {
            var v22 int8 = 1
            for fi23 := 0; fi23 < 16; fi23++ {
                delete(v4, "hello")
                v4[""] = v1
                var v24 int8 = v22
                _ = fi23
                _ = v24
            }
            _ = v22
        }
        for fi25 := 0; fi25 < 13; fi25++ {
            for fi26 := 0; fi26 < 3; fi26++ {
                if false {
                    _ = len(v4)
                    v27 := GenS3{F0: -1, F1: "a", F2: -9223372036854775808}
                    var v28 int32 = -2147483648
                    if (0.0 <= 0.1) {
                        var v29 uint64 = 251
                        _ = v29
                    }
                    _ = v27
                    _ = v28
                } else {
                    switch "\n\t" {
                    case "hello":
                        var v30 uint16 = uint16(v2)
                        _ = v30
                    case "\n\t":
                        if true {
                            if false {
                                v1 += v1
                                v31 := GenS4{F0: v21, F1: (v5 * v3)}
                                _ = v31
                            }
                        } else {
                            var v32 int64 = int64(v5)
                            _ = v32
                        }
                    default:
                        var v33 uint64 = 146
                        _ = v33
                    }
                    if (v5 > v3) {
                        v34 := GenS5{F0: v2, F1: -31}
                        if (!true) {
                            v35 := []int{(-(fi26)), (fi26 << (((v2 ^ uint(v6))) & 63))}
                            if false {
                                var v36 uint = v2
                                var v37 int16 = 32767
                                var v38 uint = v36
                                _ = v36
                                _ = v37
                                _ = v38
                            } else {
                                v39 := GenS5{F0: (v2 & (v2 - 617)), F1: -104}
                                if true {
                                    if false {
                                        switch "" {
                                        case "a":
                                            fmt.Println(v21, v21)
                                        case "hello":
                                            v13 = (((float64(1.5) * v13) * (3.14159 / 1.0)) / (v13 - ((0.1 - (((v13 - v13) - (float64((3.14159 - 1e10)) + -2.25)) + 1e10)) - (v13 * v13))))
                                        case "":
                                            v40 := GenS6{F0: v5}
                                            _ = v40
                                        default:
                                            if false {
                                                if false {
                                                    switch (("\n\t" + "hello") + "") {
                                                    case "\n\t":
                                                        v41 := v4["foo bar"]
                                                        _ = v41
                                                    case "a":
                                                        v13 -= float64(1e10)
                                                    }
                                                    fmt.Println(v13)
                                                } else {
                                                    if (!true) {
                                                        if ((v21 != v21) && false) {
                                                            v42 := GenS6{F0: v5}
                                                            v5 += v3
                                                            switch fi25 {
                                                            case 446:
                                                                if ((v2 * v2) == (v2 - (^(v2)))) {
                                                                    var v43 int8 = -64
                                                                    var v44 uint32 = v21
                                                                    _ = v43
                                                                    _ = v44
                                                                } else {
                                                                    v2 -= v2
                                                                    fmt.Println(fi25, v1, v3)
                                                                    var v45 int16 = 134
                                                                    _ = v45
                                                                }
                                                            case 32768:
                                                                fmt.Println(v21)
                                                            case 1:
                                                                v3 += int32(v21)
                                                            default:
                                                                var v46 int32 = (v5 << (((v2 + 42)) & 63))
                                                                _ = v46
                                                            }
                                                            _ = v42
                                                        } else {
                                                            v6 -= 180
                                                            if len(v35) >= 2 {
                                                                v35 = v35[1:len(v35)]
                                                            }
                                                            var v47 float64 = v13
                                                            _ = v47
                                                        }
                                                        v48 := GenS7{F0: 0, F1: -77, F2: 0}
                                                        _ = v48
                                                    } else {
                                                        v2 = v2
                                                        fmt.Println(v3, v1)
                                                    }
                                                    var v49 float32 = ((float32(-2.25) - (float32(v13) * ((0.1 / -2.25) / (0.0 - (-2.25 * float32(v13)))))) - (3.14159 / (((1.5 * ((((float32(v13) - 1e10) / 3.14159) + -2.25) - ((3.14159 / ((-2.25 - float32(v13)) + (1.5 / (float32(float64((1.5 / 3.14159))) * float32(v13))))) / -2.25))) + -2.25) + -2.25)))
                                                    _ = v49
                                                }
                                                fmt.Println(v2)
                                                switch "hello" {
                                                case "foo bar":
                                                    fmt.Println(fi25, fi26)
                                                default:
                                                    fmt.Println(fi25, v3, v21)
                                                }
                                                var v50 float32 = 3.14159
                                                _ = v50
                                            }
                                        }
                                        v1 += v1
                                        var v51 uint64 = 4294967295
                                        _ = v51
                                    } else {
                                        v52 := v4["\n\t"]
                                        fmt.Println(v21, v21)
                                        switch fi25 {
                                        case 128:
                                            if true {
                                                v4["hello"] = (v1 ^ fi25)
                                            }
                                        case 0:
                                            if true {
                                                fmt.Println(v2)
                                            }
                                        case 42:
                                            v13 = v13
                                        default:
                                            _ = len(v35)
                                        }
                                        _ = v52
                                    }
                                }
                                _ = v39
                            }
                            v3 += v3
                            if (18446744073709551615 != 0) {
                                v1 = fi25
                                if true {
                                    if ((v1 < 616) || false) {
                                        v5 -= v3
                                        if (1 <= 32767) {
                                            v53 := GenS6{F0: (v5 << (((v2 % ((v2) | 1))) & 63))}
                                            v1 -= fi26
                                            _ = v53
                                        } else {
                                            v6 = v6
                                        }
                                        var v54 uint64 = 861
                                        _ = v54
                                    } else {
                                        if false {
                                            var v55 float32 = -2.25
                                            v4[(("a" + "\n\t") + ("" + "a"))] = (v1 & (v1 | fi25))
                                            _ = len(v4)
                                            _ = len(v35)
                                            _ = v55
                                        }
                                        v56 := GenS8{F0: v21, F1: v2}
                                        _ = v56
                                    }
                                } else {
                                    v21 -= (^(v21))
                                    if len(v35) >= 2 {
                                        v35 = v35[1:len(v35)]
                                    }
                                }
                            } else {
                                _ = len(v4)
                                var v57 int8 = 100
                                _ = v57
                            }
                            _ = v35
                        }
                        if true {
                            v5 = v3
                            if true {
                                v3 = v5
                                v58 := GenS2{F0: v21, F1: 4294967295}
                                var v59 uint8 = (v6 ^ 128)
                                switch 32768 {
                                case -2147483648:
                                    if true {
                                        v60 := []int{}
                                        v61 := GenS3{F0: -71, F1: "foo bar", F2: -9223372036854775808}
                                        if (v21 == 2147483647) {
                                            v62 := GenS9{F0: 1}
                                            v5 -= int32(v2)
                                            v3 += v3
                                            var v63 uint64 = 2
                                            var v64 int8 = 31
                                            _ = v62
                                            _ = v63
                                            _ = v64
                                        }
                                        _ = v60
                                        _ = v61
                                    }
                                case 4294967295:
                                    fmt.Println(fi25, v1, fi25)
                                    v2 = uint(v21)
                                    v59 += (v6 >> ((v2) & 63))
                                default:
                                    fmt.Println(v21, v21, v59)
                                }
                                if false {
                                    v65 := []int{(fi26 % ((127) | 1)), fi25, 2147483647, -440}
                                    _ = v65
                                } else {
                                    var v66 uint16 = 42
                                    v67 := []int{}
                                    _ = v66
                                    _ = v67
                                }
                                _ = v58
                                _ = v59
                            } else {
                                fmt.Println(v6)
                                v68 := []int{(^(fi25)), (v1 + fi26), (fi25 * 1)}
                                _ = v68
                            }
                            if (0.0 < ((v13 - v13) / 1.5)) {
                                if false {
                                    if true {
                                        fmt.Println(v13, v3, v6)
                                        v69 := []int{32767, int(v5), fi26}
                                        v70 := GenS10{F0: int(v6)}
                                        _ = v69
                                        _ = v70
                                    }
                                    fmt.Println(v5)
                                    v13 = v13
                                } else {
                                    fmt.Println(v3, v21, v3)
                                }
                            } else {
                                if false {
                                    var v71 int = 169
                                    v21 -= uint32(v2)
                                    var v72 int8 = 1
                                    _ = v71
                                    _ = v72
                                }
                                v21 = v21
                                v73 := []int{fi25, (fi26 / ((int(v5)) | 1))}
                                _ = v73
                            }
                        } else {
                            var v74 uint64 = 2
                            fmt.Println(fi26)
                            v21 = (v21 + 2147483648)
                            delete(v4, "foo bar")
                            _ = v74
                        }
                        _ = v34
                    } else {
                        if (-1 == 42) {
                            var v75 int8 = 1
                            _ = v75
                        }
                        fmt.Println(v2, v13, v6)
                        var v76 float64 = v13
                        _ = v76
                    }
                }
                _ = fi26
            }
            for fi77 := 0; fi77 < 4; fi77++ {
                fmt.Println(v2, v2)
                var v78 int32 = 1
                _ = fi77
                _ = v78
            }
            for fi79 := 0; fi79 < 6; fi79++ {
                v80 := GenS7{F0: 774, F1: 0, F2: 2}
                if true {
                    v81 := []int{-352, v1, (fi25 + (fi25 | 0)), v1}
                    var v82 int = (^(fi25))
                    _ = v81
                    _ = v82
                } else {
                    v1 += fi79
                }
                v83 := GenS7{F0: 2, F1: 9223372036854775807, F2: 2}
                _ = len(v4)
                _ = fi79
                _ = v80
                _ = v83
            }
            _ = fi25
        }
        switch v1 {
        case 741:
            for fi84 := 0; fi84 < 2; fi84++ {
                switch fi84 {
                case -9223372036854775808:
                    for fi85 := 0; fi85 < 1; fi85++ {
                        var v86 int32 = (^(v5))
                        _ = fi85
                        _ = v86
                    }
                default:
                    v87 := []int{(fi84 | (fi84 >> (((v2 % (((^(v2))) | 1))) & 63))), (fi84 / ((469) | 1))}
                    _ = v87
                }
                v88 := GenS1{F0: 596}
                fmt.Println(v6, v2)
                _ = fi84
                _ = v88
            }
        case 193:
            for fi89 := 0; fi89 < 7; fi89++ {
                fmt.Println(v21)
                for fi90 := 0; fi90 < 13; fi90++ {
                    var v91 uint = (v2 << (((v2 | 42)) & 63))
                    v92 := GenS1{F0: 42}
                    v93 := []int{int(v3)}
                    _ = fi90
                    _ = v91
                    _ = v92
                    _ = v93
                }
                v94 := []int{-595, (^(v1)), fi89}
                _ = fi89
                _ = v94
            }
        default:
            var v95 uint16 = 555
            _ = v95
        }
    } else {
        v3 = (v5 | -865)
        v96 := []int{}
        _ = v96
    }
    var v97 uint32 = (v21 << ((1) & 63))
    for fi98 := 0; fi98 < 16; fi98++ {
        for fi99 := 0; fi99 < 8; fi99++ {
            switch "" {
            case "a":
                var v100 int8 = 127
                _ = v100
            default:
                var v101 bool = false
                _ = v101
            }
            var v102 uint64 = 42
            var v103 uint8 = 255
            v4[""] = 350
            var v104 uint16 = 436
            _ = fi99
            _ = v102
            _ = v103
            _ = v104
        }
        _ = fi98
    }
    for fi105 := 0; fi105 < 7; fi105++ {
        for fi106 := 0; fi106 < 4; fi106++ {
            v107 := v4[""]
            _ = fi106
            _ = v107
        }
        switch "a" {
        case "hello":
            v21 -= (v97 << ((584) & 63))
        case "foo bar":
            v108 := GenS6{F0: (v5 ^ (v5 << ((v2) & 63)))}
            _ = v108
        }
        for fi109 := 0; fi109 < 18; fi109++ {
            if false {
                if (1 < v6) {
                    if true {
                        v110 := []int{}
                        v5 = (v3 / (((^(v5))) | 1))
                        if (1000 <= 0) {
                            switch ("foo bar" + "") {
                            case "":
                                var v111 uint = 458
                                _ = v111
                            case "\n\t":
                                var v112 uint16 = uint16(v2)
                                _ = v112
                            default:
                                v113 := GenS11{F0: 50}
                                v21 -= v97
                                _ = v113
                            }
                            var v114 bool = (!(127 < 127))
                            _ = v114
                        }
                        fmt.Println(fi105, v21, v13)
                        if false {
                            v115 := GenS6{F0: (v3 >> ((v2) & 63))}
                            var v116 uint32 = v97
                            v21 -= v116
                            v6 -= (^(v6))
                            _ = v115
                            _ = v116
                        } else {
                            v2 = v2
                            v2 += uint(v6)
                        }
                        _ = v110
                    } else {
                        v117 := GenS6{F0: (^(v5))}
                        _ = v117
                    }
                    fmt.Println(v13)
                    v118 := []int{2147483647}
                    _ = v118
                } else {
                    if false {
                        if true {
                            var v119 float32 = (3.14159 - 3.14159)
                            _ = v119
                        }
                        v5 += 42
                        var v120 int8 = 127
                        _ = v120
                    }
                    v121 := []int{(fi105 * (^(fi105)))}
                    _ = v121
                }
                var v122 int8 = 11
                _ = v122
            } else {
                var v123 int = fi109
                switch (-(fi105)) {
                case 4294967295:
                    var v124 uint = v2
                    _ = v124
                case 0:
                    if true {
                        v5 -= v5
                        fmt.Println(v123, fi109)
                    }
                case 9223372036854775807:
                    v125 := []int{1, (v123 << (((^(v2))) & 63))}
                    _ = v125
                }
                v126 := GenS10{F0: -499}
                _ = v123
                _ = v126
            }
            var v127 string = (("a" + ((("hello" + "hello") + "") + "a")) + "foo bar")
            v13 = 0.1
            _ = fi109
            _ = v127
        }
        _ = fi105
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v97
}
