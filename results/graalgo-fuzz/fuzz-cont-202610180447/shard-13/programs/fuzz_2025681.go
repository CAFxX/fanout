package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint64; F1 string; F2 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8; F1 uint64; F2 uint64 }
type GenS5 struct { F0 bool; F1 uint; F2 int16 }
type GenS6 struct { F0 uint64; F1 int }
type GenS7 struct { F0 uint32; F1 string }
type GenS8 struct { F0 int8; F1 int8 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint; F1 uint32; F2 int64 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 uint8; F1 int8; F2 int32 }
type GenS13 struct { F0 uint; F1 int8 }
type GenS14 struct { F0 uint64 }
type GenS15 struct { F0 int32 }
type GenS16 struct { F0 uint32 }

func main() {
    v1 := []int{1, 2147483647}
    v2 := []int{42}
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]uint32)
    for fi5 := 0; fi5 < 18; fi5++ {
        _ = len(v2)
        _ = fi5
    }
    v1 = append(v1, 9223372036854775807)
    if (42 >= 32767) {
        var v6 uint64 = 0
        switch "foo bar" {
        case "\n\t":
            v3[32768] = 2147483647
        case "a":
            v7 := GenS1{F0: 207, F1: 988}
            _ = v7
        case "hello":
            var v8 int8 = 1
            _ = v8
        default:
            v6 -= 18446744073709551615
        }
        v6 += (v6 - 2)
        _ = v6
    } else {
        for fi9 := 0; fi9 < 12; fi9++ {
            var v10 uint8 = 1
            if (-2.25 != (0.0 * ((((0.1 / -2.25) / ((3.14159 + (0.0 + 1e10)) * -2.25)) + 0.0) + 0.0))) {
                fmt.Println(v10)
            } else {
                var v11 float32 = ((1e10 - 0.1) - float32(0.1))
                v11 += (float32(0.0) * 3.14159)
                var v12 int32 = 0
                _ = v11
                _ = v12
            }
            switch ("hello" + ("hello" + ("\n\t" + "foo bar"))) {
            case "hello":
                for fi13 := 0; fi13 < 20; fi13++ {
                    fmt.Println(fi9)
                    _ = fi13
                }
            }
            _ = fi9
            _ = v10
        }
        _ = len(v4)
    }
    v14 := GenS1{F0: 89, F1: 42}
    switch "hello" {
    case "foo bar":
        v3[128] = 42
    case "a":
        var v15 uint64 = 42
        _ = v15
    case "hello":
        switch "" {
        case "foo bar":
            var v16 string = "a"
            _ = v16
        default:
            var v17 uint16 = 42
            _ = v17
        }
    }
    var v18 bool = true
    var v19 float64 = 0.0
    v20 := GenS2{F0: 2, F1: "foo bar", F2: 1}
    v19 += 0.0
    for fi21 := 0; fi21 < 3; fi21++ {
        for fi22 := 0; fi22 < 7; fi22++ {
            if (42 <= 42) {
                if v18 {
                    v23 := GenS2{F0: 42, F1: "a", F2: 1}
                    fmt.Println(v19, fi22)
                    _ = v23
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    fmt.Println(v19, fi22)
                }
            } else {
                v24 := GenS3{F0: 128}
                _ = v24
            }
            var v25 float64 = (((3.14159 * v19) * float64(1e10)) / float64((0.0 - -2.25)))
            if (0.1 > ((3.14159 - 3.14159) / 1.0)) {
                var v26 float32 = 1e10
                _ = v26
            }
            v19 = v25
            delete(v3, (fi22 + 9223372036854775807))
            switch (("\n\t" + "foo bar") + "\n\t") {
            case "foo bar":
                var v27 float64 = float64(((0.1 / ((0.1 + (((float32(v19) - 1.5) - 1.5) * 3.14159)) + 3.14159)) * 1.5))
                _ = v27
            }
            _ = fi22
            _ = v25
        }
        switch "hello" {
        case "":
            v28 := GenS3{F0: 127}
            _ = v28
        }
        _ = fi21
    }
    fmt.Println(v19, v19)
    for fi29 := 0; fi29 < 7; fi29++ {
        switch -2147483648 {
        case 32768:
            var v30 float32 = (1e10 - (-2.25 + 3.14159))
            _ = v30
        case 4294967295:
            fmt.Println(v18)
        }
        _ = fi29
    }
    for fi31 := 0; fi31 < 5; fi31++ {
        if (4294967295 > 1) {
            if (v19 == v19) {
                for fi32 := 0; fi32 < 14; fi32++ {
                    v18 = ((!(949 == 931)) || v18)
                    var v33 float32 = float32(((1e10 + v19) + v19))
                    _ = fi32
                    _ = v33
                }
                var v34 bool = v18
                if (42 == 0) {
                    var v35 int32 = 42
                    if (128 >= 0) {
                        v18 = v34
                        var v36 int32 = v35
                        v37 := GenS4{F0: 1, F1: 2, F2: 18446744073709551615}
                        _ = v36
                        _ = v37
                    } else {
                        var v38 int = 32768
                        if (1 < 140) {
                            if v18 {
                                for fi39 := 0; fi39 < 5; fi39++ {
                                    v40 := GenS5{F0: v18, F1: 672, F2: 0}
                                    _ = fi39
                                    _ = v40
                                }
                                if len(v2) > 0 {
                                    v2[0] = v38
                                }
                                var v41 float32 = (1e10 / 1.0)
                                _ = v41
                            } else {
                                if false {
                                    v18 = (2147483647 == 4294967295)
                                    var v42 uint32 = 560
                                    for fi43 := 0; fi43 < 19; fi43++ {
                                        delete(v4, 582)
                                        var v44 int64 = 731
                                        v42 = (^(v42))
                                        v45 := GenS4{F0: 128, F1: 18446744073709551615, F2: 1}
                                        _ = fi43
                                        _ = v44
                                        _ = v45
                                    }
                                    _ = v42
                                } else {
                                    if (int64(fi31) == 42) {
                                        if false {
                                            v46 := v4[42]
                                            _ = v46
                                        } else {
                                            var v47 bool = v18
                                            fmt.Println(v47, fi31)
                                            _ = v47
                                        }
                                        v19 = (1e10 + 1e10)
                                        fmt.Println(v18, v19, v34)
                                    } else {
                                        if v34 {
                                            for fi48 := 0; fi48 < 6; fi48++ {
                                                switch "foo bar" {
                                                case "":
                                                    if (v34 && false) {
                                                        if v18 {
                                                            switch v38 {
                                                            case -210:
                                                                if v18 {
                                                                    var v49 int8 = -1
                                                                    if (42 <= 1) {
                                                                        v19 -= 0.0
                                                                        if ((v18 && v18) && v34) {
                                                                            _ = len(v2)
                                                                            fmt.Println(v18)
                                                                        } else {
                                                                            if (v19 > (float64((0.0 / 1.5)) + (1e10 - ((1.5 - 0.0) - 0.1)))) {
                                                                                var v50 int32 = (v35 | v35)
                                                                                var v51 uint8 = 42
                                                                                var v52 int64 = 9223372036854775807
                                                                                _ = v50
                                                                                _ = v51
                                                                                _ = v52
                                                                            } else {
                                                                                v53 := GenS1{F0: 127, F1: 2}
                                                                                if v18 {
                                                                                    fmt.Println(fi31)
                                                                                    var v54 int32 = (v35 / ((-2147483648) | 1))
                                                                                    fmt.Println(v54, v54, v35)
                                                                                    var v55 uint8 = 42
                                                                                    var v56 float32 = 0.0
                                                                                    if (!((((v19 - -2.25) <= (0.1 - (v19 + (1e10 - 1.5)))) || true) || true)) {
                                                                                        if (!v34) {
                                                                                            switch (fi31 & -106) {
                                                                                            case -801:
                                                                                                if true {
                                                                                                    if v18 {
                                                                                                        if v18 {
                                                                                                            switch "hello" {
                                                                                                            case "a":
                                                                                                                v19 -= v19
                                                                                                            case "\n\t":
                                                                                                                var v57 int8 = 42
                                                                                                                _ = v57
                                                                                                            }
                                                                                                            fmt.Println(v55)
                                                                                                            v35 -= v35
                                                                                                            _ = len(v2)
                                                                                                            switch (^(v38)) {
                                                                                                            case 901:
                                                                                                                delete(v3, (fi48 / (((fi48 << ((948) & 63))) | 1)))
                                                                                                            case -425:
                                                                                                                if (!v34) {
                                                                                                                    _ = len(v4)
                                                                                                                    if len(v2) >= 2 {
                                                                                                                        v2 = v2[1:len(v2)]
                                                                                                                    }
                                                                                                                    v55 = (v55 ^ v55)
                                                                                                                } else {
                                                                                                                    var v58 bool = (-716 <= 42)
                                                                                                                    switch ((("a" + "foo bar") + "\n\t") + "a") {
                                                                                                                    case "":
                                                                                                                        _ = len(v2)
                                                                                                                    case "hello":
                                                                                                                        v58 = true
                                                                                                                    }
                                                                                                                    if len(v2) >= 2 {
                                                                                                                        v2 = v2[1:len(v2)]
                                                                                                                    }
                                                                                                                    _ = v58
                                                                                                                }
                                                                                                            default:
                                                                                                                v49 = 1
                                                                                                            }
                                                                                                        } else {
                                                                                                            if len(v2) >= 2 {
                                                                                                                v2 = v2[1:len(v2)]
                                                                                                            }
                                                                                                        }
                                                                                                        v55 += v55
                                                                                                        if v34 {
                                                                                                            v19 = v19
                                                                                                            var v59 uint64 = 2
                                                                                                            if (!false) {
                                                                                                                if ("" <= "foo bar") {
                                                                                                                    var v60 string = "hello"
                                                                                                                    var v61 uint8 = v55
                                                                                                                    if true {
                                                                                                                        fmt.Println(v55, fi31)
                                                                                                                    } else {
                                                                                                                        var v62 uint8 = v55
                                                                                                                        var v63 string = "\n\t"
                                                                                                                        _ = v62
                                                                                                                        _ = v63
                                                                                                                    }
                                                                                                                    _ = v60
                                                                                                                    _ = v61
                                                                                                                } else {
                                                                                                                    if true {
                                                                                                                        v59 -= 1
                                                                                                                        v64 := GenS6{F0: 1, F1: (v38 & (fi48 + 42))}
                                                                                                                        fmt.Println(v55)
                                                                                                                        if (128 >= fi48) {
                                                                                                                            v65 := GenS4{F0: v55, F1: 185, F2: (^(v59))}
                                                                                                                            _ = v65
                                                                                                                        } else {
                                                                                                                            v35 = -994
                                                                                                                            v55 = v55
                                                                                                                            v66 := GenS7{F0: 1, F1: ((("" + ("a" + "\n\t")) + "foo bar") + (("foo bar" + ("\n\t" + "a")) + ("\n\t" + "hello")))}
                                                                                                                            _ = v66
                                                                                                                        }
                                                                                                                        v67 := GenS4{F0: 255, F1: 4294967295, F2: (^(v59))}
                                                                                                                        delete(v3, v38)
                                                                                                                        var v68 int8 = 100
                                                                                                                        _ = v64
                                                                                                                        _ = v67
                                                                                                                        _ = v68
                                                                                                                    } else {
                                                                                                                        _ = len(v3)
                                                                                                                        v55 = uint8(v59)
                                                                                                                    }
                                                                                                                }
                                                                                                            } else {
                                                                                                                fmt.Println(v49, v49, fi48)
                                                                                                                var v69 int = v38
                                                                                                                _ = v69
                                                                                                            }
                                                                                                            _ = v59
                                                                                                        } else {
                                                                                                            var v70 int64 = -9223372036854775808
                                                                                                            _ = v70
                                                                                                        }
                                                                                                    } else {
                                                                                                        if len(v1) > 0 {
                                                                                                            v1[0] = int(v49)
                                                                                                        }
                                                                                                        var v71 int32 = v54
                                                                                                        _ = v71
                                                                                                    }
                                                                                                } else {
                                                                                                    var v72 int8 = -128
                                                                                                    if ((1 >= 0) || ((v54 / (((v54 ^ v35)) | 1)) >= (v54 << ((4294967295) & 63)))) {
                                                                                                        fmt.Println(v19, v35, v19)
                                                                                                    }
                                                                                                    v38 += -2147483648
                                                                                                    _ = v72
                                                                                                }
                                                                                            case -632:
                                                                                                if true {
                                                                                                    var v73 int16 = int16(fi31)
                                                                                                    _ = v73
                                                                                                }
                                                                                            case 128:
                                                                                                if true {
                                                                                                    var v74 int8 = v49
                                                                                                    _ = v74
                                                                                                }
                                                                                            }
                                                                                            v75 := GenS8{F0: (v49 * 42), F1: (-(v49))}
                                                                                            fmt.Println(v35)
                                                                                            _ = v75
                                                                                        }
                                                                                        if (610 > 18446744073709551615) {
                                                                                            v18 = ((128 <= v55) || true)
                                                                                            if ((v38 >> ((65535) & 63)) >= v38) {
                                                                                                var v76 uint8 = (v55 & 1)
                                                                                                _ = v76
                                                                                            } else {
                                                                                                if v34 {
                                                                                                    var v77 bool = ((v38 * (fi48 << ((65535) & 63))) != (v38 - v38))
                                                                                                    v78 := GenS2{F0: 18446744073709551615, F1: "", F2: fi48}
                                                                                                    _ = v77
                                                                                                    _ = v78
                                                                                                }
                                                                                                switch "" {
                                                                                                case "hello":
                                                                                                    if len(v2) > 0 {
                                                                                                        v2[0] = v38
                                                                                                    }
                                                                                                case "":
                                                                                                    fmt.Println(v54, v38)
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        fmt.Println(v18)
                                                                                    } else {
                                                                                        var v79 uint16 = 2
                                                                                        v18 = (v35 != (v35 / ((int32(v38)) | 1)))
                                                                                        _ = v79
                                                                                    }
                                                                                    v35 = v54
                                                                                    if v18 {
                                                                                        v54 = 42
                                                                                        fmt.Println(v38, v35, v54)
                                                                                    }
                                                                                    v80 := GenS9{F0: 2}
                                                                                    fmt.Println(v54)
                                                                                    _ = v54
                                                                                    _ = v55
                                                                                    _ = v56
                                                                                    _ = v80
                                                                                } else {
                                                                                    var v81 int32 = (^(v35))
                                                                                    v34 = v34
                                                                                    v38 = 128
                                                                                    v81 = -1
                                                                                    fmt.Println(v49, v35, v49)
                                                                                    _ = v81
                                                                                }
                                                                                var v82 string = ("" + (((("\n\t" + ("\n\t" + "a")) + "") + "\n\t") + (("hello" + ("\n\t" + "a")) + "")))
                                                                                _ = v53
                                                                                _ = v82
                                                                            }
                                                                            switch (("hello" + "") + "\n\t") {
                                                                            case "":
                                                                                v49 += v49
                                                                            default:
                                                                                v83 := GenS2{F0: 1, F1: (("hello" + "") + "hello"), F2: fi31}
                                                                                _ = v83
                                                                            }
                                                                        }
                                                                        _ = len(v2)
                                                                    } else {
                                                                        switch "" {
                                                                        case "hello":
                                                                            _ = len(v1)
                                                                        case "":
                                                                            switch fi48 {
                                                                            case -1:
                                                                                if (!v18) {
                                                                                    var v84 string = (((("a" + ((((("hello" + "") + (("\n\t" + "a") + ("a" + ("hello" + (((("a" + ("a" + (("a" + "hello") + "foo bar"))) + "hello") + "hello") + "foo bar"))))) + ("" + "foo bar")) + "foo bar") + "a")) + "") + ((("\n\t" + "") + "a") + ("hello" + "hello"))) + "a")
                                                                                    var v85 uint16 = 42
                                                                                    v84 = ((("hello" + "\n\t") + v84) + "hello")
                                                                                    _ = v84
                                                                                    _ = v85
                                                                                } else {
                                                                                    v38 = fi31
                                                                                }
                                                                            default:
                                                                                fmt.Println(v34, v38)
                                                                            }
                                                                        }
                                                                        fmt.Println(v34)
                                                                        var v86 int32 = v35
                                                                        _ = v86
                                                                    }
                                                                    _ = v49
                                                                }
                                                            case 1:
                                                                v18 = v34
                                                            }
                                                            v34 = (!false)
                                                            fmt.Println(fi31, v38)
                                                        } else {
                                                            if (0.1 != 0.0) {
                                                                v35 -= (v35 | (^(v35)))
                                                            }
                                                            v34 = v18
                                                        }
                                                    }
                                                    v87 := GenS10{F0: 539, F1: 1, F2: 9223372036854775807}
                                                    fmt.Println(v18, v38)
                                                    v34 = (false || (!false))
                                                    _ = v87
                                                }
                                                _ = fi48
                                            }
                                        } else {
                                            for fi88 := 0; fi88 < 16; fi88++ {
                                                _ = len(v1)
                                                v89 := v4[575]
                                                if (!false) {
                                                    var v90 int16 = -32768
                                                    switch ("a" + "foo bar") {
                                                    case "hello":
                                                        switch ("\n\t" + ("" + "foo bar")) {
                                                        case "a":
                                                            v90 = v90
                                                        default:
                                                            var v91 float64 = v19
                                                            _ = v91
                                                        }
                                                    case "a":
                                                        if (true && (((true || (!(!("foo bar" > (("foo bar" + "hello") + "hello"))))) || v18) && true)) {
                                                            switch "a" {
                                                            case "a":
                                                                v92 := GenS11{F0: (("hello" + ("foo bar" + "a")) + "")}
                                                                _ = v92
                                                            case "hello":
                                                                if true {
                                                                    v90 = (v90 % (((v90 & v90)) | 1))
                                                                } else {
                                                                    var v93 int64 = 528
                                                                    v94 := GenS6{F0: 593, F1: v38}
                                                                    switch v38 {
                                                                    case 0:
                                                                        fmt.Println(fi88, v38, v90)
                                                                        var v95 uint16 = 1
                                                                        _ = v95
                                                                    case -2147483648:
                                                                        if true {
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                            v35 = (v35 % (((v35 - (v35 * (-(v35))))) | 1))
                                                                            delete(v4, (v89 % (((v89 / ((v89) | 1))) | 1)))
                                                                        }
                                                                    case 1:
                                                                        if false {
                                                                            v1 = append(v1, int(v90))
                                                                        } else {
                                                                            v96 := GenS10{F0: 4294967295, F1: (v89 / ((v89) | 1)), F2: v93}
                                                                            v97 := GenS9{F0: 824}
                                                                            _ = len(v4)
                                                                            _ = v96
                                                                            _ = v97
                                                                        }
                                                                    }
                                                                    _ = v93
                                                                    _ = v94
                                                                }
                                                            case "":
                                                                var v98 uint16 = 855
                                                                _ = v98
                                                            default:
                                                                v99 := GenS7{F0: 1, F1: "hello"}
                                                                _ = v99
                                                            }
                                                            var v100 int8 = -128
                                                            _ = v100
                                                        }
                                                    case "\n\t":
                                                        fmt.Println(fi31)
                                                    default:
                                                        var v101 string = "a"
                                                        _ = v101
                                                    }
                                                    fmt.Println(fi31, fi31, v34)
                                                    _ = v90
                                                }
                                                v89 = v89
                                                v102 := GenS12{F0: 128, F1: -1, F2: v35}
                                                _ = fi88
                                                _ = v102
                                                _ = v89
                                            }
                                            v38 = v38
                                        }
                                    }
                                    if v18 {
                                        var v103 int8 = 0
                                        _ = v103
                                    }
                                    v35 -= v35
                                }
                                switch ("" + "hello") {
                                case "\n\t":
                                    if (-115 > 42) {
                                        if ((v18 || v18) && v34) {
                                            for fi104 := 0; fi104 < 7; fi104++ {
                                                var v105 uint64 = 42
                                                _ = fi104
                                                _ = v105
                                            }
                                            v106 := GenS3{F0: 1}
                                            _ = v106
                                        } else {
                                            v107 := v4[899]
                                            if ("hello" > "hello") {
                                                v107 += (v107 ^ 820)
                                                var v108 int32 = 0
                                                var v109 uint8 = 127
                                                _ = v108
                                                _ = v109
                                            }
                                            v110 := GenS13{F0: 65535, F1: int8(v107)}
                                            _ = v107
                                            _ = v110
                                        }
                                    }
                                }
                                for fi111 := 0; fi111 < 6; fi111++ {
                                    v18 = true
                                    var v112 uint64 = 474
                                    switch (v38 % (((fi111 / ((fi31) | 1))) | 1)) {
                                    case 602:
                                        v18 = v34
                                    case 1:
                                        var v113 uint32 = 42
                                        _ = v113
                                    case 1000:
                                        var v114 int32 = v35
                                        if ((!((184 <= 9223372036854775807) && (!v34))) && v34) {
                                            v115 := GenS4{F0: 1, F1: (v112 >> ((255) & 63)), F2: 4294967295}
                                            v38 = -34
                                            _ = v115
                                        }
                                        var v116 int64 = -266
                                        _ = v114
                                        _ = v116
                                    }
                                    _ = fi111
                                    _ = v112
                                }
                            }
                        } else {
                            for fi117 := 0; fi117 < 14; fi117++ {
                                v34 = false
                                v34 = v34
                                _ = fi117
                            }
                            v38 = (-(v38))
                        }
                        _ = v38
                    }
                    if false {
                        v34 = v18
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    } else {
                        var v118 int64 = -9223372036854775808
                        _ = v118
                    }
                    _ = v35
                } else {
                    if (v19 > (((-2.25 + v19) / ((v19 - (v19 / (3.14159 / (v19 * (v19 * 0.1))))) * ((v19 - 0.1) / 1.0))) * 1.5)) {
                        var v119 int64 = 9223372036854775807
                        _ = v119
                    }
                    fmt.Println(v18)
                    v3[fi31] = 1
                }
                _ = v34
            } else {
                for fi120 := 0; fi120 < 16; fi120++ {
                    if ((1.5 - -2.25) <= -2.25) {
                        v19 -= (((0.0 * (float64(1.5) / ((v19 + v19) + -2.25))) / 1e10) / 1.0)
                        switch "" {
                        case "\n\t":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        case "":
                            var v121 uint8 = 255
                            _ = v121
                        case "a":
                            fmt.Println(fi120)
                        }
                    }
                    var v122 float32 = 3.14159
                    _ = fi120
                    _ = v122
                }
                v4[2] = 2147483647
                var v123 int16 = 1000
                _ = v123
            }
            if true {
                v18 = (2147483647 < 1)
            }
        } else {
            v19 = v19
            switch (((("foo bar" + "\n\t") + "a") + "hello") + "foo bar") {
            case "hello":
                fmt.Println(v18, v18, v18)
            case "foo bar":
                if (1e10 < 3.14159) {
                    v19 += (0.0 + 3.14159)
                }
            case "a":
                if (2 <= 811) {
                    v18 = (false || v18)
                    v124 := GenS10{F0: 822, F1: 65535, F2: 617}
                    v125 := GenS3{F0: 2}
                    _ = v124
                    _ = v125
                }
            }
            v126 := GenS14{F0: 18446744073709551615}
            _ = v126
        }
        _ = fi31
    }
    v18 = (1000 == 1)
    switch ("foo bar" + "") {
    case "\n\t":
        var v127 uint32 = 772
        _ = v127
    case "a":
        if v18 {
            switch "" {
            case "hello":
                if (9223372036854775807 <= -9223372036854775808) {
                    var v128 uint = 730
                    _ = v128
                } else {
                    v18 = (-2147483648 >= -9223372036854775808)
                }
            }
            switch -790 {
            case 9223372036854775807:
                for fi129 := 0; fi129 < 14; fi129++ {
                    for fi130 := 0; fi130 < 3; fi130++ {
                        if true {
                            fmt.Println(v18)
                            v131 := GenS11{F0: ""}
                            var v132 float32 = 3.14159
                            _ = v131
                            _ = v132
                        } else {
                            var v133 int32 = 638
                            _ = v133
                        }
                        _ = fi130
                    }
                    v19 = v19
                    v19 = v19
                    _ = fi129
                }
            default:
                v18 = ("hello" == "\n\t")
            }
            var v134 uint32 = 2147483647
            _ = v134
        }
    case "foo bar":
        if v18 {
            switch "foo bar" {
            case "\n\t":
                for fi135 := 0; fi135 < 7; fi135++ {
                    var v136 uint64 = 918
                    fmt.Println(fi135)
                    v19 = v19
                    _ = fi135
                    _ = v136
                }
            default:
                var v137 bool = ((3.14159 <= (((0.1 - float32((float64(float32(float64(float32(1e10)))) * (float64((3.14159 + 1.5)) * ((float64(1e10) * v19) / v19))))) + ((0.0 - -2.25) + 0.0)) + ((0.1 + 1.5) - 3.14159))) && v18)
                _ = v137
            }
            v138 := GenS15{F0: 1}
            for fi139 := 0; fi139 < 3; fi139++ {
                var v140 uint8 = 128
                v141 := GenS15{F0: 1}
                for fi142 := 0; fi142 < 14; fi142++ {
                    fmt.Println(fi139, fi139, v19)
                    v143 := GenS2{F0: 290, F1: "\n\t", F2: (^(fi139))}
                    if v18 {
                        v140 = (v140 | v140)
                        var v144 uint8 = v140
                        v145 := GenS16{F0: 42}
                        _ = v144
                        _ = v145
                    } else {
                        v140 -= 255
                        var v146 bool = (1 == -32768)
                        _ = v146
                    }
                    _ = fi142
                    _ = v143
                }
                _ = fi139
                _ = v140
                _ = v141
            }
            _ = v138
        } else {
            fmt.Println(v18)
        }
    default:
        var v147 int64 = -1
        _ = v147
    }
    v148 := GenS13{F0: 42, F1: -1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v148
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
