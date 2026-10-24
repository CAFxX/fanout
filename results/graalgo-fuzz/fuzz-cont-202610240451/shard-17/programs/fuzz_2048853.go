package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int8; F1 int16; F2 uint32 }
type GenS6 struct { F0 int64; F1 int16 }
type GenS7 struct { F0 int64; F1 int64 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int16; F1 int8 }

func main() {
    v1 := make(map[string]bool)
    var v2 bool = (true || false)
    var v3 uint16 = 369
    v4 := []uint32{16, 1, 2147483647, 42}
    v5 := make(map[int]uint32)
    var v6 uint = 511
    v2 = true
    v6 = uint(v3)
    v6 += 251
    var v7 uint = 18446744073709551615
    v8 := GenS1{F0: v6, F1: 2147483647, F2: (^(v6))}
    var v9 uint64 = 513
    for fi10 := 0; fi10 < 11; fi10++ {
        v11 := GenS2{F0: v9}
        _ = fi10
        _ = v11
    }
    switch "a" {
    case "\n\t":
        var v12 float64 = 1.5
        _ = v12
    default:
        var v13 int64 = 0
        _ = v13
    }
    if ((false && ((v3 <= (v3 & (^(v3)))) && false)) && true) {
        var v14 int16 = -89
        v3 = (^(v3))
        _ = v14
    } else {
        v7 = uint(v9)
        for fi15 := 0; fi15 < 10; fi15++ {
            for fi16 := 0; fi16 < 9; fi16++ {
                v17 := v1[("foo bar" + "hello")]
                switch "" {
                case "foo bar":
                    if (128 > 128) {
                        _ = len(v4)
                    }
                case "hello":
                    var v18 float32 = 3.14159
                    _ = v18
                case "":
                    var v19 int = fi16
                    _ = v19
                default:
                    v7 -= (v6 >> ((v6) & 63))
                }
                _ = len(v4)
                _ = fi16
                _ = v17
            }
            v9 = (v9 + v9)
            switch fi15 {
            case 2147483647:
                v2 = (2147483647 != 42)
            }
            _ = fi15
        }
        if v2 {
            v20 := GenS1{F0: (v7 >> (((v6 - uint(v9))) & 63)), F1: int32(v9), F2: v6}
            if (v9 == v9) {
                for fi21 := 0; fi21 < 2; fi21++ {
                    var v22 bool = (uint8(v9) < uint8(v6))
                    _ = fi21
                    _ = v22
                }
                for fi23 := 0; fi23 < 19; fi23++ {
                    v6 = (v6 ^ (v7 | 42))
                    _ = fi23
                }
            } else {
                v2 = (false || v2)
                switch -588 {
                case -9223372036854775808:
                    var v24 uint32 = 0
                    _ = v24
                case 32767:
                    v25 := GenS3{F0: ("foo bar" + "\n\t")}
                    _ = v25
                default:
                    switch 128 {
                    case 42:
                        v26 := GenS4{F0: 0}
                        _ = v26
                    case -2147483648:
                        var v27 uint16 = v3
                        _ = v27
                    case -406:
                        var v28 bool = v2
                        _ = v28
                    default:
                        v29 := GenS5{F0: 44, F1: 1000, F2: 2147483648}
                        _ = v29
                    }
                }
                v3 += v3
            }
            _ = v20
        }
    }
    v7 = 255
    for fi30 := 0; fi30 < 3; fi30++ {
        for fi31 := 0; fi31 < 17; fi31++ {
            v5[fi30] = 2147483647
            if (false && (fi31 != (fi30 * 2147483647))) {
                switch (fi31 << ((v7) & 63)) {
                case 32768:
                    fmt.Println(v2)
                case 2147483647:
                    var v32 int32 = 0
                    v6 = (v6 >> ((v6) & 63))
                    _ = v32
                }
            } else {
                v3 = (v3 % ((190) | 1))
                var v33 string = "\n\t"
                v34 := GenS6{F0: 87, F1: int16(v9)}
                v35 := GenS2{F0: v9}
                switch v33 {
                case "hello":
                    v7 += v7
                case "foo bar":
                    var v36 uint32 = 2147483647
                    _ = v36
                case "a":
                    switch v33 {
                    case "a":
                        v7 += (v6 - v7)
                    default:
                        var v37 int8 = 100
                        switch 127 {
                        case 127:
                            v3 += (v3 - (v3 << (((v6 / ((462) | 1))) & 63)))
                        }
                        if (!((fi30 == (fi31 - (fi30 ^ int(v37)))) && (42 > int32(fi31)))) {
                            if len(v4) > 0 {
                                v4[0] = 865
                            }
                            fmt.Println(v2, v9)
                        }
                        _ = v37
                    }
                    if (v2 && (v2 || ((v2 && (!(1e10 < (0.1 + (1.5 / 1.0))))) && (1e10 != 3.14159)))) {
                        if (v2 && (1 < 9223372036854775807)) {
                            switch fi30 {
                            case -1:
                                if v2 {
                                    v2 = (v2 || ((v6 & (v6 << ((v6) & 63))) >= (v6 / ((uint(v3)) | 1))))
                                    v9 = v9
                                    fmt.Println(v2)
                                    v2 = (v2 || (-9223372036854775808 != 0))
                                    if (!((v33 + "a") == "foo bar")) {
                                        var v38 float64 = (1.5 + 1.5)
                                        _ = v38
                                    }
                                    fmt.Println(fi31, v7, v6)
                                } else {
                                    if (!v2) {
                                        var v39 int16 = 1
                                        _ = v39
                                    } else {
                                        v7 += 255
                                    }
                                    if (-124 > -128) {
                                        switch v33 {
                                        case "":
                                            if (-1 <= 0) {
                                                v2 = v2
                                                var v40 int8 = 42
                                                _ = v40
                                            }
                                        }
                                        if v2 {
                                            v6 = (v6 * v6)
                                        }
                                    } else {
                                        v6 = (v6 / ((v6) | 1))
                                    }
                                    switch 127 {
                                    case 42:
                                        v4 = append(v4, 642)
                                    case 9223372036854775807:
                                        v41 := GenS3{F0: v33}
                                        _ = v41
                                    default:
                                        if true {
                                            _ = len(v4)
                                            fmt.Println(v7)
                                            if false {
                                                fmt.Println(v3)
                                                v9 -= 42
                                            } else {
                                                v42 := GenS6{F0: 0, F1: -32768}
                                                fmt.Println(v7)
                                                var v43 uint = 169
                                                _ = v42
                                                _ = v43
                                            }
                                        } else {
                                            var v44 int64 = 42
                                            v6 -= (^(v7))
                                            _ = v44
                                        }
                                    }
                                }
                                var v45 string = (v33 + "hello")
                                switch (v33 + ("a" + "foo bar")) {
                                case "":
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                case "hello":
                                    if false {
                                        var v46 bool = (true && (!(100 == -77)))
                                        if v2 {
                                            _ = len(v4)
                                        } else {
                                            var v47 float64 = 0.0
                                            _ = v47
                                        }
                                        fmt.Println(v6, v7, fi31)
                                        _ = v46
                                    } else {
                                        if v2 {
                                            if v2 {
                                                var v48 float64 = ((1.5 * 1e10) + -2.25)
                                                var v49 int8 = 35
                                                fmt.Println(v33, v7)
                                                _ = v48
                                                _ = v49
                                            } else {
                                                v2 = (!v2)
                                                var v50 bool = (-2147483648 < fi30)
                                                _ = v50
                                            }
                                        } else {
                                            var v51 uint32 = 497
                                            v45 = ("hello" + "a")
                                            var v52 float32 = float32(0.1)
                                            _ = v51
                                            _ = v52
                                        }
                                    }
                                case "\n\t":
                                    fmt.Println(v45, v6)
                                }
                                v53 := GenS1{F0: 2, F1: 833, F2: v7}
                                fmt.Println(fi30)
                                _ = v45
                                _ = v53
                            case 128:
                                if v2 {
                                    v9 += v9
                                    v33 = ("hello" + v33)
                                    var v54 int64 = 1
                                    v55 := GenS7{F0: v54, F1: (-(v54))}
                                    v2 = v2
                                    _ = v54
                                    _ = v55
                                } else {
                                    var v56 int8 = 42
                                    if false {
                                        v33 = v33
                                    }
                                    _ = v56
                                }
                                var v57 uint8 = 0
                                var v58 string = ((v33 + v33) + "")
                                var v59 uint = v7
                                _ = v57
                                _ = v58
                                _ = v59
                            }
                        }
                        v9 += v9
                        var v60 uint = 658
                        _ = v60
                    } else {
                        if (128 == 255) {
                            v61 := GenS8{F0: v2}
                            _ = v61
                        }
                    }
                default:
                    if true {
                        v7 -= v6
                        var v62 int32 = -1
                        fmt.Println(v7, fi31, v3)
                        _ = v62
                    } else {
                        var v63 int16 = 32767
                        delete(v5, (^(fi30)))
                        v64 := GenS7{F0: 778, F1: 1}
                        _ = v63
                        _ = v64
                    }
                }
                _ = v33
                _ = v34
                _ = v35
            }
            v65 := GenS1{F0: v6, F1: 42, F2: 294}
            if (-456 >= 1) {
                if (v2 || true) {
                    if v2 {
                        switch 501 {
                        case 580:
                            var v66 int64 = 9223372036854775807
                            _ = v66
                        case 127:
                            if len(v4) > 0 {
                                v4[0] = 161
                            }
                        case 4294967295:
                            if (109 > 873) {
                                _ = len(v5)
                            }
                        }
                    }
                }
                v5[fi31] = 65535
                switch "\n\t" {
                case "foo bar":
                    v6 -= (v7 & (v6 % ((v7) | 1)))
                default:
                    _ = len(v5)
                    v5[(fi31 / ((32767) | 1))] = 0
                }
                var v67 int8 = 0
                var v68 uint = v6
                var v69 uint = v6
                _ = v67
                _ = v68
                _ = v69
            } else {
                v70 := GenS6{F0: 518, F1: int16(fi31)}
                v6 -= (v7 & (^(v6)))
                v71 := GenS9{F0: 853, F1: 14}
                _ = v70
                _ = v71
            }
            _ = fi31
            _ = v65
        }
        v72 := GenS3{F0: "a"}
        var v73 int8 = 100
        _ = fi30
        _ = v72
        _ = v73
    }
    v3 -= 340
    v3 += v3
    var v74 string = ""
    var v75 int32 = 42
    var v76 float64 = (-2.25 + 1.5)
    v77 := GenS3{F0: (("hello" + ((v74 + "a") + "")) + (v74 + (v74 + v74)))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v74
    _ = v75
    _ = v76
    _ = v77
    _ = v8
    _ = v9
}
