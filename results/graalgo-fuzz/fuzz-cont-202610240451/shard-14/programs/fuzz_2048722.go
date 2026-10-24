package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 uint64; F1 uint32; F2 int64 }
type GenS3 struct { F0 uint32; F1 int8; F2 uint64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint8; F1 uint32 }
type GenS7 struct { F0 int16; F1 uint16 }
type GenS8 struct { F0 int8 }

func main() {
    v1 := []string{"foo bar", ("\n\t" + (("foo bar" + "") + "foo bar"))}
    var v2 uint8 = 127
    var v3 uint8 = v2
    if false {
        v2 = 2
        if len(v1) > 0 {
            v1[0] = ("\n\t" + ("foo bar" + "foo bar"))
        }
        if true {
            v4 := GenS1{F0: 914, F1: 2147483647}
            for fi5 := 0; fi5 < 18; fi5++ {
                var v6 int16 = 1000
                v3 -= v2
                v7 := GenS1{F0: 9223372036854775807, F1: 42}
                _ = fi5
                _ = v6
                _ = v7
            }
            var v8 uint8 = (v3 << ((42) & 63))
            _ = v4
            _ = v8
        } else {
            var v9 uint32 = 688
            _ = v9
        }
    } else {
        var v10 float64 = float64(0.0)
        _ = v10
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        v12 := GenS2{F0: 812, F1: 46, F2: 112}
        if (!false) {
            if false {
                var v13 float32 = ((float32(((((((3.14159 - 0.0) + 0.0) * 3.14159) * 0.1) + (1e10 / 3.14159)) / -2.25)) - -2.25) - 3.14159)
                _ = v13
            } else {
                v3 = v2
                if (false || false) {
                    if true {
                        _ = len(v1)
                        if ((uint16(v2) == 1) || (false || false)) {
                            var v14 int32 = 42
                            _ = v14
                        } else {
                            if false {
                                for fi15 := 0; fi15 < 11; fi15++ {
                                    v3 = (^(v2))
                                    _ = fi15
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            } else {
                                for fi16 := 0; fi16 < 15; fi16++ {
                                    var v17 uint32 = 103
                                    v18 := GenS2{F0: 798, F1: v17, F2: -9223372036854775808}
                                    _ = fi16
                                    _ = v17
                                    _ = v18
                                }
                            }
                        }
                    } else {
                        v1 = append(v1, "\n\t")
                    }
                }
            }
        } else {
            var v19 bool = true
            switch ("foo bar" + (("\n\t" + "hello") + (("hello" + "") + ("foo bar" + "")))) {
            case "hello":
                v2 -= 0
            default:
                if len(v1) > 0 {
                    v1[0] = ("hello" + "foo bar")
                }
            }
            var v20 uint16 = 159
            _ = v19
            _ = v20
        }
        for fi21 := 0; fi21 < 11; fi21++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if len(v1) > 0 {
                v1[0] = ""
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            switch fi21 {
            case -596:
                var v22 uint8 = v3
                _ = v22
            }
            v2 += 164
            _ = fi21
        }
        _ = fi11
        _ = v12
    }
    if (false || false) {
        if false {
            if false {
                for fi23 := 0; fi23 < 11; fi23++ {
                    v24 := GenS3{F0: 42, F1: 104, F2: 1}
                    var v25 int16 = -595
                    var v26 bool = (-713 != (fi23 << ((4) & 63)))
                    _ = fi23
                    _ = v24
                    _ = v25
                    _ = v26
                }
                var v27 int16 = 32767
                var v28 int64 = -9223372036854775808
                _ = v27
                _ = v28
            } else {
                for fi29 := 0; fi29 < 3; fi29++ {
                    v30 := GenS4{F0: 65535}
                    fmt.Println(v3)
                    _ = fi29
                    _ = v30
                }
            }
            switch "foo bar" {
            case "foo bar":
                fmt.Println(v2, v2)
            case "a":
                v2 = 0
            case "\n\t":
                v3 = (v3 & v3)
            default:
                var v31 int64 = 0
                _ = v31
            }
            if (-2147483648 >= 1) {
                v2 += (v3 * 0)
            } else {
                var v32 uint = 4294967295
                _ = v32
            }
        } else {
            fmt.Println(v3, v3)
            v3 += v3
        }
        switch -2147483648 {
        case -988:
            for fi33 := 0; fi33 < 18; fi33++ {
                if false {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    fmt.Println(v2, v3)
                }
                fmt.Println(v3, v3)
                _ = fi33
            }
        case 32767:
            if ((!false) && true) {
                for fi34 := 0; fi34 < 18; fi34++ {
                    if false {
                        v3 -= (v2 | (^(v2)))
                        for fi35 := 0; fi35 < 18; fi35++ {
                            if true {
                                v2 = 1
                            } else {
                                v36 := GenS5{F0: 4294967295}
                                if (!true) {
                                    if (v2 < v3) {
                                        var v37 int64 = -9223372036854775808
                                        v1 = append(v1, ("hello" + "a"))
                                        _ = v37
                                    } else {
                                        v1 = append(v1, "foo bar")
                                        if false {
                                            v38 := GenS6{F0: v3, F1: 276}
                                            _ = len(v1)
                                            v2 -= (v2 | (v2 + v3))
                                            _ = v38
                                        }
                                        v3 = 42
                                    }
                                } else {
                                    if false {
                                        var v39 int16 = -32768
                                        _ = v39
                                    }
                                }
                                var v40 uint32 = 980
                                _ = v36
                                _ = v40
                            }
                            if (-2147483648 < fi35) {
                                v41 := GenS2{F0: 0, F1: 42, F2: 9223372036854775807}
                                v2 += (^(v2))
                                switch (fi34 + (fi34 % ((fi34) | 1))) {
                                case 32767:
                                    if (!true) {
                                        v42 := GenS1{F0: 42, F1: 2147483647}
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        fmt.Println(fi35)
                                        _ = v42
                                    }
                                default:
                                    v43 := GenS1{F0: -842, F1: uint32(v3)}
                                    _ = v43
                                }
                                _ = v41
                            }
                            v3 = (v2 % ((54) | 1))
                            _ = fi35
                        }
                        var v44 int64 = -9223372036854775808
                        _ = v44
                    } else {
                        v45 := GenS4{F0: uint16(v3)}
                        if false {
                            switch "foo bar" {
                            case "a":
                                if len(v1) > 0 {
                                    v1[0] = "hello"
                                }
                            case "\n\t":
                                var v46 int8 = -1
                                _ = v46
                            case "foo bar":
                                v2 += v2
                            }
                        } else {
                            var v47 int32 = -929
                            v3 -= (v3 - 128)
                            _ = v47
                        }
                        switch fi34 {
                        case -9223372036854775808:
                            v48 := GenS5{F0: 1}
                            _ = v48
                        }
                        _ = v45
                    }
                    for fi49 := 0; fi49 < 2; fi49++ {
                        if len(v1) > 0 {
                            v1[0] = "a"
                        }
                        var v50 int = (-(fi34))
                        if false {
                            var v51 uint16 = 0
                            v52 := GenS2{F0: 621, F1: 203, F2: 42}
                            _ = v51
                            _ = v52
                        } else {
                            var v53 float32 = (((3.14159 - ((float32(0.1) * 0.0) / ((0.0 * ((-2.25 + ((0.0 + -2.25) + -2.25)) + -2.25)) - 0.1))) - (float32(0.1) + 0.0)) * (0.1 - 0.1))
                            if false {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                var v54 int8 = 100
                                _ = v54
                            } else {
                                var v55 uint64 = 2
                                var v56 float32 = (3.14159 * -2.25)
                                _ = v55
                                _ = v56
                            }
                            if (((1e10 + 1e10) == 3.14159) && true) {
                                fmt.Println(fi49, v3, fi34)
                                var v57 int64 = -1
                                v58 := GenS5{F0: 1}
                                _ = len(v1)
                                _ = v57
                                _ = v58
                            } else {
                                var v59 int8 = 42
                                var v60 string = "\n\t"
                                _ = v59
                                _ = v60
                            }
                            _ = v53
                        }
                        switch fi49 {
                        case -145:
                            if true {
                                switch (v50 ^ 2147483647) {
                                case 128:
                                    switch ("\n\t" + "\n\t") {
                                    case "foo bar":
                                        v61 := GenS2{F0: 42, F1: 65535, F2: 9223372036854775807}
                                        _ = v61
                                    case "a":
                                        if len(v1) > 0 {
                                            v1[0] = "a"
                                        }
                                        v50 = (v50 ^ v50)
                                        if true {
                                            v2 = (^(v2))
                                            v1 = append(v1, (("a" + "a") + "foo bar"))
                                            fmt.Println(v3)
                                            var v62 int64 = -261
                                            _ = v62
                                        }
                                    case "\n\t":
                                        switch -2147483648 {
                                        case 32768:
                                            var v63 int8 = 88
                                            _ = v63
                                        default:
                                            switch ("hello" + "foo bar") {
                                            case "":
                                                switch 855 {
                                                case -767:
                                                    if ((fi34 >> ((343) & 63)) == 1) {
                                                        v1 = append(v1, "")
                                                        v1 = append(v1, ("hello" + "hello"))
                                                        v3 = 128
                                                    }
                                                case 127:
                                                    if (true || (162 >= v3)) {
                                                        _ = len(v1)
                                                    }
                                                default:
                                                    switch (fi34 + 1) {
                                                    case -2147483648:
                                                        v64 := GenS3{F0: 2147483647, F1: 1, F2: 1}
                                                        _ = v64
                                                    case 0:
                                                        if len(v1) > 0 {
                                                            v1[0] = (("foo bar" + ("a" + "hello")) + "")
                                                        }
                                                    }
                                                }
                                            case "hello":
                                                var v65 int64 = 42
                                                _ = v65
                                            default:
                                                var v66 uint64 = 0
                                                _ = v66
                                            }
                                        }
                                        var v67 uint64 = 18446744073709551615
                                        var v68 int = (^(v50))
                                        switch (v68 & 32768) {
                                        case 1:
                                            fmt.Println(v67, fi49)
                                        }
                                        var v69 int16 = 922
                                        v70 := GenS4{F0: 42}
                                        _ = v67
                                        _ = v68
                                        _ = v69
                                        _ = v70
                                    default:
                                        if true {
                                            if (!false) {
                                                var v71 uint = 4294967295
                                                var v72 float64 = 0.0
                                                _ = v71
                                                _ = v72
                                            } else {
                                                if false {
                                                    if (((0.1 - 0.1) <= 1e10) && (2147483647 >= 1)) {
                                                        fmt.Println(v50)
                                                    } else {
                                                        v50 -= fi49
                                                        v2 = v3
                                                    }
                                                    fmt.Println(v2, v50, fi34)
                                                } else {
                                                    v73 := GenS3{F0: 2, F1: int8(fi49), F2: 1}
                                                    if (3.14159 < 3.14159) {
                                                        fmt.Println(v3, v50, v50)
                                                    }
                                                    var v74 uint16 = 0
                                                    _ = v73
                                                    _ = v74
                                                }
                                            }
                                            v3 = v2
                                            var v75 uint32 = 24
                                            v76 := GenS3{F0: v75, F1: 66, F2: 2}
                                            _ = v75
                                            _ = v76
                                        }
                                    }
                                case 0:
                                    v77 := GenS1{F0: 1, F1: 0}
                                    _ = v77
                                }
                                v78 := GenS7{F0: -306, F1: 255}
                                var v79 uint8 = v3
                                _ = v78
                                _ = v79
                            } else {
                                v80 := GenS2{F0: 796, F1: 506, F2: 0}
                                var v81 uint32 = 83
                                _ = v80
                                _ = v81
                            }
                        case 32768:
                            v50 = fi49
                        default:
                            var v82 uint16 = 65535
                            _ = v82
                        }
                        _ = fi49
                        _ = v50
                    }
                    for fi83 := 0; fi83 < 18; fi83++ {
                        if true {
                            fmt.Println(v2, v2)
                            v2 -= v2
                            v3 += 13
                        } else {
                            var v84 uint8 = 215
                            fmt.Println(v84)
                            var v85 int = fi83
                            var v86 uint8 = v84
                            switch ("" + ("foo bar" + "hello")) {
                            case "":
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            default:
                                v84 += 168
                            }
                            _ = v84
                            _ = v85
                            _ = v86
                        }
                        v2 = (v2 * (v2 + v2))
                        fmt.Println(v2)
                        _ = fi83
                    }
                    _ = fi34
                }
            }
        default:
            v87 := GenS1{F0: 9223372036854775807, F1: 311}
            _ = v87
        }
        v88 := GenS8{F0: -128}
        _ = v88
    } else {
        for fi89 := 0; fi89 < 9; fi89++ {
            for fi90 := 0; fi90 < 8; fi90++ {
                var v91 int16 = 1000
                _ = fi90
                _ = v91
            }
            v2 = (v3 | 127)
            if (true && (1 < 65535)) {
                v3 -= v2
            }
            _ = fi89
        }
    }
    var v92 float64 = 0.0
    var v93 uint32 = 0
    v3 += 2
    var v94 bool = ((-1 != int64(v93)) && false)
    for fi95 := 0; fi95 < 11; fi95++ {
        for fi96 := 0; fi96 < 19; fi96++ {
            fmt.Println(v93, fi95)
            _ = fi96
        }
        var v97 int8 = 127
        _ = fi95
        _ = v97
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v2
    _ = v3
    _ = v92
    _ = v93
    _ = v94
}
