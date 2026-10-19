package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 uint8; F1 uint; F2 int64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 string; F1 uint16 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int; F1 bool }
type GenS10 struct { F0 int16; F1 uint16 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 int64; F1 uint32 }
func genFn0(p0 uint8, p1 int8, p2 int16) int16 {
    fmt.Println(p2, p2)
    var v7 string = "\n\t"
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 uint32 = 4294967295
        for fi10 := 0; fi10 < 8; fi10++ {
            var v11 int32 = 42
            var v12 float32 = float32(float64(1.5))
            if true {
                fmt.Println(v9)
            } else {
                v13 := GenS1{F0: 42}
                var v14 int64 = 1
                var v15 float64 = (1e10 / 1.0)
                _ = v13
                _ = v14
                _ = v15
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        v16 := []int{(fi8 >> ((42) & 63)), (fi8 >> ((2) & 63)), fi8}
        _ = fi8
        _ = v16
        _ = v9
    }
    fmt.Println(v7)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p2
}

func main() {
    var v1 string = "a"
    v2 := []uint32{1}
    v3 := []uint32{7, 2147483648, 0}
    var v4 string = (v1 + ("" + (v1 + ((v1 + v1) + v1))))
    v5 := make(map[string]bool)
    v6 := make(map[uint32]int)
    delete(v5, ("" + "hello"))
    for fi17 := 0; fi17 < 7; fi17++ {
        var v18 int64 = 9223372036854775807
        for fi19 := 0; fi19 < 3; fi19++ {
            if false {
                v18 += (v18 | (v18 & 42))
            } else {
                if false {
                    switch 1 {
                    case 128:
                        if (!false) {
                            v20 := GenS1{F0: -2147483648}
                            v18 -= v18
                            _ = v20
                        } else {
                            var v21 float32 = (1e10 + float32(float64(((-2.25 / 1.0) / -2.25))))
                            _ = v21
                        }
                    case 234:
                        v1 = ("\n\t" + ((v4 + (v1 + "")) + v1))
                        if (1 == -714) {
                            var v22 uint = 4294967295
                            _ = v22
                        }
                        v18 = v18
                    case -9223372036854775808:
                        if true {
                            v1 = "foo bar"
                            fmt.Println(v4, fi17)
                            v3 = append(v3, 4294967295)
                        } else {
                            if ((1 == 100) && true) {
                                switch "a" {
                                case "":
                                    var v23 uint64 = 44
                                    _ = v23
                                case "\n\t":
                                    if ((v18 >> ((4294967295) & 63)) <= 42) {
                                        v5[("" + "hello")] = false
                                        var v24 uint32 = 0
                                        v18 += v18
                                        _ = v24
                                    } else {
                                        v25 := GenS1{F0: -2147483648}
                                        if (32767 == 339) {
                                            var v26 int16 = -135
                                            v27 := GenS2{F0: 130, F1: fi17}
                                            _ = v26
                                            _ = v27
                                        } else {
                                            v1 = (("foo bar" + v1) + "")
                                            v1 = (v1 + v4)
                                            switch fi17 {
                                            case 2147483647:
                                                v4 = v1
                                            case -9223372036854775808:
                                                v28 := GenS3{F0: 42, F1: 65535, F2: v18}
                                                _ = v28
                                            default:
                                                v5[(v1 + v1)] = ((!true) && (((1.5 + (0.0 / (((1.5 - 1e10) * 0.1) * (1.5 + 0.1)))) >= -2.25) && (true || (false || (!(!true))))))
                                            }
                                        }
                                        fmt.Println(v4, v1, fi17)
                                        if ("a" <= "hello") {
                                            v29 := GenS4{F0: 531}
                                            _ = v29
                                        }
                                        v30 := GenS5{F0: 766, F1: 30, F2: 2}
                                        _ = v25
                                        _ = v30
                                    }
                                }
                                fmt.Println(fi19)
                            }
                        }
                    default:
                        v4 = (("hello" + "") + (v4 + v4))
                    }
                } else {
                    var v31 int = fi19
                    _ = v31
                }
                if true {
                    v2 = append(v2, 2147483647)
                    var v32 int8 = -128
                    if (!(32767 >= 32767)) {
                        var v33 int8 = 0
                        v34 := GenS4{F0: 2}
                        _ = v33
                        _ = v34
                    }
                    v4 = v4
                    v35 := GenS6{F0: 42}
                    _ = v32
                    _ = v35
                }
                v36 := GenS4{F0: 353}
                if ((!(!(!true))) && (!true)) {
                    fmt.Println(v4, fi19)
                    v1 = "foo bar"
                    if ((float32(float64((1.5 * 0.1))) + 3.14159) != 0.0) {
                        v18 = v18
                    }
                }
                _ = v36
            }
            var v37 int8 = 80
            _ = fi19
            _ = v37
        }
        _ = fi17
        _ = v18
    }
    for fi38 := 0; fi38 < 19; fi38++ {
        var v39 int16 = -482
        _ = len(v3)
        if false {
            var v40 int8 = 127
            v4 = v1
            _ = v40
        }
        _ = fi38
        _ = v39
    }
    for fi41 := 0; fi41 < 2; fi41++ {
        v4 = ("" + "a")
        _ = fi41
    }
    v4 = v1
    var v42 int32 = -1
    if (776 <= 0) {
        v4 = v1
    } else {
        switch "hello" {
        case "foo bar":
            var v43 float32 = (3.14159 + float32((float64((0.1 - -2.25)) - 1.5)))
            _ = v43
        case "\n\t":
            fmt.Println(v4, v42, v42)
        default:
            v44 := GenS2{F0: 0, F1: 32767}
            _ = v44
        }
        switch 100 {
        case 9223372036854775807:
            for fi45 := 0; fi45 < 11; fi45++ {
                for fi46 := 0; fi46 < 15; fi46++ {
                    if (9223372036854775807 > 42) {
                        var v47 int8 = 0
                        _ = v47
                    } else {
                        v4 = (v1 + "")
                        v48 := GenS6{F0: -1}
                        var v49 uint8 = 2
                        switch (("hello" + (v1 + ("foo bar" + v1))) + "a") {
                        case "hello":
                            v2 = append(v2, 4294967295)
                        case "a":
                            if ((true || false) && false) {
                                if (fi46 > fi45) {
                                    var v50 int = (fi45 & (fi46 - fi46))
                                    var v51 int = v50
                                    _ = v50
                                    _ = v51
                                }
                                if len(v2) > 0 {
                                    v2[0] = 1
                                }
                                v4 = "hello"
                            }
                        default:
                            v49 -= (^(v49))
                        }
                        if ((409 != 462) && (!true)) {
                            var v52 bool = true
                            v53 := GenS4{F0: 654}
                            switch v4 {
                            case "foo bar":
                                var v54 uint = 42
                                _ = v54
                            case "\n\t":
                                v55 := GenS7{F0: ("" + (v1 + v4)), F1: 501}
                                _ = v55
                            case "":
                                var v56 uint32 = 383
                                _ = v56
                            default:
                                fmt.Println(v42, v42, fi45)
                            }
                            _ = v52
                            _ = v53
                        }
                        _ = v48
                        _ = v49
                    }
                    fmt.Println(v4, fi46)
                    v3 = append(v3, 65535)
                    _ = fi46
                }
                if false {
                    var v57 int16 = 744
                    if false {
                        var v58 string = "\n\t"
                        _ = v58
                    } else {
                        for fi59 := 0; fi59 < 13; fi59++ {
                            v57 -= 32767
                            fmt.Println(v57, v57)
                            _ = fi59
                        }
                        if ((953 >= 255) && false) {
                            v4 = ((v4 + (v4 + (v4 + v4))) + (((v4 + v4) + (((v4 + v1) + (v1 + (v4 + ("a" + v4)))) + "a")) + ""))
                            for fi60 := 0; fi60 < 12; fi60++ {
                                delete(v6, 2)
                                v61 := GenS4{F0: 993}
                                v1 = ((v1 + v1) + ("foo bar" + v1))
                                _ = fi60
                                _ = v61
                            }
                            v57 += v57
                        }
                    }
                    _ = v57
                }
                _ = fi45
            }
        case 128:
            switch -2147483648 {
            case -747:
                v42 = 0
            case -9223372036854775808:
                var v62 float64 = 0.0
                _ = v62
            default:
                var v63 int64 = -1
                _ = v63
            }
        default:
            var v64 int32 = (v42 % ((v42) | 1))
            _ = v64
        }
        var v65 int16 = 32767
        _ = v65
    }
    switch -2147483648 {
    case 0:
        v4 = "\n\t"
    case 32768:
        for fi66 := 0; fi66 < 16; fi66++ {
            v67 := GenS8{F0: ((fi66 * (-(fi66))) == fi66)}
            _ = fi66
            _ = v67
        }
    case 128:
        var v68 int = 2147483647
        _ = v68
    default:
        if ((-2.25 - 0.1) == (((1e10 / (0.0 - 3.14159)) / -2.25) / ((1e10 * 3.14159) - 1.5))) {
            var v69 int = 2147483647
            _ = v69
        } else {
            if (v4 == v4) {
                var v70 int64 = 9223372036854775807
                _ = v70
            } else {
                v4 = ("\n\t" + v4)
                if (290 != 388) {
                    v1 = "hello"
                    var v71 float64 = 0.1
                    for fi72 := 0; fi72 < 20; fi72++ {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        if (!true) {
                            fmt.Println(v1, v71, v42)
                            switch (fi72 % ((-1) | 1)) {
                            case -2147483648:
                                for fi73 := 0; fi73 < 17; fi73++ {
                                    fmt.Println(v71, v71)
                                    v4 = v4
                                    _ = fi73
                                }
                            default:
                                if true {
                                    v1 = ""
                                } else {
                                    if ((true || false) || false) {
                                        fmt.Println(v42, v1)
                                        switch (v1 + "\n\t") {
                                        case "\n\t":
                                            v42 += (v42 - v42)
                                        case "foo bar":
                                            v74 := GenS5{F0: -32768, F1: 255, F2: 255}
                                            _ = v74
                                        }
                                    } else {
                                        for fi75 := 0; fi75 < 11; fi75++ {
                                            var v76 float32 = float32((0.1 / v71))
                                            fmt.Println(fi72, v42)
                                            v2 = append(v2, 2)
                                            _ = fi75
                                            _ = v76
                                        }
                                    }
                                    for fi77 := 0; fi77 < 6; fi77++ {
                                        v78 := GenS9{F0: fi77, F1: (163 > 163)}
                                        _ = fi77
                                        _ = v78
                                    }
                                    switch fi72 {
                                    case -384:
                                        if ((((0.0 * 1.5) / 1.5) / 1.5) != float32(float64(0.1))) {
                                            switch 4294967295 {
                                            case 2147483647:
                                                for fi79 := 0; fi79 < 18; fi79++ {
                                                    var v80 float32 = -2.25
                                                    _ = fi79
                                                    _ = v80
                                                }
                                            default:
                                                var v81 int32 = (v42 - (v42 | int32(fi72)))
                                                _ = v81
                                            }
                                            if false {
                                                for fi82 := 0; fi82 < 20; fi82++ {
                                                    if false {
                                                        fmt.Println(v4, fi82, fi72)
                                                        var v83 int = fi82
                                                        v83 += (v83 | fi72)
                                                        _ = v83
                                                    }
                                                    if (442 == 2) {
                                                        fmt.Println(v42, v4)
                                                        if len(v3) > 0 {
                                                            v3[0] = 301
                                                        }
                                                        var v84 int16 = genFn0(128, -1, 42)
                                                        _ = v84
                                                    } else {
                                                        v4 = "hello"
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        if (-9223372036854775808 >= -9223372036854775808) {
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            switch (fi82 - (fi72 << ((133) & 63))) {
                                                            case -9223372036854775808:
                                                                v1 = v4
                                                            }
                                                        } else {
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            var v85 uint = 635
                                                            v4 = ("a" + ((v1 + "\n\t") + v1))
                                                            var v86 uint16 = 221
                                                            _ = v85
                                                            _ = v86
                                                        }
                                                        if (!(true || (0 <= -1))) {
                                                            var v87 uint = 1
                                                            _ = v87
                                                        }
                                                        var v88 float32 = ((1.5 + (0.0 * 0.0)) - ((((0.0 / 1.0) + (-2.25 * (((1e10 * -2.25) / float32((v71 - v71))) * (3.14159 / 1.5)))) * ((3.14159 - 1.5) * float32((float64((1.5 - (((0.0 - 1e10) / (-2.25 / 0.1)) * 0.0))) / 1e10)))) / (float32(-2.25) * float32(0.1))))
                                                        v1 = v1
                                                        var v89 uint16 = 293
                                                        var v90 int8 = 127
                                                        fmt.Println(fi82)
                                                        fmt.Println(v89, v90, v88)
                                                        v91 := GenS10{F0: 32767, F1: (v89 & (^(v89)))}
                                                        var v92 uint8 = 232
                                                        _ = v88
                                                        _ = v89
                                                        _ = v90
                                                        _ = v91
                                                        _ = v92
                                                    }
                                                    _ = fi82
                                                }
                                                v2 = append(v2, 4294967295)
                                            }
                                        }
                                    }
                                }
                            }
                            v93 := v5[v1]
                            _ = v93
                        }
                        if (fi72 != (fi72 / ((fi72) | 1))) {
                            if (9223372036854775807 < 9223372036854775807) {
                                if (true && (((v71 + float64((1e10 * float32(v71)))) < v71) && (((!((fi72 | fi72) >= (fi72 & fi72))) && false) || true))) {
                                    var v94 int32 = (v42 * (v42 - (-(v42))))
                                    _ = v94
                                }
                            }
                            if (!false) {
                                var v95 int64 = 0
                                if (1 == 100) {
                                    if (65535 != 65535) {
                                        if (1 <= fi72) {
                                            v96 := GenS11{F0: 128}
                                            fmt.Println(v71, v1, v71)
                                            _ = v96
                                        } else {
                                            v4 = v4
                                        }
                                    }
                                } else {
                                    v97 := GenS4{F0: 794}
                                    var v98 int64 = (v95 >> ((819) & 63))
                                    var v99 int64 = -9223372036854775808
                                    _ = v97
                                    _ = v98
                                    _ = v99
                                }
                                _ = v95
                            } else {
                                if true {
                                    var v100 int16 = -32768
                                    fmt.Println(v100, v100)
                                    _ = v100
                                } else {
                                    if true {
                                        v4 = v1
                                        var v101 uint16 = 539
                                        _ = v101
                                    } else {
                                        fmt.Println(fi72, v1, v1)
                                        v102 := GenS10{F0: -1, F1: 255}
                                        v103 := GenS12{F0: 1, F1: 208}
                                        _ = v102
                                        _ = v103
                                    }
                                }
                                if (!(!true)) {
                                    v104 := GenS12{F0: 731, F1: 65535}
                                    _ = len(v6)
                                    var v105 int8 = 89
                                    _ = v104
                                    _ = v105
                                } else {
                                    if (243 < 863) {
                                        if (!true) {
                                            v42 = (^(v42))
                                            _ = len(v6)
                                        }
                                        fmt.Println(v71, v4)
                                    }
                                    var v106 int16 = 532
                                    _ = v106
                                }
                            }
                            delete(v5, (v1 + v1))
                        } else {
                            v107 := v6[2147483647]
                            switch v107 {
                            case 2147483647:
                                if (203 == 206) {
                                    var v108 int32 = v42
                                    _ = v108
                                }
                            case 277:
                                v107 = (v107 + (v107 << ((699) & 63)))
                            case 478:
                                for fi109 := 0; fi109 < 17; fi109++ {
                                    var v110 float64 = float64(-2.25)
                                    _ = len(v3)
                                    v111 := GenS6{F0: -993}
                                    _ = fi109
                                    _ = v110
                                    _ = v111
                                }
                            default:
                                var v112 int = 436
                                _ = v112
                            }
                            v113 := GenS8{F0: true}
                            _ = v107
                            _ = v113
                        }
                        _ = fi72
                    }
                    _ = v71
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v1 = v4
                }
                for fi114 := 0; fi114 < 9; fi114++ {
                    v4 = "a"
                    var v115 float64 = (0.1 * 1e10)
                    _ = fi114
                    _ = v115
                }
            }
            v42 += v42
            v1 = ((v4 + "") + v4)
        }
    }
    switch 4294967295 {
    case 0:
        v4 = "hello"
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
