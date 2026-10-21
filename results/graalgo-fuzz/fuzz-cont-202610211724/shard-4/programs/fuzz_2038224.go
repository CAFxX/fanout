package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 uint16; F1 uint; F2 int64 }
type GenS5 struct { F0 uint64; F1 uint; F2 int64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int64; F1 int64; F2 int16 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint64; F1 int8; F2 int16 }
type GenS11 struct { F0 int64; F1 int }
type GenS12 struct { F0 int64; F1 uint64 }
type GenS13 struct { F0 int }
type GenS14 struct { F0 uint8; F1 uint32 }
type GenS15 struct { F0 int16; F1 bool }
type GenS16 struct { F0 string; F1 uint8; F2 int }
type GenS17 struct { F0 int32; F1 int64 }
type GenS18 struct { F0 bool; F1 uint64 }
type GenS19 struct { F0 uint64; F1 int32 }
type GenS20 struct { F0 int64 }
type GenS21 struct { F0 int64; F1 int32 }
type GenS22 struct { F0 int16; F1 int32 }
type GenS23 struct { F0 string }
func genFn0() int8 {
    v6 := []int{}
    var v7 bool = (840 <= 1)
    if (v7 && v7) {
        var v8 int64 = 1
        v9 := GenS1{F0: -2147483648, F1: 7}
        _ = v8
        _ = v9
    }
    _ = v6
    _ = v7
    return 33
}

func main() {
    v1 := []string{(("foo bar" + "\n\t") + ""), "hello", ((("" + "") + "") + (("\n\t" + "\n\t") + "hello")), (("" + "\n\t") + "")}
    var v2 uint = 4294967295
    var v3 int32 = 1
    var v4 uint32 = 42
    var v5 string = ""
    v4 += (v4 + 910)
    _ = len(v1)
    for fi10 := 0; fi10 < 18; fi10++ {
        if ((!true) || (1 != 127)) {
            v11 := GenS2{F0: int16(fi10)}
            for fi12 := 0; fi12 < 7; fi12++ {
                if true {
                    var v13 int64 = -371
                    var v14 float64 = ((3.14159 + 1.5) * 1.5)
                    _ = v13
                    _ = v14
                }
                _ = fi12
            }
            var v15 float32 = 1e10
            _ = v11
            _ = v15
        }
        _ = fi10
    }
    switch ((v5 + (v5 + (v5 + "foo bar"))) + v5) {
    case "foo bar":
        v1 = append(v1, "foo bar")
    default:
        for fi16 := 0; fi16 < 1; fi16++ {
            v1 = append(v1, (((("\n\t" + ("foo bar" + ("foo bar" + v5))) + v5) + "hello") + "a"))
            switch ("hello" + v5) {
            case "":
                switch v5 {
                case "\n\t":
                    if (!(!(2 > 154))) {
                        if (-1 > 32767) {
                            v2 -= (v2 % ((0) | 1))
                        } else {
                            v4 -= v4
                        }
                    } else {
                        fmt.Println(fi16)
                        if len(v1) > 0 {
                            v1[0] = "foo bar"
                        }
                        for fi17 := 0; fi17 < 15; fi17++ {
                            var v18 int32 = v3
                            v1 = append(v1, ((v5 + (("foo bar" + v5) + ((v5 + "hello") + ("hello" + v5)))) + (("foo bar" + v5) + v5)))
                            _ = fi17
                            _ = v18
                        }
                    }
                default:
                    v5 = v5
                }
            case "hello":
                v3 -= 1
            case "\n\t":
                if true {
                    var v19 string = "a"
                    v20 := GenS2{F0: 0}
                    _ = v19
                    _ = v20
                }
            }
            var v21 int = -1
            _ = fi16
            _ = v21
        }
    }
    if len(v1) > 0 {
        v1[0] = v5
    }
    v5 = (v5 + v5)
    var v22 uint = 4294967295
    if false {
        for fi23 := 0; fi23 < 4; fi23++ {
            switch (((v5 + "") + (v5 + v5)) + v5) {
            case "\n\t":
                if false {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = len(v1)
                    switch int(v3) {
                    case 1:
                        v1 = append(v1, "")
                    case 672:
                        fmt.Println(v22)
                    case 550:
                        if (((v5 + (((("foo bar" + ((v5 + v5) + "\n\t")) + v5) + "a") + "a")) + v5) <= (v5 + (v5 + "a"))) {
                            fmt.Println(v3, v3, fi23)
                            v2 = 156
                            v24 := GenS2{F0: 42}
                            _ = v24
                        }
                    default:
                        var v25 int8 = 127
                        _ = v25
                    }
                }
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if true {
                for fi26 := 0; fi26 < 1; fi26++ {
                    if true {
                        fmt.Println(v3, v3)
                    } else {
                        v27 := GenS2{F0: -32768}
                        _ = v27
                    }
                    var v28 float32 = (0.1 - 1.5)
                    switch "a" {
                    case "a":
                        v22 += (v2 >> (((v22 / (((^(v2))) | 1))) & 63))
                    case "foo bar":
                        switch v5 {
                        case "\n\t":
                            var v29 uint = 4294967295
                            var v30 int8 = -38
                            _ = v29
                            _ = v30
                        case "foo bar":
                            v31 := GenS2{F0: 0}
                            _ = v31
                        default:
                            if len(v1) > 0 {
                                v1[0] = "a"
                            }
                        }
                    case "":
                        fmt.Println(fi26)
                    }
                    fmt.Println(v22)
                    _ = fi26
                    _ = v28
                }
                if ((false && (!(1 == -1))) && true) {
                    for fi32 := 0; fi32 < 20; fi32++ {
                        fmt.Println(v22, v5)
                        switch v5 {
                        case "\n\t":
                            if false {
                                if true {
                                    v33 := GenS3{F0: genFn0(), F1: (v22 >> ((2) & 63))}
                                    var v34 int32 = v3
                                    _ = v33
                                    _ = v34
                                } else {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                v35 := GenS3{F0: -1, F1: 255}
                                _ = v35
                            }
                        default:
                            fmt.Println(fi32, v3)
                        }
                        switch 256 {
                        case -2147483648:
                            var v36 string = "a"
                            _ = v36
                        }
                        _ = fi32
                    }
                    for fi37 := 0; fi37 < 8; fi37++ {
                        fmt.Println(v22)
                        if false {
                            var v38 string = v5
                            fmt.Println(v5)
                            v3 = 42
                            _ = v38
                        } else {
                            v4 = (v4 - 2)
                            switch fi23 {
                            case -9223372036854775808:
                                if true {
                                    v3 -= v3
                                } else {
                                    fmt.Println(v4, fi37, v3)
                                    var v39 float64 = 3.14159
                                    v39 += ((v39 - (float64(1e10) + (-2.25 - v39))) / float64(float32(v39)))
                                    _ = v39
                                }
                            case 32768:
                                v4 = (v4 / (((^(v4))) | 1))
                            default:
                                if ((32767 < 32767) || true) {
                                    if true {
                                        v5 = v5
                                    } else {
                                        var v40 uint32 = v4
                                        var v41 int64 = -9223372036854775808
                                        v1 = append(v1, v5)
                                        _ = v40
                                        _ = v41
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = ((v5 + v5) + v5)
                                    }
                                    var v42 int = (-(fi23))
                                    _ = v42
                                } else {
                                    fmt.Println(v22, v5)
                                }
                            }
                        }
                        v4 += 161
                        fmt.Println(fi23)
                        switch (fi37 / ((fi23) | 1)) {
                        case 128:
                            v4 = v4
                            var v43 uint = 65535
                            if true {
                                if true {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                } else {
                                    var v44 int8 = genFn0()
                                    fmt.Println(v5, v3)
                                    switch v5 {
                                    case "foo bar":
                                        if false {
                                            fmt.Println(v44, v5)
                                            var v45 uint64 = 119
                                            fmt.Println(v45, v3)
                                            _ = v45
                                        } else {
                                            if len(v1) > 0 {
                                                v1[0] = "foo bar"
                                            }
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            if (!(18446744073709551615 < uint64(v4))) {
                                                v46 := GenS3{F0: (^(v44)), F1: v2}
                                                v1 = append(v1, (v5 + v5))
                                                var v47 string = v5
                                                var v48 int32 = v3
                                                switch fi23 {
                                                case 2147483647:
                                                    v49 := GenS1{F0: (v3 * (v48 & v48)), F1: 128}
                                                    _ = v49
                                                }
                                                switch v5 {
                                                case "foo bar":
                                                    v50 := GenS2{F0: 1000}
                                                    _ = v50
                                                case "":
                                                    if false {
                                                        if false {
                                                            var v51 int = (fi37 ^ int(v48))
                                                            _ = len(v1)
                                                            _ = v51
                                                        }
                                                        switch int(v44) {
                                                        case -457:
                                                            v48 -= 0
                                                        case -651:
                                                            if ((!(!false)) || false) {
                                                                v1 = append(v1, "")
                                                                var v52 int64 = 871
                                                                if true {
                                                                    fmt.Println(v47, v22)
                                                                    fmt.Println(v48, fi23)
                                                                }
                                                                _ = v52
                                                            }
                                                        case 127:
                                                            v5 = "a"
                                                        default:
                                                            v22 -= (^(v43))
                                                            if (408 >= v43) {
                                                                v1 = append(v1, "\n\t")
                                                                v1 = append(v1, ("\n\t" + (v5 + "")))
                                                                v2 -= v22
                                                                v48 -= (v48 / (((v3 << ((75) & 63))) | 1))
                                                                v1 = append(v1, "foo bar")
                                                                switch fi37 {
                                                                case 720:
                                                                    fmt.Println(v22, v2)
                                                                case 127:
                                                                    if (true && false) {
                                                                        if ((false || true) && true) {
                                                                            var v53 int32 = (^(v3))
                                                                            var v54 uint8 = 180
                                                                            v55 := GenS4{F0: 1, F1: 42, F2: -973}
                                                                            _ = v53
                                                                            _ = v54
                                                                            _ = v55
                                                                        }
                                                                        if (0 >= v4) {
                                                                            var v56 uint64 = 1
                                                                            if (888 == 3) {
                                                                                if false {
                                                                                    v5 = ("\n\t" + v5)
                                                                                    if true {
                                                                                        var v57 int16 = 42
                                                                                        if len(v1) >= 2 {
                                                                                            v1 = v1[1:len(v1)]
                                                                                        }
                                                                                        _ = v57
                                                                                    } else {
                                                                                        var v58 int8 = (v44 + (^(v44)))
                                                                                        if len(v1) >= 2 {
                                                                                            v1 = v1[1:len(v1)]
                                                                                        }
                                                                                        _ = v58
                                                                                    }
                                                                                    switch -9223372036854775808 {
                                                                                    case -2147483648:
                                                                                        if (!false) {
                                                                                            v22 = 2
                                                                                        } else {
                                                                                            if len(v1) > 0 {
                                                                                                v1[0] = (v5 + "a")
                                                                                            }
                                                                                            var v59 float32 = ((1.5 / (0.0 + ((3.14159 * 1e10) * (0.1 * 3.14159)))) - (float32(0.1) * 1.5))
                                                                                            _ = v59
                                                                                        }
                                                                                    case -750:
                                                                                        var v60 int32 = (v3 % ((v3) | 1))
                                                                                        _ = v60
                                                                                    case 128:
                                                                                        if true {
                                                                                            var v61 uint32 = v4
                                                                                            _ = v61
                                                                                        }
                                                                                    default:
                                                                                        var v62 uint16 = 2
                                                                                        _ = v62
                                                                                    }
                                                                                    v3 += -2147483648
                                                                                }
                                                                                v1 = append(v1, ((v5 + "\n\t") + ((("a" + v5) + (v5 + "foo bar")) + (v47 + "a"))))
                                                                                v22 = v22
                                                                                var v63 int32 = (-(v48))
                                                                                if true {
                                                                                    v3 = v48
                                                                                    var v64 int8 = genFn0()
                                                                                    var v65 uint16 = 42
                                                                                    _ = v64
                                                                                    _ = v65
                                                                                } else {
                                                                                    if len(v1) >= 2 {
                                                                                        v1 = v1[1:len(v1)]
                                                                                    }
                                                                                }
                                                                                v1 = append(v1, (("a" + (v47 + (v47 + ""))) + v47))
                                                                                _ = v63
                                                                            }
                                                                            v66 := GenS5{F0: v56, F1: v43, F2: -27}
                                                                            switch (v5 + "") {
                                                                            case "foo bar":
                                                                                switch fi23 {
                                                                                case 0:
                                                                                    v48 = v48
                                                                                default:
                                                                                    switch (-(fi23)) {
                                                                                    case 9223372036854775807:
                                                                                        v1 = append(v1, v5)
                                                                                    case 128:
                                                                                        if len(v1) > 0 {
                                                                                            v1[0] = v5
                                                                                        }
                                                                                    }
                                                                                }
                                                                            default:
                                                                                v67 := GenS3{F0: v44, F1: (v43 | v2)}
                                                                                _ = v67
                                                                            }
                                                                            _ = v56
                                                                            _ = v66
                                                                        }
                                                                    }
                                                                }
                                                                v47 = ""
                                                                fmt.Println(v2, fi23, v47)
                                                            }
                                                        }
                                                    }
                                                case "a":
                                                    v44 -= (v44 << (((v2 >> ((4294967295) & 63))) & 63))
                                                default:
                                                    if (1000 != 32767) {
                                                        switch v47 {
                                                        case "":
                                                            switch fi23 {
                                                            case 4294967295:
                                                                if false {
                                                                    var v68 int = (fi23 ^ (fi37 * fi37))
                                                                    _ = v68
                                                                } else {
                                                                    if true {
                                                                        var v69 string = v47
                                                                        if false {
                                                                            var v70 int16 = 0
                                                                            _ = v70
                                                                        }
                                                                        v71 := GenS1{F0: (v48 % ((v3) | 1)), F1: 2}
                                                                        _ = v69
                                                                        _ = v71
                                                                    }
                                                                    v22 = v22
                                                                }
                                                            case -2147483648:
                                                                var v72 float32 = -2.25
                                                                _ = v72
                                                            default:
                                                                v48 = (-(v48))
                                                            }
                                                        case "hello":
                                                            if false {
                                                                var v73 int8 = v44
                                                                _ = v73
                                                            }
                                                        }
                                                        v22 = v2
                                                        _ = len(v1)
                                                    } else {
                                                        fmt.Println(v2, v22, v2)
                                                        if len(v1) > 0 {
                                                            v1[0] = ((v5 + v47) + (v5 + "\n\t"))
                                                        }
                                                    }
                                                }
                                                _ = v46
                                                _ = v47
                                                _ = v48
                                            } else {
                                                switch v5 {
                                                case "\n\t":
                                                    var v74 int16 = 790
                                                    _ = v74
                                                case "a":
                                                    v2 = v22
                                                case "hello":
                                                    v75 := GenS3{F0: 0, F1: (v22 + v2)}
                                                    _ = v75
                                                }
                                            }
                                        }
                                    }
                                    _ = v44
                                }
                            }
                            _ = v43
                        }
                        v3 -= -506
                        _ = fi37
                    }
                    if (false || false) {
                        v2 = 65535
                        v2 = (v22 % ((430) | 1))
                        if true {
                            v76 := GenS6{F0: 65535}
                            fmt.Println(v4, fi23, v5)
                            v77 := GenS5{F0: 1, F1: (v2 & (v2 << (((^(v2))) & 63))), F2: 42}
                            _ = v76
                            _ = v77
                        } else {
                            if (0.1 == (0.1 + 0.0)) {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                for fi78 := 0; fi78 < 4; fi78++ {
                                    v3 = v3
                                    v1 = append(v1, "")
                                    var v79 uint64 = 253
                                    _ = fi78
                                    _ = v79
                                }
                            }
                        }
                    } else {
                        switch fi23 {
                        case 32767:
                            var v80 float32 = (-2.25 * (1e10 + 3.14159))
                            _ = v80
                        }
                    }
                } else {
                    switch fi23 {
                    case -350:
                        var v81 int32 = -1
                        _ = v81
                    case -976:
                        var v82 int32 = (v3 ^ v3)
                        _ = v82
                    case 32768:
                        if false {
                            if false {
                                if true {
                                    switch (fi23 + fi23) {
                                    case -418:
                                        fmt.Println(v4)
                                    case 814:
                                        v83 := GenS2{F0: 32767}
                                        _ = v83
                                    }
                                }
                                if (false || (false && (42 >= v3))) {
                                    fmt.Println(v22, v3, v4)
                                    v5 = (("a" + "a") + v5)
                                } else {
                                    var v84 float64 = 0.0
                                    var v85 int64 = -530
                                    _ = v84
                                    _ = v85
                                }
                            } else {
                                v5 = v5
                            }
                            for fi86 := 0; fi86 < 16; fi86++ {
                                v87 := GenS1{F0: (v3 + 540), F1: 255}
                                _ = fi86
                                _ = v87
                            }
                            v88 := GenS7{F0: int64(fi23), F1: 250, F2: 32767}
                            _ = v88
                        } else {
                            v5 = ("" + ((v5 + (v5 + "a")) + v5))
                            var v89 int16 = -335
                            _ = v89
                        }
                    default:
                        var v90 bool = (756 < (v3 >> ((42) & 63)))
                        _ = v90
                    }
                    var v91 uint8 = 1
                    _ = v91
                }
                if len(v1) > 0 {
                    v1[0] = v5
                }
            }
            _ = fi23
        }
        for fi92 := 0; fi92 < 6; fi92++ {
            switch (fi92 - fi92) {
            case 854:
                for fi93 := 0; fi93 < 18; fi93++ {
                    v94 := GenS8{F0: 0}
                    if true {
                        v22 += 65535
                        v1 = append(v1, "")
                    } else {
                        var v95 uint8 = 2
                        _ = len(v1)
                        switch 9223372036854775807 {
                        case 32768:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        case -1:
                            v96 := GenS9{F0: 0}
                            _ = v96
                        default:
                            var v97 int = (fi93 / (((fi93 * (fi92 << ((2) & 63)))) | 1))
                            _ = v97
                        }
                        switch v5 {
                        case "\n\t":
                            v5 = (("\n\t" + v5) + v5)
                        case "a":
                            var v98 float32 = (0.0 + 1.5)
                            _ = v98
                        case "hello":
                            v99 := GenS10{F0: 68, F1: 25, F2: 278}
                            _ = v99
                        }
                        if false {
                            fmt.Println(fi93)
                            if (0 > 18446744073709551615) {
                                var v100 int64 = -916
                                fmt.Println(v22, v2)
                                if false {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                v101 := GenS5{F0: 42, F1: (v2 | (v2 | 2)), F2: (v100 % ((42) | 1))}
                                var v102 uint = 18446744073709551615
                                if false {
                                    switch v5 {
                                    case "hello":
                                        v4 += (v4 - (v4 - 0))
                                    case "a":
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    case "":
                                        v3 += v3
                                    }
                                    if (((^(v4)) <= (v4 >> (((v102 + (v2 ^ v22))) & 63))) || false) {
                                        v22 += (v2 & 641)
                                        var v103 uint32 = v4
                                        var v104 int16 = -32768
                                        _ = v103
                                        _ = v104
                                    }
                                }
                                _ = v100
                                _ = v101
                                _ = v102
                            } else {
                                switch (fi92 >> (((v22 ^ v2)) & 63)) {
                                case 9223372036854775807:
                                    var v105 int8 = genFn0()
                                    _ = v105
                                }
                            }
                            switch (fi93 + fi92) {
                            case 415:
                                v95 = (v95 << ((uint(v95)) & 63))
                            case 665:
                                var v106 int32 = (v3 + v3)
                                _ = v106
                            }
                        } else {
                            v1 = append(v1, v5)
                            if ((-(v3)) < (v3 % (((v3 - (v3 - (v3 * (v3 % ((-1) | 1)))))) | 1))) {
                                if true {
                                    switch fi92 {
                                    case -975:
                                        _ = len(v1)
                                        v107 := GenS4{F0: 1, F1: v2, F2: 0}
                                        _ = v107
                                    case -9223372036854775808:
                                        _ = len(v1)
                                    default:
                                        var v108 uint16 = 65535
                                        _ = v108
                                    }
                                }
                                v109 := GenS9{F0: 255}
                                _ = v109
                            } else {
                                var v110 int64 = 819
                                var v111 int8 = -115
                                var v112 bool = true
                                _ = v110
                                _ = v111
                                _ = v112
                            }
                        }
                        _ = v95
                    }
                    if (-2.25 <= 0.1) {
                        var v113 int64 = 1
                        _ = v113
                    }
                    _ = fi93
                    _ = v94
                }
            case -9223372036854775808:
                for fi114 := 0; fi114 < 14; fi114++ {
                    switch (v5 + (("a" + v5) + v5)) {
                    case "foo bar":
                        switch "a" {
                        case "foo bar":
                            if (-360 >= -353) {
                                switch (^(fi114)) {
                                case -1:
                                    v3 = -2147483648
                                case 32768:
                                    v115 := GenS7{F0: 42, F1: 623, F2: 945}
                                    v4 = 1
                                    if len(v1) > 0 {
                                        v1[0] = v5
                                    }
                                    _ = v115
                                case 128:
                                    fmt.Println(v3, fi114)
                                default:
                                    v3 = -835
                                }
                                if (-66 <= 100) {
                                    if (v2 >= (v2 >> ((2) & 63))) {
                                        fmt.Println(v4, v2, fi92)
                                    }
                                    var v116 int16 = -660
                                    v117 := GenS8{F0: 531}
                                    _ = v116
                                    _ = v117
                                } else {
                                    if (!(!(32767 > 992))) {
                                        _ = len(v1)
                                        var v118 int64 = 1
                                        v22 = v2
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        var v119 uint16 = 1
                                        if true {
                                            var v120 int8 = 0
                                            v1 = append(v1, (v5 + ""))
                                            v2 = (v2 * 0)
                                            v119 = 32767
                                            _ = v120
                                        } else {
                                            switch "" {
                                            case "foo bar":
                                                var v121 bool = false
                                                _ = v121
                                            default:
                                                if ((!false) && false) {
                                                    var v122 float64 = 1.5
                                                    v123 := GenS11{F0: v118, F1: (fi114 + fi114)}
                                                    _ = v122
                                                    _ = v123
                                                }
                                            }
                                        }
                                        _ = v118
                                        _ = v119
                                    }
                                    var v124 int32 = 42
                                    switch fi92 {
                                    case -9223372036854775808:
                                        v5 = (v5 + "a")
                                    }
                                    _ = v124
                                }
                                v5 = v5
                                v4 -= 2
                            } else {
                                switch 1 {
                                case -835:
                                    v125 := GenS1{F0: (v3 & (v3 ^ (-(v3)))), F1: 2}
                                    _ = v125
                                case 9223372036854775807:
                                    v22 += v22
                                }
                                if (true || false) {
                                    var v126 uint32 = (v4 % ((uint32(v22)) | 1))
                                    fmt.Println(v126, fi92, fi92)
                                    _ = v126
                                }
                                v3 -= 2147483647
                            }
                        }
                        var v127 bool = (!false)
                        v5 = v5
                        var v128 int32 = (v3 << ((v22) & 63))
                        _ = v127
                        _ = v128
                    }
                    _ = fi114
                }
            }
            _ = fi92
        }
        var v129 uint = (^(v2))
        _ = v129
    } else {
        v3 -= v3
        if false {
            for fi130 := 0; fi130 < 9; fi130++ {
                var v131 int32 = (v3 << (((v2 << (((v2 ^ v22)) & 63))) & 63))
                _ = fi130
                _ = v131
            }
            switch (v5 + (v5 + (v5 + v5))) {
            case "\n\t":
                var v132 uint64 = 2
                _ = v132
            case "hello":
                switch -9223372036854775808 {
                case 0:
                    v2 -= v22
                case -845:
                    var v133 uint = 273
                    _ = v133
                case 529:
                    for fi134 := 0; fi134 < 11; fi134++ {
                        if (!true) {
                            v135 := GenS3{F0: genFn0(), F1: v22}
                            _ = v135
                        }
                        var v136 bool = false
                        var v137 uint8 = 128
                        _ = fi134
                        _ = v136
                        _ = v137
                    }
                default:
                    switch -706 {
                    case 21:
                        switch 55 {
                        case -9223372036854775808:
                            v138 := GenS5{F0: uint64(v4), F1: v22, F2: 1}
                            _ = v138
                        case 359:
                            switch 431 {
                            case 2147483647:
                                v5 = ((v5 + v5) + ((v5 + ((v5 + "\n\t") + (v5 + "a"))) + ((v5 + "foo bar") + v5)))
                            case -340:
                                var v139 uint64 = 260
                                _ = v139
                            case 32768:
                                var v140 uint8 = 2
                                _ = v140
                            default:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                        case 128:
                            v141 := GenS12{F0: 809, F1: 42}
                            _ = v141
                        default:
                            if (!(-1 <= 42)) {
                                if (42 > 255) {
                                    v2 -= (^(v22))
                                    if (42 != 42) {
                                        for fi142 := 0; fi142 < 12; fi142++ {
                                            var v143 int32 = (^(v3))
                                            var v144 int64 = -90
                                            for fi145 := 0; fi145 < 19; fi145++ {
                                                v146 := GenS9{F0: 402}
                                                if (v5 == v5) {
                                                    if (!true) {
                                                        v143 = (v3 | (v3 + v143))
                                                        v1 = append(v1, v5)
                                                        v147 := GenS12{F0: (v144 * 9223372036854775807), F1: 0}
                                                        _ = v147
                                                    } else {
                                                        v148 := GenS13{F0: fi145}
                                                        _ = v148
                                                    }
                                                    if true {
                                                        v2 += (v22 << ((v22) & 63))
                                                        if false {
                                                            var v149 string = v5
                                                            var v150 uint8 = 127
                                                            switch (fi142 ^ fi145) {
                                                            case 1:
                                                                var v151 int64 = 42
                                                                _ = v151
                                                            case -1:
                                                                var v152 float32 = 0.1
                                                                _ = v152
                                                            default:
                                                                v1 = append(v1, v149)
                                                            }
                                                            _ = v149
                                                            _ = v150
                                                        }
                                                    }
                                                    v153 := GenS14{F0: 193, F1: v4}
                                                    v4 -= v4
                                                    v154 := GenS15{F0: 32767, F1: (!true)}
                                                    _ = v153
                                                    _ = v154
                                                }
                                                v4 += (v4 ^ v4)
                                                _ = fi145
                                                _ = v146
                                            }
                                            _ = fi142
                                            _ = v143
                                            _ = v144
                                        }
                                        v4 -= (v4 % ((v4) | 1))
                                    } else {
                                        var v155 uint64 = 42
                                        for fi156 := 0; fi156 < 9; fi156++ {
                                            v2 -= (v22 % ((v2) | 1))
                                            for fi157 := 0; fi157 < 5; fi157++ {
                                                fmt.Println(v4, v22, v155)
                                                v158 := GenS14{F0: 42, F1: v4}
                                                _ = fi157
                                                _ = v158
                                            }
                                            v159 := GenS5{F0: v155, F1: (v22 << ((v22) & 63)), F2: 0}
                                            _ = fi156
                                            _ = v159
                                        }
                                        _ = v155
                                    }
                                }
                                var v160 int32 = v3
                                for fi161 := 0; fi161 < 4; fi161++ {
                                    if ((v5 == "a") && false) {
                                        if false {
                                            var v162 int64 = 1
                                            _ = v162
                                        } else {
                                            var v163 int = 924
                                            _ = v163
                                        }
                                        var v164 float32 = 0.1
                                        v160 = (v160 / ((1) | 1))
                                        _ = v164
                                    }
                                    v165 := GenS2{F0: 830}
                                    v160 = 42
                                    _ = fi161
                                    _ = v165
                                }
                                _ = v160
                            } else {
                                _ = len(v1)
                                v5 = "hello"
                            }
                        }
                    case -1:
                        if len(v1) > 0 {
                            v1[0] = (v5 + "")
                        }
                    default:
                        var v166 uint16 = 1
                        _ = v166
                    }
                }
            case "a":
                v1 = append(v1, v5)
            default:
                v167 := GenS16{F0: v5, F1: uint8(v4), F2: 4294967295}
                _ = v167
            }
        } else {
            var v168 uint16 = 1
            switch 4294967295 {
            case 127:
                for fi169 := 0; fi169 < 20; fi169++ {
                    var v170 uint = uint(v168)
                    _ = fi169
                    _ = v170
                }
            default:
                v2 = v22
            }
            _ = v168
        }
    }
    v2 += (v2 ^ (^(v22)))
    var v171 int = -1
    switch v171 {
    case 2147483647:
        v2 += (v2 - 18446744073709551615)
    default:
        if ((((v171 + v171) <= 32767) || true) || true) {
            v1 = append(v1, "")
            if false {
                v5 = ((v5 + (((v5 + (("\n\t" + "foo bar") + v5)) + "hello") + ("hello" + "a"))) + v5)
            } else {
                for fi172 := 0; fi172 < 14; fi172++ {
                    if (false || false) {
                        var v173 uint64 = 1
                        _ = v173
                    } else {
                        fmt.Println(v22, v4)
                    }
                    fmt.Println(v171, fi172)
                    if (!(false && (1 == 0))) {
                        switch fi172 {
                        case 127:
                            v174 := GenS17{F0: 629, F1: 0}
                            _ = v174
                        case 128:
                            fmt.Println(v2, v4, v22)
                        }
                        var v175 int16 = 43
                        _ = v175
                    }
                    _ = fi172
                }
                _ = len(v1)
            }
            if (!(!(v171 > (v171 << ((v2) & 63))))) {
                if true {
                    var v176 int = 128
                    _ = v176
                }
                for fi177 := 0; fi177 < 8; fi177++ {
                    v5 = (v5 + v5)
                    v5 = ""
                    _ = fi177
                }
                var v178 string = "foo bar"
                _ = v178
            } else {
                switch v5 {
                case "foo bar":
                    if (1.5 < 1e10) {
                        for fi179 := 0; fi179 < 2; fi179++ {
                            fmt.Println(fi179, v3)
                            v2 += (v22 - (v2 | (v22 - 255)))
                            _ = fi179
                        }
                    }
                }
                if (true || (-2.25 > float64(1.5))) {
                    if false {
                        if (42 != 787) {
                            v180 := GenS1{F0: v3, F1: 114}
                            _ = v180
                        } else {
                            if len(v1) > 0 {
                                v1[0] = v5
                            }
                            for fi181 := 0; fi181 < 18; fi181++ {
                                switch "foo bar" {
                                case "\n\t":
                                    switch (fi181 % (((fi181 / ((v171) | 1))) | 1)) {
                                    case -9223372036854775808:
                                        v182 := GenS18{F0: (9223372036854775807 <= 2147483647), F1: 4294967295}
                                        _ = v182
                                    case -22:
                                        for fi183 := 0; fi183 < 16; fi183++ {
                                            v3 -= 42
                                            var v184 int64 = -778
                                            var v185 uint16 = 0
                                            _ = fi183
                                            _ = v184
                                            _ = v185
                                        }
                                    case 1:
                                        if len(v1) > 0 {
                                            v1[0] = v5
                                        }
                                    }
                                }
                                _ = fi181
                            }
                            for fi186 := 0; fi186 < 19; fi186++ {
                                if (!(v4 != 0)) {
                                    v171 += v171
                                    v2 = (v2 ^ (v22 | (v22 * (v22 & 980))))
                                    v187 := GenS2{F0: -111}
                                    _ = v187
                                } else {
                                    for fi188 := 0; fi188 < 12; fi188++ {
                                        v22 -= uint(v4)
                                        _ = fi188
                                    }
                                    var v189 int = v171
                                    _ = v189
                                }
                                _ = len(v1)
                                _ = fi186
                            }
                        }
                        for fi190 := 0; fi190 < 18; fi190++ {
                            switch (v171 / (((v171 * (fi190 + (v171 | fi190)))) | 1)) {
                            case 4294967295:
                                if len(v1) > 0 {
                                    v1[0] = "\n\t"
                                }
                            }
                            _ = fi190
                        }
                    } else {
                        if len(v1) > 0 {
                            v1[0] = v5
                        }
                        switch ((v5 + (v5 + v5)) + v5) {
                        case "hello":
                            var v191 bool = (!true)
                            _ = v191
                        case "":
                            var v192 float32 = (float32(0.1) / (1.5 * 1.5))
                            _ = v192
                        default:
                            v193 := GenS13{F0: int(v4)}
                            _ = v193
                        }
                        v171 = 4294967295
                    }
                    switch v5 {
                    case "a":
                        if true {
                            v194 := GenS19{F0: 412, F1: (v3 ^ v3)}
                            _ = v194
                        } else {
                            if len(v1) > 0 {
                                v1[0] = v5
                            }
                        }
                    case "foo bar":
                        var v195 int64 = 9223372036854775807
                        _ = v195
                    case "\n\t":
                        for fi196 := 0; fi196 < 7; fi196++ {
                            fmt.Println(v5)
                            var v197 int64 = 910
                            v1 = append(v1, ((v5 + "foo bar") + ""))
                            _ = fi196
                            _ = v197
                        }
                    }
                } else {
                    var v198 int64 = 42
                    var v199 int64 = 9223372036854775807
                    _ = v198
                    _ = v199
                }
            }
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v200 := GenS20{F0: -1}
    var v201 uint32 = (v4 + v4)
    var v202 uint64 = 2
    _ = len(v1)
    if true {
        var v203 int32 = 0
        if true {
            for fi204 := 0; fi204 < 4; fi204++ {
                v171 = 9223372036854775807
                var v205 uint16 = 42
                switch 42 {
                case 656:
                    v171 = (fi204 >> ((v22) & 63))
                case -1:
                    v203 = v3
                case 1:
                    var v206 uint16 = (v205 | (^(v205)))
                    _ = v206
                }
                _ = fi204
                _ = v205
            }
            v5 = v5
            if false {
                for fi207 := 0; fi207 < 6; fi207++ {
                    v1 = append(v1, v5)
                    for fi208 := 0; fi208 < 13; fi208++ {
                        v209 := GenS21{F0: 42, F1: v3}
                        _ = fi208
                        _ = v209
                    }
                    for fi210 := 0; fi210 < 14; fi210++ {
                        if ((v4 + 2) >= (v4 << ((v22) & 63))) {
                            _ = len(v1)
                            var v211 uint32 = 2147483648
                            v212 := GenS22{F0: 42, F1: 42}
                            _ = v211
                            _ = v212
                        }
                        var v213 int8 = 127
                        if len(v1) > 0 {
                            v1[0] = ""
                        }
                        _ = fi210
                        _ = v213
                    }
                    _ = fi207
                }
                v2 -= (^(v2))
                fmt.Println(v171)
            } else {
                v171 = v171
                fmt.Println(v22, v3)
                v1 = append(v1, (v5 + ((v5 + (v5 + v5)) + v5)))
            }
        }
        for fi214 := 0; fi214 < 20; fi214++ {
            if true {
                if len(v1) > 0 {
                    v1[0] = "\n\t"
                }
                for fi215 := 0; fi215 < 13; fi215++ {
                    var v216 uint32 = v4
                    var v217 int8 = -1
                    _ = fi215
                    _ = v216
                    _ = v217
                }
            } else {
                v202 += (v202 & (v202 + v202))
            }
            _ = fi214
        }
        _ = v203
    }
    v218 := GenS23{F0: v5}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v171:", v171)
    fmt.Println("v2:", v2)
    fmt.Println("v201:", v201)
    fmt.Println("v202:", v202)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v171
    _ = v2
    _ = v200
    _ = v201
    _ = v202
    _ = v218
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
