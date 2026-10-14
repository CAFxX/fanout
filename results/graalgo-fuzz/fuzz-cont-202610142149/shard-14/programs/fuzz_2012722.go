package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 uint32; F1 int8; F2 int }
type GenS6 struct { F0 int; F1 uint64; F2 uint8 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint8; F1 int16 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 uint16; F1 int32; F2 int64 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 uint; F1 int16 }
type GenS14 struct { F0 uint64; F1 int }
type GenS15 struct { F0 bool; F1 int16 }
func genFn0(p0 string, p1 int16, p2 uint) int32 {
    var v4 bool = (2 >= 255)
    switch "foo bar" {
    case "\n\t":
        p0 = "\n\t"
    case "a":
        var v5 float32 = 0.1
        _ = v5
    }
    fmt.Println(v4, p0, v4)
    fmt.Println(v4, p1, p1)
    var v6 uint64 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v6
    return 1
}

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{170, 2147483648}
    v3 := make(map[string]string)
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 uint64 = 0
        _ = fi7
        _ = v8
    }
    var v9 int16 = 155
    fmt.Println(v9)
    v10 := GenS1{F0: 127}
    var v11 int8 = 1
    switch 707 {
    case -546:
        v1["a"] = 42
    case -2147483648:
        switch 561 {
        case -376:
            if len(v2) > 0 {
                v2[0] = 4294967295
            }
        }
    case 128:
        var v12 int64 = -141
        _ = v12
    default:
        var v13 float32 = 3.14159
        _ = v13
    }
    if (v9 <= -32768) {
        var v14 int16 = (^(v9))
        _ = v14
    } else {
        if ((v9 >> ((18446744073709551615) & 63)) == (v9 % ((v9) | 1))) {
            var v15 uint64 = 774
            _ = v15
        } else {
            for fi16 := 0; fi16 < 7; fi16++ {
                v17 := GenS2{F0: fi16, F1: "\n\t"}
                for fi18 := 0; fi18 < 13; fi18++ {
                    fmt.Println(fi18, fi16, fi18)
                    v19 := GenS2{F0: (fi16 >> ((2) & 63)), F1: ("\n\t" + ("\n\t" + ("foo bar" + (("a" + "a") + "\n\t"))))}
                    if (!((v9 + (v9 ^ v9)) != (^(v9)))) {
                        if ((0.1 - 0.0) < 3.14159) {
                            switch "hello" {
                            case "\n\t":
                                v20 := GenS3{F0: true}
                                var v21 float32 = (float32(((1e10 / 1.0) * float64((float32(3.14159) * 3.14159)))) * 1e10)
                                switch "foo bar" {
                                case "\n\t":
                                    var v22 int8 = (^(v11))
                                    _ = v22
                                case "":
                                    if len(v2) > 0 {
                                        v2[0] = 65535
                                    }
                                case "foo bar":
                                    if true {
                                        var v23 uint32 = 41
                                        v24 := GenS3{F0: (9223372036854775807 != 0)}
                                        _ = v23
                                        _ = v24
                                    }
                                }
                                _ = v20
                                _ = v21
                            }
                            switch ("foo bar" + ("foo bar" + (("" + "hello") + "foo bar"))) {
                            case "":
                                if true {
                                    v25 := GenS4{F0: 1, F1: false}
                                    v9 = v9
                                    _ = v25
                                } else {
                                    if (0 < -1) {
                                        if (672 == 136) {
                                            switch ("\n\t" + "") {
                                            case "a":
                                                switch "a" {
                                                case "hello":
                                                    v9 += int16(v11)
                                                case "\n\t":
                                                    v9 += v9
                                                case "foo bar":
                                                    v26 := v3["hello"]
                                                    _ = v26
                                                default:
                                                    v11 -= 72
                                                }
                                            case "":
                                                if ((0 < 4294967295) && true) {
                                                    if false {
                                                        v27 := GenS4{F0: 830, F1: (!(2147483647 != 2147483647))}
                                                        v9 = -1
                                                        v9 += v9
                                                        _ = v27
                                                    } else {
                                                        v28 := GenS2{F0: (fi18 & -2147483648), F1: "foo bar"}
                                                        _ = v28
                                                    }
                                                    fmt.Println(v11, fi16)
                                                }
                                            case "foo bar":
                                                var v29 uint64 = 0
                                                _ = v29
                                            default:
                                                fmt.Println(v11, fi18, v11)
                                            }
                                            var v30 int32 = 1
                                            v31 := GenS5{F0: 2147483648, F1: v11, F2: (fi16 * (fi16 - fi16))}
                                            _ = v30
                                            _ = v31
                                        } else {
                                            var v32 int16 = (v9 & (v9 << ((42) & 63)))
                                            if true {
                                                fmt.Println(v32, v9)
                                            } else {
                                                if true {
                                                    fmt.Println(fi16, fi18, v11)
                                                    if false {
                                                        if (false || (!false)) {
                                                            var v33 float64 = 0.1
                                                            v9 += (^(v32))
                                                            v11 += (v11 / ((v11) | 1))
                                                            switch fi18 {
                                                            case -47:
                                                                fmt.Println(v11, fi18)
                                                            }
                                                            var v34 uint32 = 5
                                                            _ = v33
                                                            _ = v34
                                                        } else {
                                                            switch 0 {
                                                            case 456:
                                                                var v35 uint32 = 42
                                                                _ = v35
                                                            case 32767:
                                                                v36 := GenS1{F0: v11}
                                                                _ = v36
                                                            }
                                                            switch fi16 {
                                                            case -863:
                                                                fmt.Println(v9, fi18)
                                                            default:
                                                                if true {
                                                                    v2 = append(v2, 4294967295)
                                                                    var v37 float32 = 1e10
                                                                    if ("" != "foo bar") {
                                                                        fmt.Println(v11)
                                                                    } else {
                                                                        var v38 uint8 = 1
                                                                        _ = v38
                                                                    }
                                                                    _ = v37
                                                                } else {
                                                                    v11 += (v11 >> ((601) & 63))
                                                                    if (true || false) {
                                                                        switch ((((("a" + "a") + ("a" + ((("a" + "hello") + ("hello" + ((("" + "hello") + "a") + "\n\t"))) + ((("a" + "a") + "a") + "hello")))) + ("a" + ((("a" + (("foo bar" + "foo bar") + ("hello" + ""))) + "a") + "hello"))) + "hello") + "\n\t") {
                                                                        case "hello":
                                                                            v32 -= v9
                                                                        case "":
                                                                            v32 = v32
                                                                        case "a":
                                                                            v39 := GenS6{F0: int(v11), F1: 413, F2: 127}
                                                                            _ = v39
                                                                        default:
                                                                            var v40 int8 = 23
                                                                            _ = v40
                                                                        }
                                                                        if len(v2) > 0 {
                                                                            v2[0] = 4294967295
                                                                        }
                                                                        fmt.Println(fi18)
                                                                        v9 = int16(v11)
                                                                        v41 := GenS2{F0: 42, F1: (("foo bar" + "foo bar") + "foo bar")}
                                                                        _ = v41
                                                                    } else {
                                                                        v42 := GenS7{F0: 0}
                                                                        var v43 string = "\n\t"
                                                                        var v44 int = (fi16 + (fi16 + 2147483647))
                                                                        _ = v42
                                                                        _ = v43
                                                                        _ = v44
                                                                    }
                                                                }
                                                            }
                                                            v45 := GenS8{F0: 255, F1: v9}
                                                            _ = v45
                                                        }
                                                        var v46 int32 = 54
                                                        _ = v46
                                                    } else {
                                                        if len(v2) > 0 {
                                                            v2[0] = 65535
                                                        }
                                                        var v47 uint64 = 1
                                                        _ = v47
                                                    }
                                                    var v48 int32 = 0
                                                    var v49 bool = true
                                                    _ = v48
                                                    _ = v49
                                                } else {
                                                    v50 := GenS7{F0: -9223372036854775808}
                                                    switch "\n\t" {
                                                    case "a":
                                                        var v51 uint32 = 263
                                                        _ = v51
                                                    case "\n\t":
                                                        _ = len(v2)
                                                    }
                                                    _ = len(v1)
                                                    switch (fi18 << ((4294967295) & 63)) {
                                                    case 127:
                                                        if (18446744073709551615 < 499) {
                                                            var v52 uint8 = 128
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            _ = v52
                                                        }
                                                    default:
                                                        if false {
                                                            var v53 uint = 65535
                                                            _ = v53
                                                        } else {
                                                            v11 -= v11
                                                        }
                                                    }
                                                    _ = v50
                                                }
                                            }
                                            v9 = v32
                                            _ = v32
                                        }
                                    }
                                    if true {
                                        var v54 bool = (!false)
                                        var v55 float64 = float64(0.0)
                                        v9 += (-(v9))
                                        _ = v54
                                        _ = v55
                                    }
                                    var v56 float64 = -2.25
                                    _ = v56
                                }
                            case "\n\t":
                                v11 += v11
                            case "hello":
                                v9 = (-(v9))
                            }
                            v9 = (-(v9))
                        }
                        _ = len(v2)
                    }
                    _ = fi18
                    _ = v19
                }
                _ = fi16
                _ = v17
            }
            v9 += 629
            for fi57 := 0; fi57 < 2; fi57++ {
                switch "a" {
                case "":
                    for fi58 := 0; fi58 < 4; fi58++ {
                        if true {
                            v9 = (^(v9))
                            var v59 int64 = int64(v11)
                            _ = v59
                        }
                        if true {
                            if false {
                                fmt.Println(v11, v11, v11)
                            }
                            if (65535 != 18446744073709551615) {
                                v60 := v1["a"]
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                if (255 != 614) {
                                    v61 := GenS9{F0: 18446744073709551615}
                                    if true {
                                        v62 := GenS10{F0: (v9 >> ((18446744073709551615) & 63))}
                                        fmt.Println(fi57, v11, fi58)
                                        v63 := GenS6{F0: 32767, F1: 4294967295, F2: 143}
                                        _ = v62
                                        _ = v63
                                    }
                                    var v64 int64 = -9223372036854775808
                                    _ = v61
                                    _ = v64
                                }
                                _ = v60
                            } else {
                                if len(v2) > 0 {
                                    v2[0] = 2147483647
                                }
                                switch "hello" {
                                case "foo bar":
                                    var v65 int16 = v9
                                    _ = v65
                                case "":
                                    switch ("a" + "a") {
                                    case "foo bar":
                                        v11 = 100
                                    }
                                default:
                                    v1[(("" + "") + ("" + "\n\t"))] = (fi58 / (((fi58 % (((fi58 * (fi57 - fi57))) | 1))) | 1))
                                }
                                v9 = v9
                            }
                            v11 = v11
                            v11 -= v11
                            _ = len(v2)
                            v11 += (v11 * v11)
                        }
                        var v66 uint64 = 367
                        _ = fi58
                        _ = v66
                    }
                }
                v67 := GenS11{F0: 65535, F1: -2147483648, F2: -1}
                v9 += v9
                _ = fi57
                _ = v67
            }
        }
        v68 := GenS12{F0: ("foo bar" + "a")}
        _ = v68
    }
    for fi69 := 0; fi69 < 15; fi69++ {
        for fi70 := 0; fi70 < 19; fi70++ {
            switch 2147483647 {
            case -1:
                if (("foo bar" + "a") < "\n\t") {
                    switch fi70 {
                    case -987:
                        var v71 bool = (!(!(!(true || true))))
                        _ = v71
                    case 1:
                        if (!false) {
                            v11 += v11
                            if false {
                                v11 = (v11 + (v11 << ((4294967295) & 63)))
                                switch "foo bar" {
                                case "hello":
                                    if false {
                                        v9 -= v9
                                        var v72 bool = (16 >= (^(v11)))
                                        _ = v72
                                    }
                                case "":
                                    _ = len(v2)
                                case "\n\t":
                                    v11 = (v11 >> ((42) & 63))
                                }
                                var v73 int32 = -1
                                if (!false) {
                                    fmt.Println(v9, fi69)
                                    var v74 float64 = 1.5
                                    var v75 int16 = (v9 - (v9 ^ (-(v9))))
                                    _ = v74
                                    _ = v75
                                } else {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    v76 := GenS13{F0: 217, F1: (v9 + v9)}
                                    _ = v76
                                }
                                _ = v73
                            }
                            v77 := GenS14{F0: 42, F1: fi69}
                            _ = v77
                        }
                    case 4294967295:
                        fmt.Println(v11)
                    default:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                }
            case 42:
                var v78 int16 = -1
                _ = v78
            case 32768:
                var v79 int32 = -1
                _ = v79
            }
            if (!(1.5 == 0.0)) {
                v11 = (^(v11))
                v9 = (v9 & v9)
            }
            _ = fi70
        }
        var v80 uint16 = 827
        if ((255 > 773) && (1.5 < (1.5 + -2.25))) {
            for fi81 := 0; fi81 < 14; fi81++ {
                var v82 string = "foo bar"
                _ = fi81
                _ = v82
            }
            v3[((("" + "a") + "") + "hello")] = "foo bar"
        }
        _ = fi69
        _ = v80
    }
    if (42 > 227) {
        v83 := GenS15{F0: true, F1: v9}
        v84 := GenS13{F0: 18446744073709551615, F1: (-(v9))}
        _ = v83
        _ = v84
    }
    v9 -= v9
    for fi85 := 0; fi85 < 4; fi85++ {
        v86 := GenS2{F0: fi85, F1: "a"}
        v9 += 1000
        var v87 uint16 = 249
        _ = fi85
        _ = v86
        _ = v87
    }
    var v88 uint64 = 584
    fmt.Println(v9)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v88:", v88)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v88
    _ = v9
}
