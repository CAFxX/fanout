package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 string; F2 uint }
type GenS3 struct { F0 int64; F1 int }
type GenS4 struct { F0 bool; F1 uint8; F2 int32 }

func main() {
    var v1 int64 = 1
    var v2 uint64 = 0
    v3 := []int{-824, -2147483648, 42}
    var v4 bool = false
    var v5 uint = 1
    fmt.Println(v4, v2)
    switch (("hello" + "foo bar") + (("foo bar" + "foo bar") + "\n\t")) {
    case "":
        v5 = v5
    }
    switch 2147483647 {
    case 127:
        var v6 uint16 = 1
        _ = v6
    }
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 uint16 = 255
        v9 := GenS1{F0: (fi7 ^ (-(fi7)))}
        var v10 bool = true
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 int16 = int16(v1)
    switch 128 {
    case 435:
        v5 = 383
    default:
        v2 = 4294967295
    }
    v12 := GenS2{F0: v2, F1: "foo bar", F2: v5}
    var v13 bool = v4
    fmt.Println(v13, v4, v4)
    switch "a" {
    case "a":
        v11 -= (-(v11))
    case "\n\t":
        v13 = (-116 >= 100)
    case "":
        for fi14 := 0; fi14 < 2; fi14++ {
            if v13 {
                fmt.Println(v4)
                for fi15 := 0; fi15 < 12; fi15++ {
                    v16 := GenS2{F0: v2, F1: (("" + "") + "foo bar"), F2: v5}
                    var v17 bool = (!v13)
                    if v4 {
                        if (true || ((100 == -45) && (42 != 659))) {
                            v18 := GenS3{F0: (v1 + (v1 << (((v5 + v5)) & 63))), F1: 4294967295}
                            v13 = v17
                            _ = len(v3)
                            fmt.Println(v11, v5)
                            v17 = ((0.1 * (3.14159 - float64(0.1))) == (0.0 + 0.1))
                            if (v1 == v1) {
                                switch 589 {
                                case 9223372036854775807:
                                    v3 = append(v3, -483)
                                case -533:
                                    var v19 int64 = v1
                                    _ = v19
                                case 127:
                                    switch (^(fi14)) {
                                    case -9223372036854775808:
                                        var v20 uint64 = (v2 << ((v5) & 63))
                                        _ = v20
                                    case 0:
                                        var v21 int64 = 329
                                        _ = v21
                                    case 977:
                                        if v4 {
                                            v4 = true
                                            v22 := GenS4{F0: false, F1: 1, F2: -2147483648}
                                            _ = v22
                                        }
                                    default:
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                }
                                v1 = v1
                                fmt.Println(v5)
                            }
                            v1 += (-(v1))
                            _ = v18
                        } else {
                            if (2 <= 229) {
                                var v23 int8 = -128
                                _ = v23
                            } else {
                                var v24 uint = 2
                                _ = len(v3)
                                var v25 uint8 = 42
                                if (float32(float64(0.0)) == 3.14159) {
                                    var v26 int16 = v11
                                    var v27 uint32 = 1
                                    v24 -= (v24 / (((v5 % ((v5) | 1))) | 1))
                                    _ = v26
                                    _ = v27
                                }
                                _ = v24
                                _ = v25
                            }
                            v3 = append(v3, (^(fi15)))
                        }
                        var v28 int32 = -2147483648
                        v5 = (v5 | v5)
                        fmt.Println(v5)
                        _ = v28
                    } else {
                        v2 = 0
                    }
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            } else {
                if (237 >= 32767) {
                    if (1e10 != 3.14159) {
                        fmt.Println(v4)
                        v5 -= 4294967295
                        v29 := GenS4{F0: false, F1: 128, F2: 2147483647}
                        _ = v29
                    }
                    if v4 {
                        var v30 uint = v5
                        _ = v30
                    } else {
                        switch (-(fi14)) {
                        case 32767:
                            if false {
                                v13 = ((false && (!v13)) && false)
                            }
                        case 9223372036854775807:
                            v5 -= 4294967295
                        }
                        switch 42 {
                        case 32768:
                            fmt.Println(v2, v5)
                        case 2147483647:
                            for fi31 := 0; fi31 < 4; fi31++ {
                                v13 = (v5 == (^(v5)))
                                var v32 int32 = 42
                                _ = fi31
                                _ = v32
                            }
                        }
                    }
                }
                fmt.Println(v13, v2)
            }
            _ = fi14
        }
    }
    for fi33 := 0; fi33 < 1; fi33++ {
        v5 += v5
        _ = fi33
    }
    v1 = (v1 % ((-304) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
