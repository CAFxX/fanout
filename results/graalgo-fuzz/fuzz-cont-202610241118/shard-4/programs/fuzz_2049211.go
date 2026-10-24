package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16; F1 bool; F2 uint64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 string }
func genFn0(p0 int64, p1 int8) int8 {
    v4 := GenS1{F0: p1, F1: 1}
    if (32767 != 560) {
        var v5 int = 32768
        _ = v5
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = v4
    return -116
}
func genFn1(p0 uint16, p1 uint32) int32 {
    fmt.Println(p0, p0)
    for fi6 := 0; fi6 < 9; fi6++ {
        switch "foo bar" {
        case "hello":
            fmt.Println(p0, p1)
        case "foo bar":
            var v7 int = (^(fi6))
            _ = v7
        default:
            for fi8 := 0; fi8 < 3; fi8++ {
                var v9 int16 = 1000
                v10 := GenS2{F0: fi8, F1: 938}
                switch (^(fi8)) {
                case 128:
                    if true {
                        switch -2147483648 {
                        case 639:
                            fmt.Println(fi6, fi8, p0)
                        case 2147483647:
                            switch fi6 {
                            case 127:
                                v11 := []int{}
                                _ = v11
                            case 9223372036854775807:
                                v9 = (v9 - (v9 | 72))
                            case -9223372036854775808:
                                var v12 int32 = -718
                                _ = v12
                            }
                        }
                    }
                }
                _ = fi8
                _ = v10
                _ = v9
            }
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    return -716
}

func main() {
    var v1 int8 = 100
    v2 := []string{"", (((("foo bar" + "a") + "hello") + "foo bar") + "foo bar")}
    var v3 uint = 4294967295
    if (-9223372036854775808 < 0) {
        var v13 bool = (true || true)
        v3 -= 542
        _ = v13
    } else {
        v2 = append(v2, "\n\t")
        for fi14 := 0; fi14 < 9; fi14++ {
            if ((-2147483648 > -99) && (false || false)) {
                switch "\n\t" {
                case "":
                    if true {
                        var v15 int64 = -1
                        v1 = 100
                        for fi16 := 0; fi16 < 7; fi16++ {
                            var v17 int64 = (v15 / ((-9223372036854775808) | 1))
                            _ = fi16
                            _ = v17
                        }
                        _ = v15
                    }
                case "hello":
                    fmt.Println(v3, fi14)
                default:
                    if (-32768 >= 1) {
                        if (!(uint16(v3) == 42)) {
                            if false {
                                var v18 float32 = 3.14159
                                switch "a" {
                                case "hello":
                                    switch (fi14 >> ((1) & 63)) {
                                    case 128:
                                        switch "a" {
                                        case "":
                                            v3 = (^(v3))
                                        case "hello":
                                            var v19 uint16 = 1
                                            _ = v19
                                        case "a":
                                            v20 := GenS1{F0: (v1 - v1), F1: 42}
                                            _ = v20
                                        default:
                                            v21 := GenS3{F0: 240}
                                            _ = v21
                                        }
                                    default:
                                        var v22 int8 = v1
                                        _ = v22
                                    }
                                case "\n\t":
                                    var v23 uint64 = 575
                                    _ = v23
                                case "foo bar":
                                    v24 := GenS2{F0: (fi14 >> (((v3 - 65535)) & 63)), F1: 981}
                                    _ = v24
                                default:
                                    var v25 uint32 = 2147483647
                                    _ = v25
                                }
                                for fi26 := 0; fi26 < 17; fi26++ {
                                    switch ("a" + ("a" + "\n\t")) {
                                    case "foo bar":
                                        switch "\n\t" {
                                        case "foo bar":
                                            v18 -= v18
                                        case "hello":
                                            _ = len(v2)
                                        default:
                                            var v27 int32 = 0
                                            _ = v27
                                        }
                                    default:
                                        v1 = v1
                                    }
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    fmt.Println(fi26)
                                    fmt.Println(fi26, v1)
                                    _ = fi26
                                }
                                _ = v18
                            }
                        }
                        v1 += 0
                    }
                }
                if (!true) {
                    var v28 float64 = (3.14159 / -2.25)
                    v29 := GenS2{F0: 128, F1: -1}
                    for fi30 := 0; fi30 < 15; fi30++ {
                        switch (^(fi30)) {
                        case 32767:
                            var v31 int64 = -931
                            _ = v31
                        case 127:
                            if (true && ((v3 * (v3 | (^(v3)))) < v3)) {
                                if true {
                                    v32 := GenS4{F0: -209, F1: false, F2: 42}
                                    v28 -= ((v28 - (v28 * 1.5)) + ((1.5 - v28) * ((v28 * float64(3.14159)) + (((0.1 - v28) / 1.0) * ((-2.25 / v28) / v28)))))
                                    v2 = append(v2, "hello")
                                    _ = v32
                                }
                                switch (fi14 << (((v3 >> (((v3 ^ (v3 / ((255) | 1)))) & 63))) & 63)) {
                                case 128:
                                    if true {
                                        var v33 uint = (^(v3))
                                        v2 = append(v2, "foo bar")
                                        _ = v33
                                    } else {
                                        fmt.Println(v1, fi14, fi30)
                                        var v34 string = ("a" + ("" + (("a" + ("foo bar" + ("a" + "a"))) + "hello")))
                                        v1 += (v1 & (v1 ^ v1))
                                        _ = v34
                                    }
                                case 444:
                                    if ("\n\t" >= "a") {
                                        v35 := GenS4{F0: 42, F1: true, F2: 134}
                                        var v36 uint = (v3 | v3)
                                        _ = v35
                                        _ = v36
                                    } else {
                                        v1 = v1
                                        if false {
                                            v37 := GenS2{F0: (^(fi30)), F1: 42}
                                            switch "foo bar" {
                                            case "\n\t":
                                                var v38 int = fi30
                                                _ = v38
                                            case "foo bar":
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            case "":
                                                v3 = v3
                                            }
                                            _ = v37
                                        }
                                    }
                                    var v39 float32 = 1.5
                                    _ = v39
                                case 243:
                                    var v40 uint64 = 2
                                    _ = v40
                                }
                                var v41 int16 = 1000
                                _ = v41
                            }
                        case 0:
                            var v42 uint8 = 128
                            _ = v42
                        default:
                            if false {
                                var v43 int32 = 2147483647
                                _ = v43
                            }
                        }
                        fmt.Println(v1, fi30, v28)
                        var v44 bool = ("a" > "foo bar")
                        _ = fi30
                        _ = v44
                    }
                    _ = v28
                    _ = v29
                } else {
                    var v45 bool = (42 != 309)
                    fmt.Println(v3)
                    var v46 uint8 = 255
                    _ = v45
                    _ = v46
                }
            } else {
                v47 := GenS5{F0: 736}
                v1 = (-(v1))
                _ = v47
            }
            _ = fi14
        }
    }
    v1 = (v1 + v1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v48 uint16 = 2
    switch (("\n\t" + "a") + "hello") {
    case "\n\t":
        var v49 int32 = -606
        _ = v49
    }
    if (42 <= (v48 << ((v3) & 63))) {
        v3 = v3
    }
    for fi50 := 0; fi50 < 9; fi50++ {
        if (true || (true || false)) {
            var v51 int16 = -1
            _ = v51
        }
        for fi52 := 0; fi52 < 7; fi52++ {
            var v53 int8 = -14
            var v54 uint8 = 72
            if true {
                v1 = v1
                var v55 int = fi50
                if false {
                    switch "a" {
                    case "foo bar":
                        v2 = append(v2, ("" + "foo bar"))
                    }
                    var v56 uint32 = 460
                    v57 := GenS1{F0: int8(v55), F1: 42}
                    _ = v56
                    _ = v57
                } else {
                    switch (fi52 >> ((v3) & 63)) {
                    case 42:
                        if len(v2) > 0 {
                            v2[0] = ""
                        }
                    case 32767:
                        var v58 uint = (v3 | v3)
                        _ = v58
                    case 488:
                        v55 -= (v55 * v55)
                    default:
                        var v59 bool = true
                        _ = v59
                    }
                    v2 = append(v2, ("foo bar" + ("" + "\n\t")))
                }
                _ = v55
            }
            _ = fi52
            _ = v53
            _ = v54
        }
        if len(v2) > 0 {
            v2[0] = ("" + ("" + ("\n\t" + "a")))
        }
        _ = fi50
    }
    v60 := GenS6{F0: ("hello" + ("hello" + "hello"))}
    for fi61 := 0; fi61 < 16; fi61++ {
        if ((((3.14159 / float64(0.1)) - (3.14159 * 1.5)) != (0.1 + 3.14159)) || false) {
            v2 = append(v2, "a")
        } else {
            fmt.Println(v3, fi61, v3)
            fmt.Println(fi61)
        }
        _ = fi61
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v48
    _ = v60
}
