package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int64; F1 int16 }
type GenS8 struct { F0 uint16; F1 uint32 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 bool; F1 int64 }
type GenS11 struct { F0 uint8; F1 int16 }
type GenS12 struct { F0 string; F1 int8; F2 int8 }
func genFn0(p0 int) string {
    var v6 float32 = -2.25
    for fi7 := 0; fi7 < 20; fi7++ {
        for fi8 := 0; fi8 < 6; fi8++ {
            v9 := GenS1{F0: (fi8 >= fi8), F1: 32767}
            v6 += v6
            var v10 string = "foo bar"
            _ = fi8
            _ = v10
            _ = v9
        }
        var v11 int64 = 42
        switch "a" {
        case "":
            for fi12 := 0; fi12 < 1; fi12++ {
                v11 = 1
                _ = fi12
            }
        }
        _ = fi7
        _ = v11
    }
    var v13 float64 = (3.14159 * 1e10)
    _ = p0
    _ = v13
    _ = v6
    return (("" + "a") + "a")
}
func genFn1(p0 int, p1 int, p2 int8) uint64 {
    for fi14 := 0; fi14 < 7; fi14++ {
        v15 := []int{(^(p1))}
        _ = fi14
        _ = v15
    }
    if ((3.14159 < ((-2.25 / 1e10) + 0.0)) || (-1 > 0)) {
        if false {
            switch "\n\t" {
            case "":
                if false {
                    p1 = (p0 * 0)
                    var v16 float32 = ((-2.25 - 0.1) / -2.25)
                    _ = v16
                } else {
                    var v17 uint64 = 42
                    switch ("\n\t" + "\n\t") {
                    case "\n\t":
                        if true {
                            var v18 int32 = 1
                            switch (p0 % ((p1) | 1)) {
                            case 0:
                                var v19 int16 = -414
                                _ = v19
                            case 32768:
                                v20 := GenS1{F0: true, F1: 824}
                                _ = v20
                            }
                            _ = v18
                        }
                    case "hello":
                        var v21 int32 = 0
                        _ = v21
                    case "foo bar":
                        v22 := GenS1{F0: ((!false) || (157 < 2)), F1: 65535}
                        _ = v22
                    default:
                        switch "" {
                        case "foo bar":
                            var v23 int64 = -55
                            _ = v23
                        case "hello":
                            var v24 uint64 = v17
                            _ = v24
                        }
                    }
                    _ = v17
                }
            }
            v25 := []int{4294967295, 128}
            _ = v25
        }
        v26 := GenS1{F0: true, F1: 65535}
        var v27 float32 = 3.14159
        _ = v26
        _ = v27
    }
    if true {
        var v28 uint64 = 2
        if false {
            for fi29 := 0; fi29 < 15; fi29++ {
                for fi30 := 0; fi30 < 11; fi30++ {
                    v31 := []int{fi30, 9223372036854775807, -998}
                    if len(v31) > 0 {
                        v31[0] = (p1 * p1)
                    }
                    switch "a" {
                    case "foo bar":
                        v32 := GenS1{F0: (!false), F1: 42}
                        _ = v32
                    case "\n\t":
                        if true {
                            if len(v31) > 0 {
                                v31[0] = (^(fi29))
                            }
                            p0 -= -176
                        } else {
                            p1 += 1
                        }
                    case "":
                        fmt.Println(p2)
                    default:
                        var v33 float64 = 0.0
                        p2 = -119
                        _ = v33
                    }
                    if false {
                        switch "foo bar" {
                        case "foo bar":
                            var v34 uint = uint(v28)
                            _ = v34
                        case "":
                            var v35 bool = (!(!false))
                            _ = v35
                        }
                    }
                    var v36 string = ""
                    var v37 string = ("foo bar" + (v36 + v36))
                    _ = fi30
                    _ = v31
                    _ = v36
                    _ = v37
                }
                var v38 int16 = -22
                v39 := []int{fi29, 127, p0}
                _ = fi29
                _ = v38
                _ = v39
            }
            var v40 int16 = int16(p2)
            _ = v40
        } else {
            v41 := GenS2{F0: -744, F1: 0}
            v28 -= 0
            _ = v41
        }
        _ = v28
    } else {
        v42 := GenS2{F0: 370, F1: 0}
        _ = v42
    }
    var v43 uint32 = 316
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v43
    return 18446744073709551615
}

func main() {
    var v1 bool = ((true && false) || (2147483647 < -972))
    var v2 uint32 = 942
    var v3 uint32 = (v2 + (v2 % (((v2 % ((v2) | 1))) | 1)))
    var v4 int = 32768
    var v5 int = v4
    v44 := []int{(v4 & v5)}
    for fi45 := 0; fi45 < 1; fi45++ {
        v5 += v4
        _ = fi45
    }
    v44 = append(v44, 898)
    var v46 bool = (42 != 0)
    if true {
        for fi47 := 0; fi47 < 19; fi47++ {
            switch fi47 {
            case 1:
                v48 := GenS3{F0: (0 > 236), F1: 2}
                _ = v48
            case 99:
                if v1 {
                    if true {
                        var v49 int32 = -1
                        switch "hello" {
                        case "":
                            v44 = append(v44, 577)
                        case "a":
                            switch fi47 {
                            case -782:
                                if v1 {
                                    v5 = v4
                                    for fi50 := 0; fi50 < 2; fi50++ {
                                        if len(v44) > 0 {
                                            v44[0] = (-(v4))
                                        }
                                        v51 := GenS4{F0: (v4 << ((4294967295) & 63))}
                                        v1 = (65 <= 128)
                                        _ = fi50
                                        _ = v51
                                    }
                                } else {
                                    var v52 bool = true
                                    var v53 int8 = 127
                                    v3 = 1
                                    _ = v52
                                    _ = v53
                                }
                            case -515:
                                v3 += v2
                            case -165:
                                var v54 int = (v5 - (fi47 + (fi47 * 268)))
                                _ = v54
                            default:
                                var v55 int = (fi47 & (fi47 * fi47))
                                _ = v55
                            }
                        }
                        if (579 < -1) {
                            if true {
                                for fi56 := 0; fi56 < 15; fi56++ {
                                    var v57 int32 = (v49 - v49)
                                    _ = fi56
                                    _ = v57
                                }
                                v5 = v5
                            }
                            v46 = false
                            switch (-(v4)) {
                            case 908:
                                var v58 uint16 = 93
                                _ = v58
                            case 4294967295:
                                var v59 uint16 = 1
                                _ = v59
                            case -9223372036854775808:
                                for fi60 := 0; fi60 < 20; fi60++ {
                                    if (4294967295 > 42) {
                                        switch (^(fi47)) {
                                        case 127:
                                            switch v4 {
                                            case -835:
                                                if false {
                                                    if true {
                                                        v61 := GenS2{F0: -32768, F1: (v2 + v3)}
                                                        _ = v61
                                                    }
                                                    v62 := GenS5{F0: v2}
                                                    if len(v44) >= 2 {
                                                        v44 = v44[1:len(v44)]
                                                    }
                                                    var v63 uint8 = 127
                                                    _ = v62
                                                    _ = v63
                                                } else {
                                                    v1 = (!("hello" <= (((((("\n\t" + ("\n\t" + "")) + "\n\t") + "foo bar") + "") + "hello") + (("\n\t" + "") + "foo bar"))))
                                                    if len(v44) > 0 {
                                                        v44[0] = v4
                                                    }
                                                    switch ("\n\t" + "") {
                                                    case "foo bar":
                                                        switch "hello" {
                                                        case "foo bar":
                                                            fmt.Println(v46, v5, v3)
                                                        case "":
                                                            fmt.Println(v5, fi60)
                                                        case "\n\t":
                                                            var v64 int = -2147483648
                                                            _ = v64
                                                        }
                                                    default:
                                                        switch (^(fi47)) {
                                                        case 2147483647:
                                                            v65 := GenS6{F0: (^(v3))}
                                                            _ = v65
                                                        case -2147483648:
                                                            if false {
                                                                v46 = (18446744073709551615 < 4294967295)
                                                                switch (v4 >> ((18446744073709551615) & 63)) {
                                                                case -883:
                                                                    var v66 uint32 = 1
                                                                    _ = v66
                                                                }
                                                                var v67 uint16 = 32767
                                                                _ = v67
                                                            } else {
                                                                fmt.Println(v2, v49)
                                                            }
                                                        default:
                                                            switch "hello" {
                                                            case "a":
                                                                _ = len(v44)
                                                            case "foo bar":
                                                                var v68 uint8 = 1
                                                                _ = v68
                                                            case "\n\t":
                                                                switch (v5 << ((255) & 63)) {
                                                                case -2147483648:
                                                                    var v69 bool = (4294967295 > (^(v2)))
                                                                    _ = v69
                                                                case 741:
                                                                    var v70 float64 = -2.25
                                                                    _ = v70
                                                                case -1:
                                                                    v3 += v3
                                                                }
                                                            }
                                                        }
                                                        if ((!false) || (118 == 325)) {
                                                            var v71 int = (v5 | 127)
                                                            _ = v71
                                                        }
                                                        fmt.Println(v46, v49)
                                                        _ = len(v44)
                                                        v49 += (v49 << ((4294967295) & 63))
                                                        v4 = (fi60 + 32768)
                                                        var v72 int8 = 100
                                                        switch "" {
                                                        case "a":
                                                            fmt.Println(fi60, fi47)
                                                        }
                                                        _ = v72
                                                    }
                                                }
                                            }
                                        default:
                                            v73 := GenS7{F0: 42, F1: 1}
                                            _ = v73
                                        }
                                        var v74 int32 = (^(v49))
                                        var v75 uint8 = 255
                                        if ((((-2.25 / float32(0.1)) + (((1e10 + ((1.5 * 1.5) + 0.1)) - 1.5) / float32((0.1 * 3.14159)))) == -2.25) || true) {
                                            v76 := GenS8{F0: 32767, F1: v2}
                                            _ = len(v44)
                                            _ = v76
                                        }
                                        _ = v74
                                        _ = v75
                                    } else {
                                        v1 = (1 <= 130)
                                        fmt.Println(fi47, v49)
                                        var v77 uint32 = v3
                                        _ = v77
                                    }
                                    v49 = v49
                                    var v78 int32 = v49
                                    _ = fi60
                                    _ = v78
                                }
                            }
                        } else {
                            v79 := GenS4{F0: (v5 ^ (fi47 | 32767))}
                            for fi80 := 0; fi80 < 18; fi80++ {
                                var v81 string = "foo bar"
                                _ = fi80
                                _ = v81
                            }
                            for fi82 := 0; fi82 < 9; fi82++ {
                                fmt.Println(fi82, v1)
                                if v1 {
                                    var v83 float32 = 3.14159
                                    if len(v44) >= 2 {
                                        v44 = v44[1:len(v44)]
                                    }
                                    switch "foo bar" {
                                    case "foo bar":
                                        v49 -= v49
                                    default:
                                        switch ("\n\t" + "hello") {
                                        case "":
                                            v84 := GenS3{F0: (!((("hello" + "") != "hello") && v46)), F1: v2}
                                            _ = v84
                                        case "foo bar":
                                            var v85 float32 = v83
                                            _ = v85
                                        }
                                    }
                                    _ = v83
                                } else {
                                    v46 = v1
                                }
                                _ = fi82
                            }
                            _ = v79
                        }
                        _ = v49
                    }
                    v86 := GenS9{F0: "a"}
                    var v87 uint32 = (v2 | (v3 << ((2) & 63)))
                    _ = v86
                    _ = v87
                } else {
                    for fi88 := 0; fi88 < 14; fi88++ {
                        var v89 uint = 892
                        if len(v44) >= 2 {
                            v44 = v44[1:len(v44)]
                        }
                        v90 := GenS10{F0: ((v89 * v89) == (v89 & 4294967295)), F1: 0}
                        _ = fi88
                        _ = v89
                        _ = v90
                    }
                }
            default:
                v2 = 0
            }
            fmt.Println(v1, v2, v1)
            _ = fi47
        }
        var v91 int = v4
        fmt.Println(v4)
        _ = v91
    }
    if false {
        v5 += v5
        if len(v44) >= 2 {
            v44 = v44[1:len(v44)]
        }
        for fi92 := 0; fi92 < 3; fi92++ {
            for fi93 := 0; fi93 < 9; fi93++ {
                switch "foo bar" {
                case "hello":
                    v4 = v5
                }
                _ = len(v44)
                _ = fi93
            }
            var v94 int = -2147483648
            v95 := GenS6{F0: v3}
            _ = fi92
            _ = v94
            _ = v95
        }
    }
    if v46 {
        var v96 int64 = 9223372036854775807
        var v97 uint8 = 229
        for fi98 := 0; fi98 < 8; fi98++ {
            switch (("foo bar" + "") + ("" + "")) {
            case "a":
                var v99 float32 = 1.5
                _ = v99
            case "":
                fmt.Println(v5, v3, v46)
            }
            var v100 uint16 = 662
            _ = fi98
            _ = v100
        }
        _ = v96
        _ = v97
    } else {
        switch "a" {
        case "hello":
            fmt.Println(v5, v3, v3)
        case "":
            v5 -= v5
        default:
            v4 = v5
        }
        var v101 int = 1
        _ = v101
    }
    v44 = append(v44, (v5 | (v5 - (v5 >> ((255) & 63)))))
    for fi102 := 0; fi102 < 15; fi102++ {
        switch v4 {
        case 4294967295:
            switch fi102 {
            case 4294967295:
                fmt.Println(v5)
            case -1:
                var v103 uint64 = 42
                _ = v103
            case 32768:
                v44 = append(v44, fi102)
            }
        case 517:
            fmt.Println(fi102, fi102)
        default:
            v4 = 127
        }
        _ = fi102
    }
    v46 = (uint8(v2) <= 255)
    switch 864 {
    case 32768:
        v104 := GenS11{F0: 146, F1: -1}
        _ = v104
    default:
        v105 := GenS12{F0: "\n\t", F1: 127, F2: -128}
        _ = v105
    }
    if ((v2 % (((^(v2))) | 1)) != 9) {
        v3 -= (v3 / ((v2) | 1))
    } else {
        for fi106 := 0; fi106 < 15; fi106++ {
            v3 = 0
            _ = fi106
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v46
    _ = v5
}
