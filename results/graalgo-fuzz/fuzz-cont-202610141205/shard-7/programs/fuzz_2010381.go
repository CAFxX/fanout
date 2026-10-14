package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint32; F1 int8; F2 uint32 }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int32; F1 int8; F2 int32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int64; F1 int8; F2 uint8 }
type GenS8 struct { F0 uint8; F1 uint32; F2 uint8 }
func genFn0() int32 {
    for fi6 := 0; fi6 < 11; fi6++ {
        for fi7 := 0; fi7 < 7; fi7++ {
            v8 := []int{fi7}
            var v9 uint8 = 164
            _ = fi7
            _ = v8
            _ = v9
        }
        fmt.Println(fi6)
        _ = fi6
    }
    var v10 float64 = float64(3.14159)
    for fi11 := 0; fi11 < 11; fi11++ {
        v10 += ((0.0 - 3.14159) * -2.25)
        var v12 uint32 = 65535
        var v13 int64 = 9223372036854775807
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := []int{9223372036854775807, 4294967295, -393}
    _ = v10
    _ = v14
    return 1
}
func genFn1(p0 uint, p1 int64) int {
    p0 += p0
    fmt.Println(p0)
    _ = p0
    _ = p1
    return 855
}

func main() {
    v1 := []int{734, 127, -63, -40}
    v2 := []uint32{}
    var v3 uint8 = 31
    var v4 bool = (true && false)
    var v5 float64 = (0.0 - (float64(float32(((-2.25 / (0.0 + float64((((-2.25 - 0.1) * float32(0.1)) / (1e10 + (0.0 + float32(1e10))))))) / (1e10 + 3.14159)))) * 3.14159))
    switch ("a" + ("a" + "hello")) {
    case "":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "foo bar":
        fmt.Println(v4, v4)
    }
    v5 = v5
    switch "\n\t" {
    case "foo bar":
        for fi15 := 0; fi15 < 3; fi15++ {
            var v16 int32 = -367
            switch ("a" + "hello") {
            case "a":
                for fi17 := 0; fi17 < 11; fi17++ {
                    fmt.Println(fi15)
                    v4 = v4
                    if len(v1) > 0 {
                        v1[0] = (fi15 << ((0) & 63))
                    }
                    _ = fi17
                }
            case "foo bar":
                switch "hello" {
                case "hello":
                    var v18 uint64 = 18446744073709551615
                    _ = v18
                case "\n\t":
                    var v19 uint8 = v3
                    _ = v19
                default:
                    v16 = 0
                }
            case "hello":
                for fi20 := 0; fi20 < 9; fi20++ {
                    var v21 uint8 = v3
                    _ = fi20
                    _ = v21
                }
            }
            _ = fi15
            _ = v16
        }
    case "hello":
        v1 = append(v1, 128)
    case "\n\t":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    fmt.Println(v5, v4, v5)
    if (!false) {
        v1 = append(v1, -9223372036854775808)
        var v22 uint32 = 333
        if v4 {
            v5 += ((((v5 + v5) - v5) - v5) * 1e10)
            for fi23 := 0; fi23 < 12; fi23++ {
                v2 = append(v2, (^(v22)))
                _ = fi23
            }
            var v24 int16 = -1
            _ = v24
        } else {
            var v25 int16 = 42
            _ = v25
        }
        _ = v22
    }
    switch ("a" + (("a" + "a") + "")) {
    case "hello":
        var v26 float64 = v5
        _ = v26
    }
    var v27 uint64 = 4294967295
    switch ("foo bar" + ("\n\t" + ("hello" + "a"))) {
    case "a":
        if (!(0 != genFn0())) {
            if (v4 && v4) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                for fi28 := 0; fi28 < 14; fi28++ {
                    for fi29 := 0; fi29 < 3; fi29++ {
                        v5 = ((v5 / float64(1e10)) * (1.5 * float64(3.14159)))
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v30 int8 = -89
                        var v31 int16 = -1
                        var v32 float64 = float64(0.1)
                        v3 = (v3 | (v3 - (v3 % ((v3) | 1))))
                        _ = fi29
                        _ = v30
                        _ = v31
                        _ = v32
                    }
                    v3 = v3
                    _ = fi28
                }
                v5 += (v5 / float64(float32(3.14159)))
            }
            if (v4 && (!((float32(0.1) >= 1e10) && true))) {
                var v33 bool = v4
                _ = v33
            } else {
                for fi34 := 0; fi34 < 13; fi34++ {
                    v35 := GenS1{F0: 1, F1: 127}
                    v36 := GenS2{F0: 611, F1: -128, F2: 554}
                    _ = fi34
                    _ = v35
                    _ = v36
                }
                v37 := GenS1{F0: -60, F1: -104}
                for fi38 := 0; fi38 < 20; fi38++ {
                    v39 := GenS3{F0: 2, F1: ""}
                    v27 = v27
                    _ = fi38
                    _ = v39
                }
                _ = v37
            }
            var v40 bool = v4
            _ = v40
        } else {
            v41 := GenS4{F0: 4294967295}
            var v42 int64 = 0
            var v43 uint32 = 1
            _ = v41
            _ = v42
            _ = v43
        }
    case "foo bar":
        if v4 {
            v27 -= 1
        } else {
            v3 = v3
            var v44 string = "\n\t"
            _ = v44
        }
    default:
        for fi45 := 0; fi45 < 20; fi45++ {
            if v4 {
                v3 = uint8(v27)
            }
            v46 := GenS5{F0: -564, F1: 70, F2: -536}
            _ = fi45
            _ = v46
        }
    }
    for fi47 := 0; fi47 < 14; fi47++ {
        v4 = (!false)
        _ = fi47
    }
    var v48 uint8 = v3
    var v49 string = ("hello" + "hello")
    if true {
        if (0 <= 255) {
            if (!(2147483647 >= genFn0())) {
                fmt.Println(v4)
            }
        } else {
            if (v4 || (9223372036854775807 >= 0)) {
                switch v49 {
                case "foo bar":
                    if len(v2) > 0 {
                        v2[0] = 65535
                    }
                case "a":
                    if len(v2) > 0 {
                        v2[0] = 949
                    }
                }
                var v50 int = int(v48)
                v51 := GenS5{F0: 42, F1: -1, F2: -1}
                _ = v50
                _ = v51
            } else {
                switch 0 {
                case -2147483648:
                    v52 := GenS5{F0: 42, F1: -24, F2: -1}
                    _ = v52
                case 128:
                    var v53 uint = 919
                    _ = v53
                case 2147483647:
                    if (!false) {
                        v54 := GenS6{F0: 42}
                        if len(v2) > 0 {
                            v2[0] = 844
                        }
                        for fi55 := 0; fi55 < 19; fi55++ {
                            if ((1e10 * -2.25) != (1e10 + 1e10)) {
                                var v56 uint32 = 65535
                                _ = v56
                            }
                            if ((!v4) && v4) {
                                for fi57 := 0; fi57 < 2; fi57++ {
                                    if (!(42 != 993)) {
                                        v58 := GenS7{F0: 362, F1: int8(v3), F2: v48}
                                        _ = v58
                                    } else {
                                        var v59 int32 = 492
                                        var v60 uint8 = 127
                                        _ = v59
                                        _ = v60
                                    }
                                    switch int(v3) {
                                    case 1:
                                        var v61 uint16 = 32767
                                        _ = v61
                                    }
                                    var v62 uint = 255
                                    _ = fi57
                                    _ = v62
                                }
                                v4 = v4
                            }
                            var v63 bool = v4
                            _ = fi55
                            _ = v63
                        }
                        _ = v54
                    } else {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                default:
                    if v4 {
                        v48 += v48
                    }
                }
                var v64 uint64 = v27
                _ = v64
            }
            for fi65 := 0; fi65 < 3; fi65++ {
                var v66 float32 = 1e10
                if (!v4) {
                    var v67 uint16 = 381
                    var v68 int64 = 0
                    if (!((true || ((-(fi65)) != fi65)) || v4)) {
                        if len(v2) > 0 {
                            v2[0] = 2147483648
                        }
                    }
                    _ = v67
                    _ = v68
                } else {
                    v69 := GenS8{F0: (v3 << ((357) & 63)), F1: 2147483648, F2: (^(v48))}
                    var v70 float64 = (v5 * 3.14159)
                    _ = v69
                    _ = v70
                }
                for fi71 := 0; fi71 < 15; fi71++ {
                    v5 = (v5 - v5)
                    v4 = false
                    _ = fi71
                }
                _ = fi65
                _ = v66
            }
            var v72 int64 = -1
            _ = v72
        }
        var v73 int8 = 42
        if len(v1) > 0 {
            v1[0] = -9223372036854775808
        }
        _ = v73
    }
    if ((v4 || true) && v4) {
        fmt.Println(v49)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
}
