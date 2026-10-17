package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 int; F1 uint }
type GenS4 struct { F0 uint; F1 uint32 }
type GenS5 struct { F0 uint64; F1 uint8 }
type GenS6 struct { F0 int32; F1 int8 }
type GenS7 struct { F0 int16; F1 string; F2 int }
type GenS8 struct { F0 uint8; F1 int8 }

func main() {
    var v1 float32 = -2.25
    v2 := []string{""}
    var v3 int = 55
    var v4 int32 = 929
    var v5 int32 = v4
    var v6 uint8 = 223
    fmt.Println(v3, v1, v4)
    _ = len(v2)
    if true {
        var v7 bool = (((v3 + int(v4)) != v3) && (0.0 <= ((v1 + v1) / (v1 / (v1 + v1)))))
        v3 += 42
        _ = v7
    } else {
        v5 = v5
    }
    v6 += (v6 ^ v6)
    for fi8 := 0; fi8 < 5; fi8++ {
        if (1 < 2147483648) {
            if true {
                switch "foo bar" {
                case "a":
                    if true {
                        if false {
                            v2 = append(v2, ("hello" + "a"))
                        } else {
                            fmt.Println(v5, v4)
                        }
                        if true {
                            if len(v2) > 0 {
                                v2[0] = (("a" + ("a" + (("" + "a") + "\n\t"))) + ("" + "hello"))
                            }
                        } else {
                            for fi9 := 0; fi9 < 6; fi9++ {
                                v10 := GenS1{F0: (("hello" + "") + "\n\t")}
                                _ = fi9
                                _ = v10
                            }
                        }
                    }
                case "hello":
                    v3 -= (v3 | (fi8 - -876))
                case "foo bar":
                    v2 = append(v2, (("foo bar" + "\n\t") + ("a" + "")))
                }
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if true {
                if false {
                    v6 = (v6 << ((0) & 63))
                    v6 = (v6 + (v6 << ((2) & 63)))
                    switch (v3 * (fi8 + (v3 >> ((185) & 63)))) {
                    case -239:
                        var v11 uint64 = 4294967295
                        _ = v11
                    case 630:
                        fmt.Println(v5, v3)
                    case 2147483647:
                        var v12 int32 = -833
                        _ = v12
                    }
                } else {
                    fmt.Println(v6, fi8)
                    if false {
                        v13 := GenS1{F0: "hello"}
                        _ = v13
                    } else {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if ((fi8 - v3) > v3) {
                            for fi14 := 0; fi14 < 18; fi14++ {
                                v6 -= (^(v6))
                                v2 = append(v2, ("foo bar" + ("hello" + "a")))
                                _ = fi14
                            }
                        } else {
                            v15 := GenS1{F0: "hello"}
                            _ = v15
                        }
                        if true {
                            var v16 float32 = v1
                            if (!true) {
                                var v17 int16 = 1000
                                if ("foo bar" <= "\n\t") {
                                    var v18 uint16 = 329
                                    var v19 float32 = (0.1 - 0.1)
                                    _ = v18
                                    _ = v19
                                } else {
                                    fmt.Println(v6, v5)
                                    if false {
                                        v20 := GenS1{F0: "a"}
                                        _ = v20
                                    } else {
                                        var v21 int = 32768
                                        if true {
                                            v22 := GenS2{F0: 42, F1: 0}
                                            _ = v22
                                        } else {
                                            if (!true) {
                                                var v23 int32 = 0
                                                _ = v23
                                            }
                                        }
                                        if false {
                                            for fi24 := 0; fi24 < 11; fi24++ {
                                                if (true || (4294967295 <= 408)) {
                                                    if false {
                                                        v25 := GenS1{F0: (("foo bar" + (("foo bar" + "\n\t") + ("" + "foo bar"))) + "foo bar")}
                                                        v2 = append(v2, ("a" + ("foo bar" + ((("a" + "\n\t") + "hello") + (("foo bar" + "a") + ("hello" + (("foo bar" + "hello") + "hello")))))))
                                                        if false {
                                                            if ("foo bar" <= (("hello" + "hello") + ("a" + "\n\t"))) {
                                                                var v26 int32 = -2147483648
                                                                _ = v26
                                                            } else {
                                                                var v27 int16 = v17
                                                                fmt.Println(v1, v4, v4)
                                                                _ = v27
                                                            }
                                                            var v28 int = (fi24 - int(v6))
                                                            _ = v28
                                                        } else {
                                                            var v29 string = ""
                                                            _ = v29
                                                        }
                                                        _ = v25
                                                    } else {
                                                        if (!((fi8 & 128) > (^(fi24)))) {
                                                            switch "" {
                                                            case "hello":
                                                                if true {
                                                                    fmt.Println(v1)
                                                                    v2 = append(v2, "\n\t")
                                                                    if false {
                                                                        var v30 uint16 = 32767
                                                                        _ = v30
                                                                    } else {
                                                                        v4 -= v4
                                                                        switch (("" + "\n\t") + "foo bar") {
                                                                        case "\n\t":
                                                                            switch "foo bar" {
                                                                            case "\n\t":
                                                                                v16 -= (v1 * v16)
                                                                            }
                                                                        case "foo bar":
                                                                            v31 := GenS1{F0: ""}
                                                                            _ = v31
                                                                        case "hello":
                                                                            v32 := GenS1{F0: "foo bar"}
                                                                            _ = v32
                                                                        }
                                                                        var v33 bool = true
                                                                        if (255 == 42) {
                                                                            if (0 <= 42) {
                                                                                v17 = (v17 % (((-(v17))) | 1))
                                                                            } else {
                                                                                var v34 uint = 57
                                                                                v3 += 128
                                                                                v16 -= v16
                                                                                _ = v34
                                                                            }
                                                                            v6 += 5
                                                                            var v35 uint = 509
                                                                            _ = v35
                                                                        }
                                                                        _ = v33
                                                                    }
                                                                } else {
                                                                    v36 := GenS3{F0: v21, F1: 0}
                                                                    _ = v36
                                                                }
                                                            case "a":
                                                                if (0 > 127) {
                                                                    v37 := GenS4{F0: 42, F1: 4294967295}
                                                                    v3 += (v21 / ((v3) | 1))
                                                                    _ = v37
                                                                }
                                                            case "\n\t":
                                                                switch "\n\t" {
                                                                case "foo bar":
                                                                    v38 := GenS5{F0: 42, F1: (v6 + v6)}
                                                                    _ = v38
                                                                case "a":
                                                                    if false {
                                                                        switch "hello" {
                                                                        case "a":
                                                                            fmt.Println(fi8, v4, v6)
                                                                        case "hello":
                                                                            var v39 int16 = v17
                                                                            _ = v39
                                                                        default:
                                                                            v6 += (v6 / (((v6 % (((^(v6))) | 1))) | 1))
                                                                        }
                                                                        var v40 int8 = 127
                                                                        _ = v40
                                                                    } else {
                                                                        var v41 uint64 = 18446744073709551615
                                                                        _ = v41
                                                                    }
                                                                case "hello":
                                                                    var v42 bool = (!(102 != 42))
                                                                    _ = v42
                                                                default:
                                                                    var v43 float32 = ((((v1 + (float32(0.0) / (3.14159 * 0.1))) + v16) * v1) / ((((-2.25 + (3.14159 * -2.25)) / -2.25) / v1) - v16))
                                                                    _ = v43
                                                                }
                                                            default:
                                                                v16 += v1
                                                            }
                                                            v44 := GenS2{F0: (-(v4)), F1: 255}
                                                            var v45 uint64 = 994
                                                            _ = v44
                                                            _ = v45
                                                        } else {
                                                            if true {
                                                                if (42 > v17) {
                                                                    switch (v21 & (v3 % ((v3) | 1))) {
                                                                    case -2147483648:
                                                                        switch "foo bar" {
                                                                        case "hello":
                                                                            if (true || false) {
                                                                                v46 := GenS6{F0: (v5 % ((42) | 1)), F1: -1}
                                                                                var v47 uint = 1
                                                                                _ = v46
                                                                                _ = v47
                                                                            } else {
                                                                                var v48 int = v3
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                                v49 := GenS1{F0: ""}
                                                                                _ = v48
                                                                                _ = v49
                                                                            }
                                                                        case "a":
                                                                            if false {
                                                                                fmt.Println(v17, v16, v21)
                                                                                var v50 int16 = v17
                                                                                _ = v50
                                                                            } else {
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                                var v51 int8 = -85
                                                                                _ = v51
                                                                            }
                                                                        default:
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                        }
                                                                    case -826:
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                    case 539:
                                                                        if ((true && true) || (!(175 == 1))) {
                                                                            v52 := GenS2{F0: (^(v4)), F1: 0}
                                                                            v1 += v16
                                                                            _ = v52
                                                                        }
                                                                    }
                                                                    v17 += 42
                                                                    fmt.Println(fi8)
                                                                } else {
                                                                    var v53 float32 = (float32((0.0 * 3.14159)) * v1)
                                                                    if (true || (fi8 > 2147483647)) {
                                                                        if (0 == -1) {
                                                                            v2 = append(v2, ("a" + ""))
                                                                            switch ("" + "") {
                                                                            case "":
                                                                                v16 -= (v1 + ((v53 - 1e10) - v53))
                                                                            }
                                                                            fmt.Println(v4, v3, v53)
                                                                            var v54 int8 = -20
                                                                            if (-2.25 >= (0.1 + 3.14159)) {
                                                                                var v55 int8 = v54
                                                                                if (59 <= 42) {
                                                                                    if (18446744073709551615 != 870) {
                                                                                        v56 := GenS2{F0: v4, F1: 1}
                                                                                        _ = v56
                                                                                    }
                                                                                } else {
                                                                                    var v57 uint8 = v6
                                                                                    _ = v57
                                                                                }
                                                                                v58 := GenS7{F0: -32768, F1: "foo bar", F2: fi24}
                                                                                var v59 int8 = (v55 | (v54 + v55))
                                                                                v60 := GenS5{F0: 809, F1: (v6 >> ((42) & 63))}
                                                                                fmt.Println(fi24, v53)
                                                                                v3 -= (fi8 / ((fi24) | 1))
                                                                                _ = v55
                                                                                _ = v58
                                                                                _ = v59
                                                                                _ = v60
                                                                            }
                                                                            _ = v54
                                                                        } else {
                                                                            fmt.Println(v21, v3, v21)
                                                                            fmt.Println(v16, v3, fi8)
                                                                        }
                                                                    }
                                                                    v4 = v4
                                                                    _ = v53
                                                                }
                                                                var v61 uint8 = 128
                                                                var v62 uint8 = 128
                                                                _ = v61
                                                                _ = v62
                                                            } else {
                                                                fmt.Println(v16, fi24)
                                                            }
                                                        }
                                                        fmt.Println(v1, fi8)
                                                        v63 := GenS3{F0: v21, F1: 656}
                                                        _ = v63
                                                    }
                                                    fmt.Println(v17)
                                                } else {
                                                    var v64 string = "\n\t"
                                                    v16 += v1
                                                    _ = v64
                                                }
                                                _ = fi24
                                            }
                                            if (!(-9223372036854775808 != 518)) {
                                                for fi65 := 0; fi65 < 11; fi65++ {
                                                    v66 := GenS3{F0: (v21 + (fi65 & fi65)), F1: 879}
                                                    v67 := GenS3{F0: (v3 % ((0) | 1)), F1: 572}
                                                    _ = fi65
                                                    _ = v66
                                                    _ = v67
                                                }
                                            }
                                            v17 -= (-(v17))
                                        }
                                        _ = v21
                                    }
                                }
                                if (-868 >= 42) {
                                    if (true && (-1 <= v4)) {
                                        switch "\n\t" {
                                        case "a":
                                            v1 = v16
                                        case "\n\t":
                                            var v68 uint32 = 717
                                            _ = v68
                                        default:
                                            v2 = append(v2, ("foo bar" + ("\n\t" + "a")))
                                        }
                                    }
                                    var v69 float32 = v1
                                    _ = v69
                                }
                                _ = v17
                            } else {
                                var v70 int64 = -773
                                var v71 int32 = 2147483647
                                var v72 int32 = 150
                                _ = v70
                                _ = v71
                                _ = v72
                            }
                            _ = v16
                        }
                    }
                }
                v5 = (v5 >> ((18446744073709551615) & 63))
                var v73 int64 = 42
                _ = v73
            }
        } else {
            v3 -= v3
            v2 = append(v2, "foo bar")
            if (-324 >= 0) {
                for fi74 := 0; fi74 < 17; fi74++ {
                    switch ("hello" + "") {
                    case "":
                        v1 = v1
                    case "a":
                        if ((v5 <= (v4 - v5)) && ((416 <= 832) || true)) {
                            if len(v2) > 0 {
                                v2[0] = ""
                            }
                            v75 := GenS6{F0: (^(v5)), F1: 1}
                            fmt.Println(v5)
                            _ = v75
                        }
                        v2 = append(v2, "")
                    case "\n\t":
                        if (65535 == 484) {
                            if false {
                                v4 += v4
                                var v76 uint32 = 747
                                var v77 string = "foo bar"
                                _ = v76
                                _ = v77
                            } else {
                                var v78 int = fi74
                                var v79 uint8 = (^(v6))
                                if len(v2) > 0 {
                                    v2[0] = ("a" + ("hello" + (("\n\t" + "a") + "a")))
                                }
                                if (false && (2 < 2)) {
                                    v5 -= (v4 << ((18446744073709551615) & 63))
                                    v5 -= -1
                                    v80 := GenS6{F0: (v5 & (^(v4))), F1: 100}
                                    _ = v80
                                }
                                _ = v78
                                _ = v79
                            }
                            var v81 uint16 = 32767
                            _ = v81
                        }
                    }
                    if true {
                        v2 = append(v2, "a")
                    } else {
                        fmt.Println(v4)
                        var v82 int = fi8
                        _ = v82
                    }
                    var v83 int64 = -605
                    _ = fi74
                    _ = v83
                }
                for fi84 := 0; fi84 < 11; fi84++ {
                    if true {
                        if (975 < 1) {
                            v85 := GenS6{F0: 2147483647, F1: 1}
                            v1 = 3.14159
                            v6 = 169
                            var v86 int = (fi8 & (fi84 * -1))
                            _ = v85
                            _ = v86
                        }
                        var v87 uint32 = 899
                        _ = v87
                    } else {
                        v6 = v6
                        fmt.Println(v5, v5, fi8)
                    }
                    if (1 < 9223372036854775807) {
                        v6 = v6
                        v4 -= -2147483648
                        if false {
                            var v88 uint8 = (v6 >> ((423) & 63))
                            v89 := GenS8{F0: (v88 + v6), F1: -1}
                            v88 -= v88
                            _ = v88
                            _ = v89
                        } else {
                            if ((315 < 255) && true) {
                                var v90 bool = (true || false)
                                _ = v90
                            } else {
                                var v91 int16 = 906
                                _ = v91
                            }
                        }
                    }
                    _ = fi84
                }
            } else {
                v3 = fi8
                switch fi8 {
                case 4294967295:
                    v2 = append(v2, "hello")
                case 9223372036854775807:
                    if (((fi8 | (v3 & (fi8 | 42))) >= (v3 | fi8)) && false) {
                        var v92 uint8 = v6
                        v4 -= (v5 / (((v4 & 1)) | 1))
                        _ = v92
                    } else {
                        v3 -= fi8
                        var v93 bool = false
                        _ = v93
                    }
                default:
                    if ((1 != 2) && true) {
                        v4 = (^(v5))
                        v2 = append(v2, "\n\t")
                        var v94 bool = false
                        _ = v94
                    }
                }
                switch -261 {
                case 0:
                    switch "hello" {
                    case "hello":
                        switch fi8 {
                        case 127:
                            for fi95 := 0; fi95 < 1; fi95++ {
                                var v96 int64 = 710
                                _ = fi95
                                _ = v96
                            }
                        case 128:
                            v3 = fi8
                        default:
                            var v97 uint8 = (^(v6))
                            _ = v97
                        }
                    default:
                        v5 = (-(v4))
                    }
                case 2147483647:
                    var v98 uint = 741
                    _ = v98
                case 32768:
                    fmt.Println(v4, v3, v4)
                }
            }
        }
        v5 = 1
        v1 = 0.1
        _ = fi8
    }
    v99 := GenS5{F0: 2, F1: v6}
    v3 -= (-(v3))
    fmt.Println(v6)
    _ = len(v2)
    v6 -= v6
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v99
}
