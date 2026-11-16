package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 bool; F1 int16; F2 uint32 }
type GenS3 struct { F0 uint64; F1 uint; F2 uint }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint64; F1 int32 }
type GenS6 struct { F0 bool; F1 int64; F2 string }
type GenS7 struct { F0 int64; F1 bool }
type GenS8 struct { F0 bool; F1 int16 }
type GenS9 struct { F0 int16; F1 uint32 }
type GenS10 struct { F0 uint32; F1 int8 }
type GenS11 struct { F0 uint; F1 int32; F2 int8 }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 int32; F1 uint16 }
type GenS14 struct { F0 uint; F1 string; F2 uint }
type GenS15 struct { F0 bool; F1 int16 }
type GenS16 struct { F0 int64; F1 uint64 }
type GenS17 struct { F0 int8; F1 uint; F2 int32 }
func genFn0() int8 {
    var v5 int8 = -94
    if true {
        v6 := GenS1{F0: (("foo bar" + "foo bar") + "\n\t"), F1: 32767}
        var v7 uint8 = 255
        _ = v6
        _ = v7
    } else {
        var v8 bool = (115 == 128)
        if (v8 && true) {
            if v8 {
                v9 := []int{-2147483648, 32767}
                _ = v9
            } else {
                if (-38 < v5) {
                    for fi10 := 0; fi10 < 16; fi10++ {
                        fmt.Println(fi10, v5)
                        v11 := GenS2{F0: (592 >= 1), F1: 42, F2: 2147483648}
                        fmt.Println(v8, v5)
                        _ = fi10
                        _ = v11
                    }
                    v12 := []int{0, 0, 1}
                    _ = v12
                } else {
                    if true {
                        if (903 == 218) {
                            var v13 uint8 = 255
                            _ = v13
                        }
                    }
                    if (32767 > -32768) {
                        v14 := GenS1{F0: "foo bar", F1: 1000}
                        _ = v14
                    } else {
                        v15 := []int{9223372036854775807, -976}
                        for fi16 := 0; fi16 < 20; fi16++ {
                            var v17 float64 = 0.0
                            var v18 uint32 = 774
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                        _ = v15
                    }
                    if ((2147483647 < 42) || false) {
                        v19 := []int{2147483647, 9223372036854775807}
                        var v20 int8 = (v5 + v5)
                        v8 = true
                        _ = v19
                        _ = v20
                    } else {
                        v21 := GenS3{F0: 4294967295, F1: 18446744073709551615, F2: 459}
                        switch "\n\t" {
                        case "hello":
                            switch "foo bar" {
                            case "foo bar":
                                v8 = v8
                            default:
                                var v22 string = ""
                                _ = v22
                            }
                        }
                        _ = v21
                    }
                }
                for fi23 := 0; fi23 < 19; fi23++ {
                    v8 = (18446744073709551615 <= 18446744073709551615)
                    v5 = (^(v5))
                    switch int(v5) {
                    case 431:
                        for fi24 := 0; fi24 < 3; fi24++ {
                            fmt.Println(v8)
                            _ = fi24
                        }
                    case 1:
                        for fi25 := 0; fi25 < 8; fi25++ {
                            fmt.Println(v5)
                            _ = fi25
                        }
                    case -9223372036854775808:
                        var v26 bool = true
                        _ = v26
                    }
                    _ = fi23
                }
                v27 := GenS4{F0: -201}
                _ = v27
            }
            var v28 int8 = 1
            _ = v28
        }
        v5 = v5
        _ = v8
    }
    v5 -= (v5 * -28)
    _ = v5
    return (v5 % (((v5 << ((0) & 63))) | 1))
}
func genFn1(p0 int, p1 uint8, p2 int32) int64 {
    var v29 int32 = (-(p2))
    var v30 float32 = 3.14159
    switch "" {
    case "hello":
        var v31 int32 = p2
        _ = v31
    case "":
        var v32 float64 = -2.25
        _ = v32
    case "foo bar":
        v29 += p2
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v29
    _ = v30
    return 1
}

func main() {
    var v1 uint64 = 42
    var v2 int32 = -458
    var v3 float32 = 0.1
    var v4 float32 = (v3 + 1e10)
    var v33 int16 = 0
    if true {
        var v34 uint32 = 589
        var v35 string = ("\n\t" + "hello")
        v36 := GenS2{F0: true, F1: -32768, F2: (v34 ^ (v34 << ((250) & 63)))}
        _ = v34
        _ = v35
        _ = v36
    } else {
        v37 := []int{-966, -1}
        _ = v37
    }
    v33 = v33
    v38 := GenS1{F0: "", F1: (v33 / ((v33) | 1))}
    fmt.Println(v4, v2, v2)
    v3 = (((v3 / float32(((-2.25 + 3.14159) + 1.5))) + ((v3 - v3) + 0.1)) - 1.5)
    v39 := []int{9223372036854775807, 9223372036854775807, 127, 915}
    if len(v39) > 0 {
        v39[0] = 2147483647
    }
    v2 -= v2
    v3 = float32(-2.25)
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v1)
    case "hello":
        var v40 string = ("\n\t" + "hello")
        _ = v40
    case "foo bar":
        var v41 int32 = (v2 ^ (v2 * (v2 << ((896) & 63))))
        _ = v41
    default:
        var v42 uint = 18446744073709551615
        _ = v42
    }
    for fi43 := 0; fi43 < 7; fi43++ {
        fmt.Println(v2, v2)
        v44 := GenS5{F0: (v1 ^ v1), F1: (v2 & -2147483648)}
        v39 = append(v39, (fi43 & fi43))
        _ = fi43
        _ = v44
    }
    switch "" {
    case "a":
        if (!false) {
            v45 := GenS1{F0: "a", F1: (^(v33))}
            fmt.Println(v1)
            _ = v45
        } else {
            v46 := GenS5{F0: v1, F1: v2}
            switch -2147483648 {
            case 1:
                v4 -= -2.25
            default:
                v47 := GenS6{F0: (true && true), F1: 1, F2: "hello"}
                _ = v47
            }
            v4 = v3
            _ = v46
        }
    case "":
        if true {
            switch -531 {
            case -81:
                v3 = v4
            case 127:
                var v48 int8 = genFn0()
                _ = v48
            case -9223372036854775808:
                if len(v39) >= 2 {
                    v39 = v39[1:len(v39)]
                }
            }
            v3 = v4
            if ((3.14159 / 1e10) < -2.25) {
                for fi49 := 0; fi49 < 3; fi49++ {
                    switch "\n\t" {
                    case "":
                        if false {
                            for fi50 := 0; fi50 < 10; fi50++ {
                                var v51 int32 = (v2 & 766)
                                _ = fi50
                                _ = v51
                            }
                            if false {
                                v3 = v3
                                v52 := GenS7{F0: -9223372036854775808, F1: true}
                                if (!true) {
                                    var v53 int = fi49
                                    _ = v53
                                }
                                _ = v52
                            }
                            var v54 int8 = 1
                            _ = v54
                        }
                    case "\n\t":
                        v33 = v33
                    case "a":
                        for fi55 := 0; fi55 < 2; fi55++ {
                            _ = len(v39)
                            switch ("a" + "foo bar") {
                            case "foo bar":
                                if true {
                                    fmt.Println(fi55)
                                    v33 = (^(v33))
                                } else {
                                    if ((true && (false && (!true))) || (0 <= 116)) {
                                        v56 := GenS6{F0: (42 >= (v1 & v1)), F1: -1, F2: "a"}
                                        v57 := GenS4{F0: 42}
                                        _ = v56
                                        _ = v57
                                    }
                                    if true {
                                        var v58 uint16 = 42
                                        var v59 int16 = (v33 & v33)
                                        _ = v58
                                        _ = v59
                                    } else {
                                        fmt.Println(v3, v33, v2)
                                        v4 += 1e10
                                        fmt.Println(v3, v2, fi55)
                                    }
                                    switch fi49 {
                                    case 4294967295:
                                        var v60 uint16 = 32767
                                        fmt.Println(v3, v60)
                                        _ = v60
                                    case 0:
                                        v61 := GenS3{F0: (^(v1)), F1: 2, F2: 1}
                                        _ = v61
                                    default:
                                        v39 = append(v39, fi49)
                                    }
                                }
                            case "":
                                v62 := GenS8{F0: ((v2 / (((^(v2))) | 1)) < (v2 % (((v2 & (v2 << ((42) & 63)))) | 1))), F1: (v33 << ((65535) & 63))}
                                _ = v62
                            }
                            _ = fi55
                        }
                    }
                    if false {
                        fmt.Println(v2, v3)
                        v4 += (v3 / v3)
                    } else {
                        var v63 int64 = -1
                        if false {
                            v63 += (-(v63))
                            v4 += (v4 - v4)
                        } else {
                            switch "" {
                            case "\n\t":
                                for fi64 := 0; fi64 < 4; fi64++ {
                                    v3 = float32(1e10)
                                    var v65 float64 = (1.5 / 0.1)
                                    _ = fi64
                                    _ = v65
                                }
                            }
                            if false {
                                v33 = (v33 << ((1) & 63))
                                if len(v39) > 0 {
                                    v39[0] = (-(fi49))
                                }
                            } else {
                                if len(v39) >= 2 {
                                    v39 = v39[1:len(v39)]
                                }
                                var v66 uint = 508
                                v2 = -709
                                _ = v66
                            }
                            if (0 <= 255) {
                                if true {
                                    v63 = v63
                                    var v67 float32 = 3.14159
                                    _ = v67
                                } else {
                                    var v68 string = ("" + "\n\t")
                                    _ = v68
                                }
                                if (2 != 127) {
                                    v33 = v33
                                    v63 -= genFn1(fi49, 255, v2)
                                } else {
                                    fmt.Println(fi49, v1, fi49)
                                    for fi69 := 0; fi69 < 8; fi69++ {
                                        fmt.Println(fi49)
                                        switch "hello" {
                                        case "":
                                            if len(v39) >= 2 {
                                                v39 = v39[1:len(v39)]
                                            }
                                            if false {
                                                if (1e10 == float64(v4)) {
                                                    var v70 int = int(v33)
                                                    v39 = append(v39, (fi49 ^ fi69))
                                                    var v71 int64 = v63
                                                    _ = v70
                                                    _ = v71
                                                } else {
                                                    v39 = append(v39, fi69)
                                                    _ = len(v39)
                                                    var v72 uint64 = (v1 >> ((42) & 63))
                                                    fmt.Println(v1, v1)
                                                    var v73 uint8 = 255
                                                    _ = v72
                                                    _ = v73
                                                }
                                                v74 := GenS9{F0: v33, F1: 396}
                                                var v75 float64 = float64(v3)
                                                _ = v74
                                                _ = v75
                                            } else {
                                                fmt.Println(v4, v33, fi49)
                                                if true {
                                                    var v76 int32 = 0
                                                    fmt.Println(v4, v33)
                                                    var v77 uint = 8
                                                    _ = v76
                                                    _ = v77
                                                }
                                                var v78 int = fi69
                                                _ = v78
                                            }
                                        }
                                        v79 := GenS10{F0: 4294967295, F1: int8(v2)}
                                        _ = fi69
                                        _ = v79
                                    }
                                }
                                for fi80 := 0; fi80 < 20; fi80++ {
                                    v39 = append(v39, fi49)
                                    _ = fi80
                                }
                            } else {
                                fmt.Println(v63, v33, fi49)
                                switch fi49 {
                                case -9223372036854775808:
                                    v81 := GenS10{F0: 2147483648, F1: -54}
                                    _ = v81
                                case -698:
                                    switch fi49 {
                                    case -9223372036854775808:
                                        _ = len(v39)
                                    case 1:
                                        v3 += (v3 - float32(float64(v3)))
                                    }
                                default:
                                    switch "hello" {
                                    case "a":
                                        var v82 bool = true
                                        _ = v82
                                    }
                                }
                                switch "foo bar" {
                                case "foo bar":
                                    v83 := GenS11{F0: 327, F1: (v2 & v2), F2: genFn0()}
                                    _ = v83
                                case "hello":
                                    for fi84 := 0; fi84 < 19; fi84++ {
                                        v85 := GenS4{F0: -32768}
                                        _ = fi84
                                        _ = v85
                                    }
                                }
                            }
                        }
                        _ = v63
                    }
                    _ = fi49
                }
                if true {
                    var v86 float32 = v3
                    _ = v86
                } else {
                    v87 := GenS12{F0: (v33 >> ((65535) & 63))}
                    _ = v87
                }
            }
        } else {
            v88 := GenS12{F0: v33}
            _ = v88
        }
    case "hello":
        switch 683 {
        case 1:
            if false {
                if true {
                    if len(v39) > 0 {
                        v39[0] = -130
                    }
                    if (3.14159 > 3.14159) {
                        if false {
                            switch ("a" + ("a" + "a")) {
                            case "hello":
                                v89 := GenS13{F0: v2, F1: 2}
                                _ = v89
                            default:
                                v1 = v1
                            }
                        }
                        if false {
                            switch 32767 {
                            case -1:
                                for fi90 := 0; fi90 < 3; fi90++ {
                                    if (v2 >= v2) {
                                        v4 += (3.14159 * v3)
                                        for fi91 := 0; fi91 < 18; fi91++ {
                                            _ = len(v39)
                                            if true {
                                                v3 -= v3
                                                v3 += (v4 / (1.5 / 1.0))
                                                if (!true) {
                                                    v3 += v3
                                                }
                                                var v92 int16 = v33
                                                _ = v92
                                            } else {
                                                switch fi90 {
                                                case 42:
                                                    fmt.Println(fi90)
                                                case -9223372036854775808:
                                                    if len(v39) > 0 {
                                                        v39[0] = 9223372036854775807
                                                    }
                                                case 32767:
                                                    var v93 uint = 42
                                                    _ = v93
                                                }
                                                v4 = ((v4 - 1e10) + ((0.1 / v3) - -2.25))
                                                v33 = (v33 | (v33 & (v33 >> ((4294967295) & 63))))
                                            }
                                            if ((("a" + "a") + "a") <= "foo bar") {
                                                switch -1 {
                                                case 85:
                                                    if ((v4 < float32((-2.25 / (((0.1 / 0.1) / -2.25) - (0.1 * (float64((-2.25 + v3)) + (((1.5 + 0.0) / 3.14159) + ((-2.25 + 3.14159) + (-2.25 + (0.1 / 0.1)))))))))) && false) {
                                                        v33 = v33
                                                    }
                                                case 32767:
                                                    var v94 uint64 = (v1 ^ 1)
                                                    _ = v94
                                                default:
                                                    if false {
                                                        if ((!false) && false) {
                                                            var v95 string = "hello"
                                                            switch v95 {
                                                            case "foo bar":
                                                                if len(v39) >= 2 {
                                                                    v39 = v39[1:len(v39)]
                                                                }
                                                            case "":
                                                                switch ((v95 + v95) + "") {
                                                                case "hello":
                                                                    var v96 uint32 = 1
                                                                    _ = v96
                                                                case "\n\t":
                                                                    var v97 int16 = (v33 - (-(v33)))
                                                                    _ = v97
                                                                default:
                                                                    var v98 float32 = float32((0.1 - 0.0))
                                                                    _ = v98
                                                                }
                                                            default:
                                                                switch fi90 {
                                                                case -9223372036854775808:
                                                                    v99 := GenS2{F0: false, F1: (-(v33)), F2: 631}
                                                                    _ = v99
                                                                }
                                                            }
                                                            fmt.Println(fi90)
                                                            _ = v95
                                                        }
                                                    }
                                                }
                                                var v100 string = "a"
                                                v101 := GenS13{F0: v2, F1: 65535}
                                                _ = v100
                                                _ = v101
                                            }
                                            _ = fi91
                                        }
                                    } else {
                                        for fi102 := 0; fi102 < 10; fi102++ {
                                            var v103 uint16 = 1
                                            fmt.Println(v4)
                                            fmt.Println(v103)
                                            if true {
                                                var v104 bool = (!((!(fi90 <= fi102)) || ((((2 >= 2147483647) || (-9223372036854775808 == 9223372036854775807)) && (v1 == (v1 * 42))) && false)))
                                                v1 = (^(v1))
                                                v103 = 65535
                                                if false {
                                                    var v105 int64 = -168
                                                    var v106 int16 = v33
                                                    _ = v105
                                                    _ = v106
                                                } else {
                                                    fmt.Println(v2, v2)
                                                    switch fi90 {
                                                    case -759:
                                                        var v107 uint32 = 2147483648
                                                        _ = v107
                                                    case 32767:
                                                        v108 := GenS14{F0: 42, F1: "foo bar", F2: 2}
                                                        _ = v108
                                                    case -744:
                                                        if (100 < 1) {
                                                            fmt.Println(v103, v33, v1)
                                                            switch ("a" + "\n\t") {
                                                            case "foo bar":
                                                                var v109 uint = 648
                                                                _ = v109
                                                            }
                                                            _ = len(v39)
                                                        }
                                                    }
                                                }
                                                _ = v104
                                            } else {
                                                v103 = v103
                                            }
                                            v110 := GenS10{F0: 2, F1: 1}
                                            v39 = append(v39, (-(fi102)))
                                            _ = fi102
                                            _ = v103
                                            _ = v110
                                        }
                                    }
                                    var v111 uint8 = 2
                                    v112 := GenS15{F0: true, F1: v33}
                                    _ = fi90
                                    _ = v111
                                    _ = v112
                                }
                            case 9223372036854775807:
                                v3 -= v3
                            }
                            for fi113 := 0; fi113 < 5; fi113++ {
                                if true {
                                    for fi114 := 0; fi114 < 1; fi114++ {
                                        v33 = v33
                                        v1 = (v1 * v1)
                                        if (!true) {
                                            v1 = (v1 % ((42) | 1))
                                            fmt.Println(v3)
                                        } else {
                                            var v115 int64 = 9223372036854775807
                                            var v116 uint64 = 4294967295
                                            switch fi113 {
                                            case -814:
                                                switch ("\n\t" + "foo bar") {
                                                case "":
                                                    v3 = (((v4 - 0.1) / (v4 - v4)) / (0.1 - 0.0))
                                                case "foo bar":
                                                    v2 = v2
                                                case "\n\t":
                                                    v2 = (v2 * 0)
                                                default:
                                                    var v117 int32 = 1
                                                    _ = v117
                                                }
                                            default:
                                                v39 = append(v39, (fi114 | (fi114 ^ 42)))
                                                v118 := GenS16{F0: v115, F1: (v1 / ((v1) | 1))}
                                                var v119 uint8 = 255
                                                _ = v118
                                                _ = v119
                                            }
                                            _ = v115
                                            _ = v116
                                        }
                                        _ = fi114
                                    }
                                    v120 := GenS13{F0: -99, F1: 70}
                                    _ = v120
                                } else {
                                    var v121 int8 = 91
                                    v122 := GenS17{F0: v121, F1: 65535, F2: v2}
                                    _ = v121
                                    _ = v122
                                }
                                for fi123 := 0; fi123 < 13; fi123++ {
                                    fmt.Println(v4, v1, fi123)
                                    v124 := GenS16{F0: 9223372036854775807, F1: v1}
                                    _ = fi123
                                    _ = v124
                                }
                                var v125 uint8 = 42
                                _ = fi113
                                _ = v125
                            }
                            fmt.Println(v4, v4, v2)
                        }
                    }
                    for fi126 := 0; fi126 < 16; fi126++ {
                        v1 = 18446744073709551615
                        fmt.Println(v3, v4)
                        if false {
                            var v127 bool = false
                            v2 += 42
                            _ = v127
                        }
                        _ = fi126
                    }
                }
                switch "hello" {
                case "":
                    v2 = v2
                case "foo bar":
                    fmt.Println(v4, v3, v3)
                default:
                    var v128 int8 = 127
                    _ = v128
                }
                fmt.Println(v4, v4, v1)
            }
        }
    default:
        v1 -= (^(v1))
    }
    v1 = v1
    var v129 float32 = (3.14159 / v4)
    for fi130 := 0; fi130 < 7; fi130++ {
        v39 = append(v39, (fi130 % (((fi130 & -1)) | 1)))
        v2 += (-(v2))
        if true {
            v2 = (-(v2))
            fmt.Println(v33, fi130, fi130)
        } else {
            fmt.Println(v2)
            v131 := GenS17{F0: 100, F1: 834, F2: v2}
            _ = v131
        }
        _ = fi130
    }
    v33 = -32768
    var v132 uint16 = 665
    fmt.Println("v1:", v1)
    fmt.Println("v129:", v129)
    fmt.Println("v132:", v132)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v129
    _ = v132
    _ = v2
    _ = v3
    _ = v33
    _ = v38
    _ = v39
    _ = v4
}
