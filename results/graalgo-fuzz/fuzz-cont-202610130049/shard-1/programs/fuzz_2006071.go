package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint8; F1 int64; F2 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int; F1 int64 }
type GenS6 struct { F0 int64; F1 bool; F2 uint8 }
type GenS7 struct { F0 int16; F1 uint32 }
type GenS8 struct { F0 int8; F1 string }
type GenS9 struct { F0 int16; F1 int; F2 int }
type GenS10 struct { F0 int64; F1 int }
type GenS11 struct { F0 int; F1 uint8; F2 int64 }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 int64 }
func genFn0() int32 {
    var v5 uint8 = 2
    if false {
        v5 = (v5 - (^(v5)))
    }
    v5 -= (^(v5))
    fmt.Println(v5)
    _ = v5
    return -1
}
func genFn1(p0 int16, p1 int8, p2 uint64) string {
    var v6 bool = (42 < 42)
    v7 := []int{-547}
    fmt.Println(p2, p2, p0)
    v7 = append(v7, 127)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return ("hello" + "")
}

func main() {
    v1 := []int{-2147483648, -9223372036854775808, 127, -907}
    var v2 int32 = 0
    v3 := []string{}
    var v4 uint16 = 2
    var v8 int16 = 42
    v4 = (v4 * v4)
    if false {
        for fi9 := 0; fi9 < 2; fi9++ {
            fmt.Println(v8, v8)
            v10 := GenS1{F0: false, F1: 596}
            var v11 int16 = (v8 | (v8 & -1))
            _ = fi9
            _ = v10
            _ = v11
        }
        var v12 bool = false
        if v12 {
            fmt.Println(v2, v2, v8)
            fmt.Println(v4, v12)
        } else {
            var v13 float64 = (0.0 * 1.5)
            fmt.Println(v2, v13)
            _ = v13
        }
        _ = v12
    } else {
        var v14 uint32 = 2147483648
        v14 = 4294967295
        var v15 int64 = 0
        _ = v14
        _ = v15
    }
    v8 = v8
    v16 := GenS1{F0: true, F1: v2}
    for fi17 := 0; fi17 < 12; fi17++ {
        v18 := GenS2{F0: 1, F1: -507, F2: v4}
        _ = fi17
        _ = v18
    }
    switch "hello" {
    case "a":
        if (!true) {
            var v19 string = ("a" + ("hello" + ""))
            fmt.Println(v8, v4)
            _ = v19
        } else {
            fmt.Println(v8, v4)
            v2 += genFn0()
            v2 = (v2 >> ((709) & 63))
        }
    default:
        if false {
            v1 = append(v1, -2147483648)
            v2 = (v2 ^ 0)
        } else {
            if (1 > 42) {
                v20 := GenS3{F0: "\n\t"}
                var v21 int = -1
                if false {
                    v21 = (v21 & 2147483647)
                } else {
                    switch (v21 | v21) {
                    case 317:
                        if ((true && (361 >= 31)) && true) {
                            var v22 int32 = (v2 - v2)
                            if (1 == 2) {
                                for fi23 := 0; fi23 < 6; fi23++ {
                                    if len(v3) > 0 {
                                        v3[0] = ("" + "hello")
                                    }
                                    _ = fi23
                                }
                                var v24 float64 = ((float64(1.5) + 0.1) + 0.1)
                                v22 -= (v22 - (v22 / (((v2 + (v22 * v22))) | 1)))
                                _ = v24
                            } else {
                                var v25 int32 = (v22 / ((-2147483648) | 1))
                                for fi26 := 0; fi26 < 10; fi26++ {
                                    switch "foo bar" {
                                    case "":
                                        v27 := GenS2{F0: 255, F1: -9223372036854775808, F2: 1}
                                        _ = v27
                                    case "hello":
                                        for fi28 := 0; fi28 < 8; fi28++ {
                                            var v29 uint = 65535
                                            _ = fi28
                                            _ = v29
                                        }
                                    case "a":
                                        v30 := GenS4{F0: 200}
                                        _ = v30
                                    default:
                                        var v31 int8 = 100
                                        _ = v31
                                    }
                                    v32 := GenS5{F0: (v21 >> ((18446744073709551615) & 63)), F1: 9223372036854775807}
                                    _ = fi26
                                    _ = v32
                                }
                                v33 := GenS4{F0: 128}
                                _ = v25
                                _ = v33
                            }
                            _ = v22
                        }
                    case -559:
                        for fi34 := 0; fi34 < 19; fi34++ {
                            if true {
                                for fi35 := 0; fi35 < 16; fi35++ {
                                    _ = len(v3)
                                    var v36 int8 = 127
                                    if true {
                                        if (false && (false || ((v2 % ((v2) | 1)) >= (v2 - (v2 << ((317) & 63)))))) {
                                            var v37 uint64 = 832
                                            v38 := GenS3{F0: "hello"}
                                            _ = v37
                                            _ = v38
                                        } else {
                                            v39 := GenS3{F0: ""}
                                            _ = v39
                                        }
                                        if (164 != 1) {
                                            v40 := GenS1{F0: true, F1: (v2 % (((v2 % ((v2) | 1))) | 1))}
                                            fmt.Println(v21, v2)
                                            if false {
                                                fmt.Println(fi34, v8)
                                                v4 += (v4 << ((4294967295) & 63))
                                                v4 -= (^(v4))
                                            }
                                            _ = v40
                                        } else {
                                            var v41 int = (-(fi35))
                                            _ = v41
                                        }
                                    } else {
                                        switch (fi35 | int(v4)) {
                                        case 42:
                                            var v42 uint8 = 128
                                            _ = v42
                                        }
                                        if (false && false) {
                                            v8 = (v8 * 1)
                                            var v43 uint8 = 229
                                            if false {
                                                var v44 int64 = 1
                                                var v45 uint8 = (v43 & (v43 % ((v43) | 1)))
                                                switch fi35 {
                                                case 2147483647:
                                                    v36 = -128
                                                default:
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                }
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                _ = v44
                                                _ = v45
                                            } else {
                                                var v46 string = "a"
                                                if (!true) {
                                                    var v47 int8 = int8(v8)
                                                    v3 = append(v3, "")
                                                    var v48 uint = 0
                                                    _ = v47
                                                    _ = v48
                                                } else {
                                                    if (v46 >= v46) {
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        switch int(v4) {
                                                        case 32768:
                                                            var v49 float64 = (float64((1.5 - 1.5)) + -2.25)
                                                            _ = v49
                                                        case 277:
                                                            if (!false) {
                                                                if true {
                                                                    var v50 int = (fi34 - (-(v21)))
                                                                    _ = v50
                                                                }
                                                            }
                                                        case 0:
                                                            var v51 int32 = (v2 >> ((0) & 63))
                                                            _ = v51
                                                        }
                                                        v2 -= v2
                                                    } else {
                                                        var v52 int16 = -32768
                                                        if ((v8 - -4) <= -111) {
                                                            v53 := GenS4{F0: (v43 % ((v43) | 1))}
                                                            v2 = -613
                                                            switch v46 {
                                                            case "foo bar":
                                                                fmt.Println(fi35, v2)
                                                            case "":
                                                                v2 -= 0
                                                            case "hello":
                                                                var v54 uint32 = 2147483647
                                                                _ = v54
                                                            }
                                                            _ = v53
                                                        }
                                                        if (false && false) {
                                                            v43 -= 1
                                                        }
                                                        _ = v52
                                                    }
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    var v55 string = "a"
                                                    v55 = ((v46 + (v46 + "\n\t")) + v46)
                                                    _ = v55
                                                }
                                                _ = v46
                                            }
                                            _ = v43
                                        }
                                        v36 = (v36 >> ((255) & 63))
                                        _ = len(v1)
                                        var v56 uint32 = 1
                                        _ = v56
                                    }
                                    _ = fi35
                                    _ = v36
                                }
                            } else {
                                switch (fi34 % (((v21 | (v21 >> ((4294967295) & 63)))) | 1)) {
                                case 42:
                                    if (-1 < -9223372036854775808) {
                                        v21 += (fi34 | fi34)
                                    } else {
                                        v8 -= 42
                                        for fi57 := 0; fi57 < 17; fi57++ {
                                            v2 += v2
                                            _ = fi57
                                        }
                                    }
                                case -634:
                                    for fi58 := 0; fi58 < 13; fi58++ {
                                        v8 += v8
                                        var v59 uint64 = 700
                                        _ = fi58
                                        _ = v59
                                    }
                                }
                            }
                            var v60 uint16 = 892
                            v21 -= fi34
                            _ = fi34
                            _ = v60
                        }
                    default:
                        v2 = v2
                    }
                    if true {
                        for fi61 := 0; fi61 < 7; fi61++ {
                            var v62 int64 = -1
                            fmt.Println(v2)
                            if (!false) {
                                var v63 int64 = -9223372036854775808
                                v63 = (v63 * 42)
                                var v64 uint16 = 255
                                _ = v63
                                _ = v64
                            } else {
                                switch fi61 {
                                case 32767:
                                    if ((fi61 / ((42) | 1)) > v21) {
                                        v3 = append(v3, "hello")
                                        var v65 int8 = -128
                                        _ = v65
                                    } else {
                                        var v66 int16 = v8
                                        var v67 uint16 = (v4 * 1)
                                        _ = v66
                                        _ = v67
                                    }
                                case -665:
                                    if true {
                                        var v68 uint = 297
                                        for fi69 := 0; fi69 < 13; fi69++ {
                                            v3 = append(v3, ("hello" + "\n\t"))
                                            fmt.Println(fi69, v8, v4)
                                            var v70 int16 = v8
                                            var v71 string = ((("a" + ("\n\t" + ("" + "\n\t"))) + ("hello" + ("\n\t" + "\n\t"))) + "hello")
                                            _ = fi69
                                            _ = v70
                                            _ = v71
                                        }
                                        v2 = 0
                                        _ = v68
                                    }
                                case 1:
                                    switch -9223372036854775808 {
                                    case 42:
                                        v72 := GenS6{F0: (v62 >> ((0) & 63)), F1: true, F2: 1}
                                        _ = v72
                                    case 127:
                                        v62 = v62
                                    case 32768:
                                        v8 += v8
                                    default:
                                        for fi73 := 0; fi73 < 3; fi73++ {
                                            var v74 float32 = (float32(1.5) + (1e10 - -2.25))
                                            v62 += (v62 + -9223372036854775808)
                                            fmt.Println(v2, v4, v2)
                                            fmt.Println(v8, v21, v21)
                                            v4 = v4
                                            v62 -= 42
                                            _ = fi73
                                            _ = v74
                                        }
                                    }
                                default:
                                    var v75 float32 = float32(0.0)
                                    _ = v75
                                }
                                v1 = append(v1, v21)
                                v76 := GenS7{F0: 287, F1: 65535}
                                _ = v76
                            }
                            _ = fi61
                            _ = v62
                        }
                        v21 = (v21 / ((795) | 1))
                    } else {
                        fmt.Println(v2, v21)
                    }
                }
                _ = v20
                _ = v21
            } else {
                fmt.Println(v2, v8, v4)
            }
            for fi77 := 0; fi77 < 18; fi77++ {
                v8 += (^(v8))
                if ((490 <= 417) || false) {
                    if ((v2 ^ (v2 % (((v2 / ((-355) | 1))) | 1))) < (v2 * -2147483648)) {
                        if false {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            if (true && (2 <= 18446744073709551615)) {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v2 = (v2 * v2)
                                fmt.Println(v2)
                            } else {
                                for fi78 := 0; fi78 < 20; fi78++ {
                                    _ = len(v1)
                                    _ = fi78
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                var v79 int = 32767
                                _ = v79
                            }
                        } else {
                            v3 = append(v3, ("hello" + (("a" + (("a" + ("hello" + "foo bar")) + "\n\t")) + (("a" + "a") + "hello"))))
                        }
                    }
                }
                _ = fi77
            }
            for fi80 := 0; fi80 < 20; fi80++ {
                if false {
                    if len(v3) > 0 {
                        v3[0] = ((("a" + "foo bar") + "a") + ("" + "\n\t"))
                    }
                    v8 = 193
                    if ((v4 / (((v4 / ((v4) | 1))) | 1)) < (v4 | v4)) {
                        switch (fi80 & (fi80 * (fi80 - fi80))) {
                        case 127:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        for fi81 := 0; fi81 < 14; fi81++ {
                            v82 := GenS3{F0: "foo bar"}
                            _ = fi81
                            _ = v82
                        }
                        v83 := GenS4{F0: 127}
                        _ = v83
                    } else {
                        switch "foo bar" {
                        case "foo bar":
                            if true {
                                var v84 int = (^(fi80))
                                _ = v84
                            } else {
                                if (genFn0() < (v2 & v2)) {
                                    for fi85 := 0; fi85 < 18; fi85++ {
                                        if (!false) {
                                            if (0 == -74) {
                                                var v86 float64 = -2.25
                                                _ = v86
                                            }
                                            var v87 uint64 = 344
                                            _ = v87
                                        } else {
                                            fmt.Println(v8, v4, fi80)
                                            var v88 int16 = (v8 << ((255) & 63))
                                            var v89 uint = 18446744073709551615
                                            _ = v88
                                            _ = v89
                                        }
                                        fmt.Println(v2)
                                        if (402 < fi85) {
                                            v90 := GenS3{F0: ("hello" + "foo bar")}
                                            var v91 uint32 = 0
                                            _ = v90
                                            _ = v91
                                        } else {
                                            v8 += v8
                                            v4 = v4
                                            var v92 int8 = 42
                                            _ = v92
                                        }
                                        fmt.Println(fi80, v2, v8)
                                        if false {
                                            if false {
                                                if (v4 > (v4 * (v4 | (v4 % ((v4) | 1))))) {
                                                    switch (fi85 & -487) {
                                                    case -2147483648:
                                                        v93 := GenS8{F0: 1, F1: ""}
                                                        _ = v93
                                                    case 127:
                                                        if (!false) {
                                                            if false {
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                            } else {
                                                                switch "hello" {
                                                                case "a":
                                                                    var v94 float32 = 0.1
                                                                    _ = v94
                                                                case "hello":
                                                                    v95 := GenS9{F0: v8, F1: (fi85 + (fi85 >> ((18446744073709551615) & 63))), F2: (fi80 / (((-(fi85))) | 1))}
                                                                    _ = v95
                                                                default:
                                                                    var v96 string = ("" + "foo bar")
                                                                    _ = v96
                                                                }
                                                                v8 = v8
                                                            }
                                                            _ = len(v1)
                                                        }
                                                    default:
                                                        v97 := GenS10{F0: -254, F1: -9223372036854775808}
                                                        _ = v97
                                                    }
                                                } else {
                                                    var v98 uint16 = (v4 & (v4 ^ v4))
                                                    var v99 uint64 = 0
                                                    _ = v98
                                                    _ = v99
                                                }
                                            }
                                            v8 = 0
                                            v4 += 0
                                        }
                                        var v100 float32 = -2.25
                                        _ = fi85
                                        _ = v100
                                    }
                                    switch "\n\t" {
                                    case "\n\t":
                                        fmt.Println(fi80, v2, v4)
                                    case "foo bar":
                                        if (0 > 3) {
                                            v101 := GenS6{F0: 176, F1: true, F2: 1}
                                            v102 := GenS7{F0: (-(v8)), F1: 1}
                                            _ = v101
                                            _ = v102
                                        }
                                    case "a":
                                        fmt.Println(fi80)
                                    }
                                }
                                var v103 int = 4294967295
                                v103 = fi80
                                _ = v103
                            }
                        }
                        v104 := GenS11{F0: (fi80 >> ((33) & 63)), F1: 2, F2: 0}
                        _ = v104
                    }
                }
                switch "" {
                case "a":
                    if (1 <= -1) {
                        v105 := GenS12{F0: v8}
                        _ = v105
                    } else {
                        if true {
                            v2 += (v2 << ((949) & 63))
                        }
                    }
                case "foo bar":
                    var v106 int8 = -110
                    _ = v106
                }
                _ = fi80
            }
        }
    }
    fmt.Println(v4, v4, v8)
    fmt.Println(v2, v4, v8)
    v107 := GenS13{F0: 1}
    switch "a" {
    case "hello":
        v4 = (v4 << ((65535) & 63))
    case "a":
        v108 := GenS6{F0: -1, F1: true, F2: 139}
        _ = v108
    }
    for fi109 := 0; fi109 < 5; fi109++ {
        if false {
            for fi110 := 0; fi110 < 1; fi110++ {
                var v111 float64 = ((3.14159 + (-2.25 - 0.1)) - 3.14159)
                _ = fi110
                _ = v111
            }
            v8 -= (-(v8))
        } else {
            switch 395 {
            case -9223372036854775808:
                if false {
                    for fi112 := 0; fi112 < 1; fi112++ {
                        var v113 uint16 = 32767
                        _ = fi112
                        _ = v113
                    }
                } else {
                    var v114 bool = false
                    _ = v114
                }
            case -2147483648:
                if len(v1) > 0 {
                    v1[0] = fi109
                }
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            for fi115 := 0; fi115 < 16; fi115++ {
                fmt.Println(v2, v8, fi109)
                v116 := GenS14{F0: 0}
                _ = fi115
                _ = v116
            }
        }
        _ = fi109
    }
    var v117 uint64 = 787
    v8 -= 32767
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v117:", v117)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v107
    _ = v117
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
