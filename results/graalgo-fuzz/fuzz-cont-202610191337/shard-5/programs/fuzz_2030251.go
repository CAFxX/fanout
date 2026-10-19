package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 uint8; F2 uint }
type GenS4 struct { F0 int16; F1 int32 }
type GenS5 struct { F0 int64; F1 int32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 string; F1 int32; F2 int }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int; F1 int; F2 uint }
type GenS10 struct { F0 uint32; F1 string; F2 int }
type GenS11 struct { F0 int }

func main() {
    var v1 uint = 65535
    var v2 uint16 = 953
    var v3 uint16 = (v2 * (v2 % (((v2 >> ((v1) & 63))) | 1)))
    v4 := []uint32{1, 42}
    v5 := make(map[string]string)
    switch -820 {
    case -9223372036854775808:
        var v6 int64 = 958
        _ = v6
    case 128:
        for fi7 := 0; fi7 < 9; fi7++ {
            var v8 int16 = -1
            _ = fi7
            _ = v8
        }
    case 766:
        fmt.Println(v1, v2)
    }
    switch "" {
    case "hello":
        if false {
            if true {
                var v9 uint = 4294967295
                _ = v9
            }
            if len(v4) > 0 {
                v4[0] = 4294967295
            }
            if true {
                if false {
                    if (!false) {
                        v4 = append(v4, 2)
                    } else {
                        for fi10 := 0; fi10 < 15; fi10++ {
                            v11 := GenS1{F0: -492}
                            _ = fi10
                            _ = v11
                        }
                        v12 := GenS1{F0: 1}
                        _ = v12
                    }
                } else {
                    v2 = 1
                    var v13 int32 = 0
                    v3 = 175
                    _ = v13
                }
                v14 := GenS1{F0: 0}
                switch 2147483647 {
                case 385:
                    var v15 int = 0
                    _ = v15
                case 127:
                    var v16 float64 = ((-2.25 - (0.0 * -2.25)) - 1e10)
                    _ = v16
                case -468:
                    v17 := GenS1{F0: -1}
                    _ = v17
                }
                _ = v14
            } else {
                fmt.Println(v3, v3, v1)
                if false {
                    if true {
                        var v18 int64 = 42
                        _ = v18
                    } else {
                        v1 = v1
                        v19 := GenS2{F0: true}
                        _ = v19
                    }
                    v20 := GenS1{F0: -927}
                    _ = v20
                } else {
                    switch 9223372036854775807 {
                    case -289:
                        var v21 string = "foo bar"
                        _ = v21
                    case 1:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    case -1:
                        v5[("hello" + (("hello" + "hello") + ""))] = "a"
                    default:
                        var v22 int32 = int32(v2)
                        _ = v22
                    }
                    v2 += 169
                    v3 -= (^(v3))
                }
            }
        } else {
            var v23 uint16 = (v2 ^ uint16(v1))
            _ = v23
        }
    case "\n\t":
        v2 += 74
    case "a":
        if (0.1 != 3.14159) {
            for fi24 := 0; fi24 < 19; fi24++ {
                v3 += (v2 & v3)
                var v25 float32 = float32((1.5 - 1e10))
                for fi26 := 0; fi26 < 3; fi26++ {
                    v1 = v1
                    switch fi24 {
                    case 32768:
                        var v27 uint16 = uint16(v1)
                        _ = v27
                    case 126:
                        v28 := GenS3{F0: 2, F1: 42, F2: 2}
                        _ = v28
                    case 4294967295:
                        v25 += v25
                    }
                    _ = fi26
                }
                _ = fi24
                _ = v25
            }
        } else {
            var v29 int64 = -1
            v30 := GenS2{F0: ((2147483647 == 2147483647) && true)}
            _ = v29
            _ = v30
        }
    default:
        if false {
            if false {
                v1 += (v1 % (((v1 | (v1 / ((v1) | 1)))) | 1))
                if (!((42 >= 42) && true)) {
                    v5["hello"] = "\n\t"
                    if ((!true) && (0.1 != ((0.0 / -2.25) * 0.0))) {
                        v31 := GenS2{F0: (!true)}
                        _ = v31
                    }
                } else {
                    for fi32 := 0; fi32 < 20; fi32++ {
                        fmt.Println(v3, v1, v2)
                        v33 := GenS3{F0: 127, F1: 255, F2: 42}
                        switch (^(fi32)) {
                        case 127:
                            for fi34 := 0; fi34 < 13; fi34++ {
                                fmt.Println(v3, v3)
                                if false {
                                    if true {
                                        var v35 int64 = 1
                                        switch (fi32 % (((fi34 / ((fi32) | 1))) | 1)) {
                                        case -9223372036854775808:
                                            var v36 uint64 = 1
                                            _ = v36
                                        case 128:
                                            switch ("" + (("\n\t" + "hello") + ((((("" + "a") + "") + "foo bar") + "") + (("hello" + ("foo bar" + "foo bar")) + "\n\t")))) {
                                            case "foo bar":
                                                v37 := GenS4{F0: -1, F1: 42}
                                                _ = v37
                                            }
                                        case 32768:
                                            var v38 bool = ((3.14159 <= (-2.25 * 3.14159)) && ((v35 & v35) >= (-(v35))))
                                            _ = v38
                                        }
                                        var v39 string = "\n\t"
                                        _ = v35
                                        _ = v39
                                    }
                                    v1 = (v1 & 415)
                                    v40 := GenS3{F0: 216, F1: 255, F2: v1}
                                    _ = v40
                                }
                                _ = fi34
                            }
                        case 9223372036854775807:
                            v41 := GenS5{F0: 1, F1: 0}
                            _ = v41
                        }
                        _ = fi32
                        _ = v33
                    }
                    switch "foo bar" {
                    case "\n\t":
                        var v42 bool = false
                        _ = v42
                    case "a":
                        for fi43 := 0; fi43 < 4; fi43++ {
                            var v44 int8 = -78
                            _ = fi43
                            _ = v44
                        }
                    case "foo bar":
                        fmt.Println(v3)
                    }
                    var v45 int = 42
                    _ = v45
                }
            } else {
                var v46 string = (("\n\t" + (("" + ((("" + "hello") + "foo bar") + ("hello" + "hello"))) + "\n\t")) + "\n\t")
                _ = v46
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v47 float32 = 1e10
            _ = v47
        }
    }
    v1 += v1
    fmt.Println(v2, v2)
    fmt.Println(v2, v3, v1)
    v48 := v5[(("hello" + ("\n\t" + ("\n\t" + "a"))) + "a")]
    v1 = v1
    fmt.Println(v2)
    var v49 uint = (^(v1))
    if len(v4) > 0 {
        v4[0] = 65535
    }
    switch (("a" + v48) + "\n\t") {
    case "a":
        for fi50 := 0; fi50 < 11; fi50++ {
            fmt.Println(fi50, fi50, v1)
            _ = fi50
        }
    case "\n\t":
        var v51 uint32 = 166
        _ = v51
    case "hello":
        if ((v48 >= v48) || (!false)) {
            if false {
                fmt.Println(v2)
                v3 -= v3
                v1 += (v1 & 781)
            }
            v52 := v5[("hello" + "\n\t")]
            if (-1 >= 804) {
                for fi53 := 0; fi53 < 9; fi53++ {
                    if true {
                        var v54 uint16 = v3
                        var v55 int64 = 1
                        v52 = ("a" + v48)
                        _ = v54
                        _ = v55
                    } else {
                        v2 -= v2
                        fmt.Println(v1)
                    }
                    if (1 == -159) {
                        var v56 int16 = -16
                        _ = v56
                    } else {
                        var v57 uint64 = 420
                        switch int(v57) {
                        case -9223372036854775808:
                            switch (fi53 * 772) {
                            case 32768:
                                for fi58 := 0; fi58 < 19; fi58++ {
                                    v1 = (v1 - (v1 + (v49 << (((v1 % ((v1) | 1))) & 63))))
                                    var v59 uint64 = (v57 + v57)
                                    _ = fi58
                                    _ = v59
                                }
                            case 4294967295:
                                v60 := GenS1{F0: 1}
                                _ = v60
                            case 127:
                                var v61 float32 = -2.25
                                _ = v61
                            }
                        case 659:
                            v49 -= 0
                        case 32768:
                            var v62 int64 = 0
                            _ = v62
                        default:
                            var v63 string = "\n\t"
                            _ = v63
                        }
                        v64 := GenS1{F0: 1000}
                        _ = v57
                        _ = v64
                    }
                    var v65 int16 = 579
                    _ = fi53
                    _ = v65
                }
            }
            _ = v52
        }
    }
    v66 := GenS6{F0: 100}
    switch -1 {
    case 2147483647:
        v2 = 0
    default:
        v2 = 603
    }
    var v67 int = -832
    for fi68 := 0; fi68 < 18; fi68++ {
        v2 += (v3 * (v3 * (v3 >> (((v49 % ((v49) | 1))) & 63))))
        var v69 int8 = 127
        _ = fi68
        _ = v69
    }
    switch (v48 + "hello") {
    case "foo bar":
        var v70 float64 = float64(0.0)
        _ = v70
    case "\n\t":
        v71 := GenS7{F0: v48, F1: 392, F2: (v67 % ((-52) | 1))}
        _ = v71
    case "a":
        var v72 float32 = -2.25
        _ = v72
    }
    var v73 string = "hello"
    if true {
        if (!false) {
            switch v48 {
            case "":
                v74 := GenS8{F0: -128}
                _ = v74
            }
            var v75 uint = v49
            switch "" {
            case "a":
                for fi76 := 0; fi76 < 19; fi76++ {
                    var v77 int8 = 100
                    _ = fi76
                    _ = v77
                }
            }
            _ = v75
        } else {
            if (2 <= v49) {
                var v78 float32 = ((float32((3.14159 / 3.14159)) / 1.0) / 1e10)
                v79 := v5[v48]
                _ = v78
                _ = v79
            } else {
                v5[v48] = "\n\t"
                v67 = v67
            }
            if true {
                var v80 uint64 = 42
                var v81 int = (v67 % (((v67 / ((v67) | 1))) | 1))
                _ = v80
                _ = v81
            }
            for fi82 := 0; fi82 < 11; fi82++ {
                switch "\n\t" {
                case "foo bar":
                    var v83 uint64 = 784
                    _ = v83
                case "hello":
                    switch 2147483647 {
                    case 343:
                        if ((!(!(128 <= 1))) || true) {
                            for fi84 := 0; fi84 < 15; fi84++ {
                                var v85 float32 = float32(float64(1e10))
                                _ = fi84
                                _ = v85
                            }
                            for fi86 := 0; fi86 < 10; fi86++ {
                                v49 = 4294967295
                                _ = fi86
                            }
                            if (!true) {
                                for fi87 := 0; fi87 < 12; fi87++ {
                                    var v88 int32 = -2147483648
                                    v88 = (v88 / ((v88) | 1))
                                    if false {
                                        v5[""] = v48
                                        switch int(v49) {
                                        case 128:
                                            var v89 float32 = 0.1
                                            _ = v89
                                        default:
                                            var v90 float32 = 3.14159
                                            _ = v90
                                        }
                                    }
                                    _ = fi87
                                    _ = v88
                                }
                                v48 = "foo bar"
                            } else {
                                if false {
                                    _ = len(v4)
                                }
                                for fi91 := 0; fi91 < 17; fi91++ {
                                    if false {
                                        v92 := GenS6{F0: 127}
                                        var v93 int32 = 42
                                        fmt.Println(v67)
                                        _ = v92
                                        _ = v93
                                    } else {
                                        v94 := v5["foo bar"]
                                        _ = v94
                                    }
                                    var v95 uint = v49
                                    switch 128 {
                                    case -389:
                                        if false {
                                            if true {
                                                switch (v67 & -1) {
                                                case -244:
                                                    v73 = v73
                                                case 473:
                                                    if (-1 == 42) {
                                                        _ = len(v5)
                                                    } else {
                                                        v49 -= v95
                                                        if true {
                                                            var v96 int8 = -80
                                                            _ = v96
                                                        }
                                                        v2 -= v2
                                                    }
                                                }
                                            }
                                            if ((0.1 + ((1e10 + (0.0 - -2.25)) / 1.0)) == 1.5) {
                                                v1 = v49
                                                if ((^(v95)) != (v95 + 349)) {
                                                    var v97 bool = false
                                                    _ = v97
                                                } else {
                                                    v98 := GenS8{F0: -1}
                                                    v48 = "hello"
                                                    if (0 != uint64(v3)) {
                                                        var v99 int8 = 127
                                                        fmt.Println(fi82, v2, v95)
                                                        delete(v5, v48)
                                                        _ = v99
                                                    }
                                                    _ = v98
                                                }
                                                _ = len(v5)
                                            }
                                            if ((0 <= 9223372036854775807) || true) {
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                var v100 uint8 = 128
                                                _ = v100
                                            } else {
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            }
                                            v3 -= v2
                                            v73 = ((v48 + "foo bar") + "a")
                                        }
                                    case 4294967295:
                                        v101 := GenS6{F0: 0}
                                        _ = v101
                                    default:
                                        v102 := GenS9{F0: fi91, F1: -382, F2: (v95 % (((^(v49))) | 1))}
                                        _ = v102
                                    }
                                    _ = fi91
                                    _ = v95
                                }
                                for fi103 := 0; fi103 < 4; fi103++ {
                                    var v104 string = ("hello" + (v73 + (v48 + "foo bar")))
                                    _ = fi103
                                    _ = v104
                                }
                            }
                        }
                    case 32768:
                        v2 = v2
                    case 1:
                        if false {
                            var v105 int = 374
                            var v106 int16 = 1
                            switch v48 {
                            case "hello":
                                for fi107 := 0; fi107 < 2; fi107++ {
                                    v108 := GenS10{F0: 7, F1: "a", F2: (fi82 << ((4294967295) & 63))}
                                    _ = len(v4)
                                    if len(v4) > 0 {
                                        v4[0] = 4294967295
                                    }
                                    _ = fi107
                                    _ = v108
                                }
                            case "\n\t":
                                v109 := GenS9{F0: (fi82 ^ v67), F1: 127, F2: 67}
                                _ = v109
                            case "":
                                if (true || (50 <= -38)) {
                                    v3 = (^(v2))
                                    if (v2 > v2) {
                                        for fi110 := 0; fi110 < 6; fi110++ {
                                            v105 = fi82
                                            v1 += 1
                                            var v111 uint8 = 1
                                            _ = fi110
                                            _ = v111
                                        }
                                        fmt.Println(v48, v49)
                                    } else {
                                        v49 = (v1 % ((v1) | 1))
                                        v48 = v48
                                        for fi112 := 0; fi112 < 4; fi112++ {
                                            var v113 uint16 = v3
                                            var v114 int64 = -411
                                            fmt.Println(v48)
                                            _ = fi112
                                            _ = v113
                                            _ = v114
                                        }
                                    }
                                } else {
                                    _ = len(v4)
                                    for fi115 := 0; fi115 < 17; fi115++ {
                                        var v116 uint16 = v3
                                        _ = fi115
                                        _ = v116
                                    }
                                }
                            }
                            _ = v105
                            _ = v106
                        }
                    default:
                        var v117 uint = v49
                        _ = v117
                    }
                case "a":
                    if (42 == 42) {
                        v118 := GenS11{F0: v67}
                        _ = v118
                    }
                }
                for fi119 := 0; fi119 < 5; fi119++ {
                    var v120 uint16 = 65535
                    v121 := GenS5{F0: 151, F1: 63}
                    _ = fi119
                    _ = v120
                    _ = v121
                }
                _ = fi82
            }
        }
        v67 = 32767
        var v122 string = (((v73 + "hello") + v73) + v73)
        _ = v122
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v67:", v67)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v66
    _ = v67
    _ = v73
}
