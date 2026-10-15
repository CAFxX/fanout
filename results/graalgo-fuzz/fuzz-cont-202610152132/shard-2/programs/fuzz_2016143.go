package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 bool; F2 int }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int32; F1 int64; F2 uint32 }
type GenS6 struct { F0 string; F1 uint64; F2 uint }
type GenS7 struct { F0 int32; F1 string; F2 int64 }
type GenS8 struct { F0 bool; F1 uint16; F2 string }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint64; F1 bool }
type GenS11 struct { F0 int; F1 uint16; F2 int64 }

func main() {
    var v1 float64 = (3.14159 + (((0.1 / 1e10) + 1e10) / float64(0.1)))
    var v2 int = 0
    v3 := []int{(v2 ^ 32768)}
    v4 := make(map[int]bool)
    var v5 int32 = 659
    v6 := make(map[int]int)
    for fi7 := 0; fi7 < 1; fi7++ {
        for fi8 := 0; fi8 < 2; fi8++ {
            if (!(!(-2.25 != (1.5 + (float32(1e10) / 1.5))))) {
                var v9 int32 = 1
                if true {
                    var v10 uint32 = 677
                    var v11 int32 = (^(v5))
                    _ = v10
                    _ = v11
                } else {
                    v12 := GenS1{F0: 4294967295}
                    v13 := v6[fi8]
                    _ = v12
                    _ = v13
                }
                if false {
                    v14 := GenS2{F0: false}
                    v6[(-(fi7))] = 4294967295
                    v15 := v6[v2]
                    _ = v14
                    _ = v15
                } else {
                    delete(v4, fi8)
                    var v16 uint8 = 42
                    switch "\n\t" {
                    case "\n\t":
                        v2 += (-(fi8))
                    default:
                        v5 -= (v9 << ((0) & 63))
                    }
                    _ = v16
                }
                _ = v9
            } else {
                if (true && ((!(true && ((^(fi7)) == (fi8 & v2)))) || (32767 != 1))) {
                    if ((-2.25 >= -2.25) || (!(!(42 <= 2)))) {
                        if (42 > 115) {
                            var v17 uint64 = 787
                            switch "" {
                            case "foo bar":
                                var v18 uint = 42
                                _ = v18
                            case "hello":
                                v2 += v2
                                v19 := GenS2{F0: false}
                                v3 = append(v3, (fi7 ^ (fi8 + fi7)))
                                if (9223372036854775807 != (fi8 * 545)) {
                                    var v20 bool = false
                                    var v21 bool = true
                                    if (-491 != 93) {
                                        var v22 int = -720
                                        _ = v22
                                    } else {
                                        switch v2 {
                                        case -1:
                                            v23 := GenS3{F0: -837, F1: (!v21), F2: (-(v2))}
                                            _ = v23
                                        case -9223372036854775808:
                                            var v24 int8 = int8(v2)
                                            _ = v24
                                        case 470:
                                            var v25 int8 = 100
                                            _ = v25
                                        }
                                    }
                                    _ = v20
                                    _ = v21
                                } else {
                                    v4[fi8] = (!(!(2 <= v17)))
                                }
                                switch fi8 {
                                case -9223372036854775808:
                                    if false {
                                        v17 -= 1
                                    } else {
                                        if len(v3) > 0 {
                                            v3[0] = (fi8 * (-(v2)))
                                        }
                                        v17 = v17
                                    }
                                case -909:
                                    v26 := GenS3{F0: -994, F1: (32767 < 0), F2: (^(v2))}
                                    _ = v26
                                }
                                _ = v19
                            }
                            if len(v3) > 0 {
                                v3[0] = fi7
                            }
                            var v27 uint = 4294967295
                            v28 := GenS4{F0: 934}
                            v29 := GenS1{F0: 145}
                            switch (("hello" + ("foo bar" + "a")) + "foo bar") {
                            case "foo bar":
                                if (false || (true && true)) {
                                    v30 := GenS4{F0: 2}
                                    v3 = append(v3, fi8)
                                    _ = v30
                                } else {
                                    v31 := v6[(v2 | (fi8 * (fi8 | (fi8 & -9223372036854775808))))]
                                    v17 += (v17 - v17)
                                    var v32 int = (v2 >> ((v27) & 63))
                                    if (v17 <= v17) {
                                        v17 += v17
                                    } else {
                                        var v33 uint16 = 42
                                        v1 = v1
                                        _ = len(v6)
                                        var v34 string = "foo bar"
                                        _ = v33
                                        _ = v34
                                    }
                                    if false {
                                        _ = len(v6)
                                        switch "" {
                                        case "\n\t":
                                            v35 := v4[(v32 % ((int(v27)) | 1))]
                                            _ = v35
                                        case "a":
                                            if (!true) {
                                                var v36 uint = (v27 + (^(v27)))
                                                if true {
                                                    _ = len(v4)
                                                    switch -663 {
                                                    case 4294967295:
                                                        v37 := GenS5{F0: v5, F1: -42, F2: 2}
                                                        _ = v37
                                                    case -860:
                                                        v38 := GenS6{F0: "hello", F1: 237, F2: (v27 + (v36 & v36))}
                                                        _ = v38
                                                    case 32767:
                                                        v32 = v32
                                                    default:
                                                        if (1 >= 1) {
                                                            v39 := GenS3{F0: 0, F1: false, F2: (v31 & v32)}
                                                            _ = v39
                                                        } else {
                                                            var v40 int8 = 42
                                                            _ = v40
                                                        }
                                                        v41 := GenS5{F0: (-(v5)), F1: -1, F2: 0}
                                                        _ = v41
                                                    }
                                                }
                                                var v42 bool = (42 == 428)
                                                _ = v36
                                                _ = v42
                                            }
                                        case "":
                                            var v43 uint8 = 2
                                            _ = v43
                                        }
                                    } else {
                                        var v44 int16 = 32767
                                        v44 = v44
                                        var v45 int32 = -2147483648
                                        _ = v44
                                        _ = v45
                                    }
                                    _ = v31
                                    _ = v32
                                }
                            case "a":
                                var v46 int64 = 1
                                _ = v46
                            case "\n\t":
                                fmt.Println(fi7, v2)
                                var v47 bool = (true || (!(false || (-2.25 > (float32(1.5) + float32(v1))))))
                                _ = v47
                            default:
                                var v48 uint64 = (v17 ^ (^(v17)))
                                _ = v48
                            }
                            v49 := GenS7{F0: int32(v27), F1: "a", F2: -1}
                            v27 -= (v27 & (v27 % (((v27 - v27)) | 1)))
                            switch int(v17) {
                            case 32767:
                                v50 := GenS6{F0: "hello", F1: (v17 | (v17 ^ (^(v17)))), F2: (v27 >> (((v27 >> ((v27) & 63))) & 63))}
                                _ = v50
                            case 1:
                                var v51 string = ((("a" + ("a" + "hello")) + "") + "a")
                                _ = v51
                            }
                            _ = v17
                            _ = v27
                            _ = v28
                            _ = v29
                            _ = v49
                        }
                        fmt.Println(fi7, v2, v2)
                        v52 := GenS8{F0: false, F1: 0, F2: ""}
                        _ = v52
                    } else {
                        var v53 int32 = (v5 % ((1) | 1))
                        v5 += v5
                        if (v2 == 443) {
                            if true {
                                if true {
                                    if ((float32(v1) * 3.14159) != float32(v1)) {
                                        fmt.Println(v1, v1)
                                        delete(v4, 1)
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        v1 = float64(float32(((v1 + ((v1 / (v1 / 1.0)) - v1)) - 3.14159)))
                                    } else {
                                        v5 = (^(v53))
                                    }
                                }
                                var v54 uint = 414
                                v3 = append(v3, (fi7 + (^(v2))))
                                _ = v54
                            }
                            var v55 int8 = 1
                            var v56 string = (("foo bar" + (("\n\t" + "foo bar") + ("" + (("" + "") + "a")))) + ("foo bar" + "\n\t"))
                            _ = v55
                            _ = v56
                        }
                        _ = v53
                    }
                    v57 := GenS1{F0: 368}
                    _ = v57
                } else {
                    if false {
                        var v58 int8 = -1
                        if (v5 != v5) {
                            switch (fi7 / ((int(v58)) | 1)) {
                            case -2147483648:
                                v2 = (fi7 & 9223372036854775807)
                            }
                        } else {
                            switch "hello" {
                            case "a":
                                if true {
                                    if (0 <= 65535) {
                                        if true {
                                            v2 -= (-(fi7))
                                            fmt.Println(fi7)
                                            fmt.Println(v5)
                                        } else {
                                            var v59 uint32 = 816
                                            _ = v59
                                        }
                                        v60 := GenS6{F0: "hello", F1: 1, F2: 2}
                                        _ = v60
                                    } else {
                                        fmt.Println(fi8)
                                        v5 = v5
                                    }
                                    fmt.Println(v58, fi7)
                                    v61 := v6[v2]
                                    _ = v61
                                } else {
                                    var v62 uint32 = 2
                                    _ = v62
                                }
                            case "\n\t":
                                fmt.Println(fi8, v5, v58)
                            }
                        }
                        var v63 int8 = 1
                        _ = v58
                        _ = v63
                    } else {
                        if (1 != 789) {
                            v64 := GenS1{F0: 42}
                            if false {
                                delete(v6, v2)
                                v65 := GenS6{F0: "a", F1: 4294967295, F2: 51}
                                var v66 int32 = 0
                                _ = v65
                                _ = v66
                            } else {
                                _ = len(v3)
                            }
                            v2 = int(v5)
                            _ = v64
                        } else {
                            if (3.14159 <= 1e10) {
                                var v67 int32 = 411
                                v68 := GenS1{F0: 2}
                                _ = v67
                                _ = v68
                            } else {
                                var v69 int8 = -128
                                if false {
                                    v5 += (v5 << ((854) & 63))
                                    switch (fi7 | (fi8 % (((fi8 * (v2 - fi8))) | 1))) {
                                    case 127:
                                        v3 = append(v3, (v2 ^ (^(fi8))))
                                    case -9223372036854775808:
                                        v70 := v6[fi8]
                                        _ = v70
                                    }
                                } else {
                                    v5 -= -2147483648
                                }
                                v71 := GenS9{F0: 784}
                                _ = v69
                                _ = v71
                            }
                        }
                    }
                    v2 -= 2147483647
                }
                if len(v3) > 0 {
                    v3[0] = -487
                }
            }
            fmt.Println(fi8, fi8)
            _ = fi8
        }
        v72 := GenS1{F0: 18446744073709551615}
        _ = fi7
        _ = v72
    }
    if (0 >= 32767) {
        switch "foo bar" {
        case "foo bar":
            v73 := GenS10{F0: 804, F1: (v2 == (v2 & v2))}
            _ = v73
        case "":
            delete(v6, v2)
        }
    } else {
        if len(v3) > 0 {
            v3[0] = (v2 + v2)
        }
        fmt.Println(v5, v5)
        var v74 string = "a"
        _ = v74
    }
    if true {
        var v75 bool = true
        _ = v75
    }
    if len(v3) > 0 {
        v3[0] = (v2 - v2)
    }
    _ = len(v4)
    v5 = (v5 | -686)
    v76 := GenS5{F0: 42, F1: 42, F2: 42}
    if (42 > 0) {
        v1 = (0.1 + ((float64(-2.25) * (1.5 / 1.5)) / v1))
        v2 += (-(v2))
    }
    if false {
        var v77 uint16 = 255
        var v78 int16 = 32767
        v77 = 65535
        _ = v77
        _ = v78
    }
    _ = len(v3)
    if (true && ((1e10 <= 1e10) && (668 < 880))) {
        v2 = v2
    } else {
        var v79 int32 = v5
        _ = v79
    }
    for fi80 := 0; fi80 < 6; fi80++ {
        var v81 int8 = 100
        v1 -= -2.25
        v82 := GenS11{F0: 128, F1: 1, F2: int64(v5)}
        _ = fi80
        _ = v81
        _ = v82
    }
    for fi83 := 0; fi83 < 19; fi83++ {
        var v84 string = ("" + "hello")
        _ = fi83
        _ = v84
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v76
}
