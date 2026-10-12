package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 string; F1 bool; F2 int }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint64; F1 int32 }
type GenS8 struct { F0 string; F1 int }
type GenS9 struct { F0 uint16; F1 int64; F2 uint8 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 uint32 }

func main() {
    v1 := []int{}
    var v2 uint8 = 42
    var v3 float64 = 1e10
    var v4 int16 = -1
    v5 := make(map[string]int)
    var v6 float32 = 1.5
    var v7 uint64 = 0
    var v8 int16 = v4
    if ((v7 & (v7 << ((65535) & 63))) >= (^(v7))) {
        v9 := GenS1{F0: 0}
        for fi10 := 0; fi10 < 12; fi10++ {
            _ = len(v1)
            var v11 int16 = v4
            _ = fi10
            _ = v11
        }
        switch 32768 {
        case -109:
            v3 = (float64((0.0 / 1e10)) * 1.5)
        case 2147483647:
            var v12 float32 = 0.1
            _ = v12
        case -325:
            v6 = ((3.14159 + (0.1 / float32(1.5))) - v6)
        default:
            for fi13 := 0; fi13 < 2; fi13++ {
                v4 -= (v4 - (v4 + (v4 / ((v8) | 1))))
                _ = len(v1)
                if (true || ((false || false) || (2147483647 == -1))) {
                    if true {
                        fmt.Println(v3, fi13, v3)
                        v3 = 1e10
                        v2 = v2
                    } else {
                        if (!false) {
                            switch "hello" {
                            case "a":
                                if true {
                                    if false {
                                        fmt.Println(v3)
                                        v14 := GenS1{F0: 2147483647}
                                        var v15 int8 = 127
                                        _ = v14
                                        _ = v15
                                    } else {
                                        switch "a" {
                                        case "hello":
                                            fmt.Println(v4, v4, v8)
                                        case "foo bar":
                                            var v16 int16 = v4
                                            _ = v16
                                        }
                                        v17 := GenS1{F0: 415}
                                        _ = v17
                                    }
                                }
                            }
                            switch "a" {
                            case "hello":
                                switch (((("hello" + ("hello" + ("" + ""))) + ("hello" + "")) + "hello") + "foo bar") {
                                case "foo bar":
                                    if false {
                                        if (true && true) {
                                            switch "hello" {
                                            case "a":
                                                fmt.Println(v6, v6, fi13)
                                            case "foo bar":
                                                fmt.Println(v7)
                                            case "":
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            default:
                                                for fi18 := 0; fi18 < 9; fi18++ {
                                                    var v19 int64 = -9223372036854775808
                                                    _ = fi18
                                                    _ = v19
                                                }
                                            }
                                        } else {
                                            var v20 int = (fi13 >> ((42) & 63))
                                            _ = v20
                                        }
                                    }
                                default:
                                    fmt.Println(v6)
                                }
                            case "\n\t":
                                v21 := GenS1{F0: 2147483648}
                                _ = v21
                            case "":
                                var v22 uint16 = 372
                                _ = v22
                            default:
                                v23 := GenS1{F0: 380}
                                _ = v23
                            }
                        }
                    }
                    if false {
                        _ = len(v1)
                        if len(v1) > 0 {
                            v1[0] = int(v7)
                        }
                        switch "\n\t" {
                        case "":
                            v24 := GenS1{F0: 4294967295}
                            _ = v24
                        case "hello":
                            var v25 int32 = 2147483647
                            _ = v25
                        case "foo bar":
                            if true {
                                var v26 uint = 500
                                var v27 int8 = 86
                                v28 := GenS2{F0: 65535, F1: 518}
                                _ = v26
                                _ = v27
                                _ = v28
                            }
                        default:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                    } else {
                        if (0 < 1) {
                            for fi29 := 0; fi29 < 6; fi29++ {
                                if (42 < -62) {
                                    var v30 int16 = -32768
                                    var v31 float32 = 0.1
                                    _ = v30
                                    _ = v31
                                }
                                var v32 bool = true
                                _ = fi29
                                _ = v32
                            }
                            var v33 uint32 = 1
                            _ = v33
                        } else {
                            if true {
                                v7 = (v7 / ((v7) | 1))
                            }
                            switch "\n\t" {
                            case "\n\t":
                                var v34 int16 = (^(v8))
                                _ = v34
                            default:
                                v1 = append(v1, (fi13 >> ((18446744073709551615) & 63)))
                            }
                            if false {
                                for fi35 := 0; fi35 < 10; fi35++ {
                                    v5["foo bar"] = int(v7)
                                    fmt.Println(v8)
                                    var v36 float32 = ((0.0 / (3.14159 / v6)) - ((v6 / -2.25) + v6))
                                    var v37 int = (-(fi13))
                                    fmt.Println(fi13, v3)
                                    _ = fi35
                                    _ = v36
                                    _ = v37
                                }
                            } else {
                                v7 = (v7 << ((0) & 63))
                                var v38 int64 = 860
                                _ = v38
                            }
                        }
                        for fi39 := 0; fi39 < 16; fi39++ {
                            if ((v3 == ((float64((v6 * float32(v3))) / float64((v6 / 1.5))) * v3)) || (1 < 65535)) {
                                switch (fi39 % (((fi39 | 0)) | 1)) {
                                case 127:
                                    switch fi39 {
                                    case -2147483648:
                                        v4 = -1
                                    }
                                }
                                switch "a" {
                                case "a":
                                    fmt.Println(v3, v6)
                                case "hello":
                                    var v40 string = ""
                                    _ = v40
                                case "":
                                    v41 := GenS3{F0: v2, F1: 41}
                                    _ = v41
                                }
                            }
                            _ = fi39
                        }
                    }
                }
                _ = fi13
            }
        }
        _ = v9
    } else {
        if true {
            v42 := GenS1{F0: 2147483648}
            v3 -= 1.5
            _ = v42
        } else {
            for fi43 := 0; fi43 < 2; fi43++ {
                for fi44 := 0; fi44 < 12; fi44++ {
                    var v45 float32 = v6
                    if (!true) {
                        var v46 int16 = (v8 + (^(v8)))
                        if false {
                            var v47 int16 = v8
                            v46 += (v8 >> ((18446744073709551615) & 63))
                            fmt.Println(v7)
                            _ = v47
                        }
                        if false {
                            if (v7 > v7) {
                                switch fi44 {
                                case -1:
                                    var v48 int64 = 1
                                    _ = v48
                                case 1:
                                    v49 := GenS4{F0: true}
                                    _ = v49
                                }
                                switch (("foo bar" + "\n\t") + "hello") {
                                case "a":
                                    var v50 uint32 = 0
                                    _ = v50
                                case "foo bar":
                                    switch 32768 {
                                    case 1:
                                        var v51 string = "\n\t"
                                        _ = v51
                                    default:
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                default:
                                    switch -343 {
                                    case 32768:
                                        v8 = v4
                                    default:
                                        fmt.Println(v45, v3, v2)
                                    }
                                }
                                v52 := GenS3{F0: (^(v2)), F1: 2147483647}
                                _ = v52
                            } else {
                                v1 = append(v1, fi43)
                            }
                        }
                        _ = v46
                    }
                    var v53 int8 = 100
                    _ = fi44
                    _ = v45
                    _ = v53
                }
                _ = fi43
            }
        }
        v54 := GenS1{F0: 0}
        _ = v54
    }
    v55 := GenS5{F0: "foo bar", F1: true, F2: 32767}
    v3 -= v3
    if (!true) {
        v2 = v2
    } else {
        switch "hello" {
        case "hello":
            var v56 uint16 = 65535
            _ = v56
        case "":
            for fi57 := 0; fi57 < 2; fi57++ {
                v7 = (v7 * 300)
                _ = fi57
            }
        case "\n\t":
            _ = len(v5)
        }
    }
    var v58 uint8 = 255
    for fi59 := 0; fi59 < 17; fi59++ {
        for fi60 := 0; fi60 < 1; fi60++ {
            _ = len(v1)
            v61 := GenS3{F0: v2, F1: 613}
            fmt.Println(v8, v7)
            _ = fi60
            _ = v61
        }
        fmt.Println(v8, v4)
        var v62 int16 = (v4 | v8)
        _ = fi59
        _ = v62
    }
    v6 = v6
    var v63 uint = 0
    switch ("hello" + "\n\t") {
    case "foo bar":
        if true {
            for fi64 := 0; fi64 < 12; fi64++ {
                var v65 bool = false
                var v66 int32 = 387
                if true {
                    var v67 bool = (!v65)
                    var v68 uint64 = v7
                    switch (-(fi64)) {
                    case 128:
                        if ((100 >= int8(v58)) || ((!false) || (!false))) {
                            for fi69 := 0; fi69 < 8; fi69++ {
                                v63 -= 0
                                switch 32767 {
                                case 0:
                                    v68 = uint64(v63)
                                case 278:
                                    v7 -= v7
                                    if len(v1) > 0 {
                                        v1[0] = fi64
                                    }
                                    switch "a" {
                                    case "hello":
                                        v7 = 4294967295
                                    case "foo bar":
                                        var v70 uint16 = 2
                                        _ = v70
                                    case "a":
                                        fmt.Println(v66, v8, v7)
                                    }
                                    v7 -= v7
                                case 2147483647:
                                    var v71 int16 = 1000
                                    _ = v71
                                default:
                                    v67 = v65
                                }
                                switch 4294967295 {
                                case 9223372036854775807:
                                    v8 = v8
                                case 42:
                                    fmt.Println(v2, v66, fi64)
                                case -9223372036854775808:
                                    fmt.Println(v66, v66)
                                    if (255 <= 2) {
                                        fmt.Println(fi64)
                                        var v72 int64 = 0
                                        v7 -= 2
                                        _ = v72
                                    }
                                    if (v67 && (int16(v66) > (^(v8)))) {
                                        if (9223372036854775807 != int64(fi64)) {
                                            fmt.Println(fi69, v58, v4)
                                            v73 := v5["hello"]
                                            v5["\n\t"] = (v73 >> ((v63) & 63))
                                            _ = v73
                                        } else {
                                            fmt.Println(v3, fi64)
                                            v74 := v5[""]
                                            var v75 uint8 = 127
                                            v2 += v58
                                            v6 = (float32(v3) * (((v6 / v6) + (v6 / v6)) / (v6 * 1.5)))
                                            var v76 uint32 = 42
                                            switch "foo bar" {
                                            case "hello":
                                                if (v67 || (true && (0 < v7))) {
                                                    v77 := v5["\n\t"]
                                                    var v78 uint = (v63 + (v63 / ((v63) | 1)))
                                                    if (!(0.1 > (v3 / v3))) {
                                                        v79 := GenS1{F0: v76}
                                                        _ = v79
                                                    }
                                                    _ = v77
                                                    _ = v78
                                                }
                                            case "foo bar":
                                                var v80 int32 = 0
                                                _ = v80
                                            case "":
                                                fmt.Println(v65)
                                            }
                                            _ = v74
                                            _ = v75
                                            _ = v76
                                        }
                                    }
                                    if (v65 && (!((((!v67) || v65) || (987 >= 775)) || v67))) {
                                        var v81 int8 = -1
                                        if (!(0.1 > 0.1)) {
                                            v4 -= (v8 - 1)
                                            var v82 bool = true
                                            v83 := GenS6{F0: v58}
                                            _ = v82
                                            _ = v83
                                        } else {
                                            v84 := GenS6{F0: 45}
                                            if v67 {
                                                delete(v5, "foo bar")
                                                v81 -= -128
                                                v67 = v67
                                            }
                                            _ = v84
                                        }
                                        _ = v81
                                    } else {
                                        if (378 >= 2) {
                                            v6 += v6
                                            v85 := v5[(("" + "\n\t") + "hello")]
                                            _ = v85
                                        } else {
                                            v86 := GenS7{F0: (v7 | 2), F1: (v66 | (v66 * (^(v66))))}
                                            _ = v86
                                        }
                                    }
                                }
                                _ = fi69
                            }
                            v87 := GenS8{F0: (("" + "hello") + "hello"), F1: fi64}
                            var v88 uint64 = 1
                            _ = v87
                            _ = v88
                        } else {
                            var v89 bool = true
                            v67 = false
                            v90 := GenS9{F0: 2, F1: -1, F2: uint8(v7)}
                            _ = v89
                            _ = v90
                        }
                    case 0:
                        v91 := GenS10{F0: 42}
                        _ = v91
                    }
                    _ = v67
                    _ = v68
                } else {
                    v92 := GenS11{F0: 2}
                    _ = v92
                }
                _ = fi64
                _ = v65
                _ = v66
            }
        }
    }
    for fi93 := 0; fi93 < 16; fi93++ {
        v6 = 1.5
        if true {
            for fi94 := 0; fi94 < 11; fi94++ {
                var v95 int64 = 42
                _ = fi94
                _ = v95
            }
        } else {
            if true {
                for fi96 := 0; fi96 < 6; fi96++ {
                    if true {
                        var v97 bool = true
                        _ = v97
                    }
                    v63 += (v63 << ((v63) & 63))
                    v5["foo bar"] = 127
                    var v98 int = fi93
                    v99 := GenS9{F0: 0, F1: 9223372036854775807, F2: (v58 >> (((v63 * (v63 & (v63 & v63)))) & 63))}
                    _ = fi96
                    _ = v98
                    _ = v99
                }
            }
            switch 742 {
            case 128:
                v7 -= 4294967295
            case -335:
                for fi100 := 0; fi100 < 2; fi100++ {
                    var v101 uint = (v63 ^ v63)
                    v101 = v101
                    var v102 uint64 = 1
                    var v103 uint8 = 128
                    var v104 string = ""
                    v63 -= v63
                    v105 := GenS12{F0: 4294967295}
                    _ = fi100
                    _ = v101
                    _ = v102
                    _ = v103
                    _ = v104
                    _ = v105
                }
            }
        }
        _ = fi93
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v58
    _ = v6
    _ = v63
    _ = v7
    _ = v8
}
