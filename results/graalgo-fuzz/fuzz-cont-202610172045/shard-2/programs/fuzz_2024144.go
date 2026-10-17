package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 int; F2 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64; F1 int32; F2 uint }
type GenS5 struct { F0 uint16; F1 uint8 }
type GenS6 struct { F0 uint32; F1 int8 }
type GenS7 struct { F0 uint16 }
func genFn0(p0 int16, p1 bool, p2 uint) int8 {
    fmt.Println(p1, p0, p1)
    var v5 bool = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 127
}
func genFn1(p0 uint32) int16 {
    var v6 uint = 0
    v6 = v6
    for fi7 := 0; fi7 < 17; fi7++ {
        p0 = (p0 / (((p0 << (((v6 << (((v6 & (v6 & (^(v6))))) & 63))) & 63))) | 1))
        if (int(p0) != int(p0)) {
            var v8 bool = ((false || true) && (true && (true || false)))
            v8 = ((v8 || (-32768 != -937)) || true)
            _ = v8
        } else {
            for fi9 := 0; fi9 < 1; fi9++ {
                var v10 int16 = 32767
                _ = fi9
                _ = v10
            }
            v11 := []int{0, fi7, (fi7 ^ -2147483648)}
            _ = v11
        }
        fmt.Println(v6, fi7, fi7)
        _ = fi7
    }
    _ = p0
    _ = v6
    return 1000
}

func main() {
    var v1 uint32 = 797
    var v2 int16 = 1000
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]uint32)
    for fi12 := 0; fi12 < 11; fi12++ {
        if (!false) {
            v13 := GenS1{F0: (3.14159 == (0.0 * (3.14159 - 0.0)))}
            fmt.Println(fi12, v1, fi12)
            fmt.Println(fi12)
            _ = v13
        } else {
            var v14 uint64 = 757
            _ = len(v4)
            _ = v14
        }
        if (-9223372036854775808 < 42) {
            v15 := []int{fi12, fi12, (fi12 - 32767), 32767}
            _ = v15
        }
        _ = fi12
    }
    if (42 >= 142) {
        for fi16 := 0; fi16 < 9; fi16++ {
            v17 := GenS1{F0: (2 > 1)}
            delete(v3, -2147483648)
            if ((v1 + (v1 * (^(v1)))) > (v1 >> ((2) & 63))) {
                switch "hello" {
                case "hello":
                    fmt.Println(v2, v2, v1)
                }
            }
            _ = fi16
            _ = v17
        }
        fmt.Println(v2)
        switch ("a" + "\n\t") {
        case "\n\t":
            v2 = (v2 % (((v2 << ((384) & 63))) | 1))
        }
    } else {
        v1 = v1
        if (!false) {
            switch -1 {
            case -1:
                v1 = (v1 & (^(v1)))
            case 404:
                v1 -= v1
            }
            v1 += 531
        }
        v2 = (v2 & (v2 | (v2 >> ((18446744073709551615) & 63))))
    }
    v18 := GenS2{F0: 1, F1: -9223372036854775808, F2: 77}
    if false {
        if true {
            if (false || ((-2.25 * 1e10) > 1.5)) {
                var v19 int32 = 0
                _ = v19
            }
            v2 -= genFn1(2147483648)
        }
        for fi20 := 0; fi20 < 7; fi20++ {
            switch fi20 {
            case 32768:
                var v21 int16 = (^(v2))
                _ = v21
            }
            _ = fi20
        }
    } else {
        if (-966 == 1) {
            for fi22 := 0; fi22 < 19; fi22++ {
                fmt.Println(v1)
                var v23 int16 = v2
                _ = fi22
                _ = v23
            }
        }
    }
    v24 := GenS3{F0: 837}
    switch "\n\t" {
    case "a":
        if true {
            v2 -= (^(v2))
            switch ("\n\t" + "") {
            case "hello":
                var v25 float32 = (0.0 - (3.14159 + (1e10 - 3.14159)))
                _ = v25
            case "a":
                switch "hello" {
                case "a":
                    if false {
                        if true {
                            v2 -= (v2 * v2)
                            v2 += genFn1(1)
                            v26 := GenS2{F0: 1, F1: 9223372036854775807, F2: -1}
                            _ = v26
                        } else {
                            v27 := []int{-763}
                            if (!true) {
                                switch "" {
                                case "a":
                                    switch 372 {
                                    case 944:
                                        if false {
                                            v1 = (v1 % (((v1 | v1)) | 1))
                                        }
                                    case 2147483647:
                                        for fi28 := 0; fi28 < 16; fi28++ {
                                            v1 = (v1 * (v1 ^ v1))
                                            if (!(255 != 1)) {
                                                for fi29 := 0; fi29 < 6; fi29++ {
                                                    v3[fi29] = 2
                                                    _ = fi29
                                                }
                                            }
                                            _ = fi28
                                        }
                                    case 1:
                                        if len(v27) > 0 {
                                            v27[0] = -1
                                        }
                                    default:
                                        if (false && (2147483647 >= 58)) {
                                            switch "foo bar" {
                                            case "":
                                                for fi30 := 0; fi30 < 6; fi30++ {
                                                    if (776 != 610) {
                                                        v31 := GenS2{F0: 42, F1: fi30, F2: 127}
                                                        v2 += (v2 >> ((18446744073709551615) & 63))
                                                        _ = v31
                                                    }
                                                    v32 := GenS4{F0: 42, F1: 2147483647, F2: 255}
                                                    for fi33 := 0; fi33 < 7; fi33++ {
                                                        var v34 string = "\n\t"
                                                        v34 = v34
                                                        v35 := GenS3{F0: (fi33 / ((fi30) | 1))}
                                                        _ = fi33
                                                        _ = v34
                                                        _ = v35
                                                    }
                                                    _ = fi30
                                                    _ = v32
                                                }
                                            case "\n\t":
                                                var v36 float64 = 3.14159
                                                _ = v36
                                            case "foo bar":
                                                var v37 uint64 = 803
                                                _ = v37
                                            }
                                        } else {
                                            var v38 uint16 = 42
                                            _ = v38
                                        }
                                    }
                                case "":
                                    for fi39 := 0; fi39 < 19; fi39++ {
                                        v2 += -32768
                                        for fi40 := 0; fi40 < 7; fi40++ {
                                            var v41 uint32 = v1
                                            var v42 int32 = 1
                                            v1 -= v41
                                            _ = fi40
                                            _ = v41
                                            _ = v42
                                        }
                                        _ = fi39
                                    }
                                }
                            }
                            _ = v27
                        }
                        v2 = v2
                        v2 = -242
                    } else {
                        v43 := GenS1{F0: false}
                        for fi44 := 0; fi44 < 3; fi44++ {
                            var v45 int8 = -1
                            if false {
                                for fi46 := 0; fi46 < 11; fi46++ {
                                    if (-9223372036854775808 != -9223372036854775808) {
                                        switch 1 {
                                        case 32768:
                                            delete(v3, (^(fi44)))
                                        case 4294967295:
                                            v47 := GenS5{F0: 42, F1: 255}
                                            if false {
                                                if true {
                                                    var v48 int = (fi44 << ((319) & 63))
                                                    _ = v48
                                                }
                                            }
                                            var v49 uint8 = 128
                                            switch "\n\t" {
                                            case "foo bar":
                                                var v50 float64 = 1e10
                                                _ = v50
                                            case "":
                                                v49 -= v49
                                            }
                                            v51 := GenS1{F0: (2 < 32767)}
                                            _ = v47
                                            _ = v49
                                            _ = v51
                                        case 2147483647:
                                            if true {
                                                fmt.Println(fi44)
                                            }
                                        default:
                                            var v52 uint8 = 2
                                            _ = v52
                                        }
                                        v53 := []int{}
                                        delete(v4, (v1 ^ v1))
                                        _ = v53
                                    } else {
                                        v45 -= v45
                                        if (4294967295 >= uint64(v1)) {
                                            delete(v3, fi44)
                                            if (!true) {
                                                switch (fi46 - fi46) {
                                                case -748:
                                                    var v54 string = "\n\t"
                                                    v45 = v45
                                                    _ = v54
                                                case 1:
                                                    if (32767 > 55) {
                                                        var v55 int16 = 1000
                                                        _ = v55
                                                    }
                                                case 42:
                                                    if true {
                                                        v56 := []int{fi44, fi46, (fi46 >> ((692) & 63))}
                                                        if true {
                                                            var v57 uint16 = 32767
                                                            switch "" {
                                                            case "":
                                                                v1 = (v1 << ((716) & 63))
                                                            case "foo bar":
                                                                v58 := GenS3{F0: fi46}
                                                                _ = v58
                                                            case "hello":
                                                                fmt.Println(v2, fi46, fi46)
                                                            default:
                                                                fmt.Println(v1, v1, v57)
                                                            }
                                                            v57 = v57
                                                            if (4294967295 <= 18446744073709551615) {
                                                                var v59 uint = 255
                                                                fmt.Println(v1)
                                                                if false {
                                                                    var v60 int64 = 863
                                                                    _ = v60
                                                                } else {
                                                                    v61 := v3[123]
                                                                    var v62 bool = (128 < 255)
                                                                    _ = v61
                                                                    _ = v62
                                                                }
                                                                _ = v59
                                                            }
                                                            _ = v57
                                                        }
                                                        _ = v56
                                                    } else {
                                                        var v63 uint8 = 0
                                                        v2 += v2
                                                        if true {
                                                            var v64 bool = (0 != 871)
                                                            var v65 int16 = (v2 / ((v2) | 1))
                                                            _ = v64
                                                            _ = v65
                                                        } else {
                                                            v66 := GenS5{F0: 32767, F1: (v63 & (v63 % (((^(v63))) | 1)))}
                                                            var v67 int32 = 42
                                                            _ = v66
                                                            _ = v67
                                                        }
                                                        _ = v63
                                                    }
                                                }
                                                v68 := []int{}
                                                v1 = 42
                                                _ = v68
                                            }
                                        }
                                        var v69 string = "hello"
                                        _ = v69
                                    }
                                    var v70 int64 = 432
                                    _ = fi46
                                    _ = v70
                                }
                                v1 = v1
                            }
                            _ = fi44
                            _ = v45
                        }
                        _ = v43
                    }
                case "":
                    var v71 uint64 = 1
                    _ = v71
                case "hello":
                    v2 -= (v2 & 32767)
                default:
                    for fi72 := 0; fi72 < 16; fi72++ {
                        v2 = v2
                        for fi73 := 0; fi73 < 13; fi73++ {
                            fmt.Println(fi73)
                            _ = fi73
                        }
                        _ = fi72
                    }
                }
            }
            switch 97 {
            case 42:
                switch 770 {
                case 1:
                    fmt.Println(v1, v1)
                case 0:
                    v1 -= (^(v1))
                default:
                    if false {
                        v74 := GenS4{F0: 1, F1: 609, F2: 0}
                        v4[(v1 * v1)] = v1
                        _ = v74
                    }
                }
            case -2147483648:
                v75 := GenS1{F0: (127 >= 78)}
                _ = v75
            }
        }
    default:
        switch 32768 {
        case -176:
            for fi76 := 0; fi76 < 17; fi76++ {
                if ((1.5 / 1.0) < (3.14159 / 1.0)) {
                    var v77 uint16 = 237
                    if true {
                        var v78 int8 = -83
                        _ = v78
                    }
                    for fi79 := 0; fi79 < 19; fi79++ {
                        v77 = 864
                        var v80 bool = true
                        v77 += (v77 % (((v77 >> ((42) & 63))) | 1))
                        _ = fi79
                        _ = v80
                    }
                    _ = v77
                } else {
                    var v81 int8 = 0
                    delete(v3, (fi76 & 32767))
                    v2 = genFn1(820)
                    _ = v81
                }
                _ = fi76
            }
        case 9223372036854775807:
            if (855 >= 2147483647) {
                for fi82 := 0; fi82 < 5; fi82++ {
                    switch 1 {
                    case 127:
                        switch "hello" {
                        case "\n\t":
                            if true {
                                if (2 >= 42) {
                                    switch (fi82 % ((795) | 1)) {
                                    case 572:
                                        var v83 int16 = int16(v1)
                                        _ = v83
                                    }
                                }
                                fmt.Println(v2, v2, v1)
                            }
                        case "hello":
                            v84 := GenS3{F0: fi82}
                            _ = v84
                        case "a":
                            for fi85 := 0; fi85 < 15; fi85++ {
                                if (1e10 < 0.1) {
                                    v86 := GenS1{F0: ("\n\t" == "foo bar")}
                                    var v87 string = ("foo bar" + "hello")
                                    v2 -= v2
                                    _ = v86
                                    _ = v87
                                } else {
                                    v88 := GenS5{F0: 255, F1: 1}
                                    v89 := GenS6{F0: (v1 | 42), F1: 42}
                                    _ = v88
                                    _ = v89
                                }
                                _ = fi85
                            }
                        }
                    case 32768:
                        v90 := GenS2{F0: 694, F1: 42, F2: -128}
                        _ = v90
                    default:
                        v1 += v1
                    }
                    var v91 bool = ((v2 != genFn1(2)) && (false || false))
                    _ = fi82
                    _ = v91
                }
            }
        default:
            v2 -= v2
        }
    }
    v92 := GenS1{F0: false}
    if false {
        v2 += v2
        var v93 uint8 = 128
        _ = v93
    } else {
        v94 := GenS7{F0: 255}
        v95 := GenS4{F0: -878, F1: -849, F2: 1}
        _ = v94
        _ = v95
    }
    v96 := []int{0, 32768, -263}
    v1 = 1
    switch ("foo bar" + ("" + (("\n\t" + "\n\t") + ("" + "")))) {
    case "":
        v1 = 652
    case "foo bar":
        v1 += (v1 >> ((4294967295) & 63))
    default:
        for fi97 := 0; fi97 < 15; fi97++ {
            var v98 uint8 = 42
            var v99 int = (fi97 << ((1) & 63))
            _ = fi97
            _ = v98
            _ = v99
        }
    }
    if false {
        v100 := GenS2{F0: 255, F1: 4294967295, F2: genFn0(v2, false, 188)}
        for fi101 := 0; fi101 < 2; fi101++ {
            for fi102 := 0; fi102 < 6; fi102++ {
                v3[fi102] = v1
                v1 += 407
                if false {
                    delete(v3, (fi101 >> ((65535) & 63)))
                }
                _ = fi102
            }
            v2 = v2
            _ = fi101
        }
        _ = v100
    } else {
        if false {
            fmt.Println(v1, v1)
            v103 := GenS5{F0: 255, F1: 127}
            _ = v103
        }
    }
    v104 := GenS4{F0: 42, F1: 2147483647, F2: 1}
    switch (("foo bar" + ("a" + "foo bar")) + "\n\t") {
    case "hello":
        if len(v96) > 0 {
            v96[0] = -1
        }
    case "a":
        if true {
            for fi105 := 0; fi105 < 13; fi105++ {
                v106 := GenS2{F0: 42, F1: fi105, F2: int8(v1)}
                _ = fi105
                _ = v106
            }
            if (2 > 18446744073709551615) {
                if len(v96) >= 2 {
                    v96 = v96[1:len(v96)]
                }
                v2 -= (v2 % (((^(v2))) | 1))
                switch 42 {
                case 127:
                    v1 -= v1
                case 128:
                    _ = len(v96)
                case 9223372036854775807:
                    var v107 int64 = -1
                    _ = v107
                }
            }
        } else {
            var v108 uint64 = 226
            v108 += 4294967295
            _ = v108
        }
    case "":
        var v109 int16 = (v2 + int16(v1))
        _ = v109
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v96:", len(v96), v96)
    _ = v1
    _ = v104
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v92
    _ = v96
}
