package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 int16; F1 int8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint8; F1 uint; F2 int16 }
type GenS7 struct { F0 int64; F1 int16; F2 int32 }
type GenS8 struct { F0 uint64; F1 uint64; F2 uint32 }
type GenS9 struct { F0 int32; F1 uint8 }
type GenS10 struct { F0 int16 }
func genFn0() int32 {
    if true {
        var v6 int = 9223372036854775807
        if true {
            v7 := GenS1{F0: 100, F1: 272}
            var v8 int32 = 2147483647
            _ = v7
            _ = v8
        } else {
            v6 = 958
        }
        _ = v6
    } else {
        var v9 float64 = 1e10
        _ = v9
    }
    v10 := GenS2{F0: 0}
    v11 := GenS3{F0: 346, F1: "hello"}
    v12 := GenS4{F0: 1, F1: -128}
    for fi13 := 0; fi13 < 17; fi13++ {
        var v14 uint16 = 255
        var v15 uint64 = 4294967295
        _ = fi13
        _ = v14
        _ = v15
    }
    _ = v10
    _ = v11
    _ = v12
    return 700
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int8 = 127
    var v3 string = "foo bar"
    v4 := make(map[string]int)
    var v5 int64 = -1
    var v16 bool = ((!true) && false)
    var v17 int32 = -325
    v2 -= (v2 << ((672) & 63))
    for fi18 := 0; fi18 < 2; fi18++ {
        for fi19 := 0; fi19 < 10; fi19++ {
            fmt.Println(fi18, fi18)
            fmt.Println(fi19)
            v17 = 42
            fmt.Println(fi18, v2)
            _ = fi19
        }
        v16 = v16
        for fi20 := 0; fi20 < 10; fi20++ {
            if (!((v5 + (v5 % (((v5 * v5)) | 1))) == 9223372036854775807)) {
                v21 := GenS3{F0: 42, F1: v3}
                v22 := GenS5{F0: 42}
                switch (fi20 % ((fi20) | 1)) {
                case -937:
                    switch v3 {
                    case "":
                        var v23 uint64 = 0
                        _ = v23
                    case "foo bar":
                        var v24 int16 = -1
                        _ = v24
                    case "a":
                        switch -1 {
                        case 32768:
                            _ = len(v4)
                        case -1:
                            v16 = (65535 != 32767)
                        }
                    default:
                        switch (fi20 ^ (-(fi20))) {
                        case 0:
                            var v25 uint32 = 347
                            _ = v25
                        case 2147483647:
                            if v16 {
                                switch "" {
                                case "hello":
                                    switch (v3 + "") {
                                    case "":
                                        v26 := GenS1{F0: (-(v2)), F1: 255}
                                        _ = v26
                                    case "a":
                                        if false {
                                            v16 = (v5 >= v5)
                                            v27 := []int{42}
                                            fmt.Println(fi20, v17, v3)
                                            _ = v27
                                        } else {
                                            var v28 uint = 235
                                            switch ((v3 + (v3 + (v3 + (v3 + "\n\t")))) + v3) {
                                            case "hello":
                                                var v29 int64 = v5
                                                _ = v29
                                            case "":
                                                v30 := GenS6{F0: 255, F1: 18446744073709551615, F2: 1000}
                                                _ = v30
                                            case "a":
                                                v31 := GenS7{F0: -1, F1: 1, F2: (-(v17))}
                                                _ = v31
                                            }
                                            v28 += v28
                                            v32 := GenS3{F0: 42, F1: v3}
                                            fmt.Println(v3, v5, v16)
                                            _ = v28
                                            _ = v32
                                        }
                                    case "hello":
                                        v17 -= (v17 % ((genFn0()) | 1))
                                    }
                                case "a":
                                    if false {
                                        var v33 uint16 = 255
                                        _ = v33
                                    }
                                }
                            } else {
                                if v16 {
                                    v16 = true
                                    if (v16 || true) {
                                        v34 := []int{(-(fi18)), -9223372036854775808, fi20}
                                        if v16 {
                                            v35 := GenS8{F0: 515, F1: 4294967295, F2: 42}
                                            _ = v35
                                        } else {
                                            fmt.Println(v2, v17, v5)
                                        }
                                        _ = v34
                                    }
                                    v36 := GenS9{F0: -1, F1: 120}
                                    _ = v36
                                }
                                var v37 uint = 18446744073709551615
                                var v38 int = (fi20 - (fi18 >> ((v37) & 63)))
                                v17 = (v17 << (((^(v37))) & 63))
                                var v39 uint8 = uint8(v37)
                                v40 := []int{}
                                _ = v37
                                _ = v38
                                _ = v39
                                _ = v40
                            }
                        case -24:
                            v41 := []int{968, (fi18 - (-(fi20)))}
                            _ = v41
                        }
                    }
                case 2147483647:
                    var v42 int64 = (-(v5))
                    _ = v42
                case 9223372036854775807:
                    if (4294967295 < 255) {
                        fmt.Println(v3)
                        var v43 uint8 = 210
                        v44 := []int{128, fi18}
                        if (-9223372036854775808 <= int(v5)) {
                            fmt.Println(fi20, v17)
                            if true {
                                var v45 uint64 = 38
                                _ = v45
                            }
                            if (!v16) {
                                v5 = (v5 << ((764) & 63))
                                v46 := GenS9{F0: (v17 >> ((18446744073709551615) & 63)), F1: (v43 | v43)}
                                _ = v46
                            }
                        }
                        _ = v43
                        _ = v44
                    } else {
                        if false {
                            fmt.Println(fi18)
                            v47 := []int{fi18, fi20, (-(fi20))}
                            var v48 uint32 = 979
                            _ = v47
                            _ = v48
                        }
                        var v49 string = ("hello" + v3)
                        v3 = v49
                        var v50 uint64 = 92
                        _ = v49
                        _ = v50
                    }
                    var v51 int64 = (v5 / (((-(v5))) | 1))
                    _ = v51
                }
                _ = v21
                _ = v22
            }
            _ = fi20
        }
        _ = fi18
    }
    switch -9223372036854775808 {
    case 523:
        v52 := v1[1]
        _ = v52
    case -226:
        if v16 {
            var v53 int = -1
            if v16 {
                for fi54 := 0; fi54 < 18; fi54++ {
                    fmt.Println(v2, fi54)
                    switch 42 {
                    case 82:
                        for fi55 := 0; fi55 < 12; fi55++ {
                            v5 = v5
                            v56 := GenS2{F0: 42}
                            var v57 bool = (42 < -1)
                            _ = fi55
                            _ = v56
                            _ = v57
                        }
                    case 2147483647:
                        switch v53 {
                        case 660:
                            for fi58 := 0; fi58 < 9; fi58++ {
                                var v59 int32 = (v17 + v17)
                                _ = fi58
                                _ = v59
                            }
                        case 718:
                            if (166 >= 42) {
                                for fi60 := 0; fi60 < 1; fi60++ {
                                    var v61 int64 = (v5 << ((65535) & 63))
                                    var v62 uint16 = 1
                                    fmt.Println(fi60, v62)
                                    _ = fi60
                                    _ = v61
                                    _ = v62
                                }
                                switch (v53 ^ (fi54 / (((v53 + fi54)) | 1))) {
                                case 9223372036854775807:
                                    if (int8(v5) == v2) {
                                        if false {
                                            if v16 {
                                                var v63 int16 = -1
                                                var v64 uint64 = 42
                                                var v65 int32 = genFn0()
                                                _ = v63
                                                _ = v64
                                                _ = v65
                                            } else {
                                                var v66 int8 = v2
                                                _ = v66
                                            }
                                        }
                                        for fi67 := 0; fi67 < 1; fi67++ {
                                            fmt.Println(v17, v16, v53)
                                            fmt.Println(fi67, v17, fi67)
                                            v68 := GenS6{F0: 69, F1: 2, F2: 1000}
                                            _ = fi67
                                            _ = v68
                                        }
                                    }
                                case -2147483648:
                                    var v69 int64 = v5
                                    _ = v69
                                }
                            } else {
                                v3 = ("" + "hello")
                            }
                        default:
                            var v70 int32 = genFn0()
                            _ = v70
                        }
                    case 32767:
                        v71 := []int{7, fi54}
                        _ = v71
                    }
                    v3 = (v3 + (("" + v3) + ("a" + v3)))
                    _ = fi54
                }
            }
            for fi72 := 0; fi72 < 20; fi72++ {
                switch v53 {
                case -1:
                    fmt.Println(fi72, fi72)
                }
                if ((v2 * v2) > 1) {
                    if true {
                        v73 := GenS5{F0: 232}
                        switch (-(v53)) {
                        case 0:
                            v2 -= v2
                        case 893:
                            if true {
                                for fi74 := 0; fi74 < 10; fi74++ {
                                    var v75 uint32 = 1
                                    v76 := v4[v3]
                                    var v77 int16 = int16(v2)
                                    _ = fi74
                                    _ = v75
                                    _ = v76
                                    _ = v77
                                }
                            }
                        case 266:
                            var v78 uint32 = 42
                            _ = v78
                        default:
                            var v79 int16 = -678
                            _ = v79
                        }
                        var v80 uint32 = 67
                        _ = v73
                        _ = v80
                    } else {
                        var v81 int8 = v2
                        v53 = fi72
                        _ = v81
                    }
                }
                switch 9223372036854775807 {
                case 0:
                    fmt.Println(v17)
                default:
                    switch v3 {
                    case "":
                        var v82 float64 = float64(3.14159)
                        _ = v82
                    case "a":
                        v16 = ((42 > 127) || ((v5 >> ((65535) & 63)) < -9223372036854775808))
                    default:
                        v17 -= (v17 % ((v17) | 1))
                    }
                }
                _ = fi72
            }
            _ = v53
        } else {
            if true {
                var v83 uint8 = 1
                _ = v83
            } else {
                v84 := GenS2{F0: 42}
                v3 = v3
                _ = v84
            }
        }
    }
    if (!v16) {
        v3 = v3
        if (v17 < (v17 ^ (^(v17)))) {
            fmt.Println(v2, v16, v3)
            if v16 {
                fmt.Println(v3)
                var v85 float32 = 3.14159
                switch 42 {
                case 128:
                    switch -1 {
                    case 128:
                        var v86 int32 = (v17 ^ (v17 ^ v17))
                        _ = v86
                    default:
                        v2 -= v2
                    }
                }
                _ = v85
            }
        } else {
            v2 = (v2 % (((v2 / ((-79) | 1))) | 1))
            v87 := GenS10{F0: 1000}
            _ = v87
        }
    } else {
        for fi88 := 0; fi88 < 19; fi88++ {
            v89 := GenS7{F0: v5, F1: 42, F2: (^(v17))}
            for fi90 := 0; fi90 < 17; fi90++ {
                fmt.Println(v17, v16, v2)
                _ = len(v1)
                var v91 uint = 0
                _ = fi90
                _ = v91
            }
            var v92 float64 = float64(1.5)
            _ = fi88
            _ = v89
            _ = v92
        }
        var v93 int16 = 1
        fmt.Println(v16)
        _ = v93
    }
    if (1e10 <= 1e10) {
        for fi94 := 0; fi94 < 7; fi94++ {
            var v95 uint = 903
            v96 := GenS5{F0: 0}
            _ = fi94
            _ = v95
            _ = v96
        }
        if false {
            v2 = (v2 >> ((2) & 63))
        } else {
            var v97 uint32 = 65535
            v98 := []int{-241, 128, 32767, 1}
            _ = v97
            _ = v98
        }
    } else {
        switch v3 {
        case "hello":
            v2 = (-(v2))
        case "a":
            v3 = (v3 + ("hello" + (v3 + "a")))
        default:
            v16 = (1 != 32767)
        }
        v99 := GenS10{F0: 32767}
        _ = v99
    }
    v3 = (("\n\t" + (v3 + "")) + ((v3 + (v3 + "\n\t")) + ""))
    v17 = v17
    for fi100 := 0; fi100 < 15; fi100++ {
        var v101 uint32 = 2147483648
        var v102 float64 = (float64(1e10) / ((1.5 + float64((1e10 - (3.14159 + 3.14159)))) / 1.0))
        if false {
            switch "foo bar" {
            case "hello":
                for fi103 := 0; fi103 < 18; fi103++ {
                    switch -769 {
                    case -244:
                        switch "" {
                        case "\n\t":
                            fmt.Println(v2, v17)
                        }
                    case -601:
                        if (!((-2.25 == float32((v102 + v102))) || (!v16))) {
                            v104 := v4["\n\t"]
                            v102 += v102
                            _ = v104
                        } else {
                            var v105 int = (fi103 + fi100)
                            fmt.Println(v17, fi100)
                            v3 = ("hello" + (v3 + v3))
                            _ = v105
                        }
                    }
                    v3 = v3
                    _ = fi103
                }
            }
            delete(v4, v3)
        }
        _ = fi100
        _ = v101
        _ = v102
    }
    v3 = "a"
    v106 := GenS1{F0: (^(v2)), F1: 2}
    var v107 int64 = v5
    for fi108 := 0; fi108 < 17; fi108++ {
        for fi109 := 0; fi109 < 15; fi109++ {
            v110 := GenS8{F0: 18446744073709551615, F1: 18446744073709551615, F2: 552}
            var v111 int8 = -124
            _ = len(v1)
            _ = fi109
            _ = v110
            _ = v111
        }
        if v16 {
            v112 := []int{fi108, fi108, fi108, fi108}
            switch v3 {
            case "":
                delete(v1, 2147483648)
            case "\n\t":
                for fi113 := 0; fi113 < 7; fi113++ {
                    var v114 int64 = -902
                    var v115 bool = (-529 != 368)
                    _ = fi113
                    _ = v114
                    _ = v115
                }
            case "foo bar":
                for fi116 := 0; fi116 < 9; fi116++ {
                    fmt.Println(v107, v5)
                    var v117 int64 = (^(v107))
                    _ = fi116
                    _ = v117
                }
            }
            v5 -= (v5 + v107)
            _ = v112
        }
        _ = fi108
    }
    if (!true) {
        var v118 uint8 = 198
        _ = v118
    } else {
        if false {
            v119 := []int{int(v2), 497, -864}
            var v120 int64 = (v5 % ((42) | 1))
            _ = v119
            _ = v120
        }
        fmt.Println(v5)
    }
    v4["foo bar"] = -1
    v2 += v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v107:", v107)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v106
    _ = v107
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
