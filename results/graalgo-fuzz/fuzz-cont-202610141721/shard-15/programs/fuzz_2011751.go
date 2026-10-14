package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32; F1 uint16; F2 int }
type GenS5 struct { F0 int64; F1 uint; F2 uint8 }
type GenS6 struct { F0 uint16; F1 int64; F2 uint32 }
type GenS7 struct { F0 uint64; F1 uint32; F2 uint32 }
type GenS8 struct { F0 int16; F1 bool }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int32 }

func main() {
    var v1 int = -9223372036854775808
    var v2 string = (("" + "foo bar") + "foo bar")
    v3 := []int{(v1 ^ v1), (v1 - (v1 | v1)), 6, (v1 + v1)}
    v4 := GenS1{F0: 42, F1: 65535}
    var v5 uint8 = 159
    if len(v3) > 0 {
        v3[0] = v1
    }
    switch (v2 + v2) {
    case "foo bar":
        v6 := GenS1{F0: (v5 * v5), F1: 977}
        _ = v6
    case "hello":
        v5 -= (v5 | 42)
    case "\n\t":
        switch "a" {
        case "\n\t":
            if (0 > -1) {
                v5 += 2
            } else {
                if (18446744073709551615 != 18446744073709551615) {
                    v7 := GenS2{F0: 2}
                    _ = v7
                }
                for fi8 := 0; fi8 < 20; fi8++ {
                    v9 := GenS1{F0: v5, F1: 0}
                    _ = fi8
                    _ = v9
                }
            }
        case "foo bar":
            for fi10 := 0; fi10 < 1; fi10++ {
                if (!((((("" + v2) + ("hello" + v2)) + v2) + ((v2 + ("foo bar" + v2)) + "")) > v2)) {
                    v11 := GenS2{F0: 790}
                    _ = v11
                } else {
                    switch ("a" + v2) {
                    case "":
                        fmt.Println(v1)
                    }
                    switch fi10 {
                    case 127:
                        fmt.Println(fi10)
                    case 605:
                        if true {
                            fmt.Println(v2, v2, v2)
                            v12 := GenS3{F0: false}
                            for fi13 := 0; fi13 < 14; fi13++ {
                                var v14 int8 = 74
                                _ = fi13
                                _ = v14
                            }
                            _ = v12
                        } else {
                            switch (^(v1)) {
                            case 32768:
                                v5 -= v5
                            case 1:
                                if (true || (!(false && (702 > 0)))) {
                                    var v15 uint32 = 2147483647
                                    for fi16 := 0; fi16 < 14; fi16++ {
                                        if len(v3) > 0 {
                                            v3[0] = v1
                                        }
                                        v1 += fi16
                                        fmt.Println(fi16, v2)
                                        _ = fi16
                                    }
                                    var v17 uint = 120
                                    _ = v15
                                    _ = v17
                                } else {
                                    var v18 uint = 136
                                    _ = v18
                                }
                            default:
                                if true {
                                    v2 = "a"
                                    v1 -= (^(v1))
                                    for fi19 := 0; fi19 < 18; fi19++ {
                                        if true {
                                            if false {
                                                var v20 uint32 = 4294967295
                                                v21 := GenS4{F0: v20, F1: 32767, F2: (-(fi19))}
                                                if true {
                                                    var v22 int = (fi10 ^ (-(v1)))
                                                    if len(v3) > 0 {
                                                        v3[0] = 2147483647
                                                    }
                                                    var v23 float64 = (0.0 / 1.5)
                                                    var v24 int64 = -9223372036854775808
                                                    _ = v22
                                                    _ = v23
                                                    _ = v24
                                                }
                                                _ = v20
                                                _ = v21
                                            } else {
                                                v2 = (("a" + (v2 + ((v2 + v2) + ""))) + "a")
                                                if true {
                                                    if (((-2.25 * ((-2.25 - 1.5) / 1.0)) * (3.14159 / ((1e10 / float64(1e10)) - -2.25))) == (-2.25 + (1.5 * 1e10))) {
                                                        switch (v2 + v2) {
                                                        case "a":
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        case "hello":
                                                            var v25 int8 = -1
                                                            _ = v25
                                                        default:
                                                            var v26 int8 = 0
                                                            _ = v26
                                                        }
                                                    }
                                                    v5 = v5
                                                }
                                                switch (fi19 ^ (fi10 / (((v1 + 9223372036854775807)) | 1))) {
                                                case 32767:
                                                    switch fi19 {
                                                    case 1:
                                                        fmt.Println(v2)
                                                    case 32768:
                                                        v27 := GenS4{F0: 773, F1: 32767, F2: fi19}
                                                        _ = v27
                                                    case -2147483648:
                                                        v3 = append(v3, fi10)
                                                    default:
                                                        v3 = append(v3, (-(fi10)))
                                                    }
                                                case 0:
                                                    var v28 string = v2
                                                    _ = v28
                                                }
                                                var v29 int32 = 2147483647
                                                _ = v29
                                            }
                                        } else {
                                            if (!(0 >= 4294967295)) {
                                                var v30 int8 = 0
                                                var v31 uint64 = 2
                                                v32 := GenS3{F0: (-2.25 < 0.0)}
                                                v33 := GenS3{F0: false}
                                                v2 = v2
                                                var v34 int8 = 42
                                                var v35 uint64 = v31
                                                var v36 uint32 = 4294967295
                                                _ = v30
                                                _ = v31
                                                _ = v32
                                                _ = v33
                                                _ = v34
                                                _ = v35
                                                _ = v36
                                            }
                                            v1 = 85
                                            v37 := GenS4{F0: 4294967295, F1: 65535, F2: v1}
                                            fmt.Println(fi10, v2, v1)
                                            var v38 uint64 = 18446744073709551615
                                            _ = v37
                                            _ = v38
                                        }
                                        fmt.Println(v1)
                                        _ = fi19
                                    }
                                } else {
                                    if (11 == 127) {
                                        v39 := GenS2{F0: 32767}
                                        for fi40 := 0; fi40 < 13; fi40++ {
                                            v3 = append(v3, (v1 * v1))
                                            v1 = fi10
                                            _ = fi40
                                        }
                                        for fi41 := 0; fi41 < 1; fi41++ {
                                            v3 = append(v3, fi41)
                                            v2 = v2
                                            v1 = 9223372036854775807
                                            _ = fi41
                                        }
                                        _ = v39
                                    }
                                }
                            }
                            if (v2 >= (v2 + ("" + ((((v2 + v2) + v2) + ("hello" + ("a" + v2))) + "")))) {
                                var v42 int8 = 0
                                _ = v42
                            } else {
                                v2 = "\n\t"
                            }
                            if (1 > 840) {
                                var v43 uint = 2
                                var v44 float32 = 3.14159
                                if false {
                                    v5 = 0
                                    v45 := GenS5{F0: -171, F1: v43, F2: v5}
                                    v44 = v44
                                    _ = v45
                                }
                                _ = v43
                                _ = v44
                            } else {
                                if (0 >= 1) {
                                    if (!false) {
                                        v1 = 295
                                        v46 := GenS6{F0: 623, F1: -1, F2: 2}
                                        _ = v46
                                    } else {
                                        var v47 float32 = 1e10
                                        fmt.Println(v5, v1)
                                        var v48 uint8 = v5
                                        _ = v47
                                        _ = v48
                                    }
                                    for fi49 := 0; fi49 < 14; fi49++ {
                                        var v50 uint8 = 127
                                        var v51 uint16 = 2
                                        v2 = v2
                                        var v52 uint32 = 589
                                        _ = fi49
                                        _ = v50
                                        _ = v51
                                        _ = v52
                                    }
                                    v2 = v2
                                }
                                if (913 > 42) {
                                    var v53 uint32 = 150
                                    fmt.Println(v5, fi10)
                                    v1 -= fi10
                                    _ = v53
                                }
                                var v54 int8 = 100
                                _ = v54
                            }
                        }
                    }
                    if true {
                        fmt.Println(fi10, fi10, v2)
                    }
                }
                v55 := GenS7{F0: 2, F1: 2147483648, F2: 2147483647}
                _ = fi10
                _ = v55
            }
        case "":
            v56 := GenS1{F0: v5, F1: 0}
            _ = v56
        }
    }
    fmt.Println(v5, v2)
    v57 := GenS4{F0: 993, F1: 255, F2: v1}
    for fi58 := 0; fi58 < 6; fi58++ {
        v59 := GenS8{F0: 32767, F1: (false && (18446744073709551615 == 839))}
        if false {
            fmt.Println(v2)
            switch fi58 {
            case 0:
                v60 := GenS9{F0: 2}
                _ = v60
            case 306:
                fmt.Println(v1)
            case 128:
                v2 = v2
            default:
                v5 = 68
            }
            for fi61 := 0; fi61 < 5; fi61++ {
                v5 = (v5 | v5)
                fmt.Println(v1, v5, fi58)
                var v62 int = (-(v1))
                _ = fi61
                _ = v62
            }
        } else {
            v1 = -2147483648
        }
        v2 = "\n\t"
        _ = fi58
        _ = v59
    }
    v63 := GenS5{F0: 228, F1: 4294967295, F2: (v5 << ((954) & 63))}
    var v64 int64 = 9223372036854775807
    var v65 float32 = 0.0
    var v66 int8 = -1
    for fi67 := 0; fi67 < 15; fi67++ {
        switch ("" + (v2 + v2)) {
        case "foo bar":
            v5 = v5
        }
        v68 := GenS1{F0: 128, F1: 0}
        var v69 uint16 = 89
        _ = fi67
        _ = v68
        _ = v69
    }
    for fi70 := 0; fi70 < 20; fi70++ {
        v1 = fi70
        _ = fi70
    }
    for fi71 := 0; fi71 < 20; fi71++ {
        fmt.Println(v2)
        _ = fi71
    }
    v72 := GenS9{F0: 1}
    v73 := GenS10{F0: 42}
    var v74 uint8 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v63
    _ = v64
    _ = v65
    _ = v66
    _ = v72
    _ = v73
    _ = v74
}
