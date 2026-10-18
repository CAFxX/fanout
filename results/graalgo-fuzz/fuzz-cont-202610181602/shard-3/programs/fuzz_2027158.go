package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint; F1 uint; F2 uint }
type GenS5 struct { F0 uint8; F1 uint }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 string; F1 uint }
type GenS8 struct { F0 uint16; F1 int64; F2 uint32 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int8; F1 int }
func genFn0(p0 int32) int16 {
    var v5 bool = (!(-32768 < 0))
    v6 := GenS1{F0: 0, F1: 128, F2: 28}
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 int64 = 1
        v5 = ((v5 || (164 != 2)) || (!((0.1 / -2.25) > 0.1)))
        for fi9 := 0; fi9 < 6; fi9++ {
            fmt.Println(v5)
            v8 = 0
            v8 = v8
            p0 = p0
            v10 := []int{-1, 42}
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    if (42 == 1) {
        p0 -= -1
        p0 = (p0 - -2147483648)
        if v5 {
            if (!v5) {
                var v11 int8 = 1
                var v12 int64 = -1
                _ = v11
                _ = v12
            }
        } else {
            if false {
                if ((p0 - p0) <= p0) {
                    for fi13 := 0; fi13 < 20; fi13++ {
                        p0 -= (-(p0))
                        fmt.Println(fi13)
                        v14 := GenS1{F0: -128, F1: 2, F2: -128}
                        _ = fi13
                        _ = v14
                    }
                    v15 := GenS1{F0: 0, F1: 127, F2: -128}
                    var v16 uint = 18446744073709551615
                    _ = v15
                    _ = v16
                } else {
                    if (v5 && true) {
                        v17 := GenS2{F0: 2}
                        _ = v17
                    } else {
                        var v18 int = 128
                        _ = v18
                    }
                    for fi19 := 0; fi19 < 13; fi19++ {
                        v20 := []int{(fi19 & (fi19 * (fi19 / (((fi19 / (((^(fi19))) | 1))) | 1))))}
                        var v21 int64 = 42
                        fmt.Println(v5, p0)
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    v22 := GenS2{F0: 4294967295}
                    _ = v22
                }
            }
            v5 = ((("" + ("\n\t" + "foo bar")) + "foo bar") <= "hello")
        }
    }
    _ = p0
    _ = v5
    _ = v6
    return -1
}
func genFn1() uint16 {
    var v23 int8 = -1
    switch "a" {
    case "hello":
        for fi24 := 0; fi24 < 10; fi24++ {
            for fi25 := 0; fi25 < 5; fi25++ {
                v26 := []int{(fi24 << ((1) & 63))}
                _ = fi25
                _ = v26
            }
            var v27 uint8 = 0
            v28 := GenS2{F0: 0}
            _ = fi24
            _ = v27
            _ = v28
        }
    case "\n\t":
        v29 := GenS2{F0: 8}
        _ = v29
    default:
        var v30 float32 = -2.25
        _ = v30
    }
    v23 = -128
    v23 = (v23 * v23)
    _ = v23
    return 32767
}

func main() {
    var v1 int16 = 42
    var v2 bool = ((235 >= 1) && ((0.0 / 3.14159) == 1e10))
    var v3 string = "hello"
    v4 := make(map[string]bool)
    v31 := []int{-901, 128, -9223372036854775808}
    var v32 string = v3
    if (((v3 + v3) < "a") && v2) {
        v2 = (103 < 247)
        var v33 uint = 1
        for fi34 := 0; fi34 < 8; fi34++ {
            switch "" {
            case "":
                v31 = append(v31, 127)
            case "a":
                v32 = "foo bar"
            case "\n\t":
                var v35 uint = v33
                _ = v35
            }
            switch (v32 + v3) {
            case "\n\t":
                v36 := v4[""]
                _ = v36
            default:
                _ = len(v31)
            }
            var v37 int32 = -72
            _ = fi34
            _ = v37
        }
        _ = v33
    } else {
        for fi38 := 0; fi38 < 1; fi38++ {
            v32 = v32
            v32 = v3
            _ = fi38
        }
    }
    for fi39 := 0; fi39 < 18; fi39++ {
        if (((2 == -1) && (!((-2.25 - 3.14159) <= (1e10 * (float64(1.5) * 1e10))))) || ((2 == 65535) || v2)) {
            var v40 int = -2147483648
            var v41 bool = true
            _ = v40
            _ = v41
        }
        v1 -= -239
        _ = fi39
    }
    if true {
        v2 = true
        v42 := GenS3{F0: 18446744073709551615}
        fmt.Println(v32)
        _ = v42
    }
    v32 = v3
    for fi43 := 0; fi43 < 10; fi43++ {
        v2 = v2
        if v2 {
            v3 = "foo bar"
            v44 := GenS4{F0: 42, F1: 1, F2: 65535}
            _ = v44
        }
        _ = fi43
    }
    for fi45 := 0; fi45 < 20; fi45++ {
        var v46 bool = ((-2.25 / (1.5 / 1.5)) != 1e10)
        v46 = (255 == 149)
        _ = fi45
        _ = v46
    }
    if (1 == 2) {
        switch ("" + v32) {
        case "foo bar":
            for fi47 := 0; fi47 < 15; fi47++ {
                v48 := GenS5{F0: 2, F1: 120}
                _ = fi47
                _ = v48
            }
        case "a":
            for fi49 := 0; fi49 < 12; fi49++ {
                v4["a"] = (!false)
                if (1 > 294) {
                    if len(v31) >= 2 {
                        v31 = v31[1:len(v31)]
                    }
                    if (1 < 1) {
                        for fi50 := 0; fi50 < 2; fi50++ {
                            switch "" {
                            case "\n\t":
                                v51 := GenS6{F0: -1}
                                _ = v51
                            case "a":
                                if (3.14159 >= 0.1) {
                                    v32 = ((("hello" + "") + v32) + "a")
                                } else {
                                    if len(v31) >= 2 {
                                        v31 = v31[1:len(v31)]
                                    }
                                    var v52 float32 = ((1e10 + 3.14159) - (1e10 + 3.14159))
                                    _ = v52
                                }
                            default:
                                if (-9223372036854775808 <= fi50) {
                                    switch (fi49 / (((fi49 / ((128) | 1))) | 1)) {
                                    case -735:
                                        if (v2 || (0 > 56)) {
                                            v4["a"] = (!false)
                                            _ = len(v31)
                                        } else {
                                            v3 = (v32 + ("\n\t" + v3))
                                        }
                                    case 127:
                                        var v53 float32 = float32((float64(1.5) + (1e10 + float64(0.1))))
                                        _ = v53
                                    case 9223372036854775807:
                                        var v54 int64 = -9223372036854775808
                                        _ = v54
                                    default:
                                        v32 = v32
                                    }
                                    var v55 int8 = -57
                                    _ = v55
                                } else {
                                    var v56 int32 = -639
                                    var v57 uint64 = 0
                                    if (!false) {
                                        v32 = ("a" + "hello")
                                        var v58 float64 = (0.1 * (((1.5 + 3.14159) / float64((3.14159 * (3.14159 * 1e10)))) + ((0.1 / -2.25) / -2.25)))
                                        _ = v58
                                    } else {
                                        fmt.Println(fi50)
                                        fmt.Println(v56, fi50)
                                    }
                                    _ = v56
                                    _ = v57
                                }
                            }
                            v59 := GenS5{F0: 220, F1: 258}
                            v32 = "foo bar"
                            _ = fi50
                            _ = v59
                        }
                        v2 = (v2 && (423 != 255))
                    } else {
                        v60 := GenS7{F0: v32, F1: 308}
                        fmt.Println(v2)
                        _ = v60
                    }
                } else {
                    if v2 {
                        switch "foo bar" {
                        case "a":
                            var v61 uint = 2
                            _ = v61
                        case "\n\t":
                            for fi62 := 0; fi62 < 7; fi62++ {
                                v3 = "hello"
                                _ = fi62
                            }
                        }
                        v1 += (v1 / ((genFn0(2147483647)) | 1))
                        for fi63 := 0; fi63 < 11; fi63++ {
                            var v64 uint64 = 2
                            _ = fi63
                            _ = v64
                        }
                    }
                    v32 = (v3 + v3)
                }
                _ = fi49
            }
        case "hello":
            var v65 uint64 = 699
            _ = v65
        default:
            var v66 uint8 = 128
            _ = v66
        }
        var v67 uint = 2
        _ = v67
    }
    var v68 string = v3
    fmt.Println(v2, v2)
    if v2 {
        if len(v31) > 0 {
            v31[0] = -179
        }
        if (!true) {
            for fi69 := 0; fi69 < 15; fi69++ {
                fmt.Println(v32, v32)
                _ = fi69
            }
            if (!v2) {
                for fi70 := 0; fi70 < 5; fi70++ {
                    for fi71 := 0; fi71 < 20; fi71++ {
                        var v72 uint = 4294967295
                        if ("hello" <= (v32 + (("\n\t" + "") + "hello"))) {
                            v68 = v32
                            v32 = "foo bar"
                        }
                        v73 := GenS8{F0: 32767, F1: 0, F2: 287}
                        _ = fi71
                        _ = v72
                        _ = v73
                    }
                    _ = fi70
                }
                for fi74 := 0; fi74 < 16; fi74++ {
                    v75 := GenS9{F0: -127}
                    switch (fi74 << ((65535) & 63)) {
                    case 42:
                        if (0 == 42) {
                            v76 := GenS10{F0: (v68 + "\n\t")}
                            _ = v76
                        }
                    case 672:
                        for fi77 := 0; fi77 < 13; fi77++ {
                            var v78 uint8 = 125
                            _ = fi77
                            _ = v78
                        }
                    }
                    var v79 int = fi74
                    _ = fi74
                    _ = v75
                    _ = v79
                }
            }
            v80 := GenS11{F0: 100, F1: -315}
            _ = v80
        }
        fmt.Println(v1, v1, v32)
    }
    v81 := GenS9{F0: 1}
    v82 := v4["hello"]
    v83 := GenS8{F0: 42, F1: -800, F2: 2147483647}
    var v84 uint32 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v68:", v68)
    fmt.Println("v82:", v82)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v68
    _ = v81
    _ = v82
    _ = v83
    _ = v84
}
