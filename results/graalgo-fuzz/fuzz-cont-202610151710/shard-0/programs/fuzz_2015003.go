package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int8; F1 uint16; F2 uint64 }
type GenS3 struct { F0 string; F1 int64; F2 int16 }
type GenS4 struct { F0 int }
func genFn0(p0 uint16, p1 uint) string {
    v7 := []int{4294967295, 9223372036854775807, -773}
    for fi8 := 0; fi8 < 3; fi8++ {
        if false {
            if true {
                for fi9 := 0; fi9 < 12; fi9++ {
                    if true {
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    } else {
                        v10 := GenS1{F0: 515, F1: p0, F2: 2}
                        var v11 uint64 = 860
                        if false {
                            v12 := GenS2{F0: -1, F1: (p0 >> ((p1) & 63)), F2: v11}
                            _ = v12
                        } else {
                            var v13 int32 = 798
                            _ = v13
                        }
                        _ = v10
                        _ = v11
                    }
                    _ = fi9
                }
                v7 = append(v7, fi8)
            } else {
                for fi14 := 0; fi14 < 10; fi14++ {
                    var v15 uint64 = 18446744073709551615
                    p0 += (p0 % ((42) | 1))
                    fmt.Println(p1, v15)
                    _ = fi14
                    _ = v15
                }
            }
            if len(v7) > 0 {
                v7[0] = fi8
            }
        }
        for fi16 := 0; fi16 < 19; fi16++ {
            if (!false) {
                switch 32767 {
                case 2147483647:
                    p0 = p0
                    if true {
                        p0 += p0
                        if (!(-123 < -1)) {
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                            p0 -= p0
                        } else {
                            var v17 bool = false
                            v18 := GenS1{F0: 42, F1: 42, F2: 2147483648}
                            v17 = v17
                            if (!(42 == 714)) {
                                var v19 uint16 = p0
                                if false {
                                    var v20 float32 = 3.14159
                                    p1 = p1
                                    _ = v20
                                } else {
                                    if (-819 == -690) {
                                        v19 = p0
                                    } else {
                                        var v21 uint16 = 305
                                        _ = v21
                                    }
                                }
                                var v22 uint8 = 0
                                _ = v19
                                _ = v22
                            } else {
                                switch "a" {
                                case "foo bar":
                                    var v23 int32 = -890
                                    _ = v23
                                }
                                p0 = 1
                                fmt.Println(fi8, p0, p1)
                            }
                            p0 = p0
                            _ = v17
                            _ = v18
                        }
                        if true {
                            switch fi16 {
                            case -796:
                                fmt.Println(fi16, p0)
                                v24 := GenS3{F0: (("foo bar" + ("a" + "foo bar")) + "a"), F1: -1, F2: 1}
                                var v25 float32 = 1.5
                                _ = v24
                                _ = v25
                            case -9223372036854775808:
                                var v26 int32 = 0
                                _ = v26
                            case 127:
                                if true {
                                    switch "" {
                                    case "foo bar":
                                        var v27 uint16 = (p0 << ((p1) & 63))
                                        _ = v27
                                    }
                                    p1 += 451
                                    var v28 string = "a"
                                    _ = v28
                                } else {
                                    if true {
                                        var v29 uint16 = 255
                                        var v30 int64 = -1
                                        switch "" {
                                        case "":
                                            fmt.Println(v29, fi16, v30)
                                            p0 += (p0 % (((p0 / ((uint16(p1)) | 1))) | 1))
                                            switch ("\n\t" + ("" + (("" + "foo bar") + "foo bar"))) {
                                            case "foo bar":
                                                var v31 int16 = 156
                                                _ = v31
                                            case "":
                                                if len(v7) > 0 {
                                                    v7[0] = (fi8 % ((fi16) | 1))
                                                }
                                                var v32 bool = ("\n\t" == "\n\t")
                                                _ = len(v7)
                                                _ = v32
                                            default:
                                                p1 = (^(p1))
                                            }
                                        default:
                                            switch "a" {
                                            case "hello":
                                                v33 := GenS2{F0: 98, F1: (p0 + 335), F2: 1}
                                                _ = v33
                                            }
                                        }
                                        _ = v29
                                        _ = v30
                                    } else {
                                        p0 = 432
                                        var v34 uint = 520
                                        if true {
                                            v7 = append(v7, fi8)
                                            v34 += (v34 | (^(v34)))
                                        }
                                        _ = v34
                                    }
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                    v35 := GenS3{F0: "", F1: 0, F2: -32768}
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                    _ = v35
                                }
                            default:
                                _ = len(v7)
                            }
                        } else {
                            var v36 float64 = 1.5
                            _ = v36
                        }
                    }
                    switch "hello" {
                    case "hello":
                        var v37 uint64 = 1
                        _ = v37
                    case "foo bar":
                        var v38 uint64 = 2
                        _ = v38
                    case "a":
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    }
                case 32768:
                    p1 = p1
                default:
                    v7 = append(v7, (fi8 << (((p1 ^ 690)) & 63)))
                }
            } else {
                var v39 bool = (1e10 <= (0.1 / (1.5 + (-2.25 * 1.5))))
                _ = v39
            }
            var v40 bool = (0.0 != ((0.1 + 0.0) + float32(((0.0 / 1.5) / -2.25))))
            v41 := GenS2{F0: 42, F1: p0, F2: 435}
            _ = fi16
            _ = v40
            _ = v41
        }
        v42 := GenS3{F0: "foo bar", F1: 1, F2: -1}
        _ = fi8
        _ = v42
    }
    switch 32767 {
    case 42:
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    case 633:
        var v43 int32 = 0
        _ = v43
    }
    var v44 int64 = -9223372036854775808
    p1 = (p1 | (p1 - p1))
    _ = p0
    _ = p1
    _ = v44
    _ = v7
    return ""
}
func genFn1(p0 int, p1 uint8, p2 int) int16 {
    fmt.Println(p0, p1, p1)
    v45 := GenS1{F0: -1, F1: 287, F2: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v45
    return 32767
}

func main() {
    var v1 uint32 = 781
    var v2 bool = (!(((1.5 - 0.1) <= (1.5 - 0.1)) && (74 > 42)))
    var v3 uint32 = (^(v1))
    var v4 float32 = (1e10 - 0.0)
    var v5 string = "a"
    var v6 int32 = 0
    v46 := GenS2{F0: 42, F1: 255, F2: 376}
    if (791 != 0) {
        v3 = v3
        v1 = v1
        v47 := GenS3{F0: v5, F1: 9223372036854775807, F2: genFn1(32768, 0, 32768)}
        _ = v47
    } else {
        if true {
            v1 -= v1
            var v48 int8 = 100
            _ = v48
        } else {
            var v49 int32 = 2147483647
            switch "hello" {
            case "":
                v50 := []int{554, -467, 127, 128}
                _ = v50
            case "\n\t":
                if v2 {
                    var v51 int16 = 32767
                    var v52 string = v5
                    _ = v51
                    _ = v52
                } else {
                    v53 := []int{}
                    _ = v53
                }
            case "hello":
                switch 127 {
                case 32768:
                    v54 := []int{-674}
                    _ = v54
                }
            }
            _ = v49
        }
    }
    v55 := []int{}
    if (((1.5 * 1.5) - (-2.25 * 1e10)) >= (3.14159 - -2.25)) {
        for fi56 := 0; fi56 < 14; fi56++ {
            var v57 uint16 = 32767
            _ = len(v55)
            fmt.Println(v2)
            _ = fi56
            _ = v57
        }
    } else {
        fmt.Println(v6, v4)
    }
    v58 := GenS4{F0: -9223372036854775808}
    if (false && true) {
        var v59 int8 = 100
        v3 -= (^(v3))
        v4 = v4
        _ = v59
    }
    v1 = v3
    switch -9223372036854775808 {
    case 2147483647:
        if true {
            if v2 {
                v55 = append(v55, 821)
                v1 += (v3 & (^(v3)))
                for fi60 := 0; fi60 < 20; fi60++ {
                    fmt.Println(v4, v4, v4)
                    _ = fi60
                }
            } else {
                var v61 uint = 4294967295
                _ = v61
            }
        }
    default:
        var v62 int32 = 306
        _ = v62
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v55:", len(v55), v55)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v55
    _ = v58
    _ = v6
}
