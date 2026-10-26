package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 uint32 }
type GenS2 struct { F0 string; F1 string; F2 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint32; F1 int16; F2 uint }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int8; F1 uint32 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint8; F1 int; F2 uint8 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 uint; F1 uint16; F2 uint8 }
type GenS14 struct { F0 string; F1 int }
type GenS15 struct { F0 string }
type GenS16 struct { F0 bool; F1 bool; F2 int32 }
type GenS17 struct { F0 int8; F1 uint64; F2 int8 }
type GenS18 struct { F0 uint32; F1 uint64 }
type GenS19 struct { F0 int64; F1 uint64; F2 int }
type GenS20 struct { F0 int32; F1 int; F2 uint64 }

func main() {
    var v1 float64 = ((1e10 * float64(1.5)) + float64(0.1))
    var v2 string = "\n\t"
    v3 := make(map[string]string)
    v4 := []uint32{42, 388, 132}
    fmt.Println(v1)
    _ = len(v3)
    v1 = float64(0.1)
    switch 32767 {
    case 127:
        if false {
            var v5 int16 = 1000
            _ = len(v3)
            _ = v5
        }
    case 32768:
        switch (v2 + v2) {
        case "":
            delete(v3, v2)
        case "\n\t":
            v1 = (0.0 - ((v1 - float64(1.5)) * v1))
        }
    default:
        v6 := GenS1{F0: 1, F1: ((v2 + v2) + ("a" + "\n\t")), F2: 292}
        _ = v6
    }
    if false {
        v1 = 0.1
    }
    v2 = v2
    var v7 int64 = 143
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 float32 = (((3.14159 * ((0.1 * (-2.25 / (float32(float64(-2.25)) / 1.0))) + -2.25)) * -2.25) - 1e10)
        _ = fi8
        _ = v9
    }
    if (!(4294967295 > 851)) {
        v7 = v7
        v1 = 3.14159
    } else {
        if false {
            v7 += (v7 & (v7 % ((v7) | 1)))
            var v10 uint32 = 351
            var v11 uint64 = 751
            _ = v10
            _ = v11
        } else {
            v7 -= v7
            v1 = (float64(float32(-2.25)) - (v1 - (0.1 - (v1 - v1))))
        }
    }
    if true {
        v2 = ""
        v12 := GenS2{F0: "foo bar", F1: v2, F2: true}
        _ = v12
    } else {
        switch 127 {
        case -1:
            var v13 float32 = ((((float32(v1) + 0.1) + -2.25) / 1.0) / 0.1)
            _ = v13
        case 127:
            switch v2 {
            case "foo bar":
                v14 := GenS3{F0: v7}
                _ = v14
            default:
                var v15 int64 = 892
                _ = v15
            }
        default:
            if ((v2 + v2) <= "a") {
                fmt.Println(v7, v2, v2)
            } else {
                if false {
                    switch v2 {
                    case "\n\t":
                        if true {
                            if true {
                                if true {
                                    var v16 bool = true
                                    if (98 > 2) {
                                        var v17 uint32 = 4294967295
                                        _ = v17
                                    } else {
                                        v18 := v3["hello"]
                                        if false {
                                            fmt.Println(v16)
                                            v19 := GenS3{F0: v7}
                                            fmt.Println(v18, v7)
                                            _ = v19
                                        } else {
                                            for fi20 := 0; fi20 < 1; fi20++ {
                                                v7 -= (v7 << ((0) & 63))
                                                switch v18 {
                                                case "a":
                                                    var v21 uint32 = 2147483647
                                                    _ = v21
                                                case "":
                                                    v22 := GenS4{F0: v7}
                                                    _ = v22
                                                }
                                                _ = fi20
                                            }
                                            v23 := GenS5{F0: 928, F1: 186, F2: 2}
                                            if false {
                                                if false {
                                                    fmt.Println(v7)
                                                    v7 -= v7
                                                    switch ("" + "\n\t") {
                                                    case "\n\t":
                                                        if ((v16 || v16) && v16) {
                                                            v24 := GenS6{F0: 34}
                                                            switch -812 {
                                                            case -350:
                                                                if v16 {
                                                                    switch 128 {
                                                                    case 2147483647:
                                                                        switch ((v2 + (("\n\t" + "\n\t") + (v18 + ("foo bar" + ("foo bar" + v2))))) + v2) {
                                                                        case "hello":
                                                                            var v25 float32 = (((((3.14159 * (((0.1 + float32((v1 / 3.14159))) * 1.5) * 1.5)) * ((float32(0.0) / 1.0) * 0.0)) * 1.5) * -2.25) / ((0.0 + -2.25) / ((-2.25 - (-2.25 / 0.1)) / (1.5 + 1.5))))
                                                                            _ = v25
                                                                        case "":
                                                                            v26 := GenS7{F0: 42, F1: 347}
                                                                            _ = v26
                                                                        case "a":
                                                                            v4 = append(v4, 84)
                                                                        default:
                                                                            v27 := GenS8{F0: 2147483647}
                                                                            _ = v27
                                                                        }
                                                                    case 32767:
                                                                        if v16 {
                                                                            switch v18 {
                                                                            case "foo bar":
                                                                                _ = len(v3)
                                                                            }
                                                                            for fi28 := 0; fi28 < 11; fi28++ {
                                                                                delete(v3, v18)
                                                                                _ = fi28
                                                                            }
                                                                            v29 := GenS8{F0: 1}
                                                                            _ = v29
                                                                        } else {
                                                                            for fi30 := 0; fi30 < 14; fi30++ {
                                                                                v18 = v18
                                                                                _ = fi30
                                                                            }
                                                                        }
                                                                    case 249:
                                                                        for fi31 := 0; fi31 < 14; fi31++ {
                                                                            v32 := v3[v2]
                                                                            var v33 int16 = -989
                                                                            _ = fi31
                                                                            _ = v32
                                                                            _ = v33
                                                                        }
                                                                    default:
                                                                        if v16 {
                                                                            v18 = (v18 + ("" + (v2 + v18)))
                                                                            var v34 int8 = -128
                                                                            v1 += (1.5 / v1)
                                                                            _ = v34
                                                                        } else {
                                                                            v18 = v18
                                                                            v35 := GenS2{F0: ("foo bar" + v18), F1: v18, F2: v16}
                                                                            var v36 int64 = (^(v7))
                                                                            _ = v35
                                                                            _ = v36
                                                                        }
                                                                    }
                                                                    switch 127 {
                                                                    case -366:
                                                                        if (42 < 1) {
                                                                            v4 = append(v4, 931)
                                                                            v4 = append(v4, 65535)
                                                                        } else {
                                                                            if false {
                                                                                if (65535 > 158) {
                                                                                    if false {
                                                                                        v1 = v1
                                                                                    }
                                                                                    if false {
                                                                                        delete(v3, "a")
                                                                                        var v37 int8 = -113
                                                                                        _ = v37
                                                                                    }
                                                                                } else {
                                                                                    v38 := GenS4{F0: (v7 + v7)}
                                                                                    var v39 uint64 = 0
                                                                                    _ = v38
                                                                                    _ = v39
                                                                                }
                                                                                var v40 uint8 = 127
                                                                                _ = v40
                                                                            }
                                                                        }
                                                                    default:
                                                                        var v41 int64 = 1
                                                                        _ = v41
                                                                    }
                                                                } else {
                                                                    var v42 int8 = -4
                                                                    _ = v42
                                                                }
                                                            case 32767:
                                                                var v43 float32 = -2.25
                                                                _ = v43
                                                            }
                                                            var v44 float32 = ((float32(0.0) * float32(v1)) + (0.0 / -2.25))
                                                            _ = v24
                                                            _ = v44
                                                        } else {
                                                            if v16 {
                                                                switch -1 {
                                                                case 517:
                                                                    for fi45 := 0; fi45 < 11; fi45++ {
                                                                        for fi46 := 0; fi46 < 7; fi46++ {
                                                                            v4 = append(v4, 2147483647)
                                                                            switch (v18 + v2) {
                                                                            case "hello":
                                                                                if v16 {
                                                                                    if len(v4) >= 2 {
                                                                                        v4 = v4[1:len(v4)]
                                                                                    }
                                                                                    var v47 int16 = 32767
                                                                                    _ = v47
                                                                                } else {
                                                                                    var v48 int8 = 51
                                                                                    v18 = v18
                                                                                    v49 := v3[""]
                                                                                    _ = v48
                                                                                    _ = v49
                                                                                }
                                                                            case "foo bar":
                                                                                var v50 float64 = (v1 - v1)
                                                                                _ = v50
                                                                            case "\n\t":
                                                                                fmt.Println(fi46)
                                                                            }
                                                                            _ = fi46
                                                                        }
                                                                        var v51 bool = (65535 == 4294967295)
                                                                        for fi52 := 0; fi52 < 1; fi52++ {
                                                                            if ((v7 & 945) >= 1) {
                                                                                fmt.Println(v7, fi52, v18)
                                                                                var v53 uint64 = 671
                                                                                var v54 int64 = v7
                                                                                _ = v53
                                                                                _ = v54
                                                                            } else {
                                                                                if v51 {
                                                                                    v18 = ("foo bar" + v2)
                                                                                    v55 := GenS8{F0: 42}
                                                                                    var v56 uint64 = 42
                                                                                    _ = v55
                                                                                    _ = v56
                                                                                }
                                                                                var v57 uint = 1
                                                                                if (141 > 32767) {
                                                                                    if ((false && v16) || ((1 > 0) || ((true && (1e10 <= 1.5)) || v51))) {
                                                                                        var v58 uint16 = 0
                                                                                        v1 -= 0.0
                                                                                        switch ("a" + v2) {
                                                                                        case "foo bar":
                                                                                            v59 := GenS3{F0: v7}
                                                                                            _ = v59
                                                                                        case "hello":
                                                                                            switch v2 {
                                                                                            case "":
                                                                                                v60 := GenS8{F0: 65535}
                                                                                                _ = v60
                                                                                            case "a":
                                                                                                var v61 uint8 = 0
                                                                                                _ = v61
                                                                                            case "foo bar":
                                                                                                switch v18 {
                                                                                                case "hello":
                                                                                                    switch (fi52 % ((9223372036854775807) | 1)) {
                                                                                                    case 4294967295:
                                                                                                        if len(v4) > 0 {
                                                                                                            v4[0] = 0
                                                                                                        }
                                                                                                    case 9223372036854775807:
                                                                                                        if v16 {
                                                                                                            var v62 int8 = int8(v7)
                                                                                                            _ = v62
                                                                                                        } else {
                                                                                                            switch (fi45 & -801) {
                                                                                                            case -160:
                                                                                                                v18 = ""
                                                                                                            default:
                                                                                                                if v16 {
                                                                                                                    if v16 {
                                                                                                                        fmt.Println(fi45)
                                                                                                                    }
                                                                                                                    v51 = false
                                                                                                                } else {
                                                                                                                    fmt.Println(v51, v57, v58)
                                                                                                                    v63 := GenS9{F0: 0, F1: -189, F2: 128}
                                                                                                                    fmt.Println(v57)
                                                                                                                    _ = v63
                                                                                                                }
                                                                                                                var v64 uint = 100
                                                                                                                v3[v2] = ("a" + (v2 + v18))
                                                                                                                _ = v64
                                                                                                            }
                                                                                                            v57 = (v57 / ((v57) | 1))
                                                                                                            v65 := GenS4{F0: 42}
                                                                                                            _ = v65
                                                                                                        }
                                                                                                    default:
                                                                                                        v2 = "hello"
                                                                                                    }
                                                                                                    v66 := GenS10{F0: 45}
                                                                                                    _ = len(v4)
                                                                                                    var v67 uint32 = 358
                                                                                                    _ = v66
                                                                                                    _ = v67
                                                                                                default:
                                                                                                    v57 += (v57 >> (((v57 - (v57 | 18446744073709551615))) & 63))
                                                                                                }
                                                                                                _ = len(v3)
                                                                                                v68 := GenS1{F0: -32768, F1: (v18 + (("" + v18) + v2)), F2: 386}
                                                                                                _ = v68
                                                                                            }
                                                                                        case "\n\t":
                                                                                            var v69 string = v2
                                                                                            _ = v69
                                                                                        default:
                                                                                            fmt.Println(v51, v51)
                                                                                        }
                                                                                        _ = v58
                                                                                    } else {
                                                                                        v18 = v18
                                                                                        if (42 != 55) {
                                                                                            v70 := GenS9{F0: 135, F1: fi45, F2: 0}
                                                                                            v71 := GenS1{F0: -1, F1: v2, F2: 2147483648}
                                                                                            if (!((v7 % ((-9223372036854775808) | 1)) <= 9223372036854775807)) {
                                                                                                var v72 uint32 = 921
                                                                                                _ = v72
                                                                                            }
                                                                                            if len(v4) > 0 {
                                                                                                v4[0] = 227
                                                                                            }
                                                                                            var v73 uint64 = 4294967295
                                                                                            _ = v70
                                                                                            _ = v71
                                                                                            _ = v73
                                                                                        } else {
                                                                                            var v74 bool = (2147483648 > 0)
                                                                                            _ = v74
                                                                                        }
                                                                                    }
                                                                                    if v16 {
                                                                                        if ((v1 / 0.1) != 1e10) {
                                                                                            v3[(("a" + "hello") + (v2 + "foo bar"))] = (v2 + (v18 + (v18 + v2)))
                                                                                            var v75 uint64 = 645
                                                                                            _ = v75
                                                                                        }
                                                                                        if v16 {
                                                                                            if (!v16) {
                                                                                                switch v2 {
                                                                                                case "foo bar":
                                                                                                    switch "" {
                                                                                                    case "foo bar":
                                                                                                        v18 = "foo bar"
                                                                                                    case "\n\t":
                                                                                                        _ = len(v4)
                                                                                                        if v51 {
                                                                                                            fmt.Println(v18)
                                                                                                        } else {
                                                                                                            _ = len(v3)
                                                                                                        }
                                                                                                        switch (fi45 % ((fi52) | 1)) {
                                                                                                        case 2147483647:
                                                                                                            v51 = (!(!(42 > 31)))
                                                                                                        case -160:
                                                                                                            _ = len(v4)
                                                                                                        case 128:
                                                                                                            delete(v3, "foo bar")
                                                                                                        default:
                                                                                                            v1 += v1
                                                                                                        }
                                                                                                        v76 := GenS11{F0: 42}
                                                                                                        v1 = v1
                                                                                                        _ = v76
                                                                                                    case "":
                                                                                                        var v77 float32 = float32(0.1)
                                                                                                        _ = v77
                                                                                                    }
                                                                                                }
                                                                                                var v78 int8 = 49
                                                                                                v57 = (v57 % (((v57 | v57)) | 1))
                                                                                                _ = v78
                                                                                            } else {
                                                                                                v79 := GenS6{F0: -48}
                                                                                                fmt.Println(v51)
                                                                                                _ = v79
                                                                                            }
                                                                                            var v80 int8 = 100
                                                                                            _ = v80
                                                                                        }
                                                                                        v81 := GenS12{F0: (v7 == (v7 << ((v57) & 63)))}
                                                                                        _ = v81
                                                                                    }
                                                                                } else {
                                                                                    v3[v18] = (v2 + v2)
                                                                                    v1 += (((v1 * (((-2.25 * float64((1.5 + (1e10 * 0.1)))) * (v1 * v1)) - float64((3.14159 / (1e10 / 1.0))))) + 0.0) * v1)
                                                                                }
                                                                                _ = v57
                                                                            }
                                                                            _ = fi52
                                                                        }
                                                                        _ = fi45
                                                                        _ = v51
                                                                    }
                                                                case 42:
                                                                    if (32767 > 1) {
                                                                        if (42 < 0) {
                                                                            _ = len(v4)
                                                                            for fi82 := 0; fi82 < 3; fi82++ {
                                                                                delete(v3, "hello")
                                                                                if (true || (32767 <= -1)) {
                                                                                    switch v18 {
                                                                                    case "\n\t":
                                                                                        v16 = v16
                                                                                    case "foo bar":
                                                                                        var v83 int = 1
                                                                                        _ = v83
                                                                                    default:
                                                                                        var v84 int16 = -1
                                                                                        _ = v84
                                                                                    }
                                                                                    var v85 string = ((((v2 + "hello") + "\n\t") + v18) + v2)
                                                                                    var v86 uint8 = 0
                                                                                    _ = v85
                                                                                    _ = v86
                                                                                }
                                                                                if v16 {
                                                                                    for fi87 := 0; fi87 < 6; fi87++ {
                                                                                        v88 := GenS13{F0: 18446744073709551615, F1: 1, F2: 128}
                                                                                        var v89 int = 0
                                                                                        var v90 bool = v16
                                                                                        _ = fi87
                                                                                        _ = v88
                                                                                        _ = v89
                                                                                        _ = v90
                                                                                    }
                                                                                    v18 = v18
                                                                                }
                                                                                _ = fi82
                                                                            }
                                                                            if v16 {
                                                                                for fi91 := 0; fi91 < 12; fi91++ {
                                                                                    switch v2 {
                                                                                    case "hello":
                                                                                        _ = len(v3)
                                                                                    case "\n\t":
                                                                                        for fi92 := 0; fi92 < 18; fi92++ {
                                                                                            fmt.Println(v1)
                                                                                            v2 = ("foo bar" + v18)
                                                                                            v16 = (508 <= 65535)
                                                                                            _ = fi92
                                                                                        }
                                                                                    case "foo bar":
                                                                                        for fi93 := 0; fi93 < 4; fi93++ {
                                                                                            switch fi91 {
                                                                                            case -73:
                                                                                                if (fi91 == 32767) {
                                                                                                    v4 = append(v4, 0)
                                                                                                }
                                                                                            case 128:
                                                                                                v1 = v1
                                                                                            case 986:
                                                                                                v18 = "foo bar"
                                                                                                v94 := GenS14{F0: "\n\t", F1: fi91}
                                                                                                _ = v94
                                                                                            default:
                                                                                                if len(v4) >= 2 {
                                                                                                    v4 = v4[1:len(v4)]
                                                                                                }
                                                                                            }
                                                                                            v7 = (v7 ^ v7)
                                                                                            var v95 int = -9223372036854775808
                                                                                            _ = fi93
                                                                                            _ = v95
                                                                                        }
                                                                                    }
                                                                                    if len(v4) > 0 {
                                                                                        v4[0] = 1
                                                                                    }
                                                                                    _ = fi91
                                                                                }
                                                                                v18 = v2
                                                                            } else {
                                                                                switch 0 {
                                                                                case -465:
                                                                                    for fi96 := 0; fi96 < 15; fi96++ {
                                                                                        v1 += (v1 - (v1 + float64((1.5 + (1e10 * (-2.25 - 1.5))))))
                                                                                        v1 = (((v1 * 1e10) * (((v1 + (v1 - 0.0)) / -2.25) - (v1 + (v1 / v1)))) - v1)
                                                                                        v2 = v2
                                                                                        _ = fi96
                                                                                    }
                                                                                case -9223372036854775808:
                                                                                    v97 := GenS13{F0: 2, F1: 0, F2: 0}
                                                                                    _ = v97
                                                                                default:
                                                                                    if ("foo bar" < ("" + ((v2 + v18) + v18))) {
                                                                                        v98 := GenS10{F0: 0}
                                                                                        _ = v98
                                                                                    }
                                                                                }
                                                                                if (42 <= 2) {
                                                                                    switch 128 {
                                                                                    case -980:
                                                                                        for fi99 := 0; fi99 < 13; fi99++ {
                                                                                            var v100 int16 = -1
                                                                                            _ = fi99
                                                                                            _ = v100
                                                                                        }
                                                                                    }
                                                                                    if v16 {
                                                                                        if (((1e10 - 3.14159) != ((3.14159 * 1.5) + float32(1e10))) && (!((true && true) || (-1 == 42)))) {
                                                                                            var v101 int32 = -895
                                                                                            _ = v101
                                                                                        } else {
                                                                                            switch -9223372036854775808 {
                                                                                            case 560:
                                                                                                _ = len(v3)
                                                                                            case 4294967295:
                                                                                                var v102 int = -1
                                                                                                _ = v102
                                                                                            default:
                                                                                                v1 = 0.1
                                                                                            }
                                                                                        }
                                                                                        v16 = ((0.0 * float32(v1)) < (1e10 + 0.0))
                                                                                    } else {
                                                                                        for fi103 := 0; fi103 < 17; fi103++ {
                                                                                            for fi104 := 0; fi104 < 20; fi104++ {
                                                                                                fmt.Println(fi103)
                                                                                                _ = fi104
                                                                                            }
                                                                                            v18 = v2
                                                                                            _ = fi103
                                                                                        }
                                                                                        v7 = v7
                                                                                    }
                                                                                }
                                                                            }
                                                                        } else {
                                                                            v4 = append(v4, 4294967295)
                                                                        }
                                                                    }
                                                                case 781:
                                                                    var v105 string = v2
                                                                    _ = v105
                                                                }
                                                                var v106 uint16 = 1
                                                                fmt.Println(v1, v18, v2)
                                                                _ = v106
                                                            }
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        }
                                                    case "hello":
                                                        if false {
                                                            for fi107 := 0; fi107 < 20; fi107++ {
                                                                switch v2 {
                                                                case "":
                                                                    if false {
                                                                        var v108 int32 = 1
                                                                        for fi109 := 0; fi109 < 18; fi109++ {
                                                                            if (!((3.14159 + ((v1 + (-2.25 / float64((-2.25 - (0.1 / 1.0))))) * (1.5 + 3.14159))) >= 1.5)) {
                                                                                switch v2 {
                                                                                case "\n\t":
                                                                                    if v16 {
                                                                                        if true {
                                                                                            switch (-(fi109)) {
                                                                                            case 127:
                                                                                                fmt.Println(v7)
                                                                                            case 32767:
                                                                                                var v110 uint8 = 0
                                                                                                _ = v110
                                                                                            case 0:
                                                                                                var v111 uint64 = 729
                                                                                                _ = v111
                                                                                            default:
                                                                                                var v112 float32 = 1.5
                                                                                                _ = v112
                                                                                            }
                                                                                            switch v2 {
                                                                                            case "\n\t":
                                                                                                delete(v3, "a")
                                                                                                v113 := v3["hello"]
                                                                                                v114 := GenS4{F0: (v7 - v7)}
                                                                                                _ = v113
                                                                                                _ = v114
                                                                                            }
                                                                                        }
                                                                                        var v115 uint = 42
                                                                                        if len(v4) >= 2 {
                                                                                            v4 = v4[1:len(v4)]
                                                                                        }
                                                                                        _ = v115
                                                                                    }
                                                                                    fmt.Println(v1, v16, v108)
                                                                                case "a":
                                                                                    v18 = v18
                                                                                }
                                                                            } else {
                                                                                var v116 string = v18
                                                                                _ = v116
                                                                            }
                                                                            _ = fi109
                                                                        }
                                                                        v1 -= v1
                                                                        _ = v108
                                                                    }
                                                                case "a":
                                                                    var v117 int32 = 2147483647
                                                                    _ = v117
                                                                case "\n\t":
                                                                    v118 := GenS14{F0: v18, F1: fi107}
                                                                    _ = v118
                                                                }
                                                                var v119 uint64 = 4294967295
                                                                _ = fi107
                                                                _ = v119
                                                            }
                                                            v18 = v18
                                                            v120 := GenS5{F0: 42, F1: 176, F2: 12}
                                                            _ = v120
                                                        }
                                                    default:
                                                        v18 = v18
                                                    }
                                                } else {
                                                    v7 += v7
                                                    v16 = true
                                                }
                                            }
                                            _ = v23
                                        }
                                        v7 += (v7 >> ((65535) & 63))
                                        _ = v18
                                    }
                                    v2 = "a"
                                    _ = v16
                                }
                            }
                        }
                    }
                    v7 -= (v7 / ((0) | 1))
                    if (2147483647 <= 100) {
                        var v121 int32 = -449
                        v2 = (v2 + (((v2 + v2) + v2) + v2))
                        var v122 int8 = 127
                        _ = v121
                        _ = v122
                    } else {
                        var v123 uint8 = 194
                        _ = v123
                    }
                } else {
                    fmt.Println(v1)
                }
            }
        }
        if false {
            var v124 int8 = 111
            fmt.Println(v1)
            fmt.Println(v124, v2, v1)
            _ = v124
        } else {
            _ = len(v4)
            var v125 float64 = 1.5
            switch 0 {
            case 371:
                switch 505 {
                case 0:
                    switch v2 {
                    case "\n\t":
                        switch v2 {
                        case "hello":
                            v126 := GenS15{F0: "hello"}
                            _ = v126
                        case "a":
                            if true {
                                var v127 int32 = 1
                                switch -9223372036854775808 {
                                case -2147483648:
                                    v128 := GenS6{F0: int8(v7)}
                                    _ = v128
                                default:
                                    var v129 uint8 = 0
                                    _ = v129
                                }
                                for fi130 := 0; fi130 < 5; fi130++ {
                                    if ((true || ((fi130 << ((4294967295) & 63)) == (fi130 | (fi130 & (fi130 >> ((4294967295) & 63)))))) && (fi130 >= fi130)) {
                                        for fi131 := 0; fi131 < 10; fi131++ {
                                            v7 += 939
                                            var v132 int16 = 983
                                            v133 := GenS4{F0: v7}
                                            _ = fi131
                                            _ = v132
                                            _ = v133
                                        }
                                        v134 := GenS16{F0: (!(1 <= v7)), F1: true, F2: 42}
                                        _ = v134
                                    } else {
                                        var v135 int32 = -847
                                        _ = v135
                                    }
                                    var v136 int16 = -291
                                    var v137 int = fi130
                                    _ = fi130
                                    _ = v136
                                    _ = v137
                                }
                                _ = v127
                            } else {
                                v138 := v3[v2]
                                var v139 int32 = -1
                                switch v2 {
                                case "hello":
                                    switch "foo bar" {
                                    case "foo bar":
                                        v4 = append(v4, 2147483648)
                                    case "a":
                                        v140 := v3[v2]
                                        _ = v140
                                    case "hello":
                                        for fi141 := 0; fi141 < 11; fi141++ {
                                            for fi142 := 0; fi142 < 9; fi142++ {
                                                if len(v4) > 0 {
                                                    v4[0] = 42
                                                }
                                                fmt.Println(v1, v125, v138)
                                                v7 += v7
                                                v139 -= (v139 + v139)
                                                var v143 float64 = float64(0.1)
                                                _ = fi142
                                                _ = v143
                                            }
                                            for fi144 := 0; fi144 < 11; fi144++ {
                                                if (!(false || false)) {
                                                    var v145 int = 128
                                                    _ = v145
                                                }
                                                v2 = "a"
                                                _ = fi144
                                            }
                                            v146 := GenS15{F0: ((v2 + v2) + v2)}
                                            _ = fi141
                                            _ = v146
                                        }
                                    }
                                case "a":
                                    if false {
                                        for fi147 := 0; fi147 < 13; fi147++ {
                                            var v148 uint8 = 212
                                            var v149 uint16 = uint16(v148)
                                            var v150 uint = uint(v148)
                                            _ = fi147
                                            _ = v148
                                            _ = v149
                                            _ = v150
                                        }
                                        switch -9223372036854775808 {
                                        case 32767:
                                            for fi151 := 0; fi151 < 9; fi151++ {
                                                var v152 string = v2
                                                v153 := GenS3{F0: -514}
                                                if false {
                                                    var v154 int32 = v139
                                                    _ = v154
                                                } else {
                                                    for fi155 := 0; fi155 < 9; fi155++ {
                                                        v156 := GenS17{F0: 100, F1: 1, F2: 100}
                                                        v157 := GenS11{F0: 189}
                                                        _ = fi155
                                                        _ = v156
                                                        _ = v157
                                                    }
                                                    switch 242 {
                                                    case 552:
                                                        var v158 uint16 = 1
                                                        _ = v158
                                                    case 42:
                                                        for fi159 := 0; fi159 < 19; fi159++ {
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            switch "\n\t" {
                                                            case "hello":
                                                                fmt.Println(v125, v152, v139)
                                                            case "foo bar":
                                                                v7 += (-(v7))
                                                            case "a":
                                                                if false {
                                                                    v1 = float64(float32((v1 / ((v125 + 0.1) - float64(1e10)))))
                                                                } else {
                                                                    var v160 int8 = 100
                                                                    _ = v160
                                                                }
                                                            default:
                                                                v152 = ("foo bar" + "a")
                                                            }
                                                            _ = fi159
                                                        }
                                                    case 9223372036854775807:
                                                        v161 := GenS9{F0: 185, F1: fi151, F2: 255}
                                                        _ = v161
                                                    default:
                                                        v162 := GenS18{F0: 1, F1: 1}
                                                        _ = v162
                                                    }
                                                }
                                                _ = fi151
                                                _ = v152
                                                _ = v153
                                            }
                                        case 32768:
                                            for fi163 := 0; fi163 < 20; fi163++ {
                                                v7 += 42
                                                _ = fi163
                                            }
                                        default:
                                            _ = len(v3)
                                        }
                                        var v164 float64 = v1
                                        _ = v164
                                    }
                                default:
                                    v1 = (3.14159 * (0.0 + 3.14159))
                                }
                                _ = v138
                                _ = v139
                            }
                        default:
                            v2 = (v2 + "")
                        }
                    }
                }
            case -9223372036854775808:
                fmt.Println(v7, v1)
            case 32767:
                fmt.Println(v1, v125, v2)
            default:
                for fi165 := 0; fi165 < 13; fi165++ {
                    for fi166 := 0; fi166 < 18; fi166++ {
                        switch v2 {
                        case "":
                            v7 += -1
                        case "a":
                            v167 := GenS14{F0: (v2 + "a"), F1: (fi165 >> ((42) & 63))}
                            _ = v167
                        case "foo bar":
                            v168 := GenS8{F0: 42}
                            _ = v168
                        }
                        _ = fi166
                    }
                    _ = fi165
                }
            }
            _ = v125
        }
        for fi169 := 0; fi169 < 5; fi169++ {
            if (v1 <= ((float64((1e10 - 1e10)) / ((0.0 / (v1 / v1)) * (1e10 * float64((1e10 * 1e10))))) - v1)) {
                fmt.Println(v7)
            }
            switch fi169 {
            case -686:
                v170 := GenS19{F0: v7, F1: 984, F2: fi169}
                _ = v170
            case 311:
                switch "\n\t" {
                case "":
                    for fi171 := 0; fi171 < 9; fi171++ {
                        _ = len(v4)
                        if (18446744073709551615 != 255) {
                            if (false && true) {
                                var v172 uint8 = 242
                                if (-244 >= 1) {
                                    v7 = v7
                                    var v173 uint64 = 2
                                    _ = v173
                                }
                                _ = v172
                            } else {
                                if (!(("" + "") != v2)) {
                                    var v174 int8 = 127
                                    if false {
                                        v1 -= 0.0
                                        var v175 int = fi171
                                        v176 := GenS17{F0: v174, F1: 1, F2: (v174 | (^(v174)))}
                                        _ = v175
                                        _ = v176
                                    } else {
                                        var v177 float64 = 0.0
                                        var v178 uint32 = 0
                                        _ = v177
                                        _ = v178
                                    }
                                    var v179 bool = (int64(v174) > (v7 ^ 0))
                                    if v179 {
                                        v2 = v2
                                    } else {
                                        v180 := GenS3{F0: 9223372036854775807}
                                        var v181 int16 = 42
                                        fmt.Println(v1, fi169)
                                        _ = v180
                                        _ = v181
                                    }
                                    _ = v174
                                    _ = v179
                                } else {
                                    v2 = v2
                                    var v182 float64 = ((((-2.25 / v1) - float64(1e10)) + v1) + (v1 + 3.14159))
                                    _ = v182
                                }
                                v4 = append(v4, 65535)
                            }
                            switch fi169 {
                            case 466:
                                v183 := GenS8{F0: 65535}
                                if ((v7 % (((v7 / ((v7) | 1))) | 1)) != v7) {
                                    var v184 int8 = 127
                                    v184 = (-(v184))
                                    v4 = append(v4, 394)
                                    _ = v184
                                } else {
                                    var v185 float32 = 1.5
                                    var v186 int = (-(fi171))
                                    if true {
                                        if (!(0 <= 2)) {
                                            fmt.Println(v1, v2, fi171)
                                        } else {
                                            if false {
                                                v1 += (float64((v185 / (float32(float64(((((((3.14159 * 0.0) + v185) + float32(v1)) * 3.14159) - v185) - v185))) * v185))) + ((v1 * 0.0) - v1))
                                                fmt.Println(v186, v185, fi171)
                                                fmt.Println(v7)
                                            }
                                            var v187 int64 = v7
                                            _ = v187
                                        }
                                        v188 := v3[(v2 + v2)]
                                        v2 = v2
                                        _ = v188
                                    }
                                    _ = v185
                                    _ = v186
                                }
                                _ = v183
                            }
                            var v189 uint = 683
                            _ = v189
                        } else {
                            v1 -= v1
                            _ = len(v4)
                            var v190 uint32 = 65535
                            _ = v190
                        }
                        if true {
                            switch 0 {
                            case 1:
                                var v191 float64 = 1e10
                                _ = v191
                            case 32768:
                                delete(v3, (v2 + v2))
                            case 9223372036854775807:
                                v3[v2] = v2
                            }
                            if (2 <= 2) {
                                v2 = "\n\t"
                                var v192 uint16 = 2
                                _ = v192
                            } else {
                                v7 -= v7
                                if false {
                                    var v193 string = ""
                                    if (!false) {
                                        v7 -= (^(v7))
                                        var v194 int32 = -1
                                        if true {
                                            fmt.Println(v7)
                                            v7 = 42
                                        } else {
                                            if true {
                                                var v195 float32 = float32(((v1 - (v1 * v1)) + (1.5 / (float64(0.0) * v1))))
                                                if (((1.5 / (((((0.1 / -2.25) + v195) + float32(1e10)) * 1e10) - float32((0.0 - (-2.25 + 0.0))))) - float32(-2.25)) == -2.25) {
                                                    v196 := GenS14{F0: "foo bar", F1: (-(fi171))}
                                                    v195 = v195
                                                    _ = v196
                                                } else {
                                                    fmt.Println(fi171, v2)
                                                    var v197 int = (^(fi169))
                                                    if true {
                                                        if false {
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        }
                                                        if false {
                                                            delete(v3, v2)
                                                        }
                                                    }
                                                    _ = v197
                                                }
                                                var v198 float32 = (v195 - float32((((float64(v195) / v1) / 3.14159) + 0.0)))
                                                if false {
                                                    switch (v2 + (("foo bar" + v193) + (v193 + v193))) {
                                                    case "a":
                                                        switch (^(fi171)) {
                                                        case -2147483648:
                                                            v2 = ""
                                                        }
                                                    default:
                                                        if false {
                                                            switch (fi171 - fi169) {
                                                            case 42:
                                                                if (32767 >= 42) {
                                                                    var v199 string = v193
                                                                    if len(v4) > 0 {
                                                                        v4[0] = 0
                                                                    }
                                                                    _ = v199
                                                                }
                                                            case -1:
                                                                switch "hello" {
                                                                case "hello":
                                                                    if (2147483647 > 42) {
                                                                        v1 -= (v1 - v1)
                                                                    } else {
                                                                        var v200 bool = true
                                                                        _ = v200
                                                                    }
                                                                case "\n\t":
                                                                    fmt.Println(v198, v193, v195)
                                                                case "a":
                                                                    fmt.Println(v1, v194, v1)
                                                                    var v201 int8 = 127
                                                                    v198 -= (v198 / v195)
                                                                    fmt.Println(v193, v201, v7)
                                                                    v202 := GenS11{F0: 42}
                                                                    _ = v201
                                                                    _ = v202
                                                                }
                                                                var v203 int32 = 0
                                                                _ = v203
                                                            default:
                                                                v204 := GenS1{F0: -1, F1: "hello", F2: 300}
                                                                _ = v204
                                                            }
                                                            var v205 bool = false
                                                            v206 := GenS20{F0: v194, F1: (fi169 >> ((964) & 63)), F2: 36}
                                                            _ = v205
                                                            _ = v206
                                                        }
                                                    }
                                                    v193 = ("\n\t" + (v193 + "hello"))
                                                }
                                                _ = v195
                                                _ = v198
                                            }
                                            v1 = v1
                                        }
                                        var v207 uint8 = 104
                                        v207 = v207
                                        _ = v194
                                        _ = v207
                                    }
                                    _ = v193
                                } else {
                                    switch v2 {
                                    case "\n\t":
                                        v2 = (v2 + v2)
                                        v2 = (v2 + v2)
                                    case "hello":
                                        v1 -= (float64(0.0) - 1.5)
                                    }
                                }
                            }
                            v208 := GenS3{F0: v7}
                            _ = v208
                        }
                        _ = fi171
                    }
                case "\n\t":
                    v209 := v3[v2]
                    _ = v209
                default:
                    var v210 int8 = -1
                    _ = v210
                }
            case 9223372036854775807:
                for fi211 := 0; fi211 < 16; fi211++ {
                    v7 = v7
                    _ = fi211
                }
            default:
                var v212 int8 = 127
                _ = v212
            }
            _ = fi169
        }
    }
    v2 = (v2 + v2)
    v3[(v2 + "hello")] = ((v2 + "foo bar") + v2)
    switch "foo bar" {
    case "foo bar":
        if (1 == -1) {
            switch 829 {
            case 32768:
                var v213 int16 = -1
                _ = v213
            }
        }
    }
    v3[(("hello" + "\n\t") + ((v2 + v2) + "foo bar"))] = ("" + ((v2 + "") + (v2 + ((v2 + ("\n\t" + (("\n\t" + ("\n\t" + "")) + v2))) + ""))))
    v3[v2] = ""
    if ((v7 - (^(v7))) < 638) {
        v1 += v1
    }
    v1 = (-2.25 - float64(1.5))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
