package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 bool; F2 uint }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int8; F1 int16; F2 uint8 }
type GenS7 struct { F0 int16; F1 uint8 }
type GenS8 struct { F0 uint; F1 int32 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 int32 }

func main() {
    v1 := make(map[int]bool)
    v2 := []string{"a", "hello", ("" + "")}
    var v3 int8 = 100
    v4 := make(map[string]uint32)
    var v5 uint16 = 68
    var v6 float64 = 1.5
    v7 := GenS1{F0: true}
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 uint32 = 2147483647
        v10 := GenS1{F0: true}
        _ = fi8
        _ = v10
        _ = v9
    }
    v5 = v5
    v11 := GenS2{F0: 1000, F1: (-595 < 0), F2: 0}
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        fmt.Println(v6)
        if (!false) {
            var v13 int = fi12
            _ = v13
        } else {
            var v14 float64 = v6
            v6 = (v6 / ((3.14159 / 1.5) + (float64(1.5) * (((float64(0.1) / (float64(1e10) / -2.25)) / (v6 * v14)) + ((((v6 / (v6 * v6)) * (float64(float32(float64(3.14159))) - float64(0.1))) - v6) - 1e10)))))
            _ = v14
        }
        if (3.14159 < float64(float32(v6))) {
            v6 += 0.0
        } else {
            var v15 int = 9223372036854775807
            var v16 uint32 = 0
            switch 128 {
            case 128:
                var v17 float32 = (1e10 + -2.25)
                _ = v17
            case 2147483647:
                var v18 uint8 = uint8(v16)
                _ = v18
            case 32767:
                var v19 int16 = 603
                _ = v19
            default:
                _ = len(v1)
            }
            _ = v15
            _ = v16
        }
        _ = fi12
    }
    for fi20 := 0; fi20 < 2; fi20++ {
        for fi21 := 0; fi21 < 6; fi21++ {
            v22 := GenS3{F0: 2147483647, F1: 71}
            var v23 uint64 = 0
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = fi20
    }
    if true {
        switch "" {
        case "hello":
            for fi24 := 0; fi24 < 13; fi24++ {
                v25 := GenS1{F0: true}
                fmt.Println(fi24, v6)
                _ = fi24
                _ = v25
            }
        case "a":
            fmt.Println(v5)
        case "foo bar":
            if (!(!true)) {
                if (42 < 677) {
                    switch ("" + "foo bar") {
                    case "foo bar":
                        var v26 bool = true
                        _ = v26
                    case "a":
                        if (969 <= 1) {
                            fmt.Println(v6, v3)
                            if len(v2) > 0 {
                                v2[0] = "hello"
                            }
                            for fi27 := 0; fi27 < 6; fi27++ {
                                var v28 int32 = 809
                                var v29 uint16 = 32767
                                v29 += v5
                                _ = fi27
                                _ = v28
                                _ = v29
                            }
                        }
                    case "hello":
                        v6 = v6
                    default:
                        for fi30 := 0; fi30 < 10; fi30++ {
                            _ = len(v1)
                            v31 := GenS3{F0: 0, F1: 127}
                            if true {
                                if true {
                                    fmt.Println(v5)
                                    v2 = append(v2, "foo bar")
                                } else {
                                    if false {
                                        v32 := GenS3{F0: -2147483648, F1: 2}
                                        v3 = (v3 | v3)
                                        v33 := v1[940]
                                        _ = v32
                                        _ = v33
                                    }
                                    switch "a" {
                                    case "a":
                                        v2 = append(v2, "")
                                    case "foo bar":
                                        v5 = (^(v5))
                                    }
                                    var v34 uint32 = 0
                                    _ = v34
                                }
                                v35 := GenS3{F0: -1, F1: 128}
                                _ = v35
                            }
                            _ = fi30
                            _ = v31
                        }
                    }
                    v5 = (v5 / ((v5) | 1))
                    fmt.Println(v6, v5, v5)
                }
                if (320 <= 2147483647) {
                    v36 := v4["foo bar"]
                    _ = v36
                }
            } else {
                v37 := GenS4{F0: 255}
                v38 := GenS3{F0: -736, F1: 127}
                var v39 float32 = 1.5
                _ = v37
                _ = v38
                _ = v39
            }
        default:
            v3 += (v3 >> ((255) & 63))
        }
        for fi40 := 0; fi40 < 3; fi40++ {
            switch (fi40 - 284) {
            case 32768:
                for fi41 := 0; fi41 < 20; fi41++ {
                    var v42 int16 = 42
                    v43 := v1[fi40]
                    if (!v43) {
                        if v43 {
                            var v44 float32 = 3.14159
                            fmt.Println(v6, v42, v42)
                            var v45 uint64 = 0
                            fmt.Println(fi40, v44, v43)
                            _ = len(v1)
                            if ("hello" > (("a" + ((("foo bar" + ("hello" + "")) + ("\n\t" + "")) + ("a" + "\n\t"))) + "")) {
                                switch "\n\t" {
                                case "a":
                                    if (!(!((^(v42)) != v42))) {
                                        v42 = (v42 | 42)
                                    }
                                case "\n\t":
                                    v4["a"] = 2147483648
                                case "hello":
                                    v3 -= (v3 / ((v3) | 1))
                                }
                            } else {
                                var v46 bool = v43
                                v42 -= v42
                                _ = v46
                            }
                            v45 -= (v45 >> ((4294967295) & 63))
                            switch (fi41 ^ fi40) {
                            case 127:
                                fmt.Println(v5, v3)
                            case 4294967295:
                                var v47 int16 = (v42 * (v42 % ((v42) | 1)))
                                _ = v47
                            default:
                                switch fi41 {
                                case 128:
                                    if ((v42 >> ((18446744073709551615) & 63)) != -823) {
                                        if v43 {
                                            v2 = append(v2, ("\n\t" + ("\n\t" + "foo bar")))
                                            var v48 float64 = v6
                                            _ = v48
                                        } else {
                                            v3 = v3
                                        }
                                        fmt.Println(v5, v45)
                                        switch "\n\t" {
                                        case "":
                                            var v49 uint = 255
                                            _ = v49
                                        }
                                    }
                                case 127:
                                    if v43 {
                                        v45 = (v45 | 0)
                                        var v50 int8 = v3
                                        _ = v50
                                    }
                                }
                            }
                            _ = v44
                            _ = v45
                        } else {
                            if (!v43) {
                                var v51 int8 = -1
                                _ = v51
                            } else {
                                v52 := GenS4{F0: 1}
                                _ = v52
                            }
                            v5 = v5
                        }
                        if (v43 || v43) {
                            var v53 uint32 = 2147483647
                            _ = v53
                        } else {
                            v54 := GenS5{F0: (v3 & v3)}
                            var v55 string = "foo bar"
                            _ = v54
                            _ = v55
                        }
                        v56 := GenS6{F0: int8(fi40), F1: 32767, F2: 128}
                        _ = v56
                    }
                    _ = fi41
                    _ = v42
                    _ = v43
                }
            case 313:
                v6 = 1e10
            case 2147483647:
                for fi57 := 0; fi57 < 9; fi57++ {
                    if ((false || true) || ((false || false) && false)) {
                        fmt.Println(v6)
                    }
                    _ = fi57
                }
            default:
                if (!true) {
                    v3 = v3
                }
            }
            v58 := GenS2{F0: 1000, F1: true, F2: 42}
            switch (fi40 / ((fi40) | 1)) {
            case 32767:
                v59 := GenS4{F0: 18446744073709551615}
                _ = v59
            case -331:
                v60 := GenS5{F0: 1}
                _ = v60
            case -24:
                for fi61 := 0; fi61 < 20; fi61++ {
                    if len(v2) > 0 {
                        v2[0] = ((("" + ("foo bar" + "foo bar")) + "foo bar") + "")
                    }
                    var v62 uint64 = 596
                    if (v5 <= (v5 & 184)) {
                        fmt.Println(fi40, v3)
                    } else {
                        var v63 bool = ((42 > 1) && (1 >= (fi40 % ((32768) | 1))))
                        _ = v63
                    }
                    if (4294967295 >= 840) {
                        v62 -= (v62 >> ((2) & 63))
                        var v64 uint8 = 0
                        delete(v4, ("hello" + "\n\t"))
                        v65 := GenS7{F0: 42, F1: (v64 / ((v64) | 1))}
                        v6 = v6
                        _ = v64
                        _ = v65
                    }
                    _ = fi61
                    _ = v62
                }
            default:
                for fi66 := 0; fi66 < 5; fi66++ {
                    var v67 bool = false
                    _ = fi66
                    _ = v67
                }
            }
            _ = fi40
            _ = v58
        }
    } else {
        var v68 uint16 = (^(v5))
        var v69 int64 = 0
        _ = v68
        _ = v69
    }
    var v70 int16 = -1
    var v71 uint16 = v5
    if false {
        var v72 float64 = v6
        fmt.Println(v71, v5, v72)
        v3 = v3
        _ = v72
    } else {
        v73 := GenS6{F0: (v3 ^ (-(v3))), F1: 42, F2: uint8(v71)}
        v74 := GenS8{F0: 550, F1: 841}
        if false {
            switch 42 {
            case 518:
                v6 = 1e10
            default:
                var v75 uint8 = 42
                _ = v75
            }
            var v76 int16 = 0
            _ = v76
        } else {
            switch ("a" + ((("hello" + ("foo bar" + "")) + "foo bar") + "a")) {
            case "":
                v5 = (v5 | (v5 * v5))
            case "foo bar":
                if (42 <= int64(v71)) {
                    v70 = (-(v70))
                } else {
                    var v77 float64 = float64((float32(v6) + (1.5 + 1e10)))
                    for fi78 := 0; fi78 < 4; fi78++ {
                        for fi79 := 0; fi79 < 8; fi79++ {
                            v77 += (float64(3.14159) + float64(-2.25))
                            var v80 uint32 = 2
                            var v81 float64 = v77
                            var v82 int32 = 0
                            v83 := GenS9{F0: 1}
                            _ = fi79
                            _ = v80
                            _ = v81
                            _ = v82
                            _ = v83
                        }
                        if true {
                            for fi84 := 0; fi84 < 1; fi84++ {
                                if true {
                                    fmt.Println(v3, v5)
                                }
                                var v85 int8 = v3
                                _ = len(v4)
                                _ = fi84
                                _ = v85
                            }
                        } else {
                            if false {
                                var v86 int64 = 1
                                _ = v86
                            } else {
                                if false {
                                    var v87 string = (("a" + "foo bar") + "")
                                    if (9223372036854775807 <= 1) {
                                        for fi88 := 0; fi88 < 7; fi88++ {
                                            var v89 int64 = 893
                                            _ = fi88
                                            _ = v89
                                        }
                                        switch v87 {
                                        case "foo bar":
                                            fmt.Println(v70, v3)
                                        case "hello":
                                            v90 := GenS10{F0: v70}
                                            _ = v90
                                        }
                                    }
                                    fmt.Println(v3, v3, v70)
                                    _ = v87
                                } else {
                                    var v91 uint32 = 0
                                    _ = len(v2)
                                    _ = v91
                                }
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v77 = 0.1
                            }
                        }
                        _ = fi78
                    }
                    _ = v77
                }
            case "hello":
                if false {
                    if false {
                        switch "\n\t" {
                        case "a":
                            v3 = v3
                        case "foo bar":
                            v5 -= (v5 - (v5 * (^(v5))))
                        default:
                            var v92 uint32 = 427
                            _ = v92
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    } else {
                        fmt.Println(v71, v6)
                        var v93 int32 = -2147483648
                        switch 128 {
                        case 47:
                            v94 := GenS11{F0: (v70 ^ (v70 + (-(v70))))}
                            _ = v94
                        case 32767:
                            var v95 string = ("foo bar" + "hello")
                            _ = v95
                        default:
                            v96 := GenS12{F0: (v93 | 389)}
                            _ = v96
                        }
                        _ = v93
                    }
                } else {
                    delete(v4, "a")
                    var v97 uint = 782
                    _ = v97
                }
            default:
                v71 -= 4
            }
            v71 -= v71
            v3 = -1
        }
        _ = v73
        _ = v74
    }
    for fi98 := 0; fi98 < 2; fi98++ {
        v70 = -32768
        _ = fi98
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v70
    _ = v71
}
