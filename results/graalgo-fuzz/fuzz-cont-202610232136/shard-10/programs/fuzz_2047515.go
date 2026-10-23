package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint8 }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int; F1 uint8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int32; F1 uint8; F2 uint8 }
type GenS7 struct { F0 bool; F1 int }
type GenS8 struct { F0 int64; F1 int32; F2 int16 }

func main() {
    var v1 int = 0
    var v2 string = ("a" + "a")
    var v3 uint32 = 4294967295
    var v4 int8 = 111
    v5 := make(map[string]int)
    var v6 string = ((v2 + v2) + v2)
    v2 = (v2 + v2)
    v1 = (-(v1))
    v3 = v3
    for fi7 := 0; fi7 < 5; fi7++ {
        fmt.Println(v4, v6)
        _ = fi7
    }
    v8 := GenS1{F0: true, F1: 1, F2: 127}
    var v9 float32 = 3.14159
    v2 = ((((v2 + v6) + v2) + ((((("hello" + v2) + (v6 + v2)) + v6) + (((v2 + "a") + v2) + "")) + v6)) + ((v2 + "hello") + (v2 + ("foo bar" + v2))))
    v10 := GenS1{F0: true, F1: 4294967295, F2: 96}
    v11 := []int{4294967295, v1, (v1 | 912), -755}
    var v12 float64 = -2.25
    switch "\n\t" {
    case "foo bar":
        if (!false) {
            v2 = v6
            if (-195 != -32768) {
                v3 += v3
                var v13 string = ""
                v2 = "\n\t"
                _ = v13
            }
            for fi14 := 0; fi14 < 2; fi14++ {
                var v15 int16 = -1
                _ = fi14
                _ = v15
            }
        }
    default:
        v16 := GenS2{F0: true, F1: 242}
        _ = v16
    }
    var v17 uint64 = 18446744073709551615
    if len(v11) > 0 {
        v11[0] = (v1 ^ (v1 + (-(v1))))
    }
    v3 -= v3
    if true {
        v18 := GenS2{F0: (false && false), F1: (v17 - (v17 << ((18446744073709551615) & 63)))}
        for fi19 := 0; fi19 < 17; fi19++ {
            switch 374 {
            case -13:
                if true {
                    v20 := GenS1{F0: false, F1: 917, F2: 127}
                    switch v2 {
                    case "\n\t":
                        v11 = append(v11, (v1 + (v1 << ((65535) & 63))))
                    case "a":
                        fmt.Println(v9, v3, v9)
                    case "":
                        v21 := GenS2{F0: false, F1: v17}
                        _ = v21
                    default:
                        if false {
                            for fi22 := 0; fi22 < 19; fi22++ {
                                fmt.Println(v1)
                                v4 -= (^(v4))
                                var v23 int = (v1 + (v1 + (-(fi19))))
                                var v24 bool = true
                                switch (v6 + v2) {
                                case "hello":
                                    if v24 {
                                        _ = len(v11)
                                        if v24 {
                                            fmt.Println(v3, v4, v24)
                                            if (v3 == (v3 * (v3 | (v3 << ((2) & 63))))) {
                                                _ = len(v5)
                                                v25 := GenS2{F0: (!v24), F1: 42}
                                                _ = v25
                                            }
                                        } else {
                                            v24 = v24
                                            switch "\n\t" {
                                            case "\n\t":
                                                v9 -= v9
                                            default:
                                                switch -2147483648 {
                                                case -84:
                                                    var v26 uint16 = 65535
                                                    _ = v26
                                                default:
                                                    var v27 int = 1
                                                    var v28 int8 = 42
                                                    v4 = v4
                                                    var v29 bool = true
                                                    _ = v27
                                                    _ = v28
                                                    _ = v29
                                                }
                                            }
                                        }
                                        switch (^(v1)) {
                                        case 128:
                                            var v30 uint64 = v17
                                            _ = v30
                                        case 32767:
                                            if (v12 >= float64((v9 * v9))) {
                                                var v31 uint16 = 32767
                                                _ = v31
                                            } else {
                                                _ = len(v5)
                                            }
                                        }
                                    }
                                case "\n\t":
                                    if len(v11) > 0 {
                                        v11[0] = (fi22 % (((fi19 & v1)) | 1))
                                    }
                                case "a":
                                    v32 := GenS2{F0: (true || v24), F1: v17}
                                    _ = v32
                                }
                                _ = fi22
                                _ = v23
                                _ = v24
                            }
                        } else {
                            switch "a" {
                            case "hello":
                                if (-1 <= -32768) {
                                    for fi33 := 0; fi33 < 15; fi33++ {
                                        switch "" {
                                        case "foo bar":
                                            v6 = v6
                                        case "":
                                            var v34 float64 = 0.0
                                            fmt.Println(v1, fi19)
                                            fmt.Println(fi33)
                                            _ = v34
                                        default:
                                            v35 := GenS1{F0: false, F1: v17, F2: 0}
                                            _ = v35
                                        }
                                        fmt.Println(v3, v1)
                                        var v36 float32 = v9
                                        _ = fi33
                                        _ = v36
                                    }
                                    v12 = v12
                                } else {
                                    if len(v11) >= 2 {
                                        v11 = v11[1:len(v11)]
                                    }
                                    if len(v11) >= 2 {
                                        v11 = v11[1:len(v11)]
                                    }
                                    var v37 int16 = 1000
                                    _ = v37
                                }
                            default:
                                v1 += (v1 << ((42) & 63))
                            }
                            for fi38 := 0; fi38 < 8; fi38++ {
                                switch ("foo bar" + v6) {
                                case "":
                                    switch v6 {
                                    case "foo bar":
                                        if (v2 > v6) {
                                            var v39 uint8 = 2
                                            fmt.Println(v6)
                                            var v40 uint16 = 0
                                            _ = v39
                                            _ = v40
                                        }
                                    case "hello":
                                        switch ("a" + v2) {
                                        case "hello":
                                            v41 := GenS3{F0: 42}
                                            _ = v41
                                        case "foo bar":
                                            v12 -= (0.1 * 1e10)
                                        default:
                                            var v42 uint8 = 68
                                            _ = v42
                                        }
                                    default:
                                        var v43 uint8 = 42
                                        _ = v43
                                    }
                                case "hello":
                                    switch v2 {
                                    case "":
                                        if ((^(v4)) >= (v4 % ((v4) | 1))) {
                                            var v44 float64 = (v12 - v12)
                                            _ = v44
                                        }
                                    }
                                case "\n\t":
                                    fmt.Println(v3)
                                default:
                                    v4 += (v4 >> ((1) & 63))
                                }
                                v45 := GenS1{F0: (false && (!(!true))), F1: (^(v17)), F2: 0}
                                if true {
                                    v46 := GenS2{F0: false, F1: v17}
                                    _ = v46
                                }
                                switch 128 {
                                case 32768:
                                    switch (fi19 * v1) {
                                    case 32767:
                                        fmt.Println(v12, fi38)
                                    default:
                                        var v47 string = "\n\t"
                                        _ = v47
                                    }
                                    v48 := GenS4{F0: (v1 * (-(fi38))), F1: 127}
                                    if (255 > 2) {
                                        var v49 uint = 0
                                        if false {
                                            var v50 string = (("hello" + v6) + v6)
                                            delete(v5, (((("hello" + v50) + v6) + v6) + v6))
                                            if (1000 <= 0) {
                                                switch fi19 {
                                                case 128:
                                                    v50 = v6
                                                }
                                                v51 := GenS3{F0: 0}
                                                v52 := v5[v2]
                                                _ = v51
                                                _ = v52
                                            }
                                            _ = v50
                                        } else {
                                            v1 += fi19
                                        }
                                        var v53 int32 = 815
                                        _ = v49
                                        _ = v53
                                    }
                                    _ = v48
                                case 127:
                                    v4 -= v4
                                case 4294967295:
                                    var v54 uint32 = v3
                                    _ = v54
                                default:
                                    if ((v3 / (((^(v3))) | 1)) >= (v3 * v3)) {
                                        fmt.Println(v17, v3, fi38)
                                        v5["a"] = 2147483647
                                        var v55 int = v1
                                        switch (fi19 ^ v1) {
                                        case 32768:
                                            v12 -= 3.14159
                                        case -9223372036854775808:
                                            if (((v12 * (((0.1 * v12) - float64(-2.25)) - 1.5)) - float64(v9)) < v12) {
                                                switch (((("" + (v6 + ((v2 + v2) + v6))) + v6) + "a") + "hello") {
                                                case "foo bar":
                                                    v1 = v55
                                                }
                                            }
                                            if (true || false) {
                                                switch v6 {
                                                case "\n\t":
                                                    v5[v2] = v1
                                                }
                                                v56 := GenS2{F0: ((!true) && false), F1: v17}
                                                v57 := GenS4{F0: 42, F1: 128}
                                                var v58 uint8 = 2
                                                _ = v56
                                                _ = v57
                                                _ = v58
                                            } else {
                                                v59 := GenS2{F0: true, F1: (v17 / ((v17) | 1))}
                                                var v60 uint = 0
                                                _ = v59
                                                _ = v60
                                            }
                                            v3 = v3
                                        case 32767:
                                            var v61 bool = (!((1 != 1) || false))
                                            _ = v61
                                        default:
                                            v62 := GenS5{F0: 127}
                                            _ = v62
                                        }
                                        if (-1 < int32(v3)) {
                                            delete(v5, v6)
                                            switch (fi19 >> ((0) & 63)) {
                                            case 4294967295:
                                                var v63 uint32 = 569
                                                _ = v63
                                            default:
                                                var v64 uint16 = 255
                                                _ = v64
                                            }
                                        }
                                        _ = v55
                                    } else {
                                        if (0.0 >= v12) {
                                            v65 := GenS6{F0: 2147483647, F1: 128, F2: 1}
                                            var v66 uint = 186
                                            if len(v11) >= 2 {
                                                v11 = v11[1:len(v11)]
                                            }
                                            _ = v65
                                            _ = v66
                                        }
                                        if len(v11) > 0 {
                                            v11[0] = (v1 * 1)
                                        }
                                        v67 := GenS7{F0: true, F1: (v1 % (((v1 | (fi38 >> ((42) & 63)))) | 1))}
                                        if false {
                                            fmt.Println(v4, v1)
                                            var v68 float32 = v9
                                            v4 = (v4 % ((v4) | 1))
                                            delete(v5, (v6 + "hello"))
                                            _ = v68
                                        } else {
                                            var v69 uint64 = (^(v17))
                                            _ = v69
                                        }
                                        switch "\n\t" {
                                        case "a":
                                            if len(v11) > 0 {
                                                v11[0] = -481
                                            }
                                        case "":
                                            fmt.Println(v12)
                                        case "\n\t":
                                            var v70 string = "hello"
                                            _ = v70
                                        default:
                                            var v71 uint = 271
                                            _ = v71
                                        }
                                        _ = v67
                                    }
                                }
                                _ = fi38
                                _ = v45
                            }
                            for fi72 := 0; fi72 < 20; fi72++ {
                                var v73 uint8 = 137
                                v3 += v3
                                var v74 float64 = (-2.25 + v12)
                                _ = fi72
                                _ = v73
                                _ = v74
                            }
                        }
                    }
                    _ = v20
                } else {
                    if (false || false) {
                        fmt.Println(v12, v17, v6)
                        if false {
                            v75 := v5[v6]
                            switch v75 {
                            case -2147483648:
                                var v76 float64 = v12
                                _ = v76
                            }
                            var v77 uint = 65535
                            _ = v75
                            _ = v77
                        } else {
                            var v78 string = v6
                            _ = v78
                        }
                        v12 += float64(1e10)
                    } else {
                        var v79 uint16 = 116
                        v2 = "hello"
                        if false {
                            var v80 uint64 = 183
                            switch (v1 - fi19) {
                            case 4294967295:
                                v3 = v3
                            case 42:
                                v12 += 1e10
                            case 0:
                                v81 := GenS5{F0: 84}
                                _ = v81
                            default:
                                fmt.Println(v2, v80, v4)
                            }
                            _ = v80
                        }
                        _ = v79
                    }
                    fmt.Println(fi19, v1)
                }
            }
            v17 = v17
            v4 = (v4 >> ((18446744073709551615) & 63))
            _ = fi19
        }
        _ = v18
    } else {
        delete(v5, ((v2 + "a") + "\n\t"))
        for fi82 := 0; fi82 < 19; fi82++ {
            v4 = (v4 << ((18446744073709551615) & 63))
            var v83 float64 = (v12 + v12)
            _ = fi82
            _ = v83
        }
        var v84 float64 = float64((1.5 - v9))
        _ = v84
    }
    v9 -= (v9 + (-2.25 + float32(v12)))
    for fi85 := 0; fi85 < 14; fi85++ {
        v86 := GenS1{F0: (false && false), F1: v17, F2: 42}
        v4 = 127
        for fi87 := 0; fi87 < 5; fi87++ {
            v12 -= v12
            var v88 uint32 = v3
            v89 := GenS8{F0: -1, F1: 0, F2: 806}
            _ = fi87
            _ = v88
            _ = v89
        }
        _ = fi85
        _ = v86
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
