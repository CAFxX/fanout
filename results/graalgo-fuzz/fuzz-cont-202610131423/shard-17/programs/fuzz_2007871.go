package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint16; F1 string; F2 int8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 bool; F1 uint8; F2 bool }

func main() {
    v1 := []string{}
    var v2 float64 = 1e10
    var v3 string = ""
    var v4 uint16 = 647
    var v5 int64 = -1
    switch ("a" + ("" + "\n\t")) {
    case "\n\t":
        switch 127 {
        case 32768:
            v6 := GenS1{F0: 63}
            _ = v6
        }
    case "hello":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        if false {
            var v7 uint32 = 88
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            switch "foo bar" {
            case "a":
                v2 = 3.14159
            case "hello":
                v8 := GenS1{F0: 121}
                _ = v8
            case "":
                for fi9 := 0; fi9 < 6; fi9++ {
                    if false {
                        var v10 int16 = 42
                        var v11 uint16 = v4
                        fmt.Println(v7)
                        _ = v10
                        _ = v11
                    } else {
                        v12 := GenS2{F0: -81}
                        for fi13 := 0; fi13 < 14; fi13++ {
                            v3 = (v3 + "a")
                            _ = fi13
                        }
                        _ = v12
                    }
                    var v14 bool = ((fi9 << ((646) & 63)) < fi9)
                    _ = fi9
                    _ = v14
                }
            }
            _ = v7
        } else {
            v4 += (v4 * 255)
        }
    }
    var v15 uint16 = v4
    var v16 float64 = float64(-2.25)
    var v17 float64 = (v16 - -2.25)
    v18 := GenS1{F0: 58}
    if (false && ((v4 | v4) > v4)) {
        var v19 int16 = -899
        for fi20 := 0; fi20 < 10; fi20++ {
            var v21 uint16 = (v15 ^ (v15 & v15))
            switch v3 {
            case "a":
                switch (fi20 & 2147483647) {
                case -2147483648:
                    v19 -= (v19 ^ v19)
                case 678:
                    v22 := GenS3{F0: 42}
                    _ = v22
                case 42:
                    v2 = v2
                default:
                    v15 = 1
                }
            case "foo bar":
                if false {
                    var v23 uint64 = 684
                    _ = v23
                }
            default:
                if (!((fi20 % ((fi20) | 1)) >= (fi20 % (((fi20 * (fi20 | -9223372036854775808))) | 1)))) {
                    v17 = v16
                } else {
                    fmt.Println(v21)
                    v21 -= 32767
                }
            }
            _ = fi20
            _ = v21
        }
        _ = v19
    }
    _ = len(v1)
    _ = len(v1)
    v17 -= ((v16 / (-2.25 / (v16 / float64((((1e10 * (1.5 - 1.5)) / 3.14159) + 1.5))))) * float64(float32(3.14159)))
    switch -2147483648 {
    case 32767:
        if true {
            _ = len(v1)
        } else {
            v24 := GenS4{F0: v15, F1: "", F2: -13}
            var v25 float32 = 3.14159
            _ = v24
            _ = v25
        }
    case 42:
        fmt.Println(v15, v15, v3)
    default:
        var v26 uint16 = (^(v15))
        _ = v26
    }
    if len(v1) > 0 {
        v1[0] = v3
    }
    if (!true) {
        switch 128 {
        case 9223372036854775807:
            if false {
                fmt.Println(v4, v15, v2)
            } else {
                if (v15 <= v15) {
                    if true {
                        for fi27 := 0; fi27 < 7; fi27++ {
                            var v28 uint16 = 731
                            _ = fi27
                            _ = v28
                        }
                        v15 += (v15 << ((991) & 63))
                        v29 := GenS5{F0: 0}
                        _ = v29
                    }
                    v16 = ((v2 - (1e10 - v16)) * (0.0 / v2))
                    if (-128 >= 123) {
                        v4 += 255
                    } else {
                        v30 := GenS6{F0: (v5 > v5), F1: 2, F2: (0 > 32767)}
                        v31 := GenS4{F0: 0, F1: "hello", F2: 1}
                        v5 = (v5 | 46)
                        _ = v30
                        _ = v31
                    }
                } else {
                    var v32 uint64 = 645
                    _ = v32
                }
                v4 = (v15 / (((^(v15))) | 1))
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        }
    } else {
        v4 = (v4 % (((v4 >> ((18446744073709551615) & 63))) | 1))
        fmt.Println(v2)
        fmt.Println(v15, v16)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
