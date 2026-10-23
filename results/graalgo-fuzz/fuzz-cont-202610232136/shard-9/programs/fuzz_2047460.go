package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 int64; F1 uint32 }
type GenS5 struct { F0 uint32; F1 uint16; F2 uint8 }
type GenS6 struct { F0 int64; F1 int16; F2 int64 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint16; F1 uint64 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 int8; F1 uint64 }
func genFn0(p0 int, p1 uint, p2 uint64) int8 {
    var v7 uint8 = 42
    var v8 bool = true
    v9 := GenS1{F0: 42, F1: 240, F2: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return -128
}

func main() {
    var v1 bool = (!false)
    v2 := []string{"hello", "foo bar", ""}
    v3 := []uint32{2}
    var v4 int8 = 100
    v5 := make(map[int]uint32)
    var v6 uint32 = 173
    v10 := v5[-170]
    if v1 {
        if (160 < -2147483648) {
            if false {
                var v11 int64 = 931
                _ = v11
            }
            if (1 == -414) {
                v4 += (v4 % ((100) | 1))
                var v12 uint64 = 466
                if v1 {
                    var v13 int = -288
                    switch "a" {
                    case "hello":
                        switch v13 {
                        case 765:
                            delete(v5, 32767)
                        case -215:
                            if true {
                                if false {
                                    switch "a" {
                                    case "":
                                        for fi14 := 0; fi14 < 8; fi14++ {
                                            for fi15 := 0; fi15 < 5; fi15++ {
                                                switch "hello" {
                                                case "a":
                                                    v16 := GenS2{F0: (237 <= 0)}
                                                    _ = v16
                                                case "":
                                                    v10 = 1
                                                }
                                                v1 = v1
                                                _ = fi15
                                            }
                                            _ = fi14
                                        }
                                    }
                                    if (true || (v1 && (2 >= v12))) {
                                        var v17 int64 = 572
                                        _ = v17
                                    } else {
                                        v13 -= 9223372036854775807
                                        var v18 string = (("hello" + (("a" + "") + "\n\t")) + "")
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = v18
                                    }
                                } else {
                                    for fi19 := 0; fi19 < 2; fi19++ {
                                        var v20 bool = (18446744073709551615 != (^(v12)))
                                        _ = fi19
                                        _ = v20
                                    }
                                    v1 = true
                                    if false {
                                        switch v13 {
                                        case -491:
                                            v21 := GenS3{F0: 1, F1: v1}
                                            _ = v21
                                        case 2147483647:
                                            v1 = v1
                                        case 32768:
                                            v22 := v5[(-(v13))]
                                            _ = v22
                                        default:
                                            v1 = (806 < 42)
                                        }
                                        v23 := GenS4{F0: 1, F1: 633}
                                        _ = v23
                                    }
                                }
                                v5[32768] = (v6 ^ v6)
                                var v24 int8 = (^(v4))
                                _ = v24
                            } else {
                                v25 := GenS1{F0: int64(v12), F1: 32767, F2: 98}
                                v26 := GenS2{F0: true}
                                switch "foo bar" {
                                case "":
                                    if (v10 != v6) {
                                        v27 := GenS4{F0: 315, F1: (v6 + (^(v6)))}
                                        _ = v27
                                    } else {
                                        v28 := GenS5{F0: v6, F1: 42, F2: 12}
                                        _ = v28
                                    }
                                case "hello":
                                    var v29 string = ("foo bar" + (("" + "\n\t") + "foo bar"))
                                    _ = v29
                                case "\n\t":
                                    fmt.Println(v13)
                                default:
                                    var v30 string = "foo bar"
                                    _ = v30
                                }
                                _ = v25
                                _ = v26
                            }
                        case 127:
                            if (!(false || (!false))) {
                                if v1 {
                                    v31 := v5[(v13 >> ((70) & 63))]
                                    v12 -= (v12 * v12)
                                    _ = v31
                                } else {
                                    v32 := GenS6{F0: 0, F1: 42, F2: int64(v4)}
                                    var v33 string = "hello"
                                    _ = v32
                                    _ = v33
                                }
                                fmt.Println(v12, v1, v12)
                                var v34 int16 = 32767
                                _ = v34
                            } else {
                                var v35 int8 = genFn0(-2147483648, 959, (v12 % ((2) | 1)))
                                var v36 uint64 = 438
                                _ = v35
                                _ = v36
                            }
                        }
                    case "\n\t":
                        var v37 uint8 = 255
                        _ = v37
                    default:
                        v38 := GenS7{F0: 0}
                        _ = v38
                    }
                    _ = v13
                } else {
                    _ = len(v3)
                    switch 4294967295 {
                    case -798:
                        if (-9223372036854775808 == 42) {
                            var v39 int = 128
                            v6 += (^(v6))
                            var v40 float32 = -2.25
                            _ = v39
                            _ = v40
                        } else {
                            var v41 uint8 = 1
                            _ = v41
                        }
                    case -1:
                        for fi42 := 0; fi42 < 15; fi42++ {
                            var v43 int8 = (v4 & genFn0(fi42, 42, v12))
                            switch (fi42 >> ((65535) & 63)) {
                            case -612:
                                v44 := GenS8{F0: 494}
                                _ = v44
                            default:
                                v45 := v5[fi42]
                                _ = v45
                            }
                            _ = fi42
                            _ = v43
                        }
                    case 4294967295:
                        if ((!(218 > 0)) && v1) {
                            v46 := v5[32768]
                            _ = v46
                        } else {
                            var v47 uint8 = 120
                            fmt.Println(v10, v1)
                            _ = v47
                        }
                    default:
                        for fi48 := 0; fi48 < 15; fi48++ {
                            v1 = false
                            for fi49 := 0; fi49 < 12; fi49++ {
                                var v50 int64 = 42
                                _ = fi49
                                _ = v50
                            }
                            _ = fi48
                        }
                    }
                    v12 = (v12 & 18446744073709551615)
                }
                _ = v12
            }
        } else {
            v4 = (-(v4))
            var v51 bool = v1
            fmt.Println(v1)
            _ = v51
        }
        var v52 int32 = -2147483648
        v53 := v5[-972]
        _ = v52
        _ = v53
    }
    v1 = v1
    if true {
        if false {
            for fi54 := 0; fi54 < 7; fi54++ {
                var v55 float64 = 0.1
                var v56 string = "\n\t"
                for fi57 := 0; fi57 < 20; fi57++ {
                    v58 := GenS8{F0: 0}
                    v59 := GenS9{F0: 755, F1: 686}
                    v60 := GenS10{F0: 42}
                    v61 := GenS11{F0: 0}
                    _ = fi57
                    _ = v58
                    _ = v59
                    _ = v60
                    _ = v61
                }
                _ = fi54
                _ = v55
                _ = v56
            }
            var v62 int16 = 1
            v4 = v4
            _ = v62
        }
        var v63 float64 = 3.14159
        v64 := GenS9{F0: 32767, F1: 42}
        _ = v63
        _ = v64
    } else {
        for fi65 := 0; fi65 < 1; fi65++ {
            var v66 int = (-(fi65))
            _ = fi65
            _ = v66
        }
    }
    var v67 int32 = -1
    var v68 uint8 = 1
    for fi69 := 0; fi69 < 18; fi69++ {
        for fi70 := 0; fi70 < 20; fi70++ {
            v71 := v5[0]
            var v72 int16 = 516
            if v1 {
                delete(v5, fi70)
                var v73 float64 = (0.1 / float64((1e10 * 3.14159)))
                _ = v73
            }
            _ = fi70
            _ = v71
            _ = v72
        }
        v6 = 87
        _ = fi69
    }
    v74 := GenS12{F0: "foo bar"}
    if v1 {
        var v75 float32 = 3.14159
        _ = v75
    }
    fmt.Println(v4, v68, v67)
    for fi76 := 0; fi76 < 3; fi76++ {
        for fi77 := 0; fi77 < 3; fi77++ {
            var v78 int = fi77
            v6 = (^(v6))
            v79 := GenS8{F0: -32768}
            _ = fi77
            _ = v78
            _ = v79
        }
        v6 -= 64
        if true {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if true {
                if v1 {
                    v6 = (v10 | v6)
                    if (!(v68 < v68)) {
                        for fi80 := 0; fi80 < 15; fi80++ {
                            delete(v5, fi76)
                            _ = fi80
                        }
                        v81 := GenS8{F0: 1000}
                        _ = v81
                    } else {
                        for fi82 := 0; fi82 < 8; fi82++ {
                            if len(v3) > 0 {
                                v3[0] = (v10 >> ((974) & 63))
                            }
                            _ = fi82
                        }
                    }
                    v67 = -699
                } else {
                    v83 := GenS13{F0: genFn0((fi76 + (fi76 & -973)), 851, 0), F1: 470}
                    switch (fi76 ^ -9223372036854775808) {
                    case -9223372036854775808:
                        if v1 {
                            v84 := GenS5{F0: 647, F1: 107, F2: v68}
                            _ = v84
                        }
                    case 537:
                        var v85 uint16 = 1
                        _ = v85
                    default:
                        fmt.Println(v4, fi76)
                    }
                    v10 = (v6 | (v6 & 2147483647))
                    _ = v83
                }
                var v86 float32 = 1e10
                _ = v86
            }
        } else {
            for fi87 := 0; fi87 < 18; fi87++ {
                var v88 uint8 = 204
                var v89 float32 = ((-2.25 + -2.25) - (float32((-2.25 / 1.5)) + 0.1))
                _ = fi87
                _ = v88
                _ = v89
            }
        }
        _ = fi76
    }
    var v90 uint16 = 1
    v91 := GenS2{F0: (32767 == 1000)}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v67
    _ = v68
    _ = v74
    _ = v90
    _ = v91
}
