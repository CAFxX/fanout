package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 string; F1 int; F2 int16 }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int8; F1 uint64 }
type GenS10 struct { F0 string; F1 uint32 }
type GenS11 struct { F0 int8; F1 int8 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]bool)
    var v3 uint32 = 922
    v4 := GenS1{F0: 2, F1: 450}
    switch "hello" {
    case "a":
        if true {
            var v5 uint16 = 0
            switch -18 {
            case 437:
                for fi6 := 0; fi6 < 6; fi6++ {
                    var v7 uint = 1
                    if false {
                        v8 := GenS2{F0: "foo bar", F1: fi6, F2: 0}
                        _ = v8
                    }
                    fmt.Println(v3, v3)
                    _ = fi6
                    _ = v7
                }
            case 128:
                v3 -= v3
            default:
                for fi9 := 0; fi9 < 12; fi9++ {
                    v10 := GenS3{F0: (fi9 >> ((881) & 63)), F1: 740}
                    _ = fi9
                    _ = v10
                }
            }
            _ = len(v1)
            _ = v5
        } else {
            v11 := GenS1{F0: 32767, F1: -547}
            if (v3 == v3) {
                for fi12 := 0; fi12 < 17; fi12++ {
                    if (v3 != (v3 ^ v3)) {
                        v13 := []int{fi12, fi12}
                        _ = v13
                    } else {
                        v3 = v3
                        v14 := v2[(fi12 | -237)]
                        _ = v14
                    }
                    v3 = (v3 * (^(v3)))
                    v3 = 65535
                    _ = fi12
                }
            }
            if false {
                v15 := []int{-968, -457}
                if false {
                    for fi16 := 0; fi16 < 11; fi16++ {
                        v1[""] = "a"
                        var v17 int64 = 9223372036854775807
                        _ = fi16
                        _ = v17
                    }
                    _ = len(v2)
                } else {
                    v18 := GenS4{F0: 557}
                    switch 4294967295 {
                    case 0:
                        v3 = v3
                    case -779:
                        if (!false) {
                            v19 := GenS3{F0: 369, F1: 18446744073709551615}
                            for fi20 := 0; fi20 < 11; fi20++ {
                                var v21 string = ""
                                v22 := GenS2{F0: v21, F1: fi20, F2: -1}
                                _ = fi20
                                _ = v21
                                _ = v22
                            }
                            if false {
                                switch "" {
                                case "":
                                    var v23 int16 = 1
                                    _ = v23
                                case "hello":
                                    for fi24 := 0; fi24 < 19; fi24++ {
                                        for fi25 := 0; fi25 < 11; fi25++ {
                                            v26 := GenS5{F0: 953}
                                            if (1 > -1) {
                                                if ((fi25 ^ fi24) >= (fi25 - -227)) {
                                                    var v27 string = "a"
                                                    if true {
                                                        var v28 uint = 0
                                                        _ = v28
                                                    }
                                                    _ = v27
                                                } else {
                                                    if (-2147483648 >= -497) {
                                                        v29 := GenS5{F0: 42}
                                                        var v30 float64 = (float64(0.0) / -2.25)
                                                        v3 = v3
                                                        _ = v29
                                                        _ = v30
                                                    } else {
                                                        _ = len(v15)
                                                        v3 -= (v3 | 2147483648)
                                                        var v31 int32 = -1
                                                        _ = v31
                                                    }
                                                    var v32 int32 = 885
                                                    v32 = 2147483647
                                                    _ = v32
                                                }
                                                v33 := GenS1{F0: 255, F1: fi25}
                                                v3 += v3
                                                if true {
                                                    var v34 int = (fi25 / (((fi24 / ((fi25) | 1))) | 1))
                                                    _ = len(v2)
                                                    v35 := GenS6{F0: 100}
                                                    _ = v34
                                                    _ = v35
                                                }
                                                v3 += (v3 - (v3 & v3))
                                                _ = v33
                                            }
                                            _ = fi25
                                            _ = v26
                                        }
                                        for fi36 := 0; fi36 < 3; fi36++ {
                                            var v37 uint64 = 18446744073709551615
                                            var v38 int = (fi24 / (((fi24 + fi36)) | 1))
                                            _ = fi36
                                            _ = v37
                                            _ = v38
                                        }
                                        delete(v1, "a")
                                        _ = fi24
                                    }
                                case "a":
                                    switch (("a" + ("" + "\n\t")) + ((("a" + ("a" + "foo bar")) + "\n\t") + ("" + ("" + ("\n\t" + ""))))) {
                                    case "a":
                                        if (false && (-38 > -2147483648)) {
                                            if true {
                                                var v39 int = 42
                                                _ = v39
                                            }
                                            v3 += v3
                                            if (("hello" + "foo bar") < "") {
                                                v40 := GenS1{F0: 2, F1: 0}
                                                v3 -= v3
                                                v3 = v3
                                                _ = v40
                                            } else {
                                                if (42 <= 42) {
                                                    v41 := GenS4{F0: 382}
                                                    _ = v41
                                                }
                                            }
                                        }
                                    }
                                }
                                v3 += v3
                            } else {
                                v42 := GenS7{F0: 16}
                                for fi43 := 0; fi43 < 2; fi43++ {
                                    if len(v15) > 0 {
                                        v15[0] = -1
                                    }
                                    _ = fi43
                                }
                                for fi44 := 0; fi44 < 7; fi44++ {
                                    if (!(388 == 194)) {
                                        if true {
                                            v45 := v2[fi44]
                                            _ = v45
                                        } else {
                                            fmt.Println(fi44)
                                            fmt.Println(fi44)
                                        }
                                        var v46 float64 = (((-2.25 - -2.25) / 0.1) + (0.0 - -2.25))
                                        v3 += (v3 >> ((2) & 63))
                                        _ = v46
                                    }
                                    if ((v3 + v3) <= v3) {
                                        var v47 int16 = int16(v3)
                                        _ = v47
                                    } else {
                                        fmt.Println(v3, v3)
                                        v48 := GenS8{F0: 2}
                                        if false {
                                            var v49 uint16 = 42
                                            v3 = (v3 | (v3 - (v3 % ((1) | 1))))
                                            v49 = v49
                                            _ = v49
                                        }
                                        _ = v48
                                    }
                                    if true {
                                        if len(v15) > 0 {
                                            v15[0] = fi44
                                        }
                                    }
                                    _ = fi44
                                }
                                _ = v42
                            }
                            _ = v19
                        }
                    case -9223372036854775808:
                        fmt.Println(v3)
                    }
                    _ = v18
                }
                _ = v15
            } else {
                _ = len(v1)
            }
            _ = v11
        }
    }
    v3 = v3
    var v50 float64 = (0.0 / (0.0 + (3.14159 - float64(-2.25))))
    v3 = v3
    if ((v50 < (v50 * 3.14159)) || (42 == (^(v3)))) {
        v2[127] = true
        var v51 uint16 = 2
        var v52 int64 = 0
        _ = v51
        _ = v52
    }
    var v53 float64 = ((-2.25 - (((1e10 + v50) - 3.14159) + 3.14159)) * v50)
    if true {
        v50 = 3.14159
        v53 -= v50
    }
    var v54 int8 = 42
    switch 782 {
    case 936:
        var v55 float64 = (1.5 + v50)
        _ = v55
    case 2147483647:
        v53 += v50
    default:
        v56 := []int{}
        _ = v56
    }
    v57 := GenS9{F0: (v54 + (v54 - -128)), F1: 595}
    if true {
        v58 := GenS10{F0: "hello", F1: v3}
        v59 := []int{}
        _ = v58
        _ = v59
    } else {
        if false {
            v60 := GenS3{F0: 32768, F1: 18446744073709551615}
            var v61 int16 = 1
            _ = v60
            _ = v61
        } else {
            for fi62 := 0; fi62 < 1; fi62++ {
                switch ("a" + "") {
                case "\n\t":
                    if ((0.0 + 0.0) == 0.1) {
                        for fi63 := 0; fi63 < 8; fi63++ {
                            var v64 bool = false
                            var v65 int16 = -241
                            _ = fi63
                            _ = v64
                            _ = v65
                        }
                        var v66 int64 = 660
                        _ = v66
                    }
                case "foo bar":
                    var v67 uint8 = 0
                    _ = v67
                case "a":
                    v68 := GenS11{F0: v54, F1: v54}
                    _ = v68
                default:
                    v69 := GenS4{F0: 528}
                    _ = v69
                }
                _ = fi62
            }
            fmt.Println(v54)
            for fi70 := 0; fi70 < 19; fi70++ {
                v71 := GenS11{F0: v54, F1: v54}
                if true {
                    fmt.Println(v50, v3)
                    fmt.Println(v50, v53, fi70)
                    v54 += -128
                }
                _ = fi70
                _ = v71
            }
        }
        v53 = -2.25
        fmt.Println(v53, v53)
    }
    switch 42 {
    case 128:
        if (!true) {
            var v72 uint64 = 1
            fmt.Println(v3, v50, v50)
            var v73 uint8 = 255
            _ = v72
            _ = v73
        }
    case 860:
        var v74 uint32 = 541
        _ = v74
    case -429:
        if (!(!true)) {
            if false {
                switch 429 {
                case -511:
                    var v75 uint8 = 250
                    _ = v75
                case 977:
                    var v76 float64 = 1e10
                    _ = v76
                }
            } else {
                switch "" {
                case "foo bar":
                    v77 := GenS11{F0: int8(v3), F1: v54}
                    _ = v77
                case "\n\t":
                    var v78 string = "hello"
                    _ = v78
                default:
                    v2[2147483647] = true
                }
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v50:", v50)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v53
    _ = v54
    _ = v57
}
