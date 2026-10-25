package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 int; F1 int32; F2 uint8 }
type GenS4 struct { F0 uint64; F1 int64; F2 int8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint16; F1 uint64 }
type GenS8 struct { F0 uint32; F1 int32; F2 int32 }
func genFn0(p0 uint64, p1 uint16) uint8 {
    fmt.Println(p1, p1)
    v5 := []int{42, 32767, -235}
    for fi6 := 0; fi6 < 13; fi6++ {
        for fi7 := 0; fi7 < 3; fi7++ {
            var v8 uint8 = uint8(p0)
            v9 := GenS1{F0: -655, F1: 0}
            switch -1 {
            case 4294967295:
                var v10 uint16 = 32767
                _ = v10
            case 82:
                var v11 uint = 42
                _ = v11
            case -361:
                var v12 string = "a"
                _ = v12
            }
            _ = fi7
            _ = v8
            _ = v9
        }
        switch "\n\t" {
        case "\n\t":
            var v13 int8 = 1
            _ = v13
        case "a":
            switch "\n\t" {
            case "foo bar":
                var v14 uint64 = p0
                _ = v14
            default:
                switch ("a" + "") {
                case "":
                    var v15 uint32 = 65535
                    _ = v15
                case "foo bar":
                    var v16 bool = true
                    _ = v16
                case "a":
                    for fi17 := 0; fi17 < 8; fi17++ {
                        p1 = p1
                        var v18 int16 = 32767
                        fmt.Println(fi6)
                        v19 := GenS2{F0: 2147483648, F1: 2147483647}
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                default:
                    var v20 int = fi6
                    _ = v20
                }
            }
        default:
            switch ("" + "") {
            case "foo bar":
                fmt.Println(p0)
            case "hello":
                p1 -= p1
            case "":
                v21 := GenS3{F0: fi6, F1: -798, F2: 28}
                _ = v21
            }
        }
        _ = fi6
    }
    p1 += (p1 - p1)
    var v22 int32 = 2147483647
    _ = p0
    _ = p1
    _ = v22
    _ = v5
    return 42
}
func genFn1(p0 string, p1 uint64, p2 int64) int8 {
    for fi23 := 0; fi23 < 12; fi23++ {
        v24 := GenS2{F0: 0, F1: uint32(p1)}
        _ = fi23
        _ = v24
    }
    v25 := GenS2{F0: 65535, F1: 335}
    for fi26 := 0; fi26 < 8; fi26++ {
        if (491 != 0) {
            p1 = p1
            v27 := GenS4{F0: 18446744073709551615, F1: (p2 | p2), F2: 42}
            if ((((0.1 + 0.0) - 1e10) >= ((1e10 * -2.25) * 0.1)) || false) {
                for fi28 := 0; fi28 < 16; fi28++ {
                    if true {
                        p0 = p0
                        var v29 int = (^(fi28))
                        v30 := GenS1{F0: p2, F1: (p2 >> ((18446744073709551615) & 63))}
                        p2 = p2
                        _ = v29
                        _ = v30
                    } else {
                        v31 := GenS5{F0: 500}
                        _ = v31
                    }
                    _ = fi28
                }
                v32 := []int{fi26, 9223372036854775807, (fi26 & fi26)}
                _ = v32
            }
            _ = v27
        } else {
            p1 = (p1 + (^(p1)))
        }
        for fi33 := 0; fi33 < 4; fi33++ {
            v34 := []int{}
            if ((0.0 - (0.0 - ((1.5 * 1.5) + 3.14159))) <= 3.14159) {
                v35 := GenS6{F0: 0}
                _ = v35
            } else {
                var v36 int = int(p2)
                p0 = p0
                _ = v36
            }
            if len(v34) >= 2 {
                v34 = v34[1:len(v34)]
            }
            _ = fi33
            _ = v34
        }
        _ = fi26
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    return -3
}

func main() {
    v1 := []int{32767, 9223372036854775807}
    v2 := []int{127, 128}
    v3 := []string{}
    v4 := make(map[string]bool)
    var v37 uint8 = genFn0(4294967295, 993)
    var v38 int64 = 42
    if ((42 > -938) || true) {
        if false {
            for fi39 := 0; fi39 < 6; fi39++ {
                switch (-(fi39)) {
                case 0:
                    if ((138 <= 70) && true) {
                        for fi40 := 0; fi40 < 12; fi40++ {
                            v37 = 127
                            _ = fi40
                        }
                        if (255 <= 0) {
                            if true {
                                switch (fi39 - 42) {
                                case -930:
                                    v41 := GenS2{F0: 42, F1: 4294967295}
                                    _ = v41
                                case 9223372036854775807:
                                    v37 = genFn0(4294967295, 255)
                                case -1:
                                    fmt.Println(v37)
                                }
                                v42 := GenS7{F0: 0, F1: 18446744073709551615}
                                switch "a" {
                                case "\n\t":
                                    var v43 uint = 1
                                    _ = v43
                                case "foo bar":
                                    var v44 string = ("hello" + "\n\t")
                                    _ = v44
                                default:
                                    if false {
                                        _ = len(v4)
                                        v37 += v37
                                        v38 -= 1
                                    }
                                }
                                _ = v42
                            }
                            v37 = (^(v37))
                        } else {
                            var v45 int = fi39
                            var v46 int8 = -89
                            _ = v45
                            _ = v46
                        }
                        var v47 int = (^(fi39))
                        _ = v47
                    } else {
                        var v48 uint8 = (v37 + (v37 * genFn0(0, 989)))
                        var v49 int8 = -128
                        _ = v48
                        _ = v49
                    }
                default:
                    delete(v4, ("foo bar" + "foo bar"))
                }
                v4[""] = (999 != (fi39 >> ((42) & 63)))
                fmt.Println(fi39)
                _ = fi39
            }
            v50 := v4["foo bar"]
            fmt.Println(v38, v37)
            _ = v50
        }
        if (18446744073709551615 <= 2) {
            fmt.Println(v37)
            switch 32768 {
            case -335:
                v38 -= (v38 % ((v38) | 1))
            }
            for fi51 := 0; fi51 < 14; fi51++ {
                var v52 bool = true
                _ = fi51
                _ = v52
            }
        }
    } else {
        v1 = append(v1, -1)
    }
    var v53 float64 = 3.14159
    var v54 int = 9223372036854775807
    if false {
        var v55 uint8 = v37
        var v56 float64 = v53
        var v57 uint16 = 255
        _ = v55
        _ = v56
        _ = v57
    }
    v58 := GenS8{F0: 0, F1: -1, F2: int32(v38)}
    if (true || (!(false && false))) {
        delete(v4, "\n\t")
        v54 = (v54 | v54)
    }
    var v59 uint8 = (^(v37))
    for fi60 := 0; fi60 < 6; fi60++ {
        v61 := GenS7{F0: 208, F1: 0}
        v62 := GenS4{F0: 4294967295, F1: (v38 << ((55) & 63)), F2: 100}
        v63 := GenS3{F0: 4294967295, F1: -118, F2: v59}
        _ = fi60
        _ = v61
        _ = v62
        _ = v63
    }
    v64 := GenS8{F0: 2, F1: 1, F2: 364}
    _ = len(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v53
    _ = v54
    _ = v58
    _ = v59
    _ = v64
}
