package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int32 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 string; F1 uint16; F2 uint64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint8; F1 int16; F2 bool }
type GenS7 struct { F0 uint16; F1 int16 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint8; F1 int64; F2 uint64 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 int32; F1 uint64 }
type GenS12 struct { F0 int32; F1 bool }
type GenS13 struct { F0 bool; F1 uint8 }
type GenS14 struct { F0 string }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 uint64; F1 int; F2 uint64 }
func genFn0(p0 int16) int16 {
    v6 := GenS1{F0: 42, F1: 604, F2: -1}
    if false {
        v7 := []int{1, 128}
        _ = v7
    }
    v8 := GenS2{F0: "hello", F1: true}
    _ = p0
    _ = v6
    _ = v8
    return 32767
}

func main() {
    v1 := []string{}
    var v2 int8 = 0
    var v3 uint32 = 65535
    v4 := make(map[int]bool)
    var v5 float64 = (0.0 + ((float64((-2.25 - float32(1.5))) - 0.0) * ((0.1 * 3.14159) - 0.0)))
    for fi9 := 0; fi9 < 8; fi9++ {
        for fi10 := 0; fi10 < 19; fi10++ {
            fmt.Println(fi10, v2)
            v11 := v4[2147483647]
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    switch "hello" {
    case "foo bar":
        for fi12 := 0; fi12 < 7; fi12++ {
            var v13 uint8 = 2
            fmt.Println(fi12)
            var v14 int16 = -692
            _ = fi12
            _ = v13
            _ = v14
        }
    case "hello":
        for fi15 := 0; fi15 < 11; fi15++ {
            v16 := GenS3{F0: 894, F1: 255}
            var v17 int16 = 42
            switch fi15 {
            case -2147483648:
                v17 = (-(v17))
            }
            _ = fi15
            _ = v16
            _ = v17
        }
    default:
        var v18 float64 = v5
        _ = v18
    }
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    if (0 == int8(v3)) {
        for fi19 := 0; fi19 < 9; fi19++ {
            for fi20 := 0; fi20 < 5; fi20++ {
                v2 -= -45
                _ = fi20
            }
            if (0 == 42) {
                if (!true) {
                    if (1 >= 431) {
                        v21 := v4[9223372036854775807]
                        _ = v21
                    } else {
                        v22 := GenS4{F0: "\n\t", F1: 1, F2: 0}
                        v23 := GenS3{F0: 0, F1: 65535}
                        _ = v22
                        _ = v23
                    }
                    v24 := GenS4{F0: "a", F1: 688, F2: 87}
                    _ = v24
                } else {
                    var v25 int16 = -860
                    _ = v25
                }
            }
            _ = fi19
        }
        for fi26 := 0; fi26 < 9; fi26++ {
            v1 = append(v1, "")
            _ = fi26
        }
        v3 = 65535
    }
    if (1 > 0) {
        var v27 uint8 = 0
        _ = v27
    }
    fmt.Println(v2, v5)
    if (255 >= 1) {
        for fi28 := 0; fi28 < 16; fi28++ {
            for fi29 := 0; fi29 < 4; fi29++ {
                switch ("hello" + "\n\t") {
                case "foo bar":
                    switch fi28 {
                    case -1:
                        if true {
                            switch 4294967295 {
                            case 2147483647:
                                var v30 int16 = -1
                                _ = v30
                            case 32767:
                                if (1000 >= 4294967295) {
                                    switch (("a" + ("a" + "")) + "\n\t") {
                                    case "\n\t":
                                        var v31 float32 = 1e10
                                        var v32 int8 = 95
                                        v33 := GenS5{F0: fi28}
                                        _ = v31
                                        _ = v32
                                        _ = v33
                                    default:
                                        v34 := GenS6{F0: 126, F1: 1, F2: true}
                                        _ = v34
                                    }
                                    if true {
                                        fmt.Println(fi28, fi29, fi28)
                                        if ((v2 | v2) > (-(v2))) {
                                            v3 -= (v3 - (^(v3)))
                                            v2 += v2
                                            v35 := GenS4{F0: "\n\t", F1: 2, F2: 2}
                                            switch "foo bar" {
                                            case "a":
                                                fmt.Println(v3, fi28)
                                            case "\n\t":
                                                if (false || true) {
                                                    v3 = v3
                                                    var v36 uint8 = 42
                                                    v36 -= v36
                                                    _ = v36
                                                }
                                            }
                                            _ = v35
                                        } else {
                                            fmt.Println(v3, v2, v3)
                                        }
                                        v37 := GenS7{F0: 65535, F1: 443}
                                        var v38 uint32 = 2
                                        var v39 uint16 = 1
                                        _ = v37
                                        _ = v38
                                        _ = v39
                                    }
                                } else {
                                    v40 := GenS1{F0: 744, F1: (v3 - 0), F2: -416}
                                    _ = v40
                                }
                            }
                            v41 := GenS6{F0: 252, F1: 1000, F2: (v3 < v3)}
                            fmt.Println(fi29, fi29, v2)
                            v42 := GenS8{F0: 127}
                            _ = v41
                            _ = v42
                        } else {
                            var v43 string = ""
                            v43 = v43
                            _ = v43
                        }
                    case 925:
                        v44 := GenS9{F0: uint8(v3), F1: 829, F2: 416}
                        _ = v44
                    }
                case "":
                    var v45 int64 = 890
                    _ = v45
                default:
                    fmt.Println(fi29, fi29, fi29)
                }
                var v46 uint64 = 18446744073709551615
                v47 := v4[(fi29 ^ -932)]
                _ = fi29
                _ = v46
                _ = v47
            }
            v5 += float64(-2.25)
            switch -2147483648 {
            case -1:
                v3 -= v3
            case 4294967295:
                v3 += (v3 + v3)
            case 0:
                if false {
                    v4[fi28] = (!(0 > 1000))
                    v1 = append(v1, ("hello" + "hello"))
                    v3 = v3
                }
            default:
                var v48 int = int(v3)
                _ = v48
            }
            _ = fi28
        }
        for fi49 := 0; fi49 < 5; fi49++ {
            if (9223372036854775807 != 1) {
                var v50 int16 = -1
                _ = v50
            } else {
                if true {
                    v51 := GenS8{F0: 2}
                    _ = v51
                }
                v5 += (0.0 / (v5 * float64(-2.25)))
                if (42 > 18446744073709551615) {
                    var v52 uint32 = (v3 / ((248) | 1))
                    if true {
                        switch fi49 {
                        case 1:
                            v52 -= v52
                        case 127:
                            var v53 float64 = 3.14159
                            _ = v53
                        default:
                            v54 := v4[(fi49 << ((65535) & 63))]
                            _ = v54
                        }
                        switch 9223372036854775807 {
                        case 32768:
                            switch "hello" {
                            case "a":
                                if (!false) {
                                    var v55 uint32 = (v52 ^ (v3 << ((699) & 63)))
                                    for fi56 := 0; fi56 < 9; fi56++ {
                                        fmt.Println(v52, v5)
                                        switch (fi49 ^ (fi49 - (fi56 % ((853) | 1)))) {
                                        case 313:
                                            v4[fi49] = (false || false)
                                        case -9223372036854775808:
                                            switch (fi56 & (fi56 - fi49)) {
                                            case 1:
                                                if len(v1) > 0 {
                                                    v1[0] = "hello"
                                                }
                                            case 127:
                                                v55 = v52
                                            case 32767:
                                                if true {
                                                    fmt.Println(v2)
                                                    if true {
                                                        fmt.Println(fi56)
                                                        fmt.Println(fi49)
                                                    }
                                                } else {
                                                    var v57 float32 = (0.1 / 3.14159)
                                                    _ = v57
                                                }
                                            }
                                        }
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        fmt.Println(v2, fi49)
                                        _ = fi56
                                    }
                                    switch "\n\t" {
                                    case "":
                                        v52 += (v52 | (v3 ^ 854))
                                    }
                                    _ = v55
                                }
                            case "hello":
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            default:
                                if (float64(1.5) < v5) {
                                    v5 = float64(3.14159)
                                }
                            }
                        case 2147483647:
                            for fi58 := 0; fi58 < 16; fi58++ {
                                v5 = v5
                                var v59 uint16 = 1
                                var v60 uint32 = 65535
                                var v61 int8 = v2
                                fmt.Println(fi49, v61, fi49)
                                _ = fi58
                                _ = v59
                                _ = v60
                                _ = v61
                            }
                        }
                        for fi62 := 0; fi62 < 11; fi62++ {
                            v63 := GenS10{F0: fi49}
                            v64 := GenS5{F0: fi49}
                            _ = fi62
                            _ = v63
                            _ = v64
                        }
                    }
                    _ = v52
                }
            }
            if false {
                v65 := GenS11{F0: 877, F1: 4294967295}
                if true {
                    if false {
                        fmt.Println(fi49, v5, fi49)
                        var v66 uint32 = (v3 / (((v3 / (((v3 | v3)) | 1))) | 1))
                        _ = v66
                    }
                    for fi67 := 0; fi67 < 6; fi67++ {
                        if (v3 == (v3 % (((v3 % ((4294967295) | 1))) | 1))) {
                            if true {
                                fmt.Println(v2, v2)
                                var v68 bool = true
                                if v68 {
                                    var v69 string = ("" + "hello")
                                    switch (^(fi67)) {
                                    case -1:
                                        var v70 uint8 = 0
                                        _ = v70
                                    case 0:
                                        if v68 {
                                            v5 += 1e10
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        }
                                    }
                                    _ = v69
                                } else {
                                    if len(v1) > 0 {
                                        v1[0] = "a"
                                    }
                                    fmt.Println(fi67)
                                }
                                _ = v68
                            } else {
                                if false {
                                    var v71 uint32 = v3
                                    var v72 float32 = 3.14159
                                    _ = v71
                                    _ = v72
                                } else {
                                    v73 := v4[1]
                                    _ = v73
                                }
                                if len(v1) > 0 {
                                    v1[0] = ("\n\t" + "")
                                }
                                switch "foo bar" {
                                case "\n\t":
                                    delete(v4, fi49)
                                    var v74 int32 = -2147483648
                                    v3 -= (v3 ^ (v3 | 2))
                                    fmt.Println(fi49, v74, fi67)
                                    var v75 uint32 = v3
                                    _ = v74
                                    _ = v75
                                case "foo bar":
                                    if len(v1) > 0 {
                                        v1[0] = "a"
                                    }
                                case "a":
                                    var v76 uint32 = 42
                                    if (127 > 0) {
                                        var v77 int16 = -942
                                        _ = v77
                                    }
                                    _ = v76
                                }
                            }
                            v4[-712] = (42 >= 18446744073709551615)
                            v3 += (v3 / (((v3 | (^(v3)))) | 1))
                            if true {
                                if (0.0 > float32(((0.1 - float64(1e10)) * v5))) {
                                    v2 = (v2 ^ v2)
                                    var v78 uint32 = 1
                                    _ = v78
                                } else {
                                    v79 := GenS9{F0: 2, F1: 42, F2: 18446744073709551615}
                                    v2 += (v2 << ((42) & 63))
                                    _ = v79
                                }
                                v80 := GenS12{F0: 42, F1: (65535 != v3)}
                                _ = v80
                            }
                            if true {
                                var v81 uint32 = v3
                                _ = v81
                            }
                            switch (fi49 + 32768) {
                            case -807:
                                delete(v4, fi49)
                            default:
                                switch "hello" {
                                case "hello":
                                    v2 -= -22
                                case "":
                                    v3 = (^(v3))
                                case "\n\t":
                                    if ((float32(-2.25) / (float32(-2.25) * 1.5)) <= (-2.25 - 1e10)) {
                                        var v82 float32 = 1.5
                                        var v83 uint8 = 1
                                        v84 := GenS13{F0: false, F1: 48}
                                        _ = v82
                                        _ = v83
                                        _ = v84
                                    } else {
                                        v3 += v3
                                        _ = len(v1)
                                        var v85 uint64 = 18446744073709551615
                                        _ = v85
                                    }
                                }
                            }
                            fmt.Println(v2, v3, v5)
                        }
                        _ = fi67
                    }
                }
                _ = v65
            }
            _ = fi49
        }
    } else {
        v86 := GenS11{F0: 1, F1: 952}
        v2 = v2
        var v87 int64 = 1
        _ = v86
        _ = v87
    }
    if false {
        fmt.Println(v2, v5)
        var v88 uint16 = 65535
        switch ("\n\t" + "hello") {
        case "":
            v5 = (v5 / v5)
        case "\n\t":
            v2 = 120
        }
        _ = v88
    } else {
        _ = len(v1)
        if (!true) {
            v89 := GenS14{F0: "a"}
            _ = v89
        } else {
            v3 = v3
            for fi90 := 0; fi90 < 5; fi90++ {
                v91 := GenS15{F0: -1}
                _ = fi90
                _ = v91
            }
        }
    }
    fmt.Println(v2, v5, v2)
    v92 := GenS13{F0: false, F1: 46}
    for fi93 := 0; fi93 < 7; fi93++ {
        var v94 string = ("" + "a")
        _ = fi93
        _ = v94
    }
    v2 += v2
    for fi95 := 0; fi95 < 11; fi95++ {
        switch fi95 {
        case 9223372036854775807:
            if (42 > 0) {
                var v96 int64 = 606
                for fi97 := 0; fi97 < 2; fi97++ {
                    var v98 int8 = -128
                    v99 := v4[fi95]
                    _ = fi97
                    _ = v98
                    _ = v99
                }
                _ = v96
            } else {
                for fi100 := 0; fi100 < 7; fi100++ {
                    v1 = append(v1, "")
                    _ = fi100
                }
                v5 = 0.0
            }
        case 127:
            v3 -= v3
        default:
            fmt.Println(v2, v3)
        }
        _ = fi95
    }
    fmt.Println(v5)
    if (!(v5 < (1e10 + v5))) {
        _ = len(v1)
    } else {
        if (32767 == -1) {
            if false {
                v101 := GenS11{F0: -1, F1: 1}
                _ = v101
            } else {
                switch ("a" + "\n\t") {
                case "\n\t":
                    var v102 float64 = v5
                    _ = v102
                default:
                    v103 := GenS16{F0: 850, F1: 32768, F2: 1}
                    _ = v103
                }
                switch -1 {
                case 4294967295:
                    v5 = float64(((-2.25 * (0.0 * -2.25)) / 0.1))
                case 128:
                    v3 += v3
                case -9223372036854775808:
                    v104 := v4[127]
                    _ = v104
                default:
                    var v105 uint = 18446744073709551615
                    _ = v105
                }
            }
            if false {
                fmt.Println(v5, v3)
            }
        }
    }
    if false {
        v3 += (v3 - (v3 >> ((999) & 63)))
        for fi106 := 0; fi106 < 11; fi106++ {
            var v107 int64 = -1
            _ = fi106
            _ = v107
        }
    } else {
        v2 -= v2
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v92
}
