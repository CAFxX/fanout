package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int64 }
type GenS2 struct { F0 int8; F1 uint32; F2 int16 }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 uint; F1 int16 }
type GenS5 struct { F0 uint; F1 int16; F2 int32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 string; F1 string }
type GenS8 struct { F0 bool; F1 bool }
type GenS9 struct { F0 int8; F1 uint16 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 uint; F1 int32 }
type GenS12 struct { F0 uint32; F1 uint16 }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 uint; F1 uint64; F2 uint8 }
type GenS16 struct { F0 int16 }
type GenS17 struct { F0 int8; F1 uint16 }
type GenS18 struct { F0 int32 }
type GenS19 struct { F0 bool }

func main() {
    var v1 string = ("a" + "")
    var v2 float64 = 0.1
    v3 := []string{v1}
    var v4 int32 = -2147483648
    var v5 bool = (true && (!(1 <= 42)))
    if len(v3) > 0 {
        v3[0] = (v1 + (v1 + v1))
    }
    if len(v3) > 0 {
        v3[0] = (v1 + (v1 + v1))
    }
    for fi6 := 0; fi6 < 4; fi6++ {
        fmt.Println(v5, v1, v4)
        v1 = "hello"
        v7 := GenS1{F0: 0, F1: 122, F2: 9223372036854775807}
        _ = fi6
        _ = v7
    }
    fmt.Println(v1, v2)
    if (v5 || v5) {
        v8 := GenS2{F0: -62, F1: 1, F2: 42}
        _ = v8
    }
    var v9 uint64 = 0
    for fi10 := 0; fi10 < 6; fi10++ {
        fmt.Println(fi10)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v11 uint16 = 42
        _ = fi10
        _ = v11
    }
    if (404 == 2147483648) {
        v12 := GenS3{F0: true, F1: 100}
        _ = v12
    } else {
        if (!(255 <= 506)) {
            v1 = ("hello" + "foo bar")
            if true {
                v9 = v9
                if (!v5) {
                    v13 := GenS4{F0: 2, F1: int16(v9)}
                    _ = v13
                }
            }
        } else {
            if (3.14159 <= 0.1) {
                for fi14 := 0; fi14 < 7; fi14++ {
                    var v15 int64 = 190
                    _ = fi14
                    _ = v15
                }
            } else {
                switch "" {
                case "a":
                    var v16 int32 = -2147483648
                    _ = v16
                case "":
                    switch 0 {
                    case -9223372036854775808:
                        switch v1 {
                        case "a":
                            var v17 string = "\n\t"
                            _ = v17
                        case "hello":
                            fmt.Println(v9, v4, v2)
                        case "\n\t":
                            v3 = append(v3, (v1 + (v1 + v1)))
                        default:
                            for fi18 := 0; fi18 < 10; fi18++ {
                                fmt.Println(v2, v4, v4)
                                if len(v3) > 0 {
                                    v3[0] = (((((((v1 + "hello") + v1) + (v1 + v1)) + v1) + (((v1 + v1) + v1) + v1)) + "a") + v1)
                                }
                                for fi19 := 0; fi19 < 20; fi19++ {
                                    if len(v3) > 0 {
                                        v3[0] = v1
                                    }
                                    _ = fi19
                                }
                                _ = fi18
                            }
                        }
                    case 32767:
                        v20 := GenS5{F0: 2, F1: -39, F2: (v4 - 562)}
                        _ = v20
                    case 702:
                        var v21 uint8 = 0
                        _ = v21
                    }
                case "\n\t":
                    for fi22 := 0; fi22 < 14; fi22++ {
                        if len(v3) > 0 {
                            v3[0] = ("" + v1)
                        }
                        if false {
                            v2 -= v2
                            switch v1 {
                            case "a":
                                switch "\n\t" {
                                case "hello":
                                    v23 := GenS5{F0: 308, F1: 923, F2: v4}
                                    _ = v23
                                default:
                                    fmt.Println(fi22)
                                }
                            case "":
                                var v24 float32 = ((0.0 * ((3.14159 * 0.0) * (-2.25 + (1e10 + float32(((float64(0.1) + v2) * -2.25)))))) / 0.1)
                                _ = v24
                            case "\n\t":
                                v25 := GenS3{F0: v5, F1: -1}
                                _ = v25
                            default:
                                for fi26 := 0; fi26 < 8; fi26++ {
                                    if false {
                                        if false {
                                            var v27 uint = 1
                                            v28 := GenS6{F0: -1}
                                            if v5 {
                                                if v5 {
                                                    fmt.Println(v4)
                                                    var v29 bool = (((v9 >> (((^(v27))) & 63)) >= (v9 & v9)) && (!(!v5)))
                                                    switch v1 {
                                                    case "":
                                                        v9 += (v9 + v9)
                                                    case "foo bar":
                                                        var v30 uint16 = 65535
                                                        _ = v30
                                                    case "hello":
                                                        var v31 uint16 = 65535
                                                        _ = len(v3)
                                                        _ = v31
                                                    default:
                                                        v29 = (1e10 < 0.0)
                                                    }
                                                    _ = v29
                                                }
                                                v3 = append(v3, (("\n\t" + "hello") + ("" + v1)))
                                                var v32 uint64 = (^(v9))
                                                _ = v32
                                            } else {
                                                fmt.Println(v4, v27, v4)
                                                var v33 uint64 = v9
                                                _ = v33
                                            }
                                            _ = v27
                                            _ = v28
                                        } else {
                                            if ((128 == 1) && v5) {
                                                switch (^(fi22)) {
                                                case -1:
                                                    if (2 > 4294967295) {
                                                        v34 := GenS2{F0: -1, F1: 19, F2: 42}
                                                        v35 := GenS2{F0: -9, F1: 2, F2: -921}
                                                        var v36 int64 = 9223372036854775807
                                                        _ = v34
                                                        _ = v35
                                                        _ = v36
                                                    }
                                                case -907:
                                                    v5 = v5
                                                case 0:
                                                    _ = len(v3)
                                                default:
                                                    if false {
                                                        v5 = v5
                                                        if len(v3) > 0 {
                                                            v3[0] = v1
                                                        }
                                                        var v37 uint16 = uint16(v9)
                                                        _ = v37
                                                    } else {
                                                        fmt.Println(v4, v1)
                                                    }
                                                }
                                                v9 += (^(v9))
                                            }
                                        }
                                        v5 = true
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    } else {
                                        switch v1 {
                                        case "\n\t":
                                            v4 = -967
                                            if len(v3) > 0 {
                                                v3[0] = "hello"
                                            }
                                            if v5 {
                                                if false {
                                                    if v5 {
                                                        v38 := GenS7{F0: "hello", F1: (v1 + (v1 + v1))}
                                                        var v39 float32 = -2.25
                                                        _ = v38
                                                        _ = v39
                                                    } else {
                                                        var v40 string = "a"
                                                        var v41 int64 = 42
                                                        v9 += (v9 % (((v9 | 0)) | 1))
                                                        _ = v40
                                                        _ = v41
                                                    }
                                                    var v42 int8 = 1
                                                    _ = v42
                                                }
                                                if len(v3) > 0 {
                                                    v3[0] = (v1 + v1)
                                                }
                                            }
                                        case "":
                                            var v43 string = ("foo bar" + "")
                                            _ = v43
                                        case "a":
                                            switch -1000 {
                                            case -1:
                                                var v44 uint16 = 292
                                                _ = v44
                                            }
                                        default:
                                            if len(v3) > 0 {
                                                v3[0] = v1
                                            }
                                        }
                                        v45 := GenS8{F0: (498 == 42), F1: v5}
                                        _ = v45
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = v1
                                    }
                                    _ = fi26
                                }
                            }
                        }
                        if v5 {
                            switch v1 {
                            case "\n\t":
                                v2 += v2
                            case "a":
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            default:
                                for fi46 := 0; fi46 < 15; fi46++ {
                                    if true {
                                        v47 := GenS9{F0: 127, F1: 567}
                                        v9 = v9
                                        var v48 uint32 = 596
                                        _ = v47
                                        _ = v48
                                    } else {
                                        if false {
                                            v4 = (v4 >> ((47) & 63))
                                            v9 = 55
                                        } else {
                                            var v49 int = (fi46 * fi22)
                                            if (((!(!((v9 + v9) <= 2))) && v5) && (v1 <= (v1 + v1))) {
                                                var v50 uint16 = 581
                                                v3 = append(v3, v1)
                                                switch v1 {
                                                case "foo bar":
                                                    fmt.Println(fi46)
                                                case "a":
                                                    _ = len(v3)
                                                case "hello":
                                                    var v51 int8 = 0
                                                    _ = v51
                                                }
                                                if true {
                                                    v4 -= v4
                                                    switch "hello" {
                                                    case "\n\t":
                                                        var v52 int32 = (v4 & (v4 - v4))
                                                        _ = v52
                                                    case "a":
                                                        switch v1 {
                                                        case "hello":
                                                            if ((0.1 == 1e10) || v5) {
                                                                if v5 {
                                                                    var v53 float64 = float64((3.14159 + 1.5))
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                    if true {
                                                                        switch v1 {
                                                                        case "\n\t":
                                                                            fmt.Println(fi46, v2, v53)
                                                                        }
                                                                    } else {
                                                                        v3 = append(v3, v1)
                                                                        v5 = false
                                                                        if v5 {
                                                                            if true {
                                                                                fmt.Println(v53)
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v53
                                                                } else {
                                                                    if len(v3) > 0 {
                                                                        v3[0] = (((v1 + v1) + ("" + ("a" + (("hello" + v1) + v1)))) + (v1 + "foo bar"))
                                                                    }
                                                                    var v54 int32 = v4
                                                                    switch (v1 + v1) {
                                                                    case "\n\t":
                                                                        if v5 {
                                                                            v4 += v4
                                                                            if len(v3) > 0 {
                                                                                v3[0] = v1
                                                                            }
                                                                            v4 -= (v54 & (-(v4)))
                                                                        }
                                                                        v55 := GenS5{F0: 997, F1: 1, F2: v4}
                                                                        if len(v3) > 0 {
                                                                            v3[0] = (v1 + v1)
                                                                        }
                                                                        _ = v55
                                                                    }
                                                                    v9 = v9
                                                                    var v56 uint8 = 128
                                                                    _ = v54
                                                                    _ = v56
                                                                }
                                                                var v57 uint16 = 255
                                                                switch v49 {
                                                                case -1:
                                                                    v58 := GenS4{F0: 282, F1: -987}
                                                                    _ = v58
                                                                default:
                                                                    v57 += v50
                                                                }
                                                                switch fi46 {
                                                                case -4:
                                                                    v2 -= 1e10
                                                                }
                                                                v5 = v5
                                                                _ = v57
                                                            }
                                                        default:
                                                            fmt.Println(fi22, v4, v4)
                                                        }
                                                    case "":
                                                        var v59 string = (v1 + v1)
                                                        _ = v59
                                                    }
                                                } else {
                                                    v3 = append(v3, v1)
                                                    v3 = append(v3, v1)
                                                    v50 = (^(v50))
                                                    if (1e10 != ((1.5 + v2) - v2)) {
                                                        switch v1 {
                                                        case "a":
                                                            _ = len(v3)
                                                        case "":
                                                            v50 += v50
                                                            if false {
                                                                if (((v1 + v1) + v1) != (v1 + (((v1 + ("a" + v1)) + (("hello" + "\n\t") + v1)) + v1))) {
                                                                    v60 := GenS10{F0: 0}
                                                                    if v5 {
                                                                        var v61 string = ("foo bar" + (v1 + (v1 + (v1 + ""))))
                                                                        v3 = append(v3, ((v1 + (((v1 + v1) + v61) + v61)) + "hello"))
                                                                        _ = v61
                                                                    }
                                                                    var v62 uint64 = v9
                                                                    _ = v60
                                                                    _ = v62
                                                                }
                                                                if (v1 <= ((v1 + v1) + v1)) {
                                                                    switch ("\n\t" + v1) {
                                                                    case "":
                                                                        v1 = "\n\t"
                                                                        v63 := GenS6{F0: 9223372036854775807}
                                                                        if ((^(v9)) > v9) {
                                                                            v64 := GenS8{F0: (65535 >= 42), F1: false}
                                                                            _ = v64
                                                                        }
                                                                        switch 128 {
                                                                        case 0:
                                                                            v65 := GenS11{F0: 278, F1: v4}
                                                                            _ = v65
                                                                        case 128:
                                                                            _ = len(v3)
                                                                        default:
                                                                            if true {
                                                                                v66 := GenS3{F0: v5, F1: -123}
                                                                                _ = v66
                                                                            } else {
                                                                                v67 := GenS11{F0: 42, F1: v4}
                                                                                var v68 uint = 18446744073709551615
                                                                                _ = v67
                                                                                _ = v68
                                                                            }
                                                                        }
                                                                        var v69 int64 = 9223372036854775807
                                                                        v50 -= v50
                                                                        _ = v63
                                                                        _ = v69
                                                                    default:
                                                                        var v70 float64 = 3.14159
                                                                        _ = v70
                                                                    }
                                                                    if ((809 <= 18446744073709551615) || v5) {
                                                                        v9 -= v9
                                                                        if (100 >= -128) {
                                                                            if (-9223372036854775808 == 42) {
                                                                                v1 = v1
                                                                                var v71 uint16 = (v50 >> ((642) & 63))
                                                                                v50 = (v50 << ((0) & 63))
                                                                                _ = v71
                                                                            }
                                                                            _ = len(v3)
                                                                            var v72 uint32 = 210
                                                                            _ = v72
                                                                        }
                                                                        fmt.Println(fi22, v1)
                                                                    } else {
                                                                        if v5 {
                                                                            v4 += v4
                                                                            v73 := GenS12{F0: 0, F1: 2}
                                                                            var v74 float32 = 0.0
                                                                            _ = v73
                                                                            _ = v74
                                                                        } else {
                                                                            if v5 {
                                                                                v75 := GenS13{F0: -9223372036854775808}
                                                                                if true {
                                                                                    v50 = v50
                                                                                } else {
                                                                                    v49 = (fi46 / ((v49) | 1))
                                                                                    v76 := GenS10{F0: -126}
                                                                                    _ = v76
                                                                                }
                                                                                _ = v75
                                                                            }
                                                                            switch 32767 {
                                                                            case -9223372036854775808:
                                                                                fmt.Println(fi22, v9)
                                                                            case 4294967295:
                                                                                var v77 float32 = 0.0
                                                                                _ = v77
                                                                            case 0:
                                                                                if (-128 == -1) {
                                                                                    var v78 int16 = 1000
                                                                                    v5 = false
                                                                                    _ = v78
                                                                                } else {
                                                                                    v49 = (v49 >> ((720) & 63))
                                                                                    var v79 uint32 = 1
                                                                                    _ = v79
                                                                                }
                                                                            default:
                                                                                var v80 int32 = 2147483647
                                                                                _ = v80
                                                                            }
                                                                            v81 := GenS9{F0: 100, F1: v50}
                                                                            _ = v81
                                                                        }
                                                                        v82 := GenS2{F0: -66, F1: 677, F2: int16(v50)}
                                                                        _ = v82
                                                                    }
                                                                    v50 -= 255
                                                                } else {
                                                                    var v83 uint64 = v9
                                                                    _ = v83
                                                                }
                                                                v2 = (float64(0.0) - ((v2 / v2) + float64((0.1 - 1.5))))
                                                            }
                                                            v84 := GenS14{F0: (^(v50))}
                                                            v85 := GenS15{F0: 1, F1: v9, F2: 255}
                                                            v3 = append(v3, v1)
                                                            _ = v84
                                                            _ = v85
                                                        case "hello":
                                                            v5 = (-1 == 42)
                                                        default:
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        }
                                                        fmt.Println(v50)
                                                        switch 1 {
                                                        case 0:
                                                            switch v49 {
                                                            case -365:
                                                                var v86 uint32 = 438
                                                                _ = v86
                                                            case -262:
                                                                v50 = v50
                                                            default:
                                                                v87 := GenS5{F0: 255, F1: 877, F2: -2147483648}
                                                                _ = v87
                                                            }
                                                        case 500:
                                                            if v5 {
                                                                switch v1 {
                                                                case "a":
                                                                    var v88 int8 = 100
                                                                    _ = v88
                                                                case "foo bar":
                                                                    var v89 int64 = 9223372036854775807
                                                                    _ = v89
                                                                default:
                                                                    fmt.Println(fi46)
                                                                }
                                                                v3 = append(v3, (v1 + "hello"))
                                                            }
                                                        default:
                                                            var v90 int8 = 100
                                                            _ = v90
                                                        }
                                                    } else {
                                                        v3 = append(v3, (v1 + v1))
                                                        if (0 <= v9) {
                                                            switch -9223372036854775808 {
                                                            case -1:
                                                                v4 += v4
                                                            case -9223372036854775808:
                                                                v91 := GenS16{F0: 1000}
                                                                _ = v91
                                                            }
                                                            fmt.Println(v2, fi46)
                                                            _ = len(v3)
                                                            switch "a" {
                                                            case "\n\t":
                                                                v49 = v49
                                                            default:
                                                                fmt.Println(v1, v4)
                                                            }
                                                            fmt.Println(v4, v9)
                                                        } else {
                                                            var v92 float32 = -2.25
                                                            _ = v92
                                                        }
                                                    }
                                                }
                                                _ = v50
                                            }
                                            _ = v49
                                        }
                                        var v93 int32 = (^(v4))
                                        _ = v93
                                    }
                                    v94 := GenS13{F0: -314}
                                    _ = fi46
                                    _ = v94
                                }
                            }
                            v9 = 365
                            var v95 float64 = float64(1.5)
                            _ = v95
                        }
                        _ = fi22
                    }
                }
                if v5 {
                    var v96 uint = 65535
                    _ = v96
                }
            }
        }
    }
    fmt.Println(v2, v9, v5)
    if v5 {
        var v97 uint32 = 937
        v98 := GenS17{F0: 0, F1: 255}
        if (v5 && v5) {
            if (!(1 == -1)) {
                fmt.Println(v97, v4)
                v4 += 550
            }
        } else {
            var v99 uint16 = 333
            if ((^(v9)) >= uint64(v97)) {
                v100 := GenS18{F0: v4}
                switch (v1 + ((("" + v1) + v1) + (v1 + "a"))) {
                case "":
                    _ = len(v3)
                default:
                    v9 = v9
                }
                _ = v100
            }
            v1 = (v1 + ("hello" + (v1 + v1)))
            _ = v99
        }
        _ = v97
        _ = v98
    }
    v1 = "a"
    if false {
        var v101 string = v1
        if len(v3) > 0 {
            v3[0] = v1
        }
        for fi102 := 0; fi102 < 20; fi102++ {
            switch fi102 {
            case -1:
                fmt.Println(v9, v2)
            case -205:
                v103 := GenS19{F0: v5}
                _ = v103
            default:
                v2 = ((v2 + v2) / v2)
            }
            v4 += v4
            _ = fi102
        }
        _ = v101
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
