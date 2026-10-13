package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 int32; F1 uint; F2 int64 }
type GenS4 struct { F0 string; F1 int8; F2 int32 }
type GenS5 struct { F0 int16; F1 int32 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int; F1 uint; F2 string }
type GenS8 struct { F0 uint8; F1 int32; F2 uint32 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int64; F1 uint32 }
type GenS11 struct { F0 uint16 }
func genFn0(p0 uint64) uint8 {
    var v6 float64 = 0.0
    var v7 uint8 = 255
    _ = p0
    _ = v6
    _ = v7
    return 255
}

func main() {
    var v1 int8 = 100
    v2 := []int{2147483647, -9223372036854775808, 907}
    v3 := make(map[uint32]string)
    var v4 int = -253
    var v5 string = ""
    var v8 int = -539
    v5 = v5
    var v9 float32 = (-2.25 * (0.0 + -2.25))
    fmt.Println(v9)
    for fi10 := 0; fi10 < 20; fi10++ {
        if true {
            _ = len(v2)
        } else {
            v11 := GenS1{F0: 42, F1: 2147483648}
            if true {
                var v12 bool = false
                v13 := GenS1{F0: 4294967295, F1: 0}
                var v14 int = v8
                _ = v12
                _ = v13
                _ = v14
            } else {
                v3[2147483648] = "a"
                var v15 uint = 42
                v4 += v4
                _ = v15
            }
            _ = len(v3)
            _ = v11
        }
        for fi16 := 0; fi16 < 18; fi16++ {
            v4 -= v8
            if false {
                v17 := GenS1{F0: 18446744073709551615, F1: 2147483648}
                _ = v17
            }
            _ = fi16
        }
        v1 = v1
        _ = fi10
    }
    if true {
        var v18 float32 = v9
        if (v1 <= v1) {
            if (1 < int16(v1)) {
                _ = len(v2)
            } else {
                fmt.Println(v1)
                for fi19 := 0; fi19 < 15; fi19++ {
                    v20 := GenS2{F0: 893, F1: (v5 >= ("a" + v5))}
                    for fi21 := 0; fi21 < 15; fi21++ {
                        v22 := GenS2{F0: 4294967295, F1: (127 > 127)}
                        _ = fi21
                        _ = v22
                    }
                    var v23 int16 = 1
                    _ = fi19
                    _ = v20
                    _ = v23
                }
                v4 -= v8
            }
        }
        var v24 int32 = 2147483647
        _ = v18
        _ = v24
    }
    var v25 uint8 = 104
    if (-1 > 922) {
        v26 := GenS3{F0: 0, F1: 255, F2: 9223372036854775807}
        if true {
            for fi27 := 0; fi27 < 17; fi27++ {
                if true {
                    if (!(95 <= 4294967295)) {
                        if true {
                            v1 = 1
                        } else {
                            for fi28 := 0; fi28 < 10; fi28++ {
                                v29 := GenS3{F0: 96, F1: 2, F2: int64(v25)}
                                switch fi28 {
                                case 128:
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                                _ = fi28
                                _ = v29
                            }
                        }
                        for fi30 := 0; fi30 < 10; fi30++ {
                            var v31 uint8 = 1
                            switch 4294967295 {
                            case 9223372036854775807:
                                v8 += 32768
                            case -434:
                                fmt.Println(v8)
                            case 42:
                                if false {
                                    v32 := GenS3{F0: -1, F1: 2, F2: -1}
                                    if (true || false) {
                                        var v33 float64 = 1.5
                                        v34 := v3[1]
                                        if true {
                                            if (!((-2147483648 > 295) || ((true && false) && false))) {
                                                v8 += 370
                                                var v35 string = (("a" + "") + "\n\t")
                                                _ = v35
                                            } else {
                                                v9 = (((v9 + ((v9 + v9) * v9)) * float32(1.5)) - float32(v33))
                                                v36 := v3[1]
                                                _ = v36
                                            }
                                        }
                                        if true {
                                            v37 := GenS3{F0: 1, F1: 255, F2: 826}
                                            fmt.Println(fi27, v34, v31)
                                            var v38 uint32 = 4294967295
                                            _ = v37
                                            _ = v38
                                        } else {
                                            var v39 uint16 = 808
                                            if true {
                                                fmt.Println(v1)
                                            } else {
                                                var v40 uint8 = (v25 % ((102) | 1))
                                                var v41 string = (v5 + "hello")
                                                _ = v40
                                                _ = v41
                                            }
                                            _ = v39
                                        }
                                        _ = v33
                                        _ = v34
                                    } else {
                                        var v42 uint64 = 2
                                        _ = v42
                                    }
                                    v1 += (v1 >> ((201) & 63))
                                    _ = v32
                                }
                            default:
                                fmt.Println(v31, v5, v8)
                            }
                            _ = fi30
                            _ = v31
                        }
                        var v43 uint = 0
                        _ = v43
                    }
                } else {
                    v44 := v3[1]
                    v45 := GenS4{F0: (v44 + v5), F1: v1, F2: int32(v1)}
                    v46 := GenS1{F0: 1, F1: 149}
                    _ = v44
                    _ = v45
                    _ = v46
                }
                var v47 bool = (0 == 0)
                _ = fi27
                _ = v47
            }
        } else {
            var v48 int8 = v1
            _ = v48
        }
        var v49 string = (((("" + ("\n\t" + v5)) + v5) + v5) + v5)
        _ = v26
        _ = v49
    } else {
        switch ("hello" + v5) {
        case "hello":
            var v50 string = (v5 + (v5 + "a"))
            _ = v50
        }
        v9 += v9
        _ = len(v3)
    }
    for fi51 := 0; fi51 < 9; fi51++ {
        switch 32767 {
        case 32768:
            v9 = v9
        }
        _ = len(v2)
        _ = fi51
    }
    switch ((("hello" + "\n\t") + v5) + v5) {
    case "foo bar":
        var v52 uint = 65535
        _ = v52
    case "\n\t":
        switch -9223372036854775808 {
        case 32768:
            for fi53 := 0; fi53 < 5; fi53++ {
                v54 := GenS5{F0: 42, F1: -1}
                var v55 float32 = ((v9 / 3.14159) - -2.25)
                _ = fi53
                _ = v54
                _ = v55
            }
        case 9223372036854775807:
            if false {
                v8 = v4
            }
        }
    case "hello":
        var v56 string = ((v5 + ((("foo bar" + ("\n\t" + v5)) + (v5 + "\n\t")) + ((v5 + v5) + "foo bar"))) + "\n\t")
        _ = v56
    default:
        var v57 uint64 = 399
        _ = v57
    }
    switch (v8 << ((65535) & 63)) {
    case 42:
        if (v1 == 127) {
            for fi58 := 0; fi58 < 18; fi58++ {
                v59 := GenS2{F0: 18446744073709551615, F1: true}
                var v60 int32 = -2147483648
                _ = fi58
                _ = v59
                _ = v60
            }
            switch v5 {
            case "a":
                v61 := GenS3{F0: 331, F1: uint(v25), F2: 915}
                _ = v61
            case "foo bar":
                var v62 uint8 = 42
                _ = v62
            default:
                switch "hello" {
                case "":
                    v4 -= (v8 >> ((42) & 63))
                case "\n\t":
                    v9 = -2.25
                }
            }
            var v63 uint32 = 356
            _ = v63
        } else {
            v64 := GenS4{F0: ("a" + v5), F1: v1, F2: 1}
            fmt.Println(v8, v9, v25)
            v1 = (v1 | v1)
            _ = v64
        }
    }
    v8 = (-(v4))
    if true {
        v65 := GenS2{F0: 458, F1: ((-1 != -1) && ((1000 <= 1000) || (1 != 4294967295)))}
        _ = len(v2)
        if (false && (!(1.5 > (3.14159 - 1.5)))) {
            v66 := GenS6{F0: v4}
            if (936 < 902) {
                v4 = v8
                v67 := GenS2{F0: 723, F1: false}
                for fi68 := 0; fi68 < 9; fi68++ {
                    var v69 float64 = 1.5
                    _ = fi68
                    _ = v69
                }
                _ = v67
            }
            var v70 float32 = 1.5
            _ = v66
            _ = v70
        }
        _ = v65
    }
    if true {
        if (2147483648 > 4294967295) {
            var v71 float32 = 0.0
            _ = v71
        } else {
            v8 = (v4 & v4)
            switch v5 {
            case "":
                for fi72 := 0; fi72 < 18; fi72++ {
                    var v73 int8 = (v1 * v1)
                    if (767 > 636) {
                        v74 := GenS7{F0: -1, F1: 2, F2: v5}
                        if len(v2) > 0 {
                            v2[0] = 334
                        }
                        _ = v74
                    } else {
                        v1 = 42
                    }
                    _ = len(v2)
                    _ = fi72
                    _ = v73
                }
            case "foo bar":
                var v75 uint8 = v25
                _ = v75
            case "\n\t":
                v76 := GenS6{F0: v4}
                _ = v76
            default:
                v77 := GenS8{F0: (v25 | (v25 & 1)), F1: 0, F2: 2}
                _ = v77
            }
        }
        v78 := v3[2147483647]
        for fi79 := 0; fi79 < 9; fi79++ {
            if false {
                if (!(!(v1 > v1))) {
                    fmt.Println(v1, fi79, v4)
                    v80 := GenS9{F0: 2}
                    _ = v80
                } else {
                    var v81 bool = true
                    if ((0.0 - 1e10) <= v9) {
                        switch (v8 | (v4 * v4)) {
                        case -9223372036854775808:
                            if v81 {
                                v82 := v3[329]
                                _ = v82
                            } else {
                                var v83 int64 = 0
                                _ = v83
                            }
                        default:
                            for fi84 := 0; fi84 < 20; fi84++ {
                                var v85 int32 = -1
                                v86 := GenS10{F0: 42, F1: 2147483647}
                                _ = fi84
                                _ = v85
                                _ = v86
                            }
                        }
                        v87 := GenS11{F0: 255}
                        v88 := GenS3{F0: 1, F1: 0, F2: 0}
                        _ = v87
                        _ = v88
                    }
                    v89 := GenS7{F0: (fi79 & (fi79 ^ 9223372036854775807)), F1: 468, F2: "foo bar"}
                    _ = v81
                    _ = v89
                }
                if (true || false) {
                    var v90 int = 4294967295
                    _ = v90
                }
            }
            if false {
                var v91 uint16 = 65535
                _ = v91
            } else {
                switch "a" {
                case "a":
                    v1 -= v1
                case "":
                    v25 = v25
                default:
                    switch "foo bar" {
                    case "a":
                        fmt.Println(v78)
                    }
                }
            }
            v3[42] = v78
            _ = fi79
        }
        _ = v78
    } else {
        var v92 uint16 = 911
        fmt.Println(v5, v4, v4)
        _ = v92
    }
    for fi93 := 0; fi93 < 12; fi93++ {
        v5 = v5
        v8 = (v4 - (fi93 + v8))
        _ = fi93
    }
    var v94 int8 = 0
    if ((((!(((v25 | v25) < 12) && (-2147483648 == int32(v8)))) && ((v4 / ((-275) | 1)) > 4294967295)) && true) || ((v9 / float32(float64(float32(0.1)))) >= v9)) {
        var v95 string = ("" + v5)
        var v96 float64 = -2.25
        _ = v95
        _ = v96
    }
    var v97 int64 = 770
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    fmt.Println("v94:", v94)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
    _ = v94
    _ = v97
}
