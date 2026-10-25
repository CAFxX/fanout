package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 uint8; F1 uint16; F2 uint32 }
type GenS5 struct { F0 string; F1 int64; F2 string }
type GenS6 struct { F0 bool; F1 string; F2 int64 }

func main() {
    var v1 int32 = -1
    v2 := make(map[uint32]uint32)
    var v3 float32 = ((float32((float64(1e10) * (-2.25 * 3.14159))) * (3.14159 + (3.14159 + float32(1e10)))) * (float32(1e10) + 1.5))
    var v4 float32 = (((v3 / v3) + float32(0.1)) * v3)
    v4 = v3
    var v5 int64 = 1
    var v6 int64 = (v5 + (v5 % ((373) | 1)))
    fmt.Println(v1, v1)
    v1 = 2147483647
    var v7 string = "a"
    switch 2147483647 {
    case 32767:
        v1 += (v1 ^ -2147483648)
    }
    var v8 uint64 = 1
    fmt.Println(v7)
    var v9 int = 2147483647
    var v10 int64 = (^(v6))
    v3 -= (0.0 / v4)
    v11 := GenS1{F0: 255, F1: v9}
    var v12 int = 1
    switch "foo bar" {
    case "\n\t":
        for fi13 := 0; fi13 < 10; fi13++ {
            if (!false) {
                v7 = "\n\t"
            }
            var v14 uint8 = 128
            switch v9 {
            case 128:
                v15 := v2[878]
                _ = v15
            case 9223372036854775807:
                fmt.Println(v3, v7)
            }
            _ = fi13
            _ = v14
        }
    case "foo bar":
        v2[918] = 4294967295
    default:
        if (true && (((v7 + ("foo bar" + v7)) + v7) < (("a" + v7) + v7))) {
            switch v7 {
            case "foo bar":
                if (true && true) {
                    var v16 uint32 = 4294967295
                    for fi17 := 0; fi17 < 17; fi17++ {
                        if true {
                            v9 -= (fi17 % ((int(v10)) | 1))
                        }
                        v5 = -1
                        _ = fi17
                    }
                    if true {
                        fmt.Println(v5, v4)
                        v18 := v2[0]
                        for fi19 := 0; fi19 < 12; fi19++ {
                            var v20 int8 = -1
                            if (false && (true && ((false && (((float32((-2.25 - 0.1)) - (float32(0.1) / (v3 / (v4 - v3)))) == ((v4 * v3) * (float32((3.14159 / (0.1 / float64(v4)))) / ((v3 / (v4 + float32(3.14159))) - v3)))) && false)) && false))) {
                                switch (v7 + "foo bar") {
                                case "hello":
                                    v21 := GenS1{F0: 255, F1: int(v10)}
                                    _ = v21
                                case "":
                                    if true {
                                        switch v7 {
                                        case "\n\t":
                                            v22 := []int{v12, (-(v9)), (v9 % ((fi19) | 1))}
                                            _ = v22
                                        case "hello":
                                            var v23 uint = 2
                                            _ = v23
                                        default:
                                            var v24 bool = ((1.5 / -2.25) > -2.25)
                                            _ = v24
                                        }
                                    } else {
                                        for fi25 := 0; fi25 < 4; fi25++ {
                                            v18 += v18
                                            _ = fi25
                                        }
                                        if true {
                                            for fi26 := 0; fi26 < 6; fi26++ {
                                                var v27 bool = (!(0 >= 1))
                                                switch (((v7 + v7) + (v7 + ("foo bar" + ""))) + "hello") {
                                                case "":
                                                    v28 := GenS2{F0: (v1 << ((2) & 63)), F1: (!(895 <= 32767))}
                                                    _ = v28
                                                case "foo bar":
                                                    var v29 int = fi19
                                                    _ = v29
                                                }
                                                if false {
                                                    fmt.Println(v3, v27, v16)
                                                } else {
                                                    v20 = -128
                                                    fmt.Println(v6, v16)
                                                }
                                                _ = fi26
                                                _ = v27
                                            }
                                            v3 = (1.5 + 0.0)
                                        } else {
                                            v3 = v4
                                            for fi30 := 0; fi30 < 17; fi30++ {
                                                v31 := GenS1{F0: 4294967295, F1: (v12 << ((486) & 63))}
                                                fmt.Println(v18)
                                                _ = fi30
                                                _ = v31
                                            }
                                            for fi32 := 0; fi32 < 3; fi32++ {
                                                var v33 float64 = (((1e10 - (((0.1 - 0.0) / 3.14159) * 1e10)) - -2.25) * float64(v3))
                                                var v34 uint32 = v18
                                                if false {
                                                    v1 -= 413
                                                    v35 := GenS3{F0: 2, F1: 0}
                                                    v5 -= v10
                                                    _ = v35
                                                }
                                                fmt.Println(v8, v20, v34)
                                                _ = fi32
                                                _ = v33
                                                _ = v34
                                            }
                                        }
                                        for fi36 := 0; fi36 < 9; fi36++ {
                                            var v37 uint8 = 128
                                            fmt.Println(v1, v5, v4)
                                            fmt.Println(v3, v1, v20)
                                            _ = fi36
                                            _ = v37
                                        }
                                    }
                                default:
                                    var v38 uint8 = 23
                                    _ = v38
                                }
                            }
                            var v39 uint = 65535
                            _ = fi19
                            _ = v20
                            _ = v39
                        }
                        _ = v18
                    } else {
                        for fi40 := 0; fi40 < 17; fi40++ {
                            switch "foo bar" {
                            case "foo bar":
                                fmt.Println(v4, v12)
                            case "\n\t":
                                for fi41 := 0; fi41 < 2; fi41++ {
                                    var v42 uint32 = 4294967295
                                    _ = fi41
                                    _ = v42
                                }
                            default:
                                if (uint16(v8) >= uint16(v8)) {
                                    var v43 bool = (v8 >= v8)
                                    if (!v43) {
                                        delete(v2, (v16 * (^(v16))))
                                        for fi44 := 0; fi44 < 2; fi44++ {
                                            var v45 bool = v43
                                            _ = fi44
                                            _ = v45
                                        }
                                        v46 := []int{128}
                                        _ = v46
                                    } else {
                                        var v47 uint32 = v16
                                        var v48 uint64 = v8
                                        v49 := GenS4{F0: 1, F1: 0, F2: 596}
                                        _ = v47
                                        _ = v48
                                        _ = v49
                                    }
                                    _ = v43
                                } else {
                                    if (!true) {
                                        v5 = v5
                                    } else {
                                        v50 := GenS2{F0: (-(v1)), F1: true}
                                        _ = v50
                                    }
                                }
                            }
                            for fi51 := 0; fi51 < 6; fi51++ {
                                v4 -= 3.14159
                                fmt.Println(v6)
                                _ = fi51
                            }
                            _ = fi40
                        }
                        for fi52 := 0; fi52 < 12; fi52++ {
                            var v53 float64 = 0.1
                            fmt.Println(v7, v53)
                            _ = fi52
                            _ = v53
                        }
                        v5 = (v10 / (((v10 | (v10 & v6))) | 1))
                    }
                    _ = v16
                } else {
                    if true {
                        fmt.Println(v5, v4, v7)
                        if false {
                            v1 += 42
                            v54 := GenS4{F0: 128, F1: 2, F2: 42}
                            _ = v54
                        }
                    }
                    switch ("a" + (v7 + v7)) {
                    case "a":
                        var v55 uint8 = 1
                        _ = v55
                    case "foo bar":
                        fmt.Println(v3)
                    default:
                        fmt.Println(v1)
                    }
                    var v56 uint8 = 42
                    _ = v56
                }
            case "\n\t":
                var v57 int64 = v6
                _ = v57
            }
            if true {
                v7 = "hello"
                v2[2147483648] = 907
            } else {
                for fi58 := 0; fi58 < 18; fi58++ {
                    for fi59 := 0; fi59 < 20; fi59++ {
                        var v60 int64 = (v10 >> ((643) & 63))
                        _ = fi59
                        _ = v60
                    }
                    _ = fi58
                }
            }
        } else {
            var v61 float32 = v4
            if (((v61 / v4) * (v3 / float32(((1.5 + 3.14159) * ((1.5 / 1.0) - 1e10))))) <= -2.25) {
                fmt.Println(v8, v7)
                switch v12 {
                case 127:
                    switch "foo bar" {
                    case "foo bar":
                        switch (^(v9)) {
                        case 127:
                            var v62 float64 = 1.5
                            _ = v62
                        case 0:
                            for fi63 := 0; fi63 < 19; fi63++ {
                                var v64 uint32 = 42
                                _ = fi63
                                _ = v64
                            }
                        case -278:
                            var v65 int64 = 42
                            _ = v65
                        }
                    case "hello":
                        v66 := GenS4{F0: 23, F1: 65535, F2: 1}
                        _ = v66
                    case "a":
                        v2[42] = 48
                    default:
                        v8 = v8
                    }
                case -942:
                    v67 := GenS4{F0: 168, F1: 255, F2: 0}
                    _ = v67
                }
            } else {
                var v68 uint16 = 904
                var v69 int = 2147483647
                for fi70 := 0; fi70 < 18; fi70++ {
                    switch v7 {
                    case "foo bar":
                        var v71 string = "a"
                        _ = v71
                    case "hello":
                        fmt.Println(v1, fi70)
                    case "":
                        if true {
                            _ = len(v2)
                            if ((v1 | v1) > -1) {
                                if true {
                                    for fi72 := 0; fi72 < 16; fi72++ {
                                        v73 := v2[4294967295]
                                        var v74 uint = 18446744073709551615
                                        var v75 string = v7
                                        if (!(!true)) {
                                            v76 := []int{(fi72 | (^(v69))), (fi70 ^ v9)}
                                            fmt.Println(v10)
                                            switch (v12 >> (((v74 & v74)) & 63)) {
                                            case 9223372036854775807:
                                                var v77 string = v75
                                                _ = v77
                                            case -723:
                                                v78 := GenS2{F0: v1, F1: true}
                                                _ = v78
                                            case 32768:
                                                var v79 int64 = -9223372036854775808
                                                _ = v79
                                            default:
                                                if false {
                                                    if len(v76) > 0 {
                                                        v76[0] = (fi72 & v9)
                                                    }
                                                } else {
                                                    v73 -= v73
                                                }
                                            }
                                            var v80 int64 = 1
                                            _ = v76
                                            _ = v80
                                        }
                                        if (true && true) {
                                            v74 -= v74
                                        } else {
                                            v81 := GenS5{F0: v75, F1: 0, F2: "hello"}
                                            var v82 int8 = 42
                                            switch ("" + "\n\t") {
                                            case "\n\t":
                                                var v83 uint32 = 767
                                                _ = v83
                                            }
                                            _ = v81
                                            _ = v82
                                        }
                                        var v84 string = "\n\t"
                                        _ = fi72
                                        _ = v73
                                        _ = v74
                                        _ = v75
                                        _ = v84
                                    }
                                } else {
                                    switch v7 {
                                    case "\n\t":
                                        fmt.Println(v10)
                                    case "a":
                                        for fi85 := 0; fi85 < 5; fi85++ {
                                            var v86 uint32 = 65535
                                            fmt.Println(v5, v5)
                                            v87 := GenS5{F0: "hello", F1: v5, F2: ""}
                                            v88 := GenS6{F0: true, F1: v7, F2: (v6 / ((v5) | 1))}
                                            v89 := GenS5{F0: "a", F1: v5, F2: "\n\t"}
                                            _ = fi85
                                            _ = v86
                                            _ = v87
                                            _ = v88
                                            _ = v89
                                        }
                                    default:
                                        fmt.Println(v12, v10, v5)
                                    }
                                    v90 := v2[4294967295]
                                    _ = v90
                                }
                            } else {
                                v8 = (v8 << ((559) & 63))
                                if (0.1 > 1e10) {
                                    fmt.Println(v9, v3, v7)
                                    switch v12 {
                                    case 4294967295:
                                        for fi91 := 0; fi91 < 16; fi91++ {
                                            v92 := []int{}
                                            _ = fi91
                                            _ = v92
                                        }
                                    }
                                    var v93 bool = ((v5 & v10) == (v6 & (-(v5))))
                                    _ = v93
                                } else {
                                    if ((!true) || (true && (!(!(2147483647 <= 875))))) {
                                        fmt.Println(v5, v69)
                                        var v94 int32 = v1
                                        _ = v94
                                    }
                                    var v95 float32 = float32(3.14159)
                                    v96 := GenS5{F0: v7, F1: 409, F2: (v7 + (v7 + v7))}
                                    _ = v95
                                    _ = v96
                                }
                            }
                            for fi97 := 0; fi97 < 7; fi97++ {
                                v2[2] = 38
                                var v98 bool = true
                                var v99 uint = 65535
                                _ = fi97
                                _ = v98
                                _ = v99
                            }
                        } else {
                            v100 := []int{9223372036854775807, int(v6), (v69 + v12), v69}
                            _ = v100
                        }
                    default:
                        if (!((v5 & v6) != (v6 & v5))) {
                            if (!((128 > 49) && ((false || (242 >= 877)) || (false || true)))) {
                                if true {
                                    switch "hello" {
                                    case "":
                                        v6 -= v10
                                    case "hello":
                                        v8 -= v8
                                    }
                                } else {
                                    if (-37 == -128) {
                                        var v101 int32 = (v1 << ((0) & 63))
                                        v102 := v2[uint32(v8)]
                                        switch v69 {
                                        case 9223372036854775807:
                                            for fi103 := 0; fi103 < 9; fi103++ {
                                                v3 = v61
                                                _ = len(v2)
                                                _ = fi103
                                            }
                                        default:
                                            if (v1 != (v101 + (v101 << ((833) & 63)))) {
                                                v104 := []int{}
                                                _ = v104
                                            } else {
                                                for fi105 := 0; fi105 < 16; fi105++ {
                                                    fmt.Println(v102)
                                                    var v106 int64 = 9223372036854775807
                                                    v107 := []int{-826, (^(v12)), int(v68), (^(v12))}
                                                    _ = fi105
                                                    _ = v106
                                                    _ = v107
                                                }
                                                var v108 uint = 637
                                                _ = v108
                                            }
                                        }
                                        _ = v101
                                        _ = v102
                                    }
                                    v4 -= v3
                                    _ = len(v2)
                                }
                            } else {
                                v68 = (v68 ^ (v68 + v68))
                            }
                            v109 := []int{fi70, (fi70 / ((v12) | 1)), (v12 * v9), (v69 - (v9 & (^(v69))))}
                            _ = v109
                        } else {
                            v110 := v2[390]
                            for fi111 := 0; fi111 < 16; fi111++ {
                                var v112 uint = 18446744073709551615
                                fmt.Println(v10, fi70)
                                _ = fi111
                                _ = v112
                            }
                            if false {
                                for fi113 := 0; fi113 < 14; fi113++ {
                                    if (true && false) {
                                        v8 = (^(v8))
                                        if (v110 > (v110 * (v110 | (v110 | (v110 % (((v110 | v110)) | 1)))))) {
                                            v4 = 1e10
                                            var v114 int8 = -1
                                            v61 = v4
                                            _ = v114
                                        }
                                        var v115 int16 = 32767
                                        _ = v115
                                    }
                                    if false {
                                        delete(v2, 2147483647)
                                        var v116 uint8 = 146
                                        var v117 int8 = -128
                                        _ = v116
                                        _ = v117
                                    }
                                    _ = fi113
                                }
                            } else {
                                switch 2147483647 {
                                case -631:
                                    v1 = (v1 >> ((471) & 63))
                                case 4294967295:
                                    v4 -= v4
                                case 32768:
                                    for fi118 := 0; fi118 < 4; fi118++ {
                                        switch (v9 ^ (v9 + (v9 / ((fi118) | 1)))) {
                                        case -810:
                                            v5 -= (v6 + (-(v10)))
                                        case 4294967295:
                                            v69 = fi118
                                        case -111:
                                            var v119 uint64 = (v8 | (v8 ^ uint64(v110)))
                                            _ = v119
                                        default:
                                            var v120 int32 = (v1 & int32(fi118))
                                            _ = v120
                                        }
                                        _ = fi118
                                    }
                                }
                                var v121 float64 = float64(v61)
                                _ = v121
                            }
                            _ = v110
                        }
                    }
                    fmt.Println(v12, v61, v7)
                    var v122 uint8 = 255
                    _ = fi70
                    _ = v122
                }
                _ = v68
                _ = v69
            }
            v123 := []int{(v12 / ((32767) | 1)), v12, int(v8), (v12 | v9)}
            _ = v123
            _ = v61
        }
    }
    var v124 int16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v124:", v124)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v124
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
