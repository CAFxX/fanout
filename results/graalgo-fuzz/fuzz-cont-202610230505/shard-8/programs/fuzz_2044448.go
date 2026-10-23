package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 int }
type GenS2 struct { F0 int16; F1 int; F2 int16 }
type GenS3 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS4 struct { F0 uint32; F1 uint16; F2 int }
type GenS5 struct { F0 string }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int8; F1 bool; F2 int32 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint8; F1 uint64 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 uint8; F1 bool; F2 uint8 }
type GenS13 struct { F0 int; F1 int32 }
type GenS14 struct { F0 int; F1 string }

func main() {
    var v1 int32 = 256
    var v2 uint8 = 53
    var v3 int16 = 755
    var v4 uint32 = 608
    v5 := make(map[string]bool)
    var v6 int8 = 95
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 uint64 = 974
        if ((2 > v8) || (!false)) {
            v9 := GenS1{F0: v4, F1: 813, F2: fi7}
            _ = v9
        } else {
            for fi10 := 0; fi10 < 9; fi10++ {
                switch ("foo bar" + ("" + "foo bar")) {
                case "\n\t":
                    v4 = v4
                case "":
                    var v11 float32 = 1.5
                    _ = v11
                }
                fmt.Println(v8)
                fmt.Println(v2, v6)
                _ = fi10
            }
        }
        _ = fi7
        _ = v8
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        v4 = 0
        var v13 uint16 = 991
        _ = fi12
        _ = v13
    }
    if true {
        for fi14 := 0; fi14 < 3; fi14++ {
            if (1 != -9223372036854775808) {
                fmt.Println(v4, v3, v6)
            } else {
                for fi15 := 0; fi15 < 3; fi15++ {
                    var v16 uint8 = v2
                    var v17 int8 = v6
                    switch ("foo bar" + "a") {
                    case "foo bar":
                        switch "a" {
                        case "hello":
                            v18 := GenS2{F0: v3, F1: (fi15 / (((fi15 & fi14)) | 1)), F2: v3}
                            _ = v18
                        case "a":
                            var v19 uint64 = 4294967295
                            _ = v19
                        case "foo bar":
                            var v20 string = ("hello" + "hello")
                            _ = v20
                        }
                        v21 := GenS1{F0: (v4 & 0), F1: 87, F2: fi14}
                        _ = v21
                    default:
                        if (42 < v17) {
                            v22 := []int{1, (fi14 | (fi14 * (fi14 - fi14))), (-(fi14))}
                            v23 := GenS1{F0: (v4 << ((4294967295) & 63)), F1: 18446744073709551615, F2: 127}
                            var v24 uint32 = v4
                            _ = v22
                            _ = v23
                            _ = v24
                        } else {
                            v1 = v1
                        }
                    }
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                var v25 int16 = int16(v2)
                if false {
                    var v26 uint = 4294967295
                    v27 := GenS3{F0: (v26 / (((v26 - 255)) | 1)), F1: 32767, F2: 2}
                    _ = v26
                    _ = v27
                } else {
                    v28 := GenS3{F0: 198, F1: 65535, F2: 18446744073709551615}
                    var v29 uint8 = v2
                    fmt.Println(v6, v1)
                    _ = v28
                    _ = v29
                }
                _ = v25
            }
            if true {
                fmt.Println(fi14, v1, v2)
            } else {
                var v30 float64 = -2.25
                switch -566 {
                case 4294967295:
                    v6 = v6
                case -2147483648:
                    v5[("" + "foo bar")] = true
                default:
                    if ((v4 | (v4 & 934)) >= v4) {
                        switch (fi14 + (fi14 & (-(fi14)))) {
                        case -758:
                            v2 += (v2 % (((v2 / (((v2 - v2)) | 1))) | 1))
                        default:
                            if (((fi14 & fi14) != fi14) || false) {
                                if false {
                                    if true {
                                        v1 += v1
                                        v31 := []int{fi14, 631}
                                        _ = v31
                                    }
                                    v32 := GenS4{F0: v4, F1: 2, F2: fi14}
                                    if true {
                                        v33 := GenS2{F0: (v3 * (v3 + (v3 * v3))), F1: 32767, F2: (v3 | v3)}
                                        delete(v5, ("" + ("\n\t" + "\n\t")))
                                        var v34 uint = 65535
                                        _ = v33
                                        _ = v34
                                    }
                                    _ = v32
                                } else {
                                    switch (fi14 * fi14) {
                                    case 32767:
                                        var v35 string = ""
                                        _ = v35
                                    default:
                                        v2 = v2
                                    }
                                }
                            }
                        }
                        v3 = v3
                        var v36 int = 174
                        _ = v36
                    } else {
                        if (false && (!true)) {
                            v2 += v2
                            switch (fi14 / (((fi14 >> ((42) & 63))) | 1)) {
                            case -970:
                                for fi37 := 0; fi37 < 7; fi37++ {
                                    _ = len(v5)
                                    v30 += v30
                                    if (!true) {
                                        var v38 int16 = v3
                                        v39 := GenS5{F0: "a"}
                                        switch (fi14 << ((408) & 63)) {
                                        case -1000:
                                            var v40 uint = 1
                                            _ = v40
                                        case -2147483648:
                                            v30 += ((1.5 * 1.5) - 1e10)
                                        default:
                                            switch (fi37 ^ -768) {
                                            case 1:
                                                if (((1 <= v38) || true) || false) {
                                                    var v41 string = "foo bar"
                                                    v3 += -399
                                                    if true {
                                                        var v42 uint = 2
                                                        v43 := []int{fi14}
                                                        _ = v42
                                                        _ = v43
                                                    }
                                                    _ = v41
                                                } else {
                                                    var v44 bool = false
                                                    switch (fi37 % (((fi14 % ((fi14) | 1))) | 1)) {
                                                    case -1:
                                                        v2 = v2
                                                    case 32767:
                                                        switch "\n\t" {
                                                        case "a":
                                                            fmt.Println(v1, v30, v4)
                                                        case "\n\t":
                                                            if false {
                                                                fmt.Println(fi14)
                                                            } else {
                                                                v45 := GenS6{F0: ((("foo bar" + (("\n\t" + "\n\t") + "a")) + "\n\t") + "hello")}
                                                                var v46 uint = uint(v4)
                                                                _ = v45
                                                                _ = v46
                                                            }
                                                        }
                                                    case -2147483648:
                                                        switch (fi14 + fi37) {
                                                        case -9223372036854775808:
                                                            var v47 uint = 42
                                                            _ = v47
                                                        case 42:
                                                            var v48 uint32 = (v4 << ((874) & 63))
                                                            _ = v48
                                                        case 194:
                                                            delete(v5, "\n\t")
                                                        }
                                                    default:
                                                        fmt.Println(v3, v3)
                                                    }
                                                    var v49 uint = 633
                                                    _ = v44
                                                    _ = v49
                                                }
                                            case 32767:
                                                v50 := GenS7{F0: (v6 % (((v6 >> ((uint(v4)) & 63))) | 1)), F1: (!(((true && (0 < 1000)) && false) || true)), F2: v1}
                                                _ = v50
                                            default:
                                                fmt.Println(v4, v1)
                                            }
                                        }
                                        _ = v38
                                        _ = v39
                                    } else {
                                        var v51 string = (("hello" + "\n\t") + ("foo bar" + (((("hello" + "") + (("" + "\n\t") + "a")) + (("a" + "a") + ("foo bar" + "hello"))) + "a")))
                                        _ = v51
                                    }
                                    _ = fi37
                                }
                            case 0:
                                var v52 uint64 = 7
                                _ = v52
                            default:
                                for fi53 := 0; fi53 < 4; fi53++ {
                                    switch ("foo bar" + "foo bar") {
                                    case "foo bar":
                                        var v54 uint64 = 2
                                        _ = v54
                                    case "a":
                                        var v55 bool = true
                                        _ = v55
                                    case "\n\t":
                                        v4 = 2147483648
                                    }
                                    var v56 uint64 = 696
                                    _ = fi53
                                    _ = v56
                                }
                            }
                        } else {
                            v57 := GenS8{F0: (("hello" + "a") + "foo bar")}
                            var v58 int = (fi14 ^ -329)
                            _ = v57
                            _ = v58
                        }
                    }
                }
                switch 127 {
                case -2147483648:
                    fmt.Println(v2)
                case 4294967295:
                    if true {
                        for fi59 := 0; fi59 < 2; fi59++ {
                            v6 = (v6 | v6)
                            if false {
                                v3 -= (v3 << ((740) & 63))
                                v60 := []int{537, (-(fi14)), fi59, (fi59 ^ fi59)}
                                v3 = v3
                                if (!((-(fi59)) > 867)) {
                                    v4 = v4
                                    v61 := GenS9{F0: (v2 >> ((0) & 63)), F1: 376}
                                    v62 := GenS3{F0: 2, F1: 824, F2: 4294967295}
                                    _ = v61
                                    _ = v62
                                } else {
                                    v63 := GenS6{F0: "hello"}
                                    if (((v2 - v2) >= v2) && (!(2 > 255))) {
                                        if true {
                                            if (int64(fi14) >= 9223372036854775807) {
                                                v6 += (v6 & v6)
                                                if true {
                                                    var v64 uint32 = (v4 | (v4 * (v4 >> ((1) & 63))))
                                                    switch "a" {
                                                    case "\n\t":
                                                        v65 := GenS10{F0: 32768}
                                                        _ = v65
                                                    case "foo bar":
                                                        var v66 string = ""
                                                        var v67 float32 = (0.1 / 0.1)
                                                        _ = v66
                                                        _ = v67
                                                    }
                                                    if len(v60) >= 2 {
                                                        v60 = v60[1:len(v60)]
                                                    }
                                                    fmt.Println(v1)
                                                    _ = v64
                                                }
                                                fmt.Println(v6, v2, v1)
                                                v68 := GenS11{F0: 2}
                                                fmt.Println(v4, v4)
                                                _ = v68
                                            }
                                        } else {
                                            var v69 int16 = v3
                                            if false {
                                                v70 := GenS11{F0: uint64(v2)}
                                                fmt.Println(fi14)
                                                _ = v70
                                            }
                                            _ = v69
                                        }
                                        var v71 string = ("\n\t" + "\n\t")
                                        _ = v71
                                    } else {
                                        switch fi59 {
                                        case 32767:
                                            v72 := GenS2{F0: (v3 & 32767), F1: (^(fi59)), F2: v3}
                                            _ = v72
                                        case 32768:
                                            if len(v60) > 0 {
                                                v60[0] = fi14
                                            }
                                        case -766:
                                            v4 -= 365
                                            v4 = (v4 << ((65535) & 63))
                                        }
                                        var v73 bool = (1 >= v6)
                                        var v74 int32 = v1
                                        delete(v5, "foo bar")
                                        if ((v3 < v3) || v73) {
                                            fmt.Println(fi14, fi14, fi59)
                                            if (1e10 <= -2.25) {
                                                fmt.Println(fi59)
                                                if v73 {
                                                    var v75 uint32 = v4
                                                    _ = v75
                                                } else {
                                                    v74 += v74
                                                    fmt.Println(fi14, v30, v4)
                                                    v76 := GenS7{F0: (-(v6)), F1: v73, F2: v74}
                                                    _ = v76
                                                }
                                            } else {
                                                var v77 uint64 = 4294967295
                                                _ = v77
                                            }
                                            var v78 int = fi59
                                            _ = v78
                                        }
                                        _ = v73
                                        _ = v74
                                    }
                                    _ = v63
                                }
                                switch "foo bar" {
                                case "\n\t":
                                    fmt.Println(v2)
                                case "foo bar":
                                    var v79 uint64 = 284
                                    _ = v79
                                default:
                                    if false {
                                        var v80 uint8 = 111
                                        _ = v80
                                    }
                                }
                                _ = v60
                            }
                            v81 := []int{int(v1), 4294967295, (fi59 & (fi59 + -726)), -2147483648}
                            var v82 int16 = (v3 | 145)
                            v82 -= v3
                            _ = fi59
                            _ = v81
                            _ = v82
                        }
                        for fi83 := 0; fi83 < 6; fi83++ {
                            var v84 bool = false
                            var v85 uint64 = 0
                            _ = fi83
                            _ = v84
                            _ = v85
                        }
                        var v86 string = ("hello" + (("hello" + "a") + ""))
                        _ = v86
                    } else {
                        switch "foo bar" {
                        case "a":
                            var v87 uint16 = 0
                            _ = v87
                        }
                        fmt.Println(v3, v30, v3)
                        var v88 int32 = 0
                        _ = v88
                    }
                default:
                    switch ("\n\t" + "foo bar") {
                    case "hello":
                        v4 = 14
                    }
                }
                _ = v30
            }
            _ = fi14
        }
        if (true && false) {
            fmt.Println(v2, v4)
            fmt.Println(v3, v6)
        }
        var v89 uint64 = 0
        _ = v89
    }
    v90 := []int{127}
    var v91 int = 2147483647
    for fi92 := 0; fi92 < 4; fi92++ {
        var v93 int64 = 9223372036854775807
        for fi94 := 0; fi94 < 19; fi94++ {
            if ((v1 + v1) > v1) {
                v4 += v4
            } else {
                v95 := v5["foo bar"]
                if (407 == 32767) {
                    _ = len(v5)
                } else {
                    v96 := GenS8{F0: ""}
                    v2 += v2
                    v97 := v5["a"]
                    if false {
                        var v98 uint = 636
                        v99 := GenS8{F0: ""}
                        _ = v98
                        _ = v99
                    }
                    v90 = append(v90, (v91 / ((v91) | 1)))
                    _ = v96
                    _ = v97
                }
                _ = v95
            }
            _ = fi94
        }
        v3 += 42
        _ = fi92
        _ = v93
    }
    switch "a" {
    case "":
        _ = len(v5)
    default:
        var v100 int64 = -9223372036854775808
        _ = v100
    }
    if true {
        v4 -= (^(v4))
        v91 -= (v91 % (((v91 | (v91 ^ 0))) | 1))
        if false {
            v101 := GenS1{F0: uint32(v2), F1: 18446744073709551615, F2: (^(v91))}
            switch v91 {
            case 42:
                if true {
                    v91 += (v91 << ((2) & 63))
                    v2 -= 128
                } else {
                    v102 := GenS8{F0: "hello"}
                    v4 = (v4 % (((v4 - 2147483647)) | 1))
                    var v103 uint8 = v2
                    _ = v102
                    _ = v103
                }
            case 0:
                if true {
                    var v104 int32 = 0
                    switch (v91 | 0) {
                    case 4294967295:
                        v104 += v1
                    case 9223372036854775807:
                        v1 = int32(v2)
                    case 127:
                        v105 := GenS10{F0: v91}
                        _ = v105
                    }
                    var v106 int64 = -9223372036854775808
                    _ = v104
                    _ = v106
                } else {
                    v1 = v1
                }
            case -9223372036854775808:
                v5["foo bar"] = (!true)
            default:
                v3 = (v3 & (v3 << ((65535) & 63)))
            }
            _ = v101
        } else {
            for fi107 := 0; fi107 < 19; fi107++ {
                fmt.Println(v91, v4)
                _ = fi107
            }
            var v108 bool = (v3 < -1)
            if (v108 && true) {
                for fi109 := 0; fi109 < 15; fi109++ {
                    v110 := GenS12{F0: v2, F1: (!(v91 == 1)), F2: 0}
                    if ((!false) && v108) {
                        if (!true) {
                            v111 := GenS2{F0: v3, F1: v91, F2: 1000}
                            _ = v111
                        }
                    }
                    if len(v90) >= 2 {
                        v90 = v90[1:len(v90)]
                    }
                    _ = fi109
                    _ = v110
                }
                switch (v91 | v91) {
                case 128:
                    var v112 int16 = v3
                    _ = v112
                case 9223372036854775807:
                    v113 := GenS6{F0: "a"}
                    _ = v113
                default:
                    var v114 uint64 = 0
                    _ = v114
                }
                v115 := GenS13{F0: (v91 % (((v91 - (v91 << ((1) & 63)))) | 1)), F1: int32(v6)}
                _ = v115
            }
            _ = v108
        }
    } else {
        for fi116 := 0; fi116 < 4; fi116++ {
            var v117 int8 = (v6 ^ (v6 >> ((65535) & 63)))
            _ = fi116
            _ = v117
        }
    }
    if ((^(v3)) < v3) {
        if len(v90) >= 2 {
            v90 = v90[1:len(v90)]
        }
        var v118 float64 = 3.14159
        var v119 int8 = (^(v6))
        _ = v118
        _ = v119
    } else {
        if (!false) {
            _ = len(v5)
            for fi120 := 0; fi120 < 7; fi120++ {
                if (!true) {
                    if ("hello" != "a") {
                        if false {
                            v4 = (^(v4))
                            fmt.Println(v1, v6)
                            for fi121 := 0; fi121 < 10; fi121++ {
                                v2 = (^(v2))
                                var v122 uint = 255
                                v122 = 255
                                _ = fi121
                                _ = v122
                            }
                        } else {
                            _ = len(v90)
                            v1 = -1
                        }
                        var v123 int = 128
                        v2 = v2
                        _ = v123
                    }
                    for fi124 := 0; fi124 < 5; fi124++ {
                        fmt.Println(fi120, v6, v4)
                        v5[""] = (v1 > v1)
                        _ = fi124
                    }
                    var v125 int = 32767
                    _ = v125
                }
                var v126 bool = ((v3 >> ((255) & 63)) > 32767)
                _ = fi120
                _ = v126
            }
        }
        v127 := GenS3{F0: 255, F1: 0, F2: 773}
        _ = v127
    }
    var v128 uint64 = 979
    v128 += (v128 * uint64(v2))
    v2 = v2
    for fi129 := 0; fi129 < 1; fi129++ {
        delete(v5, "a")
        for fi130 := 0; fi130 < 18; fi130++ {
            v5["foo bar"] = (uint64(v4) < (^(v128)))
            var v131 int = -1
            _ = fi130
            _ = v131
        }
        _ = fi129
    }
    v91 -= v91
    if (v3 == int16(v91)) {
        if false {
            if (false || (!true)) {
                v5[("hello" + ("a" + ("a" + "\n\t")))] = false
                if len(v90) >= 2 {
                    v90 = v90[1:len(v90)]
                }
            } else {
                for fi132 := 0; fi132 < 1; fi132++ {
                    var v133 float64 = (0.0 + 0.0)
                    _ = fi132
                    _ = v133
                }
            }
        } else {
            var v134 float64 = 1e10
            for fi135 := 0; fi135 < 8; fi135++ {
                var v136 float32 = float32(float64(-2.25))
                if ((v134 + 1e10) > float64(v136)) {
                    delete(v5, ("" + "foo bar"))
                } else {
                    var v137 string = ((("\n\t" + "\n\t") + "hello") + (((("a" + "\n\t") + "foo bar") + "\n\t") + "a"))
                    _ = v137
                }
                for fi138 := 0; fi138 < 13; fi138++ {
                    if (692 >= 32) {
                        v139 := GenS14{F0: int(v6), F1: "hello"}
                        _ = v139
                    }
                    _ = fi138
                }
                _ = fi135
                _ = v136
            }
            _ = v134
        }
        var v140 uint16 = 32767
        _ = v140
    } else {
        v141 := GenS12{F0: (v2 / ((v2) | 1)), F1: ((v91 | v91) != int(v128)), F2: v2}
        _ = v141
    }
    _ = len(v90)
    fmt.Println("v1:", v1)
    fmt.Println("v128:", v128)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v90:", len(v90), v90)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v128
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v90
    _ = v91
}
