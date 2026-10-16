package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int8 }
type GenS2 struct { F0 int16; F1 uint16; F2 int8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8; F1 uint64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int; F1 int16; F2 int64 }
type GenS8 struct { F0 uint16; F1 uint; F2 bool }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int8; F1 uint8 }
type GenS11 struct { F0 string; F1 int16; F2 int }
type GenS12 struct { F0 int16 }
func genFn0() int32 {
    if false {
        v5 := GenS1{F0: 0, F1: 42, F2: -128}
        _ = v5
    } else {
        var v6 uint32 = 0
        var v7 int8 = -128
        _ = v6
        _ = v7
    }
    if (false || true) {
        if (1 >= 1) {
            v8 := GenS2{F0: 741, F1: 1, F2: -128}
            var v9 uint8 = 6
            var v10 int8 = -1
            _ = v10
            _ = v8
            _ = v9
        }
        if (-50 > 0) {
            var v11 uint = 794
            if (!true) {
                v11 -= 42
            } else {
                var v12 bool = true
                _ = v12
            }
            _ = v11
        }
        v13 := []int{56, 295}
        _ = v13
    } else {
        var v14 string = "hello"
        _ = v14
    }
    var v15 uint16 = 302
    _ = v15
    return 2147483647
}
func genFn1(p0 uint16, p1 int32) bool {
    for fi16 := 0; fi16 < 12; fi16++ {
        if true {
            if ((!(true || true)) && false) {
                v17 := GenS3{F0: p1}
                v18 := GenS4{F0: 59, F1: 4294967295}
                _ = v17
                _ = v18
            } else {
                v19 := []int{}
                var v20 uint32 = 96
                _ = v19
                _ = v20
            }
            for fi21 := 0; fi21 < 4; fi21++ {
                var v22 int8 = 42
                var v23 int16 = 42
                var v24 bool = ((false || false) && (2 > 255))
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
            }
            if true {
                var v25 int = 1
                for fi26 := 0; fi26 < 10; fi26++ {
                    v25 += 127
                    v27 := []int{618, fi26, (^(fi16))}
                    _ = fi26
                    _ = v27
                }
                _ = v25
            }
        }
        _ = fi16
    }
    var v28 float64 = 3.14159
    v28 = (-2.25 / v28)
    p0 -= p0
    v29 := []int{-560, -568, -850}
    _ = p0
    _ = p1
    _ = v28
    _ = v29
    return false
}

func main() {
    var v1 uint16 = 111
    v2 := make(map[int]uint32)
    var v3 string = (((((("a" + "hello") + "\n\t") + "") + ("" + "")) + "a") + ("hello" + "foo bar"))
    var v4 uint64 = 2
    fmt.Println(v1, v4)
    v1 = 1
    v3 = ("a" + "hello")
    var v30 uint32 = 2
    for fi31 := 0; fi31 < 7; fi31++ {
        v32 := []int{fi31, (fi31 >> ((711) & 63)), (fi31 * -14)}
        if (!(false || (true || false))) {
            if true {
                switch v3 {
                case "":
                    fmt.Println(fi31, v4)
                case "a":
                    var v33 float32 = 3.14159
                    _ = v33
                }
                for fi34 := 0; fi34 < 6; fi34++ {
                    if true {
                        if false {
                            v35 := GenS5{F0: 127}
                            if false {
                                switch "a" {
                                case "\n\t":
                                    if (!false) {
                                        fmt.Println(v4, fi31, v1)
                                        var v36 int64 = 1
                                        fmt.Println(fi34, fi34, fi31)
                                        _ = v36
                                    }
                                case "":
                                    v30 = 93
                                case "a":
                                    var v37 int64 = -1
                                    _ = v37
                                default:
                                    var v38 int64 = -115
                                    _ = v38
                                }
                            }
                            _ = v35
                        }
                        delete(v2, 470)
                    }
                    switch v3 {
                    case "\n\t":
                        v39 := GenS6{F0: v30}
                        _ = v39
                    default:
                        fmt.Println(fi34)
                    }
                    var v40 uint8 = 188
                    _ = fi34
                    _ = v40
                }
            } else {
                for fi41 := 0; fi41 < 3; fi41++ {
                    var v42 uint64 = 141
                    _ = fi41
                    _ = v42
                }
                v43 := GenS1{F0: 4294967295, F1: 42, F2: -1}
                _ = v43
            }
        } else {
            var v44 uint = 65535
            _ = v44
        }
        _ = fi31
        _ = v32
    }
    v45 := GenS7{F0: -1, F1: -603, F2: 1}
    var v46 uint32 = (v30 + 736)
    var v47 uint = 453
    for fi48 := 0; fi48 < 15; fi48++ {
        v49 := []int{fi48, fi48}
        for fi50 := 0; fi50 < 4; fi50++ {
            v51 := GenS4{F0: 1, F1: 993}
            var v52 float32 = (float32(0.0) - -2.25)
            _ = fi50
            _ = v51
            _ = v52
        }
        _ = fi48
        _ = v49
    }
    var v53 bool = (1e10 > (0.1 - (-2.25 + 0.0)))
    switch ("\n\t" + "hello") {
    case "":
        var v54 uint8 = 46
        _ = v54
    default:
        if false {
            for fi55 := 0; fi55 < 4; fi55++ {
                if (65535 >= 0) {
                    switch ((v3 + "hello") + v3) {
                    case "a":
                        if (!(3.14159 < (0.1 - 0.1))) {
                            for fi56 := 0; fi56 < 9; fi56++ {
                                if (uint8(v47) >= 2) {
                                    fmt.Println(v4, v46)
                                }
                                _ = fi56
                            }
                        }
                    case "\n\t":
                        fmt.Println(v46, v30, v3)
                    case "hello":
                        if (fi55 != -9223372036854775808) {
                            v57 := GenS8{F0: v1, F1: (v47 % ((v47) | 1)), F2: (0 < 9223372036854775807)}
                            if (!v53) {
                                var v58 uint64 = v4
                                for fi59 := 0; fi59 < 11; fi59++ {
                                    fmt.Println(fi59, fi55)
                                    v60 := GenS9{F0: 1}
                                    _ = fi59
                                    _ = v60
                                }
                                switch "foo bar" {
                                case "foo bar":
                                    for fi61 := 0; fi61 < 5; fi61++ {
                                        v62 := GenS2{F0: 1, F1: (v1 - (v1 << ((v47) & 63))), F2: -128}
                                        var v63 bool = true
                                        _ = fi61
                                        _ = v62
                                        _ = v63
                                    }
                                default:
                                    v64 := []int{(fi55 ^ (fi55 - 336))}
                                    _ = v64
                                }
                                _ = v58
                            }
                            _ = v57
                        }
                    }
                } else {
                    var v65 bool = (0.1 != float32(-2.25))
                    _ = v65
                }
                if (!(0.1 > 0.1)) {
                    v4 += v4
                    var v66 int8 = 42
                    for fi67 := 0; fi67 < 18; fi67++ {
                        v68 := GenS1{F0: (v47 ^ (v47 << ((739) & 63))), F1: 31, F2: v66}
                        if v53 {
                            v69 := GenS10{F0: 100, F1: 214}
                            if true {
                                switch (fi55 & fi55) {
                                case 334:
                                    var v70 string = v3
                                    _ = v70
                                case 127:
                                    v46 = (v30 << ((v47) & 63))
                                case 0:
                                    var v71 int16 = 32767
                                    _ = v71
                                default:
                                    switch (v3 + v3) {
                                    case "a":
                                        v72 := GenS6{F0: uint32(v47)}
                                        _ = v72
                                    case "":
                                        delete(v2, (fi67 * (fi67 | 42)))
                                    case "hello":
                                        var v73 uint64 = (v4 & v4)
                                        _ = v73
                                    }
                                }
                            }
                            if v53 {
                                v4 += v4
                                fmt.Println(v30, v53)
                                v2[fi67] = uint32(v47)
                            }
                            _ = v69
                        }
                        _ = fi67
                        _ = v68
                    }
                    _ = v66
                } else {
                    v74 := GenS11{F0: (v3 + v3), F1: 0, F2: (fi55 >> (((v47 & v47)) & 63))}
                    v75 := GenS5{F0: 1}
                    _ = v74
                    _ = v75
                }
                _ = fi55
            }
            v76 := []int{127}
            _ = v76
        }
    }
    switch 4294967295 {
    case 1:
        fmt.Println(v53, v47)
    }
    var v77 uint8 = 127
    if (!(!(0 != (^(v46))))) {
        if v53 {
            if (v53 && v53) {
                for fi78 := 0; fi78 < 13; fi78++ {
                    v79 := GenS10{F0: int8(v30), F1: v77}
                    _ = fi78
                    _ = v79
                }
            }
        } else {
            v30 -= (v30 + v30)
            if v53 {
                v80 := GenS8{F0: 2, F1: (v47 - 293), F2: ((v1 > (v1 % (((v1 ^ v1)) | 1))) && true)}
                v81 := GenS11{F0: "hello", F1: 0, F2: 9223372036854775807}
                _ = v80
                _ = v81
            } else {
                fmt.Println(v47, v1, v3)
                for fi82 := 0; fi82 < 7; fi82++ {
                    v83 := GenS12{F0: -1}
                    v84 := []int{fi82, (fi82 | fi82), fi82, (fi82 | fi82)}
                    _ = fi82
                    _ = v83
                    _ = v84
                }
            }
        }
    }
    v77 -= v77
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v53:", v53)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v53
    _ = v77
}
