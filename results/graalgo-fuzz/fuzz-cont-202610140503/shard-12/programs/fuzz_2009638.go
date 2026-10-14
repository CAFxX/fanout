package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 uint16; F2 int32 }
type GenS3 struct { F0 uint16; F1 string; F2 string }
type GenS4 struct { F0 int64; F1 int8; F2 uint32 }
type GenS5 struct { F0 int8; F1 int32; F2 uint64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int8; F1 int32 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint64; F1 int32; F2 uint32 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 uint8; F1 int8; F2 uint }
type GenS12 struct { F0 uint64; F1 uint8 }
type GenS13 struct { F0 bool; F1 string; F2 int8 }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 uint64; F1 int; F2 uint8 }

func main() {
    v1 := []int{579}
    v2 := []string{"a", "", ("\n\t" + "hello")}
    v3 := make(map[int]uint32)
    var v4 int64 = 800
    v5 := GenS1{F0: 2147483647}
    switch ("" + "\n\t") {
    case "hello":
        if true {
            for fi6 := 0; fi6 < 9; fi6++ {
                if true {
                    var v7 float64 = 0.1
                    _ = v7
                } else {
                    delete(v3, -17)
                    v8 := GenS2{F0: ("\n\t" + ""), F1: 42, F2: 1}
                    v4 = v4
                    _ = v8
                }
                _ = fi6
            }
            delete(v3, 42)
        } else {
            var v9 uint8 = 128
            var v10 int64 = (v4 % (((v4 & v4)) | 1))
            for fi11 := 0; fi11 < 14; fi11++ {
                var v12 int32 = 2147483647
                switch ("a" + "a") {
                case "":
                    v13 := GenS2{F0: "", F1: 320, F2: 215}
                    _ = v13
                case "\n\t":
                    v14 := GenS3{F0: 32767, F1: ("\n\t" + "foo bar"), F2: "\n\t"}
                    _ = v14
                case "foo bar":
                    for fi15 := 0; fi15 < 1; fi15++ {
                        v16 := GenS4{F0: 42, F1: 1, F2: 0}
                        _ = fi15
                        _ = v16
                    }
                default:
                    v12 = v12
                }
                var v17 uint16 = 0
                _ = fi11
                _ = v12
                _ = v17
            }
            _ = v10
            _ = v9
        }
    case "a":
        if false {
            switch 9223372036854775807 {
            case 356:
                v18 := GenS1{F0: 2}
                _ = v18
            default:
                switch "hello" {
                case "hello":
                    v4 -= v4
                case "foo bar":
                    v2 = append(v2, "hello")
                case "a":
                    if (false || ((!true) && (-704 != -1))) {
                        if ((-9223372036854775808 < -9223372036854775808) && true) {
                            for fi19 := 0; fi19 < 19; fi19++ {
                                for fi20 := 0; fi20 < 11; fi20++ {
                                    var v21 uint32 = 2
                                    switch (fi19 % ((fi20) | 1)) {
                                    case -2147483648:
                                        fmt.Println(fi19)
                                    case -154:
                                        v4 -= v4
                                    case 42:
                                        fmt.Println(v21, v4)
                                    default:
                                        if (v4 > v4) {
                                            v4 = (v4 | v4)
                                            var v22 int64 = 42
                                            v4 = (v4 << ((870) & 63))
                                            _ = v22
                                        }
                                    }
                                    v21 = (v21 / ((v21) | 1))
                                    v23 := GenS4{F0: v4, F1: 8, F2: v21}
                                    _ = fi20
                                    _ = v21
                                    _ = v23
                                }
                                _ = fi19
                            }
                            var v24 uint16 = 65535
                            _ = v24
                        } else {
                            v25 := GenS5{F0: -53, F1: -2147483648, F2: 1}
                            _ = v25
                        }
                        if (-2147483648 >= -371) {
                            v4 -= v4
                            for fi26 := 0; fi26 < 20; fi26++ {
                                v27 := GenS3{F0: 412, F1: ("\n\t" + ("\n\t" + ("foo bar" + (("a" + ("\n\t" + ("foo bar" + ("" + "")))) + "")))), F2: ""}
                                v1 = append(v1, fi26)
                                _ = fi26
                                _ = v27
                            }
                        } else {
                            delete(v3, 850)
                            for fi28 := 0; fi28 < 10; fi28++ {
                                if false {
                                    fmt.Println(fi28, fi28)
                                    v4 += v4
                                } else {
                                    v4 = v4
                                    v29 := v3[fi28]
                                    for fi30 := 0; fi30 < 2; fi30++ {
                                        v4 -= v4
                                        switch "" {
                                        case "a":
                                            v31 := v3[(fi30 + (fi30 & fi30))]
                                            _ = v31
                                        case "\n\t":
                                            switch "\n\t" {
                                            case "\n\t":
                                                v4 = v4
                                            case "":
                                                v29 = 2147483648
                                            }
                                            fmt.Println(v29)
                                            var v32 int = fi30
                                            _ = v32
                                        default:
                                            switch "" {
                                            case "a":
                                                var v33 uint32 = (v29 / (((v29 % ((v29) | 1))) | 1))
                                                _ = v33
                                            case "foo bar":
                                                if len(v1) > 0 {
                                                    v1[0] = (fi30 ^ fi30)
                                                }
                                            case "\n\t":
                                                v29 -= v29
                                            default:
                                                var v34 uint = 172
                                                _ = v34
                                            }
                                        }
                                        var v35 uint32 = v29
                                        v1 = append(v1, 1)
                                        _ = fi30
                                        _ = v35
                                    }
                                    _ = v29
                                }
                                _ = fi28
                            }
                        }
                        switch -931 {
                        case -2147483648:
                            _ = len(v3)
                        }
                    } else {
                        v36 := GenS2{F0: "a", F1: 32767, F2: -924}
                        var v37 int = 32768
                        _ = v36
                        _ = v37
                    }
                default:
                    for fi38 := 0; fi38 < 14; fi38++ {
                        v39 := GenS1{F0: 1}
                        if (v4 < (v4 & (v4 + (v4 | (v4 - (v4 << ((255) & 63))))))) {
                            v40 := GenS5{F0: 42, F1: -193, F2: 964}
                            _ = len(v3)
                            for fi41 := 0; fi41 < 5; fi41++ {
                                v4 += 9223372036854775807
                                _ = fi41
                            }
                            _ = v40
                        }
                        if false {
                            if false {
                                switch "\n\t" {
                                case "hello":
                                    v3[fi38] = 2
                                case "foo bar":
                                    v4 -= (v4 ^ v4)
                                }
                            }
                        }
                        _ = fi38
                        _ = v39
                    }
                }
            }
        }
    }
    fmt.Println(v4)
    var v42 uint8 = 43
    switch -662 {
    case 0:
        switch "a" {
        case "hello":
            v43 := GenS6{F0: v42}
            _ = v43
        case "a":
            fmt.Println(v42, v4)
        default:
            fmt.Println(v42)
        }
    case 32767:
        var v44 int = -772
        _ = v44
    case 32768:
        var v45 uint32 = 737
        _ = v45
    default:
        if ((136 == 87) || (1.5 != 3.14159)) {
            for fi46 := 0; fi46 < 13; fi46++ {
                v47 := GenS1{F0: 0}
                v48 := GenS7{F0: -84, F1: 895}
                _ = fi46
                _ = v47
                _ = v48
            }
            if (!(!(true || (!true)))) {
                if true {
                    v42 = (v42 << ((18446744073709551615) & 63))
                    v49 := v3[27]
                    _ = v49
                } else {
                    switch 42 {
                    case 0:
                        var v50 uint = 875
                        _ = v50
                    case -99:
                        var v51 int16 = 32767
                        _ = v51
                    default:
                        if (336 > 438) {
                            switch 360 {
                            case -365:
                                fmt.Println(v42, v42)
                            default:
                                v4 = -522
                            }
                        } else {
                            var v52 uint = 0
                            for fi53 := 0; fi53 < 15; fi53++ {
                                if true {
                                    v4 = (v4 ^ v4)
                                } else {
                                    for fi54 := 0; fi54 < 17; fi54++ {
                                        v52 -= 728
                                        fmt.Println(fi54, fi53, v52)
                                        if false {
                                            var v55 float32 = 0.0
                                            v52 = 1
                                            _ = v55
                                        } else {
                                            switch "\n\t" {
                                            case "\n\t":
                                                var v56 int64 = (v4 >> ((v52) & 63))
                                                _ = v56
                                            case "":
                                                switch "foo bar" {
                                                case "":
                                                    var v57 int64 = int64(v52)
                                                    _ = v57
                                                case "hello":
                                                    switch fi54 {
                                                    case 32768:
                                                        var v58 int32 = 276
                                                        _ = v58
                                                    case -119:
                                                        v59 := GenS8{F0: (v52 / ((v52) | 1))}
                                                        _ = v59
                                                    default:
                                                        fmt.Println(v42)
                                                    }
                                                }
                                            case "a":
                                                var v60 float64 = 1e10
                                                _ = v60
                                            default:
                                                var v61 float64 = 0.0
                                                _ = v61
                                            }
                                            switch "hello" {
                                            case "foo bar":
                                                delete(v3, (fi54 / (((fi54 << ((v52) & 63))) | 1)))
                                            }
                                            if true {
                                                if ((42 <= 127) && (("foo bar" != ("a" + "\n\t")) && (-1 == 599))) {
                                                    fmt.Println(fi53, v52, fi54)
                                                    var v62 uint8 = 223
                                                    if ((v42 / ((v42) | 1)) < (v42 & v62)) {
                                                        v63 := GenS9{F0: 54, F1: 2147483647, F2: 42}
                                                        if (4294967295 > 984) {
                                                            if (!false) {
                                                                v4 += v4
                                                                switch (fi53 << ((v52) & 63)) {
                                                                case -2147483648:
                                                                    var v64 float32 = 3.14159
                                                                    var v65 string = ("\n\t" + "\n\t")
                                                                    _ = v64
                                                                    _ = v65
                                                                case 0:
                                                                    var v66 int64 = (v4 + (v4 | (v4 << (((v52 % ((42) | 1))) & 63))))
                                                                    _ = v66
                                                                case 42:
                                                                    v67 := GenS2{F0: ("a" + ""), F1: 42, F2: 42}
                                                                    v62 += 42
                                                                    v68 := GenS2{F0: "hello", F1: 0, F2: -340}
                                                                    _ = v67
                                                                    _ = v68
                                                                }
                                                                fmt.Println(fi53, v42)
                                                            } else {
                                                                if (false && (!(133 < 65535))) {
                                                                    delete(v3, -887)
                                                                    if (-128 != int8(v62)) {
                                                                        if true {
                                                                            if (-1 <= -32768) {
                                                                                v62 += v62
                                                                                v42 = v42
                                                                                switch (fi54 % ((fi54) | 1)) {
                                                                                case -133:
                                                                                    v62 += v42
                                                                                case -835:
                                                                                    if (false || (-1 != 1)) {
                                                                                        v42 -= v62
                                                                                        fmt.Println(v4)
                                                                                        var v69 string = "hello"
                                                                                        _ = v69
                                                                                    }
                                                                                    var v70 int64 = v4
                                                                                    _ = v70
                                                                                }
                                                                            }
                                                                        }
                                                                        fmt.Println(v62, v42)
                                                                    } else {
                                                                        v71 := GenS10{F0: 2147483648}
                                                                        v4 = (v4 % (((v4 * (v4 | 0))) | 1))
                                                                        if ((-(v4)) == v4) {
                                                                            var v72 uint32 = 265
                                                                            var v73 int64 = 0
                                                                            var v74 int32 = -2147483648
                                                                            _ = v72
                                                                            _ = v73
                                                                            _ = v74
                                                                        } else {
                                                                            v42 = v62
                                                                            v52 = v52
                                                                        }
                                                                        _ = v71
                                                                    }
                                                                    v75 := GenS11{F0: v42, F1: 127, F2: v52}
                                                                    _ = v75
                                                                }
                                                            }
                                                        } else {
                                                            switch fi53 {
                                                            case -850:
                                                                _ = len(v3)
                                                                if (false || ((1e10 - 1.5) <= ((((0.0 / 3.14159) - float32((0.1 - float64((((0.1 / (-2.25 - 3.14159)) + 0.1) - 0.1))))) - (0.1 * (0.1 + ((((-2.25 / 1.0) / 1e10) - 3.14159) * 1e10)))) / (-2.25 / 3.14159)))) {
                                                                    v62 += (v42 * (v42 >> (((v52 & (^(v52)))) & 63)))
                                                                    var v76 uint8 = (v62 >> ((255) & 63))
                                                                    _ = v76
                                                                }
                                                                delete(v3, (-(fi54)))
                                                                v77 := GenS12{F0: 630, F1: 255}
                                                                _ = v77
                                                            case -1:
                                                                v78 := GenS13{F0: (false || true), F1: ("\n\t" + "hello"), F2: 0}
                                                                _ = v78
                                                            case 2147483647:
                                                                v79 := GenS11{F0: 108, F1: -1, F2: v52}
                                                                _ = v79
                                                            default:
                                                                switch (fi53 * (fi54 | fi53)) {
                                                                case -282:
                                                                    v52 -= 4294967295
                                                                }
                                                            }
                                                        }
                                                        delete(v3, fi54)
                                                        _ = v63
                                                    } else {
                                                        if len(v2) > 0 {
                                                            v2[0] = ""
                                                        }
                                                        v80 := GenS14{F0: -679}
                                                        var v81 int16 = 529
                                                        _ = v80
                                                        _ = v81
                                                    }
                                                    var v82 int32 = 420
                                                    _ = v62
                                                    _ = v82
                                                } else {
                                                    v83 := GenS15{F0: 4294967295, F1: fi54, F2: (v42 * 0)}
                                                    _ = v83
                                                }
                                                fmt.Println(v4, fi54)
                                            } else {
                                                delete(v3, fi53)
                                            }
                                            if false {
                                                var v84 float64 = 3.14159
                                                var v85 string = "hello"
                                                v4 += v4
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                v42 = (^(v42))
                                                fmt.Println(v85, v85)
                                                _ = v84
                                                _ = v85
                                            } else {
                                                if false {
                                                    var v86 uint32 = 2
                                                    var v87 int32 = 1
                                                    _ = v86
                                                    _ = v87
                                                }
                                            }
                                            switch "a" {
                                            case "hello":
                                                if (!true) {
                                                    v42 -= v42
                                                }
                                            case "foo bar":
                                                var v88 float32 = float32(-2.25)
                                                _ = v88
                                            case "":
                                                fmt.Println(v4, v52)
                                            }
                                        }
                                        _ = fi54
                                    }
                                    var v89 float32 = 1.5
                                    delete(v3, 4294967295)
                                    _ = v89
                                }
                                _ = fi53
                            }
                            _ = v52
                        }
                    }
                }
            } else {
                for fi90 := 0; fi90 < 5; fi90++ {
                    var v91 int64 = (v4 / ((v4) | 1))
                    for fi92 := 0; fi92 < 18; fi92++ {
                        fmt.Println(fi90, v42, fi92)
                        _ = fi92
                    }
                    _ = fi90
                    _ = v91
                }
                var v93 float64 = 3.14159
                _ = v93
            }
            for fi94 := 0; fi94 < 9; fi94++ {
                if ((((fi94 / ((fi94) | 1)) >= fi94) || true) && ((2147483647 <= -998) && (fi94 > fi94))) {
                    switch "foo bar" {
                    case "\n\t":
                        var v95 int = fi94
                        _ = v95
                    }
                    v96 := GenS11{F0: 1, F1: 100, F2: 742}
                    _ = v96
                }
                _ = len(v3)
                _ = fi94
            }
        } else {
            v42 += (v42 >> ((4294967295) & 63))
        }
    }
    for fi97 := 0; fi97 < 10; fi97++ {
        v4 -= 9223372036854775807
        _ = fi97
    }
    if (247 < -601) {
        var v98 uint32 = 1
        _ = len(v3)
        _ = v98
    } else {
        v42 -= v42
        _ = len(v2)
        if (917 <= 1000) {
            v99 := GenS8{F0: 2}
            var v100 int32 = 803
            _ = v100
            _ = v99
        }
    }
    v4 = v4
    if (-321 <= 0) {
        fmt.Println(v4, v4)
        var v101 uint32 = 2147483647
        v102 := GenS15{F0: 218, F1: 4294967295, F2: (v42 ^ (v42 & 42))}
        _ = v101
        _ = v102
    }
    v103 := GenS15{F0: 68, F1: -9223372036854775808, F2: v42}
    var v104 int16 = -990
    v4 += v4
    v4 = (v4 - 42)
    var v105 bool = false
    for fi106 := 0; fi106 < 7; fi106++ {
        v107 := GenS11{F0: 42, F1: 42, F2: 2}
        fmt.Println(fi106, v105, v42)
        _ = fi106
        _ = v107
    }
    v42 = (v42 ^ 42)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v104:", v104)
    fmt.Println("v105:", v105)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v103
    _ = v104
    _ = v105
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
}
