package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 int; F1 int16; F2 uint16 }
type GenS4 struct { F0 int64; F1 bool }
type GenS5 struct { F0 uint32; F1 int8 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int16; F1 int; F2 int64 }
type GenS8 struct { F0 uint; F1 string }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int32; F1 uint16; F2 string }
type GenS11 struct { F0 uint; F1 uint64; F2 bool }
type GenS12 struct { F0 uint8 }

func main() {
    var v1 int64 = 0
    v2 := []uint32{}
    v3 := []string{"\n\t", ("a" + "")}
    var v4 float64 = (1e10 / 1.0)
    var v5 uint32 = 983
    switch ("a" + "") {
    case "a":
        _ = len(v3)
    }
    v2 = append(v2, v5)
    v4 = float64(float32(v4))
    for fi6 := 0; fi6 < 7; fi6++ {
        if true {
            v7 := GenS1{F0: 98}
            fmt.Println(fi6)
            _ = v7
        } else {
            if false {
                v8 := GenS2{F0: "foo bar", F1: 229}
                if true {
                    var v9 int64 = (v1 << ((675) & 63))
                    v4 = (((v4 / ((v4 + v4) * float64(1e10))) - v4) / v4)
                    for fi10 := 0; fi10 < 11; fi10++ {
                        var v11 uint16 = 575
                        if true {
                            v4 += 1e10
                        } else {
                            switch 442 {
                            case -503:
                                fmt.Println(v9)
                            default:
                                var v12 int64 = -267
                                _ = v12
                            }
                        }
                        switch ("a" + "hello") {
                        case "foo bar":
                            var v13 uint64 = 600
                            v5 += (^(v5))
                            if false {
                                if false {
                                    if false {
                                        if (true && true) {
                                            v14 := GenS2{F0: "a", F1: uint8(v13)}
                                            switch ("\n\t" + "") {
                                            case "hello":
                                                _ = len(v3)
                                            case "foo bar":
                                                var v15 float32 = 1.5
                                                _ = v15
                                            }
                                            _ = v14
                                        } else {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            v16 := GenS3{F0: fi6, F1: 98, F2: 70}
                                            v5 = (v5 - (v5 % (((v5 + 42)) | 1)))
                                            _ = v16
                                        }
                                        v17 := GenS4{F0: (v1 >> ((18446744073709551615) & 63)), F1: (!true)}
                                        var v18 int32 = -1
                                        _ = v17
                                        _ = v18
                                    } else {
                                        var v19 int16 = 32767
                                        _ = v19
                                    }
                                    if false {
                                        v20 := GenS5{F0: (v5 << ((65535) & 63)), F1: 0}
                                        switch ("\n\t" + "") {
                                        case "a":
                                            var v21 bool = (v11 > v11)
                                            _ = v21
                                        default:
                                            if len(v3) > 0 {
                                                v3[0] = ""
                                            }
                                        }
                                        _ = v20
                                    } else {
                                        var v22 uint16 = 42
                                        var v23 uint32 = v5
                                        _ = v22
                                        _ = v23
                                    }
                                    v13 = v13
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    var v24 string = "foo bar"
                                    _ = v24
                                } else {
                                    var v25 int32 = 2147483647
                                    if (!false) {
                                        var v26 int32 = (^(v25))
                                        _ = v26
                                    }
                                    _ = v25
                                }
                            } else {
                                var v27 int64 = 0
                                fmt.Println(v5, v1, fi10)
                                switch (fi6 * fi6) {
                                case 32768:
                                    if (221 < 0) {
                                        if false {
                                            if true {
                                                var v28 int = (fi10 + (fi10 / ((304) | 1)))
                                                v3 = append(v3, (("" + (("foo bar" + "") + "hello")) + "hello"))
                                                fmt.Println(v1)
                                                fmt.Println(fi10, v1, fi10)
                                                var v29 int8 = -128
                                                _ = v28
                                                _ = v29
                                            }
                                        }
                                        v5 -= (v5 * (v5 / (((v5 / ((42) | 1))) | 1)))
                                        var v30 uint64 = v13
                                        _ = v30
                                    }
                                default:
                                    v31 := GenS3{F0: -27, F1: int16(v5), F2: v11}
                                    _ = v31
                                }
                                _ = v27
                            }
                            _ = v13
                        case "hello":
                            if true {
                                v32 := GenS1{F0: 128}
                                _ = v32
                            } else {
                                var v33 int16 = 37
                                if len(v2) > 0 {
                                    v2[0] = (v5 & v5)
                                }
                                if false {
                                    var v34 uint16 = v11
                                    _ = v34
                                }
                                _ = v33
                            }
                        case "a":
                            v3 = append(v3, ((("hello" + "") + "hello") + "foo bar"))
                            var v35 float32 = -2.25
                            _ = v35
                        default:
                            if false {
                                fmt.Println(fi10, v4)
                                if false {
                                    var v36 float32 = -2.25
                                    var v37 int8 = 3
                                    v2 = append(v2, (^(v5)))
                                    var v38 int16 = 32767
                                    _ = v36
                                    _ = v37
                                    _ = v38
                                }
                                switch (fi6 & fi10) {
                                case 127:
                                    if true {
                                        var v39 uint16 = (v11 % ((v11) | 1))
                                        v39 += v39
                                        v5 += (v5 >> ((661) & 63))
                                        _ = v39
                                    }
                                case 898:
                                    var v40 bool = true
                                    v9 -= int64(fi6)
                                    _ = v40
                                default:
                                    if false {
                                        var v41 float64 = v4
                                        if false {
                                            fmt.Println(v41, v1)
                                        } else {
                                            v9 += v9
                                            v42 := GenS4{F0: -9223372036854775808, F1: false}
                                            _ = v42
                                        }
                                        _ = v41
                                    } else {
                                        if false {
                                            if len(v2) > 0 {
                                                v2[0] = 0
                                            }
                                            v2 = append(v2, v5)
                                            if ((32767 <= 352) || (790 > 2)) {
                                                if (127 != 61) {
                                                    v43 := GenS6{F0: -941}
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        switch fi6 {
                                                        case 1:
                                                            var v44 int16 = 1000
                                                            _ = v44
                                                        default:
                                                            if (0 >= 255) {
                                                                var v45 int64 = v1
                                                                if (true && (uint8(v11) > 23)) {
                                                                    if (!false) {
                                                                        if false {
                                                                            switch 767 {
                                                                            case -67:
                                                                                v46 := GenS6{F0: -116}
                                                                                _ = v46
                                                                            case 128:
                                                                                var v47 int = fi6
                                                                                _ = v47
                                                                            case 4294967295:
                                                                                v48 := GenS4{F0: v9, F1: (false && (-1 == 1))}
                                                                                _ = v48
                                                                            default:
                                                                                if true {
                                                                                    switch (("a" + (("hello" + ("a" + "foo bar")) + "foo bar")) + ("hello" + ("foo bar" + "a"))) {
                                                                                    case "foo bar":
                                                                                        v4 = (v4 / 1e10)
                                                                                        switch "foo bar" {
                                                                                        case "\n\t":
                                                                                            if len(v2) > 0 {
                                                                                                v2[0] = v5
                                                                                            }
                                                                                        case "hello":
                                                                                            v5 = (v5 << ((18446744073709551615) & 63))
                                                                                        }
                                                                                    case "hello":
                                                                                        v49 := GenS1{F0: 128}
                                                                                        _ = v49
                                                                                    }
                                                                                    var v50 string = ("hello" + "hello")
                                                                                    var v51 float32 = (-2.25 + ((float32(-2.25) + 1.5) / 3.14159))
                                                                                    _ = v50
                                                                                    _ = v51
                                                                                }
                                                                            }
                                                                        } else {
                                                                            v5 += 548
                                                                        }
                                                                        v52 := GenS4{F0: (-(v45)), F1: true}
                                                                        _ = v52
                                                                    }
                                                                    var v53 bool = (true || (!(true && false)))
                                                                    v9 -= (-(v1))
                                                                    _ = v53
                                                                }
                                                                _ = v45
                                                            } else {
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                v54 := GenS7{F0: 1000, F1: fi6, F2: (v1 >> ((42) & 63))}
                                                                _ = v54
                                                            }
                                                        }
                                                        v11 = v11
                                                        v4 += float64((1.5 - 0.1))
                                                        var v55 uint16 = 163
                                                        v11 = (v55 % (((v11 + 2)) | 1))
                                                        _ = len(v3)
                                                        if true {
                                                            v11 = v55
                                                            switch (fi10 ^ (fi10 | (fi10 | fi6))) {
                                                            case 32768:
                                                                switch "a" {
                                                                case "a":
                                                                    switch "" {
                                                                    case "\n\t":
                                                                        if (true || true) {
                                                                            v56 := GenS8{F0: 2, F1: "foo bar"}
                                                                            v57 := GenS5{F0: v5, F1: 83}
                                                                            _ = v56
                                                                            _ = v57
                                                                        }
                                                                    case "foo bar":
                                                                        v1 = (-(v1))
                                                                    }
                                                                case "hello":
                                                                    if (!(!true)) {
                                                                        var v58 uint = 1
                                                                        _ = v58
                                                                    } else {
                                                                        v1 += 9223372036854775807
                                                                        var v59 int32 = -2147483648
                                                                        v9 -= (v9 << ((42) & 63))
                                                                        fmt.Println(v4, v5, v55)
                                                                        _ = len(v3)
                                                                        _ = v59
                                                                    }
                                                                }
                                                            case -302:
                                                                if (!(!(590 == int32(v55)))) {
                                                                    fmt.Println(fi6)
                                                                    v2 = append(v2, v5)
                                                                    v11 = (v55 + (v11 >> ((1) & 63)))
                                                                    v60 := GenS9{F0: 18446744073709551615}
                                                                    _ = v60
                                                                }
                                                            }
                                                            v55 += v11
                                                        } else {
                                                            fmt.Println(fi10, v55, v55)
                                                            fmt.Println(v11, v55, v55)
                                                            v5 += v5
                                                        }
                                                        var v61 uint8 = 255
                                                        _ = v55
                                                        _ = v61
                                                    case "foo bar":
                                                        var v62 bool = (872 != 2)
                                                        fmt.Println(v1)
                                                        switch "a" {
                                                        case "":
                                                            v4 -= (v4 * (((v4 + (v4 + v4)) * ((3.14159 + ((v4 / 1.5) + v4)) + (0.1 * -2.25))) * v4))
                                                        case "foo bar":
                                                            fmt.Println(v1, v62, fi6)
                                                        }
                                                        var v63 int = fi10
                                                        if (-681 != (v1 / (((-(v1))) | 1))) {
                                                            v2 = append(v2, v5)
                                                            switch "a" {
                                                            case "a":
                                                                if (1e10 != float64((3.14159 - (float32(v4) * (((1.5 * 1e10) * 3.14159) / float32(v4)))))) {
                                                                    v62 = ((65535 >= 2) && (42 >= uint8(v5)))
                                                                    var v64 int32 = 732
                                                                    v62 = (0.0 < (3.14159 / -2.25))
                                                                    v63 += fi6
                                                                    _ = v64
                                                                }
                                                            case "foo bar":
                                                                var v65 uint = 2
                                                                _ = v65
                                                            case "hello":
                                                                v66 := GenS1{F0: 153}
                                                                _ = v66
                                                            default:
                                                                if false {
                                                                    var v67 string = "hello"
                                                                    var v68 uint16 = (^(v11))
                                                                    _ = v67
                                                                    _ = v68
                                                                } else {
                                                                    v4 += v4
                                                                    if v62 {
                                                                        var v69 bool = (!(1 > 586))
                                                                        if (v62 && true) {
                                                                            v70 := GenS4{F0: v1, F1: true}
                                                                            _ = v70
                                                                        } else {
                                                                            v71 := GenS10{F0: 833, F1: (^(v11)), F2: ("\n\t" + (("" + "a") + ("a" + "\n\t")))}
                                                                            _ = v71
                                                                        }
                                                                        _ = v69
                                                                    } else {
                                                                        v1 = (v1 ^ (-(v1)))
                                                                        var v72 uint32 = 2
                                                                        if (0 < 2) {
                                                                            var v73 int64 = (v1 ^ -919)
                                                                            _ = len(v2)
                                                                            fmt.Println(fi10, fi10)
                                                                            _ = v73
                                                                        }
                                                                        _ = v72
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        _ = v62
                                                        _ = v63
                                                    }
                                                    var v74 uint16 = (^(v11))
                                                    _ = v43
                                                    _ = v74
                                                } else {
                                                    v75 := GenS11{F0: 0, F1: 292, F2: true}
                                                    fmt.Println(fi10, v11)
                                                    var v76 uint8 = 2
                                                    _ = v75
                                                    _ = v76
                                                }
                                                var v77 string = "hello"
                                                _ = v77
                                            } else {
                                                var v78 string = ""
                                                _ = v78
                                            }
                                        } else {
                                            v1 += (^(v1))
                                            v79 := GenS2{F0: "foo bar", F1: 1}
                                            if (!true) {
                                                if ((23 == 47) && (-306 >= 42)) {
                                                    switch "foo bar" {
                                                    case "":
                                                        fmt.Println(v11, v9, fi10)
                                                    case "\n\t":
                                                        if true {
                                                            if true {
                                                                v1 = v9
                                                                fmt.Println(v11, v4, v11)
                                                            }
                                                        } else {
                                                            v4 -= ((v4 / v4) - 1.5)
                                                            var v80 int8 = int8(v5)
                                                            _ = v80
                                                        }
                                                        _ = len(v3)
                                                        v9 = v9
                                                        if true {
                                                            v81 := GenS1{F0: 255}
                                                            v82 := GenS12{F0: 255}
                                                            fmt.Println(v1)
                                                            var v83 float64 = v4
                                                            v1 = v1
                                                            if len(v2) > 0 {
                                                                v2[0] = v5
                                                            }
                                                            _ = v81
                                                            _ = v82
                                                            _ = v83
                                                        }
                                                    }
                                                    var v84 uint32 = 4294967295
                                                    _ = v84
                                                }
                                            } else {
                                                var v85 int8 = 28
                                                var v86 int = (fi10 | int(v85))
                                                switch ((("foo bar" + "foo bar") + "") + "hello") {
                                                case "\n\t":
                                                    if ((v85 * (v85 / ((v85) | 1))) == -61) {
                                                        _ = len(v2)
                                                        switch ("hello" + "foo bar") {
                                                        case "":
                                                            var v87 float32 = 0.1
                                                            _ = v87
                                                        case "foo bar":
                                                            var v88 uint16 = 2
                                                            _ = v88
                                                        case "\n\t":
                                                            fmt.Println(v1, v11, v11)
                                                        default:
                                                            var v89 float32 = (0.1 + ((float32(((0.1 + (v4 + v4)) * 0.0)) * 0.1) + (3.14159 * 1e10)))
                                                            _ = v89
                                                        }
                                                        if (!false) {
                                                            var v90 int32 = 0
                                                            fmt.Println(fi6, v5, v4)
                                                            _ = v90
                                                        } else {
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        }
                                                    } else {
                                                        var v91 int32 = 1
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        _ = v91
                                                    }
                                                case "hello":
                                                    v92 := GenS8{F0: 42, F1: "a"}
                                                    _ = v92
                                                }
                                                _ = v85
                                                _ = v86
                                            }
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            _ = v79
                                        }
                                    }
                                }
                            } else {
                                if false {
                                    v93 := GenS2{F0: "foo bar", F1: 69}
                                    v4 += 1e10
                                    _ = v93
                                } else {
                                    v94 := GenS10{F0: 908, F1: 1, F2: "a"}
                                    fmt.Println(v5, fi6)
                                    _ = v94
                                }
                            }
                        }
                        _ = fi10
                        _ = v11
                    }
                    _ = v9
                } else {
                    v4 -= v4
                }
                v95 := GenS1{F0: 17}
                _ = v8
                _ = v95
            }
        }
        _ = fi6
    }
    var v96 uint64 = 111
    v5 = v5
    fmt.Println(v4)
    for fi97 := 0; fi97 < 17; fi97++ {
        fmt.Println(v1)
        switch -9223372036854775808 {
        case 9223372036854775807:
            for fi98 := 0; fi98 < 18; fi98++ {
                v99 := GenS7{F0: 1000, F1: (fi97 >> ((453) & 63)), F2: (v1 / ((0) | 1))}
                _ = fi98
                _ = v99
            }
        }
        var v100 int = -2147483648
        _ = fi97
        _ = v100
    }
    var v101 uint16 = 255
    v101 = (v101 / (((v101 + v101)) | 1))
    v102 := GenS5{F0: (v5 >> ((0) & 63)), F1: 39}
    fmt.Println("v1:", v1)
    fmt.Println("v101:", v101)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v101
    _ = v102
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v96
}
