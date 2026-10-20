package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 int32; F2 int }
type GenS4 struct { F0 int64; F1 int16; F2 uint }
type GenS5 struct { F0 uint32; F1 int32; F2 int16 }
type GenS6 struct { F0 int32; F1 int32 }
type GenS7 struct { F0 int8; F1 int8; F2 string }
type GenS8 struct { F0 uint8 }

func main() {
    var v1 int32 = 1
    var v2 uint64 = 1
    var v3 uint64 = (v2 ^ (v2 << ((815) & 63)))
    v4 := make(map[uint32]int)
    var v5 int64 = 0
    delete(v4, 1)
    for fi6 := 0; fi6 < 11; fi6++ {
        for fi7 := 0; fi7 < 20; fi7++ {
            if true {
                v5 = v5
            }
            v8 := GenS1{F0: v3, F1: (v1 * (v1 - (v1 * (v1 + 340)))), F2: "foo bar"}
            v2 = 2
            _ = fi7
            _ = v8
        }
        if ((160 >= 801) || false) {
            v3 += v3
            var v9 uint32 = 27
            _ = v9
        }
        _ = fi6
    }
    for fi10 := 0; fi10 < 15; fi10++ {
        switch ("foo bar" + "") {
        case "hello":
            v4[4294967295] = (fi10 % (((fi10 << ((42) & 63))) | 1))
        default:
            v11 := v4[387]
            _ = v11
        }
        v12 := GenS1{F0: 599, F1: (v1 & (^(v1))), F2: ("hello" + "foo bar")}
        v5 = v5
        _ = fi10
        _ = v12
    }
    v5 = (v5 - (-(v5)))
    var v13 uint64 = (^(v3))
    switch "a" {
    case "\n\t":
        v14 := []int{4294967295, -9223372036854775808, 42, 42}
        _ = v14
    case "a":
        if false {
            var v15 string = ("foo bar" + "")
            _ = v15
        } else {
            if ((0 > 1) && (0.0 < 1.5)) {
                switch "hello" {
                case "hello":
                    for fi16 := 0; fi16 < 5; fi16++ {
                        for fi17 := 0; fi17 < 20; fi17++ {
                            if (true || (255 <= 32767)) {
                                v18 := GenS2{F0: 532}
                                _ = v18
                            }
                            if false {
                                v5 -= int64(fi16)
                                var v19 uint = 622
                                v1 = (v1 << ((463) & 63))
                                v2 -= v2
                                var v20 int32 = (v1 | (v1 << ((42) & 63)))
                                v3 += v3
                                v21 := GenS3{F0: 0, F1: (v1 + v1), F2: fi16}
                                _ = v19
                                _ = v20
                                _ = v21
                            } else {
                                var v22 bool = (true && true)
                                v3 = v2
                                v23 := GenS3{F0: 127, F1: (v1 >> ((543) & 63)), F2: fi16}
                                _ = v22
                                _ = v23
                            }
                            _ = fi17
                        }
                        fmt.Println(v2, v3)
                        _ = fi16
                    }
                default:
                    fmt.Println(v2, v1)
                }
                if ((^(v1)) > (v1 ^ v1)) {
                    for fi24 := 0; fi24 < 16; fi24++ {
                        var v25 uint32 = 430
                        for fi26 := 0; fi26 < 10; fi26++ {
                            v25 -= v25
                            var v27 uint8 = 1
                            var v28 int8 = 0
                            v28 -= v28
                            var v29 uint64 = v2
                            if false {
                                switch 1 {
                                case 9223372036854775807:
                                    fmt.Println(v13)
                                }
                                v3 = v2
                            }
                            fmt.Println(v25, v29, fi26)
                            _ = fi26
                            _ = v27
                            _ = v28
                            _ = v29
                        }
                        _ = fi24
                        _ = v25
                    }
                    var v30 int8 = -7
                    if (false && false) {
                        switch 1 {
                        case -293:
                            var v31 uint8 = 2
                            _ = v31
                        case 42:
                            switch "a" {
                            case "\n\t":
                                if (!true) {
                                    v3 = v13
                                }
                            }
                        case 4294967295:
                            _ = len(v4)
                        default:
                            if (true || ((0.1 / (0.1 * 0.1)) == 1e10)) {
                                v4[65535] = 112
                                switch "\n\t" {
                                case "\n\t":
                                    fmt.Println(v1, v30, v3)
                                case "a":
                                    var v32 float32 = (-2.25 + (0.1 - ((0.0 - 1.5) - ((1e10 / ((float32(1e10) - ((1e10 + (-2.25 * ((0.1 * float32(float64(((3.14159 / (0.0 - 0.1)) / 1.0)))) - 0.0))) - (0.1 + 0.1))) * (0.1 / ((0.1 - 0.1) - (0.0 - 0.1))))) * (-2.25 * float32(-2.25))))))
                                    _ = v32
                                case "hello":
                                    if false {
                                        var v33 uint = 2
                                        var v34 int = -9223372036854775808
                                        _ = v33
                                        _ = v34
                                    }
                                }
                                for fi35 := 0; fi35 < 14; fi35++ {
                                    _ = len(v4)
                                    switch ("\n\t" + "\n\t") {
                                    case "foo bar":
                                        for fi36 := 0; fi36 < 6; fi36++ {
                                            v37 := []int{(fi36 >> ((1) & 63)), (^(fi35)), fi35, fi36}
                                            if true {
                                                if false {
                                                    var v38 float32 = (((1e10 - (0.1 - (0.0 / 1.5))) / 1.5) + ((0.1 - ((3.14159 - 0.0) + (1e10 / -2.25))) + 3.14159))
                                                    var v39 float64 = (3.14159 + 1.5)
                                                    _ = v38
                                                    _ = v39
                                                }
                                                v5 = (v5 / ((9223372036854775807) | 1))
                                                switch fi35 {
                                                case -229:
                                                    var v40 float64 = float64(float32((float64(0.1) + 1e10)))
                                                    _ = v40
                                                case -892:
                                                    fmt.Println(v5)
                                                    var v41 int8 = v30
                                                    v2 -= (v13 / (((v3 & v2)) | 1))
                                                    fmt.Println(v30)
                                                    _ = v41
                                                case 128:
                                                    fmt.Println(v13, v30, fi35)
                                                }
                                                var v42 float32 = float32(1.5)
                                                _ = v42
                                            }
                                            _ = fi36
                                            _ = v37
                                        }
                                    case "":
                                        for fi43 := 0; fi43 < 2; fi43++ {
                                            if false {
                                                var v44 uint64 = 288
                                                if (!(85 != 47)) {
                                                    if true {
                                                        var v45 int16 = 39
                                                        _ = v45
                                                    } else {
                                                        switch -1 {
                                                        case 128:
                                                            var v46 uint8 = 0
                                                            _ = v46
                                                        case 0:
                                                            v3 = v13
                                                            v47 := GenS4{F0: (v5 | v5), F1: -483, F2: 18446744073709551615}
                                                            v48 := GenS4{F0: v5, F1: -1, F2: 42}
                                                            _ = v47
                                                            _ = v48
                                                        default:
                                                            switch ("" + "foo bar") {
                                                            case "a":
                                                                v2 = v44
                                                            case "":
                                                                var v49 int8 = v30
                                                                _ = v49
                                                            default:
                                                                v50 := GenS5{F0: 42, F1: (-(v1)), F2: -21}
                                                                _ = v50
                                                            }
                                                        }
                                                        v51 := GenS6{F0: (v1 % ((1) | 1)), F1: v1}
                                                        v1 = (-(v1))
                                                        _ = v51
                                                    }
                                                } else {
                                                    switch (fi35 << ((4294967295) & 63)) {
                                                    case -2147483648:
                                                        v52 := GenS1{F0: 42, F1: (-(v1)), F2: "hello"}
                                                        _ = v52
                                                    case 2147483647:
                                                        fmt.Println(fi43, fi35, v5)
                                                        if (v30 < v30) {
                                                            fmt.Println(v2, fi43)
                                                            if true {
                                                                v53 := GenS6{F0: (v1 * (v1 ^ -2147483648)), F1: v1}
                                                                var v54 int32 = v1
                                                                v55 := GenS5{F0: uint32(v44), F1: (v1 | v54), F2: 32767}
                                                                v56 := []int{int(v5), (fi35 | fi35), (fi43 ^ (fi35 << ((42) & 63))), (fi43 * fi35)}
                                                                _ = v53
                                                                _ = v54
                                                                _ = v55
                                                                _ = v56
                                                            } else {
                                                                var v57 bool = (true || ((v30 + (v30 / ((v30) | 1))) == (v30 + (v30 & v30))))
                                                                _ = v57
                                                            }
                                                            var v58 int16 = -1
                                                            _ = v58
                                                        }
                                                        fmt.Println(v3, v3)
                                                    case -1:
                                                        switch (fi35 / ((fi43) | 1)) {
                                                        case -873:
                                                            var v59 int16 = -1
                                                            _ = v59
                                                        default:
                                                            var v60 uint16 = 346
                                                            _ = v60
                                                        }
                                                    }
                                                    fmt.Println(v44)
                                                    if true {
                                                        v61 := []int{4294967295}
                                                        var v62 int32 = (v1 % ((594) | 1))
                                                        v13 = 18446744073709551615
                                                        _ = v61
                                                        _ = v62
                                                    }
                                                }
                                                _ = v44
                                            } else {
                                                if (("a" >= ("a" + ("a" + ""))) && false) {
                                                    var v63 uint = 0
                                                    v13 = (^(v3))
                                                    _ = v63
                                                } else {
                                                    v30 = v30
                                                    v64 := GenS5{F0: uint32(v2), F1: v1, F2: 42}
                                                    _ = v64
                                                }
                                            }
                                            if true {
                                                v65 := GenS7{F0: -1, F1: -1, F2: (("a" + ("" + "\n\t")) + "\n\t")}
                                                _ = v65
                                            }
                                            _ = fi43
                                        }
                                    }
                                    _ = fi35
                                }
                            } else {
                                for fi66 := 0; fi66 < 8; fi66++ {
                                    var v67 float32 = float32(3.14159)
                                    v30 -= v30
                                    _ = fi66
                                    _ = v67
                                }
                            }
                        }
                        delete(v4, 1)
                        v68 := GenS2{F0: 65535}
                        _ = v68
                    } else {
                        if true {
                            switch -845 {
                            case 220:
                                var v69 float64 = 0.0
                                _ = v69
                            default:
                                v3 += (^(v3))
                            }
                            v70 := GenS8{F0: 255}
                            var v71 int8 = (v30 + (-(v30)))
                            _ = v70
                            _ = v71
                        }
                        v72 := GenS2{F0: 255}
                        _ = v72
                    }
                    _ = v30
                }
            } else {
                if (1.5 >= 1e10) {
                    v73 := v4[4294967295]
                    switch "\n\t" {
                    case "a":
                        v74 := v4[695]
                        _ = v74
                    case "":
                        v75 := []int{}
                        _ = v75
                    }
                    _ = v73
                }
                v76 := []int{4294967295, -772, 2147483647, 127}
                var v77 uint16 = 255
                _ = v76
                _ = v77
            }
            if false {
                var v78 string = ""
                fmt.Println(v13)
                _ = v78
            }
        }
    default:
        v79 := []int{128, 1, 322}
        _ = v79
    }
    _ = len(v4)
    v80 := []int{-2147483648, 127, -2147483648}
    v81 := v4[4294967295]
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v80:", len(v80), v80)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v80
    _ = v81
}
