package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 bool; F2 bool }
type GenS4 struct { F0 int32; F1 int; F2 int16 }
type GenS5 struct { F0 bool; F1 uint16; F2 uint16 }

func main() {
    v1 := []int{-822, 2147483647, 2147483647, 1}
    var v2 uint16 = 963
    v3 := make(map[string]bool)
    if (((-32768 < -906) && false) || (!false)) {
        v2 = 2
    } else {
        for fi4 := 0; fi4 < 11; fi4++ {
            switch "foo bar" {
            case "a":
                switch "a" {
                case "":
                    var v5 bool = false
                    _ = v5
                case "foo bar":
                    if true {
                        fmt.Println(fi4, fi4)
                    } else {
                        fmt.Println(v2)
                        v2 += (v2 << ((42) & 63))
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                case "a":
                    switch (("hello" + (("a" + "\n\t") + "\n\t")) + (("\n\t" + "\n\t") + "")) {
                    case "hello":
                        var v6 int16 = 1000
                        _ = v6
                    }
                default:
                    v2 += 1
                }
            case "foo bar":
                var v7 bool = false
                _ = v7
            }
            _ = len(v3)
            if (0.1 <= 1.5) {
                v2 = (v2 ^ 2)
                var v8 int32 = 0
                _ = v8
            } else {
                var v9 uint8 = 67
                fmt.Println(v2, fi4, v2)
                _ = v9
            }
            _ = fi4
        }
        switch 2147483647 {
        case 2147483647:
            v10 := GenS1{F0: 1, F1: 1}
            _ = v10
        }
        v2 -= v2
    }
    v11 := GenS1{F0: v2, F1: 0}
    if (0.1 >= 3.14159) {
        v1 = append(v1, 9223372036854775807)
        v2 = (v2 % (((v2 % ((42) | 1))) | 1))
    }
    var v12 int64 = 9223372036854775807
    var v13 float64 = (float64(float32((0.1 - 0.1))) * 1.5)
    for fi14 := 0; fi14 < 11; fi14++ {
        if false {
            var v15 int16 = 32767
            _ = v15
        }
        if (894 >= 2) {
            var v16 bool = (!true)
            _ = len(v1)
            _ = v16
        }
        v12 = (v12 % ((9223372036854775807) | 1))
        _ = fi14
    }
    fmt.Println(v13)
    v12 = 9223372036854775807
    v3[""] = true
    v17 := GenS1{F0: (v2 * (^(v2))), F1: 42}
    switch 209 {
    case 42:
        var v18 bool = false
        _ = v18
    case 1:
        fmt.Println(v12, v13)
    case 32767:
        for fi19 := 0; fi19 < 9; fi19++ {
            for fi20 := 0; fi20 < 18; fi20++ {
                if len(v1) > 0 {
                    v1[0] = (^(fi19))
                }
                switch fi20 {
                case 2147483647:
                    var v21 uint = 2
                    _ = v21
                case -267:
                    switch ("foo bar" + "a") {
                    case "foo bar":
                        if true {
                            v13 = 1e10
                            v22 := GenS2{F0: 4294967295}
                            v2 += v2
                            _ = v22
                        }
                    case "":
                        v23 := GenS3{F0: (fi20 * (fi20 + 9223372036854775807)), F1: true, F2: true}
                        _ = v23
                    }
                default:
                    v2 = (v2 ^ 910)
                }
                var v24 int32 = -168
                _ = fi20
                _ = v24
            }
            v25 := GenS2{F0: (fi19 + fi19)}
            fmt.Println(v13, v12, v12)
            _ = fi19
            _ = v25
        }
    default:
        fmt.Println(v2, v2)
    }
    if (false && ((v12 << ((156) & 63)) <= -9223372036854775808)) {
        v13 = (float64(3.14159) - v13)
    }
    v26 := GenS4{F0: -2147483648, F1: 274, F2: -618}
    if len(v1) > 0 {
        v1[0] = 32768
    }
    v13 += (1.5 - 0.1)
    if (!true) {
        fmt.Println(v13, v12, v12)
    } else {
        fmt.Println(v2)
        for fi27 := 0; fi27 < 15; fi27++ {
            for fi28 := 0; fi28 < 14; fi28++ {
                v29 := v3[(("a" + "foo bar") + "foo bar")]
                var v30 int = (fi27 >> ((748) & 63))
                v13 = v13
                _ = fi28
                _ = v29
                _ = v30
            }
            _ = fi27
        }
        fmt.Println(v12, v12, v2)
    }
    v31 := GenS5{F0: (!(!(false || ((-717 >= -168) && false)))), F1: (v2 & (^(v2))), F2: (v2 >> ((255) & 63))}
    for fi32 := 0; fi32 < 18; fi32++ {
        switch "" {
        case "hello":
            if (0.1 <= ((((0.1 * 1.5) + (((0.0 * (1e10 + 0.0)) * float32(v13)) / -2.25)) + 1.5) / 1e10)) {
                var v33 bool = true
                _ = v33
            } else {
                var v34 int8 = -1
                _ = v34
            }
        }
        v35 := GenS2{F0: 32768}
        _ = fi32
        _ = v35
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v26
    _ = v3
    _ = v31
}
