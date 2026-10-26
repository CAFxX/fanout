package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 uint32; F2 bool }
type GenS3 struct { F0 int; F1 int32; F2 uint }
type GenS4 struct { F0 int64; F1 uint8; F2 uint64 }
type GenS5 struct { F0 uint; F1 int8 }
type GenS6 struct { F0 int64; F1 uint16; F2 uint16 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 string; F1 bool; F2 uint32 }
type GenS9 struct { F0 int16; F1 string }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 int; F1 uint32; F2 uint32 }
type GenS12 struct { F0 uint16 }
func genFn0(p0 uint16, p1 uint32) int32 {
    p0 += (p0 * p0)
    p1 = (p1 >> ((42) & 63))
    _ = p0
    _ = p1
    return int32(p0)
}

func main() {
    var v1 uint64 = 1
    var v2 bool = false
    v3 := []string{"foo bar", "hello", ((("\n\t" + "a") + "hello") + "\n\t")}
    var v4 uint16 = 65535
    if len(v3) > 0 {
        v3[0] = (("" + "hello") + (("a" + "hello") + ""))
    }
    v5 := GenS1{F0: 2}
    if (18446744073709551615 < 2) {
        if v2 {
            switch 101 {
            case 127:
                _ = len(v3)
            case 32767:
                _ = len(v3)
            }
            var v6 int = -1
            switch v6 {
            case -1:
                v3 = append(v3, ("foo bar" + ("foo bar" + "a")))
            case -503:
                v7 := GenS1{F0: 42}
                _ = v7
            default:
                if (v2 && (!(1000 != 655))) {
                    switch int(v1) {
                    case 128:
                        for fi8 := 0; fi8 < 16; fi8++ {
                            v1 += v1
                            if v2 {
                                for fi9 := 0; fi9 < 3; fi9++ {
                                    v10 := GenS1{F0: 2}
                                    if true {
                                        v2 = true
                                        if (fi9 != fi9) {
                                            if false {
                                                if (!(255 != 127)) {
                                                    _ = len(v3)
                                                } else {
                                                    var v11 uint32 = 992
                                                    _ = v11
                                                }
                                            } else {
                                                v12 := GenS2{F0: 98, F1: 4294967295, F2: ((!(-26 != 127)) && (genFn0((v4 + 42), 2) == genFn0((v4 * v4), 0)))}
                                                _ = v12
                                            }
                                            v4 = v4
                                            v13 := GenS3{F0: v6, F1: -2147483648, F2: 4294967295}
                                            v6 = -9223372036854775808
                                            switch fi8 {
                                            case 32767:
                                                switch fi8 {
                                                case -497:
                                                    v14 := GenS4{F0: -1, F1: 89, F2: (v1 * (v1 | 2))}
                                                    _ = v14
                                                case -653:
                                                    v1 += 2
                                                }
                                            case -1:
                                                var v15 int64 = 0
                                                _ = v15
                                            default:
                                                if true {
                                                    if (v2 && (!(4294967295 >= v1))) {
                                                        var v16 uint = 42
                                                        _ = v16
                                                    }
                                                    if v2 {
                                                        var v17 uint8 = 255
                                                        var v18 bool = (42 <= 9223372036854775807)
                                                        if v2 {
                                                            var v19 uint8 = v17
                                                            switch ("\n\t" + "hello") {
                                                            case "a":
                                                                v18 = (!(168 <= 2))
                                                            case "":
                                                                if (875 >= 633) {
                                                                    var v20 int = fi8
                                                                    _ = v20
                                                                }
                                                                if v18 {
                                                                    v4 = 0
                                                                }
                                                                v21 := GenS5{F0: 405, F1: -1}
                                                                _ = v21
                                                            default:
                                                                v4 = v4
                                                            }
                                                            _ = v19
                                                        }
                                                        fmt.Println(v6, v1)
                                                        _ = v17
                                                        _ = v18
                                                    }
                                                } else {
                                                    var v22 uint = 65535
                                                    _ = v22
                                                }
                                                switch "hello" {
                                                case "foo bar":
                                                    if v2 {
                                                        var v23 uint32 = 983
                                                        if (95 < uint(v23)) {
                                                            v6 -= (fi8 ^ (^(fi8)))
                                                            _ = len(v3)
                                                        } else {
                                                            v1 += v1
                                                        }
                                                        _ = v23
                                                    } else {
                                                        switch ("" + "hello") {
                                                        case "foo bar":
                                                            v24 := GenS4{F0: 483, F1: 127, F2: (^(v1))}
                                                            _ = v24
                                                        case "hello":
                                                            if v2 {
                                                                switch "" {
                                                                case "":
                                                                    var v25 int32 = -2147483648
                                                                    _ = v25
                                                                case "a":
                                                                    if v2 {
                                                                        v2 = (v2 && (42 <= 2))
                                                                        v2 = (fi8 >= 1)
                                                                        switch (v6 & (-(fi9))) {
                                                                        case 684:
                                                                            v3 = append(v3, "a")
                                                                        }
                                                                        var v26 uint = 620
                                                                        _ = v26
                                                                    } else {
                                                                        switch (v6 & (v6 << ((2) & 63))) {
                                                                        case 32768:
                                                                            _ = len(v3)
                                                                        case -2147483648:
                                                                            var v27 uint64 = v1
                                                                            _ = v27
                                                                        default:
                                                                            v4 = 2
                                                                        }
                                                                        switch "" {
                                                                        case "a":
                                                                            v3 = append(v3, ("" + ("a" + ("foo bar" + ""))))
                                                                        case "\n\t":
                                                                            v6 = 824
                                                                        default:
                                                                            var v28 uint32 = 4294967295
                                                                            _ = v28
                                                                        }
                                                                        var v29 float32 = 0.0
                                                                        _ = v29
                                                                    }
                                                                case "hello":
                                                                    v2 = ((v4 * (^(v4))) <= v4)
                                                                }
                                                                if false {
                                                                    switch "a" {
                                                                    case "":
                                                                        if false {
                                                                            var v30 int16 = int16(v1)
                                                                            _ = v30
                                                                        } else {
                                                                            fmt.Println(v4, v2, fi8)
                                                                        }
                                                                    case "a":
                                                                        switch fi8 {
                                                                        case 356:
                                                                            v3 = append(v3, ("\n\t" + "\n\t"))
                                                                            fmt.Println(v4)
                                                                            v31 := GenS6{F0: 794, F1: (^(v4)), F2: 149}
                                                                            _ = v31
                                                                        default:
                                                                            v1 -= 18446744073709551615
                                                                        }
                                                                    case "\n\t":
                                                                        switch ("a" + "\n\t") {
                                                                        case "foo bar":
                                                                            var v32 int8 = 1
                                                                            _ = v32
                                                                        default:
                                                                            _ = len(v3)
                                                                        }
                                                                        fmt.Println(v4, fi9)
                                                                    default:
                                                                        v1 += v1
                                                                    }
                                                                }
                                                                switch ((("hello" + "hello") + ("\n\t" + "hello")) + "foo bar") {
                                                                case "":
                                                                    if false {
                                                                        v1 = v1
                                                                        var v33 int32 = -857
                                                                        _ = v33
                                                                    } else {
                                                                        var v34 uint32 = 42
                                                                        v35 := GenS5{F0: 1, F1: 100}
                                                                        v36 := GenS3{F0: (fi8 ^ -2147483648), F1: 353, F2: 293}
                                                                        _ = v34
                                                                        _ = v35
                                                                        _ = v36
                                                                    }
                                                                case "a":
                                                                    var v37 float32 = (0.1 + 0.0)
                                                                    _ = v37
                                                                default:
                                                                    v4 = (v4 & (v4 | v4))
                                                                }
                                                            }
                                                        default:
                                                            if true {
                                                                v38 := GenS4{F0: 42, F1: 127, F2: v1}
                                                                v3 = append(v3, "a")
                                                                _ = v38
                                                            } else {
                                                                var v39 int16 = 1
                                                                fmt.Println(v39)
                                                                _ = v39
                                                            }
                                                            v40 := GenS7{F0: 106}
                                                            _ = v40
                                                        }
                                                        var v41 uint = 18446744073709551615
                                                        _ = v41
                                                    }
                                                default:
                                                    v4 = 48
                                                }
                                                v42 := GenS8{F0: ("hello" + "a"), F1: (!(fi8 > 128)), F2: 1}
                                                _ = v42
                                            }
                                            _ = v13
                                        } else {
                                            if ((1.5 > 3.14159) && v2) {
                                                var v43 int8 = int8(v4)
                                                _ = v43
                                            }
                                            if v2 {
                                                v3 = append(v3, ("\n\t" + "a"))
                                                if (v4 <= 688) {
                                                    v2 = v2
                                                }
                                                v3 = append(v3, "\n\t")
                                            }
                                        }
                                        if (723 > 65535) {
                                            var v44 int32 = -1
                                            _ = v44
                                        } else {
                                            var v45 int8 = -128
                                            v6 += fi8
                                            _ = v45
                                        }
                                    } else {
                                        if (v2 || v2) {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        } else {
                                            var v46 uint8 = 35
                                            var v47 float64 = (1e10 / (float64((3.14159 - 1e10)) / 3.14159))
                                            v48 := GenS1{F0: 0}
                                            v1 = v1
                                            _ = v46
                                            _ = v47
                                            _ = v48
                                        }
                                        v6 += -640
                                        var v49 uint64 = v1
                                        _ = v49
                                    }
                                    _ = fi9
                                    _ = v10
                                }
                                if ((v2 || true) || v2) {
                                    for fi50 := 0; fi50 < 13; fi50++ {
                                        if v2 {
                                            fmt.Println(v6, v6)
                                            v6 = v6
                                            if true {
                                                v51 := GenS6{F0: 0, F1: (v4 % (((v4 * (v4 | v4))) | 1)), F2: v4}
                                                var v52 float32 = 1e10
                                                _ = v51
                                                _ = v52
                                            } else {
                                                v53 := GenS3{F0: (fi50 >> ((uint(v1)) & 63)), F1: -790, F2: 719}
                                                v4 = (v4 % ((65535) | 1))
                                                _ = v53
                                            }
                                        } else {
                                            fmt.Println(v6, v1)
                                            if v2 {
                                                switch v6 {
                                                case 0:
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                default:
                                                    if len(v3) > 0 {
                                                        v3[0] = ""
                                                    }
                                                }
                                                var v54 bool = ((false || (v6 <= v6)) || (!(("" + ("\n\t" + "hello")) > ("\n\t" + "a"))))
                                                _ = v54
                                            }
                                            if ((fi50 >> ((867) & 63)) != (^(fi50))) {
                                                v6 -= fi8
                                                fmt.Println(v2, v6)
                                                switch "foo bar" {
                                                case "hello":
                                                    v55 := GenS9{F0: 32767, F1: "a"}
                                                    _ = v55
                                                case "\n\t":
                                                    var v56 int16 = 42
                                                    _ = v56
                                                }
                                                if ((fi50 / (((fi50 % (((fi50 & 42)) | 1))) | 1)) >= v6) {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    v57 := GenS10{F0: (v1 << ((878) & 63))}
                                                    v2 = (((!false) || (0.1 == 0.1)) || v2)
                                                    _ = v57
                                                }
                                            }
                                        }
                                        if v2 {
                                            v58 := GenS9{F0: -594, F1: "foo bar"}
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            fmt.Println(v6, v4, v2)
                                            _ = v58
                                        } else {
                                            v3 = append(v3, (("a" + ("foo bar" + "a")) + "a"))
                                            var v59 int64 = -9223372036854775808
                                            _ = v59
                                        }
                                        fmt.Println(v6)
                                        _ = fi50
                                    }
                                }
                            }
                            v2 = (v2 || false)
                            _ = fi8
                        }
                    default:
                        v60 := GenS9{F0: 32767, F1: (("a" + "\n\t") + ("hello" + ("a" + "")))}
                        _ = v60
                    }
                    _ = len(v3)
                }
            }
            _ = v6
        } else {
            fmt.Println(v4)
        }
        fmt.Println(v1, v1)
    } else {
        v61 := GenS11{F0: 108, F1: 65535, F2: 78}
        v62 := GenS12{F0: (^(v4))}
        for fi63 := 0; fi63 < 16; fi63++ {
            if v2 {
                _ = len(v3)
                v1 += v1
            } else {
                var v64 uint64 = v1
                v64 = v64
                _ = v64
            }
            var v65 int32 = 42
            _ = fi63
            _ = v65
        }
        _ = v61
        _ = v62
    }
    fmt.Println(v4)
    v1 -= v1
    var v66 int32 = 2147483647
    v2 = false
    v67 := GenS7{F0: uint8(v1)}
    for fi68 := 0; fi68 < 11; fi68++ {
        v2 = v2
        fmt.Println(v1, v66, v1)
        _ = fi68
    }
    for fi69 := 0; fi69 < 5; fi69++ {
        v2 = (((1.5 < float32(1.5)) && true) || (!false))
        _ = fi69
    }
    for fi70 := 0; fi70 < 9; fi70++ {
        var v71 uint8 = 1
        v2 = (!(!(127 == 42)))
        fmt.Println(fi70, v2)
        _ = fi70
        _ = v71
    }
    _ = len(v3)
    switch -1 {
    case -150:
        for fi72 := 0; fi72 < 9; fi72++ {
            var v73 int = 328
            for fi74 := 0; fi74 < 4; fi74++ {
                fmt.Println(v73)
                _ = fi74
            }
            for fi75 := 0; fi75 < 18; fi75++ {
                var v76 int64 = -1
                v3 = append(v3, "foo bar")
                _ = fi75
                _ = v76
            }
            _ = fi72
            _ = v73
        }
    default:
        v4 = v4
    }
    var v77 int = 504
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v66:", v66)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v66
    _ = v67
    _ = v77
}
