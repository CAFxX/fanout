package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 uint16; F1 uint16; F2 bool }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string; F1 string }

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = -2.25
    var v3 string = "foo bar"
    v4 := make(map[string]int)
    v5 := []uint32{2147483647, 0, 986, 4294967295}
    var v6 uint16 = 42
    switch 9223372036854775807 {
    case 4294967295:
        _ = len(v1)
    default:
        for fi7 := 0; fi7 < 13; fi7++ {
            v3 = "\n\t"
            _ = fi7
        }
    }
    switch ("hello" + v3) {
    case "":
        switch 1 {
        case -804:
            fmt.Println(v3, v6, v6)
        case 0:
            v2 = float32(1e10)
        case 2147483647:
            if (((!false) || false) || (-1 > 1)) {
                for fi8 := 0; fi8 < 14; fi8++ {
                    v4[v3] = fi8
                    fmt.Println(fi8)
                    _ = fi8
                }
                _ = len(v1)
            }
        default:
            for fi9 := 0; fi9 < 4; fi9++ {
                var v10 int16 = 0
                if false {
                    var v11 string = v3
                    for fi12 := 0; fi12 < 17; fi12++ {
                        if ((v10 ^ 32767) != 1000) {
                            var v13 int32 = -2147483648
                            delete(v4, (("\n\t" + ((v3 + v3) + v3)) + (("foo bar" + v11) + v3)))
                            _ = v13
                        }
                        _ = len(v5)
                        if len(v5) > 0 {
                            v5[0] = 365
                        }
                        _ = fi12
                    }
                    _ = v11
                }
                v6 += v6
                _ = fi9
                _ = v10
            }
        }
    case "a":
        switch v3 {
        case "\n\t":
            v6 += v6
        case "foo bar":
            for fi14 := 0; fi14 < 10; fi14++ {
                var v15 int8 = 1
                v3 = v3
                _ = fi14
                _ = v15
            }
        }
    }
    v6 = v6
    for fi16 := 0; fi16 < 17; fi16++ {
        for fi17 := 0; fi17 < 19; fi17++ {
            if (false && true) {
                if (269 >= 209) {
                    if ((0.1 / 0.1) >= 3.14159) {
                        if len(v5) > 0 {
                            v5[0] = 1
                        }
                        if (0 == 42) {
                            if ((937 == 269) || false) {
                                if true {
                                    v3 = ("" + ("\n\t" + v3))
                                    v18 := GenS1{F0: false, F1: 924}
                                    v19 := GenS1{F0: false, F1: 4294967295}
                                    _ = v18
                                    _ = v19
                                } else {
                                    v3 = v3
                                    v20 := GenS1{F0: ((v3 + "a") >= "\n\t"), F1: 42}
                                    var v21 int32 = 1
                                    var v22 int = 32767
                                    if ((!true) || true) {
                                        fmt.Println(fi16)
                                        switch 32768 {
                                        case 128:
                                            if (!(!false)) {
                                                v21 = v21
                                                if false {
                                                    if len(v5) > 0 {
                                                        v5[0] = 909
                                                    }
                                                    v21 += v21
                                                }
                                                fmt.Println(v21, fi17)
                                            }
                                        case 336:
                                            var v23 int8 = -128
                                            _ = v23
                                        }
                                    }
                                    _ = v20
                                    _ = v21
                                    _ = v22
                                }
                                switch (v3 + "") {
                                case "a":
                                    var v24 int8 = 1
                                    _ = v24
                                }
                            } else {
                                var v25 uint = 42
                                if false {
                                    if len(v5) > 0 {
                                        v5[0] = 278
                                    }
                                    if ((560 >= 42) || false) {
                                        if true {
                                            switch v3 {
                                            case "foo bar":
                                                var v26 float64 = 1.5
                                                _ = v26
                                            }
                                            var v27 string = v3
                                            _ = v27
                                        } else {
                                            var v28 float64 = (0.0 / -2.25)
                                            _ = v28
                                        }
                                        if (!false) {
                                            if (true && (false && (!((v6 & v6) >= (v6 | (v6 & 911)))))) {
                                                v3 = "hello"
                                                if (!true) {
                                                    v25 -= v25
                                                    switch ("foo bar" + ("\n\t" + v3)) {
                                                    case "foo bar":
                                                        var v29 float64 = -2.25
                                                        _ = v29
                                                    default:
                                                        var v30 float64 = 3.14159
                                                        _ = v30
                                                    }
                                                    v31 := GenS1{F0: true, F1: 4294967295}
                                                    _ = v31
                                                } else {
                                                    fmt.Println(v25)
                                                    if true {
                                                        v25 = (v25 << (((^(v25))) & 63))
                                                    } else {
                                                        var v32 bool = ((false || ("foo bar" >= (("hello" + (v3 + "foo bar")) + "\n\t"))) && (1e10 > 1.5))
                                                        _ = v32
                                                    }
                                                    var v33 int32 = -2147483648
                                                    if (886 >= 818) {
                                                        if true {
                                                            v34 := GenS1{F0: (-2.25 == (v2 + 1.5)), F1: 2147483648}
                                                            v35 := GenS1{F0: true, F1: 4294967295}
                                                            _ = v34
                                                            _ = v35
                                                        }
                                                        v2 += ((0.1 - v2) + (float32((1e10 / 1e10)) - v2))
                                                        fmt.Println(fi17, v3, v3)
                                                    } else {
                                                        var v36 uint16 = (v6 + 65535)
                                                        _ = v36
                                                    }
                                                    v37 := GenS2{F0: v6, F1: (v6 % ((v6) | 1)), F2: (v2 > float32(((3.14159 * ((-2.25 - float64(float32(1e10))) - -2.25)) - 1.5)))}
                                                    _ = v33
                                                    _ = v37
                                                }
                                            } else {
                                                if false {
                                                    if (fi17 <= fi17) {
                                                        v6 = 65535
                                                    } else {
                                                        fmt.Println(fi16, v25)
                                                    }
                                                    v5 = append(v5, 0)
                                                }
                                            }
                                            if len(v5) > 0 {
                                                v5[0] = 136
                                            }
                                            v38 := GenS1{F0: false, F1: 2147483648}
                                            _ = v38
                                        }
                                        switch (fi16 * (fi16 % ((fi16) | 1))) {
                                        case 897:
                                            fmt.Println(fi17)
                                        }
                                    }
                                    v39 := GenS2{F0: 2, F1: v6, F2: ((true || false) || false)}
                                    _ = v39
                                }
                                v2 = (v2 * (1.5 / v2))
                                _ = len(v5)
                                _ = v25
                            }
                        } else {
                            var v40 uint16 = v6
                            if len(v5) > 0 {
                                v5[0] = 65535
                            }
                            _ = v40
                        }
                    } else {
                        var v41 uint64 = 560
                        _ = v41
                    }
                    fmt.Println(fi16, v2)
                    fmt.Println(v2, fi17, fi17)
                } else {
                    v42 := GenS1{F0: false, F1: 0}
                    _ = v42
                }
                if (-2.25 != (0.1 / ((float64(((float32((3.14159 * 0.1)) / -2.25) + (v2 / 1.5))) * float64(v2)) * (0.1 - 1e10)))) {
                    var v43 int32 = 13
                    _ = v43
                }
            } else {
                v3 = v3
            }
            v44 := GenS2{F0: v6, F1: v6, F2: (2 <= 0)}
            v3 = v3
            switch (fi16 >> ((42) & 63)) {
            case 32767:
                v2 -= 3.14159
            case -2147483648:
                v2 = v2
            case 0:
                _ = len(v5)
            }
            _ = fi17
            _ = v44
        }
        if false {
            for fi45 := 0; fi45 < 6; fi45++ {
                var v46 uint8 = 128
                var v47 uint8 = v46
                v47 = v47
                _ = fi45
                _ = v46
                _ = v47
            }
            v2 -= v2
            fmt.Println(v6)
        }
        _ = fi16
    }
    if true {
        for fi48 := 0; fi48 < 1; fi48++ {
            v5 = append(v5, 2147483647)
            switch v3 {
            case "a":
                v2 = ((v2 - v2) - (float32(0.1) + (v2 * (float32(3.14159) * (1.5 * -2.25)))))
            case "\n\t":
                fmt.Println(fi48, v6)
            case "":
                switch v3 {
                case "foo bar":
                    for fi49 := 0; fi49 < 11; fi49++ {
                        if (((!true) || true) || (((-1 <= -97) || ("hello" < ("hello" + (("" + v3) + "a")))) || false)) {
                            if false {
                                v4[(((v3 + v3) + v3) + ("foo bar" + "foo bar"))] = (fi48 % ((0) | 1))
                                if false {
                                    v6 += (v6 ^ v6)
                                    var v50 int64 = -889
                                    if true {
                                        v2 -= v2
                                        switch (fi48 << ((615) & 63)) {
                                        case 127:
                                            _ = len(v1)
                                        case 516:
                                            fmt.Println(v50, fi48, v6)
                                        case 137:
                                            var v51 string = (("foo bar" + "\n\t") + v3)
                                            _ = v51
                                        }
                                    } else {
                                        if false {
                                            v3 = v3
                                            v2 = (v2 / v2)
                                            _ = len(v4)
                                        }
                                        var v52 uint = 0
                                        _ = v52
                                    }
                                    var v53 float64 = ((1.5 + 0.1) * 1.5)
                                    fmt.Println(v53, v50, v3)
                                    _ = v50
                                    _ = v53
                                }
                                var v54 float64 = 1e10
                                _ = v54
                            } else {
                                v3 = (v3 + (("a" + "foo bar") + "\n\t"))
                            }
                        } else {
                            v3 = (v3 + v3)
                        }
                        if true {
                            var v55 string = "foo bar"
                            var v56 uint = 18446744073709551615
                            var v57 uint64 = 42
                            _ = v55
                            _ = v56
                            _ = v57
                        }
                        v3 = "foo bar"
                        _ = fi49
                    }
                case "\n\t":
                    var v58 int8 = 127
                    _ = v58
                case "hello":
                    v59 := GenS2{F0: (v6 % ((0) | 1)), F1: 42, F2: ("" <= v3)}
                    _ = v59
                }
            default:
                v4[v3] = (fi48 * fi48)
            }
            if len(v5) > 0 {
                v5[0] = 192
            }
            _ = fi48
        }
        if false {
            if true {
                if false {
                    if false {
                        if false {
                            var v60 uint8 = 1
                            _ = v60
                        }
                    } else {
                        v2 = v2
                    }
                    v61 := GenS1{F0: (false && (uint(v6) != 4294967295)), F1: 446}
                    var v62 uint = 18446744073709551615
                    _ = v61
                    _ = v62
                } else {
                    var v63 int8 = 100
                    for fi64 := 0; fi64 < 15; fi64++ {
                        v63 += (v63 % (((v63 + (-(v63)))) | 1))
                        _ = fi64
                    }
                    v3 = "\n\t"
                    _ = v63
                }
                for fi65 := 0; fi65 < 10; fi65++ {
                    if (!false) {
                        var v66 uint32 = 253
                        _ = v66
                    }
                    for fi67 := 0; fi67 < 2; fi67++ {
                        v3 = "a"
                        _ = fi67
                    }
                    fmt.Println(v6, v2, v2)
                    _ = fi65
                }
            } else {
                if ((v6 / ((v6) | 1)) >= 416) {
                    v1[1] = (908 == 0)
                    var v68 bool = false
                    _ = v68
                } else {
                    for fi69 := 0; fi69 < 5; fi69++ {
                        var v70 string = v3
                        v71 := GenS3{F0: v6}
                        _ = fi69
                        _ = v70
                        _ = v71
                    }
                    for fi72 := 0; fi72 < 15; fi72++ {
                        v2 = v2
                        v6 -= 0
                        _ = fi72
                    }
                    v73 := GenS1{F0: (!(!(false || (!false)))), F1: 664}
                    _ = v73
                }
                v74 := GenS3{F0: v6}
                v2 = float32(float64(float32(1.5)))
                _ = v74
            }
            v6 = v6
            v75 := GenS3{F0: (v6 % (((v6 + (v6 << ((18446744073709551615) & 63)))) | 1))}
            _ = v75
        }
    }
    for fi76 := 0; fi76 < 1; fi76++ {
        fmt.Println(v2, v6)
        fmt.Println(v6)
        _ = fi76
    }
    for fi77 := 0; fi77 < 17; fi77++ {
        switch 127 {
        case 579:
            v78 := GenS4{F0: v3, F1: (v3 + v3)}
            _ = v78
        case 2147483647:
            var v79 uint = 255
            _ = v79
        default:
            fmt.Println(v3, v2, fi77)
        }
        _ = fi77
    }
    for fi80 := 0; fi80 < 14; fi80++ {
        if len(v5) > 0 {
            v5[0] = 2147483647
        }
        _ = fi80
    }
    fmt.Println(v2, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
