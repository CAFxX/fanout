package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 int8; F2 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint; F1 uint8; F2 uint64 }
type GenS5 struct { F0 bool; F1 int }
type GenS6 struct { F0 uint32; F1 int64; F2 bool }
type GenS7 struct { F0 uint16; F1 int64; F2 bool }
type GenS8 struct { F0 uint64; F1 uint32; F2 uint16 }
type GenS9 struct { F0 int64; F1 int32 }
type GenS10 struct { F0 uint16 }

func main() {
    var v1 bool = (119 >= -2147483648)
    v2 := []string{"\n\t", "hello", ("foo bar" + "hello"), "a"}
    var v3 string = "hello"
    var v4 uint16 = 65535
    v5 := []int{32768, -111, 2147483647}
    switch "hello" {
    case "a":
        for fi6 := 0; fi6 < 2; fi6++ {
            var v7 int8 = -107
            switch "\n\t" {
            case "a":
                var v8 uint8 = 42
                _ = v8
            case "hello":
                fmt.Println(fi6, v3, v7)
            default:
                var v9 float64 = (0.1 + (1e10 * 0.1))
                _ = v9
            }
            for fi10 := 0; fi10 < 11; fi10++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi10
            }
            _ = fi6
            _ = v7
        }
    case "hello":
        switch 4294967295 {
        case 0:
            if v1 {
                v1 = v1
                for fi11 := 0; fi11 < 20; fi11++ {
                    var v12 int32 = -787
                    _ = fi11
                    _ = v12
                }
            } else {
                var v13 string = (v3 + v3)
                v14 := GenS1{F0: 870}
                _ = v13
                _ = v14
            }
        }
    default:
        if true {
            fmt.Println(v1)
        } else {
            var v15 uint = 2
            v16 := GenS2{F0: 42, F1: -3, F2: (^(v15))}
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v15
            _ = v16
        }
    }
    var v17 int8 = -1
    v3 = (v3 + (v3 + (("a" + v3) + v3)))
    var v18 int = 1
    for fi19 := 0; fi19 < 3; fi19++ {
        for fi20 := 0; fi20 < 5; fi20++ {
            v3 = ""
            _ = fi20
        }
        fmt.Println(v17, v4, v1)
        switch "foo bar" {
        case "":
            switch (fi19 % ((-458) | 1)) {
            case 0:
                v3 = v3
            case -128:
                v2 = append(v2, v3)
            case 128:
                for fi21 := 0; fi21 < 10; fi21++ {
                    if (2 >= 0) {
                        if v1 {
                            if len(v2) > 0 {
                                v2[0] = "hello"
                            }
                            v22 := GenS1{F0: v4}
                            v2 = append(v2, (v3 + (((((v3 + ((("hello" + (v3 + "hello")) + v3) + (((v3 + v3) + v3) + ((v3 + v3) + ((("hello" + ("foo bar" + (v3 + "a"))) + (v3 + "foo bar")) + "foo bar"))))) + (((v3 + "") + (v3 + "a")) + ("" + (("foo bar" + (v3 + v3)) + v3)))) + v3) + ("a" + v3)) + v3)))
                            _ = v22
                        }
                        v23 := GenS1{F0: (^(v4))}
                        var v24 int16 = 1
                        _ = v23
                        _ = v24
                    }
                    v2 = append(v2, (((((v3 + "") + (v3 + v3)) + v3) + v3) + v3))
                    v2 = append(v2, "a")
                    _ = fi21
                }
            }
        case "foo bar":
            v3 = ("\n\t" + v3)
        case "\n\t":
            v17 = (v17 ^ (v17 ^ v17))
        }
        _ = fi19
    }
    var v25 uint8 = 0
    switch (v18 | (-(v18))) {
    case -893:
        var v26 int32 = 42
        _ = v26
    case 178:
        fmt.Println(v1, v1, v4)
    case 32767:
        var v27 float32 = 1.5
        _ = v27
    default:
        if ((0 > 42) || ((!v1) || false)) {
            var v28 uint8 = 127
            _ = v28
        } else {
            switch (v18 % (((v18 ^ 32767)) | 1)) {
            case -1:
                for fi29 := 0; fi29 < 9; fi29++ {
                    var v30 bool = v1
                    switch 42 {
                    case -2147483648:
                        for fi31 := 0; fi31 < 16; fi31++ {
                            switch "\n\t" {
                            case "foo bar":
                                _ = len(v5)
                            case "\n\t":
                                if ((false || v1) && false) {
                                    if (3.14159 != 0.0) {
                                        v32 := GenS3{F0: 1000}
                                        _ = len(v2)
                                        _ = len(v5)
                                        if v30 {
                                            v33 := GenS3{F0: 32767}
                                            switch (v3 + ("hello" + v3)) {
                                            case "\n\t":
                                                _ = len(v5)
                                            }
                                            _ = v33
                                        }
                                        _ = len(v5)
                                        v34 := GenS1{F0: v4}
                                        _ = v32
                                        _ = v34
                                    }
                                    _ = len(v5)
                                    if v1 {
                                        v3 = (v3 + "foo bar")
                                        if (483 < 985) {
                                            if (-208 < -1) {
                                                if len(v5) > 0 {
                                                    v5[0] = v18
                                                }
                                                switch 0 {
                                                case 752:
                                                    _ = len(v2)
                                                }
                                            } else {
                                                v35 := GenS4{F0: 958, F1: v25, F2: 424}
                                                switch v3 {
                                                case "foo bar":
                                                    if (!(false || false)) {
                                                        if v1 {
                                                            v25 = (v25 << ((18446744073709551615) & 63))
                                                            var v36 int8 = v17
                                                            _ = v36
                                                        } else {
                                                            var v37 int16 = 42
                                                            v37 = int16(v17)
                                                            switch v3 {
                                                            case "a":
                                                                var v38 uint16 = uint16(v25)
                                                                _ = v38
                                                            case "":
                                                                v39 := GenS5{F0: false, F1: (fi31 * (fi31 / (((fi31 ^ (v18 << ((255) & 63)))) | 1)))}
                                                                _ = v39
                                                            default:
                                                                v18 = 1
                                                            }
                                                            _ = v37
                                                        }
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                    } else {
                                                        v17 -= v17
                                                        if (0.0 >= 1.5) {
                                                            v4 = (v4 >> ((129) & 63))
                                                            if v30 {
                                                                var v40 uint16 = v4
                                                                v30 = true
                                                                _ = v40
                                                            } else {
                                                                v4 -= v4
                                                                switch v3 {
                                                                case "a":
                                                                    v2 = append(v2, v3)
                                                                case "":
                                                                    v41 := GenS6{F0: 1, F1: -9223372036854775808, F2: (-2147483648 != -2147483648)}
                                                                    _ = v41
                                                                default:
                                                                    var v42 float32 = 1e10
                                                                    _ = v42
                                                                }
                                                            }
                                                            v18 = v18
                                                        }
                                                    }
                                                default:
                                                    v43 := GenS7{F0: v4, F1: -9223372036854775808, F2: false}
                                                    if true {
                                                        switch v18 {
                                                        case -2147483648:
                                                            v4 = 136
                                                        }
                                                        v1 = (float64(float32(0.0)) < (1e10 / float64(0.1)))
                                                        if v1 {
                                                            var v44 uint16 = 1
                                                            if false {
                                                                var v45 int8 = (v17 << ((18446744073709551615) & 63))
                                                                _ = v45
                                                            }
                                                            _ = v44
                                                        }
                                                    }
                                                    _ = v43
                                                }
                                                _ = v35
                                            }
                                            v25 = v25
                                            if (-662 != -32768) {
                                                var v46 int32 = -2147483648
                                                switch (v3 + v3) {
                                                case "a":
                                                    fmt.Println(v25)
                                                }
                                                v30 = true
                                                var v47 uint16 = v4
                                                v17 += (v17 / (((-(v17))) | 1))
                                                v48 := GenS5{F0: v30, F1: (fi29 * v18)}
                                                if ((-2.25 * 1.5) < (-2.25 * 0.0)) {
                                                    v1 = true
                                                    v46 = v46
                                                    var v49 string = (v3 + v3)
                                                    if (!v30) {
                                                        if (3.14159 < 0.1) {
                                                            switch v49 {
                                                            case "hello":
                                                                switch v3 {
                                                                case "foo bar":
                                                                    v50 := GenS2{F0: 1, F1: (v17 + v17), F2: 2}
                                                                    _ = v50
                                                                }
                                                            case "foo bar":
                                                                fmt.Println(v3, v17)
                                                            default:
                                                                v25 = (^(v25))
                                                                var v51 uint16 = v4
                                                                _ = v51
                                                            }
                                                            var v52 bool = ((v3 + ((v3 + (((("\n\t" + v49) + (v49 + (((v3 + ("foo bar" + (v3 + "a"))) + ("\n\t" + v3)) + v49))) + ((("\n\t" + v3) + "") + v3)) + v49)) + v49)) > v49)
                                                            v25 -= (^(v25))
                                                            _ = v52
                                                        }
                                                        switch v18 {
                                                        case 32768:
                                                            if (2 == uint64(v25)) {
                                                                var v53 int16 = -313
                                                                if v30 {
                                                                    var v54 float32 = -2.25
                                                                    v55 := GenS8{F0: 4294967295, F1: 42, F2: 32767}
                                                                    _ = v54
                                                                    _ = v55
                                                                }
                                                                _ = v53
                                                            } else {
                                                                var v56 float32 = 3.14159
                                                                _ = v56
                                                            }
                                                        case -697:
                                                            if v1 {
                                                                if v30 {
                                                                    _ = len(v5)
                                                                    v57 := GenS6{F0: 1, F1: 9223372036854775807, F2: v1}
                                                                    _ = v57
                                                                }
                                                                var v58 float32 = 1.5
                                                                var v59 uint16 = (^(v47))
                                                                _ = v58
                                                                _ = v59
                                                            }
                                                        case 9223372036854775807:
                                                            fmt.Println(fi31, v25, v4)
                                                        }
                                                        switch (v18 / ((fi31) | 1)) {
                                                        case 1:
                                                            var v60 int8 = (v17 * v17)
                                                            _ = v60
                                                        }
                                                    }
                                                    var v61 bool = ("hello" <= v3)
                                                    _ = v49
                                                    _ = v61
                                                } else {
                                                    if v1 {
                                                        if v30 {
                                                            switch (-(v18)) {
                                                            case -713:
                                                                v17 = v17
                                                            case 128:
                                                                var v62 uint32 = 2147483647
                                                                _ = v62
                                                            case 0:
                                                                var v63 int16 = -32768
                                                                _ = v63
                                                            }
                                                            v17 = (v17 ^ int8(v46))
                                                        }
                                                    }
                                                }
                                                var v64 uint16 = v47
                                                v65 := GenS9{F0: -1, F1: 1}
                                                _ = v46
                                                _ = v47
                                                _ = v48
                                                _ = v64
                                                _ = v65
                                            }
                                        } else {
                                            if len(v5) > 0 {
                                                v5[0] = 2147483647
                                            }
                                        }
                                        var v66 int64 = -9223372036854775808
                                        var v67 bool = ((v66 - (v66 % (((v66 | -1)) | 1))) < (v66 % (((v66 ^ v66)) | 1)))
                                        _ = v66
                                        _ = v67
                                    } else {
                                        v25 = 255
                                        v68 := GenS2{F0: (v17 | 0), F1: (v17 / ((v17) | 1)), F2: 255}
                                        _ = v68
                                    }
                                }
                            case "a":
                                var v69 uint = 2
                                _ = v69
                            default:
                                if (!true) {
                                    var v70 uint = 1000
                                    v25 = v25
                                    v5 = append(v5, (v18 ^ 0))
                                    _ = v70
                                }
                            }
                            switch "foo bar" {
                            case "hello":
                                if (uint(v4) >= 4294967295) {
                                    v71 := GenS8{F0: 4294967295, F1: uint32(v4), F2: 255}
                                    _ = v71
                                } else {
                                    v4 = 174
                                    if len(v5) > 0 {
                                        v5[0] = 0
                                    }
                                    var v72 uint8 = v25
                                    _ = v72
                                }
                            default:
                                var v73 uint8 = (v25 >> ((0) & 63))
                                _ = v73
                            }
                            var v74 uint64 = 999
                            _ = fi31
                            _ = v74
                        }
                    default:
                        for fi75 := 0; fi75 < 16; fi75++ {
                            var v76 string = ((v3 + v3) + (v3 + (v3 + v3)))
                            if true {
                                v1 = ((0.0 * 3.14159) >= 1.5)
                            } else {
                                v4 += (v4 + (v4 | (v4 | v4)))
                                switch fi75 {
                                case 0:
                                    v77 := GenS4{F0: 42, F1: (v25 % (((v25 % (((^(v25))) | 1))) | 1)), F2: 18446744073709551615}
                                    _ = v77
                                case -990:
                                    var v78 int32 = -1
                                    _ = v78
                                default:
                                    switch (fi29 + (fi29 - (^(v18)))) {
                                    case 32767:
                                        fmt.Println(v30, fi29, v3)
                                    case 42:
                                        var v79 float32 = 1.5
                                        var v80 int32 = 1
                                        if v30 {
                                            v81 := GenS5{F0: (!v1), F1: fi75}
                                            v82 := GenS10{F0: (v4 ^ (v4 & v4))}
                                            if (1 != 847) {
                                                fmt.Println(v79)
                                                if v30 {
                                                    _ = len(v2)
                                                }
                                            } else {
                                                var v83 int8 = v17
                                                switch fi75 {
                                                case 4294967295:
                                                    if true {
                                                        v80 = v80
                                                    } else {
                                                        var v84 int16 = 827
                                                        v25 -= 128
                                                        _ = v84
                                                    }
                                                }
                                                v18 = (fi75 << ((4294967295) & 63))
                                                _ = v83
                                            }
                                            _ = len(v2)
                                            v80 = 962
                                            _ = v81
                                            _ = v82
                                        } else {
                                            v1 = (v30 || (v1 || (v79 > v79)))
                                            fmt.Println(v4)
                                        }
                                        var v85 int16 = 0
                                        var v86 float64 = (1.5 / 0.1)
                                        v87 := GenS2{F0: int8(v80), F1: -128, F2: 376}
                                        _ = v79
                                        _ = v80
                                        _ = v85
                                        _ = v86
                                        _ = v87
                                    default:
                                        var v88 int8 = v17
                                        _ = v88
                                    }
                                }
                                v5 = append(v5, v18)
                            }
                            var v89 int64 = 1
                            _ = fi75
                            _ = v76
                            _ = v89
                        }
                    }
                    _ = fi29
                    _ = v30
                }
            case 0:
                fmt.Println(v18, v4)
            }
        }
    }
    v4 = (v4 | (v4 << ((2) & 63)))
    v90 := GenS7{F0: 255, F1: 9223372036854775807, F2: (true || (true && v1))}
    var v91 int16 = -977
    var v92 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v91:", v91)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v90
    _ = v91
    _ = v92
}
