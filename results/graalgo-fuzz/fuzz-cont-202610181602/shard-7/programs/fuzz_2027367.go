package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 uint8; F1 uint32; F2 uint64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint64; F1 bool; F2 uint32 }
type GenS7 struct { F0 uint16; F1 int16 }
type GenS8 struct { F0 int; F1 bool }
func genFn0(p0 string, p1 string) bool {
    v4 := GenS1{F0: 0}
    p1 = (("" + p0) + p0)
    if ((false || true) && (false || (1 == 0))) {
        if false {
            var v5 bool = (true || true)
            v5 = v5
            if false {
                for fi6 := 0; fi6 < 10; fi6++ {
                    p0 = "foo bar"
                    _ = fi6
                }
                if (!true) {
                    var v7 float64 = 1e10
                    _ = v7
                }
                v8 := GenS1{F0: -1}
                _ = v8
            }
            _ = v5
        } else {
            if false {
                var v9 int32 = 0
                fmt.Println(v9, v9, p0)
                _ = v9
            }
            v10 := []int{32767}
            p0 = p1
            _ = v10
        }
    } else {
        v11 := GenS1{F0: 127}
        for fi12 := 0; fi12 < 3; fi12++ {
            var v13 uint8 = 127
            v13 += v13
            var v14 float32 = 0.0
            _ = fi12
            _ = v13
            _ = v14
        }
        for fi15 := 0; fi15 < 7; fi15++ {
            for fi16 := 0; fi16 < 3; fi16++ {
                p0 = p0
                var v17 bool = false
                _ = fi16
                _ = v17
            }
            for fi18 := 0; fi18 < 16; fi18++ {
                v19 := []int{fi18, (^(fi15))}
                var v20 int16 = 127
                _ = fi18
                _ = v19
                _ = v20
            }
            _ = fi15
        }
        _ = v11
    }
    var v21 float64 = (1.5 / (3.14159 + 3.14159))
    v22 := []int{128, 128, 0}
    _ = p0
    _ = p1
    _ = v21
    _ = v22
    _ = v4
    return (0 > 42)
}

func main() {
    v1 := []int{704, -1, 44}
    var v2 string = ("" + "hello")
    v3 := make(map[string]int)
    if true {
        var v23 int32 = 0
        if true {
            var v24 int32 = 1
            v25 := GenS1{F0: int(v23)}
            v24 = v23
            _ = v24
            _ = v25
        } else {
            if true {
                v2 = v2
                var v26 uint16 = 0
                _ = v26
            } else {
                v27 := GenS1{F0: 721}
                var v28 uint8 = 2
                var v29 int8 = 42
                _ = v27
                _ = v28
                _ = v29
            }
        }
        _ = v23
    } else {
        switch v2 {
        case "":
            fmt.Println(v2)
        case "hello":
            for fi30 := 0; fi30 < 6; fi30++ {
                switch 2147483647 {
                case 128:
                    if (255 >= 2) {
                        if true {
                            for fi31 := 0; fi31 < 2; fi31++ {
                                var v32 int32 = 0
                                switch "a" {
                                case "a":
                                    v32 = (-(v32))
                                default:
                                    var v33 uint = 381
                                    _ = v33
                                }
                                _ = fi31
                                _ = v32
                            }
                            for fi34 := 0; fi34 < 4; fi34++ {
                                var v35 uint32 = 2
                                v36 := v3[v2]
                                _ = fi34
                                _ = v35
                                _ = v36
                            }
                        } else {
                            v37 := GenS1{F0: (fi30 >> ((80) & 63))}
                            v2 = v2
                            _ = v37
                        }
                        var v38 int = fi30
                        _ = v38
                    } else {
                        switch (fi30 << ((1) & 63)) {
                        case 882:
                            var v39 int16 = 305
                            _ = v39
                        default:
                            v40 := GenS1{F0: fi30}
                            _ = v40
                        }
                    }
                case -271:
                    var v41 string = ""
                    _ = v41
                case 127:
                    fmt.Println(fi30, fi30)
                }
                v3[(v2 + v2)] = (^(fi30))
                if false {
                    if true {
                        for fi42 := 0; fi42 < 1; fi42++ {
                            fmt.Println(fi42, fi42, fi42)
                            var v43 float32 = (-2.25 / -2.25)
                            switch ("foo bar" + "") {
                            case "hello":
                                fmt.Println(v43, v2, fi30)
                            }
                            _ = fi42
                            _ = v43
                        }
                        for fi44 := 0; fi44 < 18; fi44++ {
                            var v45 float64 = (float64((0.1 * float32(0.1))) + (-2.25 - (3.14159 - ((3.14159 * 1e10) / (((0.0 * 1e10) - 1e10) * 1e10)))))
                            v1 = append(v1, -488)
                            var v46 int16 = 263
                            _ = fi44
                            _ = v45
                            _ = v46
                        }
                        _ = len(v3)
                    } else {
                        v47 := GenS2{F0: 255, F1: 0}
                        v48 := GenS2{F0: 1, F1: 32767}
                        switch 0 {
                        case -2147483648:
                            v2 = v2
                        case 0:
                            v2 = v2
                        case 128:
                            for fi49 := 0; fi49 < 15; fi49++ {
                                v2 = "foo bar"
                                v50 := GenS3{F0: 65535, F1: 0}
                                var v51 int8 = 0
                                _ = fi49
                                _ = v50
                                _ = v51
                            }
                        default:
                            v52 := GenS4{F0: 212, F1: 2, F2: 0}
                            _ = v52
                        }
                        _ = v47
                        _ = v48
                    }
                    var v53 uint32 = 0
                    if false {
                        if ((fi30 >> ((705) & 63)) != -1) {
                            var v54 int32 = 2147483647
                            fmt.Println(v2, v54, v2)
                            fmt.Println(v2)
                            _ = v54
                        } else {
                            v53 += v53
                            for fi55 := 0; fi55 < 8; fi55++ {
                                v56 := GenS3{F0: (v53 >> ((2) & 63)), F1: 800}
                                v2 = (v2 + "foo bar")
                                var v57 bool = ((128 >= 128) || (1 < -1))
                                _ = fi55
                                _ = v56
                                _ = v57
                            }
                        }
                        for fi58 := 0; fi58 < 14; fi58++ {
                            var v59 string = v2
                            _ = fi58
                            _ = v59
                        }
                        fmt.Println(v53, v2)
                    } else {
                        for fi60 := 0; fi60 < 15; fi60++ {
                            v2 = "\n\t"
                            _ = fi60
                        }
                    }
                    _ = v53
                } else {
                    var v61 uint8 = 209
                    if ((v2 + ((((v2 + v2) + ("hello" + v2)) + v2) + v2)) != "foo bar") {
                        for fi62 := 0; fi62 < 15; fi62++ {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = fi62
                        }
                        var v63 uint16 = 0
                        _ = len(v1)
                        _ = v63
                    }
                    _ = v61
                }
                _ = fi30
            }
        }
    }
    var v64 int = 2147483647
    if (!true) {
        delete(v3, "\n\t")
        if (655 >= 42) {
            if ((-(v64)) > v64) {
                fmt.Println(v64, v2)
            } else {
                switch v2 {
                case "\n\t":
                    v64 += v64
                case "":
                    var v65 uint = 894
                    _ = v65
                default:
                    for fi66 := 0; fi66 < 19; fi66++ {
                        switch (((v2 + v2) + ("\n\t" + "")) + v2) {
                        case "\n\t":
                            v64 -= fi66
                        case "hello":
                            v2 = "hello"
                        default:
                            switch ("foo bar" + ("hello" + "\n\t")) {
                            case "a":
                                v67 := v3[(v2 + "foo bar")]
                                _ = v67
                            default:
                                v1 = append(v1, 128)
                            }
                        }
                        if false {
                            v2 = ((v2 + (v2 + ("a" + (v2 + ((v2 + v2) + v2))))) + ("a" + v2))
                            var v68 uint16 = 137
                            switch v64 {
                            case 42:
                                var v69 uint16 = 0
                                _ = v69
                            }
                            _ = v68
                        } else {
                            switch v2 {
                            case "hello":
                                var v70 string = ("hello" + v2)
                                _ = v70
                            case "\n\t":
                                if true {
                                    if len(v1) > 0 {
                                        v1[0] = fi66
                                    }
                                } else {
                                    for fi71 := 0; fi71 < 3; fi71++ {
                                        if true {
                                            v64 = (fi66 + fi66)
                                            v72 := GenS2{F0: 1, F1: 1000}
                                            v2 = ((v2 + v2) + v2)
                                            if (18446744073709551615 >= 2) {
                                                v64 -= (v64 + (fi66 / (((fi66 & fi71)) | 1)))
                                                var v73 int8 = 42
                                                _ = v73
                                            } else {
                                                v74 := GenS5{F0: 328}
                                                var v75 float64 = 0.1
                                                var v76 int64 = 0
                                                v77 := GenS5{F0: v76}
                                                var v78 uint8 = 127
                                                _ = v74
                                                _ = v75
                                                _ = v76
                                                _ = v77
                                                _ = v78
                                            }
                                            _ = v72
                                        }
                                        switch "\n\t" {
                                        case "foo bar":
                                            v2 = (v2 + "")
                                        }
                                        _ = fi71
                                    }
                                }
                            case "a":
                                v79 := GenS3{F0: 623, F1: 2}
                                _ = v79
                            default:
                                if len(v1) > 0 {
                                    v1[0] = v64
                                }
                            }
                            var v80 uint16 = 32767
                            _ = v80
                        }
                        _ = fi66
                    }
                }
            }
            v81 := GenS4{F0: 2, F1: 65535, F2: 2}
            _ = v81
        }
        for fi82 := 0; fi82 < 5; fi82++ {
            fmt.Println(v64, v64)
            if false {
                v2 = v2
            } else {
                var v83 int8 = 100
                _ = v83
            }
            for fi84 := 0; fi84 < 17; fi84++ {
                var v85 uint16 = 967
                switch v2 {
                case "hello":
                    fmt.Println(fi84, v2, v85)
                case "a":
                    v64 = (fi84 ^ (fi84 & 2147483647))
                case "foo bar":
                    delete(v3, v2)
                default:
                    fmt.Println(v2)
                }
                switch (fi84 * (fi84 >> ((4294967295) & 63))) {
                case 349:
                    v86 := GenS3{F0: 65535, F1: 532}
                    _ = v86
                case 127:
                    v64 += v64
                default:
                    v64 += v64
                }
                switch v2 {
                case "\n\t":
                    if (v64 <= fi84) {
                        if false {
                            _ = len(v3)
                            var v87 bool = (((!true) || true) && ((float64((0.0 * 0.0)) / (1.5 - (1.5 - (1e10 / 1.5)))) > float64((3.14159 - (3.14159 / 3.14159)))))
                            _ = v87
                        } else {
                            switch v2 {
                            case "hello":
                                if false {
                                    fmt.Println(v85)
                                    v64 = (fi82 * -2147483648)
                                    var v88 string = (v2 + v2)
                                    if (true && (!false)) {
                                        if (!false) {
                                            _ = len(v1)
                                        }
                                        var v89 int32 = 42
                                        _ = v89
                                    }
                                    switch 32767 {
                                    case 115:
                                        if false {
                                            switch v88 {
                                            case "foo bar":
                                                v1 = append(v1, int(v85))
                                            case "\n\t":
                                                fmt.Println(v88, fi84)
                                            default:
                                                v90 := GenS6{F0: 799, F1: false, F2: 2147483648}
                                                v3[("foo bar" + v88)] = (fi84 + fi84)
                                                v64 = -487
                                                _ = v90
                                            }
                                        }
                                    case -611:
                                        if (!(!((v2 != v88) || (0.0 == float64(3.14159))))) {
                                            if (v88 < ("hello" + "hello")) {
                                                switch v64 {
                                                case 2147483647:
                                                    v88 = "foo bar"
                                                case 32768:
                                                    fmt.Println(fi84, fi84)
                                                }
                                            }
                                        }
                                    }
                                    var v91 uint8 = 150
                                    fmt.Println(fi82, fi82, v91)
                                    _ = v88
                                    _ = v91
                                } else {
                                    switch "foo bar" {
                                    case "\n\t":
                                        v92 := GenS2{F0: 960, F1: 859}
                                        _ = v92
                                    case "hello":
                                        var v93 float32 = (((0.0 - (0.0 / 0.1)) - ((3.14159 / -2.25) / (float32((float64((-2.25 - -2.25)) + 0.1)) + 3.14159))) * 3.14159)
                                        _ = v93
                                    case "":
                                        if false {
                                            if true {
                                                v2 = ("hello" + (((v2 + v2) + ("a" + v2)) + (v2 + ((v2 + v2) + v2))))
                                                var v94 uint32 = 4294967295
                                                if true {
                                                    var v95 float64 = (float64((float32(0.0) - 1.5)) * ((1.5 + (float64(((float32(-2.25) + 3.14159) * (0.0 / 1e10))) + -2.25)) * ((3.14159 - float64(0.0)) - 0.1)))
                                                    _ = v95
                                                }
                                                _ = v94
                                            } else {
                                                var v96 uint16 = (^(v85))
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                if len(v1) > 0 {
                                                    v1[0] = (fi84 * 32768)
                                                }
                                                _ = v96
                                            }
                                        }
                                    default:
                                        v64 -= fi82
                                    }
                                    v64 += (fi82 >> ((65535) & 63))
                                }
                            case "":
                                var v97 int32 = 520
                                _ = v97
                            default:
                                var v98 uint32 = 299
                                _ = v98
                            }
                            v99 := v3[""]
                            _ = v99
                        }
                        var v100 uint32 = 884
                        v2 = (v2 + v2)
                        _ = v100
                    }
                }
                if (v64 == -1) {
                    fmt.Println(fi84, v2)
                }
                v85 += (v85 >> ((550) & 63))
                if (!true) {
                    var v101 int = fi82
                    v64 = 32768
                    _ = v101
                }
                _ = fi84
                _ = v85
            }
            _ = fi82
        }
    }
    if (-508 >= 42) {
        if false {
            v102 := GenS1{F0: 2147483647}
            _ = v102
        } else {
            v64 -= (v64 + 0)
            if (0 >= 42) {
                v1 = append(v1, -1)
                var v103 float64 = 3.14159
                _ = v103
            }
            var v104 uint = 2
            _ = v104
        }
        fmt.Println(v64)
    }
    v2 = ((("hello" + ((("hello" + v2) + (v2 + v2)) + v2)) + "") + v2)
    for fi105 := 0; fi105 < 15; fi105++ {
        var v106 int = fi105
        v106 -= (v64 - v106)
        v107 := GenS7{F0: 532, F1: -130}
        _ = fi105
        _ = v106
        _ = v107
    }
    v108 := GenS6{F0: 1, F1: false, F2: 65535}
    v109 := GenS5{F0: -114}
    v110 := GenS8{F0: v64, F1: false}
    if true {
        v1 = append(v1, (v64 % ((9223372036854775807) | 1)))
        var v111 int32 = 950
        _ = v111
    } else {
        v64 += v64
        for fi112 := 0; fi112 < 4; fi112++ {
            v64 = (v64 - (v64 ^ fi112))
            switch v2 {
            case "a":
                switch "foo bar" {
                case "a":
                    v2 = v2
                case "hello":
                    if (false && (!false)) {
                        switch (v64 % (((v64 / (((-(v64))) | 1))) | 1)) {
                        case 128:
                            fmt.Println(v2, v2, v2)
                        }
                    }
                default:
                    for fi113 := 0; fi113 < 1; fi113++ {
                        var v114 int32 = 103
                        switch "foo bar" {
                        case "foo bar":
                            v2 = v2
                        case "\n\t":
                            v115 := GenS3{F0: 518, F1: 463}
                            _ = v115
                        case "hello":
                            var v116 uint64 = 18446744073709551615
                            v117 := v3[v2]
                            v1 = append(v1, -506)
                            _ = v116
                            _ = v117
                        default:
                            _ = len(v1)
                        }
                        _ = fi113
                        _ = v114
                    }
                }
            case "\n\t":
                switch "foo bar" {
                case "\n\t":
                    var v118 uint16 = 32767
                    _ = v118
                }
            default:
                v64 = (fi112 >> ((255) & 63))
            }
            _ = fi112
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v64:", v64)
    _ = v1
    _ = v108
    _ = v109
    _ = v110
    _ = v2
    _ = v3
    _ = v64
}
