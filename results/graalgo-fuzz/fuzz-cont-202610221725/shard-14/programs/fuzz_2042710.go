package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 uint8; F1 string }
type GenS5 struct { F0 string; F1 uint8; F2 int64 }
type GenS6 struct { F0 int; F1 uint }
type GenS7 struct { F0 int; F1 int; F2 uint16 }

func main() {
    var v1 bool = ((true || false) || false)
    v2 := []uint32{1}
    var v3 int = -624
    v4 := []string{}
    var v5 int64 = 515
    v6 := make(map[int]int)
    v7 := GenS1{F0: 42, F1: ((!false) && (!false))}
    v8 := v6[v3]
    _ = len(v6)
    if (!true) {
        v2 = append(v2, 0)
        for fi9 := 0; fi9 < 6; fi9++ {
            v8 -= (v3 | -84)
            delete(v6, (^(fi9)))
            _ = fi9
        }
    }
    fmt.Println(v1, v8, v1)
    switch "\n\t" {
    case "":
        if v1 {
            if v1 {
                switch v3 {
                case -550:
                    var v10 bool = (!v1)
                    _ = v10
                default:
                    v1 = true
                }
            } else {
                var v11 int8 = 1
                _ = v11
            }
        } else {
            if (!false) {
                var v12 bool = ((("a" + "") + ("hello" + "")) < "")
                _ = v12
            } else {
                if (-736 <= v5) {
                    for fi13 := 0; fi13 < 11; fi13++ {
                        v8 = v3
                        _ = fi13
                    }
                } else {
                    switch (("hello" + ("hello" + ("foo bar" + "a"))) + ("" + ("\n\t" + ("a" + "foo bar")))) {
                    case "a":
                        var v14 float64 = float64((0.1 + 0.0))
                        _ = v14
                    default:
                        if v1 {
                            v8 = (v8 * -1)
                            switch 183 {
                            case 9223372036854775807:
                                v1 = ("hello" > "hello")
                            }
                            var v15 uint16 = 860
                            _ = v15
                        } else {
                            if (v1 || (!(!(!((266 != 42) && false))))) {
                                v16 := GenS2{F0: 0}
                                if (!(1.5 != float64(1e10))) {
                                    switch 32767 {
                                    case 2147483647:
                                        v17 := GenS1{F0: -1, F1: v1}
                                        _ = v17
                                    case 4294967295:
                                        if v1 {
                                            for fi18 := 0; fi18 < 19; fi18++ {
                                                if v1 {
                                                    v8 -= v8
                                                    if (1 == 2) {
                                                        v3 = v8
                                                        v19 := GenS3{F0: 1, F1: 406}
                                                        _ = v19
                                                    }
                                                } else {
                                                    v3 -= fi18
                                                    delete(v6, fi18)
                                                }
                                                v20 := GenS1{F0: 124, F1: v1}
                                                var v21 uint8 = 2
                                                _ = fi18
                                                _ = v20
                                                _ = v21
                                            }
                                        }
                                    default:
                                        switch (v8 >> ((187) & 63)) {
                                        case -396:
                                            switch (v8 ^ (-(v8))) {
                                            case 9223372036854775807:
                                                var v22 uint64 = 42
                                                _ = v22
                                            default:
                                                if v1 {
                                                    v23 := GenS2{F0: 42}
                                                    var v24 int16 = 0
                                                    _ = v23
                                                    _ = v24
                                                }
                                            }
                                        case -243:
                                            fmt.Println(v5)
                                        case 127:
                                            if (false || (-128 > 1)) {
                                                for fi25 := 0; fi25 < 18; fi25++ {
                                                    switch (fi25 | (v3 & (v3 ^ (-(fi25))))) {
                                                    case 32767:
                                                        var v26 float32 = ((-2.25 / (((-2.25 * 3.14159) * 0.1) * 1e10)) + 1.5)
                                                        _ = v26
                                                    default:
                                                        for fi27 := 0; fi27 < 1; fi27++ {
                                                            var v28 uint8 = 53
                                                            _ = fi27
                                                            _ = v28
                                                        }
                                                    }
                                                    v29 := v6[554]
                                                    _ = fi25
                                                    _ = v29
                                                }
                                                var v30 uint16 = 42
                                                _ = v30
                                            }
                                        }
                                    }
                                    v31 := GenS1{F0: 100, F1: (127 < 127)}
                                    if true {
                                        fmt.Println(v5)
                                        switch v3 {
                                        case 128:
                                            v3 -= (v3 / ((v3) | 1))
                                        default:
                                            if true {
                                                v1 = (!v1)
                                            } else {
                                                if false {
                                                    switch 42 {
                                                    case 9223372036854775807:
                                                        if v1 {
                                                            switch v3 {
                                                            case 852:
                                                                _ = len(v4)
                                                            case 32768:
                                                                v32 := GenS4{F0: 163, F1: "foo bar"}
                                                                _ = v32
                                                            }
                                                        }
                                                    default:
                                                        for fi33 := 0; fi33 < 10; fi33++ {
                                                            _ = len(v6)
                                                            var v34 uint = 2
                                                            _ = fi33
                                                            _ = v34
                                                        }
                                                    }
                                                    var v35 int32 = -748
                                                    _ = v35
                                                }
                                            }
                                        }
                                    } else {
                                        v36 := GenS3{F0: 340, F1: v3}
                                        var v37 string = ("" + "a")
                                        v38 := GenS2{F0: 100}
                                        _ = v36
                                        _ = v37
                                        _ = v38
                                    }
                                    _ = v31
                                }
                                _ = v16
                            } else {
                                for fi39 := 0; fi39 < 12; fi39++ {
                                    switch (v3 * v8) {
                                    case 1:
                                        v1 = false
                                    case 9223372036854775807:
                                        v40 := GenS5{F0: "a", F1: 42, F2: 852}
                                        _ = v40
                                    case 0:
                                        if (!false) {
                                            var v41 int64 = -921
                                            var v42 int64 = (v5 >> ((42) & 63))
                                            _ = v41
                                            _ = v42
                                        }
                                    }
                                    _ = fi39
                                }
                                for fi43 := 0; fi43 < 20; fi43++ {
                                    switch ("a" + (("\n\t" + ("foo bar" + "a")) + "hello")) {
                                    case "":
                                        v4 = append(v4, (("a" + "hello") + ("a" + ("foo bar" + "foo bar"))))
                                    case "\n\t":
                                        v5 += (v5 / ((42) | 1))
                                    }
                                    fmt.Println(fi43, fi43, fi43)
                                    v5 = -390
                                    _ = fi43
                                }
                            }
                            v44 := GenS3{F0: 42, F1: v3}
                            for fi45 := 0; fi45 < 15; fi45++ {
                                for fi46 := 0; fi46 < 3; fi46++ {
                                    if v1 {
                                        if v1 {
                                            v47 := GenS5{F0: "a", F1: 218, F2: v5}
                                            if v1 {
                                                if v1 {
                                                    switch (fi46 / (((v3 + (v3 >> ((255) & 63)))) | 1)) {
                                                    case 127:
                                                        v4 = append(v4, "a")
                                                    }
                                                    var v48 uint32 = 170
                                                    _ = v48
                                                } else {
                                                    v1 = v1
                                                    if false {
                                                        if false {
                                                            v1 = v1
                                                            if len(v4) > 0 {
                                                                v4[0] = ""
                                                            }
                                                        }
                                                        if v1 {
                                                            v6[fi46] = -419
                                                            switch v8 {
                                                            case -722:
                                                                v8 += fi46
                                                                var v49 float32 = 3.14159
                                                                _ = v49
                                                            case -2147483648:
                                                                switch "a" {
                                                                case "a":
                                                                    v1 = (465 == 42)
                                                                }
                                                            }
                                                            fmt.Println(v1, v8)
                                                        } else {
                                                            v3 = v3
                                                            switch (v8 - (fi45 + (fi45 * -920))) {
                                                            case 128:
                                                                if v1 {
                                                                    var v50 int16 = 1000
                                                                    if ((0.0 - float32(0.0)) == 3.14159) {
                                                                        _ = len(v4)
                                                                        if ((1e10 - 1.5) <= 0.1) {
                                                                            if (!true) {
                                                                                var v51 float64 = (-2.25 - -2.25)
                                                                                var v52 int8 = -1
                                                                                v53 := GenS4{F0: 255, F1: "hello"}
                                                                                _ = v51
                                                                                _ = v52
                                                                                _ = v53
                                                                            }
                                                                        } else {
                                                                            var v54 int8 = 77
                                                                            fmt.Println(v54)
                                                                            _ = v54
                                                                        }
                                                                    } else {
                                                                        v50 -= -1
                                                                        v50 = (-(v50))
                                                                        var v55 string = (("" + "a") + ("a" + "a"))
                                                                        delete(v6, 161)
                                                                        var v56 int32 = -1
                                                                        _ = v55
                                                                        _ = v56
                                                                    }
                                                                    _ = v50
                                                                } else {
                                                                    delete(v6, fi46)
                                                                    switch (("foo bar" + (("a" + "") + "")) + ("hello" + ("foo bar" + "a"))) {
                                                                    case "\n\t":
                                                                        delete(v6, fi46)
                                                                    case "hello":
                                                                        var v57 uint8 = 127
                                                                        _ = v57
                                                                    case "foo bar":
                                                                        v5 -= (v5 | -9223372036854775808)
                                                                    }
                                                                    v58 := GenS6{F0: fi46, F1: 1}
                                                                    _ = v58
                                                                }
                                                            case 32768:
                                                                if v1 {
                                                                    var v59 int16 = 1000
                                                                    var v60 uint8 = 127
                                                                    _ = v59
                                                                    _ = v60
                                                                }
                                                            default:
                                                                v1 = v1
                                                            }
                                                        }
                                                    }
                                                    var v61 int32 = 293
                                                    _ = v61
                                                }
                                                var v62 int = (^(v3))
                                                v1 = (2 != 2)
                                                _ = v62
                                            } else {
                                                fmt.Println(v3, v8)
                                                v63 := v6[(^(fi46))]
                                                _ = v63
                                            }
                                            var v64 int64 = 365
                                            _ = v47
                                            _ = v64
                                        }
                                        var v65 int32 = 0
                                        _ = v65
                                    }
                                    if ((v5 % ((v5) | 1)) != v5) {
                                        v3 = 382
                                    }
                                    _ = fi46
                                }
                                for fi66 := 0; fi66 < 9; fi66++ {
                                    v3 = 32768
                                    v1 = (!v1)
                                    var v67 uint32 = 822
                                    _ = fi66
                                    _ = v67
                                }
                                fmt.Println(fi45, fi45, v5)
                                _ = fi45
                            }
                            _ = v44
                        }
                    }
                }
                var v68 uint32 = 661
                switch (v8 / ((v8) | 1)) {
                case 1:
                    v5 = 1
                case -38:
                    _ = len(v6)
                case 0:
                    v3 -= (v3 << ((148) & 63))
                }
                _ = v68
            }
            switch (v8 | v8) {
            case 237:
                for fi69 := 0; fi69 < 17; fi69++ {
                    v70 := GenS7{F0: (v3 & v3), F1: v8, F2: 42}
                    _ = fi69
                    _ = v70
                }
            case 32768:
                for fi71 := 0; fi71 < 8; fi71++ {
                    for fi72 := 0; fi72 < 16; fi72++ {
                        v5 = (v5 << ((4294967295) & 63))
                        _ = fi72
                    }
                    v1 = v1
                    _ = fi71
                }
            }
        }
    case "\n\t":
        v73 := GenS4{F0: 246, F1: "foo bar"}
        _ = v73
    }
    var v74 uint8 = 98
    var v75 int32 = 0
    v5 += v5
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v74
    _ = v75
    _ = v8
}
