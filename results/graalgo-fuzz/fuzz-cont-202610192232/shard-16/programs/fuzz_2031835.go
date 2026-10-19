package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32; F1 int16 }
type GenS5 struct { F0 int8; F1 uint16; F2 int8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint8; F1 int8 }
type GenS8 struct { F0 int8; F1 uint; F2 bool }
type GenS9 struct { F0 int; F1 bool; F2 bool }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 string; F1 uint; F2 uint64 }
type GenS12 struct { F0 uint32 }

func main() {
    var v1 uint16 = 42
    var v2 bool = (false && true)
    var v3 uint32 = 65535
    v4 := make(map[int]uint32)
    v5 := []string{"a", "a", ("hello" + "hello")}
    v6 := []int{}
    v1 -= 840
    if false {
        v3 -= 505
        var v7 float64 = (-2.25 / -2.25)
        _ = v7
    }
    for fi8 := 0; fi8 < 16; fi8++ {
        fmt.Println(v3, fi8)
        v2 = true
        _ = fi8
    }
    var v9 string = "a"
    switch "a" {
    case "":
        var v10 float32 = (1e10 * (float32(0.0) * 0.1))
        _ = v10
    case "a":
        v11 := GenS1{F0: 4294967295, F1: (v3 << ((65535) & 63))}
        _ = v11
    default:
        v9 = v9
    }
    for fi12 := 0; fi12 < 1; fi12++ {
        v6 = append(v6, (-(fi12)))
        v13 := GenS1{F0: 18446744073709551615, F1: v3}
        fmt.Println(v1, v1)
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 3; fi14++ {
        v15 := GenS1{F0: 0, F1: v3}
        _ = fi14
        _ = v15
    }
    switch 116 {
    case 78:
        v16 := GenS2{F0: -379, F1: ("a" + v9)}
        _ = v16
    case -433:
        var v17 uint16 = (v1 ^ (v1 - (v1 - v1)))
        _ = v17
    case -940:
        v18 := GenS3{F0: 100}
        _ = v18
    }
    v9 = ("" + ("foo bar" + (v9 + ((v9 + "a") + (v9 + (v9 + ("foo bar" + v9)))))))
    var v19 uint32 = v3
    var v20 int64 = -1
    var v21 int64 = v20
    switch 863 {
    case -1:
        if v2 {
            if (v9 <= ("foo bar" + ((v9 + "") + (v9 + v9)))) {
                v21 = v21
                fmt.Println(v19, v19, v19)
            }
            var v22 string = ("hello" + ("a" + ""))
            fmt.Println(v1)
            _ = v22
        }
    case 0:
        var v23 uint8 = 1
        _ = v23
    case 32767:
        switch 32767 {
        case -9223372036854775808:
            for fi24 := 0; fi24 < 5; fi24++ {
                var v25 int8 = 127
                if false {
                    fmt.Println(v1)
                    switch 0 {
                    case -9223372036854775808:
                        for fi26 := 0; fi26 < 4; fi26++ {
                            var v27 int64 = (v21 % (((v20 | (^(v20)))) | 1))
                            var v28 string = "foo bar"
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                    }
                }
                _ = fi24
                _ = v25
            }
        case 321:
            var v29 float32 = -2.25
            _ = v29
        default:
            v30 := GenS4{F0: 65535, F1: 460}
            _ = v30
        }
    default:
        v31 := GenS1{F0: 1, F1: (v3 & 655)}
        _ = v31
    }
    for fi32 := 0; fi32 < 7; fi32++ {
        if v2 {
            var v33 int16 = -923
            if (!v2) {
                var v34 uint16 = (v1 | (v1 & 286))
                _ = v34
            } else {
                if (float32((0.1 / (1.5 / 1.0))) == (1.5 * ((((0.0 + 3.14159) - 1e10) + -2.25) * 0.0))) {
                    var v35 uint = uint(v1)
                    switch fi32 {
                    case 532:
                        delete(v4, fi32)
                    case 9223372036854775807:
                        if ((!v2) && v2) {
                            for fi36 := 0; fi36 < 14; fi36++ {
                                var v37 int16 = 42
                                var v38 string = ((((v9 + "a") + v9) + "a") + v9)
                                switch fi32 {
                                case 9223372036854775807:
                                    v1 = uint16(v35)
                                default:
                                    switch (-(fi32)) {
                                    case -1:
                                        v37 += v37
                                    }
                                }
                                _ = fi36
                                _ = v37
                                _ = v38
                            }
                            for fi39 := 0; fi39 < 10; fi39++ {
                                v40 := GenS4{F0: (v3 << ((18446744073709551615) & 63)), F1: (v33 >> (((v35 << ((255) & 63))) & 63))}
                                v41 := v4[(fi32 ^ (fi32 - (-(fi32))))]
                                if (!(false || v2)) {
                                    v2 = false
                                    fmt.Println(v9)
                                    var v42 bool = v2
                                    _ = v42
                                }
                                _ = fi39
                                _ = v40
                                _ = v41
                            }
                        }
                    case 32768:
                        for fi43 := 0; fi43 < 6; fi43++ {
                            if v2 {
                                fmt.Println(v3)
                            } else {
                                var v44 int64 = (v21 - (v21 << ((0) & 63)))
                                _ = v44
                            }
                            if (false || false) {
                                var v45 bool = v2
                                var v46 int32 = 2147483647
                                if (!v45) {
                                    v47 := GenS3{F0: -1}
                                    _ = v47
                                }
                                _ = len(v4)
                                _ = v45
                                _ = v46
                            } else {
                                fmt.Println(v19, v35, v19)
                            }
                            v6 = append(v6, (fi32 % ((fi43) | 1)))
                            _ = fi43
                        }
                    default:
                        var v48 uint8 = 127
                        _ = v48
                    }
                    if (v2 || ("" != v9)) {
                        v2 = (v35 > v35)
                        switch fi32 {
                        case -9223372036854775808:
                            var v49 string = v9
                            _ = v49
                        }
                    }
                    _ = v35
                } else {
                    var v50 string = v9
                    if false {
                        v1 += (v1 - v1)
                        if (255 != 65535) {
                            v3 += v19
                            var v51 bool = (!v2)
                            _ = len(v4)
                            _ = v51
                        } else {
                            v52 := GenS1{F0: 0, F1: (v3 & v3)}
                            var v53 int16 = 42
                            _ = v52
                            _ = v53
                        }
                        v54 := v4[fi32]
                        _ = v54
                    } else {
                        for fi55 := 0; fi55 < 10; fi55++ {
                            v50 = "foo bar"
                            switch v9 {
                            case "hello":
                                v56 := GenS5{F0: 127, F1: v1, F2: -18}
                                _ = v56
                            case "foo bar":
                                v57 := GenS6{F0: (v1 | v1)}
                                _ = v57
                            }
                            var v58 int8 = -23
                            v3 = v19
                            if len(v5) > 0 {
                                v5[0] = (("foo bar" + ((v50 + v9) + v50)) + v9)
                            }
                            _ = fi55
                            _ = v58
                        }
                        fmt.Println(v33, v3, v3)
                    }
                    var v59 int32 = 0
                    _ = v50
                    _ = v59
                }
                var v60 string = "\n\t"
                v21 = (v20 & v20)
                _ = v60
            }
            v9 = "foo bar"
            _ = v33
        } else {
            if v2 {
                fmt.Println(v2)
            }
            if (!(!true)) {
                var v61 uint32 = (v19 >> ((869) & 63))
                var v62 int = (fi32 ^ fi32)
                switch (fi32 / ((v62) | 1)) {
                case 730:
                    if (v2 || (v2 || v2)) {
                        var v63 uint64 = 2
                        for fi64 := 0; fi64 < 10; fi64++ {
                            switch (v62 & fi64) {
                            case 159:
                                if (true || true) {
                                    var v65 bool = ((v20 | (-(v20))) > -706)
                                    var v66 uint16 = v1
                                    var v67 int = 127
                                    _ = v65
                                    _ = v66
                                    _ = v67
                                } else {
                                    v21 -= v20
                                    switch (("foo bar" + "\n\t") + (((v9 + (v9 + v9)) + v9) + "foo bar")) {
                                    case "foo bar":
                                        if v2 {
                                            v1 = v1
                                            v68 := v4[fi64]
                                            v2 = v2
                                            _ = v68
                                        } else {
                                            var v69 float64 = (-2.25 * ((0.1 + 0.1) * (1e10 - 1e10)))
                                            var v70 int32 = 1
                                            fmt.Println(v70, v20)
                                            _ = v69
                                            _ = v70
                                        }
                                    case "hello":
                                        var v71 uint16 = (^(v1))
                                        _ = v71
                                    case "\n\t":
                                        switch (("a" + v9) + v9) {
                                        case "foo bar":
                                            _ = len(v4)
                                        case "a":
                                            var v72 uint32 = (v61 % ((v19) | 1))
                                            _ = v72
                                        case "\n\t":
                                            v4[1] = 2
                                        }
                                    default:
                                        v73 := GenS7{F0: 127, F1: 127}
                                        _ = v73
                                    }
                                    if ((^(v3)) == v61) {
                                        if (-1 > int8(v1)) {
                                            var v74 int8 = 1
                                            _ = v74
                                        } else {
                                            v75 := GenS8{F0: int8(v20), F1: 42, F2: (("hello" + v9) > v9)}
                                            _ = v75
                                        }
                                    }
                                }
                                var v76 string = "foo bar"
                                v61 += 2147483648
                                _ = v76
                            default:
                                var v77 int8 = -128
                                _ = v77
                            }
                            var v78 uint = 255
                            _ = fi64
                            _ = v78
                        }
                        if (228 >= 682) {
                            fmt.Println(v61)
                            v79 := GenS8{F0: -24, F1: 0, F2: ((v9 + "\n\t") < ("a" + v9))}
                            _ = v79
                        }
                        _ = v63
                    }
                case 172:
                    if (true || (!(((v20 - (v21 ^ (v20 | 790))) >= int64(v3)) || false))) {
                        switch v9 {
                        case "hello":
                            v80 := GenS9{F0: (fi32 / (((fi32 | 42)) | 1)), F1: (!(!(!false))), F2: false}
                            _ = v80
                        }
                        var v81 float64 = (3.14159 - -2.25)
                        v82 := GenS6{F0: (v1 & v1)}
                        _ = v81
                        _ = v82
                    } else {
                        var v83 int32 = -222
                        var v84 float32 = 0.0
                        if true {
                            if v2 {
                                v85 := GenS4{F0: (v3 & (v3 >> ((2) & 63))), F1: 42}
                                if len(v6) > 0 {
                                    v6[0] = fi32
                                }
                                v5 = append(v5, v9)
                                _ = v85
                            } else {
                                var v86 bool = (!(32767 == (v62 * (fi32 % ((-506) | 1)))))
                                _ = v86
                            }
                            for fi87 := 0; fi87 < 18; fi87++ {
                                v88 := GenS10{F0: -105}
                                _ = fi87
                                _ = v88
                            }
                            switch v9 {
                            case "a":
                                v89 := GenS2{F0: 307, F1: "\n\t"}
                                _ = v89
                            case "\n\t":
                                fmt.Println(v1, v21)
                            case "foo bar":
                                if v2 {
                                    if (-205 > 813) {
                                        if v2 {
                                            v61 = v61
                                        }
                                        if true {
                                            v83 += v83
                                            delete(v4, 0)
                                            v4[-198] = (v19 | (v3 / (((v3 | v3)) | 1)))
                                        }
                                        var v90 float32 = float32(0.1)
                                        _ = v90
                                    } else {
                                        switch 32767 {
                                        case -81:
                                            v62 = (-(fi32))
                                        case -906:
                                            _ = len(v5)
                                        case -890:
                                            v83 -= (v83 - (v83 * v83))
                                        default:
                                            v91 := GenS5{F0: 100, F1: v1, F2: 127}
                                            _ = v91
                                        }
                                        v92 := GenS11{F0: v9, F1: 65535, F2: 0}
                                        _ = v92
                                    }
                                }
                            default:
                                v19 += 1
                            }
                        }
                        _ = v83
                        _ = v84
                    }
                case -9223372036854775808:
                    for fi93 := 0; fi93 < 1; fi93++ {
                        fmt.Println(v1)
                        v94 := GenS7{F0: 104, F1: 83}
                        _ = fi93
                        _ = v94
                    }
                }
                _ = v61
                _ = v62
            }
        }
        v95 := GenS8{F0: int8(v1), F1: 42, F2: (-72 >= -85)}
        for fi96 := 0; fi96 < 2; fi96++ {
            if v2 {
                fmt.Println(v1, v20)
                v97 := GenS12{F0: 2}
                var v98 int16 = 1000
                _ = v97
                _ = v98
            } else {
                if (!v2) {
                    var v99 uint = 933
                    _ = v99
                }
                var v100 float32 = -2.25
                v101 := GenS10{F0: 42}
                _ = v100
                _ = v101
            }
            switch (^(fi32)) {
            case -1:
                delete(v4, fi32)
            case -186:
                v21 += (v20 * (v21 % ((v20) | 1)))
            }
            v19 = v3
            _ = fi96
        }
        _ = fi32
        _ = v95
    }
    v20 = v20
    var v102 float32 = (float32((3.14159 + 0.0)) / 1.5)
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v102:", v102)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v102
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
