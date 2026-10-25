package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 uint16 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 int64; F1 bool; F2 int }
type GenS4 struct { F0 string; F1 uint32; F2 uint8 }
type GenS5 struct { F0 uint16; F1 int32 }
type GenS6 struct { F0 uint64; F1 uint32; F2 uint8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int64 }
func genFn0() uint {
    if (4294967295 == 42) {
        var v5 uint = 415
        v5 += v5
        _ = v5
    } else {
        var v6 float32 = ((-2.25 + 3.14159) * (((float32(1.5) - 0.0) - (0.1 - 3.14159)) + -2.25))
        var v7 string = "\n\t"
        for fi8 := 0; fi8 < 16; fi8++ {
            var v9 int = (-(fi8))
            switch v7 {
            case "\n\t":
                switch (v7 + "\n\t") {
                case "a":
                    v7 = v7
                case "":
                    switch (^(v9)) {
                    case 42:
                        v7 = ""
                    default:
                        var v10 uint8 = 196
                        _ = v10
                    }
                default:
                    var v11 int = (-(v9))
                    _ = v11
                }
            default:
                fmt.Println(v7)
            }
            _ = fi8
            _ = v9
        }
        _ = v6
        _ = v7
    }
    var v12 uint8 = 190
    if ((!true) || (138 >= 18446744073709551615)) {
        var v13 string = "hello"
        _ = v13
    }
    _ = v12
    return 18446744073709551615
}
func genFn1(p0 uint16, p1 uint64) bool {
    v14 := GenS1{F0: 4294967295, F1: -1, F2: (p0 - (p0 + p0))}
    p1 = (^(p1))
    _ = p0
    _ = p1
    _ = v14
    return (p0 < (p0 / ((p0) | 1)))
}

func main() {
    var v1 int8 = 98
    var v2 uint64 = 18446744073709551615
    var v3 float64 = 0.0
    v4 := make(map[int]uint32)
    v15 := GenS2{F0: uint32(v2), F1: 42}
    var v16 int8 = -128
    v17 := []int{-9223372036854775808}
    if false {
        var v18 uint64 = v2
        var v19 float64 = (v3 / v3)
        v20 := GenS1{F0: genFn0(), F1: 32768, F2: 0}
        _ = v18
        _ = v19
        _ = v20
    }
    if (false || true) {
        v21 := GenS1{F0: genFn0(), F1: 454, F2: 1}
        _ = v21
    } else {
        var v22 int16 = 32767
        v23 := GenS3{F0: 1, F1: true, F2: 4294967295}
        _ = v22
        _ = v23
    }
    v24 := v4[4294967295]
    if true {
        switch ("foo bar" + "") {
        case "a":
            for fi25 := 0; fi25 < 12; fi25++ {
                v1 = 127
                v26 := GenS1{F0: 42, F1: fi25, F2: 594}
                _ = fi25
                _ = v26
            }
        case "":
            v27 := GenS1{F0: 42, F1: 1, F2: 32767}
            _ = v27
        case "hello":
            var v28 int32 = -2147483648
            _ = v28
        default:
            for fi29 := 0; fi29 < 6; fi29++ {
                fmt.Println(v3)
                _ = fi29
            }
        }
        v2 = v2
    } else {
        v24 -= (v24 % ((v24) | 1))
        if (false && true) {
            v1 = (v16 - v16)
            switch -464 {
            case 9223372036854775807:
                var v30 uint = 0
                _ = v30
            }
        }
    }
    if (v24 == (v24 % ((0) | 1))) {
        if false {
            if (9223372036854775807 > int64(v1)) {
                fmt.Println(v16)
                for fi31 := 0; fi31 < 11; fi31++ {
                    switch (fi31 << ((18446744073709551615) & 63)) {
                    case -890:
                        var v32 int = -2147483648
                        _ = v32
                    case 128:
                        v24 = 2147483647
                    case -925:
                        v33 := GenS4{F0: "hello", F1: 65535, F2: 0}
                        _ = v33
                    default:
                        if false {
                            fmt.Println(v3, v2)
                            var v34 float64 = (1.5 - ((v3 - (1e10 / (float64(0.1) * (v3 * float64((((1.5 + (((3.14159 + 0.1) + 1e10) * float32((v3 + 1.5)))) + (float32(((v3 * (-2.25 * 1e10)) * (float64(0.1) + v3))) + -2.25)) / (((1e10 / (float32(3.14159) - ((1e10 + 0.1) + float32((0.1 * 1e10))))) / (float32(-2.25) + 0.1)) + 1.5))))))) + v3))
                            _ = v34
                        } else {
                            v35 := GenS2{F0: (v24 & (v24 >> ((genFn0()) & 63))), F1: (v24 | 2147483648)}
                            _ = v35
                        }
                    }
                    _ = fi31
                }
                _ = len(v4)
            } else {
                switch "\n\t" {
                case "":
                    switch "\n\t" {
                    case "":
                        v2 += 42
                    case "hello":
                        switch 128 {
                        case -9223372036854775808:
                            v36 := GenS2{F0: v24, F1: (v24 % ((42) | 1))}
                            _ = v36
                        case 128:
                            if (false && false) {
                                var v37 string = ("\n\t" + "foo bar")
                                if len(v17) >= 2 {
                                    v17 = v17[1:len(v17)]
                                }
                                for fi38 := 0; fi38 < 19; fi38++ {
                                    var v39 float32 = float32(v3)
                                    if false {
                                        v16 += (-(v1))
                                    } else {
                                        if true {
                                            var v40 float64 = (v3 * 0.1)
                                            v3 = v3
                                            var v41 float32 = (3.14159 / float32(1e10))
                                            _ = v40
                                            _ = v41
                                        }
                                        if ("\n\t" < "foo bar") {
                                            if (v24 <= (v24 / (((v24 << ((0) & 63))) | 1))) {
                                                var v42 uint32 = v24
                                                if false {
                                                    var v43 uint64 = 42
                                                    var v44 bool = false
                                                    var v45 uint8 = 255
                                                    _ = v43
                                                    _ = v44
                                                    _ = v45
                                                } else {
                                                    var v46 uint = genFn0()
                                                    _ = v46
                                                }
                                                v47 := GenS4{F0: v37, F1: (v42 | v42), F2: 178}
                                                _ = v42
                                                _ = v47
                                            } else {
                                                switch fi38 {
                                                case 32768:
                                                    var v48 uint64 = v2
                                                    _ = v48
                                                }
                                            }
                                            v24 += v24
                                        } else {
                                            v39 = v39
                                            v49 := GenS5{F0: 32767, F1: 1}
                                            delete(v4, fi38)
                                            _ = v49
                                        }
                                    }
                                    _ = fi38
                                    _ = v39
                                }
                                _ = v37
                            }
                        }
                    case "a":
                        v4[int(v2)] = 65535
                    default:
                        for fi50 := 0; fi50 < 6; fi50++ {
                            for fi51 := 0; fi51 < 3; fi51++ {
                                if true {
                                    fmt.Println(fi51)
                                    v3 = (v3 + (float64((0.1 / -2.25)) + (v3 - v3)))
                                    fmt.Println(v24, v2)
                                }
                                _ = fi51
                            }
                            fmt.Println(fi50, v16)
                            if true {
                                var v52 string = ("\n\t" + "hello")
                                v2 += (^(v2))
                                for fi53 := 0; fi53 < 4; fi53++ {
                                    switch v52 {
                                    case "":
                                        switch v52 {
                                        case "hello":
                                            v54 := GenS5{F0: 255, F1: -1}
                                            _ = v54
                                        case "\n\t":
                                            v52 = (v52 + "")
                                        }
                                    case "a":
                                        var v55 uint8 = 31
                                        _ = v55
                                    case "foo bar":
                                        if (!(((v16 < v16) || (false && true)) && false)) {
                                            v16 += v1
                                        } else {
                                            v56 := GenS3{F0: 9223372036854775807, F1: false, F2: fi53}
                                            var v57 int8 = 1
                                            _ = v56
                                            _ = v57
                                        }
                                    default:
                                        v2 -= 42
                                        var v58 string = v52
                                        _ = v58
                                    }
                                    var v59 uint32 = 4294967295
                                    if (false && (102 >= uint8(v24))) {
                                        v3 += (v3 * float64(float32(v3)))
                                    } else {
                                        fmt.Println(v52, fi53)
                                    }
                                    v60 := GenS6{F0: (v2 % ((v2) | 1)), F1: v59, F2: 128}
                                    v61 := GenS7{F0: v24}
                                    var v62 uint8 = 127
                                    v63 := GenS2{F0: v59, F1: (v24 | v24)}
                                    _ = fi53
                                    _ = v59
                                    _ = v60
                                    _ = v61
                                    _ = v62
                                    _ = v63
                                }
                                _ = v52
                            }
                            _ = fi50
                        }
                    }
                }
            }
            switch "\n\t" {
            case "\n\t":
                switch 4294967295 {
                case -982:
                    fmt.Println(v24)
                case -1:
                    if true {
                        v16 -= v16
                        for fi64 := 0; fi64 < 1; fi64++ {
                            if (0.0 >= (float32((v3 * float64(3.14159))) - (3.14159 - (3.14159 / (1.5 - (0.1 + 1e10)))))) {
                                for fi65 := 0; fi65 < 19; fi65++ {
                                    var v66 uint64 = v2
                                    var v67 uint8 = 0
                                    _ = fi65
                                    _ = v66
                                    _ = v67
                                }
                            } else {
                                v2 -= (v2 * (v2 * v2))
                                v24 -= (v24 & v24)
                            }
                            var v68 uint8 = 98
                            _ = len(v17)
                            _ = fi64
                            _ = v68
                        }
                        v3 = -2.25
                    }
                case 199:
                    var v69 int8 = v16
                    _ = v69
                }
            case "foo bar":
                var v70 uint = 65535
                _ = v70
            case "a":
                v16 = 123
            }
        }
        v71 := GenS8{F0: v24}
        _ = v71
    } else {
        var v72 uint16 = 65535
        for fi73 := 0; fi73 < 18; fi73++ {
            if (!(18446744073709551615 > 932)) {
                fmt.Println(v1, v1)
                switch (fi73 | fi73) {
                case 128:
                    fmt.Println(v3, v3)
                case 0:
                    v74 := GenS9{F0: 42}
                    _ = v74
                case 242:
                    var v75 bool = (uint32(v2) <= (v24 * (^(v24))))
                    _ = v75
                }
                v3 -= (float64((3.14159 / (1.5 - 3.14159))) * (((float64(((0.1 - 3.14159) - 0.1)) + ((v3 / v3) + (0.1 / 1.0))) * (((((0.1 / (v3 * (v3 * (v3 + v3)))) + 0.1) + v3) + (v3 / v3)) * v3)) / float64(-2.25)))
            }
            var v76 uint = 2
            _ = fi73
            _ = v76
        }
        v77 := GenS9{F0: 197}
        _ = v72
        _ = v77
    }
    if len(v17) > 0 {
        v17[0] = -212
    }
    var v78 uint64 = (v2 * v2)
    v16 = v1
    var v79 int16 = 1000
    v1 = v1
    if true {
        v4[32768] = v24
    } else {
        fmt.Println(v78)
    }
    v3 = (float64(-2.25) - -2.25)
    switch "foo bar" {
    case "a":
        if (genFn0() >= 65535) {
            v1 -= 111
        }
    case "\n\t":
        var v80 float32 = -2.25
        _ = v80
    default:
        if false {
            v24 = (v24 - 1)
            var v81 bool = true
            if (0 < uint16(v24)) {
                var v82 string = "a"
                v3 = 1e10
                for fi83 := 0; fi83 < 14; fi83++ {
                    var v84 int = fi83
                    if len(v17) >= 2 {
                        v17 = v17[1:len(v17)]
                    }
                    _ = fi83
                    _ = v84
                }
                _ = v82
            }
            _ = v81
        }
    }
    if true {
        var v85 bool = (-44 <= -9223372036854775808)
        for fi86 := 0; fi86 < 12; fi86++ {
            switch "\n\t" {
            case "a":
                if len(v17) > 0 {
                    v17[0] = fi86
                }
            case "\n\t":
                if v85 {
                    fmt.Println(v78)
                    var v87 uint8 = 3
                    switch fi86 {
                    case 1:
                        v2 = v78
                    case 42:
                        for fi88 := 0; fi88 < 5; fi88++ {
                            v3 -= (3.14159 - 0.1)
                            _ = fi88
                        }
                    }
                    _ = v87
                }
            default:
                for fi89 := 0; fi89 < 10; fi89++ {
                    fmt.Println(v1)
                    v90 := GenS3{F0: 673, F1: ((!(1e10 <= 3.14159)) || (0.1 == (v3 - (1e10 + 1e10)))), F2: fi86}
                    _ = fi89
                    _ = v90
                }
            }
            _ = fi86
        }
        if v85 {
            v91 := GenS3{F0: 42, F1: (-9223372036854775808 == 42), F2: 2147483647}
            var v92 int16 = (v79 / ((1000) | 1))
            _ = v91
            _ = v92
        } else {
            v78 -= (v78 ^ v78)
        }
        _ = v85
    } else {
        v78 += (v2 | v78)
        v93 := GenS1{F0: uint(v24), F1: -1, F2: 255}
        _ = v93
    }
    for fi94 := 0; fi94 < 18; fi94++ {
        _ = len(v4)
        _ = fi94
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v78
    _ = v79
}
