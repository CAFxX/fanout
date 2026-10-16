package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 int; F1 uint64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int8; F1 int64 }
type GenS8 struct { F0 int16; F1 int8 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint; F1 int64; F2 uint8 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 int8; F1 int }
type GenS13 struct { F0 uint; F1 bool }
type GenS14 struct { F0 uint16; F1 uint64 }
type GenS15 struct { F0 uint32; F1 uint }
func genFn0(p0 bool) uint8 {
    var v7 uint8 = 90
    var v8 bool = (42 > 531)
    var v9 uint16 = 2
    for fi10 := 0; fi10 < 20; fi10++ {
        switch ("hello" + "") {
        case "\n\t":
            for fi11 := 0; fi11 < 14; fi11++ {
                if false {
                    var v12 string = ""
                    _ = v12
                }
                var v13 int32 = 2147483647
                _ = fi11
                _ = v13
            }
        case "":
            v9 -= (v9 << ((502) & 63))
        }
        v14 := []int{1}
        v7 -= uint8(v9)
        _ = fi10
        _ = v14
    }
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return (v7 * (^(v7)))
}

func main() {
    var v1 uint64 = 784
    var v2 int8 = -128
    var v3 int = -506
    var v4 int = (-(v3))
    var v5 int32 = 42
    var v6 uint16 = 32767
    v15 := []int{(v4 * v3), v3, v4}
    if false {
        if ((v1 >> ((758) & 63)) >= (v1 << ((255) & 63))) {
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
            v5 = 603
            v6 -= (v6 * (v6 - v6))
        } else {
            v15 = append(v15, v3)
            var v16 uint64 = (v1 | 4294967295)
            _ = v16
        }
    } else {
        switch "" {
        case "a":
            var v17 string = ("foo bar" + "foo bar")
            _ = v17
        }
        if false {
            if true {
                var v18 int = v4
                var v19 int8 = 100
                if false {
                    switch (("a" + "") + ("" + "foo bar")) {
                    case "hello":
                        v2 -= v19
                    default:
                        var v20 int16 = -1
                        _ = v20
                    }
                }
                _ = v18
                _ = v19
            }
        } else {
            fmt.Println(v3, v6, v5)
        }
        v21 := GenS1{F0: v4, F1: -128}
        _ = v21
    }
    switch 42 {
    case 0:
        var v22 string = ("foo bar" + "foo bar")
        _ = v22
    }
    v23 := GenS1{F0: (v3 * v4), F1: (v2 % (((^(v2))) | 1))}
    for fi24 := 0; fi24 < 3; fi24++ {
        for fi25 := 0; fi25 < 6; fi25++ {
            if len(v15) > 0 {
                v15[0] = (fi24 % (((v4 + int(v5))) | 1))
            }
            v15 = append(v15, (fi24 << ((65535) & 63)))
            if len(v15) > 0 {
                v15[0] = fi25
            }
            _ = fi25
        }
        var v26 float64 = 1e10
        _ = fi24
        _ = v26
    }
    for fi27 := 0; fi27 < 10; fi27++ {
        v28 := GenS2{F0: (-(v3))}
        _ = fi27
        _ = v28
    }
    switch (v3 << ((255) & 63)) {
    case 0:
        switch v4 {
        case -9223372036854775808:
            for fi29 := 0; fi29 < 12; fi29++ {
                v30 := GenS3{F0: (v1 < 0), F1: 161}
                for fi31 := 0; fi31 < 16; fi31++ {
                    v2 = (v2 << ((18446744073709551615) & 63))
                    var v32 int64 = 1
                    v5 = (-(v5))
                    _ = fi31
                    _ = v32
                }
                _ = fi29
                _ = v30
            }
        default:
            if (1000 < -1) {
                for fi33 := 0; fi33 < 3; fi33++ {
                    switch -423 {
                    case 0:
                        v34 := GenS4{F0: 363, F1: (v1 / ((v1) | 1))}
                        _ = v34
                    case 42:
                        var v35 float32 = (((float32(float64(1.5)) / -2.25) - -2.25) - ((0.0 + 0.0) / 1e10))
                        _ = v35
                    case -2147483648:
                        if true {
                            if (("hello" + "hello") == ("foo bar" + "foo bar")) {
                                fmt.Println(v5)
                                if (579 != -1) {
                                    v36 := GenS5{F0: -96}
                                    var v37 uint = 2
                                    _ = v36
                                    _ = v37
                                } else {
                                    v38 := GenS4{F0: v4, F1: (v1 - (^(v1)))}
                                    _ = v38
                                }
                            } else {
                                v5 += -1
                                if len(v15) > 0 {
                                    v15[0] = -1
                                }
                            }
                        } else {
                            for fi39 := 0; fi39 < 13; fi39++ {
                                var v40 uint64 = v1
                                if len(v15) > 0 {
                                    v15[0] = v4
                                }
                                _ = fi39
                                _ = v40
                            }
                            var v41 string = "a"
                            _ = v41
                        }
                    default:
                        for fi42 := 0; fi42 < 14; fi42++ {
                            fmt.Println(v2, v6)
                            v43 := GenS4{F0: -308, F1: (v1 << ((65535) & 63))}
                            _ = fi42
                            _ = v43
                        }
                    }
                    _ = fi33
                }
                v44 := GenS2{F0: (^(v4))}
                _ = v44
            } else {
                v45 := GenS6{F0: v2}
                var v46 int16 = 1
                if true {
                    for fi47 := 0; fi47 < 1; fi47++ {
                        var v48 bool = ((^(v46)) == v46)
                        var v49 uint32 = 65535
                        _ = fi47
                        _ = v48
                        _ = v49
                    }
                    v3 += (-(v4))
                }
                _ = v45
                _ = v46
            }
        }
    default:
        for fi50 := 0; fi50 < 11; fi50++ {
            if (!false) {
                v6 -= 930
                if false {
                    for fi51 := 0; fi51 < 6; fi51++ {
                        fmt.Println(v2, fi50)
                        fmt.Println(fi51, v4, v4)
                        _ = fi51
                    }
                    v1 = (v1 - (^(v1)))
                } else {
                    v2 = (v2 % ((42) | 1))
                }
                for fi52 := 0; fi52 < 15; fi52++ {
                    if true {
                        var v53 int32 = -1
                        if false {
                            var v54 int16 = -869
                            _ = v54
                        }
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                        _ = v53
                    } else {
                        v5 = v5
                        var v55 int = v4
                        var v56 string = "foo bar"
                        _ = v55
                        _ = v56
                    }
                    _ = fi52
                }
            }
            v57 := GenS7{F0: (v2 & (-(v2))), F1: -9223372036854775808}
            _ = fi50
            _ = v57
        }
    }
    if true {
        for fi58 := 0; fi58 < 10; fi58++ {
            var v59 bool = (-1 < 42)
            if false {
                fmt.Println(v59, v5)
                if false {
                    v60 := GenS6{F0: (v2 << ((4294967295) & 63))}
                    v61 := GenS2{F0: (fi58 & fi58)}
                    _ = v60
                    _ = v61
                }
                v1 = 4294967295
            } else {
                var v62 int32 = (v5 ^ 42)
                v59 = (("a" + ("a" + "\n\t")) == "a")
                _ = v62
            }
            if ((v1 % (((v1 | v1)) | 1)) < v1) {
                var v63 uint64 = v1
                v15 = append(v15, 128)
                _ = v63
            }
            _ = fi58
            _ = v59
        }
        v6 = v6
    } else {
        if (23 > 2) {
            v64 := GenS8{F0: -1, F1: (v2 >> ((677) & 63))}
            switch "foo bar" {
            case "foo bar":
                for fi65 := 0; fi65 < 16; fi65++ {
                    if false {
                        if true {
                            switch "" {
                            case "hello":
                                v66 := GenS9{F0: 0}
                                _ = v66
                            case "a":
                                var v67 uint = 2
                                _ = v67
                            case "foo bar":
                                v68 := GenS10{F0: 67, F1: 838, F2: 42}
                                _ = v68
                            default:
                                if (0.1 >= (float32(float64(1.5)) * (3.14159 - (((0.0 * 3.14159) / 1.0) + ((-2.25 / (-2.25 - (3.14159 - (0.0 / 1e10)))) * (3.14159 * -2.25)))))) {
                                    if (9223372036854775807 >= 9223372036854775807) {
                                        v6 = (v6 ^ v6)
                                    }
                                    for fi69 := 0; fi69 < 20; fi69++ {
                                        _ = len(v15)
                                        v5 = 2147483647
                                        v1 = 18446744073709551615
                                        _ = fi69
                                    }
                                }
                            }
                            for fi70 := 0; fi70 < 7; fi70++ {
                                switch v3 {
                                case -731:
                                    switch int(v1) {
                                    case -9223372036854775808:
                                        var v71 string = "\n\t"
                                        _ = v71
                                    case 650:
                                        v72 := GenS5{F0: v2}
                                        _ = v72
                                    case 128:
                                        v73 := GenS11{F0: 1000}
                                        _ = v73
                                    default:
                                        var v74 int32 = v5
                                        _ = v74
                                    }
                                case 952:
                                    if true {
                                        _ = len(v15)
                                        v4 = (v3 >> ((255) & 63))
                                    } else {
                                        v75 := GenS3{F0: (-416 != fi70), F1: 2}
                                        switch (v4 + v3) {
                                        case 506:
                                            if (608 < 42) {
                                                var v76 int16 = 32767
                                                _ = v76
                                            } else {
                                                v6 = 2
                                                var v77 float64 = ((3.14159 + (-2.25 - 1e10)) / (float64(0.1) * (1.5 + 0.0)))
                                                _ = v77
                                            }
                                        case 128:
                                            v6 = uint16(v1)
                                        default:
                                            switch "a" {
                                            case "\n\t":
                                                if false {
                                                    v78 := GenS2{F0: fi65}
                                                    _ = v78
                                                } else {
                                                    switch (v4 | 32768) {
                                                    case -251:
                                                        v5 = v5
                                                    case 128:
                                                        switch int(v1) {
                                                        case 32767:
                                                            var v79 string = "a"
                                                            _ = v79
                                                        case 127:
                                                            v80 := GenS12{F0: v2, F1: (v4 + v3)}
                                                            _ = v80
                                                        }
                                                    case 32768:
                                                        if false {
                                                            var v81 int16 = 1000
                                                            v1 = v1
                                                            var v82 string = ("a" + (("foo bar" + "hello") + "foo bar"))
                                                            _ = v81
                                                            _ = v82
                                                        } else {
                                                            switch ("\n\t" + "") {
                                                            case "\n\t":
                                                                switch "\n\t" {
                                                                case "":
                                                                    if (true || ((fi65 % ((42) | 1)) <= (fi65 / ((int(v2)) | 1)))) {
                                                                        v83 := GenS7{F0: v2, F1: 1}
                                                                        v15 = append(v15, v4)
                                                                        _ = v83
                                                                    } else {
                                                                        v5 += v5
                                                                    }
                                                                case "hello":
                                                                    fmt.Println(v3, v3, v3)
                                                                default:
                                                                    v1 = (v1 >> ((0) & 63))
                                                                }
                                                            case "":
                                                                if true {
                                                                    if (1e10 != (0.1 * 1.5)) {
                                                                        fmt.Println(v2)
                                                                        var v84 float32 = (-2.25 - (float32((float64(0.1) * 0.1)) + 0.0))
                                                                        if true {
                                                                            if ((v5 % ((1) | 1)) != (v5 ^ (^(v5)))) {
                                                                                var v85 int64 = 857
                                                                                if false {
                                                                                    v85 = v85
                                                                                    fmt.Println(v1, v5)
                                                                                }
                                                                                if len(v15) >= 2 {
                                                                                    v15 = v15[1:len(v15)]
                                                                                }
                                                                                _ = v85
                                                                            }
                                                                            v5 -= (v5 | (^(v5)))
                                                                        } else {
                                                                            fmt.Println(v3, fi65)
                                                                            v86 := GenS13{F0: 756, F1: true}
                                                                            _ = v86
                                                                        }
                                                                        fmt.Println(v2)
                                                                        v87 := GenS4{F0: (^(v3)), F1: v1}
                                                                        v4 -= fi65
                                                                        fmt.Println(fi70, v4, v4)
                                                                        _ = v84
                                                                        _ = v87
                                                                    }
                                                                    var v88 string = ""
                                                                    var v89 float64 = (0.1 + 0.1)
                                                                    _ = v88
                                                                    _ = v89
                                                                } else {
                                                                    var v90 bool = false
                                                                    switch v4 {
                                                                    case 128:
                                                                        var v91 float32 = 0.0
                                                                        if (v5 == (v5 / ((v5) | 1))) {
                                                                            var v92 int = fi65
                                                                            if ((v3 % ((v4) | 1)) == -175) {
                                                                                switch v92 {
                                                                                case -9223372036854775808:
                                                                                    v2 -= v2
                                                                                }
                                                                                v91 -= (1.5 + (v91 - 0.0))
                                                                                if len(v15) > 0 {
                                                                                    v15[0] = 32768
                                                                                }
                                                                            }
                                                                            var v93 int64 = 42
                                                                            _ = v92
                                                                            _ = v93
                                                                        }
                                                                        var v94 float32 = float32(-2.25)
                                                                        _ = v91
                                                                        _ = v94
                                                                    default:
                                                                        var v95 int = 127
                                                                        _ = v95
                                                                    }
                                                                    if (v90 || false) {
                                                                        var v96 float64 = (-2.25 + 1e10)
                                                                        v3 = (^(fi70))
                                                                        _ = v96
                                                                    } else {
                                                                        v15 = append(v15, (v4 + v3))
                                                                        v15 = append(v15, (v4 >> ((1) & 63)))
                                                                        if v90 {
                                                                            _ = len(v15)
                                                                        } else {
                                                                            if false {
                                                                                v97 := GenS14{F0: v6, F1: v1}
                                                                                _ = v97
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v90
                                                                }
                                                            default:
                                                                if true {
                                                                    var v98 float64 = 0.0
                                                                    _ = v98
                                                                } else {
                                                                    v5 = v5
                                                                    v99 := GenS2{F0: (^(v4))}
                                                                    var v100 float32 = (float32(0.1) / 1.0)
                                                                    _ = v100
                                                                    _ = v99
                                                                }
                                                            }
                                                            if (v1 >= 4294967295) {
                                                                _ = len(v15)
                                                            } else {
                                                                switch "\n\t" {
                                                                case "foo bar":
                                                                    fmt.Println(v2)
                                                                case "a":
                                                                    fmt.Println(fi65, v6)
                                                                case "":
                                                                    var v101 string = ("" + "\n\t")
                                                                    _ = v101
                                                                default:
                                                                    v1 = v1
                                                                }
                                                                var v102 uint64 = 468
                                                                _ = v102
                                                            }
                                                            switch (v4 - (fi65 & v4)) {
                                                            case -338:
                                                                v103 := GenS4{F0: (v3 * v4), F1: uint64(v6)}
                                                                _ = v103
                                                            case 32768:
                                                                var v104 int16 = -1
                                                                _ = v104
                                                            default:
                                                                fmt.Println(v6, v4, v2)
                                                            }
                                                            if false {
                                                                fmt.Println(v4, v1)
                                                                var v105 uint64 = v1
                                                                _ = v105
                                                            }
                                                        }
                                                    }
                                                }
                                            case "hello":
                                                _ = len(v15)
                                            }
                                        }
                                        _ = v75
                                    }
                                default:
                                    var v106 int64 = 1
                                    _ = v106
                                }
                                _ = fi70
                            }
                            for fi107 := 0; fi107 < 13; fi107++ {
                                switch "" {
                                case "a":
                                    var v108 int8 = 0
                                    v5 -= (v5 << ((255) & 63))
                                    _ = v108
                                case "":
                                    var v109 int = (fi65 / ((v4) | 1))
                                    _ = v109
                                }
                                v6 += (v6 ^ v6)
                                var v110 uint = 879
                                v111 := GenS11{F0: 1}
                                _ = fi107
                                _ = v110
                                _ = v111
                            }
                        }
                        v15 = append(v15, (v4 >> ((255) & 63)))
                    }
                    _ = fi65
                }
            case "a":
                if ((!false) && false) {
                    if len(v15) > 0 {
                        v15[0] = 32768
                    }
                    v112 := GenS1{F0: 32768, F1: (^(v2))}
                    if true {
                        var v113 uint16 = v6
                        fmt.Println(v1)
                        var v114 int = 0
                        _ = v113
                        _ = v114
                    }
                    _ = v112
                } else {
                    v1 = (v1 ^ 4294967295)
                    var v115 int = v4
                    _ = v115
                }
            case "hello":
                for fi116 := 0; fi116 < 11; fi116++ {
                    v117 := GenS1{F0: (-(fi116)), F1: (v2 / ((20) | 1))}
                    _ = fi116
                    _ = v117
                }
            }
            _ = v64
        }
        fmt.Println(v3, v6)
    }
    switch (-(v3)) {
    case 2147483647:
        var v118 uint64 = v1
        _ = v118
    default:
        var v119 uint8 = 255
        _ = v119
    }
    v6 = v6
    var v120 uint = 0
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    for fi121 := 0; fi121 < 15; fi121++ {
        if false {
            switch (fi121 + -9223372036854775808) {
            case 42:
                var v122 int32 = v5
                _ = v122
            case 391:
                fmt.Println(v2)
            default:
                v1 -= (v1 >> ((v120) & 63))
            }
            if ((((v1 != v1) && ((v1 & (^(v1))) <= v1)) && ((v4 >= v3) || true)) && true) {
                v5 -= (v5 - 138)
                v6 = 42
                if ((v120 | 4294967295) == (v120 / ((4294967295) | 1))) {
                    v5 += -1
                    var v123 uint8 = 255
                    _ = v123
                }
            } else {
                for fi124 := 0; fi124 < 6; fi124++ {
                    v125 := GenS15{F0: 42, F1: 4294967295}
                    switch "a" {
                    case "":
                        switch v3 {
                        case 9223372036854775807:
                            var v126 uint32 = 2
                            _ = v126
                        }
                    default:
                        v127 := GenS5{F0: v2}
                        _ = v127
                    }
                    _ = fi124
                    _ = v125
                }
            }
        }
        _ = fi121
    }
    fmt.Println("v1:", v1)
    fmt.Println("v120:", v120)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v120
    _ = v15
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
