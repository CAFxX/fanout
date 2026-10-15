package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint; F1 uint; F2 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint32; F1 int32 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 bool }

func main() {
    var v1 string = ("a" + ("" + "a"))
    var v2 int64 = -1
    var v3 uint32 = 2147483648
    v1 = ("a" + ((("" + v1) + v1) + v1))
    for fi4 := 0; fi4 < 18; fi4++ {
        v3 = v3
        v3 += (v3 ^ (v3 % ((2) | 1)))
        _ = fi4
    }
    if (false && false) {
        if true {
            fmt.Println(v2, v1)
            var v5 int = 2147483647
            _ = v5
        }
        var v6 bool = false
        v7 := GenS1{F0: 779, F1: 1}
        _ = v6
        _ = v7
    }
    var v8 float32 = 1e10
    var v9 int = -489
    v9 = (v9 + 4294967295)
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 int64 = v2
        v11 -= v11
        v2 += 0
        _ = fi10
        _ = v11
    }
    v12 := GenS2{F0: 559, F1: 0, F2: v3}
    var v13 bool = false
    for fi14 := 0; fi14 < 19; fi14++ {
        if (false && (v13 || (false && (!true)))) {
            if ("\n\t" <= "foo bar") {
                for fi15 := 0; fi15 < 4; fi15++ {
                    v16 := []int{(v9 & -1), fi14, 2147483647, 32767}
                    var v17 uint = 2
                    if ((fi14 % ((v9) | 1)) != (fi15 * 2147483647)) {
                        v18 := GenS1{F0: (v17 << (((v17 >> ((v17) & 63))) & 63)), F1: -506}
                        if len(v16) > 0 {
                            v16[0] = fi14
                        }
                        v3 = v3
                        _ = v18
                    } else {
                        if v13 {
                            v16 = append(v16, (fi14 >> (((v17 - v17)) & 63)))
                            v8 = v8
                            var v19 float32 = float32(0.1)
                            _ = v19
                        } else {
                            v1 = v1
                            if v13 {
                                switch v9 {
                                case -2147483648:
                                    if v13 {
                                        switch ("" + "a") {
                                        case "":
                                            _ = len(v16)
                                        case "foo bar":
                                            if (241 != 1) {
                                                if (!(-9223372036854775808 != (v2 >> ((v17) & 63)))) {
                                                    if (v13 || (float64(v8) < 1e10)) {
                                                        v20 := GenS2{F0: 586, F1: v17, F2: (^(v3))}
                                                        v13 = (!(600 <= -9223372036854775808))
                                                        _ = v20
                                                    }
                                                    v21 := GenS3{F0: 2}
                                                    fmt.Println(fi14)
                                                    _ = v21
                                                }
                                                if (!(("foo bar" + v1) != (v1 + "foo bar"))) {
                                                    var v22 uint = 195
                                                    _ = v22
                                                }
                                                v23 := GenS2{F0: 797, F1: (v17 % ((781) | 1)), F2: v3}
                                                _ = v23
                                            } else {
                                                switch v1 {
                                                case "a":
                                                    fmt.Println(fi14, v9)
                                                    v13 = (v2 > (v2 / ((v2) | 1)))
                                                    if v13 {
                                                        var v24 uint64 = 2
                                                        var v25 int64 = (^(v2))
                                                        if len(v16) > 0 {
                                                            v16[0] = (fi14 * (fi14 % ((v9) | 1)))
                                                        }
                                                        _ = v24
                                                        _ = v25
                                                    }
                                                case "":
                                                    v26 := GenS3{F0: uint64(v17)}
                                                    _ = v26
                                                }
                                                v27 := GenS4{F0: 20}
                                                v17 = 929
                                                _ = v27
                                            }
                                        case "hello":
                                            v2 = v2
                                        default:
                                            fmt.Println(v13)
                                        }
                                        v1 = v1
                                    }
                                case 127:
                                    v3 += (v3 & (v3 + 2147483648))
                                default:
                                    switch fi14 {
                                    case 42:
                                        var v28 int = 42
                                        _ = v28
                                    }
                                }
                                if v13 {
                                    var v29 uint64 = 445
                                    var v30 bool = true
                                    _ = v29
                                    _ = v30
                                } else {
                                    v31 := GenS1{F0: (v17 >> ((1) & 63)), F1: 2147483647}
                                    fmt.Println(v1, v13)
                                    _ = v31
                                }
                            } else {
                                fmt.Println(fi14, v2, fi15)
                                v32 := GenS3{F0: 2}
                                _ = v32
                            }
                            switch "hello" {
                            case "hello":
                                _ = len(v16)
                            case "":
                                v9 -= (^(fi14))
                            }
                        }
                        var v33 uint = v17
                        v33 -= v33
                        _ = v33
                    }
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                switch v1 {
                case "foo bar":
                    if v13 {
                        if (-128 != -111) {
                            var v34 int64 = (v2 / ((9223372036854775807) | 1))
                            _ = v34
                        } else {
                            for fi35 := 0; fi35 < 20; fi35++ {
                                v8 = 1.5
                                v36 := GenS5{F0: 127}
                                if (v1 != v1) {
                                    switch v1 {
                                    case "foo bar":
                                        v37 := []int{fi14, (fi14 % ((int(v2)) | 1)), fi35}
                                        _ = v37
                                    default:
                                        var v38 uint8 = 2
                                        v39 := []int{(v9 & v9), (fi35 + fi14), (fi35 * fi14), fi35}
                                        var v40 int32 = -787
                                        _ = v38
                                        _ = v39
                                        _ = v40
                                    }
                                    fmt.Println(v1, v13, v3)
                                    fmt.Println(v1, v2)
                                } else {
                                    fmt.Println(v2, fi35)
                                    var v41 uint64 = 717
                                    fmt.Println(v13, v9, fi35)
                                    if v13 {
                                        fmt.Println(v3, fi14, v1)
                                        fmt.Println(v3, v13)
                                        v1 = v1
                                    }
                                    if v13 {
                                        v42 := GenS3{F0: 42}
                                        switch 32768 {
                                        case -142:
                                            var v43 uint8 = 255
                                            _ = v43
                                        }
                                        v44 := GenS6{F0: v41}
                                        _ = v42
                                        _ = v44
                                    }
                                    _ = v41
                                }
                                _ = fi35
                                _ = v36
                            }
                            var v45 uint16 = 65535
                            _ = v45
                        }
                        switch -9223372036854775808 {
                        case 127:
                            if ((("\n\t" + (v1 + v1)) + (("hello" + v1) + "hello")) != v1) {
                                for fi46 := 0; fi46 < 13; fi46++ {
                                    var v47 int16 = 32767
                                    if false {
                                        var v48 uint16 = 32767
                                        v2 -= v2
                                        v3 -= uint32(v48)
                                        _ = v48
                                    } else {
                                        v49 := []int{(^(fi14))}
                                        v50 := GenS4{F0: 0}
                                        if (((float64(v8) + 3.14159) + 1e10) < 0.1) {
                                            var v51 int16 = v47
                                            _ = v51
                                        }
                                        v47 = (v47 ^ v47)
                                        _ = v49
                                        _ = v50
                                    }
                                    v52 := GenS7{F0: 0, F1: 908}
                                    v53 := []int{}
                                    v47 = (v47 >> ((18446744073709551615) & 63))
                                    var v54 int16 = 1
                                    var v55 int32 = -1
                                    _ = fi46
                                    _ = v47
                                    _ = v52
                                    _ = v53
                                    _ = v54
                                    _ = v55
                                }
                                var v56 int = (fi14 & (fi14 + 1))
                                _ = v56
                            }
                        case 32767:
                            fmt.Println(fi14, v3)
                        case 32768:
                            var v57 uint16 = uint16(v3)
                            _ = v57
                        default:
                            var v58 int64 = 42
                            _ = v58
                        }
                    } else {
                        v9 = 417
                        var v59 bool = false
                        _ = v59
                    }
                case "":
                    var v60 uint16 = 183
                    _ = v60
                }
                var v61 uint16 = 238
                _ = v61
            }
            switch fi14 {
            case 128:
                if (!true) {
                    v62 := []int{fi14}
                    switch v1 {
                    case "\n\t":
                        if (((554 == 499) && ((!false) || v13)) && v13) {
                            var v63 int8 = -72
                            v64 := GenS8{F0: 926}
                            _ = v63
                            _ = v64
                        }
                    case "a":
                        var v65 int64 = 272
                        _ = v65
                    case "foo bar":
                        if (-1 <= 127) {
                            switch fi14 {
                            case 32768:
                                v13 = v13
                            case 0:
                                v66 := GenS1{F0: 582, F1: 937}
                                _ = v66
                            }
                            v8 = float32(((((0.0 * float64(0.0)) - (1.5 - 1.5)) / 1.0) * ((0.0 + 1e10) / 1e10)))
                            v13 = (-1 < -1)
                        }
                    }
                    switch 4294967295 {
                    case 425:
                        var v67 int = (fi14 + fi14)
                        _ = v67
                    case 0:
                        v68 := GenS9{F0: (-2147483648 < -1)}
                        _ = v68
                    case 32767:
                        var v69 int8 = -27
                        _ = v69
                    default:
                        var v70 int32 = 302
                        _ = v70
                    }
                    _ = v62
                } else {
                    var v71 string = v1
                    v72 := GenS1{F0: 269, F1: 1}
                    _ = v71
                    _ = v72
                }
            case -314:
                v73 := []int{}
                _ = v73
            default:
                var v74 int8 = 100
                _ = v74
            }
        } else {
            for fi75 := 0; fi75 < 5; fi75++ {
                v9 = fi75
                _ = fi75
            }
        }
        _ = fi14
    }
    fmt.Println(v8, v1)
    var v76 float32 = float32((3.14159 - 0.0))
    if (v76 <= (0.1 - (v8 + (float32(float64((((v76 + (float32(0.0) + float32(1e10))) / 1.0) + v8))) / v76)))) {
        for fi77 := 0; fi77 < 13; fi77++ {
            switch 0 {
            case 42:
                var v78 int16 = -301
                _ = v78
            case 127:
                var v79 float64 = (-2.25 - (0.1 / (float64(v8) - 0.1)))
                _ = v79
            default:
                if true {
                    if (!v13) {
                        fmt.Println(v9)
                        v80 := []int{fi77, 0}
                        _ = v80
                    } else {
                        fmt.Println(v76)
                    }
                }
            }
            _ = fi77
        }
        for fi81 := 0; fi81 < 1; fi81++ {
            v1 = v1
            var v82 string = "hello"
            _ = fi81
            _ = v82
        }
    }
    for fi83 := 0; fi83 < 4; fi83++ {
        if v13 {
            var v84 uint16 = 255
            _ = v84
        } else {
            var v85 uint8 = 59
            _ = v85
        }
        v86 := []int{(v9 << ((2) & 63)), (v9 & (fi83 | fi83)), fi83}
        _ = fi83
        _ = v86
    }
    if v13 {
        v76 = 1e10
        v87 := GenS9{F0: v13}
        _ = v87
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v76:", v76)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v76
    _ = v8
    _ = v9
}
