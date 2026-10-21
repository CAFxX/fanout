package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 string }
type GenS2 struct { F0 int32; F1 int32; F2 string }
type GenS3 struct { F0 int8; F1 int16; F2 int64 }
type GenS4 struct { F0 string; F1 int16; F2 uint64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int16; F1 bool; F2 int8 }
type GenS7 struct { F0 uint64; F1 int16; F2 uint }
type GenS8 struct { F0 uint; F1 uint64 }
type GenS9 struct { F0 string }

func main() {
    var v1 uint64 = 4294967295
    var v2 uint = 1
    v3 := make(map[int]string)
    var v4 uint16 = 0
    var v5 uint = v2
    var v6 int32 = 0
    v1 -= v1
    switch "" {
    case "hello":
        for fi7 := 0; fi7 < 20; fi7++ {
            v1 -= (v1 - (v1 | v1))
            v8 := GenS1{F0: (^(v6)), F1: (^(v1)), F2: "hello"}
            var v9 uint16 = 42
            _ = fi7
            _ = v8
            _ = v9
        }
    case "":
        v10 := GenS1{F0: -2147483648, F1: (v1 % ((v1) | 1)), F2: ("hello" + "foo bar")}
        _ = v10
    default:
        if (9223372036854775807 > 128) {
            v2 = 4294967295
        } else {
            var v11 uint = v2
            v12 := GenS1{F0: v6, F1: v1, F2: "a"}
            _ = v11
            _ = v12
        }
    }
    v13 := GenS2{F0: v6, F1: -279, F2: "hello"}
    var v14 uint8 = 1
    for fi15 := 0; fi15 < 13; fi15++ {
        switch fi15 {
        case 1:
            v6 -= (v6 + (v6 | v6))
        case 2147483647:
            v16 := GenS1{F0: (v6 + 1), F1: 0, F2: ""}
            _ = v16
        case 128:
            if ((!((65535 <= 2147483647) && true)) || true) {
                for fi17 := 0; fi17 < 16; fi17++ {
                    v5 = 4294967295
                    var v18 int16 = 42
                    _ = fi17
                    _ = v18
                }
                v19 := GenS1{F0: 2147483647, F1: (^(v1)), F2: "a"}
                _ = v19
            } else {
                v20 := []int{}
                fmt.Println(v14)
                _ = v20
            }
        }
        _ = fi15
    }
    for fi21 := 0; fi21 < 6; fi21++ {
        v14 = v14
        for fi22 := 0; fi22 < 19; fi22++ {
            fmt.Println(v1, v2)
            switch fi22 {
            case 128:
                var v23 string = ""
                _ = v23
            case 127:
                v1 = v1
            case -1:
                v24 := GenS2{F0: v6, F1: (v6 % ((v6) | 1)), F2: ""}
                _ = v24
            default:
                fmt.Println(v4, v4)
            }
            v25 := []int{}
            var v26 int16 = 42
            _ = fi22
            _ = v25
            _ = v26
        }
        _ = fi21
    }
    for fi27 := 0; fi27 < 16; fi27++ {
        var v28 int8 = 100
        fmt.Println(v5)
        if false {
            if true {
                v29 := GenS2{F0: v6, F1: (v6 << ((v2) & 63)), F2: ""}
                _ = len(v3)
                fmt.Println(v14)
                _ = v29
            } else {
                v30 := GenS1{F0: (v6 & -1), F1: 0, F2: ("foo bar" + (("foo bar" + "") + ("hello" + "a")))}
                if true {
                    switch (("a" + "a") + "foo bar") {
                    case "a":
                        for fi31 := 0; fi31 < 18; fi31++ {
                            v32 := GenS1{F0: (v6 & v6), F1: v1, F2: "a"}
                            _ = fi31
                            _ = v32
                        }
                    case "\n\t":
                        for fi33 := 0; fi33 < 1; fi33++ {
                            v34 := GenS1{F0: -1, F1: v1, F2: "foo bar"}
                            if true {
                                v35 := []int{fi27, 9223372036854775807}
                                _ = v35
                            } else {
                                v36 := GenS3{F0: 1, F1: 32767, F2: -1}
                                if (v6 == v6) {
                                    v37 := []int{int(v14), 674, (fi33 << ((v2) & 63)), fi33}
                                    _ = v37
                                }
                                if ((fi33 > (fi33 % ((int(v4)) | 1))) && true) {
                                    var v38 int64 = 1
                                    _ = v38
                                } else {
                                    v5 += 4294967295
                                    var v39 uint64 = (v1 % ((4294967295) | 1))
                                    _ = v39
                                }
                                fmt.Println(v28, v28, v14)
                                v40 := GenS4{F0: "\n\t", F1: 1, F2: v1}
                                var v41 int16 = 294
                                var v42 int16 = v41
                                var v43 int32 = (v6 << (((v2 ^ (v2 * (v2 * (v2 + v5))))) & 63))
                                _ = v36
                                _ = v40
                                _ = v41
                                _ = v42
                                _ = v43
                            }
                            if false {
                                if true {
                                    switch ("\n\t" + "a") {
                                    case "a":
                                        var v44 int = int(v6)
                                        _ = v44
                                    case "\n\t":
                                        if true {
                                            v45 := GenS5{F0: 582}
                                            v5 += v5
                                            _ = v45
                                        } else {
                                            v46 := GenS4{F0: "\n\t", F1: 0, F2: 18446744073709551615}
                                            _ = v46
                                        }
                                    default:
                                        v47 := []int{fi27, (fi33 / (((-(fi27))) | 1)), int(v4)}
                                        _ = v47
                                    }
                                    fmt.Println(v2)
                                } else {
                                    v5 += 4294967295
                                    fmt.Println(v2, fi33, v4)
                                    v48 := []int{42, fi33, fi27, (fi27 >> ((v2) & 63))}
                                    _ = v48
                                }
                            } else {
                                switch -758 {
                                case 2147483647:
                                    v49 := GenS2{F0: v6, F1: (-(v6)), F2: ("" + "hello")}
                                    v50 := GenS3{F0: v28, F1: -1, F2: -9223372036854775808}
                                    v5 = 482
                                    _ = v49
                                    _ = v50
                                case 0:
                                    if true {
                                        if true {
                                            if false {
                                                if true {
                                                    switch ((("hello" + "\n\t") + "") + "\n\t") {
                                                    case "":
                                                        v6 += (v6 << (((v5 >> ((2) & 63))) & 63))
                                                    case "\n\t":
                                                        fmt.Println(v4)
                                                    default:
                                                        delete(v3, fi33)
                                                    }
                                                    fmt.Println(v5, fi27, fi33)
                                                    v28 = (v28 / ((v28) | 1))
                                                    switch "\n\t" {
                                                    case "":
                                                        delete(v3, 32768)
                                                    }
                                                }
                                            } else {
                                                _ = len(v3)
                                                if false {
                                                    v51 := GenS6{F0: -32768, F1: ((true || (42 < 990)) && true), F2: (v28 & (v28 >> ((v2) & 63)))}
                                                    v4 += (v4 | 0)
                                                    _ = v51
                                                }
                                                switch (fi27 * (-(fi33))) {
                                                case 0:
                                                    if false {
                                                        var v52 uint64 = uint64(v4)
                                                        _ = v52
                                                    } else {
                                                        var v53 uint = 524
                                                        fmt.Println(v1, v6, v14)
                                                        _ = v53
                                                    }
                                                default:
                                                    fmt.Println(v6, fi27, fi27)
                                                }
                                                v3[(fi27 << ((2) & 63))] = (("foo bar" + ("\n\t" + "foo bar")) + (("foo bar" + "a") + ""))
                                                var v54 int32 = 1
                                                _ = v54
                                            }
                                            switch ("foo bar" + "") {
                                            case "\n\t":
                                                v55 := []int{}
                                                _ = v55
                                            case "foo bar":
                                                v56 := GenS7{F0: v1, F1: 42, F2: v5}
                                                _ = v56
                                            case "a":
                                                if false {
                                                    var v57 string = ("\n\t" + "a")
                                                    var v58 float32 = float32(3.14159)
                                                    _ = v57
                                                    _ = v58
                                                } else {
                                                    v5 += 2
                                                    fmt.Println(v28, v1)
                                                }
                                            default:
                                                v6 -= 2147483647
                                            }
                                            if true {
                                                var v59 bool = (1000 == -768)
                                                v60 := GenS8{F0: v5, F1: 18446744073709551615}
                                                var v61 uint16 = v4
                                                v62 := GenS4{F0: (("a" + "") + "a"), F1: -32768, F2: (^(v1))}
                                                v1 = 94
                                                v61 += v4
                                                _ = v59
                                                _ = v60
                                                _ = v61
                                                _ = v62
                                            } else {
                                                fmt.Println(v6)
                                                if (true && true) {
                                                    v6 -= (v6 / (((v6 / (((v6 % (((v6 * int32(v1))) | 1))) | 1))) | 1))
                                                    v4 = (v4 >> (((v5 / ((1) | 1))) & 63))
                                                    fmt.Println(fi27)
                                                    if true {
                                                        v28 = 119
                                                        var v63 bool = true
                                                        switch "" {
                                                        case "a":
                                                            v1 = v1
                                                        default:
                                                            v64 := []int{(fi27 << ((0) & 63)), fi27}
                                                            _ = v64
                                                        }
                                                        _ = v63
                                                    } else {
                                                        var v65 uint8 = v14
                                                        v66 := []int{fi27}
                                                        fmt.Println(v2, fi33)
                                                        _ = v65
                                                        _ = v66
                                                    }
                                                    var v67 int16 = 1000
                                                    if ((v1 % ((42) | 1)) > v1) {
                                                        fmt.Println(v1, v14, v2)
                                                        var v68 float32 = 0.0
                                                        _ = v68
                                                    }
                                                    _ = v67
                                                }
                                                var v69 uint16 = 358
                                                _ = v69
                                            }
                                        } else {
                                            var v70 int16 = 1000
                                            var v71 bool = true
                                            _ = v70
                                            _ = v71
                                        }
                                        v72 := []int{fi27, (fi27 + fi27)}
                                        if len(v72) >= 2 {
                                            v72 = v72[1:len(v72)]
                                        }
                                        _ = v72
                                    } else {
                                        v73 := []int{(-(fi33)), 4294967295, (fi33 % (((^(fi27))) | 1))}
                                        fmt.Println(v4, v4)
                                        _ = v73
                                    }
                                case -1:
                                    var v74 int8 = -128
                                    _ = v74
                                }
                                switch "" {
                                case "hello":
                                    switch (("hello" + ("" + "foo bar")) + "hello") {
                                    case "hello":
                                        var v75 uint = 2
                                        _ = v75
                                    }
                                case "\n\t":
                                    _ = len(v3)
                                case "":
                                    if (((v6 / ((v6) | 1)) != (v6 ^ (^(v6)))) && false) {
                                        v76 := GenS8{F0: 2, F1: v1}
                                        var v77 int = 32767
                                        _ = v76
                                        _ = v77
                                    }
                                }
                                v78 := GenS4{F0: "foo bar", F1: 32767, F2: (^(v1))}
                                v14 = v14
                                v1 = 18446744073709551615
                                _ = v78
                            }
                            switch "foo bar" {
                            case "\n\t":
                                v79 := GenS2{F0: -2147483648, F1: 42, F2: "\n\t"}
                                _ = v79
                            default:
                                fmt.Println(fi33, fi27, v5)
                            }
                            v5 -= v5
                            _ = fi33
                            _ = v34
                        }
                    }
                } else {
                    v1 -= (^(v1))
                    for fi80 := 0; fi80 < 17; fi80++ {
                        v14 -= v14
                        _ = fi80
                    }
                }
                v3[(fi27 >> ((v2) & 63))] = "hello"
                _ = v30
            }
            var v81 float64 = 3.14159
            _ = v81
        } else {
            for fi82 := 0; fi82 < 9; fi82++ {
                v3[(fi82 & (-(fi82)))] = ("foo bar" + ("a" + ("\n\t" + "a")))
                v83 := []int{(fi82 & -2147483648), (fi27 - fi27)}
                fmt.Println(v28, v4, v6)
                _ = fi82
                _ = v83
            }
            var v84 int32 = (v6 - (v6 & v6))
            if false {
                var v85 string = "a"
                for fi86 := 0; fi86 < 11; fi86++ {
                    var v87 bool = false
                    if v87 {
                        fmt.Println(fi86, v6)
                        if (fi86 < 4294967295) {
                            var v88 uint8 = (v14 / ((42) | 1))
                            if v87 {
                                v89 := GenS9{F0: "hello"}
                                v90 := []int{32767, (^(fi86)), 128}
                                var v91 int8 = v28
                                _ = v89
                                _ = v90
                                _ = v91
                            }
                            _ = v88
                        } else {
                            switch v85 {
                            case "hello":
                                var v92 int = (^(fi27))
                                _ = v92
                            case "":
                                var v93 int8 = (^(v28))
                                _ = v93
                            case "foo bar":
                                v6 = v6
                            }
                            switch v85 {
                            case "":
                                if (!(v1 == (^(v1)))) {
                                    if (v87 || v87) {
                                        if false {
                                            if v87 {
                                                var v94 int64 = 156
                                                var v95 int8 = (v28 & v28)
                                                var v96 string = "\n\t"
                                                _ = v94
                                                _ = v95
                                                _ = v96
                                            } else {
                                                v97 := GenS5{F0: -1}
                                                fmt.Println(v85)
                                                _ = v97
                                            }
                                            v98 := GenS4{F0: ("a" + ""), F1: 42, F2: v1}
                                            var v99 float32 = 1.5
                                            _ = v98
                                            _ = v99
                                        } else {
                                            v100 := []int{}
                                            v101 := GenS5{F0: (^(fi27))}
                                            switch fi86 {
                                            case 0:
                                                fmt.Println(v28, v28)
                                            case 128:
                                                var v102 uint8 = v14
                                                switch (("foo bar" + ((("\n\t" + ("\n\t" + "\n\t")) + "hello") + "hello")) + v85) {
                                                case "":
                                                    v103 := GenS4{F0: v85, F1: 32767, F2: 1}
                                                    _ = v103
                                                case "\n\t":
                                                    var v104 float32 = ((0.1 - -2.25) + float32(0.0))
                                                    _ = v104
                                                case "hello":
                                                    v6 = 2147483647
                                                default:
                                                    var v105 string = v85
                                                    _ = v105
                                                }
                                                _ = v102
                                            default:
                                                switch (fi86 & (fi86 - (fi27 ^ int(v1)))) {
                                                case -418:
                                                    switch 128 {
                                                    case 0:
                                                        fmt.Println(v28)
                                                    }
                                                case 128:
                                                    var v106 uint32 = 2
                                                    _ = v106
                                                }
                                            }
                                            _ = v100
                                            _ = v101
                                        }
                                        v107 := []int{(fi27 % (((^(fi27))) | 1)), fi27, (fi86 - (fi86 << ((v2) & 63))), fi27}
                                        _ = v107
                                    } else {
                                        if true {
                                            v14 -= 0
                                            var v108 int = fi27
                                            var v109 uint64 = v1
                                            v84 = v6
                                            _ = v108
                                            _ = v109
                                        } else {
                                            delete(v3, (fi86 & (fi86 | fi86)))
                                            var v110 int8 = 1
                                            v111 := []int{int(v110), fi86, int(v84)}
                                            fmt.Println(v84, v6, v1)
                                            _ = v110
                                            _ = v111
                                        }
                                    }
                                    v84 -= (v84 - (v84 & int32(fi27)))
                                }
                            case "foo bar":
                                var v112 bool = ((3.14159 - (1.5 * 1.5)) == 1.5)
                                _ = v112
                            case "\n\t":
                                if (((-1 != (-(v84))) && v87) && (((v14 != (^(v14))) && false) || v87)) {
                                    fmt.Println(v5)
                                    var v113 string = "a"
                                    _ = v113
                                }
                            }
                            switch (fi27 >> ((v2) & 63)) {
                            case -889:
                                if (false && v87) {
                                    var v114 uint16 = (^(v4))
                                    var v115 int32 = (^(v6))
                                    var v116 int = (fi27 * 1)
                                    _ = v114
                                    _ = v115
                                    _ = v116
                                } else {
                                    fmt.Println(v85, v85, v1)
                                }
                            default:
                                var v117 uint = 834
                                _ = v117
                            }
                        }
                    }
                    var v118 int = (fi27 << ((v5) & 63))
                    var v119 uint16 = (^(v4))
                    var v120 int8 = 100
                    _ = fi86
                    _ = v118
                    _ = v119
                    _ = v120
                    _ = v87
                }
                switch (fi27 ^ fi27) {
                case -782:
                    if true {
                        for fi121 := 0; fi121 < 9; fi121++ {
                            var v122 int64 = -9223372036854775808
                            var v123 int32 = v6
                            v124 := GenS3{F0: 42, F1: 1000, F2: (v122 ^ (^(v122)))}
                            _ = fi121
                            _ = v122
                            _ = v123
                            _ = v124
                        }
                        var v125 int8 = 103
                        for fi126 := 0; fi126 < 7; fi126++ {
                            var v127 int = fi27
                            v6 += (v6 ^ int32(v125))
                            _ = fi126
                            _ = v127
                        }
                        _ = v125
                    }
                case -26:
                    var v128 uint = 42
                    _ = v128
                }
                _ = v85
            }
            _ = v84
        }
        _ = fi27
        _ = v28
    }
    for fi129 := 0; fi129 < 2; fi129++ {
        var v130 int32 = (v6 / ((int32(fi129)) | 1))
        if false {
            v131 := []int{-1, fi129, fi129, fi129}
            for fi132 := 0; fi132 < 8; fi132++ {
                v4 -= 255
                if ((v4 % ((v4) | 1)) > v4) {
                    var v133 int16 = 294
                    var v134 int16 = (v133 ^ (v133 + (^(v133))))
                    _ = v133
                    _ = v134
                }
                if true {
                    fmt.Println(v6, v130)
                }
                _ = fi132
            }
            _ = v131
        }
        _ = fi129
        _ = v130
    }
    fmt.Println(v1, v1, v4)
    var v135 uint = v2
    fmt.Println("v1:", v1)
    fmt.Println("v135:", v135)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v135
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
