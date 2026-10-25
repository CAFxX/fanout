package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 bool }
type GenS5 struct { F0 uint; F1 int8; F2 string }

func main() {
    v1 := []string{"hello", ("\n\t" + "\n\t"), ("" + "foo bar"), ("" + ("" + ""))}
    var v2 string = ((("hello" + "hello") + "a") + "")
    var v3 int = 544
    var v4 float32 = (-2.25 + 3.14159)
    var v5 int8 = 33
    var v6 string = v2
    _ = len(v1)
    for fi7 := 0; fi7 < 3; fi7++ {
        v6 = v6
        v5 = (v5 + 1)
        _ = fi7
    }
    var v8 float64 = 0.1
    v8 = v8
    for fi9 := 0; fi9 < 13; fi9++ {
        v10 := GenS1{F0: 42, F1: 943}
        if true {
            switch "a" {
            case "\n\t":
                v11 := GenS1{F0: 9223372036854775807, F1: -783}
                _ = v11
            case "foo bar":
                v12 := GenS2{F0: 440}
                _ = v12
            case "":
                v4 = v4
            default:
                v13 := GenS2{F0: 0}
                _ = v13
            }
            switch -2147483648 {
            case -861:
                _ = len(v1)
            case -1:
                switch v6 {
                case "a":
                    for fi14 := 0; fi14 < 19; fi14++ {
                        if true {
                            v5 -= 42
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            switch (v3 % ((fi14) | 1)) {
                            case 1:
                                if true {
                                    if false {
                                        switch (fi9 + fi9) {
                                        case 271:
                                            v15 := GenS2{F0: -9223372036854775808}
                                            _ = v15
                                        default:
                                            if (float64(v4) > 1e10) {
                                                if (609 == 255) {
                                                    fmt.Println(fi9, v4, v5)
                                                } else {
                                                    var v16 float32 = v4
                                                    _ = v16
                                                }
                                            }
                                        }
                                        v2 = ("\n\t" + (v6 + ""))
                                    }
                                    v17 := GenS2{F0: -726}
                                    if ((true || false) || true) {
                                        var v18 float32 = 0.1
                                        v3 = v3
                                        _ = v18
                                    }
                                    _ = v17
                                } else {
                                    if ((v4 * v4) == (1e10 + 0.0)) {
                                        var v19 int64 = -1
                                        v20 := GenS3{F0: (v5 - 127)}
                                        _ = v19
                                        _ = v20
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    fmt.Println(v3)
                                }
                            case 4294967295:
                                fmt.Println(v4, fi9, v2)
                            case 32767:
                                fmt.Println(fi14, fi14, v4)
                            default:
                                v2 = "foo bar"
                            }
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = fi14
                    }
                }
            case 128:
                _ = len(v1)
            default:
                var v21 int = v3
                _ = v21
            }
            v4 += v4
        } else {
            for fi22 := 0; fi22 < 2; fi22++ {
                if false {
                    var v23 uint32 = 2147483647
                    switch (v6 + ("" + "foo bar")) {
                    case "":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    switch ("" + "a") {
                    case "foo bar":
                        var v24 uint32 = v23
                        _ = v24
                    case "":
                        var v25 int16 = -1
                        _ = v25
                    case "hello":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    v26 := GenS1{F0: 977, F1: 862}
                    var v27 float64 = float64(1e10)
                    _ = v23
                    _ = v26
                    _ = v27
                }
                v28 := GenS3{F0: v5}
                if (166 >= 21) {
                    v5 = (-(v5))
                }
                _ = fi22
                _ = v28
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        _ = fi9
        _ = v10
    }
    var v29 int32 = -2147483648
    var v30 int = 128
    var v31 int16 = -32768
    switch (("\n\t" + (v6 + v2)) + "\n\t") {
    case "foo bar":
        for fi32 := 0; fi32 < 3; fi32++ {
            if len(v1) > 0 {
                v1[0] = ((v2 + "") + ("" + v2))
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi32
        }
    default:
        v1 = append(v1, "\n\t")
    }
    switch (v6 + "hello") {
    case "\n\t":
        var v33 uint8 = 127
        _ = v33
    case "foo bar":
        if ("a" >= v6) {
            var v34 bool = (394 <= 65535)
            _ = v34
        }
    case "hello":
        if false {
            v4 = v4
            switch (v6 + v2) {
            case "hello":
                var v35 uint = 255
                _ = v35
            case "":
                var v36 int8 = (v5 * (v5 << ((0) & 63)))
                _ = v36
            case "foo bar":
                for fi37 := 0; fi37 < 8; fi37++ {
                    switch "foo bar" {
                    case "\n\t":
                        switch (fi37 - 0) {
                        case 380:
                            v38 := GenS1{F0: 9223372036854775807, F1: (v29 - -2147483648)}
                            _ = v38
                        }
                    case "foo bar":
                        fmt.Println(v2, v6, v31)
                    }
                    var v39 int = 905
                    var v40 int8 = (v5 % ((v5) | 1))
                    _ = fi37
                    _ = v39
                    _ = v40
                }
            }
        }
    }
    var v41 uint = 65535
    _ = len(v1)
    switch -993 {
    case -202:
        _ = len(v1)
    case 901:
        for fi42 := 0; fi42 < 14; fi42++ {
            var v43 string = "a"
            _ = fi42
            _ = v43
        }
    case -44:
        switch v6 {
        case "":
            fmt.Println(v2, v6, v3)
        case "a":
            v44 := GenS2{F0: 432}
            _ = v44
        default:
            switch ((("foo bar" + (v2 + ((v2 + ("foo bar" + v2)) + "a"))) + ((v6 + (v2 + v6)) + "a")) + v6) {
            case "hello":
                v41 = v41
            }
        }
    default:
        switch v30 {
        case 725:
            var v45 uint16 = 2
            _ = v45
        case -540:
            var v46 uint = 842
            _ = v46
        default:
            for fi47 := 0; fi47 < 12; fi47++ {
                if (1 != 65535) {
                    if true {
                        var v48 bool = false
                        for fi49 := 0; fi49 < 18; fi49++ {
                            v50 := GenS1{F0: 42, F1: -868}
                            switch v3 {
                            case 255:
                                fmt.Println(v6, v6)
                            default:
                                if true {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    var v51 uint32 = 4294967295
                                    fmt.Println(v51, v5, fi49)
                                    _ = v51
                                } else {
                                    var v52 uint = 0
                                    var v53 float64 = (1.5 - float64((v4 + v4)))
                                    _ = len(v1)
                                    _ = v52
                                    _ = v53
                                }
                            }
                            v8 += v8
                            _ = fi49
                            _ = v50
                        }
                        var v54 uint16 = 739
                        _ = v48
                        _ = v54
                    } else {
                        for fi55 := 0; fi55 < 9; fi55++ {
                            var v56 int = fi47
                            _ = fi55
                            _ = v56
                        }
                        for fi57 := 0; fi57 < 20; fi57++ {
                            fmt.Println(fi47, v6, v41)
                            _ = fi57
                        }
                    }
                    if true {
                        v30 = 9223372036854775807
                        var v58 uint8 = 1
                        fmt.Println(v8)
                        _ = v58
                    }
                } else {
                    if (-9223372036854775808 <= 0) {
                        switch v2 {
                        case "hello":
                            var v59 int8 = 127
                            _ = v59
                        case "\n\t":
                            v60 := GenS4{F0: ((v5 == v5) && (18446744073709551615 > 42)), F1: (-1 < v29)}
                            _ = v60
                        }
                        fmt.Println(v31, v6, v29)
                        v8 += (v8 + 1.5)
                    }
                }
                v61 := GenS5{F0: (v41 | (v41 << ((v41) & 63))), F1: v5, F2: ""}
                _ = fi47
                _ = v61
            }
        }
    }
    if (float64(((v4 / (float32(3.14159) * v4)) + v4)) < v8) {
        if len(v1) > 0 {
            v1[0] = v2
        }
    } else {
        fmt.Println(v3, v29, v30)
        var v62 int8 = int8(v29)
        _ = v62
    }
    var v63 uint32 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v41
    _ = v5
    _ = v6
    _ = v63
    _ = v8
}
