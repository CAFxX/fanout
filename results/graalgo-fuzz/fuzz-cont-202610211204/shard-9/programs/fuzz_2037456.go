package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint16; F1 uint8; F2 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16; F1 uint64 }
type GenS5 struct { F0 string; F1 string }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int64; F1 uint64; F2 int8 }
type GenS8 struct { F0 string; F1 bool }

func main() {
    v1 := []int{32767, 9223372036854775807, 1, -2147483648}
    v2 := make(map[int]uint32)
    v3 := []uint32{0, 2147483647, 144, 366}
    var v4 string = (("" + "\n\t") + "a")
    v5 := []int{32767, 1}
    delete(v2, -898)
    for fi6 := 0; fi6 < 1; fi6++ {
        for fi7 := 0; fi7 < 20; fi7++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v4 = ""
            _ = fi7
        }
        for fi8 := 0; fi8 < 14; fi8++ {
            var v9 int16 = 740
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 5; fi10++ {
            if (!(true && true)) {
                var v11 uint = 18446744073709551615
                _ = v11
            } else {
                v4 = v4
                v4 = (v4 + v4)
                _ = len(v2)
                v4 = (("" + "hello") + v4)
            }
            _ = fi10
        }
        _ = fi6
    }
    v12 := GenS1{F0: -88, F1: 19}
    v4 = "hello"
    if true {
        var v13 uint = 0
        fmt.Println(v4, v4)
        _ = v13
    }
    if (1.5 <= 0.1) {
        for fi14 := 0; fi14 < 16; fi14++ {
            _ = len(v5)
            v4 = v4
            v15 := v2[4294967295]
            _ = fi14
            _ = v15
        }
    } else {
        switch -551 {
        case 836:
            _ = len(v2)
        case 32768:
            if true {
                switch v4 {
                case "":
                    for fi16 := 0; fi16 < 20; fi16++ {
                        var v17 uint8 = 0
                        if false {
                            _ = len(v2)
                        } else {
                            switch (fi16 << ((0) & 63)) {
                            case -453:
                                v4 = "\n\t"
                            case -2147483648:
                                v4 = (v4 + (v4 + ("a" + "hello")))
                            case -1:
                                v4 = ("hello" + ("\n\t" + (("hello" + ("\n\t" + v4)) + "\n\t")))
                            }
                        }
                        _ = fi16
                        _ = v17
                    }
                case "\n\t":
                    for fi18 := 0; fi18 < 9; fi18++ {
                        for fi19 := 0; fi19 < 3; fi19++ {
                            switch (^(fi18)) {
                            case 128:
                                switch (fi19 * fi19) {
                                case -1:
                                    var v20 uint = 42
                                    _ = v20
                                default:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                            }
                            v4 = v4
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            switch v4 {
                            case "a":
                                v21 := v2[(fi19 ^ fi18)]
                                v4 = "a"
                                v22 := GenS2{F0: 1, F1: 1, F2: 1}
                                _ = v21
                                _ = v22
                            }
                            _ = fi19
                        }
                        var v23 uint64 = 42
                        _ = fi18
                        _ = v23
                    }
                case "a":
                    if (!true) {
                        var v24 float64 = 0.0
                        fmt.Println(v4, v24)
                        if false {
                            v25 := v2[9223372036854775807]
                            fmt.Println(v24, v4, v4)
                            if false {
                                var v26 int32 = 0
                                fmt.Println(v25, v26)
                                _ = v26
                            }
                            _ = v25
                        }
                        _ = v24
                    } else {
                        v4 = v4
                    }
                default:
                    switch ("a" + v4) {
                    case "hello":
                        for fi27 := 0; fi27 < 16; fi27++ {
                            var v28 int = (fi27 + (fi27 >> ((2) & 63)))
                            v28 += v28
                            _ = fi27
                            _ = v28
                        }
                    case "\n\t":
                        v29 := GenS2{F0: 32767, F1: 1, F2: 1}
                        _ = v29
                    default:
                        var v30 float64 = ((0.0 / 3.14159) * 0.0)
                        _ = v30
                    }
                }
                delete(v2, 651)
            } else {
                var v31 string = (v4 + "foo bar")
                var v32 uint32 = 1
                _ = v31
                _ = v32
            }
        case 875:
            switch "a" {
            case "hello":
                switch ("" + v4) {
                case "foo bar":
                    var v33 int16 = -32768
                    _ = v33
                case "\n\t":
                    if (true && (!false)) {
                        if len(v1) > 0 {
                            v1[0] = 591
                        }
                        var v34 int8 = 120
                        _ = v34
                    }
                }
            case "foo bar":
                v4 = v4
            case "":
                switch (v4 + "\n\t") {
                case "\n\t":
                    for fi35 := 0; fi35 < 9; fi35++ {
                        v4 = v4
                        _ = fi35
                    }
                }
            }
        }
    }
    fmt.Println(v4)
    var v36 uint32 = 2147483648
    fmt.Println(v36, v36)
    fmt.Println(v4, v4)
    v36 = v36
    if (v4 > "\n\t") {
        var v37 uint16 = 32767
        _ = v37
    } else {
        for fi38 := 0; fi38 < 2; fi38++ {
            for fi39 := 0; fi39 < 3; fi39++ {
                fmt.Println(v36, v4, v4)
                switch (((v4 + ((("hello" + ("foo bar" + (("\n\t" + v4) + ("\n\t" + "")))) + v4) + ("hello" + (v4 + v4)))) + v4) + v4) {
                case "foo bar":
                    if len(v5) > 0 {
                        v5[0] = 0
                    }
                case "":
                    var v40 uint8 = 60
                    _ = v40
                default:
                    switch v4 {
                    case "\n\t":
                        switch (fi39 + (-(fi39))) {
                        case 128:
                            v36 -= (^(v36))
                        case 441:
                            fmt.Println(fi39)
                        case 524:
                            var v41 int = (fi38 / ((-1) | 1))
                            if false {
                                if (18446744073709551615 == 4294967295) {
                                    if false {
                                        v36 = 930
                                        switch (((v4 + v4) + "\n\t") + v4) {
                                        case "foo bar":
                                            switch v4 {
                                            case "foo bar":
                                                if true {
                                                    if (!(!(!(!(42 <= 127))))) {
                                                        var v42 uint16 = 2
                                                        _ = v42
                                                    }
                                                } else {
                                                    if true {
                                                        v43 := GenS3{F0: 1}
                                                        var v44 float32 = (float32(((1.5 + 0.1) + 1.5)) / (((float32(1.5) * (float32(float64((0.1 / (((1.5 * ((0.1 * (3.14159 / float32((1e10 + 0.0)))) - 1e10)) - 0.0) - float32(0.1))))) - -2.25)) / (((0.0 / (-2.25 - 3.14159)) - 3.14159) / 1.5)) * 0.1))
                                                        v45 := GenS4{F0: 178, F1: 2}
                                                        _ = v43
                                                        _ = v44
                                                        _ = v45
                                                    } else {
                                                        var v46 uint64 = 18
                                                        var v47 uint8 = 2
                                                        v46 += (v46 >> ((65535) & 63))
                                                        _ = v46
                                                        _ = v47
                                                    }
                                                    v36 = 65535
                                                    v4 = ("foo bar" + v4)
                                                }
                                            default:
                                                var v48 bool = (!(1 != -818))
                                                _ = v48
                                            }
                                        }
                                    } else {
                                        var v49 int64 = 677
                                        var v50 uint64 = 42
                                        _ = v49
                                        _ = v50
                                    }
                                    _ = len(v3)
                                    fmt.Println(fi38)
                                }
                            }
                            var v51 int8 = 127
                            _ = v41
                            _ = v51
                        }
                    case "a":
                        if (!true) {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        } else {
                            var v52 string = "\n\t"
                            v53 := GenS1{F0: 0, F1: 2}
                            _ = v52
                            _ = v53
                        }
                    default:
                        v36 = 4294967295
                    }
                }
                _ = fi39
            }
            if false {
                if (!((9223372036854775807 <= 0) || ((!(423 == (^(v36)))) || true))) {
                    for fi54 := 0; fi54 < 2; fi54++ {
                        v36 = (v36 | v36)
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        v55 := GenS5{F0: (v4 + "hello"), F1: "hello"}
                        _ = fi54
                        _ = v55
                    }
                    if false {
                        v4 = "foo bar"
                        v5 = append(v5, (^(fi38)))
                    } else {
                        if (4294967295 >= 65535) {
                            v36 -= (v36 & v36)
                            var v56 uint16 = 518
                            var v57 int = fi38
                            _ = v56
                            _ = v57
                        }
                        v58 := v2[(fi38 << ((1) & 63))]
                        _ = v58
                    }
                }
                if (!(42 < 2)) {
                    v36 -= (v36 ^ 617)
                    v36 = 655
                    if ("a" <= v4) {
                        v4 = ((v4 + (v4 + v4)) + v4)
                        for fi59 := 0; fi59 < 18; fi59++ {
                            v36 += 463
                            v2[(fi38 >> ((18446744073709551615) & 63))] = v36
                            v60 := GenS1{F0: -376, F1: 127}
                            v61 := GenS4{F0: 255, F1: 0}
                            _ = fi59
                            _ = v60
                            _ = v61
                        }
                        v4 = "\n\t"
                    } else {
                        v62 := GenS6{F0: 1000}
                        var v63 float32 = (-2.25 + ((0.1 / (3.14159 + float32(1.5))) + (3.14159 * -2.25)))
                        var v64 uint = 65535
                        _ = v62
                        _ = v63
                        _ = v64
                    }
                }
            } else {
                for fi65 := 0; fi65 < 10; fi65++ {
                    var v66 uint32 = v36
                    switch (^(fi38)) {
                    case 127:
                        if (true || true) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        } else {
                            v67 := GenS7{F0: -9223372036854775808, F1: 0, F2: 0}
                            var v68 uint16 = 0
                            _ = v67
                            _ = v68
                        }
                    case -207:
                        if (("foo bar" < ("\n\t" + v4)) || true) {
                            var v69 int64 = 9223372036854775807
                            _ = v69
                        } else {
                            if ((-1 != -1) && (42 > 0)) {
                                _ = len(v2)
                            } else {
                                fmt.Println(fi38)
                                v66 += 2147483647
                            }
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                    case 32767:
                        var v70 int32 = -1
                        _ = v70
                    default:
                        _ = len(v3)
                    }
                    _ = fi65
                    _ = v66
                }
            }
            var v71 int32 = -2147483648
            _ = fi38
            _ = v71
        }
        var v72 int8 = 127
        var v73 uint64 = 18446744073709551615
        _ = v72
        _ = v73
    }
    for fi74 := 0; fi74 < 14; fi74++ {
        if true {
            v75 := GenS8{F0: "a", F1: (v4 == v4)}
            _ = v75
        }
        var v76 int32 = 1
        for fi77 := 0; fi77 < 17; fi77++ {
            v36 = v36
            v4 = v4
            _ = fi77
        }
        _ = fi74
        _ = v76
    }
    v4 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
