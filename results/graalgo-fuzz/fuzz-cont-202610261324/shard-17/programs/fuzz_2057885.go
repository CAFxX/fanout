package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint; F1 uint16 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint8; F1 int16 }
type GenS10 struct { F0 int32; F1 bool; F2 int }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 uint; F1 int16 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = ("foo bar" + "a")
    var v3 float64 = (0.0 - 3.14159)
    var v4 uint64 = 42
    var v5 int32 = int32(v4)
    var v6 int32 = v5
    var v7 int16 = -288
    fmt.Println(v7)
    for fi8 := 0; fi8 < 6; fi8++ {
        v6 = 1
        var v9 int64 = 9223372036854775807
        _ = fi8
        _ = v9
    }
    var v10 int64 = -980
    if true {
        for fi11 := 0; fi11 < 19; fi11++ {
            for fi12 := 0; fi12 < 10; fi12++ {
                v13 := []int{1, fi11}
                var v14 int32 = v5
                var v15 int32 = -2147483648
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
            }
            if (248 != 398) {
                v3 = (v3 * (v3 - ((v3 - 0.1) * 1e10)))
                if (true || (1 > v10)) {
                    fmt.Println(v7, v3, v4)
                    v2 = "hello"
                }
            } else {
                v16 := GenS1{F0: "foo bar", F1: ((true && false) || (0 > 100))}
                for fi17 := 0; fi17 < 2; fi17++ {
                    v18 := []int{fi11, fi11, fi17}
                    switch v2 {
                    case "hello":
                        switch "\n\t" {
                        case "foo bar":
                            v19 := GenS1{F0: "a", F1: (-1 < -9223372036854775808)}
                            _ = v19
                        }
                    case "":
                        switch (-(fi11)) {
                        case 128:
                            if false {
                                var v20 int16 = v7
                                var v21 float32 = float32((v3 / v3))
                                if false {
                                    v4 -= v4
                                    v22 := GenS1{F0: (("a" + v2) + (v2 + ((((("hello" + "hello") + ("a" + v2)) + ((v2 + v2) + "hello")) + ("\n\t" + "foo bar")) + v2))), F1: false}
                                    switch (v2 + (v2 + ("foo bar" + ""))) {
                                    case "a":
                                        if true {
                                            var v23 string = v2
                                            var v24 float32 = float32(((1e10 + v3) + v3))
                                            var v25 int32 = v6
                                            _ = v23
                                            _ = v24
                                            _ = v25
                                        } else {
                                            v21 = 1.5
                                            v20 = (v7 & (v7 - v20))
                                            fmt.Println(v21, v3, v7)
                                        }
                                        var v26 uint = 42
                                        v27 := GenS2{F0: (^(v4)), F1: -665}
                                        _ = v26
                                        _ = v27
                                    case "foo bar":
                                        v10 += (v10 >> ((65535) & 63))
                                    }
                                    _ = v22
                                } else {
                                    if false {
                                        v28 := GenS1{F0: v2, F1: false}
                                        _ = len(v18)
                                        _ = v28
                                    } else {
                                        var v29 uint32 = 48
                                        v4 = 4294967295
                                        switch 4294967295 {
                                        case 0:
                                            if true {
                                                fmt.Println(fi17)
                                                switch 128 {
                                                case 1:
                                                    var v30 uint64 = 472
                                                    _ = v30
                                                case -1:
                                                    v10 += (v10 & (^(v10)))
                                                    switch (fi11 & 2147483647) {
                                                    case 32768:
                                                        if ((1000 == (v7 << ((65535) & 63))) && (21 >= 42)) {
                                                            var v31 int8 = -1
                                                            if true {
                                                                fmt.Println(v7, v29, v6)
                                                                var v32 uint16 = 65535
                                                                _ = v32
                                                            } else {
                                                                var v33 uint8 = 1
                                                                v4 -= v4
                                                                _ = v33
                                                            }
                                                            _ = v31
                                                        }
                                                    case 32767:
                                                        v34 := GenS3{F0: v29, F1: 255}
                                                        _ = v34
                                                    case -1:
                                                        var v35 int16 = (^(v20))
                                                        _ = v35
                                                    default:
                                                        v18 = append(v18, (fi11 + 42))
                                                    }
                                                    var v36 bool = (!true)
                                                    _ = v36
                                                default:
                                                    v20 -= -307
                                                }
                                                if ((v4 - v4) >= (v4 % ((v4) | 1))) {
                                                    switch v2 {
                                                    case "hello":
                                                        var v37 uint = 18446744073709551615
                                                        _ = v37
                                                    case "foo bar":
                                                        var v38 int = 42
                                                        _ = v38
                                                    case "":
                                                        if false {
                                                            v39 := GenS2{F0: 0, F1: 1}
                                                            var v40 string = ((((v2 + "") + v2) + "hello") + ("" + (((v2 + v2) + v2) + (v2 + v2))))
                                                            v41 := GenS3{F0: v29, F1: 32767}
                                                            _ = v39
                                                            _ = v40
                                                            _ = v41
                                                        } else {
                                                            var v42 uint32 = (v29 - v29)
                                                            _ = v42
                                                        }
                                                    }
                                                    var v43 bool = (713 >= 256)
                                                    _ = v43
                                                }
                                                var v44 uint16 = 407
                                                fmt.Println(v20, v6, v2)
                                                var v45 string = "foo bar"
                                                _ = v44
                                                _ = v45
                                            }
                                        case 2147483647:
                                            fmt.Println(v5, v5)
                                        case -9223372036854775808:
                                            switch fi17 {
                                            case 703:
                                                v46 := GenS2{F0: (v4 % ((2) | 1)), F1: (v10 & 42)}
                                                _ = v46
                                            case 564:
                                                switch fi17 {
                                                case 32767:
                                                    var v47 int16 = v7
                                                    _ = v47
                                                case 20:
                                                    if true {
                                                        var v48 int = (fi17 % ((fi17) | 1))
                                                        _ = v48
                                                    }
                                                case 2147483647:
                                                    var v49 uint8 = 145
                                                    _ = v49
                                                }
                                            }
                                            v10 += v10
                                        default:
                                            v10 = -1
                                        }
                                        _ = v29
                                    }
                                }
                                var v50 int16 = v20
                                v51 := GenS4{F0: 1}
                                var v52 bool = (((!(true || (v4 == v4))) && (false && true)) || true)
                                fmt.Println(fi11)
                                var v53 int = fi11
                                v7 = (^(v20))
                                v54 := v1[42]
                                _ = v20
                                _ = v21
                                _ = v50
                                _ = v51
                                _ = v52
                                _ = v53
                                _ = v54
                            } else {
                                switch (-(fi17)) {
                                case 32767:
                                    switch fi17 {
                                    case 0:
                                        var v55 uint = 406
                                        _ = v55
                                    default:
                                        v56 := GenS3{F0: 4294967295, F1: 32767}
                                        _ = v56
                                    }
                                case -2147483648:
                                    v57 := GenS2{F0: v4, F1: v10}
                                    _ = v57
                                case 42:
                                    if len(v18) > 0 {
                                        v18[0] = fi11
                                    }
                                    if ((v6 | (v6 | (v6 + v6))) >= v5) {
                                        if true {
                                            var v58 uint64 = (v4 % (((v4 / ((1) | 1))) | 1))
                                            _ = v58
                                        } else {
                                            if len(v18) >= 2 {
                                                v18 = v18[1:len(v18)]
                                            }
                                            if false {
                                                var v59 int64 = (v10 | v10)
                                                if false {
                                                    var v60 uint8 = 50
                                                    if false {
                                                        if (v3 > (1.5 * v3)) {
                                                            var v61 string = v2
                                                            v2 = v61
                                                            _ = v61
                                                        } else {
                                                            v62 := GenS5{F0: (^(v7))}
                                                            if true {
                                                                var v63 int64 = -929
                                                                _ = v63
                                                            }
                                                            _ = len(v1)
                                                            _ = v62
                                                        }
                                                        var v64 uint = 42
                                                        _ = len(v1)
                                                        if true {
                                                            v65 := GenS6{F0: (v64 << (((v64 << ((32) & 63))) & 63)), F1: 1}
                                                            v18 = append(v18, 9223372036854775807)
                                                            v64 += 4294967295
                                                            _ = v65
                                                        }
                                                        _ = v64
                                                    }
                                                    _ = v60
                                                }
                                                v66 := GenS7{F0: (-(v59))}
                                                _ = v59
                                                _ = v66
                                            } else {
                                                v10 = (v10 | v10)
                                                v5 -= (v6 << ((42) & 63))
                                            }
                                            v7 -= v7
                                        }
                                        if (((!(65535 == 65535)) || false) && (65535 > 42)) {
                                            v3 += v3
                                            var v67 int64 = v10
                                            _ = v67
                                        } else {
                                            v68 := GenS8{F0: -2147483648}
                                            if true {
                                                v4 -= (v4 - 2)
                                                var v69 bool = (true || ((1e10 == float32((v3 * 0.1))) || (((v2 + v2) + v2) <= ((v2 + "\n\t") + ("" + ("foo bar" + (((v2 + "\n\t") + ("" + "hello")) + v2)))))))
                                                var v70 int64 = (^(v10))
                                                _ = v69
                                                _ = v70
                                            }
                                            _ = v68
                                        }
                                        fmt.Println(v5, v10, v3)
                                    }
                                    v18 = append(v18, fi11)
                                    var v71 uint16 = 255
                                    if (!true) {
                                        if (((v7 ^ v7) <= (v7 >> ((860) & 63))) && true) {
                                            if false {
                                                if (307 <= 1) {
                                                    switch "a" {
                                                    case "hello":
                                                        v72 := GenS4{F0: (v71 / ((v71) | 1))}
                                                        _ = v72
                                                    case "\n\t":
                                                        fmt.Println(v4, fi17, v10)
                                                    }
                                                }
                                                var v73 string = "a"
                                                _ = v73
                                            }
                                            fmt.Println(v6, v71)
                                        }
                                    } else {
                                        v7 = (v7 & (v7 << ((4294967295) & 63)))
                                    }
                                    _ = v71
                                default:
                                    switch v2 {
                                    case "foo bar":
                                        v5 -= 2147483647
                                    case "a":
                                        switch (fi11 - (fi11 - (fi17 << ((42) & 63)))) {
                                        case 42:
                                            _ = len(v18)
                                        default:
                                            if (((fi11 & (fi17 >> ((65535) & 63))) <= (fi17 - fi17)) && ((128 < 42) || (!(true || true)))) {
                                                if (v6 != 1) {
                                                    _ = len(v18)
                                                    fmt.Println(v2, fi17, v6)
                                                    if len(v18) > 0 {
                                                        v18[0] = (fi17 - (fi17 % ((fi11) | 1)))
                                                    }
                                                } else {
                                                    v74 := GenS4{F0: 969}
                                                    var v75 int = (fi17 - fi11)
                                                    if false {
                                                        v7 += (v7 | (v7 * -92))
                                                        v4 -= v4
                                                    } else {
                                                        v76 := GenS4{F0: 1}
                                                        _ = v76
                                                    }
                                                    _ = v74
                                                    _ = v75
                                                }
                                                if (v2 == (v2 + v2)) {
                                                    switch fi11 {
                                                    case -890:
                                                        delete(v1, 42)
                                                    case 766:
                                                        switch fi11 {
                                                        case -2147483648:
                                                            var v77 int32 = -203
                                                            _ = v77
                                                        case 1:
                                                            var v78 bool = (true || true)
                                                            _ = v78
                                                        case -154:
                                                            var v79 string = v2
                                                            _ = v79
                                                        }
                                                    default:
                                                        v7 -= (v7 - v7)
                                                    }
                                                    v18 = append(v18, fi17)
                                                    if (v10 >= v10) {
                                                        switch (v2 + v2) {
                                                        case "a":
                                                            if false {
                                                                v1[2] = 964
                                                            }
                                                        case "foo bar":
                                                            v80 := GenS1{F0: v2, F1: false}
                                                            _ = v80
                                                        case "":
                                                            fmt.Println(v7, v6, v7)
                                                            var v81 int = fi17
                                                            _ = v81
                                                        }
                                                        var v82 float64 = v3
                                                        _ = v82
                                                    }
                                                } else {
                                                    fmt.Println(fi17, v6)
                                                }
                                            }
                                        }
                                    case "hello":
                                        var v83 string = (v2 + "a")
                                        _ = v83
                                    }
                                }
                                v4 += v4
                            }
                        }
                    default:
                        switch v2 {
                        case "":
                            var v84 string = v2
                            _ = v84
                        case "a":
                            switch (fi17 << ((42) & 63)) {
                            case 0:
                                if (uint(v4) >= 4294967295) {
                                    v5 = -2147483648
                                    var v85 bool = (2 < 65535)
                                    var v86 int32 = v6
                                    v87 := GenS9{F0: 127, F1: v7}
                                    _ = v85
                                    _ = v86
                                    _ = v87
                                } else {
                                    if true {
                                        v4 = v4
                                    }
                                    if ((v4 % ((v4) | 1)) <= (v4 >> ((4294967295) & 63))) {
                                        v7 += (v7 ^ 1000)
                                    } else {
                                        if ((-2.25 + 3.14159) == float32(float64((((1e10 / ((0.0 * (1e10 / 1e10)) + 1.5)) + -2.25) - float32((v3 + v3)))))) {
                                            switch int(v4) {
                                            case -906:
                                                v88 := GenS10{F0: v5, F1: false, F2: (fi17 ^ (fi17 ^ (fi11 - (fi11 % ((fi11) | 1)))))}
                                                _ = v88
                                            default:
                                                v89 := GenS11{F0: 42}
                                                _ = v89
                                            }
                                        } else {
                                            delete(v1, 2147483648)
                                        }
                                    }
                                    var v90 uint64 = v4
                                    v5 = v5
                                    if len(v18) >= 2 {
                                        v18 = v18[1:len(v18)]
                                    }
                                    v6 -= (v6 / (((v6 >> ((262) & 63))) | 1))
                                    fmt.Println(fi17, v90, fi17)
                                    _ = v90
                                }
                            default:
                                switch 2147483647 {
                                case 495:
                                    if ((v10 - (v10 | v10)) == (v10 * v10)) {
                                        v7 -= (v7 / (((v7 * -1)) | 1))
                                        var v91 float32 = 3.14159
                                        var v92 uint16 = 209
                                        _ = v91
                                        _ = v92
                                    }
                                case -683:
                                    var v93 int = fi11
                                    _ = v93
                                }
                            }
                            fmt.Println(v2, v7, fi11)
                            var v94 int64 = -1
                            _ = v94
                        case "\n\t":
                            v95 := GenS12{F0: 441, F1: v7}
                            _ = v95
                        default:
                            switch fi11 {
                            case -9223372036854775808:
                                v10 = 9223372036854775807
                            case 4294967295:
                                var v96 uint16 = 32767
                                _ = v96
                            default:
                                if len(v18) > 0 {
                                    v18[0] = (fi17 % ((fi11) | 1))
                                }
                            }
                        }
                    }
                    switch v2 {
                    case "\n\t":
                        v7 -= (v7 - v7)
                    }
                    if true {
                        var v97 uint16 = 907
                        if true {
                            fmt.Println(v5)
                            var v98 uint = 255
                            _ = v98
                        }
                        _ = v97
                    } else {
                        var v99 int32 = v5
                        _ = v99
                    }
                    _ = fi17
                    _ = v18
                }
                var v100 uint16 = 1
                _ = v100
                _ = v16
            }
            v6 += (-(v5))
            _ = fi11
        }
    } else {
        if (!(!((v4 + 42) > 306))) {
            if (-410 > -436) {
                v5 += v6
            } else {
                v10 = (v10 * v10)
                var v101 int8 = 1
                if true {
                    v102 := GenS8{F0: -430}
                    _ = v102
                } else {
                    var v103 string = v2
                    _ = v103
                }
                _ = v101
            }
            v104 := GenS8{F0: v5}
            var v105 string = ""
            _ = v104
            _ = v105
        }
        v4 = v4
    }
    v3 -= 3.14159
    var v106 int8 = 100
    v107 := GenS12{F0: 4294967295, F1: v7}
    v4 = (v4 - v4)
    if true {
        v108 := GenS1{F0: v2, F1: true}
        fmt.Println(v4, v3)
        _ = v108
    } else {
        v4 -= 4294967295
        var v109 int = 9223372036854775807
        _ = v109
    }
    var v110 uint64 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v106:", v106)
    fmt.Println("v110:", v110)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v106
    _ = v107
    _ = v110
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
