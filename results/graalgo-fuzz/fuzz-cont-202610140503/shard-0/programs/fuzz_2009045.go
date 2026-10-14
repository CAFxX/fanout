package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 uint; F2 int16 }
type GenS4 struct { F0 uint16; F1 int32; F2 int32 }
func genFn0() uint16 {
    v7 := []int{-181, -9223372036854775808, -371}
    for fi8 := 0; fi8 < 14; fi8++ {
        switch "a" {
        case "hello":
            fmt.Println(fi8)
        case "\n\t":
            switch "a" {
            case "foo bar":
                fmt.Println(fi8)
            case "hello":
                v9 := GenS1{F0: 266, F1: 2, F2: "a"}
                _ = v9
            }
        case "foo bar":
            for fi10 := 0; fi10 < 6; fi10++ {
                var v11 uint64 = 42
                v11 = (^(v11))
                v11 = 4294967295
                var v12 uint32 = 65535
                v12 -= (v12 >> ((4294967295) & 63))
                if len(v7) > 0 {
                    v7[0] = fi8
                }
                _ = fi10
                _ = v11
                _ = v12
            }
        default:
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        }
        _ = fi8
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        if (3.14159 != 3.14159) {
            var v14 string = "hello"
            fmt.Println(fi13, v14)
            _ = v14
        } else {
            fmt.Println(fi13)
        }
        v15 := GenS1{F0: 42, F1: 2147483648, F2: "hello"}
        _ = fi13
        _ = v15
    }
    var v16 float32 = 3.14159
    _ = v16
    _ = v7
    return 0
}
func genFn1() int {
    switch "foo bar" {
    case "\n\t":
        switch "foo bar" {
        case "\n\t":
            var v17 int16 = 9
            _ = v17
        }
    case "hello":
        v18 := []int{}
        _ = v18
    }
    if (239 != 1) {
        v19 := GenS1{F0: -9223372036854775808, F1: 42, F2: "foo bar"}
        var v20 uint = 417
        for fi21 := 0; fi21 < 15; fi21++ {
            if true {
                if (!(1e10 >= float32(1e10))) {
                    var v22 bool = false
                    _ = v22
                } else {
                    v23 := []int{}
                    _ = v23
                }
                var v24 uint32 = 0
                v20 = (v20 + (^(v20)))
                _ = v24
            }
            _ = fi21
        }
        _ = v19
        _ = v20
    } else {
        var v25 float32 = ((1.5 * 3.14159) * 0.0)
        _ = v25
    }
    var v26 uint32 = 2147483648
    _ = v26
    return 2147483647
}

func main() {
    var v1 uint64 = 1
    var v2 uint32 = 65535
    var v3 uint16 = 672
    var v4 uint16 = (v3 ^ 255)
    var v5 int16 = 42
    v6 := make(map[int]string)
    fmt.Println(v4)
    for fi27 := 0; fi27 < 15; fi27++ {
        v3 -= (v4 % ((303) | 1))
        _ = fi27
    }
    switch (("" + ("hello" + ("" + "hello"))) + ("" + "a")) {
    case "\n\t":
        var v28 uint32 = (v2 << ((4294967295) & 63))
        _ = v28
    default:
        if true {
            var v29 uint64 = (v1 & 2)
            var v30 uint32 = 730
            _ = v29
            _ = v30
        }
    }
    for fi31 := 0; fi31 < 3; fi31++ {
        v1 = (v1 * (^(v1)))
        v3 = genFn0()
        var v32 bool = (false && (("foo bar" + "foo bar") >= "hello"))
        _ = fi31
        _ = v32
    }
    if (!(true || true)) {
        v33 := []int{252, 778, int(v4)}
        if (false || false) {
            v1 = 0
            for fi34 := 0; fi34 < 14; fi34++ {
                if false {
                    v4 -= (v4 ^ uint16(v1))
                    v35 := GenS1{F0: -9223372036854775808, F1: (^(v2)), F2: "foo bar"}
                    _ = v35
                }
                fmt.Println(v3)
                _ = fi34
            }
            var v36 bool = false
            _ = v36
        }
        var v37 uint64 = (^(v1))
        _ = v33
        _ = v37
    } else {
        switch 4294967295 {
        case 41:
            fmt.Println(v1, v3)
        case 53:
            v4 = v4
        case 27:
            v5 -= v5
        default:
            var v38 int16 = v5
            _ = v38
        }
    }
    if true {
        v39 := []int{27, -770}
        _ = v39
    } else {
        if false {
            v40 := GenS2{F0: "a"}
            v41 := []int{-1, -1}
            v2 = v2
            _ = v40
            _ = v41
        }
        for fi42 := 0; fi42 < 1; fi42++ {
            var v43 float32 = 3.14159
            _ = fi42
            _ = v43
        }
    }
    for fi44 := 0; fi44 < 19; fi44++ {
        var v45 float32 = 3.14159
        _ = fi44
        _ = v45
    }
    if (32767 != -385) {
        if (-2147483648 <= 42) {
            if false {
                var v46 int8 = -11
                fmt.Println(v2, v1)
                _ = v46
            } else {
                v47 := GenS2{F0: ("\n\t" + (("" + "hello") + ""))}
                var v48 float64 = (((-2.25 * -2.25) + 1.5) * (((0.1 + 3.14159) / 1.5) * 3.14159))
                _ = v47
                _ = v48
            }
            switch "" {
            case "a":
                if ((v1 > v1) && true) {
                    if (0 != 255) {
                        switch 2147483647 {
                        case -113:
                            v49 := v6[1]
                            _ = v49
                        case -9223372036854775808:
                            var v50 float32 = (0.0 - float32(-2.25))
                            _ = v50
                        }
                        var v51 uint = 2
                        v52 := []int{127, -15, 128}
                        _ = v51
                        _ = v52
                    }
                    for fi53 := 0; fi53 < 18; fi53++ {
                        for fi54 := 0; fi54 < 19; fi54++ {
                            if false {
                                v55 := GenS2{F0: ""}
                                if (!(true && ((((69 == 1) || true) || ((3.14159 * 1e10) == 0.0)) || false))) {
                                    if ((^(v5)) != (v5 % ((1000) | 1))) {
                                        switch -2147483648 {
                                        case -826:
                                            var v56 float32 = (0.1 - float32(-2.25))
                                            _ = v56
                                        case 4294967295:
                                            var v57 int8 = 127
                                            _ = v57
                                        default:
                                            switch "\n\t" {
                                            case "\n\t":
                                                fmt.Println(v3, fi54)
                                            case "a":
                                                if (!true) {
                                                    switch "foo bar" {
                                                    case "":
                                                        delete(v6, fi54)
                                                    case "\n\t":
                                                        var v58 uint = 4294967295
                                                        _ = v58
                                                    case "foo bar":
                                                        var v59 string = ""
                                                        _ = v59
                                                    default:
                                                        v1 += (v1 >> ((4294967295) & 63))
                                                    }
                                                    v60 := GenS2{F0: (("foo bar" + "foo bar") + "hello")}
                                                    _ = v60
                                                } else {
                                                    fmt.Println(v2, fi54, v1)
                                                    if false {
                                                        switch "hello" {
                                                        case "":
                                                            v61 := []int{}
                                                            _ = v61
                                                        case "\n\t":
                                                            var v62 string = "\n\t"
                                                            _ = v62
                                                        case "a":
                                                            if false {
                                                                v3 = v3
                                                                var v63 int64 = 9223372036854775807
                                                                _ = v63
                                                            } else {
                                                                v3 += v3
                                                                v64 := GenS3{F0: 23, F1: 255, F2: (v5 - (v5 | 32767))}
                                                                _ = v64
                                                            }
                                                        }
                                                        var v65 float64 = float64(1.5)
                                                        _ = v65
                                                    }
                                                }
                                            case "foo bar":
                                                var v66 string = "hello"
                                                _ = v66
                                            default:
                                                v3 += (v3 ^ genFn0())
                                            }
                                        }
                                        fmt.Println(v5)
                                    } else {
                                        fmt.Println(v4)
                                        fmt.Println(v5, fi53)
                                    }
                                }
                                var v67 string = ("hello" + "\n\t")
                                _ = v55
                                _ = v67
                            }
                            v68 := GenS3{F0: 127, F1: uint(v4), F2: v5}
                            var v69 uint = 18446744073709551615
                            _ = fi54
                            _ = v68
                            _ = v69
                        }
                        if (127 < 197) {
                            v70 := GenS4{F0: (v4 - uint16(v1)), F1: 2147483647, F2: 913}
                            delete(v6, 32767)
                            _ = len(v6)
                            _ = v70
                        }
                        _ = fi53
                    }
                    var v71 int8 = -128
                    _ = v71
                } else {
                    v5 -= v5
                }
            }
            v5 += 42
        } else {
            var v72 int = 32767
            var v73 float64 = 1e10
            v6[v72] = "\n\t"
            _ = v72
            _ = v73
        }
        v74 := []int{genFn1(), -2147483648, -9223372036854775808}
        _ = v74
    } else {
        fmt.Println(v5, v4)
        v75 := GenS3{F0: 255, F1: 65535, F2: (-(v5))}
        var v76 bool = (-1 >= -1)
        _ = v75
        _ = v76
    }
    v1 = (v1 * (^(v1)))
    v1 = v1
    var v77 int16 = (v5 * -962)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v77
}
