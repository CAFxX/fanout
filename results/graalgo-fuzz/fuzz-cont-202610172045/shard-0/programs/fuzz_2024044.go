package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 uint16; F1 uint64; F2 bool }
type GenS5 struct { F0 uint8; F1 int64 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]bool)
    v3 := make(map[string]int)
    v4 := make(map[string]string)
    var v5 int = -1
    if (("foo bar" <= "hello") || (!true)) {
        v6 := []int{v5, 128}
        _ = v6
    }
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 3; fi8++ {
            if true {
                v5 = (-(fi7))
            }
            v5 -= v5
            _ = fi8
        }
        _ = fi7
    }
    v9 := GenS1{F0: 255}
    v5 += 9223372036854775807
    v10 := GenS2{F0: 0}
    v11 := GenS3{F0: (v5 | -821), F1: (v5 - v5)}
    if false {
        v5 += (v5 >> ((1) & 63))
        v5 = 4294967295
        v12 := GenS1{F0: 1}
        _ = v12
    } else {
        v5 -= (v5 + -974)
        if false {
            fmt.Println(v5)
        } else {
            switch "\n\t" {
            case "foo bar":
                switch "hello" {
                case "hello":
                    var v13 float32 = -2.25
                    _ = v13
                case "a":
                    if (!false) {
                        switch (v5 ^ (v5 | 2147483647)) {
                        case 84:
                            var v14 float32 = ((float32(((0.1 / (0.0 + 0.1)) * (0.1 / (1.5 - (0.1 - 0.0))))) / 1e10) + -2.25)
                            _ = v14
                        }
                        for fi15 := 0; fi15 < 9; fi15++ {
                            if false {
                                v16 := GenS1{F0: 684}
                                _ = v16
                            }
                            _ = fi15
                        }
                    }
                case "foo bar":
                    fmt.Println(v5)
                default:
                    var v17 uint8 = 128
                    _ = v17
                }
            }
            var v18 string = "\n\t"
            _ = v18
        }
    }
    switch "a" {
    case "hello":
        var v19 int64 = 847
        _ = v19
    case "":
        for fi20 := 0; fi20 < 11; fi20++ {
            var v21 uint64 = 1
            switch "\n\t" {
            case "a":
                for fi22 := 0; fi22 < 6; fi22++ {
                    v23 := v1[2147483647]
                    _ = fi22
                    _ = v23
                }
            default:
                v24 := GenS4{F0: 653, F1: v21, F2: (!false)}
                _ = v24
            }
            v25 := GenS5{F0: uint8(v21), F1: -9223372036854775808}
            _ = fi20
            _ = v21
            _ = v25
        }
    case "\n\t":
        v26 := []int{}
        _ = v26
    default:
        v27 := v2[("a" + "a")]
        _ = v27
    }
    switch 32767 {
    case 524:
        fmt.Println(v5)
    default:
        if true {
            switch ("\n\t" + ("" + "")) {
            case "":
                fmt.Println(v5)
            default:
                var v28 int16 = -419
                _ = v28
            }
            v5 -= 866
        }
    }
    for fi29 := 0; fi29 < 13; fi29++ {
        fmt.Println(fi29)
        _ = fi29
    }
    var v30 int8 = -128
    var v31 uint8 = 128
    v31 += v31
    v2["a"] = true
    if (((!(0.1 >= ((3.14159 * (1.5 + float32(0.0))) * (0.0 / 3.14159)))) && false) || false) {
        for fi32 := 0; fi32 < 2; fi32++ {
            v5 = (fi32 + v5)
            v31 -= (v31 << ((255) & 63))
            var v33 int16 = 0
            _ = fi32
            _ = v33
        }
        for fi34 := 0; fi34 < 4; fi34++ {
            for fi35 := 0; fi35 < 4; fi35++ {
                v30 = (v30 & 6)
                _ = fi35
            }
            _ = fi34
        }
    } else {
        var v36 uint = 65535
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v9
}
