package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 uint; F1 uint32; F2 int64 }
type GenS5 struct { F0 uint16; F1 int16; F2 uint8 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint16; F1 int8 }
type GenS8 struct { F0 int; F1 uint }
type GenS9 struct { F0 bool; F1 uint8; F2 uint8 }
type GenS10 struct { F0 string; F1 string; F2 bool }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 bool; F1 int32; F2 bool }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 int; F1 uint }
type GenS16 struct { F0 int8 }
func genFn0(p0 uint16) int64 {
    if false {
        v7 := GenS1{F0: -316}
        if true {
            for fi8 := 0; fi8 < 18; fi8++ {
                v9 := []int{(fi8 | -2147483648), 4294967295, (fi8 / (((fi8 * 32768)) | 1)), 128}
                var v10 uint64 = 824
                v10 = v10
                _ = fi8
                _ = v10
                _ = v9
            }
            var v11 uint = 65535
            if (true && (true || (0 == 125))) {
                v12 := []int{128, 42}
                if false {
                    for fi13 := 0; fi13 < 7; fi13++ {
                        v14 := GenS1{F0: -460}
                        if len(v12) >= 2 {
                            v12 = v12[1:len(v12)]
                        }
                        _ = fi13
                        _ = v14
                    }
                    fmt.Println(p0, v11)
                } else {
                    var v15 uint8 = 57
                    _ = v15
                }
                if false {
                    var v16 uint32 = 615
                    switch 9223372036854775807 {
                    case 127:
                        v17 := GenS2{F0: -779, F1: 42}
                        _ = v17
                    case 4294967295:
                        fmt.Println(p0, p0)
                    case -93:
                        for fi18 := 0; fi18 < 12; fi18++ {
                            fmt.Println(v16)
                            _ = fi18
                        }
                    }
                    _ = v16
                } else {
                    var v19 int8 = 0
                    _ = len(v12)
                    v20 := GenS2{F0: -514, F1: 42}
                    _ = v19
                    _ = v20
                }
                _ = v12
            }
            _ = v11
        }
        _ = v7
    } else {
        p0 = p0
        p0 -= (^(p0))
        v21 := GenS3{F0: 915, F1: 42}
        _ = v21
    }
    for fi22 := 0; fi22 < 6; fi22++ {
        var v23 float32 = -2.25
        v24 := GenS4{F0: 42, F1: 1, F2: -437}
        v25 := []int{}
        _ = fi22
        _ = v23
        _ = v24
        _ = v25
    }
    switch "\n\t" {
    case "foo bar":
        if true {
            if false {
                v26 := GenS5{F0: 357, F1: -739, F2: 1}
                _ = v26
            }
            fmt.Println(p0)
        }
    case "":
        for fi27 := 0; fi27 < 2; fi27++ {
            for fi28 := 0; fi28 < 20; fi28++ {
                if true {
                    v29 := []int{fi27, (fi27 << ((0) & 63)), fi28, fi28}
                    p0 += (p0 >> ((18446744073709551615) & 63))
                    var v30 float32 = (0.1 * ((1.5 * -2.25) + float32(1e10)))
                    _ = v29
                    _ = v30
                }
                p0 += (p0 - (p0 - p0))
                _ = fi28
            }
            v31 := GenS6{F0: 33}
            if true {
                v32 := []int{(fi27 % (((^(fi27))) | 1)), (^(fi27)), (fi27 / ((fi27) | 1)), (fi27 * (-(fi27)))}
                for fi33 := 0; fi33 < 4; fi33++ {
                    v34 := GenS6{F0: 42}
                    p0 = p0
                    _ = fi33
                    _ = v34
                }
                var v35 uint = 1
                _ = v32
                _ = v35
            } else {
                var v36 uint8 = uint8(p0)
                var v37 float64 = 3.14159
                var v38 uint16 = 42
                _ = v36
                _ = v37
                _ = v38
            }
            _ = fi27
            _ = v31
        }
    }
    var v39 string = "hello"
    for fi40 := 0; fi40 < 16; fi40++ {
        p0 += 42
        _ = fi40
    }
    _ = p0
    _ = v39
    return -557
}
func genFn1(p0 string) int32 {
    var v41 bool = true
    for fi42 := 0; fi42 < 3; fi42++ {
        switch (fi42 - 32767) {
        case 4294967295:
            var v43 uint64 = 0
            _ = v43
        case 2147483647:
            var v44 int32 = -2147483648
            _ = v44
        }
        var v45 uint8 = 2
        for fi46 := 0; fi46 < 17; fi46++ {
            v41 = v41
            var v47 float32 = (3.14159 - 1.5)
            v48 := []int{(fi42 - (fi46 + fi46)), (fi42 ^ -9223372036854775808), fi42}
            _ = fi46
            _ = v47
            _ = v48
        }
        _ = fi42
        _ = v45
    }
    v41 = true
    _ = p0
    _ = v41
    return 0
}

func main() {
    var v1 int64 = 222
    v2 := []int{127, -9223372036854775808, 127, 740}
    v3 := []uint32{568, 1}
    var v4 bool = (0 == 1)
    var v5 float32 = float32((1e10 / -2.25))
    v6 := []int{9223372036854775807}
    v49 := GenS7{F0: 0, F1: -128}
    v50 := GenS4{F0: 4294967295, F1: 2, F2: (v1 ^ (^(v1)))}
    switch 9223372036854775807 {
    case 434:
        var v51 int = -1
        _ = v51
    }
    if (-849 >= 1) {
        fmt.Println(v1, v5, v1)
        var v52 int8 = -128
        _ = v52
    } else {
        v53 := GenS8{F0: -563, F1: 44}
        for fi54 := 0; fi54 < 7; fi54++ {
            var v55 bool = v4
            _ = fi54
            _ = v55
        }
        v5 += 3.14159
        _ = v53
    }
    if v4 {
        v56 := GenS9{F0: ((v5 * 0.0) < ((1e10 + (v5 / 1e10)) * v5)), F1: 127, F2: 127}
        _ = v56
    }
    v2 = append(v2, 0)
    var v57 float32 = v5
    v5 += v57
    v58 := GenS3{F0: 1, F1: 117}
    fmt.Println(v57)
    if v4 {
        var v59 int32 = -2147483648
        if (157 <= 2) {
            switch -557 {
            case 32768:
                fmt.Println(v57)
            case 42:
                fmt.Println(v5)
            }
        } else {
            switch 665 {
            case 127:
                switch "foo bar" {
                case "":
                    for fi60 := 0; fi60 < 1; fi60++ {
                        v5 = 0.1
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        for fi61 := 0; fi61 < 7; fi61++ {
                            _ = len(v3)
                            switch (fi60 / (((^(fi60))) | 1)) {
                            case -2147483648:
                                if v4 {
                                    switch (fi60 - (fi61 / ((fi60) | 1))) {
                                    case 9223372036854775807:
                                        v62 := GenS10{F0: ("a" + ("\n\t" + "a")), F1: ("a" + ("hello" + "\n\t")), F2: (42 != 42)}
                                        _ = v62
                                    case 549:
                                        if true {
                                            v63 := GenS11{F0: 65535}
                                            switch "\n\t" {
                                            case "hello":
                                                if (!false) {
                                                    v57 = (float32(0.0) / (v57 + 1e10))
                                                    fmt.Println(v57, v5, v5)
                                                }
                                            }
                                            v64 := GenS5{F0: 65535, F1: 582, F2: 186}
                                            _ = v63
                                            _ = v64
                                        }
                                    default:
                                        if (!(!v4)) {
                                            var v65 int16 = 32767
                                            _ = v65
                                        } else {
                                            v1 -= genFn0(255)
                                            v66 := GenS6{F0: int32(v1)}
                                            _ = v66
                                        }
                                    }
                                    fmt.Println(v59, v1, v57)
                                    fmt.Println(v57, v5)
                                    v4 = (false && v4)
                                    var v67 float32 = v5
                                    _ = len(v2)
                                    _ = v67
                                }
                            default:
                                var v68 uint32 = 65535
                                var v69 int = fi60
                                var v70 float32 = v5
                                _ = v68
                                _ = v69
                                _ = v70
                            }
                            v1 = v1
                            _ = fi61
                        }
                        _ = fi60
                    }
                case "a":
                    v1 += 11
                default:
                    switch -9223372036854775808 {
                    case -9223372036854775808:
                        var v71 bool = (((1.5 / 1.5) > -2.25) && false)
                        _ = v71
                    case -750:
                        v72 := GenS8{F0: 32767, F1: 70}
                        _ = v72
                    case 9223372036854775807:
                        switch 9223372036854775807 {
                        case -87:
                            v4 = false
                        case 670:
                            v1 = -9223372036854775808
                        case 2147483647:
                            v4 = false
                        default:
                            fmt.Println(v57)
                        }
                    }
                }
            default:
                v6 = append(v6, 942)
            }
            if (4294967295 >= 42) {
                v4 = v4
            } else {
                v4 = (0 != 18446744073709551615)
            }
            var v73 uint16 = 2
            _ = v73
        }
        v74 := GenS12{F0: v4, F1: (v59 >> ((1) & 63)), F2: true}
        _ = v59
        _ = v74
    } else {
        fmt.Println(v5)
        _ = len(v2)
        v4 = v4
    }
    var v75 string = ""
    switch "\n\t" {
    case "a":
        var v76 float32 = v57
        _ = v76
    case "foo bar":
        if (216 <= 192) {
            _ = len(v3)
            v77 := GenS3{F0: 4294967295, F1: 4294967295}
            _ = v77
        } else {
            for fi78 := 0; fi78 < 10; fi78++ {
                if ((v1 >> ((65535) & 63)) > v1) {
                    for fi79 := 0; fi79 < 13; fi79++ {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v80 int8 = 90
                        _ = len(v3)
                        v75 = ""
                        _ = fi79
                        _ = v80
                    }
                    fmt.Println(fi78, fi78)
                } else {
                    var v81 int32 = 42
                    for fi82 := 0; fi82 < 4; fi82++ {
                        v83 := GenS13{F0: 1}
                        fmt.Println(fi82)
                        _ = fi82
                        _ = v83
                    }
                    _ = v81
                }
                switch fi78 {
                case 1:
                    fmt.Println(v1, v57, v5)
                case -1:
                    _ = len(v3)
                default:
                    v84 := GenS14{F0: -1}
                    _ = v84
                }
                _ = fi78
            }
            var v85 uint32 = 2147483647
            fmt.Println(v4)
            _ = v85
        }
    case "\n\t":
        var v86 float32 = (v57 * v57)
        _ = v86
    default:
        v1 -= genFn0(351)
    }
    for fi87 := 0; fi87 < 8; fi87++ {
        for fi88 := 0; fi88 < 3; fi88++ {
            fmt.Println(v57, v5, v1)
            if false {
                var v89 uint64 = 42
                v1 = -738
                _ = v89
            }
            if (42 < -679) {
                if (!(1 == 42)) {
                    _ = len(v2)
                } else {
                    switch v75 {
                    case "\n\t":
                        v90 := GenS11{F0: 661}
                        _ = v90
                    default:
                        v91 := GenS15{F0: (-(fi88)), F1: 65535}
                        _ = v91
                    }
                    _ = len(v3)
                    v2 = append(v2, (^(fi87)))
                }
                v57 += (v5 / float32(1.5))
                fmt.Println(fi88, v75)
            }
            _ = len(v6)
            switch (fi88 % ((fi87) | 1)) {
            case 128:
                switch (v75 + v75) {
                case "foo bar":
                    var v92 uint8 = 1
                    _ = v92
                case "":
                    var v93 float32 = (3.14159 + v5)
                    v4 = false
                    v94 := GenS10{F0: ("\n\t" + v75), F1: "hello", F2: true}
                    _ = v93
                    _ = v94
                default:
                    var v95 bool = v4
                    switch 42 {
                    case 4294967295:
                        v75 = (v75 + v75)
                    case -2147483648:
                        if (!(!v4)) {
                            var v96 int16 = 501
                            _ = v96
                        } else {
                            fmt.Println(v5, v5)
                            var v97 string = v75
                            if false {
                                if v4 {
                                    fmt.Println(v75, v1, v95)
                                    v57 -= (v5 * -2.25)
                                    var v98 int32 = -2147483648
                                    var v99 uint8 = 128
                                    _ = v98
                                    _ = v99
                                } else {
                                    v5 = 1e10
                                    if (v95 && (v1 >= (-(v1)))) {
                                        var v100 int64 = (^(v1))
                                        var v101 int64 = 9223372036854775807
                                        _ = v100
                                        _ = v101
                                    }
                                    var v102 float32 = ((((v57 - float32((((((((((1e10 * float64(((v5 + v5) - (float32(0.1) - v5)))) * 0.1) + (0.0 / 1.0)) / 1e10) * 0.0) / 1.0) - 0.0) + 1.5) - -2.25))) / v57) + 1.5) - ((v5 - (1e10 + v57)) - (0.1 * (1.5 * 0.0))))
                                    _ = v102
                                }
                                v5 = v5
                            }
                            _ = v97
                        }
                    default:
                        v75 = ((v75 + "a") + "\n\t")
                    }
                    v1 = v1
                    _ = v95
                }
            case -1:
                var v103 int = fi87
                _ = v103
            default:
                v4 = (-80 == 42)
            }
            _ = fi88
        }
        v104 := GenS7{F0: 42, F1: 63}
        v105 := GenS16{F0: 127}
        _ = fi87
        _ = v104
        _ = v105
    }
    switch 32768 {
    case 9223372036854775807:
        if len(v6) > 0 {
            v6[0] = 32768
        }
    case -274:
        fmt.Println(v4, v57)
    case -606:
        var v106 uint = 0
        _ = v106
    }
    v75 = v75
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v57
    _ = v58
    _ = v6
    _ = v75
}
