package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint; F1 int32; F2 bool }
type GenS3 struct { F0 uint8; F1 uint; F2 int32 }
type GenS4 struct { F0 uint; F1 uint8; F2 uint }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int16; F1 int64 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint8; F1 uint64 }
func genFn0() int {
    for fi4 := 0; fi4 < 12; fi4++ {
        var v5 string = "a"
        _ = fi4
        _ = v5
    }
    if false {
        v6 := GenS1{F0: 84, F1: 1}
        _ = v6
    }
    v7 := GenS2{F0: 591, F1: -2147483648, F2: false}
    var v8 int16 = -32768
    _ = v7
    _ = v8
    return -827
}
func genFn1(p0 uint) uint {
    p0 = p0
    var v9 int = 32767
    switch "foo bar" {
    case "hello":
        var v10 bool = (false || true)
        _ = v10
    case "a":
        v9 = v9
    }
    var v11 uint = 4294967295
    v12 := GenS2{F0: (v11 ^ p0), F1: 0, F2: true}
    _ = p0
    _ = v11
    _ = v12
    _ = v9
    return (p0 - 18446744073709551615)
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int = -9223372036854775808
    v3 := make(map[int]bool)
    v13 := GenS2{F0: 400, F1: 42, F2: (-166 >= 808)}
    v3[v2] = (0 > 65535)
    switch (v2 - -226) {
    case 771:
        var v14 int16 = 809
        _ = v14
    case 767:
        v2 -= (v2 & v2)
    default:
        fmt.Println(v2)
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        v1[393] = ("foo bar" + ("\n\t" + "hello"))
        _ = fi15
    }
    if (65535 != 68) {
        v2 = (v2 >> ((4294967295) & 63))
        v2 = v2
        for fi16 := 0; fi16 < 17; fi16++ {
            var v17 uint8 = 115
            _ = fi16
            _ = v17
        }
    } else {
        fmt.Println(v2)
        switch "hello" {
        case "\n\t":
            var v18 uint32 = 2147483648
            _ = v18
        case "hello":
            switch 958 {
            case 1:
                var v19 int32 = -480
                _ = v19
            case 127:
                if true {
                    v20 := GenS2{F0: 65535, F1: 1, F2: true}
                    _ = v20
                }
            }
        default:
            _ = len(v3)
        }
    }
    if (!(true || true)) {
        var v21 uint8 = 128
        _ = v21
    } else {
        for fi22 := 0; fi22 < 11; fi22++ {
            switch (fi22 - v2) {
            case 32767:
                fmt.Println(v2, fi22)
            case 4294967295:
                fmt.Println(fi22, fi22)
            case 0:
                switch fi22 {
                case 732:
                    v23 := GenS3{F0: 42, F1: 65535, F2: 35}
                    _ = v23
                case 9223372036854775807:
                    v24 := []int{(^(fi22)), 876, fi22}
                    _ = v24
                }
            default:
                var v25 int16 = 433
                _ = v25
            }
            v2 -= (v2 * fi22)
            if (-96 <= 335) {
                if true {
                    v26 := []int{114, genFn0(), fi22, fi22}
                    var v27 uint16 = 409
                    var v28 uint32 = 630
                    _ = v26
                    _ = v27
                    _ = v28
                }
            } else {
                switch "" {
                case "foo bar":
                    switch (v2 >> ((701) & 63)) {
                    case -1:
                        var v29 int32 = 582
                        _ = v29
                    }
                case "":
                    fmt.Println(v2)
                case "\n\t":
                    fmt.Println(fi22, v2)
                default:
                    switch "\n\t" {
                    case "foo bar":
                        v2 = -9223372036854775808
                    case "hello":
                        var v30 float64 = 0.1
                        _ = v30
                    case "\n\t":
                        v31 := GenS4{F0: 4294967295, F1: 128, F2: 881}
                        _ = v31
                    }
                }
            }
            _ = fi22
        }
        fmt.Println(v2)
        var v32 int32 = 2147483647
        _ = v32
    }
    if true {
        v33 := GenS5{F0: 42}
        var v34 uint64 = 1
        delete(v1, 2147483647)
        _ = v33
        _ = v34
    } else {
        if ((0 >= 44) || ((true || true) || true)) {
            if true {
                v2 = genFn0()
                if (v2 <= v2) {
                    fmt.Println(v2)
                    var v35 uint64 = 1
                    if (!false) {
                        v36 := GenS4{F0: 65535, F1: 193, F2: genFn1(717)}
                        _ = v36
                    } else {
                        v37 := v3[(^(v2))]
                        fmt.Println(v2, v2)
                        var v38 int8 = 0
                        _ = v37
                        _ = v38
                    }
                    _ = v35
                } else {
                    v2 = genFn0()
                    delete(v1, 455)
                    var v39 float32 = (1e10 + 3.14159)
                    _ = v39
                }
            } else {
                var v40 int16 = -1
                for fi41 := 0; fi41 < 17; fi41++ {
                    delete(v1, 2147483648)
                    v2 = (v2 - fi41)
                    _ = fi41
                }
                switch (v2 / ((v2) | 1)) {
                case 9223372036854775807:
                    var v42 int = 382
                    _ = v42
                case 4294967295:
                    switch 32767 {
                    case 2147483647:
                        v2 += (v2 & -498)
                    }
                case -2147483648:
                    var v43 int16 = v40
                    _ = v43
                default:
                    var v44 string = ""
                    _ = v44
                }
                _ = v40
            }
        } else {
            for fi45 := 0; fi45 < 14; fi45++ {
                for fi46 := 0; fi46 < 7; fi46++ {
                    v2 = (fi45 | -2147483648)
                    v2 -= (fi46 + fi46)
                    if true {
                        switch fi46 {
                        case -197:
                            v2 += genFn0()
                            v47 := GenS1{F0: 65, F1: 2}
                            var v48 bool = false
                            _ = v47
                            _ = v48
                        case 978:
                            v3[0] = true
                        default:
                            v2 = (fi46 | (fi46 / (((fi46 / ((v2) | 1))) | 1)))
                            v2 -= 1
                            v2 = fi45
                        }
                        var v49 int32 = -492
                        var v50 uint32 = 0
                        _ = v49
                        _ = v50
                    }
                    _ = fi46
                }
                v2 += v2
                switch (fi45 / ((-2147483648) | 1)) {
                case 2147483647:
                    for fi51 := 0; fi51 < 15; fi51++ {
                        var v52 int64 = 1
                        fmt.Println(fi45)
                        switch "\n\t" {
                        case "hello":
                            v2 -= (fi45 >> ((0) & 63))
                        case "a":
                            v53 := GenS1{F0: 42, F1: 0}
                            _ = v53
                        case "foo bar":
                            v52 -= v52
                        default:
                            var v54 float64 = float64(float32(1.5))
                            _ = v54
                        }
                        _ = fi51
                        _ = v52
                    }
                case -891:
                    if false {
                        v55 := v1[1]
                        _ = v55
                    } else {
                        if false {
                            v56 := GenS1{F0: 144, F1: 42}
                            for fi57 := 0; fi57 < 12; fi57++ {
                                var v58 uint = 18446744073709551615
                                _ = fi57
                                _ = v58
                            }
                            for fi59 := 0; fi59 < 9; fi59++ {
                                var v60 uint8 = 255
                                if (561 == 9223372036854775807) {
                                    v61 := GenS3{F0: (v60 - 2), F1: 4294967295, F2: 42}
                                    _ = v61
                                }
                                v60 = 1
                                var v62 bool = (-862 < 1000)
                                _ = fi59
                                _ = v60
                                _ = v62
                            }
                            _ = v56
                        } else {
                            v63 := GenS4{F0: 0, F1: 128, F2: 18446744073709551615}
                            var v64 bool = false
                            _ = v63
                            _ = v64
                        }
                    }
                default:
                    v65 := GenS2{F0: 42, F1: -1, F2: true}
                    _ = v65
                }
                _ = fi45
            }
        }
        if true {
            var v66 uint8 = 42
            _ = v66
        } else {
            v2 -= -621
            if false {
                for fi67 := 0; fi67 < 14; fi67++ {
                    v2 = fi67
                    _ = fi67
                }
                v68 := GenS3{F0: 74, F1: 123, F2: -1}
                var v69 float32 = (-2.25 * float32(1.5))
                _ = v68
                _ = v69
            } else {
                for fi70 := 0; fi70 < 3; fi70++ {
                    v71 := GenS2{F0: genFn1(403), F1: -2147483648, F2: (-1 < 962)}
                    _ = fi70
                    _ = v71
                }
                fmt.Println(v2)
                v2 = genFn0()
            }
            var v72 float64 = ((3.14159 - 0.1) / 1.5)
            _ = v72
        }
        fmt.Println(v2)
    }
    delete(v1, 42)
    v73 := GenS6{F0: 1000, F1: -713}
    fmt.Println(v2)
    v2 -= (v2 * (^(v2)))
    var v74 uint32 = 915
    v75 := GenS7{F0: 18446744073709551615}
    var v76 int16 = -1000
    v74 -= v74
    v77 := GenS8{F0: 0, F1: 4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v74:", v74)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v73
    _ = v74
    _ = v75
    _ = v76
    _ = v77
}
