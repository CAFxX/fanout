package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int32 }
type GenS2 struct { F0 uint8; F1 string; F2 uint }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 uint32; F1 int32; F2 uint }
type GenS5 struct { F0 string; F1 string }
type GenS6 struct { F0 uint; F1 int8 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int32; F1 int32; F2 int32 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int; F1 string }
type GenS11 struct { F0 string; F1 uint }
type GenS12 struct { F0 uint8; F1 uint32; F2 int16 }
type GenS13 struct { F0 uint64; F1 int64 }
type GenS14 struct { F0 uint64; F1 int8; F2 uint8 }
type GenS15 struct { F0 uint8; F1 uint16 }
type GenS16 struct { F0 uint8; F1 int8 }
type GenS17 struct { F0 int16; F1 uint64 }
type GenS18 struct { F0 string; F1 int; F2 uint }
func genFn0() int64 {
    var v7 float64 = (3.14159 / 1e10)
    v8 := GenS1{F0: 1000, F1: -98, F2: 42}
    if (!((true || (4294967295 < 1)) || (0 > 1))) {
        var v9 bool = false
        if v9 {
            if v9 {
                switch 32768 {
                case 32767:
                    fmt.Println(v9, v7)
                case 127:
                    v10 := GenS2{F0: 148, F1: "foo bar", F2: 65535}
                    _ = v10
                }
                v7 = float64(3.14159)
                var v11 uint32 = 65535
                _ = v11
            } else {
                var v12 bool = v9
                v7 = 0.0
                _ = v12
            }
            if v9 {
                v13 := GenS1{F0: 1, F1: 123, F2: -1}
                for fi14 := 0; fi14 < 16; fi14++ {
                    v9 = false
                    v9 = false
                    _ = fi14
                }
                _ = v13
            } else {
                for fi15 := 0; fi15 < 6; fi15++ {
                    v9 = (!(3.14159 == float32(v7)))
                    if true {
                        var v16 int8 = -128
                        _ = v16
                    } else {
                        v17 := GenS1{F0: 0, F1: -119, F2: 1}
                        for fi18 := 0; fi18 < 12; fi18++ {
                            v19 := GenS1{F0: 1000, F1: 1, F2: 42}
                            var v20 uint = 796
                            v21 := []int{(fi15 / ((fi15) | 1))}
                            _ = fi18
                            _ = v19
                            _ = v20
                            _ = v21
                        }
                        for fi22 := 0; fi22 < 12; fi22++ {
                            if v9 {
                                var v23 string = "a"
                                _ = v23
                            }
                            var v24 int8 = 127
                            v9 = v9
                            _ = fi22
                            _ = v24
                        }
                        _ = v17
                    }
                    _ = fi15
                }
                var v25 uint16 = 65535
                _ = v25
            }
            switch "a" {
            case "\n\t":
                v7 += 1e10
            case "":
                v9 = v9
            case "a":
                if v9 {
                    v26 := []int{}
                    if (!v9) {
                        fmt.Println(v7)
                        var v27 int16 = -32768
                        _ = v27
                    }
                    _ = v26
                }
            default:
                fmt.Println(v9, v9)
            }
        }
        _ = v9
    } else {
        for fi28 := 0; fi28 < 3; fi28++ {
            for fi29 := 0; fi29 < 15; fi29++ {
                v30 := []int{fi29, (fi28 * (fi28 * fi29)), (fi29 - 0), (fi29 >> ((0) & 63))}
                _ = fi29
                _ = v30
            }
            v31 := GenS3{F0: "hello", F1: 670}
            _ = fi28
            _ = v31
        }
    }
    _ = v7
    _ = v8
    return 782
}

func main() {
    var v1 uint64 = 1
    v2 := make(map[int]int)
    v3 := make(map[string]bool)
    var v4 bool = false
    var v5 uint = 186
    var v6 uint8 = 128
    v32 := []int{127}
    v33 := GenS2{F0: 19, F1: ("a" + "\n\t"), F2: 260}
    v34 := GenS1{F0: 1, F1: 1, F2: -2147483648}
    var v35 string = "foo bar"
    for fi36 := 0; fi36 < 9; fi36++ {
        v3["\n\t"] = v4
        var v37 int32 = 0
        _ = fi36
        _ = v37
    }
    v38 := GenS4{F0: 65535, F1: 579, F2: uint(v6)}
    v32 = append(v32, -1)
    v6 += v6
    if ((((v1 << ((553) & 63)) < uint64(v5)) || (((0.1 - (((0.1 + ((0.0 * 1.5) - (3.14159 * 0.0))) + -2.25) * 0.1)) == (0.1 * 1e10)) && (69 <= 100))) && false) {
        if (773 <= -2147483648) {
            if (!true) {
                var v39 int8 = 1
                _ = v39
            } else {
                _ = len(v2)
                v40 := GenS4{F0: 0, F1: -402, F2: (^(v5))}
                _ = v40
            }
            v4 = (42 > int8(v5))
        } else {
            v35 = ((v35 + ((((v35 + v35) + v35) + v35) + v35)) + v35)
        }
        _ = len(v3)
        v41 := GenS5{F0: v35, F1: v35}
        _ = v41
    } else {
        var v42 bool = (float64(3.14159) <= -2.25)
        fmt.Println(v1, v6)
        _ = v42
    }
    if (42 < 0) {
        var v43 int8 = 42
        v44 := GenS6{F0: 65535, F1: (v43 % ((-128) | 1))}
        _ = v43
        _ = v44
    }
    if true {
        if false {
            var v45 int = 2147483647
            if (false || (654 <= -145)) {
                v1 -= 1
                v46 := GenS5{F0: v35, F1: v35}
                _ = v46
            } else {
                v35 = v35
            }
            switch (v35 + v35) {
            case "":
                v35 = v35
            case "a":
                if v4 {
                    var v47 uint16 = 65535
                    v45 = v45
                    var v48 string = "foo bar"
                    _ = v47
                    _ = v48
                } else {
                    v49 := GenS7{F0: 42}
                    switch ("a" + "\n\t") {
                    case "foo bar":
                        for fi50 := 0; fi50 < 14; fi50++ {
                            v51 := GenS8{F0: 1, F1: 0, F2: 42}
                            var v52 int64 = -9223372036854775808
                            _ = fi50
                            _ = v51
                            _ = v52
                        }
                    case "\n\t":
                        for fi53 := 0; fi53 < 1; fi53++ {
                            switch v35 {
                            case "":
                                for fi54 := 0; fi54 < 19; fi54++ {
                                    if v4 {
                                        fmt.Println(v45, v35)
                                    } else {
                                        if v4 {
                                            var v55 float32 = 0.0
                                            _ = v55
                                        } else {
                                            if v4 {
                                                var v56 float64 = 1e10
                                                _ = v56
                                            } else {
                                                fmt.Println(v45)
                                                fmt.Println(v45)
                                                v57 := GenS2{F0: (v6 << (((v5 + (v5 ^ (v5 >> ((v5) & 63))))) & 63)), F1: v35, F2: v5}
                                                _ = v57
                                            }
                                            if (!((true && false) || (-2.25 != 0.0))) {
                                                if v4 {
                                                    v5 += 255
                                                    if v4 {
                                                        v5 += v5
                                                    } else {
                                                        var v58 float64 = float64(0.0)
                                                        if len(v32) > 0 {
                                                            v32[0] = fi53
                                                        }
                                                        _ = v58
                                                    }
                                                    _ = len(v32)
                                                } else {
                                                    v45 -= (v45 * v45)
                                                    v59 := GenS3{F0: ("foo bar" + ("\n\t" + ("foo bar" + v35))), F1: v5}
                                                    var v60 int32 = 42
                                                    v35 = v35
                                                    _ = v59
                                                    _ = v60
                                                }
                                                if false {
                                                    v61 := GenS1{F0: -1, F1: 76, F2: 1}
                                                    _ = v61
                                                } else {
                                                    switch v35 {
                                                    case "\n\t":
                                                        if v4 {
                                                            var v62 uint64 = 2
                                                            if v4 {
                                                                if (true && (1e10 == (float32(1e10) - ((0.1 + 1e10) + (0.0 / (1.5 / 1e10)))))) {
                                                                    v62 = 4294967295
                                                                    var v63 string = v35
                                                                    v64 := GenS9{F0: 42}
                                                                    delete(v3, v63)
                                                                    _ = v63
                                                                    _ = v64
                                                                }
                                                                v65 := GenS3{F0: "foo bar", F1: (v5 / ((65535) | 1))}
                                                                _ = v65
                                                            }
                                                            if len(v32) >= 2 {
                                                                v32 = v32[1:len(v32)]
                                                            }
                                                            _ = v62
                                                        } else {
                                                            var v66 int32 = -187
                                                            v4 = v4
                                                            _ = v66
                                                        }
                                                    case "a":
                                                        if (745 >= 32767) {
                                                            var v67 int16 = -1
                                                            if false {
                                                                var v68 uint8 = v6
                                                                var v69 uint32 = 2147483648
                                                                if false {
                                                                    if (float64(0.0) <= float64(1e10)) {
                                                                        fmt.Println(v67, v69)
                                                                        if len(v32) >= 2 {
                                                                            v32 = v32[1:len(v32)]
                                                                        }
                                                                    } else {
                                                                        var v70 uint32 = (v69 - v69)
                                                                        v1 = v1
                                                                        _ = v70
                                                                    }
                                                                }
                                                                _ = v68
                                                                _ = v69
                                                            }
                                                            v3["\n\t"] = (0 != 100)
                                                            _ = v67
                                                        } else {
                                                            if (268 != v1) {
                                                                if v4 {
                                                                    if v4 {
                                                                        var v71 string = v35
                                                                        switch fi54 {
                                                                        case 1:
                                                                            fmt.Println(fi54, v35)
                                                                        case -240:
                                                                            if (!false) {
                                                                                var v72 int64 = 1
                                                                                v72 = genFn0()
                                                                                v1 += (^(v1))
                                                                                v73 := GenS1{F0: -184, F1: -128, F2: 752}
                                                                                _ = v72
                                                                                _ = v73
                                                                            }
                                                                        }
                                                                        if true {
                                                                            v45 += 9223372036854775807
                                                                            var v74 float64 = 3.14159
                                                                            _ = v74
                                                                        } else {
                                                                            v75 := GenS2{F0: v6, F1: "a", F2: (v5 << ((907) & 63))}
                                                                            var v76 uint16 = 349
                                                                            var v77 float64 = (1.5 / (3.14159 + (-2.25 + 0.1)))
                                                                            var v78 int32 = 0
                                                                            var v79 float64 = 1.5
                                                                            var v80 int = v45
                                                                            if (!true) {
                                                                                var v81 int64 = 9223372036854775807
                                                                                _ = v81
                                                                            }
                                                                            _ = v75
                                                                            _ = v76
                                                                            _ = v77
                                                                            _ = v78
                                                                            _ = v79
                                                                            _ = v80
                                                                        }
                                                                        v1 = v1
                                                                        v45 = -203
                                                                        switch v35 {
                                                                        case "\n\t":
                                                                            delete(v3, v35)
                                                                        case "":
                                                                            v82 := GenS1{F0: -712, F1: -128, F2: 690}
                                                                            _ = v82
                                                                        case "hello":
                                                                            var v83 int = fi53
                                                                            _ = v83
                                                                        default:
                                                                            if v4 {
                                                                                switch "hello" {
                                                                                case "\n\t":
                                                                                    v2[fi53] = (-(fi53))
                                                                                default:
                                                                                    v45 = fi53
                                                                                }
                                                                                if (1 > genFn0()) {
                                                                                    v2[(fi54 % ((fi53) | 1))] = (v45 * int(v6))
                                                                                    v84 := v3[("a" + v35)]
                                                                                    _ = v84
                                                                                } else {
                                                                                    var v85 int32 = -1
                                                                                    v71 = ((v71 + "foo bar") + "foo bar")
                                                                                    _ = v85
                                                                                }
                                                                                if (772 != 704) {
                                                                                    var v86 int64 = genFn0()
                                                                                    _ = v86
                                                                                }
                                                                            } else {
                                                                                if (972 == 830) {
                                                                                    fmt.Println(v5)
                                                                                    if false {
                                                                                        v4 = ((1 >= 700) || ((1e10 > 0.1) && true))
                                                                                        v87 := v2[fi54]
                                                                                        switch v71 {
                                                                                        case "foo bar":
                                                                                            var v88 uint32 = 62
                                                                                            _ = v88
                                                                                        }
                                                                                        _ = v87
                                                                                    } else {
                                                                                        v71 = v71
                                                                                        switch fi53 {
                                                                                        case 127:
                                                                                            v89 := GenS10{F0: (v45 + (-(fi54))), F1: (v35 + v71)}
                                                                                            _ = v89
                                                                                        }
                                                                                        if len(v32) > 0 {
                                                                                            v32[0] = v45
                                                                                        }
                                                                                        v90 := GenS3{F0: v35, F1: v5}
                                                                                        _ = v90
                                                                                    }
                                                                                    v91 := GenS1{F0: -1, F1: -101, F2: -862}
                                                                                    if v4 {
                                                                                        var v92 int16 = -32768
                                                                                        v4 = v4
                                                                                        v71 = v35
                                                                                        _ = v92
                                                                                    }
                                                                                    var v93 uint16 = 16
                                                                                    _ = v91
                                                                                    _ = v93
                                                                                } else {
                                                                                    v94 := GenS11{F0: "foo bar", F1: (v5 << ((909) & 63))}
                                                                                    _ = v94
                                                                                }
                                                                            }
                                                                        }
                                                                        _ = v71
                                                                    } else {
                                                                        switch (fi54 | v45) {
                                                                        case 42:
                                                                            var v95 float64 = ((0.0 - float64((float32(1e10) - 1e10))) / 1.0)
                                                                            _ = v95
                                                                        }
                                                                    }
                                                                    if ((!v4) || (1.5 <= 1e10)) {
                                                                        switch fi54 {
                                                                        case -918:
                                                                            v45 = (-(fi53))
                                                                        case 970:
                                                                            if v4 {
                                                                                v6 = v6
                                                                                var v96 string = (("" + ("\n\t" + v35)) + v35)
                                                                                _ = v96
                                                                            }
                                                                        case 205:
                                                                            delete(v3, "foo bar")
                                                                        default:
                                                                            switch "hello" {
                                                                            case "":
                                                                                v97 := GenS7{F0: 0}
                                                                                _ = v97
                                                                            case "foo bar":
                                                                                fmt.Println(v4)
                                                                            case "\n\t":
                                                                                fmt.Println(v35, v45)
                                                                            default:
                                                                                var v98 int16 = -32768
                                                                                _ = v98
                                                                            }
                                                                        }
                                                                        fmt.Println(v6)
                                                                    }
                                                                    var v99 uint = 1
                                                                    _ = v99
                                                                }
                                                                var v100 float64 = ((float64(-2.25) + (0.1 + 0.0)) * ((1e10 + 0.1) * (1e10 / ((1e10 * (1.5 / ((float64(-2.25) / 1.5) - 3.14159))) + (0.0 + 0.1)))))
                                                                _ = v100
                                                            }
                                                            var v101 float32 = (1e10 - ((-2.25 / 1.5) * (float32((1.5 + 1.5)) + 0.1)))
                                                            v32 = append(v32, (fi53 + v45))
                                                            _ = v101
                                                        }
                                                    case "hello":
                                                        v102 := GenS3{F0: ("hello" + "hello"), F1: v5}
                                                        _ = v102
                                                    }
                                                    var v103 uint16 = 0
                                                    _ = v103
                                                }
                                                v5 = 42
                                            } else {
                                                switch v35 {
                                                case "foo bar":
                                                    switch 9223372036854775807 {
                                                    case -1:
                                                        v4 = (!v4)
                                                    case -2147483648:
                                                        fmt.Println(v45, v35)
                                                        var v104 uint8 = (v6 << (((v5 | v5)) & 63))
                                                        _ = v104
                                                    }
                                                case "\n\t":
                                                    var v105 string = (("\n\t" + ("" + v35)) + "a")
                                                    _ = v105
                                                }
                                                if ((^(v5)) > (v5 / (((v5 + (v5 - (^(v5))))) | 1))) {
                                                    var v106 uint32 = 1
                                                    v106 = 2147483647
                                                    if len(v32) >= 2 {
                                                        v32 = v32[1:len(v32)]
                                                    }
                                                    _ = v106
                                                }
                                            }
                                            var v107 uint64 = uint64(v6)
                                            v108 := GenS11{F0: v35, F1: (v5 | (v5 / ((2) | 1)))}
                                            _ = v107
                                            _ = v108
                                        }
                                    }
                                    _ = fi54
                                }
                            case "foo bar":
                                if len(v32) >= 2 {
                                    v32 = v32[1:len(v32)]
                                }
                            }
                            var v109 uint8 = (v6 | 127)
                            _ = len(v32)
                            _ = fi53
                            _ = v109
                        }
                    case "":
                        if v4 {
                            if len(v32) > 0 {
                                v32[0] = -9223372036854775808
                            }
                        }
                    }
                    if (42 >= v1) {
                        for fi110 := 0; fi110 < 19; fi110++ {
                            fmt.Println(v6, v1, v5)
                            _ = fi110
                        }
                        if true {
                            if v4 {
                                var v111 int16 = -513
                                if v4 {
                                    for fi112 := 0; fi112 < 9; fi112++ {
                                        v113 := GenS12{F0: (v6 + v6), F1: 0, F2: v111}
                                        var v114 float32 = (3.14159 * 3.14159)
                                        if v4 {
                                            for fi115 := 0; fi115 < 2; fi115++ {
                                                if ((v35 + v35) == "foo bar") {
                                                    v116 := GenS3{F0: v35, F1: v5}
                                                    _ = v116
                                                }
                                                v6 += (v6 / ((v6) | 1))
                                                if (!(!v4)) {
                                                    switch ("foo bar" + "\n\t") {
                                                    case "hello":
                                                        v117 := v3[v35]
                                                        v2[(fi115 - v45)] = fi115
                                                        v118 := GenS10{F0: 42, F1: (v35 + "a")}
                                                        _ = v117
                                                        _ = v118
                                                    }
                                                    v119 := GenS13{F0: v1, F1: 1}
                                                    _ = v119
                                                } else {
                                                    v32 = append(v32, fi112)
                                                    v4 = ((v111 - -32768) < 1000)
                                                }
                                                _ = fi115
                                            }
                                        }
                                        _ = fi112
                                        _ = v113
                                        _ = v114
                                    }
                                    var v120 int8 = 127
                                    _ = v120
                                }
                                _ = v111
                            } else {
                                _ = len(v32)
                            }
                            for fi121 := 0; fi121 < 17; fi121++ {
                                if v4 {
                                    _ = len(v2)
                                    var v122 int = fi121
                                    _ = v122
                                } else {
                                    var v123 float32 = 1.5
                                    var v124 int16 = 1000
                                    _ = len(v32)
                                    _ = v123
                                    _ = v124
                                }
                                v1 -= 42
                                var v125 bool = v4
                                _ = fi121
                                _ = v125
                            }
                            for fi126 := 0; fi126 < 10; fi126++ {
                                v6 = (v6 | (v6 << ((255) & 63)))
                                _ = fi126
                            }
                        }
                    }
                    _ = v49
                }
            case "foo bar":
                fmt.Println(v6, v6, v5)
            }
            _ = v45
        } else {
            if v4 {
                var v127 uint64 = 1
                v1 -= v1
                v128 := GenS14{F0: v127, F1: -43, F2: 127}
                _ = v127
                _ = v128
            }
        }
        if ((127 > -1) && (-591 == 1)) {
            for fi129 := 0; fi129 < 18; fi129++ {
                v35 = v35
                for fi130 := 0; fi130 < 19; fi130++ {
                    v131 := GenS4{F0: 826, F1: 1, F2: v5}
                    v132 := GenS15{F0: 255, F1: 255}
                    switch v35 {
                    case "hello":
                        var v133 uint8 = (^(v6))
                        _ = v133
                    case "":
                        if len(v32) >= 2 {
                            v32 = v32[1:len(v32)]
                        }
                    }
                    _ = fi130
                    _ = v131
                    _ = v132
                }
                _ = fi129
            }
        } else {
            var v134 string = "\n\t"
            var v135 float64 = (1.5 / (3.14159 - 0.1))
            v6 += (v6 % ((uint8(v1)) | 1))
            _ = v134
            _ = v135
        }
        var v136 float32 = (1.5 + 0.0)
        _ = v136
    } else {
        for fi137 := 0; fi137 < 13; fi137++ {
            switch "a" {
            case "hello":
                if v4 {
                    var v138 uint = v5
                    _ = v138
                }
            case "":
                for fi139 := 0; fi139 < 17; fi139++ {
                    switch (fi137 & fi137) {
                    case 1:
                        if v4 {
                            var v140 uint8 = (v6 | v6)
                            _ = v140
                        } else {
                            var v141 float32 = 3.14159
                            _ = v141
                        }
                    }
                    v142 := v2[(fi137 + 32768)]
                    if (v4 || v4) {
                        if (-2147483648 <= 42) {
                            v143 := GenS16{F0: v6, F1: 0}
                            fmt.Println(v6, v35)
                            _ = v143
                        } else {
                            var v144 float32 = 1.5
                            var v145 uint64 = (v1 >> ((v5) & 63))
                            _ = v144
                            _ = v145
                        }
                        var v146 int = fi137
                        _ = v146
                    }
                    _ = fi139
                    _ = v142
                }
            case "foo bar":
                fmt.Println(fi137)
            }
            v35 = v35
            _ = fi137
        }
        for fi147 := 0; fi147 < 4; fi147++ {
            var v148 uint8 = v6
            switch 2147483647 {
            case -787:
                for fi149 := 0; fi149 < 4; fi149++ {
                    v5 = (v5 - (^(v5)))
                    _ = fi149
                }
            }
            v148 -= (v148 ^ 128)
            _ = fi147
            _ = v148
        }
        var v150 int16 = 1
        _ = v150
    }
    fmt.Println(v5, v4)
    v151 := GenS17{F0: 634, F1: v1}
    fmt.Println(v6, v6)
    v4 = (v5 <= v5)
    v2[128] = 32767
    v152 := GenS18{F0: "", F1: -162, F2: (v5 >> (((^(v5))) & 63))}
    v5 = v5
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v151
    _ = v152
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
