package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint16; F1 uint32 }
type GenS6 struct { F0 int16; F1 uint32 }
type GenS7 struct { F0 int8; F1 int64; F2 uint8 }
type GenS8 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS9 struct { F0 int8; F1 int64; F2 int16 }
func genFn0(p0 string) bool {
    v6 := []int{-101}
    for fi7 := 0; fi7 < 20; fi7++ {
        if false {
            if (!(2 > 701)) {
                var v8 uint8 = 255
                var v9 uint32 = 521
                fmt.Println(fi7)
                _ = v8
                _ = v9
            } else {
                v10 := GenS1{F0: 1}
                _ = v10
            }
            if true {
                for fi11 := 0; fi11 < 20; fi11++ {
                    var v12 float64 = (float64(0.0) * ((0.1 - -2.25) * (0.1 - 3.14159)))
                    switch p0 {
                    case "\n\t":
                        fmt.Println(fi7, fi11, fi7)
                    case "a":
                        p0 = p0
                    case "foo bar":
                        var v13 uint = 65535
                        _ = v13
                    }
                    v12 += v12
                    _ = fi11
                    _ = v12
                }
                if (-146 < 1) {
                    for fi14 := 0; fi14 < 18; fi14++ {
                        if true {
                            var v15 uint16 = 255
                            v15 = (v15 ^ v15)
                            var v16 uint32 = 1
                            _ = v15
                            _ = v16
                        }
                        if (686 != 42) {
                            v17 := GenS2{F0: 65535, F1: false}
                            _ = v17
                        } else {
                            var v18 float32 = 0.1
                            _ = v18
                        }
                        var v19 uint64 = 593
                        _ = fi14
                        _ = v19
                    }
                    var v20 float64 = 3.14159
                    _ = v20
                }
                var v21 int16 = -1
                _ = v21
            }
            for fi22 := 0; fi22 < 12; fi22++ {
                var v23 uint32 = 0
                _ = fi22
                _ = v23
            }
        } else {
            var v24 uint32 = 339
            _ = v24
        }
        _ = fi7
    }
    _ = p0
    _ = v6
    return true
}

func main() {
    var v1 int16 = 1
    var v2 int8 = 100
    v3 := make(map[string]string)
    var v4 int64 = 42
    v5 := make(map[uint32]uint32)
    var v25 bool = false
    v26 := GenS3{F0: (((0.0 - 0.1) < -2.25) || v25), F1: 42}
    var v27 string = ""
    for fi28 := 0; fi28 < 18; fi28++ {
        var v29 int8 = 43
        var v30 float32 = -2.25
        switch (("foo bar" + (("hello" + "foo bar") + "foo bar")) + v27) {
        case "foo bar":
            v31 := []int{(fi28 * (fi28 * fi28))}
            _ = v31
        case "hello":
            v27 = v27
        case "a":
            v29 -= v2
        default:
            v1 -= -643
        }
        _ = fi28
        _ = v29
        _ = v30
    }
    switch -56 {
    case 4294967295:
        v2 -= (v2 >> ((1) & 63))
    default:
        v25 = v25
    }
    v32 := GenS3{F0: false, F1: 4294967295}
    fmt.Println(v2, v1, v4)
    var v33 int = -50
    for fi34 := 0; fi34 < 2; fi34++ {
        if ((-32768 >= (^(v1))) || false) {
            if (-9223372036854775808 <= (v4 + v4)) {
                if v25 {
                    switch v27 {
                    case "foo bar":
                        v35 := []int{(-(v33))}
                        _ = v35
                    case "a":
                        for fi36 := 0; fi36 < 20; fi36++ {
                            v33 += (-(fi34))
                            switch v27 {
                            case "\n\t":
                                v37 := []int{-371, fi34}
                                _ = v37
                            case "a":
                                if ((1 == 776) && (false || true)) {
                                    v4 += (v4 + (v4 | 1))
                                    v1 = 139
                                    v38 := GenS1{F0: 42}
                                    _ = v38
                                }
                            }
                            _ = fi36
                        }
                    case "\n\t":
                        delete(v3, (v27 + (v27 + v27)))
                    }
                    if v25 {
                        v1 = -1
                        if false {
                            v39 := []int{}
                            v40 := GenS2{F0: 2, F1: ((v25 && (!(65535 != 2))) || (v25 && ((fi34 % ((-118) | 1)) >= 4294967295)))}
                            _ = v39
                            _ = v40
                        }
                    }
                } else {
                    v41 := []int{128, (v33 << ((2) & 63)), 32767}
                    _ = v41
                }
                v2 = (v2 >> ((2) & 63))
            } else {
                for fi42 := 0; fi42 < 5; fi42++ {
                    v1 = -186
                    switch v27 {
                    case "foo bar":
                        if (!v25) {
                            switch (fi42 << ((2) & 63)) {
                            case -1:
                                fmt.Println(v4, v1, v2)
                            default:
                                switch v27 {
                                case "\n\t":
                                    v27 = (v27 + (v27 + v27))
                                case "a":
                                    fmt.Println(fi42, v33)
                                case "":
                                    var v43 int16 = v1
                                    _ = v43
                                default:
                                    v44 := []int{(v33 % (((v33 & -9223372036854775808)) | 1)), 9223372036854775807, int(v4), fi42}
                                    _ = v44
                                }
                            }
                            v4 = v4
                            if ((^(fi42)) > (fi34 % (((v33 * 253)) | 1))) {
                                var v45 float32 = float32(float64(float32(float64((float32(1e10) + 3.14159)))))
                                fmt.Println(v4, v4, v2)
                                _ = v45
                            }
                        } else {
                            fmt.Println(v27, v4, v25)
                        }
                        var v46 float64 = 0.0
                        if v25 {
                            var v47 float32 = (((0.0 * (3.14159 + float32((v46 - ((v46 + (v46 * v46)) * (0.0 * (v46 + v46))))))) * 1.5) / 1.0)
                            fmt.Println(v2, v33)
                            if v25 {
                                var v48 int = -423
                                switch v27 {
                                case "a":
                                    if (0 < 1) {
                                        switch fi34 {
                                        case 4294967295:
                                            switch 153 {
                                            case 2147483647:
                                                v49 := GenS2{F0: 2147483648, F1: true}
                                                _ = v49
                                            case -9223372036854775808:
                                                v50 := GenS1{F0: 0}
                                                v51 := GenS1{F0: 128}
                                                switch (v27 + "a") {
                                                case "":
                                                    switch v48 {
                                                    case 4294967295:
                                                        var v52 float32 = 1e10
                                                        _ = v52
                                                    case 128:
                                                        var v53 uint32 = 434
                                                        _ = v53
                                                    }
                                                default:
                                                    if true {
                                                        switch "hello" {
                                                        case "foo bar":
                                                            var v54 float32 = (v47 - v47)
                                                            _ = v54
                                                        default:
                                                            v55 := []int{1, v33, v33, v33}
                                                            _ = v55
                                                        }
                                                    } else {
                                                        var v56 int32 = -1
                                                        v25 = v25
                                                        var v57 int16 = v1
                                                        fmt.Println(v48, v33)
                                                        _ = v56
                                                        _ = v57
                                                    }
                                                    v58 := []int{}
                                                    _ = v58
                                                }
                                                _ = v50
                                                _ = v51
                                            }
                                        case 2147483647:
                                            v59 := GenS4{F0: 2}
                                            _ = v59
                                        }
                                        var v60 uint = 65535
                                        _ = v60
                                    } else {
                                        fmt.Println(v1)
                                        if v25 {
                                            v61 := GenS1{F0: 42}
                                            v62 := []int{(v33 & 1), fi42, -9223372036854775808}
                                            _ = v61
                                            _ = v62
                                        } else {
                                            var v63 float64 = v46
                                            var v64 int16 = (v1 >> ((65535) & 63))
                                            v27 = ((v27 + v27) + "foo bar")
                                            _ = v63
                                            _ = v64
                                        }
                                        switch (fi34 % ((789) | 1)) {
                                        case 2147483647:
                                            switch (v27 + (v27 + v27)) {
                                            case "\n\t":
                                                v65 := GenS4{F0: 65535}
                                                _ = v65
                                            case "hello":
                                                v47 += (1e10 - v47)
                                            }
                                        default:
                                            if v25 {
                                                delete(v5, 2147483647)
                                            }
                                            v48 -= fi34
                                            switch (fi34 - fi42) {
                                            case 2147483647:
                                                fmt.Println(v47)
                                            case -366:
                                                var v66 uint32 = 42
                                                _ = v66
                                            case -630:
                                                if v25 {
                                                    if ((v25 && (!(!false))) || v25) {
                                                        var v67 uint = 1
                                                        fmt.Println(v67, v47, v27)
                                                        _ = v67
                                                    }
                                                    var v68 uint = 2
                                                    _ = v68
                                                } else {
                                                    var v69 int32 = 2147483647
                                                    _ = v69
                                                }
                                            default:
                                                v33 += (v33 & (fi34 ^ fi42))
                                            }
                                            var v70 uint32 = 925
                                            v71 := GenS5{F0: 688, F1: (v70 << ((255) & 63))}
                                            _ = v70
                                            _ = v71
                                        }
                                    }
                                    switch (fi42 << ((18446744073709551615) & 63)) {
                                    case 127:
                                        var v72 bool = (1 == 934)
                                        _ = v72
                                    case 42:
                                        v73 := []int{(^(fi42)), -2147483648}
                                        _ = v73
                                    default:
                                        var v74 int8 = (v2 << ((4294967295) & 63))
                                        _ = v74
                                    }
                                    var v75 float32 = v47
                                    _ = v75
                                }
                                _ = len(v5)
                                _ = v48
                            } else {
                                fmt.Println(fi34, v47, v47)
                                v76 := []int{}
                                if v25 {
                                    delete(v3, (v27 + v27))
                                    fmt.Println(v27, v27, fi42)
                                    var v77 string = (v27 + (("hello" + v27) + ("" + (v27 + (v27 + "")))))
                                    _ = v77
                                }
                                _ = v76
                            }
                            _ = v47
                        }
                        _ = v46
                    case "\n\t":
                        v1 = v1
                    }
                    _ = fi42
                }
                fmt.Println(v25)
                v27 = (v27 + v27)
            }
            if (!v25) {
                var v78 float64 = 3.14159
                _ = v78
            } else {
                v27 = "hello"
                v79 := GenS6{F0: 1, F1: 425}
                v80 := GenS7{F0: -38, F1: v4, F2: 13}
                _ = v79
                _ = v80
            }
            v81 := GenS7{F0: (v2 << ((0) & 63)), F1: -1, F2: 42}
            _ = v81
        } else {
            fmt.Println(v1, v27)
        }
        switch fi34 {
        case 2147483647:
            v33 = v33
        default:
            v82 := []int{-10, -1}
            _ = v82
        }
        _ = fi34
    }
    var v83 uint = 576
    var v84 float64 = float64(((float32(1e10) * (-2.25 * 1e10)) / (0.1 / float32((((0.0 + (1e10 - 0.0)) / float64((3.14159 + 1e10))) - float64(float32(-2.25)))))))
    fmt.Println(v2, v27)
    if (!v25) {
        var v85 int16 = (^(v1))
        v86 := GenS8{F0: (v85 & v1), F1: 190, F2: 42}
        _ = v85
        _ = v86
    }
    var v87 int = int(v83)
    v4 -= -229
    v83 -= 827
    for fi88 := 0; fi88 < 7; fi88++ {
        for fi89 := 0; fi89 < 5; fi89++ {
            if v25 {
                if true {
                    var v90 float64 = v84
                    var v91 int16 = 32767
                    v92 := GenS3{F0: v25, F1: 2}
                    _ = v90
                    _ = v91
                    _ = v92
                } else {
                    if (3.14159 == 1.5) {
                        var v93 int8 = (v2 * (v2 * v2))
                        v94 := GenS3{F0: true, F1: 2}
                        fmt.Println(v27, v25)
                        _ = v93
                        _ = v94
                    }
                    if (!(18446744073709551615 == 18446744073709551615)) {
                        fmt.Println(v27)
                        v95 := v3["a"]
                        _ = v95
                    } else {
                        if (-2147483648 >= 0) {
                            v4 += (v4 << ((4294967295) & 63))
                            var v96 uint32 = 2
                            _ = v96
                        } else {
                            v97 := GenS9{F0: (v2 - (v2 ^ (^(v2)))), F1: (v4 % (((v4 * int64(fi88))) | 1)), F2: (v1 % (((v1 & v1)) | 1))}
                            fmt.Println(v2)
                            _ = v97
                        }
                        v3[(v27 + ("\n\t" + "foo bar"))] = (v27 + "a")
                        v98 := v5[987]
                        _ = v98
                    }
                    if v25 {
                        if (4294967295 <= uint64(v83)) {
                            v87 += v33
                        }
                    }
                }
                var v99 string = v27
                var v100 uint32 = 2147483648
                v101 := []int{fi88, 691}
                _ = v100
                _ = v101
                _ = v99
            }
            _ = fi89
        }
        _ = fi88
    }
    var v102 int8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v102:", v102)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v102
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v83
    _ = v84
    _ = v87
}
