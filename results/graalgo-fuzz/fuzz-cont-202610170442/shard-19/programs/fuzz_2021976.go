package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 uint32; F1 int64 }
type GenS5 struct { F0 int64; F1 uint8; F2 int }
type GenS6 struct { F0 int32; F1 int32 }
type GenS7 struct { F0 string; F1 int16 }
type GenS8 struct { F0 int8; F1 int16; F2 uint32 }

func main() {
    var v1 int16 = 906
    var v2 int = 65
    var v3 int = (v2 * (v2 - v2))
    v4 := make(map[string]bool)
    var v5 bool = (42 >= 121)
    var v6 float64 = (float64(float32(0.1)) * (1.5 + -2.25))
    fmt.Println(v6, v6)
    var v7 int8 = 0
    v6 -= v6
    if (v1 < (v1 - v1)) {
        v6 -= (((((float64((float32(float64(3.14159)) * 1.5)) + 1.5) + float64((0.1 * 1.5))) / v6) / float64(1.5)) / v6)
        v6 = 0.1
        var v8 int8 = (v7 - 42)
        _ = v8
    }
    v9 := GenS1{F0: 42, F1: 4294967295}
    if (209 != 18446744073709551615) {
        var v10 int16 = v1
        v11 := GenS1{F0: 4294967295, F1: 4294967295}
        _ = v10
        _ = v11
    } else {
        fmt.Println(v1, v5)
    }
    if (!v5) {
        if (!(!(965 >= 65535))) {
            for fi12 := 0; fi12 < 3; fi12++ {
                var v13 uint = 65535
                if v5 {
                    v14 := v4["hello"]
                    for fi15 := 0; fi15 < 11; fi15++ {
                        v16 := GenS1{F0: 4294967295, F1: 65535}
                        _ = fi15
                        _ = v16
                    }
                    fmt.Println(fi12)
                    _ = v14
                }
                switch v3 {
                case 4294967295:
                    for fi17 := 0; fi17 < 19; fi17++ {
                        if true {
                            switch (v3 & 42) {
                            case -608:
                                v18 := []int{}
                                _ = v18
                            case 4294967295:
                                v6 += v6
                            default:
                                if false {
                                    v19 := GenS1{F0: 18446744073709551615, F1: (v13 << ((v13) & 63))}
                                    _ = v19
                                }
                            }
                        }
                        v20 := GenS1{F0: 2, F1: v13}
                        var v21 bool = true
                        v1 = (v1 & v1)
                        switch "" {
                        case "hello":
                            delete(v4, ("a" + "\n\t"))
                            v22 := v4[("hello" + "a")]
                            switch -936 {
                            case -64:
                                v23 := GenS2{F0: 801}
                                _ = v23
                            case 0:
                                v1 += 42
                            }
                            var v24 uint = 267
                            _ = v22
                            _ = v24
                        case "":
                            v25 := GenS1{F0: 21, F1: 18446744073709551615}
                            fmt.Println(v7, v1)
                            _ = v25
                        case "\n\t":
                            var v26 float64 = v6
                            _ = v26
                        }
                        _ = fi17
                        _ = v20
                        _ = v21
                    }
                }
                _ = fi12
                _ = v13
            }
            _ = len(v4)
            v27 := []int{v3}
            _ = v27
        } else {
            if v5 {
                for fi28 := 0; fi28 < 2; fi28++ {
                    v7 = v7
                    if v5 {
                        v29 := GenS1{F0: 4294967295, F1: 388}
                        _ = v29
                    } else {
                        var v30 float32 = 1.5
                        switch (("" + "a") + ("hello" + ((("a" + "foo bar") + "hello") + (("hello" + "") + "foo bar")))) {
                        case "a":
                            var v31 int = (v3 / (((fi28 ^ 4294967295)) | 1))
                            _ = v31
                        case "\n\t":
                            v32 := GenS3{F0: -9223372036854775808, F1: 42}
                            _ = v32
                        }
                        _ = v30
                    }
                    v2 += v3
                    _ = fi28
                }
            } else {
                var v33 int32 = -1
                _ = v33
            }
            v34 := GenS1{F0: 18446744073709551615, F1: 2}
            _ = v34
        }
        for fi35 := 0; fi35 < 12; fi35++ {
            if true {
                if v5 {
                    v5 = (32767 < 918)
                    fmt.Println(v3, v2)
                }
            } else {
                if v5 {
                    fmt.Println(v7, v7)
                }
                var v36 string = ""
                var v37 int8 = (v7 + 100)
                _ = v36
                _ = v37
            }
            if (!(!v5)) {
                v38 := v4["foo bar"]
                var v39 string = ""
                if (!(!v5)) {
                    v40 := GenS1{F0: 1, F1: 65535}
                    var v41 uint16 = 509
                    var v42 uint16 = v41
                    _ = v40
                    _ = v41
                    _ = v42
                } else {
                    var v43 float32 = 1e10
                    _ = v43
                }
                _ = v38
                _ = v39
            } else {
                v44 := GenS1{F0: 997, F1: 0}
                var v45 uint8 = 46
                _ = v44
                _ = v45
            }
            v5 = (42 == 223)
            _ = fi35
        }
    } else {
        switch "" {
        case "":
            v7 = (v7 & (-(v7)))
        case "hello":
            v3 = v2
        case "foo bar":
            var v46 uint8 = 128
            _ = v46
        }
        var v47 float64 = (1e10 + v6)
        for fi48 := 0; fi48 < 15; fi48++ {
            var v49 uint16 = 301
            v2 -= (v2 ^ v3)
            _ = fi48
            _ = v49
        }
        _ = v47
    }
    for fi50 := 0; fi50 < 11; fi50++ {
        switch 32768 {
        case 251:
            for fi51 := 0; fi51 < 2; fi51++ {
                switch (((("" + "a") + "") + ("a" + "\n\t")) + "foo bar") {
                case "a":
                    v52 := []int{}
                    _ = v52
                default:
                    switch "hello" {
                    case "a":
                        v53 := []int{(fi51 ^ fi50), 692}
                        _ = v53
                    default:
                        v54 := []int{(-(fi50)), -26, 42, fi50}
                        v1 = -1
                        var v55 float32 = (((-2.25 - -2.25) * (0.0 - 1e10)) * 0.1)
                        _ = v54
                        _ = v55
                    }
                }
                v1 = (v1 ^ 1000)
                _ = fi51
            }
        default:
            for fi56 := 0; fi56 < 6; fi56++ {
                v57 := GenS4{F0: 711, F1: 9223372036854775807}
                _ = fi56
                _ = v57
            }
        }
        if false {
            for fi58 := 0; fi58 < 10; fi58++ {
                if v5 {
                    switch fi50 {
                    case -2147483648:
                        if (v5 || ((!false) && true)) {
                            var v59 uint16 = 42
                            fmt.Println(v59)
                            _ = v59
                        }
                    case 128:
                        if (v5 && false) {
                            var v60 float64 = (v6 * 3.14159)
                            fmt.Println(v60, v2, fi50)
                            v61 := GenS5{F0: 1, F1: 128, F2: (-(fi58))}
                            _ = v60
                            _ = v61
                        }
                    case 32767:
                        var v62 int64 = 0
                        _ = v62
                    default:
                        var v63 uint = 255
                        _ = v63
                    }
                    fmt.Println(fi50, v6, v3)
                    v7 = 42
                    switch (v2 * 531) {
                    case 42:
                        fmt.Println(v5, v1, v3)
                    }
                    var v64 int16 = -684
                    _ = v64
                } else {
                    switch ("\n\t" + ("hello" + "\n\t")) {
                    case "\n\t":
                        fmt.Println(fi58, fi50)
                    case "hello":
                        fmt.Println(v3, fi58)
                    default:
                        var v65 float64 = (float64(float32((v6 / (v6 * (v6 + float64(((((0.1 * (0.1 - ((0.0 / float32(v6)) / (-2.25 * 3.14159)))) * 0.1) + -2.25) / 0.1))))))) / v6)
                        _ = v65
                    }
                }
                var v66 int8 = (v7 + (v7 % (((v7 + (v7 - v7))) | 1)))
                v1 += (v1 % ((v1) | 1))
                _ = fi58
                _ = v66
            }
            v67 := GenS3{F0: 893, F1: int32(fi50)}
            if v5 {
                if ((v2 << ((1) & 63)) >= v3) {
                    v3 += 30
                    var v68 uint32 = 65535
                    _ = v68
                } else {
                    switch (("hello" + "") + "foo bar") {
                    case "":
                        if ((v5 && (324 == 0)) && false) {
                            if ((v7 == 13) && (18446744073709551615 == 1)) {
                                if v5 {
                                    if (2 > 42) {
                                        var v69 int64 = 42
                                        _ = v69
                                    } else {
                                        v70 := []int{(v2 ^ v2), 2147483647}
                                        _ = v70
                                    }
                                    v71 := []int{}
                                    v2 += (v2 ^ v3)
                                    _ = v71
                                } else {
                                    var v72 bool = (("foo bar" + "\n\t") < "a")
                                    switch ("a" + "hello") {
                                    case "a":
                                        if v72 {
                                            v73 := GenS6{F0: -1, F1: 483}
                                            v4[("foo bar" + ("a" + "\n\t"))] = v72
                                            _ = v73
                                        } else {
                                            switch "a" {
                                            case "foo bar":
                                                v6 += v6
                                            case "\n\t":
                                                if (609 <= 2) {
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        if true {
                                                            v74 := []int{v3, v3, (-(v2)), v3}
                                                            v75 := GenS6{F0: 0, F1: -648}
                                                            _ = v74
                                                            _ = v75
                                                        }
                                                    case "":
                                                        for fi76 := 0; fi76 < 19; fi76++ {
                                                            v72 = (v72 || ("" >= "foo bar"))
                                                            _ = fi76
                                                        }
                                                    }
                                                }
                                            }
                                        }
                                    }
                                    _ = v72
                                }
                                if true {
                                    if (!v5) {
                                        v77 := GenS6{F0: 1, F1: -1}
                                        var v78 float64 = (((float64(0.1) - (((v6 + 1e10) - 1e10) + v6)) / float64((3.14159 * 1e10))) * v6)
                                        _ = v77
                                        _ = v78
                                    } else {
                                        if false {
                                            v79 := v4["\n\t"]
                                            v80 := GenS3{F0: 0, F1: 1}
                                            for fi81 := 0; fi81 < 5; fi81++ {
                                                v82 := []int{}
                                                if v79 {
                                                    fmt.Println(v1)
                                                } else {
                                                    _ = len(v82)
                                                    var v83 bool = v5
                                                    _ = v83
                                                }
                                                fmt.Println(v5, v1, v6)
                                                _ = fi81
                                                _ = v82
                                            }
                                            _ = v79
                                            _ = v80
                                        }
                                        if false {
                                            for fi84 := 0; fi84 < 15; fi84++ {
                                                v85 := []int{}
                                                if (!(255 <= 1)) {
                                                    if (v5 || v5) {
                                                        if v5 {
                                                            v2 -= (fi84 - -105)
                                                            _ = len(v85)
                                                        } else {
                                                            v3 -= (v3 ^ 4294967295)
                                                            v86 := v4[("" + "\n\t")]
                                                            _ = v86
                                                        }
                                                        switch "" {
                                                        case "":
                                                            var v87 int = fi50
                                                            _ = v87
                                                        default:
                                                            v88 := GenS7{F0: "foo bar", F1: v1}
                                                            _ = v88
                                                        }
                                                        fmt.Println(fi84)
                                                    } else {
                                                        if ((3.14159 / float32((v6 + 3.14159))) < 1e10) {
                                                            v1 = v1
                                                        } else {
                                                            if (-884 > v1) {
                                                                if (((1e10 / v6) * 3.14159) == 0.1) {
                                                                    _ = len(v85)
                                                                    var v89 int8 = -1
                                                                    _ = v89
                                                                } else {
                                                                    switch fi50 {
                                                                    case -494:
                                                                        v85 = append(v85, (v3 | fi84))
                                                                    case -590:
                                                                        if v5 {
                                                                            if (v6 < v6) {
                                                                                v5 = (!true)
                                                                            } else {
                                                                                v7 -= 1
                                                                                var v90 float32 = (((1.5 / 0.1) * 0.0) / 0.1)
                                                                                _ = v90
                                                                            }
                                                                        } else {
                                                                            fmt.Println(v5)
                                                                            v3 = fi50
                                                                            v3 -= fi84
                                                                            fmt.Println(fi50, v3)
                                                                        }
                                                                    default:
                                                                        var v91 uint8 = 42
                                                                        _ = v91
                                                                    }
                                                                }
                                                            }
                                                            v92 := GenS4{F0: 42, F1: 0}
                                                            _ = v92
                                                        }
                                                    }
                                                    switch (v3 * (v2 - fi84)) {
                                                    case 4294967295:
                                                        fmt.Println(v6, v5, v5)
                                                    }
                                                }
                                                _ = fi84
                                                _ = v85
                                            }
                                            fmt.Println(v5, v5, v2)
                                            var v93 uint32 = 65535
                                            _ = v93
                                        } else {
                                            v1 += v1
                                            var v94 uint8 = 1
                                            var v95 float64 = float64((0.0 * float32(v6)))
                                            _ = v94
                                            _ = v95
                                        }
                                    }
                                } else {
                                    if v5 {
                                        v5 = true
                                        if (917 == 354) {
                                            fmt.Println(v1, fi50)
                                            for fi96 := 0; fi96 < 1; fi96++ {
                                                if v5 {
                                                    fmt.Println(v7, fi96, v5)
                                                    switch ("foo bar" + "foo bar") {
                                                    case "foo bar":
                                                        v97 := []int{(v2 ^ fi96), v2}
                                                        _ = v97
                                                    }
                                                    var v98 uint16 = 551
                                                    _ = v98
                                                }
                                                var v99 float32 = float32(1.5)
                                                var v100 int16 = v1
                                                v101 := GenS8{F0: (v7 ^ (v7 / (((v7 >> ((0) & 63))) | 1))), F1: v1, F2: 597}
                                                _ = fi96
                                                _ = v100
                                                _ = v101
                                                _ = v99
                                            }
                                            if v5 {
                                                var v102 float32 = 0.1
                                                _ = v102
                                            } else {
                                                v4[("" + "\n\t")] = (((1e10 + 0.1) != 1.5) && v5)
                                                var v103 bool = (18446744073709551615 == 65535)
                                                fmt.Println(v2, v7, v7)
                                                _ = v103
                                            }
                                        } else {
                                            if (!(2 <= 287)) {
                                                v104 := []int{(-(v2)), (^(v3)), (-(fi50))}
                                                _ = v104
                                            } else {
                                                v105 := GenS3{F0: 42, F1: -1}
                                                fmt.Println(v3)
                                                v2 = (fi50 ^ fi50)
                                                _ = v105
                                            }
                                        }
                                    } else {
                                        v5 = ((!(1.5 == (3.14159 + v6))) && false)
                                    }
                                    switch v2 {
                                    case 1:
                                        var v106 string = "foo bar"
                                        _ = v106
                                    case 127:
                                        v5 = (65535 < 42)
                                    default:
                                        var v107 uint32 = 166
                                        _ = v107
                                    }
                                    if (!v5) {
                                        var v108 float64 = (v6 + (v6 * float64(1.5)))
                                        _ = v108
                                    } else {
                                        v109 := []int{v3, v2, fi50}
                                        _ = v109
                                    }
                                }
                                for fi110 := 0; fi110 < 5; fi110++ {
                                    var v111 float32 = 1.5
                                    v112 := []int{0, (fi50 | 2147483647), 32768, (fi50 << ((0) & 63))}
                                    _ = fi110
                                    _ = v111
                                    _ = v112
                                }
                            } else {
                                var v113 uint32 = 65535
                                _ = v113
                            }
                        }
                    case "a":
                        v7 = v7
                    default:
                        var v114 uint8 = 163
                        _ = v114
                    }
                    var v115 int64 = -1
                    v6 += v6
                    _ = v115
                }
                v2 = (fi50 % ((32767) | 1))
            } else {
                switch v2 {
                case 2147483647:
                    for fi116 := 0; fi116 < 8; fi116++ {
                        v2 += 42
                        _ = fi116
                    }
                case 0:
                    v117 := []int{(fi50 ^ v3), v2}
                    _ = v117
                }
            }
            _ = v67
        } else {
            var v118 int8 = -43
            _ = v118
        }
        v5 = (false || v5)
        _ = fi50
    }
    for fi119 := 0; fi119 < 3; fi119++ {
        switch fi119 {
        case 32768:
            v3 = (v2 >> ((255) & 63))
        }
        _ = fi119
    }
    v7 += v7
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
