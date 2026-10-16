package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 uint16 }
type GenS2 struct { F0 uint8; F1 bool; F2 uint32 }
type GenS3 struct { F0 int8; F1 int; F2 uint16 }
type GenS4 struct { F0 int64; F1 uint8; F2 uint }
type GenS5 struct { F0 int8; F1 uint8; F2 int64 }
type GenS6 struct { F0 int32; F1 uint16 }
type GenS7 struct { F0 int64; F1 uint32 }
type GenS8 struct { F0 uint; F1 uint8; F2 int8 }
type GenS9 struct { F0 int; F1 int64; F2 uint8 }
func genFn0() int32 {
    v6 := GenS1{F0: (true && (49 != 4294967295)), F1: 2147483647, F2: 65535}
    var v7 uint64 = 0
    v8 := []int{0, -177, 127}
    fmt.Println(v7)
    _ = v6
    _ = v7
    _ = v8
    return -2147483648
}
func genFn1(p0 int64, p1 uint16, p2 uint16) int64 {
    if (2147483647 >= int(p2)) {
        for fi9 := 0; fi9 < 4; fi9++ {
            fmt.Println(fi9)
            p2 = (p2 / ((32767) | 1))
            _ = fi9
        }
        var v10 int32 = 1
        for fi11 := 0; fi11 < 10; fi11++ {
            v10 = v10
            if (203 < 1) {
                fmt.Println(v10)
                v12 := GenS2{F0: 42, F1: (524 < 42), F2: 870}
                var v13 uint = 65535
                _ = v12
                _ = v13
            } else {
                fmt.Println(v10, p0, p0)
            }
            var v14 int = fi11
            _ = fi11
            _ = v14
        }
        _ = v10
    }
    p2 += p1
    _ = p0
    _ = p1
    _ = p2
    return -1
}

func main() {
    var v1 int16 = 1000
    var v2 int8 = -128
    var v3 bool = ((true || false) && ((("" <= ("" + "a")) && true) && ((!false) && false)))
    v4 := []uint32{}
    var v5 int32 = 2147483647
    var v15 uint16 = 65535
    if (18446744073709551615 >= 18446744073709551615) {
        if (202 != 2) {
            v15 = v15
            var v16 uint32 = 4294967295
            if v3 {
                var v17 string = "foo bar"
                _ = v17
            }
            _ = v16
        } else {
            if v3 {
                if false {
                    v15 = (v15 << ((729) & 63))
                    for fi18 := 0; fi18 < 11; fi18++ {
                        var v19 float32 = (0.1 * (0.1 * (3.14159 + -2.25)))
                        if (int64(v15) == 9223372036854775807) {
                            for fi20 := 0; fi20 < 5; fi20++ {
                                if v3 {
                                    _ = len(v4)
                                    var v21 string = "foo bar"
                                    _ = v21
                                } else {
                                    var v22 uint32 = 163
                                    _ = v22
                                }
                                v3 = v3
                                fmt.Println(v15)
                                _ = fi20
                            }
                            var v23 uint16 = 42
                            _ = v23
                        }
                        var v24 int = fi18
                        _ = fi18
                        _ = v19
                        _ = v24
                    }
                    for fi25 := 0; fi25 < 10; fi25++ {
                        v15 += 65535
                        fmt.Println(v3, v15)
                        v26 := GenS2{F0: 9, F1: v3, F2: 1}
                        _ = fi25
                        _ = v26
                    }
                }
                if (v3 || (v3 || (-9223372036854775808 < -568))) {
                    v27 := GenS2{F0: 42, F1: v3, F2: 42}
                    if true {
                        v5 = genFn0()
                        var v28 uint8 = 255
                        _ = v28
                    }
                    _ = v27
                }
            } else {
                fmt.Println(v3, v1, v5)
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v2 = (v2 * v2)
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            fmt.Println(v1)
        }
        var v29 uint16 = (^(v15))
        _ = v29
    } else {
        if v3 {
            if (!v3) {
                if (v2 < (v2 * (v2 >> ((65535) & 63)))) {
                    switch 2147483647 {
                    case 4294967295:
                        if true {
                            _ = len(v4)
                            var v30 uint64 = 0
                            _ = v30
                        }
                    case -9223372036854775808:
                        for fi31 := 0; fi31 < 7; fi31++ {
                            var v32 int8 = 42
                            var v33 string = "a"
                            _ = fi31
                            _ = v32
                            _ = v33
                        }
                    }
                    var v34 int64 = 9223372036854775807
                    _ = v34
                } else {
                    var v35 uint16 = 42
                    switch "a" {
                    case "foo bar":
                        if v3 {
                            if (-2.25 == 3.14159) {
                                var v36 string = "\n\t"
                                v37 := GenS1{F0: v3, F1: v5, F2: (^(v35))}
                                _ = v36
                                _ = v37
                            }
                            for fi38 := 0; fi38 < 11; fi38++ {
                                v35 -= v35
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                for fi39 := 0; fi39 < 9; fi39++ {
                                    var v40 int16 = (v1 | (v1 + 42))
                                    _ = fi39
                                    _ = v40
                                }
                                _ = fi38
                            }
                        } else {
                            var v41 int = int(v15)
                            _ = v41
                        }
                    case "hello":
                        if (-2.25 <= 0.1) {
                            fmt.Println(v1)
                            var v42 float64 = 0.0
                            _ = v42
                        }
                    case "\n\t":
                        if (1 != 507) {
                            var v43 int32 = (v5 << ((0) & 63))
                            if false {
                                for fi44 := 0; fi44 < 5; fi44++ {
                                    fmt.Println(v15, fi44, v5)
                                    v45 := GenS3{F0: v2, F1: 857, F2: (v15 | v35)}
                                    v46 := GenS4{F0: int64(v2), F1: 175, F2: 255}
                                    _ = fi44
                                    _ = v45
                                    _ = v46
                                }
                            }
                            _ = v43
                        }
                    default:
                        fmt.Println(v35, v5)
                    }
                    fmt.Println(v1)
                    _ = v35
                }
                var v47 uint32 = 1
                var v48 string = ("hello" + "")
                _ = v47
                _ = v48
            } else {
                v15 += 42
                var v49 uint16 = 2
                _ = v49
            }
            if false {
                for fi50 := 0; fi50 < 4; fi50++ {
                    _ = len(v4)
                    fmt.Println(v1, v15)
                    if ((-2.25 <= 0.0) && true) {
                        fmt.Println(fi50, v15)
                        fmt.Println(v5, v5)
                    } else {
                        if (961 <= 1) {
                            var v51 uint16 = (v15 * (v15 >> ((0) & 63)))
                            _ = v51
                        } else {
                            if v3 {
                                if (!(2 < 253)) {
                                    if v3 {
                                        for fi52 := 0; fi52 < 7; fi52++ {
                                            v15 = 345
                                            fmt.Println(fi52, v15)
                                            v3 = (1.5 < ((3.14159 - -2.25) / 3.14159))
                                            _ = fi52
                                        }
                                        var v53 float32 = 1e10
                                        _ = v53
                                    }
                                } else {
                                    for fi54 := 0; fi54 < 16; fi54++ {
                                        v55 := GenS5{F0: v2, F1: 246, F2: -81}
                                        _ = fi54
                                        _ = v55
                                    }
                                    if (2 <= 128) {
                                        v2 += 62
                                        if true {
                                            if true {
                                                fmt.Println(v2)
                                            } else {
                                                switch "" {
                                                case "a":
                                                    if ((v15 ^ (v15 / ((323) | 1))) == (v15 >> ((0) & 63))) {
                                                        var v56 string = ""
                                                        _ = v56
                                                    }
                                                case "foo bar":
                                                    if v3 {
                                                        v1 = v1
                                                        switch ("\n\t" + "") {
                                                        case "hello":
                                                            if v3 {
                                                                if v3 {
                                                                    for fi57 := 0; fi57 < 10; fi57++ {
                                                                        fmt.Println(fi50)
                                                                        _ = fi57
                                                                    }
                                                                } else {
                                                                    var v58 int32 = (v5 >> ((uint(v15)) & 63))
                                                                    for fi59 := 0; fi59 < 12; fi59++ {
                                                                        var v60 uint8 = 0
                                                                        _ = fi59
                                                                        _ = v60
                                                                    }
                                                                    _ = v58
                                                                }
                                                                v15 = v15
                                                            }
                                                        case "\n\t":
                                                            if ((float32(0.0) == float32(((1e10 * (float64(0.1) + (3.14159 - (1e10 * 1e10)))) / 0.1))) || (255 <= 249)) {
                                                                v4 = append(v4, 2147483648)
                                                            } else {
                                                                for fi61 := 0; fi61 < 12; fi61++ {
                                                                    if (v3 && (1.5 != -2.25)) {
                                                                        var v62 int8 = (v2 / ((v2) | 1))
                                                                        v3 = (v3 || true)
                                                                        _ = v62
                                                                    }
                                                                    _ = fi61
                                                                }
                                                                v63 := GenS6{F0: -238, F1: (v15 | v15)}
                                                                _ = v63
                                                            }
                                                        default:
                                                            for fi64 := 0; fi64 < 19; fi64++ {
                                                                v65 := GenS6{F0: 2147483647, F1: v15}
                                                                _ = fi64
                                                                _ = v65
                                                            }
                                                        }
                                                        for fi66 := 0; fi66 < 3; fi66++ {
                                                            var v67 string = ("foo bar" + ("a" + ("foo bar" + "foo bar")))
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            _ = fi66
                                                            _ = v67
                                                        }
                                                    }
                                                case "hello":
                                                    if (!v3) {
                                                        var v68 int8 = v2
                                                        _ = v68
                                                    }
                                                default:
                                                    var v69 string = ""
                                                    _ = v69
                                                }
                                                if false {
                                                    for fi70 := 0; fi70 < 6; fi70++ {
                                                        var v71 uint8 = 42
                                                        var v72 int8 = v2
                                                        _ = fi70
                                                        _ = v71
                                                        _ = v72
                                                    }
                                                    switch (fi50 + 202) {
                                                    case -2147483648:
                                                        _ = len(v4)
                                                    case -1:
                                                        var v73 string = "foo bar"
                                                        _ = v73
                                                    default:
                                                        v3 = v3
                                                    }
                                                    switch ("a" + "hello") {
                                                    case "a":
                                                        v2 = v2
                                                    case "hello":
                                                        v4 = append(v4, 1)
                                                    default:
                                                        var v74 int = fi50
                                                        _ = v74
                                                    }
                                                } else {
                                                    v1 = (v1 / (((v1 - (-(v1)))) | 1))
                                                    var v75 int64 = 1
                                                    var v76 int8 = 127
                                                    _ = v75
                                                    _ = v76
                                                }
                                                for fi77 := 0; fi77 < 14; fi77++ {
                                                    v2 += v2
                                                    fmt.Println(v3)
                                                    switch (("\n\t" + "hello") + ("" + "\n\t")) {
                                                    case "":
                                                        v78 := GenS1{F0: v3, F1: genFn0(), F2: 0}
                                                        _ = v78
                                                    case "hello":
                                                        v79 := GenS4{F0: 510, F1: 84, F2: 686}
                                                        _ = v79
                                                    case "foo bar":
                                                        fmt.Println(fi50, fi77)
                                                    default:
                                                        v80 := GenS4{F0: -9223372036854775808, F1: 128, F2: uint(v15)}
                                                        _ = v80
                                                    }
                                                    _ = fi77
                                                }
                                            }
                                            v5 = v5
                                            v5 += -1
                                        }
                                        var v81 int8 = v2
                                        _ = v81
                                    }
                                    var v82 int64 = genFn1(1, (v15 + (v15 >> ((1) & 63))), (v15 + 1))
                                    _ = v82
                                }
                                for fi83 := 0; fi83 < 4; fi83++ {
                                    fmt.Println(v15, fi50)
                                    v84 := GenS7{F0: 9223372036854775807, F1: 776}
                                    _ = fi83
                                    _ = v84
                                }
                                var v85 int8 = (v2 * v2)
                                _ = v85
                            }
                            v2 = v2
                            for fi86 := 0; fi86 < 20; fi86++ {
                                var v87 uint16 = 1
                                _ = fi86
                                _ = v87
                            }
                        }
                        var v88 int32 = (v5 / (((v5 & v5)) | 1))
                        _ = v88
                    }
                    _ = fi50
                }
                var v89 int = 42
                var v90 uint = 0
                _ = v89
                _ = v90
            } else {
                var v91 uint16 = v15
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v92 uint32 = 65535
                _ = v91
                _ = v92
            }
            if v3 {
                for fi93 := 0; fi93 < 20; fi93++ {
                    if v3 {
                        if true {
                            v5 = v5
                        } else {
                            fmt.Println(v3, v2)
                            v1 = v1
                        }
                        v94 := GenS8{F0: 0, F1: 1, F2: v2}
                        switch (fi93 & (fi93 | 128)) {
                        case 126:
                            var v95 uint32 = 2147483647
                            _ = v95
                        }
                        _ = v94
                    }
                    var v96 uint16 = v15
                    _ = fi93
                    _ = v96
                }
                var v97 int16 = (^(v1))
                _ = v97
            }
        }
    }
    if false {
        v3 = v3
        v4 = append(v4, 4294967295)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        v1 = (v1 - v1)
        var v98 float64 = 0.0
        for fi99 := 0; fi99 < 20; fi99++ {
            v100 := GenS9{F0: fi99, F1: 9223372036854775807, F2: 255}
            _ = fi99
            _ = v100
        }
        _ = v98
    }
    v101 := GenS4{F0: 1, F1: 2, F2: 18446744073709551615}
    var v102 uint64 = 18446744073709551615
    v103 := GenS9{F0: -705, F1: -9223372036854775808, F2: 42}
    v104 := GenS5{F0: (-(v2)), F1: 159, F2: 167}
    for fi105 := 0; fi105 < 10; fi105++ {
        fmt.Println(v3, v1, v3)
        _ = fi105
    }
    v106 := GenS1{F0: true, F1: -291, F2: v15}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch "\n\t" {
    case "hello":
        var v107 int8 = (v2 % (((v2 << ((610) & 63))) | 1))
        _ = v107
    }
    fmt.Println("v1:", v1)
    fmt.Println("v102:", v102)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v101
    _ = v102
    _ = v103
    _ = v104
    _ = v106
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
