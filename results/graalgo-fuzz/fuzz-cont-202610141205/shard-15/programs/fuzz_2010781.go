package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 bool; F1 int64; F2 uint32 }
type GenS7 struct { F0 bool; F1 int32 }
type GenS8 struct { F0 int64; F1 bool; F2 uint8 }
type GenS9 struct { F0 int64; F1 int64 }
type GenS10 struct { F0 uint32 }

func main() {
    v1 := make(map[int]int)
    var v2 int16 = 765
    var v3 uint8 = 25
    v4 := make(map[string]string)
    var v5 uint8 = (v3 | v3)
    var v6 uint = 0
    switch 32768 {
    case -495:
        v7 := GenS1{F0: 0, F1: 32767}
        _ = v7
    case -9223372036854775808:
        fmt.Println(v6)
    default:
        v6 = (^(v6))
    }
    v6 = (v6 + (v6 >> ((1) & 63)))
    v8 := v1[42]
    var v9 bool = false
    if (615 > 171) {
        v5 -= 127
        if (!(-128 < 100)) {
            v10 := []int{128}
            for fi11 := 0; fi11 < 2; fi11++ {
                v9 = v9
                switch (v8 / ((fi11) | 1)) {
                case -862:
                    var v12 int32 = 2147483647
                    _ = v12
                case 1:
                    var v13 float32 = 3.14159
                    _ = v13
                default:
                    for fi14 := 0; fi14 < 1; fi14++ {
                        v3 -= uint8(v6)
                        var v15 int = fi11
                        if (42 != 100) {
                            if true {
                                if false {
                                    switch v15 {
                                    case -1:
                                        v9 = true
                                    case -2147483648:
                                        var v16 uint = 1
                                        _ = v16
                                    case 424:
                                        var v17 uint16 = 257
                                        var v18 int16 = (v2 >> ((2) & 63))
                                        v2 -= v18
                                        _ = v17
                                        _ = v18
                                    }
                                } else {
                                    if v9 {
                                        fmt.Println(v5, v8, fi14)
                                        v15 += 9223372036854775807
                                    } else {
                                        v10 = append(v10, -452)
                                    }
                                    v19 := GenS2{F0: 80}
                                    delete(v1, v8)
                                    _ = v19
                                }
                            } else {
                                v20 := GenS2{F0: -1}
                                _ = v20
                            }
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                            v21 := GenS3{F0: 842}
                            _ = v21
                        } else {
                            v5 -= v3
                            var v22 uint32 = 0
                            v22 -= 2147483648
                            v23 := GenS2{F0: 42}
                            _ = v22
                            _ = v23
                        }
                        _ = fi14
                        _ = v15
                    }
                }
                v5 += (v3 << ((572) & 63))
                _ = fi11
            }
            if (-9223372036854775808 >= 0) {
                for fi24 := 0; fi24 < 12; fi24++ {
                    if true {
                        for fi25 := 0; fi25 < 12; fi25++ {
                            v2 -= v2
                            switch "" {
                            case "":
                                v26 := GenS1{F0: 947, F1: v2}
                                _ = v26
                            case "a":
                                var v27 int32 = 0
                                _ = v27
                            default:
                                v9 = (false || (-917 >= 0))
                            }
                            switch "" {
                            case "a":
                                if len(v10) >= 2 {
                                    v10 = v10[1:len(v10)]
                                }
                            case "foo bar":
                                if (v9 || (1 == (v5 & v5))) {
                                    v2 = v2
                                    var v28 float32 = 3.14159
                                    if (!(902 != -632)) {
                                        if false {
                                            v2 -= 824
                                            fmt.Println(fi24, v2, v28)
                                            v28 += v28
                                        } else {
                                            v8 -= 276
                                            switch ("a" + (("foo bar" + "a") + ("a" + ("a" + "foo bar")))) {
                                            case "":
                                                if v9 {
                                                    if false {
                                                        v3 -= (v5 ^ v3)
                                                        if (2147483647 > 1) {
                                                            if v9 {
                                                                v29 := GenS1{F0: 1, F1: (v2 << (((v6 >> (((^(v6))) & 63))) & 63))}
                                                                _ = v29
                                                            }
                                                        } else {
                                                            v3 = v3
                                                        }
                                                    } else {
                                                        var v30 string = "foo bar"
                                                        if (false && v9) {
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                if ((v3 >> ((v6) & 63)) == (v5 - 1)) {
                                                                    v31 := GenS3{F0: 32767}
                                                                    switch (^(fi25)) {
                                                                    case 127:
                                                                        v3 = v5
                                                                    }
                                                                    v32 := GenS2{F0: -86}
                                                                    delete(v1, v8)
                                                                    _ = v31
                                                                    _ = v32
                                                                } else {
                                                                    var v33 float64 = 0.1
                                                                    v1[(fi25 >> (((v6 + (v6 % ((328) | 1)))) & 63))] = 4294967295
                                                                    _ = v33
                                                                }
                                                            case "foo bar":
                                                                v34 := GenS1{F0: 0, F1: v2}
                                                                _ = v34
                                                            default:
                                                                v2 = (v2 >> ((v6) & 63))
                                                            }
                                                            var v35 uint8 = v5
                                                            v10 = append(v10, fi24)
                                                            _ = v35
                                                        } else {
                                                            var v36 float64 = 3.14159
                                                            v36 += (v36 + v36)
                                                            v37 := GenS4{F0: -1}
                                                            _ = v36
                                                            _ = v37
                                                        }
                                                        _ = v30
                                                    }
                                                    switch fi24 {
                                                    case 0:
                                                        var v38 int8 = -128
                                                        _ = v38
                                                    case 413:
                                                        v39 := GenS4{F0: 898}
                                                        _ = v39
                                                    default:
                                                        _ = len(v4)
                                                    }
                                                } else {
                                                    v3 += v5
                                                    var v40 uint16 = 296
                                                    v41 := GenS5{F0: -1}
                                                    _ = v40
                                                    _ = v41
                                                }
                                            case "a":
                                                fmt.Println(v3)
                                            case "\n\t":
                                                v8 = -9223372036854775808
                                            default:
                                                v8 += (fi24 / ((fi24) | 1))
                                            }
                                            v4["a"] = (("foo bar" + "a") + "foo bar")
                                        }
                                        var v42 float32 = 1.5
                                        _ = v42
                                    } else {
                                        v2 += (v2 / (((v2 ^ -474)) | 1))
                                        if v9 {
                                            if ((805 < 341) && false) {
                                                var v43 float32 = (3.14159 / v28)
                                                _ = v43
                                            }
                                            fmt.Println(v5, fi24, v9)
                                        }
                                    }
                                    _ = v28
                                } else {
                                    fmt.Println(v9, v6)
                                    switch "a" {
                                    case "hello":
                                        v44 := GenS6{F0: (!("" > ((("foo bar" + "\n\t") + (("foo bar" + "") + "hello")) + ("\n\t" + "")))), F1: 42, F2: 4294967295}
                                        _ = v44
                                    default:
                                        v5 = v5
                                        v45 := GenS2{F0: 64}
                                        _ = v45
                                    }
                                }
                            case "hello":
                                v46 := GenS7{F0: ("\n\t" == ("hello" + "foo bar")), F1: -428}
                                _ = v46
                            }
                            _ = fi25
                        }
                    }
                    v47 := GenS3{F0: 65535}
                    switch "\n\t" {
                    case "\n\t":
                        var v48 uint16 = 42
                        _ = v48
                    default:
                        v49 := GenS1{F0: 4294967295, F1: v2}
                        _ = v49
                    }
                    _ = fi24
                    _ = v47
                }
                v50 := GenS5{F0: 707}
                if v9 {
                    v51 := GenS8{F0: -1, F1: v9, F2: 42}
                    if len(v10) > 0 {
                        v10[0] = v8
                    }
                    _ = v51
                } else {
                    for fi52 := 0; fi52 < 11; fi52++ {
                        v3 += v3
                        _ = fi52
                    }
                    var v53 uint64 = 0
                    _ = v53
                }
                _ = v50
            }
            _ = v10
        }
    } else {
        switch (v8 >> ((476) & 63)) {
        case 128:
            var v54 int = (^(v8))
            _ = v54
        case -81:
            v1[(v8 << (((v6 - v6)) & 63))] = (-(v8))
        }
    }
    if (35 > -1) {
        var v55 int = -842
        switch "foo bar" {
        case "a":
            for fi56 := 0; fi56 < 20; fi56++ {
                for fi57 := 0; fi57 < 11; fi57++ {
                    if v9 {
                        if (!((v2 | (v2 ^ int16(v8))) >= v2)) {
                            var v58 uint32 = 1
                            _ = v58
                        } else {
                            v5 = 1
                        }
                        var v59 uint = (v6 | (^(v6)))
                        v4["hello"] = ""
                        _ = v59
                    } else {
                        var v60 int16 = 1
                        v5 = v5
                        _ = len(v4)
                        _ = v60
                    }
                    var v61 uint16 = 2
                    _ = fi57
                    _ = v61
                }
                for fi62 := 0; fi62 < 3; fi62++ {
                    var v63 int = (fi62 ^ 4294967295)
                    fmt.Println(v9, v5, fi56)
                    _ = fi62
                    _ = v63
                }
                _ = fi56
            }
        case "":
            if (!v9) {
                var v64 uint64 = 685
                v65 := []int{(v8 + (v55 | v55)), (v55 >> ((v6) & 63)), (v55 | (v8 & v55)), (-(v8))}
                v1[v55] = (v55 - (v55 << ((v6) & 63)))
                _ = v64
                _ = v65
            }
        case "hello":
            v66 := GenS2{F0: -128}
            _ = v66
        }
        if v9 {
            if v9 {
                var v67 uint64 = 773
                if v9 {
                    v68 := []int{(v8 & (v8 * v55))}
                    switch "\n\t" {
                    case "hello":
                        for fi69 := 0; fi69 < 10; fi69++ {
                            v67 = 35
                            v9 = ((v5 << ((v6) & 63)) > (^(v5)))
                            v70 := GenS1{F0: 911, F1: 987}
                            _ = fi69
                            _ = v70
                        }
                    case "\n\t":
                        if false {
                            switch ("hello" + "") {
                            case "a":
                                for fi71 := 0; fi71 < 4; fi71++ {
                                    var v72 int8 = 0
                                    _ = fi71
                                    _ = v72
                                }
                            }
                            v55 = 127
                        } else {
                            v67 -= v67
                            var v73 int32 = 2147483647
                            _ = v73
                        }
                    }
                    fmt.Println(v55)
                    _ = v68
                }
                v74 := GenS1{F0: 18446744073709551615, F1: v2}
                _ = v67
                _ = v74
            }
            var v75 int = (v55 | v55)
            v9 = (v9 && (((1.5 + ((((-2.25 * (((0.0 - (0.0 * float32(((1.5 * (-2.25 / 1.0)) + 0.0)))) / 1e10) / ((0.0 + ((float32(1e10) / 1.0) * float32(0.1))) / (float32((1.5 + (1.5 - 0.0))) - 0.0)))) - 1.5) * (float32(float64(((((float32(0.0) * 1e10) + float32(((-2.25 + 0.0) + (0.0 / 0.1)))) - 1e10) - 1e10))) - 0.1)) / float32(0.1))) != 1e10) || ((v2 ^ (v2 + 42)) != (v2 << (((v6 & v6)) & 63)))))
            _ = v75
        } else {
            v55 += (-(v55))
            fmt.Println(v8, v55)
        }
        _ = v55
    } else {
        v76 := []int{}
        if ("foo bar" <= "") {
            var v77 int16 = v2
            switch (-(v8)) {
            case 9223372036854775807:
                if (0 >= -546) {
                    v78 := v4["foo bar"]
                    v79 := GenS5{F0: 42}
                    _ = v78
                    _ = v79
                } else {
                    if len(v76) >= 2 {
                        v76 = v76[1:len(v76)]
                    }
                }
            case 970:
                v77 = v77
            case -463:
                v80 := GenS9{F0: 42, F1: 42}
                _ = v80
            default:
                for fi81 := 0; fi81 < 13; fi81++ {
                    if (!v9) {
                        var v82 uint = 875
                        delete(v1, (^(v8)))
                        v83 := GenS10{F0: 2}
                        _ = v82
                        _ = v83
                    }
                    _ = fi81
                }
            }
            _ = v77
        } else {
            fmt.Println(v9)
            v84 := GenS7{F0: v9, F1: -2147483648}
            if v9 {
                v8 = (-(v8))
            } else {
                var v85 bool = true
                var v86 float32 = 1.5
                _ = v85
                _ = v86
            }
            _ = v84
        }
        v6 += (v6 / ((255) | 1))
        _ = v76
    }
    var v87 string = "\n\t"
    switch v87 {
    case "a":
        v8 += v8
    case "foo bar":
        var v88 int8 = 42
        _ = v88
    }
    fmt.Println(v5, v6, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v87:", v87)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v87
    _ = v9
}
