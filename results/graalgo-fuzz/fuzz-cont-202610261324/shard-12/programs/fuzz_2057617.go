package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 uint32; F1 int64; F2 bool }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint16; F1 string }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint; F1 int8 }
func genFn0(p0 uint32, p1 int32, p2 int64) uint16 {
    var v5 uint32 = p0
    if true {
        if false {
            switch ("\n\t" + "\n\t") {
            case "\n\t":
                v6 := []int{32767, 2147483647}
                _ = v6
            }
            for fi7 := 0; fi7 < 20; fi7++ {
                v8 := []int{fi7, (fi7 * 127)}
                if (false && (!true)) {
                    fmt.Println(p0, p1, p2)
                } else {
                    p0 = v5
                    v5 = v5
                }
                var v9 int = (fi7 >> ((0) & 63))
                _ = fi7
                _ = v8
                _ = v9
            }
        }
        v10 := []int{42, -999}
        _ = v10
    }
    for fi11 := 0; fi11 < 3; fi11++ {
        v12 := GenS1{F0: (fi11 / ((fi11) | 1)), F1: -402}
        var v13 int64 = (-(p2))
        v14 := GenS2{F0: 255}
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    v15 := []int{}
    if (!true) {
        v5 = v5
        for fi16 := 0; fi16 < 16; fi16++ {
            fmt.Println(p1, v5)
            if true {
                var v17 uint16 = 61
                _ = v17
            } else {
                v18 := GenS2{F0: 65535}
                switch "" {
                case "hello":
                    for fi19 := 0; fi19 < 19; fi19++ {
                        switch "foo bar" {
                        case "hello":
                            if (100 != 0) {
                                v5 = (v5 * p0)
                            } else {
                                if len(v15) > 0 {
                                    v15[0] = fi19
                                }
                                switch (fi16 >> ((255) & 63)) {
                                case 179:
                                    v5 -= v5
                                case 4294967295:
                                    var v20 int16 = -32768
                                    _ = v20
                                case 0:
                                    p0 = (v5 ^ 2)
                                default:
                                    fmt.Println(fi16, fi19)
                                }
                            }
                        case "foo bar":
                            v21 := GenS1{F0: fi19, F1: (fi19 >> ((155) & 63))}
                            _ = v21
                        case "\n\t":
                            v22 := GenS1{F0: fi19, F1: (-(fi16))}
                            _ = v22
                        }
                        if (!(-48 < 127)) {
                            if false {
                                if (42 <= 4294967295) {
                                    if false {
                                        switch "" {
                                        case "\n\t":
                                            v23 := GenS2{F0: 18446744073709551615}
                                            _ = v23
                                        case "a":
                                            fmt.Println(fi16, p1, v5)
                                        default:
                                            v5 = 4294967295
                                        }
                                    } else {
                                        p2 = (p2 % ((-106) | 1))
                                        if len(v15) > 0 {
                                            v15[0] = 42
                                        }
                                        fmt.Println(p2, fi16, fi16)
                                    }
                                    if len(v15) >= 2 {
                                        v15 = v15[1:len(v15)]
                                    }
                                } else {
                                    v5 += (v5 & v5)
                                    if (2 <= 9) {
                                        if false {
                                            if false {
                                                var v24 float32 = -2.25
                                                v25 := GenS3{F0: 74, F1: 255}
                                                fmt.Println(v5, v5)
                                                fmt.Println(fi16, p0, p2)
                                                _ = v24
                                                _ = v25
                                            } else {
                                                v26 := GenS2{F0: 0}
                                                if len(v15) > 0 {
                                                    v15[0] = fi16
                                                }
                                                if false {
                                                    var v27 float64 = (3.14159 + 1.5)
                                                    v27 -= float64(((0.0 - 1e10) / -2.25))
                                                    v5 -= v5
                                                    v28 := GenS4{F0: 0, F1: (p2 >> ((1) & 63)), F2: (127 != -47)}
                                                    _ = v27
                                                    _ = v28
                                                }
                                                _ = v26
                                            }
                                        }
                                    }
                                    var v29 uint8 = 128
                                    _ = v29
                                }
                            }
                            p1 += p1
                        }
                        _ = fi19
                    }
                default:
                    var v30 bool = (65535 >= 255)
                    _ = v30
                }
                switch (fi16 + fi16) {
                case -9223372036854775808:
                    var v31 int64 = p2
                    _ = v31
                default:
                    for fi32 := 0; fi32 < 7; fi32++ {
                        switch fi32 {
                        case -588:
                            var v33 uint64 = uint64(p0)
                            if len(v15) >= 2 {
                                v15 = v15[1:len(v15)]
                            }
                            var v34 uint32 = (p0 % ((683) | 1))
                            _ = v33
                            _ = v34
                        case -714:
                            v35 := GenS5{F0: 4294967295}
                            _ = v35
                        case 225:
                            p2 = p2
                        }
                        p0 = (v5 | p0)
                        _ = fi32
                    }
                }
                _ = v18
            }
            _ = fi16
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v5
    return 401
}

func main() {
    var v1 int32 = 751
    v2 := make(map[uint32]string)
    var v3 uint64 = 0
    v4 := []uint32{2147483647}
    if len(v4) > 0 {
        v4[0] = 94
    }
    for fi36 := 0; fi36 < 3; fi36++ {
        var v37 uint64 = v3
        delete(v2, 2)
        fmt.Println(v37, v3, v3)
        _ = fi36
        _ = v37
    }
    switch 32767 {
    case -284:
        switch "\n\t" {
        case "":
            if ((v3 <= (^(v3))) || (v1 < (v1 << ((1) & 63)))) {
                if ((!true) || (-1 == 42)) {
                    var v38 uint = 2
                    _ = v38
                }
                switch -657 {
                case 1:
                    for fi39 := 0; fi39 < 10; fi39++ {
                        v3 -= 18446744073709551615
                        var v40 int32 = v1
                        v3 = v3
                        _ = fi39
                        _ = v40
                    }
                case 0:
                    v1 += -1
                default:
                    switch 32768 {
                    case 127:
                        v41 := GenS1{F0: 42, F1: -2147483648}
                        _ = v41
                    case 32767:
                        delete(v2, 42)
                    case 32768:
                        switch "foo bar" {
                        case "a":
                            var v42 uint64 = v3
                            _ = v42
                        case "":
                            var v43 int32 = int32(v3)
                            _ = v43
                        case "hello":
                            switch 0 {
                            case 1:
                                if (-2.25 == 3.14159) {
                                    v1 += -49
                                } else {
                                    for fi44 := 0; fi44 < 16; fi44++ {
                                        var v45 uint32 = 419
                                        _ = fi44
                                        _ = v45
                                    }
                                    v2[978] = "foo bar"
                                }
                            case 4294967295:
                                var v46 string = ("a" + "a")
                                _ = v46
                            default:
                                switch -482 {
                                case -594:
                                    v3 += v3
                                default:
                                    switch -9223372036854775808 {
                                    case -178:
                                        var v47 uint32 = 531
                                        _ = v47
                                    case -9223372036854775808:
                                        fmt.Println(v1, v1)
                                    case 4294967295:
                                        switch 128 {
                                        case -910:
                                            var v48 bool = ((true && false) || true)
                                            _ = v48
                                        case 127:
                                            switch 38 {
                                            case 4294967295:
                                                switch "\n\t" {
                                                case "\n\t":
                                                    var v49 uint = 255
                                                    _ = v49
                                                }
                                            }
                                        default:
                                            fmt.Println(v1)
                                        }
                                    default:
                                        v3 = v3
                                    }
                                }
                            }
                        }
                    }
                }
            }
        case "foo bar":
            var v50 bool = (((("\n\t" + ("hello" + "a")) + "\n\t") != ("hello" + "hello")) && true)
            _ = v50
        case "a":
            var v51 string = "hello"
            _ = v51
        }
    }
    if false {
        var v52 uint8 = 0
        var v53 uint32 = 2147483647
        _ = v52
        _ = v53
    } else {
        if false {
            _ = len(v4)
        }
        v1 = v1
    }
    if false {
        var v54 string = ""
        v54 = "foo bar"
        _ = v54
    }
    for fi55 := 0; fi55 < 5; fi55++ {
        v56 := GenS1{F0: int(v1), F1: (fi55 % ((fi55) | 1))}
        _ = fi55
        _ = v56
    }
    for fi57 := 0; fi57 < 17; fi57++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if ((^(fi57)) != fi57) {
            if (v1 != 0) {
                var v58 string = "hello"
                if true {
                    var v59 bool = false
                    v3 = (v3 * v3)
                    _ = v59
                }
                _ = v58
            } else {
                var v60 int16 = int16(v3)
                fmt.Println(fi57, fi57, v1)
                if (((v60 % ((-1) | 1)) != v60) && true) {
                    v61 := GenS6{F0: 255, F1: ((("hello" + "foo bar") + ("foo bar" + (("hello" + "\n\t") + "hello"))) + (("a" + "") + ("" + ("foo bar" + ("" + "")))))}
                    _ = v61
                }
                _ = v60
            }
            var v62 int8 = 127
            _ = v62
        } else {
            v1 += (v1 * v1)
            v63 := GenS6{F0: 65535, F1: ("foo bar" + ("" + ""))}
            switch 32768 {
            case -121:
                var v64 uint16 = 65535
                _ = v64
            }
            _ = v63
        }
        _ = fi57
    }
    var v65 float32 = 1e10
    var v66 uint16 = 0
    v3 += (v3 & (v3 | (v3 % ((v3) | 1))))
    var v67 bool = true
    v4 = append(v4, 2147483648)
    var v68 uint64 = 2
    switch "foo bar" {
    case "a":
        var v69 int32 = -2147483648
        _ = v69
    case "hello":
        if len(v4) > 0 {
            v4[0] = 42
        }
    case "":
        if v67 {
            switch 128 {
            case 32767:
                if (false && v67) {
                    for fi70 := 0; fi70 < 20; fi70++ {
                        for fi71 := 0; fi71 < 4; fi71++ {
                            if v67 {
                                v3 += (v3 + v3)
                                if v67 {
                                    v68 = (v68 / ((v68) | 1))
                                    v4 = append(v4, 1)
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                            } else {
                                switch "hello" {
                                case "":
                                    v2[4294967295] = "hello"
                                case "hello":
                                    fmt.Println(v66)
                                }
                                v66 = v66
                            }
                            var v72 bool = (0 < -707)
                            _ = fi71
                            _ = v72
                        }
                        for fi73 := 0; fi73 < 7; fi73++ {
                            switch "hello" {
                            case "hello":
                                var v74 float64 = 3.14159
                                _ = len(v4)
                                _ = v74
                            }
                            v75 := GenS7{F0: 41}
                            var v76 string = "hello"
                            _ = fi73
                            _ = v75
                            _ = v76
                        }
                        _ = fi70
                    }
                    _ = len(v4)
                    v77 := GenS7{F0: 0}
                    _ = v77
                } else {
                    v66 = (v66 & (^(v66)))
                    if v67 {
                        v65 -= float32(-2.25)
                        switch 710 {
                        case 994:
                            var v78 int32 = 1
                            _ = v78
                        case 32768:
                            var v79 int64 = 9223372036854775807
                            _ = v79
                        case 42:
                            _ = len(v2)
                        }
                        switch "foo bar" {
                        case "a":
                            v80 := GenS3{F0: 65535, F1: 0}
                            _ = v80
                        case "hello":
                            var v81 int64 = int64(v68)
                            _ = v81
                        case "foo bar":
                            if v67 {
                                v82 := GenS1{F0: -314, F1: 666}
                                var v83 uint8 = 127
                                v3 = v3
                                _ = v82
                                _ = v83
                            } else {
                                var v84 float64 = (3.14159 - float64(v65))
                                fmt.Println(v66)
                                _ = v84
                            }
                        default:
                            v2[42] = "\n\t"
                        }
                    } else {
                        var v85 uint64 = (v68 + v68)
                        fmt.Println(v65)
                        _ = len(v4)
                        _ = v85
                    }
                    v65 = v65
                }
            }
            fmt.Println(v1, v65)
            for fi86 := 0; fi86 < 19; fi86++ {
                if false {
                    for fi87 := 0; fi87 < 7; fi87++ {
                        v67 = v67
                        v88 := GenS2{F0: 340}
                        _ = fi87
                        _ = v88
                    }
                } else {
                    v67 = true
                    fmt.Println(v66, v68)
                }
                var v89 uint16 = (^(v66))
                v67 = (("a" + "a") <= "a")
                _ = fi86
                _ = v89
            }
        } else {
            fmt.Println(v1, v68, v3)
            for fi90 := 0; fi90 < 19; fi90++ {
                v91 := GenS5{F0: uint64(v66)}
                fmt.Println(v67, v65, fi90)
                v2[93] = "\n\t"
                _ = fi90
                _ = v91
            }
        }
    default:
        v92 := GenS1{F0: 9223372036854775807, F1: -75}
        _ = v92
    }
    if false {
        v67 = (float64((v65 - (float32(1.5) * (0.1 - v65)))) < 0.1)
        v93 := GenS6{F0: uint16(v3), F1: "\n\t"}
        for fi94 := 0; fi94 < 20; fi94++ {
            v3 = (v3 | (v68 % ((766) | 1)))
            v95 := GenS8{F0: 42, F1: 0}
            _ = fi94
            _ = v95
        }
        _ = v93
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v65
    _ = v66
    _ = v67
    _ = v68
}
