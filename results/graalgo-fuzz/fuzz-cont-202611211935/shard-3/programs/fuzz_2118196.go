package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 uint16; F1 int32; F2 string }

func main() {
    v1 := []string{"foo bar"}
    v2 := make(map[string]string)
    v3 := []int{42, -9223372036854775808, 128, -1}
    var v4 int16 = 666
    switch 42 {
    case -871:
        if false {
            v4 += v4
        }
    case 4294967295:
        switch 1 {
        case -2147483648:
            v5 := GenS1{F0: (!(true || false)), F1: 128}
            _ = v5
        case 1:
            v6 := GenS1{F0: true, F1: 1}
            _ = v6
        case 2147483647:
            v4 = v4
        }
    case 32768:
        if len(v3) > 0 {
            v3[0] = 1
        }
    default:
        var v7 float64 = -2.25
        _ = v7
    }
    v4 = 727
    v8 := GenS2{F0: 0, F1: 42, F2: "\n\t"}
    _ = len(v2)
    v4 -= (v4 * (v4 * (v4 * v4)))
    v4 = v4
    for fi9 := 0; fi9 < 18; fi9++ {
        if false {
            v4 = (v4 << ((1) & 63))
            if (("hello" + "\n\t") >= "hello") {
                switch fi9 {
                case -2147483648:
                    v2["hello"] = "\n\t"
                }
                switch "a" {
                case "hello":
                    v4 = (v4 - (v4 ^ v4))
                }
            } else {
                var v10 uint8 = 0
                _ = v10
            }
        }
        fmt.Println(fi9, fi9)
        _ = fi9
    }
    if (!true) {
        for fi11 := 0; fi11 < 16; fi11++ {
            var v12 int32 = 2147483647
            v13 := GenS1{F0: ((!false) || true), F1: 251}
            _ = fi11
            _ = v12
            _ = v13
        }
        v4 += (^(v4))
        for fi14 := 0; fi14 < 10; fi14++ {
            v4 = v4
            fmt.Println(fi14, fi14)
            _ = fi14
        }
    } else {
        v2[("hello" + "hello")] = ("a" + ("" + "hello"))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
