package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 int8; F2 uint32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32; F1 uint16; F2 int32 }
type GenS6 struct { F0 uint32; F1 uint }
type GenS7 struct { F0 int8; F1 string }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint8; F1 int64; F2 int16 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 int16; F1 uint; F2 int32 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 int8; F1 uint64 }
type GenS14 struct { F0 int32; F1 bool }
type GenS15 struct { F0 int64 }
func genFn0() string {
    switch ("foo bar" + "a") {
    case "foo bar":
        switch "hello" {
        case "":
            var v4 int = -9223372036854775808
            _ = v4
        case "hello":
            var v5 int64 = 9223372036854775807
            _ = v5
        default:
            var v6 uint = 519
            _ = v6
        }
    case "":
        switch "\n\t" {
        case "hello":
            v7 := []int{32768, -2147483648, 128}
            _ = v7
        case "\n\t":
            v8 := GenS1{F0: 2, F1: -32768}
            _ = v8
        default:
            switch -9223372036854775808 {
            case 32767:
                var v9 int8 = 6
                _ = v9
            }
        }
    case "hello":
        for fi10 := 0; fi10 < 13; fi10++ {
            fmt.Println(fi10)
            var v11 float64 = ((-2.25 / 1e10) / -2.25)
            for fi12 := 0; fi12 < 14; fi12++ {
                if true {
                    var v13 int8 = 1
                    if false {
                        v13 -= (v13 << ((42) & 63))
                    } else {
                        v11 -= v11
                        v14 := GenS1{F0: 1, F1: 42}
                        fmt.Println(v11, fi10)
                        _ = v14
                    }
                    var v15 uint32 = 4294967295
                    var v16 uint64 = 0
                    v17 := GenS1{F0: 168, F1: 42}
                    _ = v13
                    _ = v15
                    _ = v16
                    _ = v17
                }
                if (v11 < v11) {
                    switch fi12 {
                    case 127:
                        if (((-32768 < 1) && true) && (false && (!false))) {
                            v18 := GenS1{F0: 65535, F1: -1}
                            fmt.Println(fi12, v11, fi10)
                            var v19 uint64 = 4294967295
                            var v20 bool = (517 != int16(fi10))
                            if false {
                                var v21 uint8 = 2
                                _ = v21
                            }
                            _ = v18
                            _ = v19
                            _ = v20
                        }
                    case 32768:
                        v22 := GenS1{F0: 65535, F1: -418}
                        _ = v22
                    }
                }
                var v23 float32 = (-2.25 - float32(v11))
                _ = fi12
                _ = v23
            }
            _ = fi10
            _ = v11
        }
    default:
        var v24 int = 42
        _ = v24
    }
    var v25 float64 = 1.5
    switch ("a" + "") {
    case "foo bar":
        var v26 uint64 = 125
        _ = v26
    case "hello":
        v27 := GenS1{F0: 2, F1: 558}
        _ = v27
    default:
        for fi28 := 0; fi28 < 5; fi28++ {
            v25 = ((1.5 + v25) * (3.14159 * float64(float32(0.0))))
            _ = fi28
        }
    }
    if (11 <= 42) {
        var v29 uint32 = 0
        var v30 uint64 = 0
        _ = v29
        _ = v30
    } else {
        for fi31 := 0; fi31 < 8; fi31++ {
            if true {
                switch -1 {
                case -669:
                    v25 = float64(3.14159)
                case 42:
                    v32 := GenS2{F0: 0}
                    _ = v32
                case 512:
                    switch (-(fi31)) {
                    case 0:
                        var v33 int8 = -6
                        _ = v33
                    case 2147483647:
                        for fi34 := 0; fi34 < 14; fi34++ {
                            var v35 int16 = 0
                            v25 += 3.14159
                            _ = fi34
                            _ = v35
                        }
                    case 1:
                        v36 := []int{(fi31 << ((4294967295) & 63)), 42}
                        _ = v36
                    default:
                        v37 := []int{}
                        _ = v37
                    }
                default:
                    if true {
                        v38 := GenS3{F0: 0, F1: -1, F2: 0}
                        _ = v38
                    } else {
                        v39 := GenS1{F0: 18446744073709551615, F1: 32767}
                        _ = v39
                    }
                }
                v40 := GenS4{F0: (!(310 == 87))}
                _ = v40
            } else {
                var v41 float64 = float64(0.1)
                _ = v41
            }
            v42 := []int{fi31, 103, fi31}
            for fi43 := 0; fi43 < 14; fi43++ {
                if true {
                    v44 := GenS5{F0: 931, F1: 2, F2: -1}
                    _ = v44
                }
                if len(v42) > 0 {
                    v42[0] = -713
                }
                v45 := GenS6{F0: 2, F1: 4294967295}
                _ = fi43
                _ = v45
            }
            _ = fi31
            _ = v42
        }
        v46 := []int{128, -9223372036854775808, -1, -2147483648}
        switch "hello" {
        case "\n\t":
            if len(v46) >= 2 {
                v46 = v46[1:len(v46)]
            }
        default:
            v25 -= 0.1
        }
        _ = v46
    }
    if false {
        v25 -= v25
        for fi47 := 0; fi47 < 9; fi47++ {
            if true {
                v25 = (0.0 / (v25 - v25))
                var v48 string = ("foo bar" + "\n\t")
                switch v48 {
                case "hello":
                    fmt.Println(v48)
                default:
                    v25 = (v25 * (v25 / (0.1 + v25)))
                }
                _ = v48
            } else {
                if (!true) {
                    if true {
                        for fi49 := 0; fi49 < 2; fi49++ {
                            var v50 int8 = -128
                            var v51 float64 = v25
                            _ = fi49
                            _ = v50
                            _ = v51
                        }
                        v25 = 0.0
                        v25 = -2.25
                    } else {
                        for fi52 := 0; fi52 < 4; fi52++ {
                            v53 := GenS7{F0: 42, F1: "\n\t"}
                            _ = fi52
                            _ = v53
                        }
                    }
                    v54 := GenS4{F0: true}
                    _ = v54
                } else {
                    v55 := GenS7{F0: -1, F1: ""}
                    _ = v55
                }
                var v56 float64 = v25
                _ = v56
            }
            v57 := GenS8{F0: 95}
            _ = fi47
            _ = v57
        }
    } else {
        v25 = v25
        for fi58 := 0; fi58 < 1; fi58++ {
            for fi59 := 0; fi59 < 9; fi59++ {
                var v60 uint64 = 1
                _ = fi59
                _ = v60
            }
            _ = fi58
        }
    }
    _ = v25
    return ("" + ("a" + ""))
}
func genFn1(p0 string, p1 string, p2 int32) uint {
    var v61 uint64 = 581
    if false {
        var v62 uint = 972
        _ = v62
    } else {
        var v63 int32 = p2
        if true {
            v64 := GenS9{F0: 199, F1: 42, F2: 915}
            v65 := []int{2147483647, -9223372036854775808, 4294967295, -226}
            _ = v64
            _ = v65
        } else {
            v63 = -313
            var v66 int16 = 32767
            _ = v66
        }
        v67 := []int{}
        _ = v63
        _ = v67
    }
    switch -9223372036854775808 {
    case 910:
        var v68 uint32 = 42
        _ = v68
    case 2147483647:
        p1 = (p1 + (p1 + "hello"))
    case 0:
        fmt.Println(p2, p2)
    default:
        v61 = (v61 / ((1) | 1))
    }
    fmt.Println(v61, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v61
    return 1
}

func main() {
    var v1 uint64 = 0
    var v2 uint16 = 32767
    var v3 float64 = 0.0
    for fi69 := 0; fi69 < 4; fi69++ {
        v1 = (v1 - 1)
        var v70 uint32 = 2147483647
        var v71 uint16 = 449
        _ = fi69
        _ = v70
        _ = v71
    }
    if (!false) {
        if (-2.25 >= (0.0 * ((0.1 + 1e10) / 1.0))) {
            switch "" {
            case "":
                v1 = v1
            case "hello":
                v72 := GenS10{F0: 42}
                _ = v72
            case "foo bar":
                if false {
                    for fi73 := 0; fi73 < 2; fi73++ {
                        v74 := []int{(^(fi73)), (^(fi73)), (-(fi73)), (fi73 - fi73)}
                        _ = fi73
                        _ = v74
                    }
                    fmt.Println(v1)
                    v2 = v2
                } else {
                    v75 := GenS10{F0: 1}
                    if ((1 >= 9223372036854775807) && ("\n\t" != "")) {
                        if false {
                            v2 = (v2 - (v2 % ((32767) | 1)))
                        }
                        if (("foo bar" + "\n\t") != (("a" + "") + "\n\t")) {
                            v2 = (v2 >> ((42) & 63))
                        } else {
                            var v76 bool = (-1 == 1)
                            v77 := []int{42}
                            var v78 int = 32768
                            _ = v76
                            _ = v77
                            _ = v78
                        }
                    } else {
                        for fi79 := 0; fi79 < 9; fi79++ {
                            v1 = (v1 + (v1 % ((18446744073709551615) | 1)))
                            _ = fi79
                        }
                        fmt.Println(v2)
                    }
                    _ = v75
                }
            default:
                fmt.Println(v1, v2, v1)
            }
            v1 += v1
            switch 9223372036854775807 {
            case -1:
                v80 := GenS11{F0: 1, F1: 0, F2: -345}
                _ = v80
            case 9223372036854775807:
                v81 := GenS12{F0: ("" + (("foo bar" + "\n\t") + "foo bar"))}
                _ = v81
            case -9223372036854775808:
                var v82 int32 = -2147483648
                _ = v82
            }
        }
    } else {
        for fi83 := 0; fi83 < 1; fi83++ {
            var v84 uint = 705
            if true {
                v85 := []int{}
                for fi86 := 0; fi86 < 10; fi86++ {
                    if (((100 >= 1) && ((352 == -32768) || (!false))) || ((^(fi83)) != (-(fi86)))) {
                        switch fi86 {
                        case -577:
                            v87 := GenS5{F0: 42, F1: (v2 + v2), F2: 64}
                            _ = v87
                        default:
                            v84 = 65535
                        }
                    }
                    fmt.Println(v2)
                    v84 += (v84 | (v84 + (v84 | 42)))
                    _ = fi86
                }
                if true {
                    if len(v85) >= 2 {
                        v85 = v85[1:len(v85)]
                    }
                    switch ((("hello" + "hello") + "\n\t") + "foo bar") {
                    case "a":
                        var v88 uint8 = 0
                        _ = v88
                    case "":
                        v84 += v84
                    }
                }
                _ = v85
            }
            _ = fi83
            _ = v84
        }
        v89 := []int{1, 42, -649}
        _ = v89
    }
    v90 := []int{}
    var v91 int8 = 0
    for fi92 := 0; fi92 < 15; fi92++ {
        var v93 int = 42
        if (-1 >= -32768) {
            v3 -= ((0.0 + v3) * v3)
            if ((-9223372036854775808 != 42) && (127 >= 42)) {
                fmt.Println(v1, v1)
            } else {
                if ((-2147483648 == -257) && (!((((float64(float32(v3)) * 1.5) + 0.1) / v3) < (v3 / 0.1)))) {
                    v3 -= ((((float64((1e10 - ((3.14159 - ((0.0 + -2.25) - (((float32(v3) - (1e10 + (0.0 * 1.5))) * 0.0) * 1e10))) / 1e10))) / v3) - v3) - 0.1) / float64(-2.25))
                    v94 := GenS3{F0: v2, F1: int8(v2), F2: 595}
                    v1 += v1
                    _ = v94
                } else {
                    var v95 int = fi92
                    _ = v95
                }
                switch "" {
                case "a":
                    switch "" {
                    case "":
                        var v96 float64 = v3
                        _ = v96
                    case "foo bar":
                        if false {
                            v97 := GenS13{F0: (-(v91)), F1: (v1 / ((v1) | 1))}
                            if false {
                                v2 = v2
                                fmt.Println(fi92, v91)
                                v3 = ((0.0 - 0.0) - 3.14159)
                            }
                            _ = v97
                        }
                    case "hello":
                        var v98 int8 = v91
                        _ = v98
                    default:
                        v3 = (v3 + v3)
                    }
                }
            }
        } else {
            if false {
                if (127 >= 2) {
                    var v99 uint64 = (v1 | (v1 * v1))
                    fmt.Println(v93, v1, v91)
                    _ = v99
                } else {
                    v1 -= 1
                    var v100 uint8 = 42
                    _ = v100
                }
                for fi101 := 0; fi101 < 6; fi101++ {
                    if len(v90) > 0 {
                        v90[0] = 4294967295
                    }
                    fmt.Println(v91)
                    v90 = append(v90, 4294967295)
                    _ = fi101
                }
                if false {
                    if true {
                        switch -580 {
                        case -47:
                            v102 := GenS7{F0: (v91 << ((0) & 63)), F1: "a"}
                            _ = v102
                        case 127:
                            if len(v90) > 0 {
                                v90[0] = (fi92 * fi92)
                            }
                        default:
                            for fi103 := 0; fi103 < 7; fi103++ {
                                v1 += v1
                                _ = fi103
                            }
                        }
                    }
                    if (!false) {
                        if (!(!true)) {
                            if len(v90) >= 2 {
                                v90 = v90[1:len(v90)]
                            }
                            v104 := GenS2{F0: v2}
                            _ = v104
                        }
                    }
                    v105 := GenS9{F0: 128, F1: 244, F2: -32768}
                    _ = v105
                } else {
                    _ = len(v90)
                    fmt.Println(v1, fi92)
                    for fi106 := 0; fi106 < 10; fi106++ {
                        v107 := GenS9{F0: 2, F1: 0, F2: 1}
                        fmt.Println(fi106)
                        v3 = v3
                        var v108 int32 = 0
                        var v109 string = "a"
                        _ = fi106
                        _ = v107
                        _ = v108
                        _ = v109
                    }
                }
            } else {
                var v110 uint16 = (v2 >> ((1) & 63))
                var v111 int32 = -198
                _ = v110
                _ = v111
            }
            v91 = (v91 * (v91 - -1))
            v1 += (v1 >> ((814) & 63))
        }
        _ = fi92
        _ = v93
    }
    for fi112 := 0; fi112 < 13; fi112++ {
        if false {
            v3 -= 0.1
        } else {
            var v113 float32 = ((-2.25 + 0.1) * (0.1 + 3.14159))
            fmt.Println(v1)
            switch (-(fi112)) {
            case -1:
                if (!(!(-32768 > 283))) {
                    v114 := GenS3{F0: 240, F1: (^(v91)), F2: 42}
                    v3 += float64(v113)
                    _ = v114
                }
            case 0:
                switch fi112 {
                case -31:
                    var v115 int = (^(fi112))
                    _ = v115
                case 0:
                    v113 += 0.0
                case 1:
                    v116 := GenS14{F0: 42, F1: false}
                    _ = v116
                }
            case -9223372036854775808:
                switch (fi112 - (-(fi112))) {
                case -402:
                    v117 := GenS15{F0: -1}
                    _ = v117
                }
            default:
                v3 -= float64(v113)
            }
            _ = v113
        }
        for fi118 := 0; fi118 < 18; fi118++ {
            fmt.Println(v3, fi118)
            _ = fi118
        }
        _ = fi112
    }
    v119 := GenS4{F0: (2 != 1)}
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v90:", len(v90), v90)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v119
    _ = v2
    _ = v3
    _ = v90
    _ = v91
}
