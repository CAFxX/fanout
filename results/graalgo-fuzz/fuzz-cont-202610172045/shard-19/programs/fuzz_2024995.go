package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint; F1 uint; F2 int32 }
type GenS5 struct { F0 bool; F1 bool; F2 uint32 }
type GenS6 struct { F0 string; F1 string }

func main() {
    var v1 uint32 = 752
    v2 := []string{("hello" + "a"), "hello"}
    v3 := []uint32{(^(v1)), (v1 << ((1) & 63))}
    var v4 int16 = 32767
    v1 = v1
    v5 := GenS1{F0: 911}
    if (0 == 42) {
        for fi6 := 0; fi6 < 6; fi6++ {
            v7 := GenS1{F0: (v1 - (v1 >> ((255) & 63)))}
            _ = fi6
            _ = v7
        }
        var v8 uint = 781
        _ = v8
    } else {
        for fi9 := 0; fi9 < 6; fi9++ {
            switch ("" + "\n\t") {
            case "a":
                for fi10 := 0; fi10 < 5; fi10++ {
                    var v11 bool = ((0.1 != (1.5 - 1e10)) || true)
                    _ = fi10
                    _ = v11
                }
            default:
                var v12 uint = 237
                _ = v12
            }
            var v13 int64 = 679
            _ = fi9
            _ = v13
        }
    }
    if true {
        for fi14 := 0; fi14 < 18; fi14++ {
            for fi15 := 0; fi15 < 13; fi15++ {
                var v16 uint32 = v1
                if (1000 <= int16(fi14)) {
                    var v17 bool = ((-2.25 <= 1.5) && false)
                    v18 := GenS1{F0: v16}
                    _ = v17
                    _ = v18
                }
                if (!(true || false)) {
                    v19 := GenS1{F0: v16}
                    _ = v19
                }
                _ = fi15
                _ = v16
            }
            fmt.Println(v1)
            var v20 int = (fi14 << ((2) & 63))
            _ = fi14
            _ = v20
        }
        v21 := GenS2{F0: 1}
        _ = v21
    }
    switch "a" {
    case "a":
        v22 := GenS3{F0: ((!false) || false)}
        _ = v22
    default:
        switch 0 {
        case -2147483648:
            for fi23 := 0; fi23 < 18; fi23++ {
                v1 = (^(v1))
                v4 = 733
                _ = fi23
            }
        default:
            if len(v3) > 0 {
                v3[0] = 2
            }
        }
    }
    for fi24 := 0; fi24 < 11; fi24++ {
        switch "" {
        case "a":
            fmt.Println(fi24, fi24, v4)
        case "\n\t":
            v4 -= v4
        case "foo bar":
            if (-1 == 2147483647) {
                v25 := GenS2{F0: 829}
                v1 += (v1 << ((0) & 63))
                fmt.Println(fi24)
                _ = v25
            }
        default:
            switch fi24 {
            case 729:
                var v26 int8 = -1
                _ = v26
            case -209:
                switch (("" + "foo bar") + "hello") {
                case "hello":
                    var v27 int32 = 321
                    _ = v27
                case "":
                    var v28 uint16 = 1
                    _ = v28
                default:
                    if false {
                        var v29 float64 = 0.0
                        v1 += v1
                        _ = v29
                    }
                }
            case -2147483648:
                switch 729 {
                case 128:
                    v4 = 0
                }
            }
        }
        switch "" {
        case "a":
            v4 = (v4 >> ((65535) & 63))
        }
        _ = len(v3)
        _ = fi24
    }
    switch "foo bar" {
    case "a":
        fmt.Println(v1, v1)
    default:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    fmt.Println(v1, v4)
    for fi30 := 0; fi30 < 20; fi30++ {
        for fi31 := 0; fi31 < 12; fi31++ {
            switch 78 {
            case -637:
                v32 := GenS4{F0: 65535, F1: 2, F2: 42}
                _ = v32
            case -2147483648:
                switch fi31 {
                case 127:
                    switch -9223372036854775808 {
                    case 42:
                        switch fi31 {
                        case 17:
                            if (-16 < 1) {
                                if true {
                                    switch "foo bar" {
                                    case "foo bar":
                                        v33 := GenS5{F0: (true && ((("foo bar" + "\n\t") + "hello") == "")), F1: false, F2: 2147483647}
                                        _ = v33
                                    case "":
                                        if ((v4 | v4) == (v4 << ((50) & 63))) {
                                            fmt.Println(fi31, fi30)
                                            if len(v3) > 0 {
                                                v3[0] = 971
                                            }
                                        }
                                    default:
                                        fmt.Println(v1, v4)
                                    }
                                } else {
                                    fmt.Println(fi31, fi30, fi30)
                                }
                            } else {
                                v1 -= v1
                                switch 4294967295 {
                                case -552:
                                    v34 := GenS5{F0: false, F1: true, F2: v1}
                                    _ = v34
                                }
                            }
                        case -591:
                            fmt.Println(v4, fi31)
                        case -1:
                            if (((fi31 + fi31) <= 42) && (v4 >= v4)) {
                                v4 = (^(v4))
                                v4 = v4
                                v35 := GenS1{F0: v1}
                                _ = v35
                            }
                        default:
                            var v36 bool = (691 < 951)
                            _ = v36
                        }
                    }
                default:
                    if false {
                        _ = len(v2)
                        var v37 int64 = 9223372036854775807
                        _ = len(v3)
                        _ = v37
                    } else {
                        if (("\n\t" + "\n\t") != (("hello" + ("\n\t" + "hello")) + ("\n\t" + ("a" + "")))) {
                            fmt.Println(fi31, fi31)
                            fmt.Println(v1)
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v38 int = fi31
                            switch "" {
                            case "":
                                var v39 int8 = 100
                                _ = v39
                            }
                            _ = v38
                        }
                    }
                }
            case 42:
                v40 := GenS4{F0: 539, F1: 255, F2: -2147483648}
                _ = v40
            default:
                v41 := GenS6{F0: (((("" + "a") + ("hello" + "")) + "") + (("\n\t" + (("hello" + "hello") + (("hello" + "hello") + ("hello" + "")))) + "a")), F1: ""}
                _ = v41
            }
            v42 := GenS4{F0: 65535, F1: 514, F2: -2147483648}
            v43 := GenS5{F0: true, F1: false, F2: 65535}
            _ = fi31
            _ = v42
            _ = v43
        }
        v4 = (v4 | (v4 / (((v4 ^ (v4 & (v4 % (((v4 % (((^(v4))) | 1))) | 1))))) | 1)))
        switch fi30 {
        case -90:
            for fi44 := 0; fi44 < 2; fi44++ {
                _ = len(v3)
                _ = fi44
            }
        case 4294967295:
            fmt.Println(v4, v1, fi30)
        }
        _ = fi30
    }
    if false {
        v45 := GenS5{F0: false, F1: true, F2: (v1 >> ((597) & 63))}
        _ = v45
    } else {
        var v46 uint32 = (v1 & v1)
        for fi47 := 0; fi47 < 13; fi47++ {
            switch "foo bar" {
            case "a":
                v1 += v1
            case "foo bar":
                v4 += v4
            }
            var v48 string = "hello"
            _ = fi47
            _ = v48
        }
        var v49 int8 = -40
        _ = v46
        _ = v49
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v50 := GenS6{F0: "foo bar", F1: ("hello" + "foo bar")}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
}
