package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 uint32 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 uint16; F1 int16; F2 int32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 bool; F1 uint32 }
type GenS6 struct { F0 int32; F1 bool }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int64; F1 string }
func genFn0(p0 uint8) int8 {
    var v6 uint8 = p0
    v7 := GenS1{F0: p0, F1: -128, F2: 174}
    if true {
        var v8 uint8 = p0
        _ = v8
    }
    var v9 int = 0
    _ = p0
    _ = v6
    _ = v7
    _ = v9
    return 1
}
func genFn1(p0 int32) int32 {
    switch ("a" + "hello") {
    case "":
        v10 := []int{32768, 32767, -690, -1}
        _ = v10
    case "a":
        for fi11 := 0; fi11 < 8; fi11++ {
            v12 := GenS1{F0: 128, F1: 42, F2: 2}
            if (42 == 691) {
                p0 -= 1
                v13 := GenS2{F0: "foo bar", F1: -1}
                _ = v13
            } else {
                for fi14 := 0; fi14 < 18; fi14++ {
                    switch ((("\n\t" + "foo bar") + ("" + ("\n\t" + ""))) + "foo bar") {
                    case "\n\t":
                        v15 := GenS1{F0: 108, F1: 127, F2: 4294967295}
                        _ = v15
                    }
                    if (1 < genFn0(172)) {
                        if (false || true) {
                            switch fi11 {
                            case 32767:
                                v16 := GenS3{F0: 65535, F1: 42, F2: 2147483647}
                                _ = v16
                            case 42:
                                var v17 int16 = 32767
                                _ = v17
                            default:
                                p0 = (p0 / ((375) | 1))
                            }
                            p0 -= p0
                        } else {
                            p0 += int32(fi11)
                            v18 := GenS2{F0: "hello", F1: (^(p0))}
                            var v19 int = fi11
                            _ = v18
                            _ = v19
                        }
                        if true {
                            v20 := GenS4{F0: 255}
                            p0 -= 1
                            _ = v20
                        }
                        if true {
                            switch (fi11 >> ((349) & 63)) {
                            case 42:
                                var v21 uint8 = 42
                                _ = v21
                            case -903:
                                var v22 int32 = -136
                                _ = v22
                            default:
                                if (!true) {
                                    if (!false) {
                                        if (true || true) {
                                            if ("\n\t" >= (("" + "") + "foo bar")) {
                                                if (0 <= 42) {
                                                    if ((!(!(true && true))) && (652 > 2147483647)) {
                                                        p0 += -2147483648
                                                        if (1 <= -648) {
                                                            if (0 <= 18446744073709551615) {
                                                                v23 := []int{(fi14 / ((fi14) | 1)), fi11, (fi11 | (fi14 + (fi11 - (fi11 & fi11)))), 32768}
                                                                if (1 == 0) {
                                                                    v24 := GenS1{F0: 152, F1: -128, F2: 2147483647}
                                                                    _ = v24
                                                                }
                                                                var v25 float64 = (0.1 / 3.14159)
                                                                switch "hello" {
                                                                case "hello":
                                                                    fmt.Println(v25)
                                                                case "\n\t":
                                                                    v26 := GenS3{F0: 65535, F1: -32768, F2: 2147483647}
                                                                    _ = v26
                                                                }
                                                                if (!false) {
                                                                    if (fi11 < 128) {
                                                                        v25 = v25
                                                                        _ = len(v23)
                                                                        switch (fi14 >> ((2) & 63)) {
                                                                        case -9223372036854775808:
                                                                            switch "a" {
                                                                            case "":
                                                                                p0 -= 696
                                                                            }
                                                                        default:
                                                                            switch fi14 {
                                                                            case 32768:
                                                                                p0 = (^(p0))
                                                                                switch fi11 {
                                                                                case 42:
                                                                                    fmt.Println(fi14, v25)
                                                                                case 128:
                                                                                    _ = len(v23)
                                                                                case 546:
                                                                                    var v27 uint32 = 4294967295
                                                                                    _ = v27
                                                                                default:
                                                                                    fmt.Println(fi11, v25, v25)
                                                                                }
                                                                            case 1:
                                                                                if false {
                                                                                    v25 = (v25 - ((0.0 * ((float64((0.0 / 1.0)) / v25) * v25)) / v25))
                                                                                    p0 = (^(p0))
                                                                                }
                                                                            }
                                                                        }
                                                                    } else {
                                                                        var v28 float64 = (0.1 + 1.5)
                                                                        var v29 uint = 65535
                                                                        _ = v28
                                                                        _ = v29
                                                                    }
                                                                    if (false || false) {
                                                                        var v30 int = (fi14 / ((-9223372036854775808) | 1))
                                                                        var v31 uint64 = 42
                                                                        if false {
                                                                            var v32 int8 = genFn0(127)
                                                                            _ = v32
                                                                        } else {
                                                                            fmt.Println(fi14, p0)
                                                                            var v33 bool = true
                                                                            _ = v33
                                                                        }
                                                                        _ = v30
                                                                        _ = v31
                                                                    }
                                                                    p0 = (p0 - -453)
                                                                } else {
                                                                    var v34 bool = (false || false)
                                                                    var v35 uint = 255
                                                                    _ = v34
                                                                    _ = v35
                                                                }
                                                                _ = v23
                                                                _ = v25
                                                            } else {
                                                                switch "" {
                                                                case "foo bar":
                                                                    var v36 bool = (864 > 213)
                                                                    _ = v36
                                                                case "a":
                                                                    var v37 uint8 = 148
                                                                    var v38 float64 = (1.5 * float64(1.5))
                                                                    _ = v37
                                                                    _ = v38
                                                                case "\n\t":
                                                                    switch 128 {
                                                                    case 4294967295:
                                                                        v39 := []int{(fi14 - (fi14 - fi14)), (^(fi14)), (^(fi11)), fi14}
                                                                        _ = v39
                                                                    case 127:
                                                                        p0 += (p0 % (((-(p0))) | 1))
                                                                    case 32767:
                                                                        var v40 bool = (!false)
                                                                        _ = v40
                                                                    }
                                                                default:
                                                                    var v41 int32 = (p0 >> ((1) & 63))
                                                                    _ = v41
                                                                }
                                                                p0 -= 1
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    p0 = (p0 | (p0 / (((p0 & -358)) | 1)))
                                                }
                                                var v42 int16 = 42
                                                _ = v42
                                            } else {
                                                var v43 uint = 18446744073709551615
                                                v44 := GenS1{F0: 255, F1: 100, F2: 2147483648}
                                                if (false || true) {
                                                    v45 := []int{fi11, 42, fi11}
                                                    switch ("hello" + "hello") {
                                                    case "hello":
                                                        fmt.Println(fi14, v43)
                                                    default:
                                                        v46 := GenS2{F0: ("" + "foo bar"), F1: p0}
                                                        _ = v46
                                                    }
                                                    _ = v45
                                                }
                                                v47 := GenS3{F0: 66, F1: 32767, F2: p0}
                                                _ = v43
                                                _ = v44
                                                _ = v47
                                            }
                                            p0 += (-(p0))
                                        } else {
                                            switch ("\n\t" + "hello") {
                                            case "foo bar":
                                                fmt.Println(fi11, fi14, p0)
                                            case "a":
                                                fmt.Println(p0, fi11, fi11)
                                            }
                                            var v48 float32 = 0.1
                                            var v49 uint = 198
                                            _ = v48
                                            _ = v49
                                        }
                                    } else {
                                        if (1e10 < (1.5 * -2.25)) {
                                            var v50 int32 = -1
                                            _ = v50
                                        } else {
                                            if true {
                                                var v51 uint8 = 127
                                                var v52 string = (("\n\t" + ("" + "")) + "foo bar")
                                                var v53 uint64 = 4294967295
                                                v54 := GenS2{F0: "a", F1: p0}
                                                _ = v51
                                                _ = v52
                                                _ = v53
                                                _ = v54
                                            }
                                        }
                                        var v55 uint16 = 65535
                                        _ = v55
                                    }
                                    fmt.Println(fi14, fi14)
                                }
                            }
                            fmt.Println(fi14)
                        } else {
                            fmt.Println(fi14, p0)
                        }
                    } else {
                        var v56 int64 = -536
                        var v57 uint16 = 42
                        v58 := GenS5{F0: (!false), F1: 628}
                        _ = v56
                        _ = v57
                        _ = v58
                    }
                    _ = fi14
                }
                var v59 float32 = 3.14159
                v60 := []int{fi11, 11, -9223372036854775808}
                _ = v59
                _ = v60
            }
            p0 = p0
            _ = fi11
            _ = v12
        }
    default:
        v61 := []int{-814}
        _ = v61
    }
    var v62 uint32 = 2147483647
    var v63 uint16 = 65535
    fmt.Println(v63, p0)
    p0 -= -196
    _ = p0
    _ = v62
    _ = v63
    return p0
}

func main() {
    v1 := []int{42, 128, -138}
    var v2 float32 = 1e10
    v3 := []int{}
    var v4 string = "hello"
    var v5 uint64 = 0
    for fi64 := 0; fi64 < 3; fi64++ {
        switch v4 {
        case "foo bar":
            v4 = v4
        case "hello":
            if ((18446744073709551615 >= 255) && (!false)) {
                var v65 int32 = -2147483648
                for fi66 := 0; fi66 < 20; fi66++ {
                    var v67 uint8 = 166
                    var v68 int8 = 127
                    _ = fi66
                    _ = v67
                    _ = v68
                }
                for fi69 := 0; fi69 < 4; fi69++ {
                    v2 += 0.0
                    _ = fi69
                }
                _ = v65
            } else {
                v1 = append(v1, fi64)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        }
        var v70 uint64 = v5
        _ = fi64
        _ = v70
    }
    switch 466 {
    case 176:
        if true {
            var v71 uint16 = 410
            fmt.Println(v71)
            _ = v71
        }
    default:
        _ = len(v1)
    }
    v2 -= 3.14159
    if len(v1) > 0 {
        v1[0] = 0
    }
    for fi72 := 0; fi72 < 2; fi72++ {
        if (128 <= 128) {
            var v73 bool = true
            if (!v73) {
                var v74 float32 = ((v2 / v2) + (v2 - (v2 + ((v2 / v2) / (v2 + v2)))))
                fmt.Println(v2, v73, v5)
                for fi75 := 0; fi75 < 3; fi75++ {
                    var v76 float32 = -2.25
                    _ = fi75
                    _ = v76
                }
                _ = v74
            }
            _ = v73
        } else {
            var v77 int16 = 1000
            var v78 uint64 = (v5 >> ((4294967295) & 63))
            _ = v77
            _ = v78
        }
        _ = fi72
    }
    v2 = -2.25
    switch -1 {
    case -283:
        v4 = v4
    case -1:
        for fi79 := 0; fi79 < 7; fi79++ {
            if len(v3) > 0 {
                v3[0] = fi79
            }
            _ = fi79
        }
    case 361:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    switch -678 {
    case -899:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case 4294967295:
        v80 := GenS6{F0: 42, F1: (-1 > -1)}
        _ = v80
    default:
        switch -9223372036854775808 {
        case 32768:
            v81 := GenS7{F0: v4}
            _ = v81
        case 42:
            if false {
                for fi82 := 0; fi82 < 1; fi82++ {
                    if (int32(fi82) > -2147483648) {
                        for fi83 := 0; fi83 < 16; fi83++ {
                            if len(v3) > 0 {
                                v3[0] = fi82
                            }
                            _ = fi83
                        }
                    }
                    var v84 uint8 = 127
                    for fi85 := 0; fi85 < 19; fi85++ {
                        var v86 uint = 255
                        var v87 string = "hello"
                        _ = len(v3)
                        _ = fi85
                        _ = v86
                        _ = v87
                    }
                    _ = fi82
                    _ = v84
                }
                switch -1 {
                case -660:
                    fmt.Println(v5)
                case -2147483648:
                    v88 := GenS3{F0: 1, F1: -32768, F2: -906}
                    _ = v88
                default:
                    v3 = append(v3, -292)
                }
            } else {
                v4 = v4
            }
        case -1:
            var v89 bool = false
            _ = v89
        default:
            var v90 uint16 = 0
            _ = v90
        }
    }
    v1 = append(v1, 194)
    if true {
        var v91 uint32 = 2
        if true {
            if (198 >= 1) {
                v91 = v91
            } else {
                var v92 uint = 70
                _ = v92
            }
            fmt.Println(v2, v5)
            if len(v3) > 0 {
                v3[0] = 2147483647
            }
        }
        if false {
            switch ("hello" + ("" + v4)) {
            case "hello":
                fmt.Println(v91, v91, v5)
            case "a":
                v3 = append(v3, -563)
            default:
                var v93 uint8 = 2
                _ = v93
            }
            if ((!true) || (!true)) {
                for fi94 := 0; fi94 < 9; fi94++ {
                    v2 = v2
                    v91 -= (v91 >> ((644) & 63))
                    if false {
                        for fi95 := 0; fi95 < 19; fi95++ {
                            if len(v1) > 0 {
                                v1[0] = 286
                            }
                            if false {
                                fmt.Println(v2, v5, v4)
                                var v96 int16 = 42
                                _ = v96
                            } else {
                                var v97 float32 = 3.14159
                                if ((!(v97 <= v97)) && false) {
                                    if false {
                                        if false {
                                            v4 = v4
                                            v91 = (v91 ^ uint32(v5))
                                        }
                                        if (2 == 42) {
                                            v97 = v2
                                            var v98 uint8 = 255
                                            _ = len(v3)
                                            _ = v98
                                        }
                                    } else {
                                        v5 -= 329
                                    }
                                    v99 := GenS8{F0: 1, F1: ((v4 + v4) + ("\n\t" + v4))}
                                    var v100 uint16 = 32767
                                    v101 := GenS5{F0: (!true), F1: (^(v91))}
                                    v102 := GenS3{F0: v100, F1: 0, F2: 2147483647}
                                    _ = v100
                                    _ = v101
                                    _ = v102
                                    _ = v99
                                }
                                v103 := GenS3{F0: 921, F1: 0, F2: -2147483648}
                                _ = v103
                                _ = v97
                            }
                            var v104 int64 = -1
                            var v105 int64 = (v104 >> ((255) & 63))
                            _ = fi95
                            _ = v104
                            _ = v105
                        }
                    } else {
                        if len(v1) > 0 {
                            v1[0] = fi94
                        }
                    }
                    _ = fi94
                }
            } else {
                var v106 int64 = 0
                v107 := GenS5{F0: false, F1: 4294967295}
                v91 = (v91 % (((v91 % (((v91 % ((196) | 1))) | 1))) | 1))
                _ = v106
                _ = v107
            }
            for fi108 := 0; fi108 < 5; fi108++ {
                if false {
                    var v109 uint64 = 149
                    _ = v109
                } else {
                    v91 += 0
                }
                _ = fi108
            }
        }
        _ = v91
    } else {
        v4 = v4
        v2 = v2
    }
    var v110 bool = false
    var v111 uint = 1
    var v112 uint16 = 518
    switch v4 {
    case "":
        var v113 int8 = -27
        _ = v113
    case "foo bar":
        v3 = append(v3, 32767)
    default:
        var v114 uint = (^(v111))
        _ = v114
    }
    var v115 int32 = 1
    for fi116 := 0; fi116 < 20; fi116++ {
        v117 := GenS5{F0: (32767 != 42), F1: 1}
        if v110 {
            v5 = v5
        }
        var v118 bool = v110
        _ = fi116
        _ = v117
        _ = v118
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v110:", v110)
    fmt.Println("v111:", v111)
    fmt.Println("v112:", v112)
    fmt.Println("v115:", v115)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v110
    _ = v111
    _ = v112
    _ = v115
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
