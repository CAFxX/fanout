package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int8 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 bool }
type GenS5 struct { F0 int64; F1 uint32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int16 = -32768
    v3 := []string{"hello", ""}
    var v4 int64 = -9223372036854775808
    var v5 int16 = (v2 ^ v2)
    var v6 string = "a"
    var v7 int64 = v4
    fmt.Println(v7, v2, v4)
    if (((v2 != int16(v7)) && (false && true)) && false) {
        v8 := GenS1{F0: 451, F1: (v5 << ((65535) & 63)), F2: -12}
        _ = v8
    } else {
        v9 := GenS1{F0: 4294967295, F1: (^(v5)), F2: -52}
        var v10 int64 = v4
        _ = v10
        _ = v9
    }
    if (("" == "a") || (2147483648 > 4294967295)) {
        var v11 uint8 = 128
        var v12 float32 = 0.0
        v6 = v6
        _ = v11
        _ = v12
    } else {
        switch "foo bar" {
        case "a":
            for fi13 := 0; fi13 < 15; fi13++ {
                switch int(v4) {
                case -890:
                    v14 := GenS1{F0: 292, F1: (v2 % (((v5 ^ (-(v5)))) | 1)), F2: -72}
                    _ = v14
                case -2147483648:
                    for fi15 := 0; fi15 < 14; fi15++ {
                        switch (fi13 >> ((65535) & 63)) {
                        case 42:
                            v2 = v5
                        case 4294967295:
                            var v16 int8 = 0
                            v1[v6] = 0
                            if true {
                                v6 = ((v6 + "a") + v6)
                                _ = len(v1)
                            } else {
                                var v17 uint32 = 0
                                if true {
                                    v18 := GenS1{F0: 0, F1: v2, F2: v16}
                                    _ = v18
                                }
                                v1[(v6 + v6)] = (v17 % (((v17 - v17)) | 1))
                                v16 = v16
                                fmt.Println(fi15, fi15)
                                switch fi15 {
                                case 32767:
                                    v6 = v6
                                case -296:
                                    fmt.Println(fi15, fi15, fi15)
                                    v16 += 42
                                    var v19 bool = (42 == 124)
                                    _ = v19
                                default:
                                    v7 = (v7 - v7)
                                }
                                _ = v17
                            }
                            _ = v16
                        }
                        _ = fi15
                    }
                case -988:
                    v5 += 42
                }
                v20 := GenS1{F0: 608, F1: -1, F2: -98}
                _ = fi13
                _ = v20
            }
        case "\n\t":
            v21 := GenS2{F0: -1, F1: (((v6 + v6) + ("\n\t" + ("a" + (("a" + "a") + v6)))) + "foo bar")}
            _ = v21
        default:
            v6 = "hello"
        }
        switch -2147483648 {
        case -525:
            var v22 float64 = (3.14159 / 3.14159)
            _ = v22
        case -1:
            if (v5 < (v2 << ((761) & 63))) {
                v1["foo bar"] = 1
                v7 -= (^(v7))
                var v23 float64 = 3.14159
                _ = v23
            }
        case 2147483647:
            if (266 > (v5 + (v5 / ((962) | 1)))) {
                var v24 int16 = v5
                v7 = -1
                _ = v24
            }
        default:
            var v25 uint64 = 0
            _ = v25
        }
    }
    if true {
        var v26 uint32 = 4294967295
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v26
    } else {
        for fi27 := 0; fi27 < 5; fi27++ {
            switch ((v6 + (("foo bar" + v6) + ("foo bar" + (v6 + (v6 + v6))))) + "hello") {
            case "\n\t":
                var v28 bool = true
                _ = v28
            case "hello":
                switch (-(fi27)) {
                case 32767:
                    _ = len(v1)
                case -752:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            case "a":
                v7 += v4
            default:
                switch 63 {
                case 288:
                    for fi29 := 0; fi29 < 19; fi29++ {
                        v30 := GenS1{F0: 0, F1: v2, F2: -1}
                        delete(v1, v6)
                        v1[("" + v6)] = 65535
                        _ = fi29
                        _ = v30
                    }
                case 32767:
                    var v31 float32 = 1.5
                    _ = v31
                case 667:
                    v4 += -894
                default:
                    switch fi27 {
                    case 32768:
                        if true {
                            if (!(v7 < (v7 << ((4294967295) & 63)))) {
                                var v32 float64 = 0.0
                                _ = v32
                            }
                            v1["hello"] = 2147483648
                            var v33 int32 = -820
                            _ = v33
                        }
                    }
                }
            }
            _ = len(v3)
            switch fi27 {
            case -700:
                v2 -= 32767
            case 2147483647:
                if true {
                    fmt.Println(v6, v4)
                    for fi34 := 0; fi34 < 8; fi34++ {
                        if (!(42 > 1)) {
                            v35 := GenS1{F0: 42, F1: (v5 - v2), F2: -81}
                            var v36 string = v6
                            _ = v35
                            _ = v36
                        } else {
                            v1[v6] = 111
                            if (!true) {
                                if (v2 > v2) {
                                    switch -2147483648 {
                                    case 444:
                                        v37 := GenS3{F0: 89}
                                        _ = v37
                                    case 231:
                                        fmt.Println(fi34, v2, v6)
                                    case 248:
                                        delete(v1, v6)
                                    }
                                    fmt.Println(v2)
                                    var v38 bool = (v4 <= v7)
                                    v7 = v7
                                    var v39 int64 = v7
                                    _ = v38
                                    _ = v39
                                }
                            } else {
                                if (32767 != (v2 % ((v5) | 1))) {
                                    v40 := v1[(v6 + "hello")]
                                    var v41 uint32 = 42
                                    if true {
                                        var v42 int32 = -1
                                        v43 := GenS3{F0: 100}
                                        v3 = append(v3, v6)
                                        _ = v42
                                        _ = v43
                                    }
                                    _ = v40
                                    _ = v41
                                } else {
                                    if false {
                                        var v44 uint64 = 145
                                        if true {
                                            v45 := GenS4{F0: true, F1: true}
                                            fmt.Println(v7, v2, v7)
                                            v44 -= (v44 << ((925) & 63))
                                            _ = v45
                                        }
                                        if false {
                                            if (1 == (-(v4))) {
                                                if ((337 > 32767) && false) {
                                                    v7 -= -9223372036854775808
                                                }
                                                v2 += 42
                                            }
                                            var v46 int32 = -452
                                            _ = v46
                                        }
                                        if false {
                                            switch ("foo bar" + v6) {
                                            case "hello":
                                                v47 := GenS4{F0: ((v6 + "\n\t") >= v6), F1: true}
                                                _ = v47
                                            case "a":
                                                v1[v6] = 4294967295
                                            }
                                            v1[("\n\t" + v6)] = 0
                                            if (v44 > v44) {
                                                var v48 int64 = -9223372036854775808
                                                var v49 uint16 = 255
                                                v7 = v48
                                                switch "\n\t" {
                                                case "":
                                                    fmt.Println(v6)
                                                    v7 = v7
                                                    _ = len(v3)
                                                }
                                                _ = v48
                                                _ = v49
                                            }
                                            var v50 uint64 = (^(v44))
                                            switch v6 {
                                            case "hello":
                                                v2 = 1
                                            }
                                            v50 = v44
                                            v1[v6] = 2147483648
                                            var v51 uint64 = (^(v50))
                                            _ = v50
                                            _ = v51
                                        } else {
                                            fmt.Println(v44, fi34)
                                            var v52 uint16 = 32767
                                            _ = v52
                                        }
                                        v6 = (v6 + "a")
                                        _ = v44
                                    } else {
                                        v53 := GenS5{F0: -187, F1: 65535}
                                        _ = v53
                                    }
                                    v54 := GenS3{F0: 127}
                                    var v55 string = (("" + "") + ("hello" + (v6 + v6)))
                                    _ = v54
                                    _ = v55
                                }
                                _ = len(v3)
                            }
                            _ = len(v1)
                        }
                        _ = fi34
                    }
                }
            default:
                v6 = ("foo bar" + v6)
            }
            _ = fi27
        }
    }
    for fi56 := 0; fi56 < 14; fi56++ {
        for fi57 := 0; fi57 < 15; fi57++ {
            switch ("\n\t" + (v6 + "")) {
            case "\n\t":
                var v58 string = v6
                _ = v58
            case "a":
                v4 = v4
            }
            v59 := GenS2{F0: 609, F1: v6}
            v7 += v7
            _ = fi57
            _ = v59
        }
        v6 = v6
        if len(v3) > 0 {
            v3[0] = ("\n\t" + "foo bar")
        }
        _ = fi56
    }
    fmt.Println(v7, v5, v4)
    if (65535 != 171) {
        v4 -= -9223372036854775808
        _ = len(v1)
        if true {
            v60 := GenS1{F0: 2, F1: v2, F2: 0}
            _ = v60
        }
    }
    var v61 int16 = v5
    switch ("hello" + "a") {
    case "foo bar":
        switch v6 {
        case "a":
            var v62 int = -887
            _ = v62
        }
    case "":
        v63 := GenS2{F0: -1, F1: "foo bar"}
        _ = v63
    }
    switch "\n\t" {
    case "foo bar":
        v64 := GenS4{F0: ((v7 * -1) != v7), F1: (!false)}
        _ = v64
    }
    var v65 int = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v61:", v61)
    fmt.Println("v65:", v65)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
    _ = v65
    _ = v7
}
