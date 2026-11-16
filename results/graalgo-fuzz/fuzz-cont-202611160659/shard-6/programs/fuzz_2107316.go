package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 uint16; F1 bool; F2 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 uint16; F1 string; F2 uint }

func main() {
    var v1 bool = ((1 < 847) || ("" > "foo bar"))
    var v2 bool = v1
    var v3 float32 = 1e10
    v4 := make(map[string]uint32)
    var v5 int16 = -32768
    var v6 int8 = 48
    if v1 {
        if v2 {
            for fi7 := 0; fi7 < 19; fi7++ {
                v5 = v5
                if (v1 && (!v1)) {
                    var v8 int = 2147483647
                    for fi9 := 0; fi9 < 8; fi9++ {
                        var v10 bool = (42 != 338)
                        v10 = ((2 >= 2) || v1)
                        _ = fi9
                        _ = v10
                    }
                    _ = v8
                }
                _ = fi7
            }
            v5 += v5
        }
    } else {
        v1 = (!(!((2 == 0) && (v2 && v1))))
    }
    if (v2 && false) {
        v11 := GenS1{F0: 42, F1: (v5 % (((v5 >> ((18446744073709551615) & 63))) | 1))}
        switch -1 {
        case -2147483648:
            var v12 uint8 = 1
            _ = v12
        case -404:
            fmt.Println(v2, v6, v5)
        default:
            if true {
                for fi13 := 0; fi13 < 20; fi13++ {
                    switch "" {
                    case "foo bar":
                        if v2 {
                            var v14 int8 = v6
                            v2 = (v5 != (v5 * v5))
                            for fi15 := 0; fi15 < 11; fi15++ {
                                if (((-(fi15)) < fi13) && v2) {
                                    var v16 uint = 187
                                    var v17 uint8 = 26
                                    fmt.Println(fi15, v2)
                                    _ = v16
                                    _ = v17
                                }
                                var v18 int = (-(fi13))
                                v19 := GenS1{F0: (^(v6)), F1: v5}
                                v3 = v3
                                v14 += 1
                                v18 -= (fi13 ^ (fi15 & fi15))
                                _ = fi15
                                _ = v18
                                _ = v19
                            }
                            _ = v14
                        }
                    default:
                        if false {
                            fmt.Println(v1, v5, fi13)
                            var v20 uint8 = 0
                            _ = v20
                        } else {
                            switch "a" {
                            case "hello":
                                switch -1 {
                                case -1:
                                    var v21 int64 = 312
                                    _ = v21
                                }
                            case "\n\t":
                                var v22 int16 = v5
                                _ = v22
                            case "a":
                                var v23 uint32 = 1
                                _ = v23
                            }
                        }
                    }
                    for fi24 := 0; fi24 < 16; fi24++ {
                        var v25 int8 = v6
                        _ = fi24
                        _ = v25
                    }
                    _ = fi13
                }
                delete(v4, "foo bar")
                v26 := v4["hello"]
                _ = v26
            }
        }
        _ = v11
    } else {
        delete(v4, "foo bar")
        switch "hello" {
        case "":
            for fi27 := 0; fi27 < 16; fi27++ {
                var v28 int64 = -9223372036854775808
                if v1 {
                    fmt.Println(v28)
                    switch "hello" {
                    case "a":
                        var v29 bool = (!v1)
                        _ = v29
                    }
                } else {
                    v5 -= (v5 ^ 1)
                }
                _ = fi27
                _ = v28
            }
        }
        for fi30 := 0; fi30 < 17; fi30++ {
            var v31 uint16 = 32767
            delete(v4, "foo bar")
            if (4294967295 >= uint(v31)) {
                v32 := []int{}
                _ = v32
            }
            _ = fi30
            _ = v31
        }
    }
    var v33 string = "foo bar"
    var v34 int32 = -1
    _ = len(v4)
    for fi35 := 0; fi35 < 20; fi35++ {
        v36 := GenS1{F0: v6, F1: v5}
        _ = fi35
        _ = v36
    }
    for fi37 := 0; fi37 < 17; fi37++ {
        switch ("a" + "") {
        case "foo bar":
            v38 := []int{}
            _ = v38
        case "":
            switch "" {
            case "a":
                v34 += 2147483647
            case "foo bar":
                for fi39 := 0; fi39 < 13; fi39++ {
                    if v1 {
                        v2 = (4294967295 < 633)
                        fmt.Println(fi39, fi39)
                        v4[(("\n\t" + "hello") + "foo bar")] = 2147483648
                    } else {
                        v2 = (v1 && ((v3 - 1e10) < (float32((1.5 / (1e10 + (((((1.5 / 1.0) / 0.1) / 1.5) + ((1e10 + (0.0 / ((1.5 / 3.14159) - float64(((3.14159 + v3) * v3))))) + (3.14159 * (((float64(1e10) / (float64(v3) * 0.0)) * 0.0) / -2.25)))) + 1.5)))) + v3)))
                    }
                    _ = fi39
                }
            case "":
                var v40 int = -1
                _ = v40
            }
        case "a":
            if (((v34 == (v34 + v34)) && (v2 || (v2 || (!v1)))) && ((false && v1) || false)) {
                if false {
                    v2 = true
                    switch "" {
                    case "\n\t":
                        v41 := GenS1{F0: (v6 ^ (v6 * (-(v6)))), F1: v5}
                        _ = v41
                    default:
                        var v42 int32 = -500
                        _ = v42
                    }
                    fmt.Println(v33)
                } else {
                    for fi43 := 0; fi43 < 15; fi43++ {
                        fmt.Println(v1, v1)
                        switch 586 {
                        case 0:
                            _ = len(v4)
                        case -69:
                            if true {
                                var v44 string = "a"
                                _ = v44
                            }
                        case 4294967295:
                            v6 += -78
                        default:
                            fmt.Println(v34, fi37, v33)
                        }
                        v5 = 0
                        _ = fi43
                    }
                }
                v34 -= v34
                for fi45 := 0; fi45 < 8; fi45++ {
                    var v46 uint16 = 313
                    _ = fi45
                    _ = v46
                }
            } else {
                for fi47 := 0; fi47 < 1; fi47++ {
                    v34 = (v34 & (v34 / ((203) | 1)))
                    v48 := GenS1{F0: 106, F1: (v5 - v5)}
                    var v49 float32 = v3
                    _ = fi47
                    _ = v48
                    _ = v49
                }
            }
        default:
            v1 = false
        }
        _ = fi37
    }
    fmt.Println(v3, v2, v2)
    var v50 uint8 = 42
    var v51 int64 = 1
    for fi52 := 0; fi52 < 19; fi52++ {
        if false {
            if (42 < 2) {
                var v53 bool = v1
                var v54 int16 = (v5 % ((42) | 1))
                _ = v53
                _ = v54
            } else {
                if v2 {
                    switch fi52 {
                    case 0:
                        v33 = v33
                    case 32768:
                        v2 = ((v51 % (((v51 * 821)) | 1)) < int64(fi52))
                    case -9223372036854775808:
                        for fi55 := 0; fi55 < 4; fi55++ {
                            v56 := GenS1{F0: v6, F1: v5}
                            v6 = -107
                            fmt.Println(v5, v51, v34)
                            _ = fi55
                            _ = v56
                        }
                    default:
                        for fi57 := 0; fi57 < 10; fi57++ {
                            if v2 {
                                v6 = (v6 & v6)
                                var v58 int32 = v34
                                switch v33 {
                                case "":
                                    switch ((v33 + (v33 + v33)) + (v33 + "")) {
                                    case "a":
                                        v6 = 1
                                    case "":
                                        switch "\n\t" {
                                        case "\n\t":
                                            var v59 int16 = 0
                                            _ = v59
                                        default:
                                            var v60 uint8 = 2
                                            _ = v60
                                        }
                                        v61 := GenS1{F0: v6, F1: v5}
                                        _ = v61
                                    default:
                                        switch fi52 {
                                        case -2147483648:
                                            var v62 bool = v2
                                            _ = v62
                                        }
                                    }
                                    v1 = true
                                    if (v33 < ("a" + (v33 + "hello"))) {
                                        if (2 == 4294967295) {
                                            var v63 float64 = 1.5
                                            v64 := []int{(^(fi52)), (fi52 / (((fi52 >> ((698) & 63))) | 1)), 32768, 127}
                                            _ = len(v4)
                                            _ = v63
                                            _ = v64
                                        } else {
                                            v65 := []int{}
                                            switch (("hello" + ((v33 + v33) + (("" + v33) + v33))) + "foo bar") {
                                            case "a":
                                                v66 := GenS2{F0: 428, F1: v2, F2: true}
                                                _ = v66
                                            default:
                                                var v67 bool = v1
                                                _ = v67
                                            }
                                            v68 := GenS3{F0: 929}
                                            var v69 uint8 = (v50 >> ((65535) & 63))
                                            v2 = v2
                                            fmt.Println(v6, fi57, v58)
                                            _ = v65
                                            _ = v68
                                            _ = v69
                                        }
                                        if false {
                                            delete(v4, "")
                                            fmt.Println(v50, v58)
                                        }
                                        var v70 string = "a"
                                        fmt.Println(v34, v70)
                                        _ = v70
                                    }
                                    if (v6 != v6) {
                                        v71 := GenS1{F0: (^(v6)), F1: (-(v5))}
                                        var v72 string = (v33 + "")
                                        switch v33 {
                                        case "a":
                                            var v73 int32 = (-(v34))
                                            _ = v73
                                        case "":
                                            if (1000 >= v5) {
                                                var v74 string = ("hello" + ((v33 + "a") + "foo bar"))
                                                v75 := GenS1{F0: 127, F1: v5}
                                                v76 := GenS4{F0: v5, F1: -243}
                                                _ = v74
                                                _ = v75
                                                _ = v76
                                            }
                                        default:
                                            v5 = v5
                                        }
                                        _ = v71
                                        _ = v72
                                    } else {
                                        v5 += (v5 >> ((825) & 63))
                                        switch fi52 {
                                        case 127:
                                            _ = len(v4)
                                        case 32767:
                                            v5 += (v5 ^ v5)
                                        case -776:
                                            if true {
                                                var v77 float64 = -2.25
                                                switch fi52 {
                                                case 4294967295:
                                                    delete(v4, (v33 + "hello"))
                                                case -9223372036854775808:
                                                    var v78 int64 = v51
                                                    _ = v78
                                                case 536:
                                                    _ = len(v4)
                                                default:
                                                    v77 -= (v77 + v77)
                                                }
                                                v34 += (v58 * (v58 * v34))
                                                _ = v77
                                            }
                                        default:
                                            switch v33 {
                                            case "a":
                                                if (!((v6 ^ v6) >= (v6 >> ((4294967295) & 63)))) {
                                                    v79 := []int{}
                                                    _ = v79
                                                } else {
                                                    v80 := v4[(v33 + (v33 + (v33 + v33)))]
                                                    v81 := []int{}
                                                    v82 := GenS1{F0: v6, F1: v5}
                                                    _ = v80
                                                    _ = v81
                                                    _ = v82
                                                }
                                            case "hello":
                                                v6 += (v6 + 42)
                                            }
                                            var v83 uint8 = v50
                                            var v84 uint8 = v50
                                            _ = v83
                                            _ = v84
                                        }
                                    }
                                case "a":
                                    var v85 int64 = v51
                                    _ = v85
                                case "foo bar":
                                    v50 -= (v50 | v50)
                                }
                                _ = v58
                            }
                            if false {
                                var v86 uint16 = 2
                                fmt.Println(v1, fi57)
                                _ = v86
                            }
                            _ = fi57
                        }
                    }
                }
                fmt.Println(v1, v5)
            }
        } else {
            for fi87 := 0; fi87 < 1; fi87++ {
                v2 = true
                _ = fi87
            }
        }
        v88 := []int{}
        _ = fi52
        _ = v88
    }
    v34 = -2147483648
    var v89 uint64 = 18446744073709551615
    v90 := []int{533}
    var v91 int32 = int32(v6)
    if ((v5 >> ((402) & 63)) >= 477) {
        var v92 int32 = -705
        fmt.Println(v89, v1)
        _ = v92
    } else {
        for fi93 := 0; fi93 < 9; fi93++ {
            if true {
                if (65535 > uint16(v50)) {
                    for fi94 := 0; fi94 < 18; fi94++ {
                        v51 = (v51 / ((v51) | 1))
                        fmt.Println(v1)
                        var v95 uint32 = 2147483648
                        _ = fi94
                        _ = v95
                    }
                    fmt.Println(v3)
                }
                for fi96 := 0; fi96 < 19; fi96++ {
                    v97 := GenS4{F0: 1, F1: -1}
                    v98 := GenS5{F0: 32767, F1: ((v33 + v33) + v33), F2: 255}
                    _ = fi96
                    _ = v97
                    _ = v98
                }
            }
            _ = fi93
        }
        v90 = append(v90, -1)
        if len(v90) >= 2 {
            v90 = v90[1:len(v90)]
        }
    }
    delete(v4, v33)
    fmt.Println(v34)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    fmt.Println("v89:", v89)
    fmt.Println("slicelen v90:", len(v90), v90)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v6
    _ = v89
    _ = v90
    _ = v91
}
