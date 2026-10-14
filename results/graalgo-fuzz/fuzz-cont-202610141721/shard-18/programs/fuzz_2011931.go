package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 uint16; F1 uint; F2 bool }
type GenS5 struct { F0 bool; F1 string }
type GenS6 struct { F0 bool; F1 int }
type GenS7 struct { F0 string; F1 int8; F2 string }
type GenS8 struct { F0 uint8; F1 int64; F2 int16 }
type GenS9 struct { F0 uint8; F1 uint64; F2 int64 }
type GenS10 struct { F0 int; F1 uint32; F2 string }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 uint; F1 string }
type GenS13 struct { F0 bool; F1 uint16 }
type GenS14 struct { F0 uint64 }
type GenS15 struct { F0 uint32; F1 int32 }

func main() {
    var v1 string = ("hello" + ("hello" + ("a" + (("a" + "") + "foo bar"))))
    var v2 bool = (0.0 != 0.1)
    v3 := make(map[uint32]int)
    v4 := make(map[string]string)
    var v5 uint8 = 0
    for fi6 := 0; fi6 < 7; fi6++ {
        delete(v4, "a")
        var v7 uint32 = 65535
        v8 := GenS1{F0: 676}
        _ = fi6
        _ = v7
        _ = v8
    }
    switch -1 {
    case -248:
        v9 := GenS1{F0: 1}
        _ = v9
    case -9223372036854775808:
        v10 := GenS2{F0: "foo bar"}
        _ = v10
    default:
        var v11 int = 4294967295
        _ = v11
    }
    if (9223372036854775807 != 9223372036854775807) {
        if (v2 && ((float32(0.1) + -2.25) <= (-2.25 + ((-2.25 - 0.1) / 1e10)))) {
            for fi12 := 0; fi12 < 12; fi12++ {
                v13 := GenS3{F0: 9223372036854775807, F1: 131}
                if (-9223372036854775808 <= 0) {
                    v14 := v3[2147483648]
                    if (!v2) {
                        v15 := GenS4{F0: 42, F1: 392, F2: v2}
                        _ = v15
                    }
                    _ = v14
                }
                if (1000 >= 42) {
                    var v16 uint = 65535
                    _ = v16
                }
                _ = fi12
                _ = v13
            }
            for fi17 := 0; fi17 < 15; fi17++ {
                if v2 {
                    var v18 int = (-(fi17))
                    for fi19 := 0; fi19 < 5; fi19++ {
                        if (2 != 32767) {
                            var v20 uint8 = v5
                            _ = v20
                        }
                        v1 = ("hello" + v1)
                        var v21 uint32 = 2147483648
                        _ = fi19
                        _ = v21
                    }
                    v2 = v2
                    _ = v18
                }
                _ = fi17
            }
        }
        for fi22 := 0; fi22 < 13; fi22++ {
            v2 = v2
            v23 := GenS5{F0: ((18446744073709551615 <= 648) && ((4294967295 <= 2147483648) && (v2 || v2))), F1: v1}
            _ = fi22
            _ = v23
        }
        delete(v4, v1)
    } else {
        for fi24 := 0; fi24 < 17; fi24++ {
            for fi25 := 0; fi25 < 9; fi25++ {
                switch v1 {
                case "a":
                    fmt.Println(v1, v2)
                case "foo bar":
                    if (true || false) {
                        switch fi24 {
                        case 128:
                            var v26 uint8 = 2
                            _ = v26
                        }
                        v5 += v5
                    }
                default:
                    var v27 int8 = 42
                    _ = v27
                }
                _ = fi25
            }
            v28 := GenS5{F0: v2, F1: v1}
            switch "foo bar" {
            case "\n\t":
                switch fi24 {
                case -831:
                    var v29 uint8 = (v5 - (v5 & (v5 >> ((735) & 63))))
                    _ = v29
                default:
                    if v2 {
                        switch v1 {
                        case "a":
                            for fi30 := 0; fi30 < 17; fi30++ {
                                var v31 uint64 = 2
                                v32 := []int{}
                                _ = fi30
                                _ = v31
                                _ = v32
                            }
                        }
                        if (1 < -1) {
                            v5 = v5
                            v1 = v1
                        } else {
                            for fi33 := 0; fi33 < 6; fi33++ {
                                var v34 uint8 = 2
                                _ = len(v4)
                                var v35 float64 = float64(3.14159)
                                fmt.Println(fi24, v35)
                                _ = fi33
                                _ = v34
                                _ = v35
                            }
                            v36 := GenS2{F0: v1}
                            _ = v36
                        }
                    } else {
                        for fi37 := 0; fi37 < 17; fi37++ {
                            v38 := GenS2{F0: v1}
                            if (!v2) {
                                v5 = (v5 * (v5 << ((789) & 63)))
                                v5 = (v5 ^ 139)
                                v39 := GenS6{F0: (((1 != 32767) || ((!(!(v2 || true))) && v2)) && (0.0 < float32(0.1))), F1: (fi37 - -786)}
                                _ = v39
                            } else {
                                var v40 uint64 = 42
                                v5 += v5
                                fmt.Println(v2, v40, v2)
                                _ = v40
                            }
                            if (!(fi37 == (-(fi24)))) {
                                fmt.Println(v5, fi24)
                                var v41 int32 = 670
                                if v2 {
                                    v42 := v4[(v1 + ("foo bar" + "foo bar"))]
                                    switch (v42 + "") {
                                    case "":
                                        var v43 bool = v2
                                        _ = v43
                                    case "hello":
                                        switch ((v1 + v1) + "a") {
                                        case "\n\t":
                                            v1 = "a"
                                            var v44 string = "foo bar"
                                            switch -587 {
                                            case 32767:
                                                if false {
                                                    v45 := GenS7{F0: "a", F1: 100, F2: v42}
                                                    if (-2.25 == 0.0) {
                                                        fmt.Println(v2)
                                                        v1 = "hello"
                                                        if v2 {
                                                            v46 := GenS8{F0: 255, F1: -822, F2: 0}
                                                            switch fi37 {
                                                            case 32767:
                                                                var v47 int = fi37
                                                                _ = v47
                                                            case -2147483648:
                                                                delete(v4, "a")
                                                            case 42:
                                                                v48 := v3[2]
                                                                _ = v48
                                                            default:
                                                                fmt.Println(fi37, fi37)
                                                            }
                                                            _ = v46
                                                        } else {
                                                            v1 = (v44 + v44)
                                                            if v2 {
                                                                v49 := GenS9{F0: v5, F1: 0, F2: 337}
                                                                _ = v49
                                                            } else {
                                                                if v2 {
                                                                    if v2 {
                                                                        if (42 < 255) {
                                                                            v50 := GenS6{F0: v2, F1: 32767}
                                                                            _ = v50
                                                                        } else {
                                                                            switch 9223372036854775807 {
                                                                            case 2147483647:
                                                                                v51 := GenS10{F0: (fi37 << ((4294967295) & 63)), F1: 594, F2: (v44 + "hello")}
                                                                                _ = v51
                                                                            }
                                                                        }
                                                                        if v2 {
                                                                            if (!v2) {
                                                                                fmt.Println(v2, v42, v2)
                                                                            } else {
                                                                                v52 := []int{42}
                                                                                fmt.Println(v5, v42, v1)
                                                                                if ((v2 && true) && true) {
                                                                                    v2 = (!(v44 > "hello"))
                                                                                    _ = len(v4)
                                                                                    if (0 <= -9223372036854775808) {
                                                                                        v53 := GenS4{F0: 32767, F1: 22, F2: v2}
                                                                                        _ = v53
                                                                                    }
                                                                                } else {
                                                                                    var v54 uint32 = 21
                                                                                    _ = v54
                                                                                }
                                                                                if (v2 && v2) {
                                                                                    v5 += v5
                                                                                } else {
                                                                                    var v55 float32 = (0.1 * -2.25)
                                                                                    v56 := GenS11{F0: 598}
                                                                                    v44 = (v1 + ("foo bar" + "hello"))
                                                                                    _ = v55
                                                                                    _ = v56
                                                                                }
                                                                                if len(v52) > 0 {
                                                                                    v52[0] = fi24
                                                                                }
                                                                                _ = v52
                                                                            }
                                                                            fmt.Println(fi24, v5)
                                                                            fmt.Println(v42, v2, v1)
                                                                        } else {
                                                                            v5 += (v5 / (((v5 ^ v5)) | 1))
                                                                            switch "foo bar" {
                                                                            case "a":
                                                                                v2 = (v2 && (fi37 != fi37))
                                                                            case "\n\t":
                                                                                _ = len(v4)
                                                                            case "hello":
                                                                                var v57 string = v44
                                                                                _ = v57
                                                                            default:
                                                                                v41 += (v41 & (v41 >> ((65535) & 63)))
                                                                            }
                                                                        }
                                                                        v58 := GenS7{F0: v1, F1: -1, F2: (v42 + (v1 + (("hello" + (v42 + "\n\t")) + "")))}
                                                                        fmt.Println(v42, fi24, fi24)
                                                                        _ = v58
                                                                    }
                                                                    fmt.Println(v2, v42)
                                                                    var v59 int32 = v41
                                                                    _ = v59
                                                                } else {
                                                                    switch ("hello" + "hello") {
                                                                    case "\n\t":
                                                                        _ = len(v3)
                                                                    }
                                                                    _ = len(v3)
                                                                }
                                                                switch fi37 {
                                                                case 1:
                                                                    var v60 float64 = (0.0 - 0.1)
                                                                    _ = v60
                                                                case 2147483647:
                                                                    var v61 int = fi37
                                                                    _ = v61
                                                                default:
                                                                    var v62 uint64 = 1
                                                                    _ = v62
                                                                }
                                                            }
                                                            fmt.Println(v5, v44, v44)
                                                        }
                                                        var v63 float32 = ((1e10 + 1e10) / 1.0)
                                                        switch (v44 + "\n\t") {
                                                        case "a":
                                                            var v64 int16 = 42
                                                            _ = v64
                                                        case "\n\t":
                                                            v41 += v41
                                                        }
                                                        _ = v63
                                                    } else {
                                                        var v65 bool = false
                                                        var v66 uint = 2
                                                        v2 = v2
                                                        if (556 >= -269) {
                                                            v65 = v2
                                                            v2 = v2
                                                            if (1000 != 1000) {
                                                                v67 := GenS12{F0: (v66 << (((v66 ^ 860)) & 63)), F1: v44}
                                                                _ = v67
                                                            }
                                                        } else {
                                                            var v68 string = (v1 + ("\n\t" + v1))
                                                            _ = v68
                                                        }
                                                        _ = v65
                                                        _ = v66
                                                    }
                                                    _ = v45
                                                } else {
                                                    switch (fi24 - 400) {
                                                    case 0:
                                                        _ = len(v3)
                                                    case -133:
                                                        v69 := GenS4{F0: 2, F1: 745, F2: false}
                                                        fmt.Println(fi24)
                                                        var v70 int32 = (v41 * (v41 + 1))
                                                        _ = v69
                                                        _ = v70
                                                    }
                                                }
                                            case 921:
                                                v41 -= v41
                                            case -151:
                                                v71 := GenS10{F0: (^(fi24)), F1: 1, F2: "a"}
                                                _ = v71
                                            default:
                                                v2 = false
                                            }
                                            _ = v44
                                        case "foo bar":
                                            var v72 bool = v2
                                            _ = v72
                                        case "hello":
                                            v73 := GenS4{F0: 0, F1: 281, F2: (-1 <= 32767)}
                                            _ = v73
                                        }
                                    case "foo bar":
                                        var v74 int = (^(fi37))
                                        _ = v74
                                    default:
                                        v41 = (v41 >> ((4294967295) & 63))
                                    }
                                    _ = v42
                                } else {
                                    v41 = v41
                                    if false {
                                        v41 -= v41
                                    }
                                    switch v1 {
                                    case "\n\t":
                                        var v75 int8 = 127
                                        _ = v75
                                    case "hello":
                                        v76 := GenS13{F0: v2, F1: 255}
                                        _ = v76
                                    case "a":
                                        v5 -= (^(v5))
                                    default:
                                        fmt.Println(v5, fi24, fi24)
                                    }
                                }
                                switch (v1 + v1) {
                                case "a":
                                    v77 := GenS14{F0: 2}
                                    _ = v77
                                }
                                var v78 uint16 = 42
                                switch v1 {
                                case "foo bar":
                                    v79 := GenS6{F0: v2, F1: fi24}
                                    _ = v79
                                case "hello":
                                    v41 += -52
                                case "":
                                    v2 = (!((-1 < 1) && v2))
                                default:
                                    switch ("hello" + "\n\t") {
                                    case "":
                                        var v80 int8 = -110
                                        _ = v80
                                    case "foo bar":
                                        var v81 int64 = 9223372036854775807
                                        _ = v81
                                    case "a":
                                        switch fi37 {
                                        case 128:
                                            var v82 uint = 18446744073709551615
                                            _ = v82
                                        case 32768:
                                            fmt.Println(fi37)
                                        case 4294967295:
                                            v41 = (v41 >> ((255) & 63))
                                            var v83 uint16 = 608
                                            _ = v83
                                        }
                                    }
                                }
                                var v84 int = 668
                                _ = v41
                                _ = v78
                                _ = v84
                            }
                            _ = fi37
                            _ = v38
                        }
                        v3[2147483647] = (fi24 - (-(fi24)))
                        switch v1 {
                        case "a":
                            switch fi24 {
                            case 455:
                                if (!(319 >= uint(v5))) {
                                    v85 := []int{}
                                    v1 = "a"
                                    if (v2 || (42 < -9223372036854775808)) {
                                        var v86 float64 = -2.25
                                        _ = v86
                                    } else {
                                        for fi87 := 0; fi87 < 12; fi87++ {
                                            var v88 uint64 = 18446744073709551615
                                            v2 = (v1 < v1)
                                            var v89 int32 = -1
                                            _ = fi87
                                            _ = v88
                                            _ = v89
                                        }
                                        fmt.Println(v5)
                                    }
                                    _ = v85
                                }
                            case 127:
                                v1 = ""
                            case -1:
                                _ = len(v4)
                            default:
                                for fi90 := 0; fi90 < 10; fi90++ {
                                    switch fi24 {
                                    case 9223372036854775807:
                                        v91 := GenS1{F0: 65535}
                                        _ = v91
                                    default:
                                        v92 := []int{}
                                        _ = v92
                                    }
                                    delete(v3, 65535)
                                    _ = fi90
                                }
                            }
                        case "\n\t":
                            for fi93 := 0; fi93 < 17; fi93++ {
                                v2 = v2
                                _ = fi93
                            }
                        case "":
                            for fi94 := 0; fi94 < 14; fi94++ {
                                if v2 {
                                    if (!(true || v2)) {
                                        v1 = ("hello" + (v1 + v1))
                                        var v95 int8 = 97
                                        v96 := GenS8{F0: 2, F1: 0, F2: -1}
                                        _ = v95
                                        _ = v96
                                    } else {
                                        if false {
                                            v1 = (v1 + (v1 + v1))
                                            if (v2 || (42 != 1)) {
                                                if v2 {
                                                    v1 = (("hello" + ("foo bar" + v1)) + (v1 + "hello"))
                                                    v1 = v1
                                                    v5 += (^(v5))
                                                }
                                                var v97 bool = v2
                                                v5 += v5
                                                fmt.Println(v1)
                                                var v98 float64 = 1e10
                                                v98 = (v98 + 1.5)
                                                _ = v97
                                                _ = v98
                                            } else {
                                                fmt.Println(v1, v2, fi94)
                                                switch v1 {
                                                case "hello":
                                                    if (-1 >= -1) {
                                                        switch (v1 + (((v1 + "") + (v1 + "")) + "a")) {
                                                        case "hello":
                                                            if (v5 >= (v5 * v5)) {
                                                                switch "\n\t" {
                                                                case "a":
                                                                    v5 = 127
                                                                case "\n\t":
                                                                    v99 := []int{}
                                                                    _ = v99
                                                                default:
                                                                    v1 = v1
                                                                }
                                                                if false {
                                                                    v1 = v1
                                                                } else {
                                                                    var v100 string = ("\n\t" + "a")
                                                                    var v101 uint16 = 0
                                                                    _ = v100
                                                                    _ = v101
                                                                }
                                                                var v102 int32 = 446
                                                                var v103 float64 = 3.14159
                                                                _ = v102
                                                                _ = v103
                                                            }
                                                            var v104 int64 = int64(fi94)
                                                            var v105 float32 = 0.1
                                                            _ = v104
                                                            _ = v105
                                                        case "":
                                                            v2 = true
                                                        case "a":
                                                            v106 := []int{fi94}
                                                            _ = v106
                                                        }
                                                    }
                                                case "foo bar":
                                                    fmt.Println(fi94, v2, fi94)
                                                case "a":
                                                    var v107 uint16 = 65535
                                                    _ = v107
                                                }
                                                v3[1] = int(v5)
                                            }
                                        }
                                        v108 := []int{127, (^(fi94)), fi94, 2147483647}
                                        v5 -= v5
                                        _ = v108
                                    }
                                    switch "foo bar" {
                                    case "a":
                                        fmt.Println(v2)
                                    case "":
                                        fmt.Println(v2, fi94, v2)
                                    case "hello":
                                        var v109 int32 = -1
                                        _ = v109
                                    default:
                                        v110 := []int{fi94, 32767, fi24, fi24}
                                        _ = v110
                                    }
                                    switch v1 {
                                    case "hello":
                                        var v111 int8 = -2
                                        _ = v111
                                    case "foo bar":
                                        switch fi24 {
                                        case 128:
                                            v112 := GenS15{F0: 1, F1: 0}
                                            _ = v112
                                        }
                                    default:
                                        if v2 {
                                            var v113 int = fi24
                                            _ = v113
                                        } else {
                                            fmt.Println(v2, v1, fi24)
                                            fmt.Println(v2, v2)
                                        }
                                    }
                                }
                                var v114 uint64 = 1
                                v115 := GenS14{F0: 2}
                                _ = fi94
                                _ = v114
                                _ = v115
                            }
                        default:
                            var v116 uint64 = 13
                            _ = v116
                        }
                    }
                }
            case "foo bar":
                for fi117 := 0; fi117 < 3; fi117++ {
                    if (0 != 100) {
                        v2 = v2
                        var v118 uint = 426
                        _ = v118
                    }
                    if (!((!true) || (((0.0 * (0.1 - (0.1 - 0.0))) / 0.1) < -2.25))) {
                        fmt.Println(v5)
                        v119 := GenS15{F0: 2, F1: -2147483648}
                        var v120 bool = v2
                        _ = v119
                        _ = v120
                    } else {
                        v1 = (v1 + ("" + "foo bar"))
                        v121 := GenS14{F0: 42}
                        var v122 bool = (v2 || false)
                        switch (fi117 - (^(fi24))) {
                        case 32767:
                            var v123 uint = 767
                            v124 := []int{(fi24 << ((119) & 63)), 32767, fi24}
                            v125 := GenS4{F0: 2, F1: v123, F2: true}
                            _ = v123
                            _ = v124
                            _ = v125
                        }
                        _ = v121
                        _ = v122
                    }
                    _ = fi117
                }
            }
            _ = fi24
            _ = v28
        }
    }
    var v126 int = -2147483648
    v126 += v126
    v127 := []int{-1, 2147483647, v126}
    var v128 int32 = -366
    fmt.Println(v5, v1, v1)
    fmt.Println(v1)
    v5 -= 0
    fmt.Println("v1:", v1)
    fmt.Println("v126:", v126)
    fmt.Println("slicelen v127:", len(v127), v127)
    fmt.Println("v128:", v128)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v126
    _ = v127
    _ = v128
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
