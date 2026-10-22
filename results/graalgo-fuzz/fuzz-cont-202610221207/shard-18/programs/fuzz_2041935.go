package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32; F1 int64; F2 uint32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int8; F1 uint; F2 string }
type GenS7 struct { F0 int; F1 int16 }
type GenS8 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint16; F1 string }
type GenS11 struct { F0 bool; F1 bool }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 int16; F1 uint32 }
type GenS16 struct { F0 int16 }
type GenS17 struct { F0 string }
type GenS18 struct { F0 int8 }
func genFn0(p0 uint16, p1 uint16) int32 {
    v5 := []int{127, 2147483647, -716}
    v6 := GenS1{F0: 0, F1: true}
    var v7 int = 1
    v8 := GenS1{F0: -128, F1: ("foo bar" < ("a" + ("\n\t" + "hello")))}
    v5 = append(v5, -449)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return -1
}
func genFn1(p0 int64) int64 {
    var v9 float32 = 0.0
    for fi10 := 0; fi10 < 5; fi10++ {
        v11 := GenS2{F0: 2, F1: 832}
        var v12 int32 = 0
        var v13 int = (^(fi10))
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        p0 += p0
        var v15 bool = (42 < 96)
        for fi16 := 0; fi16 < 17; fi16++ {
            if (!(fi14 == (fi14 + (fi16 / ((fi16) | 1))))) {
                v17 := GenS3{F0: 999}
                var v18 uint32 = 4294967295
                v19 := []int{fi16}
                _ = v17
                _ = v18
                _ = v19
            }
            v20 := GenS4{F0: 2147483648, F1: (p0 + p0), F2: 42}
            _ = fi16
            _ = v20
        }
        _ = fi14
        _ = v15
    }
    var v21 uint = 255
    _ = p0
    _ = v21
    _ = v9
    return p0
}

func main() {
    v1 := []uint32{2147483647, 42}
    v2 := []uint32{428, 378, 2147483647, 704}
    var v3 uint8 = 255
    v4 := []string{"hello", "hello"}
    v3 -= 42
    if (4294967295 > 746) {
        fmt.Println(v3)
    }
    var v22 int32 = 2147483647
    switch "\n\t" {
    case "":
        for fi23 := 0; fi23 < 12; fi23++ {
            var v24 uint32 = 2147483647
            _ = fi23
            _ = v24
        }
    case "\n\t":
        fmt.Println(v3, v3)
    case "hello":
        if false {
            for fi25 := 0; fi25 < 5; fi25++ {
                v3 -= 255
                _ = fi25
            }
        } else {
            v3 = (v3 | (v3 * v3))
            var v26 int16 = 970
            _ = v26
        }
    default:
        var v27 string = ("foo bar" + "hello")
        _ = v27
    }
    switch "hello" {
    case "\n\t":
        if (true && true) {
            v28 := GenS4{F0: 0, F1: -1, F2: 4294967295}
            switch ("\n\t" + "a") {
            case "foo bar":
                if (false && (v3 != v3)) {
                    var v29 int = -591
                    switch "a" {
                    case "\n\t":
                        v30 := GenS4{F0: 2, F1: -632, F2: 670}
                        _ = v30
                    case "":
                        for fi31 := 0; fi31 < 19; fi31++ {
                            if len(v1) > 0 {
                                v1[0] = 2
                            }
                            _ = fi31
                        }
                    case "a":
                        if true {
                            if true {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                if true {
                                    if true {
                                        if true {
                                            if true {
                                                var v32 float64 = ((3.14159 + (-2.25 + (0.1 * 1e10))) - 0.0)
                                                if (!(!false)) {
                                                    v33 := GenS3{F0: 2}
                                                    _ = v33
                                                } else {
                                                    if ((1.5 >= v32) && (770 == 1)) {
                                                        v29 = 9223372036854775807
                                                    }
                                                }
                                                var v34 float32 = (0.1 - float32(v32))
                                                _ = v32
                                                _ = v34
                                            }
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            for fi35 := 0; fi35 < 2; fi35++ {
                                                switch (fi35 * -534) {
                                                case 4294967295:
                                                    for fi36 := 0; fi36 < 4; fi36++ {
                                                        if (-2.25 <= 1e10) {
                                                            if false {
                                                                var v37 int64 = -9223372036854775808
                                                                v22 -= (v22 << ((65535) & 63))
                                                                if false {
                                                                    var v38 uint32 = 245
                                                                    var v39 string = (("a" + "a") + "a")
                                                                    _ = v38
                                                                    _ = v39
                                                                } else {
                                                                    var v40 uint8 = v3
                                                                    fmt.Println(fi36, v29, v3)
                                                                    v41 := GenS5{F0: v37}
                                                                    _ = v40
                                                                    _ = v41
                                                                }
                                                                v4 = append(v4, ("\n\t" + "foo bar"))
                                                                _ = v37
                                                            }
                                                            switch 32767 {
                                                            case 128:
                                                                v2 = append(v2, 2147483647)
                                                            case -1:
                                                                var v42 bool = false
                                                                _ = v42
                                                            default:
                                                                if (-32768 == 841) {
                                                                    fmt.Println(v22)
                                                                    v43 := GenS1{F0: -128, F1: false}
                                                                    switch "foo bar" {
                                                                    case "\n\t":
                                                                        switch (v29 / ((2147483647) | 1)) {
                                                                        case 2147483647:
                                                                            var v44 uint16 = 0
                                                                            _ = v44
                                                                        case 4294967295:
                                                                            if len(v1) > 0 {
                                                                                v1[0] = 65535
                                                                            }
                                                                        case 1:
                                                                            var v45 float64 = 0.1
                                                                            _ = v45
                                                                        default:
                                                                            if (0 == 281) {
                                                                                fmt.Println(fi35)
                                                                                switch (v29 >> ((4294967295) & 63)) {
                                                                                case 9223372036854775807:
                                                                                    if (148 > (v3 & (^(v3)))) {
                                                                                        switch (v29 & (fi35 % ((128) | 1))) {
                                                                                        case 2147483647:
                                                                                            v2 = append(v2, 2)
                                                                                        case 42:
                                                                                            switch fi35 {
                                                                                            case 32767:
                                                                                                if false {
                                                                                                    switch (^(fi35)) {
                                                                                                    case 128:
                                                                                                        var v46 bool = true
                                                                                                        _ = v46
                                                                                                    case 1:
                                                                                                        v29 -= -675
                                                                                                    default:
                                                                                                        v2 = append(v2, 0)
                                                                                                    }
                                                                                                    v3 = v3
                                                                                                }
                                                                                            default:
                                                                                                var v47 float32 = 0.1
                                                                                                _ = v47
                                                                                            }
                                                                                            v48 := GenS3{F0: 42}
                                                                                            v22 += (v22 & v22)
                                                                                            _ = v48
                                                                                        case 400:
                                                                                            if len(v2) >= 2 {
                                                                                                v2 = v2[1:len(v2)]
                                                                                            }
                                                                                        }
                                                                                        var v49 uint32 = 1
                                                                                        _ = v49
                                                                                    } else {
                                                                                        v3 += v3
                                                                                        var v50 uint = 706
                                                                                        _ = v50
                                                                                    }
                                                                                case 32767:
                                                                                    var v51 int = fi36
                                                                                    _ = v51
                                                                                default:
                                                                                    var v52 float64 = 1e10
                                                                                    _ = v52
                                                                                }
                                                                            } else {
                                                                                v22 = 1
                                                                                v29 += 350
                                                                                v53 := GenS1{F0: 100, F1: ((v29 << ((42) & 63)) <= 2147483647)}
                                                                                _ = v53
                                                                            }
                                                                        }
                                                                    case "a":
                                                                        _ = len(v2)
                                                                    }
                                                                    _ = v43
                                                                }
                                                            }
                                                        } else {
                                                            var v54 int16 = 0
                                                            var v55 uint64 = 313
                                                            _ = v54
                                                            _ = v55
                                                        }
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        switch (^(fi36)) {
                                                        case 685:
                                                            if true {
                                                                fmt.Println(fi36)
                                                                var v56 uint16 = 2
                                                                v57 := GenS3{F0: 2}
                                                                _ = v56
                                                                _ = v57
                                                            } else {
                                                                var v58 uint = 4294967295
                                                                _ = v58
                                                            }
                                                        case 4294967295:
                                                            var v59 uint = 1
                                                            _ = v59
                                                        default:
                                                            v2 = append(v2, 42)
                                                        }
                                                        _ = fi36
                                                    }
                                                case 583:
                                                    v29 = v29
                                                case 128:
                                                    var v60 uint8 = 2
                                                    _ = v60
                                                default:
                                                    v61 := GenS6{F0: -57, F1: 1, F2: "a"}
                                                    _ = v61
                                                }
                                                v22 -= (v22 >> ((18446744073709551615) & 63))
                                                _ = fi35
                                            }
                                        } else {
                                            if false {
                                                v3 = (v3 - 255)
                                                v62 := GenS6{F0: 100, F1: 65535, F2: "foo bar"}
                                                fmt.Println(v3, v22, v22)
                                                _ = v62
                                            } else {
                                                switch 0 {
                                                case 32767:
                                                    var v63 int32 = (v22 >> ((587) & 63))
                                                    _ = v63
                                                case 2147483647:
                                                    var v64 int = v29
                                                    _ = v64
                                                case 127:
                                                    v65 := GenS5{F0: -500}
                                                    _ = v65
                                                }
                                                if (-921 != 1) {
                                                    _ = len(v4)
                                                    v3 -= 0
                                                    for fi66 := 0; fi66 < 7; fi66++ {
                                                        if false {
                                                            if false {
                                                                if (42 <= 516) {
                                                                    var v67 string = "\n\t"
                                                                    _ = v67
                                                                }
                                                                fmt.Println(fi66, v3, v22)
                                                            } else {
                                                                for fi68 := 0; fi68 < 11; fi68++ {
                                                                    v29 += (fi68 - -2147483648)
                                                                    _ = fi68
                                                                }
                                                                v29 -= (fi66 >> ((4294967295) & 63))
                                                            }
                                                            fmt.Println(fi66)
                                                            v29 -= fi66
                                                        }
                                                        v69 := GenS7{F0: (-(v29)), F1: -1}
                                                        var v70 uint = 243
                                                        _ = fi66
                                                        _ = v69
                                                        _ = v70
                                                    }
                                                } else {
                                                    v71 := GenS8{F0: 1000, F1: 42, F2: 255}
                                                    switch "hello" {
                                                    case "foo bar":
                                                        fmt.Println(v29, v22, v3)
                                                    default:
                                                        v3 = (v3 ^ 42)
                                                    }
                                                    _ = v71
                                                }
                                            }
                                            if true {
                                                switch (v29 - v29) {
                                                case 343:
                                                    v72 := GenS9{F0: "foo bar"}
                                                    _ = v72
                                                }
                                            }
                                            v73 := GenS10{F0: 65535, F1: ("" + "\n\t")}
                                            _ = v73
                                        }
                                        for fi74 := 0; fi74 < 19; fi74++ {
                                            _ = len(v2)
                                            _ = fi74
                                        }
                                        if (!(!true)) {
                                            if false {
                                                if (!(394 != -32768)) {
                                                    if true {
                                                        var v75 int16 = 1
                                                        _ = v75
                                                    }
                                                }
                                            }
                                            fmt.Println(v22, v22, v22)
                                            var v76 uint16 = 65535
                                            _ = v76
                                        }
                                    } else {
                                        if false {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            var v77 uint16 = 42
                                            _ = v77
                                        }
                                    }
                                    var v78 float32 = (0.0 * ((0.1 - (1e10 / (1e10 / 1e10))) * float32(0.1)))
                                    _ = v78
                                }
                            } else {
                                var v79 int64 = -756
                                for fi80 := 0; fi80 < 5; fi80++ {
                                    for fi81 := 0; fi81 < 8; fi81++ {
                                        switch (fi81 / ((fi81) | 1)) {
                                        case 32767:
                                            var v82 string = "\n\t"
                                            _ = v82
                                        case -536:
                                            if (2147483648 >= 362) {
                                                _ = len(v4)
                                            }
                                        case 42:
                                            var v83 uint = 626
                                            _ = v83
                                        default:
                                            switch fi81 {
                                            case -659:
                                                v84 := GenS1{F0: 127, F1: true}
                                                _ = v84
                                            case 9223372036854775807:
                                                var v85 int = (fi80 | fi81)
                                                _ = v85
                                            }
                                        }
                                        v86 := GenS8{F0: 1000, F1: 1, F2: 65535}
                                        _ = fi81
                                        _ = v86
                                    }
                                    v2 = append(v2, 773)
                                    _ = fi80
                                }
                                switch (v29 / ((32767) | 1)) {
                                case 127:
                                    v87 := GenS3{F0: 65535}
                                    _ = v87
                                default:
                                    v29 -= (^(v29))
                                }
                                _ = v79
                            }
                        }
                    default:
                        v22 -= int32(v29)
                    }
                    var v88 uint = 968
                    _ = v29
                    _ = v88
                }
            case "":
                switch -9223372036854775808 {
                case 32767:
                    fmt.Println(v3, v3)
                case -217:
                    v89 := GenS11{F0: false, F1: ((!false) || (2147483647 < 2))}
                    _ = v89
                }
            default:
                var v90 int64 = 1
                _ = v90
            }
            _ = v28
        } else {
            switch -9223372036854775808 {
            case -562:
                v3 = 255
            case 128:
                for fi91 := 0; fi91 < 5; fi91++ {
                    var v92 int32 = (v22 ^ (^(v22)))
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi91
                    _ = v92
                }
            }
            v93 := GenS4{F0: 1, F1: 42, F2: 1}
            v94 := GenS6{F0: 1, F1: 65535, F2: "hello"}
            _ = v93
            _ = v94
        }
    default:
        var v95 float64 = -2.25
        _ = v95
    }
    var v96 uint32 = 631
    if true {
        if true {
            if true {
                switch "hello" {
                case "a":
                    if (100 == 42) {
                        v97 := GenS3{F0: 0}
                        v98 := GenS12{F0: v22}
                        _ = v97
                        _ = v98
                    }
                case "":
                    for fi99 := 0; fi99 < 1; fi99++ {
                        v100 := GenS13{F0: v3}
                        _ = fi99
                        _ = v100
                    }
                default:
                    for fi101 := 0; fi101 < 4; fi101++ {
                        switch "a" {
                        case "foo bar":
                            switch "a" {
                            case "":
                                var v102 uint8 = (^(v3))
                                _ = v102
                            case "hello":
                                v103 := GenS14{F0: 9223372036854775807}
                                _ = v103
                            case "a":
                                var v104 uint8 = 2
                                _ = v104
                            default:
                                var v105 uint8 = (v3 + 2)
                                _ = v105
                            }
                        case "hello":
                            var v106 float32 = float32(0.1)
                            _ = v106
                        }
                        var v107 uint = 18446744073709551615
                        for fi108 := 0; fi108 < 4; fi108++ {
                            v109 := GenS15{F0: 0, F1: uint32(v3)}
                            fmt.Println(fi108)
                            var v110 uint32 = uint32(v107)
                            _ = fi108
                            _ = v109
                            _ = v110
                        }
                        _ = fi101
                        _ = v107
                    }
                }
                v3 = v3
            } else {
                switch 32767 {
                case 127:
                    fmt.Println(v96, v3, v22)
                case 258:
                    fmt.Println(v3, v3)
                }
                v3 = (v3 ^ (v3 - v3))
            }
            var v111 int = 1
            v3 = (v3 * (v3 & (v3 % (((v3 | v3)) | 1))))
            _ = v111
        } else {
            v112 := GenS3{F0: (v96 ^ v96)}
            _ = v112
        }
    }
    for fi113 := 0; fi113 < 17; fi113++ {
        fmt.Println(fi113, v3, v96)
        v114 := GenS4{F0: (v96 % ((v96) | 1)), F1: 9223372036854775807, F2: 2147483647}
        _ = fi113
        _ = v114
    }
    for fi115 := 0; fi115 < 19; fi115++ {
        v116 := GenS16{F0: -849}
        v3 = (v3 - v3)
        switch ("foo bar" + (("foo bar" + "a") + "hello")) {
        case "foo bar":
            v22 = v22
        case "\n\t":
            v22 += (v22 << ((4294967295) & 63))
        case "":
            v117 := GenS13{F0: 128}
            _ = v117
        default:
            _ = len(v1)
        }
        _ = fi115
        _ = v116
    }
    v118 := GenS17{F0: ("foo bar" + "")}
    var v119 int8 = 1
    v3 += uint8(v96)
    if (1 > 9223372036854775807) {
        if (((!false) && true) || false) {
            v96 = (v96 >> ((65535) & 63))
            switch "" {
            case "hello":
                v120 := GenS18{F0: 50}
                _ = v120
            case "foo bar":
                var v121 uint = 958
                _ = v121
            case "\n\t":
                v22 = (v22 >> ((65535) & 63))
            default:
                for fi122 := 0; fi122 < 2; fi122++ {
                    if (false || (1.5 > ((1e10 * float64((((1e10 + (1e10 * float32(0.1))) + (1.5 - float32(((float64(float32(float64(1.5))) / 1.0) + 1.5)))) + 0.1))) + 0.0))) {
                        v123 := GenS14{F0: genFn1(-9223372036854775808)}
                        v119 = (-(v119))
                        for fi124 := 0; fi124 < 14; fi124++ {
                            switch (^(fi122)) {
                            case -366:
                                if (-1 >= -32768) {
                                    switch fi122 {
                                    case -9223372036854775808:
                                        switch 293 {
                                        case 32767:
                                            var v125 int = (-(fi124))
                                            _ = v125
                                        case 747:
                                            var v126 string = ""
                                            _ = v126
                                        }
                                    }
                                    if false {
                                        v3 += 76
                                        v127 := GenS15{F0: int16(v22), F1: 4294967295}
                                        _ = v127
                                    } else {
                                        v2 = append(v2, (v96 + v96))
                                        v119 -= v119
                                    }
                                    v96 -= v96
                                }
                            case 32767:
                                var v128 uint64 = 0
                                _ = v128
                            case -1:
                                fmt.Println(v96, fi122)
                            }
                            _ = fi124
                        }
                        _ = v123
                    }
                    v119 -= 42
                    var v129 uint = 254
                    _ = fi122
                    _ = v129
                }
            }
        }
        var v130 uint64 = 0
        _ = v130
    }
    if ((v3 % (((v3 & v3)) | 1)) < (v3 << ((uint(v3)) & 63))) {
        v96 += (v96 / ((4294967295) | 1))
        for fi131 := 0; fi131 < 3; fi131++ {
            _ = len(v1)
            _ = fi131
        }
    }
    for fi132 := 0; fi132 < 19; fi132++ {
        v96 -= (v96 - (v96 << ((1) & 63)))
        if len(v2) > 0 {
            v2[0] = 42
        }
        _ = fi132
    }
    v119 = (v119 << ((255) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v119:", v119)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v118
    _ = v119
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v96
}
