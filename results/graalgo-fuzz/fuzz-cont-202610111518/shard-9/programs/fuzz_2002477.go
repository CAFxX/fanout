package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 string; F1 uint32; F2 string }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint64; F1 string; F2 int32 }
func genFn0(p0 uint8, p1 uint64) int8 {
    for fi4 := 0; fi4 < 2; fi4++ {
        var v5 uint8 = (p0 & (p0 % ((p0) | 1)))
        v5 -= 30
        _ = fi4
        _ = v5
    }
    if (127 != 1) {
        for fi6 := 0; fi6 < 16; fi6++ {
            for fi7 := 0; fi7 < 3; fi7++ {
                v8 := []int{fi6, fi6, fi6, fi6}
                _ = fi7
                _ = v8
            }
            if true {
                var v9 int = (fi6 + (fi6 + 42))
                _ = v9
            } else {
                var v10 int16 = 42
                for fi11 := 0; fi11 < 18; fi11++ {
                    p0 -= uint8(p1)
                    p0 = (p0 + 0)
                    _ = fi11
                }
                _ = v10
            }
            _ = fi6
        }
    }
    _ = p0
    _ = p1
    return 1
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 32767
    v3 := []uint32{42, 1, 272}
    _ = len(v1)
    fmt.Println(v2)
    var v12 int32 = -65
    v13 := v1["hello"]
    if true {
        var v14 float32 = 1e10
        _ = v14
    }
    var v15 string = "a"
    if true {
        switch "a" {
        case "foo bar":
            for fi16 := 0; fi16 < 1; fi16++ {
                var v17 int = fi16
                if (2147483647 != v12) {
                    var v18 uint64 = 564
                    for fi19 := 0; fi19 < 1; fi19++ {
                        fmt.Println(fi19, v15, v12)
                        _ = fi19
                    }
                    _ = v18
                }
                for fi20 := 0; fi20 < 3; fi20++ {
                    if (924 != 876) {
                        v21 := v1[("a" + v15)]
                        _ = v21
                    } else {
                        v15 = ""
                    }
                    switch (v15 + "") {
                    case "foo bar":
                        fmt.Println(v12, v2)
                    case "hello":
                        fmt.Println(v12)
                        v22 := GenS1{F0: 195, F1: 18446744073709551615}
                        _ = v22
                    case "\n\t":
                        delete(v1, v15)
                    default:
                        var v23 uint64 = 2
                        switch v15 {
                        case "":
                            v17 += 202
                        }
                        v13 -= (v13 - v13)
                        var v24 int8 = 127
                        if (((v2 % (((v2 >> ((0) & 63))) | 1)) >= v2) || true) {
                            v25 := v1["\n\t"]
                            if ("hello" <= "hello") {
                                var v26 int32 = (v12 ^ v12)
                                _ = len(v3)
                                _ = v26
                            } else {
                                if (!false) {
                                    var v27 string = "\n\t"
                                    if true {
                                        v27 = ("a" + v15)
                                    } else {
                                        var v28 bool = false
                                        if len(v3) > 0 {
                                            v3[0] = (v25 & 4294967295)
                                        }
                                        v29 := GenS1{F0: (v2 + v2), F1: 418}
                                        _ = v28
                                        _ = v29
                                    }
                                    _ = v27
                                }
                                switch ((v15 + (v15 + v15)) + ("a" + ("" + v15))) {
                                case "a":
                                    if false {
                                        _ = len(v1)
                                        var v30 uint8 = 1
                                        if ((!false) && ((1.5 / (1.5 * (0.1 - -2.25))) <= -2.25)) {
                                            switch fi16 {
                                            case -2147483648:
                                                v13 = 235
                                            case -275:
                                                switch (^(fi16)) {
                                                case -2147483648:
                                                    fmt.Println(v15, v30)
                                                    switch 125 {
                                                    case 128:
                                                        var v31 int64 = 1
                                                        _ = v31
                                                    case 42:
                                                        var v32 int16 = -568
                                                        _ = v32
                                                    case 902:
                                                        if len(v3) > 0 {
                                                            v3[0] = v13
                                                        }
                                                    }
                                                    var v33 float64 = (0.0 / 1.0)
                                                    _ = v33
                                                }
                                                v1[v15] = 3
                                            case -1:
                                                if false {
                                                    var v34 uint64 = v23
                                                    _ = v34
                                                }
                                            default:
                                                var v35 int16 = 180
                                                _ = v35
                                            }
                                        }
                                        _ = v30
                                    } else {
                                        v36 := GenS1{F0: v2, F1: 255}
                                        if false {
                                            if false {
                                                v1[v15] = (v13 / (((v25 + (^(v25)))) | 1))
                                            } else {
                                                v13 -= v13
                                                var v37 int16 = 32767
                                                v37 = 0
                                                _ = v37
                                            }
                                            v23 += v23
                                        }
                                        _ = len(v1)
                                        if (true || true) {
                                            if false {
                                                fmt.Println(v15, fi16)
                                                v2 = 2
                                                switch fi16 {
                                                case 600:
                                                    fmt.Println(v23)
                                                case 1:
                                                    switch int(v12) {
                                                    case 0:
                                                        fmt.Println(v23, fi20)
                                                    }
                                                }
                                            } else {
                                                switch -270 {
                                                case 32768:
                                                    var v38 uint8 = 0
                                                    _ = v38
                                                case -1:
                                                    v39 := GenS2{F0: (v2 >> ((18446744073709551615) & 63)), F1: 127}
                                                    _ = v39
                                                case 33:
                                                    var v40 int = v17
                                                    _ = v40
                                                }
                                            }
                                            if (!false) {
                                                var v41 int32 = (-(v12))
                                                _ = len(v1)
                                                var v42 bool = ((false && false) || ((v12 != int32(v25)) && (0 == -1)))
                                                var v43 int8 = (-(v24))
                                                _ = v41
                                                _ = v42
                                                _ = v43
                                            }
                                        } else {
                                            var v44 uint8 = 255
                                            switch fi20 {
                                            case -298:
                                                var v45 int64 = 82
                                                _ = v45
                                            case -2147483648:
                                                v46 := GenS3{F0: v15, F1: (v25 + v13), F2: (v15 + "\n\t")}
                                                if (!true) {
                                                    if (9223372036854775807 < -451) {
                                                        _ = len(v1)
                                                        var v47 uint64 = 751
                                                        v3 = append(v3, 122)
                                                        _ = v47
                                                    } else {
                                                        var v48 int64 = 463
                                                        if false {
                                                            var v49 uint8 = v44
                                                            _ = v49
                                                        }
                                                        if (false || false) {
                                                            v23 -= 599
                                                            fmt.Println(v12, v24)
                                                        } else {
                                                            var v50 float64 = 1e10
                                                            _ = v50
                                                        }
                                                        v51 := GenS2{F0: (v2 ^ 0), F1: 0}
                                                        delete(v1, (v15 + v15))
                                                        v23 += (v23 * v23)
                                                        _ = len(v3)
                                                        _ = v48
                                                        _ = v51
                                                    }
                                                } else {
                                                    v12 = (v12 * v12)
                                                    fmt.Println(v13, v13)
                                                    v44 = v44
                                                }
                                                v15 = (("hello" + "hello") + (v15 + (v15 + "a")))
                                                _ = v46
                                            default:
                                                if true {
                                                    var v52 int64 = -9223372036854775808
                                                    switch "" {
                                                    case "hello":
                                                        v24 = (-(v24))
                                                    case "":
                                                        delete(v1, v15)
                                                    case "a":
                                                        if len(v3) > 0 {
                                                            v3[0] = (v13 << ((0) & 63))
                                                        }
                                                    default:
                                                        v53 := GenS2{F0: uint16(v44), F1: (v44 - (v44 << ((2) & 63)))}
                                                        _ = v53
                                                    }
                                                    _ = v52
                                                } else {
                                                    _ = len(v1)
                                                    v12 = 0
                                                    v23 = uint64(v25)
                                                    fmt.Println(fi16)
                                                }
                                            }
                                            v15 = (("" + (v15 + (v15 + (v15 + v15)))) + v15)
                                            _ = v44
                                        }
                                        if (!false) {
                                            v15 = v15
                                            var v54 string = ("hello" + "\n\t")
                                            _ = v54
                                        } else {
                                            v55 := GenS1{F0: (v2 >> ((2) & 63)), F1: 42}
                                            fmt.Println(fi20)
                                            _ = v55
                                        }
                                        switch (fi20 / ((fi20) | 1)) {
                                        case -9223372036854775808:
                                            var v56 uint = 42
                                            _ = v56
                                        case 0:
                                            v23 = uint64(v2)
                                        case 127:
                                            fmt.Println(v15, v23, v17)
                                        default:
                                            v12 -= (v12 << ((501) & 63))
                                        }
                                        _ = v36
                                    }
                                case "":
                                    v57 := v1[(("foo bar" + ("foo bar" + v15)) + v15)]
                                    _ = v57
                                case "foo bar":
                                    if false {
                                        v23 -= v23
                                        v12 += v12
                                        if len(v3) > 0 {
                                            v3[0] = v13
                                        }
                                        var v58 int8 = (v24 & int8(v13))
                                        var v59 int = (fi20 >> ((7) & 63))
                                        _ = v58
                                        _ = v59
                                    } else {
                                        v60 := GenS2{F0: 5, F1: 52}
                                        v13 = 2147483648
                                        _ = v60
                                    }
                                }
                            }
                            _ = v25
                        }
                        switch v15 {
                        case "\n\t":
                            if false {
                                switch v15 {
                                case "foo bar":
                                    v13 -= 0
                                case "":
                                    var v61 int32 = (v12 & (v12 | v12))
                                    _ = v61
                                }
                                var v62 uint = 1
                                var v63 string = v15
                                _ = v62
                                _ = v63
                            }
                        }
                        var v64 uint = 2
                        v65 := GenS4{F0: v13}
                        var v66 uint64 = 820
                        _ = v23
                        _ = v24
                        _ = v64
                        _ = v65
                        _ = v66
                    }
                    _ = fi20
                }
                _ = fi16
                _ = v17
            }
        default:
            fmt.Println(v2, v15)
        }
        v67 := GenS2{F0: 255, F1: 127}
        switch 42 {
        case 32767:
            v68 := v1[(("foo bar" + ((("foo bar" + ("a" + (v15 + (v15 + v15)))) + "hello") + ((("hello" + ("a" + v15)) + v15) + "\n\t"))) + "a")]
            _ = v68
        case -602:
            if true {
                var v69 int32 = (^(v12))
                v70 := v1[((v15 + v15) + v15)]
                var v71 bool = false
                _ = v69
                _ = v70
                _ = v71
            }
        default:
            var v72 float64 = (1.5 + 1.5)
            _ = v72
        }
        _ = v67
    } else {
        var v73 int8 = genFn0(uint8(v13), 2)
        _ = v73
    }
    var v74 string = v15
    var v75 uint = 255
    if true {
        for fi76 := 0; fi76 < 16; fi76++ {
            if (!(false && false)) {
                if ((4294967295 != (v13 | v13)) || (!(1e10 < ((0.1 + 3.14159) + float64(0.1))))) {
                    var v77 int8 = 0
                    v12 += (v12 - (v12 ^ int32(v75)))
                    _ = v77
                } else {
                    for fi78 := 0; fi78 < 13; fi78++ {
                        v79 := GenS3{F0: v15, F1: 863, F2: ("" + v74)}
                        _ = fi78
                        _ = v79
                    }
                }
                if false {
                    if true {
                        v80 := GenS2{F0: 608, F1: 128}
                        _ = v80
                    }
                    switch v74 {
                    case "hello":
                        var v81 int64 = 0
                        _ = v81
                    case "\n\t":
                        for fi82 := 0; fi82 < 17; fi82++ {
                            fmt.Println(v13)
                            v74 = v74
                            _ = fi82
                        }
                    default:
                        v83 := GenS3{F0: (v74 + v15), F1: 11, F2: "\n\t"}
                        _ = v83
                    }
                    v84 := GenS3{F0: (("foo bar" + "a") + v15), F1: 42, F2: ("a" + (("foo bar" + ("foo bar" + v74)) + ((v15 + (v15 + (v15 + "foo bar"))) + (v74 + (v15 + (((("a" + (v15 + "")) + (v15 + (v74 + v15))) + "a") + "foo bar"))))))}
                    _ = v84
                } else {
                    v85 := GenS1{F0: v2, F1: 65535}
                    _ = v85
                }
                if false {
                    switch v74 {
                    case "":
                        v86 := GenS3{F0: v15, F1: 1, F2: v74}
                        _ = v86
                    case "hello":
                        var v87 uint = 255
                        _ = v87
                    }
                }
            } else {
                var v88 int32 = v12
                v13 -= v13
                _ = v88
            }
            v89 := GenS3{F0: v15, F1: v13, F2: v74}
            _ = fi76
            _ = v89
        }
    }
    if (0 == (v75 | (v75 - (v75 << ((v75) & 63))))) {
        switch "hello" {
        case "hello":
            for fi90 := 0; fi90 < 13; fi90++ {
                _ = len(v3)
                _ = fi90
            }
        case "foo bar":
            var v91 float64 = ((3.14159 / 1.0) * float64((1.5 - (0.0 / 3.14159))))
            _ = v91
        case "\n\t":
            var v92 float64 = 0.0
            _ = v92
        default:
            v13 = 254
        }
        v93 := GenS3{F0: "a", F1: (v13 >> (((v75 * (^(v75)))) & 63)), F2: ((v74 + "\n\t") + (("hello" + "") + "\n\t"))}
        if (v12 == v12) {
            fmt.Println(v13, v13, v75)
            if ((0 > 128) || false) {
                var v94 bool = (1e10 <= -2.25)
                _ = v94
            } else {
                var v95 bool = true
                fmt.Println(v15, v2, v12)
                _ = v95
            }
        } else {
            if (!true) {
                for fi96 := 0; fi96 < 6; fi96++ {
                    switch int(v75) {
                    case 32768:
                        v97 := GenS5{F0: ("" + "a")}
                        _ = v97
                    case 773:
                        var v98 uint8 = 2
                        _ = v98
                    default:
                        switch v15 {
                        case "":
                            for fi99 := 0; fi99 < 4; fi99++ {
                                _ = len(v1)
                                fmt.Println(v15)
                                if false {
                                    var v100 bool = (v75 > 42)
                                    v13 = v13
                                    fmt.Println(fi96, v74, v12)
                                    var v101 int32 = v12
                                    _ = v100
                                    _ = v101
                                } else {
                                    fmt.Println(v15, v75, v74)
                                    fmt.Println(fi99, v75)
                                }
                                _ = fi99
                            }
                        case "foo bar":
                            fmt.Println(v74, v15)
                        default:
                            if true {
                                var v102 uint = 835
                                _ = v102
                            } else {
                                switch fi96 {
                                case 127:
                                    var v103 int = (fi96 << ((v75) & 63))
                                    _ = v103
                                case 32767:
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                case 2147483647:
                                    if true {
                                        v3 = append(v3, 2)
                                        v2 = v2
                                    } else {
                                        v104 := v1[""]
                                        v15 = v74
                                        _ = v104
                                    }
                                }
                                var v105 uint16 = v2
                                fmt.Println(fi96, v74)
                                _ = v105
                            }
                        }
                    }
                    var v106 int32 = 1
                    for fi107 := 0; fi107 < 19; fi107++ {
                        switch fi96 {
                        case 0:
                            if true {
                                if true {
                                    var v108 uint8 = 128
                                    var v109 uint64 = 218
                                    _ = v108
                                    _ = v109
                                }
                                if ((false && true) && true) {
                                    v74 = ("a" + v15)
                                    v110 := GenS6{F0: 1, F1: (v74 + v15), F2: v12}
                                    if (-622 >= 9223372036854775807) {
                                        var v111 uint64 = 18446744073709551615
                                        _ = v111
                                    }
                                    _ = v110
                                }
                            } else {
                                v74 = (v15 + "")
                            }
                        case 42:
                            var v112 int = 888
                            _ = v112
                        case 1:
                            v1[v74] = (v13 & v13)
                        default:
                            v15 = (("a" + ("hello" + v74)) + (("foo bar" + ("a" + "foo bar")) + ""))
                            var v113 float32 = 1.5
                            fmt.Println(v75, fi96, v106)
                            _ = v113
                        }
                        v1["\n\t"] = (v13 | v13)
                        v74 = ((v74 + "hello") + v15)
                        _ = fi107
                    }
                    _ = fi96
                    _ = v106
                }
                v75 -= (^(v75))
                if true {
                    var v114 int32 = v12
                    if true {
                        _ = len(v1)
                        for fi115 := 0; fi115 < 10; fi115++ {
                            v13 = 65535
                            v74 = "hello"
                            v2 = v2
                            _ = fi115
                        }
                        var v116 uint32 = (v13 ^ (v13 << (((v75 ^ 2)) & 63)))
                        _ = v116
                    }
                    _ = v114
                }
            }
        }
        _ = v93
    } else {
        var v117 uint = v75
        switch 32767 {
        case -327:
            v74 = ""
        case 32768:
            switch v15 {
            case "":
                v2 = (v2 & 42)
            case "\n\t":
                v15 = "a"
            }
        case -97:
            var v118 uint64 = 18446744073709551615
            _ = v118
        }
        _ = v117
    }
    var v119 uint8 = 255
    var v120 string = v74
    v120 = v120
    v1[("\n\t" + "foo bar")] = (v13 * 2)
    _ = len(v1)
    fmt.Println(v75, v74)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v119:", v119)
    fmt.Println("v12:", v12)
    fmt.Println("v120:", v120)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v119
    _ = v12
    _ = v120
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v74
    _ = v75
}
