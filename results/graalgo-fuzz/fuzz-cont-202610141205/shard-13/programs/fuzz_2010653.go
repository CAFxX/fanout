package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 bool; F2 uint8 }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 uint32; F1 uint8; F2 int }
type GenS5 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS6 struct { F0 bool; F1 uint32; F2 uint8 }
type GenS7 struct { F0 uint8; F1 uint32 }
type GenS8 struct { F0 bool; F1 uint16; F2 int64 }
type GenS9 struct { F0 uint32; F1 int32 }
type GenS10 struct { F0 int8; F1 uint }
type GenS11 struct { F0 uint; F1 string; F2 uint32 }
func genFn0(p0 bool, p1 uint8, p2 uint16) int32 {
    if (p2 <= 576) {
        var v6 uint8 = p1
        _ = v6
    }
    if p0 {
        v7 := GenS1{F0: (3.14159 < float32(0.0))}
        p1 = (^(p1))
        p0 = ((!true) || (0 > 42))
        _ = v7
    }
    var v8 int16 = 42
    p2 = (p2 ^ 255)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return -1
}
func genFn1() int16 {
    var v9 uint64 = 2
    v10 := []int{1, 2147483647}
    _ = v10
    _ = v9
    return -32768
}

func main() {
    v1 := make(map[int]int)
    var v2 int16 = 32767
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]int)
    var v5 uint = 148
    for fi11 := 0; fi11 < 7; fi11++ {
        v5 = (v5 << ((v5) & 63))
        v2 = 1000
        _ = fi11
    }
    if ((false && false) && false) {
        var v12 uint8 = 255
        if (42 < (v5 * v5)) {
            var v13 string = (("foo bar" + "a") + ("foo bar" + "hello"))
            _ = len(v3)
            v12 = 42
            _ = v13
        } else {
            var v14 uint8 = v12
            _ = v14
        }
        var v15 string = ""
        _ = v12
        _ = v15
    }
    v2 += 0
    for fi16 := 0; fi16 < 10; fi16++ {
        v5 += (v5 - 1)
        _ = fi16
    }
    v2 = genFn1()
    var v17 uint = v5
    var v18 int = 9223372036854775807
    if false {
        if (!(2 <= 777)) {
            switch (v18 * (^(v18))) {
            case 32767:
                v19 := v4[0]
                _ = v19
            case -1:
                if (true || true) {
                    if false {
                        if true {
                            var v20 int64 = -9223372036854775808
                            if (-1 != 42) {
                                v21 := GenS1{F0: false}
                                _ = v21
                            }
                            var v22 int16 = (v2 / (((^(v2))) | 1))
                            _ = v20
                            _ = v22
                        }
                        var v23 uint8 = 255
                        _ = v23
                    } else {
                        for fi24 := 0; fi24 < 14; fi24++ {
                            if false {
                                for fi25 := 0; fi25 < 8; fi25++ {
                                    v18 -= fi25
                                    var v26 int64 = -536
                                    var v27 int64 = (v26 * v26)
                                    _ = fi25
                                    _ = v26
                                    _ = v27
                                }
                                switch ("hello" + "hello") {
                                case "foo bar":
                                    for fi28 := 0; fi28 < 20; fi28++ {
                                        switch ("hello" + (("a" + "\n\t") + ("hello" + "foo bar"))) {
                                        case "\n\t":
                                            v18 -= 4294967295
                                        case "foo bar":
                                            if true {
                                                v29 := []int{}
                                                switch fi24 {
                                                case 32768:
                                                    if true {
                                                        var v30 uint32 = 0
                                                        _ = v30
                                                    }
                                                case 32767:
                                                    if true {
                                                        var v31 int = (v18 >> (((^(v17))) & 63))
                                                        v17 -= v17
                                                        if (-281 < 777) {
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                var v32 uint32 = 2147483647
                                                                _ = v32
                                                            default:
                                                                fmt.Println(fi28, v2)
                                                            }
                                                            var v33 uint8 = 127
                                                            if (!((0.0 + 1e10) >= (((1.5 - (0.0 - ((0.0 - float64(((0.1 * (-2.25 * float32(1e10))) - (3.14159 + 0.0)))) - (float64(float32(-2.25)) * (((((0.0 * 0.0) / 1e10) * (3.14159 / 1.5)) / (float64((float32(1.5) / 1e10)) - (((-2.25 - 0.0) * 0.0) * (((1e10 - -2.25) * 3.14159) + -2.25)))) + 0.0))))) / 0.1) / 1.5))) {
                                                                switch ("a" + "foo bar") {
                                                                case "\n\t":
                                                                    fmt.Println(v18, v2, v5)
                                                                case "hello":
                                                                    var v34 int = 9223372036854775807
                                                                    _ = v34
                                                                }
                                                            }
                                                            fmt.Println(v31)
                                                            var v35 float32 = 1.5
                                                            _ = v33
                                                            _ = v35
                                                        } else {
                                                            var v36 int32 = 737
                                                            v37 := GenS1{F0: false}
                                                            v31 = (^(v31))
                                                            _ = v36
                                                            _ = v37
                                                        }
                                                        _ = v31
                                                    }
                                                case 127:
                                                    fmt.Println(fi28, fi28)
                                                }
                                                _ = v29
                                            }
                                        case "a":
                                            if true {
                                                var v38 int16 = 1000
                                                var v39 int8 = -128
                                                _ = v38
                                                _ = v39
                                            } else {
                                                v40 := GenS1{F0: ((!false) || (!(((0.0 - 1.5) + 3.14159) >= (0.0 / 3.14159))))}
                                                if false {
                                                    var v41 bool = true
                                                    v17 = 180
                                                    _ = v41
                                                }
                                                v42 := GenS1{F0: (!false)}
                                                _ = v40
                                                _ = v42
                                            }
                                        }
                                        _ = fi28
                                    }
                                case "\n\t":
                                    v43 := []int{(fi24 | 128), fi24, 32767}
                                    _ = v43
                                }
                                switch (fi24 + v18) {
                                case -980:
                                    v17 += v5
                                case -9223372036854775808:
                                    var v44 uint64 = 2
                                    _ = v44
                                default:
                                    if false {
                                        for fi45 := 0; fi45 < 7; fi45++ {
                                            v46 := GenS1{F0: (-1 > 0)}
                                            if (!((v2 % ((int16(fi45)) | 1)) <= int16(v5))) {
                                                if (398 == 764) {
                                                    var v47 string = ("a" + ("hello" + "\n\t"))
                                                    _ = len(v1)
                                                    var v48 int8 = 0
                                                    _ = v47
                                                    _ = v48
                                                } else {
                                                    v49 := []int{-1}
                                                    var v50 float64 = float64(float32(0.0))
                                                    _ = v49
                                                    _ = v50
                                                }
                                                v17 += 4294967295
                                                v5 -= v17
                                                v18 = (v18 + -446)
                                            } else {
                                                var v51 bool = (-93 < -82)
                                                v52 := GenS1{F0: (3.14159 < (((-2.25 * float32((3.14159 + 1e10))) / (((0.1 * float32(1e10)) + (0.1 - 0.1)) / 0.1)) - 3.14159))}
                                                _ = v51
                                                _ = v52
                                            }
                                            var v53 int16 = int16(v17)
                                            v54 := GenS1{F0: (792 != 42)}
                                            _ = fi45
                                            _ = v46
                                            _ = v53
                                            _ = v54
                                        }
                                        v55 := GenS2{F0: "a", F1: (false || true), F2: 245}
                                        for fi56 := 0; fi56 < 9; fi56++ {
                                            switch "" {
                                            case "hello":
                                                v57 := GenS3{F0: 9223372036854775807, F1: -128}
                                                _ = v57
                                            case "foo bar":
                                                switch "\n\t" {
                                                case "a":
                                                    if (!(2147483647 < 0)) {
                                                        v17 = v17
                                                        var v58 int64 = 1
                                                        _ = len(v4)
                                                        _ = v58
                                                    } else {
                                                        var v59 uint32 = 4294967295
                                                        v4[(v59 - (v59 % ((v59) | 1)))] = fi24
                                                        fmt.Println(fi24, v18)
                                                        _ = v59
                                                    }
                                                case "foo bar":
                                                    if true {
                                                        v60 := GenS4{F0: 544, F1: 167, F2: (fi24 | (v18 & 616))}
                                                        _ = v60
                                                    }
                                                case "\n\t":
                                                    var v61 uint8 = 248
                                                    _ = v61
                                                default:
                                                    switch "a" {
                                                    case "foo bar":
                                                        if ((!(false || (false || ((fi56 & (fi24 - (fi24 ^ fi24))) > 0)))) && true) {
                                                            var v62 uint32 = 2
                                                            var v63 float32 = ((float32(-2.25) * ((-2.25 * 3.14159) + 1.5)) / 1.0)
                                                            _ = v62
                                                            _ = v63
                                                        }
                                                    default:
                                                        v17 -= v5
                                                    }
                                                }
                                            case "":
                                                fmt.Println(v17, v18, v2)
                                            default:
                                                var v64 float64 = 3.14159
                                                _ = v64
                                            }
                                            v65 := GenS5{F0: int16(v5), F1: 0, F2: 864}
                                            _ = fi56
                                            _ = v65
                                        }
                                        _ = v55
                                    }
                                }
                            } else {
                                v2 += (v2 - 0)
                                v66 := GenS6{F0: false, F1: 0, F2: 36}
                                v67 := GenS7{F0: 255, F1: 575}
                                _ = v66
                                _ = v67
                            }
                            v18 = v18
                            _ = fi24
                        }
                        if true {
                            v68 := GenS6{F0: ((3.14159 * 1e10) <= float32((3.14159 * 3.14159))), F1: 0, F2: 69}
                            _ = v68
                        } else {
                            for fi69 := 0; fi69 < 12; fi69++ {
                                var v70 int32 = 42
                                var v71 uint = (v5 * 4294967295)
                                _ = fi69
                                _ = v70
                                _ = v71
                            }
                        }
                        v17 += (v17 << ((v17) & 63))
                    }
                } else {
                    for fi72 := 0; fi72 < 20; fi72++ {
                        v73 := v3[173]
                        _ = fi72
                        _ = v73
                    }
                }
            case 97:
                var v74 float64 = float64(3.14159)
                _ = v74
            default:
                v75 := GenS8{F0: true, F1: 42, F2: -1}
                _ = v75
            }
        }
        fmt.Println(v2, v17, v18)
        var v76 uint = v5
        _ = v76
    }
    switch (v18 * (v18 ^ (v18 << (((v17 >> ((0) & 63))) & 63)))) {
    case 843:
        v77 := []int{(v18 + (v18 / ((v18) | 1)))}
        _ = v77
    default:
        if (!(!(625 < 1))) {
            switch (^(v18)) {
            case -733:
                v18 -= int(v17)
            default:
                if false {
                    v5 += v5
                    switch v18 {
                    case 42:
                        if (!true) {
                            if false {
                                v18 += (v18 / ((v18) | 1))
                            } else {
                                if false {
                                    v78 := []int{v18, v18, (v18 * v18)}
                                    _ = v78
                                }
                            }
                            var v79 int64 = 0
                            _ = v79
                        }
                    case -684:
                        v80 := GenS9{F0: 2147483647, F1: 824}
                        _ = v80
                    default:
                        v81 := GenS10{F0: -19, F1: v17}
                        _ = v81
                    }
                    switch "hello" {
                    case "\n\t":
                        for fi82 := 0; fi82 < 15; fi82++ {
                            v5 = (v5 ^ v5)
                            v83 := GenS11{F0: (v5 << (((v5 << (((^(v17))) & 63))) & 63)), F1: "foo bar", F2: 2}
                            _ = fi82
                            _ = v83
                        }
                    case "":
                        var v84 int16 = (v2 >> (((v5 - (v17 >> ((v5) & 63)))) & 63))
                        _ = v84
                    }
                } else {
                    var v85 string = "foo bar"
                    v86 := []int{}
                    for fi87 := 0; fi87 < 10; fi87++ {
                        fmt.Println(v85, v2, v18)
                        _ = fi87
                    }
                    _ = v85
                    _ = v86
                }
            }
            if false {
                v5 += v17
            } else {
                for fi88 := 0; fi88 < 16; fi88++ {
                    fmt.Println(v17)
                    _ = fi88
                }
                if (1.5 < 0.0) {
                    var v89 bool = (!((("\n\t" + ("foo bar" + "foo bar")) <= "a") && false))
                    for fi90 := 0; fi90 < 19; fi90++ {
                        v5 -= (v5 ^ (v17 >> ((v5) & 63)))
                        fmt.Println(v2, fi90)
                        _ = fi90
                    }
                    for fi91 := 0; fi91 < 2; fi91++ {
                        var v92 int16 = v2
                        var v93 float32 = (3.14159 + 1.5)
                        _ = fi91
                        _ = v92
                        _ = v93
                    }
                    _ = v89
                } else {
                    v2 -= genFn1()
                    v94 := GenS6{F0: (((-2.25 / 1.0) / 0.1) == (1.5 / float32((-2.25 - 1e10)))), F1: 2147483647, F2: 255}
                    _ = v94
                }
            }
        }
    }
    v95 := []int{(v18 + (v18 & v18)), v18, (v18 ^ (^(v18))), (v18 / ((v18) | 1))}
    _ = len(v4)
    if (false || (true || (uint64(v5) > 1))) {
        var v96 int16 = (v2 % ((-724) | 1))
        var v97 uint8 = 1
        if (uint8(v5) < (v97 ^ v97)) {
            for fi98 := 0; fi98 < 13; fi98++ {
                v17 = (^(v17))
                _ = fi98
            }
            _ = len(v1)
            var v99 int = (v18 * -221)
            _ = v99
        } else {
            if true {
                for fi100 := 0; fi100 < 10; fi100++ {
                    v17 = 65535
                    v18 += (fi100 >> (((v17 + v17)) & 63))
                    _ = fi100
                }
                v101 := GenS2{F0: "\n\t", F1: (false && (v18 <= v18)), F2: v97}
                _ = v101
            } else {
                var v102 uint64 = uint64(v17)
                if (!(42 == 7)) {
                    fmt.Println(v18, v102, v17)
                    if false {
                        switch "hello" {
                        case "\n\t":
                            for fi103 := 0; fi103 < 6; fi103++ {
                                var v104 int16 = v96
                                _ = fi103
                                _ = v104
                            }
                        case "a":
                            for fi105 := 0; fi105 < 3; fi105++ {
                                var v106 string = "hello"
                                for fi107 := 0; fi107 < 5; fi107++ {
                                    var v108 int32 = -1
                                    var v109 float64 = 3.14159
                                    _ = fi107
                                    _ = v108
                                    _ = v109
                                }
                                _ = fi105
                                _ = v106
                            }
                        case "foo bar":
                            if false {
                                fmt.Println(v96)
                            } else {
                                v3[0] = (("" + "") + (("hello" + "hello") + "\n\t"))
                                var v110 string = (("" + ("" + "foo bar")) + "a")
                                _ = v110
                            }
                        }
                    } else {
                        switch v18 {
                        case 0:
                            delete(v3, 0)
                        }
                        switch ("a" + "a") {
                        case "hello":
                            _ = len(v1)
                        case "":
                            switch -1 {
                            case 128:
                                if len(v95) > 0 {
                                    v95[0] = v18
                                }
                            }
                        default:
                            for fi111 := 0; fi111 < 11; fi111++ {
                                fmt.Println(v17, v97)
                                var v112 uint8 = v97
                                v96 += v96
                                _ = fi111
                                _ = v112
                            }
                        }
                    }
                    var v113 uint32 = 2147483647
                    _ = v113
                }
                if false {
                    for fi114 := 0; fi114 < 4; fi114++ {
                        fmt.Println(v5, v2, v2)
                        v115 := GenS7{F0: v97, F1: 0}
                        _ = fi114
                        _ = v115
                    }
                    var v116 string = "\n\t"
                    _ = v116
                }
                _ = v102
            }
        }
        _ = v96
        _ = v97
    } else {
        for fi117 := 0; fi117 < 10; fi117++ {
            v2 = (-(v2))
            v118 := GenS9{F0: 65535, F1: -935}
            for fi119 := 0; fi119 < 4; fi119++ {
                v2 += 1000
                if false {
                    if len(v95) > 0 {
                        v95[0] = 9223372036854775807
                    }
                }
                var v120 float64 = 1e10
                _ = fi119
                _ = v120
            }
            _ = fi117
            _ = v118
        }
        for fi121 := 0; fi121 < 16; fi121++ {
            for fi122 := 0; fi122 < 8; fi122++ {
                var v123 int16 = v2
                _ = fi122
                _ = v123
            }
            _ = fi121
        }
    }
    fmt.Println(v5, v17, v18)
    fmt.Println(v17, v18)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v95:", len(v95), v95)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v95
}
