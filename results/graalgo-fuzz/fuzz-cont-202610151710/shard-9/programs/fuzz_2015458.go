package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 int; F1 uint32; F2 string }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint8; F1 int8 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 bool; F1 uint8 }
type GenS9 struct { F0 uint16; F1 uint16; F2 bool }
type GenS10 struct { F0 uint32; F1 uint64; F2 int }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 uint; F1 int }
func genFn0() string {
    var v7 float64 = (0.0 / ((0.1 + 3.14159) / 0.1))
    v8 := []int{963, -748, 4294967295}
    v9 := GenS1{F0: 832}
    _ = v7
    _ = v8
    _ = v9
    return "\n\t"
}

func main() {
    var v1 int64 = -243
    var v2 uint16 = 65535
    var v3 uint16 = v2
    var v4 int = 2147483647
    var v5 int32 = 793
    v6 := make(map[string]bool)
    var v10 float64 = 3.14159
    var v11 uint8 = 128
    fmt.Println(v3, v4)
    switch "foo bar" {
    case "foo bar":
        switch v4 {
        case 131:
            if (!false) {
                for fi12 := 0; fi12 < 11; fi12++ {
                    var v13 uint32 = 42
                    v14 := GenS2{F0: (v11 << ((16) & 63))}
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                v10 += v10
            } else {
                _ = len(v6)
                if (v3 > v3) {
                    var v15 int16 = -32768
                    v16 := GenS2{F0: (v11 + 127)}
                    v17 := GenS3{F0: 2147483647, F1: 1}
                    _ = v15
                    _ = v16
                    _ = v17
                }
                v18 := []int{v4, v4}
                _ = v18
            }
        case 9223372036854775807:
            if false {
                var v19 uint = 18446744073709551615
                var v20 uint16 = v3
                for fi21 := 0; fi21 < 8; fi21++ {
                    var v22 uint = (v19 / (((v19 * v19)) | 1))
                    if (1 >= 1) {
                        if false {
                            v23 := []int{fi21}
                            if false {
                                for fi24 := 0; fi24 < 17; fi24++ {
                                    v4 -= 1
                                    fmt.Println(v5, v1)
                                    v25 := GenS1{F0: (v20 >> ((v22) & 63))}
                                    _ = fi24
                                    _ = v25
                                }
                            }
                            v26 := GenS4{F0: (fi21 - fi21), F1: 1, F2: ("a" + (("a" + "foo bar") + "foo bar"))}
                            _ = v23
                            _ = v26
                        }
                        if true {
                            if (!(v5 != v5)) {
                                v4 = 32768
                                for fi27 := 0; fi27 < 10; fi27++ {
                                    _ = len(v6)
                                    if true {
                                        switch "foo bar" {
                                        case "a":
                                            v4 = fi21
                                        case "foo bar":
                                            v22 += 0
                                        case "\n\t":
                                            switch (^(fi27)) {
                                            case -2147483648:
                                                v28 := GenS5{F0: (fi27 >> ((v22) & 63))}
                                                _ = v28
                                            default:
                                                var v29 uint8 = v11
                                                _ = v29
                                            }
                                        }
                                        v1 -= -226
                                    }
                                    _ = fi27
                                }
                                var v30 string = (("a" + "\n\t") + "foo bar")
                                _ = v30
                            } else {
                                if (0.0 <= 3.14159) {
                                    v31 := GenS3{F0: 2, F1: 66}
                                    var v32 int16 = 0
                                    switch v4 {
                                    case 128:
                                        for fi33 := 0; fi33 < 17; fi33++ {
                                            v34 := GenS2{F0: (v11 << (((v22 | 255)) & 63))}
                                            _ = fi33
                                            _ = v34
                                        }
                                    case 0:
                                        var v35 int64 = (v1 | (v1 | -259))
                                        _ = v35
                                    case 2147483647:
                                        v36 := GenS6{F0: 0, F1: -52}
                                        _ = v36
                                    default:
                                        if ((!(464 <= (v19 + (v22 | v22)))) && ((v11 << ((v22) & 63)) == 0)) {
                                            var v37 int64 = (^(v1))
                                            v38 := []int{fi21}
                                            _ = v37
                                            _ = v38
                                        }
                                    }
                                    _ = v31
                                    _ = v32
                                }
                                v1 = (-(v1))
                                fmt.Println(v11, v10, v5)
                            }
                            for fi39 := 0; fi39 < 15; fi39++ {
                                v40 := []int{fi21, fi39}
                                var v41 int8 = 4
                                switch fi39 {
                                case 0:
                                    var v42 uint16 = v2
                                    v1 -= v1
                                    if (2 == 895) {
                                        var v43 int32 = v5
                                        if false {
                                            if false {
                                                if true {
                                                    v44 := v6["foo bar"]
                                                    _ = v44
                                                } else {
                                                    if (true || false) {
                                                        switch (fi21 ^ (fi21 * (fi21 + 32767))) {
                                                        case -243:
                                                            switch "" {
                                                            case "hello":
                                                                if len(v40) >= 2 {
                                                                    v40 = v40[1:len(v40)]
                                                                }
                                                            case "":
                                                                var v45 uint32 = 2
                                                                _ = v45
                                                            }
                                                        case 9223372036854775807:
                                                            v11 += (v11 >> (((v22 ^ 65535)) & 63))
                                                        default:
                                                            var v46 string = "foo bar"
                                                            _ = v46
                                                        }
                                                        if true {
                                                            v10 -= float64(3.14159)
                                                        } else {
                                                            var v47 float32 = 3.14159
                                                            var v48 int16 = -922
                                                            var v49 int32 = v43
                                                            v50 := GenS4{F0: fi39, F1: 502, F2: ("a" + "foo bar")}
                                                            fmt.Println(v3)
                                                            _ = v47
                                                            _ = v48
                                                            _ = v49
                                                            _ = v50
                                                        }
                                                        fmt.Println(v42, v5)
                                                    }
                                                }
                                            } else {
                                                _ = len(v6)
                                                if (uint8(v22) >= (v11 & v11)) {
                                                    v51 := GenS4{F0: 127, F1: 150, F2: ""}
                                                    _ = v51
                                                }
                                            }
                                        } else {
                                            v52 := GenS7{F0: 304}
                                            fmt.Println(fi39)
                                            if ((^(v11)) == v11) {
                                                _ = len(v40)
                                            } else {
                                                if true {
                                                    fmt.Println(fi39)
                                                }
                                            }
                                            _ = v52
                                        }
                                        fmt.Println(v11)
                                        switch fi39 {
                                        case -2147483648:
                                            var v53 float64 = v10
                                            _ = v53
                                        case 9223372036854775807:
                                            v41 += (v41 & 72)
                                        case 1:
                                            v54 := GenS2{F0: v11}
                                            _ = v54
                                        }
                                        _ = v43
                                    }
                                    _ = v42
                                default:
                                    var v55 uint = (v19 ^ v22)
                                    _ = v55
                                }
                                _ = fi39
                                _ = v40
                                _ = v41
                            }
                            var v56 bool = false
                            _ = v56
                        }
                    }
                    _ = fi21
                    _ = v22
                }
                _ = v19
                _ = v20
            } else {
                v57 := GenS5{F0: v4}
                v58 := GenS5{F0: v4}
                _ = v57
                _ = v58
            }
        case 4294967295:
            v59 := GenS8{F0: (((!false) && ((v4 >= v4) || false)) && true), F1: 128}
            _ = v59
        default:
            if (true || false) {
                v60 := []int{}
                _ = len(v6)
                _ = v60
            } else {
                var v61 bool = true
                for fi62 := 0; fi62 < 7; fi62++ {
                    v63 := []int{(v4 / ((42) | 1)), (v4 % ((v4) | 1))}
                    if ((111 > 62) && (0 == 0)) {
                        v4 = -869
                        if true {
                            v10 -= 1e10
                            var v64 uint8 = 114
                            var v65 int16 = 42
                            _ = v64
                            _ = v65
                        } else {
                            switch "foo bar" {
                            case "hello":
                                var v66 float32 = ((0.0 * 1e10) + 3.14159)
                                _ = v66
                            case "a":
                                var v67 string = ("\n\t" + "\n\t")
                                _ = v67
                            case "\n\t":
                                v10 = 1.5
                            default:
                                v4 += fi62
                            }
                            v3 = 288
                            v68 := GenS9{F0: 65535, F1: v3, F2: (42 > 0)}
                            _ = v68
                        }
                        if ((v5 >> ((42) & 63)) > -575) {
                            if (v61 || (v2 == v2)) {
                                v1 = v1
                                v63 = append(v63, v4)
                            }
                        }
                    } else {
                        for fi69 := 0; fi69 < 17; fi69++ {
                            if len(v63) >= 2 {
                                v63 = v63[1:len(v63)]
                            }
                            _ = fi69
                        }
                        for fi70 := 0; fi70 < 3; fi70++ {
                            var v71 int64 = (v1 ^ v1)
                            var v72 bool = v61
                            _ = len(v63)
                            _ = fi70
                            _ = v71
                            _ = v72
                        }
                        v73 := GenS10{F0: 1, F1: 18446744073709551615, F2: (v4 << ((2) & 63))}
                        _ = v73
                    }
                    v61 = (!(false && v61))
                    _ = fi62
                    _ = v63
                }
                var v74 int8 = 1
                _ = v61
                _ = v74
            }
        }
    case "a":
        v3 -= 99
    }
    fmt.Println(v2, v3)
    var v75 int32 = v5
    v3 += v3
    switch v4 {
    case -1:
        if (v2 < v2) {
            fmt.Println(v10)
            v5 += 2147483647
            if true {
                if (!false) {
                    var v76 float32 = (3.14159 - 0.0)
                    _ = v76
                }
                v11 = uint8(v2)
            } else {
                if (480 <= 602) {
                    if true {
                        v4 += (v4 - v4)
                        fmt.Println(v2, v5)
                    }
                } else {
                    for fi77 := 0; fi77 < 5; fi77++ {
                        var v78 int64 = v1
                        v2 += v3
                        _ = fi77
                        _ = v78
                    }
                    for fi79 := 0; fi79 < 20; fi79++ {
                        if false {
                            for fi80 := 0; fi80 < 7; fi80++ {
                                v5 += (v5 - v5)
                                switch ("" + ("" + "hello")) {
                                case "\n\t":
                                    if (!true) {
                                        fmt.Println(v75, v5, v5)
                                        v81 := GenS11{F0: 845}
                                        var v82 uint = 42
                                        _ = v81
                                        _ = v82
                                    }
                                    v5 -= v5
                                    fmt.Println(v5, v1)
                                case "":
                                    v83 := GenS12{F0: 42}
                                    _ = v83
                                default:
                                    v84 := GenS11{F0: 534}
                                    _ = v84
                                }
                                _ = fi80
                            }
                            for fi85 := 0; fi85 < 16; fi85++ {
                                v2 = v2
                                v86 := []int{(fi85 | (fi85 | -2147483648)), fi85}
                                _ = fi85
                                _ = v86
                            }
                        }
                        _ = fi79
                    }
                    for fi87 := 0; fi87 < 9; fi87++ {
                        v2 = (v3 & v2)
                        _ = len(v6)
                        fmt.Println(v10, fi87, v5)
                        _ = fi87
                    }
                }
                fmt.Println(v11, v5)
            }
        }
    case -623:
        fmt.Println(v5, v2)
    default:
        var v88 int16 = 32767
        _ = v88
    }
    v1 += v1
    for fi89 := 0; fi89 < 18; fi89++ {
        var v90 int64 = (-(v1))
        for fi91 := 0; fi91 < 16; fi91++ {
            if false {
                switch "foo bar" {
                case "\n\t":
                    v90 = (v90 & v90)
                default:
                    v92 := GenS2{F0: v11}
                    _ = v92
                }
            }
            var v93 int = 128
            var v94 int32 = v5
            _ = fi91
            _ = v93
            _ = v94
        }
        _ = fi89
        _ = v90
    }
    v95 := []int{(v4 | 128), (^(v4)), -9223372036854775808, v4}
    v96 := GenS7{F0: (v75 * v5)}
    switch "a" {
    case "":
        for fi97 := 0; fi97 < 3; fi97++ {
            v98 := GenS8{F0: true, F1: (^(v11))}
            switch (^(v4)) {
            case -2147483648:
                v99 := GenS2{F0: (v11 - v11)}
                _ = v99
            }
            _ = fi97
            _ = v98
        }
    case "hello":
        var v100 bool = (!(((true || true) && true) && false))
        _ = v100
    }
    switch (v4 - (v4 | v4)) {
    case 0:
        fmt.Println(v75, v3, v3)
    case 32768:
        switch "a" {
        case "a":
            v3 = 1
        case "\n\t":
            var v101 float32 = (float32((float64((0.1 / 3.14159)) * (v10 - v10))) * 3.14159)
            _ = v101
        }
    default:
        var v102 uint32 = 4294967295
        _ = v102
    }
    v75 = v5
    switch v4 {
    case -2147483648:
        for fi103 := 0; fi103 < 12; fi103++ {
            var v104 bool = true
            _ = fi103
            _ = v104
        }
    case -1:
        if (v10 > v10) {
            fmt.Println(v2, v3)
            for fi105 := 0; fi105 < 7; fi105++ {
                v3 = 42
                _ = fi105
            }
        }
    default:
        if true {
            v106 := GenS2{F0: v11}
            var v107 int16 = -1
            _ = v106
            _ = v107
        } else {
            if false {
                v108 := GenS13{F0: 4294967295, F1: (v4 >> ((255) & 63))}
                for fi109 := 0; fi109 < 7; fi109++ {
                    switch ((("foo bar" + (("" + "a") + "\n\t")) + (("" + "a") + ("" + "foo bar"))) + (("foo bar" + "") + "hello")) {
                    case "":
                        v1 += (v1 << ((18446744073709551615) & 63))
                    case "\n\t":
                        switch "hello" {
                        case "\n\t":
                            fmt.Println(v1)
                        default:
                            v110 := GenS7{F0: (v5 ^ v5)}
                            _ = v110
                        }
                    }
                    var v111 uint64 = 2
                    fmt.Println(v10)
                    _ = fi109
                    _ = v111
                }
                _ = v108
            }
            for fi112 := 0; fi112 < 4; fi112++ {
                fmt.Println(v5, v11)
                _ = fi112
            }
        }
    }
    v6["a"] = (uint64(v11) > 212)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v75:", v75)
    fmt.Println("slicelen v95:", len(v95), v95)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v75
    _ = v95
    _ = v96
}
