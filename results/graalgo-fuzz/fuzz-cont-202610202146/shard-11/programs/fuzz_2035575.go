package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 string; F1 uint64 }
type GenS6 struct { F0 uint64; F1 int32 }
type GenS7 struct { F0 uint8; F1 uint }
type GenS8 struct { F0 uint8; F1 int32 }
func genFn0(p0 int8, p1 uint8, p2 int8) bool {
    v4 := GenS1{F0: 1}
    v5 := GenS2{F0: -1}
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return false
}

func main() {
    var v1 uint8 = 42
    var v2 int = 127
    v3 := []uint32{512, 65535, 321}
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := GenS3{F0: true}
        var v8 uint = 18446744073709551615
        _ = fi6
        _ = v7
        _ = v8
    }
    v2 = v2
    var v9 uint8 = v1
    for fi10 := 0; fi10 < 17; fi10++ {
        v2 = (fi10 | (v2 & (fi10 >> ((255) & 63))))
        _ = fi10
    }
    v1 = 166
    v2 -= 1
    for fi11 := 0; fi11 < 11; fi11++ {
        v12 := GenS1{F0: 1}
        _ = fi11
        _ = v12
    }
    if true {
        _ = len(v3)
        if (24 <= 18446744073709551615) {
            var v13 uint = 42
            if true {
                for fi14 := 0; fi14 < 3; fi14++ {
                    v15 := GenS1{F0: v13}
                    for fi16 := 0; fi16 < 9; fi16++ {
                        var v17 uint32 = 2147483648
                        v2 = (-(fi16))
                        v1 -= v9
                        _ = fi16
                        _ = v17
                    }
                    _ = fi14
                    _ = v15
                }
                var v18 float32 = float32(0.1)
                _ = v18
            } else {
                var v19 uint8 = (v9 | 255)
                if ("hello" <= "hello") {
                    for fi20 := 0; fi20 < 3; fi20++ {
                        v21 := GenS3{F0: true}
                        _ = fi20
                        _ = v21
                    }
                }
                _ = v19
            }
            _ = v13
        } else {
            _ = len(v3)
        }
        switch (v2 & (v2 | v2)) {
        case 42:
            if (!true) {
                v1 = 127
                if (true || true) {
                    if ((0.0 / 0.1) >= 1.5) {
                        switch "hello" {
                        case "foo bar":
                            v9 = (v1 / ((2) | 1))
                        case "":
                            var v22 uint8 = (v1 | v1)
                            _ = v22
                        default:
                            var v23 uint16 = 65535
                            _ = v23
                        }
                    } else {
                        _ = len(v3)
                        var v24 uint16 = 65535
                        _ = v24
                    }
                    if (2147483647 < 2147483647) {
                        var v25 uint32 = 1
                        fmt.Println(v2)
                        _ = v25
                    }
                    if false {
                        var v26 string = ("" + ("\n\t" + ("foo bar" + ("foo bar" + "hello"))))
                        switch (v2 >> ((2) & 63)) {
                        case 9223372036854775807:
                            v1 = v9
                        case -140:
                            fmt.Println(v9)
                        default:
                            switch v2 {
                            case 1:
                                if (25 > 0) {
                                    for fi27 := 0; fi27 < 1; fi27++ {
                                        for fi28 := 0; fi28 < 14; fi28++ {
                                            var v29 uint32 = 504
                                            _ = fi28
                                            _ = v29
                                        }
                                        v30 := GenS4{F0: 304}
                                        var v31 float32 = float32(3.14159)
                                        _ = fi27
                                        _ = v30
                                        _ = v31
                                    }
                                    var v32 bool = (((0.0 - 3.14159) / 3.14159) <= ((0.0 + (0.1 + (0.0 - (((float32(3.14159) - 0.1) - -2.25) / 3.14159)))) - ((((0.1 + 3.14159) - 1e10) + (1e10 / 1.0)) * 0.0)))
                                    var v33 int64 = -1
                                    _ = v32
                                    _ = v33
                                } else {
                                    for fi34 := 0; fi34 < 4; fi34++ {
                                        if true {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            v35 := GenS5{F0: v26, F1: 2}
                                            _ = v35
                                        }
                                        _ = fi34
                                    }
                                    fmt.Println(v1, v2)
                                    switch ("foo bar" + v26) {
                                    case "a":
                                        switch "" {
                                        case "a":
                                            fmt.Println(v1, v9, v1)
                                        case "foo bar":
                                            v2 += (-(v2))
                                        default:
                                            var v36 float64 = 0.0
                                            _ = v36
                                        }
                                    case "foo bar":
                                        for fi37 := 0; fi37 < 10; fi37++ {
                                            v9 = (^(v9))
                                            var v38 float32 = (float32(((-2.25 * -2.25) * 0.0)) - 1.5)
                                            v26 = ("foo bar" + "foo bar")
                                            _ = fi37
                                            _ = v38
                                        }
                                    }
                                }
                            default:
                                v26 = (v26 + "a")
                            }
                        }
                        _ = v26
                    }
                } else {
                    fmt.Println(v2)
                    v39 := GenS6{F0: 2, F1: -1}
                    _ = v39
                }
            }
        default:
            fmt.Println(v2)
        }
    }
    switch "foo bar" {
    case "foo bar":
        v40 := GenS7{F0: v1, F1: 65535}
        _ = v40
    case "hello":
        v41 := GenS8{F0: v9, F1: 799}
        _ = v41
    case "":
        v9 += 255
    }
    for fi42 := 0; fi42 < 3; fi42++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        switch ("foo bar" + ("" + "")) {
        case "\n\t":
            fmt.Println(v9, v2, v9)
        }
        _ = len(v3)
        _ = fi42
    }
    if (!(false || (!false))) {
        if false {
            for fi43 := 0; fi43 < 19; fi43++ {
                var v44 uint8 = v9
                switch -9223372036854775808 {
                case -664:
                    if false {
                        var v45 int32 = 42
                        switch (v2 ^ (fi43 & (-(fi43)))) {
                        case -1:
                            var v46 string = "\n\t"
                            _ = v46
                        case 902:
                            if false {
                                v1 += (v1 % (((v44 << ((18446744073709551615) & 63))) | 1))
                                v47 := GenS8{F0: (v44 >> ((18446744073709551615) & 63)), F1: 42}
                                _ = v47
                            }
                        default:
                            for fi48 := 0; fi48 < 15; fi48++ {
                                var v49 uint64 = 4294967295
                                if len(v3) > 0 {
                                    v3[0] = uint32(v49)
                                }
                                _ = fi48
                                _ = v49
                            }
                        }
                        _ = v45
                    }
                }
                var v50 int8 = 1
                _ = fi43
                _ = v44
                _ = v50
            }
        } else {
            v9 = 0
            v3 = append(v3, 2147483647)
            if len(v3) > 0 {
                v3[0] = uint32(v1)
            }
        }
    } else {
        v2 = (v2 >> ((723) & 63))
        var v51 int = -370
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v9
}
