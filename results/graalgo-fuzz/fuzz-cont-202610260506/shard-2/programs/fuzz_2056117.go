package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 uint32; F1 uint32; F2 uint16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint; F1 int16 }
type GenS8 struct { F0 uint32; F1 uint16; F2 uint }
type GenS9 struct { F0 string; F1 int; F2 string }

func main() {
    var v1 uint8 = 65
    v2 := make(map[string]int)
    v3 := []string{((("hello" + "\n\t") + "\n\t") + ("foo bar" + "a")), ("foo bar" + "foo bar"), "a", "\n\t"}
    var v4 int8 = -68
    var v5 uint = uint(v1)
    if false {
        var v6 int = 1
        v4 = v4
        var v7 int8 = v4
        _ = v6
        _ = v7
    }
    var v8 int64 = 458
    var v9 uint64 = 659
    if false {
        if ((-(v8)) == (v8 * (v8 + v8))) {
            v1 = v1
            if false {
                for fi10 := 0; fi10 < 4; fi10++ {
                    switch "" {
                    case "foo bar":
                        for fi11 := 0; fi11 < 19; fi11++ {
                            v9 = v9
                            v9 += (v9 | (v9 ^ 18446744073709551615))
                            _ = fi11
                        }
                    default:
                        fmt.Println(fi10)
                    }
                    _ = fi10
                }
                v12 := GenS1{F0: 38, F1: (v9 << ((v5) & 63)), F2: v5}
                _ = v12
            } else {
                for fi13 := 0; fi13 < 14; fi13++ {
                    v14 := GenS1{F0: (v5 | (^(v5))), F1: v9, F2: v5}
                    switch fi13 {
                    case 4294967295:
                        v4 = (^(v4))
                    default:
                        var v15 int8 = v4
                        _ = v15
                    }
                    _ = len(v2)
                    _ = fi13
                    _ = v14
                }
                fmt.Println(v1, v8, v9)
            }
            v2[(("foo bar" + "\n\t") + "")] = 500
        } else {
            for fi16 := 0; fi16 < 9; fi16++ {
                var v17 float64 = 1.5
                var v18 uint16 = 255
                fmt.Println(v5, v18)
                _ = fi16
                _ = v17
                _ = v18
            }
        }
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        var v20 uint32 = 4294967295
        var v21 int = 128
        _ = fi19
        _ = v20
        _ = v21
    }
    delete(v2, "\n\t")
    for fi22 := 0; fi22 < 19; fi22++ {
        fmt.Println(v9, v8)
        if true {
            v23 := v2[(("" + "a") + "hello")]
            delete(v2, ("" + "hello"))
            _ = v23
        } else {
            _ = len(v3)
            v1 = (^(v1))
            v24 := GenS1{F0: (v5 >> ((v5) & 63)), F1: v9, F2: 4294967295}
            _ = v24
        }
        _ = fi22
    }
    var v25 float32 = 1.5
    var v26 bool = false
    v1 = v1
    var v27 string = (("a" + "hello") + "\n\t")
    if len(v3) > 0 {
        v3[0] = v27
    }
    if v26 {
        if (true && v26) {
            if (0.1 <= (3.14159 - (v25 / (1.5 - float32(0.0))))) {
                if (false && v26) {
                    if (0.1 <= -2.25) {
                        var v28 uint64 = v9
                        if v26 {
                            for fi29 := 0; fi29 < 2; fi29++ {
                                v1 += 41
                                for fi30 := 0; fi30 < 16; fi30++ {
                                    v1 = (^(v1))
                                    v31 := GenS2{F0: v27}
                                    switch v27 {
                                    case "\n\t":
                                        v8 += (-(v8))
                                    }
                                    _ = fi30
                                    _ = v31
                                }
                                for fi32 := 0; fi32 < 13; fi32++ {
                                    v26 = (v9 > (v9 << (((v5 / ((uint(v1)) | 1))) & 63)))
                                    fmt.Println(v8)
                                    if (2 <= (v28 + v28)) {
                                        if (true || v26) {
                                            v9 -= v28
                                            v28 = 0
                                            v33 := GenS1{F0: 42, F1: v9, F2: (v5 & (^(v5)))}
                                            _ = v33
                                        } else {
                                            _ = len(v2)
                                            var v34 uint16 = 607
                                            _ = v34
                                        }
                                    }
                                    if v26 {
                                        v27 = "\n\t"
                                        v8 += (v8 / ((int64(v4)) | 1))
                                        v2["foo bar"] = fi32
                                    } else {
                                        fmt.Println(fi32, v28)
                                    }
                                    if ((v1 % (((v1 & (v1 / (((v1 & 122)) | 1)))) | 1)) == (v1 - (^(v1)))) {
                                        v35 := GenS2{F0: (v27 + v27)}
                                        if (1 < 0) {
                                            v4 = v4
                                            if true {
                                                if (-2147483648 != 42) {
                                                    var v36 uint16 = 42
                                                    if (v26 && (uint32(v5) != 516)) {
                                                        v5 = v5
                                                        v26 = true
                                                        var v37 uint = 255
                                                        _ = v37
                                                    } else {
                                                        v38 := v2["a"]
                                                        _ = v38
                                                    }
                                                    _ = v36
                                                } else {
                                                    fmt.Println(v25, v9)
                                                    var v39 int16 = -23
                                                    var v40 int64 = (v8 + 9223372036854775807)
                                                    _ = v39
                                                    _ = v40
                                                }
                                                if v26 {
                                                    v26 = (42 > 0)
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    var v41 uint64 = (^(v9))
                                                    _ = v41
                                                } else {
                                                    switch 32767 {
                                                    case 32768:
                                                        v25 -= (v25 * (v25 * (v25 - 1.5)))
                                                    default:
                                                        v42 := GenS3{F0: 42, F1: v27}
                                                        _ = v42
                                                    }
                                                }
                                            }
                                            v8 = v8
                                        } else {
                                            var v43 bool = false
                                            var v44 bool = (((0 != -23) || (v5 < (^(v5)))) || (!(!v43)))
                                            switch fi32 {
                                            case 32767:
                                                var v45 uint8 = v1
                                                _ = v45
                                            case 826:
                                                fmt.Println(v43, v28, v43)
                                            case 531:
                                                switch (((v27 + v27) + v27) + v27) {
                                                case "hello":
                                                    var v46 int = fi32
                                                    _ = v46
                                                }
                                            }
                                            _ = v43
                                            _ = v44
                                        }
                                        _ = len(v2)
                                        v28 = v28
                                        v8 = v8
                                        _ = v35
                                    }
                                    fmt.Println(v28, v25, v1)
                                    _ = fi32
                                }
                                _ = fi29
                            }
                        } else {
                            if true {
                                if false {
                                    if ((("hello" + (v27 + v27)) + ("a" + "foo bar")) != (v27 + ("hello" + ("\n\t" + v27)))) {
                                        fmt.Println(v9)
                                    } else {
                                        if (2147483647 > 2147483648) {
                                            for fi47 := 0; fi47 < 5; fi47++ {
                                                v9 = 357
                                                v25 -= ((v25 - 0.1) / float32((float64((v25 * v25)) * 1e10)))
                                                _ = fi47
                                            }
                                            var v48 uint64 = 1
                                            _ = v48
                                        } else {
                                            switch v27 {
                                            case "a":
                                                var v49 int64 = v8
                                                _ = v49
                                            default:
                                                v5 -= (^(v5))
                                            }
                                            if (!v26) {
                                                fmt.Println(v9)
                                                if true {
                                                    v27 = "a"
                                                    if v26 {
                                                        if v26 {
                                                            fmt.Println(v28)
                                                        } else {
                                                            var v50 float64 = 1.5
                                                            if v26 {
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                            } else {
                                                                if ((v1 >> (((v5 ^ 184)) & 63)) < 251) {
                                                                    v51 := v2["a"]
                                                                    var v52 uint64 = uint64(v1)
                                                                    _ = v51
                                                                    _ = v52
                                                                } else {
                                                                    if (!v26) {
                                                                        if ((v1 >> ((v5) & 63)) > v1) {
                                                                            v28 = (^(v28))
                                                                            v53 := GenS2{F0: v27}
                                                                            _ = v53
                                                                        } else {
                                                                            _ = len(v2)
                                                                            var v54 int8 = (-(v4))
                                                                            _ = v54
                                                                        }
                                                                        var v55 int32 = -705
                                                                        var v56 float32 = 1e10
                                                                        _ = v55
                                                                        _ = v56
                                                                    }
                                                                    switch 32768 {
                                                                    case -2147483648:
                                                                        for fi57 := 0; fi57 < 13; fi57++ {
                                                                            v50 = (v50 - v50)
                                                                            _ = fi57
                                                                        }
                                                                    default:
                                                                        switch v27 {
                                                                        case "foo bar":
                                                                            for fi58 := 0; fi58 < 13; fi58++ {
                                                                                fmt.Println(v28)
                                                                                v50 -= float64((v25 * float32(float64(0.1))))
                                                                                var v59 uint64 = (^(v9))
                                                                                _ = fi58
                                                                                _ = v59
                                                                            }
                                                                        case "hello":
                                                                            var v60 int16 = 0
                                                                            _ = v60
                                                                        default:
                                                                            switch 702 {
                                                                            case 181:
                                                                                for fi61 := 0; fi61 < 9; fi61++ {
                                                                                    _ = len(v2)
                                                                                    if v26 {
                                                                                        v25 = 1.5
                                                                                        v4 -= (v4 | (v4 * v4))
                                                                                    }
                                                                                    var v62 int64 = v8
                                                                                    _ = fi61
                                                                                    _ = v62
                                                                                }
                                                                            case 933:
                                                                                var v63 bool = v26
                                                                                _ = v63
                                                                            case 0:
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                            default:
                                                                                var v64 float64 = v50
                                                                                _ = v64
                                                                            }
                                                                        }
                                                                    }
                                                                    v65 := GenS3{F0: (v5 & (v5 << (((v5 ^ (v5 - v5))) & 63))), F1: v27}
                                                                    _ = v65
                                                                }
                                                            }
                                                            _ = v50
                                                        }
                                                    }
                                                }
                                                v8 = v8
                                            } else {
                                                fmt.Println(v8)
                                                switch 2147483647 {
                                                case 42:
                                                    var v66 uint8 = (v1 * v1)
                                                    _ = v66
                                                case 127:
                                                    if v26 {
                                                        v9 = 4294967295
                                                    }
                                                }
                                            }
                                            if (v26 && v26) {
                                                v25 -= v25
                                                v67 := GenS4{F0: 377, F1: 899, F2: 252}
                                                v8 += (v8 - -9223372036854775808)
                                                _ = v67
                                            }
                                        }
                                    }
                                    v3 = append(v3, ("hello" + v27))
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                delete(v2, v27)
                            }
                            switch v27 {
                            case "foo bar":
                                v3 = append(v3, ("\n\t" + "foo bar"))
                            case "\n\t":
                                var v68 bool = v26
                                _ = v68
                            }
                            for fi69 := 0; fi69 < 1; fi69++ {
                                for fi70 := 0; fi70 < 11; fi70++ {
                                    var v71 bool = v26
                                    switch ("" + "\n\t") {
                                    case "a":
                                        if ((!v26) || ((v4 ^ v4) <= (v4 / ((15) | 1)))) {
                                            var v72 uint64 = v28
                                            v3 = append(v3, "foo bar")
                                            _ = v72
                                        } else {
                                            switch int(v1) {
                                            case 128:
                                                _ = len(v3)
                                                fmt.Println(v28)
                                            }
                                            switch int(v8) {
                                            case 128:
                                                if v71 {
                                                    fmt.Println(v27)
                                                }
                                            default:
                                                v73 := GenS5{F0: false}
                                                _ = v73
                                            }
                                            if v71 {
                                                if (v71 && (v71 || v71)) {
                                                    v74 := v2[v27]
                                                    _ = v74
                                                } else {
                                                    _ = len(v3)
                                                    fmt.Println(v4, fi70)
                                                    var v75 uint8 = (v1 % (((v1 | 2)) | 1))
                                                    _ = v75
                                                }
                                                v76 := GenS4{F0: 0, F1: 4294967295, F2: 698}
                                                _ = v76
                                            }
                                        }
                                    }
                                    var v77 float64 = 1e10
                                    _ = fi70
                                    _ = v71
                                    _ = v77
                                }
                                var v78 uint8 = 15
                                v79 := GenS6{F0: (v1 % ((v78) | 1))}
                                _ = fi69
                                _ = v78
                                _ = v79
                            }
                        }
                        _ = v28
                    } else {
                        var v80 float64 = 1e10
                        var v81 int16 = 32767
                        _ = v80
                        _ = v81
                    }
                    var v82 int8 = (v4 | (v4 >> (((^(v5))) & 63)))
                    if (-680 < -32768) {
                        if (!((0.0 / (v25 + (((3.14159 * v25) + float32(float64(float32(0.0)))) + v25))) >= (3.14159 + float32((0.1 * (1.5 + 3.14159)))))) {
                            var v83 int64 = v8
                            var v84 uint32 = 1
                            _ = v83
                            _ = v84
                        } else {
                            switch 4294967295 {
                            case -773:
                                switch (v27 + ("\n\t" + "")) {
                                case "":
                                    switch 9223372036854775807 {
                                    case 310:
                                        for fi85 := 0; fi85 < 10; fi85++ {
                                            if len(v3) > 0 {
                                                v3[0] = ("\n\t" + v27)
                                            }
                                            _ = fi85
                                        }
                                    case -172:
                                        v3 = append(v3, v27)
                                    }
                                case "hello":
                                    v27 = (v27 + v27)
                                default:
                                    v26 = false
                                }
                            case 32768:
                                var v86 int = 367
                                _ = v86
                            default:
                                fmt.Println(v1)
                            }
                        }
                    } else {
                        for fi87 := 0; fi87 < 1; fi87++ {
                            _ = len(v3)
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi87
                        }
                        for fi88 := 0; fi88 < 10; fi88++ {
                            var v89 uint32 = 0
                            _ = fi88
                            _ = v89
                        }
                    }
                    _ = v82
                }
                for fi90 := 0; fi90 < 18; fi90++ {
                    switch (fi90 % ((4294967295) | 1)) {
                    case 32767:
                        v25 += v25
                    case 231:
                        if false {
                            fmt.Println(v1, v27, v27)
                            var v91 uint16 = 1
                            var v92 float64 = (0.1 - (0.0 + (0.0 / 1e10)))
                            _ = v91
                            _ = v92
                        } else {
                            v26 = v26
                        }
                    default:
                        for fi93 := 0; fi93 < 17; fi93++ {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi93
                        }
                    }
                    if (v4 < -97) {
                        fmt.Println(v8, v8)
                        switch (fi90 + fi90) {
                        case 1:
                            for fi94 := 0; fi94 < 20; fi94++ {
                                v8 -= v8
                                fmt.Println(v9)
                                _ = fi94
                            }
                        case 42:
                            v9 = (v9 / ((4294967295) | 1))
                        case 4294967295:
                            v95 := GenS4{F0: 2, F1: 0, F2: 1}
                            _ = v95
                        default:
                            v96 := GenS2{F0: "hello"}
                            _ = v96
                        }
                        v27 = v27
                    }
                    v8 = 72
                    _ = fi90
                }
            }
            if (3.14159 <= 3.14159) {
                for fi97 := 0; fi97 < 9; fi97++ {
                    var v98 string = "hello"
                    _ = fi97
                    _ = v98
                }
            } else {
                if false {
                    switch -314 {
                    case -9223372036854775808:
                        var v99 bool = (345 < int(v1))
                        _ = v99
                    }
                } else {
                    for fi100 := 0; fi100 < 6; fi100++ {
                        var v101 string = "foo bar"
                        var v102 int64 = 0
                        fmt.Println(v101, v4)
                        _ = fi100
                        _ = v101
                        _ = v102
                    }
                }
                v25 -= 1e10
                fmt.Println(v4, v27, v25)
            }
        } else {
            v8 = (v8 << ((v5) & 63))
            var v103 int32 = -1
            var v104 int8 = v4
            _ = v103
            _ = v104
        }
        fmt.Println(v9, v8)
    } else {
        v105 := GenS7{F0: 867, F1: 1}
        for fi106 := 0; fi106 < 13; fi106++ {
            if false {
                for fi107 := 0; fi107 < 11; fi107++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    fmt.Println(v4)
                    _ = fi107
                }
                var v108 int = (fi106 & 701)
                _ = len(v2)
                _ = v108
            } else {
                var v109 float32 = (((0.0 * 0.0) * v25) * v25)
                v110 := GenS8{F0: 2, F1: 32767, F2: v5}
                _ = v109
                _ = v110
            }
            _ = fi106
        }
        _ = v105
    }
    if (!true) {
        var v111 int32 = int32(v9)
        var v112 uint16 = 0
        v111 = (v111 << (((v5 | (^(v5)))) & 63))
        _ = v111
        _ = v112
    } else {
        var v113 uint64 = (v9 << ((v5) & 63))
        v2[((("hello" + v27) + ((v27 + v27) + ("foo bar" + (((v27 + ((v27 + (v27 + "\n\t")) + v27)) + v27) + "")))) + "")] = 103
        _ = v113
    }
    fmt.Println(v9, v27)
    v1 += v1
    v114 := GenS9{F0: (v27 + v27), F1: 610, F2: v27}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v114
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
