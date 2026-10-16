package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint }
type GenS3 struct { F0 uint64; F1 int64; F2 uint64 }
type GenS4 struct { F0 int; F1 int16; F2 bool }
type GenS5 struct { F0 uint64; F1 int8 }
type GenS6 struct { F0 uint8; F1 int8 }
type GenS7 struct { F0 int16; F1 int; F2 int16 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint8 }

func main() {
    v1 := make(map[int]string)
    var v2 float64 = 3.14159
    var v3 uint32 = 75
    v4 := []int{}
    var v5 bool = true
    var v6 string = ("\n\t" + "a")
    for fi7 := 0; fi7 < 18; fi7++ {
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 float32 = 0.1
            _ = fi8
            _ = v9
        }
        fmt.Println(v6, fi7, v2)
        _ = fi7
    }
    switch v6 {
    case "a":
        fmt.Println(v2, v6, v3)
    case "hello":
        for fi10 := 0; fi10 < 10; fi10++ {
            switch (v6 + v6) {
            case "":
                if v5 {
                    if (42 < 277) {
                        for fi11 := 0; fi11 < 18; fi11++ {
                            var v12 int32 = -2147483648
                            _ = fi11
                            _ = v12
                        }
                        for fi13 := 0; fi13 < 10; fi13++ {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = len(v4)
                            v5 = false
                            _ = fi13
                        }
                        v5 = (!(435 < 431))
                    }
                    for fi14 := 0; fi14 < 11; fi14++ {
                        v2 = v2
                        _ = fi14
                    }
                    if (0 > -1) {
                        var v15 uint32 = v3
                        if (100 != 127) {
                            if (v6 == "foo bar") {
                                v5 = v5
                                var v16 uint32 = v15
                                _ = v16
                            }
                            for fi17 := 0; fi17 < 15; fi17++ {
                                var v18 uint8 = 128
                                v5 = ((0.1 / 0.1) >= 1e10)
                                switch ((v6 + v6) + v6) {
                                case "foo bar":
                                    v18 -= v18
                                case "a":
                                    var v19 uint64 = 875
                                    _ = v19
                                case "\n\t":
                                    switch fi10 {
                                    case 42:
                                        if (true && v5) {
                                            v20 := GenS1{F0: v18}
                                            _ = v20
                                        }
                                    }
                                    var v21 uint16 = uint16(v15)
                                    switch 0 {
                                    case 32768:
                                        var v22 string = "a"
                                        v15 = 111
                                        v4 = append(v4, (fi10 & fi17))
                                        switch fi17 {
                                        case 42:
                                            v23 := GenS2{F0: (v3 >> ((1) & 63)), F1: v15, F2: 65535}
                                            _ = v23
                                        case 32768:
                                            var v24 float64 = (v2 - 1e10)
                                            _ = v24
                                        default:
                                            var v25 int8 = 42
                                            _ = v25
                                        }
                                        var v26 string = v22
                                        _ = v22
                                        _ = v26
                                    case 127:
                                        v27 := GenS3{F0: 0, F1: -9223372036854775808, F2: 196}
                                        _ = v27
                                    case 186:
                                        switch v6 {
                                        case "a":
                                            var v28 float64 = v2
                                            _ = v28
                                        default:
                                            var v29 uint = 0
                                            _ = v29
                                        }
                                    default:
                                        v30 := GenS3{F0: 42, F1: -937, F2: 42}
                                        _ = v30
                                    }
                                    _ = v21
                                }
                                _ = fi17
                                _ = v18
                            }
                            if (-9223372036854775808 >= -1) {
                                var v31 int8 = 127
                                _ = v31
                            }
                        }
                        fmt.Println(v3)
                        _ = v15
                    }
                }
            }
            _ = fi10
        }
    case "foo bar":
        if (42 > 255) {
            v5 = (-244 >= -1)
        }
    }
    var v32 float64 = v2
    fmt.Println(v6, v32)
    for fi33 := 0; fi33 < 10; fi33++ {
        v4 = append(v4, fi33)
        _ = fi33
    }
    var v34 int = -977
    v6 = "\n\t"
    v35 := GenS2{F0: (v3 << ((18446744073709551615) & 63)), F1: (v3 & 207), F2: 123}
    var v36 float64 = v2
    if true {
        for fi37 := 0; fi37 < 3; fi37++ {
            var v38 int8 = -56
            v39 := GenS4{F0: (^(fi37)), F1: 1000, F2: v5}
            _ = fi37
            _ = v38
            _ = v39
        }
        fmt.Println(v3, v3)
    }
    for fi40 := 0; fi40 < 10; fi40++ {
        fmt.Println(v3, fi40, v2)
        _ = fi40
    }
    var v41 uint64 = 144
    fmt.Println(v41)
    var v42 float32 = float32(-2.25)
    for fi43 := 0; fi43 < 20; fi43++ {
        if (false && (true || false)) {
            for fi44 := 0; fi44 < 18; fi44++ {
                var v45 int8 = int8(v41)
                _ = len(v1)
                _ = fi44
                _ = v45
            }
            fmt.Println(v41)
            fmt.Println(v5, fi43, v3)
        } else {
            var v46 uint = 42
            if (fi43 >= (fi43 * 425)) {
                v1[(v34 + (fi43 & -228))] = v6
                for fi47 := 0; fi47 < 19; fi47++ {
                    v41 += v41
                    var v48 uint32 = 439
                    var v49 int = (v34 * v34)
                    var v50 bool = v5
                    v1[v34] = "a"
                    if (-128 > 42) {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v51 uint64 = 0
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v52 := GenS4{F0: (v49 + fi43), F1: -133, F2: true}
                        _ = v51
                        _ = v52
                    } else {
                        if v50 {
                            var v53 int8 = 31
                            if (1e10 <= v2) {
                                if (((^(v41)) <= (v41 ^ v41)) || (254 > 53)) {
                                    fmt.Println(fi43, v36)
                                } else {
                                    v54 := GenS5{F0: 384, F1: v53}
                                    if v50 {
                                        switch v49 {
                                        case 4294967295:
                                            if false {
                                                v36 -= (float64(v42) * float64(v42))
                                                fmt.Println(v34)
                                            }
                                        case -210:
                                            if v5 {
                                                v55 := GenS6{F0: 0, F1: (^(v53))}
                                                v56 := GenS7{F0: 42, F1: v34, F2: int16(v53)}
                                                v2 += float64(3.14159)
                                                _ = v55
                                                _ = v56
                                            }
                                        }
                                        switch (v6 + v6) {
                                        case "foo bar":
                                            var v57 bool = v5
                                            _ = v57
                                        case "\n\t":
                                            if (true || (int16(v34) <= -32768)) {
                                                if ((v5 || false) && v50) {
                                                    v46 = (v46 - (v46 & (v46 * (v46 << (((v46 & 829)) & 63)))))
                                                    v4 = append(v4, (v34 / ((-9223372036854775808) | 1)))
                                                } else {
                                                    fmt.Println(v41, v49)
                                                    if false {
                                                        v2 -= (1e10 * v2)
                                                        if (4294967295 > 18446744073709551615) {
                                                            switch "\n\t" {
                                                            case "a":
                                                                var v58 float32 = ((((v42 / (v42 - v42)) * v42) + (v42 + v42)) - 1.5)
                                                                _ = v58
                                                            case "\n\t":
                                                                var v59 uint32 = 1
                                                                _ = v59
                                                            }
                                                        }
                                                        v46 += (v46 * 42)
                                                    } else {
                                                        v60 := GenS4{F0: 42, F1: 42, F2: (v50 || (18446744073709551615 >= (v46 ^ uint(v41))))}
                                                        var v61 uint16 = 42
                                                        var v62 bool = v5
                                                        var v63 string = v6
                                                        v64 := GenS1{F0: 128}
                                                        _ = v60
                                                        _ = v61
                                                        _ = v62
                                                        _ = v63
                                                        _ = v64
                                                    }
                                                    v36 = v32
                                                }
                                            }
                                        case "a":
                                            if ((v6 != ((v6 + "foo bar") + "foo bar")) && (false || true)) {
                                                v50 = ((v50 && true) && (-1 <= -281))
                                                switch v6 {
                                                case "foo bar":
                                                    v34 -= (fi43 ^ fi43)
                                                case "\n\t":
                                                    v36 = v2
                                                case "":
                                                    var v65 int8 = (v53 / ((v53) | 1))
                                                    _ = v65
                                                default:
                                                    var v66 float64 = float64((v42 - 3.14159))
                                                    if ((1.5 <= float64(((v42 + v42) / (v42 * 0.0)))) && false) {
                                                        if (!((^(v46)) == (^(v46)))) {
                                                            v50 = true
                                                        }
                                                        v67 := GenS2{F0: v3, F1: (v3 ^ uint32(v46)), F2: v46}
                                                        _ = v67
                                                    }
                                                    fmt.Println(v48, v34)
                                                    _ = v66
                                                }
                                            }
                                        }
                                    } else {
                                        v4 = append(v4, 2147483647)
                                        switch -1 {
                                        case 855:
                                            if false {
                                                if (702 >= -1) {
                                                    v68 := v1[v49]
                                                    if v50 {
                                                        var v69 uint64 = (v41 & v41)
                                                        var v70 int = -733
                                                        _ = v69
                                                        _ = v70
                                                    } else {
                                                        var v71 uint32 = (v3 * v3)
                                                        fmt.Println(v41, v36, v41)
                                                        _ = v71
                                                    }
                                                    switch (v68 + "\n\t") {
                                                    case "a":
                                                        v32 = float64(v42)
                                                        v49 -= v49
                                                        var v72 int16 = 1
                                                        _ = v72
                                                    case "\n\t":
                                                        v73 := GenS4{F0: (-(v34)), F1: 1, F2: (v34 < (fi47 - 568))}
                                                        _ = v73
                                                    }
                                                    _ = v68
                                                }
                                                v49 += 4294967295
                                                var v74 uint8 = 42
                                                fmt.Println(v46, v3, v36)
                                                var v75 float64 = float64(v42)
                                                v75 = v32
                                                _ = v74
                                                _ = v75
                                            } else {
                                                var v76 int16 = 1
                                                _ = v76
                                            }
                                        }
                                        v77 := GenS8{F0: (((v46 * 0) <= (v46 / (((v46 ^ v46)) | 1))) || (42 > 1000))}
                                        _ = v77
                                    }
                                    var v78 uint8 = 127
                                    _ = v54
                                    _ = v78
                                }
                                v79 := GenS2{F0: v48, F1: v3, F2: 633}
                                v4 = append(v4, -646)
                                if v5 {
                                    if len(v4) > 0 {
                                        v4[0] = fi47
                                    }
                                }
                                _ = v79
                            }
                            fmt.Println(v53, fi43)
                            v48 = 1
                            v36 = (v2 + v2)
                            if (!v5) {
                                var v80 uint8 = 5
                                _ = v80
                            } else {
                                v53 += v53
                                var v81 int32 = 837
                                switch ("" + "\n\t") {
                                case "a":
                                    switch 938 {
                                    case 9223372036854775807:
                                        var v82 uint8 = 1
                                        _ = v82
                                    case 2147483647:
                                        v83 := GenS7{F0: 428, F1: 9223372036854775807, F2: 0}
                                        _ = v83
                                    case -2147483648:
                                        v48 += v48
                                    default:
                                        fmt.Println(v6, v41, v49)
                                    }
                                case "hello":
                                    if false {
                                        v53 += v53
                                        v53 += (v53 - v53)
                                    } else {
                                        var v84 uint16 = uint16(v46)
                                        var v85 float32 = v42
                                        var v86 float64 = (v2 - ((3.14159 - v32) + 3.14159))
                                        _ = v84
                                        _ = v85
                                        _ = v86
                                    }
                                case "foo bar":
                                    var v87 float32 = (((0.1 - ((float32(((v36 - 0.0) - ((1.5 + 1.5) + v36))) / 1e10) / (1e10 + float32(float64((v42 + float32(0.0))))))) + 0.0) + ((v42 * -2.25) / (v42 * v42)))
                                    _ = v87
                                default:
                                    _ = len(v4)
                                }
                                _ = v81
                            }
                            var v88 bool = v50
                            _ = v53
                            _ = v88
                        } else {
                            if v50 {
                                switch v49 {
                                case -1:
                                    var v89 int64 = -9223372036854775808
                                    var v90 uint32 = (v48 % ((451) | 1))
                                    v3 = 755
                                    var v91 uint8 = 128
                                    _ = v89
                                    _ = v90
                                    _ = v91
                                case 32768:
                                    switch (v34 & v49) {
                                    case -2147483648:
                                        switch "a" {
                                        case "":
                                            v50 = (v41 < (v41 / (((v41 % ((v41) | 1))) | 1)))
                                        case "hello":
                                            v2 += v2
                                        case "\n\t":
                                            switch ((((v6 + ((v6 + (v6 + (v6 + ""))) + v6)) + "foo bar") + ("hello" + (v6 + (v6 + (v6 + v6))))) + v6) {
                                            case "hello":
                                                switch (v6 + v6) {
                                                case "\n\t":
                                                    v36 -= 0.1
                                                case "a":
                                                    v6 = "\n\t"
                                                }
                                            default:
                                                fmt.Println(v49, v3, fi43)
                                            }
                                        }
                                    default:
                                        var v92 int32 = 42
                                        _ = v92
                                    }
                                default:
                                    var v93 string = (("hello" + (v6 + v6)) + "a")
                                    _ = v93
                                }
                                var v94 uint8 = 0
                                _ = v94
                            }
                            var v95 bool = true
                            var v96 uint = v46
                            _ = v95
                            _ = v96
                        }
                    }
                    var v97 float32 = v42
                    _ = fi47
                    _ = v48
                    _ = v49
                    _ = v50
                    _ = v97
                }
                fmt.Println(v36)
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v6 = "foo bar"
            }
            v98 := GenS9{F0: 2}
            _ = v46
            _ = v98
        }
        v3 += 1
        _ = fi43
    }
    v34 -= (-(v34))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
