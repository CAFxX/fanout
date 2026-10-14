package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 int64; F1 int64; F2 uint }
type GenS4 struct { F0 int64 }

func main() {
    var v1 int = 127
    v2 := make(map[int]bool)
    var v3 int32 = 0
    v4 := []string{"foo bar", (((("\n\t" + "foo bar") + ("hello" + "foo bar")) + (("foo bar" + "") + "\n\t")) + ("hello" + "hello")), ("a" + "hello"), ("a" + "")}
    var v5 int16 = 0
    v6 := make(map[uint32]uint32)
    var v7 uint16 = 0
    switch ("" + "") {
    case "":
        fmt.Println(v5, v5, v1)
    case "foo bar":
        for fi8 := 0; fi8 < 16; fi8++ {
            for fi9 := 0; fi9 < 11; fi9++ {
                fmt.Println(v5)
                switch 2 {
                case -124:
                    v7 -= v7
                case 127:
                    v10 := GenS1{F0: v1, F1: 9223372036854775807, F2: (fi8 | fi8)}
                    _ = v10
                case 4294967295:
                    if false {
                        if ((0.1 - 3.14159) <= 3.14159) {
                            v7 = (^(v7))
                        }
                        v5 -= (v5 / ((1000) | 1))
                        v1 = fi8
                    } else {
                        if true {
                            v11 := GenS1{F0: (-(fi8)), F1: int64(fi9), F2: fi8}
                            _ = v11
                        }
                    }
                }
                v12 := GenS2{F0: (v7 * v7), F1: 683}
                _ = fi9
                _ = v12
            }
            v3 += int32(v5)
            _ = fi8
        }
    case "\n\t":
        v3 = -1
    default:
        v5 = (v5 % ((1000) | 1))
    }
    v13 := GenS2{F0: v7, F1: 42}
    for fi14 := 0; fi14 < 8; fi14++ {
        if (false || false) {
            var v15 int32 = (v3 & v3)
            for fi16 := 0; fi16 < 11; fi16++ {
                delete(v2, v1)
                switch fi14 {
                case -760:
                    var v17 int64 = 0
                    _ = v17
                default:
                    v1 -= (-(fi14))
                }
                _ = fi16
            }
            _ = v15
        } else {
            _ = len(v4)
            switch ("" + "\n\t") {
            case "a":
                v1 += v1
            default:
                for fi18 := 0; fi18 < 12; fi18++ {
                    v5 += v5
                    _ = fi18
                }
            }
        }
        fmt.Println(v1, v3, v7)
        _ = fi14
    }
    v19 := GenS3{F0: 0, F1: -1, F2: 79}
    for fi20 := 0; fi20 < 14; fi20++ {
        if false {
            var v21 int16 = v5
            for fi22 := 0; fi22 < 18; fi22++ {
                fmt.Println(v3)
                v21 = 1
                _ = fi22
            }
            v23 := GenS4{F0: -395}
            _ = v21
            _ = v23
        } else {
            var v24 uint16 = 42
            _ = v24
        }
        var v25 uint64 = 2
        _ = fi20
        _ = v25
    }
    if true {
        v3 = (v3 >> ((4294967295) & 63))
    }
    v7 -= v7
    var v26 int = (v1 << ((65535) & 63))
    var v27 uint = 64
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
