package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 int32; F1 uint32 }
type GenS6 struct { F0 string; F1 uint64 }
type GenS7 struct { F0 string; F1 bool; F2 string }
type GenS8 struct { F0 int; F1 uint8 }
type GenS9 struct { F0 int }

func main() {
    v1 := []int{-116}
    var v2 uint = 2
    v3 := []uint32{970, 0, 4294967295}
    var v4 uint32 = 2147483648
    v5 := GenS1{F0: 1, F1: v4, F2: 2147483647}
    fmt.Println(v4, v4)
    v6 := GenS1{F0: 1, F1: (v4 - v4), F2: -470}
    var v7 string = "foo bar"
    var v8 int8 = 0
    v9 := GenS2{F0: 42}
    for fi10 := 0; fi10 < 5; fi10++ {
        if (!(!(true || (!false)))) {
            v11 := GenS1{F0: 42, F1: 51, F2: (fi10 + 712)}
            _ = v11
        } else {
            if ((v2 + (v2 ^ v2)) >= (^(v2))) {
                for fi12 := 0; fi12 < 3; fi12++ {
                    if true {
                        if (-375 <= 0) {
                            v2 = v2
                            v2 -= v2
                        } else {
                            _ = len(v3)
                            switch 128 {
                            case 9223372036854775807:
                                if ((!true) || true) {
                                    var v13 float64 = 0.1
                                    _ = v13
                                } else {
                                    var v14 uint32 = v4
                                    if len(v1) > 0 {
                                        v1[0] = 1
                                    }
                                    if true {
                                        fmt.Println(fi10, v7, fi12)
                                        fmt.Println(v4)
                                        if ((v8 >> (((v2 + uint(v14))) & 63)) >= v8) {
                                            if len(v3) > 0 {
                                                v3[0] = v14
                                            }
                                        } else {
                                            fmt.Println(fi10, fi10, v14)
                                            switch v7 {
                                            case "\n\t":
                                                v15 := GenS1{F0: uint8(v2), F1: 0, F2: -9223372036854775808}
                                                _ = v15
                                            }
                                        }
                                    } else {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        switch (^(fi12)) {
                                        case 42:
                                            v3 = append(v3, 320)
                                        case -1:
                                            switch v7 {
                                            case "\n\t":
                                                var v16 uint = v2
                                                _ = v16
                                            case "a":
                                                if len(v1) > 0 {
                                                    v1[0] = (fi12 - 127)
                                                }
                                            case "hello":
                                                var v17 uint = (v2 * v2)
                                                _ = v17
                                            }
                                        default:
                                            v4 = v14
                                        }
                                    }
                                    _ = v14
                                }
                                var v18 uint32 = 0
                                var v19 uint32 = (v18 / ((v18) | 1))
                                _ = v18
                                _ = v19
                            case 128:
                                if ((true || ((53 != 169) || (1 != 1))) && true) {
                                    if false {
                                        v2 -= uint(v4)
                                    }
                                } else {
                                    v4 -= 2147483647
                                    fmt.Println(fi12, fi12)
                                }
                            case 196:
                                var v20 float64 = 1e10
                                _ = v20
                            default:
                                v21 := GenS1{F0: 0, F1: 42, F2: (-(fi10))}
                                _ = v21
                            }
                        }
                        if (false || false) {
                            var v22 bool = true
                            switch v7 {
                            case "hello":
                                _ = len(v1)
                            case "":
                                var v23 uint64 = 523
                                _ = v23
                            case "\n\t":
                                if (true || (966 < 171)) {
                                    if false {
                                        v8 = v8
                                    } else {
                                        v24 := GenS3{F0: v2, F1: (0 <= 42)}
                                        v7 = v7
                                        v8 = 1
                                        _ = v24
                                    }
                                    v25 := GenS1{F0: 1, F1: v4, F2: (fi10 >> (((v2 - v2)) & 63))}
                                    _ = v25
                                }
                            default:
                                if v22 {
                                    var v26 int32 = 42
                                    v1 = append(v1, (^(fi10)))
                                    _ = v26
                                }
                            }
                            _ = v22
                        }
                        switch (v7 + v7) {
                        case "\n\t":
                            v27 := GenS4{F0: (v4 * (v4 >> ((v2) & 63))), F1: 0}
                            _ = v27
                        case "foo bar":
                            v7 = (v7 + v7)
                        case "a":
                            v4 -= v4
                        }
                        var v28 uint8 = 0
                        _ = v28
                    }
                    if len(v3) > 0 {
                        v3[0] = 659
                    }
                    _ = fi12
                }
                if (v2 == v2) {
                    var v29 string = v7
                    var v30 float32 = (0.0 / 3.14159)
                    v4 = (v4 ^ (v4 % ((42) | 1)))
                    _ = v29
                    _ = v30
                }
                v31 := GenS5{F0: -1, F1: (^(v4))}
                _ = v31
            }
            v2 = 65535
            if false {
                if false {
                    switch fi10 {
                    case 4294967295:
                        switch (-(fi10)) {
                        case -2147483648:
                            var v32 float32 = 3.14159
                            _ = v32
                        case -319:
                            var v33 uint = (v2 | v2)
                            _ = v33
                        case 9223372036854775807:
                            fmt.Println(v4, v8)
                        default:
                            for fi34 := 0; fi34 < 8; fi34++ {
                                v7 = v7
                                _ = fi34
                            }
                        }
                    case 42:
                        switch (fi10 << (((v2 % (((v2 - (v2 % ((v2) | 1)))) | 1))) & 63)) {
                        case -433:
                            v35 := GenS2{F0: 479}
                            _ = v35
                        case 32767:
                            var v36 uint8 = 2
                            _ = v36
                        case -632:
                            fmt.Println(fi10)
                        default:
                            _ = len(v3)
                        }
                    case 0:
                        for fi37 := 0; fi37 < 7; fi37++ {
                            v7 = v7
                            _ = fi37
                        }
                    default:
                        switch "hello" {
                        case "foo bar":
                            var v38 int64 = -1
                            _ = v38
                        default:
                            for fi39 := 0; fi39 < 12; fi39++ {
                                if (true || false) {
                                    if len(v1) > 0 {
                                        v1[0] = 329
                                    }
                                    var v40 uint16 = 42
                                    if (("hello" + ("a" + v7)) <= "foo bar") {
                                        v4 -= (v4 - v4)
                                    } else {
                                        switch "hello" {
                                        case "":
                                            if len(v1) > 0 {
                                                v1[0] = -346
                                            }
                                        case "a":
                                            fmt.Println(v40, v7)
                                        default:
                                            if len(v1) > 0 {
                                                v1[0] = (fi39 & (-(fi10)))
                                            }
                                        }
                                        var v41 string = ("" + v7)
                                        fmt.Println(v2, v4, v7)
                                        _ = v41
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    switch "hello" {
                                    case "hello":
                                        if false {
                                            var v42 float32 = 1.5
                                            var v43 int = fi39
                                            switch ((v7 + v7) + "") {
                                            case "hello":
                                                if len(v1) > 0 {
                                                    v1[0] = -592
                                                }
                                            }
                                            _ = v42
                                            _ = v43
                                        } else {
                                            v2 += v2
                                        }
                                        _ = len(v3)
                                    }
                                    var v44 int = 127
                                    _ = v40
                                    _ = v44
                                }
                                var v45 int32 = 0
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                fmt.Println(v4, v2, fi10)
                                _ = fi39
                                _ = v45
                            }
                        }
                    }
                } else {
                    var v46 uint8 = 1
                    _ = v46
                }
                if false {
                    var v47 int32 = 0
                    _ = v47
                } else {
                    if (0 >= 162) {
                        v48 := GenS6{F0: v7, F1: 0}
                        fmt.Println(v2, v2)
                        _ = v48
                    } else {
                        var v49 float64 = 3.14159
                        _ = v49
                    }
                    v50 := GenS2{F0: -9223372036854775808}
                    _ = v50
                }
            }
        }
        var v51 bool = true
        v52 := GenS7{F0: v7, F1: true, F2: "\n\t"}
        _ = fi10
        _ = v51
        _ = v52
    }
    for fi53 := 0; fi53 < 20; fi53++ {
        if false {
            if false {
                if ((v2 <= v2) || false) {
                    var v54 uint8 = 255
                    v55 := GenS7{F0: ("foo bar" + (v7 + v7)), F1: true, F2: (v7 + "foo bar")}
                    _ = v54
                    _ = v55
                }
            } else {
                for fi56 := 0; fi56 < 18; fi56++ {
                    if ((-1 != 42) || (true || false)) {
                        if (!(841 < 1000)) {
                            var v57 uint = 1
                            v58 := GenS8{F0: fi56, F1: 255}
                            _ = v57
                            _ = v58
                        } else {
                            fmt.Println(v7, v7)
                        }
                        switch v7 {
                        case "hello":
                            switch (v7 + v7) {
                            case "\n\t":
                                v4 += 0
                                fmt.Println(fi53)
                                var v59 int = (fi56 % (((fi56 & 4294967295)) | 1))
                                if (639 > (^(v4))) {
                                    var v60 uint32 = v4
                                    switch v7 {
                                    case "hello":
                                        if true {
                                            if true {
                                                if (-32768 == 1000) {
                                                    v8 -= v8
                                                    if (v59 >= int(v8)) {
                                                        switch v7 {
                                                        case "":
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        case "hello":
                                                            _ = len(v1)
                                                        case "foo bar":
                                                            fmt.Println(v8)
                                                        default:
                                                            v1 = append(v1, v59)
                                                        }
                                                        var v61 int = (fi56 >> (((^(v2))) & 63))
                                                        _ = v61
                                                    }
                                                    v60 += (^(v4))
                                                    fmt.Println(v8, fi56, v8)
                                                    if len(v3) > 0 {
                                                        v3[0] = v60
                                                    }
                                                    switch v7 {
                                                    case "":
                                                        var v62 int8 = v8
                                                        _ = v62
                                                    case "hello":
                                                        var v63 int8 = (v8 + (v8 % (((v8 * 100)) | 1)))
                                                        _ = v63
                                                    case "foo bar":
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    default:
                                                        var v64 int8 = 1
                                                        _ = v64
                                                    }
                                                } else {
                                                    switch v7 {
                                                    case "":
                                                        var v65 string = ("a" + v7)
                                                        _ = v65
                                                    case "hello":
                                                        var v66 float32 = 1.5
                                                        _ = v66
                                                    default:
                                                        var v67 uint64 = 377
                                                        _ = v67
                                                    }
                                                    if false {
                                                        var v68 float32 = 0.1
                                                        var v69 uint32 = (v60 / ((v60) | 1))
                                                        v60 = 42
                                                        _ = v68
                                                        _ = v69
                                                    }
                                                    switch fi56 {
                                                    case 128:
                                                        var v70 float32 = (((0.0 + float32(1e10)) - (0.1 / 1e10)) + (0.1 / 1.0))
                                                        _ = v70
                                                    case 32768:
                                                        v71 := GenS3{F0: (^(v2)), F1: (true && false)}
                                                        _ = v71
                                                    case 4294967295:
                                                        v72 := GenS4{F0: v60, F1: 2}
                                                        _ = v72
                                                    }
                                                }
                                            }
                                            var v73 uint32 = (^(v60))
                                            var v74 uint16 = 228
                                            v7 = v7
                                            if true {
                                                v8 = (v8 * int8(v60))
                                                var v75 int8 = v8
                                                if len(v1) > 0 {
                                                    v1[0] = fi53
                                                }
                                                _ = v75
                                            } else {
                                                var v76 uint64 = 632
                                                fmt.Println(v2, v74)
                                                v60 = uint32(v74)
                                                _ = v76
                                            }
                                            _ = v73
                                            _ = v74
                                        } else {
                                            if len(v1) > 0 {
                                                v1[0] = (fi53 & v59)
                                            }
                                        }
                                    case "a":
                                        v8 = -128
                                    case "foo bar":
                                        fmt.Println(v7, v7, fi53)
                                    default:
                                        v3 = append(v3, 0)
                                    }
                                    _ = v60
                                }
                                var v77 string = (("foo bar" + "hello") + v7)
                                _ = v59
                                _ = v77
                            }
                        case "a":
                            v1 = append(v1, 128)
                        default:
                            var v78 int16 = 0
                            _ = v78
                        }
                    }
                    switch "a" {
                    case "foo bar":
                        v8 += (v8 ^ (v8 - (v8 ^ 0)))
                    }
                    _ = fi56
                }
                var v79 string = ("" + v7)
                fmt.Println(v2, v8)
                _ = v79
            }
            if len(v1) > 0 {
                v1[0] = (-(fi53))
            }
            var v80 int16 = -456
            _ = v80
        }
        v4 = 889
        _ = fi53
    }
    v4 = (v4 + (v4 >> (((v2 * (v2 << (((v2 / ((uint(v4)) | 1))) & 63)))) & 63)))
    var v81 string = ("a" + "foo bar")
    var v82 float32 = float32(0.0)
    var v83 string = v7
    switch (("" + "hello") + v81) {
    case "foo bar":
        v84 := GenS9{F0: 128}
        _ = v84
    case "":
        v4 = v4
    case "\n\t":
        if false {
            v85 := GenS4{F0: (v4 * v4), F1: 18446744073709551615}
            _ = v85
        } else {
            v86 := GenS7{F0: v83, F1: (float32(1e10) == 1.5), F2: "hello"}
            var v87 uint32 = v4
            _ = v86
            _ = v87
        }
    }
    _ = len(v1)
    fmt.Println(v2, v8)
    switch 32767 {
    case 128:
        _ = len(v1)
    case 2147483647:
        var v88 int32 = -1
        _ = v88
    }
    if false {
        if false {
            var v89 uint32 = v4
            v90 := GenS2{F0: 42}
            _ = v89
            _ = v90
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        if false {
            fmt.Println(v82)
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v81
    _ = v82
    _ = v83
    _ = v9
}
