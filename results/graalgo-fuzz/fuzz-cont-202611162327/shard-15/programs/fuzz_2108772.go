package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int64 }
type GenS2 struct { F0 int64; F1 uint8; F2 int }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 uint16; F1 int16; F2 string }
type GenS5 struct { F0 int32; F1 uint; F2 uint8 }
type GenS6 struct { F0 int; F1 int32; F2 uint16 }
type GenS7 struct { F0 string; F1 int32; F2 uint }
type GenS8 struct { F0 int8; F1 int; F2 uint32 }
type GenS9 struct { F0 uint16; F1 int32 }

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 4294967295
    v3 := make(map[int]string)
    v4 := make(map[int]bool)
    var v5 uint16 = 32767
    if (-1 <= int32(v5)) {
        switch ("foo bar" + ("" + "hello")) {
        case "\n\t":
            var v6 float32 = 1e10
            _ = v6
        case "":
            if (false || (0.0 < (0.0 * (1e10 * 0.1)))) {
                fmt.Println(v2, v5)
                v7 := GenS1{F0: uint(v5), F1: -9223372036854775808, F2: -1}
                _ = v7
            }
        case "a":
            v8 := []int{9223372036854775807, 42, 4294967295}
            _ = v8
        }
        if (false || true) {
            var v9 uint16 = (v5 << ((65535) & 63))
            var v10 float32 = 0.0
            _ = v10
            _ = v9
        }
    } else {
        v11 := v1[("foo bar" + (("\n\t" + "foo bar") + (("a" + "a") + "foo bar")))]
        var v12 int16 = -778
        _ = v11
        _ = v12
    }
    var v13 uint64 = 656
    var v14 int = -1
    for fi15 := 0; fi15 < 10; fi15++ {
        v2 -= (v2 >> ((2) & 63))
        v16 := []int{fi15, v14}
        v5 -= (^(v5))
        _ = fi15
        _ = v16
    }
    delete(v4, (^(v14)))
    if false {
        v17 := GenS2{F0: 209, F1: 255, F2: v14}
        switch v14 {
        case -900:
            if (0 > 1) {
                switch "hello" {
                case "a":
                    var v18 int32 = 1
                    _ = v18
                case "":
                    v2 -= 2
                case "hello":
                    switch 2147483647 {
                    case 32768:
                        var v19 string = ("foo bar" + "hello")
                        _ = v19
                    case -9223372036854775808:
                        v14 = v14
                    case 1:
                        var v20 int8 = 127
                        _ = v20
                    default:
                        switch (v14 ^ v14) {
                        case 127:
                            var v21 int = v14
                            _ = v21
                        case 9223372036854775807:
                            if true {
                                for fi22 := 0; fi22 < 10; fi22++ {
                                    v23 := GenS1{F0: 4294967295, F1: -1, F2: -1}
                                    if true {
                                        var v24 uint8 = 2
                                        _ = v24
                                    }
                                    _ = fi22
                                    _ = v23
                                }
                            }
                        case 159:
                            fmt.Println(v13, v5)
                        default:
                            v25 := GenS2{F0: 0, F1: 0, F2: (v14 ^ v14)}
                            _ = v25
                        }
                    }
                default:
                    for fi26 := 0; fi26 < 6; fi26++ {
                        if false {
                            v5 = v5
                        }
                        v2 = 65535
                        v27 := GenS3{F0: -1, F1: v2}
                        _ = fi26
                        _ = v27
                    }
                }
            } else {
                var v28 uint16 = (v5 * (v5 - v5))
                switch 128 {
                case -928:
                    for fi29 := 0; fi29 < 2; fi29++ {
                        var v30 int = fi29
                        if false {
                            v31 := GenS2{F0: -1, F1: 0, F2: (fi29 | 9223372036854775807)}
                            fmt.Println(fi29)
                            switch 128 {
                            case 0:
                                if false {
                                    for fi32 := 0; fi32 < 11; fi32++ {
                                        fmt.Println(v30, v13, v28)
                                        _ = fi32
                                    }
                                } else {
                                    v33 := GenS2{F0: 0, F1: 166, F2: 806}
                                    _ = v33
                                }
                            case -873:
                                v14 = fi29
                            case 9223372036854775807:
                                for fi34 := 0; fi34 < 3; fi34++ {
                                    var v35 int = 32768
                                    _ = fi34
                                    _ = v35
                                }
                            default:
                                v13 -= v13
                            }
                            _ = v31
                        }
                        _ = fi29
                        _ = v30
                    }
                case 128:
                    switch (("hello" + (("" + ("hello" + (("" + "foo bar") + ("hello" + ("" + ((("hello" + ("hello" + ("\n\t" + "\n\t"))) + "foo bar") + "a")))))) + ((("hello" + ((("\n\t" + (("a" + ("" + ("\n\t" + ""))) + ("\n\t" + "a"))) + (("a" + "") + "\n\t")) + "\n\t")) + ("hello" + "a")) + "hello"))) + "\n\t") {
                    case "":
                        var v36 int32 = -2147483648
                        _ = v36
                    default:
                        v13 = v13
                    }
                case 2147483647:
                    switch 754 {
                    case 554:
                        for fi37 := 0; fi37 < 9; fi37++ {
                            for fi38 := 0; fi38 < 5; fi38++ {
                                v2 += (v2 * v2)
                                _ = len(v1)
                                if true {
                                    if (0 == 9223372036854775807) {
                                        if false {
                                            v2 -= (v2 ^ (v2 % (((v2 + 0)) | 1)))
                                            v13 = 18446744073709551615
                                            v14 = fi38
                                        }
                                        var v39 int32 = 42
                                        if (int64(v28) < 42) {
                                            v40 := v1[("foo bar" + "hello")]
                                            _ = v40
                                        } else {
                                            fmt.Println(fi38, v5)
                                        }
                                        _ = v39
                                    } else {
                                        var v41 string = "\n\t"
                                        _ = v41
                                    }
                                    v28 = (v28 / ((10) | 1))
                                } else {
                                    var v42 float64 = 0.1
                                    var v43 float64 = 3.14159
                                    var v44 int8 = -1
                                    _ = v42
                                    _ = v43
                                    _ = v44
                                }
                                _ = fi38
                            }
                            v28 = 32767
                            for fi45 := 0; fi45 < 6; fi45++ {
                                if true {
                                    var v46 uint32 = v2
                                    var v47 int16 = -32768
                                    _ = v46
                                    _ = v47
                                } else {
                                    var v48 int = (fi45 ^ v14)
                                    if (!true) {
                                        v28 += v5
                                        v28 -= v28
                                    }
                                    v5 = v5
                                    switch -523 {
                                    case 315:
                                        fmt.Println(v2)
                                    case 127:
                                        var v49 float64 = 3.14159
                                        _ = v49
                                    default:
                                        if (243 != 255) {
                                            switch v14 {
                                            case 0:
                                                v50 := []int{v14, (v48 - v14), fi45, (fi37 & -160)}
                                                _ = v50
                                            case 1:
                                                v5 -= (v5 * (^(v5)))
                                            }
                                        } else {
                                            var v51 uint = 65535
                                            v5 = v28
                                            v52 := v1[""]
                                            _ = v51
                                            _ = v52
                                        }
                                    }
                                    if (0 < -820) {
                                        switch fi37 {
                                        case 128:
                                            fmt.Println(v48)
                                            v53 := GenS4{F0: (^(v5)), F1: 1000, F2: "hello"}
                                            var v54 int8 = 0
                                            _ = v53
                                            _ = v54
                                        case 2147483647:
                                            v55 := v4[(-(fi37))]
                                            _ = v55
                                        case -546:
                                            switch 257 {
                                            case -499:
                                                var v56 bool = true
                                                _ = v56
                                            case 2147483647:
                                                v57 := []int{fi37, (v48 | fi45)}
                                                _ = v57
                                            case 128:
                                                v13 += 18446744073709551615
                                            }
                                        }
                                        if (("a" >= ("" + "hello")) || (1e10 != ((0.1 - -2.25) / float32(-2.25)))) {
                                            v58 := GenS5{F0: 42, F1: 548, F2: 20}
                                            v28 -= 1
                                            _ = v58
                                        } else {
                                            var v59 string = "\n\t"
                                            _ = v59
                                        }
                                        v60 := GenS6{F0: v48, F1: 42, F2: 948}
                                        var v61 float32 = 3.14159
                                        _ = v60
                                        _ = v61
                                    } else {
                                        if (2147483647 > 167) {
                                            var v62 int32 = 2147483647
                                            v63 := GenS7{F0: "\n\t", F1: v62, F2: 1}
                                            _ = v62
                                            _ = v63
                                        } else {
                                            var v64 uint8 = 128
                                            v65 := GenS8{F0: -1, F1: (fi45 + v14), F2: (v2 | (v2 & (^(v2))))}
                                            var v66 uint64 = (v13 - (v13 >> ((uint(v13)) & 63)))
                                            _ = len(v3)
                                            v66 = (v13 % (((v66 - (v13 | v66))) | 1))
                                            _ = v64
                                            _ = v65
                                            _ = v66
                                        }
                                    }
                                    switch int(v5) {
                                    case -1:
                                        var v67 int32 = -1
                                        _ = v67
                                    case 1:
                                        var v68 float32 = 0.1
                                        _ = v68
                                    case -9223372036854775808:
                                        v13 = v13
                                    }
                                    var v69 int32 = -2147483648
                                    v5 += (v28 / ((uint16(v13)) | 1))
                                    _ = v48
                                    _ = v69
                                }
                                _ = fi45
                            }
                            _ = fi37
                        }
                    case -1:
                        if ((-2.25 / (1.5 / 0.1)) == 0.1) {
                            fmt.Println(v14, v2)
                        }
                    }
                }
                _ = v28
            }
        case 580:
            v70 := GenS8{F0: 1, F1: (-(v14)), F2: (v2 + uint32(v13))}
            _ = v70
        }
        _ = v17
    }
    if false {
        v3[(v14 - (v14 - (v14 ^ (v14 << ((957) & 63)))))] = "hello"
        var v71 string = ("a" + "a")
        _ = v71
    }
    v5 = v5
    var v72 int64 = 42
    delete(v3, (v14 & int(v2)))
    for fi73 := 0; fi73 < 3; fi73++ {
        v2 = v2
        if true {
            if (!(v2 <= (^(v2)))) {
                if (!(float64((0.0 * 0.1)) == (0.0 + (((1e10 + 3.14159) / -2.25) * 0.0)))) {
                    if (1000 != 862) {
                        if ((v5 % ((v5) | 1)) != (^(v5))) {
                            if (!false) {
                                v72 = v72
                                var v74 bool = false
                                _ = v74
                            } else {
                                var v75 int16 = -759
                                for fi76 := 0; fi76 < 9; fi76++ {
                                    switch 128 {
                                    case 185:
                                        if (((v72 & v72) < (v72 ^ (v72 | (v72 & (^(v72)))))) && false) {
                                            fmt.Println(v13, v2, v5)
                                            var v77 float64 = 0.1
                                            _ = v77
                                        }
                                    case 42:
                                        var v78 float32 = (1e10 + 1e10)
                                        v79 := []int{}
                                        if true {
                                            if false {
                                                v5 = 646
                                            } else {
                                                v80 := GenS8{F0: 1, F1: -1, F2: (v2 / (((v2 >> ((0) & 63))) | 1))}
                                                var v81 float64 = 1.5
                                                v79 = append(v79, (fi73 + (v14 + 1)))
                                                _ = v80
                                                _ = v81
                                            }
                                            var v82 uint8 = 2
                                            fmt.Println(v14, fi76, v78)
                                            _ = v82
                                        }
                                        _ = v78
                                        _ = v79
                                    default:
                                        v83 := []int{}
                                        _ = v83
                                    }
                                    var v84 float32 = (((float32(1e10) + 1e10) * -2.25) / 3.14159)
                                    var v85 string = "\n\t"
                                    if true {
                                        v75 = int16(v13)
                                    }
                                    v72 = v72
                                    _ = fi76
                                    _ = v84
                                    _ = v85
                                }
                                _ = v75
                            }
                        } else {
                            var v86 float64 = 3.14159
                            v87 := GenS9{F0: uint16(v13), F1: int32(v72)}
                            _ = v86
                            _ = v87
                        }
                        v72 -= (v72 * (v72 + v72))
                        v2 -= uint32(v13)
                    } else {
                        var v88 int = (fi73 | (^(fi73)))
                        _ = v88
                    }
                    if false {
                        var v89 uint8 = 128
                        for fi90 := 0; fi90 < 16; fi90++ {
                            fmt.Println(v13)
                            fmt.Println(fi90)
                            switch (("foo bar" + ("\n\t" + ("" + ((("\n\t" + "a") + "hello") + "")))) + "") {
                            case "foo bar":
                                fmt.Println(fi73, v72)
                            default:
                                if true {
                                    var v91 float64 = -2.25
                                    _ = v91
                                } else {
                                    fmt.Println(v13, v14)
                                    v92 := []int{1, (^(v14)), 647, (fi73 & (v14 - v14))}
                                    _ = v92
                                }
                            }
                            v72 += 1
                            v72 = 859
                            _ = fi90
                        }
                        _ = v89
                    }
                }
            } else {
                fmt.Println(v14, v5, v14)
            }
        }
        v14 += 619
        _ = fi73
    }
    v13 -= (v13 + v13)
    var v93 int64 = v72
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v72:", v72)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v72
    _ = v93
}
