package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 uint64; F2 int64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 bool; F1 int32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int; F1 int16; F2 int }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS11 struct { F0 int8; F1 int }
type GenS12 struct { F0 uint16; F1 int }
type GenS13 struct { F0 uint8; F1 int32; F2 int64 }
func genFn0(p0 int, p1 uint16, p2 string) int8 {
    p1 += p1
    v4 := GenS1{F0: "hello"}
    v5 := []int{24, (p0 >> ((4294967295) & 63)), p0, p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 125
}
func genFn1() int32 {
    for fi6 := 0; fi6 < 2; fi6++ {
        for fi7 := 0; fi7 < 4; fi7++ {
            var v8 float32 = (((-2.25 + 3.14159) - ((0.1 / 1e10) - -2.25)) + -2.25)
            switch ("hello" + "") {
            case "a":
                switch (("hello" + "a") + ("\n\t" + "")) {
                case "hello":
                    var v9 uint = 1
                    _ = v9
                case "":
                    fmt.Println(fi7)
                case "\n\t":
                    v10 := GenS1{F0: ("a" + ("\n\t" + (("\n\t" + "foo bar") + "hello")))}
                    _ = v10
                }
            case "\n\t":
                if false {
                    var v11 uint16 = 2
                    var v12 int = fi6
                    _ = v11
                    _ = v12
                }
            case "foo bar":
                if false {
                    v8 += ((((v8 / v8) - v8) * v8) + v8)
                } else {
                    var v13 uint = 42
                    var v14 int = (fi6 + (^(fi7)))
                    fmt.Println(fi6, v14)
                    _ = v13
                    _ = v14
                }
                v8 = v8
            }
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    var v15 int16 = 32767
    v16 := []int{2147483647}
    _ = len(v16)
    for fi17 := 0; fi17 < 8; fi17++ {
        var v18 int = (fi17 << ((255) & 63))
        _ = fi17
        _ = v18
    }
    _ = v15
    _ = v16
    return -1
}

func main() {
    var v1 uint = 4294967295
    var v2 uint16 = 869
    var v3 bool = (1 > 18446744073709551615)
    v19 := GenS1{F0: ("\n\t" + "foo bar")}
    var v20 int32 = 350
    v21 := []int{0, 931}
    v3 = (uint32(v2) <= 911)
    switch "foo bar" {
    case "\n\t":
        var v22 int32 = genFn1()
        _ = v22
    default:
        v21 = append(v21, 128)
    }
    for fi23 := 0; fi23 < 16; fi23++ {
        var v24 uint8 = 2
        _ = fi23
        _ = v24
    }
    var v25 uint64 = 0
    for fi26 := 0; fi26 < 6; fi26++ {
        _ = len(v21)
        var v27 int16 = int16(v25)
        for fi28 := 0; fi28 < 14; fi28++ {
            v1 = (v1 / ((v1) | 1))
            _ = fi28
        }
        _ = fi26
        _ = v27
    }
    fmt.Println(v3)
    var v29 uint = (v1 | (v1 & v1))
    if v3 {
        for fi30 := 0; fi30 < 17; fi30++ {
            if true {
                var v31 uint16 = v2
                if len(v21) >= 2 {
                    v21 = v21[1:len(v21)]
                }
                for fi32 := 0; fi32 < 6; fi32++ {
                    if v3 {
                        v20 = genFn1()
                        v21 = append(v21, (^(fi30)))
                        fmt.Println(v25, fi30)
                    } else {
                        if (-1 == 127) {
                            var v33 int64 = -9223372036854775808
                            if ((^(v29)) >= uint(v25)) {
                                v34 := GenS2{F0: fi32, F1: (v25 | (v25 + (^(v25)))), F2: 723}
                                var v35 uint8 = 127
                                _ = v34
                                _ = v35
                            } else {
                                v21 = append(v21, fi30)
                                _ = len(v21)
                            }
                            var v36 uint8 = 106
                            v37 := GenS1{F0: "\n\t"}
                            v36 -= v36
                            _ = v33
                            _ = v36
                            _ = v37
                        } else {
                            if (65535 <= 0) {
                                v1 += v1
                                v25 = (v25 * (v25 << ((v29) & 63)))
                                var v38 int64 = 1
                                if v3 {
                                    fmt.Println(v31, v1)
                                    if len(v21) >= 2 {
                                        v21 = v21[1:len(v21)]
                                    }
                                    v29 = (v29 % ((v1) | 1))
                                } else {
                                    v21 = append(v21, (fi32 >> (((v29 + 18446744073709551615)) & 63)))
                                    var v39 int64 = v38
                                    v21 = append(v21, (fi32 ^ fi32))
                                    _ = v39
                                }
                                _ = v38
                            } else {
                                if v3 {
                                    var v40 bool = false
                                    var v41 float64 = float64(1.5)
                                    _ = v40
                                    _ = v41
                                } else {
                                    var v42 uint = 255
                                    v31 -= 255
                                    _ = v42
                                }
                                if (v25 >= (v25 << (((v1 % ((v1) | 1))) & 63))) {
                                    v2 = 32767
                                    v29 = v29
                                    v43 := GenS1{F0: "foo bar"}
                                    _ = v43
                                } else {
                                    v2 = v2
                                    if len(v21) >= 2 {
                                        v21 = v21[1:len(v21)]
                                    }
                                    if (0 == 127) {
                                        v21 = append(v21, fi32)
                                    } else {
                                        if len(v21) >= 2 {
                                            v21 = v21[1:len(v21)]
                                        }
                                        var v44 int8 = 1
                                        _ = v44
                                    }
                                }
                            }
                        }
                        if (true && v3) {
                            var v45 uint = (v1 % ((564) | 1))
                            v46 := GenS1{F0: "hello"}
                            if (!(v3 && true)) {
                                v29 += v1
                                var v47 int16 = -190
                                if v3 {
                                    var v48 int = (fi32 - fi32)
                                    switch ("hello" + "a") {
                                    case "\n\t":
                                        if len(v21) >= 2 {
                                            v21 = v21[1:len(v21)]
                                        }
                                    case "foo bar":
                                        var v49 uint8 = uint8(v1)
                                        _ = v49
                                    case "":
                                        v29 -= 392
                                    default:
                                        var v50 int64 = 0
                                        _ = v50
                                    }
                                    v51 := GenS3{F0: -397}
                                    _ = v48
                                    _ = v51
                                } else {
                                    if len(v21) >= 2 {
                                        v21 = v21[1:len(v21)]
                                    }
                                    var v52 float64 = 1e10
                                    _ = v52
                                }
                                v53 := GenS4{F0: 1}
                                v25 -= v25
                                v29 -= v1
                                _ = v47
                                _ = v53
                            } else {
                                v54 := GenS2{F0: fi30, F1: v25, F2: -1}
                                _ = v54
                            }
                            _ = v45
                            _ = v46
                        }
                    }
                    fmt.Println(v20, v1, v25)
                    var v55 float32 = 3.14159
                    if v3 {
                        fmt.Println(v2, v3)
                        if (v3 || false) {
                            var v56 int8 = 100
                            var v57 uint32 = 573
                            _ = v56
                            _ = v57
                        }
                        v58 := GenS2{F0: (fi32 / (((fi32 >> ((v29) & 63))) | 1)), F1: v25, F2: 777}
                        _ = v58
                    }
                    _ = fi32
                    _ = v55
                }
                _ = v31
            } else {
                for fi59 := 0; fi59 < 17; fi59++ {
                    v60 := GenS5{F0: v3, F1: (v20 - v20)}
                    switch "a" {
                    case "a":
                        var v61 uint8 = 127
                        _ = v61
                    case "foo bar":
                        _ = len(v21)
                        v62 := GenS1{F0: "\n\t"}
                        var v63 float64 = 1.5
                        _ = v62
                        _ = v63
                    }
                    _ = fi59
                    _ = v60
                }
                var v64 float64 = 0.0
                _ = v64
            }
            if ("" == "foo bar") {
                var v65 int = (-(fi30))
                var v66 bool = ((v2 / (((v2 + (v2 & v2))) | 1)) != (v2 >> (((v29 >> (((v1 - (v29 ^ (v29 % (((v29 & (v1 - v29))) | 1))))) & 63))) & 63)))
                _ = v65
                _ = v66
            } else {
                v20 = -215
            }
            var v67 int32 = v20
            _ = fi30
            _ = v67
        }
    } else {
        if (v25 >= uint64(v1)) {
            for fi68 := 0; fi68 < 16; fi68++ {
                if (!(222 < 42)) {
                    if (uint8(v2) >= 128) {
                        var v69 uint16 = v2
                        var v70 int32 = 42
                        if v3 {
                            v1 = v29
                            if (425 <= 2147483647) {
                                v71 := GenS2{F0: (^(fi68)), F1: (v25 | v25), F2: -822}
                                if (2147483648 < 4294967295) {
                                    for fi72 := 0; fi72 < 10; fi72++ {
                                        v25 = (v25 + 4294967295)
                                        _ = fi72
                                    }
                                    v2 -= 255
                                    for fi73 := 0; fi73 < 13; fi73++ {
                                        if (float32(0.1) == float32(-2.25)) {
                                            var v74 float64 = ((float64(float32(float64(float32((1.5 * 0.0))))) + ((((1.5 / ((1e10 - 0.1) * float64((1.5 + 1e10)))) - 1.5) + (0.1 - (0.1 - 3.14159))) + 0.1)) / 1.0)
                                            fmt.Println(fi73, v70)
                                            _ = v74
                                        }
                                        if (42 <= v25) {
                                            var v75 uint = (v29 / ((42) | 1))
                                            v2 = v69
                                            _ = v75
                                        } else {
                                            if true {
                                                v20 -= (v20 >> ((uint(v69)) & 63))
                                                var v76 int32 = (v70 * (v70 / (((^(v20))) | 1)))
                                                fmt.Println(fi68, v70)
                                                _ = v76
                                            }
                                        }
                                        _ = fi73
                                    }
                                } else {
                                    if false {
                                        var v77 int32 = 589
                                        v21 = append(v21, (fi68 ^ 128))
                                        _ = v77
                                    } else {
                                        var v78 uint8 = 42
                                        v79 := GenS2{F0: (fi68 >> ((540) & 63)), F1: (v25 * v25), F2: 68}
                                        _ = v78
                                        _ = v79
                                    }
                                }
                                v3 = (0 == -9223372036854775808)
                                _ = v71
                            } else {
                                if v3 {
                                    if len(v21) >= 2 {
                                        v21 = v21[1:len(v21)]
                                    }
                                    if len(v21) > 0 {
                                        v21[0] = (fi68 << (((v1 % (((v29 - (v1 & (v29 / (((v29 | (v29 * (^(v29))))) | 1))))) | 1))) & 63))
                                    }
                                    v21 = append(v21, fi68)
                                }
                                switch "\n\t" {
                                case "hello":
                                    v80 := GenS6{F0: 1}
                                    _ = v80
                                }
                                v70 -= 42
                            }
                            for fi81 := 0; fi81 < 10; fi81++ {
                                var v82 uint16 = 605
                                v21 = append(v21, fi68)
                                _ = fi81
                                _ = v82
                            }
                        } else {
                            for fi83 := 0; fi83 < 12; fi83++ {
                                _ = len(v21)
                                _ = fi83
                            }
                        }
                        _ = v69
                        _ = v70
                    } else {
                        if ((!v3) && (-43 < -62)) {
                            for fi84 := 0; fi84 < 4; fi84++ {
                                if (1 != 42) {
                                    fmt.Println(v1)
                                    v25 += v25
                                    v1 -= v29
                                } else {
                                    if (!(!(3.14159 > 0.1))) {
                                        var v85 bool = v3
                                        _ = v85
                                    }
                                    v2 -= (^(v2))
                                }
                                var v86 uint8 = 2
                                _ = fi84
                                _ = v86
                            }
                            v3 = false
                            v2 += 65535
                        } else {
                            var v87 uint64 = (v25 * v25)
                            _ = v87
                        }
                    }
                    v88 := GenS7{F0: fi68, F1: 261, F2: fi68}
                    _ = v88
                } else {
                    for fi89 := 0; fi89 < 11; fi89++ {
                        if true {
                            v90 := GenS1{F0: (("hello" + "a") + "\n\t")}
                            _ = v90
                        } else {
                            switch "foo bar" {
                            case "\n\t":
                                v91 := GenS3{F0: 1}
                                if (0 < 451) {
                                    v25 = v25
                                }
                                if v3 {
                                    _ = len(v21)
                                    if false {
                                        v2 -= v2
                                        var v92 int64 = 0
                                        switch ("\n\t" + "foo bar") {
                                        case "a":
                                            v2 += 2
                                        case "\n\t":
                                            if (!true) {
                                                fmt.Println(fi89, v29)
                                            } else {
                                                var v93 uint64 = v25
                                                var v94 uint8 = 127
                                                v29 = (v1 % (((v1 / ((65535) | 1))) | 1))
                                                _ = v93
                                                _ = v94
                                            }
                                        case "hello":
                                            var v95 uint16 = (v2 + 0)
                                            _ = v95
                                        default:
                                            if len(v21) > 0 {
                                                v21[0] = (fi89 >> ((v29) & 63))
                                            }
                                        }
                                        _ = v92
                                    } else {
                                        fmt.Println(v20, v25)
                                    }
                                    v20 = v20
                                    var v96 uint16 = (v2 - v2)
                                    var v97 uint = (^(v1))
                                    _ = v96
                                    _ = v97
                                }
                                _ = v91
                            case "hello":
                                v25 -= 534
                            }
                        }
                        v98 := GenS8{F0: fi89}
                        _ = fi89
                        _ = v98
                    }
                    for fi99 := 0; fi99 < 4; fi99++ {
                        var v100 int16 = int16(v20)
                        v101 := GenS9{F0: v100}
                        v2 = 65535
                        _ = fi99
                        _ = v100
                        _ = v101
                    }
                }
                v1 = (v1 & (^(v1)))
                _ = fi68
            }
            v3 = (42 <= 1000)
            v3 = (128 >= 9223372036854775807)
        } else {
            switch -9223372036854775808 {
            case -862:
                for fi102 := 0; fi102 < 16; fi102++ {
                    v1 -= v1
                    _ = fi102
                }
            case 1:
                v21 = append(v21, 128)
            case -9223372036854775808:
                if len(v21) > 0 {
                    v21[0] = 576
                }
            }
            if v3 {
                v103 := GenS10{F0: (v1 % ((v1) | 1)), F1: (v2 & v2), F2: 1}
                var v104 float32 = 0.0
                _ = v103
                _ = v104
            }
        }
        if (4294967295 <= -1) {
            for fi105 := 0; fi105 < 19; fi105++ {
                if ("a" < ((("foo bar" + "foo bar") + "foo bar") + ("" + ("hello" + "hello")))) {
                    var v106 bool = (-2147483648 == 1)
                    _ = v106
                }
                if true {
                    fmt.Println(v25, v2, v2)
                    var v107 bool = (!(v3 || ((v2 % ((v2) | 1)) != (v2 ^ v2))))
                    v2 = (v2 | uint16(v25))
                    _ = v107
                } else {
                    for fi108 := 0; fi108 < 17; fi108++ {
                        var v109 float64 = (0.0 * 0.1)
                        var v110 int8 = -126
                        var v111 uint16 = v2
                        _ = fi108
                        _ = v109
                        _ = v110
                        _ = v111
                    }
                }
                _ = fi105
            }
            v112 := GenS11{F0: 100, F1: 1}
            v113 := GenS12{F0: v2, F1: 0}
            _ = v112
            _ = v113
        } else {
            if v3 {
                var v114 int = 32768
                v114 = (v114 & v114)
                _ = v114
            } else {
                v20 -= v20
            }
            if (-804 < 565) {
                switch -1 {
                case 758:
                    var v115 int64 = -9223372036854775808
                    _ = v115
                }
                _ = len(v21)
            }
            for fi116 := 0; fi116 < 2; fi116++ {
                var v117 string = "a"
                _ = fi116
                _ = v117
            }
        }
        var v118 float32 = -2.25
        _ = v118
    }
    var v119 int16 = 0
    var v120 int64 = 1
    fmt.Println(v120, v1, v1)
    v119 -= 962
    v121 := GenS13{F0: 135, F1: (v20 << ((uint(v25)) & 63)), F2: (v120 - v120)}
    v1 = (v1 - (v29 << ((1) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v119:", v119)
    fmt.Println("v120:", v120)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v119
    _ = v120
    _ = v121
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v29
    _ = v3
}
