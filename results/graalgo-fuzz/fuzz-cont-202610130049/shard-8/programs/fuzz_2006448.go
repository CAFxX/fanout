package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 int64; F2 int }
type GenS5 struct { F0 uint8; F1 uint16; F2 int16 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int; F1 string }
type GenS8 struct { F0 int32; F1 int; F2 int }
type GenS9 struct { F0 int64; F1 uint; F2 int }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint16; F1 uint8 }
type GenS12 struct { F0 uint; F1 int; F2 string }
func genFn0() string {
    v6 := GenS1{F0: 0, F1: 255, F2: 786}
    v7 := GenS1{F0: 372, F1: 255, F2: 4294967295}
    var v8 int = 32768
    var v9 string = "foo bar"
    var v10 int64 = -412
    _ = v10
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return ("hello" + (v9 + "foo bar"))
}
func genFn1(p0 int, p1 string) int64 {
    if false {
        for fi11 := 0; fi11 < 11; fi11++ {
            v12 := []int{(p0 * (p0 & fi11)), (p0 >> ((65535) & 63)), p0, 1}
            if true {
                for fi13 := 0; fi13 < 6; fi13++ {
                    if len(v12) > 0 {
                        v12[0] = -2147483648
                    }
                    _ = fi13
                }
                fmt.Println(p1, p1)
            } else {
                if (2147483647 >= 42) {
                    switch ((p1 + (p1 + p1)) + p1) {
                    case "\n\t":
                        if len(v12) >= 2 {
                            v12 = v12[1:len(v12)]
                        }
                    default:
                        v14 := GenS1{F0: -1, F1: 0, F2: 438}
                        _ = v14
                    }
                    fmt.Println(fi11, fi11)
                }
                var v15 int16 = 311
                var v16 int8 = 42
                _ = v15
                _ = v16
            }
            if true {
                for fi17 := 0; fi17 < 6; fi17++ {
                    fmt.Println(fi17, fi17, fi17)
                    var v18 bool = (1e10 < 0.0)
                    v12 = append(v12, fi11)
                    v19 := GenS1{F0: 1, F1: 255, F2: 2147483647}
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                v12 = append(v12, fi11)
                if (!false) {
                    var v20 float32 = 1.5
                    switch fi11 {
                    case 1:
                        if false {
                            var v21 float32 = v20
                            _ = v21
                        }
                    }
                    _ = v20
                } else {
                    var v22 uint64 = 0
                    if true {
                        v23 := GenS1{F0: 1, F1: 2, F2: 0}
                        p1 = (((p1 + "foo bar") + p1) + p1)
                        var v24 int = (p0 << ((199) & 63))
                        _ = v23
                        _ = v24
                    }
                    p1 = p1
                    _ = v22
                }
            } else {
                var v25 uint16 = 622
                _ = v25
            }
            _ = fi11
            _ = v12
        }
    } else {
        var v26 int = 42
        switch p1 {
        case "a":
            for fi27 := 0; fi27 < 11; fi27++ {
                v28 := []int{(fi27 % ((-1) | 1)), p0}
                _ = fi27
                _ = v28
            }
        case "\n\t":
            for fi29 := 0; fi29 < 1; fi29++ {
                if false {
                    v30 := []int{fi29, p0}
                    p0 += fi29
                    _ = v30
                } else {
                    var v31 int32 = 0
                    _ = v31
                }
                if true {
                    fmt.Println(p0)
                } else {
                    if true {
                        var v32 uint16 = 804
                        fmt.Println(v26, p1)
                        p0 += (v26 | fi29)
                        _ = v32
                    }
                    p1 = (p1 + ("" + ""))
                }
                _ = fi29
            }
        default:
            var v33 int16 = -1
            _ = v33
        }
        fmt.Println(v26, p1, v26)
        _ = v26
    }
    if (-1 <= -65) {
        for fi34 := 0; fi34 < 17; fi34++ {
            fmt.Println(p0, p0)
            if true {
                switch "a" {
                case "a":
                    switch "hello" {
                    case "a":
                        v35 := GenS1{F0: 42, F1: 0, F2: 83}
                        _ = v35
                    }
                default:
                    var v36 uint32 = 384
                    _ = v36
                }
                fmt.Println(fi34, p1)
                v37 := []int{}
                _ = v37
            }
            fmt.Println(p0)
            _ = fi34
        }
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint = 65535
    v3 := []int{120}
    var v4 int8 = -42
    v5 := []uint32{2147483647}
    fmt.Println(v2, v4)
    var v38 int16 = 1000
    fmt.Println(v4)
    v39 := GenS1{F0: genFn1(-65, "hello"), F1: 42, F2: 2}
    var v40 uint = 977
    for fi41 := 0; fi41 < 3; fi41++ {
        fmt.Println(v2, v38)
        _ = fi41
    }
    if (!false) {
        var v42 float64 = 0.1
        _ = v42
    } else {
        v43 := GenS2{F0: 18446744073709551615, F1: 4294967295}
        if false {
            if len(v5) > 0 {
                v5[0] = 858
            }
            for fi44 := 0; fi44 < 13; fi44++ {
                if false {
                    for fi45 := 0; fi45 < 14; fi45++ {
                        v4 = (^(v4))
                        _ = fi45
                    }
                    fmt.Println(v40, v40)
                    v38 = 143
                } else {
                    _ = len(v1)
                }
                v40 += 18446744073709551615
                if (187 == -1) {
                    for fi46 := 0; fi46 < 3; fi46++ {
                        fmt.Println(v38)
                        if false {
                            fmt.Println(v40, v38, v4)
                        }
                        _ = fi46
                    }
                    var v47 float64 = -2.25
                    if true {
                        if true {
                            var v48 int = fi44
                            fmt.Println(fi44, v48, v48)
                            delete(v1, 2147483647)
                            _ = v48
                        }
                        var v49 float32 = ((((0.0 / 3.14159) + 0.0) * float32(0.1)) + 0.0)
                        _ = v49
                    } else {
                        for fi50 := 0; fi50 < 8; fi50++ {
                            v51 := GenS2{F0: 4294967295, F1: 103}
                            _ = fi50
                            _ = v51
                        }
                        switch (fi44 & 32768) {
                        case 32767:
                            v38 = (v38 & 1)
                        case -566:
                            if true {
                                if (!((v38 - (v38 - 1000)) <= 42)) {
                                    v38 -= v38
                                    v2 -= v2
                                } else {
                                    for fi52 := 0; fi52 < 20; fi52++ {
                                        v4 = v4
                                        v2 = v40
                                        if false {
                                            switch (("foo bar" + "foo bar") + "\n\t") {
                                            case "hello":
                                                fmt.Println(v4, v47)
                                            case "":
                                                if false {
                                                    if false {
                                                        v40 -= v40
                                                    }
                                                    fmt.Println(fi44)
                                                    fmt.Println(v2, v2, v4)
                                                } else {
                                                    var v53 float32 = (1e10 + 0.0)
                                                    var v54 float64 = (float64(((v53 / 1.5) + v53)) - v47)
                                                    _ = v53
                                                    _ = v54
                                                }
                                            case "a":
                                                v55 := GenS2{F0: 42, F1: 1}
                                                _ = v55
                                            }
                                            v56 := v1[1]
                                            v40 -= (^(v2))
                                            v57 := GenS2{F0: 45, F1: (v56 << ((42) & 63))}
                                            v40 = v40
                                            v56 = v56
                                            _ = v56
                                            _ = v57
                                        }
                                        v38 = 1
                                        var v58 uint64 = 1
                                        v59 := GenS3{F0: ((("\n\t" + "\n\t") + "a") + (("\n\t" + "") + "\n\t"))}
                                        _ = fi52
                                        _ = v58
                                        _ = v59
                                    }
                                }
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            } else {
                                if false {
                                    for fi60 := 0; fi60 < 15; fi60++ {
                                        switch "hello" {
                                        case "foo bar":
                                            if false {
                                                v61 := GenS4{F0: true, F1: -1, F2: (fi60 * (fi60 - int(v38)))}
                                                var v62 float32 = (1.5 - 1.5)
                                                v63 := GenS2{F0: 18446744073709551615, F1: 42}
                                                _ = v61
                                                _ = v62
                                                _ = v63
                                            } else {
                                                var v64 int16 = int16(v2)
                                                var v65 string = "a"
                                                v38 = v64
                                                _ = v64
                                                _ = v65
                                            }
                                        case "hello":
                                            switch 4294967295 {
                                            case 42:
                                                v66 := GenS3{F0: "hello"}
                                                _ = v66
                                            case -2147483648:
                                                _ = len(v1)
                                            case -657:
                                                var v67 int8 = v4
                                                _ = v67
                                            }
                                        case "a":
                                            if true {
                                                v3 = append(v3, fi44)
                                                switch (fi60 / ((fi60) | 1)) {
                                                case 1:
                                                    v2 = v40
                                                case -661:
                                                    v1[2147483647] = 715
                                                case 32768:
                                                    var v68 int16 = v38
                                                    _ = v68
                                                default:
                                                    fmt.Println(v4, v38)
                                                }
                                                fmt.Println(v47, v40, v47)
                                            }
                                        default:
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        }
                                        _ = len(v1)
                                        v38 += v38
                                        _ = fi60
                                    }
                                    fmt.Println(v38)
                                    v69 := v1[543]
                                    _ = v69
                                } else {
                                    v70 := GenS3{F0: ("foo bar" + "foo bar")}
                                    var v71 float32 = 3.14159
                                    if (genFn1(fi44, "hello") != -1) {
                                        for fi72 := 0; fi72 < 13; fi72++ {
                                            v73 := GenS5{F0: 128, F1: 32767, F2: -32768}
                                            if true {
                                                var v74 int8 = v4
                                                if true {
                                                    v2 = (v40 - (v40 + v40))
                                                    v75 := GenS6{F0: ("" + "hello")}
                                                    var v76 int16 = v38
                                                    _ = len(v1)
                                                    _ = v75
                                                    _ = v76
                                                }
                                                _ = v74
                                            } else {
                                                v77 := GenS7{F0: fi44, F1: ("" + (("\n\t" + "\n\t") + "\n\t"))}
                                                if false {
                                                    if true {
                                                        var v78 int8 = v4
                                                        if (2 == 4294967295) {
                                                            v2 = v40
                                                            var v79 float32 = float32(v47)
                                                            var v80 float64 = (-2.25 / 0.1)
                                                            _ = v79
                                                            _ = v80
                                                        }
                                                        var v81 int64 = 9223372036854775807
                                                        _ = v78
                                                        _ = v81
                                                    }
                                                    if (((255 >= 1) && false) && true) {
                                                        fmt.Println(fi44, v40)
                                                        if ((1 <= 174) || true) {
                                                            if false {
                                                                var v82 float64 = float64(((v71 + (v71 / 1.5)) / v71))
                                                                if false {
                                                                    v82 = ((v82 - ((v82 / float64(v71)) * v82)) / 1e10)
                                                                    v83 := GenS8{F0: -1, F1: fi44, F2: fi44}
                                                                    _ = v83
                                                                } else {
                                                                    v47 = v47
                                                                    delete(v1, 0)
                                                                    var v84 int32 = 0
                                                                    if (!(!true)) {
                                                                        if false {
                                                                            v84 += v84
                                                                            switch (("" + (((("hello" + "a") + "hello") + "") + "hello")) + "a") {
                                                                            case "\n\t":
                                                                                v38 = int16(v40)
                                                                            default:
                                                                                if false {
                                                                                    _ = len(v1)
                                                                                    var v85 int64 = genFn1(fi72, (("a" + "foo bar") + ("\n\t" + "hello")))
                                                                                    if ((v38 & v38) > (v38 & v38)) {
                                                                                        switch ("hello" + "a") {
                                                                                        case "hello":
                                                                                            fmt.Println(v82)
                                                                                        default:
                                                                                            var v86 int64 = (v85 & (^(v85)))
                                                                                            _ = v86
                                                                                        }
                                                                                        if true {
                                                                                            var v87 int16 = (v38 / ((v38) | 1))
                                                                                            v87 += 1000
                                                                                            if false {
                                                                                                fmt.Println(v47)
                                                                                            }
                                                                                            _ = v87
                                                                                        }
                                                                                    } else {
                                                                                        var v88 uint16 = 255
                                                                                        _ = v88
                                                                                    }
                                                                                    _ = v85
                                                                                } else {
                                                                                    switch (fi72 ^ (fi72 ^ fi44)) {
                                                                                    case 2147483647:
                                                                                        v89 := GenS9{F0: 397, F1: (v40 + v2), F2: 0}
                                                                                        _ = v89
                                                                                    case -216:
                                                                                        v71 += ((0.0 + (1.5 / ((0.0 + float32(v82)) - v71))) / (((v71 - float32(-2.25)) + (-2.25 - (v71 * -2.25))) / ((0.0 * 3.14159) * (v71 * 1.5))))
                                                                                    case -1:
                                                                                        fmt.Println(v71)
                                                                                        if true {
                                                                                            v2 = 42
                                                                                            fmt.Println(fi72)
                                                                                        }
                                                                                        if true {
                                                                                            var v90 uint64 = 42
                                                                                            var v91 uint16 = 42
                                                                                            if len(v3) > 0 {
                                                                                                v3[0] = fi44
                                                                                            }
                                                                                            _ = v90
                                                                                            _ = v91
                                                                                        } else {
                                                                                            v82 = (0.1 / v47)
                                                                                            var v92 float32 = float32(((((v47 / (float64(-2.25) * v47)) * (1.5 + float64((((1e10 + v71) + ((v71 / (((v71 / v71) - (v71 / (v71 - v71))) / 0.1)) - (v71 + ((v71 * v71) * 1e10)))) * ((3.14159 / v71) - v71))))) * v47) * 1e10))
                                                                                            if (0 < 65535) {
                                                                                                v82 += v47
                                                                                                var v93 int64 = int64(fi72)
                                                                                                v84 = 42
                                                                                                _ = v93
                                                                                            } else {
                                                                                                var v94 float64 = v82
                                                                                                _ = v94
                                                                                            }
                                                                                            _ = v92
                                                                                        }
                                                                                        if (true && (((!(-203 != 0)) || ("" >= "foo bar")) || true)) {
                                                                                            fmt.Println(v40, v84)
                                                                                            v84 = (v84 / ((v84) | 1))
                                                                                        }
                                                                                        v38 = (v38 - v38)
                                                                                    }
                                                                                    v82 = (1e10 - 0.1)
                                                                                }
                                                                                fmt.Println(fi72)
                                                                                if false {
                                                                                    if ("a" == "foo bar") {
                                                                                        if true {
                                                                                            v2 -= v40
                                                                                        } else {
                                                                                            v95 := GenS2{F0: 4294967295, F1: 4294967295}
                                                                                            _ = v95
                                                                                        }
                                                                                    } else {
                                                                                        switch "hello" {
                                                                                        case "\n\t":
                                                                                            if false {
                                                                                                if false {
                                                                                                    v96 := GenS10{F0: v2}
                                                                                                    v97 := GenS4{F0: true, F1: 1, F2: fi44}
                                                                                                    _ = v96
                                                                                                    _ = v97
                                                                                                }
                                                                                                if (v40 >= 4294967295) {
                                                                                                    var v98 uint = (v2 >> ((18446744073709551615) & 63))
                                                                                                    switch "" {
                                                                                                    case "\n\t":
                                                                                                        if len(v5) >= 2 {
                                                                                                            v5 = v5[1:len(v5)]
                                                                                                        }
                                                                                                    }
                                                                                                    v47 -= 0.1
                                                                                                    _ = v98
                                                                                                }
                                                                                                var v99 int = int(v4)
                                                                                                _ = v99
                                                                                            } else {
                                                                                                v100 := GenS1{F0: 0, F1: 0, F2: 2}
                                                                                                _ = v100
                                                                                            }
                                                                                        case "a":
                                                                                            if false {
                                                                                                switch "\n\t" {
                                                                                                case "a":
                                                                                                    v1[895] = 65535
                                                                                                case "\n\t":
                                                                                                    if false {
                                                                                                        v47 -= 3.14159
                                                                                                        if ("a" < "a") {
                                                                                                            var v101 uint16 = 32767
                                                                                                            v102 := GenS5{F0: 0, F1: (v101 * (^(v101))), F2: (v38 ^ 1)}
                                                                                                            var v103 uint8 = 174
                                                                                                            var v104 int32 = v84
                                                                                                            v105 := GenS2{F0: 435, F1: 2}
                                                                                                            _ = v101
                                                                                                            _ = v102
                                                                                                            _ = v103
                                                                                                            _ = v104
                                                                                                            _ = v105
                                                                                                        } else {
                                                                                                            v106 := GenS6{F0: "a"}
                                                                                                            v107 := v1[2147483648]
                                                                                                            _ = v106
                                                                                                            _ = v107
                                                                                                        }
                                                                                                        v1[2147483647] = 1
                                                                                                    }
                                                                                                case "hello":
                                                                                                    v2 = (^(v2))
                                                                                                default:
                                                                                                    v71 += float32(v47)
                                                                                                }
                                                                                            } else {
                                                                                                var v108 uint8 = uint8(v40)
                                                                                                if true {
                                                                                                    if len(v3) >= 2 {
                                                                                                        v3 = v3[1:len(v3)]
                                                                                                    }
                                                                                                    fmt.Println(v4)
                                                                                                    var v109 uint = v40
                                                                                                    _ = v109
                                                                                                } else {
                                                                                                    v4 = v4
                                                                                                }
                                                                                                if (false || ("" <= "foo bar")) {
                                                                                                    switch ("\n\t" + "hello") {
                                                                                                    case "hello":
                                                                                                        if true {
                                                                                                            var v110 bool = true
                                                                                                            v2 = uint(v108)
                                                                                                            fmt.Println(v47, v47)
                                                                                                            _ = v110
                                                                                                        } else {
                                                                                                            v111 := GenS9{F0: 0, F1: 2, F2: (fi72 | -125)}
                                                                                                            _ = v111
                                                                                                        }
                                                                                                    }
                                                                                                    if false {
                                                                                                        var v112 float32 = float32(v47)
                                                                                                        var v113 uint = (^(v40))
                                                                                                        _ = v112
                                                                                                        _ = v113
                                                                                                    } else {
                                                                                                        var v114 float32 = ((0.0 * (1e10 * v71)) * v71)
                                                                                                        var v115 uint8 = (v108 & v108)
                                                                                                        _ = v114
                                                                                                        _ = v115
                                                                                                    }
                                                                                                    _ = len(v3)
                                                                                                }
                                                                                                v38 = v38
                                                                                                v71 -= (3.14159 / (v71 + v71))
                                                                                                _ = v108
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    fmt.Println(fi72, v71)
                                                                                    _ = len(v5)
                                                                                } else {
                                                                                    v116 := GenS11{F0: 0, F1: 230}
                                                                                    _ = v116
                                                                                }
                                                                                fmt.Println(v82, fi44)
                                                                            }
                                                                            v117 := GenS12{F0: 456, F1: (fi44 - (fi72 >> ((2) & 63))), F2: "foo bar"}
                                                                            v71 += ((v71 - v71) - v71)
                                                                            if (false && ((v4 + 0) < v4)) {
                                                                                v118 := GenS2{F0: 0, F1: 2147483648}
                                                                                switch (fi72 * 9223372036854775807) {
                                                                                case -242:
                                                                                    v40 += (v40 | (^(v40)))
                                                                                case 4294967295:
                                                                                    v4 += -1
                                                                                case 806:
                                                                                    var v119 float32 = 0.1
                                                                                    _ = v119
                                                                                default:
                                                                                    var v120 int32 = (^(v84))
                                                                                    _ = v120
                                                                                }
                                                                                var v121 uint = v40
                                                                                _ = v118
                                                                                _ = v121
                                                                            } else {
                                                                                fmt.Println(v82, v47, v82)
                                                                                var v122 string = "foo bar"
                                                                                _ = v122
                                                                            }
                                                                            _ = v117
                                                                        } else {
                                                                            v40 += 255
                                                                            delete(v1, 498)
                                                                            v71 -= v71
                                                                        }
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                        var v123 string = ("\n\t" + "\n\t")
                                                                        fmt.Println(v84)
                                                                        var v124 int16 = v38
                                                                        var v125 uint16 = 0
                                                                        var v126 uint16 = v125
                                                                        _ = v123
                                                                        _ = v124
                                                                        _ = v125
                                                                        _ = v126
                                                                    }
                                                                    _ = v84
                                                                }
                                                                v38 = (v38 & (^(v38)))
                                                                if false {
                                                                    v40 = v40
                                                                    var v127 float64 = 3.14159
                                                                    _ = v127
                                                                }
                                                                _ = v82
                                                            }
                                                        }
                                                        v2 = v2
                                                    } else {
                                                        _ = len(v5)
                                                        v128 := GenS4{F0: true, F1: -9223372036854775808, F2: -2147483648}
                                                        _ = v128
                                                    }
                                                } else {
                                                    var v129 float64 = -2.25
                                                    _ = v129
                                                }
                                                v71 += float32(float64((v71 - (0.0 + v71))))
                                                _ = v77
                                            }
                                            var v130 int16 = (v38 - (v38 << ((47) & 63)))
                                            _ = fi72
                                            _ = v130
                                            _ = v73
                                        }
                                        if true {
                                            var v131 int = 32767
                                            switch 586 {
                                            case 970:
                                                v2 += (v40 * (^(v2)))
                                            case 128:
                                                if true {
                                                    v132 := GenS12{F0: (v2 ^ (v2 - (v2 ^ (v40 << ((v2) & 63))))), F1: v131, F2: ("hello" + "a")}
                                                    _ = v132
                                                } else {
                                                    var v133 bool = true
                                                    delete(v1, 219)
                                                    _ = v133
                                                }
                                            }
                                            fmt.Println(v47, v40)
                                            _ = v131
                                        } else {
                                            for fi134 := 0; fi134 < 1; fi134++ {
                                                v40 = 4294967295
                                                _ = fi134
                                            }
                                            if true {
                                                v3 = append(v3, fi44)
                                                if len(v3) > 0 {
                                                    v3[0] = fi44
                                                }
                                            } else {
                                                v4 = v4
                                                var v135 int = fi44
                                                var v136 string = ("" + "hello")
                                                _ = v135
                                                _ = v136
                                            }
                                        }
                                    } else {
                                        v38 += v38
                                        v71 = (-2.25 + 3.14159)
                                    }
                                    _ = v70
                                    _ = v71
                                }
                                for fi137 := 0; fi137 < 15; fi137++ {
                                    switch "a" {
                                    case "\n\t":
                                        var v138 uint = (v40 - (v40 / ((731) | 1)))
                                        _ = v138
                                    case "":
                                        v47 += 0.0
                                    default:
                                        v38 = (-(v38))
                                    }
                                    var v139 uint32 = 2147483647
                                    _ = fi137
                                    _ = v139
                                }
                            }
                        }
                        var v140 uint16 = 0
                        _ = v140
                    }
                    _ = v47
                }
                _ = fi44
            }
        } else {
            v2 -= v2
            v141 := v1[2]
            _ = v141
        }
        _ = v43
    }
    for fi142 := 0; fi142 < 1; fi142++ {
        v4 -= (-(v4))
        _ = fi142
    }
    _ = len(v5)
    _ = len(v5)
    for fi143 := 0; fi143 < 2; fi143++ {
        if true {
            var v144 int8 = 127
            var v145 bool = false
            _ = v144
            _ = v145
        } else {
            if true {
                switch fi143 {
                case -9223372036854775808:
                    v40 += v2
                case 1:
                    v2 += v2
                }
                var v146 int = fi143
                _ = v146
            } else {
                v147 := GenS1{F0: genFn1(2147483647, "hello"), F1: 32767, F2: 190}
                _ = v147
            }
        }
        var v148 int16 = v38
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi143
        _ = v148
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
