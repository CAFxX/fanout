package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 bool; F1 uint64; F2 string }
type GenS6 struct { F0 uint32; F1 bool; F2 int8 }

func main() {
    var v1 float64 = (3.14159 + 1.5)
    v2 := []string{""}
    v3 := []string{(("\n\t" + "foo bar") + "foo bar"), "foo bar", ("hello" + "hello")}
    var v4 bool = (float64(-2.25) == ((3.14159 * v1) / v1))
    var v5 int = -9223372036854775808
    var v6 float64 = 0.1
    if (0 >= 17) {
        if (2 <= 127) {
            for fi7 := 0; fi7 < 16; fi7++ {
                var v8 int8 = 0
                var v9 int16 = -370
                fmt.Println(v8, v1, fi7)
                _ = fi7
                _ = v8
                _ = v9
            }
            switch "foo bar" {
            case "":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            default:
                v10 := GenS1{F0: 853}
                _ = v10
            }
            switch v5 {
            case 227:
                v11 := GenS1{F0: 595}
                _ = v11
            case 1:
                for fi12 := 0; fi12 < 7; fi12++ {
                    switch (v5 ^ (fi12 | fi12)) {
                    case 831:
                        var v13 bool = v4
                        _ = v13
                    }
                    _ = fi12
                }
            case 4294967295:
                fmt.Println(v4)
            default:
                if true {
                    var v14 float64 = -2.25
                    _ = len(v3)
                    _ = v14
                } else {
                    v15 := GenS1{F0: 42}
                    _ = v15
                }
            }
        }
    } else {
        var v16 int8 = -128
        switch -869 {
        case 32767:
            v1 -= 0.0
        }
        _ = v16
    }
    for fi17 := 0; fi17 < 9; fi17++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi18 := 0; fi18 < 10; fi18++ {
            v19 := GenS1{F0: 255}
            var v20 bool = (!(4294967295 == 617))
            var v21 uint64 = 1
            if false {
                if (!(0 != -9223372036854775808)) {
                    v5 = (fi18 ^ fi18)
                } else {
                    var v22 int16 = 1000
                    if (498 < 1) {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    switch v5 {
                    case -372:
                        var v23 float64 = float64(0.1)
                        _ = v23
                    case -2147483648:
                        v24 := GenS2{F0: 255, F1: 0}
                        if v20 {
                            switch 127 {
                            case 9223372036854775807:
                                v3 = append(v3, "")
                            case 4294967295:
                                if false {
                                    v21 = 2
                                    v21 = v21
                                    _ = len(v2)
                                } else {
                                    var v25 int16 = v22
                                    _ = v25
                                }
                            case 1:
                                var v26 float32 = 0.1
                                _ = v26
                            }
                            v4 = (((v21 * (^(v21))) <= 2) || v20)
                        }
                        var v27 float32 = (1e10 / 1.5)
                        _ = v24
                        _ = v27
                    case 2147483647:
                        v28 := GenS1{F0: 2}
                        _ = v28
                    }
                    _ = v22
                }
                fmt.Println(fi17)
            } else {
                if (!((("foo bar" >= ("" + ("hello" + "foo bar"))) || true) || v4)) {
                    v4 = false
                    v1 += float64(-2.25)
                } else {
                    if false {
                        var v29 string = "foo bar"
                        v6 = v6
                        _ = v29
                    }
                    v30 := GenS1{F0: 65535}
                    if (127 != 232) {
                        if v4 {
                            v21 += 805
                        }
                    }
                    var v31 uint32 = 2147483647
                    var v32 int32 = 1
                    _ = v30
                    _ = v31
                    _ = v32
                }
            }
            _ = fi18
            _ = v19
            _ = v20
            _ = v21
        }
        for fi33 := 0; fi33 < 3; fi33++ {
            v4 = (!v4)
            if false {
                v6 -= 1.5
                v2 = append(v2, ("\n\t" + ""))
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            switch (("foo bar" + ("hello" + "\n\t")) + "hello") {
            case "\n\t":
                v6 -= 3.14159
            case "hello":
                v34 := GenS3{F0: 2147483647, F1: 18446744073709551615}
                _ = v34
            case "foo bar":
                switch (^(fi33)) {
                case 9223372036854775807:
                    v35 := GenS4{F0: 715}
                    _ = v35
                case 0:
                    if v4 {
                        if v4 {
                            v36 := GenS2{F0: 42, F1: 1}
                            _ = v36
                        }
                    } else {
                        v37 := GenS3{F0: 895, F1: 42}
                        v1 = float64(1.5)
                        if (-128 < 0) {
                            if v4 {
                                var v38 uint64 = 0
                                _ = v38
                            }
                            var v39 uint64 = 159
                            var v40 int16 = 32767
                            _ = v39
                            _ = v40
                        } else {
                            switch (fi17 + -9223372036854775808) {
                            case -1:
                                switch "a" {
                                case "hello":
                                    v4 = (2147483647 <= 629)
                                case "foo bar":
                                    var v41 uint8 = 128
                                    _ = v41
                                default:
                                    var v42 uint16 = 291
                                    _ = v42
                                }
                                var v43 uint32 = 688
                                _ = v43
                            default:
                                switch (fi17 >> ((18446744073709551615) & 63)) {
                                case -2147483648:
                                    v44 := GenS2{F0: 2, F1: 255}
                                    _ = v44
                                case -852:
                                    if len(v3) > 0 {
                                        v3[0] = "foo bar"
                                    }
                                }
                            }
                            fmt.Println(v6)
                            var v45 uint32 = 0
                            if (127 <= 42) {
                                var v46 int32 = 42
                                v1 = v6
                                _ = len(v3)
                                _ = v46
                            } else {
                                v1 += 1e10
                                fmt.Println(v6, v1)
                            }
                            _ = v45
                        }
                        _ = v37
                    }
                case 2147483647:
                    if len(v3) > 0 {
                        v3[0] = "foo bar"
                    }
                default:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            }
            _ = fi33
        }
        _ = fi17
    }
    fmt.Println(v5, v4, v1)
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    for fi47 := 0; fi47 < 7; fi47++ {
        for fi48 := 0; fi48 < 7; fi48++ {
            fmt.Println(v6)
            _ = fi48
        }
        v4 = v4
        _ = fi47
    }
    v6 += ((1e10 / ((-2.25 * 3.14159) * (0.0 / (v1 * v6)))) + v1)
    var v49 int32 = -2147483648
    for fi50 := 0; fi50 < 12; fi50++ {
        var v51 float64 = float64(float32(v6))
        v1 = float64(3.14159)
        _ = fi50
        _ = v51
    }
    var v52 float32 = -2.25
    for fi53 := 0; fi53 < 18; fi53++ {
        v4 = v4
        var v54 uint = 2
        _ = fi53
        _ = v54
    }
    v55 := GenS2{F0: 255, F1: 254}
    for fi56 := 0; fi56 < 16; fi56++ {
        v57 := GenS3{F0: fi56, F1: 2}
        switch v5 {
        case 42:
            if ((v4 || ((1 < 18446744073709551615) && v4)) && false) {
                v4 = (v1 >= 1e10)
                v58 := GenS4{F0: 2}
                for fi59 := 0; fi59 < 11; fi59++ {
                    v1 += ((0.0 + float64(v52)) + 1.5)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v60 := GenS5{F0: ((true || (1 >= (v49 & (-(v49))))) && ("hello" <= "\n\t")), F1: 18446744073709551615, F2: "a"}
                    v2 = append(v2, ("" + "a"))
                    var v61 bool = (((1 > 42) || (!v4)) && true)
                    _ = fi59
                    _ = v60
                    _ = v61
                }
                _ = v58
            }
        case 551:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        default:
            if ((v52 - v52) < (0.0 + ((v52 - (-2.25 / v52)) * float32(v1)))) {
                fmt.Println(v5)
                for fi62 := 0; fi62 < 17; fi62++ {
                    if (!v4) {
                        if false {
                            v2 = append(v2, "foo bar")
                            _ = len(v2)
                        } else {
                            v4 = v4
                            v63 := GenS5{F0: (!(!(!v4))), F1: 42, F2: ""}
                            _ = v63
                        }
                        var v64 uint32 = 275
                        var v65 string = "a"
                        v66 := GenS2{F0: 2, F1: 42}
                        v67 := GenS5{F0: (!v4), F1: 18446744073709551615, F2: v65}
                        _ = v64
                        _ = v65
                        _ = v66
                        _ = v67
                    }
                    _ = fi62
                }
            } else {
                v49 -= (v49 / ((v49) | 1))
                if len(v3) > 0 {
                    v3[0] = ""
                }
                if (!("\n\t" == (("foo bar" + "foo bar") + "hello"))) {
                    if len(v3) > 0 {
                        v3[0] = "foo bar"
                    }
                    var v68 int16 = 377
                    switch (("a" + "hello") + "hello") {
                    case "":
                        for fi69 := 0; fi69 < 4; fi69++ {
                            fmt.Println(v1)
                            v5 = -726
                            var v70 int16 = (v68 << ((563) & 63))
                            _ = fi69
                            _ = v70
                        }
                    }
                    _ = v68
                } else {
                    v71 := GenS6{F0: 2147483648, F1: (!((!true) || (42 == v49))), F2: 42}
                    _ = v71
                }
            }
        }
        _ = fi56
        _ = v57
    }
    for fi72 := 0; fi72 < 2; fi72++ {
        var v73 float64 = v1
        _ = fi72
        _ = v73
    }
    v74 := GenS6{F0: 2, F1: (1 == 236), F2: 100}
    for fi75 := 0; fi75 < 7; fi75++ {
        for fi76 := 0; fi76 < 7; fi76++ {
            var v77 int32 = v49
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi76
            _ = v77
        }
        _ = fi75
    }
    var v78 int64 = 9223372036854775807
    if v4 {
        if len(v2) > 0 {
            v2[0] = "foo bar"
        }
        switch ("\n\t" + ("hello" + ("" + "a"))) {
        case "":
            fmt.Println(v78)
        case "foo bar":
            if v4 {
                var v79 uint = 2
                var v80 int64 = -1
                var v81 int = -223
                _ = v79
                _ = v80
                _ = v81
            }
        case "hello":
            v78 = v78
        default:
            fmt.Println(v52, v1)
        }
        if len(v3) > 0 {
            v3[0] = "\n\t"
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v52
    _ = v55
    _ = v6
    _ = v74
    _ = v78
}
