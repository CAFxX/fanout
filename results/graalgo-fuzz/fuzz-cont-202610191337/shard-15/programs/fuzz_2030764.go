package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64; F1 uint; F2 uint8 }
func genFn0() uint8 {
    var v6 int32 = 696
    if false {
        v6 = (^(v6))
        for fi7 := 0; fi7 < 15; fi7++ {
            var v8 float64 = ((-2.25 / float64(-2.25)) * (((-2.25 / float64(0.1)) - -2.25) - 0.1))
            _ = fi7
            _ = v8
        }
        v6 += v6
    } else {
        if (65535 < 831) {
            var v9 int8 = -128
            var v10 int16 = 1000
            _ = v10
            _ = v9
        } else {
            var v11 int32 = v6
            v6 = v11
            switch 4294967295 {
            case 2147483647:
                v12 := []int{-493, 32767, 32767, 1}
                _ = v12
            case -2147483648:
                var v13 int = 128
                _ = v13
            case 874:
                v11 -= 0
            default:
                var v14 float32 = 0.0
                _ = v14
            }
            _ = v11
        }
        v15 := GenS1{F0: 97}
        for fi16 := 0; fi16 < 15; fi16++ {
            var v17 string = "hello"
            if ((255 == 2) || (1 > 1)) {
                var v18 float64 = 1e10
                switch "\n\t" {
                case "":
                    var v19 int16 = 1000
                    _ = v19
                case "hello":
                    if (0 > 255) {
                        v20 := []int{}
                        var v21 int32 = v6
                        v20 = append(v20, 127)
                        _ = v20
                        _ = v21
                    }
                case "foo bar":
                    v22 := []int{578}
                    _ = v22
                default:
                    var v23 uint32 = 869
                    _ = v23
                }
                _ = v18
            }
            _ = fi16
            _ = v17
        }
        _ = v15
    }
    if (42 <= 0) {
        fmt.Println(v6)
        var v24 uint16 = 65535
        var v25 string = "hello"
        _ = v24
        _ = v25
    } else {
        v6 = 854
    }
    v6 = 694
    var v26 int64 = 1
    _ = v26
    _ = v6
    return 0
}
func genFn1(p0 uint64, p1 int16, p2 uint8) uint {
    if false {
        v27 := GenS1{F0: genFn0()}
        _ = v27
    }
    var v28 uint32 = 232
    _ = p0
    _ = p1
    _ = p2
    _ = v28
    return 2
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int8 = -1
    v3 := make(map[int]uint32)
    var v4 float64 = 3.14159
    var v5 float32 = 3.14159
    fmt.Println(v4, v2)
    v1 = (v1 ^ v1)
    v5 += v5
    if ((42 >= 2147483647) || (65535 >= 150)) {
        v2 += (v2 % ((v2) | 1))
        _ = len(v3)
    } else {
        var v29 float32 = 1e10
        if (((v2 & v2) > -84) && ((255 == 2) && (-646 > 0))) {
            for fi30 := 0; fi30 < 6; fi30++ {
                var v31 int16 = 32767
                _ = fi30
                _ = v31
            }
            v2 += (v2 - -128)
            switch "hello" {
            case "hello":
                v32 := GenS1{F0: 128}
                _ = v32
            case "a":
                v5 -= float32((v4 / ((1.5 + (0.1 - v4)) * 3.14159)))
            }
        }
        var v33 uint32 = 930
        _ = v29
        _ = v33
    }
    for fi34 := 0; fi34 < 19; fi34++ {
        switch "\n\t" {
        case "a":
            v35 := []int{(fi34 | (fi34 | fi34)), (fi34 >> ((42) & 63)), 1}
            _ = v35
        case "":
            var v36 uint32 = 42
            _ = v36
        case "hello":
            var v37 float32 = (v5 / (v5 / (v5 / v5)))
            _ = v37
        default:
            v5 += (v5 + v5)
        }
        var v38 uint32 = 42
        for fi39 := 0; fi39 < 10; fi39++ {
            if (("a" + "hello") <= "foo bar") {
                fmt.Println(fi34, fi39, v38)
            }
            v38 = v38
            v40 := v3[(fi34 ^ fi39)]
            _ = fi39
            _ = v40
        }
        _ = fi34
        _ = v38
    }
    fmt.Println(v5)
    v1 -= v1
    var v41 float32 = 0.0
    for fi42 := 0; fi42 < 4; fi42++ {
        v43 := []int{-869, (fi42 * (fi42 ^ fi42)), (fi42 | (fi42 ^ -9223372036854775808))}
        _ = fi42
        _ = v43
    }
    switch -9223372036854775808 {
    case 127:
        for fi44 := 0; fi44 < 7; fi44++ {
            fmt.Println(v2, v2, v41)
            fmt.Println(v4, v1, v2)
            _ = fi44
        }
    case 0:
        for fi45 := 0; fi45 < 7; fi45++ {
            switch "foo bar" {
            case "foo bar":
                fmt.Println(v1)
            case "\n\t":
                v2 = -128
            case "hello":
                switch "" {
                case "":
                    v2 = v2
                }
            default:
                var v46 uint8 = 0
                _ = v46
            }
            var v47 uint8 = 0
            _ = fi45
            _ = v47
        }
    case 9223372036854775807:
        v1 = (v1 - 329)
    default:
        var v48 string = (("hello" + ("foo bar" + "\n\t")) + (("foo bar" + "\n\t") + ("foo bar" + ("a" + (("" + (("a" + ("hello" + ("a" + ("hello" + "foo bar")))) + "foo bar")) + "")))))
        _ = v48
    }
    for fi49 := 0; fi49 < 17; fi49++ {
        for fi50 := 0; fi50 < 5; fi50++ {
            v2 = v2
            switch "a" {
            case "hello":
                v51 := GenS2{F0: -872}
                _ = v51
            default:
                fmt.Println(v41)
            }
            var v52 int16 = 1000
            _ = fi50
            _ = v52
        }
        for fi53 := 0; fi53 < 8; fi53++ {
            v54 := GenS3{F0: (v2 | (v2 - v2))}
            _ = fi53
            _ = v54
        }
        _ = fi49
    }
    var v55 bool = (-1 < 1)
    for fi56 := 0; fi56 < 15; fi56++ {
        var v57 int16 = 0
        var v58 int = fi56
        _ = fi56
        _ = v57
        _ = v58
    }
    switch "hello" {
    case "a":
        var v59 uint32 = 530
        _ = v59
    case "hello":
        var v60 float32 = ((v5 * float32((v4 + (((float64((v41 * v5)) * 3.14159) + 0.1) - 0.0)))) * -2.25)
        _ = v60
    }
    v61 := GenS4{F0: (v1 << ((0) & 63)), F1: 65535, F2: 255}
    switch "" {
    case "foo bar":
        var v62 uint32 = 2147483647
        _ = v62
    case "a":
        switch 2147483647 {
        case 0:
            if (-51 >= 42) {
                if (0 > -27) {
                    v63 := GenS3{F0: (v2 % (((v2 & v2)) | 1))}
                    _ = v63
                }
            } else {
                v1 -= v1
                v64 := GenS3{F0: v2}
                for fi65 := 0; fi65 < 13; fi65++ {
                    var v66 string = "hello"
                    for fi67 := 0; fi67 < 8; fi67++ {
                        v2 = (v2 | v2)
                        var v68 string = "hello"
                        v68 = "\n\t"
                        var v69 uint8 = 134
                        var v70 string = v68
                        _ = fi67
                        _ = v68
                        _ = v69
                        _ = v70
                    }
                    _ = fi65
                    _ = v66
                }
                _ = v64
            }
        case 295:
            switch "hello" {
            case "\n\t":
                delete(v3, -255)
            case "foo bar":
                switch "\n\t" {
                case "foo bar":
                    fmt.Println(v41, v1)
                case "\n\t":
                    v3[-9223372036854775808] = 0
                }
            }
        }
    }
    for fi71 := 0; fi71 < 18; fi71++ {
        v2 = v2
        v5 -= (0.1 - 0.0)
        for fi72 := 0; fi72 < 11; fi72++ {
            var v73 uint16 = 65535
            v55 = true
            switch "a" {
            case "\n\t":
                var v74 bool = (!v55)
                _ = v74
            case "hello":
                var v75 int16 = -1
                _ = v75
            case "foo bar":
                var v76 uint32 = 151
                _ = v76
            default:
                var v77 uint8 = 186
                _ = v77
            }
            _ = fi72
            _ = v73
        }
        _ = fi71
    }
    v78 := []int{788}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("slicelen v78:", len(v78), v78)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v55
    _ = v61
    _ = v78
}
