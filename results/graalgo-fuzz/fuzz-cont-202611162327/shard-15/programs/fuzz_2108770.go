package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 uint16; F1 int64 }
type GenS5 struct { F0 uint; F1 uint }
type GenS6 struct { F0 int; F1 uint8 }
type GenS7 struct { F0 int16; F1 bool }
type GenS8 struct { F0 bool; F1 uint8 }
type GenS9 struct { F0 int32; F1 string }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 int }
type GenS15 struct { F0 int16; F1 int16; F2 string }
func genFn0(p0 int8, p1 uint16) uint8 {
    v5 := GenS1{F0: (!(((0.1 + float32(0.0)) * 1.5) == ((1.5 / (((0.0 - -2.25) * 1.5) / 1.0)) - 1e10)))}
    switch 128 {
    case 588:
        v6 := []int{-1}
        _ = v6
    default:
        p0 -= (p0 - 0)
    }
    var v7 uint32 = 0
    _ = p0
    _ = p1
    _ = v5
    _ = v7
    return 1
}

func main() {
    var v1 int8 = 42
    var v2 float64 = 0.0
    v3 := make(map[string]int)
    var v4 uint64 = 964
    switch 32767 {
    case -568:
        v8 := GenS1{F0: false}
        _ = v8
    default:
        v9 := []int{}
        _ = v9
    }
    var v10 uint = 18446744073709551615
    var v11 string = "\n\t"
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 uint64 = (v4 - uint64(v10))
        _ = fi12
        _ = v13
    }
    switch v11 {
    case "\n\t":
        v2 = v2
    case "a":
        if (false || ((v1 / ((v1) | 1)) != (v1 - (^(v1))))) {
            var v14 uint8 = 42
            _ = v14
        } else {
            var v15 bool = true
            for fi16 := 0; fi16 < 2; fi16++ {
                v17 := GenS2{F0: (^(v10)), F1: 362}
                v18 := GenS2{F0: (v10 * v10), F1: uint32(v4)}
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = v15
        }
    case "foo bar":
        for fi19 := 0; fi19 < 2; fi19++ {
            var v20 int64 = 9223372036854775807
            switch (v11 + (v11 + v11)) {
            case "hello":
                fmt.Println(v11, v10)
            default:
                if true {
                    if (788 == 2147483648) {
                        v21 := v3[v11]
                        _ = v21
                    }
                    v22 := GenS1{F0: false}
                    _ = v22
                } else {
                    v2 = v2
                }
            }
            for fi23 := 0; fi23 < 5; fi23++ {
                fmt.Println(v2)
                v24 := GenS1{F0: false}
                if true {
                    switch (fi23 / ((fi23) | 1)) {
                    case 32767:
                        switch v11 {
                        case "a":
                            switch fi19 {
                            case 127:
                                v2 += v2
                            case -754:
                                v25 := GenS3{F0: 32767, F1: (v11 + v11)}
                                _ = v25
                            default:
                                switch fi19 {
                                case 32767:
                                    v26 := v3[v11]
                                    var v27 int = fi19
                                    _ = v26
                                    _ = v27
                                default:
                                    var v28 uint32 = 103
                                    _ = len(v3)
                                    _ = v28
                                }
                            }
                        }
                    case 128:
                        var v29 int64 = (v20 ^ (v20 + 0))
                        v11 = v11
                        v1 = v1
                        _ = v29
                    }
                    var v30 string = v11
                    if true {
                        var v31 int = (fi19 - fi19)
                        _ = v31
                    } else {
                        var v32 uint = (v10 << ((v10) & 63))
                        var v33 string = (v11 + (v11 + v11))
                        v20 = v20
                        fmt.Println(fi23, v2, v2)
                        v11 = v30
                        v10 = v10
                        _ = v32
                        _ = v33
                    }
                    _ = v30
                } else {
                    v34 := GenS4{F0: 2, F1: (^(v20))}
                    fmt.Println(fi23)
                    var v35 float32 = (-2.25 - 0.0)
                    v20 = v20
                    _ = v34
                    _ = v35
                }
                _ = fi23
                _ = v24
            }
            _ = fi19
            _ = v20
        }
    default:
        for fi36 := 0; fi36 < 7; fi36++ {
            v37 := v3[v11]
            var v38 uint64 = v4
            var v39 uint8 = 251
            _ = fi36
            _ = v37
            _ = v38
            _ = v39
        }
    }
    switch -9223372036854775808 {
    case 32768:
        v40 := []int{32767, -2147483648, 0, int(v4)}
        _ = v40
    case 1:
        v11 = "foo bar"
    }
    v41 := GenS5{F0: 1, F1: v10}
    for fi42 := 0; fi42 < 17; fi42++ {
        _ = len(v3)
        _ = fi42
    }
    for fi43 := 0; fi43 < 1; fi43++ {
        delete(v3, "\n\t")
        fmt.Println(fi43)
        _ = fi43
    }
    if true {
        if ((3.14159 - 1.5) != (v2 - v2)) {
            v44 := GenS3{F0: 0, F1: v11}
            _ = v44
        }
        if false {
            for fi45 := 0; fi45 < 8; fi45++ {
                var v46 float64 = 3.14159
                if (!(57 < 128)) {
                    switch fi45 {
                    case 1:
                        for fi47 := 0; fi47 < 20; fi47++ {
                            v48 := GenS5{F0: v10, F1: v10}
                            _ = fi47
                            _ = v48
                        }
                    case -527:
                        fmt.Println(v11, v1, v1)
                    default:
                        for fi49 := 0; fi49 < 4; fi49++ {
                            fmt.Println(v46, v46, fi49)
                            if (42 <= v4) {
                                if false {
                                    delete(v3, v11)
                                    switch v11 {
                                    case "a":
                                        v4 -= 878
                                        var v50 int16 = -219
                                        v51 := []int{262}
                                        _ = v50
                                        _ = v51
                                    case "\n\t":
                                        delete(v3, "\n\t")
                                        if false {
                                            v11 = v11
                                            v10 += (v10 << (((v10 * (v10 - v10))) & 63))
                                            switch -1 {
                                            case 9223372036854775807:
                                                v2 += float64(-2.25)
                                                fmt.Println(v4)
                                                if true {
                                                    var v52 bool = ((fi49 >> ((v10) & 63)) < (-(fi45)))
                                                    _ = v52
                                                }
                                                v3[v11] = (^(fi49))
                                                v4 = v4
                                                v53 := []int{9223372036854775807, 32768}
                                                var v54 bool = true
                                                _ = v53
                                                _ = v54
                                            case -9223372036854775808:
                                                v55 := []int{(fi45 >> (((^(v10))) & 63)), (^(fi45))}
                                                _ = v55
                                            }
                                            fmt.Println(v11)
                                            if false {
                                                v4 = (v4 | v4)
                                            } else {
                                                v56 := v3["foo bar"]
                                                _ = v56
                                            }
                                            if (!true) {
                                                if true {
                                                    v11 = v11
                                                    if (!true) {
                                                        if true {
                                                            v57 := []int{4294967295, -9223372036854775808, fi49}
                                                            var v58 int16 = 1000
                                                            var v59 string = "hello"
                                                            if false {
                                                                fmt.Println(v11, fi49, v1)
                                                            } else {
                                                                v2 = v46
                                                            }
                                                            if false {
                                                                fmt.Println(v10, v4)
                                                                _ = len(v3)
                                                            } else {
                                                                _ = len(v57)
                                                            }
                                                            var v60 uint8 = 42
                                                            _ = v57
                                                            _ = v58
                                                            _ = v59
                                                            _ = v60
                                                        } else {
                                                            v10 += v10
                                                            v61 := GenS6{F0: fi45, F1: 255}
                                                            if (false || (!(v1 > int8(fi45)))) {
                                                                v62 := []int{9223372036854775807, fi45, (fi49 >> (((v10 ^ (v10 ^ v10))) & 63)), (-(fi49))}
                                                                switch (fi45 / ((32768) | 1)) {
                                                                case -831:
                                                                    if (!true) {
                                                                        v11 = v11
                                                                        if (false && (4294967295 >= (fi45 >> ((0) & 63)))) {
                                                                            if ((!true) || (v2 < ((v46 * float64(0.0)) / (((3.14159 - 1.5) + (v2 / (0.1 / ((((0.1 + v46) + (v2 * ((v46 * v46) / v2))) * (v2 * (v2 + -2.25))) * ((v2 / v46) / float64(3.14159)))))) / v2)))) {
                                                                                if (v1 != v1) {
                                                                                    var v63 uint32 = 97
                                                                                    v64 := GenS5{F0: 42, F1: (^(v10))}
                                                                                    _ = v63
                                                                                    _ = v64
                                                                                }
                                                                            } else {
                                                                                var v65 uint64 = (^(v4))
                                                                                v66 := GenS7{F0: -434, F1: false}
                                                                                v65 = v65
                                                                                _ = v65
                                                                                _ = v66
                                                                            }
                                                                            if (!false) {
                                                                                var v67 uint16 = 380
                                                                                v4 = (v4 - v4)
                                                                                _ = v67
                                                                            }
                                                                        }
                                                                        fmt.Println(v10)
                                                                    }
                                                                case 308:
                                                                    v4 = 237
                                                                case -427:
                                                                    v68 := GenS8{F0: true, F1: uint8(v4)}
                                                                    _ = v68
                                                                }
                                                                _ = v62
                                                            } else {
                                                                var v69 float64 = v46
                                                                _ = v69
                                                            }
                                                            v10 = (^(v10))
                                                            _ = v61
                                                        }
                                                        var v70 string = v11
                                                        _ = v70
                                                    } else {
                                                        v71 := GenS9{F0: 13, F1: v11}
                                                        _ = v71
                                                    }
                                                    var v72 uint8 = 42
                                                    var v73 uint64 = 42
                                                    fmt.Println(v10, v1)
                                                    _ = v72
                                                    _ = v73
                                                }
                                                fmt.Println(fi49)
                                            } else {
                                                var v74 int8 = v1
                                                var v75 int16 = 42
                                                _ = v74
                                                _ = v75
                                            }
                                        } else {
                                            switch fi49 {
                                            case 32768:
                                                switch fi49 {
                                                case -116:
                                                    v10 = (v10 >> ((v10) & 63))
                                                case 4294967295:
                                                    v1 += v1
                                                }
                                            case 42:
                                                v2 -= -2.25
                                            case 4294967295:
                                                v76 := GenS4{F0: 782, F1: 1}
                                                _ = v76
                                            }
                                            v77 := []int{(fi45 * fi49), int(v10), fi49, fi49}
                                            v11 = (v11 + v11)
                                            _ = v77
                                        }
                                    case "foo bar":
                                        v78 := GenS10{F0: 255}
                                        _ = v78
                                    default:
                                        if true {
                                            v4 = v4
                                        }
                                    }
                                } else {
                                    var v79 float32 = 0.1
                                    if true {
                                        v1 = (v1 ^ (v1 % ((int8(v4)) | 1)))
                                        fmt.Println(v10, v46)
                                        switch -109 {
                                        case 9223372036854775807:
                                            if ((v10 >> ((v10) & 63)) > (^(v10))) {
                                                if (!(v1 != v1)) {
                                                    delete(v3, v11)
                                                    switch "" {
                                                    case "a":
                                                        fmt.Println(v79, v4, v79)
                                                    default:
                                                        if (42 > (fi49 & fi49)) {
                                                            v10 -= v10
                                                            fmt.Println(v46, v1, v2)
                                                            v4 = 181
                                                            switch v11 {
                                                            case "hello":
                                                                fmt.Println(v4, v10, v10)
                                                            case "":
                                                                if (1000 > 166) {
                                                                    switch "hello" {
                                                                    case "a":
                                                                        v80 := GenS3{F0: 32767, F1: "\n\t"}
                                                                        _ = v80
                                                                    case "foo bar":
                                                                        if (!(1 != 30)) {
                                                                            v79 -= v79
                                                                            v46 = v46
                                                                        }
                                                                    default:
                                                                        v81 := GenS5{F0: (v10 ^ 18446744073709551615), F1: (v10 << (((^(v10))) & 63))}
                                                                        _ = v81
                                                                    }
                                                                    v82 := GenS2{F0: 2, F1: 928}
                                                                    _ = v82
                                                                }
                                                            case "foo bar":
                                                                var v83 uint16 = 1
                                                                _ = v83
                                                            }
                                                            v84 := []int{fi45, fi49, fi45, -679}
                                                            _ = v84
                                                        }
                                                    }
                                                    var v85 bool = false
                                                    _ = v85
                                                }
                                                var v86 uint16 = 65535
                                                var v87 string = ((v11 + v11) + v11)
                                                _ = v86
                                                _ = v87
                                            }
                                        }
                                    }
                                    var v88 float64 = (((float64(((-2.25 - v79) - (3.14159 * v79))) / v2) / (v2 * float64(v79))) - 0.1)
                                    _ = v79
                                    _ = v88
                                }
                                v89 := GenS5{F0: 65535, F1: v10}
                                v90 := GenS11{F0: v4}
                                switch fi45 {
                                case 1:
                                    v11 = (v11 + "hello")
                                case 32767:
                                    var v91 int16 = -32768
                                    if (true || true) {
                                        var v92 float64 = v46
                                        v3["foo bar"] = 4294967295
                                        if true {
                                            if false {
                                                var v93 int16 = (v91 & v91)
                                                _ = v93
                                            } else {
                                                v94 := GenS12{F0: v4}
                                                v95 := []int{(fi45 / ((32768) | 1)), (fi49 + -532)}
                                                if ((0.1 + float64((1.5 + 0.0))) > float64(1.5)) {
                                                    if len(v95) >= 2 {
                                                        v95 = v95[1:len(v95)]
                                                    }
                                                    var v96 bool = false
                                                    _ = v96
                                                }
                                                _ = v94
                                                _ = v95
                                            }
                                            v11 = "hello"
                                            var v97 uint = (v10 >> ((v10) & 63))
                                            if (false || (0 == 0)) {
                                                var v98 int32 = -413
                                                var v99 bool = true
                                                _ = v98
                                                _ = v99
                                            } else {
                                                var v100 int64 = 1
                                                switch ("a" + v11) {
                                                case "\n\t":
                                                    switch fi45 {
                                                    case -2147483648:
                                                        v101 := []int{(fi45 << ((4294967295) & 63))}
                                                        _ = v101
                                                    case -439:
                                                        var v102 uint64 = 1
                                                        _ = v102
                                                    case -107:
                                                        v103 := GenS13{F0: -362}
                                                        _ = v103
                                                    }
                                                case "hello":
                                                    var v104 string = "hello"
                                                    _ = v104
                                                default:
                                                    var v105 string = v11
                                                    _ = v105
                                                }
                                                var v106 int8 = (v1 << (((v10 ^ (v10 * (v10 & v97)))) & 63))
                                                _ = v100
                                                _ = v106
                                            }
                                            v107 := []int{fi45, fi45}
                                            if len(v107) >= 2 {
                                                v107 = v107[1:len(v107)]
                                            }
                                            _ = v107
                                            _ = v97
                                        }
                                        _ = v92
                                    }
                                    v4 -= v4
                                    _ = v91
                                }
                                _ = v89
                                _ = v90
                            } else {
                                var v108 int8 = -3
                                _ = v108
                            }
                            fmt.Println(v46)
                            _ = fi49
                        }
                    }
                } else {
                    v109 := GenS1{F0: true}
                    _ = v109
                }
                v110 := GenS14{F0: (fi45 % ((-2147483648) | 1))}
                _ = fi45
                _ = v110
                _ = v46
            }
        }
        var v111 uint64 = v4
        _ = v111
    } else {
        for fi112 := 0; fi112 < 10; fi112++ {
            delete(v3, v11)
            switch v11 {
            case "\n\t":
                for fi113 := 0; fi113 < 20; fi113++ {
                    v10 = (v10 - 84)
                    v1 = 0
                    _ = fi113
                }
            case "a":
                var v114 int8 = (-(v1))
                _ = v114
            case "foo bar":
                delete(v3, (v11 + "foo bar"))
            default:
                v4 = v4
            }
            _ = fi112
        }
        for fi115 := 0; fi115 < 15; fi115++ {
            v3["hello"] = (fi115 * fi115)
            v116 := []int{}
            for fi117 := 0; fi117 < 12; fi117++ {
                fmt.Println(v10)
                _ = fi117
            }
            _ = fi115
            _ = v116
        }
    }
    for fi118 := 0; fi118 < 6; fi118++ {
        if ((v11 + v11) >= v11) {
            var v119 int8 = v1
            fmt.Println(fi118, v119, v4)
            var v120 uint = (v10 | v10)
            _ = v119
            _ = v120
        } else {
            switch "foo bar" {
            case "":
                if (!(!false)) {
                    v10 -= (v10 << (((v10 * 553)) & 63))
                    v2 += v2
                } else {
                    v121 := []int{fi118}
                    var v122 float64 = (0.0 / v2)
                    _ = v121
                    _ = v122
                }
            case "a":
                var v123 uint8 = 1
                _ = v123
            }
        }
        v124 := []int{fi118, (-(fi118)), fi118}
        if true {
            v125 := GenS7{F0: 1000, F1: true}
            _ = v125
        } else {
            var v126 int16 = -1
            v126 -= (^(v126))
            v127 := GenS15{F0: v126, F1: v126, F2: v11}
            _ = v126
            _ = v127
        }
        _ = fi118
        _ = v124
    }
    v11 = v11
    v128 := GenS11{F0: v4}
    for fi129 := 0; fi129 < 2; fi129++ {
        v11 = v11
        fmt.Println(v11, v10)
        _ = fi129
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v128
    _ = v2
    _ = v3
    _ = v4
    _ = v41
}
