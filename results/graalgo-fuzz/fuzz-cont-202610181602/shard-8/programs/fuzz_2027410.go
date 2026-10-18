package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint }
type GenS2 struct { F0 uint16; F1 int16; F2 uint32 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int8; F1 int; F2 int16 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint; F1 int8 }
type GenS10 struct { F0 int; F1 int8; F2 int64 }
type GenS11 struct { F0 int; F1 uint; F2 int16 }
type GenS12 struct { F0 uint64; F1 int32; F2 string }
type GenS13 struct { F0 uint; F1 int; F2 uint }
type GenS14 struct { F0 int32; F1 uint32 }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 int32; F1 int8 }
type GenS17 struct { F0 uint64 }
type GenS18 struct { F0 bool; F1 int }
func genFn0(p0 int) uint8 {
    p0 = (^(p0))
    for fi5 := 0; fi5 < 12; fi5++ {
        switch p0 {
        case 502:
            var v6 string = "hello"
            _ = v6
        case 42:
            var v7 int16 = 32767
            _ = v7
        }
        switch "hello" {
        case "":
            if false {
                var v8 uint8 = 1
                var v9 float32 = ((0.0 + float32(1e10)) / (((1.5 / 1.0) + (((3.14159 + (((0.0 + (0.1 + 0.1)) - (float32((1.5 / 1.0)) / 1.5)) / (0.0 - 3.14159))) + (1e10 / 3.14159)) / -2.25)) * (float32(1.5) - 0.0)))
                var v10 bool = false
                _ = v10
                _ = v8
                _ = v9
            } else {
                switch "\n\t" {
                case "a":
                    var v11 float32 = 0.0
                    _ = v11
                case "\n\t":
                    p0 = (fi5 / ((fi5) | 1))
                case "hello":
                    for fi12 := 0; fi12 < 14; fi12++ {
                        switch p0 {
                        case 424:
                            p0 -= 42
                        case -359:
                            p0 = (p0 * (fi12 & 32768))
                        case -1:
                            var v13 int64 = 923
                            _ = v13
                        }
                        _ = fi12
                    }
                default:
                    var v14 uint16 = 1
                    _ = v14
                }
            }
        default:
            if (1 != -236) {
                if (141 != 127) {
                    p0 = (-(fi5))
                    v15 := []int{fi5, (p0 << ((18446744073709551615) & 63)), (p0 << ((1) & 63))}
                    _ = v15
                }
                var v16 uint = 42
                _ = v16
            } else {
                switch (^(fi5)) {
                case 0:
                    switch "hello" {
                    case "\n\t":
                        if true {
                            for fi17 := 0; fi17 < 15; fi17++ {
                                switch (p0 % (((-(fi5))) | 1)) {
                                case 9223372036854775807:
                                    var v18 uint8 = 111
                                    _ = v18
                                case 127:
                                    switch fi5 {
                                    case -617:
                                        v19 := []int{}
                                        switch p0 {
                                        case 4294967295:
                                            switch (fi17 & fi5) {
                                            case 282:
                                                switch ("hello" + (("a" + "") + "")) {
                                                case "":
                                                    switch 9223372036854775807 {
                                                    case 32768:
                                                        if (-447 != 0) {
                                                            var v20 uint = 2
                                                            _ = v20
                                                        } else {
                                                            _ = len(v19)
                                                        }
                                                    case -9223372036854775808:
                                                        p0 -= fi5
                                                    default:
                                                        var v21 uint32 = 0
                                                        _ = v21
                                                    }
                                                default:
                                                    if false {
                                                        if ((2147483647 <= 2) && true) {
                                                            var v22 uint64 = 18446744073709551615
                                                            _ = v22
                                                        } else {
                                                            switch "hello" {
                                                            case "hello":
                                                                p0 = fi17
                                                            case "":
                                                                v23 := GenS1{F0: 255, F1: 1, F2: 623}
                                                                _ = v23
                                                            case "foo bar":
                                                                p0 -= (fi17 << ((255) & 63))
                                                            }
                                                            var v24 uint8 = 42
                                                            v25 := GenS2{F0: 0, F1: 32767, F2: 748}
                                                            _ = v24
                                                            _ = v25
                                                        }
                                                        if true {
                                                            fmt.Println(fi17, fi17)
                                                            v26 := GenS3{F0: "", F1: 88}
                                                            if len(v19) >= 2 {
                                                                v19 = v19[1:len(v19)]
                                                            }
                                                            _ = v26
                                                        } else {
                                                            if false {
                                                                fmt.Println(p0, fi17)
                                                                var v27 uint16 = 255
                                                                _ = v27
                                                            }
                                                            v19 = append(v19, 32767)
                                                        }
                                                    } else {
                                                        var v28 float64 = 1.5
                                                        if (false || false) {
                                                            v29 := GenS4{F0: 876, F1: "a"}
                                                            if false {
                                                                var v30 int16 = 0
                                                                _ = v30
                                                            } else {
                                                                v31 := GenS5{F0: 42}
                                                                v32 := GenS4{F0: -2147483648, F1: ""}
                                                                v33 := GenS5{F0: 419}
                                                                var v34 uint16 = 32767
                                                                _ = v31
                                                                _ = v32
                                                                _ = v33
                                                                _ = v34
                                                            }
                                                            _ = v29
                                                        }
                                                        switch "hello" {
                                                        case "foo bar":
                                                            var v35 float64 = v28
                                                            _ = v35
                                                        case "\n\t":
                                                            p0 += (-(fi17))
                                                        }
                                                        if (true && (true && ((0 == 1) && false))) {
                                                            fmt.Println(fi5)
                                                            switch "" {
                                                            case "foo bar":
                                                                p0 -= 329
                                                            default:
                                                                var v36 float64 = float64(float32(v28))
                                                                _ = v36
                                                            }
                                                            v28 -= v28
                                                        } else {
                                                            var v37 float32 = -2.25
                                                            var v38 int64 = 0
                                                            if len(v19) >= 2 {
                                                                v19 = v19[1:len(v19)]
                                                            }
                                                            var v39 string = "hello"
                                                            v39 = ((v39 + "hello") + (v39 + v39))
                                                            _ = v37
                                                            _ = v38
                                                            _ = v39
                                                        }
                                                        _ = v28
                                                    }
                                                    if len(v19) >= 2 {
                                                        v19 = v19[1:len(v19)]
                                                    }
                                                    p0 -= p0
                                                }
                                                if (52 != -128) {
                                                    p0 = (fi17 >> ((669) & 63))
                                                    v40 := GenS1{F0: 32767, F1: 42, F2: 255}
                                                    var v41 uint32 = 2147483647
                                                    _ = v40
                                                    _ = v41
                                                }
                                            case 127:
                                                if (!true) {
                                                    var v42 uint8 = 208
                                                    _ = v42
                                                }
                                            default:
                                                var v43 int16 = 1000
                                                _ = v43
                                            }
                                        default:
                                            v44 := GenS5{F0: 4294967295}
                                            _ = v44
                                        }
                                        _ = v19
                                    case 740:
                                        if true {
                                            v45 := GenS5{F0: 588}
                                            if (true || false) {
                                                p0 -= (fi17 % (((^(fi17))) | 1))
                                            } else {
                                                v46 := GenS6{F0: 0}
                                                v47 := []int{(fi17 + 2147483647)}
                                                _ = v46
                                                _ = v47
                                            }
                                            p0 -= fi5
                                            _ = v45
                                        }
                                    case 0:
                                        if (!true) {
                                            v48 := GenS3{F0: "foo bar", F1: 0}
                                            fmt.Println(fi5)
                                            _ = v48
                                        }
                                    }
                                default:
                                    fmt.Println(p0, p0, p0)
                                }
                                _ = fi17
                            }
                        }
                    case "a":
                        var v49 int = 127
                        _ = v49
                    case "foo bar":
                        switch (p0 & fi5) {
                        case 309:
                            p0 = (fi5 + (fi5 >> ((42) & 63)))
                        case -1:
                            p0 += (p0 - (^(fi5)))
                        }
                    default:
                        var v50 int16 = -32768
                        _ = v50
                    }
                case -2147483648:
                    v51 := GenS7{F0: 0, F1: (fi5 + (p0 / (((p0 + fi5)) | 1))), F2: 42}
                    _ = v51
                default:
                    p0 += (fi5 * (p0 / (((p0 + (p0 - -1))) | 1)))
                }
                fmt.Println(p0, fi5)
            }
        }
        _ = fi5
    }
    switch "hello" {
    case "hello":
        v52 := GenS3{F0: "hello", F1: 42}
        _ = v52
    case "a":
        v53 := GenS8{F0: 18446744073709551615}
        _ = v53
    case "\n\t":
        p0 = p0
    }
    _ = p0
    return 210
}

func main() {
    var v1 int = -9223372036854775808
    var v2 uint8 = 14
    v3 := []int{-350, 1, 42, (v1 + -2147483648)}
    v4 := make(map[string]string)
    if len(v3) > 0 {
        v3[0] = v1
    }
    v2 = v2
    v54 := GenS9{F0: uint(v2), F1: 102}
    if ((-2.25 + 1e10) <= -2.25) {
        if true {
            switch "foo bar" {
            case "hello":
                v1 -= int(v2)
            case "a":
                var v55 int64 = -536
                _ = v55
            case "":
                var v56 uint8 = 210
                _ = v56
            default:
                if false {
                    fmt.Println(v2, v1)
                    if false {
                        v2 += (v2 + genFn0(v1))
                        switch v1 {
                        case 32767:
                            var v57 uint64 = 2
                            _ = v57
                        }
                        var v58 string = (("hello" + "foo bar") + ((("a" + "\n\t") + "hello") + ""))
                        _ = v58
                    }
                    delete(v4, ((("foo bar" + "a") + "hello") + "\n\t"))
                }
            }
        } else {
            v2 = v2
            v2 = (v2 | 127)
            for fi59 := 0; fi59 < 14; fi59++ {
                switch (v1 | (fi59 - (fi59 - -766))) {
                case 1:
                    v2 += genFn0(fi59)
                case 9223372036854775807:
                    v1 = v1
                }
                v1 += fi59
                _ = len(v3)
                _ = fi59
            }
        }
    } else {
        for fi60 := 0; fi60 < 1; fi60++ {
            _ = len(v3)
            v2 = 0
            v1 += (fi60 << ((638) & 63))
            _ = fi60
        }
    }
    var v61 uint64 = 386
    var v62 string = "hello"
    v1 = v1
    if (65535 == 42) {
        if (!(((v62 + "\n\t") + v62) <= "")) {
            for fi63 := 0; fi63 < 7; fi63++ {
                switch v62 {
                case "\n\t":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case "hello":
                    for fi64 := 0; fi64 < 3; fi64++ {
                        v61 += v61
                        switch (v62 + v62) {
                        case "":
                            v65 := GenS10{F0: 100, F1: 42, F2: 42}
                            _ = v65
                        }
                        _ = fi64
                    }
                case "foo bar":
                    _ = len(v4)
                default:
                    switch (v62 + v62) {
                    case "a":
                        if (!(!((true || false) || false))) {
                            v2 -= v2
                            var v66 int = (v1 + -989)
                            _ = v66
                        } else {
                            switch "hello" {
                            case "":
                                if (!(0.0 >= (0.1 / float32(-2.25)))) {
                                    if (!(v61 < 42)) {
                                        v67 := GenS2{F0: 32767, F1: -32768, F2: 2147483648}
                                        _ = v67
                                    }
                                    if true {
                                        var v68 uint = 735
                                        switch "\n\t" {
                                        case "foo bar":
                                            for fi69 := 0; fi69 < 10; fi69++ {
                                                v62 = (("hello" + "a") + v62)
                                                _ = fi69
                                            }
                                        case "hello":
                                            v70 := GenS11{F0: fi63, F1: (^(v68)), F2: -465}
                                            _ = v70
                                        }
                                        _ = v68
                                    }
                                    v1 = (fi63 + (fi63 / ((42) | 1)))
                                }
                            }
                            v2 -= uint8(v61)
                            var v71 int8 = 100
                            _ = v71
                        }
                    default:
                        v2 = v2
                    }
                }
                v72 := GenS2{F0: 65535, F1: 6, F2: 0}
                _ = fi63
                _ = v72
            }
            switch v62 {
            case "a":
                switch 42 {
                case 1:
                    for fi73 := 0; fi73 < 13; fi73++ {
                        if true {
                            fmt.Println(v1)
                            var v74 uint32 = 593
                            _ = len(v4)
                            _ = v74
                        }
                        v75 := GenS12{F0: v61, F1: -2147483648, F2: v62}
                        _ = fi73
                        _ = v75
                    }
                case 557:
                    switch v62 {
                    case "hello":
                        v76 := GenS4{F0: 706, F1: v62}
                        _ = v76
                    case "a":
                        for fi77 := 0; fi77 < 3; fi77++ {
                            var v78 float64 = 1e10
                            v62 = v62
                            _ = fi77
                            _ = v78
                        }
                    case "":
                        for fi79 := 0; fi79 < 20; fi79++ {
                            if ((1 >= -939) || ((v1 + (fi79 << ((65535) & 63))) > -759)) {
                                if false {
                                    if false {
                                        if true {
                                            fmt.Println(fi79)
                                            v1 = v1
                                            switch (v1 - (fi79 ^ (fi79 * (fi79 | v1)))) {
                                            case 1:
                                                v2 = v2
                                            }
                                        } else {
                                            var v80 float32 = 0.0
                                            var v81 uint = 65535
                                            if false {
                                                if false {
                                                    if true {
                                                        fmt.Println(v62, v61, fi79)
                                                    }
                                                    fmt.Println(v80)
                                                } else {
                                                    v82 := GenS13{F0: 65535, F1: -2147483648, F2: (v81 % (((v81 & (v81 % ((uint(v2)) | 1)))) | 1))}
                                                    _ = v82
                                                }
                                            } else {
                                                for fi83 := 0; fi83 < 3; fi83++ {
                                                    if len(v3) > 0 {
                                                        v3[0] = (v1 | fi83)
                                                    }
                                                    _ = fi83
                                                }
                                                _ = len(v3)
                                            }
                                            _ = v80
                                            _ = v81
                                        }
                                    }
                                    for fi84 := 0; fi84 < 10; fi84++ {
                                        var v85 bool = (v1 < -2147483648)
                                        v1 = (v1 + 4294967295)
                                        var v86 int16 = 42
                                        _ = fi84
                                        _ = v85
                                        _ = v86
                                    }
                                }
                                for fi87 := 0; fi87 < 16; fi87++ {
                                    v4[v62] = v62
                                    v1 = (^(v1))
                                    v2 = 19
                                    if len(v3) > 0 {
                                        v3[0] = (fi87 / ((v1) | 1))
                                    }
                                    _ = fi87
                                }
                            }
                            fmt.Println(fi79)
                            _ = fi79
                        }
                    }
                case -609:
                    v88 := GenS14{F0: -310, F1: 2}
                    _ = v88
                default:
                    if false {
                        v4[v62] = v62
                    }
                }
            case "\n\t":
                fmt.Println(v61, v1)
            case "foo bar":
                var v89 uint32 = 2147483647
                _ = v89
            }
            var v90 int16 = 365
            _ = v90
        } else {
            for fi91 := 0; fi91 < 4; fi91++ {
                v2 -= (v2 % (((v2 % ((genFn0((v1 - fi91))) | 1))) | 1))
                v62 = "foo bar"
                fmt.Println(v62, v62, v61)
                _ = fi91
            }
        }
        fmt.Println(v62)
        v61 = (v61 / (((v61 ^ (v61 ^ v61))) | 1))
    } else {
        if false {
            fmt.Println(v62, v62)
        }
        for fi92 := 0; fi92 < 15; fi92++ {
            v3 = append(v3, (fi92 % (((^(v1))) | 1)))
            for fi93 := 0; fi93 < 1; fi93++ {
                v94 := GenS6{F0: v61}
                _ = fi93
                _ = v94
            }
            var v95 uint16 = 839
            _ = fi92
            _ = v95
        }
    }
    if true {
        v2 = uint8(v61)
        v96 := GenS5{F0: 0}
        _ = v96
    }
    if ((v61 / ((v61) | 1)) >= (v61 >> ((0) & 63))) {
        v62 = ""
        for fi97 := 0; fi97 < 2; fi97++ {
            fmt.Println(v61)
            v98 := GenS1{F0: 65535, F1: -885, F2: 688}
            _ = fi97
            _ = v98
        }
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    if (!(false && ((v2 >> ((0) & 63)) != v2))) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        fmt.Println(v61, v61, v62)
    }
    var v99 uint = uint(v61)
    v99 -= v99
    fmt.Println(v1)
    v2 = (v2 - v2)
    var v100 uint16 = 255
    for fi101 := 0; fi101 < 13; fi101++ {
        var v102 int16 = -572
        fmt.Println(v102, v102, v100)
        fmt.Println(v1)
        _ = fi101
        _ = v102
    }
    switch v1 {
    case -608:
        if (false || (!true)) {
            var v103 int32 = 251
            _ = v103
        }
    case -9223372036854775808:
        v104 := GenS12{F0: v61, F1: 1, F2: ""}
        _ = v104
    default:
        if (!((v2 < 187) && ((v2 << (((v99 ^ (v99 / ((65535) | 1)))) & 63)) < (v2 % (((v2 << ((2) & 63))) | 1))))) {
            v105 := GenS10{F0: v1, F1: 127, F2: 1}
            _ = v105
        } else {
            for fi106 := 0; fi106 < 16; fi106++ {
                for fi107 := 0; fi107 < 13; fi107++ {
                    if true {
                        v108 := GenS15{F0: -1}
                        v109 := GenS16{F0: -2147483648, F1: 1}
                        var v110 int16 = 1000
                        _ = v108
                        _ = v109
                        _ = v110
                    } else {
                        v4[v62] = ""
                        var v111 int8 = 42
                        _ = v111
                    }
                    var v112 bool = true
                    _ = fi107
                    _ = v112
                }
                _ = fi106
            }
            v113 := GenS17{F0: v61}
            for fi114 := 0; fi114 < 17; fi114++ {
                var v115 uint64 = (v61 * 697)
                if false {
                    var v116 float32 = (3.14159 + 3.14159)
                    if (false || true) {
                        var v117 int32 = -305
                        switch ("a" + (v62 + (v62 + v62))) {
                        case "\n\t":
                            v118 := GenS18{F0: true, F1: (v1 / ((v1) | 1))}
                            _ = v118
                        case "a":
                            fmt.Println(fi114)
                        default:
                            _ = len(v4)
                        }
                        _ = v117
                    } else {
                        for fi119 := 0; fi119 < 15; fi119++ {
                            v61 -= v61
                            _ = fi119
                        }
                    }
                    _ = len(v3)
                    _ = v116
                } else {
                    var v120 int = fi114
                    _ = v120
                }
                _ = fi114
                _ = v115
            }
            _ = v113
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v2
    _ = v3
    _ = v4
    _ = v54
    _ = v61
    _ = v62
    _ = v99
}
