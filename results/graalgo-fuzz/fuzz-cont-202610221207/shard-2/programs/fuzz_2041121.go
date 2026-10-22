package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 int; F2 uint32 }
type GenS4 struct { F0 int8; F1 uint8; F2 int16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int32; F1 int32; F2 bool }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int64; F1 int; F2 int }
type GenS9 struct { F0 uint; F1 uint64; F2 uint16 }
type GenS10 struct { F0 uint64; F1 uint8 }
type GenS11 struct { F0 string; F1 int; F2 int16 }
func genFn0(p0 uint32, p1 string, p2 int16) uint32 {
    switch ((p1 + (p1 + "a")) + p1) {
    case "foo bar":
        if true {
            var v7 float32 = float32(float64(1.5))
            for fi8 := 0; fi8 < 20; fi8++ {
                if (0 != 1) {
                    if false {
                        var v9 float64 = 1e10
                        fmt.Println(p1, v7)
                        var v10 int32 = -2147483648
                        _ = v10
                        _ = v9
                    } else {
                        for fi11 := 0; fi11 < 13; fi11++ {
                            v12 := GenS1{F0: ((p1 + (("" + "") + (p1 + p1))) + "hello")}
                            p0 = p0
                            var v13 int8 = 100
                            _ = fi11
                            _ = v12
                            _ = v13
                        }
                    }
                    v14 := []int{(-(fi8)), fi8, (fi8 | (fi8 % ((fi8) | 1))), fi8}
                    _ = v14
                }
                switch 32767 {
                case 32768:
                    var v15 uint16 = 335
                    _ = v15
                case -1:
                    switch int(p2) {
                    case -338:
                        v16 := []int{0, (fi8 - (^(fi8))), (fi8 | (fi8 & fi8))}
                        _ = v16
                    case -923:
                        v7 = -2.25
                    case -5:
                        v17 := []int{}
                        _ = v17
                    }
                }
                v18 := []int{742}
                _ = fi8
                _ = v18
            }
            v19 := []int{}
            _ = v19
            _ = v7
        } else {
            for fi20 := 0; fi20 < 16; fi20++ {
                p1 = "\n\t"
                v21 := GenS2{F0: 27}
                v22 := []int{fi20, fi20, fi20}
                _ = fi20
                _ = v21
                _ = v22
            }
            p0 = 4294967295
            p1 = "foo bar"
        }
    case "\n\t":
        switch 832 {
        case 32767:
            var v23 uint16 = 348
            _ = v23
        case 0:
            v24 := GenS3{F0: 109, F1: 42, F2: 1}
            _ = v24
        case 9223372036854775807:
            v25 := GenS1{F0: p1}
            _ = v25
        }
    default:
        v26 := []int{}
        _ = v26
    }
    if false {
        p0 = (^(p0))
    } else {
        var v27 uint = 1
        _ = v27
    }
    for fi28 := 0; fi28 < 9; fi28++ {
        p2 = 739
        if true {
            p1 = p1
            p1 = p1
        } else {
            if false {
                switch "" {
                case "foo bar":
                    fmt.Println(p0)
                default:
                    var v29 uint64 = 0
                    _ = v29
                }
                v30 := []int{fi28, (fi28 % (((fi28 & fi28)) | 1)), (fi28 % ((fi28) | 1)), (fi28 >> ((0) & 63))}
                _ = v30
            } else {
                v31 := []int{(fi28 >> ((65535) & 63)), (fi28 + (fi28 & 32768)), (fi28 - fi28), fi28}
                _ = v31
            }
        }
        _ = fi28
    }
    _ = p0
    _ = p1
    _ = p2
    return (p0 >> ((4294967295) & 63))
}

func main() {
    v1 := []int{-679, -477, -2147483648, 1}
    var v2 uint8 = 128
    v3 := make(map[string]string)
    var v4 uint8 = v2
    v5 := make(map[uint32]bool)
    var v6 uint64 = 42
    v4 -= v4
    var v32 uint = uint(v6)
    v6 += v6
    switch "" {
    case "\n\t":
        v33 := GenS4{F0: -1, F1: (v4 >> ((v32) & 63)), F2: 1}
        _ = v33
    default:
        v34 := GenS3{F0: -1, F1: 1, F2: 459}
        _ = v34
    }
    for fi35 := 0; fi35 < 7; fi35++ {
        var v36 float64 = -2.25
        if false {
            v5[4294967295] = (v6 > (v6 % ((543) | 1)))
        } else {
            v4 = (v4 + v4)
            v4 += (v2 / ((255) | 1))
        }
        _ = fi35
        _ = v36
    }
    for fi37 := 0; fi37 < 5; fi37++ {
        v2 += uint8(v6)
        switch fi37 {
        case 42:
            for fi38 := 0; fi38 < 4; fi38++ {
                var v39 uint8 = (^(v2))
                _ = fi38
                _ = v39
            }
        }
        if (1e10 <= (((float64(0.0) * (3.14159 - ((3.14159 - 1e10) * float64(-2.25)))) / 3.14159) - 0.1)) {
            v4 += (^(v4))
        } else {
            v4 = 128
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if true {
                var v40 int8 = -128
                var v41 string = "hello"
                _ = v40
                _ = v41
            } else {
                switch "hello" {
                case "a":
                    v42 := GenS5{F0: genFn0(65535, "hello", -1)}
                    _ = v42
                case "hello":
                    v6 -= (v6 ^ (^(v6)))
                default:
                    switch "\n\t" {
                    case "hello":
                        var v43 uint16 = 658
                        _ = v43
                    case "foo bar":
                        switch 32767 {
                        case 32767:
                            if true {
                                delete(v5, 461)
                                var v44 uint = (v32 * v32)
                                _ = v44
                            }
                        }
                    default:
                        if false {
                            var v45 float32 = 3.14159
                            _ = v45
                        }
                    }
                }
            }
        }
        _ = fi37
    }
    if ((-216 < 722) && true) {
        var v46 uint32 = 42
        var v47 int16 = 821
        var v48 float32 = (1.5 + -2.25)
        _ = v46
        _ = v47
        _ = v48
    } else {
        var v49 uint16 = 0
        _ = v49
    }
    if (-32768 < 0) {
        if true {
            var v50 string = "hello"
            if (682 != -844) {
                if ((false && ((v50 != v50) || (!false))) || ((0.0 != 1e10) || (42 > 0))) {
                    v51 := GenS2{F0: (v4 & v2)}
                    v4 -= 60
                    var v52 float32 = 0.0
                    _ = v51
                    _ = v52
                }
            } else {
                for fi53 := 0; fi53 < 10; fi53++ {
                    v4 = 255
                    if len(v1) > 0 {
                        v1[0] = -59
                    }
                    if true {
                        fmt.Println(v4, v2)
                    } else {
                        var v54 int16 = 1000
                        _ = v54
                    }
                    _ = fi53
                }
                v32 -= (v32 ^ (v32 - (^(v32))))
                v2 = (v4 >> ((v32) & 63))
            }
            v5[2147483647] = (255 == v32)
            _ = v50
        } else {
            if true {
                if true {
                    for fi55 := 0; fi55 < 13; fi55++ {
                        for fi56 := 0; fi56 < 16; fi56++ {
                            if (!false) {
                                v4 -= (v4 & (v4 % ((v2) | 1)))
                                v5[0] = (-759 != int32(v32))
                            } else {
                                v57 := GenS1{F0: ("foo bar" + (("hello" + "") + ""))}
                                v32 = 0
                                _ = v57
                            }
                            _ = fi56
                        }
                        delete(v5, genFn0(51, "\n\t", 1000))
                        _ = fi55
                    }
                    var v58 uint64 = v6
                    _ = v58
                } else {
                    if ((^(v32)) != 38) {
                        v32 += 18446744073709551615
                        v2 -= v4
                    } else {
                        if (true && true) {
                            v59 := GenS6{F0: 2147483647, F1: 143, F2: false}
                            v4 = (^(v4))
                            v32 -= v32
                            _ = v59
                        }
                        var v60 int8 = 42
                        var v61 uint8 = (v2 << (((^(v32))) & 63))
                        _ = v60
                        _ = v61
                    }
                    var v62 uint8 = v2
                    var v63 int16 = -32768
                    _ = v62
                    _ = v63
                }
            } else {
                for fi64 := 0; fi64 < 3; fi64++ {
                    var v65 string = ("a" + "")
                    v66 := GenS7{F0: 431}
                    if len(v1) > 0 {
                        v1[0] = (fi64 << (((v32 << ((v32) & 63))) & 63))
                    }
                    _ = fi64
                    _ = v65
                    _ = v66
                }
                switch "hello" {
                case "hello":
                    switch 702 {
                    case -9223372036854775808:
                        if true {
                            switch "a" {
                            case "hello":
                                var v67 uint64 = (^(v6))
                                _ = v67
                            default:
                                v2 = (v4 >> (((v32 << ((899) & 63))) & 63))
                            }
                        }
                    case 0:
                        if (true || false) {
                            for fi68 := 0; fi68 < 19; fi68++ {
                                if false {
                                    v69 := GenS7{F0: 32767}
                                    _ = v69
                                }
                                var v70 float32 = 3.14159
                                for fi71 := 0; fi71 < 3; fi71++ {
                                    var v72 uint = (^(v32))
                                    if (true && (1e10 >= 0.1)) {
                                        var v73 bool = false
                                        var v74 uint8 = v4
                                        _ = v73
                                        _ = v74
                                    } else {
                                        if (120 <= uint32(v72)) {
                                            v2 = v2
                                            if (0.1 < 3.14159) {
                                                var v75 int32 = -1
                                                _ = v75
                                            }
                                        } else {
                                            switch fi71 {
                                            case 127:
                                                v72 = 444
                                            case 856:
                                                v1 = append(v1, fi68)
                                            default:
                                                fmt.Println(v70, v72, v2)
                                            }
                                            if true {
                                                switch fi68 {
                                                case -356:
                                                    v76 := GenS4{F0: 127, F1: (v4 | v2), F2: 1}
                                                    _ = v76
                                                case 418:
                                                    v4 += v4
                                                    v77 := GenS8{F0: -371, F1: fi71, F2: (fi68 << ((225) & 63))}
                                                    _ = v77
                                                default:
                                                    _ = len(v1)
                                                }
                                            } else {
                                                switch fi68 {
                                                case 32767:
                                                    if len(v1) > 0 {
                                                        v1[0] = (fi71 & fi68)
                                                    }
                                                case 32768:
                                                    v78 := GenS7{F0: 0}
                                                    _ = v78
                                                case 976:
                                                    v79 := v5[323]
                                                    _ = v79
                                                default:
                                                    fmt.Println(v6, v72, v72)
                                                }
                                                v6 = v6
                                            }
                                        }
                                    }
                                    v80 := GenS3{F0: 1, F1: fi68, F2: 190}
                                    v2 = 113
                                    switch (fi68 ^ fi68) {
                                    case 32767:
                                        v70 -= v70
                                    case 1:
                                        v70 = float32((-2.25 / 3.14159))
                                    default:
                                        fmt.Println(v32, v70, v70)
                                    }
                                    _ = fi71
                                    _ = v72
                                    _ = v80
                                }
                                _ = fi68
                                _ = v70
                            }
                        } else {
                            if false {
                                var v81 bool = (18 == (^(v32)))
                                if len(v1) > 0 {
                                    v1[0] = 128
                                }
                                _ = v81
                            } else {
                                var v82 uint16 = 834
                                var v83 int64 = 693
                                _ = v82
                                _ = v83
                            }
                        }
                    case 996:
                        _ = len(v5)
                    default:
                        v84 := GenS9{F0: (v32 << ((v32) & 63)), F1: (^(v6)), F2: 65535}
                        _ = v84
                    }
                case "":
                    switch "" {
                    case "foo bar":
                        v85 := GenS10{F0: 12, F1: (v2 + v4)}
                        _ = v85
                    case "":
                        var v86 uint16 = 485
                        _ = v86
                    case "\n\t":
                        for fi87 := 0; fi87 < 18; fi87++ {
                            v6 = (^(v6))
                            v88 := GenS9{F0: v32, F1: (v6 * 42), F2: 42}
                            _ = fi87
                            _ = v88
                        }
                    default:
                        v5[0] = true
                    }
                default:
                    switch 42 {
                    case -362:
                        var v89 int = 9223372036854775807
                        _ = v89
                    default:
                        switch 1 {
                        case 315:
                            if (42 < -1) {
                                v32 = v32
                            }
                        }
                    }
                }
            }
            v90 := GenS11{F0: "\n\t", F1: 316, F2: 85}
            v32 = v32
            _ = v90
        }
    }
    fmt.Println(v32, v4)
    v2 -= 42
    var v91 uint16 = 2
    switch 127 {
    case -4:
        v5[334] = (!true)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v91
}
