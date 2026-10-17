package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 uint32; F2 uint32 }
type GenS3 struct { F0 uint64; F1 uint32; F2 uint32 }
type GenS4 struct { F0 string; F1 int64; F2 uint64 }
type GenS5 struct { F0 int64; F1 int16; F2 int8 }
type GenS6 struct { F0 bool; F1 int32; F2 uint16 }
type GenS7 struct { F0 int64; F1 uint64; F2 int }
type GenS8 struct { F0 int32; F1 bool; F2 int8 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int; F1 int32 }
type GenS11 struct { F0 string; F1 string }

func main() {
    var v1 float32 = -2.25
    var v2 string = "hello"
    var v3 uint16 = 32767
    v4 := make(map[string]int)
    v5 := make(map[int]int)
    fmt.Println(v2, v2, v1)
    if false {
        v6 := GenS1{F0: 42}
        _ = v6
    } else {
        var v7 int16 = 901
        v8 := GenS1{F0: int64(v7)}
        _ = v7
        _ = v8
    }
    if (42 != -32768) {
        v9 := GenS1{F0: -1}
        if false {
            var v10 int32 = -1
            for fi11 := 0; fi11 < 6; fi11++ {
                v1 -= 0.1
                if (!(!((fi11 | fi11) >= (fi11 % ((fi11) | 1))))) {
                    v12 := GenS2{F0: -128, F1: 2147483647, F2: 774}
                    _ = v12
                } else {
                    var v13 float32 = (v1 - v1)
                    for fi14 := 0; fi14 < 7; fi14++ {
                        _ = len(v5)
                        v3 += 65535
                        _ = fi14
                    }
                    _ = v13
                }
                _ = fi11
            }
            v10 = 1
            _ = v10
        }
        _ = v9
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        v3 -= (v3 / ((v3) | 1))
        v2 = "\n\t"
        _ = fi15
    }
    var v16 uint = 1
    for fi17 := 0; fi17 < 4; fi17++ {
        switch (((("a" + "foo bar") + v2) + v2) + v2) {
        case "\n\t":
            switch "foo bar" {
            case "hello":
                for fi18 := 0; fi18 < 12; fi18++ {
                    var v19 uint8 = 42
                    v5[fi17] = fi17
                    v19 += uint8(v3)
                    v20 := []int{(^(fi17)), (fi18 | 128), fi18}
                    if (-96 == 0) {
                        var v21 string = "hello"
                        var v22 uint = 667
                        _ = v21
                        _ = v22
                    } else {
                        v16 -= (v16 % (((v16 - 18446744073709551615)) | 1))
                        if (836 > 1) {
                            if true {
                                if (v3 <= v3) {
                                    v3 -= 42
                                    var v23 uint = uint(v3)
                                    if len(v20) >= 2 {
                                        v20 = v20[1:len(v20)]
                                    }
                                    _ = v23
                                }
                                if true {
                                    fmt.Println(v3, v16, v2)
                                    v24 := GenS1{F0: 824}
                                    _ = v24
                                }
                                v25 := GenS3{F0: 42, F1: 2, F2: 744}
                                _ = v25
                            }
                        } else {
                            v16 = uint(v19)
                        }
                        var v26 uint32 = 56
                        if ("a" > ((v2 + ((v2 + v2) + ("" + (v2 + v2)))) + v2)) {
                            v27 := GenS4{F0: (v2 + v2), F1: 0, F2: 1}
                            fmt.Println(v2, v3, v3)
                            _ = v27
                        } else {
                            var v28 uint = (v16 + (^(v16)))
                            _ = v28
                        }
                        _ = v26
                    }
                    _ = fi18
                    _ = v19
                    _ = v20
                }
            case "\n\t":
                var v29 int16 = 1000
                _ = v29
            case "a":
                var v30 uint16 = v3
                _ = v30
            default:
                var v31 float32 = 1.5
                _ = v31
            }
        default:
            switch v2 {
            case "a":
                v3 = 337
            case "foo bar":
                fmt.Println(fi17)
            case "\n\t":
                v2 = ""
            default:
                switch fi17 {
                case 42:
                    v1 = (float32(float64(1.5)) + v1)
                case 1:
                    v32 := GenS1{F0: -9223372036854775808}
                    _ = v32
                case 9223372036854775807:
                    fmt.Println(v1, fi17, v2)
                default:
                    switch "" {
                    case "foo bar":
                        v3 = (v3 * v3)
                    case "\n\t":
                        fmt.Println(v3, v3)
                    }
                }
            }
        }
        _ = fi17
    }
    var v33 bool = true
    for fi34 := 0; fi34 < 14; fi34++ {
        switch fi34 {
        case 9223372036854775807:
            fmt.Println(fi34, v1, v16)
        case 32767:
            if v33 {
                var v35 uint = (v16 / (((v16 + (v16 / ((v16) | 1)))) | 1))
                var v36 uint16 = (v3 % (((v3 - (v3 << ((v35) & 63)))) | 1))
                _ = v35
                _ = v36
            } else {
                for fi37 := 0; fi37 < 5; fi37++ {
                    if v33 {
                        var v38 int8 = 42
                        var v39 float64 = (1.5 + 3.14159)
                        _ = v38
                        _ = v39
                    } else {
                        v16 += 255
                        v40 := GenS3{F0: 2, F1: 2147483647, F2: 554}
                        v1 -= (0.1 + 1e10)
                        _ = v40
                    }
                    v1 = (v1 / ((3.14159 + float32((float64(v1) + (0.0 / 3.14159)))) * v1))
                    var v41 int8 = 0
                    if false {
                        v42 := GenS4{F0: "hello", F1: 42, F2: 976}
                        v2 = "foo bar"
                        _ = len(v5)
                        if true {
                            switch fi34 {
                            case -2147483648:
                                switch fi37 {
                                case -64:
                                    var v43 uint32 = 299
                                    _ = v43
                                }
                            case -280:
                                v16 -= v16
                            case -6:
                                var v44 uint16 = 32767
                                _ = v44
                            default:
                                v33 = (float32(-2.25) >= v1)
                            }
                        } else {
                            v45 := GenS5{F0: 42, F1: -32768, F2: (v41 / ((v41) | 1))}
                            v46 := GenS6{F0: v33, F1: 37, F2: (^(v3))}
                            var v47 uint16 = v3
                            _ = v45
                            _ = v46
                            _ = v47
                        }
                        _ = v42
                    }
                    var v48 bool = (!((!true) || (1 >= 1)))
                    _ = fi37
                    _ = v41
                    _ = v48
                }
            }
        case 2147483647:
            for fi49 := 0; fi49 < 19; fi49++ {
                v1 = v1
                var v50 float64 = 0.1
                _ = fi49
                _ = v50
            }
        default:
            switch fi34 {
            case 127:
                fmt.Println(v16, v16)
            }
        }
        _ = fi34
    }
    for fi51 := 0; fi51 < 2; fi51++ {
        v52 := v4["a"]
        for fi53 := 0; fi53 < 6; fi53++ {
            var v54 int32 = -2147483648
            _ = fi53
            _ = v54
        }
        v55 := v5[42]
        _ = fi51
        _ = v52
        _ = v55
    }
    switch "hello" {
    case "foo bar":
        v56 := []int{}
        _ = v56
    case "hello":
        for fi57 := 0; fi57 < 17; fi57++ {
            var v58 uint = uint(v3)
            v16 -= (v16 >> ((v16) & 63))
            _ = fi57
            _ = v58
        }
    }
    var v59 int16 = -32768
    if (2147483648 >= 1) {
        if true {
            v2 = (("\n\t" + "hello") + ((v2 + v2) + (("hello" + "") + "foo bar")))
        }
        v3 += v3
    }
    v60 := []int{-1}
    v16 += v16
    v61 := GenS7{F0: -454, F1: 4294967295, F2: -554}
    if len(v60) >= 2 {
        v60 = v60[1:len(v60)]
    }
    if true {
        if (v33 && (uint16(v16) != 299)) {
            for fi62 := 0; fi62 < 5; fi62++ {
                for fi63 := 0; fi63 < 13; fi63++ {
                    if (!(-64 != (fi63 + (^(fi63))))) {
                        _ = len(v4)
                        var v64 uint32 = 2
                        _ = v64
                    } else {
                        fmt.Println(v3, fi62, v1)
                        if (0 <= 127) {
                            fmt.Println(fi63, fi62, v1)
                            var v65 float64 = ((0.0 / (3.14159 - -2.25)) * 1e10)
                            v16 = v16
                            _ = v65
                        }
                        fmt.Println(v33)
                        if v33 {
                            if v33 {
                                switch fi62 {
                                case 853:
                                    var v66 uint16 = v3
                                    _ = v66
                                }
                                switch ("hello" + "foo bar") {
                                case "hello":
                                    fmt.Println(v59, v16)
                                    var v67 float32 = ((v1 / 1.0) / 1.5)
                                    _ = v67
                                default:
                                    v68 := GenS8{F0: -2147483648, F1: (!(!(v33 || (1 < 255)))), F2: -1}
                                    _ = v68
                                }
                                v2 = ""
                            }
                            var v69 int8 = 26
                            _ = v69
                        } else {
                            switch (v2 + v2) {
                            case "\n\t":
                                var v70 uint64 = 756
                                _ = v70
                            case "":
                                v16 += v16
                            }
                            switch (fi62 | (fi62 % (((fi62 * 996)) | 1))) {
                            case -2147483648:
                                v33 = ((-95 != 42) || ((v33 || true) || v33))
                            case -702:
                                switch (fi62 + fi62) {
                                case 0:
                                    var v71 float64 = (0.1 / 1.0)
                                    _ = v71
                                case -970:
                                    if v33 {
                                        v2 = v2
                                    }
                                case -285:
                                    _ = len(v4)
                                default:
                                    switch -1 {
                                    case 4294967295:
                                        v2 = ("a" + v2)
                                    case 42:
                                        delete(v4, v2)
                                    case 32767:
                                        if len(v60) > 0 {
                                            v60[0] = fi62
                                        }
                                    default:
                                        v33 = false
                                    }
                                }
                            default:
                                v3 = v3
                            }
                            if v33 {
                                if (-409 < 1) {
                                    v2 = ("" + v2)
                                    var v72 uint16 = v3
                                    _ = v72
                                }
                            } else {
                                fmt.Println(v16, v1)
                                v16 = (v16 * (v16 << ((255) & 63)))
                            }
                        }
                        v2 = (("a" + "") + "foo bar")
                    }
                    _ = fi63
                }
                switch "hello" {
                case "":
                    if len(v60) >= 2 {
                        v60 = v60[1:len(v60)]
                    }
                }
                _ = fi62
            }
        }
    } else {
        if false {
            if v33 {
                fmt.Println(v33, v2, v59)
                v2 = "hello"
            } else {
                switch "a" {
                case "\n\t":
                    v73 := v5[42]
                    _ = v73
                case "hello":
                    if true {
                        var v74 uint = v16
                        _ = v74
                    }
                case "":
                    v75 := GenS9{F0: 0}
                    _ = v75
                }
                if len(v60) > 0 {
                    v60[0] = -649
                }
                if true {
                    v76 := v4["foo bar"]
                    _ = v76
                } else {
                    _ = len(v4)
                    var v77 int = 481
                    var v78 uint16 = v3
                    _ = v77
                    _ = v78
                }
            }
            switch v2 {
            case "hello":
                v60 = append(v60, 42)
            case "\n\t":
                if true {
                    if true {
                        fmt.Println(v33, v3)
                        var v79 uint16 = (v3 + v3)
                        _ = v79
                    }
                    for fi80 := 0; fi80 < 9; fi80++ {
                        var v81 int = 32767
                        fmt.Println(v59, v81, v59)
                        _ = fi80
                        _ = v81
                    }
                } else {
                    v82 := GenS8{F0: 129, F1: v33, F2: 42}
                    var v83 float64 = 3.14159
                    _ = v82
                    _ = v83
                }
            default:
                if v33 {
                    v59 -= 283
                    if v33 {
                        switch ("" + v2) {
                        case "\n\t":
                            if len(v60) >= 2 {
                                v60 = v60[1:len(v60)]
                            }
                        default:
                            v84 := GenS3{F0: 1, F1: 65535, F2: 0}
                            _ = v84
                        }
                        if v33 {
                            if ((-(v59)) != 42) {
                                switch 9223372036854775807 {
                                case 42:
                                    if (("hello" != "") || v33) {
                                        v16 -= (v16 | (v16 & v16))
                                    }
                                case 128:
                                    v1 = 1.5
                                default:
                                    var v85 int64 = 352
                                    _ = v85
                                }
                                var v86 float64 = 0.1
                                _ = v86
                            } else {
                                switch -456 {
                                case 764:
                                    var v87 uint8 = 0
                                    _ = v87
                                case -825:
                                    switch 683 {
                                    case 32768:
                                        if (69 <= 967) {
                                            v59 -= (v59 % ((v59) | 1))
                                            var v88 int64 = 42
                                            if ((^(v16)) < v16) {
                                                v59 -= (v59 / ((223) | 1))
                                                var v89 uint = v16
                                                _ = v89
                                            }
                                            _ = v88
                                        } else {
                                            var v90 float32 = v1
                                            _ = v90
                                        }
                                    }
                                case 2147483647:
                                    for fi91 := 0; fi91 < 10; fi91++ {
                                        if (v3 > 2) {
                                            var v92 uint64 = uint64(v16)
                                            var v93 string = ""
                                            switch "\n\t" {
                                            case "foo bar":
                                                if false {
                                                    if true {
                                                        for fi94 := 0; fi94 < 20; fi94++ {
                                                            v16 -= v16
                                                            _ = fi94
                                                        }
                                                        for fi95 := 0; fi95 < 8; fi95++ {
                                                            fmt.Println(v2, fi91)
                                                            v3 -= 624
                                                            if v33 {
                                                                v16 += 704
                                                                switch v93 {
                                                                case "\n\t":
                                                                    _ = len(v60)
                                                                    v59 += v59
                                                                    v96 := GenS10{F0: fi95, F1: -1}
                                                                    _ = v96
                                                                case "foo bar":
                                                                    var v97 int32 = 1
                                                                    _ = v97
                                                                }
                                                                var v98 uint64 = (v92 << ((18446744073709551615) & 63))
                                                                if v33 {
                                                                    v98 += 260
                                                                } else {
                                                                    var v99 int8 = 42
                                                                    v16 = v16
                                                                    switch (fi91 - fi95) {
                                                                    case -1:
                                                                        var v100 uint64 = (v98 / (((v98 * 9)) | 1))
                                                                        _ = v100
                                                                    case 0:
                                                                        var v101 uint16 = (v3 ^ (v3 ^ v3))
                                                                        _ = v101
                                                                    default:
                                                                        v33 = (false && false)
                                                                    }
                                                                    _ = v99
                                                                }
                                                                _ = v98
                                                            } else {
                                                                switch "a" {
                                                                case "hello":
                                                                    if (!v33) {
                                                                        var v102 uint8 = 163
                                                                        fmt.Println(v3)
                                                                        _ = v102
                                                                    } else {
                                                                        v59 = (v59 << (((v16 & (v16 << ((v16) & 63)))) & 63))
                                                                        var v103 float32 = (1.5 / v1)
                                                                        _ = v103
                                                                    }
                                                                case "a":
                                                                    v3 = uint16(v16)
                                                                case "\n\t":
                                                                    if len(v60) > 0 {
                                                                        v60[0] = 128
                                                                    }
                                                                }
                                                                if v33 {
                                                                    var v104 uint32 = 2
                                                                    if v33 {
                                                                        v60 = append(v60, (-(fi95)))
                                                                    }
                                                                    var v105 uint8 = 42
                                                                    _ = v104
                                                                    _ = v105
                                                                } else {
                                                                    if true {
                                                                        if v33 {
                                                                            switch 0 {
                                                                            case 2147483647:
                                                                                v106 := GenS11{F0: "hello", F1: "a"}
                                                                                _ = v106
                                                                            case 556:
                                                                                var v107 int8 = -15
                                                                                _ = v107
                                                                            default:
                                                                                fmt.Println(fi95, v16, v93)
                                                                            }
                                                                            fmt.Println(v16)
                                                                        }
                                                                        v108 := v4[v93]
                                                                        _ = v108
                                                                    } else {
                                                                        var v109 uint8 = 48
                                                                        _ = v109
                                                                    }
                                                                    switch v93 {
                                                                    case "":
                                                                        fmt.Println(v92)
                                                                    case "hello":
                                                                        var v110 int64 = 126
                                                                        _ = v110
                                                                    }
                                                                    var v111 uint64 = v92
                                                                    _ = v111
                                                                }
                                                            }
                                                            _ = fi95
                                                        }
                                                        switch (fi91 >> ((v16) & 63)) {
                                                        case 9223372036854775807:
                                                            var v112 float32 = (v1 - v1)
                                                            _ = v112
                                                        case 87:
                                                            v93 = ("\n\t" + v2)
                                                        default:
                                                            for fi113 := 0; fi113 < 7; fi113++ {
                                                                v4[v2] = (fi113 >> ((uint(v92)) & 63))
                                                                _ = fi113
                                                            }
                                                        }
                                                    } else {
                                                        delete(v5, fi91)
                                                    }
                                                }
                                            case "hello":
                                                switch v93 {
                                                case "hello":
                                                    var v114 bool = (127 >= 100)
                                                    _ = v114
                                                case "foo bar":
                                                    var v115 int8 = int8(v16)
                                                    _ = v115
                                                default:
                                                    switch v93 {
                                                    case "foo bar":
                                                        v16 = (^(v16))
                                                    case "":
                                                        v33 = (((v93 + (v2 + v93)) >= v2) && (v16 != (v16 >> (((v16 << ((710) & 63))) & 63))))
                                                    default:
                                                        fmt.Println(v33)
                                                    }
                                                }
                                            case "":
                                                v4[v2] = 4294967295
                                            default:
                                                var v116 int16 = 32767
                                                _ = v116
                                            }
                                            _ = v92
                                            _ = v93
                                        } else {
                                            if (true && ((v33 || (fi91 > (fi91 >> ((604) & 63)))) || (2147483648 < 65535))) {
                                                v117 := GenS7{F0: 9223372036854775807, F1: 623, F2: (fi91 - fi91)}
                                                _ = v117
                                            }
                                            for fi118 := 0; fi118 < 9; fi118++ {
                                                v2 = ((("foo bar" + "hello") + "a") + (v2 + v2))
                                                if len(v60) >= 2 {
                                                    v60 = v60[1:len(v60)]
                                                }
                                                v1 = 0.0
                                                _ = fi118
                                            }
                                        }
                                        _ = fi91
                                    }
                                }
                                v119 := GenS9{F0: 32767}
                                _ = v119
                            }
                            delete(v5, -1)
                            if (42 != -829) {
                                var v120 int16 = v59
                                _ = v120
                            }
                        } else {
                            var v121 uint8 = 42
                            _ = v121
                        }
                    }
                }
            }
            var v122 uint32 = 65535
            _ = v122
        }
    }
    for fi123 := 0; fi123 < 2; fi123++ {
        v124 := GenS6{F0: (9223372036854775807 >= 1), F1: 2147483647, F2: (v3 & (v3 * v3))}
        v125 := v5[fi123]
        fmt.Println(v2, v125, v59)
        _ = fi123
        _ = v124
        _ = v125
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v59:", v59)
    fmt.Println("slicelen v60:", len(v60), v60)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v59
    _ = v60
    _ = v61
}
