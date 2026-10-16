package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 int8; F1 uint; F2 int64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int8; F1 int8; F2 int16 }
type GenS7 struct { F0 bool; F1 uint64; F2 int64 }
type GenS8 struct { F0 int64; F1 bool; F2 uint8 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 uint8; F1 uint; F2 int16 }
type GenS12 struct { F0 uint16; F1 int16 }
func genFn0() int64 {
    if true {
        var v4 int64 = 42
        _ = v4
    } else {
        v5 := GenS1{F0: 2147483647, F1: 8}
        var v6 uint16 = 65535
        _ = v5
        _ = v6
    }
    v7 := GenS2{F0: 42, F1: 2, F2: 42}
    v8 := GenS2{F0: 0, F1: 65535, F2: 42}
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 int, p1 uint8) bool {
    v9 := GenS1{F0: (p0 / ((p0) | 1)), F1: 2}
    v10 := []int{p0, -9223372036854775808, p0, -18}
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return (float32(1e10) >= 1.5)
}

func main() {
    var v1 uint = 2
    v2 := make(map[uint32]string)
    v3 := make(map[int]bool)
    if true {
        if (1e10 > 1e10) {
            var v11 int16 = 1000
            if (1 == int8(v1)) {
                v12 := []int{-109, -9223372036854775808}
                v11 = (v11 ^ (v11 / ((v11) | 1)))
                _ = v12
            }
            _ = v11
        } else {
            delete(v3, -686)
            var v13 bool = (!(!true))
            _ = v13
        }
    } else {
        v1 = (v1 % ((0) | 1))
    }
    var v14 uint16 = 255
    for fi15 := 0; fi15 < 19; fi15++ {
        for fi16 := 0; fi16 < 18; fi16++ {
            var v17 uint = v1
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    var v18 float64 = 1.5
    var v19 uint = 366
    for fi20 := 0; fi20 < 12; fi20++ {
        v21 := GenS1{F0: (fi20 * (-(fi20))), F1: v14}
        if true {
            fmt.Println(v1, v1, v14)
            fmt.Println(fi20, v19, fi20)
        }
        _ = fi20
        _ = v21
    }
    switch "" {
    case "\n\t":
        v22 := []int{632, 342, -9223372036854775808, 427}
        _ = v22
    }
    v23 := []int{-2147483648}
    if (!(0 == 0)) {
        v14 -= v14
        fmt.Println(v18)
        var v24 uint16 = v14
        _ = v24
    } else {
        fmt.Println(v1)
    }
    v23 = append(v23, -450)
    if false {
        v18 += ((1.5 * (((1e10 + v18) - v18) / (3.14159 + (v18 * (v18 / (1.5 - (((v18 / v18) + float64(1e10)) - v18))))))) * v18)
        var v25 uint = v19
        if true {
            v1 = v19
            var v26 bool = (!false)
            _ = v26
        } else {
            var v27 uint = v1
            var v28 uint8 = 1
            v29 := GenS2{F0: -1, F1: (^(v25)), F2: -1}
            _ = v27
            _ = v28
            _ = v29
        }
        _ = v25
    } else {
        var v30 int = 2147483647
        var v31 uint = 255
        v14 += 32767
        _ = v30
        _ = v31
    }
    if false {
        if true {
            v32 := GenS3{F0: 639}
            var v33 int32 = -493
            v33 -= 548
            _ = v32
            _ = v33
        }
        v23 = append(v23, 171)
    }
    fmt.Println(v19, v14)
    for fi34 := 0; fi34 < 3; fi34++ {
        var v35 uint = v19
        _ = fi34
        _ = v35
    }
    for fi36 := 0; fi36 < 19; fi36++ {
        for fi37 := 0; fi37 < 9; fi37++ {
            v14 = 65535
            _ = len(v2)
            _ = fi37
        }
        _ = fi36
    }
    fmt.Println(v1, v19, v14)
    switch "" {
    case "hello":
        fmt.Println(v18, v19)
    case "foo bar":
        if (true && true) {
            if (42 <= 0) {
                switch 128 {
                case -9223372036854775808:
                    v38 := GenS4{F0: 200}
                    _ = v38
                default:
                    v39 := GenS5{F0: v1}
                    _ = v39
                }
                var v40 uint8 = 1
                _ = v40
            }
            v41 := GenS1{F0: -176, F1: 2}
            _ = v41
        }
    }
    switch (("" + "") + "a") {
    case "\n\t":
        if true {
            v18 += v18
            for fi42 := 0; fi42 < 14; fi42++ {
                for fi43 := 0; fi43 < 9; fi43++ {
                    var v44 uint8 = 2
                    _ = fi43
                    _ = v44
                }
                _ = fi42
            }
            for fi45 := 0; fi45 < 12; fi45++ {
                var v46 int16 = -32768
                v47 := GenS4{F0: 0}
                _ = fi45
                _ = v46
                _ = v47
            }
        }
    case "":
        for fi48 := 0; fi48 < 6; fi48++ {
            v49 := GenS3{F0: -1}
            if ((^(fi48)) != -530) {
                v50 := GenS1{F0: (fi48 >> (((v1 << ((v1) & 63))) & 63)), F1: v14}
                _ = v50
            } else {
                v1 = v19
                v14 = v14
            }
            var v51 string = "\n\t"
            _ = fi48
            _ = v49
            _ = v51
        }
    case "a":
        if (true && ((!(!(!true))) && (127 < 255))) {
            if (((42 == 96) && (v14 <= 2)) && (false || ((0 > 42) || (!true)))) {
                if (42 < 0) {
                    v52 := GenS2{F0: 1, F1: uint(v14), F2: genFn0()}
                    v14 += (v14 * (^(v14)))
                    var v53 uint64 = 0
                    _ = v52
                    _ = v53
                } else {
                    if (!(true || (!((v18 - (float64(float32((v18 / 1.0))) * v18)) >= (1e10 - 1e10))))) {
                        if false {
                            v23 = append(v23, 4294967295)
                            v19 = uint(v14)
                            _ = len(v23)
                        } else {
                            for fi54 := 0; fi54 < 20; fi54++ {
                                for fi55 := 0; fi55 < 10; fi55++ {
                                    v19 -= 65535
                                    switch (-(fi54)) {
                                    case 32768:
                                        if true {
                                            var v56 float32 = 3.14159
                                            _ = v56
                                        }
                                        var v57 float64 = float64((0.1 - -2.25))
                                        if false {
                                            v58 := GenS6{F0: 87, F1: 93, F2: 622}
                                            _ = v58
                                        } else {
                                            v18 = 1.5
                                            var v59 string = "a"
                                            _ = v59
                                        }
                                        _ = v57
                                    default:
                                        if false {
                                            if ((("\n\t" + "") + "hello") <= ("a" + "")) {
                                                fmt.Println(v18, fi55)
                                                fmt.Println(v19)
                                                v19 -= v19
                                            } else {
                                                v60 := GenS3{F0: 42}
                                                fmt.Println(fi54, v1)
                                                _ = v60
                                            }
                                            var v61 int = 0
                                            _ = v61
                                        }
                                    }
                                    v1 = 4294967295
                                    fmt.Println(v19, fi55)
                                    _ = fi55
                                }
                                v18 = 3.14159
                                if ((42 != 42) || true) {
                                    v62 := GenS7{F0: false, F1: 848, F2: 1}
                                    var v63 string = "\n\t"
                                    _ = v62
                                    _ = v63
                                }
                                _ = fi54
                            }
                            if true {
                                switch "" {
                                case "":
                                    v19 += v19
                                case "\n\t":
                                    if true {
                                        v14 += v14
                                        var v64 string = "hello"
                                        _ = v64
                                    } else {
                                        v19 = 4294967295
                                    }
                                case "foo bar":
                                    v65 := GenS4{F0: 374}
                                    _ = v65
                                }
                                if (2 >= 4294967295) {
                                    v18 -= v18
                                    var v66 int8 = -24
                                    _ = v66
                                }
                            }
                        }
                        switch (("foo bar" + "foo bar") + "\n\t") {
                        case "\n\t":
                            if (false || (1e10 >= 0.0)) {
                                var v67 bool = ((201 > 2147483647) && false)
                                if v67 {
                                    if (-128 >= 53) {
                                        v67 = v67
                                        v68 := GenS4{F0: 42}
                                        switch "\n\t" {
                                        case "\n\t":
                                            var v69 float64 = v18
                                            _ = v69
                                        case "":
                                            v18 = v18
                                        case "a":
                                            v70 := GenS8{F0: 42, F1: ((v18 / float64((1e10 * 3.14159))) < (float64(3.14159) / 0.1)), F2: 63}
                                            _ = v70
                                        }
                                        _ = v68
                                    } else {
                                        v71 := GenS9{F0: (((0.0 - (((3.14159 / 0.1) - -2.25) - (-2.25 + 1.5))) / (((((0.1 + ((-2.25 / 1e10) * 1e10)) / ((float32(v18) + ((((1.5 - float32(v18)) - (0.1 + 1e10)) * 0.1) + (float32((float64(1e10) + v18)) + 1e10))) / ((-2.25 / 1.5) - 0.0))) / 1e10) + 1e10) * -2.25)) < 0.0)}
                                        v19 += v1
                                        for fi72 := 0; fi72 < 20; fi72++ {
                                            for fi73 := 0; fi73 < 12; fi73++ {
                                                v14 = v14
                                                switch ("" + (("\n\t" + (((("\n\t" + "foo bar") + "foo bar") + "foo bar") + "foo bar")) + ("\n\t" + "hello"))) {
                                                case "\n\t":
                                                    if (false || (3.14159 == float32((v18 - v18)))) {
                                                        fmt.Println(v1, v18)
                                                        v74 := GenS10{F0: 42}
                                                        _ = v74
                                                    }
                                                case "a":
                                                    v75 := v2[4294967295]
                                                    if ((-2.25 - 3.14159) != 3.14159) {
                                                        v75 = "\n\t"
                                                        if (!v67) {
                                                            v67 = v67
                                                            if len(v23) >= 2 {
                                                                v23 = v23[1:len(v23)]
                                                            }
                                                            var v76 int32 = 2147483647
                                                            _ = v76
                                                        } else {
                                                            var v77 uint32 = 4294967295
                                                            v19 = (v1 * (^(v19)))
                                                            _ = v77
                                                        }
                                                        v19 = (^(v1))
                                                        if len(v23) >= 2 {
                                                            v23 = v23[1:len(v23)]
                                                        }
                                                        var v78 uint8 = 255
                                                        _ = v78
                                                    } else {
                                                        var v79 float64 = v18
                                                        switch fi72 {
                                                        case -2147483648:
                                                            fmt.Println(v75, v75, fi73)
                                                        case 9223372036854775807:
                                                            var v80 uint16 = 2
                                                            _ = v80
                                                        }
                                                        v19 += (^(v19))
                                                        if ((!v67) || v67) {
                                                            fmt.Println(fi73)
                                                            v81 := GenS8{F0: 1, F1: false, F2: 255}
                                                            _ = v81
                                                        }
                                                        var v82 uint = (v1 << ((v1) & 63))
                                                        _ = v79
                                                        _ = v82
                                                    }
                                                    var v83 float64 = v18
                                                    _ = v75
                                                    _ = v83
                                                default:
                                                    fmt.Println(v1, v19, v19)
                                                }
                                                fmt.Println(v1, fi72, v18)
                                                _ = fi73
                                            }
                                            switch ("a" + (("" + "\n\t") + "hello")) {
                                            case "hello":
                                                fmt.Println(v19, v14, v14)
                                            default:
                                                switch (fi72 + (fi72 << (((v1 << ((0) & 63))) & 63))) {
                                                case 482:
                                                    if (genFn0() < 813) {
                                                        v67 = false
                                                        v84 := GenS8{F0: 42, F1: v67, F2: 236}
                                                        _ = v84
                                                    }
                                                case 128:
                                                    v85 := GenS10{F0: 42}
                                                    _ = v85
                                                case 9223372036854775807:
                                                    if len(v23) >= 2 {
                                                        v23 = v23[1:len(v23)]
                                                    }
                                                default:
                                                    var v86 int = (-(fi72))
                                                    _ = v86
                                                }
                                            }
                                            _ = fi72
                                        }
                                        _ = v71
                                    }
                                    v87 := GenS11{F0: 247, F1: (^(v19)), F2: 1}
                                    switch 42 {
                                    case 42:
                                        var v88 float32 = 0.1
                                        _ = v88
                                    }
                                    _ = v87
                                }
                                var v89 int8 = 79
                                _ = v67
                                _ = v89
                            }
                        default:
                            var v90 uint32 = 0
                            _ = v90
                        }
                        v18 = v18
                    } else {
                        v14 -= 2
                        switch 248 {
                        case 127:
                            fmt.Println(v1)
                        case 2147483647:
                            for fi91 := 0; fi91 < 7; fi91++ {
                                if false {
                                    if len(v23) > 0 {
                                        v23[0] = (fi91 << (((^(v19))) & 63))
                                    }
                                }
                                fmt.Println(v1, v18)
                                _ = fi91
                            }
                        }
                        if true {
                            fmt.Println(v14)
                            for fi92 := 0; fi92 < 18; fi92++ {
                                if false {
                                    for fi93 := 0; fi93 < 6; fi93++ {
                                        var v94 int64 = 42
                                        fmt.Println(v19, v18, v1)
                                        var v95 string = "foo bar"
                                        _ = fi93
                                        _ = v94
                                        _ = v95
                                    }
                                    for fi96 := 0; fi96 < 7; fi96++ {
                                        var v97 int = 42
                                        if len(v23) >= 2 {
                                            v23 = v23[1:len(v23)]
                                        }
                                        v1 += 65535
                                        _ = fi96
                                        _ = v97
                                    }
                                    v98 := GenS10{F0: 255}
                                    _ = v98
                                }
                                for fi99 := 0; fi99 < 19; fi99++ {
                                    v14 -= (v14 | (v14 % ((2) | 1)))
                                    var v100 int32 = 1
                                    _ = fi99
                                    _ = v100
                                }
                                v18 = float64(((3.14159 + float32(v18)) + 0.1))
                                _ = fi92
                            }
                        } else {
                            var v101 uint64 = 1
                            if true {
                                fmt.Println(v19)
                                if (int64(v14) != 9223372036854775807) {
                                    var v102 uint32 = 4294967295
                                    _ = v102
                                } else {
                                    var v103 uint32 = 1
                                    _ = v103
                                }
                            } else {
                                var v104 int16 = 42
                                fmt.Println(v14, v101)
                                _ = v104
                            }
                            v105 := GenS12{F0: 1, F1: -32768}
                            _ = v101
                            _ = v105
                        }
                    }
                }
                var v106 string = ("hello" + ("" + "a"))
                _ = v106
            } else {
                for fi107 := 0; fi107 < 8; fi107++ {
                    v18 = ((v18 - (v18 * (3.14159 - (v18 / ((float64(-2.25) * (3.14159 / v18)) * (3.14159 - (1.5 - 0.1))))))) * (v18 + v18))
                    _ = fi107
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v3
}
