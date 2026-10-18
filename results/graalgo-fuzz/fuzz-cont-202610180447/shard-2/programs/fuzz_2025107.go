package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 int16; F2 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS5 struct { F0 int; F1 uint64 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint16 }

func main() {
    var v1 uint8 = 255
    var v2 uint16 = 0
    var v3 int64 = 1
    var v4 float32 = 1.5
    v3 = 64
    switch 0 {
    case 0:
        v5 := GenS1{F0: 42}
        _ = v5
    case -9223372036854775808:
        if ((0.1 * float64(v4)) >= 0.1) {
            v6 := GenS2{F0: 0, F1: 42, F2: 211}
            if false {
                var v7 uint8 = (v1 ^ (v1 << ((42) & 63)))
                for fi8 := 0; fi8 < 11; fi8++ {
                    var v9 int8 = 1
                    _ = fi8
                    _ = v9
                }
                v10 := []int{-119, int(v7), 2147483647}
                _ = v10
                _ = v7
            }
            if (!false) {
                if false {
                    fmt.Println(v3, v3)
                }
                if (393 < 489) {
                    v11 := GenS2{F0: 1, F1: -154, F2: 9223372036854775807}
                    v12 := GenS3{F0: (v3 - (v3 * (v3 / ((9223372036854775807) | 1))))}
                    _ = v11
                    _ = v12
                } else {
                    v13 := []int{0, 32767}
                    v14 := GenS4{F0: 2, F1: (false && true), F2: (v2 - v2)}
                    _ = v13
                    _ = v14
                }
                if true {
                    switch ((("\n\t" + "hello") + ("foo bar" + "a")) + "a") {
                    case "a":
                        v2 -= v2
                    case "foo bar":
                        v15 := []int{32768, 704, 32767}
                        _ = v15
                    default:
                        v3 -= 1
                    }
                    fmt.Println(v4, v2, v3)
                }
            }
            _ = v6
        }
    default:
        fmt.Println(v4, v4)
    }
    var v16 float64 = (-2.25 * 0.0)
    fmt.Println(v4, v16, v2)
    if (false && (false || ((!(793 >= 1)) && true))) {
        v1 += (v1 & 29)
        if (42 < (v1 | v1)) {
            v17 := GenS1{F0: 2147483647}
            _ = v17
        } else {
            v18 := GenS3{F0: v3}
            v19 := []int{510, -1}
            _ = v18
            _ = v19
        }
        switch 127 {
        case 4294967295:
            v1 += v1
        case -1:
            for fi20 := 0; fi20 < 7; fi20++ {
                if (1.5 <= v16) {
                    v21 := []int{}
                    _ = v21
                }
                var v22 bool = ((v1 * v1) < v1)
                _ = fi20
                _ = v22
            }
        case 0:
            for fi23 := 0; fi23 < 14; fi23++ {
                fmt.Println(v4, v1)
                _ = fi23
            }
        default:
            if true {
                fmt.Println(v3)
                var v24 uint32 = 680
                _ = v24
            } else {
                v25 := GenS3{F0: 9223372036854775807}
                _ = v25
            }
        }
    }
    switch ("hello" + "\n\t") {
    case "a":
        for fi26 := 0; fi26 < 20; fi26++ {
            v2 = 2
            _ = fi26
        }
    case "\n\t":
        v27 := GenS3{F0: v3}
        _ = v27
    default:
        v28 := GenS4{F0: 2, F1: false, F2: uint16(v1)}
        _ = v28
    }
    v29 := []int{}
    v3 = int64(v2)
    v16 -= 1e10
    var v30 string = "foo bar"
    for fi31 := 0; fi31 < 5; fi31++ {
        switch ((("foo bar" + v30) + v30) + (v30 + "hello")) {
        case "":
            v32 := GenS5{F0: fi31, F1: 42}
            _ = v32
        case "a":
            if (true && (-602 == int16(v2))) {
                var v33 int16 = 32767
                v1 -= v1
                fmt.Println(fi31)
                _ = v33
            } else {
                fmt.Println(v3)
            }
        default:
            var v34 int32 = -1
            _ = v34
        }
        if len(v29) >= 2 {
            v29 = v29[1:len(v29)]
        }
        var v35 float32 = (-2.25 / v4)
        _ = fi31
        _ = v35
    }
    if true {
        var v36 int16 = 697
        switch "" {
        case "foo bar":
            var v37 uint = 4294967295
            _ = v37
        case "\n\t":
            v1 -= 0
        default:
            if true {
                var v38 string = v30
                v4 -= 1e10
                for fi39 := 0; fi39 < 19; fi39++ {
                    for fi40 := 0; fi40 < 13; fi40++ {
                        var v41 float32 = v4
                        fmt.Println(v36)
                        _ = fi40
                        _ = v41
                    }
                    v30 = v30
                    if true {
                        if (v16 >= v16) {
                            if (false || (127 <= -128)) {
                                var v42 uint = 18446744073709551615
                                _ = v42
                            }
                        }
                    } else {
                        switch "\n\t" {
                        case "\n\t":
                            switch 42 {
                            case 2147483647:
                                for fi43 := 0; fi43 < 8; fi43++ {
                                    if false {
                                        var v44 uint32 = 2147483647
                                        _ = v44
                                    }
                                    var v45 bool = (v4 != v4)
                                    v38 = "\n\t"
                                    _ = fi43
                                    _ = v45
                                }
                            case 32767:
                                switch (fi39 + fi39) {
                                case 128:
                                    var v46 bool = ((v1 ^ v1) != v1)
                                    _ = v46
                                case -2147483648:
                                    for fi47 := 0; fi47 < 16; fi47++ {
                                        var v48 uint = 787
                                        fmt.Println(fi47, v38)
                                        var v49 int = 2147483647
                                        _ = fi47
                                        _ = v48
                                        _ = v49
                                    }
                                }
                            case 9223372036854775807:
                                switch fi39 {
                                case 32768:
                                    switch "a" {
                                    case "\n\t":
                                        fmt.Println(v3, fi39)
                                    case "foo bar":
                                        fmt.Println(v36, v38, v16)
                                    case "hello":
                                        for fi50 := 0; fi50 < 11; fi50++ {
                                            var v51 uint = 4294967295
                                            _ = fi50
                                            _ = v51
                                        }
                                    }
                                case 1:
                                    var v52 bool = true
                                    _ = v52
                                default:
                                    _ = len(v29)
                                }
                            }
                        case "hello":
                            var v53 bool = (-32768 >= -32768)
                            _ = v53
                        }
                        v54 := GenS2{F0: (-(v36)), F1: v36, F2: 4294967295}
                        _ = v54
                    }
                    _ = fi39
                }
                _ = v38
            }
        }
        v55 := GenS5{F0: 128, F1: 18446744073709551615}
        _ = v36
        _ = v55
    }
    var v56 int32 = 569
    v57 := GenS6{F0: 32767}
    switch 127 {
    case 127:
        v1 = 2
    default:
        if (-32768 != 0) {
            v58 := GenS4{F0: 4294967295, F1: (v30 >= (v30 + ((v30 + "\n\t") + v30))), F2: v2}
            var v59 int = 2147483647
            v1 -= 1
            _ = v58
            _ = v59
        }
    }
    var v60 int = 128
    if false {
        if false {
            if false {
                var v61 int16 = -1
                v4 = v4
                v62 := GenS5{F0: (v60 | (v60 * v60)), F1: 0}
                _ = v61
                _ = v62
            } else {
                if len(v29) >= 2 {
                    v29 = v29[1:len(v29)]
                }
            }
            v63 := GenS7{F0: (v2 + v2)}
            v29 = append(v29, 827)
            _ = v63
        }
        fmt.Println(v2)
        for fi64 := 0; fi64 < 10; fi64++ {
            switch ("foo bar" + ("\n\t" + ("hello" + "a"))) {
            case "\n\t":
                var v65 uint64 = 0
                _ = v65
            case "foo bar":
                switch (fi64 ^ (fi64 * (v60 >> ((65535) & 63)))) {
                case 128:
                    for fi66 := 0; fi66 < 18; fi66++ {
                        v56 = v56
                        fmt.Println(v3, fi64)
                        if len(v29) >= 2 {
                            v29 = v29[1:len(v29)]
                        }
                        _ = fi66
                    }
                default:
                    var v67 int64 = v3
                    _ = v67
                }
            default:
                for fi68 := 0; fi68 < 17; fi68++ {
                    v1 = (v1 + (v1 << ((879) & 63)))
                    v4 -= v4
                    if false {
                        v60 = (fi68 >> ((42) & 63))
                        var v69 uint32 = 23
                        _ = v69
                    }
                    _ = fi68
                }
            }
            var v70 int16 = 868
            _ = fi64
            _ = v70
        }
    } else {
        for fi71 := 0; fi71 < 7; fi71++ {
            if len(v29) > 0 {
                v29[0] = v60
            }
            v16 = float64(v4)
            fmt.Println(v16)
            _ = fi71
        }
        var v72 int32 = v56
        v1 = 153
        _ = v72
    }
    if false {
        var v73 int8 = 42
        if len(v29) > 0 {
            v29[0] = v60
        }
        _ = v73
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v56:", v56)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v56
    _ = v57
    _ = v60
}
