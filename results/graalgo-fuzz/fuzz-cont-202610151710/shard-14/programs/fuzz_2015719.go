package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 bool; F1 int; F2 uint16 }
type GenS3 struct { F0 int64; F1 uint8; F2 int32 }
type GenS4 struct { F0 int; F1 uint64 }
type GenS5 struct { F0 int8; F1 uint16; F2 int }
type GenS6 struct { F0 int16; F1 int8; F2 uint32 }

func main() {
    var v1 uint8 = 2
    v2 := make(map[string]uint32)
    var v3 uint64 = 1
    var v4 uint64 = v3
    var v5 uint8 = 42
    var v6 int16 = 1000
    v7 := GenS1{F0: 9223372036854775807, F1: 42}
    v3 -= (v4 << ((1) & 63))
    fmt.Println(v1, v3)
    for fi8 := 0; fi8 < 9; fi8++ {
        for fi9 := 0; fi9 < 9; fi9++ {
            if (!false) {
                if false {
                    var v10 uint8 = (v5 % ((v5) | 1))
                    v1 = v10
                    v1 = (v1 & v5)
                    if false {
                        v1 = v1
                        v4 -= (v3 >> ((65535) & 63))
                        v11 := GenS1{F0: fi8, F1: (fi8 - fi9)}
                        _ = v11
                    }
                    _ = v10
                }
                fmt.Println(v5, fi9)
                var v12 uint8 = 1
                _ = v12
            } else {
                v13 := GenS2{F0: false, F1: -537, F2: 65535}
                if false {
                    fmt.Println(v3, v3)
                    var v14 bool = false
                    _ = v14
                }
                var v15 int32 = -1
                _ = v13
                _ = v15
            }
            if true {
                var v16 int16 = (v6 | v6)
                var v17 uint32 = 1
                v6 = (v16 - v6)
                _ = v16
                _ = v17
            } else {
                var v18 float64 = (-2.25 / 3.14159)
                _ = v18
            }
            _ = fi9
        }
        _ = fi8
    }
    if ((1 == 9223372036854775807) && false) {
        for fi19 := 0; fi19 < 9; fi19++ {
            v3 -= v4
            _ = fi19
        }
        if true {
            v20 := GenS3{F0: -1, F1: v5, F2: 0}
            _ = v20
        }
        v4 -= (v3 * v3)
    }
    _ = len(v2)
    v5 += (v5 / ((v5) | 1))
    v21 := v2["hello"]
    _ = len(v2)
    if (1.5 >= 1.5) {
        v6 = v6
        if (v21 >= v21) {
            v4 += (v3 & (v3 % (((v4 / ((v3) | 1))) | 1)))
            v22 := GenS4{F0: -2147483648, F1: v4}
            var v23 int16 = v6
            _ = v22
            _ = v23
        }
    }
    if false {
        if false {
            var v24 int32 = int32(v6)
            for fi25 := 0; fi25 < 7; fi25++ {
                switch fi25 {
                case -2147483648:
                    if false {
                        var v26 string = "hello"
                        _ = v26
                    } else {
                        _ = len(v2)
                        v27 := GenS2{F0: ((fi25 % ((fi25) | 1)) != fi25), F1: fi25, F2: 255}
                        _ = v27
                    }
                default:
                    if (!(!(true && (!false)))) {
                        switch fi25 {
                        case 32767:
                            if (303 < (v21 * v21)) {
                                for fi28 := 0; fi28 < 14; fi28++ {
                                    v1 = v1
                                    _ = fi28
                                }
                                var v29 int64 = -1
                                _ = v29
                            } else {
                                v2["hello"] = v21
                                if true {
                                    v30 := GenS1{F0: fi25, F1: (fi25 ^ fi25)}
                                    _ = v30
                                }
                            }
                        default:
                            v31 := []int{fi25}
                            _ = v31
                        }
                        fmt.Println(fi25)
                        var v32 uint32 = v21
                        _ = v32
                    } else {
                        fmt.Println(v5, v4, v3)
                        fmt.Println(v1, v6)
                        var v33 int64 = 0
                        _ = v33
                    }
                }
                v6 = v6
                _ = fi25
            }
            var v34 uint64 = v3
            _ = v24
            _ = v34
        } else {
            for fi35 := 0; fi35 < 13; fi35++ {
                var v36 int = fi35
                v6 = v6
                switch "hello" {
                case "foo bar":
                    for fi37 := 0; fi37 < 15; fi37++ {
                        switch (v36 ^ fi35) {
                        case 127:
                            if false {
                                fmt.Println(fi37)
                                v2["foo bar"] = v21
                                v38 := []int{4294967295, fi35, (fi35 - 2147483647), 670}
                                _ = v38
                            } else {
                                v1 = (v5 & v1)
                                var v39 string = "hello"
                                _ = v39
                            }
                        case 86:
                            if false {
                                v1 = (v1 & v1)
                                switch "foo bar" {
                                case "hello":
                                    v40 := []int{32768, fi37}
                                    _ = v40
                                case "a":
                                    switch fi37 {
                                    case 32768:
                                        fmt.Println(v5, fi35)
                                    case -2147483648:
                                        v41 := []int{v36, int(v1), 42, int(v1)}
                                        _ = v41
                                    case -140:
                                        var v42 int16 = (v6 >> ((453) & 63))
                                        _ = v42
                                    default:
                                        v43 := []int{(fi37 - fi35), fi35, 2147483647, fi37}
                                        _ = v43
                                    }
                                case "\n\t":
                                    if true {
                                        if ((!((float32((0.1 / (0.1 / -2.25))) != -2.25) || ((322 < 988) || true))) && (!(fi35 != fi37))) {
                                            v5 = (^(v1))
                                        } else {
                                            var v44 uint64 = (v4 << ((255) & 63))
                                            _ = v44
                                        }
                                        fmt.Println(v21, v6)
                                        var v45 uint = 961
                                        _ = v45
                                    } else {
                                        var v46 uint32 = (v21 ^ v21)
                                        _ = v46
                                    }
                                }
                                fmt.Println(fi37)
                            }
                        }
                        fmt.Println(v3, fi37, v3)
                        v47 := []int{fi35, (-(fi37)), int(v6), -1}
                        _ = fi37
                        _ = v47
                    }
                case "hello":
                    v5 = v5
                case "a":
                    if false {
                        switch ("foo bar" + "") {
                        case "":
                            switch 127 {
                            case 373:
                                v3 -= (v3 & v4)
                            case 32768:
                                v48 := GenS2{F0: true, F1: (-(fi35)), F2: 0}
                                _ = v48
                            default:
                                v3 = (v3 / ((0) | 1))
                            }
                        case "hello":
                            fmt.Println(fi35, fi35, v36)
                        default:
                            switch "hello" {
                            case "":
                                var v49 int = v36
                                _ = v49
                            case "a":
                                var v50 float32 = 0.0
                                _ = v50
                            default:
                                v3 = 4294967295
                            }
                        }
                    }
                default:
                    _ = len(v2)
                }
                _ = fi35
                _ = v36
            }
            var v51 uint16 = 298
            _ = v51
        }
        for fi52 := 0; fi52 < 6; fi52++ {
            if true {
                v53 := []int{32768, (fi52 % ((2147483647) | 1)), fi52, fi52}
                var v54 string = "hello"
                _ = v53
                _ = v54
            }
            v6 += (^(v6))
            for fi55 := 0; fi55 < 19; fi55++ {
                fmt.Println(fi52, fi52, v5)
                fmt.Println(v5, fi52, v21)
                if false {
                    v6 = -32768
                }
                if (-9223372036854775808 <= 42) {
                    fmt.Println(v21)
                    if (((((0.1 * 3.14159) - float32(-2.25)) + 0.1) * -2.25) >= (1e10 * 1e10)) {
                        var v56 uint = 810
                        switch ("foo bar" + "a") {
                        case "hello":
                            v57 := []int{(^(fi55)), (fi52 | 958), (fi55 % (((fi52 | (fi55 | fi52))) | 1))}
                            _ = v57
                        case "a":
                            switch ("" + (("" + (((((("a" + "") + "") + "") + "foo bar") + "foo bar") + "")) + "\n\t")) {
                            case "":
                                var v58 uint16 = 65535
                                _ = v58
                            case "hello":
                                var v59 int32 = 2147483647
                                _ = v59
                            default:
                                var v60 int8 = 100
                                _ = v60
                            }
                        case "\n\t":
                            v61 := GenS5{F0: 42, F1: 264, F2: (fi52 / ((fi55) | 1))}
                            _ = v61
                        default:
                            var v62 uint32 = 737
                            _ = v62
                        }
                        var v63 uint16 = 255
                        _ = v56
                        _ = v63
                    }
                    v3 = 82
                }
                v5 = v5
                _ = fi55
            }
            _ = fi52
        }
    }
    switch "foo bar" {
    case "\n\t":
        if (!((1.5 == 3.14159) && (-1 > 116))) {
            switch "\n\t" {
            case "hello":
                switch "hello" {
                case "a":
                    for fi64 := 0; fi64 < 17; fi64++ {
                        var v65 uint16 = 778
                        _ = fi64
                        _ = v65
                    }
                case "foo bar":
                    for fi66 := 0; fi66 < 15; fi66++ {
                        var v67 bool = (!false)
                        v3 -= (v4 & 549)
                        var v68 bool = true
                        _ = fi66
                        _ = v67
                        _ = v68
                    }
                }
            case "\n\t":
                if false {
                    v69 := GenS2{F0: (!((255 <= 1) || true)), F1: -1, F2: 221}
                    v70 := GenS4{F0: 32768, F1: (^(v4))}
                    _ = v69
                    _ = v70
                }
            default:
                var v71 bool = (-64 == -172)
                _ = v71
            }
        }
    default:
        for fi72 := 0; fi72 < 20; fi72++ {
            v5 = 128
            _ = fi72
        }
    }
    v73 := GenS6{F0: v6, F1: -128, F2: v21}
    for fi74 := 0; fi74 < 12; fi74++ {
        v6 += (v6 / ((v6) | 1))
        v21 = (^(v21))
        _ = fi74
    }
    var v75 uint8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v73
    _ = v75
}
