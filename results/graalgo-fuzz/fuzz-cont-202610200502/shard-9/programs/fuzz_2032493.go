package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 uint8 = 128
    var v2 string = "foo bar"
    var v3 uint = 165
    v4 := make(map[string]string)
    switch -2147483648 {
    case 137:
        v3 = v3
    case 42:
        var v5 float64 = 0.1
        _ = v5
    default:
        var v6 float64 = 3.14159
        _ = v6
    }
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := GenS1{F0: 764, F1: (v1 * (v1 + (v1 >> (((v3 / (((v3 ^ (v3 | (^(v3))))) | 1))) & 63))))}
        v9 := []int{987, -306}
        fmt.Println(v1, v1, v2)
        _ = fi7
        _ = v8
        _ = v9
    }
    if true {
        switch v2 {
        case "a":
            v10 := []int{-618, 127}
            _ = v10
        case "\n\t":
            v3 += (v3 ^ (v3 - (v3 + 65535)))
        }
    }
    v1 = 162
    v11 := GenS1{F0: 667, F1: (v1 << (((v3 / ((42) | 1))) & 63))}
    v12 := GenS1{F0: uint64(v1), F1: v1}
    switch -58 {
    case 9223372036854775807:
        if true {
            v13 := []int{-2147483648, 0, -1}
            v4[v2] = "hello"
            _ = v13
        }
    default:
        for fi14 := 0; fi14 < 5; fi14++ {
            if true {
                switch ((v2 + v2) + "hello") {
                case "foo bar":
                    var v15 uint64 = 2
                    _ = v15
                case "":
                    for fi16 := 0; fi16 < 7; fi16++ {
                        v17 := GenS1{F0: 42, F1: (v1 - v1)}
                        var v18 string = v2
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                }
                v1 += v1
                for fi19 := 0; fi19 < 18; fi19++ {
                    switch "\n\t" {
                    case "\n\t":
                        delete(v4, "\n\t")
                    default:
                        var v20 int64 = -1
                        _ = v20
                    }
                    v21 := GenS2{F0: -9223372036854775808}
                    _ = fi19
                    _ = v21
                }
            }
            v22 := GenS2{F0: 0}
            _ = fi14
            _ = v22
        }
    }
    var v23 uint8 = v1
    var v24 uint64 = 951
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
