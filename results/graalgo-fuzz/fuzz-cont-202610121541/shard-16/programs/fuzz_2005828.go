package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 int; F2 string }
type GenS3 struct { F0 int64; F1 string; F2 bool }

func main() {
    var v1 uint = 853
    var v2 uint32 = 4294967295
    var v3 string = ""
    fmt.Println(v2, v1)
    if true {
        switch -432 {
        case -2147483648:
            if (280 > 2) {
                if false {
                    fmt.Println(v2, v3, v1)
                    v4 := GenS1{F0: -1}
                    if false {
                        v5 := GenS2{F0: 158, F1: -9223372036854775808, F2: v3}
                        _ = v5
                    } else {
                        var v6 uint = (^(v1))
                        _ = v6
                    }
                    _ = v4
                } else {
                    var v7 float64 = (0.1 - ((-2.25 / -2.25) + 1.5))
                    fmt.Println(v3)
                    v2 = v2
                    _ = v7
                }
            }
        case -135:
            var v8 uint16 = 0
            _ = v8
        case 1:
            v3 = ("\n\t" + (v3 + v3))
        default:
            v9 := GenS1{F0: 1}
            _ = v9
        }
    }
    v10 := []int{9223372036854775807, 32767, 42, -647}
    _ = len(v10)
    for fi11 := 0; fi11 < 13; fi11++ {
        var v12 uint = (v1 << ((v1) & 63))
        _ = fi11
        _ = v12
    }
    fmt.Println(v3, v1, v3)
    switch -72 {
    case 92:
        for fi13 := 0; fi13 < 9; fi13++ {
            switch "foo bar" {
            case "a":
                switch -585 {
                case 261:
                    var v14 uint8 = 128
                    _ = v14
                }
            case "hello":
                var v15 uint32 = (v2 & v2)
                _ = v15
            case "\n\t":
                fmt.Println(v1, fi13)
            }
            if true {
                v3 = (v3 + "")
                v1 = (^(v1))
            } else {
                var v16 uint32 = (v2 * (v2 + v2))
                _ = v16
            }
            _ = fi13
        }
    case -1:
        fmt.Println(v3, v3, v2)
    }
    var v17 int8 = 42
    v18 := GenS3{F0: -607, F1: v3, F2: true}
    v17 = v17
    v3 = "\n\t"
    fmt.Println(v2, v2)
    v17 = v17
    var v19 int8 = (v17 / ((34) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
