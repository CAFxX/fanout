package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 bool; F1 bool }

func main() {
    var v1 int8 = -45
    v2 := []string{"", ("hello" + "a")}
    v3 := make(map[uint32]string)
    var v4 float64 = (-2.25 - 0.0)
    v5 := make(map[string]string)
    v6 := make(map[uint32]uint32)
    v7 := GenS1{F0: 255}
    delete(v6, 2147483647)
    if true {
        fmt.Println(v1, v1)
        for fi8 := 0; fi8 < 9; fi8++ {
            if (false && true) {
                fmt.Println(fi8, v4)
                if false {
                    v9 := GenS1{F0: 1}
                    _ = v9
                } else {
                    v10 := v5[""]
                    v4 = (1e10 * 0.0)
                    var v11 uint8 = 22
                    _ = v10
                    _ = v11
                }
                _ = len(v6)
            }
            _ = fi8
        }
    } else {
        for fi12 := 0; fi12 < 7; fi12++ {
            if false {
                var v13 uint = 2
                _ = v13
            }
            _ = fi12
        }
        v14 := GenS2{F0: v1, F1: 434}
        v1 -= (v1 / ((v1) | 1))
        _ = v14
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        if true {
            v16 := GenS2{F0: (-(v1)), F1: 2}
            _ = v16
        }
        var v17 uint32 = 852
        _ = fi15
        _ = v17
    }
    var v18 int = 4294967295
    v19 := GenS1{F0: 0}
    switch v18 {
    case 193:
        var v20 int16 = -119
        _ = v20
    case 1:
        v2 = append(v2, ("" + "foo bar"))
    case -2147483648:
        switch "a" {
        case "\n\t":
            v1 = v1
        case "foo bar":
            v21 := GenS2{F0: v1, F1: 2}
            _ = v21
        default:
            switch "foo bar" {
            case "foo bar":
                delete(v6, 859)
            case "":
                v4 -= (v4 + (float64(1e10) / (1e10 / float64(3.14159))))
            case "hello":
                if len(v2) > 0 {
                    v2[0] = "hello"
                }
            default:
                var v22 uint16 = 2
                _ = v22
            }
        }
    default:
        if true {
            switch 477 {
            case -938:
                var v23 uint = 255
                _ = v23
            case -2147483648:
                for fi24 := 0; fi24 < 18; fi24++ {
                    for fi25 := 0; fi25 < 13; fi25++ {
                        v1 += (v1 >> ((0) & 63))
                        _ = fi25
                    }
                    fmt.Println(v18, v18)
                    _ = fi24
                }
            case 886:
                var v26 int16 = 0
                _ = v26
            default:
                v27 := GenS3{F0: true, F1: (255 <= 0)}
                _ = v27
            }
        }
    }
    v18 = (v18 | (^(v18)))
    var v28 uint64 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
