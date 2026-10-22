package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int64; F1 uint64 }
type GenS6 struct { F0 int32; F1 uint64 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int16; F1 uint64 }
type GenS10 struct { F0 int32; F1 uint8; F2 uint }
type GenS11 struct { F0 uint8; F1 int; F2 string }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 uint64; F1 int8; F2 uint64 }
type GenS14 struct { F0 uint; F1 int8; F2 uint32 }
type GenS15 struct { F0 string; F1 int; F2 int }
type GenS16 struct { F0 int16 }
type GenS17 struct { F0 int8; F1 bool; F2 string }
type GenS18 struct { F0 uint64; F1 string }
type GenS19 struct { F0 int; F1 int16; F2 uint32 }
type GenS20 struct { F0 uint8 }
func genFn0(p0 uint) uint16 {
    switch -9223372036854775808 {
    case 32767:
        var v5 int16 = 278
        _ = v5
    case 1:
        v6 := []int{}
        _ = v6
    }
    v7 := GenS1{F0: 1000}
    _ = p0
    _ = v7
    return 255
}
func genFn1() int64 {
    v8 := []int{1, 597, 32768, 1}
    v9 := GenS2{F0: (((float64(0.1) != (1e10 / (1.5 * 1e10))) || (384 >= 4294967295)) || (0 <= 0))}
    v10 := GenS3{F0: 1, F1: (("a" + "\n\t") + ("hello" + "a"))}
    if false {
        var v11 int16 = -32768
        for fi12 := 0; fi12 < 7; fi12++ {
            var v13 uint8 = 42
            v13 = v13
            _ = fi12
            _ = v13
        }
        _ = v11
    }
    if true {
        if (true && false) {
            v14 := GenS2{F0: (1.5 == ((0.0 + 1e10) - 1e10))}
            _ = v14
        }
    } else {
        var v15 int8 = -128
        _ = v15
    }
    _ = v10
    _ = v8
    _ = v9
    return 573
}

func main() {
    v1 := []int{1, 32767, 583}
    var v2 uint64 = 1
    v3 := []int{-767}
    var v4 uint = 18446744073709551615
    v16 := GenS1{F0: -1}
    if (true && (!false)) {
        var v17 uint16 = 566
        v4 += 1
        _ = v17
    } else {
        var v18 string = "hello"
        v19 := GenS4{F0: -1}
        _ = len(v3)
        _ = v18
        _ = v19
    }
    if true {
        if (42 == 554) {
            fmt.Println(v4)
            v4 -= (^(v4))
        }
    }
    var v20 int8 = -1
    for fi21 := 0; fi21 < 5; fi21++ {
        if ((2 > uint8(v2)) || (false && (((v4 ^ (v4 & v4)) == (v4 & (v4 % (((v4 >> (((v4 ^ v4)) & 63))) | 1)))) && true))) {
            v1 = append(v1, fi21)
        } else {
            v20 += v20
            if (-9223372036854775808 <= genFn1()) {
                switch fi21 {
                case 0:
                    fmt.Println(v2)
                default:
                    v22 := GenS5{F0: 948, F1: uint64(v4)}
                    _ = v22
                }
            }
            for fi23 := 0; fi23 < 2; fi23++ {
                if ((v20 & 100) <= int8(v2)) {
                    if (150 <= 0) {
                        if false {
                            v2 -= 42
                            switch "a" {
                            case "":
                                fmt.Println(fi23)
                                if ((-2.25 - -2.25) == float64(float32(1.5))) {
                                    if false {
                                        fmt.Println(v2, v20)
                                        if (-61 == genFn1()) {
                                            v24 := GenS6{F0: 1, F1: 1}
                                            var v25 int32 = 2147483647
                                            switch "\n\t" {
                                            case "\n\t":
                                                v25 = (v25 >> (((v4 & v4)) & 63))
                                            }
                                            _ = v24
                                            _ = v25
                                        }
                                    } else {
                                        var v26 bool = false
                                        _ = v26
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = fi23
                                    }
                                }
                            case "a":
                                v27 := GenS2{F0: ("foo bar" <= ("hello" + "a"))}
                                _ = v27
                            default:
                                v28 := GenS5{F0: -9223372036854775808, F1: v2}
                                _ = v28
                            }
                            if (!(0 > 255)) {
                                var v29 string = "foo bar"
                                switch (fi23 ^ fi23) {
                                case 0:
                                    v1 = append(v1, -448)
                                default:
                                    if false {
                                        var v30 uint8 = 127
                                        v4 += v4
                                        fmt.Println(v2, v2)
                                        _ = v30
                                    }
                                }
                                _ = v29
                            } else {
                                switch "a" {
                                case "\n\t":
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    v3 = append(v3, fi21)
                                }
                                v31 := GenS7{F0: (fi23 | fi21)}
                                _ = v31
                            }
                        } else {
                            var v32 uint16 = genFn0(255)
                            switch fi21 {
                            case 437:
                                var v33 uint16 = (v32 ^ (v32 % ((v32) | 1)))
                                _ = v33
                            case 32767:
                                v20 -= int8(v4)
                            case 508:
                                var v34 string = ""
                                _ = v34
                            default:
                                switch "" {
                                case "hello":
                                    switch fi23 {
                                    case 4294967295:
                                        v32 = (^(v32))
                                        if (!false) {
                                            v35 := GenS1{F0: 32767}
                                            _ = v35
                                        }
                                        if true {
                                            v4 += v4
                                            switch (-(fi21)) {
                                            case 178:
                                                v36 := GenS2{F0: true}
                                                _ = v36
                                            case 42:
                                                v32 = (v32 | (v32 - v32))
                                            case 1:
                                                v20 += (^(v20))
                                            default:
                                                var v37 uint16 = (v32 / ((v32) | 1))
                                                _ = v37
                                            }
                                        }
                                    case 2147483647:
                                        fmt.Println(v4)
                                    case 32768:
                                        v4 = v4
                                    default:
                                        if true {
                                            if (((3.14159 / 3.14159) + (((-2.25 + -2.25) / 1.5) / float32(((float64(1e10) * (0.0 / 1.0)) + (1e10 / (3.14159 * 0.1)))))) >= 0.0) {
                                                v20 -= 4
                                                if true {
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    v20 = -128
                                                } else {
                                                    if true {
                                                        _ = len(v1)
                                                        if (true && false) {
                                                            if (42 <= 2147483647) {
                                                                var v38 uint = 4294967295
                                                                _ = v38
                                                            }
                                                        }
                                                        v2 = (v2 | v2)
                                                    } else {
                                                        if true {
                                                            v39 := GenS8{F0: uint8(v32)}
                                                            _ = v39
                                                        } else {
                                                            v40 := GenS5{F0: 9223372036854775807, F1: uint64(v4)}
                                                            if (v32 < (^(v32))) {
                                                                v41 := GenS7{F0: fi23}
                                                                fmt.Println(v2, fi23)
                                                                v2 = (v2 - v2)
                                                                fmt.Println(fi23)
                                                                if true {
                                                                    if (v2 > (v2 * (v2 & v2))) {
                                                                        var v42 int8 = 1
                                                                        var v43 uint32 = 2
                                                                        v42 = v42
                                                                        v44 := GenS6{F0: 680, F1: v2}
                                                                        if len(v1) > 0 {
                                                                            v1[0] = (fi23 ^ (^(fi23)))
                                                                        }
                                                                        _ = v42
                                                                        _ = v43
                                                                        _ = v44
                                                                    }
                                                                    v4 = v4
                                                                } else {
                                                                    var v45 float64 = float64(0.1)
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                    fmt.Println(v2, v20, fi21)
                                                                    v20 += v20
                                                                    fmt.Println(v4, fi21, fi23)
                                                                    v46 := GenS9{F0: 1, F1: v2}
                                                                    _ = v45
                                                                    _ = v46
                                                                }
                                                                _ = v41
                                                            }
                                                            v47 := GenS10{F0: 1, F1: 127, F2: (v4 | v4)}
                                                            _ = v40
                                                            _ = v47
                                                        }
                                                    }
                                                    _ = len(v3)
                                                    var v48 int16 = -32768
                                                    _ = v48
                                                }
                                                v32 = 2
                                            }
                                            switch fi23 {
                                            case 128:
                                                if false {
                                                    var v49 float32 = 0.1
                                                    v32 = 1
                                                    if true {
                                                        v1 = append(v1, (fi21 - fi23))
                                                        var v50 int = 32767
                                                        var v51 uint32 = 566
                                                        _ = v50
                                                        _ = v51
                                                    } else {
                                                        var v52 float64 = float64(v49)
                                                        var v53 float32 = v49
                                                        v49 -= (3.14159 + v49)
                                                        _ = v52
                                                        _ = v53
                                                    }
                                                    _ = v49
                                                } else {
                                                    v20 -= 1
                                                }
                                            case 599:
                                                var v54 float64 = (0.0 - float64(3.14159))
                                                _ = v54
                                            }
                                        } else {
                                            v55 := GenS11{F0: 188, F1: (-(fi23)), F2: "a"}
                                            v56 := GenS3{F0: fi21, F1: "hello"}
                                            _ = v55
                                            _ = v56
                                        }
                                    }
                                default:
                                    _ = len(v1)
                                }
                            }
                            _ = v32
                        }
                        v20 -= v20
                        var v57 int16 = 1000
                        _ = v57
                    } else {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v58 int32 = 42
                        _ = v58
                    }
                    if len(v1) > 0 {
                        v1[0] = (fi21 - (fi23 + (fi23 << ((0) & 63))))
                    }
                    v20 += (v20 & 0)
                } else {
                    var v59 int32 = int32(v4)
                    var v60 int64 = 9223372036854775807
                    _ = v59
                    _ = v60
                }
                switch "\n\t" {
                case "a":
                    v4 = (v4 >> (((v4 / (((v4 - (v4 + (v4 - 0)))) | 1))) & 63))
                case "\n\t":
                    var v61 uint64 = v2
                    _ = v61
                }
                if true {
                    if true {
                        v3 = append(v3, fi21)
                        fmt.Println(fi21, v20, v20)
                    }
                    v4 = v4
                } else {
                    var v62 string = "foo bar"
                    _ = v62
                }
                if (255 <= 128) {
                    var v63 uint32 = 42
                    var v64 uint = 871
                    var v65 string = ""
                    v63 = (^(v63))
                    v66 := GenS12{F0: 65535}
                    _ = v63
                    _ = v64
                    _ = v65
                    _ = v66
                }
                _ = fi23
            }
        }
        _ = fi21
    }
    var v67 int8 = v20
    var v68 uint64 = (v2 / (((v2 & (v2 ^ v2))) | 1))
    for fi69 := 0; fi69 < 17; fi69++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if true {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            switch (fi69 * fi69) {
            case 0:
                v67 = v67
            case 128:
                var v70 int16 = 32767
                _ = v70
            case 894:
                switch fi69 {
                case 339:
                    for fi71 := 0; fi71 < 8; fi71++ {
                        fmt.Println(v20, v20, fi69)
                        fmt.Println(fi69)
                        _ = len(v3)
                        _ = fi71
                    }
                case 933:
                    var v72 uint32 = 2147483647
                    _ = v72
                case -586:
                    v73 := GenS13{F0: uint64(v4), F1: -128, F2: v68}
                    _ = v73
                default:
                    if (!false) {
                        switch "\n\t" {
                        case "\n\t":
                            var v74 int16 = 0
                            _ = v74
                        }
                    } else {
                        for fi75 := 0; fi75 < 3; fi75++ {
                            v76 := GenS14{F0: v4, F1: v67, F2: 1}
                            v20 -= (^(v20))
                            _ = fi75
                            _ = v76
                        }
                        var v77 uint16 = 42
                        fmt.Println(v4)
                        _ = v77
                    }
                }
            default:
                switch fi69 {
                case -1:
                    _ = len(v1)
                case 52:
                    v1 = append(v1, (^(fi69)))
                case -274:
                    switch "hello" {
                    case "hello":
                        for fi78 := 0; fi78 < 6; fi78++ {
                            if len(v3) > 0 {
                                v3[0] = (^(fi78))
                            }
                            if len(v3) > 0 {
                                v3[0] = (^(fi69))
                            }
                            if (151 >= 127) {
                                if len(v1) > 0 {
                                    v1[0] = (fi69 >> (((^(v4))) & 63))
                                }
                            }
                            _ = fi78
                        }
                    default:
                        for fi79 := 0; fi79 < 19; fi79++ {
                            v80 := GenS15{F0: "foo bar", F1: (fi79 & 433), F2: fi69}
                            if false {
                                var v81 int16 = -5
                                if false {
                                    fmt.Println(v20, v68)
                                }
                                v81 -= 42
                                _ = v81
                            }
                            _ = fi79
                            _ = v80
                        }
                    }
                default:
                    v1 = append(v1, fi69)
                }
            }
            var v82 int32 = 1
            _ = v82
        } else {
            for fi83 := 0; fi83 < 13; fi83++ {
                v4 = 1
                v84 := GenS5{F0: 247, F1: v2}
                fmt.Println(v20, v4)
                _ = fi83
                _ = v84
            }
        }
        _ = fi69
    }
    switch ("\n\t" + "\n\t") {
    case "foo bar":
        var v85 uint64 = v2
        _ = v85
    case "":
        v3 = append(v3, 964)
    default:
        switch "hello" {
        case "":
            for fi86 := 0; fi86 < 3; fi86++ {
                v1 = append(v1, (^(fi86)))
                _ = fi86
            }
        case "\n\t":
            for fi87 := 0; fi87 < 1; fi87++ {
                v88 := GenS16{F0: -1}
                v2 = uint64(v4)
                _ = fi87
                _ = v88
            }
        case "hello":
            v2 = v2
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi89 := 0; fi89 < 18; fi89++ {
        for fi90 := 0; fi90 < 5; fi90++ {
            var v91 bool = (false && true)
            v92 := GenS8{F0: 235}
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            switch fi90 {
            case 4294967295:
                var v93 bool = true
                _ = v93
            case 411:
                fmt.Println(v91, v2)
            case 42:
                if (false && v91) {
                    v91 = true
                    v94 := GenS17{F0: v20, F1: true, F2: ""}
                    _ = v94
                } else {
                    v4 = 2
                    var v95 bool = (!(42 < (v2 ^ 1)))
                    v4 -= (v4 * (^(v4)))
                    _ = v95
                }
            default:
                var v96 int = 32767
                _ = v96
            }
            if ((-2.25 - ((float64(1e10) / -2.25) / float64(float32(3.14159)))) == 3.14159) {
                if v91 {
                    _ = len(v1)
                } else {
                    v67 = (v67 ^ (-(v20)))
                    _ = len(v3)
                    if v91 {
                        var v97 bool = ((0.0 - 1.5) < 0.1)
                        if v91 {
                            v98 := GenS18{F0: 0, F1: ""}
                            switch fi90 {
                            case 128:
                                if v91 {
                                    switch 97 {
                                    case -477:
                                        v91 = (42 > 42)
                                    case 595:
                                        v20 = (v20 >> (((v4 ^ 0)) & 63))
                                    case -9223372036854775808:
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    default:
                                        if true {
                                            fmt.Println(v2, v91)
                                        }
                                    }
                                } else {
                                    if (!(42 < 531)) {
                                        v91 = true
                                        fmt.Println(v91, v2)
                                    } else {
                                        v67 -= (v67 & v20)
                                        v3 = append(v3, (fi89 >> ((uint(v68)) & 63)))
                                    }
                                }
                            case 32767:
                                if (3.14159 < (1.5 - 1.5)) {
                                    v20 += v20
                                    var v99 int64 = -362
                                    v91 = (-1 != -620)
                                    _ = v99
                                }
                            case -1:
                                if len(v1) > 0 {
                                    v1[0] = -947
                                }
                            }
                            fmt.Println(v67, v91)
                            v97 = v91
                            fmt.Println(v67, fi89)
                            if false {
                                v20 -= (v20 >> ((v4) & 63))
                            }
                            switch fi89 {
                            case -2147483648:
                                fmt.Println(v4, v4, v20)
                            }
                            _ = v98
                        }
                        _ = v97
                    } else {
                        v91 = ((1.5 / (-2.25 / 1.5)) >= ((((0.0 + (((((0.0 - 3.14159) - -2.25) / float64(3.14159)) * -2.25) / 1e10)) * 1e10) + 3.14159) - (0.0 - -2.25)))
                        var v100 float64 = float64(float32(1.5))
                        _ = v100
                    }
                }
                v4 = v4
            }
            _ = fi90
            _ = v91
            _ = v92
        }
        v101 := GenS19{F0: (fi89 << ((v4) & 63)), F1: 737, F2: 238}
        if false {
            switch 846 {
            case 127:
                v102 := GenS9{F0: -995, F1: v68}
                _ = v102
            }
            if (468 != -492) {
                fmt.Println(fi89, v20, v68)
                var v103 bool = (-1 <= -985)
                _ = v103
            } else {
                v104 := GenS20{F0: 127}
                _ = v104
            }
            if len(v3) > 0 {
                v3[0] = fi89
            }
        }
        _ = fi89
        _ = v101
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v67
    _ = v68
}
