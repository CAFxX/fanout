package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 uint64; F1 int16; F2 uint }
type GenS5 struct { F0 uint8 }

func main() {
    v1 := []int{964, -1, 9223372036854775807, 887}
    v2 := []string{("\n\t" + ("" + ("hello" + "")))}
    v3 := []string{"\n\t", "\n\t", "a"}
    var v4 float32 = ((1.5 / (0.1 / float32(3.14159))) - 1.5)
    v5 := make(map[int]bool)
    v6 := []string{"hello", "a", (("\n\t" + (("" + "hello") + "")) + "a"), ""}
    v4 -= v4
    _ = len(v5)
    v7 := GenS1{F0: ("\n\t" + "\n\t"), F1: -9223372036854775808}
    v8 := GenS1{F0: "", F1: 694}
    if false {
        var v9 int8 = 42
        v3 = append(v3, "a")
        _ = v9
    }
    for fi10 := 0; fi10 < 15; fi10++ {
        v4 = 0.0
        if (((fi10 / (((-(fi10))) | 1)) == fi10) || (18446744073709551615 == 18446744073709551615)) {
            _ = len(v5)
        } else {
            v4 += v4
        }
        _ = fi10
    }
    switch -602 {
    case 127:
        v11 := GenS2{F0: 18446744073709551615}
        _ = v11
    case 32768:
        var v12 uint64 = 2
        _ = v12
    case 1:
        v13 := GenS1{F0: ("\n\t" + "\n\t"), F1: 9223372036854775807}
        _ = v13
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        if false {
            delete(v5, (^(fi14)))
            for fi15 := 0; fi15 < 1; fi15++ {
                var v16 uint64 = 0
                if (fi14 > fi15) {
                    switch (("a" + ((("\n\t" + ("\n\t" + "foo bar")) + "hello") + (("foo bar" + "foo bar") + "hello"))) + "") {
                    case "":
                        switch (fi14 - fi14) {
                        case 2147483647:
                            _ = len(v5)
                            if false {
                                v17 := v5[fi15]
                                fmt.Println(v4, v16)
                                var v18 int32 = 801
                                _ = v17
                                _ = v18
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v19 := GenS2{F0: 4294967295}
                            var v20 uint16 = 1
                            v21 := GenS2{F0: v16}
                            v22 := GenS3{F0: -890, F1: 390}
                            switch "" {
                            case "a":
                                switch "" {
                                case "":
                                    fmt.Println(v20, fi14)
                                case "\n\t":
                                    switch "" {
                                    case "hello":
                                        v6 = append(v6, "")
                                    default:
                                        v2 = append(v2, "")
                                    }
                                }
                                v23 := GenS4{F0: 161, F1: 0, F2: 2}
                                _ = v23
                            case "hello":
                                v4 += v4
                            default:
                                var v24 uint32 = uint32(v20)
                                _ = v24
                            }
                            _ = v19
                            _ = v20
                            _ = v21
                            _ = v22
                        default:
                            if true {
                                var v25 uint8 = 71
                                if true {
                                    v25 = (^(v25))
                                } else {
                                    var v26 int8 = 127
                                    if (1 <= 0) {
                                        switch int(v25) {
                                        case 127:
                                            v27 := v5[(-(fi15))]
                                            _ = v27
                                        case 42:
                                            v16 -= (v16 & v16)
                                        case -109:
                                            v16 += 4294967295
                                        }
                                        switch fi14 {
                                        case 128:
                                            switch "a" {
                                            case "\n\t":
                                                if (907 < v16) {
                                                    if true {
                                                        v25 += (v25 << ((847) & 63))
                                                        v28 := GenS5{F0: uint8(v16)}
                                                        switch ("hello" + "") {
                                                        case "a":
                                                            _ = len(v3)
                                                        case "hello":
                                                            fmt.Println(v25, v26, v26)
                                                        }
                                                        _ = v28
                                                    } else {
                                                        v29 := GenS4{F0: 1, F1: -1, F2: 255}
                                                        v5[(^(fi15))] = (fi14 < (fi14 + (fi15 ^ fi14)))
                                                        switch "\n\t" {
                                                        case "hello":
                                                            var v30 string = ((("" + "") + ((("foo bar" + "hello") + "a") + "a")) + "hello")
                                                            _ = v30
                                                        default:
                                                            if len(v3) > 0 {
                                                                v3[0] = ""
                                                            }
                                                        }
                                                        if true {
                                                            var v31 int = fi15
                                                            v31 += v31
                                                            v32 := GenS4{F0: (^(v16)), F1: 32767, F2: 521}
                                                            _ = v31
                                                            _ = v32
                                                        }
                                                        v16 += v16
                                                        if false {
                                                            switch 4294967295 {
                                                            case 9223372036854775807:
                                                                v16 -= (v16 | (v16 + v16))
                                                            case 0:
                                                                var v33 uint32 = 664
                                                                _ = v33
                                                            }
                                                            v4 -= v4
                                                            var v34 int = fi14
                                                            v5[(fi14 << ((42) & 63))] = false
                                                            _ = v34
                                                        }
                                                        _ = v29
                                                    }
                                                }
                                            case "foo bar":
                                                v26 += v26
                                            default:
                                                var v35 int32 = 236
                                                _ = v35
                                            }
                                        }
                                        var v36 bool = false
                                        _ = v36
                                    } else {
                                        var v37 uint8 = uint8(v16)
                                        v26 = v26
                                        fmt.Println(fi14, fi14)
                                        switch (^(fi14)) {
                                        case 1:
                                            var v38 uint64 = (v16 - v16)
                                            _ = v38
                                        case -2147483648:
                                            var v39 uint = 65535
                                            _ = v39
                                        case -150:
                                            var v40 uint16 = 611
                                            _ = v40
                                        default:
                                            v16 -= (v16 ^ (^(v16)))
                                        }
                                        _ = v37
                                    }
                                    _ = v26
                                }
                                var v41 int = int(v25)
                                _ = v25
                                _ = v41
                            }
                            v42 := GenS3{F0: -884, F1: -9223372036854775808}
                            _ = v42
                        }
                    default:
                        if (true && (-1 <= 42)) {
                            v16 += v16
                            _ = len(v5)
                            if (!(!(!((-2.25 * (1.5 / 1.5)) <= 1.5)))) {
                                v4 = 1.5
                            }
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        var v43 uint16 = 32767
                        _ = v43
                    }
                    var v44 float32 = v4
                    var v45 int = (fi14 | fi14)
                    _ = v44
                    _ = v45
                }
                _ = fi15
                _ = v16
            }
            var v46 uint32 = 4294967295
            _ = v46
        }
        for fi47 := 0; fi47 < 4; fi47++ {
            v4 += 0.1
            var v48 uint64 = 1
            _ = fi47
            _ = v48
        }
        if ((42 <= 42) || true) {
            for fi49 := 0; fi49 < 11; fi49++ {
                if false {
                    v3 = append(v3, "\n\t")
                    switch ("a" + "") {
                    case "":
                        v4 += (0.1 / v4)
                    case "foo bar":
                        if false {
                            var v50 uint64 = 1
                            _ = v50
                        }
                        v51 := GenS1{F0: "foo bar", F1: int64(fi14)}
                        fmt.Println(v4, fi14)
                        _ = v51
                    default:
                        var v52 float32 = (3.14159 - -2.25)
                        _ = v52
                    }
                    v4 = -2.25
                    var v53 uint = 18446744073709551615
                    _ = v53
                } else {
                    fmt.Println(fi49, fi14, fi14)
                }
                _ = fi49
            }
            var v54 uint8 = 0
            var v55 uint = 42
            _ = v54
            _ = v55
        }
        _ = fi14
    }
    if false {
        v4 = 1.5
    }
    if true {
        var v56 uint64 = 18446744073709551615
        v4 = 0.0
        switch "\n\t" {
        case "\n\t":
            for fi57 := 0; fi57 < 1; fi57++ {
                if len(v2) > 0 {
                    v2[0] = ""
                }
                _ = fi57
            }
        case "hello":
            var v58 int64 = -1
            _ = v58
        }
        _ = v56
    } else {
        fmt.Println(v4)
        for fi59 := 0; fi59 < 17; fi59++ {
            if ((539 > 47) || false) {
                fmt.Println(v4)
                var v60 uint = 813
                _ = v60
            } else {
                if len(v1) > 0 {
                    v1[0] = (fi59 % ((fi59) | 1))
                }
                if (((1.5 - (0.0 + 0.1)) * float64(((v4 - 0.1) + v4))) <= (-2.25 * (((0.1 / 3.14159) / 1.0) / 1e10))) {
                    fmt.Println(fi59, v4)
                    v61 := GenS4{F0: 1, F1: 1, F2: 42}
                    var v62 uint8 = 193
                    _ = v61
                    _ = v62
                } else {
                    if true {
                        var v63 int16 = 985
                        fmt.Println(v4, v4)
                        _ = v63
                    } else {
                        for fi64 := 0; fi64 < 19; fi64++ {
                            _ = len(v5)
                            v4 += v4
                            _ = fi64
                        }
                        if true {
                            switch ((("foo bar" + ("foo bar" + "foo bar")) + "a") + "\n\t") {
                            case "foo bar":
                                fmt.Println(fi59, fi59)
                            case "":
                                if false {
                                    fmt.Println(fi59)
                                } else {
                                    if (127 > -128) {
                                        for fi65 := 0; fi65 < 13; fi65++ {
                                            if (1 >= -306) {
                                                v4 = v4
                                                var v66 string = "a"
                                                _ = v66
                                            }
                                            v4 -= float32(-2.25)
                                            var v67 int8 = -1
                                            _ = fi65
                                            _ = v67
                                        }
                                    }
                                    v4 += 0.0
                                }
                            case "a":
                                var v68 uint8 = 1
                                _ = v68
                            default:
                                v4 = (v4 + 0.1)
                            }
                            v4 = 1e10
                            var v69 uint32 = 2147483647
                            _ = v69
                        }
                        fmt.Println(fi59, v4)
                    }
                    var v70 float32 = 0.0
                    for fi71 := 0; fi71 < 10; fi71++ {
                        var v72 int64 = 42
                        _ = fi71
                        _ = v72
                    }
                    _ = v70
                }
            }
            _ = fi59
        }
        v73 := GenS4{F0: 586, F1: 0, F2: 928}
        _ = v73
    }
    var v74 uint16 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v74
    _ = v8
}
