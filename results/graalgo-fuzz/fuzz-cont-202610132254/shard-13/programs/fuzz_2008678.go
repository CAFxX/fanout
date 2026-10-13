package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 bool; F1 int8 }
type GenS5 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS6 struct { F0 uint64; F1 uint16 }
type GenS7 struct { F0 uint64; F1 uint64; F2 uint32 }
type GenS8 struct { F0 int64 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]bool)
    v3 := []int{371}
    var v4 int = 9223372036854775807
    v5 := GenS1{F0: -9, F1: 8, F2: 319}
    var v6 int16 = 32767
    fmt.Println(v4)
    fmt.Println(v4)
    for fi7 := 0; fi7 < 14; fi7++ {
        v4 += (^(fi7))
        _ = fi7
    }
    var v8 int8 = -1
    for fi9 := 0; fi9 < 10; fi9++ {
        v10 := v1[-9223372036854775808]
        var v11 float32 = 0.1
        for fi12 := 0; fi12 < 15; fi12++ {
            fmt.Println(fi12)
            switch 1 {
            case 4294967295:
                if false {
                    v3 = append(v3, (^(v4)))
                    v3 = append(v3, v4)
                    if true {
                        if (!false) {
                            v4 = 20
                            var v13 uint8 = 2
                            _ = len(v1)
                            _ = v13
                        } else {
                            delete(v2, 65535)
                            v14 := GenS1{F0: (v8 << ((316) & 63)), F1: 42, F2: 88}
                            v15 := GenS2{F0: fi12}
                            _ = len(v1)
                            _ = v14
                            _ = v15
                        }
                        switch (-(v4)) {
                        case 712:
                            if false {
                                var v16 int16 = 1000
                                _ = v16
                            } else {
                                if true {
                                    v6 = (^(v6))
                                    v4 += v4
                                    v11 += 1e10
                                    fmt.Println(v6, fi12, v11)
                                    var v17 bool = (0 < 4294967295)
                                    _ = v17
                                } else {
                                    _ = len(v1)
                                }
                                v3 = append(v3, (fi9 << ((2) & 63)))
                                fmt.Println(v6)
                            }
                        case 579:
                            var v18 bool = true
                            if ((v6 & (v6 << ((42) & 63))) == v6) {
                                var v19 int64 = -9223372036854775808
                                switch ("foo bar" + v10) {
                                case "":
                                    var v20 uint8 = 255
                                    _ = v20
                                }
                                v4 = fi9
                                _ = v19
                            } else {
                                v21 := GenS1{F0: -1, F1: 4294967295, F2: 4294967295}
                                fmt.Println(v6, v8)
                                v10 = v10
                                _ = v21
                            }
                            var v22 uint32 = 4294967295
                            delete(v1, (fi12 + (^(fi9))))
                            if (("foo bar" < (v10 + v10)) || false) {
                                if v18 {
                                    fmt.Println(v22)
                                } else {
                                    var v23 int8 = (v8 - (v8 - v8))
                                    _ = v23
                                }
                                v3 = append(v3, 2147483647)
                            }
                            _ = v18
                            _ = v22
                        case -2147483648:
                            var v24 float32 = (((float32(float64((v11 / v11))) * (0.0 + v11)) - 0.0) * v11)
                            _ = v24
                        default:
                            if true {
                                switch v10 {
                                case "foo bar":
                                    var v25 int64 = -679
                                    _ = v25
                                case "hello":
                                    v26 := GenS3{F0: 42}
                                    _ = v26
                                default:
                                    if (-1 == 0) {
                                        if len(v3) > 0 {
                                            v3[0] = -903
                                        }
                                    } else {
                                        if (v6 > (v6 / ((v6) | 1))) {
                                            var v27 int64 = 777
                                            v6 -= (v6 - v6)
                                            _ = v27
                                        } else {
                                            v28 := GenS4{F0: (202 >= 0), F1: (v8 - (v8 - v8))}
                                            if false {
                                                v10 = v10
                                            }
                                            if false {
                                                var v29 string = v10
                                                if true {
                                                    var v30 uint8 = 213
                                                    var v31 int32 = 408
                                                    if (!(4294967295 != 358)) {
                                                        _ = len(v2)
                                                        v29 = (v29 + v29)
                                                    }
                                                    _ = v30
                                                    _ = v31
                                                } else {
                                                    v4 = fi9
                                                }
                                                v32 := GenS1{F0: -128, F1: 786, F2: 42}
                                                _ = v29
                                                _ = v32
                                            } else {
                                                if true {
                                                    if ((!(v11 < ((3.14159 - v11) - v11))) || (false || ((float64(float32(-2.25)) / (0.1 + -2.25)) != 1.5))) {
                                                        v33 := GenS2{F0: (fi9 % (((v4 << ((4294967295) & 63))) | 1))}
                                                        _ = v33
                                                    }
                                                    if (v6 <= (v6 - (v6 - -1))) {
                                                        if (true && false) {
                                                            var v34 int = v4
                                                            var v35 uint8 = 34
                                                            var v36 int = (v34 | 32768)
                                                            var v37 uint8 = v35
                                                            fmt.Println(v37, v37, v4)
                                                            v38 := v2[65535]
                                                            v39 := GenS5{F0: v6, F1: 4294967295, F2: 881}
                                                            _ = v34
                                                            _ = v35
                                                            _ = v36
                                                            _ = v37
                                                            _ = v38
                                                            _ = v39
                                                        }
                                                        v8 -= v8
                                                    }
                                                    var v40 int8 = 1
                                                    _ = v40
                                                } else {
                                                    var v41 uint64 = 968
                                                    var v42 bool = true
                                                    v11 += 3.14159
                                                    _ = v41
                                                    _ = v42
                                                }
                                                fmt.Println(fi9, v6, v10)
                                                v10 = "foo bar"
                                                _ = len(v2)
                                            }
                                            var v43 float32 = (v11 / v11)
                                            _ = v28
                                            _ = v43
                                        }
                                    }
                                }
                                v11 = 3.14159
                                if (-787 < 0) {
                                    _ = len(v1)
                                }
                            } else {
                                v44 := GenS6{F0: 1, F1: 0}
                                _ = v44
                            }
                            var v45 int16 = v6
                            _ = v45
                        }
                    }
                    switch (v4 / (((v4 >> ((694) & 63))) | 1)) {
                    case -430:
                        if true {
                            var v46 int64 = -9223372036854775808
                            _ = v46
                        }
                    case 9223372036854775807:
                        v10 = v10
                    case -405:
                        var v47 uint64 = 0
                        _ = v47
                    default:
                        if len(v3) > 0 {
                            v3[0] = fi12
                        }
                    }
                } else {
                    v11 -= v11
                }
            default:
                v48 := GenS5{F0: v6, F1: 18446744073709551615, F2: 2}
                v6 -= 626
                _ = v48
            }
            _ = fi12
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    var v49 bool = false
    if v49 {
        switch "a" {
        case "hello":
            fmt.Println(v6, v6, v49)
        case "a":
            switch -1 {
            case -540:
                switch (v4 | (v4 * v4)) {
                case 1:
                    v49 = (!(((!v49) || v49) || v49))
                case 2147483647:
                    if v49 {
                        var v50 int16 = 1000
                        var v51 int32 = 2147483647
                        var v52 uint = 65535
                        _ = v50
                        _ = v51
                        _ = v52
                    } else {
                        v53 := GenS7{F0: 0, F1: 2, F2: 803}
                        var v54 uint64 = 599
                        if ((-(v6)) == v6) {
                            var v55 float32 = 1e10
                            _ = v55
                        } else {
                            var v56 int32 = 0
                            var v57 int16 = v6
                            _ = v56
                            _ = v57
                        }
                        _ = v53
                        _ = v54
                    }
                default:
                    v58 := GenS6{F0: 0, F1: 65535}
                    _ = v58
                }
            }
        case "\n\t":
            var v59 int16 = (v6 % ((v6) | 1))
            _ = v59
        }
        for fi60 := 0; fi60 < 5; fi60++ {
            switch "hello" {
            case "a":
                switch fi60 {
                case 32767:
                    if true {
                        v61 := GenS4{F0: v49, F1: v8}
                        _ = v61
                    } else {
                        var v62 int = 492
                        v8 = -1
                        _ = len(v2)
                        _ = v62
                    }
                case -2147483648:
                    _ = len(v3)
                }
            case "foo bar":
                var v63 uint64 = 759
                _ = v63
            case "":
                switch ("hello" + "hello") {
                case "hello":
                    v6 -= 467
                case "a":
                    for fi64 := 0; fi64 < 6; fi64++ {
                        if v49 {
                            fmt.Println(v6, v49, fi60)
                            if v49 {
                                v65 := GenS3{F0: 1}
                                var v66 int16 = (^(v6))
                                switch (("a" + "") + (("" + "hello") + (((("a" + ("a" + "foo bar")) + "foo bar") + "a") + "hello"))) {
                                case "foo bar":
                                    if v49 {
                                        v49 = false
                                        v4 = fi60
                                        v4 = fi64
                                    }
                                default:
                                    switch "" {
                                    case "":
                                        if v49 {
                                            delete(v1, fi64)
                                        } else {
                                            if len(v3) > 0 {
                                                v3[0] = 0
                                            }
                                        }
                                    case "foo bar":
                                        v8 -= 42
                                    default:
                                        var v67 uint32 = 500
                                        _ = v67
                                    }
                                }
                                _ = v65
                                _ = v66
                            } else {
                                v68 := v1[fi60]
                                _ = v68
                            }
                        }
                        switch "\n\t" {
                        case "hello":
                            v49 = ((0 >= -1) || true)
                        case "a":
                            v69 := GenS6{F0: 4294967295, F1: 592}
                            _ = v69
                        case "foo bar":
                            if v49 {
                                fmt.Println(v4)
                                v49 = v49
                            } else {
                                fmt.Println(fi64, fi60)
                                v8 = (v8 % ((int8(v4)) | 1))
                                v49 = false
                            }
                        }
                        _ = fi64
                    }
                }
            default:
                if false {
                    for fi70 := 0; fi70 < 12; fi70++ {
                        if len(v3) > 0 {
                            v3[0] = fi70
                        }
                        v1[(-(fi60))] = "hello"
                        if true {
                            v49 = v49
                            var v71 uint64 = 1
                            v6 = (v6 + v6)
                            _ = v71
                        } else {
                            switch 0 {
                            case 0:
                                v6 += (v6 % ((-1) | 1))
                            case 128:
                                var v72 uint8 = 213
                                _ = v72
                            case 281:
                                v49 = true
                            }
                            var v73 int64 = 42
                            var v74 uint32 = 721
                            _ = v73
                            _ = v74
                        }
                        var v75 float32 = 0.0
                        var v76 uint8 = 69
                        _ = fi70
                        _ = v75
                        _ = v76
                    }
                } else {
                    fmt.Println(v8)
                }
            }
            for fi77 := 0; fi77 < 18; fi77++ {
                var v78 uint = 65535
                v49 = true
                _ = fi77
                _ = v78
            }
            fmt.Println(fi60, fi60)
            _ = fi60
        }
    } else {
        v8 += (^(v8))
        fmt.Println(v49, v4)
    }
    v79 := GenS8{F0: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v6
    _ = v79
    _ = v8
}
