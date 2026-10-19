package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64; F1 int32 }
type GenS5 struct { F0 string; F1 bool; F2 int64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint64; F1 int64 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int32; F1 int16; F2 int16 }
type GenS10 struct { F0 bool; F1 int64 }
type GenS11 struct { F0 int8; F1 int; F2 int }
type GenS12 struct { F0 string; F1 uint8 }
func genFn0() int64 {
    var v7 float64 = 0.0
    var v8 uint64 = 18446744073709551615
    _ = v7
    _ = v8
    return 9223372036854775807
}

func main() {
    v1 := []int{0}
    v2 := make(map[int]bool)
    v3 := []string{"hello", ("" + "\n\t"), "\n\t", "\n\t"}
    var v4 int = -2147483648
    var v5 int16 = 42
    v6 := []string{("" + "\n\t")}
    for fi9 := 0; fi9 < 11; fi9++ {
        fmt.Println(v4)
        if false {
            if (!(false && false)) {
                for fi10 := 0; fi10 < 15; fi10++ {
                    if len(v1) > 0 {
                        v1[0] = int(v5)
                    }
                    if (-457 > 0) {
                        switch "\n\t" {
                        case "a":
                            if true {
                                var v11 int8 = 0
                                _ = len(v2)
                                var v12 float64 = 0.0
                                _ = v11
                                _ = v12
                            }
                        case "foo bar":
                            if false {
                                fmt.Println(fi9, fi10)
                                var v13 float64 = 1.5
                                var v14 int8 = 19
                                v13 -= float64((0.0 - 1.5))
                                var v15 bool = true
                                _ = v13
                                _ = v14
                                _ = v15
                            } else {
                                var v16 uint = 18446744073709551615
                                var v17 uint32 = 2147483648
                                _ = v16
                                _ = v17
                            }
                        case "hello":
                            v18 := GenS1{F0: 32767, F1: 42}
                            _ = v18
                        default:
                            _ = len(v2)
                        }
                        var v19 uint = 4294967295
                        _ = v19
                    }
                    switch 2147483647 {
                    case 0:
                        if len(v1) > 0 {
                            v1[0] = -9223372036854775808
                        }
                    case 984:
                        var v20 uint = 0
                        _ = v20
                    }
                    _ = fi10
                }
                var v21 uint16 = 580
                _ = v21
            } else {
                for fi22 := 0; fi22 < 19; fi22++ {
                    fmt.Println(fi9)
                    _ = fi22
                }
            }
            for fi23 := 0; fi23 < 6; fi23++ {
                delete(v2, fi23)
                v24 := GenS1{F0: 667, F1: 42}
                _ = fi23
                _ = v24
            }
            for fi25 := 0; fi25 < 9; fi25++ {
                v26 := GenS1{F0: 32767, F1: 871}
                if len(v3) > 0 {
                    v3[0] = "hello"
                }
                var v27 uint16 = 255
                _ = fi25
                _ = v26
                _ = v27
            }
        } else {
            var v28 int16 = (v5 ^ (v5 - v5))
            var v29 int = (-(fi9))
            _ = v28
            _ = v29
        }
        if (0 < 2) {
            if true {
                v5 = (v5 - (v5 / (((v5 ^ v5)) | 1)))
                if false {
                    var v30 uint16 = 32767
                    switch fi9 {
                    case -233:
                        var v31 int8 = 127
                        _ = v31
                    case 42:
                        var v32 float64 = 0.1
                        _ = v32
                    default:
                        v33 := GenS2{F0: false}
                        _ = v33
                    }
                    _ = v30
                }
                if true {
                    var v34 bool = (2147483647 != 65535)
                    var v35 int64 = 42
                    _ = v34
                    _ = v35
                } else {
                    fmt.Println(fi9, v4, v4)
                    delete(v2, (v4 / ((v4) | 1)))
                }
            }
            var v36 string = "a"
            _ = v36
        }
        _ = fi9
    }
    fmt.Println(v4, v5)
    if (-1 <= int64(v4)) {
        switch 2147483647 {
        case 9223372036854775807:
            for fi37 := 0; fi37 < 4; fi37++ {
                var v38 int64 = genFn0()
                if (((1.5 + 1e10) > ((1.5 + float32((0.1 / 1.5))) / 1.5)) && (42 == 245)) {
                    v4 = int(v5)
                    v2[32767] = (1 <= 216)
                    for fi39 := 0; fi39 < 12; fi39++ {
                        v38 -= v38
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        v40 := GenS3{F0: (2147483647 > -1)}
                        _ = fi39
                        _ = v40
                    }
                }
                if true {
                    var v41 float32 = 1.5
                    var v42 uint64 = 2
                    var v43 uint64 = (^(v42))
                    _ = v41
                    _ = v42
                    _ = v43
                } else {
                    v3 = append(v3, "a")
                    v2[v4] = (true && (!false))
                }
                _ = fi37
                _ = v38
            }
        case 42:
            v44 := GenS2{F0: (false || (!((v5 >> ((1) & 63)) != -1)))}
            _ = v44
        default:
            var v45 uint8 = 42
            _ = v45
        }
        var v46 uint8 = 255
        v47 := GenS3{F0: (1.5 > 0.1)}
        _ = v46
        _ = v47
    } else {
        v48 := GenS4{F0: 783, F1: 727}
        _ = v48
    }
    for fi49 := 0; fi49 < 12; fi49++ {
        var v50 float64 = 1e10
        delete(v2, fi49)
        if true {
            if false {
                var v51 uint8 = 0
                var v52 int = (v4 << ((469) & 63))
                v53 := GenS2{F0: true}
                _ = v51
                _ = v52
                _ = v53
            }
        }
        _ = fi49
        _ = v50
    }
    v4 = v4
    var v54 int8 = 0
    for fi55 := 0; fi55 < 8; fi55++ {
        var v56 float32 = (1.5 + ((1.5 / 1.0) - 0.1))
        v56 -= v56
        var v57 uint16 = 32767
        _ = fi55
        _ = v56
        _ = v57
    }
    for fi58 := 0; fi58 < 8; fi58++ {
        if true {
            if false {
                v59 := GenS5{F0: "", F1: false, F2: -9223372036854775808}
                v60 := GenS4{F0: 1, F1: -855}
                _ = v59
                _ = v60
            } else {
                v1 = append(v1, fi58)
                if true {
                    v54 += v54
                    fmt.Println(v4, v54, v4)
                } else {
                    for fi61 := 0; fi61 < 4; fi61++ {
                        v62 := GenS6{F0: 647}
                        v63 := GenS7{F0: 453, F1: 746}
                        var v64 uint = 694
                        _ = fi61
                        _ = v62
                        _ = v63
                        _ = v64
                    }
                    _ = len(v2)
                    v5 = (v5 / ((v5) | 1))
                }
            }
            var v65 string = "\n\t"
            _ = v65
        } else {
            delete(v2, -499)
            if (true || (42 <= -362)) {
                v5 = (v5 & -32768)
            }
        }
        v5 = v5
        _ = fi58
    }
    v54 -= 1
    v66 := v2[(v4 << ((4294967295) & 63))]
    v5 += (v5 ^ v5)
    v54 -= -128
    var v67 string = "a"
    if (v66 || v66) {
        v54 = v54
        if v66 {
            fmt.Println(v5)
            _ = len(v6)
            if false {
                for fi68 := 0; fi68 < 1; fi68++ {
                    v69 := GenS8{F0: 42}
                    _ = fi68
                    _ = v69
                }
                for fi70 := 0; fi70 < 15; fi70++ {
                    v71 := GenS9{F0: 0, F1: v5, F2: (-(v5))}
                    var v72 bool = v66
                    fmt.Println(v66)
                    _ = fi70
                    _ = v71
                    _ = v72
                }
            }
        }
        switch v4 {
        case 32768:
            if ((v5 >= (v5 << ((65535) & 63))) && ((2147483647 < 1) && ((v5 ^ -32768) > v5))) {
                if ((((-2.25 * (1.5 * 0.1)) - float32(1e10)) <= 1.5) && (65535 >= 659)) {
                    if ((127 != (^(v54))) && (32767 == -2147483648)) {
                        var v73 string = v67
                        for fi74 := 0; fi74 < 13; fi74++ {
                            v75 := GenS8{F0: 32767}
                            _ = fi74
                            _ = v75
                        }
                        if v66 {
                            v76 := v2[v4]
                            var v77 int16 = v5
                            v78 := GenS1{F0: 0, F1: 255}
                            _ = v76
                            _ = v77
                            _ = v78
                        }
                        _ = v73
                    } else {
                        if v66 {
                            if (247 > 7) {
                                v66 = v66
                                for fi79 := 0; fi79 < 12; fi79++ {
                                    if v66 {
                                        v80 := v2[v4]
                                        switch "a" {
                                        case "foo bar":
                                            if v80 {
                                                fmt.Println(fi79)
                                                v54 -= (v54 | v54)
                                                fmt.Println(v66, v67, v67)
                                            } else {
                                                if (!true) {
                                                    v4 += (^(v4))
                                                } else {
                                                    var v81 string = ("\n\t" + ((v67 + v67) + (("hello" + v67) + "a")))
                                                    var v82 uint = 478
                                                    for fi83 := 0; fi83 < 1; fi83++ {
                                                        v84 := GenS6{F0: 65535}
                                                        v54 = v54
                                                        _ = fi83
                                                        _ = v84
                                                    }
                                                    _ = v81
                                                    _ = v82
                                                }
                                                if v80 {
                                                    var v85 uint16 = 0
                                                    switch v67 {
                                                    case "":
                                                        delete(v2, (fi79 + v4))
                                                    case "hello":
                                                        fmt.Println(v67, v67)
                                                    case "a":
                                                        var v86 bool = (1 > -2147483648)
                                                        _ = v86
                                                    default:
                                                        _ = len(v2)
                                                    }
                                                    var v87 int8 = v54
                                                    _ = v85
                                                    _ = v87
                                                }
                                            }
                                        case "\n\t":
                                            var v88 int = -2147483648
                                            _ = v88
                                        case "":
                                            var v89 int64 = genFn0()
                                            _ = v89
                                        default:
                                            var v90 int32 = 0
                                            _ = v90
                                        }
                                        _ = v80
                                    } else {
                                        fmt.Println(v54, fi79)
                                    }
                                    var v91 float32 = (1.5 + 3.14159)
                                    if (-2147483648 == -2147483648) {
                                        v92 := GenS1{F0: 691, F1: 65535}
                                        var v93 bool = (!((812 == 42) && (!v66)))
                                        switch -2147483648 {
                                        case 791:
                                            if false {
                                                switch "a" {
                                                case "foo bar":
                                                    v94 := GenS10{F0: (false && v93), F1: 42}
                                                    _ = v94
                                                default:
                                                    v67 = ""
                                                }
                                            } else {
                                                switch 32767 {
                                                case 128:
                                                    if len(v1) > 0 {
                                                        v1[0] = -1
                                                    }
                                                }
                                                fmt.Println(v54)
                                            }
                                        case 485:
                                            for fi95 := 0; fi95 < 14; fi95++ {
                                                var v96 int16 = (v5 & (v5 + (v5 + v5)))
                                                _ = fi95
                                                _ = v96
                                            }
                                        case 0:
                                            var v97 uint64 = 0
                                            _ = v97
                                        default:
                                            if v66 {
                                                fmt.Println(fi79, v91)
                                            }
                                        }
                                        _ = v92
                                        _ = v93
                                    }
                                    _ = fi79
                                    _ = v91
                                }
                            }
                        }
                        v98 := GenS8{F0: 1}
                        if v66 {
                            for fi99 := 0; fi99 < 6; fi99++ {
                                switch (v67 + "\n\t") {
                                case "\n\t":
                                    v100 := GenS11{F0: v54, F1: -2147483648, F2: 0}
                                    _ = v100
                                case "a":
                                    if (float32(float64(3.14159)) < -2.25) {
                                        var v101 uint16 = 680
                                        if v66 {
                                            v67 = "hello"
                                        }
                                        _ = v101
                                    }
                                case "hello":
                                    v102 := GenS12{F0: v67, F1: 2}
                                    _ = v102
                                }
                                _ = fi99
                            }
                            for fi103 := 0; fi103 < 4; fi103++ {
                                var v104 int32 = -1
                                _ = fi103
                                _ = v104
                            }
                            v66 = (0.1 != (0.0 * 0.0))
                        }
                        _ = v98
                    }
                    v105 := GenS3{F0: (18446744073709551615 < 0)}
                    _ = v105
                } else {
                    var v106 int16 = v5
                    var v107 int8 = 42
                    _ = v106
                    _ = v107
                }
                for fi108 := 0; fi108 < 4; fi108++ {
                    fmt.Println(v5, v67, v4)
                    _ = fi108
                }
                v67 = (v67 + v67)
            } else {
                v66 = (("\n\t" + v67) > "")
            }
        }
    }
    if true {
        v67 = v67
        v109 := GenS8{F0: 1}
        fmt.Println(v5, v5, v5)
        _ = v109
    } else {
        v66 = (v66 || v66)
        var v110 int64 = 1
        v66 = (false && false)
        _ = v110
    }
    var v111 float32 = 0.0
    var v112 int32 = 0
    v5 += v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v111:", v111)
    fmt.Println("v112:", v112)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v111
    _ = v112
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v6
    _ = v66
    _ = v67
}
