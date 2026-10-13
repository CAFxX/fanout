package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 string; F2 int16 }
type GenS3 struct { F0 int8; F1 bool; F2 uint16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int64; F1 int64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint16; F1 uint64; F2 string }
type GenS8 struct { F0 uint64; F1 int; F2 int8 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int; F1 string; F2 uint32 }
type GenS12 struct { F0 uint; F1 uint64 }
type GenS13 struct { F0 string; F1 int64; F2 uint32 }
type GenS14 struct { F0 int32; F1 uint64; F2 int16 }
type GenS15 struct { F0 int8; F1 int16 }
type GenS16 struct { F0 uint16; F1 int8 }
type GenS17 struct { F0 int }
func genFn0(p0 string, p1 string) int64 {
    v7 := []int{-1, 2147483647, -476, 32768}
    if len(v7) > 0 {
        v7[0] = 1
    }
    var v8 uint16 = 42
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 1
}
func genFn1() uint16 {
    var v9 uint32 = 2147483647
    var v10 float32 = -2.25
    var v11 string = "a"
    _ = v10
    _ = v11
    _ = v9
    return 938
}

func main() {
    var v1 int32 = -150
    v2 := []uint32{579, 2147483648}
    var v3 uint = 1
    var v4 bool = false
    var v5 int = -9223372036854775808
    v6 := []string{("\n\t" + "a"), "a", "a", "foo bar"}
    v12 := GenS1{F0: 0}
    for fi13 := 0; fi13 < 16; fi13++ {
        v14 := GenS2{F0: (((1e10 / -2.25) <= 0.0) && v4), F1: "a", F2: -548}
        var v15 int16 = -32768
        _ = len(v6)
        _ = fi13
        _ = v14
        _ = v15
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v16 uint8 = 11
    v17 := GenS3{F0: 127, F1: (!v4), F2: 2}
    v2 = append(v2, 0)
    for fi18 := 0; fi18 < 4; fi18++ {
        v3 = v3
        v4 = true
        _ = fi18
    }
    v19 := GenS1{F0: 0}
    v20 := GenS2{F0: v4, F1: "\n\t", F2: 823}
    v16 += 255
    switch v5 {
    case -2147483648:
        var v21 int16 = 1
        _ = v21
    case 4294967295:
        for fi22 := 0; fi22 < 4; fi22++ {
            if (v3 >= (v3 | v3)) {
                if v4 {
                    for fi23 := 0; fi23 < 16; fi23++ {
                        var v24 int = fi23
                        switch ("hello" + ("hello" + ("foo bar" + "\n\t"))) {
                        case "":
                            if ((0 == 0) && v4) {
                                var v25 float64 = -2.25
                                var v26 int32 = 42
                                if (255 <= v16) {
                                    switch int(v1) {
                                    case -1:
                                        if len(v6) > 0 {
                                            v6[0] = ""
                                        }
                                    }
                                    if false {
                                        var v27 float32 = 1e10
                                        var v28 int16 = -720
                                        _ = v27
                                        _ = v28
                                    } else {
                                        v29 := GenS4{F0: v26}
                                        _ = v29
                                    }
                                    if v4 {
                                        v30 := GenS5{F0: -519, F1: 46}
                                        _ = v30
                                    }
                                    switch "hello" {
                                    case "a":
                                        v31 := GenS6{F0: 100}
                                        _ = v31
                                    case "\n\t":
                                        v32 := GenS1{F0: 0}
                                        _ = v32
                                    default:
                                        switch "a" {
                                        case "foo bar":
                                            _ = len(v2)
                                            v26 = (v1 * (v1 - -1))
                                            v33 := GenS7{F0: 65535, F1: 295, F2: ("\n\t" + "\n\t")}
                                            _ = v33
                                        case "hello":
                                            var v34 uint32 = 771
                                            _ = v34
                                        }
                                    }
                                    if len(v2) > 0 {
                                        v2[0] = 441
                                    }
                                } else {
                                    v3 = (v3 & (v3 | v3))
                                    var v35 int32 = v1
                                    fmt.Println(fi22, v24)
                                    _ = v35
                                }
                                _ = v25
                                _ = v26
                            }
                            fmt.Println(v3, fi22, v3)
                            switch fi23 {
                            case -9223372036854775808:
                                fmt.Println(v1, fi22)
                            case 849:
                                var v36 int = (v24 * v5)
                                _ = v36
                            case 725:
                                var v37 float32 = 1e10
                                _ = v37
                            }
                        case "\n\t":
                            if (2 <= 810) {
                                switch fi22 {
                                case 1:
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                case -2147483648:
                                    var v38 uint8 = (v16 - 255)
                                    _ = v38
                                }
                                switch "" {
                                case "\n\t":
                                    var v39 int = v5
                                    _ = v39
                                case "foo bar":
                                    var v40 int8 = 11
                                    _ = v40
                                case "hello":
                                    switch ("foo bar" + (("\n\t" + ("hello" + "")) + "\n\t")) {
                                    case "hello":
                                        var v41 int16 = 1000
                                        _ = v41
                                    }
                                default:
                                    switch fi22 {
                                    case -1:
                                        v16 -= (v16 * (v16 % (((v16 * v16)) | 1)))
                                    }
                                }
                                v42 := GenS8{F0: 263, F1: (^(fi23)), F2: 0}
                                v43 := GenS9{F0: 0}
                                _ = v42
                                _ = v43
                            }
                        case "foo bar":
                            if (980 < 819) {
                                if (42 < uint64(v16)) {
                                    switch "\n\t" {
                                    case "":
                                        _ = len(v6)
                                    case "foo bar":
                                        if v4 {
                                            if v4 {
                                                var v44 uint64 = 42
                                                if v4 {
                                                    if v4 {
                                                        v45 := GenS10{F0: v4}
                                                        v6 = append(v6, (("\n\t" + "\n\t") + "hello"))
                                                        v3 = (v3 << (((^(v3))) & 63))
                                                        _ = v45
                                                    } else {
                                                        var v46 int16 = 42
                                                        _ = v46
                                                    }
                                                } else {
                                                    if (!v4) {
                                                        v47 := GenS11{F0: fi22, F1: ("hello" + (("" + "foo bar") + "\n\t")), F2: 4294967295}
                                                        var v48 int = v24
                                                        v49 := GenS12{F0: (v3 >> ((v3) & 63)), F1: (v44 + v44)}
                                                        var v50 uint8 = v16
                                                        _ = v47
                                                        _ = v48
                                                        _ = v49
                                                        _ = v50
                                                    } else {
                                                        if true {
                                                            var v51 float64 = ((-2.25 / 3.14159) * (3.14159 - (((-2.25 / (0.0 + (float64(0.1) * (float64((0.1 - 0.1)) - float64((1.5 - (float32(-2.25) - float32((0.0 / 1.0))))))))) + 1e10) + (1e10 * (1.5 - 0.1)))))
                                                            var v52 uint16 = 413
                                                            fmt.Println(v24, v44, v51)
                                                            _ = v51
                                                            _ = v52
                                                        }
                                                        var v53 int16 = int16(v44)
                                                        var v54 uint64 = v44
                                                        _ = v53
                                                        _ = v54
                                                    }
                                                    var v55 bool = ("" != "a")
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                    _ = v55
                                                }
                                                if v4 {
                                                    if true {
                                                        v56 := GenS5{F0: -821, F1: 9223372036854775807}
                                                        if (0 == v44) {
                                                            if (9223372036854775807 >= -82) {
                                                                v57 := GenS13{F0: "a", F1: 42, F2: 65535}
                                                                _ = v57
                                                            }
                                                            if false {
                                                                var v58 int16 = 42
                                                                var v59 int64 = 9223372036854775807
                                                                _ = v58
                                                                _ = v59
                                                            }
                                                        }
                                                        v60 := GenS14{F0: v1, F1: 2, F2: -881}
                                                        _ = v56
                                                        _ = v60
                                                    } else {
                                                        switch v24 {
                                                        case -479:
                                                            switch "a" {
                                                            case "a":
                                                                var v61 int8 = -122
                                                                _ = v61
                                                            }
                                                        case 9223372036854775807:
                                                            v1 -= (-(v1))
                                                        }
                                                    }
                                                    var v62 uint8 = v16
                                                    _ = v62
                                                } else {
                                                    v1 -= (-(v1))
                                                }
                                                _ = v44
                                            } else {
                                                var v63 float64 = 0.1
                                                _ = v63
                                            }
                                            var v64 int8 = 0
                                            _ = v64
                                        }
                                    }
                                    _ = len(v2)
                                    v65 := GenS8{F0: 517, F1: fi23, F2: 23}
                                    _ = v65
                                }
                                var v66 int8 = 100
                                _ = v66
                            } else {
                                var v67 uint32 = uint32(v3)
                                var v68 int8 = 100
                                _ = v67
                                _ = v68
                            }
                        default:
                            v69 := GenS15{F0: 82, F1: -991}
                            _ = v69
                        }
                        _ = fi23
                        _ = v24
                    }
                    if len(v2) > 0 {
                        v2[0] = 42
                    }
                    v1 += (v1 >> ((v3) & 63))
                }
            } else {
                for fi70 := 0; fi70 < 3; fi70++ {
                    v71 := GenS3{F0: 127, F1: (-1 < 62), F2: 143}
                    v16 -= 2
                    _ = fi70
                    _ = v71
                }
                v4 = ((-9223372036854775808 != 1) && v4)
                _ = len(v2)
            }
            _ = fi22
        }
    }
    v16 += v16
    for fi72 := 0; fi72 < 4; fi72++ {
        var v73 string = "\n\t"
        var v74 uint = (v3 * (v3 % ((v3) | 1)))
        _ = fi72
        _ = v73
        _ = v74
    }
    for fi75 := 0; fi75 < 18; fi75++ {
        switch "hello" {
        case "\n\t":
            v76 := GenS16{F0: genFn1(), F1: -68}
            _ = v76
        }
        if true {
            v4 = (-2147483648 < 94)
            var v77 int = v5
            for fi78 := 0; fi78 < 8; fi78++ {
                _ = len(v6)
                if true {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
                switch (v5 >> ((18446744073709551615) & 63)) {
                case 60:
                    var v79 uint8 = 77
                    _ = v79
                case 22:
                    var v80 uint8 = (v16 - 2)
                    _ = v80
                default:
                    var v81 uint64 = 4294967295
                    _ = v81
                }
                _ = fi78
            }
            _ = v77
        }
        if true {
            for fi82 := 0; fi82 < 12; fi82++ {
                v4 = true
                var v83 int32 = (v1 - 1)
                var v84 float32 = 1.5
                if (v3 <= uint(v16)) {
                    var v85 int = v5
                    _ = v85
                } else {
                    var v86 int32 = v83
                    if v4 {
                        if len(v6) > 0 {
                            v6[0] = "a"
                        }
                        var v87 int16 = 0
                        _ = v87
                    }
                    _ = v86
                }
                _ = fi82
                _ = v83
                _ = v84
            }
            var v88 int8 = -1
            _ = v88
        }
        _ = fi75
    }
    switch "hello" {
    case "foo bar":
        var v89 int16 = -1
        _ = v89
    case "\n\t":
        switch (v5 % ((-9223372036854775808) | 1)) {
        case 1:
            v90 := GenS17{F0: v5}
            _ = v90
        case -1:
            var v91 int32 = v1
            _ = v91
        default:
            for fi92 := 0; fi92 < 1; fi92++ {
                v1 = -2147483648
                if len(v6) > 0 {
                    v6[0] = ("\n\t" + ("foo bar" + (("" + "") + ((("" + "a") + (("" + "hello") + ("\n\t" + ("foo bar" + "foo bar")))) + "a"))))
                }
                _ = fi92
            }
        }
    }
    if false {
        var v93 float64 = float64(0.1)
        _ = v93
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
