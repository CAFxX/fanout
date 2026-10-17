package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint32; F1 uint16; F2 int64 }
type GenS6 struct { F0 string; F1 int32 }
type GenS7 struct { F0 bool; F1 int32 }
func genFn0(p0 int16, p1 int32) int64 {
    fmt.Println(p1, p0)
    var v7 int8 = 100
    v8 := []int{9223372036854775807, 0}
    if false {
        for fi9 := 0; fi9 < 11; fi9++ {
            fmt.Println(v7, p0, v7)
            var v10 uint16 = 2
            _ = fi9
            _ = v10
        }
        for fi11 := 0; fi11 < 20; fi11++ {
            var v12 int = -9223372036854775808
            for fi13 := 0; fi13 < 3; fi13++ {
                fmt.Println(fi11)
                _ = fi13
            }
            _ = fi11
            _ = v12
        }
    } else {
        var v14 uint8 = 116
        _ = v14
    }
    _ = len(v8)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 341
}
func genFn1(p0 bool) int32 {
    v15 := []int{-1, 0, 127}
    var v16 bool = p0
    var v17 int = -852
    if (!(-2147483648 != 2147483647)) {
        var v18 int8 = -2
        _ = v18
    } else {
        p0 = (1 >= 0)
        var v19 bool = p0
        _ = v19
    }
    _ = p0
    _ = v15
    _ = v16
    _ = v17
    return 295
}

func main() {
    var v1 uint64 = 4294967295
    var v2 string = ("foo bar" + "\n\t")
    var v3 int8 = 1
    var v4 string = v2
    v5 := make(map[string]int)
    v6 := make(map[string]string)
    v20 := GenS1{F0: 128, F1: -1}
    v4 = ""
    for fi21 := 0; fi21 < 18; fi21++ {
        if (42 < 255) {
            if true {
                v22 := GenS1{F0: (fi21 ^ (-(fi21))), F1: 903}
                _ = len(v6)
                _ = v22
            } else {
                v4 = v4
                fmt.Println(fi21, fi21)
            }
            switch fi21 {
            case 128:
                v23 := []int{fi21}
                _ = v23
            case 127:
                if true {
                    v2 = "foo bar"
                    fmt.Println(v1, v4)
                    var v24 uint = 42
                    _ = v24
                } else {
                    var v25 int16 = -11
                    _ = v25
                }
            case -435:
                var v26 uint16 = 105
                _ = v26
            }
        }
        v27 := GenS1{F0: fi21, F1: int16(v1)}
        _ = fi21
        _ = v27
    }
    var v28 float64 = ((-2.25 + -2.25) / (1.5 + 1e10))
    v29 := GenS1{F0: 212, F1: 1}
    for fi30 := 0; fi30 < 5; fi30++ {
        _ = len(v5)
        v31 := []int{}
        v32 := GenS2{F0: 0}
        _ = fi30
        _ = v31
        _ = v32
    }
    var v33 int64 = -1
    for fi34 := 0; fi34 < 8; fi34++ {
        v35 := []int{(fi34 ^ fi34)}
        switch (-(fi34)) {
        case 1:
            if true {
                switch (("hello" + v2) + v4) {
                case "":
                    if (848 < 2) {
                        if (false || false) {
                            var v36 int8 = (v3 | (v3 / ((127) | 1)))
                            fmt.Println(v2, v4, v36)
                            if (true || (false || (541 != 92))) {
                                if (((-(v33)) != v33) || (32767 > 32767)) {
                                    if len(v35) >= 2 {
                                        v35 = v35[1:len(v35)]
                                    }
                                }
                                var v37 int64 = v33
                                _ = v37
                            } else {
                                for fi38 := 0; fi38 < 3; fi38++ {
                                    if (1000 > 0) {
                                        var v39 float64 = ((1e10 / ((v28 * -2.25) + ((-2.25 - v28) - 3.14159))) - float64(((1e10 + (-2.25 - 1e10)) / 0.1)))
                                        var v40 uint8 = 42
                                        if false {
                                            v41 := GenS2{F0: -498}
                                            _ = v41
                                        } else {
                                            if true {
                                                v42 := GenS1{F0: fi38, F1: -829}
                                                _ = v42
                                            } else {
                                                v2 = (v2 + v4)
                                                if false {
                                                    var v43 int32 = 1
                                                    _ = v43
                                                }
                                            }
                                            var v44 int = fi38
                                            _ = v44
                                        }
                                        _ = v39
                                        _ = v40
                                    } else {
                                        v33 = v33
                                        var v45 int8 = 51
                                        v46 := GenS2{F0: -32768}
                                        v47 := v5[("hello" + "foo bar")]
                                        _ = v45
                                        _ = v46
                                        _ = v47
                                    }
                                    var v48 int = fi34
                                    switch 941 {
                                    case 0:
                                        var v49 string = v4
                                        _ = v49
                                    case 4294967295:
                                        if len(v35) > 0 {
                                            v35[0] = (v48 / ((127) | 1))
                                        }
                                    default:
                                        if true {
                                            var v50 int32 = 42
                                            var v51 int = (v48 & v48)
                                            var v52 int = fi34
                                            v53 := GenS3{F0: 1}
                                            _ = v50
                                            _ = v51
                                            _ = v52
                                            _ = v53
                                        } else {
                                            var v54 bool = false
                                            if true {
                                                fmt.Println(v48, v4)
                                                if (int64(v48) <= 1) {
                                                    _ = len(v35)
                                                    v55 := GenS1{F0: (fi38 * 32768), F1: 1}
                                                    _ = v55
                                                }
                                                if len(v35) > 0 {
                                                    v35[0] = fi34
                                                }
                                                switch fi38 {
                                                case -250:
                                                    fmt.Println(v33, v48)
                                                case 2147483647:
                                                    if (v54 || v54) {
                                                        var v56 uint = 42
                                                        if ((0 <= 917) && v54) {
                                                            var v57 uint32 = 2147483648
                                                            var v58 int64 = 0
                                                            var v59 uint16 = 255
                                                            _ = v57
                                                            _ = v58
                                                            _ = v59
                                                        }
                                                        _ = v56
                                                    } else {
                                                        if v54 {
                                                            v3 += 100
                                                        } else {
                                                            switch "foo bar" {
                                                            case "hello":
                                                                switch 101 {
                                                                case 9223372036854775807:
                                                                    v2 = "a"
                                                                case -549:
                                                                    fmt.Println(v33, v36, fi34)
                                                                default:
                                                                    if (v54 || false) {
                                                                        switch "a" {
                                                                        case "hello":
                                                                            v1 -= v1
                                                                        case "":
                                                                            switch -9223372036854775808 {
                                                                            case -314:
                                                                                switch v4 {
                                                                                case "hello":
                                                                                    var v60 int = (v48 % ((-2147483648) | 1))
                                                                                    _ = v60
                                                                                }
                                                                            case 2147483647:
                                                                                v1 = 18446744073709551615
                                                                            default:
                                                                                v61 := v6[("foo bar" + (v4 + "foo bar"))]
                                                                                _ = v61
                                                                            }
                                                                        default:
                                                                            if true {
                                                                                var v62 uint16 = 1
                                                                                _ = v62
                                                                            } else {
                                                                                switch (fi34 & 0) {
                                                                                case 9223372036854775807:
                                                                                    fmt.Println(v33, fi38)
                                                                                case 42:
                                                                                    v63 := GenS3{F0: 4294967295}
                                                                                    _ = v63
                                                                                default:
                                                                                    v1 += v1
                                                                                }
                                                                                v1 += (^(v1))
                                                                            }
                                                                        }
                                                                        if v54 {
                                                                            if ((123 == 42) || (!((v1 ^ v1) < v1))) {
                                                                                if v54 {
                                                                                    v36 = v3
                                                                                } else {
                                                                                    v48 -= fi38
                                                                                    var v64 int64 = 1
                                                                                    var v65 bool = (1e10 == 1.5)
                                                                                    _ = v64
                                                                                    _ = v65
                                                                                }
                                                                            }
                                                                            if (!(-128 == -128)) {
                                                                                v3 += (-(v36))
                                                                            }
                                                                            if len(v35) > 0 {
                                                                                v35[0] = 42
                                                                            }
                                                                        }
                                                                        v66 := GenS4{F0: v33}
                                                                        _ = v66
                                                                    }
                                                                }
                                                            }
                                                            var v67 uint = 4294967295
                                                            fmt.Println(v3)
                                                            _ = v67
                                                        }
                                                        v3 -= 100
                                                    }
                                                case -1:
                                                    v4 = (("hello" + "a") + v4)
                                                }
                                                fmt.Println(v4)
                                                v28 += (1e10 - -2.25)
                                                v68 := GenS5{F0: 333, F1: 2, F2: (v33 ^ v33)}
                                                _ = v68
                                            }
                                            _ = len(v5)
                                            _ = v54
                                        }
                                        var v69 bool = ((-(fi38)) == v48)
                                        _ = v69
                                    }
                                    _ = fi38
                                    _ = v48
                                }
                                _ = len(v35)
                                v70 := GenS3{F0: 557}
                                _ = v70
                            }
                            _ = v36
                        } else {
                            v71 := GenS3{F0: 42}
                            _ = v71
                        }
                        if (2147483648 >= 296) {
                            for fi72 := 0; fi72 < 1; fi72++ {
                                v35 = append(v35, fi72)
                                var v73 uint32 = 42
                                _ = fi72
                                _ = v73
                            }
                            var v74 bool = true
                            _ = v74
                        }
                        v1 = v1
                    }
                case "\n\t":
                    fmt.Println(v28, v4)
                }
                var v75 uint32 = 0
                _ = v75
            } else {
                v76 := GenS6{F0: (v2 + v4), F1: -1}
                var v77 int16 = 32767
                _ = v76
                _ = v77
            }
        default:
            v78 := GenS5{F0: 65535, F1: 900, F2: (v33 & -9223372036854775808)}
            _ = v78
        }
        _ = fi34
        _ = v35
    }
    v28 = (0.0 - v28)
    v79 := []int{-2147483648}
    if (((true || false) || (!false)) && ("" >= (("a" + v4) + (("foo bar" + "hello") + (("foo bar" + "foo bar") + "a"))))) {
        for fi80 := 0; fi80 < 2; fi80++ {
            fmt.Println(v28, v2, v4)
            _ = fi80
        }
    } else {
        if (1 > -2147483648) {
            for fi81 := 0; fi81 < 9; fi81++ {
                var v82 float32 = 1e10
                if true {
                    v82 -= v82
                }
                _ = fi81
                _ = v82
            }
            for fi83 := 0; fi83 < 18; fi83++ {
                if len(v79) >= 2 {
                    v79 = v79[1:len(v79)]
                }
                if true {
                    if ((v33 & v33) <= (v33 + (v33 ^ -378))) {
                        var v84 uint8 = 128
                        _ = v84
                    } else {
                        v2 = ("\n\t" + "a")
                    }
                    fmt.Println(v2, v1)
                    v85 := GenS7{F0: (42 > 858), F1: 303}
                    _ = v85
                }
                if true {
                    for fi86 := 0; fi86 < 3; fi86++ {
                        var v87 uint16 = 572
                        _ = fi86
                        _ = v87
                    }
                } else {
                    var v88 uint8 = 128
                    for fi89 := 0; fi89 < 16; fi89++ {
                        var v90 float64 = v28
                        delete(v5, v4)
                        v28 += (v28 * 0.0)
                        _ = fi89
                        _ = v90
                    }
                    var v91 float64 = v28
                    _ = v88
                    _ = v91
                }
                _ = fi83
            }
        }
        v4 = "a"
        var v92 uint16 = 65535
        _ = v92
    }
    v4 = "hello"
    v93 := GenS7{F0: (!(false || false)), F1: 0}
    var v94 float64 = float64(3.14159)
    v2 = "hello"
    if (!(((-32768 > 1000) || true) || false)) {
        fmt.Println(v33)
        v95 := v5["foo bar"]
        switch "" {
        case "a":
            var v96 uint64 = v1
            _ = v96
        case "hello":
            delete(v5, v4)
        default:
            var v97 uint64 = (^(v1))
            _ = v97
        }
        _ = v95
    }
    v98 := GenS3{F0: 18446744073709551615}
    var v99 int16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v79:", len(v79), v79)
    fmt.Println("v94:", v94)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v2
    _ = v20
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v79
    _ = v93
    _ = v94
    _ = v98
    _ = v99
}
