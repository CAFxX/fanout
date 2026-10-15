package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 bool; F2 uint64 }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int32; F1 int32 }
type GenS7 struct { F0 bool; F1 string; F2 uint }

func main() {
    v1 := []uint32{886}
    var v2 uint16 = 448
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]int)
    v5 := []uint32{}
    var v6 int = -2147483648
    for fi7 := 0; fi7 < 11; fi7++ {
        for fi8 := 0; fi8 < 17; fi8++ {
            v9 := GenS1{F0: 32767}
            fmt.Println(v2, fi7)
            var v10 bool = false
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    if true {
        if (true || true) {
            if false {
                if false {
                    switch ("hello" + "foo bar") {
                    case "hello":
                        var v11 string = "hello"
                        _ = v11
                    case "foo bar":
                        var v12 int16 = -32768
                        _ = v12
                    case "\n\t":
                        v6 = (v6 & v6)
                    }
                    for fi13 := 0; fi13 < 5; fi13++ {
                        var v14 uint64 = 42
                        v14 = 1
                        _ = fi13
                        _ = v14
                    }
                    for fi15 := 0; fi15 < 17; fi15++ {
                        var v16 int32 = 264
                        fmt.Println(v6)
                        var v17 float32 = 0.1
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                } else {
                    v18 := GenS1{F0: 127}
                    switch (-(v6)) {
                    case -455:
                        switch "foo bar" {
                        case "foo bar":
                            v6 += v6
                        case "":
                            if true {
                                v19 := v4[0]
                                for fi20 := 0; fi20 < 10; fi20++ {
                                    var v21 bool = true
                                    var v22 bool = v21
                                    if false {
                                        if len(v5) > 0 {
                                            v5[0] = 2147483647
                                        }
                                    }
                                    _ = fi20
                                    _ = v21
                                    _ = v22
                                }
                                v23 := GenS1{F0: -1}
                                _ = v19
                                _ = v23
                            } else {
                                for fi24 := 0; fi24 < 17; fi24++ {
                                    var v25 string = "hello"
                                    switch fi24 {
                                    case -492:
                                        if (1 > 4294967295) {
                                            for fi26 := 0; fi26 < 12; fi26++ {
                                                var v27 int32 = -711
                                                v28 := GenS1{F0: -583}
                                                _ = fi26
                                                _ = v27
                                                _ = v28
                                            }
                                            if true {
                                                fmt.Println(v6, v6)
                                            } else {
                                                if len(v5) > 0 {
                                                    v5[0] = 671
                                                }
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                v29 := GenS1{F0: 42}
                                                _ = v29
                                            }
                                        } else {
                                            v4[0] = fi24
                                            v30 := GenS1{F0: 219}
                                            _ = v30
                                        }
                                    case 32768:
                                        if (false || (((float32(3.14159) - 1.5) / 1e10) != (1e10 * 0.1))) {
                                            if ((!true) && (true || false)) {
                                                switch ((v25 + (v25 + "\n\t")) + (((("foo bar" + v25) + "hello") + "") + (((v25 + v25) + "a") + (v25 + v25)))) {
                                                case "a":
                                                    var v31 string = v25
                                                    _ = v31
                                                case "hello":
                                                    for fi32 := 0; fi32 < 15; fi32++ {
                                                        switch (-(fi24)) {
                                                        case 0:
                                                            if true {
                                                                v6 += 32768
                                                                var v33 int64 = 9223372036854775807
                                                                switch (("hello" + "\n\t") + v25) {
                                                                case "a":
                                                                    var v34 int = fi24
                                                                    _ = v34
                                                                default:
                                                                    v6 -= (fi32 >> ((255) & 63))
                                                                    v25 = "\n\t"
                                                                    v35 := v3[v25]
                                                                    fmt.Println(v6, v25)
                                                                    _ = v35
                                                                }
                                                                _ = v33
                                                            }
                                                        case 32767:
                                                            if true {
                                                                v36 := GenS1{F0: -32768}
                                                                var v37 uint = 328
                                                                _ = v36
                                                                _ = v37
                                                            } else {
                                                                var v38 uint8 = 127
                                                                v39 := GenS1{F0: -32768}
                                                                _ = v38
                                                                _ = v39
                                                            }
                                                        case 4294967295:
                                                            if (!(995 < 1)) {
                                                                var v40 uint8 = 2
                                                                v2 += v2
                                                                _ = v40
                                                            } else {
                                                                if true {
                                                                    var v41 string = v25
                                                                    fmt.Println(v25)
                                                                    switch (-(fi24)) {
                                                                    case -2147483648:
                                                                        if ((115 >= 4294967295) && true) {
                                                                            v42 := GenS1{F0: 0}
                                                                            _ = len(v3)
                                                                            _ = v42
                                                                        }
                                                                    default:
                                                                        if false {
                                                                            v2 = (v2 * (v2 * (v2 / (((^(v2))) | 1))))
                                                                            v43 := GenS1{F0: 1000}
                                                                            _ = v43
                                                                        }
                                                                    }
                                                                    _ = v41
                                                                }
                                                            }
                                                        }
                                                        if true {
                                                            var v44 string = ("" + v25)
                                                            v6 = v6
                                                            v45 := GenS2{F0: (0.1 != 1e10), F1: (42 <= 2147483647), F2: 18446744073709551615}
                                                            _ = v44
                                                            _ = v45
                                                        } else {
                                                            var v46 float32 = (0.1 * 0.0)
                                                            _ = v46
                                                        }
                                                        fmt.Println(v25, v2)
                                                        _ = fi32
                                                    }
                                                case "foo bar":
                                                    fmt.Println(fi24, fi24)
                                                }
                                            } else {
                                                v2 = (v2 << ((65535) & 63))
                                                var v47 int64 = 1
                                                v25 = "a"
                                                _ = v47
                                            }
                                        }
                                    case 239:
                                        v48 := GenS3{F0: fi24, F1: -1}
                                        _ = v48
                                    default:
                                        v49 := GenS1{F0: 605}
                                        _ = v49
                                    }
                                    if (true && false) {
                                        switch "\n\t" {
                                        case "foo bar":
                                            v50 := GenS2{F0: false, F1: (1e10 < 0.0), F2: 450}
                                            _ = v50
                                        case "":
                                            switch v25 {
                                            case "hello":
                                                var v51 bool = true
                                                _ = v51
                                            case "foo bar":
                                                delete(v3, v25)
                                            }
                                        case "\n\t":
                                            v52 := GenS4{F0: 42}
                                            _ = v52
                                        }
                                        if (0 != 2) {
                                            v2 = 927
                                            if (!((1.5 + (-2.25 - float64((3.14159 / (0.0 + (3.14159 + (0.0 * 3.14159))))))) == 0.0)) {
                                                var v53 int = int(v2)
                                                _ = v53
                                            }
                                        } else {
                                            for fi54 := 0; fi54 < 4; fi54++ {
                                                v55 := GenS4{F0: -540}
                                                _ = fi54
                                                _ = v55
                                            }
                                        }
                                        if (1.5 >= (1.5 - 1.5)) {
                                            var v56 uint64 = 2
                                            _ = v56
                                        }
                                    }
                                    _ = fi24
                                    _ = v25
                                }
                                var v57 int = 304
                                _ = v57
                            }
                        default:
                            var v58 int32 = 1
                            _ = v58
                        }
                    case 42:
                        v59 := GenS2{F0: true, F1: (!true), F2: 42}
                        _ = v59
                    }
                    _ = v18
                }
                if (!true) {
                    for fi60 := 0; fi60 < 12; fi60++ {
                        if true {
                            for fi61 := 0; fi61 < 20; fi61++ {
                                v2 -= 255
                                if true {
                                    v62 := GenS1{F0: 32767}
                                    v6 = (fi60 % ((128) | 1))
                                    v6 = -759
                                    v2 = (^(v2))
                                    _ = v62
                                }
                                _ = fi61
                            }
                            var v63 float64 = ((1e10 / 1e10) - 0.0)
                            v64 := GenS5{F0: int(v2)}
                            _ = v63
                            _ = v64
                        }
                        _ = fi60
                    }
                    var v65 int16 = 1000
                    _ = v65
                } else {
                    var v66 uint32 = 942
                    _ = v66
                }
                var v67 bool = (1 <= 100)
                _ = v67
            }
        }
    }
    var v68 int32 = 840
    v69 := GenS5{F0: (v6 ^ (v6 & (v6 ^ (v6 * -1))))}
    fmt.Println(v68)
    for fi70 := 0; fi70 < 12; fi70++ {
        switch "" {
        case "foo bar":
            var v71 int8 = -1
            _ = v71
        case "\n\t":
            for fi72 := 0; fi72 < 1; fi72++ {
                delete(v4, 2147483647)
                v73 := GenS1{F0: -32768}
                _ = fi72
                _ = v73
            }
        }
        _ = fi70
    }
    var v74 int64 = -1
    var v75 bool = true
    for fi76 := 0; fi76 < 7; fi76++ {
        v74 += (v74 - v74)
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi76
    }
    switch -9223372036854775808 {
    case 9223372036854775807:
        v2 += (v2 | v2)
    case -2147483648:
        switch -2147483648 {
        case 2147483647:
            switch (v6 ^ v6) {
            case -2147483648:
                v77 := GenS3{F0: v6, F1: (^(v68))}
                _ = v77
            case 1:
                var v78 uint32 = 42
                _ = v78
            default:
                v79 := GenS4{F0: (v74 + v74)}
                _ = v79
            }
        case -766:
            if v75 {
                v68 = 42
                v68 -= v68
            }
        default:
            if v75 {
                if false {
                    if v75 {
                        delete(v3, ("foo bar" + "a"))
                        var v80 uint = 515
                        _ = v80
                    } else {
                        v75 = false
                        fmt.Println(v74)
                        if false {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                    }
                } else {
                    var v81 int = (v6 - v6)
                    v82 := GenS1{F0: int16(v74)}
                    _ = v81
                    _ = v82
                }
                v83 := v4[18]
                _ = v83
            } else {
                switch "\n\t" {
                case "a":
                    var v84 int64 = int64(v6)
                    _ = v84
                }
                switch (v6 % ((v6) | 1)) {
                case 524:
                    v6 += v6
                case -1:
                    v68 = int32(v6)
                case 42:
                    for fi85 := 0; fi85 < 11; fi85++ {
                        v86 := GenS1{F0: 42}
                        _ = fi85
                        _ = v86
                    }
                default:
                    var v87 uint16 = v2
                    _ = v87
                }
                var v88 int = v6
                _ = v88
            }
        }
    case -1:
        var v89 bool = (!(42 >= 4294967295))
        _ = v89
    }
    var v90 int = (v6 - v6)
    var v91 int8 = -30
    switch "a" {
    case "a":
        var v92 int = (v6 << ((232) & 63))
        _ = v92
    case "foo bar":
        if (true && ((((1.5 * -2.25) - 0.1) < -2.25) && v75)) {
            var v93 int64 = (v74 % ((v74) | 1))
            _ = v93
        } else {
            v4[445] = 42
            v3["hello"] = 1
            v2 = v2
        }
    case "":
        v94 := GenS1{F0: -1}
        _ = v94
    default:
        var v95 float32 = (float32((0.0 + float64(-2.25))) + 0.0)
        _ = v95
    }
    v96 := GenS6{F0: v68, F1: int32(v74)}
    v97 := GenS7{F0: v75, F1: "hello", F2: 65535}
    var v98 string = ("hello" + ("foo bar" + ("\n\t" + "foo bar")))
    for fi99 := 0; fi99 < 18; fi99++ {
        var v100 int = v6
        v68 = v68
        _ = fi99
        _ = v100
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v68:", v68)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    fmt.Println("v90:", v90)
    fmt.Println("v91:", v91)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v68
    _ = v69
    _ = v74
    _ = v75
    _ = v90
    _ = v91
    _ = v96
    _ = v97
    _ = v98
}
