package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 uint16; F2 bool }
type GenS4 struct { F0 int64; F1 uint; F2 uint64 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]string)
    v3 := make(map[string]bool)
    var v4 bool = (0 >= 0)
    v4 = (4294967295 != 2)
    for fi5 := 0; fi5 < 14; fi5++ {
        fmt.Println(v4)
        if v4 {
            for fi6 := 0; fi6 < 14; fi6++ {
                v4 = (!false)
                _ = fi6
            }
        }
        v7 := GenS1{F0: -1, F1: 0}
        _ = fi5
        _ = v7
    }
    var v8 uint16 = 0
    fmt.Println(v4, v8)
    switch "foo bar" {
    case "foo bar":
        switch -794 {
        case -681:
            v9 := GenS1{F0: -862, F1: 255}
            _ = v9
        case 9223372036854775807:
            for fi10 := 0; fi10 < 2; fi10++ {
                v2[(fi10 | -1)] = "foo bar"
                for fi11 := 0; fi11 < 3; fi11++ {
                    _ = len(v2)
                    v4 = (18446744073709551615 <= 4294967295)
                    _ = fi11
                }
                _ = fi10
            }
        case -851:
            delete(v2, 0)
        default:
            v4 = v4
        }
    case "\n\t":
        v12 := GenS2{F0: 9223372036854775807}
        _ = v12
    }
    var v13 int8 = 42
    for fi14 := 0; fi14 < 9; fi14++ {
        v15 := GenS2{F0: -659}
        if ((((111 <= 0) || ((1.5 > -2.25) || false)) || ("a" == "\n\t")) && v4) {
            var v16 float64 = 0.0
            _ = v16
        }
        _ = fi14
        _ = v15
    }
    if v4 {
        fmt.Println(v13, v4, v4)
        if v4 {
            v17 := GenS3{F0: -1, F1: v8, F2: v4}
            delete(v3, ("a" + ""))
            _ = v17
        } else {
            switch 4294967295 {
            case 128:
                fmt.Println(v4)
            case 2147483647:
                var v18 int64 = -25
                _ = v18
            default:
                _ = len(v3)
            }
            var v19 uint64 = 0
            _ = v19
        }
        for fi20 := 0; fi20 < 15; fi20++ {
            for fi21 := 0; fi21 < 13; fi21++ {
                v22 := GenS1{F0: 42, F1: 122}
                _ = fi21
                _ = v22
            }
            v23 := []int{1, (-(fi20)), 32768, (fi20 | 32768)}
            _ = fi20
            _ = v23
        }
    } else {
        v8 = (v8 % ((v8) | 1))
        switch (((("" + "\n\t") + "") + "") + "") {
        case "a":
            v13 = v13
        default:
            v24 := GenS4{F0: 0, F1: 0, F2: 565}
            _ = v24
        }
    }
    fmt.Println(v13, v13)
    v13 = v13
    v25 := GenS3{F0: 60, F1: v8, F2: v4}
    v26 := GenS2{F0: 1}
    var v27 int = 42
    v28 := []int{v27, (v27 - (-(v27))), v27}
    v13 = -116
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v8
}
