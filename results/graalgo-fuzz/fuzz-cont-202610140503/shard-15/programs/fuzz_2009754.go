package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 int16 }
type GenS2 struct { F0 uint32; F1 uint16; F2 int }
type GenS3 struct { F0 string; F1 string; F2 uint }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 string; F1 int32; F2 bool }

func main() {
    var v1 uint64 = 0
    var v2 int16 = -32768
    var v3 uint8 = 127
    var v4 uint8 = v3
    fmt.Println(v4)
    v5 := GenS1{F0: 127, F1: (v2 - (v2 / (((v2 ^ (-(v2)))) | 1))), F2: (^(v2))}
    var v6 uint8 = 137
    var v7 uint8 = (v3 >> ((2) & 63))
    var v8 uint64 = (^(v1))
    if ((v7 + (v4 / ((v6) | 1))) > (^(v3))) {
        v7 -= (v6 - 128)
    } else {
        v9 := []int{42, 32767, 1, 1}
        if true {
            switch -2147483648 {
            case 32767:
                if ("hello" == "hello") {
                    if ((false || false) || (true || true)) {
                        v10 := GenS2{F0: 42, F1: 5, F2: -19}
                        v4 = v4
                        v6 -= (v4 | 42)
                        _ = v10
                    } else {
                        var v11 float64 = -2.25
                        _ = v11
                    }
                    var v12 int16 = -831
                    if false {
                        v13 := GenS3{F0: "", F1: "a", F2: 556}
                        for fi14 := 0; fi14 < 2; fi14++ {
                            for fi15 := 0; fi15 < 6; fi15++ {
                                v8 += v1
                                v8 += v8
                                switch "\n\t" {
                                case "a":
                                    _ = len(v9)
                                case "\n\t":
                                    v16 := GenS3{F0: (((("foo bar" + (("\n\t" + "hello") + "hello")) + "\n\t") + ("foo bar" + "\n\t")) + "foo bar"), F1: ("" + "a"), F2: 4294967295}
                                    _ = v16
                                case "hello":
                                    switch 32768 {
                                    case -9223372036854775808:
                                        if (-38 <= -128) {
                                            var v17 uint = 2
                                            _ = v17
                                        }
                                    case -1:
                                        v9 = append(v9, 333)
                                    case 32768:
                                        fmt.Println(v4)
                                    default:
                                        if (3.14159 == ((0.0 + (1.5 / (1.5 / (((3.14159 + (-2.25 / 1e10)) * (1.5 * 3.14159)) / 1.0)))) + 0.0)) {
                                            v1 = (v1 | (v1 << ((255) & 63)))
                                            v9 = append(v9, (-(fi14)))
                                            v8 = v1
                                            var v18 uint32 = 2147483648
                                            _ = v18
                                        } else {
                                            v9 = append(v9, (fi14 * (fi15 - (fi15 ^ fi14))))
                                        }
                                    }
                                default:
                                    v19 := GenS4{F0: -185}
                                    _ = v19
                                }
                                v9 = append(v9, 1)
                                var v20 int32 = 2147483647
                                _ = fi15
                                _ = v20
                            }
                            v4 = (v3 - v7)
                            v9 = append(v9, (fi14 - (fi14 & -57)))
                            _ = fi14
                        }
                        _ = v13
                    } else {
                        fmt.Println(v3, v2)
                        for fi21 := 0; fi21 < 19; fi21++ {
                            fmt.Println(v1)
                            var v22 uint16 = 522
                            if true {
                                var v23 uint16 = 65535
                                v24 := GenS1{F0: 0, F1: v12, F2: v2}
                                _ = v23
                                _ = v24
                            }
                            _ = fi21
                            _ = v22
                        }
                        v4 += 2
                    }
                    _ = v12
                } else {
                    if true {
                        fmt.Println(v6, v1)
                        if true {
                            v2 = (v2 + v2)
                            var v25 uint = 18446744073709551615
                            _ = v25
                        } else {
                            var v26 uint16 = 982
                            _ = v26
                        }
                        v6 -= v6
                    }
                }
            case 4294967295:
                switch -1 {
                case 2147483647:
                    fmt.Println(v1)
                case 597:
                    for fi27 := 0; fi27 < 5; fi27++ {
                        for fi28 := 0; fi28 < 4; fi28++ {
                            v3 += v3
                            _ = fi28
                        }
                        _ = fi27
                    }
                case 9223372036854775807:
                    if true {
                        switch "hello" {
                        case "a":
                            _ = len(v9)
                        case "hello":
                            v4 -= v7
                        case "":
                            var v29 uint8 = (^(v6))
                            _ = v29
                        }
                        switch 127 {
                        case 2147483647:
                            var v30 int8 = 127
                            _ = v30
                        case 32767:
                            if (132 < 178) {
                                _ = len(v9)
                            }
                        }
                        v31 := GenS3{F0: "\n\t", F1: "hello", F2: 458}
                        _ = v31
                    }
                default:
                    for fi32 := 0; fi32 < 14; fi32++ {
                        fmt.Println(v3)
                        _ = fi32
                    }
                }
            default:
                v9 = append(v9, 848)
            }
            if true {
                v33 := GenS1{F0: 99, F1: (-(v2)), F2: (v2 / ((v2) | 1))}
                for fi34 := 0; fi34 < 19; fi34++ {
                    v8 = 2
                    switch -992 {
                    case -789:
                        var v35 bool = false
                        _ = v35
                    case 42:
                        v36 := GenS5{F0: "\n\t", F1: -1, F2: true}
                        _ = v36
                    }
                    _ = fi34
                }
                _ = v33
            }
        }
        switch int(v8) {
        case -859:
            v37 := GenS1{F0: -1, F1: (^(v2)), F2: v2}
            _ = v37
        case 9223372036854775807:
            fmt.Println(v2)
        }
        _ = v9
    }
    var v38 int16 = -32768
    v2 = (v2 & v2)
    var v39 uint16 = uint16(v3)
    for fi40 := 0; fi40 < 7; fi40++ {
        v38 -= v38
        _ = fi40
    }
    var v41 int8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v41
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
