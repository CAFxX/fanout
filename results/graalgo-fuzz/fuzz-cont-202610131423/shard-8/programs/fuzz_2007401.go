package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 string }
type GenS2 struct { F0 uint8; F1 int8; F2 uint64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int64; F1 int }
type GenS6 struct { F0 uint64; F1 int8 }
type GenS7 struct { F0 int; F1 uint; F2 int64 }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 float32 = 0.0
    var v3 int16 = -458
    v4 := []int{-542, -1, -936}
    var v5 uint16 = 0
    for fi6 := 0; fi6 < 18; fi6++ {
        if ("foo bar" >= "foo bar") {
            v1 = (-(v1))
        }
        switch 133 {
        case 2147483647:
            v3 = (v3 & (v3 | (v3 * (v3 - 32767))))
        case 4294967295:
            if true {
                if len(v4) > 0 {
                    v4[0] = 32768
                }
            }
        default:
            var v7 uint64 = 0
            _ = v7
        }
        _ = fi6
    }
    if (0 >= -14) {
        v8 := GenS1{F0: 768, F1: v1, F2: "\n\t"}
        for fi9 := 0; fi9 < 20; fi9++ {
            if false {
                var v10 string = ("foo bar" + "\n\t")
                fmt.Println(v5)
                fmt.Println(v3, fi9, v3)
                _ = v10
            } else {
                switch fi9 {
                case 0:
                    fmt.Println(v1, v3, v3)
                default:
                    v4 = append(v4, (fi9 * (fi9 | fi9)))
                }
                if ((237 > 1) && ((255 < 1) && (18446744073709551615 >= 2))) {
                    v11 := GenS1{F0: (v5 | (v5 % ((29) | 1))), F1: v1, F2: "foo bar"}
                    _ = v11
                }
                var v12 uint16 = v5
                _ = v12
            }
            _ = fi9
        }
        _ = v8
    } else {
        if (798 != 32768) {
            v4 = append(v4, 42)
        } else {
            for fi13 := 0; fi13 < 2; fi13++ {
                fmt.Println(fi13, v2)
                for fi14 := 0; fi14 < 9; fi14++ {
                    switch "" {
                    case "\n\t":
                        v2 = v2
                    default:
                        v15 := GenS1{F0: (v5 % ((v5) | 1)), F1: v1, F2: "foo bar"}
                        _ = v15
                    }
                    if (42 != 18446744073709551615) {
                        var v16 uint = 1
                        v3 = 1000
                        switch "hello" {
                        case "\n\t":
                            var v17 float32 = float32(float64(float32(-2.25)))
                            _ = v17
                        default:
                            v2 = (0.1 - v2)
                        }
                        _ = v16
                    }
                    _ = fi14
                }
                _ = fi13
            }
        }
        switch 655 {
        case 128:
            var v18 int = 32767
            _ = v18
        case 638:
            v4 = append(v4, -9223372036854775808)
        case 9223372036854775807:
            switch "a" {
            case "\n\t":
                v5 = v5
            case "foo bar":
                v1 = (-(v1))
            }
        default:
            fmt.Println(v1, v2)
        }
    }
    var v19 uint8 = 1
    var v20 uint32 = 975
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi21 := 0; fi21 < 17; fi21++ {
        if (0.0 > float64(v2)) {
            var v22 int32 = 299
            _ = v22
        }
        _ = fi21
    }
    if true {
        for fi23 := 0; fi23 < 18; fi23++ {
            var v24 int16 = (v3 % ((v3) | 1))
            _ = fi23
            _ = v24
        }
        var v25 uint8 = 30
        _ = v25
    }
    for fi26 := 0; fi26 < 5; fi26++ {
        var v27 uint16 = (^(v5))
        switch "" {
        case "a":
            v28 := GenS1{F0: (v27 & (v27 | uint16(v19))), F1: int64(v3), F2: ""}
            _ = v28
        default:
            for fi29 := 0; fi29 < 9; fi29++ {
                var v30 uint16 = 32767
                _ = fi29
                _ = v30
            }
        }
        if (true || true) {
            if true {
                for fi31 := 0; fi31 < 2; fi31++ {
                    v4 = append(v4, fi26)
                    _ = fi31
                }
                v2 = v2
            }
        }
        _ = fi26
        _ = v27
    }
    _ = len(v4)
    fmt.Println(v19, v1)
    v32 := GenS1{F0: (v5 * v5), F1: 9223372036854775807, F2: "a"}
    switch "a" {
    case "a":
        fmt.Println(v2)
    case "\n\t":
        for fi33 := 0; fi33 < 20; fi33++ {
            var v34 int16 = 32767
            _ = fi33
            _ = v34
        }
    case "foo bar":
        v35 := GenS2{F0: (v19 / ((v19) | 1)), F1: 100, F2: 4294967295}
        _ = v35
    default:
        var v36 uint16 = 2
        _ = v36
    }
    for fi37 := 0; fi37 < 5; fi37++ {
        v5 -= (^(v5))
        if (!true) {
            var v38 float64 = (-2.25 - 3.14159)
            v39 := GenS3{F0: 186}
            switch fi37 {
            case 2147483647:
                if false {
                    if (698 >= -2147483648) {
                        var v40 uint32 = (v20 * 654)
                        _ = v40
                    } else {
                        if false {
                            var v41 uint64 = 1
                            _ = v41
                        }
                        if true {
                            fmt.Println(v2, v19)
                            for fi42 := 0; fi42 < 7; fi42++ {
                                if false {
                                    v43 := GenS4{F0: ""}
                                    _ = v43
                                }
                                _ = fi42
                            }
                        } else {
                            switch "" {
                            case "\n\t":
                                v44 := GenS4{F0: "\n\t"}
                                _ = v44
                            }
                            v5 = 65535
                            var v45 string = (("foo bar" + "\n\t") + ("hello" + ""))
                            _ = v45
                        }
                        v1 -= -304
                    }
                    switch fi37 {
                    case 128:
                        if len(v4) > 0 {
                            v4[0] = (-(fi37))
                        }
                    case -494:
                        if true {
                            v4 = append(v4, -928)
                            var v46 uint8 = (v19 ^ (v19 - v19))
                            var v47 uint = 0
                            _ = v46
                            _ = v47
                        } else {
                            v4 = append(v4, (fi37 >> ((559) & 63)))
                            var v48 uint32 = v20
                            _ = v48
                        }
                    case 13:
                        var v49 uint8 = v19
                        _ = v49
                    default:
                        v50 := GenS5{F0: (v1 << ((208) & 63)), F1: -438}
                        _ = v50
                    }
                    v3 = (v3 >> ((273) & 63))
                }
            case -411:
                switch "a" {
                case "foo bar":
                    v51 := GenS5{F0: 1, F1: (^(fi37))}
                    _ = v51
                case "\n\t":
                    var v52 uint16 = 65535
                    _ = v52
                case "hello":
                    var v53 string = (("\n\t" + "\n\t") + "a")
                    _ = v53
                default:
                    var v54 int32 = -428
                    _ = v54
                }
            default:
                v1 = -1
            }
            _ = v38
            _ = v39
        } else {
            if true {
                v55 := GenS6{F0: 637, F1: 1}
                _ = v55
            } else {
                v56 := GenS5{F0: v1, F1: fi37}
                _ = v56
            }
        }
        _ = fi37
    }
    var v57 int = 2147483647
    for fi58 := 0; fi58 < 12; fi58++ {
        switch (("a" + "hello") + ("\n\t" + "")) {
        case "foo bar":
            v20 = (v20 | v20)
        case "a":
            if true {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        default:
            v1 = (v1 << ((65535) & 63))
        }
        v59 := GenS2{F0: (v19 & v19), F1: 127, F2: 2}
        switch "hello" {
        case "":
            v60 := GenS7{F0: (v57 >> ((407) & 63)), F1: 1, F2: 42}
            _ = v60
        case "foo bar":
            var v61 int8 = 75
            _ = v61
        case "\n\t":
            for fi62 := 0; fi62 < 2; fi62++ {
                switch fi58 {
                case 765:
                    v19 -= (v19 * v19)
                }
                v20 -= 65535
                _ = fi62
            }
        }
        _ = fi58
        _ = v59
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v57
}
