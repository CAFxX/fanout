package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 uint32; F1 uint64; F2 uint64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int8; F1 uint8 }
type GenS7 struct { F0 uint32; F1 string; F2 bool }
type GenS8 struct { F0 int; F1 int }
type GenS9 struct { F0 int8; F1 int16 }

func main() {
    var v1 float64 = 1.5
    var v2 uint8 = 1
    v3 := []int{573, -2147483648, -377, 32767}
    v4 := GenS1{F0: false, F1: false}
    if ((v2 <= 2) && true) {
        _ = len(v3)
    }
    v3 = append(v3, 1)
    for fi5 := 0; fi5 < 14; fi5++ {
        var v6 uint8 = v2
        v2 = 122
        if len(v3) > 0 {
            v3[0] = (fi5 / ((fi5) | 1))
        }
        _ = fi5
        _ = v6
    }
    var v7 uint = 987
    v3 = append(v3, -158)
    for fi8 := 0; fi8 < 18; fi8++ {
        v7 += (^(v7))
        _ = fi8
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        if false {
            v2 = (^(v2))
        } else {
            v3 = append(v3, -829)
        }
        _ = fi9
    }
    if (!(-1 == -57)) {
        for fi10 := 0; fi10 < 6; fi10++ {
            var v11 int = (fi10 % ((fi10) | 1))
            var v12 uint = (v7 | 4294967295)
            v13 := GenS2{F0: 0, F1: -891}
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
    } else {
        for fi14 := 0; fi14 < 16; fi14++ {
            fmt.Println(v2, v1)
            _ = fi14
        }
    }
    v1 -= 1.5
    v15 := GenS2{F0: 1000, F1: 0}
    switch ("\n\t" + ("\n\t" + "\n\t")) {
    case "a":
        for fi16 := 0; fi16 < 19; fi16++ {
            v3 = append(v3, -9223372036854775808)
            v3 = append(v3, fi16)
            for fi17 := 0; fi17 < 7; fi17++ {
                switch fi17 {
                case -1:
                    fmt.Println(v1, v2, v2)
                case -9223372036854775808:
                    if false {
                        var v18 string = "foo bar"
                        var v19 uint8 = (v2 & v2)
                        v7 = v7
                        switch "foo bar" {
                        case "a":
                            v20 := GenS3{F0: v2, F1: 42}
                            _ = v20
                        case "hello":
                            if len(v3) > 0 {
                                v3[0] = (fi16 * (^(fi16)))
                            }
                            var v21 int16 = -32768
                            fmt.Println(fi16, v1, fi17)
                            _ = v21
                        case "foo bar":
                            var v22 int64 = -1
                            _ = v22
                        }
                        _ = v18
                        _ = v19
                    } else {
                        _ = len(v3)
                        var v23 float32 = 0.1
                        switch fi16 {
                        case -2147483648:
                            fmt.Println(fi17, fi17, v23)
                        }
                        _ = v23
                    }
                default:
                    if (1 == 1000) {
                        v24 := GenS4{F0: 0, F1: 2, F2: 1}
                        v25 := GenS2{F0: -934, F1: -1}
                        v26 := GenS5{F0: 695}
                        _ = v24
                        _ = v25
                        _ = v26
                    } else {
                        v2 -= 255
                        v2 = 197
                        v27 := GenS2{F0: 42, F1: -1}
                        _ = v27
                    }
                }
                var v28 int16 = int16(fi16)
                v2 += v2
                v29 := GenS4{F0: 0, F1: 4294967295, F2: 18446744073709551615}
                _ = fi17
                _ = v28
                _ = v29
            }
            _ = fi16
        }
    case "foo bar":
        v30 := GenS6{F0: 100, F1: uint8(v7)}
        _ = v30
    default:
        for fi31 := 0; fi31 < 9; fi31++ {
            v2 += (v2 << (((v7 / ((v7) | 1))) & 63))
            if (!(1 < 4294967295)) {
                if (176 <= int16(v2)) {
                    v7 += v7
                    var v32 uint = (^(v7))
                    _ = v32
                }
                var v33 bool = false
                for fi34 := 0; fi34 < 6; fi34++ {
                    if true {
                        var v35 uint8 = v2
                        var v36 int64 = 42
                        switch ("hello" + "a") {
                        case "a":
                            v33 = (!((-2.25 + (1e10 / 1.0)) <= (1e10 * ((((3.14159 + ((1.5 + 1e10) + 0.0)) / float32(1e10)) - 0.0) - 1.5))))
                        }
                        _ = v35
                        _ = v36
                    }
                    var v37 uint64 = 1
                    fmt.Println(v37, v1, v2)
                    _ = fi34
                    _ = v37
                }
                _ = v33
            }
            if true {
                v38 := GenS4{F0: 204, F1: 18446744073709551615, F2: 2}
                for fi39 := 0; fi39 < 17; fi39++ {
                    var v40 bool = false
                    if true {
                        switch "a" {
                        case "a":
                            switch "foo bar" {
                            case "foo bar":
                                v3 = append(v3, 5)
                            case "\n\t":
                                v2 = v2
                            default:
                                fmt.Println(v7, v2, fi31)
                            }
                        case "":
                            fmt.Println(v40, v7)
                        }
                        switch ("" + "\n\t") {
                        case "":
                            if ((!false) || ((0.1 / 1e10) >= -2.25)) {
                                if (1 <= 18446744073709551615) {
                                    if (-32768 != -32768) {
                                        switch fi31 {
                                        case -1:
                                            v41 := GenS7{F0: 2147483648, F1: "foo bar", F2: v40}
                                            _ = v41
                                        }
                                        v3 = append(v3, (fi31 % ((32767) | 1)))
                                        var v42 float32 = ((3.14159 * (1.5 + ((0.0 / 1.5) + 0.1))) - 0.1)
                                        _ = v42
                                    }
                                }
                            }
                        default:
                            var v43 int32 = -2147483648
                            _ = v43
                        }
                        v40 = v40
                        v44 := GenS8{F0: fi39, F1: (fi39 / ((fi39) | 1))}
                        _ = v44
                    } else {
                        var v45 int8 = 127
                        _ = v45
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi39
                    _ = v40
                }
                _ = v38
            } else {
                v46 := GenS9{F0: 0, F1: 612}
                fmt.Println(fi31, v1)
                _ = v46
            }
            _ = fi31
        }
    }
    var v47 int = 9223372036854775807
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if len(v3) > 0 {
        v3[0] = 128
    }
    v1 += float64((1e10 * ((0.0 + (float32((((-2.25 + (float64((1e10 * 3.14159)) - ((0.1 * (v1 + v1)) + (v1 - v1)))) * (float64(3.14159) / -2.25)) - v1)) + -2.25)) / (float32((v1 - 0.1)) + float32(1.5)))))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v47:", v47)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v7
}
