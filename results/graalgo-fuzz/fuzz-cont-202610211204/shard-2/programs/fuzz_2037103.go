package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 int; F2 uint8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint16; F1 int32; F2 int }
type GenS6 struct { F0 uint; F1 int32 }
type GenS7 struct { F0 int16; F1 uint64 }
type GenS8 struct { F0 uint64 }
func genFn0() uint16 {
    var v6 uint8 = 159
    var v7 int64 = -203
    var v8 int16 = -826
    _ = v6
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 uint16) int32 {
    var v9 uint32 = 2147483648
    switch 32767 {
    case -1:
        switch 128 {
        case -597:
            var v10 uint8 = 128
            _ = v10
        }
    case 2147483647:
        v11 := []int{9223372036854775807, 4294967295, 800}
        _ = v11
    case 32768:
        var v12 uint16 = uint16(v9)
        _ = v12
    default:
        v13 := GenS1{F0: v9}
        _ = v13
    }
    _ = p0
    _ = v9
    return -183
}

func main() {
    var v1 int64 = 991
    var v2 uint8 = 226
    var v3 uint8 = v2
    var v4 int64 = (v1 << ((0) & 63))
    v5 := []string{"a", (("foo bar" + (("hello" + (("\n\t" + "foo bar") + ("foo bar" + (("hello" + "a") + "hello")))) + "a")) + ("a" + ("foo bar" + "hello"))), "a"}
    v3 += v2
    if true {
        if false {
            fmt.Println(v3, v1)
            for fi14 := 0; fi14 < 5; fi14++ {
                switch "a" {
                case "hello":
                    if (259 <= 749) {
                        var v15 string = ("hello" + "")
                        _ = v15
                    } else {
                        if len(v5) > 0 {
                            v5[0] = "\n\t"
                        }
                    }
                case "\n\t":
                    v16 := GenS1{F0: 2}
                    _ = v16
                case "a":
                    var v17 uint32 = 1
                    _ = v17
                }
                v2 += v2
                _ = fi14
            }
        } else {
            switch (("\n\t" + "hello") + "a") {
            case "hello":
                if false {
                    for fi18 := 0; fi18 < 2; fi18++ {
                        switch (fi18 << ((18446744073709551615) & 63)) {
                        case 0:
                            v5 = append(v5, "a")
                        case 9223372036854775807:
                            var v19 bool = (0 == 65535)
                            _ = v19
                        case 42:
                            if len(v5) > 0 {
                                v5[0] = "\n\t"
                            }
                        }
                        _ = fi18
                    }
                    var v20 string = ""
                    fmt.Println(v1, v2)
                    _ = v20
                } else {
                    switch ("hello" + ("\n\t" + "\n\t")) {
                    case "hello":
                        var v21 uint32 = 943
                        _ = v21
                    default:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                    for fi22 := 0; fi22 < 8; fi22++ {
                        var v23 int16 = -218
                        _ = fi22
                        _ = v23
                    }
                }
            case "":
                for fi24 := 0; fi24 < 4; fi24++ {
                    if true {
                        v1 -= -378
                        var v25 int16 = 1000
                        var v26 bool = false
                        _ = v25
                        _ = v26
                    }
                    var v27 int = 32768
                    _ = fi24
                    _ = v27
                }
            case "foo bar":
                for fi28 := 0; fi28 < 6; fi28++ {
                    if (4294967295 == 65535) {
                        var v29 uint32 = 42
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        switch (fi28 / (((fi28 >> ((0) & 63))) | 1)) {
                        case -2147483648:
                            var v30 uint32 = 409
                            _ = v30
                        default:
                            if false {
                                if (32767 == 1) {
                                    for fi31 := 0; fi31 < 16; fi31++ {
                                        if false {
                                            v32 := GenS2{F0: 1}
                                            _ = v32
                                        }
                                        _ = len(v5)
                                        fmt.Println(fi28)
                                        _ = fi31
                                    }
                                } else {
                                    v33 := GenS3{F0: fi28, F1: (-(fi28)), F2: (^(v3))}
                                    if (!((float32(0.0) - 0.0) > (1.5 * (3.14159 + ((3.14159 / 3.14159) * 1.5))))) {
                                        switch (fi28 >> ((uint(v29)) & 63)) {
                                        case 127:
                                            var v34 uint32 = (v29 % ((v29) | 1))
                                            _ = v34
                                        case -507:
                                            for fi35 := 0; fi35 < 3; fi35++ {
                                                v4 = 9223372036854775807
                                                fmt.Println(v29)
                                                switch "" {
                                                case "foo bar":
                                                    switch (fi35 ^ int(v29)) {
                                                    case 127:
                                                        v4 += (v1 & 325)
                                                    case 2147483647:
                                                        if (fi35 != fi28) {
                                                            fmt.Println(v4)
                                                            var v36 int16 = 1
                                                            _ = v36
                                                        }
                                                    }
                                                case "hello":
                                                    v1 -= v4
                                                }
                                                _ = fi35
                                            }
                                        case -1:
                                            for fi37 := 0; fi37 < 17; fi37++ {
                                                v2 = (v3 ^ 127)
                                                var v38 int16 = -32768
                                                if true {
                                                    v5 = append(v5, "\n\t")
                                                    fmt.Println(v3, fi28)
                                                }
                                                _ = fi37
                                                _ = v38
                                            }
                                        default:
                                            if false {
                                                for fi39 := 0; fi39 < 3; fi39++ {
                                                    fmt.Println(v4, v4, v29)
                                                    _ = fi39
                                                }
                                            } else {
                                                for fi40 := 0; fi40 < 7; fi40++ {
                                                    if len(v5) > 0 {
                                                        v5[0] = "foo bar"
                                                    }
                                                    if true {
                                                        v2 -= (v2 % ((127) | 1))
                                                    }
                                                    var v41 float32 = -2.25
                                                    fmt.Println(v41)
                                                    _ = fi40
                                                    _ = v41
                                                }
                                            }
                                        }
                                    }
                                    fmt.Println(v4, fi28, v4)
                                    _ = v33
                                }
                                var v42 int = fi28
                                var v43 int32 = 0
                                _ = v42
                                _ = v43
                            } else {
                                _ = len(v5)
                                for fi44 := 0; fi44 < 10; fi44++ {
                                    fmt.Println(v29)
                                    _ = len(v5)
                                    fmt.Println(v2, v1, fi44)
                                    switch "" {
                                    case "foo bar":
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    case "\n\t":
                                        fmt.Println(fi44, v29, v1)
                                    case "hello":
                                        if len(v5) > 0 {
                                            v5[0] = (("hello" + ("hello" + "\n\t")) + ("foo bar" + ("foo bar" + "hello")))
                                        }
                                    default:
                                        v5 = append(v5, ("hello" + "\n\t"))
                                    }
                                    v45 := GenS1{F0: 686}
                                    _ = fi44
                                    _ = v45
                                }
                            }
                        }
                        _ = v29
                    }
                    _ = fi28
                }
            }
        }
    } else {
        v2 -= v3
        v2 += (^(v3))
        if false {
            for fi46 := 0; fi46 < 7; fi46++ {
                if len(v5) > 0 {
                    v5[0] = "\n\t"
                }
                if ((!(0.1 <= 0.1)) && (4294967295 < 0)) {
                    var v47 uint = 255
                    _ = v47
                } else {
                    v4 -= v4
                    switch "\n\t" {
                    case "foo bar":
                        v48 := GenS3{F0: fi46, F1: (fi46 + (fi46 << ((0) & 63))), F2: v3}
                        _ = v48
                    case "\n\t":
                        if true {
                            var v49 float64 = 1e10
                            _ = v49
                        } else {
                            v50 := GenS1{F0: 65535}
                            _ = v50
                        }
                    }
                }
                _ = fi46
            }
            if ("\n\t" >= (((("hello" + "a") + ("a" + ("hello" + "hello"))) + "a") + ("a" + ""))) {
                if (!(false || false)) {
                    switch "a" {
                    case "hello":
                        v51 := GenS4{F0: 4294967295}
                        _ = v51
                    }
                } else {
                    fmt.Println(v2, v1, v3)
                    if true {
                        v52 := GenS2{F0: 92}
                        v53 := GenS5{F0: 32767, F1: 42, F2: 4294967295}
                        _ = v52
                        _ = v53
                    }
                    v3 += v2
                }
                if false {
                    var v54 uint16 = 1
                    switch 128 {
                    case -288:
                        var v55 int16 = 71
                        _ = v55
                    case 0:
                        fmt.Println(v4, v54)
                    case 9223372036854775807:
                        switch 738 {
                        case 9223372036854775807:
                            var v56 uint = 18446744073709551615
                            _ = v56
                        case -78:
                            var v57 int16 = int16(v54)
                            _ = v57
                        }
                    }
                    v54 -= 0
                    _ = v54
                }
                if (!(!(!false))) {
                    v4 += 0
                    if (true || true) {
                        switch "hello" {
                        case "a":
                            var v58 uint32 = 2
                            _ = v58
                        }
                        for fi59 := 0; fi59 < 6; fi59++ {
                            if false {
                                v60 := GenS5{F0: 1, F1: -1, F2: 128}
                                v1 = v4
                                var v61 float32 = 1.5
                                _ = v60
                                _ = v61
                            }
                            var v62 int16 = -312
                            _ = fi59
                            _ = v62
                        }
                        if len(v5) > 0 {
                            v5[0] = "a"
                        }
                    } else {
                        var v63 int16 = 1000
                        _ = v63
                    }
                } else {
                    if true {
                        switch -641 {
                        case -9223372036854775808:
                            var v64 string = ("" + "a")
                            _ = v64
                        case 552:
                            _ = len(v5)
                        default:
                            switch 32768 {
                            case -9223372036854775808:
                                v3 = v2
                            }
                        }
                        var v65 int64 = (v1 * v4)
                        if true {
                            v66 := GenS2{F0: 100}
                            _ = v66
                        }
                        _ = v65
                    } else {
                        fmt.Println(v4)
                        v67 := GenS4{F0: 858}
                        _ = v67
                    }
                    v2 = 0
                    fmt.Println(v4)
                }
            }
            v68 := GenS3{F0: 627, F1: 395, F2: (^(v3))}
            _ = v68
        }
    }
    fmt.Println(v2)
    v1 = (v1 & v1)
    _ = len(v5)
    v1 += v4
    if true {
        fmt.Println(v4)
        fmt.Println(v1)
        v3 = v2
    }
    var v69 uint32 = 703
    v70 := GenS6{F0: 4294967295, F1: 2147483647}
    v71 := GenS7{F0: 36, F1: 42}
    v72 := GenS8{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v69
    _ = v70
    _ = v71
    _ = v72
}
