package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 int64; F1 int64 }
type GenS6 struct { F0 uint; F1 uint }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint32; F1 uint16 }
type GenS9 struct { F0 uint64; F1 int; F2 uint8 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 uint8; F1 int }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 int64; F1 uint16; F2 bool }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 uint32; F1 string }
type GenS16 struct { F0 uint16 }
type GenS17 struct { F0 uint16 }
type GenS18 struct { F0 int64; F1 uint32; F2 int16 }

func main() {
    var v1 uint8 = 2
    var v2 uint8 = 2
    v3 := make(map[string]uint32)
    v1 = 128
    v1 = (v2 * (v1 - 128))
    fmt.Println(v1, v2)
    v1 -= (v2 ^ v2)
    for fi4 := 0; fi4 < 20; fi4++ {
        v5 := GenS1{F0: 435, F1: (^(v2))}
        for fi6 := 0; fi6 < 4; fi6++ {
            switch "" {
            case "a":
                if true {
                    switch (((("" + ("\n\t" + (("\n\t" + "foo bar") + ("hello" + ("hello" + "hello"))))) + ("hello" + "foo bar")) + ((("hello" + "hello") + ("a" + "foo bar")) + "\n\t")) + ("a" + ("hello" + "foo bar"))) {
                    case "foo bar":
                        v7 := GenS2{F0: 876, F1: 1}
                        _ = v7
                    case "":
                        v1 -= 2
                    case "\n\t":
                        switch "hello" {
                        case "hello":
                            v8 := GenS1{F0: 42, F1: 177}
                            _ = v8
                        case "":
                            v1 -= v2
                        default:
                            v9 := []int{fi4}
                            _ = v9
                        }
                    }
                    var v10 uint32 = 120
                    fmt.Println(fi6, v10)
                    _ = v10
                }
            default:
                var v11 string = ("\n\t" + "hello")
                _ = v11
            }
            switch 142 {
            case 681:
                var v12 float64 = 0.0
                _ = v12
            case 127:
                if (("a" + "a") <= "hello") {
                    var v13 float64 = 3.14159
                    fmt.Println(v1, fi6, v13)
                    _ = v13
                } else {
                    switch (fi6 << ((2) & 63)) {
                    case 97:
                        v14 := GenS2{F0: 42, F1: -1}
                        _ = v14
                    case 32768:
                        if (9223372036854775807 == 0) {
                            v15 := []int{9223372036854775807, fi4, fi4}
                            var v16 uint32 = 4294967295
                            v17 := GenS3{F0: 0}
                            _ = v15
                            _ = v16
                            _ = v17
                        } else {
                            if false {
                                var v18 int32 = 2147483647
                                if true {
                                    if (1e10 > (1e10 / 1.0)) {
                                        v3["foo bar"] = 144
                                        if true {
                                            v19 := GenS4{F0: 65535, F1: true}
                                            _ = v19
                                        } else {
                                            if false {
                                                v20 := []int{1, int(v18), 128, (fi4 % ((42) | 1))}
                                                v1 += (v2 + v1)
                                                v20 = append(v20, (fi6 >> ((828) & 63)))
                                                _ = len(v3)
                                                _ = v20
                                            } else {
                                                v18 = int32(fi6)
                                                v21 := GenS2{F0: 18446744073709551615, F1: -969}
                                                v22 := v3[(("hello" + "hello") + "a")]
                                                _ = v21
                                                _ = v22
                                            }
                                            if true {
                                                var v23 int16 = 0
                                                var v24 int = (fi4 + fi4)
                                                if (32767 <= 120) {
                                                    var v25 uint64 = 377
                                                    v23 = (v23 << ((0) & 63))
                                                    var v26 int32 = (v18 << ((266) & 63))
                                                    var v27 uint32 = uint32(v2)
                                                    _ = v25
                                                    _ = v26
                                                    _ = v27
                                                }
                                                _ = v23
                                                _ = v24
                                            }
                                        }
                                    } else {
                                        v1 += v1
                                        var v28 int8 = 1
                                        _ = v28
                                    }
                                    v29 := GenS4{F0: 1, F1: (0.0 == (1e10 - float32(0.1)))}
                                    var v30 int16 = 0
                                    _ = v29
                                    _ = v30
                                } else {
                                    v31 := GenS5{F0: 0, F1: 9223372036854775807}
                                    fmt.Println(v18, v2)
                                    _ = v31
                                }
                                v2 = 195
                                if ((1e10 > (-2.25 * 0.0)) || (true && true)) {
                                    v3[""] = 42
                                    v18 = v18
                                } else {
                                    var v32 int8 = -72
                                    v18 = 2147483647
                                    v32 = -1
                                    switch fi4 {
                                    case 127:
                                        if (0 >= 42) {
                                            v1 = v1
                                        } else {
                                            v18 -= v18
                                        }
                                    case 2147483647:
                                        fmt.Println(v1, fi4)
                                    case 376:
                                        fmt.Println(fi4)
                                    }
                                    v33 := GenS2{F0: 4294967295, F1: -1}
                                    v34 := GenS5{F0: 832, F1: -381}
                                    v18 -= (v18 - 235)
                                    _ = v32
                                    _ = v33
                                    _ = v34
                                }
                                v18 += v18
                                _ = v18
                            } else {
                                var v35 int64 = -1
                                v36 := GenS1{F0: 32767, F1: v1}
                                switch "hello" {
                                case "":
                                    delete(v3, "\n\t")
                                case "hello":
                                    if true {
                                        fmt.Println(fi4, v35)
                                        var v37 int64 = v35
                                        var v38 uint64 = 18446744073709551615
                                        _ = v37
                                        _ = v38
                                    } else {
                                        if (1 < 32767) {
                                            v1 += (v1 * v2)
                                            v39 := GenS6{F0: 2, F1: 1}
                                            var v40 int32 = 42
                                            _ = v39
                                            _ = v40
                                        } else {
                                            v41 := GenS2{F0: 0, F1: -1}
                                            _ = v41
                                        }
                                        _ = len(v3)
                                        var v42 int64 = 42
                                        fmt.Println(fi6, v35)
                                        _ = v42
                                    }
                                default:
                                    var v43 uint = 2
                                    _ = v43
                                }
                                var v44 string = "\n\t"
                                switch 32767 {
                                case 1:
                                    fmt.Println(fi4, v35, fi6)
                                case -1:
                                    v45 := []int{int(v35)}
                                    _ = v45
                                }
                                _ = v35
                                _ = v36
                                _ = v44
                            }
                            var v46 uint = 134
                            v2 += 0
                            _ = v46
                        }
                    }
                    v47 := []int{}
                    v48 := GenS7{F0: 18446744073709551615}
                    if false {
                        switch fi4 {
                        case 128:
                            switch "" {
                            case "\n\t":
                                v49 := GenS7{F0: 65535}
                                _ = v49
                            }
                        }
                        var v50 int = (fi6 ^ (fi6 << ((65535) & 63)))
                        switch "a" {
                        case "\n\t":
                            if len(v47) > 0 {
                                v47[0] = -333
                            }
                        }
                        _ = v50
                    } else {
                        fmt.Println(v1)
                        v51 := GenS8{F0: 42, F1: 1}
                        switch (^(fi4)) {
                        case 0:
                            v1 = v2
                        case 1:
                            _ = len(v3)
                        case -2147483648:
                            var v52 float64 = float64((3.14159 + 0.1))
                            _ = v52
                        }
                        var v53 int16 = 1
                        _ = v51
                        _ = v53
                    }
                    _ = v47
                    _ = v48
                }
            case -440:
                fmt.Println(fi6, fi4)
            default:
                switch "\n\t" {
                case "a":
                    var v54 int16 = 42
                    _ = v54
                default:
                    var v55 float32 = 0.1
                    if true {
                        switch (-(fi6)) {
                        case 32768:
                            delete(v3, "a")
                        case 42:
                            v56 := GenS7{F0: 255}
                            _ = v56
                        default:
                            v55 -= ((((-2.25 + (v55 - v55)) - float32((3.14159 / (0.0 + 1e10)))) / v55) + float32((((1.5 + 3.14159) - float64(v55)) + (-2.25 + (((((0.1 * float64(v55)) * 1.5) + (float64(v55) - 0.0)) / ((0.0 - float64(((((v55 + -2.25) / v55) * v55) * v55))) - float64((float32(((1.5 / float64(float32((1.5 - float64((1.5 + (((v55 * v55) - 0.1) + (1.5 / v55)))))))) / 1e10)) - v55)))) + 3.14159)))))
                        }
                    } else {
                        var v57 uint64 = 42
                        var v58 int32 = -403
                        _ = v57
                        _ = v58
                    }
                    if true {
                        v59 := []int{fi6, fi6, fi4, (-(fi6))}
                        v55 -= v55
                        if (82 != 992) {
                            fmt.Println(v55)
                            if true {
                                v60 := GenS3{F0: -1}
                                v61 := v3[((("a" + ("" + "hello")) + "hello") + "foo bar")]
                                _ = v60
                                _ = v61
                            } else {
                                switch (("foo bar" + ("hello" + "a")) + ("foo bar" + "foo bar")) {
                                case "hello":
                                    var v62 int32 = 618
                                    _ = v62
                                case "\n\t":
                                    v55 += 3.14159
                                case "a":
                                    switch ("foo bar" + "a") {
                                    case "hello":
                                        if true {
                                            if ((2147483647 < 65535) || true) {
                                                if len(v59) > 0 {
                                                    v59[0] = (^(fi4))
                                                }
                                            } else {
                                                var v63 float64 = (-2.25 - 0.1)
                                                _ = v63
                                            }
                                            if (2 < 866) {
                                                v3["hello"] = 935
                                            }
                                            switch "" {
                                            case "a":
                                                _ = len(v3)
                                            default:
                                                if (!(("foo bar" > "a") && (true || (42 <= 2)))) {
                                                    v2 = 127
                                                    if true {
                                                        switch (fi6 >> ((2) & 63)) {
                                                        case -2147483648:
                                                            fmt.Println(v55, v1, v1)
                                                        case 32768:
                                                            if (0 < 505) {
                                                                var v64 int8 = 127
                                                                var v65 uint64 = 1
                                                                var v66 uint64 = 4294967295
                                                                _ = v64
                                                                _ = v65
                                                                _ = v66
                                                            } else {
                                                                if false {
                                                                    _ = len(v3)
                                                                    v67 := GenS8{F0: 42, F1: 32767}
                                                                    if len(v59) >= 2 {
                                                                        v59 = v59[1:len(v59)]
                                                                    }
                                                                    _ = v67
                                                                }
                                                                var v68 int32 = -259
                                                                _ = v68
                                                            }
                                                        case 4294967295:
                                                            v1 = 42
                                                        }
                                                        v55 -= 3.14159
                                                    }
                                                    fmt.Println(v2, v2)
                                                } else {
                                                    v1 += v2
                                                    v69 := GenS8{F0: 42, F1: 2}
                                                    v70 := GenS9{F0: 819, F1: -9223372036854775808, F2: v2}
                                                    v2 = v2
                                                    _ = v69
                                                    _ = v70
                                                }
                                                v3["hello"] = 65535
                                            }
                                        } else {
                                            var v71 int8 = 65
                                            if (!false) {
                                                var v72 uint16 = 0
                                                var v73 string = "\n\t"
                                                v71 = v71
                                                _ = v72
                                                _ = v73
                                            } else {
                                                _ = len(v3)
                                                var v74 int = fi6
                                                _ = v74
                                            }
                                            _ = v71
                                        }
                                    default:
                                        v55 = (((3.14159 / v55) - float32((0.0 + 3.14159))) - (float32((3.14159 / (1e10 + ((0.0 - 3.14159) + 0.1)))) * ((v55 / 0.1) / v55)))
                                    }
                                }
                                if (827 == 255) {
                                    switch "" {
                                    case "\n\t":
                                        var v75 uint64 = 42
                                        _ = v75
                                    default:
                                        v2 += v2
                                    }
                                    var v76 int = 9223372036854775807
                                    _ = v76
                                }
                                delete(v3, "a")
                                fmt.Println(v1, fi6)
                                if true {
                                    var v77 int = fi4
                                    if len(v59) >= 2 {
                                        v59 = v59[1:len(v59)]
                                    }
                                    v2 = 128
                                    _ = v77
                                }
                                switch ("foo bar" + "\n\t") {
                                case "hello":
                                    fmt.Println(v2, v55, v55)
                                case "\n\t":
                                    var v78 float32 = v55
                                    _ = v78
                                case "foo bar":
                                    v79 := GenS10{F0: 2147483647}
                                    if (false && (fi6 != (fi6 % ((fi4) | 1)))) {
                                        v80 := GenS8{F0: 65535, F1: 898}
                                        _ = v80
                                    }
                                    if true {
                                        v81 := GenS11{F0: v2, F1: 694}
                                        v82 := GenS12{F0: (v1 / ((v1) | 1))}
                                        var v83 uint = 0
                                        _ = v81
                                        _ = v82
                                        _ = v83
                                    } else {
                                        v84 := GenS13{F0: 42, F1: 1, F2: true}
                                        v85 := GenS14{F0: 1}
                                        _ = v84
                                        _ = v85
                                    }
                                    _ = v79
                                default:
                                    switch fi4 {
                                    case 1:
                                        v2 = v2
                                    case 32768:
                                        v86 := GenS15{F0: 42, F1: (("a" + "hello") + "\n\t")}
                                        _ = v86
                                    default:
                                        delete(v3, "hello")
                                    }
                                }
                            }
                            v55 -= (3.14159 / (v55 * ((0.1 * (-2.25 * (float32(0.1) / 1.0))) + (-2.25 + 0.0))))
                        }
                        _ = v59
                    } else {
                        fmt.Println(v2, fi6)
                        v2 = v2
                        if false {
                            var v87 uint64 = 1
                            v88 := GenS16{F0: 2}
                            fmt.Println(fi6)
                            _ = v87
                            _ = v88
                        } else {
                            var v89 string = "a"
                            v89 = (("" + v89) + "hello")
                            var v90 int64 = 1
                            _ = v89
                            _ = v90
                        }
                    }
                    _ = v55
                }
            }
            v91 := []int{(fi6 % ((-9223372036854775808) | 1)), 32767, fi6, fi4}
            _ = fi6
            _ = v91
        }
        v1 += (v2 * v1)
        _ = fi4
        _ = v5
    }
    v92 := GenS17{F0: 1}
    v2 += v1
    v2 = 163
    fmt.Println(v2)
    for fi93 := 0; fi93 < 8; fi93++ {
        var v94 bool = false
        fmt.Println(v2)
        _ = fi93
        _ = v94
    }
    var v95 uint64 = 86
    var v96 int64 = 42
    v2 = v2
    v2 -= (v2 & v2)
    var v97 int = 9223372036854775807
    for fi98 := 0; fi98 < 2; fi98++ {
        for fi99 := 0; fi99 < 11; fi99++ {
            v100 := []int{fi98, (fi99 + v97)}
            v101 := GenS18{F0: (^(v96)), F1: 0, F2: 1}
            var v102 int8 = 0
            _ = fi99
            _ = v100
            _ = v101
            _ = v102
        }
        var v103 uint16 = 1
        fmt.Println(v97)
        _ = fi98
        _ = v103
    }
    v104 := v3["\n\t"]
    var v105 int16 = -43
    fmt.Println("v1:", v1)
    fmt.Println("v104:", v104)
    fmt.Println("v105:", v105)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v95:", v95)
    fmt.Println("v96:", v96)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v104
    _ = v105
    _ = v2
    _ = v3
    _ = v92
    _ = v95
    _ = v96
    _ = v97
}
