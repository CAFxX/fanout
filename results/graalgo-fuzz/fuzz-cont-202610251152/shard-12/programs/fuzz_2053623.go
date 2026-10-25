package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 int8 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int; F1 int16; F2 string }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 string; F1 uint8; F2 int16 }
type GenS9 struct { F0 uint32; F1 int64 }
type GenS10 struct { F0 uint; F1 int16; F2 string }

func main() {
    v1 := []uint32{1, 704, 2}
    v2 := []uint32{914}
    var v3 uint8 = 255
    v3 = (^(v3))
    var v4 float32 = (-2.25 / -2.25)
    var v5 uint16 = 32767
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi6 := 0; fi6 < 11; fi6++ {
        if len(v1) > 0 {
            v1[0] = 65535
        }
        _ = fi6
    }
    if (!false) {
        var v7 uint = 0
        _ = v7
    } else {
        switch -48 {
        case 4294967295:
            for fi8 := 0; fi8 < 16; fi8++ {
                v9 := GenS1{F0: (v5 >> ((385) & 63)), F1: false, F2: 0}
                fmt.Println(v5, v5)
                _ = fi8
                _ = v9
            }
        case 9223372036854775807:
            var v10 uint32 = 65535
            _ = v10
        case 127:
            var v11 uint16 = (v5 >> ((42) & 63))
            _ = v11
        }
    }
    if (777 < (v5 / (((^(v5))) | 1))) {
        fmt.Println(v3, v3)
        for fi12 := 0; fi12 < 19; fi12++ {
            if (42 <= -1) {
                if (true && false) {
                    v13 := GenS1{F0: v5, F1: false, F2: 127}
                    v3 = (v3 * (v3 % (((v3 ^ v3)) | 1)))
                    v2 = append(v2, 1)
                    _ = v13
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v14 := GenS1{F0: (v5 << ((2) & 63)), F1: (("" != (("" + ("\n\t" + "foo bar")) + "hello")) && true), F2: 1}
                    _ = v14
                }
                if ((v5 + v5) >= (^(v5))) {
                    var v15 int32 = 0
                    v3 += v3
                    for fi16 := 0; fi16 < 2; fi16++ {
                        v17 := GenS2{F0: -846, F1: (fi16 + -879)}
                        if (false && true) {
                            fmt.Println(v3, v3, fi16)
                        } else {
                            v18 := GenS3{F0: -457}
                            if (true && ((uint8(v5) >= (v3 + v3)) && false)) {
                                var v19 uint16 = v5
                                v20 := GenS4{F0: fi12, F1: 42, F2: ("hello" + "")}
                                var v21 bool = ((!false) && (1 != 1))
                                _ = v19
                                _ = v20
                                _ = v21
                            } else {
                                var v22 int32 = v15
                                v23 := GenS5{F0: (fi16 | (^(fi12)))}
                                _ = v22
                                _ = v23
                            }
                            _ = v18
                        }
                        fmt.Println(fi16, fi16)
                        _ = fi16
                        _ = v17
                    }
                    _ = v15
                } else {
                    if true {
                        fmt.Println(v4, v3, v5)
                        for fi24 := 0; fi24 < 15; fi24++ {
                            switch "a" {
                            case "hello":
                                if false {
                                    if (0.0 >= (0.1 * ((0.1 + 3.14159) + 0.0))) {
                                        v25 := GenS1{F0: v5, F1: (2 == 0), F2: -12}
                                        _ = v25
                                    } else {
                                        fmt.Println(fi24)
                                        fmt.Println(fi12)
                                    }
                                    v4 = -2.25
                                } else {
                                    v4 -= ((v4 + v4) + (3.14159 / v4))
                                    if (65535 != v5) {
                                        v26 := GenS6{F0: 4294967295, F1: false, F2: v5}
                                        if (4294967295 < 424) {
                                            if ((v5 / ((32767) | 1)) == (^(v5))) {
                                                if true {
                                                    switch 127 {
                                                    case 578:
                                                        if (((v3 % ((141) | 1)) <= v3) && false) {
                                                            if false {
                                                                v1 = append(v1, 853)
                                                                var v27 float64 = (-2.25 + -2.25)
                                                                if (!(2 != uint8(v5))) {
                                                                    v27 = 1.5
                                                                    var v28 uint16 = (v5 % (((v5 * 741)) | 1))
                                                                    switch ((("hello" + "foo bar") + "\n\t") + ("" + "hello")) {
                                                                    case "\n\t":
                                                                        var v29 float32 = v4
                                                                        _ = v29
                                                                    case "foo bar":
                                                                        fmt.Println(v27, v3)
                                                                    default:
                                                                        v30 := GenS7{F0: 65535}
                                                                        _ = v30
                                                                    }
                                                                    _ = v28
                                                                }
                                                                v3 -= v3
                                                                v4 += v4
                                                                v3 = 255
                                                                if true {
                                                                    switch "\n\t" {
                                                                    case "hello":
                                                                        var v31 int8 = 127
                                                                        _ = v31
                                                                    case "":
                                                                        var v32 string = "\n\t"
                                                                        v33 := GenS2{F0: 80, F1: -444}
                                                                        fmt.Println(v4, fi24)
                                                                        _ = v32
                                                                        _ = v33
                                                                    }
                                                                    if (-128 == -1) {
                                                                        if len(v1) > 0 {
                                                                            v1[0] = 0
                                                                        }
                                                                    }
                                                                }
                                                                _ = v27
                                                            }
                                                            fmt.Println(fi12)
                                                        } else {
                                                            var v34 uint = 255
                                                            _ = v34
                                                        }
                                                    case -1:
                                                        var v35 int = -885
                                                        _ = v35
                                                    case -405:
                                                        var v36 int32 = 585
                                                        _ = v36
                                                    }
                                                    v5 = (^(v5))
                                                } else {
                                                    v4 -= 1.5
                                                    if false {
                                                        fmt.Println(fi24, fi12, v5)
                                                        var v37 uint8 = (v3 - (v3 / (((v3 / ((v3) | 1))) | 1)))
                                                        v2 = append(v2, 0)
                                                        var v38 uint = 943
                                                        v39 := GenS6{F0: uint64(v37), F1: false, F2: v5}
                                                        var v40 int16 = 42
                                                        _ = v37
                                                        _ = v38
                                                        _ = v39
                                                        _ = v40
                                                    } else {
                                                        var v41 float64 = -2.25
                                                        var v42 float64 = (float64(1.5) + v41)
                                                        _ = v41
                                                        _ = v42
                                                    }
                                                    fmt.Println(fi12)
                                                    v43 := GenS8{F0: "hello", F1: v3, F2: 1}
                                                    _ = v43
                                                }
                                                if (2 > 4294967295) {
                                                    v44 := GenS5{F0: 42}
                                                    if false {
                                                        fmt.Println(fi12, fi24)
                                                    } else {
                                                        switch (fi24 % ((128) | 1)) {
                                                        case 32768:
                                                            var v45 int16 = 32767
                                                            _ = v45
                                                        default:
                                                            v5 = (v5 - (^(v5)))
                                                        }
                                                        var v46 int64 = 9223372036854775807
                                                        _ = v46
                                                    }
                                                    _ = v44
                                                }
                                            }
                                            _ = len(v1)
                                        }
                                        if len(v2) > 0 {
                                            v2[0] = 157
                                        }
                                        _ = v26
                                    } else {
                                        v4 = v4
                                        var v47 int64 = 9223372036854775807
                                        var v48 uint32 = 2147483647
                                        _ = v47
                                        _ = v48
                                    }
                                }
                            case "foo bar":
                                if true {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    if (v4 > float32((3.14159 / 1.0))) {
                                        v4 -= v4
                                        v5 = v5
                                        switch (fi24 << ((912) & 63)) {
                                        case 128:
                                            var v49 bool = (677 > 42)
                                            _ = v49
                                        case 482:
                                            switch ("hello" + "foo bar") {
                                            case "foo bar":
                                                v4 = 3.14159
                                            default:
                                                v50 := GenS8{F0: "\n\t", F1: v3, F2: -408}
                                                _ = v50
                                            }
                                            if true {
                                                v4 += v4
                                            }
                                        default:
                                            if (uint32(v5) >= 65535) {
                                                if true {
                                                    v51 := GenS4{F0: -2147483648, F1: 42, F2: "hello"}
                                                    _ = v51
                                                }
                                                var v52 uint = 0
                                                v1 = append(v1, 1)
                                                v3 = (v3 % ((v3) | 1))
                                                _ = v52
                                            }
                                        }
                                    } else {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                    v2 = append(v2, 42)
                                } else {
                                    if true {
                                        switch fi24 {
                                        case 127:
                                            v5 = 967
                                        default:
                                            v53 := GenS6{F0: 2, F1: (988 <= uint(v5)), F2: (^(v5))}
                                            _ = v53
                                        }
                                    }
                                    v3 = v3
                                    v3 = 243
                                }
                            }
                            v54 := GenS8{F0: (("a" + ("hello" + "a")) + "a"), F1: v3, F2: 0}
                            _ = fi24
                            _ = v54
                        }
                        var v55 uint32 = 1
                        _ = v55
                    } else {
                        var v56 int64 = 42
                        v57 := GenS9{F0: 751, F1: v56}
                        v1 = append(v1, 627)
                        _ = v56
                        _ = v57
                    }
                    v58 := GenS10{F0: 0, F1: 879, F2: "hello"}
                    switch int(v3) {
                    case -1:
                        v59 := GenS5{F0: fi12}
                        _ = v59
                    case 128:
                        if len(v1) > 0 {
                            v1[0] = uint32(v5)
                        }
                    }
                    _ = v58
                }
            }
            var v60 float32 = (0.1 * float32(0.1))
            _ = fi12
            _ = v60
        }
    }
    if len(v2) > 0 {
        v2[0] = 65535
    }
    for fi61 := 0; fi61 < 3; fi61++ {
        var v62 float64 = 1e10
        _ = fi61
        _ = v62
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
