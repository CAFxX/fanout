package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 uint16; F1 int16 }
type GenS6 struct { F0 int16; F1 uint16; F2 uint32 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint8; F1 int8 }
type GenS9 struct { F0 int32; F1 int; F2 int16 }

func main() {
    var v1 uint16 = 42
    var v2 uint8 = 42
    var v3 uint16 = (v1 + v1)
    v4 := make(map[string]string)
    var v5 string = "foo bar"
    var v6 uint32 = 85
    if true {
        var v7 float64 = 3.14159
        _ = v7
    }
    for fi8 := 0; fi8 < 16; fi8++ {
        v5 = ""
        _ = fi8
    }
    v9 := GenS1{F0: -2147483648}
    v6 = (v6 << ((255) & 63))
    if true {
        v5 = v5
    } else {
        var v10 int64 = -1
        if (1 >= (^(v10))) {
            switch "a" {
            case "\n\t":
                var v11 bool = false
                _ = v11
            default:
                var v12 float32 = 3.14159
                _ = v12
            }
            var v13 bool = (127 < 42)
            _ = v13
        }
        _ = v10
    }
    v4["\n\t"] = ("\n\t" + v5)
    switch 32767 {
    case 128:
        var v14 int64 = -785
        _ = v14
    case 9223372036854775807:
        if (v5 == (("\n\t" + v5) + "\n\t")) {
            for fi15 := 0; fi15 < 17; fi15++ {
                var v16 uint8 = (^(v2))
                _ = fi15
                _ = v16
            }
        } else {
            for fi17 := 0; fi17 < 3; fi17++ {
                var v18 string = ("hello" + "")
                if false {
                    if ((2 > 862) || (32768 <= (fi17 ^ (fi17 << ((2) & 63))))) {
                        switch "hello" {
                        case "a":
                            v19 := GenS1{F0: 2147483647}
                            _ = v19
                        default:
                            if false {
                                fmt.Println(v5, v5)
                                if false {
                                    if false {
                                        var v20 uint16 = v3
                                        if true {
                                            v21 := []int{fi17, fi17, fi17}
                                            v5 = "\n\t"
                                            _ = v21
                                        } else {
                                            for fi22 := 0; fi22 < 12; fi22++ {
                                                v2 = (v2 << ((1) & 63))
                                                if true {
                                                    v23 := []int{fi22, int(v20), -966}
                                                    _ = v23
                                                } else {
                                                    v24 := GenS2{F0: (^(v2))}
                                                    switch 128 {
                                                    case -74:
                                                        var v25 float32 = 3.14159
                                                        _ = v25
                                                    case -9223372036854775808:
                                                        v26 := []int{-9223372036854775808, -848, 42, 32768}
                                                        _ = v26
                                                    case 0:
                                                        v6 += (v6 ^ uint32(v2))
                                                    default:
                                                        v27 := GenS2{F0: v2}
                                                        _ = v27
                                                    }
                                                    _ = v24
                                                }
                                                var v28 string = v5
                                                v4[(v5 + (v28 + (((v5 + (v28 + "a")) + v5) + v18)))] = (v5 + ((v28 + v5) + v18))
                                                _ = fi22
                                                _ = v28
                                            }
                                        }
                                        _ = v20
                                    }
                                } else {
                                    v2 += v2
                                    fmt.Println(v3, v5, v3)
                                }
                                var v29 uint16 = v1
                                _ = v29
                            }
                        }
                        v30 := GenS3{F0: 18446744073709551615, F1: 565}
                        _ = v30
                    }
                    v18 = (v18 + ((v5 + "foo bar") + "hello"))
                } else {
                    fmt.Println(v6)
                    v3 = v3
                    v31 := []int{(^(fi17)), fi17, fi17, 635}
                    _ = v31
                }
                v32 := []int{(^(fi17))}
                _ = fi17
                _ = v18
                _ = v32
            }
            switch v5 {
            case "a":
                if true {
                    if (false || false) {
                        v1 -= 1
                        var v33 string = "foo bar"
                        v34 := GenS4{F0: 18446744073709551615, F1: 4294967295}
                        _ = v33
                        _ = v34
                    } else {
                        for fi35 := 0; fi35 < 14; fi35++ {
                            v36 := GenS5{F0: 255, F1: -28}
                            if false {
                                fmt.Println(v1, v2)
                                if false {
                                    var v37 uint64 = 19
                                    var v38 uint32 = v6
                                    _ = v37
                                    _ = v38
                                } else {
                                    for fi39 := 0; fi39 < 16; fi39++ {
                                        var v40 uint = 695
                                        var v41 bool = (true || true)
                                        _ = fi39
                                        _ = v40
                                        _ = v41
                                    }
                                }
                            }
                            if true {
                                var v42 float32 = float32((((float64(1.5) - 1e10) - 1e10) * (float64(0.1) / -2.25)))
                                _ = v42
                            } else {
                                for fi43 := 0; fi43 < 4; fi43++ {
                                    delete(v4, (v5 + v5))
                                    v2 -= 0
                                    _ = fi43
                                }
                            }
                            _ = fi35
                            _ = v36
                        }
                        v1 = v3
                    }
                    var v44 string = "a"
                    v45 := []int{-9223372036854775808, 0, 32767}
                    _ = v44
                    _ = v45
                }
            case "\n\t":
                v46 := []int{32767, -949}
                _ = v46
            }
        }
    case 528:
        var v47 uint64 = 20
        _ = v47
    }
    v48 := []int{-1}
    v49 := GenS6{F0: -32768, F1: (v1 | uint16(v6)), F2: (^(v6))}
    for fi50 := 0; fi50 < 6; fi50++ {
        v3 += v3
        switch fi50 {
        case 0:
            v51 := GenS7{F0: fi50}
            _ = v51
        case -2147483648:
            for fi52 := 0; fi52 < 19; fi52++ {
                var v53 string = v5
                _ = fi52
                _ = v53
            }
        default:
            var v54 int16 = 32767
            _ = v54
        }
        _ = fi50
    }
    fmt.Println(v1, v3)
    for fi55 := 0; fi55 < 10; fi55++ {
        v56 := GenS8{F0: (v2 ^ 1), F1: 127}
        _ = fi55
        _ = v56
    }
    _ = len(v48)
    switch v5 {
    case "foo bar":
        switch v5 {
        case "a":
            for fi57 := 0; fi57 < 14; fi57++ {
                for fi58 := 0; fi58 < 13; fi58++ {
                    if (v5 == v5) {
                        if (17 > 0) {
                            v5 = v5
                        } else {
                            v6 = (v6 - v6)
                            var v59 uint64 = 4294967295
                            v4[v5] = (v5 + v5)
                            _ = v59
                        }
                        if len(v48) >= 2 {
                            v48 = v48[1:len(v48)]
                        }
                        fmt.Println(fi57, fi57)
                        switch fi58 {
                        case -2147483648:
                            v1 = (v1 << ((uint(v2)) & 63))
                        default:
                            v60 := GenS4{F0: 1, F1: 0}
                            _ = v60
                        }
                        v5 = v5
                        if (!false) {
                            var v61 uint = 954
                            var v62 int8 = 127
                            var v63 uint16 = (v1 & 1)
                            _ = v61
                            _ = v62
                            _ = v63
                        } else {
                            var v64 int32 = -2147483648
                            _ = v64
                        }
                        var v65 uint = 70
                        _ = v65
                    } else {
                        if false {
                            if len(v48) >= 2 {
                                v48 = v48[1:len(v48)]
                            }
                        } else {
                            fmt.Println(v6, v5)
                            var v66 int8 = -74
                            _ = v66
                        }
                        v6 = 299
                    }
                    v67 := GenS2{F0: 1}
                    v1 -= uint16(v6)
                    _ = fi58
                    _ = v67
                }
                _ = fi57
            }
        case "\n\t":
            v2 = uint8(v3)
        default:
            if (9223372036854775807 != -892) {
                var v68 uint8 = v2
                v3 = (v3 << ((414) & 63))
                _ = v68
            } else {
                var v69 int8 = 0
                var v70 uint64 = 0
                if ((((1e10 - (0.1 - ((1.5 * 1e10) - -2.25))) - 3.14159) + 0.1) > 0.1) {
                    v6 += v6
                    var v71 int64 = 0
                    if ((v71 * (v71 ^ v71)) > v71) {
                        var v72 int16 = 42
                        for fi73 := 0; fi73 < 7; fi73++ {
                            var v74 uint = uint(v70)
                            var v75 int8 = (v69 % ((-40) | 1))
                            _ = fi73
                            _ = v74
                            _ = v75
                        }
                        _ = v72
                    }
                    _ = v71
                } else {
                    if (161 == (v2 * 194)) {
                        for fi76 := 0; fi76 < 2; fi76++ {
                            for fi77 := 0; fi77 < 14; fi77++ {
                                v48 = append(v48, fi76)
                                var v78 int = (fi76 * -2147483648)
                                _ = fi77
                                _ = v78
                            }
                            _ = fi76
                        }
                        v70 += (v70 % (((v70 / ((18446744073709551615) | 1))) | 1))
                        var v79 float32 = 0.0
                        _ = v79
                    } else {
                        var v80 bool = true
                        _ = v80
                    }
                    fmt.Println(v70)
                }
                _ = v69
                _ = v70
            }
        }
    case "":
        v1 -= v3
    default:
        if false {
            if true {
                v81 := GenS6{F0: 0, F1: v3, F2: 4294967295}
                _ = v81
            } else {
                v82 := GenS9{F0: -2147483648, F1: 42, F2: 32767}
                _ = v82
            }
            v5 = ("\n\t" + (v5 + ""))
        }
    }
    v83 := GenS8{F0: v2, F1: 47}
    for fi84 := 0; fi84 < 11; fi84++ {
        if true {
            fmt.Println(v3, v5, v2)
            var v85 float64 = 0.1
            var v86 int8 = -128
            _ = v85
            _ = v86
        }
        _ = fi84
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v48:", len(v48), v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v6
    _ = v83
    _ = v9
}
