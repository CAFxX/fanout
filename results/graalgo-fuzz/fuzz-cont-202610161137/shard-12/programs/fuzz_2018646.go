package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 uint64; F2 uint64 }
type GenS4 struct { F0 uint8; F1 int64; F2 uint32 }
type GenS5 struct { F0 int; F1 bool; F2 uint64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int32; F1 string; F2 uint8 }
type GenS8 struct { F0 int }

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{-9223372036854775808, 9223372036854775807, -9223372036854775808, -261}
    var v3 float64 = 1.5
    v4 := GenS1{F0: 803, F1: 255, F2: 413}
    switch "a" {
    case "\n\t":
        v3 += (-2.25 / v3)
    case "foo bar":
        var v5 int = -1
        _ = v5
    case "hello":
        v3 = v3
    default:
        for fi6 := 0; fi6 < 8; fi6++ {
            v3 -= v3
            v3 = (v3 * v3)
            var v7 int8 = 0
            _ = fi6
            _ = v7
        }
    }
    v3 -= float64(1e10)
    v3 -= v3
    v8 := GenS1{F0: 94, F1: 127, F2: 32767}
    if len(v2) > 0 {
        v2[0] = 809
    }
    v9 := GenS2{F0: 1}
    if true {
        fmt.Println(v3)
    } else {
        if false {
            for fi10 := 0; fi10 < 1; fi10++ {
                v3 -= 0.0
                _ = fi10
            }
        }
        switch -489 {
        case 2147483647:
            for fi11 := 0; fi11 < 11; fi11++ {
                v12 := GenS1{F0: 2, F1: 128, F2: -539}
                _ = fi11
                _ = v12
            }
        }
    }
    fmt.Println(v3)
    if (2147483647 >= -326) {
        switch 4294967295 {
        case 9223372036854775807:
            v3 = ((v3 / 1.0) / float64(3.14159))
        }
        v3 -= v3
        if (9223372036854775807 != -1) {
            for fi13 := 0; fi13 < 3; fi13++ {
                fmt.Println(fi13, fi13)
                _ = fi13
            }
            if true {
                _ = len(v2)
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    } else {
        fmt.Println(v3)
        for fi14 := 0; fi14 < 19; fi14++ {
            switch "\n\t" {
            case "foo bar":
                var v15 int8 = 76
                _ = v15
            case "\n\t":
                var v16 int16 = -1
                _ = v16
            case "hello":
                if ((!(false && false)) && (-2.25 < (0.1 / (0.0 - (3.14159 / 1.5))))) {
                    for fi17 := 0; fi17 < 1; fi17++ {
                        v3 += v3
                        var v18 uint32 = 0
                        _ = fi17
                        _ = v18
                    }
                    fmt.Println(fi14, v3)
                }
            default:
                if (true && (1e10 > v3)) {
                    var v19 uint64 = 0
                    _ = v19
                }
            }
            fmt.Println(fi14, v3)
            _ = fi14
        }
        var v20 uint = 604
        _ = v20
    }
    for fi21 := 0; fi21 < 10; fi21++ {
        v22 := GenS3{F0: 128, F1: 18446744073709551615, F2: 18446744073709551615}
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 13; fi23++ {
        switch "a" {
        case "":
            var v24 string = ("\n\t" + "foo bar")
            _ = v24
        case "hello":
            v25 := GenS4{F0: 0, F1: -1, F2: 412}
            _ = v25
        default:
            v3 = ((float64(-2.25) - 0.0) * v3)
        }
        var v26 uint16 = 32767
        _ = fi23
        _ = v26
    }
    switch "\n\t" {
    case "foo bar":
        switch 1 {
        case 128:
            switch ("\n\t" + "") {
            case "\n\t":
                v3 = float64(((-2.25 * 1.5) + 0.1))
            }
        default:
            v27 := GenS1{F0: 42, F1: 127, F2: 1}
            _ = v27
        }
    case "hello":
        for fi28 := 0; fi28 < 4; fi28++ {
            v3 -= v3
            v2 = append(v2, (-(fi28)))
            if ((fi28 >> ((1) & 63)) != (fi28 + (fi28 - fi28))) {
                v3 += (0.1 - v3)
                v29 := GenS4{F0: 255, F1: 1, F2: 42}
                switch "a" {
                case "a":
                    fmt.Println(fi28, fi28)
                }
                _ = v29
            }
            _ = fi28
        }
    }
    if true {
        switch 32768 {
        case 42:
            v30 := GenS1{F0: 785, F1: 48, F2: 748}
            _ = v30
        case 128:
            v3 = v3
        case 9223372036854775807:
            if len(v2) > 0 {
                v2[0] = 32767
            }
        }
        if (!(127 < -1)) {
            if false {
                for fi31 := 0; fi31 < 20; fi31++ {
                    var v32 int = 128
                    _ = fi31
                    _ = v32
                }
            } else {
                fmt.Println(v3)
            }
        }
        if false {
            if (666 <= 407) {
                for fi33 := 0; fi33 < 4; fi33++ {
                    if (v3 != float64(0.1)) {
                        switch (fi33 * (fi33 | (fi33 << ((923) & 63)))) {
                        case 2147483647:
                            switch (("foo bar" + "\n\t") + "foo bar") {
                            case "foo bar":
                                fmt.Println(fi33)
                            case "":
                                v34 := GenS2{F0: fi33}
                                _ = v34
                            default:
                                v3 -= (float64(0.0) * (v3 / (float64(0.0) * v3)))
                            }
                        case -529:
                            v3 = (3.14159 + ((1.5 * (v3 / 0.1)) * float64(float32(float64((float32(v3) + 3.14159))))))
                        case 42:
                            var v35 uint = 4294967295
                            _ = v35
                        default:
                            var v36 bool = true
                            _ = v36
                        }
                        var v37 uint8 = 0
                        v38 := GenS3{F0: (v37 / ((v37) | 1)), F1: 42, F2: 0}
                        _ = v37
                        _ = v38
                    } else {
                        v3 = v3
                        var v39 string = "\n\t"
                        _ = v39
                    }
                    switch -9223372036854775808 {
                    case 32767:
                        var v40 uint32 = 42
                        _ = v40
                    }
                    _ = fi33
                }
                fmt.Println(v3)
            } else {
                v3 = 3.14159
                v41 := GenS5{F0: 462, F1: false, F2: 0}
                if (1 == 1) {
                    var v42 uint8 = 128
                    var v43 uint = 2
                    v44 := GenS5{F0: -482, F1: true, F2: 4294967295}
                    _ = v42
                    _ = v43
                    _ = v44
                } else {
                    switch -1 {
                    case 128:
                        for fi45 := 0; fi45 < 12; fi45++ {
                            var v46 uint8 = 33
                            v3 -= (-2.25 / v3)
                            _ = fi45
                            _ = v46
                        }
                    case 1:
                        if (65535 >= 42) {
                            var v47 bool = ((!((((0.1 * v3) + 3.14159) - v3) >= v3)) || (false && true))
                            if ((!(2 == 585)) || (!((4294967295 < 42) && (!true)))) {
                                var v48 uint64 = 42
                                if (356 == v48) {
                                    fmt.Println(v3, v47, v47)
                                    if (((!(v3 >= v3)) || false) && false) {
                                        v47 = v47
                                        switch 42 {
                                        case 871:
                                            switch "hello" {
                                            case "hello":
                                                for fi49 := 0; fi49 < 3; fi49++ {
                                                    for fi50 := 0; fi50 < 2; fi50++ {
                                                        switch ("\n\t" + "foo bar") {
                                                        case "foo bar":
                                                            _ = len(v1)
                                                        case "hello":
                                                            v3 = 0.0
                                                        default:
                                                            var v51 uint64 = v48
                                                            _ = v51
                                                        }
                                                        _ = fi50
                                                    }
                                                    for fi52 := 0; fi52 < 9; fi52++ {
                                                        v2 = append(v2, 2147483647)
                                                        v2 = append(v2, 42)
                                                        var v53 int16 = -1
                                                        v54 := GenS5{F0: fi49, F1: v47, F2: 4294967295}
                                                        _ = fi52
                                                        _ = v53
                                                        _ = v54
                                                    }
                                                    _ = fi49
                                                }
                                            case "foo bar":
                                                fmt.Println(v3)
                                            default:
                                                v48 += (v48 | v48)
                                            }
                                        case 9223372036854775807:
                                            var v55 string = ""
                                            _ = v55
                                        }
                                        v56 := GenS5{F0: 128, F1: v47, F2: (v48 >> ((4294967295) & 63))}
                                        _ = v56
                                    } else {
                                        for fi57 := 0; fi57 < 5; fi57++ {
                                            if true {
                                                if true {
                                                    v58 := GenS2{F0: fi57}
                                                    v59 := GenS1{F0: 62, F1: 255, F2: 32767}
                                                    fmt.Println(fi57, v48, fi57)
                                                    _ = v58
                                                    _ = v59
                                                }
                                                switch fi57 {
                                                case 137:
                                                    var v60 string = "hello"
                                                    _ = v60
                                                default:
                                                    for fi61 := 0; fi61 < 2; fi61++ {
                                                        switch "hello" {
                                                        case "foo bar":
                                                            if v47 {
                                                                fmt.Println(fi61)
                                                                v3 = v3
                                                                var v62 int32 = 42
                                                                _ = v62
                                                            }
                                                        case "a":
                                                            fmt.Println(fi61)
                                                        default:
                                                            v47 = (65535 < 65535)
                                                        }
                                                        v63 := GenS6{F0: 42}
                                                        _ = fi61
                                                        _ = v63
                                                    }
                                                }
                                            } else {
                                                var v64 int16 = 57
                                                if true {
                                                    var v65 int16 = (-(v64))
                                                    var v66 uint64 = (v48 + (v48 ^ 1))
                                                    if false {
                                                        _ = len(v1)
                                                    } else {
                                                        v47 = (false || false)
                                                    }
                                                    _ = v65
                                                    _ = v66
                                                }
                                                _ = v64
                                            }
                                            switch ((("" + "hello") + ("a" + "hello")) + ("hello" + "hello")) {
                                            case "":
                                                var v67 uint = 4294967295
                                                _ = v67
                                            }
                                            _ = fi57
                                        }
                                        v68 := GenS2{F0: 9223372036854775807}
                                        v3 = -2.25
                                        _ = v68
                                    }
                                    _ = len(v1)
                                } else {
                                    v48 += v48
                                }
                                var v69 int64 = -472
                                _ = v48
                                _ = v69
                            } else {
                                v70 := v1["hello"]
                                v71 := v1[("" + (("\n\t" + "\n\t") + "\n\t"))]
                                if (42 >= 1) {
                                    v3 -= (float64(1.5) - (v3 - v3))
                                    var v72 uint8 = 127
                                    _ = v72
                                }
                                _ = v70
                                _ = v71
                            }
                            _ = v47
                        } else {
                            v73 := GenS5{F0: 32768, F1: true, F2: 42}
                            if (0.0 > -2.25) {
                                var v74 uint = 4294967295
                                v74 = (v74 << ((v74) & 63))
                                _ = v74
                            } else {
                                v75 := GenS7{F0: -1, F1: "\n\t", F2: 4}
                                v76 := GenS8{F0: 0}
                                _ = v75
                                _ = v76
                            }
                            v3 += ((float64(-2.25) - v3) + ((float64(0.0) * 1.5) - v3))
                            _ = v73
                        }
                    case 42:
                        for fi77 := 0; fi77 < 14; fi77++ {
                            v3 -= float64(1.5)
                            _ = fi77
                        }
                    default:
                        v1["a"] = 4294967295
                    }
                    var v78 float64 = -2.25
                    v79 := GenS3{F0: 0, F1: 126, F2: 499}
                    _ = v78
                    _ = v79
                }
                _ = v41
            }
        }
    } else {
        v1[""] = 2
        _ = len(v1)
        v1[("foo bar" + "foo bar")] = 1
    }
    v3 = v3
    fmt.Println(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
