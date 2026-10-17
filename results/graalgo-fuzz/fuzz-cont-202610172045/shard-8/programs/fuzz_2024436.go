package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 bool; F1 bool; F2 uint8 }
type GenS5 struct { F0 string; F1 int8; F2 uint }
type GenS6 struct { F0 uint8; F1 uint; F2 int32 }
type GenS7 struct { F0 uint; F1 uint8; F2 uint16 }
type GenS8 struct { F0 uint64 }

func main() {
    var v1 int = 127
    var v2 int8 = 0
    var v3 bool = false
    var v4 string = "\n\t"
    v5 := make(map[int]int)
    fmt.Println(v2, v2, v2)
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 uint16 = 825
        _ = fi6
        _ = v7
    }
    v4 = ("a" + v4)
    for fi8 := 0; fi8 < 16; fi8++ {
        v3 = v3
        _ = fi8
    }
    if false {
        fmt.Println(v1)
        switch v4 {
        case "\n\t":
            for fi9 := 0; fi9 < 2; fi9++ {
                var v10 string = (((v4 + "\n\t") + v4) + v4)
                switch "hello" {
                case "":
                    fmt.Println(v10, fi9, v10)
                case "hello":
                    if (((18446744073709551615 <= 4294967295) || v3) || (!(!(1 > 42)))) {
                        v3 = true
                        v2 -= (-(v2))
                    } else {
                        if v3 {
                            v10 = ((v10 + "") + (v10 + v10))
                            for fi11 := 0; fi11 < 11; fi11++ {
                                if v3 {
                                    var v12 string = (v4 + v10)
                                    v13 := GenS1{F0: 311}
                                    v14 := GenS2{F0: v10, F1: 116}
                                    v15 := GenS2{F0: "foo bar", F1: 181}
                                    var v16 uint16 = 32767
                                    _ = v12
                                    _ = v13
                                    _ = v14
                                    _ = v15
                                    _ = v16
                                } else {
                                    if (float32((-2.25 + 1e10)) <= -2.25) {
                                        v17 := []int{-1, (^(v1)), (fi9 / ((814) | 1))}
                                        v4 = v10
                                        v18 := GenS3{F0: 212, F1: 0}
                                        v1 -= (^(fi11))
                                        if v3 {
                                            var v19 uint64 = 42
                                            fmt.Println(v1)
                                            _ = v19
                                        }
                                        _ = v17
                                        _ = v18
                                    }
                                    switch 9223372036854775807 {
                                    case 0:
                                        v20 := v5[(^(v1))]
                                        _ = v20
                                    case 95:
                                        v21 := GenS2{F0: "\n\t", F1: 1000}
                                        var v22 float32 = (0.0 + 0.0)
                                        if (false && v3) {
                                            switch fi9 {
                                            case 160:
                                                fmt.Println(v1, fi11, v3)
                                            case 2147483647:
                                                var v23 uint = 1
                                                _ = v23
                                            }
                                            if (0 > 1) {
                                                v5[(fi11 % (((v1 ^ (fi11 & (fi11 / ((9223372036854775807) | 1))))) | 1))] = fi9
                                                switch (-(fi9)) {
                                                case 985:
                                                    v24 := GenS3{F0: v1, F1: 42}
                                                    _ = v24
                                                case 9223372036854775807:
                                                    fmt.Println(v10, v3, v22)
                                                }
                                            } else {
                                                v22 = v22
                                                v2 = v2
                                                var v25 uint32 = 223
                                                _ = v25
                                            }
                                        } else {
                                            _ = len(v5)
                                            v26 := GenS4{F0: (((3.14159 / (1e10 / (1e10 + (1.5 * -2.25)))) / -2.25) == 3.14159), F1: true, F2: 42}
                                            switch 585 {
                                            case 9223372036854775807:
                                                v22 -= 1e10
                                            default:
                                                var v27 int16 = 1
                                                _ = v27
                                            }
                                            _ = v26
                                        }
                                        _ = v21
                                        _ = v22
                                    default:
                                        v28 := GenS2{F0: "a", F1: 535}
                                        _ = v28
                                    }
                                    if v3 {
                                        v29 := []int{}
                                        _ = v29
                                    }
                                    v10 = v4
                                    switch (v10 + "\n\t") {
                                    case "\n\t":
                                        var v30 int16 = 768
                                        _ = v30
                                    default:
                                        v31 := []int{fi9, 468}
                                        _ = v31
                                    }
                                }
                                var v32 int32 = -1
                                _ = fi11
                                _ = v32
                            }
                        } else {
                            var v33 int16 = 0
                            _ = v33
                        }
                        if false {
                            v34 := GenS5{F0: v4, F1: v2, F2: 65535}
                            var v35 string = (((v4 + v10) + (((v10 + v4) + ("foo bar" + (v4 + "\n\t"))) + v10)) + ((v4 + "foo bar") + v10))
                            for fi36 := 0; fi36 < 10; fi36++ {
                                v37 := GenS5{F0: v10, F1: (-(v2)), F2: 65535}
                                v38 := GenS6{F0: 15, F1: 65535, F2: -469}
                                v2 -= 1
                                _ = fi36
                                _ = v37
                                _ = v38
                            }
                            _ = v34
                            _ = v35
                        }
                    }
                default:
                    v5[(v1 + fi9)] = 1
                }
                for fi39 := 0; fi39 < 17; fi39++ {
                    var v40 int8 = 99
                    v41 := GenS4{F0: ((1 != 42) || (!false)), F1: false, F2: 239}
                    _ = fi39
                    _ = v40
                    _ = v41
                }
                _ = fi9
                _ = v10
            }
        case "hello":
            v42 := []int{}
            _ = v42
        case "foo bar":
            fmt.Println(v3, v2)
        default:
            v1 = (v1 - (v1 << ((4294967295) & 63)))
        }
    }
    switch 32768 {
    case 2147483647:
        if true {
            v43 := GenS7{F0: 18446744073709551615, F1: 133, F2: 473}
            v3 = true
            _ = v43
        } else {
            if v3 {
                if (!v3) {
                    var v44 int16 = 1
                    fmt.Println(v44, v1)
                    v4 = v4
                    _ = v44
                } else {
                    var v45 string = (((v4 + "\n\t") + "a") + "hello")
                    var v46 int32 = 0
                    v4 = (v4 + v4)
                    _ = v45
                    _ = v46
                }
                v3 = ((937 != 720) && (1 >= -1))
                for fi47 := 0; fi47 < 6; fi47++ {
                    var v48 float32 = -2.25
                    _ = fi47
                    _ = v48
                }
            }
            v49 := []int{v1, 191, (v1 - -690), (v1 % ((42) | 1))}
            if (436 < 42) {
                var v50 int = (v1 + (^(v1)))
                _ = v50
            }
            _ = v49
        }
    case -1:
        v4 = (((("\n\t" + "hello") + v4) + v4) + v4)
    default:
        v1 += v1
    }
    if (42 < 186) {
        if v3 {
            switch "a" {
            case "foo bar":
                fmt.Println(v1, v2, v4)
            }
            v3 = (1 == 2147483647)
        } else {
            v5[v1] = v1
            var v51 int32 = -745
            for fi52 := 0; fi52 < 15; fi52++ {
                v2 -= v2
                v5[128] = (v1 >> ((600) & 63))
                v53 := GenS4{F0: ((2147483647 <= 791) || v3), F1: (1000 < 0), F2: 255}
                _ = fi52
                _ = v53
            }
            _ = v51
        }
        v4 = v4
        var v54 uint = 0
        _ = v54
    } else {
        for fi55 := 0; fi55 < 3; fi55++ {
            if v3 {
                for fi56 := 0; fi56 < 11; fi56++ {
                    var v57 bool = v3
                    switch "" {
                    case "hello":
                        v2 = 127
                    case "a":
                        var v58 uint64 = 0
                        _ = v58
                    }
                    _ = fi56
                    _ = v57
                }
                if v3 {
                    if false {
                        var v59 int = (v1 << ((54) & 63))
                        var v60 uint16 = 1
                        _ = v59
                        _ = v60
                    }
                } else {
                    _ = len(v5)
                }
                if (1e10 <= 3.14159) {
                    for fi61 := 0; fi61 < 2; fi61++ {
                        if v3 {
                            v4 = ("" + (v4 + "a"))
                            fmt.Println(v2)
                        }
                        var v62 int32 = 1
                        _ = fi61
                        _ = v62
                    }
                }
            }
            _ = fi55
        }
    }
    switch v1 {
    case -9223372036854775808:
        switch (v1 & (v1 << ((255) & 63))) {
        case 0:
            var v63 string = v4
            _ = v63
        default:
            switch v4 {
            case "a":
                for fi64 := 0; fi64 < 3; fi64++ {
                    if false {
                        if true {
                            var v65 int16 = -907
                            fmt.Println(v65, v3)
                            var v66 int64 = 287
                            _ = v65
                            _ = v66
                        }
                        switch v1 {
                        case 32767:
                            delete(v5, v1)
                        }
                        v2 = (v2 & (v2 >> ((355) & 63)))
                    }
                    for fi67 := 0; fi67 < 10; fi67++ {
                        var v68 uint16 = 65535
                        v2 -= (v2 ^ 51)
                        v69 := GenS8{F0: 102}
                        _ = fi67
                        _ = v68
                        _ = v69
                    }
                    v3 = v3
                    _ = fi64
                }
            case "\n\t":
                fmt.Println(v4, v4, v3)
            default:
                var v70 int32 = 101
                _ = v70
            }
        }
    case -17:
        var v71 uint64 = 1
        _ = v71
    case 42:
        switch (-(v1)) {
        case -9223372036854775808:
            for fi72 := 0; fi72 < 9; fi72++ {
                switch (^(v1)) {
                case 4294967295:
                    v73 := []int{fi72, 128}
                    _ = v73
                case -9223372036854775808:
                    v5[fi72] = v1
                case 32767:
                    for fi74 := 0; fi74 < 9; fi74++ {
                        v75 := []int{127, (-(v1))}
                        if v3 {
                            delete(v5, v1)
                        }
                        _ = len(v75)
                        _ = fi74
                        _ = v75
                    }
                default:
                    switch (fi72 | (-(fi72))) {
                    case -9223372036854775808:
                        v4 = ((v4 + (("" + (((v4 + v4) + v4) + "foo bar")) + ((("a" + ("hello" + "foo bar")) + (((("\n\t" + v4) + v4) + v4) + v4)) + v4))) + v4)
                    case 32767:
                        v1 = (-(v1))
                    default:
                        v76 := GenS5{F0: "", F1: (v2 % ((v2) | 1)), F2: 620}
                        _ = v76
                    }
                }
                var v77 float64 = 0.0
                if (v2 == (v2 ^ (v2 ^ (v2 - v2)))) {
                    var v78 uint = 65535
                    var v79 float64 = -2.25
                    _ = v78
                    _ = v79
                }
                _ = fi72
                _ = v77
            }
        case 128:
            var v80 int32 = -677
            _ = v80
        case -250:
            switch v4 {
            case "a":
                for fi81 := 0; fi81 < 4; fi81++ {
                    for fi82 := 0; fi82 < 15; fi82++ {
                        var v83 uint16 = 445
                        v84 := v5[fi82]
                        if true {
                            v84 = v84
                        }
                        _ = fi82
                        _ = v83
                        _ = v84
                    }
                    fmt.Println(v2, v3, v3)
                    _ = fi81
                }
            }
        default:
            v1 += v1
        }
    }
    v85 := v5[128]
    v2 = (v2 >> ((277) & 63))
    fmt.Println(v85, v3)
    fmt.Println(v4, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v85:", v85)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v85
}
