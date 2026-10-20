package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 int8; F1 int16; F2 uint64 }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 uint64; F1 uint16; F2 string }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int16; F1 uint8; F2 int8 }
type GenS9 struct { F0 int32; F1 int8; F2 int }
func genFn0(p0 uint32) uint16 {
    for fi6 := 0; fi6 < 9; fi6++ {
        v7 := GenS1{F0: 65535, F1: 1}
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 255, F1: 588}
    var v9 int64 = -434
    v10 := GenS2{F0: -57, F1: 1000, F2: 18446744073709551615}
    for fi11 := 0; fi11 < 16; fi11++ {
        p0 = p0
        v12 := GenS1{F0: 588, F1: 548}
        var v13 uint32 = 2147483647
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return 539
}
func genFn1(p0 uint8, p1 uint8, p2 uint8) int8 {
    v14 := GenS2{F0: -13, F1: -1, F2: 2}
    v15 := GenS1{F0: 0, F1: 158}
    if (1e10 < float64(1.5)) {
        v16 := []int{553}
        switch -949 {
        case -1:
            for fi17 := 0; fi17 < 9; fi17++ {
                for fi18 := 0; fi18 < 7; fi18++ {
                    switch "foo bar" {
                    case "hello":
                        v19 := GenS2{F0: 1, F1: 1, F2: 0}
                        _ = v19
                    case "":
                        var v20 uint16 = 1
                        _ = v20
                    }
                    if (32767 == genFn0(2147483648)) {
                        if false {
                            switch 9223372036854775807 {
                            case 247:
                                var v21 float64 = 1e10
                                _ = v21
                            }
                        } else {
                            var v22 int16 = -520
                            switch "\n\t" {
                            case "a":
                                if true {
                                    p0 = (p0 + p2)
                                } else {
                                    switch "\n\t" {
                                    case "\n\t":
                                        switch fi18 {
                                        case 4294967295:
                                            v22 = (v22 >> ((65535) & 63))
                                        }
                                    }
                                    switch ("a" + "") {
                                    case "a":
                                        v16 = append(v16, fi17)
                                    case "\n\t":
                                        p1 -= p1
                                    case "":
                                        var v23 int64 = 1
                                        _ = v23
                                    default:
                                        v24 := GenS1{F0: 949, F1: 2}
                                        _ = v24
                                    }
                                }
                            case "\n\t":
                                v22 += (-(v22))
                            default:
                                v25 := GenS3{F0: 2147483648, F1: 2147483647}
                                _ = v25
                            }
                            _ = v22
                        }
                        if true {
                            switch -1 {
                            case -9223372036854775808:
                                var v26 uint16 = genFn0(108)
                                _ = v26
                            case -809:
                                var v27 bool = true
                                _ = v27
                            case -1:
                                if len(v16) > 0 {
                                    v16[0] = fi17
                                }
                            }
                            fmt.Println(fi18)
                        }
                        if false {
                            switch "\n\t" {
                            case "foo bar":
                                var v28 int8 = 73
                                _ = v28
                            case "hello":
                                if true {
                                    p0 = p0
                                    fmt.Println(p2, p2, p0)
                                    v29 := GenS4{F0: 143, F1: 255}
                                    _ = v29
                                } else {
                                    fmt.Println(p2, p0, p0)
                                }
                            case "\n\t":
                                v30 := GenS2{F0: -1, F1: -32768, F2: 4294967295}
                                _ = v30
                            }
                        }
                    }
                    _ = fi18
                }
                p2 = (p2 | (p1 % ((p2) | 1)))
                for fi31 := 0; fi31 < 18; fi31++ {
                    fmt.Println(fi31)
                    p1 += p2
                    fmt.Println(p2)
                    _ = fi31
                }
                _ = fi17
            }
        case 4294967295:
            if (!((((1e10 + 1.5) / 1.0) - 1.5) <= -2.25)) {
                for fi32 := 0; fi32 < 5; fi32++ {
                    switch ("" + "") {
                    case "foo bar":
                        for fi33 := 0; fi33 < 20; fi33++ {
                            fmt.Println(fi32, p0, fi32)
                            p1 = p0
                            _ = fi33
                        }
                    default:
                        switch (-(fi32)) {
                        case 9223372036854775807:
                            p0 = (p1 >> ((4294967295) & 63))
                        case -2147483648:
                            var v34 float64 = 0.0
                            _ = v34
                        }
                    }
                    if ("foo bar" != "hello") {
                        fmt.Println(p1)
                    }
                    p2 -= p1
                    _ = fi32
                }
            }
        case 2147483647:
            if (577 < 362) {
                p1 = p0
            } else {
                switch -1 {
                case -585:
                    v35 := GenS5{F0: 1, F1: 42, F2: "hello"}
                    _ = v35
                case 1:
                    if ((p2 / ((p1) | 1)) == p1) {
                        p2 += p0
                        v36 := GenS4{F0: 469, F1: uint(p2)}
                        for fi37 := 0; fi37 < 6; fi37++ {
                            for fi38 := 0; fi38 < 4; fi38++ {
                                _ = len(v16)
                                _ = len(v16)
                                v16 = append(v16, -1)
                                _ = fi38
                            }
                            var v39 uint8 = (p1 + 127)
                            _ = fi37
                            _ = v39
                        }
                        _ = v36
                    } else {
                        var v40 string = (("hello" + "hello") + "")
                        _ = v40
                    }
                default:
                    for fi41 := 0; fi41 < 18; fi41++ {
                        if (825 <= 2) {
                            switch ("a" + ("foo bar" + "foo bar")) {
                            case "\n\t":
                                var v42 float32 = float32(3.14159)
                                _ = v42
                            case "a":
                                if false {
                                    if (true && true) {
                                        p0 -= 16
                                    } else {
                                        p0 = (p0 << ((42) & 63))
                                        var v43 int64 = 9223372036854775807
                                        _ = v43
                                    }
                                }
                            case "hello":
                                for fi44 := 0; fi44 < 12; fi44++ {
                                    p1 += p1
                                    var v45 uint16 = genFn0(4294967295)
                                    _ = fi44
                                    _ = v45
                                }
                            }
                            fmt.Println(p2, p2)
                        }
                        p0 += (p2 % ((p2) | 1))
                        switch (fi41 ^ (fi41 << ((255) & 63))) {
                        case -9223372036854775808:
                            switch ((("foo bar" + ((("a" + "foo bar") + "") + "hello")) + "a") + "foo bar") {
                            case "\n\t":
                                fmt.Println(p0, p1, fi41)
                            default:
                                _ = len(v16)
                            }
                        case 32768:
                            fmt.Println(p1, p1, fi41)
                        case -282:
                            var v46 int16 = 42
                            _ = v46
                        }
                        _ = fi41
                    }
                }
                if false {
                    if len(v16) > 0 {
                        v16[0] = 102
                    }
                    p0 = 127
                    if true {
                        for fi47 := 0; fi47 < 10; fi47++ {
                            var v48 uint8 = (p0 % ((p0) | 1))
                            _ = len(v16)
                            if len(v16) >= 2 {
                                v16 = v16[1:len(v16)]
                            }
                            _ = fi47
                            _ = v48
                        }
                        v49 := GenS6{F0: 585}
                        _ = v49
                    } else {
                        p1 -= (^(p0))
                        var v50 int32 = 978
                        _ = v50
                    }
                } else {
                    p0 -= (p1 - (^(p1)))
                }
                v51 := GenS5{F0: 352, F1: 2, F2: "\n\t"}
                _ = v51
            }
        default:
            switch ("a" + ("foo bar" + "")) {
            case "a":
                if false {
                    p0 += 42
                } else {
                    p1 = (p2 / ((p1) | 1))
                }
            case "":
                switch 1 {
                case -465:
                    fmt.Println(p2, p1, p2)
                default:
                    switch 2147483647 {
                    case -1:
                        var v52 string = ("a" + ("\n\t" + "\n\t"))
                        _ = v52
                    default:
                        var v53 int32 = -2147483648
                        _ = v53
                    }
                }
            default:
                for fi54 := 0; fi54 < 9; fi54++ {
                    if (2147483647 < 116) {
                        var v55 int32 = 42
                        _ = v55
                    } else {
                        var v56 int = -1
                        _ = v56
                    }
                    switch -547 {
                    case -124:
                        _ = len(v16)
                    default:
                        var v57 float32 = ((float32(-2.25) / 1.0) - ((0.0 / 1.5) * (float32(1.5) - 1e10)))
                        _ = v57
                    }
                    v58 := GenS4{F0: 0, F1: 4294967295}
                    _ = fi54
                    _ = v58
                }
            }
        }
        _ = v16
    }
    for fi59 := 0; fi59 < 18; fi59++ {
        for fi60 := 0; fi60 < 9; fi60++ {
            p1 += 127
            if true {
                var v61 uint32 = 1
                fmt.Println(v61)
                v62 := GenS3{F0: (v61 + (v61 + (v61 - (v61 | (v61 * v61))))), F1: 825}
                _ = v61
                _ = v62
            } else {
                var v63 bool = false
                v64 := []int{(fi60 + 4294967295)}
                v65 := GenS7{F0: (fi59 >> ((4294967295) & 63))}
                _ = v63
                _ = v64
                _ = v65
            }
            var v66 bool = (!(true && false))
            _ = fi60
            _ = v66
        }
        _ = fi59
    }
    v67 := []int{862, 2147483647, 2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    _ = v67
    return -128
}

func main() {
    v1 := []int{}
    v2 := []int{666, 4294967295, 970}
    var v3 int32 = 1
    v4 := make(map[int]int)
    v5 := []uint32{4294967295}
    for fi68 := 0; fi68 < 6; fi68++ {
        for fi69 := 0; fi69 < 18; fi69++ {
            var v70 int16 = 1000
            v71 := v4[fi69]
            _ = fi69
            _ = v70
            _ = v71
        }
        _ = fi68
    }
    v72 := GenS8{F0: int16(v3), F1: 2, F2: 100}
    var v73 int64 = -9223372036854775808
    v3 = 2147483647
    fmt.Println(v3, v3)
    var v74 string = (("\n\t" + "") + "hello")
    var v75 uint64 = 2
    fmt.Println(v75, v75)
    v74 = (((v74 + (v74 + v74)) + "a") + v74)
    var v76 int32 = (v3 & (v3 ^ -2147483648))
    for fi77 := 0; fi77 < 3; fi77++ {
        var v78 float64 = 0.0
        if (((v74 + (v74 + v74)) + v74) == ("hello" + v74)) {
            for fi79 := 0; fi79 < 2; fi79++ {
                v73 = v73
                fmt.Println(v74)
                switch (fi79 & (fi79 ^ (fi77 % ((fi77) | 1)))) {
                case 127:
                    v73 -= (-(v73))
                    if true {
                        var v80 int = int(v73)
                        var v81 float32 = 0.0
                        v82 := GenS6{F0: 2}
                        _ = v80
                        _ = v81
                        _ = v82
                    }
                case 4294967295:
                    switch fi77 {
                    case 4294967295:
                        var v83 bool = true
                        _ = v83
                    case -516:
                        v75 = v75
                    case 30:
                        v76 += (v3 + v76)
                    default:
                        v84 := GenS5{F0: v75, F1: 2, F2: ""}
                        _ = v84
                    }
                }
                v85 := GenS9{F0: 0, F1: 100, F2: fi77}
                v86 := v4[fi79]
                _ = fi79
                _ = v85
                _ = v86
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        _ = fi77
        _ = v78
    }
    var v87 int8 = 1
    var v88 int32 = (^(v3))
    v88 = v3
    v74 = (v74 + "hello")
    for fi89 := 0; fi89 < 13; fi89++ {
        if true {
            v75 -= 713
            v90 := GenS4{F0: (v75 / (((v75 + 42)) | 1)), F1: 767}
            var v91 float64 = 0.1
            _ = v90
            _ = v91
        }
        for fi92 := 0; fi92 < 5; fi92++ {
            var v93 uint32 = 931
            var v94 uint32 = (^(v93))
            var v95 int = fi89
            _ = fi92
            _ = v93
            _ = v94
            _ = v95
        }
        _ = fi89
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    fmt.Println("v87:", v87)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v72
    _ = v73
    _ = v74
    _ = v75
    _ = v76
    _ = v87
    _ = v88
}
