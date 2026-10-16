package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 int; F1 uint64; F2 int8 }
type GenS5 struct { F0 uint8; F1 int32 }
type GenS6 struct { F0 uint64; F1 uint16 }
type GenS7 struct { F0 uint16; F1 int; F2 int16 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 string; F1 int16; F2 int32 }
type GenS10 struct { F0 bool; F1 int8; F2 uint64 }
type GenS11 struct { F0 int8; F1 uint }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 uint32; F1 int32 }
type GenS14 struct { F0 int8; F1 string }
type GenS15 struct { F0 int64; F1 int32 }
type GenS16 struct { F0 int8; F1 uint64 }

func main() {
    var v1 float64 = (0.1 + 1e10)
    v2 := []int{32767, -2147483648}
    var v3 int32 = -872
    v4 := []uint32{42}
    v5 := make(map[string]int)
    v6 := GenS1{F0: -1}
    v7 := GenS2{F0: 4294967295, F1: 4294967295}
    v8 := GenS3{F0: 2, F1: 2147483647}
    for fi9 := 0; fi9 < 6; fi9++ {
        v10 := GenS2{F0: 2, F1: 18446744073709551615}
        if false {
            v1 = v1
            var v11 uint64 = 2
            _ = v11
        }
        _ = fi9
        _ = v10
    }
    var v12 int = 265
    for fi13 := 0; fi13 < 15; fi13++ {
        v3 -= -1
        var v14 uint16 = 2
        fmt.Println(v14)
        _ = fi13
        _ = v14
    }
    if false {
        switch ("hello" + ("" + "\n\t")) {
        case "foo bar":
            if true {
                var v15 bool = false
                _ = v15
            }
        case "":
            for fi16 := 0; fi16 < 8; fi16++ {
                fmt.Println(v12, v3, v12)
                for fi17 := 0; fi17 < 13; fi17++ {
                    switch (("hello" + "foo bar") + "") {
                    case "hello":
                        if true {
                            var v18 int8 = -128
                            _ = v18
                        } else {
                            if (-1 < 42) {
                                v19 := GenS4{F0: (fi17 | (^(fi16))), F1: 2, F2: 0}
                                var v20 string = ("foo bar" + "a")
                                fmt.Println(fi16, v3, v20)
                                _ = v19
                                _ = v20
                            }
                        }
                    case "foo bar":
                        var v21 int64 = -863
                        _ = v21
                    }
                    switch fi16 {
                    case 4294967295:
                        v12 = v12
                    case -601:
                        if (!((true || (750 >= 992)) || false)) {
                            v12 = fi16
                            _ = len(v5)
                            if (!(false || true)) {
                                v2 = append(v2, 42)
                                if false {
                                    v3 = v3
                                    switch (v12 - v12) {
                                    case -2147483648:
                                        if true {
                                            if false {
                                                var v22 int16 = 0
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                var v23 int32 = (v3 * v3)
                                                _ = v22
                                                _ = v23
                                            } else {
                                                var v24 uint = 4294967295
                                                _ = v24
                                            }
                                            delete(v5, "hello")
                                            var v25 int = (^(fi16))
                                            _ = v25
                                        } else {
                                            v26 := GenS5{F0: 128, F1: 1}
                                            fmt.Println(fi16, fi16)
                                            var v27 float32 = 1.5
                                            v1 = -2.25
                                            var v28 int32 = (v3 | 997)
                                            _ = v26
                                            _ = v27
                                            _ = v28
                                        }
                                        if true {
                                            switch (v12 * 128) {
                                            case -566:
                                                switch 832 {
                                                case -465:
                                                    if false {
                                                        var v29 int32 = v3
                                                        var v30 int32 = (v3 >> ((1) & 63))
                                                        switch fi17 {
                                                        case -9223372036854775808:
                                                            v31 := GenS3{F0: 127, F1: -986}
                                                            v29 += v29
                                                            _ = v31
                                                        }
                                                        fmt.Println(v29, v3)
                                                        v29 += v3
                                                        _ = v29
                                                        _ = v30
                                                    }
                                                default:
                                                    v32 := GenS6{F0: 2, F1: 512}
                                                    _ = v32
                                                }
                                            case -766:
                                                v3 -= (v3 - 2147483647)
                                            default:
                                                var v33 string = ("a" + "foo bar")
                                                _ = v33
                                            }
                                        }
                                    case 4294967295:
                                        v34 := GenS6{F0: 823, F1: 0}
                                        _ = v34
                                    case -1:
                                        v35 := GenS5{F0: 255, F1: v3}
                                        _ = v35
                                    }
                                    fmt.Println(fi17, v12)
                                }
                                var v36 uint64 = 42
                                _ = v36
                            }
                        }
                    case 1:
                        v12 = fi17
                        _ = len(v4)
                        if false {
                            v37 := GenS7{F0: 255, F1: v12, F2: -21}
                            _ = v37
                        }
                        var v38 int32 = (v3 ^ v3)
                        _ = v38
                    default:
                        v12 = (fi16 >> ((1) & 63))
                    }
                    _ = fi17
                }
                _ = fi16
            }
        default:
            v1 = (v1 / -2.25)
        }
        _ = len(v5)
        if false {
            fmt.Println(v3)
        } else {
            var v39 int64 = 1
            if (1 <= 1) {
                if true {
                    fmt.Println(v39, v39, v12)
                    var v40 int8 = int8(v39)
                    _ = v40
                }
            }
            v41 := GenS1{F0: v39}
            _ = v39
            _ = v41
        }
    }
    if (!(!(true && true))) {
        v42 := GenS8{F0: 2}
        var v43 float64 = v1
        var v44 int32 = v3
        _ = v42
        _ = v43
        _ = v44
    }
    if false {
        if (!false) {
            var v45 int16 = 821
            _ = v45
        } else {
            v3 = (v3 % ((v3) | 1))
            if ((568 <= 65535) && ((707 == 368) || false)) {
                var v46 int8 = 42
                if (!(("foo bar" == "") || (false || true))) {
                    switch "" {
                    case "hello":
                        v47 := GenS2{F0: 18446744073709551615, F1: 4294967295}
                        _ = v47
                    default:
                        if (-31 < (^(v46))) {
                            v48 := GenS7{F0: 2, F1: 0, F2: 32767}
                            v46 -= -53
                            _ = v48
                        }
                    }
                    fmt.Println(v1, v12, v12)
                } else {
                    v49 := GenS9{F0: "", F1: -32768, F2: v3}
                    v46 += v46
                    for fi50 := 0; fi50 < 16; fi50++ {
                        if (!false) {
                            switch (-(v12)) {
                            case 32768:
                                v1 -= (float64(-2.25) / v1)
                            case 32767:
                                for fi51 := 0; fi51 < 1; fi51++ {
                                    var v52 float64 = v1
                                    _ = fi51
                                    _ = v52
                                }
                            }
                            delete(v5, "")
                        }
                        _ = fi50
                    }
                    _ = v49
                }
                _ = v46
            } else {
                var v53 int16 = 42
                fmt.Println(v53)
                _ = v53
            }
            if (!false) {
                v54 := GenS10{F0: true, F1: -71, F2: 2}
                if (475 >= -1) {
                    fmt.Println(v12)
                } else {
                    if true {
                        if (-332 >= 1) {
                            if true {
                                if true {
                                    switch v12 {
                                    case 2147483647:
                                        v55 := GenS4{F0: (v12 * (v12 / ((32768) | 1))), F1: 934, F2: 53}
                                        _ = v55
                                    case 127:
                                        for fi56 := 0; fi56 < 17; fi56++ {
                                            var v57 uint64 = 0
                                            _ = fi56
                                            _ = v57
                                        }
                                    case 42:
                                        var v58 float32 = (1.5 + 0.1)
                                        _ = v58
                                    }
                                    if (!true) {
                                        fmt.Println(v1, v3, v12)
                                    } else {
                                        v12 -= v12
                                    }
                                }
                                _ = len(v4)
                            } else {
                                var v59 uint8 = 127
                                for fi60 := 0; fi60 < 6; fi60++ {
                                    for fi61 := 0; fi61 < 16; fi61++ {
                                        fmt.Println(fi61)
                                        _ = fi61
                                    }
                                    var v62 uint64 = 42
                                    _ = fi60
                                    _ = v62
                                }
                                v2 = append(v2, -2147483648)
                                _ = v59
                            }
                        } else {
                            if false {
                                for fi63 := 0; fi63 < 5; fi63++ {
                                    if (false || (!(!(!false)))) {
                                        fmt.Println(fi63, v1, v12)
                                    } else {
                                        for fi64 := 0; fi64 < 8; fi64++ {
                                            v12 -= (fi64 + 136)
                                            var v65 int32 = 926
                                            _ = fi64
                                            _ = v65
                                        }
                                        for fi66 := 0; fi66 < 7; fi66++ {
                                            _ = len(v4)
                                            if true {
                                                v67 := GenS4{F0: (fi66 + 32767), F1: 42, F2: 83}
                                                _ = v67
                                            } else {
                                                if (687 > (fi66 << ((2) & 63))) {
                                                    if (false || (!false)) {
                                                        fmt.Println(v1)
                                                    } else {
                                                        var v68 float32 = (0.0 * (-2.25 / ((0.0 / -2.25) / (0.1 / (float32(v1) - -2.25)))))
                                                        v1 = float64((v68 + ((v68 + (v68 - v68)) - ((((((((3.14159 - (v68 - v68)) / v68) / v68) + float32((v1 + ((v1 * (3.14159 - 1.5)) * v1)))) * 1.5) * 1e10) / (0.1 + 0.0)) + v68))))
                                                        _ = v68
                                                    }
                                                    var v69 int32 = v3
                                                    if ((int16(v69) == 1000) && (128 != 203)) {
                                                        fmt.Println(v3, v3, fi66)
                                                        var v70 int = (fi66 ^ v12)
                                                        _ = v70
                                                    } else {
                                                        v3 = -330
                                                        fmt.Println(v12)
                                                        var v71 uint64 = 81
                                                        _ = v71
                                                    }
                                                    _ = v69
                                                }
                                                v72 := GenS6{F0: 1, F1: 709}
                                                _ = v72
                                            }
                                            if false {
                                                if (696 >= 2) {
                                                    v5["a"] = 0
                                                    v3 -= (-(v3))
                                                } else {
                                                    fmt.Println(fi63)
                                                    if (true || (!true)) {
                                                        v12 -= fi66
                                                        v73 := GenS11{F0: -33, F1: 65535}
                                                        switch 1 {
                                                        case 127:
                                                            v12 += fi66
                                                        }
                                                        _ = v73
                                                    }
                                                }
                                            } else {
                                                v12 = 9223372036854775807
                                                switch fi63 {
                                                case 881:
                                                    fmt.Println(fi66, fi66, fi63)
                                                default:
                                                    var v74 int8 = -83
                                                    _ = v74
                                                }
                                            }
                                            _ = fi66
                                        }
                                        v75 := GenS12{F0: 185}
                                        _ = v75
                                    }
                                    if (v3 < -2147483648) {
                                        for fi76 := 0; fi76 < 13; fi76++ {
                                            if false {
                                                v77 := GenS13{F0: 1, F1: -1}
                                                switch (("hello" + ("foo bar" + "hello")) + "hello") {
                                                case "foo bar":
                                                    var v78 uint8 = 1
                                                    _ = v78
                                                case "":
                                                    fmt.Println(v12)
                                                }
                                                var v79 int16 = 527
                                                _ = v77
                                                _ = v79
                                            } else {
                                                if (100 >= -31) {
                                                    v12 -= 273
                                                    v1 -= 1.5
                                                } else {
                                                    fmt.Println(v1, fi76, v3)
                                                    var v80 uint16 = 111
                                                    _ = v80
                                                }
                                                var v81 uint = 42
                                                _ = v81
                                            }
                                            if ((((-2.25 / (1e10 * 0.1)) - ((1e10 + 3.14159) - float32((v1 / ((v1 / float64(1e10)) * -2.25))))) + (((3.14159 - (-2.25 - 1e10)) + (1e10 * (3.14159 - (0.0 * (3.14159 * float32(v1)))))) * (1.5 / -2.25))) > float32(((v1 + (-2.25 / 0.1)) / (float64(0.1) * v1)))) {
                                                v4 = append(v4, 4294967295)
                                                v5["a"] = (fi76 >> ((0) & 63))
                                                _ = len(v5)
                                            } else {
                                                v3 = (v3 + -1)
                                                if ((true && false) && true) {
                                                    if (42 == 0) {
                                                        var v82 uint16 = 533
                                                        var v83 int = 32768
                                                        v12 -= v83
                                                        _ = v82
                                                        _ = v83
                                                    }
                                                } else {
                                                    v84 := GenS14{F0: -64, F1: "\n\t"}
                                                    _ = v84
                                                }
                                            }
                                            v85 := GenS15{F0: -724, F1: (v3 * v3)}
                                            _ = fi76
                                            _ = v85
                                        }
                                        var v86 uint64 = 4294967295
                                        v12 += (fi63 * v12)
                                        _ = v86
                                    } else {
                                        for fi87 := 0; fi87 < 11; fi87++ {
                                            delete(v5, "foo bar")
                                            _ = fi87
                                        }
                                        var v88 uint64 = 42
                                        _ = v88
                                    }
                                    _ = fi63
                                }
                                v89 := GenS16{F0: -128, F1: 1}
                                if true {
                                    v12 += v12
                                    switch ("foo bar" + "a") {
                                    case "":
                                        v1 = 0.0
                                    }
                                } else {
                                    var v90 int32 = v3
                                    switch "\n\t" {
                                    case "hello":
                                        var v91 uint64 = 18446744073709551615
                                        _ = v91
                                    case "\n\t":
                                        v12 -= (v12 >> ((1) & 63))
                                    case "foo bar":
                                        var v92 int8 = 1
                                        _ = v92
                                    }
                                    delete(v5, "hello")
                                    _ = v90
                                }
                                _ = v89
                            }
                            for fi93 := 0; fi93 < 14; fi93++ {
                                _ = len(v5)
                                _ = fi93
                            }
                            v3 = 1
                        }
                        var v94 int32 = (v3 / (((v3 / (((v3 | v3)) | 1))) | 1))
                        _ = v94
                    } else {
                        v3 += (v3 & v3)
                        fmt.Println(v1, v3)
                        if false {
                            switch (v12 * (v12 - v12)) {
                            case 9223372036854775807:
                                if true {
                                    for fi95 := 0; fi95 < 6; fi95++ {
                                        for fi96 := 0; fi96 < 17; fi96++ {
                                            var v97 float32 = 1.5
                                            _ = fi96
                                            _ = v97
                                        }
                                        var v98 int32 = v3
                                        fmt.Println(v98, v3)
                                        _ = fi95
                                        _ = v98
                                    }
                                    var v99 uint32 = 4294967295
                                    if (!true) {
                                        _ = len(v4)
                                    }
                                    _ = v99
                                } else {
                                    delete(v5, "foo bar")
                                }
                            case -9223372036854775808:
                                v1 -= (v1 / -2.25)
                            }
                            var v100 string = ""
                            _ = v100
                        } else {
                            if true {
                                for fi101 := 0; fi101 < 7; fi101++ {
                                    v3 += v3
                                    v3 = -2147483648
                                    _ = fi101
                                }
                            }
                        }
                    }
                }
                v12 -= v12
                _ = v54
            }
        }
    }
    v1 = float64(-2.25)
    switch "foo bar" {
    case "foo bar":
        var v102 float64 = (v1 - 1.5)
        _ = v102
    case "":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    default:
        v103 := v5[""]
        _ = v103
    }
    var v104 uint64 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v104:", v104)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v104
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
