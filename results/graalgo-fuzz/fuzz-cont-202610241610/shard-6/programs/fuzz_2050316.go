package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 bool; F1 int16 }
type GenS5 struct { F0 uint16; F1 int16 }
type GenS6 struct { F0 int16; F1 bool; F2 int8 }
type GenS7 struct { F0 int16; F1 int32 }
func genFn0() string {
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 bool = (!false)
        v7 = (18446744073709551615 > 251)
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        for fi9 := 0; fi9 < 19; fi9++ {
            var v10 int8 = -114
            v10 = (-(v10))
            var v11 bool = ((((float64(1.5) / 1.0) + 0.1) + ((0.0 / (3.14159 / -2.25)) * 1e10)) >= float64((((3.14159 - 1.5) + 0.1) / float32(1e10))))
            _ = fi9
            _ = v10
            _ = v11
        }
        if (247 <= 42) {
            v12 := GenS1{F0: (float64(1.5) <= 0.1), F1: 127, F2: -85}
            var v13 int = fi8
            for fi14 := 0; fi14 < 4; fi14++ {
                fmt.Println(v13)
                v15 := GenS1{F0: (false || (0 <= 4294967295)), F1: 127, F2: 127}
                v16 := []int{(-(fi8)), fi8, fi8}
                v13 -= 128
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = v12
            _ = v13
        } else {
            var v17 int8 = 0
            var v18 uint8 = 128
            _ = v17
            _ = v18
        }
        for fi19 := 0; fi19 < 8; fi19++ {
            var v20 bool = true
            v20 = true
            v20 = v20
            var v21 bool = v20
            _ = fi19
            _ = v20
            _ = v21
        }
        _ = fi8
    }
    v22 := []int{-936, 42, 32768}
    var v23 int64 = -9223372036854775808
    _ = v22
    _ = v23
    return ("foo bar" + ("foo bar" + ("\n\t" + "hello")))
}
func genFn1() uint64 {
    switch 42 {
    case -136:
        switch "a" {
        case "":
            v24 := GenS1{F0: false, F1: 128, F2: -1}
            _ = v24
        case "foo bar":
            for fi25 := 0; fi25 < 7; fi25++ {
                var v26 int8 = -1
                _ = fi25
                _ = v26
            }
        case "a":
            var v27 uint64 = 1
            _ = v27
        }
    default:
        var v28 bool = false
        _ = v28
    }
    var v29 uint32 = 0
    if true {
        v30 := GenS1{F0: ("" == ("a" + "hello")), F1: 1, F2: -115}
        _ = v30
    } else {
        v29 = (v29 * (^(v29)))
    }
    _ = v29
    return 1
}

func main() {
    var v1 uint32 = 1
    var v2 float32 = ((3.14159 / 1.0) * float32(1.5))
    var v3 uint64 = 4294967295
    v4 := []int{32767, int(v3)}
    var v5 float32 = (v2 / v2)
    if (true || (128 == 42)) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    for fi31 := 0; fi31 < 17; fi31++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi32 := 0; fi32 < 19; fi32++ {
            switch 4294967295 {
            case -9223372036854775808:
                v33 := GenS2{F0: ""}
                _ = v33
            case 42:
                fmt.Println(v1, fi32)
            case 2147483647:
                v2 = -2.25
            default:
                var v34 bool = (true || true)
                _ = v34
            }
            v2 -= (-2.25 * (1.5 + (v5 * ((0.0 + v5) + (v5 - float32(1e10))))))
            if (255 != 0) {
                var v35 int64 = -534
                fmt.Println(v2)
                _ = v35
            } else {
                fmt.Println(v1, v2)
                var v36 int8 = -36
                switch fi31 {
                case -2147483648:
                    if (!((v3 & (v3 - (v3 % (((v3 & uint64(v1))) | 1)))) > (v3 + 24))) {
                        v36 += v36
                        v2 += v2
                        v37 := GenS3{F0: -358}
                        _ = v37
                    } else {
                        v38 := GenS3{F0: 1}
                        _ = v38
                    }
                case 32768:
                    var v39 uint = 18446744073709551615
                    _ = v39
                case 127:
                    v5 += (v5 / v5)
                default:
                    v4 = append(v4, -2147483648)
                    v2 = (v2 - ((0.0 / (float32(float64(v5)) - (1.5 / (v5 - v5)))) + (v5 + (v2 / (((-2.25 - (v5 * 1.5)) + v2) * 0.1)))))
                }
                _ = v36
            }
            _ = fi32
        }
        _ = fi31
    }
    fmt.Println(v3)
    fmt.Println(v1, v5, v1)
    v1 += v1
    v40 := GenS3{F0: 1000}
    var v41 uint64 = (v3 & (v3 * v3))
    for fi42 := 0; fi42 < 8; fi42++ {
        var v43 int = fi42
        v3 = 42
        var v44 uint8 = 127
        _ = fi42
        _ = v43
        _ = v44
    }
    if true {
        var v45 float32 = float32(0.1)
        _ = len(v4)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v45
    } else {
        v46 := GenS2{F0: ("" + "hello")}
        fmt.Println(v3, v3, v3)
        switch 128 {
        case 127:
            var v47 uint8 = 0
            _ = v47
        case -9223372036854775808:
            var v48 int8 = 42
            _ = v48
        default:
            switch "foo bar" {
            case "foo bar":
                var v49 int32 = -911
                _ = v49
            case "hello":
                var v50 int = 931
                _ = v50
            default:
                v51 := GenS4{F0: false, F1: 42}
                _ = v51
            }
        }
        _ = v46
    }
    var v52 float32 = 3.14159
    for fi53 := 0; fi53 < 2; fi53++ {
        if ((216 >= -92) && (((!false) || (2147483647 == -2147483648)) || (!(2 <= 128)))) {
            var v54 uint = 2
            fmt.Println(fi53)
            v55 := GenS4{F0: false, F1: 0}
            _ = v54
            _ = v55
        }
        v4 = append(v4, 95)
        _ = fi53
    }
    for fi56 := 0; fi56 < 15; fi56++ {
        if false {
            for fi57 := 0; fi57 < 18; fi57++ {
                fmt.Println(v2)
                _ = fi57
            }
            v1 -= (v1 % ((uint32(v41)) | 1))
        }
        _ = len(v4)
        _ = fi56
    }
    var v58 float64 = (float64((v5 + v5)) * 0.1)
    if true {
        _ = len(v4)
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if true {
        switch 1 {
        case 283:
            v59 := GenS4{F0: (42 < 537), F1: 1}
            _ = v59
        default:
            v60 := GenS5{F0: 42, F1: 42}
            _ = v60
        }
        if (0 != 1) {
            fmt.Println(v52, v3)
            var v61 uint32 = 291
            v4 = append(v4, 42)
            _ = v61
        }
        v41 -= 0
    } else {
        switch "hello" {
        case "\n\t":
            if true {
                v62 := GenS6{F0: 0, F1: false, F2: 13}
                fmt.Println(v58, v58, v58)
                _ = v62
            } else {
                if false {
                    var v63 int8 = -128
                    var v64 int16 = 42
                    var v65 int8 = int8(v41)
                    _ = v63
                    _ = v64
                    _ = v65
                }
            }
        default:
            var v66 uint8 = 0
            _ = v66
        }
        var v67 bool = false
        v68 := GenS7{F0: -32768, F1: 0}
        _ = v67
        _ = v68
    }
    for fi69 := 0; fi69 < 20; fi69++ {
        var v70 string = ("\n\t" + "foo bar")
        if true {
            v1 = v1
        }
        fmt.Println(v70, v5)
        _ = fi69
        _ = v70
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v52
    _ = v58
}
