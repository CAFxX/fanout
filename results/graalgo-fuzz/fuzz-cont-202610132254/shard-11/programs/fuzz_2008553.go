package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int }
type GenS2 struct { F0 uint64; F1 int64; F2 string }
type GenS3 struct { F0 uint8; F1 int32; F2 int8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 bool; F1 uint32 }
type GenS7 struct { F0 bool; F1 int32; F2 uint32 }
type GenS8 struct { F0 uint16; F1 uint64; F2 uint16 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int16; F1 uint32; F2 bool }
type GenS11 struct { F0 int8; F1 bool }

func main() {
    v1 := []string{"", ("foo bar" + "hello"), ("a" + ""), ("foo bar" + "a")}
    v2 := []uint32{2, 925, 2}
    var v3 int64 = -946
    if len(v1) > 0 {
        v1[0] = ("hello" + "")
    }
    var v4 int32 = 0
    if (1e10 != 1.5) {
        var v5 int32 = v4
        _ = v5
    } else {
        var v6 uint64 = 18446744073709551615
        _ = v6
    }
    if false {
        if true {
            for fi7 := 0; fi7 < 19; fi7++ {
                v4 -= v4
                var v8 uint32 = 42
                _ = len(v2)
                _ = fi7
                _ = v8
            }
            var v9 uint16 = 2
            _ = v9
        }
        v10 := GenS1{F0: 2, F1: 235, F2: -907}
        var v11 int16 = 0
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 13; fi12++ {
        v4 += (v4 / ((0) | 1))
        v4 += (v4 >> ((2) & 63))
        for fi13 := 0; fi13 < 17; fi13++ {
            v4 = (v4 % ((v4) | 1))
            switch "\n\t" {
            case "a":
                fmt.Println(v4)
            default:
                v14 := GenS2{F0: 303, F1: v3, F2: "\n\t"}
                _ = v14
            }
            _ = fi13
        }
        _ = fi12
    }
    switch "a" {
    case "foo bar":
        fmt.Println(v4, v4)
    }
    v15 := GenS3{F0: 42, F1: (^(v4)), F2: 1}
    if len(v1) > 0 {
        v1[0] = ("a" + "a")
    }
    for fi16 := 0; fi16 < 18; fi16++ {
        for fi17 := 0; fi17 < 17; fi17++ {
            v18 := GenS2{F0: 2, F1: (v3 << ((769) & 63)), F2: "\n\t"}
            v19 := GenS4{F0: 2}
            if true {
                var v20 int = -2147483648
                v20 -= 463
                var v21 string = ""
                if (42 == 0) {
                    v3 = (v3 + (v3 >> ((65535) & 63)))
                }
                _ = v20
                _ = v21
            } else {
                v4 = v4
                var v22 uint16 = 655
                _ = v22
            }
            _ = fi17
            _ = v18
            _ = v19
        }
        var v23 uint16 = 0
        fmt.Println(v3, fi16)
        _ = fi16
        _ = v23
    }
    for fi24 := 0; fi24 < 19; fi24++ {
        if (-128 != 0) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            if ((fi24 - 32767) <= fi24) {
                switch (fi24 + (fi24 | (fi24 << ((1) & 63)))) {
                case 42:
                    if true {
                        fmt.Println(fi24, v3, fi24)
                    }
                default:
                    v1 = append(v1, ("" + "\n\t"))
                }
                if (!(4294967295 > 42)) {
                    for fi25 := 0; fi25 < 18; fi25++ {
                        if true {
                            v3 = 1
                        } else {
                            v3 += -9223372036854775808
                        }
                        v26 := GenS5{F0: v4}
                        _ = fi25
                        _ = v26
                    }
                } else {
                    v27 := GenS1{F0: 0, F1: 1, F2: 127}
                    v3 -= 42
                    v2 = append(v2, 570)
                    _ = v27
                }
                v3 += v3
            }
        }
        _ = fi24
    }
    if false {
        fmt.Println(v3)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi28 := 0; fi28 < 10; fi28++ {
            if false {
                if false {
                    for fi29 := 0; fi29 < 19; fi29++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if ((-2.25 * 0.0) < 1.5) {
                            if len(v2) > 0 {
                                v2[0] = 42
                            }
                            if (419 >= 4294967295) {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                        }
                        var v30 int32 = (v4 % ((v4) | 1))
                        _ = fi29
                        _ = v30
                    }
                    if ((v4 ^ 1) != v4) {
                        v3 -= v3
                        v4 = 0
                        var v31 float32 = -2.25
                        _ = v31
                    } else {
                        for fi32 := 0; fi32 < 14; fi32++ {
                            if false {
                                fmt.Println(fi32, v4, fi28)
                                v33 := GenS4{F0: 42}
                                var v34 uint32 = 429
                                _ = v33
                                _ = v34
                            }
                            v4 += v4
                            if len(v1) > 0 {
                                v1[0] = ("\n\t" + ((("hello" + "a") + "\n\t") + "hello"))
                            }
                            var v35 int = -1
                            if ((-(fi28)) <= fi28) {
                                if len(v2) > 0 {
                                    v2[0] = 2147483647
                                }
                                v36 := GenS2{F0: 4294967295, F1: (v3 << ((42) & 63)), F2: ((("hello" + "\n\t") + "foo bar") + "hello")}
                                fmt.Println(v4)
                                fmt.Println(v4)
                                var v37 uint = 42
                                v38 := GenS6{F0: true, F1: 2147483647}
                                _ = v36
                                _ = v37
                                _ = v38
                            } else {
                                if true {
                                    var v39 uint32 = 46
                                    v40 := GenS6{F0: (!false), F1: v39}
                                    _ = v39
                                    _ = v40
                                }
                                if (69 != 128) {
                                    if len(v1) > 0 {
                                        v1[0] = ""
                                    }
                                    switch "hello" {
                                    case "foo bar":
                                        if (-2.25 != 3.14159) {
                                            if (18446744073709551615 == 224) {
                                                if (!(!(127 != int(v4)))) {
                                                    var v41 float64 = -2.25
                                                    v42 := GenS1{F0: 65535, F1: 42, F2: (-(fi28))}
                                                    var v43 float64 = float64(((float32((3.14159 * v41)) - 0.1) + 0.0))
                                                    _ = v41
                                                    _ = v42
                                                    _ = v43
                                                }
                                            } else {
                                                v35 -= (fi32 - v35)
                                                v35 = (fi28 % ((32767) | 1))
                                                if ((^(v3)) == (v3 << ((132) & 63))) {
                                                    if len(v1) > 0 {
                                                        v1[0] = "foo bar"
                                                    }
                                                    var v44 uint8 = 102
                                                    var v45 uint16 = 65535
                                                    _ = v44
                                                    _ = v45
                                                } else {
                                                    v3 += v3
                                                    v35 -= fi28
                                                }
                                                v3 = (v3 + (^(v3)))
                                                v46 := GenS2{F0: 4294967295, F1: (v3 | (v3 + v3)), F2: ("" + "hello")}
                                                fmt.Println(v4, v4)
                                                _ = v46
                                            }
                                        } else {
                                            switch "a" {
                                            case "foo bar":
                                                var v47 int16 = -1
                                                fmt.Println(fi28)
                                                _ = v47
                                            default:
                                                v48 := GenS7{F0: true, F1: -913, F2: 2147483647}
                                                _ = v48
                                            }
                                        }
                                    default:
                                        var v49 uint8 = 0
                                        _ = v49
                                    }
                                }
                            }
                            _ = fi32
                            _ = v35
                        }
                        if len(v1) > 0 {
                            v1[0] = ("\n\t" + ("" + "hello"))
                        }
                    }
                    if false {
                        fmt.Println(v4)
                        v50 := GenS8{F0: 106, F1: 129, F2: 156}
                        _ = v50
                    }
                } else {
                    if true {
                        switch (fi28 >> ((42) & 63)) {
                        case -520:
                            var v51 int = (-(fi28))
                            _ = v51
                        case -461:
                            switch fi28 {
                            case 215:
                                var v52 int = fi28
                                _ = v52
                            case 32767:
                                for fi53 := 0; fi53 < 6; fi53++ {
                                    if true {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    } else {
                                        var v54 int = fi28
                                        v55 := GenS5{F0: v4}
                                        v54 = v54
                                        _ = v54
                                        _ = v55
                                    }
                                    var v56 int = fi53
                                    _ = fi53
                                    _ = v56
                                }
                            }
                        case 0:
                            v3 -= 1
                        }
                        for fi57 := 0; fi57 < 8; fi57++ {
                            v58 := GenS6{F0: false, F1: 398}
                            v59 := GenS9{F0: 0}
                            _ = fi57
                            _ = v58
                            _ = v59
                        }
                    } else {
                        fmt.Println(v3)
                        var v60 bool = true
                        if v60 {
                            v61 := GenS1{F0: 2, F1: 32767, F2: fi28}
                            _ = len(v2)
                            v2 = append(v2, 987)
                            _ = v61
                        }
                        _ = v60
                    }
                    var v62 string = (("a" + "a") + "")
                    _ = v62
                }
            } else {
                fmt.Println(v4)
                fmt.Println(v3, v3, v3)
                v2 = append(v2, 42)
            }
            _ = fi28
        }
    }
    var v63 string = ("" + ("a" + "a"))
    v3 = (v3 ^ (^(v3)))
    switch (v63 + "\n\t") {
    case "a":
        switch v63 {
        case "foo bar":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        default:
            for fi64 := 0; fi64 < 7; fi64++ {
                for fi65 := 0; fi65 < 16; fi65++ {
                    var v66 uint8 = 184
                    if true {
                        v67 := GenS10{F0: -1, F1: 65535, F2: ((v66 < (v66 | (v66 * 4))) && (true || true))}
                        fmt.Println(v3, v3)
                        v68 := GenS5{F0: v4}
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = v67
                        _ = v68
                    } else {
                        v69 := GenS2{F0: 2, F1: 1, F2: "\n\t"}
                        fmt.Println(v63)
                        _ = v69
                    }
                    _ = fi65
                    _ = v66
                }
                _ = fi64
            }
        }
    case "":
        if (!(!(1 >= 42))) {
            var v70 int8 = 59
            _ = v70
        }
    default:
        for fi71 := 0; fi71 < 19; fi71++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v72 := GenS8{F0: 65535, F1: 1, F2: 255}
            switch (v63 + v63) {
            case "foo bar":
                for fi73 := 0; fi73 < 14; fi73++ {
                    v3 = -490
                    switch (fi73 >> ((620) & 63)) {
                    case 953:
                        v63 = "hello"
                    default:
                        if false {
                            var v74 float32 = -2.25
                            var v75 uint = 2
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = v74
                            _ = v75
                        }
                    }
                    _ = fi73
                }
            case "hello":
                for fi76 := 0; fi76 < 13; fi76++ {
                    fmt.Println(fi76)
                    _ = fi76
                }
            default:
                v77 := GenS11{F0: 0, F1: true}
                _ = v77
            }
            _ = fi71
            _ = v72
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v63
}
