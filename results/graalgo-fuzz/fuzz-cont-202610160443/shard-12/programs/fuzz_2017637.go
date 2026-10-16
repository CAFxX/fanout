package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 uint8; F1 int8; F2 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8; F1 int32 }
type GenS6 struct { F0 string; F1 int8 }
type GenS7 struct { F0 int8; F1 string; F2 int8 }
type GenS8 struct { F0 uint8; F1 bool }
type GenS9 struct { F0 int32 }

func main() {
    v1 := []uint32{117, 1, 647}
    v2 := []int{32767, -9223372036854775808, 2147483647}
    var v3 uint8 = 1
    v4 := make(map[int]bool)
    switch "\n\t" {
    case "foo bar":
        if len(v2) > 0 {
            v2[0] = 360
        }
    }
    fmt.Println(v3)
    v5 := GenS1{F0: -9223372036854775808, F1: 1}
    var v6 float32 = 1.5
    switch "" {
    case "hello":
        if true {
            v4[-590] = (v3 <= (v3 % (((v3 + (v3 ^ v3))) | 1)))
            fmt.Println(v3)
        }
    case "a":
        switch "foo bar" {
        case "foo bar":
            var v7 uint = 255
            _ = v7
        case "":
            var v8 float32 = v6
            _ = v8
        }
    default:
        v9 := GenS1{F0: 0, F1: 65535}
        _ = v9
    }
    if false {
        for fi10 := 0; fi10 < 12; fi10++ {
            v11 := GenS2{F0: (v3 & v3), F1: -44, F2: 4294967295}
            v6 += v6
            if (9223372036854775807 < 97) {
                for fi12 := 0; fi12 < 14; fi12++ {
                    var v13 int8 = 42
                    _ = fi12
                    _ = v13
                }
                fmt.Println(fi10, v3)
            } else {
                v14 := GenS3{F0: fi10}
                _ = v14
            }
            _ = fi10
            _ = v11
        }
        for fi15 := 0; fi15 < 11; fi15++ {
            switch ("a" + "") {
            case "hello":
                for fi16 := 0; fi16 < 1; fi16++ {
                    v17 := GenS1{F0: -9223372036854775808, F1: 65535}
                    var v18 int8 = 127
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            case "foo bar":
                fmt.Println(fi15, v3, v6)
            case "":
                for fi19 := 0; fi19 < 2; fi19++ {
                    if (4294967295 > 1) {
                        v6 -= v6
                    }
                    _ = fi19
                }
            }
            if true {
                if false {
                    v6 = (v6 - (0.0 + 0.1))
                    if true {
                        v6 = v6
                        v6 -= (1e10 * 0.1)
                    }
                    v20 := GenS4{F0: 281}
                    _ = v20
                }
                for fi21 := 0; fi21 < 19; fi21++ {
                    fmt.Println(fi21, v3, v3)
                    _ = fi21
                }
                var v22 float64 = (0.0 - float64((v6 / 1.5)))
                _ = v22
            }
            v23 := GenS1{F0: 1, F1: 5}
            _ = fi15
            _ = v23
        }
        var v24 uint16 = 677
        _ = v24
    } else {
        for fi25 := 0; fi25 < 9; fi25++ {
            var v26 uint16 = 1
            _ = fi25
            _ = v26
        }
        v3 = v3
    }
    for fi27 := 0; fi27 < 16; fi27++ {
        switch "a" {
        case "a":
            _ = len(v2)
        case "":
            var v28 bool = true
            _ = v28
        case "\n\t":
            _ = len(v4)
        default:
            switch -2147483648 {
            case 1:
                for fi29 := 0; fi29 < 13; fi29++ {
                    v30 := GenS5{F0: v3, F1: 0}
                    fmt.Println(fi29, fi27, v6)
                    _ = fi29
                    _ = v30
                }
            case -449:
                fmt.Println(fi27, fi27, v6)
            }
        }
        v3 = v3
        _ = fi27
    }
    for fi31 := 0; fi31 < 12; fi31++ {
        v32 := GenS5{F0: 128, F1: -2147483648}
        var v33 float64 = 0.0
        _ = fi31
        _ = v32
        _ = v33
    }
    v3 = 240
    if (-1 > -997) {
        for fi34 := 0; fi34 < 13; fi34++ {
            var v35 float64 = float64(float32(0.1))
            switch "a" {
            case "":
                v3 -= (v3 % (((v3 + v3)) | 1))
            case "a":
                var v36 int = fi34
                _ = v36
            case "\n\t":
                v37 := GenS6{F0: "a", F1: -49}
                _ = v37
            }
            _ = fi34
            _ = v35
        }
    } else {
        var v38 int8 = -1
        fmt.Println(v6)
        for fi39 := 0; fi39 < 4; fi39++ {
            for fi40 := 0; fi40 < 13; fi40++ {
                switch (fi39 ^ (fi39 % (((fi39 / (((fi40 / ((128) | 1))) | 1))) | 1))) {
                case 1:
                    var v41 bool = true
                    _ = v41
                case -765:
                    _ = len(v4)
                case 9223372036854775807:
                    var v42 string = "a"
                    var v43 uint8 = 128
                    _ = v42
                    _ = v43
                }
                _ = fi40
            }
            fmt.Println(fi39, v3, v38)
            for fi44 := 0; fi44 < 7; fi44++ {
                if (!false) {
                    v38 = v38
                    var v45 float32 = v6
                    _ = v45
                }
                fmt.Println(v6, v6)
                _ = fi44
            }
            _ = fi39
        }
        _ = v38
    }
    v3 = (^(v3))
    if (-3 == 100) {
        var v46 uint = 198
        if (!true) {
            v3 = (v3 >> ((v46) & 63))
            switch "" {
            case "":
                for fi47 := 0; fi47 < 8; fi47++ {
                    var v48 int64 = 9223372036854775807
                    v6 += v6
                    switch (("foo bar" + "foo bar") + ("" + (("\n\t" + ("foo bar" + "a")) + "a"))) {
                    case "\n\t":
                        v3 = v3
                    case "foo bar":
                        v49 := GenS7{F0: 127, F1: "", F2: 100}
                        _ = v49
                    default:
                        v46 = (v46 ^ v46)
                    }
                    _ = fi47
                    _ = v48
                }
            case "hello":
                v50 := v4[128]
                _ = v50
            case "foo bar":
                v46 += (v46 | 257)
            }
            for fi51 := 0; fi51 < 11; fi51++ {
                v1 = append(v1, 2147483647)
                for fi52 := 0; fi52 < 7; fi52++ {
                    if ((!true) && (false && false)) {
                        switch "" {
                        case "a":
                            fmt.Println(v46)
                        default:
                            if len(v1) > 0 {
                                v1[0] = 791
                            }
                        }
                        var v53 float64 = (1.5 + (0.1 + 0.1))
                        v54 := GenS8{F0: 190, F1: ((true && true) || true)}
                        _ = v53
                        _ = v54
                    } else {
                        fmt.Println(v3, fi51, fi52)
                        v55 := GenS2{F0: 25, F1: 127, F2: v46}
                        _ = v55
                    }
                    v56 := GenS1{F0: 9223372036854775807, F1: uint16(v3)}
                    switch "hello" {
                    case "foo bar":
                        fmt.Println(v6, v6, fi52)
                    default:
                        fmt.Println(v46, v3, v6)
                    }
                    var v57 int8 = 126
                    _ = len(v4)
                    _ = fi52
                    _ = v56
                    _ = v57
                }
                var v58 string = "\n\t"
                _ = fi51
                _ = v58
            }
        } else {
            delete(v4, -1)
        }
        v59 := GenS9{F0: 2147483647}
        _ = v46
        _ = v59
    } else {
        var v60 float64 = (float64(v6) * 3.14159)
        fmt.Println(v3)
        _ = v60
    }
    v3 = (v3 << ((4294967295) & 63))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v61 := GenS2{F0: v3, F1: -1, F2: 1}
    var v62 string = ("foo bar" + "hello")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
    _ = v62
}
