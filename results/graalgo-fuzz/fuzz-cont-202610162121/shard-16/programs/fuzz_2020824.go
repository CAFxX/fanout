package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 int64; F1 uint16; F2 int }
type GenS5 struct { F0 uint32; F1 string; F2 string }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 bool; F1 uint64; F2 int }
type GenS10 struct { F0 int; F1 uint64; F2 int8 }
type GenS11 struct { F0 int16; F1 int32 }
type GenS12 struct { F0 int8; F1 uint16; F2 int64 }
type GenS13 struct { F0 string; F1 int64; F2 int64 }
type GenS14 struct { F0 int32; F1 bool }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]bool)
    var v3 int = -1
    v4 := make(map[uint32]uint32)
    var v5 string = "\n\t"
    v5 = v5
    fmt.Println(v5, v5)
    if false {
        if false {
            v6 := GenS1{F0: 1}
            _ = v6
        } else {
            switch "a" {
            case "\n\t":
                for fi7 := 0; fi7 < 13; fi7++ {
                    if ((!(2 > 65535)) && true) {
                        fmt.Println(fi7, v3)
                    }
                    _ = fi7
                }
            case "":
                var v8 uint16 = 857
                _ = v8
            default:
                v4[42] = 42
            }
            for fi9 := 0; fi9 < 18; fi9++ {
                v10 := GenS2{F0: 4294967295, F1: 2147483647}
                _ = fi9
                _ = v10
            }
        }
        var v11 bool = false
        if v11 {
            switch (v5 + (v5 + v5)) {
            case "foo bar":
                switch (v3 - (v3 - (v3 | v3))) {
                case 634:
                    v11 = v11
                default:
                    v12 := GenS3{F0: 32767, F1: 65535}
                    _ = v12
                }
            case "a":
                fmt.Println(v5, v11, v11)
            case "\n\t":
                v13 := GenS3{F0: 341, F1: 2147483647}
                _ = v13
            default:
                v5 = ("\n\t" + v5)
            }
            for fi14 := 0; fi14 < 18; fi14++ {
                if v11 {
                    v11 = v11
                    fmt.Println(fi14, fi14)
                }
                _ = fi14
            }
            v15 := GenS4{F0: 0, F1: 0, F2: (v3 >> ((462) & 63))}
            _ = v15
        } else {
            v16 := GenS1{F0: 0}
            _ = v16
        }
        _ = v11
    } else {
        v5 = (("\n\t" + v5) + ((("\n\t" + ("" + v5)) + v5) + ((v5 + (v5 + v5)) + v5)))
        v5 = v5
        switch v3 {
        case -802:
            var v17 bool = false
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 6; fi18++ {
        switch -1 {
        case 42:
            if false {
                v3 = (v3 / ((fi18) | 1))
            }
        default:
            delete(v2, (v5 + v5))
        }
        for fi19 := 0; fi19 < 7; fi19++ {
            var v20 uint64 = 18446744073709551615
            v3 = int(v20)
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    if false {
        switch (v3 * 724) {
        case 32767:
            v21 := GenS5{F0: 92, F1: ("\n\t" + v5), F2: v5}
            _ = v21
        case 42:
            v22 := GenS5{F0: 2147483648, F1: v5, F2: v5}
            _ = v22
        default:
            for fi23 := 0; fi23 < 5; fi23++ {
                if true {
                    var v24 string = ((((v5 + v5) + v5) + v5) + v5)
                    var v25 bool = false
                    _ = v24
                    _ = v25
                }
                _ = fi23
            }
        }
        v3 = -724
    }
    switch v3 {
    case -122:
        var v26 uint16 = 255
        _ = v26
    case -1:
        fmt.Println(v3, v3)
    case -2147483648:
        var v27 uint = 2
        _ = v27
    }
    if true {
        var v28 uint32 = 2147483647
        v3 += v3
        if (!(42 < -1)) {
            v3 += v3
            for fi29 := 0; fi29 < 7; fi29++ {
                var v30 int32 = -981
                v31 := GenS5{F0: 1, F1: v5, F2: (v5 + v5)}
                v3 += (v3 % ((v3) | 1))
                _ = fi29
                _ = v30
                _ = v31
            }
            if false {
                v5 = v5
                for fi32 := 0; fi32 < 6; fi32++ {
                    var v33 int64 = -910
                    _ = fi32
                    _ = v33
                }
                v5 = v5
            } else {
                v34 := []int{}
                var v35 uint = 2
                _ = v34
                _ = v35
            }
        } else {
            var v36 int64 = 818
            v37 := GenS2{F0: 103, F1: 1}
            var v38 int64 = (v36 + (v36 % (((v36 % (((v36 ^ v36)) | 1))) | 1)))
            _ = v36
            _ = v37
            _ = v38
        }
        _ = v28
    } else {
        v5 = "foo bar"
        switch (v3 / (((v3 | v3)) | 1)) {
        case -9223372036854775808:
            var v39 uint64 = 0
            _ = v39
        case 32768:
            var v40 uint16 = 1
            _ = v40
        case 32767:
            v41 := GenS6{F0: 65535}
            _ = v41
        default:
            v42 := []int{}
            _ = v42
        }
        if true {
            v43 := GenS6{F0: 2}
            fmt.Println(v5)
            _ = v43
        }
    }
    v5 = v5
    var v44 uint64 = 1
    for fi45 := 0; fi45 < 20; fi45++ {
        v46 := GenS7{F0: 1}
        if (32767 != 1) {
            v5 = (v5 + "foo bar")
            for fi47 := 0; fi47 < 8; fi47++ {
                var v48 int32 = 42
                if (32767 >= -1) {
                    if false {
                        fmt.Println(v3)
                        fmt.Println(fi45)
                    }
                } else {
                    if (!true) {
                        if false {
                            var v49 uint32 = 0
                            if false {
                                var v50 int32 = (v48 >> ((163) & 63))
                                if (!((!false) || true)) {
                                    var v51 float32 = (float32(0.0) - 0.1)
                                    _ = v51
                                } else {
                                    v49 -= 42
                                    v52 := v2[v5]
                                    _ = v52
                                }
                                var v53 string = v5
                                var v54 uint = 65535
                                var v55 int32 = -515
                                _ = v50
                                _ = v53
                                _ = v54
                                _ = v55
                            }
                            _ = v49
                        }
                        v48 += v48
                        if false {
                            v3 += (fi47 + 127)
                        }
                    } else {
                        v3 = fi47
                        fmt.Println(v44)
                        var v56 int64 = 9223372036854775807
                        if true {
                            if false {
                                var v57 int = (fi47 >> ((0) & 63))
                                _ = v57
                            } else {
                                var v58 float64 = 3.14159
                                _ = v58
                            }
                            if true {
                                var v59 float64 = ((((0.0 * 0.1) - (0.0 / (3.14159 - 0.0))) - float64((((-2.25 * (1e10 + -2.25)) / ((float32(((float64(1e10) / float64((-2.25 + ((-2.25 / 1e10) - ((3.14159 + -2.25) - ((((3.14159 - (3.14159 / 1.5)) * 0.1) + (float32(0.0) - 0.1)) * (-2.25 / 1.0))))))) + 0.0)) - 0.1) + 0.1)) - 1e10))) + 3.14159)
                                switch "a" {
                                case "hello":
                                    fmt.Println(v48, v44)
                                case "a":
                                    if (-514 < 0) {
                                        v60 := GenS8{F0: 4294967295}
                                        var v61 int8 = 1
                                        v48 = v48
                                        v62 := []int{(^(v3)), 32768, fi45, (v3 | (^(v3)))}
                                        if ((4294967295 == 0) || (!true)) {
                                            switch v5 {
                                            case "":
                                                if true {
                                                    v44 += (v44 << ((65535) & 63))
                                                }
                                            case "foo bar":
                                                var v63 uint64 = 345
                                                _ = v63
                                            case "a":
                                                var v64 int8 = 70
                                                _ = v64
                                            default:
                                                v3 = fi47
                                            }
                                        }
                                        _ = v60
                                        _ = v61
                                        _ = v62
                                    } else {
                                        _ = len(v2)
                                        switch "foo bar" {
                                        case "foo bar":
                                            fmt.Println(v3, v5, v44)
                                        case "hello":
                                            v65 := GenS9{F0: true, F1: 407, F2: (fi45 | -895)}
                                            _ = v65
                                        case "":
                                            v59 += v59
                                        }
                                    }
                                }
                                _ = v59
                            }
                        } else {
                            var v66 uint64 = (v44 << ((4294967295) & 63))
                            v67 := []int{v3, fi47, -630, fi47}
                            switch "a" {
                            case "hello":
                                v3 = fi47
                            case "a":
                                v68 := GenS1{F0: 2147483648}
                                _ = v68
                            default:
                                var v69 bool = false
                                _ = v69
                            }
                            _ = v66
                            _ = v67
                        }
                        if false {
                            v70 := GenS3{F0: 255, F1: 42}
                            if true {
                                _ = len(v4)
                                fmt.Println(v48)
                                if false {
                                    if (!(fi47 != v3)) {
                                        fmt.Println(v44)
                                        var v71 int16 = 1
                                        switch v5 {
                                        case "\n\t":
                                            _ = len(v1)
                                        case "":
                                            var v72 int8 = int8(fi47)
                                            _ = v72
                                        }
                                        var v73 uint = 4294967295
                                        v56 += v56
                                        v71 -= v71
                                        _ = v71
                                        _ = v73
                                    } else {
                                        switch v5 {
                                        case "hello":
                                            if false {
                                                v74 := []int{(fi45 % ((298) | 1)), v3, (fi47 & v3)}
                                                var v75 uint64 = v44
                                                var v76 uint = 4294967295
                                                _ = v74
                                                _ = v75
                                                _ = v76
                                            }
                                        case "a":
                                            if (4294967295 != 255) {
                                                var v77 uint64 = (v44 % (((^(v44))) | 1))
                                                var v78 uint16 = 0
                                                if false {
                                                    v44 -= 1
                                                }
                                                _ = v77
                                                _ = v78
                                            } else {
                                                _ = len(v2)
                                                switch (v5 + v5) {
                                                case "":
                                                    fmt.Println(v56, v56)
                                                default:
                                                    var v79 float64 = 0.1
                                                    _ = v79
                                                }
                                            }
                                        case "foo bar":
                                            switch "a" {
                                            case "a":
                                                var v80 int = v3
                                                _ = v80
                                            default:
                                                v81 := []int{fi47}
                                                _ = v81
                                            }
                                            var v82 uint8 = 102
                                            _ = v82
                                        }
                                    }
                                }
                            } else {
                                if false {
                                    var v83 uint32 = 65535
                                    switch v5 {
                                    case "":
                                        v83 += 4294967295
                                    case "a":
                                        _ = len(v4)
                                    }
                                    _ = v83
                                } else {
                                    if true {
                                        v3 += fi47
                                        var v84 uint8 = 0
                                        v85 := GenS5{F0: 1, F1: ("foo bar" + v5), F2: v5}
                                        _ = v84
                                        _ = v85
                                    } else {
                                        v86 := GenS10{F0: (^(v3)), F1: v44, F2: 42}
                                        v87 := GenS9{F0: (((v56 ^ 1) != -361) && (42 != v56)), F1: (v44 - 1), F2: (v3 - 2147483647)}
                                        var v88 int16 = 32767
                                        _ = v86
                                        _ = v87
                                        _ = v88
                                    }
                                }
                                if true {
                                    if false {
                                        v56 += 639
                                    } else {
                                        if true {
                                            switch ("\n\t" + v5) {
                                            case "hello":
                                                v89 := []int{v3, (fi45 * fi47)}
                                                _ = v89
                                            case "":
                                                switch "hello" {
                                                case "":
                                                    var v90 float32 = (3.14159 + 1e10)
                                                    _ = v90
                                                case "foo bar":
                                                    v91 := GenS9{F0: true, F1: v44, F2: fi47}
                                                    _ = v91
                                                }
                                            }
                                            v5 = ""
                                        } else {
                                            v56 += (v56 & (v56 | 970))
                                            delete(v1, 2)
                                            delete(v4, 1)
                                        }
                                    }
                                    switch (fi47 - fi47) {
                                    case 868:
                                        v92 := GenS11{F0: 669, F1: v48}
                                        _ = v92
                                    case -9223372036854775808:
                                        v93 := GenS12{F0: -128, F1: 642, F2: v56}
                                        _ = v93
                                    }
                                    switch ("foo bar" + v5) {
                                    case "":
                                        if ((0.1 - -2.25) <= 1.5) {
                                            v94 := []int{(fi47 & (fi47 % ((128) | 1)))}
                                            v95 := GenS13{F0: (v5 + v5), F1: v56, F2: (^(v56))}
                                            var v96 uint8 = 90
                                            _ = v94
                                            _ = v95
                                            _ = v96
                                        }
                                    default:
                                        v97 := GenS6{F0: 2}
                                        _ = v97
                                    }
                                    v98 := []int{0, 9223372036854775807}
                                    var v99 int64 = v56
                                    _ = v98
                                    _ = v99
                                } else {
                                    fmt.Println(v5)
                                    v44 += (v44 * 4294967295)
                                }
                                var v100 int16 = -1
                                _ = v100
                            }
                            _ = v70
                        } else {
                            if ((65535 > 2) || ((65535 == 65535) && false)) {
                                v101 := GenS14{F0: (^(v48)), F1: false}
                                v1[314] = true
                                v48 -= v48
                                _ = v101
                            }
                            v2[v5] = (32767 != 32767)
                        }
                        _ = v56
                    }
                }
                _ = fi47
                _ = v48
            }
        }
        v3 += v3
        _ = fi45
        _ = v46
    }
    switch (^(v3)) {
    case 1:
        v1[2147483648] = false
    case -896:
        fmt.Println(v3, v44)
    case -273:
        for fi102 := 0; fi102 < 20; fi102++ {
            var v103 uint = 42
            fmt.Println(fi102, v3)
            _ = fi102
            _ = v103
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
}
