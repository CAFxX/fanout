package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint; F1 uint32; F2 int32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint16; F1 int }
type GenS8 struct { F0 uint16 }
func genFn0(p0 uint32, p1 int8) int64 {
    var v7 float64 = (-2.25 * 0.0)
    var v8 bool = false
    switch -1 {
    case -181:
        if (0 > -179) {
            v9 := []int{-1, 32768}
            _ = v9
        }
    case 128:
        if (true || ((1 >= 1) || (false || ((!v8) && true)))) {
            var v10 uint32 = 0
            v11 := GenS1{F0: 2}
            fmt.Println(v10)
            _ = v10
            _ = v11
        } else {
            v12 := GenS2{F0: -1}
            v13 := GenS1{F0: 36}
            v7 = (v7 / float64(((3.14159 + ((0.0 + -2.25) + (((((float32(v7) / 1.5) * (1.5 + 0.1)) / 1.5) * -2.25) - 3.14159))) - 0.0)))
            _ = v12
            _ = v13
        }
    default:
        var v14 int32 = -2147483648
        _ = v14
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        for fi16 := 0; fi16 < 1; fi16++ {
            switch "\n\t" {
            case "hello":
                v17 := GenS3{F0: 1000}
                _ = v17
            case "\n\t":
                p1 += -128
            }
            _ = fi16
        }
        for fi18 := 0; fi18 < 15; fi18++ {
            v19 := GenS1{F0: 143}
            p0 = p0
            fmt.Println(p0, p0)
            _ = fi18
            _ = v19
        }
        for fi20 := 0; fi20 < 4; fi20++ {
            var v21 int16 = 1
            if v8 {
                if (18446744073709551615 >= 42) {
                    v22 := GenS2{F0: 1}
                    v23 := GenS2{F0: 1}
                    var v24 int8 = p1
                    _ = v22
                    _ = v23
                    _ = v24
                }
                if (true || v8) {
                    v7 += v7
                    v25 := []int{}
                    var v26 uint8 = uint8(p0)
                    _ = v25
                    _ = v26
                } else {
                    var v27 uint = 42
                    v27 = v27
                    _ = v27
                }
                v28 := []int{}
                _ = v28
            }
            _ = fi20
            _ = v21
        }
        _ = fi15
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := make(map[string]bool)
    v2 := []int{416, 42, 32767, 472}
    v3 := make(map[int]uint32)
    var v4 int8 = 127
    v5 := make(map[int]string)
    var v6 uint64 = 296
    var v29 bool = false
    var v30 int16 = -845
    var v31 int16 = v30
    v30 -= 32767
    fmt.Println(v6)
    var v32 uint16 = 32767
    v2 = append(v2, -572)
    fmt.Println(v30, v4, v29)
    if (!(v4 > v4)) {
        for fi33 := 0; fi33 < 9; fi33++ {
            v30 = v30
            var v34 uint = 42
            if v29 {
                var v35 int32 = -2147483648
                switch int(v32) {
                case 9223372036854775807:
                    v4 = int8(v30)
                case -9223372036854775808:
                    for fi36 := 0; fi36 < 9; fi36++ {
                        fmt.Println(fi33, v34)
                        v34 = (v34 % ((v34) | 1))
                        _ = fi36
                    }
                default:
                    if v29 {
                        if (v31 < (v31 ^ (v30 - -1))) {
                            v31 += v31
                            switch "\n\t" {
                            case "foo bar":
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            case "a":
                                var v37 bool = v29
                                _ = v37
                            case "":
                                var v38 uint16 = v32
                                _ = v38
                            default:
                                v29 = v29
                            }
                            for fi39 := 0; fi39 < 20; fi39++ {
                                _ = len(v2)
                                var v40 uint = 65535
                                v6 -= v6
                                _ = fi39
                                _ = v40
                            }
                        }
                        v41 := GenS4{F0: (^(v34)), F1: 2, F2: (v35 >> ((v34) & 63))}
                        v42 := GenS5{F0: 32767}
                        _ = v41
                        _ = v42
                    } else {
                        v43 := GenS3{F0: (-(v30))}
                        _ = v43
                    }
                }
                v32 -= (v32 * 1)
                _ = v35
            } else {
                if false {
                    var v44 bool = (v32 >= v32)
                    _ = v44
                } else {
                    v6 = (v6 % ((v6) | 1))
                }
                fmt.Println(v6, v31, v30)
                for fi45 := 0; fi45 < 3; fi45++ {
                    var v46 uint8 = 42
                    var v47 uint8 = 1
                    if true {
                        if (uint8(v32) == (v46 / ((v47) | 1))) {
                            fmt.Println(v47)
                        }
                        v30 = (v31 ^ v31)
                        v48 := GenS4{F0: v34, F1: 2147483648, F2: -2147483648}
                        _ = v48
                    } else {
                        var v49 uint64 = (v6 / ((v6) | 1))
                        _ = v49
                    }
                    _ = fi45
                    _ = v46
                    _ = v47
                }
            }
            _ = fi33
            _ = v34
        }
    } else {
        for fi50 := 0; fi50 < 10; fi50++ {
            var v51 uint16 = (v32 ^ v32)
            switch (fi50 & fi50) {
            case 42:
                switch "hello" {
                case "":
                    var v52 int8 = v4
                    _ = v52
                case "\n\t":
                    v4 = v4
                case "hello":
                    v51 += (v32 + 65535)
                default:
                    fmt.Println(v51, v31, v51)
                }
            case 9223372036854775807:
                var v53 int64 = 0
                _ = v53
            case 259:
                if (v29 || false) {
                    v54 := v1[("\n\t" + "\n\t")]
                    for fi55 := 0; fi55 < 14; fi55++ {
                        if (0.1 > (1.5 / 1.0)) {
                            switch fi55 {
                            case 42:
                                var v56 int16 = (v30 ^ v30)
                                _ = v56
                            default:
                                v4 -= v4
                            }
                            switch fi55 {
                            case 32768:
                                var v57 float64 = (1e10 - (-2.25 + 0.1))
                                _ = v57
                            case 1:
                                switch ("" + "") {
                                case "hello":
                                    v6 = uint64(v51)
                                case "\n\t":
                                    switch "\n\t" {
                                    case "hello":
                                        var v58 float32 = 3.14159
                                        _ = v58
                                    case "a":
                                        var v59 int = fi55
                                        _ = v59
                                    default:
                                        v5[(fi50 ^ -1)] = "\n\t"
                                    }
                                case "foo bar":
                                    v60 := v1[("hello" + ("" + (((("a" + "") + "a") + "a") + ("hello" + "a"))))]
                                    _ = v60
                                default:
                                    v61 := GenS6{F0: ((224 >= v6) && true)}
                                    _ = v61
                                }
                            case 9223372036854775807:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v62 int32 = int32(fi50)
                                if false {
                                    if ((v4 ^ v4) != (v4 / ((95) | 1))) {
                                        v5[9223372036854775807] = "hello"
                                    } else {
                                        if (!((!v29) || ((2 == (v51 + v51)) || v29))) {
                                            v63 := GenS7{F0: uint16(v6), F1: 1}
                                            var v64 float64 = float64((1.5 + 3.14159))
                                            var v65 uint8 = 255
                                            _ = v63
                                            _ = v64
                                            _ = v65
                                        }
                                        var v66 int32 = (-(v62))
                                        switch "hello" {
                                        case "hello":
                                            var v67 float64 = float64(0.0)
                                            _ = v67
                                        }
                                        _ = v66
                                    }
                                    if len(v2) > 0 {
                                        v2[0] = -46
                                    }
                                } else {
                                    var v68 string = "hello"
                                    fmt.Println(v29, fi50, v32)
                                    if (0.0 > 1.5) {
                                        var v69 int8 = 42
                                        fmt.Println(v32, v69, fi50)
                                        fmt.Println(v30)
                                        var v70 int = (fi50 << ((2) & 63))
                                        v31 = (v31 << ((4294967295) & 63))
                                        v71 := GenS8{F0: (^(v32))}
                                        _ = v69
                                        _ = v70
                                        _ = v71
                                    } else {
                                        if v54 {
                                            v72 := GenS4{F0: 42, F1: 203, F2: (-(v62))}
                                            v1[((v68 + ((v68 + ("\n\t" + v68)) + "")) + ((v68 + v68) + v68))] = (112 < 157)
                                            var v73 int8 = 100
                                            _ = v72
                                            _ = v73
                                        }
                                    }
                                    _ = v68
                                }
                                _ = v62
                            }
                        }
                        _ = fi55
                    }
                    _ = v54
                } else {
                    for fi74 := 0; fi74 < 9; fi74++ {
                        v30 = v31
                        v31 -= int16(fi74)
                        fmt.Println(v32)
                        switch ("\n\t" + "") {
                        case "a":
                            v75 := GenS7{F0: (v32 / ((v32) | 1)), F1: (fi74 / ((32767) | 1))}
                            _ = v75
                        case "":
                            v51 += 42
                        default:
                            v76 := GenS7{F0: v32, F1: fi74}
                            _ = v76
                        }
                        v77 := GenS2{F0: 1}
                        _ = fi74
                        _ = v77
                    }
                    var v78 uint32 = 65535
                    _ = v78
                }
            }
            _ = fi50
            _ = v51
        }
        for fi79 := 0; fi79 < 8; fi79++ {
            var v80 int16 = -1
            for fi81 := 0; fi81 < 1; fi81++ {
                v29 = (42 == 42)
                var v82 bool = false
                switch -9223372036854775808 {
                case 1:
                    var v83 int32 = -2147483648
                    _ = v83
                case -205:
                    var v84 int8 = v4
                    _ = v84
                default:
                    v82 = (3.14159 < (1.5 * 0.1))
                }
                _ = fi81
                _ = v82
            }
            _ = fi79
            _ = v80
        }
        var v85 uint16 = v32
        _ = v85
    }
    v86 := v5[128]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v86
}
