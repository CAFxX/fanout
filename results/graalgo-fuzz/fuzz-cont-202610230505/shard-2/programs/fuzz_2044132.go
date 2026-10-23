package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 int16; F2 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int; F1 uint32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint16; F1 int8 }
type GenS7 struct { F0 uint64; F1 uint8; F2 uint16 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 bool; F1 uint }
type GenS10 struct { F0 uint16 }
func genFn0(p0 int32, p1 int16, p2 uint32) int16 {
    var v4 string = "\n\t"
    switch ("" + "") {
    case "hello":
        if true {
            if (100 != -16) {
                v5 := GenS1{F0: 459}
                p0 = p0
                _ = v5
            }
        } else {
            for fi6 := 0; fi6 < 15; fi6++ {
                var v7 float64 = 1.5
                v8 := []int{(fi6 & fi6), 32768}
                _ = fi6
                _ = v7
                _ = v8
            }
            var v9 uint16 = 0
            var v10 float32 = 1e10
            _ = v10
            _ = v9
        }
    case "":
        if true {
            var v11 float32 = 1.5
            _ = v11
        }
    default:
        v12 := GenS1{F0: 0}
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return p1
}
func genFn1(p0 string, p1 int64, p2 uint32) uint32 {
    if (p0 == "\n\t") {
        p2 -= 2
    } else {
        if true {
            v13 := GenS2{F0: 79, F1: 850, F2: 1}
            _ = v13
        }
        var v14 uint32 = p2
        var v15 uint64 = 42
        _ = v14
        _ = v15
    }
    v16 := GenS2{F0: 127, F1: genFn0(1, 1, (^(p2))), F2: 42}
    var v17 int = -283
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    return (p2 - (^(p2)))
}

func main() {
    v1 := []int{923}
    var v2 uint32 = 65535
    var v3 float64 = -2.25
    v3 = ((v3 - (0.1 + v3)) + (v3 - (v3 + v3)))
    for fi18 := 0; fi18 < 17; fi18++ {
        if true {
            if false {
                var v19 int32 = -2147483648
                switch ((("" + "a") + "\n\t") + "") {
                case "":
                    if false {
                        fmt.Println(v2, v19)
                    }
                case "a":
                    v19 = v19
                default:
                    v20 := GenS1{F0: (-(fi18))}
                    _ = v20
                }
                _ = v19
            } else {
                var v21 uint32 = v2
                _ = v21
            }
            for fi22 := 0; fi22 < 1; fi22++ {
                var v23 float64 = (v3 + v3)
                v24 := GenS1{F0: fi18}
                v25 := GenS3{F0: (fi18 & fi22)}
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
            }
            var v26 int16 = 0
            _ = v26
        }
        _ = fi18
    }
    for fi27 := 0; fi27 < 16; fi27++ {
        fmt.Println(v2, v3, v3)
        _ = len(v1)
        _ = fi27
    }
    var v28 int64 = 1
    if ((false && false) && ((("a" + "foo bar") + "hello") == (("hello" + ("" + "")) + "foo bar"))) {
        if (!true) {
            for fi29 := 0; fi29 < 4; fi29++ {
                v2 -= v2
                _ = fi29
            }
            _ = len(v1)
        }
    } else {
        v30 := GenS2{F0: 0, F1: 32767, F2: 0}
        _ = v30
    }
    switch 128 {
    case 42:
        var v31 uint32 = 222
        _ = v31
    case 32768:
        if (true || false) {
            v3 = 3.14159
        } else {
            v1 = append(v1, 9223372036854775807)
            switch 128 {
            case 2147483647:
                for fi32 := 0; fi32 < 12; fi32++ {
                    for fi33 := 0; fi33 < 16; fi33++ {
                        v28 += 9223372036854775807
                        v3 += float64(0.0)
                        _ = fi33
                    }
                    v2 = 2147483647
                    var v34 int8 = 31
                    _ = fi32
                    _ = v34
                }
            }
        }
    case 4294967295:
        var v35 uint = 1
        _ = v35
    default:
        for fi36 := 0; fi36 < 17; fi36++ {
            fmt.Println(v3, fi36)
            v2 = v2
            if false {
                var v37 uint8 = 42
                switch fi36 {
                case 73:
                    if true {
                        if true {
                            v2 -= genFn1(("\n\t" + "a"), 1, 1)
                        } else {
                            var v38 uint16 = 255
                            _ = v38
                        }
                        if false {
                            v39 := GenS4{F0: (fi36 % (((fi36 * 127)) | 1)), F1: (v2 * (v2 * genFn1(("" + ""), int64(v2), 65535)))}
                            _ = v39
                        }
                    } else {
                        if true {
                            v2 = v2
                            v3 += float64((((-2.25 * -2.25) - 1.5) / 0.1))
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        } else {
                            for fi40 := 0; fi40 < 9; fi40++ {
                                fmt.Println(fi36)
                                var v41 uint16 = 255
                                v2 += genFn1("", (v28 - v28), v2)
                                _ = fi40
                                _ = v41
                            }
                        }
                    }
                case 0:
                    for fi42 := 0; fi42 < 11; fi42++ {
                        if len(v1) > 0 {
                            v1[0] = fi42
                        }
                        if false {
                            fmt.Println(v37, v3, fi42)
                            if true {
                                var v43 int = fi42
                                var v44 bool = true
                                switch fi42 {
                                case 32767:
                                    v3 += ((v3 / v3) * v3)
                                case 2147483647:
                                    v1 = append(v1, fi42)
                                    v45 := GenS5{F0: 1}
                                    _ = v45
                                case 658:
                                    v28 -= v28
                                default:
                                    var v46 int16 = 1
                                    if v44 {
                                        if v44 {
                                            v47 := GenS3{F0: v43}
                                            _ = v47
                                        }
                                        if false {
                                            if (3.14159 <= 3.14159) {
                                                if len(v1) > 0 {
                                                    v1[0] = 0
                                                }
                                            }
                                            fmt.Println(v37, v28)
                                            v46 = (v46 | (-(v46)))
                                        } else {
                                            v48 := GenS6{F0: 255, F1: -128}
                                            switch 501 {
                                            case -9223372036854775808:
                                                var v49 uint = 65535
                                                _ = v49
                                            case 32768:
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                if ((v2 / ((0) | 1)) <= (v2 >> ((0) & 63))) {
                                                    var v50 uint16 = uint16(v37)
                                                    v1 = append(v1, fi42)
                                                    _ = v50
                                                }
                                            case 474:
                                                fmt.Println(v46, v3, fi42)
                                            }
                                            fmt.Println(v46)
                                            _ = v48
                                        }
                                        var v51 int16 = (v46 * (v46 + (v46 ^ (v46 % ((-634) | 1)))))
                                        _ = v51
                                    } else {
                                        v44 = v44
                                        v37 += (v37 + (v37 ^ 1))
                                        fmt.Println(v37, v3, fi36)
                                    }
                                    _ = v46
                                }
                                _ = v43
                                _ = v44
                            } else {
                                v2 -= (^(v2))
                            }
                        }
                        switch "a" {
                        case "":
                            v2 = (v2 - 0)
                        case "a":
                            if true {
                                v2 = (v2 >> ((255) & 63))
                                v52 := GenS7{F0: 18446744073709551615, F1: (v37 + (v37 % ((2) | 1))), F2: uint16(v2)}
                                v28 += (v28 % ((v28) | 1))
                                _ = v52
                            } else {
                                if (32767 >= 1000) {
                                    if true {
                                        var v53 int8 = -1
                                        _ = v53
                                    }
                                }
                                var v54 float32 = -2.25
                                _ = v54
                            }
                            switch (^(fi42)) {
                            case 31:
                                v55 := GenS8{F0: 127}
                                _ = v55
                            case 128:
                                switch "hello" {
                                case "hello":
                                    fmt.Println(v37, fi42, v2)
                                case "a":
                                    var v56 int16 = 727
                                    _ = v56
                                case "foo bar":
                                    var v57 int32 = 610
                                    _ = v57
                                }
                                if true {
                                    fmt.Println(fi42, v2)
                                    var v58 uint64 = 0
                                    _ = v58
                                } else {
                                    var v59 bool = false
                                    _ = v59
                                }
                            case 32767:
                                var v60 int = (fi36 << ((778) & 63))
                                _ = v60
                            default:
                                v37 = (^(v37))
                            }
                        default:
                            var v61 int32 = -1
                            _ = v61
                        }
                        _ = fi42
                    }
                default:
                    switch "" {
                    case "\n\t":
                        var v62 int = fi36
                        _ = v62
                    case "foo bar":
                        if (((("" <= "\n\t") && (!((!false) || true))) || true) || true) {
                            for fi63 := 0; fi63 < 13; fi63++ {
                                if (1000 < 0) {
                                    fmt.Println(fi63)
                                    v3 = ((1e10 - v3) + v3)
                                }
                                switch (("\n\t" + (("\n\t" + "hello") + "")) + ("\n\t" + ("hello" + ("foo bar" + "foo bar")))) {
                                case "foo bar":
                                    var v64 string = ""
                                    _ = v64
                                case "a":
                                    var v65 uint = 18446744073709551615
                                    v66 := GenS9{F0: true, F1: uint(v37)}
                                    v3 = v3
                                    _ = v65
                                    _ = v66
                                case "":
                                    var v67 uint = 0
                                    _ = v67
                                default:
                                    v68 := GenS3{F0: (fi63 / ((-489) | 1))}
                                    _ = v68
                                }
                                var v69 int32 = -723
                                _ = fi63
                                _ = v69
                            }
                            v3 -= 3.14159
                        } else {
                            for fi70 := 0; fi70 < 16; fi70++ {
                                v3 -= v3
                                _ = fi70
                            }
                        }
                    case "":
                        var v71 uint32 = v2
                        _ = v71
                    }
                }
                var v72 int16 = 1000
                _ = v37
                _ = v72
            } else {
                fmt.Println(fi36, fi36, v28)
            }
            _ = fi36
        }
    }
    switch ("foo bar" + "a") {
    case "a":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "hello":
        switch 1 {
        case 129:
            v3 = (v3 + 1.5)
        case 2147483647:
            v2 -= 185
        case 42:
            v73 := GenS10{F0: 255}
            _ = v73
        }
    }
    for fi74 := 0; fi74 < 10; fi74++ {
        for fi75 := 0; fi75 < 16; fi75++ {
            _ = len(v1)
            v28 += v28
            _ = fi75
        }
        var v76 string = (("\n\t" + "\n\t") + ("hello" + "foo bar"))
        _ = fi74
        _ = v76
    }
    var v77 int64 = 42
    if len(v1) > 0 {
        v1[0] = 783
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v28 = v77
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v77
}
