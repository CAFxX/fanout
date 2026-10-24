package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32; F1 uint32; F2 int32 }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 uint16; F1 uint32; F2 uint16 }

func main() {
    var v1 uint32 = 2
    var v2 int = 1
    var v3 int32 = 322
    var v4 uint = 0
    var v5 int8 = -128
    v5 = -1
    if false {
        v3 = v3
        if false {
            fmt.Println(v5)
            for fi6 := 0; fi6 < 11; fi6++ {
                var v7 uint8 = 111
                for fi8 := 0; fi8 < 19; fi8++ {
                    var v9 uint = 65535
                    _ = fi8
                    _ = v9
                }
                fmt.Println(v5)
                _ = fi6
                _ = v7
            }
        }
        var v10 int8 = (-(v5))
        _ = v10
    } else {
        v11 := []int{(v2 + (v2 + (v2 & (v2 & 781)))), v2, (^(v2)), (v2 * (v2 >> (((v4 - 18446744073709551615)) & 63)))}
        _ = v11
    }
    if true {
        v3 += v3
        var v12 uint16 = 1
        _ = v12
    } else {
        v3 = v3
        v13 := GenS1{F0: 9223372036854775807}
        v1 -= (v1 | v1)
        _ = v13
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        switch "hello" {
        case "":
            fmt.Println(v5)
        default:
            for fi15 := 0; fi15 < 7; fi15++ {
                v2 = (v2 | 127)
                v16 := GenS1{F0: 921}
                _ = fi15
                _ = v16
            }
        }
        v2 = (^(fi14))
        _ = fi14
    }
    switch v2 {
    case -780:
        var v17 uint8 = uint8(v4)
        _ = v17
    case -6:
        v18 := GenS2{F0: v2}
        _ = v18
    case 32768:
        if ((v1 + (^(v1))) <= 2) {
            fmt.Println(v5, v2)
            v19 := GenS3{F0: (v1 - v1), F1: 1, F2: (v3 - v3)}
            v1 -= 763
            _ = v19
        } else {
            for fi20 := 0; fi20 < 17; fi20++ {
                v4 -= v4
                switch fi20 {
                case 32767:
                    v21 := GenS1{F0: 42}
                    _ = v21
                case 726:
                    switch "" {
                    case "foo bar":
                        v22 := GenS1{F0: 1}
                        _ = v22
                    default:
                        if (9223372036854775807 >= 0) {
                            v1 -= (v1 >> ((v4) & 63))
                            v5 += (v5 << (((v4 / ((v4) | 1))) & 63))
                        } else {
                            switch "" {
                            case "a":
                                v2 = fi20
                            }
                            switch (v2 * (fi20 - (fi20 << (((v4 ^ v4)) & 63)))) {
                            case 1:
                                var v23 uint32 = (v1 | (v1 >> (((^(v4))) & 63)))
                                _ = v23
                            }
                        }
                    }
                case -734:
                    var v24 uint16 = 0
                    _ = v24
                default:
                    for fi25 := 0; fi25 < 12; fi25++ {
                        v26 := GenS4{F0: ((true && ((-1 > 0) && true)) || true), F1: 32767}
                        _ = fi25
                        _ = v26
                    }
                }
                v27 := GenS5{F0: 1, F1: v1, F2: 65535}
                _ = fi20
                _ = v27
            }
            v2 = (v2 - v2)
        }
    }
    v28 := GenS2{F0: (v2 / (((v2 >> (((v4 + v4)) & 63))) | 1))}
    v3 += (v3 >> ((v4) & 63))
    v2 = (v2 * v2)
    if false {
        v29 := []int{-217}
        fmt.Println(v1, v2, v1)
        fmt.Println(v3)
        _ = v29
    }
    var v30 int16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
