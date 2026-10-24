package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 uint8; F1 uint16; F2 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8; F1 string }
type GenS6 struct { F0 uint; F1 int; F2 uint16 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int32; F1 uint8 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int }
func genFn0(p0 int16, p1 uint16) int16 {
    v5 := []int{-1, 0, -927}
    for fi6 := 0; fi6 < 8; fi6++ {
        for fi7 := 0; fi7 < 7; fi7++ {
            if (((1.5 < ((3.14159 / float32(0.1)) * 0.0)) || true) && (2 >= 569)) {
                var v8 int8 = 127
                if false {
                    fmt.Println(v8, fi7)
                    var v9 uint16 = (p1 | p1)
                    if true {
                        if ((-2.25 + 0.0) != 0.1) {
                            var v10 uint32 = 65535
                            _ = v10
                        } else {
                            if (42 < 127) {
                                var v11 string = "\n\t"
                                _ = v11
                            } else {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                            v8 = v8
                            fmt.Println(p1)
                        }
                        v12 := GenS1{F0: fi6, F1: "\n\t"}
                        if false {
                            var v13 string = "foo bar"
                            _ = v13
                        }
                        _ = v12
                    } else {
                        v14 := GenS1{F0: fi6, F1: "hello"}
                        _ = v14
                    }
                    _ = v9
                }
                v15 := GenS1{F0: 107, F1: "hello"}
                _ = v15
                _ = v8
            } else {
                v5 = append(v5, fi6)
                p1 = 255
            }
            if false {
                if (901 > 1) {
                    if ((2147483648 <= 4294967295) || true) {
                        var v16 int16 = (p0 - p0)
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        v16 -= (p0 ^ 1000)
                        _ = v16
                    }
                    _ = len(v5)
                    if true {
                        p0 = -32768
                    }
                } else {
                    p1 -= 255
                }
                var v17 uint8 = 26
                v18 := GenS1{F0: (fi7 / (((fi6 & (fi6 ^ (fi6 | fi7)))) | 1)), F1: "a"}
                _ = v17
                _ = v18
            }
            _ = fi7
        }
        p1 -= 960
        _ = fi6
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v5
    return p0
}
func genFn1(p0 int, p1 int8, p2 uint8) int16 {
    var v19 uint8 = (p2 >> ((1) & 63))
    v20 := GenS2{F0: v19, F1: 394, F2: ""}
    v21 := []int{p0, 0, -87, (^(p0))}
    fmt.Println(p2, p0, p1)
    var v22 int = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v20
    _ = v21
    _ = v22
    return 1
}

func main() {
    v1 := []int{123}
    var v2 int8 = -128
    var v3 bool = true
    var v4 int16 = 1000
    if (-1 > 2147483647) {
        if (934 == 353) {
            v2 += (v2 / ((52) | 1))
        }
    } else {
        for fi23 := 0; fi23 < 15; fi23++ {
            switch "hello" {
            case "a":
                if (-459 >= -9223372036854775808) {
                    switch "\n\t" {
                    case "hello":
                        fmt.Println(fi23, v3)
                    case "\n\t":
                        var v24 string = ("foo bar" + "\n\t")
                        _ = v24
                    case "foo bar":
                        for fi25 := 0; fi25 < 3; fi25++ {
                            _ = len(v1)
                            switch -590 {
                            case 128:
                                var v26 uint64 = 0
                                _ = v26
                            case 2147483647:
                                fmt.Println(v2)
                            default:
                                v2 = v2
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = fi25
                        }
                    }
                    for fi27 := 0; fi27 < 16; fi27++ {
                        fmt.Println(v4, v4)
                        v28 := GenS3{F0: ("a" + "")}
                        _ = fi27
                        _ = v28
                    }
                    var v29 uint64 = 2
                    _ = v29
                } else {
                    if false {
                        switch (fi23 / (((fi23 & (fi23 - 100))) | 1)) {
                        case -495:
                            v1 = append(v1, fi23)
                        case 4294967295:
                            fmt.Println(fi23)
                        case 678:
                            fmt.Println(v2, v2, fi23)
                        default:
                            fmt.Println(v3, v4)
                        }
                        fmt.Println(v3, v2)
                        fmt.Println(v4, v3)
                    }
                    var v30 uint = 328
                    v31 := GenS4{F0: 1}
                    _ = v30
                    _ = v31
                }
            default:
                if true {
                    v32 := GenS4{F0: 825}
                    switch "" {
                    case "hello":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case "a":
                        v2 += v2
                    case "foo bar":
                        var v33 int = (^(fi23))
                        _ = v33
                    default:
                        fmt.Println(v3)
                    }
                    _ = v32
                } else {
                    switch (("foo bar" + "foo bar") + "a") {
                    case "hello":
                        var v34 uint64 = 1
                        _ = v34
                    default:
                        if false {
                            for fi35 := 0; fi35 < 6; fi35++ {
                                v4 = genFn0(-1, 65535)
                                if (!true) {
                                    v36 := GenS1{F0: 32768, F1: ("\n\t" + "")}
                                    var v37 int8 = (v2 / ((v2) | 1))
                                    _ = len(v1)
                                    _ = v36
                                    _ = v37
                                } else {
                                    if v3 {
                                        var v38 uint = 2
                                        fmt.Println(fi35, v2)
                                        var v39 int32 = 0
                                        _ = v38
                                        _ = v39
                                    } else {
                                        var v40 int32 = 1
                                        if false {
                                            v41 := GenS2{F0: 0, F1: 65535, F2: "foo bar"}
                                            var v42 float32 = -2.25
                                            _ = v41
                                            _ = v42
                                        }
                                        if v3 {
                                            var v43 uint32 = 0
                                            switch ("a" + ("foo bar" + ("\n\t" + "a"))) {
                                            case "hello":
                                                v40 += (v40 ^ (v40 & (v40 >> ((255) & 63))))
                                            case "":
                                                v3 = ((-9223372036854775808 >= -534) || ((((v2 << ((2) & 63)) <= (v2 % ((100) | 1))) || v3) || ((v4 | v4) < (^(v4)))))
                                            default:
                                                v2 = (-(v2))
                                            }
                                            var v44 uint8 = 128
                                            _ = v43
                                            _ = v44
                                        }
                                        v45 := GenS2{F0: 104, F1: 2, F2: "\n\t"}
                                        _ = v40
                                        _ = v45
                                    }
                                }
                                _ = fi35
                            }
                        } else {
                            fmt.Println(fi23, v3, fi23)
                            _ = len(v1)
                        }
                    }
                    if (111 == (v2 & (v2 / ((v2) | 1)))) {
                        var v46 int = fi23
                        if false {
                            v47 := GenS5{F0: (v2 + v2), F1: ("" + "\n\t")}
                            for fi48 := 0; fi48 < 20; fi48++ {
                                var v49 uint8 = 1
                                v50 := GenS3{F0: "foo bar"}
                                _ = fi48
                                _ = v49
                                _ = v50
                            }
                            for fi51 := 0; fi51 < 10; fi51++ {
                                if v3 {
                                    v52 := GenS3{F0: (("hello" + "a") + "")}
                                    _ = v52
                                }
                                v53 := GenS4{F0: 0}
                                var v54 bool = ((v4 / ((v4) | 1)) >= (^(v4)))
                                _ = fi51
                                _ = v53
                                _ = v54
                            }
                            _ = v47
                        }
                        _ = v46
                    }
                }
            }
            switch fi23 {
            case 2147483647:
                fmt.Println(v4, fi23)
            case -194:
                if (!(!(("hello" + "") > "\n\t"))) {
                    v3 = v3
                }
            }
            var v55 int = 32767
            _ = fi23
            _ = v55
        }
        for fi56 := 0; fi56 < 9; fi56++ {
            var v57 uint = 434
            v58 := GenS6{F0: 256, F1: fi56, F2: 32767}
            var v59 string = "foo bar"
            _ = fi56
            _ = v57
            _ = v58
            _ = v59
        }
        fmt.Println(v3, v4)
    }
    v2 -= 24
    v4 += v4
    if ((v2 * int8(v4)) >= v2) {
        var v60 int = 32768
        v60 = v60
        v61 := GenS1{F0: (-(v60)), F1: "a"}
        _ = v60
        _ = v61
    } else {
        v2 += (v2 | v2)
        for fi62 := 0; fi62 < 11; fi62++ {
            switch (("a" + "foo bar") + ("\n\t" + "hello")) {
            case "a":
                fmt.Println(fi62, v2, v2)
            case "hello":
                var v63 int8 = 127
                _ = v63
            case "":
                fmt.Println(v4, fi62)
            default:
                if false {
                    v3 = false
                    var v64 bool = (0 < 642)
                    fmt.Println(v4, v2, v2)
                    _ = v64
                }
            }
            if true {
                var v65 int64 = -1
                var v66 uint64 = 18446744073709551615
                _ = v65
                _ = v66
            }
            for fi67 := 0; fi67 < 5; fi67++ {
                var v68 int16 = (v4 ^ (v4 + 1))
                _ = fi67
                _ = v68
            }
            _ = fi62
        }
    }
    for fi69 := 0; fi69 < 5; fi69++ {
        v70 := GenS7{F0: 65535}
        for fi71 := 0; fi71 < 15; fi71++ {
            if (2 != 255) {
                if v3 {
                    fmt.Println(fi69)
                    if true {
                        v4 = v4
                        var v72 int16 = v4
                        v73 := GenS8{F0: 816, F1: 42}
                        _ = v72
                        _ = v73
                    }
                } else {
                    v3 = v3
                }
                if v3 {
                    v3 = (((("" + "hello") + "\n\t") + "hello") != ("" + "foo bar"))
                    v4 += (^(v4))
                    switch (^(fi69)) {
                    case -9223372036854775808:
                        fmt.Println(fi71, v2)
                    case 1:
                        fmt.Println(fi69)
                    case -354:
                        v74 := GenS6{F0: 2, F1: (fi71 | (fi69 >> ((0) & 63))), F2: 644}
                        _ = v74
                    }
                }
            } else {
                v3 = v3
                if len(v1) > 0 {
                    v1[0] = fi71
                }
            }
            _ = fi71
        }
        _ = fi69
        _ = v70
    }
    var v75 uint8 = 1
    var v76 float64 = 0.1
    fmt.Println(v76, v2, v76)
    v2 -= v2
    switch 32767 {
    case -720:
        for fi77 := 0; fi77 < 12; fi77++ {
            v75 = v75
            switch fi77 {
            case -591:
                v78 := GenS3{F0: (("hello" + "") + "hello")}
                _ = v78
            default:
                if false {
                    var v79 uint8 = v75
                    for fi80 := 0; fi80 < 20; fi80++ {
                        fmt.Println(v3, v3)
                        _ = fi80
                    }
                    v3 = true
                    _ = v79
                } else {
                    for fi81 := 0; fi81 < 10; fi81++ {
                        var v82 uint16 = 255
                        _ = fi81
                        _ = v82
                    }
                }
            }
            var v83 float32 = (float32(1.5) * 0.0)
            _ = fi77
            _ = v83
        }
    case 42:
        if true {
            fmt.Println(v75, v4, v3)
            var v84 int16 = (-(v4))
            _ = v84
        } else {
            var v85 uint8 = 184
            var v86 uint16 = 255
            _ = v85
            _ = v86
        }
    default:
        v76 = 0.0
    }
    switch 32768 {
    case 9223372036854775807:
        if (3.14159 < -2.25) {
            for fi87 := 0; fi87 < 4; fi87++ {
                var v88 int64 = 42
                _ = fi87
                _ = v88
            }
            if (((3.14159 / float32(1.5)) + 0.0) != 0.0) {
                var v89 uint32 = 2147483647
                var v90 int32 = 0
                for fi91 := 0; fi91 < 13; fi91++ {
                    if false {
                        var v92 string = ("\n\t" + "foo bar")
                        for fi93 := 0; fi93 < 11; fi93++ {
                            fmt.Println(v4)
                            _ = fi93
                        }
                        v94 := GenS9{F0: (^(v4))}
                        _ = v92
                        _ = v94
                    } else {
                        for fi95 := 0; fi95 < 4; fi95++ {
                            v96 := GenS1{F0: 1, F1: "foo bar"}
                            var v97 bool = v3
                            _ = fi95
                            _ = v96
                            _ = v97
                        }
                    }
                    fmt.Println(v4, v2, v3)
                    _ = len(v1)
                    _ = fi91
                }
                _ = v89
                _ = v90
            } else {
                var v98 uint64 = 2
                v76 = 1.5
                _ = v98
            }
        }
    case -475:
        v99 := GenS10{F0: 0}
        _ = v99
    case 525:
        var v100 int8 = v2
        _ = v100
    default:
        switch 9223372036854775807 {
        case 32768:
            var v101 string = "a"
            _ = v101
        case 0:
            v75 = v75
        default:
            fmt.Println(v76, v76, v75)
        }
    }
    if len(v1) > 0 {
        v1[0] = 128
    }
    var v102 string = ("" + "foo bar")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v102:", v102)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v102
    _ = v2
    _ = v3
    _ = v4
    _ = v75
    _ = v76
}
