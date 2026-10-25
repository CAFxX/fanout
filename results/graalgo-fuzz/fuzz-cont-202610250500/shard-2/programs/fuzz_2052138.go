package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint32 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 int; F1 uint64; F2 uint32 }
type GenS4 struct { F0 int16; F1 string; F2 string }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int32; F1 int64 }
type GenS8 struct { F0 uint8; F1 int32; F2 int }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int32; F1 int32 }

func main() {
    var v1 int = -858
    var v2 uint8 = 0
    var v3 uint16 = 42
    v4 := []uint32{857, 4294967295}
    v5 := make(map[uint32]int)
    var v6 int16 = -1
    v7 := GenS1{F0: -1, F1: ("" + "a"), F2: 42}
    v8 := GenS2{F0: 121, F1: (v1 % ((127) | 1))}
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 uint32 = 4294967295
        var v11 float32 = 1.5
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 float32 = 3.14159
    var v13 int16 = v6
    var v14 int32 = 2147483647
    switch (v1 * (v1 >> ((42) & 63))) {
    case -49:
        if true {
            if true {
                v15 := GenS3{F0: int(v6), F1: 2, F2: 2147483648}
                _ = v15
            }
            for fi16 := 0; fi16 < 17; fi16++ {
                if true {
                    var v17 uint8 = (v2 & 255)
                    _ = v17
                } else {
                    var v18 int64 = -886
                    var v19 uint64 = 703
                    _ = v18
                    _ = v19
                }
                if false {
                    for fi20 := 0; fi20 < 16; fi20++ {
                        fmt.Println(fi20)
                        if (496 >= -1) {
                            if (v12 != (v12 * v12)) {
                                switch v1 {
                                case 32768:
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                case -9223372036854775808:
                                    if (!(false && (v13 > v6))) {
                                        switch ((("\n\t" + "foo bar") + "a") + "hello") {
                                        case "\n\t":
                                            fmt.Println(fi16, v12, v6)
                                        case "foo bar":
                                            var v21 uint8 = (v2 - 245)
                                            _ = v21
                                        default:
                                            switch "foo bar" {
                                            case "hello":
                                                if true {
                                                    v22 := GenS1{F0: v14, F1: "", F2: 0}
                                                    v4 = append(v4, 1)
                                                    if ("hello" > ("\n\t" + ("" + "hello"))) {
                                                        v23 := GenS2{F0: -1, F1: fi20}
                                                        _ = v23
                                                    } else {
                                                        if (((-1 < 42) && true) && true) {
                                                            var v24 uint = 721
                                                            var v25 uint = v24
                                                            _ = v24
                                                            _ = v25
                                                        } else {
                                                            v2 = v2
                                                            if true {
                                                                fmt.Println(v1)
                                                            } else {
                                                                v1 -= fi20
                                                                if (!false) {
                                                                    fmt.Println(v14)
                                                                }
                                                                switch -2147483648 {
                                                                case -1:
                                                                    v14 -= (v14 >> ((4294967295) & 63))
                                                                default:
                                                                    v26 := GenS1{F0: (v14 & v14), F1: "a", F2: 42}
                                                                    _ = v26
                                                                }
                                                            }
                                                        }
                                                        switch ("foo bar" + "") {
                                                        case "hello":
                                                            if ((3.14159 + float64(v12)) == (3.14159 - (-2.25 - (-2.25 / float64((v12 + float32(1e10))))))) {
                                                                v6 = (v13 ^ v13)
                                                                switch v1 {
                                                                case -9223372036854775808:
                                                                    var v27 int8 = 1
                                                                    _ = v27
                                                                default:
                                                                    if true {
                                                                        v3 += (^(v3))
                                                                        var v28 bool = (!false)
                                                                        _ = v28
                                                                    } else {
                                                                        v6 = v6
                                                                    }
                                                                }
                                                            } else {
                                                                var v29 int = v1
                                                                if len(v4) > 0 {
                                                                    v4[0] = 408
                                                                }
                                                                _ = v29
                                                            }
                                                        case "a":
                                                            if (!(!(!true))) {
                                                                var v30 uint64 = 4294967295
                                                                _ = v30
                                                            }
                                                        default:
                                                            v12 += v12
                                                            switch "a" {
                                                            case "\n\t":
                                                                v13 = -298
                                                            default:
                                                                v5[0] = fi16
                                                            }
                                                        }
                                                    }
                                                    _ = v22
                                                } else {
                                                    v14 = 0
                                                }
                                            default:
                                                fmt.Println(v2)
                                            }
                                        }
                                        fmt.Println(v13)
                                        var v31 int64 = 42
                                        _ = v31
                                    } else {
                                        if true {
                                            var v32 uint16 = (v3 & (v3 >> ((25) & 63)))
                                            switch "\n\t" {
                                            case "foo bar":
                                                if (!true) {
                                                    v6 = (v6 << ((18446744073709551615) & 63))
                                                    _ = len(v5)
                                                    switch (v1 - fi16) {
                                                    case 4294967295:
                                                        fmt.Println(v3, v1)
                                                    }
                                                }
                                            case "":
                                                fmt.Println(fi20, v1)
                                            }
                                            switch "a" {
                                            case "foo bar":
                                                if true {
                                                    v33 := GenS1{F0: (v14 / ((-1) | 1)), F1: "foo bar", F2: 42}
                                                    v34 := GenS4{F0: v13, F1: "foo bar", F2: "a"}
                                                    var v35 uint16 = v32
                                                    _ = v33
                                                    _ = v34
                                                    _ = v35
                                                } else {
                                                    fmt.Println(v3)
                                                    var v36 int16 = v6
                                                    _ = v36
                                                }
                                            }
                                            var v37 int32 = (v14 / (((v14 + (v14 << ((18446744073709551615) & 63)))) | 1))
                                            var v38 int16 = 634
                                            if false {
                                                if false {
                                                    v37 += (v14 ^ int32(v13))
                                                    v39 := GenS4{F0: int16(v2), F1: "\n\t", F2: ("foo bar" + "\n\t")}
                                                    var v40 float64 = 0.1
                                                    _ = v39
                                                    _ = v40
                                                }
                                                switch "" {
                                                case "":
                                                    switch (-(v1)) {
                                                    case 42:
                                                        var v41 int = (v1 << ((4294967295) & 63))
                                                        _ = v41
                                                    case 9223372036854775807:
                                                        if (v6 < v38) {
                                                            if true {
                                                                v5[2147483648] = int(v37)
                                                            }
                                                            v6 = (v13 % (((^(v13))) | 1))
                                                        }
                                                        var v42 uint32 = 1
                                                        var v43 float64 = float64((1e10 / 1e10))
                                                        v44 := GenS3{F0: v1, F1: uint64(v3), F2: (v42 << ((42) & 63))}
                                                        var v45 uint64 = 1
                                                        _ = v42
                                                        _ = v43
                                                        _ = v44
                                                        _ = v45
                                                    case -9223372036854775808:
                                                        v46 := GenS2{F0: int8(v1), F1: 160}
                                                        _ = v46
                                                    default:
                                                        v4 = append(v4, 65535)
                                                    }
                                                default:
                                                    v13 = v6
                                                }
                                            }
                                            v47 := GenS5{F0: v32}
                                            _ = v32
                                            _ = v37
                                            _ = v38
                                            _ = v47
                                        } else {
                                            var v48 uint8 = (v2 * v2)
                                            var v49 int32 = (v14 | (v14 - (v14 >> ((4294967295) & 63))))
                                            fmt.Println(v2)
                                            _ = v48
                                            _ = v49
                                        }
                                    }
                                }
                                var v50 string = "foo bar"
                                _ = v50
                            } else {
                                v12 = v12
                                v13 += v13
                                if (2 == 42) {
                                    v51 := GenS4{F0: 32767, F1: "\n\t", F2: ""}
                                    v1 += fi16
                                    _ = v51
                                } else {
                                    fmt.Println(fi20, v13, v12)
                                    var v52 int8 = 42
                                    v53 := GenS3{F0: fi16, F1: 4294967295, F2: 207}
                                    _ = v52
                                    _ = v53
                                }
                                v54 := v5[39]
                                v55 := GenS2{F0: 1, F1: (fi16 ^ 4294967295)}
                                _ = v54
                                _ = v55
                            }
                        }
                        if (!(-1 >= 1)) {
                            var v56 bool = true
                            if v56 {
                                _ = len(v4)
                            } else {
                                v13 += 243
                            }
                            v1 = (^(fi16))
                            _ = v56
                        } else {
                            var v57 int16 = (v13 / ((v6) | 1))
                            if true {
                                fmt.Println(v57, v2, v2)
                                var v58 float64 = 1.5
                                switch (^(fi16)) {
                                case -9223372036854775808:
                                    if len(v4) > 0 {
                                        v4[0] = uint32(v2)
                                    }
                                case -2147483648:
                                    v59 := GenS2{F0: 1, F1: (-(v1))}
                                    _ = v59
                                case 130:
                                    fmt.Println(v57, v3, v2)
                                }
                                _ = v58
                            } else {
                                v57 = v13
                                var v60 uint = 65535
                                _ = v60
                            }
                            if (1.5 < 1e10) {
                                v61 := GenS3{F0: fi20, F1: 975, F2: uint32(v2)}
                                fmt.Println(fi20)
                                _ = v61
                            } else {
                                v62 := GenS6{F0: true}
                                _ = v62
                            }
                            _ = v57
                        }
                        _ = fi20
                    }
                    switch ("\n\t" + "a") {
                    case "\n\t":
                        v63 := GenS7{F0: v14, F1: 42}
                        _ = v63
                    case "a":
                        v2 += v2
                    default:
                        for fi64 := 0; fi64 < 7; fi64++ {
                            v65 := GenS3{F0: (v1 << ((0) & 63)), F1: 1, F2: 679}
                            v66 := GenS8{F0: v2, F1: v14, F2: fi16}
                            var v67 string = ("foo bar" + ("hello" + "foo bar"))
                            if (!false) {
                                v67 = (("hello" + "hello") + ((v67 + ((((v67 + v67) + v67) + ((v67 + v67) + "a")) + v67)) + v67))
                                v5[666] = (fi64 ^ fi16)
                                var v68 int32 = (v14 % (((-(v14))) | 1))
                                _ = v68
                            }
                            _ = fi64
                            _ = v65
                            _ = v66
                            _ = v67
                        }
                    }
                } else {
                    switch ("\n\t" + "") {
                    case "hello":
                        var v69 float64 = 0.0
                        _ = v69
                    case "a":
                        var v70 int8 = -18
                        _ = v70
                    default:
                        var v71 int = (fi16 * (v1 >> ((4294967295) & 63)))
                        _ = v71
                    }
                }
                _ = fi16
            }
            for fi72 := 0; fi72 < 6; fi72++ {
                switch fi72 {
                case -449:
                    var v73 float64 = float64((0.0 + (1e10 + (v12 / v12))))
                    _ = v73
                case 153:
                    v14 = (v14 + -345)
                }
                v14 = v14
                for fi74 := 0; fi74 < 6; fi74++ {
                    var v75 int16 = (v13 >> ((uint(v3)) & 63))
                    _ = fi74
                    _ = v75
                }
                _ = fi72
            }
        } else {
            for fi76 := 0; fi76 < 18; fi76++ {
                var v77 uint8 = v2
                _ = fi76
                _ = v77
            }
            if (1e10 == 1.5) {
                v78 := GenS1{F0: -2147483648, F1: "\n\t", F2: 0}
                var v79 int32 = v14
                var v80 float32 = ((float32(0.1) - ((v12 - v12) * v12)) * v12)
                _ = v78
                _ = v79
                _ = v80
            } else {
                if len(v4) > 0 {
                    v4[0] = 2147483647
                }
                v81 := GenS2{F0: 4, F1: v1}
                _ = v81
            }
        }
    case -1:
        for fi82 := 0; fi82 < 1; fi82++ {
            v2 += v2
            var v83 int16 = -1
            _ = fi82
            _ = v83
        }
    case 4294967295:
        var v84 int = (^(v1))
        _ = v84
    default:
        if (1e10 >= float64(v12)) {
            if (true && false) {
                switch "" {
                case "hello":
                    var v85 int32 = (-(v14))
                    _ = v85
                case "":
                    var v86 int64 = -799
                    _ = v86
                }
                var v87 int32 = v14
                v4 = append(v4, 117)
                _ = v87
            } else {
                v14 = 0
                v14 = -1
                v13 = (v6 ^ v6)
            }
        } else {
            if ((0.0 > 0.1) || (!(0 >= -304))) {
                v88 := GenS9{F0: (v2 * v2)}
                v13 = (v6 & v6)
                _ = v88
            } else {
                var v89 uint64 = 18446744073709551615
                if true {
                    var v90 float64 = (0.1 * (1.5 * (-2.25 * ((0.0 + 1.5) + 1.5))))
                    var v91 uint16 = (v3 % (((v3 | v3)) | 1))
                    _ = v90
                    _ = v91
                } else {
                    v14 = v14
                }
                var v92 uint = 18446744073709551615
                _ = v89
                _ = v92
            }
            v4 = append(v4, 65535)
            fmt.Println(v2, v6, v2)
        }
    }
    switch (^(v1)) {
    case -822:
        for fi93 := 0; fi93 < 14; fi93++ {
            if (-159 < -604) {
                v94 := v5[0]
                if (!false) {
                    var v95 uint32 = 424
                    _ = v95
                } else {
                    var v96 uint8 = (v2 * (v2 * 255))
                    if false {
                        if (true || true) {
                            v97 := GenS9{F0: 218}
                            var v98 uint16 = v3
                            _ = v97
                            _ = v98
                        } else {
                            var v99 bool = (0 > 1)
                            var v100 string = (("a" + "") + "a")
                            _ = v100
                            _ = v99
                        }
                        v94 -= v94
                    } else {
                        var v101 float64 = 3.14159
                        if false {
                            switch 1 {
                            case 127:
                                v14 -= (^(v14))
                            }
                        } else {
                            switch "foo bar" {
                            case "hello":
                                var v102 int16 = v6
                                _ = v102
                            case "":
                                switch "hello" {
                                case "":
                                    var v103 float32 = (((3.14159 + 0.0) + 1.5) + v12)
                                    _ = v103
                                case "a":
                                    if len(v4) > 0 {
                                        v4[0] = 2147483648
                                    }
                                case "\n\t":
                                    var v104 uint32 = 2147483647
                                    _ = v104
                                }
                            }
                        }
                        _ = v101
                    }
                    _ = v96
                }
                _ = v94
            }
            _ = fi93
        }
    case -366:
        var v105 uint16 = v3
        _ = v105
    case -9223372036854775808:
        switch "hello" {
        case "":
            switch v1 {
            case 0:
                var v106 float32 = 0.1
                _ = v106
            default:
                var v107 float32 = (-2.25 / 1.5)
                _ = v107
            }
        default:
            v108 := GenS6{F0: (-54 <= -1)}
            _ = v108
        }
    default:
        switch v1 {
        case 42:
            delete(v5, 65535)
        case 2147483647:
            switch (("" + "foo bar") + "foo bar") {
            case "foo bar":
                fmt.Println(v2)
            case "a":
                var v109 float64 = (1e10 / ((3.14159 + (-2.25 / (0.1 + (float64((v12 / float32(1.5))) - 1e10)))) - 0.1))
                _ = v109
            }
        case 128:
            if (1 != v14) {
                _ = len(v4)
                v110 := GenS2{F0: -91, F1: v1}
                _ = v110
            } else {
                switch v1 {
                case 0:
                    var v111 int64 = 0
                    _ = v111
                case -886:
                    fmt.Println(v13)
                }
                var v112 int16 = -1
                if true {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    fmt.Println(v3, v6)
                }
                _ = v112
            }
        default:
            for fi113 := 0; fi113 < 20; fi113++ {
                for fi114 := 0; fi114 < 13; fi114++ {
                    if true {
                        var v115 int32 = -179
                        _ = v115
                    }
                    v3 -= 959
                    _ = fi114
                }
                var v116 int32 = v14
                delete(v5, 2147483647)
                _ = fi113
                _ = v116
            }
        }
    }
    if (true && (float32(-2.25) != (((v12 + float32(3.14159)) / (1e10 * (v12 + 0.0))) - v12))) {
        switch -305 {
        case 9223372036854775807:
            v1 = (-(v1))
        case -9223372036854775808:
            fmt.Println(v13, v14)
        default:
            for fi117 := 0; fi117 < 14; fi117++ {
                for fi118 := 0; fi118 < 20; fi118++ {
                    v1 += v1
                    v13 = (v13 / (((v6 / ((v13) | 1))) | 1))
                    _ = fi118
                }
                _ = fi117
            }
        }
        switch -2147483648 {
        case 2147483647:
            var v119 uint64 = 1
            _ = v119
        }
    }
    v6 = int16(v2)
    v120 := GenS3{F0: v1, F1: 0, F2: 879}
    var v121 uint8 = (v2 / ((128) | 1))
    switch "foo bar" {
    case "\n\t":
        v6 -= v13
    default:
        if (18446744073709551615 != uint(v3)) {
            for fi122 := 0; fi122 < 10; fi122++ {
                v123 := GenS4{F0: v13, F1: ("\n\t" + ("\n\t" + "hello")), F2: ("a" + "\n\t")}
                _ = fi122
                _ = v123
            }
            v124 := GenS10{F0: -86, F1: (v14 + v14)}
            v5[2147483648] = 4294967295
            _ = v124
        }
    }
    var v125 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v121:", v121)
    fmt.Println("v125:", v125)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v120
    _ = v121
    _ = v125
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
