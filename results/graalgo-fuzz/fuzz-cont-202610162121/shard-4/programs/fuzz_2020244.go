package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 int32; F2 uint }
type GenS4 struct { F0 uint16; F1 int8; F2 bool }
type GenS5 struct { F0 uint32; F1 string }
type GenS6 struct { F0 uint; F1 uint32 }
type GenS7 struct { F0 int; F1 uint8 }
func genFn0(p0 int8, p1 uint32) uint {
    var v4 int64 = -9223372036854775808
    for fi5 := 0; fi5 < 6; fi5++ {
        v6 := []int{32767}
        v7 := GenS1{F0: (2 == 1)}
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 string = "a"
        _ = fi8
        _ = v9
    }
    switch ((("a" + "") + "foo bar") + "a") {
    case "\n\t":
        if (-54 >= int8(v4)) {
            if true {
                var v10 float32 = 1.5
                var v11 int16 = 32767
                if false {
                    if (!true) {
                        if (!(2 <= 65535)) {
                            v12 := GenS2{F0: (-(v4))}
                            _ = v12
                        }
                    } else {
                        v10 -= 1e10
                    }
                    var v13 uint16 = uint16(p1)
                    v14 := GenS3{F0: "a", F1: 42, F2: 38}
                    _ = v13
                    _ = v14
                }
                _ = v10
                _ = v11
            }
        } else {
            v15 := GenS2{F0: v4}
            _ = v15
        }
    case "foo bar":
        p1 = (p1 & (^(p1)))
    case "a":
        var v16 uint16 = 2
        _ = v16
    }
    var v17 bool = false
    _ = p0
    _ = p1
    _ = v17
    _ = v4
    return 2
}

func main() {
    var v1 float32 = 0.1
    var v2 bool = true
    var v3 int64 = 42
    for fi18 := 0; fi18 < 16; fi18++ {
        v19 := GenS1{F0: (!((("" + "\n\t") > ("" + "hello")) || (0 >= -10)))}
        v20 := GenS3{F0: ("\n\t" + "foo bar"), F1: 679, F2: 2}
        _ = fi18
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 6; fi21++ {
        switch ("foo bar" + "a") {
        case "":
            v1 = ((v1 / 1.0) / v1)
        case "\n\t":
            if (fi21 > fi21) {
                var v22 float64 = 1.5
                for fi23 := 0; fi23 < 2; fi23++ {
                    var v24 uint64 = 0
                    _ = fi23
                    _ = v24
                }
                _ = v22
            }
        case "hello":
            v2 = v2
        }
        _ = fi21
    }
    if v2 {
        v25 := []int{}
        v2 = (1 < -1)
        if v2 {
            fmt.Println(v1, v3)
            if v2 {
                var v26 int16 = 32767
                for fi27 := 0; fi27 < 13; fi27++ {
                    for fi28 := 0; fi28 < 3; fi28++ {
                        v29 := GenS2{F0: v3}
                        fmt.Println(fi27, v3)
                        switch fi27 {
                        case -9223372036854775808:
                            var v30 uint16 = 255
                            _ = v30
                        case -39:
                            v25 = append(v25, (fi28 ^ (fi27 - (fi27 % ((-9223372036854775808) | 1)))))
                        case 42:
                            if false {
                                switch ("foo bar" + "") {
                                case "foo bar":
                                    switch "foo bar" {
                                    case "a":
                                        switch (fi28 ^ fi28) {
                                        case 32767:
                                            var v31 uint64 = 4294967295
                                            _ = v31
                                        default:
                                            v32 := GenS1{F0: v2}
                                            _ = v32
                                        }
                                    case "foo bar":
                                        var v33 int64 = v3
                                        _ = v33
                                    case "":
                                        switch (("a" + "\n\t") + "a") {
                                        case "\n\t":
                                            if (-1 > -1) {
                                                fmt.Println(fi28, v2, v26)
                                                v1 = float32((1.5 / 1e10))
                                            }
                                            if true {
                                                var v34 bool = v2
                                                _ = v34
                                            }
                                        case "a":
                                            switch (fi28 << ((0) & 63)) {
                                            case 32768:
                                                var v35 bool = ((false && v2) || v2)
                                                _ = v35
                                            }
                                        }
                                        if len(v25) >= 2 {
                                            v25 = v25[1:len(v25)]
                                        }
                                    default:
                                        var v36 int8 = -40
                                        if len(v25) > 0 {
                                            v25[0] = (^(fi28))
                                        }
                                        if (!(255 != 65535)) {
                                            if len(v25) > 0 {
                                                v25[0] = (fi28 - fi27)
                                            }
                                            if len(v25) > 0 {
                                                v25[0] = fi28
                                            }
                                        } else {
                                            if (65535 == genFn0(0, 950)) {
                                                v2 = (!v2)
                                                if len(v25) > 0 {
                                                    v25[0] = (fi27 & (-(fi28)))
                                                }
                                            } else {
                                                v36 += v36
                                                switch fi28 {
                                                case 127:
                                                    var v37 int32 = 42
                                                    _ = v37
                                                case 32768:
                                                    var v38 uint = 255
                                                    _ = v38
                                                case 577:
                                                    v36 = (v36 >> ((460) & 63))
                                                default:
                                                    var v39 uint = genFn0(v36, 695)
                                                    _ = v39
                                                }
                                                switch (fi27 + fi28) {
                                                case 32768:
                                                    v25 = append(v25, (-(fi27)))
                                                case 4294967295:
                                                    var v40 int16 = (^(v26))
                                                    _ = v40
                                                case -1:
                                                    var v41 string = ""
                                                    _ = v41
                                                }
                                            }
                                        }
                                        _ = v36
                                    }
                                }
                                v2 = true
                                fmt.Println(v2, v1)
                                v26 += (v26 * (v26 ^ v26))
                            } else {
                                if len(v25) > 0 {
                                    v25[0] = fi27
                                }
                            }
                        }
                        v42 := GenS4{F0: 442, F1: -1, F2: false}
                        fmt.Println(v1, v1)
                        var v43 int64 = v3
                        _ = fi28
                        _ = v29
                        _ = v42
                        _ = v43
                    }
                    switch fi27 {
                    case 9223372036854775807:
                        v44 := GenS4{F0: 65535, F1: 60, F2: ((((!(2 < 18446744073709551615)) && ((!v2) || (1 < 1))) || (((0.1 * 0.1) + ((3.14159 + 0.1) - 3.14159)) >= (0.1 - ((1.5 - 0.1) - 0.1)))) || false)}
                        _ = v44
                    default:
                        v45 := GenS5{F0: 0, F1: ("" + "\n\t")}
                        _ = v45
                    }
                    _ = fi27
                }
                _ = v26
            } else {
                v46 := GenS4{F0: 42, F1: 0, F2: v2}
                var v47 int16 = -1
                if ("\n\t" < ("hello" + (("a" + ("" + "\n\t")) + (("" + ("" + (("hello" + "hello") + ""))) + ("hello" + ("hello" + "a")))))) {
                    var v48 int64 = v3
                    for fi49 := 0; fi49 < 2; fi49++ {
                        v47 += v47
                        switch "foo bar" {
                        case "":
                            for fi50 := 0; fi50 < 15; fi50++ {
                                if ((255 == 221) || v2) {
                                    v1 = (3.14159 - (v1 + (1.5 * 3.14159)))
                                    switch (("" + "\n\t") + "") {
                                    case "":
                                        v51 := GenS6{F0: 0, F1: 959}
                                        _ = v51
                                    case "foo bar":
                                        switch (fi49 >> ((genFn0(100, 2147483647)) & 63)) {
                                        case 355:
                                            var v52 uint64 = 1
                                            _ = v52
                                        case 0:
                                            if (false || (3.14159 <= 0.1)) {
                                                if len(v25) > 0 {
                                                    v25[0] = fi49
                                                }
                                                var v53 string = "a"
                                                switch fi50 {
                                                case 4294967295:
                                                    var v54 uint8 = 45
                                                    _ = v54
                                                case -2147483648:
                                                    fmt.Println(fi49, fi49, v47)
                                                case -9223372036854775808:
                                                    v47 = -944
                                                default:
                                                    v48 += (v3 & v3)
                                                    fmt.Println(v1, v3)
                                                    var v55 int = fi50
                                                    var v56 bool = ((0 <= -1) && v2)
                                                    v57 := GenS3{F0: "foo bar", F1: 0, F2: 2}
                                                    fmt.Println(v3, fi49, v56)
                                                    switch (fi49 - 591) {
                                                    case 935:
                                                        fmt.Println(v48, fi50)
                                                    case 891:
                                                        var v58 uint16 = 42
                                                        _ = v58
                                                    case 0:
                                                        if true {
                                                            _ = len(v25)
                                                            var v59 float64 = 1.5
                                                            v48 = v3
                                                            _ = v59
                                                        }
                                                    default:
                                                        v53 = (v53 + v53)
                                                    }
                                                    _ = v55
                                                    _ = v56
                                                    _ = v57
                                                }
                                                v2 = v2
                                                if v2 {
                                                    v25 = append(v25, 32767)
                                                    if len(v25) > 0 {
                                                        v25[0] = fi50
                                                    }
                                                    v53 = ("hello" + "\n\t")
                                                    v47 += v47
                                                    switch fi49 {
                                                    case -1:
                                                        var v60 string = "hello"
                                                        _ = v60
                                                    case -9223372036854775808:
                                                        v3 -= (-(v48))
                                                    case 0:
                                                        v53 = ""
                                                    default:
                                                        v1 -= float32(float64(((v1 + 1.5) * (v1 * 1e10))))
                                                    }
                                                }
                                                _ = v53
                                            }
                                        }
                                    }
                                    fmt.Println(v47)
                                }
                                fmt.Println(v47, v47, v3)
                                _ = fi50
                            }
                        }
                        var v61 float32 = 1.5
                        _ = fi49
                        _ = v61
                    }
                    v47 -= v47
                    _ = v48
                }
                _ = v46
                _ = v47
            }
            if len(v25) >= 2 {
                v25 = v25[1:len(v25)]
            }
        } else {
            var v62 float64 = float64(v1)
            _ = v62
        }
        _ = v25
    }
    for fi63 := 0; fi63 < 4; fi63++ {
        v3 += -222
        v64 := []int{}
        var v65 int8 = 127
        _ = fi63
        _ = v64
        _ = v65
    }
    if false {
        if v2 {
            var v66 int16 = -90
            v67 := []int{}
            _ = v66
            _ = v67
        }
        switch "\n\t" {
        case "foo bar":
            var v68 int16 = 1000
            _ = v68
        case "\n\t":
            for fi69 := 0; fi69 < 10; fi69++ {
                if (((((fi69 - 174) > fi69) && ((4294967295 == 838) || false)) || v2) || (!v2)) {
                    var v70 uint = 1
                    for fi71 := 0; fi71 < 18; fi71++ {
                        v2 = (!(!((v70 >> ((262) & 63)) != v70)))
                        v3 -= v3
                        v72 := GenS5{F0: 2, F1: ""}
                        _ = fi71
                        _ = v72
                    }
                    v2 = (true || v2)
                    _ = v70
                }
                v1 += v1
                _ = fi69
            }
        case "a":
            v73 := []int{42, 42, 32767}
            _ = v73
        }
        var v74 bool = (-128 != -1)
        _ = v74
    }
    switch "a" {
    case "a":
        switch ("a" + "") {
        case "\n\t":
            fmt.Println(v1, v3, v3)
        }
    case "foo bar":
        if false {
            var v75 float32 = v1
            _ = v75
        } else {
            var v76 uint8 = 42
            _ = v76
        }
    case "hello":
        v77 := GenS2{F0: (v3 ^ (v3 / ((v3) | 1)))}
        _ = v77
    default:
        v1 += (float32(-2.25) - (float32(float64((v1 - v1))) - ((v1 + v1) / (1e10 + 1.5))))
    }
    v78 := GenS5{F0: 0, F1: ("foo bar" + "foo bar")}
    var v79 uint32 = 1
    fmt.Println(v1, v3, v79)
    v2 = v2
    v80 := GenS7{F0: -1, F1: 255}
    v2 = (1 <= 758)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v78
    _ = v79
    _ = v80
}
