package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64; F1 string; F2 int32 }
type GenS5 struct { F0 uint16; F1 int8 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int32; F1 uint16 }
type GenS8 struct { F0 uint8; F1 uint32; F2 string }
type GenS9 struct { F0 uint16 }
func genFn0(p0 int, p1 uint64, p2 uint16) int {
    v7 := []int{int(p1), p0, p0, p0}
    var v8 uint64 = p1
    v9 := GenS1{F0: 2}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return (p0 & p0)
}
func genFn1(p0 uint8, p1 string) int16 {
    var v10 uint = 270
    if (!false) {
        p0 -= p0
    }
    if (229 > 32767) {
        if false {
            for fi11 := 0; fi11 < 7; fi11++ {
                p1 = p1
                var v12 bool = ((("a" + ("hello" + p1)) + p1) > "a")
                _ = fi11
                _ = v12
            }
            p1 = "\n\t"
        } else {
            v10 = (v10 << (((v10 | v10)) & 63))
            var v13 float64 = (((-2.25 - -2.25) - 0.1) + 1e10)
            p0 = 43
            _ = v13
        }
    }
    _ = p0
    _ = p1
    _ = v10
    return 1
}

func main() {
    var v1 uint32 = 0
    v2 := make(map[uint32]uint32)
    v3 := []string{}
    var v4 uint16 = 1
    var v5 int = -9223372036854775808
    v6 := make(map[string]bool)
    if (true && ((861 < 42) || true)) {
        if true {
            v1 += 4294967295
        } else {
            switch "" {
            case "a":
                v1 = 1
            case "hello":
                for fi14 := 0; fi14 < 16; fi14++ {
                    v15 := GenS1{F0: 255}
                    _ = fi14
                    _ = v15
                }
            case "\n\t":
                v4 = v4
            default:
                if true {
                    v4 = (v4 ^ v4)
                    if (!false) {
                        v16 := v2[(^(v1))]
                        if ("" > ("hello" + "")) {
                            var v17 uint16 = v4
                            for fi18 := 0; fi18 < 5; fi18++ {
                                v19 := GenS2{F0: v4, F1: 513}
                                var v20 int32 = 1
                                v20 -= 0
                                _ = fi18
                                _ = v19
                                _ = v20
                            }
                            fmt.Println(v16, v5, v17)
                            _ = v17
                        } else {
                            if false {
                                v1 = v1
                            } else {
                                if true {
                                    for fi21 := 0; fi21 < 14; fi21++ {
                                        if (42 <= 255) {
                                            fmt.Println(v1, v1, v1)
                                            var v22 int32 = 1
                                            for fi23 := 0; fi23 < 19; fi23++ {
                                                v24 := GenS3{F0: 405}
                                                v1 = v1
                                                v25 := GenS4{F0: -1, F1: "foo bar", F2: -2147483648}
                                                _ = fi23
                                                _ = v24
                                                _ = v25
                                            }
                                            _ = v22
                                        } else {
                                            v16 -= (^(v16))
                                            switch "foo bar" {
                                            case "":
                                                v1 = (^(v1))
                                            case "hello":
                                                fmt.Println(v16, v5, v4)
                                            case "a":
                                                for fi26 := 0; fi26 < 6; fi26++ {
                                                    var v27 bool = ((255 == 255) && (-53 <= -1))
                                                    _ = fi26
                                                    _ = v27
                                                }
                                            default:
                                                var v28 uint32 = uint32(v4)
                                                _ = v28
                                            }
                                        }
                                        _ = fi21
                                    }
                                    var v29 string = "\n\t"
                                    switch genFn0(int(v16), 4294967295, v4) {
                                    case -2147483648:
                                        var v30 float64 = ((0.0 * (1.5 + 1.5)) / 1.0)
                                        _ = v30
                                    case -1:
                                        fmt.Println(v16, v16)
                                    default:
                                        v31 := GenS5{F0: v4, F1: 0}
                                        _ = v31
                                    }
                                    _ = v29
                                } else {
                                    for fi32 := 0; fi32 < 14; fi32++ {
                                        v5 -= (v5 - (v5 / (((v5 | (v5 & genFn0(genFn0(genFn0(127, uint64(v16), v4), 1, (v4 | v4)), 562, (v4 << ((1) & 63)))))) | 1)))
                                        _ = fi32
                                    }
                                    v33 := GenS4{F0: -1, F1: "hello", F2: -303}
                                    v34 := GenS6{F0: 1}
                                    _ = v33
                                    _ = v34
                                }
                                var v35 string = "foo bar"
                                _ = v35
                            }
                            if false {
                                for fi36 := 0; fi36 < 15; fi36++ {
                                    v2[(v1 ^ (^(v16)))] = (^(v1))
                                    _ = fi36
                                }
                                for fi37 := 0; fi37 < 15; fi37++ {
                                    v38 := GenS2{F0: (v4 + (v4 << ((255) & 63))), F1: 18446744073709551615}
                                    _ = fi37
                                    _ = v38
                                }
                            } else {
                                switch "foo bar" {
                                case "hello":
                                    var v39 uint = 4294967295
                                    _ = v39
                                case "foo bar":
                                    v16 = v16
                                case "a":
                                    v1 = 0
                                default:
                                    var v40 uint8 = 2
                                    _ = v40
                                }
                                switch ("hello" + "foo bar") {
                                case "hello":
                                    v1 -= v1
                                }
                            }
                            switch genFn0((v5 % (((v5 - (v5 % ((v5) | 1)))) | 1)), 197, v4) {
                            case -9223372036854775808:
                                v41 := v2[(v16 + v1)]
                                _ = v41
                            case 81:
                                if false {
                                    if ((1e10 != 0.0) && true) {
                                        for fi42 := 0; fi42 < 1; fi42++ {
                                            fmt.Println(v4, v4, v1)
                                            var v43 float64 = 1.5
                                            _ = fi42
                                            _ = v43
                                        }
                                    }
                                } else {
                                    var v44 int8 = 100
                                    for fi45 := 0; fi45 < 17; fi45++ {
                                        if len(v3) > 0 {
                                            v3[0] = "\n\t"
                                        }
                                        v1 += v1
                                        _ = fi45
                                    }
                                    var v46 uint8 = 0
                                    _ = v44
                                    _ = v46
                                }
                            default:
                                var v47 int64 = 584
                                _ = v47
                            }
                        }
                        _ = len(v2)
                        _ = v16
                    } else {
                        var v48 uint8 = 0
                        _ = v48
                    }
                    delete(v6, "a")
                }
            }
            v49 := GenS7{F0: 0, F1: v4}
            var v50 bool = (((-1 < 42) && false) || false)
            _ = v49
            _ = v50
        }
        var v51 int32 = 42
        _ = v51
    }
    v5 = v5
    v4 = v4
    v1 -= (v1 / ((0) | 1))
    if (((v4 ^ v4) != v4) || false) {
        var v52 uint64 = 18446744073709551615
        for fi53 := 0; fi53 < 5; fi53++ {
            v5 -= fi53
            if true {
                if (!true) {
                    v54 := GenS1{F0: 66}
                    fmt.Println(v4, v52, fi53)
                    _ = v54
                }
                v5 += fi53
                var v55 int16 = 42
                _ = v55
            }
            _ = fi53
        }
        var v56 uint = 42
        _ = v52
        _ = v56
    } else {
        switch "" {
        case "foo bar":
            var v57 int8 = 0
            _ = v57
        case "\n\t":
            var v58 uint = 42
            _ = v58
        case "":
            _ = len(v3)
        }
        fmt.Println(v5, v1, v4)
    }
    for fi59 := 0; fi59 < 11; fi59++ {
        v4 -= v4
        _ = len(v3)
        v60 := GenS2{F0: v4, F1: 2}
        _ = fi59
        _ = v60
    }
    v1 = 2147483648
    for fi61 := 0; fi61 < 7; fi61++ {
        v62 := GenS4{F0: int64(v1), F1: ("foo bar" + "\n\t"), F2: -2147483648}
        v5 += (v5 << ((414) & 63))
        var v63 float32 = ((0.0 / -2.25) + 3.14159)
        _ = fi61
        _ = v62
        _ = v63
    }
    v1 = 0
    _ = len(v2)
    v5 = genFn0((-(v5)), 4294967295, v4)
    v1 = (^(v1))
    var v64 float64 = 3.14159
    if (0 != -73) {
        v65 := GenS8{F0: 255, F1: (v1 >> ((25) & 63)), F2: "\n\t"}
        for fi66 := 0; fi66 < 5; fi66++ {
            v4 += (v4 - v4)
            if true {
                v64 += 1.5
                v67 := GenS4{F0: 614, F1: "foo bar", F2: -2147483648}
                fmt.Println(v1, fi66, fi66)
                _ = v67
            } else {
                v3 = append(v3, "a")
                var v68 string = (("hello" + ("\n\t" + "\n\t")) + ("\n\t" + ("\n\t" + ("foo bar" + "\n\t"))))
                _ = v68
            }
            var v69 float64 = 1.5
            _ = fi66
            _ = v69
        }
        _ = v65
    } else {
        var v70 int8 = 0
        if (true && false) {
            var v71 uint16 = v4
            for fi72 := 0; fi72 < 8; fi72++ {
                v71 += 921
                var v73 uint8 = 3
                _ = fi72
                _ = v73
            }
            v1 += 143
            _ = v71
        } else {
            v70 = v70
            if (0 <= (v5 << ((687) & 63))) {
                if false {
                    var v74 uint8 = 2
                    v75 := GenS9{F0: uint16(v74)}
                    _ = v74
                    _ = v75
                } else {
                    v76 := GenS5{F0: v4, F1: v70}
                    if len(v3) > 0 {
                        v3[0] = ("foo bar" + ("" + "foo bar"))
                    }
                    if (!false) {
                        fmt.Println(v4, v5)
                        fmt.Println(v64)
                        fmt.Println(v64)
                    } else {
                        var v77 float32 = (float32(v64) / 1.0)
                        if true {
                            fmt.Println(v5, v5)
                            if false {
                                var v78 uint = 18446744073709551615
                                _ = v78
                            }
                        } else {
                            v77 = 3.14159
                            fmt.Println(v70)
                            _ = len(v2)
                        }
                        for fi79 := 0; fi79 < 3; fi79++ {
                            fmt.Println(v1)
                            for fi80 := 0; fi80 < 20; fi80++ {
                                var v81 uint32 = 686
                                _ = fi80
                                _ = v81
                            }
                            _ = fi79
                        }
                        _ = v77
                    }
                    _ = v76
                }
                v5 -= v5
                fmt.Println(v4, v5)
            } else {
                var v82 uint = 2
                fmt.Println(v82, v4, v5)
                fmt.Println(v5, v1, v1)
                _ = v82
            }
            var v83 uint8 = 0
            _ = v83
        }
        _ = v70
    }
    if true {
        if false {
            v5 = v5
        }
        var v84 int16 = 1
        _ = v84
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v64
}
