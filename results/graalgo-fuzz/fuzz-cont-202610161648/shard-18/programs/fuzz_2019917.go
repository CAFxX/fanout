package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint16 }
type GenS2 struct { F0 uint32; F1 int16; F2 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int; F1 uint16; F2 bool }
type GenS5 struct { F0 bool; F1 uint }
type GenS6 struct { F0 uint64; F1 uint64; F2 string }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint; F1 bool; F2 int64 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 uint64; F1 int64 }
type GenS12 struct { F0 int32; F1 uint16 }
type GenS13 struct { F0 int; F1 int }
func genFn0() int {
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 float64 = 1e10
        var v8 float64 = v7
        v7 = v7
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        v10 := GenS1{F0: 2, F1: "foo bar", F2: 2}
        _ = fi9
        _ = v10
    }
    var v11 int16 = 42
    var v12 uint16 = 32767
    for fi13 := 0; fi13 < 13; fi13++ {
        if false {
            if false {
                switch "hello" {
                case "foo bar":
                    v11 = v11
                case "hello":
                    v12 -= (^(v12))
                case "":
                    switch (fi13 ^ fi13) {
                    case -864:
                        v11 = v11
                    case 0:
                        var v14 float32 = (0.1 + -2.25)
                        _ = v14
                    case 922:
                        fmt.Println(v12, fi13, fi13)
                    }
                }
                if (!true) {
                    var v15 bool = (573 > 0)
                    v11 = int16(fi13)
                    fmt.Println(v11)
                    _ = v15
                } else {
                    for fi16 := 0; fi16 < 6; fi16++ {
                        if (-2147483648 <= -364) {
                            v17 := GenS2{F0: 114, F1: v11, F2: 0}
                            if true {
                                var v18 uint16 = v12
                                switch 32768 {
                                case 1:
                                    if ((1.5 - 0.0) < (1.5 * 0.0)) {
                                        v19 := GenS1{F0: 1, F1: "hello", F2: (^(v12))}
                                        fmt.Println(v11, v18, fi13)
                                        _ = v19
                                    } else {
                                        v20 := GenS3{F0: false}
                                        v12 -= v12
                                        var v21 string = (((("foo bar" + "") + "hello") + "a") + "foo bar")
                                        fmt.Println(v18, v21, v11)
                                        v22 := GenS3{F0: (true || true)}
                                        _ = v20
                                        _ = v21
                                        _ = v22
                                    }
                                case 32768:
                                    var v23 uint = 279
                                    _ = v23
                                case -356:
                                    v12 -= v18
                                }
                                if false {
                                    fmt.Println(v12, fi16, v12)
                                    var v24 uint32 = 2147483647
                                    _ = v24
                                }
                                _ = v18
                            }
                            _ = v17
                        } else {
                            v25 := GenS4{F0: 32768, F1: (v12 % ((v12) | 1)), F2: false}
                            var v26 int32 = 592
                            v12 += (v12 * (v12 * v12))
                            var v27 int64 = 1
                            v28 := []int{519, 236, (fi13 & (fi16 / (((fi13 >> ((894) & 63))) | 1)))}
                            _ = len(v28)
                            _ = v25
                            _ = v26
                            _ = v27
                            _ = v28
                        }
                        v29 := []int{(fi13 * fi13), -53, fi16, fi13}
                        v30 := GenS5{F0: (!((!(0 == 1)) || false)), F1: 1}
                        _ = fi16
                        _ = v29
                        _ = v30
                    }
                    var v31 uint16 = (v12 / (((v12 ^ v12)) | 1))
                    _ = v31
                }
                if false {
                    var v32 float32 = (1.5 - 1e10)
                    var v33 uint8 = uint8(v12)
                    if true {
                        v12 -= v12
                        var v34 float32 = -2.25
                        _ = v34
                    }
                    _ = v32
                    _ = v33
                }
            } else {
                v35 := []int{-377, fi13, fi13, fi13}
                _ = v35
            }
            v11 += (v11 * v11)
        } else {
            switch 188 {
            case 0:
                switch "foo bar" {
                case "\n\t":
                    v36 := GenS2{F0: 2147483647, F1: int16(v12), F2: 103}
                    _ = v36
                }
            case -2147483648:
                var v37 int8 = 0
                _ = v37
            case 128:
                v38 := GenS5{F0: (!false), F1: 255}
                _ = v38
            default:
                var v39 uint = 255
                _ = v39
            }
            v40 := []int{}
            v41 := GenS3{F0: false}
            _ = v40
            _ = v41
        }
        var v42 float64 = 1.5
        v11 = 1000
        _ = fi13
        _ = v42
    }
    _ = v11
    _ = v12
    return 0
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]bool)
    var v3 int16 = 32767
    var v4 int16 = (-(v3))
    var v5 bool = true
    var v43 uint64 = 18446744073709551615
    switch (("a" + "a") + (("foo bar" + ("\n\t" + "hello")) + "\n\t")) {
    case "hello":
        v44 := GenS6{F0: 4294967295, F1: (v43 ^ v43), F2: "a"}
        _ = v44
    case "":
        if v5 {
            if v5 {
                v5 = ((978 > 2) || (821 != genFn0()))
                for fi45 := 0; fi45 < 19; fi45++ {
                    v46 := GenS6{F0: v43, F1: (v43 + (v43 ^ v43)), F2: "\n\t"}
                    if (255 < 98) {
                        if ((0 > 255) || (true && (v5 || (v5 && (!v5))))) {
                            delete(v2, (fi45 >> ((74) & 63)))
                            fmt.Println(v4, v3)
                            v5 = v5
                        }
                        var v47 int32 = -1
                        _ = v47
                    } else {
                        v48 := GenS3{F0: (!v5)}
                        for fi49 := 0; fi49 < 5; fi49++ {
                            delete(v2, fi49)
                            fmt.Println(v3)
                            _ = fi49
                        }
                        v50 := []int{32768, 2147483647, (fi45 - fi45)}
                        _ = v48
                        _ = v50
                    }
                    if ((v43 / ((367) | 1)) >= (^(v43))) {
                        for fi51 := 0; fi51 < 20; fi51++ {
                            var v52 bool = false
                            _ = fi51
                            _ = v52
                        }
                        switch fi45 {
                        case 9223372036854775807:
                            v53 := GenS5{F0: (194 > 214), F1: uint(v43)}
                            _ = v53
                        case 0:
                            var v54 int8 = -128
                            _ = v54
                        default:
                            v55 := []int{-1, (fi45 << ((2) & 63)), 4294967295, (fi45 << ((65535) & 63))}
                            _ = v55
                        }
                        fmt.Println(v43, fi45)
                    } else {
                        v56 := GenS2{F0: 42, F1: 42, F2: 0}
                        var v57 int32 = -86
                        _ = v56
                        _ = v57
                    }
                    _ = fi45
                    _ = v46
                }
            }
        }
    case "\n\t":
        if (-1 >= -222) {
            v58 := []int{}
            _ = v58
        }
    default:
        var v59 bool = (-116 >= 42)
        _ = v59
    }
    var v60 uint8 = 47
    fmt.Println(v60, v5)
    for fi61 := 0; fi61 < 13; fi61++ {
        if (!v5) {
            for fi62 := 0; fi62 < 19; fi62++ {
                switch "foo bar" {
                case "foo bar":
                    var v63 int = (fi62 & fi61)
                    _ = v63
                default:
                    v3 = (v4 >> ((2) & 63))
                }
                v64 := []int{-9223372036854775808}
                _ = fi62
                _ = v64
            }
            var v65 int64 = -9223372036854775808
            v66 := GenS7{F0: 474}
            _ = v65
            _ = v66
        }
        if (((v5 || (false || (!("foo bar" < "hello")))) || (!v5)) && (65535 <= 641)) {
            v43 += (v43 << ((72) & 63))
            v3 -= 42
        } else {
            var v67 string = "\n\t"
            fmt.Println(v43)
            _ = v67
        }
        v4 += v3
        _ = fi61
    }
    switch 4294967295 {
    case 2147483647:
        v68 := GenS5{F0: (int(v43) <= 127), F1: 106}
        _ = v68
    case -23:
        for fi69 := 0; fi69 < 12; fi69++ {
            var v70 uint16 = 802
            var v71 float64 = (-2.25 - (-2.25 * (float64(-2.25) * (1e10 - 0.1))))
            v72 := v1[2147483647]
            _ = fi69
            _ = v70
            _ = v71
            _ = v72
        }
    default:
        v43 += v43
    }
    for fi73 := 0; fi73 < 10; fi73++ {
        if (2147483647 <= 69) {
            fmt.Println(v43, v3)
        } else {
            if v5 {
                var v74 int32 = -1
                _ = v74
            }
        }
        for fi75 := 0; fi75 < 12; fi75++ {
            v5 = (!v5)
            var v76 uint64 = 2
            switch "a" {
            case "\n\t":
                var v77 int16 = v4
                _ = v77
            case "hello":
                if v5 {
                    if ((fi73 ^ (fi75 | -220)) > 0) {
                        switch (fi73 >> ((229) & 63)) {
                        case -9223372036854775808:
                            v3 = v4
                        case 1:
                            fmt.Println(fi73)
                        default:
                            v78 := []int{}
                            fmt.Println(fi73, fi73, v3)
                            var v79 string = "hello"
                            _ = v78
                            _ = v79
                        }
                        v80 := GenS8{F0: 255, F1: false, F2: -9223372036854775808}
                        _ = v80
                    }
                    v81 := GenS3{F0: v5}
                    v82 := []int{(fi73 - fi75), fi73}
                    _ = v81
                    _ = v82
                }
            default:
                v83 := GenS3{F0: v5}
                _ = v83
            }
            _ = fi75
            _ = v76
        }
        _ = fi73
    }
    v84 := []int{int(v60), 128, -18, 842}
    for fi85 := 0; fi85 < 7; fi85++ {
        if true {
            v3 = int16(v43)
            for fi86 := 0; fi86 < 18; fi86++ {
                switch (fi86 << ((18446744073709551615) & 63)) {
                case 32768:
                    v87 := GenS9{F0: v43}
                    _ = v87
                default:
                    var v88 uint = 2
                    _ = v88
                }
                v89 := GenS1{F0: (v43 << ((255) & 63)), F1: "\n\t", F2: 657}
                v90 := GenS8{F0: 319, F1: (!(2147483648 >= 65535)), F2: 42}
                _ = fi86
                _ = v89
                _ = v90
            }
            if (!((-2147483648 >= -1) || ((!((("\n\t" + ((("\n\t" + ("foo bar" + "a")) + "\n\t") + "a")) > ("\n\t" + "")) && true)) || v5))) {
                switch ("a" + ("\n\t" + "\n\t")) {
                case "\n\t":
                    var v91 int16 = 0
                    _ = v91
                case "a":
                    v92 := GenS1{F0: (v43 % ((v43) | 1)), F1: "a", F2: 0}
                    _ = v92
                case "foo bar":
                    v93 := GenS5{F0: (-2147483648 < -2147483648), F1: uint(v43)}
                    _ = v93
                }
                for fi94 := 0; fi94 < 9; fi94++ {
                    fmt.Println(fi94)
                    v4 = (v4 % ((0) | 1))
                    _ = fi94
                }
                v95 := GenS10{F0: 42}
                _ = v95
            } else {
                v96 := GenS9{F0: (v43 * 42)}
                var v97 string = "a"
                fmt.Println(v4, v43, v4)
                _ = v96
                _ = v97
            }
        } else {
            if v5 {
                if (((v5 || (-1 != 492)) || (fi85 == fi85)) || (142 == 1)) {
                    v1[1] = ""
                } else {
                    if v5 {
                        var v98 uint8 = v60
                        _ = v98
                    }
                }
                if (-1 != 42) {
                    for fi99 := 0; fi99 < 8; fi99++ {
                        v60 = v60
                        _ = fi99
                    }
                    var v100 int32 = -1
                    v43 = (v43 - (v43 | (^(v43))))
                    _ = v100
                }
            } else {
                v3 += (^(v4))
            }
            v3 += (v3 & 42)
            v3 = v3
        }
        if v5 {
            v101 := GenS11{F0: (v43 % ((1) | 1)), F1: -9223372036854775808}
            if false {
                delete(v1, 2147483647)
            } else {
                v1[2] = ((("foo bar" + "") + "") + ("\n\t" + "foo bar"))
                if v5 {
                    v43 = (v43 + v43)
                    if v5 {
                        switch "\n\t" {
                        case "foo bar":
                            for fi102 := 0; fi102 < 13; fi102++ {
                                if len(v84) >= 2 {
                                    v84 = v84[1:len(v84)]
                                }
                                v103 := v2[(^(fi85))]
                                _ = fi102
                                _ = v103
                            }
                        case "":
                            v3 = 42
                        case "a":
                            var v104 float64 = 0.0
                            _ = v104
                        }
                    } else {
                        v43 = 2
                    }
                    for fi105 := 0; fi105 < 16; fi105++ {
                        if false {
                            delete(v2, (fi105 - (^(fi105))))
                        }
                        fmt.Println(fi85, v5, v60)
                        _ = fi105
                    }
                }
                delete(v1, 977)
            }
            var v106 bool = (!v5)
            _ = v101
            _ = v106
        } else {
            var v107 uint8 = v60
            _ = v107
        }
        if true {
            v4 = (v3 + v3)
        }
        _ = fi85
    }
    _ = len(v1)
    if len(v84) > 0 {
        v84[0] = 42
    }
    v108 := GenS8{F0: 1, F1: v5, F2: 42}
    if (1 > 56) {
        if len(v84) > 0 {
            v84[0] = 1
        }
        if len(v84) >= 2 {
            v84 = v84[1:len(v84)]
        }
    } else {
        var v109 int16 = v3
        switch 32768 {
        case 208:
            for fi110 := 0; fi110 < 9; fi110++ {
                var v111 int64 = 42
                var v112 float64 = -2.25
                _ = fi110
                _ = v111
                _ = v112
            }
        case 9223372036854775807:
            v113 := GenS8{F0: 564, F1: v5, F2: -9223372036854775808}
            _ = v113
        }
        v114 := GenS12{F0: 2147483647, F1: 65535}
        _ = v109
        _ = v114
    }
    v115 := GenS9{F0: v43}
    v116 := GenS5{F0: true, F1: 2}
    if len(v84) >= 2 {
        v84 = v84[1:len(v84)]
    }
    switch "\n\t" {
    case "foo bar":
        v60 += (v60 - v60)
    case "":
        v3 -= v3
    case "\n\t":
        if (!(36 > 42)) {
            var v117 uint8 = 225
            _ = v117
        }
    default:
        if (!false) {
            for fi118 := 0; fi118 < 3; fi118++ {
                v119 := GenS10{F0: 88}
                if v5 {
                    for fi120 := 0; fi120 < 18; fi120++ {
                        var v121 int16 = (v4 - 32767)
                        v122 := GenS13{F0: -1, F1: fi120}
                        _ = fi120
                        _ = v121
                        _ = v122
                    }
                }
                if (!false) {
                    v123 := GenS5{F0: ((!v5) || (-32 > 59)), F1: 729}
                    _ = v123
                }
                _ = fi118
                _ = v119
            }
            var v124 int32 = 42
            if len(v84) >= 2 {
                v84 = v84[1:len(v84)]
            }
            _ = v124
        } else {
            for fi125 := 0; fi125 < 5; fi125++ {
                if false {
                    var v126 int8 = 100
                    if v5 {
                        if (!(3.14159 < ((-2.25 * (0.1 - 1e10)) / 1.5))) {
                            v3 = v4
                            var v127 uint = uint(v43)
                            _ = v127
                        }
                    }
                    fmt.Println(v3, v5, v3)
                    _ = v126
                } else {
                    v3 += v3
                    var v128 bool = v5
                    v128 = v5
                    _ = v128
                }
                var v129 uint8 = v60
                _ = fi125
                _ = v129
            }
            v60 = (v60 - 40)
            v4 = (v3 % ((v3) | 1))
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v60:", v60)
    fmt.Println("slicelen v84:", len(v84), v84)
    _ = v1
    _ = v108
    _ = v115
    _ = v116
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v60
    _ = v84
}
