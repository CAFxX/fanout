package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 string; F1 uint32; F2 uint32 }
type GenS5 struct { F0 uint64; F1 uint32 }
type GenS6 struct { F0 int8; F1 uint; F2 uint32 }
type GenS7 struct { F0 int16; F1 uint32 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int; F1 uint64; F2 int16 }
type GenS10 struct { F0 uint; F1 int64; F2 int16 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 int64; F1 uint32 }
type GenS13 struct { F0 int16; F1 int8; F2 int16 }
type GenS14 struct { F0 string }
type GenS15 struct { F0 string; F1 bool }
type GenS16 struct { F0 int32; F1 uint64 }
type GenS17 struct { F0 int; F1 uint8; F2 string }
type GenS18 struct { F0 int32; F1 int }
type GenS19 struct { F0 bool; F1 int64; F2 bool }
func genFn0(p0 int64) bool {
    p0 += (-(p0))
    if (0 > 0) {
        v5 := []int{-9223372036854775808, 228, 2147483647}
        v6 := GenS1{F0: 2, F1: 255}
        _ = v5
        _ = v6
    } else {
        fmt.Println(p0)
        p0 += p0
        var v7 uint16 = 0
        _ = v7
    }
    _ = p0
    return true
}
func genFn1(p0 int8) int16 {
    var v8 uint32 = 188
    switch "hello" {
    case "hello":
        v8 += (v8 * v8)
    default:
        var v9 int32 = 1
        _ = v9
    }
    _ = p0
    _ = v8
    return 0
}

func main() {
    v1 := make(map[int]bool)
    var v2 int = 51
    var v3 uint32 = 127
    var v4 float32 = (1.5 / (((1e10 / float32(3.14159)) / 3.14159) / 1e10))
    var v10 float32 = (v4 - (v4 - 1.5))
    for fi11 := 0; fi11 < 9; fi11++ {
        fmt.Println(v3, v2)
        if (((42 < 516) && (true || (!(true && true)))) && (626 == 0)) {
            v10 -= v4
            var v12 uint64 = 18446744073709551615
            _ = v12
        }
        _ = fi11
    }
    var v13 uint8 = 128
    switch v2 {
    case 9223372036854775807:
        fmt.Println(v13, v3, v10)
    default:
        v14 := GenS2{F0: v2}
        _ = v14
    }
    switch (v2 / (((-(v2))) | 1)) {
    case 0:
        v2 = (v2 + (v2 >> ((2) & 63)))
    default:
        fmt.Println(v10, v10, v3)
    }
    switch "foo bar" {
    case "":
        if true {
            if (("" + "") >= "foo bar") {
                switch (v2 ^ v2) {
                case 1:
                    v15 := v1[(v2 - -9223372036854775808)]
                    _ = v15
                default:
                    fmt.Println(v10, v13)
                }
                switch -1 {
                case 1:
                    var v16 int64 = -9223372036854775808
                    _ = v16
                case 32767:
                    switch v2 {
                    case 1:
                        fmt.Println(v3, v13)
                    case 127:
                        for fi17 := 0; fi17 < 4; fi17++ {
                            if ("\n\t" >= "\n\t") {
                                for fi18 := 0; fi18 < 5; fi18++ {
                                    fmt.Println(v10)
                                    switch 128 {
                                    case 1:
                                        v4 -= (v10 / v4)
                                    case 42:
                                        var v19 float32 = v4
                                        v20 := GenS1{F0: 1, F1: uint(v13)}
                                        switch (v2 * fi18) {
                                        case 32767:
                                            var v21 float64 = 3.14159
                                            _ = v21
                                        case -9223372036854775808:
                                            v22 := GenS3{F0: 100, F1: (^(v2))}
                                            _ = v22
                                        default:
                                            fmt.Println(v10)
                                        }
                                        _ = v19
                                        _ = v20
                                    }
                                    v4 -= (0.1 - v4)
                                    if (!true) {
                                        var v23 string = ("" + "a")
                                        v24 := []int{}
                                        switch v2 {
                                        case 32767:
                                            v4 = 1e10
                                        default:
                                            _ = len(v1)
                                        }
                                        _ = v23
                                        _ = v24
                                    } else {
                                        switch ("hello" + "a") {
                                        case "foo bar":
                                            v25 := GenS4{F0: "\n\t", F1: 2, F2: (v3 % (((^(v3))) | 1))}
                                            _ = v25
                                        case "a":
                                            var v26 bool = (("" + "foo bar") > "\n\t")
                                            _ = v26
                                        case "\n\t":
                                            if false {
                                                var v27 uint8 = 128
                                                _ = v27
                                            } else {
                                                v10 = v4
                                                v28 := GenS1{F0: v3, F1: 1}
                                                if true {
                                                    v29 := GenS5{F0: 18446744073709551615, F1: v3}
                                                    var v30 int64 = 0
                                                    _ = v29
                                                    _ = v30
                                                } else {
                                                    v10 = (((v4 + float32((1.5 / ((1e10 / 3.14159) / float64(v4))))) - (0.1 / v4)) * 0.0)
                                                    v3 += 2147483648
                                                    var v31 int = (fi17 | (fi18 >> ((42) & 63)))
                                                    _ = v31
                                                }
                                                _ = v28
                                            }
                                        default:
                                            if ((-1 < 9223372036854775807) || false) {
                                                if false {
                                                    delete(v1, fi17)
                                                    v32 := v1[v2]
                                                    if v32 {
                                                        var v33 int8 = 100
                                                        if true {
                                                            if false {
                                                                fmt.Println(v2, fi17, v4)
                                                                v34 := GenS4{F0: "hello", F1: 171, F2: (v3 ^ (v3 + (v3 * 320)))}
                                                                _ = v34
                                                            }
                                                            var v35 uint8 = (v13 ^ 1)
                                                            var v36 string = "hello"
                                                            if ((128 > 42) || (((!v32) && (!((fi17 + 1) >= (fi17 & v2)))) || (!(2 < 42)))) {
                                                                v35 += 255
                                                            } else {
                                                                if (v3 >= v3) {
                                                                    if ((65535 == 389) && true) {
                                                                        switch fi17 {
                                                                        case -951:
                                                                            v4 -= ((((v4 - (v10 + 0.1)) * 0.0) + v10) * (float32(1.5) * v10))
                                                                            v13 = (v35 & v13)
                                                                        default:
                                                                            if v32 {
                                                                                v37 := GenS1{F0: 65535, F1: 0}
                                                                                _ = v37
                                                                            } else {
                                                                                fmt.Println(v4, v4)
                                                                                if (v32 || false) {
                                                                                    v38 := GenS5{F0: 18446744073709551615, F1: v3}
                                                                                    v39 := GenS4{F0: (((("\n\t" + v36) + v36) + v36) + ("hello" + (v36 + "\n\t"))), F1: (v3 << ((711) & 63)), F2: v3}
                                                                                    _ = v38
                                                                                    _ = v39
                                                                                } else {
                                                                                    v3 -= v3
                                                                                    v40 := []int{-156, v2, (fi18 * fi17), v2}
                                                                                    var v41 float64 = ((1.5 * 0.1) * -2.25)
                                                                                    _ = v40
                                                                                    _ = v41
                                                                                }
                                                                                fmt.Println(v3, v4, v4)
                                                                            }
                                                                        }
                                                                        var v42 bool = false
                                                                        if v42 {
                                                                            v2 = (fi18 >> ((2) & 63))
                                                                            if (v36 == "hello") {
                                                                                var v43 uint = 65535
                                                                                _ = v43
                                                                            } else {
                                                                                switch (v36 + v36) {
                                                                                case "foo bar":
                                                                                    v10 += 1e10
                                                                                case "a":
                                                                                    v44 := []int{(fi17 + (fi18 | (fi17 * 9223372036854775807)))}
                                                                                    _ = v44
                                                                                case "":
                                                                                    if v42 {
                                                                                        v42 = (false && true)
                                                                                        v13 += 127
                                                                                        if v32 {
                                                                                            switch v36 {
                                                                                            case "foo bar":
                                                                                                v35 = v13
                                                                                            case "a":
                                                                                                var v45 int64 = 42
                                                                                                _ = v45
                                                                                            }
                                                                                            v3 -= 4294967295
                                                                                            v3 = (v3 / (((v3 + (^(v3)))) | 1))
                                                                                        } else {
                                                                                            if true {
                                                                                                switch "" {
                                                                                                case "\n\t":
                                                                                                    var v46 int32 = 2147483647
                                                                                                    _ = v46
                                                                                                case "hello":
                                                                                                    var v47 bool = (1 >= 856)
                                                                                                    _ = v47
                                                                                                default:
                                                                                                    v1[-368] = (((((v10 + 0.0) + (1.5 / v10)) <= float32(0.0)) && false) || ((0.1 * 0.0) != 1e10))
                                                                                                }
                                                                                                var v48 uint = 18446744073709551615
                                                                                                _ = v48
                                                                                            }
                                                                                            if true {
                                                                                                var v49 uint64 = 0
                                                                                                var v50 int8 = v33
                                                                                                var v51 int64 = 42
                                                                                                v3 -= (v3 << ((224) & 63))
                                                                                                _ = v49
                                                                                                _ = v50
                                                                                                _ = v51
                                                                                            }
                                                                                        }
                                                                                        if true {
                                                                                            v52 := GenS2{F0: v2}
                                                                                            _ = v52
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    var v53 uint8 = v13
                                                                                    _ = v53
                                                                                }
                                                                                fmt.Println(fi18)
                                                                            }
                                                                            delete(v1, 32767)
                                                                        }
                                                                        _ = v42
                                                                    }
                                                                    var v54 int8 = (v33 | -1)
                                                                    _ = v54
                                                                }
                                                                if v32 {
                                                                    v32 = (668 < 89)
                                                                    if v32 {
                                                                        var v55 uint = 0
                                                                        _ = v55
                                                                    }
                                                                    v2 += fi18
                                                                } else {
                                                                    switch v36 {
                                                                    case "\n\t":
                                                                        v56 := GenS2{F0: (fi18 + (fi17 << ((0) & 63)))}
                                                                        _ = v56
                                                                    case "a":
                                                                        v36 = (v36 + v36)
                                                                    default:
                                                                        v57 := GenS6{F0: (v33 << ((255) & 63)), F1: 71, F2: v3}
                                                                        _ = v57
                                                                    }
                                                                    switch (v2 << ((42) & 63)) {
                                                                    case 281:
                                                                        v32 = (v36 > ("foo bar" + ""))
                                                                    case 886:
                                                                        switch v36 {
                                                                        case "hello":
                                                                            var v58 bool = (533 < 2)
                                                                            _ = v58
                                                                        default:
                                                                            var v59 uint = 515
                                                                            _ = v59
                                                                        }
                                                                    case -398:
                                                                        var v60 float32 = v10
                                                                        _ = v60
                                                                    default:
                                                                        if v32 {
                                                                            var v61 float64 = 1.5
                                                                            var v62 int = v2
                                                                            _ = v61
                                                                            _ = v62
                                                                        } else {
                                                                            var v63 uint8 = (v35 << ((2) & 63))
                                                                            _ = v63
                                                                        }
                                                                        fmt.Println(v36, v36)
                                                                        v1[fi18] = v32
                                                                        switch v2 {
                                                                        case 42:
                                                                            v4 -= (v4 - 0.0)
                                                                        default:
                                                                            var v64 uint32 = (v3 + 1)
                                                                            v64 = (v3 / (((v64 << ((0) & 63))) | 1))
                                                                            _ = v64
                                                                        }
                                                                        switch -2147483648 {
                                                                        case 42:
                                                                            var v65 uint32 = v3
                                                                            _ = v65
                                                                        case -2147483648:
                                                                            fmt.Println(v2, v4)
                                                                        case 1:
                                                                            v3 += v3
                                                                        default:
                                                                            if ("" >= ("" + v36)) {
                                                                                if (-1 != 695) {
                                                                                    v13 = 2
                                                                                    v4 = ((3.14159 / 0.1) + v4)
                                                                                } else {
                                                                                    v66 := GenS1{F0: v3, F1: 65535}
                                                                                    fmt.Println(v2)
                                                                                    _ = v66
                                                                                }
                                                                            } else {
                                                                                var v67 bool = (("" + v36) != v36)
                                                                                var v68 int = (fi17 & (-(fi18)))
                                                                                _ = v67
                                                                                _ = v68
                                                                            }
                                                                        }
                                                                    }
                                                                    var v69 int64 = -799
                                                                    _ = v69
                                                                }
                                                            }
                                                            if false {
                                                                var v70 int64 = -567
                                                                switch v36 {
                                                                case "":
                                                                    if (v32 || ("foo bar" == v36)) {
                                                                        v71 := GenS3{F0: (v33 ^ v33), F1: v2}
                                                                        var v72 float32 = v4
                                                                        switch -257 {
                                                                        case 42:
                                                                            switch (v2 ^ (^(v2))) {
                                                                            case -237:
                                                                                if true {
                                                                                    var v73 int64 = v70
                                                                                    v73 -= -336
                                                                                    fmt.Println(v73)
                                                                                    _ = v73
                                                                                }
                                                                            case -995:
                                                                                if false {
                                                                                    if (int16(fi18) > 480) {
                                                                                        var v74 int = v2
                                                                                        if ((v33 * v33) >= v33) {
                                                                                            fmt.Println(v33, v36, v74)
                                                                                            v75 := GenS3{F0: (v33 ^ v33), F1: (v2 * 42)}
                                                                                            _ = v75
                                                                                        }
                                                                                        v33 = v33
                                                                                        _ = v74
                                                                                    }
                                                                                    switch fi17 {
                                                                                    case 32768:
                                                                                        v76 := GenS1{F0: 689, F1: 1}
                                                                                        _ = v76
                                                                                    }
                                                                                    switch v36 {
                                                                                    case "":
                                                                                        v13 = (v35 >> ((30) & 63))
                                                                                    }
                                                                                }
                                                                            }
                                                                        default:
                                                                            fmt.Println(v3, v33)
                                                                            fmt.Println(fi17, v10)
                                                                        }
                                                                        var v77 string = v36
                                                                        _ = v71
                                                                        _ = v72
                                                                        _ = v77
                                                                    }
                                                                case "hello":
                                                                    v35 = (^(v35))
                                                                }
                                                                _ = v70
                                                            } else {
                                                                v10 -= 1e10
                                                            }
                                                            fmt.Println(v4)
                                                            if v32 {
                                                                var v78 uint32 = v3
                                                                _ = v78
                                                            }
                                                            _ = v35
                                                            _ = v36
                                                        } else {
                                                            var v79 uint32 = (v3 | (v3 >> ((255) & 63)))
                                                            var v80 uint = 477
                                                            var v81 uint64 = 127
                                                            v2 -= 127
                                                            fmt.Println(v13, fi18, v2)
                                                            delete(v1, -384)
                                                            if v32 {
                                                                switch "hello" {
                                                                case "":
                                                                    v82 := []int{v2, -9223372036854775808, fi17}
                                                                    _ = v82
                                                                }
                                                            } else {
                                                                v32 = true
                                                                v33 += (v33 % ((v33) | 1))
                                                                if v32 {
                                                                    var v83 int64 = -9223372036854775808
                                                                    _ = v83
                                                                }
                                                            }
                                                            if ((^(v81)) != (v81 / (((v81 | v81)) | 1))) {
                                                                v84 := GenS2{F0: (fi17 >> ((v80) & 63))}
                                                                switch (fi17 + 9223372036854775807) {
                                                                case 0:
                                                                    v85 := GenS2{F0: (fi18 * fi17)}
                                                                    _ = v85
                                                                case 2147483647:
                                                                    if v32 {
                                                                        v13 = (v13 * 128)
                                                                    }
                                                                case -9223372036854775808:
                                                                    v3 += (^(v79))
                                                                }
                                                                v1[(-(fi17))] = (!true)
                                                                switch fi17 {
                                                                case 42:
                                                                    v2 -= (v2 % (((fi17 & -827)) | 1))
                                                                case 653:
                                                                    v33 = 100
                                                                }
                                                                v4 = ((v10 / v4) * (v4 - 3.14159))
                                                                _ = v84
                                                            } else {
                                                                v79 = 0
                                                                v86 := GenS5{F0: (v81 + uint64(v3)), F1: (v79 % (((v3 + v79)) | 1))}
                                                                _ = len(v1)
                                                                _ = v86
                                                            }
                                                            switch int(v80) {
                                                            case 4294967295:
                                                                v32 = true
                                                            case 0:
                                                                if (!v32) {
                                                                    v80 = v80
                                                                    var v87 bool = (fi18 > fi18)
                                                                    switch fi17 {
                                                                    case -1:
                                                                        v81 = v81
                                                                    case 2147483647:
                                                                        switch "\n\t" {
                                                                        case "\n\t":
                                                                            var v88 float64 = 0.0
                                                                            _ = v88
                                                                        default:
                                                                            var v89 float32 = v10
                                                                            _ = v89
                                                                        }
                                                                    }
                                                                    _ = v87
                                                                } else {
                                                                    v13 += 128
                                                                    if (true && (-9223372036854775808 <= -1)) {
                                                                        var v90 int64 = 0
                                                                        var v91 bool = v32
                                                                        _ = v90
                                                                        _ = v91
                                                                    }
                                                                }
                                                            case -1:
                                                                var v92 float32 = (v4 - v10)
                                                                _ = v92
                                                            default:
                                                                switch fi18 {
                                                                case -1:
                                                                    var v93 int = (^(v2))
                                                                    _ = v93
                                                                case 2147483647:
                                                                    var v94 int32 = 42
                                                                    _ = v94
                                                                case 0:
                                                                    v80 += v80
                                                                    var v95 string = "foo bar"
                                                                    var v96 uint32 = v3
                                                                    _ = v95
                                                                    _ = v96
                                                                }
                                                            }
                                                            _ = v79
                                                            _ = v80
                                                            _ = v81
                                                        }
                                                        _ = v33
                                                    } else {
                                                        v97 := []int{}
                                                        _ = v97
                                                    }
                                                    _ = v32
                                                } else {
                                                    v98 := GenS7{F0: 1000, F1: (v3 ^ v3)}
                                                    v99 := GenS3{F0: 0, F1: v2}
                                                    var v100 float32 = v10
                                                    if true {
                                                        if false {
                                                            v4 -= ((v10 / 1.5) / 1e10)
                                                        } else {
                                                            v10 = v10
                                                        }
                                                        v100 -= 1e10
                                                        var v101 uint16 = 2
                                                        _ = v101
                                                    } else {
                                                        if false {
                                                            v2 -= fi17
                                                            var v102 int32 = 1
                                                            v3 = (v3 - 2147483647)
                                                            _ = v102
                                                        }
                                                        v103 := GenS8{F0: "hello"}
                                                        if false {
                                                            v10 += v100
                                                            v3 += (v3 / ((0) | 1))
                                                        } else {
                                                            switch "" {
                                                            case "foo bar":
                                                                if (18446744073709551615 <= 18446744073709551615) {
                                                                    fmt.Println(v100)
                                                                    switch "\n\t" {
                                                                    case "foo bar":
                                                                        var v104 int = fi17
                                                                        _ = v104
                                                                    case "a":
                                                                        v1[(fi17 * (v2 << ((2) & 63)))] = (true || false)
                                                                    }
                                                                }
                                                            case "a":
                                                                if (v3 >= 2) {
                                                                    v2 -= fi18
                                                                    var v105 uint16 = 32767
                                                                    _ = v105
                                                                }
                                                            case "\n\t":
                                                                var v106 int8 = 42
                                                                _ = v106
                                                            default:
                                                                if false {
                                                                    switch (^(fi17)) {
                                                                    case -2147483648:
                                                                        if true {
                                                                            switch "a" {
                                                                            case "a":
                                                                                var v107 uint8 = (^(v13))
                                                                                _ = v107
                                                                            default:
                                                                                switch fi18 {
                                                                                case 1:
                                                                                    var v108 int32 = 2147483647
                                                                                    _ = v108
                                                                                }
                                                                            }
                                                                            if false {
                                                                                fmt.Println(v2)
                                                                                _ = len(v1)
                                                                            } else {
                                                                                v109 := GenS6{F0: 127, F1: 18446744073709551615, F2: (v3 & (^(v3)))}
                                                                                _ = v109
                                                                            }
                                                                        } else {
                                                                            v100 -= 0.0
                                                                        }
                                                                    default:
                                                                        var v110 float32 = v4
                                                                        _ = v110
                                                                    }
                                                                    v111 := []int{fi17, (v2 + (fi18 << ((42) & 63))), (v2 << ((65535) & 63)), (^(fi17))}
                                                                    _ = v111
                                                                }
                                                            }
                                                            switch (fi18 % (((v2 ^ fi17)) | 1)) {
                                                            case -352:
                                                                _ = len(v1)
                                                            case 1:
                                                                v13 += (v13 & 128)
                                                            }
                                                        }
                                                        var v112 float32 = (1.5 * 0.1)
                                                        _ = v103
                                                        _ = v112
                                                    }
                                                    _ = v100
                                                    _ = v98
                                                    _ = v99
                                                }
                                                delete(v1, 32768)
                                                v113 := GenS9{F0: (v2 ^ 964), F1: 4294967295, F2: -896}
                                                var v114 int64 = 42
                                                _ = v113
                                                _ = v114
                                            }
                                        }
                                        fmt.Println(v2, v2, v3)
                                        v115 := GenS10{F0: 343, F1: int64(v3), F2: 32767}
                                        var v116 string = ("a" + "hello")
                                        v3 = v3
                                        _ = v115
                                        _ = v116
                                    }
                                    _ = fi18
                                }
                                var v117 float64 = 1.5
                                _ = v117
                            } else {
                                var v118 int16 = 32767
                                v119 := GenS11{F0: (true && false)}
                                _ = v118
                                _ = v119
                            }
                            v120 := []int{(fi17 & 2147483647), (fi17 & v2), (fi17 / ((v2) | 1)), fi17}
                            var v121 float64 = 0.1
                            _ = fi17
                            _ = v120
                            _ = v121
                        }
                    }
                case 42:
                    v3 += (v3 % (((v3 + v3)) | 1))
                default:
                    v122 := []int{42}
                    _ = v122
                }
            }
            v123 := GenS12{F0: 42, F1: v3}
            _ = v123
        } else {
            v124 := GenS7{F0: 1, F1: 2147483647}
            if (-113 == genFn1(0)) {
                var v125 float64 = 0.1
                fmt.Println(v4, v3, v2)
                _ = v125
            } else {
                for fi126 := 0; fi126 < 17; fi126++ {
                    v13 -= (v13 * (v13 >> ((18446744073709551615) & 63)))
                    _ = fi126
                }
                v127 := GenS5{F0: 1000, F1: 736}
                _ = v127
            }
            _ = v124
        }
    case "foo bar":
        switch -2147483648 {
        case 129:
            v3 = 2147483647
        case -2147483648:
            for fi128 := 0; fi128 < 20; fi128++ {
                v3 = 42
                switch "hello" {
                case "\n\t":
                    v2 += v2
                case "hello":
                    for fi129 := 0; fi129 < 19; fi129++ {
                        v2 = fi128
                        _ = fi129
                    }
                }
                _ = fi128
            }
        case -523:
            var v130 float64 = 0.1
            _ = v130
        }
    case "hello":
        for fi131 := 0; fi131 < 13; fi131++ {
            v132 := v1[-778]
            fmt.Println(v3, v132, v2)
            v133 := GenS12{F0: 1, F1: 0}
            _ = fi131
            _ = v132
            _ = v133
        }
    default:
        var v134 string = "a"
        _ = v134
    }
    for fi135 := 0; fi135 < 6; fi135++ {
        var v136 int8 = 100
        fmt.Println(v13, v13, v2)
        _ = fi135
        _ = v136
    }
    v3 = (v3 - 65535)
    _ = len(v1)
    switch v2 {
    case 127:
        for fi137 := 0; fi137 < 14; fi137++ {
            if (false || false) {
                var v138 uint32 = v3
                switch fi137 {
                case -1:
                    v10 += (v10 + 0.1)
                }
                _ = v138
            } else {
                v139 := []int{fi137, 478, (fi137 << ((807) & 63))}
                _ = v139
            }
            var v140 uint8 = 2
            v2 = int(v13)
            _ = fi137
            _ = v140
        }
    }
    for fi141 := 0; fi141 < 1; fi141++ {
        var v142 int8 = -19
        switch 32768 {
        case 8:
            if (true && false) {
                if (-1 < -1) {
                    v4 -= v4
                }
                var v143 uint = 65535
                switch 127 {
                case 4294967295:
                    fmt.Println(v3, v4, v10)
                }
                _ = v143
            }
        case 963:
            for fi144 := 0; fi144 < 4; fi144++ {
                if false {
                    fmt.Println(v4)
                    switch ("" + "") {
                    case "foo bar":
                        switch (("foo bar" + (("foo bar" + ("foo bar" + (("a" + "\n\t") + "hello"))) + "a")) + "hello") {
                        case "hello":
                            var v145 int64 = -9223372036854775808
                            _ = v145
                        case "\n\t":
                            v4 -= ((v4 + v4) * ((1.5 - (0.0 - v4)) / v4))
                        default:
                            v146 := GenS13{F0: 32767, F1: 0, F2: int16(v13)}
                            v147 := GenS14{F0: "\n\t"}
                            _ = v146
                            _ = v147
                        }
                    case "\n\t":
                        switch fi141 {
                        case 32768:
                            v2 += (fi144 * fi141)
                        case -2147483648:
                            if (529 < 2) {
                                if (18446744073709551615 == 812) {
                                    var v148 float64 = (float64((v10 - v4)) * (0.0 + 0.0))
                                    var v149 int8 = 127
                                    v148 = v148
                                    var v150 uint32 = v3
                                    _ = v148
                                    _ = v149
                                    _ = v150
                                } else {
                                    var v151 float64 = -2.25
                                    v152 := []int{781}
                                    _ = v151
                                    _ = v152
                                }
                                var v153 bool = false
                                _ = v153
                            } else {
                                v3 = 8
                            }
                        case 1:
                            var v154 int32 = 0
                            _ = v154
                        }
                        fmt.Println(v10, v13, v13)
                        if (!false) {
                            var v155 int64 = -1
                            if ((v155 & v155) >= v155) {
                                v4 -= (v10 * (0.1 / 1.0))
                                fmt.Println(v142, v10, v10)
                                switch (v2 | int(v13)) {
                                case -919:
                                    if true {
                                        if (239 == 0) {
                                            if true {
                                                switch ("\n\t" + "a") {
                                                case "\n\t":
                                                    if true {
                                                        switch (^(v2)) {
                                                        case -1:
                                                            v156 := GenS15{F0: "foo bar", F1: true}
                                                            _ = v156
                                                        case 4294967295:
                                                            switch (v2 * (fi141 % ((fi141) | 1))) {
                                                            case 32767:
                                                                if (127 < (^(v13))) {
                                                                    var v157 uint32 = 354
                                                                    switch -99 {
                                                                    case -9223372036854775808:
                                                                        v155 -= v155
                                                                    case 42:
                                                                        v158 := v1[(v2 * 128)]
                                                                        _ = v158
                                                                    }
                                                                    v10 = float32((1e10 * 1.5))
                                                                    _ = v157
                                                                }
                                                            default:
                                                                v4 += ((v4 - -2.25) / v4)
                                                            }
                                                        case 2147483647:
                                                            switch "\n\t" {
                                                            case "":
                                                                v1[-9223372036854775808] = (421 <= 665)
                                                            }
                                                        }
                                                        var v159 uint32 = (v3 / ((2147483648) | 1))
                                                        v142 = -1
                                                        _ = v159
                                                    }
                                                }
                                                v160 := GenS7{F0: int16(v155), F1: (v3 | (v3 % (((v3 / ((816) | 1))) | 1)))}
                                                _ = v160
                                            }
                                            switch "hello" {
                                            case "hello":
                                                if false {
                                                    if (true || true) {
                                                        v161 := GenS16{F0: 2147483647, F1: 1}
                                                        delete(v1, (-(fi141)))
                                                        _ = v161
                                                    }
                                                    fmt.Println(v155, fi144, fi144)
                                                } else {
                                                    v162 := []int{(fi144 & v2)}
                                                    v162 = append(v162, 9223372036854775807)
                                                    v163 := GenS17{F0: v2, F1: (v13 >> ((uint(v13)) & 63)), F2: ""}
                                                    var v164 int16 = 0
                                                    _ = v162
                                                    _ = v163
                                                    _ = v164
                                                }
                                            default:
                                                v142 = v142
                                            }
                                        }
                                    }
                                }
                            } else {
                                v165 := GenS2{F0: (fi141 * (fi144 - (-(fi141))))}
                                fmt.Println(v10, v3)
                                _ = v165
                            }
                            _ = v155
                        }
                    default:
                        v166 := GenS8{F0: ("" + "hello")}
                        _ = v166
                    }
                    if true {
                        v2 = fi141
                        v167 := []int{(v2 ^ (fi144 / ((fi141) | 1))), (v2 << ((18446744073709551615) & 63)), fi141}
                        v168 := GenS18{F0: -2147483648, F1: v2}
                        switch "" {
                        case "a":
                            fmt.Println(v10, v3, v13)
                        case "hello":
                            switch (("" + ("hello" + "hello")) + ("" + ("" + ("\n\t" + (("a" + "hello") + "a"))))) {
                            case "a":
                                v13 = (v13 % (((^(v13))) | 1))
                            case "":
                                _ = len(v1)
                            case "\n\t":
                                var v169 int32 = 0
                                _ = v169
                            default:
                                delete(v1, (fi141 + (v2 >> ((65535) & 63))))
                            }
                        default:
                            if false {
                                var v170 float64 = float64((v4 * (v10 + (0.0 + (v10 / v10)))))
                                var v171 uint32 = (v3 ^ (v3 | v3))
                                var v172 float64 = v170
                                _ = v170
                                _ = v171
                                _ = v172
                            }
                        }
                        if true {
                            if true {
                                v2 = -409
                            } else {
                                var v173 string = "hello"
                                _ = v173
                            }
                            v2 = fi144
                        }
                        _ = v167
                        _ = v168
                    } else {
                        v174 := GenS2{F0: (-(v2))}
                        _ = v174
                    }
                }
                fmt.Println(v13)
                _ = fi144
            }
        case 516:
            fmt.Println(v13, v10, v142)
        default:
            switch "foo bar" {
            case "a":
                var v175 float32 = ((float32(3.14159) * v4) / (v10 / 1.0))
                _ = v175
            case "":
                v13 -= (^(v13))
            }
        }
        _ = fi141
        _ = v142
    }
    fmt.Println(v10, v13)
    if true {
        fmt.Println(v3)
        for fi176 := 0; fi176 < 2; fi176++ {
            var v177 int16 = 289
            v178 := GenS13{F0: 667, F1: -128, F2: 32767}
            switch (^(fi176)) {
            case 9223372036854775807:
                v179 := []int{9223372036854775807, fi176}
                _ = v179
            case 381:
                var v180 int = (^(v2))
                _ = v180
            case 1:
                switch "foo bar" {
                case "\n\t":
                    v10 -= (v10 + v4)
                case "a":
                    v2 -= v2
                default:
                    v181 := []int{42, (-(v2))}
                    _ = v181
                }
            default:
                v10 = v10
            }
            _ = fi176
            _ = v177
            _ = v178
        }
        switch 767 {
        case 0:
            v182 := GenS7{F0: genFn1(1), F1: (v3 * 4294967295)}
            _ = v182
        case 1:
            fmt.Println(v13, v4, v10)
        default:
            fmt.Println(v2, v10)
        }
    }
    v183 := GenS19{F0: false, F1: 0, F2: true}
    v10 -= v10
    var v184 int64 = 9223372036854775807
    v184 += v184
    for fi185 := 0; fi185 < 11; fi185++ {
        var v186 int32 = -608
        var v187 uint = 18446744073709551615
        var v188 int32 = 2147483647
        _ = fi185
        _ = v186
        _ = v187
        _ = v188
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v184:", v184)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v13
    _ = v183
    _ = v184
    _ = v2
    _ = v3
    _ = v4
}
