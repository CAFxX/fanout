package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 bool; F2 int8 }
type GenS4 struct { F0 int8; F1 string }
type GenS5 struct { F0 uint; F1 int8; F2 int64 }
type GenS6 struct { F0 uint32; F1 uint8 }
type GenS7 struct { F0 uint32 }

func main() {
    var v1 string = ("a" + "a")
    var v2 uint64 = 2
    var v3 uint16 = 827
    var v4 float64 = (1e10 / 1.5)
    v5 := []string{v1, "a", (((v1 + (v1 + "a")) + ("foo bar" + v1)) + "hello"), "a"}
    v6 := []uint32{922, 2147483647}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v7 bool = false
    for fi8 := 0; fi8 < 16; fi8++ {
        switch "foo bar" {
        case "hello":
            var v9 int16 = 1000
            _ = v9
        case "a":
            v2 = v2
        default:
            var v10 uint = 4294967295
            _ = v10
        }
        var v11 uint32 = 303
        _ = fi8
        _ = v11
    }
    if len(v6) > 0 {
        v6[0] = uint32(v3)
    }
    if len(v6) > 0 {
        v6[0] = 4294967295
    }
    v3 = v3
    if (!(!false)) {
        v3 = (v3 - v3)
        if len(v5) > 0 {
            v5[0] = "hello"
        }
        if len(v6) > 0 {
            v6[0] = 4294967295
        }
    }
    v7 = v7
    v1 = ""
    if (!v7) {
        var v12 uint16 = v3
        _ = v12
    }
    if (42 < 0) {
        switch -9223372036854775808 {
        case 32768:
            v13 := GenS1{F0: 0}
            _ = v13
        case 32767:
            if (((0.0 - v4) - v4) > ((0.1 * (v4 * v4)) * (((v4 - ((float64(3.14159) + v4) / v4)) / 1e10) / v4))) {
                v14 := GenS1{F0: 28}
                v15 := GenS1{F0: -58}
                var v16 string = "\n\t"
                _ = v14
                _ = v15
                _ = v16
            }
        }
        v3 = v3
        v17 := GenS1{F0: 9223372036854775807}
        _ = v17
    } else {
        if v7 {
            fmt.Println(v4, v4)
            if v7 {
                for fi18 := 0; fi18 < 16; fi18++ {
                    v7 = (v2 <= (^(v2)))
                    _ = fi18
                }
                v1 = v1
            } else {
                if (!(true || true)) {
                    var v19 bool = (v1 == v1)
                    for fi20 := 0; fi20 < 20; fi20++ {
                        if true {
                            v6 = append(v6, 65535)
                            if (3.14159 != 1e10) {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                if len(v6) > 0 {
                                    v6[0] = 0
                                }
                            } else {
                                if ((((373 > 764) && v7) && v19) || (255 <= 255)) {
                                    fmt.Println(v1, v3)
                                } else {
                                    v21 := GenS2{F0: v3}
                                    var v22 int64 = 42
                                    _ = v21
                                    _ = v22
                                }
                                v2 = v2
                            }
                        } else {
                            var v23 int64 = -9223372036854775808
                            _ = v23
                        }
                        v3 = (v3 & 42)
                        _ = fi20
                    }
                    v24 := GenS3{F0: 1, F1: (v3 >= v3), F2: 100}
                    _ = v19
                    _ = v24
                }
                for fi25 := 0; fi25 < 18; fi25++ {
                    v1 = v1
                    if v7 {
                        var v26 uint8 = 255
                        for fi27 := 0; fi27 < 18; fi27++ {
                            v28 := GenS1{F0: 487}
                            v3 = (v3 % ((652) | 1))
                            _ = fi27
                            _ = v28
                        }
                        _ = v26
                    } else {
                        for fi29 := 0; fi29 < 16; fi29++ {
                            v1 = v1
                            v2 = (^(v2))
                            var v30 uint64 = (v2 % (((v2 & (v2 % (((v2 << ((65535) & 63))) | 1)))) | 1))
                            _ = fi29
                            _ = v30
                        }
                        _ = len(v5)
                    }
                    switch "foo bar" {
                    case "\n\t":
                        if true {
                            for fi31 := 0; fi31 < 20; fi31++ {
                                switch fi31 {
                                case 127:
                                    v3 = v3
                                case 18:
                                    v4 += ((1.5 - ((v4 * v4) + (float64((-2.25 - 0.0)) / 1.0))) + ((0.0 / (3.14159 - 0.0)) - v4))
                                case 759:
                                    v32 := GenS3{F0: uint8(v3), F1: ((!(1 >= 255)) || (2 >= 527)), F2: 6}
                                    v3 = v3
                                    var v33 uint8 = 24
                                    if v7 {
                                        v34 := GenS4{F0: 100, F1: ("foo bar" + "hello")}
                                        var v35 uint16 = (^(v3))
                                        v3 -= (v3 >> ((18446744073709551615) & 63))
                                        _ = v34
                                        _ = v35
                                    } else {
                                        var v36 int = 9223372036854775807
                                        v5 = append(v5, "hello")
                                        _ = v36
                                    }
                                    _ = v32
                                    _ = v33
                                }
                                v2 = uint64(v3)
                                _ = fi31
                            }
                        }
                    case "a":
                        for fi37 := 0; fi37 < 4; fi37++ {
                            v4 = v4
                            var v38 bool = (v7 || (!false))
                            var v39 uint8 = 127
                            v1 = (v1 + "hello")
                            fmt.Println(v7)
                            _ = fi37
                            _ = v38
                            _ = v39
                        }
                    default:
                        if true {
                            for fi40 := 0; fi40 < 20; fi40++ {
                                v1 = (v1 + "hello")
                                _ = fi40
                            }
                            if (int16(v2) > 176) {
                                var v41 int32 = 0
                                v2 += 712
                                _ = v41
                            }
                        } else {
                            v42 := GenS2{F0: v3}
                            if v7 {
                                var v43 int8 = 127
                                if false {
                                    for fi44 := 0; fi44 < 4; fi44++ {
                                        v2 += (v2 / (((^(v2))) | 1))
                                        _ = fi44
                                    }
                                    v4 = float64(3.14159)
                                }
                                _ = v43
                            } else {
                                fmt.Println(v7)
                                v2 = v2
                                var v45 uint64 = uint64(v3)
                                _ = v45
                            }
                            _ = v42
                        }
                    }
                    _ = fi25
                }
            }
            var v46 string = v1
            _ = v46
        }
        v47 := GenS4{F0: 42, F1: "a"}
        var v48 int16 = int16(v3)
        _ = v47
        _ = v48
    }
    if v7 {
        if ((v3 | (^(v3))) < v3) {
            switch (("a" + v1) + v1) {
            case "hello":
                v49 := GenS5{F0: 865, F1: 0, F2: 42}
                _ = v49
            case "":
                var v50 int16 = 42
                _ = v50
            default:
                switch ((v1 + v1) + v1) {
                case "foo bar":
                    var v51 int = 32767
                    _ = v51
                case "\n\t":
                    for fi52 := 0; fi52 < 10; fi52++ {
                        for fi53 := 0; fi53 < 9; fi53++ {
                            switch -305 {
                            case 2147483647:
                                var v54 int16 = -32768
                                _ = v54
                            default:
                                v4 -= (v4 + v4)
                            }
                            switch "\n\t" {
                            case "":
                                switch (v1 + "") {
                                case "":
                                    switch (fi52 * 478) {
                                    case 9223372036854775807:
                                        if (-1 != 74) {
                                            v6 = append(v6, 2)
                                            fmt.Println(fi53)
                                            v4 -= 1e10
                                        } else {
                                            fmt.Println(v4, v3)
                                            var v55 float32 = 1e10
                                            var v56 uint16 = (v3 % (((v3 / (((v3 ^ 255)) | 1))) | 1))
                                            v57 := GenS6{F0: 65535, F1: 127}
                                            switch v1 {
                                            case "hello":
                                                if len(v6) > 0 {
                                                    v6[0] = 965
                                                }
                                            case "a":
                                                switch fi52 {
                                                case 996:
                                                    _ = len(v5)
                                                case 867:
                                                    v58 := GenS4{F0: 124, F1: (v1 + v1)}
                                                    _ = v58
                                                }
                                            }
                                            fmt.Println(v3, v4)
                                            _ = v55
                                            _ = v56
                                            _ = v57
                                        }
                                    case 114:
                                        v59 := GenS5{F0: 4294967295, F1: -1, F2: 1}
                                        _ = v59
                                    case 288:
                                        if v7 {
                                            var v60 uint64 = v2
                                            _ = v60
                                        }
                                    }
                                }
                            case "a":
                                fmt.Println(fi52, v1, fi53)
                            case "\n\t":
                                var v61 uint = 65535
                                _ = v61
                            default:
                                switch (v1 + v1) {
                                case "\n\t":
                                    v1 = ("hello" + "")
                                case "":
                                    v62 := GenS5{F0: 255, F1: 0, F2: 42}
                                    _ = v62
                                case "a":
                                    v63 := GenS2{F0: (v3 << ((255) & 63))}
                                    _ = v63
                                default:
                                    v2 = (v2 / ((v2) | 1))
                                }
                            }
                            _ = fi53
                        }
                        _ = fi52
                    }
                case "a":
                    switch (v1 + "a") {
                    case "\n\t":
                        fmt.Println(v3, v1, v4)
                    }
                }
            }
            _ = len(v5)
        }
    } else {
        fmt.Println(v2, v2)
        v1 = (v1 + (v1 + v1))
    }
    if v7 {
        v64 := GenS7{F0: 2}
        _ = v64
    }
    if (float32(v4) < 1e10) {
        var v65 int = 128
        if (128 != 128) {
            for fi66 := 0; fi66 < 9; fi66++ {
                var v67 uint16 = 65535
                if false {
                    var v68 string = "foo bar"
                    for fi69 := 0; fi69 < 2; fi69++ {
                        var v70 string = (v1 + "a")
                        if len(v5) > 0 {
                            v5[0] = v70
                        }
                        _ = fi69
                        _ = v70
                    }
                    _ = v68
                }
                _ = fi66
                _ = v67
            }
        } else {
            var v71 uint32 = 42
            _ = v71
        }
        var v72 uint16 = v3
        _ = v65
        _ = v72
    } else {
        fmt.Println(v7, v4, v2)
        if (1000 < -638) {
            v2 = v2
            v3 = v3
            _ = len(v5)
        }
    }
    if true {
        for fi73 := 0; fi73 < 20; fi73++ {
            var v74 bool = (1000 < 1)
            if ((0.1 - (float32(float64(-2.25)) / (float32((v4 * (0.0 + 1.5))) / 3.14159))) <= float32(float64(float32(v4)))) {
                fmt.Println(v3, v74)
                v3 += (^(v3))
                v75 := GenS6{F0: 4294967295, F1: 127}
                _ = v75
            }
            _ = fi73
            _ = v74
        }
        if (!false) {
            v5 = append(v5, v1)
            switch 18 {
            case -819:
                fmt.Println(v4, v4, v1)
            }
            var v76 int8 = -1
            _ = v76
        } else {
            if (true && ((!v7) && v7)) {
                var v77 int = -820
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                v5 = append(v5, v1)
                _ = v77
            }
            if v7 {
                v7 = true
                v2 += (v2 << ((509) & 63))
                v78 := GenS1{F0: 0}
                _ = v78
            } else {
                var v79 int8 = 75
                _ = v79
            }
        }
        var v80 uint32 = 42
        _ = v80
    }
    if v7 {
        v3 += v3
    } else {
        var v81 int8 = 1
        var v82 uint = 169
        _ = v81
        _ = v82
    }
    var v83 float64 = (1.5 / float64(-2.25))
    var v84 int32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v83
    _ = v84
}
