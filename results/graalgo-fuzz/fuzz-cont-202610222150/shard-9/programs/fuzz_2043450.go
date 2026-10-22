package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 uint8 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 uint; F1 uint64; F2 uint8 }
type GenS5 struct { F0 uint16; F1 uint64; F2 uint64 }
type GenS6 struct { F0 uint; F1 uint8; F2 uint32 }
type GenS7 struct { F0 int16; F1 int64; F2 int32 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int64; F1 int16; F2 int64 }
type GenS10 struct { F0 int8; F1 int16 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 uint64; F1 int }
type GenS13 struct { F0 int32; F1 uint16 }
type GenS14 struct { F0 uint; F1 string }
type GenS15 struct { F0 int }

func main() {
    var v1 float32 = -2.25
    var v2 string = ("foo bar" + ("\n\t" + "hello"))
    v3 := []uint32{437, 719}
    var v4 int16 = 70
    var v5 int = 42
    var v6 uint8 = 42
    var v7 int32 = -2147483648
    v5 += v5
    v4 -= 32767
    v8 := GenS1{F0: 524, F1: 2147483647, F2: (v6 * (v6 % ((255) | 1)))}
    switch 0 {
    case 9223372036854775807:
        if false {
            if true {
                if (201 != 1) {
                    var v9 int16 = 32767
                    _ = v9
                } else {
                    for fi10 := 0; fi10 < 3; fi10++ {
                        if (!(1.5 != 1e10)) {
                            v11 := GenS2{F0: "foo bar", F1: v6}
                            switch 0 {
                            case -2147483648:
                                for fi12 := 0; fi12 < 14; fi12++ {
                                    var v13 uint16 = uint16(v6)
                                    if (!((v5 | v5) >= v5)) {
                                        var v14 int = fi10
                                        var v15 float32 = ((v1 + 0.0) * (float32(1e10) + (v1 / ((v1 + v1) * (v1 + 0.0)))))
                                        v1 -= v15
                                        _ = v14
                                        _ = v15
                                    } else {
                                        v16 := GenS3{F0: (v6 / ((v6) | 1)), F1: v6}
                                        _ = v16
                                    }
                                    switch ("\n\t" + "") {
                                    case "hello":
                                        v17 := GenS1{F0: 1, F1: 4294967295, F2: 90}
                                        _ = v17
                                    default:
                                        v18 := GenS2{F0: v2, F1: (v6 - 0)}
                                        _ = v18
                                    }
                                    _ = fi12
                                    _ = v13
                                }
                            case 9223372036854775807:
                                v1 = (v1 * ((float32((1e10 + (float64(v1) / -2.25))) * (1.5 * ((v1 / float32(-2.25)) - v1))) / -2.25))
                            }
                            _ = v11
                        }
                        _ = fi10
                    }
                    var v19 int64 = 321
                    var v20 int32 = (v7 >> ((60) & 63))
                    _ = v19
                    _ = v20
                }
                if (!(((^(v5)) > (v5 & (v5 >> ((255) & 63)))) && (false || (((v5 | 1) >= (-(v5))) && false)))) {
                    v1 = v1
                }
                var v21 bool = ((((v1 + 1e10) - (v1 + v1)) * 0.0) != float32(1.5))
                _ = v21
            }
            v22 := GenS4{F0: 4294967295, F1: 440, F2: (v6 ^ (v6 % ((v6) | 1)))}
            _ = v22
        } else {
            var v23 int16 = int16(v7)
            if len(v3) > 0 {
                v3[0] = 937
            }
            var v24 bool = false
            _ = v23
            _ = v24
        }
    case 2147483647:
        if len(v3) > 0 {
            v3[0] = 250
        }
    case 127:
        var v25 int = (v5 << ((42) & 63))
        _ = v25
    default:
        switch (v5 / (((v5 + (v5 - 128))) | 1)) {
        case -9223372036854775808:
            switch v5 {
            case -925:
                for fi26 := 0; fi26 < 17; fi26++ {
                    v4 = (^(v4))
                    v6 += 0
                    _ = fi26
                }
            case 32768:
                switch ("hello" + (v2 + ((("\n\t" + v2) + "foo bar") + "a"))) {
                case "":
                    if len(v3) > 0 {
                        v3[0] = 42
                    }
                case "a":
                    for fi27 := 0; fi27 < 9; fi27++ {
                        if true {
                            v7 -= v7
                            fmt.Println(v6)
                            if true {
                                switch 4294967295 {
                                case -409:
                                    for fi28 := 0; fi28 < 19; fi28++ {
                                        v6 = (^(v6))
                                        var v29 string = "a"
                                        _ = fi28
                                        _ = v29
                                    }
                                case 376:
                                    if true {
                                        switch (fi27 << ((255) & 63)) {
                                        case -445:
                                            v4 = (v4 & (-(v4)))
                                        case 32767:
                                            var v30 int64 = -882
                                            _ = v30
                                        case -9223372036854775808:
                                            for fi31 := 0; fi31 < 11; fi31++ {
                                                if len(v3) > 0 {
                                                    v3[0] = 103
                                                }
                                                _ = fi31
                                            }
                                        default:
                                            v6 = v6
                                        }
                                    }
                                }
                            } else {
                                for fi32 := 0; fi32 < 18; fi32++ {
                                    switch ("hello" + v2) {
                                    case "hello":
                                        v33 := GenS4{F0: 2, F1: 0, F2: (v6 | 128)}
                                        _ = v33
                                    case "":
                                        var v34 int64 = -1
                                        _ = v34
                                    default:
                                        v6 -= v6
                                    }
                                    switch (fi32 / (((-(v5))) | 1)) {
                                    case 128:
                                        var v35 uint64 = 1
                                        _ = v35
                                    case 1:
                                        switch v2 {
                                        case "":
                                            var v36 int8 = 0
                                            if len(v3) > 0 {
                                                v3[0] = 462
                                            }
                                            if true {
                                                v37 := GenS4{F0: 400, F1: 576, F2: v6}
                                                _ = v37
                                            }
                                            _ = v36
                                        case "foo bar":
                                            var v38 float64 = 0.0
                                            _ = v38
                                        case "\n\t":
                                            var v39 int = fi27
                                            _ = v39
                                        }
                                    default:
                                        v40 := GenS5{F0: 2, F1: 1, F2: 1}
                                        switch (fi32 / (((^(v5))) | 1)) {
                                        case 75:
                                            var v41 uint32 = 0
                                            _ = v41
                                        case 32768:
                                            _ = len(v3)
                                        case 32767:
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        default:
                                            v5 = fi27
                                        }
                                        _ = len(v3)
                                        var v42 int32 = (v7 / (((v7 >> ((2) & 63))) | 1))
                                        _ = v40
                                        _ = v42
                                    }
                                    fmt.Println(v5, v5, v6)
                                    v2 = ((("\n\t" + v2) + v2) + (((v2 + "a") + "foo bar") + v2))
                                    fmt.Println(v6, v5)
                                    _ = fi32
                                }
                                v5 -= int(v4)
                            }
                        } else {
                            v2 = v2
                        }
                        fmt.Println(v7, v2)
                        fmt.Println(v4)
                        _ = fi27
                    }
                }
            case -416:
                v43 := GenS1{F0: 4294967295, F1: 4294967295, F2: 2}
                _ = v43
            }
        case 32767:
            v44 := GenS6{F0: 0, F1: v6, F2: 209}
            _ = v44
        case -934:
            v45 := GenS7{F0: (v4 << ((498) & 63)), F1: int64(v7), F2: (v7 / (((v7 ^ (v7 & v7))) | 1))}
            _ = v45
        default:
            for fi46 := 0; fi46 < 7; fi46++ {
                fmt.Println(v7, fi46)
                v3 = append(v3, 2147483647)
                v47 := GenS4{F0: 4294967295, F1: 471, F2: v6}
                _ = fi46
                _ = v47
            }
        }
    }
    v48 := GenS4{F0: 1, F1: 811, F2: 185}
    v49 := GenS8{F0: true}
    if (2147483647 != 239) {
        var v50 float64 = (0.1 * (1.5 * 1.5))
        _ = v50
    } else {
        for fi51 := 0; fi51 < 1; fi51++ {
            _ = len(v3)
            fmt.Println(fi51)
            _ = fi51
        }
        v3 = append(v3, 2147483647)
        v52 := GenS9{F0: 0, F1: v4, F2: 9223372036854775807}
        _ = v52
    }
    if (!true) {
        fmt.Println(v6)
    } else {
        if ((^(v6)) >= (v6 - v6)) {
            v4 = v4
            for fi53 := 0; fi53 < 2; fi53++ {
                if (18446744073709551615 <= 363) {
                    for fi54 := 0; fi54 < 2; fi54++ {
                        switch ((v2 + v2) + (v2 + ("hello" + "a"))) {
                        case "foo bar":
                            v3 = append(v3, 1)
                        case "":
                            if false {
                                var v55 uint8 = (v6 * (v6 & v6))
                                if (!((42 == (v55 << ((65535) & 63))) && false)) {
                                    if true {
                                        v55 = 128
                                        var v56 uint8 = v55
                                        v57 := GenS10{F0: 42, F1: v4}
                                        var v58 uint8 = v6
                                        v6 -= v55
                                        _ = v56
                                        _ = v57
                                        _ = v58
                                    }
                                    v55 = (v6 + (v55 / ((v55) | 1)))
                                }
                                v59 := GenS11{F0: false}
                                _ = v55
                                _ = v59
                            } else {
                                switch 9223372036854775807 {
                                case 4294967295:
                                    v4 = 539
                                }
                            }
                        case "hello":
                            if true {
                                v2 = v2
                                switch v2 {
                                case "a":
                                    var v60 uint32 = 0
                                    _ = v60
                                case "foo bar":
                                    var v61 uint = 42
                                    _ = v61
                                case "hello":
                                    if false {
                                        var v62 int16 = v4
                                        v62 -= v62
                                        _ = v62
                                    }
                                }
                            } else {
                                v4 += -1
                                v63 := GenS12{F0: 217, F1: 1}
                                v64 := GenS7{F0: v4, F1: 42, F2: (v7 | 2147483647)}
                                _ = v63
                                _ = v64
                            }
                        default:
                            if true {
                                if (!true) {
                                    switch "hello" {
                                    case "":
                                        var v65 uint8 = v6
                                        _ = v65
                                    }
                                } else {
                                    fmt.Println(fi54)
                                    var v66 uint32 = 2147483648
                                    if true {
                                        switch v2 {
                                        case "a":
                                            var v67 int = -721
                                            _ = v67
                                        default:
                                            if (!true) {
                                                var v68 int = (fi54 % ((v5) | 1))
                                                switch v68 {
                                                case 32768:
                                                    v3 = append(v3, v66)
                                                case 0:
                                                    if len(v3) > 0 {
                                                        v3[0] = 10
                                                    }
                                                }
                                                var v69 bool = true
                                                _ = v68
                                                _ = v69
                                            } else {
                                                if true {
                                                    var v70 uint8 = 87
                                                    if (false || true) {
                                                        v4 = (^(v4))
                                                    } else {
                                                        fmt.Println(fi54, v4)
                                                    }
                                                    _ = v70
                                                } else {
                                                    var v71 int = 128
                                                    v3 = append(v3, 1)
                                                    _ = v71
                                                }
                                            }
                                        }
                                        fmt.Println(fi53, v1)
                                    } else {
                                        var v72 uint = 4294967295
                                        v4 -= v4
                                        _ = v72
                                    }
                                    var v73 int = 4294967295
                                    _ = v66
                                    _ = v73
                                }
                                if len(v3) > 0 {
                                    v3[0] = 224
                                }
                                var v74 bool = true
                                _ = v74
                            } else {
                                switch 2147483647 {
                                case -1:
                                    v75 := GenS2{F0: (v2 + (((v2 + v2) + v2) + ((v2 + "foo bar") + (v2 + "hello")))), F1: 136}
                                    _ = v75
                                case -25:
                                    fmt.Println(fi54, v1)
                                    fmt.Println(fi54, v6, v2)
                                }
                            }
                        }
                        fmt.Println(v4, fi54, v5)
                        _ = fi54
                    }
                }
                _ = fi53
            }
            var v76 int16 = 1000
            _ = v76
        }
        var v77 bool = false
        _ = v77
    }
    switch (v5 * 0) {
    case 32767:
        v78 := GenS10{F0: 127, F1: v4}
        _ = v78
    }
    if true {
        if false {
            var v79 int32 = (v7 + (v7 + (v7 ^ v7)))
            var v80 uint32 = 2147483647
            _ = v79
            _ = v80
        }
        v2 = "\n\t"
        for fi81 := 0; fi81 < 1; fi81++ {
            for fi82 := 0; fi82 < 12; fi82++ {
                _ = len(v3)
                switch v5 {
                case 128:
                    v83 := GenS13{F0: v7, F1: 42}
                    _ = v83
                case 333:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case -501:
                    fmt.Println(v2, v4)
                default:
                    v1 = (v1 - ((v1 + v1) + (v1 / v1)))
                }
                v2 = (v2 + ("hello" + ((v2 + ((v2 + ((("foo bar" + v2) + "a") + v2)) + ("" + v2))) + v2)))
                _ = fi82
            }
            for fi84 := 0; fi84 < 17; fi84++ {
                v6 -= (v6 % ((v6) | 1))
                var v85 int32 = v7
                switch "foo bar" {
                case "a":
                    if ((!(false || (65535 != 27))) && (-1 <= 0)) {
                        if (-1 >= (v85 * 0)) {
                            if false {
                                var v86 float64 = ((1.5 + (0.0 - (0.1 / float64(v1)))) - -2.25)
                                v5 = fi81
                                _ = v86
                            } else {
                                if (-2.25 <= 0.0) {
                                    var v87 uint = 14
                                    v3 = append(v3, 4294967295)
                                    v4 = v4
                                    switch ("foo bar" + "") {
                                    case "\n\t":
                                        if (225 > -264) {
                                            v85 = int32(v87)
                                            v5 += fi81
                                            var v88 int32 = (v85 / (((v7 - 1)) | 1))
                                            var v89 uint32 = 6
                                            var v90 int64 = 0
                                            v91 := GenS14{F0: (v87 + v87), F1: "a"}
                                            _ = v88
                                            _ = v89
                                            _ = v90
                                            _ = v91
                                        } else {
                                            if len(v3) > 0 {
                                                v3[0] = 2
                                            }
                                            v4 += v4
                                            v92 := GenS15{F0: 0}
                                            _ = v92
                                        }
                                    default:
                                        fmt.Println(v87, v7, v87)
                                    }
                                    v6 = v6
                                    _ = len(v3)
                                    _ = v87
                                }
                                v1 -= v1
                            }
                        } else {
                            _ = len(v3)
                        }
                    }
                case "foo bar":
                    if true {
                        var v93 int8 = 100
                        var v94 int32 = (v7 + (v7 / (((v85 * -1)) | 1)))
                        _ = v93
                        _ = v94
                    }
                default:
                    var v95 uint16 = uint16(v6)
                    _ = v95
                }
                _ = fi84
                _ = v85
            }
            _ = fi81
        }
    }
    fmt.Println(v2, v6, v5)
    var v96 int16 = -615
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v96
}
