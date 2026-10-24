package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 int16; F2 uint32 }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 uint16; F1 string }
type GenS5 struct { F0 int16; F1 uint }

func main() {
    var v1 uint16 = 1
    var v2 int = 32768
    var v3 float64 = 1e10
    for fi4 := 0; fi4 < 19; fi4++ {
        fmt.Println(v2, v1)
        v2 = fi4
        _ = fi4
    }
    v5 := []int{(v2 % ((v2) | 1)), 9223372036854775807, 2147483647, v2}
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 int16 = 42
        var v8 int32 = -1
        _ = fi6
        _ = v7
        _ = v8
    }
    if true {
        var v9 int64 = -506
        _ = v9
    }
    fmt.Println(v1, v1, v3)
    v10 := GenS1{F0: "\n\t"}
    var v11 bool = false
    for fi12 := 0; fi12 < 19; fi12++ {
        var v13 int64 = -9223372036854775808
        switch "a" {
        case "\n\t":
            if v11 {
                v5 = append(v5, v2)
            }
        }
        _ = fi12
        _ = v13
    }
    v3 = v3
    v14 := GenS2{F0: 1000, F1: 1000, F2: 1}
    switch "\n\t" {
    case "\n\t":
        if (v11 || v11) {
            v11 = (20 < -1)
            v11 = true
        }
    }
    v15 := GenS3{F0: (v2 % ((v2) | 1)), F1: v1}
    if len(v5) > 0 {
        v5[0] = v2
    }
    var v16 int32 = 2147483647
    switch v2 {
    case 0:
        switch v2 {
        case -9:
            if ((!v11) || false) {
                for fi17 := 0; fi17 < 2; fi17++ {
                    v1 += 1
                    _ = len(v5)
                    if true {
                        if v11 {
                            switch (fi17 | (v2 % ((v2) | 1))) {
                            case -370:
                                v11 = (-681 > (fi17 | v2))
                            default:
                                fmt.Println(v3, v1, fi17)
                            }
                            var v18 uint32 = 2147483648
                            _ = v18
                        } else {
                            fmt.Println(fi17)
                            var v19 int8 = 1
                            var v20 float64 = float64(1e10)
                            _ = v19
                            _ = v20
                        }
                        _ = len(v5)
                        v5 = append(v5, (fi17 / ((fi17) | 1)))
                    } else {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        fmt.Println(v16, v2)
                        switch ("a" + ((("\n\t" + "a") + "") + "\n\t")) {
                        case "hello":
                            v1 = v1
                        case "\n\t":
                            v2 = v2
                        case "":
                            var v21 bool = v11
                            _ = v21
                        default:
                            var v22 int = (v2 << ((42) & 63))
                            _ = v22
                        }
                    }
                    _ = fi17
                }
                _ = len(v5)
            }
        case 9223372036854775807:
            v16 += (v16 % ((v16) | 1))
        }
    case 42:
        _ = len(v5)
    case -662:
        for fi23 := 0; fi23 < 16; fi23++ {
            if (!true) {
                for fi24 := 0; fi24 < 4; fi24++ {
                    if v11 {
                        var v25 uint64 = 18446744073709551615
                        var v26 int = fi24
                        var v27 bool = (v11 && false)
                        switch (v26 - -1) {
                        case 4294967295:
                            var v28 uint8 = 1
                            _ = v28
                        case 128:
                            var v29 bool = (v11 || (42 <= 0))
                            _ = v29
                        case -469:
                            var v30 int32 = v16
                            _ = v30
                        default:
                            if ((938 <= 0) || (!(v1 != 65535))) {
                                if (!("" <= "a")) {
                                    if ((1e10 * 0.0) <= 0.0) {
                                        switch 726 {
                                        case 128:
                                            var v31 float32 = 0.0
                                            _ = v31
                                        default:
                                            var v32 uint32 = 2147483648
                                            _ = v32
                                        }
                                    } else {
                                        fmt.Println(v25)
                                        var v33 uint64 = (v25 % (((^(v25))) | 1))
                                        _ = v33
                                    }
                                }
                            }
                            v1 += 2
                            var v34 uint32 = 42
                            _ = v34
                        }
                        var v35 int = (fi23 * 452)
                        _ = v25
                        _ = v26
                        _ = v27
                        _ = v35
                    } else {
                        if v11 {
                            v36 := GenS2{F0: 32767, F1: 1, F2: 65535}
                            var v37 float32 = (3.14159 / 0.1)
                            _ = v36
                            _ = v37
                        }
                    }
                    _ = fi24
                }
                for fi38 := 0; fi38 < 2; fi38++ {
                    fmt.Println(v3)
                    _ = fi38
                }
            } else {
                var v39 int = v2
                if (v3 > -2.25) {
                    v3 = 1e10
                    v1 += (v1 ^ 65535)
                }
                v39 = (v39 * 0)
                _ = v39
            }
            var v40 int64 = 9223372036854775807
            var v41 uint64 = 42
            _ = fi23
            _ = v40
            _ = v41
        }
    default:
        for fi42 := 0; fi42 < 16; fi42++ {
            if (255 != 1) {
                _ = len(v5)
            }
            _ = fi42
        }
    }
    if (1e10 >= (-2.25 - 0.0)) {
        fmt.Println(v16, v16, v3)
        for fi43 := 0; fi43 < 15; fi43++ {
            switch "hello" {
            case "foo bar":
                var v44 string = "\n\t"
                _ = v44
            default:
                for fi45 := 0; fi45 < 5; fi45++ {
                    if ((18446744073709551615 != 2) || (true && true)) {
                        fmt.Println(v1, v3, v11)
                    } else {
                        v46 := GenS1{F0: ""}
                        _ = v46
                    }
                    switch "\n\t" {
                    case "hello":
                        if (v11 && (42 > -128)) {
                            v2 += fi45
                        } else {
                            if len(v5) > 0 {
                                v5[0] = 32767
                            }
                            if ("" <= "\n\t") {
                                var v47 int8 = 127
                                if v11 {
                                    v2 = (fi45 / ((fi43) | 1))
                                    var v48 string = ""
                                    var v49 string = "\n\t"
                                    _ = v48
                                    _ = v49
                                } else {
                                    fmt.Println(fi45, v16, v47)
                                    if (!false) {
                                        fmt.Println(v3, v1, fi45)
                                        switch ((("foo bar" + "foo bar") + "hello") + ("hello" + "a")) {
                                        case "foo bar":
                                            fmt.Println(v2, v1, v3)
                                        default:
                                            v2 -= 4294967295
                                        }
                                    }
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    v50 := GenS1{F0: ("hello" + ("foo bar" + "a"))}
                                    _ = v50
                                }
                                _ = v47
                            } else {
                                fmt.Println(v3)
                            }
                            var v51 int = (v2 / ((fi43) | 1))
                            _ = v51
                        }
                    case "foo bar":
                        switch 1 {
                        case -1:
                            switch "\n\t" {
                            case "":
                                v52 := GenS4{F0: (v1 >> ((42) & 63)), F1: "foo bar"}
                                _ = v52
                            default:
                                var v53 string = ("\n\t" + "foo bar")
                                _ = v53
                            }
                        case -232:
                            _ = len(v5)
                        case -2147483648:
                            var v54 uint64 = 1
                            _ = v54
                        default:
                            if (127 == 42) {
                                if (v11 || false) {
                                    switch fi43 {
                                    case 20:
                                        var v55 int = (fi43 - 1)
                                        _ = v55
                                    case 4294967295:
                                        fmt.Println(v11)
                                    default:
                                        if (2147483647 < 65535) {
                                            var v56 uint64 = 985
                                            if true {
                                                fmt.Println(v2)
                                                switch (v2 / ((-1) | 1)) {
                                                case 32767:
                                                    var v57 uint64 = (^(v56))
                                                    _ = v57
                                                default:
                                                    v2 -= (v2 % (((fi45 & (fi43 << ((42) & 63)))) | 1))
                                                }
                                                var v58 string = (("" + "a") + "")
                                                _ = v58
                                            }
                                            var v59 int8 = 127
                                            _ = v56
                                            _ = v59
                                        }
                                    }
                                } else {
                                    fmt.Println(fi45, fi43)
                                    var v60 int64 = 900
                                    _ = v60
                                }
                            } else {
                                var v61 int16 = -940
                                var v62 int16 = (^(v61))
                                _ = v61
                                _ = v62
                            }
                        }
                    }
                    _ = len(v5)
                    _ = fi45
                }
            }
            if len(v5) > 0 {
                v5[0] = (v2 % (((fi43 | v2)) | 1))
            }
            if v11 {
                for fi63 := 0; fi63 < 10; fi63++ {
                    v5 = append(v5, (fi63 >> ((0) & 63)))
                    v5 = append(v5, (v2 % (((fi43 - (^(fi43)))) | 1)))
                    var v64 float32 = (1e10 * 1e10)
                    _ = fi63
                    _ = v64
                }
                fmt.Println(fi43, v16)
            }
            _ = fi43
        }
    }
    if (v11 && false) {
        v16 = v16
        v3 -= v3
    } else {
        if v11 {
            switch ((("a" + "hello") + ((("" + "hello") + "a") + "\n\t")) + ("foo bar" + ("" + (("a" + (("" + "\n\t") + "\n\t")) + ((((((("a" + "") + "foo bar") + "a") + "") + ("hello" + (("a" + "foo bar") + (("foo bar" + "") + "\n\t")))) + "foo bar") + ("" + ("foo bar" + (("foo bar" + "") + ("foo bar" + (("a" + "foo bar") + "\n\t")))))))))) {
            case "a":
                for fi65 := 0; fi65 < 2; fi65++ {
                    if false {
                        var v66 string = "foo bar"
                        var v67 string = v66
                        _ = v66
                        _ = v67
                    } else {
                        if ((0.1 * 1.5) > 0.0) {
                            v68 := GenS1{F0: "\n\t"}
                            v1 = v1
                            _ = v68
                        }
                        v69 := GenS5{F0: -32768, F1: 2}
                        v11 = v11
                        _ = v69
                    }
                    switch ("" + "hello") {
                    case "foo bar":
                        v3 += (((3.14159 - v3) - (v3 / 1e10)) * (((v3 + ((v3 + 3.14159) - (v3 / (1.5 * v3)))) / v3) - -2.25))
                    case "hello":
                        if len(v5) > 0 {
                            v5[0] = (v2 ^ (fi65 % (((v2 ^ (v2 % ((fi65) | 1)))) | 1)))
                        }
                    default:
                        for fi70 := 0; fi70 < 17; fi70++ {
                            var v71 int8 = 100
                            v16 += v16
                            _ = fi70
                            _ = v71
                        }
                    }
                    _ = fi65
                }
            default:
                v16 -= v16
            }
            for fi72 := 0; fi72 < 19; fi72++ {
                var v73 int32 = 42
                _ = fi72
                _ = v73
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v5
}
