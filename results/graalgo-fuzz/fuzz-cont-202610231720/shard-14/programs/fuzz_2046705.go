package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 bool; F1 uint8; F2 uint16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS7 struct { F0 int; F1 uint }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int64; F1 uint32 }
func genFn0(p0 bool, p1 uint64, p2 uint) uint32 {
    var v4 float32 = 3.14159
    var v5 int = 1
    for fi6 := 0; fi6 < 2; fi6++ {
        for fi7 := 0; fi7 < 8; fi7++ {
            switch int(p1) {
            case -1:
                v8 := []int{v5, 1, fi6}
                _ = v8
            case 32768:
                if p0 {
                    switch (v5 / ((127) | 1)) {
                    case 737:
                        v4 = ((v4 * -2.25) + 1.5)
                    case -229:
                        if (!false) {
                            if false {
                                var v9 int16 = 1000
                                _ = v9
                            } else {
                                var v10 uint32 = 187
                                var v11 int = (^(v5))
                                _ = v10
                                _ = v11
                            }
                            v12 := GenS1{F0: 0}
                            var v13 int64 = 1
                            _ = v12
                            _ = v13
                        }
                    case 185:
                        v5 += (fi6 - (fi6 + (fi6 << ((p2) & 63))))
                    default:
                        v14 := GenS2{F0: 42, F1: 1}
                        _ = v14
                    }
                    var v15 int16 = -1
                    v16 := GenS1{F0: 0}
                    _ = v15
                    _ = v16
                }
            case 9223372036854775807:
                v17 := GenS3{F0: 2147483647, F1: -1}
                _ = v17
            }
            _ = fi7
        }
        var v18 int = 32768
        _ = fi6
        _ = v18
    }
    if true {
        var v19 int8 = -128
        fmt.Println(v19, p0, v19)
        var v20 int32 = 2147483647
        _ = v19
        _ = v20
    } else {
        v21 := []int{-1, (v5 >> ((0) & 63)), (v5 + (v5 >> ((255) & 63)))}
        _ = v21
    }
    for fi22 := 0; fi22 < 7; fi22++ {
        for fi23 := 0; fi23 < 19; fi23++ {
            fmt.Println(p0)
            _ = fi23
        }
        if (p2 == p2) {
            p1 -= (p1 / ((p1) | 1))
            v4 = (v4 * v4)
        } else {
            if false {
                v24 := GenS1{F0: -128}
                v25 := GenS2{F0: 42, F1: 2}
                v26 := []int{v5, fi22, v5}
                _ = v24
                _ = v25
                _ = v26
            }
        }
        fmt.Println(fi22, p0)
        _ = fi22
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 489
}
func genFn1(p0 uint, p1 string) uint32 {
    var v27 float64 = float64(0.1)
    var v28 uint16 = 1
    v29 := GenS4{F0: (42 != -2147483648), F1: 128, F2: (^(v28))}
    v28 = v28
    _ = p0
    _ = p1
    _ = v27
    _ = v28
    _ = v29
    return 2147483648
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []int{-97, -1}
    v3 := make(map[int]string)
    if (!true) {
        v30 := GenS5{F0: 32767}
        _ = len(v1)
        for fi31 := 0; fi31 < 20; fi31++ {
            v32 := GenS3{F0: -329, F1: -74}
            v33 := GenS3{F0: -283, F1: 127}
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = v30
    }
    switch ((("hello" + "") + "\n\t") + ("foo bar" + (("foo bar" + "hello") + "a"))) {
    case "a":
        var v34 int32 = 42
        _ = v34
    default:
        var v35 uint32 = 2
        _ = v35
    }
    for fi36 := 0; fi36 < 14; fi36++ {
        switch "" {
        case "\n\t":
            for fi37 := 0; fi37 < 15; fi37++ {
                v38 := GenS5{F0: 223}
                v3[fi36] = ""
                if (!true) {
                    _ = len(v2)
                    var v39 int8 = 42
                    if (true && (65535 != 899)) {
                        fmt.Println(v39, v39, v39)
                        v40 := GenS2{F0: 1, F1: 42}
                        _ = v40
                    } else {
                        if (v39 >= v39) {
                            var v41 uint64 = 82
                            v41 = v41
                            v41 = v41
                            _ = len(v2)
                            _ = v41
                        } else {
                            var v42 bool = (("foo bar" + "a") != "")
                            _ = v42
                        }
                        v39 = (v39 * v39)
                        v39 -= int8(fi37)
                    }
                    _ = v39
                } else {
                    fmt.Println(fi37)
                    fmt.Println(fi36)
                }
                _ = fi37
                _ = v38
            }
        case "a":
            v43 := GenS6{F0: 2, F1: 32767, F2: 285}
            _ = v43
        case "foo bar":
            fmt.Println(fi36)
        }
        v44 := GenS6{F0: 2, F1: 42, F2: 695}
        v2 = append(v2, fi36)
        _ = fi36
        _ = v44
    }
    v45 := GenS1{F0: -1}
    var v46 uint32 = 741
    switch "hello" {
    case "hello":
        v2 = append(v2, 701)
    case "\n\t":
        v47 := GenS2{F0: 2147483647, F1: v46}
        _ = v47
    case "a":
        fmt.Println(v46)
    }
    var v48 float32 = 0.1
    if false {
        v3[0] = "a"
        for fi49 := 0; fi49 < 17; fi49++ {
            var v50 uint16 = 32767
            for fi51 := 0; fi51 < 10; fi51++ {
                fmt.Println(fi51, fi49, v48)
                var v52 int16 = 32767
                v52 += int16(fi49)
                _ = fi51
                _ = v52
            }
            _ = fi49
            _ = v50
        }
    }
    for fi53 := 0; fi53 < 16; fi53++ {
        v46 += v46
        _ = fi53
    }
    if (2 > 127) {
        fmt.Println(v48)
        v48 += 1e10
        if false {
            switch "hello" {
            case "foo bar":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case "\n\t":
                fmt.Println(v48)
            case "a":
                fmt.Println(v48, v48)
            default:
                var v54 bool = true
                _ = v54
            }
            v46 = (v46 * 1)
            for fi55 := 0; fi55 < 14; fi55++ {
                var v56 int = fi55
                if false {
                    for fi57 := 0; fi57 < 20; fi57++ {
                        if false {
                            v56 = 127
                        } else {
                            var v58 int64 = 0
                            _ = v58
                        }
                        _ = fi57
                    }
                    if true {
                        var v59 bool = false
                        for fi60 := 0; fi60 < 12; fi60++ {
                            if ((v56 * 42) != -1) {
                                fmt.Println(v46, v48)
                                var v61 uint32 = (v46 & v46)
                                _ = v61
                            }
                            var v62 int64 = 42
                            if true {
                                v63 := GenS7{F0: fi60, F1: 4294967295}
                                _ = v63
                            } else {
                                v64 := GenS1{F0: 42}
                                var v65 string = ""
                                var v66 int16 = 32767
                                _ = v64
                                _ = v65
                                _ = v66
                            }
                            _ = fi60
                            _ = v62
                        }
                        v67 := GenS3{F0: 0, F1: 0}
                        _ = v59
                        _ = v67
                    }
                }
                if false {
                    if (true || (!((v46 + v46) < v46))) {
                        var v68 uint64 = 0
                        var v69 uint64 = 760
                        v70 := GenS3{F0: 2147483647, F1: 127}
                        _ = v68
                        _ = v69
                        _ = v70
                    }
                }
                _ = fi55
                _ = v56
            }
        } else {
            v2 = append(v2, -9223372036854775808)
            for fi71 := 0; fi71 < 7; fi71++ {
                v46 -= v46
                if (true || false) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if true {
                        var v72 uint = 4294967295
                        if true {
                            var v73 int8 = -119
                            for fi74 := 0; fi74 < 15; fi74++ {
                                switch (-(fi74)) {
                                case -787:
                                    v75 := GenS8{F0: (-(v73))}
                                    _ = v75
                                }
                                var v76 int = (fi74 / (((fi74 / (((fi71 * fi71)) | 1))) | 1))
                                switch ("a" + "") {
                                case "foo bar":
                                    v1[724] = true
                                }
                                _ = fi74
                                _ = v76
                            }
                            _ = v73
                        } else {
                            var v77 uint32 = v46
                            v48 = v48
                            _ = v77
                        }
                        _ = v72
                    } else {
                        _ = len(v2)
                    }
                    if (!false) {
                        v48 -= 1.5
                        var v78 uint = 2
                        _ = v78
                    }
                }
                _ = fi71
            }
            v79 := GenS7{F0: -514, F1: 1}
            _ = v79
        }
    }
    v46 = (v46 << ((2) & 63))
    v48 = 0.1
    fmt.Println(v48, v46)
    fmt.Println(v48, v46)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if false {
        if false {
            if ((-1 == 9223372036854775807) || false) {
                switch 127 {
                case -662:
                    v46 += v46
                case 32768:
                    delete(v3, -77)
                case 128:
                    var v80 uint8 = 42
                    _ = v80
                }
                v48 += (0.1 - -2.25)
            } else {
                if ((v46 | (^(v46))) != (v46 + (v46 ^ (v46 + v46)))) {
                    v81 := GenS9{F0: 9223372036854775807, F1: (v46 >> ((42) & 63))}
                    _ = v81
                }
            }
            if (772 <= 18446744073709551615) {
                v48 += (3.14159 * 0.0)
                var v82 string = ""
                _ = v82
            }
        } else {
            for fi83 := 0; fi83 < 7; fi83++ {
                for fi84 := 0; fi84 < 3; fi84++ {
                    _ = len(v3)
                    var v85 uint32 = (v46 % ((genFn0((true && true), 18446744073709551615, uint(v46))) | 1))
                    _ = fi84
                    _ = v85
                }
                if false {
                    v86 := GenS8{F0: -105}
                    var v87 uint = 623
                    _ = v86
                    _ = v87
                }
                _ = fi83
            }
            v88 := GenS8{F0: -128}
            _ = v88
        }
        _ = len(v3)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v45
    _ = v46
    _ = v48
}
