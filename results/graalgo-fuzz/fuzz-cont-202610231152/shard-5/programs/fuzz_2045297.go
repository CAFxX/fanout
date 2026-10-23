package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 int16; F1 bool; F2 uint }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int64; F1 uint32; F2 int8 }
type GenS7 struct { F0 uint16; F1 int; F2 uint64 }
type GenS8 struct { F0 uint8; F1 uint; F2 int64 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 bool; F1 int8 }
type GenS12 struct { F0 uint16; F1 int }
type GenS13 struct { F0 string; F1 bool; F2 uint64 }
type GenS14 struct { F0 int }
type GenS15 struct { F0 uint64 }
type GenS16 struct { F0 int; F1 uint64 }
type GenS17 struct { F0 int32 }
type GenS18 struct { F0 bool; F1 uint; F2 int16 }
type GenS19 struct { F0 uint8 }
type GenS20 struct { F0 uint; F1 uint16 }
type GenS21 struct { F0 int8; F1 int }
type GenS22 struct { F0 int8 }
type GenS23 struct { F0 uint64; F1 bool }
type GenS24 struct { F0 string; F1 uint }
type GenS25 struct { F0 uint8; F1 int; F2 int16 }
type GenS26 struct { F0 int }
type GenS27 struct { F0 uint32; F1 uint; F2 uint }
func genFn0(p0 int16, p1 uint32) uint8 {
    p1 = (p1 % ((42) | 1))
    var v5 float64 = (float64(((3.14159 + -2.25) - 1.5)) * float64((((0.0 * float32((0.0 - (0.0 * (3.14159 / 1.5))))) + float32((1.5 + (1.5 / 1e10)))) * (-2.25 / 1.0))))
    _ = p0
    _ = p1
    _ = v5
    return 128
}

func main() {
    var v1 uint64 = 1
    v2 := make(map[string]string)
    var v3 uint8 = 127
    var v4 int32 = 887
    fmt.Println(v4, v3, v1)
    v6 := GenS1{F0: 551}
    switch "foo bar" {
    case "a":
        v7 := GenS2{F0: -51, F1: (v1 ^ v1)}
        _ = v7
    case "hello":
        fmt.Println(v3)
    case "":
        fmt.Println(v3)
    default:
        v4 += int32(v3)
    }
    v8 := GenS3{F0: 9223372036854775807, F1: 0}
    for fi9 := 0; fi9 < 3; fi9++ {
        for fi10 := 0; fi10 < 3; fi10++ {
            fmt.Println(fi9, v4)
            var v11 bool = (true && (false && (((1 <= 0) || false) || ((!true) && (1 > -333)))))
            var v12 int32 = -865
            _ = fi10
            _ = v11
            _ = v12
        }
        for fi13 := 0; fi13 < 8; fi13++ {
            var v14 int32 = 2147483647
            _ = fi13
            _ = v14
        }
        _ = fi9
    }
    v1 += (v1 % (((v1 % ((v1) | 1))) | 1))
    v15 := GenS2{F0: 100, F1: (v1 & v1)}
    fmt.Println(v4)
    v4 += v4
    v1 = (v1 & v1)
    fmt.Println(v4, v4, v3)
    switch -2147483648 {
    case -917:
        fmt.Println(v3, v1)
    }
    if true {
        v1 += 1
        if false {
            for fi16 := 0; fi16 < 4; fi16++ {
                if true {
                    if true {
                        for fi17 := 0; fi17 < 17; fi17++ {
                            v4 = v4
                            v18 := []int{790, fi17}
                            v3 = (v3 - (^(v3)))
                            if false {
                                var v19 int16 = 1000
                                v19 -= -897
                                var v20 int16 = (v19 ^ v19)
                                _ = v19
                                _ = v20
                            }
                            _ = fi17
                            _ = v18
                        }
                        switch "foo bar" {
                        case "\n\t":
                            switch 9223372036854775807 {
                            case 9223372036854775807:
                                v1 -= 27
                            case -598:
                                for fi21 := 0; fi21 < 17; fi21++ {
                                    switch -2147483648 {
                                    case 32767:
                                        switch "a" {
                                        case "":
                                            v2["\n\t"] = ""
                                        }
                                    default:
                                        if true {
                                            var v22 uint = 1
                                            var v23 int32 = v4
                                            _ = v22
                                            _ = v23
                                        } else {
                                            switch (fi21 / ((fi21) | 1)) {
                                            case 1:
                                                _ = len(v2)
                                            case -643:
                                                fmt.Println(v1, v1)
                                            case 646:
                                                v24 := GenS4{F0: 32767, F1: (float32(0.1) >= (-2.25 + 0.1)), F2: 18446744073709551615}
                                                _ = v24
                                            }
                                            v1 = (^(v1))
                                            switch ("hello" + "a") {
                                            case "foo bar":
                                                var v25 float64 = 0.1
                                                v1 += 0
                                                _ = v25
                                            case "":
                                                v3 = (v3 * (v3 / ((0) | 1)))
                                            case "a":
                                                if (float32(float64(0.0)) <= 1.5) {
                                                    switch "" {
                                                    case "foo bar":
                                                        switch 127 {
                                                        case 127:
                                                            v3 = 2
                                                        case 2147483647:
                                                            var v26 bool = (false && false)
                                                            _ = v26
                                                        case -9223372036854775808:
                                                            if true {
                                                                var v27 int64 = -1
                                                                var v28 float32 = 0.0
                                                                v28 = v28
                                                                _ = v27
                                                                _ = v28
                                                            }
                                                        default:
                                                            switch -2147483648 {
                                                            case -483:
                                                                switch "hello" {
                                                                case "\n\t":
                                                                    var v29 float32 = 0.0
                                                                    _ = v29
                                                                case "foo bar":
                                                                    switch fi16 {
                                                                    case 127:
                                                                        fmt.Println(v1, v1, v1)
                                                                    case -1:
                                                                        if (1e10 != ((1e10 + 3.14159) + 0.1)) {
                                                                            var v30 float64 = ((1e10 - float64(float32((float64((((1.5 / ((0.1 + 3.14159) - 0.0)) - 3.14159) * 3.14159)) * 0.1)))) - float64((-2.25 / 1e10)))
                                                                            if true {
                                                                                var v31 int = (fi16 << ((2) & 63))
                                                                                v30 += 0.0
                                                                                _ = v31
                                                                            } else {
                                                                                v32 := GenS3{F0: 406, F1: 282}
                                                                                _ = v32
                                                                            }
                                                                            _ = v30
                                                                        }
                                                                        var v33 uint32 = 2147483648
                                                                        fmt.Println(fi16)
                                                                        _ = v33
                                                                    default:
                                                                        v3 -= v3
                                                                    }
                                                                    v3 = 255
                                                                    v34 := GenS3{F0: 0, F1: 42}
                                                                    _ = v34
                                                                case "hello":
                                                                    delete(v2, "")
                                                                }
                                                            case 1:
                                                                v35 := GenS5{F0: (v4 % ((v4) | 1))}
                                                                _ = v35
                                                            }
                                                        }
                                                        if false {
                                                            switch "" {
                                                            case "\n\t":
                                                                if false {
                                                                    var v36 bool = true
                                                                    _ = v36
                                                                }
                                                            case "hello":
                                                                v1 = (v1 & (v1 * (^(v1))))
                                                            default:
                                                                v37 := GenS2{F0: 55, F1: v1}
                                                                _ = v37
                                                            }
                                                        } else {
                                                            v2["\n\t"] = "a"
                                                            v1 = v1
                                                            var v38 uint = 0
                                                            _ = v38
                                                        }
                                                    case "\n\t":
                                                        if (!true) {
                                                            switch ("hello" + "") {
                                                            case "\n\t":
                                                                var v39 int32 = 1
                                                                fmt.Println(fi21)
                                                                _ = v39
                                                            case "foo bar":
                                                                v3 = 128
                                                            default:
                                                                v40 := GenS1{F0: 128}
                                                                _ = v40
                                                            }
                                                            var v41 string = "hello"
                                                            v42 := []int{fi21, (fi16 ^ 557)}
                                                            _ = v41
                                                            _ = v42
                                                        } else {
                                                            var v43 uint8 = 42
                                                            fmt.Println(v4, v43)
                                                            v44 := GenS6{F0: -686, F1: 921, F2: 87}
                                                            _ = v43
                                                            _ = v44
                                                        }
                                                    case "":
                                                        if true {
                                                            if false {
                                                                v3 += (v3 >> ((1) & 63))
                                                            } else {
                                                                if true {
                                                                    if true {
                                                                        v2[""] = ""
                                                                        var v45 string = "foo bar"
                                                                        if ((0 == 42) && true) {
                                                                            v45 = "hello"
                                                                            var v46 string = v45
                                                                            _ = v46
                                                                        } else {
                                                                            v1 -= v1
                                                                            switch fi16 {
                                                                            case 406:
                                                                                fmt.Println(fi16, fi21, v1)
                                                                            default:
                                                                                v47 := v2["foo bar"]
                                                                                _ = v47
                                                                            }
                                                                            v48 := GenS6{F0: -9223372036854775808, F1: 2147483648, F2: 15}
                                                                            _ = v48
                                                                        }
                                                                        _ = v45
                                                                    } else {
                                                                        var v49 string = ("foo bar" + "\n\t")
                                                                        _ = v49
                                                                    }
                                                                    var v50 uint32 = 2147483647
                                                                    _ = v50
                                                                } else {
                                                                    v51 := GenS1{F0: (fi21 % (((^(fi16))) | 1))}
                                                                    v52 := []int{fi16, fi16, fi16}
                                                                    _ = v51
                                                                    _ = v52
                                                                }
                                                                if ((0 == 390) && ((470 > 32767) || (true || true))) {
                                                                    if (76 != 2) {
                                                                        v4 = v4
                                                                    } else {
                                                                        if true {
                                                                            var v53 int32 = (v4 >> ((315) & 63))
                                                                            _ = v53
                                                                        } else {
                                                                            v54 := GenS7{F0: 231, F1: (fi21 * (fi16 * (fi21 / ((409) | 1)))), F2: 224}
                                                                            v55 := GenS3{F0: 158, F1: uint32(v3)}
                                                                            _ = v54
                                                                            _ = v55
                                                                        }
                                                                        v56 := v2["\n\t"]
                                                                        v4 += v4
                                                                        _ = v56
                                                                    }
                                                                    v4 = v4
                                                                }
                                                                v57 := GenS8{F0: genFn0(1, 42), F1: 18446744073709551615, F2: 9223372036854775807}
                                                                if true {
                                                                    v4 -= v4
                                                                }
                                                                var v58 uint64 = (v1 << ((357) & 63))
                                                                _ = v57
                                                                _ = v58
                                                            }
                                                            fmt.Println(v1)
                                                            v59 := GenS9{F0: 42}
                                                            var v60 int32 = (v4 / (((v4 << ((18446744073709551615) & 63))) | 1))
                                                            v1 = (v1 | 4294967295)
                                                            switch "a" {
                                                            case "hello":
                                                                var v61 int16 = 32767
                                                                _ = v61
                                                            }
                                                            if false {
                                                                var v62 float32 = 1e10
                                                                v4 += v60
                                                                v2[("a" + "foo bar")] = "a"
                                                                _ = v62
                                                            }
                                                            _ = v59
                                                            _ = v60
                                                        } else {
                                                            v3 = (v3 % ((127) | 1))
                                                            v63 := GenS10{F0: true}
                                                            _ = v63
                                                        }
                                                    }
                                                    var v64 uint8 = 42
                                                    _ = v64
                                                }
                                            default:
                                                if ((-(fi16)) != fi16) {
                                                    v65 := GenS11{F0: ((!true) || false), F1: -1}
                                                    v3 = genFn0(-720, 42)
                                                    _ = v65
                                                } else {
                                                    if false {
                                                        v4 -= (^(v4))
                                                        if false {
                                                            fmt.Println(v4, fi16, v3)
                                                        }
                                                        v1 = v1
                                                    } else {
                                                        var v66 uint = 136
                                                        _ = v66
                                                    }
                                                    if false {
                                                        if true {
                                                            v67 := []int{fi21}
                                                            _ = v67
                                                        } else {
                                                            var v68 uint8 = v3
                                                            v1 -= (v1 % ((v1) | 1))
                                                            _ = v68
                                                        }
                                                    }
                                                }
                                                if false {
                                                    var v69 float32 = (0.0 - (1.5 * 0.0))
                                                    if false {
                                                        var v70 uint64 = (v1 | 563)
                                                        switch fi21 {
                                                        case 128:
                                                            fmt.Println(v3, v1, v70)
                                                            if (false && false) {
                                                                switch ("foo bar" + "a") {
                                                                case "foo bar":
                                                                    v70 += (v70 & uint64(v3))
                                                                    var v71 uint16 = uint16(v1)
                                                                    _ = v71
                                                                case "hello":
                                                                    var v72 int64 = int64(v4)
                                                                    _ = v72
                                                                default:
                                                                    v73 := GenS12{F0: 255, F1: 32767}
                                                                    _ = v73
                                                                }
                                                            } else {
                                                                switch ("hello" + "a") {
                                                                case "a":
                                                                    if (0.0 == 3.14159) {
                                                                        switch -9223372036854775808 {
                                                                        case -2147483648:
                                                                            v1 -= 18446744073709551615
                                                                        case 42:
                                                                            if ("foo bar" == "foo bar") {
                                                                                v74 := []int{(-(fi16)), fi16, fi16, fi21}
                                                                                var v75 uint8 = v3
                                                                                switch 896 {
                                                                                case 127:
                                                                                    v76 := v2[(("foo bar" + "\n\t") + "a")]
                                                                                    _ = v76
                                                                                case 2147483647:
                                                                                    v77 := GenS12{F0: 32767, F1: (fi16 % ((-186) | 1))}
                                                                                    _ = v77
                                                                                case 9223372036854775807:
                                                                                    if false {
                                                                                        v1 = v70
                                                                                    }
                                                                                default:
                                                                                    v2["hello"] = ("a" + ((("a" + "") + "a") + "hello"))
                                                                                }
                                                                                if false {
                                                                                    if true {
                                                                                        if (3.14159 <= 0.0) {
                                                                                            var v78 bool = ((v70 == (v1 ^ (v1 * (v70 * v1)))) || false)
                                                                                            var v79 int8 = 57
                                                                                            v80 := v2[("a" + "\n\t")]
                                                                                            fmt.Println(v4, fi16)
                                                                                            var v81 string = "a"
                                                                                            fmt.Println(v69, v4, fi16)
                                                                                            fmt.Println(fi16, v4)
                                                                                            switch -2147483648 {
                                                                                            case 0:
                                                                                                var v82 uint64 = (^(v70))
                                                                                                _ = v82
                                                                                            case 903:
                                                                                                v70 = (^(v1))
                                                                                            case -178:
                                                                                                var v83 int64 = -1
                                                                                                _ = v83
                                                                                            }
                                                                                            switch "a" {
                                                                                            case "a":
                                                                                                v80 = (v81 + v80)
                                                                                            default:
                                                                                                v70 = (v70 / ((v1) | 1))
                                                                                            }
                                                                                            switch (fi16 ^ (fi16 << ((1) & 63))) {
                                                                                            case 885:
                                                                                                v70 += v1
                                                                                            }
                                                                                            v3 -= v75
                                                                                            _ = v78
                                                                                            _ = v79
                                                                                            _ = v80
                                                                                            _ = v81
                                                                                        }
                                                                                        if (float32(float64(1e10)) < float32((1e10 + 0.1))) {
                                                                                            v3 -= 127
                                                                                        } else {
                                                                                            var v84 uint64 = (^(v1))
                                                                                            var v85 uint64 = v84
                                                                                            _ = v84
                                                                                            _ = v85
                                                                                        }
                                                                                    } else {
                                                                                        v86 := GenS9{F0: 100}
                                                                                        _ = v86
                                                                                    }
                                                                                    v69 -= (v69 / v69)
                                                                                    switch fi21 {
                                                                                    case -638:
                                                                                        v87 := GenS9{F0: 61}
                                                                                        _ = v87
                                                                                    case 128:
                                                                                        v74 = append(v74, fi21)
                                                                                    default:
                                                                                        v70 -= (v70 % ((uint64(v75)) | 1))
                                                                                    }
                                                                                } else {
                                                                                    if (v4 <= (v4 * v4)) {
                                                                                        if false {
                                                                                            v69 -= float32(3.14159)
                                                                                            var v88 int32 = v4
                                                                                            _ = v88
                                                                                        }
                                                                                        switch ("foo bar" + "\n\t") {
                                                                                        case "foo bar":
                                                                                            v89 := GenS10{F0: false}
                                                                                            _ = v89
                                                                                        }
                                                                                    }
                                                                                    if (!true) {
                                                                                        var v90 float32 = 0.1
                                                                                        if ("a" >= "foo bar") {
                                                                                            if (v4 <= v4) {
                                                                                                delete(v2, "foo bar")
                                                                                                fmt.Println(v3)
                                                                                                if (127 > 2147483647) {
                                                                                                    if (false || false) {
                                                                                                        v4 = -360
                                                                                                    } else {
                                                                                                        var v91 float32 = v69
                                                                                                        if (true && false) {
                                                                                                            v3 = genFn0(32767, 42)
                                                                                                            v90 -= float32(1.5)
                                                                                                        }
                                                                                                        _ = v91
                                                                                                    }
                                                                                                } else {
                                                                                                    fmt.Println(v3, v4)
                                                                                                    v69 = (3.14159 / ((v69 + (v90 * ((1e10 / v90) * v90))) + -2.25))
                                                                                                    _ = len(v2)
                                                                                                }
                                                                                            }
                                                                                            _ = len(v2)
                                                                                            switch (-(fi16)) {
                                                                                            case -937:
                                                                                                v3 -= 103
                                                                                            default:
                                                                                                v90 -= (-2.25 - v90)
                                                                                            }
                                                                                        }
                                                                                        var v92 uint8 = v75
                                                                                        _ = v90
                                                                                        _ = v92
                                                                                    }
                                                                                    var v93 int32 = v4
                                                                                    _ = v93
                                                                                }
                                                                                var v94 string = "\n\t"
                                                                                _ = len(v2)
                                                                                _ = v74
                                                                                _ = v75
                                                                                _ = v94
                                                                            }
                                                                        case 587:
                                                                            switch fi21 {
                                                                            case 32768:
                                                                                var v95 float32 = v69
                                                                                _ = v95
                                                                            case 2147483647:
                                                                                var v96 int32 = (v4 << ((286) & 63))
                                                                                _ = v96
                                                                            case -9223372036854775808:
                                                                                switch -274 {
                                                                                case -385:
                                                                                    _ = len(v2)
                                                                                case 559:
                                                                                    v97 := GenS4{F0: 1000, F1: false, F2: 52}
                                                                                    _ = v97
                                                                                case 44:
                                                                                    v1 -= v70
                                                                                default:
                                                                                    var v98 string = ("a" + "\n\t")
                                                                                    _ = v98
                                                                                }
                                                                            }
                                                                            if (-9223372036854775808 < -675) {
                                                                                var v99 bool = (int16(fi16) > 481)
                                                                                v1 = (v1 | v70)
                                                                                _ = v99
                                                                            } else {
                                                                                v3 -= genFn0(1, 0)
                                                                                if (false || ("\n\t" >= "\n\t")) {
                                                                                    v100 := []int{(fi16 - fi16), (fi21 & fi16), fi16}
                                                                                    v70 = 2
                                                                                    _ = v100
                                                                                }
                                                                                v3 = v3
                                                                            }
                                                                        default:
                                                                            switch "a" {
                                                                            case "a":
                                                                                var v101 float32 = v69
                                                                                _ = v101
                                                                            case "hello":
                                                                                v69 = 1.5
                                                                            default:
                                                                                v4 = v4
                                                                            }
                                                                        }
                                                                        v4 += v4
                                                                        v102 := []int{fi21, fi21, (fi16 * fi21), 2147483647}
                                                                        var v103 string = "a"
                                                                        fmt.Println(fi16, v69)
                                                                        var v104 int8 = 82
                                                                        _ = v102
                                                                        _ = v103
                                                                        _ = v104
                                                                    } else {
                                                                        _ = len(v2)
                                                                        switch fi21 {
                                                                        case 0:
                                                                            v105 := v2["a"]
                                                                            _ = v105
                                                                        case 1:
                                                                            fmt.Println(v69, v69, fi21)
                                                                        case -784:
                                                                            var v106 int8 = 0
                                                                            _ = v106
                                                                        default:
                                                                            switch ("hello" + ("\n\t" + "\n\t")) {
                                                                            case "a":
                                                                                if true {
                                                                                    v69 = 3.14159
                                                                                } else {
                                                                                    fmt.Println(fi21, v3, v70)
                                                                                    v70 += v70
                                                                                    var v107 uint = 935
                                                                                    _ = v107
                                                                                }
                                                                            case "\n\t":
                                                                                v108 := GenS4{F0: 1000, F1: ((-(fi16)) != 32767), F2: 2}
                                                                                _ = v108
                                                                            default:
                                                                                fmt.Println(fi16, fi16)
                                                                            }
                                                                        }
                                                                        if ((v3 - v3) > v3) {
                                                                            v3 -= v3
                                                                            v4 = 295
                                                                            v109 := GenS6{F0: -9223372036854775808, F1: 973, F2: 100}
                                                                            _ = v109
                                                                        } else {
                                                                            if true {
                                                                                switch fi21 {
                                                                                case 2147483647:
                                                                                    fmt.Println(v70)
                                                                                case 755:
                                                                                    var v110 float32 = 1e10
                                                                                    _ = v110
                                                                                default:
                                                                                    if false {
                                                                                        v111 := []int{int(v4), fi21, (fi21 / ((fi21) | 1)), fi16}
                                                                                        if false {
                                                                                            v112 := GenS13{F0: "foo bar", F1: true, F2: v70}
                                                                                            v1 -= v70
                                                                                            _ = v112
                                                                                        }
                                                                                        delete(v2, "")
                                                                                        if (!true) {
                                                                                            if true {
                                                                                                var v113 uint = 18446744073709551615
                                                                                                switch -659 {
                                                                                                case -128:
                                                                                                    v114 := GenS14{F0: fi21}
                                                                                                    _ = v114
                                                                                                case 851:
                                                                                                    delete(v2, "foo bar")
                                                                                                case 0:
                                                                                                    fmt.Println(v3, v113)
                                                                                                }
                                                                                                _ = v113
                                                                                            }
                                                                                            switch (fi16 ^ (fi16 ^ fi21)) {
                                                                                            case -1:
                                                                                                if false {
                                                                                                    v115 := GenS4{F0: 42, F1: true, F2: 648}
                                                                                                    _ = v115
                                                                                                } else {
                                                                                                    fmt.Println(v4, v70, v1)
                                                                                                }
                                                                                                _ = len(v111)
                                                                                                var v116 uint16 = 42
                                                                                                _ = v116
                                                                                            case 1:
                                                                                                if (!true) {
                                                                                                    switch 32767 {
                                                                                                    case -2147483648:
                                                                                                        if false {
                                                                                                            var v117 int32 = v4
                                                                                                            v1 = v70
                                                                                                            _ = v117
                                                                                                        }
                                                                                                    case 127:
                                                                                                        v70 -= 0
                                                                                                    case 1:
                                                                                                        var v118 bool = true
                                                                                                        _ = v118
                                                                                                    }
                                                                                                    var v119 uint8 = uint8(v1)
                                                                                                    var v120 uint64 = (v70 / (((v70 - (v1 % ((v70) | 1)))) | 1))
                                                                                                    _ = v119
                                                                                                    _ = v120
                                                                                                } else {
                                                                                                    v1 = (^(v70))
                                                                                                    if true {
                                                                                                        switch "hello" {
                                                                                                        case "foo bar":
                                                                                                            _ = len(v2)
                                                                                                        case "":
                                                                                                            if false {
                                                                                                                var v121 uint64 = v70
                                                                                                                v70 -= v70
                                                                                                                _ = v121
                                                                                                            }
                                                                                                        case "\n\t":
                                                                                                            v122 := GenS6{F0: -90, F1: 42, F2: 117}
                                                                                                            _ = v122
                                                                                                        }
                                                                                                        v69 = v69
                                                                                                        v4 -= (v4 >> ((2) & 63))
                                                                                                        var v123 int32 = 2147483647
                                                                                                        _ = v123
                                                                                                    } else {
                                                                                                        if true {
                                                                                                            if false {
                                                                                                                v124 := GenS1{F0: (-(fi21))}
                                                                                                                _ = v124
                                                                                                            }
                                                                                                            fmt.Println(fi16)
                                                                                                            var v125 float32 = v69
                                                                                                            _ = v125
                                                                                                        } else {
                                                                                                            var v126 int = (fi21 / (((fi16 >> ((18446744073709551615) & 63))) | 1))
                                                                                                            var v127 uint16 = 42
                                                                                                            v128 := GenS5{F0: v4}
                                                                                                            _ = v126
                                                                                                            _ = v127
                                                                                                            _ = v128
                                                                                                        }
                                                                                                        fmt.Println(fi16, v3)
                                                                                                        v129 := GenS2{F0: -1, F1: (v70 << ((1) & 63))}
                                                                                                        var v130 uint32 = 753
                                                                                                        _ = v129
                                                                                                        _ = v130
                                                                                                    }
                                                                                                    fmt.Println(v3, v70)
                                                                                                }
                                                                                            case 127:
                                                                                                v131 := GenS15{F0: (v70 * 0)}
                                                                                                _ = v131
                                                                                            }
                                                                                        } else {
                                                                                            v132 := v2[("" + "hello")]
                                                                                            var v133 float64 = 1.5
                                                                                            _ = v132
                                                                                            _ = v133
                                                                                        }
                                                                                        _ = v111
                                                                                    } else {
                                                                                        v134 := GenS7{F0: uint16(v70), F1: 42, F2: (^(v70))}
                                                                                        v135 := GenS4{F0: -32768, F1: false, F2: 4294967295}
                                                                                        _ = v134
                                                                                        _ = v135
                                                                                    }
                                                                                }
                                                                                v1 += 1
                                                                                var v136 int16 = -1
                                                                                fmt.Println(v1)
                                                                                _ = v136
                                                                            } else {
                                                                                v69 = -2.25
                                                                                if true {
                                                                                    switch "foo bar" {
                                                                                    case "hello":
                                                                                        var v137 uint = 4294967295
                                                                                        _ = v137
                                                                                    case "":
                                                                                        v70 -= (v1 & v70)
                                                                                    case "foo bar":
                                                                                        if (!false) {
                                                                                            fmt.Println(v3)
                                                                                            v138 := []int{}
                                                                                            _ = v138
                                                                                        } else {
                                                                                            var v139 uint8 = (^(v3))
                                                                                            v140 := v2[("foo bar" + "foo bar")]
                                                                                            _ = v139
                                                                                            _ = v140
                                                                                        }
                                                                                    default:
                                                                                        var v141 float64 = (float64((v69 - v69)) - 0.0)
                                                                                        _ = v141
                                                                                    }
                                                                                    v142 := []int{fi16, fi21, fi21}
                                                                                    v3 -= 127
                                                                                    _ = v142
                                                                                }
                                                                                v69 = 0.1
                                                                                var v143 int64 = 0
                                                                                _ = v143
                                                                            }
                                                                            v3 += v3
                                                                            if ((!((^(v4)) == v4)) && false) {
                                                                                v1 = v70
                                                                            }
                                                                        }
                                                                    }
                                                                case "\n\t":
                                                                    v144 := GenS3{F0: 318, F1: 423}
                                                                    _ = v144
                                                                case "":
                                                                    v145 := []int{(fi16 | (fi21 / ((-794) | 1)))}
                                                                    _ = v145
                                                                }
                                                                v146 := v2["hello"]
                                                                switch fi21 {
                                                                case 32767:
                                                                    var v147 int8 = 100
                                                                    _ = v147
                                                                case -2147483648:
                                                                    switch ((v146 + v146) + ("" + "\n\t")) {
                                                                    case "foo bar":
                                                                        v148 := GenS13{F0: "\n\t", F1: false, F2: (v70 & (^(v1)))}
                                                                        _ = v148
                                                                    case "a":
                                                                        if ((v70 / (((v1 / ((v1) | 1))) | 1)) <= uint64(v3)) {
                                                                            fmt.Println(v4)
                                                                            v70 = uint64(v3)
                                                                            v69 = float32(3.14159)
                                                                            switch v146 {
                                                                            case "foo bar":
                                                                                switch fi21 {
                                                                                case -166:
                                                                                    if true {
                                                                                        if false {
                                                                                            if true {
                                                                                                v2[(v146 + v146)] = "\n\t"
                                                                                                if false {
                                                                                                    switch ((("foo bar" + v146) + v146) + "foo bar") {
                                                                                                    case "a":
                                                                                                        var v149 uint64 = v1
                                                                                                        _ = v149
                                                                                                    case "hello":
                                                                                                        delete(v2, ("hello" + v146))
                                                                                                    case "foo bar":
                                                                                                        if false {
                                                                                                            if (false || true) {
                                                                                                                fmt.Println(v69, v70)
                                                                                                                v1 += (v1 | v1)
                                                                                                                v150 := GenS1{F0: (fi16 & 32768)}
                                                                                                                _ = v150
                                                                                                            } else {
                                                                                                                if false {
                                                                                                                    fmt.Println(v70, v1, v4)
                                                                                                                    var v151 float32 = (-2.25 + 1e10)
                                                                                                                    var v152 int = -2147483648
                                                                                                                    _ = v151
                                                                                                                    _ = v152
                                                                                                                } else {
                                                                                                                    v4 -= 1
                                                                                                                    v153 := GenS6{F0: 42, F1: 140, F2: 25}
                                                                                                                    _ = v153
                                                                                                                }
                                                                                                            }
                                                                                                            var v154 float32 = (v69 - (v69 * (((v69 / v69) * -2.25) * v69)))
                                                                                                            var v155 uint = 65535
                                                                                                            _ = v154
                                                                                                            _ = v155
                                                                                                        } else {
                                                                                                            fmt.Println(fi21, v4, fi16)
                                                                                                            fmt.Println(v4)
                                                                                                            v4 -= 42
                                                                                                        }
                                                                                                    default:
                                                                                                        if true {
                                                                                                            if false {
                                                                                                                var v156 int = fi16
                                                                                                                _ = len(v2)
                                                                                                                var v157 uint16 = 2
                                                                                                                switch (("foo bar" + v146) + v146) {
                                                                                                                case "hello":
                                                                                                                    var v158 uint8 = v3
                                                                                                                    _ = v158
                                                                                                                case "foo bar":
                                                                                                                    delete(v2, "hello")
                                                                                                                case "\n\t":
                                                                                                                    v69 = v69
                                                                                                                default:
                                                                                                                    if false {
                                                                                                                        v69 += float32(((0.1 / 1.0) - float64((1.5 - v69))))
                                                                                                                        if (fi21 <= (-(fi21))) {
                                                                                                                            v159 := GenS1{F0: (v156 & (fi21 << ((42) & 63)))}
                                                                                                                            v160 := GenS11{F0: false, F1: 58}
                                                                                                                            _ = v159
                                                                                                                            _ = v160
                                                                                                                        } else {
                                                                                                                            fmt.Println(v157, v4)
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                                v161 := GenS12{F0: v157, F1: v156}
                                                                                                                v162 := []int{(fi21 << ((4294967295) & 63))}
                                                                                                                _ = v156
                                                                                                                _ = v157
                                                                                                                _ = v161
                                                                                                                _ = v162
                                                                                                            } else {
                                                                                                                v163 := GenS10{F0: ("" <= v146)}
                                                                                                                v164 := GenS16{F0: fi21, F1: 4294967295}
                                                                                                                _ = v163
                                                                                                                _ = v164
                                                                                                            }
                                                                                                        }
                                                                                                        var v165 uint16 = 65535
                                                                                                        v1 -= 18446744073709551615
                                                                                                        _ = v165
                                                                                                    }
                                                                                                    if ("foo bar" > v146) {
                                                                                                        fmt.Println(fi16, v4)
                                                                                                    }
                                                                                                    v146 = ((((v146 + v146) + v146) + "") + v146)
                                                                                                } else {
                                                                                                    switch (fi16 - 42) {
                                                                                                    case -35:
                                                                                                        v4 -= (v4 >> ((3) & 63))
                                                                                                    default:
                                                                                                        v166 := []int{(fi16 % ((fi16) | 1)), (fi21 / ((fi21) | 1)), 9223372036854775807}
                                                                                                        _ = v166
                                                                                                    }
                                                                                                }
                                                                                                v69 += (v69 / v69)
                                                                                            } else {
                                                                                                if true {
                                                                                                    v167 := v2[v146]
                                                                                                    var v168 bool = false
                                                                                                    var v169 int = (fi21 | -2147483648)
                                                                                                    _ = v167
                                                                                                    _ = v168
                                                                                                    _ = v169
                                                                                                }
                                                                                                var v170 int64 = -908
                                                                                                _ = v170
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                case 127:
                                                                                    v171 := []int{fi21, (fi16 * int(v1))}
                                                                                    _ = v171
                                                                                case -2147483648:
                                                                                    var v172 uint = 65535
                                                                                    var v173 uint32 = 1
                                                                                    _ = v172
                                                                                    _ = v173
                                                                                }
                                                                            }
                                                                            switch fi16 {
                                                                            case 42:
                                                                                if true {
                                                                                    if true {
                                                                                        v174 := []int{4294967295}
                                                                                        if ("" < "a") {
                                                                                            var v175 float32 = v69
                                                                                            if (0.1 > v69) {
                                                                                                if (true && false) {
                                                                                                    _ = len(v174)
                                                                                                }
                                                                                            } else {
                                                                                                var v176 string = v146
                                                                                                v175 = (3.14159 / v175)
                                                                                                v177 := GenS17{F0: 2147483647}
                                                                                                _ = v176
                                                                                                _ = v177
                                                                                            }
                                                                                            if false {
                                                                                                if false {
                                                                                                    var v178 bool = true
                                                                                                    var v179 uint8 = (v3 / ((genFn0(32767, 631)) | 1))
                                                                                                    _ = v178
                                                                                                    _ = v179
                                                                                                } else {
                                                                                                    if len(v174) >= 2 {
                                                                                                        v174 = v174[1:len(v174)]
                                                                                                    }
                                                                                                    if len(v174) >= 2 {
                                                                                                        v174 = v174[1:len(v174)]
                                                                                                    }
                                                                                                }
                                                                                                if false {
                                                                                                    if (((v3 > (^(v3))) && true) && (v4 > v4)) {
                                                                                                        switch "foo bar" {
                                                                                                        case "hello":
                                                                                                            delete(v2, "hello")
                                                                                                        case "":
                                                                                                            var v180 int8 = 1
                                                                                                            _ = v180
                                                                                                        }
                                                                                                        if (false || false) {
                                                                                                            var v181 float64 = (float64(v69) / 1.0)
                                                                                                            _ = v181
                                                                                                        } else {
                                                                                                            v3 -= genFn0(-1, 65535)
                                                                                                        }
                                                                                                        switch fi21 {
                                                                                                        case 32768:
                                                                                                            v69 -= v69
                                                                                                        case 9223372036854775807:
                                                                                                            var v182 int8 = 0
                                                                                                            if (!(2 >= 1)) {
                                                                                                                var v183 float64 = 1e10
                                                                                                                fmt.Println(v69, v1, v183)
                                                                                                                if true {
                                                                                                                    v69 -= (v69 - (1e10 + v69))
                                                                                                                }
                                                                                                                _ = v183
                                                                                                            }
                                                                                                            v2["a"] = "foo bar"
                                                                                                            _ = v182
                                                                                                        default:
                                                                                                            fmt.Println(v69, v4, v69)
                                                                                                        }
                                                                                                    } else {
                                                                                                        if ((fi21 * fi21) <= fi16) {
                                                                                                            var v184 float32 = ((0.1 - v69) + 0.1)
                                                                                                            _ = v184
                                                                                                        }
                                                                                                        fmt.Println(fi16, fi21)
                                                                                                        var v185 int16 = -32768
                                                                                                        _ = v185
                                                                                                    }
                                                                                                    if true {
                                                                                                        fmt.Println(v1, v4)
                                                                                                        switch (v146 + v146) {
                                                                                                        case "a":
                                                                                                            v1 = v70
                                                                                                        default:
                                                                                                            fmt.Println(v1, v69, v70)
                                                                                                        }
                                                                                                        v175 -= 3.14159
                                                                                                        v186 := GenS12{F0: 2, F1: fi21}
                                                                                                        _ = v186
                                                                                                    } else {
                                                                                                        var v187 int64 = -1
                                                                                                        v146 = "\n\t"
                                                                                                        _ = v187
                                                                                                    }
                                                                                                } else {
                                                                                                    v1 = (v1 - v70)
                                                                                                    if false {
                                                                                                        v69 = v69
                                                                                                    } else {
                                                                                                        var v188 float64 = 3.14159
                                                                                                        v189 := GenS7{F0: 1, F1: (fi21 % ((61) | 1)), F2: 233}
                                                                                                        _ = v188
                                                                                                        _ = v189
                                                                                                    }
                                                                                                }
                                                                                                v3 = 44
                                                                                                v69 = (v175 + 1.5)
                                                                                                var v190 uint64 = 42
                                                                                                _ = v190
                                                                                            }
                                                                                            _ = v175
                                                                                        }
                                                                                        if false {
                                                                                            if true {
                                                                                                v3 += genFn0(32767, 42)
                                                                                                v70 += 82
                                                                                                fmt.Println(v70, v4)
                                                                                                switch "\n\t" {
                                                                                                case "\n\t":
                                                                                                    v4 = 847
                                                                                                case "hello":
                                                                                                    if (70 < 761) {
                                                                                                        v191 := GenS12{F0: 484, F1: fi16}
                                                                                                        var v192 float32 = 0.0
                                                                                                        v4 -= v4
                                                                                                        _ = v191
                                                                                                        _ = v192
                                                                                                    }
                                                                                                }
                                                                                                switch (fi21 / ((-2147483648) | 1)) {
                                                                                                case 127:
                                                                                                    if len(v174) >= 2 {
                                                                                                        v174 = v174[1:len(v174)]
                                                                                                    }
                                                                                                case -948:
                                                                                                    v174 = append(v174, (fi16 << ((2) & 63)))
                                                                                                }
                                                                                                v70 -= (v70 % (((v70 * (^(v1)))) | 1))
                                                                                                var v193 int8 = -1
                                                                                                _ = v193
                                                                                            } else {
                                                                                                v194 := GenS4{F0: 32767, F1: false, F2: 42}
                                                                                                _ = v194
                                                                                            }
                                                                                        } else {
                                                                                            var v195 uint16 = 455
                                                                                            _ = v195
                                                                                        }
                                                                                        switch (fi16 / ((9223372036854775807) | 1)) {
                                                                                        case 814:
                                                                                            if false {
                                                                                                var v196 bool = false
                                                                                                v197 := GenS4{F0: 500, F1: (!true), F2: 1}
                                                                                                v146 = v146
                                                                                                _ = v196
                                                                                                _ = v197
                                                                                            } else {
                                                                                                switch fi21 {
                                                                                                case 2147483647:
                                                                                                    switch "hello" {
                                                                                                    case "":
                                                                                                        v198 := GenS18{F0: true, F1: 2, F2: 608}
                                                                                                        _ = v198
                                                                                                    case "hello":
                                                                                                        v199 := GenS18{F0: ((fi16 / ((206) | 1)) >= -869), F1: 1, F2: 268}
                                                                                                        _ = v199
                                                                                                    default:
                                                                                                        if len(v174) >= 2 {
                                                                                                            v174 = v174[1:len(v174)]
                                                                                                        }
                                                                                                    }
                                                                                                case 127:
                                                                                                    var v200 uint64 = (v70 / ((2) | 1))
                                                                                                    _ = v200
                                                                                                default:
                                                                                                    var v201 int8 = int8(fi21)
                                                                                                    _ = v201
                                                                                                }
                                                                                                _ = len(v174)
                                                                                                v202 := GenS1{F0: -9223372036854775808}
                                                                                                _ = v202
                                                                                            }
                                                                                        case 650:
                                                                                            v70 = v1
                                                                                        case -510:
                                                                                            var v203 uint64 = 1
                                                                                            _ = v203
                                                                                        default:
                                                                                            fmt.Println(fi16)
                                                                                        }
                                                                                        _ = v174
                                                                                    } else {
                                                                                        v70 = (^(v1))
                                                                                        var v204 uint64 = v70
                                                                                        var v205 int64 = 975
                                                                                        _ = v204
                                                                                        _ = v205
                                                                                    }
                                                                                }
                                                                            case 4294967295:
                                                                                v70 = (v1 & v70)
                                                                            }
                                                                        }
                                                                    case "":
                                                                        switch v146 {
                                                                        case "hello":
                                                                            v2[v146] = ("hello" + v146)
                                                                        }
                                                                        var v206 uint16 = 42
                                                                        fmt.Println(v3, v4, fi16)
                                                                        _ = v206
                                                                    }
                                                                case 4294967295:
                                                                    if (false || true) {
                                                                        if false {
                                                                            v207 := GenS19{F0: (v3 + (v3 | v3))}
                                                                            switch "foo bar" {
                                                                            case "\n\t":
                                                                                if (!true) {
                                                                                    v4 -= -2147483648
                                                                                    if (!true) {
                                                                                        switch "a" {
                                                                                        case "foo bar":
                                                                                            v2["foo bar"] = v146
                                                                                        case "a":
                                                                                            fmt.Println(v3)
                                                                                        case "":
                                                                                            v208 := GenS20{F0: 65535, F1: uint16(v70)}
                                                                                            _ = v208
                                                                                        default:
                                                                                            v3 -= v3
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    var v209 uint = 65535
                                                                                    _ = v209
                                                                                }
                                                                            default:
                                                                                v3 = genFn0(1000, 2147483647)
                                                                            }
                                                                            _ = v207
                                                                        } else {
                                                                            v69 = 0.0
                                                                            fmt.Println(fi16)
                                                                            fmt.Println(v1, v70, v3)
                                                                        }
                                                                        var v210 float64 = ((float64((1.5 * ((float32(1.5) - (v69 - 0.1)) + ((v69 + (v69 - float32((1.5 / 3.14159)))) * (3.14159 * v69))))) - (-2.25 * 1e10)) / -2.25)
                                                                        _ = v210
                                                                    } else {
                                                                        v211 := v2[v146]
                                                                        v4 = (v4 % ((v4) | 1))
                                                                        _ = v211
                                                                    }
                                                                }
                                                                _ = v146
                                                            }
                                                        default:
                                                            switch "hello" {
                                                            case "\n\t":
                                                                v212 := GenS1{F0: (fi21 * (fi16 | (fi21 << ((746) & 63))))}
                                                                _ = v212
                                                            case "foo bar":
                                                                var v213 int = (fi16 - (fi16 ^ (-(fi16))))
                                                                _ = v213
                                                            case "":
                                                                v214 := []int{}
                                                                _ = v214
                                                            }
                                                            v69 = ((1e10 + (v69 + 1.5)) - v69)
                                                            if true {
                                                                v215 := GenS21{F0: -128, F1: fi16}
                                                                v70 = (v1 ^ v70)
                                                                _ = v215
                                                            } else {
                                                                fmt.Println(fi21)
                                                                if false {
                                                                    var v216 float32 = v69
                                                                    v216 = -2.25
                                                                    _ = v216
                                                                } else {
                                                                    var v217 string = "foo bar"
                                                                    _ = v217
                                                                }
                                                                var v218 uint8 = (^(v3))
                                                                _ = v218
                                                            }
                                                            v1 -= (v1 | (v1 / (((^(v1))) | 1)))
                                                        }
                                                        v219 := GenS6{F0: -129, F1: 643, F2: -128}
                                                        var v220 uint32 = 0
                                                        _ = v219
                                                        _ = v220
                                                        _ = v70
                                                    }
                                                    _ = v69
                                                } else {
                                                    switch ("a" + "a") {
                                                    case "hello":
                                                        v4 = v4
                                                    case "a":
                                                        v221 := GenS22{F0: -91}
                                                        _ = v221
                                                    case "":
                                                        fmt.Println(v3, fi21, fi21)
                                                    default:
                                                        var v222 uint32 = 4294967295
                                                        _ = v222
                                                    }
                                                    v1 += (v1 - v1)
                                                    v223 := GenS21{F0: 1, F1: int(v3)}
                                                    _ = v223
                                                }
                                                switch "\n\t" {
                                                case "\n\t":
                                                    fmt.Println(fi16, fi16, fi16)
                                                case "":
                                                    switch "a" {
                                                    case "hello":
                                                        v224 := []int{1, -9223372036854775808, 127}
                                                        _ = v224
                                                    case "foo bar":
                                                        v225 := GenS19{F0: v3}
                                                        _ = v225
                                                    }
                                                case "foo bar":
                                                    fmt.Println(v3, fi21)
                                                }
                                            }
                                            v226 := []int{fi21, fi21, (fi16 - fi16)}
                                            _ = v226
                                        }
                                    }
                                    var v227 float32 = 0.1
                                    if (!(!false)) {
                                        switch ("foo bar" + "") {
                                        case "foo bar":
                                            if (((fi16 >> ((1) & 63)) > fi21) || ((float32(0.1) == 1.5) && false)) {
                                                var v228 int = 2147483647
                                                v229 := v2["\n\t"]
                                                _ = v228
                                                _ = v229
                                            }
                                        default:
                                            var v230 uint16 = 271
                                            _ = v230
                                        }
                                    } else {
                                        v1 = (v1 * (v1 << ((42) & 63)))
                                        v4 = v4
                                        switch (fi21 & fi16) {
                                        case -1:
                                            fmt.Println(v1, v3)
                                        case 2147483647:
                                            var v231 float64 = 1e10
                                            _ = v231
                                        }
                                    }
                                    if (!(879 != 1)) {
                                        if true {
                                            v227 -= float32((0.0 - (1.5 * (((0.1 + 3.14159) - float64(v227)) * 1e10))))
                                            v232 := GenS21{F0: 123, F1: 691}
                                            var v233 int8 = -1
                                            var v234 int8 = (v233 ^ 1)
                                            v4 += 2147483647
                                            v234 = (v233 ^ (v233 - v233))
                                            _ = v232
                                            _ = v233
                                            _ = v234
                                        } else {
                                            v3 = (^(v3))
                                            v1 -= (v1 * 18446744073709551615)
                                            v4 = v4
                                        }
                                        if (491 != 255) {
                                            if (!false) {
                                                if (9223372036854775807 >= 9223372036854775807) {
                                                    fmt.Println(v227)
                                                    v235 := GenS22{F0: 0}
                                                    _ = v235
                                                } else {
                                                    v236 := []int{}
                                                    _ = v236
                                                }
                                                fmt.Println(fi16)
                                                var v237 uint32 = 42
                                                _ = v237
                                            }
                                            switch fi16 {
                                            case 603:
                                                if true {
                                                    switch "a" {
                                                    case "a":
                                                        v238 := GenS20{F0: 18446744073709551615, F1: 255}
                                                        _ = v238
                                                    }
                                                    if true {
                                                        v239 := []int{fi16, fi21, (fi16 % (((fi16 % ((-70) | 1))) | 1))}
                                                        v227 -= v227
                                                        var v240 uint16 = 65535
                                                        _ = v239
                                                        _ = v240
                                                    }
                                                    var v241 uint64 = 42
                                                    v1 += v1
                                                    fmt.Println(fi16, v4, fi21)
                                                    _ = v241
                                                }
                                            }
                                            v242 := GenS23{F0: (^(v1)), F1: (-9223372036854775808 > -1)}
                                            _ = v242
                                        }
                                    }
                                    var v243 int16 = 42
                                    _ = fi21
                                    _ = v227
                                    _ = v243
                                }
                            default:
                                if true {
                                    var v244 uint16 = 2
                                    v4 -= v4
                                    for fi245 := 0; fi245 < 15; fi245++ {
                                        switch fi16 {
                                        case 811:
                                            v246 := GenS6{F0: -1, F1: 0, F2: 127}
                                            _ = v246
                                        case -719:
                                            v247 := []int{(fi245 << ((255) & 63)), (fi16 >> ((2) & 63)), fi245}
                                            _ = v247
                                        case 226:
                                            v248 := v2["foo bar"]
                                            _ = v248
                                        }
                                        _ = fi245
                                    }
                                    _ = v244
                                } else {
                                    delete(v2, "a")
                                    for fi249 := 0; fi249 < 3; fi249++ {
                                        v250 := []int{fi16, fi249}
                                        _ = fi249
                                        _ = v250
                                    }
                                    fmt.Println(fi16, v1)
                                }
                            }
                        case "foo bar":
                            fmt.Println(v1, v4)
                        case "":
                            fmt.Println(fi16, v1, v4)
                        default:
                            v3 = (v3 + genFn0(404, 388))
                        }
                        var v251 float32 = 0.0
                        _ = v251
                    } else {
                        for fi252 := 0; fi252 < 16; fi252++ {
                            var v253 int32 = (^(v4))
                            v2[("a" + "")] = "\n\t"
                            _ = fi252
                            _ = v253
                        }
                        for fi254 := 0; fi254 < 5; fi254++ {
                            v255 := GenS1{F0: fi254}
                            var v256 int = fi254
                            var v257 uint16 = 660
                            _ = fi254
                            _ = v255
                            _ = v256
                            _ = v257
                        }
                    }
                    delete(v2, "a")
                }
                _ = fi16
            }
        } else {
            v4 = 0
            for fi258 := 0; fi258 < 18; fi258++ {
                var v259 float64 = 0.1
                if (false || true) {
                    v260 := GenS24{F0: ("foo bar" + (("hello" + "a") + ("foo bar" + "foo bar"))), F1: 4294967295}
                    for fi261 := 0; fi261 < 10; fi261++ {
                        if (0 <= -32768) {
                            if (false && ((1 > uint(v3)) && false)) {
                                var v262 uint8 = (^(v3))
                                var v263 uint = 42
                                _ = v262
                                _ = v263
                            } else {
                                v264 := GenS25{F0: genFn0(-1, 2147483647), F1: (fi261 / ((4294967295) | 1)), F2: 0}
                                fmt.Println(v1, fi261)
                                _ = v264
                            }
                            var v265 bool = (((1.5 + (float32(v259) / 1.5)) * -2.25) != ((float32(-2.25) + float32(0.0)) / 1.0))
                            _ = len(v2)
                            _ = v265
                        } else {
                            fmt.Println(fi261)
                        }
                        var v266 int32 = (v4 % (((v4 ^ -1)) | 1))
                        _ = fi261
                        _ = v266
                    }
                    _ = v260
                } else {
                    v267 := GenS26{F0: (fi258 | fi258)}
                    var v268 float64 = (v259 - v259)
                    fmt.Println(v1, fi258)
                    _ = v267
                    _ = v268
                }
                var v269 int64 = -1
                _ = fi258
                _ = v259
                _ = v269
            }
            v270 := GenS17{F0: (v4 >> ((65535) & 63))}
            _ = v270
        }
    } else {
        v271 := GenS27{F0: 2147483648, F1: 65535, F2: 255}
        _ = v271
    }
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v8
}
