package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 int32 }
type GenS2 struct { F0 int32; F1 uint32; F2 uint16 }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint; F1 uint8 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int; F1 uint16; F2 uint }
type GenS9 struct { F0 uint32; F1 uint8 }
type GenS10 struct { F0 uint8; F1 string }
type GenS11 struct { F0 int64; F1 int64 }
type GenS12 struct { F0 int64; F1 bool; F2 string }
type GenS13 struct { F0 uint32; F1 uint32; F2 int64 }
type GenS14 struct { F0 bool; F1 uint }
func genFn0(p0 uint, p1 int64, p2 int8) int16 {
    v7 := []int{44, -9223372036854775808}
    p0 = (p0 - 221)
    if ((true && (!true)) && (4294967295 > 2)) {
        var v8 int64 = 0
        p1 = -422
        v9 := GenS1{F0: 4294967295, F1: -1, F2: 42}
        _ = v8
        _ = v9
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -297
}
func genFn1(p0 int16) bool {
    v10 := GenS1{F0: 506, F1: -833, F2: 2147483647}
    v11 := GenS1{F0: 722, F1: 0, F2: int32(p0)}
    _ = p0
    _ = v10
    _ = v11
    return (-2147483648 <= 907)
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{0, 600, 2147483647}
    var v3 int64 = 42
    var v4 string = "a"
    var v5 float32 = float32((3.14159 / 1.0))
    v6 := make(map[uint32]uint32)
    v12 := GenS2{F0: 2147483647, F1: 718, F2: 183}
    if (true || true) {
        v4 = v4
    }
    fmt.Println(v4, v5)
    switch "a" {
    case "a":
        if (309 != 255) {
            v5 -= 0.1
        } else {
            v4 = v4
            switch v4 {
            case "\n\t":
                var v13 int = 9223372036854775807
                _ = v13
            case "hello":
                fmt.Println(v4, v3)
            }
            v3 -= -9223372036854775808
        }
    case "hello":
        var v14 int = -966
        _ = v14
    default:
        for fi15 := 0; fi15 < 12; fi15++ {
            switch fi15 {
            case 0:
                var v16 string = (v4 + ((v4 + "foo bar") + "a"))
                _ = v16
            case 14:
                v17 := GenS2{F0: 349, F1: 65535, F2: 255}
                _ = v17
            case 32768:
                var v18 int16 = 1
                _ = v18
            }
            _ = fi15
        }
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        switch (-(fi19)) {
        case 32767:
            v1[fi19] = 165
        case -785:
            var v20 int16 = 234
            _ = v20
        default:
            fmt.Println(v4)
        }
        if ((!false) || false) {
            if true {
                v4 = "hello"
                switch "a" {
                case "\n\t":
                    for fi21 := 0; fi21 < 18; fi21++ {
                        var v22 float64 = 1.5
                        _ = fi21
                        _ = v22
                    }
                case "hello":
                    var v23 int32 = -464
                    _ = v23
                }
            } else {
                var v24 uint32 = 2147483648
                switch fi19 {
                case 127:
                    v4 = v4
                default:
                    for fi25 := 0; fi25 < 6; fi25++ {
                        v5 = (v5 + v5)
                        v4 = v4
                        v3 = v3
                        _ = fi25
                    }
                }
                _ = v24
            }
            var v26 int64 = (v3 ^ v3)
            var v27 uint8 = 250
            _ = v26
            _ = v27
        } else {
            fmt.Println(v4, v3)
            v4 = "hello"
            _ = len(v2)
        }
        v4 = v4
        _ = fi19
    }
    var v28 uint64 = 2
    v6[963] = 42
    if ((v3 / (((-(v3))) | 1)) >= v3) {
        v29 := GenS3{F0: 1, F1: (((float64(v5) - 0.1) * (3.14159 + (-2.25 * 0.0))) > 1e10)}
        var v30 bool = (!false)
        _ = v29
        _ = v30
    }
    if true {
        var v31 uint16 = 42
        if false {
            v28 = (v28 * v28)
            v3 += (v3 / ((v3) | 1))
            v5 -= (0.1 / 1.5)
        } else {
            v31 = (^(v31))
        }
        _ = v31
    } else {
        v4 = ""
        for fi32 := 0; fi32 < 3; fi32++ {
            v3 += v3
            v5 = v5
            _ = fi32
        }
    }
    v4 = (v4 + (("a" + ("\n\t" + "foo bar")) + v4))
    switch 128 {
    case 32767:
        v28 -= 42
    case -559:
        if (uint(v28) < 963) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if ((-507 != 32767) && false) {
                for fi33 := 0; fi33 < 8; fi33++ {
                    v34 := GenS4{F0: v28}
                    var v35 uint8 = 128
                    _ = len(v2)
                    _ = fi33
                    _ = v34
                    _ = v35
                }
            } else {
                _ = len(v2)
                if (true && true) {
                    for fi36 := 0; fi36 < 13; fi36++ {
                        v1[(fi36 / (((fi36 - fi36)) | 1))] = 108
                        _ = fi36
                    }
                    if true {
                        switch 42 {
                        case 500:
                            switch (v4 + (v4 + "\n\t")) {
                            case "a":
                                if (((!true) || (!true)) && true) {
                                    v37 := GenS5{F0: 2}
                                    _ = v37
                                } else {
                                    for fi38 := 0; fi38 < 10; fi38++ {
                                        if ((v3 - v3) != v3) {
                                            if ((v4 >= (v4 + "\n\t")) || true) {
                                                var v39 uint16 = 255
                                                v40 := GenS6{F0: 1, F1: 92}
                                                v41 := GenS4{F0: 998}
                                                _ = v39
                                                _ = v40
                                                _ = v41
                                            } else {
                                                if (961 < 2) {
                                                    for fi42 := 0; fi42 < 12; fi42++ {
                                                        var v43 uint32 = 524
                                                        v44 := GenS7{F0: 0}
                                                        v4 = v4
                                                        _ = fi42
                                                        _ = v43
                                                        _ = v44
                                                    }
                                                    fmt.Println(v5, v3)
                                                    for fi45 := 0; fi45 < 13; fi45++ {
                                                        fmt.Println(v4)
                                                        _ = fi45
                                                    }
                                                } else {
                                                    if true {
                                                        for fi46 := 0; fi46 < 6; fi46++ {
                                                            var v47 int8 = -48
                                                            v3 += v3
                                                            delete(v6, 65535)
                                                            v6[292] = 475
                                                            var v48 int32 = 1
                                                            _ = fi46
                                                            _ = v47
                                                            _ = v48
                                                        }
                                                        switch fi38 {
                                                        case -2147483648:
                                                            v49 := GenS8{F0: -471, F1: 42, F2: 4294967295}
                                                            _ = v49
                                                        }
                                                        fmt.Println(v4)
                                                    } else {
                                                        for fi50 := 0; fi50 < 19; fi50++ {
                                                            v51 := GenS9{F0: 4294967295, F1: 127}
                                                            var v52 int = (fi38 ^ fi38)
                                                            v53 := GenS10{F0: 128, F1: v4}
                                                            _ = fi50
                                                            _ = v51
                                                            _ = v52
                                                            _ = v53
                                                        }
                                                    }
                                                }
                                                if false {
                                                    var v54 float32 = float32(3.14159)
                                                    v54 = v5
                                                    var v55 uint = 90
                                                    _ = v54
                                                    _ = v55
                                                } else {
                                                    v3 -= v3
                                                    var v56 int8 = 127
                                                    for fi57 := 0; fi57 < 19; fi57++ {
                                                        fmt.Println(v28, v28, v28)
                                                        switch -716 {
                                                        case 42:
                                                            fmt.Println(v56)
                                                        case -985:
                                                            if true {
                                                                v2 = append(v2, 42)
                                                                if false {
                                                                    v3 = v3
                                                                    v4 = v4
                                                                    if (!false) {
                                                                        v58 := GenS11{F0: v3, F1: (-(v3))}
                                                                        _ = v58
                                                                    } else {
                                                                        fmt.Println(v3, v4)
                                                                        v59 := GenS5{F0: 2}
                                                                        v5 = (v5 - v5)
                                                                        _ = v59
                                                                    }
                                                                }
                                                            }
                                                        default:
                                                            var v60 int32 = 2147483647
                                                            _ = v60
                                                        }
                                                        _ = fi57
                                                    }
                                                    _ = v56
                                                }
                                            }
                                        } else {
                                            v61 := GenS8{F0: fi38, F1: 2, F2: 0}
                                            _ = v61
                                        }
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = fi38
                                    }
                                }
                            case "":
                                v62 := GenS12{F0: v3, F1: true, F2: v4}
                                _ = v62
                            default:
                                v28 = v28
                            }
                        case 32767:
                            var v63 uint8 = 2
                            _ = v63
                        case 1:
                            if (2 < 2) {
                                var v64 string = "hello"
                                _ = v64
                            } else {
                                _ = len(v2)
                                fmt.Println(v4, v28, v28)
                            }
                        }
                    } else {
                        if (42 >= 39) {
                            fmt.Println(v28, v3, v3)
                        }
                    }
                } else {
                    switch v4 {
                    case "foo bar":
                        if (32767 > 32767) {
                            var v65 uint16 = 311
                            _ = v65
                        } else {
                            v66 := GenS13{F0: 150, F1: 42, F2: -9223372036854775808}
                            var v67 int = 127
                            v4 = ""
                            _ = v66
                            _ = v67
                        }
                    case "":
                        v5 -= 0.0
                    case "a":
                        var v68 float32 = 1e10
                        _ = v68
                    default:
                        if true {
                            v69 := GenS2{F0: -2147483648, F1: 2147483648, F2: 801}
                            v70 := GenS1{F0: (v28 << ((4294967295) & 63)), F1: -357, F2: 1}
                            delete(v6, 0)
                            _ = v69
                            _ = v70
                        } else {
                            switch -548 {
                            case 32768:
                                v71 := GenS1{F0: v28, F1: -1, F2: 949}
                                _ = v71
                            }
                            _ = len(v2)
                        }
                    }
                }
                v5 -= -2.25
            }
            v72 := v1[1]
            _ = v72
        }
    case -9223372036854775808:
        v73 := GenS1{F0: 2, F1: -2147483648, F2: -1}
        _ = v73
    default:
        switch 32768 {
        case 2147483647:
            if ((!false) || (-547 > -555)) {
                fmt.Println(v3, v28, v28)
            } else {
                v74 := GenS14{F0: (((!(0 != 127)) || ((!true) || (!true))) || true), F1: 2}
                for fi75 := 0; fi75 < 17; fi75++ {
                    var v76 bool = true
                    _ = fi75
                    _ = v76
                }
                _ = v74
            }
        }
    }
    v5 += v5
    for fi77 := 0; fi77 < 20; fi77++ {
        v28 = v28
        _ = len(v1)
        _ = fi77
    }
    for fi78 := 0; fi78 < 8; fi78++ {
        v28 += (v28 | 445)
        var v79 int8 = 100
        if false {
            v28 -= v28
            var v80 uint = 404
            _ = v80
        }
        _ = fi78
        _ = v79
    }
    var v81 string = (v4 + v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v81:", v81)
    _ = v1
    _ = v12
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v81
}
