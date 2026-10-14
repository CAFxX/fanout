package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 uint; F1 int; F2 uint8 }
type GenS4 struct { F0 bool; F1 int16 }
type GenS5 struct { F0 uint32; F1 bool }
type GenS6 struct { F0 string; F1 uint16 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 string; F1 int64 }
type GenS9 struct { F0 int16; F1 int32; F2 int16 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int; F1 uint }
func genFn0() uint64 {
    if false {
        v6 := GenS1{F0: 42}
        _ = v6
    } else {
        v7 := []int{-1, -2147483648, 4294967295, -9223372036854775808}
        _ = v7
    }
    var v8 int16 = -343
    var v9 int16 = (v8 >> ((59) & 63))
    for fi10 := 0; fi10 < 12; fi10++ {
        v11 := GenS2{F0: 65535, F1: -729}
        v12 := GenS3{F0: 2, F1: -407, F2: 169}
        var v13 uint8 = 42
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    v9 = 905
    _ = v8
    _ = v9
    return 18446744073709551615
}
func genFn1() uint64 {
    switch 362 {
    case 0:
        for fi14 := 0; fi14 < 3; fi14++ {
            var v15 int8 = 42
            v16 := []int{(fi14 - fi14), fi14}
            for fi17 := 0; fi17 < 2; fi17++ {
                v15 += (^(v15))
                v15 -= 0
                var v18 int = fi17
                _ = fi17
                _ = v18
            }
            _ = fi14
            _ = v15
            _ = v16
        }
    default:
        switch -40 {
        case 2147483647:
            switch -2147483648 {
            case 927:
                if true {
                    var v19 uint = 4294967295
                    v20 := []int{-2147483648, 32768}
                    _ = v19
                    _ = v20
                }
            }
        case 127:
            var v21 float64 = 0.1
            _ = v21
        }
    }
    v22 := []int{-670, 42, 4294967295}
    _ = v22
    return 18446744073709551615
}

func main() {
    v1 := []uint32{4294967295}
    var v2 bool = false
    var v3 int = 4294967295
    var v4 float64 = 0.0
    v5 := []uint32{65535}
    for fi23 := 0; fi23 < 4; fi23++ {
        switch "\n\t" {
        case "foo bar":
            var v24 uint64 = 1
            _ = v24
        case "a":
            for fi25 := 0; fi25 < 11; fi25++ {
                v5 = append(v5, 0)
                _ = fi25
            }
        case "hello":
            if (-1 == -2147483648) {
                _ = len(v1)
                if (!true) {
                    switch "hello" {
                    case "":
                        v5 = append(v5, 42)
                    case "a":
                        v5 = append(v5, 2147483648)
                    default:
                        switch "" {
                        case "":
                            for fi26 := 0; fi26 < 17; fi26++ {
                                if (68 < 42) {
                                    if v2 {
                                        v2 = v2
                                        v27 := GenS1{F0: 71}
                                        _ = v27
                                    } else {
                                        fmt.Println(v4)
                                    }
                                    if (v2 || v2) {
                                        fmt.Println(v2, v4)
                                        if (v2 || v2) {
                                            v1 = append(v1, 2147483648)
                                            v28 := GenS4{F0: (v2 || v2), F1: 1}
                                            _ = v28
                                        }
                                        var v29 bool = v2
                                        if (0 >= -13) {
                                            fmt.Println(v4)
                                            if false {
                                                v1 = append(v1, 1)
                                            }
                                            if v29 {
                                                v30 := GenS5{F0: 2147483648, F1: (!(("" + "a") < "foo bar"))}
                                                _ = v30
                                            } else {
                                                v31 := GenS1{F0: 200}
                                                v32 := GenS2{F0: 42, F1: 2147483647}
                                                switch "foo bar" {
                                                case "a":
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                case "\n\t":
                                                    var v33 uint = 42
                                                    _ = v33
                                                case "hello":
                                                    var v34 float64 = (((v4 + v4) + v4) + -2.25)
                                                    _ = v34
                                                }
                                                v2 = v29
                                                _ = v31
                                                _ = v32
                                            }
                                        } else {
                                            v29 = (true || ((!(255 <= 42)) || v2))
                                            var v35 bool = ((!v2) || (v29 && (!(!(-48 > 1)))))
                                            _ = v35
                                        }
                                        _ = v29
                                    }
                                    v36 := GenS4{F0: v2, F1: -32768}
                                    _ = v36
                                } else {
                                    fmt.Println(fi23, fi26)
                                }
                                v37 := GenS6{F0: "", F1: 255}
                                _ = fi26
                                _ = v37
                            }
                        default:
                            fmt.Println(fi23, v2)
                        }
                    }
                }
            }
        }
        v3 = (^(fi23))
        _ = fi23
    }
    v38 := GenS7{F0: 747}
    v3 = 252
    v39 := GenS8{F0: "", F1: 9223372036854775807}
    for fi40 := 0; fi40 < 1; fi40++ {
        v41 := GenS5{F0: 4294967295, F1: false}
        fmt.Println(v4)
        _ = fi40
        _ = v41
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v42 uint16 = 2
    if v2 {
        v3 = 32767
        var v43 uint = 42
        _ = v43
    } else {
        v44 := GenS9{F0: 1000, F1: 2147483647, F2: 484}
        v42 = 255
        _ = v44
    }
    if (86 < 42) {
        for fi45 := 0; fi45 < 10; fi45++ {
            for fi46 := 0; fi46 < 13; fi46++ {
                fmt.Println(v2, v2, v2)
                fmt.Println(v2, v4)
                v4 += ((v4 - v4) * 1.5)
                _ = fi46
            }
            _ = fi45
        }
    }
    for fi47 := 0; fi47 < 7; fi47++ {
        fmt.Println(v3, v42)
        _ = fi47
    }
    if (!(-587 >= -1)) {
        var v48 int32 = -2147483648
        for fi49 := 0; fi49 < 19; fi49++ {
            if (0 <= 42) {
                v50 := GenS10{F0: ((4294967295 > 255) || (!false))}
                var v51 uint32 = 756
                v2 = v2
                _ = v50
                _ = v51
            } else {
                for fi52 := 0; fi52 < 9; fi52++ {
                    var v53 uint16 = (v42 + v42)
                    var v54 int = int(v48)
                    if false {
                        v55 := GenS2{F0: v42, F1: (v48 | -2147483648)}
                        if len(v1) > 0 {
                            v1[0] = 0
                        }
                        _ = v55
                    } else {
                        var v56 int32 = v48
                        if (23 >= 0) {
                            fmt.Println(v54)
                            v1 = append(v1, 0)
                            var v57 int8 = 100
                            _ = v57
                        } else {
                            fmt.Println(fi49, v48, fi52)
                            var v58 int32 = 1
                            v42 -= 32767
                            _ = v58
                        }
                        switch "\n\t" {
                        case "":
                            v59 := GenS8{F0: ("\n\t" + "hello"), F1: -631}
                            _ = v59
                        default:
                            v56 -= v56
                        }
                        _ = v56
                    }
                    _ = fi52
                    _ = v53
                    _ = v54
                }
                fmt.Println(v2)
                v42 = (v42 + 0)
            }
            fmt.Println(v42, v3, v4)
            var v60 uint16 = v42
            _ = fi49
            _ = v60
        }
        _ = v48
    } else {
        var v61 float64 = v4
        v3 += (v3 / ((v3) | 1))
        _ = v61
    }
    for fi62 := 0; fi62 < 15; fi62++ {
        var v63 uint16 = v42
        if v2 {
            v42 -= (v42 / ((v63) | 1))
            v5 = append(v5, 2147483648)
        }
        _ = fi62
        _ = v63
    }
    if (1 != uint64(v42)) {
        v64 := GenS11{F0: (v3 & (v3 | v3)), F1: 42}
        _ = v64
    } else {
        var v65 uint16 = (v42 % (((v42 / (((^(v42))) | 1))) | 1))
        _ = v65
    }
    v4 = (v4 + (((1e10 - v4) / (v4 + v4)) + float64((3.14159 - 0.1))))
    var v66 int16 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v42
    _ = v5
    _ = v66
}
