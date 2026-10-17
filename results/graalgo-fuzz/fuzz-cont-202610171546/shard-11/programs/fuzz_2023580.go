package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint }
type GenS2 struct { F0 int8; F1 string; F2 int8 }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 int32; F1 int32 }
type GenS5 struct { F0 string; F1 int32 }
type GenS6 struct { F0 uint64; F1 uint8; F2 string }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 string; F1 uint; F2 bool }
type GenS11 struct { F0 uint64; F1 int8; F2 bool }
type GenS12 struct { F0 uint8; F1 bool }
type GenS13 struct { F0 uint32; F1 int32 }

func main() {
    v1 := []string{("foo bar" + ""), "foo bar"}
    var v2 bool = (1 <= -128)
    var v3 uint64 = 0
    var v4 uint64 = v3
    v5 := make(map[string]bool)
    v2 = v2
    if false {
        v2 = true
        var v6 uint = 1
        _ = v6
    }
    var v7 uint32 = 171
    var v8 int16 = 17
    v4 += (v4 << ((132) & 63))
    if v2 {
        var v9 uint32 = (^(v7))
        v10 := GenS1{F0: (v8 & v8), F1: 4294967295, F2: 4294967295}
        if (!v2) {
            for fi11 := 0; fi11 < 8; fi11++ {
                var v12 uint = 4294967295
                _ = fi11
                _ = v12
            }
        } else {
            if ((!v2) && true) {
                for fi13 := 0; fi13 < 6; fi13++ {
                    var v14 int8 = 1
                    v4 = (v3 / ((v3) | 1))
                    _ = fi13
                    _ = v14
                }
                fmt.Println(v2, v3)
                for fi15 := 0; fi15 < 5; fi15++ {
                    var v16 int = fi15
                    v4 = 0
                    var v17 int8 = 0
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            }
            if v2 {
                if (v9 != uint32(v4)) {
                    if v2 {
                        if ((v9 & (v7 % (((v7 >> ((868) & 63))) | 1))) < v9) {
                            var v18 uint = 255
                            _ = v18
                        }
                        v19 := GenS2{F0: -128, F1: (("\n\t" + "a") + "\n\t"), F2: 100}
                        if v2 {
                            _ = len(v1)
                            var v20 int32 = 767
                            _ = v20
                        }
                        _ = v19
                    }
                } else {
                    v5[(("a" + ("a" + ("hello" + ("foo bar" + "\n\t")))) + "\n\t")] = true
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v5["\n\t"] = (v4 < (v4 + (v4 % ((18446744073709551615) | 1))))
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            } else {
                v21 := GenS2{F0: -88, F1: ("foo bar" + "a"), F2: -48}
                switch 594 {
                case 32768:
                    v22 := GenS2{F0: 0, F1: "a", F2: 51}
                    _ = v22
                case 0:
                    if (42 != -413) {
                        var v23 float32 = 0.1
                        fmt.Println(v3)
                        _ = v23
                    } else {
                        if false {
                            v24 := v5["\n\t"]
                            _ = v24
                        }
                        for fi25 := 0; fi25 < 19; fi25++ {
                            var v26 uint64 = (v4 - v3)
                            if true {
                                var v27 uint = 255
                                _ = v27
                            } else {
                                for fi28 := 0; fi28 < 14; fi28++ {
                                    fmt.Println(v8, v3, fi28)
                                    var v29 float32 = (1e10 + -2.25)
                                    _ = fi28
                                    _ = v29
                                }
                            }
                            _ = fi25
                            _ = v26
                        }
                        for fi30 := 0; fi30 < 5; fi30++ {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            fmt.Println(v8, v9)
                            v31 := v5[(("hello" + "a") + "hello")]
                            _ = fi30
                            _ = v31
                        }
                    }
                default:
                    fmt.Println(v2)
                }
                _ = v21
            }
        }
        _ = v10
        _ = v9
    }
    switch 2147483647 {
    case 42:
        v32 := GenS1{F0: 32767, F1: 65535, F2: 18446744073709551615}
        _ = v32
    case 32767:
        if v2 {
            v4 -= (v4 >> ((18) & 63))
            if false {
                if v2 {
                    v2 = ("" == "foo bar")
                    v33 := GenS3{F0: 18446744073709551615, F1: 0}
                    for fi34 := 0; fi34 < 1; fi34++ {
                        for fi35 := 0; fi35 < 1; fi35++ {
                            v36 := GenS4{F0: 1, F1: 1}
                            var v37 float32 = -2.25
                            v38 := GenS5{F0: "", F1: int32(v4)}
                            _ = fi35
                            _ = v36
                            _ = v37
                            _ = v38
                        }
                        v39 := GenS6{F0: 42, F1: 127, F2: ("hello" + "")}
                        if (!(1 > 1)) {
                            v2 = ((((!v2) && true) || (true || (-2147483648 <= 0))) || v2)
                            v7 = (v7 << ((65535) & 63))
                        } else {
                            for fi40 := 0; fi40 < 14; fi40++ {
                                if (((float32(3.14159) + 0.0) - 0.1) == (3.14159 / ((0.0 / float32(1.5)) + -2.25))) {
                                    if (369 != 518) {
                                        if false {
                                            if v2 {
                                                var v41 uint16 = 0
                                                _ = v41
                                            }
                                            v42 := GenS3{F0: 2, F1: 9223372036854775807}
                                            if v2 {
                                                v2 = true
                                                var v43 int8 = 1
                                                v44 := GenS7{F0: 65535}
                                                _ = v43
                                                _ = v44
                                            }
                                            var v45 bool = v2
                                            _ = v42
                                            _ = v45
                                        }
                                    }
                                }
                                if false {
                                    var v46 uint64 = (v3 * uint64(v7))
                                    _ = v46
                                } else {
                                    v47 := GenS2{F0: 1, F1: "", F2: -36}
                                    fmt.Println(fi40, fi40, v2)
                                    _ = v47
                                }
                                if (v8 < v8) {
                                    v8 = v8
                                    var v48 float32 = (-2.25 + ((0.1 - 0.1) + -2.25))
                                    _ = v48
                                } else {
                                    fmt.Println(v7, v3, fi34)
                                }
                                _ = fi40
                            }
                            if v2 {
                                var v49 uint8 = uint8(v4)
                                _ = v49
                            }
                            v2 = (v2 && false)
                        }
                        _ = fi34
                        _ = v39
                    }
                    _ = v33
                } else {
                    v4 -= (v3 ^ v4)
                    var v50 int64 = -1
                    var v51 int32 = -876
                    _ = v50
                    _ = v51
                }
                for fi52 := 0; fi52 < 9; fi52++ {
                    for fi53 := 0; fi53 < 8; fi53++ {
                        switch "a" {
                        case "hello":
                            v54 := GenS8{F0: 100}
                            _ = v54
                        default:
                            if false {
                                v55 := GenS2{F0: 0, F1: "foo bar", F2: 100}
                                v56 := GenS9{F0: (-1 == -34)}
                                if v2 {
                                    if (32767 == 32767) {
                                        var v57 int = (-(fi52))
                                        fmt.Println(v2)
                                        _ = v57
                                    }
                                    fmt.Println(v2, v3, fi53)
                                }
                                v2 = false
                                if true {
                                    var v58 float32 = 3.14159
                                    fmt.Println(v58, v2, fi53)
                                    if (true && ((^(v3)) < (v3 << ((42) & 63)))) {
                                        var v59 float64 = (float64(1.5) - 0.1)
                                        _ = len(v1)
                                        var v60 int = fi53
                                        _ = v59
                                        _ = v60
                                    }
                                    v58 -= (((v58 * v58) + (v58 - (float32(1e10) / (1e10 * -2.25)))) / v58)
                                    _ = v58
                                }
                                switch ("\n\t" + "hello") {
                                case "\n\t":
                                    v2 = (v7 < (v7 >> ((1) & 63)))
                                }
                                _ = v55
                                _ = v56
                            } else {
                                v8 = v8
                            }
                        }
                        v4 = (v4 % ((4294967295) | 1))
                        if v2 {
                            if v2 {
                                fmt.Println(v7, v3)
                                var v61 bool = v2
                                _ = v61
                            }
                            v62 := GenS6{F0: v4, F1: 128, F2: ("foo bar" + "")}
                            switch "a" {
                            case "":
                                v63 := GenS10{F0: "foo bar", F1: uint(v4), F2: (!(v2 && v2))}
                                _ = v63
                            }
                            _ = v62
                        } else {
                            v3 += v3
                            v2 = v2
                            fmt.Println(v7)
                        }
                        _ = fi53
                    }
                    var v64 float32 = (((0.0 - (3.14159 - float32(1e10))) - 1.5) / -2.25)
                    _ = fi52
                    _ = v64
                }
            } else {
                var v65 uint = 42
                v66 := v5["a"]
                _ = v65
                _ = v66
            }
            switch -658 {
            case 9223372036854775807:
                if v2 {
                    var v67 uint64 = 2
                    v68 := GenS11{F0: 405, F1: 100, F2: v2}
                    _ = v67
                    _ = v68
                }
            case -222:
                v7 += v7
            case 4294967295:
                var v69 int16 = (v8 & v8)
                _ = v69
            default:
                var v70 int8 = -78
                _ = v70
            }
        }
    }
    switch "foo bar" {
    case "a":
        switch -826 {
        case -2147483648:
            v2 = v2
        }
    case "":
        var v71 int = 730
        _ = v71
    case "hello":
        for fi72 := 0; fi72 < 1; fi72++ {
            v3 -= v4
            _ = fi72
        }
    default:
        fmt.Println(v7, v7, v7)
    }
    v4 = (v3 + v4)
    var v73 uint32 = v7
    fmt.Println(v73, v3, v2)
    v3 += 42
    if ((42 == 797) && (255 >= 1)) {
        switch 127 {
        case -384:
            var v74 int16 = 143
            _ = v74
        default:
            v75 := GenS7{F0: 194}
            _ = v75
        }
        if ((v4 | v4) >= (v3 * (v4 | v4))) {
            if (127 < 255) {
                if (65535 > 586) {
                    v76 := GenS12{F0: 52, F1: true}
                    _ = v76
                } else {
                    v3 -= v3
                }
                var v77 uint32 = v7
                v8 -= v8
                _ = v77
            } else {
                v4 = uint64(v73)
                fmt.Println(v8, v73, v8)
                for fi78 := 0; fi78 < 10; fi78++ {
                    fmt.Println(fi78, v4, v8)
                    var v79 float64 = 3.14159
                    var v80 uint64 = (v3 ^ uint64(v7))
                    _ = fi78
                    _ = v79
                    _ = v80
                }
            }
            if len(v1) > 0 {
                v1[0] = "a"
            }
            v81 := GenS13{F0: (v7 * 425), F1: 422}
            _ = v81
        }
    }
    for fi82 := 0; fi82 < 18; fi82++ {
        switch ("a" + "a") {
        case "hello":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        _ = len(v5)
        _ = fi82
    }
    v73 -= 806
    for fi83 := 0; fi83 < 7; fi83++ {
        v8 = v8
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v8, v7, v8)
        case "":
            for fi84 := 0; fi84 < 2; fi84++ {
                var v85 string = "a"
                delete(v5, v85)
                if len(v1) > 0 {
                    v1[0] = v85
                }
                _ = fi84
                _ = v85
            }
        }
        var v86 uint32 = v7
        _ = fi83
        _ = v86
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v73:", v73)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v73
    _ = v8
}
