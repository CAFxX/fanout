package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 uint64; F2 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int; F1 uint64; F2 int16 }
type GenS5 struct { F0 int; F1 uint32; F2 uint16 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint16; F1 uint64; F2 int16 }
type GenS8 struct { F0 int; F1 int }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int8; F1 int; F2 uint8 }
type GenS12 struct { F0 int32; F1 uint64 }
type GenS13 struct { F0 uint16; F1 bool }
func genFn0(p0 int32) uint16 {
    for fi6 := 0; fi6 < 5; fi6++ {
        for fi7 := 0; fi7 < 19; fi7++ {
            var v8 float32 = 1e10
            if (!true) {
                v9 := []int{}
                _ = v9
            }
            if true {
                var v10 uint64 = 18446744073709551615
                _ = v10
            } else {
                var v11 int64 = 9223372036854775807
                var v12 uint = 305
                _ = v11
                _ = v12
            }
            _ = fi7
            _ = v8
        }
        for fi13 := 0; fi13 < 4; fi13++ {
            var v14 string = "hello"
            p0 -= p0
            v14 = v14
            _ = fi13
            _ = v14
        }
        _ = fi6
    }
    v15 := []int{-9223372036854775808, 9223372036854775807, 128}
    if true {
        var v16 uint16 = 1
        for fi17 := 0; fi17 < 9; fi17++ {
            v18 := GenS1{F0: 687}
            _ = fi17
            _ = v18
        }
        _ = v16
    } else {
        p0 += -1
    }
    if true {
        fmt.Println(p0)
        var v19 int16 = 32767
        v20 := GenS1{F0: 603}
        _ = v19
        _ = v20
    } else {
        var v21 uint16 = 65535
        for fi22 := 0; fi22 < 7; fi22++ {
            switch "hello" {
            case "hello":
                v21 = (^(v21))
            default:
                v23 := GenS2{F0: (fi22 ^ (fi22 + -1)), F1: 0, F2: p0}
                _ = v23
            }
            v15 = append(v15, fi22)
            _ = fi22
        }
        p0 -= p0
        _ = v21
    }
    fmt.Println(p0)
    _ = p0
    _ = v15
    return 2
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint16 = 0
    var v3 float32 = 1.5
    var v4 int16 = 42
    var v5 int64 = 0
    var v24 uint16 = v2
    var v25 int32 = 1
    switch "\n\t" {
    case "foo bar":
        var v26 float32 = v3
        _ = v26
    case "a":
        for fi27 := 0; fi27 < 15; fi27++ {
            switch "foo bar" {
            case "foo bar":
                for fi28 := 0; fi28 < 13; fi28++ {
                    var v29 bool = (784 > 531)
                    _ = fi28
                    _ = v29
                }
            case "\n\t":
                delete(v1, 42)
            }
            if (!(0.0 != float64(((float32(float64(-2.25)) + 3.14159) + (v3 / float32(1e10)))))) {
                switch (("foo bar" + ("foo bar" + "")) + "\n\t") {
                case "hello":
                    var v30 int = -311
                    _ = v30
                case "\n\t":
                    v31 := GenS2{F0: fi27, F1: 855, F2: v25}
                    _ = v31
                }
            }
            v24 += 258
            _ = fi27
        }
    default:
        v5 -= (v5 ^ (v5 >> ((42) & 63)))
    }
    var v32 int32 = 1
    v32 += v32
    for fi33 := 0; fi33 < 18; fi33++ {
        if true {
            fmt.Println(v4)
        }
        if (!(2 == 2)) {
            for fi34 := 0; fi34 < 15; fi34++ {
                var v35 float32 = 1e10
                v36 := GenS3{F0: int8(fi33)}
                if true {
                    v37 := GenS1{F0: 0}
                    v3 = v35
                    var v38 uint64 = 0
                    _ = v37
                    _ = v38
                } else {
                    v39 := GenS1{F0: 0}
                    if false {
                        v40 := GenS4{F0: (fi33 / ((9223372036854775807) | 1)), F1: 161, F2: -1}
                        v41 := []int{(fi33 % (((-(fi33))) | 1)), -2147483648, -702}
                        if (0 > 84) {
                            var v42 int64 = (^(v5))
                            var v43 int = fi34
                            v2 = genFn0(-1)
                            _ = v42
                            _ = v43
                        } else {
                            v44 := GenS5{F0: (fi33 ^ -141), F1: 761, F2: (v2 >> ((571) & 63))}
                            if false {
                                var v45 bool = false
                                delete(v1, 4294967295)
                                _ = v45
                            } else {
                                switch "a" {
                                case "foo bar":
                                    var v46 uint = 65535
                                    _ = v46
                                default:
                                    v47 := GenS5{F0: fi34, F1: 156, F2: v2}
                                    _ = v47
                                }
                                switch fi34 {
                                case 758:
                                    var v48 int32 = (v25 % ((-616) | 1))
                                    _ = v48
                                default:
                                    if ("a" < ("a" + "a")) {
                                        switch ("hello" + ("hello" + ("a" + "a"))) {
                                        case "foo bar":
                                            v49 := GenS6{F0: 4294967295}
                                            _ = v49
                                        case "a":
                                            switch (("hello" + "") + "") {
                                            case "a":
                                                var v50 int64 = v5
                                                _ = v50
                                            case "hello":
                                                v32 += (v25 % (((v32 | -2147483648)) | 1))
                                            case "\n\t":
                                                if (128 <= 1) {
                                                    v51 := GenS7{F0: v2, F1: 321, F2: (v4 | (v4 + v4))}
                                                    _ = v51
                                                }
                                            }
                                        case "":
                                            var v52 float64 = (0.0 / 3.14159)
                                            _ = v52
                                        }
                                    }
                                    if (fi34 != (-(fi34))) {
                                        v32 = (v25 >> ((0) & 63))
                                        v53 := GenS1{F0: 0}
                                        _ = v53
                                    }
                                    v3 -= 3.14159
                                    switch (fi34 * fi33) {
                                    case 370:
                                        fmt.Println(fi33, v5)
                                    case 128:
                                        if len(v41) > 0 {
                                            v41[0] = (-(fi33))
                                        }
                                    }
                                    var v54 int64 = 134
                                    switch (-(fi34)) {
                                    case 711:
                                        fmt.Println(v35, v4)
                                    case 32767:
                                        if ((v54 / (((v5 % ((v5) | 1))) | 1)) > v5) {
                                            var v55 string = "hello"
                                            _ = v55
                                        } else {
                                            var v56 float64 = (1e10 - ((1e10 + (0.1 + 1e10)) + (0.0 - 0.0)))
                                            _ = v56
                                        }
                                    default:
                                        if (!false) {
                                            v57 := GenS8{F0: -990, F1: fi34}
                                            _ = v57
                                        } else {
                                            var v58 uint32 = uint32(v24)
                                            switch fi34 {
                                            case 127:
                                                switch ("foo bar" + "a") {
                                                case "hello":
                                                    var v59 int8 = -1
                                                    _ = v59
                                                case "":
                                                    var v60 uint16 = (v2 + (^(v24)))
                                                    _ = v60
                                                case "a":
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        fmt.Println(v32)
                                                    case "":
                                                        if ((v54 / ((v5) | 1)) != (v5 & (v5 % (((v54 << ((65535) & 63))) | 1)))) {
                                                            var v61 float64 = 0.1
                                                            _ = v61
                                                        } else {
                                                            var v62 int = fi33
                                                            var v63 int64 = v54
                                                            fmt.Println(v54, v2, v35)
                                                            fmt.Println(v63)
                                                            _ = v62
                                                            _ = v63
                                                        }
                                                    case "a":
                                                        v25 -= 2147483647
                                                    default:
                                                        v2 = 65535
                                                    }
                                                default:
                                                    v3 = (float32(1.5) + ((v35 - (v3 - (v3 + v35))) / 1e10))
                                                }
                                            case 9223372036854775807:
                                                fmt.Println(v25, v25, v3)
                                            case -372:
                                                if true {
                                                    var v64 int = (fi34 << ((488) & 63))
                                                    if (true && true) {
                                                        fmt.Println(v25, v54, v25)
                                                        var v65 int64 = -993
                                                        v66 := v1[586]
                                                        var v67 string = v66
                                                        v68 := GenS1{F0: 441}
                                                        _ = v65
                                                        _ = v66
                                                        _ = v67
                                                        _ = v68
                                                    }
                                                    _ = v64
                                                } else {
                                                    switch fi33 {
                                                    case 621:
                                                        var v69 int64 = v5
                                                        _ = v69
                                                    case 32767:
                                                        fmt.Println(v5, v25, v24)
                                                    default:
                                                        var v70 uint16 = v24
                                                        _ = v70
                                                    }
                                                    v1[695] = "a"
                                                    v25 -= v25
                                                }
                                            default:
                                                _ = len(v41)
                                                fmt.Println(v24)
                                                if false {
                                                    fmt.Println(v25, v54)
                                                }
                                            }
                                            var v71 int = fi34
                                            _ = v58
                                            _ = v71
                                        }
                                    }
                                    _ = v54
                                }
                            }
                            v25 = 824
                            _ = v44
                        }
                        _ = v40
                        _ = v41
                    } else {
                        if (18446744073709551615 < 768) {
                            fmt.Println(v3, v2, v4)
                            v3 = v3
                            if false {
                                v72 := []int{}
                                _ = v72
                            } else {
                                v35 = (v3 / ((float32(3.14159) / float32((3.14159 / ((0.0 + 1e10) / ((((1.5 / 1.5) / 1.0) / (float64(v35) / -2.25)) * ((float64((v3 - v35)) - 1e10) * (0.0 / (-2.25 + -2.25)))))))) + float32(3.14159)))
                            }
                            switch "\n\t" {
                            case "foo bar":
                                if (false && (!((148 <= 18446744073709551615) || ((v2 << ((2) & 63)) == genFn0(-558))))) {
                                    if (true && (2 != 48)) {
                                        v73 := v1[4294967295]
                                        var v74 float64 = 1.5
                                        if (v4 >= (v4 - (-(v4)))) {
                                            v75 := GenS6{F0: 752}
                                            var v76 float32 = (float32(((v74 + (float64(v35) - float64(v35))) - ((float64((1e10 * v35)) * v74) - v74))) / v35)
                                            v77 := []int{}
                                            _ = v75
                                            _ = v76
                                            _ = v77
                                        }
                                        v78 := GenS9{F0: v24}
                                        v1[119] = (v73 + v73)
                                        _ = v73
                                        _ = v74
                                        _ = v78
                                    } else {
                                        var v79 int32 = 1
                                        switch (("\n\t" + "\n\t") + "") {
                                        case "hello":
                                            var v80 int16 = -1
                                            _ = v80
                                        default:
                                            switch "\n\t" {
                                            case "foo bar":
                                                v81 := GenS10{F0: (true || true)}
                                                _ = v81
                                            }
                                        }
                                        _ = v79
                                    }
                                    switch (fi34 + (^(fi33))) {
                                    case 42:
                                        v32 += v32
                                    case -1:
                                        v24 += (v24 + (v24 * 32767))
                                    case 32767:
                                        if (!((false && false) && (1 >= -26))) {
                                            _ = len(v1)
                                            if (v24 < 32767) {
                                                var v82 int32 = 42
                                                fmt.Println(v5)
                                                var v83 float32 = 1.5
                                                _ = v82
                                                _ = v83
                                            } else {
                                                var v84 uint8 = uint8(v2)
                                                v84 = v84
                                                v85 := GenS7{F0: 702, F1: 0, F2: v4}
                                                _ = v84
                                                _ = v85
                                            }
                                        }
                                    }
                                    var v86 uint = 42
                                    var v87 string = ((("hello" + "a") + ("foo bar" + "foo bar")) + "hello")
                                    fmt.Println(v86, v3)
                                    _ = v86
                                    _ = v87
                                } else {
                                    v88 := GenS4{F0: -1, F1: 311, F2: 1}
                                    v89 := GenS4{F0: int(v2), F1: 2, F2: 32767}
                                    _ = v88
                                    _ = v89
                                }
                            case "\n\t":
                                var v90 float64 = 1.5
                                _ = v90
                            case "a":
                                v3 -= (v3 * 0.0)
                            default:
                                v91 := []int{fi33, (fi33 | fi34), (fi33 >> ((0) & 63))}
                                _ = v91
                            }
                        } else {
                            switch "a" {
                            case "hello":
                                v1[2] = "a"
                            default:
                                var v92 uint16 = (v2 % ((v2) | 1))
                                _ = v92
                            }
                        }
                        v24 = 1
                        v25 = (v32 - 172)
                    }
                    v32 -= v25
                    v93 := GenS8{F0: fi34, F1: fi34}
                    _ = v39
                    _ = v93
                }
                _ = fi34
                _ = v35
                _ = v36
            }
        }
        _ = fi33
    }
    if (true || true) {
        fmt.Println(v3)
    }
    var v94 uint32 = 2
    v95 := GenS7{F0: (v24 % ((2) | 1)), F1: 18446744073709551615, F2: 32767}
    v3 -= -2.25
    if false {
        fmt.Println(v3, v32, v5)
    } else {
        fmt.Println(v3, v4)
        for fi96 := 0; fi96 < 5; fi96++ {
            v97 := []int{fi96, (^(fi96))}
            v3 = 1e10
            fmt.Println(fi96, v2)
            _ = fi96
            _ = v97
        }
    }
    for fi98 := 0; fi98 < 5; fi98++ {
        v5 = (v5 >> ((194) & 63))
        for fi99 := 0; fi99 < 15; fi99++ {
            var v100 float32 = float32(-2.25)
            _ = fi99
            _ = v100
        }
        v4 = (v4 + (v4 - v4))
        _ = fi98
    }
    for fi101 := 0; fi101 < 17; fi101++ {
        v102 := []int{}
        if false {
            switch 350 {
            case 42:
                v2 = (v2 + v24)
            case 2147483647:
                v103 := GenS11{F0: -1, F1: fi101, F2: 127}
                _ = v103
            case 1:
                v24 -= 42
            default:
                if true {
                    fmt.Println(v32)
                    var v104 int32 = 42
                    _ = v104
                } else {
                    var v105 int32 = 42
                    if ((v32 >= v105) || false) {
                        for fi106 := 0; fi106 < 4; fi106++ {
                            if false {
                                v107 := GenS12{F0: -1, F1: 18446744073709551615}
                                fmt.Println(v4, v2)
                                _ = v107
                            }
                            _ = fi106
                        }
                        v3 += 3.14159
                        switch ("\n\t" + "") {
                        case "\n\t":
                            fmt.Println(v25)
                        }
                    } else {
                        switch fi101 {
                        case 1:
                            var v108 uint16 = v2
                            _ = v108
                        default:
                            v109 := GenS13{F0: (v24 << ((1) & 63)), F1: true}
                            _ = v109
                        }
                    }
                    _ = v105
                }
            }
        }
        _ = fi101
        _ = v102
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v94
    _ = v95
}
