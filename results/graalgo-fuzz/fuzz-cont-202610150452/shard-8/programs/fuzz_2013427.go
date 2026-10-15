package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 int16; F2 int32 }
type GenS3 struct { F0 uint32; F1 int }
type GenS4 struct { F0 int8; F1 int8; F2 string }
func genFn0(p0 int16, p1 uint) uint16 {
    if (!true) {
        for fi5 := 0; fi5 < 20; fi5++ {
            if (((!true) || true) || true) {
                for fi6 := 0; fi6 < 6; fi6++ {
                    p0 -= p0
                    _ = fi6
                }
                if (fi5 >= fi5) {
                    fmt.Println(p0, p1, fi5)
                    v7 := []int{(fi5 * (fi5 | (^(fi5)))), (fi5 >> ((p1) & 63))}
                    _ = v7
                } else {
                    for fi8 := 0; fi8 < 1; fi8++ {
                        var v9 uint32 = 2
                        _ = fi8
                        _ = v9
                    }
                }
                fmt.Println(p0, p0)
            } else {
                var v10 float32 = ((3.14159 + (1e10 * 0.1)) / 1.0)
                _ = v10
            }
            v11 := []int{(fi5 ^ (fi5 >> (((^(p1))) & 63)))}
            if false {
                var v12 uint16 = 42
                if len(v11) > 0 {
                    v11[0] = (fi5 | (fi5 - (fi5 + fi5)))
                }
                _ = v12
            }
            _ = fi5
            _ = v11
        }
        p1 += (p1 | 0)
    } else {
        v13 := []int{-364}
        var v14 int16 = -32768
        _ = v13
        _ = v14
    }
    var v15 uint64 = 324
    if true {
        v16 := GenS1{F0: 100}
        _ = v16
    }
    v17 := GenS2{F0: "a", F1: 0, F2: 1}
    _ = p0
    _ = p1
    _ = v15
    _ = v17
    return 32767
}
func genFn1(p0 uint64) int16 {
    switch "hello" {
    case "hello":
        if (false || false) {
            p0 -= 42
            switch "\n\t" {
            case "a":
                for fi18 := 0; fi18 < 9; fi18++ {
                    p0 = 0
                    var v19 int = fi18
                    switch ("" + "\n\t") {
                    case "a":
                        p0 = (p0 << ((726) & 63))
                    default:
                        if true {
                            fmt.Println(p0)
                        }
                    }
                    _ = fi18
                    _ = v19
                }
            }
            p0 = (^(p0))
        } else {
            for fi20 := 0; fi20 < 17; fi20++ {
                p0 = p0
                v21 := GenS2{F0: "", F1: 32767, F2: -1}
                _ = fi20
                _ = v21
            }
        }
    }
    var v22 int8 = 100
    _ = p0
    _ = v22
    return 1
}

func main() {
    var v1 int8 = 100
    var v2 uint = 65535
    var v3 string = "foo bar"
    var v4 int = 1
    if (2 < 127) {
        switch 4294967295 {
        case 128:
            v23 := []int{}
            _ = v23
        case -722:
            switch v3 {
            case "hello":
                var v24 int = v4
                _ = v24
            default:
                var v25 uint64 = 669
                _ = v25
            }
        default:
            v3 = (v3 + "hello")
        }
    }
    v26 := []int{}
    if (true && (2 != 127)) {
        switch 127 {
        case 4294967295:
            for fi27 := 0; fi27 < 6; fi27++ {
                var v28 uint = 4294967295
                var v29 uint8 = 42
                _ = fi27
                _ = v28
                _ = v29
            }
        case 591:
            fmt.Println(v1, v4, v4)
        }
        fmt.Println(v1, v4)
    } else {
        if (v2 <= (^(v2))) {
            v26 = append(v26, v4)
        }
        switch (v4 % ((32768) | 1)) {
        case 354:
            for fi30 := 0; fi30 < 2; fi30++ {
                var v31 int16 = -32768
                v32 := GenS2{F0: "\n\t", F1: v31, F2: -1}
                for fi33 := 0; fi33 < 19; fi33++ {
                    if (9223372036854775807 > 42) {
                        fmt.Println(fi33, v4)
                        v1 = (v1 & v1)
                        var v34 int32 = -1
                        _ = v34
                    } else {
                        v35 := GenS3{F0: 1, F1: (fi33 << (((v2 + 255)) & 63))}
                        fmt.Println(v1, fi33, v4)
                        _ = v35
                    }
                    _ = fi33
                }
                _ = fi30
                _ = v31
                _ = v32
            }
        default:
            if true {
                var v36 int64 = 1
                v36 = v36
                _ = v36
            } else {
                var v37 string = v3
                switch v3 {
                case "":
                    for fi38 := 0; fi38 < 5; fi38++ {
                        switch ((v3 + v3) + v37) {
                        case "hello":
                            for fi39 := 0; fi39 < 1; fi39++ {
                                v26 = append(v26, 9223372036854775807)
                                fmt.Println(v37, v3, v3)
                                switch fi39 {
                                case 42:
                                    v40 := GenS3{F0: 4294967295, F1: 1}
                                    _ = v40
                                case 32768:
                                    var v41 int32 = 0
                                    _ = v41
                                }
                                _ = fi39
                            }
                        case "":
                            var v42 int64 = -856
                            _ = v42
                        default:
                            fmt.Println(v1, v4, fi38)
                        }
                        var v43 float64 = (((0.1 * (1.5 - (((0.1 - (1e10 / 1.0)) / float64(1.5)) / -2.25))) - (0.0 / 1e10)) - 0.0)
                        v44 := GenS2{F0: v3, F1: genFn1(18446744073709551615), F2: -2147483648}
                        _ = fi38
                        _ = v43
                        _ = v44
                    }
                default:
                    var v45 int32 = -516
                    _ = v45
                }
                _ = v37
            }
        }
        if true {
            fmt.Println(v1, v3)
        } else {
            fmt.Println(v3, v4)
        }
    }
    v26 = append(v26, (v4 * v4))
    var v46 int16 = 42
    v47 := GenS3{F0: 4294967295, F1: (v4 / (((v4 & 4294967295)) | 1))}
    v48 := GenS1{F0: (v1 >> ((65535) & 63))}
    v1 -= (-(v1))
    if (!((580 != 2) && false)) {
        v4 = (v4 * v4)
        var v49 int = (v4 - (^(v4)))
        fmt.Println(v4)
        _ = v49
    } else {
        if (!false) {
            var v50 uint8 = 42
            if ((false && true) || (255 > (^(v2)))) {
                fmt.Println(v2, v50, v3)
                v2 = (v2 + 18446744073709551615)
                if (v3 != v3) {
                    switch (v4 | -348) {
                    case -2147483648:
                        if (1 >= 42) {
                            fmt.Println(v2, v3, v4)
                        }
                    case 32768:
                        fmt.Println(v50, v2, v1)
                    case 349:
                        fmt.Println(v50, v46)
                    default:
                        switch (v3 + "\n\t") {
                        case "a":
                            if (true || (true && (-2.25 <= 0.0))) {
                                if len(v26) >= 2 {
                                    v26 = v26[1:len(v26)]
                                }
                                var v51 uint64 = 42
                                v3 = v3
                                _ = v51
                            }
                        case "":
                            _ = len(v26)
                        case "hello":
                            v4 = (v4 * v4)
                        default:
                            var v52 float32 = float32(((1.5 / (1e10 - (-2.25 - 1.5))) / 1.5))
                            _ = v52
                        }
                    }
                    if (false || (!false)) {
                        if true {
                            if (true && (true && (!(((!(!(!(("\n\t" + (v3 + ("\n\t" + ((v3 + v3) + ("" + (v3 + v3)))))) <= (v3 + ("\n\t" + "\n\t")))))) || false) && (2147483647 > 4294967295))))) {
                                if (42 >= 0) {
                                    v46 = 710
                                }
                                fmt.Println(v4)
                                for fi53 := 0; fi53 < 15; fi53++ {
                                    for fi54 := 0; fi54 < 19; fi54++ {
                                        fmt.Println(v1, v3, v1)
                                        _ = fi54
                                    }
                                    v2 -= (v2 + (^(v2)))
                                    var v55 float64 = 1e10
                                    _ = fi53
                                    _ = v55
                                }
                            } else {
                                fmt.Println(v4, v4, v4)
                                v2 = 4294967295
                                var v56 uint32 = 2147483648
                                _ = v56
                            }
                            if len(v26) >= 2 {
                                v26 = v26[1:len(v26)]
                            }
                            if false {
                                var v57 uint8 = v50
                                _ = v57
                            }
                        } else {
                            if len(v26) > 0 {
                                v26[0] = (v4 >> (((v2 << (((v2 & v2)) & 63))) & 63))
                            }
                            v4 += v4
                        }
                        for fi58 := 0; fi58 < 5; fi58++ {
                            var v59 float32 = 0.1
                            v2 -= v2
                            v4 = (fi58 & (v4 & 42))
                            _ = fi58
                            _ = v59
                        }
                        v26 = append(v26, (v4 - 128))
                    }
                    if (-1 < 0) {
                        for fi60 := 0; fi60 < 5; fi60++ {
                            var v61 uint64 = 365
                            var v62 int8 = v1
                            _ = fi60
                            _ = v61
                            _ = v62
                        }
                    } else {
                        v2 = (^(v2))
                    }
                } else {
                    if false {
                        v63 := GenS4{F0: (v1 << ((4294967295) & 63)), F1: 1, F2: (((("hello" + (v3 + v3)) + "a") + (v3 + v3)) + (v3 + "hello"))}
                        _ = len(v26)
                        _ = v63
                    } else {
                        v64 := GenS2{F0: "hello", F1: -1, F2: 1}
                        _ = v64
                    }
                    if len(v26) > 0 {
                        v26[0] = (v4 | -2147483648)
                    }
                    v46 = genFn1(0)
                }
            }
            v65 := GenS2{F0: "hello", F1: genFn1(4294967295), F2: 2147483647}
            _ = v50
            _ = v65
        }
        switch (v3 + (v3 + (v3 + v3))) {
        case "\n\t":
            v4 = (v4 << ((65535) & 63))
        case "foo bar":
            var v66 uint64 = 0
            _ = v66
        default:
            fmt.Println(v2, v3)
        }
        if len(v26) > 0 {
            v26[0] = -356
        }
    }
    _ = len(v26)
    if false {
        for fi67 := 0; fi67 < 2; fi67++ {
            var v68 uint16 = 0
            _ = fi67
            _ = v68
        }
        var v69 float64 = 1e10
        _ = v69
    }
    switch 4294967295 {
    case -2147483648:
        var v70 uint64 = 752
        _ = v70
    }
    switch v4 {
    case -1:
        for fi71 := 0; fi71 < 11; fi71++ {
            if true {
                v3 = v3
                for fi72 := 0; fi72 < 15; fi72++ {
                    v73 := GenS4{F0: (v1 | 42), F1: (v1 % (((v1 | v1)) | 1)), F2: "\n\t"}
                    var v74 float64 = 0.1
                    _ = fi72
                    _ = v73
                    _ = v74
                }
            } else {
                _ = len(v26)
            }
            switch "hello" {
            case "a":
                if len(v26) >= 2 {
                    v26 = v26[1:len(v26)]
                }
            }
            _ = len(v26)
            _ = fi71
        }
    case -415:
        v26 = append(v26, (-(v4)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
}
