package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 int16 }
type GenS2 struct { F0 int32; F1 int8; F2 int }
type GenS3 struct { F0 uint64; F1 uint }
type GenS4 struct { F0 uint16; F1 uint32; F2 int32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int16 }
func genFn0() int64 {
    v7 := GenS1{F0: -1, F1: 938, F2: 384}
    v8 := GenS1{F0: -32768, F1: 127, F2: 1000}
    var v9 uint64 = 42
    var v10 int16 = -995
    var v11 uint = 7
    _ = v10
    _ = v11
    _ = v7
    _ = v8
    _ = v9
    return -621
}

func main() {
    var v1 uint8 = 127
    var v2 int = 635
    v3 := []uint32{0, 1, 1, 579}
    var v4 string = ""
    v5 := make(map[string]int)
    v6 := []int{42}
    v4 = ((((v4 + v4) + "foo bar") + "\n\t") + v4)
    var v12 string = (v4 + "a")
    var v13 uint32 = 478
    v6 = append(v6, (v2 << ((849) & 63)))
    if (595 >= 156) {
        v4 = (v12 + "hello")
        v2 -= 810
        for fi14 := 0; fi14 < 5; fi14++ {
            if false {
                var v15 bool = ((!true) || false)
                delete(v5, "hello")
                _ = v15
            }
            _ = fi14
        }
    }
    v2 -= v2
    if (!true) {
        if true {
            var v16 bool = false
            _ = v16
        } else {
            if false {
                if (4294967295 > 18446744073709551615) {
                    v12 = "foo bar"
                }
            } else {
                if false {
                    for fi17 := 0; fi17 < 16; fi17++ {
                        fmt.Println(fi17, v12, v4)
                        for fi18 := 0; fi18 < 18; fi18++ {
                            var v19 int32 = 42
                            switch 32767 {
                            case -529:
                                _ = len(v3)
                            case -9223372036854775808:
                                var v20 uint16 = 36
                                _ = v20
                            case 0:
                                v21 := GenS2{F0: v19, F1: -1, F2: v2}
                                _ = v21
                            default:
                                if true {
                                    v22 := GenS3{F0: 797, F1: uint(v13)}
                                    v2 = 127
                                    _ = v22
                                } else {
                                    delete(v5, v12)
                                    var v23 bool = (-9223372036854775808 > 0)
                                    _ = v23
                                }
                            }
                            switch (v2 % ((v2) | 1)) {
                            case -841:
                                var v24 bool = (!true)
                                _ = v24
                            case -1:
                                if (!false) {
                                    v25 := v5[(v4 + v4)]
                                    _ = v25
                                } else {
                                    v2 += fi18
                                    v26 := GenS1{F0: 1000, F1: (v2 * fi17), F2: -449}
                                    fmt.Println(v1)
                                    _ = v26
                                }
                            }
                            _ = fi18
                            _ = v19
                        }
                        _ = fi17
                    }
                    switch v12 {
                    case "hello":
                        var v27 uint = 65535
                        _ = v27
                    case "a":
                        v12 = v4
                    case "\n\t":
                        var v28 uint = 18446744073709551615
                        _ = v28
                    default:
                        v29 := GenS3{F0: 42, F1: 564}
                        _ = v29
                    }
                }
            }
        }
        v5["a"] = v2
        fmt.Println(v4, v12)
    }
    switch (("foo bar" + (v4 + ((v4 + v4) + v4))) + "a") {
    case "hello":
        v30 := GenS4{F0: 255, F1: 2147483648, F2: -602}
        _ = v30
    }
    for fi31 := 0; fi31 < 16; fi31++ {
        delete(v5, ((("hello" + v4) + v4) + "hello"))
        _ = len(v3)
        _ = fi31
    }
    if true {
        v32 := GenS3{F0: 727, F1: 42}
        var v33 float64 = ((1.5 + (0.1 - 1e10)) - 0.1)
        if (255 == 766) {
            if false {
                var v34 bool = true
                for fi35 := 0; fi35 < 8; fi35++ {
                    if v34 {
                        switch fi35 {
                        case 42:
                            fmt.Println(v34, v13)
                        default:
                            if ((!((!true) || false)) || true) {
                                for fi36 := 0; fi36 < 9; fi36++ {
                                    var v37 int8 = int8(v2)
                                    var v38 float64 = v33
                                    _ = fi36
                                    _ = v37
                                    _ = v38
                                }
                            } else {
                                var v39 uint32 = 65535
                                switch (^(fi35)) {
                                case -9223372036854775808:
                                    if false {
                                        if (v33 == v33) {
                                            var v40 bool = false
                                            _ = v40
                                        }
                                        v41 := GenS2{F0: 42, F1: -1, F2: (fi35 ^ (^(fi35)))}
                                        v42 := GenS2{F0: 2147483647, F1: -1, F2: fi35}
                                        _ = v41
                                        _ = v42
                                    }
                                default:
                                    delete(v5, ((((("foo bar" + v12) + ("a" + v12)) + (("\n\t" + ((v12 + v4) + "a")) + v4)) + (v12 + "hello")) + v12))
                                }
                                switch v12 {
                                case "a":
                                    var v43 int16 = 0
                                    _ = v43
                                case "foo bar":
                                    v44 := GenS1{F0: 1, F1: (fi35 << ((42) & 63)), F2: -32768}
                                    _ = v44
                                default:
                                    for fi45 := 0; fi45 < 20; fi45++ {
                                        if len(v6) > 0 {
                                            v6[0] = -9223372036854775808
                                        }
                                        _ = fi45
                                    }
                                }
                                _ = v39
                            }
                        }
                        fmt.Println(v1)
                        fmt.Println(v1)
                    } else {
                        for fi46 := 0; fi46 < 17; fi46++ {
                            if false {
                                switch "hello" {
                                case "a":
                                    switch v4 {
                                    case "\n\t":
                                        var v47 uint8 = uint8(v13)
                                        _ = v47
                                    case "":
                                        v13 = v13
                                    case "a":
                                        v33 = (v33 - v33)
                                    default:
                                        var v48 float64 = v33
                                        v49 := GenS4{F0: 32767, F1: v13, F2: 1}
                                        var v50 int8 = -128
                                        _ = v48
                                        _ = v49
                                        _ = v50
                                    }
                                case "":
                                    switch (fi46 * (fi46 + fi35)) {
                                    case 127:
                                        var v51 int64 = int64(v1)
                                        _ = v51
                                    }
                                case "foo bar":
                                    if ((!true) || (!(-128 != 127))) {
                                        switch ("hello" + "") {
                                        case "":
                                            var v52 bool = true
                                            _ = v52
                                        }
                                    }
                                default:
                                    var v53 int32 = 1
                                    _ = v53
                                }
                                switch v4 {
                                case "":
                                    var v54 float32 = 1.5
                                    _ = v54
                                case "\n\t":
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                default:
                                    v3 = append(v3, v13)
                                }
                            } else {
                                fmt.Println(v2, v1, v4)
                                if len(v6) > 0 {
                                    v6[0] = 0
                                }
                            }
                            _ = fi46
                        }
                        if false {
                            switch 32767 {
                            case 128:
                                v4 = ("foo bar" + (v4 + v12))
                            case -1:
                                for fi55 := 0; fi55 < 14; fi55++ {
                                    var v56 uint = 18446744073709551615
                                    _ = fi55
                                    _ = v56
                                }
                            }
                        } else {
                            v3 = append(v3, (^(v13)))
                        }
                        for fi57 := 0; fi57 < 16; fi57++ {
                            v58 := GenS4{F0: 1, F1: (v13 - 2147483648), F2: -917}
                            _ = fi57
                            _ = v58
                        }
                    }
                    if (((!(!false)) || (255 == 65535)) && v34) {
                        var v59 float32 = -2.25
                        v6 = append(v6, (^(v2)))
                        if (v34 || true) {
                            fmt.Println(v12)
                            switch v2 {
                            case 941:
                                v60 := GenS1{F0: -517, F1: v2, F2: 257}
                                _ = v60
                            case 1:
                                if len(v3) > 0 {
                                    v3[0] = 0
                                }
                            case 9223372036854775807:
                                fmt.Println(v2, fi35, v12)
                            default:
                                for fi61 := 0; fi61 < 8; fi61++ {
                                    fmt.Println(v2, fi61)
                                    _ = len(v5)
                                    _ = fi61
                                }
                            }
                            if true {
                                _ = len(v6)
                            } else {
                                var v62 uint8 = (^(v1))
                                if (v33 == v33) {
                                    for fi63 := 0; fi63 < 16; fi63++ {
                                        var v64 bool = false
                                        _ = fi63
                                        _ = v64
                                    }
                                    if v34 {
                                        var v65 int8 = 100
                                        _ = v65
                                    }
                                }
                                v3 = append(v3, v13)
                                _ = v62
                            }
                        }
                        _ = v59
                    }
                    _ = fi35
                }
                _ = v34
            } else {
                for fi66 := 0; fi66 < 11; fi66++ {
                    v67 := GenS5{F0: (v2 * fi66)}
                    var v68 float64 = ((v33 - v33) - v33)
                    v69 := v5[v12]
                    _ = fi66
                    _ = v67
                    _ = v68
                    _ = v69
                }
                for fi70 := 0; fi70 < 12; fi70++ {
                    v1 -= v1
                    v71 := GenS2{F0: 2147483647, F1: 0, F2: v2}
                    _ = fi70
                    _ = v71
                }
            }
            for fi72 := 0; fi72 < 16; fi72++ {
                var v73 bool = (-2.25 < v33)
                if len(v3) > 0 {
                    v3[0] = 65535
                }
                if v73 {
                    v73 = ((!false) || (0 > 704))
                } else {
                    var v74 float64 = v33
                    fmt.Println(v74, v33)
                    _ = v74
                }
                _ = fi72
                _ = v73
            }
        }
        _ = v32
        _ = v33
    } else {
        v75 := GenS2{F0: 0, F1: -128, F2: v2}
        _ = v75
    }
    for fi76 := 0; fi76 < 16; fi76++ {
        var v77 int = (fi76 % (((fi76 >> ((255) & 63))) | 1))
        _ = fi76
        _ = v77
    }
    v78 := GenS3{F0: 100, F1: 1}
    v79 := GenS6{F0: 1}
    v80 := GenS1{F0: 1, F1: 774, F2: -69}
    v13 = v13
    var v81 float32 = (1.5 * (3.14159 / 3.14159))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v78
    _ = v79
    _ = v80
    _ = v81
}
