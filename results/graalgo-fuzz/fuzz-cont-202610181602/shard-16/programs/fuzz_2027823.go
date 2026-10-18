package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 int16; F1 bool; F2 uint }
type GenS4 struct { F0 uint64; F1 int64; F2 int }
type GenS5 struct { F0 uint; F1 int; F2 int }
type GenS6 struct { F0 uint64; F1 int16 }
type GenS7 struct { F0 bool; F1 int16 }
type GenS8 struct { F0 string; F1 uint32; F2 uint8 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int32 = -1
    var v3 bool = (true || false)
    var v4 uint32 = 42
    var v5 uint8 = 112
    var v6 uint32 = (^(v4))
    v3 = (!(803 > 65535))
    delete(v1, (v4 & 553))
    fmt.Println(v3, v6, v3)
    var v7 uint64 = 2
    if (672 < 4294967295) {
        var v8 bool = true
        _ = v8
    } else {
        v9 := v1[v6]
        _ = v9
    }
    var v10 bool = v3
    var v11 uint16 = 65535
    var v12 int16 = 0
    v5 = (^(v5))
    for fi13 := 0; fi13 < 13; fi13++ {
        v7 += (^(v7))
        _ = fi13
    }
    fmt.Println(v7, v12)
    for fi14 := 0; fi14 < 12; fi14++ {
        v12 -= v12
        if v10 {
            if v3 {
                if (v2 > 42) {
                    switch "foo bar" {
                    case "foo bar":
                        v12 = (v12 ^ v12)
                    case "a":
                        var v15 int16 = 0
                        _ = v15
                    default:
                        v16 := []int{}
                        _ = v16
                    }
                } else {
                    v17 := []int{(fi14 << ((255) & 63)), (fi14 * (fi14 % ((fi14) | 1))), (-(fi14)), (fi14 & fi14)}
                    v18 := GenS1{F0: (v12 | (-(v12)))}
                    fmt.Println(v4)
                    _ = v17
                    _ = v18
                }
                var v19 uint16 = 2
                _ = v19
            }
        } else {
            v5 = (v5 | uint8(v6))
            for fi20 := 0; fi20 < 2; fi20++ {
                fmt.Println(v10, v5, v2)
                var v21 uint64 = v7
                if false {
                    v22 := GenS1{F0: v12}
                    var v23 uint32 = (v4 << ((4294967295) & 63))
                    v24 := GenS1{F0: v12}
                    _ = v22
                    _ = v23
                    _ = v24
                } else {
                    fmt.Println(fi20, v10, v12)
                    if false {
                        v25 := GenS2{F0: ("a" + "hello"), F1: ("\n\t" < "a")}
                        v21 = uint64(v5)
                        _ = v25
                    } else {
                        var v26 int = fi14
                        var v27 uint16 = 255
                        _ = v26
                        _ = v27
                    }
                }
                switch (^(fi14)) {
                case 42:
                    v28 := GenS1{F0: v12}
                    _ = v28
                }
                fmt.Println(v10, v4, fi14)
                _ = fi20
                _ = v21
            }
        }
        var v29 uint = 22
        _ = fi14
        _ = v29
    }
    for fi30 := 0; fi30 < 10; fi30++ {
        v31 := GenS3{F0: v12, F1: false, F2: 255}
        v2 -= (v2 - v2)
        switch "" {
        case "a":
            if v10 {
                for fi32 := 0; fi32 < 15; fi32++ {
                    v10 = true
                    var v33 bool = (v10 || (fi30 < int(v5)))
                    var v34 int = 32767
                    _ = fi32
                    _ = v33
                    _ = v34
                }
            }
        case "hello":
            if v10 {
                switch "" {
                case "\n\t":
                    var v35 string = "hello"
                    _ = v35
                case "a":
                    for fi36 := 0; fi36 < 18; fi36++ {
                        var v37 uint8 = v5
                        if ((v3 && (v6 > v4)) && (!false)) {
                            if true {
                                if false {
                                    switch ("" + "hello") {
                                    case "a":
                                        switch 0 {
                                        case -2147483648:
                                            fmt.Println(v6, v7)
                                        case -125:
                                            v38 := GenS1{F0: -81}
                                            _ = v38
                                        default:
                                            var v39 float64 = (1e10 - (((0.1 * 0.1) + (0.1 / (((1.5 / 1.0) + float64(-2.25)) / (((3.14159 / (1e10 * 1e10)) / 3.14159) + (float64(float32(-2.25)) + -2.25))))) - 1.5))
                                            _ = v39
                                        }
                                    default:
                                        var v40 float64 = ((1e10 - 3.14159) - 1.5)
                                        _ = v40
                                    }
                                } else {
                                    delete(v1, (v4 - uint32(v5)))
                                    switch 127 {
                                    case 128:
                                        v37 = v5
                                        v12 = v12
                                        var v41 int32 = (^(v2))
                                        _ = v41
                                    case -935:
                                        _ = len(v1)
                                    case -745:
                                        switch "hello" {
                                        case "":
                                            switch ("a" + "hello") {
                                            case "foo bar":
                                                var v42 int = (fi36 + (fi36 - 325))
                                                _ = v42
                                            case "a":
                                                if ((int(v37) == 0) && v10) {
                                                    v37 = (v37 * (v37 / (((v37 / (((v37 & (^(v5)))) | 1))) | 1)))
                                                    var v43 uint32 = uint32(v37)
                                                    _ = v43
                                                } else {
                                                    v44 := []int{fi36, fi36}
                                                    var v45 float32 = (3.14159 + -2.25)
                                                    v46 := GenS4{F0: (v7 / (((^(v7))) | 1)), F1: -1, F2: fi36}
                                                    _ = v44
                                                    _ = v45
                                                    _ = v46
                                                }
                                            default:
                                                if ((v4 << ((1) & 63)) > (v6 / ((v6) | 1))) {
                                                    if v10 {
                                                        v2 += 1
                                                    }
                                                    v2 = (v2 * -2147483648)
                                                    if true {
                                                        var v47 int64 = -9223372036854775808
                                                        if (42 >= 80) {
                                                            var v48 int16 = 1000
                                                            _ = v48
                                                        }
                                                        _ = v47
                                                    } else {
                                                        fmt.Println(v10)
                                                        v5 = 255
                                                        v7 = v7
                                                    }
                                                }
                                            }
                                        case "foo bar":
                                            var v49 uint32 = 2147483648
                                            _ = v49
                                        case "\n\t":
                                            switch (fi30 * 4294967295) {
                                            case -9223372036854775808:
                                                var v50 int32 = v2
                                                _ = v50
                                            }
                                        }
                                    }
                                    var v51 uint32 = uint32(v7)
                                    v52 := GenS4{F0: uint64(v37), F1: 595, F2: fi30}
                                    var v53 float64 = ((float64(0.1) - (float64(0.0) * (((3.14159 + 0.0) * float64(-2.25)) / 1.5))) + (-2.25 * 0.0))
                                    switch "" {
                                    case "hello":
                                        v6 = 604
                                    case "\n\t":
                                        switch (fi36 - -266) {
                                        case 127:
                                            v3 = false
                                        case 32768:
                                            v5 += (^(v5))
                                        }
                                    }
                                    _ = v51
                                    _ = v52
                                    _ = v53
                                }
                                switch "foo bar" {
                                case "hello":
                                    v54 := GenS5{F0: 42, F1: (fi30 / (((fi30 >> ((42) & 63))) | 1)), F2: (fi36 % ((32768) | 1))}
                                    _ = v54
                                case "a":
                                    v55 := []int{(fi30 >> ((1) & 63)), (-(fi30)), (fi30 / ((9223372036854775807) | 1)), fi30}
                                    _ = v55
                                default:
                                    v10 = false
                                }
                                if ((^(v5)) > 127) {
                                    v10 = (v4 > (^(v6)))
                                } else {
                                    var v56 uint64 = v7
                                    v57 := GenS5{F0: 737, F1: (fi36 * (^(fi36))), F2: (fi36 - (^(fi36)))}
                                    var v58 bool = v10
                                    _ = v56
                                    _ = v57
                                    _ = v58
                                }
                                switch (("a" + "foo bar") + "\n\t") {
                                case "foo bar":
                                    var v59 uint16 = (^(v11))
                                    _ = v59
                                }
                                fmt.Println(v12, v12, v10)
                            } else {
                                v2 += int32(v7)
                                v5 = 128
                            }
                        }
                        if ((v3 && true) || v10) {
                            if v3 {
                                var v60 int32 = v2
                                _ = v60
                            }
                            v61 := []int{-9223372036854775808, (-(fi30)), (-(fi30)), fi36}
                            v37 += (v37 & (v37 << ((1) & 63)))
                            _ = v61
                        }
                        _ = fi36
                        _ = v37
                    }
                case "hello":
                    var v62 string = ("a" + "")
                    _ = v62
                }
                for fi63 := 0; fi63 < 8; fi63++ {
                    switch "\n\t" {
                    case "\n\t":
                        v4 -= v4
                    }
                    v64 := []int{}
                    var v65 float64 = -2.25
                    _ = fi63
                    _ = v64
                    _ = v65
                }
                v10 = ((1.5 > 1.5) && (v7 > uint64(v11)))
            }
        }
        _ = fi30
        _ = v31
    }
    v5 = 42
    fmt.Println(v3, v7, v10)
    var v66 bool = (((0.0 * (-2.25 * (3.14159 + 1.5))) * ((((float64(1e10) + (1e10 * 1.5)) + float64(0.0)) / 1e10) * 0.0)) <= 1.5)
    if v10 {
        var v67 int32 = (v2 & (v2 / ((v2) | 1)))
        fmt.Println(v2, v5, v5)
        for fi68 := 0; fi68 < 12; fi68++ {
            v69 := GenS6{F0: v7, F1: v12}
            var v70 uint32 = 2147483648
            _ = fi68
            _ = v69
            _ = v70
        }
        _ = v67
    } else {
        v71 := GenS7{F0: v3, F1: (v12 * (^(v12)))}
        v2 = (^(v2))
        v72 := GenS8{F0: "\n\t", F1: (^(v4)), F2: 194}
        _ = v71
        _ = v72
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v66:", v66)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v66
    _ = v7
}
