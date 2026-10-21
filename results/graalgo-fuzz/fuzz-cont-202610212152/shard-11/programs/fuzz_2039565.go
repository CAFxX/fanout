package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int; F1 uint32 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int16; F1 int32; F2 uint }
type GenS9 struct { F0 uint64 }

func main() {
    v1 := []string{"foo bar", "\n\t"}
    v2 := []uint32{465}
    v3 := make(map[int]string)
    v4 := make(map[uint32]string)
    switch 9223372036854775807 {
    case 32767:
        switch "" {
        case "a":
            var v5 int16 = 42
            _ = v5
        }
    default:
        var v6 uint16 = 1
        _ = v6
    }
    if (-322 == 998) {
        var v7 uint8 = 58
        v7 -= v7
        _ = v7
    }
    var v8 int64 = 9223372036854775807
    v8 = (v8 + v8)
    v9 := GenS1{F0: -389, F1: 0}
    fmt.Println(v8)
    var v10 int32 = -1
    for fi11 := 0; fi11 < 2; fi11++ {
        for fi12 := 0; fi12 < 20; fi12++ {
            if (-25 < 64) {
                fmt.Println(v10, v8, fi11)
                v10 += (-(v10))
                v10 = (v10 | -1)
            } else {
                var v13 uint16 = 42
                var v14 bool = (false || (110 >= -1))
                var v15 uint64 = 878
                _ = v13
                _ = v14
                _ = v15
            }
            var v16 int = (fi11 | 4294967295)
            if false {
                var v17 uint64 = 42
                _ = v17
            } else {
                v10 = 408
                if (0 > 18446744073709551615) {
                    v8 -= (v8 % ((v8) | 1))
                    var v18 int16 = 0
                    _ = v18
                } else {
                    switch ("foo bar" + "a") {
                    case "a":
                        v19 := v4[244]
                        _ = v19
                    }
                    if (("foo bar" + "foo bar") < ("hello" + "foo bar")) {
                        switch (-(v16)) {
                        case 4294967295:
                            fmt.Println(v16, v8)
                        case 817:
                            v20 := GenS2{F0: 42, F1: true}
                            _ = v20
                        case 42:
                            switch "\n\t" {
                            case "foo bar":
                                switch "\n\t" {
                                case "":
                                    var v21 float32 = 1e10
                                    _ = v21
                                case "\n\t":
                                    switch -749 {
                                    case 1:
                                        fmt.Println(v16)
                                    }
                                case "hello":
                                    var v22 float64 = 0.0
                                    _ = v22
                                }
                            case "a":
                                if false {
                                    switch "a" {
                                    case "\n\t":
                                        v23 := GenS3{F0: "foo bar"}
                                        _ = v23
                                    case "a":
                                        v3[117] = "foo bar"
                                    case "foo bar":
                                        var v24 int32 = (v10 / (((^(v10))) | 1))
                                        _ = v24
                                    }
                                    v10 -= int32(v8)
                                    v25 := GenS1{F0: 1000, F1: 2147483648}
                                    _ = len(v1)
                                    _ = v25
                                } else {
                                    v16 = (fi11 >> ((42) & 63))
                                }
                            }
                            var v26 uint32 = 4294967295
                            if (44 > 128) {
                                var v27 uint64 = 0
                                fmt.Println(fi12, fi11, fi11)
                                _ = v27
                            }
                            _ = v26
                        }
                    } else {
                        var v28 int = -9223372036854775808
                        var v29 float64 = 0.1
                        var v30 float32 = 0.1
                        fmt.Println(fi12)
                        _ = v28
                        _ = v29
                        _ = v30
                    }
                }
                v8 += 1
                _ = len(v4)
            }
            v31 := GenS2{F0: 42, F1: (v16 <= fi12)}
            _ = fi12
            _ = v16
            _ = v31
        }
        _ = fi11
    }
    v8 = v8
    if false {
        v32 := GenS1{F0: 1, F1: 2}
        v33 := GenS3{F0: ""}
        _ = v32
        _ = v33
    } else {
        if true {
            if ((!(42 >= 0)) && (18446744073709551615 <= 478)) {
                for fi34 := 0; fi34 < 12; fi34++ {
                    for fi35 := 0; fi35 < 1; fi35++ {
                        v36 := GenS4{F0: 2147483648}
                        fmt.Println(fi34)
                        v1 = append(v1, "\n\t")
                        if (false && false) {
                            v10 = 42
                            var v37 string = "foo bar"
                            _ = v37
                        } else {
                            v38 := GenS5{F0: -23}
                            if (!true) {
                                v39 := GenS4{F0: 4294967295}
                                if (fi34 < fi35) {
                                    v10 -= 0
                                } else {
                                    if (1 < 106) {
                                        fmt.Println(fi34, v8)
                                        switch "hello" {
                                        case "hello":
                                            v40 := GenS6{F0: 127, F1: 447}
                                            var v41 int32 = (-(v10))
                                            _ = v40
                                            _ = v41
                                        default:
                                            v8 = v8
                                        }
                                        if (0 != 4294967295) {
                                            v1 = append(v1, (("\n\t" + "\n\t") + "hello"))
                                            switch "foo bar" {
                                            case "a":
                                                v10 = (-(v10))
                                                if (1 >= 428) {
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        var v42 int32 = 1
                                                        _ = v42
                                                    case "hello":
                                                        switch fi35 {
                                                        case -9223372036854775808:
                                                            v8 -= v8
                                                        case 32768:
                                                            v2 = append(v2, 153)
                                                            v10 = v10
                                                            var v43 uint16 = 255
                                                            _ = v43
                                                        case 42:
                                                            v8 += v8
                                                        default:
                                                            _ = len(v4)
                                                        }
                                                    default:
                                                        var v44 bool = (-128 > -128)
                                                        _ = v44
                                                    }
                                                    v10 += (v10 >> ((351) & 63))
                                                } else {
                                                    v10 -= (v10 - v10)
                                                }
                                                switch 32768 {
                                                case 120:
                                                    v8 += 42
                                                case 2147483647:
                                                    if true {
                                                        var v45 uint32 = 4294967295
                                                        _ = len(v4)
                                                        fmt.Println(fi35, v45)
                                                        fmt.Println(fi34)
                                                        v46 := v4[v45]
                                                        _ = v45
                                                        _ = v46
                                                    }
                                                case -685:
                                                    var v47 int16 = 32767
                                                    _ = v47
                                                }
                                                fmt.Println(fi35, v10)
                                            case "hello":
                                                var v48 string = ("a" + "")
                                                _ = v48
                                            default:
                                                switch ("a" + "a") {
                                                case "\n\t":
                                                    if (!false) {
                                                        fmt.Println(v8, v10)
                                                    }
                                                }
                                            }
                                            fmt.Println(v8, fi35)
                                            var v49 float32 = ((-2.25 / (0.1 - 1e10)) / 1e10)
                                            var v50 string = ""
                                            _ = v49
                                            _ = v50
                                        }
                                    }
                                    v51 := GenS4{F0: 42}
                                    if (963 > 0) {
                                        fmt.Println(fi34, fi35)
                                    } else {
                                        v10 += (v10 << ((4294967295) & 63))
                                        if (false || (!false)) {
                                            switch "\n\t" {
                                            case "hello":
                                                v8 += 358
                                            case "foo bar":
                                                v10 -= (v10 * v10)
                                            case "\n\t":
                                                v8 = 0
                                            }
                                            switch "hello" {
                                            case "":
                                                fmt.Println(fi34, v10)
                                            }
                                        } else {
                                            var v52 float32 = 0.0
                                            var v53 uint = 39
                                            var v54 float32 = float32(1.5)
                                            _ = v52
                                            _ = v53
                                            _ = v54
                                        }
                                        if (v8 < 42) {
                                            if (0 <= 764) {
                                                switch (((("a" + "a") + ("a" + "\n\t")) + (("a" + "a") + "hello")) + ("\n\t" + "")) {
                                                case "\n\t":
                                                    v55 := GenS4{F0: 2147483648}
                                                    _ = v55
                                                case "a":
                                                    var v56 int16 = -904
                                                    var v57 int = (fi35 >> ((42) & 63))
                                                    v58 := GenS4{F0: 1}
                                                    v59 := GenS7{F0: 42}
                                                    _ = v56
                                                    _ = v57
                                                    _ = v58
                                                    _ = v59
                                                case "":
                                                    var v60 uint64 = 42
                                                    _ = v60
                                                default:
                                                    v4[756] = (((("\n\t" + ("a" + "\n\t")) + "hello") + "") + (("hello" + "a") + ("hello" + "a")))
                                                }
                                            }
                                            v61 := GenS7{F0: 232}
                                            _ = v61
                                        } else {
                                            v8 = 716
                                            v62 := v3[1]
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            var v63 uint64 = 1
                                            _ = v62
                                            _ = v63
                                        }
                                    }
                                    _ = v51
                                }
                                _ = v39
                            } else {
                                if (1 < 0) {
                                    var v64 uint32 = 42
                                    if false {
                                        v8 -= (v8 & -179)
                                    } else {
                                        var v65 int16 = 42
                                        var v66 uint32 = v64
                                        var v67 float64 = 1.5
                                        _ = v65
                                        _ = v66
                                        _ = v67
                                    }
                                    switch (fi35 | -9223372036854775808) {
                                    case 711:
                                        v3[4294967295] = "a"
                                    case 128:
                                        v68 := GenS6{F0: 9223372036854775807, F1: v64}
                                        _ = v68
                                    case 137:
                                        if (-9223372036854775808 <= (-(v8))) {
                                            switch fi35 {
                                            case 0:
                                                var v69 string = "foo bar"
                                                _ = v69
                                            case 9223372036854775807:
                                                var v70 uint64 = 2
                                                _ = v70
                                            }
                                            switch (^(fi35)) {
                                            case 1:
                                                var v71 float64 = float64(((0.1 - (((0.0 / 1.0) - float32((0.0 - 0.1))) + float32(((((1e10 - 3.14159) * -2.25) / 1.5) * (0.0 * 0.1))))) - 3.14159))
                                                _ = v71
                                            default:
                                                var v72 uint16 = 586
                                                _ = v72
                                            }
                                        } else {
                                            var v73 int = 128
                                            fmt.Println(v73, v64, v64)
                                            _ = v73
                                        }
                                    default:
                                        v8 += (v8 ^ -835)
                                    }
                                    _ = v64
                                } else {
                                    fmt.Println(fi35, v10)
                                    v8 -= -1
                                    v74 := GenS8{F0: 1, F1: 42, F2: 1}
                                    _ = v74
                                }
                                switch ("foo bar" + "\n\t") {
                                case "\n\t":
                                    var v75 int16 = -32768
                                    _ = v75
                                default:
                                    var v76 int16 = -557
                                    _ = v76
                                }
                                switch (fi34 - 9223372036854775807) {
                                case 9223372036854775807:
                                    if false {
                                        switch ("hello" + "a") {
                                        case "a":
                                            var v77 float32 = float32(3.14159)
                                            _ = v77
                                        case "foo bar":
                                            var v78 int32 = v10
                                            _ = v78
                                        }
                                    }
                                }
                            }
                            fmt.Println(fi35, v8, fi34)
                            _ = v38
                        }
                        var v79 uint16 = 26
                        _ = fi35
                        _ = v36
                        _ = v79
                    }
                    switch (fi34 ^ -2147483648) {
                    case 283:
                        fmt.Println(v8, fi34, fi34)
                    }
                    v10 = v10
                    _ = fi34
                }
                for fi80 := 0; fi80 < 10; fi80++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if len(v1) > 0 {
                        v1[0] = ""
                    }
                    if false {
                        v2 = append(v2, 2147483647)
                        if ((-2.25 / -2.25) == 1.5) {
                            v81 := GenS9{F0: 42}
                            v8 = v8
                            _ = v81
                        } else {
                            delete(v4, 4294967295)
                        }
                    }
                    _ = fi80
                }
            }
            var v82 uint8 = 185
            v8 = v8
            _ = v82
        }
    }
    var v83 bool = true
    fmt.Println(v83, v10)
    v83 = v83
    if false {
        v84 := GenS5{F0: 127}
        v8 = 9223372036854775807
        _ = v84
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v83
    _ = v9
}
