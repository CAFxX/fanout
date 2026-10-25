package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 uint64; F2 int32 }
type GenS4 struct { F0 uint16; F1 int64 }
type GenS5 struct { F0 int32; F1 bool; F2 uint8 }
type GenS6 struct { F0 uint16; F1 string }

func main() {
    var v1 float64 = (1e10 * 0.1)
    var v2 float64 = 0.0
    var v3 int32 = 2147483647
    if (!true) {
        var v4 uint8 = 124
        _ = v4
    }
    var v5 uint16 = 790
    fmt.Println(v1)
    for fi6 := 0; fi6 < 20; fi6++ {
        for fi7 := 0; fi7 < 3; fi7++ {
            if (false || false) {
                if (9223372036854775807 == fi7) {
                    if (0 != -128) {
                        switch (("hello" + ("a" + "\n\t")) + "a") {
                        case "a":
                            switch "" {
                            case "hello":
                                fmt.Println(v5, v2, fi7)
                            }
                        default:
                            var v8 float32 = (0.1 / 3.14159)
                            _ = v8
                        }
                        v9 := GenS1{F0: (128 <= 1), F1: 255}
                        if (42 < 42) {
                            v5 = v5
                            if false {
                                v10 := GenS2{F0: (v5 | (v5 % (((^(v5))) | 1)))}
                                fmt.Println(v5)
                                _ = v10
                            }
                            v11 := []int{fi7, (fi7 + fi7), fi7, -9223372036854775808}
                            _ = v11
                        }
                        _ = v9
                    }
                    var v12 string = "hello"
                    v13 := []int{(fi7 % ((fi7) | 1))}
                    _ = v12
                    _ = v13
                }
                v14 := GenS1{F0: true, F1: 255}
                v15 := []int{(-(fi7)), int(v5), 9223372036854775807}
                v16 := GenS3{F0: 65535, F1: 2, F2: v3}
                if len(v15) > 0 {
                    v15[0] = fi7
                }
                _ = v14
                _ = v15
                _ = v16
            } else {
                var v17 float32 = (-2.25 * 1.5)
                switch "a" {
                case "a":
                    if (3.14159 < float32(0.1)) {
                        fmt.Println(v17)
                        v18 := []int{}
                        if false {
                            if true {
                                v5 = (v5 >> ((1) & 63))
                                if len(v18) >= 2 {
                                    v18 = v18[1:len(v18)]
                                }
                            }
                            if len(v18) > 0 {
                                v18[0] = fi6
                            }
                            v1 = (0.1 - ((float64(v17) - (v1 / ((((v1 / v1) - v2) / 1.0) - v2))) * v1))
                            switch 4294967295 {
                            case -956:
                                if (720 == (v3 * v3)) {
                                    var v19 int64 = 0
                                    _ = v19
                                }
                            case 32768:
                                var v20 float64 = v1
                                switch ("" + ("hello" + "hello")) {
                                case "":
                                    v1 += v1
                                default:
                                    if (1 == -1) {
                                        v3 -= v3
                                    } else {
                                        var v21 string = "a"
                                        _ = v21
                                    }
                                }
                                _ = v20
                            }
                        } else {
                            var v22 bool = (v5 < (v5 | (^(v5))))
                            var v23 int64 = 42
                            _ = v22
                            _ = v23
                        }
                        _ = v18
                    } else {
                        v17 = -2.25
                        var v24 float64 = v2
                        v25 := []int{(fi7 & 42), (fi7 / (((fi6 >> ((1) & 63))) | 1))}
                        switch ("hello" + ("a" + "")) {
                        case "hello":
                            var v26 uint8 = 26
                            _ = v26
                        case "\n\t":
                            if (981 != 1) {
                                switch ((("\n\t" + "\n\t") + "\n\t") + "a") {
                                case "a":
                                    switch "\n\t" {
                                    case "foo bar":
                                        if len(v25) > 0 {
                                            v25[0] = -1
                                        }
                                    case "":
                                        if false {
                                            switch 2147483647 {
                                            case 127:
                                                if ((1e10 >= (-2.25 * float32((v24 / (v2 / 1.0))))) || true) {
                                                    if len(v25) > 0 {
                                                        v25[0] = (-(fi6))
                                                    }
                                                }
                                            case 2147483647:
                                                var v27 uint = 1
                                                _ = v27
                                            case -857:
                                                v1 = v2
                                            default:
                                                var v28 float64 = v2
                                                _ = v28
                                            }
                                        }
                                        v2 = v24
                                    default:
                                        v24 -= v1
                                    }
                                default:
                                    v29 := GenS4{F0: (v5 & 642), F1: 9223372036854775807}
                                    _ = v29
                                }
                                if (false && (-567 <= 0)) {
                                    v3 += (v3 << ((151) & 63))
                                    v5 = v5
                                }
                            }
                        case "":
                            v25 = append(v25, fi6)
                        default:
                            v3 = -1
                            switch 128 {
                            case 32767:
                                var v30 uint = 2
                                _ = v30
                            case 127:
                                if ((v5 >> ((673) & 63)) == 65535) {
                                    v31 := GenS4{F0: 1, F1: 1}
                                    var v32 int64 = 1
                                    if len(v25) >= 2 {
                                        v25 = v25[1:len(v25)]
                                    }
                                    _ = v31
                                    _ = v32
                                } else {
                                    if true {
                                        v33 := GenS4{F0: (v5 >> ((0) & 63)), F1: 9223372036854775807}
                                        v24 += 1.5
                                        v34 := GenS2{F0: v5}
                                        _ = v33
                                        _ = v34
                                    } else {
                                        v3 += (^(v3))
                                    }
                                    if len(v25) >= 2 {
                                        v25 = v25[1:len(v25)]
                                    }
                                    var v35 uint = 885
                                    _ = v35
                                }
                            case -624:
                                var v36 uint16 = 2
                                _ = v36
                            default:
                                v5 = v5
                            }
                        }
                        _ = v24
                        _ = v25
                    }
                }
                _ = v17
            }
            _ = fi7
        }
        v37 := []int{}
        var v38 int32 = 42
        _ = fi6
        _ = v37
        _ = v38
    }
    v1 -= v2
    if ((0.0 / 3.14159) < (3.14159 / 0.1)) {
        v39 := []int{32768}
        _ = v39
    } else {
        if (18446744073709551615 != 255) {
            var v40 int64 = 42
            v40 = (v40 & -1)
            fmt.Println(v5)
            _ = v40
        } else {
            switch "" {
            case "":
                if true {
                    switch 4294967295 {
                    case 133:
                        fmt.Println(v1, v5)
                    case 50:
                        v41 := GenS3{F0: 2, F1: 4294967295, F2: 2147483647}
                        _ = v41
                    case -2147483648:
                        switch 127 {
                        case 128:
                            for fi42 := 0; fi42 < 15; fi42++ {
                                v1 = -2.25
                                _ = fi42
                            }
                        default:
                            v5 = v5
                        }
                    default:
                        for fi43 := 0; fi43 < 1; fi43++ {
                            for fi44 := 0; fi44 < 20; fi44++ {
                                if ((((1 < 2147483647) && (("" + "\n\t") > "hello")) && (fi43 < (fi43 | 9223372036854775807))) && (("\n\t" + (("" + "foo bar") + ("hello" + "a"))) >= ("a" + "hello"))) {
                                    v1 = (float64(1.5) + (3.14159 / v2))
                                    v5 = v5
                                } else {
                                    var v45 int16 = 335
                                    _ = v45
                                }
                                if (!true) {
                                    var v46 int64 = 0
                                    var v47 float64 = (v1 * float64(0.1))
                                    _ = v46
                                    _ = v47
                                } else {
                                    v5 += (v5 >> ((72) & 63))
                                    if false {
                                        var v48 uint16 = (v5 >> ((65535) & 63))
                                        _ = v48
                                    }
                                    if (9223372036854775807 > -837) {
                                        var v49 int = 967
                                        _ = v49
                                    }
                                }
                                if (2 == 910) {
                                    v3 -= (v3 | 287)
                                    var v50 int32 = v3
                                    v3 = v50
                                    _ = v50
                                }
                                switch "\n\t" {
                                case "":
                                    switch -134 {
                                    case 4294967295:
                                        if true {
                                            if (v5 >= v5) {
                                                v3 = v3
                                            } else {
                                                v51 := []int{fi43, 2147483647, fi43, 593}
                                                v3 += (v3 % ((891) | 1))
                                                _ = v51
                                            }
                                        }
                                    case -2147483648:
                                        fmt.Println(v1)
                                    default:
                                        fmt.Println(fi43)
                                        v52 := GenS3{F0: 2, F1: 4294967295, F2: -1}
                                        _ = v52
                                    }
                                case "\n\t":
                                    v53 := GenS4{F0: v5, F1: 470}
                                    _ = v53
                                }
                                _ = fi44
                            }
                            for fi54 := 0; fi54 < 11; fi54++ {
                                fmt.Println(v1, v2)
                                v55 := GenS1{F0: (0 == 18446744073709551615), F1: 0}
                                fmt.Println(v5, v3)
                                _ = fi54
                                _ = v55
                            }
                            _ = fi43
                        }
                    }
                }
            default:
                v56 := GenS4{F0: 42, F1: 653}
                _ = v56
            }
        }
        switch "" {
        case "hello":
            for fi57 := 0; fi57 < 5; fi57++ {
                v58 := GenS5{F0: v3, F1: true, F2: 1}
                if true {
                    v5 = v5
                    fmt.Println(v1, fi57)
                } else {
                    v2 = 0.1
                }
                if false {
                    v59 := []int{(fi57 - (fi57 / ((2147483647) | 1))), (fi57 % ((-201) | 1))}
                    if true {
                        for fi60 := 0; fi60 < 6; fi60++ {
                            if len(v59) >= 2 {
                                v59 = v59[1:len(v59)]
                            }
                            _ = fi60
                        }
                        v1 += (v2 / (((v1 * (v2 - ((0.1 / v2) * v2))) * v2) / v1))
                    } else {
                        if true {
                            v61 := GenS2{F0: (^(v5))}
                            switch fi57 {
                            case 127:
                                v3 += (v3 + 0)
                            case 4294967295:
                                v62 := GenS6{F0: (v5 >> ((2) & 63)), F1: "a"}
                                _ = v62
                            }
                            v2 -= float64(1e10)
                            _ = v61
                        }
                        v59 = append(v59, fi57)
                        var v63 uint32 = 0
                        _ = v63
                    }
                    _ = v59
                } else {
                    for fi64 := 0; fi64 < 18; fi64++ {
                        if (!(437 != -570)) {
                            v3 = (v3 + v3)
                            fmt.Println(fi57)
                            v65 := []int{(fi57 * (fi57 + fi57)), (fi64 << ((650) & 63)), -623, fi57}
                            _ = v65
                        }
                        switch fi57 {
                        case -2147483648:
                            v66 := []int{}
                            _ = v66
                        case -182:
                            v67 := GenS4{F0: 65535, F1: -190}
                            _ = v67
                        }
                        v68 := []int{(-(fi64)), (fi57 * fi64), (fi64 >> ((255) & 63))}
                        _ = fi64
                        _ = v68
                    }
                }
                _ = fi57
                _ = v58
            }
        }
        var v69 int8 = -91
        _ = v69
    }
    fmt.Println(v5, v5, v1)
    v70 := []int{549, -960, 888}
    var v71 bool = (int16(v3) > -32768)
    if len(v70) > 0 {
        v70[0] = 32768
    }
    var v72 string = "\n\t"
    v71 = true
    if len(v70) > 0 {
        v70[0] = -151
    }
    v71 = v71
    v70 = append(v70, 128)
    if len(v70) > 0 {
        v70[0] = 128
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v70:", len(v70), v70)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v70
    _ = v71
    _ = v72
}
