package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 int32; F2 int16 }
type GenS3 struct { F0 string; F1 int16; F2 int64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int; F1 int16 }
func genFn0(p0 int64, p1 int32, p2 uint8) string {
    var v5 uint8 = p2
    v6 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return ("foo bar" + ("foo bar" + "hello"))
}
func genFn1(p0 uint, p1 uint64, p2 uint32) int32 {
    var v7 uint32 = uint32(p1)
    for fi8 := 0; fi8 < 2; fi8++ {
        p1 += p1
        v9 := []int{21, 32768}
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    return -959
}

func main() {
    v1 := []uint32{871, 16}
    var v2 uint16 = 255
    var v3 int64 = -1
    var v4 int = 55
    fmt.Println(v2, v2)
    if (42 == 549) {
        if (255 != 255) {
            if (float64(float32(1.5)) != (0.1 * (1.5 - -2.25))) {
                var v11 int8 = 100
                v4 = (-(v4))
                for fi12 := 0; fi12 < 14; fi12++ {
                    v4 += (fi12 % ((fi12) | 1))
                    var v13 float64 = (1.5 / -2.25)
                    for fi14 := 0; fi14 < 17; fi14++ {
                        if true {
                            var v15 float64 = v13
                            if ("\n\t" > ((("hello" + "\n\t") + ("a" + "hello")) + "foo bar")) {
                                v2 = (v2 % ((v2) | 1))
                            }
                            _ = v15
                        } else {
                            if (0 == -1) {
                                v16 := GenS1{F0: 586}
                                _ = v16
                            } else {
                                fmt.Println(v13, fi14, v2)
                            }
                            v11 += v11
                        }
                        _ = fi14
                    }
                    _ = fi12
                    _ = v13
                }
                _ = v11
            } else {
                fmt.Println(v2, v2, v3)
                switch v4 {
                case -2147483648:
                    v2 += (v2 * 42)
                case 0:
                    var v17 uint64 = 0
                    _ = v17
                case 1:
                    var v18 uint16 = (v2 | 42)
                    _ = v18
                default:
                    if (false || (!true)) {
                        v3 = (v3 ^ 9223372036854775807)
                        v2 -= v2
                    }
                }
            }
            var v19 int16 = 1000
            _ = v19
        }
        var v20 uint8 = 2
        var v21 int = (v4 % ((v4) | 1))
        _ = v20
        _ = v21
    }
    if (42 > 426) {
        var v22 float32 = (-2.25 + ((-2.25 / (1e10 + (-2.25 / 1.5))) / ((((-2.25 * 0.1) * 1e10) - 3.14159) / (-2.25 / (float32(3.14159) / 1e10)))))
        _ = v22
    } else {
        v2 -= v2
        if true {
            for fi23 := 0; fi23 < 19; fi23++ {
                fmt.Println(fi23)
                var v24 string = (((("" + "foo bar") + "hello") + "\n\t") + ("" + ("\n\t" + "\n\t")))
                _ = fi23
                _ = v24
            }
            fmt.Println(v4, v3)
            switch (^(v4)) {
            case 9223372036854775807:
                v3 = (v3 - v3)
            case 4294967295:
                var v25 int32 = -2147483648
                _ = v25
            }
        }
        v4 -= v4
    }
    if len(v1) > 0 {
        v1[0] = 0
    }
    var v26 string = "hello"
    if true {
        v27 := GenS2{F0: v26, F1: -2147483648, F2: -254}
        switch v26 {
        case "a":
            if ((((v26 + ((v26 + ("\n\t" + v26)) + ((v26 + v26) + "a"))) + "") + v26) != ("a" + (v26 + "\n\t"))) {
                v4 -= (v4 | int(v3))
                if false {
                    var v28 uint32 = 42
                    _ = v28
                }
                switch v4 {
                case 2147483647:
                    var v29 int = v4
                    _ = v29
                case 9223372036854775807:
                    v30 := GenS3{F0: "a", F1: 42, F2: v3}
                    _ = v30
                case -1:
                    for fi31 := 0; fi31 < 13; fi31++ {
                        v3 += 9223372036854775807
                        if ((v26 >= "hello") || (3.14159 == (1.5 / 1e10))) {
                            for fi32 := 0; fi32 < 13; fi32++ {
                                if (-128 <= 100) {
                                    var v33 float32 = 1e10
                                    var v34 uint32 = 0
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = v33
                                    _ = v34
                                }
                                if (0 == 1) {
                                    var v35 uint = 42
                                    _ = v35
                                } else {
                                    v36 := GenS3{F0: "a", F1: 1, F2: (v3 & v3)}
                                    if true {
                                        v37 := GenS2{F0: (v26 + v26), F1: 448, F2: -1}
                                        _ = v37
                                    } else {
                                        _ = len(v1)
                                    }
                                    fmt.Println(v4, v3)
                                    _ = v36
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = fi32
                            }
                            var v38 bool = false
                            _ = v38
                        } else {
                            v2 = (v2 >> ((101) & 63))
                            v4 += 989
                            if len(v1) > 0 {
                                v1[0] = 2147483647
                            }
                        }
                        _ = len(v1)
                        _ = fi31
                    }
                default:
                    v39 := GenS2{F0: ((v26 + v26) + v26), F1: -2147483648, F2: -791}
                    _ = v39
                }
            } else {
                v3 -= (v3 << ((uint(v2)) & 63))
                if true {
                    var v40 int64 = 0
                    v4 += (-(v4))
                    _ = v40
                } else {
                    var v41 uint16 = v2
                    v26 = "a"
                    if (((3.14159 / 3.14159) - 1.5) <= 3.14159) {
                        for fi42 := 0; fi42 < 5; fi42++ {
                            v1 = append(v1, 1)
                            _ = fi42
                        }
                    }
                    _ = v41
                }
                fmt.Println(v2)
            }
        case "":
            v2 = (v2 / (((^(v2))) | 1))
        case "\n\t":
            v1 = append(v1, uint32(v2))
        }
        _ = len(v1)
        _ = v27
    } else {
        v1 = append(v1, 392)
        v43 := GenS3{F0: "", F1: -32768, F2: (v3 ^ (v3 ^ v3))}
        _ = len(v1)
        _ = v43
    }
    var v44 uint64 = 693
    v45 := GenS2{F0: v26, F1: -336, F2: 387}
    var v46 bool = (v26 <= v26)
    var v47 uint16 = (^(v2))
    var v48 float32 = 1.5
    var v49 int16 = 32767
    v3 -= v3
    if (0 != int32(v3)) {
        for fi50 := 0; fi50 < 5; fi50++ {
            for fi51 := 0; fi51 < 3; fi51++ {
                var v52 uint = 137
                v53 := GenS1{F0: uint64(v2)}
                v46 = v46
                _ = fi51
                _ = v52
                _ = v53
            }
            fmt.Println(v47, v44)
            _ = fi50
        }
        fmt.Println(v26, v26, v47)
        for fi54 := 0; fi54 < 4; fi54++ {
            v49 = (v49 >> ((868) & 63))
            _ = fi54
        }
    } else {
        for fi55 := 0; fi55 < 2; fi55++ {
            fmt.Println(v4)
            _ = fi55
        }
    }
    for fi56 := 0; fi56 < 7; fi56++ {
        v26 = (v26 + (v26 + (v26 + (v26 + (v26 + ("\n\t" + (v26 + ("\n\t" + "\n\t"))))))))
        fmt.Println(v3, v47, v3)
        _ = fi56
    }
    v1 = append(v1, 65535)
    for fi57 := 0; fi57 < 13; fi57++ {
        switch (fi57 - v4) {
        case 805:
            v58 := GenS1{F0: v44}
            _ = v58
        case 0:
            var v59 int64 = v3
            _ = v59
        case 32767:
            if v46 {
                if v46 {
                    v60 := GenS4{F0: 42}
                    var v61 int = fi57
                    fmt.Println(v46, v26, v47)
                    _ = v60
                    _ = v61
                } else {
                    var v62 int32 = 0
                    fmt.Println(v4)
                    _ = v62
                }
                switch v26 {
                case "\n\t":
                    switch (v4 ^ v4) {
                    case -79:
                        switch fi57 {
                        case -9223372036854775808:
                            for fi63 := 0; fi63 < 10; fi63++ {
                                if false {
                                    var v64 uint32 = 2
                                    if v46 {
                                        var v65 uint = 335
                                        _ = v65
                                    }
                                    v66 := GenS1{F0: uint64(v2)}
                                    if v46 {
                                        var v67 int32 = 335
                                        fmt.Println(v44, v3)
                                        var v68 uint16 = v47
                                        _ = v67
                                        _ = v68
                                    }
                                    _ = v64
                                    _ = v66
                                }
                                _ = fi63
                            }
                        case 2147483647:
                            var v69 uint64 = v44
                            _ = v69
                        default:
                            for fi70 := 0; fi70 < 2; fi70++ {
                                switch (^(fi57)) {
                                case 143:
                                    switch ((v26 + v26) + v26) {
                                    case "foo bar":
                                        var v71 string = "a"
                                        _ = v71
                                    default:
                                        switch fi57 {
                                        case 59:
                                            v26 = (v26 + v26)
                                        case 1:
                                            v2 += 32767
                                        case -633:
                                            switch fi70 {
                                            case -469:
                                                if (!v46) {
                                                    fmt.Println(v3, v3, fi57)
                                                    if (!(4294967295 == 0)) {
                                                        var v72 int8 = -117
                                                        _ = len(v1)
                                                        _ = v72
                                                    }
                                                    v48 += (v48 - v48)
                                                }
                                            case -1:
                                                var v73 int8 = -14
                                                _ = v73
                                            }
                                        default:
                                            v4 += fi70
                                        }
                                    }
                                case 128:
                                    if len(v1) > 0 {
                                        v1[0] = 497
                                    }
                                case 127:
                                    var v74 float32 = v48
                                    _ = v74
                                default:
                                    v75 := GenS2{F0: (v26 + (v26 + "foo bar")), F1: -1, F2: -32768}
                                    _ = v75
                                }
                                _ = len(v1)
                                fmt.Println(v2, fi70, v44)
                                _ = fi70
                            }
                        }
                    case 127:
                        var v76 string = (v26 + "foo bar")
                        _ = v76
                    case -2147483648:
                        v49 += v49
                    default:
                        switch v26 {
                        case "":
                            v77 := GenS5{F0: v4, F1: (v49 + 0)}
                            _ = v77
                        case "foo bar":
                            v48 = -2.25
                        default:
                            fmt.Println(v44, v26, v4)
                        }
                    }
                }
                fmt.Println(v49)
            } else {
                switch "a" {
                case "\n\t":
                    v4 -= (^(fi57))
                default:
                    v2 += 65535
                }
                switch (fi57 >> ((2) & 63)) {
                case 4294967295:
                    switch (("\n\t" + "foo bar") + v26) {
                    case "a":
                        v4 = fi57
                    default:
                        _ = len(v1)
                    }
                case 127:
                    var v78 string = v26
                    _ = v78
                default:
                    fmt.Println(v47, v4)
                }
            }
        }
        v49 = v49
        _ = fi57
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
}
