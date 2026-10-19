package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 int16; F2 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string; F1 uint16 }
type GenS5 struct { F0 bool; F1 int8 }
type GenS6 struct { F0 int; F1 uint64 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 string; F1 uint8; F2 int16 }
func genFn0(p0 bool, p1 uint) uint32 {
    fmt.Println(p1)
    v4 := GenS1{F0: 2}
    _ = p0
    _ = p1
    _ = v4
    return 4294967295
}
func genFn1(p0 uint8, p1 uint32, p2 bool) uint64 {
    if true {
        p0 = (p0 << ((2) & 63))
    } else {
        var v5 int64 = 881
        for fi6 := 0; fi6 < 11; fi6++ {
            for fi7 := 0; fi7 < 20; fi7++ {
                v5 += v5
                v5 = (-(v5))
                v8 := GenS1{F0: (p1 / ((p1) | 1))}
                var v9 uint32 = (^(p1))
                p2 = false
                if (((0.0 < 0.1) || p2) || ((p2 || p2) || p2)) {
                    var v10 uint8 = (^(p0))
                    var v11 float32 = (0.0 - 0.1)
                    _ = v10
                    _ = v11
                }
                _ = fi7
                _ = v8
                _ = v9
            }
            _ = fi6
        }
        _ = v5
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        var v13 uint = 65535
        _ = fi12
        _ = v13
    }
    var v14 uint16 = 730
    v15 := []int{9223372036854775807, int(p0), 4294967295, -61}
    var v16 uint32 = genFn0(p2, 2)
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    _ = v16
    return 486
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []uint32{2147483647}
    var v3 uint32 = 1
    switch "hello" {
    case "":
        if (true && true) {
            for fi17 := 0; fi17 < 16; fi17++ {
                v3 = (v3 * genFn0(false, 18446744073709551615))
                _ = len(v2)
                switch "hello" {
                case "\n\t":
                    var v18 int64 = -9223372036854775808
                    _ = v18
                default:
                    v19 := GenS2{F0: -9223372036854775808, F1: 0, F2: 255}
                    _ = v19
                }
                _ = fi17
            }
            for fi20 := 0; fi20 < 9; fi20++ {
                v3 = 942
                _ = fi20
            }
        } else {
            v1 = 317
            if ((v3 << ((18446744073709551615) & 63)) == genFn0(false, 647)) {
                var v21 int64 = -1
                switch 32767 {
                case 42:
                    v22 := GenS2{F0: -22, F1: 1, F2: 817}
                    _ = v22
                case -9223372036854775808:
                    var v23 uint16 = 0
                    _ = v23
                }
                _ = v21
            } else {
                for fi24 := 0; fi24 < 10; fi24++ {
                    v3 = (v3 >> ((42) & 63))
                    _ = fi24
                }
                v3 = (v3 >> ((4294967295) & 63))
            }
        }
    case "a":
        fmt.Println(v1)
    case "foo bar":
        if (!false) {
            for fi25 := 0; fi25 < 11; fi25++ {
                if len(v2) > 0 {
                    v2[0] = 0
                }
                _ = fi25
            }
            v1 = genFn1(128, genFn0((true && (!false)), 65535), (969 <= 65535))
            if true {
                if (4294967295 != 42) {
                    switch 1 {
                    case -1:
                        var v26 uint64 = v1
                        _ = v26
                    }
                    if true {
                        var v27 int16 = 1
                        if true {
                            if len(v2) > 0 {
                                v2[0] = genFn0(true, 0)
                            }
                            if false {
                                if false {
                                    for fi28 := 0; fi28 < 13; fi28++ {
                                        v1 -= v1
                                        if (!true) {
                                            var v29 bool = false
                                            var v30 uint8 = 166
                                            _ = v29
                                            _ = v30
                                        }
                                        for fi31 := 0; fi31 < 10; fi31++ {
                                            var v32 bool = (!(!(32767 != 676)))
                                            v1 = v1
                                            _ = fi31
                                            _ = v32
                                        }
                                        _ = fi28
                                    }
                                    switch int(v27) {
                                    case 128:
                                        v33 := GenS1{F0: v3}
                                        _ = v33
                                    case -693:
                                        if true {
                                            if true {
                                                for fi34 := 0; fi34 < 10; fi34++ {
                                                    var v35 float32 = (3.14159 / 1.0)
                                                    if ("\n\t" <= "a") {
                                                        v36 := GenS2{F0: 9223372036854775807, F1: v27, F2: 0}
                                                        _ = len(v2)
                                                        switch "a" {
                                                        case "foo bar":
                                                            for fi37 := 0; fi37 < 13; fi37++ {
                                                                var v38 bool = false
                                                                var v39 bool = v38
                                                                v40 := GenS3{F0: 65535}
                                                                _ = fi37
                                                                _ = v38
                                                                _ = v39
                                                                _ = v40
                                                            }
                                                        case "":
                                                            for fi41 := 0; fi41 < 18; fi41++ {
                                                                var v42 int16 = v27
                                                                var v43 uint16 = 32767
                                                                v44 := GenS3{F0: (^(v3))}
                                                                var v45 int64 = 9223372036854775807
                                                                _ = fi41
                                                                _ = v42
                                                                _ = v43
                                                                _ = v44
                                                                _ = v45
                                                            }
                                                        }
                                                        _ = v36
                                                    }
                                                    _ = fi34
                                                    _ = v35
                                                }
                                            } else {
                                                var v46 float64 = 1.5
                                                if false {
                                                    v27 = 32767
                                                    if true {
                                                        v3 = (v3 | genFn0(((!true) || false), 18446744073709551615))
                                                        if len(v2) > 0 {
                                                            v2[0] = (v3 + (v3 % ((v3) | 1)))
                                                        }
                                                        v27 = 1
                                                    }
                                                    if false {
                                                        v3 = (v3 >> ((18446744073709551615) & 63))
                                                        if false {
                                                            fmt.Println(v1, v3, v27)
                                                            var v47 int16 = v27
                                                            fmt.Println(v27)
                                                            _ = v47
                                                        } else {
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        }
                                                    }
                                                }
                                                _ = len(v2)
                                                _ = v46
                                            }
                                            if false {
                                                if len(v2) > 0 {
                                                    v2[0] = genFn0((!true), 833)
                                                }
                                                var v48 uint64 = (v1 / ((v1) | 1))
                                                for fi49 := 0; fi49 < 12; fi49++ {
                                                    v48 += (v48 - v48)
                                                    for fi50 := 0; fi50 < 18; fi50++ {
                                                        if false {
                                                            v1 += v1
                                                            v1 -= (v48 >> ((255) & 63))
                                                            var v51 uint = 4294967295
                                                            switch 128 {
                                                            case 266:
                                                                v52 := GenS1{F0: v3}
                                                                _ = v52
                                                            case 127:
                                                                v53 := GenS3{F0: (v3 + (v3 | (v3 | (^(v3)))))}
                                                                var v54 int32 = 803
                                                                switch "a" {
                                                                case "foo bar":
                                                                    v3 += (v3 / ((4294967295) | 1))
                                                                case "a":
                                                                    var v55 uint32 = 562
                                                                    _ = v55
                                                                }
                                                                switch fi49 {
                                                                case 941:
                                                                    if (int16(v54) != v27) {
                                                                        fmt.Println(v1, v51)
                                                                    } else {
                                                                        var v56 int8 = -1
                                                                        _ = v56
                                                                    }
                                                                case 9223372036854775807:
                                                                    var v57 bool = true
                                                                    _ = v57
                                                                default:
                                                                    _ = len(v2)
                                                                    var v58 float32 = (1e10 * (3.14159 - (3.14159 - 3.14159)))
                                                                    _ = len(v2)
                                                                    _ = v58
                                                                }
                                                                var v59 uint32 = v3
                                                                _ = v53
                                                                _ = v54
                                                                _ = v59
                                                            }
                                                            var v60 uint64 = 18446744073709551615
                                                            _ = v51
                                                            _ = v60
                                                        } else {
                                                            if (!((55 <= int32(fi49)) && false)) {
                                                                v61 := GenS4{F0: "a", F1: 0}
                                                                if (18446744073709551615 > 255) {
                                                                    if (!true) {
                                                                        v48 = (v1 + (^(v1)))
                                                                    }
                                                                    var v62 float32 = 0.0
                                                                    _ = v62
                                                                } else {
                                                                    v63 := GenS1{F0: (v3 >> ((356) & 63))}
                                                                    _ = v63
                                                                }
                                                                _ = v61
                                                            }
                                                            v3 += (v3 * 2147483648)
                                                        }
                                                        if (false && (v3 >= (v3 ^ (^(v3))))) {
                                                            switch "hello" {
                                                            case "foo bar":
                                                                var v64 uint32 = (v3 - (v3 * (v3 ^ (v3 * (v3 * v3)))))
                                                                _ = len(v2)
                                                                _ = v64
                                                            case "":
                                                                if (true || (!false)) {
                                                                    v65 := GenS5{F0: true, F1: -34}
                                                                    switch fi50 {
                                                                    case -95:
                                                                        switch ("hello" + (((("hello" + "\n\t") + (("foo bar" + "") + ("\n\t" + "foo bar"))) + "foo bar") + "hello")) {
                                                                        case "\n\t":
                                                                            var v66 uint16 = 2
                                                                            _ = v66
                                                                        }
                                                                    }
                                                                    _ = v65
                                                                }
                                                            }
                                                            switch -320 {
                                                            case 4294967295:
                                                                v3 = (^(v3))
                                                            case 164:
                                                                switch fi50 {
                                                                case 127:
                                                                    var v67 string = "\n\t"
                                                                    _ = v67
                                                                case 561:
                                                                    var v68 float32 = 0.0
                                                                    _ = v68
                                                                default:
                                                                    v48 += (v48 / ((2) | 1))
                                                                }
                                                            case -9223372036854775808:
                                                                switch 9223372036854775807 {
                                                                case -1:
                                                                    var v69 int8 = 107
                                                                    _ = v69
                                                                case 406:
                                                                    fmt.Println(v48, v27)
                                                                case 32767:
                                                                    fmt.Println(v1)
                                                                default:
                                                                    fmt.Println(fi50, v27)
                                                                }
                                                            }
                                                        } else {
                                                            fmt.Println(v27, fi49, v27)
                                                        }
                                                        v48 += 2
                                                        _ = fi50
                                                    }
                                                    v3 = genFn0((255 > 255), 255)
                                                    _ = fi49
                                                }
                                                _ = v48
                                            }
                                        }
                                    }
                                    var v70 uint8 = 255
                                    _ = v70
                                }
                            } else {
                                var v71 int32 = 42
                                v27 += v27
                                fmt.Println(v3, v27, v71)
                                _ = v71
                            }
                        }
                        _ = v27
                    }
                }
                v3 -= v3
            } else {
                fmt.Println(v1, v3)
            }
        }
    }
    v72 := GenS4{F0: (("" + ((("a" + "\n\t") + "\n\t") + ((("a" + "\n\t") + "\n\t") + "a"))) + (("foo bar" + "\n\t") + "a")), F1: 539}
    if true {
        _ = len(v2)
        v1 = v1
        fmt.Println(v3, v3)
    }
    for fi73 := 0; fi73 < 18; fi73++ {
        v3 += 2147483647
        _ = fi73
    }
    v3 = 2147483647
    v1 += (v1 >> ((1) & 63))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v1, v1)
    var v74 int8 = 127
    v74 -= -1
    _ = len(v2)
    _ = len(v2)
    v2 = append(v2, v3)
    v74 = (v74 * (^(v74)))
    for fi75 := 0; fi75 < 18; fi75++ {
        switch fi75 {
        case 0:
            if (!(true || true)) {
                for fi76 := 0; fi76 < 1; fi76++ {
                    switch "" {
                    case "foo bar":
                        if true {
                            var v77 int = (fi75 * (fi76 + fi76))
                            v78 := GenS6{F0: (fi75 | (fi75 ^ v77)), F1: v1}
                            var v79 int64 = 29
                            _ = v77
                            _ = v78
                            _ = v79
                        } else {
                            if true {
                                if false {
                                    var v80 uint64 = 18446744073709551615
                                    _ = v80
                                } else {
                                    fmt.Println(fi75)
                                }
                                var v81 string = "\n\t"
                                _ = v81
                            }
                            switch "a" {
                            case "\n\t":
                                v74 -= (v74 << ((579) & 63))
                            case "foo bar":
                                v3 += genFn0(false, 255)
                            case "":
                                var v82 int32 = -2147483648
                                _ = v82
                            default:
                                if (591 < 907) {
                                    v83 := GenS7{F0: 375}
                                    if ((v3 + v3) != (v3 - (v3 * (v3 << ((423) & 63))))) {
                                        var v84 uint32 = (v3 ^ 1)
                                        if (false || (42 == -1)) {
                                            var v85 bool = (!true)
                                            _ = v85
                                        }
                                        v84 -= (^(v84))
                                        v74 -= 0
                                        _ = v84
                                    }
                                    fmt.Println(v1, v3)
                                    if ((v1 / ((v1) | 1)) < uint64(v3)) {
                                        var v86 float32 = 0.1
                                        if false {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                        v1 = (v1 >> ((65535) & 63))
                                        _ = v86
                                    }
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = v83
                                } else {
                                    switch (fi76 & (fi75 >> ((1) & 63))) {
                                    case -826:
                                        if len(v2) > 0 {
                                            v2[0] = (v3 * v3)
                                        }
                                    case -762:
                                        v3 -= genFn0(true, 2)
                                    default:
                                        fmt.Println(fi76)
                                    }
                                }
                            }
                        }
                    default:
                        v3 = v3
                    }
                    _ = len(v2)
                    v3 += (v3 % ((genFn0(true, 2)) | 1))
                    _ = fi76
                }
                v2 = append(v2, v3)
            } else {
                switch 1 {
                case 4294967295:
                    v87 := GenS6{F0: fi75, F1: v1}
                    _ = v87
                case -9223372036854775808:
                    fmt.Println(fi75, v1)
                }
                var v88 float32 = 0.0
                v89 := GenS2{F0: 380, F1: -1, F2: 2}
                _ = v88
                _ = v89
            }
        case -362:
            if len(v2) > 0 {
                v2[0] = (v3 * (v3 * uint32(v1)))
            }
        case 32768:
            v90 := GenS8{F0: "hello", F1: 1, F2: -65}
            _ = v90
        }
        fmt.Println(v1)
        fmt.Println(v1)
        _ = fi75
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v72
    _ = v74
}
