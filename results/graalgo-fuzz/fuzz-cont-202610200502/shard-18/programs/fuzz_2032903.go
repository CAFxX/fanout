package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 string; F1 int64 }
type GenS8 struct { F0 int8; F1 uint }
type GenS9 struct { F0 int16; F1 uint32 }
type GenS10 struct { F0 int32; F1 bool }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]uint32)
    v3 := make(map[int]uint32)
    v4 := make(map[string]string)
    v5 := make(map[uint32]int)
    v6 := make(map[uint32]string)
    var v7 int = -1
    v8 := GenS1{F0: 0, F1: v7}
    var v9 uint = 344
    var v10 int8 = 127
    _ = len(v3)
    fmt.Println(v9, v9, v7)
    v11 := GenS1{F0: 1, F1: int(v9)}
    var v12 float32 = 1.5
    var v13 uint16 = 558
    var v14 int = v7
    if false {
        if false {
            switch "" {
            case "\n\t":
                if true {
                    switch v14 {
                    case 1:
                        var v15 int64 = 1
                        _ = v15
                    }
                    for fi16 := 0; fi16 < 8; fi16++ {
                        if true {
                            v12 -= 1.5
                            v17 := GenS2{F0: (v13 << ((uint(v13)) & 63)), F1: 0}
                            if true {
                                if (v10 >= v10) {
                                    v9 += (v9 * v9)
                                    v18 := GenS3{F0: (v9 * (v9 + (v9 >> ((2) & 63)))), F1: 2147483647}
                                    _ = v18
                                } else {
                                    v19 := []int{}
                                    _ = v19
                                }
                                if (!(877 < 4294967295)) {
                                    var v20 bool = true
                                    if v20 {
                                        v3[(v14 / ((42) | 1))] = 2147483648
                                    }
                                    v21 := []int{(fi16 ^ v7), 32767}
                                    _ = v20
                                    _ = v21
                                } else {
                                    if (false && (!(!(-1 <= (v10 + (-(v10))))))) {
                                        var v22 uint8 = 128
                                        v23 := []int{v14, 1}
                                        for fi24 := 0; fi24 < 12; fi24++ {
                                            if (!false) {
                                                switch (("\n\t" + (((("hello" + "a") + "a") + "foo bar") + ("foo bar" + (("foo bar" + "a") + "hello")))) + ("foo bar" + "a")) {
                                                case "foo bar":
                                                    if true {
                                                        if len(v23) >= 2 {
                                                            v23 = v23[1:len(v23)]
                                                        }
                                                        var v25 int32 = 258
                                                        _ = v25
                                                    } else {
                                                        v9 += (^(v9))
                                                        var v26 string = "a"
                                                        _ = v26
                                                    }
                                                case "a":
                                                    switch "hello" {
                                                    case "\n\t":
                                                        if (0 != 741) {
                                                            var v27 string = "a"
                                                            switch v27 {
                                                            case "a":
                                                                if true {
                                                                    fmt.Println(v9, v13, v7)
                                                                } else {
                                                                    var v28 bool = false
                                                                    v29 := GenS2{F0: (v13 / ((v13) | 1)), F1: 18446744073709551615}
                                                                    v30 := GenS1{F0: (v22 << ((v9) & 63)), F1: -30}
                                                                    _ = v28
                                                                    _ = v29
                                                                    _ = v30
                                                                }
                                                                var v31 float64 = ((0.0 + float64(-2.25)) - 0.1)
                                                                v32 := GenS4{F0: 2}
                                                                _ = v31
                                                                _ = v32
                                                            case "hello":
                                                                v33 := GenS2{F0: (v13 % ((32767) | 1)), F1: 59}
                                                                _ = v33
                                                            default:
                                                                delete(v4, v27)
                                                            }
                                                            _ = v27
                                                        } else {
                                                            switch -9223372036854775808 {
                                                            case 4294967295:
                                                                fmt.Println(v13)
                                                            case 32767:
                                                                if (-884 <= 766) {
                                                                    _ = len(v1)
                                                                    _ = len(v3)
                                                                } else {
                                                                    v10 = (-(v10))
                                                                    var v34 int16 = 42
                                                                    _ = v34
                                                                }
                                                            }
                                                            v35 := v1[568]
                                                            var v36 float64 = 3.14159
                                                            var v37 uint8 = (v22 * 27)
                                                            _ = v35
                                                            _ = v36
                                                            _ = v37
                                                        }
                                                    }
                                                case "\n\t":
                                                    if true {
                                                        v12 += v12
                                                    }
                                                default:
                                                    _ = len(v23)
                                                }
                                                var v38 float64 = 3.14159
                                                v38 -= (1e10 - 0.1)
                                                _ = v38
                                            }
                                            _ = fi24
                                        }
                                        _ = v22
                                        _ = v23
                                    }
                                }
                                v39 := GenS2{F0: (v13 / (((v13 >> ((4294967295) & 63))) | 1)), F1: 1}
                                _ = v39
                            }
                            _ = v17
                        }
                        switch "a" {
                        case "\n\t":
                            switch "" {
                            case "foo bar":
                                fmt.Println(v10, v13, v13)
                            }
                        default:
                            if (!false) {
                                var v40 uint64 = 360
                                _ = v40
                            }
                        }
                        _ = fi16
                    }
                    for fi41 := 0; fi41 < 3; fi41++ {
                        if false {
                            var v42 uint = v9
                            v42 = v9
                            v3[32768] = 2
                            _ = v42
                        }
                        _ = fi41
                    }
                } else {
                    switch (v14 | (v14 - v7)) {
                    case -96:
                        v43 := GenS5{F0: v10}
                        _ = v43
                    case 352:
                        _ = len(v1)
                    case -502:
                        fmt.Println(v14, v12)
                    }
                    if true {
                        if (uint8(v13) < 127) {
                            var v44 uint64 = 2
                            var v45 string = ("" + "")
                            fmt.Println(v12)
                            _ = v44
                            _ = v45
                        }
                    } else {
                        if false {
                            if (!false) {
                                var v46 int = (-(v14))
                                var v47 bool = (-1 == 127)
                                _ = v46
                                _ = v47
                            } else {
                                var v48 int64 = 1
                                fmt.Println(v48, v48)
                                v13 -= 139
                                _ = v48
                            }
                            v49 := []int{(v14 % (((v7 ^ 929)) | 1)), (v14 ^ v7)}
                            _ = v49
                        }
                    }
                }
            case "":
                v50 := GenS6{F0: 1}
                _ = v50
            }
            for fi51 := 0; fi51 < 6; fi51++ {
                v52 := GenS6{F0: 0}
                v9 += 18446744073709551615
                _ = fi51
                _ = v52
            }
        } else {
            v53 := []int{2147483647, 42, (v7 & (v14 % (((-(v7))) | 1))), v7}
            for fi54 := 0; fi54 < 10; fi54++ {
                v9 -= (^(v9))
                v53 = append(v53, (v7 | v14))
                var v55 string = ("" + "hello")
                _ = fi54
                _ = v55
            }
            _ = v53
        }
        for fi56 := 0; fi56 < 18; fi56++ {
            var v57 int32 = 1
            v58 := GenS6{F0: 42}
            v57 -= 1
            _ = fi56
            _ = v57
            _ = v58
        }
        if (((((3.14159 * (0.0 / 1e10)) - ((((-2.25 - 1e10) + 3.14159) - 1e10) + 1e10)) + (-2.25 - (-2.25 - float64((float32(0.1) * -2.25))))) == float64(v12)) && (-1 < -365)) {
            var v59 uint8 = 42
            v60 := GenS2{F0: (^(v13)), F1: 42}
            if (42 <= 2147483647) {
                if false {
                    var v61 uint16 = (^(v13))
                    _ = v61
                } else {
                    if true {
                        if true {
                            v62 := GenS5{F0: v10}
                            v63 := GenS4{F0: 61}
                            v64 := []int{(v14 << ((v9) & 63)), v7}
                            _ = v62
                            _ = v63
                            _ = v64
                        } else {
                            for fi65 := 0; fi65 < 8; fi65++ {
                                v66 := GenS7{F0: "", F1: -490}
                                _ = fi65
                                _ = v66
                            }
                        }
                        switch (^(v7)) {
                        case -557:
                            var v67 uint16 = (^(v13))
                            _ = v67
                        case -1:
                            var v68 float32 = (v12 - v12)
                            _ = v68
                        case 106:
                            if false {
                                if (342 == 65535) {
                                    var v69 int64 = 0
                                    _ = v69
                                }
                                v70 := []int{(v14 + (v7 >> ((v9) & 63))), (v7 & int(v13)), 32767, 127}
                                _ = v70
                            }
                        default:
                            if (true && (true && (!true))) {
                                var v71 uint8 = (v59 - (v59 / ((255) | 1)))
                                var v72 int32 = 42
                                _ = v71
                                _ = v72
                            }
                        }
                        v73 := GenS5{F0: v10}
                        _ = v73
                    } else {
                        v13 = v13
                        if false {
                            for fi74 := 0; fi74 < 1; fi74++ {
                                var v75 int64 = -1
                                _ = fi74
                                _ = v75
                            }
                        } else {
                            var v76 int32 = int32(v13)
                            var v77 float32 = v12
                            if (149 != 4294967295) {
                                var v78 uint = (^(v9))
                                var v79 float32 = float32((0.1 / (1.5 - (-2.25 / 1.5))))
                                for fi80 := 0; fi80 < 4; fi80++ {
                                    var v81 uint32 = 42
                                    v1[uint32(v59)] = "hello"
                                    fmt.Println(v77)
                                    _ = fi80
                                    _ = v81
                                }
                                _ = v78
                                _ = v79
                            }
                            _ = v76
                            _ = v77
                        }
                        v9 = (v9 / ((v9) | 1))
                    }
                    for fi82 := 0; fi82 < 11; fi82++ {
                        var v83 int = (fi82 * v7)
                        var v84 bool = false
                        _ = fi82
                        _ = v83
                        _ = v84
                    }
                    fmt.Println(v14, v13, v14)
                }
                v85 := GenS3{F0: 222, F1: 0}
                var v86 uint32 = 1
                _ = v85
                _ = v86
            }
            _ = v59
            _ = v60
        } else {
            var v87 uint16 = 42
            if true {
                for fi88 := 0; fi88 < 20; fi88++ {
                    v89 := GenS7{F0: "hello", F1: -823}
                    _ = fi88
                    _ = v89
                }
                var v90 int8 = (-(v10))
                switch "hello" {
                case "a":
                    v12 = v12
                case "hello":
                    for fi91 := 0; fi91 < 17; fi91++ {
                        var v92 uint64 = 2
                        _ = fi91
                        _ = v92
                    }
                case "":
                    if false {
                        switch "foo bar" {
                        case "a":
                            v87 = v87
                        case "\n\t":
                            var v93 int = 812
                            _ = v93
                        default:
                            if ((v12 + 1e10) >= ((float32(((3.14159 - float64(v12)) * 3.14159)) * v12) - (((1e10 + (float32(((1e10 / 1.0) - 1e10)) / -2.25)) / (float32(1.5) * v12)) + v12))) {
                                var v94 uint16 = (v13 + (v87 + (v13 + v87)))
                                if (0.1 == float64(((v12 * v12) / (v12 / 1e10)))) {
                                    v95 := GenS6{F0: 0}
                                    v96 := []int{(v14 | v14), 9223372036854775807}
                                    var v97 bool = false
                                    _ = v95
                                    _ = v96
                                    _ = v97
                                } else {
                                    v98 := []int{}
                                    fmt.Println(v90, v9, v94)
                                    _ = v98
                                }
                                v10 += (v90 / ((100) | 1))
                                _ = v94
                            }
                        }
                        for fi99 := 0; fi99 < 2; fi99++ {
                            var v100 int32 = -1
                            _ = fi99
                            _ = v100
                        }
                    }
                default:
                    v101 := []int{}
                    _ = v101
                }
                _ = v90
            } else {
                _ = len(v2)
                if true {
                    v102 := GenS8{F0: int8(v9), F1: (^(v9))}
                    for fi103 := 0; fi103 < 2; fi103++ {
                        if ((false || true) && (2147483647 == 2147483648)) {
                            var v104 uint = 357
                            v13 = 508
                            v105 := GenS9{F0: 32767, F1: 2147483647}
                            _ = v104
                            _ = v105
                        } else {
                            for fi106 := 0; fi106 < 7; fi106++ {
                                var v107 int16 = -673
                                switch "hello" {
                                case "hello":
                                    v108 := GenS10{F0: 1, F1: false}
                                    var v109 uint = 42
                                    _ = v108
                                    _ = v109
                                case "":
                                    v14 = (v7 | 170)
                                case "foo bar":
                                    v110 := GenS4{F0: 240}
                                    var v111 int64 = 1
                                    switch "\n\t" {
                                    case "a":
                                        switch int(v13) {
                                        case 9223372036854775807:
                                            var v112 uint = 18446744073709551615
                                            _ = v112
                                        }
                                    default:
                                        var v113 uint16 = 0
                                        _ = v113
                                    }
                                    fmt.Println(v107, fi103)
                                    _ = v110
                                    _ = v111
                                }
                                if false {
                                    v9 += 42
                                    if (-321 != 145) {
                                        fmt.Println(fi103, v14)
                                        var v114 int64 = -897
                                        _ = v114
                                    } else {
                                        var v115 float32 = v12
                                        v116 := GenS5{F0: v10}
                                        _ = len(v4)
                                        _ = v115
                                        _ = v116
                                    }
                                }
                                _ = fi106
                                _ = v107
                            }
                            v7 = v14
                            v14 += v7
                        }
                        for fi117 := 0; fi117 < 12; fi117++ {
                            v118 := []int{42, -441, (-(fi117))}
                            _ = fi117
                            _ = v118
                        }
                        _ = fi103
                    }
                    _ = v102
                }
            }
            _ = v87
        }
    }
    v13 = (v13 / ((v13) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
