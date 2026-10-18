package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int32 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64; F1 int32 }
type GenS5 struct { F0 int32; F1 uint16; F2 bool }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint32; F1 int8; F2 uint16 }
type GenS8 struct { F0 uint8; F1 string; F2 string }
type GenS9 struct { F0 int16; F1 string }
type GenS10 struct { F0 string }
type GenS11 struct { F0 uint16 }
func genFn0() int8 {
    var v5 uint64 = 42
    v6 := GenS1{F0: true, F1: 0, F2: 42}
    for fi7 := 0; fi7 < 18; fi7++ {
        if false {
            for fi8 := 0; fi8 < 8; fi8++ {
                v9 := []int{}
                var v10 float64 = 1e10
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                var v11 uint8 = 222
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            switch "" {
            case "foo bar":
                v12 := GenS1{F0: (3.14159 != (0.0 / 0.1)), F1: 1, F2: -1}
                _ = v12
            case "\n\t":
                switch "" {
                case "hello":
                    fmt.Println(v5)
                }
            case "hello":
                if false {
                    if ("\n\t" <= "\n\t") {
                        for fi13 := 0; fi13 < 16; fi13++ {
                            var v14 int16 = -1
                            _ = fi13
                            _ = v14
                        }
                        v5 = (v5 % (((^(v5))) | 1))
                    }
                }
            }
            for fi15 := 0; fi15 < 2; fi15++ {
                if (816 != 742) {
                    switch "hello" {
                    case "\n\t":
                        v16 := []int{(fi15 | (fi15 >> ((255) & 63))), (fi7 | (fi15 / ((-2147483648) | 1))), (fi15 & fi15)}
                        _ = v16
                    case "hello":
                        v5 += v5
                    }
                    v5 = (^(v5))
                    v5 += v5
                    v17 := GenS2{F0: 0, F1: -23}
                    v18 := GenS3{F0: int8(v5)}
                    v5 = (v5 % ((v5) | 1))
                    if (42 >= 1) {
                        if false {
                            var v19 int64 = -846
                            var v20 uint8 = 128
                            _ = v19
                            _ = v20
                        }
                        v5 += (v5 + v5)
                    }
                    _ = v17
                    _ = v18
                }
                v21 := []int{}
                var v22 bool = false
                _ = fi15
                _ = v21
                _ = v22
            }
        }
        var v23 bool = (426 <= 1)
        _ = fi7
        _ = v23
    }
    _ = v5
    _ = v6
    return -1
}
func genFn1() int32 {
    var v24 float64 = 0.0
    var v25 uint8 = 75
    if false {
        var v26 uint64 = 794
        v27 := []int{9223372036854775807, 669, 42, -461}
        if (!false) {
            var v28 float32 = 1.5
            v29 := GenS4{F0: v26, F1: -1}
            for fi30 := 0; fi30 < 11; fi30++ {
                v24 = v24
                _ = fi30
            }
            _ = v28
            _ = v29
        } else {
            var v31 int = 9223372036854775807
            _ = len(v27)
            for fi32 := 0; fi32 < 10; fi32++ {
                if (((!(539 <= int16(v25))) || ((1 == 1) || false)) || (uint32(v25) == 929)) {
                    if (1.5 != 1.5) {
                        var v33 uint8 = (v25 & v25)
                        v31 -= v31
                        _ = v33
                    }
                    fmt.Println(v24, v24)
                    for fi34 := 0; fi34 < 3; fi34++ {
                        if len(v27) > 0 {
                            v27[0] = fi34
                        }
                        fmt.Println(v24, v31)
                        _ = fi34
                    }
                } else {
                    if ((0 <= 0) || true) {
                        var v35 int16 = -471
                        _ = v35
                    }
                    v25 = v25
                }
                _ = len(v27)
                v36 := GenS3{F0: 127}
                _ = fi32
                _ = v36
            }
            _ = v31
        }
        _ = v26
        _ = v27
    }
    v24 = (v24 + 1e10)
    _ = v24
    _ = v25
    return -875
}

func main() {
    v1 := make(map[string]string)
    var v2 float32 = (0.1 + float32((((0.1 / ((0.1 + -2.25) / ((0.0 * ((1e10 * 0.0) + 3.14159)) + (((1.5 / (0.1 * 1e10)) / 1.0) * -2.25)))) - 0.0) - 0.0)))
    var v3 uint8 = 170
    var v4 float32 = 3.14159
    for fi37 := 0; fi37 < 4; fi37++ {
        v3 = (v3 * (v3 ^ v3))
        fmt.Println(v4, fi37, v4)
        var v38 uint32 = 65535
        _ = fi37
        _ = v38
    }
    var v39 int16 = -1
    v39 += int16(v3)
    v2 = (v2 / (float32(1e10) * (v2 - (v4 + float32(((1.5 / -2.25) + (float64(v2) + (float64(1.5) * 0.0))))))))
    v40 := GenS4{F0: 42, F1: -2147483648}
    var v41 float64 = 0.1
    if ((v3 - 0) != v3) {
        var v42 int16 = (v39 - -1)
        _ = v42
    } else {
        v39 = 1000
        for fi43 := 0; fi43 < 6; fi43++ {
            if false {
                for fi44 := 0; fi44 < 1; fi44++ {
                    if false {
                        var v45 uint32 = 2147483648
                        switch "\n\t" {
                        case "\n\t":
                            switch fi44 {
                            case 32767:
                                v46 := []int{fi43}
                                if true {
                                    switch "" {
                                    case "hello":
                                        v2 = 1.5
                                    default:
                                        switch fi43 {
                                        case -856:
                                            v4 = float32(-2.25)
                                        case -1:
                                            switch ("hello" + ("\n\t" + (("hello" + "") + ("a" + "hello")))) {
                                            case "foo bar":
                                                v47 := GenS2{F0: -2147483648, F1: -128}
                                                _ = v47
                                            case "a":
                                                switch fi44 {
                                                case 32768:
                                                    var v48 float64 = float64(v4)
                                                    _ = v48
                                                case 832:
                                                    var v49 int8 = -1
                                                    _ = v49
                                                }
                                            case "hello":
                                                v46 = append(v46, 42)
                                            }
                                            var v50 bool = true
                                            switch fi44 {
                                            case 349:
                                                var v51 uint16 = 42
                                                _ = v51
                                            case 128:
                                                var v52 bool = (v50 && v50)
                                                _ = v52
                                            }
                                            v53 := GenS1{F0: (4294967295 != 377), F1: uint(v45), F2: 1}
                                            var v54 string = "\n\t"
                                            _ = v50
                                            _ = v53
                                            _ = v54
                                        default:
                                            v2 = -2.25
                                        }
                                    }
                                    v4 += v2
                                } else {
                                    v39 = v39
                                }
                                var v55 uint16 = 73
                                _ = v46
                                _ = v55
                            case 9223372036854775807:
                                if false {
                                    switch "hello" {
                                    case "foo bar":
                                        if (("a" + ("\n\t" + "a")) <= "foo bar") {
                                            switch "hello" {
                                            case "hello":
                                                if (46 > 55) {
                                                    var v56 float32 = v4
                                                    var v57 uint64 = 42
                                                    var v58 string = "foo bar"
                                                    _ = v56
                                                    _ = v57
                                                    _ = v58
                                                }
                                            case "":
                                                v59 := GenS2{F0: 1, F1: 87}
                                                _ = v59
                                            default:
                                                v45 += v45
                                            }
                                            switch ("a" + ("\n\t" + "a")) {
                                            case "":
                                                _ = len(v1)
                                            }
                                        } else {
                                            if (-71 != 1) {
                                                if true {
                                                    v60 := GenS5{F0: 0, F1: 775, F2: false}
                                                    v61 := GenS6{F0: 1}
                                                    var v62 float64 = (v41 - (float64(0.0) - (3.14159 + 1.5)))
                                                    _ = v60
                                                    _ = v61
                                                    _ = v62
                                                } else {
                                                    v39 += (-(v39))
                                                    var v63 uint32 = v45
                                                    v64 := GenS7{F0: v45, F1: 1, F2: 0}
                                                    _ = v63
                                                    _ = v64
                                                }
                                                var v65 int = fi43
                                                _ = v65
                                            }
                                            fmt.Println(fi43)
                                        }
                                    case "\n\t":
                                        switch (fi43 / ((int(v3)) | 1)) {
                                        case -94:
                                            v39 -= (v39 | 1)
                                        case 0:
                                            var v66 int8 = 42
                                            _ = v66
                                        case 2147483647:
                                            var v67 bool = (2 >= 1)
                                            _ = v67
                                        }
                                    }
                                    if true {
                                        v2 += v2
                                        v45 += 42
                                    }
                                    if (888 <= 638) {
                                        v68 := GenS7{F0: (v45 % ((v45) | 1)), F1: 1, F2: 9}
                                        _ = v68
                                    }
                                }
                            case -637:
                                if false {
                                    v2 = (((v2 + (0.0 * (v2 + v4))) - (v4 * v2)) * ((3.14159 / v2) - (1e10 / v2)))
                                }
                            }
                        case "":
                            if (fi44 <= 766) {
                                switch ("" + "hello") {
                                case "hello":
                                    fmt.Println(v39)
                                case "\n\t":
                                    v4 = (v2 + v4)
                                default:
                                    if true {
                                        v1[("\n\t" + "foo bar")] = ("foo bar" + "foo bar")
                                        if (!false) {
                                            var v69 uint32 = v45
                                            _ = v69
                                        }
                                        if false {
                                            var v70 uint8 = (^(v3))
                                            _ = v70
                                        }
                                        if (!true) {
                                            v71 := GenS6{F0: -67}
                                            if ((true && true) && false) {
                                                fmt.Println(v45)
                                                fmt.Println(fi43)
                                                v39 += (v39 & (v39 % (((v39 + v39)) | 1)))
                                            } else {
                                                v41 = -2.25
                                                fmt.Println(fi43)
                                                var v72 int8 = 1
                                                _ = v72
                                            }
                                            v73 := GenS8{F0: 186, F1: "hello", F2: "a"}
                                            switch ("\n\t" + "hello") {
                                            case "foo bar":
                                                var v74 int64 = -46
                                                _ = v74
                                            case "a":
                                                var v75 int8 = 127
                                                _ = v75
                                            }
                                            switch fi44 {
                                            case 1:
                                                var v76 int16 = (^(v39))
                                                _ = v76
                                            case -26:
                                                if (true || ((v41 * ((float64(((v4 + v4) / (v2 - v4))) / v41) + v41)) == (v41 - v41))) {
                                                    delete(v1, "hello")
                                                }
                                            case 2147483647:
                                                if (true || false) {
                                                    v2 -= v4
                                                } else {
                                                    v77 := GenS9{F0: (^(v39)), F1: ("foo bar" + ("" + ("hello" + "hello")))}
                                                    v78 := v1[""]
                                                    _ = v77
                                                    _ = v78
                                                }
                                            default:
                                                var v79 uint = 833
                                                _ = v79
                                            }
                                            _ = v71
                                            _ = v73
                                        } else {
                                            if false {
                                                v39 += (v39 & (-(v39)))
                                                var v80 uint64 = 1
                                                _ = v80
                                            } else {
                                                v2 -= 0.0
                                                if false {
                                                    var v81 float32 = 0.0
                                                    v3 = (v3 % (((^(v3))) | 1))
                                                    _ = v81
                                                } else {
                                                    v4 = v2
                                                    delete(v1, "hello")
                                                    switch (^(fi43)) {
                                                    case 42:
                                                        v82 := GenS9{F0: (v39 + v39), F1: (("hello" + "hello") + ("hello" + "\n\t"))}
                                                        _ = v82
                                                    case 1:
                                                        delete(v1, "\n\t")
                                                    }
                                                }
                                            }
                                        }
                                        if (v41 < (-2.25 * -2.25)) {
                                            v3 = v3
                                        }
                                    }
                                    v83 := []int{fi43}
                                    _ = v83
                                }
                                delete(v1, ("foo bar" + ""))
                            }
                        case "a":
                            var v84 string = ("foo bar" + ("foo bar" + "hello"))
                            _ = v84
                        default:
                            if ((!(fi43 == fi43)) || (104 > (v3 ^ (v3 - (^(v3)))))) {
                                v85 := GenS4{F0: 185, F1: 2147483647}
                                var v86 int64 = 42
                                _ = v85
                                _ = v86
                            } else {
                                v87 := GenS5{F0: 42, F1: 2, F2: false}
                                switch 4294967295 {
                                case 128:
                                    v88 := GenS8{F0: (v3 | 0), F1: "hello", F2: ""}
                                    _ = v88
                                default:
                                    if false {
                                        fmt.Println(fi43, fi43, fi43)
                                    }
                                }
                                _ = v87
                            }
                        }
                        if false {
                            v89 := GenS10{F0: (("a" + ("a" + "\n\t")) + ("" + ""))}
                            _ = v89
                        } else {
                            var v90 string = ""
                            _ = v90
                        }
                        _ = v45
                    } else {
                        var v91 int64 = 9223372036854775807
                        v39 -= (v39 + v39)
                        v92 := GenS11{F0: 0}
                        _ = v91
                        _ = v92
                    }
                    v93 := v1["hello"]
                    var v94 int32 = -1
                    _ = fi44
                    _ = v93
                    _ = v94
                }
            }
            _ = fi43
        }
    }
    var v95 float32 = -2.25
    var v96 float64 = v41
    v39 = -134
    var v97 uint32 = 731
    switch 32768 {
    case 1:
        v98 := []int{}
        _ = v98
    case 205:
        v99 := GenS7{F0: (^(v97)), F1: -99, F2: 1}
        _ = v99
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v95:", v95)
    fmt.Println("v96:", v96)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v95
    _ = v96
    _ = v97
}
