package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int64; F1 uint32; F2 uint8 }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 bool; F1 int16; F2 uint32 }
type GenS6 struct { F0 int32; F1 int32; F2 string }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int; F1 uint32 }

func main() {
    var v1 int32 = 0
    var v2 uint32 = 2147483647
    var v3 float64 = ((((3.14159 - -2.25) * 3.14159) * 3.14159) / 1.5)
    fmt.Println(v3, v1)
    v4 := GenS1{F0: 42, F1: 4294967295}
    var v5 string = ("\n\t" + "hello")
    fmt.Println(v3, v3, v1)
    for fi6 := 0; fi6 < 15; fi6++ {
        if ((v1 / (((v1 - -1)) | 1)) != (v1 & v1)) {
            v5 = v5
            fmt.Println(v2, v2)
            v7 := GenS1{F0: 18446744073709551615, F1: (v2 - v2)}
            _ = v7
        } else {
            var v8 uint64 = 101
            _ = v8
        }
        if (1 > 1000) {
            var v9 uint8 = 0
            if (1 > -413) {
                if false {
                    if true {
                        if false {
                            v5 = v5
                        } else {
                            v10 := GenS2{F0: -1, F1: (v2 >> ((42) & 63)), F2: (v9 * 1)}
                            switch "hello" {
                            case "":
                                if true {
                                    v3 += 1.5
                                    if true {
                                        if (false && true) {
                                            var v11 string = (v5 + ("foo bar" + v5))
                                            switch 0 {
                                            case 127:
                                                var v12 float32 = (-2.25 / 1.5)
                                                _ = v12
                                            case 728:
                                                var v13 int = int(v9)
                                                _ = v13
                                            }
                                            _ = v11
                                        }
                                        var v14 float32 = (1e10 - (-2.25 * 1e10))
                                        if (-410 > 32767) {
                                            if (((float64(1.5) / v3) != (v3 + 0.1)) || false) {
                                                var v15 uint8 = (v9 - v9)
                                                var v16 bool = (!true)
                                                _ = v15
                                                _ = v16
                                            }
                                        } else {
                                            v17 := []int{}
                                            _ = v17
                                        }
                                        _ = v14
                                    }
                                    switch (fi6 % (((fi6 % (((fi6 - (fi6 | (-(fi6))))) | 1))) | 1)) {
                                    case 2147483647:
                                        var v18 uint = 255
                                        _ = v18
                                    case 32767:
                                        if true {
                                            v9 -= v9
                                            if (float32(v3) >= (0.0 + float32(float64(1.5)))) {
                                                for fi19 := 0; fi19 < 13; fi19++ {
                                                    v1 = (v1 & (^(v1)))
                                                    var v20 uint = 255
                                                    _ = fi19
                                                    _ = v20
                                                }
                                                var v21 uint32 = (v2 >> ((2) & 63))
                                                v9 -= (v9 / (((^(v9))) | 1))
                                                _ = v21
                                            }
                                        }
                                    default:
                                        if false {
                                            var v22 float64 = v3
                                            if (!(v5 >= "hello")) {
                                                v3 -= 3.14159
                                            }
                                            _ = v22
                                        } else {
                                            switch (fi6 >> ((409) & 63)) {
                                            case -721:
                                                switch v5 {
                                                case "a":
                                                    var v23 int32 = 376
                                                    _ = v23
                                                case "\n\t":
                                                    for fi24 := 0; fi24 < 9; fi24++ {
                                                        v25 := GenS2{F0: 0, F1: (v2 % ((2147483647) | 1)), F2: 2}
                                                        fmt.Println(v3)
                                                        var v26 int = fi6
                                                        _ = fi24
                                                        _ = v25
                                                        _ = v26
                                                    }
                                                case "hello":
                                                    fmt.Println(v3, v2)
                                                }
                                            default:
                                                for fi27 := 0; fi27 < 6; fi27++ {
                                                    v5 = "\n\t"
                                                    v28 := []int{(fi6 | fi6)}
                                                    _ = fi27
                                                    _ = v28
                                                }
                                            }
                                            var v29 uint8 = v9
                                            _ = v29
                                        }
                                    }
                                }
                            case "hello":
                                fmt.Println(v9, v3)
                            case "\n\t":
                                v30 := GenS1{F0: 65535, F1: 744}
                                _ = v30
                            }
                            _ = v10
                        }
                    } else {
                        v31 := []int{(fi6 >> ((42) & 63))}
                        if false {
                            if (!false) {
                                if true {
                                    var v32 bool = (fi6 < fi6)
                                    v33 := GenS3{F0: v32, F1: 2}
                                    _ = v32
                                    _ = v33
                                } else {
                                    v31 = append(v31, (fi6 << ((1) & 63)))
                                }
                                var v34 float64 = v3
                                v34 = 0.1
                                _ = v34
                            } else {
                                v35 := GenS3{F0: false, F1: 255}
                                switch (fi6 + (fi6 / (((-(fi6))) | 1))) {
                                case 127:
                                    if true {
                                        fmt.Println(fi6)
                                    }
                                default:
                                    if false {
                                        for fi36 := 0; fi36 < 18; fi36++ {
                                            fmt.Println(v5)
                                            v9 -= (^(v9))
                                            _ = fi36
                                        }
                                        fmt.Println(v5)
                                    } else {
                                        if false {
                                            v31 = append(v31, fi6)
                                        } else {
                                            var v37 float32 = -2.25
                                            v5 = v5
                                            _ = v37
                                        }
                                        switch 939 {
                                        case 4294967295:
                                            if (!true) {
                                                var v38 int16 = -767
                                                _ = v38
                                            }
                                        case -176:
                                            v2 = 814
                                        case 32768:
                                            v2 -= 4294967295
                                        default:
                                            fmt.Println(v9)
                                        }
                                    }
                                }
                                _ = v35
                            }
                        } else {
                            v39 := GenS4{F0: true}
                            _ = len(v31)
                            _ = v39
                        }
                        _ = v31
                    }
                    var v40 float32 = (-2.25 + (0.1 - 3.14159))
                    var v41 uint = 91
                    _ = v40
                    _ = v41
                } else {
                    var v42 float32 = (3.14159 / 3.14159)
                    switch (fi6 - (fi6 / ((fi6) | 1))) {
                    case 127:
                        switch (fi6 + fi6) {
                        case -2147483648:
                            v43 := []int{fi6, fi6, (fi6 >> ((18446744073709551615) & 63)), 32767}
                            _ = v43
                        case 127:
                            fmt.Println(v5, v5, v3)
                        case 0:
                            v44 := []int{fi6, fi6, (^(fi6))}
                            _ = v44
                        }
                    case 32768:
                        if (!false) {
                            v3 -= 0.0
                            for fi45 := 0; fi45 < 20; fi45++ {
                                var v46 string = v5
                                switch (fi6 ^ 4294967295) {
                                case 1:
                                    v47 := GenS3{F0: false, F1: 2}
                                    _ = v47
                                case 42:
                                    fmt.Println(fi6, fi6)
                                }
                                v48 := []int{}
                                _ = fi45
                                _ = v46
                                _ = v48
                            }
                            v1 -= v1
                        } else {
                            fmt.Println(v2, v42)
                            for fi49 := 0; fi49 < 18; fi49++ {
                                v1 -= v1
                                v3 += (3.14159 * v3)
                                _ = fi49
                            }
                        }
                    }
                    v50 := GenS1{F0: 295, F1: v2}
                    _ = v42
                    _ = v50
                }
                if (false && false) {
                    switch v5 {
                    case "\n\t":
                        switch 4294967295 {
                        case 0:
                            for fi51 := 0; fi51 < 16; fi51++ {
                                if false {
                                    switch v5 {
                                    case "hello":
                                        v52 := []int{fi51}
                                        _ = v52
                                    case "":
                                        var v53 int = -9223372036854775808
                                        _ = v53
                                    }
                                } else {
                                    var v54 float64 = -2.25
                                    if false {
                                        v55 := []int{127}
                                        if false {
                                            var v56 uint64 = 711
                                            v57 := GenS3{F0: (v56 != (v56 * v56)), F1: 42}
                                            _ = v56
                                            _ = v57
                                        }
                                        var v58 bool = false
                                        _ = v55
                                        _ = v58
                                    }
                                    _ = v54
                                }
                                v5 = (v5 + (((v5 + "a") + (v5 + "")) + (v5 + (v5 + v5))))
                                v59 := GenS2{F0: -832, F1: 42, F2: (v9 - 1)}
                                _ = fi51
                                _ = v59
                            }
                        case 4294967295:
                            v9 = (v9 >> ((989) & 63))
                        case -1:
                            for fi60 := 0; fi60 < 2; fi60++ {
                                var v61 string = v5
                                v62 := GenS5{F0: (!true), F1: int16(v2), F2: v2}
                                fmt.Println(v61, v5)
                                _ = fi60
                                _ = v61
                                _ = v62
                            }
                        default:
                            v63 := []int{(^(fi6)), fi6, fi6, fi6}
                            _ = v63
                        }
                    default:
                        if (127 == -128) {
                            v64 := GenS2{F0: 42, F1: v2, F2: 0}
                            if false {
                                v5 = "a"
                            }
                            _ = v64
                        }
                    }
                } else {
                    v65 := GenS5{F0: false, F1: 32767, F2: (v2 & v2)}
                    for fi66 := 0; fi66 < 20; fi66++ {
                        if true {
                            v2 += (^(v2))
                            switch v5 {
                            case "foo bar":
                                v5 = v5
                            case "hello":
                                v67 := GenS4{F0: (-1 > 42)}
                                _ = v67
                            default:
                                if true {
                                    v68 := GenS3{F0: (763 <= 998), F1: 909}
                                    v69 := []int{}
                                    var v70 string = ("\n\t" + (((v5 + v5) + "hello") + v5))
                                    _ = v68
                                    _ = v69
                                    _ = v70
                                } else {
                                    var v71 string = "foo bar"
                                    v2 -= (v2 * (v2 * v2))
                                    var v72 uint16 = 523
                                    _ = v71
                                    _ = v72
                                }
                                v73 := GenS5{F0: (!(!true)), F1: 613, F2: v2}
                                _ = v73
                            }
                        }
                        fmt.Println(v1, v3)
                        v74 := []int{fi6, int(v9)}
                        _ = fi66
                        _ = v74
                    }
                    fmt.Println(v5, v1, v3)
                    _ = v65
                }
                if false {
                    v75 := []int{(fi6 ^ (^(fi6))), fi6, fi6}
                    _ = v75
                } else {
                    switch (fi6 % (((^(fi6))) | 1)) {
                    case 4294967295:
                        v76 := []int{}
                        _ = v76
                    case 32768:
                        v5 = v5
                    default:
                        if (false && ((!(fi6 >= (fi6 & fi6))) && ((42 == 381) || (!false)))) {
                            for fi77 := 0; fi77 < 6; fi77++ {
                                fmt.Println(v9, v2, v1)
                                _ = fi77
                            }
                            var v78 string = v5
                            v79 := []int{}
                            _ = v78
                            _ = v79
                        }
                    }
                    for fi80 := 0; fi80 < 18; fi80++ {
                        switch v5 {
                        case "":
                            var v81 uint = 2
                            _ = v81
                        default:
                            if true {
                                v82 := GenS6{F0: (v1 << ((4294967295) & 63)), F1: 1, F2: v5}
                                var v83 uint64 = 137
                                switch v5 {
                                case "foo bar":
                                    v84 := GenS6{F0: -1, F1: v1, F2: "a"}
                                    v5 = v5
                                    var v85 string = v5
                                    _ = v84
                                    _ = v85
                                }
                                _ = v82
                                _ = v83
                            }
                            v86 := GenS7{F0: false}
                            _ = v86
                        }
                        var v87 uint32 = v2
                        _ = fi80
                        _ = v87
                    }
                }
            } else {
                v88 := GenS1{F0: 255, F1: (v2 << ((42) & 63))}
                for fi89 := 0; fi89 < 1; fi89++ {
                    var v90 float64 = (0.0 - (((v3 - v3) + v3) + v3))
                    _ = fi89
                    _ = v90
                }
                _ = v88
            }
            _ = v9
        }
        v91 := GenS5{F0: (!true), F1: -32768, F2: 2147483647}
        _ = fi6
        _ = v91
    }
    v92 := []int{899, -829, 720}
    switch 32768 {
    case 989:
        fmt.Println(v1, v3, v1)
    case 32767:
        for fi93 := 0; fi93 < 2; fi93++ {
            v5 = v5
            _ = fi93
        }
    case 127:
        for fi94 := 0; fi94 < 13; fi94++ {
            fmt.Println(v1)
            v95 := GenS8{F0: (fi94 & fi94), F1: 0}
            var v96 float32 = 0.0
            _ = fi94
            _ = v95
            _ = v96
        }
    default:
        var v97 int = 637
        _ = v97
    }
    var v98 int64 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v92:", len(v92), v92)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v92
    _ = v98
}
