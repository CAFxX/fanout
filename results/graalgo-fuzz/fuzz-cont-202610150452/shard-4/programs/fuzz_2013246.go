package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 int32; F2 int64 }
type GenS4 struct { F0 bool; F1 uint64; F2 int8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint8; F1 int16; F2 uint32 }
type GenS7 struct { F0 int64 }

func main() {
    v1 := make(map[string]uint32)
    v2 := []uint32{2147483647}
    v3 := []int{}
    var v4 float32 = (0.0 / 1.5)
    v5 := make(map[uint32]string)
    var v6 uint16 = 42
    var v7 float64 = 3.14159
    v7 += (v7 - v7)
    switch (((("" + ("a" + "")) + (((("\n\t" + (("foo bar" + "hello") + "\n\t")) + ((("" + "") + "\n\t") + "a")) + "a") + "foo bar")) + "hello") + "a") {
    case "\n\t":
        v6 = 678
    case "hello":
        var v8 string = "\n\t"
        _ = v8
    case "":
        v4 += 0.0
    default:
        v6 += (v6 / ((189) | 1))
    }
    v7 = v7
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 int16 = 1
        v7 += v7
        v10 -= (-(v10))
        _ = fi9
        _ = v10
    }
    var v11 int8 = 60
    for fi12 := 0; fi12 < 9; fi12++ {
        v7 = (1.5 / v7)
        switch "\n\t" {
        case "hello":
            v13 := GenS1{F0: 2147483647}
            _ = v13
        case "":
            if true {
                v4 += v4
                var v14 string = ("\n\t" + (((("foo bar" + "hello") + "hello") + "\n\t") + (("hello" + ("\n\t" + (("hello" + "hello") + "foo bar"))) + "\n\t")))
                for fi15 := 0; fi15 < 12; fi15++ {
                    if true {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if false {
                            if false {
                                v6 = 32767
                                v7 += 0.1
                                if false {
                                    v11 += 77
                                }
                            } else {
                                if (true && false) {
                                    v14 = v14
                                    v16 := GenS1{F0: 4294967295}
                                    var v17 int64 = -85
                                    _ = v16
                                    _ = v17
                                } else {
                                    var v18 uint8 = 255
                                    v19 := GenS2{F0: true}
                                    _ = v18
                                    _ = v19
                                }
                                v20 := GenS2{F0: false}
                                _ = v20
                            }
                            v4 = v4
                        }
                    } else {
                        if false {
                            if (false && ((113 <= -1) || false)) {
                                switch v14 {
                                case "a":
                                    v4 -= (v4 - v4)
                                case "hello":
                                    v21 := GenS3{F0: 42, F1: -1, F2: -1}
                                    _ = v21
                                case "\n\t":
                                    fmt.Println(v4, fi15)
                                default:
                                    fmt.Println(v4)
                                }
                                var v22 float32 = (1e10 * v4)
                                var v23 uint = 200
                                _ = v22
                                _ = v23
                            }
                            if true {
                                var v24 float32 = v4
                                var v25 uint32 = 4294967295
                                v26 := GenS1{F0: v25}
                                v27 := GenS1{F0: (v25 ^ (v25 << ((0) & 63)))}
                                var v28 bool = true
                                _ = v24
                                _ = v25
                                _ = v26
                                _ = v27
                                _ = v28
                            } else {
                                if false {
                                    if true {
                                        v29 := GenS1{F0: 42}
                                        var v30 bool = (!true)
                                        _ = v29
                                        _ = v30
                                    }
                                    fmt.Println(v11)
                                    fmt.Println(v4, v11)
                                }
                            }
                        }
                        fmt.Println(v14, v7)
                    }
                    _ = fi15
                }
                _ = v14
            }
        }
        _ = len(v5)
        _ = fi12
    }
    if (true || true) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        var v31 uint64 = 0
        for fi32 := 0; fi32 < 20; fi32++ {
            var v33 uint64 = 42
            for fi34 := 0; fi34 < 9; fi34++ {
                var v35 uint64 = v31
                if false {
                    switch (-(fi34)) {
                    case 0:
                        if (!true) {
                            var v36 uint = 917
                            v37 := GenS4{F0: (true && (227 <= 2)), F1: 340, F2: v11}
                            switch fi32 {
                            case 595:
                                v31 = (^(v31))
                            case 32:
                                fmt.Println(fi34, v31, v36)
                            case 412:
                                var v38 uint8 = 42
                                _ = v38
                            }
                            _ = v36
                            _ = v37
                        } else {
                            v39 := GenS1{F0: 2147483647}
                            if true {
                                var v40 int8 = v11
                                if (((false && (!(true && false))) && (9223372036854775807 >= -9223372036854775808)) && true) {
                                    v11 += v11
                                    fmt.Println(v33)
                                    if (true || false) {
                                        fmt.Println(v40, fi34, v4)
                                        if (false || (69 != 2)) {
                                            fmt.Println(v7, v6, fi32)
                                        } else {
                                            if true {
                                                v41 := GenS3{F0: 139, F1: 2147483647, F2: 1}
                                                _ = v41
                                            }
                                            if ((fi32 % ((fi34) | 1)) != (fi34 | (-(fi32)))) {
                                                v42 := GenS5{F0: -1}
                                                _ = v42
                                            }
                                        }
                                        var v43 uint32 = 0
                                        _ = v43
                                    } else {
                                        fmt.Println(v33, v33)
                                        v40 -= (v11 ^ (v40 % ((v11) | 1)))
                                        switch "" {
                                        case "\n\t":
                                            v6 += (v6 % ((1) | 1))
                                        default:
                                            var v44 bool = (v11 <= 127)
                                            _ = v44
                                        }
                                    }
                                }
                                fmt.Println(v11, v35, v33)
                                var v45 int64 = 9223372036854775807
                                _ = v40
                                _ = v45
                            } else {
                                _ = len(v1)
                                if true {
                                    switch (fi32 & fi34) {
                                    case 32767:
                                        var v46 int8 = -128
                                        _ = v46
                                    default:
                                        fmt.Println(fi32, v33)
                                    }
                                    fmt.Println(v4, fi32, v11)
                                    if true {
                                        fmt.Println(fi32, v35, fi32)
                                    } else {
                                        var v47 uint16 = (v6 % ((42) | 1))
                                        fmt.Println(fi32, v6, v6)
                                        fmt.Println(v33, v33)
                                        _ = v47
                                    }
                                    var v48 float64 = float64(-2.25)
                                    _ = v48
                                }
                                v7 += v7
                            }
                            if (false && (570 == 42)) {
                                v33 = v31
                                switch (fi34 >> ((0) & 63)) {
                                case 2147483647:
                                    if true {
                                        if true {
                                            if false {
                                                v1["a"] = 42
                                            }
                                        } else {
                                            switch (^(fi34)) {
                                            case -823:
                                                var v49 int = 32767
                                                _ = v49
                                            }
                                            fmt.Println(v4)
                                        }
                                        v50 := GenS4{F0: false, F1: uint64(v6), F2: v11}
                                        if true {
                                            if (!((!(!false)) || (!(!(!((^(fi34)) >= fi34)))))) {
                                                var v51 uint8 = 1
                                                fmt.Println(v33, v4, fi32)
                                                if (0 >= 9223372036854775807) {
                                                    v11 -= 127
                                                    switch (fi32 >> ((104) & 63)) {
                                                    case -1:
                                                        _ = len(v5)
                                                    }
                                                } else {
                                                    v52 := GenS2{F0: (-1 >= -661)}
                                                    var v53 uint32 = 65535
                                                    _ = v52
                                                    _ = v53
                                                }
                                                _ = v51
                                            }
                                            fmt.Println(v7, v6)
                                        } else {
                                            var v54 uint = 247
                                            _ = v54
                                        }
                                        _ = v50
                                    } else {
                                        var v55 string = "hello"
                                        if false {
                                            if (v4 >= v4) {
                                                var v56 int64 = 1
                                                _ = v56
                                            } else {
                                                if true {
                                                    v5[226] = v55
                                                    if true {
                                                        if (v55 == v55) {
                                                            v33 = v35
                                                            fmt.Println(v31, v4, v33)
                                                        } else {
                                                            v4 = 0.1
                                                            var v57 uint32 = 1
                                                            _ = v57
                                                        }
                                                        if false {
                                                            var v58 uint = 1
                                                            switch ("a" + v55) {
                                                            case "hello":
                                                                if (!((^(fi34)) <= fi32)) {
                                                                    fmt.Println(v6)
                                                                    v11 -= (v11 * v11)
                                                                } else {
                                                                    var v59 float64 = float64(v4)
                                                                    var v60 bool = true
                                                                    _ = v59
                                                                    _ = v60
                                                                }
                                                            case "foo bar":
                                                                v2 = append(v2, 2147483647)
                                                            case "\n\t":
                                                                var v61 int64 = 9223372036854775807
                                                                _ = v61
                                                            }
                                                            fmt.Println(v55)
                                                            _ = v58
                                                        } else {
                                                            var v62 float64 = 1.5
                                                            v63 := GenS3{F0: 193, F1: -1, F2: 42}
                                                            v4 -= 0.1
                                                            _ = v62
                                                            _ = v63
                                                        }
                                                        var v64 int8 = (v11 / (((v11 - 1)) | 1))
                                                        _ = v64
                                                    }
                                                    var v65 float32 = v4
                                                    _ = v65
                                                }
                                                v2 = append(v2, 1)
                                                var v66 int8 = (^(v11))
                                                v67 := GenS1{F0: 128}
                                                _ = v66
                                                _ = v67
                                            }
                                            fmt.Println(v6)
                                            var v68 uint16 = v6
                                            _ = v68
                                        }
                                        switch v55 {
                                        case "foo bar":
                                            v5[0] = v55
                                        case "":
                                            v69 := GenS2{F0: (0.0 >= v7)}
                                            _ = v69
                                        default:
                                            var v70 float64 = ((v7 + (0.1 - v7)) / 1e10)
                                            _ = v70
                                        }
                                        _ = v55
                                    }
                                default:
                                    _ = len(v1)
                                }
                            } else {
                                switch fi32 {
                                case 0:
                                    v71 := GenS3{F0: 4294967295, F1: 1, F2: 9223372036854775807}
                                    _ = v71
                                case 32767:
                                    v5[2147483647] = ""
                                case 191:
                                    v35 = 2
                                }
                            }
                            _ = v39
                        }
                        fmt.Println(v31)
                    }
                }
                _ = fi34
                _ = v35
            }
            for fi72 := 0; fi72 < 7; fi72++ {
                var v73 int64 = -27
                var v74 int = (fi72 | (fi72 << ((760) & 63)))
                _ = fi72
                _ = v73
                _ = v74
            }
            _ = fi32
            _ = v33
        }
        if true {
            for fi75 := 0; fi75 < 19; fi75++ {
                _ = len(v3)
                var v76 int = fi75
                _ = fi75
                _ = v76
            }
            v77 := v5[0]
            if (32767 == 393) {
                if ("foo bar" < "\n\t") {
                    if (-980 <= int64(v31)) {
                        v77 = v77
                        v11 = -36
                    }
                    v4 = v4
                    switch ("\n\t" + v77) {
                    case "a":
                        var v78 int32 = -2147483648
                        _ = v78
                    case "foo bar":
                        v31 = 237
                    default:
                        for fi79 := 0; fi79 < 15; fi79++ {
                            v4 += ((v4 / ((((0.0 / float32(((float64(((v4 / (float32(v7) - (v4 * 0.0))) / float32(((v7 - 1e10) / v7)))) * 1e10) + 0.0))) * (1e10 * (((0.1 + float32(v7)) + v4) + (v4 / 1.0)))) / v4) + v4)) - ((3.14159 + float32(v7)) - (0.1 + (float32(v7) * v4))))
                            for fi80 := 0; fi80 < 3; fi80++ {
                                switch v77 {
                                case "hello":
                                    if ("\n\t" > (("a" + "\n\t") + v77)) {
                                        if (false || true) {
                                            if (false && true) {
                                                v81 := GenS6{F0: 42, F1: 42, F2: 978}
                                                v82 := GenS1{F0: 2147483648}
                                                _ = v81
                                                _ = v82
                                            } else {
                                                v83 := GenS7{F0: -9223372036854775808}
                                                _ = v83
                                            }
                                            v84 := GenS5{F0: -373}
                                            var v85 int = 190
                                            _ = v84
                                            _ = v85
                                        }
                                    } else {
                                        v31 -= (^(v31))
                                    }
                                    v7 += ((v7 / float64(float32(v7))) + (v7 / ((v7 + 1.5) - (1e10 * v7))))
                                default:
                                    v7 += (0.0 * 0.0)
                                }
                                _ = fi80
                            }
                            _ = fi79
                        }
                    }
                }
                if true {
                    v31 -= (v31 >> ((42) & 63))
                    switch (v77 + (v77 + (v77 + v77))) {
                    case "a":
                        v86 := GenS7{F0: 230}
                        _ = v86
                    case "foo bar":
                        v87 := GenS1{F0: 971}
                        _ = v87
                    case "":
                        if true {
                            var v88 int8 = 27
                            _ = v88
                        } else {
                            v4 = float32((((((v7 * v7) / 1e10) * v7) - v7) - v7))
                            if (132 == 138) {
                                var v89 uint64 = 42
                                _ = v89
                            } else {
                                v77 = "hello"
                                v90 := GenS3{F0: 0, F1: -406, F2: 0}
                                for fi91 := 0; fi91 < 14; fi91++ {
                                    fmt.Println(v7)
                                    var v92 int32 = 688
                                    v1[v77] = 2147483647
                                    _ = fi91
                                    _ = v92
                                }
                                _ = v90
                            }
                            var v93 int32 = 268
                            _ = v93
                        }
                    default:
                        v4 -= -2.25
                    }
                    v11 += -128
                } else {
                    delete(v1, "")
                    v94 := GenS2{F0: ((32767 <= 255) && false)}
                    _ = v94
                }
                switch 127 {
                case 127:
                    v95 := GenS4{F0: (2147483647 > -9223372036854775808), F1: v31, F2: (v11 << ((42) & 63))}
                    _ = v95
                case 128:
                    if true {
                        fmt.Println(v77, v77)
                    } else {
                        v77 = v77
                        var v96 uint = 18446744073709551615
                        var v97 string = v77
                        _ = v96
                        _ = v97
                    }
                default:
                    v98 := GenS1{F0: 42}
                    _ = v98
                }
            }
            _ = v77
        } else {
            if (255 != 207) {
                v99 := v1["foo bar"]
                v31 = (v31 - (v31 / (((v31 / (((v31 ^ 0)) | 1))) | 1)))
                _ = v99
            }
            if ("\n\t" > (("hello" + "a") + "hello")) {
                _ = len(v3)
            } else {
                v7 = v7
                var v100 uint8 = 114
                if true {
                    for fi101 := 0; fi101 < 16; fi101++ {
                        switch (fi101 - (fi101 - (^(fi101)))) {
                        case 285:
                            var v102 uint64 = v31
                            _ = v102
                        case 32768:
                            var v103 uint = uint(v6)
                            _ = v103
                        }
                        _ = fi101
                    }
                } else {
                    var v104 uint64 = v31
                    _ = v104
                }
                _ = v100
            }
        }
        _ = v31
    }
    var v105 float32 = (((v4 - float32(v7)) / (((((v4 - -2.25) - float32(v7)) - (-2.25 + v4)) / v4) / ((1e10 / v4) / v4))) / 3.14159)
    v4 += (v4 / ((v105 / (v4 - (v4 + 3.14159))) - ((v105 * v4) - v105)))
    if (2 > 2) {
        v106 := GenS2{F0: true}
        if (4294967295 == 65535) {
            fmt.Println(v7, v105, v11)
        }
        _ = v106
    }
    var v107 int16 = int16(v6)
    v2 = append(v2, 0)
    if false {
        var v108 int8 = (v11 / ((v11) | 1))
        if (false && false) {
            v4 -= float32((float64(v4) + 1e10))
        }
        var v109 int8 = (v108 << ((0) & 63))
        _ = v108
        _ = v109
    }
    var v110 uint64 = 42
    var v111 int16 = v107
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v105:", v105)
    fmt.Println("v107:", v107)
    fmt.Println("v11:", v11)
    fmt.Println("v110:", v110)
    fmt.Println("v111:", v111)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v105
    _ = v107
    _ = v11
    _ = v110
    _ = v111
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
