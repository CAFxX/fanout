package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool; F1 int8; F2 uint16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int8; F1 bool }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint16; F1 uint32; F2 bool }
type GenS10 struct { F0 uint; F1 uint64 }
type GenS11 struct { F0 uint32; F1 int8 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 int; F1 bool; F2 int16 }
func genFn0() int32 {
    v4 := GenS1{F0: 0, F1: 576}
    v5 := GenS1{F0: 65535, F1: -318}
    var v6 float64 = float64(0.1)
    v7 := []int{2147483647, -52, 4294967295, 32767}
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return -673
}
func genFn1(p0 int32, p1 uint32, p2 int32) uint8 {
    fmt.Println(p2, p1, p1)
    p1 += (^(p1))
    v8 := []int{32767, 532, 2147483647, 42}
    if ((true && (false || false)) || false) {
        if false {
            switch 32767 {
            case 128:
                fmt.Println(p0, p1)
            default:
                v8 = append(v8, -672)
            }
        }
        fmt.Println(p0)
        fmt.Println(p1, p2)
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 uint = 18446744073709551615
        for fi11 := 0; fi11 < 19; fi11++ {
            if true {
                var v12 int32 = p0
                _ = v12
            }
            v13 := GenS1{F0: 65535, F1: 940}
            if (1 != p2) {
                var v14 bool = true
                p0 += p0
                var v15 int8 = 1
                _ = v14
                _ = v15
            }
            if ((true || (0.0 >= 1e10)) && ((42 > 1) && false)) {
                fmt.Println(p2)
                var v16 uint = v10
                if (p1 >= (p1 & 193)) {
                    v17 := GenS2{F0: (fi9 << ((v10) & 63))}
                    v18 := GenS3{F0: 255}
                    v10 += v10
                    _ = v17
                    _ = v18
                } else {
                    if true {
                        if (true || true) {
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                            v8 = append(v8, fi11)
                            if len(v8) > 0 {
                                v8[0] = fi9
                            }
                        }
                        p2 -= (-(p2))
                        p2 = (p2 % (((p2 & 2147483647)) | 1))
                    }
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                }
                _ = v16
            } else {
                fmt.Println(p1, fi11, p2)
            }
            var v19 string = ((("a" + "foo bar") + "\n\t") + "hello")
            _ = fi11
            _ = v13
            _ = v19
        }
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 255
}

func main() {
    var v1 string = ""
    var v2 float32 = 0.0
    v3 := make(map[string]int)
    var v20 uint8 = 128
    v21 := GenS3{F0: 1}
    switch -2147483648 {
    case 0:
        var v22 int32 = -2147483648
        _ = v22
    case -9223372036854775808:
        var v23 string = v1
        _ = v23
    }
    var v24 int32 = 281
    for fi25 := 0; fi25 < 19; fi25++ {
        v26 := GenS4{F0: (true && false), F1: 50, F2: 2}
        _ = fi25
        _ = v26
    }
    switch (((v1 + ((v1 + v1) + "")) + v1) + v1) {
    case "\n\t":
        for fi27 := 0; fi27 < 10; fi27++ {
            var v28 int16 = 1
            fmt.Println(v1, v20)
            fmt.Println(v28, v24)
            _ = fi27
            _ = v28
        }
    case "":
        v20 = (v20 / ((v20) | 1))
    case "hello":
        if (192 == 599) {
            var v29 int64 = 9223372036854775807
            for fi30 := 0; fi30 < 1; fi30++ {
                if false {
                    v31 := GenS5{F0: v20}
                    for fi32 := 0; fi32 < 15; fi32++ {
                        if false {
                            v24 -= genFn0()
                            if (v1 < ("" + v1)) {
                                v33 := GenS4{F0: ((-744 >= -529) || (true || (-100 >= 51))), F1: 1, F2: 32767}
                                if ((2 != 825) || (18446744073709551615 >= 0)) {
                                    v1 = ""
                                    delete(v3, (v1 + ((v1 + (("" + "a") + "")) + v1)))
                                    var v34 int16 = 1000
                                    _ = v34
                                } else {
                                    v1 = v1
                                }
                                fmt.Println(fi32, fi30, v2)
                                _ = v33
                            }
                            switch "a" {
                            case "\n\t":
                                switch "a" {
                                case "\n\t":
                                    v35 := GenS6{F0: -1, F1: true}
                                    _ = v35
                                case "hello":
                                    v36 := GenS3{F0: 42}
                                    _ = v36
                                case "foo bar":
                                    v2 = (float32(1e10) - (v2 - v2))
                                    delete(v3, ("a" + (((v1 + v1) + (v1 + ((v1 + "hello") + (v1 + (v1 + "foo bar"))))) + v1)))
                                default:
                                    _ = len(v3)
                                }
                            case "":
                                fmt.Println(v1, v1)
                                fmt.Println(v2, v2, fi30)
                                var v37 int = (fi32 - int(v20))
                                _ = v37
                            case "foo bar":
                                v38 := []int{}
                                _ = v38
                            }
                            v39 := GenS5{F0: v20}
                            _ = v39
                        }
                        switch 2147483647 {
                        case 32767:
                            _ = len(v3)
                        case 127:
                            v40 := GenS7{F0: 42}
                            _ = v40
                        }
                        v41 := GenS6{F0: 100, F1: true}
                        fmt.Println(v24)
                        _ = fi32
                        _ = v41
                    }
                    v42 := GenS6{F0: -1, F1: false}
                    _ = v31
                    _ = v42
                }
                _ = fi30
            }
            v43 := []int{4294967295}
            _ = v29
            _ = v43
        }
    }
    v1 = "a"
    switch v1 {
    case "":
        switch -2147483648 {
        case 42:
            fmt.Println(v2)
        case 0:
            if ((v20 - v20) < 255) {
                var v44 uint16 = 0
                _ = v44
            } else {
                switch -2147483648 {
                case 127:
                    var v45 int8 = 42
                    _ = v45
                case 223:
                    for fi46 := 0; fi46 < 10; fi46++ {
                        fmt.Println(v24, v1)
                        fmt.Println(v2)
                        v47 := []int{}
                        _ = fi46
                        _ = v47
                    }
                case -1:
                    var v48 uint8 = v20
                    _ = v48
                default:
                    for fi49 := 0; fi49 < 8; fi49++ {
                        v2 += v2
                        var v50 uint8 = (v20 * (^(v20)))
                        _ = fi49
                        _ = v50
                    }
                }
                v51 := GenS8{F0: -824}
                _ = v51
            }
        case 127:
            for fi52 := 0; fi52 < 11; fi52++ {
                switch (fi52 / (((-(fi52))) | 1)) {
                case 42:
                    switch (("a" + v1) + "\n\t") {
                    case "foo bar":
                        for fi53 := 0; fi53 < 1; fi53++ {
                            var v54 string = v1
                            if (42 < -128) {
                                v1 = v54
                                if true {
                                    v55 := v3[("\n\t" + "foo bar")]
                                    if true {
                                        v56 := GenS9{F0: 255, F1: 2, F2: true}
                                        if false {
                                            var v57 string = "hello"
                                            var v58 uint8 = (v20 ^ 127)
                                            v1 = "foo bar"
                                            _ = v57
                                            _ = v58
                                        }
                                        v59 := GenS10{F0: 65535, F1: 42}
                                        _ = v56
                                        _ = v59
                                    } else {
                                        var v60 int = fi53
                                        v61 := GenS10{F0: 65535, F1: 2}
                                        fmt.Println(v54, v1, fi52)
                                        if (!(("a" + "hello") < (v54 + "a"))) {
                                            v54 = (("" + (v1 + v54)) + v54)
                                        } else {
                                            _ = len(v3)
                                        }
                                        _ = v60
                                        _ = v61
                                    }
                                    _ = v55
                                } else {
                                    v54 = "a"
                                    v62 := GenS11{F0: 1, F1: 84}
                                    if false {
                                        var v63 uint64 = 1
                                        v2 = v2
                                        fmt.Println(v63)
                                        _ = v63
                                    }
                                    var v64 int16 = 1
                                    _ = v62
                                    _ = v64
                                }
                                delete(v3, v54)
                            } else {
                                fmt.Println(v54, fi53)
                            }
                            var v65 int8 = 42
                            var v66 string = (("a" + "foo bar") + (v54 + "foo bar"))
                            switch v66 {
                            case "a":
                                v54 = ("a" + "hello")
                            case "\n\t":
                                v65 = (v65 ^ v65)
                            case "foo bar":
                                switch v54 {
                                case "":
                                    v67 := GenS11{F0: 1, F1: (^(v65))}
                                    _ = v67
                                case "hello":
                                    v68 := GenS12{F0: 65535}
                                    _ = v68
                                case "a":
                                    v20 = (v20 | 128)
                                default:
                                    v69 := []int{32767, fi52, (fi52 * fi53)}
                                    _ = v69
                                }
                            default:
                                if false {
                                    v20 = genFn1(genFn0(), 42, v24)
                                    var v70 string = (("" + "") + (v66 + "foo bar"))
                                    switch v70 {
                                    case "\n\t":
                                        fmt.Println(v70)
                                    default:
                                        v71 := GenS8{F0: 0}
                                        _ = v71
                                    }
                                    v70 = "\n\t"
                                    _ = v70
                                }
                            }
                            switch v54 {
                            case "\n\t":
                                var v72 uint16 = 675
                                _ = v72
                            case "":
                                v54 = (v1 + v54)
                                v24 -= 968
                            case "a":
                                var v73 uint = 451
                                _ = v73
                            }
                            _ = fi53
                            _ = v54
                            _ = v65
                            _ = v66
                        }
                    case "a":
                        if ((^(v20)) <= (v20 & (v20 % (((v20 * 128)) | 1)))) {
                            v74 := []int{fi52, (^(fi52)), fi52}
                            _ = v74
                        } else {
                            if (!(v20 != (^(v20)))) {
                                v75 := GenS13{F0: 1}
                                if false {
                                    v76 := GenS14{F0: v24}
                                    _ = v76
                                }
                                _ = v75
                            }
                            v24 = (v24 >> ((451) & 63))
                        }
                    case "":
                        v77 := []int{-1, fi52, fi52}
                        _ = v77
                    default:
                        v24 = v24
                    }
                case 128:
                    v2 += v2
                default:
                    for fi78 := 0; fi78 < 13; fi78++ {
                        v1 = v1
                        _ = fi78
                    }
                }
                if (!true) {
                    if true {
                        var v79 int16 = 1000
                        v20 = (v20 - (v20 / (((v20 / ((genFn1(v24, 2147483648, 2147483647)) | 1))) | 1)))
                        if true {
                            v20 += v20
                            v79 = v79
                        }
                        _ = v79
                    }
                    var v80 uint = 550
                    _ = v80
                } else {
                    if true {
                        for fi81 := 0; fi81 < 16; fi81++ {
                            v82 := GenS14{F0: v24}
                            v20 = v20
                            if false {
                                var v83 uint = 365
                                var v84 int32 = (v24 >> (((v83 % ((v83) | 1))) & 63))
                                _ = v83
                                _ = v84
                            }
                            _ = fi81
                            _ = v82
                        }
                    } else {
                        v85 := []int{fi52, (fi52 | -9223372036854775808)}
                        var v86 uint = 703
                        _ = v85
                        _ = v86
                    }
                }
                _ = fi52
            }
        default:
            v3[("" + v1)] = 32767
        }
    case "foo bar":
        v24 -= genFn0()
    case "\n\t":
        switch (v1 + (v1 + v1)) {
        case "a":
            for fi87 := 0; fi87 < 2; fi87++ {
                v88 := GenS15{F0: fi87, F1: (2 > 42), F2: 303}
                _ = fi87
                _ = v88
            }
        default:
            var v89 uint32 = 2147483647
            _ = v89
        }
    }
    v90 := GenS6{F0: 127, F1: false}
    var v91 uint64 = 0
    v1 = ("\n\t" + v1)
    var v92 uint64 = v91
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v91:", v91)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v90
    _ = v91
    _ = v92
}
