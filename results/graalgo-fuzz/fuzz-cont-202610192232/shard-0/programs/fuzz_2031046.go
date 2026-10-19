package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 int; F2 uint16 }
type GenS3 struct { F0 int32; F1 string; F2 bool }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int64; F1 string; F2 int64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int; F1 bool }
func genFn0(p0 int32) uint {
    var v7 bool = true
    fmt.Println(v7)
    fmt.Println(v7, p0)
    switch "foo bar" {
    case "":
        if (255 >= 10) {
            v8 := GenS1{F0: -2}
            _ = v8
        }
    case "foo bar":
        var v9 float32 = float32(((float64(1.5) * 1e10) * 1e10))
        _ = v9
    default:
        var v10 uint8 = 0
        _ = v10
    }
    var v11 bool = v7
    _ = p0
    _ = v11
    _ = v7
    return 18446744073709551615
}
func genFn1(p0 int64, p1 int32) uint64 {
    fmt.Println(p0, p0)
    v12 := []int{}
    p0 = (^(p0))
    if (115 == 100) {
        if false {
            switch ("" + "hello") {
            case "a":
                fmt.Println(p1)
            case "\n\t":
                p0 = p0
            default:
                v13 := GenS2{F0: 42, F1: 2147483647, F2: 688}
                _ = v13
            }
        } else {
            for fi14 := 0; fi14 < 8; fi14++ {
                fmt.Println(fi14, p0)
                if (!(!(12 < 2))) {
                    if false {
                        v12 = append(v12, -786)
                    } else {
                        for fi15 := 0; fi15 < 5; fi15++ {
                            if false {
                                var v16 bool = ((true && false) || true)
                                if false {
                                    if (1.5 < ((1e10 - 0.1) - 3.14159)) {
                                        v16 = v16
                                        v17 := GenS2{F0: fi14, F1: fi15, F2: 733}
                                        p1 = (p1 | 1)
                                        _ = v17
                                    }
                                } else {
                                    p0 -= p0
                                    var v18 uint32 = 2
                                    var v19 bool = true
                                    if true {
                                        p1 -= p1
                                        var v20 uint8 = 0
                                        _ = v20
                                    }
                                    if len(v12) > 0 {
                                        v12[0] = fi15
                                    }
                                    _ = v18
                                    _ = v19
                                }
                                _ = v16
                            } else {
                                p1 -= -2147483648
                                var v21 uint8 = 128
                                switch 2147483647 {
                                case 128:
                                    if true {
                                        v22 := GenS2{F0: fi15, F1: (fi15 | fi14), F2: 884}
                                        _ = v22
                                    } else {
                                        var v23 uint8 = (v21 | (v21 | 0))
                                        fmt.Println(v23, p1, p0)
                                        _ = len(v12)
                                        _ = v23
                                    }
                                case -2147483648:
                                    v24 := GenS3{F0: (p1 / ((p1) | 1)), F1: ("" + "foo bar"), F2: (false || true)}
                                    _ = v24
                                default:
                                    switch fi15 {
                                    case 550:
                                        if (1e10 <= (float32(((((1.5 + 1e10) * (0.1 * 3.14159)) / (-2.25 + 1e10)) * float64((float32(1.5) * 0.0)))) - 3.14159)) {
                                            fmt.Println(fi14, fi15)
                                        } else {
                                            p1 -= p1
                                            var v25 int32 = p1
                                            v26 := GenS4{F0: 42}
                                            v12 = append(v12, fi14)
                                            _ = len(v12)
                                            p1 = v25
                                            v27 := GenS2{F0: int(v21), F1: (fi15 % ((fi14) | 1)), F2: 32767}
                                            if true {
                                                fmt.Println(p1, fi15, v25)
                                                v28 := GenS3{F0: -26, F1: "foo bar", F2: (float32((1.5 / (((3.14159 * 3.14159) + (0.0 + 0.0)) / 1.0))) <= ((0.0 + ((((0.0 / 1.5) / 1.0) / float32(1.5)) - ((1e10 - float32(1.5)) + ((1e10 / -2.25) + float32(0.0))))) / 3.14159))}
                                                _ = v28
                                            }
                                            _ = v25
                                            _ = v26
                                            _ = v27
                                        }
                                    case 127:
                                        var v29 uint64 = 740
                                        _ = v29
                                    }
                                }
                                _ = v21
                            }
                            if len(v12) > 0 {
                                v12[0] = (-(fi14))
                            }
                            _ = fi15
                        }
                        p1 += p1
                    }
                    v12 = append(v12, fi14)
                    for fi30 := 0; fi30 < 6; fi30++ {
                        p1 = (p1 >> ((255) & 63))
                        var v31 float64 = -2.25
                        _ = fi30
                        _ = v31
                    }
                }
                _ = fi14
            }
            p0 = (^(p0))
            for fi32 := 0; fi32 < 19; fi32++ {
                for fi33 := 0; fi33 < 12; fi33++ {
                    v34 := GenS3{F0: -397, F1: "hello", F2: (false && true)}
                    _ = fi33
                    _ = v34
                }
                _ = fi32
            }
        }
        p0 = p0
        p0 = (-(p0))
    }
    _ = p0
    _ = p1
    _ = v12
    return 42
}

func main() {
    var v1 uint64 = 2
    var v2 bool = false
    v3 := []string{"", "foo bar"}
    var v4 float64 = 1.5
    var v5 int8 = -109
    var v6 float32 = 3.14159
    var v35 uint64 = (^(v1))
    var v36 uint64 = 4294967295
    switch -2147483648 {
    case -1:
        fmt.Println(v4, v35)
    case 32767:
        v6 = (((v6 + 1e10) / 0.1) - (0.0 / v6))
    case -2147483648:
        v37 := GenS4{F0: 2}
        _ = v37
    default:
        var v38 int8 = v5
        _ = v38
    }
    v4 = (((v4 / float64(3.14159)) - 1.5) * v4)
    if (v1 != genFn1(42, -1)) {
        var v39 int8 = 100
        var v40 int8 = v5
        _ = v39
        _ = v40
    }
    if (128 < 0) {
        _ = len(v3)
        v41 := GenS5{F0: 42, F1: "\n\t", F2: 9223372036854775807}
        _ = v41
    } else {
        switch 2147483647 {
        case 376:
            var v42 uint8 = 1
            _ = v42
        case 32768:
            v43 := GenS5{F0: 9223372036854775807, F1: "", F2: -1}
            _ = v43
        case 567:
            v2 = v2
        default:
            var v44 float64 = 1.5
            _ = v44
        }
    }
    for fi45 := 0; fi45 < 5; fi45++ {
        v46 := GenS6{F0: 127}
        _ = fi45
        _ = v46
    }
    var v47 float32 = (3.14159 + 0.1)
    var v48 uint64 = v35
    var v49 int8 = 104
    if v2 {
        if ((0 != 1) || ("foo bar" == "hello")) {
            fmt.Println(v47)
            switch 402 {
            case 32767:
                if (v2 && true) {
                    for fi50 := 0; fi50 < 19; fi50++ {
                        switch "a" {
                        case "":
                            fmt.Println(v2, v35, v47)
                        case "hello":
                            for fi51 := 0; fi51 < 13; fi51++ {
                                fmt.Println(v2, v1, v5)
                                if ((^(v48)) > (^(v48))) {
                                    fmt.Println(v47, v35, v1)
                                } else {
                                    var v52 int = -9223372036854775808
                                    var v53 uint = 2
                                    var v54 uint32 = 42
                                    _ = v52
                                    _ = v53
                                    _ = v54
                                }
                                var v55 uint8 = 128
                                _ = fi51
                                _ = v55
                            }
                        default:
                            fmt.Println(v5, v47, v35)
                        }
                        fmt.Println(v1, v1, v5)
                        v3 = append(v3, ("foo bar" + "hello"))
                        _ = fi50
                    }
                }
            case 42:
                switch -9223372036854775808 {
                case -2147483648:
                    if (v2 || v2) {
                        fmt.Println(v6, v49)
                        var v56 uint = 0
                        var v57 int64 = -9223372036854775808
                        _ = v56
                        _ = v57
                    } else {
                        if v2 {
                            _ = len(v3)
                        } else {
                            if v2 {
                                switch 9223372036854775807 {
                                case 32767:
                                    v1 -= v48
                                case -2147483648:
                                    v58 := GenS7{F0: 2}
                                    _ = v58
                                }
                            }
                            v47 += 0.0
                            var v59 uint16 = 65535
                            _ = v59
                        }
                        for fi60 := 0; fi60 < 15; fi60++ {
                            v3 = append(v3, "a")
                            _ = fi60
                        }
                        if false {
                            for fi61 := 0; fi61 < 13; fi61++ {
                                var v62 string = ("hello" + "\n\t")
                                _ = fi61
                                _ = v62
                            }
                            fmt.Println(v2)
                            if (!((65535 < 1) && (!v2))) {
                                v63 := GenS7{F0: 128}
                                v36 += (v48 * 0)
                                switch ("foo bar" + ("\n\t" + ("" + "foo bar"))) {
                                case "":
                                    if ((((!v2) && (genFn0(431) <= 65535)) || true) || v2) {
                                        for fi64 := 0; fi64 < 11; fi64++ {
                                            v6 = float32(-2.25)
                                            _ = fi64
                                        }
                                        v2 = true
                                    } else {
                                        var v65 uint16 = 65535
                                        var v66 int64 = 42
                                        _ = v65
                                        _ = v66
                                    }
                                case "a":
                                    var v67 float64 = float64((3.14159 - v6))
                                    _ = v67
                                default:
                                    fmt.Println(v6, v48)
                                }
                                _ = v63
                            }
                        } else {
                            switch "hello" {
                            case "":
                                if (!(-9223372036854775808 >= 1)) {
                                    v3 = append(v3, "")
                                } else {
                                    v3 = append(v3, ("\n\t" + ""))
                                }
                            }
                        }
                    }
                default:
                    v68 := GenS8{F0: 128, F1: (-9223372036854775808 != 128)}
                    _ = v68
                }
            case 608:
                fmt.Println(v1, v2, v1)
            }
        } else {
            if len(v3) > 0 {
                v3[0] = ("foo bar" + "")
            }
            var v69 int = 128
            if len(v3) > 0 {
                v3[0] = "hello"
            }
            _ = v69
        }
        for fi70 := 0; fi70 < 11; fi70++ {
            var v71 int64 = -9223372036854775808
            for fi72 := 0; fi72 < 8; fi72++ {
                v73 := GenS8{F0: fi70, F1: v2}
                v1 -= (v48 | genFn1(1, 2147483647))
                v6 += -2.25
                _ = fi72
                _ = v73
            }
            v35 = (v48 * v36)
            _ = fi70
            _ = v71
        }
    } else {
        var v74 int = 909
        for fi75 := 0; fi75 < 14; fi75++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi75
        }
        _ = v74
    }
    var v76 uint32 = 523
    if v2 {
        fmt.Println(v6)
        switch -1 {
        case 9223372036854775807:
            v77 := GenS7{F0: 42}
            _ = v77
        case 0:
            v47 = v6
        case -2147483648:
            for fi78 := 0; fi78 < 17; fi78++ {
                v48 = v35
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v79 uint32 = v76
                _ = fi78
                _ = v79
            }
        default:
            for fi80 := 0; fi80 < 20; fi80++ {
                v4 += v4
                fmt.Println(v48)
                _ = fi80
            }
        }
        var v81 int16 = 0
        _ = v81
    } else {
        v4 = v4
        v5 += v5
    }
    var v82 uint8 = 42
    if v2 {
        fmt.Println(v82)
        fmt.Println(v35, v35)
    } else {
        v35 += (v48 + 671)
        fmt.Println(v49)
        var v83 uint64 = v36
        _ = v83
    }
    fmt.Println(v2, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v76:", v76)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v6
    _ = v76
    _ = v82
}
