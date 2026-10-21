package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 int64; F1 uint; F2 uint16 }
type GenS5 struct { F0 uint; F1 uint8; F2 uint }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int16; F1 string }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int32; F1 uint32; F2 uint64 }
type GenS11 struct { F0 int; F1 int64; F2 int32 }
type GenS12 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{"foo bar", (("a" + "foo bar") + "foo bar"), "foo bar", "a"}
    var v3 bool = ((32767 > 0) || false)
    v3 = false
    fmt.Println(v3)
    var v4 bool = false
    var v5 int32 = -2147483648
    v3 = ((!false) || false)
    fmt.Println(v5, v4, v5)
    switch 32767 {
    case -1:
        for fi6 := 0; fi6 < 3; fi6++ {
            fmt.Println(v5)
            var v7 int64 = -9223372036854775808
            for fi8 := 0; fi8 < 11; fi8++ {
                if (v3 || (true && (!(3.14159 >= (1.5 + 1e10))))) {
                    fmt.Println(v5)
                    if true {
                        if (0.0 > (3.14159 * float32(0.1))) {
                            if (65535 > 1) {
                                switch 0 {
                                case 0:
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                case 4294967295:
                                    v9 := GenS1{F0: (v4 || false), F1: ("" + ("foo bar" + ((("a" + "foo bar") + "hello") + ("foo bar" + "foo bar"))))}
                                    _ = v9
                                case -1:
                                    fmt.Println(fi8)
                                default:
                                    fmt.Println(v4, fi8)
                                }
                                v10 := GenS1{F0: (v3 && true), F1: "a"}
                                if ("\n\t" <= "hello") {
                                    if (42 == -32768) {
                                        var v11 uint = 42
                                        var v12 bool = (3.14159 <= ((1e10 * 1e10) - 3.14159))
                                        fmt.Println(fi6, v5, v12)
                                        v13 := GenS2{F0: (v7 % ((int64(fi6)) | 1))}
                                        switch "a" {
                                        case "foo bar":
                                            if len(v2) > 0 {
                                                v2[0] = ""
                                            }
                                        case "hello":
                                            var v14 uint32 = 430
                                            _ = v14
                                        }
                                        var v15 float32 = float32(1.5)
                                        v4 = v12
                                        v16 := GenS3{F0: -80, F1: 42}
                                        _ = v11
                                        _ = v12
                                        _ = v13
                                        _ = v15
                                        _ = v16
                                    }
                                    var v17 int64 = 1
                                    var v18 uint8 = 2
                                    v2 = append(v2, "a")
                                    _ = v17
                                    _ = v18
                                } else {
                                    fmt.Println(fi6)
                                    v1[42] = (^(fi6))
                                    var v19 int = fi8
                                    _ = v19
                                }
                                var v20 float32 = 1e10
                                _ = v10
                                _ = v20
                            }
                            switch "hello" {
                            case "a":
                                var v21 int8 = int8(v5)
                                _ = v21
                            case "":
                                if len(v2) > 0 {
                                    v2[0] = "foo bar"
                                }
                            default:
                                switch "a" {
                                case "\n\t":
                                    v7 = 9223372036854775807
                                case "a":
                                    v3 = ((1e10 * (0.0 * 0.0)) >= 1e10)
                                default:
                                    v22 := GenS1{F0: v3, F1: ("foo bar" + ("foo bar" + ((("foo bar" + "") + "\n\t") + "hello")))}
                                    var v23 uint32 = 2147483648
                                    v4 = true
                                    _ = v22
                                    _ = v23
                                }
                            }
                            if true {
                                var v24 int32 = -506
                                var v25 uint8 = 223
                                _ = v24
                                _ = v25
                            } else {
                                fmt.Println(v3, fi8)
                                v26 := GenS2{F0: (v7 << ((65535) & 63))}
                                _ = v26
                            }
                        } else {
                            var v27 uint8 = 105
                            if v3 {
                                v2 = append(v2, "\n\t")
                                v28 := GenS4{F0: (v7 >> ((65535) & 63)), F1: 18446744073709551615, F2: 1}
                                v4 = v4
                                _ = v28
                            } else {
                                if false {
                                    switch "\n\t" {
                                    case "hello":
                                        switch "" {
                                        case "foo bar":
                                            var v29 float32 = (0.0 - 1e10)
                                            _ = v29
                                        case "\n\t":
                                            var v30 string = "hello"
                                            _ = v30
                                        default:
                                            v31 := GenS1{F0: v3, F1: "a"}
                                            _ = v31
                                        }
                                    case "a":
                                        v27 = v27
                                    }
                                }
                            }
                            var v32 string = (("hello" + (((("hello" + "") + "hello") + "a") + "foo bar")) + "\n\t")
                            var v33 uint64 = 18446744073709551615
                            _ = v27
                            _ = v32
                            _ = v33
                        }
                        v4 = false
                        var v34 string = ("" + ("hello" + "\n\t"))
                        _ = v34
                    }
                    v5 = v5
                    var v35 float64 = -2.25
                    if v3 {
                        var v36 int8 = 100
                        v37 := GenS4{F0: v7, F1: 255, F2: 62}
                        v38 := GenS5{F0: 255, F1: 128, F2: 255}
                        _ = v36
                        _ = v37
                        _ = v38
                    }
                    var v39 uint16 = 1
                    var v40 int = fi6
                    switch ("\n\t" + ("hello" + "foo bar")) {
                    case "":
                        switch "foo bar" {
                        case "a":
                            var v41 bool = (42 != -32768)
                            _ = v41
                        }
                    case "a":
                        var v42 int8 = 21
                        _ = v42
                    case "hello":
                        var v43 int = (fi8 / (((fi6 | 0)) | 1))
                        _ = v43
                    }
                    var v44 uint = 4294967295
                    _ = v35
                    _ = v39
                    _ = v40
                    _ = v44
                } else {
                    v4 = false
                }
                _ = fi8
            }
            _ = fi6
            _ = v7
        }
    case 0:
        v4 = v3
    case 712:
        var v45 uint32 = 905
        _ = v45
    }
    for fi46 := 0; fi46 < 1; fi46++ {
        if (0 < 62) {
            for fi47 := 0; fi47 < 1; fi47++ {
                switch (fi47 >> ((255) & 63)) {
                case 9223372036854775807:
                    v48 := GenS6{F0: 245}
                    _ = v48
                case 32767:
                    v5 -= v5
                default:
                    v3 = v3
                    switch ("a" + (((((("" + "\n\t") + (("a" + ("hello" + "a")) + "foo bar")) + ("\n\t" + "hello")) + ((("a" + "") + "hello") + "hello")) + "hello") + ((("foo bar" + (("" + "foo bar") + ("" + "a"))) + "\n\t") + "hello"))) {
                    case "\n\t":
                        if (v4 && ("\n\t" == (("\n\t" + "a") + ("foo bar" + "foo bar")))) {
                            var v49 int16 = 42
                            v50 := GenS7{F0: v49, F1: (("hello" + "\n\t") + "")}
                            _ = v49
                            _ = v50
                        } else {
                            v51 := GenS1{F0: v4, F1: "foo bar"}
                            _ = v51
                        }
                    default:
                        v3 = (128 > 228)
                    }
                }
                v52 := GenS3{F0: 0, F1: 42}
                _ = fi47
                _ = v52
            }
            var v53 uint8 = 127
            fmt.Println(v53, v5, v5)
            _ = v53
        } else {
            var v54 bool = (v3 && ((v4 || false) && (!false)))
            if true {
                var v55 uint64 = 2
                for fi56 := 0; fi56 < 11; fi56++ {
                    switch fi46 {
                    case -244:
                        switch fi46 {
                        case 456:
                            if v3 {
                                v54 = (((-(v5)) >= (v5 & v5)) && ((float32((float64(3.14159) + 3.14159)) - 1.5) < (((0.0 * -2.25) / float32((-2.25 + float64(1e10)))) - (float32(((0.0 / (0.1 - 1e10)) / (1e10 / 1.0))) - float32(1e10)))))
                                delete(v1, 2147483647)
                                if (875 >= 0) {
                                    switch "hello" {
                                    case "\n\t":
                                        fmt.Println(v55, v4)
                                    default:
                                        v54 = v54
                                    }
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                } else {
                                    if (42 == 42) {
                                        switch fi46 {
                                        case -197:
                                            switch ("hello" + "a") {
                                            case "foo bar":
                                                switch "hello" {
                                                case "\n\t":
                                                    if ((1.5 - 0.0) == 1.5) {
                                                        var v57 float64 = 1.5
                                                        _ = v57
                                                    }
                                                case "":
                                                    var v58 uint = 0
                                                    _ = v58
                                                }
                                            default:
                                                v59 := v1[2147483648]
                                                _ = v59
                                            }
                                        case 128:
                                            if (!(false && v3)) {
                                                var v60 int16 = -1
                                                var v61 uint64 = (v55 * 4294967295)
                                                v60 -= v60
                                                _ = v60
                                                _ = v61
                                            }
                                        case -1:
                                            var v62 uint8 = 1
                                            _ = v62
                                        }
                                        var v63 bool = v4
                                        _ = v63
                                    }
                                    if v3 {
                                        switch "\n\t" {
                                        case "a":
                                            var v64 string = "foo bar"
                                            _ = v64
                                        case "":
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        case "foo bar":
                                            var v65 string = ("\n\t" + "\n\t")
                                            _ = v65
                                        }
                                        v66 := GenS7{F0: 771, F1: "hello"}
                                        if (1 != 255) {
                                            v5 -= (-(v5))
                                            var v67 float32 = 0.0
                                            v55 -= v55
                                            _ = v67
                                        }
                                        v3 = (!((^(v5)) == (v5 - (v5 * (v5 % (((^(v5))) | 1))))))
                                        _ = v66
                                    }
                                }
                            } else {
                                v4 = v54
                            }
                        default:
                            v5 = (v5 % ((v5) | 1))
                        }
                    }
                    _ = fi56
                }
                var v68 int8 = 0
                _ = v55
                _ = v68
            } else {
                for fi69 := 0; fi69 < 7; fi69++ {
                    switch "foo bar" {
                    case "\n\t":
                        var v70 bool = (1 >= 4294967295)
                        _ = v70
                    case "a":
                        fmt.Println(fi46, v54, v4)
                    case "foo bar":
                        switch fi69 {
                        case -15:
                            switch (-(fi69)) {
                            case 779:
                                switch "a" {
                                case "foo bar":
                                    v71 := GenS8{F0: int8(v5)}
                                    _ = v71
                                case "":
                                    v72 := GenS8{F0: -1}
                                    _ = v72
                                }
                            default:
                                v73 := GenS9{F0: (0 == 65535)}
                                _ = v73
                            }
                        case 874:
                            if len(v2) > 0 {
                                v2[0] = "hello"
                            }
                        case 127:
                            v74 := GenS2{F0: 686}
                            _ = v74
                        }
                        v5 += (v5 * (v5 ^ v5))
                    default:
                        v3 = true
                    }
                    if true {
                        switch 127 {
                        case -1:
                            _ = len(v1)
                        case -9223372036854775808:
                            if (!v3) {
                                v4 = (!(198 <= 1))
                                var v75 int64 = -9223372036854775808
                                _ = len(v1)
                                switch ("foo bar" + "a") {
                                case "a":
                                    delete(v1, 1)
                                case "hello":
                                    if v3 {
                                        v5 = (v5 - (v5 - (v5 ^ (v5 >> ((4294967295) & 63)))))
                                    }
                                case "\n\t":
                                    if v4 {
                                        if (32767 < -14) {
                                            if (v54 || true) {
                                                _ = len(v1)
                                                v75 = (^(v75))
                                                fmt.Println(fi46, v4)
                                                v76 := GenS4{F0: (v75 + v75), F1: 42, F2: 864}
                                                _ = v76
                                            }
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            switch "" {
                                            case "a":
                                                var v77 int8 = 100
                                                _ = v77
                                            }
                                        } else {
                                            v75 = v75
                                            if v4 {
                                                var v78 bool = v3
                                                _ = v78
                                            } else {
                                                if v54 {
                                                    _ = len(v2)
                                                } else {
                                                    var v79 float64 = (0.0 / 1.0)
                                                    _ = v79
                                                }
                                            }
                                        }
                                        v80 := GenS10{F0: 0, F1: 42, F2: 42}
                                        v75 += (v75 & 951)
                                        _ = v80
                                    }
                                }
                                v3 = true
                                _ = v75
                            }
                        case -563:
                            v5 += v5
                        default:
                            v1[674] = fi46
                        }
                        v81 := GenS11{F0: -527, F1: 0, F2: -2147483648}
                        var v82 bool = v4
                        _ = v81
                        _ = v82
                    } else {
                        fmt.Println(fi69, v54)
                        v3 = v54
                        v83 := GenS11{F0: fi69, F1: -1, F2: (v5 - 2147483647)}
                        if (v4 || v4) {
                            v3 = (v4 || ((0.1 >= 0.0) || v54))
                            fmt.Println(fi46)
                            if true {
                                fmt.Println(v3, v3)
                                var v84 bool = v54
                                v4 = true
                                v54 = (42 <= 42)
                                _ = v84
                            }
                        }
                        var v85 int32 = v5
                        _ = v83
                        _ = v85
                    }
                    var v86 uint8 = 128
                    fmt.Println(v4, fi69)
                    _ = fi69
                    _ = v86
                }
            }
            if (!v54) {
                switch (fi46 | (fi46 & fi46)) {
                case -1:
                    var v87 uint16 = 65535
                    _ = v87
                case 4294967295:
                    var v88 bool = (fi46 >= fi46)
                    _ = v88
                case 42:
                    v89 := GenS12{F0: 0}
                    _ = v89
                default:
                    v90 := GenS7{F0: -698, F1: ("foo bar" + ((((("\n\t" + ((("hello" + ("" + (("foo bar" + "") + ("hello" + "")))) + (("\n\t" + "") + "\n\t")) + "a")) + "a") + "foo bar") + ("hello" + "foo bar")) + "hello"))}
                    _ = v90
                }
            }
            _ = v54
        }
        _ = fi46
    }
    var v91 string = "foo bar"
    var v92 int = 556
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v91:", v91)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v91
    _ = v92
}
