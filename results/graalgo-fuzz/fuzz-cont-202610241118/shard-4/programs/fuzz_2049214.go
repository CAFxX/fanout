package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 uint64; F1 int8; F2 uint }
type GenS5 struct { F0 uint16; F1 int32; F2 int64 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int64; F1 string }
type GenS8 struct { F0 bool; F1 int64 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 int; F1 uint64 }
type GenS12 struct { F0 int; F1 string }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 int16 }
func genFn0() uint64 {
    v4 := GenS1{F0: false}
    v5 := GenS1{F0: (-1 >= 9223372036854775807)}
    v6 := GenS2{F0: 598}
    var v7 bool = (false && (42 == 0))
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 uint64) int16 {
    switch ("\n\t" + "foo bar") {
    case "":
        if (1 < 255) {
            switch "a" {
            case "":
                p0 = 79
            case "\n\t":
                switch "a" {
                case "hello":
                    for fi8 := 0; fi8 < 15; fi8++ {
                        for fi9 := 0; fi9 < 2; fi9++ {
                            v10 := []int{fi8}
                            _ = fi9
                            _ = v10
                        }
                        _ = fi8
                    }
                default:
                    v11 := []int{9223372036854775807, 32767, 4294967295, 134}
                    _ = v11
                }
            case "a":
                var v12 bool = false
                _ = v12
            default:
                var v13 float64 = 3.14159
                _ = v13
            }
            var v14 uint16 = 629
            v15 := []int{920, 127}
            _ = v14
            _ = v15
        } else {
            fmt.Println(p0)
        }
    case "a":
        for fi16 := 0; fi16 < 12; fi16++ {
            var v17 bool = true
            fmt.Println(fi16, fi16)
            var v18 string = "foo bar"
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    switch 954 {
    case 32767:
        if (!true) {
            p0 = 4294967295
            v19 := []int{-770, -2147483648}
            v20 := GenS3{F0: (32767 > 195), F1: (!(1 <= -37))}
            _ = v19
            _ = v20
        } else {
            v21 := []int{4294967295}
            _ = len(v21)
            fmt.Println(p0)
            _ = v21
        }
    case -189:
        v22 := GenS2{F0: 637}
        _ = v22
    case 0:
        var v23 uint32 = 42
        _ = v23
    }
    _ = p0
    return -367
}

func main() {
    var v1 uint64 = 684
    v2 := make(map[string]int)
    var v3 int64 = 42
    delete(v2, (("" + "a") + "\n\t"))
    if true {
        fmt.Println(v1)
        fmt.Println(v1, v1)
    } else {
        if true {
            v24 := GenS4{F0: (^(v1)), F1: 1, F2: 2}
            v1 += 18446744073709551615
            var v25 uint8 = 255
            _ = v24
            _ = v25
        }
    }
    if (!true) {
        var v26 int64 = -1
        _ = len(v2)
        v1 = v1
        _ = v26
    } else {
        v27 := GenS5{F0: 32767, F1: 42, F2: v3}
        _ = v27
    }
    if (!((v3 < (v3 % ((-1) | 1))) && (127 < -9223372036854775808))) {
        for fi28 := 0; fi28 < 4; fi28++ {
            v1 = (v1 << ((422) & 63))
            for fi29 := 0; fi29 < 19; fi29++ {
                if false {
                    fmt.Println(fi28, fi28)
                    v30 := GenS6{F0: (v1 - 42)}
                    var v31 int = fi29
                    v1 -= (v1 % ((v1) | 1))
                    if ((1.5 + (float32(3.14159) * float32((0.0 / 1.5)))) != (float32(3.14159) / 1.0)) {
                        if (!false) {
                            var v32 int16 = 649
                            v3 = v3
                            switch fi29 {
                            case 9223372036854775807:
                                switch 128 {
                                case 128:
                                    if false {
                                        var v33 float64 = (((1e10 - (((((((1e10 / ((0.1 / (3.14159 + 1.5)) - 1e10)) * (1.5 - 0.1)) + (float64(float32(-2.25)) * (-2.25 - 0.1))) + 0.1) + (0.1 * 3.14159)) + float64((1e10 - 1e10))) * 0.1)) - -2.25) - (-2.25 / 3.14159))
                                        if (1 != 420) {
                                            v34 := GenS1{F0: (255 != 213)}
                                            if true {
                                                v32 -= int16(v1)
                                            } else {
                                                v3 = 42
                                            }
                                            _ = v34
                                        } else {
                                            delete(v2, (("hello" + ("\n\t" + "a")) + "hello"))
                                            fmt.Println(v31, fi28)
                                        }
                                        var v35 bool = true
                                        _ = v33
                                        _ = v35
                                    }
                                case -9223372036854775808:
                                    v36 := GenS1{F0: (!(false || (!(!((42 >= 279) && true)))))}
                                    if true {
                                        var v37 uint8 = 255
                                        _ = v37
                                    }
                                    if (!(!true)) {
                                        _ = len(v2)
                                        fmt.Println(v3, v1, v3)
                                        var v38 string = "\n\t"
                                        _ = v38
                                    }
                                    _ = v36
                                default:
                                    var v39 string = ""
                                    _ = v39
                                }
                            }
                            delete(v2, "")
                            if true {
                                var v40 int8 = 111
                                _ = v40
                            }
                            _ = v32
                        } else {
                            v41 := GenS7{F0: (v3 ^ (^(v3))), F1: "hello"}
                            _ = v41
                        }
                        v42 := GenS6{F0: v1}
                        if false {
                            v1 = v1
                            v31 -= v31
                        }
                        _ = v42
                    } else {
                        if (0.1 >= -2.25) {
                            var v43 bool = true
                            var v44 bool = (!(false && (((((1.5 - 1.5) * 1.5) * (3.14159 - -2.25)) * 0.1) != 3.14159)))
                            if false {
                                var v45 float32 = (-2.25 * 1.5)
                                switch "" {
                                case "a":
                                    v46 := []int{4294967295}
                                    _ = v46
                                case "\n\t":
                                    if (genFn0() != 584) {
                                        switch "\n\t" {
                                        case "a":
                                            if (957 > 32767) {
                                                v47 := GenS6{F0: (v1 | genFn0())}
                                                fmt.Println(v44)
                                                if ((v3 >> ((4294967295) & 63)) > v3) {
                                                    if false {
                                                        fmt.Println(v44, v31)
                                                        if (v31 < 127) {
                                                            if (-32768 <= -32768) {
                                                                var v48 bool = false
                                                                v44 = true
                                                                v31 -= v31
                                                                v3 = (v3 - v3)
                                                                v3 += (v3 / ((v3) | 1))
                                                                _ = v48
                                                            } else {
                                                                var v49 int64 = v3
                                                                v50 := GenS5{F0: 0, F1: 42, F2: 861}
                                                                fmt.Println(fi28, fi28, v49)
                                                                _ = v49
                                                                _ = v50
                                                            }
                                                            v51 := []int{fi29}
                                                            _ = v51
                                                        }
                                                        v52 := GenS4{F0: (v1 % ((genFn0()) | 1)), F1: 1, F2: 255}
                                                        _ = v52
                                                    } else {
                                                        v53 := []int{(fi29 ^ fi28), fi28, (fi29 % ((127) | 1))}
                                                        _ = v53
                                                    }
                                                    var v54 uint32 = 42
                                                    _ = v54
                                                } else {
                                                    delete(v2, "a")
                                                    _ = len(v2)
                                                    fmt.Println(fi28)
                                                }
                                                _ = v47
                                            }
                                        case "\n\t":
                                            v55 := GenS3{F0: v44, F1: (("\n\t" + "\n\t") > "foo bar")}
                                            _ = v55
                                        default:
                                            var v56 uint64 = genFn0()
                                            _ = len(v2)
                                            _ = v56
                                        }
                                        v57 := GenS8{F0: v44, F1: -1}
                                        delete(v2, ("a" + "a"))
                                        _ = v57
                                    }
                                default:
                                    if true {
                                        delete(v2, "\n\t")
                                        switch fi28 {
                                        case -454:
                                            v58 := GenS4{F0: (v1 & (^(v1))), F1: 1, F2: 1}
                                            _ = v58
                                        default:
                                            fmt.Println(v1, fi29, fi29)
                                        }
                                    }
                                }
                                _ = v45
                            }
                            var v59 uint16 = 32767
                            v1 = v1
                            var v60 string = ("foo bar" + "\n\t")
                            if v43 {
                                var v61 int = fi29
                                v59 = 0
                                v62 := v2[(v60 + "foo bar")]
                                _ = v61
                                _ = v62
                            }
                            _ = v43
                            _ = v44
                            _ = v59
                            _ = v60
                        }
                        switch (v31 * v31) {
                        case 42:
                            v1 = v1
                            v2["foo bar"] = (^(v31))
                            fmt.Println(v31, v31, v3)
                        case 0:
                            v63 := GenS9{F0: v1}
                            if (4294967295 > 670) {
                                switch "hello" {
                                case "foo bar":
                                    var v64 bool = false
                                    _ = v64
                                }
                                var v65 uint = 255
                                _ = v65
                            }
                            v31 = -9223372036854775808
                            _ = v63
                        case 591:
                            v66 := GenS8{F0: (42 == v3), F1: (v3 - (v3 * v3))}
                            _ = v66
                        default:
                            v67 := GenS4{F0: (v1 / ((v1) | 1)), F1: 48, F2: 893}
                            _ = v67
                        }
                    }
                    var v68 bool = (true && true)
                    fmt.Println(fi28)
                    switch fi29 {
                    case 939:
                        var v69 bool = v68
                        _ = v69
                    }
                    var v70 float32 = 1.5
                    var v71 int = 128
                    if (!(true || v68)) {
                        var v72 int = 32767
                        var v73 int = v72
                        _ = v72
                        _ = v73
                    }
                    if v68 {
                        v71 = (-(fi28))
                        v71 = (-(fi28))
                        var v74 string = ("hello" + "\n\t")
                        var v75 uint8 = 255
                        fmt.Println(v70, v74)
                        fmt.Println(v31)
                        _ = v74
                        _ = v75
                    } else {
                        if false {
                            v1 -= (v1 & 4294967295)
                            if v68 {
                                var v76 int64 = -9223372036854775808
                                v77 := GenS1{F0: (0.0 > 0.0)}
                                var v78 float32 = v70
                                _ = v76
                                _ = v77
                                _ = v78
                            } else {
                                v79 := GenS10{F0: 7}
                                v80 := GenS9{F0: (v1 | (v1 - 0))}
                                switch (fi29 >> ((0) & 63)) {
                                case 32768:
                                    v81 := v2["\n\t"]
                                    _ = v81
                                case -9223372036854775808:
                                    var v82 int32 = -55
                                    _ = v82
                                case 9223372036854775807:
                                    var v83 bool = (!v68)
                                    _ = v83
                                }
                                _ = v79
                                _ = v80
                            }
                        } else {
                            var v84 int16 = -1
                            _ = v84
                        }
                    }
                    _ = v30
                    _ = v31
                    _ = v68
                    _ = v70
                    _ = v71
                }
                v85 := GenS11{F0: int(v1), F1: (^(v1))}
                v86 := v2[((((("\n\t" + "a") + "foo bar") + "\n\t") + (("" + "a") + "foo bar")) + (((("\n\t" + (("foo bar" + "hello") + (((("" + "hello") + (("" + ("hello" + "hello")) + (((("" + "foo bar") + (("a" + "foo bar") + "a")) + "") + ("a" + "foo bar")))) + ("hello" + "hello")) + ""))) + "\n\t") + "a") + "hello"))]
                _ = fi29
                _ = v85
                _ = v86
            }
            var v87 float64 = float64(3.14159)
            _ = fi28
            _ = v87
        }
        v3 = v3
        v1 = v1
    }
    v88 := GenS6{F0: 855}
    for fi89 := 0; fi89 < 8; fi89++ {
        if false {
            fmt.Println(fi89, v3)
            var v90 float32 = 3.14159
            _ = v90
        } else {
            fmt.Println(fi89, fi89, v1)
            v91 := GenS6{F0: (v1 | (v1 / ((v1) | 1)))}
            _ = v91
        }
        var v92 int16 = 42
        v93 := GenS7{F0: v3, F1: (("a" + "") + "")}
        _ = fi89
        _ = v92
        _ = v93
    }
    v3 += 0
    v94 := []int{-9223372036854775808}
    if (-2147483648 == 9223372036854775807) {
        if (464 <= 713) {
            if false {
                if (!(1 > 127)) {
                    v1 = (^(v1))
                    v1 -= 2
                    if (2 >= 255) {
                        for fi95 := 0; fi95 < 1; fi95++ {
                            for fi96 := 0; fi96 < 19; fi96++ {
                                v97 := GenS12{F0: 127, F1: "hello"}
                                v3 = (v3 % ((v3) | 1))
                                _ = fi96
                                _ = v97
                            }
                            var v98 string = "foo bar"
                            if true {
                                fmt.Println(v1)
                            } else {
                                v1 -= v1
                                v99 := GenS3{F0: true, F1: (0 == 255)}
                                _ = v99
                            }
                            _ = fi95
                            _ = v98
                        }
                    } else {
                        if true {
                            v3 -= (-(v3))
                        }
                    }
                } else {
                    v100 := GenS13{F0: 0}
                    _ = v100
                }
            } else {
                v101 := GenS4{F0: genFn0(), F1: -89, F2: 255}
                fmt.Println(v3, v1)
                _ = v101
            }
            fmt.Println(v1, v3)
        } else {
            if false {
                switch 9223372036854775807 {
                case 9223372036854775807:
                    v102 := GenS14{F0: 18446744073709551615}
                    _ = v102
                }
                if (true || false) {
                    var v103 uint32 = 4294967295
                    _ = v103
                } else {
                    for fi104 := 0; fi104 < 4; fi104++ {
                        var v105 int64 = (v3 / ((v3) | 1))
                        _ = len(v94)
                        var v106 uint = 2
                        _ = fi104
                        _ = v105
                        _ = v106
                    }
                }
            } else {
                v1 -= genFn0()
                v3 = v3
            }
            if (!(!(2147483648 >= 42))) {
                fmt.Println(v3, v3)
            }
            var v107 uint64 = v1
            _ = v107
        }
        if true {
            fmt.Println(v3)
            _ = len(v2)
            for fi108 := 0; fi108 < 20; fi108++ {
                for fi109 := 0; fi109 < 5; fi109++ {
                    v110 := GenS15{F0: 32767}
                    var v111 uint = 751
                    _ = fi109
                    _ = v110
                    _ = v111
                }
                _ = fi108
            }
        } else {
            if (9223372036854775807 <= -230) {
                var v112 uint8 = 230
                _ = v112
            } else {
                _ = len(v94)
                if len(v94) >= 2 {
                    v94 = v94[1:len(v94)]
                }
                if (-2.25 != 1e10) {
                    _ = len(v2)
                } else {
                    v1 -= v1
                }
            }
            v94 = append(v94, 1)
            v1 = v1
        }
    }
    fmt.Println(v1, v1)
    var v113 uint16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v113:", v113)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v94:", len(v94), v94)
    _ = v1
    _ = v113
    _ = v2
    _ = v3
    _ = v88
    _ = v94
}
