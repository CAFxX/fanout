package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int32; F1 int16 }
type GenS5 struct { F0 string; F1 int }
type GenS6 struct { F0 int32; F1 bool; F2 int32 }

func main() {
    var v1 bool = true
    var v2 string = "foo bar"
    var v3 string = (("" + v2) + "foo bar")
    v4 := make(map[string]int)
    v5 := make(map[int]int)
    var v6 float64 = float64(1.5)
    v7 := GenS1{F0: 42, F1: "", F2: ""}
    fmt.Println(v6, v2)
    if v1 {
        var v8 string = v2
        _ = v8
    } else {
        var v9 uint = 338
        _ = v9
    }
    if (4294967295 < 988) {
        var v10 int16 = -69
        switch 871 {
        case -2147483648:
            var v11 bool = (!v1)
            _ = v11
        case 32767:
            var v12 int16 = 42
            _ = v12
        case -1:
            switch "foo bar" {
            case "a":
                if (-128 > 127) {
                    var v13 uint32 = 42
                    v1 = false
                    for fi14 := 0; fi14 < 7; fi14++ {
                        for fi15 := 0; fi15 < 3; fi15++ {
                            fmt.Println(fi15, fi14, v13)
                            _ = fi15
                        }
                        _ = fi14
                    }
                    _ = v13
                } else {
                    if v1 {
                        if v1 {
                            v1 = (((float32(1e10) / (1.5 - 0.1)) < 3.14159) || v1)
                        } else {
                            v16 := v4[(v3 + v2)]
                            _ = v16
                        }
                    }
                }
            default:
                v17 := GenS1{F0: 0, F1: v2, F2: (("foo bar" + v3) + v2)}
                _ = v17
            }
        }
        _ = v10
    } else {
        v6 = v6
        v18 := []int{9223372036854775807, -908, 2147483647, 0}
        _ = v18
    }
    v6 = v6
    for fi19 := 0; fi19 < 4; fi19++ {
        switch v3 {
        case "\n\t":
            v20 := v5[-156]
            _ = v20
        case "foo bar":
            v21 := GenS1{F0: 42, F1: v2, F2: (v2 + ("\n\t" + v3))}
            _ = v21
        case "hello":
            fmt.Println(fi19)
        default:
            v22 := GenS1{F0: 42, F1: v3, F2: v3}
            _ = v22
        }
        switch v3 {
        case "":
            for fi23 := 0; fi23 < 5; fi23++ {
                v2 = v2
                v24 := GenS1{F0: 2, F1: "a", F2: (v2 + (v3 + (v3 + v2)))}
                v25 := GenS1{F0: 311, F1: ((((v3 + v2) + "a") + v3) + "hello"), F2: "a"}
                _ = fi23
                _ = v24
                _ = v25
            }
        case "foo bar":
            for fi26 := 0; fi26 < 8; fi26++ {
                var v27 int = (fi19 * fi19)
                v28 := GenS1{F0: 16, F1: v2, F2: ("a" + "hello")}
                _ = fi26
                _ = v27
                _ = v28
            }
        case "a":
            for fi29 := 0; fi29 < 12; fi29++ {
                v30 := []int{(fi29 + (fi29 | 32768)), fi19}
                var v31 int8 = -1
                if ((v31 << ((18446744073709551615) & 63)) == (^(v31))) {
                    if v1 {
                        if len(v30) >= 2 {
                            v30 = v30[1:len(v30)]
                        }
                    }
                    var v32 float64 = (0.0 / v6)
                    switch (("hello" + (("foo bar" + (v2 + ((v3 + ("hello" + v2)) + v3))) + "\n\t")) + "hello") {
                    case "hello":
                        v1 = v1
                    }
                    v33 := v4[(v3 + "a")]
                    _ = v32
                    _ = v33
                } else {
                    fmt.Println(v6, v1)
                    v2 = v3
                    var v34 uint32 = 979
                    _ = v34
                }
                _ = fi29
                _ = v30
                _ = v31
            }
        }
        v2 = "a"
        _ = fi19
    }
    v3 = v3
    v3 = (v2 + "hello")
    v35 := GenS1{F0: 49, F1: (v3 + (v2 + v3)), F2: ((v2 + "foo bar") + v3)}
    for fi36 := 0; fi36 < 18; fi36++ {
        for fi37 := 0; fi37 < 13; fi37++ {
            v5[32767] = fi37
            _ = fi37
        }
        fmt.Println(v2, v3, v6)
        fmt.Println(v2)
        _ = fi36
    }
    for fi38 := 0; fi38 < 7; fi38++ {
        v3 = v2
        v3 = v3
        _ = fi38
    }
    if false {
        v39 := GenS1{F0: 278, F1: (v2 + v2), F2: ("hello" + v2)}
        var v40 int = -9223372036854775808
        _ = v39
        _ = v40
    }
    switch (v2 + "") {
    case "":
        if v1 {
            if (2 >= 348) {
                var v41 uint32 = 747
                var v42 uint32 = (^(v41))
                if ((((3.14159 + 0.1) * ((float32(float64(0.0)) - (3.14159 - 0.0)) * (0.0 / 0.1))) / (0.0 + 1.5)) == float32(float64(0.0))) {
                    v43 := []int{72, 32767, -520}
                    if v1 {
                        var v44 int16 = 0
                        v45 := GenS2{F0: 1}
                        _ = v44
                        _ = v45
                    }
                    v41 -= v41
                    _ = v43
                }
                _ = v41
                _ = v42
            } else {
                v1 = v1
                if (15 != 42) {
                    if false {
                        if v1 {
                            for fi46 := 0; fi46 < 11; fi46++ {
                                if false {
                                    v6 = ((float64(float32(float64(3.14159))) * (float64(-2.25) + (v6 + 1e10))) / 1.5)
                                    if v1 {
                                        for fi47 := 0; fi47 < 12; fi47++ {
                                            if (true || v1) {
                                                var v48 int8 = -80
                                                switch ("a" + "a") {
                                                case "foo bar":
                                                    v1 = (2 != 2)
                                                case "hello":
                                                    var v49 uint16 = 513
                                                    _ = v49
                                                }
                                                fmt.Println(fi46)
                                                _ = v48
                                            }
                                            _ = fi47
                                        }
                                        v50 := []int{fi46}
                                        for fi51 := 0; fi51 < 19; fi51++ {
                                            if (v1 || (127 <= 42)) {
                                                if (396 >= 32767) {
                                                    fmt.Println(v1, fi46, v2)
                                                } else {
                                                    switch "a" {
                                                    case "foo bar":
                                                        fmt.Println(v2)
                                                    case "hello":
                                                        var v52 uint64 = 304
                                                        _ = v52
                                                    case "\n\t":
                                                        if (v1 && v1) {
                                                            var v53 int64 = 0
                                                            if v1 {
                                                                var v54 int = -2147483648
                                                                var v55 int32 = -491
                                                                _ = v54
                                                                _ = v55
                                                            } else {
                                                                v56 := GenS3{F0: 18446744073709551615}
                                                                _ = v56
                                                            }
                                                            _ = v53
                                                        } else {
                                                            if true {
                                                                if v1 {
                                                                    switch v3 {
                                                                    case "\n\t":
                                                                        switch (v3 + v3) {
                                                                        case "":
                                                                            v3 = "foo bar"
                                                                        }
                                                                    case "a":
                                                                        if true {
                                                                            var v57 int = (fi46 | fi51)
                                                                            _ = v57
                                                                        } else {
                                                                            if (!(!(681 != -32768))) {
                                                                                switch fi46 {
                                                                                case 128:
                                                                                    _ = len(v50)
                                                                                case -228:
                                                                                    if v1 {
                                                                                        fmt.Println(v3, v3, fi51)
                                                                                        v58 := v5[(fi46 & 4294967295)]
                                                                                        _ = v58
                                                                                    } else {
                                                                                        v59 := GenS3{F0: 527}
                                                                                        v6 -= v6
                                                                                        var v60 uint32 = 544
                                                                                        _ = len(v50)
                                                                                        switch v3 {
                                                                                        case "":
                                                                                            v61 := GenS2{F0: -1}
                                                                                            _ = v61
                                                                                        default:
                                                                                            var v62 int32 = 0
                                                                                            _ = v62
                                                                                        }
                                                                                        _ = v59
                                                                                        _ = v60
                                                                                    }
                                                                                case 0:
                                                                                    v63 := GenS1{F0: 314, F1: ("a" + (v3 + v3)), F2: (v2 + "")}
                                                                                    _ = v63
                                                                                }
                                                                                v64 := GenS2{F0: 42}
                                                                                _ = v64
                                                                            }
                                                                            var v65 float64 = float64(1e10)
                                                                            v2 = "\n\t"
                                                                            _ = v65
                                                                        }
                                                                    }
                                                                    v66 := v4[("\n\t" + "hello")]
                                                                    v67 := GenS2{F0: 2147483647}
                                                                    var v68 int32 = 0
                                                                    _ = v66
                                                                    _ = v67
                                                                    _ = v68
                                                                } else {
                                                                    v2 = v2
                                                                    fmt.Println(v1, fi46, v1)
                                                                    var v69 int64 = 0
                                                                    _ = v69
                                                                }
                                                            }
                                                            if (-1 <= -1) {
                                                                if (2 != 1) {
                                                                    v70 := GenS4{F0: -143, F1: -879}
                                                                    switch 2147483647 {
                                                                    case 121:
                                                                        switch v3 {
                                                                        case "\n\t":
                                                                            v71 := GenS5{F0: (v3 + "foo bar"), F1: 4294967295}
                                                                            _ = v71
                                                                        case "hello":
                                                                            var v72 uint32 = 787
                                                                            v5[(fi51 << ((289) & 63))] = fi51
                                                                            _ = v72
                                                                        case "a":
                                                                            v3 = v3
                                                                        default:
                                                                            v73 := GenS4{F0: 0, F1: -1}
                                                                            _ = v73
                                                                        }
                                                                    case 2147483647:
                                                                        var v74 uint16 = 221
                                                                        _ = v74
                                                                    case 4294967295:
                                                                        v6 -= (v6 * 1e10)
                                                                    }
                                                                    switch (fi46 >> ((4294967295) & 63)) {
                                                                    case 32768:
                                                                        v75 := GenS2{F0: -2147483648}
                                                                        _ = v75
                                                                    case 127:
                                                                        var v76 int = fi51
                                                                        _ = v76
                                                                    case 32767:
                                                                        var v77 int32 = -889
                                                                        _ = v77
                                                                    }
                                                                    v78 := GenS1{F0: 2, F1: "hello", F2: v2}
                                                                    var v79 float32 = 0.0
                                                                    _ = v70
                                                                    _ = v78
                                                                    _ = v79
                                                                } else {
                                                                    switch 32768 {
                                                                    case -744:
                                                                        var v80 bool = (341 != 2)
                                                                        _ = v80
                                                                    case -9223372036854775808:
                                                                        var v81 string = v2
                                                                        _ = v81
                                                                    case 4294967295:
                                                                        delete(v4, (v2 + v3))
                                                                    }
                                                                    v82 := GenS2{F0: 42}
                                                                    _ = v82
                                                                }
                                                            }
                                                            v4[("\n\t" + v2)] = fi51
                                                            v3 = (v2 + (v3 + "a"))
                                                        }
                                                    default:
                                                        v1 = v1
                                                    }
                                                    if (149 >= 1) {
                                                        switch 203 {
                                                        case 2147483647:
                                                            if (0 <= 817) {
                                                                switch fi51 {
                                                                case -1:
                                                                    if v1 {
                                                                        v3 = (v2 + "hello")
                                                                        switch (fi51 ^ fi46) {
                                                                        case 42:
                                                                            fmt.Println(v1, v3)
                                                                        default:
                                                                            var v83 uint8 = 1
                                                                            _ = v83
                                                                        }
                                                                    }
                                                                case 806:
                                                                    var v84 bool = (!(false || v1))
                                                                    _ = v84
                                                                default:
                                                                    fmt.Println(v2)
                                                                }
                                                                v50 = append(v50, (fi51 / (((fi51 >> ((4294967295) & 63))) | 1)))
                                                                if (-1 != -1) {
                                                                    var v85 string = v2
                                                                    _ = v85
                                                                }
                                                            }
                                                        case 841:
                                                            var v86 int16 = 1000
                                                            _ = v86
                                                        }
                                                        fmt.Println(v2, fi51)
                                                    } else {
                                                        var v87 int16 = -32768
                                                        var v88 uint16 = 255
                                                        _ = v87
                                                        _ = v88
                                                    }
                                                    v2 = (v3 + v3)
                                                }
                                            } else {
                                                delete(v4, v3)
                                            }
                                            var v89 uint32 = 65535
                                            var v90 uint16 = 310
                                            _ = fi51
                                            _ = v89
                                            _ = v90
                                        }
                                        _ = v50
                                    }
                                }
                                for fi91 := 0; fi91 < 16; fi91++ {
                                    fmt.Println(v2)
                                    _ = fi91
                                }
                                var v92 string = (v2 + (v2 + v2))
                                _ = fi46
                                _ = v92
                            }
                        }
                        for fi93 := 0; fi93 < 11; fi93++ {
                            fmt.Println(v2, fi93, v2)
                            v1 = (122 <= 194)
                            v2 = "a"
                            _ = fi93
                        }
                    } else {
                        for fi94 := 0; fi94 < 5; fi94++ {
                            v4[v2] = 1
                            _ = fi94
                        }
                        switch ((v3 + ((v2 + ("a" + "\n\t")) + "\n\t")) + "a") {
                        case "foo bar":
                            switch (v2 + "") {
                            case "foo bar":
                                v95 := GenS2{F0: 0}
                                _ = v95
                            default:
                                fmt.Println(v1, v2)
                            }
                        case "a":
                            var v96 int32 = 730
                            _ = v96
                        default:
                            for fi97 := 0; fi97 < 7; fi97++ {
                                var v98 int8 = -128
                                fmt.Println(fi97)
                                _ = fi97
                                _ = v98
                            }
                        }
                        v1 = (true && v1)
                    }
                } else {
                    v99 := GenS6{F0: 752, F1: false, F2: -467}
                    _ = v99
                }
            }
        }
    }
    v6 = v6
    v3 = v3
    for fi100 := 0; fi100 < 11; fi100++ {
        var v101 int32 = -773
        _ = fi100
        _ = v101
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
