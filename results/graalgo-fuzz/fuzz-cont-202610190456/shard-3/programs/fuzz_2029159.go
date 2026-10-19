package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint16; F1 uint; F2 uint16 }
type GenS4 struct { F0 uint; F1 uint16; F2 bool }
type GenS5 struct { F0 string; F1 int16; F2 uint16 }
type GenS6 struct { F0 uint32; F1 uint; F2 uint32 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 string; F1 uint }
type GenS9 struct { F0 uint8 }

func main() {
    var v1 int32 = 0
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]string)
    v4 := GenS1{F0: 42, F1: 100}
    var v5 uint64 = 42
    v1 = v1
    v6 := GenS1{F0: -9223372036854775808, F1: -113}
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 int32 = v1
        for fi9 := 0; fi9 < 19; fi9++ {
            var v10 float32 = float32(((1e10 + (float64((((1e10 - float32(0.1)) / 0.1) / 1.0)) + 1e10)) + (0.0 + 0.1)))
            v11 := GenS2{F0: 579, F1: 1}
            fmt.Println(fi9, fi7, v10)
            switch fi7 {
            case 42:
                if ((!true) && true) {
                    if true {
                        switch fi9 {
                        case 962:
                            var v12 uint64 = (v5 - 4294967295)
                            if (!(!false)) {
                                v3[65535] = ("foo bar" + "\n\t")
                                if (98 < 100) {
                                    v5 = 4294967295
                                    var v13 int = (fi9 * fi7)
                                    _ = v13
                                } else {
                                    v12 -= (^(v12))
                                    v10 = (float32(3.14159) / v10)
                                    switch fi7 {
                                    case -2147483648:
                                        v14 := GenS3{F0: 65535, F1: 172, F2: 255}
                                        _ = v14
                                    default:
                                        v15 := GenS1{F0: 9223372036854775807, F1: -1}
                                        _ = v15
                                    }
                                    if false {
                                        var v16 uint16 = 2
                                        _ = v16
                                    }
                                    switch "a" {
                                    case "":
                                        var v17 int64 = -9223372036854775808
                                        _ = v17
                                    default:
                                        v5 = (v5 % (((v5 / (((^(v5))) | 1))) | 1))
                                    }
                                    var v18 uint8 = 127
                                    var v19 int32 = (v8 | v1)
                                    _ = v18
                                    _ = v19
                                }
                            }
                            _ = v12
                        case 32768:
                            fmt.Println(v8, v1, v1)
                        case 1:
                            fmt.Println(v5, fi7, v1)
                        default:
                            var v20 string = "foo bar"
                            _ = v20
                        }
                    }
                    v1 += v8
                } else {
                    switch "a" {
                    case "foo bar":
                        v8 += v8
                    default:
                        v5 = (v5 & 2)
                    }
                    fmt.Println(fi9)
                    if (42 != 1) {
                        if false {
                            v8 += (v1 * -1)
                        }
                        if false {
                            var v21 int64 = 42
                            var v22 uint = 692
                            _ = v21
                            _ = v22
                        } else {
                            v23 := GenS2{F0: 842, F1: 0}
                            if false {
                                v24 := v3[2]
                                var v25 uint32 = 929
                                _ = v24
                                _ = v25
                            } else {
                                var v26 uint8 = 1
                                if (((-(fi7)) <= (fi9 * fi9)) && (!(!true))) {
                                    v27 := []int{fi7, (fi9 ^ fi7)}
                                    v1 = v8
                                    var v28 float64 = 3.14159
                                    switch ("foo bar" + ((("\n\t" + "\n\t") + "hello") + (("foo bar" + (("hello" + ("" + "foo bar")) + "")) + (("foo bar" + "foo bar") + "foo bar")))) {
                                    case "":
                                        if (498 <= 259) {
                                            if false {
                                                if false {
                                                    if ((1e10 > (v28 * float64(-2.25))) || (!(false && true))) {
                                                        var v29 bool = (127 > (fi9 % (((fi9 << ((711) & 63))) | 1)))
                                                        if (!false) {
                                                            if (252 > (v5 * v5)) {
                                                                v26 -= (v26 + (v26 ^ 42))
                                                                if v29 {
                                                                    var v30 uint16 = 101
                                                                    _ = v30
                                                                } else {
                                                                    switch "\n\t" {
                                                                    case "":
                                                                        if (v29 || v29) {
                                                                            var v31 uint8 = (^(v26))
                                                                            _ = v31
                                                                        } else {
                                                                            var v32 uint16 = 112
                                                                            _ = v32
                                                                        }
                                                                    case "\n\t":
                                                                        var v33 uint = 0
                                                                        _ = v33
                                                                    case "hello":
                                                                        v8 += -1
                                                                    }
                                                                    fmt.Println(v8)
                                                                    var v34 float32 = 1.5
                                                                    _ = v34
                                                                }
                                                                if true {
                                                                    v5 += (v5 / ((v5) | 1))
                                                                } else {
                                                                    if v29 {
                                                                        v29 = v29
                                                                        var v35 uint16 = 595
                                                                        var v36 int8 = 0
                                                                        _ = v35
                                                                        _ = v36
                                                                    }
                                                                    if true {
                                                                        if (98 == -128) {
                                                                            v37 := GenS4{F0: 2, F1: 725, F2: v29}
                                                                            fmt.Println(v10, v28, v29)
                                                                            _ = v37
                                                                        }
                                                                        var v38 int16 = 1
                                                                        _ = v38
                                                                    }
                                                                    if len(v27) > 0 {
                                                                        v27[0] = (fi9 % ((fi9) | 1))
                                                                    }
                                                                    fmt.Println(v8, v5, fi9)
                                                                    v26 += (v26 ^ (v26 & (v26 * (v26 << ((65535) & 63)))))
                                                                    v39 := GenS1{F0: 0, F1: -96}
                                                                    _ = v39
                                                                }
                                                                v28 = v28
                                                                v27 = append(v27, (fi9 | 616))
                                                            } else {
                                                                v5 = 528
                                                                v40 := GenS5{F0: "\n\t", F1: 1000, F2: 42}
                                                                if true {
                                                                    delete(v3, 65535)
                                                                }
                                                                _ = v40
                                                            }
                                                        } else {
                                                            switch (-(fi7)) {
                                                            case 887:
                                                                fmt.Println(fi7, v5, fi9)
                                                            case 32768:
                                                                v8 = (v8 >> ((65535) & 63))
                                                            }
                                                            v3[65535] = ""
                                                            fmt.Println(v5)
                                                        }
                                                        v29 = v29
                                                        fmt.Println(v5, v1)
                                                        var v41 int8 = 42
                                                        if (v29 || (-76 > -1)) {
                                                            switch 0 {
                                                            case 32768:
                                                                switch "" {
                                                                case "hello":
                                                                    fmt.Println(v5, v28)
                                                                case "foo bar":
                                                                    switch "hello" {
                                                                    case "a":
                                                                        fmt.Println(v29)
                                                                    case "":
                                                                        var v42 int64 = 42
                                                                        _ = v42
                                                                    case "foo bar":
                                                                        v27 = append(v27, fi9)
                                                                    default:
                                                                        v10 += v10
                                                                    }
                                                                    if len(v27) >= 2 {
                                                                        v27 = v27[1:len(v27)]
                                                                    }
                                                                case "a":
                                                                    var v43 uint8 = 68
                                                                    _ = v43
                                                                default:
                                                                    var v44 int8 = int8(fi9)
                                                                    var v45 uint32 = 1
                                                                    v46 := GenS5{F0: (("" + "a") + "a"), F1: 1, F2: 65535}
                                                                    _ = v44
                                                                    _ = v45
                                                                    _ = v46
                                                                }
                                                            case 2147483647:
                                                                v47 := GenS6{F0: 2147483647, F1: 65535, F2: 0}
                                                                var v48 uint32 = 2147483648
                                                                var v49 int8 = 127
                                                                _ = len(v3)
                                                                v48 = (v48 % (((v48 / (((^(v48))) | 1))) | 1))
                                                                _ = v47
                                                                _ = v48
                                                                _ = v49
                                                            case 42:
                                                                if len(v27) >= 2 {
                                                                    v27 = v27[1:len(v27)]
                                                                }
                                                            default:
                                                                var v50 int32 = v8
                                                                _ = v50
                                                            }
                                                            var v51 int32 = (v8 ^ v1)
                                                            _ = v51
                                                        }
                                                        switch fi9 {
                                                        case 42:
                                                            v27 = append(v27, (fi7 << ((2) & 63)))
                                                        case 1:
                                                            v28 += v28
                                                        }
                                                        _ = v29
                                                        _ = v41
                                                    }
                                                    var v52 int16 = 1000
                                                    v53 := GenS1{F0: -1, F1: 1}
                                                    v5 += 18446744073709551615
                                                    if false {
                                                        if ((^(v52)) >= 0) {
                                                            var v54 int8 = 0
                                                            _ = len(v27)
                                                            _ = v54
                                                        } else {
                                                            var v55 bool = false
                                                            switch ("hello" + ("\n\t" + "foo bar")) {
                                                            case "\n\t":
                                                                if len(v27) >= 2 {
                                                                    v27 = v27[1:len(v27)]
                                                                }
                                                            case "a":
                                                                _ = len(v2)
                                                            }
                                                            var v56 int8 = -128
                                                            var v57 float64 = 1e10
                                                            _ = v55
                                                            _ = v56
                                                            _ = v57
                                                        }
                                                        v28 -= v28
                                                        if len(v27) >= 2 {
                                                            v27 = v27[1:len(v27)]
                                                        }
                                                    }
                                                    _ = v52
                                                    _ = v53
                                                }
                                                switch fi9 {
                                                case -881:
                                                    if (1 <= 887) {
                                                        v58 := GenS1{F0: 0, F1: int8(v5)}
                                                        var v59 int8 = -127
                                                        _ = v58
                                                        _ = v59
                                                    } else {
                                                        if (true && true) {
                                                            v60 := GenS7{F0: (("hello" + "") + (("" + ("foo bar" + "a")) + "hello"))}
                                                            v5 -= v5
                                                            _ = v60
                                                        }
                                                    }
                                                case 127:
                                                    v5 = v5
                                                case 0:
                                                    v61 := GenS6{F0: 42, F1: 937, F2: 2147483648}
                                                    _ = v61
                                                }
                                            }
                                            if false {
                                                v26 += v26
                                            }
                                            if (!(164 < 42)) {
                                                fmt.Println(fi9)
                                            }
                                        } else {
                                            if (!true) {
                                                var v62 uint8 = uint8(v5)
                                                v62 = v26
                                                _ = v62
                                            }
                                            v63 := GenS3{F0: 65535, F1: 487, F2: 670}
                                            fmt.Println(v8, v8)
                                            if (v10 <= v10) {
                                                v8 = (v1 << ((477) & 63))
                                                var v64 int64 = 530
                                                _ = v64
                                            } else {
                                                v65 := GenS6{F0: 267, F1: 42, F2: 2}
                                                var v66 int8 = 0
                                                _ = v65
                                                _ = v66
                                            }
                                            _ = v63
                                        }
                                    default:
                                        v1 -= (v1 ^ (v1 >> ((0) & 63)))
                                    }
                                    _ = v27
                                    _ = v28
                                } else {
                                    v1 -= v8
                                }
                                _ = v26
                            }
                            _ = len(v2)
                            _ = v23
                        }
                        if (2 != 1) {
                            fmt.Println(v10, v5, v10)
                            var v67 uint = 992
                            var v68 float32 = 0.0
                            _ = v67
                            _ = v68
                        } else {
                            v10 = -2.25
                            v69 := GenS8{F0: ("a" + ("\n\t" + "a")), F1: 42}
                            _ = v69
                        }
                    }
                }
            default:
                fmt.Println(v10)
            }
            fmt.Println(v8, v10)
            var v70 int64 = -135
            v71 := GenS5{F0: "a", F1: -987, F2: 32767}
            var v72 uint64 = (v5 + (v5 >> ((1) & 63)))
            if (true || ("foo bar" == "\n\t")) {
                var v73 uint64 = v5
                var v74 int32 = 1
                _ = v73
                _ = v74
            } else {
                fmt.Println(v10, fi7)
                switch fi9 {
                case -942:
                    v75 := []int{}
                    _ = v75
                }
                var v76 float64 = 0.1
                v77 := GenS9{F0: 1}
                _ = v76
                _ = v77
            }
            _ = fi9
            _ = v10
            _ = v11
            _ = v70
            _ = v71
            _ = v72
        }
        _ = fi7
        _ = v8
    }
    var v78 int64 = -1
    v79 := GenS5{F0: "foo bar", F1: 32767, F2: 65535}
    var v80 string = ""
    v81 := []int{128, 127, 473, 9223372036854775807}
    if false {
        v1 = 42
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v78:", v78)
    fmt.Println("v80:", v80)
    fmt.Println("slicelen v81:", len(v81), v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v78
    _ = v79
    _ = v80
    _ = v81
}
