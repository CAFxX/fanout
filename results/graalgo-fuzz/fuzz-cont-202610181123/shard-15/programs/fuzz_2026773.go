package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 uint; F2 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64; F1 int64; F2 int16 }
type GenS5 struct { F0 uint; F1 bool; F2 bool }
type GenS6 struct { F0 int }
type GenS7 struct { F0 bool; F1 uint64; F2 int8 }

func main() {
    v1 := []uint32{0, 215, 304}
    var v2 uint32 = 2147483648
    v3 := []int{0, 575, -1}
    v2 = 2
    switch 1 {
    case 920:
        for fi4 := 0; fi4 < 4; fi4++ {
            for fi5 := 0; fi5 < 16; fi5++ {
                var v6 int8 = 1
                v6 = (v6 + (^(v6)))
                _ = fi5
                _ = v6
            }
            for fi7 := 0; fi7 < 10; fi7++ {
                v2 = (v2 | (^(v2)))
                _ = fi7
            }
            _ = fi4
        }
    }
    fmt.Println(v2)
    for fi8 := 0; fi8 < 4; fi8++ {
        if len(v1) > 0 {
            v1[0] = 42
        }
        _ = fi8
    }
    v9 := GenS1{F0: 26}
    fmt.Println(v2)
    v10 := GenS2{F0: "foo bar", F1: uint(v2), F2: 2147483647}
    switch "" {
    case "foo bar":
        for fi11 := 0; fi11 < 2; fi11++ {
            for fi12 := 0; fi12 < 9; fi12++ {
                v2 -= v2
                _ = fi12
            }
            v13 := GenS1{F0: -32768}
            _ = fi11
            _ = v13
        }
    case "":
        for fi14 := 0; fi14 < 16; fi14++ {
            v2 = v2
            if (1 >= 4294967295) {
                switch "\n\t" {
                case "a":
                    for fi15 := 0; fi15 < 16; fi15++ {
                        var v16 int32 = -365
                        var v17 float64 = 3.14159
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                case "":
                    v18 := GenS2{F0: "\n\t", F1: 865, F2: 42}
                    _ = v18
                case "foo bar":
                    v2 = (v2 * (^(v2)))
                }
                for fi19 := 0; fi19 < 8; fi19++ {
                    if false {
                        v20 := GenS2{F0: "\n\t", F1: 0, F2: -759}
                        v2 += (v2 / ((v2) | 1))
                        v21 := GenS2{F0: "foo bar", F1: 255, F2: -147}
                        if true {
                            if true {
                                v3 = append(v3, fi19)
                                fmt.Println(fi19, fi19, fi19)
                            }
                            v22 := GenS2{F0: ("foo bar" + ""), F1: 255, F2: 42}
                            var v23 bool = true
                            _ = v22
                            _ = v23
                        } else {
                            v2 = v2
                        }
                        _ = v20
                        _ = v21
                    } else {
                        switch "\n\t" {
                        case "\n\t":
                            v2 = (v2 - (v2 >> ((1) & 63)))
                        case "hello":
                            v3 = append(v3, (fi19 / ((9223372036854775807) | 1)))
                        default:
                            v24 := GenS3{F0: 127}
                            _ = v24
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    if (!(0.0 == (-2.25 - float32(3.14159)))) {
                        v2 -= (v2 << ((960) & 63))
                        v2 = v2
                    } else {
                        if true {
                            v25 := GenS1{F0: 0}
                            var v26 uint64 = 644
                            _ = v25
                            _ = v26
                        } else {
                            if true {
                                var v27 int32 = 0
                                var v28 int32 = v27
                                _ = v27
                                _ = v28
                            } else {
                                v29 := GenS3{F0: 1}
                                fmt.Println(v2, fi19, v2)
                                _ = v29
                            }
                        }
                        if false {
                            if (!(!(0 >= 654))) {
                                if len(v3) > 0 {
                                    v3[0] = fi19
                                }
                                v30 := GenS4{F0: 625, F1: 250, F2: 1000}
                                _ = v30
                            } else {
                                _ = len(v3)
                                switch ((((("a" + (("\n\t" + "") + "")) + ("a" + "foo bar")) + "a") + "") + ("" + "\n\t")) {
                                case "\n\t":
                                    v2 += 830
                                default:
                                    v2 -= (^(v2))
                                }
                                v2 -= v2
                            }
                            v2 += v2
                        } else {
                            fmt.Println(fi19, fi19)
                        }
                        v2 = v2
                    }
                    _ = fi19
                }
            }
            _ = fi14
        }
    }
    for fi31 := 0; fi31 < 13; fi31++ {
        for fi32 := 0; fi32 < 3; fi32++ {
            if (fi32 > fi31) {
                var v33 uint32 = (v2 % (((^(v2))) | 1))
                v34 := GenS5{F0: 1, F1: false, F2: (18446744073709551615 != 42)}
                var v35 uint = 502
                _ = v33
                _ = v34
                _ = v35
            }
            switch "\n\t" {
            case "hello":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            default:
                switch fi31 {
                case -1:
                    switch (-(fi32)) {
                    case -854:
                        _ = len(v1)
                    }
                case 9223372036854775807:
                    v2 = v2
                case 1:
                    if (1 != -9223372036854775808) {
                        switch ("\n\t" + "") {
                        case "a":
                            v2 = v2
                        case "\n\t":
                            v2 = (v2 / ((v2) | 1))
                            var v36 int32 = 0
                            _ = v36
                        case "hello":
                            v2 += v2
                        }
                    }
                }
            }
            _ = fi32
        }
        var v37 uint = 4294967295
        if len(v1) > 0 {
            v1[0] = (^(v2))
        }
        _ = fi31
        _ = v37
    }
    if (false || (-562 >= 1)) {
        var v38 uint32 = 888
        _ = v38
    } else {
        fmt.Println(v2)
        v2 = v2
        if true {
            if ("\n\t" > ("" + "foo bar")) {
                switch 0 {
                case 32768:
                    v2 = v2
                }
                v2 -= (v2 / ((612) | 1))
                v39 := GenS4{F0: 42, F1: 622, F2: -489}
                _ = v39
            }
            for fi40 := 0; fi40 < 12; fi40++ {
                var v41 int8 = 36
                if (v2 < (v2 * (v2 % ((42) | 1)))) {
                    for fi42 := 0; fi42 < 7; fi42++ {
                        v41 -= -128
                        var v43 uint64 = 1
                        _ = fi42
                        _ = v43
                    }
                    v41 -= v41
                } else {
                    switch "hello" {
                    case "a":
                        v44 := GenS4{F0: 18446744073709551615, F1: -1, F2: 1}
                        _ = v44
                    case "":
                        fmt.Println(v2, v41)
                    case "hello":
                        v1 = append(v1, (v2 << ((1) & 63)))
                    default:
                        v2 = v2
                    }
                    var v45 int8 = (v41 * 79)
                    _ = v45
                }
                _ = fi40
                _ = v41
            }
        }
    }
    if (!(104 < -115)) {
        var v46 uint32 = v2
        _ = v46
    }
    for fi47 := 0; fi47 < 10; fi47++ {
        if (1e10 > (1.5 / float64((1e10 * -2.25)))) {
            v48 := GenS6{F0: (fi47 + (fi47 & fi47))}
            var v49 uint8 = 128
            var v50 uint = 18446744073709551615
            _ = v48
            _ = v49
            _ = v50
        }
        for fi51 := 0; fi51 < 5; fi51++ {
            v2 = v2
            fmt.Println(v2)
            v52 := GenS7{F0: (!false), F1: 2, F2: -128}
            _ = fi51
            _ = v52
        }
        var v53 uint64 = 4294967295
        _ = fi47
        _ = v53
    }
    fmt.Println(v2)
    var v54 int64 = -9223372036854775808
    var v55 uint64 = 168
    v54 += -9223372036854775808
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v54
    _ = v55
    _ = v9
}
