package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 int64; F2 bool }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 uint16; F1 int8; F2 int8 }
type GenS5 struct { F0 int }
func genFn0(p0 uint32) bool {
    p0 -= (p0 + (^(p0)))
    p0 += p0
    fmt.Println(p0)
    var v7 uint64 = 18446744073709551615
    _ = p0
    _ = v7
    return false
}
func genFn1() uint32 {
    v8 := GenS1{F0: 2}
    v9 := []int{}
    switch 988 {
    case 998:
        var v10 int8 = 1
        _ = v10
    case 42:
        var v11 uint64 = 706
        _ = v11
    default:
        for fi12 := 0; fi12 < 3; fi12++ {
            var v13 uint32 = 31
            if len(v9) > 0 {
                v9[0] = (fi12 ^ fi12)
            }
            for fi14 := 0; fi14 < 14; fi14++ {
                fmt.Println(fi12, v13, v13)
                fmt.Println(fi12, fi14, v13)
                v15 := GenS1{F0: 406}
                _ = fi14
                _ = v15
            }
            _ = fi12
            _ = v13
        }
    }
    _ = v8
    _ = v9
    return 747
}

func main() {
    v1 := []uint32{0, 2147483647}
    var v2 float32 = -2.25
    var v3 float64 = float64(1.5)
    var v4 int16 = 158
    v5 := []uint32{373, 2, 2147483648, 2}
    var v6 uint8 = 78
    switch "\n\t" {
    case "hello":
        switch (((("hello" + "\n\t") + "\n\t") + "foo bar") + "a") {
        case "\n\t":
            v6 = 0
        case "foo bar":
            if false {
                v3 = -2.25
                for fi16 := 0; fi16 < 11; fi16++ {
                    v2 -= v2
                    for fi17 := 0; fi17 < 8; fi17++ {
                        v6 += (v6 + (^(v6)))
                        switch "" {
                        case "foo bar":
                            fmt.Println(fi16, v3, fi17)
                        case "hello":
                            v4 = 1000
                            if true {
                                fmt.Println(v2, v3)
                                var v18 string = "foo bar"
                                if false {
                                    var v19 float64 = v3
                                    _ = v19
                                } else {
                                    var v20 string = ("\n\t" + ((v18 + "\n\t") + v18))
                                    switch ((((v20 + "hello") + v18) + v20) + v20) {
                                    case "":
                                        v20 = (v20 + (((v18 + "\n\t") + v18) + v18))
                                    case "a":
                                        var v21 int32 = -2147483648
                                        v20 = v18
                                        _ = v21
                                    case "\n\t":
                                        var v22 bool = (508 >= fi17)
                                        _ = v22
                                    }
                                    if true {
                                        if false {
                                            v3 = 1.5
                                            if false {
                                                v4 = v4
                                                fmt.Println(v3, v20, fi16)
                                                v20 = "\n\t"
                                            } else {
                                                var v23 int8 = -1
                                                v1 = append(v1, 595)
                                                _ = v23
                                            }
                                            if len(v1) > 0 {
                                                v1[0] = 270
                                            }
                                        }
                                        switch v20 {
                                        case "hello":
                                            if true {
                                                v6 = 81
                                                var v24 uint32 = 2
                                                v2 = 1.5
                                                _ = v24
                                            } else {
                                                var v25 int16 = (v4 * -316)
                                                v26 := GenS1{F0: uint(v6)}
                                                if false {
                                                    var v27 float32 = ((v2 * ((v2 / v2) * 0.0)) - (float32(3.14159) + (3.14159 + 0.0)))
                                                    _ = v27
                                                } else {
                                                    fmt.Println(v6)
                                                }
                                                _ = v25
                                                _ = v26
                                            }
                                            switch v20 {
                                            case "a":
                                                v4 = 0
                                            case "\n\t":
                                                fmt.Println(v4)
                                            case "hello":
                                                v28 := GenS1{F0: 18446744073709551615}
                                                _ = v28
                                            }
                                            if len(v1) > 0 {
                                                v1[0] = 42
                                            }
                                        case "foo bar":
                                            if true {
                                                if (int16(fi16) != (v4 / (((v4 << ((985) & 63))) | 1))) {
                                                    v3 = ((v3 + (3.14159 - 0.1)) / float64(v2))
                                                    v4 -= v4
                                                    v2 += ((v2 * v2) * v2)
                                                } else {
                                                    switch fi16 {
                                                    case 127:
                                                        var v29 int = fi16
                                                        _ = v29
                                                    }
                                                    switch (fi17 ^ (fi17 * fi16)) {
                                                    case 9223372036854775807:
                                                        v20 = (("hello" + ("" + v18)) + ((("\n\t" + "\n\t") + (("" + "") + (("hello" + "a") + v20))) + (v18 + ("hello" + v18))))
                                                    default:
                                                        switch ("\n\t" + "") {
                                                        case "hello":
                                                            var v30 uint64 = 42
                                                            if ((true || false) || (255 <= 89)) {
                                                                fmt.Println(v18, fi16, v4)
                                                                v31 := GenS1{F0: 4294967295}
                                                                if (1 != 1) {
                                                                    var v32 int64 = 42
                                                                    _ = v32
                                                                }
                                                                _ = v31
                                                            }
                                                            _ = len(v5)
                                                            _ = v30
                                                        case "\n\t":
                                                            v20 = v18
                                                        }
                                                    }
                                                    var v33 int32 = -378
                                                    _ = v33
                                                }
                                            }
                                        default:
                                            fmt.Println(v2, v18, v4)
                                        }
                                        v34 := GenS1{F0: 65535}
                                        v1 = append(v1, 0)
                                        switch fi17 {
                                        case 9223372036854775807:
                                            fmt.Println(v18, v3)
                                        case 1:
                                            switch (fi17 | -1) {
                                            case 32768:
                                                var v35 uint8 = 28
                                                _ = v35
                                            case 136:
                                                v4 -= (v4 - v4)
                                            case -914:
                                                var v36 bool = false
                                                _ = v36
                                            default:
                                                v4 = -32768
                                            }
                                        case 127:
                                            var v37 float32 = 0.0
                                            _ = v37
                                        }
                                        _ = v34
                                    } else {
                                        v38 := GenS1{F0: 675}
                                        _ = v38
                                    }
                                    _ = v20
                                }
                                _ = v18
                            } else {
                                fmt.Println(fi16)
                                switch ("" + "") {
                                case "\n\t":
                                    if false {
                                        var v39 float64 = (float64((v2 - v2)) - v3)
                                        _ = v39
                                    } else {
                                        var v40 float32 = v2
                                        var v41 string = ""
                                        _ = v40
                                        _ = v41
                                    }
                                default:
                                    v5 = append(v5, 2147483647)
                                }
                                var v42 float32 = 3.14159
                                _ = v42
                            }
                            switch fi16 {
                            case -2147483648:
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            case 4294967295:
                                switch (("hello" + "a") + ("a" + "\n\t")) {
                                case "foo bar":
                                    if true {
                                        v5 = append(v5, 744)
                                        var v43 uint64 = 1
                                        _ = v43
                                    }
                                case "":
                                    v4 += v4
                                default:
                                    v44 := GenS2{F0: -128, F1: 9223372036854775807, F2: false}
                                    _ = v44
                                }
                            case 127:
                                v45 := GenS2{F0: 69, F1: 0, F2: false}
                                _ = v45
                            default:
                                if false {
                                    v4 += v4
                                    if ((-339 >= -885) && ((0.0 == -2.25) && ((1 > 18446744073709551615) || true))) {
                                        v6 = (^(v6))
                                        fmt.Println(fi17, v2)
                                    } else {
                                        v46 := GenS2{F0: 0, F1: -1, F2: true}
                                        var v47 int8 = 87
                                        switch fi16 {
                                        case -91:
                                            var v48 int8 = -128
                                            _ = v48
                                        case 128:
                                            v3 = v3
                                        case 9223372036854775807:
                                            fmt.Println(v2)
                                        }
                                        _ = v46
                                        _ = v47
                                    }
                                    if (!((!true) || false)) {
                                        switch "" {
                                        case "a":
                                            switch ("" + ("\n\t" + "hello")) {
                                            case "a":
                                                var v49 int64 = 414
                                                _ = v49
                                            }
                                        case "":
                                            fmt.Println(v6)
                                            if ((105 == -1) && (v6 < (v6 >> ((0) & 63)))) {
                                                var v50 uint64 = 2
                                                _ = v50
                                            } else {
                                                var v51 int8 = 100
                                                if (("foo bar" + "\n\t") < ("foo bar" + "hello")) {
                                                    var v52 bool = (255 <= 65535)
                                                    var v53 int8 = (v51 + (-(v51)))
                                                    var v54 int32 = 2147483647
                                                    _ = v52
                                                    _ = v53
                                                    _ = v54
                                                }
                                                _ = v51
                                            }
                                        case "foo bar":
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            v3 = -2.25
                                        default:
                                            var v55 int16 = (v4 % (((v4 ^ -1)) | 1))
                                            _ = v55
                                        }
                                        var v56 uint16 = 2
                                        _ = v56
                                    }
                                }
                            }
                            var v57 uint8 = (^(v6))
                            _ = v57
                        case "\n\t":
                            v6 -= (v6 >> ((65535) & 63))
                            var v58 int = (fi17 - (fi16 >> ((2) & 63)))
                            _ = v58
                        }
                        if (42 == 127) {
                            var v59 bool = true
                            switch fi17 {
                            case 4294967295:
                                var v60 float64 = v3
                                _ = v60
                            }
                            _ = v59
                        } else {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            v1 = append(v1, 149)
                            fmt.Println(fi16, v2, fi17)
                        }
                        _ = fi17
                    }
                    _ = fi16
                }
                var v61 float32 = v2
                _ = v61
            } else {
                v4 = (v4 % ((42) | 1))
                for fi62 := 0; fi62 < 5; fi62++ {
                    var v63 int = fi62
                    var v64 uint16 = 0
                    _ = fi62
                    _ = v63
                    _ = v64
                }
            }
        case "a":
            v65 := GenS2{F0: 127, F1: -9223372036854775808, F2: (515 != 65535)}
            _ = v65
        }
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v66 := GenS1{F0: 1}
    switch 32768 {
    case 32767:
        var v67 string = ("foo bar" + "\n\t")
        _ = v67
    case -377:
        var v68 int64 = -119
        _ = v68
    default:
        v69 := GenS3{F0: 81, F1: 100}
        _ = v69
    }
    fmt.Println(v4)
    switch 127 {
    case 9223372036854775807:
        switch "" {
        case "\n\t":
            if (1.5 <= float32(float64((v2 - (3.14159 / (float32((v3 / 0.1)) * v2)))))) {
                if true {
                    v70 := GenS4{F0: 0, F1: 42, F2: 91}
                    _ = v70
                } else {
                    fmt.Println(v6)
                }
            } else {
                for fi71 := 0; fi71 < 17; fi71++ {
                    v2 = v2
                    _ = fi71
                }
            }
        }
    case -2147483648:
        var v72 int8 = -108
        _ = v72
    case -1:
        fmt.Println(v6)
    }
    switch -1 {
    case 4294967295:
        if len(v5) > 0 {
            v5[0] = genFn1()
        }
    case 42:
        v6 = v6
    default:
        fmt.Println(v4, v3)
    }
    v2 -= 0.0
    var v73 uint16 = 731
    if true {
        var v74 int = 371
        if false {
            for fi75 := 0; fi75 < 16; fi75++ {
                v76 := GenS3{F0: 42, F1: 127}
                _ = fi75
                _ = v76
            }
            var v77 uint = 255
            _ = v77
        }
        _ = v74
    }
    if len(v5) > 0 {
        v5[0] = genFn1()
    }
    var v78 int32 = -16
    switch 788 {
    case -375:
        if len(v1) > 0 {
            v1[0] = 42
        }
    case -494:
        v1 = append(v1, genFn1())
    }
    v79 := GenS5{F0: 363}
    v2 = v2
    var v80 int64 = -1
    var v81 string = "foo bar"
    var v82 uint32 = 0
    fmt.Println(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v73:", v73)
    fmt.Println("v78:", v78)
    fmt.Println("v80:", v80)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v66
    _ = v73
    _ = v78
    _ = v79
    _ = v80
    _ = v81
    _ = v82
}
