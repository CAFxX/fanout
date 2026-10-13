package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 int8; F2 uint }
type GenS3 struct { F0 int64; F1 int; F2 int64 }
type GenS4 struct { F0 int16; F1 uint8 }
type GenS5 struct { F0 uint32; F1 uint32; F2 int8 }
type GenS6 struct { F0 string; F1 int; F2 string }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int32; F1 bool; F2 int8 }
type GenS9 struct { F0 int64 }
func genFn0(p0 int, p1 int) uint32 {
    switch 9223372036854775807 {
    case -9223372036854775808:
        p0 += 863
    }
    var v5 uint8 = 71
    var v6 int16 = 42
    p1 = (p0 ^ 32768)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 154
}
func genFn1(p0 uint8, p1 int32, p2 uint64) uint64 {
    for fi7 := 0; fi7 < 15; fi7++ {
        if false {
            if (!false) {
                if false {
                    fmt.Println(p1, fi7)
                    if (1 <= -1) {
                        for fi8 := 0; fi8 < 11; fi8++ {
                            v9 := []int{0, fi7, fi8, (fi7 ^ 127)}
                            if len(v9) > 0 {
                                v9[0] = 2147483647
                            }
                            _ = fi8
                            _ = v9
                        }
                        var v10 uint16 = 2
                        var v11 uint64 = p2
                        _ = v10
                        _ = v11
                    }
                }
                for fi12 := 0; fi12 < 10; fi12++ {
                    if (!true) {
                        p0 = 245
                    }
                    if false {
                        p1 += 2147483647
                        if false {
                            if (0.1 >= float32(3.14159)) {
                                var v13 float32 = 1e10
                                p2 = (p2 + p2)
                                _ = v13
                            } else {
                                fmt.Println(fi12, fi7)
                                v14 := GenS1{F0: -173}
                                if ((-1 <= 0) && false) {
                                    v15 := GenS2{F0: 438, F1: 0, F2: 1}
                                    v16 := []int{}
                                    v17 := GenS3{F0: 0, F1: (fi7 >> ((255) & 63)), F2: -9223372036854775808}
                                    _ = v15
                                    _ = v16
                                    _ = v17
                                } else {
                                    v18 := GenS4{F0: 0, F1: uint8(p2)}
                                    _ = v18
                                }
                                v19 := GenS4{F0: 32767, F1: uint8(p2)}
                                v20 := []int{(fi12 / (((fi7 << ((4294967295) & 63))) | 1)), fi7, 199, fi7}
                                var v21 int = (^(fi12))
                                _ = v14
                                _ = v19
                                _ = v20
                                _ = v21
                            }
                            p0 = uint8(p2)
                            var v22 float64 = 0.0
                            _ = v22
                        }
                        p0 += 0
                        var v23 int = (fi12 << ((255) & 63))
                        _ = v23
                    } else {
                        var v24 uint = 0
                        var v25 float32 = (((3.14159 / (0.1 * 1e10)) * (0.0 + -2.25)) - -2.25)
                        var v26 uint64 = p2
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                    v27 := GenS2{F0: 479, F1: 1, F2: 0}
                    _ = fi12
                    _ = v27
                }
                var v28 uint16 = 2
                _ = v28
            } else {
                v29 := GenS1{F0: 0}
                fmt.Println(p1, p0)
                _ = v29
            }
        } else {
            var v30 string = (((("foo bar" + "\n\t") + "hello") + (("foo bar" + ((("" + "") + ("hello" + "")) + "foo bar")) + "foo bar")) + (("" + "foo bar") + "hello"))
            _ = v30
        }
        v31 := GenS2{F0: 833, F1: 0, F2: 73}
        _ = fi7
        _ = v31
    }
    var v32 string = "a"
    _ = p0
    _ = p1
    _ = p2
    _ = v32
    return p2
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = "hello"
    v3 := make(map[string]int)
    var v4 float32 = (0.1 + 3.14159)
    var v33 int8 = 127
    v34 := []int{}
    v33 -= -1
    var v35 int8 = v33
    var v36 uint32 = 0
    v2 = (v2 + v2)
    var v37 uint8 = 1
    _ = len(v3)
    _ = len(v34)
    switch v2 {
    case "hello":
        for fi38 := 0; fi38 < 20; fi38++ {
            if false {
                var v39 uint32 = (v36 | v36)
                _ = v39
            } else {
                var v40 uint = 2
                for fi41 := 0; fi41 < 15; fi41++ {
                    var v42 int64 = 1
                    v3[((v2 + (("\n\t" + ((v2 + (v2 + "")) + (v2 + "a"))) + ("a" + (v2 + ("hello" + (((((((v2 + (v2 + "a")) + v2) + "hello") + v2) + v2) + (v2 + (v2 + "a"))) + ("hello" + ("hello" + v2)))))))) + (v2 + (("a" + v2) + ((v2 + v2) + v2))))] = fi41
                    var v43 uint32 = (v36 >> ((v40) & 63))
                    _ = fi41
                    _ = v42
                    _ = v43
                }
                _ = v40
            }
            var v44 int8 = (v33 * v35)
            if false {
                fmt.Println(v2, v2, v33)
                v3[((v2 + (v2 + ("\n\t" + "foo bar"))) + "foo bar")] = (fi38 + (fi38 + fi38))
            } else {
                switch v2 {
                case "\n\t":
                    v45 := GenS5{F0: 65535, F1: 929, F2: (v33 / (((v44 + -128)) | 1))}
                    _ = v45
                case "a":
                    switch ("\n\t" + v2) {
                    case "foo bar":
                        var v46 int16 = 0
                        _ = v46
                    case "\n\t":
                        v4 += 1.5
                    default:
                        v47 := GenS4{F0: 1, F1: v37}
                        _ = v47
                    }
                }
                if ((255 > 1) && (false || (-372 < 1))) {
                    switch v2 {
                    case "foo bar":
                        var v48 uint64 = 0
                        _ = v48
                    case "a":
                        if (4294967295 > 921) {
                            v3[""] = fi38
                            for fi49 := 0; fi49 < 12; fi49++ {
                                var v50 uint32 = 576
                                if false {
                                    v4 = 0.1
                                } else {
                                    switch (^(fi38)) {
                                    case 0:
                                        v51 := GenS6{F0: "", F1: (fi49 ^ (fi49 - fi49)), F2: v2}
                                        _ = v51
                                    case -9223372036854775808:
                                        v1[v36] = v50
                                    case 127:
                                        v52 := GenS2{F0: 42, F1: 100, F2: 4294967295}
                                        _ = v52
                                    }
                                    var v53 float64 = 1e10
                                    delete(v1, v50)
                                    _ = v53
                                }
                                delete(v3, v2)
                                _ = fi49
                                _ = v50
                            }
                            switch (v2 + "") {
                            case "a":
                                var v54 int16 = 832
                                _ = v54
                            case "\n\t":
                                fmt.Println(v44, v37)
                            default:
                                if (!false) {
                                    var v55 uint8 = (^(v37))
                                    v56 := GenS7{F0: 948}
                                    _ = v55
                                    _ = v56
                                } else {
                                    var v57 int16 = 1
                                    for fi58 := 0; fi58 < 17; fi58++ {
                                        v3[("hello" + ((v2 + ((v2 + ((v2 + v2) + (v2 + "foo bar"))) + v2)) + (v2 + (v2 + v2))))] = int(v36)
                                        if true {
                                            var v59 int = (fi38 / ((fi58) | 1))
                                            var v60 uint64 = 4294967295
                                            var v61 uint64 = (v60 / ((genFn1(127, -2147483648, (v60 ^ genFn1(v37, 2147483647, v60)))) | 1))
                                            _ = v59
                                            _ = v60
                                            _ = v61
                                        }
                                        v34 = append(v34, 42)
                                        switch v2 {
                                        case "foo bar":
                                            v62 := v3[v2]
                                            _ = v62
                                        case "":
                                            v4 = (v4 + v4)
                                        case "a":
                                            if false {
                                                var v63 string = "hello"
                                                var v64 int16 = v57
                                                _ = v63
                                                _ = v64
                                            } else {
                                                var v65 uint = 270
                                                _ = v65
                                            }
                                            var v66 int64 = 115
                                            _ = v66
                                        }
                                        if ((fi38 * -142) >= fi38) {
                                            v35 -= v35
                                        }
                                        if (!(true || false)) {
                                            switch (fi38 ^ fi38) {
                                            case 4294967295:
                                                var v67 int = -77
                                                _ = v67
                                            case 2147483647:
                                                v36 -= (^(v36))
                                            case 32768:
                                                var v68 uint32 = (v36 >> ((uint(v36)) & 63))
                                                _ = v68
                                            default:
                                                var v69 bool = true
                                                _ = v69
                                            }
                                            if (!(true || (v37 != v37))) {
                                                if len(v34) >= 2 {
                                                    v34 = v34[1:len(v34)]
                                                }
                                                _ = len(v34)
                                            }
                                            fmt.Println(v2, v36, v44)
                                            if true {
                                                var v70 int32 = -982
                                                switch (fi58 | int(v57)) {
                                                case 9223372036854775807:
                                                    var v71 float64 = float64(v4)
                                                    _ = v71
                                                case 0:
                                                    var v72 float64 = -2.25
                                                    _ = v72
                                                case 821:
                                                    var v73 uint16 = 0
                                                    _ = v73
                                                default:
                                                    var v74 uint = 4294967295
                                                    _ = v74
                                                }
                                                _ = v70
                                            }
                                        }
                                        _ = fi58
                                    }
                                    _ = v57
                                }
                            }
                        }
                    case "hello":
                        var v75 uint = 18446744073709551615
                        _ = v75
                    default:
                        switch (v2 + (((v2 + (v2 + v2)) + (v2 + "")) + v2)) {
                        case "foo bar":
                            var v76 float32 = 3.14159
                            _ = v76
                        case "\n\t":
                            if (!(true && false)) {
                                v77 := GenS8{F0: 42, F1: ((0.0 < (0.1 / 1e10)) && true), F2: (v44 >> ((4294967295) & 63))}
                                v33 += -128
                                v4 -= v4
                                _ = v77
                            } else {
                                var v78 uint = 18446744073709551615
                                v79 := GenS3{F0: 9223372036854775807, F1: fi38, F2: -1}
                                _ = v78
                                _ = v79
                            }
                        default:
                            if len(v34) >= 2 {
                                v34 = v34[1:len(v34)]
                            }
                        }
                    }
                } else {
                    var v80 uint64 = 42
                    var v81 float64 = -2.25
                    var v82 uint = 4294967295
                    _ = v80
                    _ = v81
                    _ = v82
                }
                v83 := GenS9{F0: 9223372036854775807}
                _ = v83
            }
            _ = fi38
            _ = v44
        }
    case "":
        for fi84 := 0; fi84 < 15; fi84++ {
            var v85 int16 = 537
            var v86 bool = false
            _ = fi84
            _ = v85
            _ = v86
        }
    }
    var v87 int8 = 0
    v33 += (v33 & v33)
    for fi88 := 0; fi88 < 8; fi88++ {
        for fi89 := 0; fi89 < 3; fi89++ {
            fmt.Println(v87, v33)
            _ = fi89
        }
        v87 = (v33 + v87)
        _ = fi88
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v87
}
