package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 bool; F1 uint16; F2 int }

func main() {
    var v1 bool = (1 > 18446744073709551615)
    var v2 uint16 = 42
    v3 := []uint32{2147483647, 2147483648, 0}
    var v4 int16 = 0
    var v5 uint = 18446744073709551615
    var v6 int = -346
    for fi7 := 0; fi7 < 20; fi7++ {
        fmt.Println(v2, v5)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi7
    }
    fmt.Println(v6, v5)
    if (1 > 2) {
        switch v6 {
        case 1:
            switch "foo bar" {
            case "hello":
                v3 = append(v3, 2147483647)
            case "foo bar":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "a":
                var v8 int32 = 0
                _ = v8
            default:
                v9 := GenS1{F0: (v6 % ((v6) | 1)), F1: 65535, F2: 1}
                _ = v9
            }
        case 2147483647:
            v3 = append(v3, 565)
        default:
            v5 -= (v5 + v5)
        }
    }
    _ = len(v3)
    var v10 float64 = 1.5
    v1 = (32768 > v6)
    if v1 {
        v4 = (v4 & (v4 ^ -360))
    }
    var v11 int8 = 18
    var v12 int64 = 9223372036854775807
    var v13 int = (-(v6))
    if v1 {
        fmt.Println(v6, v12)
        fmt.Println(v13, v6)
        if (((1e10 * (float32(v10) * ((float32(1e10) - (3.14159 * 3.14159)) + -2.25))) - -2.25) <= (1.5 / -2.25)) {
            switch (("" + "\n\t") + "a") {
            case "\n\t":
                if (false || ((v4 ^ v4) <= (-(v4)))) {
                    for fi14 := 0; fi14 < 14; fi14++ {
                        if false {
                            v10 += 1.5
                            var v15 float64 = float64((0.0 + (-2.25 * 3.14159)))
                            for fi16 := 0; fi16 < 13; fi16++ {
                                var v17 int16 = (v4 * (v4 - (^(v4))))
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v18 := GenS1{F0: -868, F1: v2, F2: 42}
                                _ = fi16
                                _ = v17
                                _ = v18
                            }
                            _ = v15
                        } else {
                            for fi19 := 0; fi19 < 6; fi19++ {
                                var v20 uint32 = 0
                                _ = fi19
                                _ = v20
                            }
                            for fi21 := 0; fi21 < 7; fi21++ {
                                fmt.Println(v2)
                                if false {
                                    fmt.Println(v2, v1, v10)
                                    if (!((v4 != v4) || v1)) {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        var v22 uint32 = 129
                                        v1 = (!(v12 == -839))
                                        var v23 float32 = (1e10 / 3.14159)
                                        _ = v22
                                        _ = v23
                                    }
                                    v2 += (^(v2))
                                    v3 = append(v3, 65535)
                                    fmt.Println(v1, v2)
                                } else {
                                    var v24 uint = v5
                                    var v25 int16 = (v4 - (v4 * (v4 - (v4 >> ((v24) & 63)))))
                                    _ = v24
                                    _ = v25
                                }
                                switch ("\n\t" + "") {
                                case "\n\t":
                                    v10 += (float64((-2.25 / 0.1)) - v10)
                                    switch (fi21 ^ fi21) {
                                    case 127:
                                        _ = len(v3)
                                    default:
                                        v1 = v1
                                    }
                                    v1 = false
                                case "":
                                    v1 = false
                                default:
                                    v26 := GenS1{F0: (v13 - (fi21 * 4294967295)), F1: v2, F2: 18446744073709551615}
                                    _ = v26
                                }
                                _ = fi21
                            }
                            if ((0.1 * v10) < 0.1) {
                                v10 += (v10 * 1.5)
                                v10 = (v10 - (0.1 - (v10 / v10)))
                            } else {
                                for fi27 := 0; fi27 < 8; fi27++ {
                                    v10 = 0.0
                                    _ = fi27
                                }
                                var v28 bool = false
                                _ = v28
                            }
                        }
                        var v29 uint32 = 2147483648
                        _ = fi14
                        _ = v29
                    }
                } else {
                    var v30 int = v6
                    _ = v30
                }
            case "foo bar":
                for fi31 := 0; fi31 < 9; fi31++ {
                    if v1 {
                        for fi32 := 0; fi32 < 20; fi32++ {
                            fmt.Println(v6, v5)
                            if v1 {
                                v33 := GenS1{F0: fi31, F1: v2, F2: 0}
                                var v34 bool = false
                                _ = v33
                                _ = v34
                            }
                            v11 = -128
                            _ = fi32
                        }
                    }
                    if (((fi31 + 698) != 42) || (42 >= v5)) {
                        var v35 float32 = -2.25
                        var v36 int8 = (v11 >> ((v5) & 63))
                        for fi37 := 0; fi37 < 18; fi37++ {
                            v36 = v36
                            v2 -= v2
                            v6 -= fi37
                            _ = fi37
                        }
                        _ = v35
                        _ = v36
                    }
                    _ = fi31
                }
            default:
                var v38 uint8 = 202
                _ = v38
            }
            switch (^(v13)) {
            case 678:
                var v39 uint32 = 65535
                _ = v39
            case 128:
                for fi40 := 0; fi40 < 1; fi40++ {
                    v41 := GenS1{F0: -1, F1: v2, F2: 395}
                    for fi42 := 0; fi42 < 10; fi42++ {
                        v10 -= (v10 / (float64(((float32(((0.1 * v10) + v10)) + 1.5) * 1e10)) * 1e10))
                        v12 = v12
                        v6 -= v6
                        if (!true) {
                            switch "\n\t" {
                            case "a":
                                v3 = append(v3, 898)
                            case "hello":
                                fmt.Println(v5, v5, v4)
                                v2 -= 1
                                v43 := GenS2{F0: v13}
                                _ = v43
                            case "foo bar":
                                if v1 {
                                    switch "a" {
                                    case "":
                                        fmt.Println(v10, v13)
                                    }
                                    v44 := GenS1{F0: -761, F1: v2, F2: uint64(v5)}
                                    if v1 {
                                        fmt.Println(v1, v4)
                                        v45 := GenS3{F0: -960, F1: v5}
                                        if true {
                                            if v1 {
                                                _ = len(v3)
                                                v6 += (v6 + (fi40 | v13))
                                            }
                                        }
                                        _ = v45
                                    }
                                    _ = v44
                                }
                                v46 := GenS4{F0: (!(2147483648 != 2)), F1: (v2 + 2), F2: 42}
                                var v47 uint8 = 2
                                _ = v46
                                _ = v47
                            }
                            v2 = 32767
                            var v48 string = ("\n\t" + (("a" + "a") + "\n\t"))
                            _ = v48
                        }
                        v10 = 1.5
                        _ = fi42
                    }
                    _ = fi40
                    _ = v41
                }
            case 2147483647:
                v5 = (^(v5))
            }
        } else {
            var v49 uint16 = (v2 | 0)
            _ = len(v3)
            var v50 int8 = v11
            _ = v49
            _ = v50
        }
    }
    v51 := GenS2{F0: v6}
    var v52 int32 = 0
    var v53 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v53
    _ = v6
}
