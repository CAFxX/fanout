package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 int }
type GenS2 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 int32; F1 int8; F2 int16 }
type GenS6 struct { F0 int8; F1 uint16 }

func main() {
    var v1 int8 = 0
    v2 := make(map[uint32]uint32)
    var v3 int64 = 497
    v4 := make(map[string]uint32)
    if (1 >= 32767) {
        if (9223372036854775807 < v3) {
            var v5 float32 = 1.5
            _ = v5
        } else {
            fmt.Println(v3, v3)
        }
    }
    var v6 float32 = 1e10
    var v7 int8 = (v1 >> ((2) & 63))
    v8 := GenS1{F0: -710, F1: 127, F2: -642}
    for fi9 := 0; fi9 < 6; fi9++ {
        for fi10 := 0; fi10 < 8; fi10++ {
            v6 = 1.5
            _ = fi10
        }
        v11 := []int{9223372036854775807, 1}
        _ = fi9
        _ = v11
    }
    if ((343 < 2) || (((1.5 - (1.5 - 0.1)) > (3.14159 * float64(v6))) && (726 != 1))) {
        if (255 > 86) {
            var v12 uint = 2
            _ = v12
        } else {
            if (!false) {
                fmt.Println(v3, v1, v3)
                if (!(false || true)) {
                    v13 := GenS1{F0: -9223372036854775808, F1: 1, F2: -1}
                    for fi14 := 0; fi14 < 8; fi14++ {
                        v3 -= v3
                        _ = fi14
                    }
                    _ = v13
                } else {
                    v1 += 1
                    var v15 bool = (922 > 42)
                    _ = v15
                }
            } else {
                v16 := GenS1{F0: -1, F1: 1, F2: 42}
                _ = v16
            }
            for fi17 := 0; fi17 < 14; fi17++ {
                v6 -= 1e10
                var v18 uint8 = 2
                v19 := GenS2{F0: 2, F1: 32767, F2: 2}
                _ = fi17
                _ = v18
                _ = v19
            }
        }
    }
    if false {
        v20 := []int{1, 32768, -9223372036854775808, 2147483647}
        _ = v20
    }
    fmt.Println(v3, v6, v1)
    switch "foo bar" {
    case "\n\t":
        switch "hello" {
        case "hello":
            var v21 uint16 = 65535
            _ = v21
        case "":
            fmt.Println(v7, v6, v7)
        default:
            for fi22 := 0; fi22 < 3; fi22++ {
                fmt.Println(fi22, v7)
                _ = fi22
            }
        }
    case "a":
        v23 := GenS2{F0: 798, F1: 32767, F2: 1}
        _ = v23
    }
    for fi24 := 0; fi24 < 20; fi24++ {
        var v25 string = ("" + ("\n\t" + "hello"))
        v1 = (v1 ^ v1)
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 8; fi26++ {
        if true {
            var v27 uint32 = 351
            fmt.Println(v27, v27, v3)
            v7 += int8(v3)
            _ = v27
        }
        fmt.Println(fi26)
        _ = fi26
    }
    var v28 int16 = -1
    _ = len(v4)
    if true {
        if (1 < v3) {
            v29 := GenS3{F0: "", F1: 65535}
            if (0 != 1) {
                for fi30 := 0; fi30 < 13; fi30++ {
                    switch fi30 {
                    case 783:
                        if (false && (2 <= 42)) {
                            for fi31 := 0; fi31 < 13; fi31++ {
                                fmt.Println(fi30, v6)
                                var v32 int16 = int16(v3)
                                _ = fi31
                                _ = v32
                            }
                            switch fi30 {
                            case 1:
                                for fi33 := 0; fi33 < 19; fi33++ {
                                    v34 := GenS4{F0: 1, F1: 0}
                                    v35 := []int{-725, -965, fi33}
                                    var v36 int16 = v28
                                    _ = fi33
                                    _ = v34
                                    _ = v35
                                    _ = v36
                                }
                            case 32767:
                                var v37 float32 = (v6 * 1e10)
                                _ = v37
                            }
                            if (!(!(v3 <= 1))) {
                                var v38 bool = (!true)
                                _ = v38
                            } else {
                                if false {
                                    fmt.Println(v1, v1)
                                } else {
                                    v7 = 0
                                    var v39 bool = (2 != 304)
                                    _ = v39
                                }
                                v1 -= v1
                            }
                        } else {
                            for fi40 := 0; fi40 < 10; fi40++ {
                                var v41 int64 = v3
                                v6 = ((v6 + ((v6 * (v6 / ((float32((0.0 + (3.14159 / ((((1e10 * ((3.14159 + 1.5) * (((3.14159 + 1e10) + 0.0) - 1e10))) / float64(((((v6 * 1e10) - 1.5) * v6) * v6))) + -2.25) * (0.1 - 0.0))))) * 1e10) - (v6 + ((1e10 - 0.1) / (v6 + 1.5)))))) + v6)) / v6)
                                v3 += 80
                                _ = fi40
                                _ = v41
                            }
                        }
                    default:
                        if (fi30 != 107) {
                            v42 := []int{4294967295, int(v28), (^(fi30))}
                            _ = v42
                        } else {
                            v28 += v28
                            var v43 int = fi30
                            _ = v43
                        }
                    }
                    v28 = v28
                    for fi44 := 0; fi44 < 13; fi44++ {
                        switch "a" {
                        case "\n\t":
                            v28 = (^(v28))
                        default:
                            if (((false && true) && false) && (42 == 4294967295)) {
                                fmt.Println(v3, v7, v3)
                                var v45 int16 = (-(v28))
                                v46 := []int{(fi30 - (fi30 % ((fi30) | 1))), (fi30 - fi30), (fi44 / ((fi44) | 1))}
                                _ = v45
                                _ = v46
                            } else {
                                v47 := GenS5{F0: -1, F1: (v1 - v1), F2: (-(v28))}
                                v48 := GenS6{F0: v7, F1: 253}
                                v49 := GenS5{F0: 0, F1: (v7 % (((v1 ^ (v1 ^ 42))) | 1)), F2: v28}
                                _ = v47
                                _ = v48
                                _ = v49
                            }
                        }
                        _ = fi44
                    }
                    _ = fi30
                }
                for fi50 := 0; fi50 < 5; fi50++ {
                    v28 += 1000
                    _ = fi50
                }
                _ = len(v2)
            }
            v51 := GenS4{F0: 255, F1: -1}
            _ = v29
            _ = v51
        }
        v52 := GenS6{F0: v1, F1: 706}
        v53 := GenS1{F0: 32767, F1: 42, F2: 4294967295}
        _ = v52
        _ = v53
    } else {
        if (2 == 18446744073709551615) {
            v6 = v6
            var v54 int16 = 95
            switch "hello" {
            case "\n\t":
                var v55 float64 = (0.1 * 0.1)
                _ = v55
            case "hello":
                v7 += v7
            case "a":
                if (v7 <= v1) {
                    if false {
                        switch int(v1) {
                        case -2147483648:
                            if (float32(float64(v6)) != v6) {
                                switch -2147483648 {
                                case 127:
                                    for fi56 := 0; fi56 < 7; fi56++ {
                                        v1 = v1
                                        switch (fi56 / ((fi56) | 1)) {
                                        case -589:
                                            fmt.Println(v28)
                                        }
                                        fmt.Println(v1, v6)
                                        _ = fi56
                                    }
                                case 1:
                                    var v57 uint = 42
                                    _ = v57
                                case 9223372036854775807:
                                    if (!(!false)) {
                                        var v58 uint8 = 210
                                        _ = v58
                                    } else {
                                        var v59 float32 = v6
                                        if true {
                                            v60 := GenS1{F0: 9223372036854775807, F1: 144, F2: 945}
                                            _ = v60
                                        }
                                        v7 += v1
                                        _ = v59
                                    }
                                }
                                for fi61 := 0; fi61 < 7; fi61++ {
                                    var v62 string = "hello"
                                    v3 += (v3 << ((637) & 63))
                                    _ = fi61
                                    _ = v62
                                }
                            } else {
                                v63 := []int{}
                                if len(v63) > 0 {
                                    v63[0] = 2147483647
                                }
                                var v64 uint = 1
                                _ = v63
                                _ = v64
                            }
                        case 4294967295:
                            v65 := GenS4{F0: 0, F1: -2147483648}
                            _ = v65
                        case 127:
                            switch 1 {
                            case 128:
                                if true {
                                    v1 += 42
                                } else {
                                    v7 += 42
                                }
                            case 0:
                                if false {
                                    fmt.Println(v7)
                                }
                            }
                        default:
                            var v66 bool = (!true)
                            _ = v66
                        }
                        v67 := []int{0, -2147483648, 9223372036854775807}
                        for fi68 := 0; fi68 < 10; fi68++ {
                            v69 := GenS5{F0: 0, F1: v7, F2: (v54 >> ((18446744073709551615) & 63))}
                            v3 = (^(v3))
                            _ = fi68
                            _ = v69
                        }
                        _ = v67
                    }
                    v1 -= 127
                    v70 := GenS5{F0: -2147483648, F1: v7, F2: (v28 >> ((255) & 63))}
                    _ = v70
                }
            }
            _ = v54
        } else {
            v71 := GenS3{F0: "", F1: 2147483648}
            v1 += v1
            _ = v71
        }
        if true {
            v7 += (^(v1))
            v72 := []int{32768, 32768}
            if ((2 >= 0) && true) {
                var v73 bool = true
                _ = v73
            }
            _ = v72
        } else {
            var v74 int64 = (v3 >> ((0) & 63))
            _ = v74
        }
        if (!((true && (false || true)) || false)) {
            fmt.Println(v28)
            v75 := GenS1{F0: -195, F1: 0, F2: -539}
            _ = v75
        }
    }
    fmt.Println(v28, v28, v28)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
