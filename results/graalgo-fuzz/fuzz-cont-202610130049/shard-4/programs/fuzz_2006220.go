package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 uint; F1 bool; F2 bool }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint16; F1 int16; F2 uint64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int; F1 int16 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint16; F1 bool }
type GenS11 struct { F0 uint8; F1 string; F2 int64 }
func genFn0(p0 uint32, p1 int64) int16 {
    var v7 float64 = 0.0
    p1 += -170
    p1 -= 9223372036854775807
    v8 := []int{}
    for fi9 := 0; fi9 < 14; fi9++ {
        for fi10 := 0; fi10 < 9; fi10++ {
            v8 = append(v8, fi10)
            _ = fi10
        }
        var v11 int64 = p1
        v12 := GenS1{F0: (p0 >> ((1) & 63)), F1: 1, F2: p0}
        _ = fi9
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 bool, p1 int) int32 {
    if p0 {
        p0 = (42 < 42)
        var v13 int32 = 1
        _ = v13
    } else {
        switch "a" {
        case "\n\t":
            var v14 int64 = -195
            _ = v14
        default:
            v15 := GenS2{F0: 849, F1: 883}
            _ = v15
        }
        for fi16 := 0; fi16 < 20; fi16++ {
            v17 := []int{p1, -9223372036854775808, (fi16 >> ((216) & 63)), (fi16 ^ (fi16 % ((-930) | 1)))}
            p0 = true
            p0 = p0
            _ = fi16
            _ = v17
        }
        for fi18 := 0; fi18 < 6; fi18++ {
            var v19 uint = 4294967295
            _ = fi18
            _ = v19
        }
    }
    if p0 {
        for fi20 := 0; fi20 < 14; fi20++ {
            if (!(!((3.14159 + 3.14159) > 1.5))) {
                var v21 uint = 0
                _ = v21
            } else {
                v22 := GenS3{F0: 1, F1: ("foo bar" < "\n\t"), F2: false}
                p0 = p0
                for fi23 := 0; fi23 < 8; fi23++ {
                    p1 = 128
                    _ = fi23
                }
                _ = v22
            }
            for fi24 := 0; fi24 < 16; fi24++ {
                p0 = (-128 >= -1)
                v25 := []int{}
                switch (fi24 - fi20) {
                case 9223372036854775807:
                    p0 = true
                default:
                    if len(v25) >= 2 {
                        v25 = v25[1:len(v25)]
                    }
                }
                _ = fi24
                _ = v25
            }
            switch p1 {
            case -1:
                for fi26 := 0; fi26 < 15; fi26++ {
                    var v27 int = (fi26 + (^(fi20)))
                    if (!((1e10 * float64(-2.25)) >= -2.25)) {
                        v27 += (^(fi20))
                        v28 := []int{p1}
                        switch 42 {
                        case 127:
                            switch "hello" {
                            case "a":
                                switch 9223372036854775807 {
                                case 32767:
                                    if (genFn0(65535, -9223372036854775808) > 42) {
                                        if (9223372036854775807 > 1) {
                                            fmt.Println(fi20, fi20, v27)
                                            p0 = (-9223372036854775808 == (fi20 ^ (fi26 << ((255) & 63))))
                                            if false {
                                                if false {
                                                    if (-41 == -628) {
                                                        v29 := GenS3{F0: 2, F1: true, F2: (42 != 65535)}
                                                        _ = v29
                                                    }
                                                    var v30 int8 = 100
                                                    var v31 uint = 65535
                                                    _ = v30
                                                    _ = v31
                                                }
                                                v27 -= fi26
                                                v32 := GenS4{F0: -128}
                                                switch (("foo bar" + "a") + "") {
                                                case "":
                                                    if p0 {
                                                        fmt.Println(v27, fi20)
                                                    }
                                                case "a":
                                                    if len(v28) >= 2 {
                                                        v28 = v28[1:len(v28)]
                                                    }
                                                default:
                                                    v33 := GenS5{F0: 982, F1: genFn0(25, 473), F2: 1}
                                                    _ = v33
                                                }
                                                _ = v32
                                            } else {
                                                p0 = p0
                                                v27 -= (-(fi20))
                                            }
                                            fmt.Println(p1, fi20)
                                        } else {
                                            var v34 int = 463
                                            var v35 int64 = 42
                                            _ = v34
                                            _ = v35
                                        }
                                        fmt.Println(fi26, fi26)
                                        v36 := GenS4{F0: -1}
                                        _ = v36
                                    } else {
                                        p1 = (fi26 % ((v27) | 1))
                                    }
                                default:
                                    p0 = (false || (false || (true || (("foo bar" + "a") > ""))))
                                }
                            case "hello":
                                if p0 {
                                    v37 := GenS5{F0: 0, F1: -1, F2: 1}
                                    if (p0 && (!(0 < 161))) {
                                        p1 = v27
                                    }
                                    switch 127 {
                                    case 1:
                                        v38 := GenS6{F0: 42}
                                        _ = v38
                                    case 42:
                                        fmt.Println(fi26, v27)
                                    default:
                                        if true {
                                            fmt.Println(fi26)
                                        }
                                    }
                                    _ = v37
                                } else {
                                    if false {
                                        p1 = (p1 - 164)
                                        var v39 uint32 = 4294967295
                                        _ = v39
                                    }
                                }
                            case "foo bar":
                                p1 -= v27
                            }
                        default:
                            p1 += (fi26 << ((4294967295) & 63))
                        }
                        _ = v28
                    }
                    switch fi20 {
                    case -9223372036854775808:
                        v40 := GenS7{F0: -1, F1: genFn0(2, 392)}
                        _ = v40
                    default:
                        var v41 uint = 4294967295
                        _ = v41
                    }
                    v42 := []int{(v27 & fi20)}
                    p0 = (500 != 4294967295)
                    _ = fi26
                    _ = v27
                    _ = v42
                }
            default:
                var v43 uint = 108
                _ = v43
            }
            _ = fi20
        }
    }
    p0 = p0
    for fi44 := 0; fi44 < 7; fi44++ {
        var v45 uint32 = 464
        if (p0 && (!(2147483647 == -1))) {
            if ((0 <= 18446744073709551615) && (0 >= 2)) {
                v46 := []int{(p1 / ((761) | 1)), (-(fi44))}
                for fi47 := 0; fi47 < 5; fi47++ {
                    if true {
                        if p0 {
                            fmt.Println(p1)
                            var v48 uint16 = 42
                            p1 += (p1 << ((65535) & 63))
                            _ = v48
                        } else {
                            v49 := GenS8{F0: 0}
                            _ = v49
                        }
                    }
                    v50 := GenS9{F0: 937}
                    var v51 int16 = -1
                    p0 = (141 <= 42)
                    var v52 uint64 = 18446744073709551615
                    v53 := GenS10{F0: 255, F1: true}
                    _ = fi47
                    _ = v50
                    _ = v51
                    _ = v52
                    _ = v53
                }
                _ = v46
            }
        } else {
            p1 += -2147483648
        }
        v45 = 65535
        _ = fi44
        _ = v45
    }
    p0 = (!((!p0) && ((2 != 2) || (!(0 <= 635)))))
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 int8 = 1
    var v2 float32 = -2.25
    v3 := []int{}
    v4 := make(map[string]bool)
    var v5 float32 = v2
    v6 := []int{485}
    fmt.Println(v1, v5, v5)
    v1 -= 1
    for fi54 := 0; fi54 < 8; fi54++ {
        var v55 uint64 = 2
        v1 += 127
        _ = len(v4)
        _ = fi54
        _ = v55
    }
    switch "" {
    case "hello":
        var v56 bool = (0 <= 804)
        _ = v56
    case "":
        fmt.Println(v2)
    case "foo bar":
        var v57 float64 = (-2.25 * 3.14159)
        _ = v57
    default:
        var v58 uint32 = 288
        _ = v58
    }
    switch 910 {
    case -9223372036854775808:
        delete(v4, "\n\t")
    case 42:
        for fi59 := 0; fi59 < 3; fi59++ {
            fmt.Println(v1, fi59)
            _ = fi59
        }
    case 9223372036854775807:
        v2 += v5
    default:
        fmt.Println(v1)
    }
    v5 = -2.25
    var v60 int = -2147483648
    for fi61 := 0; fi61 < 15; fi61++ {
        v62 := v4[""]
        _ = fi61
        _ = v62
    }
    v60 = (v60 % (((^(v60))) | 1))
    if true {
        var v63 string = "\n\t"
        _ = v63
    } else {
        if (v1 > v1) {
            switch ("" + "a") {
            case "foo bar":
                var v64 float32 = 0.1
                _ = v64
            case "a":
                v4["a"] = false
            case "\n\t":
                v5 += 0.1
            }
            if (0.1 >= float32((1.5 / (1.5 * (-2.25 + (float64(((1.5 * v2) / 0.1)) * 0.1)))))) {
                var v65 float64 = (0.1 * (((float64(v2) * float64((v5 - 1e10))) + 1.5) - 0.1))
                fmt.Println(v5)
                for fi66 := 0; fi66 < 6; fi66++ {
                    fmt.Println(v5, v2)
                    var v67 uint8 = 255
                    v68 := GenS11{F0: (v67 + (v67 >> ((948) & 63))), F1: ("foo bar" + "a"), F2: 1}
                    _ = fi66
                    _ = v67
                    _ = v68
                }
                _ = v65
            } else {
                var v69 float64 = 0.0
                v5 += v5
                _ = v69
            }
            for fi70 := 0; fi70 < 1; fi70++ {
                v2 = v5
                v60 = (v60 / ((fi70) | 1))
                _ = fi70
            }
        } else {
            var v71 uint8 = 155
            v60 = (v60 % ((v60) | 1))
            _ = v71
        }
        var v72 string = (("foo bar" + "\n\t") + (("a" + "") + ("" + "a")))
        _ = v72
    }
    for fi73 := 0; fi73 < 6; fi73++ {
        var v74 int16 = 42
        _ = fi73
        _ = v74
    }
    v75 := GenS1{F0: 42, F1: 840, F2: 284}
    _ = len(v3)
    var v76 string = ""
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v60:", v60)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v75
    _ = v76
}
