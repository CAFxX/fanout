package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 int; F1 int8; F2 int64 }
type GenS4 struct { F0 int32; F1 int64; F2 string }
type GenS5 struct { F0 uint; F1 int; F2 int64 }
type GenS6 struct { F0 string; F1 uint32; F2 uint }
type GenS7 struct { F0 uint16; F1 int }
type GenS8 struct { F0 uint8; F1 int; F2 int8 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int; F1 int }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 uint64; F1 uint32; F2 uint16 }
type GenS13 struct { F0 int16; F1 int }
type GenS14 struct { F0 int64; F1 int32 }
type GenS15 struct { F0 bool; F1 string }
type GenS16 struct { F0 uint8; F1 uint; F2 uint32 }
type GenS17 struct { F0 int64 }
type GenS18 struct { F0 uint; F1 int }
type GenS19 struct { F0 uint32 }
type GenS20 struct { F0 bool }
type GenS21 struct { F0 int32; F1 string; F2 uint32 }
func genFn0(p0 bool) int32 {
    for fi6 := 0; fi6 < 2; fi6++ {
        if p0 {
            p0 = p0
        }
        p0 = (0.0 <= 0.0)
        var v7 float32 = 0.1
        _ = fi6
        _ = v7
    }
    p0 = (-327 < -273)
    if p0 {
        if true {
            for fi8 := 0; fi8 < 3; fi8++ {
                v9 := []int{(fi8 ^ (fi8 << ((65535) & 63)))}
                v10 := GenS1{F0: 65535, F1: "\n\t", F2: fi8}
                var v11 string = "a"
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
        } else {
            v12 := GenS1{F0: 1, F1: ("a" + "\n\t"), F2: -1}
            _ = v12
        }
        var v13 uint = 302
        var v14 int32 = -1
        _ = v13
        _ = v14
    }
    p0 = p0
    _ = p0
    return 1
}
func genFn1(p0 int32, p1 int8) uint {
    v15 := []int{32767, 32767}
    if (42 <= -473) {
        switch 128 {
        case 230:
            var v16 uint8 = 0
            _ = v16
        default:
            if true {
                for fi17 := 0; fi17 < 13; fi17++ {
                    var v18 uint8 = 255
                    _ = fi17
                    _ = v18
                }
                v19 := GenS1{F0: 18446744073709551615, F1: "\n\t", F2: 32768}
                v20 := GenS1{F0: 1, F1: "", F2: 32767}
                _ = v19
                _ = v20
            }
        }
    }
    v21 := GenS1{F0: 232, F1: ("a" + ("\n\t" + "foo bar")), F2: int(p1)}
    _ = p0
    _ = p1
    _ = v15
    _ = v21
    return 255
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]bool)
    v3 := []uint32{65535}
    var v4 bool = false
    v5 := make(map[string]string)
    if (128 > 10) {
        v22 := GenS1{F0: genFn1(1, -128), F1: "hello", F2: -764}
        v23 := GenS1{F0: 65535, F1: ("\n\t" + (("" + ("" + ("foo bar" + "a"))) + "hello")), F2: 86}
        fmt.Println(v4)
        _ = v22
        _ = v23
    }
    if v4 {
        fmt.Println(v4)
        v4 = (-2147483648 >= 2147483647)
    }
    if (42 >= 65535) {
        var v24 float32 = 1.5
        v4 = v4
        switch -956 {
        case 99:
            if (v4 || true) {
                var v25 uint = 379
                delete(v2, "")
                for fi26 := 0; fi26 < 4; fi26++ {
                    v27 := GenS1{F0: (v25 << (((v25 ^ 18446744073709551615)) & 63)), F1: (("\n\t" + ("" + "\n\t")) + "hello"), F2: fi26}
                    v4 = false
                    if v4 {
                        v24 -= (0.1 * (v24 * ((((float32(1.5) / ((v24 + (v24 + v24)) - v24)) / v24) * (float32(0.0) * v24)) * float32(float64(v24)))))
                    }
                    _ = fi26
                    _ = v27
                }
                _ = v25
            }
        }
        _ = v24
    } else {
        var v28 int8 = -1
        _ = v28
    }
    var v29 int32 = -1
    var v30 bool = (255 != 1)
    if (false && (true && v30)) {
        switch 0 {
        case 128:
            var v31 bool = (false || (1 < 2))
            _ = v31
        case -9223372036854775808:
            delete(v2, "")
        case 4294967295:
            v30 = (42 < 0)
        default:
            switch ("\n\t" + "foo bar") {
            case "a":
                v32 := GenS1{F0: 2, F1: "a", F2: 127}
                _ = v32
            case "foo bar":
                v33 := GenS1{F0: 744, F1: "\n\t", F2: 128}
                _ = v33
            }
        }
        v30 = (!(207 >= 0))
        if v30 {
            var v34 uint32 = 0
            v35 := GenS1{F0: 709, F1: "\n\t", F2: 32767}
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v34
            _ = v35
        }
    }
    if v30 {
        var v36 uint = 0
        _ = len(v5)
        _ = v36
    }
    for fi37 := 0; fi37 < 13; fi37++ {
        v29 = (v29 / ((v29) | 1))
        var v38 float32 = (float32((((1e10 * ((3.14159 + (0.0 * 0.1)) / 1.5)) * (float64(0.1) * float64(((-2.25 - (0.1 / 3.14159)) / 3.14159)))) * (-2.25 + 0.1))) * 0.0)
        switch -1 {
        case -640:
            var v39 uint = 1
            _ = v39
        case 2147483647:
            switch (fi37 * fi37) {
            case 844:
                v40 := GenS1{F0: 4294967295, F1: ("foo bar" + ""), F2: (-(fi37))}
                _ = v40
            case 4294967295:
                for fi41 := 0; fi41 < 6; fi41++ {
                    v38 += v38
                    _ = fi41
                }
            case 2147483647:
                for fi42 := 0; fi42 < 18; fi42++ {
                    v29 = -2147483648
                    _ = fi42
                }
            }
        case -9223372036854775808:
            for fi43 := 0; fi43 < 8; fi43++ {
                if len(v3) > 0 {
                    v3[0] = 4294967295
                }
                if v4 {
                    var v44 uint64 = 18446744073709551615
                    var v45 int = (fi43 - fi43)
                    if v30 {
                        if (!(662 <= 379)) {
                            var v46 uint32 = 1
                            switch ("" + ("\n\t" + ("" + "\n\t"))) {
                            case "":
                                switch (-(fi43)) {
                                case 9223372036854775807:
                                    var v47 int16 = -924
                                    _ = v47
                                case 587:
                                    v48 := GenS1{F0: genFn1(v29, 1), F1: "hello", F2: (fi43 ^ (fi43 | fi43))}
                                    _ = v48
                                case 4294967295:
                                    if v30 {
                                        switch "a" {
                                        case "foo bar":
                                            var v49 int64 = 465
                                            _ = v49
                                        case "a":
                                            v30 = v4
                                        case "\n\t":
                                            switch "a" {
                                            case "\n\t":
                                                v44 -= (v44 + 911)
                                            case "a":
                                                var v50 uint = 0
                                                _ = v50
                                            case "":
                                                v44 = (v44 / (((v44 / (((v44 + v44)) | 1))) | 1))
                                            }
                                        }
                                        var v51 uint16 = 65535
                                        var v52 uint = 711
                                        _ = v51
                                        _ = v52
                                    } else {
                                        var v53 uint32 = 716
                                        if (0.1 >= v38) {
                                            var v54 int = 7
                                            _ = v54
                                        }
                                        _ = v53
                                    }
                                default:
                                    var v55 int32 = 414
                                    _ = v55
                                }
                            case "\n\t":
                                v56 := GenS1{F0: 255, F1: "a", F2: (v45 / ((fi37) | 1))}
                                _ = v56
                            case "foo bar":
                                v30 = (v44 <= v44)
                            default:
                                fmt.Println(v29, v45, v38)
                            }
                            v4 = false
                            switch v45 {
                            case 32767:
                                var v57 float32 = v38
                                _ = v57
                            }
                            switch (^(fi43)) {
                            case 32768:
                                var v58 int64 = 1
                                _ = v58
                            case -2147483648:
                                var v59 int = (^(fi43))
                                _ = v59
                            }
                            _ = v46
                        } else {
                            fmt.Println(fi37, fi43)
                            if ((482 >= 18446744073709551615) || v4) {
                                if (18446744073709551615 != 1) {
                                    switch ("\n\t" + "a") {
                                    case "hello":
                                        var v60 float32 = v38
                                        _ = v60
                                    case "\n\t":
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    case "a":
                                        _ = len(v3)
                                    default:
                                        var v61 uint64 = (v44 % ((v44) | 1))
                                        _ = v61
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    v3 = append(v3, 2147483648)
                                } else {
                                    v30 = (202 > -2147483648)
                                    v45 -= v45
                                    if v4 {
                                        v62 := GenS1{F0: 255, F1: "hello", F2: (v45 & (v45 & (fi37 << ((779) & 63))))}
                                        _ = v62
                                    } else {
                                        fmt.Println(v4, v4, v4)
                                        if v30 {
                                            delete(v5, (("foo bar" + (("\n\t" + "\n\t") + "\n\t")) + "\n\t"))
                                            if (v4 || v4) {
                                                v30 = v30
                                                if (!v30) {
                                                    switch 993 {
                                                    case -529:
                                                        var v63 float64 = (1.5 * (1e10 + 3.14159))
                                                        _ = v63
                                                    case 459:
                                                        var v64 int16 = 946
                                                        _ = v64
                                                    case 128:
                                                        if v4 {
                                                            fmt.Println(v30, v29)
                                                        }
                                                    }
                                                    v44 += v44
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    if v4 {
                                                        v65 := GenS1{F0: 4294967295, F1: "foo bar", F2: 42}
                                                        _ = v65
                                                    } else {
                                                        var v66 uint32 = 1
                                                        fmt.Println(fi43, v30)
                                                        var v67 int = (fi43 ^ v45)
                                                        _ = v66
                                                        _ = v67
                                                    }
                                                    var v68 uint32 = 250
                                                    var v69 int16 = 1
                                                    _ = v68
                                                    _ = v69
                                                }
                                                if ((v29 + genFn0((32767 > 142))) < genFn0((0 > 9223372036854775807))) {
                                                    if v4 {
                                                        var v70 int16 = -32768
                                                        _ = v70
                                                    } else {
                                                        if v4 {
                                                            switch fi43 {
                                                            case -453:
                                                                v30 = (v29 >= v29)
                                                            case -527:
                                                                v71 := GenS2{F0: 47, F1: 42}
                                                                _ = v71
                                                            case -2147483648:
                                                                var v72 int32 = (v29 * v29)
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                var v73 int64 = 490
                                                                _ = v72
                                                                _ = v73
                                                            default:
                                                                v30 = (!(2147483647 <= v29))
                                                            }
                                                        } else {
                                                            v74 := GenS3{F0: 128, F1: -52, F2: -9223372036854775808}
                                                            v4 = (v4 && (("a" + ("a" + "\n\t")) > "foo bar"))
                                                            var v75 bool = false
                                                            var v76 uint32 = 73
                                                            _ = v74
                                                            _ = v75
                                                            _ = v76
                                                        }
                                                        switch ("a" + "foo bar") {
                                                        case "":
                                                            var v77 float32 = float32((3.14159 * 1e10))
                                                            _ = v77
                                                        case "a":
                                                            if v4 {
                                                                v45 = (v45 % ((fi37) | 1))
                                                                switch (("\n\t" + "hello") + "\n\t") {
                                                                case "\n\t":
                                                                    if ((2 >= 2) || true) {
                                                                        if v30 {
                                                                            v78 := GenS4{F0: int32(v44), F1: -880, F2: "hello"}
                                                                            _ = v78
                                                                        }
                                                                        v79 := GenS5{F0: 4294967295, F1: fi43, F2: 0}
                                                                        switch (fi37 * (fi37 + (fi43 * 443))) {
                                                                        case 32768:
                                                                            if v4 {
                                                                                v44 = v44
                                                                                v80 := GenS6{F0: "hello", F1: 4294967295, F2: 4294967295}
                                                                                delete(v2, "\n\t")
                                                                                var v81 uint64 = v44
                                                                                switch "" {
                                                                                case "\n\t":
                                                                                    var v82 uint16 = 42
                                                                                    _ = v82
                                                                                default:
                                                                                    switch (^(v45)) {
                                                                                    case -872:
                                                                                        if true {
                                                                                            v83 := GenS3{F0: (v45 / ((fi43) | 1)), F1: 100, F2: 9223372036854775807}
                                                                                            v30 = v30
                                                                                            _ = v83
                                                                                        }
                                                                                    default:
                                                                                        switch v45 {
                                                                                        case -174:
                                                                                            v84 := GenS5{F0: genFn1(1, -89), F1: (-(fi37)), F2: -147}
                                                                                            _ = v84
                                                                                        case -414:
                                                                                            if v4 {
                                                                                                switch "foo bar" {
                                                                                                case "":
                                                                                                    v44 -= v81
                                                                                                case "\n\t":
                                                                                                    switch -92 {
                                                                                                    case 2147483647:
                                                                                                        switch "foo bar" {
                                                                                                        case "hello":
                                                                                                            if v30 {
                                                                                                                _ = len(v3)
                                                                                                            }
                                                                                                        case "\n\t":
                                                                                                            if true {
                                                                                                                v44 = v44
                                                                                                                v44 = v81
                                                                                                                v85 := GenS7{F0: 918, F1: -2147483648}
                                                                                                                _ = v85
                                                                                                            }
                                                                                                        }
                                                                                                    case -821:
                                                                                                        var v86 float64 = float64(-2.25)
                                                                                                        _ = v86
                                                                                                    default:
                                                                                                        v45 = (^(v45))
                                                                                                    }
                                                                                                case "foo bar":
                                                                                                    v44 += (v81 | v44)
                                                                                                default:
                                                                                                    var v87 int32 = (v29 / (((v29 | 305)) | 1))
                                                                                                    _ = v87
                                                                                                }
                                                                                            }
                                                                                        case 593:
                                                                                            if true {
                                                                                                var v88 uint64 = v44
                                                                                                if v30 {
                                                                                                    var v89 float64 = -2.25
                                                                                                    _ = v89
                                                                                                }
                                                                                                var v90 string = ("hello" + ("a" + "foo bar"))
                                                                                                _ = v88
                                                                                                _ = v90
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    if v30 {
                                                                                        v44 += 2
                                                                                        var v91 float32 = float32((float64((1.5 / 3.14159)) + (0.1 * float64(float32(1e10)))))
                                                                                        _ = v91
                                                                                    } else {
                                                                                        v45 = -158
                                                                                        var v92 uint8 = 255
                                                                                        switch v45 {
                                                                                        case 540:
                                                                                            if v30 {
                                                                                                if v4 {
                                                                                                    switch fi43 {
                                                                                                    case -9223372036854775808:
                                                                                                        v93 := GenS8{F0: (v92 >> ((18446744073709551615) & 63)), F1: (fi37 / ((fi43) | 1)), F2: 1}
                                                                                                        _ = v93
                                                                                                    case -807:
                                                                                                        _ = len(v3)
                                                                                                    }
                                                                                                } else {
                                                                                                    _ = len(v3)
                                                                                                    fmt.Println(v92)
                                                                                                    fmt.Println(v29)
                                                                                                }
                                                                                                v4 = v30
                                                                                            } else {
                                                                                                var v94 bool = ((9223372036854775807 <= -1) && v4)
                                                                                                _ = v94
                                                                                            }
                                                                                        case 32768:
                                                                                            var v95 string = ""
                                                                                            _ = v95
                                                                                        default:
                                                                                            if (v30 || v30) {
                                                                                                switch "hello" {
                                                                                                case "":
                                                                                                    v29 -= (v29 + genFn0((-9223372036854775808 > -9223372036854775808)))
                                                                                                case "foo bar":
                                                                                                    v38 += v38
                                                                                                    v96 := GenS9{F0: 972}
                                                                                                    _ = v96
                                                                                                default:
                                                                                                    if v30 {
                                                                                                        v44 -= (v44 << ((18446744073709551615) & 63))
                                                                                                        fmt.Println(v30, v30)
                                                                                                    }
                                                                                                }
                                                                                                var v97 int = (-(fi37))
                                                                                                if ((true && true) || (v92 == v92)) {
                                                                                                    v2["hello"] = (!(v44 != 164))
                                                                                                }
                                                                                                _ = v97
                                                                                            }
                                                                                        }
                                                                                        _ = v92
                                                                                    }
                                                                                    v29 = v29
                                                                                    if len(v3) > 0 {
                                                                                        v3[0] = 42
                                                                                    }
                                                                                }
                                                                                var v98 string = ("hello" + "a")
                                                                                _ = v80
                                                                                _ = v81
                                                                                _ = v98
                                                                            }
                                                                        default:
                                                                            v29 = (^(v29))
                                                                        }
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                        v44 = (v44 % ((2) | 1))
                                                                        _ = v79
                                                                    }
                                                                case "":
                                                                    switch fi37 {
                                                                    case 127:
                                                                        fmt.Println(v4, v38, v29)
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    }
                                                    var v99 uint = 65535
                                                    var v100 float32 = float32(1.5)
                                                    _ = v100
                                                    _ = v99
                                                }
                                                if false {
                                                    var v101 uint = 65535
                                                    var v102 uint32 = 65535
                                                    _ = v101
                                                    _ = v102
                                                }
                                                switch "\n\t" {
                                                case "hello":
                                                    if v30 {
                                                        var v103 string = "a"
                                                        fmt.Println(fi37, fi43, v30)
                                                        v4 = v4
                                                        switch v103 {
                                                        case "\n\t":
                                                            switch -1 {
                                                            case 32768:
                                                                switch "foo bar" {
                                                                case "foo bar":
                                                                    switch "" {
                                                                    case "":
                                                                        var v104 uint = 591
                                                                        _ = v104
                                                                    case "hello":
                                                                        v103 = "\n\t"
                                                                    default:
                                                                        v105 := GenS9{F0: -1}
                                                                        _ = v105
                                                                    }
                                                                case "\n\t":
                                                                    v5[v103] = ""
                                                                case "a":
                                                                    fmt.Println(v29, fi43, v30)
                                                                default:
                                                                    if false {
                                                                        switch v103 {
                                                                        case "hello":
                                                                            if v30 {
                                                                                v1[v103] = 1
                                                                            }
                                                                        }
                                                                        if true {
                                                                            delete(v1, ((v103 + (("foo bar" + "foo bar") + (v103 + "a"))) + "\n\t"))
                                                                            v106 := GenS10{F0: (fi43 << ((65535) & 63)), F1: fi37}
                                                                            var v107 uint32 = 2
                                                                            var v108 bool = ((fi43 * 4294967295) <= 42)
                                                                            v30 = (128 >= 127)
                                                                            v103 = v103
                                                                            _ = v106
                                                                            _ = v107
                                                                            _ = v108
                                                                        } else {
                                                                            var v109 int = (v45 % ((1) | 1))
                                                                            var v110 uint32 = 42
                                                                            _ = v109
                                                                            _ = v110
                                                                        }
                                                                        switch (-(fi37)) {
                                                                        case 127:
                                                                            v29 = v29
                                                                        case 42:
                                                                            switch v103 {
                                                                            case "":
                                                                                if false {
                                                                                    if v4 {
                                                                                        switch fi43 {
                                                                                        case -959:
                                                                                            v1[(v103 + v103)] = 1
                                                                                        case -917:
                                                                                            v111 := v1["a"]
                                                                                            _ = v111
                                                                                        default:
                                                                                            if len(v3) > 0 {
                                                                                                v3[0] = 2
                                                                                            }
                                                                                        }
                                                                                        var v112 uint32 = 2147483648
                                                                                        _ = v112
                                                                                    } else {
                                                                                        switch 143 {
                                                                                        case 4294967295:
                                                                                            if len(v3) >= 2 {
                                                                                                v3 = v3[1:len(v3)]
                                                                                            }
                                                                                        case -991:
                                                                                            v113 := GenS8{F0: 255, F1: (^(v45)), F2: 0}
                                                                                            _ = v113
                                                                                        case 127:
                                                                                            _ = len(v3)
                                                                                        }
                                                                                        if false {
                                                                                            var v114 float32 = (((-2.25 * float32(0.1)) + v38) + (float32((0.0 + 0.0)) + ((v38 - 1.5) - 0.1)))
                                                                                            if (!(true && ((1 <= 9223372036854775807) || true))) {
                                                                                                v29 -= (v29 >> ((0) & 63))
                                                                                                var v115 bool = false
                                                                                                v116 := GenS11{F0: 65535}
                                                                                                _ = v115
                                                                                                _ = v116
                                                                                            } else {
                                                                                                fmt.Println(v29, v114)
                                                                                                v45 = v45
                                                                                            }
                                                                                            _ = v114
                                                                                        } else {
                                                                                            v30 = (3.14159 > 3.14159)
                                                                                            switch 42 {
                                                                                            case -9223372036854775808:
                                                                                                v117 := GenS4{F0: v29, F1: -231, F2: ("hello" + (v103 + v103))}
                                                                                                _ = v117
                                                                                            case 127:
                                                                                                v118 := GenS1{F0: 0, F1: v103, F2: 74}
                                                                                                _ = v118
                                                                                            default:
                                                                                                var v119 int8 = -128
                                                                                                switch "\n\t" {
                                                                                                case "":
                                                                                                    if v4 {
                                                                                                        switch v103 {
                                                                                                        case "hello":
                                                                                                            var v120 int = (^(fi37))
                                                                                                            _ = v120
                                                                                                        default:
                                                                                                            v119 += (v119 << ((42) & 63))
                                                                                                        }
                                                                                                        v121 := GenS10{F0: (v45 % (((fi37 / (((fi43 - (fi43 * v45))) | 1))) | 1)), F1: fi37}
                                                                                                        _ = v121
                                                                                                    } else {
                                                                                                        fmt.Println(v45, v44, v103)
                                                                                                        v2[v103] = ((!v4) || v30)
                                                                                                        v4 = v4
                                                                                                    }
                                                                                                case "hello":
                                                                                                    if v30 {
                                                                                                        var v122 int16 = -32768
                                                                                                        var v123 float64 = 1.5
                                                                                                        fmt.Println(v38)
                                                                                                        _ = v122
                                                                                                        _ = v123
                                                                                                    }
                                                                                                }
                                                                                                v45 = (v45 & (v45 - fi43))
                                                                                                _ = v119
                                                                                            }
                                                                                            v124 := GenS12{F0: (v44 / (((v44 ^ 2)) | 1)), F1: 679, F2: 0}
                                                                                            _ = v124
                                                                                        }
                                                                                    }
                                                                                    v125 := GenS6{F0: v103, F1: 2147483648, F2: 255}
                                                                                    if (v4 && (false && (genFn1((v29 + (v29 / ((v29) | 1))), 100) <= 42))) {
                                                                                        v45 = fi43
                                                                                    } else {
                                                                                        fmt.Println(v4)
                                                                                        v3 = append(v3, 0)
                                                                                        v3 = append(v3, 4294967295)
                                                                                        var v126 int32 = v29
                                                                                        _ = v126
                                                                                    }
                                                                                    _ = v125
                                                                                } else {
                                                                                    v2[""] = (!false)
                                                                                }
                                                                            case "hello":
                                                                                v3 = append(v3, 1)
                                                                            case "a":
                                                                                switch (v103 + v103) {
                                                                                case "a":
                                                                                    v45 = 4294967295
                                                                                case "foo bar":
                                                                                    v127 := GenS9{F0: -966}
                                                                                    _ = v127
                                                                                case "":
                                                                                    switch "\n\t" {
                                                                                    case "\n\t":
                                                                                        if len(v3) >= 2 {
                                                                                            v3 = v3[1:len(v3)]
                                                                                        }
                                                                                    case "":
                                                                                        v128 := GenS4{F0: -1, F1: 0, F2: "hello"}
                                                                                        _ = v128
                                                                                    default:
                                                                                        v129 := GenS10{F0: 4294967295, F1: 0}
                                                                                        _ = v129
                                                                                    }
                                                                                }
                                                                            default:
                                                                                if false {
                                                                                    if v30 {
                                                                                        switch (v103 + (("" + "") + "\n\t")) {
                                                                                        case "hello":
                                                                                            var v130 uint16 = 42
                                                                                            _ = v130
                                                                                        case "\n\t":
                                                                                            var v131 uint32 = 2147483647
                                                                                            switch (v103 + v103) {
                                                                                            case "":
                                                                                                if (4294967295 < (v44 % (((^(v44))) | 1))) {
                                                                                                    delete(v1, "a")
                                                                                                }
                                                                                            case "\n\t":
                                                                                                if false {
                                                                                                    switch v103 {
                                                                                                    case "\n\t":
                                                                                                        var v132 uint = 4294967295
                                                                                                        _ = v132
                                                                                                    }
                                                                                                    var v133 float64 = (1.5 * 1.5)
                                                                                                    _ = v133
                                                                                                }
                                                                                            case "hello":
                                                                                                var v134 bool = v4
                                                                                                _ = v134
                                                                                            default:
                                                                                                v2[v103] = (!v30)
                                                                                            }
                                                                                            var v135 uint = genFn1(v29, 0)
                                                                                            _ = v131
                                                                                            _ = v135
                                                                                        default:
                                                                                            fmt.Println(v38, fi43)
                                                                                        }
                                                                                    } else {
                                                                                        if (true && v30) {
                                                                                            var v136 int32 = (v29 >> ((65535) & 63))
                                                                                            var v137 string = v103
                                                                                            _ = v136
                                                                                            _ = v137
                                                                                        } else {
                                                                                            fmt.Println(v38, fi43)
                                                                                        }
                                                                                        var v138 int = -576
                                                                                        if len(v3) >= 2 {
                                                                                            v3 = v3[1:len(v3)]
                                                                                        }
                                                                                        _ = v138
                                                                                    }
                                                                                    v29 = -2147483648
                                                                                } else {
                                                                                    var v139 float32 = ((v38 - v38) / float32((0.1 / -2.25)))
                                                                                    var v140 string = v103
                                                                                    _ = v139
                                                                                    _ = v140
                                                                                }
                                                                                var v141 uint32 = 0
                                                                                fmt.Println(v30)
                                                                                _ = v141
                                                                            }
                                                                        case -369:
                                                                            fmt.Println(fi37)
                                                                        }
                                                                        v142 := GenS4{F0: (v29 >> ((4294967295) & 63)), F1: -689, F2: v103}
                                                                        _ = v142
                                                                    }
                                                                }
                                                            }
                                                        case "":
                                                            var v143 uint32 = 90
                                                            _ = v143
                                                        default:
                                                            v29 = v29
                                                        }
                                                        v45 = (fi43 >> ((genFn1(v29, -1)) & 63))
                                                        _ = v103
                                                    } else {
                                                        var v144 uint64 = 1
                                                        _ = v144
                                                    }
                                                case "foo bar":
                                                    fmt.Println(v38)
                                                }
                                                _ = len(v1)
                                                v145 := GenS4{F0: 440, F1: 9223372036854775807, F2: ""}
                                                _ = v145
                                            }
                                        } else {
                                            switch (fi43 - fi37) {
                                            case 32768:
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            case 512:
                                                v44 += (v44 >> ((genFn1(2147483647, -1)) & 63))
                                                if ("a" < "foo bar") {
                                                    switch (v45 & -1) {
                                                    case 32768:
                                                        if (165 >= 1000) {
                                                            if (!(false && v30)) {
                                                                v30 = (83 >= 219)
                                                            }
                                                            switch "" {
                                                            case "":
                                                                v146 := GenS13{F0: -1, F1: v45}
                                                                _ = v146
                                                            case "foo bar":
                                                                if v4 {
                                                                    var v147 int16 = 1
                                                                    var v148 float32 = (float32((-2.25 * (1e10 * 0.0))) / v38)
                                                                    v45 = (fi43 % ((1) | 1))
                                                                    _ = v147
                                                                    _ = v148
                                                                }
                                                            }
                                                        }
                                                    default:
                                                        if (!v4) {
                                                            fmt.Println(v30, v45, fi43)
                                                            if (!v4) {
                                                                if true {
                                                                    var v149 int32 = (v29 << ((252) & 63))
                                                                    var v150 string = ""
                                                                    fmt.Println(fi37)
                                                                    _ = v149
                                                                    _ = v150
                                                                }
                                                                v30 = v4
                                                                switch "foo bar" {
                                                                case "a":
                                                                    var v151 uint8 = 215
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                    _ = v151
                                                                default:
                                                                    if v4 {
                                                                        var v152 uint = genFn1((-(v29)), 59)
                                                                        v38 = (v38 * (3.14159 / (v38 - (-2.25 / v38))))
                                                                        _ = v152
                                                                    }
                                                                }
                                                            } else {
                                                                var v153 uint16 = 255
                                                                v29 = 42
                                                                var v154 uint = 4294967295
                                                                switch fi43 {
                                                                case 32767:
                                                                    if false {
                                                                        fmt.Println(v4, v38)
                                                                        switch "foo bar" {
                                                                        case "hello":
                                                                            v155 := GenS9{F0: -1}
                                                                            _ = v155
                                                                        default:
                                                                            v153 += 32767
                                                                        }
                                                                    } else {
                                                                        var v156 uint = v154
                                                                        fmt.Println(v156, v4, v154)
                                                                        var v157 float64 = 1e10
                                                                        _ = v156
                                                                        _ = v157
                                                                    }
                                                                case 128:
                                                                    if ("a" != "hello") {
                                                                        var v158 bool = (0 < (-(v29)))
                                                                        _ = v158
                                                                    }
                                                                    var v159 int32 = v29
                                                                    if ("\n\t" < "a") {
                                                                        v1["\n\t"] = uint32(v154)
                                                                        v160 := GenS8{F0: uint8(v154), F1: (fi43 % (((fi37 & (^(fi43)))) | 1)), F2: -92}
                                                                        v154 += 0
                                                                        if v4 {
                                                                            fmt.Println(v30, v29)
                                                                            if ((fi37 - fi43) <= fi43) {
                                                                                var v161 int = (fi43 ^ (fi37 * (fi43 ^ -807)))
                                                                                _ = v161
                                                                            } else {
                                                                                v44 = v44
                                                                                v45 = (fi43 >> ((v154) & 63))
                                                                            }
                                                                            v153 = 65535
                                                                        } else {
                                                                            v44 = 610
                                                                        }
                                                                        _ = v160
                                                                    }
                                                                    _ = v159
                                                                case 32768:
                                                                    fmt.Println(v154, v30)
                                                                }
                                                                _ = v153
                                                                _ = v154
                                                            }
                                                            v4 = v30
                                                        } else {
                                                            if true {
                                                                v30 = ((v29 ^ (^(v29))) >= genFn0(v30))
                                                            }
                                                            v30 = false
                                                            v45 = v45
                                                        }
                                                    }
                                                } else {
                                                    if true {
                                                        fmt.Println(v29, v44)
                                                        v162 := GenS14{F0: 761, F1: genFn0(v30)}
                                                        v38 = v38
                                                        _ = v162
                                                    }
                                                    _ = len(v3)
                                                    if v4 {
                                                        v4 = (4294967295 != 1)
                                                    } else {
                                                        v163 := GenS12{F0: (^(v44)), F1: 1, F2: 705}
                                                        var v164 string = "hello"
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            var v165 uint8 = 2
                                                            _ = v165
                                                        }
                                                        _ = v163
                                                        _ = v164
                                                    }
                                                }
                                                v166 := GenS14{F0: 9223372036854775807, F1: (v29 / (((v29 * v29)) | 1))}
                                                _ = v166
                                            case 0:
                                                fmt.Println(v30, fi37)
                                            default:
                                                v30 = (v38 >= ((0.0 / ((v38 + 3.14159) - -2.25)) / v38))
                                            }
                                            if v30 {
                                                v30 = (v4 && ((183 >= 128) && (!false)))
                                                var v167 uint64 = 998
                                                switch "" {
                                                case "a":
                                                    if v4 {
                                                        var v168 float64 = 0.0
                                                        switch "" {
                                                        case "hello":
                                                            var v169 uint8 = 1
                                                            _ = v169
                                                        case "\n\t":
                                                            v168 -= -2.25
                                                        default:
                                                            v4 = v30
                                                        }
                                                        var v170 uint64 = v44
                                                        v38 = 0.1
                                                        switch (v45 * (fi43 / (((^(v45))) | 1))) {
                                                        case 4294967295:
                                                            v38 = v38
                                                        case 42:
                                                            var v171 uint = 18446744073709551615
                                                            _ = v171
                                                        case 127:
                                                            v38 -= v38
                                                        }
                                                        _ = v168
                                                        _ = v170
                                                    } else {
                                                        v30 = false
                                                        var v172 int = (fi43 << ((0) & 63))
                                                        _ = v172
                                                    }
                                                case "foo bar":
                                                    v167 = v44
                                                case "":
                                                    var v173 uint64 = (v44 / ((v167) | 1))
                                                    _ = v173
                                                }
                                                var v174 int16 = 0
                                                _ = v167
                                                _ = v174
                                            } else {
                                                fmt.Println(v45, v38)
                                                var v175 uint16 = 0
                                                _ = v175
                                            }
                                            if ((!v4) && (2147483647 > (v29 - (v29 / ((2147483647) | 1))))) {
                                                switch -9223372036854775808 {
                                                case 32768:
                                                    var v176 int64 = 0
                                                    _ = v176
                                                case -324:
                                                    switch "a" {
                                                    case "a":
                                                        v30 = (2147483647 < 2147483647)
                                                    case "":
                                                        fmt.Println(v29)
                                                    }
                                                }
                                                v38 -= v38
                                            }
                                            switch (v45 << ((0) & 63)) {
                                            case -9223372036854775808:
                                                if ((729 == 367) || v4) {
                                                    v177 := GenS15{F0: (0.1 >= (1.5 * 3.14159)), F1: "a"}
                                                    v44 = v44
                                                    _ = v177
                                                } else {
                                                    if true {
                                                        v30 = ((v44 * v44) != v44)
                                                    } else {
                                                        if (((((v29 + v29) <= v29) || (int(v29) >= v45)) || v4) || (827 < genFn1(-2147483648, 1))) {
                                                            fmt.Println(v38, fi43)
                                                        }
                                                        if v30 {
                                                            v45 = -1
                                                            var v178 int8 = 1
                                                            if v4 {
                                                                fmt.Println(v38)
                                                                var v179 float32 = v38
                                                                _ = v179
                                                            } else {
                                                                v180 := GenS5{F0: 651, F1: (fi37 * (v45 % ((fi43) | 1))), F2: 753}
                                                                v181 := GenS1{F0: 2, F1: ("a" + ("foo bar" + "a")), F2: 32767}
                                                                _ = v180
                                                                _ = v181
                                                            }
                                                            _ = v178
                                                        }
                                                    }
                                                }
                                            case 9223372036854775807:
                                                if (32767 > 406) {
                                                    var v182 float64 = 1.5
                                                    var v183 float64 = v182
                                                    var v184 uint = 65535
                                                    _ = v182
                                                    _ = v183
                                                    _ = v184
                                                }
                                            }
                                            var v185 float32 = (((((v38 - (((1.5 / ((float32(1e10) + (1e10 + (0.1 - v38))) + v38)) + v38) + float32((1.5 / float64(3.14159))))) / float32(1.5)) * (v38 / v38)) - v38) - v38)
                                            _ = v185
                                        }
                                    }
                                    v3 = append(v3, 720)
                                    var v186 int8 = 100
                                    fmt.Println(fi37, v186, v30)
                                    _ = v186
                                }
                                var v187 int = v45
                                v38 -= 0.1
                                _ = v187
                            }
                            if (!v30) {
                                if (-112 == 9223372036854775807) {
                                    if len(v3) > 0 {
                                        v3[0] = 2147483647
                                    }
                                    v30 = v4
                                    var v188 uint64 = 2
                                    _ = v188
                                } else {
                                    v189 := GenS5{F0: 65535, F1: v45, F2: -1}
                                    var v190 float64 = float64(v38)
                                    var v191 uint8 = 2
                                    if (("a" + ("a" + "foo bar")) >= "foo bar") {
                                        v45 -= int(v29)
                                        if v30 {
                                            v1["\n\t"] = 2147483648
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        } else {
                                            var v192 int8 = -1
                                            fmt.Println(fi43, v29, fi37)
                                            _ = v192
                                        }
                                    }
                                    v45 += (fi43 / (((v45 + v45)) | 1))
                                    _ = v189
                                    _ = v190
                                    _ = v191
                                }
                                var v193 uint = genFn1((v29 | -808), 32)
                                v29 = v29
                                _ = v193
                            }
                        }
                        v194 := GenS14{F0: -1, F1: (v29 - (v29 | (v29 / ((genFn0((v30 || v4))) | 1))))}
                        switch 66 {
                        case 846:
                            if true {
                                v44 = (v44 | (v44 << ((genFn1((-(v29)), 80)) & 63)))
                                var v195 int64 = 9223372036854775807
                                v38 = (v38 * v38)
                                _ = v195
                            }
                        case 0:
                            var v196 int32 = genFn0((1.5 <= v38))
                            _ = v196
                        default:
                            var v197 int8 = -46
                            _ = v197
                        }
                        _ = v194
                    } else {
                        fmt.Println(v4, v38, v30)
                        switch "\n\t" {
                        case "a":
                            var v198 int = (-(fi43))
                            _ = v198
                        default:
                            v38 = (v38 - v38)
                        }
                    }
                    _ = v44
                    _ = v45
                } else {
                    v4 = (200 <= 2)
                    var v199 bool = (!v30)
                    fmt.Println(v199)
                    _ = v199
                }
                v3 = append(v3, 2147483648)
                _ = fi43
            }
        default:
            _ = len(v1)
        }
        _ = fi37
        _ = v38
    }
    var v200 float32 = float32((3.14159 * -2.25))
    v4 = ("foo bar" >= ((("a" + "\n\t") + "a") + "\n\t"))
    for fi201 := 0; fi201 < 7; fi201++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi201
    }
    v202 := GenS11{F0: 2}
    if v4 {
        v4 = v4
        switch "a" {
        case "":
            if v4 {
                if (!true) {
                    var v203 int8 = 42
                    var v204 float64 = 1.5
                    _ = v203
                    _ = v204
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        case "\n\t":
            if v4 {
                v5["\n\t"] = ("hello" + "a")
                switch "hello" {
                case "":
                    switch ("\n\t" + "foo bar") {
                    case "":
                        v205 := GenS6{F0: "a", F1: 2147483648, F2: 4294967295}
                        _ = v205
                    case "a":
                        v206 := v1["foo bar"]
                        _ = v206
                    case "hello":
                        v207 := GenS16{F0: 42, F1: genFn1((v29 + genFn0(v30)), -128), F2: 2147483648}
                        _ = v207
                    }
                default:
                    fmt.Println(v200, v4, v200)
                }
            }
        default:
            if (!v30) {
                var v208 int32 = 42
                for fi209 := 0; fi209 < 14; fi209++ {
                    v210 := GenS17{F0: -894}
                    if ((!(!v30)) && v4) {
                        switch fi209 {
                        case -210:
                            delete(v5, "\n\t")
                        default:
                            for fi211 := 0; fi211 < 19; fi211++ {
                                v212 := GenS18{F0: 18446744073709551615, F1: 116}
                                v213 := GenS19{F0: 1}
                                if (0 >= 1) {
                                    var v214 int16 = 1000
                                    switch "foo bar" {
                                    case "":
                                        if (775 <= 0) {
                                            v215 := GenS20{F0: v30}
                                            _ = v215
                                        } else {
                                            var v216 float64 = 1.5
                                            _ = v216
                                        }
                                    default:
                                        fmt.Println(fi209)
                                    }
                                    v4 = true
                                    v4 = v4
                                    var v217 uint32 = 222
                                    _ = v214
                                    _ = v217
                                } else {
                                    v218 := GenS17{F0: -363}
                                    var v219 int8 = 17
                                    _ = v218
                                    _ = v219
                                }
                                v4 = v4
                                _ = fi211
                                _ = v212
                                _ = v213
                            }
                        }
                    }
                    _ = fi209
                    _ = v210
                }
                _ = v208
            }
        }
    }
    var v220 bool = (false || true)
    fmt.Println(v30)
    v29 += v29
    switch 626 {
    case -47:
        for fi221 := 0; fi221 < 16; fi221++ {
            var v222 int16 = 934
            fmt.Println(v220, v29)
            v223 := GenS21{F0: v29, F1: "hello", F2: 2147483647}
            _ = fi221
            _ = v222
            _ = v223
        }
    case 907:
        var v224 int32 = (^(v29))
        _ = v224
    case 32767:
        for fi225 := 0; fi225 < 11; fi225++ {
            v220 = true
            v29 = (v29 ^ v29)
            _ = fi225
        }
    default:
        if ((((((float64(1.5) + 0.0) - 0.0) * (1.5 + 0.1)) - (float64(v200) - 3.14159)) / 1e10) < (((((float64(float32(0.0)) + -2.25) / 1e10) - -2.25) + 0.1) + (float64(0.1) / float64((float32((((float64(0.0) / (0.1 / 1e10)) / 0.1) * 1.5)) + (v200 / -2.25)))))) {
            for fi226 := 0; fi226 < 3; fi226++ {
                if v30 {
                    if v30 {
                        v227 := GenS16{F0: 2, F1: 2, F2: 65535}
                        fmt.Println(v29)
                        _ = v227
                    }
                }
                var v228 uint64 = 42
                v229 := GenS14{F0: 1, F1: (v29 ^ genFn0(v30))}
                _ = fi226
                _ = v228
                _ = v229
            }
            if false {
                switch 743 {
                case -116:
                    if (-9223372036854775808 < 9223372036854775807) {
                        v4 = (0 != 0)
                        v220 = (v4 && (1 > 702))
                    } else {
                        var v230 uint8 = 2
                        _ = v230
                    }
                case -9223372036854775808:
                    v220 = v30
                }
            } else {
                if true {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v29 -= v29
                    fmt.Println(v200, v4)
                } else {
                    fmt.Println(v200)
                    var v231 bool = (2147483648 >= 126)
                    var v232 uint = 42
                    _ = v231
                    _ = v232
                }
                for fi233 := 0; fi233 < 17; fi233++ {
                    if v220 {
                        v30 = v30
                    }
                    var v234 string = "a"
                    switch (fi233 / ((127) | 1)) {
                    case -553:
                        if (965 > 1) {
                            var v235 uint32 = 1
                            _ = v235
                        }
                    case 0:
                        v220 = v4
                    case 128:
                        fmt.Println(v200, v220, v4)
                    default:
                        v236 := GenS2{F0: 0, F1: -1}
                        _ = v236
                    }
                    _ = fi233
                    _ = v234
                }
            }
        }
    }
    v30 = (3.14159 <= 1e10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v200:", v200)
    fmt.Println("v220:", v220)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v200
    _ = v202
    _ = v220
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
