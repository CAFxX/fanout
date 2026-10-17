package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int64; F1 uint32; F2 int16 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int32; F1 uint }
type GenS9 struct { F0 int16; F1 int16 }
func genFn0(p0 string) string {
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 uint8 = 147
        if (255 <= 4294967295) {
            var v8 int64 = -9223372036854775808
            for fi9 := 0; fi9 < 8; fi9++ {
                switch p0 {
                case "":
                    v7 += v7
                case "foo bar":
                    if (!false) {
                        v10 := []int{-9223372036854775808, 2147483647}
                        _ = v10
                    }
                    v11 := GenS1{F0: fi6, F1: -1}
                    if false {
                        v12 := []int{(-(fi6)), fi6, (-(fi6))}
                        _ = v12
                    } else {
                        fmt.Println(v7, fi6)
                        v8 = -9223372036854775808
                        v13 := []int{32768, fi6}
                        _ = v13
                    }
                    _ = v11
                }
                _ = fi9
            }
            _ = v8
        }
        fmt.Println(v7)
        _ = fi6
        _ = v7
    }
    var v14 int16 = 1
    if true {
        fmt.Println(v14, p0)
    } else {
        for fi15 := 0; fi15 < 20; fi15++ {
            if (0 != 600) {
                for fi16 := 0; fi16 < 13; fi16++ {
                    switch (p0 + "foo bar") {
                    case "a":
                        var v17 uint8 = 38
                        _ = v17
                    case "":
                        v18 := []int{fi16, fi16, fi15, fi16}
                        _ = v18
                    default:
                        fmt.Println(p0, p0)
                    }
                    v19 := []int{fi15, fi16, fi16}
                    fmt.Println(fi15, v14, fi16)
                    var v20 uint = 295
                    var v21 uint64 = 18446744073709551615
                    _ = fi16
                    _ = v19
                    _ = v20
                    _ = v21
                }
            }
            for fi22 := 0; fi22 < 2; fi22++ {
                v23 := GenS2{F0: 4294967295, F1: 0}
                _ = fi22
                _ = v23
            }
            var v24 int16 = v14
            _ = fi15
            _ = v24
        }
        v14 = 32767
    }
    v25 := []int{}
    v14 += v14
    _ = p0
    _ = v14
    _ = v25
    return p0
}

func main() {
    var v1 bool = false
    v2 := []int{184}
    var v3 uint = 2
    v4 := make(map[uint32]string)
    v5 := []uint32{4294967295, 742, 901}
    if (42 != 2147483648) {
        v3 += v3
        v3 = (v3 ^ v3)
        var v26 uint8 = 127
        _ = v26
    }
    var v27 bool = (!v1)
    fmt.Println(v1, v1, v27)
    if false {
        v1 = v27
    }
    fmt.Println(v1, v1, v27)
    for fi28 := 0; fi28 < 6; fi28++ {
        if true {
            if v27 {
                var v29 uint = 2
                for fi30 := 0; fi30 < 5; fi30++ {
                    if (fi28 >= (fi28 * fi30)) {
                        var v31 uint32 = 2147483647
                        v32 := GenS1{F0: (fi30 / ((int(v31)) | 1)), F1: -565}
                        fmt.Println(v29, v1, v29)
                        _ = v31
                        _ = v32
                    } else {
                        v3 -= v3
                    }
                    var v33 int8 = 81
                    _ = fi30
                    _ = v33
                }
                if false {
                    fmt.Println(fi28, v27, v1)
                    v29 = 2
                }
                _ = v29
            }
            for fi34 := 0; fi34 < 9; fi34++ {
                v1 = v1
                fmt.Println(v27, v1)
                _ = fi34
            }
        } else {
            var v35 float32 = (0.0 + 1.5)
            _ = len(v4)
            delete(v4, 1)
            _ = v35
        }
        for fi36 := 0; fi36 < 15; fi36++ {
            if v27 {
                switch (fi36 >> (((v3 ^ v3)) & 63)) {
                case 460:
                    if v27 {
                        fmt.Println(v27)
                    } else {
                        var v37 int32 = 0
                        _ = len(v4)
                        var v38 int16 = -94
                        _ = v37
                        _ = v38
                    }
                case -9223372036854775808:
                    _ = len(v5)
                case 9223372036854775807:
                    if (v1 || (v27 || (191 != 42))) {
                        v3 -= v3
                    } else {
                        switch (fi28 ^ (fi28 - fi28)) {
                        case 4294967295:
                            if len(v5) > 0 {
                                v5[0] = 746
                            }
                        case -9223372036854775808:
                            v27 = (!(0 < -1))
                        }
                        fmt.Println(v1)
                    }
                }
            }
            if v27 {
                if true {
                    if (2 > 1) {
                        v1 = (42 >= int32(v3))
                        if v27 {
                            v39 := GenS3{F0: -1}
                            _ = v39
                        } else {
                            if v27 {
                                v1 = v1
                                var v40 int64 = -1
                                _ = v40
                            }
                            v41 := GenS3{F0: 2147483647}
                            _ = v41
                        }
                        var v42 int32 = 42
                        _ = v42
                    }
                    v3 -= 343
                    _ = len(v4)
                }
            } else {
                if (int64(v3) <= 7) {
                    var v43 bool = false
                    var v44 int16 = 1
                    v3 = v3
                    _ = v43
                    _ = v44
                } else {
                    if v27 {
                        var v45 bool = (1 < -9223372036854775808)
                        fmt.Println(v1, fi36, v45)
                        v45 = true
                        _ = v45
                    }
                }
                if v27 {
                    var v46 uint32 = 42
                    if v1 {
                        if (!v1) {
                            fmt.Println(fi36, v3, v46)
                            fmt.Println(v46)
                        } else {
                            v47 := GenS4{F0: -128, F1: 81}
                            v48 := GenS4{F0: 3, F1: 0}
                            var v49 int64 = 251
                            _ = v47
                            _ = v48
                            _ = v49
                        }
                        fmt.Println(v3)
                    }
                    _ = v46
                } else {
                    if (!((fi28 >> (((v3 & (v3 - 612))) & 63)) > 127)) {
                        var v50 int = fi36
                        v2 = append(v2, fi36)
                        v27 = true
                        _ = v50
                    }
                    v4[635] = ("hello" + "hello")
                    if v27 {
                        v51 := GenS4{F0: -128, F1: -128}
                        switch "" {
                        case "hello":
                            fmt.Println(v27, v3, fi28)
                        case "":
                            if (1 < -128) {
                                var v52 uint8 = 127
                                _ = len(v4)
                                if (v27 || (!false)) {
                                    v3 = (v3 % ((1) | 1))
                                    _ = len(v2)
                                } else {
                                    var v53 int32 = 2147483647
                                    v53 = v53
                                    if (((((3.14159 + 1e10) * (0.0 + (-2.25 - 0.0))) * (1.5 + -2.25)) >= ((3.14159 + -2.25) * (3.14159 - 0.0))) && (v27 || true)) {
                                        var v54 uint = (v3 >> ((v3) & 63))
                                        _ = v54
                                    } else {
                                        var v55 float64 = (float64((((float32((((3.14159 + 0.1) - (float64(1e10) + 0.1)) + ((float64((3.14159 - 3.14159)) * 0.1) * ((1e10 - (((-2.25 / (float64((1e10 + -2.25)) + 3.14159)) * ((float64((-2.25 / ((float32(0.0) / 1.5) + 1.5))) / 3.14159) - 1e10)) - (((0.1 - 0.1) * ((0.0 / 1.0) - 1.5)) * float64((float32(1.5) / 1e10))))) - float64(1.5))))) * (float32(1.5) - 0.0)) + 1.5) - ((1e10 + 3.14159) / 1.5))) + ((0.1 / (3.14159 - 1.5)) - 3.14159))
                                        _ = v55
                                    }
                                    _ = v53
                                }
                                _ = v52
                            } else {
                                switch "foo bar" {
                                case "":
                                    delete(v4, 0)
                                case "\n\t":
                                    switch fi28 {
                                    case -1:
                                        v3 = v3
                                    }
                                case "a":
                                    fmt.Println(v3)
                                default:
                                    if false {
                                        var v56 float64 = 3.14159
                                        v56 = 1.5
                                        v56 = 1.5
                                        var v57 uint16 = 32767
                                        _ = v56
                                        _ = v57
                                    } else {
                                        switch ("foo bar" + "foo bar") {
                                        case "foo bar":
                                            var v58 uint32 = 0
                                            _ = v58
                                        case "\n\t":
                                            if v1 {
                                                v2 = append(v2, fi28)
                                            } else {
                                                v3 -= 947
                                                switch "" {
                                                case "":
                                                    var v59 float32 = 3.14159
                                                    _ = v59
                                                }
                                                var v60 string = ("hello" + "foo bar")
                                                _ = v60
                                            }
                                            switch (fi28 + (fi36 << ((v3) & 63))) {
                                            case -753:
                                                var v61 bool = v27
                                                _ = v61
                                            default:
                                                fmt.Println(v27, v27, v27)
                                                v62 := GenS5{F0: "foo bar"}
                                                _ = v62
                                            }
                                            _ = len(v5)
                                        default:
                                            v63 := GenS6{F0: 0, F1: 1, F2: -1}
                                            _ = v63
                                        }
                                        v64 := v4[42]
                                        _ = v64
                                    }
                                    var v65 int16 = -1
                                    var v66 float64 = (0.0 - (1.5 - 3.14159))
                                    _ = v65
                                    _ = v66
                                }
                                v27 = v27
                                switch (fi28 * (fi28 << ((v3) & 63))) {
                                case 255:
                                    var v67 float32 = 3.14159
                                    _ = v67
                                case 4294967295:
                                    v3 = v3
                                case 127:
                                    fmt.Println(fi36, v1, fi36)
                                }
                            }
                        default:
                            switch "a" {
                            case "hello":
                                if true {
                                    v1 = (1.5 > 3.14159)
                                }
                            case "":
                                if (-32768 >= -741) {
                                    var v68 int16 = -32768
                                    var v69 uint = v3
                                    var v70 bool = false
                                    _ = v68
                                    _ = v69
                                    _ = v70
                                } else {
                                    if v1 {
                                        v3 -= 2
                                        var v71 int32 = 1
                                        var v72 float64 = ((0.0 / 1.5) + 0.1)
                                        _ = v71
                                        _ = v72
                                    } else {
                                        v3 += 333
                                    }
                                    fmt.Println(v27, fi28)
                                    v3 -= v3
                                }
                            case "\n\t":
                                if (1 > 1) {
                                    delete(v4, 65535)
                                    var v73 uint64 = 4294967295
                                    v74 := GenS7{F0: (^(v73))}
                                    _ = v73
                                    _ = v74
                                }
                            default:
                                var v75 int8 = -128
                                _ = v75
                            }
                        }
                        v76 := GenS4{F0: -110, F1: 42}
                        _ = v51
                        _ = v76
                    } else {
                        _ = len(v5)
                        v77 := GenS1{F0: (fi28 % ((int(v3)) | 1)), F1: 1000}
                        v27 = v1
                        v3 -= (v3 << ((65535) & 63))
                        _ = v77
                    }
                    var v78 int32 = 42
                    fmt.Println(v78, v78)
                    _ = v78
                }
            }
            fmt.Println(fi36)
            v27 = true
            _ = fi36
        }
        if true {
            if v1 {
                var v79 float32 = 1e10
                _ = v79
            }
        }
        _ = fi28
    }
    var v80 float32 = -2.25
    if (!(true && v1)) {
        v81 := v4[0]
        var v82 int8 = 1
        _ = v81
        _ = v82
    }
    for fi83 := 0; fi83 < 3; fi83++ {
        v3 = 255
        _ = fi83
    }
    v84 := GenS1{F0: -5, F1: -1}
    v85 := GenS8{F0: 1, F1: v3}
    v27 = true
    for fi86 := 0; fi86 < 11; fi86++ {
        switch (fi86 % ((-612) | 1)) {
        case 0:
            v87 := GenS9{F0: 1000, F1: 42}
            _ = v87
        case 42:
            for fi88 := 0; fi88 < 6; fi88++ {
                fmt.Println(fi88)
                _ = fi88
            }
        default:
            fmt.Println(fi86, v27)
        }
        _ = fi86
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v80
    _ = v84
    _ = v85
}
