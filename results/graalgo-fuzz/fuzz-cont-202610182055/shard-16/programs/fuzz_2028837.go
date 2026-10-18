package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 int8; F2 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int; F1 int16; F2 uint16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint; F1 bool }
type GenS9 struct { F0 uint32 }

func main() {
    var v1 float32 = (float32(1.5) - 1e10)
    v2 := []string{("foo bar" + "hello"), "foo bar", ("hello" + "")}
    var v3 uint8 = 133
    var v4 int64 = -119
    var v5 bool = (!false)
    var v6 int64 = (v4 ^ (v4 >> ((0) & 63)))
    switch 128 {
    case -9223372036854775808:
        var v7 float64 = ((float64((v1 / float32(float64(float32(1.5))))) + 1e10) * -2.25)
        _ = v7
    }
    if v5 {
        if (!(v5 && v5)) {
            for fi8 := 0; fi8 < 7; fi8++ {
                if len(v2) > 0 {
                    v2[0] = ((("foo bar" + "a") + "") + ("\n\t" + (("hello" + ("" + ((("\n\t" + "foo bar") + "hello") + ""))) + ("a" + "foo bar"))))
                }
                _ = fi8
            }
            if false {
                fmt.Println(v3, v6, v5)
            } else {
                switch (("" + ("a" + (("" + ("a" + "\n\t")) + ("a" + "hello")))) + ("a" + "")) {
                case "hello":
                    switch 0 {
                    case -1:
                        var v9 int8 = -86
                        _ = v9
                    case -2147483648:
                        if ((!true) && v5) {
                            v5 = v5
                            for fi10 := 0; fi10 < 16; fi10++ {
                                fmt.Println(v4)
                                _ = len(v2)
                                fmt.Println(v6, v5)
                                _ = fi10
                            }
                            v3 = v3
                        }
                    default:
                        for fi11 := 0; fi11 < 18; fi11++ {
                            var v12 float64 = 0.0
                            v2 = append(v2, "hello")
                            v4 = (v6 << ((42) & 63))
                            _ = fi11
                            _ = v12
                        }
                    }
                default:
                    _ = len(v2)
                }
                if len(v2) > 0 {
                    v2[0] = ("" + "a")
                }
                var v13 uint32 = 2147483647
                _ = v13
            }
        }
        if len(v2) > 0 {
            v2[0] = (("a" + "a") + ("a" + ""))
        }
        if len(v2) > 0 {
            v2[0] = "a"
        }
    } else {
        var v14 uint64 = 233
        _ = v14
    }
    v15 := GenS1{F0: 0}
    switch "a" {
    case "":
        switch "a" {
        case "hello":
            for fi16 := 0; fi16 < 2; fi16++ {
                v17 := GenS1{F0: 626}
                fmt.Println(v4, v5)
                _ = fi16
                _ = v17
            }
        case "":
            v1 = float32(1.5)
        case "a":
            if v5 {
                if false {
                    for fi18 := 0; fi18 < 13; fi18++ {
                        switch (fi18 & int(v3)) {
                        case -683:
                            var v19 uint32 = 1
                            _ = v19
                        case -652:
                            var v20 int = (fi18 - (fi18 & (fi18 & (fi18 >> ((2) & 63)))))
                            _ = v20
                        case -437:
                            v4 += (v4 / ((int64(fi18)) | 1))
                        }
                        for fi21 := 0; fi21 < 11; fi21++ {
                            if (0.0 < float32(0.0)) {
                                fmt.Println(fi21, v5, v5)
                            }
                            if true {
                                var v22 int32 = 507
                                _ = v22
                            } else {
                                var v23 int8 = 42
                                v2 = append(v2, ((("" + "a") + ("\n\t" + "hello")) + "foo bar"))
                                fmt.Println(v6, v1, v4)
                                _ = v23
                            }
                            _ = fi21
                        }
                        _ = fi18
                    }
                    var v24 int32 = -768
                    var v25 uint64 = 18446744073709551615
                    _ = v24
                    _ = v25
                }
                var v26 int16 = int16(v6)
                v3 = v3
                _ = v26
            }
        }
    default:
        v3 -= (v3 & v3)
    }
    for fi27 := 0; fi27 < 17; fi27++ {
        if v5 {
            _ = len(v2)
            v4 -= v6
            v3 += 42
        } else {
            v28 := GenS2{F0: 0, F1: 104, F2: 168}
            fmt.Println(v5)
            _ = v28
        }
        if (4294967295 < 2) {
            v2 = append(v2, "foo bar")
        }
        _ = fi27
    }
    for fi29 := 0; fi29 < 4; fi29++ {
        v5 = v5
        _ = fi29
    }
    for fi30 := 0; fi30 < 12; fi30++ {
        if (2 < 723) {
            v4 += (v4 >> ((4294967295) & 63))
            switch ("" + "") {
            case "a":
                if (!(v3 < v3)) {
                    _ = len(v2)
                    v6 = v4
                    for fi31 := 0; fi31 < 4; fi31++ {
                        switch 243 {
                        case -78:
                            var v32 string = ("\n\t" + "a")
                            _ = v32
                        case 128:
                            v5 = ((!v5) || true)
                        case -1:
                            switch ("foo bar" + "foo bar") {
                            case "":
                                fmt.Println(v6)
                            case "\n\t":
                                v33 := GenS1{F0: 1}
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v34 := GenS2{F0: -1, F1: 100, F2: 255}
                                _ = v33
                                _ = v34
                            default:
                                fmt.Println(v1, fi30, v3)
                            }
                        }
                        v5 = v5
                        var v35 int32 = 1
                        _ = fi31
                        _ = v35
                    }
                } else {
                    for fi36 := 0; fi36 < 7; fi36++ {
                        switch "\n\t" {
                        case "foo bar":
                            switch ("foo bar" + "\n\t") {
                            case "":
                                if (1.5 > float64((0.1 * ((v1 / v1) + 3.14159)))) {
                                    if v5 {
                                        v37 := GenS3{F0: 215}
                                        v38 := GenS4{F0: fi30, F1: 1, F2: 65535}
                                        _ = v37
                                        _ = v38
                                    } else {
                                        v39 := GenS5{F0: 2}
                                        _ = v39
                                    }
                                    var v40 int = fi30
                                    if v5 {
                                        v2 = append(v2, ("foo bar" + "a"))
                                        fmt.Println(v4, v5)
                                    } else {
                                        var v41 uint32 = 175
                                        _ = v41
                                    }
                                    _ = v40
                                } else {
                                    v5 = v5
                                    var v42 int8 = 0
                                    switch (fi36 * fi36) {
                                    case 2147483647:
                                        v1 = v1
                                    default:
                                        if (!("foo bar" != ("foo bar" + "hello"))) {
                                            switch "hello" {
                                            case "\n\t":
                                                var v43 int16 = 472
                                                _ = v43
                                            case "hello":
                                                if len(v2) > 0 {
                                                    v2[0] = ""
                                                }
                                            case "a":
                                                _ = len(v2)
                                            default:
                                                fmt.Println(v3, fi30, v42)
                                            }
                                            fmt.Println(v1, v6, v3)
                                        } else {
                                            fmt.Println(v4, v1, v5)
                                        }
                                    }
                                    _ = v42
                                }
                            default:
                                var v44 uint16 = 255
                                fmt.Println(v1, v6)
                                var v45 int32 = -145
                                _ = v44
                                _ = v45
                            }
                        case "hello":
                            var v46 uint32 = 569
                            _ = v46
                        }
                        var v47 uint16 = 65535
                        _ = fi36
                        _ = v47
                    }
                    for fi48 := 0; fi48 < 14; fi48++ {
                        var v49 string = "hello"
                        v50 := GenS1{F0: 1}
                        if v5 {
                            fmt.Println(v3)
                        }
                        fmt.Println(v1, fi48, v5)
                        _ = fi48
                        _ = v49
                        _ = v50
                    }
                }
            case "foo bar":
                _ = len(v2)
            default:
                var v51 float64 = 0.1
                _ = v51
            }
        } else {
            if (v5 || (2 == 822)) {
                var v52 int = fi30
                _ = v52
            } else {
                v53 := GenS6{F0: 1}
                switch "\n\t" {
                case "a":
                    v3 -= v3
                case "foo bar":
                    fmt.Println(fi30, v3)
                case "\n\t":
                    v54 := GenS2{F0: 1, F1: 0, F2: uint16(v3)}
                    _ = v54
                default:
                    switch "a" {
                    case "foo bar":
                        if (!(((0.0 - 3.14159) + (-2.25 - 0.0)) >= -2.25)) {
                            fmt.Println(fi30)
                            var v55 int8 = 1
                            v2 = append(v2, ((("hello" + "a") + "") + (("foo bar" + (("\n\t" + "foo bar") + (("a" + "a") + ""))) + ("" + "\n\t"))))
                            _ = v55
                        } else {
                            if (-2147483648 >= -1) {
                                v5 = (2 <= 18446744073709551615)
                                fmt.Println(v3)
                            }
                        }
                    default:
                        v6 = v4
                    }
                }
                _ = v53
            }
        }
        _ = fi30
    }
    var v56 string = ("a" + ("a" + "\n\t"))
    if len(v2) > 0 {
        v2[0] = ""
    }
    v1 = (v1 + -2.25)
    for fi57 := 0; fi57 < 8; fi57++ {
        if false {
            for fi58 := 0; fi58 < 6; fi58++ {
                v1 -= (float32(float64(v1)) * v1)
                var v59 int64 = v6
                var v60 int = 42
                if (!v5) {
                    if true {
                        var v61 int32 = -837
                        _ = v61
                    } else {
                        fmt.Println(v6, v60)
                    }
                    if false {
                        fmt.Println(fi57, v4)
                        var v62 float32 = v1
                        _ = v62
                    }
                } else {
                    if v5 {
                        var v63 int16 = 42
                        v64 := GenS7{F0: 2}
                        if false {
                            v60 = (v60 + (fi57 ^ (fi57 + (v60 | (fi58 >> ((18446744073709551615) & 63))))))
                            v3 -= v3
                        }
                        _ = v63
                        _ = v64
                    } else {
                        if v5 {
                            var v65 uint32 = 743
                            v2 = append(v2, (v56 + v56))
                            _ = v65
                        }
                    }
                    v2 = append(v2, "")
                    fmt.Println(v6, v5, v1)
                }
                _ = fi58
                _ = v59
                _ = v60
            }
        } else {
            for fi66 := 0; fi66 < 6; fi66++ {
                switch v56 {
                case "":
                    var v67 int32 = -283
                    _ = v67
                case "hello":
                    var v68 uint32 = 2147483647
                    _ = v68
                case "a":
                    v56 = "\n\t"
                default:
                    var v69 uint16 = 342
                    _ = v69
                }
                v70 := GenS7{F0: 981}
                var v71 uint32 = 2147483647
                switch v56 {
                case "":
                    v72 := GenS8{F0: 255, F1: false}
                    var v73 int64 = (v4 ^ 965)
                    if ("a" != (v56 + (v56 + "hello"))) {
                        if v5 {
                            var v74 string = v56
                            var v75 int = (fi57 + 128)
                            _ = v74
                            _ = v75
                        }
                        switch fi57 {
                        case 1:
                            _ = len(v2)
                        case -2147483648:
                            v6 += v4
                        }
                        v5 = (false && (!((v6 % ((v4) | 1)) < 9223372036854775807)))
                    } else {
                        v71 -= v71
                        switch fi66 {
                        case 32768:
                            switch fi66 {
                            case -2147483648:
                                v2 = append(v2, (v56 + v56))
                            case 127:
                                var v76 int = 32767
                                switch (v76 >> ((4294967295) & 63)) {
                                case 127:
                                    var v77 int8 = 1
                                    _ = v77
                                }
                                fmt.Println(fi66, v76)
                                _ = v76
                            case 916:
                                v78 := GenS8{F0: 18446744073709551615, F1: (332 <= 18446744073709551615)}
                                _ = v78
                            default:
                                fmt.Println(v5, v6)
                            }
                        }
                        v5 = true
                    }
                    _ = v72
                    _ = v73
                case "a":
                    var v79 uint32 = v71
                    _ = v79
                case "hello":
                    v5 = v5
                }
                _ = fi66
                _ = v70
                _ = v71
            }
            v56 = (v56 + "hello")
        }
        for fi80 := 0; fi80 < 16; fi80++ {
            v1 = ((float32(3.14159) * v1) - 3.14159)
            if v5 {
                if (((v4 ^ (v4 >> ((2) & 63))) <= (v4 << ((497) & 63))) || v5) {
                    var v81 int64 = 98
                    var v82 int = -757
                    v6 += v6
                    if (255 != 1) {
                        if v5 {
                            v4 -= (^(v81))
                            fmt.Println(v4)
                        }
                        switch v56 {
                        case "foo bar":
                            var v83 bool = false
                            _ = v83
                        }
                    }
                    switch ("" + v56) {
                    case "a":
                        if (v81 > 9223372036854775807) {
                            v84 := GenS4{F0: (v82 >> ((255) & 63)), F1: -114, F2: 2}
                            fmt.Println(v82, v1, fi80)
                            if false {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v4 += int64(fi80)
                                switch (("foo bar" + (v56 + "a")) + v56) {
                                case "\n\t":
                                    v6 = v4
                                }
                                if v5 {
                                    if ((v3 != v3) && ((v5 && (true || true)) || false)) {
                                        v2 = append(v2, "")
                                        switch (v82 * (v82 >> ((65535) & 63))) {
                                        case -1:
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        case 4294967295:
                                            v85 := GenS6{F0: 42}
                                            _ = v85
                                        case -199:
                                            if (((!v5) || (476 >= 2)) || (v6 > -1)) {
                                                var v86 float32 = float32(((3.14159 * float64(((v1 + 3.14159) + -2.25))) - ((0.0 * (-2.25 * -2.25)) * ((0.0 + (3.14159 + 0.0)) - 3.14159))))
                                                _ = v86
                                            } else {
                                                if (42 == -32768) {
                                                    v87 := GenS4{F0: v82, F1: 1000, F2: 32767}
                                                    v3 += (v3 ^ 122)
                                                    _ = v87
                                                } else {
                                                    var v88 int64 = v4
                                                    _ = v88
                                                }
                                                fmt.Println(fi57, v56, v56)
                                            }
                                        }
                                    }
                                    var v89 uint32 = 0
                                    _ = v89
                                }
                            } else {
                                v5 = v5
                                fmt.Println(fi80, v3, v5)
                                var v90 uint8 = (v3 & v3)
                                _ = v90
                            }
                            v6 -= 42
                            _ = len(v2)
                            _ = len(v2)
                            _ = v84
                        } else {
                            if false {
                                v56 = v56
                                v6 = (v81 | (v4 | 1))
                            } else {
                                if len(v2) > 0 {
                                    v2[0] = v56
                                }
                            }
                            v82 = v82
                            v5 = false
                        }
                    case "":
                        fmt.Println(v56, v82)
                    case "\n\t":
                        var v91 uint16 = 997
                        _ = v91
                    default:
                        fmt.Println(v5, v56)
                    }
                    if v5 {
                        switch "" {
                        case "hello":
                            switch "\n\t" {
                            case "a":
                                var v92 uint64 = 1
                                _ = v92
                            }
                        default:
                            v4 -= (v6 % (((v4 & v6)) | 1))
                        }
                        var v93 int16 = 1
                        v3 -= (^(v3))
                        _ = v93
                    } else {
                        var v94 int32 = 42
                        var v95 uint8 = (v3 + (v3 ^ v3))
                        switch (v82 / (((fi80 ^ 9223372036854775807)) | 1)) {
                        case -584:
                            if len(v2) > 0 {
                                v2[0] = (v56 + v56)
                            }
                        case 42:
                            if v5 {
                                v3 = 126
                            } else {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                if v5 {
                                    v5 = true
                                    v96 := GenS9{F0: 305}
                                    var v97 float64 = (0.0 - (float64((v1 / v1)) * 1.5))
                                    var v98 int8 = -91
                                    _ = v96
                                    _ = v97
                                    _ = v98
                                } else {
                                    switch v56 {
                                    case "":
                                        if (!false) {
                                            v2 = append(v2, "foo bar")
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            v94 += (-(v94))
                                        }
                                        fmt.Println(fi57, v82)
                                        v95 = (^(v95))
                                    case "\n\t":
                                        v94 -= (v94 ^ v94)
                                    case "foo bar":
                                        var v99 int32 = (v94 >> ((18446744073709551615) & 63))
                                        _ = v99
                                    default:
                                        v6 = (v6 - (v81 + -9223372036854775808))
                                    }
                                }
                                v100 := GenS7{F0: 18446744073709551615}
                                _ = v100
                            }
                        case 859:
                            v95 = 250
                        default:
                            v95 = v3
                        }
                        _ = v94
                        _ = v95
                    }
                    fmt.Println(v5)
                    var v101 uint64 = 18446744073709551615
                    _ = v101
                    _ = v81
                    _ = v82
                }
                fmt.Println(v56, fi80, v4)
                v6 -= v4
                var v102 int8 = 42
                if len(v2) > 0 {
                    v2[0] = v56
                }
                _ = v102
            }
            _ = fi80
        }
        fmt.Println(v1, v1)
        _ = fi57
    }
    var v103 int16 = 368
    fmt.Println(v103, v5)
    switch v56 {
    case "foo bar":
        v103 = (v103 & v103)
    }
    if v5 {
        v2 = append(v2, ((v56 + (v56 + "\n\t")) + v56))
    }
    var v104 int64 = (v6 & 42)
    v104 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v103:", v103)
    fmt.Println("v104:", v104)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v103
    _ = v104
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v6
}
