package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint64; F1 int64 }
type GenS7 struct { F0 bool; F1 int16 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint8; F1 int }
type GenS10 struct { F0 uint8; F1 int8; F2 int }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 uint16 }

func main() {
    v1 := []string{("hello" + "hello"), "foo bar", "", "a"}
    var v2 uint32 = 783
    v3 := make(map[uint32]uint32)
    var v4 int64 = -509
    v4 -= v4
    v5 := GenS1{F0: 1, F1: 4294967295}
    var v6 int = 42
    v7 := GenS1{F0: 18446744073709551615, F1: 707}
    var v8 float64 = 0.1
    var v9 int64 = -9223372036854775808
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 uint8 = 46
        if (245 <= 0) {
            if (2 > 18446744073709551615) {
                var v12 int16 = 389
                _ = v12
            } else {
                var v13 string = (("foo bar" + "\n\t") + ("\n\t" + ("" + "hello")))
                if true {
                    switch fi10 {
                    case 302:
                        if false {
                            var v14 int32 = -540
                            _ = v14
                        } else {
                            for fi15 := 0; fi15 < 5; fi15++ {
                                var v16 uint = 0
                                _ = fi15
                                _ = v16
                            }
                            if (((true || false) && false) && (!true)) {
                                fmt.Println(v13)
                            }
                            var v17 uint8 = (^(v11))
                            _ = v17
                        }
                    default:
                        switch (v6 ^ 42) {
                        case 0:
                            var v18 string = "a"
                            _ = v18
                        default:
                            var v19 uint32 = (v2 & v2)
                            _ = v19
                        }
                    }
                    v4 = 1
                    switch fi10 {
                    case -9223372036854775808:
                        v2 += (v2 + (v2 | v2))
                    case -118:
                        v6 -= fi10
                    default:
                        if (v11 < v11) {
                            for fi20 := 0; fi20 < 16; fi20++ {
                                v3[(v2 & (v2 << ((255) & 63)))] = (^(v2))
                                if (((658 == 32767) || (1.5 == 0.1)) || false) {
                                    fmt.Println(v9, v9, v6)
                                    v9 -= 9223372036854775807
                                    if true {
                                        switch fi20 {
                                        case 752:
                                            v11 += 2
                                        case 9223372036854775807:
                                            if false {
                                                var v21 int8 = 1
                                                v13 = v13
                                                if false {
                                                    v22 := v3[4294967295]
                                                    v22 -= (^(v22))
                                                    _ = v22
                                                }
                                                _ = v21
                                            } else {
                                                switch 121 {
                                                case 2147483647:
                                                    fmt.Println(fi20, fi20, v9)
                                                case 128:
                                                    switch v6 {
                                                    case 42:
                                                        if (false || false) {
                                                            var v23 int8 = 100
                                                            _ = v23
                                                        }
                                                    case -2147483648:
                                                        switch (v13 + (v13 + (v13 + v13))) {
                                                        case "hello":
                                                            v2 += v2
                                                        default:
                                                            if (42 > 255) {
                                                                switch (v13 + v13) {
                                                                case "foo bar":
                                                                    var v24 uint32 = 1
                                                                    v4 -= v4
                                                                    v8 = float64(1.5)
                                                                    v25 := GenS1{F0: 65535, F1: 255}
                                                                    var v26 bool = false
                                                                    _ = v24
                                                                    _ = v25
                                                                    _ = v26
                                                                case "":
                                                                    switch "\n\t" {
                                                                    case "":
                                                                        switch fi10 {
                                                                        case 1:
                                                                            v27 := GenS2{F0: 32767, F1: v6}
                                                                            _ = v27
                                                                        case -371:
                                                                            v28 := GenS3{F0: (-(fi10)), F1: 963}
                                                                            _ = v28
                                                                        case 667:
                                                                            var v29 int8 = -1
                                                                            _ = v29
                                                                        }
                                                                    case "a":
                                                                        if false {
                                                                            v4 = int64(v11)
                                                                            fmt.Println(fi20, fi10)
                                                                            if false {
                                                                                var v30 uint32 = 42
                                                                                var v31 int = (-(v6))
                                                                                _ = v30
                                                                                _ = v31
                                                                            } else {
                                                                                if (364 <= 0) {
                                                                                    v8 = v8
                                                                                    v32 := GenS4{F0: 32767}
                                                                                    v2 -= (v2 + v2)
                                                                                    _ = v32
                                                                                } else {
                                                                                    v33 := GenS1{F0: 2, F1: 65535}
                                                                                    v13 = v13
                                                                                    _ = v33
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                                switch (^(fi20)) {
                                                                case -9223372036854775808:
                                                                    v8 += (3.14159 + v8)
                                                                case 4294967295:
                                                                    var v34 uint64 = 1
                                                                    _ = v34
                                                                case 127:
                                                                    var v35 uint = 58
                                                                    _ = v35
                                                                }
                                                                switch v13 {
                                                                case "a":
                                                                    v36 := GenS2{F0: 443, F1: (fi10 | -782)}
                                                                    _ = v36
                                                                case "":
                                                                    if (false && false) {
                                                                        var v37 bool = false
                                                                        delete(v3, v2)
                                                                        v38 := GenS5{F0: int16(v2)}
                                                                        v11 = v11
                                                                        v39 := GenS4{F0: 42}
                                                                        _ = v37
                                                                        _ = v38
                                                                        _ = v39
                                                                    }
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                default:
                                                                    if true {
                                                                        var v40 float32 = 1e10
                                                                        var v41 uint16 = 65535
                                                                        _ = v40
                                                                        _ = v41
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    }
                                                case -9223372036854775808:
                                                    delete(v3, v2)
                                                default:
                                                    var v42 uint = 4294967295
                                                    _ = v42
                                                }
                                            }
                                        }
                                        switch fi10 {
                                        case -9223372036854775808:
                                            if true {
                                                if ((("a" + (((v13 + (((v13 + v13) + v13) + "a")) + ((((v13 + (v13 + "\n\t")) + "foo bar") + ((v13 + v13) + v13)) + ((v13 + (v13 + "a")) + v13))) + v13)) + v13) < ((("\n\t" + ("hello" + v13)) + v13) + (((("\n\t" + (v13 + v13)) + ("" + ("" + v13))) + "a") + ""))) {
                                                    v43 := GenS6{F0: 4294967295, F1: (v9 << ((309) & 63))}
                                                    _ = len(v3)
                                                    _ = v43
                                                }
                                                v9 = (v9 * (^(v9)))
                                                v44 := GenS7{F0: true, F1: -1}
                                                _ = v44
                                            } else {
                                                if (!(false && false)) {
                                                    var v45 float64 = float64(((1e10 / 1.0) - 1e10))
                                                    var v46 string = ""
                                                    _ = v45
                                                    _ = v46
                                                }
                                                if (!true) {
                                                    fmt.Println(v6, v11, fi20)
                                                    v47 := GenS7{F0: false, F1: 335}
                                                    v4 -= (v9 >> ((1) & 63))
                                                    var v48 int32 = 2147483647
                                                    _ = v47
                                                    _ = v48
                                                } else {
                                                    if false {
                                                        fmt.Println(v4)
                                                        v1 = append(v1, "hello")
                                                    }
                                                    var v49 uint64 = 1
                                                    var v50 bool = false
                                                    v1 = append(v1, v13)
                                                    v51 := GenS5{F0: 0}
                                                    if (!(false || false)) {
                                                        switch (fi20 ^ (v6 + -136)) {
                                                        case 0:
                                                            var v52 int8 = 42
                                                            _ = v52
                                                        case 609:
                                                            _ = len(v3)
                                                        }
                                                        if (((1 != 211) || (v50 || (9223372036854775807 <= v9))) || (int(v11) != -499)) {
                                                            switch ("" + (v13 + v13)) {
                                                            case "hello":
                                                                var v53 float64 = v8
                                                                _ = v53
                                                            }
                                                        }
                                                        fmt.Println(v8)
                                                    } else {
                                                        if (4294967295 > 1) {
                                                            var v54 uint16 = 42
                                                            var v55 uint64 = (^(v49))
                                                            _ = v54
                                                            _ = v55
                                                        }
                                                        switch v13 {
                                                        case "foo bar":
                                                            switch "" {
                                                            case "a":
                                                                if ((v2 & (^(v2))) <= (^(v2))) {
                                                                    fmt.Println(v50, v11)
                                                                } else {
                                                                    fmt.Println(v50)
                                                                    v56 := GenS6{F0: v49, F1: (v4 >> ((517) & 63))}
                                                                    fmt.Println(v4, fi20, v49)
                                                                    var v57 uint32 = 13
                                                                    _ = v56
                                                                    _ = v57
                                                                }
                                                            case "foo bar":
                                                                v58 := GenS8{F0: 1000}
                                                                _ = v58
                                                            case "":
                                                                var v59 uint8 = uint8(v2)
                                                                _ = v59
                                                            default:
                                                                switch (v6 % (((^(fi10))) | 1)) {
                                                                case -1:
                                                                    var v60 int8 = -1
                                                                    _ = v60
                                                                case 32768:
                                                                    var v61 int32 = -645
                                                                    _ = v61
                                                                }
                                                            }
                                                        default:
                                                            v62 := GenS9{F0: (v11 ^ v11), F1: -1}
                                                            _ = v62
                                                        }
                                                        if (-34 != 133) {
                                                            v49 -= (v49 ^ v49)
                                                        }
                                                    }
                                                    _ = v49
                                                    _ = v50
                                                    _ = v51
                                                }
                                            }
                                        default:
                                            switch (fi20 >> ((4294967295) & 63)) {
                                            case 2147483647:
                                                fmt.Println(v11, v8, fi20)
                                            case 0:
                                                if (false || true) {
                                                    v2 += (v2 >> ((1) & 63))
                                                    var v63 int = v6
                                                    v2 = v2
                                                    _ = v63
                                                }
                                            case -2147483648:
                                                v6 -= int(v9)
                                            default:
                                                var v64 float64 = v8
                                                _ = v64
                                            }
                                        }
                                    } else {
                                        v6 -= v6
                                        v4 = v4
                                        _ = len(v1)
                                    }
                                } else {
                                    v65 := GenS10{F0: (^(v11)), F1: 1, F2: fi20}
                                    fmt.Println(v9)
                                    _ = v65
                                }
                                v8 = (float64(1e10) - 3.14159)
                                v4 = (^(v9))
                                _ = fi20
                            }
                        }
                    }
                } else {
                    for fi66 := 0; fi66 < 16; fi66++ {
                        v6 = int(v2)
                        _ = fi66
                    }
                }
                _ = v13
            }
        } else {
            v6 = (v6 >> ((4294967295) & 63))
            var v67 bool = (42 == -2147483648)
            for fi68 := 0; fi68 < 10; fi68++ {
                v69 := GenS7{F0: true, F1: 1}
                var v70 int16 = int16(v9)
                _ = fi68
                _ = v69
                _ = v70
            }
            _ = v67
        }
        switch "a" {
        case "":
            switch v6 {
            case 2147483647:
                var v71 float64 = (((float64(0.1) - v8) / v8) * float64(3.14159))
                _ = v71
            }
        case "hello":
            var v72 int32 = -1
            _ = v72
        default:
            _ = len(v3)
        }
        _ = fi10
        _ = v11
    }
    switch "\n\t" {
    case "\n\t":
        v73 := GenS11{F0: 18446744073709551615}
        _ = v73
    case "foo bar":
        for fi74 := 0; fi74 < 16; fi74++ {
            v4 = (v9 - (v9 - -1))
            if len(v1) > 0 {
                v1[0] = (((((("a" + ("" + ("" + ("foo bar" + "foo bar")))) + "hello") + "foo bar") + "") + "") + (("foo bar" + "a") + "hello"))
            }
            _ = fi74
        }
    case "hello":
        v75 := GenS12{F0: 42}
        _ = v75
    }
    var v76 int32 = 1
    var v77 int64 = (-(v4))
    if true {
        var v78 int = (v6 << ((1) & 63))
        var v79 int8 = 127
        _ = v78
        _ = v79
    }
    for fi80 := 0; fi80 < 14; fi80++ {
        for fi81 := 0; fi81 < 1; fi81++ {
            _ = len(v3)
            v2 -= 42
            fmt.Println(v6, v77, v9)
            _ = fi81
        }
        _ = fi80
    }
    v82 := GenS9{F0: 0, F1: v6}
    var v83 int = (v6 ^ v6)
    for fi84 := 0; fi84 < 10; fi84++ {
        var v85 uint64 = 1
        _ = fi84
        _ = v85
    }
    v77 -= (v4 >> ((2) & 63))
    v1 = append(v1, "a")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    fmt.Println("v8:", v8)
    fmt.Println("v83:", v83)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v76
    _ = v77
    _ = v8
    _ = v82
    _ = v83
    _ = v9
}
