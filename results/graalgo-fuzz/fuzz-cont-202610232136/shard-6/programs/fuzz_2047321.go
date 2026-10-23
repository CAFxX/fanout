package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 string; F1 int32; F2 string }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint32; F1 int; F2 uint16 }
type GenS6 struct { F0 bool; F1 int; F2 uint8 }
type GenS7 struct { F0 uint16; F1 int16 }
type GenS8 struct { F0 int; F1 string }
type GenS9 struct { F0 string; F1 uint16; F2 uint16 }
func genFn0(p0 uint, p1 int64, p2 int64) int {
    var v5 float64 = float64(3.14159)
    v6 := GenS1{F0: 0}
    p2 -= (p1 & -1)
    var v7 uint16 = 587
    var v8 uint8 = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 128
}
func genFn1(p0 int64, p1 uint) uint16 {
    fmt.Println(p1, p0)
    v9 := GenS2{F0: 1, F1: "\n\t"}
    _ = p0
    _ = p1
    _ = v9
    return 325
}

func main() {
    var v1 uint = 42
    var v2 float32 = 3.14159
    var v3 uint64 = 0
    v4 := []string{"foo bar", "a"}
    v2 = -2.25
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (v3 >= (v3 >> (((v1 % ((0) | 1))) & 63))) {
        for fi10 := 0; fi10 < 13; fi10++ {
            v1 -= 1
            v11 := GenS2{F0: -370, F1: "\n\t"}
            _ = fi10
            _ = v11
        }
        if ((v3 << (((v1 ^ (v1 << ((v1) & 63)))) & 63)) >= 4294967295) {
            switch 4294967295 {
            case 542:
                v12 := GenS2{F0: -73, F1: "foo bar"}
                _ = v12
            default:
                v1 = v1
            }
            var v13 int32 = -185
            v14 := GenS1{F0: 2147483648}
            _ = v13
            _ = v14
        } else {
            v1 -= 18446744073709551615
            var v15 uint16 = genFn1(1, v1)
            v2 = 1.5
            _ = v15
        }
        _ = len(v4)
    } else {
        fmt.Println(v2, v2)
        if false {
            for fi16 := 0; fi16 < 10; fi16++ {
                var v17 int16 = 1
                _ = fi16
                _ = v17
            }
            for fi18 := 0; fi18 < 8; fi18++ {
                var v19 float64 = 1e10
                if ((v19 / v19) != (1.5 + 1e10)) {
                    v2 += 1.5
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v20 uint64 = v3
                    _ = v20
                } else {
                    v4 = append(v4, "a")
                    for fi21 := 0; fi21 < 15; fi21++ {
                        fmt.Println(v19)
                        _ = len(v4)
                        _ = fi21
                    }
                    if true {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v22 int16 = 857
                        _ = v22
                    } else {
                        for fi23 := 0; fi23 < 1; fi23++ {
                            if (32767 >= genFn1(42, v1)) {
                                v19 -= (((v19 / 1.5) / (v19 / v19)) * (1e10 + v19))
                                var v24 bool = false
                                _ = v24
                            }
                            var v25 bool = (false && (38 >= 42))
                            _ = fi23
                            _ = v25
                        }
                    }
                }
                if false {
                    v1 = (v1 << (((v1 | 2)) & 63))
                    for fi26 := 0; fi26 < 9; fi26++ {
                        switch ("a" + "foo bar") {
                        case "":
                            var v27 uint8 = 15
                            _ = v27
                        case "\n\t":
                            v28 := GenS3{F0: "", F1: -911, F2: "a"}
                            _ = v28
                        }
                        var v29 uint32 = 2
                        _ = fi26
                        _ = v29
                    }
                    for fi30 := 0; fi30 < 12; fi30++ {
                        if (!(1 > 127)) {
                            if (455 != 2) {
                                if true {
                                    v31 := GenS1{F0: 4294967295}
                                    if (4294967295 > 65535) {
                                        switch "a" {
                                        case "foo bar":
                                            v3 = 2
                                        case "hello":
                                            _ = len(v4)
                                        case "":
                                            var v32 float64 = v19
                                            _ = v32
                                        }
                                    } else {
                                        v33 := GenS4{F0: genFn0((v1 & 2), 9223372036854775807, -678)}
                                        var v34 int64 = 42
                                        v19 += 0.0
                                        _ = v33
                                        _ = v34
                                    }
                                    if ((v2 / v2) > (1e10 * ((v2 - float32(v19)) - v2))) {
                                        var v35 int64 = 42
                                        _ = v35
                                    }
                                    _ = v31
                                } else {
                                    switch "foo bar" {
                                    case "":
                                        if true {
                                            v36 := GenS5{F0: 42, F1: genFn0(v1, 165, 42), F2: 1}
                                            var v37 uint = 408
                                            v1 = v1
                                            if false {
                                                if true {
                                                    _ = len(v4)
                                                    if (0 != 0) {
                                                        v4 = append(v4, ("foo bar" + "a"))
                                                    } else {
                                                        v37 = v1
                                                        v3 += v3
                                                        fmt.Println(v2, fi18)
                                                    }
                                                }
                                            }
                                            _ = v36
                                            _ = v37
                                        } else {
                                            fmt.Println(fi18, v19)
                                            fmt.Println(fi30, fi30, fi18)
                                            var v38 uint32 = 806
                                            fmt.Println(v1, v19)
                                            _ = v38
                                        }
                                    }
                                    switch fi18 {
                                    case 396:
                                        if true {
                                            _ = len(v4)
                                            if true {
                                                v2 = (-2.25 + v2)
                                                var v39 uint16 = 566
                                                _ = v39
                                            }
                                            if (fi18 > genFn0((v1 + v1), -439, int64(v3))) {
                                                if true {
                                                    var v40 uint8 = 255
                                                    fmt.Println(v19)
                                                    var v41 uint32 = 2147483648
                                                    fmt.Println(v1)
                                                    _ = v40
                                                    _ = v41
                                                } else {
                                                    v42 := GenS5{F0: 2, F1: (^(fi30)), F2: 484}
                                                    v3 += v3
                                                    _ = v42
                                                }
                                                var v43 uint32 = 2147483647
                                                var v44 int = genFn0((v1 >> (((^(v1))) & 63)), 9223372036854775807, 844)
                                                _ = v43
                                                _ = v44
                                            } else {
                                                fmt.Println(fi18, v19, v1)
                                            }
                                        } else {
                                            if (false && (598 >= 18446744073709551615)) {
                                                switch fi18 {
                                                case 9223372036854775807:
                                                    var v45 float32 = v2
                                                    _ = v45
                                                default:
                                                    fmt.Println(v19, v1)
                                                }
                                                if (false && ((!true) && false)) {
                                                    var v46 float64 = float64(0.1)
                                                    if (383 != 492) {
                                                        switch fi30 {
                                                        case 4294967295:
                                                            v47 := GenS6{F0: (!false), F1: genFn0((^(v1)), 9223372036854775807, -1), F2: uint8(v3)}
                                                            _ = v47
                                                        case -725:
                                                            var v48 float32 = ((v2 + v2) * (float32(v46) * v2))
                                                            _ = v48
                                                        }
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                    }
                                                    _ = v46
                                                }
                                                v2 = 1e10
                                            }
                                            if ((^(v1)) > (v1 % (((v1 * v1)) | 1))) {
                                                var v49 int64 = 648
                                                _ = v49
                                            }
                                        }
                                    default:
                                        v4 = append(v4, "hello")
                                    }
                                    _ = len(v4)
                                }
                                v1 = v1
                                v50 := GenS3{F0: ((("hello" + "\n\t") + ("\n\t" + ("a" + "\n\t"))) + ("hello" + "\n\t")), F1: 0, F2: ("" + "\n\t")}
                                switch "foo bar" {
                                case "\n\t":
                                    var v51 bool = true
                                    _ = v51
                                case "foo bar":
                                    v52 := GenS5{F0: 1, F1: (fi18 ^ fi30), F2: 32767}
                                    _ = v52
                                default:
                                    v53 := GenS7{F0: 65535, F1: -181}
                                    if true {
                                        v1 -= v1
                                    } else {
                                        var v54 int = fi18
                                        v55 := GenS8{F0: (^(fi18)), F1: "foo bar"}
                                        _ = v54
                                        _ = v55
                                    }
                                    v4 = append(v4, ("a" + "foo bar"))
                                    var v56 int8 = 17
                                    switch 9223372036854775807 {
                                    case 447:
                                        var v57 uint16 = 304
                                        _ = v57
                                    case 127:
                                        v3 += 730
                                    }
                                    _ = v53
                                    _ = v56
                                }
                                var v58 uint64 = 42
                                _ = v50
                                _ = v58
                            }
                            v1 = (v1 / (((v1 << ((v1) & 63))) | 1))
                            if true {
                                if len(v4) > 0 {
                                    v4[0] = "foo bar"
                                }
                            }
                            switch 1 {
                            case 42:
                                var v59 bool = false
                                _ = v59
                            }
                        }
                        var v60 uint64 = v3
                        if len(v4) > 0 {
                            v4[0] = "hello"
                        }
                        _ = fi30
                        _ = v60
                    }
                }
                _ = fi18
                _ = v19
            }
        } else {
            if true {
                if len(v4) > 0 {
                    v4[0] = ("a" + "\n\t")
                }
                var v61 uint8 = 131
                var v62 uint32 = 2
                _ = v61
                _ = v62
            }
        }
    }
    v63 := GenS1{F0: 0}
    if len(v4) > 0 {
        v4[0] = ("" + ("" + ("foo bar" + "hello")))
    }
    var v64 uint32 = 65535
    fmt.Println(v64, v3, v1)
    fmt.Println(v2, v3)
    if (false || ((1 >= -717) || true)) {
        for fi65 := 0; fi65 < 10; fi65++ {
            for fi66 := 0; fi66 < 2; fi66++ {
                v64 -= 2147483647
                v67 := GenS8{F0: (fi66 % (((fi65 % (((fi66 >> (((v1 ^ 4294967295)) & 63))) | 1))) | 1)), F1: (("a" + "") + ("" + "a"))}
                if (((v2 + v2) / ((-2.25 / ((-2.25 - 1e10) + v2)) - float32(((1e10 * float64(v2)) / 3.14159)))) != -2.25) {
                    var v68 bool = true
                    switch fi66 {
                    case 0:
                        var v69 int32 = -910
                        _ = v69
                    case 701:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    _ = v68
                } else {
                    var v70 float64 = 1.5
                    var v71 uint32 = 65535
                    _ = v70
                    _ = v71
                }
                _ = fi66
                _ = v67
            }
            if len(v4) > 0 {
                v4[0] = ""
            }
            if len(v4) > 0 {
                v4[0] = ("\n\t" + "foo bar")
            }
            _ = fi65
        }
        var v72 float64 = -2.25
        _ = v72
    }
    v3 -= v3
    if ((!true) && true) {
        if (true || (!false)) {
            v73 := GenS4{F0: int(v1)}
            _ = v73
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v74 := GenS3{F0: "foo bar", F1: 0, F2: "foo bar"}
            _ = v74
        }
        if true {
            var v75 int64 = 42
            v76 := GenS5{F0: (v64 & (^(v64))), F1: 2147483647, F2: 336}
            _ = v75
            _ = v76
        } else {
            v77 := GenS4{F0: 4294967295}
            _ = v77
        }
    }
    switch -2147483648 {
    case 9223372036854775807:
        _ = len(v4)
    case 2147483647:
        if (-2.25 < (3.14159 * (3.14159 * 1e10))) {
            v78 := GenS7{F0: 255, F1: -32768}
            v79 := GenS9{F0: (("foo bar" + ("a" + "foo bar")) + "hello"), F1: 2, F2: 0}
            for fi80 := 0; fi80 < 15; fi80++ {
                if false {
                    switch (("" + "foo bar") + "hello") {
                    case "a":
                        v2 -= (v2 + 3.14159)
                    case "hello":
                        var v81 int16 = -154
                        _ = v81
                    }
                    var v82 uint = v1
                    var v83 float64 = ((-2.25 / ((1e10 / 1e10) / float64(1e10))) * (((1.5 + (-2.25 + 3.14159)) - 1e10) * float64(float32(-2.25))))
                    _ = v82
                    _ = v83
                }
                if len(v4) > 0 {
                    v4[0] = "hello"
                }
                for fi84 := 0; fi84 < 19; fi84++ {
                    v4 = append(v4, "")
                    v64 -= (v64 << (((v1 / ((4294967295) | 1))) & 63))
                    _ = fi84
                }
                _ = fi80
            }
            _ = v78
            _ = v79
        }
    case -573:
        var v85 float32 = v2
        _ = v85
    default:
        v3 = 428
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v63
    _ = v64
}
