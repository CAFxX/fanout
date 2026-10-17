package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint64; F1 int8 }
type GenS8 struct { F0 string; F1 uint }
type GenS9 struct { F0 uint16; F1 int64 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 string; F1 uint8; F2 int64 }

func main() {
    var v1 int = 1
    var v2 int16 = 502
    v3 := make(map[uint32]string)
    var v4 bool = true
    var v5 bool = true
    v6 := make(map[uint32]bool)
    v7 := GenS1{F0: -144, F1: 65535}
    var v8 float32 = 1.5
    if (1e10 >= (1e10 / 1e10)) {
        if (!((^(v2)) != v2)) {
            switch (("hello" + "hello") + "\n\t") {
            case "hello":
                v5 = v4
            case "foo bar":
                var v9 uint64 = 4294967295
                _ = v9
            case "":
                v2 = (v2 | v2)
            }
            v10 := GenS2{F0: 431, F1: 18446744073709551615}
            _ = v10
        } else {
            var v11 uint32 = 0
            if v5 {
                v12 := GenS3{F0: -30}
                _ = v12
            } else {
                for fi13 := 0; fi13 < 6; fi13++ {
                    v14 := GenS4{F0: (v2 >> ((18446744073709551615) & 63))}
                    if v4 {
                        fmt.Println(v4, v5, v1)
                        for fi15 := 0; fi15 < 2; fi15++ {
                            if (!v5) {
                                switch int(v11) {
                                case 136:
                                    var v16 float32 = (1e10 * v8)
                                    _ = v16
                                case -367:
                                    if v5 {
                                        if (855 == 32768) {
                                            var v17 uint32 = v11
                                            v8 = (v8 / v8)
                                            var v18 bool = false
                                            _ = v17
                                            _ = v18
                                        } else {
                                            v11 += (v11 + v11)
                                            v19 := []int{42, v1}
                                            _ = v19
                                        }
                                        switch (fi13 << ((65535) & 63)) {
                                        case 9223372036854775807:
                                            v20 := GenS3{F0: 74}
                                            _ = v20
                                        default:
                                            v21 := GenS5{F0: 1}
                                            _ = v21
                                        }
                                    } else {
                                        fmt.Println(v8, v2)
                                        switch (fi13 + v1) {
                                        case 32768:
                                            switch "hello" {
                                            case "\n\t":
                                                v22 := GenS5{F0: 209}
                                                _ = v22
                                            case "hello":
                                                var v23 uint32 = (^(v11))
                                                _ = v23
                                            case "a":
                                                if false {
                                                    var v24 bool = false
                                                    if true {
                                                        _ = len(v3)
                                                        if ((float64(v8) / 1.0) < 3.14159) {
                                                            v25 := []int{-1, fi13}
                                                            v8 -= -2.25
                                                            v2 = (v2 ^ v2)
                                                            switch (v1 / ((fi15) | 1)) {
                                                            case -1:
                                                                if false {
                                                                    if len(v25) >= 2 {
                                                                        v25 = v25[1:len(v25)]
                                                                    }
                                                                    v3[(^(v11))] = "hello"
                                                                }
                                                            default:
                                                                v11 += (v11 << ((255) & 63))
                                                                var v26 float32 = v8
                                                                _ = v26
                                                            }
                                                            _ = v25
                                                        }
                                                        if true {
                                                            if v5 {
                                                                var v27 string = "foo bar"
                                                                v28 := []int{(v1 | 1), fi13, int(v11)}
                                                                _ = v27
                                                                _ = v28
                                                            } else {
                                                                if ((!v4) || ("" != ((((("hello" + "") + "") + "") + "hello") + (((("hello" + ("foo bar" + "foo bar")) + "") + ("foo bar" + "hello")) + (("\n\t" + (("hello" + "a") + (((("" + "a") + ("hello" + ("a" + ("foo bar" + ("foo bar" + "hello"))))) + "\n\t") + ((("hello" + "\n\t") + "a") + "\n\t")))) + "\n\t"))))) {
                                                                    switch "a" {
                                                                    case "foo bar":
                                                                        if (2 > 12) {
                                                                            v29 := GenS6{F0: (v1 ^ 127)}
                                                                            v2 = v2
                                                                            var v30 uint8 = 42
                                                                            _ = v29
                                                                            _ = v30
                                                                        } else {
                                                                            var v31 uint = 2
                                                                            v32 := GenS4{F0: (-(v2))}
                                                                            _ = len(v6)
                                                                            _ = v31
                                                                            _ = v32
                                                                        }
                                                                    }
                                                                    switch fi13 {
                                                                    case -719:
                                                                        switch 127 {
                                                                        case 850:
                                                                            v33 := GenS7{F0: 8, F1: 72}
                                                                            _ = v33
                                                                        default:
                                                                            v2 -= v2
                                                                        }
                                                                    default:
                                                                        var v34 int32 = -1
                                                                        _ = v34
                                                                    }
                                                                }
                                                                v35 := v3[(v11 - v11)]
                                                                _ = v35
                                                            }
                                                            var v36 int16 = (v2 & -32768)
                                                            _ = v36
                                                        } else {
                                                            if v5 {
                                                                var v37 string = "hello"
                                                                _ = v37
                                                            } else {
                                                                v8 += float32(0.1)
                                                                if false {
                                                                    v2 = (v2 | (v2 % ((-32768) | 1)))
                                                                    if (v8 != 0.0) {
                                                                        if (!true) {
                                                                            if false {
                                                                                v38 := []int{(fi13 | -859), 32767, fi13}
                                                                                v24 = false
                                                                                _ = v38
                                                                            } else {
                                                                                var v39 int = (-(fi15))
                                                                                if (270 != 255) {
                                                                                    switch "\n\t" {
                                                                                    case "hello":
                                                                                        fmt.Println(v2)
                                                                                    default:
                                                                                        fmt.Println(v39, v8, v5)
                                                                                    }
                                                                                }
                                                                                _ = v39
                                                                            }
                                                                            v40 := GenS7{F0: 4294967295, F1: -117}
                                                                            switch "a" {
                                                                            case "foo bar":
                                                                                if (!(512 > 4294967295)) {
                                                                                    if ((1 > 42) && v5) {
                                                                                        switch (-(fi13)) {
                                                                                        case 4294967295:
                                                                                            switch "\n\t" {
                                                                                            case "":
                                                                                                var v41 uint64 = 424
                                                                                                _ = v41
                                                                                            default:
                                                                                                v11 -= 42
                                                                                            }
                                                                                        case 0:
                                                                                            v4 = v5
                                                                                        default:
                                                                                            if v4 {
                                                                                                v42 := GenS8{F0: "\n\t", F1: 19}
                                                                                                v24 = v5
                                                                                                _ = v42
                                                                                            } else {
                                                                                                v11 = (v11 >> ((360) & 63))
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        if ((((3.14159 - float64(v8)) + 0.0) == 0.0) || v5) {
                                                                                            fmt.Println(v24, v4)
                                                                                        }
                                                                                    }
                                                                                }
                                                                            case "a":
                                                                                v11 = 4294967295
                                                                            case "":
                                                                                var v43 uint64 = 530
                                                                                _ = v43
                                                                            }
                                                                            var v44 uint8 = 5
                                                                            _ = v40
                                                                            _ = v44
                                                                        }
                                                                    } else {
                                                                        v45 := GenS2{F0: 65535, F1: 0}
                                                                        _ = v45
                                                                    }
                                                                    v11 += 42
                                                                } else {
                                                                    var v46 string = "hello"
                                                                    var v47 int32 = 2147483647
                                                                    _ = v46
                                                                    _ = v47
                                                                }
                                                                if v5 {
                                                                    v5 = false
                                                                    v2 = v2
                                                                    v8 = ((v8 * (v8 * ((v8 * (float32(1.5) - v8)) + 0.1))) - float32(0.1))
                                                                }
                                                                delete(v3, v11)
                                                                if v5 {
                                                                    v11 = 2147483647
                                                                    if v5 {
                                                                        delete(v3, (v11 >> ((907) & 63)))
                                                                        v2 = -1
                                                                        var v48 int64 = 1
                                                                        _ = v48
                                                                    }
                                                                    var v49 int = v1
                                                                    _ = v49
                                                                }
                                                                v2 = 685
                                                                if v5 {
                                                                    v2 -= v2
                                                                }
                                                                var v50 string = ""
                                                                _ = v50
                                                            }
                                                        }
                                                    } else {
                                                        v51 := []int{fi15, (fi15 - -1), (-(v1))}
                                                        v11 = v11
                                                        var v52 float32 = (0.0 - (v8 * 0.0))
                                                        _ = v51
                                                        _ = v52
                                                    }
                                                    _ = v24
                                                } else {
                                                    if v5 {
                                                        fmt.Println(fi15, v5)
                                                        switch fi13 {
                                                        case -2147483648:
                                                            v11 += v11
                                                        default:
                                                            var v53 int = v1
                                                            _ = v53
                                                        }
                                                    } else {
                                                        if ((v2 % (((v2 % (((v2 & v2)) | 1))) | 1)) != 32767) {
                                                            var v54 float64 = (1.5 - ((0.1 * (0.0 / ((3.14159 / 0.1) + 3.14159))) + 0.1))
                                                            switch ("foo bar" + "hello") {
                                                            case "foo bar":
                                                                v55 := GenS9{F0: 227, F1: 1}
                                                                _ = v55
                                                            case "\n\t":
                                                                if (!false) {
                                                                    fmt.Println(fi15, v5)
                                                                    var v56 int64 = 1
                                                                    _ = v56
                                                                } else {
                                                                    var v57 int32 = 0
                                                                    var v58 int8 = 15
                                                                    v2 -= v2
                                                                    switch "hello" {
                                                                    case "hello":
                                                                        v59 := GenS3{F0: (v58 + v58)}
                                                                        _ = v59
                                                                    case "a":
                                                                        v57 -= (v57 >> ((18446744073709551615) & 63))
                                                                    }
                                                                    _ = len(v3)
                                                                    v60 := []int{fi13, (v1 % ((fi13) | 1)), 365}
                                                                    _ = v57
                                                                    _ = v58
                                                                    _ = v60
                                                                }
                                                            case "":
                                                                v4 = (v4 || (42 > 127))
                                                            default:
                                                                switch "a" {
                                                                case "\n\t":
                                                                    if v5 {
                                                                        v5 = v4
                                                                        v61 := GenS6{F0: fi13}
                                                                        _ = v61
                                                                    }
                                                                case "a":
                                                                    switch "hello" {
                                                                    case "foo bar":
                                                                        v62 := GenS4{F0: (v2 & v2)}
                                                                        _ = v62
                                                                    default:
                                                                        v63 := []int{(fi15 << ((663) & 63)), 127, v1, fi15}
                                                                        _ = v63
                                                                    }
                                                                case "hello":
                                                                    var v64 uint8 = 1
                                                                    _ = v64
                                                                }
                                                            }
                                                            _ = v54
                                                        } else {
                                                            var v65 uint = 18446744073709551615
                                                            v8 -= (v8 * (0.0 - 1e10))
                                                            v8 = float32(0.1)
                                                            _ = v65
                                                        }
                                                        if (-2147483648 > 2147483647) {
                                                            v2 -= v2
                                                        }
                                                    }
                                                    v5 = (42 == -128)
                                                }
                                            }
                                        }
                                        v11 = (v11 ^ 2147483647)
                                    }
                                case 221:
                                    v66 := []int{v1, v1, fi13, int(v11)}
                                    _ = v66
                                default:
                                    var v67 int64 = -420
                                    _ = v67
                                }
                            } else {
                                v2 = 1
                                _ = len(v6)
                                var v68 string = "hello"
                                _ = v68
                            }
                            v8 -= ((v8 / v8) / 3.14159)
                            _ = fi15
                        }
                        v69 := GenS10{F0: (v1 >> ((911) & 63))}
                        _ = v69
                    } else {
                        v11 = (v11 | v11)
                        for fi70 := 0; fi70 < 13; fi70++ {
                            v71 := GenS6{F0: fi13}
                            switch fi13 {
                            case 32767:
                                if (259 == 42) {
                                    v72 := GenS11{F0: (("a" + "hello") + ""), F1: 53, F2: -569}
                                    _ = v72
                                }
                            default:
                                var v73 float64 = 1e10
                                _ = v73
                            }
                            _ = fi70
                            _ = v71
                        }
                        v74 := GenS10{F0: (v1 >> ((2) & 63))}
                        _ = v74
                    }
                    _ = fi13
                    _ = v14
                }
                if (2 < 651) {
                    var v75 int8 = 125
                    switch 128 {
                    case 769:
                        fmt.Println(v4, v11)
                    }
                    v2 = (^(v2))
                    _ = v75
                } else {
                    fmt.Println(v2, v11)
                }
            }
            _ = v11
        }
    } else {
        fmt.Println(v8, v5)
        fmt.Println(v5)
        v2 = v2
    }
    v4 = (189 != 255)
    v5 = v5
    v8 = 3.14159
    fmt.Println(v8, v8, v8)
    v76 := GenS4{F0: (v2 >> ((266) & 63))}
    switch (^(v1)) {
    case 995:
        if (42 < -186) {
            v77 := []int{128, (v1 * 128)}
            switch "a" {
            case "\n\t":
                v8 = 0.1
            case "foo bar":
                v5 = (1.5 < 0.0)
            case "a":
                v77 = append(v77, v1)
            default:
                var v78 int32 = 1
                _ = v78
            }
            v1 -= v1
            _ = v77
        }
    case 42:
        v2 = 275
    case 9223372036854775807:
        if (65535 >= 255) {
            if v5 {
                v79 := v3[0]
                fmt.Println(v1, v1, v8)
                _ = v79
            } else {
                v80 := GenS9{F0: 158, F1: -792}
                var v81 uint64 = 4294967295
                fmt.Println(v4, v2, v1)
                _ = v80
                _ = v81
            }
            fmt.Println(v4, v5, v4)
        }
    }
    v82 := []int{-9223372036854775808, v1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v82:", len(v82), v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v76
    _ = v8
    _ = v82
}
