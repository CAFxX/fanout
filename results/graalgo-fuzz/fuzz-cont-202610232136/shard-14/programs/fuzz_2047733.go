package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 uint64; F2 int64 }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 bool; F1 string }
type GenS5 struct { F0 int8; F1 uint; F2 int8 }
type GenS6 struct { F0 string; F1 string }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint; F1 int8 }
type GenS11 struct { F0 uint16; F1 uint; F2 int64 }
type GenS12 struct { F0 int }
type GenS13 struct { F0 uint64; F1 int16 }
type GenS14 struct { F0 int16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = (-80 == -37)
    var v3 int8 = -1
    var v4 int16 = 394
    var v5 float32 = ((((0.0 + 0.0) * 0.0) / 1.0) - float32((1.5 - 1e10)))
    var v6 uint64 = 1
    v2 = (!(v2 || (v2 && v2)))
    v4 = (v4 - (^(v4)))
    for fi7 := 0; fi7 < 7; fi7++ {
        v4 = -1
        fmt.Println(v4)
        _ = fi7
    }
    if true {
        var v8 uint = 1
        _ = v8
    }
    if (!v2) {
        for fi9 := 0; fi9 < 16; fi9++ {
            v10 := GenS1{F0: 574}
            switch ("hello" + ("\n\t" + (("hello" + "a") + "a"))) {
            case "\n\t":
                switch (-(fi9)) {
                case 128:
                    var v11 float32 = v5
                    _ = v11
                case 4294967295:
                    switch fi9 {
                    case -133:
                        var v12 uint64 = (v6 >> ((721) & 63))
                        _ = v12
                    case 127:
                        var v13 int32 = 1
                        _ = v13
                    default:
                        v14 := GenS2{F0: 737, F1: (v6 >> ((42) & 63)), F2: 0}
                        _ = v14
                    }
                case -384:
                    var v15 bool = (!(!v2))
                    _ = v15
                default:
                    var v16 int = (fi9 | (fi9 ^ fi9))
                    _ = v16
                }
            case "foo bar":
                fmt.Println(v4, v6, fi9)
            default:
                var v17 bool = v2
                _ = v17
            }
            _ = fi9
            _ = v10
        }
        if ((0.0 - 0.0) == 0.1) {
            for fi18 := 0; fi18 < 11; fi18++ {
                fmt.Println(v6, v4)
                if (!true) {
                    var v19 int64 = -9223372036854775808
                    var v20 uint64 = (v6 % ((v6) | 1))
                    for fi21 := 0; fi21 < 20; fi21++ {
                        var v22 float32 = v5
                        var v23 int = fi21
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                    _ = v19
                    _ = v20
                }
                for fi24 := 0; fi24 < 17; fi24++ {
                    v25 := GenS1{F0: 10}
                    _ = fi24
                    _ = v25
                }
                _ = fi18
            }
            var v26 bool = ((v2 || ((-9223372036854775808 > -9223372036854775808) || (!((220 <= (v6 % ((v6) | 1))) && false)))) && v2)
            switch (("hello" + ((("\n\t" + "") + ("" + "")) + ((("hello" + ("" + "\n\t")) + "foo bar") + ""))) + "") {
            case "\n\t":
                _ = len(v1)
            }
            _ = v26
        } else {
            v27 := []int{}
            if (255 < 42) {
                var v28 uint16 = 255
                _ = v28
            } else {
                var v29 int32 = -531
                var v30 int = 2147483647
                _ = v29
                _ = v30
            }
            _ = v27
        }
    } else {
        v6 += v6
        for fi31 := 0; fi31 < 4; fi31++ {
            v32 := GenS3{F0: 1, F1: "\n\t"}
            _ = fi31
            _ = v32
        }
    }
    v33 := GenS4{F0: (-9223372036854775808 == 0), F1: "\n\t"}
    if (v2 && false) {
        if true {
            fmt.Println(v4, v5, v3)
            v34 := GenS3{F0: int(v4), F1: ("a" + (("" + (("a" + ("" + ((("hello" + "a") + ("foo bar" + "foo bar")) + ("a" + "")))) + "\n\t")) + "hello"))}
            var v35 uint8 = 39
            _ = v34
            _ = v35
        } else {
            var v36 uint32 = 2
            v5 -= v5
            var v37 uint = 4294967295
            _ = v36
            _ = v37
        }
        v38 := []int{}
        _ = v38
    } else {
        var v39 int16 = 32767
        if false {
            v40 := GenS1{F0: 2147483648}
            for fi41 := 0; fi41 < 5; fi41++ {
                for fi42 := 0; fi42 < 1; fi42++ {
                    var v43 int8 = (v3 ^ v3)
                    fmt.Println(v43, fi42, v2)
                    fmt.Println(v3, fi42, v5)
                    _ = fi42
                    _ = v43
                }
                v4 = (v4 | (v4 * (^(v4))))
                _ = fi41
            }
            var v44 int64 = -916
            _ = v40
            _ = v44
        } else {
            if (255 == 642) {
                if v2 {
                    switch "foo bar" {
                    case "a":
                        for fi45 := 0; fi45 < 12; fi45++ {
                            var v46 float32 = float32(-2.25)
                            var v47 int32 = -2147483648
                            v48 := []int{(fi45 >> ((0) & 63))}
                            _ = fi45
                            _ = v46
                            _ = v47
                            _ = v48
                        }
                    default:
                        if ((!true) || (true || v2)) {
                            for fi49 := 0; fi49 < 17; fi49++ {
                                if v2 {
                                    var v50 int64 = -1
                                    for fi51 := 0; fi51 < 5; fi51++ {
                                        var v52 uint8 = 42
                                        var v53 int8 = 52
                                        _ = fi51
                                        _ = v52
                                        _ = v53
                                    }
                                    _ = v50
                                } else {
                                    fmt.Println(v3, v6)
                                }
                                v54 := []int{}
                                _ = fi49
                                _ = v54
                            }
                        } else {
                            v55 := GenS5{F0: v3, F1: 888, F2: v3}
                            var v56 uint64 = v6
                            _ = v55
                            _ = v56
                        }
                    }
                } else {
                    if true {
                        v57 := []int{-2147483648, 0, 32767}
                        if v2 {
                            if len(v57) >= 2 {
                                v57 = v57[1:len(v57)]
                            }
                        }
                        _ = v57
                    } else {
                        v3 = v3
                    }
                    var v58 int32 = int32(v6)
                    _ = v58
                }
                if true {
                    var v59 uint8 = 125
                    if v2 {
                        if true {
                            var v60 int16 = 1000
                            if (v2 && v2) {
                                _ = len(v1)
                                v3 = v3
                                for fi61 := 0; fi61 < 16; fi61++ {
                                    if v2 {
                                        v62 := GenS6{F0: "\n\t", F1: ("foo bar" + "hello")}
                                        v63 := []int{9223372036854775807}
                                        if false {
                                            v2 = (42 <= 1)
                                        }
                                        _ = v62
                                        _ = v63
                                    }
                                    fmt.Println(v60)
                                    if (v2 && (!(v2 && (v2 && v2)))) {
                                        _ = len(v1)
                                        var v64 uint16 = 42
                                        _ = v64
                                    }
                                    _ = fi61
                                }
                            } else {
                                for fi65 := 0; fi65 < 10; fi65++ {
                                    if v2 {
                                        if false {
                                            if v2 {
                                                for fi66 := 0; fi66 < 19; fi66++ {
                                                    v60 = (v60 - (-(v39)))
                                                    _ = fi66
                                                }
                                                for fi67 := 0; fi67 < 3; fi67++ {
                                                    if v2 {
                                                        fmt.Println(v6)
                                                        switch 4294967295 {
                                                        case 128:
                                                            var v68 int16 = (v4 - v60)
                                                            _ = v68
                                                        case -2147483648:
                                                            var v69 int32 = 133
                                                            _ = v69
                                                        case 127:
                                                            v6 += v6
                                                            var v70 uint16 = uint16(v59)
                                                            var v71 float64 = float64(v5)
                                                            _ = v70
                                                            _ = v71
                                                        default:
                                                            if ((v6 | (v6 >> ((1) & 63))) > 2) {
                                                                var v72 uint16 = 42
                                                                fmt.Println(v72, v39)
                                                                v73 := GenS5{F0: 9, F1: 2, F2: -1}
                                                                _ = v72
                                                                _ = v73
                                                            } else {
                                                                v74 := GenS3{F0: fi67, F1: "foo bar"}
                                                                _ = v74
                                                            }
                                                        }
                                                    }
                                                    _ = fi67
                                                }
                                            } else {
                                                switch 924 {
                                                case 4294967295:
                                                    var v75 int = 32767
                                                    _ = v75
                                                case -137:
                                                    v5 = float32(0.1)
                                                default:
                                                    switch ("" + ("\n\t" + ("a" + "\n\t"))) {
                                                    case "":
                                                        for fi76 := 0; fi76 < 10; fi76++ {
                                                            switch "hello" {
                                                            case "":
                                                                v39 -= v60
                                                            case "foo bar":
                                                                v39 -= (v4 << ((42) & 63))
                                                            }
                                                            var v77 float64 = -2.25
                                                            v78 := GenS7{F0: (v59 << ((42) & 63))}
                                                            _ = fi76
                                                            _ = v77
                                                            _ = v78
                                                        }
                                                    case "foo bar":
                                                        if (v2 || (0 <= 18446744073709551615)) {
                                                            fmt.Println(v39)
                                                        }
                                                    case "hello":
                                                        for fi79 := 0; fi79 < 4; fi79++ {
                                                            var v80 float64 = (0.1 / (0.0 + -2.25))
                                                            v81 := []int{(fi65 << ((65535) & 63)), (-(fi79)), fi65, fi65}
                                                            _ = fi79
                                                            _ = v80
                                                            _ = v81
                                                        }
                                                    default:
                                                        fmt.Println(v2)
                                                    }
                                                }
                                            }
                                        }
                                    }
                                    _ = fi65
                                }
                            }
                            for fi82 := 0; fi82 < 13; fi82++ {
                                v83 := []int{(-(fi82)), fi82}
                                _ = fi82
                                _ = v83
                            }
                            _ = v60
                        }
                    } else {
                        v6 = v6
                        for fi84 := 0; fi84 < 5; fi84++ {
                            v4 = v4
                            if v2 {
                                var v85 int = fi84
                                fmt.Println(v59, v2)
                                _ = len(v1)
                                _ = v85
                            } else {
                                v59 += v59
                                if v2 {
                                    if v2 {
                                        v86 := []int{}
                                        if len(v86) >= 2 {
                                            v86 = v86[1:len(v86)]
                                        }
                                        v2 = true
                                        _ = v86
                                    }
                                } else {
                                    fmt.Println(v3, v59)
                                }
                                if v2 {
                                    fmt.Println(v5, v59, v2)
                                    v87 := GenS7{F0: (v59 + v59)}
                                    _ = v87
                                } else {
                                    v88 := GenS8{F0: v3}
                                    for fi89 := 0; fi89 < 11; fi89++ {
                                        if true {
                                            switch "hello" {
                                            case "a":
                                                switch ("foo bar" + "foo bar") {
                                                case "foo bar":
                                                    v90 := GenS4{F0: v2, F1: "foo bar"}
                                                    _ = v90
                                                case "\n\t":
                                                    v39 -= v39
                                                    fmt.Println(v39)
                                                    v2 = (449 <= 42)
                                                    if true {
                                                        v91 := GenS2{F0: 1, F1: v6, F2: 0}
                                                        if (39 >= -18) {
                                                            fmt.Println(fi89, v39)
                                                            var v92 uint = 2
                                                            v93 := v1["\n\t"]
                                                            _ = v92
                                                            _ = v93
                                                        } else {
                                                            fmt.Println(v3, v5, fi84)
                                                            v2 = (9223372036854775807 < 997)
                                                            v94 := GenS9{F0: (-(v3))}
                                                            var v95 float64 = 1.5
                                                            _ = v94
                                                            _ = v95
                                                        }
                                                        v96 := GenS5{F0: 78, F1: 355, F2: -1}
                                                        _ = v91
                                                        _ = v96
                                                    }
                                                case "a":
                                                    var v97 float64 = (((((0.1 - float64(float32((-2.25 + 1e10)))) / 1e10) - (0.0 + 1e10)) - (1.5 + (float64(v5) / 1.0))) + 0.0)
                                                    _ = v97
                                                }
                                            default:
                                                var v98 int16 = 0
                                                _ = v98
                                            }
                                            v2 = true
                                            v99 := []int{fi89, int(v3)}
                                            _ = v99
                                        } else {
                                            switch "foo bar" {
                                            case "":
                                                switch "a" {
                                                case "":
                                                    var v100 int = 712
                                                    _ = v100
                                                }
                                            case "a":
                                                v1["hello"] = 4294967295
                                            }
                                        }
                                        v101 := []int{fi89}
                                        switch "a" {
                                        case "\n\t":
                                            fmt.Println(v59, v59, fi89)
                                        case "a":
                                            switch ("a" + "") {
                                            case "\n\t":
                                                v102 := GenS8{F0: v3}
                                                _ = v102
                                            case "foo bar":
                                                if v2 {
                                                    if (528 > 65535) {
                                                        var v103 int64 = int64(v59)
                                                        var v104 int16 = (-(v39))
                                                        _ = v103
                                                        _ = v104
                                                    }
                                                    var v105 uint8 = v59
                                                    var v106 float32 = 1.5
                                                    _ = v105
                                                    _ = v106
                                                }
                                            default:
                                                _ = len(v101)
                                            }
                                        case "foo bar":
                                            v107 := GenS8{F0: (v3 | (v3 | 100))}
                                            _ = v107
                                        default:
                                            var v108 uint64 = (v6 / ((v6) | 1))
                                            _ = v108
                                        }
                                        _ = fi89
                                        _ = v101
                                    }
                                    v2 = v2
                                    _ = v88
                                }
                            }
                            var v109 int8 = 100
                            _ = fi84
                            _ = v109
                        }
                        for fi110 := 0; fi110 < 8; fi110++ {
                            v111 := []int{}
                            var v112 int16 = v39
                            var v113 string = ("\n\t" + "a")
                            _ = fi110
                            _ = v111
                            _ = v112
                            _ = v113
                        }
                    }
                    if (v2 || (float32(float64(1.5)) > float32(1e10))) {
                        if v2 {
                            var v114 string = ("a" + ("hello" + ("a" + "\n\t")))
                            v114 = "foo bar"
                            if (v2 && v2) {
                                for fi115 := 0; fi115 < 17; fi115++ {
                                    for fi116 := 0; fi116 < 3; fi116++ {
                                        v117 := GenS6{F0: "foo bar", F1: "\n\t"}
                                        if (v6 == v6) {
                                            if ((-(v4)) > (v39 * v39)) {
                                                v114 = v114
                                                if (2 > 51) {
                                                    if (156 < 396) {
                                                        v6 -= v6
                                                    } else {
                                                        v1[v114] = 407
                                                        var v118 int = fi116
                                                        v59 = (v59 ^ 127)
                                                        v119 := GenS10{F0: 65535, F1: (v3 | v3)}
                                                        var v120 uint8 = (v59 + (v59 | 95))
                                                        _ = v118
                                                        _ = v119
                                                        _ = v120
                                                    }
                                                } else {
                                                    if (v2 && v2) {
                                                        fmt.Println(v6)
                                                    } else {
                                                        var v121 bool = true
                                                        v1[v114] = uint32(v59)
                                                        _ = v121
                                                    }
                                                    fmt.Println(fi115)
                                                    v5 = v5
                                                    v122 := GenS11{F0: 32767, F1: 42, F2: 954}
                                                    switch v114 {
                                                    case "\n\t":
                                                        var v123 float64 = 0.1
                                                        _ = v123
                                                    default:
                                                        v124 := []int{(fi115 * (fi116 % ((fi116) | 1)))}
                                                        _ = v124
                                                    }
                                                    _ = v122
                                                }
                                            }
                                            if (1e10 == (1e10 * 0.0)) {
                                                switch fi115 {
                                                case 0:
                                                    if v2 {
                                                        var v125 bool = (v2 && v2)
                                                        _ = v125
                                                    } else {
                                                        v1[((v114 + ("hello" + v114)) + v114)] = 1
                                                        var v126 float64 = (-2.25 / 1.5)
                                                        var v127 int16 = (v4 % (((v39 - v4)) | 1))
                                                        _ = v126
                                                        _ = v127
                                                    }
                                                    v2 = true
                                                    v39 -= v4
                                                case -291:
                                                    _ = len(v1)
                                                }
                                            } else {
                                                var v128 int16 = (v39 ^ (v39 % ((561) | 1)))
                                                _ = v128
                                            }
                                            v129 := GenS12{F0: (fi116 | 90)}
                                            var v130 uint8 = (v59 / ((v59) | 1))
                                            v131 := GenS6{F0: v114, F1: v114}
                                            _ = v129
                                            _ = v130
                                            _ = v131
                                        }
                                        fmt.Println(v4)
                                        _ = fi116
                                        _ = v117
                                    }
                                    _ = fi115
                                }
                                if (false || v2) {
                                    v132 := v1["a"]
                                    _ = v132
                                }
                                v133 := []int{}
                                _ = v133
                            }
                            _ = v114
                        } else {
                            v1[(("a" + ("" + ("a" + (("" + "hello") + "a")))) + (("" + "hello") + ""))] = 368
                            var v134 int = 4294967295
                            for fi135 := 0; fi135 < 4; fi135++ {
                                fmt.Println(v3, fi135)
                                v136 := GenS13{F0: v6, F1: v39}
                                _ = fi135
                                _ = v136
                            }
                            _ = v134
                        }
                        fmt.Println(v6, v59)
                        switch -2147483648 {
                        case 32768:
                            var v137 int8 = 1
                            _ = v137
                        case 128:
                            v2 = ((127 <= -1) || true)
                        case 1:
                            var v138 uint8 = 98
                            _ = v138
                        default:
                            for fi139 := 0; fi139 < 18; fi139++ {
                                var v140 float64 = ((1e10 / 1.0) / 1e10)
                                var v141 uint = 255
                                _ = fi139
                                _ = v140
                                _ = v141
                            }
                        }
                    } else {
                        for fi142 := 0; fi142 < 20; fi142++ {
                            var v143 bool = v2
                            if true {
                                var v144 int64 = -786
                                v39 += v4
                                _ = v144
                            }
                            _ = fi142
                            _ = v143
                        }
                    }
                    _ = v59
                } else {
                    delete(v1, "")
                    var v145 int = 0
                    fmt.Println(v6)
                    _ = v145
                }
            } else {
                v146 := []int{-945}
                v147 := GenS9{F0: v3}
                _ = v146
                _ = v147
            }
        }
        _ = len(v1)
        _ = v39
    }
    v148 := GenS14{F0: v4}
    v6 = v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v148
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
