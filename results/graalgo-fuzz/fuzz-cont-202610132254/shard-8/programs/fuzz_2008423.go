package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 uint64; F1 int16; F2 int64 }
type GenS4 struct { F0 int8; F1 string }
type GenS5 struct { F0 int32; F1 uint64; F2 bool }
type GenS6 struct { F0 int16; F1 uint16 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS9 struct { F0 int; F1 int16 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint64; F1 uint16 }

func main() {
    var v1 int16 = 0
    var v2 uint32 = 0
    var v3 uint64 = 97
    var v4 int16 = (v1 ^ 42)
    var v5 bool = false
    v5 = false
    var v6 bool = v5
    switch -2147483648 {
    case -297:
        v7 := GenS1{F0: 58}
        _ = v7
    case -9223372036854775808:
        if (true && (1 == 255)) {
            var v8 uint8 = 128
            _ = v8
        }
    default:
        v9 := GenS2{F0: 9223372036854775807, F1: 998}
        _ = v9
    }
    v10 := []int{}
    v11 := GenS3{F0: (v3 ^ v3), F1: 0, F2: int64(v2)}
    v5 = (80 <= -1)
    if ((0.1 / 0.1) == 1e10) {
        for fi12 := 0; fi12 < 5; fi12++ {
            v3 = v3
            if (v3 < v3) {
                v4 -= (v4 ^ v4)
                fmt.Println(v3, v3)
                fmt.Println(v2)
            } else {
                var v13 float32 = 1e10
                var v14 int = fi12
                _ = v13
                _ = v14
            }
            var v15 uint16 = 975
            _ = fi12
            _ = v15
        }
    } else {
        v16 := GenS4{F0: -1, F1: ("\n\t" + "hello")}
        switch 649 {
        case 32767:
            v1 += (v1 | (v1 + -1))
        case 330:
            if len(v10) > 0 {
                v10[0] = -241
            }
        default:
            if false {
                var v17 int = 482
                v10 = append(v10, (v17 << ((454) & 63)))
                _ = v17
            }
        }
        v18 := GenS4{F0: 90, F1: ""}
        _ = v16
        _ = v18
    }
    v19 := GenS3{F0: v3, F1: (v1 - (v1 * v4)), F2: 9223372036854775807}
    for fi20 := 0; fi20 < 13; fi20++ {
        v21 := GenS5{F0: 2147483647, F1: v3, F2: (0 > 705)}
        v22 := GenS2{F0: 1, F1: (v3 ^ (v3 - 653))}
        var v23 int64 = -9223372036854775808
        _ = fi20
        _ = v21
        _ = v22
        _ = v23
    }
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    if true {
        v24 := GenS2{F0: 1, F1: v3}
        if v5 {
            var v25 uint16 = 113
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            _ = v25
        }
        _ = v24
    }
    fmt.Println(v6, v6, v2)
    v1 = 42
    for fi26 := 0; fi26 < 16; fi26++ {
        fmt.Println(v4, v1, v3)
        _ = fi26
    }
    var v27 int8 = 1
    switch "foo bar" {
    case "a":
        switch 32768 {
        case -116:
            var v28 uint8 = 255
            _ = v28
        case 352:
            v5 = (-1 > 2147483647)
        case 336:
            switch 42 {
            case -853:
                switch -1 {
                case 128:
                    v27 = (v27 * -1)
                case 9223372036854775807:
                    for fi29 := 0; fi29 < 9; fi29++ {
                        if v5 {
                            v5 = true
                        }
                        _ = fi29
                    }
                case 139:
                    v4 = v4
                default:
                    switch "foo bar" {
                    case "\n\t":
                        if len(v10) > 0 {
                            v10[0] = -835
                        }
                    case "hello":
                        switch "\n\t" {
                        case "foo bar":
                            v30 := GenS6{F0: v4, F1: 0}
                            _ = v30
                        case "a":
                            v31 := GenS4{F0: 42, F1: "hello"}
                            _ = v31
                        case "\n\t":
                            if (1.5 >= 1e10) {
                                for fi32 := 0; fi32 < 14; fi32++ {
                                    if (v6 || (true && v6)) {
                                        if v6 {
                                            var v33 int = fi32
                                            _ = v33
                                        } else {
                                            for fi34 := 0; fi34 < 16; fi34++ {
                                                if ((-577 >= 552) || v5) {
                                                    var v35 float64 = ((float64(1e10) + ((1e10 - 0.0) / -2.25)) * (((3.14159 / 1.0) + 1e10) * -2.25))
                                                    v36 := GenS2{F0: -9223372036854775808, F1: (v3 | (^(v3)))}
                                                    v27 = (^(v27))
                                                    _ = v35
                                                    _ = v36
                                                }
                                                _ = fi34
                                            }
                                            if ((fi32 | fi32) <= -1) {
                                                var v37 int64 = 42
                                                _ = len(v10)
                                                switch "hello" {
                                                case "hello":
                                                    v1 -= 42
                                                case "foo bar":
                                                    v10 = append(v10, (fi32 + int(v37)))
                                                case "\n\t":
                                                    v37 -= (v37 << ((255) & 63))
                                                default:
                                                    v38 := GenS7{F0: 0}
                                                    _ = v38
                                                }
                                                _ = v37
                                            }
                                            var v39 int16 = v4
                                            _ = v39
                                        }
                                    } else {
                                        if (255 > 18446744073709551615) {
                                            if len(v10) >= 2 {
                                                v10 = v10[1:len(v10)]
                                            }
                                            v2 = (v2 / (((v2 + 2)) | 1))
                                            var v40 uint32 = v2
                                            _ = v40
                                        } else {
                                            v41 := GenS8{F0: 0, F1: 2147483647, F2: v27}
                                            for fi42 := 0; fi42 < 10; fi42++ {
                                                if len(v10) >= 2 {
                                                    v10 = v10[1:len(v10)]
                                                }
                                                _ = fi42
                                            }
                                            _ = v41
                                        }
                                    }
                                    v5 = false
                                    if (947 < 42) {
                                        fmt.Println(v2, v3)
                                        v5 = v6
                                    } else {
                                        v2 += (^(v2))
                                    }
                                    _ = fi32
                                }
                                fmt.Println(v4, v4, v5)
                            } else {
                                for fi43 := 0; fi43 < 1; fi43++ {
                                    var v44 int32 = 42
                                    v45 := GenS7{F0: 65535}
                                    for fi46 := 0; fi46 < 1; fi46++ {
                                        var v47 float64 = 1.5
                                        var v48 int32 = 42
                                        v49 := GenS8{F0: 406, F1: v2, F2: v27}
                                        _ = fi46
                                        _ = v47
                                        _ = v48
                                        _ = v49
                                    }
                                    _ = fi43
                                    _ = v44
                                    _ = v45
                                }
                            }
                        }
                    case "foo bar":
                        v5 = v5
                    }
                }
            }
        default:
            fmt.Println(v4, v3, v4)
        }
    case "":
        for fi50 := 0; fi50 < 14; fi50++ {
            var v51 bool = v6
            v52 := GenS9{F0: (fi50 & (fi50 | (fi50 + -9223372036854775808))), F1: (v1 % ((v4) | 1))}
            _ = fi50
            _ = v51
            _ = v52
        }
    }
    if (!v6) {
        for fi53 := 0; fi53 < 3; fi53++ {
            v54 := GenS10{F0: -362}
            _ = fi53
            _ = v54
        }
        v5 = (((42 > 593) || ((1 == v2) || v5)) || (255 < 63))
    } else {
        var v55 uint = 255
        if false {
            var v56 float32 = 0.0
            if (((!true) && false) && v6) {
                if ((42 < -928) || false) {
                    for fi57 := 0; fi57 < 6; fi57++ {
                        fmt.Println(v1, v56)
                        if v6 {
                            v2 = 4294967295
                        }
                        _ = fi57
                    }
                    switch ("" + ("foo bar" + "hello")) {
                    case "hello":
                        var v58 int64 = 42
                        _ = v58
                    }
                } else {
                    v59 := GenS11{F0: v3, F1: 0}
                    switch -1 {
                    case -348:
                        var v60 float64 = -2.25
                        _ = v60
                    case 128:
                        v61 := GenS11{F0: 18446744073709551615, F1: 0}
                        _ = v61
                    case 127:
                        if v5 {
                            switch ("hello" + ("\n\t" + (("" + "a") + "foo bar"))) {
                            case "a":
                                v62 := GenS7{F0: 42}
                                _ = v62
                            case "\n\t":
                                v1 -= v1
                            case "":
                                var v63 uint8 = 1
                                _ = v63
                            }
                            for fi64 := 0; fi64 < 6; fi64++ {
                                fmt.Println(v3, fi64, v56)
                                v5 = v5
                                for fi65 := 0; fi65 < 6; fi65++ {
                                    v6 = (((!(!(42 <= 0))) || (v1 > 827)) || (v3 < 162))
                                    v56 += 3.14159
                                    v3 = (^(v3))
                                    _ = fi65
                                }
                                _ = fi64
                            }
                            if (((0.1 * ((0.0 - 3.14159) * 1.5)) + float64((v56 * v56))) != (0.1 - (1.5 - 0.0))) {
                                var v66 int64 = 1
                                switch -620 {
                                case 55:
                                    if len(v10) >= 2 {
                                        v10 = v10[1:len(v10)]
                                    }
                                case -9223372036854775808:
                                    var v67 uint64 = (v3 >> ((v55) & 63))
                                    _ = v67
                                case 713:
                                    switch 544 {
                                    case -895:
                                        var v68 float64 = -2.25
                                        _ = v68
                                    case 0:
                                        v1 = v1
                                    }
                                default:
                                    v55 += 42
                                }
                                if false {
                                    var v69 float64 = 1.5
                                    _ = v69
                                }
                                _ = v66
                            }
                        }
                    default:
                        var v70 uint16 = uint16(v55)
                        _ = v70
                    }
                    _ = v59
                }
            }
            _ = v56
        } else {
            fmt.Println(v1)
            var v71 int32 = -2147483648
            var v72 int64 = 42
            _ = v71
            _ = v72
        }
        _ = v55
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
