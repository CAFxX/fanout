package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 bool; F1 int16; F2 int }
type GenS3 struct { F0 int16; F1 int8; F2 int64 }
type GenS4 struct { F0 bool; F1 int8; F2 uint32 }
type GenS5 struct { F0 int; F1 int8 }
type GenS6 struct { F0 int16; F1 int32; F2 uint }
type GenS7 struct { F0 string; F1 bool }
type GenS8 struct { F0 int64; F1 int64 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 int8 }
func genFn0(p0 int8, p1 int, p2 int) bool {
    v7 := GenS1{F0: 42, F1: 96}
    fmt.Println(p2, p1)
    for fi8 := 0; fi8 < 2; fi8++ {
        v9 := []int{fi8, (p2 - (p2 | 9223372036854775807)), p2, p2}
        for fi10 := 0; fi10 < 8; fi10++ {
            fmt.Println(fi8, p1, fi10)
            _ = fi10
        }
        _ = fi8
        _ = v9
    }
    p2 = -682
    if false {
        v11 := GenS2{F0: true, F1: -32768, F2: (p2 + p2)}
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return true
}

func main() {
    var v1 uint8 = 1
    var v2 float64 = 0.0
    v3 := make(map[uint32]int)
    var v4 string = "\n\t"
    var v5 uint8 = v1
    var v6 uint32 = 0
    var v12 uint32 = v6
    if true {
        switch v4 {
        case "hello":
            if false {
                for fi13 := 0; fi13 < 13; fi13++ {
                    v12 = v12
                    _ = fi13
                }
                var v14 int16 = 1
                _ = v14
            } else {
                v15 := []int{}
                if false {
                    switch v4 {
                    case "hello":
                        var v16 uint8 = (v5 | v5)
                        _ = v16
                    case "":
                        _ = len(v3)
                    default:
                        fmt.Println(v12, v6, v12)
                    }
                    v2 += 0.0
                }
                _ = v15
            }
        case "":
            v17 := GenS3{F0: 42, F1: 127, F2: 1}
            _ = v17
        }
        for fi18 := 0; fi18 < 18; fi18++ {
            v3[v12] = 127
            if ((fi18 + (-(fi18))) != fi18) {
                var v19 float64 = v2
                v1 += (v5 | (v1 << ((769) & 63)))
                v20 := v3[(v6 << ((382) & 63))]
                _ = v19
                _ = v20
            } else {
                var v21 float64 = v2
                switch (v4 + "") {
                case "\n\t":
                    if ((v5 & 127) < 219) {
                        var v22 int32 = 42
                        var v23 int8 = 100
                        _ = v22
                        _ = v23
                    }
                case "hello":
                    v24 := GenS3{F0: 1, F1: 42, F2: 28}
                    _ = v24
                case "":
                    if true {
                        if ((v5 >= (v5 + (v1 / ((221) | 1)))) && false) {
                            v25 := GenS2{F0: ((42 == -2147483648) || false), F1: 0, F2: 745}
                            _ = v25
                        } else {
                            switch fi18 {
                            case -249:
                                v1 += (v1 + v5)
                            case 42:
                                switch (v4 + (v4 + "hello")) {
                                case "":
                                    if false {
                                        v26 := v3[(v6 >> ((277) & 63))]
                                        for fi27 := 0; fi27 < 3; fi27++ {
                                            var v28 bool = ((v6 >> ((573) & 63)) <= (v12 + (^(v12))))
                                            if (v2 == 1.5) {
                                                v21 = float64(float32(3.14159))
                                                v29 := GenS4{F0: (1 < 18446744073709551615), F1: 1, F2: v6}
                                                _ = v29
                                            }
                                            if (-1 != 1) {
                                                v30 := GenS4{F0: (403 != 255), F1: 79, F2: v12}
                                                if (v26 != -2147483648) {
                                                    v31 := []int{42, fi18, fi18}
                                                    _ = v31
                                                }
                                                v32 := GenS1{F0: 157, F1: v5}
                                                _ = v30
                                                _ = v32
                                            }
                                            _ = fi27
                                            _ = v28
                                        }
                                        _ = v26
                                    }
                                default:
                                    v33 := GenS3{F0: -622, F1: -127, F2: 907}
                                    _ = v33
                                }
                            default:
                                switch (fi18 % ((10) | 1)) {
                                case -9223372036854775808:
                                    switch (fi18 << ((550) & 63)) {
                                    case 1:
                                        switch fi18 {
                                        case 32767:
                                            v34 := GenS5{F0: (fi18 | fi18), F1: 0}
                                            _ = v34
                                        case 128:
                                            v3[(^(v6))] = (fi18 / ((fi18) | 1))
                                        }
                                    case -590:
                                        for fi35 := 0; fi35 < 17; fi35++ {
                                            v36 := v3[2147483648]
                                            _ = fi35
                                            _ = v36
                                        }
                                    case 946:
                                        for fi37 := 0; fi37 < 10; fi37++ {
                                            if (!((0.1 / float32((v21 / v21))) > (0.0 - 1.5))) {
                                                if (true || false) {
                                                    if false {
                                                        var v38 uint16 = 1
                                                        v39 := GenS5{F0: -9223372036854775808, F1: -32}
                                                        switch ("" + ((v4 + v4) + "foo bar")) {
                                                        case "":
                                                            fmt.Println(fi18)
                                                        }
                                                        _ = v38
                                                        _ = v39
                                                    }
                                                } else {
                                                    v21 -= -2.25
                                                    v40 := GenS3{F0: 1, F1: -1, F2: -9223372036854775808}
                                                    if (!true) {
                                                        switch 32767 {
                                                        case 32767:
                                                            v21 = v2
                                                        case 9223372036854775807:
                                                            var v41 uint8 = (v1 - v1)
                                                            _ = v41
                                                        default:
                                                            if (float32(float64(1e10)) == 1e10) {
                                                                v42 := []int{}
                                                                v43 := GenS6{F0: 1000, F1: -1, F2: 137}
                                                                _ = v42
                                                                _ = v43
                                                            } else {
                                                                v44 := GenS5{F0: 42, F1: 1}
                                                                v4 = v4
                                                                _ = v44
                                                            }
                                                            v12 += (^(v12))
                                                            var v45 uint16 = 255
                                                            _ = v45
                                                        }
                                                        fmt.Println(v21, v5)
                                                    } else {
                                                        if (42 >= -1) {
                                                            fmt.Println(fi37, fi37, v12)
                                                        }
                                                    }
                                                    _ = v40
                                                }
                                                v46 := []int{(fi37 - fi37), (fi37 + fi18), fi37, fi37}
                                                _ = v46
                                            } else {
                                                if (false || false) {
                                                    switch (fi18 / ((fi18) | 1)) {
                                                    case 32767:
                                                        v5 -= v5
                                                    default:
                                                        var v47 float64 = (1.5 / ((v2 - v21) + v2))
                                                        if ((false && (127 < 622)) && false) {
                                                            if (170 <= 65535) {
                                                                _ = len(v3)
                                                                var v48 bool = false
                                                                _ = v48
                                                            } else {
                                                                switch 174 {
                                                                case -9223372036854775808:
                                                                    v49 := []int{(-(fi37)), (fi18 & (fi18 & fi37)), (fi18 / (((fi37 * (fi37 >> ((4294967295) & 63)))) | 1)), (fi18 ^ (fi37 + 0))}
                                                                    _ = v49
                                                                case 112:
                                                                    var v50 uint32 = 4294967295
                                                                    _ = v50
                                                                default:
                                                                    v12 -= 0
                                                                }
                                                                v6 += (^(v12))
                                                            }
                                                            if false {
                                                                fmt.Println(v6)
                                                            }
                                                            if false {
                                                                var v51 uint64 = 790
                                                                _ = v51
                                                            } else {
                                                                if (true && (!(-1 == 1000))) {
                                                                    v52 := []int{fi37, fi37}
                                                                    var v53 uint8 = uint8(v6)
                                                                    _ = v52
                                                                    _ = v53
                                                                }
                                                                var v54 float64 = (float64((-2.25 / 1.0)) * float64(float32(3.14159)))
                                                                _ = v54
                                                            }
                                                        }
                                                        var v55 int = (fi18 << ((1) & 63))
                                                        _ = v47
                                                        _ = v55
                                                    }
                                                    fmt.Println(v12)
                                                } else {
                                                    v12 += (v6 & (v12 * 2147483647))
                                                    var v56 bool = (3.14159 >= (0.1 / float32(v21)))
                                                    if (v56 && true) {
                                                        var v57 uint16 = 255
                                                        v58 := []int{-2147483648, fi18, fi18}
                                                        v5 = (v5 << ((886) & 63))
                                                        _ = v57
                                                        _ = v58
                                                    }
                                                    v2 += ((-2.25 + 0.0) / (v21 + v21))
                                                    _ = v56
                                                }
                                                var v59 float32 = (0.0 + 3.14159)
                                                if false {
                                                    fmt.Println(v21, v21, v12)
                                                    v4 = (((v4 + v4) + v4) + "hello")
                                                    v1 = v5
                                                }
                                                _ = v59
                                            }
                                            if false {
                                                v2 -= (1e10 / (v21 - (1.5 / float64(1.5))))
                                                fmt.Println(fi37, v6)
                                                v1 -= 28
                                            }
                                            v3[335] = (fi37 >> ((255) & 63))
                                            v60 := GenS7{F0: "foo bar", F1: (1 != 237)}
                                            var v61 int64 = -1
                                            _ = fi37
                                            _ = v60
                                            _ = v61
                                        }
                                    default:
                                        var v62 bool = true
                                        _ = v62
                                    }
                                case -16:
                                    fmt.Println(v12)
                                default:
                                    v63 := GenS1{F0: 0, F1: v5}
                                    _ = v63
                                }
                            }
                            var v64 int = (fi18 & (fi18 ^ fi18))
                            _ = v64
                        }
                        for fi65 := 0; fi65 < 5; fi65++ {
                            v66 := GenS8{F0: 9223372036854775807, F1: 875}
                            switch v4 {
                            case "":
                                v12 = (v6 ^ 2)
                            case "foo bar":
                                var v67 int64 = 0
                                _ = v67
                            default:
                                if true {
                                    if true {
                                        v4 = v4
                                    } else {
                                        var v68 bool = true
                                        var v69 int64 = 0
                                        var v70 uint8 = v5
                                        if v68 {
                                            var v71 int64 = 9223372036854775807
                                            v72 := GenS2{F0: (float64(0.0) > v21), F1: -699, F2: fi65}
                                            v73 := GenS3{F0: 0, F1: 0, F2: 42}
                                            _ = v71
                                            _ = v72
                                            _ = v73
                                        }
                                        v21 = (v2 * 0.0)
                                        v6 = (v6 << ((2) & 63))
                                        v12 += (v6 * v6)
                                        _ = v68
                                        _ = v69
                                        _ = v70
                                    }
                                    v74 := GenS4{F0: false, F1: -114, F2: v12}
                                    _ = v74
                                }
                            }
                            v75 := GenS1{F0: 65535, F1: v5}
                            _ = fi65
                            _ = v66
                            _ = v75
                        }
                    }
                default:
                    switch v4 {
                    case "\n\t":
                        v76 := GenS9{F0: 1}
                        _ = v76
                    default:
                        if (int8(v1) != 100) {
                            if true {
                                v5 -= (v5 % (((v1 - v1)) | 1))
                                v77 := GenS2{F0: (2 > 117), F1: 32767, F2: (fi18 / ((fi18) | 1))}
                                v78 := GenS9{F0: 128}
                                _ = v77
                                _ = v78
                            }
                        }
                    }
                }
                _ = v21
            }
            fmt.Println(v4, v12)
            _ = fi18
        }
    } else {
        var v79 int64 = 42
        v5 = v1
        switch "\n\t" {
        case "a":
            var v80 float32 = (float32(0.1) / -2.25)
            _ = v80
        case "":
            switch 32768 {
            case -9223372036854775808:
                delete(v3, uint32(v5))
            default:
                var v81 uint64 = 18446744073709551615
                _ = v81
            }
        case "foo bar":
            v79 = -470
        }
        _ = v79
    }
    if (-2.25 >= -2.25) {
        for fi82 := 0; fi82 < 9; fi82++ {
            switch v4 {
            case "hello":
                if (1 >= 0) {
                    v1 -= 0
                } else {
                    for fi83 := 0; fi83 < 20; fi83++ {
                        var v84 int = fi82
                        _ = len(v3)
                        switch 199 {
                        case -9223372036854775808:
                            v85 := []int{(v84 / (((fi82 - 127)) | 1)), 4294967295}
                            _ = v85
                        }
                        _ = fi83
                        _ = v84
                    }
                    var v86 int16 = 42
                    _ = v86
                }
            case "":
                v12 -= v12
            default:
                v87 := GenS10{F0: v12}
                _ = v87
            }
            _ = fi82
        }
    }
    if true {
        v88 := GenS11{F0: 100}
        if false {
            switch v4 {
            case "a":
                for fi89 := 0; fi89 < 4; fi89++ {
                    for fi90 := 0; fi90 < 5; fi90++ {
                        v1 -= (v1 * 0)
                        _ = fi90
                    }
                    var v91 uint64 = 1
                    _ = fi89
                    _ = v91
                }
            case "\n\t":
                if false {
                    var v92 int16 = 933
                    _ = v92
                }
            }
        } else {
            switch v4 {
            case "foo bar":
                if true {
                    var v93 uint16 = 42
                    _ = v93
                }
            case "\n\t":
                v94 := []int{1, 42, 32768, 405}
                _ = v94
            }
            var v95 string = v4
            _ = v95
        }
        var v96 int8 = -128
        _ = v88
        _ = v96
    } else {
        var v97 int = 199
        for fi98 := 0; fi98 < 20; fi98++ {
            var v99 uint16 = 2
            v100 := GenS4{F0: true, F1: -74, F2: v6}
            _ = fi98
            _ = v100
            _ = v99
        }
        if true {
            var v101 int8 = 54
            v5 = v5
            v2 = 0.1
            _ = v101
        }
        _ = v97
    }
    v12 -= uint32(v5)
    fmt.Println(v1, v6)
    switch "\n\t" {
    case "a":
        v12 += (v12 * (v12 - 65535))
    default:
        var v102 int = -9223372036854775808
        _ = v102
    }
    switch ((v4 + v4) + (v4 + "")) {
    case "":
        var v103 int16 = int16(v1)
        _ = v103
    }
    for fi104 := 0; fi104 < 4; fi104++ {
        v12 = v6
        var v105 int8 = 55
        _ = fi104
        _ = v105
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
