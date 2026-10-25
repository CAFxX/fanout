package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 uint; F2 bool }
type GenS4 struct { F0 int8; F1 string; F2 int64 }
type GenS5 struct { F0 string; F1 int32; F2 bool }
type GenS6 struct { F0 bool; F1 uint; F2 bool }
type GenS7 struct { F0 uint; F1 bool }
type GenS8 struct { F0 uint32; F1 int16 }
type GenS9 struct { F0 int64; F1 bool }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int8; F1 int32; F2 int32 }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 uint64; F1 uint32 }

func main() {
    var v1 float32 = float32(1e10)
    v2 := []string{"foo bar"}
    v3 := []uint32{2, 791, 42, 2147483648}
    for fi4 := 0; fi4 < 8; fi4++ {
        for fi5 := 0; fi5 < 10; fi5++ {
            switch fi4 {
            case 4294967295:
                if true {
                    switch 0 {
                    case -345:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    default:
                        fmt.Println(fi4, fi4, fi5)
                    }
                    fmt.Println(v1)
                } else {
                    fmt.Println(v1, fi5)
                    if (96 < 42) {
                        fmt.Println(v1, v1)
                        v1 = float32(0.1)
                        fmt.Println(v1, fi4, v1)
                    } else {
                        v6 := GenS1{F0: fi5, F1: 32767}
                        if (false && true) {
                            v1 = (3.14159 - (1.5 + (v1 / (((((v1 / (v1 - 1.5)) * 0.1) - v1) + (float32((-2.25 * 0.0)) - float32(1.5))) + v1))))
                            v2 = append(v2, (((("hello" + "\n\t") + "a") + "foo bar") + "\n\t"))
                            fmt.Println(fi5, v1)
                        } else {
                            if (!(!true)) {
                                fmt.Println(fi4, fi5, v1)
                                var v7 uint32 = 2
                                _ = v7
                            }
                            if (1 <= 80) {
                                v8 := GenS2{F0: 32}
                                v1 = -2.25
                                v1 = ((v1 / float32(0.1)) + (((v1 / v1) / v1) / v1))
                                _ = v8
                            } else {
                                var v9 uint32 = 0
                                var v10 bool = true
                                _ = v10
                                _ = v9
                            }
                            var v11 float32 = v1
                            _ = v11
                        }
                        v1 = 1.5
                        if (42 >= 42) {
                            if (-2147483648 == 0) {
                                switch (fi5 % (((fi5 * (fi5 | fi5))) | 1)) {
                                case -1:
                                    fmt.Println(fi5)
                                case 32768:
                                    if len(v2) > 0 {
                                        v2[0] = ("\n\t" + "a")
                                    }
                                case 4294967295:
                                    var v12 uint64 = 212
                                    _ = v12
                                default:
                                    var v13 int8 = -128
                                    _ = v13
                                }
                                fmt.Println(fi4, v1)
                                v1 -= 1e10
                            }
                        }
                        _ = v6
                    }
                }
            case 42:
                var v14 float32 = ((((v1 * v1) + (0.1 * float32((((3.14159 + float64(v1)) * 0.0) * 0.0)))) - (v1 - 1.5)) + 1e10)
                _ = v14
            case -9223372036854775808:
                var v15 uint8 = 127
                _ = v15
            default:
                if len(v3) > 0 {
                    v3[0] = 756
                }
            }
            var v16 uint64 = 774
            _ = fi5
            _ = v16
        }
        switch (fi4 - (fi4 + 1)) {
        case 2147483647:
            var v17 int8 = 127
            _ = v17
        case 127:
            var v18 uint16 = 42
            _ = v18
        default:
            for fi19 := 0; fi19 < 16; fi19++ {
                fmt.Println(v1, v1)
                _ = fi19
            }
        }
        _ = fi4
    }
    var v20 int32 = -929
    switch -1 {
    case 0:
        var v21 int32 = (v20 * v20)
        _ = v21
    case -1:
        for fi22 := 0; fi22 < 1; fi22++ {
            var v23 int8 = -128
            v24 := GenS1{F0: fi22, F1: 108}
            if true {
                fmt.Println(v1)
                for fi25 := 0; fi25 < 20; fi25++ {
                    if (511 != -245) {
                        switch fi22 {
                        case -549:
                            var v26 string = "a"
                            _ = v26
                        case -1:
                            if false {
                                if true {
                                    if (!false) {
                                        var v27 int = (fi22 << ((4294967295) & 63))
                                        _ = v27
                                    }
                                    var v28 int16 = 0
                                    var v29 int32 = v20
                                    _ = v28
                                    _ = v29
                                } else {
                                    fmt.Println(fi25, v1)
                                    var v30 uint16 = 1
                                    if true {
                                        if (!false) {
                                            if (697 > 4294967295) {
                                                v31 := GenS2{F0: 102}
                                                v1 -= (((3.14159 / v1) + 3.14159) + (v1 + 0.1))
                                                _ = v31
                                            } else {
                                                if true {
                                                    if (true && ((!(42 <= 33)) && false)) {
                                                        fmt.Println(fi25, v30)
                                                        v32 := GenS3{F0: -730, F1: 1, F2: (2147483648 == 4294967295)}
                                                        fmt.Println(v30, v23)
                                                        _ = v32
                                                    } else {
                                                        fmt.Println(v23)
                                                    }
                                                } else {
                                                    var v33 uint8 = 127
                                                    _ = v33
                                                }
                                            }
                                            v34 := GenS4{F0: v23, F1: "a", F2: 0}
                                            var v35 int = -938
                                            _ = v34
                                            _ = v35
                                        } else {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            var v36 float64 = 3.14159
                                            if (v30 == v30) {
                                                v30 = (v30 | (v30 & 825))
                                                v1 += (0.0 - (3.14159 / 1.0))
                                                v36 += 1e10
                                                v36 += float64((float32(v36) + ((1e10 + (v1 / v1)) - (((v1 + (((float32((v36 / v36)) + v1) * float32((v36 / (((v36 * v36) / 1.0) + v36)))) / -2.25)) / ((float32(-2.25) - v1) - v1)) / (v1 / (v1 * v1))))))
                                                if len(v2) > 0 {
                                                    v2[0] = ""
                                                }
                                            }
                                            _ = v36
                                        }
                                        switch (-(fi22)) {
                                        case -9223372036854775808:
                                            v23 += -128
                                        }
                                        v37 := GenS4{F0: (v23 / ((v23) | 1)), F1: "foo bar", F2: 42}
                                        v1 += 1.5
                                        _ = v37
                                    } else {
                                        if true {
                                            if (v30 > (v30 >> ((58) & 63))) {
                                                v38 := GenS5{F0: ("foo bar" + "foo bar"), F1: 1, F2: true}
                                                var v39 uint = 255
                                                _ = v38
                                                _ = v39
                                            } else {
                                                if false {
                                                    v23 += (v23 - 1)
                                                }
                                            }
                                            v1 += v1
                                            var v40 int64 = -699
                                            _ = v40
                                        } else {
                                            var v41 int64 = 834
                                            v42 := GenS2{F0: 246}
                                            _ = v41
                                            _ = v42
                                        }
                                        if (0 <= 65535) {
                                            var v43 int16 = 1
                                            _ = v43
                                        }
                                    }
                                    _ = v30
                                }
                                _ = len(v2)
                                v1 -= v1
                            } else {
                                fmt.Println(fi22, v20, v1)
                                v23 -= (^(v23))
                            }
                            switch fi22 {
                            case 9223372036854775807:
                                if len(v3) > 0 {
                                    v3[0] = 1
                                }
                            case 42:
                                var v44 uint8 = 127
                                _ = v44
                            case -1:
                                v23 = v23
                            default:
                                v45 := GenS4{F0: v23, F1: "foo bar", F2: -1}
                                _ = v45
                            }
                            fmt.Println(v1, v20, fi22)
                        default:
                            v1 -= v1
                        }
                    } else {
                        var v46 int8 = 42
                        switch (fi22 - fi25) {
                        case -9223372036854775808:
                            v23 = (v23 / (((v23 << ((1) & 63))) | 1))
                        default:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        v47 := GenS4{F0: 20, F1: "a", F2: -470}
                        _ = v46
                        _ = v47
                    }
                    v1 += 0.0
                    _ = fi25
                }
                v3 = append(v3, 679)
            }
            _ = fi22
            _ = v23
            _ = v24
        }
    case 1:
        switch 9223372036854775807 {
        case 4294967295:
            for fi48 := 0; fi48 < 20; fi48++ {
                for fi49 := 0; fi49 < 8; fi49++ {
                    v1 += 0.1
                    if true {
                        if (4294967295 != 42) {
                            var v50 int64 = -904
                            _ = v50
                        } else {
                            v20 = v20
                            switch (fi48 + (fi48 & fi49)) {
                            case 9223372036854775807:
                                v51 := GenS2{F0: 127}
                                _ = v51
                            case -2147483648:
                                v20 = -945
                            case 2147483647:
                                if (!false) {
                                    v20 = v20
                                    if len(v2) > 0 {
                                        v2[0] = (("\n\t" + (("\n\t" + ("" + "a")) + "foo bar")) + "a")
                                    }
                                    v2 = append(v2, "")
                                }
                            }
                            v20 += (v20 + v20)
                            v52 := GenS6{F0: (255 < 2), F1: 18446744073709551615, F2: ((0.0 + 1e10) != v1)}
                            var v53 uint = 1
                            _ = v52
                            _ = v53
                        }
                    }
                    _ = fi49
                }
                v20 = (v20 % ((v20) | 1))
                _ = fi48
            }
        default:
            switch "hello" {
            case "foo bar":
                switch -485 {
                case -370:
                    if (true || false) {
                        for fi54 := 0; fi54 < 19; fi54++ {
                            if true {
                                var v55 int8 = -101
                                var v56 int16 = 0
                                _ = v55
                                _ = v56
                            } else {
                                if (!(40 >= 42)) {
                                    v57 := GenS1{F0: fi54, F1: 994}
                                    v1 = (1.5 * v1)
                                    _ = v57
                                }
                                for fi58 := 0; fi58 < 7; fi58++ {
                                    if (false && (2 != 494)) {
                                        var v59 uint64 = 2
                                        v60 := GenS5{F0: ("a" + "hello"), F1: (v20 >> ((1) & 63)), F2: false}
                                        _ = v59
                                        _ = v60
                                    }
                                    var v61 int = (fi58 ^ -1)
                                    var v62 int8 = 127
                                    _ = fi58
                                    _ = v61
                                    _ = v62
                                }
                            }
                            v63 := GenS7{F0: 18446744073709551615, F1: false}
                            if (0 == 50) {
                                v64 := GenS6{F0: false, F1: 42, F2: ("a" == "a")}
                                _ = v64
                            } else {
                                var v65 uint32 = 2147483647
                                v65 -= 65535
                                _ = v65
                            }
                            _ = fi54
                            _ = v63
                        }
                        v2 = append(v2, ("foo bar" + "\n\t"))
                    } else {
                        if len(v3) > 0 {
                            v3[0] = 4294967295
                        }
                        v66 := GenS5{F0: "foo bar", F1: 116, F2: (false && (!(!((false && true) || (29 != -128)))))}
                        _ = v66
                    }
                default:
                    var v67 int8 = -128
                    _ = v67
                }
            case "a":
                if (-46 > 42) {
                    var v68 uint16 = 0
                    _ = v68
                } else {
                    v1 -= ((float32((1.5 * -2.25)) * v1) * v1)
                    fmt.Println(v1, v1)
                    for fi69 := 0; fi69 < 10; fi69++ {
                        if len(v3) > 0 {
                            v3[0] = 98
                        }
                        var v70 int = (fi69 - (^(fi69)))
                        _ = fi69
                        _ = v70
                    }
                }
            default:
                var v71 int32 = v20
                _ = v71
            }
        }
    }
    v72 := GenS8{F0: 634, F1: 0}
    fmt.Println(v20)
    switch 127 {
    case 2147483647:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    switch "foo bar" {
    case "foo bar":
        switch ("hello" + ("\n\t" + ("" + "\n\t"))) {
        case "foo bar":
            for fi73 := 0; fi73 < 1; fi73++ {
                switch "foo bar" {
                case "hello":
                    if len(v2) > 0 {
                        v2[0] = ""
                    }
                }
                _ = len(v3)
                _ = fi73
            }
        case "":
            for fi74 := 0; fi74 < 19; fi74++ {
                v75 := GenS9{F0: 42, F1: (-1 < fi74)}
                var v76 uint8 = 42
                if (322 <= 18446744073709551615) {
                    fmt.Println(v76)
                    v76 = v76
                } else {
                    v77 := GenS10{F0: v76}
                    v76 = 86
                    _ = v77
                }
                _ = fi74
                _ = v75
                _ = v76
            }
        }
    case "a":
        var v78 int8 = 42
        _ = v78
    default:
        v20 = v20
    }
    fmt.Println(v20, v1)
    if (false || false) {
        if (89 != int(v20)) {
            v1 += 1.5
        } else {
            fmt.Println(v20)
            for fi79 := 0; fi79 < 9; fi79++ {
                v80 := GenS5{F0: "hello", F1: (v20 | v20), F2: true}
                var v81 int16 = int16(fi79)
                var v82 int = fi79
                _ = fi79
                _ = v80
                _ = v81
                _ = v82
            }
            v2 = append(v2, "\n\t")
        }
        var v83 int = 300
        _ = v83
    }
    var v84 uint64 = 4294967295
    for fi85 := 0; fi85 < 17; fi85++ {
        v86 := GenS3{F0: v20, F1: 290, F2: (true && true)}
        var v87 int64 = 944
        for fi88 := 0; fi88 < 13; fi88++ {
            if (-1 <= -858) {
                var v89 bool = (true || (46 >= 1))
                v84 = v84
                v90 := GenS11{F0: 42, F1: v20, F2: (v20 - v20)}
                _ = v89
                _ = v90
            }
            var v91 int16 = -1
            if (-9223372036854775808 > -1) {
                switch fi88 {
                case 42:
                    v87 -= v87
                case -851:
                    var v92 uint8 = 128
                    _ = v92
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v93 := GenS9{F0: (v87 / ((int64(fi85)) | 1)), F1: true}
                switch "foo bar" {
                case "foo bar":
                    v94 := GenS1{F0: -1, F1: 65535}
                    _ = v94
                default:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                _ = v93
            }
            _ = fi88
            _ = v91
        }
        _ = fi85
        _ = v86
        _ = v87
    }
    if true {
        var v95 bool = (false || (true || (127 != 1)))
        v95 = (42 <= 2147483647)
        _ = v95
    } else {
        var v96 uint = 65535
        var v97 int64 = 1
        _ = v96
        _ = v97
    }
    switch "" {
    case "\n\t":
        fmt.Println(v20, v84, v1)
    case "":
        var v98 string = "foo bar"
        _ = v98
    }
    if (-320 == (-(v20))) {
        v99 := GenS12{F0: 42}
        v100 := GenS2{F0: 127}
        _ = v100
        _ = v99
    } else {
        if false {
            switch "hello" {
            case "\n\t":
                v84 = v84
            case "":
                var v101 string = ((("" + "foo bar") + ("" + "a")) + ("\n\t" + ("foo bar" + "hello")))
                _ = v101
            case "foo bar":
                v84 -= v84
            default:
                fmt.Println(v84, v20)
            }
            if len(v2) > 0 {
                v2[0] = (("a" + "a") + ("hello" + "a"))
            }
        }
        for fi102 := 0; fi102 < 8; fi102++ {
            v84 = (^(v84))
            if true {
                v103 := GenS13{F0: v84, F1: 0}
                _ = len(v2)
                switch (("" + "hello") + "hello") {
                case "foo bar":
                    if (65535 == uint(v84)) {
                        v2 = append(v2, "foo bar")
                    } else {
                        v3 = append(v3, 2)
                        v84 += v84
                    }
                }
                _ = v103
            }
            _ = fi102
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v72
    _ = v84
}
