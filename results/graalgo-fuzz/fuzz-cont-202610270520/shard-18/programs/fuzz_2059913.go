package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 int32 }
type GenS2 struct { F0 int32; F1 uint16; F2 bool }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint16; F1 int8; F2 int16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint8 }
func genFn0() string {
    if (251 > 255) {
        if false {
            var v4 int16 = -923
            _ = v4
        } else {
            var v5 uint = 65535
            var v6 uint8 = 127
            v7 := []int{}
            _ = v5
            _ = v6
            _ = v7
        }
    }
    for fi8 := 0; fi8 < 5; fi8++ {
        fmt.Println(fi8)
        fmt.Println(fi8)
        if ((true && ((!false) || true)) && false) {
            var v9 bool = true
            _ = v9
        } else {
            v10 := GenS1{F0: 1, F1: 32767, F2: 1}
            var v11 int16 = -1
            v11 -= v11
            _ = v10
            _ = v11
        }
        _ = fi8
    }
    var v12 string = (("\n\t" + "a") + "")
    v13 := GenS1{F0: 1, F1: -32768, F2: -2147483648}
    _ = v12
    _ = v13
    return ("foo bar" + ("\n\t" + ("\n\t" + v12)))
}
func genFn1(p0 uint32, p1 uint16) int16 {
    v14 := GenS2{F0: -1, F1: p1, F2: (127 > -1)}
    fmt.Println(p1, p1)
    for fi15 := 0; fi15 < 7; fi15++ {
        p0 += (p0 / ((p0) | 1))
        var v16 uint32 = 2147483648
        v17 := GenS2{F0: -1, F1: (p1 & 42), F2: true}
        _ = fi15
        _ = v16
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v14
    return -32768
}

func main() {
    var v1 int16 = 42
    var v2 int16 = v1
    v3 := []string{(((("foo bar" + "hello") + ("a" + ("\n\t" + ""))) + "a") + "hello"), "a"}
    _ = len(v3)
    var v18 float64 = (3.14159 * ((3.14159 / 3.14159) - -2.25))
    var v19 int8 = -63
    var v20 bool = (225 != 127)
    for fi21 := 0; fi21 < 7; fi21++ {
        v22 := GenS1{F0: -930, F1: v1, F2: 42}
        _ = fi21
        _ = v22
    }
    switch (("hello" + ((("foo bar" + (("" + ("foo bar" + "hello")) + ("foo bar" + ""))) + ((((("a" + "hello") + "hello") + ("a" + "foo bar")) + "\n\t") + ("\n\t" + "foo bar"))) + "a")) + "") {
    case "\n\t":
        v20 = ((42 < -1) && v20)
    case "foo bar":
        _ = len(v3)
    case "a":
        var v23 uint8 = 0
        _ = v23
    default:
        if true {
            var v24 int16 = -366
            v18 = v18
            if (v20 && true) {
                fmt.Println(v18, v20)
            }
            _ = v24
        }
    }
    if (v20 || (true && v20)) {
        switch -906 {
        case -208:
            if (2 != 95) {
                for fi25 := 0; fi25 < 9; fi25++ {
                    switch "" {
                    case "foo bar":
                        v2 -= (v2 * 32767)
                    default:
                        if (!v20) {
                            for fi26 := 0; fi26 < 5; fi26++ {
                                var v27 float64 = float64(0.1)
                                if (858 != 511) {
                                    var v28 uint16 = 306
                                    fmt.Println(v20, v1)
                                    v27 += 0.1
                                    _ = len(v3)
                                    _ = v28
                                }
                                var v29 bool = false
                                _ = fi26
                                _ = v27
                                _ = v29
                            }
                            v2 = v1
                        }
                    }
                    _ = fi25
                }
                switch 127 {
                case -416:
                    v20 = v20
                case 42:
                    v19 = v19
                }
            }
        case 875:
            var v30 int16 = 1000
            _ = v30
        case -9223372036854775808:
            v31 := GenS3{F0: 942}
            _ = v31
        default:
            v3 = append(v3, ("\n\t" + ("foo bar" + "")))
        }
        v32 := GenS1{F0: 42, F1: int16(v19), F2: 755}
        _ = v32
    }
    switch ("foo bar" + "a") {
    case "":
        if v20 {
            v33 := GenS4{F0: 65535, F1: 127, F2: v1}
            v19 += v19
            for fi34 := 0; fi34 < 15; fi34++ {
                switch fi34 {
                case 32768:
                    if v20 {
                        v2 += (v2 | genFn1(106, 32767))
                        for fi35 := 0; fi35 < 16; fi35++ {
                            fmt.Println(v18, v2, fi35)
                            _ = fi35
                        }
                    }
                case 2147483647:
                    for fi36 := 0; fi36 < 20; fi36++ {
                        if true {
                            v1 = v2
                        } else {
                            fmt.Println(v18, fi34)
                            switch (fi34 % (((fi36 * -1)) | 1)) {
                            case -1:
                                switch "foo bar" {
                                case "foo bar":
                                    v1 += (v2 / (((v2 / (((v2 / ((v2) | 1))) | 1))) | 1))
                                }
                                v37 := GenS3{F0: (fi34 ^ 32768)}
                                v2 = v1
                                switch 32767 {
                                case -2147483648:
                                    switch "" {
                                    case "a":
                                        var v38 int16 = v1
                                        _ = v38
                                    default:
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                case 42:
                                    var v39 uint16 = 687
                                    _ = v39
                                default:
                                    if true {
                                        if true {
                                            if len(v3) > 0 {
                                                v3[0] = "foo bar"
                                            }
                                            switch "foo bar" {
                                            case "":
                                                var v40 float32 = (float32(v18) / 1.0)
                                                _ = v40
                                            }
                                            fmt.Println(v2, v19)
                                        } else {
                                            var v41 float32 = (1.5 * 1.5)
                                            if v20 {
                                                fmt.Println(fi36, v1, v20)
                                            } else {
                                                if (v20 || (true || true)) {
                                                    if v20 {
                                                        if (!v20) {
                                                            if ((0 > 721) && v20) {
                                                                _ = len(v3)
                                                                var v42 uint32 = 1
                                                                var v43 float64 = v18
                                                                _ = v42
                                                                _ = v43
                                                            } else {
                                                                if true {
                                                                    v19 -= -128
                                                                    var v44 float64 = (3.14159 - float64(v41))
                                                                    _ = v44
                                                                }
                                                                v41 -= v41
                                                                var v45 uint64 = 343
                                                                if ((128 == fi36) && (!v20)) {
                                                                    fmt.Println(v41)
                                                                    v41 = v41
                                                                    _ = len(v3)
                                                                } else {
                                                                    var v46 uint16 = 248
                                                                    v45 += v45
                                                                    _ = v46
                                                                }
                                                                var v47 string = "a"
                                                                _ = v45
                                                                _ = v47
                                                            }
                                                            switch "\n\t" {
                                                            case "a":
                                                                if len(v3) > 0 {
                                                                    v3[0] = "\n\t"
                                                                }
                                                            case "\n\t":
                                                                var v48 float32 = v41
                                                                _ = v48
                                                            default:
                                                                var v49 int32 = 2147483647
                                                                _ = v49
                                                            }
                                                        } else {
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            fmt.Println(fi36, v19)
                                                            fmt.Println(fi34)
                                                        }
                                                    } else {
                                                        switch ("\n\t" + "hello") {
                                                        case "":
                                                            fmt.Println(v20, v2)
                                                        case "foo bar":
                                                            var v50 float64 = v18
                                                            _ = v50
                                                        default:
                                                            v51 := GenS4{F0: 40, F1: (v19 >> ((255) & 63)), F2: -32768}
                                                            _ = v51
                                                        }
                                                    }
                                                    v20 = (-9223372036854775808 == 0)
                                                }
                                                var v52 uint32 = 90
                                                var v53 float64 = (3.14159 / ((((v18 / 1e10) + (v18 * v18)) + v18) - ((1.5 - float64((v41 / float32((0.1 - (3.14159 - v18)))))) * 1.5)))
                                                _ = v52
                                                _ = v53
                                            }
                                            _ = v41
                                        }
                                        var v54 bool = (v19 == int8(v1))
                                        if true {
                                            var v55 string = "foo bar"
                                            _ = v55
                                        }
                                        _ = v54
                                    } else {
                                        switch "foo bar" {
                                        case "hello":
                                            switch (("" + "a") + "\n\t") {
                                            case "a":
                                                if (!v20) {
                                                    var v56 uint32 = 4294967295
                                                    fmt.Println(v20, v19)
                                                    _ = v56
                                                } else {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    v20 = (true || (false && (((628 > 598) || (!v20)) || true)))
                                                    var v57 bool = v20
                                                    fmt.Println(v1, fi34)
                                                    var v58 float32 = float32((0.1 / v18))
                                                    v19 = v19
                                                    _ = v57
                                                    _ = v58
                                                }
                                            case "":
                                                switch "foo bar" {
                                                case "":
                                                    var v59 string = ("foo bar" + "hello")
                                                    _ = v59
                                                case "a":
                                                    if ((255 >= 65535) || (127 <= 128)) {
                                                        var v60 int = -2147483648
                                                        var v61 int8 = v19
                                                        _ = v60
                                                        _ = v61
                                                    } else {
                                                        fmt.Println(v2, v18, v1)
                                                    }
                                                case "foo bar":
                                                    v62 := GenS5{F0: (v19 / ((v19) | 1))}
                                                    _ = v62
                                                }
                                            }
                                        }
                                        if true {
                                            var v63 float64 = v18
                                            if v20 {
                                                if true {
                                                    fmt.Println(v19, v63)
                                                    var v64 int32 = 761
                                                    _ = v64
                                                }
                                                if (!(v20 || v20)) {
                                                    switch fi34 {
                                                    case 42:
                                                        var v65 string = "foo bar"
                                                        _ = v65
                                                    }
                                                    switch fi36 {
                                                    case -668:
                                                        v63 = v18
                                                    default:
                                                        if (592 != 1) {
                                                            switch fi34 {
                                                            case 1:
                                                                v63 = (-2.25 - ((v18 - v63) + -2.25))
                                                            }
                                                            v66 := GenS2{F0: 42, F1: 255, F2: true}
                                                            _ = v66
                                                        }
                                                    }
                                                    v20 = (!(0 >= -858))
                                                } else {
                                                    fmt.Println(v19)
                                                    v67 := GenS6{F0: 232}
                                                    _ = len(v3)
                                                    _ = v67
                                                }
                                            }
                                            _ = v63
                                        }
                                        fmt.Println(fi36, fi36)
                                    }
                                }
                                if false {
                                    v19 = 0
                                    switch (fi36 % ((-606) | 1)) {
                                    case 32768:
                                        var v68 int64 = 1
                                        _ = v68
                                    default:
                                        fmt.Println(v2)
                                    }
                                    fmt.Println(v18, v18, v18)
                                } else {
                                    var v69 int32 = -2147483648
                                    v70 := GenS1{F0: v69, F1: (v2 + (v2 | (v2 / ((v2) | 1)))), F2: (v69 + (v69 >> ((805) & 63)))}
                                    _ = v69
                                    _ = v70
                                }
                                var v71 uint = 18446744073709551615
                                _ = v37
                                _ = v71
                            case 9223372036854775807:
                                if true {
                                    fmt.Println(v20, fi36, v1)
                                    if (true || (fi34 != fi34)) {
                                        fmt.Println(v2, v18, fi34)
                                    } else {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        var v72 int16 = int16(v19)
                                        v2 = (v72 & 1)
                                        _ = v72
                                    }
                                    switch "" {
                                    case "hello":
                                        var v73 int = fi36
                                        _ = v73
                                    case "":
                                        fmt.Println(v18, v1, fi36)
                                    case "foo bar":
                                        _ = len(v3)
                                    default:
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                    _ = len(v3)
                                    if ((-128 > (v19 / ((1) | 1))) && (((13 != 2) || false) && (2 <= 0))) {
                                        switch "a" {
                                        case "":
                                            var v74 bool = v20
                                            _ = v74
                                        case "\n\t":
                                            var v75 float64 = (v18 / v18)
                                            _ = v75
                                        case "a":
                                            v76 := GenS3{F0: (-(fi34))}
                                            _ = v76
                                        default:
                                            v18 = ((v18 - (v18 + 3.14159)) + (1e10 + (float64(3.14159) - (float64(((1e10 + 1.5) + 0.0)) - (v18 + v18)))))
                                        }
                                        switch (("a" + ("a" + ("a" + "a"))) + "") {
                                        case "foo bar":
                                            var v77 uint32 = 4294967295
                                            _ = v77
                                        default:
                                            if true {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                var v78 bool = (712 == 2)
                                                _ = v78
                                            } else {
                                                fmt.Println(v1, v20)
                                            }
                                        }
                                    }
                                }
                            }
                            v19 = (v19 * v19)
                        }
                        _ = fi36
                    }
                }
                var v79 int8 = (v19 - (v19 % ((v19) | 1)))
                _ = fi34
                _ = v79
            }
            _ = v33
        }
    case "a":
        fmt.Println(v19, v1)
    case "hello":
        var v80 int = 128
        _ = v80
    }
    var v81 int16 = (v1 + genFn1(42, 1))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v81
}
