package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8; F1 int64 }
type GenS5 struct { F0 int8; F1 int; F2 string }
type GenS6 struct { F0 int16; F1 uint64; F2 int8 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int }
func genFn0() int32 {
    if (4294967295 > 2) {
        v4 := GenS1{F0: -2147483648}
        var v5 float32 = 1.5
        _ = v4
        _ = v5
    } else {
        var v6 uint = 65535
        _ = v6
    }
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 uint16 = 65535
        v9 := GenS2{F0: true}
        v8 += (v8 + v8)
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint = 308
    for fi11 := 0; fi11 < 6; fi11++ {
        var v12 uint64 = 4294967295
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 1; fi13++ {
        v14 := GenS1{F0: -1}
        var v15 int8 = 100
        v10 = v10
        _ = fi13
        _ = v14
        _ = v15
    }
    _ = v10
    return 1
}

func main() {
    var v1 int32 = 1
    var v2 uint16 = 1
    v3 := []string{"hello", ""}
    var v16 int8 = 127
    var v17 int32 = v1
    for fi18 := 0; fi18 < 5; fi18++ {
        if (608 <= -9223372036854775808) {
            v19 := GenS3{F0: 2147483647}
            if (!true) {
                var v20 uint64 = 18446744073709551615
                _ = v20
            } else {
                for fi21 := 0; fi21 < 12; fi21++ {
                    v22 := GenS1{F0: v1}
                    _ = fi21
                    _ = v22
                }
                fmt.Println(v16, v2, fi18)
                for fi23 := 0; fi23 < 1; fi23++ {
                    fmt.Println(v16)
                    _ = fi23
                }
            }
            var v24 int64 = 828
            _ = v19
            _ = v24
        } else {
            for fi25 := 0; fi25 < 18; fi25++ {
                v1 = 365
                v26 := GenS4{F0: 255, F1: -388}
                v27 := GenS1{F0: int32(v2)}
                v2 = 65535
                _ = fi25
                _ = v26
                _ = v27
            }
            var v28 float64 = (-2.25 + 3.14159)
            v3 = append(v3, ("a" + "\n\t"))
            _ = v28
        }
        _ = fi18
    }
    for fi29 := 0; fi29 < 5; fi29++ {
        v1 -= v17
        v30 := GenS5{F0: 45, F1: (fi29 % (((fi29 << ((210) & 63))) | 1)), F2: (("foo bar" + ("foo bar" + ("hello" + ("" + ("" + "a"))))) + "a")}
        _ = fi29
        _ = v30
    }
    v16 = int8(v2)
    v3 = append(v3, "a")
    var v31 int32 = v17
    for fi32 := 0; fi32 < 10; fi32++ {
        for fi33 := 0; fi33 < 13; fi33++ {
            if (32768 <= 32767) {
                v3 = append(v3, "a")
                switch 128 {
                case 4294967295:
                    var v34 int64 = -1
                    _ = v34
                case 128:
                    if true {
                        v35 := GenS3{F0: fi33}
                        if (!true) {
                            v1 += v31
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            var v36 uint32 = 2147483647
                            var v37 int64 = 42
                            if (127 > 21) {
                                switch ("" + (("" + ((("hello" + ("foo bar" + "")) + ("" + "")) + ("foo bar" + "hello"))) + "")) {
                                case "\n\t":
                                    var v38 int16 = 42
                                    _ = v38
                                case "":
                                    var v39 string = "\n\t"
                                    _ = v39
                                default:
                                    fmt.Println(v17, fi33, v16)
                                    switch "" {
                                    case "\n\t":
                                        switch fi32 {
                                        case 1:
                                            v37 = (v37 % ((197) | 1))
                                        case 0:
                                            var v40 int8 = (v16 - (^(v16)))
                                            _ = v40
                                        case 609:
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                    case "":
                                        v41 := GenS6{F0: 120, F1: 1, F2: (-(v16))}
                                        _ = v41
                                    default:
                                        var v42 int32 = v1
                                        _ = v42
                                    }
                                    var v43 int8 = 11
                                    _ = v43
                                }
                                _ = len(v3)
                            }
                            _ = v36
                            _ = v37
                        } else {
                            var v44 float64 = ((0.1 - (-2.25 - float64(0.1))) + 1e10)
                            v2 = 32767
                            if (255 > 44) {
                                if false {
                                    v1 -= v31
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                } else {
                                    if (!(592 <= v2)) {
                                        switch (-(fi32)) {
                                        case -75:
                                            switch fi33 {
                                            case 467:
                                                switch "a" {
                                                case "a":
                                                    if true {
                                                        fmt.Println(v1, v16)
                                                        fmt.Println(v1)
                                                    }
                                                case "foo bar":
                                                    fmt.Println(v2, v17)
                                                }
                                            case -9223372036854775808:
                                                if true {
                                                    fmt.Println(fi33, v31, v31)
                                                } else {
                                                    if false {
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        switch -9223372036854775808 {
                                                        case 486:
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        default:
                                                            v45 := GenS3{F0: 32768}
                                                            _ = v45
                                                        }
                                                        var v46 string = "\n\t"
                                                        fmt.Println(v1, fi32, v16)
                                                        _ = v46
                                                    } else {
                                                        v3 = append(v3, "foo bar")
                                                        if len(v3) > 0 {
                                                            v3[0] = ("" + "foo bar")
                                                        }
                                                    }
                                                    v17 -= 2147483647
                                                }
                                            }
                                            fmt.Println(fi32, v2)
                                            fmt.Println(v16)
                                        case 9223372036854775807:
                                            if len(v3) > 0 {
                                                v3[0] = ""
                                            }
                                        case 887:
                                            if (1 <= 4294967295) {
                                                v44 = (1.5 + 1e10)
                                                if false {
                                                    if (v16 == -1) {
                                                        var v47 uint16 = 0
                                                        if (v16 > -128) {
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            var v48 int16 = 411
                                                            _ = v48
                                                        }
                                                        _ = v47
                                                    }
                                                    if ((fi32 % ((0) | 1)) == 9223372036854775807) {
                                                        var v49 string = ("a" + "foo bar")
                                                        _ = v49
                                                    } else {
                                                        if (42 < 1) {
                                                            var v50 int64 = 729
                                                            if len(v3) > 0 {
                                                                v3[0] = "hello"
                                                            }
                                                            var v51 uint64 = 4294967295
                                                            _ = v50
                                                            _ = v51
                                                        }
                                                        if (!false) {
                                                            var v52 int16 = 79
                                                            v53 := GenS3{F0: 2147483647}
                                                            _ = v52
                                                            _ = v53
                                                        }
                                                    }
                                                    switch "\n\t" {
                                                    case "hello":
                                                        v54 := GenS7{F0: v31}
                                                        _ = v54
                                                    case "":
                                                        if true {
                                                            switch -218 {
                                                            case 4294967295:
                                                                if len(v3) > 0 {
                                                                    v3[0] = "a"
                                                                }
                                                            case 127:
                                                                if len(v3) > 0 {
                                                                    v3[0] = (("hello" + "\n\t") + "\n\t")
                                                                }
                                                            }
                                                        } else {
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            v2 = (v2 * (^(v2)))
                                                            if false {
                                                                if false {
                                                                    v55 := GenS4{F0: 2, F1: -9223372036854775808}
                                                                    if len(v3) > 0 {
                                                                        v3[0] = (("a" + "\n\t") + "")
                                                                    }
                                                                    _ = v55
                                                                }
                                                            }
                                                        }
                                                        var v56 int64 = -1
                                                        switch "a" {
                                                        case "foo bar":
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        }
                                                        var v57 int8 = v16
                                                        _ = v56
                                                        _ = v57
                                                    case "\n\t":
                                                        var v58 int8 = (v16 & (v16 % ((v16) | 1)))
                                                        if false {
                                                            if len(v3) > 0 {
                                                                v3[0] = "foo bar"
                                                            }
                                                            switch fi33 {
                                                            case -1:
                                                                v2 = (v2 >> ((2) & 63))
                                                            default:
                                                                var v59 uint64 = 0
                                                                _ = v59
                                                            }
                                                        } else {
                                                            var v60 int = (fi32 >> ((65535) & 63))
                                                            _ = len(v3)
                                                            _ = v60
                                                        }
                                                        fmt.Println(fi32, v58, fi32)
                                                        _ = v58
                                                    }
                                                    var v61 int8 = int8(fi33)
                                                    var v62 int = (fi33 - (^(fi33)))
                                                    if false {
                                                        var v63 uint16 = 0
                                                        v44 -= (float64(1.5) / (float64(1.5) * (v44 * 3.14159)))
                                                        v3 = append(v3, "")
                                                        fmt.Println(v16, v31, fi33)
                                                        v64 := GenS2{F0: false}
                                                        _ = v63
                                                        _ = v64
                                                    } else {
                                                        v1 = v17
                                                    }
                                                    v65 := GenS6{F0: 42, F1: 26, F2: v61}
                                                    _ = v61
                                                    _ = v62
                                                    _ = v65
                                                } else {
                                                    var v66 int8 = 121
                                                    var v67 int16 = -1
                                                    var v68 int64 = -783
                                                    _ = v66
                                                    _ = v67
                                                    _ = v68
                                                }
                                                if true {
                                                    _ = len(v3)
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                } else {
                                                    if true {
                                                        var v69 int16 = 610
                                                        v17 = -183
                                                        if false {
                                                            v44 += (1.5 / 1.0)
                                                            if true {
                                                                v3 = append(v3, "a")
                                                            } else {
                                                                fmt.Println(v44, v2)
                                                                var v70 float64 = (0.1 / float64(1.5))
                                                                _ = v70
                                                            }
                                                        }
                                                        v71 := GenS1{F0: -52}
                                                        _ = v69
                                                        _ = v71
                                                    }
                                                    var v72 uint8 = 1
                                                    if (1000 <= 1000) {
                                                        v1 = v31
                                                        if (true && (-2.25 != ((1.5 * ((((0.0 - -2.25) * float32((v44 - (v44 * v44)))) / (3.14159 + 3.14159)) - 3.14159)) - ((1.5 + 1e10) / (0.1 * (0.1 / (((1.5 + 3.14159) * ((1.5 - (3.14159 / 0.1)) * -2.25)) / 1.5))))))) {
                                                            if true {
                                                                if false {
                                                                    if (!false) {
                                                                        if (!((true || true) && true)) {
                                                                            var v73 uint32 = 4294967295
                                                                            switch ("foo bar" + "") {
                                                                            case "a":
                                                                                var v74 uint = 880
                                                                                v44 += v44
                                                                                _ = v74
                                                                            case "hello":
                                                                                switch 128 {
                                                                                case 413:
                                                                                    var v75 int16 = 1000
                                                                                    fmt.Println(v1)
                                                                                    var v76 float64 = 0.1
                                                                                    if (uint64(v72) <= 4294967295) {
                                                                                        var v77 bool = true
                                                                                        var v78 uint16 = (^(v2))
                                                                                        _ = v77
                                                                                        _ = v78
                                                                                    }
                                                                                    _ = len(v3)
                                                                                    var v79 uint8 = v72
                                                                                    _ = v75
                                                                                    _ = v76
                                                                                    _ = v79
                                                                                case -9223372036854775808:
                                                                                    if false {
                                                                                        v73 = v73
                                                                                        var v80 bool = (true || false)
                                                                                        v80 = (0 == (^(v73)))
                                                                                        _ = v80
                                                                                    } else {
                                                                                        switch (fi33 << ((530) & 63)) {
                                                                                        case 4294967295:
                                                                                            fmt.Println(v2, fi32)
                                                                                        case 0:
                                                                                            v1 = (v31 * v1)
                                                                                        case 1:
                                                                                            switch (fi33 + (fi33 % ((fi32) | 1))) {
                                                                                            case 231:
                                                                                                switch "foo bar" {
                                                                                                case "hello":
                                                                                                    var v81 uint = 992
                                                                                                    _ = v81
                                                                                                case "\n\t":
                                                                                                    fmt.Println(v73, v16, v31)
                                                                                                }
                                                                                            default:
                                                                                                if false {
                                                                                                    v44 = 1e10
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        var v82 string = "foo bar"
                                                                                        var v83 int64 = 0
                                                                                        v84 := GenS1{F0: v1}
                                                                                        _ = v82
                                                                                        _ = v83
                                                                                        _ = v84
                                                                                    }
                                                                                case 2147483647:
                                                                                    v1 -= v1
                                                                                }
                                                                            case "foo bar":
                                                                                if (!(1 > 1)) {
                                                                                    var v85 string = ((("hello" + "a") + "foo bar") + "hello")
                                                                                    v16 -= 127
                                                                                    _ = v85
                                                                                }
                                                                            default:
                                                                                if (206 <= 804) {
                                                                                    v86 := GenS5{F0: v16, F1: fi32, F2: "foo bar"}
                                                                                    _ = v86
                                                                                }
                                                                            }
                                                                            if (4294967295 == 4294967295) {
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                            } else {
                                                                                v87 := GenS8{F0: 795}
                                                                                v16 = 1
                                                                                fmt.Println(v17)
                                                                                _ = v87
                                                                            }
                                                                            _ = v73
                                                                        }
                                                                        v31 += int32(v16)
                                                                        v3 = append(v3, "foo bar")
                                                                    } else {
                                                                        v88 := GenS1{F0: v17}
                                                                        var v89 string = "a"
                                                                        _ = v88
                                                                        _ = v89
                                                                    }
                                                                    v72 += v72
                                                                    v16 += (^(v16))
                                                                }
                                                                if len(v3) > 0 {
                                                                    v3[0] = "foo bar"
                                                                }
                                                                v16 = (v16 ^ v16)
                                                                switch fi33 {
                                                                case 128:
                                                                    v90 := GenS9{F0: (^(v2))}
                                                                    _ = v90
                                                                case -589:
                                                                    fmt.Println(v16, v44)
                                                                case 42:
                                                                    var v91 uint32 = 1
                                                                    v31 = (v17 >> ((18446744073709551615) & 63))
                                                                    v91 -= 42
                                                                    v3 = append(v3, "foo bar")
                                                                    _ = v91
                                                                }
                                                            } else {
                                                                v17 = int32(fi33)
                                                                v44 += (v44 / 1e10)
                                                                switch (("\n\t" + ("foo bar" + ("foo bar" + ("a" + ("foo bar" + "hello"))))) + "a") {
                                                                case "\n\t":
                                                                    v3 = append(v3, "foo bar")
                                                                case "hello":
                                                                    v92 := GenS2{F0: false}
                                                                    _ = v92
                                                                default:
                                                                    if (!false) {
                                                                        v31 -= 709
                                                                        var v93 int32 = genFn0()
                                                                        if len(v3) > 0 {
                                                                            v3[0] = "hello"
                                                                        }
                                                                        _ = v93
                                                                    } else {
                                                                        if (v1 > v31) {
                                                                            var v94 uint32 = 2147483647
                                                                            _ = len(v3)
                                                                            v94 += v94
                                                                            _ = v94
                                                                        } else {
                                                                            if ((v44 + v44) > float64(float32(v44))) {
                                                                                var v95 int64 = int64(v72)
                                                                                fmt.Println(v31, v95)
                                                                                _ = v95
                                                                            }
                                                                            if false {
                                                                                v31 -= 425
                                                                                if ((!true) || (2 >= 4294967295)) {
                                                                                    fmt.Println(v44, v44, v31)
                                                                                }
                                                                                fmt.Println(fi32)
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        var v96 float64 = -2.25
                                                        _ = v96
                                                    }
                                                    var v97 string = ""
                                                    var v98 bool = (true && (fi33 <= (fi32 + fi33)))
                                                    _ = v72
                                                    _ = v97
                                                    _ = v98
                                                }
                                            }
                                        default:
                                            v99 := GenS10{F0: int64(v2)}
                                            _ = v99
                                        }
                                    }
                                }
                            }
                            fmt.Println(v16, v31, fi32)
                            v3 = append(v3, ("\n\t" + (("hello" + "a") + ("" + ""))))
                            _ = v44
                        }
                        _ = v35
                    }
                case 646:
                    var v100 int64 = 1
                    _ = v100
                default:
                    v17 = v17
                }
                var v101 uint16 = (v2 >> ((18446744073709551615) & 63))
                _ = v101
            }
            _ = fi33
        }
        var v102 int32 = v31
        var v103 int8 = (v16 / ((-91) | 1))
        _ = fi32
        _ = v102
        _ = v103
    }
    if false {
        if false {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            switch -361 {
            case 2147483647:
                var v104 uint8 = 127
                _ = v104
            }
        } else {
            v3 = append(v3, "")
            v105 := GenS7{F0: (v17 % (((v31 % ((v1) | 1))) | 1))}
            _ = v105
        }
        v106 := GenS1{F0: v31}
        v3 = append(v3, ("\n\t" + "hello"))
        _ = v106
    } else {
        for fi107 := 0; fi107 < 13; fi107++ {
            var v108 int16 = 1000
            _ = fi107
            _ = v108
        }
        v17 = (v31 - v17)
        fmt.Println(v31, v17)
    }
    if false {
        fmt.Println(v31, v1, v16)
    } else {
        if false {
            var v109 int16 = int16(v31)
            v1 = genFn0()
            _ = v109
        } else {
            for fi110 := 0; fi110 < 6; fi110++ {
                for fi111 := 0; fi111 < 17; fi111++ {
                    v2 -= v2
                    _ = len(v3)
                    v17 += (^(v1))
                    if false {
                        if (2 > 2147483647) {
                            var v112 uint = 0
                            _ = v112
                        } else {
                            v1 = genFn0()
                            v113 := GenS7{F0: (v17 * v31)}
                            _ = v113
                        }
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if (true || true) {
                        v3 = append(v3, "hello")
                        v114 := GenS11{F0: 665}
                        _ = v114
                    }
                    _ = fi111
                }
                _ = fi110
            }
            switch "hello" {
            case "\n\t":
                v115 := GenS8{F0: 784}
                _ = v115
            case "a":
                if false {
                    for fi116 := 0; fi116 < 5; fi116++ {
                        if true {
                            v1 -= 42
                            var v117 int64 = 42
                            _ = v117
                        }
                        fmt.Println(v31, v31)
                        _ = len(v3)
                        _ = fi116
                    }
                }
            }
        }
    }
    for fi118 := 0; fi118 < 8; fi118++ {
        if false {
            if ("foo bar" < "\n\t") {
                v119 := GenS9{F0: 255}
                for fi120 := 0; fi120 < 20; fi120++ {
                    var v121 uint16 = v2
                    var v122 int16 = 0
                    v123 := GenS10{F0: -1}
                    _ = fi120
                    _ = v121
                    _ = v122
                    _ = v123
                }
                var v124 int32 = (^(v1))
                _ = v119
                _ = v124
            }
        }
        var v125 uint8 = 127
        _ = fi118
        _ = v125
    }
    if len(v3) > 0 {
        v3[0] = ("\n\t" + "")
    }
    for fi126 := 0; fi126 < 1; fi126++ {
        for fi127 := 0; fi127 < 16; fi127++ {
            var v128 int64 = -9223372036854775808
            fmt.Println(v128)
            var v129 int64 = 0
            fmt.Println(v129, v31)
            _ = fi127
            _ = v128
            _ = v129
        }
        v130 := GenS12{F0: -786}
        _ = fi126
        _ = v130
    }
    v31 -= 220
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v31
}
