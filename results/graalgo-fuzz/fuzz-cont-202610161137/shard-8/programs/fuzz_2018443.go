package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 int8; F1 uint16; F2 uint }
type GenS4 struct { F0 bool; F1 int64; F2 uint }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint64; F1 int64 }

func main() {
    var v1 int8 = 100
    var v2 uint64 = 542
    v3 := make(map[int]uint32)
    v4 := []int{}
    if (1000 <= 32767) {
        var v5 float64 = -2.25
        v1 = (v1 ^ 42)
        v6 := v3[-554]
        _ = v5
        _ = v6
    } else {
        var v7 uint32 = 407
        _ = v7
    }
    var v8 uint8 = 11
    var v9 uint64 = v2
    v10 := GenS1{F0: 255, F1: ("hello" + "foo bar")}
    var v11 uint = 918
    v1 = v1
    for fi12 := 0; fi12 < 16; fi12++ {
        v11 -= (v11 >> ((v11) & 63))
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi12
    }
    for fi13 := 0; fi13 < 1; fi13++ {
        v8 = (v8 % (((v8 % ((uint8(v11)) | 1))) | 1))
        var v14 int = fi13
        _ = fi13
        _ = v14
    }
    if (false || true) {
        var v15 uint = 929
        for fi16 := 0; fi16 < 9; fi16++ {
            if true {
                v4 = append(v4, (fi16 ^ (fi16 % ((-888) | 1))))
            } else {
                v17 := GenS1{F0: 426, F1: ""}
                _ = v17
            }
            for fi18 := 0; fi18 < 2; fi18++ {
                var v19 uint16 = 411
                if false {
                    switch 156 {
                    case -452:
                        switch fi16 {
                        case -9223372036854775808:
                            delete(v3, 1)
                        case 2147483647:
                            v20 := GenS1{F0: uint16(v8), F1: "hello"}
                            _ = v20
                        }
                    default:
                        switch "foo bar" {
                        case "":
                            if (32768 <= fi18) {
                                v1 = 101
                                if false {
                                    if false {
                                        if ((!(-698 != -9223372036854775808)) || true) {
                                            fmt.Println(fi18, v15, v19)
                                        } else {
                                            v21 := GenS2{F0: 42, F1: (v8 - (v8 % ((v8) | 1)))}
                                            v11 -= v11
                                            v4 = append(v4, fi18)
                                            _ = v21
                                        }
                                        var v22 int32 = -307
                                        var v23 int16 = 491
                                        _ = v22
                                        _ = v23
                                    }
                                    if true {
                                        if (!true) {
                                            v1 += (v1 & v1)
                                        }
                                    }
                                    _ = len(v4)
                                }
                                if false {
                                    switch "hello" {
                                    case "":
                                        var v24 int8 = 1
                                        _ = v24
                                    }
                                    v8 = (v8 >> (((v15 / (((^(v11))) | 1))) & 63))
                                    v25 := v3[int(v2)]
                                    _ = v25
                                } else {
                                    fmt.Println(fi18, v15, fi16)
                                    switch fi16 {
                                    case 32768:
                                        var v26 uint32 = 889
                                        _ = v26
                                    case -9223372036854775808:
                                        v1 += v1
                                    }
                                    if (((true && false) || false) && ((v1 * -128) > v1)) {
                                        v2 = (^(v2))
                                        v8 += (^(v8))
                                        if true {
                                            fmt.Println(v8, v15)
                                            var v27 float64 = (0.0 + 0.1)
                                            _ = v27
                                        } else {
                                            var v28 uint32 = 299
                                            switch "hello" {
                                            case "a":
                                                fmt.Println(v2, v1, v2)
                                            default:
                                                fmt.Println(v28, v19)
                                            }
                                            _ = v28
                                        }
                                    }
                                    switch "\n\t" {
                                    case "hello":
                                        if false {
                                            v2 = uint64(v8)
                                        } else {
                                            if false {
                                                v29 := GenS2{F0: int8(v8), F1: 163}
                                                _ = v29
                                            }
                                            if true {
                                                v2 += v9
                                                if (v15 > uint(v8)) {
                                                    fmt.Println(v9)
                                                } else {
                                                    v11 = v15
                                                    fmt.Println(fi16)
                                                    var v30 int32 = 2147483647
                                                    _ = v30
                                                }
                                            }
                                            v31 := GenS2{F0: -1, F1: (v8 % ((243) | 1))}
                                            _ = v31
                                        }
                                    case "a":
                                        fmt.Println(v8, fi16)
                                    }
                                    switch "\n\t" {
                                    case "foo bar":
                                        delete(v3, fi18)
                                    }
                                }
                            } else {
                                switch (fi16 ^ fi18) {
                                case -237:
                                    var v32 float32 = 3.14159
                                    _ = v32
                                default:
                                    var v33 float64 = (-2.25 + 0.0)
                                    _ = v33
                                }
                            }
                        case "\n\t":
                            v11 += (v11 % ((v15) | 1))
                        default:
                            if true {
                                var v34 int64 = 42
                                v35 := GenS2{F0: (v1 | int8(v8)), F1: v8}
                                _ = v34
                                _ = v35
                            } else {
                                switch "hello" {
                                case "foo bar":
                                    if (-1 == int16(v1)) {
                                        if true {
                                            switch -9223372036854775808 {
                                            case 0:
                                                switch (fi18 >> ((v15) & 63)) {
                                                case -2147483648:
                                                    var v36 uint16 = 32767
                                                    _ = v36
                                                case 2147483647:
                                                    v2 -= v9
                                                case -751:
                                                    if false {
                                                        if (!(599 != 1)) {
                                                            v9 = v2
                                                            var v37 int = fi16
                                                            if (v9 < v2) {
                                                                if true {
                                                                    if len(v4) > 0 {
                                                                        v4[0] = fi18
                                                                    }
                                                                } else {
                                                                    v15 += (v15 | uint(v9))
                                                                }
                                                                var v38 uint64 = 592
                                                                _ = v38
                                                            } else {
                                                                switch v37 {
                                                                case 2147483647:
                                                                    delete(v3, (fi16 & (-(fi18))))
                                                                case 32767:
                                                                    fmt.Println(v15, v19, fi18)
                                                                case 32768:
                                                                    v4 = append(v4, 4294967295)
                                                                }
                                                            }
                                                            _ = v37
                                                        }
                                                        switch "a" {
                                                        case "foo bar":
                                                            fmt.Println(v11)
                                                        }
                                                        var v39 int16 = 1000
                                                        _ = v39
                                                    } else {
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            switch ("hello" + "\n\t") {
                                                            case "foo bar":
                                                                var v40 string = (("hello" + "") + "hello")
                                                                _ = v40
                                                            }
                                                        }
                                                        var v41 uint8 = 146
                                                        var v42 uint8 = (v8 % ((128) | 1))
                                                        var v43 float32 = 0.1
                                                        _ = v41
                                                        _ = v42
                                                        _ = v43
                                                    }
                                                }
                                            }
                                        } else {
                                            var v44 string = "hello"
                                            var v45 uint32 = uint32(v8)
                                            var v46 string = "a"
                                            _ = v44
                                            _ = v45
                                            _ = v46
                                        }
                                        var v47 string = "foo bar"
                                        _ = v47
                                    }
                                case "":
                                    fmt.Println(fi18, v19)
                                case "hello":
                                    v48 := GenS1{F0: 2, F1: (("" + "\n\t") + ("hello" + "foo bar"))}
                                    _ = v48
                                }
                                var v49 int16 = -32768
                                if false {
                                    v50 := GenS2{F0: (v1 & -1), F1: 3}
                                    fmt.Println(v8)
                                    _ = v50
                                } else {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    switch "" {
                                    case "\n\t":
                                        v9 = v9
                                    case "":
                                        v51 := v3[292]
                                        _ = v51
                                    default:
                                        fmt.Println(v19)
                                    }
                                    var v52 string = ((("foo bar" + "") + "hello") + "a")
                                    _ = v52
                                }
                                switch ("foo bar" + "\n\t") {
                                case "hello":
                                    v53 := GenS3{F0: (v1 >> ((255) & 63)), F1: v19, F2: 1}
                                    _ = v53
                                case "\n\t":
                                    switch (("\n\t" + ((("hello" + "a") + "hello") + "\n\t")) + "a") {
                                    case "\n\t":
                                        var v54 int16 = 1
                                        _ = v54
                                    case "a":
                                        var v55 uint16 = v19
                                        _ = v55
                                    case "foo bar":
                                        var v56 bool = true
                                        _ = v56
                                    }
                                case "foo bar":
                                    fmt.Println(v9, v15)
                                }
                                var v57 string = "hello"
                                if false {
                                    switch v57 {
                                    case "foo bar":
                                        var v58 uint32 = 2147483647
                                        _ = v58
                                    case "a":
                                        var v59 int32 = -2147483648
                                        _ = v59
                                    case "\n\t":
                                        if (!(v49 >= (v49 & v49))) {
                                            if true {
                                                v60 := GenS2{F0: 42, F1: (v8 & (v8 / (((v8 ^ (v8 + 171))) | 1)))}
                                                fmt.Println(v57, v2, v2)
                                                _ = v60
                                            }
                                            v11 += v15
                                        }
                                    }
                                    switch fi18 {
                                    case 9223372036854775807:
                                        switch ((v57 + "\n\t") + "foo bar") {
                                        case "":
                                            var v61 int32 = -1
                                            _ = v61
                                        case "hello":
                                            v49 += (v49 / (((v49 << ((v15) & 63))) | 1))
                                        case "foo bar":
                                            v11 -= v15
                                        }
                                    case 127:
                                        _ = len(v4)
                                    case 2147483647:
                                        v62 := GenS2{F0: v1, F1: v8}
                                        _ = v62
                                    default:
                                        var v63 int16 = (^(v49))
                                        if (false && (42 < 19)) {
                                            fmt.Println(v19, v15)
                                            var v64 uint64 = v9
                                            _ = v64
                                        } else {
                                            switch 1 {
                                            case 32768:
                                                _ = len(v4)
                                            case 42:
                                                fmt.Println(v1, v11, v8)
                                            }
                                            if (158 >= v8) {
                                                var v65 int16 = (v63 ^ (^(v49)))
                                                v66 := GenS4{F0: (false || (false && true)), F1: 9223372036854775807, F2: (v15 << ((255) & 63))}
                                                v8 += 128
                                                _ = v65
                                                _ = v66
                                            } else {
                                                v67 := GenS2{F0: 0, F1: (v8 % ((42) | 1))}
                                                var v68 uint = v15
                                                _ = v67
                                                _ = v68
                                            }
                                            _ = len(v3)
                                        }
                                        _ = v63
                                    }
                                } else {
                                    if true {
                                        v1 -= (v1 - (v1 % ((-1) | 1)))
                                    } else {
                                        v69 := GenS2{F0: v1, F1: 2}
                                        if false {
                                            v70 := GenS5{F0: v1}
                                            delete(v3, fi16)
                                            v15 = v15
                                            _ = v70
                                        }
                                        var v71 uint16 = v19
                                        v19 = (v19 / ((2) | 1))
                                        _ = v69
                                        _ = v71
                                    }
                                }
                                _ = v49
                                _ = v57
                            }
                        }
                    }
                    v72 := GenS4{F0: true, F1: 9223372036854775807, F2: 255}
                    v73 := GenS6{F0: (v19 >> ((42) & 63))}
                    var v74 string = "a"
                    _ = v72
                    _ = v73
                    _ = v74
                } else {
                    v15 -= (v11 | 18446744073709551615)
                    switch fi18 {
                    case 128:
                        var v75 int16 = 13
                        _ = v75
                    case 159:
                        v76 := GenS1{F0: 42, F1: "a"}
                        _ = v76
                    case -2147483648:
                        var v77 int16 = int16(v9)
                        _ = v77
                    default:
                        var v78 bool = false
                        _ = v78
                    }
                }
                switch fi18 {
                case 42:
                    var v79 int16 = 295
                    _ = v79
                default:
                    v15 -= 226
                }
                if (3.14159 == (0.1 + (0.0 / 1.0))) {
                    if true {
                        _ = len(v4)
                        var v80 uint16 = (v19 % (((^(v19))) | 1))
                        v81 := GenS1{F0: 32767, F1: "\n\t"}
                        v15 = (v15 + v15)
                        _ = v80
                        _ = v81
                    }
                    v82 := GenS2{F0: v1, F1: 128}
                    v11 = v15
                    _ = v82
                } else {
                    if false {
                        v15 -= v15
                        v3[(fi16 >> ((2) & 63))] = 2147483648
                        var v83 uint = (v15 * (v11 ^ (v15 ^ (v15 / ((42) | 1)))))
                        _ = v83
                    } else {
                        v84 := GenS3{F0: v1, F1: 2, F2: (v15 + (^(v11)))}
                        fmt.Println(v2, fi18)
                        _ = v84
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                _ = fi18
                _ = v19
            }
            if ("hello" == "foo bar") {
                v9 += (v9 % ((18446744073709551615) | 1))
                v85 := GenS7{F0: 974, F1: 1}
                var v86 bool = ((1 == 1) || (true && (!(2147483648 < 4294967295))))
                _ = v85
                _ = v86
            } else {
                var v87 int32 = 146
                if len(v4) > 0 {
                    v4[0] = int(v8)
                }
                fmt.Println(v8)
                _ = v87
            }
            _ = fi16
        }
        v88 := GenS5{F0: v1}
        _ = v15
        _ = v88
    } else {
        v1 = (v1 & v1)
        fmt.Println(v1, v9)
        v1 += v1
    }
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
