package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint32; F1 int; F2 int }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 uint16; F1 uint16; F2 uint64 }
type GenS5 struct { F0 int; F1 uint32; F2 uint16 }
type GenS6 struct { F0 string; F1 int8 }
type GenS7 struct { F0 uint8; F1 bool }
type GenS8 struct { F0 uint64 }
func genFn0(p0 uint32) uint32 {
    var v7 uint = 42
    v7 += (^(v7))
    _ = p0
    _ = v7
    return p0
}
func genFn1(p0 int, p1 int32, p2 bool) uint16 {
    var v8 uint = 255
    v8 -= v8
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 0
}

func main() {
    var v1 bool = ((2 > 2) && (!false))
    v2 := make(map[string]string)
    v3 := make(map[uint32]string)
    var v4 int8 = 42
    v5 := make(map[int]int)
    v6 := make(map[uint32]uint32)
    switch "hello" {
    case "":
        v9 := GenS1{F0: -737, F1: 42}
        _ = v9
    case "foo bar":
        for fi10 := 0; fi10 < 16; fi10++ {
            var v11 uint32 = 2147483648
            _ = fi10
            _ = v11
        }
    }
    v4 += 42
    var v12 uint8 = 2
    v13 := []int{419}
    delete(v3, 2147483647)
    var v14 int16 = -1
    v12 = v12
    if (-2147483648 > 0) {
        var v15 int64 = -311
        if (v1 && v1) {
            v6[1] = 65535
        }
        _ = v15
    } else {
        if (!(0.0 != -2.25)) {
            switch (("\n\t" + ("" + "a")) + "") {
            case "foo bar":
                for fi16 := 0; fi16 < 2; fi16++ {
                    v17 := GenS1{F0: (v14 << ((18446744073709551615) & 63)), F1: 65535}
                    if true {
                        for fi18 := 0; fi18 < 14; fi18++ {
                            _ = len(v13)
                            _ = fi18
                        }
                    } else {
                        for fi19 := 0; fi19 < 7; fi19++ {
                            fmt.Println(v1, fi19, fi19)
                            _ = fi19
                        }
                    }
                    v20 := GenS2{F0: 42, F1: fi16, F2: int(v14)}
                    _ = fi16
                    _ = v17
                    _ = v20
                }
            case "a":
                switch "foo bar" {
                case "foo bar":
                    if false {
                        fmt.Println(v14, v14)
                        var v21 uint8 = (v12 / (((v12 + (v12 * v12))) | 1))
                        if (!v1) {
                            v22 := GenS3{F0: v4, F1: 370}
                            for fi23 := 0; fi23 < 19; fi23++ {
                                var v24 uint32 = genFn0(215)
                                if v1 {
                                    v25 := GenS3{F0: v4, F1: 633}
                                    _ = v25
                                }
                                if v1 {
                                    var v26 int8 = v4
                                    var v27 uint64 = 609
                                    _ = v26
                                    _ = v27
                                }
                                _ = fi23
                                _ = v24
                            }
                            v1 = v1
                            _ = v22
                        } else {
                            var v28 int64 = -425
                            _ = v28
                        }
                        _ = v21
                    }
                case "a":
                    var v29 uint = 0
                    _ = v29
                }
            case "":
                switch "a" {
                case "":
                    v13 = append(v13, 0)
                default:
                    var v30 uint8 = (v12 | (^(v12)))
                    _ = v30
                }
            }
            if (116 == 2147483648) {
                v1 = ((((float64(1.5) / 1e10) - (1e10 * 1e10)) + 1e10) != 0.0)
                if v1 {
                    for fi31 := 0; fi31 < 9; fi31++ {
                        if (fi31 >= fi31) {
                            var v32 uint8 = v12
                            if v1 {
                                for fi33 := 0; fi33 < 5; fi33++ {
                                    v32 -= v12
                                    if (42 == 65535) {
                                        v34 := GenS4{F0: 0, F1: 2, F2: 0}
                                        fmt.Println(v4, fi33)
                                        _ = v34
                                    } else {
                                        var v35 int32 = -1
                                        v14 -= -909
                                        v12 = (^(v32))
                                        switch "a" {
                                        case "a":
                                            v36 := GenS1{F0: v14, F1: 18446744073709551615}
                                            _ = v36
                                        }
                                        _ = v35
                                    }
                                    _ = fi33
                                }
                                var v37 int = (fi31 % ((fi31) | 1))
                                _ = v37
                            }
                            v38 := GenS2{F0: 2147483648, F1: 1, F2: fi31}
                            _ = v32
                            _ = v38
                        } else {
                            for fi39 := 0; fi39 < 18; fi39++ {
                                v4 = v4
                                fmt.Println(v1, v12)
                                var v40 int32 = int32(v4)
                                _ = fi39
                                _ = v40
                            }
                            var v41 uint16 = 2
                            _ = v41
                        }
                        for fi42 := 0; fi42 < 6; fi42++ {
                            v14 += (v14 - 1)
                            switch ("\n\t" + "a") {
                            case "hello":
                                if v1 {
                                    var v43 uint8 = 2
                                    fmt.Println(v1)
                                    _ = v43
                                } else {
                                    var v44 int8 = v4
                                    fmt.Println(v44, v14, v14)
                                    v13 = append(v13, (fi42 + fi31))
                                    _ = v44
                                }
                            case "":
                                if (true || true) {
                                    var v45 uint32 = 2
                                    _ = v45
                                } else {
                                    if v1 {
                                        var v46 int8 = (v4 >> ((2) & 63))
                                        fmt.Println(v4, v46)
                                        if (1.5 > -2.25) {
                                            switch fi42 {
                                            case 32767:
                                                fmt.Println(fi42, v1)
                                            default:
                                                v14 += (v14 << ((255) & 63))
                                            }
                                            v12 = v12
                                            var v47 int32 = 1
                                            _ = v47
                                        } else {
                                            var v48 int16 = (-(v14))
                                            fmt.Println(fi42)
                                            v12 = (v12 - v12)
                                            _ = v48
                                        }
                                        switch "" {
                                        case "\n\t":
                                            switch "a" {
                                            case "a":
                                                if false {
                                                    delete(v2, ("" + "\n\t"))
                                                    v46 = 0
                                                    switch "hello" {
                                                    case "\n\t":
                                                        v46 += -1
                                                    }
                                                }
                                            case "foo bar":
                                                v14 = v14
                                            }
                                        case "hello":
                                            var v49 bool = (-76 != -1)
                                            _ = v49
                                        default:
                                            fmt.Println(fi42, v12, v46)
                                        }
                                        fmt.Println(v1)
                                        _ = v46
                                    } else {
                                        v50 := GenS5{F0: fi31, F1: 162, F2: 212}
                                        _ = v50
                                    }
                                    var v51 uint64 = 523
                                    v51 = (^(v51))
                                    _ = v51
                                }
                            case "foo bar":
                                var v52 float32 = 0.0
                                _ = v52
                            }
                            _ = fi42
                        }
                        switch (("a" + "foo bar") + "\n\t") {
                        case "":
                            for fi53 := 0; fi53 < 3; fi53++ {
                                switch ("" + "a") {
                                case "a":
                                    if v1 {
                                        v54 := v5[fi53]
                                        fmt.Println(v1, v12)
                                        v55 := GenS4{F0: genFn1(fi31, 2147483647, (v1 && false)), F1: 42, F2: 42}
                                        if (genFn1(fi31, -688, v1) >= 65535) {
                                            var v56 uint8 = (v12 & 255)
                                            if v1 {
                                                _ = len(v13)
                                                var v57 float32 = 1e10
                                                _ = v57
                                            }
                                            var v58 uint32 = 591
                                            _ = v56
                                            _ = v58
                                        }
                                        _ = v54
                                        _ = v55
                                    }
                                default:
                                    var v59 int16 = (v14 * (v14 & -1))
                                    _ = v59
                                }
                                v14 = 42
                                _ = fi53
                            }
                        case "hello":
                            if false {
                                for fi60 := 0; fi60 < 6; fi60++ {
                                    if false {
                                        var v61 int = fi60
                                        if (!((v12 * 128) >= v12)) {
                                            v2[("foo bar" + "")] = "foo bar"
                                        } else {
                                            fmt.Println(v4, v1, v14)
                                            if false {
                                                switch (^(fi60)) {
                                                case 1:
                                                    switch ("" + "hello") {
                                                    case "a":
                                                        v3[2147483648] = "foo bar"
                                                    case "hello":
                                                        v4 += int8(v14)
                                                        v62 := v3[genFn0(42)]
                                                        _ = v62
                                                    case "\n\t":
                                                        if false {
                                                            v63 := GenS3{F0: (v4 >> ((811) & 63)), F1: genFn1(fi60, -791, (float64(0.1) == (1e10 / (-2.25 + 3.14159))))}
                                                            if v1 {
                                                                var v64 uint16 = 0
                                                                var v65 int16 = -1
                                                                fmt.Println(v64, v1)
                                                                _ = v64
                                                                _ = v65
                                                            }
                                                            _ = v63
                                                        } else {
                                                            v4 = v4
                                                        }
                                                    }
                                                case -810:
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        v66 := v6[241]
                                                        _ = v66
                                                    case "":
                                                        var v67 float32 = 0.0
                                                        _ = v67
                                                    case "a":
                                                        v4 += v4
                                                    default:
                                                        var v68 int16 = -1
                                                        _ = v68
                                                    }
                                                default:
                                                    var v69 uint8 = 73
                                                    _ = v69
                                                }
                                            } else {
                                                v61 = fi60
                                                v14 = (v14 * v14)
                                            }
                                            v4 = v4
                                            if (1 >= 255) {
                                                var v70 uint8 = 106
                                                v71 := v2["\n\t"]
                                                _ = v70
                                                _ = v71
                                            }
                                            v14 += v14
                                        }
                                        _ = v61
                                    } else {
                                        var v72 uint16 = 0
                                        var v73 uint16 = genFn1((fi31 >> ((18446744073709551615) & 63)), 128, true)
                                        _ = v72
                                        _ = v73
                                    }
                                    if (genFn1(fi60, -213, v1) == genFn1(811, 1, (true && (!v1)))) {
                                        if false {
                                            var v74 int16 = v14
                                            fmt.Println(v4)
                                            _ = v74
                                        } else {
                                            _ = len(v3)
                                            var v75 int8 = 16
                                            v12 = 1
                                            _ = v75
                                        }
                                        switch (fi31 | fi31) {
                                        case 42:
                                            fmt.Println(v1)
                                        case 0:
                                            if (v14 == v14) {
                                                if v1 {
                                                    v1 = v1
                                                }
                                                v14 = v14
                                                if v1 {
                                                    if len(v13) > 0 {
                                                        v13[0] = fi31
                                                    }
                                                    v1 = (!true)
                                                    var v76 int16 = (v14 / ((0) | 1))
                                                    _ = v76
                                                }
                                                if v1 {
                                                    var v77 float32 = (1.5 + -2.25)
                                                    if len(v13) >= 2 {
                                                        v13 = v13[1:len(v13)]
                                                    }
                                                    _ = v77
                                                } else {
                                                    switch 42 {
                                                    case 4294967295:
                                                        var v78 uint64 = 106
                                                        _ = v78
                                                    }
                                                    fmt.Println(v14)
                                                }
                                                if len(v13) > 0 {
                                                    v13[0] = fi60
                                                }
                                            } else {
                                                v79 := GenS2{F0: 511, F1: (-(fi31)), F2: 1}
                                                var v80 float32 = 1e10
                                                _ = v79
                                                _ = v80
                                            }
                                        case -9223372036854775808:
                                            var v81 uint32 = 65535
                                            v12 = v12
                                            v4 += v4
                                            _ = v81
                                        }
                                        delete(v2, "\n\t")
                                        if len(v13) >= 2 {
                                            v13 = v13[1:len(v13)]
                                        }
                                        if len(v13) > 0 {
                                            v13[0] = (^(fi60))
                                        }
                                    } else {
                                        v82 := GenS6{F0: "a", F1: (v4 % ((v4) | 1))}
                                        _ = v82
                                    }
                                    _ = fi60
                                }
                                for fi83 := 0; fi83 < 14; fi83++ {
                                    v84 := GenS2{F0: 65535, F1: (fi31 >> ((uint(v12)) & 63)), F2: fi31}
                                    _ = fi83
                                    _ = v84
                                }
                            } else {
                                switch "\n\t" {
                                case "foo bar":
                                    switch "\n\t" {
                                    case "foo bar":
                                        v4 = (v4 << ((42) & 63))
                                    case "a":
                                        var v85 int = fi31
                                        _ = v85
                                    }
                                }
                                for fi86 := 0; fi86 < 15; fi86++ {
                                    switch "" {
                                    case "foo bar":
                                        v1 = v1
                                    case "\n\t":
                                        var v87 int = (-(fi86))
                                        _ = v87
                                    default:
                                        if (-9223372036854775808 != -9223372036854775808) {
                                            v4 += (v4 / ((v4) | 1))
                                            var v88 uint64 = 42
                                            if len(v13) >= 2 {
                                                v13 = v13[1:len(v13)]
                                            }
                                            _ = v88
                                        }
                                    }
                                    _ = fi86
                                }
                                if ((0.0 + (0.0 * (((((3.14159 / 1e10) + 0.1) / 1.0) - 0.1) + (float32(3.14159) / 1e10)))) < float32(float64((float32(3.14159) - -2.25)))) {
                                    var v89 float32 = -2.25
                                    v4 += -1
                                    _ = v89
                                }
                            }
                        default:
                            var v90 uint32 = 65535
                            _ = v90
                        }
                        _ = fi31
                    }
                    v91 := GenS7{F0: (^(v12)), F1: true}
                    var v92 float64 = (1e10 - float64((((float32(-2.25) - float32((-2.25 + ((1.5 * 0.0) - -2.25)))) / ((0.1 - -2.25) - 1.5)) - -2.25)))
                    _ = v91
                    _ = v92
                } else {
                    var v93 bool = v1
                    _ = v93
                }
            } else {
                if len(v13) > 0 {
                    v13[0] = -2147483648
                }
                v12 = (v12 / (((v12 * (v12 & (v12 / ((42) | 1))))) | 1))
                var v94 uint = 18446744073709551615
                _ = v94
            }
            var v95 uint16 = 165
            _ = v95
        } else {
            for fi96 := 0; fi96 < 1; fi96++ {
                var v97 uint16 = 80
                v98 := v2[""]
                _ = fi96
                _ = v97
                _ = v98
            }
            var v99 int32 = 42
            _ = v99
        }
        var v100 uint = 255
        _ = v100
    }
    fmt.Println(v14)
    v12 = (v12 + 255)
    v101 := GenS5{F0: 32767, F1: 859, F2: 2}
    v102 := GenS6{F0: "foo bar", F1: v4}
    v103 := GenS5{F0: 32767, F1: 872, F2: genFn1(4294967295, 602, true)}
    v1 = false
    var v104 uint = 255
    var v105 int32 = 0
    switch "foo bar" {
    case "\n\t":
        _ = len(v13)
    case "":
        v106 := v6[1]
        _ = v106
    case "foo bar":
        var v107 int = 37
        _ = v107
    default:
        switch "\n\t" {
        case "a":
            v108 := v6[genFn0(uint32(v104))]
            _ = v108
        }
    }
    v109 := GenS8{F0: 780}
    fmt.Println("v1:", v1)
    fmt.Println("v104:", v104)
    fmt.Println("v105:", v105)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v101
    _ = v102
    _ = v103
    _ = v104
    _ = v105
    _ = v109
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
