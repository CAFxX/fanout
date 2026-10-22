package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 bool; F1 string }

func main() {
    var v1 bool = (1.5 <= -2.25)
    v2 := []uint32{}
    var v3 int16 = 1
    var v4 uint32 = 1
    var v5 uint16 = 869
    var v6 int16 = 1000
    v4 -= (v4 >> ((2) & 63))
    v7 := GenS1{F0: (v6 ^ v6)}
    v3 -= (v3 | v3)
    for fi8 := 0; fi8 < 3; fi8++ {
        v5 = 923
        var v9 bool = false
        _ = len(v2)
        _ = fi8
        _ = v9
    }
    v2 = append(v2, 65535)
    v5 = (^(v5))
    if (18446744073709551615 != 4294967295) {
        if v1 {
            v3 -= 1000
        }
    }
    var v10 bool = v1
    v11 := GenS1{F0: (-(v6))}
    var v12 float32 = 1.5
    var v13 string = "a"
    for fi14 := 0; fi14 < 16; fi14++ {
        if v10 {
            v15 := GenS1{F0: 1}
            if v10 {
                var v16 int = fi14
                _ = v16
            }
            if v1 {
                v17 := GenS2{F0: -250}
                fmt.Println(v6)
                _ = v17
            } else {
                v2 = append(v2, v4)
            }
            _ = v15
        }
        _ = fi14
    }
    for fi18 := 0; fi18 < 11; fi18++ {
        switch v13 {
        case "\n\t":
            v19 := GenS2{F0: -767}
            _ = v19
        case "a":
            var v20 uint16 = 65535
            _ = v20
        default:
            switch fi18 {
            case -812:
                v3 = -1
            case 287:
                fmt.Println(v6, v3, v5)
            default:
                v4 += v4
            }
        }
        _ = fi18
    }
    for fi21 := 0; fi21 < 6; fi21++ {
        for fi22 := 0; fi22 < 14; fi22++ {
            fmt.Println(v10)
            var v23 int8 = 42
            _ = fi22
            _ = v23
        }
        v6 += 1000
        for fi24 := 0; fi24 < 20; fi24++ {
            var v25 uint32 = 42
            switch "a" {
            case "foo bar":
                fmt.Println(v5, fi21)
            case "a":
                switch v13 {
                case "foo bar":
                    switch 9223372036854775807 {
                    case 127:
                        v13 = "a"
                    case 42:
                        v4 = (v4 * (v25 / ((895) | 1)))
                    case -271:
                        v4 += (^(v4))
                    }
                    v5 += uint16(v4)
                    var v26 int = (-(fi21))
                    _ = v26
                case "hello":
                    v27 := GenS1{F0: v6}
                    _ = v27
                }
            case "hello":
                v28 := GenS3{F0: uint64(v5), F1: -2147483648}
                _ = v28
            }
            var v29 uint = 54
            _ = len(v2)
            _ = fi24
            _ = v25
            _ = v29
        }
        _ = fi21
    }
    if true {
        v30 := GenS2{F0: 0}
        fmt.Println(v12, v12)
        var v31 bool = (v13 != ("foo bar" + "hello"))
        _ = v30
        _ = v31
    } else {
        for fi32 := 0; fi32 < 6; fi32++ {
            v3 -= (v3 % ((v6) | 1))
            _ = fi32
        }
        if true {
            v1 = (-1 > int64(v4))
        } else {
            v33 := GenS4{F0: 42}
            _ = v33
        }
        v34 := GenS5{F0: v1, F1: "foo bar"}
        _ = v34
    }
    var v35 int16 = -32768
    v35 = (v3 - v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
