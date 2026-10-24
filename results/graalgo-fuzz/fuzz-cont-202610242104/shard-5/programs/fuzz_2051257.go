package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint32; F1 uint32 }
type GenS6 struct { F0 int; F1 bool }
type GenS7 struct { F0 uint16; F1 uint64; F2 uint64 }
type GenS8 struct { F0 uint; F1 uint16; F2 int8 }
type GenS9 struct { F0 uint8; F1 uint64; F2 bool }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 bool; F1 uint8; F2 int8 }
type GenS13 struct { F0 uint64; F1 int32; F2 int64 }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 int; F1 uint; F2 int }
type GenS16 struct { F0 int; F1 int32; F2 int }
type GenS17 struct { F0 int32; F1 bool }
type GenS18 struct { F0 uint16; F1 int64 }
type GenS19 struct { F0 int16; F1 int }
type GenS20 struct { F0 string }
type GenS21 struct { F0 int64 }
type GenS22 struct { F0 int64; F1 int32; F2 int16 }
type GenS23 struct { F0 uint32 }
type GenS24 struct { F0 int; F1 uint }
type GenS25 struct { F0 int8; F1 int }
type GenS26 struct { F0 int16; F1 uint8; F2 uint16 }
type GenS27 struct { F0 uint8; F1 int16; F2 bool }
type GenS28 struct { F0 string; F1 int }
type GenS29 struct { F0 uint64; F1 int; F2 uint32 }

func main() {
    var v1 float32 = 3.14159
    var v2 uint16 = 65535
    v3 := make(map[string]int)
    var v4 uint8 = 47
    var v5 int64 = 388
    v6 := []uint32{}
    v7 := GenS1{F0: (^(v2)), F1: 42}
    v3["\n\t"] = -2147483648
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 bool = true
        v10 := GenS1{F0: 32767, F1: 100}
        fmt.Println(v4)
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        fmt.Println(v2, fi11)
        var v12 uint32 = 42
        if len(v6) > 0 {
            v6[0] = 65535
        }
        _ = fi11
        _ = v12
    }
    fmt.Println(v4, v4, v4)
    v2 = 0
    v4 = 42
    v2 += v2
    v1 += (v1 - float32(1.5))
    var v13 uint64 = 42
    v14 := GenS2{F0: 42}
    if false {
        v15 := v3["foo bar"]
        switch (v15 * 1) {
        case 122:
            v16 := GenS2{F0: (v5 | 0)}
            _ = v16
        }
        if (42 <= 853) {
            v15 = (v15 + 932)
        }
        _ = v15
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        var v18 float32 = v1
        for fi19 := 0; fi19 < 9; fi19++ {
            var v20 int8 = -128
            var v21 float64 = 3.14159
            switch ((("hello" + ("a" + ((("hello" + "") + "") + "\n\t"))) + ("a" + ("a" + ""))) + "hello") {
            case "\n\t":
                v20 = 0
            case "foo bar":
                if true {
                    v22 := v3["\n\t"]
                    _ = v22
                } else {
                    fmt.Println(v1, v1, fi17)
                    v1 -= (float32(3.14159) / float32(((3.14159 + ((v21 + float64((0.1 / v1))) / -2.25)) / (v21 * 0.1))))
                    v23 := GenS1{F0: (v2 | v2), F1: v20}
                    _ = v23
                }
            case "hello":
                v24 := GenS3{F0: true, F1: 32767}
                _ = v24
            default:
                v13 -= v13
            }
            var v25 float64 = v21
            var v26 float32 = (3.14159 + 0.1)
            v27 := GenS4{F0: 42}
            _ = fi19
            _ = v20
            _ = v21
            _ = v25
            _ = v26
            _ = v27
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi17
        _ = v18
    }
    fmt.Println(v4, v1, v13)
    fmt.Println(v13)
    v4 += 255
    v2 = (^(v2))
    if (!true) {
        var v28 float64 = 3.14159
        for fi29 := 0; fi29 < 4; fi29++ {
            for fi30 := 0; fi30 < 8; fi30++ {
                v31 := GenS5{F0: 765, F1: 2147483647}
                switch 127 {
                case 327:
                    switch fi29 {
                    case 32767:
                        switch (fi30 & (fi30 - (fi30 % (((-(fi30))) | 1)))) {
                        case 128:
                            v4 = v4
                        case -637:
                            if true {
                                var v32 int16 = 740
                                var v33 bool = (v28 <= float64(((float32(v28) / 1e10) / 1.5)))
                                switch (fi30 * int(v5)) {
                                case -1:
                                    if (!(-128 < 1)) {
                                        var v34 uint = 42
                                        v3["foo bar"] = (fi29 / (((fi30 * int(v4))) | 1))
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        _ = v34
                                    } else {
                                        v35 := v3["hello"]
                                        _ = v35
                                    }
                                case 4294967295:
                                    switch 1 {
                                    case -2147483648:
                                        if (false || false) {
                                            if v33 {
                                                _ = len(v6)
                                                if (("hello" + "\n\t") < "foo bar") {
                                                    var v36 string = "a"
                                                    v32 += (v32 - v32)
                                                    _ = v36
                                                } else {
                                                    if (0.0 < (v1 + v1)) {
                                                        var v37 string = "\n\t"
                                                        v4 -= (v4 << ((4294967295) & 63))
                                                        _ = v37
                                                    } else {
                                                        var v38 uint32 = 0
                                                        if v33 {
                                                            var v39 int64 = (v5 ^ 1)
                                                            fmt.Println(v1)
                                                            _ = v39
                                                        } else {
                                                            if true {
                                                                var v40 int = fi30
                                                                switch ("\n\t" + "foo bar") {
                                                                case "hello":
                                                                    switch "\n\t" {
                                                                    case "foo bar":
                                                                        v41 := GenS3{F0: true, F1: 0}
                                                                        _ = v41
                                                                    case "\n\t":
                                                                        v42 := GenS3{F0: v33, F1: (^(v2))}
                                                                        _ = v42
                                                                    default:
                                                                        switch "hello" {
                                                                        case "":
                                                                            if len(v6) > 0 {
                                                                                v6[0] = v38
                                                                            }
                                                                        default:
                                                                            switch "" {
                                                                            case "a":
                                                                                v43 := GenS6{F0: -242, F1: v33}
                                                                                _ = v43
                                                                            }
                                                                        }
                                                                        var v44 uint = uint(v4)
                                                                        v2 -= (v2 >> ((v44) & 63))
                                                                        _ = v44
                                                                    }
                                                                case "foo bar":
                                                                    if v33 {
                                                                        v6 = append(v6, (v38 | (v38 - (v38 & (v38 & (v38 >> ((65535) & 63)))))))
                                                                        if (42 == 1) {
                                                                            var v45 float64 = (v28 - (-2.25 - (((-2.25 + (((v28 - 0.0) * 1.5) + (v28 + v28))) + v28) + (v28 - 1.5))))
                                                                            _ = v45
                                                                        }
                                                                        if false {
                                                                            var v46 int64 = (v5 / ((-1) | 1))
                                                                            fmt.Println(v13, v28, v1)
                                                                            if (true || (!false)) {
                                                                                v33 = ((v4 % ((v4) | 1)) > v4)
                                                                                if len(v6) > 0 {
                                                                                    v6[0] = 65535
                                                                                }
                                                                                if (v13 == v13) {
                                                                                    var v47 uint32 = (v38 << ((200) & 63))
                                                                                    if len(v6) > 0 {
                                                                                        v6[0] = 4294967295
                                                                                    }
                                                                                    v48 := v3[""]
                                                                                    v49 := GenS7{F0: 605, F1: (v13 | 920), F2: 2}
                                                                                    if (v4 >= (v4 | v4)) {
                                                                                        var v50 float64 = 1.5
                                                                                        switch (v40 ^ 539) {
                                                                                        case 4294967295:
                                                                                            v51 := GenS8{F0: 795, F1: v2, F2: 1}
                                                                                            _ = v51
                                                                                        case 128:
                                                                                            if ((42 > 255) && v33) {
                                                                                                fmt.Println(v13, v46)
                                                                                            } else {
                                                                                                if len(v6) >= 2 {
                                                                                                    v6 = v6[1:len(v6)]
                                                                                                }
                                                                                                if true {
                                                                                                    var v52 int16 = v32
                                                                                                    fmt.Println(v13, v28)
                                                                                                    _ = v52
                                                                                                } else {
                                                                                                    var v53 float32 = 3.14159
                                                                                                    if (false && (((((-2.25 * (v53 / -2.25)) == (v1 / v53)) || (!((v50 + (((float64(v1) * v50) * (((v28 + v50) + float64(v53)) * v28)) + v50)) <= v28))) || (768 < v2)) && (!(!v33)))) {
                                                                                                        v28 += 0.1
                                                                                                        _ = len(v3)
                                                                                                        v54 := GenS1{F0: (v2 >> ((768) & 63)), F1: 0}
                                                                                                        switch v48 {
                                                                                                        case -1:
                                                                                                            fmt.Println(v33, v13, v47)
                                                                                                        case 2147483647:
                                                                                                            if ("a" != "hello") {
                                                                                                                var v55 uint = 1
                                                                                                                switch "foo bar" {
                                                                                                                case "foo bar":
                                                                                                                    var v56 int64 = v5
                                                                                                                    _ = v56
                                                                                                                case "a":
                                                                                                                    if (42 == -240) {
                                                                                                                        var v57 string = ""
                                                                                                                        v6 = append(v6, uint32(v4))
                                                                                                                        v55 = v55
                                                                                                                        v58 := GenS8{F0: (^(v55)), F1: 65535, F2: 42}
                                                                                                                        _ = v57
                                                                                                                        _ = v58
                                                                                                                    } else {
                                                                                                                        v40 = v40
                                                                                                                        fmt.Println(v32, v4, v13)
                                                                                                                    }
                                                                                                                    v55 = v55
                                                                                                                case "":
                                                                                                                    var v59 int32 = 212
                                                                                                                    _ = v59
                                                                                                                }
                                                                                                                switch (v40 << ((0) & 63)) {
                                                                                                                case 32767:
                                                                                                                    v6 = append(v6, 2147483647)
                                                                                                                case 32768:
                                                                                                                    var v60 string = "a"
                                                                                                                    _ = v60
                                                                                                                default:
                                                                                                                    v40 = (fi29 / ((int(v55)) | 1))
                                                                                                                }
                                                                                                                _ = v55
                                                                                                            }
                                                                                                        }
                                                                                                        v61 := v3["a"]
                                                                                                        _ = v54
                                                                                                        _ = v61
                                                                                                    } else {
                                                                                                        _ = len(v3)
                                                                                                        switch ("hello" + "") {
                                                                                                        case "a":
                                                                                                            if ((v13 & v13) != (v13 % ((511) | 1))) {
                                                                                                                if len(v6) > 0 {
                                                                                                                    v6[0] = (^(v38))
                                                                                                                }
                                                                                                                v32 += v32
                                                                                                            }
                                                                                                        case "foo bar":
                                                                                                            v33 = ((-9223372036854775808 < (v5 + v46)) && (v4 < v4))
                                                                                                            if ((v4 * v4) > (v4 >> ((1) & 63))) {
                                                                                                                var v62 uint64 = v13
                                                                                                                v32 -= v32
                                                                                                                var v63 int64 = (v46 & -557)
                                                                                                                _ = v62
                                                                                                                _ = v63
                                                                                                            } else {
                                                                                                                switch v48 {
                                                                                                                case -1:
                                                                                                                    if len(v6) >= 2 {
                                                                                                                        v6 = v6[1:len(v6)]
                                                                                                                    }
                                                                                                                case 28:
                                                                                                                    fmt.Println(v38, v50)
                                                                                                                }
                                                                                                                switch (fi29 - -9223372036854775808) {
                                                                                                                case 32768:
                                                                                                                    if (((v13 - v13) <= 4294967295) && (2 <= (v4 >> ((42) & 63)))) {
                                                                                                                        var v64 int = v48
                                                                                                                        if true {
                                                                                                                            switch 32767 {
                                                                                                                            case 32767:
                                                                                                                                var v65 uint16 = 255
                                                                                                                                _ = v65
                                                                                                                            case 6:
                                                                                                                                v66 := GenS9{F0: v4, F1: (v13 * (v13 / ((v13) | 1))), F2: ((true || (54 >= -128)) && (v13 > (^(v13))))}
                                                                                                                                _ = v66
                                                                                                                            default:
                                                                                                                                v53 = (v53 - v53)
                                                                                                                            }
                                                                                                                            v67 := GenS10{F0: v13}
                                                                                                                            if v33 {
                                                                                                                                var v68 uint16 = 373
                                                                                                                                switch 42 {
                                                                                                                                case 32767:
                                                                                                                                    fmt.Println(fi29)
                                                                                                                                case -1:
                                                                                                                                    if (!(v33 && v33)) {
                                                                                                                                        var v69 uint16 = (v68 % ((v2) | 1))
                                                                                                                                        v70 := GenS11{F0: 101}
                                                                                                                                        _ = v69
                                                                                                                                        _ = v70
                                                                                                                                    } else {
                                                                                                                                        v40 -= (fi29 << ((18446744073709551615) & 63))
                                                                                                                                    }
                                                                                                                                case 0:
                                                                                                                                    switch "hello" {
                                                                                                                                    case "":
                                                                                                                                        switch 758 {
                                                                                                                                        case 32768:
                                                                                                                                            var v71 uint16 = (v2 - 65535)
                                                                                                                                            _ = v71
                                                                                                                                        case -713:
                                                                                                                                            fmt.Println(v48)
                                                                                                                                        case -9223372036854775808:
                                                                                                                                            var v72 bool = v33
                                                                                                                                            _ = v72
                                                                                                                                        default:
                                                                                                                                            var v73 uint32 = v47
                                                                                                                                            _ = v73
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                }
                                                                                                                                _ = v68
                                                                                                                            } else {
                                                                                                                                v74 := GenS6{F0: (-(v64)), F1: v33}
                                                                                                                                v64 = (^(fi30))
                                                                                                                                v75 := GenS5{F0: (^(v47)), F1: v38}
                                                                                                                                v2 -= uint16(v38)
                                                                                                                                _ = v74
                                                                                                                                _ = v75
                                                                                                                            }
                                                                                                                            var v76 float32 = (v53 * (v53 / v53))
                                                                                                                            _ = v67
                                                                                                                            _ = v76
                                                                                                                        }
                                                                                                                        if len(v6) > 0 {
                                                                                                                            v6[0] = 65535
                                                                                                                        }
                                                                                                                        _ = v64
                                                                                                                    } else {
                                                                                                                        _ = len(v3)
                                                                                                                        v4 = v4
                                                                                                                        v5 = 1
                                                                                                                    }
                                                                                                                case -1:
                                                                                                                    v40 = (v40 + 32767)
                                                                                                                }
                                                                                                                v77 := GenS12{F0: ((v47 ^ v47) < (v38 << ((uint(v2)) & 63))), F1: 38, F2: int8(v32)}
                                                                                                                if (!v33) {
                                                                                                                    if false {
                                                                                                                        v1 = v53
                                                                                                                    }
                                                                                                                    var v78 int16 = (v32 << ((197) & 63))
                                                                                                                    var v79 int8 = 0
                                                                                                                    _ = v78
                                                                                                                    _ = v79
                                                                                                                }
                                                                                                                _ = v77
                                                                                                            }
                                                                                                        case "hello":
                                                                                                            var v80 float64 = (v28 * (0.1 / (float64(v53) + (0.1 * ((v28 - 1e10) + 1.5)))))
                                                                                                            _ = v80
                                                                                                        }
                                                                                                    }
                                                                                                    v81 := GenS13{F0: v13, F1: 0, F2: (v5 / ((v46) | 1))}
                                                                                                    _ = v53
                                                                                                    _ = v81
                                                                                                }
                                                                                            }
                                                                                        case 2147483647:
                                                                                            v28 += v50
                                                                                        default:
                                                                                            if (v33 || v33) {
                                                                                                var v82 string = "hello"
                                                                                                v83 := GenS14{F0: (v2 ^ 42)}
                                                                                                var v84 int = (v48 * fi29)
                                                                                                v85 := GenS14{F0: (v2 + 42)}
                                                                                                var v86 float32 = -2.25
                                                                                                v3[("" + v82)] = (-(v40))
                                                                                                var v87 uint16 = (v2 & (v2 / (((v2 + v2)) | 1)))
                                                                                                _ = v82
                                                                                                _ = v83
                                                                                                _ = v84
                                                                                                _ = v85
                                                                                                _ = v86
                                                                                                _ = v87
                                                                                            }
                                                                                        }
                                                                                        var v88 int64 = -949
                                                                                        _ = v50
                                                                                        _ = v88
                                                                                    } else {
                                                                                        fmt.Println(v32, v46, v46)
                                                                                        var v89 uint = 0
                                                                                        _ = v89
                                                                                    }
                                                                                    _ = v47
                                                                                    _ = v48
                                                                                    _ = v49
                                                                                } else {
                                                                                    var v90 uint8 = (^(v4))
                                                                                    if ((1 <= 16) && v33) {
                                                                                        var v91 int64 = (v46 % ((v46) | 1))
                                                                                        v46 = v91
                                                                                        _ = v91
                                                                                    }
                                                                                    v1 = ((((0.0 / 1.0) - (0.0 - v1)) + v1) - v1)
                                                                                    v92 := GenS9{F0: v4, F1: v13, F2: (227 < (fi29 | -9223372036854775808))}
                                                                                    _ = len(v6)
                                                                                    switch (fi30 >> ((74) & 63)) {
                                                                                    case -2147483648:
                                                                                        v46 = v46
                                                                                    }
                                                                                    _ = v90
                                                                                    _ = v92
                                                                                }
                                                                            }
                                                                            fmt.Println(v33)
                                                                            v93 := GenS1{F0: (v2 - v2), F1: 127}
                                                                            _ = v46
                                                                            _ = v93
                                                                        } else {
                                                                            switch -235 {
                                                                            case 32767:
                                                                                v94 := GenS11{F0: 127}
                                                                                _ = v94
                                                                            case 127:
                                                                                var v95 uint64 = (v13 ^ (^(v13)))
                                                                                _ = v95
                                                                            default:
                                                                                v40 = (fi29 * fi29)
                                                                            }
                                                                        }
                                                                        if true {
                                                                            fmt.Println(v40, v28, fi30)
                                                                            var v96 float32 = (v1 / v1)
                                                                            if ((v28 - (v28 * ((v28 + v28) / (float64(1.5) / (v28 * v28))))) > (float64(float32(((3.14159 - (float64(-2.25) + (1e10 + 1.5))) + float64(-2.25)))) + v28)) {
                                                                                var v97 int = (fi30 ^ fi30)
                                                                                v28 -= (0.1 + ((((v28 * 0.1) * v28) + (v28 + (float64(0.0) - float64(((v96 - v96) - 1.5))))) * float64((v96 * (v96 / (v1 + v96))))))
                                                                                v5 = (v5 % ((42) | 1))
                                                                                _ = v97
                                                                            } else {
                                                                                var v98 uint64 = uint64(v38)
                                                                                _ = v98
                                                                            }
                                                                            var v99 string = "a"
                                                                            if ((v33 || (!((int8(v2) <= 0) || v33))) && v33) {
                                                                                if len(v6) >= 2 {
                                                                                    v6 = v6[1:len(v6)]
                                                                                }
                                                                                v100 := GenS1{F0: 0, F1: 1}
                                                                                v101 := GenS11{F0: 42}
                                                                                var v102 int64 = v5
                                                                                _ = v100
                                                                                _ = v101
                                                                                _ = v102
                                                                            }
                                                                            v96 += ((v1 - -2.25) / (v1 + (float32((((v28 / v28) - 3.14159) * 0.1)) / float32(v28))))
                                                                            v103 := GenS10{F0: v13}
                                                                            v33 = (v5 < (-(v5)))
                                                                            _ = v103
                                                                            _ = v96
                                                                            _ = v99
                                                                        } else {
                                                                            var v104 float64 = v28
                                                                            _ = v104
                                                                        }
                                                                        v38 = (v38 * 4294967295)
                                                                    } else {
                                                                        var v105 uint32 = v38
                                                                        var v106 int64 = (v5 >> ((65535) & 63))
                                                                        _ = v105
                                                                        _ = v106
                                                                    }
                                                                }
                                                                if (v33 || v33) {
                                                                    if ((v38 << ((156) & 63)) != 0) {
                                                                        fmt.Println(v40, v1)
                                                                    } else {
                                                                        var v107 uint16 = v2
                                                                        v4 = 128
                                                                        switch "hello" {
                                                                        case "\n\t":
                                                                            var v108 int8 = int8(v13)
                                                                            _ = v108
                                                                        }
                                                                        _ = v107
                                                                    }
                                                                } else {
                                                                    v33 = (0 >= 0)
                                                                    var v109 int64 = v5
                                                                    _ = v109
                                                                }
                                                                _ = v40
                                                            } else {
                                                                if false {
                                                                    var v110 float32 = ((v1 + 1.5) / ((float32(float64(0.1)) + 1.5) / (0.0 + float32(3.14159))))
                                                                    if true {
                                                                        var v111 uint8 = v4
                                                                        fmt.Println(v5, v2)
                                                                        fmt.Println(v13, v111, v2)
                                                                        _ = v111
                                                                    }
                                                                    if true {
                                                                        v112 := GenS7{F0: (v2 ^ (v2 / ((uint16(v4)) | 1))), F1: 36, F2: (^(v13))}
                                                                        var v113 uint8 = 119
                                                                        v1 = v110
                                                                        _ = v112
                                                                        _ = v113
                                                                    } else {
                                                                        var v114 int16 = (v32 / ((v32) | 1))
                                                                        _ = v114
                                                                    }
                                                                    _ = v110
                                                                } else {
                                                                    v115 := GenS1{F0: (v2 & (v2 << ((42) & 63))), F1: 1}
                                                                    if v33 {
                                                                        if len(v6) >= 2 {
                                                                            v6 = v6[1:len(v6)]
                                                                        }
                                                                    }
                                                                    v116 := v3[((("a" + "") + "foo bar") + "\n\t")]
                                                                    _ = v115
                                                                    _ = v116
                                                                }
                                                            }
                                                            _ = len(v3)
                                                        }
                                                        v32 = -539
                                                        v3["foo bar"] = 9223372036854775807
                                                        v33 = (-572 > -697)
                                                        if ((^(v4)) > (v4 >> ((456) & 63))) {
                                                            v33 = ((v5 % ((int64(v4)) | 1)) == v5)
                                                            switch (fi29 | -9223372036854775808) {
                                                            case 32767:
                                                                var v117 int64 = v5
                                                                _ = v117
                                                            case -2147483648:
                                                                v13 += (v13 >> ((0) & 63))
                                                            case -1:
                                                                v2 = v2
                                                                _ = len(v6)
                                                            default:
                                                                var v118 int = 968
                                                                _ = v118
                                                            }
                                                        }
                                                        switch (^(fi30)) {
                                                        case 0:
                                                            v28 = v28
                                                        }
                                                        _ = v38
                                                    }
                                                }
                                                v119 := GenS8{F0: 18446744073709551615, F1: v2, F2: -1}
                                                _ = v119
                                            }
                                            v120 := GenS15{F0: fi30, F1: 921, F2: (-(fi29))}
                                            _ = v120
                                        } else {
                                            var v121 uint32 = 2147483647
                                            var v122 uint = 255
                                            var v123 int = -9223372036854775808
                                            _ = v121
                                            _ = v122
                                            _ = v123
                                        }
                                    case 0:
                                        var v124 int32 = -1
                                        _ = v124
                                    case 2147483647:
                                        if v33 {
                                            switch 127 {
                                            case 32768:
                                                if v33 {
                                                    delete(v3, "a")
                                                    v2 += (^(v2))
                                                    fmt.Println(fi30)
                                                }
                                            case 0:
                                                var v125 int8 = 55
                                                _ = v125
                                            default:
                                                v1 = (v1 * float32(((v28 * 0.0) * 1.5)))
                                            }
                                        }
                                    default:
                                        switch (fi30 ^ (fi29 % (((^(fi29))) | 1))) {
                                        case 0:
                                            fmt.Println(v2, fi30, v5)
                                        case 4294967295:
                                            fmt.Println(v13, v13)
                                        default:
                                            switch (fi30 % ((fi30) | 1)) {
                                            case 92:
                                                v13 -= 1
                                                var v126 bool = v33
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                                _ = v126
                                            case 176:
                                                v1 -= float32(1.5)
                                            }
                                        }
                                    }
                                default:
                                    _ = len(v3)
                                }
                                v127 := GenS16{F0: fi30, F1: 2147483647, F2: (fi30 | -1)}
                                if false {
                                    var v128 uint64 = 0
                                    fmt.Println(v1)
                                    v129 := GenS11{F0: -39}
                                    _ = v128
                                    _ = v129
                                }
                                _ = v127
                                _ = v32
                                _ = v33
                            } else {
                                var v130 float64 = (((v28 / ((-2.25 * (float64(float32((float64(float32(v28)) + 0.1))) - v28)) / float64(v1))) * v28) - (v28 * ((v28 + v28) - v28)))
                                v2 -= v2
                                v131 := GenS17{F0: 1, F1: false}
                                v130 += 0.1
                                switch -968 {
                                case 9223372036854775807:
                                    if false {
                                        v132 := GenS18{F0: (v2 % ((255) | 1)), F1: (v5 % ((-9223372036854775808) | 1))}
                                        var v133 float32 = v1
                                        _ = v132
                                        _ = v133
                                    }
                                case 4294967295:
                                    fmt.Println(fi30, fi30, v130)
                                case 959:
                                    switch (fi30 * fi30) {
                                    case -9223372036854775808:
                                        v13 += (v13 % (((v13 ^ v13)) | 1))
                                    case 128:
                                        v134 := GenS19{F0: 189, F1: fi30}
                                        _ = v134
                                    case -839:
                                        v135 := GenS2{F0: 9223372036854775807}
                                        _ = v135
                                    default:
                                        if false {
                                            v136 := GenS20{F0: "a"}
                                            v2 -= v2
                                            _ = v136
                                        } else {
                                            var v137 int16 = 1000
                                            switch "foo bar" {
                                            case "a":
                                                if ((v2 % (((v2 * v2)) | 1)) < v2) {
                                                    fmt.Println(v2, v130, v5)
                                                    var v138 string = "hello"
                                                    v1 += v1
                                                    _ = v138
                                                } else {
                                                    fmt.Println(v28, v13)
                                                }
                                            case "foo bar":
                                                if (v1 < ((v1 * v1) + float32(3.14159))) {
                                                    v1 += v1
                                                    var v139 int16 = (v137 << ((1) & 63))
                                                    if (fi29 == (fi29 >> ((18446744073709551615) & 63))) {
                                                        var v140 float64 = v28
                                                        v141 := v3["hello"]
                                                        var v142 int = (fi29 * 1)
                                                        _ = v140
                                                        _ = v141
                                                        _ = v142
                                                    } else {
                                                        v4 += (^(v4))
                                                        if false {
                                                            v28 -= v130
                                                            v4 = (v4 / (((v4 & v4)) | 1))
                                                            var v143 int32 = 227
                                                            _ = v143
                                                        } else {
                                                            if (v1 <= v1) {
                                                                if true {
                                                                    switch "hello" {
                                                                    case "hello":
                                                                        var v144 float32 = 3.14159
                                                                        _ = v144
                                                                    case "\n\t":
                                                                        var v145 uint16 = uint16(v4)
                                                                        _ = v145
                                                                    }
                                                                } else {
                                                                    var v146 int16 = (v139 & (-(v137)))
                                                                    var v147 float32 = 3.14159
                                                                    _ = v146
                                                                    _ = v147
                                                                }
                                                                switch 98 {
                                                                case 2147483647:
                                                                    v148 := v3["\n\t"]
                                                                    _ = v148
                                                                case -353:
                                                                    switch "" {
                                                                    case "a":
                                                                        switch "a" {
                                                                        case "foo bar":
                                                                            delete(v3, ("\n\t" + "\n\t"))
                                                                        case "hello":
                                                                            fmt.Println(v137, v139)
                                                                        case "a":
                                                                            if (false || true) {
                                                                                var v149 string = (("\n\t" + "\n\t") + "a")
                                                                                _ = v149
                                                                            } else {
                                                                                switch 4294967295 {
                                                                                case 9223372036854775807:
                                                                                    v150 := v3["a"]
                                                                                    _ = len(v6)
                                                                                    var v151 int16 = (v139 & (v139 << ((18446744073709551615) & 63)))
                                                                                    _ = v150
                                                                                    _ = v151
                                                                                }
                                                                                v6 = append(v6, 698)
                                                                                v6 = append(v6, 2147483648)
                                                                            }
                                                                        default:
                                                                            v152 := v3[(((((("\n\t" + "foo bar") + "\n\t") + "hello") + "") + "foo bar") + "foo bar")]
                                                                            v137 += (v137 - (-(v139)))
                                                                            _ = v152
                                                                        }
                                                                    default:
                                                                        if (!false) {
                                                                            switch 2147483647 {
                                                                            case -9223372036854775808:
                                                                                v153 := GenS21{F0: v5}
                                                                                _ = v153
                                                                            case 4294967295:
                                                                                v2 = v2
                                                                            case 55:
                                                                                var v154 uint = 42
                                                                                _ = v154
                                                                            }
                                                                            if (v130 >= (v130 / v28)) {
                                                                                var v155 int64 = (-(v5))
                                                                                _ = v155
                                                                            }
                                                                        }
                                                                    }
                                                                case -2147483648:
                                                                    v3["\n\t"] = (^(fi29))
                                                                }
                                                                v137 = v139
                                                            } else {
                                                                switch "\n\t" {
                                                                case "hello":
                                                                    var v156 string = ("\n\t" + "\n\t")
                                                                    _ = v156
                                                                }
                                                                v157 := GenS22{F0: 42, F1: 42, F2: v139}
                                                                if true {
                                                                    fmt.Println(v13)
                                                                } else {
                                                                    v1 -= 0.1
                                                                    v139 -= (v139 - (-(v137)))
                                                                }
                                                                _ = v157
                                                            }
                                                            if len(v6) > 0 {
                                                                v6[0] = uint32(v2)
                                                            }
                                                        }
                                                        if (!(!(v5 != v5))) {
                                                            if false {
                                                                fmt.Println(v2)
                                                                var v158 int64 = v5
                                                                v130 += float64(v1)
                                                                _ = v158
                                                            }
                                                        }
                                                        if (true && false) {
                                                            if len(v6) > 0 {
                                                                v6[0] = 2147483648
                                                            }
                                                            v13 = 0
                                                        } else {
                                                            var v159 int16 = (v139 ^ 1)
                                                            _ = v159
                                                        }
                                                    }
                                                    var v160 int64 = (-(v5))
                                                    _ = v139
                                                    _ = v160
                                                } else {
                                                    v161 := GenS23{F0: 2}
                                                    switch "a" {
                                                    case "foo bar":
                                                        if len(v6) >= 2 {
                                                            v6 = v6[1:len(v6)]
                                                        }
                                                    case "":
                                                        if true {
                                                            var v162 int8 = -1
                                                            v163 := GenS24{F0: fi29, F1: 4294967295}
                                                            v28 = -2.25
                                                            _ = v162
                                                            _ = v163
                                                        }
                                                    case "hello":
                                                        switch 9223372036854775807 {
                                                        case -144:
                                                            fmt.Println(fi29, v5, v1)
                                                        case 871:
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                        }
                                                    }
                                                    var v164 uint64 = v13
                                                    _ = v161
                                                    _ = v164
                                                }
                                            default:
                                                var v165 int32 = 1
                                                _ = v165
                                            }
                                            v166 := GenS25{F0: int8(v4), F1: 1}
                                            _ = v137
                                            _ = v166
                                        }
                                    }
                                default:
                                    if true {
                                        fmt.Println(v130, v4, v1)
                                    } else {
                                        switch ("\n\t" + "foo bar") {
                                        case "\n\t":
                                            var v167 uint32 = 2147483648
                                            _ = v167
                                        }
                                    }
                                }
                                _ = v130
                                _ = v131
                            }
                        case -762:
                            v168 := GenS26{F0: 1000, F1: (v4 % (((^(v4))) | 1)), F2: v2}
                            _ = v168
                        default:
                            v169 := GenS1{F0: v2, F1: -1}
                            _ = v169
                        }
                    default:
                        v170 := GenS4{F0: 42}
                        v171 := GenS19{F0: 138, F1: (fi29 | fi29)}
                        _ = v170
                        _ = v171
                    }
                    var v172 uint32 = 65535
                    _ = v172
                default:
                    switch fi30 {
                    case 828:
                        if true {
                            v173 := v3[""]
                            _ = v173
                        } else {
                            var v174 bool = (v28 <= v28)
                            switch "foo bar" {
                            case "":
                                var v175 int16 = -245
                                _ = v175
                            default:
                                v5 = -1
                                switch "\n\t" {
                                case "hello":
                                    v174 = (v1 == (v1 / 0.1))
                                default:
                                    var v176 int8 = 127
                                    _ = v176
                                }
                                if v174 {
                                    v3[("a" + "\n\t")] = int(v13)
                                } else {
                                    v177 := GenS6{F0: 128, F1: (v4 != 1)}
                                    var v178 string = "hello"
                                    var v179 int8 = 127
                                    switch (-(fi30)) {
                                    case -2147483648:
                                        var v180 uint = 572
                                        _ = v180
                                    case 9223372036854775807:
                                        v179 += (v179 + (v179 / ((v179) | 1)))
                                    default:
                                        switch 32767 {
                                        case -9223372036854775808:
                                            v181 := GenS27{F0: (v4 >> ((2) & 63)), F1: 370, F2: (42 <= 223)}
                                            _ = v181
                                        case -337:
                                            if v174 {
                                                var v182 bool = (4294967295 > 2)
                                                if v174 {
                                                    fmt.Println(v2, v1)
                                                }
                                                _ = v182
                                            }
                                        default:
                                            v183 := GenS28{F0: "hello", F1: fi29}
                                            _ = v183
                                        }
                                    }
                                    switch v178 {
                                    case "\n\t":
                                        var v184 int8 = -66
                                        _ = v184
                                    }
                                    _ = v177
                                    _ = v178
                                    _ = v179
                                }
                            }
                            switch fi30 {
                            case 104:
                                var v185 float32 = (v1 + (v1 - v1))
                                _ = v185
                            }
                            v2 += (^(v2))
                            _ = v174
                        }
                    case -166:
                        v186 := GenS29{F0: (v13 + v13), F1: (^(fi29)), F2: 0}
                        _ = v186
                    }
                }
                _ = fi30
                _ = v31
            }
            _ = len(v6)
            var v187 uint8 = (v4 | 128)
            _ = fi29
            _ = v187
        }
        _ = v28
    } else {
        for fi188 := 0; fi188 < 8; fi188++ {
            var v189 float32 = ((v1 + v1) - v1)
            var v190 uint = 18446744073709551615
            _ = fi188
            _ = v189
            _ = v190
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        var v191 int8 = 100
        _ = v191
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
