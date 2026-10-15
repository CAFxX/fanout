package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 int64 }
type GenS2 struct { F0 string; F1 int; F2 int32 }
type GenS3 struct { F0 uint32; F1 string; F2 string }
type GenS4 struct { F0 uint8; F1 uint8; F2 uint }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int32; F1 uint64; F2 bool }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 string; F1 uint; F2 string }
type GenS9 struct { F0 int32; F1 uint8; F2 uint16 }
type GenS10 struct { F0 uint32; F1 uint32 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int32; F1 uint16; F2 int64 }
type GenS13 struct { F0 int64; F1 int }
type GenS14 struct { F0 uint64; F1 uint32 }
type GenS15 struct { F0 bool; F1 int32; F2 uint8 }
type GenS16 struct { F0 uint64 }
func genFn0() int8 {
    var v5 float64 = -2.25
    var v6 int = 9223372036854775807
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := []int{}
        var v9 uint64 = 0
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = v5
    _ = v6
    return -83
}

func main() {
    var v1 int8 = -128
    var v2 uint64 = 648
    var v3 uint32 = 42
    var v4 float64 = float64((0.1 - -2.25))
    v10 := []int{4294967295}
    var v11 float32 = ((0.1 / 1e10) * (-2.25 + (float32((float64(float32(v4)) - (0.1 + 1.5))) + (1e10 + -2.25))))
    v3 = 1
    if false {
        _ = len(v10)
    }
    v4 -= ((v4 + (v4 - v4)) * v4)
    for fi12 := 0; fi12 < 20; fi12++ {
        if ("hello" == ("" + "hello")) {
            v13 := GenS1{F0: v1, F1: -32768, F2: -1}
            _ = v13
        } else {
            switch (fi12 | int(v2)) {
            case -9223372036854775808:
                var v14 bool = false
                _ = v14
            case 428:
                var v15 int32 = 0
                _ = v15
            default:
                v11 += float32(3.14159)
            }
            if false {
                v16 := GenS1{F0: genFn0(), F1: -399, F2: 243}
                var v17 float64 = v4
                fmt.Println(v11, v11)
                _ = v16
                _ = v17
            } else {
                v18 := GenS2{F0: ("a" + "foo bar"), F1: fi12, F2: 1}
                v19 := GenS3{F0: 65535, F1: (("hello" + "hello") + ((("\n\t" + "\n\t") + ((("a" + ((("hello" + "") + "a") + "")) + "a") + "hello")) + (("a" + "") + "a"))), F2: "foo bar"}
                v1 = int8(v3)
                _ = v18
                _ = v19
            }
            var v20 uint16 = 65535
            _ = v20
        }
        for fi21 := 0; fi21 < 15; fi21++ {
            v22 := GenS1{F0: genFn0(), F1: -95, F2: 0}
            v23 := GenS4{F0: 226, F1: 127, F2: 4294967295}
            v1 -= (v1 & v1)
            v10 = append(v10, (-(fi12)))
            if ("\n\t" < "\n\t") {
                if true {
                    switch "\n\t" {
                    case "hello":
                        var v24 int8 = 42
                        _ = v24
                    case "":
                        if false {
                            var v25 int8 = genFn0()
                            var v26 float64 = 0.0
                            _ = v25
                            _ = v26
                        }
                    case "\n\t":
                        var v27 int = -761
                        _ = v27
                    default:
                        if true {
                            v28 := GenS3{F0: (^(v3)), F1: "hello", F2: "a"}
                            var v29 int8 = 7
                            v30 := GenS1{F0: (v1 + (v29 * v1)), F1: 966, F2: 388}
                            v1 -= 0
                            _ = v28
                            _ = v29
                            _ = v30
                        }
                    }
                    v2 += (v2 & v2)
                    var v31 float64 = float64(1e10)
                    _ = v31
                }
                if (42 <= 207) {
                    if false {
                        switch (^(fi21)) {
                        case 885:
                            if (-1 <= -357) {
                                if (1 == 32767) {
                                    var v32 uint16 = 0
                                    v32 = (v32 / (((^(v32))) | 1))
                                    _ = v32
                                } else {
                                    v3 = 94
                                    v4 = 0.1
                                    var v33 uint8 = 2
                                    var v34 int = (fi21 << ((2) & 63))
                                    fmt.Println(v11)
                                    _ = v33
                                    _ = v34
                                }
                                fmt.Println(fi21, fi12, v11)
                            } else {
                                v35 := GenS1{F0: -1, F1: 85, F2: 0}
                                var v36 int16 = -38
                                _ = v35
                                _ = v36
                            }
                        case 0:
                            if (true || false) {
                                if false {
                                    fmt.Println(v11, v2, v1)
                                    if (true && true) {
                                        v37 := GenS5{F0: 629}
                                        if false {
                                            var v38 uint32 = (v3 * (v3 >> ((18446744073709551615) & 63)))
                                            v3 = (v38 % ((875) | 1))
                                            fmt.Println(fi12, v2)
                                            _ = v38
                                        } else {
                                            if len(v10) >= 2 {
                                                v10 = v10[1:len(v10)]
                                            }
                                            switch (("a" + "") + "hello") {
                                            case "hello":
                                                v10 = append(v10, 127)
                                                v4 += v4
                                            default:
                                                if (1000 == 32767) {
                                                    switch (fi12 * 32768) {
                                                    case -1:
                                                        v11 = -2.25
                                                    }
                                                    var v39 string = "\n\t"
                                                    _ = v39
                                                }
                                            }
                                            var v40 bool = ((v2 << ((249) & 63)) <= (^(v2)))
                                            var v41 uint8 = 0
                                            v42 := GenS6{F0: -1, F1: v2, F2: v40}
                                            v3 = v3
                                            _ = v40
                                            _ = v41
                                            _ = v42
                                        }
                                        _ = v37
                                    }
                                } else {
                                    var v43 uint32 = (v3 << ((65535) & 63))
                                    _ = v43
                                }
                                v11 -= (v11 / float32((1e10 + -2.25)))
                                switch "hello" {
                                case "":
                                    v3 = (v3 << ((1) & 63))
                                case "foo bar":
                                    var v44 int16 = 0
                                    _ = v44
                                case "hello":
                                    v45 := GenS7{F0: (v1 >> ((65535) & 63))}
                                    _ = v45
                                default:
                                    fmt.Println(fi12, v1, v2)
                                }
                            }
                        case 967:
                            fmt.Println(v4, v3, v3)
                            v2 -= (v2 >> ((255) & 63))
                            var v46 string = "\n\t"
                            _ = v46
                        default:
                            v47 := GenS3{F0: v3, F1: ((((("" + "") + "hello") + ((("foo bar" + "hello") + "foo bar") + "hello")) + ("a" + ("" + ""))) + ""), F2: (("a" + "a") + (("foo bar" + "a") + "a"))}
                            _ = v47
                        }
                    }
                    if (-32768 < 1000) {
                        if true {
                            var v48 int64 = 42
                            v48 += 0
                            var v49 float32 = (3.14159 / 0.1)
                            if ((((1e10 - (v4 / 3.14159)) * (((v4 * 1e10) + v4) / float64(v11))) == (v4 * 0.0)) && true) {
                                if len(v10) >= 2 {
                                    v10 = v10[1:len(v10)]
                                }
                                switch 32767 {
                                case 2147483647:
                                    v3 = v3
                                case 127:
                                    v2 = uint64(v3)
                                case 4294967295:
                                    fmt.Println(v2, v2)
                                }
                                var v50 uint16 = 0
                                _ = v50
                            } else {
                                if len(v10) >= 2 {
                                    v10 = v10[1:len(v10)]
                                }
                                switch (^(fi21)) {
                                case 4294967295:
                                    switch (fi21 << ((0) & 63)) {
                                    case 42:
                                        v51 := GenS8{F0: "a", F1: 1, F2: ("foo bar" + ((("hello" + "foo bar") + "") + ("\n\t" + ("a" + "foo bar"))))}
                                        _ = v51
                                    case 9223372036854775807:
                                        var v52 uint16 = 874
                                        _ = v52
                                    }
                                case 660:
                                    if (int(v3) == (fi12 >> ((65535) & 63))) {
                                        var v53 uint64 = (^(v2))
                                        _ = v53
                                    } else {
                                        v48 -= (-(v48))
                                        v54 := GenS9{F0: 989, F1: 1, F2: 65535}
                                        fmt.Println(v11)
                                        _ = v54
                                    }
                                }
                                v55 := GenS2{F0: "a", F1: (^(fi12)), F2: 476}
                                _ = v55
                            }
                            v10 = append(v10, -726)
                            _ = v48
                            _ = v49
                        } else {
                            v56 := GenS10{F0: (v3 >> ((319) & 63)), F1: v3}
                            v11 = v11
                            var v57 int8 = genFn0()
                            v10 = append(v10, (^(fi21)))
                            var v58 uint64 = (v2 + 4294967295)
                            v1 += v57
                            switch "" {
                            case "":
                                switch (fi12 ^ 42) {
                                case -978:
                                    if (!(364 <= uint(v2))) {
                                        var v59 string = "a"
                                        var v60 uint16 = 810
                                        if len(v10) >= 2 {
                                            v10 = v10[1:len(v10)]
                                        }
                                        _ = v59
                                        _ = v60
                                    } else {
                                        fmt.Println(v2, v57, fi12)
                                    }
                                case 156:
                                    if (true || false) {
                                        switch fi12 {
                                        case 2147483647:
                                            v57 = 0
                                        default:
                                            v11 = 0.1
                                        }
                                        var v61 int32 = 2147483647
                                        fmt.Println(v58)
                                        _ = v61
                                    }
                                    fmt.Println(fi21)
                                case 4294967295:
                                    if (genFn0() == (v1 << ((18446744073709551615) & 63))) {
                                        var v62 int16 = 32767
                                        switch "\n\t" {
                                        case "hello":
                                            v63 := GenS8{F0: "a", F1: 0, F2: ("\n\t" + "a")}
                                            _ = v63
                                        }
                                        if (v3 < (v3 & v3)) {
                                            v64 := GenS11{F0: 0}
                                            fmt.Println(v1)
                                            var v65 bool = false
                                            _ = v64
                                            _ = v65
                                        } else {
                                            fmt.Println(v4)
                                            v58 = 18446744073709551615
                                            v66 := GenS12{F0: 42, F1: 26, F2: -9223372036854775808}
                                            v11 = v11
                                            switch 9223372036854775807 {
                                            case 128:
                                                var v67 string = ""
                                                _ = v67
                                            case -9223372036854775808:
                                                switch fi21 {
                                                case 2147483647:
                                                    var v68 float32 = v11
                                                    _ = v68
                                                case -9223372036854775808:
                                                    _ = len(v10)
                                                case 32768:
                                                    v4 += (v4 / 1.5)
                                                default:
                                                    var v69 int = fi21
                                                    _ = v69
                                                }
                                            case 1:
                                                if len(v10) > 0 {
                                                    v10[0] = fi12
                                                }
                                            default:
                                                var v70 string = "hello"
                                                _ = v70
                                            }
                                            _ = v66
                                        }
                                        _ = v62
                                    } else {
                                        if len(v10) >= 2 {
                                            v10 = v10[1:len(v10)]
                                        }
                                    }
                                default:
                                    if (!(42 != 18446744073709551615)) {
                                        v4 -= ((v4 + v4) * v4)
                                        var v71 float32 = ((v11 - (v11 * (((((v11 * (0.1 + 0.0)) * v11) + 1.5) / v11) - v11))) + v11)
                                        _ = v71
                                    } else {
                                        if false {
                                            if len(v10) > 0 {
                                                v10[0] = (fi12 >> ((42) & 63))
                                            }
                                            if len(v10) > 0 {
                                                v10[0] = (fi21 + fi12)
                                            }
                                            var v72 float64 = (1.5 + v4)
                                            if (-2147483648 >= 0) {
                                                if len(v10) > 0 {
                                                    v10[0] = (fi21 - fi12)
                                                }
                                                var v73 int16 = 42
                                                v74 := GenS13{F0: 386, F1: fi21}
                                                _ = v73
                                                _ = v74
                                            } else {
                                                var v75 int32 = 0
                                                _ = v75
                                            }
                                            _ = v72
                                        }
                                        fmt.Println(v4)
                                        fmt.Println(v57, fi12)
                                    }
                                }
                            case "foo bar":
                                var v76 float64 = 0.1
                                _ = v76
                            case "\n\t":
                                fmt.Println(v11)
                            }
                            switch (fi21 & fi12) {
                            case -1:
                                fmt.Println(fi12)
                                if ((v58 >> ((18446744073709551615) & 63)) != 42) {
                                    v77 := GenS14{F0: (v58 | v58), F1: v3}
                                    v78 := GenS14{F0: v2, F1: 564}
                                    v79 := GenS1{F0: v1, F1: 32767, F2: 819}
                                    var v80 int8 = v1
                                    if (-2147483648 == 2147483647) {
                                        var v81 uint64 = v58
                                        v82 := GenS15{F0: (0 > 42), F1: 0, F2: 2}
                                        if (v3 >= uint32(v2)) {
                                            v83 := GenS10{F0: v3, F1: (^(v3))}
                                            _ = v83
                                        }
                                        fmt.Println(v3, v80, fi21)
                                        v4 += float64(v11)
                                        _ = v81
                                        _ = v82
                                    }
                                    v10 = append(v10, fi21)
                                    var v84 uint64 = v2
                                    _ = v77
                                    _ = v78
                                    _ = v79
                                    _ = v80
                                    _ = v84
                                } else {
                                    v4 -= (1e10 + v4)
                                    v85 := GenS16{F0: v58}
                                    _ = len(v10)
                                    _ = v85
                                }
                            case 127:
                                if len(v10) >= 2 {
                                    v10 = v10[1:len(v10)]
                                }
                            }
                            _ = v56
                            _ = v57
                            _ = v58
                        }
                        var v86 float32 = 3.14159
                        _ = v86
                    } else {
                        var v87 uint64 = v2
                        fmt.Println(v4, v87, v1)
                        _ = v87
                    }
                    v4 = (1e10 * 0.1)
                } else {
                    fmt.Println(v3, v1, fi21)
                }
            } else {
                v4 = v4
            }
            v11 += v11
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = fi12
    }
    v4 -= 1.5
    if true {
        var v88 float32 = v11
        var v89 int = int(v3)
        _ = v88
        _ = v89
    } else {
        var v90 int = 32767
        _ = v90
    }
    for fi91 := 0; fi91 < 20; fi91++ {
        fmt.Println(fi91, v3, v2)
        if true {
            for fi92 := 0; fi92 < 1; fi92++ {
                switch fi91 {
                case 32768:
                    var v93 float32 = (((0.1 / 3.14159) - 0.1) + 1e10)
                    _ = v93
                default:
                    v3 += (v3 | 2)
                }
                var v94 int16 = 42
                _ = fi92
                _ = v94
            }
            var v95 int64 = 0
            var v96 uint = 360
            _ = v95
            _ = v96
        }
        if (18446744073709551615 < 841) {
            var v97 int64 = -218
            fmt.Println(v3, v97, v4)
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            _ = v97
        }
        _ = fi91
    }
    v3 = (^(v3))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
