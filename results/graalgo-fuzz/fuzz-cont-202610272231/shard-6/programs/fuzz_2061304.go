package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 bool; F1 uint8; F2 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 string }
type GenS6 struct { F0 string; F1 int8 }
type GenS7 struct { F0 uint64; F1 uint64 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint8; F1 int32 }
type GenS10 struct { F0 int; F1 int16; F2 bool }
func genFn0() uint8 {
    var v4 int = -81
    if true {
        if true {
            fmt.Println(v4)
            var v5 float64 = ((((3.14159 / 0.1) - -2.25) * 0.0) - float64((1.5 / 1.0)))
            v6 := GenS1{F0: (^(v4)), F1: 111}
            _ = v5
            _ = v6
        } else {
            var v7 uint8 = 173
            _ = v7
        }
    }
    switch "hello" {
    case "\n\t":
        v8 := GenS2{F0: 42, F1: 2147483648}
        _ = v8
    case "a":
        switch (v4 << ((255) & 63)) {
        case -1:
            var v9 uint16 = 255
            _ = v9
        case -2147483648:
            for fi10 := 0; fi10 < 16; fi10++ {
                var v11 uint16 = 625
                fmt.Println(fi10)
                var v12 int8 = 127
                _ = fi10
                _ = v11
                _ = v12
            }
        }
    default:
        for fi13 := 0; fi13 < 11; fi13++ {
            v4 -= fi13
            _ = fi13
        }
    }
    var v14 int32 = 0
    _ = v14
    _ = v4
    return 255
}

func main() {
    var v1 int = 127
    v2 := []uint32{0}
    v3 := []string{"\n\t", ((("hello" + "hello") + ("\n\t" + ("a" + "hello"))) + "hello"), "foo bar"}
    v1 = (v1 * (v1 - v1))
    v15 := GenS1{F0: (^(v1)), F1: 255}
    var v16 int64 = -1
    if (0 >= 969) {
        var v17 uint8 = 59
        if false {
            var v18 uint8 = (v17 >> ((1) & 63))
            _ = v18
        }
        _ = v17
    }
    if true {
        v1 += (v1 - (v1 % ((v1) | 1)))
    } else {
        var v19 float32 = -2.25
        var v20 bool = true
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        fmt.Println(v16, v16)
        for fi22 := 0; fi22 < 13; fi22++ {
            var v23 bool = (!(0 <= 42))
            if v23 {
                if v23 {
                    var v24 int16 = 32767
                    if v23 {
                        fmt.Println(fi21, v23)
                    } else {
                        v25 := GenS1{F0: fi21, F1: 4294967295}
                        _ = v25
                    }
                    _ = v24
                }
                var v26 uint64 = 42
                if (-528 < 563) {
                    if v23 {
                        v16 = (^(v16))
                        if v23 {
                            v16 += (^(v16))
                            if v23 {
                                v16 = 9223372036854775807
                                var v27 uint16 = 0
                                if (953 <= 63) {
                                    v27 = 1
                                    v26 = 18446744073709551615
                                }
                                _ = v27
                            } else {
                                var v28 float64 = float64(-2.25)
                                if v23 {
                                    v26 -= (v26 + 833)
                                    var v29 bool = v23
                                    v26 = v26
                                    _ = v29
                                } else {
                                    _ = len(v3)
                                }
                                v30 := GenS2{F0: 42, F1: 65535}
                                _ = v28
                                _ = v30
                            }
                        } else {
                            v31 := GenS3{F0: (-128 == 1), F1: genFn0(), F2: genFn0()}
                            _ = v31
                        }
                    } else {
                        v32 := GenS4{F0: 4294967295}
                        v33 := GenS5{F0: "hello"}
                        _ = v32
                        _ = v33
                    }
                    if v23 {
                        if (!v23) {
                            v16 += (-(v16))
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v34 uint16 = uint16(v26)
                            v1 += 135
                            v35 := GenS6{F0: "\n\t", F1: -105}
                            _ = v34
                            _ = v35
                        } else {
                            v23 = false
                        }
                        _ = len(v3)
                        if v23 {
                            v36 := GenS7{F0: 711, F1: v26}
                            v37 := GenS7{F0: 1, F1: v26}
                            _ = v36
                            _ = v37
                        }
                    } else {
                        var v38 uint8 = 1
                        _ = v38
                    }
                    var v39 float64 = 3.14159
                    _ = v39
                } else {
                    v1 += (v1 ^ v1)
                    switch "a" {
                    case "\n\t":
                        v26 -= (v26 + v26)
                    default:
                        v1 = v1
                    }
                }
                _ = v26
            } else {
                if (v23 && false) {
                    v23 = (1000 < 42)
                    var v40 bool = false
                    _ = v40
                }
                fmt.Println(v16, v23)
                fmt.Println(fi22)
            }
            _ = fi22
            _ = v23
        }
        fmt.Println(v16)
        _ = fi21
    }
    switch 127 {
    case 4294967295:
        fmt.Println(v1, v1)
    case 2147483647:
        if true {
            v1 -= v1
            if false {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                switch "foo bar" {
                case "":
                    switch v1 {
                    case 32768:
                        var v41 int16 = 529
                        _ = v41
                    }
                }
            } else {
                v16 += (^(v16))
                v1 = v1
                for fi42 := 0; fi42 < 13; fi42++ {
                    for fi43 := 0; fi43 < 14; fi43++ {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi43
                    }
                    var v44 string = "foo bar"
                    if true {
                        if true {
                            var v45 float32 = 0.0
                            _ = v45
                        }
                        if (!true) {
                            v44 = (v44 + (v44 + (v44 + "a")))
                        } else {
                            var v46 string = v44
                            switch v1 {
                            case 165:
                                v44 = v46
                            case 42:
                                v47 := GenS8{F0: fi42}
                                _ = v47
                            case 128:
                                var v48 int64 = v16
                                _ = v48
                            default:
                                _ = len(v2)
                            }
                            switch v46 {
                            case "hello":
                                fmt.Println(fi42, v1, v46)
                            case "\n\t":
                                _ = len(v2)
                            case "a":
                                v1 = 0
                            default:
                                if (0.1 != float32((3.14159 - 0.1))) {
                                    switch (v1 + 9223372036854775807) {
                                    case 128:
                                        for fi49 := 0; fi49 < 5; fi49++ {
                                            if len(v2) > 0 {
                                                v2[0] = 2147483647
                                            }
                                            var v50 float64 = (1.5 * 0.0)
                                            fmt.Println(v16)
                                            _ = fi49
                                            _ = v50
                                        }
                                    default:
                                        if ((((((v44 + (("a" + "a") + ("\n\t" + v44))) + (v46 + "hello")) + v44) + v46) != "a") && true) {
                                            switch v1 {
                                            case 89:
                                                v3 = append(v3, v46)
                                            }
                                        }
                                    }
                                    v51 := GenS5{F0: "hello"}
                                    fmt.Println(v1)
                                    _ = v51
                                }
                            }
                            _ = v46
                        }
                        switch v44 {
                        case "a":
                            var v52 uint32 = 24
                            _ = v52
                        default:
                            v53 := GenS1{F0: v1, F1: 243}
                            _ = v53
                        }
                    }
                    _ = fi42
                    _ = v44
                }
            }
            switch ("a" + "hello") {
            case "hello":
                var v54 int32 = 334
                _ = v54
            case "":
                fmt.Println(v1)
            case "\n\t":
                var v55 int32 = 475
                _ = v55
            }
        }
    case 32768:
        v56 := GenS2{F0: 32767, F1: 1}
        _ = v56
    }
    if false {
        if (42 <= 0) {
            v16 -= v16
        }
        var v57 int64 = v16
        for fi58 := 0; fi58 < 19; fi58++ {
            for fi59 := 0; fi59 < 4; fi59++ {
                if (-23 == -78) {
                    var v60 float32 = (float32((-2.25 + (1e10 * ((0.1 / 1.0) * 1e10)))) / -2.25)
                    switch 32767 {
                    case 4294967295:
                        var v61 string = "hello"
                        _ = v61
                    case 128:
                        if true {
                            var v62 int32 = -2147483648
                            _ = v62
                        } else {
                            var v63 int = (fi58 >> ((29) & 63))
                            v64 := GenS9{F0: genFn0(), F1: 42}
                            v16 += v16
                            v63 -= -548
                            v60 += v60
                            _ = v63
                            _ = v64
                        }
                    case 65:
                        switch "\n\t" {
                        case "foo bar":
                            v60 -= float32(-2.25)
                        case "a":
                            v65 := GenS6{F0: "\n\t", F1: 1}
                            _ = v65
                        }
                    default:
                        v66 := GenS6{F0: "", F1: 127}
                        _ = v66
                    }
                    _ = v60
                } else {
                    v57 -= v57
                    if len(v3) > 0 {
                        v3[0] = "\n\t"
                    }
                    var v67 string = "a"
                    _ = v67
                }
                _ = fi59
            }
            v16 -= 42
            var v68 float64 = 3.14159
            _ = fi58
            _ = v68
        }
        _ = v57
    }
    for fi69 := 0; fi69 < 6; fi69++ {
        if len(v2) > 0 {
            v2[0] = 0
        }
        if false {
            switch ((("a" + "foo bar") + "\n\t") + ("\n\t" + (("" + "a") + ("" + "a")))) {
            case "foo bar":
                if false {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v16 -= v16
                    v16 = (v16 & v16)
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if true {
                        if (!(true || true)) {
                            v16 += v16
                            v1 = v1
                        }
                    }
                }
            case "":
                _ = len(v3)
            }
            v70 := GenS2{F0: 0, F1: 2}
            _ = v70
        }
        var v71 int8 = 32
        _ = fi69
        _ = v71
    }
    var v72 float64 = (3.14159 / 1e10)
    switch "foo bar" {
    case "\n\t":
        v72 -= float64(1.5)
    case "foo bar":
        fmt.Println(v72, v72, v72)
    }
    for fi73 := 0; fi73 < 16; fi73++ {
        var v74 uint = 18446744073709551615
        for fi75 := 0; fi75 < 5; fi75++ {
            v76 := GenS1{F0: fi73, F1: 255}
            v77 := GenS10{F0: (fi73 & (-(fi73))), F1: 42, F2: (19 > 32767)}
            _ = fi75
            _ = v76
            _ = v77
        }
        _ = fi73
        _ = v74
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v72
}
