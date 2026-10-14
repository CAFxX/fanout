package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int }
type GenS2 struct { F0 string; F1 int32; F2 int64 }
type GenS3 struct { F0 uint32; F1 uint8; F2 int16 }
type GenS4 struct { F0 int64; F1 string; F2 uint }
type GenS5 struct { F0 uint32; F1 int }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int8; F1 int; F2 int8 }
func genFn0(p0 int, p1 int64) uint64 {
    for fi5 := 0; fi5 < 3; fi5++ {
        if (true || true) {
            var v6 int8 = 0
            for fi7 := 0; fi7 < 1; fi7++ {
                p1 -= p1
                _ = fi7
            }
            _ = v6
        } else {
            for fi8 := 0; fi8 < 12; fi8++ {
                p0 = (p0 & (fi5 + (fi8 & (fi5 / ((p0) | 1)))))
                fmt.Println(fi8, fi8, p1)
                _ = fi8
            }
        }
        _ = fi5
    }
    v9 := []int{(p0 ^ (p0 >> ((18446744073709551615) & 63)))}
    var v10 uint8 = 196
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 4294967295
}

func main() {
    var v1 uint64 = 663
    var v2 uint = 331
    var v3 int8 = 13
    v4 := make(map[uint32]uint32)
    if (127 != 51) {
        if false {
            var v11 uint32 = 236
            for fi12 := 0; fi12 < 15; fi12++ {
                for fi13 := 0; fi13 < 1; fi13++ {
                    v14 := GenS1{F0: 4294967295, F1: ((^(fi12)) <= (fi12 + (fi13 * (fi12 ^ fi12)))), F2: fi13}
                    v15 := GenS2{F0: "hello", F1: -1, F2: 9223372036854775807}
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                if (false && (!("" != ("foo bar" + "a")))) {
                    v16 := GenS3{F0: (v11 | 1), F1: 42, F2: -749}
                    if true {
                        for fi17 := 0; fi17 < 4; fi17++ {
                            v11 = (^(v11))
                            var v18 int64 = 42
                            var v19 string = ("foo bar" + "\n\t")
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                        if true {
                            fmt.Println(fi12, v3, v11)
                            for fi20 := 0; fi20 < 19; fi20++ {
                                v1 -= (v1 & genFn0(fi12, 0))
                                if false {
                                    switch (fi12 | 128) {
                                    case 252:
                                        if false {
                                            if (385 != -1) {
                                                var v21 float64 = (((((float64(0.1) / ((3.14159 + (-2.25 / 1.0)) / (1.5 * 0.1))) + 3.14159) + 1e10) * (float64(0.0) + 1.5)) + 0.0)
                                                if ((-(v3)) >= v3) {
                                                    switch fi20 {
                                                    case 4294967295:
                                                        var v22 uint32 = (v11 >> (((^(v2))) & 63))
                                                        _ = v22
                                                    case -350:
                                                        v1 = v1
                                                    default:
                                                        var v23 int16 = 459
                                                        _ = v23
                                                    }
                                                    v24 := []int{(fi20 * -2147483648), (fi12 / ((fi12) | 1)), -626, fi20}
                                                    v25 := GenS2{F0: ("\n\t" + "hello"), F1: -2147483648, F2: 42}
                                                    _ = v24
                                                    _ = v25
                                                }
                                                if true {
                                                    fmt.Println(v1, v2, fi20)
                                                    v11 += 917
                                                    v26 := GenS1{F0: v2, F1: (true && false), F2: 128}
                                                    _ = v26
                                                }
                                                _ = v21
                                            }
                                            v3 -= -104
                                        } else {
                                            var v27 uint16 = 255
                                            _ = v27
                                        }
                                    case 127:
                                        fmt.Println(fi20)
                                    default:
                                        v28 := GenS3{F0: v11, F1: 127, F2: 0}
                                        _ = v28
                                    }
                                    _ = len(v4)
                                }
                                v1 += genFn0(128, 42)
                                switch 9223372036854775807 {
                                case -39:
                                    switch "hello" {
                                    case "\n\t":
                                        if false {
                                            var v29 int = (fi12 % ((fi20) | 1))
                                            if ("foo bar" < (("hello" + "hello") + ("a" + "a"))) {
                                                v30 := GenS1{F0: v2, F1: (false || ((v1 / ((v1) | 1)) > 260)), F2: fi20}
                                                v4[(^(v11))] = v11
                                                var v31 uint32 = v11
                                                _ = v30
                                                _ = v31
                                            }
                                            if false {
                                                if (!true) {
                                                    fmt.Println(fi20, v2, v3)
                                                    var v32 string = "a"
                                                    v3 = 21
                                                    _ = v32
                                                }
                                                fmt.Println(v2)
                                            } else {
                                                v2 += uint(v1)
                                            }
                                            _ = v29
                                        }
                                    case "foo bar":
                                        var v33 uint16 = uint16(v11)
                                        _ = v33
                                    default:
                                        var v34 uint8 = 191
                                        _ = v34
                                    }
                                    if false {
                                        var v35 uint64 = v1
                                        var v36 uint = (v2 >> (((v2 * v2)) & 63))
                                        _ = v35
                                        _ = v36
                                    } else {
                                        v2 = 65535
                                        if (!true) {
                                            v11 = uint32(v2)
                                            v1 = 18446744073709551615
                                        }
                                        v11 = v11
                                        v11 -= v11
                                        v37 := GenS4{F0: -9223372036854775808, F1: "hello", F2: (v2 % ((v2) | 1))}
                                        _ = v37
                                    }
                                    v1 = v1
                                    v3 -= int8(fi20)
                                case -483:
                                    if true {
                                        var v38 float32 = (3.14159 / 1.5)
                                        v1 -= v1
                                        fmt.Println(v1, v1)
                                        v1 = v1
                                        v3 = v3
                                        _ = v38
                                    }
                                case 9223372036854775807:
                                    v39 := []int{fi12, fi12, -2147483648}
                                    v39 = append(v39, -845)
                                    v40 := GenS3{F0: (v11 * 937), F1: 2, F2: 1}
                                    _ = v39
                                    _ = v40
                                }
                                var v41 uint8 = uint8(v11)
                                _ = fi20
                                _ = v41
                            }
                            if (true && true) {
                                var v42 float32 = ((1.5 + ((((0.1 / -2.25) - float32(1.5)) / 1.0) - (float32((((3.14159 * 1e10) / 1.5) / (3.14159 - (1e10 * (1e10 - 1.5))))) - -2.25))) + 1e10)
                                _ = v42
                            } else {
                                for fi43 := 0; fi43 < 13; fi43++ {
                                    if (1e10 != 0.0) {
                                        switch "" {
                                        case "a":
                                            v44 := GenS3{F0: 934, F1: 227, F2: 833}
                                            _ = v44
                                        case "foo bar":
                                            var v45 uint64 = (v1 % ((v1) | 1))
                                            _ = v45
                                        case "\n\t":
                                            if (!true) {
                                                v2 -= (v2 >> ((uint(v11)) & 63))
                                            } else {
                                                v3 = (v3 - (v3 >> ((v2) & 63)))
                                                if (!true) {
                                                    v11 = (v11 >> ((18446744073709551615) & 63))
                                                    var v46 float32 = 1e10
                                                    v47 := GenS4{F0: 42, F1: "\n\t", F2: v2}
                                                    _ = v46
                                                    _ = v47
                                                }
                                                var v48 string = "\n\t"
                                                _ = v48
                                            }
                                        }
                                    } else {
                                        if true {
                                            if true {
                                                var v49 int16 = 205
                                                if (42 != 164) {
                                                    switch "" {
                                                    case "foo bar":
                                                        var v50 int8 = v3
                                                        _ = v50
                                                    case "\n\t":
                                                        if (false || false) {
                                                            fmt.Println(v11, v3, fi12)
                                                        }
                                                    default:
                                                        if false {
                                                            if false {
                                                                var v51 int16 = 0
                                                                var v52 uint64 = genFn0((^(fi43)), 1)
                                                                _ = v51
                                                                _ = v52
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    v49 = v49
                                                    if false {
                                                        v2 = (^(v2))
                                                        v53 := v4[971]
                                                        _ = v53
                                                    } else {
                                                        v4[545] = 2147483648
                                                        fmt.Println(v49, fi12)
                                                    }
                                                    fmt.Println(v11)
                                                }
                                                delete(v4, (v11 / ((v11) | 1)))
                                                _ = v49
                                            }
                                            fmt.Println(v1)
                                        } else {
                                            v54 := []int{0, (fi12 - 32768)}
                                            _ = v54
                                        }
                                    }
                                    _ = fi43
                                }
                                v55 := []int{(fi12 - 4294967295)}
                                _ = v55
                            }
                        }
                    } else {
                        v56 := GenS5{F0: 2147483647, F1: (fi12 >> (((v2 ^ v2)) & 63))}
                        var v57 string = "foo bar"
                        _ = v56
                        _ = v57
                    }
                    _ = v16
                } else {
                    var v58 uint32 = 483
                    v59 := GenS6{F0: false}
                    var v60 float32 = ((1e10 * 1.5) / 1e10)
                    _ = v58
                    _ = v59
                    _ = v60
                }
                var v61 uint = (v2 & 1)
                _ = fi12
                _ = v61
            }
            switch 128 {
            case 0:
                switch 32767 {
                case 32768:
                    v2 = v2
                case -9223372036854775808:
                    var v62 uint64 = v1
                    _ = v62
                case -191:
                    v63 := GenS7{F0: -1, F1: -9223372036854775808, F2: (v3 << (((v2 - (v2 / ((v2) | 1)))) & 63))}
                    _ = v63
                default:
                    v64 := []int{9223372036854775807, 602, 0}
                    _ = v64
                }
            }
            _ = v11
        } else {
            var v65 int32 = 2147483647
            _ = v65
        }
        var v66 uint64 = (v1 >> (((^(v2))) & 63))
        _ = v66
    } else {
        for fi67 := 0; fi67 < 5; fi67++ {
            for fi68 := 0; fi68 < 10; fi68++ {
                if ((true || (0 == -938)) && (415 == 2147483647)) {
                    v2 = (v2 << ((v2) & 63))
                } else {
                    v3 -= 1
                }
                delete(v4, 827)
                _ = fi68
            }
            _ = fi67
        }
        for fi69 := 0; fi69 < 14; fi69++ {
            if true {
                switch ("foo bar" + "\n\t") {
                case "a":
                    var v70 float32 = (-2.25 + 1.5)
                    _ = v70
                case "hello":
                    fmt.Println(fi69, v2)
                default:
                    for fi71 := 0; fi71 < 1; fi71++ {
                        var v72 int64 = 42
                        var v73 string = ("hello" + ((("a" + "") + "") + (("hello" + ("" + "hello")) + "hello")))
                        _ = fi71
                        _ = v72
                        _ = v73
                    }
                }
                var v74 int = (-(fi69))
                _ = v74
            } else {
                switch "" {
                case "hello":
                    v75 := GenS6{F0: (1000 == -985)}
                    _ = v75
                case "a":
                    var v76 uint16 = 203
                    _ = v76
                }
            }
            v77 := GenS7{F0: -1, F1: (fi69 + 263), F2: int8(fi69)}
            _ = fi69
            _ = v77
        }
    }
    v1 += (v1 % ((615) | 1))
    fmt.Println(v1, v2, v3)
    v3 += -56
    for fi78 := 0; fi78 < 16; fi78++ {
        v79 := []int{}
        v1 = (v1 << (((v2 & 18446744073709551615)) & 63))
        _ = fi78
        _ = v79
    }
    var v80 uint8 = 2
    if false {
        v81 := []int{847}
        for fi82 := 0; fi82 < 17; fi82++ {
            if (!((v1 % (((v1 | 2)) | 1)) > 426)) {
                switch "foo bar" {
                case "":
                    var v83 int = fi82
                    _ = v83
                case "hello":
                    fmt.Println(v80, v3)
                case "\n\t":
                    v4[549] = 4294967295
                }
            }
            for fi84 := 0; fi84 < 5; fi84++ {
                var v85 uint8 = (v80 * 0)
                var v86 uint16 = uint16(v2)
                _ = fi84
                _ = v85
                _ = v86
            }
            _ = fi82
        }
        _ = v81
    }
    v3 = -128
    v87 := []int{0}
    var v88 int8 = (v3 - (v3 << (((v2 | v2)) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v80:", v80)
    fmt.Println("slicelen v87:", len(v87), v87)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v80
    _ = v87
    _ = v88
}
