package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 string; F1 int16; F2 uint }
type GenS3 struct { F0 uint16; F1 int32; F2 int8 }
type GenS4 struct { F0 uint; F1 uint }
type GenS5 struct { F0 int; F1 int16; F2 string }
type GenS6 struct { F0 int16; F1 string; F2 string }
type GenS7 struct { F0 uint16; F1 int16; F2 bool }
type GenS8 struct { F0 bool; F1 string; F2 string }

func main() {
    var v1 int64 = 403
    var v2 uint = 558
    var v3 string = ("hello" + ("a" + ("" + "a")))
    v2 += 4294967295
    for fi4 := 0; fi4 < 2; fi4++ {
        v5 := GenS1{F0: false, F1: 42}
        for fi6 := 0; fi6 < 6; fi6++ {
            v7 := GenS1{F0: (!((v2 > v2) || (("hello" + "") > v3))), F1: uint64(v2)}
            v8 := GenS2{F0: ("a" + "\n\t"), F1: 726, F2: 186}
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = fi4
        _ = v5
    }
    v1 -= v1
    if true {
        fmt.Println(v1, v2)
        if false {
            if false {
                fmt.Println(v3, v2, v1)
            }
        }
        var v9 uint32 = 0
        _ = v9
    } else {
        if false {
            var v10 uint16 = 42
            if true {
                var v11 float32 = -2.25
                fmt.Println(v11, v1, v1)
                var v12 float32 = v11
                _ = v11
                _ = v12
            }
            if true {
                for fi13 := 0; fi13 < 11; fi13++ {
                    v2 -= (v2 & (v2 - 1))
                    fmt.Println(v10)
                    _ = fi13
                }
                var v14 float64 = ((0.0 + ((1e10 * (((0.1 / float64(3.14159)) + 0.0) + 1e10)) - 0.0)) + -2.25)
                v14 -= v14
                _ = v14
            } else {
                if ((-2.25 + float32(float64((0.0 + float32(1.5))))) < 0.0) {
                    if false {
                        v3 = v3
                        v15 := GenS1{F0: (false || false), F1: 4294967295}
                        _ = v15
                    }
                }
                for fi16 := 0; fi16 < 7; fi16++ {
                    v17 := []int{fi16, 127, (fi16 + 403), fi16}
                    fmt.Println(v2, v2)
                    _ = fi16
                    _ = v17
                }
            }
            _ = v10
        }
        if false {
            v2 -= (v2 + 42)
            for fi18 := 0; fi18 < 10; fi18++ {
                var v19 uint64 = 1
                switch fi18 {
                case -584:
                    if true {
                        if (false || ((v19 >> ((18446744073709551615) & 63)) != v19)) {
                            for fi20 := 0; fi20 < 17; fi20++ {
                                var v21 float64 = 3.14159
                                fmt.Println(v19)
                                v2 += (^(v2))
                                v3 = v3
                                if true {
                                    if (false && false) {
                                        var v22 string = (v3 + "a")
                                        _ = v22
                                    } else {
                                        if false {
                                            var v23 string = "a"
                                            _ = v23
                                        } else {
                                            var v24 uint = (v2 & 255)
                                            _ = v24
                                        }
                                        var v25 uint32 = 2
                                        _ = v25
                                    }
                                    if (v21 > v21) {
                                        var v26 uint = v2
                                        v26 = v2
                                        _ = v26
                                    }
                                } else {
                                    var v27 int8 = 92
                                    _ = v27
                                }
                                _ = fi20
                                _ = v21
                            }
                        }
                        v28 := GenS2{F0: v3, F1: 42, F2: (v2 ^ (v2 * v2))}
                        _ = v28
                    }
                }
                switch 1 {
                case 128:
                    switch v3 {
                    case "a":
                        switch v3 {
                        case "":
                            v29 := GenS3{F0: 42, F1: -1, F2: int8(v19)}
                            _ = v29
                        case "a":
                            v30 := []int{(fi18 % (((fi18 ^ -932)) | 1)), -9223372036854775808}
                            _ = v30
                        default:
                            v31 := []int{4294967295, (fi18 << (((v2 / (((v2 >> ((v2) & 63))) | 1))) & 63))}
                            _ = v31
                        }
                    case "foo bar":
                        fmt.Println(v3, v3, v19)
                    default:
                        if false {
                            if (1.5 < ((1.5 - 1.5) - 1.5)) {
                                var v32 float32 = -2.25
                                if (v2 <= (v2 + v2)) {
                                    fmt.Println(v1, v19, fi18)
                                } else {
                                    v33 := []int{}
                                    for fi34 := 0; fi34 < 15; fi34++ {
                                        switch ((("\n\t" + ("" + ((v3 + "foo bar") + ("a" + v3)))) + v3) + v3) {
                                        case "":
                                            v19 = v19
                                        case "\n\t":
                                            switch fi18 {
                                            case -2147483648:
                                                v32 = (v32 - 3.14159)
                                            default:
                                                v32 = 0.1
                                            }
                                        default:
                                            var v35 string = v3
                                            if ((float64(-2.25) + (float64(0.0) - float64((v32 / float32(3.14159))))) == 3.14159) {
                                                v32 = 0.0
                                                v36 := GenS2{F0: v3, F1: 32767, F2: v2}
                                                if false {
                                                    var v37 int8 = 7
                                                    _ = v37
                                                } else {
                                                    v19 = (v19 % ((v19) | 1))
                                                    var v38 int32 = 42
                                                    _ = v38
                                                }
                                                _ = v36
                                            }
                                            _ = v35
                                        }
                                        _ = fi34
                                    }
                                    var v39 uint8 = 184
                                    _ = v33
                                    _ = v39
                                }
                                _ = v32
                            } else {
                                v40 := []int{-1, (fi18 << ((v2) & 63)), fi18, fi18}
                                switch "hello" {
                                case "hello":
                                    v40 = append(v40, (fi18 | fi18))
                                case "a":
                                    v41 := GenS3{F0: uint16(v19), F1: -2147483648, F2: -128}
                                    _ = v41
                                default:
                                    if false {
                                        v40 = append(v40, (fi18 * fi18))
                                    } else {
                                        var v42 uint32 = 42
                                        v1 = 1
                                        for fi43 := 0; fi43 < 13; fi43++ {
                                            v44 := GenS2{F0: "hello", F1: -32768, F2: 18446744073709551615}
                                            var v45 string = v3
                                            _ = fi43
                                            _ = v44
                                            _ = v45
                                        }
                                        _ = v42
                                    }
                                }
                                _ = v40
                            }
                        } else {
                            fmt.Println(fi18, v1, v1)
                            for fi46 := 0; fi46 < 19; fi46++ {
                                v47 := []int{fi18, fi18, fi46, int(v1)}
                                _ = len(v47)
                                _ = fi46
                                _ = v47
                            }
                            fmt.Println(v3)
                        }
                    }
                case 2147483647:
                    fmt.Println(v3, v3, fi18)
                case 601:
                    var v48 uint16 = uint16(v19)
                    _ = v48
                }
                _ = fi18
                _ = v19
            }
        }
    }
    v49 := GenS2{F0: v3, F1: -1, F2: (v2 % (((^(v2))) | 1))}
    v50 := []int{}
    var v51 bool = true
    var v52 int16 = -666
    for fi53 := 0; fi53 < 7; fi53++ {
        if v51 {
            var v54 uint = v2
            v2 = (v54 | (v54 ^ (v54 - v54)))
            _ = v54
        } else {
            if true {
                if (true || (212 < 128)) {
                    for fi55 := 0; fi55 < 14; fi55++ {
                        var v56 uint8 = 42
                        v57 := GenS4{F0: v2, F1: v2}
                        fmt.Println(v3, v51)
                        _ = fi55
                        _ = v56
                        _ = v57
                    }
                    v52 = 1
                    v58 := GenS5{F0: fi53, F1: v52, F2: v3}
                    _ = v58
                }
                var v59 int32 = 0
                _ = v59
            }
        }
        _ = fi53
    }
    switch v3 {
    case "":
        var v60 int16 = -521
        _ = v60
    case "foo bar":
        if v51 {
            for fi61 := 0; fi61 < 17; fi61++ {
                var v62 bool = (-623 < (^(v52)))
                for fi63 := 0; fi63 < 16; fi63++ {
                    v64 := GenS6{F0: v52, F1: v3, F2: (v3 + v3)}
                    v65 := GenS2{F0: "hello", F1: (^(v52)), F2: v2}
                    switch (v3 + ((v3 + "hello") + v3)) {
                    case "foo bar":
                        var v66 int32 = -2147483648
                        _ = v66
                    case "a":
                        fmt.Println(v1, fi61, v3)
                    case "\n\t":
                        var v67 int16 = 1000
                        _ = v67
                    default:
                        var v68 uint16 = 32767
                        _ = v68
                    }
                    _ = fi63
                    _ = v64
                    _ = v65
                }
                v69 := GenS7{F0: 42, F1: (-(v52)), F2: (!(v51 || false))}
                _ = fi61
                _ = v62
                _ = v69
            }
            if (v51 && (v2 >= 42)) {
                fmt.Println(v3, v1, v51)
                v70 := GenS5{F0: -181, F1: v52, F2: "foo bar"}
                _ = v70
            }
            for fi71 := 0; fi71 < 2; fi71++ {
                v72 := GenS6{F0: -1, F1: v3, F2: "foo bar"}
                _ = fi71
                _ = v72
            }
        }
    }
    var v73 string = ""
    var v74 uint32 = 4294967295
    for fi75 := 0; fi75 < 11; fi75++ {
        v76 := GenS4{F0: v2, F1: v2}
        v52 = v52
        _ = fi75
        _ = v76
    }
    if (v2 == v2) {
        for fi77 := 0; fi77 < 19; fi77++ {
            switch (fi77 % ((32767) | 1)) {
            case -9223372036854775808:
                if v51 {
                    v78 := GenS8{F0: (255 >= 412), F1: v73, F2: "\n\t"}
                    var v79 uint = v2
                    var v80 uint = 453
                    _ = v78
                    _ = v79
                    _ = v80
                } else {
                    var v81 float64 = 1e10
                    var v82 uint64 = 42
                    v51 = v51
                    _ = v81
                    _ = v82
                }
            case -869:
                for fi83 := 0; fi83 < 9; fi83++ {
                    fmt.Println(v52, fi83)
                    _ = fi83
                }
            default:
                var v84 float64 = 1.5
                _ = v84
            }
            _ = fi77
        }
        if len(v50) >= 2 {
            v50 = v50[1:len(v50)]
        }
    } else {
        for fi85 := 0; fi85 < 5; fi85++ {
            var v86 int = fi85
            _ = fi85
            _ = v86
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v50:", len(v50), v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v49
    _ = v50
    _ = v51
    _ = v52
    _ = v73
    _ = v74
}
