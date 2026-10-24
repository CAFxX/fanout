package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 int64; F2 string }
type GenS3 struct { F0 int8; F1 int16; F2 int16 }
type GenS4 struct { F0 int32; F1 int; F2 uint64 }
type GenS5 struct { F0 int64; F1 uint }
type GenS6 struct { F0 bool; F1 uint8; F2 uint32 }
type GenS7 struct { F0 uint32; F1 uint8 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint8; F1 uint }
type GenS10 struct { F0 int; F1 int8 }
type GenS11 struct { F0 string; F1 int64 }
type GenS12 struct { F0 bool; F1 uint32 }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 uint64 }
type GenS15 struct { F0 uint8; F1 uint8 }
type GenS16 struct { F0 uint; F1 uint16; F2 string }
type GenS17 struct { F0 uint32; F1 uint16; F2 uint64 }
type GenS18 struct { F0 int }
type GenS19 struct { F0 int16 }
type GenS20 struct { F0 uint32; F1 uint8 }
type GenS21 struct { F0 uint32; F1 int64 }
type GenS22 struct { F0 uint64; F1 uint16 }
type GenS23 struct { F0 int; F1 string }
type GenS24 struct { F0 uint16; F1 bool; F2 uint64 }
type GenS25 struct { F0 int8 }
type GenS26 struct { F0 int; F1 string }
type GenS27 struct { F0 string; F1 int32; F2 int32 }
type GenS28 struct { F0 int16 }
type GenS29 struct { F0 bool; F1 int32 }
type GenS30 struct { F0 uint32; F1 uint64 }
type GenS31 struct { F0 uint16; F1 int }
type GenS32 struct { F0 uint; F1 uint8 }
type GenS33 struct { F0 uint16; F1 uint16; F2 bool }
func genFn0() int32 {
    if ((0 != 2147483647) && false) {
        var v4 uint32 = 4294967295
        v4 = (v4 | v4)
        for fi5 := 0; fi5 < 19; fi5++ {
            var v6 string = "hello"
            v4 = 1
            _ = fi5
            _ = v6
        }
        _ = v4
    }
    if (-1 == 71) {
        var v7 int8 = -15
        _ = v7
    }
    return -1
}

func main() {
    var v1 float32 = -2.25
    v2 := []string{"hello", "\n\t", ("a" + ((((("\n\t" + "a") + "a") + ("\n\t" + "a")) + "hello") + "foo bar"))}
    var v3 uint64 = 1
    switch -168 {
    case 133:
        if len(v2) > 0 {
            v2[0] = ("a" + ("foo bar" + ""))
        }
    case 128:
        switch (("a" + "a") + ("a" + ("a" + "a"))) {
        case "":
            if len(v2) > 0 {
                v2[0] = "\n\t"
            }
        case "foo bar":
            for fi8 := 0; fi8 < 19; fi8++ {
                if (((false || (!(478 >= int64(fi8)))) && (4294967295 >= 65535)) && (1 > 65535)) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v9 uint16 = 2
                    v3 += v3
                    _ = v9
                }
                v10 := GenS1{F0: 4294967295}
                if true {
                    if true {
                        var v11 uint = 2
                        v1 += v1
                        v11 = v11
                        _ = v11
                    }
                    var v12 int32 = 2147483647
                    _ = v12
                } else {
                    v3 = 536
                    if (4294967295 >= 18446744073709551615) {
                        for fi13 := 0; fi13 < 5; fi13++ {
                            v14 := GenS2{F0: 0, F1: 42, F2: "foo bar"}
                            v1 = v1
                            var v15 string = "\n\t"
                            switch fi13 {
                            case -352:
                                if (true && false) {
                                    v2 = append(v2, ((((v15 + (v15 + v15)) + (v15 + "")) + "hello") + "\n\t"))
                                }
                            case -2147483648:
                                if true {
                                    var v16 int8 = -128
                                    _ = v16
                                } else {
                                    var v17 int8 = 42
                                    fmt.Println(v15)
                                    var v18 bool = (!true)
                                    if v18 {
                                        switch (v15 + "a") {
                                        case "foo bar":
                                            v15 = v15
                                        case "a":
                                            switch fi8 {
                                            case 54:
                                                if (v3 != 875) {
                                                    if (-866 > -1) {
                                                        var v19 float64 = 0.0
                                                        _ = v19
                                                    }
                                                } else {
                                                    var v20 uint16 = 42
                                                    switch v15 {
                                                    case "foo bar":
                                                        var v21 int64 = 775
                                                        _ = v21
                                                    case "a":
                                                        v2 = append(v2, "hello")
                                                    case "hello":
                                                        v22 := GenS3{F0: 28, F1: int16(v3), F2: 617}
                                                        _ = v22
                                                    default:
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    }
                                                    switch fi8 {
                                                    case 1:
                                                        if ((!(0 >= -32768)) || (!v18)) {
                                                            if v18 {
                                                                if v18 {
                                                                    switch "foo bar" {
                                                                    case "hello":
                                                                        var v23 int8 = (v17 % (((-(v17))) | 1))
                                                                        _ = v23
                                                                    default:
                                                                        _ = len(v2)
                                                                    }
                                                                    if len(v2) > 0 {
                                                                        v2[0] = "hello"
                                                                    }
                                                                } else {
                                                                    v24 := GenS2{F0: 2, F1: 384, F2: ("hello" + "a")}
                                                                    _ = v24
                                                                }
                                                                if true {
                                                                    var v25 int8 = (v17 * v17)
                                                                    _ = v25
                                                                } else {
                                                                    var v26 int64 = 0
                                                                    v27 := GenS1{F0: 395}
                                                                    v26 -= v26
                                                                    _ = v26
                                                                    _ = v27
                                                                }
                                                                v20 = 255
                                                            }
                                                            fmt.Println(fi8, fi8, v1)
                                                            if true {
                                                                v20 += 0
                                                            } else {
                                                                switch v15 {
                                                                case "foo bar":
                                                                    var v28 int64 = 382
                                                                    v29 := GenS4{F0: 419, F1: 1, F2: v3}
                                                                    fmt.Println(fi8)
                                                                    if v18 {
                                                                        _ = len(v2)
                                                                        var v30 int16 = 0
                                                                        _ = v30
                                                                    } else {
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                        if v18 {
                                                                            if v18 {
                                                                                v3 = v3
                                                                                v17 -= 0
                                                                            } else {
                                                                                var v31 float32 = 1.5
                                                                                _ = v31
                                                                            }
                                                                            v18 = ((v28 * 42) < 9223372036854775807)
                                                                            switch v15 {
                                                                            case "":
                                                                                v3 += (v3 & v3)
                                                                            case "foo bar":
                                                                                if (!(!v18)) {
                                                                                    var v32 uint8 = 1
                                                                                    if v18 {
                                                                                        v32 = (v32 - v32)
                                                                                        switch (v15 + v15) {
                                                                                        case "":
                                                                                            switch fi8 {
                                                                                            case 269:
                                                                                                fmt.Println(v20, v32, v20)
                                                                                            case 1:
                                                                                                fmt.Println(v20)
                                                                                            }
                                                                                        default:
                                                                                            v3 = (^(v3))
                                                                                        }
                                                                                        if false {
                                                                                            switch v15 {
                                                                                            case "a":
                                                                                                fmt.Println(v15, v17)
                                                                                            case "hello":
                                                                                                v3 += 658
                                                                                            case "\n\t":
                                                                                                v33 := GenS5{F0: 0, F1: 255}
                                                                                                _ = v33
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    var v34 string = "hello"
                                                                                    v3 = v3
                                                                                    fmt.Println(v18, v15, v20)
                                                                                    _ = v32
                                                                                    _ = v34
                                                                                }
                                                                            case "hello":
                                                                                v17 += 55
                                                                            }
                                                                            fmt.Println(v1, fi13, v18)
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                        } else {
                                                                            if v18 {
                                                                                switch ("\n\t" + "foo bar") {
                                                                                case "\n\t":
                                                                                    if len(v2) >= 2 {
                                                                                        v2 = v2[1:len(v2)]
                                                                                    }
                                                                                case "hello":
                                                                                    var v35 int32 = 2147483647
                                                                                    _ = v35
                                                                                case "a":
                                                                                    fmt.Println(v3, v17, v18)
                                                                                }
                                                                                if false {
                                                                                    switch -468 {
                                                                                    case 32768:
                                                                                        v36 := GenS6{F0: ((1 != 140) || false), F1: 128, F2: 65535}
                                                                                        _ = v36
                                                                                    case 161:
                                                                                        var v37 int64 = v28
                                                                                        _ = v37
                                                                                    }
                                                                                    fmt.Println(v20, fi8)
                                                                                    if v18 {
                                                                                        v38 := GenS7{F0: 50, F1: 255}
                                                                                        fmt.Println(v20, v15, v17)
                                                                                        var v39 int16 = -1
                                                                                        if len(v2) >= 2 {
                                                                                            v2 = v2[1:len(v2)]
                                                                                        }
                                                                                        if len(v2) >= 2 {
                                                                                            v2 = v2[1:len(v2)]
                                                                                        }
                                                                                        _ = v38
                                                                                        _ = v39
                                                                                    } else {
                                                                                        var v40 uint16 = v20
                                                                                        _ = v40
                                                                                    }
                                                                                    var v41 int64 = v28
                                                                                    _ = v41
                                                                                } else {
                                                                                    v15 = v15
                                                                                    var v42 float32 = ((float32(1.5) - v1) / v1)
                                                                                    _ = v42
                                                                                }
                                                                                _ = len(v2)
                                                                            } else {
                                                                                var v43 uint16 = v20
                                                                                v44 := GenS1{F0: 255}
                                                                                _ = v43
                                                                                _ = v44
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v28
                                                                    _ = v29
                                                                case "hello":
                                                                    v1 -= 1.5
                                                                case "\n\t":
                                                                    var v45 uint64 = v3
                                                                    _ = v45
                                                                }
                                                                v3 += (v3 << ((2) & 63))
                                                                var v46 string = v15
                                                                _ = v46
                                                            }
                                                            var v47 int64 = 890
                                                            fmt.Println(v1, v15)
                                                            _ = v47
                                                        }
                                                        var v48 uint64 = (v3 - v3)
                                                        v48 += v48
                                                        _ = v48
                                                    case 2147483647:
                                                        v15 = ""
                                                    default:
                                                        if (4294967295 < 47) {
                                                            v3 += 18446744073709551615
                                                            fmt.Println(v20, fi8, v18)
                                                            if v18 {
                                                                if (!v18) {
                                                                    v49 := GenS8{F0: genFn0()}
                                                                    if (!true) {
                                                                        if v18 {
                                                                            var v50 uint32 = 65535
                                                                            _ = len(v2)
                                                                            _ = v50
                                                                        }
                                                                        switch ("\n\t" + "") {
                                                                        case "":
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                        case "hello":
                                                                            fmt.Println(v18, v20)
                                                                        default:
                                                                            switch int(v3) {
                                                                            case -373:
                                                                                fmt.Println(v18, fi8)
                                                                            case -9223372036854775808:
                                                                                var v51 bool = (!v18)
                                                                                _ = v51
                                                                            case -244:
                                                                                switch 42 {
                                                                                case -2147483648:
                                                                                    if true {
                                                                                        v17 = -12
                                                                                    } else {
                                                                                        var v52 float32 = float32(3.14159)
                                                                                        v3 = (v3 ^ v3)
                                                                                        v17 -= (v17 / ((v17) | 1))
                                                                                        _ = v52
                                                                                    }
                                                                                default:
                                                                                    v1 = (v1 * v1)
                                                                                }
                                                                            default:
                                                                                _ = len(v2)
                                                                            }
                                                                        }
                                                                    }
                                                                    v2 = append(v2, "")
                                                                    _ = v49
                                                                }
                                                                v18 = (v1 > v1)
                                                                switch v15 {
                                                                case "a":
                                                                    var v53 string = ""
                                                                    _ = v53
                                                                }
                                                            } else {
                                                                fmt.Println(v17)
                                                                var v54 uint = 65535
                                                                v20 += v20
                                                                _ = v54
                                                            }
                                                            _ = len(v2)
                                                        } else {
                                                            v15 = v15
                                                            var v55 int16 = -646
                                                            var v56 float64 = (0.0 / (0.1 / 3.14159))
                                                            _ = v55
                                                            _ = v56
                                                        }
                                                    }
                                                    v18 = false
                                                    _ = v20
                                                }
                                            case 127:
                                                if v18 {
                                                    if ((v3 << ((42) & 63)) > v3) {
                                                        if v18 {
                                                            if (true || false) {
                                                                fmt.Println(v3, v18, fi8)
                                                                v1 -= v1
                                                                if false {
                                                                    v57 := GenS9{F0: 128, F1: 42}
                                                                    var v58 int = int(v3)
                                                                    v1 = ((0.0 / v1) / ((float32(3.14159) / float32(-2.25)) / (float32(1e10) - 0.0)))
                                                                    _ = v57
                                                                    _ = v58
                                                                }
                                                            } else {
                                                                switch fi13 {
                                                                case 4294967295:
                                                                    if len(v2) > 0 {
                                                                        v2[0] = ("" + "a")
                                                                    }
                                                                }
                                                                v1 += 1e10
                                                                v59 := GenS1{F0: 54}
                                                                _ = v59
                                                            }
                                                            if v18 {
                                                                if false {
                                                                    if (797 >= 4294967295) {
                                                                        if (v18 || v18) {
                                                                            var v60 uint64 = (v3 ^ 42)
                                                                            if (v18 && (!(!v18))) {
                                                                                switch ((((("a" + v15) + (v15 + (v15 + ("foo bar" + "a")))) + "hello") + v15) + v15) {
                                                                                case "":
                                                                                    v15 = ((("foo bar" + v15) + v15) + (v15 + ("foo bar" + v15)))
                                                                                default:
                                                                                    fmt.Println(v3, v15)
                                                                                }
                                                                            }
                                                                            switch v15 {
                                                                            case "\n\t":
                                                                                v61 := GenS10{F0: fi13, F1: v17}
                                                                                _ = v61
                                                                            }
                                                                            _ = v60
                                                                        } else {
                                                                            if v18 {
                                                                                var v62 uint32 = 779
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                                _ = len(v2)
                                                                                v2 = append(v2, "")
                                                                                if (128 == 216) {
                                                                                    switch 428 {
                                                                                    case 127:
                                                                                        if (!v18) {
                                                                                            v63 := GenS11{F0: ((("foo bar" + v15) + v15) + v15), F1: int64(v62)}
                                                                                            var v64 uint8 = 2
                                                                                            if ((4294967295 > v62) || v18) {
                                                                                                v18 = ((fi13 & fi13) <= fi8)
                                                                                                var v65 uint32 = 0
                                                                                                _ = v65
                                                                                            } else {
                                                                                                if (478 > 2) {
                                                                                                    v18 = (-9223372036854775808 != 994)
                                                                                                    if v18 {
                                                                                                        if false {
                                                                                                            v18 = false
                                                                                                            if v18 {
                                                                                                                _ = len(v2)
                                                                                                            } else {
                                                                                                                fmt.Println(v17)
                                                                                                                fmt.Println(fi13)
                                                                                                            }
                                                                                                            if (((v64 * (v64 + v64)) < v64) || v18) {
                                                                                                                var v66 int32 = genFn0()
                                                                                                                var v67 int8 = (-(v17))
                                                                                                                v15 = "foo bar"
                                                                                                                _ = v66
                                                                                                                _ = v67
                                                                                                            } else {
                                                                                                                v18 = (-927 < 0)
                                                                                                                fmt.Println(v18)
                                                                                                                if true {
                                                                                                                    var v68 float64 = (1.5 + (0.1 - 0.1))
                                                                                                                    _ = v68
                                                                                                                } else {
                                                                                                                    if ((fi13 * (^(fi13))) <= fi13) {
                                                                                                                        var v69 float32 = -2.25
                                                                                                                        var v70 uint = 65535
                                                                                                                        if v18 {
                                                                                                                            if (0.0 <= 0.1) {
                                                                                                                                if true {
                                                                                                                                    v71 := GenS4{F0: int32(v17), F1: (fi8 * int(v64)), F2: (v3 ^ v3)}
                                                                                                                                    _ = v71
                                                                                                                                }
                                                                                                                                v72 := GenS1{F0: v70}
                                                                                                                                _ = v72
                                                                                                                            }
                                                                                                                            switch v15 {
                                                                                                                            case "":
                                                                                                                                fmt.Println(v69, v69)
                                                                                                                            default:
                                                                                                                                if (v18 || (((v69 + float32(0.0)) / v1) > -2.25)) {
                                                                                                                                    v73 := GenS6{F0: (v64 < v64), F1: v64, F2: (v62 / (((v62 + (v62 - (v62 / (((v62 << ((v70) & 63))) | 1))))) | 1))}
                                                                                                                                    _ = v73
                                                                                                                                } else {
                                                                                                                                    if v18 {
                                                                                                                                        v62 = v62
                                                                                                                                        switch fi13 {
                                                                                                                                        case 670:
                                                                                                                                            if (!v18) {
                                                                                                                                                var v74 uint64 = (^(v3))
                                                                                                                                                v75 := GenS12{F0: v18, F1: (v62 | (v62 ^ v62))}
                                                                                                                                                fmt.Println(v70, v3, fi13)
                                                                                                                                                _ = v74
                                                                                                                                                _ = v75
                                                                                                                                            }
                                                                                                                                        case -412:
                                                                                                                                            var v76 int8 = v17
                                                                                                                                            _ = v76
                                                                                                                                        case -723:
                                                                                                                                            if v18 {
                                                                                                                                                v3 += (v3 | (v3 % ((v3) | 1)))
                                                                                                                                                v64 -= (v64 - v64)
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                        v77 := GenS3{F0: (v17 | v17), F1: int16(v70), F2: -32768}
                                                                                                                                        _ = v77
                                                                                                                                    } else {
                                                                                                                                        switch v15 {
                                                                                                                                        case "foo bar":
                                                                                                                                            var v78 uint32 = (v62 << ((uint(v64)) & 63))
                                                                                                                                            _ = v78
                                                                                                                                        case "hello":
                                                                                                                                            v79 := GenS6{F0: v18, F1: v64, F2: 42}
                                                                                                                                            _ = v79
                                                                                                                                        case "a":
                                                                                                                                            v15 = v15
                                                                                                                                        default:
                                                                                                                                            var v80 uint16 = uint16(v64)
                                                                                                                                            _ = v80
                                                                                                                                        }
                                                                                                                                        var v81 uint16 = 65535
                                                                                                                                        _ = v81
                                                                                                                                    }
                                                                                                                                    if v18 {
                                                                                                                                        var v82 int16 = 32767
                                                                                                                                        switch v15 {
                                                                                                                                        case "hello":
                                                                                                                                            var v83 int32 = -1
                                                                                                                                            _ = v83
                                                                                                                                        case "a":
                                                                                                                                            if true {
                                                                                                                                                v15 = "foo bar"
                                                                                                                                                v84 := GenS9{F0: (v64 >> (((v70 + v70)) & 63)), F1: v70}
                                                                                                                                                fmt.Println(fi8)
                                                                                                                                                _ = v84
                                                                                                                                            }
                                                                                                                                        case "foo bar":
                                                                                                                                            v17 += -26
                                                                                                                                        default:
                                                                                                                                            var v85 int = fi8
                                                                                                                                            _ = v85
                                                                                                                                        }
                                                                                                                                        _ = v82
                                                                                                                                    }
                                                                                                                                }
                                                                                                                            }
                                                                                                                            var v86 uint16 = 65535
                                                                                                                            _ = v86
                                                                                                                        } else {
                                                                                                                            fmt.Println(fi8, v18, v70)
                                                                                                                            fmt.Println(v18, v15)
                                                                                                                            if v18 {
                                                                                                                                if ((v64 + (^(v64))) != 2) {
                                                                                                                                    if v18 {
                                                                                                                                        v3 = v3
                                                                                                                                        v70 = v70
                                                                                                                                        var v87 uint16 = 65535
                                                                                                                                        v69 += v69
                                                                                                                                        _ = v87
                                                                                                                                    } else {
                                                                                                                                        var v88 bool = v18
                                                                                                                                        _ = v88
                                                                                                                                    }
                                                                                                                                    switch (^(fi8)) {
                                                                                                                                    case 42:
                                                                                                                                        if len(v2) >= 2 {
                                                                                                                                            v2 = v2[1:len(v2)]
                                                                                                                                        }
                                                                                                                                    case 1:
                                                                                                                                        fmt.Println(v1)
                                                                                                                                    case 842:
                                                                                                                                        fmt.Println(v64)
                                                                                                                                    }
                                                                                                                                    var v89 int64 = 1
                                                                                                                                    _ = v89
                                                                                                                                }
                                                                                                                                v3 += (v3 << (((v70 + 197)) & 63))
                                                                                                                                if false {
                                                                                                                                    var v90 uint64 = (v3 - v3)
                                                                                                                                    _ = v90
                                                                                                                                } else {
                                                                                                                                    if len(v2) > 0 {
                                                                                                                                        v2[0] = (v15 + "\n\t")
                                                                                                                                    }
                                                                                                                                    if len(v2) > 0 {
                                                                                                                                        v2[0] = v15
                                                                                                                                    }
                                                                                                                                    var v91 int64 = -1
                                                                                                                                    _ = v91
                                                                                                                                }
                                                                                                                                _ = len(v2)
                                                                                                                                var v92 int64 = -873
                                                                                                                                _ = v92
                                                                                                                            }
                                                                                                                        }
                                                                                                                        _ = v69
                                                                                                                        _ = v70
                                                                                                                    }
                                                                                                                    fmt.Println(v17, v15, v1)
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    if (v18 && false) {
                                                                                                        var v93 float64 = 3.14159
                                                                                                        switch (fi8 >> ((754) & 63)) {
                                                                                                        case -134:
                                                                                                            v94 := GenS13{F0: -227}
                                                                                                            _ = v94
                                                                                                        }
                                                                                                        _ = v93
                                                                                                    } else {
                                                                                                        var v95 int64 = 0
                                                                                                        _ = v95
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            _ = v63
                                                                                            _ = v64
                                                                                        } else {
                                                                                            if true {
                                                                                                v96 := GenS14{F0: (^(v3))}
                                                                                                v3 = v3
                                                                                                var v97 int = 128
                                                                                                _ = len(v2)
                                                                                                _ = v96
                                                                                                _ = v97
                                                                                            }
                                                                                            _ = len(v2)
                                                                                        }
                                                                                    case 1:
                                                                                        fmt.Println(fi8)
                                                                                    case 128:
                                                                                        v15 = v15
                                                                                    default:
                                                                                        fmt.Println(v62)
                                                                                    }
                                                                                    v2 = append(v2, v15)
                                                                                } else {
                                                                                    if len(v2) >= 2 {
                                                                                        v2 = v2[1:len(v2)]
                                                                                    }
                                                                                    if len(v2) >= 2 {
                                                                                        v2 = v2[1:len(v2)]
                                                                                    }
                                                                                }
                                                                                v98 := GenS10{F0: (-(fi13)), F1: (^(v17))}
                                                                                var v99 string = v15
                                                                                switch fi8 {
                                                                                case -1:
                                                                                    fmt.Println(v3, v15)
                                                                                case -9223372036854775808:
                                                                                    var v100 int = (fi13 | -1)
                                                                                    _ = v100
                                                                                }
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                                _ = v62
                                                                                _ = v98
                                                                                _ = v99
                                                                            }
                                                                            var v101 float64 = 0.0
                                                                            var v102 int8 = 34
                                                                            _ = v101
                                                                            _ = v102
                                                                        }
                                                                    } else {
                                                                        if (!(false && (229 != fi13))) {
                                                                            if (-2147483648 <= -636) {
                                                                                v17 += v17
                                                                                if ((1.5 + -2.25) < (1.5 * ((v1 * (-2.25 / v1)) - 0.0))) {
                                                                                    if len(v2) >= 2 {
                                                                                        v2 = v2[1:len(v2)]
                                                                                    }
                                                                                    var v103 int8 = 55
                                                                                    var v104 uint64 = (v3 << ((42) & 63))
                                                                                    _ = v103
                                                                                    _ = v104
                                                                                } else {
                                                                                    v2 = append(v2, v15)
                                                                                    var v105 uint = 1
                                                                                    switch -1 {
                                                                                    case -972:
                                                                                        if (v15 == "") {
                                                                                            switch (v15 + "a") {
                                                                                            case "a":
                                                                                                v2 = append(v2, v15)
                                                                                            default:
                                                                                                if (true && v18) {
                                                                                                    var v106 uint16 = 0
                                                                                                    var v107 int = fi13
                                                                                                    v1 -= v1
                                                                                                    _ = v106
                                                                                                    _ = v107
                                                                                                } else {
                                                                                                    v17 += -1
                                                                                                    var v108 bool = v18
                                                                                                    var v109 uint = v105
                                                                                                    _ = v108
                                                                                                    _ = v109
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    case 550:
                                                                                        v110 := GenS15{F0: 220, F1: 0}
                                                                                        _ = v110
                                                                                    default:
                                                                                        var v111 int64 = 923
                                                                                        _ = v111
                                                                                    }
                                                                                    _ = v105
                                                                                }
                                                                            }
                                                                            _ = len(v2)
                                                                            var v112 bool = false
                                                                            _ = v112
                                                                        } else {
                                                                            v113 := GenS12{F0: v18, F1: 2147483648}
                                                                            var v114 int16 = 658
                                                                            v3 = (v3 | (v3 / (((^(v3))) | 1)))
                                                                            if ((775 >= 694) && (4294967295 < v3)) {
                                                                                v115 := GenS11{F0: ("a" + v15), F1: 341}
                                                                                if (18446744073709551615 <= 42) {
                                                                                    switch int(v114) {
                                                                                    case 32768:
                                                                                        if (v114 != -611) {
                                                                                            if v18 {
                                                                                                if (v15 <= (v15 + v15)) {
                                                                                                    v116 := GenS1{F0: 255}
                                                                                                    _ = v116
                                                                                                } else {
                                                                                                    var v117 int16 = (v114 % (((v114 % ((v114) | 1))) | 1))
                                                                                                    _ = v117
                                                                                                }
                                                                                            }
                                                                                            var v118 uint32 = 42
                                                                                            _ = len(v2)
                                                                                            _ = v118
                                                                                        } else {
                                                                                            if (v17 == (v17 / (((-(v17))) | 1))) {
                                                                                                _ = len(v2)
                                                                                                switch "foo bar" {
                                                                                                case "foo bar":
                                                                                                    v3 = (v3 - 871)
                                                                                                case "":
                                                                                                    fmt.Println(v15)
                                                                                                case "a":
                                                                                                    if (!true) {
                                                                                                        if v18 {
                                                                                                            v18 = (279 <= 0)
                                                                                                            v119 := GenS11{F0: (v15 + "hello"), F1: 42}
                                                                                                            var v120 uint8 = 255
                                                                                                            _ = v119
                                                                                                            _ = v120
                                                                                                        }
                                                                                                        var v121 float64 = 1.5
                                                                                                        v121 += v121
                                                                                                        _ = v121
                                                                                                    } else {
                                                                                                        var v122 float64 = 0.0
                                                                                                        v1 -= v1
                                                                                                        _ = v122
                                                                                                    }
                                                                                                default:
                                                                                                    _ = len(v2)
                                                                                                    fmt.Println(v114)
                                                                                                    var v123 bool = ((v3 > 0) || v18)
                                                                                                    _ = v123
                                                                                                }
                                                                                                if false {
                                                                                                    var v124 uint8 = 42
                                                                                                    var v125 bool = v18
                                                                                                    _ = v124
                                                                                                    _ = v125
                                                                                                }
                                                                                            } else {
                                                                                                if v18 {
                                                                                                    switch v15 {
                                                                                                    case "":
                                                                                                        v126 := GenS5{F0: 9223372036854775807, F1: 210}
                                                                                                        _ = v126
                                                                                                    case "foo bar":
                                                                                                        var v127 string = ((v15 + v15) + v15)
                                                                                                        _ = v127
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            if (v114 < (v114 / ((v114) | 1))) {
                                                                                                var v128 bool = v18
                                                                                                v114 = -31
                                                                                                _ = len(v2)
                                                                                                _ = v128
                                                                                            } else {
                                                                                                if len(v2) >= 2 {
                                                                                                    v2 = v2[1:len(v2)]
                                                                                                }
                                                                                                v114 += v114
                                                                                            }
                                                                                        }
                                                                                    case -590:
                                                                                        fmt.Println(v15)
                                                                                    case 2147483647:
                                                                                        v2 = append(v2, v15)
                                                                                    default:
                                                                                        switch "hello" {
                                                                                        case "foo bar":
                                                                                            var v129 string = ("a" + (v15 + "\n\t"))
                                                                                            _ = v129
                                                                                        case "":
                                                                                            var v130 bool = (-29 >= -1)
                                                                                            _ = v130
                                                                                        default:
                                                                                            v131 := GenS16{F0: 395, F1: 42, F2: v15}
                                                                                            _ = v131
                                                                                        }
                                                                                    }
                                                                                    if false {
                                                                                        if len(v2) >= 2 {
                                                                                            v2 = v2[1:len(v2)]
                                                                                        }
                                                                                        v3 -= v3
                                                                                    } else {
                                                                                        var v132 uint = 2
                                                                                        v133 := GenS4{F0: 42, F1: int(v3), F2: v3}
                                                                                        _ = v132
                                                                                        _ = v133
                                                                                    }
                                                                                } else {
                                                                                    var v134 uint = 931
                                                                                    var v135 bool = (!(!(v1 != (v1 * 3.14159))))
                                                                                    _ = v134
                                                                                    _ = v135
                                                                                }
                                                                                var v136 uint = 18446744073709551615
                                                                                v137 := GenS17{F0: 2147483647, F1: 2, F2: (^(v3))}
                                                                                var v138 uint32 = 659
                                                                                _ = v115
                                                                                _ = v136
                                                                                _ = v137
                                                                                _ = v138
                                                                            }
                                                                            var v139 bool = true
                                                                            v2 = append(v2, "")
                                                                            var v140 float32 = v1
                                                                            _ = v113
                                                                            _ = v114
                                                                            _ = v139
                                                                            _ = v140
                                                                        }
                                                                        v3 = v3
                                                                    }
                                                                    if (((0.1 / 1.5) != -2.25) || v18) {
                                                                        if ((!((2 > 1) || (63 <= 0))) || ((!false) && false)) {
                                                                            fmt.Println(fi8, v18)
                                                                        } else {
                                                                            var v141 uint16 = 32767
                                                                            var v142 int16 = 0
                                                                            _ = v141
                                                                            _ = v142
                                                                        }
                                                                    } else {
                                                                        switch (fi8 & (fi8 >> ((2) & 63))) {
                                                                        case 9223372036854775807:
                                                                            v3 += v3
                                                                        case 1:
                                                                            v15 = (("\n\t" + v15) + (((v15 + v15) + v15) + "hello"))
                                                                        }
                                                                        if (!(!false)) {
                                                                            var v143 float64 = 1e10
                                                                            v143 -= -2.25
                                                                            var v144 string = ("\n\t" + "")
                                                                            _ = v143
                                                                            _ = v144
                                                                        } else {
                                                                            v2 = append(v2, "hello")
                                                                            if (2 <= 740) {
                                                                                var v145 uint8 = 157
                                                                                v146 := GenS18{F0: 2147483647}
                                                                                fmt.Println(v18, v17)
                                                                                v147 := GenS19{F0: 42}
                                                                                _ = v145
                                                                                _ = v146
                                                                                _ = v147
                                                                            } else {
                                                                                var v148 int16 = -32768
                                                                                fmt.Println(v15, v15, fi13)
                                                                                var v149 uint64 = (v3 % ((v3) | 1))
                                                                                _ = len(v2)
                                                                                _ = v148
                                                                                _ = v149
                                                                            }
                                                                            v18 = (597 <= (fi8 & (^(fi13))))
                                                                        }
                                                                        var v150 float64 = (1.5 * (1e10 - 0.0))
                                                                        if v18 {
                                                                            v151 := GenS20{F0: 939, F1: 255}
                                                                            v2 = append(v2, "\n\t")
                                                                            if false {
                                                                                var v152 int16 = -32768
                                                                                var v153 uint16 = 198
                                                                                var v154 uint8 = 73
                                                                                _ = v152
                                                                                _ = v153
                                                                                _ = v154
                                                                            }
                                                                            _ = v151
                                                                        } else {
                                                                            _ = len(v2)
                                                                        }
                                                                        _ = v150
                                                                    }
                                                                } else {
                                                                    if len(v2) > 0 {
                                                                        v2[0] = "foo bar"
                                                                    }
                                                                }
                                                                switch fi13 {
                                                                case -981:
                                                                    v155 := GenS1{F0: 943}
                                                                    _ = v155
                                                                case -9223372036854775808:
                                                                    var v156 uint64 = v3
                                                                    _ = v156
                                                                default:
                                                                    var v157 int8 = 89
                                                                    _ = v157
                                                                }
                                                                var v158 uint64 = 42
                                                                v1 = (3.14159 + 1.5)
                                                                _ = v158
                                                            } else {
                                                                if (!true) {
                                                                    v159 := GenS8{F0: 1}
                                                                    v1 = v1
                                                                    _ = v159
                                                                }
                                                                v160 := GenS12{F0: (v1 < (v1 - 0.0)), F1: 740}
                                                                if v18 {
                                                                    var v161 int64 = -1
                                                                    var v162 uint = 598
                                                                    switch (v15 + (("a" + "") + v15)) {
                                                                    case "a":
                                                                        v18 = (0.0 >= (v1 / (-2.25 * v1)))
                                                                    }
                                                                    _ = v161
                                                                    _ = v162
                                                                }
                                                                _ = v160
                                                            }
                                                        }
                                                        switch fi13 {
                                                        case -2147483648:
                                                            switch 42 {
                                                            case 128:
                                                                v163 := GenS21{F0: 89, F1: -210}
                                                                _ = v163
                                                            case 32767:
                                                                if ((((100 > v17) || (-224 > int32(v3))) || (!(42 >= 685))) && ((false || v18) && v18)) {
                                                                    v164 := GenS22{F0: v3, F1: 2}
                                                                    _ = v164
                                                                }
                                                            case 4294967295:
                                                                v2 = append(v2, v15)
                                                            default:
                                                                if ((32767 < 732) || (v18 || (v18 || (true || (100 >= v17))))) {
                                                                    if len(v2) > 0 {
                                                                        v2[0] = v15
                                                                    }
                                                                } else {
                                                                    v165 := GenS23{F0: fi8, F1: ""}
                                                                    var v166 int64 = -1
                                                                    var v167 int32 = 1
                                                                    switch fi13 {
                                                                    case -2147483648:
                                                                        if (137 > 1) {
                                                                            switch (^(fi13)) {
                                                                            case 4294967295:
                                                                                fmt.Println(fi8)
                                                                            case 42:
                                                                                var v168 uint32 = 4294967295
                                                                                _ = v168
                                                                            default:
                                                                                var v169 int32 = (v167 ^ -17)
                                                                                _ = v169
                                                                            }
                                                                            var v170 int16 = -481
                                                                            _ = v170
                                                                        } else {
                                                                            var v171 bool = (false || (204 > uint8(v3)))
                                                                            fmt.Println(v1, v17)
                                                                            _ = v171
                                                                        }
                                                                    default:
                                                                        var v172 uint16 = 255
                                                                        _ = v172
                                                                    }
                                                                    _ = v165
                                                                    _ = v166
                                                                    _ = v167
                                                                }
                                                            }
                                                        case 9223372036854775807:
                                                            var v173 float64 = (3.14159 / 1e10)
                                                            _ = v173
                                                        case 153:
                                                            v17 = (v17 | v17)
                                                        default:
                                                            v18 = v18
                                                        }
                                                        var v174 string = "a"
                                                        _ = v174
                                                    } else {
                                                        if true {
                                                            switch (v15 + ("\n\t" + (v15 + ("hello" + (v15 + "hello"))))) {
                                                            case "hello":
                                                                switch v15 {
                                                                case "foo bar":
                                                                    v1 = 0.0
                                                                case "a":
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                }
                                                            case "\n\t":
                                                                if v18 {
                                                                    var v175 string = ((v15 + v15) + "")
                                                                    var v176 int16 = -32768
                                                                    v175 = (((v15 + (v15 + (v175 + v15))) + (v15 + v175)) + "a")
                                                                    _ = v175
                                                                    _ = v176
                                                                } else {
                                                                    switch v15 {
                                                                    case "":
                                                                        v177 := GenS20{F0: 65535, F1: 132}
                                                                        _ = v177
                                                                    case "foo bar":
                                                                        if (v3 != (v3 / (((v3 << ((1) & 63))) | 1))) {
                                                                            fmt.Println(v3, fi13, v3)
                                                                            v178 := GenS24{F0: 1, F1: v18, F2: (v3 << ((42) & 63))}
                                                                            _ = v178
                                                                        } else {
                                                                            if false {
                                                                                v179 := GenS21{F0: 729, F1: 1}
                                                                                var v180 int16 = 1000
                                                                                _ = v179
                                                                                _ = v180
                                                                            } else {
                                                                                var v181 int64 = 1
                                                                                v15 = v15
                                                                                _ = v181
                                                                            }
                                                                        }
                                                                    case "a":
                                                                        var v182 uint64 = 1
                                                                        _ = v182
                                                                    default:
                                                                        v1 = 3.14159
                                                                        switch v15 {
                                                                        case "":
                                                                            fmt.Println(v15, fi13, v17)
                                                                        case "a":
                                                                            if ((true && (32767 >= 2)) && ((-32768 < -1) && (!(0 >= 4294967295)))) {
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = ("foo bar" + v15)
                                                                                }
                                                                            }
                                                                        default:
                                                                            v18 = v18
                                                                        }
                                                                    }
                                                                    var v183 bool = ((v15 <= (v15 + "hello")) && v18)
                                                                    fmt.Println(v17, fi13, v3)
                                                                    _ = v183
                                                                }
                                                                var v184 uint = 4294967295
                                                                _ = v184
                                                            case "a":
                                                                var v185 int8 = v17
                                                                _ = v185
                                                            }
                                                            v186 := GenS25{F0: 67}
                                                            v187 := GenS13{F0: 0}
                                                            var v188 float32 = (v1 - (v1 + (1.5 / v1)))
                                                            if false {
                                                                switch fi8 {
                                                                case -9223372036854775808:
                                                                    v3 = 18446744073709551615
                                                                }
                                                                switch fi8 {
                                                                case 170:
                                                                    var v189 int8 = (v17 - (-(v17)))
                                                                    _ = v189
                                                                case -956:
                                                                    var v190 int8 = (v17 ^ 1)
                                                                    _ = v190
                                                                }
                                                                var v191 int8 = v17
                                                                switch (v15 + (v15 + ((v15 + v15) + (v15 + ((v15 + v15) + (v15 + v15)))))) {
                                                                case "":
                                                                    fmt.Println(fi13, v15, fi13)
                                                                case "\n\t":
                                                                    v17 = v191
                                                                case "foo bar":
                                                                    if ("a" <= v15) {
                                                                        v2 = append(v2, (v15 + ("" + "")))
                                                                    }
                                                                default:
                                                                    v191 = v191
                                                                }
                                                                _ = v191
                                                            } else {
                                                                var v192 uint64 = v3
                                                                var v193 float64 = 0.0
                                                                _ = v192
                                                                _ = v193
                                                            }
                                                            v194 := GenS2{F0: 0, F1: 922, F2: ("a" + v15)}
                                                            if (!false) {
                                                                if v18 {
                                                                    v2 = append(v2, ("foo bar" + v15))
                                                                }
                                                                v1 += 1.5
                                                            } else {
                                                                var v195 float32 = v1
                                                                var v196 int64 = -79
                                                                var v197 uint8 = 42
                                                                _ = v195
                                                                _ = v196
                                                                _ = v197
                                                            }
                                                            var v198 int32 = -405
                                                            if false {
                                                                var v199 uint16 = 32767
                                                                v200 := GenS26{F0: 4294967295, F1: v15}
                                                                switch fi8 {
                                                                case 1:
                                                                    v201 := GenS10{F0: fi13, F1: 0}
                                                                    if (v18 && v18) {
                                                                        var v202 uint16 = (v199 % ((698) | 1))
                                                                        v203 := GenS20{F0: 1, F1: 0}
                                                                        _ = len(v2)
                                                                        v15 = v15
                                                                        v17 += (v17 - (v17 >> ((31) & 63)))
                                                                        _ = v202
                                                                        _ = v203
                                                                    }
                                                                    if (!(774 >= 42)) {
                                                                        fmt.Println(v188, v188)
                                                                        v204 := GenS4{F0: v198, F1: (fi8 | (fi13 - fi8)), F2: v3}
                                                                        if ((!v18) || false) {
                                                                            _ = len(v2)
                                                                            if false {
                                                                                v205 := GenS22{F0: 431, F1: (v199 >> ((65535) & 63))}
                                                                                var v206 uint16 = (^(v199))
                                                                                _ = v205
                                                                                _ = v206
                                                                            } else {
                                                                                var v207 uint8 = 0
                                                                                if (v18 && true) {
                                                                                    var v208 uint = 255
                                                                                    _ = v208
                                                                                }
                                                                                v209 := GenS11{F0: "foo bar", F1: -1}
                                                                                _ = v207
                                                                                _ = v209
                                                                            }
                                                                        } else {
                                                                            if (128 <= 2) {
                                                                                v210 := GenS18{F0: fi8}
                                                                                v198 -= -1
                                                                                v211 := GenS19{F0: -32768}
                                                                                v212 := GenS13{F0: 287}
                                                                                _ = v210
                                                                                _ = v211
                                                                                _ = v212
                                                                            }
                                                                            _ = len(v2)
                                                                            v213 := GenS27{F0: "hello", F1: (v198 | (v198 * v198)), F2: (v198 | v198)}
                                                                            switch ((v15 + v15) + "hello") {
                                                                            case "\n\t":
                                                                                var v214 int = (fi13 * fi8)
                                                                                _ = v214
                                                                            case "foo bar":
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                            case "a":
                                                                                if (!(!v18)) {
                                                                                    v215 := GenS7{F0: 42, F1: 128}
                                                                                    if true {
                                                                                        fmt.Println(v3, v15)
                                                                                        v198 = (^(v198))
                                                                                    }
                                                                                    _ = v215
                                                                                }
                                                                            }
                                                                            var v216 float32 = 1e10
                                                                            _ = v213
                                                                            _ = v216
                                                                        }
                                                                        _ = v204
                                                                    }
                                                                    _ = v201
                                                                case -526:
                                                                    v199 -= (v199 | (v199 ^ (^(v199))))
                                                                case -9223372036854775808:
                                                                    v188 = 1e10
                                                                    v217 := GenS19{F0: 36}
                                                                    _ = v217
                                                                }
                                                                _ = v199
                                                                _ = v200
                                                            }
                                                            _ = v186
                                                            _ = v187
                                                            _ = v188
                                                            _ = v194
                                                            _ = v198
                                                        }
                                                    }
                                                    fmt.Println(v18)
                                                }
                                            case 0:
                                                if len(v2) > 0 {
                                                    v2[0] = (((v15 + v15) + v15) + ((v15 + v15) + v15))
                                                }
                                                if false {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    if v18 {
                                                        switch v15 {
                                                        case "":
                                                            v18 = (42 >= 44)
                                                        case "a":
                                                            v218 := GenS18{F0: (fi8 / ((128) | 1))}
                                                            _ = v218
                                                        case "\n\t":
                                                            var v219 uint8 = 128
                                                            _ = v219
                                                        default:
                                                            v17 += int8(fi8)
                                                        }
                                                    } else {
                                                        switch (^(fi8)) {
                                                        case 66:
                                                            fmt.Println(fi13)
                                                        }
                                                    }
                                                    var v220 int8 = v17
                                                    _ = v220
                                                } else {
                                                    if v18 {
                                                        var v221 int = (fi13 + fi8)
                                                        fmt.Println(v221)
                                                        _ = v221
                                                    } else {
                                                        var v222 float32 = v1
                                                        _ = v222
                                                    }
                                                }
                                            }
                                        case "hello":
                                            if v18 {
                                                switch v15 {
                                                case "hello":
                                                    var v223 float32 = v1
                                                    _ = v223
                                                case "":
                                                    v224 := GenS7{F0: 535, F1: 127}
                                                    _ = v224
                                                }
                                            } else {
                                                if (2 >= 42) {
                                                    fmt.Println(v18, v18)
                                                    var v225 uint = 255
                                                    _ = v225
                                                } else {
                                                    switch fi13 {
                                                    case 314:
                                                        var v226 uint8 = 2
                                                        _ = v226
                                                    case 500:
                                                        v227 := GenS12{F0: ((1 == 203) && true), F1: 1}
                                                        _ = v227
                                                    case -1:
                                                        v228 := GenS28{F0: 0}
                                                        _ = v228
                                                    }
                                                    v18 = (100 >= (v17 * v17))
                                                    v15 = v15
                                                }
                                                var v229 int = fi13
                                                var v230 uint = 1
                                                v15 = ((v15 + (v15 + v15)) + v15)
                                                _ = v229
                                                _ = v230
                                            }
                                        }
                                    } else {
                                        fmt.Println(v15, fi13, v15)
                                    }
                                    v15 = v15
                                    _ = v17
                                    _ = v18
                                }
                            }
                            var v231 float64 = -2.25
                            _ = fi13
                            _ = v14
                            _ = v15
                            _ = v231
                        }
                        var v232 uint8 = 42
                        var v233 uint8 = (v232 * 255)
                        _ = v232
                        _ = v233
                    }
                    var v234 int = (^(fi8))
                    _ = v234
                }
                _ = fi8
                _ = v10
            }
        case "hello":
            switch "hello" {
            case "foo bar":
                v235 := GenS29{F0: false, F1: -641}
                _ = v235
            case "a":
                var v236 int = -1
                _ = v236
            case "\n\t":
                var v237 int16 = 32767
                _ = v237
            }
        }
    case 32767:
        fmt.Println(v3)
    default:
        fmt.Println(v1, v1)
    }
    switch 32767 {
    case 127:
        v1 += 3.14159
    case 0:
        switch "foo bar" {
        case "foo bar":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            v1 = v1
        }
    case 32768:
        v3 += v3
    }
    if true {
        var v238 uint16 = 866
        for fi239 := 0; fi239 < 5; fi239++ {
            var v240 uint = 695
            v238 = (v238 + 255)
            _ = fi239
            _ = v240
        }
        _ = v238
    } else {
        v2 = append(v2, "a")
        v241 := GenS20{F0: 2147483648, F1: 127}
        v1 = 1e10
        _ = v241
    }
    switch "hello" {
    case "hello":
        for fi242 := 0; fi242 < 17; fi242++ {
            v3 = (v3 - v3)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi242
        }
    case "":
        v2 = append(v2, "hello")
    case "\n\t":
        var v243 uint16 = 32767
        _ = v243
    default:
        fmt.Println(v1, v1)
    }
    switch 344 {
    case -2147483648:
        switch -1 {
        case 690:
            v3 += (^(v3))
        case 2147483647:
            v3 -= 4294967295
        case 32767:
            switch -61 {
            case 4294967295:
                if false {
                    for fi244 := 0; fi244 < 16; fi244++ {
                        var v245 string = "a"
                        if ((832 > 288) && false) {
                            var v246 int32 = 1
                            var v247 uint32 = 65535
                            switch v245 {
                            case "a":
                                var v248 uint16 = 42
                                _ = v248
                            case "\n\t":
                                if true {
                                    fmt.Println(v246, fi244)
                                    if false {
                                        var v249 string = (v245 + ("\n\t" + "hello"))
                                        _ = v249
                                    }
                                } else {
                                    var v250 int8 = 42
                                    _ = v250
                                }
                            }
                            _ = v246
                            _ = v247
                        }
                        _ = fi244
                        _ = v245
                    }
                    switch "a" {
                    case "\n\t":
                        v3 = v3
                    case "a":
                        if true {
                            v3 -= (v3 >> ((328) & 63))
                            v3 += 42
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        } else {
                            v1 += float32(0.1)
                            var v251 int32 = 315
                            if true {
                                v252 := GenS30{F0: 822, F1: 408}
                                var v253 int32 = genFn0()
                                _ = v252
                                _ = v253
                            } else {
                                if false {
                                    v1 = ((((float32(1.5) - v1) / v1) - v1) / (v1 * (v1 * (v1 - v1))))
                                    var v254 uint = 1
                                    v254 = 291
                                    _ = v254
                                } else {
                                    v255 := GenS17{F0: 2, F1: 16, F2: 154}
                                    fmt.Println(v1)
                                    _ = v255
                                }
                                v256 := GenS21{F0: 4294967295, F1: 42}
                                v1 = 1e10
                                _ = v256
                            }
                            _ = v251
                        }
                    default:
                        v3 -= v3
                    }
                    var v257 float64 = 3.14159
                    _ = v257
                } else {
                    _ = len(v2)
                    switch "hello" {
                    case "foo bar":
                        v1 += (1e10 * -2.25)
                    case "hello":
                        if false {
                            for fi258 := 0; fi258 < 18; fi258++ {
                                if true {
                                    switch 42 {
                                    case 32768:
                                        switch ("foo bar" + "") {
                                        case "\n\t":
                                            _ = len(v2)
                                        case "":
                                            if true {
                                                var v259 int64 = 1
                                                for fi260 := 0; fi260 < 4; fi260++ {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    var v261 int64 = (-(v259))
                                                    _ = fi260
                                                    _ = v261
                                                }
                                                _ = v259
                                            }
                                        case "foo bar":
                                            v3 = v3
                                        }
                                    }
                                    var v262 uint32 = 887
                                    v1 = (v1 + v1)
                                    _ = v262
                                } else {
                                    var v263 int = fi258
                                    if len(v2) > 0 {
                                        v2[0] = ""
                                    }
                                    var v264 int64 = -1
                                    _ = v263
                                    _ = v264
                                }
                                _ = fi258
                            }
                            v265 := GenS29{F0: false, F1: 42}
                            if len(v2) > 0 {
                                v2[0] = "foo bar"
                            }
                            _ = v265
                        } else {
                            var v266 int32 = -2147483648
                            if true {
                                var v267 uint = 4294967295
                                v1 += v1
                                for fi268 := 0; fi268 < 10; fi268++ {
                                    v3 -= 42
                                    var v269 uint8 = 0
                                    switch (fi268 >> ((uint(v269)) & 63)) {
                                    case -910:
                                        if true {
                                            var v270 float32 = v1
                                            v266 -= genFn0()
                                            _ = v270
                                        } else {
                                            var v271 float64 = 0.0
                                            var v272 uint32 = uint32(v269)
                                            switch "hello" {
                                            case "":
                                                var v273 int16 = -32768
                                                _ = v273
                                            case "foo bar":
                                                if false {
                                                    switch fi268 {
                                                    case 127:
                                                        for fi274 := 0; fi274 < 4; fi274++ {
                                                            var v275 int = fi268
                                                            v271 = (v271 - 0.0)
                                                            _ = fi274
                                                            _ = v275
                                                        }
                                                    }
                                                    var v276 int32 = v266
                                                    for fi277 := 0; fi277 < 19; fi277++ {
                                                        fmt.Println(v269, v1, v269)
                                                        if len(v2) > 0 {
                                                            v2[0] = ("\n\t" + ("\n\t" + "\n\t"))
                                                        }
                                                        _ = fi277
                                                    }
                                                    _ = v276
                                                }
                                            case "\n\t":
                                                fmt.Println(v269)
                                            default:
                                                var v278 string = ("hello" + ("" + "foo bar"))
                                                _ = v278
                                            }
                                            _ = v271
                                            _ = v272
                                        }
                                    }
                                    _ = fi268
                                    _ = v269
                                }
                                _ = v267
                            }
                            var v279 bool = (42 == -32768)
                            _ = v266
                            _ = v279
                        }
                    case "\n\t":
                        if (((!(true && false)) && (!true)) && true) {
                            v3 = 726
                        } else {
                            if (393 != 18446744073709551615) {
                                if false {
                                    if ((v1 - (v1 - 3.14159)) >= (0.1 - -2.25)) {
                                        for fi280 := 0; fi280 < 12; fi280++ {
                                            switch fi280 {
                                            case 1:
                                                v3 -= v3
                                            case 32767:
                                                fmt.Println(v3, v3, v3)
                                            case -9223372036854775808:
                                                var v281 uint16 = 2
                                                _ = v281
                                            default:
                                                for fi282 := 0; fi282 < 20; fi282++ {
                                                    fmt.Println(fi282, fi282)
                                                    _ = fi282
                                                }
                                            }
                                            var v283 int16 = 1
                                            var v284 string = "a"
                                            _ = fi280
                                            _ = v283
                                            _ = v284
                                        }
                                        var v285 uint32 = 42
                                        if len(v2) > 0 {
                                            v2[0] = ((("" + ("" + "a")) + ("foo bar" + "a")) + "\n\t")
                                        }
                                        _ = v285
                                    }
                                    switch "hello" {
                                    case "\n\t":
                                        switch "\n\t" {
                                        case "hello":
                                            v1 = 0.0
                                        case "foo bar":
                                            if true {
                                                if len(v2) > 0 {
                                                    v2[0] = "\n\t"
                                                }
                                            } else {
                                                v2 = append(v2, "a")
                                                v286 := GenS31{F0: 65535, F1: 262}
                                                _ = v286
                                            }
                                        case "a":
                                            var v287 int = -1
                                            _ = v287
                                        }
                                    default:
                                        var v288 int8 = 1
                                        _ = v288
                                    }
                                } else {
                                    v1 = (v1 - v1)
                                }
                                v289 := GenS32{F0: 0, F1: 2}
                                _ = v289
                            } else {
                                switch "a" {
                                case "hello":
                                    _ = len(v2)
                                case "a":
                                    var v290 uint8 = 0
                                    _ = v290
                                case "\n\t":
                                    var v291 bool = (int(v3) <= 607)
                                    _ = v291
                                default:
                                    fmt.Println(v3, v3)
                                }
                            }
                            switch "foo bar" {
                            case "foo bar":
                                if (!true) {
                                    var v292 bool = (!(!(false || true)))
                                    if len(v2) > 0 {
                                        v2[0] = "hello"
                                    }
                                    _ = v292
                                } else {
                                    v293 := GenS22{F0: v3, F1: 85}
                                    for fi294 := 0; fi294 < 13; fi294++ {
                                        v1 += (v1 / (((v1 / v1) * v1) * v1))
                                        v2 = append(v2, "a")
                                        v295 := GenS17{F0: 1, F1: 1, F2: v3}
                                        _ = fi294
                                        _ = v295
                                    }
                                    var v296 float32 = (v1 / (float32(((1e10 / 1e10) + 1e10)) / v1))
                                    _ = v293
                                    _ = v296
                                }
                            case "":
                                if false {
                                    var v297 int = -9223372036854775808
                                    v298 := GenS12{F0: (-94 != 127), F1: 2147483647}
                                    _ = v297
                                    _ = v298
                                } else {
                                    var v299 int = -2147483648
                                    _ = len(v2)
                                    _ = v299
                                }
                            }
                        }
                    }
                    switch "" {
                    case "hello":
                        var v300 int16 = 0
                        _ = v300
                    }
                }
            default:
                switch 9223372036854775807 {
                case -522:
                    if (!(1 != 293)) {
                        v1 += (v1 * (float32((((((0.1 * 3.14159) / 1.5) / -2.25) - 3.14159) - 1e10)) / ((0.0 / v1) - v1)))
                    }
                case -1:
                    v3 = (v3 / ((v3) | 1))
                }
            }
        }
    case 128:
        if (false && false) {
            for fi301 := 0; fi301 < 20; fi301++ {
                if (9223372036854775807 < -9223372036854775808) {
                    var v302 int32 = genFn0()
                    var v303 uint8 = 143
                    if (42 >= 1) {
                        var v304 float32 = v1
                        _ = v304
                    } else {
                        v302 += v302
                        v1 -= 1.5
                        v303 -= v303
                    }
                    _ = v302
                    _ = v303
                } else {
                    if true {
                        if len(v2) > 0 {
                            v2[0] = "\n\t"
                        }
                    }
                }
                v3 = v3
                _ = fi301
            }
        }
    case -805:
        var v305 uint32 = 783
        _ = v305
    }
    v2 = append(v2, "hello")
    if true {
        v306 := GenS33{F0: 32767, F1: 468, F2: false}
        v1 += ((0.1 / (v1 / v1)) * 0.1)
        _ = v306
    } else {
        v3 += (v3 << ((1) & 63))
        var v307 uint16 = 595
        _ = v307
    }
    if false {
        fmt.Println(v3, v3)
    }
    var v308 int16 = 1000
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v308:", v308)
    _ = v1
    _ = v2
    _ = v3
    _ = v308
}
