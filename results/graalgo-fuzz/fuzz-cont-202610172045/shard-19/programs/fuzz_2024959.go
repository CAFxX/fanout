package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint8; F1 uint8; F2 uint8 }
type GenS6 struct { F0 int }

func main() {
    var v1 bool = ((!((1 > 42) && ((2 != 42) || (-2147483648 != -180)))) && (-32768 <= 118))
    v2 := []string{"", ("hello" + "a")}
    var v3 float32 = (1e10 + 0.0)
    var v4 int8 = 1
    var v5 bool = true
    v6 := GenS1{F0: 1, F1: -68}
    switch 42 {
    case 2147483647:
        if (65535 >= 42) {
            var v7 float32 = (((((v3 / ((float32(-2.25) / (1.5 + -2.25)) / v3)) / v3) - (v3 + v3)) * float32(1e10)) - ((1.5 * 0.0) * ((-2.25 + float32(((0.0 + (((float64(1e10) * float64(v3)) + 0.1) + float64(3.14159))) / float64(v3)))) + 1.5)))
            v8 := GenS1{F0: 2, F1: 0}
            switch 329 {
            case -580:
                v3 -= v7
            }
            _ = v7
            _ = v8
        }
    case -729:
        v3 -= (float32(3.14159) * (1.5 / v3))
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v9 int = 0
    switch v9 {
    case 9223372036854775807:
        var v10 string = ("" + "")
        _ = v10
    case 2147483647:
        v11 := GenS1{F0: 67, F1: 0}
        _ = v11
    case 725:
        v12 := GenS1{F0: 65535, F1: 1}
        _ = v12
    }
    if ((2 == 0) && (false || v5)) {
        switch "foo bar" {
        case "hello":
            v1 = v5
        case "foo bar":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case "\n\t":
            if v5 {
                if v1 {
                    if ((2147483648 > 65535) || v1) {
                        var v13 int64 = 1
                        _ = v13
                    }
                }
                v4 = v4
                var v14 bool = (0.1 <= 3.14159)
                _ = v14
            } else {
                var v15 uint = 18446744073709551615
                fmt.Println(v1, v15)
                v16 := GenS2{F0: -491}
                _ = v15
                _ = v16
            }
        }
        switch v9 {
        case -2147483648:
            for fi17 := 0; fi17 < 6; fi17++ {
                var v18 float32 = v3
                v19 := GenS1{F0: 518, F1: -1}
                _ = fi17
                _ = v18
                _ = v19
            }
        case 0:
            if v1 {
                for fi20 := 0; fi20 < 17; fi20++ {
                    var v21 string = ("a" + "hello")
                    v4 = (v4 - v4)
                    _ = fi20
                    _ = v21
                }
                v22 := GenS2{F0: 1}
                _ = v22
            }
        default:
            if (!(18446744073709551615 > 4294967295)) {
                v5 = true
                v23 := GenS2{F0: 427}
                _ = v23
            } else {
                v1 = true
                _ = len(v2)
                if v5 {
                    v1 = v1
                }
            }
        }
        var v24 int16 = -1
        _ = v24
    } else {
        if (42 >= -834) {
            for fi25 := 0; fi25 < 16; fi25++ {
                if ((float32(float64((v3 / v3))) + float32(3.14159)) > v3) {
                    var v26 int = fi25
                    v5 = (false || ((2 < 255) && (2 >= 67)))
                    _ = v26
                }
                _ = fi25
            }
            v27 := GenS3{F0: 129}
            _ = v27
        }
        var v28 int16 = 32767
        _ = v28
    }
    var v29 int64 = 42
    v29 -= -9223372036854775808
    var v30 int64 = v29
    var v31 int16 = -27
    v4 = v4
    for fi32 := 0; fi32 < 3; fi32++ {
        for fi33 := 0; fi33 < 18; fi33++ {
            if v1 {
                var v34 int32 = 2147483647
                if (!true) {
                    v31 += v31
                    if len(v2) > 0 {
                        v2[0] = "a"
                    }
                    if v5 {
                        v4 -= 98
                        if v1 {
                            var v35 int = (v9 + (v9 % ((fi33) | 1)))
                            _ = v35
                        }
                        var v36 int = (fi33 + (v9 + v9))
                        _ = v36
                    } else {
                        v9 = (fi32 + 32767)
                        v37 := GenS3{F0: 128}
                        _ = v37
                    }
                }
                var v38 uint = 18446744073709551615
                _ = v34
                _ = v38
            }
            if v5 {
                switch "hello" {
                case "":
                    if (v5 || (v3 == v3)) {
                        v39 := GenS3{F0: 0}
                        fmt.Println(fi33, v9, v31)
                        switch (("\n\t" + (((("hello" + "a") + "a") + (("" + ("foo bar" + ("\n\t" + "foo bar"))) + ("foo bar" + "foo bar"))) + "a")) + "foo bar") {
                        case "hello":
                            _ = len(v2)
                        case "foo bar":
                            fmt.Println(fi33, v30, v5)
                        case "":
                            if true {
                                switch "\n\t" {
                                case "a":
                                    v40 := GenS4{F0: v9}
                                    _ = v40
                                }
                                if v1 {
                                    if v5 {
                                        if (!true) {
                                            v3 += ((0.1 / v3) * ((float32(-2.25) / v3) / v3))
                                        }
                                        v41 := GenS5{F0: 0, F1: 146, F2: 0}
                                        var v42 uint64 = 4294967295
                                        v31 = v31
                                        v43 := GenS1{F0: 376, F1: v30}
                                        v30 = (v29 / (((v29 & (v29 / ((622) | 1)))) | 1))
                                        switch "hello" {
                                        case "a":
                                            switch fi33 {
                                            case 42:
                                                switch fi33 {
                                                case 9223372036854775807:
                                                    var v44 int64 = -9223372036854775808
                                                    _ = v44
                                                }
                                            default:
                                                switch fi33 {
                                                case 532:
                                                    v2 = append(v2, "a")
                                                case 32767:
                                                    var v45 float32 = v3
                                                    _ = v45
                                                }
                                                v3 = v3
                                                if len(v2) > 0 {
                                                    v2[0] = ("" + ("foo bar" + ("a" + ("" + (((("hello" + "hello") + "foo bar") + ("foo bar" + (("foo bar" + ("hello" + ("a" + "\n\t"))) + "foo bar"))) + "hello")))))
                                                }
                                                var v46 int = (v9 % ((1) | 1))
                                                _ = v46
                                            }
                                        case "":
                                            v4 -= -3
                                        case "\n\t":
                                            var v47 bool = v1
                                            _ = v47
                                        }
                                        _ = v41
                                        _ = v42
                                        _ = v43
                                    } else {
                                        var v48 uint16 = 891
                                        switch (^(v9)) {
                                        case 713:
                                            v49 := GenS3{F0: 2}
                                            _ = v49
                                        case 4294967295:
                                            v5 = (-2147483648 < 1)
                                        case 32768:
                                            v2 = append(v2, "a")
                                        }
                                        _ = v48
                                    }
                                    v9 = int(v31)
                                    v30 -= v29
                                } else {
                                    var v50 uint = 2
                                    _ = v50
                                }
                            }
                        }
                        switch "foo bar" {
                        case "\n\t":
                            switch fi32 {
                            case 9223372036854775807:
                                if len(v2) > 0 {
                                    v2[0] = "a"
                                }
                            case 4294967295:
                                if false {
                                    v3 -= (0.0 * (-2.25 / 3.14159))
                                    switch "foo bar" {
                                    case "a":
                                        var v51 int64 = (v29 * v30)
                                        _ = v51
                                    case "\n\t":
                                        switch ("\n\t" + (("foo bar" + "") + "hello")) {
                                        case "hello":
                                            var v52 int64 = 0
                                            _ = v52
                                        }
                                    case "hello":
                                        fmt.Println(v5, fi33)
                                    }
                                }
                            case 127:
                                v53 := GenS2{F0: (v30 % ((-9223372036854775808) | 1))}
                                _ = v53
                            }
                        default:
                            var v54 uint16 = 122
                            _ = v54
                        }
                        _ = len(v2)
                        _ = v39
                    } else {
                        switch 376 {
                        case 1:
                            v55 := GenS4{F0: (fi32 - fi32)}
                            _ = v55
                        }
                    }
                }
            }
            v56 := GenS3{F0: 42}
            if ((v30 >> ((0) & 63)) == -183) {
                v30 = (v30 << ((4294967295) & 63))
                if ((v1 || ((float64(float32((0.1 - (((-2.25 * 3.14159) / float64(v3)) + (3.14159 - 1e10))))) * (-2.25 + 0.1)) == 0.1)) && true) {
                    v4 = -128
                    if v1 {
                        v57 := GenS2{F0: (v30 & -9223372036854775808)}
                        if false {
                            v58 := GenS2{F0: 42}
                            _ = v58
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            fmt.Println(fi32)
                            if (127 != 146) {
                                if v5 {
                                    if len(v2) > 0 {
                                        v2[0] = (("a" + "foo bar") + ("" + "a"))
                                    }
                                    var v59 string = "\n\t"
                                    _ = v59
                                }
                                v2 = append(v2, "")
                                v60 := GenS4{F0: -9223372036854775808}
                                _ = v60
                            }
                        }
                        _ = v57
                    } else {
                        if (!v5) {
                            v1 = v5
                            var v61 int64 = (v29 + v29)
                            _ = v61
                        }
                        switch (fi32 + fi32) {
                        case -1:
                            v62 := GenS6{F0: fi33}
                            _ = v62
                        }
                    }
                    fmt.Println(v4, v9, fi33)
                }
                fmt.Println(v1, v5)
            }
            _ = fi33
            _ = v56
        }
        for fi63 := 0; fi63 < 17; fi63++ {
            if len(v2) > 0 {
                v2[0] = "\n\t"
            }
            _ = fi63
        }
        _ = fi32
    }
    for fi64 := 0; fi64 < 3; fi64++ {
        v1 = false
        for fi65 := 0; fi65 < 5; fi65++ {
            var v66 uint32 = 2
            _ = fi65
            _ = v66
        }
        v4 -= int8(v29)
        _ = fi64
    }
    fmt.Println(v3)
    _ = len(v2)
    fmt.Println(v9)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
