package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int; F1 int32; F2 int32 }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 int8; F1 int32; F2 int }
type GenS5 struct { F0 int64; F1 uint64 }
type GenS6 struct { F0 uint; F1 int64 }

func main() {
    var v1 string = "\n\t"
    var v2 uint16 = 65535
    var v3 int = 64
    var v4 uint = 4294967295
    var v5 int = v3
    var v6 uint = 1
    var v7 uint = v6
    if true {
        v8 := GenS1{F0: (v2 / ((v2) | 1)), F1: 454}
        _ = v8
    } else {
        switch (v5 & (^(v5))) {
        case -878:
            v7 -= (v7 << (((^(v4))) & 63))
        default:
            for fi9 := 0; fi9 < 20; fi9++ {
                v4 = (v6 | v4)
                v10 := GenS2{F0: v5, F1: 1, F2: 415}
                _ = fi9
                _ = v10
            }
        }
        for fi11 := 0; fi11 < 10; fi11++ {
            v12 := []int{336, v5}
            _ = fi11
            _ = v12
        }
        var v13 bool = false
        _ = v13
    }
    for fi14 := 0; fi14 < 5; fi14++ {
        v15 := []int{794}
        _ = fi14
        _ = v15
    }
    if true {
        var v16 bool = (-816 != 247)
        v17 := GenS2{F0: 310, F1: -1, F2: 2147483647}
        if false {
            v16 = v16
        } else {
            v18 := GenS2{F0: v5, F1: 42, F2: 42}
            if (!true) {
                v19 := []int{(v5 % ((9223372036854775807) | 1)), v5}
                v3 -= -1
                _ = v19
            }
            _ = v18
        }
        _ = v16
        _ = v17
    }
    var v20 string = ("hello" + v1)
    var v21 float64 = 0.1
    if true {
        var v22 int64 = 144
        if (!(!(!(!((true || false) && (1 == v4)))))) {
            var v23 int8 = 16
            var v24 string = "a"
            if false {
                fmt.Println(v22, v22)
                v1 = v1
            } else {
                v25 := GenS3{F0: 604, F1: int(v22)}
                _ = v25
            }
            _ = v23
            _ = v24
        }
        switch 0 {
        case -9223372036854775808:
            for fi26 := 0; fi26 < 9; fi26++ {
                v27 := []int{(^(fi26)), v3}
                var v28 string = "foo bar"
                _ = fi26
                _ = v27
                _ = v28
            }
        case 4294967295:
            v3 = (v5 / (((v5 << ((42) & 63))) | 1))
        case -170:
            if (42 > 488) {
                v29 := []int{v5, 9223372036854775807}
                var v30 int32 = 344
                v31 := GenS2{F0: int(v22), F1: v30, F2: 0}
                _ = v29
                _ = v30
                _ = v31
            } else {
                switch v20 {
                case "hello":
                    v32 := GenS4{F0: 100, F1: -2147483648, F2: v5}
                    _ = v32
                case "a":
                    for fi33 := 0; fi33 < 9; fi33++ {
                        switch (-(v5)) {
                        case -169:
                            v6 += v6
                        case -668:
                            v34 := GenS1{F0: v2, F1: -650}
                            _ = v34
                        case 364:
                            var v35 string = ((v1 + v1) + "foo bar")
                            _ = v35
                        }
                        v22 = -1
                        _ = fi33
                    }
                }
                v36 := GenS3{F0: -699, F1: -427}
                v4 -= (v6 + (^(v7)))
                _ = v36
            }
        }
        _ = v22
    } else {
        for fi37 := 0; fi37 < 14; fi37++ {
            var v38 uint64 = 4294967295
            v39 := []int{v3}
            for fi40 := 0; fi40 < 18; fi40++ {
                if (71 >= -83) {
                    switch (v1 + v20) {
                    case "a":
                        var v41 int8 = -89
                        _ = v41
                    }
                    v39 = append(v39, (^(fi40)))
                }
                _ = fi40
            }
            _ = fi37
            _ = v38
            _ = v39
        }
    }
    v42 := GenS3{F0: 2147483647, F1: 940}
    if (-2.25 != -2.25) {
        if (545 >= 0) {
            for fi43 := 0; fi43 < 15; fi43++ {
                for fi44 := 0; fi44 < 14; fi44++ {
                    var v45 bool = true
                    _ = fi44
                    _ = v45
                }
                fmt.Println(v3, v7, v7)
                v46 := GenS3{F0: -2147483648, F1: (-(v3))}
                _ = fi43
                _ = v46
            }
            v47 := []int{(v5 - -9223372036854775808)}
            v1 = v20
            _ = v47
        }
        v48 := []int{v5}
        _ = v48
    }
    if true {
        v49 := []int{(v3 - (v5 - -9223372036854775808))}
        if true {
            var v50 int16 = 1
            v4 = v4
            fmt.Println(v1, v7, v1)
            _ = v50
        }
        for fi51 := 0; fi51 < 11; fi51++ {
            if true {
                switch v1 {
                case "hello":
                    if true {
                        switch ("hello" + v20) {
                        case "foo bar":
                            var v52 string = v20
                            _ = v52
                        case "":
                            v3 += (v3 + (fi51 & (v5 * -2147483648)))
                        }
                    } else {
                        v53 := GenS5{F0: 365, F1: 2}
                        _ = v53
                    }
                }
                v7 = (v4 >> (((v7 / ((v7) | 1))) & 63))
                v49 = append(v49, (^(v5)))
            } else {
                if true {
                    _ = len(v49)
                    switch (v5 / (((fi51 % ((-2147483648) | 1))) | 1)) {
                    case 2147483647:
                        v54 := GenS3{F0: -2147483648, F1: int(v2)}
                        _ = v54
                    case 380:
                        for fi55 := 0; fi55 < 13; fi55++ {
                            v1 = ("foo bar" + ("hello" + v20))
                            v49 = append(v49, fi51)
                            var v56 float64 = -2.25
                            v4 += v6
                            var v57 float64 = (((v21 * (1e10 / 1.0)) + (((v56 / ((1e10 * v21) / v56)) * 0.0) / v21)) - v56)
                            _ = fi55
                            _ = v56
                            _ = v57
                        }
                    case -1:
                        var v58 string = v20
                        _ = v58
                    default:
                        for fi59 := 0; fi59 < 1; fi59++ {
                            if false {
                                switch (v5 * fi59) {
                                case 32768:
                                    var v60 uint = v7
                                    _ = v60
                                case -9223372036854775808:
                                    v6 -= v6
                                }
                                _ = len(v49)
                                if true {
                                    if len(v49) > 0 {
                                        v49[0] = (v3 % (((fi59 % (((fi59 ^ fi51)) | 1))) | 1))
                                    }
                                } else {
                                    var v61 uint16 = (^(v2))
                                    v62 := GenS2{F0: (^(fi51)), F1: 2147483647, F2: 495}
                                    var v63 int64 = 673
                                    _ = v61
                                    _ = v62
                                    _ = v63
                                }
                            }
                            v5 -= fi59
                            v1 = ((v20 + "\n\t") + "")
                            _ = fi59
                        }
                    }
                    for fi64 := 0; fi64 < 20; fi64++ {
                        fmt.Println(fi51)
                        switch v1 {
                        case "foo bar":
                            switch v1 {
                            case "hello":
                                _ = len(v49)
                            default:
                                var v65 uint = (^(v7))
                                _ = v65
                            }
                        case "\n\t":
                            v5 = 32767
                            v6 = (^(v6))
                            v4 = 4294967295
                            if false {
                                v6 = 255
                                if false {
                                    var v66 string = v20
                                    var v67 int = fi64
                                    switch v66 {
                                    case "a":
                                        var v68 uint = (^(v7))
                                        _ = v68
                                    case "foo bar":
                                        if (true || false) {
                                            v69 := GenS4{F0: 100, F1: 2147483647, F2: v67}
                                            v7 = v6
                                            switch ("hello" + v66) {
                                            case "a":
                                                _ = len(v49)
                                            case "foo bar":
                                                v5 = -701
                                            case "":
                                                v67 = v5
                                            }
                                            if ((v7 != (v6 - (^(v4)))) && ((!true) && false)) {
                                                v70 := GenS6{F0: 255, F1: 1}
                                                fmt.Println(v5, v3, v4)
                                                if (true || (9223372036854775807 <= 208)) {
                                                    v5 -= fi64
                                                    v49 = append(v49, v5)
                                                } else {
                                                    v4 += v6
                                                }
                                                _ = v70
                                            } else {
                                                fmt.Println(fi51, v4)
                                                _ = len(v49)
                                                var v71 string = v66
                                                _ = v71
                                            }
                                            v66 = "a"
                                            _ = v69
                                        }
                                    case "":
                                        if (980 == 1000) {
                                            var v72 uint32 = 42
                                            v20 = v66
                                            if ((!(135 <= int32(fi64))) || false) {
                                                if len(v49) > 0 {
                                                    v49[0] = 2147483647
                                                }
                                                if (v6 != v6) {
                                                    var v73 uint32 = (^(v72))
                                                    _ = v73
                                                }
                                                v66 = ""
                                            } else {
                                                v74 := GenS2{F0: (-(fi51)), F1: -462, F2: 2147483647}
                                                v75 := GenS4{F0: int8(v72), F1: -1, F2: (v3 % ((v5) | 1))}
                                                _ = v74
                                                _ = v75
                                            }
                                            _ = v72
                                        }
                                    }
                                    _ = v66
                                    _ = v67
                                } else {
                                    var v76 float64 = (v21 * v21)
                                    _ = v76
                                }
                                var v77 uint8 = 254
                                _ = v77
                            }
                        default:
                            if true {
                                v49 = append(v49, (v3 & fi64))
                            } else {
                                v78 := GenS5{F0: 42, F1: 754}
                                if len(v49) > 0 {
                                    v49[0] = (^(fi64))
                                }
                                fmt.Println(v3, v2, v1)
                                v6 += 65535
                                _ = v78
                            }
                        }
                        v79 := GenS4{F0: 122, F1: 1, F2: (fi51 % ((v3) | 1))}
                        _ = fi64
                        _ = v79
                    }
                }
            }
            _ = fi51
        }
        _ = v49
    } else {
        var v80 uint8 = 2
        for fi81 := 0; fi81 < 7; fi81++ {
            fmt.Println(fi81, v7)
            _ = fi81
        }
        if (!(false || ((true && (1 >= 0)) || false))) {
            v21 -= v21
            switch ("foo bar" + (v1 + "hello")) {
            case "a":
                v82 := GenS1{F0: (^(v2)), F1: -592}
                _ = v82
            }
        } else {
            v7 -= (v6 - 2)
            v83 := []int{(v3 * 2147483647), v5, 128}
            _ = v83
        }
        _ = v80
    }
    v2 = (v2 | (v2 % (((^(v2))) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
    _ = v7
}
