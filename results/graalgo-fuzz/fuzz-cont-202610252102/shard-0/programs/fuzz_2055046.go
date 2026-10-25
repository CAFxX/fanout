package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 uint16; F1 int32; F2 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS5 struct { F0 uint8; F1 int8 }
type GenS6 struct { F0 bool; F1 string; F2 uint64 }
type GenS7 struct { F0 int8; F1 int; F2 uint8 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int64; F1 int8; F2 uint8 }
type GenS10 struct { F0 uint64; F1 int64; F2 int }
type GenS11 struct { F0 bool }
func genFn0() uint {
    for fi4 := 0; fi4 < 1; fi4++ {
        for fi5 := 0; fi5 < 14; fi5++ {
            var v6 int8 = 1
            var v7 int64 = 1
            _ = fi5
            _ = v6
            _ = v7
        }
        _ = fi4
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        if true {
            var v9 bool = false
            _ = v9
        }
        v10 := GenS1{F0: (-32768 < -1), F1: 42}
        _ = fi8
        _ = v10
    }
    switch "a" {
    case "hello":
        var v11 int32 = 1
        _ = v11
    default:
        var v12 uint8 = 127
        _ = v12
    }
    var v13 uint16 = 319
    _ = v13
    return 2
}
func genFn1(p0 int64, p1 uint8, p2 int) bool {
    fmt.Println(p0, p2)
    p0 -= p0
    _ = p0
    _ = p1
    _ = p2
    return true
}

func main() {
    var v1 uint = 540
    v2 := make(map[string]bool)
    var v3 uint8 = 188
    if false {
        v14 := []int{-2147483648}
        v1 += 65535
        var v15 uint = v1
        _ = v14
        _ = v15
    } else {
        v16 := GenS2{F0: 94, F1: 42, F2: v1}
        v17 := GenS3{F0: 5}
        _ = v16
        _ = v17
    }
    if false {
        fmt.Println(v1, v3)
        if false {
            v18 := []int{-1, -497, -702, int(v3)}
            if (289 == 852) {
                var v19 int = 1
                _ = v19
            }
            if (genFn0() <= (v1 >> ((v1) & 63))) {
                var v20 float32 = -2.25
                switch ("hello" + "\n\t") {
                case "\n\t":
                    v1 = 65535
                case "":
                    var v21 int8 = -128
                    _ = v21
                case "a":
                    v3 -= v3
                default:
                    var v22 uint16 = 1
                    _ = v22
                }
                if true {
                    v1 += v1
                    var v23 int16 = 42
                    _ = v23
                } else {
                    var v24 int = -349
                    if (false || (1000 < 42)) {
                        var v25 int64 = 9223372036854775807
                        if true {
                            for fi26 := 0; fi26 < 6; fi26++ {
                                delete(v2, "hello")
                                fmt.Println(v20, v25)
                                v27 := GenS4{F0: 42, F1: 4294967295, F2: -1}
                                _ = fi26
                                _ = v27
                            }
                            v24 = (v24 - -660)
                            for fi28 := 0; fi28 < 11; fi28++ {
                                switch 128 {
                                case -781:
                                    v25 = 966
                                case 1:
                                    v3 = (v3 / (((v3 % ((v3) | 1))) | 1))
                                case 9223372036854775807:
                                    v25 = (^(v25))
                                }
                                fmt.Println(v25)
                                var v29 uint8 = v3
                                _ = fi28
                                _ = v29
                            }
                        } else {
                            var v30 uint8 = (v3 >> (((v1 | (^(v1)))) & 63))
                            _ = v30
                        }
                        _ = v25
                    } else {
                        switch (v24 ^ 0) {
                        case -790:
                            switch v24 {
                            case 0:
                                var v31 uint16 = 1
                                _ = v31
                            default:
                                if len(v18) >= 2 {
                                    v18 = v18[1:len(v18)]
                                }
                            }
                        case 164:
                            if true {
                                for fi32 := 0; fi32 < 20; fi32++ {
                                    fmt.Println(v3, v24)
                                    switch (v24 * (-(fi32))) {
                                    case 0:
                                        if (1 != 0) {
                                            v33 := GenS2{F0: 65535, F1: 0, F2: v1}
                                            _ = v33
                                        }
                                    case 341:
                                        v18 = append(v18, (v24 / (((v24 + (v24 * 9223372036854775807))) | 1)))
                                    case -2147483648:
                                        _ = len(v2)
                                    }
                                    for fi34 := 0; fi34 < 18; fi34++ {
                                        var v35 string = "hello"
                                        _ = fi34
                                        _ = v35
                                    }
                                    _ = fi32
                                }
                                for fi36 := 0; fi36 < 17; fi36++ {
                                    switch "hello" {
                                    case "hello":
                                        v3 = v3
                                    case "":
                                        var v37 float32 = v20
                                        _ = v37
                                    case "foo bar":
                                        v38 := v2["a"]
                                        _ = v38
                                    }
                                    v3 += (v3 * (v3 & (v3 ^ (v3 & v3))))
                                    _ = fi36
                                }
                                var v39 uint32 = 625
                                _ = v39
                            } else {
                                if false {
                                    var v40 uint8 = 30
                                    v41 := v2["\n\t"]
                                    var v42 uint16 = 2
                                    _ = v40
                                    _ = v41
                                    _ = v42
                                } else {
                                    v20 += v20
                                }
                                if (!true) {
                                    var v43 int8 = -1
                                    if len(v18) >= 2 {
                                        v18 = v18[1:len(v18)]
                                    }
                                    _ = v43
                                } else {
                                    for fi44 := 0; fi44 < 5; fi44++ {
                                        var v45 float64 = 3.14159
                                        if (v3 != v3) {
                                            if true {
                                                for fi46 := 0; fi46 < 16; fi46++ {
                                                    v47 := GenS5{F0: (^(v3)), F1: 127}
                                                    var v48 bool = (42 == 2147483647)
                                                    var v49 string = "\n\t"
                                                    _ = fi46
                                                    _ = v47
                                                    _ = v48
                                                    _ = v49
                                                }
                                                v2[("a" + "hello")] = (156 == 1)
                                                v45 = v45
                                            } else {
                                                var v50 float32 = v20
                                                var v51 uint64 = 4294967295
                                                _ = v50
                                                _ = v51
                                            }
                                        } else {
                                            v20 += -2.25
                                            if false {
                                                if (9223372036854775807 >= -9223372036854775808) {
                                                    v52 := GenS4{F0: 698, F1: 2, F2: -1}
                                                    v20 = v20
                                                    switch "a" {
                                                    case "a":
                                                        if ((false && true) && true) {
                                                            if false {
                                                                var v53 int8 = -13
                                                                v20 += (v20 + v20)
                                                                _ = v53
                                                            } else {
                                                                var v54 uint8 = v3
                                                                switch (("\n\t" + ("a" + "a")) + "foo bar") {
                                                                case "":
                                                                    if len(v18) >= 2 {
                                                                        v18 = v18[1:len(v18)]
                                                                    }
                                                                }
                                                                _ = v54
                                                            }
                                                        }
                                                    case "":
                                                        for fi55 := 0; fi55 < 2; fi55++ {
                                                            if false {
                                                                if (-467 <= -1) {
                                                                    v56 := GenS1{F0: true, F1: 100}
                                                                    v20 = v20
                                                                    v3 += v3
                                                                    _ = v56
                                                                } else {
                                                                    v45 -= (v45 + (v45 + (float64(1e10) + float64(v20))))
                                                                    var v57 uint = v1
                                                                    switch (fi44 / ((0) | 1)) {
                                                                    case 0:
                                                                        var v58 int8 = -85
                                                                        _ = v58
                                                                    case 631:
                                                                        v24 += (-(fi55))
                                                                    case 128:
                                                                        var v59 bool = (2 != 489)
                                                                        _ = v59
                                                                    }
                                                                    _ = v57
                                                                }
                                                                var v60 bool = ((^(v3)) == (v3 % ((131) | 1)))
                                                                v61 := GenS5{F0: v3, F1: 42}
                                                                _ = v60
                                                                _ = v61
                                                            } else {
                                                                if (857 != 35) {
                                                                    var v62 uint32 = 2147483647
                                                                    _ = v62
                                                                }
                                                                v18 = append(v18, -392)
                                                                v63 := GenS6{F0: true, F1: "\n\t", F2: 1000}
                                                                v64 := GenS7{F0: -90, F1: fi55, F2: v3}
                                                                _ = v63
                                                                _ = v64
                                                            }
                                                            if ((v3 <= v3) && ("hello" >= "")) {
                                                                if true {
                                                                    v65 := GenS2{F0: uint16(v1), F1: 2147483647, F2: 778}
                                                                    var v66 uint64 = 816
                                                                    var v67 int16 = -499
                                                                    var v68 int = 127
                                                                    if len(v18) >= 2 {
                                                                        v18 = v18[1:len(v18)]
                                                                    }
                                                                    fmt.Println(v68)
                                                                    _ = v65
                                                                    _ = v66
                                                                    _ = v67
                                                                    _ = v68
                                                                } else {
                                                                    fmt.Println(v20)
                                                                    v20 = 0.1
                                                                    v69 := GenS8{F0: "foo bar"}
                                                                    v70 := GenS9{F0: 9223372036854775807, F1: 0, F2: 2}
                                                                    _ = v69
                                                                    _ = v70
                                                                }
                                                                if true {
                                                                    var v71 int = fi55
                                                                    var v72 int16 = 42
                                                                    var v73 uint32 = 0
                                                                    _ = v71
                                                                    _ = v72
                                                                    _ = v73
                                                                }
                                                                var v74 int16 = 32767
                                                                _ = v74
                                                            } else {
                                                                v3 = v3
                                                                if true {
                                                                    v45 = (((1e10 + v45) * (1.5 - ((v45 + v45) - 1.5))) + v45)
                                                                    var v75 uint32 = 683
                                                                    _ = v75
                                                                } else {
                                                                    v3 = (v3 ^ v3)
                                                                }
                                                            }
                                                            switch "hello" {
                                                            case "a":
                                                                v3 = (v3 ^ 2)
                                                            case "hello":
                                                                if len(v18) > 0 {
                                                                    v18[0] = v24
                                                                }
                                                            default:
                                                                if true {
                                                                    switch -613 {
                                                                    case 9223372036854775807:
                                                                        if false {
                                                                            v20 += (float32((float64(v20) + v45)) * (0.1 * (v20 + v20)))
                                                                            fmt.Println(v20, v45, v24)
                                                                            v20 = (v20 + (v20 - 1e10))
                                                                        } else {
                                                                            v1 = 322
                                                                        }
                                                                    }
                                                                    fmt.Println(fi44)
                                                                    switch "a" {
                                                                    case "a":
                                                                        if (!(!false)) {
                                                                            v1 = (v1 * v1)
                                                                            v2[((("foo bar" + "foo bar") + ("hello" + "a")) + "")] = false
                                                                        } else {
                                                                            var v76 int8 = 100
                                                                            var v77 string = "hello"
                                                                            _ = v76
                                                                            _ = v77
                                                                        }
                                                                    case "":
                                                                        v1 = 374
                                                                    default:
                                                                        v3 -= v3
                                                                    }
                                                                    switch 42 {
                                                                    case 9223372036854775807:
                                                                        v18 = append(v18, (fi55 & (fi55 >> (((v1 / (((v1 << (((^(v1))) & 63))) | 1))) & 63))))
                                                                    }
                                                                }
                                                            }
                                                            _ = fi55
                                                        }
                                                    case "hello":
                                                        if (!true) {
                                                            v45 -= ((v45 + 3.14159) * float64((float32(((float64(v20) * float64(3.14159)) - (float64((0.0 + float32(-2.25))) - v45))) + (3.14159 - (((v20 / (1e10 + v20)) / (v20 * v20)) - -2.25)))))
                                                            var v78 int32 = -1
                                                            _ = v78
                                                        }
                                                    }
                                                    _ = v52
                                                }
                                                v1 = (v1 << ((v1) & 63))
                                                switch (v24 / (((v24 - v24)) | 1)) {
                                                case -9223372036854775808:
                                                    if true {
                                                        v79 := v2[("\n\t" + "foo bar")]
                                                        switch (fi44 >> (((v1 * 707)) & 63)) {
                                                        case -1:
                                                            if false {
                                                                delete(v2, "\n\t")
                                                                if (((((v20 / v20) + 0.0) - v20) - v20) == v20) {
                                                                    var v80 int16 = -32768
                                                                    fmt.Println(v1)
                                                                    v81 := GenS10{F0: 4294967295, F1: -9223372036854775808, F2: (v24 & (v24 - (v24 - (v24 % ((fi44) | 1)))))}
                                                                    _ = v80
                                                                    _ = v81
                                                                } else {
                                                                    v79 = v79
                                                                    var v82 int8 = 0
                                                                    _ = v82
                                                                }
                                                            } else {
                                                                if v79 {
                                                                    var v83 uint16 = 1
                                                                    _ = v83
                                                                }
                                                                fmt.Println(v79)
                                                                fmt.Println(v24, v24)
                                                            }
                                                        default:
                                                            var v84 uint32 = 49
                                                            _ = v84
                                                        }
                                                        v85 := GenS3{F0: -11}
                                                        _ = v79
                                                        _ = v85
                                                    }
                                                default:
                                                    var v86 int8 = 0
                                                    _ = v86
                                                }
                                            } else {
                                                v1 = 4294967295
                                                v87 := GenS3{F0: 42}
                                                _ = len(v18)
                                                _ = v87
                                            }
                                        }
                                        _ = fi44
                                        _ = v45
                                    }
                                    var v88 uint32 = 2147483648
                                    var v89 float32 = v20
                                    _ = v88
                                    _ = v89
                                }
                            }
                        }
                        var v90 uint = (^(v1))
                        var v91 int8 = 54
                        _ = v90
                        _ = v91
                    }
                    _ = v24
                }
                _ = v20
            } else {
                for fi92 := 0; fi92 < 20; fi92++ {
                    var v93 uint8 = v3
                    _ = fi92
                    _ = v93
                }
                var v94 int32 = 0
                _ = v94
            }
            _ = v18
        }
    }
    if false {
        var v95 int64 = -855
        var v96 uint32 = 4294967295
        _ = v95
        _ = v96
    }
    for fi97 := 0; fi97 < 5; fi97++ {
        for fi98 := 0; fi98 < 2; fi98++ {
            switch (("hello" + "a") + "foo bar") {
            case "hello":
                v3 = v3
            case "foo bar":
                switch 4294967295 {
                case 760:
                    var v99 int16 = -1
                    _ = v99
                case 585:
                    v100 := GenS11{F0: ((94 <= 0) && true)}
                    v1 += 255
                    v1 = v1
                    _ = v100
                }
                var v101 int = fi97
                _ = v101
            default:
                var v102 string = "\n\t"
                _ = v102
            }
            _ = fi98
        }
        _ = fi97
    }
    v103 := []int{595, -775}
    var v104 uint8 = (v3 % ((uint8(v1)) | 1))
    v2[""] = false
    for fi105 := 0; fi105 < 19; fi105++ {
        v106 := GenS11{F0: true}
        switch (-(fi105)) {
        case 0:
            switch fi105 {
            case -2147483648:
                for fi107 := 0; fi107 < 18; fi107++ {
                    var v108 int32 = 1
                    _ = fi107
                    _ = v108
                }
            default:
                switch int(v3) {
                case 2147483647:
                    v1 = 1
                case 42:
                    for fi109 := 0; fi109 < 9; fi109++ {
                        fmt.Println(v3, fi109, v1)
                        _ = fi109
                    }
                }
            }
        case -9223372036854775808:
            var v110 int8 = -1
            _ = v110
        default:
            switch (^(fi105)) {
            case 42:
                _ = len(v2)
            case -378:
                for fi111 := 0; fi111 < 3; fi111++ {
                    if len(v103) >= 2 {
                        v103 = v103[1:len(v103)]
                    }
                    _ = fi111
                }
            case 127:
                v2["foo bar"] = (-2147483648 <= -900)
            }
        }
        _ = fi105
        _ = v106
    }
    v104 = v3
    if false {
        fmt.Println(v1, v1)
        for fi112 := 0; fi112 < 17; fi112++ {
            for fi113 := 0; fi113 < 4; fi113++ {
                v2[("\n\t" + (((("\n\t" + "") + ("hello" + (((("a" + "\n\t") + ("hello" + "hello")) + "") + "a"))) + "") + ""))] = true
                v3 = 137
                _ = fi113
            }
            _ = fi112
        }
        var v114 uint64 = 2
        _ = v114
    }
    v104 -= (v104 | (v104 - (v3 + (v3 ^ (v3 / (((^(v3))) | 1))))))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v103:", len(v103), v103)
    fmt.Println("v104:", v104)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v103
    _ = v104
    _ = v2
    _ = v3
}
