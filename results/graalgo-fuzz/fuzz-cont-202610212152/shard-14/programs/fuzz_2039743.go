package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 string; F1 int16; F2 int64 }
type GenS4 struct { F0 uint16; F1 uint64; F2 uint }
type GenS5 struct { F0 uint32; F1 uint32; F2 int32 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 bool }
func genFn0() uint64 {
    switch 551 {
    case 298:
        for fi4 := 0; fi4 < 7; fi4++ {
            fmt.Println(fi4)
            _ = fi4
        }
    case -730:
        v5 := GenS1{F0: 813, F1: 404, F2: 377}
        _ = v5
    default:
        for fi6 := 0; fi6 < 15; fi6++ {
            switch "a" {
            case "foo bar":
                var v7 int64 = -9223372036854775808
                _ = v7
            case "a":
                var v8 uint64 = 629
                _ = v8
            case "hello":
                var v9 int16 = -1
                _ = v9
            }
            var v10 bool = (!true)
            _ = fi6
            _ = v10
        }
    }
    var v11 uint64 = 110
    if (false || (0 > -2147483648)) {
        for fi12 := 0; fi12 < 14; fi12++ {
            if false {
                for fi13 := 0; fi13 < 13; fi13++ {
                    var v14 uint64 = 2
                    _ = fi13
                    _ = v14
                }
                switch fi12 {
                case -1:
                    v11 = v11
                case 128:
                    var v15 uint8 = 244
                    _ = v15
                default:
                    if true {
                        v11 += 651
                        v16 := []int{}
                        _ = v16
                    }
                }
            } else {
                v11 -= (v11 % ((v11) | 1))
                var v17 uint8 = 113
                _ = v17
            }
            for fi18 := 0; fi18 < 12; fi18++ {
                var v19 uint64 = v11
                fmt.Println(fi12, v19, fi12)
                _ = fi18
                _ = v19
            }
            _ = fi12
        }
        switch "a" {
        case "\n\t":
            v11 = (^(v11))
        default:
            v11 = 234
        }
        fmt.Println(v11)
    } else {
        if ("" <= "foo bar") {
            for fi20 := 0; fi20 < 4; fi20++ {
                var v21 int8 = 110
                v11 = v11
                _ = fi20
                _ = v21
            }
        }
        var v22 uint = 40
        v23 := GenS2{F0: -1, F1: (v22 - (^(v22)))}
        _ = v22
        _ = v23
    }
    v24 := []int{}
    _ = v11
    _ = v24
    return 1
}

func main() {
    var v1 int = 128
    var v2 uint32 = 1
    v3 := []uint32{}
    var v25 int64 = 42
    v3 = append(v3, (v2 >> ((688) & 63)))
    v1 = (-(v1))
    if (!((v25 << ((2) & 63)) <= v25)) {
        for fi26 := 0; fi26 < 16; fi26++ {
            v27 := GenS3{F0: "a", F1: int16(v2), F2: v25}
            _ = fi26
            _ = v27
        }
        var v28 uint = 42
        var v29 int16 = 42
        _ = v28
        _ = v29
    } else {
        v1 -= (v1 ^ -85)
        var v30 float32 = ((((-2.25 / -2.25) * (0.1 + (3.14159 * ((0.0 + 0.1) / 1e10)))) / 1.5) * (1e10 - 1.5))
        _ = v30
    }
    for fi31 := 0; fi31 < 7; fi31++ {
        var v32 int16 = -32768
        v33 := GenS2{F0: 0, F1: 255}
        switch "\n\t" {
        case "hello":
            switch ("foo bar" + "foo bar") {
            case "foo bar":
                for fi34 := 0; fi34 < 7; fi34++ {
                    switch -9223372036854775808 {
                    case 32767:
                        v3 = append(v3, v2)
                    case -1:
                        var v35 uint32 = (v2 | v2)
                        _ = v35
                    }
                    _ = fi34
                }
            }
        case "\n\t":
            if true {
                v36 := GenS4{F0: 255, F1: 2, F2: 4294967295}
                v1 -= (v1 * v1)
                _ = v36
            }
        default:
            v37 := GenS5{F0: v2, F1: v2, F2: -1}
            _ = v37
        }
        _ = fi31
        _ = v32
        _ = v33
    }
    var v38 float64 = -2.25
    if true {
        switch v1 {
        case 78:
            var v39 uint32 = v2
            _ = v39
        case 508:
            var v40 uint16 = 42
            _ = v40
        case -337:
            v25 -= v25
        }
        _ = len(v3)
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v41 int = (v1 + (-(v1)))
    v1 = v41
    v38 = 1e10
    v41 = (v41 / ((v41) | 1))
    v38 = (v38 + ((((v38 + 1.5) - v38) / -2.25) - v38))
    switch "" {
    case "foo bar":
        if (0 < 4294967295) {
            switch "a" {
            case "hello":
                if ((true && (!((^(v41)) == (^(v41))))) && (601 < 4294967295)) {
                    v25 -= v25
                    for fi42 := 0; fi42 < 1; fi42++ {
                        fmt.Println(v38, v1)
                        _ = fi42
                    }
                    _ = len(v3)
                } else {
                    if len(v3) > 0 {
                        v3[0] = v2
                    }
                    v41 += (v1 & (^(v1)))
                }
            default:
                switch v1 {
                case 304:
                    var v43 uint16 = 65535
                    _ = v43
                case 2147483647:
                    v44 := GenS1{F0: 65535, F1: 42, F2: 2}
                    _ = v44
                }
            }
        }
    case "":
        v45 := GenS6{F0: 4294967295}
        _ = v45
    case "hello":
        v1 = 32768
    }
    v38 = v38
    v25 = 9223372036854775807
    for fi46 := 0; fi46 < 11; fi46++ {
        for fi47 := 0; fi47 < 5; fi47++ {
            switch "foo bar" {
            case "":
                fmt.Println(v25, v1)
            }
            fmt.Println(v38, v2, fi46)
            var v48 float64 = v38
            _ = fi47
            _ = v48
        }
        v49 := GenS7{F0: true}
        for fi50 := 0; fi50 < 16; fi50++ {
            v1 = fi50
            v51 := GenS4{F0: 215, F1: 348, F2: 624}
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi50
            _ = v51
        }
        _ = fi46
        _ = v49
    }
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v38
    _ = v41
}
