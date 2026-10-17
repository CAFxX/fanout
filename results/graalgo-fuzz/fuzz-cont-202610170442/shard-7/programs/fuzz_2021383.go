package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 int64; F1 int }

func main() {
    var v1 uint16 = 2
    v2 := make(map[int]bool)
    v3 := make(map[int]uint32)
    var v4 bool = true
    v5 := []uint32{0, 2, 409}
    fmt.Println(v1)
    var v6 uint = 18446744073709551615
    var v7 uint16 = 1
    for fi8 := 0; fi8 < 19; fi8++ {
        for fi9 := 0; fi9 < 7; fi9++ {
            var v10 int8 = -1
            if false {
                fmt.Println(v4, v7)
                v5 = append(v5, 75)
            } else {
                if ((("foo bar" + "a") + "foo bar") < "") {
                    if len(v5) > 0 {
                        v5[0] = 420
                    }
                    fmt.Println(v4)
                    fmt.Println(v1, v1, v4)
                }
            }
            _ = fi9
            _ = v10
        }
        var v11 uint16 = v1
        _ = fi8
        _ = v11
    }
    v7 -= 42
    _ = len(v3)
    for fi12 := 0; fi12 < 1; fi12++ {
        fmt.Println(v4, v7, v4)
        _ = fi12
    }
    if (133 == 42) {
        v7 = 42
    }
    v6 -= (v6 * v6)
    v7 = (v7 << ((v6) & 63))
    _ = len(v2)
    v13 := GenS1{F0: 448, F1: -899}
    for fi14 := 0; fi14 < 11; fi14++ {
        var v15 uint32 = 1
        switch 9223372036854775807 {
        case 127:
            v7 = (v1 % ((65535) | 1))
        case 42:
            _ = len(v5)
        case 32768:
            v16 := GenS2{F0: ("\n\t" + ""), F1: (v6 & (v6 - v6))}
            _ = v16
        default:
            var v17 uint8 = 22
            _ = v17
        }
        for fi18 := 0; fi18 < 1; fi18++ {
            var v19 int16 = 1
            _ = fi18
            _ = v19
        }
        _ = fi14
        _ = v15
    }
    if (2 != 42) {
        for fi20 := 0; fi20 < 1; fi20++ {
            v4 = v4
            if v4 {
                v21 := GenS2{F0: "\n\t", F1: 18446744073709551615}
                switch ("foo bar" + "foo bar") {
                case "":
                    v4 = false
                case "a":
                    var v22 string = "\n\t"
                    _ = v22
                case "\n\t":
                    fmt.Println(v4)
                default:
                    v23 := GenS1{F0: 2, F1: 1}
                    _ = v23
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = v21
            } else {
                var v24 uint64 = 18446744073709551615
                _ = v24
            }
            _ = len(v2)
            _ = fi20
        }
    } else {
        v25 := GenS3{F0: 0, F1: 127}
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
