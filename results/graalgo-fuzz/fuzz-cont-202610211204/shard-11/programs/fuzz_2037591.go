package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 bool; F2 bool }
type GenS4 struct { F0 int; F1 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int64; F1 string; F2 uint16 }
type GenS8 struct { F0 int; F1 int64; F2 string }
type GenS9 struct { F0 uint8; F1 bool; F2 int16 }
type GenS10 struct { F0 bool; F1 int }
type GenS11 struct { F0 int64; F1 int32 }
func genFn0(p0 bool, p1 int32, p2 int8) int16 {
    for fi4 := 0; fi4 < 3; fi4++ {
        v5 := []int{fi4, (fi4 + (fi4 ^ (fi4 ^ (fi4 + fi4)))), fi4, (fi4 >> ((4294967295) & 63))}
        _ = fi4
        _ = v5
    }
    if false {
        p1 = p1
        for fi6 := 0; fi6 < 5; fi6++ {
            switch fi6 {
            case 32767:
                v7 := GenS1{F0: p1}
                _ = v7
            default:
                var v8 int16 = -1
                _ = v8
            }
            p0 = p0
            _ = fi6
        }
        v9 := []int{365}
        _ = v9
    }
    var v10 uint32 = 2147483648
    var v11 int = 4294967295
    switch "\n\t" {
    case "hello":
        p1 -= (p1 >> ((4294967295) & 63))
    case "a":
        v12 := []int{}
        _ = v12
    case "foo bar":
        p1 = p1
    default:
        for fi13 := 0; fi13 < 15; fi13++ {
            v14 := GenS1{F0: 2147483647}
            p1 = p1
            v15 := GenS1{F0: -2147483648}
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 0
}
func genFn1(p0 uint32) uint16 {
    v16 := []int{670}
    switch 775 {
    case -9223372036854775808:
        v17 := GenS2{F0: (("\n\t" + (("foo bar" + "\n\t") + "foo bar")) + "")}
        _ = v17
    case 206:
        var v18 bool = (p0 == (^(p0)))
        _ = v18
    case 9223372036854775807:
        var v19 float32 = 3.14159
        _ = v19
    default:
        var v20 float32 = (1.5 * float32(((1.5 / float64(1.5)) * 1e10)))
        _ = v20
    }
    v21 := GenS2{F0: "a"}
    _ = p0
    _ = v16
    _ = v21
    return 42
}

func main() {
    var v1 string = "foo bar"
    var v2 uint32 = 265
    var v3 bool = false
    fmt.Println(v2, v2)
    v22 := []int{}
    if false {
        v1 = v1
        if v3 {
            var v23 uint = uint(v2)
            _ = v23
        } else {
            v3 = (!v3)
        }
        v22 = append(v22, -227)
    } else {
        var v24 uint64 = 1
        if v3 {
            v25 := GenS2{F0: "foo bar"}
            fmt.Println(v2, v2, v24)
            _ = v25
        } else {
            fmt.Println(v3, v24)
        }
        _ = v24
    }
    for fi26 := 0; fi26 < 13; fi26++ {
        var v27 int = 200
        switch "" {
        case "":
            if (!v3) {
                var v28 uint16 = 255
                _ = v28
            } else {
                v27 -= (v27 >> ((2) & 63))
                switch "" {
                case "foo bar":
                    v27 -= 9223372036854775807
                case "":
                    v27 = (v27 * (fi26 - (v27 ^ (fi26 | (fi26 & (fi26 - (fi26 | 42)))))))
                default:
                    if v3 {
                        for fi29 := 0; fi29 < 15; fi29++ {
                            v27 = -2147483648
                            _ = fi29
                        }
                    }
                }
                for fi30 := 0; fi30 < 1; fi30++ {
                    v27 -= (fi30 - int(v2))
                    _ = fi30
                }
            }
        }
        _ = fi26
        _ = v27
    }
    if v3 {
        if ((v2 ^ v2) != (v2 >> ((2) & 63))) {
            v2 -= 2147483648
            v22 = append(v22, 32)
        } else {
            var v31 uint64 = 4294967295
            var v32 float32 = (0.1 - 3.14159)
            _ = v31
            _ = v32
        }
    }
    var v33 int32 = 286
    v34 := GenS3{F0: 339, F1: (!(-32768 == -1)), F2: v3}
    v3 = (42 > 80)
    v2 += v2
    switch v1 {
    case "a":
        if v3 {
            var v35 uint64 = 18446744073709551615
            v35 = (v35 - v35)
            fmt.Println(v2, v33, v1)
            _ = v35
        } else {
            v36 := GenS3{F0: 0, F1: (!v3), F2: v3}
            var v37 int = 659
            _ = v36
            _ = v37
        }
    default:
        fmt.Println(v2, v2, v3)
    }
    v38 := GenS4{F0: 2147483647, F1: (("\n\t" + ("" + "foo bar")) + v1)}
    if v3 {
        if v3 {
            _ = len(v22)
            var v39 int16 = 42
            _ = len(v22)
            _ = v39
        } else {
            v3 = false
            v3 = (65535 == 42)
        }
        for fi40 := 0; fi40 < 10; fi40++ {
            var v41 bool = (v3 && v3)
            switch "foo bar" {
            case "hello":
                if v3 {
                    v1 = (v1 + "\n\t")
                    var v42 uint64 = 18446744073709551615
                    _ = v42
                }
            }
            _ = fi40
            _ = v41
        }
        for fi43 := 0; fi43 < 7; fi43++ {
            v2 = v2
            v44 := GenS2{F0: v1}
            _ = fi43
            _ = v44
        }
    } else {
        for fi45 := 0; fi45 < 10; fi45++ {
            var v46 int8 = 127
            _ = fi45
            _ = v46
        }
        var v47 int8 = -1
        for fi48 := 0; fi48 < 16; fi48++ {
            switch (fi48 - (-(fi48))) {
            case -1:
                var v49 float32 = (-2.25 / 1e10)
                _ = v49
            default:
                v3 = (0 <= 42)
            }
            v50 := GenS3{F0: 0, F1: (!v3), F2: true}
            _ = fi48
            _ = v50
        }
        _ = v47
    }
    var v51 int64 = 868
    if ((1 > -28) && (!false)) {
        for fi52 := 0; fi52 < 1; fi52++ {
            switch (fi52 % (((-(fi52))) | 1)) {
            case 4294967295:
                switch (fi52 >> ((255) & 63)) {
                case -57:
                    v3 = v3
                case 42:
                    if ("" == v1) {
                        v1 = ((v1 + "hello") + v1)
                        for fi53 := 0; fi53 < 13; fi53++ {
                            if false {
                                _ = len(v22)
                                if v3 {
                                    v1 = v1
                                } else {
                                    var v54 float64 = 3.14159
                                    _ = v54
                                }
                            } else {
                                var v55 int64 = v51
                                _ = v55
                            }
                            fmt.Println(v3)
                            fmt.Println(fi53)
                            var v56 uint = 0
                            var v57 float64 = (-2.25 * 0.0)
                            _ = fi53
                            _ = v56
                            _ = v57
                        }
                    }
                case 32767:
                    if (1 <= 32767) {
                        v33 = v33
                    } else {
                        if (521 == -1) {
                            for fi58 := 0; fi58 < 8; fi58++ {
                                _ = len(v22)
                                _ = fi58
                            }
                        } else {
                            v33 = (v33 % (((v33 - (v33 - v33))) | 1))
                        }
                        v33 = v33
                        var v59 string = v1
                        _ = v59
                    }
                }
            case 0:
                if (0.0 >= -2.25) {
                    v1 = "a"
                    var v60 int64 = v51
                    _ = v60
                }
            }
            v61 := GenS5{F0: (v51 ^ 9223372036854775807)}
            switch fi52 {
            case 32767:
                if true {
                    for fi62 := 0; fi62 < 14; fi62++ {
                        switch (^(fi52)) {
                        case 32767:
                            switch fi62 {
                            case 127:
                                v63 := GenS2{F0: "\n\t"}
                                _ = v63
                            case -9223372036854775808:
                                if false {
                                    if (420 == 2) {
                                        if (true && (1 > -1)) {
                                            fmt.Println(v51, v33)
                                            fmt.Println(fi52)
                                        }
                                    } else {
                                        var v64 int8 = 100
                                        v65 := GenS3{F0: 0, F1: (v1 == v1), F2: ((v64 * -128) > 100)}
                                        v66 := GenS5{F0: v51}
                                        _ = v64
                                        _ = v65
                                        _ = v66
                                    }
                                    var v67 uint32 = v2
                                    if true {
                                        var v68 string = v1
                                        var v69 uint8 = 255
                                        var v70 int16 = -186
                                        var v71 float64 = 0.0
                                        v72 := GenS6{F0: (-(fi52))}
                                        _ = v68
                                        _ = v69
                                        _ = v70
                                        _ = v71
                                        _ = v72
                                    }
                                    if len(v22) > 0 {
                                        v22[0] = (^(fi62))
                                    }
                                    v73 := GenS7{F0: (v51 << ((0) & 63)), F1: "foo bar", F2: uint16(v67)}
                                    var v74 uint = 945
                                    v67 += (v67 & (v2 & 42))
                                    if (((v74 >= v74) || (1.5 > (3.14159 + 0.0))) && (fi62 <= 510)) {
                                        switch fi52 {
                                        case 370:
                                            v33 -= v33
                                        case 0:
                                            if len(v22) >= 2 {
                                                v22 = v22[1:len(v22)]
                                            }
                                        case 580:
                                            v33 = (v33 ^ v33)
                                        }
                                        var v75 int16 = -227
                                        _ = v75
                                    }
                                    _ = v67
                                    _ = v73
                                    _ = v74
                                } else {
                                    _ = len(v22)
                                    fmt.Println(v33, fi52)
                                }
                            default:
                                var v76 int8 = 0
                                _ = v76
                            }
                        }
                        var v77 float32 = (0.1 / ((0.1 / (3.14159 + 1e10)) / ((0.1 / 0.1) * float32(1.5))))
                        _ = fi62
                        _ = v77
                    }
                    var v78 int64 = v51
                    var v79 uint8 = 131
                    _ = v78
                    _ = v79
                }
            case 193:
                switch ((v1 + v1) + (v1 + v1)) {
                case "\n\t":
                    v80 := GenS8{F0: (-(fi52)), F1: (v51 - 692), F2: "foo bar"}
                    _ = v80
                default:
                    v51 -= 42
                }
            case -9223372036854775808:
                _ = len(v22)
            default:
                var v81 float64 = (0.1 + ((1.5 + (1.5 * -2.25)) * 0.1))
                _ = v81
            }
            _ = fi52
            _ = v61
        }
    } else {
        switch v1 {
        case "a":
            var v82 int32 = v33
            _ = v82
        case "\n\t":
            v2 = (v2 << ((4294967295) & 63))
        case "foo bar":
            for fi83 := 0; fi83 < 9; fi83++ {
                switch fi83 {
                case -1:
                    if ((v51 | (v51 << ((42) & 63))) <= -817) {
                        fmt.Println(v33, v3)
                        v33 -= v33
                    } else {
                        for fi84 := 0; fi84 < 10; fi84++ {
                            v85 := GenS9{F0: 255, F1: (127 <= 57), F2: -1}
                            var v86 int16 = genFn0(false, -68, int8(v2))
                            _ = fi84
                            _ = v85
                            _ = v86
                        }
                        v87 := GenS1{F0: 0}
                        _ = v87
                    }
                case 2147483647:
                    switch fi83 {
                    case -656:
                        for fi88 := 0; fi88 < 4; fi88++ {
                            v89 := GenS10{F0: true, F1: fi83}
                            var v90 int16 = 702
                            _ = fi88
                            _ = v89
                            _ = v90
                        }
                    case 32767:
                        fmt.Println(v51, fi83)
                    case 32768:
                        if len(v22) >= 2 {
                            v22 = v22[1:len(v22)]
                        }
                    }
                }
                _ = fi83
            }
        default:
            v3 = (0 >= 42)
        }
        v2 += (v2 - v2)
        switch 128 {
        case 14:
            for fi91 := 0; fi91 < 3; fi91++ {
                var v92 uint64 = 984
                _ = fi91
                _ = v92
            }
        case 916:
            var v93 int32 = -2147483648
            _ = v93
        case -1:
            for fi94 := 0; fi94 < 2; fi94++ {
                v95 := GenS11{F0: (^(v51)), F1: 0}
                v96 := GenS8{F0: 127, F1: 9223372036854775807, F2: (v1 + (v1 + "hello"))}
                _ = fi94
                _ = v95
                _ = v96
            }
        default:
            var v97 int8 = 111
            _ = v97
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v33
    _ = v34
    _ = v38
    _ = v51
}
