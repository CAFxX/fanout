package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 uint; F2 uint16 }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int16 = 1
    v3 := []string{"", "\n\t", "foo bar"}
    _ = len(v3)
    switch -2147483648 {
    case 9223372036854775807:
        for fi4 := 0; fi4 < 10; fi4++ {
            _ = len(v3)
            _ = fi4
        }
    case -1:
        if (0 == 2147483647) {
            switch ("foo bar" + "foo bar") {
            case "hello":
                _ = len(v3)
            }
            if (false && ((1.5 - -2.25) != (0.0 / 1e10))) {
                var v5 int8 = 80
                _ = v5
            } else {
                var v6 float64 = 0.1
                _ = v6
            }
        }
    default:
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
    }
    v7 := GenS1{F0: 267}
    if true {
        if false {
            v8 := GenS1{F0: 65535}
            _ = v8
        }
        _ = len(v3)
        v1 += v1
    }
    v2 = (v2 % ((v2) | 1))
    for fi9 := 0; fi9 < 6; fi9++ {
        v2 = (v2 / (((v2 / (((-(v2))) | 1))) | 1))
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
        v2 = v2
        _ = fi9
    }
    v2 = 571
    var v10 uint16 = 255
    v1 += -9223372036854775808
    var v11 float32 = 1.5
    if (0 <= int8(v1)) {
        fmt.Println(v10, v1, v10)
        for fi12 := 0; fi12 < 12; fi12++ {
            v11 = (1.5 + v11)
            _ = fi12
        }
        var v13 float64 = 1e10
        _ = v13
    }
    var v14 uint64 = 213
    v15 := GenS2{F0: v10, F1: 283, F2: 0}
    v16 := GenS2{F0: 1, F1: 2, F2: 223}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v7
}
