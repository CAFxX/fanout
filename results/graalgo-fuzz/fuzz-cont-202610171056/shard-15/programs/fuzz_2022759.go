package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 uint8; F1 int64 }
type GenS4 struct { F0 int8; F1 uint16; F2 int16 }
type GenS5 struct { F0 int32; F1 uint; F2 uint64 }
type GenS6 struct { F0 uint; F1 int8 }
type GenS7 struct { F0 int32; F1 int8; F2 bool }
type GenS8 struct { F0 uint; F1 int32; F2 int64 }
type GenS9 struct { F0 uint32; F1 string; F2 uint }
type GenS10 struct { F0 uint64; F1 uint8; F2 string }
type GenS11 struct { F0 uint16; F1 int64 }
func genFn0() uint {
    var v6 uint32 = 187
    v6 += (v6 % (((^(v6))) | 1))
    _ = v6
    return 2
}

func main() {
    v1 := []string{"a", ("\n\t" + ""), "foo bar", ""}
    var v2 string = ((((("foo bar" + "a") + ("a" + ("hello" + "\n\t"))) + ("foo bar" + ((("\n\t" + ("foo bar" + "")) + "\n\t") + ("" + "foo bar")))) + "hello") + ("a" + ("a" + ("" + "foo bar"))))
    v3 := []uint32{}
    var v4 float32 = float32(float64((3.14159 - (1e10 / (0.0 - 1e10)))))
    v5 := make(map[uint32]string)
    var v7 int = 128
    if true {
        if (("\n\t" + (((("" + v2) + ((v2 + v2) + v2)) + "\n\t") + v2)) == "foo bar") {
            var v8 bool = (1 >= 910)
            fmt.Println(v8, v7, v4)
            switch "\n\t" {
            case "\n\t":
                v9 := GenS1{F0: 32767}
                _ = v9
            }
            _ = v8
        }
    }
    var v10 float32 = v4
    if (!(!false)) {
        if (0.1 == ((-2.25 + 0.0) * 0.0)) {
            fmt.Println(v4)
            v5[4294967295] = v2
            var v11 string = v2
            _ = v11
        }
        var v12 float32 = ((1e10 + (((v4 * float32(1.5)) - v10) * 1.5)) * v10)
        _ = v12
    } else {
        for fi13 := 0; fi13 < 10; fi13++ {
            v14 := GenS2{F0: 51, F1: fi13}
            v4 += 3.14159
            _ = fi13
            _ = v14
        }
        v15 := GenS3{F0: 42, F1: 1}
        if ((v7 < (v7 & (-(v7)))) || true) {
            v16 := v5[2]
            switch ((((v2 + v16) + "foo bar") + v16) + v16) {
            case "a":
                switch v16 {
                case "":
                    fmt.Println(v4, v16, v7)
                case "a":
                    var v17 float64 = (1e10 - (-2.25 - float64(v10)))
                    _ = v17
                case "foo bar":
                    v5[889] = v2
                }
            case "\n\t":
                switch (v7 & v7) {
                case 0:
                    switch "\n\t" {
                    case "":
                        for fi18 := 0; fi18 < 9; fi18++ {
                            delete(v5, 45)
                            v16 = ""
                            _ = fi18
                        }
                    case "foo bar":
                        if true {
                            for fi19 := 0; fi19 < 2; fi19++ {
                                v20 := v5[201]
                                v21 := GenS4{F0: 1, F1: 42, F2: -893}
                                _ = fi19
                                _ = v20
                                _ = v21
                            }
                            var v22 float64 = 0.0
                            _ = v22
                        } else {
                            for fi23 := 0; fi23 < 8; fi23++ {
                                var v24 string = (v16 + "a")
                                if len(v3) > 0 {
                                    v3[0] = 1
                                }
                                switch v2 {
                                case "":
                                    switch v24 {
                                    case "":
                                        v25 := v5[2147483647]
                                        _ = v25
                                    case "\n\t":
                                        if (251 <= 42) {
                                            delete(v5, 4294967295)
                                            v26 := GenS2{F0: 255, F1: (v7 ^ (^(fi23)))}
                                            if len(v1) > 0 {
                                                v1[0] = v24
                                            }
                                            _ = v26
                                        }
                                    case "a":
                                        var v27 uint32 = 1
                                        _ = v27
                                    default:
                                        var v28 float32 = float32(1.5)
                                        _ = v28
                                    }
                                case "a":
                                    switch ("" + (v16 + "")) {
                                    case "":
                                        var v29 uint = 255
                                        _ = v29
                                    case "a":
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                default:
                                    if (1 < 0) {
                                        switch (v7 << ((genFn0()) & 63)) {
                                        case -2147483648:
                                            var v30 uint64 = 4294967295
                                            _ = v30
                                        case 128:
                                            fmt.Println(v7)
                                        default:
                                            var v31 uint64 = 1
                                            _ = v31
                                        }
                                    } else {
                                        for fi32 := 0; fi32 < 13; fi32++ {
                                            v4 -= v4
                                            _ = fi32
                                        }
                                        var v33 int32 = 801
                                        _ = v33
                                    }
                                }
                                _ = fi23
                                _ = v24
                            }
                            switch v7 {
                            case 32768:
                                switch v7 {
                                case 32768:
                                    v34 := GenS5{F0: 1, F1: 19, F2: 1}
                                    _ = v34
                                case 163:
                                    v2 = ""
                                }
                            }
                        }
                    case "hello":
                        for fi35 := 0; fi35 < 7; fi35++ {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            for fi36 := 0; fi36 < 20; fi36++ {
                                if (-1 < 9223372036854775807) {
                                    var v37 uint64 = 2
                                    v38 := GenS6{F0: 270, F1: -36}
                                    _ = v37
                                    _ = v38
                                } else {
                                    v4 = 0.1
                                    v39 := GenS4{F0: 100, F1: 32767, F2: 1000}
                                    _ = v39
                                }
                                var v40 int64 = 9223372036854775807
                                _ = fi36
                                _ = v40
                            }
                            var v41 uint16 = 1
                            _ = fi35
                            _ = v41
                        }
                    default:
                        fmt.Println(v7, v10)
                    }
                case -223:
                    var v42 uint16 = 1
                    _ = v42
                case -1:
                    if (false && (42 < 0)) {
                        if true {
                            if len(v1) > 0 {
                                v1[0] = v2
                            }
                            for fi43 := 0; fi43 < 13; fi43++ {
                                v7 -= 1
                                v44 := GenS7{F0: 1, F1: 1, F2: false}
                                _ = fi43
                                _ = v44
                            }
                        } else {
                            v45 := GenS4{F0: 127, F1: 255, F2: -32768}
                            for fi46 := 0; fi46 < 3; fi46++ {
                                var v47 uint64 = 2
                                if false {
                                    v10 -= float32((3.14159 * 0.0))
                                    v16 = v16
                                }
                                _ = fi46
                                _ = v47
                            }
                            for fi48 := 0; fi48 < 13; fi48++ {
                                var v49 int64 = 42
                                switch "foo bar" {
                                case "a":
                                    if (2 <= 18446744073709551615) {
                                        var v50 uint32 = 2147483648
                                        switch "" {
                                        case "\n\t":
                                            var v51 uint8 = 1
                                            _ = v51
                                        case "foo bar":
                                            for fi52 := 0; fi52 < 7; fi52++ {
                                                var v53 int32 = 42
                                                _ = fi52
                                                _ = v53
                                            }
                                        case "hello":
                                            var v54 uint8 = 1
                                            _ = v54
                                        default:
                                            fmt.Println(v16, fi48, v16)
                                        }
                                        _ = v50
                                    }
                                case "":
                                    if true {
                                        if true {
                                            for fi55 := 0; fi55 < 14; fi55++ {
                                                v56 := GenS8{F0: 255, F1: -865, F2: (v49 | (v49 - (v49 | v49)))}
                                                if true {
                                                    fmt.Println(v4)
                                                    var v57 int8 = 102
                                                    _ = v57
                                                } else {
                                                    if (1 >= 32767) {
                                                        v58 := GenS8{F0: 4294967295, F1: -1, F2: (v49 + (v49 >> ((839) & 63)))}
                                                        if (false && (!((v10 > 0.0) || (v2 != v16)))) {
                                                            switch v16 {
                                                            case "hello":
                                                                var v59 uint8 = 0
                                                                _ = v59
                                                            case "":
                                                                if ((!false) || true) {
                                                                    v10 -= float32(-2.25)
                                                                    var v60 int = (fi55 * (v7 ^ (fi48 - v7)))
                                                                    v49 += (^(v49))
                                                                    _ = v60
                                                                }
                                                            }
                                                            var v61 int = 9223372036854775807
                                                            var v62 int = 2147483647
                                                            _ = v61
                                                            _ = v62
                                                        }
                                                        _ = v58
                                                    } else {
                                                        switch fi48 {
                                                        case -9223372036854775808:
                                                            fmt.Println(v49, fi48)
                                                        }
                                                    }
                                                    fmt.Println(v10)
                                                    fmt.Println(v10, v7, v16)
                                                }
                                                if true {
                                                    if false {
                                                        var v63 uint16 = 1
                                                        if ((fi55 >= (fi48 * int(v49))) && false) {
                                                            var v64 int16 = 1
                                                            var v65 uint = 1
                                                            fmt.Println(fi48)
                                                            _ = v64
                                                            _ = v65
                                                        }
                                                        _ = v63
                                                    } else {
                                                        if false {
                                                            if false {
                                                                v66 := GenS1{F0: 537}
                                                                if (!(-1 == -56)) {
                                                                    fmt.Println(v16, fi55)
                                                                } else {
                                                                    var v67 int16 = 0
                                                                    _ = v67
                                                                }
                                                                v3 = append(v3, 1)
                                                                _ = v66
                                                            } else {
                                                                if true {
                                                                    fmt.Println(v16)
                                                                    v49 -= v49
                                                                } else {
                                                                    var v68 int16 = -743
                                                                    v4 = v4
                                                                    _ = v68
                                                                }
                                                                v16 = v16
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            }
                                                        } else {
                                                            v5[0] = v16
                                                            v7 = (fi55 - fi55)
                                                            var v69 bool = (true && (!(((!(!(833 <= 1))) || (181 >= 175)) || true)))
                                                            _ = v69
                                                        }
                                                        var v70 uint = 562
                                                        _ = v70
                                                    }
                                                    v71 := v5[2147483647]
                                                    _ = v71
                                                }
                                                v72 := GenS9{F0: 1, F1: (v2 + ""), F2: 0}
                                                _ = fi55
                                                _ = v56
                                                _ = v72
                                            }
                                        }
                                        var v73 uint = 4294967295
                                        v74 := GenS6{F0: 255, F1: -1}
                                        _ = v73
                                        _ = v74
                                    } else {
                                        for fi75 := 0; fi75 < 18; fi75++ {
                                            v76 := GenS2{F0: 42, F1: (^(fi48))}
                                            switch (v7 ^ 0) {
                                            case 128:
                                                switch "hello" {
                                                case "":
                                                    var v77 int32 = -1
                                                    _ = v77
                                                case "\n\t":
                                                    v1 = append(v1, ("a" + v2))
                                                default:
                                                    v5[0] = ("" + v2)
                                                }
                                            case 9223372036854775807:
                                                v2 = (v16 + (v16 + ""))
                                            case 1:
                                                switch v2 {
                                                case "":
                                                    v78 := GenS10{F0: 0, F1: 59, F2: v16}
                                                    _ = v78
                                                default:
                                                    if (!true) {
                                                        v4 = float32(1.5)
                                                    }
                                                }
                                            }
                                            if true {
                                                if (0 == 2147483647) {
                                                    v2 = (v16 + v16)
                                                    switch 514 {
                                                    case -1:
                                                        v2 = ""
                                                    case 9223372036854775807:
                                                        v1 = append(v1, v2)
                                                    case -2147483648:
                                                        if true {
                                                            fmt.Println(fi75)
                                                            var v79 float64 = 0.0
                                                            fmt.Println(v4)
                                                            v49 = 883
                                                            _ = v79
                                                        } else {
                                                            var v80 int8 = 42
                                                            _ = v80
                                                        }
                                                    default:
                                                        if (-2.25 == (((((3.14159 / 3.14159) + (-2.25 - (0.0 * (0.1 + float64((0.0 + v10)))))) + 0.1) - 1.5) - 1e10)) {
                                                            switch v7 {
                                                            case 32768:
                                                                var v81 uint16 = 42
                                                                _ = v81
                                                            }
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            v3 = append(v3, 2147483647)
                                                        } else {
                                                            var v82 string = v2
                                                            var v83 int16 = 42
                                                            var v84 uint = 18446744073709551615
                                                            _ = v82
                                                            _ = v83
                                                            _ = v84
                                                        }
                                                    }
                                                }
                                                v85 := GenS9{F0: 92, F1: (v2 + "foo bar"), F2: 1}
                                                var v86 int = 128
                                                var v87 int = fi48
                                                v86 -= -2147483648
                                                _ = v85
                                                _ = v86
                                                _ = v87
                                            }
                                            fmt.Println(v49)
                                            _ = len(v5)
                                            _ = fi75
                                            _ = v76
                                        }
                                        for fi88 := 0; fi88 < 15; fi88++ {
                                            _ = len(v5)
                                            _ = len(v3)
                                            _ = fi88
                                        }
                                        if true {
                                            var v89 string = ((v2 + "\n\t") + v16)
                                            switch "" {
                                            case "hello":
                                                switch v7 {
                                                case 127:
                                                    switch (v89 + v16) {
                                                    case "":
                                                        v3 = append(v3, 935)
                                                    case "\n\t":
                                                        v89 = ((("\n\t" + v16) + (v16 + (v16 + v16))) + v16)
                                                    default:
                                                        fmt.Println(v10, v2)
                                                    }
                                                case 32767:
                                                    fmt.Println(v49)
                                                case 128:
                                                    var v90 uint = 2
                                                    _ = v90
                                                }
                                            case "foo bar":
                                                for fi91 := 0; fi91 < 15; fi91++ {
                                                    var v92 float64 = 3.14159
                                                    v93 := GenS4{F0: -77, F1: 2, F2: 122}
                                                    v89 = "a"
                                                    _ = fi91
                                                    _ = v92
                                                    _ = v93
                                                }
                                            default:
                                                for fi94 := 0; fi94 < 2; fi94++ {
                                                    v4 = v10
                                                    v95 := GenS1{F0: 2}
                                                    v96 := GenS3{F0: 47, F1: (v49 * (v49 >> ((42) & 63)))}
                                                    fmt.Println(v49)
                                                    _ = fi94
                                                    _ = v95
                                                    _ = v96
                                                }
                                            }
                                            _ = len(v1)
                                            _ = v89
                                        } else {
                                            fmt.Println(v4, v49)
                                        }
                                    }
                                case "\n\t":
                                    if (127 != 177) {
                                        if false {
                                            v2 = ("" + "a")
                                            var v97 int16 = 42
                                            _ = v97
                                        } else {
                                            if (-9223372036854775808 >= fi48) {
                                                switch v16 {
                                                case "":
                                                    v3 = append(v3, 0)
                                                case "foo bar":
                                                    var v98 int16 = 32767
                                                    _ = v98
                                                case "hello":
                                                    for fi99 := 0; fi99 < 6; fi99++ {
                                                        _ = len(v3)
                                                        v4 = float32(0.1)
                                                        v4 = v4
                                                        _ = fi99
                                                    }
                                                default:
                                                    switch v7 {
                                                    case 2147483647:
                                                        var v100 string = ((v16 + ("\n\t" + "")) + "foo bar")
                                                        _ = v100
                                                    case 42:
                                                        if false {
                                                            var v101 int8 = 100
                                                            var v102 float64 = 3.14159
                                                            var v103 float32 = v10
                                                            _ = v101
                                                            _ = v102
                                                            _ = v103
                                                        }
                                                    }
                                                }
                                            }
                                        }
                                    }
                                }
                                _ = fi48
                                _ = v49
                            }
                            _ = v45
                        }
                    }
                }
            }
            _ = v16
        } else {
            if (-2.25 <= 3.14159) {
                v10 += 1.5
                var v104 uint32 = 2
                _ = v104
            } else {
                v4 = (v4 - 1.5)
                var v105 int32 = 42
                _ = v105
            }
            v4 = (((v10 / (1.5 / (float32((float64(v4) + -2.25)) * v4))) + float32((0.1 - -2.25))) - v4)
            v7 = v7
        }
        _ = v15
    }
    if (!(532 < 2147483647)) {
        _ = len(v3)
        for fi106 := 0; fi106 < 20; fi106++ {
            for fi107 := 0; fi107 < 11; fi107++ {
                v108 := GenS2{F0: 92, F1: (fi107 % (((v7 * (fi107 - (v7 ^ 9223372036854775807)))) | 1))}
                if (true && true) {
                    fmt.Println(v4, v7, v10)
                }
                if len(v1) > 0 {
                    v1[0] = ("hello" + v2)
                }
                _ = fi107
                _ = v108
            }
            if false {
                if false {
                    v109 := GenS10{F0: 1, F1: 128, F2: (v2 + ((v2 + "foo bar") + (v2 + "a")))}
                    v7 += 128
                    _ = v109
                }
            }
            _ = fi106
        }
    } else {
        v4 += ((float32(0.0) * v4) - 0.0)
    }
    v2 = (v2 + v2)
    var v110 uint16 = 65535
    switch "" {
    case "foo bar":
        fmt.Println(v4, v4, v10)
    case "":
        for fi111 := 0; fi111 < 4; fi111++ {
            for fi112 := 0; fi112 < 19; fi112++ {
                fmt.Println(fi112, v10, v10)
                _ = fi112
            }
            var v113 float32 = v10
            var v114 string = "a"
            _ = fi111
            _ = v113
            _ = v114
        }
    case "hello":
        var v115 int64 = 1
        _ = v115
    default:
        v116 := GenS11{F0: v110, F1: 9223372036854775807}
        _ = v116
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v110:", v110)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v110
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
