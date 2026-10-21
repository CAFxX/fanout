package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32; F1 int16 }
type GenS5 struct { F0 uint16; F1 uint8 }
type GenS6 struct { F0 int16; F1 uint; F2 int8 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint8; F1 int8 }
type GenS11 struct { F0 int; F1 int16; F2 int8 }
type GenS12 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 float64 = 1e10
    var v3 string = (((((("" + "hello") + ((("foo bar" + "") + "a") + "a")) + "") + "a") + ("hello" + "\n\t")) + ("a" + "hello"))
    switch (v3 + "foo bar") {
    case "hello":
        var v4 int32 = 42
        _ = v4
    case "foo bar":
        for fi5 := 0; fi5 < 4; fi5++ {
            v6 := GenS1{F0: 0, F1: 1, F2: -2147483648}
            var v7 uint16 = 32767
            v7 = v7
            _ = fi5
            _ = v6
            _ = v7
        }
    }
    var v8 int64 = -9223372036854775808
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 bool = false
        _ = fi9
        _ = v10
    }
    v8 = v8
    if (1 >= 2) {
        var v11 int = -9223372036854775808
        _ = v11
    }
    if (1 <= 1) {
        var v12 int = 2147483647
        _ = v12
    }
    if (true && (255 == 62)) {
        switch -9223372036854775808 {
        case 857:
            if (-128 >= 42) {
                if true {
                    var v13 int16 = 1
                    _ = v13
                } else {
                    fmt.Println(v3, v3, v2)
                    var v14 float32 = -2.25
                    switch -81 {
                    case -1:
                        for fi15 := 0; fi15 < 20; fi15++ {
                            for fi16 := 0; fi16 < 18; fi16++ {
                                v17 := []int{27, fi16, 127}
                                if (true && ((v8 == v8) && (0.1 > ((((((v14 + 0.0) - 1e10) - 0.1) - float32((1e10 / v2))) - 0.0) * v14)))) {
                                    switch ("\n\t" + "a") {
                                    case "":
                                        v3 = (v3 + ("hello" + (v3 + v3)))
                                        if false {
                                            if true {
                                                var v18 string = ""
                                                fmt.Println(fi15, fi16, v8)
                                                v19 := GenS1{F0: 2147483647, F1: 189, F2: 241}
                                                _ = v18
                                                _ = v19
                                            } else {
                                                if ("\n\t" != v3) {
                                                    switch v3 {
                                                    case "hello":
                                                        _ = len(v1)
                                                    case "a":
                                                        if len(v17) > 0 {
                                                            v17[0] = (^(fi15))
                                                        }
                                                        var v20 uint8 = 200
                                                        _ = v20
                                                    case "\n\t":
                                                        v21 := v1[65535]
                                                        _ = v21
                                                    }
                                                } else {
                                                    switch v3 {
                                                    case "foo bar":
                                                        v8 = v8
                                                    case "":
                                                        if (!(!false)) {
                                                            fmt.Println(fi16)
                                                            v2 += (((v2 * 0.1) - 3.14159) * v2)
                                                            var v22 int32 = 862
                                                            _ = v22
                                                        } else {
                                                            v14 += ((v14 - ((v14 - ((3.14159 * 1.5) / float32(((1e10 / 1e10) / (0.0 / v2))))) * ((v14 - ((v14 / v14) / (v14 / 3.14159))) / v14))) - (v14 + ((v14 - v14) - (v14 + (float32(v2) + 1e10)))))
                                                            var v23 int16 = -1
                                                            _ = v23
                                                        }
                                                    case "hello":
                                                        var v24 int = fi15
                                                        _ = v24
                                                    }
                                                    v25 := GenS2{F0: 2147483648, F1: v3}
                                                    fmt.Println(fi16, fi15)
                                                    _ = v25
                                                }
                                                v26 := GenS3{F0: ("" + (v3 + v3))}
                                                v14 = v14
                                                _ = v26
                                            }
                                        }
                                        if true {
                                            v8 = v8
                                            switch 472 {
                                            case 9223372036854775807:
                                                fmt.Println(fi16, fi16, fi15)
                                            case -9223372036854775808:
                                                var v27 int32 = 42
                                                v28 := GenS4{F0: 344, F1: 32767}
                                                if len(v17) >= 2 {
                                                    v17 = v17[1:len(v17)]
                                                }
                                                _ = v27
                                                _ = v28
                                            case 333:
                                                if (2 != 2) {
                                                    fmt.Println(v14)
                                                    v29 := GenS5{F0: 2, F1: 166}
                                                    switch (fi15 / (((fi15 + fi16)) | 1)) {
                                                    case -1:
                                                        _ = len(v1)
                                                        if len(v17) > 0 {
                                                            v17[0] = (fi15 / (((fi15 >> ((255) & 63))) | 1))
                                                        }
                                                        var v30 uint8 = 2
                                                        _ = v30
                                                    default:
                                                        v31 := GenS3{F0: (v3 + "")}
                                                        _ = v31
                                                    }
                                                    var v32 uint = 693
                                                    _ = v29
                                                    _ = v32
                                                } else {
                                                    v1[65535] = "foo bar"
                                                    var v33 bool = true
                                                    var v34 uint32 = 408
                                                    _ = v33
                                                    _ = v34
                                                }
                                            }
                                            var v35 int64 = v8
                                            _ = v35
                                        } else {
                                            if false {
                                                switch 1 {
                                                case 9223372036854775807:
                                                    if (!false) {
                                                        var v36 int8 = -1
                                                        if false {
                                                            switch (-(fi16)) {
                                                            case 2147483647:
                                                                var v37 bool = (false || (!false))
                                                                _ = v37
                                                            case 1:
                                                                var v38 bool = (!(0 == 2))
                                                                _ = v38
                                                            case 826:
                                                                var v39 uint64 = 0
                                                                _ = v39
                                                            }
                                                        }
                                                        _ = v36
                                                    }
                                                case 127:
                                                    _ = len(v17)
                                                default:
                                                    fmt.Println(v3, v8, v8)
                                                }
                                                var v40 int = 414
                                                fmt.Println(v2, v40)
                                                _ = v40
                                            }
                                        }
                                    case "\n\t":
                                        if (-139 > 42) {
                                            v41 := GenS6{F0: 257, F1: 0, F2: 56}
                                            var v42 uint16 = 255
                                            _ = v41
                                            _ = v42
                                        } else {
                                            v8 += v8
                                            fmt.Println(fi15)
                                        }
                                    default:
                                        if len(v17) >= 2 {
                                            v17 = v17[1:len(v17)]
                                        }
                                    }
                                    var v43 bool = (875 < 255)
                                    _ = len(v17)
                                    var v44 float32 = 0.0
                                    var v45 uint8 = 184
                                    v46 := GenS1{F0: 42, F1: 2147483647, F2: -1}
                                    _ = v43
                                    _ = v44
                                    _ = v45
                                    _ = v46
                                }
                                _ = fi16
                                _ = v17
                            }
                            switch fi15 {
                            case 128:
                                if false {
                                    if true {
                                        delete(v1, 2)
                                    }
                                } else {
                                    if false {
                                        var v47 int16 = 1000
                                        var v48 float32 = -2.25
                                        v49 := v1[65535]
                                        _ = v47
                                        _ = v48
                                        _ = v49
                                    }
                                }
                            default:
                                var v50 uint32 = 0
                                _ = v50
                            }
                            _ = fi15
                        }
                    case -9223372036854775808:
                        v3 = "foo bar"
                    }
                    _ = v14
                }
            } else {
                v1[42] = v3
                var v51 int16 = 831
                v52 := GenS4{F0: 2, F1: int16(v8)}
                _ = v51
                _ = v52
            }
        default:
            if (v8 < v8) {
                v3 = ("foo bar" + ((v3 + v3) + v3))
            } else {
                fmt.Println(v8, v2, v8)
                v3 = v3
                var v53 int32 = 520
                _ = v53
            }
        }
        v54 := GenS7{F0: (^(v8))}
        _ = v54
    } else {
        delete(v1, 4294967295)
        v8 = 42
    }
    delete(v1, 1)
    v55 := GenS8{F0: 199}
    switch 1 {
    case -9223372036854775808:
        var v56 bool = false
        _ = v56
    }
    switch v3 {
    case "hello":
        fmt.Println(v3)
    }
    switch (v3 + v3) {
    case "a":
        var v57 uint16 = 255
        _ = v57
    case "\n\t":
        if (-128 <= 57) {
            for fi58 := 0; fi58 < 5; fi58++ {
                fmt.Println(v2, v3, fi58)
                if true {
                    var v59 uint8 = 232
                    for fi60 := 0; fi60 < 8; fi60++ {
                        v8 = v8
                        fmt.Println(fi58, v2)
                        _ = fi60
                    }
                    var v61 int32 = -2147483648
                    _ = v59
                    _ = v61
                } else {
                    v62 := GenS9{F0: 42}
                    _ = v62
                }
                _ = fi58
            }
            v63 := GenS10{F0: 2, F1: 100}
            v3 = (("a" + "foo bar") + "")
            _ = v63
        } else {
            for fi64 := 0; fi64 < 3; fi64++ {
                v8 -= (v8 - v8)
                _ = fi64
            }
            for fi65 := 0; fi65 < 7; fi65++ {
                v8 += (v8 + v8)
                v66 := []int{(^(fi65))}
                _ = fi65
                _ = v66
            }
        }
    case "":
        v3 = ("" + ("a" + (v3 + v3)))
    default:
        if false {
            if false {
                var v67 int = 32767
                for fi68 := 0; fi68 < 7; fi68++ {
                    switch 2147483647 {
                    case -2147483648:
                        switch 32768 {
                        case 42:
                            var v69 int8 = 42
                            _ = v69
                        case 523:
                            v70 := []int{(fi68 / (((v67 % (((fi68 | v67)) | 1))) | 1)), (^(fi68)), (-(fi68))}
                            _ = v70
                        }
                    case 4294967295:
                        if ((0 == (v8 ^ 0)) && ((true && (!(2147483648 == 42))) || (255 > 128))) {
                            if false {
                                v67 = v67
                                v67 = (fi68 >> ((1) & 63))
                            } else {
                                for fi71 := 0; fi71 < 16; fi71++ {
                                    if true {
                                        if false {
                                            v72 := GenS11{F0: -9223372036854775808, F1: 1, F2: 100}
                                            if false {
                                                var v73 string = (v3 + ("hello" + v3))
                                                v2 = (float64(3.14159) / v2)
                                                var v74 int64 = v8
                                                if true {
                                                    v75 := []int{v67, 802, (fi68 - (fi71 | (fi71 * 0)))}
                                                    _ = v75
                                                }
                                                _ = v73
                                                _ = v74
                                            } else {
                                                var v76 int = v67
                                                var v77 float32 = 0.1
                                                switch (v76 << ((4294967295) & 63)) {
                                                case 2147483647:
                                                    v3 = (v3 + v3)
                                                }
                                                _ = v76
                                                _ = v77
                                            }
                                            var v78 uint32 = 2147483648
                                            _ = v72
                                            _ = v78
                                        }
                                    }
                                    var v79 uint16 = 65535
                                    v80 := []int{(-(fi71))}
                                    var v81 float32 = -2.25
                                    if len(v80) > 0 {
                                        v80[0] = 768
                                    }
                                    _ = fi71
                                    _ = v79
                                    _ = v80
                                    _ = v81
                                }
                            }
                            v82 := []int{}
                            if false {
                                var v83 int32 = 0
                                var v84 uint8 = 149
                                _ = v83
                                _ = v84
                            }
                            _ = v82
                        }
                    default:
                        v85 := GenS4{F0: 176, F1: -602}
                        _ = v85
                    }
                    v86 := []int{}
                    v87 := GenS7{F0: v8}
                    _ = fi68
                    _ = v86
                    _ = v87
                }
                _ = v67
            }
            v88 := GenS12{F0: 694}
            v89 := GenS7{F0: (v8 << ((42) & 63))}
            _ = v88
            _ = v89
        }
    }
    var v90 int = -53
    for fi91 := 0; fi91 < 7; fi91++ {
        delete(v1, 0)
        var v92 bool = false
        _ = fi91
        _ = v92
    }
    v93 := v1[65535]
    fmt.Println(v2)
    v8 += (v8 + v8)
    switch (v90 / ((-9223372036854775808) | 1)) {
    case 127:
        switch "" {
        case "a":
            v94 := GenS5{F0: 170, F1: 127}
            _ = v94
        default:
            v8 = v8
        }
    case 128:
        var v95 float32 = ((0.1 / 3.14159) + (3.14159 + 1e10))
        _ = v95
    case -9223372036854775808:
        v96 := GenS6{F0: 914, F1: 18446744073709551615, F2: 12}
        _ = v96
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v90:", v90)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v2
    _ = v3
    _ = v55
    _ = v8
    _ = v90
    _ = v93
}
