package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 bool; F1 int }
type GenS5 struct { F0 int; F1 uint64 }
type GenS6 struct { F0 int32; F1 uint; F2 string }
type GenS7 struct { F0 uint16; F1 string; F2 string }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int8; F1 bool; F2 uint64 }
type GenS10 struct { F0 uint; F1 int64 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 uint16; F1 int8; F2 string }
type GenS13 struct { F0 string; F1 int32 }
type GenS14 struct { F0 uint8; F1 uint8; F2 int64 }
type GenS15 struct { F0 bool; F1 int8 }
type GenS16 struct { F0 int32; F1 int8; F2 int32 }
type GenS17 struct { F0 int32; F1 uint32; F2 int32 }
func genFn0(p0 int8, p1 uint) int16 {
    for fi4 := 0; fi4 < 1; fi4++ {
        v5 := GenS1{F0: 610}
        fmt.Println(p1, p0)
        v6 := []int{}
        _ = fi4
        _ = v5
        _ = v6
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    return -160
}

func main() {
    var v1 uint64 = 42
    var v2 int32 = 472
    var v3 uint8 = 127
    switch "" {
    case "hello":
        v3 -= (^(v3))
    }
    if false {
        switch "\n\t" {
        case "\n\t":
            v7 := GenS2{F0: 1, F1: 42}
            _ = v7
        }
        fmt.Println(v3, v1)
    }
    v8 := []int{32768, -9223372036854775808}
    v3 = 128
    v9 := GenS3{F0: 42, F1: "\n\t"}
    v8 = append(v8, 127)
    v10 := GenS2{F0: v3, F1: -128}
    fmt.Println(v3, v1)
    var v11 uint32 = 4294967295
    switch 0 {
    case -444:
        if ((929 > 255) || true) {
            if (!((v1 % ((2) | 1)) != (v1 + (v1 | 2)))) {
                for fi12 := 0; fi12 < 15; fi12++ {
                    v8 = append(v8, (fi12 / ((128) | 1)))
                    fmt.Println(v11, v3, v1)
                    _ = fi12
                }
            } else {
                var v13 int8 = -1
                for fi14 := 0; fi14 < 10; fi14++ {
                    var v15 uint64 = v1
                    if ((fi14 << ((290) & 63)) > (fi14 | fi14)) {
                        fmt.Println(v11, fi14, v13)
                    }
                    _ = fi14
                    _ = v15
                }
                _ = v13
            }
            for fi16 := 0; fi16 < 12; fi16++ {
                switch fi16 {
                case 42:
                    if (v1 > (v1 * v1)) {
                        switch (fi16 << ((503) & 63)) {
                        case 2147483647:
                            var v17 uint16 = 1
                            _ = v17
                        case -1:
                            for fi18 := 0; fi18 < 2; fi18++ {
                                var v19 uint8 = (v3 - v3)
                                if len(v8) > 0 {
                                    v8[0] = fi16
                                }
                                _ = fi18
                                _ = v19
                            }
                        case -884:
                            var v20 float64 = float64((1e10 + -2.25))
                            _ = v20
                        }
                        var v21 uint64 = uint64(v11)
                        if false {
                            var v22 int8 = 42
                            _ = v22
                        } else {
                            if (0 <= 1) {
                                if len(v8) > 0 {
                                    v8[0] = (fi16 << ((0) & 63))
                                }
                                var v23 int32 = (v2 + v2)
                                if false {
                                    var v24 int8 = -128
                                    fmt.Println(v3, v2, v3)
                                    _ = v24
                                } else {
                                    if (!(true && false)) {
                                        var v25 int32 = v2
                                        if (3.14159 < 3.14159) {
                                            switch -9223372036854775808 {
                                            case 9223372036854775807:
                                                switch "foo bar" {
                                                case "hello":
                                                    v25 = (v2 << ((637) & 63))
                                                case "":
                                                    v25 = (v25 - v25)
                                                case "a":
                                                    v26 := GenS2{F0: v3, F1: 82}
                                                    _ = v26
                                                default:
                                                    if (357 <= 0) {
                                                        var v27 uint8 = (v3 & (v3 >> ((472) & 63)))
                                                        var v28 string = ""
                                                        _ = v27
                                                        _ = v28
                                                    } else {
                                                        if len(v8) >= 2 {
                                                            v8 = v8[1:len(v8)]
                                                        }
                                                    }
                                                }
                                            case -2147483648:
                                                for fi29 := 0; fi29 < 6; fi29++ {
                                                    var v30 float32 = 0.1
                                                    v31 := GenS1{F0: 65535}
                                                    var v32 float32 = float32(1e10)
                                                    _ = fi29
                                                    _ = v30
                                                    _ = v31
                                                    _ = v32
                                                }
                                            }
                                            switch (fi16 + (fi16 << ((666) & 63))) {
                                            case 9223372036854775807:
                                                switch "hello" {
                                                case "a":
                                                    var v33 float64 = 0.1
                                                    _ = v33
                                                case "foo bar":
                                                    var v34 uint64 = 18446744073709551615
                                                    _ = v34
                                                case "":
                                                    v35 := GenS2{F0: v3, F1: -94}
                                                    _ = v35
                                                default:
                                                    v21 = v21
                                                }
                                            default:
                                                if true {
                                                    switch (fi16 | fi16) {
                                                    case 1:
                                                        fmt.Println(v23)
                                                    case -85:
                                                        var v36 string = ("a" + "foo bar")
                                                        _ = v36
                                                    }
                                                } else {
                                                    v11 += v11
                                                }
                                            }
                                            for fi37 := 0; fi37 < 12; fi37++ {
                                                if true {
                                                    v1 += v21
                                                }
                                                v11 += (^(v11))
                                                if len(v8) >= 2 {
                                                    v8 = v8[1:len(v8)]
                                                }
                                                fmt.Println(fi37)
                                                if (127 > 127) {
                                                    if (1 == -293) {
                                                        _ = len(v8)
                                                        if true {
                                                            if false {
                                                                fmt.Println(v2, fi16)
                                                                if false {
                                                                    if true {
                                                                        if (1 <= 42) {
                                                                            var v38 int32 = (v23 - 2147483647)
                                                                            _ = v38
                                                                        } else {
                                                                            var v39 bool = (!false)
                                                                            if len(v8) >= 2 {
                                                                                v8 = v8[1:len(v8)]
                                                                            }
                                                                            _ = v39
                                                                        }
                                                                        if true {
                                                                            v3 += v3
                                                                            v40 := GenS4{F0: false, F1: fi16}
                                                                            fmt.Println(v11)
                                                                            if (-9223372036854775808 != 9223372036854775807) {
                                                                                v41 := GenS5{F0: fi16, F1: 1}
                                                                                _ = v41
                                                                            } else {
                                                                                v11 = v11
                                                                                v42 := GenS5{F0: (fi16 >> ((65535) & 63)), F1: (v1 % ((v21) | 1))}
                                                                                var v43 int = 42
                                                                                _ = v42
                                                                                _ = v43
                                                                            }
                                                                            v11 = (v11 % ((1) | 1))
                                                                            if len(v8) > 0 {
                                                                                v8[0] = -5
                                                                            }
                                                                            _ = v40
                                                                        }
                                                                        v44 := GenS1{F0: 786}
                                                                        _ = v44
                                                                    } else {
                                                                        var v45 int16 = int16(v11)
                                                                        switch ("hello" + ("hello" + "foo bar")) {
                                                                        case "a":
                                                                            v25 = v23
                                                                        default:
                                                                            v46 := GenS6{F0: v2, F1: 255, F2: "\n\t"}
                                                                            _ = v46
                                                                        }
                                                                        _ = v45
                                                                    }
                                                                }
                                                            }
                                                            v47 := GenS7{F0: 1, F1: ("a" + (("\n\t" + "foo bar") + "\n\t")), F2: "hello"}
                                                            _ = v47
                                                        } else {
                                                            v48 := GenS8{F0: (((1 > 452) && true) && false)}
                                                            v23 -= v25
                                                            _ = v48
                                                        }
                                                    } else {
                                                        if len(v8) >= 2 {
                                                            v8 = v8[1:len(v8)]
                                                        }
                                                        if ((v23 | v23) <= (v23 % ((v25) | 1))) {
                                                            var v49 int32 = v23
                                                            if len(v8) > 0 {
                                                                v8[0] = 127
                                                            }
                                                            _ = v49
                                                        } else {
                                                            if true {
                                                                switch int(v11) {
                                                                case 505:
                                                                    if (418 <= 1000) {
                                                                        _ = len(v8)
                                                                        v50 := GenS4{F0: (false && true), F1: fi16}
                                                                        if len(v8) >= 2 {
                                                                            v8 = v8[1:len(v8)]
                                                                        }
                                                                        _ = v50
                                                                    } else {
                                                                        var v51 string = "\n\t"
                                                                        v25 = (v2 | v25)
                                                                        v52 := GenS9{F0: -78, F1: true, F2: (v21 << ((4294967295) & 63))}
                                                                        _ = v51
                                                                        _ = v52
                                                                    }
                                                                case -1:
                                                                    v21 += (^(v21))
                                                                case 4294967295:
                                                                    var v53 uint32 = v11
                                                                    _ = v53
                                                                default:
                                                                    if len(v8) > 0 {
                                                                        v8[0] = 127
                                                                    }
                                                                }
                                                                if len(v8) > 0 {
                                                                    v8[0] = 25
                                                                }
                                                                if false {
                                                                    var v54 int8 = -128
                                                                    if false {
                                                                        var v55 uint32 = 643
                                                                        if len(v8) > 0 {
                                                                            v8[0] = (fi16 / (((fi16 >> ((65535) & 63))) | 1))
                                                                        }
                                                                        fmt.Println(v54, v21, v1)
                                                                        fmt.Println(v3, v25, v3)
                                                                        if false {
                                                                            var v56 int8 = (v54 & v54)
                                                                            var v57 bool = false
                                                                            _ = v56
                                                                            _ = v57
                                                                        } else {
                                                                            v58 := GenS10{F0: 65535, F1: -1}
                                                                            _ = v58
                                                                        }
                                                                        _ = v55
                                                                    } else {
                                                                        v3 += (v3 + v3)
                                                                        var v59 int8 = 1
                                                                        if (!(0.1 != (1.5 / (1e10 - (0.0 / 1.0))))) {
                                                                            var v60 int16 = -32768
                                                                            v8 = append(v8, int(v3))
                                                                            if false {
                                                                                var v61 float64 = 3.14159
                                                                                _ = v61
                                                                            }
                                                                            _ = v60
                                                                        }
                                                                        v54 = (v54 | 100)
                                                                        _ = v59
                                                                    }
                                                                    _ = v54
                                                                } else {
                                                                    if false {
                                                                        v62 := GenS10{F0: 179, F1: 42}
                                                                        var v63 bool = true
                                                                        _ = v62
                                                                        _ = v63
                                                                    }
                                                                    var v64 float32 = (0.0 - ((float32(3.14159) - 0.1) + -2.25))
                                                                    v3 = (v3 >> ((18446744073709551615) & 63))
                                                                    _ = v64
                                                                }
                                                            } else {
                                                                fmt.Println(v2, fi16, v11)
                                                                if false {
                                                                    switch "" {
                                                                    case "":
                                                                        v25 = (v25 * v23)
                                                                    case "\n\t":
                                                                        var v65 uint16 = uint16(v11)
                                                                        _ = v65
                                                                    default:
                                                                        switch "" {
                                                                        case "":
                                                                            _ = len(v8)
                                                                        case "\n\t":
                                                                            if (v3 != v3) {
                                                                                switch (("" + "") + "foo bar") {
                                                                                case "hello":
                                                                                    var v66 int16 = 196
                                                                                    _ = v66
                                                                                case "":
                                                                                    fmt.Println(fi37, fi37)
                                                                                default:
                                                                                    v1 = (v1 & v1)
                                                                                }
                                                                                var v67 uint64 = (^(v1))
                                                                                _ = v67
                                                                            } else {
                                                                                if true {
                                                                                    if len(v8) >= 2 {
                                                                                        v8 = v8[1:len(v8)]
                                                                                    }
                                                                                } else {
                                                                                    v1 -= v1
                                                                                }
                                                                                switch (fi37 * (fi16 << ((4294967295) & 63))) {
                                                                                case 683:
                                                                                    fmt.Println(v11, v1)
                                                                                case 123:
                                                                                    if (v11 >= (v11 / ((v11) | 1))) {
                                                                                        if true {
                                                                                            if (1 <= 42) {
                                                                                                if ((fi16 - fi37) == (^(fi37))) {
                                                                                                    if (149 <= 2) {
                                                                                                        v68 := GenS11{F0: 796}
                                                                                                        if true {
                                                                                                            v21 = v21
                                                                                                            var v69 int64 = -833
                                                                                                            v23 = int32(v69)
                                                                                                            var v70 uint8 = (^(v3))
                                                                                                            _ = v69
                                                                                                            _ = v70
                                                                                                        }
                                                                                                        _ = v68
                                                                                                    }
                                                                                                    var v71 uint16 = 847
                                                                                                    v25 = -2147483648
                                                                                                    _ = v71
                                                                                                }
                                                                                                if true {
                                                                                                    if (1 <= -72) {
                                                                                                        var v72 int = (fi37 << ((65535) & 63))
                                                                                                        if (false || (0 > 255)) {
                                                                                                            _ = len(v8)
                                                                                                        }
                                                                                                        _ = len(v8)
                                                                                                        if false {
                                                                                                            var v73 float32 = 0.0
                                                                                                            _ = v73
                                                                                                        } else {
                                                                                                            v74 := GenS3{F0: uint(v11), F1: (("hello" + "a") + "foo bar")}
                                                                                                            var v75 string = "foo bar"
                                                                                                            v72 -= 32768
                                                                                                            _ = v74
                                                                                                            _ = v75
                                                                                                        }
                                                                                                        _ = v72
                                                                                                    } else {
                                                                                                        v25 += 0
                                                                                                        v25 = -1
                                                                                                    }
                                                                                                    fmt.Println(v11, v21, v23)
                                                                                                }
                                                                                                switch "a" {
                                                                                                case "hello":
                                                                                                    v2 = -36
                                                                                                case "a":
                                                                                                    if len(v8) >= 2 {
                                                                                                        v8 = v8[1:len(v8)]
                                                                                                    }
                                                                                                case "foo bar":
                                                                                                    if false {
                                                                                                        fmt.Println(v21, v21, v3)
                                                                                                        if len(v8) >= 2 {
                                                                                                            v8 = v8[1:len(v8)]
                                                                                                        }
                                                                                                        if (1 != 467) {
                                                                                                            if true {
                                                                                                                _ = len(v8)
                                                                                                                switch (fi16 ^ fi37) {
                                                                                                                case -425:
                                                                                                                    var v76 uint32 = (v11 << ((0) & 63))
                                                                                                                    _ = v76
                                                                                                                case -450:
                                                                                                                    v77 := GenS6{F0: (-(v23)), F1: 255, F2: ("a" + "")}
                                                                                                                    if (fi37 > (-(fi37))) {
                                                                                                                        v21 = (v21 % (((^(v1))) | 1))
                                                                                                                        if (!("a" >= ("hello" + "foo bar"))) {
                                                                                                                            _ = len(v8)
                                                                                                                            var v78 bool = (!false)
                                                                                                                            var v79 bool = ((0 != 127) && (!v78))
                                                                                                                            var v80 uint = 4294967295
                                                                                                                            v78 = v79
                                                                                                                            _ = v78
                                                                                                                            _ = v79
                                                                                                                            _ = v80
                                                                                                                        } else {
                                                                                                                            v81 := GenS9{F0: -128, F1: (!true), F2: v21}
                                                                                                                            v8 = append(v8, 2147483647)
                                                                                                                            _ = v81
                                                                                                                        }
                                                                                                                    }
                                                                                                                    if false {
                                                                                                                        v8 = append(v8, fi16)
                                                                                                                        v23 = v25
                                                                                                                    }
                                                                                                                    _ = v77
                                                                                                                case -9223372036854775808:
                                                                                                                    var v82 int64 = 0
                                                                                                                    _ = v82
                                                                                                                }
                                                                                                                v83 := GenS12{F0: 0, F1: -16, F2: "\n\t"}
                                                                                                                switch (fi37 ^ fi37) {
                                                                                                                case 937:
                                                                                                                    v1 += v21
                                                                                                                }
                                                                                                                _ = v83
                                                                                                            }
                                                                                                        } else {
                                                                                                            v25 += (v2 & -2147483648)
                                                                                                        }
                                                                                                    }
                                                                                                    v1 -= (^(v21))
                                                                                                    v84 := GenS7{F0: 65535, F1: "hello", F2: ""}
                                                                                                    _ = v84
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        if false {
                                                                                            v11 = uint32(v1)
                                                                                            var v85 uint8 = (^(v3))
                                                                                            _ = v85
                                                                                        } else {
                                                                                            switch (fi37 - 32768) {
                                                                                            case 4294967295:
                                                                                                switch ("hello" + "a") {
                                                                                                case "a":
                                                                                                    fmt.Println(v11)
                                                                                                default:
                                                                                                    if false {
                                                                                                        v23 = v25
                                                                                                        switch "" {
                                                                                                        case "foo bar":
                                                                                                            v8 = append(v8, (fi37 | -2147483648))
                                                                                                        }
                                                                                                        var v86 int64 = -626
                                                                                                        _ = len(v8)
                                                                                                        _ = v86
                                                                                                    } else {
                                                                                                        v87 := GenS2{F0: v3, F1: -68}
                                                                                                        switch "a" {
                                                                                                        case "hello":
                                                                                                            v3 += v3
                                                                                                            switch (("a" + "foo bar") + "\n\t") {
                                                                                                            case "hello":
                                                                                                                switch fi16 {
                                                                                                                case 224:
                                                                                                                    var v88 uint = 42
                                                                                                                    _ = v88
                                                                                                                case -2147483648:
                                                                                                                    v8 = append(v8, (fi16 ^ fi37))
                                                                                                                case 32767:
                                                                                                                    v25 += (v2 ^ v25)
                                                                                                                }
                                                                                                            case "\n\t":
                                                                                                                v23 = (^(v25))
                                                                                                            case "":
                                                                                                                var v89 int = fi37
                                                                                                                _ = v89
                                                                                                            }
                                                                                                        case "a":
                                                                                                            v25 = v2
                                                                                                        case "":
                                                                                                            var v90 int = (fi16 % ((-305) | 1))
                                                                                                            _ = v90
                                                                                                        default:
                                                                                                            v25 += (v23 ^ (v2 << ((916) & 63)))
                                                                                                        }
                                                                                                        _ = v87
                                                                                                    }
                                                                                                }
                                                                                            case 0:
                                                                                                fmt.Println(v3, v21, fi16)
                                                                                            }
                                                                                            v91 := GenS13{F0: "hello", F1: -2147483648}
                                                                                            switch (fi16 - fi16) {
                                                                                            case 9223372036854775807:
                                                                                                if true {
                                                                                                    if false {
                                                                                                        v3 -= 111
                                                                                                        _ = len(v8)
                                                                                                    }
                                                                                                } else {
                                                                                                    fmt.Println(v2, v1)
                                                                                                }
                                                                                            case -9223372036854775808:
                                                                                                if true {
                                                                                                    v3 = (v3 - v3)
                                                                                                } else {
                                                                                                    v3 += (v3 | (v3 + 127))
                                                                                                }
                                                                                            case 119:
                                                                                                v1 -= (v21 & (v1 >> ((42) & 63)))
                                                                                            }
                                                                                            var v92 bool = true
                                                                                            _ = v91
                                                                                            _ = v92
                                                                                        }
                                                                                        if ((float32(0.1) / 3.14159) == 1e10) {
                                                                                            var v93 uint16 = 351
                                                                                            fmt.Println(v1, fi16, v3)
                                                                                            _ = v93
                                                                                        }
                                                                                        var v94 int = fi37
                                                                                        _ = v94
                                                                                    }
                                                                                default:
                                                                                    v23 = v23
                                                                                }
                                                                                v11 = 65535
                                                                                if false {
                                                                                    if true {
                                                                                        if false {
                                                                                            var v95 int = (fi37 & -1)
                                                                                            _ = len(v8)
                                                                                            _ = v95
                                                                                        }
                                                                                        fmt.Println(v3, v3, v3)
                                                                                    } else {
                                                                                        _ = len(v8)
                                                                                    }
                                                                                    _ = len(v8)
                                                                                    v96 := GenS14{F0: (^(v3)), F1: uint8(v1), F2: -380}
                                                                                    _ = v96
                                                                                } else {
                                                                                    if len(v8) >= 2 {
                                                                                        v8 = v8[1:len(v8)]
                                                                                    }
                                                                                }
                                                                                v23 -= 1
                                                                            }
                                                                        }
                                                                    }
                                                                    if (!true) {
                                                                        var v97 float32 = -2.25
                                                                        fmt.Println(v23)
                                                                        fmt.Println(v1)
                                                                        switch ("hello" + ("a" + ("a" + "foo bar"))) {
                                                                        case "":
                                                                            v2 = v25
                                                                        case "\n\t":
                                                                            if len(v8) >= 2 {
                                                                                v8 = v8[1:len(v8)]
                                                                            }
                                                                        case "hello":
                                                                            if (fi37 == fi16) {
                                                                                var v98 uint8 = uint8(v11)
                                                                                switch "foo bar" {
                                                                                case "hello":
                                                                                    v99 := GenS15{F0: true, F1: 42}
                                                                                    _ = v99
                                                                                case "foo bar":
                                                                                    if true {
                                                                                        v100 := GenS16{F0: (-(v23)), F1: -1, F2: (v2 >> ((1) & 63))}
                                                                                        _ = v100
                                                                                    } else {
                                                                                        var v101 int32 = v23
                                                                                        _ = v101
                                                                                    }
                                                                                }
                                                                                if len(v8) > 0 {
                                                                                    v8[0] = fi37
                                                                                }
                                                                                _ = v98
                                                                            } else {
                                                                                var v102 int8 = 42
                                                                                _ = v102
                                                                            }
                                                                        default:
                                                                            var v103 float32 = (1.5 + 3.14159)
                                                                            _ = v103
                                                                        }
                                                                        _ = v97
                                                                    } else {
                                                                        var v104 float32 = -2.25
                                                                        switch "foo bar" {
                                                                        case "foo bar":
                                                                            var v105 uint32 = (v11 & uint32(v3))
                                                                            _ = v105
                                                                        case "hello":
                                                                            v1 += v21
                                                                        case "a":
                                                                            fmt.Println(v11, fi16, v104)
                                                                        default:
                                                                            fmt.Println(v21, v23, fi16)
                                                                        }
                                                                        var v106 int32 = v2
                                                                        _ = v104
                                                                        _ = v106
                                                                    }
                                                                } else {
                                                                    if len(v8) >= 2 {
                                                                        v8 = v8[1:len(v8)]
                                                                    }
                                                                }
                                                            }
                                                            var v107 string = (("foo bar" + "hello") + "foo bar")
                                                            var v108 int = (fi37 << ((4294967295) & 63))
                                                            _ = v107
                                                            _ = v108
                                                        }
                                                    }
                                                }
                                                _ = fi37
                                            }
                                        } else {
                                            for fi109 := 0; fi109 < 11; fi109++ {
                                                _ = len(v8)
                                                var v110 string = ("\n\t" + "hello")
                                                v25 = 42
                                                _ = fi109
                                                _ = v110
                                            }
                                            v111 := GenS4{F0: true, F1: (fi16 | 645)}
                                            _ = v111
                                        }
                                        if ((!((("" + "hello") <= "\n\t") && ((65535 == (v11 | v11)) && false))) && (!((-2.25 / 1e10) < 1e10))) {
                                            v2 += v25
                                            for fi112 := 0; fi112 < 12; fi112++ {
                                                if (float64(0.0) <= 1e10) {
                                                    fmt.Println(v21)
                                                    fmt.Println(v1)
                                                }
                                                v23 += (v23 >> ((2) & 63))
                                                var v113 uint64 = v1
                                                _ = fi112
                                                _ = v113
                                            }
                                        }
                                        _ = v25
                                    }
                                    var v114 bool = true
                                    _ = v114
                                }
                                _ = v23
                            }
                        }
                        _ = v21
                    } else {
                        fmt.Println(v3, v1, v11)
                        fmt.Println(v3, fi16, v11)
                        v115 := GenS16{F0: v2, F1: 46, F2: (^(v2))}
                        _ = v115
                    }
                case -1:
                    if ((fi16 << ((255) & 63)) == 128) {
                        if false {
                            for fi116 := 0; fi116 < 18; fi116++ {
                                fmt.Println(fi116, v3, fi116)
                                v1 = v1
                                switch "hello" {
                                case "foo bar":
                                    fmt.Println(v11, v1, v2)
                                case "a":
                                    v2 -= (v2 + v2)
                                case "":
                                    v1 = 42
                                    v117 := GenS1{F0: 1}
                                    fmt.Println(v2, v2, v1)
                                    _ = v117
                                }
                                _ = fi116
                            }
                        }
                    }
                case 32767:
                    var v118 uint32 = v11
                    _ = v118
                default:
                    var v119 uint16 = 522
                    _ = v119
                }
                _ = fi16
            }
            v120 := GenS17{F0: (-(v2)), F1: 2, F2: v2}
            _ = v120
        }
    case -2147483648:
        v3 -= (^(v3))
    case 128:
        switch "" {
        case "\n\t":
            var v121 int8 = -46
            _ = v121
        case "foo bar":
            var v122 uint32 = uint32(v1)
            _ = v122
        case "a":
            fmt.Println(v2, v11, v11)
        }
    default:
        v2 += (v2 % (((v2 | -398)) | 1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
