package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 uint8; F2 int16 }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 int64; F1 bool }
type GenS5 struct { F0 int; F1 int64 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int; F1 string }
type GenS9 struct { F0 uint; F1 int32 }
type GenS10 struct { F0 uint; F1 int64; F2 int16 }
type GenS11 struct { F0 int; F1 bool }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]int)
    v3 := make(map[string]bool)
    v4 := []int{4294967295, -2147483648}
    v5 := GenS1{F0: "a"}
    if (640 == 42) {
        _ = len(v1)
        if (!("foo bar" <= "hello")) {
            switch 2147483647 {
            case 1:
                v6 := GenS2{F0: 561, F1: 0, F2: -1}
                _ = v6
            case 278:
                v7 := GenS3{F0: 429, F1: 2}
                _ = v7
            case 271:
                var v8 uint16 = 1
                _ = v8
            default:
                var v9 int16 = 429
                _ = v9
            }
        }
    }
    if true {
        var v10 uint64 = 1
        var v11 int = -261
        var v12 int32 = 1
        _ = v10
        _ = v11
        _ = v12
    } else {
        var v13 int16 = -32768
        if true {
            if (1 <= 1) {
                if len(v4) > 0 {
                    v4[0] = 42
                }
                if (!(!(42 < -319))) {
                    fmt.Println(v13)
                    switch "hello" {
                    case "hello":
                        if (((v13 + (v13 >> ((2) & 63))) != (v13 | (v13 + (v13 * (^(v13)))))) || false) {
                            for fi14 := 0; fi14 < 19; fi14++ {
                                var v15 int64 = -9223372036854775808
                                _ = fi14
                                _ = v15
                            }
                            var v16 float32 = -2.25
                            _ = v16
                        } else {
                            v13 = (v13 * (v13 >> ((2) & 63)))
                        }
                    case "a":
                        for fi17 := 0; fi17 < 10; fi17++ {
                            switch (fi17 + fi17) {
                            case 4294967295:
                                fmt.Println(v13)
                            default:
                                if (v13 != (v13 * (v13 / (((v13 % ((v13) | 1))) | 1)))) {
                                    var v18 uint32 = 1
                                    if true {
                                        for fi19 := 0; fi19 < 12; fi19++ {
                                            var v20 string = "a"
                                            var v21 uint16 = 42
                                            var v22 uint16 = (v21 - v21)
                                            fmt.Println(v18)
                                            _ = fi19
                                            _ = v20
                                            _ = v21
                                            _ = v22
                                        }
                                        for fi23 := 0; fi23 < 10; fi23++ {
                                            v4 = append(v4, (^(fi17)))
                                            v13 += v13
                                            v13 -= v13
                                            _ = fi23
                                        }
                                    }
                                    _ = v18
                                }
                            }
                            v24 := GenS1{F0: ("\n\t" + "foo bar")}
                            if true {
                                v25 := GenS2{F0: (v13 ^ 1), F1: 1, F2: (^(v13))}
                                v13 += (v13 & v13)
                                _ = v25
                            }
                            _ = fi17
                            _ = v24
                        }
                    default:
                        if (-2.25 == 1e10) {
                            v13 = 0
                            switch -817 {
                            case 2147483647:
                                switch (((("" + ("foo bar" + "a")) + "a") + "a") + "a") {
                                case "hello":
                                    switch 2147483647 {
                                    case 382:
                                        delete(v2, "hello")
                                    case 480:
                                        v13 = 889
                                    case 2147483647:
                                        fmt.Println(v13)
                                    default:
                                        v4 = append(v4, 536)
                                    }
                                }
                            default:
                                v13 -= (v13 | v13)
                            }
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        } else {
                            v13 = 1000
                            v26 := GenS4{F0: -766, F1: true}
                            for fi27 := 0; fi27 < 13; fi27++ {
                                v28 := GenS4{F0: -1, F1: (v13 != v13)}
                                for fi29 := 0; fi29 < 14; fi29++ {
                                    v13 = 0
                                    _ = fi29
                                }
                                _ = fi27
                                _ = v28
                            }
                            _ = v26
                        }
                    }
                    switch "a" {
                    case "a":
                        var v30 bool = true
                        _ = v30
                    default:
                        fmt.Println(v13)
                    }
                } else {
                    fmt.Println(v13)
                    v31 := GenS2{F0: -1, F1: 2, F2: (v13 - v13)}
                    _ = v31
                }
            } else {
                v32 := GenS5{F0: -1, F1: -385}
                v13 = (-(v13))
                var v33 int64 = 933
                _ = v32
                _ = v33
            }
        } else {
            v34 := GenS6{F0: 90}
            v13 = v13
            _ = v34
        }
        _ = v13
    }
    var v35 float32 = float32(3.14159)
    for fi36 := 0; fi36 < 4; fi36++ {
        v37 := v3[(("a" + "\n\t") + "")]
        var v38 bool = ((42 <= 2) && v37)
        v35 = ((3.14159 / (0.1 - 3.14159)) / -2.25)
        _ = fi36
        _ = v37
        _ = v38
    }
    for fi39 := 0; fi39 < 3; fi39++ {
        if true {
            var v40 uint64 = 323
            _ = v40
        }
        var v41 uint16 = 1
        for fi42 := 0; fi42 < 11; fi42++ {
            v43 := GenS7{F0: 18446744073709551615}
            if true {
                v4 = append(v4, -434)
            } else {
                switch fi42 {
                case 32768:
                    var v44 uint32 = 2
                    _ = v44
                case 128:
                    v2["foo bar"] = fi42
                case 1:
                    if false {
                        v45 := GenS8{F0: fi39, F1: (((("" + ("\n\t" + "\n\t")) + "hello") + ("\n\t" + (("hello" + ("a" + "hello")) + "foo bar"))) + "\n\t")}
                        _ = v45
                    }
                }
                v46 := GenS9{F0: 1, F1: 1}
                if ((-1 <= 32767) && (273 != (fi39 - (fi39 ^ fi39)))) {
                    var v47 uint32 = 2147483647
                    _ = v47
                } else {
                    if (65535 > 255) {
                        v41 = 991
                        v48 := v2["hello"]
                        v35 -= (1e10 - 1e10)
                        switch v48 {
                        case -657:
                            if false {
                                var v49 string = "foo bar"
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                _ = v49
                            } else {
                                var v50 uint16 = v41
                                fmt.Println(v50)
                                switch (fi39 / ((v48) | 1)) {
                                case 0:
                                    fmt.Println(fi42)
                                case -2147483648:
                                    _ = len(v3)
                                }
                                switch "\n\t" {
                                case "a":
                                    fmt.Println(fi39, v48, fi42)
                                default:
                                    v50 = v50
                                    if (!(false || (false || true))) {
                                        v50 -= (^(v50))
                                        fmt.Println(v35, v41)
                                        switch v48 {
                                        case -502:
                                            if ((v35 > 0.0) || ("a" != ("a" + "foo bar"))) {
                                                var v51 int64 = -9223372036854775808
                                                _ = len(v1)
                                                fmt.Println(v51, v50, v51)
                                                _ = v51
                                            }
                                            v41 = (v41 - (v50 << ((uint(v41)) & 63)))
                                            v52 := GenS8{F0: fi42, F1: ("a" + "foo bar")}
                                            _ = v52
                                        case 32768:
                                            var v53 uint64 = 0
                                            _ = v53
                                        default:
                                            v41 += (^(v41))
                                        }
                                    } else {
                                        _ = len(v1)
                                        v54 := GenS5{F0: (fi42 / (((v48 / ((v48) | 1))) | 1)), F1: 42}
                                        v41 = v50
                                        _ = v54
                                    }
                                }
                                if true {
                                    var v55 float32 = (1.5 * ((v35 * v35) + float32(1e10)))
                                    _ = v55
                                } else {
                                    v56 := v3["a"]
                                    _ = v56
                                }
                                _ = v50
                            }
                        }
                        v35 = 0.0
                        _ = v48
                    }
                    if (v41 != v41) {
                        v57 := GenS10{F0: 946, F1: 0, F2: -1}
                        _ = v57
                    }
                    v41 -= 2
                }
                v41 = (v41 - v41)
                _ = v46
            }
            v58 := GenS8{F0: (-(fi39)), F1: "hello"}
            _ = fi42
            _ = v43
            _ = v58
        }
        _ = fi39
        _ = v41
    }
    v35 += 0.1
    if len(v4) > 0 {
        v4[0] = 1
    }
    if false {
        v35 = float32(3.14159)
        for fi59 := 0; fi59 < 20; fi59++ {
            var v60 bool = (803 <= 18446744073709551615)
            v60 = true
            var v61 int32 = 42
            _ = fi59
            _ = v60
            _ = v61
        }
    }
    switch 4294967295 {
    case -476:
        for fi62 := 0; fi62 < 9; fi62++ {
            if (4294967295 >= 592) {
                var v63 uint32 = 786
                var v64 uint = 42
                _ = v63
                _ = v64
            }
            var v65 uint8 = 1
            for fi66 := 0; fi66 < 6; fi66++ {
                if true {
                    var v67 int16 = -1
                    delete(v3, "a")
                    _ = v67
                }
                var v68 int32 = -1
                _ = fi66
                _ = v68
            }
            _ = fi62
            _ = v65
        }
    default:
        if true {
            var v69 float32 = (v35 - (0.1 - (0.0 / v35)))
            var v70 bool = true
            v3["a"] = v70
            _ = v69
            _ = v70
        } else {
            if false {
                var v71 int = 971
                _ = v71
            }
            var v72 uint8 = 255
            if (0 == 255) {
                switch 127 {
                case 2147483647:
                    v72 = v72
                }
                _ = len(v1)
                var v73 uint64 = 388
                _ = v73
            } else {
                fmt.Println(v35, v35)
                for fi74 := 0; fi74 < 3; fi74++ {
                    v75 := v3[(("foo bar" + (((("hello" + "\n\t") + "a") + ("hello" + "\n\t")) + "foo bar")) + ("\n\t" + (((("\n\t" + (("foo bar" + "a") + ("foo bar" + "foo bar"))) + "") + "a") + "foo bar")))]
                    _ = fi74
                    _ = v75
                }
            }
            _ = v72
        }
    }
    v35 -= (float32((0.1 - 0.0)) - v35)
    if (4294967295 >= 4294967295) {
        v76 := GenS11{F0: -863, F1: true}
        for fi77 := 0; fi77 < 11; fi77++ {
            v4 = append(v4, fi77)
            _ = fi77
        }
        _ = v76
    }
    v35 = float32(3.14159)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
