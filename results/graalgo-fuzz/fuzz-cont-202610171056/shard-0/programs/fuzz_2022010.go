package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 int64; F2 uint64 }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8; F1 int32; F2 int8 }
type GenS6 struct { F0 uint32; F1 uint; F2 int32 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int; F1 int8 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint64; F1 int64 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 int32; F1 int8 }
type GenS14 struct { F0 int16; F1 uint; F2 uint32 }
type GenS15 struct { F0 int }
type GenS16 struct { F0 int64; F1 uint; F2 uint }
type GenS17 struct { F0 uint }
type GenS18 struct { F0 uint; F1 bool; F2 uint }
func genFn0() uint8 {
    var v6 int8 = 1
    v7 := []int{9223372036854775807}
    v8 := GenS1{F0: -426}
    v7 = append(v7, -1)
    fmt.Println(v6)
    _ = v6
    _ = v7
    _ = v8
    return 146
}
func genFn1(p0 int32, p1 uint8, p2 uint64) int {
    for fi9 := 0; fi9 < 13; fi9++ {
        switch "a" {
        case "\n\t":
            var v10 uint64 = p2
            _ = v10
        case "a":
            v11 := []int{831}
            _ = v11
        case "foo bar":
            var v12 uint64 = 4294967295
            _ = v12
        }
        if true {
            if ((!false) || false) {
                var v13 float64 = (0.0 - 0.1)
                p1 -= (p1 % (((p1 | 255)) | 1))
                _ = v13
            }
        } else {
            switch "a" {
            case "hello":
                v14 := []int{42}
                _ = v14
            case "a":
                for fi15 := 0; fi15 < 9; fi15++ {
                    fmt.Println(p0, p2, fi15)
                    switch (fi9 ^ fi9) {
                    case -9223372036854775808:
                        v16 := []int{fi9}
                        _ = v16
                    }
                    _ = fi15
                }
            default:
                p2 += (^(p2))
            }
            switch "hello" {
            case "":
                var v17 uint8 = 127
                _ = v17
            default:
                var v18 uint32 = 65535
                _ = v18
            }
            switch ("" + "foo bar") {
            case "foo bar":
                var v19 uint32 = 2147483648
                _ = v19
            case "\n\t":
                v20 := []int{-2147483648, (fi9 * (fi9 | -9223372036854775808)), (fi9 >> ((18446744073709551615) & 63))}
                _ = v20
            }
        }
        _ = fi9
    }
    p0 -= (p0 % ((p0) | 1))
    _ = p0
    _ = p1
    _ = p2
    return 4294967295
}

func main() {
    var v1 int8 = 1
    v2 := []uint32{1}
    var v3 int = 32768
    v4 := []uint32{2, 2, 2147483647, 2147483648}
    v5 := []uint32{0, 2147483648, 224}
    for fi21 := 0; fi21 < 7; fi21++ {
        fmt.Println(fi21, v1)
        var v22 string = "a"
        _ = fi21
        _ = v22
    }
    fmt.Println(v3, v1)
    if ((1e10 >= (1e10 / ((0.1 - 1e10) + 3.14159))) || false) {
        var v23 int64 = -748
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        switch v3 {
        case -9223372036854775808:
            for fi24 := 0; fi24 < 19; fi24++ {
                v25 := GenS2{F0: (v3 & genFn1(int32(v3), 42, 42)), F1: (v23 % ((v23) | 1)), F2: 2}
                _ = fi24
                _ = v25
            }
        }
        _ = v23
    } else {
        for fi26 := 0; fi26 < 12; fi26++ {
            if true {
                v1 = (-(v1))
                for fi27 := 0; fi27 < 10; fi27++ {
                    v2 = append(v2, 2147483647)
                    _ = fi27
                }
                if (293 <= 289) {
                    var v28 uint32 = 2147483648
                    _ = v28
                }
            }
            _ = fi26
        }
        v3 = 4294967295
    }
    if true {
        v1 = int8(v3)
        for fi29 := 0; fi29 < 11; fi29++ {
            fmt.Println(fi29, v3, v1)
            var v30 int = (v3 | fi29)
            _ = fi29
            _ = v30
        }
    } else {
        switch "" {
        case "a":
            v1 -= (v1 << ((18446744073709551615) & 63))
        case "foo bar":
            if (0 == 42) {
                if (18446744073709551615 > 182) {
                    for fi31 := 0; fi31 < 15; fi31++ {
                        for fi32 := 0; fi32 < 5; fi32++ {
                            var v33 int8 = 100
                            v34 := GenS3{F0: ((("" + "a") + "hello") + "a"), F1: (v1 - v33)}
                            _ = fi32
                            _ = v33
                            _ = v34
                        }
                        switch fi31 {
                        case 983:
                            fmt.Println(v3, v3)
                        case 1:
                            switch (("" + (("a" + "\n\t") + "foo bar")) + ("hello" + ("a" + "hello"))) {
                            case "a":
                                var v35 int = genFn1(2147483647, 253, 50)
                                _ = v35
                            case "\n\t":
                                fmt.Println(fi31)
                            case "":
                                var v36 uint = 18446744073709551615
                                _ = v36
                            default:
                                v3 -= fi31
                            }
                        case -9223372036854775808:
                            switch "\n\t" {
                            case "hello":
                                for fi37 := 0; fi37 < 2; fi37++ {
                                    var v38 int16 = 676
                                    if (("foo bar" + "\n\t") < ("a" + "a")) {
                                        v3 += 0
                                        v39 := GenS1{F0: 330}
                                        fmt.Println(fi37, fi37, fi37)
                                        _ = v39
                                    }
                                    v40 := GenS1{F0: -1}
                                    v1 += v1
                                    _ = fi37
                                    _ = v38
                                    _ = v40
                                }
                            default:
                                var v41 string = "foo bar"
                                _ = v41
                            }
                        default:
                            if (157 < genFn0()) {
                                v42 := GenS2{F0: -1, F1: 42, F2: 2}
                                for fi43 := 0; fi43 < 17; fi43++ {
                                    var v44 bool = false
                                    v45 := GenS4{F0: -333}
                                    _ = fi43
                                    _ = v44
                                    _ = v45
                                }
                                var v46 int64 = 578
                                _ = v42
                                _ = v46
                            } else {
                                var v47 int32 = 0
                                if false {
                                    _ = len(v4)
                                    v47 -= v47
                                } else {
                                    var v48 float64 = 0.0
                                    for fi49 := 0; fi49 < 8; fi49++ {
                                        v50 := GenS5{F0: 127, F1: 1, F2: (v1 * (^(v1)))}
                                        var v51 uint8 = 170
                                        _ = fi49
                                        _ = v50
                                        _ = v51
                                    }
                                    fmt.Println(v48)
                                    _ = v48
                                }
                                fmt.Println(v3, v47, v3)
                                _ = v47
                            }
                        }
                        _ = fi31
                    }
                    v52 := GenS6{F0: 1, F1: 1, F2: 1}
                    if true {
                        var v53 uint8 = 2
                        if (!true) {
                            for fi54 := 0; fi54 < 11; fi54++ {
                                fmt.Println(fi54, v1, fi54)
                                _ = len(v5)
                                _ = fi54
                            }
                            switch v3 {
                            case 127:
                                v1 = 127
                            case -198:
                                for fi55 := 0; fi55 < 20; fi55++ {
                                    if (3.14159 < float32(float64(3.14159))) {
                                        for fi56 := 0; fi56 < 11; fi56++ {
                                            if (true && false) {
                                                var v57 float64 = ((0.0 - -2.25) / (1e10 * (0.1 * 3.14159)))
                                                var v58 uint8 = 198
                                                _ = v57
                                                _ = v58
                                            } else {
                                                if false {
                                                    var v59 float32 = 3.14159
                                                    switch ((("" + ("" + ("" + "a"))) + ("a" + "foo bar")) + "") {
                                                    case "":
                                                        switch (("\n\t" + ("" + ("\n\t" + "a"))) + ("hello" + "hello")) {
                                                        case "":
                                                            var v60 uint8 = (^(v53))
                                                            _ = v60
                                                        case "hello":
                                                            if (32767 < 0) {
                                                                if false {
                                                                    var v61 uint32 = 13
                                                                    if false {
                                                                        if (!(true || (true && true))) {
                                                                            if false {
                                                                                if len(v4) > 0 {
                                                                                    v4[0] = (v61 / (((v61 - (v61 & v61))) | 1))
                                                                                }
                                                                                switch "foo bar" {
                                                                                case "":
                                                                                    var v62 uint64 = 42
                                                                                    _ = v62
                                                                                case "foo bar":
                                                                                    v1 = v1
                                                                                case "hello":
                                                                                    var v63 int64 = 0
                                                                                    _ = v63
                                                                                }
                                                                                switch v3 {
                                                                                case -1:
                                                                                    if ((^(fi55)) < fi56) {
                                                                                        if ("\n\t" == "foo bar") {
                                                                                            if len(v2) >= 2 {
                                                                                                v2 = v2[1:len(v2)]
                                                                                            }
                                                                                        }
                                                                                        var v64 uint8 = (^(v53))
                                                                                        _ = v64
                                                                                    }
                                                                                default:
                                                                                    if len(v5) > 0 {
                                                                                        v5[0] = (v61 * 1)
                                                                                    }
                                                                                    var v65 uint16 = 967
                                                                                    if (true || true) {
                                                                                        if false {
                                                                                            v59 -= (v59 * v59)
                                                                                        } else {
                                                                                            switch (fi55 | fi56) {
                                                                                            case 278:
                                                                                                v59 += v59
                                                                                            case -2147483648:
                                                                                                if len(v2) >= 2 {
                                                                                                    v2 = v2[1:len(v2)]
                                                                                                }
                                                                                            default:
                                                                                                var v66 float32 = v59
                                                                                                _ = v66
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        var v67 uint64 = 42
                                                                                        if false {
                                                                                            v61 = 2147483648
                                                                                        }
                                                                                        _ = v67
                                                                                    }
                                                                                    _ = v65
                                                                                }
                                                                            } else {
                                                                                if false {
                                                                                    if true {
                                                                                        v1 = v1
                                                                                        v68 := GenS1{F0: 774}
                                                                                        _ = v68
                                                                                    } else {
                                                                                        if (!false) {
                                                                                            var v69 uint64 = 1
                                                                                            _ = v69
                                                                                        } else {
                                                                                            if (!true) {
                                                                                                v1 = (v1 & 42)
                                                                                                if true {
                                                                                                    v70 := GenS7{F0: 34}
                                                                                                    _ = v70
                                                                                                }
                                                                                                v71 := GenS3{F0: "hello", F1: v1}
                                                                                                _ = v71
                                                                                            }
                                                                                            v72 := GenS1{F0: 0}
                                                                                            _ = v72
                                                                                        }
                                                                                        var v73 int = (fi56 * (-(v3)))
                                                                                        switch 32767 {
                                                                                        case -2147483648:
                                                                                            v4 = append(v4, 65535)
                                                                                            var v74 uint = 255
                                                                                            _ = v74
                                                                                        case 42:
                                                                                            v75 := GenS5{F0: -1, F1: 2147483647, F2: v1}
                                                                                            _ = v75
                                                                                        case -1:
                                                                                            v1 = 67
                                                                                        default:
                                                                                            fmt.Println(v53)
                                                                                        }
                                                                                        if len(v4) >= 2 {
                                                                                            v4 = v4[1:len(v4)]
                                                                                        }
                                                                                        v3 = genFn1(-1, 242, 2)
                                                                                        _ = v73
                                                                                    }
                                                                                    var v76 string = "foo bar"
                                                                                    if len(v2) > 0 {
                                                                                        v2[0] = (v61 | (v61 / (((^(v61))) | 1)))
                                                                                    }
                                                                                    _ = v76
                                                                                } else {
                                                                                    switch 0 {
                                                                                    case 42:
                                                                                        var v77 int16 = 42
                                                                                        _ = v77
                                                                                    case 9223372036854775807:
                                                                                        if false {
                                                                                            fmt.Println(v3, v59)
                                                                                        }
                                                                                        if len(v4) > 0 {
                                                                                            v4[0] = v61
                                                                                        }
                                                                                        var v78 uint = 350
                                                                                        v78 = v78
                                                                                        if false {
                                                                                            var v79 int32 = 702
                                                                                            _ = v79
                                                                                        }
                                                                                        _ = v78
                                                                                    default:
                                                                                        v59 -= 0.1
                                                                                    }
                                                                                    fmt.Println(fi55)
                                                                                }
                                                                                if true {
                                                                                    v80 := GenS8{F0: (v3 << ((18446744073709551615) & 63)), F1: v1}
                                                                                    var v81 int8 = (v1 + (v1 * v1))
                                                                                    if (((0 > 42) || (!(1 >= 0))) && (false || true)) {
                                                                                        if false {
                                                                                            var v82 uint16 = 460
                                                                                            if (((v3 * (fi55 ^ 32768)) >= (fi56 << ((502) & 63))) && (!false)) {
                                                                                                switch "foo bar" {
                                                                                                case "hello":
                                                                                                    if (v59 < v59) {
                                                                                                        switch (fi55 + int(v53)) {
                                                                                                        case 32768:
                                                                                                            v3 = (fi56 >> ((0) & 63))
                                                                                                        case 173:
                                                                                                            fmt.Println(v53, v61)
                                                                                                        case 0:
                                                                                                            if false {
                                                                                                                v3 = (fi55 >> ((255) & 63))
                                                                                                                v82 = v82
                                                                                                            }
                                                                                                        }
                                                                                                        v61 = (v61 + v61)
                                                                                                        v83 := GenS4{F0: -9223372036854775808}
                                                                                                        _ = v83
                                                                                                    }
                                                                                                }
                                                                                                v3 = 128
                                                                                            }
                                                                                            if len(v4) >= 2 {
                                                                                                v4 = v4[1:len(v4)]
                                                                                            }
                                                                                            _ = v82
                                                                                        }
                                                                                        v1 += v1
                                                                                        if len(v4) >= 2 {
                                                                                            v4 = v4[1:len(v4)]
                                                                                        }
                                                                                    } else {
                                                                                        if true {
                                                                                            if true {
                                                                                                if (!(v61 == (v61 % (((v61 ^ (v61 % (((^(v61))) | 1)))) | 1)))) {
                                                                                                    fmt.Println(v81, v61)
                                                                                                    if (-422 > -1) {
                                                                                                        var v84 int64 = -726
                                                                                                        v85 := GenS9{F0: 1}
                                                                                                        v86 := GenS3{F0: ("foo bar" + ("" + "a")), F1: (v1 >> ((960) & 63))}
                                                                                                        _ = v84
                                                                                                        _ = v85
                                                                                                        _ = v86
                                                                                                    }
                                                                                                    var v87 uint8 = (^(v53))
                                                                                                    _ = v87
                                                                                                } else {
                                                                                                    switch fi56 {
                                                                                                    case 1:
                                                                                                        v59 = 3.14159
                                                                                                        v3 = (fi56 & 0)
                                                                                                        if len(v5) >= 2 {
                                                                                                            v5 = v5[1:len(v5)]
                                                                                                        }
                                                                                                    default:
                                                                                                        var v88 int8 = (v1 / ((15) | 1))
                                                                                                        _ = v88
                                                                                                    }
                                                                                                    switch ("foo bar" + "a") {
                                                                                                    case "hello":
                                                                                                        v3 -= fi55
                                                                                                    default:
                                                                                                        v3 = (fi55 | (v3 - (fi56 + (fi56 ^ v3))))
                                                                                                    }
                                                                                                    v89 := GenS4{F0: int64(fi56)}
                                                                                                    _ = v89
                                                                                                }
                                                                                                if false {
                                                                                                    v3 = v3
                                                                                                    var v90 int = (fi56 ^ genFn1(42, v53, 535))
                                                                                                    _ = len(v2)
                                                                                                    if len(v2) >= 2 {
                                                                                                        v2 = v2[1:len(v2)]
                                                                                                    }
                                                                                                    if len(v4) > 0 {
                                                                                                        v4[0] = 2147483648
                                                                                                    }
                                                                                                    _ = v90
                                                                                                }
                                                                                            } else {
                                                                                                if len(v4) >= 2 {
                                                                                                    v4 = v4[1:len(v4)]
                                                                                                }
                                                                                                switch int(v61) {
                                                                                                case 9223372036854775807:
                                                                                                    var v91 int16 = -1
                                                                                                    _ = v91
                                                                                                }
                                                                                                if len(v2) >= 2 {
                                                                                                    v2 = v2[1:len(v2)]
                                                                                                }
                                                                                                v1 -= (^(v81))
                                                                                                v53 = (v53 * v53)
                                                                                            }
                                                                                            if (v59 == 0.0) {
                                                                                                if (18446744073709551615 <= 11) {
                                                                                                    var v92 int64 = 774
                                                                                                    var v93 bool = true
                                                                                                    _ = v92
                                                                                                    _ = v93
                                                                                                }
                                                                                                fmt.Println(v61)
                                                                                            }
                                                                                            switch (-(v3)) {
                                                                                            case 32767:
                                                                                                switch -9223372036854775808 {
                                                                                                case 32768:
                                                                                                    if true {
                                                                                                        if false {
                                                                                                            var v94 uint8 = (v53 - v53)
                                                                                                            var v95 uint32 = (v61 / ((42) | 1))
                                                                                                            if len(v5) > 0 {
                                                                                                                v5[0] = 0
                                                                                                            }
                                                                                                            _ = v94
                                                                                                            _ = v95
                                                                                                        } else {
                                                                                                            v81 += (v81 & v81)
                                                                                                            if (true || true) {
                                                                                                                var v96 string = ((((("" + ("a" + ("" + ""))) + (("hello" + "\n\t") + "")) + ("hello" + "a")) + ("" + "a")) + "")
                                                                                                                _ = v96
                                                                                                            } else {
                                                                                                                _ = len(v4)
                                                                                                                v97 := GenS2{F0: genFn1(0, 179, 2), F1: -558, F2: 18446744073709551615}
                                                                                                                _ = v97
                                                                                                            }
                                                                                                        }
                                                                                                    } else {
                                                                                                        v61 += v61
                                                                                                        switch "foo bar" {
                                                                                                        case "foo bar":
                                                                                                            var v98 int64 = 1
                                                                                                            _ = v98
                                                                                                        }
                                                                                                        if false {
                                                                                                            var v99 uint8 = (v53 | v53)
                                                                                                            v81 = v1
                                                                                                            _ = v99
                                                                                                        } else {
                                                                                                            v81 -= (v81 ^ v1)
                                                                                                            if (2 <= 4294967295) {
                                                                                                                switch 535 {
                                                                                                                case 561:
                                                                                                                    if false {
                                                                                                                        var v100 uint16 = 823
                                                                                                                        _ = v100
                                                                                                                    }
                                                                                                                case 32768:
                                                                                                                    switch ((("a" + (("" + "\n\t") + ((("hello" + "hello") + "hello") + "a"))) + ("\n\t" + "a")) + "") {
                                                                                                                    case "a":
                                                                                                                        v81 = 1
                                                                                                                    case "":
                                                                                                                        if false {
                                                                                                                            if false {
                                                                                                                                var v101 uint32 = v61
                                                                                                                                if false {
                                                                                                                                    var v102 uint64 = 925
                                                                                                                                    v2 = append(v2, v61)
                                                                                                                                    _ = v102
                                                                                                                                }
                                                                                                                                _ = v101
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            var v103 uint8 = 2
                                                                                                                            if len(v2) > 0 {
                                                                                                                                v2[0] = v61
                                                                                                                            }
                                                                                                                            v61 -= 42
                                                                                                                            _ = v103
                                                                                                                        }
                                                                                                                    case "hello":
                                                                                                                        v61 = (v61 + (v61 * v61))
                                                                                                                    }
                                                                                                                case -314:
                                                                                                                    fmt.Println(v59, fi56)
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        if (65535 > 65535) {
                                                                                                            switch "hello" {
                                                                                                            case "hello":
                                                                                                                var v104 uint8 = v53
                                                                                                                _ = v104
                                                                                                            case "":
                                                                                                                if true {
                                                                                                                    if (((0.1 - (0.1 + float64(float32((1e10 * 0.0))))) - (1.5 - (0.1 * 1.5))) < 1e10) {
                                                                                                                        fmt.Println(v1)
                                                                                                                    } else {
                                                                                                                        var v105 string = "foo bar"
                                                                                                                        if (!(v105 <= v105)) {
                                                                                                                            v5 = append(v5, v61)
                                                                                                                        }
                                                                                                                        v106 := GenS1{F0: 489}
                                                                                                                        _ = v105
                                                                                                                        _ = v106
                                                                                                                    }
                                                                                                                    v1 = v81
                                                                                                                    if false {
                                                                                                                        fmt.Println(v53, fi56, v81)
                                                                                                                        if len(v4) >= 2 {
                                                                                                                            v4 = v4[1:len(v4)]
                                                                                                                        }
                                                                                                                        v53 -= (^(v53))
                                                                                                                    } else {
                                                                                                                        var v107 float32 = ((v59 / float32((0.0 + ((0.0 + ((3.14159 / 0.1) / (((((0.1 * 3.14159) - 0.1) + 1.5) + -2.25) * 0.1))) + ((1.5 - 3.14159) / (1.5 * -2.25)))))) - v59)
                                                                                                                        _ = v107
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    v61 = (v61 + (^(v61)))
                                                                                                                    v3 = (fi56 | v3)
                                                                                                                    v3 = (v3 & 127)
                                                                                                                    fmt.Println(v53)
                                                                                                                    if len(v2) >= 2 {
                                                                                                                        v2 = v2[1:len(v2)]
                                                                                                                    }
                                                                                                                    var v108 int64 = -1
                                                                                                                    v1 += (v81 + 0)
                                                                                                                    _ = v108
                                                                                                                }
                                                                                                            case "foo bar":
                                                                                                                if true {
                                                                                                                    fmt.Println(fi56, v81)
                                                                                                                    var v109 uint16 = 653
                                                                                                                    _ = v109
                                                                                                                } else {
                                                                                                                    v3 = fi55
                                                                                                                    if (-964 >= 2147483647) {
                                                                                                                        switch "a" {
                                                                                                                        case "\n\t":
                                                                                                                            var v110 uint8 = (^(v53))
                                                                                                                            _ = v110
                                                                                                                        case "a":
                                                                                                                            v59 -= (((float32(3.14159) + v59) * ((float32(0.1) * 0.1) * v59)) * (v59 * (((v59 / (1e10 - v59)) * (v59 + float32(((-2.25 + 0.0) * (0.0 * (1.5 * (0.1 + ((0.1 - (0.1 + 0.1)) / (-2.25 / 3.14159))))))))) / v59)))
                                                                                                                        default:
                                                                                                                            var v111 float32 = ((1e10 + v59) - (1.5 - v59))
                                                                                                                            _ = v111
                                                                                                                        }
                                                                                                                        var v112 uint64 = 1
                                                                                                                        _ = v112
                                                                                                                    } else {
                                                                                                                        var v113 uint64 = 42
                                                                                                                        v114 := GenS8{F0: 128, F1: (v81 << ((39) & 63))}
                                                                                                                        var v115 int32 = 2147483647
                                                                                                                        _ = v113
                                                                                                                        _ = v114
                                                                                                                        _ = v115
                                                                                                                    }
                                                                                                                    if (1 >= 530) {
                                                                                                                        var v116 int16 = -759
                                                                                                                        var v117 float32 = v59
                                                                                                                        v118 := GenS10{F0: 910, F1: 0}
                                                                                                                        _ = len(v4)
                                                                                                                        fmt.Println(v117)
                                                                                                                        v2 = append(v2, v61)
                                                                                                                        v117 += 0.1
                                                                                                                        switch -988 {
                                                                                                                        case 2147483647:
                                                                                                                            var v119 bool = true
                                                                                                                            _ = v119
                                                                                                                        case 127:
                                                                                                                            if false {
                                                                                                                                switch "\n\t" {
                                                                                                                                case "hello":
                                                                                                                                    if true {
                                                                                                                                        switch "foo bar" {
                                                                                                                                        case "\n\t":
                                                                                                                                            var v120 int64 = 42
                                                                                                                                            switch "a" {
                                                                                                                                            case "":
                                                                                                                                                switch "hello" {
                                                                                                                                                case "a":
                                                                                                                                                    v121 := GenS5{F0: (v1 / ((v81) | 1)), F1: 42, F2: (v1 - 67)}
                                                                                                                                                    _ = v121
                                                                                                                                                }
                                                                                                                                            case "hello":
                                                                                                                                                var v122 uint = 0
                                                                                                                                                var v123 uint = uint(v53)
                                                                                                                                                v124 := GenS5{F0: v1, F1: -778, F2: 42}
                                                                                                                                                if false {
                                                                                                                                                    fmt.Println(v3, v59, v116)
                                                                                                                                                    var v125 int8 = v81
                                                                                                                                                    if (true || ("\n\t" >= ("a" + "\n\t"))) {
                                                                                                                                                        var v126 string = ("foo bar" + "foo bar")
                                                                                                                                                        _ = v126
                                                                                                                                                    }
                                                                                                                                                    var v127 float32 = (v59 + 0.1)
                                                                                                                                                    v59 += -2.25
                                                                                                                                                    _ = v125
                                                                                                                                                    _ = v127
                                                                                                                                                }
                                                                                                                                                var v128 bool = (((v122 | v123) <= 510) && true)
                                                                                                                                                if len(v5) >= 2 {
                                                                                                                                                    v5 = v5[1:len(v5)]
                                                                                                                                                }
                                                                                                                                                fmt.Println(v116, v59)
                                                                                                                                                fmt.Println(v61)
                                                                                                                                                _ = v122
                                                                                                                                                _ = v123
                                                                                                                                                _ = v124
                                                                                                                                                _ = v128
                                                                                                                                            default:
                                                                                                                                                if ("foo bar" != "\n\t") {
                                                                                                                                                    _ = len(v5)
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                            var v129 uint8 = (v53 * (^(v53)))
                                                                                                                                            _ = v120
                                                                                                                                            _ = v129
                                                                                                                                        case "":
                                                                                                                                            v130 := GenS8{F0: (-(v3)), F1: 115}
                                                                                                                                            _ = v130
                                                                                                                                        case "hello":
                                                                                                                                            switch (fi56 + 365) {
                                                                                                                                            case 32767:
                                                                                                                                                if ((0.0 * 1.5) >= 0.0) {
                                                                                                                                                    if true {
                                                                                                                                                        if (genFn1(0, v53, 1) >= (v3 * fi56)) {
                                                                                                                                                            var v131 string = "a"
                                                                                                                                                            v81 = (^(v1))
                                                                                                                                                            _ = v131
                                                                                                                                                        }
                                                                                                                                                        if true {
                                                                                                                                                            if false {
                                                                                                                                                                v59 = v117
                                                                                                                                                            } else {
                                                                                                                                                                v61 -= 42
                                                                                                                                                                var v132 uint64 = 873
                                                                                                                                                                v116 -= v116
                                                                                                                                                                _ = v132
                                                                                                                                                            }
                                                                                                                                                            if ((-766 >= 18) && false) {
                                                                                                                                                                if len(v4) >= 2 {
                                                                                                                                                                    v4 = v4[1:len(v4)]
                                                                                                                                                                }
                                                                                                                                                                v2 = append(v2, 2147483647)
                                                                                                                                                            } else {
                                                                                                                                                                fmt.Println(fi56)
                                                                                                                                                                v116 = (v116 - v116)
                                                                                                                                                            }
                                                                                                                                                        }
                                                                                                                                                        fmt.Println(v117)
                                                                                                                                                    } else {
                                                                                                                                                        v59 += 0.0
                                                                                                                                                        var v133 int32 = -2147483648
                                                                                                                                                        _ = v133
                                                                                                                                                    }
                                                                                                                                                } else {
                                                                                                                                                    v134 := GenS8{F0: genFn1(2147483647, (^(v53)), 2), F1: (v81 * v81)}
                                                                                                                                                    v135 := GenS9{F0: 294}
                                                                                                                                                    if len(v5) > 0 {
                                                                                                                                                        v5[0] = (v61 % ((v61) | 1))
                                                                                                                                                    }
                                                                                                                                                    _ = v134
                                                                                                                                                    _ = v135
                                                                                                                                                }
                                                                                                                                            default:
                                                                                                                                                var v136 int16 = 806
                                                                                                                                                _ = v136
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                        var v137 uint8 = 69
                                                                                                                                        var v138 int16 = v116
                                                                                                                                        if (false || true) {
                                                                                                                                            var v139 float32 = (3.14159 + v117)
                                                                                                                                            v138 -= v138
                                                                                                                                            v61 = (v61 | 2147483648)
                                                                                                                                            _ = v139
                                                                                                                                        }
                                                                                                                                        v140 := GenS2{F0: fi56, F1: -9223372036854775808, F2: 202}
                                                                                                                                        if (1e10 == v117) {
                                                                                                                                            var v141 int16 = v116
                                                                                                                                            v137 = v137
                                                                                                                                            _ = v141
                                                                                                                                        }
                                                                                                                                        var v142 uint16 = 0
                                                                                                                                        _ = v137
                                                                                                                                        _ = v138
                                                                                                                                        _ = v140
                                                                                                                                        _ = v142
                                                                                                                                    }
                                                                                                                                default:
                                                                                                                                    fmt.Println(v117, v81)
                                                                                                                                }
                                                                                                                                if (true || true) {
                                                                                                                                    v4 = append(v4, v61)
                                                                                                                                    v53 = v53
                                                                                                                                    fmt.Println(fi55)
                                                                                                                                } else {
                                                                                                                                    if (!(0 == 0)) {
                                                                                                                                        var v143 int = genFn1(2147483647, v53, uint64(v53))
                                                                                                                                        if true {
                                                                                                                                            if false {
                                                                                                                                                v3 -= (fi56 / ((int(v53)) | 1))
                                                                                                                                            } else {
                                                                                                                                                v1 -= 0
                                                                                                                                                _ = len(v4)
                                                                                                                                                if true {
                                                                                                                                                    v1 -= (^(v1))
                                                                                                                                                    var v144 int8 = 1
                                                                                                                                                    v145 := GenS7{F0: 65535}
                                                                                                                                                    _ = v144
                                                                                                                                                    _ = v145
                                                                                                                                                } else {
                                                                                                                                                    switch "hello" {
                                                                                                                                                    case "":
                                                                                                                                                        v3 += (v143 ^ (^(v3)))
                                                                                                                                                    case "foo bar":
                                                                                                                                                        v81 -= (v1 << ((987) & 63))
                                                                                                                                                    default:
                                                                                                                                                        if len(v5) >= 2 {
                                                                                                                                                            v5 = v5[1:len(v5)]
                                                                                                                                                        }
                                                                                                                                                    }
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                        } else {
                                                                                                                                            v146 := GenS11{F0: -1}
                                                                                                                                            _ = v146
                                                                                                                                        }
                                                                                                                                        v1 += (v81 - 93)
                                                                                                                                        _ = v143
                                                                                                                                    } else {
                                                                                                                                        var v147 int32 = 0
                                                                                                                                        _ = v147
                                                                                                                                    }
                                                                                                                                }
                                                                                                                            }
                                                                                                                        case 476:
                                                                                                                            if ("" > "foo bar") {
                                                                                                                                switch "a" {
                                                                                                                                case "foo bar":
                                                                                                                                    fmt.Println(v53)
                                                                                                                                case "\n\t":
                                                                                                                                    fmt.Println(v117, v117, v53)
                                                                                                                                case "hello":
                                                                                                                                    if len(v2) >= 2 {
                                                                                                                                        v2 = v2[1:len(v2)]
                                                                                                                                    }
                                                                                                                                }
                                                                                                                                v53 += v53
                                                                                                                                v148 := GenS12{F0: (v61 >> ((2) & 63))}
                                                                                                                                _ = v148
                                                                                                                            } else {
                                                                                                                                v1 = (v81 & 1)
                                                                                                                            }
                                                                                                                            fmt.Println(fi56)
                                                                                                                            var v149 string = ("" + ("" + "foo bar"))
                                                                                                                            _ = v149
                                                                                                                        default:
                                                                                                                            v150 := GenS11{F0: 1}
                                                                                                                            _ = v150
                                                                                                                        }
                                                                                                                        _ = v116
                                                                                                                        _ = v117
                                                                                                                        _ = v118
                                                                                                                    } else {
                                                                                                                        v151 := GenS13{F0: 42, F1: (-(v81))}
                                                                                                                        fmt.Println(fi55, v61)
                                                                                                                        _ = v151
                                                                                                                    }
                                                                                                                }
                                                                                                            default:
                                                                                                                if len(v4) >= 2 {
                                                                                                                    v4 = v4[1:len(v4)]
                                                                                                                }
                                                                                                            }
                                                                                                            v2 = append(v2, v61)
                                                                                                            if false {
                                                                                                                if (-1 == 120) {
                                                                                                                    v152 := GenS5{F0: (v1 - v1), F1: 2147483647, F2: (v1 + v1)}
                                                                                                                    if false {
                                                                                                                        if false {
                                                                                                                            fmt.Println(fi55)
                                                                                                                            if ((false || false) || ("hello" > "foo bar")) {
                                                                                                                                v61 += 1
                                                                                                                                if true {
                                                                                                                                    var v153 string = ("\n\t" + "a")
                                                                                                                                    switch (("\n\t" + v153) + "a") {
                                                                                                                                    case "hello":
                                                                                                                                        var v154 bool = ((^(v53)) == v53)
                                                                                                                                        _ = v154
                                                                                                                                    case "\n\t":
                                                                                                                                        var v155 uint8 = 91
                                                                                                                                        _ = v155
                                                                                                                                    }
                                                                                                                                    _ = v153
                                                                                                                                }
                                                                                                                                var v156 int64 = 1
                                                                                                                                _ = v156
                                                                                                                            } else {
                                                                                                                                v3 += genFn1(1, (v53 << ((42) & 63)), 18446744073709551615)
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            v157 := GenS4{F0: -9223372036854775808}
                                                                                                                            var v158 string = ""
                                                                                                                            fmt.Println(v59, v3)
                                                                                                                            v159 := GenS4{F0: 1}
                                                                                                                            var v160 int16 = 109
                                                                                                                            _ = v157
                                                                                                                            _ = v158
                                                                                                                            _ = v159
                                                                                                                            _ = v160
                                                                                                                        }
                                                                                                                        if (0 == 32767) {
                                                                                                                            if true {
                                                                                                                                fmt.Println(v59)
                                                                                                                                var v161 int = 127
                                                                                                                                var v162 int32 = 0
                                                                                                                                _ = v161
                                                                                                                                _ = v162
                                                                                                                            } else {
                                                                                                                                if false {
                                                                                                                                    if len(v4) >= 2 {
                                                                                                                                        v4 = v4[1:len(v4)]
                                                                                                                                    }
                                                                                                                                    v59 -= v59
                                                                                                                                    v2 = append(v2, v61)
                                                                                                                                }
                                                                                                                                var v163 uint64 = 42
                                                                                                                                _ = v163
                                                                                                                            }
                                                                                                                            var v164 uint8 = v53
                                                                                                                            if (!true) {
                                                                                                                                if (154 > 0) {
                                                                                                                                    switch "foo bar" {
                                                                                                                                    case "hello":
                                                                                                                                        var v165 uint16 = 2
                                                                                                                                        _ = v165
                                                                                                                                    default:
                                                                                                                                        v164 = v53
                                                                                                                                    }
                                                                                                                                } else {
                                                                                                                                    if (!(v164 > v164)) {
                                                                                                                                        if (false && (!(false || true))) {
                                                                                                                                            v166 := GenS1{F0: 0}
                                                                                                                                            _ = v166
                                                                                                                                        } else {
                                                                                                                                            v59 = float32((((1.5 - 3.14159) - (0.1 + (0.0 * 1e10))) / 1.0))
                                                                                                                                            if (653 >= v61) {
                                                                                                                                                switch genFn1(14, genFn0(), 0) {
                                                                                                                                                case 32767:
                                                                                                                                                    var v167 string = "hello"
                                                                                                                                                    _ = v167
                                                                                                                                                case 2147483647:
                                                                                                                                                    v59 -= v59
                                                                                                                                                default:
                                                                                                                                                    if false {
                                                                                                                                                        v168 := GenS2{F0: v3, F1: 0, F2: 395}
                                                                                                                                                        _ = v168
                                                                                                                                                    }
                                                                                                                                                }
                                                                                                                                                var v169 uint64 = 2
                                                                                                                                                _ = v169
                                                                                                                                            }
                                                                                                                                            v170 := GenS14{F0: 42, F1: 18446744073709551615, F2: (v61 & v61)}
                                                                                                                                            var v171 uint8 = v164
                                                                                                                                            _ = v170
                                                                                                                                            _ = v171
                                                                                                                                        }
                                                                                                                                        switch fi55 {
                                                                                                                                        case 32768:
                                                                                                                                            switch ("hello" + (("" + "foo bar") + "\n\t")) {
                                                                                                                                            case "foo bar":
                                                                                                                                                fmt.Println(v53, v61)
                                                                                                                                            case "hello":
                                                                                                                                                switch 2147483647 {
                                                                                                                                                case 2147483647:
                                                                                                                                                    var v172 string = ""
                                                                                                                                                    _ = v172
                                                                                                                                                case -557:
                                                                                                                                                    v1 -= v81
                                                                                                                                                case -2147483648:
                                                                                                                                                    v173 := GenS15{F0: -357}
                                                                                                                                                    _ = v173
                                                                                                                                                default:
                                                                                                                                                    var v174 bool = false
                                                                                                                                                    _ = v174
                                                                                                                                                }
                                                                                                                                            case "\n\t":
                                                                                                                                                var v175 string = "a"
                                                                                                                                                _ = v175
                                                                                                                                            default:
                                                                                                                                                switch fi55 {
                                                                                                                                                case -1:
                                                                                                                                                    fmt.Println(v53)
                                                                                                                                                case -298:
                                                                                                                                                    if true {
                                                                                                                                                        fmt.Println(v164)
                                                                                                                                                    } else {
                                                                                                                                                        if false {
                                                                                                                                                            if len(v2) >= 2 {
                                                                                                                                                                v2 = v2[1:len(v2)]
                                                                                                                                                            }
                                                                                                                                                            if (1 > 0) {
                                                                                                                                                                if true {
                                                                                                                                                                    var v176 uint8 = (v164 << ((356) & 63))
                                                                                                                                                                    if (true || (2147483647 <= -367)) {
                                                                                                                                                                        switch "\n\t" {
                                                                                                                                                                        case "foo bar":
                                                                                                                                                                            var v177 int = 42
                                                                                                                                                                            _ = v177
                                                                                                                                                                        }
                                                                                                                                                                    }
                                                                                                                                                                    _ = v176
                                                                                                                                                                } else {
                                                                                                                                                                    switch (fi56 / ((fi56) | 1)) {
                                                                                                                                                                    case 2147483647:
                                                                                                                                                                        var v178 int16 = 42
                                                                                                                                                                        _ = v178
                                                                                                                                                                    default:
                                                                                                                                                                        if len(v5) >= 2 {
                                                                                                                                                                            v5 = v5[1:len(v5)]
                                                                                                                                                                        }
                                                                                                                                                                    }
                                                                                                                                                                    switch genFn1(-1, (^(v53)), 357) {
                                                                                                                                                                    case 2147483647:
                                                                                                                                                                        _ = len(v4)
                                                                                                                                                                    default:
                                                                                                                                                                        switch ((("\n\t" + (("foo bar" + "") + "\n\t")) + "\n\t") + "foo bar") {
                                                                                                                                                                        case "a":
                                                                                                                                                                            v1 = v81
                                                                                                                                                                        }
                                                                                                                                                                    }
                                                                                                                                                                }
                                                                                                                                                            }
                                                                                                                                                        }
                                                                                                                                                        var v179 int64 = -9223372036854775808
                                                                                                                                                        v53 += (^(v53))
                                                                                                                                                        _ = v179
                                                                                                                                                    }
                                                                                                                                                case 614:
                                                                                                                                                    v180 := GenS13{F0: -2147483648, F1: (v1 & v81)}
                                                                                                                                                    _ = v180
                                                                                                                                                default:
                                                                                                                                                    if false {
                                                                                                                                                        var v181 uint16 = 1
                                                                                                                                                        _ = v181
                                                                                                                                                    } else {
                                                                                                                                                        v81 = v81
                                                                                                                                                    }
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                        case 128:
                                                                                                                                            v53 -= (v53 % (((v164 & (v164 | uint8(v61)))) | 1))
                                                                                                                                        default:
                                                                                                                                            v182 := GenS8{F0: genFn1(-1, (^(v53)), 552), F1: (-(v81))}
                                                                                                                                            _ = v182
                                                                                                                                        }
                                                                                                                                        v3 = 127
                                                                                                                                    }
                                                                                                                                    switch "a" {
                                                                                                                                    case "hello":
                                                                                                                                        v3 += (fi56 % ((v3) | 1))
                                                                                                                                    default:
                                                                                                                                        v2 = append(v2, v61)
                                                                                                                                    }
                                                                                                                                    var v183 int64 = 0
                                                                                                                                    var v184 float64 = 0.1
                                                                                                                                    var v185 uint = 255
                                                                                                                                    switch (-(v3)) {
                                                                                                                                    case 524:
                                                                                                                                        if (2 <= 302) {
                                                                                                                                            var v186 uint32 = (^(v61))
                                                                                                                                            _ = v186
                                                                                                                                        }
                                                                                                                                    case 1:
                                                                                                                                        _ = len(v2)
                                                                                                                                    case 752:
                                                                                                                                        v81 = (v81 * (v1 * v1))
                                                                                                                                        if (!((1 != 0) || true)) {
                                                                                                                                            v3 = int(v164)
                                                                                                                                            var v187 string = ("a" + "hello")
                                                                                                                                            v53 = (v164 % ((genFn0()) | 1))
                                                                                                                                            _ = v187
                                                                                                                                        }
                                                                                                                                        var v188 int16 = 1
                                                                                                                                        if false {
                                                                                                                                            var v189 int64 = v183
                                                                                                                                            _ = v189
                                                                                                                                        }
                                                                                                                                        if true {
                                                                                                                                            var v190 float32 = ((((v59 / float32(((v184 - float64(v59)) * 1.5))) - v59) * v59) + (v59 - (0.1 - 1.5)))
                                                                                                                                            switch "" {
                                                                                                                                            case "a":
                                                                                                                                                switch (v3 ^ v3) {
                                                                                                                                                case 442:
                                                                                                                                                    v183 = v183
                                                                                                                                                case 32767:
                                                                                                                                                    v81 = v81
                                                                                                                                                case -2147483648:
                                                                                                                                                    if false {
                                                                                                                                                        if true {
                                                                                                                                                            v2 = append(v2, v61)
                                                                                                                                                            if (true || (2147483647 == 42)) {
                                                                                                                                                                v53 = (v53 / (((^(v53))) | 1))
                                                                                                                                                                var v191 uint16 = 698
                                                                                                                                                                if false {
                                                                                                                                                                    v4 = append(v4, uint32(v164))
                                                                                                                                                                }
                                                                                                                                                                _ = v191
                                                                                                                                                            }
                                                                                                                                                            v192 := GenS12{F0: v61}
                                                                                                                                                            _ = v192
                                                                                                                                                        } else {
                                                                                                                                                            var v193 uint16 = 1
                                                                                                                                                            var v194 int = v3
                                                                                                                                                            var v195 int32 = 2147483647
                                                                                                                                                            _ = v193
                                                                                                                                                            _ = v194
                                                                                                                                                            _ = v195
                                                                                                                                                        }
                                                                                                                                                        var v196 float32 = v59
                                                                                                                                                        var v197 uint16 = 0
                                                                                                                                                        _ = v196
                                                                                                                                                        _ = v197
                                                                                                                                                    } else {
                                                                                                                                                        _ = len(v2)
                                                                                                                                                    }
                                                                                                                                                }
                                                                                                                                            default:
                                                                                                                                                var v198 string = "\n\t"
                                                                                                                                                _ = v198
                                                                                                                                            }
                                                                                                                                            _ = v190
                                                                                                                                        }
                                                                                                                                        _ = v188
                                                                                                                                    }
                                                                                                                                    _ = v183
                                                                                                                                    _ = v184
                                                                                                                                    _ = v185
                                                                                                                                }
                                                                                                                                var v199 int64 = int64(fi55)
                                                                                                                                if true {
                                                                                                                                    v3 = (fi55 & (fi56 % ((fi55) | 1)))
                                                                                                                                } else {
                                                                                                                                    v59 -= float32((0.0 / float64(0.1)))
                                                                                                                                    if true {
                                                                                                                                        fmt.Println(v164)
                                                                                                                                        fmt.Println(v53, v59)
                                                                                                                                    } else {
                                                                                                                                        v3 = fi55
                                                                                                                                        v53 += (^(v53))
                                                                                                                                        var v200 uint16 = 535
                                                                                                                                        var v201 int16 = 855
                                                                                                                                        _ = v200
                                                                                                                                        _ = v201
                                                                                                                                    }
                                                                                                                                }
                                                                                                                                _ = v199
                                                                                                                            }
                                                                                                                            _ = v164
                                                                                                                        }
                                                                                                                        if len(v2) >= 2 {
                                                                                                                            v2 = v2[1:len(v2)]
                                                                                                                        }
                                                                                                                    }
                                                                                                                    _ = v152
                                                                                                                } else {
                                                                                                                    fmt.Println(v61, fi56, v81)
                                                                                                                }
                                                                                                            }
                                                                                                            switch v3 {
                                                                                                            case -46:
                                                                                                                var v202 int = genFn1(0, 62, 18446744073709551615)
                                                                                                                _ = v202
                                                                                                            case 127:
                                                                                                                if (false || (!(v61 != (v61 >> ((2) & 63))))) {
                                                                                                                    fmt.Println(fi55, v1)
                                                                                                                } else {
                                                                                                                    fmt.Println(v81)
                                                                                                                    fmt.Println(v53, v53, v61)
                                                                                                                    v203 := GenS12{F0: (^(v61))}
                                                                                                                    var v204 int = (fi55 & fi55)
                                                                                                                    fmt.Println(v81, v81)
                                                                                                                    _ = len(v2)
                                                                                                                    _ = v203
                                                                                                                    _ = v204
                                                                                                                }
                                                                                                            case -2147483648:
                                                                                                                if (42 < -557) {
                                                                                                                    switch ("foo bar" + "\n\t") {
                                                                                                                    case "hello":
                                                                                                                        v205 := GenS7{F0: 0}
                                                                                                                        _ = v205
                                                                                                                    case "\n\t":
                                                                                                                        fmt.Println(fi55, v61, v59)
                                                                                                                    default:
                                                                                                                        var v206 int16 = -1
                                                                                                                        fmt.Println(v3, v61, fi55)
                                                                                                                        _ = v206
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    switch "\n\t" {
                                                                                                                    case "\n\t":
                                                                                                                        if (!true) {
                                                                                                                            fmt.Println(v53, v81, v61)
                                                                                                                        } else {
                                                                                                                            if true {
                                                                                                                                var v207 uint32 = (v61 ^ 65535)
                                                                                                                                var v208 float64 = (1e10 / -2.25)
                                                                                                                                var v209 int16 = 0
                                                                                                                                _ = v207
                                                                                                                                _ = v208
                                                                                                                                _ = v209
                                                                                                                            }
                                                                                                                            var v210 int32 = 0
                                                                                                                            _ = v210
                                                                                                                        }
                                                                                                                    case "":
                                                                                                                        if true {
                                                                                                                            if len(v2) >= 2 {
                                                                                                                                v2 = v2[1:len(v2)]
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            if (!(!(("\n\t" + "a") == "a"))) {
                                                                                                                                fmt.Println(fi56, v3)
                                                                                                                            }
                                                                                                                            switch (((("\n\t" + (("hello" + "hello") + "\n\t")) + "hello") + ("" + ("a" + "foo bar"))) + "") {
                                                                                                                            case "foo bar":
                                                                                                                                fmt.Println(v81)
                                                                                                                            case "a":
                                                                                                                                var v211 uint64 = 2
                                                                                                                                _ = v211
                                                                                                                            default:
                                                                                                                                if len(v4) > 0 {
                                                                                                                                    v4[0] = (v61 - (v61 | (^(v61))))
                                                                                                                                }
                                                                                                                            }
                                                                                                                            if (v61 <= (v61 << ((2) & 63))) {
                                                                                                                                v212 := GenS6{F0: (v61 & v61), F1: 2, F2: 2147483647}
                                                                                                                                _ = v212
                                                                                                                            }
                                                                                                                        }
                                                                                                                    }
                                                                                                                    var v213 int32 = -1
                                                                                                                    _ = v213
                                                                                                                }
                                                                                                            default:
                                                                                                                var v214 uint8 = genFn0()
                                                                                                                _ = v214
                                                                                                            }
                                                                                                            switch (fi56 | 0) {
                                                                                                            case -9223372036854775808:
                                                                                                                fmt.Println(v81)
                                                                                                            case 0:
                                                                                                                v81 = (v81 | (v81 ^ v1))
                                                                                                            default:
                                                                                                                _ = len(v2)
                                                                                                            }
                                                                                                        } else {
                                                                                                            var v215 float64 = -2.25
                                                                                                            fmt.Println(v53)
                                                                                                            v216 := GenS4{F0: 1}
                                                                                                            _ = v215
                                                                                                            _ = v216
                                                                                                        }
                                                                                                    }
                                                                                                case 2147483647:
                                                                                                    var v217 int32 = 0
                                                                                                    _ = v217
                                                                                                case 0:
                                                                                                    var v218 float64 = 3.14159
                                                                                                    v219 := GenS16{F0: -1, F1: 1, F2: 52}
                                                                                                    v220 := GenS11{F0: -840}
                                                                                                    _ = v218
                                                                                                    _ = v219
                                                                                                    _ = v220
                                                                                                default:
                                                                                                    v59 = ((v59 / ((float32(1.5) * ((-2.25 + (float32(0.1) * v59)) + v59)) * ((v59 - v59) / v59))) * 0.1)
                                                                                                }
                                                                                            case 128:
                                                                                                v61 = 65535
                                                                                            case 0:
                                                                                                _ = len(v4)
                                                                                            default:
                                                                                                switch 127 {
                                                                                                case 128:
                                                                                                    v1 += (v1 / (((v81 % (((v1 + v1)) | 1))) | 1))
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        v53 = (v53 ^ v53)
                                                                                        fmt.Println(v53)
                                                                                    }
                                                                                    _ = v80
                                                                                    _ = v81
                                                                                }
                                                                            }
                                                                            v61 = (^(v61))
                                                                            v61 = (v61 | (^(v61)))
                                                                        }
                                                                        if true {
                                                                            if (3.14159 == 1.5) {
                                                                                v4 = append(v4, v61)
                                                                                switch "" {
                                                                                case "":
                                                                                    switch "foo bar" {
                                                                                    case "hello":
                                                                                        var v221 uint32 = 2147483647
                                                                                        _ = v221
                                                                                    case "a":
                                                                                        v53 = (v53 * 255)
                                                                                    case "":
                                                                                        v53 = 176
                                                                                    }
                                                                                }
                                                                                v3 = (^(fi55))
                                                                                v5 = append(v5, uint32(v53))
                                                                                v2 = append(v2, (v61 + v61))
                                                                                v1 = -31
                                                                            } else {
                                                                                v222 := GenS17{F0: 255}
                                                                                _ = v222
                                                                            }
                                                                            _ = len(v5)
                                                                            var v223 int16 = 42
                                                                            _ = v223
                                                                        }
                                                                        if ((v53 << ((uint(v53)) & 63)) == genFn0()) {
                                                                            switch "foo bar" {
                                                                            case "foo bar":
                                                                                v224 := GenS15{F0: 32767}
                                                                                _ = v224
                                                                            default:
                                                                                if len(v5) >= 2 {
                                                                                    v5 = v5[1:len(v5)]
                                                                                }
                                                                            }
                                                                            fmt.Println(v61, v3)
                                                                        }
                                                                    }
                                                                    _ = v61
                                                                }
                                                                fmt.Println(v59)
                                                            } else {
                                                                var v225 int32 = 42
                                                                fmt.Println(fi56)
                                                                _ = v225
                                                            }
                                                        }
                                                    case "hello":
                                                        switch "" {
                                                        case "foo bar":
                                                            v59 = (0.1 / 1.5)
                                                        }
                                                    default:
                                                        v226 := GenS16{F0: 42, F1: 0, F2: 0}
                                                        _ = v226
                                                    }
                                                    v1 = (v1 << ((1) & 63))
                                                    _ = v59
                                                }
                                                if true {
                                                    v53 -= v53
                                                    var v227 int = (fi56 % ((127) | 1))
                                                    _ = v227
                                                }
                                            }
                                            _ = fi56
                                        }
                                        var v228 int64 = 0
                                        _ = v228
                                    } else {
                                        v229 := GenS11{F0: 42}
                                        _ = v229
                                    }
                                    _ = fi55
                                }
                            case -1:
                                if (-604 == -705) {
                                    var v230 uint = 65535
                                    var v231 float32 = float32((((1.5 + -2.25) - (0.0 + (-2.25 - 0.1))) - 1e10))
                                    var v232 int16 = 1000
                                    _ = v230
                                    _ = v231
                                    _ = v232
                                }
                            }
                            fmt.Println(v1, v3)
                        } else {
                            v3 = -104
                            v233 := GenS3{F0: "hello", F1: (v1 << ((1) & 63))}
                            switch (^(v3)) {
                            case -9223372036854775808:
                                switch "\n\t" {
                                case "\n\t":
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                case "hello":
                                    var v234 int32 = 0
                                    _ = v234
                                case "a":
                                    fmt.Println(v3, v53)
                                }
                            case 128:
                                for fi235 := 0; fi235 < 9; fi235++ {
                                    v3 += 173
                                    var v236 float32 = (1e10 - 1e10)
                                    v3 += (v3 >> ((0) & 63))
                                    _ = fi235
                                    _ = v236
                                }
                            default:
                                v4 = append(v4, 718)
                            }
                            _ = v233
                        }
                        _ = v53
                    } else {
                        v1 += -128
                        fmt.Println(v3, v3)
                    }
                    _ = v52
                } else {
                    if (true && (269 < 65535)) {
                        v3 = (v3 & (v3 - v3))
                        for fi237 := 0; fi237 < 13; fi237++ {
                            fmt.Println(v1, v1, fi237)
                            var v238 int32 = 0
                            _ = fi237
                            _ = v238
                        }
                        if false {
                            var v239 uint32 = 746
                            fmt.Println(v3)
                            for fi240 := 0; fi240 < 1; fi240++ {
                                switch "a" {
                                case "foo bar":
                                    if true {
                                        if len(v2) > 0 {
                                            v2[0] = v239
                                        }
                                    }
                                case "hello":
                                    var v241 float32 = 0.0
                                    _ = v241
                                default:
                                    v242 := GenS18{F0: 983, F1: true, F2: 642}
                                    _ = v242
                                }
                                if (false || false) {
                                    v239 = (v239 / (((v239 ^ 65535)) | 1))
                                    var v243 string = "\n\t"
                                    _ = v243
                                } else {
                                    v3 -= (v3 & (fi240 & v3))
                                    var v244 int64 = -9223372036854775808
                                    _ = v244
                                }
                                _ = fi240
                            }
                            _ = v239
                        } else {
                            switch "foo bar" {
                            case "":
                                if false {
                                    v3 += (^(v3))
                                    var v245 float32 = 1.5
                                    _ = v245
                                } else {
                                    if true {
                                        if (!true) {
                                            v3 = (v3 ^ v3)
                                            if false {
                                                v3 += v3
                                                v4 = append(v4, 42)
                                                for fi246 := 0; fi246 < 4; fi246++ {
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    _ = fi246
                                                }
                                            } else {
                                                if (!false) {
                                                    v3 = (^(v3))
                                                    v1 -= 42
                                                    if ((0.0 / (1e10 / 0.1)) < 0.0) {
                                                        var v247 bool = true
                                                        var v248 int64 = -1
                                                        _ = v247
                                                        _ = v248
                                                    }
                                                }
                                                var v249 uint32 = 1
                                                _ = v249
                                            }
                                        }
                                        var v250 int64 = 1
                                        v251 := GenS11{F0: 0}
                                        _ = v250
                                        _ = v251
                                    } else {
                                        if (-786 != 42) {
                                            fmt.Println(v1, v1)
                                            v252 := GenS17{F0: 18446744073709551615}
                                            _ = v252
                                        } else {
                                            for fi253 := 0; fi253 < 2; fi253++ {
                                                switch (fi253 - (fi253 * fi253)) {
                                                case -734:
                                                    var v254 float64 = 1e10
                                                    _ = v254
                                                case 2147483647:
                                                    var v255 uint16 = 255
                                                    _ = v255
                                                }
                                                _ = fi253
                                            }
                                        }
                                        v256 := GenS2{F0: (-(v3)), F1: 1, F2: 42}
                                        _ = v256
                                    }
                                }
                            case "foo bar":
                                if ((1 < 255) && (1.5 < 1.5)) {
                                    v4 = append(v4, 953)
                                    if false {
                                        var v257 int8 = v1
                                        v3 = (v3 / (((v3 * v3)) | 1))
                                        switch 418 {
                                        case 32767:
                                            v258 := GenS7{F0: 741}
                                            _ = v258
                                        case 128:
                                            var v259 uint32 = 2147483647
                                            _ = v259
                                        case 257:
                                            v260 := GenS2{F0: -2147483648, F1: 1, F2: 530}
                                            _ = v260
                                        default:
                                            var v261 int = (-(v3))
                                            _ = v261
                                        }
                                        _ = v257
                                    } else {
                                        var v262 uint64 = 1
                                        _ = v262
                                    }
                                    var v263 float32 = (0.0 * ((-2.25 * -2.25) + 3.14159))
                                    _ = v263
                                } else {
                                    switch (v3 / ((genFn1(42, genFn0(), 0)) | 1)) {
                                    case -638:
                                        v3 = 128
                                    default:
                                        switch 9223372036854775807 {
                                        case -1:
                                            _ = len(v5)
                                        case -2147483648:
                                            _ = len(v2)
                                        case 901:
                                            switch (v3 & (^(v3))) {
                                            case 128:
                                                for fi264 := 0; fi264 < 10; fi264++ {
                                                    if (((3.14159 + float64(1e10)) == float64(0.1)) || ("a" == "foo bar")) {
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                    } else {
                                                        v265 := GenS6{F0: 880, F1: 534, F2: 2147483647}
                                                        _ = v265
                                                    }
                                                    _ = fi264
                                                }
                                            case 1:
                                                for fi266 := 0; fi266 < 2; fi266++ {
                                                    for fi267 := 0; fi267 < 11; fi267++ {
                                                        switch "foo bar" {
                                                        case "a":
                                                            v1 -= v1
                                                        case "hello":
                                                            v3 = 61
                                                        }
                                                        if false {
                                                            switch (("\n\t" + "a") + "") {
                                                            case "foo bar":
                                                                var v268 uint16 = 813
                                                                _ = v268
                                                            case "":
                                                                switch "\n\t" {
                                                                case "foo bar":
                                                                    v3 = genFn1(-272, 1, 4294967295)
                                                                case "\n\t":
                                                                    v3 += -2147483648
                                                                case "hello":
                                                                    var v269 float64 = float64((1e10 + (float32(((float64(1e10) - 1.5) * 1.5)) / (float32((-2.25 * -2.25)) - (0.0 + 0.1)))))
                                                                    _ = v269
                                                                default:
                                                                    v2 = append(v2, 184)
                                                                }
                                                            case "hello":
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                            }
                                                            switch "\n\t" {
                                                            case "":
                                                                var v270 uint8 = 2
                                                                _ = v270
                                                            default:
                                                                fmt.Println(v1)
                                                            }
                                                            if len(v4) > 0 {
                                                                v4[0] = 438
                                                            }
                                                        }
                                                        v5 = append(v5, 2)
                                                        _ = fi267
                                                    }
                                                    if len(v4) > 0 {
                                                        v4[0] = 287
                                                    }
                                                    v5 = append(v5, 0)
                                                    _ = fi266
                                                }
                                            case 0:
                                                if true {
                                                    v271 := GenS3{F0: "foo bar", F1: (v1 << ((281) & 63))}
                                                    _ = v271
                                                } else {
                                                    v272 := GenS1{F0: 540}
                                                    _ = v272
                                                }
                                            }
                                        default:
                                            v273 := GenS6{F0: 2, F1: 626, F2: -2147483648}
                                            _ = v273
                                        }
                                    }
                                }
                            }
                        }
                    }
                    v3 -= v3
                }
            } else {
                v274 := GenS12{F0: 974}
                v275 := GenS7{F0: 32767}
                v3 = v3
                _ = v274
                _ = v275
            }
        case "":
            for fi276 := 0; fi276 < 3; fi276++ {
                v277 := GenS2{F0: v3, F1: 0, F2: 4294967295}
                _ = fi276
                _ = v277
            }
        }
        v3 += v3
        if (!true) {
            fmt.Println(v3, v1)
            fmt.Println(v3, v3)
            if false {
                for fi278 := 0; fi278 < 9; fi278++ {
                    fmt.Println(v1, v1, v3)
                    _ = fi278
                }
            } else {
                fmt.Println(v3)
            }
        }
    }
    v3 = v3
    for fi279 := 0; fi279 < 6; fi279++ {
        for fi280 := 0; fi280 < 14; fi280++ {
            fmt.Println(v1, fi279)
            v5 = append(v5, 4294967295)
            v3 += (^(fi280))
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi280
        }
        if (true || ((("a" + "foo bar") + "") != "hello")) {
            for fi281 := 0; fi281 < 20; fi281++ {
                var v282 string = ""
                v3 -= genFn1(-385, 255, 0)
                if true {
                    fmt.Println(v1)
                }
                _ = fi281
                _ = v282
            }
        }
        v3 = (^(fi279))
        _ = fi279
    }
    var v283 int8 = v1
    fmt.Println(v3)
    if (-32768 >= 0) {
        fmt.Println(v283, v283, v1)
        var v284 bool = false
        switch "hello" {
        case "":
            if ((v3 + v3) <= (^(v3))) {
                var v285 uint8 = 0
                fmt.Println(v3, v283, v284)
                _ = v285
            } else {
                v284 = false
            }
        case "hello":
            fmt.Println(v283, v1, v284)
        case "foo bar":
            var v286 int = (-(v3))
            _ = v286
        default:
            var v287 int16 = -32768
            _ = v287
        }
        _ = v284
    }
    v283 += (v283 - (^(v1)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v283:", v283)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v283
    _ = v3
    _ = v4
    _ = v5
}
