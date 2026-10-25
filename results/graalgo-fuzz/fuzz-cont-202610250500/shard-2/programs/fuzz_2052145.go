package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 uint16 }

func main() {
    var v1 uint16 = 2
    v2 := make(map[uint32]bool)
    v3 := []uint32{223, 2, 0, 2147483647}
    var v4 uint = 534
    for fi5 := 0; fi5 < 19; fi5++ {
        v6 := v2[0]
        for fi7 := 0; fi7 < 2; fi7++ {
            var v8 int16 = 42
            var v9 int64 = 551
            v9 = v9
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi5
        _ = v6
    }
    v1 -= v1
    switch 34 {
    case 32767:
        v4 = v4
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        fmt.Println(v1)
        for fi11 := 0; fi11 < 16; fi11++ {
            if len(v3) > 0 {
                v3[0] = 1
            }
            _ = fi11
        }
        var v12 int16 = 0
        _ = fi10
        _ = v12
    }
    var v13 uint16 = v1
    if len(v3) > 0 {
        v3[0] = 0
    }
    for fi14 := 0; fi14 < 18; fi14++ {
        var v15 float64 = 1.5
        v4 += (^(v4))
        _ = fi14
        _ = v15
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v13 -= 0
    if false {
        v16 := GenS1{F0: (v4 / (((v4 << ((18446744073709551615) & 63))) | 1))}
        delete(v2, 668)
        _ = v16
    }
    v4 = v4
    if (1 > -1) {
        switch "hello" {
        case "hello":
            v1 = (^(v13))
        case "":
            var v17 uint16 = 0
            _ = v17
        case "foo bar":
            for fi18 := 0; fi18 < 5; fi18++ {
                switch "foo bar" {
                case "\n\t":
                    v1 += v1
                case "foo bar":
                    v13 = v1
                default:
                    v1 = (v1 % ((v13) | 1))
                }
                v13 -= v13
                _ = fi18
            }
        default:
            if true {
                var v19 uint = v4
                v4 = v19
                _ = v19
            } else {
                v13 = (v13 & uint16(v4))
                fmt.Println(v1, v4, v13)
            }
        }
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        if (uint16(v4) > v1) {
            v13 += 40
            var v21 bool = (!false)
            v22 := GenS1{F0: (v4 << ((v4) & 63))}
            _ = v21
            _ = v22
        }
        v23 := GenS2{F0: 255, F1: v1}
        _ = fi20
        _ = v23
    }
    var v24 int8 = -128
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}
