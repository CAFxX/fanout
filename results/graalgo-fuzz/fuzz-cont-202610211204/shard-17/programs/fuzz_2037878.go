package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 bool }
type GenS2 struct { F0 int64; F1 uint; F2 bool }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint32; F1 int; F2 uint8 }
type GenS6 struct { F0 uint32; F1 uint64; F2 uint }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint64; F1 uint16 }
type GenS9 struct { F0 string; F1 int; F2 int8 }
func genFn0(p0 int16, p1 uint64, p2 uint64) uint8 {
    v5 := []int{458}
    if len(v5) > 0 {
        v5[0] = 42
    }
    switch "\n\t" {
    case "":
        fmt.Println(p0, p1)
    }
    if true {
        var v6 uint8 = 42
        for fi7 := 0; fi7 < 11; fi7++ {
            fmt.Println(p1)
            p0 = (-(p0))
            var v8 uint8 = (v6 & 128)
            _ = fi7
            _ = v8
        }
        _ = v6
    } else {
        switch 863 {
        case 61:
            for fi9 := 0; fi9 < 10; fi9++ {
                if (true && false) {
                    p2 = p1
                    p1 = p2
                } else {
                    for fi10 := 0; fi10 < 5; fi10++ {
                        v11 := GenS1{F0: 998, F1: 71, F2: true}
                        _ = fi10
                        _ = v11
                    }
                }
                var v12 uint = 0
                _ = fi9
                _ = v12
            }
        case -9223372036854775808:
            p1 = (p2 + (p1 >> ((445) & 63)))
        default:
            for fi13 := 0; fi13 < 8; fi13++ {
                v14 := GenS1{F0: 0, F1: 53, F2: true}
                p0 += (p0 & (^(p0)))
                _ = fi13
                _ = v14
            }
        }
        p2 += (^(p2))
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        for fi16 := 0; fi16 < 11; fi16++ {
            switch (fi15 >> ((42) & 63)) {
            case -161:
                if (-971 != 0) {
                    var v17 int8 = -62
                    v18 := GenS1{F0: 778, F1: 57, F2: true}
                    v17 -= (v17 ^ v17)
                    _ = v17
                    _ = v18
                }
            case 42:
                fmt.Println(fi16, p0, p0)
            }
            var v19 string = "\n\t"
            _ = fi16
            _ = v19
        }
        if false {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            for fi20 := 0; fi20 < 1; fi20++ {
                switch (^(fi15)) {
                case 127:
                    if (!true) {
                        var v21 uint16 = 708
                        _ = v21
                    }
                default:
                    var v22 uint = 0
                    _ = v22
                }
                var v23 float64 = (0.0 * (-2.25 + (1.5 + (0.1 * (((-2.25 - (0.0 / 3.14159)) + -2.25) - ((float64(((3.14159 + ((0.0 + (3.14159 / 1e10)) - 0.1)) - (0.0 / 0.1))) + 0.1) - (1.5 + 0.0)))))))
                _ = fi20
                _ = v23
            }
            if true {
                if (false || (false && (-470 == fi15))) {
                    switch -1 {
                    case 42:
                        if len(v5) > 0 {
                            v5[0] = 58
                        }
                    case -442:
                        var v24 int32 = -971
                        _ = v24
                    case 32768:
                        for fi25 := 0; fi25 < 7; fi25++ {
                            if len(v5) > 0 {
                                v5[0] = fi25
                            }
                            fmt.Println(p1, fi25)
                            v26 := GenS2{F0: 0, F1: 1, F2: true}
                            _ = fi25
                            _ = v26
                        }
                    }
                }
                switch int(p0) {
                case 128:
                    var v27 uint32 = 913
                    _ = v27
                default:
                    p1 = p2
                }
                for fi28 := 0; fi28 < 6; fi28++ {
                    p2 = 18446744073709551615
                    if len(v5) > 0 {
                        v5[0] = (-(fi28))
                    }
                    switch "a" {
                    case "":
                        var v29 uint = 0
                        _ = v29
                    case "hello":
                        if false {
                            if len(v5) > 0 {
                                v5[0] = -822
                            }
                            v30 := GenS1{F0: 626, F1: 128, F2: ((-(p0)) != p0)}
                            switch 128 {
                            case -234:
                                p2 = 718
                            default:
                                var v31 bool = true
                                _ = v31
                            }
                            _ = v30
                        }
                        v32 := GenS3{F0: (p2 << ((42) & 63)), F1: 236}
                        _ = v32
                    }
                    if false {
                        p0 = p0
                    } else {
                        v33 := GenS4{F0: (-(fi15))}
                        _ = v33
                    }
                    _ = fi28
                }
            }
        } else {
            p1 = (^(p2))
            _ = len(v5)
        }
        _ = fi15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 42
}
func genFn1(p0 int16, p1 int32) int16 {
    switch "hello" {
    case "foo bar":
        p0 += p0
    default:
        if (float32(1e10) < 3.14159) {
            v34 := GenS4{F0: 128}
            p1 += p1
            _ = v34
        }
    }
    var v35 uint8 = 0
    v35 += genFn0(32767, 0, 401)
    p1 = p1
    for fi36 := 0; fi36 < 1; fi36++ {
        p1 = p1
        var v37 uint8 = (v35 * 1)
        _ = fi36
        _ = v37
    }
    _ = p0
    _ = p1
    _ = v35
    return p0
}

func main() {
    var v1 uint8 = 128
    var v2 uint32 = 2
    v3 := make(map[string]uint32)
    v4 := []string{}
    v1 = 127
    var v38 int8 = -128
    var v39 int32 = -2147483648
    v2 = (v2 & 905)
    fmt.Println(v38)
    for fi40 := 0; fi40 < 3; fi40++ {
        fmt.Println(v2, v38, v2)
        v2 += v2
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi40
    }
    for fi41 := 0; fi41 < 8; fi41++ {
        delete(v3, "hello")
        v42 := GenS5{F0: (v2 | v2), F1: -1, F2: v1}
        fmt.Println(v2, v1)
        _ = fi41
        _ = v42
    }
    var v43 string = ""
    switch ("a" + v43) {
    case "":
        for fi44 := 0; fi44 < 4; fi44++ {
            v1 -= v1
            v45 := GenS5{F0: 1, F1: (-(fi44)), F2: genFn0(1000, 2, uint64(v1))}
            _ = fi44
            _ = v45
        }
    case "a":
        v43 = (((v43 + (("hello" + "hello") + v43)) + ("a" + (v43 + (v43 + "hello")))) + v43)
    }
    switch ((v43 + v43) + ("a" + "foo bar")) {
    case "":
        for fi46 := 0; fi46 < 9; fi46++ {
            if false {
                v47 := GenS3{F0: 2, F1: 42}
                var v48 int64 = 42
                switch fi46 {
                case 0:
                    if (((!false) && false) && ((!false) || false)) {
                        v49 := GenS6{F0: (^(v2)), F1: 618, F2: 1}
                        v4 = append(v4, (v43 + v43))
                        var v50 float64 = ((((0.0 + 3.14159) * ((-2.25 - 1e10) * (0.0 - (0.0 - ((float64(((3.14159 + 1.5) + 0.0)) + 1.5) / 1.5))))) / ((-2.25 / 3.14159) * -2.25)) * (3.14159 / 1.5))
                        _ = v49
                        _ = v50
                    }
                case 2147483647:
                    v51 := GenS4{F0: fi46}
                    _ = v51
                case 9223372036854775807:
                    v38 += 42
                default:
                    v39 += int32(fi46)
                }
                _ = v47
                _ = v48
            } else {
                switch 4294967295 {
                case -2147483648:
                    var v52 int16 = 1
                    _ = v52
                case -157:
                    if true {
                        switch ("" + v43) {
                        case "":
                            if true {
                                for fi53 := 0; fi53 < 3; fi53++ {
                                    v1 += (v1 >> ((uint(v1)) & 63))
                                    v54 := GenS7{F0: "hello"}
                                    var v55 uint8 = (^(v1))
                                    _ = fi53
                                    _ = v54
                                    _ = v55
                                }
                                v56 := GenS8{F0: 2, F1: 255}
                                var v57 int64 = -9
                                _ = v56
                                _ = v57
                            } else {
                                fmt.Println(v43, v39, v2)
                                switch v43 {
                                case "hello":
                                    var v58 uint64 = 2
                                    _ = v58
                                case "":
                                    v3["hello"] = (v2 & 1)
                                case "foo bar":
                                    var v59 int16 = 1000
                                    _ = v59
                                default:
                                    var v60 int32 = v39
                                    _ = v60
                                }
                            }
                        case "hello":
                            v61 := GenS1{F0: 82, F1: 128, F2: true}
                            _ = v61
                        default:
                            for fi62 := 0; fi62 < 14; fi62++ {
                                if ((!(true && (-908 > 0))) || (int16(fi62) > -32768)) {
                                    if ((!(true && true)) || (128 >= 127)) {
                                        v38 = (v38 << ((4294967295) & 63))
                                    }
                                    v4 = append(v4, ("a" + "foo bar"))
                                }
                                if (v2 <= 2147483648) {
                                    switch 732 {
                                    case 32768:
                                        var v63 bool = false
                                        _ = v63
                                    case -9223372036854775808:
                                        _ = len(v3)
                                    case 2147483647:
                                        var v64 uint8 = v1
                                        _ = v64
                                    default:
                                        _ = len(v4)
                                    }
                                    var v65 uint32 = v2
                                    v1 += (^(v1))
                                    _ = v65
                                } else {
                                    var v66 uint32 = 2147483648
                                    v67 := GenS7{F0: ("a" + (v43 + "foo bar"))}
                                    fmt.Println(fi62, v39, v43)
                                    _ = v66
                                    _ = v67
                                }
                                _ = fi62
                            }
                        }
                        fmt.Println(fi46, v1, v39)
                    } else {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                case 1:
                    delete(v3, (((v43 + ((v43 + (v43 + "hello")) + v43)) + v43) + v43))
                }
                v68 := GenS9{F0: (v43 + (v43 + ("a" + v43))), F1: (-(fi46)), F2: v38}
                _ = v68
            }
            _ = fi46
        }
    }
    fmt.Println(v39, v43, v38)
    var v69 int64 = int64(v38)
    for fi70 := 0; fi70 < 19; fi70++ {
        var v71 int32 = v39
        _ = fi70
        _ = v71
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v43
    _ = v69
}
