package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 string }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 uint; F1 uint }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint16; F1 int32 }
type GenS7 struct { F0 uint8; F1 bool; F2 uint }
func genFn0(p0 int16, p1 string, p2 int32) string {
    for fi5 := 0; fi5 < 5; fi5++ {
        for fi6 := 0; fi6 < 15; fi6++ {
            fmt.Println(fi5, fi5, fi6)
            _ = fi6
        }
        _ = fi5
    }
    if (p0 == -32768) {
        var v7 string = "foo bar"
        switch 127 {
        case 4294967295:
            var v8 bool = (!false)
            _ = v8
        case 486:
            p0 = (^(p0))
        default:
            v9 := GenS1{F0: 2, F1: p2, F2: p1}
            _ = v9
        }
        _ = v7
    }
    v10 := GenS2{F0: true, F1: 42}
    p2 = (-(p2))
    p2 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return "foo bar"
}
func genFn1(p0 uint32, p1 uint16) uint8 {
    p0 = (^(p0))
    var v11 float32 = (-2.25 - ((1.5 + -2.25) * float32(1e10)))
    var v12 uint16 = (p1 >> ((948) & 63))
    v13 := GenS2{F0: (1e10 < ((0.0 / ((0.1 * 0.1) / 0.1)) - 3.14159)), F1: 183}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    return 42
}

func main() {
    var v1 int8 = -1
    v2 := make(map[uint32]uint32)
    var v3 uint16 = 255
    var v4 int32 = 694
    v14 := GenS2{F0: true, F1: 461}
    if true {
        var v15 uint32 = 149
        v16 := GenS2{F0: true, F1: 18446744073709551615}
        _ = v15
        _ = v16
    } else {
        switch 128 {
        case 4294967295:
            switch int(v3) {
            case 798:
                v2[0] = 42
            case 2147483647:
                v4 = int32(v3)
            default:
                v3 += (v3 / ((v3) | 1))
            }
        case 2147483647:
            var v17 int64 = -359
            _ = v17
        case -906:
            fmt.Println(v4, v1)
        }
        v1 += v1
        fmt.Println(v1, v1)
    }
    v18 := GenS2{F0: false, F1: 255}
    switch "" {
    case "a":
        if true {
            v19 := GenS3{F0: v1, F1: 0}
            _ = v19
        } else {
            for fi20 := 0; fi20 < 17; fi20++ {
                if true {
                    for fi21 := 0; fi21 < 8; fi21++ {
                        delete(v2, 0)
                        _ = fi21
                    }
                    fmt.Println(v4, v4)
                }
                v3 = v3
                v22 := GenS4{F0: 65535, F1: 0}
                _ = fi20
                _ = v22
            }
            v3 = (v3 * (v3 * (v3 + v3)))
            v23 := []int{128, 32767}
            _ = v23
        }
    case "\n\t":
        if false {
            if true {
                if ((v3 << ((4294967295) & 63)) >= v3) {
                    if (v1 != 53) {
                        var v24 bool = true
                        v25 := []int{204, -1}
                        for fi26 := 0; fi26 < 19; fi26++ {
                            v4 = v4
                            v3 += (^(v3))
                            _ = fi26
                        }
                        _ = v24
                        _ = v25
                    }
                    if true {
                        v27 := GenS4{F0: 2, F1: 2}
                        v3 -= v3
                        _ = v27
                    } else {
                        v4 = v4
                    }
                    var v28 string = ("" + (((("foo bar" + (((("foo bar" + "\n\t") + "a") + "\n\t") + "a")) + "") + "foo bar") + "foo bar"))
                    _ = v28
                }
                v1 = (-(v1))
                _ = len(v2)
            }
            for fi29 := 0; fi29 < 11; fi29++ {
                fmt.Println(v1, v1)
                if (!false) {
                    fmt.Println(v4, v3)
                } else {
                    v30 := v2[821]
                    v31 := GenS2{F0: (!(!((v1 | v1) != 51))), F1: 248}
                    _ = v30
                    _ = v31
                }
                if true {
                    for fi32 := 0; fi32 < 5; fi32++ {
                        if (75 >= 420) {
                            var v33 float32 = 0.0
                            _ = v33
                        } else {
                            if (18446744073709551615 < 0) {
                                fmt.Println(fi32)
                            }
                        }
                        switch 689 {
                        case -42:
                            v3 -= 1
                            if (((("a" <= "foo bar") && (genFn1(1, v3) > 128)) && (1 > 42)) && true) {
                                v34 := []int{}
                                v35 := GenS5{F0: 1}
                                var v36 int32 = (-(v4))
                                _ = v34
                                _ = v35
                                _ = v36
                            }
                        case -983:
                            v37 := GenS2{F0: (-32768 < -32768), F1: 42}
                            _ = v37
                        case -9223372036854775808:
                            v3 = 65535
                        }
                        var v38 int = (^(fi29))
                        _ = fi32
                        _ = v38
                    }
                    for fi39 := 0; fi39 < 20; fi39++ {
                        var v40 string = "a"
                        if true {
                            v41 := []int{fi29, (fi29 << ((1) & 63)), fi39}
                            _ = v41
                        } else {
                            if false {
                                _ = len(v2)
                                v42 := GenS5{F0: 198}
                                _ = v42
                            }
                        }
                        _ = fi39
                        _ = v40
                    }
                } else {
                    v43 := GenS4{F0: 780, F1: 216}
                    v4 += (v4 / (((v4 % ((v4) | 1))) | 1))
                    v3 += v3
                    _ = v43
                }
                _ = fi29
            }
        }
    }
    v4 = -1
    switch 982 {
    case 886:
        var v44 int8 = -1
        _ = v44
    case 2147483647:
        switch -804 {
        case 127:
            var v45 uint16 = 871
            _ = v45
        case 1:
            if false {
                fmt.Println(v1, v3)
                fmt.Println(v1)
                switch 906 {
                case 1:
                    v46 := GenS6{F0: v3, F1: v4}
                    _ = v46
                case 2147483647:
                    v47 := []int{197, 2147483647, 886, -92}
                    _ = v47
                }
            } else {
                var v48 int = -476
                if true {
                    for fi49 := 0; fi49 < 18; fi49++ {
                        var v50 string = ""
                        fmt.Println(v1, v4)
                        _ = fi49
                        _ = v50
                    }
                } else {
                    if true {
                        for fi51 := 0; fi51 < 13; fi51++ {
                            var v52 int32 = v4
                            _ = fi51
                            _ = v52
                        }
                        v1 = v1
                        fmt.Println(v4, v4)
                    } else {
                        var v53 int8 = (-(v1))
                        v53 = 100
                        if true {
                            v3 += v3
                        } else {
                            if false {
                                fmt.Println(v48, v53, v48)
                            } else {
                                v3 -= (^(v3))
                            }
                        }
                        _ = v53
                    }
                    switch (-(v48)) {
                    case 0:
                        switch (^(v48)) {
                        case 1:
                            v54 := []int{(-(v48))}
                            _ = v54
                        case 0:
                            if (-570 < -1) {
                                fmt.Println(v48, v1)
                            } else {
                                v4 += (^(v4))
                                _ = len(v2)
                            }
                        case -63:
                            if false {
                                var v55 int = (^(v48))
                                var v56 int16 = 42
                                v4 -= v4
                                _ = v55
                                _ = v56
                            }
                        }
                    case 9223372036854775807:
                        v4 = v4
                    case -9223372036854775808:
                        if true {
                            if true {
                                var v57 uint32 = 2147483648
                                if (!false) {
                                    var v58 float64 = 3.14159
                                    _ = v58
                                } else {
                                    switch v48 {
                                    case -69:
                                        var v59 uint8 = 128
                                        _ = v59
                                    default:
                                        v48 = 0
                                    }
                                    v60 := GenS3{F0: v1, F1: int64(v4)}
                                    _ = v60
                                }
                                _ = v57
                            } else {
                                var v61 uint16 = v3
                                fmt.Println(v61, v1)
                                delete(v2, 383)
                                _ = v61
                            }
                            var v62 int = 2147483647
                            _ = v62
                        } else {
                            if (255 < 196) {
                                if false {
                                    v2[0] = 42
                                    v1 = -128
                                    delete(v2, 65535)
                                }
                                v3 -= (v3 | v3)
                                v63 := []int{}
                                _ = v63
                            }
                            if (!false) {
                                for fi64 := 0; fi64 < 7; fi64++ {
                                    v48 += (fi64 | (fi64 ^ (fi64 * (-(fi64)))))
                                    v3 = 32767
                                    _ = fi64
                                }
                                v1 = v1
                            }
                            v2[393] = 16
                        }
                    }
                }
                _ = v48
            }
        default:
            var v65 uint64 = 326
            _ = v65
        }
    default:
        fmt.Println(v4)
    }
    switch 0 {
    case -2147483648:
        fmt.Println(v1)
    case 2147483647:
        v66 := GenS7{F0: 255, F1: false, F2: 18446744073709551615}
        _ = v66
    case 590:
        if (1e10 == (float64(3.14159) / (1e10 * ((0.1 - -2.25) - 1.5)))) {
            if (2147483647 < 195) {
                v1 = v1
            }
            var v67 int = 127
            fmt.Println(v3, v67)
            _ = v67
        }
    }
    for fi68 := 0; fi68 < 20; fi68++ {
        v3 = (v3 >> ((1) & 63))
        _ = fi68
    }
    v1 += (v1 - 127)
    v69 := []int{}
    switch "a" {
    case "foo bar":
        _ = len(v69)
    }
    var v70 bool = (true && true)
    if (1 < 1) {
        fmt.Println(v1, v70, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v69:", len(v69), v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v69
    _ = v70
}
