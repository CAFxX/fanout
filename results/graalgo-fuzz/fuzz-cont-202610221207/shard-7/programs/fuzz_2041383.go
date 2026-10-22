package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 uint8; F1 int }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 bool; F1 int32 }
type GenS7 struct { F0 uint8; F1 int16; F2 uint64 }
type GenS8 struct { F0 uint8; F1 uint16 }
type GenS9 struct { F0 int16; F1 uint16 }
func genFn0(p0 int16) int64 {
    for fi5 := 0; fi5 < 13; fi5++ {
        var v6 bool = (737 == 0)
        var v7 uint = 1
        _ = fi5
        _ = v6
        _ = v7
    }
    fmt.Println(p0)
    for fi8 := 0; fi8 < 2; fi8++ {
        v9 := GenS1{F0: -128, F1: 115}
        fmt.Println(p0, fi8)
        _ = fi8
        _ = v9
    }
    var v10 uint32 = 965
    _ = p0
    _ = v10
    return 0
}
func genFn1() uint16 {
    var v11 uint16 = 0
    v12 := GenS2{F0: 41, F1: 1}
    fmt.Println(v11)
    _ = v11
    _ = v12
    return 42
}

func main() {
    var v1 int = -9223372036854775808
    v2 := make(map[int]bool)
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]string)
    if false {
        if false {
            fmt.Println(v1)
            for fi13 := 0; fi13 < 18; fi13++ {
                v14 := []int{}
                for fi15 := 0; fi15 < 1; fi15++ {
                    v16 := GenS2{F0: -128, F1: 42}
                    _ = fi15
                    _ = v16
                }
                _ = fi13
                _ = v14
            }
            v17 := []int{9223372036854775807, v1}
            _ = v17
        }
        var v18 bool = true
        v18 = (!((v1 * v1) >= (v1 - v1)))
        _ = v18
    } else {
        fmt.Println(v1)
    }
    v1 = 32767
    v19 := GenS2{F0: -128, F1: -1}
    var v20 float64 = -2.25
    if (!false) {
        for fi21 := 0; fi21 < 8; fi21++ {
            for fi22 := 0; fi22 < 8; fi22++ {
                var v23 int32 = -2147483648
                _ = fi22
                _ = v23
            }
            _ = fi21
        }
    }
    v1 = v1
    v24 := []int{v1}
    fmt.Println(v20, v20)
    for fi25 := 0; fi25 < 8; fi25++ {
        v1 -= fi25
        fmt.Println(v1, fi25)
        _ = fi25
    }
    v20 = 3.14159
    v1 += (^(v1))
    if len(v24) >= 2 {
        v24 = v24[1:len(v24)]
    }
    v1 = (v1 & v1)
    v26 := GenS1{F0: 100, F1: -1}
    if true {
        switch (-(v1)) {
        case 4294967295:
            for fi27 := 0; fi27 < 1; fi27++ {
                if (255 < 1) {
                    v28 := GenS3{F0: (1 != 0), F1: fi27}
                    v1 -= (v1 & v1)
                    _ = v28
                } else {
                    for fi29 := 0; fi29 < 1; fi29++ {
                        var v30 int32 = -1
                        _ = fi29
                        _ = v30
                    }
                }
                fmt.Println(v20)
                if false {
                    var v31 uint16 = 42
                    if (!true) {
                        v31 = (v31 | (v31 / ((v31) | 1)))
                    } else {
                        v32 := GenS4{F0: 1, F1: v1}
                        v33 := v2[(v1 - v1)]
                        _ = v32
                        _ = v33
                    }
                    _ = v31
                }
                _ = fi27
            }
        default:
            for fi34 := 0; fi34 < 7; fi34++ {
                _ = len(v24)
                for fi35 := 0; fi35 < 8; fi35++ {
                    if (("foo bar" > "hello") && (!((126 != 127) || (42 != 42)))) {
                        if ((897 != 18446744073709551615) || false) {
                            v1 = (fi34 >> ((255) & 63))
                        } else {
                            if (2 == 1) {
                                var v36 int64 = -1
                                _ = v36
                            } else {
                                v37 := GenS5{F0: 318}
                                var v38 uint32 = 42
                                fmt.Println(v1, fi34)
                                var v39 uint = 2
                                _ = v37
                                _ = v38
                                _ = v39
                            }
                            if false {
                                switch "foo bar" {
                                case "hello":
                                    fmt.Println(v1)
                                }
                                var v40 uint16 = 0
                                _ = v40
                            }
                        }
                        var v41 float32 = ((-2.25 + -2.25) * float32((v20 + (-2.25 + ((float64(0.1) / 3.14159) + v20)))))
                        var v42 int32 = -654
                        _ = v41
                        _ = v42
                    } else {
                        v24 = append(v24, fi35)
                        v1 = -1
                    }
                    v20 += 1e10
                    _ = len(v24)
                    _ = fi35
                }
                _ = fi34
            }
        }
        var v43 int64 = 42
        _ = v43
    }
    _ = len(v24)
    for fi44 := 0; fi44 < 1; fi44++ {
        switch (-(fi44)) {
        case 4294967295:
            for fi45 := 0; fi45 < 12; fi45++ {
                var v46 uint16 = 438
                var v47 int8 = 93
                _ = fi45
                _ = v46
                _ = v47
            }
        default:
            switch (("a" + "\n\t") + (("foo bar" + "") + ("hello" + "hello"))) {
            case "foo bar":
                if false {
                    switch ("foo bar" + "\n\t") {
                    case "\n\t":
                        if (57 > 65535) {
                            if false {
                                switch "" {
                                case "foo bar":
                                    var v48 float64 = (0.1 * v20)
                                    _ = v48
                                case "a":
                                    fmt.Println(fi44)
                                case "\n\t":
                                    for fi49 := 0; fi49 < 17; fi49++ {
                                        var v50 int8 = 42
                                        _ = fi49
                                        _ = v50
                                    }
                                default:
                                    v20 = ((3.14159 / -2.25) * (v20 * 0.1))
                                }
                                var v51 string = "foo bar"
                                _ = v51
                            } else {
                                if (-533 != -637) {
                                    for fi52 := 0; fi52 < 8; fi52++ {
                                        var v53 int16 = 1
                                        _ = fi52
                                        _ = v53
                                    }
                                    v24 = append(v24, fi44)
                                    _ = len(v3)
                                } else {
                                    if (((true && ((18446744073709551615 <= 18446744073709551615) && (151 < v1))) || (float64(((3.14159 * float32(v20)) / 1.0)) <= float64(0.1))) || true) {
                                        if len(v24) >= 2 {
                                            v24 = v24[1:len(v24)]
                                        }
                                        v54 := v3[2147483648]
                                        var v55 float64 = ((v20 + 0.1) + 1e10)
                                        _ = v54
                                        _ = v55
                                    }
                                    _ = len(v4)
                                }
                            }
                            if (1 == 607) {
                                var v56 int16 = -32768
                                for fi57 := 0; fi57 < 19; fi57++ {
                                    var v58 int16 = (v56 * (v56 & (v56 * (v56 << ((0) & 63)))))
                                    v20 = 3.14159
                                    var v59 int = -940
                                    fmt.Println(fi57)
                                    var v60 int8 = -128
                                    if true {
                                        switch (("foo bar" + "\n\t") + "foo bar") {
                                        case "":
                                            var v61 int = -627
                                            _ = v61
                                        case "\n\t":
                                            fmt.Println(v20, v58, v60)
                                        case "foo bar":
                                            if true {
                                                v62 := GenS1{F0: (v60 | (v60 % (((v60 % ((v60) | 1))) | 1))), F1: 880}
                                                var v63 uint16 = 341
                                                var v64 uint64 = 2
                                                _ = v62
                                                _ = v63
                                                _ = v64
                                            }
                                        default:
                                            if ((^(v58)) > v58) {
                                                if true {
                                                    var v65 int8 = v60
                                                    _ = v65
                                                } else {
                                                    var v66 bool = (true || true)
                                                    fmt.Println(v58, fi57, v59)
                                                    _ = v66
                                                }
                                            } else {
                                                v58 += v56
                                            }
                                            switch fi44 {
                                            case 128:
                                                fmt.Println(v58, v1, v1)
                                            case 0:
                                                if (2 >= 18446744073709551615) {
                                                    var v67 int64 = 42
                                                    var v68 int = v1
                                                    if (1 != 18446744073709551615) {
                                                        v67 = (-(v67))
                                                        v69 := GenS1{F0: int8(v67), F1: 0}
                                                        fmt.Println(v60)
                                                        _ = v69
                                                    } else {
                                                        switch (v1 | v1) {
                                                        case -2147483648:
                                                            switch "foo bar" {
                                                            case "hello":
                                                                var v70 uint64 = 115
                                                                _ = v70
                                                            case "":
                                                                var v71 uint = 986
                                                                _ = v71
                                                            }
                                                        default:
                                                            v59 = 2147483647
                                                        }
                                                        v56 += v56
                                                        var v72 float64 = v20
                                                        _ = v72
                                                    }
                                                    _ = v67
                                                    _ = v68
                                                }
                                            default:
                                                fmt.Println(v20, v56, fi57)
                                            }
                                            var v73 uint16 = genFn1()
                                            _ = v73
                                        }
                                        var v74 int32 = -2147483648
                                        _ = v74
                                    } else {
                                        v56 += v56
                                        if (!false) {
                                            v60 += (-(v60))
                                            switch fi44 {
                                            case 127:
                                                v75 := v4[(("foo bar" + "\n\t") + "foo bar")]
                                                _ = v75
                                            default:
                                                v60 -= v60
                                            }
                                            v24 = append(v24, -135)
                                        }
                                    }
                                    var v76 int32 = -1
                                    var v77 uint64 = 392
                                    _ = fi57
                                    _ = v58
                                    _ = v59
                                    _ = v60
                                    _ = v76
                                    _ = v77
                                }
                                _ = v56
                            } else {
                                fmt.Println(fi44, v1)
                                v1 += fi44
                            }
                        } else {
                            v1 = 32768
                            v1 += fi44
                        }
                    }
                    v2[(fi44 << ((65535) & 63))] = (0.0 < (((3.14159 + 3.14159) * 1e10) / (float32(-2.25) / (-2.25 / 3.14159))))
                    var v78 int32 = -1
                    _ = v78
                } else {
                    v24 = append(v24, (fi44 - (v1 / (((v1 - (^(v1)))) | 1))))
                    switch "hello" {
                    case "foo bar":
                        if ((0.1 < v20) || (2147483647 > 2147483648)) {
                            v79 := GenS6{F0: true, F1: 1}
                            var v80 int = fi44
                            _ = v79
                            _ = v80
                        } else {
                            var v81 int64 = 42
                            v24 = append(v24, (fi44 << ((42) & 63)))
                            _ = v81
                        }
                    default:
                        v20 = (v20 + (float64(3.14159) * 3.14159))
                    }
                    if (1 < 42) {
                        for fi82 := 0; fi82 < 19; fi82++ {
                            if true {
                                if (!true) {
                                    v1 -= 2147483647
                                    v20 -= v20
                                }
                            } else {
                                _ = len(v24)
                                v20 -= ((float64(0.1) / (v20 * v20)) + (v20 / v20))
                            }
                            _ = fi82
                        }
                        if len(v24) > 0 {
                            v24[0] = (v1 & 32768)
                        }
                    }
                }
            }
        }
        var v83 string = "hello"
        _ = fi44
        _ = v83
    }
    if true {
        v20 = (v20 * (v20 / (3.14159 + float64(1.5))))
    } else {
        if (true || false) {
            if false {
                delete(v3, 2147483647)
                var v84 uint = 255
                _ = v84
            } else {
                v85 := GenS7{F0: 42, F1: -1, F2: 20}
                _ = v85
            }
            var v86 uint8 = 195
            v87 := GenS8{F0: v86, F1: 32767}
            _ = v86
            _ = v87
        } else {
            v88 := GenS9{F0: 0, F1: 32767}
            var v89 int16 = 0
            if true {
                var v90 float32 = 1e10
                _ = v90
            } else {
                v1 = (v1 + (^(v1)))
                for fi91 := 0; fi91 < 17; fi91++ {
                    var v92 uint32 = 65535
                    v89 -= (-(v89))
                    v89 = (v89 & (v89 << ((255) & 63)))
                    _ = fi91
                    _ = v92
                }
                if true {
                    v93 := GenS7{F0: 208, F1: v89, F2: 0}
                    for fi94 := 0; fi94 < 12; fi94++ {
                        var v95 int32 = 985
                        _ = fi94
                        _ = v95
                    }
                    switch ("\n\t" + (("hello" + "") + "")) {
                    case "hello":
                        fmt.Println(v89, v1)
                    case "a":
                        v20 = float64(1e10)
                    case "foo bar":
                        for fi96 := 0; fi96 < 13; fi96++ {
                            var v97 uint16 = 552
                            for fi98 := 0; fi98 < 12; fi98++ {
                                var v99 int8 = -128
                                _ = fi98
                                _ = v99
                            }
                            _ = fi96
                            _ = v97
                        }
                    default:
                        v100 := GenS2{F0: 1, F1: -2147483648}
                        _ = v100
                    }
                    _ = v93
                } else {
                    v101 := GenS3{F0: true, F1: (v1 ^ (^(v1)))}
                    for fi102 := 0; fi102 < 3; fi102++ {
                        for fi103 := 0; fi103 < 9; fi103++ {
                            if (18446744073709551615 > 2) {
                                switch ("" + "") {
                                case "\n\t":
                                    switch 1 {
                                    case 42:
                                        var v104 bool = (42 >= 201)
                                        _ = v104
                                    default:
                                        v20 += float64((float32(1.5) + 3.14159))
                                    }
                                case "a":
                                    _ = len(v24)
                                default:
                                    v105 := GenS5{F0: 65535}
                                    _ = v105
                                }
                                if (759 <= 628) {
                                    switch ((((("a" + ("a" + ("\n\t" + "\n\t"))) + "\n\t") + "foo bar") + ("hello" + "a")) + "foo bar") {
                                    case "foo bar":
                                        v89 = (v89 - v89)
                                        v1 = (-(v1))
                                        if len(v24) > 0 {
                                            v24[0] = 154
                                        }
                                    default:
                                        v89 = -368
                                    }
                                }
                            }
                            _ = fi103
                        }
                        _ = len(v4)
                        var v106 uint16 = 255
                        _ = fi102
                        _ = v106
                    }
                    _ = v101
                }
            }
            _ = v88
            _ = v89
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v26
    _ = v3
    _ = v4
}
