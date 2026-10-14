package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 uint64 }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 string; F1 int64; F2 string }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint64; F1 int64 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 string; F1 uint; F2 int8 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 int64; F1 uint }
type GenS12 struct { F0 int; F1 uint; F2 bool }
type GenS13 struct { F0 string }
func genFn0() uint {
    switch 566 {
    case -294:
        v5 := []int{403, 32767}
        _ = v5
    case 32767:
        v6 := GenS1{F0: 42, F1: 4294967295, F2: 18446744073709551615}
        _ = v6
    }
    var v7 uint = 9
    v7 = v7
    v8 := []int{-2147483648, -793, 562, 32768}
    _ = v7
    _ = v8
    return v7
}
func genFn1(p0 uint8) int {
    p0 -= p0
    v9 := GenS2{F0: 474, F1: 1}
    _ = p0
    _ = v9
    return 42
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = -434
    v3 := make(map[string]string)
    var v4 float64 = float64(3.14159)
    switch "a" {
    case "hello":
        _ = len(v1)
    default:
        if (627 > -1) {
            switch genFn1(1) {
            case 2147483647:
                v10 := GenS3{F0: 42}
                _ = v10
            case -2147483648:
                switch "" {
                case "\n\t":
                    if (1.5 <= v4) {
                        v11 := GenS4{F0: ("foo bar" + ("foo bar" + "hello")), F1: (-(v2)), F2: ""}
                        _ = v11
                    } else {
                        if true {
                            v12 := []int{}
                            v13 := GenS5{F0: 610}
                            if (((!(920 <= 65535)) && false) && (3.14159 < (1.5 * 0.1))) {
                                v2 = 379
                                switch "\n\t" {
                                case "hello":
                                    v2 = (v2 * (v2 + (-(v2))))
                                case "a":
                                    var v14 uint32 = 65535
                                    _ = v14
                                case "\n\t":
                                    v2 = (v2 % (((v2 << ((18446744073709551615) & 63))) | 1))
                                default:
                                    if (1.5 >= 0.1) {
                                        for fi15 := 0; fi15 < 17; fi15++ {
                                            for fi16 := 0; fi16 < 8; fi16++ {
                                                v17 := GenS6{F0: -32768}
                                                _ = fi16
                                                _ = v17
                                            }
                                            for fi18 := 0; fi18 < 6; fi18++ {
                                                if (("hello" + "a") > "a") {
                                                    if len(v12) > 0 {
                                                        v12[0] = genFn1(128)
                                                    }
                                                }
                                                v4 = 1.5
                                                var v19 uint32 = 2
                                                _ = fi18
                                                _ = v19
                                            }
                                            _ = fi15
                                        }
                                    } else {
                                        v20 := GenS5{F0: 704}
                                        _ = v20
                                    }
                                }
                                v2 -= (v2 * -847)
                            } else {
                                var v21 int8 = 127
                                v12 = append(v12, genFn1(0))
                                for fi22 := 0; fi22 < 13; fi22++ {
                                    v21 += v21
                                    _ = fi22
                                }
                                _ = v21
                            }
                            _ = v12
                            _ = v13
                        } else {
                            if (false || (false || (false && true))) {
                                v23 := []int{2147483647, 42, 9223372036854775807}
                                fmt.Println(v4, v2)
                                for fi24 := 0; fi24 < 2; fi24++ {
                                    v2 = v2
                                    var v25 bool = true
                                    _ = fi24
                                    _ = v25
                                }
                                _ = v23
                            } else {
                                v4 = ((v4 * (v4 - 1.5)) * v4)
                                for fi26 := 0; fi26 < 9; fi26++ {
                                    v27 := []int{(fi26 & (^(fi26))), (fi26 - genFn1(128))}
                                    if (true && (!(2147483647 < -469))) {
                                        v28 := GenS2{F0: (fi26 ^ 32767), F1: 597}
                                        _ = v28
                                    } else {
                                        if false {
                                            var v29 int16 = -710
                                            _ = v29
                                        }
                                        if true {
                                            var v30 uint16 = 255
                                            _ = v30
                                        }
                                    }
                                    _ = fi26
                                    _ = v27
                                }
                            }
                        }
                        if (436 >= 255) {
                            v2 -= (-(v2))
                        }
                        v4 -= float64((3.14159 - (-2.25 - 0.0)))
                    }
                default:
                    if (2147483647 > 4294967295) {
                        v31 := GenS3{F0: -2147483648}
                        if false {
                            if true {
                                v4 = v4
                                v32 := v3["a"]
                                _ = v32
                            } else {
                                fmt.Println(v2, v2)
                            }
                            for fi33 := 0; fi33 < 1; fi33++ {
                                if false {
                                    v4 += v4
                                }
                                fmt.Println(v4, v4)
                                if ((!false) || (0 != 42)) {
                                    v2 = (v2 ^ v2)
                                    var v34 int = 975
                                    _ = v34
                                } else {
                                    var v35 uint8 = 69
                                    v35 = (v35 % (((v35 & v35)) | 1))
                                    _ = v35
                                }
                                _ = fi33
                            }
                        }
                        v36 := []int{32767}
                        _ = v31
                        _ = v36
                    } else {
                        v2 = (-(v2))
                        if false {
                            for fi37 := 0; fi37 < 18; fi37++ {
                                if true {
                                    if (true || (false && (!(((v4 * (0.0 + 0.0)) - (v4 + v4)) < 0.0)))) {
                                        v38 := GenS6{F0: -786}
                                        _ = v38
                                    }
                                }
                                for fi39 := 0; fi39 < 2; fi39++ {
                                    var v40 string = "\n\t"
                                    _ = fi39
                                    _ = v40
                                }
                                v41 := GenS6{F0: -1}
                                _ = fi37
                                _ = v41
                            }
                            v42 := GenS4{F0: "", F1: (v2 & -364), F2: ("\n\t" + "\n\t")}
                            fmt.Println(v2, v2)
                            _ = v42
                        } else {
                            fmt.Println(v2, v2)
                        }
                    }
                }
            case -304:
                v4 = (float64(0.1) - (v4 - v4))
            }
        }
    }
    v1[2147483648] = ""
    var v43 string = ("a" + "hello")
    v44 := GenS7{F0: 42, F1: (v2 & (^(v2)))}
    var v45 uint = 1
    v45 -= v45
    v46 := GenS7{F0: 42, F1: v2}
    var v47 bool = true
    var v48 float64 = v4
    v43 = (((v43 + (v43 + ((v43 + v43) + v43))) + (v43 + v43)) + v43)
    switch "a" {
    case "":
        if ((^(v2)) == 42) {
            switch "foo bar" {
            case "hello":
                fmt.Println(v43, v47, v45)
            case "\n\t":
                v2 = 1
            case "a":
                if (-53 != -128) {
                    var v49 uint32 = 2147483647
                    if (genFn0() != 2) {
                        v50 := GenS3{F0: 1}
                        _ = v50
                    } else {
                        if v47 {
                            switch 611 {
                            case 127:
                                v51 := GenS8{F0: 964}
                                _ = v51
                            default:
                                delete(v3, v43)
                            }
                            if false {
                                var v52 uint64 = 131
                                _ = v52
                            }
                        } else {
                            switch -92 {
                            case 32767:
                                v53 := GenS6{F0: 0}
                                _ = v53
                            }
                            for fi54 := 0; fi54 < 11; fi54++ {
                                var v55 uint = genFn0()
                                v48 = 0.0
                                fmt.Println(v4, v49, v4)
                                _ = fi54
                                _ = v55
                            }
                            v56 := GenS9{F0: ("\n\t" + ""), F1: 0, F2: 1}
                            _ = v56
                        }
                        if v47 {
                            v57 := v3[(v43 + v43)]
                            v49 = 105
                            _ = v57
                        } else {
                            if ((1e10 * v48) <= (((-2.25 + (0.0 + ((v48 * (v48 * (-2.25 * v48))) - (v48 - v4)))) - float64((1.5 / 1.5))) * 0.1)) {
                                switch 1 {
                                case 4294967295:
                                    fmt.Println(v45)
                                case 737:
                                    if v47 {
                                        v2 = 0
                                        var v58 int16 = 32767
                                        if (255 != 42) {
                                            if (!v47) {
                                                if (-49 >= 100) {
                                                    var v59 string = ("hello" + (v43 + v43))
                                                    for fi60 := 0; fi60 < 6; fi60++ {
                                                        v4 -= 0.0
                                                        var v61 uint8 = 0
                                                        for fi62 := 0; fi62 < 16; fi62++ {
                                                            v59 = "a"
                                                            v63 := []int{(fi60 + fi60)}
                                                            var v64 uint = 4294967295
                                                            fmt.Println(v64, v49, v59)
                                                            _ = fi62
                                                            _ = v63
                                                            _ = v64
                                                        }
                                                        _ = fi60
                                                        _ = v61
                                                    }
                                                    v65 := []int{42}
                                                    _ = v59
                                                    _ = v65
                                                } else {
                                                    fmt.Println(v48)
                                                    for fi66 := 0; fi66 < 8; fi66++ {
                                                        if v47 {
                                                            var v67 uint16 = 0
                                                            var v68 uint32 = (v49 % ((v49) | 1))
                                                            _ = v67
                                                            _ = v68
                                                        } else {
                                                            var v69 uint64 = 4294967295
                                                            _ = v69
                                                        }
                                                        v48 = (v48 - (1.5 / 0.1))
                                                        _ = fi66
                                                    }
                                                    fmt.Println(v43, v43, v48)
                                                }
                                                v48 -= (((v4 * ((v48 / 1.0) + v48)) - 0.1) + (1.5 - (v4 * (float64(1e10) - (((0.0 * 0.1) + (v48 + (v48 + 1e10))) * (((((0.1 * 0.1) + v4) - v48) * v4) * v48))))))
                                                var v70 uint = (v45 * 911)
                                                _ = v70
                                            }
                                            v2 += (-(v2))
                                        } else {
                                            v49 -= v49
                                        }
                                        _ = v58
                                    }
                                default:
                                    v2 += v2
                                }
                                v43 = (("\n\t" + (("" + v43) + "a")) + "foo bar")
                                var v71 bool = v47
                                _ = v71
                            } else {
                                var v72 int16 = -1
                                var v73 int32 = 0
                                _ = v72
                                _ = v73
                            }
                            v48 = float64((float32((0.0 + v4)) / (0.1 / (3.14159 - ((-2.25 / 1.0) * 0.1)))))
                            var v74 uint = (v45 ^ 1)
                            _ = v74
                        }
                    }
                    switch "" {
                    case "a":
                        var v75 int64 = v2
                        _ = v75
                    }
                    _ = v49
                } else {
                    var v76 uint32 = 28
                    v47 = (v47 && (!((!(true || v47)) && false)))
                    if true {
                        var v77 uint64 = 1
                        _ = v77
                    } else {
                        for fi78 := 0; fi78 < 20; fi78++ {
                            for fi79 := 0; fi79 < 10; fi79++ {
                                v80 := GenS10{F0: 42}
                                _ = fi79
                                _ = v80
                            }
                            if true {
                                v47 = (1 <= genFn0())
                                v81 := GenS1{F0: 2147483647, F1: uint64(v45), F2: uint64(v45)}
                                _ = v81
                            } else {
                                v2 = -910
                                fmt.Println(v2, v76, fi78)
                                v82 := GenS1{F0: -690, F1: 0, F2: 404}
                                _ = v82
                            }
                            _ = fi78
                        }
                    }
                    _ = v76
                }
            default:
                if true {
                    fmt.Println(v47, v2)
                    for fi83 := 0; fi83 < 7; fi83++ {
                        var v84 float32 = (1e10 - -2.25)
                        fmt.Println(v2, v4)
                        _ = fi83
                        _ = v84
                    }
                    for fi85 := 0; fi85 < 13; fi85++ {
                        v43 = "hello"
                        _ = fi85
                    }
                }
            }
            switch "hello" {
            case "foo bar":
                for fi86 := 0; fi86 < 7; fi86++ {
                    _ = len(v3)
                    switch 4294967295 {
                    case -995:
                        v87 := []int{1, fi86}
                        _ = v87
                    case 42:
                        if (((-2.25 + 0.1) * (-2.25 / -2.25)) != 0.1) {
                            v48 -= float64(-2.25)
                            v45 += 344
                            for fi88 := 0; fi88 < 10; fi88++ {
                                v89 := v3[v43]
                                _ = fi88
                                _ = v89
                            }
                        } else {
                            v45 += v45
                            for fi90 := 0; fi90 < 10; fi90++ {
                                fmt.Println(v43, v48)
                                var v91 uint8 = 255
                                switch 127 {
                                case -860:
                                    v92 := []int{genFn1(v91)}
                                    _ = v92
                                case 2147483647:
                                    if (!(9223372036854775807 > (-(v2)))) {
                                        if false {
                                            v48 = ((v4 * 1e10) + 1.5)
                                            delete(v1, 691)
                                        }
                                        var v93 bool = ((true || (!v47)) || v47)
                                        if v47 {
                                            v45 += (^(v45))
                                        }
                                        _ = v93
                                    }
                                case 4294967295:
                                    v1[736] = (((v43 + "") + v43) + "hello")
                                default:
                                    switch "" {
                                    case "\n\t":
                                        if (v91 < v91) {
                                            if v47 {
                                                var v94 float32 = (0.1 - (1e10 * 0.0))
                                                v95 := []int{fi86}
                                                _ = v94
                                                _ = v95
                                            } else {
                                                fmt.Println(v91, v2)
                                                var v96 int32 = 1
                                                _ = v96
                                            }
                                        }
                                    }
                                }
                                v47 = ((-1 == 300) && v47)
                                _ = fi90
                                _ = v91
                            }
                            for fi97 := 0; fi97 < 20; fi97++ {
                                if v47 {
                                    v98 := GenS1{F0: -1, F1: 4294967295, F2: 1}
                                    _ = v98
                                } else {
                                    v48 = 3.14159
                                    v47 = v47
                                    v2 += v2
                                    if (((v43 + "\n\t") + v43) >= v43) {
                                        if v47 {
                                            fmt.Println(fi97)
                                            switch (fi86 ^ (fi97 ^ genFn1(0))) {
                                            case -812:
                                                switch (fi97 - -641) {
                                                case -9223372036854775808:
                                                    var v99 int16 = 42
                                                    _ = v99
                                                case -2147483648:
                                                    var v100 uint32 = 2147483648
                                                    _ = v100
                                                }
                                            }
                                        }
                                        if (!(170 != 815)) {
                                            v48 = 3.14159
                                        } else {
                                            if v47 {
                                                if (!v47) {
                                                    v101 := v1[777]
                                                    switch genFn1(2) {
                                                    case -531:
                                                        var v102 string = v43
                                                        _ = v102
                                                    case 127:
                                                        v43 = (((("hello" + v43) + ("hello" + v43)) + ((v101 + "a") + v101)) + v43)
                                                        v101 = ((("" + v43) + "\n\t") + "hello")
                                                        if v47 {
                                                            fmt.Println(v101, fi97)
                                                            var v103 float32 = -2.25
                                                            _ = v103
                                                        } else {
                                                            fmt.Println(v45, v101)
                                                        }
                                                        var v104 uint8 = 0
                                                        _ = v104
                                                    }
                                                    _ = v101
                                                }
                                                v105 := GenS11{F0: v2, F1: (v45 << (((v45 & (v45 << ((v45) & 63)))) & 63))}
                                                _ = v105
                                            } else {
                                                v106 := []int{fi86, fi97}
                                                fmt.Println(fi97, v2)
                                                if v47 {
                                                    switch (-(fi86)) {
                                                    case -64:
                                                        fmt.Println(v4, v47, v48)
                                                        var v107 uint16 = 584
                                                        if v47 {
                                                            if v47 {
                                                                _ = len(v106)
                                                                var v108 float32 = 3.14159
                                                                _ = v108
                                                            } else {
                                                                var v109 int64 = (v2 >> ((18446744073709551615) & 63))
                                                                var v110 int8 = 5
                                                                _ = v109
                                                                _ = v110
                                                            }
                                                            if ((fi86 ^ fi97) == fi86) {
                                                                fmt.Println(v107)
                                                                fmt.Println(v2)
                                                                v111 := GenS4{F0: v43, F1: v2, F2: v43}
                                                                _ = v111
                                                            } else {
                                                                if len(v106) >= 2 {
                                                                    v106 = v106[1:len(v106)]
                                                                }
                                                                var v112 uint8 = 127
                                                                v106 = append(v106, (fi86 >> ((v45) & 63)))
                                                                _ = v112
                                                            }
                                                        } else {
                                                            var v113 string = ""
                                                            v2 = 42
                                                            _ = len(v106)
                                                            _ = v113
                                                        }
                                                        _ = v107
                                                    case -621:
                                                        switch (v43 + v43) {
                                                        case "foo bar":
                                                            if (-2.25 == (0.0 / (float32(v48) * 1e10))) {
                                                                if ((199 < 1) && (21 < 182)) {
                                                                    var v114 uint32 = 785
                                                                    _ = len(v106)
                                                                    v48 = v48
                                                                    _ = v114
                                                                }
                                                                if v47 {
                                                                    fmt.Println(v2, v4, v47)
                                                                    if v47 {
                                                                        var v115 int64 = v2
                                                                        _ = v115
                                                                    }
                                                                    v1[2147483648] = "a"
                                                                } else {
                                                                    switch v43 {
                                                                    case "\n\t":
                                                                        _ = len(v106)
                                                                        var v116 float32 = 1.5
                                                                        v48 = 3.14159
                                                                        if (1 < 1) {
                                                                            var v117 int32 = 42
                                                                            fmt.Println(v2, v47)
                                                                            v118 := GenS12{F0: (fi86 % (((fi86 << ((805) & 63))) | 1)), F1: 255, F2: v47}
                                                                            v116 = 0.0
                                                                            v116 = v116
                                                                            _ = len(v1)
                                                                            _ = v117
                                                                            _ = v118
                                                                        }
                                                                        _ = v116
                                                                    case "foo bar":
                                                                        if (false && (!v47)) {
                                                                            v4 -= v4
                                                                            var v119 bool = (v47 && false)
                                                                            v119 = (genFn1(113) > -20)
                                                                            var v120 float32 = 0.0
                                                                            v2 = (v2 - v2)
                                                                            var v121 string = "a"
                                                                            _ = v119
                                                                            _ = v120
                                                                            _ = v121
                                                                        } else {
                                                                            v45 -= v45
                                                                            switch genFn1(128) {
                                                                            case 32768:
                                                                                if true {
                                                                                    v45 = 255
                                                                                    fmt.Println(fi86, fi97, v2)
                                                                                    v122 := GenS13{F0: v43}
                                                                                    _ = v122
                                                                                } else {
                                                                                    var v123 int8 = -1
                                                                                    _ = v123
                                                                                }
                                                                            default:
                                                                                v47 = v47
                                                                            }
                                                                            var v124 int32 = -1
                                                                            _ = v124
                                                                        }
                                                                    default:
                                                                        delete(v3, "")
                                                                    }
                                                                    var v125 int16 = int16(v45)
                                                                    v48 = ((v4 + ((3.14159 / v48) / v48)) / v48)
                                                                    _ = v125
                                                                }
                                                                var v126 float32 = 1e10
                                                                _ = v126
                                                            }
                                                        default:
                                                            v106 = append(v106, int(v45))
                                                        }
                                                    }
                                                    var v127 int8 = 126
                                                    v48 -= 0.0
                                                    _ = v127
                                                }
                                                _ = v106
                                            }
                                            var v128 uint8 = 8
                                            _ = v128
                                        }
                                    } else {
                                        if (!(1 == 191)) {
                                            var v129 uint8 = 2
                                            var v130 int32 = 0
                                            _ = v129
                                            _ = v130
                                        }
                                        var v131 int = (fi97 ^ fi97)
                                        _ = v131
                                    }
                                    var v132 float32 = 1e10
                                    switch (^(fi97)) {
                                    case -618:
                                        var v133 uint16 = 65535
                                        _ = v133
                                    case 32768:
                                        v134 := GenS13{F0: v43}
                                        _ = v134
                                    }
                                    var v135 uint64 = 4294967295
                                    _ = v132
                                    _ = v135
                                }
                                _ = fi97
                            }
                        }
                    case -191:
                        var v136 uint32 = 42
                        _ = v136
                    }
                    _ = fi86
                }
            case "\n\t":
                v3[(v43 + (v43 + v43))] = (v43 + v43)
            }
            var v137 uint32 = 489
            _ = v137
        }
    case "hello":
        v3["a"] = v43
    case "a":
        v138 := []int{}
        _ = v138
    default:
        for fi139 := 0; fi139 < 17; fi139++ {
            v48 = 1e10
            var v140 bool = v47
            _ = fi139
            _ = v140
        }
    }
    _ = len(v1)
    var v141 float32 = (-2.25 * (3.14159 * float32(float64(1e10))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v141:", v141)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v141
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
}
