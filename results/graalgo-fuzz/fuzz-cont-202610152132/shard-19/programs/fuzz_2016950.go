package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 int32; F2 int64 }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 bool; F1 int16; F2 uint8 }
type GenS5 struct { F0 int32; F1 uint32; F2 int16 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint64; F1 bool; F2 int }
type GenS8 struct { F0 string; F1 int16; F2 uint32 }

func main() {
    var v1 int16 = 42
    var v2 uint32 = 853
    v3 := []int{128, 4294967295}
    var v4 int8 = 127
    var v5 float32 = (3.14159 + 1e10)
    for fi6 := 0; fi6 < 17; fi6++ {
        v1 = v1
        _ = fi6
    }
    var v7 uint16 = 255
    for fi8 := 0; fi8 < 9; fi8++ {
        switch int(v4) {
        case 1:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        v9 := GenS1{F0: fi8}
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi8
        _ = v9
    }
    fmt.Println(v5, v1)
    if false {
        v3 = append(v3, -523)
        v10 := GenS2{F0: 2, F1: -2147483648, F2: 0}
        _ = v10
    }
    fmt.Println(v4, v4, v5)
    var v11 bool = false
    v1 -= v1
    var v12 float64 = (3.14159 + 3.14159)
    for fi13 := 0; fi13 < 5; fi13++ {
        v7 += v7
        switch (fi13 / (((fi13 & 42)) | 1)) {
        case -501:
            switch fi13 {
            case 127:
                v14 := GenS3{F0: (v4 << ((4294967295) & 63)), F1: v1}
                _ = v14
            case 0:
                fmt.Println(v7, v1)
            default:
                var v15 uint32 = (v2 % ((v2) | 1))
                _ = v15
            }
        case -369:
            var v16 bool = true
            _ = v16
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi13
    }
    var v17 uint8 = 255
    switch -86 {
    case 0:
        if (!(!(false && true))) {
            switch 9223372036854775807 {
            case 678:
                for fi18 := 0; fi18 < 5; fi18++ {
                    v19 := GenS2{F0: 2147483648, F1: -1, F2: 9223372036854775807}
                    if true {
                        switch (fi18 + 660) {
                        case 4294967295:
                            var v20 uint = 255
                            _ = v20
                        case 127:
                            for fi21 := 0; fi21 < 11; fi21++ {
                                if (!((v5 - (1.5 * (v5 * 0.1))) != v5)) {
                                    _ = len(v3)
                                    if true {
                                        if (!(!((fi21 >> ((uint(v17)) & 63)) >= (-(fi21))))) {
                                            if v11 {
                                                fmt.Println(v11)
                                                if ((^(fi21)) >= fi18) {
                                                    switch ("\n\t" + (("foo bar" + "hello") + "")) {
                                                    case "foo bar":
                                                        var v22 int8 = (v4 / ((v4) | 1))
                                                        _ = v22
                                                    }
                                                    fmt.Println(v4)
                                                }
                                            }
                                        } else {
                                            v2 -= (^(v2))
                                            v3 = append(v3, fi18)
                                            if len(v3) > 0 {
                                                v3[0] = fi21
                                            }
                                        }
                                        if v11 {
                                            if true {
                                                var v23 uint16 = (v7 % ((65535) | 1))
                                                _ = v23
                                            }
                                        } else {
                                            var v24 int8 = (-(v4))
                                            _ = v24
                                        }
                                        v25 := GenS3{F0: -66, F1: v1}
                                        switch fi18 {
                                        case 968:
                                            v12 += v12
                                        case 4294967295:
                                            v5 -= (v5 * ((float32(-2.25) * 3.14159) * v5))
                                        default:
                                            var v26 int64 = 42
                                            _ = v26
                                        }
                                        if v11 {
                                            switch "" {
                                            case "foo bar":
                                                switch ("foo bar" + "a") {
                                                case "":
                                                    v12 += v12
                                                case "hello":
                                                    v17 += (^(v17))
                                                case "\n\t":
                                                    var v27 float32 = v5
                                                    _ = v27
                                                }
                                            case "\n\t":
                                                var v28 uint32 = 1
                                                _ = v28
                                            default:
                                                var v29 uint8 = (v17 ^ 127)
                                                _ = v29
                                            }
                                            if ((v11 && (42 == 867)) || (v12 < ((v12 * v12) * 0.1))) {
                                                v12 = v12
                                                var v30 int8 = (^(v4))
                                                var v31 int32 = int32(v30)
                                                _ = v30
                                                _ = v31
                                            }
                                            fmt.Println(v1, v12)
                                        } else {
                                            if v11 {
                                                v32 := GenS2{F0: 4294967295, F1: 753, F2: 0}
                                                var v33 int64 = -1
                                                fmt.Println(v7)
                                                _ = v32
                                                _ = v33
                                            } else {
                                                var v34 uint8 = v17
                                                _ = v34
                                            }
                                        }
                                        v35 := GenS4{F0: (!(("hello" + "a") > (("" + ("" + "foo bar")) + ("hello" + "a")))), F1: (v1 % (((v1 / (((v1 - (v1 + v1))) | 1))) | 1)), F2: v17}
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        v36 := GenS1{F0: fi21}
                                        switch (("hello" + "hello") + "hello") {
                                        case "foo bar":
                                            if v11 {
                                                v4 += (v4 * v4)
                                                if true {
                                                    if (v11 || ((0 <= 166) || ((0 != 245) && (!((42 == 42) && false))))) {
                                                        switch (("a" + "foo bar") + ("\n\t" + "hello")) {
                                                        case "\n\t":
                                                            v37 := GenS2{F0: 2147483647, F1: -979, F2: 1}
                                                            _ = v37
                                                        case "foo bar":
                                                            fmt.Println(v1, v7)
                                                        default:
                                                            _ = len(v3)
                                                        }
                                                        var v38 float64 = v12
                                                        v3 = append(v3, 926)
                                                        _ = v38
                                                    }
                                                    var v39 int8 = -128
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    _ = v39
                                                } else {
                                                    var v40 int32 = 121
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    _ = v40
                                                }
                                            }
                                        case "hello":
                                            v2 -= (v2 & v2)
                                        default:
                                            v2 = v2
                                        }
                                        _ = v25
                                        _ = v35
                                        _ = v36
                                    }
                                    var v41 float32 = (0.1 + v5)
                                    v7 -= v7
                                    var v42 int64 = 0
                                    if (v1 >= v1) {
                                        var v43 uint8 = v17
                                        var v44 uint16 = (v7 & (v7 | v7))
                                        v45 := GenS2{F0: v2, F1: 42, F2: (v42 & (v42 | v42))}
                                        _ = v43
                                        _ = v44
                                        _ = v45
                                    }
                                    var v46 int64 = v42
                                    if v11 {
                                        fmt.Println(v46, v4)
                                    }
                                    v47 := GenS5{F0: 42, F1: (v2 ^ 2147483647), F2: 0}
                                    _ = v41
                                    _ = v42
                                    _ = v46
                                    _ = v47
                                }
                                var v48 uint32 = 42
                                _ = fi21
                                _ = v48
                            }
                        case 448:
                            if true {
                                if false {
                                    for fi49 := 0; fi49 < 16; fi49++ {
                                        var v50 int = (-(fi18))
                                        _ = fi49
                                        _ = v50
                                    }
                                }
                                for fi51 := 0; fi51 < 16; fi51++ {
                                    v52 := GenS6{F0: v4}
                                    _ = fi51
                                    _ = v52
                                }
                            } else {
                                var v53 uint64 = 935
                                _ = v53
                            }
                        }
                    } else {
                        if len(v3) > 0 {
                            v3[0] = fi18
                        }
                        var v54 int32 = 42
                        if ((v4 / ((127) | 1)) == (v4 * int8(fi18))) {
                            if ((v7 % ((32767) | 1)) < v7) {
                                for fi55 := 0; fi55 < 17; fi55++ {
                                    fmt.Println(v1, v54, v4)
                                    _ = fi55
                                }
                                v3 = append(v3, (fi18 << ((65535) & 63)))
                                v5 = -2.25
                            }
                            if (v7 != 65535) {
                                if (uint8(v7) != (v17 / ((126) | 1))) {
                                    var v56 int32 = 71
                                    v1 = (v1 << ((0) & 63))
                                    _ = v56
                                } else {
                                    v57 := GenS5{F0: v54, F1: (v2 + 2), F2: v1}
                                    var v58 uint32 = v2
                                    _ = v57
                                    _ = v58
                                }
                                var v59 bool = true
                                switch "" {
                                case "\n\t":
                                    v12 += (3.14159 / v12)
                                case "":
                                    if v59 {
                                        _ = len(v3)
                                        v60 := GenS1{F0: 0}
                                        _ = v60
                                    } else {
                                        var v61 int16 = 1
                                        var v62 uint16 = v7
                                        for fi63 := 0; fi63 < 10; fi63++ {
                                            if false {
                                                var v64 int16 = v61
                                                var v65 string = ((("hello" + "a") + "a") + "\n\t")
                                                _ = v64
                                                _ = v65
                                            }
                                            v66 := GenS7{F0: 497, F1: (-32768 >= v61), F2: (fi18 + fi63)}
                                            fmt.Println(v2, v61)
                                            _ = fi63
                                            _ = v66
                                        }
                                        _ = v61
                                        _ = v62
                                    }
                                default:
                                    for fi67 := 0; fi67 < 2; fi67++ {
                                        if v11 {
                                            if len(v3) > 0 {
                                                v3[0] = 4294967295
                                            }
                                            var v68 int16 = 0
                                            if (("foo bar" + "a") != "foo bar") {
                                                v59 = (1 > 2)
                                            }
                                            _ = v68
                                        } else {
                                            var v69 uint = 1
                                            _ = v69
                                        }
                                        var v70 uint = 0
                                        if v11 {
                                            if false {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                switch ("" + ("" + "\n\t")) {
                                                case "a":
                                                    v70 = v70
                                                case "foo bar":
                                                    fmt.Println(v59, v7)
                                                case "":
                                                    var v71 int16 = (v1 & 42)
                                                    _ = v71
                                                default:
                                                    var v72 float64 = -2.25
                                                    _ = v72
                                                }
                                                v7 += (v7 / (((v7 / ((71) | 1))) | 1))
                                            }
                                            var v73 bool = (v5 == v5)
                                            _ = v73
                                        } else {
                                            v4 = v4
                                            var v74 bool = (v17 <= v17)
                                            if false {
                                                v75 := GenS8{F0: "foo bar", F1: v1, F2: (v2 - v2)}
                                                v7 -= 255
                                                if true {
                                                    switch "a" {
                                                    case "":
                                                        fmt.Println(v12)
                                                        v76 := GenS6{F0: 100}
                                                        v2 = v2
                                                        _ = v76
                                                    case "foo bar":
                                                        switch fi67 {
                                                        case 9223372036854775807:
                                                            var v77 uint64 = 18446744073709551615
                                                            _ = v77
                                                        default:
                                                            v1 -= (v1 & int16(v17))
                                                        }
                                                    case "hello":
                                                        var v78 uint = 255
                                                        _ = v78
                                                    }
                                                    var v79 float32 = (v5 - 0.0)
                                                    switch (fi67 << ((v70) & 63)) {
                                                    case -917:
                                                        v70 -= (^(v70))
                                                    case 0:
                                                        v7 += v7
                                                    default:
                                                        if ((!(1 != 1)) || v11) {
                                                            v59 = ((v2 / ((v2) | 1)) != 2147483647)
                                                            v80 := GenS4{F0: false, F1: int16(v54), F2: v17}
                                                            _ = v80
                                                        }
                                                    }
                                                    _ = v79
                                                }
                                                _ = v75
                                            }
                                            if (0 == v4) {
                                                if (-128 < int8(v54)) {
                                                    var v81 int64 = 847
                                                    _ = v81
                                                } else {
                                                    var v82 uint16 = (v7 / (((^(v7))) | 1))
                                                    _ = v82
                                                }
                                                var v83 float64 = 0.0
                                                _ = v83
                                            }
                                            _ = v74
                                        }
                                        _ = fi67
                                        _ = v70
                                    }
                                }
                                _ = v59
                            }
                        } else {
                            v17 += (v17 ^ 116)
                            v12 = v12
                            var v84 float32 = 0.1
                            _ = v84
                        }
                        _ = v54
                    }
                    _ = fi18
                    _ = v19
                }
            }
            if len(v3) > 0 {
                v3[0] = 385
            }
        }
    case 4294967295:
        fmt.Println(v1, v1, v7)
    }
    v12 = v12
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
