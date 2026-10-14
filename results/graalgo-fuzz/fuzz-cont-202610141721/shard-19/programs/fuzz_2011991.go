package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int; F1 int; F2 uint16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint32; F1 int; F2 uint32 }
type GenS7 struct { F0 uint; F1 int; F2 int8 }

func main() {
    var v1 int64 = -1
    var v2 int16 = 1000
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]bool)
    for fi5 := 0; fi5 < 12; fi5++ {
        v1 = v1
        for fi6 := 0; fi6 < 20; fi6++ {
            if (2147483647 != -2147483648) {
                var v7 bool = ("foo bar" > "hello")
                v8 := GenS1{F0: v1, F1: ""}
                var v9 int64 = (v1 ^ (v1 ^ v1))
                _ = v7
                _ = v8
                _ = v9
            }
            if (float32(float64(float32((((3.14159 - -2.25) / ((-2.25 - 3.14159) * -2.25)) + (1.5 * (1.5 - (1.5 / ((1e10 / float64((3.14159 / 1.0))) - (3.14159 - float64(0.0)))))))))) >= (float32(1e10) - (3.14159 + (3.14159 * (((0.0 + 1.5) - 1e10) / 1.5))))) {
                fmt.Println(v1, v1)
            }
            v1 -= (v1 | v1)
            _ = fi6
        }
        _ = fi5
    }
    v2 = -1
    switch "a" {
    case "":
        for fi10 := 0; fi10 < 17; fi10++ {
            if (!(191 != 42)) {
                v1 = 205
            } else {
                switch -2147483648 {
                case 2147483647:
                    if (!(false && true)) {
                        var v11 uint = 65535
                        for fi12 := 0; fi12 < 10; fi12++ {
                            var v13 string = ("a" + "a")
                            if (true && true) {
                                v14 := v3[0]
                                _ = v14
                            }
                            v15 := []int{(fi12 << ((v11) & 63)), (fi12 | fi12), 32768, 127}
                            _ = fi12
                            _ = v13
                            _ = v15
                        }
                        _ = v11
                    }
                default:
                    v16 := v3[65535]
                    _ = v16
                }
            }
            _ = fi10
        }
    case "a":
        if false {
            switch 1 {
            case -9223372036854775808:
                switch -9223372036854775808 {
                case 9223372036854775807:
                    for fi17 := 0; fi17 < 20; fi17++ {
                        switch "a" {
                        case "":
                            var v18 float32 = (((float32(((0.0 + ((0.0 - (0.0 / 1.0)) / (3.14159 + -2.25))) - (3.14159 / 1e10))) + 1e10) + float32(-2.25)) * 1.5)
                            _ = v18
                        default:
                            switch -507 {
                            case -775:
                                if (((!("" >= "foo bar")) || (false && true)) && true) {
                                    v1 = -564
                                    v2 += (-(v2))
                                }
                            case 127:
                                var v19 int8 = -1
                                _ = v19
                            default:
                                for fi20 := 0; fi20 < 1; fi20++ {
                                    fmt.Println(v2)
                                    _ = fi20
                                }
                            }
                        }
                        var v21 int = (fi17 << ((472) & 63))
                        _ = fi17
                        _ = v21
                    }
                case -2147483648:
                    switch 32768 {
                    case -9223372036854775808:
                        var v22 uint8 = 60
                        _ = v22
                    case 127:
                        v1 = int64(v2)
                    case 9223372036854775807:
                        var v23 int64 = (v1 / ((v1) | 1))
                        _ = v23
                    }
                }
            case 1:
                _ = len(v3)
            case 2147483647:
                var v24 int16 = v2
                _ = v24
            default:
                var v25 uint = 65535
                _ = v25
            }
            fmt.Println(v2, v2)
            fmt.Println(v1)
        }
    }
    for fi26 := 0; fi26 < 15; fi26++ {
        v27 := []int{(fi26 | fi26)}
        var v28 uint64 = 685
        _ = fi26
        _ = v27
        _ = v28
    }
    var v29 int8 = -128
    for fi30 := 0; fi30 < 2; fi30++ {
        for fi31 := 0; fi31 < 5; fi31++ {
            if true {
                fmt.Println(fi31, v2)
                v32 := GenS1{F0: (v1 << ((65535) & 63)), F1: ""}
                v33 := GenS1{F0: (v1 << ((18446744073709551615) & 63)), F1: "a"}
                _ = v32
                _ = v33
            }
            switch (("\n\t" + "") + "hello") {
            case "\n\t":
                fmt.Println(v2, v1)
            }
            _ = len(v4)
            _ = fi31
        }
        fmt.Println(v2, v2)
        _ = fi30
    }
    if true {
        if false {
            v29 -= 127
            switch ("" + "") {
            case "hello":
                v2 = 0
            case "\n\t":
                switch 245 {
                case 0:
                    v4["foo bar"] = ((((0.0 > float32(1e10)) || (v1 != v1)) && true) && false)
                }
            case "foo bar":
                switch 2147483647 {
                case -1:
                    if false {
                        if false {
                            v1 += int64(v2)
                            switch ("" + ("" + ("\n\t" + "hello"))) {
                            case "\n\t":
                                var v34 string = ("a" + ("\n\t" + "foo bar"))
                                _ = v34
                            }
                        }
                        var v35 int64 = v1
                        _ = v35
                    }
                }
            default:
                v36 := GenS2{F0: "a", F1: 42}
                _ = v36
            }
            fmt.Println(v2)
        }
        v29 = v29
        fmt.Println(v2, v29, v29)
    } else {
        for fi37 := 0; fi37 < 1; fi37++ {
            var v38 uint64 = 1
            for fi39 := 0; fi39 < 12; fi39++ {
                if true {
                    if true {
                        if false {
                            var v40 uint = 1
                            _ = v40
                        } else {
                            if true {
                                v41 := []int{fi39}
                                _ = v41
                            }
                            var v42 uint32 = 42
                            _ = v42
                        }
                        var v43 bool = (128 < 128)
                        _ = v43
                    }
                    fmt.Println(v1)
                }
                v44 := GenS3{F0: ((65535 < 350) || false)}
                _ = fi39
                _ = v44
            }
            _ = fi37
            _ = v38
        }
    }
    if (true || (!(2 <= 456))) {
        if true {
            v1 += (^(v1))
            if true {
                var v45 int8 = v29
                for fi46 := 0; fi46 < 16; fi46++ {
                    v1 = 247
                    v47 := GenS1{F0: v1, F1: "hello"}
                    _ = fi46
                    _ = v47
                }
                _ = v45
            } else {
                if true {
                    if ((0.0 - 0.1) >= float64(0.0)) {
                        var v48 string = (("" + ("hello" + "a")) + "")
                        v49 := []int{128}
                        _ = v48
                        _ = v49
                    } else {
                        fmt.Println(v2, v1, v29)
                        for fi50 := 0; fi50 < 1; fi50++ {
                            v51 := GenS2{F0: ("" + ("hello" + "a")), F1: (^(fi50))}
                            switch (-(fi50)) {
                            case 128:
                                v2 = v2
                            }
                            v2 += v2
                            _ = fi50
                            _ = v51
                        }
                    }
                    v52 := GenS2{F0: (("\n\t" + "\n\t") + ("hello" + "\n\t")), F1: 0}
                    v2 = (v2 % (((v2 ^ v2)) | 1))
                    _ = v52
                } else {
                    switch ("foo bar" + ("foo bar" + (("hello" + "") + ""))) {
                    case "\n\t":
                        v1 -= (v1 >> ((4294967295) & 63))
                    case "a":
                        var v53 int8 = (v29 - (v29 + v29))
                        _ = v53
                    default:
                        var v54 float32 = float32((-2.25 * (3.14159 - (-2.25 - 0.0))))
                        _ = v54
                    }
                    v2 -= v2
                    fmt.Println(v29, v1)
                }
                var v55 float32 = (1e10 / -2.25)
                switch "foo bar" {
                case "":
                    if false {
                        if true {
                            v56 := []int{-9223372036854775808, 1, 9223372036854775807, 32768}
                            _ = v56
                        } else {
                            if false {
                                v29 = 0
                                v57 := GenS1{F0: (v1 ^ v1), F1: "foo bar"}
                                _ = v57
                            } else {
                                v1 = 783
                            }
                        }
                    }
                }
                _ = v55
            }
            if false {
                switch -2147483648 {
                case 9223372036854775807:
                    delete(v4, "foo bar")
                case 1:
                    v29 += v29
                }
                fmt.Println(v29, v2, v1)
                _ = len(v3)
            }
        } else {
            var v58 uint16 = 255
            _ = len(v4)
            _ = v58
        }
        var v59 string = "a"
        v59 = "hello"
        _ = v59
    } else {
        for fi60 := 0; fi60 < 16; fi60++ {
            v61 := []int{(fi60 / ((fi60) | 1))}
            _ = fi60
            _ = v61
        }
    }
    var v62 uint64 = 432
    var v63 int64 = v1
    if (65535 != 42) {
        for fi64 := 0; fi64 < 5; fi64++ {
            var v65 uint = 255
            if true {
                if false {
                    switch fi64 {
                    case -2147483648:
                        var v66 uint16 = 2
                        _ = v66
                    case 236:
                        for fi67 := 0; fi67 < 4; fi67++ {
                            if false {
                                fmt.Println(v62, fi67)
                                var v68 int64 = v63
                                var v69 int = fi64
                                _ = v68
                                _ = v69
                            }
                            if true {
                                v70 := GenS4{F0: -1, F1: (fi64 << ((v65) & 63)), F2: 32767}
                                switch fi67 {
                                case -93:
                                    v63 = (v1 / (((v1 ^ v1)) | 1))
                                case 0:
                                    if ((^(v29)) < v29) {
                                        var v71 uint32 = 819
                                        v72 := v3[v71]
                                        v73 := GenS2{F0: "\n\t", F1: fi67}
                                        v63 += v63
                                        _ = v71
                                        _ = v72
                                        _ = v73
                                    } else {
                                        var v74 bool = true
                                        v75 := []int{fi67, 42}
                                        switch "\n\t" {
                                        case "foo bar":
                                            var v76 int8 = 42
                                            _ = v76
                                        }
                                        _ = v74
                                        _ = v75
                                    }
                                case 42:
                                    var v77 uint64 = 4294967295
                                    _ = v77
                                default:
                                    if true {
                                        var v78 uint = v65
                                        var v79 float32 = 3.14159
                                        var v80 uint = 255
                                        v81 := GenS5{F0: v29}
                                        v82 := v3[571]
                                        _ = v78
                                        _ = v79
                                        _ = v80
                                        _ = v81
                                        _ = v82
                                    }
                                }
                                _ = v70
                            } else {
                                fmt.Println(fi67)
                                if true {
                                    fmt.Println(v63)
                                    var v83 uint = 18446744073709551615
                                    v2 += v2
                                    if (false || false) {
                                        fmt.Println(v62, v65)
                                        v29 += (v29 ^ v29)
                                        v63 += -644
                                    }
                                    if true {
                                        if (true || true) {
                                            if (!true) {
                                                v84 := []int{fi67}
                                                if (true || ((3.14159 == (0.0 / 1.5)) && (0 != 867))) {
                                                    v2 = -1
                                                }
                                                var v85 uint16 = 1
                                                _ = v84
                                                _ = v85
                                            }
                                        }
                                    }
                                    _ = v83
                                }
                                fmt.Println(v65)
                            }
                            _ = fi67
                        }
                    case -9223372036854775808:
                        for fi86 := 0; fi86 < 4; fi86++ {
                            v63 = v63
                            switch fi64 {
                            case 4294967295:
                                if false {
                                    v87 := GenS3{F0: false}
                                    switch (^(fi64)) {
                                    case 4294967295:
                                        var v88 uint64 = 2
                                        _ = v88
                                    case 996:
                                        var v89 int32 = -1
                                        _ = v89
                                    }
                                    if ((v29 & int8(v65)) <= v29) {
                                        v90 := GenS6{F0: 936, F1: (fi64 + -9223372036854775808), F2: 4294967295}
                                        if (((("foo bar" + "") + "") + ((("foo bar" + "a") + ("foo bar" + (("" + "foo bar") + "a"))) + "a")) >= "\n\t") {
                                            switch fi86 {
                                            case 2147483647:
                                                v91 := GenS5{F0: int8(v2)}
                                                _ = v91
                                            case 911:
                                                _ = len(v3)
                                            case 1:
                                                var v92 uint64 = v62
                                                _ = v92
                                            }
                                            v63 = v63
                                        } else {
                                            if true {
                                                fmt.Println(v1, v1)
                                                if (!false) {
                                                    switch ("hello" + ("\n\t" + ("" + "foo bar"))) {
                                                    case "\n\t":
                                                        v63 = -85
                                                    }
                                                    var v93 bool = false
                                                    v94 := []int{fi64, (fi86 | 1), fi64}
                                                    _ = v93
                                                    _ = v94
                                                } else {
                                                    v29 -= (v29 % ((v29) | 1))
                                                }
                                                var v95 uint16 = 0
                                                _ = v95
                                            } else {
                                                var v96 int32 = 970
                                                _ = v96
                                            }
                                            if false {
                                                v97 := GenS7{F0: v65, F1: (fi86 ^ 32768), F2: (^(v29))}
                                                if true {
                                                    v29 -= 42
                                                    if ((false || ((v62 != (v62 ^ (v62 | (v62 | (v62 ^ (^(v62))))))) || (-84 <= 42))) && ((v29 - 42) == (v29 % ((v29) | 1)))) {
                                                        v98 := []int{fi86, -2147483648, -507}
                                                        v99 := GenS5{F0: (v29 | v29)}
                                                        if len(v98) > 0 {
                                                            v98[0] = (fi86 | -824)
                                                        }
                                                        _ = v98
                                                        _ = v99
                                                    }
                                                    v100 := GenS2{F0: ("hello" + "\n\t"), F1: (fi86 - int(v2))}
                                                    _ = v100
                                                } else {
                                                    var v101 uint8 = 42
                                                    var v102 int16 = (v2 << ((0) & 63))
                                                    var v103 int8 = (v29 & v29)
                                                    _ = v101
                                                    _ = v102
                                                    _ = v103
                                                }
                                                _ = v97
                                            } else {
                                                v3[688] = 2147483648
                                                _ = len(v4)
                                            }
                                            fmt.Println(fi64, v63)
                                        }
                                        _ = v90
                                    }
                                    _ = v87
                                } else {
                                    var v104 int8 = (^(v29))
                                    _ = v104
                                }
                            case 2147483647:
                                _ = len(v3)
                            case -4:
                                fmt.Println(v1, v62, fi86)
                            }
                            var v105 uint8 = 42
                            _ = fi86
                            _ = v105
                        }
                    default:
                        if ((!true) && true) {
                            switch (fi64 / (((fi64 | 9223372036854775807)) | 1)) {
                            case 32767:
                                var v106 int = (fi64 * fi64)
                                _ = v106
                            }
                            var v107 int8 = (v29 & (v29 | v29))
                            if (2 > 2147483648) {
                                var v108 uint = 1
                                if true {
                                    v29 = v107
                                    var v109 uint64 = (v62 << ((v65) & 63))
                                    _ = v109
                                } else {
                                    v1 = (v63 - v1)
                                    _ = len(v4)
                                    var v110 uint = 255
                                    _ = v110
                                }
                                v63 -= v63
                                _ = v108
                            } else {
                                var v111 uint32 = 65535
                                var v112 int8 = (v29 + v107)
                                v1 += 9223372036854775807
                                _ = v111
                                _ = v112
                            }
                            _ = v107
                        }
                    }
                    var v113 int32 = 42
                    _ = v113
                }
            } else {
                for fi114 := 0; fi114 < 8; fi114++ {
                    fmt.Println(fi64, fi114)
                    var v115 uint8 = 2
                    _ = fi114
                    _ = v115
                }
            }
            _ = fi64
            _ = v65
        }
        var v116 uint16 = 707
        _ = v116
    }
    var v117 uint32 = 65535
    fmt.Println(v1)
    v2 += (v2 / ((v2) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v117:", v117)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v117
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v62
    _ = v63
}
