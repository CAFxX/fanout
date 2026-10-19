package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 string; F1 int16; F2 uint }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint; F1 int32; F2 uint64 }
type GenS6 struct { F0 uint64; F1 string; F2 bool }
type GenS7 struct { F0 string }
type GenS8 struct { F0 bool; F1 string }
type GenS9 struct { F0 int64; F1 int64; F2 uint32 }
func genFn0(p0 bool, p1 uint8, p2 int) int64 {
    var v5 bool = (3.14159 < 1e10)
    if true {
        v6 := []int{}
        if true {
            if len(v6) > 0 {
                v6[0] = p2
            }
            switch p2 {
            case -2147483648:
                var v7 int8 = 0
                _ = v7
            case 4294967295:
                switch ("foo bar" + ("hello" + "foo bar")) {
                case "a":
                    var v8 uint32 = 837
                    _ = v8
                }
            }
            if p0 {
                v9 := GenS1{F0: 2147483647}
                var v10 int64 = 784
                _ = v10
                _ = v9
            }
        } else {
            v11 := GenS2{F0: 65535, F1: "\n\t"}
            _ = len(v6)
            _ = v11
        }
        _ = v6
    }
    v12 := GenS1{F0: uint32(p1)}
    p2 = (p2 + 64)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v5
    return 109
}
func genFn1(p0 int8) string {
    var v13 bool = true
    if v13 {
        v13 = v13
        var v14 int16 = 0
        _ = v14
    } else {
        var v15 int = -838
        v16 := []int{v15, v15, (v15 >> ((314) & 63))}
        for fi17 := 0; fi17 < 11; fi17++ {
            v18 := GenS3{F0: (("\n\t" + (("foo bar" + "a") + "\n\t")) + "a"), F1: -165, F2: 2}
            fmt.Println(v13, fi17)
            if (559 >= 18446744073709551615) {
                var v19 int32 = -1
                _ = v19
            } else {
                var v20 float64 = 3.14159
                fmt.Println(v20, p0)
                if false {
                    if ("\n\t" == ((("foo bar" + "") + "\n\t") + "\n\t")) {
                        var v21 uint32 = 634
                        var v22 int16 = 0
                        _ = v21
                        _ = v22
                    }
                    switch "a" {
                    case "":
                        for fi23 := 0; fi23 < 14; fi23++ {
                            v24 := GenS3{F0: ("" + ("hello" + "\n\t")), F1: -1, F2: 0}
                            v25 := GenS1{F0: 4294967295}
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                    }
                    p0 = (p0 ^ (p0 + p0))
                }
                _ = v20
            }
            _ = fi17
            _ = v18
        }
        _ = v15
        _ = v16
    }
    switch 32768 {
    case 1:
        v26 := GenS4{F0: 2}
        _ = v26
    default:
        p0 = p0
    }
    if v13 {
        var v27 uint32 = 2
        _ = v27
    } else {
        for fi28 := 0; fi28 < 1; fi28++ {
            v13 = (((-2.25 + 3.14159) > 1e10) || (392 != 803))
            v29 := []int{}
            _ = fi28
            _ = v29
        }
    }
    for fi30 := 0; fi30 < 9; fi30++ {
        v31 := GenS5{F0: 42, F1: -2147483648, F2: 153}
        _ = fi30
        _ = v31
    }
    _ = p0
    _ = v13
    return ("\n\t" + "a")
}

func main() {
    v1 := []int{1, 42, -1, -1}
    var v2 float64 = 3.14159
    var v3 string = "foo bar"
    var v4 int8 = -16
    for fi32 := 0; fi32 < 16; fi32++ {
        v1 = append(v1, (fi32 << ((2) & 63)))
        var v33 string = v3
        _ = fi32
        _ = v33
    }
    v34 := GenS5{F0: 0, F1: -743, F2: 321}
    var v35 string = "foo bar"
    if (v3 >= v35) {
        var v36 uint8 = 127
        for fi37 := 0; fi37 < 7; fi37++ {
            if false {
                v35 = (("hello" + ("hello" + "a")) + v35)
            }
            var v38 int8 = v4
            switch v3 {
            case "a":
                var v39 int64 = 42
                _ = v39
            case "foo bar":
                if (!false) {
                    switch (^(fi37)) {
                    case 0:
                        switch ("\n\t" + v3) {
                        case "hello":
                            fmt.Println(v4)
                        case "\n\t":
                            var v40 uint64 = 870
                            _ = v40
                        default:
                            v35 = v35
                        }
                    case -9223372036854775808:
                        var v41 int32 = 0
                        _ = v41
                    }
                    v3 = ("" + ((("a" + "") + (v35 + v35)) + (v35 + ((((v3 + v3) + (("foo bar" + (v35 + v3)) + "foo bar")) + "\n\t") + "a"))))
                }
            case "\n\t":
                if false {
                    var v42 string = v3
                    _ = v42
                }
            default:
                var v43 bool = true
                _ = v43
            }
            _ = fi37
            _ = v38
        }
        _ = v36
    } else {
        for fi44 := 0; fi44 < 17; fi44++ {
            v45 := GenS6{F0: 2, F1: v3, F2: (1 != 42)}
            if (true || (!true)) {
                var v46 bool = false
                var v47 uint64 = 18446744073709551615
                var v48 uint8 = 127
                _ = v46
                _ = v47
                _ = v48
            } else {
                if true {
                    v49 := GenS4{F0: 164}
                    _ = v49
                }
                if ((248 < -2147483648) && ((v2 * (v2 / 1.0)) != (v2 + (v2 * ((v2 + 1e10) - 1.5))))) {
                    if true {
                        if false {
                            if false {
                                var v50 int32 = -1
                                _ = v50
                            }
                            fmt.Println(v2, v35)
                        } else {
                            switch "hello" {
                            case "hello":
                                var v51 uint16 = 1
                                _ = v51
                            case "a":
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            default:
                                v2 += 1e10
                            }
                            var v52 string = "foo bar"
                            for fi53 := 0; fi53 < 19; fi53++ {
                                v35 = v35
                                if len(v1) > 0 {
                                    v1[0] = (fi53 ^ -36)
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = fi53
                            }
                            _ = v52
                        }
                    } else {
                        switch fi44 {
                        case 128:
                            for fi54 := 0; fi54 < 16; fi54++ {
                                fmt.Println(v3)
                                if ((-128 > (^(v4))) || false) {
                                    var v55 uint = 4294967295
                                    _ = v55
                                } else {
                                    v56 := GenS5{F0: 829, F1: 1, F2: 42}
                                    _ = len(v1)
                                    v1 = append(v1, -9223372036854775808)
                                    _ = v56
                                }
                                _ = fi54
                            }
                        }
                    }
                    v2 -= ((((3.14159 - (v2 * 0.1)) + -2.25) * (v2 / 0.1)) * ((float64(0.1) * v2) * float64(1e10)))
                    fmt.Println(v4, v3, v3)
                } else {
                    var v57 uint16 = 0
                    var v58 uint8 = 255
                    if len(v1) > 0 {
                        v1[0] = fi44
                    }
                    _ = v57
                    _ = v58
                }
                var v59 uint32 = 2147483648
                _ = v59
            }
            var v60 int = fi44
            _ = fi44
            _ = v45
            _ = v60
        }
    }
    for fi61 := 0; fi61 < 1; fi61++ {
        v2 = (v2 + float64(1.5))
        _ = fi61
    }
    _ = len(v1)
    if true {
        v4 += v4
    }
    var v62 bool = false
    for fi63 := 0; fi63 < 17; fi63++ {
        var v64 int64 = 63
        switch 708 {
        case 2147483647:
            if ((!true) && v62) {
                if ((0.0 - ((float32(v2) + 0.0) / 0.1)) == 1e10) {
                    if v62 {
                        fmt.Println(v35, v3)
                        var v65 uint = 0
                        if true {
                            v66 := GenS5{F0: 431, F1: 174, F2: 1}
                            _ = v66
                        } else {
                            if (false || false) {
                                if (30 != (^(v4))) {
                                    var v67 int16 = 1000
                                    _ = v67
                                }
                                v64 = (v64 << ((v65) & 63))
                                var v68 uint8 = 182
                                _ = v68
                            } else {
                                v69 := GenS6{F0: 18446744073709551615, F1: v35, F2: ((fi63 - (fi63 % ((-9223372036854775808) | 1))) > 127)}
                                v70 := GenS7{F0: "\n\t"}
                                var v71 int = (fi63 * fi63)
                                _ = v69
                                _ = v70
                                _ = v71
                            }
                            var v72 uint16 = 32767
                            _ = v72
                        }
                        _ = v65
                    } else {
                        switch "a" {
                        case "\n\t":
                            v2 = (float64(-2.25) - v2)
                        case "a":
                            if (v62 || false) {
                                for fi73 := 0; fi73 < 19; fi73++ {
                                    if len(v1) > 0 {
                                        v1[0] = fi63
                                    }
                                    switch (fi73 % ((-149) | 1)) {
                                    case 42:
                                        if (0 >= 0) {
                                            if v62 {
                                                v4 = (v4 % (((v4 & (^(v4)))) | 1))
                                                fmt.Println(v4)
                                                v1 = append(v1, 4294967295)
                                                switch (v3 + v3) {
                                                case "\n\t":
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                }
                                            }
                                            var v74 uint = 4294967295
                                            var v75 uint = v74
                                            _ = v74
                                            _ = v75
                                        } else {
                                            _ = len(v1)
                                        }
                                        _ = len(v1)
                                    case -781:
                                        if v62 {
                                            if len(v1) > 0 {
                                                v1[0] = fi63
                                            }
                                        }
                                    case 128:
                                        v76 := GenS8{F0: (((v3 + "a") + v3) != ("hello" + ("foo bar" + v35))), F1: (v35 + v35)}
                                        v3 = (("\n\t" + (v35 + "a")) + "hello")
                                        _ = v76
                                    default:
                                        v3 = (("foo bar" + "hello") + ("foo bar" + "foo bar"))
                                    }
                                    var v77 bool = (0 != 983)
                                    switch fi63 {
                                    case 127:
                                        if (!true) {
                                            fmt.Println(v3, fi73, v3)
                                            if (v62 && (((4294967295 < 18446744073709551615) || (v77 && (v2 >= v2))) && (0 <= 18446744073709551615))) {
                                                var v78 string = v35
                                                if false {
                                                    fmt.Println(v35, v77)
                                                    fmt.Println(v64)
                                                    v35 = ""
                                                } else {
                                                    if len(v1) > 0 {
                                                        v1[0] = fi63
                                                    }
                                                    if v62 {
                                                        v1 = append(v1, (-(fi63)))
                                                    }
                                                }
                                                v2 += ((v2 / (0.1 * v2)) * (v2 / v2))
                                                fmt.Println(v62, v78, v4)
                                                _ = v78
                                            }
                                            fmt.Println(v62, v35)
                                        }
                                    case -9223372036854775808:
                                        var v79 string = v3
                                        _ = v79
                                    default:
                                        var v80 int32 = 1
                                        _ = v80
                                    }
                                    _ = fi73
                                    _ = v77
                                }
                            }
                        }
                    }
                    _ = len(v1)
                    v3 = ("foo bar" + v3)
                }
                var v81 int16 = -1
                v82 := GenS3{F0: v3, F1: v81, F2: 0}
                _ = v81
                _ = v82
            } else {
                for fi83 := 0; fi83 < 17; fi83++ {
                    var v84 uint = 2
                    v85 := GenS1{F0: 0}
                    v86 := GenS5{F0: v84, F1: -614, F2: 0}
                    _ = fi83
                    _ = v84
                    _ = v85
                    _ = v86
                }
                v87 := GenS9{F0: v64, F1: (v64 & genFn0((v4 >= 58), 0, (fi63 | fi63))), F2: 474}
                _ = v87
            }
        case 127:
            v88 := GenS7{F0: ""}
            _ = v88
        default:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        _ = fi63
        _ = v64
    }
    var v89 int16 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v62:", v62)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v62
    _ = v89
}
