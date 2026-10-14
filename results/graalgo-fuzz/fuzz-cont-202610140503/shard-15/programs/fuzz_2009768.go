package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 int }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]string)
    var v3 int16 = 717
    v4 := make(map[string]bool)
    var v5 string = "a"
    v5 = v5
    var v6 uint64 = 370
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 uint16 = 498
        _ = fi7
        _ = v8
    }
    if ((true || (0 == -1)) && (42 < (v3 - v3))) {
        v3 += v3
    }
    v6 = 2
    v9 := GenS1{F0: 163}
    if false {
        var v10 string = "\n\t"
        v10 = (v10 + v10)
        _ = v10
    } else {
        var v11 int = -117
        v12 := v1[2147483647]
        _ = v11
        _ = v12
    }
    v13 := GenS2{F0: v3, F1: -7}
    fmt.Println(v3, v5, v6)
    var v14 uint64 = (v6 & v6)
    var v15 uint8 = 42
    if false {
        var v16 float64 = 1e10
        v3 -= v3
        fmt.Println(v5)
        _ = v16
    } else {
        var v17 uint32 = 42
        switch "" {
        case "\n\t":
            for fi18 := 0; fi18 < 3; fi18++ {
                v5 = "foo bar"
                v3 = (v3 >> ((48) & 63))
                var v19 uint32 = 871
                _ = fi18
                _ = v19
            }
        case "hello":
            for fi20 := 0; fi20 < 5; fi20++ {
                var v21 int16 = (^(v3))
                _ = fi20
                _ = v21
            }
        case "foo bar":
            var v22 int16 = 32767
            _ = v22
        default:
            fmt.Println(v14)
        }
        _ = v17
    }
    if (447 <= 1) {
        var v23 int8 = 0
        _ = v23
    } else {
        v24 := []int{}
        _ = v24
    }
    switch 4294967295 {
    case 9223372036854775807:
        v3 -= v3
    }
    switch 788 {
    case 32767:
        switch v5 {
        case "\n\t":
            if (42 == 42) {
                v6 += v14
            } else {
                v25 := []int{-260, 32768, 1}
                v3 = v3
                if len(v25) >= 2 {
                    v25 = v25[1:len(v25)]
                }
                _ = v25
            }
        case "hello":
            v26 := []int{9223372036854775807, 128}
            _ = v26
        case "a":
            v6 = (^(v14))
        default:
            v14 += v14
        }
    case -348:
        v15 = v15
    default:
        v3 += -721
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
