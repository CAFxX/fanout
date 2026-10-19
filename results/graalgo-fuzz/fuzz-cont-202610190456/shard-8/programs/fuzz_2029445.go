package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int64 }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int16; F1 uint32; F2 int }
type GenS6 struct { F0 int8; F1 int16; F2 int32 }
type GenS7 struct { F0 uint16; F1 int64; F2 int8 }
func genFn0() uint8 {
    var v5 uint32 = 337
    var v6 uint8 = uint8(v5)
    _ = v5
    _ = v6
    return v6
}

func main() {
    var v1 uint8 = 239
    v2 := make(map[uint32]bool)
    var v3 uint16 = 2
    v4 := []int{4294967295, 32767, 1, 32768}
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 uint = 1
        _ = fi7
        _ = v8
    }
    var v9 uint8 = v1
    fmt.Println(v3, v9)
    for fi10 := 0; fi10 < 20; fi10++ {
        var v11 int = (fi10 + int(v1))
        if len(v4) > 0 {
            v4[0] = v11
        }
        v12 := GenS1{F0: v3, F1: ("hello" + ""), F2: 782}
        _ = fi10
        _ = v11
        _ = v12
    }
    if ((-9223372036854775808 != -1) || (!(0 != 2))) {
        v3 = (v3 % ((1) | 1))
    }
    if ((("\n\t" + "") + "") < ("hello" + "a")) {
        var v13 int8 = 42
        v14 := GenS2{F0: ((42 < 908) && ((true || true) && (v3 < uint16(v1)))), F1: 0}
        switch 9223372036854775807 {
        case 810:
            var v15 int32 = 3
            _ = v15
        case 9223372036854775807:
            v3 = (v3 << ((2) & 63))
        case -2147483648:
            if (-1 <= (^(v13))) {
                var v16 int = -1
                var v17 int32 = 42
                _ = v16
                _ = v17
            } else {
                v3 -= 42
                var v18 bool = true
                _ = v18
            }
        default:
            var v19 float32 = 0.1
            _ = v19
        }
        _ = v13
        _ = v14
    } else {
        var v20 float64 = float64((0.1 / -2.25))
        fmt.Println(v3, v3, v20)
        for fi21 := 0; fi21 < 2; fi21++ {
            v3 = 504
            v3 -= (^(v3))
            for fi22 := 0; fi22 < 2; fi22++ {
                if false {
                    var v23 int = 128
                    _ = v23
                } else {
                    v24 := GenS1{F0: (v3 >> ((770) & 63)), F1: "foo bar", F2: 819}
                    _ = v24
                }
                var v25 uint = 4294967295
                var v26 float64 = v20
                _ = fi22
                _ = v25
                _ = v26
            }
            _ = fi21
        }
        _ = v20
    }
    if true {
        if true {
            var v27 int64 = -199
            v27 = -216
            v1 = (^(v1))
            _ = v27
        } else {
            for fi28 := 0; fi28 < 3; fi28++ {
                fmt.Println(fi28, fi28, v9)
                v3 += (v3 >> ((0) & 63))
                v3 += (^(v3))
                _ = fi28
            }
            var v29 float32 = 1e10
            var v30 float64 = float64(v29)
            _ = v29
            _ = v30
        }
        for fi31 := 0; fi31 < 17; fi31++ {
            switch fi31 {
            case 252:
                if true {
                    switch "foo bar" {
                    case "":
                        for fi32 := 0; fi32 < 20; fi32++ {
                            fmt.Println(fi32, fi31, v9)
                            _ = fi32
                        }
                    case "a":
                        for fi33 := 0; fi33 < 13; fi33++ {
                            v4 = append(v4, (fi31 / ((fi31) | 1)))
                            _ = fi33
                        }
                    default:
                        var v34 uint64 = 42
                        _ = v34
                    }
                }
            case 32767:
                if (!true) {
                    if false {
                        v3 = (v3 << ((0) & 63))
                        var v35 int16 = 1
                        if true {
                            if ((v1 * (v1 | v9)) < v1) {
                                fmt.Println(v35, v35)
                                if true {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    for fi36 := 0; fi36 < 19; fi36++ {
                                        v3 = v3
                                        _ = fi36
                                    }
                                    var v37 int16 = -473
                                    _ = v37
                                }
                                for fi38 := 0; fi38 < 8; fi38++ {
                                    v1 += 42
                                    _ = fi38
                                }
                            }
                        } else {
                            var v39 int32 = -343
                            _ = v39
                        }
                        _ = v35
                    } else {
                        var v40 string = "a"
                        var v41 int16 = 663
                        v3 -= v3
                        _ = v40
                        _ = v41
                    }
                } else {
                    var v42 bool = true
                    var v43 uint32 = 918
                    if false {
                        v43 -= v43
                        var v44 uint32 = (v43 % (((v43 >> ((2) & 63))) | 1))
                        _ = v44
                    }
                    _ = v42
                    _ = v43
                }
            default:
                v4 = append(v4, (fi31 + fi31))
            }
            if (!true) {
                v1 = 127
                var v45 uint64 = 253
                fmt.Println(v1, v1, v45)
                _ = v45
            } else {
                v46 := GenS2{F0: false, F1: 42}
                fmt.Println(v3, v9, v3)
                var v47 uint16 = v3
                _ = v46
                _ = v47
            }
            v48 := GenS3{F0: 2}
            _ = fi31
            _ = v48
        }
        v49 := GenS3{F0: 637}
        _ = v49
    }
    v50 := GenS3{F0: 150}
    var v51 float32 = 0.0
    for fi52 := 0; fi52 < 5; fi52++ {
        for fi53 := 0; fi53 < 6; fi53++ {
            if false {
                switch int(v3) {
                case 0:
                    v3 = v3
                case 152:
                    var v54 uint32 = 2147483648
                    _ = v54
                }
                var v55 int32 = 2147483647
                var v56 uint16 = (^(v3))
                _ = v55
                _ = v56
            } else {
                v57 := GenS2{F0: (530 > 2), F1: 883}
                v58 := GenS3{F0: 329}
                _ = v57
                _ = v58
            }
            var v59 int = fi53
            v9 -= genFn0()
            var v60 float32 = ((float32(float64(float32((1.5 - float64(((((v51 - v51) - v51) / 0.1) + 3.14159)))))) / 1e10) + (v51 * v51))
            fmt.Println(v59, v59, v51)
            _ = fi53
            _ = v59
            _ = v60
        }
        _ = fi52
    }
    fmt.Println(v9, v1, v9)
    v1 = 69
    switch -1 {
    case -1:
        for fi61 := 0; fi61 < 7; fi61++ {
            var v62 uint64 = 1
            _ = fi61
            _ = v62
        }
    default:
        if false {
            if true {
                if (!false) {
                    v63 := v2[2147483647]
                    _ = len(v2)
                    _ = v63
                }
                if (v3 != 32767) {
                    if ((v3 >> ((65535) & 63)) < (v3 / ((255) | 1))) {
                        if false {
                            var v64 float64 = 1e10
                            v65 := GenS4{F0: 18446744073709551615}
                            for fi66 := 0; fi66 < 4; fi66++ {
                                v4 = append(v4, 0)
                                _ = fi66
                            }
                            _ = v64
                            _ = v65
                        } else {
                            switch "\n\t" {
                            case "":
                                var v67 uint8 = v1
                                _ = v67
                            default:
                                var v68 float64 = ((0.0 * -2.25) * ((3.14159 - float64(v51)) + (0.0 - 0.0)))
                                _ = v68
                            }
                            switch 128 {
                            case 127:
                                var v69 float64 = 1e10
                                _ = v69
                            case 42:
                                v70 := GenS1{F0: (v3 % ((v3) | 1)), F1: "hello", F2: -960}
                                _ = v70
                            case 765:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            default:
                                v51 += 0.0
                            }
                            if (42 != (v1 & v1)) {
                                var v71 uint64 = 18446744073709551615
                                v72 := GenS1{F0: (^(v3)), F1: "", F2: 42}
                                v9 = 0
                                _ = v71
                                _ = v72
                            } else {
                                for fi73 := 0; fi73 < 18; fi73++ {
                                    v74 := GenS5{F0: 239, F1: 2147483648, F2: fi73}
                                    if ((true && (0.0 > v51)) && false) {
                                        for fi75 := 0; fi75 < 11; fi75++ {
                                            v3 += (v3 | 65535)
                                            var v76 int32 = 0
                                            _ = fi75
                                            _ = v76
                                        }
                                        var v77 string = ""
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        _ = v77
                                    }
                                    _ = fi73
                                    _ = v74
                                }
                                for fi78 := 0; fi78 < 18; fi78++ {
                                    var v79 float64 = ((1e10 - 0.0) * 0.1)
                                    _ = fi78
                                    _ = v79
                                }
                            }
                        }
                        if false {
                            if true {
                                v80 := GenS6{F0: 0, F1: 0, F2: 1}
                                for fi81 := 0; fi81 < 20; fi81++ {
                                    var v82 uint64 = 18446744073709551615
                                    _ = fi81
                                    _ = v82
                                }
                                _ = v80
                            }
                            switch (("hello" + (("\n\t" + ("foo bar" + "hello")) + (("foo bar" + "hello") + "a"))) + (("\n\t" + "\n\t") + "foo bar")) {
                            case "foo bar":
                                var v83 uint = 4294967295
                                _ = v83
                            case "":
                                var v84 float32 = v51
                                _ = v84
                            case "a":
                                v85 := GenS6{F0: -1, F1: 680, F2: 0}
                                _ = v85
                            }
                            for fi86 := 0; fi86 < 6; fi86++ {
                                if false {
                                    var v87 uint64 = 4294967295
                                    v51 = ((-2.25 + v51) + (float32(float64(v51)) - v51))
                                    fmt.Println(v9, v9)
                                    _ = v87
                                }
                                for fi88 := 0; fi88 < 16; fi88++ {
                                    switch "hello" {
                                    case "foo bar":
                                        var v89 int = (fi88 - (fi88 ^ (fi86 & 411)))
                                        _ = v89
                                    case "hello":
                                        v9 = v9
                                    case "":
                                        v90 := GenS1{F0: (v3 >> ((997) & 63)), F1: "a", F2: 160}
                                        _ = v90
                                    }
                                    v91 := GenS4{F0: 1}
                                    _ = fi88
                                    _ = v91
                                }
                                _ = fi86
                            }
                        }
                    }
                    v3 += (v3 & (v3 - v3))
                } else {
                    var v92 uint8 = (v1 ^ genFn0())
                    if (871 != -9223372036854775808) {
                        v93 := GenS7{F0: (^(v3)), F1: -1, F2: -1}
                        _ = v93
                    }
                    _ = v92
                }
            } else {
                delete(v2, 0)
                fmt.Println(v1)
            }
            v51 -= 1.5
            v3 = 0
        }
    }
    var v94 int32 = -715
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v51:", v51)
    fmt.Println("v9:", v9)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v51
    _ = v9
    _ = v94
}
