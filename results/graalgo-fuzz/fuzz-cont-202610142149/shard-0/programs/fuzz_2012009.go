package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int8 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 uint64; F1 int64; F2 int16 }
type GenS4 struct { F0 uint32; F1 uint16; F2 int32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 bool; F1 uint64 }
type GenS7 struct { F0 uint32; F1 uint; F2 uint16 }
type GenS8 struct { F0 int8; F1 int8 }
type GenS9 struct { F0 int8 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = ((("\n\t" + ("hello" + ("\n\t" + ("foo bar" + "foo bar")))) + ("" + (("" + (((("a" + "") + ("hello" + (("a" + "a") + "foo bar"))) + "") + "hello")) + "\n\t"))) + ("\n\t" + ("" + "hello")))
    v3 := make(map[int]string)
    v4 := make(map[string]uint32)
    v5 := make(map[int]int)
    var v6 int16 = 32767
    v2 = (v2 + ("a" + v2))
    var v7 uint8 = 231
    if (!((true || true) || (128 <= 1))) {
        if true {
            v7 = (v7 + (v7 ^ 127))
        } else {
            var v8 uint64 = 1
            v7 = v7
            _ = v8
        }
    }
    v9 := GenS1{F0: 1, F1: 1, F2: -1}
    fmt.Println(v2, v6, v6)
    var v10 float64 = 1e10
    v10 -= v10
    var v11 float32 = 0.1
    var v12 bool = false
    v13 := GenS1{F0: 100, F1: (v7 ^ v7), F2: 42}
    switch -539 {
    case -1:
        _ = len(v5)
    case 59:
        if (1.5 <= v10) {
            fmt.Println(v12)
        } else {
            if v12 {
                v12 = v12
                for fi14 := 0; fi14 < 15; fi14++ {
                    v6 = 42
                    for fi15 := 0; fi15 < 7; fi15++ {
                        var v16 uint8 = 42
                        if false {
                            if (762 > -1) {
                                v2 = v2
                                var v17 uint16 = 821
                                _ = v17
                            } else {
                                v18 := GenS1{F0: 100, F1: (v16 ^ v16), F2: 42}
                                var v19 float32 = ((float32((0.0 + v10)) + ((float32(((((v10 / v10) / (v10 * v10)) * v10) - v10)) * 0.0) + float32(float64(v11)))) * float32(v10))
                                _ = v18
                                _ = v19
                            }
                        } else {
                            fmt.Println(v6)
                            v20 := GenS1{F0: -1, F1: v7, F2: -74}
                            fmt.Println(v10)
                            _ = v20
                        }
                        _ = fi15
                        _ = v16
                    }
                    _ = fi14
                }
            } else {
                switch v2 {
                case "a":
                    v10 += (v10 + v10)
                case "":
                    for fi21 := 0; fi21 < 19; fi21++ {
                        v22 := []int{32768, fi21}
                        var v23 uint8 = v7
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                case "\n\t":
                    if (v12 && (!v12)) {
                        v24 := []int{}
                        _ = v24
                    }
                }
            }
            for fi25 := 0; fi25 < 13; fi25++ {
                v6 = v6
                v26 := GenS2{F0: fi25, F1: 238}
                fmt.Println(v7, v6)
                _ = fi25
                _ = v26
            }
        }
    case 724:
        v27 := GenS2{F0: -9223372036854775808, F1: (v7 << ((1) & 63))}
        _ = v27
    default:
        for fi28 := 0; fi28 < 18; fi28++ {
            for fi29 := 0; fi29 < 5; fi29++ {
                var v30 string = (v2 + v2)
                switch fi29 {
                case 21:
                    if (1 == v7) {
                        switch v2 {
                        case "hello":
                            var v31 uint32 = 338
                            if true {
                                if v12 {
                                    var v32 bool = ((v12 && ((v7 / ((v7) | 1)) == 255)) && (-32768 > v6))
                                    var v33 uint = 2
                                    v10 = (float64(v11) * 0.1)
                                    var v34 int16 = -482
                                    _ = v32
                                    _ = v33
                                    _ = v34
                                }
                            } else {
                                switch int(v31) {
                                case 2147483647:
                                    if v12 {
                                        v35 := GenS2{F0: fi29, F1: v7}
                                        var v36 uint32 = (^(v31))
                                        _ = v35
                                        _ = v36
                                    }
                                case 0:
                                    v37 := v1[v31]
                                    _ = v37
                                default:
                                    switch (fi29 / ((4294967295) | 1)) {
                                    case 42:
                                        v38 := []int{(fi29 << ((42) & 63)), (fi29 / (((fi29 / ((int(v6)) | 1))) | 1)), (fi28 >> ((255) & 63))}
                                        _ = v38
                                    case 32768:
                                        var v39 float32 = (v11 + v11)
                                        _ = v39
                                    default:
                                        var v40 uint16 = 42
                                        _ = v40
                                    }
                                }
                                var v41 int = 2147483647
                                var v42 float32 = 1.5
                                _ = v41
                                _ = v42
                            }
                            switch ("\n\t" + (v2 + v2)) {
                            case "hello":
                                if v12 {
                                    switch (-(fi29)) {
                                    case -949:
                                        v10 = 0.0
                                    case -9223372036854775808:
                                        if (1.5 != 0.1) {
                                            v31 = v31
                                        }
                                    }
                                    v6 -= 244
                                    var v43 float64 = float64(((float32(v10) / ((0.1 / (v11 + v11)) - 0.1)) - v11))
                                    v44 := []int{(fi29 ^ -231)}
                                    v12 = (32767 < v6)
                                    _ = len(v1)
                                    _ = v43
                                    _ = v44
                                }
                            case "a":
                                v11 -= v11
                                v10 = float64(1e10)
                            }
                            _ = v31
                        case "foo bar":
                            v45 := GenS3{F0: 18446744073709551615, F1: 0, F2: v6}
                            _ = v45
                        case "\n\t":
                            var v46 uint64 = 42
                            _ = v46
                        default:
                            fmt.Println(v30)
                        }
                        if v12 {
                            var v47 int16 = v6
                            v48 := []int{fi29, (fi29 & (fi28 + fi28))}
                            v12 = (-49 == int8(v47))
                            _ = v47
                            _ = v48
                        }
                    } else {
                        v3[fi28] = v2
                        v6 = (-(v6))
                    }
                default:
                    if ((-1 <= 42) || (!v12)) {
                        if (65535 < 2147483647) {
                            if (v12 || v12) {
                                if (9223372036854775807 == 42) {
                                    v30 = v2
                                    v49 := []int{fi29, (-(fi29)), (fi28 << ((1) & 63)), fi29}
                                    _ = v49
                                } else {
                                    var v50 int16 = (v6 ^ (v6 - (v6 * v6)))
                                    v51 := []int{-1, (fi28 / ((fi29) | 1))}
                                    if ("foo bar" <= (v2 + (v2 + v30))) {
                                        v10 += v10
                                        if (v12 && ((!((((v10 + v10) / (((3.14159 + 1e10) + 1.5) + v10)) * float64(1.5)) <= (-2.25 / (v10 + 0.0)))) && (2 > 86))) {
                                            if (false && true) {
                                                v52 := GenS4{F0: 0, F1: 255, F2: 597}
                                                v50 += (v50 >> ((2) & 63))
                                                v30 = v30
                                                _ = v52
                                            }
                                            switch v2 {
                                            case "\n\t":
                                                if (922 >= 441) {
                                                    v1[0] = 411
                                                } else {
                                                    var v53 int = 9223372036854775807
                                                    if v12 {
                                                        var v54 string = v30
                                                        if v12 {
                                                            var v55 bool = v12
                                                            if v12 {
                                                                _ = len(v51)
                                                                v53 = 481
                                                                var v56 int32 = 0
                                                                _ = v56
                                                            } else {
                                                                if (!(9223372036854775807 >= -9223372036854775808)) {
                                                                    v1[0] = 326
                                                                }
                                                                v11 += (1e10 - 0.1)
                                                            }
                                                            if v12 {
                                                                fmt.Println(v12)
                                                            } else {
                                                                var v57 int32 = 62
                                                                v55 = v55
                                                                if ((v57 >> ((1) & 63)) < (^(v57))) {
                                                                    var v58 int16 = 32767
                                                                    _ = v58
                                                                } else {
                                                                    var v59 uint = 42
                                                                    if false {
                                                                        var v60 uint8 = (^(v7))
                                                                        v54 = (((v30 + v54) + v54) + v30)
                                                                        _ = v60
                                                                    }
                                                                    _ = v59
                                                                }
                                                                v51 = append(v51, (fi29 % ((v53) | 1)))
                                                                _ = v57
                                                            }
                                                            v12 = v12
                                                            v50 = (v50 + v6)
                                                            v61 := GenS5{F0: 32767}
                                                            _ = v55
                                                            _ = v61
                                                        }
                                                        _ = v54
                                                    } else {
                                                        var v62 float64 = v10
                                                        _ = v62
                                                    }
                                                    _ = v53
                                                }
                                            case "":
                                                v12 = v12
                                            case "a":
                                                switch fi28 {
                                                case 2147483647:
                                                    var v63 uint64 = 1
                                                    _ = v63
                                                case -410:
                                                    var v64 int8 = -93
                                                    _ = v64
                                                }
                                            }
                                            if (!v12) {
                                                v3[(fi29 * fi28)] = "a"
                                            } else {
                                                if v12 {
                                                    switch ((v30 + v30) + v30) {
                                                    case "\n\t":
                                                        var v65 uint = 151
                                                        _ = v65
                                                    }
                                                    switch fi29 {
                                                    case 9223372036854775807:
                                                        v66 := GenS3{F0: 1, F1: 0, F2: v6}
                                                        _ = v66
                                                    case -2147483648:
                                                        if false {
                                                            if len(v51) > 0 {
                                                                v51[0] = fi28
                                                            }
                                                            if (float64(v11) <= (0.0 - (((float64(float32(float64(-2.25))) / v10) + -2.25) + v10))) {
                                                                v30 = "hello"
                                                                if ((42 == 127) || (!((4294967295 <= 65535) || false))) {
                                                                    var v67 uint32 = 1
                                                                    var v68 uint64 = 4294967295
                                                                    if v12 {
                                                                        switch ((v2 + v30) + "\n\t") {
                                                                        case "hello":
                                                                            v50 -= (v50 ^ 42)
                                                                        case "":
                                                                            if v12 {
                                                                                var v69 string = ("a" + v30)
                                                                                _ = v69
                                                                            }
                                                                        default:
                                                                            v70 := GenS5{F0: 0}
                                                                            _ = v70
                                                                        }
                                                                    }
                                                                    _ = v67
                                                                    _ = v68
                                                                }
                                                            } else {
                                                                if ((v6 % ((v6) | 1)) == v50) {
                                                                    if len(v51) >= 2 {
                                                                        v51 = v51[1:len(v51)]
                                                                    }
                                                                }
                                                                var v71 float64 = v10
                                                                v12 = (v12 && v12)
                                                                _ = v71
                                                            }
                                                            if (-891 >= (^(fi28))) {
                                                                v11 = (1e10 / (v11 + v11))
                                                                switch (^(fi28)) {
                                                                case 127:
                                                                    var v72 uint32 = 42
                                                                    fmt.Println(v10, fi28)
                                                                    v50 += v6
                                                                    switch fi29 {
                                                                    case 786:
                                                                        _ = len(v51)
                                                                    }
                                                                    _ = v72
                                                                }
                                                            } else {
                                                                fmt.Println(v12, fi29)
                                                            }
                                                        } else {
                                                            var v73 string = (v2 + ("hello" + v30))
                                                            v7 -= v7
                                                            _ = v73
                                                        }
                                                    }
                                                } else {
                                                    if true {
                                                        _ = len(v51)
                                                        var v74 uint16 = 181
                                                        _ = v74
                                                    }
                                                    var v75 uint64 = 1
                                                    _ = v75
                                                }
                                                if v12 {
                                                    var v76 uint = 0
                                                    _ = v76
                                                } else {
                                                    if (!true) {
                                                        if (!(v50 > -32768)) {
                                                            v77 := GenS6{F0: v12, F1: 84}
                                                            _ = v77
                                                        } else {
                                                            if v12 {
                                                                v78 := GenS7{F0: 1, F1: 37, F2: 1}
                                                                fmt.Println(v30, fi28)
                                                                v51 = append(v51, fi28)
                                                                _ = v78
                                                            }
                                                            var v79 uint32 = 65535
                                                            fmt.Println(v11, v50)
                                                            _ = v79
                                                        }
                                                        switch (fi29 >> ((395) & 63)) {
                                                        case 588:
                                                            v7 = (v7 ^ v7)
                                                        case 619:
                                                            var v80 bool = (!v12)
                                                            _ = v80
                                                        case 4294967295:
                                                            v2 = v2
                                                        default:
                                                            v11 -= 3.14159
                                                        }
                                                        var v81 uint32 = 764
                                                        var v82 uint16 = 255
                                                        var v83 int64 = 42
                                                        _ = v81
                                                        _ = v82
                                                        _ = v83
                                                    }
                                                }
                                            }
                                        } else {
                                            v84 := GenS8{F0: 127, F1: 64}
                                            switch (fi28 | (fi29 + (^(fi29)))) {
                                            case 1:
                                                v7 = (^(v7))
                                            default:
                                                v85 := GenS8{F0: 42, F1: 1}
                                                _ = v85
                                            }
                                            v86 := GenS9{F0: 1}
                                            _ = v84
                                            _ = v86
                                        }
                                    }
                                    _ = v50
                                    _ = v51
                                }
                            }
                        }
                    }
                }
                var v87 uint8 = v7
                _ = fi29
                _ = v30
                _ = v87
            }
            _ = fi28
        }
    }
    var v88 string = v2
    v11 = 0.1
    fmt.Println(v2, v12)
    v89 := GenS4{F0: 2, F1: 2, F2: int32(v7)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v88
    _ = v89
    _ = v9
}
