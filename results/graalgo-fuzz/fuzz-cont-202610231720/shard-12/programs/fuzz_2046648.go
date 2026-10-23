package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 bool; F1 uint32; F2 uint64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int64; F1 int32 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint8; F1 bool }
type GenS9 struct { F0 uint16; F1 int64; F2 int }
type GenS10 struct { F0 bool; F1 uint64; F2 int16 }
type GenS11 struct { F0 int32; F1 int }
type GenS12 struct { F0 int64; F1 uint8; F2 int8 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 int; F1 int64 }
type GenS15 struct { F0 string; F1 string; F2 bool }
type GenS16 struct { F0 uint8; F1 uint16; F2 bool }
func genFn0(p0 int64) uint16 {
    var v7 float64 = (0.1 * 3.14159)
    p0 -= (p0 ^ (-(p0)))
    p0 = (p0 & 9223372036854775807)
    v8 := []int{42, -2147483648}
    _ = p0
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 uint32 = 2147483647
    var v2 uint8 = 2
    var v3 int8 = 127
    var v4 uint8 = 112
    v5 := []uint32{42, v1}
    v6 := make(map[uint32]uint32)
    if ((18446744073709551615 > 2) || true) {
        for fi9 := 0; fi9 < 1; fi9++ {
            var v10 int = (fi9 / ((int(v1)) | 1))
            var v11 string = "a"
            _ = fi9
            _ = v10
            _ = v11
        }
        if false {
            if (42 != 42) {
                switch ("foo bar" + (("hello" + ("" + "a")) + "a")) {
                case "foo bar":
                    var v12 int8 = (v3 % ((int8(v1)) | 1))
                    _ = v12
                }
            }
            v13 := v6[(v1 >> ((230) & 63))]
            _ = v13
        } else {
            fmt.Println(v1, v3, v4)
            if true {
                v3 -= (v3 >> ((0) & 63))
            }
            if false {
                v14 := GenS1{F0: (^(v3)), F1: ("a" + "foo bar")}
                _ = len(v6)
                v4 = v2
                _ = v14
            } else {
                for fi15 := 0; fi15 < 8; fi15++ {
                    if true {
                        var v16 uint32 = (v1 << ((65535) & 63))
                        v2 += v2
                        if ((v4 * v2) > 0) {
                            for fi17 := 0; fi17 < 18; fi17++ {
                                _ = len(v6)
                                if false {
                                    v2 = v4
                                } else {
                                    v16 = v1
                                    v18 := GenS1{F0: -128, F1: "foo bar"}
                                    v19 := GenS1{F0: (v3 - -27), F1: ("a" + ("" + (("\n\t" + "hello") + ((("\n\t" + "foo bar") + "hello") + ("a" + ("" + "\n\t"))))))}
                                    _ = v18
                                    _ = v19
                                }
                                if (0.1 == (1e10 / 1.0)) {
                                    if true {
                                        v20 := v6[v1]
                                        if false {
                                            switch (fi15 / ((fi15) | 1)) {
                                            case -1:
                                                if false {
                                                    switch fi17 {
                                                    case -9223372036854775808:
                                                        v4 = (v4 >> ((1) & 63))
                                                    case 32768:
                                                        fmt.Println(v1)
                                                    case 1:
                                                        v21 := GenS1{F0: -128, F1: "a"}
                                                        _ = v21
                                                    default:
                                                        _ = len(v6)
                                                    }
                                                    var v22 uint8 = 2
                                                    var v23 int16 = -1
                                                    v16 -= 4294967295
                                                    _ = v22
                                                    _ = v23
                                                }
                                            case 974:
                                                if (42 <= 226) {
                                                    v3 = (v3 % ((1) | 1))
                                                } else {
                                                    switch fi15 {
                                                    case 9223372036854775807:
                                                        v24 := GenS1{F0: 42, F1: ("" + "\n\t")}
                                                        _ = v24
                                                    case -2147483648:
                                                        var v25 bool = false
                                                        _ = v25
                                                    case -252:
                                                        var v26 uint32 = v16
                                                        _ = v26
                                                    default:
                                                        var v27 bool = (686 > -836)
                                                        _ = v27
                                                    }
                                                    if true {
                                                        var v28 int64 = 0
                                                        fmt.Println(v1)
                                                        if (false || (!(false || (false || true)))) {
                                                            v1 = 200
                                                            switch fi17 {
                                                            case 32768:
                                                                fmt.Println(v2, v2)
                                                                switch 0 {
                                                                case 4294967295:
                                                                    v29 := GenS1{F0: -1, F1: "a"}
                                                                    _ = v29
                                                                case 41:
                                                                    v3 = v3
                                                                case 32767:
                                                                    if len(v5) > 0 {
                                                                        v5[0] = (v16 ^ (v1 / ((v20) | 1)))
                                                                    }
                                                                default:
                                                                    v3 -= v3
                                                                }
                                                                if false {
                                                                    if len(v5) >= 2 {
                                                                        v5 = v5[1:len(v5)]
                                                                    }
                                                                }
                                                            case 1:
                                                                v28 = v28
                                                            default:
                                                                fmt.Println(v3, v3, v16)
                                                            }
                                                        } else {
                                                            if false {
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                            } else {
                                                                switch fi15 {
                                                                case 128:
                                                                    if (((32767 != 32767) && (!((2 != v4) && true))) && true) {
                                                                        if false {
                                                                            var v30 float64 = (((3.14159 - 0.1) / 1e10) / (float64(1.5) + 1e10))
                                                                            _ = v30
                                                                        } else {
                                                                            fmt.Println(v1)
                                                                            if (127 < (v2 | 255)) {
                                                                                v5 = append(v5, v1)
                                                                                var v31 uint64 = 42
                                                                                switch (^(fi15)) {
                                                                                case 2147483647:
                                                                                    v32 := GenS1{F0: v3, F1: ("hello" + "hello")}
                                                                                    _ = v32
                                                                                case 47:
                                                                                    switch fi17 {
                                                                                    case -649:
                                                                                        switch "foo bar" {
                                                                                        case "a":
                                                                                            delete(v6, v20)
                                                                                        case "\n\t":
                                                                                            v20 += v20
                                                                                            v33 := GenS1{F0: (v3 | v3), F1: "a"}
                                                                                            if false {
                                                                                                if true {
                                                                                                    v3 -= (-(v3))
                                                                                                    if false {
                                                                                                        var v34 int = (fi17 / ((fi15) | 1))
                                                                                                        _ = v34
                                                                                                    } else {
                                                                                                        v1 -= 56
                                                                                                        _ = len(v5)
                                                                                                        v35 := GenS2{F0: "foo bar"}
                                                                                                        var v36 int32 = int32(v28)
                                                                                                        v37 := GenS3{F0: (v36 * v36), F1: int(v28)}
                                                                                                        if true {
                                                                                                            v38 := GenS2{F0: "foo bar"}
                                                                                                            var v39 int64 = (v28 >> ((18446744073709551615) & 63))
                                                                                                            if false {
                                                                                                                var v40 uint = 2
                                                                                                                var v41 float64 = 1.5
                                                                                                                fmt.Println(v4, v39, v41)
                                                                                                                _ = v40
                                                                                                                _ = v41
                                                                                                            }
                                                                                                            var v42 uint = uint(v16)
                                                                                                            _ = v38
                                                                                                            _ = v39
                                                                                                            _ = v42
                                                                                                        }
                                                                                                        v16 = v1
                                                                                                        v43 := GenS3{F0: v36, F1: fi17}
                                                                                                        var v44 float32 = (-2.25 + (float32(3.14159) - 1.5))
                                                                                                        _ = v35
                                                                                                        _ = v36
                                                                                                        _ = v37
                                                                                                        _ = v43
                                                                                                        _ = v44
                                                                                                    }
                                                                                                    if true {
                                                                                                        v45 := GenS1{F0: int8(v31), F1: (("a" + "foo bar") + "")}
                                                                                                        var v46 int = (fi17 >> ((33) & 63))
                                                                                                        _ = v45
                                                                                                        _ = v46
                                                                                                    }
                                                                                                    fmt.Println(fi15, v4)
                                                                                                    var v47 uint32 = (v16 + v1)
                                                                                                    _ = v47
                                                                                                } else {
                                                                                                    v2 = 128
                                                                                                }
                                                                                            } else {
                                                                                                v48 := GenS4{F0: ((v20 * (^(v1))) < (v1 % (((v20 + v20)) | 1))), F1: (v16 >> ((18446744073709551615) & 63)), F2: (v31 + (^(v31)))}
                                                                                                if ((!false) && true) {
                                                                                                    if (42 == 406) {
                                                                                                        v20 = v16
                                                                                                    }
                                                                                                    v16 -= (v20 - v20)
                                                                                                } else {
                                                                                                    if (!false) {
                                                                                                        fmt.Println(v16, fi15)
                                                                                                    } else {
                                                                                                        var v49 uint64 = v31
                                                                                                        var v50 int64 = v28
                                                                                                        _ = v49
                                                                                                        _ = v50
                                                                                                    }
                                                                                                    v16 -= (v16 >> ((255) & 63))
                                                                                                    delete(v6, v16)
                                                                                                }
                                                                                                v3 += (v3 << ((426) & 63))
                                                                                                _ = v48
                                                                                            }
                                                                                            _ = v33
                                                                                        case "":
                                                                                            if (65535 != 165) {
                                                                                                v51 := GenS5{F0: 1}
                                                                                                v52 := GenS6{F0: -9223372036854775808, F1: 373}
                                                                                                _ = v51
                                                                                                _ = v52
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        v53 := GenS7{F0: (v4 * (v2 / ((v2) | 1)))}
                                                                                        _ = v53
                                                                                    }
                                                                                default:
                                                                                    if (0 < v4) {
                                                                                        var v54 float64 = 1.5
                                                                                        _ = v54
                                                                                    } else {
                                                                                        var v55 int8 = (v3 ^ v3)
                                                                                        fmt.Println(v20, v4)
                                                                                        var v56 float64 = 3.14159
                                                                                        _ = v55
                                                                                        _ = v56
                                                                                    }
                                                                                }
                                                                                _ = v31
                                                                            } else {
                                                                                if true {
                                                                                    v2 = (v4 << ((42) & 63))
                                                                                    v4 -= v4
                                                                                } else {
                                                                                    switch fi15 {
                                                                                    case 808:
                                                                                        var v57 uint16 = 0
                                                                                        _ = v57
                                                                                    case 127:
                                                                                        var v58 uint16 = 32767
                                                                                        _ = v58
                                                                                    }
                                                                                    v2 += v4
                                                                                    if false {
                                                                                        v16 -= v1
                                                                                    }
                                                                                }
                                                                                switch ("" + "a") {
                                                                                case "foo bar":
                                                                                    fmt.Println(v4, fi15)
                                                                                case "\n\t":
                                                                                    if len(v5) > 0 {
                                                                                        v5[0] = v20
                                                                                    }
                                                                                default:
                                                                                    if false {
                                                                                        if true {
                                                                                            var v59 float64 = 3.14159
                                                                                            var v60 uint8 = 255
                                                                                            v5 = append(v5, (^(v16)))
                                                                                            _ = v59
                                                                                            _ = v60
                                                                                        } else {
                                                                                            var v61 float64 = 3.14159
                                                                                            if false {
                                                                                                var v62 string = "foo bar"
                                                                                                if false {
                                                                                                    fmt.Println(v3, v20)
                                                                                                    var v63 float32 = -2.25
                                                                                                    v3 -= v3
                                                                                                    _ = v63
                                                                                                } else {
                                                                                                    v20 += v16
                                                                                                    v64 := GenS1{F0: (v3 | (v3 + (v3 - v3))), F1: "\n\t"}
                                                                                                    _ = v64
                                                                                                }
                                                                                                var v65 float64 = 0.1
                                                                                                _ = v62
                                                                                                _ = v65
                                                                                            }
                                                                                            var v66 int = (fi15 & (^(fi15)))
                                                                                            v67 := GenS3{F0: 0, F1: (v66 % ((559) | 1))}
                                                                                            _ = v61
                                                                                            _ = v66
                                                                                            _ = v67
                                                                                        }
                                                                                    } else {
                                                                                        v5 = append(v5, (v1 + 4294967295))
                                                                                        switch "" {
                                                                                        case "\n\t":
                                                                                            if false {
                                                                                                v3 += 127
                                                                                                var v68 uint32 = v16
                                                                                                if ("hello" == "") {
                                                                                                    switch 1 {
                                                                                                    case 127:
                                                                                                        switch fi15 {
                                                                                                        case 128:
                                                                                                            v20 -= (v68 ^ (^(v20)))
                                                                                                        case -2147483648:
                                                                                                            var v69 uint64 = 4294967295
                                                                                                            _ = v69
                                                                                                        case 32767:
                                                                                                            var v70 int64 = -1
                                                                                                            _ = v70
                                                                                                        }
                                                                                                    case -2147483648:
                                                                                                        v20 -= (^(v68))
                                                                                                    default:
                                                                                                        v71 := GenS8{F0: v4, F1: false}
                                                                                                        _ = v71
                                                                                                    }
                                                                                                }
                                                                                                _ = len(v5)
                                                                                                var v72 int = -9223372036854775808
                                                                                                _ = v68
                                                                                                _ = v72
                                                                                            } else {
                                                                                                switch (("a" + "") + "a") {
                                                                                                case "\n\t":
                                                                                                    v6[v1] = (v16 + (^(v1)))
                                                                                                case "a":
                                                                                                    if (false && (42 >= 255)) {
                                                                                                        fmt.Println(v20, fi17)
                                                                                                        switch (fi17 & 128) {
                                                                                                        case 42:
                                                                                                            switch "a" {
                                                                                                            case "":
                                                                                                                v4 = 127
                                                                                                            case "hello":
                                                                                                                if len(v5) >= 2 {
                                                                                                                    v5 = v5[1:len(v5)]
                                                                                                                }
                                                                                                            case "\n\t":
                                                                                                                fmt.Println(v4, fi15, v4)
                                                                                                            default:
                                                                                                                if len(v5) >= 2 {
                                                                                                                    v5 = v5[1:len(v5)]
                                                                                                                }
                                                                                                            }
                                                                                                        case 2147483647:
                                                                                                            var v73 uint8 = v2
                                                                                                            _ = v73
                                                                                                        case 1:
                                                                                                            var v74 string = "hello"
                                                                                                            _ = v74
                                                                                                        }
                                                                                                    } else {
                                                                                                        delete(v6, 2147483648)
                                                                                                        var v75 string = "foo bar"
                                                                                                        _ = v75
                                                                                                    }
                                                                                                case "hello":
                                                                                                    var v76 float64 = 0.1
                                                                                                    fmt.Println(v16)
                                                                                                    v77 := GenS2{F0: "\n\t"}
                                                                                                    var v78 int = (fi17 << ((65535) & 63))
                                                                                                    _ = len(v5)
                                                                                                    var v79 uint = 255
                                                                                                    _ = v76
                                                                                                    _ = v77
                                                                                                    _ = v78
                                                                                                    _ = v79
                                                                                                }
                                                                                            }
                                                                                        case "foo bar":
                                                                                            v28 += 9223372036854775807
                                                                                            v80 := GenS3{F0: -1, F1: fi15}
                                                                                            _ = v80
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                            v16 += (v20 - (v20 << ((382) & 63)))
                                                                        }
                                                                        v16 = v20
                                                                        if (false && true) {
                                                                            var v81 int64 = -979
                                                                            v2 -= uint8(v16)
                                                                            _ = v81
                                                                        }
                                                                        fmt.Println(v20)
                                                                    }
                                                                }
                                                                fmt.Println(v20, v16)
                                                            }
                                                        }
                                                        _ = v28
                                                    } else {
                                                        v3 -= (v3 | (v3 * v3))
                                                        if (-646 != -1) {
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                        } else {
                                                            if false {
                                                                v2 = v4
                                                                fmt.Println(v16, v16, v16)
                                                            } else {
                                                                if false {
                                                                    v82 := GenS9{F0: 560, F1: -9223372036854775808, F2: fi15}
                                                                    var v83 float64 = 0.1
                                                                    if true {
                                                                        v20 = 42
                                                                    }
                                                                    var v84 uint16 = 32767
                                                                    _ = v82
                                                                    _ = v83
                                                                    _ = v84
                                                                } else {
                                                                    switch "a" {
                                                                    case "foo bar":
                                                                        v85 := GenS9{F0: 255, F1: 944, F2: fi17}
                                                                        _ = v85
                                                                    default:
                                                                        if (true && (0.1 == 3.14159)) {
                                                                            var v86 uint8 = (v4 - v4)
                                                                            var v87 int = (fi15 / ((4294967295) | 1))
                                                                            fmt.Println(v1, v16, v2)
                                                                            _ = v86
                                                                            _ = v87
                                                                        } else {
                                                                            v4 = (^(v2))
                                                                            switch "" {
                                                                            case "":
                                                                                var v88 uint8 = (^(v4))
                                                                                _ = v88
                                                                            case "foo bar":
                                                                                var v89 int16 = 0
                                                                                _ = v89
                                                                            case "\n\t":
                                                                                if false {
                                                                                    var v90 int32 = 0
                                                                                    switch (fi15 * 244) {
                                                                                    case 32768:
                                                                                        switch (-(fi15)) {
                                                                                        case 508:
                                                                                            if true {
                                                                                                var v91 uint8 = v4
                                                                                                _ = v91
                                                                                            } else {
                                                                                                if (128 > (v4 + v4)) {
                                                                                                    if len(v5) >= 2 {
                                                                                                        v5 = v5[1:len(v5)]
                                                                                                    }
                                                                                                    if false {
                                                                                                        v16 = v20
                                                                                                        if true {
                                                                                                            fmt.Println(v4, v4, v1)
                                                                                                            delete(v6, (v20 / (((v20 % ((603) | 1))) | 1)))
                                                                                                        }
                                                                                                    } else {
                                                                                                        var v92 int32 = v90
                                                                                                        v93 := GenS10{F0: (int(v92) != 128), F1: 1, F2: 32767}
                                                                                                        v6[v16] = 182
                                                                                                        if len(v5) > 0 {
                                                                                                            v5[0] = 2147483648
                                                                                                        }
                                                                                                        v92 -= v90
                                                                                                        _ = v92
                                                                                                        _ = v93
                                                                                                    }
                                                                                                    fmt.Println(fi15, fi17, fi17)
                                                                                                }
                                                                                                switch -450 {
                                                                                                case -878:
                                                                                                    fmt.Println(v90, v90, v2)
                                                                                                default:
                                                                                                    var v94 uint8 = (v2 * v4)
                                                                                                    _ = v94
                                                                                                }
                                                                                            }
                                                                                        case 31:
                                                                                            fmt.Println(fi15, v2)
                                                                                        default:
                                                                                            v95 := GenS7{F0: (v2 + v4)}
                                                                                            _ = v95
                                                                                        }
                                                                                    case 46:
                                                                                        if (true || (v4 >= v4)) {
                                                                                            var v96 uint16 = 90
                                                                                            var v97 string = "a"
                                                                                            _ = v96
                                                                                            _ = v97
                                                                                        } else {
                                                                                            switch ((("hello" + ("" + "\n\t")) + "hello") + "") {
                                                                                            case "hello":
                                                                                                if (582 < 0) {
                                                                                                    v98 := GenS10{F0: (-2.25 <= 0.0), F1: 4294967295, F2: int16(fi17)}
                                                                                                    var v99 uint32 = v20
                                                                                                    _ = v98
                                                                                                    _ = v99
                                                                                                }
                                                                                            case "\n\t":
                                                                                                v3 -= int8(fi17)
                                                                                            default:
                                                                                                fmt.Println(v3)
                                                                                            }
                                                                                        }
                                                                                    case -1:
                                                                                        v6[4294967295] = v1
                                                                                    }
                                                                                    v16 = (v16 | (v20 % (((v1 ^ (v16 + (v1 / ((v20) | 1))))) | 1)))
                                                                                    v1 = v20
                                                                                    _ = v90
                                                                                }
                                                                            default:
                                                                                var v100 uint16 = 32767
                                                                                v1 -= v20
                                                                                v101 := GenS7{F0: v4}
                                                                                _ = v100
                                                                                _ = v101
                                                                            }
                                                                            fmt.Println(v16, fi17, v20)
                                                                        }
                                                                    }
                                                                    v102 := GenS1{F0: (^(v3)), F1: "a"}
                                                                    _ = v102
                                                                }
                                                                var v103 uint8 = 1
                                                                v4 -= (v4 % (((v2 / ((2) | 1))) | 1))
                                                                _ = v103
                                                            }
                                                        }
                                                    }
                                                }
                                            case -994:
                                                var v104 uint32 = (v16 - (v1 * v1))
                                                _ = v104
                                            }
                                        } else {
                                            if true {
                                                _ = len(v6)
                                            }
                                        }
                                        _ = v20
                                    } else {
                                        _ = len(v6)
                                        v2 = v4
                                        if (!(175 <= 90)) {
                                            delete(v6, 2)
                                        }
                                        v2 += 188
                                    }
                                }
                                v2 += (v4 >> ((2) & 63))
                                _ = fi17
                            }
                        }
                        _ = v16
                    }
                    _ = fi15
                }
                switch ("foo bar" + ("" + "foo bar")) {
                case "":
                    fmt.Println(v2, v4)
                default:
                    for fi105 := 0; fi105 < 4; fi105++ {
                        switch (fi105 ^ 128) {
                        case -1:
                            if true {
                                _ = len(v6)
                            }
                        case 128:
                            var v106 uint32 = v1
                            _ = v106
                        }
                        switch (("\n\t" + "") + "") {
                        case "a":
                            v107 := GenS11{F0: 42, F1: (fi105 << ((1) & 63))}
                            _ = v107
                        case "foo bar":
                            switch ("\n\t" + "a") {
                            case "":
                                var v108 bool = (v4 <= v2)
                                _ = v108
                            case "a":
                                fmt.Println(v1, v2)
                            default:
                                fmt.Println(v4, v2, v3)
                            }
                        default:
                            var v109 bool = true
                            _ = v109
                        }
                        _ = fi105
                    }
                }
                var v110 int32 = -804
                _ = v110
            }
        }
    } else {
        var v111 uint = 65535
        v112 := GenS9{F0: 2, F1: -1, F2: 9223372036854775807}
        _ = v111
        _ = v112
    }
    v113 := GenS12{F0: 9223372036854775807, F1: 128, F2: -18}
    fmt.Println(v2, v3, v4)
    if ((v3 >> ((1) & 63)) < v3) {
        for fi114 := 0; fi114 < 4; fi114++ {
            for fi115 := 0; fi115 < 5; fi115++ {
                v116 := GenS6{F0: -9223372036854775808, F1: 42}
                v5 = append(v5, (v1 + (v1 - 409)))
                _ = fi115
                _ = v116
            }
            _ = fi114
        }
        var v117 int8 = v3
        _ = v117
    } else {
        for fi118 := 0; fi118 < 11; fi118++ {
            var v119 int8 = (-(v3))
            v120 := GenS13{F0: 256}
            for fi121 := 0; fi121 < 16; fi121++ {
                var v122 uint64 = 4294967295
                var v123 int32 = 1
                _ = fi121
                _ = v122
                _ = v123
            }
            _ = fi118
            _ = v119
            _ = v120
        }
    }
    v3 -= (v3 >> ((65535) & 63))
    if (true && false) {
        v124 := GenS4{F0: ("foo bar" < ""), F1: (v1 % (((v1 - (v1 >> ((65535) & 63)))) | 1)), F2: 18446744073709551615}
        for fi125 := 0; fi125 < 12; fi125++ {
            var v126 float32 = ((1.5 + 1e10) + ((1e10 * (float32((1e10 / 0.1)) + 0.1)) + 0.0))
            for fi127 := 0; fi127 < 7; fi127++ {
                v3 -= -1
                switch "foo bar" {
                case "a":
                    v126 = 0.0
                case "":
                    delete(v6, uint32(v2))
                case "hello":
                    v126 = v126
                }
                fmt.Println(v126, v3)
                var v128 int16 = 718
                v3 = (v3 << ((1) & 63))
                _ = fi127
                _ = v128
            }
            v2 -= (v4 ^ (v2 / ((v4) | 1)))
            _ = fi125
            _ = v126
        }
        _ = v124
    } else {
        v4 = (v2 - v2)
        if (v1 <= v1) {
            for fi129 := 0; fi129 < 3; fi129++ {
                if false {
                    v3 = 1
                }
                v130 := GenS8{F0: v2, F1: (!false)}
                _ = fi129
                _ = v130
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            for fi131 := 0; fi131 < 1; fi131++ {
                v132 := GenS14{F0: fi131, F1: 311}
                switch "a" {
                case "foo bar":
                    if true {
                        var v133 uint32 = (v1 - v1)
                        _ = v133
                    }
                default:
                    v134 := GenS15{F0: "", F1: "a", F2: (v3 > (v3 - 100))}
                    _ = v134
                }
                if (0 >= 42) {
                    switch "\n\t" {
                    case "\n\t":
                        for fi135 := 0; fi135 < 3; fi135++ {
                            switch "foo bar" {
                            case "foo bar":
                                v136 := GenS3{F0: -2147483648, F1: (fi135 / (((fi135 & int(v3))) | 1))}
                                _ = v136
                            }
                            var v137 uint16 = 478
                            _ = fi135
                            _ = v137
                        }
                    default:
                        delete(v6, (v1 + 4294967295))
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                _ = fi131
                _ = v132
            }
        } else {
            if false {
                for fi138 := 0; fi138 < 2; fi138++ {
                    v139 := v6[(^(v1))]
                    fmt.Println(v3, v2)
                    if false {
                        fmt.Println(v1, v139, v1)
                        if false {
                            if false {
                                var v140 float32 = float32(0.1)
                                v141 := GenS3{F0: -2147483648, F1: fi138}
                                _ = v140
                                _ = v141
                            }
                        }
                    }
                    _ = fi138
                    _ = v139
                }
            } else {
                delete(v6, v1)
            }
            fmt.Println(v1)
            for fi142 := 0; fi142 < 12; fi142++ {
                v2 += (v2 - v4)
                v143 := GenS3{F0: 0, F1: fi142}
                v4 = v4
                _ = fi142
                _ = v143
            }
        }
    }
    switch ("a" + "hello") {
    case "":
        v144 := GenS16{F0: (v4 >> ((0) & 63)), F1: genFn0(-9223372036854775808), F2: (200 == 18446744073709551615)}
        _ = v144
    case "\n\t":
        var v145 uint8 = (^(v2))
        _ = v145
    default:
        v2 -= (v2 >> ((159) & 63))
    }
    if (true || false) {
        v4 += v2
        v3 = v3
        fmt.Println(v4, v1)
    }
    var v146 int32 = -540
    var v147 float32 = 1.5
    v148 := GenS7{F0: (v2 >> ((311) & 63))}
    fmt.Println(v147, v4)
    switch 127 {
    case 101:
        v146 -= v146
    case 32768:
        switch -2147483648 {
        case -9223372036854775808:
            var v149 int32 = 908
            _ = v149
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v146:", v146)
    fmt.Println("v147:", v147)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v113
    _ = v146
    _ = v147
    _ = v148
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
