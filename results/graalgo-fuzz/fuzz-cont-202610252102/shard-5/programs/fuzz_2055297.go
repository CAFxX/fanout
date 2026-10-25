package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint; F1 uint16; F2 uint }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int; F1 string; F2 uint32 }
type GenS7 struct { F0 int32; F1 int16 }
type GenS8 struct { F0 bool; F1 uint16 }
type GenS9 struct { F0 string; F1 bool }
type GenS10 struct { F0 int8; F1 int; F2 uint8 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 int64 }
func genFn0(p0 bool, p1 int64, p2 uint64) uint8 {
    var v7 bool = true
    p2 = 1
    var v8 string = ("" + "\n\t")
    v9 := GenS1{F0: -88, F1: p2}
    v10 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 int8) string {
    fmt.Println(p0)
    switch -1 {
    case -486:
        var v11 float64 = 1.5
        _ = v11
    case 305:
        fmt.Println(p0)
    case 542:
        if true {
            if true {
                var v12 float64 = (0.0 * 1.5)
                var v13 int8 = p0
                _ = v12
                _ = v13
            }
            v14 := []int{0, -1}
            for fi15 := 0; fi15 < 3; fi15++ {
                v14 = append(v14, fi15)
                var v16 bool = false
                _ = fi15
                _ = v16
            }
            _ = v14
        }
    }
    _ = p0
    return "hello"
}

func main() {
    var v1 int64 = 42
    var v2 float64 = 1e10
    var v3 int8 = 0
    v4 := make(map[uint32]uint32)
    var v5 string = "a"
    v6 := make(map[int]bool)
    v2 += (1.5 - v2)
    v17 := GenS1{F0: 4, F1: 711}
    if true {
        var v18 uint32 = 65535
        _ = v18
    } else {
        var v19 float32 = 1.5
        if true {
            v20 := []int{9223372036854775807, 0}
            fmt.Println(v19, v1)
            _ = v20
        } else {
            v5 = ("a" + ("foo bar" + v5))
            for fi21 := 0; fi21 < 4; fi21++ {
                for fi22 := 0; fi22 < 12; fi22++ {
                    v23 := GenS2{F0: (true && (true && (float32(v2) != 0.0))), F1: 0}
                    v19 = v19
                    if (float64(((float32(0.0) / 0.1) + (1e10 * (1e10 - ((v19 + v19) - 1.5))))) < 1e10) {
                        if true {
                            switch (-(fi22)) {
                            case 9223372036854775807:
                                if (true && false) {
                                    if (1 < 0) {
                                        if (true || (2 == 2)) {
                                            v24 := []int{fi22, (fi22 * fi21)}
                                            _ = v24
                                        } else {
                                            fmt.Println(fi22, v19, v1)
                                            var v25 uint = 2
                                            v26 := GenS3{F0: 42}
                                            _ = v25
                                            _ = v26
                                        }
                                    }
                                }
                            default:
                                if ((v1 - (v1 / ((v1) | 1))) != v1) {
                                    v2 += v2
                                    fmt.Println(v1)
                                }
                            }
                            v19 = float32(-2.25)
                        } else {
                            v2 = v2
                            if true {
                                if (v2 != ((float64((float32(v2) * float32(float64((v19 / 1.5))))) + -2.25) - 1e10)) {
                                    var v27 int64 = v1
                                    v28 := GenS3{F0: 65535}
                                    v29 := GenS3{F0: 42}
                                    _ = v27
                                    _ = v28
                                    _ = v29
                                } else {
                                    switch fi22 {
                                    case 42:
                                        if false {
                                            fmt.Println(fi22, v5, fi21)
                                            v19 = (float32((-2.25 / 1.0)) - v19)
                                        } else {
                                            v2 = v2
                                            var v30 int64 = v1
                                            _ = v30
                                        }
                                    default:
                                        var v31 uint = 65535
                                        _ = v31
                                    }
                                }
                                delete(v6, fi21)
                                if true {
                                    if ((!false) && (65535 <= 42)) {
                                        v32 := GenS4{F0: 0, F1: 2, F2: 2}
                                        v6[2147483647] = (true || (167 != 2147483648))
                                        v5 = v5
                                        _ = v32
                                    } else {
                                        v1 -= (v1 - -1)
                                        v5 = v5
                                        var v33 int32 = 381
                                        _ = v33
                                    }
                                    v3 = -127
                                }
                            }
                        }
                    } else {
                        var v34 int64 = 9223372036854775807
                        _ = len(v4)
                        switch v5 {
                        case "\n\t":
                            if (((!((1 == 42) || ((v1 << ((65535) & 63)) >= v34))) || (3.14159 > (0.0 / float64(((0.1 / v19) * (v19 * ((1e10 - (v19 * v19)) * 0.0))))))) && ((false || (4294967295 >= fi22)) || (!(764 != 2)))) {
                                delete(v4, 2147483647)
                                v35 := GenS4{F0: 65535, F1: 2, F2: 1}
                                if (244 <= 0) {
                                    v36 := GenS1{F0: (fi22 * (fi21 ^ fi22)), F1: 18446744073709551615}
                                    v37 := GenS5{F0: genFn0((!true), v34, 394)}
                                    switch 1 {
                                    case 0:
                                        v38 := GenS6{F0: fi22, F1: (v5 + v5), F2: 2147483648}
                                        _ = v38
                                    case -285:
                                        if (!true) {
                                            v19 = 0.0
                                            var v39 int8 = v3
                                            _ = v39
                                        } else {
                                            if true {
                                                fmt.Println(v2, fi22, v1)
                                            }
                                        }
                                    }
                                    _ = v36
                                    _ = v37
                                }
                                if ((((v19 * (v19 / 1.0)) * ((1.5 * float32(float64(v19))) + v19)) >= v19) || false) {
                                    v5 = "\n\t"
                                    var v40 float32 = v19
                                    var v41 int64 = (v1 + v1)
                                    _ = v40
                                    _ = v41
                                }
                                v34 = (v34 + 42)
                                _ = v35
                            }
                        case "hello":
                            switch (v5 + v5) {
                            case "\n\t":
                                v42 := []int{fi21, fi22, 2147483647, fi22}
                                _ = v42
                            default:
                                fmt.Println(v3)
                            }
                        }
                        _ = v34
                    }
                    v19 = float32(float64((((0.0 * (float32(0.0) * v19)) * (3.14159 / -2.25)) / (float32(float64(v19)) + -2.25))))
                    var v43 int64 = (v1 % (((v1 >> ((4294967295) & 63))) | 1))
                    switch v5 {
                    case "\n\t":
                        if ((v3 >= v3) && (v19 >= (float32(v2) * (((v19 - (v19 + v19)) - float32(1e10)) + 0.1)))) {
                            var v44 int = (fi22 | 42)
                            v45 := GenS7{F0: 0, F1: 0}
                            _ = v44
                            _ = v45
                        } else {
                            v46 := []int{-9223372036854775808, fi21, fi21, (fi21 / ((fi22) | 1))}
                            _ = v46
                        }
                    }
                    v43 = (v43 / (((v43 % (((^(v43))) | 1))) | 1))
                    _ = fi22
                    _ = v23
                    _ = v43
                }
                _ = fi21
            }
            v47 := GenS1{F0: 524, F1: 4294967295}
            _ = v47
        }
        _ = v19
    }
    var v48 int = 127
    if (32767 >= 0) {
        var v49 float64 = -2.25
        _ = v49
    }
    for fi50 := 0; fi50 < 8; fi50++ {
        v51 := GenS8{F0: false, F1: 0}
        var v52 float64 = v2
        _ = fi50
        _ = v51
        _ = v52
    }
    var v53 int64 = (v1 * v1)
    if (float64(-2.25) > v2) {
        if false {
            if (1e10 >= (0.0 + (-2.25 + float32(0.0)))) {
                if true {
                    if true {
                        v54 := []int{846}
                        _ = v54
                    } else {
                        v55 := []int{(v48 >> ((1) & 63)), (^(v48)), (v48 / (((v48 / (((v48 >> ((65535) & 63))) | 1))) | 1))}
                        v2 = v2
                        _ = v55
                    }
                }
                v3 = v3
                var v56 float64 = 1e10
                _ = v56
            } else {
                if true {
                    switch (v5 + v5) {
                    case "hello":
                        if true {
                            fmt.Println(v1)
                            for fi57 := 0; fi57 < 12; fi57++ {
                                if false {
                                    v3 = (v3 * 90)
                                    fmt.Println(v3, v48)
                                } else {
                                    var v58 uint64 = 104
                                    v59 := GenS7{F0: 1, F1: -1}
                                    fmt.Println(v3, v3)
                                    _ = v58
                                    _ = v59
                                }
                                v1 = v53
                                _ = fi57
                            }
                            fmt.Println(v53, v2)
                        } else {
                            if (2147483647 <= -1) {
                                var v60 uint8 = 2
                                v61 := []int{}
                                _ = v60
                                _ = v61
                            } else {
                                fmt.Println(v1, v5)
                                if (("foo bar" + "hello") >= (v5 + "\n\t")) {
                                    v62 := GenS9{F0: v5, F1: false}
                                    var v63 uint16 = 32767
                                    v5 = "\n\t"
                                    _ = v62
                                    _ = v63
                                }
                                v1 = -9223372036854775808
                            }
                            if (true && true) {
                                for fi64 := 0; fi64 < 5; fi64++ {
                                    v2 -= v2
                                    _ = fi64
                                }
                            }
                            if ((v5 + (v5 + v5)) != (("hello" + v5) + v5)) {
                                if false {
                                    v2 += v2
                                    v53 = v1
                                } else {
                                    var v65 uint16 = 32767
                                    _ = v65
                                }
                            } else {
                                if true {
                                    _ = len(v6)
                                    v53 = (-(v53))
                                } else {
                                    _ = len(v4)
                                }
                                var v66 int32 = 42
                                _ = v66
                            }
                        }
                    case "foo bar":
                        if true {
                            switch v48 {
                            case 4294967295:
                                v67 := GenS10{F0: -128, F1: (^(v48)), F2: 127}
                                _ = v67
                            case -9223372036854775808:
                                for fi68 := 0; fi68 < 16; fi68++ {
                                    fmt.Println(v3)
                                    v1 = (v53 * v53)
                                    _ = fi68
                                }
                            case 2147483647:
                                switch ("foo bar" + v5) {
                                case "hello":
                                    var v69 int = v48
                                    _ = v69
                                default:
                                    fmt.Println(v53, v1)
                                }
                            }
                        }
                    }
                } else {
                    fmt.Println(v48, v53)
                }
            }
            v70 := []int{43, (v48 - v48), v48, v48}
            _ = v70
        }
    } else {
        v48 += (-(v48))
        if ((59 < 222) || (float32(((v2 * 1e10) + v2)) >= 1.5)) {
            if false {
                v71 := GenS9{F0: ("a" + v5), F1: ((false && false) || true)}
                for fi72 := 0; fi72 < 8; fi72++ {
                    v3 += v3
                    v48 = (fi72 * fi72)
                    _ = fi72
                }
                v2 -= (v2 + float64((float32(v2) * (0.0 * 3.14159))))
                _ = v71
            } else {
                v73 := v4[42]
                if false {
                    if false {
                        v48 += (v48 | (v48 - (v48 - (v48 + v48))))
                        for fi74 := 0; fi74 < 2; fi74++ {
                            v75 := GenS11{F0: 255}
                            v76 := []int{746, fi74, 128}
                            _ = fi74
                            _ = v75
                            _ = v76
                        }
                        var v77 string = v5
                        _ = v77
                    }
                }
                if false {
                    var v78 string = "hello"
                    v79 := GenS12{F0: (v3 << ((18446744073709551615) & 63))}
                    _ = v78
                    _ = v79
                } else {
                    v2 = 0.1
                }
                _ = v73
            }
            if true {
                var v80 int8 = v3
                switch v5 {
                case "a":
                    fmt.Println(v2)
                case "hello":
                    v81 := GenS13{F0: 9223372036854775807}
                    _ = v81
                case "foo bar":
                    for fi82 := 0; fi82 < 15; fi82++ {
                        fmt.Println(v1)
                        if (true && (false && (v3 == (v80 ^ (v80 & v3))))) {
                            switch (v48 & 32767) {
                            case 32768:
                                var v83 int16 = 1
                                _ = v83
                            case 32767:
                                for fi84 := 0; fi84 < 3; fi84++ {
                                    if ((fi84 + (fi82 - v48)) > (fi84 | (v48 * v48))) {
                                        fmt.Println(fi82, v48)
                                    }
                                    v3 = v3
                                    v80 -= (v3 * v80)
                                    _ = fi84
                                }
                            }
                            v2 = v2
                        } else {
                            v85 := []int{(fi82 % ((fi82) | 1)), 428}
                            _ = v85
                        }
                        _ = fi82
                    }
                }
                fmt.Println(v80, v2)
                _ = v80
            }
        } else {
            var v86 int8 = 100
            _ = v86
        }
        v4[65535] = 65535
    }
    v87 := GenS7{F0: -269, F1: 32767}
    if true {
        for fi88 := 0; fi88 < 20; fi88++ {
            if true {
                fmt.Println(v3, v2)
                v3 += v3
                switch v5 {
                case "":
                    v48 -= fi88
                default:
                    var v89 uint16 = 1
                    _ = v89
                }
            }
            v2 = float64(((float32(v2) * 0.0) * 1.5))
            _ = fi88
        }
        fmt.Println(v53)
        if (((v5 + (v5 + "foo bar")) + v5) != ("a" + (v5 + "a"))) {
            v90 := GenS5{F0: 103}
            v53 = int64(v3)
            var v91 uint16 = 520
            _ = v90
            _ = v91
        } else {
            v1 += 347
            var v92 float64 = 1.5
            _ = v92
        }
    } else {
        var v93 int64 = 0
        _ = v93
    }
    fmt.Println(v5, v2)
    var v94 int = v48
    v94 -= (-(v94))
    var v95 uint = 668
    if true {
        fmt.Println(v53, v5)
    } else {
        v48 += v94
        var v96 uint = (^(v95))
        _ = v96
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v94:", v94)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v53
    _ = v6
    _ = v87
    _ = v94
    _ = v95
}
