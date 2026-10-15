package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 int32; F2 uint }
type GenS4 struct { F0 int }
type GenS5 struct { F0 string; F1 uint }
func genFn0(p0 int32, p1 string, p2 int16) uint8 {
    v4 := []int{711, 32768}
    v5 := GenS1{F0: false, F1: 18446744073709551615, F2: (p2 % ((0) | 1))}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 127
}
func genFn1() uint16 {
    var v6 string = "foo bar"
    if (0 <= 1) {
        var v7 uint16 = 42
        for fi8 := 0; fi8 < 6; fi8++ {
            for fi9 := 0; fi9 < 1; fi9++ {
                v6 = "foo bar"
                v10 := []int{fi8}
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
        _ = v7
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        var v12 uint64 = 42
        v13 := []int{(fi11 - 127), (fi11 / (((-(fi11))) | 1)), fi11, 485}
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = v6
    return 65535
}

func main() {
    v1 := []string{"", "hello", ("a" + "hello")}
    var v2 bool = (!false)
    var v3 int = 42
    for fi14 := 0; fi14 < 10; fi14++ {
        switch 32767 {
        case 2147483647:
            for fi15 := 0; fi15 < 20; fi15++ {
                var v16 int = (fi14 - (v3 + (v3 & (fi14 ^ 9223372036854775807))))
                _ = fi15
                _ = v16
            }
        case -2147483648:
            fmt.Println(v3)
        case 1:
            if (((((float32(0.0) / (0.1 - -2.25)) - 0.0) + -2.25) * -2.25) >= ((float32(3.14159) + (((1e10 + 0.0) - ((1e10 * 0.1) + 1e10)) + (0.1 + -2.25))) - 3.14159)) {
                v2 = false
                if v2 {
                    v3 += fi14
                    switch (("" + "hello") + "\n\t") {
                    case "foo bar":
                        var v17 string = "hello"
                        _ = v17
                    case "\n\t":
                        fmt.Println(fi14, fi14, v3)
                    case "hello":
                        v2 = v2
                    default:
                        fmt.Println(v3, fi14, v2)
                    }
                }
                for fi18 := 0; fi18 < 4; fi18++ {
                    if (v2 && ((v3 % (((v3 | fi14)) | 1)) == 42)) {
                        v2 = v2
                        fmt.Println(fi14, v2, fi14)
                        switch 1 {
                        case 0:
                            if len(v1) > 0 {
                                v1[0] = ("\n\t" + "a")
                            }
                        case 425:
                            var v19 int = v3
                            _ = v19
                        default:
                            var v20 int = 411
                            _ = v20
                        }
                    }
                    fmt.Println(fi18, v3, fi18)
                    var v21 int = fi14
                    _ = fi18
                    _ = v21
                }
            } else {
                if (0.1 > float64(1.5)) {
                    for fi22 := 0; fi22 < 14; fi22++ {
                        var v23 bool = false
                        switch fi22 {
                        case 127:
                            var v24 uint32 = 2147483647
                            _ = v24
                        case 1:
                            v3 = 127
                        default:
                            switch "a" {
                            case "\n\t":
                                v23 = v23
                            }
                        }
                        if (-2147483648 <= -51) {
                            v25 := GenS1{F0: true, F1: 42, F2: -32768}
                            if (!true) {
                                var v26 uint32 = 2147483647
                                var v27 int = fi22
                                _ = v26
                                _ = v27
                            }
                            var v28 uint16 = 355
                            _ = v25
                            _ = v28
                        }
                        var v29 uint32 = 42
                        _ = fi22
                        _ = v23
                        _ = v29
                    }
                    if v2 {
                        switch (v3 / ((fi14) | 1)) {
                        case 42:
                            if (42 == 1) {
                                var v30 uint = 42
                                v1 = append(v1, "")
                                v31 := GenS2{F0: genFn0(2147483647, (((("a" + ("\n\t" + (("\n\t" + ("foo bar" + "\n\t")) + "foo bar"))) + ("hello" + "")) + ("\n\t" + "hello")) + "\n\t"), 42)}
                                _ = v30
                                _ = v31
                            }
                        case 127:
                            v32 := GenS2{F0: 42}
                            _ = v32
                        default:
                            var v33 int16 = 381
                            _ = v33
                        }
                        fmt.Println(fi14)
                        v2 = true
                    }
                    v2 = (v2 && v2)
                }
            }
        }
        v3 -= 1
        _ = fi14
    }
    switch v3 {
    case -290:
        for fi34 := 0; fi34 < 2; fi34++ {
            if (-1 >= 1) {
                if v2 {
                    v3 = v3
                    v2 = ((fi34 | (v3 & (fi34 / ((127) | 1)))) > (v3 * (v3 << ((190) & 63))))
                } else {
                    var v35 uint8 = genFn0(63, (("hello" + "") + ("\n\t" + "\n\t")), -1)
                    v1 = append(v1, (("foo bar" + ("hello" + "foo bar")) + "\n\t"))
                    switch v3 {
                    case 2147483647:
                        var v36 uint16 = 255
                        _ = v36
                    case 645:
                        v37 := GenS2{F0: (v35 + (v35 * v35))}
                        _ = v37
                    default:
                        var v38 uint8 = 128
                        _ = v38
                    }
                    _ = v35
                }
                var v39 int64 = 42
                v39 += (v39 % (((v39 | 9223372036854775807)) | 1))
                _ = v39
            } else {
                var v40 uint32 = 641
                var v41 int16 = 144
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v40
                _ = v41
            }
            var v42 int8 = 0
            if len(v1) > 0 {
                v1[0] = ("a" + "a")
            }
            _ = fi34
            _ = v42
        }
    }
    v43 := GenS3{F0: -9223372036854775808, F1: -2147483648, F2: 255}
    for fi44 := 0; fi44 < 16; fi44++ {
        fmt.Println(fi44, fi44, v2)
        switch "\n\t" {
        case "a":
            for fi45 := 0; fi45 < 14; fi45++ {
                if (!(((0.1 + 0.0) / 1e10) != float32((float64(3.14159) - float64((0.0 - float32(-2.25))))))) {
                    var v46 uint8 = 31
                    fmt.Println(fi44, fi45, fi45)
                    v2 = true
                    _ = v46
                } else {
                    if v2 {
                        v47 := GenS2{F0: 255}
                        _ = v47
                    }
                    if len(v1) > 0 {
                        v1[0] = "hello"
                    }
                    switch fi45 {
                    case 9223372036854775807:
                        v48 := GenS4{F0: (fi45 - v3)}
                        _ = v48
                    case 127:
                        var v49 int = 127
                        _ = v49
                    case 32767:
                        v50 := GenS3{F0: (fi45 + (fi44 - fi44)), F1: 261, F2: 127}
                        _ = v50
                    default:
                        if false {
                            var v51 int64 = 421
                            var v52 int64 = v51
                            _ = v51
                            _ = v52
                        } else {
                            v3 -= fi44
                            if (!(2 < 2147483648)) {
                                var v53 float32 = (1.5 / (-2.25 * (float32(1.5) + ((float32(0.0) + 1e10) + 0.1))))
                                _ = v53
                            } else {
                                v54 := GenS2{F0: genFn0(-265, "a", 42)}
                                if (genFn1() == 2) {
                                    switch "foo bar" {
                                    case "foo bar":
                                        v2 = (2147483647 >= 4294967295)
                                    case "a":
                                        if len(v1) > 0 {
                                            v1[0] = "foo bar"
                                        }
                                    default:
                                        v55 := GenS5{F0: ("a" + "\n\t"), F1: 42}
                                        _ = v55
                                    }
                                    var v56 uint = 4294967295
                                    var v57 int64 = 42
                                    _ = len(v1)
                                    var v58 int16 = 32767
                                    _ = len(v1)
                                    v2 = (1e10 == -2.25)
                                    _ = v56
                                    _ = v57
                                    _ = v58
                                }
                                _ = v54
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                    }
                    switch 42 {
                    case 2147483647:
                        var v59 uint8 = 12
                        _ = v59
                    case 42:
                        if v2 {
                            var v60 int = 127
                            _ = v60
                        }
                    case 1:
                        var v61 float64 = 1.5
                        v2 = v2
                        v3 = fi44
                        if (!false) {
                            switch (fi44 + (fi45 - (fi44 | fi44))) {
                            case 127:
                                v2 = (v2 && (42 >= 1))
                            case 128:
                                v3 = -485
                            default:
                                if len(v1) > 0 {
                                    v1[0] = (("hello" + "a") + "a")
                                }
                            }
                        }
                        _ = v61
                    }
                }
                _ = fi45
            }
        case "":
            if (-486 > 32767) {
                for fi62 := 0; fi62 < 10; fi62++ {
                    v2 = v2
                    v2 = true
                    var v63 int16 = 879
                    v63 = v63
                    _ = fi62
                    _ = v63
                }
                v2 = (0.0 >= ((1.5 * -2.25) / (1.5 / 1e10)))
                for fi64 := 0; fi64 < 18; fi64++ {
                    var v65 int = v3
                    fmt.Println(v3)
                    _ = fi64
                    _ = v65
                }
            }
        default:
            var v66 int = fi44
            _ = v66
        }
        _ = fi44
    }
    var v67 int8 = -46
    var v68 uint16 = 374
    for fi69 := 0; fi69 < 3; fi69++ {
        v3 -= 9223372036854775807
        var v70 uint32 = 4294967295
        v68 += (v68 + (v68 >> ((326) & 63)))
        _ = fi69
        _ = v70
    }
    if v2 {
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
    } else {
        if len(v1) > 0 {
            v1[0] = "hello"
        }
    }
    for fi71 := 0; fi71 < 2; fi71++ {
        v72 := GenS2{F0: 139}
        _ = fi71
        _ = v72
    }
    _ = len(v1)
    var v73 int = -2147483648
    v73 = v3
    v68 = (v68 % (((^(v68))) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v43
    _ = v67
    _ = v68
    _ = v73
}
