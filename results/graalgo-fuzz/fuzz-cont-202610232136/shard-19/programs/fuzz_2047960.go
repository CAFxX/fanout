package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 uint; F2 int64 }
type GenS3 struct { F0 uint16; F1 int16 }
type GenS4 struct { F0 uint16; F1 int8; F2 bool }
type GenS5 struct { F0 int8; F1 string; F2 uint8 }
type GenS6 struct { F0 uint8; F1 int8 }

func main() {
    var v1 uint = 1
    var v2 uint = v1
    var v3 float32 = -2.25
    var v4 bool = true
    for fi5 := 0; fi5 < 11; fi5++ {
        fmt.Println(v3)
        var v6 float64 = 0.1
        _ = fi5
        _ = v6
    }
    var v7 int8 = 42
    fmt.Println(v2, v1)
    v4 = v4
    v4 = (false || ((419 >= 9223372036854775807) || false))
    switch 0 {
    case 32768:
        var v8 int64 = 574
        _ = v8
    }
    v1 -= v2
    if (v2 < (v1 / ((v2) | 1))) {
        v4 = v4
        switch -9223372036854775808 {
        case 0:
            v9 := GenS1{F0: (v1 - (v1 + v1))}
            _ = v9
        case 32768:
            var v10 uint64 = 185
            _ = v10
        }
    } else {
        v11 := GenS1{F0: 630}
        v2 -= 65535
        _ = v11
    }
    if false {
        v12 := GenS2{F0: "foo bar", F1: (v2 + v1), F2: -1}
        _ = v12
    }
    var v13 int8 = (v7 | v7)
    if v4 {
        if (v4 || (!(!true))) {
            v13 -= v13
            if (int64(v2) >= 42) {
                v3 = (v3 - (0.0 + v3))
            } else {
                var v14 int8 = 42
                _ = v14
            }
            if true {
                var v15 int32 = 2147483647
                if true {
                    if (!v4) {
                        v16 := GenS3{F0: 0, F1: 0}
                        v4 = v4
                        v2 = (v2 & (v2 + v1))
                        _ = v16
                    }
                    switch 42 {
                    case 127:
                        var v17 int8 = v13
                        _ = v17
                    case 9223372036854775807:
                        v2 -= v1
                    case 2147483647:
                        var v18 int8 = -61
                        _ = v18
                    }
                }
                _ = v15
            } else {
                switch ("hello" + ("a" + "a")) {
                case "foo bar":
                    var v19 float32 = 1.5
                    _ = v19
                case "\n\t":
                    v1 = 238
                default:
                    switch "" {
                    case "foo bar":
                        v20 := []int{-9223372036854775808, 1}
                        _ = v20
                    case "\n\t":
                        var v21 uint32 = 913
                        _ = v21
                    }
                }
                v3 = 0.0
                switch 4294967295 {
                case 2147483647:
                    switch 685 {
                    case 4294967295:
                        v22 := GenS4{F0: 32767, F1: (v13 << ((v2) & 63)), F2: v4}
                        _ = v22
                    case 32767:
                        var v23 int64 = 9223372036854775807
                        _ = v23
                    case 32768:
                        if v4 {
                            var v24 int8 = -32
                            var v25 int16 = -32768
                            _ = v24
                            _ = v25
                        } else {
                            if false {
                                v4 = ((v2 ^ v2) <= 241)
                                v4 = (v7 == (v13 - -128))
                            } else {
                                if true {
                                    var v26 float64 = 1.5
                                    var v27 int = int(v2)
                                    var v28 uint = 1
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                }
                                v2 -= v1
                                for fi29 := 0; fi29 < 13; fi29++ {
                                    var v30 uint8 = 217
                                    var v31 string = ("hello" + "foo bar")
                                    _ = fi29
                                    _ = v30
                                    _ = v31
                                }
                            }
                        }
                    default:
                        for fi32 := 0; fi32 < 12; fi32++ {
                            v33 := GenS5{F0: v7, F1: "a", F2: 1}
                            var v34 float64 = (0.0 / 3.14159)
                            _ = fi32
                            _ = v33
                            _ = v34
                        }
                    }
                case 383:
                    fmt.Println(v2, v1, v4)
                default:
                    switch ("\n\t" + (("a" + ("foo bar" + "")) + "foo bar")) {
                    case "":
                        if false {
                            v7 = v7
                        }
                    default:
                        if (803 < 1) {
                            for fi35 := 0; fi35 < 16; fi35++ {
                                for fi36 := 0; fi36 < 4; fi36++ {
                                    fmt.Println(fi36)
                                    switch ("\n\t" + "foo bar") {
                                    case "\n\t":
                                        if (!(!v4)) {
                                            v2 = v1
                                            switch fi35 {
                                            case 42:
                                                v7 -= 127
                                            case 0:
                                                fmt.Println(fi36)
                                            case -584:
                                                if (42 < 1) {
                                                    v37 := GenS2{F0: "", F1: (^(v1)), F2: 42}
                                                    v38 := []int{(^(fi35)), fi36}
                                                    v2 += v1
                                                    _ = v37
                                                    _ = v38
                                                }
                                                v1 = 0
                                            default:
                                                v13 += 127
                                            }
                                        }
                                    default:
                                        v13 -= (v7 ^ (^(v13)))
                                    }
                                    _ = fi36
                                }
                                v39 := GenS4{F0: 1, F1: v7, F2: v4}
                                _ = fi35
                                _ = v39
                            }
                            switch "\n\t" {
                            case "a":
                                switch ((("\n\t" + "") + (("" + "") + ("\n\t" + "\n\t"))) + "") {
                                case "foo bar":
                                    for fi40 := 0; fi40 < 6; fi40++ {
                                        for fi41 := 0; fi41 < 17; fi41++ {
                                            v13 -= 127
                                            _ = fi41
                                        }
                                        v42 := GenS5{F0: v13, F1: "", F2: 1}
                                        switch "foo bar" {
                                        case "foo bar":
                                            v43 := []int{(fi40 >> ((v2) & 63)), 0, (fi40 * fi40)}
                                            _ = v43
                                        case "hello":
                                            v44 := GenS6{F0: 128, F1: (-(v7))}
                                            _ = v44
                                        }
                                        _ = fi40
                                        _ = v42
                                    }
                                case "a":
                                    v4 = false
                                default:
                                    var v45 float64 = (0.1 * 3.14159)
                                    _ = v45
                                }
                            case "foo bar":
                                v13 -= (v7 >> ((v2) & 63))
                            default:
                                fmt.Println(v2, v7)
                            }
                        }
                    }
                }
            }
        }
    } else {
        var v46 uint64 = 2
        _ = v46
    }
    v13 -= v13
    fmt.Println(v4, v7)
    v47 := GenS3{F0: 638, F1: 0}
    var v48 uint8 = 178
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v7
}
