package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint; F1 string }
type GenS6 struct { F0 int16; F1 string }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int; F1 uint }
type GenS10 struct { F0 uint32; F1 uint32; F2 uint }
type GenS11 struct { F0 uint; F1 int; F2 uint64 }
type GenS12 struct { F0 string; F1 int32 }
func genFn0(p0 uint8) int16 {
    var v6 uint16 = 42
    var v7 int = 50
    v8 := GenS1{F0: (p0 << ((878) & 63)), F1: (v7 % (((v7 + (^(v7)))) | 1))}
    v9 := []int{}
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 0
}
func genFn1(p0 bool) uint {
    if ((!false) && true) {
        fmt.Println(p0)
    }
    p0 = (!(1 == 2147483647))
    for fi10 := 0; fi10 < 7; fi10++ {
        for fi11 := 0; fi11 < 17; fi11++ {
            p0 = p0
            v12 := []int{}
            _ = fi11
            _ = v12
        }
        p0 = (p0 || (-32768 == 35))
        _ = fi10
    }
    _ = p0
    return 65535
}

func main() {
    v1 := []int{}
    var v2 string = "\n\t"
    var v3 int16 = 72
    var v4 string = v2
    var v5 uint8 = 42
    v4 = "a"
    switch "foo bar" {
    case "a":
        v2 = v4
    }
    fmt.Println(v5, v4, v2)
    for fi13 := 0; fi13 < 13; fi13++ {
        for fi14 := 0; fi14 < 6; fi14++ {
            if len(v1) > 0 {
                v1[0] = (fi13 ^ -535)
            }
            var v15 int8 = 0
            v16 := GenS2{F0: fi14, F1: 186}
            _ = fi14
            _ = v15
            _ = v16
        }
        v4 = ((v4 + (("" + ((v4 + "foo bar") + v4)) + v2)) + (v4 + "a"))
        _ = fi13
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        var v18 float64 = 1.5
        var v19 int8 = -80
        _ = fi17
        _ = v18
        _ = v19
    }
    _ = len(v1)
    for fi20 := 0; fi20 < 18; fi20++ {
        var v21 int64 = -1
        for fi22 := 0; fi22 < 5; fi22++ {
            fmt.Println(fi20)
            _ = len(v1)
            v23 := GenS3{F0: 332, F1: -2147483648}
            _ = fi22
            _ = v23
        }
        switch 1 {
        case 9223372036854775807:
            switch (-(fi20)) {
            case -520:
                fmt.Println(v5)
            default:
                for fi24 := 0; fi24 < 4; fi24++ {
                    switch "hello" {
                    case "":
                        var v25 uint32 = 0
                        _ = v25
                    }
                    v5 += v5
                    var v26 int8 = -1
                    _ = fi24
                    _ = v26
                }
            }
        case -1:
            fmt.Println(v5, v4)
        default:
            switch (fi20 - fi20) {
            case -2147483648:
                for fi27 := 0; fi27 < 9; fi27++ {
                    switch v4 {
                    case "hello":
                        var v28 float32 = -2.25
                        _ = v28
                    case "":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    var v29 int64 = (v21 % ((-1) | 1))
                    _ = len(v1)
                    var v30 uint64 = 2
                    var v31 bool = (false && ((v5 - 127) > (v5 << ((65535) & 63))))
                    var v32 float64 = 0.1
                    fmt.Println(v4, v21, v32)
                    _ = fi27
                    _ = v29
                    _ = v30
                    _ = v31
                    _ = v32
                }
            case 32768:
                if false {
                    v33 := GenS3{F0: 644, F1: 42}
                    _ = v33
                }
            }
        }
        _ = fi20
        _ = v21
    }
    v5 = (v5 >> ((2) & 63))
    if (0 < 65535) {
        v34 := GenS3{F0: 42, F1: -2147483648}
        _ = v34
    } else {
        v1 = append(v1, 32768)
        if false {
            if true {
                for fi35 := 0; fi35 < 16; fi35++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v3 += genFn0(2)
                    fmt.Println(v2)
                    _ = fi35
                }
            }
            var v36 uint32 = 1
            _ = v36
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            switch 580 {
            case 852:
                for fi37 := 0; fi37 < 17; fi37++ {
                    if true {
                        v3 = (v3 ^ 0)
                        var v38 int32 = 1
                        v1 = append(v1, fi37)
                        _ = v38
                    } else {
                        var v39 int32 = 42
                        v40 := GenS1{F0: 0, F1: 32767}
                        _ = v39
                        _ = v40
                    }
                    _ = fi37
                }
            default:
                if false {
                    for fi41 := 0; fi41 < 19; fi41++ {
                        switch (fi41 + fi41) {
                        case -483:
                            var v42 int8 = 0
                            _ = v42
                        case -9223372036854775808:
                            for fi43 := 0; fi43 < 1; fi43++ {
                                v44 := GenS3{F0: 1, F1: -2147483648}
                                _ = len(v1)
                                fmt.Println(v5, fi43, fi43)
                                _ = fi43
                                _ = v44
                            }
                        case 128:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        var v45 int16 = 1000
                        _ = fi41
                        _ = v45
                    }
                    if (!(992 < 18446744073709551615)) {
                        if false {
                            var v46 float64 = float64((((1e10 / (3.14159 / 1.5)) + (0.1 / float32(-2.25))) - (1.5 / 0.1)))
                            _ = v46
                        }
                        if true {
                            var v47 string = v4
                            switch v47 {
                            case "\n\t":
                                var v48 int = 847
                                _ = v48
                            case "hello":
                                if (v5 >= (v5 + v5)) {
                                    v4 = ("foo bar" + v4)
                                }
                            default:
                                if true {
                                    var v49 int64 = 42
                                    v2 = "\n\t"
                                    fmt.Println(v2, v5, v49)
                                    _ = v49
                                } else {
                                    for fi50 := 0; fi50 < 2; fi50++ {
                                        v3 -= (v3 % ((genFn0(127)) | 1))
                                        var v51 bool = false
                                        if len(v1) > 0 {
                                            v1[0] = (fi50 % ((fi50) | 1))
                                        }
                                        _ = fi50
                                        _ = v51
                                    }
                                }
                            }
                            switch 128 {
                            case -1:
                                var v52 int64 = -13
                                _ = v52
                            case -364:
                                switch 872 {
                                case 0:
                                    for fi53 := 0; fi53 < 20; fi53++ {
                                        if (377 < 426) {
                                            v47 = "a"
                                        } else {
                                            var v54 uint32 = 2147483648
                                            if (float32(0.0) == 1e10) {
                                                v55 := GenS4{F0: 42}
                                                _ = v55
                                            } else {
                                                v56 := GenS5{F0: uint(v54), F1: v47}
                                                var v57 float64 = 0.0
                                                _ = v56
                                                _ = v57
                                            }
                                            v1 = append(v1, fi53)
                                            _ = v54
                                        }
                                        if (2147483647 >= 961) {
                                            if false {
                                                if (true || false) {
                                                    for fi58 := 0; fi58 < 16; fi58++ {
                                                        switch 1 {
                                                        case 42:
                                                            switch v4 {
                                                            case "":
                                                                v1 = append(v1, fi58)
                                                            case "a":
                                                                var v59 uint = 65535
                                                                _ = v59
                                                            default:
                                                                if true {
                                                                    var v60 uint16 = 0
                                                                    v3 = (-(v3))
                                                                    var v61 int32 = -2147483648
                                                                    var v62 int64 = 1
                                                                    _ = v60
                                                                    _ = v61
                                                                    _ = v62
                                                                }
                                                            }
                                                        case 32767:
                                                            var v63 float32 = float32((3.14159 / 0.1))
                                                            _ = v63
                                                        case 4294967295:
                                                            v5 = v5
                                                        }
                                                        fmt.Println(v2, v4, v3)
                                                        v5 = v5
                                                        var v64 float32 = (1e10 * ((0.0 - ((0.0 - (-2.25 / 3.14159)) - 1.5)) / float32(-2.25)))
                                                        _ = fi58
                                                        _ = v64
                                                    }
                                                    for fi65 := 0; fi65 < 11; fi65++ {
                                                        fmt.Println(fi53, v3)
                                                        _ = fi65
                                                    }
                                                    v66 := GenS6{F0: -221, F1: ((v47 + (v2 + v47)) + "hello")}
                                                    _ = v66
                                                } else {
                                                    var v67 int16 = (v3 / ((genFn0(v5)) | 1))
                                                    _ = v67
                                                }
                                                _ = len(v1)
                                            }
                                            var v68 float32 = (3.14159 + 1e10)
                                            v68 += v68
                                            _ = v68
                                        } else {
                                            v5 -= 127
                                            if true {
                                                if false {
                                                    var v69 int8 = -128
                                                    switch fi53 {
                                                    case 9223372036854775807:
                                                        v70 := GenS3{F0: 503, F1: 21}
                                                        _ = v70
                                                    }
                                                    v71 := GenS2{F0: (^(fi53)), F1: 90}
                                                    _ = v69
                                                    _ = v71
                                                }
                                                var v72 int16 = int16(fi53)
                                                _ = v72
                                            }
                                            v3 = 1000
                                        }
                                        _ = fi53
                                    }
                                }
                            default:
                                var v73 uint32 = 2147483648
                                _ = v73
                            }
                            _ = v47
                        }
                        if ((-124 < 42) && (65535 == 2147483647)) {
                            fmt.Println(v3, v5)
                            if len(v1) > 0 {
                                v1[0] = -9223372036854775808
                            }
                        }
                    } else {
                        var v74 int16 = (v3 - 0)
                        var v75 float64 = 0.0
                        _ = v74
                        _ = v75
                    }
                } else {
                    var v76 float64 = float64(float32((-2.25 * (1.5 * -2.25))))
                    for fi77 := 0; fi77 < 18; fi77++ {
                        switch "" {
                        case "foo bar":
                            var v78 int8 = 100
                            _ = v78
                        case "\n\t":
                            v3 = 1000
                        default:
                            fmt.Println(v5, v4, v76)
                        }
                        switch "foo bar" {
                        case "hello":
                            v2 = ((v2 + v2) + ("a" + v4))
                        case "\n\t":
                            var v79 uint = 255
                            _ = v79
                        default:
                            for fi80 := 0; fi80 < 5; fi80++ {
                                v81 := GenS3{F0: 175, F1: 0}
                                _ = fi80
                                _ = v81
                            }
                        }
                        var v82 int32 = 194
                        _ = fi77
                        _ = v82
                    }
                    v83 := GenS4{F0: 42}
                    _ = v76
                    _ = v83
                }
            }
            switch 42 {
            case 42:
                v3 -= int16(v5)
            }
        }
        if (0.1 >= float32(1.5)) {
            for fi84 := 0; fi84 < 4; fi84++ {
                var v85 uint8 = v5
                switch ((v2 + "\n\t") + ((v2 + (("\n\t" + "a") + "foo bar")) + v2)) {
                case "foo bar":
                    fmt.Println(v5)
                default:
                    if true {
                        v4 = "hello"
                    }
                }
                fmt.Println(v2, fi84, fi84)
                _ = fi84
                _ = v85
            }
            if true {
                var v86 uint8 = (v5 ^ 0)
                v2 = v2
                var v87 uint16 = 2
                _ = v86
                _ = v87
            }
        } else {
            v2 = (v2 + "hello")
            switch -2147483648 {
            case -291:
                v3 = -32768
            default:
                if false {
                    for fi88 := 0; fi88 < 13; fi88++ {
                        if (v5 <= (v5 * v5)) {
                            v4 = v2
                            var v89 int64 = 9223372036854775807
                            _ = v89
                        }
                        _ = fi88
                    }
                    var v90 float32 = 3.14159
                    var v91 float32 = (v90 / -2.25)
                    _ = v90
                    _ = v91
                } else {
                    var v92 int = 42
                    v2 = v4
                    _ = v92
                }
            }
            for fi93 := 0; fi93 < 8; fi93++ {
                v4 = ""
                if (!false) {
                    if (172 >= 255) {
                        var v94 float32 = ((((0.0 / 3.14159) / -2.25) / 1.0) + (1.5 + 0.0))
                        fmt.Println(v94)
                        _ = v94
                    }
                    fmt.Println(fi93)
                    switch v4 {
                    case "hello":
                        if true {
                            switch fi93 {
                            case 4294967295:
                                var v95 uint8 = v5
                                _ = v95
                            case -54:
                                v5 = v5
                            case 42:
                                var v96 string = v4
                                _ = v96
                            default:
                                var v97 uint16 = 65535
                                _ = v97
                            }
                            switch v4 {
                            case "a":
                                v4 = (v2 + ("foo bar" + "a"))
                            case "foo bar":
                                var v98 uint32 = 65535
                                _ = v98
                            default:
                                v99 := GenS7{F0: 100}
                                _ = v99
                            }
                            if true {
                                fmt.Println(v4, v5, v2)
                                v5 = (v5 - v5)
                                var v100 int32 = 1
                                _ = v100
                            }
                        } else {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                    default:
                        if (((v2 + v4) <= v4) || (1 <= -1)) {
                            var v101 float32 = ((1.5 * float32(0.0)) - (0.0 - 0.1))
                            _ = v101
                        } else {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v5 -= v5
                            fmt.Println(fi93)
                        }
                    }
                } else {
                    if (!true) {
                        v1 = append(v1, 73)
                        for fi102 := 0; fi102 < 15; fi102++ {
                            var v103 bool = true
                            _ = fi102
                            _ = v103
                        }
                        v3 = (v3 >> ((1) & 63))
                    } else {
                        var v104 uint16 = 32767
                        var v105 float64 = 3.14159
                        _ = v104
                        _ = v105
                    }
                    if (-99 != 55) {
                        fmt.Println(fi93, v4)
                        v3 = (v3 | v3)
                        var v106 int = fi93
                        _ = v106
                    }
                    v2 = v4
                }
                switch v4 {
                case "hello":
                    switch (fi93 * (fi93 * -934)) {
                    case 4294967295:
                        for fi107 := 0; fi107 < 9; fi107++ {
                            if true {
                                var v108 string = (v2 + "")
                                v109 := GenS4{F0: 0}
                                _ = v108
                                _ = v109
                            } else {
                                v5 += 60
                                v110 := GenS8{F0: 186}
                                _ = v110
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            var v111 int64 = -950
                            _ = fi107
                            _ = v111
                        }
                    case 1:
                        v5 -= 208
                    default:
                        v5 = v5
                    }
                case "":
                    v112 := GenS9{F0: fi93, F1: 18446744073709551615}
                    _ = v112
                case "\n\t":
                    var v113 uint = genFn1(true)
                    _ = v113
                default:
                    v5 -= (v5 << ((genFn1((((68 == 1) || true) || (false && true)))) & 63))
                }
                _ = fi93
            }
        }
    }
    fmt.Println(v4, v3, v3)
    v114 := GenS5{F0: 163, F1: "a"}
    fmt.Println(v4, v5, v5)
    var v115 bool = false
    switch 2147483647 {
    case 1:
        v116 := GenS7{F0: -65}
        _ = v116
    case -2147483648:
        v5 += 51
    case 42:
        if v115 {
            for fi117 := 0; fi117 < 3; fi117++ {
                switch fi117 {
                case 32767:
                    switch (v4 + "foo bar") {
                    case "a":
                        if v115 {
                            v2 = v2
                            var v118 int16 = (v3 - (v3 << ((255) & 63)))
                            for fi119 := 0; fi119 < 20; fi119++ {
                                var v120 uint8 = v5
                                v121 := GenS10{F0: 275, F1: 65535, F2: 309}
                                fmt.Println(v120)
                                _ = fi119
                                _ = v120
                                _ = v121
                            }
                            _ = v118
                        }
                    case "foo bar":
                        v1 = append(v1, (fi117 << ((0) & 63)))
                    }
                }
                _ = fi117
            }
            v115 = (2147483647 != 42)
        } else {
            if (!(0.0 >= 1e10)) {
                for fi122 := 0; fi122 < 5; fi122++ {
                    if v115 {
                        v123 := GenS11{F0: 4294967295, F1: -672, F2: 4294967295}
                        _ = v123
                    }
                    var v124 int16 = genFn0((v5 / (((v5 * v5)) | 1)))
                    _ = fi122
                    _ = v124
                }
                if true {
                    v125 := GenS12{F0: v4, F1: 1}
                    _ = v125
                } else {
                    v2 = ("\n\t" + "a")
                    v1 = append(v1, -283)
                    if len(v1) > 0 {
                        v1[0] = 0
                    }
                }
            }
            v3 += 623
        }
    default:
        var v126 bool = ((1 >= 128) || v115)
        _ = v126
    }
    switch 806 {
    case -2147483648:
        var v127 uint32 = 0
        _ = v127
    case 32767:
        for fi128 := 0; fi128 < 7; fi128++ {
            for fi129 := 0; fi129 < 16; fi129++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v130 := GenS4{F0: 2}
                if (true && true) {
                    fmt.Println(v4)
                    switch v2 {
                    case "\n\t":
                        v115 = v115
                    default:
                        var v131 uint8 = v5
                        _ = v131
                    }
                    var v132 int32 = -940
                    _ = v132
                }
                _ = fi129
                _ = v130
            }
            _ = fi128
        }
    case -537:
        v133 := GenS7{F0: -128}
        _ = v133
    }
    for fi134 := 0; fi134 < 14; fi134++ {
        v135 := GenS12{F0: ("a" + v4), F1: 0}
        var v136 float32 = ((-2.25 - 3.14159) - float32((1e10 + (0.0 * float64(1e10)))))
        _ = fi134
        _ = v135
        _ = v136
    }
    var v137 uint = 18446744073709551615
    var v138 uint8 = (v5 / (((v5 % ((v5) | 1))) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v115:", v115)
    fmt.Println("v137:", v137)
    fmt.Println("v138:", v138)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v114
    _ = v115
    _ = v137
    _ = v138
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
