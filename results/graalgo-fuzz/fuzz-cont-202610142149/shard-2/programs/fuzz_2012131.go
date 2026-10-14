package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 int16; F2 int }
type GenS4 struct { F0 uint8; F1 int8; F2 int16 }
type GenS5 struct { F0 uint64; F1 bool; F2 uint }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint; F1 string }
func genFn0() bool {
    var v4 uint64 = 2
    switch 2147483647 {
    case -1:
        switch -2147483648 {
        case -2147483648:
            var v5 int8 = 100
            _ = v5
        case 42:
            v6 := GenS1{F0: 819}
            _ = v6
        default:
            switch 127 {
            case 2147483647:
                for fi7 := 0; fi7 < 13; fi7++ {
                    var v8 int64 = 9223372036854775807
                    _ = fi7
                    _ = v8
                }
            case 60:
                if true {
                    switch -778 {
                    case 32768:
                        var v9 uint8 = 128
                        _ = v9
                    default:
                        var v10 float32 = ((0.1 - -2.25) * (-2.25 - (3.14159 + 3.14159)))
                        _ = v10
                    }
                    v11 := []int{127, 1, 42}
                    _ = v11
                } else {
                    v4 = (v4 << ((4294967295) & 63))
                    v4 += v4
                    v12 := GenS2{F0: 695}
                    _ = v12
                }
            case -164:
                var v13 string = ("a" + ("" + "a"))
                _ = v13
            default:
                v4 = (v4 + (v4 % ((271) | 1)))
            }
        }
    default:
        switch "foo bar" {
        case "a":
            var v14 uint = 255
            _ = v14
        case "hello":
            for fi15 := 0; fi15 < 13; fi15++ {
                for fi16 := 0; fi16 < 17; fi16++ {
                    v4 = v4
                    var v17 int16 = int16(fi15)
                    _ = fi16
                    _ = v17
                }
                v18 := GenS2{F0: -429}
                var v19 int16 = 1
                _ = fi15
                _ = v18
                _ = v19
            }
        default:
            v4 = (^(v4))
        }
    }
    for fi20 := 0; fi20 < 12; fi20++ {
        var v21 float32 = ((0.1 - (1.5 - 0.1)) + (0.0 + 3.14159))
        var v22 uint64 = (v4 & (v4 + 42))
        _ = fi20
        _ = v21
        _ = v22
    }
    fmt.Println(v4)
    _ = v4
    return (0.0 <= 0.0)
}

func main() {
    var v1 bool = false
    var v2 uint16 = 255
    var v3 int = 4294967295
    if (74 <= -34) {
        v1 = true
        fmt.Println(v1, v2)
        var v23 uint = 814
        _ = v23
    }
    v24 := []int{v3, 9223372036854775807}
    fmt.Println(v2, v1)
    if false {
        if len(v24) > 0 {
            v24[0] = v3
        }
    }
    switch "\n\t" {
    case "\n\t":
        var v25 bool = v1
        _ = v25
    case "a":
        v26 := GenS3{F0: 838, F1: -32768, F2: (^(v3))}
        _ = v26
    case "":
        for fi27 := 0; fi27 < 12; fi27++ {
            fmt.Println(v1, v1, v1)
            var v28 int32 = -485
            fmt.Println(fi27, v2, v2)
            _ = fi27
            _ = v28
        }
    default:
        for fi29 := 0; fi29 < 16; fi29++ {
            fmt.Println(v1, v2, v3)
            _ = fi29
        }
    }
    if (!true) {
        for fi30 := 0; fi30 < 9; fi30++ {
            for fi31 := 0; fi31 < 16; fi31++ {
                var v32 float32 = 0.1
                v2 += 1
                _ = fi31
                _ = v32
            }
            v2 -= 436
            _ = fi30
        }
    } else {
        if ("" != "a") {
            v3 -= (-(v3))
        } else {
            v33 := GenS3{F0: -411, F1: 981, F2: v3}
            _ = v33
        }
    }
    fmt.Println(v1, v2, v1)
    if v1 {
        switch (v3 ^ 4294967295) {
        case 4294967295:
            var v34 int = v3
            _ = v34
        case 128:
            v2 -= (v2 << ((1) & 63))
        case -1:
            v24 = append(v24, (v3 * (-(v3))))
        }
    } else {
        switch v3 {
        case 9223372036854775807:
            if true {
                switch (^(v3)) {
                case 32768:
                    for fi35 := 0; fi35 < 5; fi35++ {
                        switch ("hello" + "foo bar") {
                        case "hello":
                            v3 = -246
                        default:
                            v36 := GenS2{F0: 42}
                            _ = v36
                        }
                        _ = fi35
                    }
                }
                v1 = true
                v37 := GenS4{F0: 127, F1: -128, F2: 123}
                _ = v37
            }
        default:
            if v1 {
                v38 := GenS5{F0: 204, F1: v1, F2: 2}
                if false {
                    var v39 uint32 = 393
                    for fi40 := 0; fi40 < 11; fi40++ {
                        v24 = append(v24, (fi40 + v3))
                        var v41 float64 = (float64(float32(-2.25)) / float64((float32(float64(0.1)) / 1.0)))
                        _ = fi40
                        _ = v41
                    }
                    fmt.Println(v2, v2)
                    _ = v39
                } else {
                    _ = len(v24)
                    fmt.Println(v1, v1, v1)
                }
                _ = v38
            }
        }
        for fi42 := 0; fi42 < 6; fi42++ {
            for fi43 := 0; fi43 < 10; fi43++ {
                fmt.Println(v2, v2, v1)
                _ = fi43
            }
            _ = len(v24)
            v44 := GenS6{F0: 42}
            _ = fi42
            _ = v44
        }
    }
    v3 = (v3 - v3)
    v45 := GenS7{F0: 27, F1: ("\n\t" + "hello")}
    var v46 uint = 927
    if len(v24) > 0 {
        v24[0] = (v3 / ((0) | 1))
    }
    for fi47 := 0; fi47 < 3; fi47++ {
        for fi48 := 0; fi48 < 8; fi48++ {
            if true {
                var v49 uint64 = 2
                var v50 uint16 = v2
                v3 = (fi47 & fi48)
                _ = v49
                _ = v50
            }
            v46 -= 197
            _ = fi48
        }
        for fi51 := 0; fi51 < 16; fi51++ {
            v3 = (fi47 >> ((114) & 63))
            _ = fi51
        }
        _ = fi47
    }
    _ = len(v24)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v45
    _ = v46
}
