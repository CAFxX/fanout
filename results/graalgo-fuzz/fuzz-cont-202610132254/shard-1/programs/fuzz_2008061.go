package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int16; F1 int; F2 int16 }
type GenS3 struct { F0 int; F1 int64; F2 string }
type GenS4 struct { F0 int16; F1 uint8 }
type GenS5 struct { F0 uint; F1 int }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 bool; F1 int8 }
type GenS8 struct { F0 int16; F1 uint; F2 int }
type GenS9 struct { F0 uint; F1 uint32; F2 uint8 }
type GenS10 struct { F0 int32; F1 uint8 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 uint }
func genFn0() string {
    var v7 int64 = 0
    var v8 uint64 = 4294967295
    switch "foo bar" {
    case "a":
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 uint16 = 1
            v11 := GenS1{F0: 38, F1: 2, F2: 2}
            switch fi9 {
            case -9223372036854775808:
                v8 = (v8 % ((18446744073709551615) | 1))
            }
            _ = fi9
            _ = v10
            _ = v11
        }
    case "foo bar":
        fmt.Println(v7)
    case "":
        switch 127 {
        case 283:
            v12 := []int{-9223372036854775808, 85, 127}
            _ = v12
        case 32767:
            v7 = 0
        }
    default:
        v13 := GenS1{F0: 65535, F1: 2, F2: 2147483647}
        _ = v13
    }
    var v14 uint32 = 0
    _ = v14
    _ = v7
    _ = v8
    return ("hello" + "a")
}
func genFn1(p0 uint8) string {
    v15 := []int{9223372036854775807, 662}
    for fi16 := 0; fi16 < 9; fi16++ {
        v15 = append(v15, 847)
        _ = len(v15)
        _ = fi16
    }
    _ = p0
    _ = v15
    return ""
}

func main() {
    var v1 uint8 = 17
    var v2 int32 = 2147483647
    var v3 uint16 = 255
    v4 := []uint32{228, 2147483648, 4294967295}
    var v5 uint16 = v3
    var v6 uint8 = 127
    v2 -= (^(v2))
    if false {
        v17 := GenS2{F0: 1000, F1: 2147483647, F2: 1}
        v3 -= v3
        if false {
            switch ("a" + "") {
            case "hello":
                v4 = append(v4, 251)
            case "\n\t":
                v1 = (v1 ^ v6)
            }
            var v18 uint32 = 480
            _ = v18
        }
        _ = v17
    } else {
        var v19 float64 = -2.25
        var v20 int16 = 1
        _ = v19
        _ = v20
    }
    if len(v4) > 0 {
        v4[0] = 0
    }
    var v21 uint16 = v5
    var v22 uint32 = 0
    fmt.Println(v21, v3, v21)
    var v23 bool = false
    v24 := GenS3{F0: 994, F1: 983, F2: "hello"}
    v22 = (^(v22))
    for fi25 := 0; fi25 < 5; fi25++ {
        var v26 int16 = 1
        _ = fi25
        _ = v26
    }
    switch -363 {
    case 32767:
        var v27 int16 = 775
        _ = v27
    default:
        switch 127 {
        case 4294967295:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    }
    for fi28 := 0; fi28 < 7; fi28++ {
        if (-79 > 100) {
            var v29 uint16 = (v21 | v21)
            _ = v29
        }
        switch fi28 {
        case 32767:
            _ = len(v4)
        case 0:
            if (v23 && false) {
                for fi30 := 0; fi30 < 13; fi30++ {
                    v21 -= 32767
                    _ = fi30
                }
                for fi31 := 0; fi31 < 20; fi31++ {
                    var v32 uint32 = (v22 | (^(v22)))
                    v22 = v22
                    _ = fi31
                    _ = v32
                }
            } else {
                v33 := GenS4{F0: 1, F1: v6}
                if v23 {
                    switch 457 {
                    case 127:
                        switch "hello" {
                        case "a":
                            for fi34 := 0; fi34 < 5; fi34++ {
                                v22 += 65535
                                _ = fi34
                            }
                        default:
                            v1 = (v1 - v6)
                        }
                    default:
                        var v35 int32 = -1
                        _ = v35
                    }
                    v36 := GenS2{F0: -32768, F1: int(v6), F2: -746}
                    _ = v36
                } else {
                    if (0 != 42) {
                        if v23 {
                            for fi37 := 0; fi37 < 18; fi37++ {
                                v38 := GenS5{F0: 999, F1: (fi37 / (((-(fi37))) | 1))}
                                _ = fi37
                                _ = v38
                            }
                            for fi39 := 0; fi39 < 5; fi39++ {
                                _ = len(v4)
                                _ = fi39
                            }
                            if (v23 || (v2 != v2)) {
                                var v40 float64 = ((0.1 + 3.14159) - float64((1.5 * 0.0)))
                                if false {
                                    v41 := GenS6{F0: -934}
                                    _ = v41
                                } else {
                                    v42 := GenS7{F0: false, F1: 1}
                                    var v43 int32 = (v2 % ((int32(v1)) | 1))
                                    _ = v42
                                    _ = v43
                                }
                                _ = v40
                            }
                        } else {
                            if (2147483647 != -936) {
                                var v44 bool = v23
                                _ = v44
                            }
                        }
                    } else {
                        var v45 int = 32768
                        switch "a" {
                        case "a":
                            switch ("a" + "hello") {
                            case "foo bar":
                                v3 += v3
                            default:
                                v45 = 128
                            }
                        case "foo bar":
                            for fi46 := 0; fi46 < 4; fi46++ {
                                v23 = v23
                                v47 := GenS7{F0: ((v5 | (v21 - v21)) != v21), F1: int8(v45)}
                                v5 -= v3
                                _ = fi46
                                _ = v47
                            }
                        case "hello":
                            var v48 float64 = ((1e10 * (3.14159 - 0.1)) - (1.5 / (3.14159 + -2.25)))
                            _ = v48
                        }
                        for fi49 := 0; fi49 < 16; fi49++ {
                            v50 := GenS2{F0: 0, F1: fi49, F2: -116}
                            _ = fi49
                            _ = v50
                        }
                        _ = v45
                    }
                    switch (("foo bar" + "") + ("foo bar" + "")) {
                    case "a":
                        var v51 int64 = int64(v6)
                        _ = v51
                    case "hello":
                        v52 := GenS8{F0: -32768, F1: 147, F2: 32768}
                        _ = v52
                    case "\n\t":
                        v53 := GenS5{F0: 42, F1: fi28}
                        _ = v53
                    default:
                        if (32767 >= (fi28 + fi28)) {
                            switch "a" {
                            case "":
                                _ = len(v4)
                            case "a":
                                if (42 == int16(fi28)) {
                                    for fi54 := 0; fi54 < 19; fi54++ {
                                        v4 = append(v4, (^(v22)))
                                        _ = fi54
                                    }
                                } else {
                                    if len(v4) > 0 {
                                        v4[0] = 1
                                    }
                                }
                            default:
                                v55 := GenS9{F0: 0, F1: v22, F2: uint8(v5)}
                                _ = v55
                            }
                        }
                    }
                    v56 := GenS4{F0: 0, F1: (v6 & 255)}
                    _ = v56
                }
                if (true || v23) {
                    if len(v4) > 0 {
                        v4[0] = uint32(v1)
                    }
                    fmt.Println(v6, v2)
                }
                _ = v33
            }
        default:
            for fi57 := 0; fi57 < 20; fi57++ {
                if ((fi28 & -490) != fi57) {
                    var v58 uint64 = 95
                    v59 := GenS10{F0: v2, F1: v1}
                    _ = v58
                    _ = v59
                } else {
                    var v60 int8 = -62
                    switch "hello" {
                    case "":
                        switch "hello" {
                        case "foo bar":
                            switch fi57 {
                            case -2147483648:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            default:
                                v61 := GenS11{F0: (v22 * 831)}
                                _ = v61
                            }
                        default:
                            var v62 bool = ((v21 ^ uint16(v1)) == (v21 / ((v5) | 1)))
                            _ = v62
                        }
                    case "foo bar":
                        v60 = v60
                    case "\n\t":
                        fmt.Println(v23, v22, v23)
                    default:
                        if v23 {
                            var v63 float64 = 1e10
                            if false {
                                v3 += v3
                                switch "" {
                                case "\n\t":
                                    v60 -= v60
                                default:
                                    var v64 int = (^(fi57))
                                    _ = v64
                                }
                                v60 += (v60 ^ v60)
                            } else {
                                switch ((("" + "a") + ("a" + "hello")) + "") {
                                case "hello":
                                    switch ("foo bar" + "\n\t") {
                                    case "\n\t":
                                        v65 := GenS10{F0: 0, F1: v1}
                                        _ = v65
                                    case "":
                                        if (!true) {
                                            v23 = v23
                                        }
                                    default:
                                        v4 = append(v4, v22)
                                    }
                                case "foo bar":
                                    if len(v4) > 0 {
                                        v4[0] = (v22 % ((v22) | 1))
                                    }
                                case "":
                                    switch "" {
                                    case "a":
                                        var v66 int = fi28
                                        _ = v66
                                    case "foo bar":
                                        v63 -= 3.14159
                                    default:
                                        fmt.Println(v21, fi57)
                                        v5 = (v5 | v21)
                                    }
                                }
                            }
                            var v67 uint8 = (v6 << ((42) & 63))
                            v1 -= (v1 - 1)
                            _ = v63
                            _ = v67
                        } else {
                            if v23 {
                                v68 := GenS6{F0: 69}
                                _ = v68
                            } else {
                                if v23 {
                                    v69 := GenS8{F0: -468, F1: 1, F2: int(v60)}
                                    var v70 float64 = 0.0
                                    var v71 uint32 = v22
                                    _ = v69
                                    _ = v70
                                    _ = v71
                                } else {
                                    if len(v4) > 0 {
                                        v4[0] = v22
                                    }
                                    var v72 string = ("hello" + "a")
                                    if true {
                                        fmt.Println(v22, v72)
                                        _ = len(v4)
                                    } else {
                                        switch fi28 {
                                        case 243:
                                            var v73 uint16 = 255
                                            _ = v73
                                        }
                                        if (888 != 65535) {
                                            v74 := GenS12{F0: 255}
                                            _ = v74
                                        } else {
                                            var v75 uint32 = (v22 * v22)
                                            _ = v75
                                        }
                                    }
                                    var v76 int = fi57
                                    _ = v72
                                    _ = v76
                                }
                                if true {
                                    v4 = append(v4, (v22 << ((255) & 63)))
                                    v1 -= v1
                                }
                            }
                        }
                    }
                    _ = v60
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi57
            }
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi28
    }
    var v77 int64 = -889
    var v78 uint = 1
    switch 1 {
    case 127:
        v22 = v22
    case 1:
        fmt.Println(v23, v2)
    default:
        var v79 float32 = (3.14159 / -2.25)
        _ = v79
    }
    fmt.Println(v78, v22, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v77
    _ = v78
}
