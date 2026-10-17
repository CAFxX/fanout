package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 uint64; F2 uint32 }
type GenS3 struct { F0 bool; F1 uint16; F2 uint }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32; F1 uint32; F2 uint }
type GenS6 struct { F0 uint }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]string)
    v3 := []string{"hello", "a", ""}
    var v4 uint64 = 18446744073709551615
    var v5 uint = 4294967295
    v6 := []string{}
    v5 += (v5 + v5)
    for fi7 := 0; fi7 < 8; fi7++ {
        v6 = append(v6, "hello")
        _ = fi7
    }
    for fi8 := 0; fi8 < 14; fi8++ {
        v9 := GenS1{F0: 815}
        var v10 int32 = 2147483647
        _ = fi8
        _ = v10
        _ = v9
    }
    v5 = v5
    if (42 >= 100) {
        for fi11 := 0; fi11 < 13; fi11++ {
            v4 = (v4 & v4)
            _ = fi11
        }
    } else {
        _ = len(v3)
        v5 += (v5 - v5)
        if false {
            var v12 uint16 = 65535
            fmt.Println(v5, v12, v12)
            _ = v12
        }
    }
    var v13 float64 = 1e10
    switch "hello" {
    case "\n\t":
        if len(v6) > 0 {
            v6[0] = "a"
        }
    }
    switch "" {
    case "hello":
        if len(v3) > 0 {
            v3[0] = ""
        }
    case "foo bar":
        switch ("foo bar" + (("" + "a") + ("\n\t" + "hello"))) {
        case "hello":
            var v14 uint32 = 1
            _ = v14
        case "":
            switch "\n\t" {
            case "foo bar":
                v6 = append(v6, "a")
            }
        case "foo bar":
            v13 = ((v13 + (v13 / v13)) / 1.0)
        }
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        for fi16 := 0; fi16 < 10; fi16++ {
            v17 := GenS1{F0: 215}
            delete(v1, 1)
            _ = len(v6)
            _ = fi16
            _ = v17
        }
        v5 += v5
        fmt.Println(v13)
        _ = fi15
    }
    switch -1 {
    case 127:
        for fi18 := 0; fi18 < 19; fi18++ {
            fmt.Println(v5, fi18)
            if true {
                for fi19 := 0; fi19 < 17; fi19++ {
                    if (1 < 575) {
                        v20 := v1[2]
                        switch (fi19 << ((v5) & 63)) {
                        case 127:
                            if (true && false) {
                                if (1 <= 100) {
                                    var v21 int16 = -32768
                                    if ("\n\t" != "foo bar") {
                                        fmt.Println(v20)
                                        v22 := GenS1{F0: (^(v21))}
                                        _ = v22
                                    } else {
                                        var v23 uint8 = 92
                                        switch fi18 {
                                        case 127:
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                        case 2147483647:
                                            var v24 int8 = 127
                                            _ = v24
                                        case 42:
                                            if true {
                                                fmt.Println(fi18, v20, v20)
                                                if false {
                                                    switch (fi18 << ((255) & 63)) {
                                                    case 0:
                                                        v2[v20] = "foo bar"
                                                    default:
                                                        if len(v6) >= 2 {
                                                            v6 = v6[1:len(v6)]
                                                        }
                                                    }
                                                } else {
                                                    v13 = v13
                                                    var v25 string = ("foo bar" + (("\n\t" + (("hello" + ((("" + (("a" + ("hello" + "")) + ((("foo bar" + ("foo bar" + ("" + (("" + ("\n\t" + ("\n\t" + ""))) + "a")))) + "\n\t") + "foo bar"))) + "hello") + "a")) + "\n\t")) + "hello"))
                                                    _ = v25
                                                }
                                            } else {
                                                if false {
                                                    if true {
                                                        v26 := GenS1{F0: (v21 / (((-(v21))) | 1))}
                                                        _ = v26
                                                    }
                                                }
                                                fmt.Println(v5, fi19)
                                                _ = len(v6)
                                                var v27 float64 = v13
                                                _ = v27
                                            }
                                        }
                                        switch (fi19 - fi19) {
                                        case 9223372036854775807:
                                            v28 := GenS1{F0: 42}
                                            _ = v28
                                        case 2147483647:
                                            fmt.Println(v21, v23, v23)
                                        case 334:
                                            v13 = ((v13 * v13) / -2.25)
                                        default:
                                            var v29 string = ""
                                            _ = v29
                                        }
                                        var v30 int32 = 2147483647
                                        _ = v23
                                        _ = v30
                                    }
                                    switch (fi18 & (fi18 + -885)) {
                                    case 32768:
                                        var v31 uint16 = 655
                                        _ = v31
                                    case -2147483648:
                                        var v32 int64 = -9223372036854775808
                                        _ = v32
                                    default:
                                        v33 := GenS1{F0: (v21 + (v21 | v21))}
                                        _ = v33
                                    }
                                    _ = v21
                                }
                            } else {
                                var v34 int8 = 127
                                var v35 bool = false
                                fmt.Println(v20, v34)
                                _ = v34
                                _ = v35
                            }
                        case 11:
                            if false {
                                v36 := GenS2{F0: (42 > -1), F1: (^(v4)), F2: (v20 << ((887) & 63))}
                                _ = len(v1)
                                _ = v36
                            }
                        default:
                            if (42 > 65535) {
                                v20 = (v20 / (((v20 * (v20 << ((v5) & 63)))) | 1))
                                var v37 int64 = -650
                                _ = v37
                            } else {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                switch (fi18 | (fi19 / (((fi19 ^ -9223372036854775808)) | 1))) {
                                case 42:
                                    v4 = 0
                                case -993:
                                    fmt.Println(v13, v5, v4)
                                }
                                v20 += v20
                            }
                        }
                        var v38 bool = true
                        _ = v20
                        _ = v38
                    }
                    switch (("" + "foo bar") + "a") {
                    case "\n\t":
                        v39 := GenS2{F0: (((v4 % ((42) | 1)) == (v4 | v4)) || true), F1: 42, F2: uint32(v4)}
                        _ = v39
                    case "":
                        v40 := v1[2]
                        _ = v40
                    case "a":
                        v41 := GenS3{F0: true, F1: 822, F2: 4294967295}
                        _ = v41
                    }
                    if false {
                        var v42 float64 = v13
                        _ = v42
                    } else {
                        var v43 int32 = -59
                        _ = v43
                    }
                    _ = fi19
                }
            }
            _ = fi18
        }
    case -9223372036854775808:
        for fi44 := 0; fi44 < 10; fi44++ {
            if true {
                switch fi44 {
                case 42:
                    var v45 uint8 = 1
                    _ = v45
                case -760:
                    v13 = v13
                case -228:
                    if true {
                        v13 += 1e10
                        var v46 uint32 = 2147483648
                        _ = v46
                    } else {
                        var v47 int64 = -1000
                        _ = v47
                    }
                default:
                    var v48 int = -9223372036854775808
                    _ = v48
                }
            }
            for fi49 := 0; fi49 < 16; fi49++ {
                var v50 int64 = -9223372036854775808
                v51 := GenS4{F0: fi49}
                _ = fi49
                _ = v50
                _ = v51
            }
            _ = fi44
        }
    case 32767:
        var v52 int = 0
        _ = v52
    }
    v5 = (v5 >> ((v5) & 63))
    if (-32768 > 84) {
        v6 = append(v6, ("\n\t" + ("a" + "foo bar")))
        v4 -= (v4 * 0)
    } else {
        v53 := GenS2{F0: (v5 >= v5), F1: 0, F2: 4294967295}
        _ = v53
    }
    v4 = v4
    v5 -= (v5 - uint(v4))
    for fi54 := 0; fi54 < 6; fi54++ {
        for fi55 := 0; fi55 < 16; fi55++ {
            switch "hello" {
            case "hello":
                if (-498 != 0) {
                    v4 -= v4
                } else {
                    fmt.Println(v13)
                }
            case "foo bar":
                v13 = v13
            default:
                delete(v2, 2)
            }
            if false {
                _ = len(v1)
                var v56 int = fi55
                switch (((("a" + "\n\t") + "\n\t") + (("\n\t" + "a") + "hello")) + "a") {
                case "a":
                    if true {
                        if false {
                            fmt.Println(v56, v5)
                        } else {
                            v13 += v13
                            var v57 int16 = 1
                            var v58 uint = 42
                            _ = v57
                            _ = v58
                        }
                    }
                case "\n\t":
                    v13 -= v13
                default:
                    if false {
                        _ = len(v6)
                        v59 := GenS5{F0: 42, F1: 770, F2: 42}
                        if false {
                            var v60 int8 = 1
                            v5 += (v5 & (v5 % ((2) | 1)))
                            v61 := GenS2{F0: true, F1: (v4 - (^(v4))), F2: 2147483647}
                            _ = v60
                            _ = v61
                        } else {
                            fmt.Println(fi55, fi54)
                        }
                        _ = v59
                    }
                }
                _ = v56
            }
            fmt.Println(fi54)
            v5 = v5
            _ = fi55
        }
        _ = fi54
    }
    if false {
        if (v13 > (v13 / v13)) {
            fmt.Println(v13)
            var v62 int32 = 2147483647
            _ = v62
        } else {
            var v63 float32 = (1.5 - 0.0)
            var v64 int32 = 1
            _ = v63
            _ = v64
        }
        v5 = (^(v5))
        fmt.Println(v4, v5)
    }
    switch "foo bar" {
    case "a":
        var v65 uint64 = v4
        _ = v65
    case "\n\t":
        if false {
            v3 = append(v3, "")
            v66 := GenS6{F0: v5}
            v4 += v4
            _ = v66
        }
    case "":
        v4 = v4
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
