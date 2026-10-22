package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 int32; F2 uint64 }
type GenS3 struct { F0 uint8; F1 int64 }
type GenS4 struct { F0 uint32; F1 uint32; F2 bool }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int16; F1 uint32; F2 int16 }
type GenS7 struct { F0 uint; F1 uint32; F2 int }
type GenS8 struct { F0 uint64 }

func main() {
    var v1 bool = (!(42 <= 65535))
    var v2 uint32 = 4294967295
    v3 := make(map[int]string)
    var v4 int64 = 729
    v5 := make(map[uint32]string)
    if true {
        var v6 int64 = 725
        _ = v6
    } else {
        v4 += int64(v2)
        if (!v1) {
            switch "a" {
            case "\n\t":
                var v7 uint16 = 0
                _ = v7
            default:
                v8 := GenS1{F0: -2147483648}
                _ = v8
            }
            v9 := []int{-494, -9223372036854775808, -2147483648, 0}
            _ = v9
        } else {
            var v10 uint16 = 394
            v2 -= v2
            if (v1 && ((v2 << ((1) & 63)) >= 42)) {
                switch "" {
                case "":
                    for fi11 := 0; fi11 < 2; fi11++ {
                        switch -2147483648 {
                        case 1:
                            if ((v2 / ((v2) | 1)) < (v2 / ((497) | 1))) {
                                var v12 uint32 = 4294967295
                                var v13 uint8 = 42
                                _ = v12
                                _ = v13
                            }
                        case 668:
                            v2 = 733
                        }
                        for fi14 := 0; fi14 < 18; fi14++ {
                            if false {
                                fmt.Println(fi14)
                                switch fi14 {
                                case -2147483648:
                                    v1 = v1
                                case 9223372036854775807:
                                    switch "hello" {
                                    case "\n\t":
                                        v10 -= v10
                                    case "foo bar":
                                        var v15 int8 = 0
                                        _ = v15
                                    case "":
                                        var v16 int = (fi14 / ((32768) | 1))
                                        _ = v16
                                    default:
                                        if ((!v1) && v1) {
                                            v4 = v4
                                            fmt.Println(v10)
                                            var v17 uint32 = v2
                                            _ = v17
                                        } else {
                                            switch (^(fi11)) {
                                            case -1:
                                                var v18 uint8 = 255
                                                _ = v18
                                            }
                                            v19 := GenS1{F0: 2147483647}
                                            v20 := []int{(fi14 | (fi11 << ((2) & 63))), fi11, -2147483648}
                                            var v21 float32 = 1e10
                                            switch (((((("" + "") + "hello") + "a") + "a") + ("\n\t" + ("\n\t" + "hello"))) + (("\n\t" + "foo bar") + "hello")) {
                                            case "\n\t":
                                                var v22 uint = 0
                                                _ = v22
                                            default:
                                                var v23 int32 = 2147483647
                                                _ = v23
                                            }
                                            _ = v19
                                            _ = v20
                                            _ = v21
                                        }
                                        if false {
                                            var v24 string = "foo bar"
                                            _ = v24
                                        }
                                    }
                                case 127:
                                    if (!(int16(v4) != 536)) {
                                        var v25 int8 = 42
                                        switch "\n\t" {
                                        case "foo bar":
                                            var v26 uint64 = 0
                                            _ = v26
                                        case "\n\t":
                                            var v27 int32 = -1
                                            _ = v27
                                        }
                                        _ = v25
                                    } else {
                                        v4 -= v4
                                        if (!((int16(v2) != 336) && v1)) {
                                            var v28 int64 = (v4 ^ (v4 ^ 355))
                                            var v29 int32 = -330
                                            fmt.Println(v1, fi11, fi11)
                                            _ = v28
                                            _ = v29
                                        } else {
                                            var v30 int = fi11
                                            _ = v30
                                        }
                                        var v31 uint16 = v10
                                        v32 := []int{(fi14 & fi14), (^(fi11))}
                                        v33 := GenS1{F0: -2147483648}
                                        _ = v31
                                        _ = v32
                                        _ = v33
                                    }
                                }
                                switch 663 {
                                case -217:
                                    var v34 uint64 = 1
                                    _ = v34
                                default:
                                    switch "hello" {
                                    case "a":
                                        var v35 uint32 = 2
                                        _ = v35
                                    case "hello":
                                        delete(v3, fi11)
                                    case "\n\t":
                                        v4 = v4
                                    }
                                }
                                var v36 int64 = v4
                                _ = v36
                            }
                            fmt.Println(fi14, v1)
                            v1 = (!v1)
                            _ = fi14
                        }
                        fmt.Println(v1)
                        _ = fi11
                    }
                case "foo bar":
                    for fi37 := 0; fi37 < 17; fi37++ {
                        v2 = 2147483647
                        switch fi37 {
                        case 2147483647:
                            v38 := []int{fi37, (fi37 >> ((42) & 63)), (fi37 / (((-(fi37))) | 1))}
                            _ = v38
                        case 0:
                            _ = len(v3)
                        case 127:
                            fmt.Println(v1)
                        }
                        _ = fi37
                    }
                case "hello":
                    for fi39 := 0; fi39 < 8; fi39++ {
                        if v1 {
                            v40 := []int{}
                            _ = v40
                        }
                        _ = fi39
                    }
                default:
                    v41 := []int{32768, 4294967295}
                    _ = v41
                }
            }
            _ = v10
        }
        v1 = true
    }
    v5[(v2 + (v2 ^ (v2 >> ((255) & 63))))] = ("hello" + "a")
    var v42 uint = 42
    for fi43 := 0; fi43 < 15; fi43++ {
        v2 -= v2
        _ = fi43
    }
    switch 0 {
    case 32767:
        fmt.Println(v42, v1)
    case 2147483647:
        var v44 int32 = -113
        _ = v44
    default:
        fmt.Println(v2)
    }
    v45 := []int{-591, 4294967295}
    v42 = v42
    var v46 int64 = -312
    var v47 int64 = v4
    v2 += v2
    var v48 uint32 = uint32(v42)
    switch 128 {
    case 127:
        fmt.Println(v42)
    case -303:
        if v1 {
            switch "a" {
            case "hello":
                if (v48 >= 655) {
                    if ((!(!v1)) || (0.1 != 1.5)) {
                        v4 += v4
                        switch 2147483647 {
                        case 128:
                            var v49 uint = (v42 & (v42 << ((0) & 63)))
                            _ = v49
                        }
                    }
                    if v1 {
                        for fi50 := 0; fi50 < 9; fi50++ {
                            if (v1 && false) {
                                var v51 int = int(v47)
                                _ = v51
                            } else {
                                var v52 float32 = float32(0.1)
                                if v1 {
                                    switch "a" {
                                    case "":
                                        var v53 string = ""
                                        _ = v53
                                    case "hello":
                                        var v54 uint16 = 32767
                                        _ = v54
                                    case "a":
                                        if len(v45) >= 2 {
                                            v45 = v45[1:len(v45)]
                                        }
                                    }
                                    switch "a" {
                                    case "hello":
                                        fmt.Println(fi50, v4, v4)
                                    case "a":
                                        if (!(42 <= 1)) {
                                            var v55 float32 = float32((((1e10 / 1.0) - (float64((1e10 - 0.1)) / (0.0 - -2.25))) - 0.1))
                                            fmt.Println(fi50, fi50)
                                            var v56 uint32 = (v2 >> (((^(v42))) & 63))
                                            _ = v55
                                            _ = v56
                                        }
                                    case "":
                                        if (fi50 != fi50) {
                                            var v57 uint16 = 957
                                            _ = v57
                                        }
                                    default:
                                        delete(v5, v48)
                                    }
                                } else {
                                    var v58 int8 = 21
                                    _ = len(v3)
                                    if v1 {
                                        switch -527 {
                                        case -9223372036854775808:
                                            var v59 uint16 = 221
                                            _ = v59
                                        default:
                                            v4 -= 0
                                        }
                                    } else {
                                        var v60 bool = v1
                                        fmt.Println(v58, v48, v42)
                                        _ = v60
                                    }
                                    _ = v58
                                }
                                v61 := GenS1{F0: -2147483648}
                                _ = v52
                                _ = v61
                            }
                            if true {
                                if v1 {
                                    if true {
                                        switch ((("" + "foo bar") + "foo bar") + "foo bar") {
                                        case "":
                                            if (false || v1) {
                                                v62 := GenS2{F0: (v47 >> ((18446744073709551615) & 63)), F1: -471, F2: uint64(v42)}
                                                if v1 {
                                                    _ = len(v5)
                                                }
                                                var v63 float64 = -2.25
                                                _ = v62
                                                _ = v63
                                            } else {
                                                var v64 int = fi50
                                                if ((244 < 217) && v1) {
                                                    v1 = (((!v1) || (127 >= 143)) || (65535 == 2147483647))
                                                    if v1 {
                                                        for fi65 := 0; fi65 < 14; fi65++ {
                                                            v66 := GenS3{F0: 117, F1: 18}
                                                            if (((-(v4)) > v47) || v1) {
                                                                if (v1 && v1) {
                                                                    var v67 int16 = 42
                                                                    v48 += (^(v2))
                                                                    var v68 float32 = 1e10
                                                                    var v69 bool = (!v1)
                                                                    if v1 {
                                                                        var v70 int16 = v67
                                                                        v70 += v70
                                                                        _ = v70
                                                                    } else {
                                                                        v71 := GenS4{F0: (v2 | v48), F1: v2, F2: v1}
                                                                        if (!((!true) && ((v2 * v48) < uint32(v42)))) {
                                                                            var v72 float64 = 0.0
                                                                            v2 += v2
                                                                            fmt.Println(fi50)
                                                                            _ = v72
                                                                        } else {
                                                                            switch -9223372036854775808 {
                                                                            case 4294967295:
                                                                                v73 := GenS3{F0: 0, F1: (-(v4))}
                                                                                _ = v73
                                                                            }
                                                                            v74 := GenS3{F0: 248, F1: int64(v48)}
                                                                            if true {
                                                                                var v75 int32 = 116
                                                                                _ = v75
                                                                            } else {
                                                                                v45 = append(v45, 32768)
                                                                                var v76 uint = v42
                                                                                if v69 {
                                                                                    fmt.Println(v46, fi65)
                                                                                } else {
                                                                                    v77 := GenS5{F0: 2147483647}
                                                                                    var v78 string = "foo bar"
                                                                                    _ = v77
                                                                                    _ = v78
                                                                                }
                                                                                _ = v76
                                                                            }
                                                                            _ = v74
                                                                        }
                                                                        v79 := GenS2{F0: v4, F1: 2147483647, F2: 2}
                                                                        _ = v71
                                                                        _ = v79
                                                                    }
                                                                    _ = v67
                                                                    _ = v68
                                                                    _ = v69
                                                                }
                                                            } else {
                                                                if v1 {
                                                                    var v80 float64 = ((1.5 + 1.5) / 1.5)
                                                                    _ = v80
                                                                } else {
                                                                    v2 = (v2 | v48)
                                                                    if v1 {
                                                                        v2 = (v2 * v2)
                                                                        switch "\n\t" {
                                                                        case "":
                                                                            if (v1 && ((true || (v48 < 42)) || v1)) {
                                                                                var v81 int8 = -9
                                                                                v82 := GenS6{F0: 1000, F1: (v2 | v48), F2: 188}
                                                                                v47 -= int64(v81)
                                                                                delete(v3, 4294967295)
                                                                                v81 = 0
                                                                                _ = v81
                                                                                _ = v82
                                                                            } else {
                                                                                var v83 uint64 = 2
                                                                                var v84 int32 = -2147483648
                                                                                _ = v83
                                                                                _ = v84
                                                                            }
                                                                        default:
                                                                            switch (fi50 & fi50) {
                                                                            case 0:
                                                                                v42 = 186
                                                                            case -9223372036854775808:
                                                                                v2 = (v2 + (v2 % (((v48 / (((^(v48))) | 1))) | 1)))
                                                                            default:
                                                                                var v85 bool = v1
                                                                                _ = v85
                                                                            }
                                                                        }
                                                                    }
                                                                    var v86 int32 = 0
                                                                    _ = v86
                                                                }
                                                                if len(v45) > 0 {
                                                                    v45[0] = v64
                                                                }
                                                                var v87 float64 = 0.0
                                                                v88 := GenS2{F0: v46, F1: -1, F2: 2}
                                                                _ = v87
                                                                _ = v88
                                                            }
                                                            _ = fi65
                                                            _ = v66
                                                        }
                                                        if (!v1) {
                                                            v48 += (v2 & (v48 ^ 753))
                                                        }
                                                        v1 = v1
                                                    } else {
                                                        v45 = append(v45, (fi50 % ((v64) | 1)))
                                                    }
                                                } else {
                                                    var v89 uint16 = 65535
                                                    for fi90 := 0; fi90 < 18; fi90++ {
                                                        v91 := GenS6{F0: 32767, F1: 42, F2: 1000}
                                                        _ = fi90
                                                        _ = v91
                                                    }
                                                    fmt.Println(v47, v4)
                                                    _ = v89
                                                }
                                                fmt.Println(v4)
                                                _ = v64
                                            }
                                        case "hello":
                                            v48 -= (v2 >> ((v42) & 63))
                                        case "\n\t":
                                            if len(v45) >= 2 {
                                                v45 = v45[1:len(v45)]
                                            }
                                        }
                                    } else {
                                        for fi92 := 0; fi92 < 6; fi92++ {
                                            var v93 float32 = (0.1 + 0.0)
                                            v46 = (v47 - v47)
                                            _ = fi92
                                            _ = v93
                                        }
                                        v42 = 219
                                        v94 := GenS7{F0: (v42 << ((18446744073709551615) & 63)), F1: v2, F2: fi50}
                                        _ = v94
                                    }
                                    v2 = v2
                                }
                                var v95 uint32 = 0
                                v46 = -782
                                _ = v95
                            }
                            _ = fi50
                        }
                        v1 = (1e10 != 3.14159)
                    }
                } else {
                    v48 -= (v48 & v2)
                    for fi96 := 0; fi96 < 12; fi96++ {
                        v97 := GenS8{F0: 684}
                        var v98 int64 = (v47 % ((9223372036854775807) | 1))
                        if v1 {
                            fmt.Println(fi96, v48, v48)
                            for fi99 := 0; fi99 < 4; fi99++ {
                                switch "\n\t" {
                                case "":
                                    if (9223372036854775807 > fi96) {
                                        var v100 int = -9223372036854775808
                                        v101 := GenS3{F0: 128, F1: 0}
                                        _ = v100
                                        _ = v101
                                    }
                                }
                                _ = fi99
                            }
                        }
                        _ = fi96
                        _ = v97
                        _ = v98
                    }
                    if len(v45) > 0 {
                        v45[0] = -794
                    }
                }
            case "":
                fmt.Println(v4, v47)
            default:
                _ = len(v5)
            }
            var v102 uint16 = 65535
            _ = v102
        } else {
            switch "foo bar" {
            case "":
                v4 -= 0
            case "hello":
                v46 = -351
            case "foo bar":
                v1 = v1
            }
        }
    case 9223372036854775807:
        var v103 bool = v1
        _ = v103
    }
    v42 += (v42 >> ((42) & 63))
    _ = len(v45)
    switch 4294967295 {
    case 1:
        v46 = v4
    }
    delete(v3, 759)
    v47 = 1
    if v1 {
        var v104 bool = ((!true) || ((v2 == v48) || (false || (!v1))))
        switch 128 {
        case -249:
            fmt.Println(v48)
        case 452:
            v45 = append(v45, 42)
        case 9223372036854775807:
            var v105 uint = (v42 & v42)
            _ = v105
        default:
            switch ((("\n\t" + "foo bar") + ("hello" + ("\n\t" + "\n\t"))) + "\n\t") {
            case "hello":
                var v106 int16 = -32768
                _ = v106
            case "":
                v47 = v47
            }
        }
        _ = v104
    } else {
        v2 += v2
        var v107 int8 = 100
        if len(v45) > 0 {
            v45[0] = 181
        }
        _ = v107
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v5
}
