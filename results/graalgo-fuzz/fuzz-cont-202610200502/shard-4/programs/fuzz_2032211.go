package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 uint8; F1 int16; F2 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64; F1 int8; F2 bool }
type GenS5 struct { F0 int64; F1 int16 }

func main() {
    var v1 uint64 = 42
    var v2 int32 = -2147483648
    var v3 int32 = 0
    v4 := make(map[string]bool)
    v5 := []string{"foo bar", "hello"}
    v5 = append(v5, "hello")
    var v6 int8 = 31
    fmt.Println(v3)
    if false {
        if true {
            switch ("foo bar" + ("" + ("\n\t" + ""))) {
            case "\n\t":
                for fi7 := 0; fi7 < 10; fi7++ {
                    fmt.Println(v6, v6, v3)
                    v6 = int8(v1)
                    _ = fi7
                }
            }
            switch ((("a" + ("foo bar" + "")) + "a") + "hello") {
            case "a":
                switch int(v6) {
                case 2147483647:
                    if false {
                        v2 -= -2147483648
                    } else {
                        switch ((("hello" + "hello") + "foo bar") + "") {
                        case "hello":
                            if false {
                                fmt.Println(v1, v3)
                                v2 = (-(v3))
                            }
                        case "":
                            var v8 string = "\n\t"
                            _ = v8
                        case "foo bar":
                            var v9 uint32 = 2147483647
                            _ = v9
                        default:
                            v2 = v3
                        }
                    }
                default:
                    v10 := GenS1{F0: 0, F1: 0}
                    _ = v10
                }
            case "":
                delete(v4, "")
            case "foo bar":
                v2 -= v3
            }
        } else {
            var v11 bool = (!false)
            _ = v11
        }
        if ((true || true) || ((!(-578 <= 810)) || (-763 > 1))) {
            if false {
                switch 2147483647 {
                case 127:
                    fmt.Println(v6, v2, v3)
                case 2147483647:
                    switch ("\n\t" + "") {
                    case "foo bar":
                        if true {
                            for fi12 := 0; fi12 < 16; fi12++ {
                                if false {
                                    var v13 float32 = float32(float64(1e10))
                                    v5 = append(v5, ("a" + ("" + "hello")))
                                    _ = v13
                                } else {
                                    for fi14 := 0; fi14 < 7; fi14++ {
                                        v4[""] = false
                                        var v15 int16 = -32768
                                        _ = fi14
                                        _ = v15
                                    }
                                }
                                v16 := GenS2{F0: 1, F1: 1, F2: 4294967295}
                                _ = fi12
                                _ = v16
                            }
                            for fi17 := 0; fi17 < 11; fi17++ {
                                for fi18 := 0; fi18 < 8; fi18++ {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    var v19 uint32 = 42
                                    _ = fi18
                                    _ = v19
                                }
                                _ = fi17
                            }
                        } else {
                            for fi20 := 0; fi20 < 4; fi20++ {
                                fmt.Println(fi20, v1)
                                _ = fi20
                            }
                            for fi21 := 0; fi21 < 10; fi21++ {
                                var v22 uint8 = 2
                                _ = fi21
                                _ = v22
                            }
                            if (!((v2 << ((860) & 63)) == v2)) {
                                if ("\n\t" > ("" + "\n\t")) {
                                    for fi23 := 0; fi23 < 9; fi23++ {
                                        v1 = 647
                                        _ = fi23
                                    }
                                    v24 := GenS3{F0: (v1 - (v1 ^ 434))}
                                    var v25 int16 = -32768
                                    _ = v24
                                    _ = v25
                                } else {
                                    v26 := GenS2{F0: 183, F1: 1000, F2: 32767}
                                    v27 := GenS1{F0: (-(v6)), F1: 128}
                                    _ = v26
                                    _ = v27
                                }
                                var v28 float32 = 0.1
                                if (false && (!false)) {
                                    fmt.Println(v28)
                                    var v29 uint64 = v1
                                    _ = v29
                                }
                                _ = v28
                            }
                        }
                    }
                case 0:
                    fmt.Println(v1, v2, v1)
                }
                var v30 int32 = -1
                _ = v30
            } else {
                v1 += v1
                var v31 float64 = -2.25
                fmt.Println(v2)
                _ = v31
            }
        }
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    if (true || false) {
        for fi32 := 0; fi32 < 4; fi32++ {
            v3 = v3
            _ = fi32
        }
        if (!true) {
            v2 += (v2 << ((18446744073709551615) & 63))
            if (!true) {
                var v33 float64 = (3.14159 / 0.1)
                fmt.Println(v1)
                switch "hello" {
                case "\n\t":
                    v6 -= int8(v2)
                default:
                    switch 4294967295 {
                    case 0:
                        _ = len(v4)
                    case 1:
                        v34 := GenS3{F0: v1}
                        _ = v34
                    case 127:
                        var v35 int = -527
                        _ = v35
                    default:
                        var v36 int64 = 0
                        _ = v36
                    }
                }
                _ = v33
            } else {
                v37 := GenS1{F0: 0, F1: 32767}
                _ = v37
            }
        }
        v2 -= v3
    } else {
        v1 += 1
        var v38 string = "a"
        _ = v38
    }
    if (false || (false || false)) {
        v6 -= v6
        var v39 float64 = -2.25
        _ = v39
    } else {
        var v40 uint32 = 2147483648
        if (false && false) {
            switch "foo bar" {
            case "foo bar":
                for fi41 := 0; fi41 < 7; fi41++ {
                    if ((2 == 51) && (-2147483648 > 42)) {
                        var v42 uint8 = 255
                        fmt.Println(v3, v42, v2)
                        switch (fi41 >> ((42) & 63)) {
                        case 1:
                            v3 += (-(v3))
                        case -75:
                            if (489 != 1) {
                                v2 = v2
                            } else {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                v5 = append(v5, "foo bar")
                            }
                        default:
                            v40 += v40
                        }
                        _ = v42
                    } else {
                        switch 730 {
                        case 32767:
                            var v43 string = ""
                            _ = v43
                        case 9223372036854775807:
                            v44 := GenS4{F0: 9223372036854775807, F1: v6, F2: true}
                            _ = v44
                        case 127:
                            v40 = (v40 - (v40 | (v40 | (^(v40)))))
                        default:
                            var v45 uint8 = 127
                            _ = v45
                        }
                    }
                    delete(v4, "a")
                    fmt.Println(v6, v40, v40)
                    _ = fi41
                }
            case "hello":
                for fi46 := 0; fi46 < 8; fi46++ {
                    _ = len(v4)
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    if len(v5) > 0 {
                        v5[0] = "foo bar"
                    }
                    _ = fi46
                }
            }
            for fi47 := 0; fi47 < 5; fi47++ {
                switch fi47 {
                case 9223372036854775807:
                    fmt.Println(v40, v2)
                default:
                    var v48 int = 42
                    _ = v48
                }
                _ = fi47
            }
        }
        _ = v40
    }
    v1 -= (v1 | v1)
    switch "a" {
    case "a":
        var v49 bool = false
        _ = v49
    default:
        var v50 uint32 = 2
        _ = v50
    }
    v51 := GenS1{F0: 84, F1: 471}
    for fi52 := 0; fi52 < 8; fi52++ {
        if ((0.1 - 3.14159) > 1.5) {
            var v53 uint64 = (v1 >> ((33) & 63))
            for fi54 := 0; fi54 < 16; fi54++ {
                if len(v5) > 0 {
                    v5[0] = "a"
                }
                _ = fi54
            }
            _ = v53
        } else {
            v55 := GenS4{F0: -9223372036854775808, F1: (v6 * v6), F2: (1 >= (v6 - v6))}
            for fi56 := 0; fi56 < 4; fi56++ {
                if ((((1.5 - 1e10) * (1.5 * ((1e10 + (1.5 * 0.0)) - ((1e10 / ((1.5 - ((3.14159 - 1e10) - -2.25)) + 0.1)) * (1.5 + 1.5))))) / 3.14159) >= 1.5) {
                    var v57 float64 = 0.0
                    _ = v57
                } else {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v1 -= (v1 << ((uint(v1)) & 63))
                }
                var v58 int = fi52
                _ = fi56
                _ = v58
            }
            _ = v55
        }
        v6 = v6
        var v59 uint64 = v1
        _ = fi52
        _ = v59
    }
    if ((!false) || (true || true)) {
        fmt.Println(v1)
        v1 += 0
    }
    fmt.Println(v3)
    v60 := GenS5{F0: -9223372036854775808, F1: 32767}
    switch 127 {
    case -265:
        fmt.Println(v1, v3)
    case -2147483648:
        if (!(true && false)) {
            v3 = -828
        }
    case 128:
        var v61 int16 = 327
        _ = v61
    }
    if false {
        v4[((("hello" + "") + ((("hello" + "foo bar") + "hello") + "")) + "a")] = false
    }
    if (2147483648 == 65535) {
        if (443 != 2147483648) {
            if true {
                if false {
                    for fi62 := 0; fi62 < 5; fi62++ {
                        v1 += v1
                        if len(v5) > 0 {
                            v5[0] = "\n\t"
                        }
                        v63 := GenS5{F0: -795, F1: 42}
                        _ = fi62
                        _ = v63
                    }
                }
                if (v6 == (v6 - int8(v1))) {
                    switch -9223372036854775808 {
                    case -12:
                        var v64 uint = 2
                        _ = v64
                    }
                    for fi65 := 0; fi65 < 6; fi65++ {
                        var v66 int16 = 0
                        _ = fi65
                        _ = v66
                    }
                } else {
                    for fi67 := 0; fi67 < 7; fi67++ {
                        v3 = 42
                        _ = fi67
                    }
                    switch 32768 {
                    case 4294967295:
                        var v68 int8 = (v6 % ((v6) | 1))
                        _ = v68
                    case 1:
                        delete(v4, "foo bar")
                    default:
                        v1 -= (v1 >> ((65535) & 63))
                    }
                }
                fmt.Println(v1)
            }
            var v69 int8 = (-(v6))
            _ = v69
        } else {
            v4["hello"] = false
            var v70 uint8 = 255
            _ = v70
        }
        var v71 float64 = -2.25
        _ = v71
    }
    if true {
        switch 2147483647 {
        case 0:
            if len(v5) > 0 {
                v5[0] = "\n\t"
            }
        case 300:
            v72 := GenS4{F0: 9223372036854775807, F1: int8(v3), F2: (true || ((v2 << ((2) & 63)) >= -895))}
            _ = v72
        }
        _ = len(v4)
        v2 += -1
    } else {
        fmt.Println(v3, v2, v6)
    }
    for fi73 := 0; fi73 < 18; fi73++ {
        var v74 float64 = -2.25
        var v75 int32 = 42
        v2 -= (v3 << ((18446744073709551615) & 63))
        _ = fi73
        _ = v74
        _ = v75
    }
    fmt.Println("v1:", v1)
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
    _ = v51
    _ = v6
    _ = v60
}
