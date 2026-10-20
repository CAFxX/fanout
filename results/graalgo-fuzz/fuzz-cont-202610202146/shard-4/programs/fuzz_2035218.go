package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 int8 }
func genFn0(p0 int, p1 bool, p2 int64) uint8 {
    p0 += -1
    v7 := []int{p0, 2147483647, (p0 + (-(p0)))}
    for fi8 := 0; fi8 < 10; fi8++ {
        v7 = append(v7, p0)
        var v9 int = (fi8 >> ((0) & 63))
        v10 := GenS1{F0: 235, F1: 2, F2: -943}
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 128
}
func genFn1() int {
    var v11 int8 = 0
    switch "foo bar" {
    case "hello":
        v11 = 100
    }
    _ = v11
    return 42
}

func main() {
    var v1 uint8 = 127
    var v2 bool = false
    v3 := []string{((((("hello" + "a") + "") + "") + "a") + ("hello" + "\n\t")), ("hello" + "foo bar")}
    var v4 bool = true
    v5 := []string{"", "hello"}
    var v6 uint16 = 1
    var v12 int32 = 0
    if (v2 && v2) {
        if false {
            var v13 float32 = (1.5 / 3.14159)
            _ = v13
        } else {
            if (!(!(!(!true)))) {
                for fi14 := 0; fi14 < 4; fi14++ {
                    var v15 float64 = 3.14159
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    for fi16 := 0; fi16 < 1; fi16++ {
                        fmt.Println(v12, v6)
                        _ = fi16
                    }
                    _ = fi14
                    _ = v15
                }
                v6 = v6
                v17 := GenS1{F0: (v1 / ((genFn0(genFn1(), (false || false), 968)) | 1)), F1: 67, F2: 1}
                _ = v17
            } else {
                switch ("a" + "a") {
                case "\n\t":
                    if (!v2) {
                        var v18 float64 = (0.1 - (1.5 * (1.5 * 3.14159)))
                        _ = v18
                    }
                case "":
                    var v19 int32 = 2147483647
                    _ = v19
                case "hello":
                    for fi20 := 0; fi20 < 8; fi20++ {
                        switch "a" {
                        case "foo bar":
                            if len(v5) > 0 {
                                v5[0] = "a"
                            }
                        case "":
                            var v21 uint16 = 70
                            _ = v21
                        default:
                            v2 = (!true)
                        }
                        v22 := GenS1{F0: (^(v1)), F1: 2, F2: -997}
                        _ = fi20
                        _ = v22
                    }
                }
            }
        }
        for fi23 := 0; fi23 < 19; fi23++ {
            _ = len(v3)
            v24 := GenS1{F0: (v1 - v1), F1: 0, F2: 1}
            v25 := GenS1{F0: (v1 / (((v1 | v1)) | 1)), F1: 0, F2: -764}
            _ = fi23
            _ = v24
            _ = v25
        }
    }
    switch -9223372036854775808 {
    case 1:
        v6 -= v6
    }
    var v26 int = 128
    var v27 bool = ((v2 && (42 < -9223372036854775808)) && (3.14159 >= (1e10 - ((0.1 * 1.5) + float32(-2.25)))))
    switch (v26 % ((v26) | 1)) {
    case 2147483647:
        if (v4 && v27) {
            var v28 int8 = 100
            _ = v28
        } else {
            var v29 uint32 = 0
            for fi30 := 0; fi30 < 4; fi30++ {
                if ((!(1 != 100)) || (v27 || (!v27))) {
                    if ((^(v29)) < 0) {
                        v31 := GenS2{F0: (fi30 ^ (fi30 % (((v26 << ((510) & 63))) | 1)))}
                        switch "hello" {
                        case "a":
                            v3 = append(v3, ("a" + "foo bar"))
                        case "\n\t":
                            _ = len(v3)
                        case "hello":
                            switch "\n\t" {
                            case "a":
                                v32 := GenS3{F0: v12, F1: 18446744073709551615}
                                _ = v32
                            case "\n\t":
                                v6 = (v6 % ((v6) | 1))
                            default:
                                for fi33 := 0; fi33 < 3; fi33++ {
                                    if (!((v29 - (^(v29))) <= (v29 & 65535))) {
                                        if len(v3) > 0 {
                                            v3[0] = "a"
                                        }
                                    }
                                    _ = fi33
                                }
                            }
                        default:
                            v34 := GenS3{F0: (^(v12)), F1: 257}
                            _ = v34
                        }
                        var v35 int = fi30
                        _ = v31
                        _ = v35
                    }
                } else {
                    for fi36 := 0; fi36 < 3; fi36++ {
                        switch "foo bar" {
                        case "":
                            if (((3.14159 - ((-2.25 / -2.25) / (-2.25 * -2.25))) > float64((-2.25 / -2.25))) || v27) {
                                v29 = (v29 + (^(v29)))
                            }
                        }
                        _ = fi36
                    }
                    var v37 uint = 160
                    _ = v37
                }
                _ = fi30
            }
            var v38 int64 = 1
            _ = v29
            _ = v38
        }
    case 32768:
        var v39 bool = (!((v6 | 0) < uint16(v1)))
        _ = v39
    }
    var v40 uint8 = (v1 ^ (v1 << ((42) & 63)))
    fmt.Println(v6, v4)
    var v41 string = "foo bar"
    var v42 float32 = -2.25
    if v4 {
        switch genFn1() {
        case -652:
            var v43 int64 = 1
            _ = v43
        case -9223372036854775808:
            if true {
                switch -9223372036854775808 {
                case 128:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                default:
                    v5 = append(v5, (("a" + ("hello" + "")) + v41))
                }
                v41 = (v41 + "foo bar")
                if false {
                    if len(v5) > 0 {
                        v5[0] = v41
                    }
                    v5 = append(v5, ("\n\t" + v41))
                    v44 := GenS1{F0: 0, F1: 869, F2: 491}
                    _ = v44
                }
            } else {
                for fi45 := 0; fi45 < 6; fi45++ {
                    for fi46 := 0; fi46 < 4; fi46++ {
                        v47 := GenS3{F0: v12, F1: 776}
                        if v2 {
                            if true {
                                var v48 string = v41
                                _ = v48
                            } else {
                                v40 -= (v40 + v40)
                            }
                            var v49 int8 = 107
                            _ = v49
                        }
                        _ = fi46
                        _ = v47
                    }
                    _ = fi45
                }
                var v50 bool = ((1e10 - (3.14159 * (-2.25 / 3.14159))) < 1e10)
                if ((v6 * (v6 % ((v6) | 1))) == (v6 >> ((439) & 63))) {
                    if ((((("hello" + "\n\t") + "") + v41) + v41) >= ((v41 + "") + (v41 + (v41 + ("hello" + ("a" + v41)))))) {
                        var v51 float64 = 0.1
                        v3 = append(v3, (v41 + v41))
                        _ = v51
                    }
                } else {
                    var v52 uint64 = 401
                    _ = v52
                }
                _ = v50
            }
        }
        if ((2147483647 < 344) && v27) {
            switch -318 {
            case -9223372036854775808:
                switch (-(v26)) {
                case 636:
                    switch ((((v41 + v41) + v41) + (v41 + (("\n\t" + ("\n\t" + ((v41 + (v41 + v41)) + (v41 + v41)))) + "a"))) + v41) {
                    case "a":
                        if (!((v26 + -77) == int(v12))) {
                            v53 := GenS4{F0: 827, F1: -327}
                            fmt.Println(v12, v40, v12)
                            _ = v53
                        }
                    case "\n\t":
                        v4 = true
                    }
                case 127:
                    for fi54 := 0; fi54 < 15; fi54++ {
                        switch "\n\t" {
                        case "hello":
                            fmt.Println(v26, fi54)
                        case "a":
                            v55 := GenS1{F0: v1, F1: 322, F2: -1}
                            _ = v55
                        default:
                            fmt.Println(v4, v27)
                        }
                        var v56 uint8 = (v40 >> ((4294967295) & 63))
                        _ = fi54
                        _ = v56
                    }
                }
            default:
                switch (v26 % (((v26 - v26)) | 1)) {
                case 156:
                    if v27 {
                        if false {
                            v26 = (v26 / ((v26) | 1))
                            v5 = append(v5, v41)
                        } else {
                            var v57 int64 = -1
                            fmt.Println(v12, v6)
                            switch v41 {
                            case "foo bar":
                                switch 32768 {
                                case 357:
                                    var v58 uint8 = 0
                                    _ = v58
                                case 0:
                                    v59 := GenS3{F0: 42, F1: 0}
                                    _ = v59
                                case -2147483648:
                                    var v60 uint32 = uint32(v40)
                                    _ = v60
                                }
                            case "hello":
                                var v61 uint32 = 624
                                _ = v61
                            }
                            _ = v57
                        }
                        for fi62 := 0; fi62 < 4; fi62++ {
                            var v63 int8 = -94
                            for fi64 := 0; fi64 < 16; fi64++ {
                                if (!(v41 < (("hello" + "\n\t") + "hello"))) {
                                    v3 = append(v3, "hello")
                                    var v65 float64 = 1e10
                                    v2 = false
                                    _ = v65
                                }
                                _ = fi64
                            }
                            _ = fi62
                            _ = v63
                        }
                    }
                }
            }
            var v66 int8 = 67
            _ = v66
        } else {
            _ = len(v3)
            if v27 {
                for fi67 := 0; fi67 < 10; fi67++ {
                    v2 = (!v27)
                    var v68 string = "foo bar"
                    fmt.Println(v4, v41, v26)
                    _ = fi67
                    _ = v68
                }
            } else {
                switch "a" {
                case "hello":
                    v4 = true
                case "":
                    if (((((1e10 - (1e10 * (((1e10 / -2.25) / 1.0) * float64(v42)))) - (1.5 * (-2.25 * 1e10))) * float64(v42)) - 1.5) == 3.14159) {
                        if ((0.0 / 1.5) <= 1e10) {
                            for fi69 := 0; fi69 < 18; fi69++ {
                                var v70 float64 = 1e10
                                _ = fi69
                                _ = v70
                            }
                        }
                    } else {
                        v71 := GenS5{F0: 20}
                        v41 = "foo bar"
                        _ = v71
                    }
                default:
                    if (v6 == (v6 - 65535)) {
                        var v72 int64 = -9223372036854775808
                        var v73 uint32 = uint32(v1)
                        var v74 uint8 = v40
                        _ = v72
                        _ = v73
                        _ = v74
                    }
                }
                v4 = ("a" != "")
                v40 = v1
            }
            var v75 uint32 = 230
            _ = v75
        }
        if false {
            v76 := GenS2{F0: genFn1()}
            v77 := GenS3{F0: -746, F1: 2}
            v1 -= genFn0(-1, v4, -1)
            _ = v76
            _ = v77
        }
    } else {
        fmt.Println(v41, v41, v42)
        switch v41 {
        case "hello":
            var v78 int32 = 809
            _ = v78
        case "\n\t":
            v79 := GenS1{F0: genFn0((v26 + v26), ((0 != 502) && v2), 1), F1: 18446744073709551615, F2: -1}
            _ = v79
        default:
            if v2 {
                _ = len(v5)
            } else {
                var v80 float32 = ((v42 * v42) * 1.5)
                var v81 uint = 913
                v12 = v12
                _ = v80
                _ = v81
            }
        }
    }
    v26 += (v26 & v26)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v27 = true
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
