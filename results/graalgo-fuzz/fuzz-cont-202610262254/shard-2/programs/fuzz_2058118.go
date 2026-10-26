package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint }
type GenS2 struct { F0 int32; F1 uint32; F2 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int; F1 int64 }
type GenS5 struct { F0 uint8; F1 uint32 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int32; F1 uint32 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint32 }

func main() {
    var v1 uint16 = 65535
    var v2 uint64 = 0
    v3 := []uint32{4294967295, 2, 42}
    var v4 int64 = 42
    v4 += v4
    if (0 == 125) {
        v5 := GenS1{F0: "foo bar", F1: v2, F2: 2}
        v2 -= v2
        v2 = v2
        _ = v5
    } else {
        v4 -= -670
    }
    v3 = append(v3, 65535)
    var v6 string = ("foo bar" + "")
    var v7 float32 = -2.25
    v3 = append(v3, 4294967295)
    v4 = (^(v4))
    fmt.Println(v1, v1)
    if false {
        fmt.Println(v7)
        fmt.Println(v7, v4, v4)
    } else {
        fmt.Println(v7)
        for fi8 := 0; fi8 < 12; fi8++ {
            switch fi8 {
            case -22:
                for fi9 := 0; fi9 < 16; fi9++ {
                    fmt.Println(v4, v1, fi8)
                    _ = fi9
                }
            default:
                fmt.Println(fi8, v4)
            }
            v1 = 215
            if len(v3) > 0 {
                v3[0] = 2
            }
            _ = fi8
        }
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        switch "hello" {
        case "a":
            v11 := GenS1{F0: v6, F1: v2, F2: 4294967295}
            _ = v11
        default:
            if false {
                switch 127 {
                case 4294967295:
                    v1 = (^(v1))
                case 32768:
                    if (((v4 + 0) == (v4 / (((v4 % ((v4) | 1))) | 1))) && false) {
                        v3 = append(v3, 2)
                        for fi12 := 0; fi12 < 15; fi12++ {
                            var v13 int16 = -777
                            switch fi10 {
                            case 1:
                                fmt.Println(v6, fi10)
                            }
                            v14 := GenS1{F0: v6, F1: (^(v2)), F2: 255}
                            switch v6 {
                            case "a":
                                switch "hello" {
                                case "\n\t":
                                    var v15 bool = ((v13 - v13) > v13)
                                    _ = v15
                                case "a":
                                    var v16 int64 = v4
                                    _ = v16
                                case "":
                                    switch fi10 {
                                    case 804:
                                        if (!false) {
                                            v7 += ((0.1 / v7) + (v7 + v7))
                                            v3 = append(v3, 0)
                                        } else {
                                            var v17 string = v6
                                            _ = v17
                                        }
                                    case -488:
                                        var v18 uint32 = 130
                                        _ = v18
                                    case 32767:
                                        if len(v3) > 0 {
                                            v3[0] = 2147483648
                                        }
                                    }
                                default:
                                    fmt.Println(v13, fi10, v4)
                                }
                                v2 -= 706
                            default:
                                if len(v3) > 0 {
                                    v3[0] = uint32(v2)
                                }
                            }
                            if (v7 < float32((3.14159 + -2.25))) {
                                var v19 int16 = -1
                                _ = v19
                            }
                            _ = fi12
                            _ = v13
                            _ = v14
                        }
                        if (42 != 0) {
                            v20 := GenS2{F0: 489, F1: 282, F2: "a"}
                            _ = v20
                        } else {
                            fmt.Println(v2)
                            if ((!(true || (!false))) && ((42 != 2147483647) && true)) {
                                if (v7 != (0.1 + (-2.25 + -2.25))) {
                                    var v21 int = fi10
                                    _ = v21
                                } else {
                                    if false {
                                        var v22 float64 = -2.25
                                        _ = v22
                                    }
                                    v2 = 0
                                }
                            } else {
                                var v23 uint8 = 42
                                if false {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    v6 = (v6 + v6)
                                }
                                if len(v3) > 0 {
                                    v3[0] = 377
                                }
                                _ = v23
                            }
                        }
                    }
                case 0:
                    var v24 uint8 = 2
                    _ = v24
                }
                if false {
                    if true {
                        for fi25 := 0; fi25 < 6; fi25++ {
                            if false {
                                if len(v3) > 0 {
                                    v3[0] = 2
                                }
                                var v26 float64 = float64(v7)
                                _ = v26
                            } else {
                                if ((((3.14159 + -2.25) + 1.5) + (3.14159 + 3.14159)) > 1.5) {
                                    var v27 bool = false
                                    _ = v27
                                }
                                fmt.Println(v4, v6, v2)
                                var v28 int8 = 100
                                _ = v28
                            }
                            _ = fi25
                        }
                        if (!true) {
                            if true {
                                v4 = v4
                                var v29 bool = false
                                v7 = v7
                                _ = v29
                            }
                            v6 = ""
                        } else {
                            if false {
                                v1 -= (v1 << ((798) & 63))
                                fmt.Println(v6, v4, fi10)
                                v30 := GenS3{F0: 664}
                                _ = v30
                            }
                        }
                        v1 -= v1
                    }
                }
                if len(v3) > 0 {
                    v3[0] = 2147483647
                }
            } else {
                if (!(1 <= v4)) {
                    if true {
                        v6 = v6
                        v1 -= v1
                        v7 += (v7 - float32((float64(v7) - float64(((v7 / float32(-2.25)) / 1.5)))))
                    } else {
                        switch (-(fi10)) {
                        case 32767:
                            var v31 uint64 = v2
                            _ = v31
                        case 32768:
                            v32 := GenS3{F0: 42}
                            _ = v32
                        }
                        for fi33 := 0; fi33 < 14; fi33++ {
                            if (fi10 == (^(fi33))) {
                                v1 = (v1 & 2)
                                fmt.Println(v2, v2, v2)
                            }
                            v34 := GenS4{F0: (fi10 ^ (fi33 - 9223372036854775807)), F1: (v4 ^ (v4 - (v4 % ((356) | 1))))}
                            _ = fi33
                            _ = v34
                        }
                    }
                    fmt.Println(v2)
                }
                var v35 int64 = -1
                _ = v35
            }
        }
        _ = fi10
    }
    for fi36 := 0; fi36 < 1; fi36++ {
        if false {
            for fi37 := 0; fi37 < 10; fi37++ {
                if (!true) {
                    var v38 uint8 = 255
                    _ = v38
                }
                if true {
                    var v39 bool = (fi37 > fi37)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v40 uint64 = v2
                    _ = v39
                    _ = v40
                } else {
                    v41 := GenS5{F0: 128, F1: 42}
                    v1 += (v1 + (v1 & v1))
                    _ = v41
                }
                _ = fi37
            }
            switch "a" {
            case "hello":
                v1 -= v1
            case "\n\t":
                v4 += v4
            default:
                v42 := GenS1{F0: v6, F1: v2, F2: 693}
                _ = v42
            }
            for fi43 := 0; fi43 < 16; fi43++ {
                if len(v3) > 0 {
                    v3[0] = 2147483648
                }
                _ = fi43
            }
        } else {
            var v44 uint16 = v1
            var v45 int8 = -1
            _ = v44
            _ = v45
        }
        if len(v3) > 0 {
            v3[0] = 42
        }
        v1 = (v1 / ((v1) | 1))
        _ = fi36
    }
    _ = len(v3)
    var v46 int8 = 42
    switch v6 {
    case "\n\t":
        _ = len(v3)
    case "foo bar":
        if true {
            var v47 int64 = (v4 - (v4 + v4))
            _ = v47
        } else {
            v1 = (v1 ^ (v1 >> ((230) & 63)))
        }
    default:
        v48 := GenS3{F0: 584}
        _ = v48
    }
    v1 = v1
    if ((v46 * (v46 & v46)) < v46) {
        if (!true) {
            if (!true) {
                for fi49 := 0; fi49 < 19; fi49++ {
                    v50 := GenS6{F0: 636}
                    var v51 string = ""
                    var v52 int64 = (v4 ^ v4)
                    _ = fi49
                    _ = v50
                    _ = v51
                    _ = v52
                }
                v4 = v4
                switch v6 {
                case "foo bar":
                    v3 = append(v3, 297)
                default:
                    v2 = v2
                }
            }
            v1 = 2
            if ((2147483647 != 501) || true) {
                var v53 string = v6
                _ = v53
            }
        }
        for fi54 := 0; fi54 < 2; fi54++ {
            v55 := GenS7{F0: int32(v1), F1: 1}
            if len(v3) > 0 {
                v3[0] = 489
            }
            for fi56 := 0; fi56 < 20; fi56++ {
                var v57 int = 2147483647
                fmt.Println(v57, v57, v57)
                var v58 float64 = float64((0.1 + v7))
                fmt.Println(v2, v6)
                _ = fi56
                _ = v57
                _ = v58
            }
            _ = fi54
            _ = v55
        }
        var v59 int = -2147483648
        _ = v59
    } else {
        if false {
            var v60 string = "foo bar"
            var v61 int64 = v4
            _ = v60
            _ = v61
        } else {
            switch (v6 + v6) {
            case "\n\t":
                for fi62 := 0; fi62 < 6; fi62++ {
                    if ((644 <= 442) && true) {
                        v63 := GenS8{F0: v4}
                        _ = len(v3)
                        for fi64 := 0; fi64 < 9; fi64++ {
                            var v65 bool = true
                            var v66 float32 = v7
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi64
                            _ = v65
                            _ = v66
                        }
                        _ = v63
                    } else {
                        if len(v3) > 0 {
                            v3[0] = 201
                        }
                        switch (v6 + "\n\t") {
                        case "hello":
                            for fi67 := 0; fi67 < 1; fi67++ {
                                if true {
                                    switch ((v6 + "") + v6) {
                                    case "hello":
                                        v7 -= (1.5 - float32(3.14159))
                                    case "\n\t":
                                        if true {
                                            fmt.Println(v46, v46, v2)
                                        }
                                    default:
                                        v68 := GenS9{F0: 399}
                                        _ = v68
                                    }
                                }
                                var v69 int16 = -1
                                v2 -= v2
                                _ = fi67
                                _ = v69
                            }
                        case "foo bar":
                            fmt.Println(v6, fi62, fi62)
                        default:
                            v1 -= (v1 >> ((uint(v1)) & 63))
                        }
                    }
                    var v70 bool = (-2147483648 > 42)
                    fmt.Println(v7, fi62)
                    _ = fi62
                    _ = v70
                }
            }
            var v71 float64 = 0.1
            _ = v71
        }
    }
    v1 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v6
    _ = v7
}
