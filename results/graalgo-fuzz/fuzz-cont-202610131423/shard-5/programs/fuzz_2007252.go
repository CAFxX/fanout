package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 int32; F2 uint8 }
type GenS4 struct { F0 uint; F1 uint8; F2 int16 }
type GenS5 struct { F0 string; F1 uint8; F2 bool }
type GenS6 struct { F0 bool; F1 uint; F2 int8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int8; F1 int8 }
type GenS9 struct { F0 int16; F1 int64; F2 int8 }
type GenS10 struct { F0 int64; F1 uint32 }
type GenS11 struct { F0 uint16; F1 string }
func genFn0(p0 int8, p1 int16, p2 bool) uint {
    v5 := []int{}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    p1 = 0
    for fi6 := 0; fi6 < 4; fi6++ {
        v7 := GenS1{F0: 2}
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 float32 = (3.14159 + 0.1)
            if p2 {
                v5 = append(v5, (fi6 - -9223372036854775808))
                var v10 string = ""
                _ = v10
            }
            if (65535 == 2147483648) {
                v11 := GenS2{F0: 42}
                var v12 int32 = 947
                switch (fi8 << ((851) & 63)) {
                case 32767:
                    v9 = v9
                case 128:
                    p0 = -28
                }
                var v13 uint64 = 2
                _ = v11
                _ = v12
                _ = v13
            } else {
                var v14 float64 = 1e10
                _ = v14
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v15 uint32 = 2147483648
            switch (("hello" + ((("a" + ("a" + "a")) + "hello") + "")) + "hello") {
            case "a":
                v16 := GenS3{F0: (^(v15)), F1: -820, F2: 96}
                _ = v16
            case "":
                p2 = (((100 <= (p0 >> ((255) & 63))) || p2) && (3.14159 < float64(v9)))
            case "hello":
                v17 := GenS2{F0: -2147483648}
                _ = v17
            }
            var v18 float32 = float32(1.5)
            _ = fi8
            _ = v15
            _ = v18
            _ = v9
        }
        if true {
            p0 = p0
        } else {
            if p2 {
                p1 = int16(p0)
                if ("\n\t" > "a") {
                    for fi19 := 0; fi19 < 18; fi19++ {
                        var v20 uint = 4294967295
                        var v21 int32 = 2147483647
                        if p2 {
                            if (598 < 28) {
                                if (11 <= 42) {
                                    fmt.Println(p0, v21, p0)
                                    if ((3.14159 / 0.1) <= (0.0 / 1e10)) {
                                        v22 := GenS4{F0: v20, F1: 127, F2: (p1 % (((p1 - (p1 % ((p1) | 1)))) | 1))}
                                        if (v20 >= 1) {
                                            if ((p1 >> ((v20) & 63)) > (^(p1))) {
                                                var v23 int64 = 1
                                                _ = v23
                                            }
                                            var v24 string = ("hello" + "")
                                            if len(v5) > 0 {
                                                v5[0] = 128
                                            }
                                            _ = v24
                                        }
                                        fmt.Println(p1, p1)
                                        switch "hello" {
                                        case "a":
                                            v25 := GenS5{F0: "", F1: 255, F2: (((^(v20)) != (v20 + (^(v20)))) || (3.14159 == float64(((float32(-2.25) / 1.0) * -2.25))))}
                                            _ = v25
                                        case "hello":
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        default:
                                            if (0 == 196) {
                                                v26 := GenS6{F0: true, F1: (^(v20)), F2: p0}
                                                _ = len(v5)
                                                v27 := GenS7{F0: 2147483648}
                                                _ = v26
                                                _ = v27
                                            }
                                        }
                                        _ = v22
                                    } else {
                                        fmt.Println(p2, fi19)
                                        var v28 int = (fi19 << (((v20 & (^(v20)))) & 63))
                                        v29 := GenS8{F0: (p0 ^ p0), F1: p0}
                                        _ = v28
                                        _ = v29
                                    }
                                    var v30 int8 = 100
                                    p0 += (-(p0))
                                    fmt.Println(v21, v20, fi6)
                                    _ = v30
                                } else {
                                    p2 = ((1e10 <= (0.1 + 3.14159)) && true)
                                    if true {
                                        var v31 float32 = 3.14159
                                        var v32 int32 = v21
                                        switch (fi19 >> ((v20) & 63)) {
                                        case 134:
                                            var v33 int64 = 0
                                            _ = v33
                                        case 2147483647:
                                            fmt.Println(p2, fi6)
                                        default:
                                            var v34 uint = (v20 >> ((42) & 63))
                                            _ = v34
                                        }
                                        _ = v31
                                        _ = v32
                                    }
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                }
                                v35 := GenS5{F0: ((("a" + "a") + ("" + "foo bar")) + ("a" + ("hello" + (("\n\t" + "\n\t") + "foo bar")))), F1: 216, F2: true}
                                _ = len(v5)
                                _ = v35
                            }
                        } else {
                            v21 += v21
                            var v36 int = int(v21)
                            p2 = (!p2)
                            _ = v36
                        }
                        switch 127 {
                        case 0:
                            switch (fi19 << (((v20 << ((65535) & 63))) & 63)) {
                            case 2147483647:
                                _ = len(v5)
                            case -9223372036854775808:
                                p0 -= p0
                            case -1:
                                if true {
                                    switch -2147483648 {
                                    case 9223372036854775807:
                                        switch fi19 {
                                        case 4294967295:
                                            var v37 uint8 = 1
                                            _ = v37
                                        case 42:
                                            switch 1 {
                                            case 128:
                                                p2 = p2
                                            }
                                            _ = len(v5)
                                            switch -1 {
                                            case 626:
                                                v20 = v20
                                            case 4294967295:
                                                fmt.Println(v20)
                                            case 32768:
                                                var v38 int64 = 1
                                                if len(v5) > 0 {
                                                    v5[0] = fi19
                                                }
                                                p0 = p0
                                                if (42 != 18446744073709551615) {
                                                    v5 = append(v5, (fi6 << ((v20) & 63)))
                                                    switch "foo bar" {
                                                    case "a":
                                                        v5 = append(v5, fi6)
                                                    case "hello":
                                                        var v39 uint32 = 476
                                                        _ = v39
                                                    case "\n\t":
                                                        if (!p2) {
                                                            fmt.Println(v38, v21)
                                                            v21 -= (v21 / ((v21) | 1))
                                                            v40 := GenS1{F0: 683}
                                                            switch 1 {
                                                            case -197:
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                            case 32768:
                                                                _ = len(v5)
                                                                p2 = p2
                                                                if (true && ((p2 && (v20 >= 0)) || ((p1 == (-(p1))) && (773 > 4294967295)))) {
                                                                    var v41 float64 = 1e10
                                                                    switch "hello" {
                                                                    case "foo bar":
                                                                        v38 = 42
                                                                    case "hello":
                                                                        p0 = -1
                                                                    }
                                                                    _ = v41
                                                                }
                                                            default:
                                                                if p2 {
                                                                    if p2 {
                                                                        fmt.Println(v20, fi19)
                                                                    } else {
                                                                        if len(v5) >= 2 {
                                                                            v5 = v5[1:len(v5)]
                                                                        }
                                                                    }
                                                                    v42 := GenS4{F0: v20, F1: 42, F2: (^(p1))}
                                                                    var v43 uint64 = 18446744073709551615
                                                                    var v44 uint = v20
                                                                    switch (fi19 | fi19) {
                                                                    case -164:
                                                                        if (("a" + "foo bar") <= "a") {
                                                                            if (127 == p0) {
                                                                                _ = len(v5)
                                                                            }
                                                                        } else {
                                                                            if (!((p0 ^ p0) <= p0)) {
                                                                                v43 -= (v43 | (v43 - (v43 ^ 18446744073709551615)))
                                                                                v43 -= v43
                                                                                v21 = v21
                                                                            }
                                                                            if p2 {
                                                                                v45 := GenS3{F0: 65535, F1: (^(v21)), F2: 128}
                                                                                _ = v45
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v42
                                                                    _ = v43
                                                                    _ = v44
                                                                } else {
                                                                    v46 := GenS4{F0: (v20 | (v20 + v20)), F1: 23, F2: (p1 | p1)}
                                                                    v21 += (-(v21))
                                                                    _ = v46
                                                                }
                                                            }
                                                            _ = v40
                                                        }
                                                    default:
                                                        if ((1.5 * 0.0) == 0.1) {
                                                            var v47 float64 = 1.5
                                                            _ = v47
                                                        } else {
                                                            fmt.Println(v21)
                                                        }
                                                    }
                                                    if (int16(fi6) < p1) {
                                                        p2 = p2
                                                        switch "a" {
                                                        case "":
                                                            if (false && p2) {
                                                                _ = len(v5)
                                                                var v48 bool = (((("\n\t" + "hello") + (("foo bar" + "") + ("hello" + "hello"))) + "a") == "hello")
                                                                var v49 uint8 = 127
                                                                _ = v48
                                                                _ = v49
                                                            }
                                                        case "\n\t":
                                                            v50 := GenS3{F0: 4294967295, F1: 1, F2: 42}
                                                            _ = v50
                                                        }
                                                    } else {
                                                        var v51 float32 = 0.0
                                                        p0 = p0
                                                        var v52 int32 = -2147483648
                                                        fmt.Println(v21, fi6, p2)
                                                        fmt.Println(v51)
                                                        _ = v51
                                                        _ = v52
                                                    }
                                                    p1 += 1000
                                                }
                                                v53 := GenS3{F0: 65535, F1: v21, F2: 1}
                                                _ = v38
                                                _ = v53
                                            }
                                        case -1:
                                            v54 := GenS9{F0: p1, F1: 42, F2: p0}
                                            var v55 string = "\n\t"
                                            var v56 string = v55
                                            _ = v54
                                            _ = v55
                                            _ = v56
                                        }
                                    case 127:
                                        p0 += 86
                                        switch ("a" + "hello") {
                                        case "foo bar":
                                            v57 := GenS7{F0: 65535}
                                            _ = v57
                                        }
                                        if p2 {
                                            p0 += p0
                                        }
                                    case 921:
                                        p0 -= (p0 % ((p0) | 1))
                                    }
                                    fmt.Println(p2, p0)
                                }
                            }
                        default:
                            if (true || (0.1 == (1.5 / 1.5))) {
                                if len(v5) > 0 {
                                    v5[0] = fi19
                                }
                            }
                        }
                        fmt.Println(v21, p1)
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                } else {
                    if len(v5) > 0 {
                        v5[0] = fi6
                    }
                    var v58 uint = 1
                    var v59 int = (fi6 + fi6)
                    _ = v58
                    _ = v59
                }
            } else {
                _ = len(v5)
            }
        }
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 0
}
func genFn1(p0 uint32, p1 uint64) int32 {
    for fi60 := 0; fi60 < 8; fi60++ {
        if (false || (!((false || (-1 != 0)) || false))) {
            switch (fi60 - int(p1)) {
            case 42:
                v61 := []int{fi60}
                _ = v61
            }
        }
        _ = fi60
    }
    v62 := []int{-9223372036854775808}
    v63 := GenS7{F0: p0}
    _ = p0
    _ = p1
    _ = v62
    _ = v63
    return 2147483647
}

func main() {
    var v1 uint16 = 444
    v2 := []uint32{822, 712}
    var v3 string = "foo bar"
    var v4 uint64 = 18446744073709551615
    v4 += (v4 >> ((1) & 63))
    var v64 int16 = -32768
    _ = len(v2)
    if true {
        if (!true) {
            fmt.Println(v4, v3)
            var v65 uint32 = 2147483648
            for fi66 := 0; fi66 < 1; fi66++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                for fi67 := 0; fi67 < 10; fi67++ {
                    var v68 uint32 = v65
                    if ((((1.5 - 0.0) - 0.0) + (0.1 - (float64(((((1.5 - 0.1) - 1e10) / 1e10) / 1.0)) * 3.14159))) < 0.1) {
                        v69 := GenS10{F0: 9223372036854775807, F1: 4294967295}
                        switch v3 {
                        case "\n\t":
                            v2 = append(v2, (v68 >> ((42) & 63)))
                        case "a":
                            v3 = "hello"
                        case "":
                            v68 += (v68 * v68)
                            if (false || true) {
                                v70 := GenS9{F0: (v64 % ((v64) | 1)), F1: -1, F2: 100}
                                var v71 int8 = -128
                                if ((1 == 100) || false) {
                                    if false {
                                        v4 -= (^(v4))
                                    } else {
                                        var v72 int8 = -1
                                        _ = len(v2)
                                        _ = v72
                                    }
                                    if false {
                                        v65 = v68
                                    }
                                }
                                _ = v70
                                _ = v71
                            }
                            var v73 uint64 = 541
                            var v74 int = -279
                            _ = v73
                            _ = v74
                        default:
                            fmt.Println(v1, fi66, v1)
                        }
                        switch (fi67 + (fi67 + 0)) {
                        case 693:
                            v4 += v4
                        case 1:
                            v75 := GenS11{F0: v1, F1: v3}
                            _ = v75
                        default:
                            if false {
                                var v76 uint64 = v4
                                var v77 float64 = ((float64((float32(0.1) * 1e10)) * 3.14159) + 3.14159)
                                fmt.Println(v77, v68)
                                _ = v76
                                _ = v77
                            } else {
                                var v78 uint32 = v65
                                if len(v2) > 0 {
                                    v2[0] = v68
                                }
                                _ = v78
                            }
                        }
                        _ = v69
                    } else {
                        var v79 uint = genFn0(127, (v64 / (((v64 * v64)) | 1)), (!(0 <= -1)))
                        _ = v79
                    }
                    _ = len(v2)
                    _ = fi67
                    _ = v68
                }
                _ = fi66
            }
            _ = v65
        } else {
            var v80 int32 = -2147483648
            _ = v80
        }
        var v81 string = "hello"
        var v82 uint32 = 2
        _ = v81
        _ = v82
    } else {
        var v83 float64 = 0.0
        var v84 uint64 = (v4 | v4)
        for fi85 := 0; fi85 < 12; fi85++ {
            fmt.Println(v4, v64)
            var v86 float64 = ((v83 + (v83 * v83)) * (v83 / float64((0.1 + float32(v83)))))
            v84 -= v4
            _ = fi85
            _ = v86
        }
        _ = v83
        _ = v84
    }
    v2 = append(v2, 63)
    if false {
        for fi87 := 0; fi87 < 13; fi87++ {
            for fi88 := 0; fi88 < 18; fi88++ {
                v64 += (^(v64))
                fmt.Println(v1, fi87, v4)
                _ = fi88
            }
            _ = fi87
        }
        v1 = (^(v1))
        fmt.Println(v4)
    }
    var v89 uint32 = 400
    v1 += v1
    v64 -= (v64 - v64)
    var v90 uint8 = 2
    switch ("" + (((v3 + v3) + "\n\t") + "a")) {
    case "foo bar":
        v3 = ("" + (v3 + v3))
    case "a":
        fmt.Println(v90, v89, v90)
    case "":
        var v91 int32 = 0
        _ = v91
    default:
        fmt.Println(v89, v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v64:", v64)
    fmt.Println("v89:", v89)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v64
    _ = v89
    _ = v90
}
