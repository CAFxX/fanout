package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int64 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint }

func main() {
    v1 := []int{4294967295, 9223372036854775807, 42}
    var v2 uint32 = 1
    var v3 bool = ((69 <= -1) && true)
    var v4 uint = 1
    var v5 int16 = 0
    if v3 {
        v6 := GenS1{F0: true, F1: 2, F2: -1}
        for fi7 := 0; fi7 < 12; fi7++ {
            if len(v1) > 0 {
                v1[0] = -1
            }
            v2 = 640
            _ = fi7
        }
        _ = v6
    } else {
        var v8 uint16 = 65535
        if ((!v3) || (!(-128 >= 31))) {
            v9 := GenS2{F0: v5, F1: 0}
            _ = v9
        }
        _ = v8
    }
    fmt.Println(v2)
    v10 := GenS2{F0: (v5 - (v5 >> ((v4) & 63))), F1: 0}
    v1 = append(v1, -228)
    switch "foo bar" {
    case "foo bar":
        var v11 string = "\n\t"
        _ = v11
    case "":
        fmt.Println(v5)
    case "hello":
        var v12 float64 = 0.0
        _ = v12
    default:
        if (true && (true && v3)) {
            var v13 int64 = 0
            if (!(0 > 128)) {
                for fi14 := 0; fi14 < 18; fi14++ {
                    if len(v1) > 0 {
                        v1[0] = fi14
                    }
                    _ = fi14
                }
                v13 = v13
                v15 := GenS1{F0: v3, F1: 0, F2: 9223372036854775807}
                _ = v15
            } else {
                switch -224 {
                case 4294967295:
                    var v16 string = "\n\t"
                    _ = v16
                case 740:
                    if (56 == 1) {
                        for fi17 := 0; fi17 < 5; fi17++ {
                            fmt.Println(v2, v13, v5)
                            for fi18 := 0; fi18 < 17; fi18++ {
                                fmt.Println(fi18, v4)
                                switch "hello" {
                                case "":
                                    var v19 int16 = v5
                                    _ = v19
                                }
                                var v20 uint16 = 463
                                _ = fi18
                                _ = v20
                            }
                            fmt.Println(v3, v5)
                            _ = fi17
                        }
                        if false {
                            for fi21 := 0; fi21 < 6; fi21++ {
                                v22 := GenS1{F0: v3, F1: 255, F2: v13}
                                if (42 <= 18446744073709551615) {
                                    fmt.Println(v13, v13)
                                    v23 := GenS2{F0: -1, F1: 808}
                                    _ = v23
                                }
                                v1 = append(v1, fi21)
                                _ = fi21
                                _ = v22
                            }
                            for fi24 := 0; fi24 < 18; fi24++ {
                                fmt.Println(v3, v5)
                                var v25 float64 = 3.14159
                                fmt.Println(v3, v2)
                                _ = fi24
                                _ = v25
                            }
                            for fi26 := 0; fi26 < 5; fi26++ {
                                var v27 uint64 = 122
                                _ = fi26
                                _ = v27
                            }
                        }
                        if (!false) {
                            var v28 float64 = 3.14159
                            _ = v28
                        }
                    }
                default:
                    var v29 uint = (^(v4))
                    _ = v29
                }
                var v30 int32 = -2147483648
                if (-128 <= 74) {
                    v5 -= (v5 % (((v5 / ((int16(v30)) | 1))) | 1))
                    var v31 int16 = (^(v5))
                    _ = v31
                } else {
                    switch "a" {
                    case "a":
                        if v3 {
                            if (v3 && ((!true) && (0.1 >= 1e10))) {
                                v5 += v5
                                for fi32 := 0; fi32 < 15; fi32++ {
                                    v2 += v2
                                    for fi33 := 0; fi33 < 12; fi33++ {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = fi33
                                    }
                                    var v34 int64 = v13
                                    _ = fi32
                                    _ = v34
                                }
                            } else {
                                fmt.Println(v2)
                                v35 := GenS1{F0: v3, F1: uint16(v2), F2: (^(v13))}
                                _ = v35
                            }
                            v36 := GenS2{F0: v5, F1: 18446744073709551615}
                            _ = v36
                        }
                    case "":
                        v4 += v4
                    case "\n\t":
                        switch ("foo bar" + (("a" + ("hello" + "a")) + "\n\t")) {
                        case "\n\t":
                            for fi37 := 0; fi37 < 11; fi37++ {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                switch fi37 {
                                case 0:
                                    var v38 int8 = 127
                                    _ = v38
                                default:
                                    var v39 int32 = (v30 / (((v30 / (((v30 % ((v30) | 1))) | 1))) | 1))
                                    _ = v39
                                }
                                if false {
                                    v5 = (v5 + int16(v4))
                                    if v3 {
                                        switch "" {
                                        case "a":
                                            if len(v1) > 0 {
                                                v1[0] = (fi37 & -2147483648)
                                            }
                                        case "":
                                            _ = len(v1)
                                        case "\n\t":
                                            var v40 int = fi37
                                            _ = v40
                                        default:
                                            var v41 int8 = -36
                                            _ = v41
                                        }
                                    }
                                }
                                _ = fi37
                            }
                        case "foo bar":
                            for fi42 := 0; fi42 < 12; fi42++ {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = fi42
                            }
                        default:
                            v43 := GenS2{F0: v5, F1: 2}
                            _ = v43
                        }
                    }
                }
                _ = v30
            }
            if (v3 || v3) {
                fmt.Println(v13)
                var v44 int16 = v5
                v45 := GenS3{F0: 366}
                _ = v44
                _ = v45
            }
            _ = v13
        } else {
            var v46 int32 = 338
            v5 = int16(v2)
            _ = v46
        }
    }
    fmt.Println(v2)
    v47 := GenS1{F0: true, F1: 0, F2: 1}
    fmt.Println(v2, v2)
    var v48 uint = (v4 & v4)
    if len(v1) > 0 {
        v1[0] = -9223372036854775808
    }
    v49 := GenS4{F0: (^(v48))}
    v50 := GenS3{F0: -9223372036854775808}
    fmt.Println(v2)
    switch ("\n\t" + ("\n\t" + ("a" + "hello"))) {
    case "foo bar":
        for fi51 := 0; fi51 < 10; fi51++ {
            var v52 uint = v4
            v53 := GenS2{F0: v5, F1: 4294967295}
            var v54 int32 = 0
            _ = fi51
            _ = v52
            _ = v53
            _ = v54
        }
    case "hello":
        v55 := GenS1{F0: (-9223372036854775808 != -9223372036854775808), F1: 255, F2: 0}
        _ = v55
    case "":
        if v3 {
            switch -9223372036854775808 {
            case 128:
                v4 = v4
            case 9223372036854775807:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            v56 := GenS1{F0: (!v3), F1: 323, F2: 1}
            for fi57 := 0; fi57 < 12; fi57++ {
                var v58 int8 = 0
                for fi59 := 0; fi59 < 15; fi59++ {
                    var v60 int16 = 1
                    var v61 string = "foo bar"
                    _ = fi59
                    _ = v60
                    _ = v61
                }
                v48 += 42
                _ = fi57
                _ = v58
            }
            _ = v56
        }
    }
    if len(v1) > 0 {
        v1[0] = -975
    }
    var v62 int32 = 1
    v62 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v62
}
