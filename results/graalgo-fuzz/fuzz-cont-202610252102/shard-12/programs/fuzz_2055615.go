package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 uint64 = 42
    var v2 string = "hello"
    var v3 string = "hello"
    v4 := make(map[uint32]int)
    v5 := make(map[string]string)
    switch v3 {
    case "hello":
        if true {
            fmt.Println(v1, v2, v3)
            if (!(!(!false))) {
                for fi6 := 0; fi6 < 4; fi6++ {
                    v1 = (v1 | 42)
                    _ = fi6
                }
                v3 = ("" + v3)
            }
        }
    default:
        switch 2147483647 {
        case -9223372036854775808:
            var v7 int16 = 0
            _ = v7
        case 128:
            v8 := []int{-2147483648, -41}
            _ = v8
        case 834:
            v9 := GenS1{F0: "foo bar", F1: 1}
            _ = v9
        }
    }
    switch ("" + v3) {
    case "hello":
        v10 := GenS1{F0: v3, F1: -1}
        _ = v10
    }
    var v11 float32 = (((float32(((0.1 * float64(-2.25)) * -2.25)) / ((((-2.25 - (0.0 - 1e10)) + 0.0) - (1e10 - -2.25)) / 1.0)) * 0.1) / float32((0.0 - 0.1)))
    v12 := GenS1{F0: ("hello" + v2), F1: 32767}
    v11 -= v11
    v3 = v3
    var v13 int64 = 1
    for fi14 := 0; fi14 < 15; fi14++ {
        v15 := GenS1{F0: "a", F1: 619}
        switch v3 {
        case "foo bar":
            if ((!(246 < 1)) || (1.5 < 3.14159)) {
                _ = len(v5)
            } else {
                v13 = (v13 | v13)
                v11 = 0.0
                v4[672] = fi14
            }
        default:
            v11 -= (float32(0.1) / v11)
        }
        _ = fi14
        _ = v15
    }
    if ("\n\t" <= v3) {
        for fi16 := 0; fi16 < 18; fi16++ {
            switch "a" {
            case "foo bar":
                var v17 uint8 = 145
                _ = v17
            case "":
                switch (fi16 << ((2) & 63)) {
                case -397:
                    for fi18 := 0; fi18 < 14; fi18++ {
                        v11 -= v11
                        _ = fi18
                    }
                case 613:
                    var v19 int16 = -32768
                    _ = v19
                case -523:
                    for fi20 := 0; fi20 < 16; fi20++ {
                        if (42 == -1) {
                            v11 = (v11 * (v11 * v11))
                            if (true || true) {
                                v11 = v11
                            } else {
                                v1 -= 42
                            }
                        }
                        _ = fi20
                    }
                }
            }
            _ = fi16
        }
        if (!false) {
            switch "a" {
            case "\n\t":
                var v21 uint16 = 0
                _ = v21
            default:
                switch v3 {
                case "a":
                    var v22 int64 = v13
                    _ = v22
                case "\n\t":
                    v23 := GenS1{F0: (v3 + "hello"), F1: 1}
                    _ = v23
                case "hello":
                    var v24 bool = true
                    _ = v24
                default:
                    var v25 float32 = (v11 * v11)
                    _ = v25
                }
            }
            v26 := GenS2{F0: (v13 << ((252) & 63))}
            _ = v26
        } else {
            switch -1 {
            case 661:
                v4[635] = 2147483647
            default:
                if true {
                    var v27 uint64 = (v1 / (((v1 >> ((65535) & 63))) | 1))
                    _ = v27
                }
            }
            switch 128 {
            case 32767:
                fmt.Println(v11, v1, v13)
            default:
                v28 := GenS1{F0: "", F1: 42}
                _ = v28
            }
            var v29 float32 = (1e10 / v11)
            _ = v29
        }
    } else {
        if true {
            v30 := []int{}
            _ = v30
        } else {
            if (1 == 127) {
                for fi31 := 0; fi31 < 16; fi31++ {
                    fmt.Println(v13)
                    _ = fi31
                }
                v1 += v1
            }
        }
    }
    for fi32 := 0; fi32 < 2; fi32++ {
        if true {
            if ((2 == (v1 - v1)) || ((1e10 + v11) == (float32(-2.25) / float32((1e10 + 1e10))))) {
                if false {
                    v33 := GenS2{F0: (-(v13))}
                    var v34 int32 = -177
                    _ = v33
                    _ = v34
                }
                fmt.Println(v11)
                var v35 uint16 = 2
                _ = v35
            }
            delete(v5, ("\n\t" + "\n\t"))
        } else {
            var v36 uint = 594
            delete(v4, 65535)
            _ = v36
        }
        var v37 int = int(v13)
        switch v2 {
        case "":
            var v38 float32 = (1.5 * v11)
            _ = v38
        }
        _ = fi32
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
