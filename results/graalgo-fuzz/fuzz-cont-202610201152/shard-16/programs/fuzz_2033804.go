package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 uint64; F2 uint64 }
type GenS4 struct { F0 string; F1 bool; F2 int32 }
type GenS5 struct { F0 uint8; F1 int; F2 uint32 }
func genFn0(p0 string, p1 uint) string {
    fmt.Println(p0)
    var v6 bool = (42 >= 32767)
    v7 := GenS1{F0: -1, F1: 1}
    fmt.Println(v6, v6)
    if ((42 <= -2147483648) && v6) {
        v8 := GenS2{F0: 1}
        if (-666 > 1) {
            p0 = "foo bar"
            v9 := GenS2{F0: 0}
            for fi10 := 0; fi10 < 18; fi10++ {
                var v11 uint64 = 4294967295
                v12 := GenS1{F0: int64(p1), F1: int16(p1)}
                v13 := []int{fi10, 4294967295, fi10}
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
            _ = v9
        }
        _ = v8
    } else {
        switch 4294967295 {
        case 0:
            fmt.Println(p0)
        case 128:
            if (v6 && v6) {
                var v14 uint16 = 255
                for fi15 := 0; fi15 < 11; fi15++ {
                    if false {
                        var v16 int64 = 864
                        var v17 int16 = 0
                        v18 := GenS3{F0: 0, F1: 642, F2: 18446744073709551615}
                        _ = v16
                        _ = v17
                        _ = v18
                    } else {
                        v19 := []int{int(p1), fi15}
                        if (0 < 100) {
                            p0 = p0
                            if true {
                                for fi20 := 0; fi20 < 15; fi20++ {
                                    if v6 {
                                        v21 := GenS3{F0: 2147483647, F1: 411, F2: 1}
                                        _ = v21
                                    }
                                    _ = len(v19)
                                    var v22 float32 = 3.14159
                                    _ = fi20
                                    _ = v22
                                }
                                var v23 bool = (!(float64((((1.5 * float32((-2.25 / -2.25))) * -2.25) + 1.5)) == 1e10))
                                _ = v23
                            }
                            var v24 int8 = -128
                            _ = v24
                        }
                        for fi25 := 0; fi25 < 7; fi25++ {
                            if v6 {
                                var v26 float64 = (0.0 / 3.14159)
                                var v27 string = "foo bar"
                                var v28 uint16 = v14
                                var v29 float64 = float64(1e10)
                                _ = v26
                                _ = v27
                                _ = v28
                                _ = v29
                            }
                            _ = fi25
                        }
                        _ = v19
                    }
                    _ = fi15
                }
                _ = v14
            }
        case 42:
            for fi30 := 0; fi30 < 2; fi30++ {
                for fi31 := 0; fi31 < 14; fi31++ {
                    var v32 int64 = -14
                    fmt.Println(p1)
                    if (v6 && v6) {
                        v6 = (((float32((1.5 / (1e10 + 0.0))) / 1.0) / float32((-2.25 - ((0.1 - -2.25) / 1e10)))) < (-2.25 / 1.5))
                        var v33 int64 = (v32 >> ((p1) & 63))
                        _ = v33
                    }
                    _ = fi31
                    _ = v32
                }
                _ = fi30
            }
        default:
            switch 42 {
            case -1:
                v34 := GenS2{F0: -1}
                _ = v34
            case 9223372036854775807:
                for fi35 := 0; fi35 < 7; fi35++ {
                    if (42 <= 255) {
                        var v36 int = 9223372036854775807
                        _ = v36
                    }
                    for fi37 := 0; fi37 < 3; fi37++ {
                        v38 := GenS3{F0: -2147483648, F1: 4294967295, F2: 2}
                        _ = fi37
                        _ = v38
                    }
                    _ = fi35
                }
            case 2147483647:
                v6 = ((-32768 > -1) && v6)
            default:
                p0 = "hello"
            }
        }
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return "a"
}
func genFn1(p0 string) uint32 {
    switch -449 {
    case 746:
        v39 := GenS4{F0: "hello", F1: false, F2: -1}
        _ = v39
    case 977:
        v40 := GenS1{F0: -9223372036854775808, F1: 0}
        _ = v40
    case -983:
        fmt.Println(p0)
    default:
        v41 := []int{-595, -956, 32768, 336}
        _ = v41
    }
    switch -2147483648 {
    case 1:
        for fi42 := 0; fi42 < 15; fi42++ {
            v43 := GenS2{F0: 305}
            v44 := []int{fi42, (fi42 / ((-9223372036854775808) | 1)), fi42, (fi42 ^ (fi42 - (fi42 & 128)))}
            p0 = p0
            _ = fi42
            _ = v43
            _ = v44
        }
    case 18:
        if (false || true) {
            var v45 uint32 = 2147483648
            var v46 uint16 = 383
            _ = v45
            _ = v46
        }
    case -807:
        p0 = p0
    }
    var v47 uint8 = 127
    v48 := []int{128, -393}
    _ = p0
    _ = v47
    _ = v48
    return 2147483648
}

func main() {
    var v1 uint16 = 255
    v2 := []int{-439, 4294967295, 499, 32768}
    v3 := make(map[uint32]int)
    var v4 uint64 = 722
    v5 := make(map[string]bool)
    var v49 float32 = 3.14159
    v50 := v3[175]
    switch "foo bar" {
    case "\n\t":
        switch "" {
        case "":
            _ = len(v5)
        case "\n\t":
            if (!(2 != 4294967295)) {
                var v51 float32 = ((1.5 / float32((((1e10 - float64((((1.5 + v49) / (v49 + 1.5)) * (3.14159 + v49)))) / 1e10) / -2.25))) / ((v49 * 0.0) - (1.5 + v49)))
                _ = v51
            } else {
                if len(v2) > 0 {
                    v2[0] = (v50 + (v50 * v50))
                }
                v4 = 1
            }
        case "hello":
            delete(v3, genFn1("a"))
        }
    case "hello":
        if (!true) {
            if true {
                if len(v2) > 0 {
                    v2[0] = (v50 ^ (v50 + v50))
                }
            }
        } else {
            if true {
                for fi52 := 0; fi52 < 4; fi52++ {
                    for fi53 := 0; fi53 < 18; fi53++ {
                        v54 := v5["a"]
                        _ = fi53
                        _ = v54
                    }
                    switch fi52 {
                    case -2147483648:
                        delete(v3, 42)
                    case 4294967295:
                        v4 = 479
                    default:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    var v55 int64 = 0
                    _ = fi52
                    _ = v55
                }
                var v56 int8 = 100
                _ = v56
            } else {
                v57 := GenS1{F0: -9223372036854775808, F1: 42}
                var v58 uint8 = 42
                if true {
                    var v59 float64 = 1e10
                    delete(v5, "hello")
                    _ = v59
                }
                _ = v57
                _ = v58
            }
            if (uint8(v4) <= 255) {
                v49 = v49
                v50 += v50
            }
        }
    case "foo bar":
        if (true || (1e10 <= 1.5)) {
            fmt.Println(v1, v50)
            v60 := GenS2{F0: 32767}
            _ = v60
        }
    default:
        _ = len(v2)
    }
    v1 = v1
    _ = len(v5)
    v61 := GenS1{F0: 9223372036854775807, F1: 1}
    v5["\n\t"] = (420 < 18446744073709551615)
    v62 := GenS4{F0: (((("" + "") + "") + "hello") + ("a" + "foo bar")), F1: false, F2: int32(v4)}
    switch ("foo bar" + "hello") {
    case "a":
        for fi63 := 0; fi63 < 12; fi63++ {
            var v64 string = "foo bar"
            var v65 uint16 = v1
            _ = fi63
            _ = v64
            _ = v65
        }
    case "":
        if (!true) {
            if (-80 <= 100) {
                v1 += v1
            } else {
                v66 := GenS3{F0: 765, F1: uint64(v1), F2: 18446744073709551615}
                v3[545] = 42
                _ = v66
            }
        }
    default:
        var v67 int64 = 257
        _ = v67
    }
    if true {
        for fi68 := 0; fi68 < 5; fi68++ {
            for fi69 := 0; fi69 < 6; fi69++ {
                if false {
                    var v70 int16 = -107
                    var v71 uint = 1
                    v72 := GenS3{F0: 0, F1: v4, F2: (v4 << ((2) & 63))}
                    _ = v70
                    _ = v71
                    _ = v72
                } else {
                    v4 += (v4 | v4)
                    var v73 uint = 255
                    _ = v73
                }
                switch fi68 {
                case 4294967295:
                    v74 := GenS3{F0: -1, F1: 1, F2: v4}
                    _ = v74
                }
                _ = fi69
            }
            v50 = v50
            if false {
                var v75 uint = 622
                for fi76 := 0; fi76 < 18; fi76++ {
                    var v77 uint8 = 127
                    _ = fi76
                    _ = v77
                }
                v50 -= (-(v50))
                _ = v75
            } else {
                fmt.Println(v49, v1)
                if (!(1e10 >= (0.0 / float64(3.14159)))) {
                    v78 := GenS3{F0: 2147483647, F1: 666, F2: (v4 & 1)}
                    fmt.Println(v4, v4)
                    if false {
                        v49 += (1.5 - v49)
                        v79 := GenS4{F0: "hello", F1: (1 <= (v1 >> ((778) & 63))), F2: 42}
                        _ = v79
                    }
                    _ = v78
                }
            }
            _ = fi68
        }
        v1 = 1
    } else {
        var v80 float64 = -2.25
        _ = v80
    }
    v1 -= 2
    v81 := v5["\n\t"]
    if true {
        if len(v2) > 0 {
            v2[0] = v50
        }
        v82 := GenS5{F0: 174, F1: (-(v50)), F2: 517}
        _ = v82
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v61
    _ = v62
    _ = v81
}
