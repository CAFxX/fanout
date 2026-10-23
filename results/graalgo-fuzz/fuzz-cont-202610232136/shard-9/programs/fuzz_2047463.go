package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 uint16; F2 uint32 }
type GenS4 struct { F0 uint8; F1 int32; F2 int }
type GenS5 struct { F0 int; F1 bool }
type GenS6 struct { F0 string; F1 int16 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 bool; F1 uint16; F2 uint32 }
type GenS9 struct { F0 int16; F1 bool }
type GenS10 struct { F0 uint16; F1 int32; F2 uint }
func genFn0() string {
    v7 := GenS1{F0: "hello", F1: "", F2: 211}
    v8 := []int{760, 233, 4294967295}
    if false {
        var v9 int64 = 1
        if (!(-32768 < -885)) {
            fmt.Println(v9)
            var v10 int8 = 1
            _ = v10
        } else {
            switch "" {
            case "":
                var v11 int16 = -1
                _ = v11
            }
        }
        fmt.Println(v9)
        _ = v9
    }
    if len(v8) > 0 {
        v8[0] = 127
    }
    _ = v7
    _ = v8
    return "hello"
}
func genFn1(p0 int16, p1 int8, p2 int) int64 {
    for fi12 := 0; fi12 < 8; fi12++ {
        if true {
            var v13 float64 = 3.14159
            v14 := GenS2{F0: -2147483648}
            p2 = p2
            _ = v13
            _ = v14
        }
        _ = fi12
    }
    p2 = (p2 << ((65535) & 63))
    for fi15 := 0; fi15 < 18; fi15++ {
        p2 -= -246
        _ = fi15
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 uint64 = 2
    var v2 uint64 = v1
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]string)
    var v5 int16 = 1
    var v6 uint8 = 2
    v2 += v2
    v16 := GenS3{F0: 1, F1: uint16(v2), F2: 1}
    v5 = (v5 + (v5 + (v5 - v5)))
    switch -2147483648 {
    case 655:
        v6 -= v6
    default:
        switch ("\n\t" + ("foo bar" + ("\n\t" + "foo bar"))) {
        case "":
            var v17 uint32 = 2147483648
            _ = v17
        default:
            v18 := []int{32767, -2147483648, 2147483647}
            _ = v18
        }
    }
    var v19 uint64 = 0
    var v20 bool = (!true)
    for fi21 := 0; fi21 < 8; fi21++ {
        v22 := GenS4{F0: (v6 * (^(v6))), F1: 2147483647, F2: -9223372036854775808}
        v6 = 110
        for fi23 := 0; fi23 < 3; fi23++ {
            v24 := []int{(fi21 ^ 9223372036854775807), -290}
            switch (fi23 & (fi21 + 32768)) {
            case 4294967295:
                fmt.Println(v20, v1, v20)
            case -137:
                fmt.Println(fi21, fi21, v6)
            }
            _ = len(v24)
            _ = fi23
            _ = v24
        }
        _ = fi21
        _ = v22
    }
    v2 = (v19 + v2)
    for fi25 := 0; fi25 < 13; fi25++ {
        v19 = (v1 << ((2) & 63))
        v20 = v20
        v3[953] = 632
        _ = fi25
    }
    fmt.Println(v20, v1)
    for fi26 := 0; fi26 < 1; fi26++ {
        if v20 {
            if (1 <= 42) {
                v27 := GenS3{F0: 491, F1: 385, F2: 332}
                _ = v27
            }
        }
        _ = fi26
    }
    for fi28 := 0; fi28 < 19; fi28++ {
        v1 -= 834
        if v20 {
            v29 := GenS5{F0: -9223372036854775808, F1: (42 == 0)}
            _ = v29
        } else {
            for fi30 := 0; fi30 < 15; fi30++ {
                if (v20 || false) {
                    var v31 uint = uint(v1)
                    _ = v31
                }
                fmt.Println(v6, fi30)
                var v32 uint32 = 664
                _ = fi30
                _ = v32
            }
            if (1 == fi28) {
                for fi33 := 0; fi33 < 18; fi33++ {
                    v4[(-(fi33))] = "foo bar"
                    if true {
                        v34 := GenS6{F0: "foo bar", F1: v5}
                        _ = v34
                    } else {
                        var v35 bool = (1.5 < -2.25)
                        _ = v35
                    }
                    if v20 {
                        v36 := GenS1{F0: "foo bar", F1: "foo bar", F2: 255}
                        switch -755 {
                        case -9223372036854775808:
                            if ((v5 - (v5 & v5)) == v5) {
                                var v37 string = "\n\t"
                                var v38 float32 = ((0.1 - 1e10) / ((((1.5 + 1e10) - 0.0) * (1.5 - float32(0.0))) * 3.14159))
                                v39 := []int{1, fi28}
                                _ = v37
                                _ = v38
                                _ = v39
                            } else {
                                switch "" {
                                case "":
                                    v2 = (v19 - (v2 / (((v2 * (^(v1)))) | 1)))
                                case "hello":
                                    var v40 int64 = 789
                                    _ = v40
                                case "foo bar":
                                    if (!v20) {
                                        if (!v20) {
                                            if ((^(v6)) >= (v6 - (v6 / ((255) | 1)))) {
                                                v19 += v19
                                                var v41 float64 = -2.25
                                                _ = v41
                                            } else {
                                                v42 := GenS1{F0: "hello", F1: "", F2: v6}
                                                v43 := []int{(fi28 | (fi28 >> ((uint(v6)) & 63))), 921, 128, (fi33 & (^(fi33)))}
                                                _ = v42
                                                _ = v43
                                            }
                                            v6 = v6
                                            v44 := GenS7{F0: (v5 >> ((65535) & 63))}
                                            _ = v44
                                        }
                                        v45 := GenS5{F0: (-(fi33)), F1: v20}
                                        _ = v45
                                    }
                                default:
                                    fmt.Println(fi28)
                                }
                            }
                        case -2147483648:
                            v46 := []int{0, int(v6), (^(fi28))}
                            _ = v46
                        default:
                            v1 = 42
                        }
                        _ = v36
                    } else {
                        v47 := []int{}
                        _ = v47
                    }
                    v5 += (v5 / ((v5) | 1))
                    fmt.Println(fi28, v6)
                    if (815 != 255) {
                        var v48 uint = 298
                        switch "hello" {
                        case "hello":
                            var v49 int32 = 0
                            _ = v49
                        }
                        _ = v48
                    } else {
                        v3[2] = 1
                        var v50 bool = ("\n\t" == "hello")
                        fmt.Println(v19, v6, v5)
                        _ = v50
                    }
                    delete(v4, -262)
                    v2 -= v2
                    v51 := []int{(fi33 ^ (^(fi33))), (fi28 * fi33)}
                    _ = fi33
                    _ = v51
                }
                switch "a" {
                case "":
                    if (true || ((fi28 >> ((18446744073709551615) & 63)) == fi28)) {
                        v52 := GenS2{F0: -757}
                        if (!((0 == 1) && v20)) {
                            for fi53 := 0; fi53 < 11; fi53++ {
                                v20 = (v19 > v19)
                                fmt.Println(v2)
                                _ = fi53
                            }
                            for fi54 := 0; fi54 < 13; fi54++ {
                                v55 := []int{(fi54 | 42), fi54, (fi28 / ((int(v6)) | 1))}
                                v1 = 860
                                _ = fi54
                                _ = v55
                            }
                            if v20 {
                                switch fi28 {
                                case 127:
                                    if v20 {
                                        var v56 uint16 = uint16(v2)
                                        _ = v56
                                    }
                                case -1:
                                    var v57 string = "hello"
                                    _ = v57
                                case 1:
                                    var v58 int64 = 9223372036854775807
                                    _ = v58
                                default:
                                    v59 := []int{-591, fi28}
                                    _ = v59
                                }
                                if (42 > -654) {
                                    fmt.Println(v20, v5, v6)
                                }
                                var v60 int32 = -1
                                _ = v60
                            }
                        }
                        _ = v52
                    } else {
                        switch "" {
                        case "a":
                            for fi61 := 0; fi61 < 9; fi61++ {
                                if v20 {
                                    if (0 < -79) {
                                        fmt.Println(v20, v20)
                                    } else {
                                        v2 -= v1
                                        var v62 uint = 0
                                        v1 = v19
                                        _ = v62
                                    }
                                    if false {
                                        v63 := []int{(fi61 ^ fi28), -2147483648, 32767, (fi28 - 1)}
                                        var v64 int = -628
                                        fmt.Println(v64, fi28, v20)
                                        _ = v63
                                        _ = v64
                                    }
                                } else {
                                    var v65 uint32 = 4294967295
                                    _ = v65
                                }
                                _ = fi61
                            }
                        case "\n\t":
                            var v66 uint = 44
                            _ = v66
                        default:
                            fmt.Println(v5, v1, fi28)
                        }
                        var v67 int16 = v5
                        var v68 bool = (!(0 > (v6 | (v6 ^ v6))))
                        _ = v67
                        _ = v68
                    }
                case "a":
                    if true {
                        switch fi28 {
                        case 1:
                            switch "hello" {
                            case "a":
                                var v69 string = (("" + "a") + ("a" + ("foo bar" + ("a" + "foo bar"))))
                                _ = v69
                            case "":
                                v20 = false
                            case "foo bar":
                                var v70 int32 = 42
                                _ = v70
                            default:
                                v71 := GenS5{F0: fi28, F1: (("" + ("foo bar" + ("\n\t" + "\n\t"))) > (("foo bar" + "") + "a"))}
                                _ = v71
                            }
                        case 32767:
                            var v72 uint32 = 985
                            _ = v72
                        }
                    } else {
                        if v20 {
                            v73 := v4[fi28]
                            _ = v73
                        }
                    }
                case "\n\t":
                    switch "a" {
                    case "a":
                        v74 := []int{(fi28 - (fi28 >> ((42) & 63))), (^(fi28)), (fi28 / (((fi28 >> ((18446744073709551615) & 63))) | 1)), 2147483647}
                        _ = v74
                    case "":
                        v6 = (v6 >> ((0) & 63))
                    case "hello":
                        v75 := GenS3{F0: 32767, F1: 0, F2: 4294967295}
                        _ = v75
                    }
                default:
                    var v76 float64 = (0.0 / -2.25)
                    _ = v76
                }
                switch (^(fi28)) {
                case 127:
                    if ((fi28 << ((1) & 63)) > fi28) {
                        v5 = (v5 >> ((255) & 63))
                        var v77 int64 = 42
                        var v78 uint = 18446744073709551615
                        _ = v77
                        _ = v78
                    } else {
                        for fi79 := 0; fi79 < 13; fi79++ {
                            v80 := v3[523]
                            _ = fi79
                            _ = v80
                        }
                    }
                default:
                    v81 := GenS8{F0: v20, F1: 42, F2: 2147483647}
                    _ = v81
                }
            }
        }
        if v20 {
            if v20 {
                var v82 uint16 = 65535
                var v83 float32 = ((((-2.25 * (0.1 * 1e10)) + 0.0) + 1.5) * 0.0)
                for fi84 := 0; fi84 < 8; fi84++ {
                    if (v20 && (!(!(1 <= 4294967295)))) {
                        var v85 uint = 629
                        _ = v85
                    } else {
                        v20 = v20
                        v86 := GenS5{F0: fi28, F1: ((true || v20) || false)}
                        _ = v86
                    }
                    v87 := v4[(fi28 * fi28)]
                    _ = fi84
                    _ = v87
                }
                _ = v82
                _ = v83
            } else {
                v88 := GenS5{F0: (fi28 - (fi28 * fi28)), F1: false}
                v19 -= v2
                _ = v88
            }
            var v89 string = ((("\n\t" + "foo bar") + ("a" + (("foo bar" + ("a" + ("foo bar" + "a"))) + "\n\t"))) + "hello")
            var v90 int32 = -1
            _ = v89
            _ = v90
        } else {
            for fi91 := 0; fi91 < 13; fi91++ {
                fmt.Println(v19)
                v6 = v6
                v92 := []int{744, 32768, fi91, fi28}
                _ = fi91
                _ = v92
            }
            switch (fi28 << ((821) & 63)) {
            case -306:
                var v93 float32 = float32(3.14159)
                _ = v93
            case 100:
                fmt.Println(v19)
            }
            for fi94 := 0; fi94 < 13; fi94++ {
                v20 = ((!true) || ((v6 >> ((4294967295) & 63)) >= 127))
                _ = fi94
            }
        }
        _ = fi28
    }
    var v95 uint8 = 0
    for fi96 := 0; fi96 < 18; fi96++ {
        fmt.Println(v5)
        _ = fi96
    }
    for fi97 := 0; fi97 < 12; fi97++ {
        if (!((((0.1 - -2.25) + -2.25) + (0.0 / -2.25)) != (0.0 - 3.14159))) {
            for fi98 := 0; fi98 < 5; fi98++ {
                var v99 uint64 = v1
                fmt.Println(v6)
                var v100 uint16 = 42
                delete(v4, 32767)
                _ = fi98
                _ = v100
                _ = v99
            }
            for fi101 := 0; fi101 < 18; fi101++ {
                v19 = (^(v1))
                v102 := GenS9{F0: 42, F1: (!v20)}
                v6 = (v95 & v6)
                v95 = (v6 + v6)
                _ = fi101
                _ = v102
            }
        }
        if (((^(v2)) < (v1 - v1)) && true) {
            v103 := []int{fi97, (-(fi97)), int(v5)}
            var v104 int8 = int8(v95)
            _ = v103
            _ = v104
        }
        v105 := []int{(-(fi97)), (fi97 >> ((799) & 63)), (fi97 / ((42) | 1))}
        _ = fi97
        _ = v105
    }
    if v20 {
        if (v20 || (!v20)) {
            fmt.Println(v95)
            if (v5 <= v5) {
                v6 += v95
                v1 += v1
            } else {
                if false {
                    v1 = 2
                } else {
                    v106 := []int{2147483647, 0, -1, -942}
                    if len(v106) >= 2 {
                        v106 = v106[1:len(v106)]
                    }
                    v107 := GenS10{F0: 2, F1: int32(v5), F2: 42}
                    _ = v106
                    _ = v107
                }
            }
            fmt.Println(v19, v95)
        }
        for fi108 := 0; fi108 < 16; fi108++ {
            delete(v3, 258)
            fmt.Println(v95)
            _ = fi108
        }
        v1 += (v1 | (v2 % (((v19 * v19)) | 1)))
    }
    if v20 {
        v2 = v19
        v5 = v5
        if ("" <= "hello") {
            v109 := []int{-1}
            _ = v109
        }
    } else {
        v110 := []int{-9223372036854775808, -2147483648, 720}
        _ = v110
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v95
}
