package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 int64; F1 int; F2 uint8 }
type GenS4 struct { F0 int16; F1 uint32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint8; F1 int64 }
type GenS7 struct { F0 string; F1 uint64 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 uint8; F1 int16; F2 int64 }
type GenS12 struct { F0 bool; F1 int8 }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 bool }

func main() {
    v1 := make(map[string]int)
    var v2 int64 = 0
    var v3 int8 = 1
    var v4 uint16 = 32767
    switch "foo bar" {
    case "\n\t":
        if (42 == 445) {
            fmt.Println(v2)
            if (0 <= 42) {
                for fi5 := 0; fi5 < 4; fi5++ {
                    fmt.Println(v3, v4)
                    if false {
                        switch ("\n\t" + ("\n\t" + ("" + "hello"))) {
                        case "foo bar":
                            v6 := GenS1{F0: (2 > 130)}
                            _ = v6
                        case "":
                            var v7 int = fi5
                            _ = v7
                        case "hello":
                            for fi8 := 0; fi8 < 2; fi8++ {
                                v9 := GenS2{F0: 0, F1: (128 >= 1)}
                                var v10 float32 = ((float32(float64(3.14159)) - 1e10) * 1.5)
                                fmt.Println(fi8)
                                _ = fi8
                                _ = v10
                                _ = v9
                            }
                        default:
                            v11 := GenS2{F0: -264, F1: (!false)}
                            _ = v11
                        }
                        fmt.Println(v2, v2)
                        v12 := GenS3{F0: (v2 | v2), F1: fi5, F2: 3}
                        _ = v12
                    } else {
                        var v13 uint8 = 127
                        var v14 int32 = -845
                        _ = v13
                        _ = v14
                    }
                    v3 = 1
                    _ = fi5
                }
            } else {
                for fi15 := 0; fi15 < 4; fi15++ {
                    var v16 int16 = -73
                    v16 += (v16 % (((-(v16))) | 1))
                    v17 := GenS3{F0: (v2 & 0), F1: fi15, F2: 127}
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            }
            v3 += v3
        }
    case "a":
        var v18 int16 = -638
        _ = v18
    case "":
        fmt.Println(v4, v4, v2)
    }
    v19 := GenS4{F0: 42, F1: 929}
    v3 = v3
    switch -2147483648 {
    case -9223372036854775808:
        for fi20 := 0; fi20 < 11; fi20++ {
            var v21 uint8 = 219
            _ = fi20
            _ = v21
        }
    default:
        v22 := GenS5{F0: -15}
        _ = v22
    }
    if true {
        v23 := GenS5{F0: v3}
        if (true || true) {
            if (!(false && (4294967295 >= 892))) {
                for fi24 := 0; fi24 < 15; fi24++ {
                    fmt.Println(v4, v4)
                    var v25 int16 = 518
                    _ = fi24
                    _ = v25
                }
            } else {
                if (!((512 <= 38) && false)) {
                    for fi26 := 0; fi26 < 3; fi26++ {
                        fmt.Println(v4, v2, v2)
                        var v27 float32 = 0.0
                        v28 := []int{fi26, -583, (fi26 << ((42) & 63))}
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                } else {
                    switch 1 {
                    case -1:
                        v29 := GenS4{F0: 42, F1: 1}
                        _ = v29
                    case 2147483647:
                        for fi30 := 0; fi30 < 20; fi30++ {
                            if false {
                                for fi31 := 0; fi31 < 19; fi31++ {
                                    v2 = -432
                                    fmt.Println(v2, v3)
                                    var v32 uint16 = 391
                                    _ = fi31
                                    _ = v32
                                }
                                switch fi30 {
                                case 2147483647:
                                    for fi33 := 0; fi33 < 18; fi33++ {
                                        var v34 int = 1
                                        switch "hello" {
                                        case "foo bar":
                                            if false {
                                                var v35 uint8 = 72
                                                fmt.Println(fi30)
                                                if (!false) {
                                                    v4 = (v4 ^ v4)
                                                } else {
                                                    v34 = (v34 & 4294967295)
                                                    var v36 uint = 2
                                                    _ = v36
                                                }
                                                if (true && false) {
                                                    if false {
                                                        if false {
                                                            if false {
                                                                if false {
                                                                    var v37 float64 = 0.0
                                                                    v37 -= 0.0
                                                                    v34 = fi33
                                                                    _ = v37
                                                                } else {
                                                                    v38 := GenS2{F0: 789, F1: (!(-2.25 >= (0.1 + 0.0)))}
                                                                    _ = v38
                                                                }
                                                                switch int(v4) {
                                                                case 127:
                                                                    v35 -= (v35 + (v35 ^ (v35 | (v35 / ((128) | 1)))))
                                                                }
                                                                if (v35 >= uint8(v4)) {
                                                                    switch v34 {
                                                                    case 0:
                                                                        _ = len(v1)
                                                                    case 2147483647:
                                                                        if false {
                                                                            v1[(("foo bar" + "foo bar") + "a")] = (fi33 & (fi33 & fi33))
                                                                        }
                                                                    }
                                                                    var v39 string = "foo bar"
                                                                    if true {
                                                                        fmt.Println(v4, v2, v35)
                                                                        if (false || true) {
                                                                            v40 := GenS6{F0: v35, F1: v2}
                                                                            _ = v40
                                                                        } else {
                                                                            var v41 uint8 = 0
                                                                            _ = v41
                                                                        }
                                                                    } else {
                                                                        if false {
                                                                            switch v34 {
                                                                            case -864:
                                                                                switch fi30 {
                                                                                case 128:
                                                                                    _ = len(v1)
                                                                                case 2147483647:
                                                                                    var v42 uint8 = (^(v35))
                                                                                    _ = v42
                                                                                }
                                                                            }
                                                                            var v43 int32 = 181
                                                                            if ((true || false) || (2 <= (v4 * (v4 << ((4294967295) & 63))))) {
                                                                                v44 := GenS5{F0: (v3 / ((v3) | 1))}
                                                                                if (true && true) {
                                                                                    switch fi33 {
                                                                                    case -2147483648:
                                                                                        v39 = "hello"
                                                                                    case 2147483647:
                                                                                        v39 = "foo bar"
                                                                                    default:
                                                                                        v45 := GenS1{F0: (-24 <= -1)}
                                                                                        _ = v45
                                                                                    }
                                                                                    switch v39 {
                                                                                    case "foo bar":
                                                                                        if (v4 == v4) {
                                                                                            fmt.Println(v2, fi30, v3)
                                                                                            if (v2 >= 9223372036854775807) {
                                                                                                fmt.Println(v2, fi33, fi30)
                                                                                                var v46 float64 = -2.25
                                                                                                _ = v46
                                                                                            } else {
                                                                                                var v47 int64 = int64(v3)
                                                                                                v47 += 275
                                                                                                switch v39 {
                                                                                                case "foo bar":
                                                                                                    fmt.Println(v39, v47, v4)
                                                                                                default:
                                                                                                    var v48 bool = (!(fi33 < v34))
                                                                                                    _ = v48
                                                                                                }
                                                                                                _ = v47
                                                                                            }
                                                                                        }
                                                                                    case "":
                                                                                        var v49 uint32 = 2147483648
                                                                                        v50 := GenS5{F0: (v3 & v3)}
                                                                                        fmt.Println(v43, v2, v4)
                                                                                        _ = v49
                                                                                        _ = v50
                                                                                    }
                                                                                } else {
                                                                                    v51 := GenS5{F0: -128}
                                                                                    var v52 int16 = 0
                                                                                    var v53 uint32 = 784
                                                                                    v53 = 42
                                                                                    _ = v51
                                                                                    _ = v52
                                                                                    _ = v53
                                                                                }
                                                                                _ = v44
                                                                            } else {
                                                                                if true {
                                                                                    fmt.Println(fi30)
                                                                                    v35 -= v35
                                                                                }
                                                                            }
                                                                            _ = v43
                                                                        }
                                                                    }
                                                                    v54 := []int{(fi33 ^ 0)}
                                                                    _ = v39
                                                                    _ = v54
                                                                }
                                                            } else {
                                                                v55 := GenS7{F0: "hello", F1: 18446744073709551615}
                                                                _ = v55
                                                            }
                                                            _ = len(v1)
                                                            _ = len(v1)
                                                            v35 = v35
                                                        } else {
                                                            switch 32768 {
                                                            case 32768:
                                                                var v56 int = v34
                                                                _ = v56
                                                            case 0:
                                                                var v57 int8 = (^(v3))
                                                                _ = v57
                                                            default:
                                                                switch 996 {
                                                                case 127:
                                                                    v58 := GenS8{F0: -240}
                                                                    _ = v58
                                                                case 2147483647:
                                                                    var v59 uint16 = (v4 | v4)
                                                                    _ = v59
                                                                case 985:
                                                                    v3 += (v3 & (v3 >> ((42) & 63)))
                                                                default:
                                                                    v60 := GenS1{F0: (false || false)}
                                                                    _ = v60
                                                                }
                                                            }
                                                            v4 -= (v4 & 2)
                                                            var v61 uint64 = 1
                                                            _ = v61
                                                        }
                                                    } else {
                                                        delete(v1, "\n\t")
                                                        var v62 int32 = -1
                                                        _ = v62
                                                    }
                                                    v2 += (v2 | (v2 - (v2 | (v2 | (v2 | (v2 & 6))))))
                                                    var v63 int64 = (v2 / ((v2) | 1))
                                                    _ = v63
                                                } else {
                                                    switch fi30 {
                                                    case 32767:
                                                        v2 = 0
                                                    case 9223372036854775807:
                                                        v64 := []int{}
                                                        _ = v64
                                                    }
                                                    if false {
                                                        fmt.Println(v4, v3)
                                                    } else {
                                                        v2 += v2
                                                        var v65 float64 = 0.1
                                                        _ = v65
                                                    }
                                                }
                                                var v66 uint16 = v4
                                                _ = v35
                                                _ = v66
                                            }
                                        case "hello":
                                            if false {
                                                if (993 > 1) {
                                                    fmt.Println(v4, fi33)
                                                } else {
                                                    if ((!true) || false) {
                                                        v2 = (v2 ^ 0)
                                                    } else {
                                                        var v67 uint64 = 2
                                                        v34 = (-(fi33))
                                                        _ = v67
                                                    }
                                                }
                                            } else {
                                                v68 := GenS5{F0: v3}
                                                _ = v68
                                            }
                                        }
                                        var v69 int64 = 42
                                        _ = fi33
                                        _ = v34
                                        _ = v69
                                    }
                                case -946:
                                    v70 := GenS9{F0: (fi30 | fi30)}
                                    _ = v70
                                }
                            } else {
                                v3 = v3
                                v71 := []int{fi30}
                                var v72 uint = 566
                                _ = v71
                                _ = v72
                            }
                            v73 := v1[""]
                            for fi74 := 0; fi74 < 8; fi74++ {
                                v4 += (v4 - 42)
                                switch fi74 {
                                case 32768:
                                    var v75 uint32 = 837
                                    _ = v75
                                case 128:
                                    v4 += (v4 - v4)
                                case -359:
                                    if true {
                                        v3 = v3
                                    } else {
                                        v76 := GenS4{F0: 32767, F1: 1}
                                        var v77 uint8 = 255
                                        _ = v76
                                        _ = v77
                                    }
                                }
                                fmt.Println(v73, fi30)
                                _ = fi74
                            }
                            _ = fi30
                            _ = v73
                        }
                    default:
                        v78 := []int{-65, 32767, 9223372036854775807}
                        _ = v78
                    }
                    v4 += 42
                    var v79 uint = 4294967295
                    _ = v79
                }
            }
            for fi80 := 0; fi80 < 8; fi80++ {
                _ = len(v1)
                if (v3 > v3) {
                    fmt.Println(fi80)
                    var v81 float64 = 1.5
                    _ = v81
                } else {
                    for fi82 := 0; fi82 < 4; fi82++ {
                        switch ("foo bar" + "foo bar") {
                        case "hello":
                            fmt.Println(fi82)
                        case "a":
                            v3 = 1
                        }
                        v83 := GenS10{F0: uint32(v4)}
                        fmt.Println(v3)
                        _ = fi82
                        _ = v83
                    }
                    v3 = 42
                }
                _ = fi80
            }
            if true {
                v3 = (^(v3))
                var v84 int = int(v3)
                _ = v84
            } else {
                v3 = v3
            }
        } else {
            v85 := GenS1{F0: (771 >= v4)}
            if false {
                if false {
                    v3 -= 111
                    var v86 uint = 838
                    var v87 int16 = 42
                    _ = v86
                    _ = v87
                } else {
                    for fi88 := 0; fi88 < 19; fi88++ {
                        v89 := GenS11{F0: 0, F1: 32767, F2: (v2 + -1)}
                        switch fi88 {
                        case 742:
                            v90 := []int{(fi88 % (((fi88 | 127)) | 1))}
                            _ = v90
                        case 83:
                            fmt.Println(v2)
                        default:
                            if true {
                                var v91 uint16 = v4
                                _ = v91
                            } else {
                                if true {
                                    var v92 int16 = 32767
                                    v3 = v3
                                    var v93 uint16 = (^(v4))
                                    _ = v92
                                    _ = v93
                                } else {
                                    if (623 <= 212) {
                                        fmt.Println(v4)
                                    }
                                    if (true && false) {
                                        if ((!false) || true) {
                                            v94 := GenS12{F0: (!(false || (int(v3) <= fi88))), F1: (-(v3))}
                                            var v95 float64 = (((0.1 / 0.1) / (1e10 / 1.0)) + (1e10 + float64(-2.25)))
                                            var v96 int64 = (v2 >> ((0) & 63))
                                            _ = v94
                                            _ = v95
                                            _ = v96
                                        } else {
                                            v97 := []int{}
                                            _ = v97
                                        }
                                    }
                                    var v98 int8 = v3
                                    _ = v98
                                }
                            }
                        }
                        v3 += v3
                        _ = fi88
                        _ = v89
                    }
                    v4 = v4
                }
                if (100 <= 236) {
                    if true {
                        v3 = v3
                        if true {
                            v4 = 552
                            v3 += (v3 - (v3 >> ((445) & 63)))
                            v2 -= (v2 ^ v2)
                        } else {
                            switch 42 {
                            case 42:
                                var v99 uint8 = 0
                                _ = v99
                            case 32767:
                                for fi100 := 0; fi100 < 4; fi100++ {
                                    v4 = 32767
                                    var v101 int8 = v3
                                    _ = fi100
                                    _ = v101
                                }
                            case 0:
                                v102 := v1["foo bar"]
                                _ = v102
                            default:
                                var v103 int8 = (v3 * (v3 * (v3 + v3)))
                                _ = v103
                            }
                        }
                        for fi104 := 0; fi104 < 16; fi104++ {
                            v105 := []int{-969, (fi104 % (((fi104 / ((fi104) | 1))) | 1)), fi104, int(v2)}
                            _ = fi104
                            _ = v105
                        }
                    } else {
                        switch 32767 {
                        case 32767:
                            fmt.Println(v4, v4, v4)
                        case 42:
                            if (568 > -1) {
                                for fi106 := 0; fi106 < 17; fi106++ {
                                    for fi107 := 0; fi107 < 17; fi107++ {
                                        v4 -= 42
                                        if true {
                                            _ = len(v1)
                                        } else {
                                            v2 = (v2 * -1)
                                            var v108 float64 = 1e10
                                            _ = v108
                                        }
                                        _ = fi107
                                    }
                                    var v109 float64 = (0.1 / float64((0.1 - 1e10)))
                                    v110 := GenS13{F0: (v2 / (((v2 | 42)) | 1))}
                                    _ = fi106
                                    _ = v109
                                    _ = v110
                                }
                            } else {
                                v2 = v2
                                v4 -= 1
                                if (-720 >= -995) {
                                    fmt.Println(v4)
                                    v111 := GenS14{F0: false}
                                    var v112 int16 = -1
                                    _ = v111
                                    _ = v112
                                } else {
                                    var v113 uint = 18446744073709551615
                                    if true {
                                        for fi114 := 0; fi114 < 2; fi114++ {
                                            v3 += (v3 / (((-(v3))) | 1))
                                            if (1 >= 607) {
                                                v3 -= 127
                                                var v115 float64 = (float64(-2.25) * (((float64(3.14159) * float64(0.0)) / -2.25) + 1e10))
                                                for fi116 := 0; fi116 < 15; fi116++ {
                                                    var v117 uint32 = 1
                                                    if true {
                                                        fmt.Println(v2)
                                                        var v118 int8 = (v3 & v3)
                                                        var v119 int64 = (v2 + v2)
                                                        _ = v118
                                                        _ = v119
                                                    }
                                                    v117 += (v117 % ((42) | 1))
                                                    _ = fi116
                                                    _ = v117
                                                }
                                                _ = v115
                                            }
                                            v3 += -128
                                            _ = fi114
                                        }
                                        switch 32767 {
                                        case -342:
                                            switch 826 {
                                            case 0:
                                                v3 += 1
                                            }
                                        case 9223372036854775807:
                                            switch 4294967295 {
                                            case 696:
                                                var v120 bool = true
                                                _ = v120
                                            case -2147483648:
                                                v121 := []int{32767}
                                                _ = v121
                                            }
                                        case -444:
                                            var v122 int32 = 99
                                            _ = v122
                                        }
                                        var v123 int16 = -1
                                        _ = v123
                                    }
                                    _ = v113
                                }
                            }
                        default:
                            var v124 uint8 = 134
                            _ = v124
                        }
                        switch "foo bar" {
                        case "":
                            v125 := []int{-893, 42, 9223372036854775807}
                            _ = v125
                        case "\n\t":
                            fmt.Println(v2, v2, v4)
                        }
                        v126 := GenS11{F0: 42, F1: -420, F2: (^(v2))}
                        _ = v126
                    }
                    var v127 string = "\n\t"
                    _ = v127
                } else {
                    if true {
                        var v128 int16 = int16(v3)
                        var v129 uint32 = 92
                        v2 -= (v2 % (((^(v2))) | 1))
                        _ = v128
                        _ = v129
                    }
                    fmt.Println(v2, v4)
                }
            } else {
                var v130 int = 32767
                var v131 bool = true
                switch "a" {
                case "foo bar":
                    v132 := GenS5{F0: int8(v2)}
                    _ = v132
                case "":
                    fmt.Println(v131, v2, v4)
                case "\n\t":
                    for fi133 := 0; fi133 < 20; fi133++ {
                        for fi134 := 0; fi134 < 15; fi134++ {
                            if v131 {
                                var v135 int8 = (^(v3))
                                var v136 int = (v130 + 737)
                                _ = v135
                                _ = v136
                            }
                            _ = fi134
                        }
                        _ = fi133
                    }
                default:
                    var v137 float64 = (1e10 + -2.25)
                    _ = v137
                }
                _ = v130
                _ = v131
            }
            v138 := []int{}
            _ = v138
            _ = v85
        }
        v3 = (v3 >> ((493) & 63))
        _ = v23
    } else {
        v139 := []int{32768, 32767}
        fmt.Println(v4, v3)
        if len(v139) >= 2 {
            v139 = v139[1:len(v139)]
        }
        _ = v139
    }
    fmt.Println(v4, v3, v2)
    fmt.Println(v3, v2, v4)
    v140 := GenS3{F0: (v2 | v2), F1: 2147483647, F2: 0}
    var v141 uint32 = 631
    for fi142 := 0; fi142 < 10; fi142++ {
        v3 += (^(v3))
        _ = fi142
    }
    var v143 uint64 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v141:", v141)
    fmt.Println("v143:", v143)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v140
    _ = v141
    _ = v143
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
