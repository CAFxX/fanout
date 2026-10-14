package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint64 }
type GenS2 struct { F0 int32; F1 uint16; F2 uint }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 uint8; F1 uint32; F2 string }
type GenS5 struct { F0 int32; F1 int8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint16; F1 uint8 }

func main() {
    var v1 bool = (106 > 124)
    var v2 float64 = (float64(float32(float64(-2.25))) + 1.5)
    var v3 int32 = 1
    for fi4 := 0; fi4 < 7; fi4++ {
        if ((v3 * v3) >= (^(v3))) {
            var v5 uint8 = 127
            var v6 uint16 = 42
            v7 := GenS1{F0: 42, F1: 42, F2: 18446744073709551615}
            _ = v5
            _ = v6
            _ = v7
        } else {
            v8 := GenS1{F0: 65535, F1: 183, F2: 579}
            if false {
                switch 128 {
                case -989:
                    v9 := GenS1{F0: 4294967295, F1: 0, F2: 781}
                    _ = v9
                }
                v10 := GenS1{F0: 10, F1: 0, F2: 42}
                if ((v3 & 607) > v3) {
                    for fi11 := 0; fi11 < 12; fi11++ {
                        switch fi4 {
                        case 4294967295:
                            v12 := []int{fi4, -2147483648}
                            _ = v12
                        case 81:
                            var v13 uint8 = 0
                            _ = v13
                        default:
                            switch "" {
                            case "hello":
                                var v14 float32 = ((-2.25 - (1.5 - 1.5)) + (0.0 / float32(0.1)))
                                _ = v14
                            case "\n\t":
                                v15 := []int{fi4, 1, 2147483647}
                                _ = v15
                            case "a":
                                v1 = true
                            }
                        }
                        var v16 int = fi11
                        _ = fi11
                        _ = v16
                    }
                } else {
                    if (0 != 700) {
                        for fi17 := 0; fi17 < 6; fi17++ {
                            v1 = ((v1 && v1) || (46 >= -9223372036854775808))
                            _ = fi17
                        }
                        v18 := []int{(fi4 * fi4)}
                        _ = v18
                    }
                }
                _ = v10
            } else {
                v19 := []int{fi4, (fi4 & 42), fi4}
                var v20 float64 = float64((0.1 - 3.14159))
                _ = v19
                _ = v20
            }
            _ = v8
        }
        v2 -= v2
        _ = fi4
    }
    for fi21 := 0; fi21 < 20; fi21++ {
        var v22 uint8 = 57
        v3 -= v3
        _ = fi21
        _ = v22
    }
    if true {
        if false {
            fmt.Println(v2)
            v23 := GenS1{F0: 65535, F1: 371, F2: 42}
            for fi24 := 0; fi24 < 5; fi24++ {
                v3 -= v3
                if v1 {
                    if ((0 >= 1) || ((v1 && v1) || (!v1))) {
                        for fi25 := 0; fi25 < 4; fi25++ {
                            v26 := GenS2{F0: v3, F1: 2, F2: 255}
                            v2 -= 1.5
                            v27 := GenS3{F0: 139, F1: 1}
                            fmt.Println(v3, v1, fi25)
                            v3 += (v3 + (v3 * v3))
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                    }
                }
                var v28 int64 = 42
                _ = fi24
                _ = v28
            }
            _ = v23
        } else {
            for fi29 := 0; fi29 < 18; fi29++ {
                fmt.Println(fi29, v2)
                _ = fi29
            }
        }
        if true {
            if v1 {
                v2 = 0.1
                v2 = (3.14159 - v2)
                v3 = v3
            }
            if (!true) {
                v30 := GenS4{F0: 127, F1: 0, F2: "a"}
                _ = v30
            } else {
                var v31 int16 = 42
                for fi32 := 0; fi32 < 1; fi32++ {
                    for fi33 := 0; fi33 < 1; fi33++ {
                        v34 := GenS5{F0: (v3 + 935), F1: 0}
                        v35 := GenS5{F0: 2147483647, F1: 1}
                        if (!((v2 > 0.1) || v1)) {
                            var v36 int32 = v3
                            _ = v36
                        } else {
                            v37 := []int{}
                            v1 = v1
                            _ = v37
                        }
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                    v3 += 915
                    _ = fi32
                }
                _ = v31
            }
            v38 := []int{42, 42, 4294967295}
            _ = v38
        }
    }
    switch "a" {
    case "foo bar":
        if true {
            v39 := GenS6{F0: 255}
            _ = v39
        }
    }
    var v40 uint16 = 89
    v2 = (v2 + (1.5 - ((-2.25 - ((v2 * 3.14159) * ((float64((0.0 / 1e10)) - (((3.14159 - float64((1e10 - (1e10 - float32(v2))))) * (float64(1e10) * v2)) - 1.5)) + (0.0 / v2)))) - 1.5)))
    v41 := GenS1{F0: 717, F1: 508, F2: 18446744073709551615}
    v40 = (v40 + (v40 << ((65535) & 63)))
    var v42 string = "foo bar"
    v43 := []int{0, 127}
    switch 32768 {
    case 32767:
        switch v42 {
        case "a":
            var v44 int16 = -1
            _ = v44
        case "\n\t":
            if v1 {
                var v45 int32 = -895
                _ = v45
            }
        default:
            v3 += v3
        }
    default:
        switch 127 {
        case 653:
            if true {
                var v46 float32 = 0.1
                _ = v46
            } else {
                switch 2147483647 {
                case 829:
                    switch "foo bar" {
                    case "\n\t":
                        var v47 uint = 18446744073709551615
                        _ = v47
                    case "a":
                        v3 += v3
                    default:
                        if (1000 > -58) {
                            var v48 uint16 = (v40 / ((255) | 1))
                            v43 = append(v43, -147)
                            _ = v48
                        } else {
                            v42 = ""
                        }
                    }
                }
                switch -117 {
                case 32767:
                    var v49 int = -9223372036854775808
                    _ = v49
                case 127:
                    switch ((v42 + "foo bar") + v42) {
                    case "":
                        for fi50 := 0; fi50 < 6; fi50++ {
                            v51 := GenS7{F0: 1}
                            switch fi50 {
                            case 1:
                                switch "hello" {
                                case "a":
                                    var v52 int16 = 1
                                    _ = v52
                                case "\n\t":
                                    v53 := GenS5{F0: 301, F1: 1}
                                    _ = v53
                                case "hello":
                                    var v54 float64 = v2
                                    _ = v54
                                default:
                                    var v55 uint16 = (^(v40))
                                    _ = v55
                                }
                            case -469:
                                v43 = append(v43, -115)
                            }
                            v2 -= v2
                            _ = fi50
                            _ = v51
                        }
                    case "a":
                        for fi56 := 0; fi56 < 14; fi56++ {
                            v40 += (v40 << ((4294967295) & 63))
                            var v57 int8 = 127
                            v3 += v3
                            _ = fi56
                            _ = v57
                        }
                    case "\n\t":
                        if len(v43) >= 2 {
                            v43 = v43[1:len(v43)]
                        }
                    }
                case -2147483648:
                    var v58 int8 = -128
                    _ = v58
                default:
                    if v1 {
                        if ((-2.25 == 3.14159) || v1) {
                            if len(v43) > 0 {
                                v43[0] = 1
                            }
                            _ = len(v43)
                            switch -935 {
                            case 364:
                                if v1 {
                                    fmt.Println(v1, v3)
                                } else {
                                    if ((!v1) && v1) {
                                        var v59 uint64 = 2
                                        v60 := GenS8{F0: (v40 | 680), F1: 128}
                                        _ = v59
                                        _ = v60
                                    } else {
                                        if (47 <= 383) {
                                            for fi61 := 0; fi61 < 2; fi61++ {
                                                for fi62 := 0; fi62 < 11; fi62++ {
                                                    v43 = append(v43, fi61)
                                                    _ = fi62
                                                }
                                                _ = fi61
                                            }
                                            v40 -= v40
                                            for fi63 := 0; fi63 < 4; fi63++ {
                                                fmt.Println(fi63, v3, v3)
                                                v3 -= (v3 / ((v3) | 1))
                                                v64 := GenS7{F0: -79}
                                                _ = fi63
                                                _ = v64
                                            }
                                        }
                                        var v65 string = v42
                                        _ = v65
                                    }
                                    v42 = (v42 + v42)
                                }
                            case 32768:
                                fmt.Println(v2, v40)
                            case 32767:
                                _ = len(v43)
                            default:
                                fmt.Println(v3, v40)
                            }
                        }
                    } else {
                        v66 := GenS7{F0: 1}
                        _ = v66
                    }
                }
                v40 = 65535
            }
        case 42:
            var v67 int16 = 0
            _ = v67
        case 2147483647:
            if (v1 || (42 <= 0)) {
                if len(v43) > 0 {
                    v43[0] = 32768
                }
            } else {
                fmt.Println(v1)
                var v68 uint8 = 0
                for fi69 := 0; fi69 < 17; fi69++ {
                    fmt.Println(v3, v2, fi69)
                    if (504 <= 1) {
                        v42 = v42
                        var v70 string = ""
                        if (((false && (!((1 > 42) && true))) && (fi69 == fi69)) && v1) {
                            v70 = v70
                        } else {
                            v2 += v2
                            v70 = v70
                        }
                        _ = v70
                    }
                    _ = fi69
                }
                _ = v68
            }
        default:
            v42 = v42
        }
    }
    if len(v43) >= 2 {
        v43 = v43[1:len(v43)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v43:", len(v43), v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v40
    _ = v41
    _ = v42
    _ = v43
}
