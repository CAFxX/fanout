package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint; F1 bool }
type GenS5 struct { F0 int16; F1 string; F2 string }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int64 }
func genFn0(p0 int, p1 int64, p2 uint) uint16 {
    var v7 float64 = 3.14159
    for fi8 := 0; fi8 < 15; fi8++ {
        v9 := GenS1{F0: 1}
        _ = fi8
        _ = v9
    }
    p2 = p2
    var v10 uint8 = 128
    var v11 float64 = float64(1e10)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v7
    return 61
}
func genFn1() int32 {
    switch "hello" {
    case "":
        for fi12 := 0; fi12 < 8; fi12++ {
            var v13 int32 = 42
            v13 = 1
            var v14 int = fi12
            _ = fi12
            _ = v13
            _ = v14
        }
    default:
        var v15 int64 = -841
        _ = v15
    }
    switch 32768 {
    case 4294967295:
        var v16 int8 = -1
        _ = v16
    case -772:
        for fi17 := 0; fi17 < 3; fi17++ {
            v18 := []int{0}
            _ = fi17
            _ = v18
        }
    case 1:
        if true {
            v19 := GenS1{F0: 1}
            _ = v19
        }
    }
    var v20 uint16 = 498
    _ = v20
    return -1
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = -128
    var v3 int16 = -32768
    var v4 int8 = (v2 + (v2 - -128))
    v5 := []uint32{0}
    var v6 uint64 = 0
    v21 := GenS1{F0: 100}
    switch 288 {
    case -2147483648:
        var v22 uint = 1
        _ = v22
    case 4294967295:
        if false {
            v6 += v6
            v5 = append(v5, 2147483647)
            var v23 uint16 = 32767
            _ = v23
        } else {
            v1[514] = false
            if ((v6 * v6) > v6) {
                var v24 float64 = 0.0
                _ = v24
            } else {
                var v25 uint32 = 42
                if false {
                    delete(v1, 2)
                    if (!(2 > uint(v25))) {
                        switch ("hello" + "foo bar") {
                        case "\n\t":
                            v1[0] = false
                        case "a":
                            v6 -= v6
                        case "hello":
                            var v26 bool = (65535 <= (v25 << ((uint(v25)) & 63)))
                            _ = v26
                        default:
                            for fi27 := 0; fi27 < 1; fi27++ {
                                var v28 uint8 = 64
                                if (65535 != 966) {
                                    v6 = 707
                                    switch ("foo bar" + ("" + "")) {
                                    case "hello":
                                        for fi29 := 0; fi29 < 10; fi29++ {
                                            v30 := GenS1{F0: v2}
                                            _ = fi29
                                            _ = v30
                                        }
                                    case "":
                                        v6 -= (v6 + (v6 & v6))
                                    case "a":
                                        var v31 uint64 = v6
                                        _ = v31
                                    default:
                                        var v32 uint8 = (v28 % ((42) | 1))
                                        _ = v32
                                    }
                                }
                                var v33 uint8 = uint8(v6)
                                _ = fi27
                                _ = v28
                                _ = v33
                            }
                        }
                        v3 = (v3 + 0)
                    } else {
                        var v34 int16 = (v3 | v3)
                        if true {
                            var v35 int32 = 0
                            var v36 uint32 = v25
                            _ = v35
                            _ = v36
                        }
                        var v37 uint16 = genFn0(-768, 42, 18446744073709551615)
                        _ = v34
                        _ = v37
                    }
                }
                _ = v25
            }
            if true {
                var v38 int = -945
                v39 := GenS2{F0: -943}
                if (false && ((v4 >= v4) || (1 >= 255))) {
                    for fi40 := 0; fi40 < 9; fi40++ {
                        for fi41 := 0; fi41 < 12; fi41++ {
                            if (357 < genFn0((-(fi41)), 1, 147)) {
                                if ((0 >= 128) || (128 > 68)) {
                                    fmt.Println(v2)
                                    var v42 uint = 4294967295
                                    switch 127 {
                                    case 4294967295:
                                        v43 := GenS2{F0: 0}
                                        _ = v43
                                    case -1:
                                        var v44 float32 = (0.0 + 0.0)
                                        _ = v44
                                    default:
                                        v45 := GenS3{F0: 1}
                                        _ = v45
                                    }
                                    _ = v42
                                } else {
                                    var v46 uint64 = 759
                                    var v47 int32 = -883
                                    if (454 != 1) {
                                        _ = len(v5)
                                    }
                                    if true {
                                        v1[0] = false
                                        var v48 int16 = (v3 | v3)
                                        if ((v6 < v6) && true) {
                                            v48 += v48
                                        }
                                        fmt.Println(v47, v48, v38)
                                        fmt.Println(v47)
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        var v49 float32 = -2.25
                                        var v50 uint32 = 4294967295
                                        _ = v48
                                        _ = v49
                                        _ = v50
                                    }
                                    v5 = append(v5, 2147483648)
                                    var v51 int16 = (^(v3))
                                    if true {
                                        v4 = v2
                                        fmt.Println(v46, fi40, fi40)
                                    }
                                    _ = v46
                                    _ = v47
                                    _ = v51
                                }
                            } else {
                                fmt.Println(fi40, v38)
                                _ = len(v5)
                                v6 = 2
                            }
                            var v52 float64 = 0.1
                            v53 := GenS4{F0: 255, F1: (("\n\t" + "hello") >= ("foo bar" + "\n\t"))}
                            _ = fi41
                            _ = v52
                            _ = v53
                        }
                        v4 = v4
                        if false {
                            if ((v38 == (fi40 / ((42) | 1))) || true) {
                                var v54 uint = 42
                                _ = v54
                            }
                            if (false && true) {
                                v3 = 0
                                var v55 int16 = (^(v3))
                                if false {
                                    for fi56 := 0; fi56 < 10; fi56++ {
                                        var v57 uint64 = v6
                                        switch (v38 / ((32768) | 1)) {
                                        case 845:
                                            v6 = (v6 ^ (^(v6)))
                                        case 0:
                                            switch "hello" {
                                            case "":
                                                var v58 uint16 = 483
                                                _ = v58
                                            }
                                        case 642:
                                            v1[4294967295] = (true && false)
                                        }
                                        fmt.Println(v4)
                                        _ = fi56
                                        _ = v57
                                    }
                                    var v59 uint16 = 945
                                    v6 += v6
                                    _ = v59
                                } else {
                                    switch "foo bar" {
                                    case "":
                                        if false {
                                            var v60 uint64 = v6
                                            var v61 int8 = (v2 / (((v4 + v2)) | 1))
                                            v5 = append(v5, 65535)
                                            _ = v60
                                            _ = v61
                                        } else {
                                            for fi62 := 0; fi62 < 17; fi62++ {
                                                switch ("hello" + ("\n\t" + "a")) {
                                                case "hello":
                                                    if (((v3 & (v55 | (-(v3)))) != -693) && ((-279 <= v55) || false)) {
                                                        var v63 uint32 = 65535
                                                        v55 = v3
                                                        _ = v63
                                                    } else {
                                                        if (127 == fi62) {
                                                            var v64 string = "hello"
                                                            _ = v64
                                                        } else {
                                                            v65 := GenS3{F0: 2147483647}
                                                            v66 := GenS5{F0: v3, F1: "hello", F2: "\n\t"}
                                                            v67 := GenS6{F0: (^(v4))}
                                                            _ = v65
                                                            _ = v66
                                                            _ = v67
                                                        }
                                                        var v68 uint64 = 929
                                                        _ = v68
                                                    }
                                                case "\n\t":
                                                    v4 += -128
                                                case "":
                                                    _ = len(v5)
                                                default:
                                                    v2 -= 36
                                                }
                                                v69 := GenS3{F0: 2147483648}
                                                v3 = v3
                                                _ = fi62
                                                _ = v69
                                            }
                                            var v70 uint16 = genFn0((-(fi40)), 623, 0)
                                            for fi71 := 0; fi71 < 17; fi71++ {
                                                v2 = (^(v2))
                                                _ = len(v5)
                                                var v72 uint = 629
                                                _ = fi71
                                                _ = v72
                                            }
                                            _ = v70
                                        }
                                    }
                                }
                                _ = v55
                            }
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        } else {
                            for fi73 := 0; fi73 < 10; fi73++ {
                                fmt.Println(fi40)
                                _ = len(v1)
                                v74 := GenS6{F0: -128}
                                _ = fi73
                                _ = v74
                            }
                            v4 += -1
                        }
                        _ = fi40
                    }
                } else {
                    if false {
                        if true {
                            for fi75 := 0; fi75 < 16; fi75++ {
                                v76 := GenS1{F0: v4}
                                switch v38 {
                                case 44:
                                    var v77 uint16 = 65535
                                    _ = v77
                                case 32768:
                                    v1[0] = true
                                case -2147483648:
                                    var v78 bool = true
                                    _ = v78
                                default:
                                    var v79 string = (("foo bar" + (("" + ("" + "")) + "hello")) + "foo bar")
                                    _ = v79
                                }
                                _ = fi75
                                _ = v76
                            }
                            for fi80 := 0; fi80 < 17; fi80++ {
                                var v81 int64 = 290
                                v81 += (v81 + -330)
                                _ = fi80
                                _ = v81
                            }
                        }
                    } else {
                        fmt.Println(v4, v3)
                        var v82 int = v38
                        _ = v82
                    }
                    var v83 int8 = 0
                    _ = v83
                }
                _ = v38
                _ = v39
            }
        }
    case -9223372036854775808:
        v6 -= 277
    default:
        if (!false) {
            v84 := GenS7{F0: 9223372036854775807}
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v84
        }
    }
    fmt.Println(v6)
    if (3.14159 >= 3.14159) {
        var v85 uint16 = 313
        delete(v1, 2)
        v6 += v6
        _ = v85
    }
    var v86 uint = 65535
    v4 = (-(v4))
    switch "" {
    case "\n\t":
        if true {
            v6 -= v6
            var v87 int32 = 0
            _ = v87
        } else {
            v6 -= 0
            var v88 uint = v86
            _ = v88
        }
    case "a":
        v89 := GenS1{F0: 1}
        _ = v89
    default:
        if (true || false) {
            fmt.Println(v4, v86)
            v90 := v1[4294967295]
            _ = v90
        }
    }
    v2 -= (v2 & v4)
    fmt.Println(v6, v3, v86)
    for fi91 := 0; fi91 < 13; fi91++ {
        var v92 bool = false
        v4 = 110
        _ = fi91
        _ = v92
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v86
}
